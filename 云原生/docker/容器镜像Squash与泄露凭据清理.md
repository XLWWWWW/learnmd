# 容器镜像 Squash、体积治理与泄露凭据清理

> 专题导航：[容器镜像文件系统专题索引](./容器镜像文件系统专题索引.md)

本文整理自 Ivan Velichko 的《Squash a Bloated Container Image to Remove Waste and Leaked Credentials》，并补充镜像层删除语义、安全响应、多阶段构建、BuildKit secrets、squash 的局限与验证方法。

## 1. 文章在解决什么问题

示例 Python 服务在构建过程中做了两件危险的事：

1. 安装 `build-essential`、`libpq-dev` 等编译依赖，产生了约 400 MB 的 layer。
2. 使用 `COPY` 把包含私有仓库凭据的 `pip.conf` 放进镜像。

后续 Dockerfile 又执行了：

```text
apt-get purge build-essential libpq-dev
rm /root/.config/.../pip.conf
```

从最终容器 rootfs 看，构建工具和凭据文件确实消失了。但是镜像仍然臃肿且不安全，因为容器镜像 layer 是不可变的增量变化集：

```text
Layer A：添加 build-essential、libpq-dev
Layer B：添加带 token 的 pip.conf
Layer C：删除 build-essential、libpq-dev
Layer D：删除 pip.conf
```

Layer C 和 Layer D 只能表达“在最终视图中隐藏下层文件”，不能修改或擦除 Layer A、Layer B 中已经保存的字节。

因此：

```text
最终 rootfs 中看不到文件
        !=
镜像历史 layers 中不存在文件
```

这篇 challenge 的目标就是把多个历史 layer 压平成一个只包含最终文件系统状态的新 layer，从而真正丢弃已被删除的构建工具和凭据文件。

相关背景：

- [容器镜像内部结构详解](./容器镜像内部结构详解.md)
- [使用 OverlayFS 联合挂载容器镜像层](./使用OverlayFS联合挂载容器镜像层.md)
- [提取容器镜像的根文件系统](./提取容器镜像根文件系统.md)

---

## 2. 为什么 `rm` 不能减小旧 layer

假设 Dockerfile 包含：

```dockerfile
RUN apt-get update && apt-get install -y build-essential libpq-dev
COPY pip.conf /root/.config/pip/pip.conf
RUN apt-get purge -y build-essential libpq-dev
RUN rm /root/.config/pip/pip.conf
```

每条会改变文件系统的指令通常生成一个新 layer：

```text
layer 1：+ 400 MB 构建工具
layer 2：+ pip.conf
layer 3：+ 删除构建工具的 whiteout 和其他变化
layer 4：+ 删除 pip.conf 的 whiteout
```

最终文件系统由 OverlayFS 一类联合文件系统组合：

```text
上层 whiteout 遮住下层文件
```

但 registry 仍然保存并分发旧 layer blob。攻击者只需要取得镜像 manifest，逐层下载和解压 blob，就可以绕过最终 rootfs 的遮挡，找回早期 layer 中的 `pip.conf`。

同理，删除 400 MB 构建工具只会让它们在容器中不可见，拉取镜像时旧 layer 仍然需要下载，因此镜像大小不会显著下降，甚至可能因为删除操作又增加一个小 layer。

### 2.1 在同一个 RUN 中删除是否有效

如果文件只在同一条 `RUN` 指令内部出现和删除，一般不会进入该指令最终生成的 layer：

```dockerfile
RUN apt-get update \
 && apt-get install -y --no-install-recommends build-essential libpq-dev \
 && build-something \
 && apt-get purge -y --auto-remove build-essential libpq-dev \
 && rm -rf /var/lib/apt/lists/*
```

但是把秘密通过 `COPY`、`ARG`、构建上下文或不安全命令带进构建过程，仍可能在缓存、历史、日志或其他 layer 中留下痕迹。对于凭据，应使用专用 secret mount，而不是依赖同层删除。

---

## 3. 如何发现镜像浪费和泄露

先查看构建历史：

```bash
docker history --no-trunc \
  ghcr.io/iximiuz/labs/bloated/app:v1.0.0
```

重点关注：

- 突然出现的大 layer。
- 安装后又在后续 layer 中删除的软件包。
- `COPY` 进去又被删除的配置、证书、私钥和 token 文件。
- 将凭据写进 `ARG`、`ENV` 或命令行的构建指令。
- `apt`、`pip`、`npm` 等包管理器缓存。

### 3.1 示例镜像的 history 解读

实验镜像的实际输出类似：

```text
IMAGE          CREATED BY                                      SIZE
6bde184cfcb3   CMD ["python" "app.py"]                         0B
<missing>      EXPOSE [5000/tcp]                               0B
<missing>      COPY app.py ./                                  12.3kB
<missing>      RUN apt-get purge -y build-essential ...        717kB
<missing>      RUN rm -f /root/.config/pip/pip.conf            16.4kB
<missing>      RUN pip install --no-cache-dir ...               19.3MB
<missing>      COPY requirements.txt ./                        12.3kB
<missing>      COPY pip.conf /root/.config/pip/pip.conf        20.5kB
<missing>      RUN apt-get install build-essential libpq-dev   388MB
...
```

`docker history` 按照“最新层在上、基础层在下”显示，因此阅读时可以从下往上还原 Dockerfile 的构建过程。

| History 条目 | 大小 | 应该怎样理解 |
| :--- | ---: | :--- |
| 安装 `build-essential`、`libpq-dev` | 388 MB | 构建工具和相关文件被完整写入一个历史 layer |
| `COPY pip.conf` | 20.5 KB | 包含凭据的配置文件进入独立 layer，已经构成泄露 |
| `pip install` | 19.3 MB | 安装应用运行依赖，通常是最终镜像需要保留的内容 |
| `rm pip.conf` | 16.4 KB | 新 layer 记录删除语义，但旧 layer 中的文件内容仍然存在 |
| `apt-get purge` | 717 KB | 新 layer 记录大量文件删除，不能回收 388 MB 安装 layer |
| `COPY app.py` | 12.3 KB | 添加最终应用代码 |
| `CMD`、`EXPOSE`、`ENV` | 0 B | 修改 image config，不产生 filesystem layer 内容 |

最值得对比的是安装层和清理层：

```text
388 MB：添加构建工具
   |
   | 后续 apt-get purge
   v
717 KB：记录删除结果和少量元数据变化
```

717 KB 的 purge layer 并没有抵消或改写 388 MB layer。容器最终看不到这些工具，只是因为更高层的删除信息把它们遮住；拉取镜像时，客户端仍然需要取得原始 388 MB layer blob。

凭据文件同理：

```text
COPY pip.conf layer：保存真实文件内容
rm pip.conf layer：让最终 rootfs 不再显示文件
```

只查看最终容器会误以为凭据已经清理，但逐层下载并解压镜像的人仍然可以从 `COPY pip.conf` 对应的 blob 中恢复它。

### 3.2 `<missing>` 表示什么

`<missing>` 不表示 layer blob 丢失，也不表示这一层无法恢复。它通常表示本地 Docker 没有保存对应的中间镜像对象和可显示的 image ID，这在从 registry 拉取的镜像以及 BuildKit 构建的镜像中很常见。

镜像正常运行所需的 config、manifest 和 filesystem layer blobs 仍然存在。攻击者也不需要中间 image ID，只需要沿 manifest 中的 descriptor 下载 layer blob 即可分析历史文件。

### 3.3 `SIZE` 应该怎样理解

`docker history` 中的 `SIZE` 表示该历史步骤产生的文件系统增量大小，适合定位异常的大 layer，但它不等于所有场景下的实际磁盘或网络开销：

- Registry 中的 layer blob 通常经过 gzip 或 zstd 压缩，下载大小应查看 manifest 的 `.layers[].size`。
- Docker 本地会在镜像之间复用相同 layer，不能简单把多个镜像显示的大小全部相加。
- OverlayFS 展开后的磁盘占用、registry 压缩大小和最终 rootfs 可见文件总大小是三个不同指标。
- `CMD`、`EXPOSE`、`ENV` 等只改变 image config，所以通常显示为 `0B`。
- `COPY` 一个很小的文件也可能显示数 KB，因为 layer tar 还包含目录项、文件头、对齐和其他元数据。

这个示例可以得出一个粗略判断：镜像的主要浪费来自 388 MB 构建工具层；squash 后如果最终 rootfs 确实已经删除这些工具，新单层镜像应显著变小，同时保留 Python 运行环境、约 19.3 MB 的应用依赖和应用代码。

不过，`docker history` 只是线索，不是完整的安全扫描。它看不到所有文件内容，也可能显示经过简化的构建历史。进一步分析可以使用：

- `dive`：逐层观察文件增加、修改、删除和空间浪费。
- `crane manifest`、`crane blob`：直接检查 registry 对象和 layers。
- Trivy 等工具：扫描漏洞、秘密和错误配置。
- 将镜像保存为 OCI layout 后逐层检查，而不只检查最终 rootfs。

安全审计要同时回答两个问题：

```text
最终 rootfs 中现在有什么？
历史 layers 中曾经出现过什么？
```

---

## 4. Squash 做了什么

Squash，也称 flatten，核心过程是：

```text
原镜像：Layer 0 + Layer 1 + Layer 2 + Layer 3
                    |
                    | 应用所有增加、覆盖、whiteout
                    v
              最终可见 rootfs
                    |
                    | 重新打包
                    v
新镜像：一个新的 filesystem layer
```

新 layer 只包含最终 rootfs 中仍然可见的文件：

- 已被删除的构建工具不会进入新 layer。
- 已被删除的 `pip.conf` 不会进入新 layer。
- 应用运行所需的最终文件会被保留。
- 原始 layer 边界以及被遮挡的历史内容被丢弃。

challenge 推荐使用 `crane flatten`：

```bash
crane flatten \
  --tag registry.iximiuz.com/squashed/app:v1.0.0 \
  ghcr.io/iximiuz/labs/bloated/app:v1.0.0
```

`crane flatten` 面向 registry 工作：它从源 registry 拉取镜像，合并 layers，再把新镜像推到目标 registry。使用前应检查当前版本帮助和 registry 登录状态：

```bash
crane flatten --help
crane auth login registry.iximiuz.com
```

目标镜像会获得全新的 manifest digest、config digest 和 layer digest，不能继续使用原镜像的签名或摘要进行身份校验。

---

## 5. 为什么 Docker 手工 flatten 很麻烦

只使用 Docker 可以采用类似思路：

```text
docker create
  -> docker export 得到容器的扁平文件系统 tar
  -> docker import 创建单层镜像
```

但这种方式容易丢失镜像配置，例如：

- `Entrypoint` 和 `Cmd`
- `Env`
- `User`
- `WorkingDir`
- `ExposedPorts`
- `Volumes`
- labels、healthcheck 和停止信号等

`docker import --change` 只能手动补回部分配置，既繁琐又容易与原镜像不一致。此外，`docker export` 不包含 volume 内容，并可能受到容器创建过程和运行时挂载的影响。

因此，手工 `export | import` 更适合作为理解原理的实验，不是首选的生产镜像修复方案。

---

## 6. Squash 能解决什么，不能解决什么

### 6.1 可以解决

- 清除只存在于旧 filesystem layer、最终 rootfs 已删除的文件。
- 移除旧 layers 中已经不可见的构建依赖和缓存。
- 在严重层浪费的镜像中显著减小传输和存储体积。
- 把遗留的多层镜像转换成单层 rootfs 表示。

### 6.2 不能保证解决

- 凭据仍存在于最终 rootfs。
- 凭据写在 image config 的 `Env`、labels 或其他元数据中。
- 凭据出现在构建日志、CI 日志、远程缓存或 provenance 中。
- 原始泄露镜像仍保存在 registry、本地缓存或其他节点中。
- 已经被攻击者读取和使用的凭据。
- 应用依赖本身的漏洞和运行时安全问题。

因此，“squash 后扫描不到 token”不等于泄露事件已经结束。

---

## 7. 泄露凭据必须立即轮换

一旦凭据进入过已推送或可被他人访问的镜像，应视为已经泄露。正确响应顺序是：

1. 立即撤销或轮换 registry token、密码或私钥。
2. 检查凭据使用日志，判断是否出现异常访问。
3. 构建并发布不含秘密的新镜像。
4. 从 registry 中删除或限制旧 tag，但不要假设删除 tag 就能立即清除所有 blob 和缓存副本。
5. 清理 CI 构建缓存、构建节点、本地镜像缓存和制品备份。
6. 重新扫描新镜像的 layers、config 和构建元数据。

Squash 只是制品修复手段，不是凭据轮换的替代品。

---

## 8. 更好的根治方案：重新设计 Dockerfile

### 8.1 使用多阶段构建

编译依赖应该停留在 builder stage，只把运行所需产物复制到最终镜像：

```dockerfile
FROM python:3.13-slim AS builder

RUN apt-get update \
 && apt-get install -y --no-install-recommends build-essential libpq-dev \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /build
COPY requirements.txt .
RUN pip wheel --no-cache-dir --wheel-dir /wheels -r requirements.txt

FROM python:3.13-slim

RUN apt-get update \
 && apt-get install -y --no-install-recommends libpq5 \
 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /wheels /wheels
RUN pip install --no-cache-dir /wheels/* \
 && rm -rf /wheels

WORKDIR /app
COPY . .
CMD ["python", "server.py"]
```

这样编译器、头文件和 builder 的包管理器缓存根本不会成为最终镜像的 layer。

### 8.2 使用 BuildKit secret mount

不要执行：

```dockerfile
COPY pip.conf /root/.config/pip/pip.conf
```

可以使用 BuildKit secret，让凭据只在某个 `RUN` 执行期间临时挂载：

```dockerfile
# syntax=docker/dockerfile:1
FROM python:3.13-slim AS builder

RUN --mount=type=secret,id=pip_conf,target=/root/.config/pip/pip.conf \
    pip wheel --no-cache-dir --wheel-dir /wheels -r requirements.txt
```

构建时传入：

```bash
docker build \
  --secret id=pip_conf,src="$HOME/.config/pip/pip.conf" \
  -t example/app:v1.0.1 .
```

secret mount 不会被复制进最终 layer。仍应避免在命令输出、错误日志或下载产物中打印秘密。

### 8.3 控制构建上下文

使用 `.dockerignore` 排除不应发送给 builder 的文件：

```text
.git
.env
*.pem
*.key
pip.conf
node_modules
__pycache__
```

`.dockerignore` 是防误传措施，不是 secret mount 的替代品；真正需要在构建期间使用的秘密仍应通过受控 secret 机制传入。

---

## 9. Squash 的代价

Squash 并不是所有镜像都应该执行的常规优化。

| 收益 | 代价 |
| :--- | :--- |
| 删除旧 layer 中的无效文件 | 丢失原始 layer 边界和可读构建历史 |
| 严重浪费时明显减小体积 | 不同镜像之间的 layer 复用能力下降 |
| 清除仅存在于历史 layer 的文件 | 修改少量内容也可能需要重新拉取大单层 |
| 生成简单的单层 rootfs | SBOM、签名、provenance 和 digest 必须重新生成 |
| 可作为遗留镜像补救手段 | 不能替代可重复、可审计的干净重建 |

一个合理原则是：

```text
能修改 Dockerfile并重新构建时，优先干净重建；
无法重建的遗留镜像，才考虑 squash 作为补救。
```

---

## 10. 如何验证 squash 结果

### 10.1 确认 layer 数量

```bash
crane manifest registry.iximiuz.com/squashed/app:v1.0.0 \
  | jq '.layers | length'
```

预期 filesystem layer 数量为 1。多平台镜像需要先选定具体平台 manifest，不能直接把顶层 Image Index 当成单平台 manifest 检查。

### 10.2 比较体积

```bash
crane manifest ghcr.io/iximiuz/labs/bloated/app:v1.0.0 \
  | jq '[.layers[].size] | add'

crane manifest registry.iximiuz.com/squashed/app:v1.0.0 \
  | jq '[.layers[].size] | add'
```

这里比较的是 registry layer blob 的压缩大小。`docker images` 显示的本地占用、解压后的 rootfs 大小和实际网络传输量不是同一个指标。

### 10.3 检查最终 rootfs

确认应用文件、动态库、证书、时区数据等运行依赖没有丢失，并确认凭据文件不存在：

```bash
crane export registry.iximiuz.com/squashed/app:v1.0.0 - \
  | tar -tf - \
  | grep -E 'pip\.conf|\.netrc|\.npmrc|\.pypirc'
```

没有文件名不代表内容中没有秘密，还应进行 secret scanner 检查。

### 10.4 验证镜像配置和应用行为

对比新旧镜像的：

```text
Entrypoint、Cmd、Env、User、WorkingDir
端口、volume、healthcheck、停止信号
目标 OS/CPU 平台
```

然后实际运行测试：

```bash
docker pull registry.iximiuz.com/squashed/app:v1.0.0
docker run --rm registry.iximiuz.com/squashed/app:v1.0.0
```

还应重新生成 SBOM、重新运行漏洞和秘密扫描，并为新 digest 重新签名、重新生成 provenance。

---

## 11. 总结

这篇 challenge 最重要的结论是：

```text
容器中删除文件，只改变最终联合文件系统视图；
只有丢弃包含该文件的历史 layer，才能从新镜像中真正移除它。
```

Squash 的作用是把最终 rootfs 重新打成一个新 layer，因此能够清理被 whiteout 遮住的构建工具和文件系统凭据。但安全上还必须注意：

- 立即轮换已经进入过镜像的凭据。
- 检查 config、日志、缓存和 provenance，而不只检查 filesystem layers。
- 能重新构建时，优先采用多阶段构建和 BuildKit secret mount。
- Squash 后重新扫描、生成 SBOM、签名并验证应用行为。
- 不要为了追求 layer 数量少而默认 squash 所有镜像；layer 复用本身是容器镜像的重要优势。
