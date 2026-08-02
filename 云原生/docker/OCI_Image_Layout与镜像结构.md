# OCI Image Layout 与容器镜像结构

> 专题导航：[容器镜像文件系统专题索引](./容器镜像文件系统专题索引.md)

这篇 challenge 想教的核心不是“怎么导出镜像”，而是让你理解：**一个容器镜像在 registry 或 OCI layout 里，本质上是一组通过 digest 相互引用的 JSON 文档和 layer blob**。

Docker 平时把这些结构隐藏起来了。你看到的是：

```bash
docker pull nginx
docker run nginx
docker images
```

但镜像内部其实有清晰的层次：

```text
OCI Layout
  |
  +-- index.json
        |
        v
      manifest 或 image index
        |
        +-- config blob
        +-- layer blob
        +-- layer blob
        +-- ...
```

单平台镜像和多平台镜像的区别就在这里：

```text
单平台镜像：
  index.json -> image manifest -> config + layers

多平台镜像：
  index.json -> image index -> platform manifests -> config + layers
```

## 1. 什么是 OCI Image Layout

OCI Image Layout 是 OCI 规范定义的一种本地目录格式，用来把容器镜像保存到普通文件系统目录中。

典型结构：

```text
single-layout/
├── oci-layout
├── index.json
└── blobs/
    └── sha256/
        ├── <digest1>
        ├── <digest2>
        ├── <digest3>
        └── ...
```

几个关键文件：

| 文件/目录 | 作用 |
| :--- | :--- |
| `oci-layout` | 标记这是一个 OCI layout，里面通常有 layout 版本 |
| `index.json` | 入口文件，列出顶层 descriptor |
| `blobs/sha256/` | 存放所有内容 blob，包括 JSON 文档和 layer tar |

OCI layout 是 content-addressable 的。也就是说，blob 文件名就是它内容的 SHA-256 digest。

例如某个 descriptor 里写：

```json
{
  "digest": "sha256:abc123..."
}
```

那么对应文件路径就是：

```text
blobs/sha256/abc123...
```

注意路径里没有 `sha256:` 前缀，只保留后面的 hex 部分。

## 2. descriptor 是什么

OCI 里大量使用 descriptor。它是一个“指向某个 blob 的引用”。

常见字段：

```json
{
  "mediaType": "application/vnd.oci.image.manifest.v1+json",
  "digest": "sha256:...",
  "size": 1234
}
```

含义：

| 字段 | 说明 |
| :--- | :--- |
| `mediaType` | 这个 blob 的类型，例如 manifest、index、config、layer |
| `digest` | blob 内容的 sha256 摘要 |
| `size` | blob 大小 |
| `platform` | 多平台镜像里常见，说明该 descriptor 对应哪个平台 |

因此探索 OCI layout 的方式就是：

```text
打开 index.json
  |
  v
读取 manifests[0].digest
  |
  v
去 blobs/sha256/<digest> 打开下一个 JSON
  |
  v
继续跟随 config.digest 或 layers[].digest
```

## 3. 保存镜像为 OCI layout

challenge 里给了两个镜像：

```text
ghcr.io/iximiuz/labs/oci/web:v1  单平台镜像
ghcr.io/iximiuz/labs/oci/cli:v1  多平台镜像
```

如果镜像已经 pull 到本地，可以用 `docker save` 导出，再解压：

```bash
mkdir -p ~/single-layout
docker save ghcr.io/iximiuz/labs/oci/web:v1 -o /tmp/single.tar
tar -xf /tmp/single.tar -C ~/single-layout
```

多平台镜像：

```bash
mkdir -p ~/multi-layout
docker save ghcr.io/iximiuz/labs/oci/cli:v1 -o /tmp/multi.tar
tar -xf /tmp/multi.tar -C ~/multi-layout
```

检查结构：

```bash
find ~/single-layout -maxdepth 3 -type f | sort
find ~/multi-layout -maxdepth 3 -type f | sort
```

如果环境有 `jq`，后续查看 JSON 会方便很多：

```bash
jq . ~/single-layout/index.json
```

## 4. 单平台镜像结构

单平台镜像的链路通常是：

```text
index.json
  |
  v
OCI Image Manifest
  |
  +-- config digest
  |
  +-- layer digest
  +-- layer digest
  +-- ...
```

### 4.1 查看入口 index.json

```bash
jq . ~/single-layout/index.json
```

重点看：

```bash
jq -r '.manifests[0].mediaType' ~/single-layout/index.json
```

单平台镜像这里通常会是：

```text
application/vnd.oci.image.manifest.v1+json
```

也就是 `index.json` 的顶层 descriptor 直接指向一个 image manifest。

### 4.2 跟随 digest 找到 manifest

先取 digest：

```bash
MANIFEST_DIGEST=$(jq -r '.manifests[0].digest' ~/single-layout/index.json)
echo "$MANIFEST_DIGEST"
```

把 `sha256:` 去掉，得到 blob 文件路径：

```bash
MANIFEST_FILE=~/single-layout/blobs/sha256/${MANIFEST_DIGEST#sha256:}
jq . "$MANIFEST_FILE"
```

manifest 里重点字段：

```json
{
  "schemaVersion": 2,
  "mediaType": "application/vnd.oci.image.manifest.v1+json",
  "config": {
    "mediaType": "application/vnd.oci.image.config.v1+json",
    "digest": "sha256:...",
    "size": 1234
  },
  "layers": [
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "digest": "sha256:...",
      "size": 1234
    }
  ]
}
```

### 4.3 统计 filesystem layers

```bash
jq '.layers | length' "$MANIFEST_FILE"
```

这里数的是 manifest 的 `.layers` 数组长度。

每个 layer 是一个文件系统差异层。多个 layer 叠加后，才形成容器看到的 root filesystem。

### 4.4 查看 image config

manifest 里的 `.config.digest` 指向 image config blob：

```bash
CONFIG_DIGEST=$(jq -r '.config.digest' "$MANIFEST_FILE")
CONFIG_FILE=~/single-layout/blobs/sha256/${CONFIG_DIGEST#sha256:}
jq . "$CONFIG_FILE"
```

config 里包含：

| 字段 | 说明 |
| :--- | :--- |
| `architecture` | 镜像目标 CPU 架构，例如 `amd64`、`arm64` |
| `os` | 目标操作系统，例如 `linux` |
| `config` | 运行时默认配置，例如 `Entrypoint`、`Cmd`、`Env` |
| `rootfs.diff_ids` | 未压缩 layer 的 digest 列表 |
| `history` | Dockerfile 构建历史 |

查看架构：

```bash
jq -r '.architecture' "$CONFIG_FILE"
```

### 4.5 ImageID 是什么

OCI Image Spec 中，ImageID 是 image config blob 的 digest。

所以单平台镜像的 ImageID 就是：

```bash
jq -r '.config.digest' "$MANIFEST_FILE"
```

也就是：

```text
sha256:<config-blob-digest>
```

为什么 config digest 可以作为 ImageID？

因为 config 里记录了：

- 目标平台：`os`、`architecture`
- rootfs diff IDs
- 运行参数：`Entrypoint`、`Cmd`、`Env`
- 构建历史

config 内容一变，digest 就会变；digest 一变，ImageID 就不同。

注意：challenge 里提到 Docker 29 开始显示的 Image ID 可能是 manifest blob 的 digest，而不是 OCI spec 定义的 config digest。Podman、crictl、旧版 Docker 更接近 OCI spec 的定义。

## 5. 多平台镜像结构

多平台镜像比单平台多一层 image index。

结构：

```text
index.json
  |
  v
OCI Image Index
  |
  +-- linux/amd64 manifest
  |     |
  |     +-- amd64 config
  |     +-- amd64 layers
  |
  +-- linux/arm64 manifest
  |     |
  |     +-- arm64 config
  |     +-- arm64 layers
  |
  +-- ...
```

容易混淆的是：这里有两个 “index”。

| 文件 | 是否总是存在 | 作用 |
| :--- | :--- | :--- |
| 顶层 `index.json` | 是 | OCI layout 的入口文件 |
| blob 里的 OCI Image Index | 只在多平台镜像中常见 | 列出各个平台的 manifest |

它们的 mediaType 可能都是：

```text
application/vnd.oci.image.index.v1+json
```

但它们不是同一个东西。

## 6. 多平台镜像探索步骤

查看顶层入口：

```bash
jq . ~/multi-layout/index.json
```

查看顶层 entry 的 mediaType：

```bash
jq -r '.manifests[0].mediaType' ~/multi-layout/index.json
```

多平台镜像这里通常指向 OCI Image Index：

```text
application/vnd.oci.image.index.v1+json
```

跟随 digest 找到真正的 image index blob：

```bash
INDEX_DIGEST=$(jq -r '.manifests[0].digest' ~/multi-layout/index.json)
INDEX_FILE=~/multi-layout/blobs/sha256/${INDEX_DIGEST#sha256:}
jq . "$INDEX_FILE"
```

这个 blob 里的 `.manifests` 数组才是平台列表。

## 7. 统计平台 manifest 数量

challenge 特别提醒：有些 image index 里可能包含 `unknown/unknown` 的 attestation、SBOM 或 provenance descriptor，不应该把它算作普通平台镜像。

统计 platform architecture 不是 unknown 的条目：

```bash
jq '[.manifests[] | select(.platform.architecture != "unknown")] | length' "$INDEX_FILE"
```

查看平台列表：

```bash
jq -r '.manifests[]
  | select(.platform.architecture != "unknown")
  | "\(.platform.os)/\(.platform.architecture)\(if .platform.variant then "/" + .platform.variant else "" end) \(.digest)"' "$INDEX_FILE"
```

输出类似：

```text
linux/amd64 sha256:...
linux/arm64 sha256:...
linux/arm/v7 sha256:...
```

## 8. 找 linux/arm64 的 manifest digest

```bash
jq -r '.manifests[]
  | select(.platform.os == "linux" and .platform.architecture == "arm64")
  | .digest' "$INDEX_FILE"
```

这个 digest 指向的是 ARM64 平台自己的 image manifest，不是 config，也不是 layer。

打开它：

```bash
ARM64_MANIFEST_DIGEST=$(jq -r '.manifests[]
  | select(.platform.os == "linux" and .platform.architecture == "arm64")
  | .digest' "$INDEX_FILE")

ARM64_MANIFEST_FILE=~/multi-layout/blobs/sha256/${ARM64_MANIFEST_DIGEST#sha256:}
jq . "$ARM64_MANIFEST_FILE"
```

然后它的结构就和单平台 manifest 一样：

```text
platform manifest
  |
  +-- config.digest
  +-- layers[].digest
```

## 9. 多平台镜像的 ImageID 是否相同

每个平台 manifest 都有自己的 config：

```json
{
  "config": {
    "digest": "sha256:..."
  },
  "layers": [...]
}
```

ImageID 是 config digest。所以要比较不同平台的 ImageID，就要比较每个平台 manifest 的 `.config.digest`。

命令：

```bash
for digest in $(jq -r '.manifests[]
  | select(.platform.architecture != "unknown")
  | .digest' "$INDEX_FILE"); do
  file=~/multi-layout/blobs/sha256/${digest#sha256:}
  platform=$(jq -r --arg d "$digest" '.manifests[]
    | select(.digest == $d)
    | "\(.platform.os)/\(.platform.architecture)\(if .platform.variant then "/" + .platform.variant else "" end)"' "$INDEX_FILE")
  config=$(jq -r '.config.digest' "$file")
  echo "$platform $config"
done
```

如果输出的 config digest 都不同，那么答案是：

```text
no
```

通常多平台镜像的不同平台变体不会有相同 ImageID。因为 config 里至少会记录不同的：

```json
{
  "architecture": "amd64"
}
```

或者：

```json
{
  "architecture": "arm64"
}
```

只要 config 内容不同，config digest 就不同，ImageID 也就不同。

## 10. mediaType 速查

常见 OCI mediaType：

| mediaType | 含义 |
| :--- | :--- |
| `application/vnd.oci.image.index.v1+json` | OCI Image Index，多平台索引 |
| `application/vnd.oci.image.manifest.v1+json` | OCI Image Manifest，单个平台镜像描述 |
| `application/vnd.oci.image.config.v1+json` | OCI Image Configuration，镜像配置 |
| `application/vnd.oci.image.layer.v1.tar` | 未压缩 layer |
| `application/vnd.oci.image.layer.v1.tar+gzip` | gzip 压缩 layer |
| `application/vnd.docker.distribution.manifest.list.v2+json` | Docker schema2 manifest list |
| `application/vnd.docker.distribution.manifest.v2+json` | Docker schema2 image manifest |
| `application/vnd.docker.container.image.v1+json` | Docker image config |

challenge 使用 OCI layout，但实际 Docker 导出的内容可能混有 Docker mediaType。这不影响理解，核心结构仍然是：

```text
index -> manifest/index -> config + layers
```

## 11. 单平台和多平台对比

| 对比项 | 单平台镜像 | 多平台镜像 |
| :--- | :--- | :--- |
| 顶层 `index.json` 指向 | image manifest | image index |
| 是否有平台列表 | 没有 | 有 |
| 是否直接描述 layers | 是，manifest 直接有 `.layers` | 否，先选平台 manifest |
| ImageID 来源 | manifest 的 `.config.digest` | 每个平台 manifest 的 `.config.digest` |
| 不同平台 ImageID | 不涉及 | 通常不同 |

## 12. 完整排查命令汇总

单平台：

```bash
mkdir -p ~/single-layout
docker save ghcr.io/iximiuz/labs/oci/web:v1 -o /tmp/single.tar
tar -xf /tmp/single.tar -C ~/single-layout

jq -r '.manifests[0].mediaType' ~/single-layout/index.json

MANIFEST_DIGEST=$(jq -r '.manifests[0].digest' ~/single-layout/index.json)
MANIFEST_FILE=~/single-layout/blobs/sha256/${MANIFEST_DIGEST#sha256:}

jq '.layers | length' "$MANIFEST_FILE"

CONFIG_DIGEST=$(jq -r '.config.digest' "$MANIFEST_FILE")
CONFIG_FILE=~/single-layout/blobs/sha256/${CONFIG_DIGEST#sha256:}

jq -r '.architecture' "$CONFIG_FILE"
echo "$CONFIG_DIGEST"
```

多平台：

```bash
mkdir -p ~/multi-layout
docker save ghcr.io/iximiuz/labs/oci/cli:v1 -o /tmp/multi.tar
tar -xf /tmp/multi.tar -C ~/multi-layout

jq -r '.manifests[0].mediaType' ~/multi-layout/index.json

INDEX_DIGEST=$(jq -r '.manifests[0].digest' ~/multi-layout/index.json)
INDEX_FILE=~/multi-layout/blobs/sha256/${INDEX_DIGEST#sha256:}

jq '[.manifests[] | select(.platform.architecture != "unknown")] | length' "$INDEX_FILE"

jq -r '.manifests[]
  | select(.platform.os == "linux" and .platform.architecture == "arm64")
  | .digest' "$INDEX_FILE"

for digest in $(jq -r '.manifests[]
  | select(.platform.architecture != "unknown")
  | .digest' "$INDEX_FILE"); do
  file=~/multi-layout/blobs/sha256/${digest#sha256:}
  platform=$(jq -r --arg d "$digest" '.manifests[]
    | select(.digest == $d)
    | "\(.platform.os)/\(.platform.architecture)\(if .platform.variant then "/" + .platform.variant else "" end)"' "$INDEX_FILE")
  config=$(jq -r '.config.digest' "$file")
  echo "$platform $config"
done
```

## 13. 这篇 challenge 真正想建立的心智模型

容器镜像不是一个黑盒 tar 包，而是一棵 digest 引用树。

单平台：

```text
index.json
  -> image manifest
      -> config
      -> layers
```

多平台：

```text
index.json
  -> image index
      -> linux/amd64 manifest
          -> amd64 config
          -> amd64 layers
      -> linux/arm64 manifest
          -> arm64 config
          -> arm64 layers
```

`docker pull`、`docker run`、`docker buildx imagetools inspect` 背后都在处理这套结构：

- registry 存 blob。
- manifest 引用 config 和 layers。
- image index 引用多个平台 manifest。
- digest 保证内容寻址和不可变性。
- ImageID 来自 config digest。

理解这套结构后，再看多架构镜像、manifest list、ImageID、layer 复用、registry 存储，都会清楚很多。

------

# 提取容器镜像的最终文件系统

前面的 OCI layout 章节关注的是“镜像如何存储”：`index`、`manifest`、`config`、`layers` 如何通过 digest 串起来。

这一节关注另一个问题：**如何把一个镜像最终会呈现给容器的 root filesystem 提取出来**。

这两个问题不一样：

| 问题 | 关注点 | 常用工具 |
| :--- | :--- | :--- |
| 镜像是怎么存的 | manifest、config、layers、digest | `docker save`、OCI layout |
| 容器启动前能看到哪些文件 | 合并后的 rootfs | `docker export`、`docker build -o`、`ctr image mount` |

开发者经常想做的是：

```text
我不想运行这个镜像。
我只想看看它里面有哪些文件。
```

例如：

```bash
ls rootfs
cat rootfs/etc/os-release
find rootfs/usr -name nginx
```

## 14. 为什么 `docker save` 不适合提取 rootfs

`docker save` 的说明是：

```bash
docker save --help
```

```text
Save one or more images to a tar archive
```

它导出的是镜像在 Docker 存储里的 layered representation：

```text
manifest
config
layer1.tar
layer2.tar
...
```

也就是说：

```bash
docker save ghcr.io/iximiuz/labs/nginx:alpine -o nginx-image.tar
```

得到的是镜像结构，不是最终合并后的文件系统。

`docker save` 更适合：

- 研究镜像 manifest/config/layers。
- 离线搬运镜像。
- 配合 `docker load` 恢复镜像。

它不适合直接回答：

```text
这个镜像最终的 /etc、/usr、/var 目录长什么样？
```

因为镜像的 rootfs 需要把多个 layer 按顺序叠加后才能得到。

## 15. `docker export` 导出的是 container，不是 image

`docker export` 看起来很接近目标：

```bash
docker export --help
```

```text
Export a container's filesystem as a tar archive
```

但重点是：它导出的是 container filesystem，而不是 image。

所以这样会失败：

```bash
docker export ghcr.io/iximiuz/labs/nginx:alpine -o nginx.tar
```

原因：

```text
ghcr.io/iximiuz/labs/nginx:alpine 是 image，不是 container
docker export 需要的是 container id 或 container name
```

## 16. 不推荐方案：`docker run + docker export`

一种直觉方案是先运行容器，再导出：

```bash
CONT_ID=$(docker run -d ghcr.io/iximiuz/labs/nginx:alpine)
docker export "$CONT_ID" -o nginx.tar
```

然后解压：

```bash
mkdir rootfs
tar -xf nginx.tar -C rootfs
ls -l rootfs
```

这能得到类似：

```text
bin -> usr/bin
boot
dev
etc
usr
var
...
```

但它有明显问题：

| 问题 | 说明 |
| :--- | :--- |
| 容器真的启动了 | 运行任意镜像有安全风险 |
| 可能很慢 | 镜像 entrypoint 可能执行大量启动逻辑 |
| 文件系统可能被污染 | entrypoint/startup 可能修改配置文件 |
| 镜像可能无法运行 | broken image 无法启动，但仍可能想看文件 |

所以这个方案能用，但不是推荐方案。

## 17. 推荐方案：`docker create + docker export`

关键点：

```text
docker run = docker create + docker start
```

`docker create` 会创建一个 container，并准备好它的 root filesystem，但不会启动容器进程。

所以可以：

```bash
CONT_ID=$(docker create ghcr.io/iximiuz/labs/nginx:alpine)
docker export "$CONT_ID" -o nginx.tar
docker rm "$CONT_ID"
```

解压：

```bash
mkdir rootfs
tar -xf nginx.tar -C rootfs
```

或者一行：

```bash
mkdir rootfs
CONT_ID=$(docker create ghcr.io/iximiuz/labs/nginx:alpine)
docker export "$CONT_ID" | tar -xC rootfs
docker rm "$CONT_ID"
```

这个方案的优点：

| 优点 | 说明 |
| :--- | :--- |
| 不启动进程 | 避免 entrypoint 修改文件系统 |
| 只依赖 Docker 标准命令 | 不需要 containerd/ctr |
| 得到的是合并后的 rootfs | 可以直接 `ls`、`cat`、`find` |

缺点：

| 缺点 | 说明 |
| :--- | :--- |
| 仍然创建了 container | 需要记得 `docker rm` |
| 可能有运行时注入文件 | 例如 `/etc/hosts`、`/etc/resolv.conf` 等 |
| 不是纯原始 image rootfs | Docker create 会做少量容器环境准备 |

## 18. 解压时保留权限、属主和扩展属性

普通解压：

```bash
tar -xf nginx.tar -C rootfs
```

可能会让文件属主变成当前用户，也可能影响权限或 capabilities。

如果需要尽量保留原始信息，可以用：

```bash
sudo tar -xf nginx.tar \
  --numeric-owner \
  --same-permissions \
  --xattrs --xattrs-include='*' \
  -C rootfs
```

参数含义：

| 参数 | 作用 |
| :--- | :--- |
| `--numeric-owner` | 按 UID/GID 数字恢复属主，而不是映射成本机用户名 |
| `--same-permissions` | 尽量保留原始权限 |
| `--xattrs` | 恢复扩展属性 |
| `--xattrs-include='*'` | 包含所有扩展属性 |

注意：完整恢复属主、权限、xattrs 通常需要 root 权限。

## 19. 更准确方案：`docker build -o`

`docker create + docker export` 很实用，但它导出的其实是 container filesystem，不是完全纯粹的 image rootfs。Docker 创建 container 时可能额外注入一些文件，例如：

```text
/etc/hosts
/etc/resolv.conf
```

如果想更接近原始 image rootfs，可以利用 BuildKit 的 `--output`：

```bash
mkdir rootfs
cat > Dockerfile <<'EOF'
FROM ghcr.io/iximiuz/labs/nginx:alpine
EOF

docker build -o rootfs .
```

如果 Docker 版本较老，需要显式启用 BuildKit：

```bash
DOCKER_BUILDKIT=1 docker build -o rootfs .
```

也可以用 buildx：

```bash
docker buildx build -o type=local,dest=rootfs .
```

这个 Dockerfile 没有 `RUN` 指令：

```dockerfile
FROM ghcr.io/iximiuz/labs/nginx:alpine
```

因此 BuildKit 不需要执行构建命令，只需要把 `FROM` 镜像的文件系统导出到指定目录。

优点：

| 优点 | 说明 |
| :--- | :--- |
| 不需要启动容器 | 不执行 entrypoint/cmd |
| 更接近 image rootfs | 避免一部分 `docker create` 注入内容 |
| 输出就是普通目录 | 可以直接浏览文件 |

缺点：

| 缺点 | 说明 |
| :--- | :--- |
| 依赖 BuildKit | 老版本 Docker 需要 `DOCKER_BUILDKIT=1` |
| ownership 保留可能不如 export 直观 | 某些场景下文件属主信息处理需要注意 |
| 写法看起来不像“提取镜像” | 本质是借用 build output 功能 |

## 20. containerd 方案：`ctr image mount`

Docker 底层越来越多能力委托给 containerd。containerd 的命令行工具 `ctr` 可以直接 mount 镜像 rootfs。

先 pull 到 containerd：

```bash
sudo ctr image pull ghcr.io/iximiuz/labs/nginx:alpine
```

mount 镜像：

```bash
sudo ctr image mount ghcr.io/iximiuz/labs/nginx:alpine
```

命令会输出一个挂载目录，例如：

```text
/run/containerd/io.containerd.mount-manager.v1.bolt/t/1/1
```

查看：

```bash
sudo ls -l /run/containerd/io.containerd.mount-manager.v1.bolt/t/1/1
```

这个目录就是镜像合并后的 rootfs。

如果 Docker 和 containerd 使用不同 image store，可能需要用 `ctr image pull` 额外拉一次。Docker 29 之后，Docker 默认更多使用 containerd image store；在这种情况下，可能可以通过 `moby` namespace 访问 Docker 管理的镜像：

```bash
sudo ctr --namespace moby image ls
sudo ctr --namespace moby image mount ghcr.io/iximiuz/labs/nginx:alpine
```

优点：

| 优点 | 说明 |
| :--- | :--- |
| 不需要创建 Docker container | 直接操作 image |
| 结果干净 | 没有 `docker create` 的额外容器注入 |
| 更接近底层运行时 | 适合理解 containerd |

缺点：

| 缺点 | 说明 |
| :--- | :--- |
| 通常需要 root | mount 操作需要权限 |
| 不一定安装 `ctr` | 某些环境只有 Docker CLI |
| Docker image store 兼容性要看版本 | 旧版本 Docker 和 containerd 存储可能分离 |

## 21. 方法对比

| 方法 | 得到什么 | 是否运行容器 | 是否创建 container | 适合场景 |
| :--- | :--- | :--- | :--- | :--- |
| `docker save` | 镜像存储结构：manifest/config/layers | 否 | 否 | 研究 OCI 镜像结构、离线搬运 |
| `docker run + docker export` | container rootfs | 是 | 是 | 能用但不推荐 |
| `docker create + docker export` | container rootfs | 否 | 是 | Docker 标准命令下最实用 |
| `docker build -o` | image rootfs | 否 | 否 | 更干净地导出原始镜像文件系统 |
| `ctr image mount` | mounted image rootfs | 否 | 否 | containerd 底层方式，准确但要求更多 |

推荐顺序：

```text
只是想看看镜像文件系统：
  docker create + docker export

想尽量避免 Docker create 注入内容：
  docker build -o

想研究镜像结构：
  docker save / OCI layout

有 containerd 权限并想直接 mount：
  ctr image mount
```

## 22. 常用命令汇总

### 22.1 create + export

```bash
IMAGE=ghcr.io/iximiuz/labs/nginx:alpine
DEST=rootfs

mkdir -p "$DEST"
CONT_ID=$(docker create "$IMAGE")
docker export "$CONT_ID" | tar -xC "$DEST"
docker rm "$CONT_ID"
```

保留权限、属主、xattrs：

```bash
IMAGE=ghcr.io/iximiuz/labs/nginx:alpine
DEST=rootfs

mkdir -p "$DEST"
CONT_ID=$(docker create "$IMAGE")
docker export "$CONT_ID" | sudo tar -xC "$DEST" \
  --numeric-owner \
  --same-permissions \
  --xattrs --xattrs-include='*'
docker rm "$CONT_ID"
```

### 22.2 build output

```bash
IMAGE=ghcr.io/iximiuz/labs/nginx:alpine
DEST=rootfs

mkdir -p "$DEST"
cat > Dockerfile.extract <<EOF
FROM $IMAGE
EOF

docker build -f Dockerfile.extract -o "$DEST" .
```

使用 buildx：

```bash
docker buildx build -f Dockerfile.extract -o type=local,dest=rootfs .
```

### 22.3 ctr mount

```bash
IMAGE=ghcr.io/iximiuz/labs/nginx:alpine

sudo ctr image pull "$IMAGE"
sudo ctr image mount "$IMAGE"
```

如果使用 Docker 的 containerd namespace：

```bash
sudo ctr --namespace moby image ls
sudo ctr --namespace moby image mount "$IMAGE"
```

## 23. 这篇文章真正想建立的心智模型

同一个镜像可以用两种角度看：

```text
镜像存储视角：
  image = manifest + config + layers
  用 docker save / OCI layout 看

最终文件系统视角：
  image rootfs = layers 叠加后的结果
  用 docker export / docker build -o / ctr image mount 看
```

所以：

```text
docker save 不是“导出 rootfs”
docker export 是“导出 container rootfs”
docker create + docker export 是“不启动容器也能导出 rootfs”
docker build -o 是“借 BuildKit 输出 image rootfs”
ctr image mount 是“用 containerd 直接挂载 image rootfs”
```

理解这个差异后，就不会把 `docker save`、`docker export`、OCI layout、rootfs extraction 混在一起。
