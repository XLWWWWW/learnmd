# 使用 OverlayFS 联合挂载容器镜像层

> 专题导航：[容器镜像文件系统专题索引](./容器镜像文件系统专题索引.md)

本文整理自 Ivan Velichko 的 iximiuz Labs Challenge《Union Mount Container Image Layers into a Flat Filesystem Using OverlayFS》，并补充 OCI layer 顺序、OverlayFS Copy-on-Write、whiteout、容器运行时边界和常见排错方法。

## 1. 这篇文章想说明什么

这篇 challenge 要完成的事情是：

```text
远程容器镜像
  -> 保存为 OCI Image Layout
  -> 找到并解压各个 filesystem layer
  -> 使用 OverlayFS 联合挂载
  -> 得到类似容器 rootfs 的统一文件系统视图
```

它最重要的目的不是记住一条很长的 `mount` 命令，而是建立下面这个心智模型：

> 镜像不是一个已经展开的根目录，而是一组按顺序排列的文件系统变化集；容器运行时需要把这些只读层叠加起来，再加一个容器专属的可写层，才能得到容器看到的 rootfs。

可以把整个过程表示成：

```text
              merged：容器看到的统一视图
                         |
              +----------+----------+
              |                     |
       upperdir：可写层          lowerdir：镜像只读层
                                      |
                              layer 2：最高优先级
                              layer 1
                              layer 0：基础层
```

这正是 Docker、Podman 和 containerd 的 snapshotter 或 graph driver 所做工作的简化版本。准备好 rootfs 后，容器运行时还会挂载 `/proc`、`/sys`、`/dev` 等运行时文件系统，最后才由 `runc` 一类 OCI runtime 创建 namespaces、cgroups 并启动进程。

相关背景可继续阅读：

- [容器镜像内部结构详解](./容器镜像内部结构详解.md)
- [OCI Image Layout 与容器镜像结构](./OCI_Image_Layout与镜像结构.md)
- [提取容器镜像的根文件系统](./提取容器镜像根文件系统.md)

---

## 2. 实验镜像的三个 layer

文章使用的 `registry.iximiuz.com/mini-web:v1` 有三个 layer：

| Layer | 内容 | 语义 |
| :--- | :--- | :--- |
| Layer 0 | `etc/os-release`、`bin/`、`home/` 等 | 最小 Linux 风格的基础目录 |
| Layer 1 | `app/server.py`、`app/config.json`、`app/version.txt` | 添加应用文件 |
| Layer 2 | 新版本的 `app/config.json` 和 `app/version.txt` | 覆盖下层同名文件 |

镜像 layer 不是三个完整 rootfs，而是三个增量变化集：

```text
Layer 0              Layer 1                 Layer 2
etc/os-release       app/server.py           app/config.json（新版本）
bin/                 app/config.json         app/version.txt（新版本）
home/                app/version.txt
...
```

联合挂载后，用户看到的是：

```text
etc/os-release       <- Layer 0
app/server.py        <- Layer 1
app/config.json      <- Layer 2 覆盖 Layer 1
app/version.txt      <- Layer 2 覆盖 Layer 1
```

**OverlayFS 不会把这些目录预先复制成第四份完整文件系统。它在查找文件时按照优先级选择最上层可见的版本，因此创建统一视图通常比物理复制所有文件更快，也更节省空间。**

---

## 3. 保存并定位 OCI layer

### 3.1 保存为 OCI Image Layout

challenge 使用 `crane pull` 将镜像保存为 OCI 格式。不同版本的 `crane` 参数细节可能略有差异，应先查看：

```bash
crane pull --help
```

典型用法是：

```bash
mkdir -p /tmp/mini-web

crane pull \
  registry.iximiuz.com/mini-web:v1 \
  /tmp/mini-web \
  --format oci
```

保存后应能看到类似结构：

```text
/tmp/mini-web/
├── oci-layout
├── index.json
└── blobs/
    └── sha256/
        ├── <manifest-or-index>
        ├── <config>
        └── <filesystem-layers>
```

### 3.2 从 manifest 取得 layer 顺序

对于顶层 descriptor 直接指向单平台 manifest 的 OCI layout：

```bash
layout_dir=/tmp/mini-web

manifest_digest=$(jq -r '.manifests[0].digest' "$layout_dir/index.json")
manifest_file="$layout_dir/blobs/sha256/${manifest_digest#sha256:}"

jq -r '.layers[].digest' "$manifest_file"
```

manifest 中 `.layers[]` 的顺序非常重要：

```text
数组第一个元素：最老的基础层
数组最后一个元素：最新的顶层
```

不能使用 blob 文件名、文件大小或修改时间猜测 layer 顺序。多平台镜像还需要先从 Image Index 中选出目标平台的 manifest，再读取它的 `.layers[]`。

### 3.3 分别解压每个 layer

先创建目录：

```bash
sudo mkdir -p /opt/cont-1/layers/{0,1,2}
```

然后按照 manifest 顺序，将三个 blob 分别解压到对应编号目录：

```bash
sudo tar -xf /tmp/mini-web/blobs/sha256/<layer-0-digest> \
  -C /opt/cont-1/layers/0

sudo tar -xf /tmp/mini-web/blobs/sha256/<layer-1-digest> \
  -C /opt/cont-1/layers/1

sudo tar -xf /tmp/mini-web/blobs/sha256/<layer-2-digest> \
  -C /opt/cont-1/layers/2
```

layer blob 可能使用 gzip 或 zstd 压缩。GNU tar 通常能自动识别常见压缩格式，但仍可先检查：

```bash
file /tmp/mini-web/blobs/sha256/<layer-digest>
```

需要手动处理时：

```bash
gzip -dc <layer-blob> | sudo tar -xf - -C <target-directory>
zstd -dc <layer-blob> | sudo tar -xf - -C <target-directory>
```

生产场景还要关注 UID/GID、权限、硬链接、设备文件和扩展属性。普通用户直接解压，可能无法完整恢复 layer 元数据。

---

## 4. OverlayFS 的四个目录

文章让我们准备四类目录：

| 目录 | 作用 | 对应容器概念 |
| :--- | :--- | :--- |
| `lowerdir` | 一个或多个只读层 | 镜像 filesystem layers |
| `upperdir` | 接收新增、修改和删除 | 容器专属可写层 |
| `workdir` | OverlayFS 内部临时工作区 | 实现 copy-up 等操作 |
| `merged` | 对用户呈现统一视图 | 容器的 rootfs |

其中 `merged` 不是新的物理副本。它只是一个挂载点；卸载后统一视图会消失，但 `lowerdir` 和 `upperdir` 中的真实文件仍然存在。

准备目录：

```bash
sudo mkdir -p /opt/cont-1/overlayfs/{upper,work}
sudo mkdir -p /opt/cont-1/rootfs
```

挂载：

```bash
sudo mount -t overlay overlay \
  -o lowerdir=/opt/cont-1/layers/2:/opt/cont-1/layers/1:/opt/cont-1/layers/0,upperdir=/opt/cont-1/overlayfs/upper,workdir=/opt/cont-1/overlayfs/work \
  /opt/cont-1/rootfs
```

### 4.1 为什么 lowerdir 顺序是反的

OCI manifest 的 layer 顺序是：

```text
base -> middle -> top
0    -> 1      -> 2
```

OverlayFS 的 `lowerdir` 参数则要求：

```text
top:middle:base
2:1:0
```

左边优先级最高。如果多个 layer 中存在相同路径，OverlayFS 使用最靠左的版本。

因此：

```text
/opt/cont-1/layers/1/app/version.txt = 旧版本
/opt/cont-1/layers/2/app/version.txt = 新版本

/opt/cont-1/rootfs/app/version.txt   = 新版本
```

顺序写反时，旧文件反而会遮住新文件，得到与真实镜像不一致的 rootfs。

---

## 5. Copy-on-Write 如何工作

刚挂载完成时：

```bash
sudo ls -la /opt/cont-1/overlayfs/upper
```

`upperdir` 应该是空的，因为还没有对统一视图进行修改。

### 5.1 创建新文件

```bash
echo 'iximiuz Labs Rules!' | \
  sudo tee /opt/cont-1/rootfs/hello.txt >/dev/null
```

此时：

```text
rootfs/hello.txt           可见
overlayfs/upper/hello.txt  实际保存位置
layers/0、1、2             都没有变化
```

检查：

```bash
sudo find /opt/cont-1/overlayfs/upper -maxdepth 2 -ls
```

### 5.2 修改下层已有文件

如果修改只存在于 lower layer 的文件，OverlayFS 会先执行 copy-up：

```text
lowerdir 中的原文件
       |
       | copy-up
       v
upperdir 中的完整副本 -> 在副本上修改
```

例如：

```bash
echo 'locally modified' | \
  sudo tee /opt/cont-1/rootfs/app/version.txt >/dev/null
```

修改后，lower layers 仍保存镜像原始版本，新的文件副本位于 `upperdir/app/version.txt`，并在 merged 视图中遮住所有 lower 版本。

需要注意，OverlayFS 通常按文件执行 copy-up，而不是只复制被修改的数据块。修改一个很大的下层文件，可能产生接近整个文件大小的额外写入和磁盘占用。

### 5.3 删除文件

下层文件不能真的从只读 layer 中删除。删除 merged 视图中的下层文件时，OverlayFS 会在 upperdir 中记录 whiteout，用它遮住下层文件：

```text
lowerdir/file       仍然存在
upperdir/whiteout   表示“在 merged 中隐藏 file”
merged/file         不再可见
```

这解释了为什么容器中删除文件不会改变镜像层，也解释了镜像 layer 为什么需要特殊的删除标记。

---

## 6. 文章刻意省略的 whiteout 问题

该 challenge 明确说明，实验镜像只包含新增和修改，没有删除操作。这是一个重要的简化条件。

OCI layer tar 使用下面的普通文件名表达删除：

```text
.wh.filename       删除下层的 filename
.wh..wh..opq       让当前目录变成 opaque，隐藏下层同目录内容
```

OverlayFS 在磁盘上的 whiteout 表示通常是特殊设备节点或扩展属性，opaque directory 也依赖 OverlayFS 扩展属性。两者的表示方法并不相同。

因此，直接执行：

```bash
tar -xf layer.tar -C layer-directory
```

再把目录放进 `lowerdir`，不一定能正确处理真实 OCI 镜像中的删除语义。可能出现：

- `.wh.*` 文件直接暴露在 merged 视图中。
- 本应被删除的下层文件重新出现。
- opaque directory 下本应隐藏的旧文件仍然可见。

可以先检查实验 layer：

```bash
find /opt/cont-1/layers -name '.wh.*' -print
```

如果有输出，就不能把普通解压目录直接视为已经适配 OverlayFS 的 snapshot。真实容器运行时的 unpack/snapshotter 过程会解析 OCI whiteout，并转换成对应存储驱动能够理解的表示。

这也是为什么手动处理真实镜像时，通常应该使用 containerd snapshotter、Docker graph driver、`umoci` 等成熟工具，而不是简单解压 tar 后直接联合挂载。

---

## 7. OverlayFS 挂载条件与排错

挂载前应满足：

- 运行环境是 Linux，内核支持 OverlayFS。
- 执行者具有 `CAP_SYS_ADMIN`，通常意味着需要 root 或 `sudo`。
- `upperdir` 和 `workdir` 位于同一个文件系统。
- `workdir` 是 OverlayFS 专用的空目录。
- `upperdir`、`workdir`、`merged` 不要互相嵌套。
- 底层文件系统支持 OverlayFS 所需的目录项类型和扩展属性。
- `lowerdir` 使用正确的绝对路径和正确的层顺序。

查看内核是否支持：

```bash
grep overlay /proc/filesystems
```

查看挂载：

```bash
findmnt /opt/cont-1/rootfs
mount | grep '/opt/cont-1/rootfs'
```

常见的 `wrong fs type`、`invalid argument` 或 `operation not permitted` 可能来自：

- lowerdir 路径不存在或顺序拼写错误。
- `workdir` 与 `upperdir` 不在同一文件系统。
- `workdir` 正在被其他 OverlayFS 挂载使用。
- 当前环境是非特权容器，缺少挂载权限。
- lower 或 upper 位于 OverlayFS 不支持的 NFS、某些 FUSE 文件系统等后端。
- upper 文件系统不支持必要的 xattr 或有效 `d_type`。

rootless 容器环境常使用 `fuse-overlayfs`，因为普通用户通常不能直接执行内核 OverlayFS 挂载。

---

## 8. 联合挂载不等于物理扁平化

文章标题中的 flat filesystem 指“看起来像一棵普通目录树的统一视图”，并不表示所有文件已经被复制到一个新目录。

```text
挂载前：三个 lower 目录 + 一个 upper 目录
挂载后：merged 提供统一视图
卸载后：merged 重新变成普通空目录
```

卸载：

```bash
sudo umount /opt/cont-1/rootfs
```

如果确实需要生成可保存、可传输的扁平 tar，可以在挂载期间从 merged 视图重新归档：

```bash
sudo tar \
  --numeric-owner \
  --xattrs \
  --acls \
  -C /opt/cont-1/rootfs \
  -cpf /tmp/mini-web-rootfs.tar .
```

但这种扁平归档会丢失原始 layer 边界和镜像 config、manifest 等元数据。它适合查看或传输 rootfs，不再是一份完整的 OCI 镜像。

---

## 9. 它与真正容器 rootfs 的差别

完成 OverlayFS 挂载后，我们得到的是：

```text
镜像只读层 + 一个空的容器可写层
```

此时：

- Dockerfile 中的 `RUN` 结果已经存在于镜像 layers 中，因为 `RUN` 在构建时执行。
- 镜像的 `ENTRYPOINT` 和 `CMD` 尚未执行。
- 没有应用进程产生的日志、缓存和数据库文件。
- `/proc`、`/sys`、`/dev` 通常还没有按容器方式挂载。
- `/etc/hosts`、`/etc/hostname`、`/etc/resolv.conf` 等运行时内容可能尚未注入。
- OCI config 中的环境变量、工作目录、用户、capabilities 等也不是仅靠 OverlayFS 挂载就能应用的。

更完整的容器启动链路是：

```text
Registry 中的 manifest、config、layers
  -> 解压 layers，创建 snapshots
  -> OverlayFS 联合挂载，生成 rootfs
  -> 生成 OCI runtime bundle：config.json + rootfs
  -> 添加 bind mount、proc、sysfs、设备和网络配置
  -> runc 创建隔离环境并执行 Entrypoint/Cmd
```

所以本实验完成的是容器启动链路中非常关键的一步，但不是完整实现一个容器运行时。

---

## 10. 总结

这篇 challenge 主要说明了六件事：

1. 容器镜像 layer 是有顺序的增量文件系统变化集，不是多个完整 rootfs。
2. OverlayFS 可以把多个只读 layer 组合成统一文件系统视图。
3. `lowerdir` 从左到右按照“顶层到基础层”排列，与 manifest 顺序相反。
4. `upperdir` 就是容器的可写层；新增和 copy-up 后的修改都保存在这里。
5. merged 只是挂载视图，不是已经物理合并好的新目录。
6. 真实镜像还要正确处理 OCI whiteout、权限、扩展属性和运行时挂载，文章为了突出核心机制省略了这些复杂性。

最简洁的心智模型是：

```text
Image layers = lowerdir
Container writable layer = upperdir
Container rootfs view = merged
OverlayFS internal scratch space = workdir
```
