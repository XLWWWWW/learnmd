# cgroup 资源控制实践笔记

本文整理自 Ivan Velichko 的《Controlling Process Resources with Linux Control Groups》，并补充一些日常使用中的经验、排查方法和容易踩坑的细节。

这篇博客的主线很清楚：先从最底层的 `cgroupfs` 手工操作开始，再过渡到 `libcgroup` 工具，最后引出在现代 systemd 发行版上更推荐的 `systemd-run` 和 `slice`。它不是完整手册，但很适合作为理解 Docker、containerd、Kubernetes 资源限制底层机制的实践入口。

---

## 1. 对这篇博客的评价

### 1.1 优点

这篇文章的价值主要在于“从内核接口到上层工具”的路径很完整。

它没有一开始就讲 Docker 或 Kubernetes 的抽象，而是先展示：

```bash
mkdir /sys/fs/cgroup/hog_pen
echo "50000 100000" > /sys/fs/cgroup/hog_pen/cpu.max
echo "100M" > /sys/fs/cgroup/hog_pen/memory.max
echo <pid> > /sys/fs/cgroup/hog_pen/cgroup.procs
```

这样能让人直接看到 cgroup 的本质：

> cgroup 不是一个单独的后台服务，而是 Linux 内核提供的资源控制机制。用户态主要通过 `cgroupfs` 这个伪文件系统来创建、配置和移动进程。

它还把三个层次串起来了：

| 层次 | 代表方式 | 特点 |
| :--- | :--- | :--- |
| 内核接口 | 直接操作 `/sys/fs/cgroup` | 最底层、最直观、也最容易和系统管理器冲突 |
| 工具封装 | `cgcreate`、`cgset`、`cgexec` | 命令更顺手，但很多发行版已不推荐作为主流方式 |
| systemd 管理 | `systemd-run`、`.slice` | 现代 systemd 系统上更推荐，能和系统服务、容器运行时统一管理 |

这对理解 Docker 的资源限制非常有帮助，因为 Docker 的 `--memory`、`--cpus`、`--pids-limit` 等参数，最终也会落到 cgroup 控制器上。

### 1.2 不足和需要补全的地方

文章偏实践演示，有些生产和日常排查中很重要的细节没有完全展开。

主要包括：

- 没有详细解释 cgroup v2 中 `cgroup.subtree_control` 和 controller 启用规则。
- 没有强调 `sudo echo ... > file` 经常无效，因为重定向由当前 shell 执行。
- 没有区分 `cpu.max` 硬限制和 `cpu.weight` 相对权重。
- 没有展开 `memory.high`、`memory.max`、`memory.oom.group` 的不同语义。
- 对 systemd slice 的层级命名、资源继承、临时 unit 清理可以再补充。
- 对 Docker/Kubernetes 日常配置和排查命令可以再补充。

结论：

> 这篇博客适合建立 cgroup 实践直觉，但如果要用于日常运维、容器排查或生产配置，需要补上 systemd、cgroup v2、Docker/Kubernetes 相关经验。

---

## 2. cgroup 的核心理解

### 2.1 cgroup 是什么

cgroup，全称 Control Groups，是 Linux 内核提供的资源管理机制。它可以把进程组织成层级树，并对这棵树上的进程进行资源限制、统计和控制。

典型用途：

- 限制某个进程最多使用多少 CPU。
- 限制某个进程最多使用多少内存。
- 限制磁盘 IO。
- 限制最多能创建多少进程。
- 统计一组进程消耗了多少资源。
- 给 systemd service、Docker 容器、Kubernetes Pod 做资源边界。

一句话：

> namespace 负责“看到什么”，cgroup 负责“能用多少”。

### 2.2 cgroupfs 是内核接口

cgroup 的主要用户态接口是一个伪文件系统，通常挂载在：

```bash
/sys/fs/cgroup
```

查看当前系统使用的是 cgroup v1 还是 v2：

```bash
stat -fc %T /sys/fs/cgroup
```

如果输出：

```text
cgroup2fs
```

说明是 cgroup v2。

也可以看挂载信息：

```bash
mount -l | grep cgroup
findmnt /sys/fs/cgroup
```

cgroup 操作本质上是文件系统操作：

| 操作 | 对应行为 |
| :--- | :--- |
| `mkdir /sys/fs/cgroup/demo` | 创建一个 cgroup |
| 写 `cpu.max`、`memory.max` | 设置资源限制 |
| 写 `cgroup.procs` | 把进程移动到这个 cgroup |
| 读 `memory.current`、`cpu.stat` | 查看资源统计 |
| `rmdir /sys/fs/cgroup/demo` | 删除空 cgroup |

注意：`/sys/fs/cgroup` 不是普通目录。里面很多文件是内核暴露的控制文件，不能像普通文件那样随意 `rm`。

---

## 3. cgroup v2 常见文件

在 cgroup v2 中，一个 cgroup 目录里常见文件如下：

| 文件 | 作用 |
| :--- | :--- |
| `cgroup.controllers` | 当前 cgroup 可用的 controller |
| `cgroup.subtree_control` | 控制下一层子 cgroup 可以使用哪些 controller |
| `cgroup.procs` | 当前 cgroup 中的进程 PID |
| `cgroup.threads` | 当前 cgroup 中的线程 TID |
| `cgroup.events` | cgroup 状态事件 |
| `cgroup.freeze` | 冻结或恢复 cgroup 中的进程 |
| `cpu.max` | CPU quota 和 period |
| `cpu.weight` | CPU 相对权重 |
| `cpu.stat` | CPU 使用统计 |
| `memory.current` | 当前内存使用量 |
| `memory.max` | 内存硬上限 |
| `memory.high` | 内存压力阈值，超过后会被回收和限速 |
| `memory.low` | 内存保护，尽量不被回收 |
| `memory.min` | 更强的内存保护 |
| `memory.swap.max` | swap 使用上限 |
| `memory.events` | OOM、high、max 等内存事件计数 |
| `memory.oom.group` | OOM 时是否按 cgroup 整组杀掉 |
| `io.max` | 块设备 IO 限制 |
| `io.stat` | IO 统计 |
| `pids.current` | 当前进程数 |
| `pids.max` | 最大进程数 |

查看当前可用 controller：

```bash
cat /sys/fs/cgroup/cgroup.controllers
```

常见输出：

```text
cpuset cpu io memory hugetlb pids rdma misc
```

### 3.1 父子 cgroup 和 `cgroup.subtree_control`

理解 cgroup v2 时，最容易混淆的是这三个东西：

```text
父 cgroup 的 cgroup.subtree_control
子 cgroup 自己的 cpu.max / memory.max / pids.max
子 cgroup 自己的 cgroup.subtree_control
```

它们的职责完全不同。

假设目录结构是：

```text
/sys/fs/cgroup/
└── xlwwwww/
    └── app1/
```

那么关系是：

| 位置 | 控制什么 |
| :--- | :--- |
| `/sys/fs/cgroup/cgroup.subtree_control` | 控制 `/sys/fs/cgroup/xlwwwww` 这一层子 cgroup 能使用哪些 controller |
| `/sys/fs/cgroup/xlwwwww/cpu.max` | 控制 `xlwwwww` 这个 cgroup 自己的 CPU 上限 |
| `/sys/fs/cgroup/xlwwwww/memory.max` | 控制 `xlwwwww` 这个 cgroup 自己的内存上限 |
| `/sys/fs/cgroup/xlwwwww/cgroup.subtree_control` | 控制 `xlwwwww/app1` 这一层子 cgroup 能使用哪些 controller |
| `/sys/fs/cgroup/xlwwwww/app1/cpu.max` | 控制 `app1` 这个 cgroup 自己的 CPU 上限 |

一句话：

```text
谁的 subtree_control，控制谁的下一层子 cgroup。
谁的 cpu.max / memory.max，控制谁自己。
```

例如根 cgroup 中已经启用了：

```bash
cat /sys/fs/cgroup/cgroup.subtree_control
```

输出：

```text
cpu memory pids
```

这表示根 cgroup 的直接子 cgroup 可以使用 `cpu`、`memory`、`pids` controller。

所以创建：

```bash
sudo mkdir /sys/fs/cgroup/xlwwwww
```

之后，`xlwwwww` 这个子 cgroup 中通常可以使用：

```text
cpu.max
cpu.weight
memory.max
memory.current
pids.max
pids.current
```

然后真正限制 `xlwwwww` 自己资源时，写的是子 cgroup 自己的资源文件：

```bash
echo "10000 100000" | sudo tee /sys/fs/cgroup/xlwwwww/cpu.max
echo "500M" | sudo tee /sys/fs/cgroup/xlwwwww/memory.max
echo "100" | sudo tee /sys/fs/cgroup/xlwwwww/pids.max
```

如果你希望 `xlwwwww` 下面继续创建 `app1`，并且 `app1` 也能使用 `cpu` 和 `memory` controller，才需要改 `xlwwwww` 自己的 `cgroup.subtree_control`：

```bash
echo "+cpu +memory" | sudo tee /sys/fs/cgroup/xlwwwww/cgroup.subtree_control
sudo mkdir /sys/fs/cgroup/xlwwwww/app1
echo "50000 100000" | sudo tee /sys/fs/cgroup/xlwwwww/app1/cpu.max
echo "200M" | sudo tee /sys/fs/cgroup/xlwwwww/app1/memory.max
```

这里 `xlwwwww/cgroup.subtree_control` 不是用来限制 `xlwwwww` 自己的，而是用来决定 `xlwwwww` 的子 cgroup 能用哪些 controller。

### 3.2 `subtree_control` 不是资源限制值

`cgroup.subtree_control` 只决定 controller 是否分发给下一层子 cgroup，不负责设置资源大小。

它不表达：

```text
CPU 最多多少
内存最多多少
进程数最多多少
```

真正的资源限制值写在这些文件里：

```text
cpu.max
cpu.weight
memory.max
memory.high
pids.max
io.max
```

所以完整关系是：

```text
父 cgroup 的 cgroup.subtree_control
  -> 决定子 cgroup 能不能使用某类 controller

父 cgroup 的 cpu.max / memory.max
  -> 限制父 cgroup 整个子树的总资源边界

子 cgroup 的 cpu.max / memory.max
  -> 限制子 cgroup 自己的资源上限

子 cgroup 的 cgroup.subtree_control
  -> 决定孙子 cgroup 能不能使用某类 controller
```

举个边界继承的例子：

```text
parent memory.max = 1G
child  memory.max = 2G
```

即使 `child` 写了 `2G`，它实际也不能突破父 cgroup `1G` 的总边界。父级资源限制会约束整个子树。

### 3.3 `chown` 委派和 controller 启用是两件事

日常实验中，经常会这样做：

```bash
sudo mkdir /sys/fs/cgroup/xlwwwww
sudo chown -R laborant:laborant /sys/fs/cgroup/xlwwwww
```

这叫把某个 cgroup 目录委派给普通用户管理。

但要分清楚：

| 动作 | 决定什么 |
| :--- | :--- |
| 父 cgroup 的 `cgroup.subtree_control` | 子 cgroup 能使用哪些 controller |
| `chown` 子 cgroup 目录 | 普通用户有没有权限写这个目录里的控制文件 |

也就是说，`chown` 只解决 Linux 文件权限问题，不会凭空增加 controller 能力。

如果父级没有启用 `io` controller，普通用户即使拥有子目录写权限，也不能靠自己让 `io.max` 生效。

### 3.4 手动创建控制文件没有意义

`/sys/fs/cgroup` 是伪文件系统，不是普通目录。`cpu.max`、`memory.max`、`pids.max` 这些文件不是用户手写出来的配置文件，而是内核根据 controller 状态暴露出来的接口。

如果父级没有启用某个 controller，子 cgroup 通常不会出现对应的有效控制文件。手动创建同名文件一般会失败：

```bash
touch /sys/fs/cgroup/xlwwwww/io.max
```

即使在某些特殊环境中看到类似名字，也要以内核 controller 是否启用、文件是否能正常写入并影响资源行为为准。

判断方法：

```bash
cat /sys/fs/cgroup/cgroup.subtree_control
ls /sys/fs/cgroup/xlwwwww
```

一句话：

```text
cgroup 控制文件不是普通配置文件，而是内核 controller 暴露的接口。
父级没有启用 controller，子级不能手动造一个 max 文件让它生效。
```

---

## 4. 用 cgroupfs 手工限制进程资源

下面以 cgroup v2 为例。

### 4.1 创建 cgroup

```bash
sudo mkdir /sys/fs/cgroup/hog_pen
```

查看是否创建成功：

```bash
ls /sys/fs/cgroup/hog_pen
```

如果目录里没有看到 `cpu.max`、`memory.max` 等文件，通常说明父 cgroup 没有给子 cgroup 启用对应 controller。

### 4.2 补全：启用子树 controller

cgroup v2 中，父 cgroup 必须在 `cgroup.subtree_control` 中启用 controller，子 cgroup 才能使用。

查看根 cgroup 可用 controller：

```bash
cat /sys/fs/cgroup/cgroup.controllers
```

启用 CPU、memory、pids：

```bash
echo "+cpu +memory +pids" | sudo tee /sys/fs/cgroup/cgroup.subtree_control
```

然后重新检查子目录：

```bash
ls /sys/fs/cgroup/hog_pen
```

注意：

- 在 systemd 管理的系统上，不建议随意改根 cgroup。
- 有些环境里根 cgroup 的 controller 已经由 systemd 管理，普通实验最好放在可控的测试机或容器环境里。
- cgroup v2 有“no internal process”规则：非根 cgroup 如果要给子 cgroup 启用 domain controller，通常不能同时在该 cgroup 自己里面放业务进程。

### 4.3 设置 CPU 硬限制

`cpu.max` 格式：

```text
<quota> <period>
```

单位是微秒。

限制为 0.5 核：

```bash
echo "50000 100000" | sudo tee /sys/fs/cgroup/hog_pen/cpu.max
```

含义：

- `100000`：一个周期是 100000 微秒，也就是 100ms。
- `50000`：每个周期最多运行 50000 微秒。
- 整体等价于最多使用 0.5 个 CPU 核。

常见换算：

| 目标 | `cpu.max` |
| :--- | :--- |
| 不限制 | `max 100000` |
| 0.5 核 | `50000 100000` |
| 1 核 | `100000 100000` |
| 2 核 | `200000 100000` |

在多核机器上，CPU 使用率可以超过 100%。例如 4 核机器上跑满所有核可以理解为 400%。

### 4.4 设置 CPU 相对权重

`cpu.max` 是硬限制，`cpu.weight` 是相对权重。

```bash
cat /sys/fs/cgroup/hog_pen/cpu.weight
echo 200 | sudo tee /sys/fs/cgroup/hog_pen/cpu.weight
```

cgroup v2 中 `cpu.weight` 范围通常是 `1-10000`，默认值是 `100`。

举例：

| cgroup | `cpu.weight` |
| :--- | :--- |
| A | 100 |
| B | 200 |

当 CPU 紧张时，B 理论上能拿到大约 A 两倍的 CPU 时间。CPU 不紧张时，权重不会主动限制进程使用空闲 CPU。

日常经验：

- 想要“最多只能用多少”，用 `cpu.max`。
- 想要“大家抢 CPU 时谁优先”，用 `cpu.weight`。
- Docker 的 `--cpus` 更接近 CPU quota。
- Docker 的 `--cpu-shares` 更接近 CPU 权重。

### 4.5 设置内存限制

限制内存硬上限为 100MB：

```bash
echo "100M" | sudo tee /sys/fs/cgroup/hog_pen/memory.max
```

查看：

```bash
cat /sys/fs/cgroup/hog_pen/memory.max
cat /sys/fs/cgroup/hog_pen/memory.current
```

当 cgroup 中的进程接近或超过 `memory.max` 时，内核会触发 OOM，通常会杀掉其中某个进程。

### 4.6 补全：`memory.high` 和 `memory.max` 的区别

日常更推荐理解这几个内存参数：

| 参数 | 语义 |
| :--- | :--- |
| `memory.max` | 硬上限。超过后可能触发 OOM |
| `memory.high` | 压力阈值。超过后内核会加强回收并限制进程速度 |
| `memory.low` | 软保护。内存紧张时尽量少回收这组进程 |
| `memory.min` | 更强保护。比 `memory.low` 更硬 |
| `memory.swap.max` | swap 使用上限 |

例如：

```bash
echo "80M" | sudo tee /sys/fs/cgroup/hog_pen/memory.high
echo "100M" | sudo tee /sys/fs/cgroup/hog_pen/memory.max
```

含义是：

- 超过 80MB 后开始明显感受到内存压力和回收。
- 超过 100MB 后可能触发 OOM。

生产经验：

- 只设置 `memory.max` 很容易表现为“突然被 OOM kill”。
- 配合 `memory.high` 可以更早暴露内存压力。
- 排查内存问题时要看 `memory.events`，不要只看进程是否还活着。

查看内存事件：

```bash
cat /sys/fs/cgroup/hog_pen/memory.events
```

可能看到：

```text
low 0
high 12
max 3
oom 1
oom_kill 1
```

这说明 cgroup 曾经超过过 high/max，并发生过 OOM kill。

### 4.7 把已有进程加入 cgroup

假设已经启动了一个进程：

```bash
yes > /dev/null
```

找到它的 PID：

```bash
pgrep -xo yes
```

移动到 cgroup：

```bash
echo <pid> | sudo tee /sys/fs/cgroup/hog_pen/cgroup.procs
```

确认：

```bash
cat /sys/fs/cgroup/hog_pen/cgroup.procs
cat /proc/<pid>/cgroup
```

注意：

```bash
sudo echo <pid> > /sys/fs/cgroup/hog_pen/cgroup.procs
```

这种写法经常无效。原因是 `echo` 被 `sudo` 了，但 `>` 重定向是当前 shell 执行的，当前 shell 可能没有权限写目标文件。

推荐写法：

```bash
echo <pid> | sudo tee /sys/fs/cgroup/hog_pen/cgroup.procs
```

### 4.8 让进程从一开始就受限制

如果先启动进程，再把 PID 写入 `cgroup.procs`，中间会有一小段时间不受限制。

更严谨的方式是先把当前 shell 放进 cgroup，然后从这个 shell 启动目标进程：

```bash
echo $$ | sudo tee /sys/fs/cgroup/hog_pen/cgroup.procs
./your_program
```

因为子进程默认继承父进程所在的 cgroup，所以 `your_program` 从启动第一刻起就受限制。

注意：

- 这个 shell 自己也会被限制。
- 后续从这个 shell 启动的所有子进程也会继承这个 cgroup。
- 实验结束后可以退出该 shell，或者从另一个终端处理清理。

### 4.9 删除 cgroup

先确认没有进程：

```bash
cat /sys/fs/cgroup/hog_pen/cgroup.procs
```

删除：

```bash
sudo rmdir /sys/fs/cgroup/hog_pen
```

如果报错：

```text
Device or resource busy
```

通常说明里面还有进程，或者还有子 cgroup。

不要使用：

```bash
sudo rm -rf /sys/fs/cgroup/hog_pen
```

cgroupfs 里的控制文件不是普通文件，正确删除方式是 `rmdir` 空 cgroup 目录。

---

## 5. 用 libcgroup 工具管理 cgroup

`libcgroup` 提供了对 cgroupfs 的命令封装。

安装：

```bash
sudo apt-get install cgroup-tools
```

常用命令：

| 命令 | 作用 |
| :--- | :--- |
| `cgcreate` | 创建 cgroup |
| `cgset` | 设置 cgroup 参数 |
| `cgexec` | 在指定 cgroup 中运行命令 |
| `cgclassify` | 把已有进程移动到 cgroup |
| `cgdelete` | 删除 cgroup |

### 5.1 创建 cgroup

```bash
sudo cgcreate -g cpu,memory:/hog_pen2
```

含义：

```text
在 cpu 和 memory controller 下创建名为 hog_pen2 的 cgroup
```

在 cgroup v1 中，它可能对应：

```text
/sys/fs/cgroup/cpu/hog_pen2
/sys/fs/cgroup/memory/hog_pen2
```

在 cgroup v2 中，表现会受系统和工具版本影响，因为 cgroup v2 是统一层级。

### 5.2 设置限制

```bash
sudo cgset -r cpu.max="50000 100000" hog_pen2
sudo cgset -r memory.max="100M" hog_pen2
```

### 5.3 把已有进程移动进去

```bash
PID=$(pgrep -xo your_program)
sudo cgclassify -g cpu,memory:hog_pen2 "$PID"
```

### 5.4 直接在 cgroup 中启动进程

```bash
sudo cgexec -g cpu,memory:hog_pen2 ./your_program
```

这比“先启动再移动 PID”更好，因为进程从启动开始就处于目标 cgroup。

### 5.5 删除 cgroup

```bash
sudo cgdelete -g cpu,memory:/hog_pen2
```

### 5.6 日常评价

`libcgroup` 的好处是命令直观，比手写 `/sys/fs/cgroup` 文件方便。

但在现代 systemd 发行版上，它不是首选方案。原因是 systemd 自己已经在管理 cgroup 树，如果再让 `libcgroup` 或手写脚本直接操作同一棵树，容易出现多个管理者互相覆盖、清理不一致、层级不符合预期的问题。

结论：

> 学习 cgroup 原理时可以用 `libcgroup`；在 systemd 生产环境中，优先考虑 systemd 的资源控制方式。

---

## 6. 用 systemd-run 运行受限进程

在 systemd 系统上，`systemd-run` 是日常限制单个命令资源的常用方式。

### 6.1 后台运行一个受限命令

```bash
systemd-run -u hog \
  -p CPUQuota=50% \
  -p MemoryMax=100M \
  ./your_program
```

含义：

- 创建一个临时 service：`hog.service`。
- 设置 CPU 最多 50%。
- 设置内存最多 100MB。
- 由 systemd 把进程放入对应 cgroup。

查看状态：

```bash
systemctl status hog.service
```

查看 cgroup 树：

```bash
systemd-cgls --all
```

查看资源排行：

```bash
systemd-cgtop
```

### 6.2 前台运行并等待命令结束

如果希望像普通命令一样在当前终端等待结果，可以用：

```bash
systemd-run --wait --pty \
  -p CPUQuota=50% \
  -p MemoryMax=100M \
  ./your_program
```

常用参数：

| 参数 | 作用 |
| :--- | :--- |
| `--wait` | 等待 transient unit 结束并显示结果 |
| `--pty` | 分配交互式终端 |
| `--same-dir` | 在当前目录运行 |
| `--collect` | unit 不再被引用时自动回收 |
| `-u name` | 指定 unit 名 |
| `-p Key=Value` | 设置 systemd unit 属性 |

示例：

```bash
systemd-run --wait --pty --same-dir --collect \
  -u limited-shell \
  -p CPUQuota=50% \
  -p MemoryMax=512M \
  bash
```

这样可以得到一个受限的交互式 shell。

### 6.3 transient unit 清理问题

`systemd-run` 创建的是 transient unit。

如果命令成功结束，systemd 通常会自动清理这个临时 unit。但如果 one-shot 命令失败，失败状态可能会留下。

下次使用同名 unit 时可能报错：

```text
Failed to start transient service unit:
Unit <name>.service was already loaded or has a fragment file.
```

解决方式：

```bash
systemctl reset-failed <name>.service
```

或者清理所有 failed unit：

```bash
sudo systemctl reset-failed
```

更好的习惯是创建时加上：

```bash
systemd-run --collect -u hog -p CPUQuota=50% -p MemoryMax=100M ./your_program
```

`--collect` 会让 systemd 在 unit 不再被引用时进行回收，包括失败的 transient unit。

### 6.4 常见 systemd 资源属性

| systemd 属性 | 对应 cgroup 语义 |
| :--- | :--- |
| `CPUQuota=50%` | CPU 硬限制，接近 `cpu.max` |
| `CPUWeight=100` | CPU 相对权重，接近 `cpu.weight` |
| `AllowedCPUs=0-1` | 限制可用 CPU 核 |
| `MemoryMax=100M` | 内存硬上限，接近 `memory.max` |
| `MemoryHigh=80M` | 内存压力阈值，接近 `memory.high` |
| `MemorySwapMax=0` | swap 上限 |
| `TasksMax=100` | 进程/线程数量上限，接近 `pids.max` |
| `IOWeight=100` | IO 权重 |
| `IOReadBandwidthMax=` | 读带宽限制 |
| `IOWriteBandwidthMax=` | 写带宽限制 |

临时修改已有服务：

```bash
sudo systemctl set-property nginx.service CPUQuota=50% MemoryMax=512M
```

查看属性：

```bash
systemctl show nginx.service -p CPUQuotaPerSecUSec -p MemoryMax -p ControlGroup
```

---

## 7. systemd slice 怎么理解

`slice` 不是简单的“列表”。它更像是 systemd 在 cgroup 树里创建和管理的一个分组节点。

可以理解为：

```text
slice = cgroup 目录节点 + 资源边界 + 层级分组
```

典型结构：

```text
-.slice
├─system.slice
│ ├─sshd.service
│ ├─docker.service
│ └─nginx.service
├─user.slice
│ └─user-1000.slice
│   └─session-2.scope
└─machine.slice
```

常见 slice：

| slice | 作用 |
| :--- | :--- |
| `system.slice` | 系统服务 |
| `user.slice` | 用户会话 |
| `machine.slice` | 虚拟机、容器、machine 类实例 |
| 自定义 `.slice` | 自己定义的资源分组 |

### 7.1 创建持久 slice

创建文件：

```bash
sudo tee /etc/systemd/system/hog_pen.slice >/dev/null <<'EOF'
[Slice]
CPUQuota=50%
MemoryMax=100M
EOF
```

重新加载：

```bash
sudo systemctl daemon-reload
```

在这个 slice 中启动服务：

```bash
systemd-run -u hog1 --slice=hog_pen.slice ./your_program
systemd-run -u hog2 --slice=hog_pen.slice ./your_program
```

效果：

```text
hog_pen.slice
├─hog1.service
└─hog2.service
```

`hog_pen.slice` 上的限制会约束整个子树，所以 `hog1` 和 `hog2` 的总资源会受父 slice 限制。

### 7.2 slice 的资源继承

如果父 slice 设置了：

```ini
[Slice]
CPUQuota=50%
MemoryMax=100M
```

它下面所有 service/scope 的累计资源会被限制在这个边界里。

也就是说：

```text
父 slice 限制的是整个分组的总量
子 service 限制的是单个服务自己的上限
```

这是 slice 非常适合做“团队级、应用组级、实验组级”资源边界的原因。

### 7.3 slice 命名规则

systemd slice 名字体现层级。

例如：

```text
app.slice
app-frontend.slice
app-backend.slice
```

层级含义大致是：

```text
app.slice
├─app-frontend.slice
└─app-backend.slice
```

注意不是用 `/` 表示层级，而是用 `-` 编码 slice 名称层级。

---

## 8. Docker 与 cgroup

### 8.1 Docker 参数和 cgroup 的关系

常见 Docker 资源参数：

| Docker 参数 | cgroup 语义 |
| :--- | :--- |
| `--cpus=0.5` | CPU quota，接近 `cpu.max=50000 100000` |
| `--cpu-quota=50000 --cpu-period=100000` | 直接设置 quota/period |
| `--cpu-shares=512` | CPU 相对权重 |
| `--cpuset-cpus=0-1` | 绑定可用 CPU 核 |
| `--memory=100m` | 内存硬上限 |
| `--memory-reservation=80m` | 内存软限制语义 |
| `--memory-swap=100m` | 内存加 swap 总限制 |
| `--pids-limit=100` | 进程数量限制 |
| `--blkio-weight=500` | 块 IO 权重 |
| `--device-read-bps` | 设备读带宽限制 |
| `--device-write-bps` | 设备写带宽限制 |

示例：

```bash
docker run -d --name limited-nginx \
  --cpus=0.5 \
  --memory=100m \
  --pids-limit=100 \
  nginx
```

查看容器主进程 PID：

```bash
PID=$(docker inspect -f '{{.State.Pid}}' limited-nginx)
cat /proc/$PID/cgroup
```

### 8.2 Docker cgroup driver

查看 Docker cgroup 配置：

```bash
docker info | grep -i cgroup
```

可能看到：

```text
Cgroup Driver: systemd
Cgroup Version: 2
```

常见 driver：

| driver | 含义 |
| :--- | :--- |
| `cgroupfs` | Docker 直接操作 cgroupfs |
| `systemd` | Docker 通过 systemd 管理 cgroup |

在现代 systemd 发行版上，通常推荐 Docker、containerd、kubelet 统一使用 `systemd` cgroup driver。

原因：

- systemd 本身已经管理系统服务和用户会话的 cgroup。
- 容器运行时如果也直接操作 cgroupfs，容易出现多个管理者并存。
- Kubernetes 官方实践中也推荐 kubelet 和容器运行时使用一致的 cgroup driver。

### 8.3 把 Docker 容器放入指定 slice

创建一个 slice 后，可以让多个容器共享同一个上层资源边界：

```bash
docker run -d --name web --cgroup-parent=hog_pen.slice nginx
docker run -d --name redis --cgroup-parent=hog_pen.slice redis
```

效果类似：

```text
hog_pen.slice
├─docker-<web-container-id>.scope
└─docker-<redis-container-id>.scope
```

用途：

- 给一组容器设置总 CPU/内存上限。
- 模拟 Kubernetes Pod 里多个容器共享一组资源边界的思想。
- 对 Docker 没有直接暴露的 cgroup 参数进行更细粒度管理。

注意：实际路径会受到 Docker 版本、rootless 模式、cgroup driver、systemd 版本影响。

---

## 9. Kubernetes 与 cgroup

Kubernetes 的资源限制本质上也会落到 cgroup。

Pod 和容器常见资源配置：

```yaml
resources:
  requests:
    cpu: "500m"
    memory: "256Mi"
  limits:
    cpu: "1"
    memory: "512Mi"
```

大致理解：

| Kubernetes 字段 | 含义 |
| :--- | :--- |
| `requests.cpu` | 调度时的 CPU 需求，也会影响 CPU 权重 |
| `limits.cpu` | CPU 上限，通常落到 CPU quota |
| `requests.memory` | 调度时的内存需求 |
| `limits.memory` | 内存硬上限，超过可能 OOMKilled |

查看 kubelet cgroup driver：

```bash
grep -i cgroupDriver /var/lib/kubelet/config.yaml
```

常见配置：

```yaml
cgroupDriver: systemd
```

日常经验：

- kubelet 和 container runtime 的 cgroup driver 要保持一致。
- 在 systemd 机器上，优先使用 `systemd` driver。
- Pod 出现 `OOMKilled` 时，不只看应用日志，也要看节点上的 cgroup 事件和 kubelet 事件。

---

## 10. 日常排查命令

### 10.1 判断系统 cgroup 版本

```bash
stat -fc %T /sys/fs/cgroup
```

```bash
mount -l | grep cgroup
```

### 10.2 查看当前进程在哪个 cgroup

```bash
cat /proc/$$/cgroup
```

查看指定进程：

```bash
cat /proc/<pid>/cgroup
```

### 10.3 查看 systemd cgroup 树

```bash
systemd-cgls --all
```

### 10.4 查看 cgroup 资源排行

```bash
systemd-cgtop
```

### 10.5 查看某个 systemd unit 的 cgroup

```bash
systemctl show nginx.service -p ControlGroup
```

### 10.6 查看 cgroup 内进程

```bash
cat /sys/fs/cgroup/<path>/cgroup.procs
```

### 10.7 查看 CPU 统计

```bash
cat /sys/fs/cgroup/<path>/cpu.stat
```

常见字段：

| 字段 | 含义 |
| :--- | :--- |
| `usage_usec` | 总 CPU 使用时间 |
| `user_usec` | 用户态 CPU 时间 |
| `system_usec` | 内核态 CPU 时间 |
| `nr_periods` | CPU quota 周期数 |
| `nr_throttled` | 被限流次数 |
| `throttled_usec` | 被限流总时长 |

如果 `nr_throttled` 和 `throttled_usec` 增长很快，说明 CPU quota 确实在限制进程。

### 10.8 查看内存统计

```bash
cat /sys/fs/cgroup/<path>/memory.current
cat /sys/fs/cgroup/<path>/memory.max
cat /sys/fs/cgroup/<path>/memory.events
cat /sys/fs/cgroup/<path>/memory.stat
```

重点看：

- `memory.current`：当前用量。
- `memory.max`：硬上限。
- `memory.events`：是否发生 high、max、oom、oom_kill。
- `memory.stat`：匿名页、文件缓存、slab 等细分。

### 10.9 查看进程数限制

```bash
cat /sys/fs/cgroup/<path>/pids.current
cat /sys/fs/cgroup/<path>/pids.max
```

进程数达到上限时，应用可能报：

```text
fork: Resource temporarily unavailable
```

或者线程创建失败。

---

## 11. 日常使用技巧和经验

### 11.1 学习实验可以直接操作 cgroupfs，生产优先 systemd

直接写 `/sys/fs/cgroup` 最适合理解原理。

但在 systemd 系统上，生产环境更建议用：

```bash
systemd-run
systemctl set-property
systemctl edit
```

原因是 systemd 已经是事实上的 cgroup 管理者。让 systemd、Docker、Kubernetes、手写脚本同时直接管理同一棵 cgroup 树，容易产生不一致。

### 11.2 临时限制一个命令，优先用 systemd-run

例如临时跑一个可能吃资源的脚本：

```bash
systemd-run --wait --pty --same-dir --collect \
  -u risky-job \
  -p CPUQuota=100% \
  -p MemoryMax=1G \
  bash -lc './run.sh'
```

这样比手动创建 cgroup 更容易清理，也更符合 systemd 系统的管理方式。

### 11.3 限制 CPU 时先想清楚是“硬上限”还是“相对公平”

硬上限：

```bash
CPUQuota=50%
```

适合防止一个任务把 CPU 吃满。

相对权重：

```bash
CPUWeight=200
```

适合多个服务竞争 CPU 时调整优先级。

经验：

- 批处理、压测、爬虫、编译任务：常用 CPU quota。
- 多个在线服务共享机器：更常用 CPU weight 配合必要的 quota。

### 11.4 内存限制不要只看 OOM

内存问题排查顺序：

```bash
cat memory.current
cat memory.max
cat memory.high
cat memory.events
cat memory.stat
```

重点看：

- `memory.events` 中 `oom_kill` 是否增加。
- `high` 是否持续增加。
- `memory.stat` 中 file cache、anon、slab 哪个占比高。

### 11.5 OOM 时是否杀整个进程组

有些应用是多进程模型。如果其中一个进程 OOM，被杀掉一个子进程后，主进程可能还活着，但服务已经不正常。

cgroup v2 可以用：

```bash
echo 1 | sudo tee /sys/fs/cgroup/<path>/memory.oom.group
```

含义：

> 当这个 cgroup 发生 OOM 时，把整个 cgroup 当成一个整体处理，而不是只杀某一个进程。

这对容器和 Pod 语义很重要，因为很多时候我们希望“一个关键进程 OOM，整个工作负载失败并重启”。

### 11.6 pids 限制可以防 fork 炸弹

限制进程数量：

```bash
echo 100 | sudo tee /sys/fs/cgroup/<path>/pids.max
```

systemd：

```bash
systemd-run --collect -u limited \
  -p TasksMax=100 \
  ./your_program
```

Docker：

```bash
docker run --pids-limit=100 ...
```

这是防止进程无限 fork 的常见手段。

### 11.7 IO 限制要先确认设备号

cgroup v2 的 `io.max` 通常需要指定设备 major:minor。

查看设备号：

```bash
lsblk
stat -c '%t:%T' /dev/sda
```

限制示例：

```bash
echo "8:0 rbps=10485760 wbps=10485760" | sudo tee /sys/fs/cgroup/<path>/io.max
```

含义：

- `8:0`：块设备号。
- `rbps=10485760`：读带宽 10MB/s。
- `wbps=10485760`：写带宽 10MB/s。

实际生产里 IO 限制受文件系统、块设备、容器存储驱动影响较大，排查会比 CPU/内存复杂。

### 11.8 rootless 容器的路径可能不一样

rootless Docker、Podman 或用户级 systemd 下，cgroup 路径可能在 `user.slice` 下，而不是 `system.slice` 下。

排查时不要硬猜路径，先看：

```bash
PID=$(docker inspect -f '{{.State.Pid}}' <container>)
cat /proc/$PID/cgroup
```

然后根据输出去 `/sys/fs/cgroup` 查。

### 11.9 不要迷信 top 里的单个进程内存

容器或 cgroup 的内存限制是对整个 cgroup 生效，不是只看一个进程。

多进程应用里，需要看整个 cgroup：

```bash
cat /sys/fs/cgroup/<path>/memory.current
```

而不是只看：

```bash
ps aux
top
```

### 11.10 Docker Compose 里也要明确资源限制

Compose 新版本可以写：

```yaml
services:
  app:
    image: my-app
    deploy:
      resources:
        limits:
          cpus: "0.50"
          memory: 512M
```

注意：不同 Compose 版本、是否 Swarm 模式、Docker Engine 版本都会影响字段是否生效。日常要用下面命令验证：

```bash
docker inspect <container> | grep -i -E 'Memory|Cpu'
```

或者进入宿主机查对应 cgroup 文件。

---

## 12. 推荐实践路径

学习 cgroup 可以按这个顺序：

1. 先用 `cgroupfs` 手写一遍 `mkdir`、`cpu.max`、`memory.max`、`cgroup.procs`。
2. 再用 `systemd-run` 做同样的资源限制。
3. 然后查看 `systemd-cgls --all`，把 systemd unit 和 cgroup 目录对应起来。
4. 再跑一个 Docker 容器，观察容器进程在 `/proc/<pid>/cgroup` 中的位置。
5. 最后看 Kubernetes Pod 的 requests/limits 如何落到节点上的 cgroup。

这样理解最稳：

```text
Linux process
  -> cgroupfs
    -> systemd service/scope/slice
      -> Docker/containerd cgroup driver
        -> Kubernetes Pod/container resources
```

---

## 13. 速查清单

### 13.1 cgroupfs

```bash
stat -fc %T /sys/fs/cgroup
cat /sys/fs/cgroup/cgroup.controllers
sudo mkdir /sys/fs/cgroup/demo
echo "50000 100000" | sudo tee /sys/fs/cgroup/demo/cpu.max
echo "100M" | sudo tee /sys/fs/cgroup/demo/memory.max
echo <pid> | sudo tee /sys/fs/cgroup/demo/cgroup.procs
cat /sys/fs/cgroup/demo/memory.events
sudo rmdir /sys/fs/cgroup/demo
```

### 13.2 systemd

```bash
systemd-cgls --all
systemd-cgtop
systemd-run --collect -u demo -p CPUQuota=50% -p MemoryMax=100M ./app
systemctl status demo.service
systemctl reset-failed demo.service
systemctl set-property nginx.service CPUQuota=50% MemoryMax=512M
systemctl show nginx.service -p ControlGroup -p MemoryMax
```

### 13.3 Docker

```bash
docker info | grep -i cgroup
docker run -d --name limited --cpus=0.5 --memory=100m --pids-limit=100 nginx
PID=$(docker inspect -f '{{.State.Pid}}' limited)
cat /proc/$PID/cgroup
```

### 13.4 Kubernetes

```bash
grep -i cgroupDriver /var/lib/kubelet/config.yaml
kubectl describe pod <pod>
kubectl top pod
kubectl top node
```

---

## 14. 总结

这篇博客最重要的价值是让人看到：cgroup 的底层接口其实就是 `/sys/fs/cgroup` 下面的一组目录和文件。

但在日常使用中，需要进一步形成几个判断：

- 学原理时，直接操作 cgroupfs 最直观。
- 在 systemd 发行版上，生产环境优先通过 systemd 管理资源。
- 临时跑受限进程，用 `systemd-run --collect` 很方便。
- 要限制一组服务或容器的总资源，用 systemd slice。
- Docker 和 Kubernetes 的资源限制最终都会落到 cgroup。
- 排查 OOM、CPU throttling、pids 限制时，要看 cgroup 文件，不要只看应用日志。

最终可以记成一句话：

> cgroup 是内核能力，cgroupfs 是原始接口，systemd 是现代 Linux 上常见的统一管理者，Docker 和 Kubernetes 是在它之上的资源管理抽象。
