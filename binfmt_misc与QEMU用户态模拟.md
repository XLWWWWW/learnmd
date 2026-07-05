# binfmt_misc 与 QEMU 用户态模拟

这个实验的核心问题是：在 `amd64/x86_64` 主机上直接执行一个 `arm64/aarch64` Linux ELF 文件时，CPU 无法原生执行 ARM64 指令，所以内核会返回 `Exec format error`。解决方式不是把整台机器虚拟化成 ARM，而是让 Linux 内核在 `execve()` 阶段把这个 ARM64 ELF 自动交给 QEMU 用户态模拟器运行。

## 1. 先确认二进制架构

```bash
file ~/simple-server
```

典型输出会类似：

```text
/home/ubuntu/simple-server: ELF 64-bit LSB executable, ARM aarch64, ...
```

当前机器如果是 x86_64：

```bash
uname -m
```

典型输出：

```text
x86_64
```

这说明 `simple-server` 是 ARM64 程序，而宿主机 CPU 是 x86_64。未配置模拟器时直接运行：

```bash
~/simple-server
```

会报错：

```text
cannot execute binary file: Exec format error
```

这个错误不是权限问题，而是内核识别到 ELF 文件的目标机器架构和当前 CPU 不匹配。

## 2. 安装 QEMU 用户态模拟

Ubuntu 上通常安装提供静态 QEMU 用户态解释器的包：

```bash
sudo apt update
sudo apt install -y qemu-user-static binfmt-support
```

其中：

| 组件 | 作用 |
| :--- | :--- |
| `qemu-user-static` | 提供各架构的静态用户态模拟器，例如 `qemu-aarch64-static` |
| `binfmt-support` | 管理 `binfmt_misc` 注册项，帮助内核识别外来格式并调用解释器 |
| `binfmt_misc` | Linux 内核机制，根据文件魔数把某类可执行文件转交给指定解释器 |

安装后检查 ARM64 注册项：

```bash
ls /proc/sys/fs/binfmt_misc/qemu-aarch64
cat /proc/sys/fs/binfmt_misc/qemu-aarch64
```

如果注册成功，可以看到类似内容：

```text
enabled
interpreter /usr/libexec/qemu-binfmt/aarch64-binfmt-P
flags: OPF
offset 0
magic 7f454c46...
mask ffffffffff...
```

不同 Ubuntu 版本的 `interpreter` 路径可能不同，常见有：

```text
/usr/bin/qemu-aarch64-static
/usr/libexec/qemu-binfmt/aarch64-binfmt-P
```

重点不是路径完全一致，而是存在 `qemu-aarch64` 注册项，并且状态是 `enabled`。

如果 `/proc/sys/fs/binfmt_misc/` 下没有内容，可以确认 `binfmt_misc` 是否挂载：

```bash
mount | grep binfmt_misc
sudo mount -t binfmt_misc binfmt_misc /proc/sys/fs/binfmt_misc
```

## 3. 运行 ARM64 HTTP 服务

安装并注册成功后，再运行：

```bash
~/simple-server
```

如果它是前台服务，会占住当前终端。为了后续排查，可以放到后台：

```bash
~/simple-server &
```

也可以记录 shell 返回的后台 PID：

```bash
SERVER_PID=$!
echo "$SERVER_PID"
```

验证 8080 端口：

```bash
curl http://127.0.0.1:8080/
```

## 4. 找到监听 8080 的进程 PID

推荐用 `ss`：

```bash
sudo ss -ltnp 'sport = :8080'
```

典型输出：

```text
State  Recv-Q Send-Q Local Address:Port Peer Address:Port Process
LISTEN 0      128          0.0.0.0:8080      0.0.0.0:*     users:(("qemu-aarch64-st",pid=12345,fd=3))
```

也可以用 `lsof`：

```bash
sudo lsof -nP -iTCP:8080 -sTCP:LISTEN
```

典型输出：

```text
COMMAND     PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
qemu-aarc 12345 user    3u  IPv4  ...      0t0  TCP *:8080 (LISTEN)
```

这里要注意：监听端口的进程名可能显示为 `qemu-aarch64-static`、`qemu-aarch64-st`，也可能因为 `binfmt_misc` 的参数处理显示成接近原程序的名字。不要只看 `COMMAND` 列，最终要看 `/proc/<PID>/exe`。

## 5. 识别真正的可执行程序

假设上一步拿到 PID 为 `12345`：

```bash
PID=12345
readlink -f /proc/$PID/exe
```

典型结果：

```text
/usr/bin/qemu-aarch64-static
```

或者：

```text
/usr/libexec/qemu-binfmt/aarch64-binfmt-P
```

再看命令行参数：

```bash
tr '\0' ' ' < /proc/$PID/cmdline
echo
```

可能看到类似：

```text
/usr/bin/qemu-aarch64-static /home/ubuntu/simple-server
```

结论：

> 监听 8080 的业务程序逻辑来自 `~/simple-server`，但 Linux 进程真正加载并执行的宿主机可执行文件是 QEMU 的 aarch64 用户态模拟器。

也就是说，`simple-server` 没有直接在 x86_64 CPU 上运行。内核通过 `binfmt_misc` 识别出它是 ARM64 ELF，然后把它交给 `qemu-aarch64-static`。QEMU 作为实际进程，在用户态逐条翻译和执行 ARM64 指令，并把系统调用转换成当前 Linux 内核能够处理的系统调用。

## 6. 底层执行链路

整个过程可以按下面理解：

```text
用户执行 ~/simple-server
        |
        v
shell 调用 execve("~/simple-server", ...)
        |
        v
Linux 内核读取 ELF 头
        |
        v
发现架构是 AArch64，当前 CPU 是 x86_64
        |
        v
查询 binfmt_misc 注册表
        |
        v
匹配 qemu-aarch64 的 magic/mask
        |
        v
内核启动注册的解释器 qemu-aarch64-static
        |
        v
QEMU 加载并模拟执行 ~/simple-server
        |
        v
服务监听 8080
```

`binfmt_misc` 的关键不是“让 CPU 支持 ARM 指令”，而是“让内核知道遇到某种文件格式时该找谁解释执行”。

## 7. magic、mask 和 flags 怎么看

查看注册项：

```bash
cat /proc/sys/fs/binfmt_misc/qemu-aarch64
```

常见字段含义：

| 字段 | 含义 |
| :--- | :--- |
| `enabled` | 该格式处理器已启用 |
| `interpreter` | 匹配后内核要启动的解释器 |
| `magic` | 要匹配的文件头字节，ELF 文件一般以 `7f 45 4c 46` 开头 |
| `mask` | 匹配掩码，用来忽略 ELF 头里不需要固定匹配的字节 |
| `flags` | 控制内核把原文件、参数、解释器传递给 QEMU 的方式 |

常见 flags：

| flag | 说明 |
| :--- | :--- |
| `P` | preserve argv0，保留原始 `argv[0]` |
| `O` | open binary，内核打开目标二进制后把文件描述符传给解释器 |
| `C` | credentials，根据目标二进制计算凭据，和 setuid 等场景有关 |
| `F` | fix binary，注册时固定解释器文件，容器或 chroot 场景更稳定 |

实际看到哪些 flags 取决于发行版和包版本。

## 8. 排查命令汇总

```bash
# 看宿主机架构
uname -m

# 看目标二进制架构
file ~/simple-server

# 安装 QEMU 用户态模拟
sudo apt update
sudo apt install -y qemu-user-static binfmt-support

# 检查注册项
ls /proc/sys/fs/binfmt_misc/qemu-aarch64
cat /proc/sys/fs/binfmt_misc/qemu-aarch64

# 运行服务
~/simple-server &

# 找监听 8080 的 PID
sudo ss -ltnp 'sport = :8080'
sudo lsof -nP -iTCP:8080 -sTCP:LISTEN

# 识别真正执行的程序
PID=<上一步看到的PID>
readlink -f /proc/$PID/exe
tr '\0' ' ' < /proc/$PID/cmdline; echo
```

## 9. 这道题的答案应该怎么写

如果实验机输出显示 PID 是 `12345`，那么答案可以写成：

```text
监听 8080 端口的 PID 是 12345。
通过 readlink -f /proc/12345/exe 看到，真正的 executable 是 /usr/bin/qemu-aarch64-static
（或 /usr/libexec/qemu-binfmt/aarch64-binfmt-P，取决于系统）。

原因是 simple-server 是 ARM64 ELF，x86_64 内核不能直接执行它。
binfmt_misc 根据 ELF magic 匹配到 qemu-aarch64 注册项，于是内核实际启动 QEMU；
QEMU 再在用户态模拟执行 ~/simple-server。
```

所以这道题想考察的不是“哪个业务文件监听端口”，而是让你通过 `/proc/<PID>/exe` 观察到：**真正的 Linux 进程映像是 QEMU，ARM64 程序只是被 QEMU 加载和模拟执行的 guest 程序**。

## 10. Docker 中的 binfmt 与 QEMU

在 Docker 环境里，常见做法不是直接在宿主机上手动安装 `qemu-user-static`，而是运行一个特殊的特权容器：

```bash
docker run --privileged --rm tonistiigi/binfmt --install all
```

这个容器的作用不是长期运行服务，而是帮宿主机注册 `binfmt_misc` 处理器。它需要 `--privileged`，因为它要修改宿主机内核暴露出来的 `/proc/sys/fs/binfmt_misc`。

可以只安装 ARM64 支持：

```bash
docker run --privileged --rm tonistiigi/binfmt --install arm64
```

也可以安装所有常见架构：

```bash
docker run --privileged --rm tonistiigi/binfmt --install all
```

注册完成后检查：

```bash
cat /proc/sys/fs/binfmt_misc/qemu-aarch64
```

重点看：

```text
enabled
flags: ...F...
```

其中 `F` 是 `fix_binary`，Docker 官方文档要求 QEMU 解释器静态编译，并且注册项需要带 `fix_binary`，这样跨架构构建和容器环境下的执行才更稳定。

整个关系可以理解为：

```text
docker run --privileged tonistiigi/binfmt --install arm64
        |
        v
修改宿主机 /proc/sys/fs/binfmt_misc/qemu-aarch64
        |
        v
docker run --platform linux/arm64 ...
        |
        v
容器里启动 ARM64 程序
        |
        v
宿主机内核发现是 ARM64 ELF
        |
        v
转交给 qemu-aarch64 执行
```

注意：容器共享宿主机内核，所以 `tonistiigi/binfmt` 容器注册的是宿主机内核的 `binfmt_misc`，不是只影响这个安装容器自己。

## 11. 在 Docker 中运行 ARM64 镜像

安装 binfmt 后，可以在 x86_64 机器上运行 ARM64 容器：

```bash
docker run --rm --platform linux/arm64 alpine uname -m
```

如果成功，输出类似：

```text
aarch64
```

这表示 Docker 启动的是 ARM64 用户空间，底层由 QEMU 在 x86_64 宿主机上模拟执行。

如果要运行前面的 `~/simple-server`，并且它是静态链接的 ARM64 Linux 二进制，可以挂载进 ARM64 容器：

```bash
docker run --rm \
  --platform linux/arm64 \
  -p 8080:8080 \
  -v "$HOME/simple-server:/simple-server:ro" \
  alpine \
  /simple-server
```

如果报缺少动态链接器或 libc，说明这个二进制不是完全静态链接的，需要使用带有匹配 ARM64 动态库的基础镜像，例如：

```bash
docker run --rm \
  --platform linux/arm64 \
  -p 8080:8080 \
  -v "$HOME/simple-server:/simple-server:ro" \
  arm64v8/ubuntu \
  /simple-server
```

然后在宿主机访问：

```bash
curl http://127.0.0.1:8080/
```

## 12. buildx 构建跨架构镜像

Docker 中“构建跨架构镜像”通常使用 `buildx`。它构建的不是一个二进制层同时适配所有 CPU，而是一个多平台镜像索引：

```text
yourname/myapp:1.0
        |
        +-- linux/amd64 镜像
        +-- linux/arm64 镜像
```

用户拉取 `yourname/myapp:1.0` 时，Docker 会根据当前机器架构自动选择合适的镜像。

这里要先分清楚：

| 概念 | 作用 |
| :--- | :--- |
| `buildx` | 负责调度构建、生成多平台镜像和 manifest |
| QEMU/binfmt | 让构建过程能模拟执行外架构程序 |
| 交叉编译 | 不模拟外架构程序，而是用本机架构编译器直接产出目标架构二进制 |

所以 `buildx` 是构建入口，QEMU 和交叉编译是两种不同的实现路线。

### 12.1 准备 buildx builder

创建并启用 builder：

```bash
docker buildx create --name multi-builder --driver docker-container --use --bootstrap
docker buildx inspect --bootstrap
```

`--driver docker-container` 会让 BuildKit 运行在一个独立容器中，功能比默认 Docker driver 更完整，适合多平台构建。

如果后续采用 QEMU 模拟构建路线，再注册 QEMU/binfmt：

```bash
docker run --privileged --rm tonistiigi/binfmt --install all
```

如果后续采用 Go 这类交叉编译路线，并且构建阶段不需要执行目标架构程序，理论上可以不依赖 QEMU。

### 12.2 构建并推送多平台镜像

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t yourname/myapp:1.0 \
  --push \
  .
```

这里推荐使用 `--push`，因为多平台镜像最自然的存放位置是镜像仓库，例如 Docker Hub、Harbor、GHCR。

构建完成后检查：

```bash
docker buildx imagetools inspect yourname/myapp:1.0
```

应该能看到类似：

```text
Platform: linux/amd64
Platform: linux/arm64
```

如果只想构建 ARM64 镜像：

```bash
docker buildx build \
  --platform linux/arm64 \
  -t yourname/myapp:arm64 \
  --push \
  .
```

## 13. `--platform` 的基础规则

在 Docker 构建里，`--platform` 不是普通 Dockerfile `ARG`，也不是 `--build-arg`。它是 Docker 构建器接收的参数，用来说明“这次构建的目标镜像平台”。

例如：

```bash
docker build --platform=linux/arm64 -t demo .
```

或者：

```bash
docker buildx build --platform=linux/arm64 -t demo .
```

这两条命令都会告诉构建器：

```text
本次构建目标平台 = linux/arm64
```

它会影响：

1. 最终镜像的平台元数据。
2. 没写 `--platform` 的 `FROM` 默认使用哪个架构的基础镜像。
3. BuildKit 自动注入的预定义 platform args，例如 `TARGETOS`、`TARGETARCH`。

### 13.1 `FROM --platform` 的覆盖规则

`FROM` 可以显式指定当前 stage 使用哪个架构的基础镜像：

```dockerfile
FROM --platform=linux/amd64 golang:alpine AS build
```

这表示当前 stage 强制使用 `linux/amd64` 的 `golang:alpine`。

如果 `FROM` 没写 `--platform`：

```dockerfile
FROM alpine
```

它就继承构建命令里的目标平台。比如：

```bash
docker build --platform=linux/arm64 -t demo .
```

那么：

```dockerfile
FROM alpine
```

等价于：

```dockerfile
FROM --platform=linux/arm64 alpine
```

优先级：

```text
FROM --platform=xxx      最高，只影响当前 stage
docker build --platform  默认目标平台，影响没写 --platform 的 FROM
builder 默认平台          没传 --platform 时使用
```

### 13.2 `docker build` 和 `docker buildx build`

二者都可以接收 `--platform`，但定位不同：

| 命令 | 适合场景 | 说明 |
| :--- | :--- | :--- |
| `docker build --platform=linux/arm64` | 单平台构建 | 构建一个目标平台为 ARM64 的镜像 |
| `docker buildx build --platform=linux/arm64` | 单平台 BuildKit 构建 | 和上面效果接近，但走 buildx 入口 |
| `docker buildx build --platform=linux/amd64,linux/arm64` | 多平台构建 | 分别构建多个平台，并生成 manifest list |

可以这样记：

```text
docker build --platform:
  我这次要构建某一个平台的镜像。

docker buildx build --platform:
  我用 BuildKit 构建某一个或多个平台的镜像。
```

如果只是构建单个 ARM64 镜像：

```bash
docker build --platform=linux/arm64 -t demo:arm64 .
```

如果要构建多平台镜像：

```bash
docker buildx build \
  --platform=linux/amd64,linux/arm64 \
  -t yourname/demo:1.0 \
  --push \
  .
```

### 13.3 BuildKit 自动注入的 platform args

BuildKit 会根据 `--platform` 自动提供一些预定义参数：

| 变量 | 含义 | 示例 |
| :--- | :--- | :--- |
| `BUILDPLATFORM` | 执行构建的机器平台 | `linux/amd64` |
| `BUILDOS` | 执行构建的机器操作系统 | `linux` |
| `BUILDARCH` | 执行构建的机器 CPU 架构 | `amd64` |
| `TARGETPLATFORM` | 目标镜像平台 | `linux/arm64` |
| `TARGETOS` | 目标操作系统 | `linux` |
| `TARGETARCH` | 目标 CPU 架构 | `arm64` |
| `TARGETVARIANT` | 目标架构变体 | `v7` 等 |

这些变量可以在 `FROM` 中直接使用：

```dockerfile
FROM --platform=$BUILDPLATFORM golang:alpine AS build
```

如果要在某个 stage 的 `RUN` 里使用，需要先声明 `ARG`：

```dockerfile
ARG TARGETOS
ARG TARGETARCH
RUN echo "$TARGETOS $TARGETARCH"
```

这里的 `ARG TARGETOS` 和 `ARG TARGETARCH` 不是在设置默认值，而是把 BuildKit 自动注入的目标平台参数引入当前 stage 的作用域。

例如：

```bash
docker build --platform=linux/arm64 .
```

对应：

```text
TARGETOS=linux
TARGETARCH=arm64
TARGETPLATFORM=linux/arm64
```

## 14. 路线一：QEMU 模拟构建

QEMU 路线的思路是：目标平台是什么，构建阶段和运行阶段就都使用目标平台的基础镜像。如果构建过程中需要执行外架构程序，就交给 QEMU 模拟执行。

例如在 amd64 机器上构建 ARM64 镜像：

```bash
docker buildx build \
  --platform=linux/arm64 \
  -t demo:arm64 \
  --load \
  .
```

如果 Dockerfile 是：

```dockerfile
FROM alpine
RUN apk add --no-cache curl
CMD ["sh"]
```

由于 `FROM alpine` 没有显式写 `--platform`，它会继承构建目标平台：

```text
FROM alpine -> linux/arm64 alpine
```

那么 `RUN apk add ...` 执行的是 ARM64 Alpine 里的 `/bin/sh` 和 `apk`。在 amd64 宿主机上，这一步需要 QEMU/binfmt 支持，否则很容易报：

```text
exec format error
```

所以 QEMU 路线通常先做：

```bash
docker run --privileged --rm tonistiigi/binfmt --install all
```

再构建：

```bash
docker buildx build \
  --platform=linux/arm64 \
  -t demo:arm64 \
  --load \
  .
```

多平台构建时：

```bash
docker buildx build \
  --platform=linux/amd64,linux/arm64 \
  -t yourname/demo:1.0 \
  --push \
  .
```

QEMU 路线的特点：

| 特点 | 说明 |
| :--- | :--- |
| Dockerfile 改动少 | 大部分普通 Dockerfile 可以直接用 |
| 通用性好 | apt、apk、npm、pip、安装脚本等都按目标架构运行 |
| 速度较慢 | 外架构程序通过 QEMU 模拟执行 |
| 仍可能有兼容性问题 | 少数编译、测试或系统调用场景可能失败 |

这条路线适合：

- 项目没有清晰的交叉编译能力。
- 构建过程依赖目标架构系统包。
- Python/Node native addon、复杂 C/C++ 依赖、安装脚本较多。
- 想先快速得到一个能工作的多架构镜像。

## 15. 路线二：交叉编译构建

交叉编译路线的思路是：构建阶段使用构建机本身架构的工具链，让编译器直接产出目标架构的二进制；运行阶段再使用目标架构的基础镜像。

这条路线和 QEMU 的区别很大：

```text
QEMU 模拟构建：
  build stage 使用目标架构镜像
  RUN 阶段执行目标架构程序
  依赖 QEMU 模拟

交叉编译构建：
  build stage 使用构建机架构镜像
  编译器原生运行
  编译产物是目标架构
  通常不依赖 QEMU
```

### 15.1 Go 的典型写法

```dockerfile
# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o server .

FROM alpine
COPY --from=build /app/server /server
ENTRYPOINT ["/server"]
```

在 amd64 机器上执行：

```bash
docker build --platform=linux/arm64 -t demo:arm64 .
```

变量关系是：

```text
BUILDPLATFORM=linux/amd64
TARGETPLATFORM=linux/arm64
TARGETOS=linux
TARGETARCH=arm64
```

所以第一阶段：

```dockerfile
FROM --platform=$BUILDPLATFORM golang:alpine AS build
```

等价于：

```dockerfile
FROM --platform=linux/amd64 golang:alpine AS build
```

构建阶段的 Go 编译器是 amd64 程序，可以在 amd64 宿主机上原生运行。

然后：

```dockerfile
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o server .
```

等价于：

```dockerfile
RUN GOOS=linux GOARCH=arm64 go build -o server .
```

产物 `/app/server` 是 ARM64 Linux 二进制。

最后：

```dockerfile
FROM alpine
```

没有写 `--platform`，所以继承构建目标平台 `linux/arm64`，也就是 ARM64 Alpine。最终镜像里放的是 ARM64 Alpine + ARM64 server。

完整流程：

```text
docker build --platform=linux/arm64
        |
        v
目标镜像平台 = linux/arm64

FROM --platform=$BUILDPLATFORM golang:alpine AS build
        |
        v
build stage = linux/amd64 golang:alpine
Go 编译器原生运行
GOOS=linux GOARCH=arm64 go build
产出 ARM64 server

FROM alpine
        |
        v
runtime stage = linux/arm64 alpine
COPY ARM64 server 进去
最终镜像 = linux/arm64
```

### 15.2 多平台交叉编译

同一个 Dockerfile 可以构建多平台镜像：

```bash
docker buildx build \
  --platform=linux/amd64,linux/arm64 \
  -t yourname/server:1.0 \
  --push \
  .
```

BuildKit 会为每个目标平台分别执行一次构建：

```text
目标 linux/amd64:
  TARGETOS=linux
  TARGETARCH=amd64
  go build -> amd64 server
  runtime FROM alpine -> amd64 alpine

目标 linux/arm64:
  TARGETOS=linux
  TARGETARCH=arm64
  go build -> arm64 server
  runtime FROM alpine -> arm64 alpine
```

最终 `yourname/server:1.0` 是一个多平台 manifest：

```text
yourname/server:1.0
        |
        +-- linux/amd64 镜像
        +-- linux/arm64 镜像
```

### 15.3 为什么不能只复制一个固定架构二进制

如果 Dockerfile 只是复制一个已经编译好的 ARM64 二进制：

```dockerfile
FROM alpine
COPY simple-server /simple-server
CMD ["/simple-server"]
```

那么它只适合构建 `linux/arm64` 镜像。因为 `simple-server` 本身就是 ARM64 的，不能直接放进 `linux/amd64` 镜像里当 amd64 程序运行。

要构建多平台镜像，必须保证每个平台的最终镜像里都有对应架构的程序：

```text
linux/amd64 镜像里放 amd64 server
linux/arm64 镜像里放 arm64 server
```

实现方式有两种：

1. 在 Dockerfile 里根据 `TARGETARCH` 交叉编译。
2. 预先准备不同架构的二进制，然后按 `TARGETARCH` 选择复制。

Go 项目通常选择第一种。

### 15.4 哪些项目适合交叉编译

| 项目类型 | 交叉编译可行性 | 说明 |
| :--- | :--- | :--- |
| Go | 很高 | `GOOS/GOARCH` 原生支持，最适合这种模式 |
| Rust | 较高 | 需要配置 target，涉及 native 依赖时会复杂 |
| C/C++ | 中等 | 需要 cross compiler、sysroot、目标架构库 |
| Zig | 较高 | 对交叉编译支持较好 |
| Java | 特殊 | 字节码跨平台，但基础镜像/JVM 仍需目标架构 |
| Node/Python | 取决于依赖 | 纯 JS/Python 简单，有 native addon/wheel 时复杂 |

交叉编译路线的优缺点：

| 特点 | 说明 |
| :--- | :--- |
| 速度快 | 编译器在构建机上原生运行 |
| 不依赖 QEMU | 适合 CI 中高频构建 |
| Dockerfile 需要配合 | 要显式处理 `TARGETOS/TARGETARCH` |
| 对依赖要求高 | native library、CGO、系统库会增加复杂度 |

## 16. 官方文档中的多平台构建策略

Docker 官方文档把 multi-platform build 定义为：一次构建同时面向多个 OS/CPU 组合，例如 `linux/amd64`、`linux/arm64`、`windows/amd64`。最终推到 registry 的不是一个单一镜像，而是一个 manifest list，里面指向多个不同平台的镜像变体。拉取镜像时，Docker 会根据当前主机架构自动选择合适的变体。

官方文档给出的三种策略是：

| 路线 | 核心思路 | 适合场景 |
| :--- | :--- | :--- |
| QEMU 模拟构建 | 目标架构程序通过 QEMU 跑 | 通用、改动少、先跑起来 |
| 多节点原生构建 | amd64 节点构建 amd64，arm64 节点构建 arm64 | 生产 CI、复杂构建、速度和兼容性要求高 |
| 交叉编译构建 | 本机架构编译器产出目标架构二进制 | Go/Rust 等支持良好的项目，追求速度 |

实践建议：

```text
能稳定交叉编译：优先交叉编译
交叉编译成本高：使用 QEMU + buildx
追求构建速度和真实环境：使用多架构原生 builder 节点
```

Docker Build Cloud 不是第四种底层技术路线，它可以理解为“托管版多节点原生构建”。你不用自己维护 amd64/arm64 builder 集群，由 Docker 提供托管 builder 和共享缓存。

参考：Docker 官方文档 Multi-platform builds: <https://docs.docker.com/build/building/multi-platform/>

## 17. 路线三：远程 builder / 多节点原生构建

远程 builder 的核心思路是：`docker buildx build` 命令仍然在本机执行，但真正跑 Dockerfile 的 BuildKit worker 在远程机器上。

```text
dev-01 本机
  docker buildx build
        |
        | SSH / Docker context
        v
builder-01 远程机器
  BuildKit 真正执行 Dockerfile
        |
        v
构建 linux/arm64 镜像
```

如果远程机器原生就是 ARM64，那么构建 ARM64 镜像时不需要 QEMU，也不需要改 Dockerfile 做交叉编译。它是最接近真实目标环境的方式，通常也比 QEMU 快。

### 17.1 多节点 builder 的官方形态

Docker 官方文档展示的是先准备多个 Docker context，然后把它们合并到同一个 buildx builder：

```bash
docker buildx create --use --name mybuild node-amd64
docker buildx create --append --name mybuild node-arm64
docker buildx build --platform linux/amd64,linux/arm64 .
```

这里的含义是：

```text
node-amd64 负责 linux/amd64 构建
node-arm64 负责 linux/arm64 构建
buildx 负责把它们组织成一个多平台 manifest
```

这种方式适合生产 CI：

- amd64 构建在真实 amd64 节点上跑。
- arm64 构建在真实 arm64 节点上跑。
- 不依赖 QEMU，兼容性更好。
- 编译、压缩、测试这类 CPU-heavy 步骤更快。

### 17.2 通过 SSH 创建远程 builder

类似 iximiuz 这类实验环境，经常给一台远程 builder：

```text
dev-01      本机 amd64 开发机
builder-01 远程构建机，通过 SSH 访问
```

可以直接用 SSH endpoint 创建 builder：

```bash
docker buildx create \
  --name remote-arm-builder \
  --driver docker-container \
  --use \
  --bootstrap \
  ssh://laborant@builder-01
```

也可以先创建 Docker context，再用这个 context 创建 buildx builder：

```bash
docker context create builder-01 \
  --docker host=ssh://laborant@builder-01

docker buildx create \
  --name remote-arm-builder \
  --driver docker-container \
  --use \
  --bootstrap \
  builder-01
```

检查 builder：

```bash
docker buildx inspect --bootstrap
```

重点看：

```text
Endpoint:  ssh://laborant@builder-01
Platforms: linux/arm64, ...
```

只要 `Platforms` 里包含 `linux/arm64`，本机是否安装 QEMU 对这个构建就不重要了，因为真正执行构建的是远程 builder。

### 17.3 构建、加载、推送远程 builder 产物

假设应用在 `~/app`：

```bash
cd ~/app
```

构建 ARM64 镜像并加载到本机 Docker daemon：

```bash
docker buildx build \
  --builder remote-arm-builder \
  --platform linux/arm64 \
  -t registry.iximiuz.com/app:v1.0.0 \
  --load \
  .
```

然后从本机推送到 registry：

```bash
docker push registry.iximiuz.com/app:v1.0.0
```

也可以直接让 buildx 推送：

```bash
docker buildx build \
  --builder remote-arm-builder \
  --platform linux/arm64 \
  -t registry.iximiuz.com/app:v1.0.0 \
  --push \
  .
```

如果题目同时要求“load 到本地 Docker daemon”和“push 到 registry”，最稳妥的是分两步：

```text
先 --load
再 docker push
```

因为不同 buildx driver / Docker 版本对 `--load` 和 `--push` 同时使用的支持不完全一致。

### 17.4 Docker Build Cloud

Docker Build Cloud 是 Docker 官方提供的托管 builder。它本质上属于“多节点原生构建”路线，只是节点由 Docker 管理：

```bash
docker buildx create --driver cloud <ORG>/<BUILDER_NAME>

docker build \
  --builder cloud-<ORG>-<BUILDER_NAME> \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  --tag <IMAGE_NAME> \
  --push .
```

它适合：

- 不想自己维护 amd64/arm64 builder 集群。
- 构建任务 CPU 开销大，QEMU 太慢。
- 团队希望共享远程 build cache。
- CI 中需要稳定构建多平台镜像。

Docker 官方文档里的 Neovim 示例就是这条路线：用 Docker Build Cloud 编译并导出 `linux/amd64` 和 `linux/arm64` 两个平台的二进制。

## 18. 四个概念的关系

严格按 Docker 官方策略，底层是三种：

```text
1. QEMU emulation
2. Multiple native nodes
3. Cross-compilation
```

结合实际使用，可以把它们扩展理解成下面几类：

| 名称 | 底层路线 | 构建在哪里跑 | 是否需要 QEMU |
| :--- | :--- | :--- | :--- |
| 本机 QEMU 构建 | QEMU emulation | 本机 builder | 需要 |
| 本机交叉编译 | Cross-compilation | 本机 builder | 通常不需要 |
| 自建远程 builder | Multiple native nodes | 远程 Docker/BuildKit | 如果远程节点原生支持目标架构，则不需要 |
| Docker Build Cloud | Managed native nodes | Docker 托管 builder | 不需要用户维护 QEMU/节点 |

选择建议：

| 场景 | 推荐路线 |
| :--- | :--- |
| 想最快跑通，Dockerfile 不想改 | QEMU |
| Go/Rust 等容易交叉编译 | 交叉编译 |
| Node/Python/C++ 依赖复杂，QEMU 又慢 | 远程原生 builder |
| 团队 CI 长期构建多平台镜像 | 多节点 builder 或 Docker Build Cloud |
| 不想维护构建集群 | Docker Build Cloud |

## 19. `tonistiigi/binfmt` 和 buildx 的关系

可以这样区分：

| 工具 | 作用 |
| :--- | :--- |
| `tonistiigi/binfmt` | 注册 QEMU/binfmt，让宿主机内核知道如何执行外架构 ELF |
| `docker buildx` | 使用 BuildKit 构建单平台或多平台镜像 |
| `--platform` | 指定要运行或构建的平台，例如 `linux/arm64` |
| QEMU | 在用户态模拟执行外架构程序 |

也就是说：

```text
tonistiigi/binfmt 解决“能不能执行外架构程序”
buildx 解决“如何构建多平台镜像”
--platform 解决“这次要运行或构建哪个平台”
```

一次典型流程：

```bash
# 1. 注册跨架构执行支持
docker run --privileged --rm tonistiigi/binfmt --install all

# 2. 创建 buildx builder
docker buildx create --name multi-builder --driver docker-container --use --bootstrap

# 3. 构建并推送多架构镜像
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t yourname/app:1.0 \
  --push \
  .

# 4. 查看多架构 manifest
docker buildx imagetools inspect yourname/app:1.0
```

一句话总结：`tonistiigi/binfmt` 是 Docker 跨架构场景里用来注册 QEMU 模拟器的工具容器；`buildx` 是实际构建多架构镜像的工具；Docker 官方文档推荐这种方式，但 `tonistiigi/binfmt` 本身属于 BuildKit/buildx 生态中的常用项目。
