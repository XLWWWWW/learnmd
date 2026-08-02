# Docker 多阶段构建实践笔记

> 专题导航：[容器镜像文件系统专题索引](./容器镜像文件系统专题索引.md)

本文整理自 Ivan Velichko 的《How to Build Smaller Container Images: Docker Multi-Stage Builds》，并补充 Dockerfile 优化、不同技术栈模板、BuildKit 缓存、安全和日常排查经验。

这篇文章的核心观点很直接：

> 构建环境和运行环境是两个不同问题。生产镜像里应该只放运行应用真正需要的东西，不应该把编译器、构建工具、测试工具、开发依赖一起带到线上。

Multi-Stage Build 的价值就是在一个 Dockerfile 中完成“构建阶段”和“运行阶段”的分离，只把最终产物复制到生产镜像里。

---

## 1. 对这篇博客的评价

### 1.1 优点

这篇文章讲 Multi-Stage Build 的路径很清楚：

1. 先解释镜像为什么变大。
2. 再用 Go 和 Node.js 展示单阶段 Dockerfile 的问题。
3. 然后回顾 Multi-Stage Build 出现前的 Builder Pattern。
4. 最后拆解 Multi-Stage Build 的两个基础能力：多个 `FROM` 和 `COPY --from`。

它最好的地方是把“镜像瘦身”这件事讲成了一个工程边界问题：

```text
build-time dependencies != run-time dependencies
```

也就是：

- 构建时需要编译器、SDK、包管理器、开发依赖、测试工具。
- 运行时通常只需要二进制、bundle、生产依赖、运行库和配置。

如果把这两类依赖混在同一个镜像里，镜像一定会膨胀，漏洞面也会变大。

### 1.2 不足和需要补充的地方

文章的主线适合入门，但如果要真正写好生产 Dockerfile，还需要补上这些内容：

- `.dockerignore` 对构建上下文大小的影响。
- Docker layer 缓存怎么组织，尤其是依赖安装层和源码复制层的顺序。
- BuildKit 的 `--mount=type=cache`、`--mount=type=secret`。
- `scratch`、`distroless`、`alpine`、`debian-slim` 的取舍。
- 非 root 用户运行、只读文件系统、最小权限等安全实践。
- 如何验证镜像大小、层大小、CVE、运行时依赖是否完整。
- 不同语言栈更完整的 Dockerfile 模板。

结论：

> 这篇博客适合建立 Multi-Stage Build 的核心直觉，但生产使用时要和缓存优化、安全加固、基础镜像选择、依赖管理一起考虑。

---

## 2. 为什么镜像会变大

### 2.1 构建依赖和运行依赖混在一起

几乎所有应用都有两类依赖：

| 类型 | 例子 | 是否应该进入生产镜像 |
| :--- | :--- | :--- |
| 构建时依赖 | 编译器、SDK、构建工具、头文件、dev dependencies | 通常不应该 |
| 运行时依赖 | 二进制、bundle、生产依赖、运行库、证书 | 应该 |

常见构建时依赖：

- Go：`go` 编译器、module 下载缓存、测试工具。
- Node.js：`npm`、`pnpm`、`devDependencies`、构建工具链。
- Python：`gcc`、`build-essential`、`python-dev`、`*-dev` 头文件。
- Rust：`cargo`、`rustc`、registry 缓存。
- Java：JDK、Maven、Gradle、本地仓库。
- 前端：Vite、Webpack、TypeScript、ESLint、测试工具。

生产镜像里通常不需要这些东西。

### 2.2 单阶段 Dockerfile 的典型问题

很多镜像膨胀都来自一个习惯：

```text
用同一个镜像 build，也用同一个镜像 run。
```

这样写起来简单，但会把构建工具和构建缓存一起带到生产镜像。

---

## 3. 反例：Go 单阶段 Dockerfile

错误示例：

```Dockerfile
FROM golang:1.23

WORKDIR /app
COPY . .

RUN go build -o binary

CMD ["/app/binary"]
```

问题：

- `golang:1.23` 是构建镜像，不是生产运行镜像。
- 最终镜像里会包含 Go 编译器、标准库源码、构建工具链、系统包。
- 你的应用可能只有几十 MB，但最终镜像可能几百 MB 甚至更大。
- 编译器和系统包越多，CVE 数量和攻击面越大。

更合理的思路：

```text
build stage: 使用 golang 镜像编译
runtime stage: 只复制编译后的 binary
```

---

## 4. 反例：Node.js 单阶段 Dockerfile

错误示例：

```Dockerfile
FROM node:22-slim

WORKDIR /app
COPY . .

RUN npm ci
RUN npm run build

ENV NODE_ENV=production
EXPOSE 3000

CMD ["node", "/app/.output/index.mjs"]
```

这个例子比 Go 更隐蔽。

`node:22-slim` 本身可以作为生产运行镜像，但问题出在：

```Dockerfile
RUN npm ci
RUN npm run build
```

`npm ci` 默认会安装生产依赖和开发依赖。`npm run build` 又往往需要开发依赖，比如 TypeScript、Vite、Webpack、Nuxt、Next、SvelteKit 等构建工具。

所以不能简单改成：

```bash
npm ci --omit=dev
```

因为这样后面的构建命令可能会失败。

真正的问题不是 Node 镜像本身，而是：

```text
构建阶段需要 devDependencies
运行阶段不应该保留 devDependencies
```

解决方式仍然是拆阶段：

```text
build stage: npm ci + npm run build
runtime stage: 只复制 .output / dist / build 等运行产物
```

---

## 5. Multi-Stage Build 出现前：Builder Pattern

在 Docker 支持 Multi-Stage Build 之前，常见做法是写两个 Dockerfile。

### 5.1 构建 Dockerfile

例如：

```Dockerfile
# Dockerfile.build
FROM node:22-slim

WORKDIR /app
COPY . .

RUN npm ci
RUN npm run build
```

构建辅助镜像：

```bash
docker build -t build:v1 -f Dockerfile.build .
```

然后从构建容器里复制产物到宿主机：

```bash
docker cp "$(docker create build:v1):/app/.output" .
```

### 5.2 运行 Dockerfile

再写一个运行镜像：

```Dockerfile
# Dockerfile.run
FROM node:22-slim

WORKDIR /app
COPY .output ./.output

ENV NODE_ENV=production
EXPOSE 3000

CMD ["node", "/app/.output/index.mjs"]
```

构建最终镜像：

```bash
docker build -t app:v1 -f Dockerfile.run .
```

### 5.3 Builder Pattern 的问题

这种方式可行，但体验很差：

- 需要维护多个 Dockerfile。
- 需要额外脚本保证构建顺序。
- 需要把构建产物先复制到宿主机。
- 容易把上一次的旧产物打进镜像。
- 团队之间很难形成统一标准。

Multi-Stage Build 本质上就是 Docker 原生支持的 Builder Pattern。

---

## 6. Multi-Stage Build 的两个基础能力

### 6.1 `COPY --from` 可以从别的镜像复制文件

普通 `COPY` 是从宿主机构建上下文复制到镜像：

```Dockerfile
COPY ./app /app
```

但 Dockerfile 也支持从其他镜像复制文件：

```Dockerfile
COPY --from=nginx:latest /etc/nginx/nginx.conf /nginx.conf
```

这表示：

```text
从 nginx:latest 镜像里复制 /etc/nginx/nginx.conf 到当前正在构建的镜像。
```

这个能力可以绕过“先 docker cp 到宿主机，再 COPY 进运行镜像”的旧流程。

### 6.2 一个 Dockerfile 可以有多个 `FROM`

Dockerfile 里每个 `FROM` 都会开启一个新的 stage。

示例：

```Dockerfile
FROM busybox:stable AS from1
CMD ["echo", "busybox"]

FROM alpine:3 AS from2
CMD ["echo", "alpine"]

FROM debian:stable-slim AS from3
CMD ["echo", "debian"]
```

可以构建指定 stage：

```bash
docker build --target from1 -t my-busybox .
docker build --target from2 -t my-alpine .
docker build --target from3 -t my-debian .
```

如果不指定 `--target`，默认构建最后一个 stage。

### 6.3 Multi-Stage Build 的组合

把多个 `FROM` 和 `COPY --from` 组合起来，就是 Multi-Stage Build。

典型结构：

```Dockerfile
FROM sdk-image AS build

WORKDIR /src
COPY . .
RUN build-command


FROM runtime-image AS runtime

WORKDIR /app
COPY --from=build /src/artifact /app/artifact
CMD ["/app/artifact"]
```

最终镜像只包含最后一个 stage 的内容，以及显式 `COPY --from` 复制过来的文件。

---

## 7. Multi-Stage Build 的关键规则

### 7.1 每个 `FROM` 都是新 stage

例如：

```Dockerfile
FROM golang:1.23 AS build
RUN go version

FROM debian:stable-slim AS runtime
RUN which go
```

第二个 stage 是全新的文件系统，不会自动继承第一个 stage 的文件。

如果需要第一个 stage 的产物，必须显式复制：

```Dockerfile
COPY --from=build /path/in/build /path/in/runtime
```

### 7.2 stage 名称建议总是写

虽然可以用数字引用 stage：

```Dockerfile
COPY --from=0 /app/binary /app/binary
```

但不推荐。更推荐命名：

```Dockerfile
FROM golang:1.23 AS build

FROM debian:stable-slim AS runtime
COPY --from=build /app/binary /app/binary
```

好处：

- 可读性更好。
- 调整 stage 顺序时不容易出错。
- CI、调试、文档说明更清楚。

### 7.3 默认构建最后一个 stage

下面这个 Dockerfile：

```Dockerfile
FROM node:22-slim AS deps

FROM node:22-slim AS build

FROM node:22-slim AS runtime
```

执行：

```bash
docker build -t app .
```

默认构建的是 `runtime` stage。

如果只想构建到 `build`：

```bash
docker build --target build -t app-build .
```

### 7.4 BuildKit 会按依赖关系构建

现代 Docker 默认使用 BuildKit。BuildKit 会根据 `COPY --from` 的依赖关系计算构建图：

- 没有被最终 stage 依赖的 stage 可以跳过。
- 独立 stage 可以并发构建。
- 缓存命中时可以跳过重复步骤。

这也是 Multi-Stage Build 比手写 Builder Pattern 更适合标准化的原因。

### 7.5 生产 stage 建议放最后

日常写法通常是：

```text
base/deps -> build/test -> runtime
```

也就是把最终生产镜像对应的 stage 放在最后。这样执行：

```bash
docker build -t app .
```

默认就会构建生产镜像。

如果把调试 stage 或测试 stage 放在最后，忘记使用 `--target` 时就可能构建出错误镜像。

另外，虽然 Dockerfile 会被整体解析，但工程上仍建议：

```text
生产者 stage 写在前面
消费者 stage 写在后面
```

例如先定义 `build`，再在 `runtime` 里 `COPY --from=build`。这样最符合阅读顺序，也最不容易踩到不同构建器实现差异。

---

## 8. Node.js 多阶段构建

Node.js 应用分两类：

| 类型 | 运行时是否需要 Node.js |
| :--- | :--- |
| SSR 服务，例如 Next.js、Nuxt、SvelteKit server | 需要 |
| 纯静态站点，例如 Vite build 后的静态文件 | 不一定需要，可以用 nginx/caddy |

### 8.1 通用 Node.js SSR 模板

以 Nuxt/Next/SvelteKit 这类构建后仍需要 Node.js 运行的应用为例：

```Dockerfile
FROM node:22-slim AS deps

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci


FROM node:22-slim AS build

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build


FROM node:22-slim AS runtime

WORKDIR /app
ENV NODE_ENV=production

COPY --from=build /app/.output ./.output
COPY package.json ./

EXPOSE 3000
CMD ["node", ".output/index.mjs"]
```

重点：

- `deps` stage 负责安装完整依赖，包括 devDependencies。
- `build` stage 负责构建产物。
- `runtime` stage 只复制运行产物。
- 如果构建产物已经自包含，就不复制 `node_modules`。

### 8.2 Next.js standalone 模板

Next.js 推荐生产容器使用 standalone 输出。

`next.config.js` 中通常需要：

```js
module.exports = {
  output: "standalone",
};
```

Dockerfile 示例：

```Dockerfile
FROM node:22-slim AS deps

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci


FROM node:22-slim AS build

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build


FROM node:22-slim AS runtime

WORKDIR /app
ENV NODE_ENV=production

COPY --from=build /app/.next/standalone ./
COPY --from=build /app/.next/static ./.next/static
COPY --from=build /app/public ./public

EXPOSE 3000
CMD ["node", "server.js"]
```

这里不要直接复制整个项目目录。standalone 模式会把运行需要的服务端文件和依赖裁剪出来，复制 `.next/standalone` 通常比复制完整 `node_modules` 更干净。

### 8.3 需要生产 `node_modules` 的模板

有些 Node.js 服务不是完整 bundle，运行时还需要生产依赖。

可以这样：

```Dockerfile
FROM node:22-slim AS deps

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci


FROM node:22-slim AS build

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build


FROM node:22-slim AS prod-deps

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci --omit=dev


FROM node:22-slim AS runtime

WORKDIR /app
ENV NODE_ENV=production

COPY --from=prod-deps /app/node_modules ./node_modules
COPY --from=build /app/dist ./dist
COPY package.json ./

EXPOSE 3000
CMD ["node", "dist/index.js"]
```

这里分成两套依赖：

```text
deps      -> 构建用完整依赖
prod-deps -> 运行用生产依赖
```

### 8.4 pnpm 模板

```Dockerfile
FROM node:22-slim AS base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app


FROM base AS deps

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile


FROM base AS build

COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN pnpm build


FROM base AS runtime

ENV NODE_ENV=production
COPY --from=build /app/.output ./.output
EXPOSE 3000
CMD ["node", ".output/index.mjs"]
```

### 8.5 静态前端模板

例如 Vite/React/Vue 构建成静态文件后，用 nginx 运行：

```Dockerfile
FROM node:22-slim AS build

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build


FROM nginx:1.27-alpine AS runtime

COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

---

## 9. Go 多阶段构建

Go 应用总要编译。运行阶段怎么选 base image，主要取决于二进制是静态链接还是动态链接。

### 9.1 静态链接 Go 应用

```Dockerfile
FROM golang:1.23 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/app


FROM gcr.io/distroless/static-debian12 AS runtime

COPY --from=build /out/app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]
```

重点：

- `CGO_ENABLED=0` 生成静态链接二进制。
- `-trimpath` 去掉本地路径信息。
- `-ldflags="-s -w"` 减小符号表和调试信息。
- `distroless/static` 比 `golang` 小得多，也比 `scratch` 更稳妥。

### 9.2 使用 `scratch`

极简版本：

```Dockerfile
FROM golang:1.23 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd/app


FROM scratch AS runtime

COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
```

`scratch` 要谨慎使用，因为它什么都没有。

常见缺失：

- CA 证书：HTTPS 请求可能失败。
- `/etc/passwd`：非 root 用户信息没有。
- timezone 数据。
- shell：无法 `sh` 进去调试。

如果程序需要 HTTPS，可以补证书：

```Dockerfile
FROM golang:1.23 AS build

WORKDIR /src
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd/app


FROM scratch AS runtime

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
```

### 9.3 动态链接 Go 应用

如果使用 CGO，例如依赖 SQLite、某些 C 库、系统 DNS 行为，可能需要动态链接：

```Dockerfile
FROM golang:1.23 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /out/app ./cmd/app


FROM gcr.io/distroless/cc-debian12 AS runtime

COPY --from=build /out/app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]
```

如果运行时报类似：

```text
no such file or directory
```

但二进制明明存在，常见原因是动态链接器或共享库缺失。

检查方式：

```bash
ldd ./app
```

---

## 10. Python 多阶段构建

Python 应用不一定需要编译应用本身，但依赖包经常需要编译。

典型场景：

- 依赖没有当前架构的 wheel。
- 包含 C/C++/Rust 扩展。
- 需要 `gcc`、`build-essential`、`python3-dev`、`libpq-dev` 等。

目标是：

```text
build stage: 安装编译工具，构建依赖
runtime stage: 只保留 venv 和运行库
```

### 10.1 venv 模板

```Dockerfile
FROM python:3.12-slim AS build

WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends build-essential gcc \
    && rm -rf /var/lib/apt/lists/*

RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


FROM python:3.12-slim AS runtime

WORKDIR /app
ENV PATH="/opt/venv/bin:$PATH"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

COPY --from=build /opt/venv /opt/venv
COPY . .

EXPOSE 8000
CMD ["python", "-m", "app"]
```

### 10.2 FastAPI 示例

```Dockerfile
FROM python:3.12-slim AS build

WORKDIR /app
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


FROM python:3.12-slim AS runtime

WORKDIR /app
ENV PATH="/opt/venv/bin:$PATH"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

COPY --from=build /opt/venv /opt/venv
COPY app ./app

EXPOSE 8000
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

注意：

- 如果某个 Python 包运行时还需要系统动态库，runtime stage 也要安装运行库。
- build stage 安装的是 `*-dev`，runtime stage 通常只需要非 `-dev` 运行库。

例如 PostgreSQL：

```text
build stage: libpq-dev
runtime stage: libpq5
```

---

## 11. Rust 多阶段构建

Rust 官方镜像很大，因为包含 `cargo`、`rustc`、LLVM、构建工具等。

多阶段构建几乎是 Rust 生产镜像的标配。

```Dockerfile
FROM rust:1.84 AS build

WORKDIR /src
COPY Cargo.toml Cargo.lock ./
COPY src ./src
RUN cargo build --release


FROM debian:bookworm-slim AS runtime

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /src/target/release/app /usr/local/bin/app
USER 65532:65532
ENTRYPOINT ["app"]
```

如果 Rust 程序静态链接 musl，可以用更小的 runtime：

```Dockerfile
FROM rust:1.84 AS build

WORKDIR /src
RUN rustup target add x86_64-unknown-linux-musl
RUN apt-get update \
    && apt-get install -y --no-install-recommends musl-tools \
    && rm -rf /var/lib/apt/lists/*
COPY Cargo.toml Cargo.lock ./
COPY src ./src
RUN cargo build --release --target x86_64-unknown-linux-musl


FROM scratch AS runtime

COPY --from=build /src/target/x86_64-unknown-linux-musl/release/app /app
ENTRYPOINT ["/app"]
```

和 Go 一样，`scratch` 要注意证书、用户、时区和调试问题。

---

## 12. Java 多阶段构建

Java 构建阶段通常需要 JDK 和 Maven/Gradle，运行阶段通常只需要 JRE。

### 12.1 Maven 模板

```Dockerfile
FROM maven:3.9-eclipse-temurin-21 AS build

WORKDIR /src
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn package -DskipTests


FROM eclipse-temurin:21-jre AS runtime

WORKDIR /app
COPY --from=build /src/target/*.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
```

### 12.2 Gradle 模板

```Dockerfile
FROM gradle:8-jdk21 AS build

WORKDIR /src
COPY build.gradle settings.gradle ./
COPY gradle ./gradle
RUN gradle dependencies --no-daemon
COPY src ./src
RUN gradle build --no-daemon -x test


FROM eclipse-temurin:21-jre AS runtime

WORKDIR /app
COPY --from=build /src/build/libs/*.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
```

Java 补充经验：

- 构建阶段用 JDK。
- 运行阶段用 JRE。
- 可以进一步用 `jlink` 生成裁剪版 runtime，但复杂度更高。
- 容器内 Java 内存要结合 `-XX:MaxRAMPercentage`、Kubernetes limits 一起看。

---

## 13. PHP 多阶段构建

PHP 不一定需要编译应用，但 Composer 的 dev dependencies 不应该进入生产镜像。

```Dockerfile
FROM composer:2 AS vendor

WORKDIR /app
COPY composer.json composer.lock ./
RUN composer install \
    --no-dev \
    --prefer-dist \
    --no-interaction \
    --no-progress \
    --optimize-autoloader


FROM php:8.3-fpm-alpine AS runtime

WORKDIR /var/www/html
COPY --from=vendor /app/vendor ./vendor
COPY . .

CMD ["php-fpm"]
```

如果需要构建前端资源，可以再加一个 Node stage：

```Dockerfile
FROM node:22-slim AS assets

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY resources ./resources
COPY vite.config.* ./
RUN npm run build
```

然后在 PHP runtime stage 中：

```Dockerfile
COPY --from=assets /app/public/build ./public/build
```

---

## 14. 基础镜像怎么选

### 14.1 常见基础镜像对比

| 镜像 | 优点 | 缺点 | 适合场景 |
| :--- | :--- | :--- | :--- |
| `scratch` | 极小，没有包管理器和 shell | 调试困难，缺证书/用户/时区 | 完全静态链接二进制 |
| distroless | 小，攻击面低，有常见运行环境 | 没有 shell，调试不如 Debian 方便 | Go/Java/Node/Python 生产运行 |
| Alpine | 小，有包管理器 | musl libc 兼容性可能踩坑 | 简单服务、静态文件、轻量工具 |
| Debian/Ubuntu slim | 兼容性好，调试方便 | 比 distroless/Alpine 大 | 需要 glibc 或系统依赖的服务 |
| 官方 SDK 镜像 | 构建方便 | 很大，不适合生产运行 | build stage |

### 14.2 `alpine` 不是永远更好

Alpine 使用 musl libc，不是 glibc。

可能遇到的问题：

- 某些 Python wheel 不兼容，需要从源码编译。
- 某些 Node native addon 需要额外编译。
- DNS、locale、动态链接行为和 glibc 系统不同。
- 排查时工具较少。

所以不要简单认为：

```text
alpine 最小，所以一定最好。
```

更稳妥的策略：

- Go 静态二进制：distroless/static 或 scratch。
- 动态链接或依赖 glibc：debian-slim 或 distroless/cc。
- Java：JRE 镜像或 distroless/java。
- Node SSR：node slim 或 distroless Node.js 镜像，例如 `gcr.io/distroless/nodejs22-debian12`。
- Python：python slim 通常比 alpine 更省心。

---

## 15. Dockerfile 缓存优化

### 15.1 先复制依赖描述文件，再复制源码

错误写法：

```Dockerfile
COPY . .
RUN npm ci
```

只要任意源码文件变化，`npm ci` 缓存就失效。

更好写法：

```Dockerfile
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
```

Go：

```Dockerfile
COPY go.mod go.sum ./
RUN go mod download
COPY . .
```

Python：

```Dockerfile
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
```

Java Maven：

```Dockerfile
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
```

核心原则：

```text
越稳定、越耗时的步骤越放前面。
越频繁变化的源码越靠后复制。
```

### 15.2 使用 `.dockerignore`

`.dockerignore` 控制发送给 Docker daemon 的构建上下文。

常见内容：

```gitignore
.git
.github
node_modules
dist
build
.next
.nuxt
.output
coverage
*.log
.env
.DS_Store
```

作用：

- 减少构建上下文大小。
- 避免把本地缓存、日志、密钥、构建产物发送进 Docker build。
- 提高缓存稳定性。

注意：

```text
.dockerignore 不是安全边界，但能减少误打包敏感文件的概率。
```

真正的密钥不应该出现在构建上下文里。

---

## 16. BuildKit 常用能力

建议使用 BuildKit：

```bash
DOCKER_BUILDKIT=1 docker build -t app .
```

现代 Docker 通常已经默认启用 BuildKit。

### 16.1 cache mount

缓存包管理器下载内容，但不把缓存写进最终镜像层。

Node npm：

```Dockerfile
# syntax=docker/dockerfile:1.7
FROM node:22-slim AS deps

WORKDIR /app
COPY package.json package-lock.json ./
RUN --mount=type=cache,target=/root/.npm \
    npm ci
```

Go：

```Dockerfile
# syntax=docker/dockerfile:1.7
FROM golang:1.23 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /out/app ./cmd/app
```

Cargo：

```Dockerfile
# syntax=docker/dockerfile:1.7
RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/src/target \
    cargo build --release
```

### 16.2 secret mount

构建时需要私有仓库 token，不要用 `ARG` 或 `ENV` 写进镜像层。

错误做法：

```Dockerfile
ARG NPM_TOKEN
RUN npm config set //registry.npmjs.org/:_authToken=$NPM_TOKEN
```

更好做法：

```Dockerfile
# syntax=docker/dockerfile:1.7
RUN --mount=type=secret,id=npmrc,target=/root/.npmrc \
    npm ci
```

构建：

```bash
docker build --secret id=npmrc,src=$HOME/.npmrc -t app .
```

重点：

```text
secret mount 只在 RUN 期间可见，不会进入镜像层。
```

### 16.3 bind mount 构建上下文文件

BuildKit 还可以在 `RUN` 中临时挂载文件，减少不必要的层内容。

例如：

```Dockerfile
# syntax=docker/dockerfile:1.7
RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download
```

日常不一定必须用，但在优化大型构建时很有用。

---

## 17. 安全实践

### 17.1 不要用构建镜像做生产镜像

不推荐：

```Dockerfile
FROM golang:1.23
```

作为最终生产镜像。

不推荐：

```Dockerfile
FROM rust:1.84
```

作为最终生产镜像。

推荐：

```text
build stage 用 SDK 镜像
runtime stage 用 runtime/minimal 镜像
```

### 17.2 使用非 root 用户

Debian/Ubuntu slim 示例：

```Dockerfile
RUN useradd -r -u 10001 appuser
USER appuser
```

Alpine：

```Dockerfile
RUN adduser -D -u 10001 appuser
USER appuser
```

distroless 通常可用：

```Dockerfile
USER nonroot:nonroot
```

或者：

```Dockerfile
USER 65532:65532
```

### 17.3 最小化运行时包

APT 安装建议：

```Dockerfile
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*
```

注意：

- `--no-install-recommends` 避免安装推荐包。
- 清理 `/var/lib/apt/lists/*` 减小镜像层。
- 不要在一个层里 `apt-get update`，另一个层里 `apt-get install`，容易缓存出错。

### 17.4 不要把密钥打进镜像

不要在 Dockerfile 里写：

```Dockerfile
ENV AWS_SECRET_ACCESS_KEY=...
COPY .env .env
```

不要把 `.env` 放进构建上下文。

`.dockerignore` 应该包含：

```gitignore
.env
.env.*
id_rsa
*.pem
```

如果构建需要密钥，用 BuildKit secret mount。

### 17.5 镜像扫描

常用工具：

```bash
trivy image app:latest
docker scout cves app:latest
```

注意：

```text
扫描结果不是越少越代表绝对安全，但减少不必要包能明显减少漏洞面。
```

### 17.6 固定基础镜像版本

不推荐生产 Dockerfile 长期使用：

```Dockerfile
FROM node:latest
FROM python:latest
FROM ubuntu:latest
```

更推荐：

```Dockerfile
FROM node:22-slim
FROM python:3.12-slim
FROM eclipse-temurin:21-jre
```

更严格的生产环境还可以固定 digest：

```Dockerfile
FROM node:22-slim@sha256:<digest>
```

取舍：

- 固定 tag：可读性好，升级成本低。
- 固定 digest：可复现性强，但需要定期维护更新。

### 17.7 Multi-Stage 不会自动安全

Multi-Stage Build 只是工具，不是安全保证。

如果你这样写：

```Dockerfile
FROM node:22-slim AS build
WORKDIR /app
COPY . .
RUN npm ci && npm run build

FROM node:22-slim AS runtime
WORKDIR /app
COPY --from=build /app /app
CMD ["node", "server.js"]
```

最终仍然可能把源码、dev dependencies、测试文件、`.env`、构建缓存全部复制进 runtime stage。

正确思路是：

```text
只复制运行必需的最小集合。
```

例如：

```Dockerfile
COPY --from=build /app/dist ./dist
COPY --from=prod-deps /app/node_modules ./node_modules
COPY package.json ./
```

---

## 18. 如何验证镜像是否真的变小

### 18.1 查看镜像大小

```bash
docker images
```

或者：

```bash
docker image ls app
```

### 18.2 查看镜像层

```bash
docker history app:latest
```

重点看：

- 哪一层最大。
- 是否把 build tools 留进最终镜像。
- 是否有大文件先复制后删除的问题。

### 18.3 查看镜像内容

可以临时运行：

```bash
docker run --rm -it app:latest sh
```

但 distroless 和 scratch 没有 shell。

这种情况下可以用：

```bash
docker export "$(docker create app:latest)" | tar -tv | head
```

或者使用专门工具：

```bash
dive app:latest
```

### 18.4 检查最终镜像有没有构建工具

```bash
docker run --rm app:latest which gcc
docker run --rm app:latest which go
docker run --rm app:latest which cargo
docker run --rm app:latest which npm
```

如果最终镜像不应该包含这些工具，但命令能找到，说明 runtime stage 还不够干净。

### 18.5 构建到指定 stage 调试

```bash
docker build --target build -t app-build .
docker run --rm -it app-build sh
```

这很适合检查 build stage 中产物路径是否正确。

---

## 19. 常见坑

### 19.1 `COPY . .` 太早导致缓存失效

如果先复制全部源码，再安装依赖：

```Dockerfile
COPY . .
RUN npm ci
```

任何源码改动都会让依赖安装缓存失效。

应该先复制 lockfile：

```Dockerfile
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
```

### 19.2 `.dockerignore` 忘记排除 `node_modules`

如果本地 `node_modules` 被发送到构建上下文：

- 构建变慢。
- 上下文变大。
- 可能污染容器内依赖。
- 可能导致不同平台的 native addon 出问题。

应该加入：

```gitignore
node_modules
```

### 19.3 构建产物路径写错

常见错误：

```Dockerfile
COPY --from=build /app/dist ./dist
```

但实际产物可能在：

```text
/app/.output
/app/build
/app/.next/standalone
/app/target/release
```

排查：

```bash
docker build --target build -t app-build .
docker run --rm -it app-build sh
ls -lah /app
```

### 19.4 final stage 里缺动态库

症状：

```text
no such file or directory
```

但文件明明存在。

常见原因：

- 动态链接器缺失。
- glibc/musl 不匹配。
- `.so` 运行库缺失。

排查：

```bash
ldd ./binary
```

### 19.5 `scratch` 里 HTTPS 失败

报错可能类似：

```text
x509: certificate signed by unknown authority
```

原因是没有 CA 证书。

解决：

```Dockerfile
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
```

或者用 distroless/static。

### 19.6 删除文件不一定让镜像变小

错误示例：

```Dockerfile
COPY big.tar /tmp/big.tar
RUN rm /tmp/big.tar
```

`big.tar` 已经存在于上一层里，后面删除只是让它在最终文件系统视图中不可见，镜像历史层仍然包含它。

更好做法：

```Dockerfile
RUN curl -L -o /tmp/big.tar https://example.com/big.tar \
    && tar -xf /tmp/big.tar -C /opt/app \
    && rm /tmp/big.tar
```

或者用 Multi-Stage Build，只从 build stage 复制最终产物。

---

## 20. 完整模板速查

### 20.1 Go 静态二进制

```Dockerfile
FROM golang:1.23 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/app

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]
```

### 20.2 Node SSR

```Dockerfile
FROM node:22-slim AS deps
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

FROM node:22-slim AS build
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

FROM node:22-slim
WORKDIR /app
ENV NODE_ENV=production
COPY --from=build /app/.output ./.output
EXPOSE 3000
CMD ["node", ".output/index.mjs"]
```

### 20.3 Python venv

```Dockerfile
FROM python:3.12-slim AS build
WORKDIR /app
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

FROM python:3.12-slim
WORKDIR /app
ENV PATH="/opt/venv/bin:$PATH"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
COPY --from=build /opt/venv /opt/venv
COPY . .
CMD ["python", "-m", "app"]
```

### 20.4 Java Maven

```Dockerfile
FROM maven:3.9-eclipse-temurin-21 AS build
WORKDIR /src
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn package -DskipTests

FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=build /src/target/*.jar app.jar
ENTRYPOINT ["java", "-jar", "app.jar"]
```

---

## 21. 日常写 Dockerfile 的建议顺序

可以按这个顺序思考：

1. 先确认运行时真正需要什么。
2. 再选择 runtime base image。
3. 再选择 build base image。
4. 把依赖安装、编译、测试放在 build stage。
5. 只把最终产物复制到 runtime stage。
6. 使用 `.dockerignore` 减少构建上下文。
7. 调整 Dockerfile 顺序，让依赖层尽量命中缓存。
8. 使用非 root 用户运行。
9. 用 `docker history`、`docker image ls`、`trivy` 验证结果。

---

## 22. 总结

Multi-Stage Build 解决的核心问题不是“语法怎么写”，而是工程边界：

```text
构建环境应该完整，运行环境应该最小。
```

它带来的收益：

- 生产镜像更小。
- 构建工具不进入线上环境。
- CVE 数量通常更少。
- 攻击面更小。
- Dockerfile 仍然保持一个文件，构建顺序由 Docker 管理。
- 不需要额外 `docker cp` 或手写构建脚本。

最终可以记成一句话：

> Multi-Stage Build 就是在一个 Dockerfile 里创建多个阶段：前面的阶段负责构建，最后的阶段负责运行，最终镜像只保留显式复制进来的运行产物。
