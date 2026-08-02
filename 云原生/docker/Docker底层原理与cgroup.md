# Docker 底层原理与 cgroup

## 目录

- 第一章：先建立整体认识
- 第二章：Docker 依赖的 Linux 底层能力
- 第三章：Linux Namespace
- 第四章：lsns 与 nsenter
- 第五章：cgroup 基础
- 第六章：cgroup v2 详细整理
- 第七章：用 cgroupfs 手动限制进程资源
- 第八章：用 libcgroup 和 systemd 管理 cgroup
- 第九章：Docker 与 cgroup 的关系
- 第十章：Docker 文件系统原理
- 第十一章：overlay2 详细整理
- 第十二章：把容器运行过程串起来
- 第十三章：排查命令清单
- 第十四章：实验清单
- 第十五章：总结

---

## 第一章：先建立整体认识

Docker 不是一种轻量级虚拟机。更准确地说：

> Docker 是把一个普通 Linux 进程，放进一组隔离环境里运行，并用资源控制、分层文件系统、网络虚拟化等能力，让这个进程看起来像运行在一台独立的小机器中。

容器本质上仍然是宿主机上的进程。它没有自己的内核，也不会像虚拟机那样启动一个完整 Guest OS。容器中的进程直接使用宿主机 Linux Kernel。

Docker 主要依赖下面几类 Linux 能力：

| 能力 | 作用 | 对容器的意义 |
| :--- | :--- | :--- |
| Namespace | 隔离视图 | 让容器看到独立的 PID、网络、挂载点、主机名、用户等 |
| cgroup | 限制与统计资源 | 限制 CPU、内存、IO、进程数量等 |
| UnionFS / overlay2 | 分层文件系统 | 镜像多层复用，容器只写自己的可写层 |
| capabilities | 拆分 root 权限 | 容器里的 root 不一定拥有宿主机完整 root 权限 |
| seccomp | 限制系统调用 | 降低容器进程能调用的内核接口范围 |
| AppArmor / SELinux | 强制访问控制 | 进一步限制进程能访问哪些资源 |
| veth / bridge / iptables | 容器网络 | 让隔离的 network namespace 能接入宿主机和外部网络 |

一句话串起来：

> Namespace 负责“看见什么”，cgroup 负责“能用多少”，overlay2 负责“文件从哪里来、写到哪里去”，capabilities/seccomp/LSM 负责“能做什么危险动作”。

---

## 第二章：Docker 依赖的 Linux 底层能力

### 2.1 进程是容器的核心

Linux 里一切容器最终都要落到进程上。

启动一个容器：

```bash
docker run -d --name web nginx
```

宿主机上会出现 nginx 相关进程：

```bash
ps -ef | grep nginx
```

进入容器：

```bash
docker exec -it web sh
ps -ef
```

你会发现容器内看到的进程列表和宿主机不同。这不是因为进程真的跑在另一台机器上，而是因为它处于不同的 PID namespace。

### 2.2 系统调用是用户态和内核态的边界

容器进程和普通进程一样，需要通过系统调用访问内核能力，例如：

- `fork` / `clone`：创建进程或线程。
- `mount`：挂载文件系统。
- `setns`：加入某个 namespace。
- `unshare`：创建新的 namespace。
- `chroot` / `pivot_root`：切换根文件系统视图。
- `open` / `read` / `write`：文件读写。
- `socket` / `bind` / `connect`：网络通信。

容器隔离并不是靠 Docker daemon 自己“模拟”出来的，而是 Docker、containerd、runc 最终调用 Linux 内核接口完成的。

### 2.3 Docker、containerd、runc 的分工

现代 Docker 大致分为几层：

| 组件 | 作用 |
| :--- | :--- |
| Docker CLI | 用户命令行，例如 `docker run` |
| dockerd | Docker daemon，负责镜像、网络、卷、容器生命周期管理 |
| containerd | 更底层的容器运行时管理器，负责拉镜像、创建容器、管理任务 |
| runc | OCI runtime，真正调用 Linux 内核能力创建容器进程 |

大致流程：

```text
docker run
  -> dockerd
    -> containerd
      -> containerd-shim
        -> runc
          -> Linux Kernel: namespace + cgroup + mount + capabilities + seccomp
```

`runc` 是 OCI runtime 的典型实现。OCI 主要定义两个规范：

- Runtime Spec：容器怎么创建和运行。
- Image Spec：镜像格式怎么组织。

### 2.4 rootfs

容器里的 `/` 目录不是宿主机真正的 `/`，而是容器自己的 root filesystem，也就是 rootfs。

镜像提供只读 rootfs 的多层内容，容器启动时再叠加一个可写层，最后挂成容器进程看到的 `/`。

查看容器内根目录：

```bash
docker run --rm -it alpine sh
ls /
```

这套目录来自 alpine 镜像，不是宿主机原始根目录。

### 2.5 pivot_root 与 chroot

`chroot` 可以改变进程看到的根目录，但它隔离不彻底，历史上更多用于构建、救援、简单沙箱。

容器运行时更常见的是使用 `pivot_root`：

- 把新 rootfs 切换为 `/`。
- 把旧 root 移到新 rootfs 的某个临时目录。
- 卸载旧 root，避免容器进程还能访问宿主机原根目录。

容器文件系统隔离通常不是只靠一个 API，而是配合 mount namespace、rootfs、bind mount、只读挂载、masked path、readonly path 等共同完成。

---

## 第三章：Linux Namespace

### 3.1 namespace 解决什么问题

Namespace 是 Linux 内核提供的隔离机制。

它隔离的是“进程看到的系统资源视图”。不同 namespace 中的进程，看到的资源可以不同。

常见 namespace：

| Namespace | 隔离内容 | Docker 中的表现 |
| :--- | :--- | :--- |
| PID | 进程编号空间 | 容器内主进程看到自己是 PID 1 |
| NET | 网络设备、IP、路由、端口、防火墙规则视图 | 容器有自己的 eth0、IP、路由表 |
| MNT | 挂载点 | 容器有自己的 `/` 和挂载结构 |
| UTS | hostname、domainname | 容器内 hostname 独立 |
| IPC | System V IPC、POSIX message queue | 进程间通信资源隔离 |
| USER | UID、GID 映射 | 容器内 root 可映射成宿主机普通用户 |
| CGROUP | cgroup 根目录视图 | 容器内看到自己的 cgroup 层级 |
| TIME | monotonic/boottime 时钟偏移 | 较新内核支持，容器可有不同时间偏移 |

查看当前 shell 所属 namespace：

```bash
ls -l /proc/$$/ns
```

示例：

```text
cgroup -> cgroup:[4026531835]
ipc    -> ipc:[4026531839]
mnt    -> mnt:[4026531841]
net    -> net:[4026531992]
pid    -> pid:[4026531836]
user   -> user:[4026531837]
uts    -> uts:[4026531838]
```

中括号里的数字可以理解为 namespace 的 inode。两个进程对应 namespace inode 相同，说明它们在同一个 namespace 中。

### 3.2 PID namespace

PID namespace 隔离进程编号。

容器内：

```bash
docker run --rm -it alpine sh
ps
```

你通常会看到 shell 是 PID 1。

宿主机上：

```bash
docker inspect -f '{{.State.Pid}}' <container>
ps -ef | grep <pid>
```

同一个进程，在容器内和宿主机上有不同 PID：

- 容器内 PID：相对容器 PID namespace。
- 宿主机 PID：相对宿主机 PID namespace。

注意：容器内 PID 1 很特殊。Linux 中 PID 1 负责处理孤儿进程和信号。如果容器主进程不正确处理 SIGTERM、SIGCHLD，可能出现容器无法优雅退出或僵尸进程问题。

常见处理方式：

```bash
docker run --init ...
```

`--init` 会注入一个轻量 init 进程，帮助转发信号和回收子进程。

### 3.3 NET namespace

NET namespace 隔离：

- 网卡。
- IP 地址。
- 路由表。
- ARP 表。
- 端口监听空间。
- iptables/nftables 规则视图。

容器默认 bridge 网络通常是：

```text
container eth0
  <-> veth pair
  <-> host vethxxx
  <-> docker0 bridge
  <-> host network
```

容器内查看：

```bash
ip addr
ip route
ss -lntp
```

宿主机查看：

```bash
ip link
bridge link
docker network inspect bridge
```

Docker 常见网络模式和 NET namespace 关系：

| 模式 | 说明 |
| :--- | :--- |
| bridge | 默认模式，容器有独立 NET namespace，通过 docker0 出宿主机 |
| host | 容器共享宿主机 NET namespace，没有独立 IP |
| none | 有独立 NET namespace，但基本没有网络设备，通常只有 lo |
| container:<name> | 和另一个容器共享 NET namespace |

### 3.4 MNT namespace

MNT namespace 隔离挂载点视图。

容器内看到自己的 `/`、`/proc`、`/sys`、`/dev` 等挂载结构。

查看挂载：

```bash
findmnt
mount
cat /proc/self/mountinfo
```

Docker 通过 mount namespace 配合 overlay2，把镜像层和容器可写层挂成容器根文件系统。

### 3.5 UTS namespace

UTS namespace 隔离主机名。

```bash
docker run --rm -it --hostname mybox alpine sh
hostname
```

容器里看到的是 `mybox`，不会改变宿主机 hostname。

### 3.6 IPC namespace

IPC namespace 隔离传统 Linux IPC 资源：

- System V shared memory。
- System V semaphore。
- System V message queue。
- POSIX message queue。

Docker 默认给容器独立 IPC namespace。也可以共享：

```bash
docker run --ipc=host ...
docker run --ipc=container:<name> ...
```

一些需要共享内存的场景会关注它，例如数据库、多进程应用、机器学习任务。

### 3.7 USER namespace

USER namespace 可以让容器内 UID/GID 和宿主机 UID/GID 做映射。

例如容器内看起来是 root：

```text
container uid 0
```

但映射到宿主机可能是普通用户：

```text
host uid 100000
```

这可以降低容器逃逸后的破坏力。

相关命令和文件：

```bash
cat /proc/<pid>/uid_map
cat /proc/<pid>/gid_map
```

Docker rootless 模式和 userns-remap 都会涉及 USER namespace。

### 3.8 CGROUP namespace

CGROUP namespace 隔离进程看到的 cgroup 路径。

没有 cgroup namespace 时，容器内读取 `/proc/self/cgroup` 可能看到宿主机上的完整 cgroup 路径。

有 cgroup namespace 后，容器内看到的是相对它自己的 cgroup 根，减少宿主机信息暴露。

Docker 参数：

```bash
docker run --cgroupns=private ...
docker run --cgroupns=host ...
```

### 3.9 namespace 的创建和加入

相关系统调用：

| 系统调用 | 作用 |
| :--- | :--- |
| `clone` | 创建进程时顺便创建 namespace |
| `unshare` | 让当前进程脱离某些 namespace，创建新的 namespace |
| `setns` | 加入已有 namespace |

常用工具：

```bash
unshare
nsenter
lsns
```

示例：创建新的 UTS namespace：

```bash
sudo unshare --uts /bin/bash
hostname test-ns
hostname
```

另一个终端查看宿主机 hostname，不会被改变。

---

## 第四章：lsns 与 nsenter

### 4.1 lsns 是什么

`lsns` 用来列出系统上的 namespace。

常用命令：

```bash
lsns
lsns -t net
lsns -t pid
lsns -p <pid>
```

常见输出字段：

| 字段 | 含义 |
| :--- | :--- |
| NS | namespace inode |
| TYPE | namespace 类型 |
| NPROCS | namespace 中进程数量 |
| PID | 该 namespace 中某个进程在宿主机上的 PID |
| USER | 进程所属用户 |
| COMMAND | 进程命令 |

查看某个容器的 namespace：

```bash
CID=$(docker run -d nginx)
PID=$(docker inspect -f '{{.State.Pid}}' $CID)
sudo lsns -p $PID
```

也可以直接看 `/proc`：

```bash
sudo ls -l /proc/$PID/ns
```

### 4.2 nsenter 是什么

`nsenter` 用来进入另一个进程所在的 namespace。

它的关键思想是：

> 不是进入容器，而是让一个新进程加入目标进程的 namespace。

常用命令：

```bash
sudo nsenter -t <pid> -n ip addr
sudo nsenter -t <pid> -m sh
sudo nsenter -t <pid> -u hostname
sudo nsenter -t <pid> -p ps -ef
sudo nsenter -t <pid> -a sh
```

参数含义：

| 参数 | 含义 |
| :--- | :--- |
| `-t <pid>` | 指定目标进程 |
| `-m` | 进入 mount namespace |
| `-u` | 进入 UTS namespace |
| `-i` | 进入 IPC namespace |
| `-n` | 进入 network namespace |
| `-p` | 进入 PID namespace |
| `-U` | 进入 user namespace |
| `-C` | 进入 cgroup namespace |
| `-a` | 进入目标进程可进入的所有 namespace |

进入容器网络 namespace 排查网络：

```bash
CID=$(docker run -d --name web nginx)
PID=$(docker inspect -f '{{.State.Pid}}' web)

sudo nsenter -t $PID -n ip addr
sudo nsenter -t $PID -n ip route
sudo nsenter -t $PID -n ss -lntp
```

进入容器 mount namespace：

```bash
sudo nsenter -t $PID -m findmnt
```

进入容器 PID namespace：

```bash
sudo nsenter -t $PID -p ps -ef
```

更完整地进入：

```bash
sudo nsenter -t $PID -a sh
```

注意：`nsenter` 和 `docker exec` 不完全一样。

| 对比 | docker exec | nsenter |
| :--- | :--- | :--- |
| 依赖 Docker daemon | 依赖 | 不依赖 |
| 入口 | Docker API | Linux namespace |
| 环境变量/工作目录 | Docker 会设置一部分容器环境 | 更偏底层，需要自己处理 |
| 排障价值 | 常规进入容器 | Docker daemon 异常时仍可用 |

### 4.3 用 nsenter 理解容器网络

启动容器：

```bash
docker run -d --name web nginx
PID=$(docker inspect -f '{{.State.Pid}}' web)
```

查看容器网络：

```bash
sudo nsenter -t $PID -n ip addr
```

在宿主机查看 veth：

```bash
ip link
bridge link
```

容器里的 `eth0@ifX` 和宿主机的 `vethxxx@ifY` 是一对 veth pair。数据从容器 `eth0` 发出，会到宿主机 veth，再进入 docker0 网桥。

---

## 第五章：cgroup 基础

### 5.1 cgroup 解决什么问题

cgroup 全称 Control Groups，是 Linux 内核提供的资源管理机制。

它可以把进程组织成层级树，并对这棵树上的进程做资源限制、统计和控制。

典型用途：

- 限制某个进程最多使用多少 CPU。
- 限制某个进程最多使用多少内存。
- 限制磁盘 IO。
- 限制最多能创建多少进程。
- 统计某组进程消耗了多少资源。
- 在多租户环境中做资源隔离。
- Docker、Kubernetes、systemd 都依赖它管理资源。

一句话：

> namespace 让容器“以为自己很独立”，cgroup 让容器“不能无限吃资源”。

### 5.2 cgroupfs

cgroup 的内核接口是一个伪文件系统：`cgroupfs`。

现代 Linux 上通常挂载在：

```bash
/sys/fs/cgroup
```

查看：

```bash
mount -l | grep cgroup
findmnt /sys/fs/cgroup
```

如果是 cgroup v2，通常能看到：

```text
cgroup2 on /sys/fs/cgroup type cgroup2
```

cgroup 没有专门的一组“创建 cgroup 系统调用”。用户态主要通过文件系统操作来管理：

- 创建目录：创建 cgroup。
- 写配置文件：设置限制。
- 写 `cgroup.procs`：把进程移动到 cgroup。
- 删除空目录：删除 cgroup。

示例：

```bash
sudo mkdir /sys/fs/cgroup/demo
ls /sys/fs/cgroup/demo
```

### 5.3 controller

cgroup 通过 controller 控制不同资源。

常见 controller：

| controller | 作用 |
| :--- | :--- |
| cpu | CPU 时间控制 |
| cpuset | 绑定 CPU 核和 NUMA 节点 |
| memory | 内存控制 |
| io | 块设备 IO 控制 |
| pids | 进程数量控制 |
| hugetlb | 大页内存控制 |
| rdma | RDMA 资源控制 |
| misc | 其他杂项资源 |

查看当前 cgroup 支持哪些 controller：

```bash
cat /sys/fs/cgroup/cgroup.controllers
```

### 5.4 cgroup v1 和 v2

| 对比 | cgroup v1 | cgroup v2 |
| :--- | :--- | :--- |
| 层级 | 每个 controller 可以有不同层级 | 统一层级树 |
| 接口一致性 | 不同 controller 文件命名差异大 | 接口更统一 |
| 管理复杂度 | 较高 | 较低 |
| systemd 集成 | 早期支持 | 现代发行版主流方向 |
| Docker 支持 | 支持 | 支持 |

cgroup v2 的重要特点：

- 单一统一树。
- controller 需要在父 cgroup 的 `cgroup.subtree_control` 中启用，子 cgroup 才能使用。
- 资源限制文件命名更统一，例如 `memory.max`、`cpu.max`、`pids.max`。
- 更适合 systemd、容器和 Kubernetes 的统一管理。

查看版本：

```bash
stat -fc %T /sys/fs/cgroup
```

输出：

```text
cgroup2fs
```

说明是 cgroup v2。

### 5.5 进程和 cgroup 的关系

一个进程在同一个 cgroup 层级中只能属于一个 cgroup。

查看当前进程所属 cgroup：

```bash
cat /proc/$$/cgroup
```

查看某个进程：

```bash
cat /proc/<pid>/cgroup
```

把进程移动到某个 cgroup：

```bash
echo <pid> | sudo tee /sys/fs/cgroup/<group>/cgroup.procs
```

子进程默认继承父进程所在 cgroup。所以如果先把 shell 移入某个 cgroup，再从这个 shell 启动程序，新程序一开始就会受到限制。

---

## 第六章：cgroup v2 详细整理

### 6.1 cgroup v2 常见文件

在一个 cgroup v2 目录中，常见文件如下：

| 文件 | 作用 |
| :--- | :--- |
| `cgroup.controllers` | 当前 cgroup 可用 controller |
| `cgroup.subtree_control` | 给子 cgroup 启用哪些 controller |
| `cgroup.procs` | 属于该 cgroup 的进程 PID |
| `cgroup.threads` | 属于该 cgroup 的线程 TID |
| `cgroup.events` | cgroup 状态事件 |
| `cgroup.freeze` | 冻结或恢复 cgroup 中的进程 |
| `cpu.max` | CPU quota 和 period |
| `cpu.weight` | CPU 权重 |
| `cpu.stat` | CPU 使用统计 |
| `memory.current` | 当前内存使用 |
| `memory.max` | 最大内存限制 |
| `memory.high` | 内存软限制压力阈值 |
| `memory.low` | 内存保护 |
| `memory.min` | 更强的内存保护 |
| `memory.swap.max` | swap 使用限制 |
| `memory.events` | 内存相关事件计数 |
| `io.max` | IO 限制 |
| `io.stat` | IO 统计 |
| `pids.current` | 当前进程数 |
| `pids.max` | 最大进程数 |

### 6.2 启用子树 controller

在 cgroup v2 中，父 cgroup 必须显式把 controller 开给子 cgroup。

查看可用 controller：

```bash
cat /sys/fs/cgroup/cgroup.controllers
```

启用：

```bash
echo "+cpu +memory +pids +io" | sudo tee /sys/fs/cgroup/cgroup.subtree_control
```

注意：

- 不是所有环境都允许你直接改根 cgroup。
- 在 systemd 管理的系统上，直接手动操作根 cgroup 可能和 systemd 冲突。
- 实验环境可以手动玩，生产环境建议通过 systemd 或容器运行时管理。

### 6.3 CPU 限制：cpu.max

`cpu.max` 格式：

```text
<quota> <period>
```

单位是微秒。

示例：

```bash
echo "50000 100000" | sudo tee /sys/fs/cgroup/demo/cpu.max
```

含义：

- 每 100000 微秒是一个周期，也就是 100ms。
- 在每个周期中，最多使用 50000 微秒 CPU 时间。
- 等价于 0.5 个 CPU 核。

常见换算：

| 目标 | cpu.max |
| :--- | :--- |
| 不限制 | `max 100000` |
| 0.5 核 | `50000 100000` |
| 1 核 | `100000 100000` |
| 2 核 | `200000 100000` |

多核机器上 CPU 可以超过 100%。比如 4 核机器上，一个进程跑满 4 核通常可以理解为 400% CPU。

### 6.4 CPU 权重：cpu.weight

`cpu.max` 是硬限制，`cpu.weight` 是竞争时的相对权重。

```bash
cat /sys/fs/cgroup/demo/cpu.weight
echo 200 | sudo tee /sys/fs/cgroup/demo/cpu.weight
```

cgroup v2 中 `cpu.weight` 范围通常是 `1-10000`，默认 `100`。

举例：

- A cgroup：`cpu.weight=100`
- B cgroup：`cpu.weight=200`

当 CPU 紧张时，B 理论上能得到约 A 两倍的 CPU 时间。CPU 不紧张时，权重不会强行限制进程使用空闲 CPU。

### 6.5 CPU 绑定：cpuset

cpuset 可以限制进程只能在哪些 CPU 核运行。

相关文件：

```bash
cpuset.cpus
cpuset.cpus.effective
cpuset.mems
cpuset.mems.effective
```

示例：

```bash
echo "0-1" | sudo tee /sys/fs/cgroup/demo/cpuset.cpus
```

表示只允许使用 CPU 0 和 CPU 1。

Docker 对应参数：

```bash
docker run --cpuset-cpus="0-1" ...
```

### 6.6 内存限制：memory.max

设置最大内存：

```bash
echo "100M" | sudo tee /sys/fs/cgroup/demo/memory.max
```

查看当前使用：

```bash
cat /sys/fs/cgroup/demo/memory.current
```

达到 `memory.max` 后，通常会触发 cgroup 内 OOM，内核会选择进程杀掉。

Docker 对应参数：

```bash
docker run -m 100m ...
docker run --memory=100m ...
```

### 6.7 memory.high、memory.low、memory.min

`memory.max` 是硬上限，超过可能 OOM。

`memory.high` 是压力阈值：

- 超过后内核会加强回收和限速。
- 不一定立即 OOM。
- 更适合做温和限制。

```bash
echo "200M" | sudo tee /sys/fs/cgroup/demo/memory.high
```

`memory.low` 是内存保护：

- 当系统内存紧张时，低于该值的内存尽量不被回收。
- 它是尽力而为，不是绝对保证。

`memory.min` 是更强保护：

- 比 `memory.low` 更硬。
- 使用不当可能导致系统其他 cgroup 更容易 OOM。

### 6.8 swap 限制

cgroup v2 中 swap 常见文件：

```bash
memory.swap.current
memory.swap.max
```

禁用某 cgroup 使用 swap：

```bash
echo 0 | sudo tee /sys/fs/cgroup/demo/memory.swap.max
```

Docker 常见参数：

```bash
docker run --memory=300m --memory-swap=300m ...
```

含义通常是内存和内存+swap 总量都限制在 300MB，相当于不允许额外 swap。

### 6.9 OOM 相关

查看内存事件：

```bash
cat /sys/fs/cgroup/demo/memory.events
```

常见字段：

| 字段 | 含义 |
| :--- | :--- |
| `low` | 触发 low 保护相关事件 |
| `high` | 超过 high 阈值次数 |
| `max` | 触达 max 限制次数 |
| `oom` | 发生 OOM 次数 |
| `oom_kill` | OOM kill 次数 |

`memory.oom.group` 控制 OOM 时是否把 cgroup 当成整体处理：

```bash
echo 1 | sudo tee /sys/fs/cgroup/demo/memory.oom.group
```

在多进程服务中，这很重要。比如一个服务由 master 和 worker 组成，只杀某个 worker 可能导致服务处于异常半死状态；按组杀掉更符合“一个工作负载整体失败”的语义。

### 6.10 pids 限制

限制进程数量：

```bash
echo 100 | sudo tee /sys/fs/cgroup/demo/pids.max
```

查看：

```bash
cat /sys/fs/cgroup/demo/pids.current
cat /sys/fs/cgroup/demo/pids.events
```

Docker 对应参数：

```bash
docker run --pids-limit=100 ...
```

它可以防止 fork bomb 把宿主机进程表打满。

### 6.11 IO 限制

cgroup v2 的 IO 控制常用 `io.max`。

先查看设备号：

```bash
lsblk
stat -c '%t:%T' /dev/sda
```

也可以看 `io.stat` 里出现的设备号。

限制读写带宽示例：

```bash
echo "8:0 rbps=10485760 wbps=10485760" | sudo tee /sys/fs/cgroup/demo/io.max
```

含义：

- 设备 `8:0`。
- read bytes per second 限制为 10MB/s。
- write bytes per second 限制为 10MB/s。

Docker 常见参数：

```bash
docker run --device-read-bps /dev/sda:10mb ...
docker run --device-write-bps /dev/sda:10mb ...
docker run --blkio-weight 500 ...
```

注意：不同存储驱动、文件系统、内核版本、cgroup 版本会影响 IO 限制效果。

### 6.12 冻结 cgroup

cgroup v2 可以冻结一组进程：

```bash
echo 1 | sudo tee /sys/fs/cgroup/demo/cgroup.freeze
```

恢复：

```bash
echo 0 | sudo tee /sys/fs/cgroup/demo/cgroup.freeze
```

这类能力可以用于 checkpoint、调试或运行时管理。

---

## 第七章：用 cgroupfs 手动限制进程资源

下面用 cgroup v2 演示手动创建 cgroup、设置 CPU 和内存限制、加入进程。

### 7.1 准备一个消耗资源的进程

可以用 `stress-ng`：

```bash
sudo apt-get install stress-ng
```

启动一个吃 CPU 和内存的进程：

```bash
stress-ng --cpu 4 --vm 1 --vm-bytes 200M --timeout 300s
```

如果没有 `stress-ng`，也可以用简单命令观察 CPU 限制：

```bash
yes > /dev/null
```

### 7.2 创建 cgroup

```bash
sudo mkdir /sys/fs/cgroup/hog_pen
```

查看：

```bash
ls /sys/fs/cgroup/hog_pen
```

### 7.3 设置 CPU 限制

限制为 0.5 核：

```bash
echo "50000 100000" | sudo tee /sys/fs/cgroup/hog_pen/cpu.max
```

查看：

```bash
cat /sys/fs/cgroup/hog_pen/cpu.max
```

### 7.4 设置内存限制

限制为 100MB：

```bash
echo "100M" | sudo tee /sys/fs/cgroup/hog_pen/memory.max
```

查看：

```bash
cat /sys/fs/cgroup/hog_pen/memory.max
cat /sys/fs/cgroup/hog_pen/memory.current
```

### 7.5 把正在运行的进程加入 cgroup

假设目标进程 PID 是：

```bash
pgrep -xo stress-ng
```

加入：

```bash
echo <pid> | sudo tee /sys/fs/cgroup/hog_pen/cgroup.procs
```

确认：

```bash
cat /sys/fs/cgroup/hog_pen/cgroup.procs
cat /proc/<pid>/cgroup
```

### 7.6 启动前就进入 cgroup

如果先启动进程再移动 PID，中间会有一小段不受限制的时间。

更严谨的方法是：先让当前 shell 进入 cgroup，再从这个 shell 启动目标进程。

```bash
echo $$ | sudo tee /sys/fs/cgroup/hog_pen/cgroup.procs
stress-ng --cpu 4 --vm 1 --vm-bytes 200M
```

因为子进程继承父进程 cgroup，所以 `stress-ng` 从一开始就受限制。

注意：把当前交互 shell 移进 cgroup 后，这个 shell 也会受限制。实验结束后建议退出该 shell，或者从另一个终端清理。

### 7.7 观察资源统计

CPU：

```bash
cat /sys/fs/cgroup/hog_pen/cpu.stat
```

内存：

```bash
cat /sys/fs/cgroup/hog_pen/memory.current
cat /sys/fs/cgroup/hog_pen/memory.events
```

进程数：

```bash
cat /sys/fs/cgroup/hog_pen/pids.current
```

### 7.8 删除 cgroup

先确保里面没有进程：

```bash
cat /sys/fs/cgroup/hog_pen/cgroup.procs
```

删除：

```bash
sudo rmdir /sys/fs/cgroup/hog_pen
```

注意：

```bash
sudo rm -rf /sys/fs/cgroup/hog_pen
```

不适合删除 cgroup。cgroupfs 里的控制文件不是普通文件，不能像普通目录那样递归删除。正确方式是 `rmdir` 空 cgroup 目录。

---

## 第八章：用 libcgroup 和 systemd 管理 cgroup

### 8.1 libcgroup 工具

libcgroup 提供一些命令包装 cgroupfs 操作。

安装：

```bash
sudo apt-get install cgroup-tools
```

常用命令：

| 命令 | 作用 |
| :--- | :--- |
| `cgcreate` | 创建 cgroup |
| `cgset` | 设置参数 |
| `cgexec` | 在指定 cgroup 中运行命令 |
| `cgclassify` | 把已有进程移动到 cgroup |
| `cgdelete` | 删除 cgroup |

示例：

```bash
sudo cgcreate -g cpu,memory:/hog_pen2
sudo cgset -r cpu.max="50000 100000" hog_pen2
sudo cgset -r memory.max="100M" hog_pen2
sudo cgexec -g cpu,memory:hog_pen2 stress-ng --cpu 4 --vm 1 --vm-bytes 200M
```

把已有进程移动进去：

```bash
PID=$(pgrep -xo stress-ng)
sudo cgclassify -g cpu,memory:hog_pen2 $PID
```

删除：

```bash
sudo cgdelete -g cpu,memory:/hog_pen2
```

注意：一些发行版已经弱化或弃用 libcgroup，更推荐 systemd 统一管理 cgroup。

### 8.2 systemd 为什么也管 cgroup

systemd 本身就是 Linux 上的服务管理器。它会把系统服务、用户 session、scope、slice 都组织在 cgroup 树中。

查看：

```bash
systemd-cgls --all
systemd-cgtop
```

常见层级：

```text
-.slice
├─system.slice
├─user.slice
└─machine.slice
```

概念：

| systemd 概念 | 含义 |
| :--- | :--- |
| service | 长期运行的服务单元 |
| scope | 外部创建、由 systemd 管理的一组进程 |
| slice | cgroup 资源分组，类似资源控制的父级容器 |

Docker 使用 systemd cgroup driver 时，容器通常会被放到 systemd 管理的 scope 中。

### 8.3 systemd-run 创建临时受限服务

限制一个命令的 CPU 和内存：

```bash
systemd-run -u hog \
  -p CPUQuota=50% \
  -p MemoryMax=100M \
  stress-ng --cpu 4 --vm 1 --vm-bytes 200M
```

查看状态：

```bash
systemctl status hog.service
systemd-cgls --all | grep -A 20 hog
```

清理失败 unit：

```bash
systemctl reset-failed
```

或者创建 transient unit 时使用自动回收：

```bash
systemd-run --collect -u hog -p CPUQuota=50% -p MemoryMax=100M stress-ng --cpu 4
```

### 8.4 systemd slice

slice 适合创建一个可复用的资源分组。

创建：

```bash
sudo tee /etc/systemd/system/hog_pen.slice >/dev/null <<'EOF'
[Slice]
CPUQuota=50%
MemoryMax=100M
EOF
```

加载：

```bash
sudo systemctl daemon-reload
```

在 slice 中启动多个进程：

```bash
systemd-run -u hog1 --slice=hog_pen.slice stress-ng --cpu 2 --vm 1 --vm-bytes 80M
systemd-run -u hog2 --slice=hog_pen.slice stress-ng --cpu 2 --vm 1 --vm-bytes 80M
```

这两个服务会作为 `hog_pen.slice` 的子 cgroup。父 slice 上的限制会约束它们的总资源。

### 8.5 Docker 使用指定 cgroup parent

可以把 Docker 容器放到指定 systemd slice 下：

```bash
docker run -d --name web --cgroup-parent=hog_pen.slice nginx
docker run -d --name redis --cgroup-parent=hog_pen.slice redis
```

这样多个容器可以共享同一个上层资源边界，类似 Kubernetes Pod 里多个容器共享 Pod 级资源管理的思想。

---

## 第九章：Docker 与 cgroup 的关系

### 9.1 Docker 资源参数和 cgroup 文件对应

常见 Docker 参数：

| Docker 参数 | cgroup 作用 |
| :--- | :--- |
| `--cpus=0.5` | 设置 CPU quota，接近 `cpu.max=50000 100000` |
| `--cpu-quota=50000 --cpu-period=100000` | 直接设置 quota/period |
| `--cpu-shares=512` | 设置 CPU 相对权重 |
| `--cpuset-cpus=0-1` | 设置可用 CPU 核 |
| `--memory=100m` | 设置内存上限 |
| `--memory-reservation=80m` | 设置内存软限制语义 |
| `--memory-swap=100m` | 设置内存+swap 限制 |
| `--pids-limit=100` | 限制进程数量 |
| `--blkio-weight=500` | 设置块 IO 权重 |
| `--device-read-bps` | 限制设备读带宽 |
| `--device-write-bps` | 限制设备写带宽 |

示例：

```bash
docker run -d --name limited \
  --cpus=0.5 \
  --memory=100m \
  --pids-limit=100 \
  nginx
```

找到容器主进程：

```bash
PID=$(docker inspect -f '{{.State.Pid}}' limited)
cat /proc/$PID/cgroup
```

再到 `/sys/fs/cgroup` 下查看对应限制。

### 9.2 cgroup driver：systemd 与 cgroupfs

Docker 有 cgroup driver：

```bash
docker info | grep -i cgroup
```

常见：

```text
Cgroup Driver: systemd
Cgroup Version: 2
```

两种 driver：

| driver | 说明 |
| :--- | :--- |
| cgroupfs | Docker 直接操作 cgroupfs |
| systemd | Docker 通过 systemd 管理 cgroup |

在 systemd 发行版上，通常推荐 Docker、containerd、kubelet 统一使用 systemd driver，避免多个管理者同时操作 cgroup 树导致行为不一致。

### 9.3 Docker 容器的 cgroup 路径

查看容器进程：

```bash
PID=$(docker inspect -f '{{.State.Pid}}' <container>)
cat /proc/$PID/cgroup
```

cgroup v2 下可能看到类似：

```text
0::/system.slice/docker-<container-id>.scope
```

或者 rootless、不同运行时环境下路径不同。

找到实际目录：

```bash
CG=$(cat /proc/$PID/cgroup | awk -F: '{print $3}')
sudo ls /sys/fs/cgroup/$CG
```

### 9.4 容器内看到的 cgroup

进入容器：

```bash
docker exec -it limited sh
cat /proc/self/cgroup
```

如果使用 private cgroup namespace，容器内看到的路径可能被裁剪，不暴露宿主机完整层级。

---

## 第十章：Docker 文件系统原理

### 10.1 镜像和容器的文件系统关系

Docker 镜像是只读层的集合。

容器启动时，会在镜像只读层之上增加一个可写层。

```text
container writable layer
------------------------
image layer N
image layer N-1
...
image layer 1
```

容器看到的是这些层合并后的统一目录树。

### 10.2 为什么需要分层

分层带来的好处：

- 多个镜像可以复用相同基础层。
- 多个容器可以复用同一个镜像层。
- 拉取镜像时只下载本地没有的层。
- 构建镜像时可以利用 build cache。
- 容器删除时只删除自己的可写层，不影响镜像。

例如：

```Dockerfile
FROM ubuntu:22.04
RUN apt-get update && apt-get install -y nginx
COPY index.html /var/www/html/index.html
```

每条会改变文件系统的指令，通常都会产生新层。

### 10.3 写时复制 Copy-on-Write

镜像层只读，容器运行时需要修改文件怎么办？

使用 Copy-on-Write：

1. 容器读取文件时，如果可写层没有，就从下层镜像读取。
2. 容器第一次修改某个来自镜像层的文件时，先把文件复制到容器可写层。
3. 后续修改发生在可写层。
4. 下层镜像保持只读不变。

删除文件也不是直接删除下层文件，而是在上层放一个特殊标记，表示该文件在合并视图中不可见。

### 10.4 Docker storage driver

Docker 支持多种 storage driver：

| driver | 说明 |
| :--- | :--- |
| overlay2 | 当前 Linux 上最常见、推荐的驱动 |
| fuse-overlayfs | rootless 场景常见 |
| btrfs | 基于 Btrfs 特性的驱动 |
| zfs | 基于 ZFS 特性的驱动 |
| devicemapper | 历史上常见，现在较少推荐 |
| aufs | 早期 Docker 常见，主线支持和使用已减少 |

查看当前驱动：

```bash
docker info | grep -i "Storage Driver"
```

通常是：

```text
Storage Driver: overlay2
```

---

## 第十一章：overlay2 详细整理

### 11.1 overlayfs 是什么

OverlayFS 是 Linux 内核提供的联合挂载文件系统。

它可以把多个目录叠加成一个统一视图。

核心概念：

| 概念 | 作用 |
| :--- | :--- |
| lowerdir | 下层目录，只读，可以有多个 |
| upperdir | 上层目录，可写 |
| workdir | OverlayFS 内部工作目录 |
| merged | 最终合并后的挂载点 |

挂载形态：

```text
mount -t overlay overlay \
  -o lowerdir=<lower1>:<lower2>,upperdir=<upper>,workdir=<work> \
  <merged>
```

容器进程看到的是 `merged`。

### 11.2 overlay2 在 Docker 中的目录

默认目录：

```bash
/var/lib/docker/overlay2
```

查看：

```bash
sudo ls /var/lib/docker/overlay2
```

一个容器或镜像层可能对应类似目录：

```text
<layer-id>/
  diff/
  link
  lower
  merged/
  work/
```

含义：

| 目录/文件 | 含义 |
| :--- | :--- |
| `diff` | 当前层自己的文件内容 |
| `merged` | 挂载后的合并视图，容器运行时看到的根文件系统 |
| `work` | overlayfs 工作目录 |
| `lower` | 当前层依赖的下层列表 |
| `link` | 短 ID，用于缩短 lowerdir 挂载参数 |
| `l/` | Docker 维护的短链接目录，避免 mount 参数过长 |

### 11.3 容器的 GraphDriver 信息

查看容器 overlay2 路径：

```bash
docker inspect <container> --format '{{json .GraphDriver.Data}}' | jq
```

常见字段：

```json
{
  "LowerDir": "...",
  "MergedDir": "...",
  "UpperDir": "...",
  "WorkDir": "..."
}
```

含义：

| 字段 | 含义 |
| :--- | :--- |
| `LowerDir` | 镜像只读层组成的 lowerdir |
| `UpperDir` | 容器可写层 |
| `MergedDir` | 容器看到的合并目录 |
| `WorkDir` | overlayfs 工作目录 |

进入 `MergedDir` 可以看到容器的根文件系统视图。但不要在生产环境手动修改这些目录，容易破坏 Docker 对元数据和层的管理。

### 11.4 读文件过程

容器读 `/etc/nginx/nginx.conf`：

1. OverlayFS 先查 upperdir。
2. 如果 upperdir 有该文件，返回 upperdir 中的版本。
3. 如果 upperdir 没有，继续查 lowerdir。
4. lowerdir 可能有多层，按优先级从上到下查找。
5. 找到后返回给容器进程。

### 11.5 修改文件过程

容器修改镜像层已有文件：

1. 文件最开始在 lowerdir 中，只读。
2. 第一次写入时触发 copy-up。
3. OverlayFS 把该文件复制到 upperdir。
4. 修改 upperdir 中的副本。
5. 后续读取该路径时，upperdir 中的文件覆盖 lowerdir 中的版本。

这解释了一个现象：

> 容器第一次修改大文件可能比较慢，因为需要先 copy-up 整个文件。

### 11.6 新建文件过程

容器中新建文件：

```bash
echo hello > /tmp/a.txt
```

新文件直接写入 upperdir。

镜像层不会变化。

### 11.7 删除文件过程

如果删除的是 upperdir 中的文件，直接删除。

如果删除的是 lowerdir 中的文件，lowerdir 只读，不能真的删除。OverlayFS 会在 upperdir 中创建 whiteout 标记，让该文件在 merged 视图中不可见。

目录删除还可能涉及 opaque directory 标记，用来隐藏下层同名目录内容。

这解释了一个常见问题：

> Dockerfile 中先添加大文件，后面再 `RUN rm` 删除，并不会让最终镜像真正变小，因为大文件仍然存在于历史层中。

错误示例：

```Dockerfile
COPY big.tar /tmp/big.tar
RUN rm /tmp/big.tar
```

更好的做法是不要把大文件带进层，或者在同一层中下载、使用、删除：

```Dockerfile
RUN curl -L -o /tmp/big.tar https://example.com/big.tar \
    && tar -xf /tmp/big.tar -C /opt/app \
    && rm /tmp/big.tar
```

### 11.8 overlay2 和 image layer

查看镜像历史：

```bash
docker history nginx
```

每一层大致对应一次文件系统变更。

查看镜像详细信息：

```bash
docker image inspect nginx
```

其中 rootfs layers 是内容寻址的 diffID 列表。

### 11.9 容器可写层

容器运行期间的写入默认进入容器可写层。

查看容器大小：

```bash
docker ps -s
```

字段：

| 字段 | 含义 |
| :--- | :--- |
| SIZE | 容器可写层大小 |
| VIRTUAL SIZE | 镜像层 + 可写层的逻辑大小 |

容器删除后，可写层删除：

```bash
docker rm <container>
```

如果数据需要持久化，不应该依赖容器可写层，而应该使用 volume 或 bind mount。

### 11.10 volume 为什么绕开 overlay2

Docker volume 会挂载到容器内某个路径。

```bash
docker run -v mydata:/var/lib/mysql mysql
```

`/var/lib/mysql` 的读写会落到 volume，而不是容器 overlay2 可写层。

好处：

- 数据生命周期独立于容器。
- 数据库这类高写入场景性能更好。
- 避免 copy-on-write 对大文件和频繁写入的影响。

### 11.11 bind mount、volume、tmpfs

| 类型 | 数据位置 | 适合场景 |
| :--- | :--- | :--- |
| bind mount | 宿主机指定路径 | 开发调试、挂配置、挂源码 |
| volume | Docker 管理路径 | 持久化数据，数据库 |
| tmpfs | 内存 | 临时敏感数据、高速临时文件 |

示例：

```bash
docker run -v /host/path:/container/path nginx
docker run -v myvol:/data nginx
docker run --tmpfs /run nginx
```

### 11.12 overlay2 常见注意点

1. 容器可写层不适合做大量频繁写入。
2. 数据库数据应放 volume。
3. Dockerfile 中删除上一层文件不能真正减少上一层大小。
4. 构建镜像时要减少层内无用文件。
5. 不要手动修改 `/var/lib/docker/overlay2`。
6. overlay2 要求底层文件系统支持相关能力，常见是 ext4 或 xfs。
7. XFS 场景通常要求 `ftype=1`，否则 overlay2 可能不可用。

查看 XFS ftype：

```bash
xfs_info /var/lib/docker | grep ftype
```

---

## 第十二章：把容器运行过程串起来

以命令为例：

```bash
docker run -d --name web \
  --cpus=0.5 \
  --memory=256m \
  -p 8080:80 \
  nginx
```

底层大致发生这些事情：

1. Docker CLI 把请求发给 dockerd。
2. dockerd 检查本地是否有 nginx 镜像，没有就拉取。
3. 镜像层存储到 Docker 的内容存储和 overlay2 目录中。
4. dockerd 调用 containerd 创建容器任务。
5. containerd 准备 OCI runtime spec。
6. runc 创建 namespace：
   - PID namespace：容器内进程从 PID 1 开始看。
   - NET namespace：容器有自己的网卡、IP、路由。
   - MNT namespace：容器有自己的挂载视图。
   - UTS namespace：容器有自己的 hostname。
   - IPC namespace：容器有自己的 IPC 资源。
   - USER namespace：如果启用，则做 UID/GID 映射。
   - CGROUP namespace：隔离 cgroup 视图。
7. runc 配置 cgroup：
   - CPU quota 对应 `cpu.max`。
   - 内存限制对应 `memory.max`。
   - pids 限制对应 `pids.max`。
8. runc 准备 rootfs：
   - lowerdir 是 nginx 镜像层。
   - upperdir 是容器可写层。
   - merged 是最终根文件系统。
9. runc 配置安全限制：
   - capabilities。
   - seccomp profile。
   - AppArmor/SELinux label。
10. Docker 配置网络：
   - 创建 veth pair。
   - 一端放进容器 NET namespace 并命名为 eth0。
   - 一端接到 docker0 或自定义 bridge。
   - 设置 IP、路由、NAT、端口映射。
11. 容器主进程启动：
   - nginx master 成为容器内 PID 1。
   - 宿主机上它仍然是一个普通 Linux 进程。

---

## 第十三章：排查命令清单

### 13.1 查看 Docker 基本信息

```bash
docker info
docker version
docker system df
```

重点关注：

```bash
docker info | grep -E 'Storage Driver|Cgroup Driver|Cgroup Version|Docker Root Dir'
```

### 13.2 查看容器主进程

```bash
docker inspect -f '{{.State.Pid}}' <container>
```

### 13.3 查看 namespace

```bash
PID=$(docker inspect -f '{{.State.Pid}}' <container>)
sudo ls -l /proc/$PID/ns
sudo lsns -p $PID
```

比较两个容器是否共享 network namespace：

```bash
sudo readlink /proc/<pid1>/ns/net
sudo readlink /proc/<pid2>/ns/net
```

### 13.4 进入 namespace

```bash
sudo nsenter -t $PID -n ip addr
sudo nsenter -t $PID -n ip route
sudo nsenter -t $PID -m findmnt
sudo nsenter -t $PID -p ps -ef
sudo nsenter -t $PID -a sh
```

### 13.5 查看 cgroup

```bash
cat /proc/$PID/cgroup
systemd-cgls --all
systemd-cgtop
```

cgroup v2：

```bash
CG=$(cat /proc/$PID/cgroup | awk -F: '{print $3}')
sudo cat /sys/fs/cgroup/$CG/cpu.max
sudo cat /sys/fs/cgroup/$CG/memory.max
sudo cat /sys/fs/cgroup/$CG/memory.current
sudo cat /sys/fs/cgroup/$CG/memory.events
sudo cat /sys/fs/cgroup/$CG/pids.max
```

### 13.6 查看 overlay2

```bash
docker inspect <container> --format '{{json .GraphDriver.Data}}' | jq
docker ps -s
docker history <image>
docker image inspect <image>
```

### 13.7 查看容器网络

```bash
docker network ls
docker network inspect bridge
ip addr show docker0
ip link
bridge link
iptables -t nat -S
```

如果系统使用 nftables：

```bash
nft list ruleset
```

### 13.8 查看资源使用

```bash
docker stats
top
htop
free -h
iostat -xz 1
pidstat 1
```

---

## 第十四章：实验清单

### 实验 1：观察容器和宿主机 PID 差异

```bash
docker run -d --name sleep-test alpine sleep 1000
docker exec sleep-test ps
docker inspect -f '{{.State.Pid}}' sleep-test
ps -fp $(docker inspect -f '{{.State.Pid}}' sleep-test)
```

重点观察：

- 容器内 `sleep` 的 PID。
- 宿主机上同一个进程的 PID。

### 实验 2：用 lsns 查看 namespace

```bash
PID=$(docker inspect -f '{{.State.Pid}}' sleep-test)
sudo lsns -p $PID
sudo ls -l /proc/$PID/ns
```

重点观察：

- `pid`、`net`、`mnt`、`uts` namespace inode。
- 和宿主机 shell 的 `/proc/$$/ns` 对比。

### 实验 3：用 nsenter 进入容器网络

```bash
PID=$(docker inspect -f '{{.State.Pid}}' sleep-test)
sudo nsenter -t $PID -n ip addr
sudo nsenter -t $PID -n ip route
```

重点观察：

- 容器 eth0。
- 默认路由。
- IP 是否属于 Docker bridge 网段。

### 实验 4：手动 cgroup 限制 CPU

```bash
sudo mkdir /sys/fs/cgroup/cpu_demo
echo "50000 100000" | sudo tee /sys/fs/cgroup/cpu_demo/cpu.max
yes > /dev/null &
PID=$!
echo $PID | sudo tee /sys/fs/cgroup/cpu_demo/cgroup.procs
top -p $PID
kill $PID
sudo rmdir /sys/fs/cgroup/cpu_demo
```

重点观察：

- CPU 使用率是否被限制在约 50% 单核。
- `cpu.stat` 中 throttling 相关统计是否增长。

### 实验 5：Docker 参数映射到 cgroup

```bash
docker run -d --name limit-test --cpus=0.5 --memory=128m nginx
PID=$(docker inspect -f '{{.State.Pid}}' limit-test)
cat /proc/$PID/cgroup
```

到 `/sys/fs/cgroup` 对应目录查看：

```bash
cat cpu.max
cat memory.max
cat memory.current
```

### 实验 6：观察 overlay2 copy-on-write

```bash
docker run -d --name cow-test alpine sleep 1000
docker ps -s
docker exec cow-test sh -c 'dd if=/dev/zero of=/tmp/bigfile bs=1M count=50'
docker ps -s
docker inspect cow-test --format '{{json .GraphDriver.Data}}' | jq
```

重点观察：

- 容器可写层大小变化。
- `UpperDir` 中新增文件。

### 实验 7：volume 绕开容器可写层

```bash
docker volume create vol-test
docker run -d --name vol-test -v vol-test:/data alpine sleep 1000
docker exec vol-test sh -c 'dd if=/dev/zero of=/data/bigfile bs=1M count=50'
docker ps -s
docker volume inspect vol-test
```

重点观察：

- 写入 volume 后，容器可写层大小不应明显增加。
- 数据在 volume 目录中。

---

## 第十五章：总结

Docker 底层可以抓住四条主线：

1. 容器是进程，不是虚拟机。
2. Namespace 提供隔离视图，让进程看到独立的 PID、网络、挂载点、主机名等。
3. cgroup 提供资源限制和统计，让进程不能无限消耗 CPU、内存、IO、进程数。
4. overlay2 提供分层文件系统，让镜像层只读复用，容器写入落到自己的可写层。

学习顺序建议：

1. 先用 `docker inspect` 找到容器主进程 PID。
2. 再用 `/proc/<pid>/ns`、`lsns`、`nsenter` 看 namespace。
3. 再用 `/proc/<pid>/cgroup` 和 `/sys/fs/cgroup` 看资源限制。
4. 再用 `docker inspect .GraphDriver.Data` 看 overlay2 的 lower、upper、merged、work。
5. 最后把 Docker 参数、Linux 内核文件、实际运行现象对应起来。

把这些对应关系建立起来后，Docker 就不再是黑盒：

```text
docker run
  = 创建一个受 namespace 隔离的进程
  + 把它放进 cgroup 资源组
  + 给它挂一个 overlay2 rootfs
  + 给它配置网络和安全策略
```

