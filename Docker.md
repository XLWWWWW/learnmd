# Docker

- 概述
- 安装
- 命令
  - 镜像
  - 容器
  - 操作
  - 。。。。
- Docker镜像
- 容器数据卷
- DockerFile
- 网络原理
- IDEA整合Docker
- Dockers compose
- Dockers Swarm
- CI/CD



## 概述

### 为什么出现

环境配置相当麻烦，换一台机器，就要重来一次，费力费时。很多人想到，能不能从根本上解决问题，**软件可以带环境安装**？也就是说，安装的时候，把原始环境一模一样地复制过来。开发人员利用 Docker 可以消除协作编码时“在我的机器上可正常工作”的问题。

总结来说，将应用打成镜像，通过镜像成为运行在Docker容器上面的实例，而 Docker容器在任何操作系统上都是一致的，这就实现了跨平台、跨服务器。只需要一次配置好环境，换到别的机子上就可以一键部署好，大大简化了操作。



### 历史

dotCloud 开始从事云服务计算，**LXC**（Linux容器(Linux Containers，缩写为 LXC)）有关的容器技术。

在此基础上，将容器化技术 命名为 Docker

开源： 发放源代码

虚拟机（virtual machine）就是带环境安装的一种解决方案。实际上来看，就是通过软件模拟一台或者是多台电脑，

但是 虚拟机的缺点：

1    资源占用多               2    冗余步骤多                 3    启动慢



比较了 Docker 和传统虚拟化方式的不同之处：

*传统虚拟机技术是虚拟出一套硬件后，在其上运行一个完整操作系统，在该系统上再运行所需应用进程，荣誉部分较多。

*容器内的应用进程直接运行于宿主的内核，容器内没有自己的内核且也没有进行硬件虚拟。因此容器要比传统虚拟机更为轻便。

*每个容器之间互相隔离，每个容器有自己的文件系统 ，容器之间进程不会相互影响，能区分计算资源。



> Docker基于Go语言，

### 用处

虚拟机技术：

1  资源占用多		2、荣誉步骤多	 3、启动慢

容器化技术

容器不是模拟一个完整的操作系统而是对进程进行隔离。有了容器，就可以将软件运行所需的所有资源打包到一个隔离的容器中。容器与虚拟机不同，不需要捆绑一整套操作系统，只需要软件工作所需的库资源和设置。系统因此而变得高效轻量并保证部署在任何环境中的软件都能始终如一地运行。



- 传统虚拟机，虚拟出一个硬件，运行一个完整的操作系统，然后在这个操作系统上运行和安装软件

- 容器内的应用进程直接运行于宿主的内核，容器内没有自己的内核且也没有进行硬件虚拟。因此容器要比传统虚拟机更为轻便。
- 每个容器之间互相隔离，每个容器有自己的文件系统 ，容器之间进程不会相互影响，能区分计算资源。



> DevOps 开发运维
>
> - 更快的交付和部署
>
>   Docker： 打包镜像发布测试，一键运行
>
> - 更便捷的升级和扩缩容
>
>   使用Docker，应用部署就像搭积木一样
>
> - 更简单的系统运维
>
>   容器化后，开发测试环境都是高度一致的
>
> - 更高效的计算资源利用
>
>   内核级别的虚拟化，可以在一个主机上运行很多的容器实例。



## 安装

<img src="https://ts1.cn.mm.bing.net/th/id/R-C.1774676184117a7651da679303f7f2a5?rik=tD8GxVFzILxUTA&riu=http%3a%2f%2fwww.yiibai.com%2fuploads%2fimages%2f201706%2f0206%2f262150629_86976.png&ehk=b08QZOJOFKWBBfczWiSLKjdCb8XG5Ow2kHoxAUT2F%2fI%3d&risl=&pid=ImgRaw&r=0&sres=1&sresct=1" alt="Docker架构 - Docker教程" style="zoom:67%;" />

- 镜像(image)：

Docker 镜像（Image）就是一个只读的模板。镜像可以用来创建 Docker 容器，一个镜像可以创建很多容器。最终服务运行或者是项目运行就是在容器中的

- 容器(container):

**1 从面向对象角度**

Docker 利用容器（Container）独立运行的一个或一组应用，**应用程序或服务运行在容器里面**，容器就类似于一个虚拟化的运行环境，**容器是用镜像创建的运行实例**。就像是Java中的类和实例对象一样，

镜像是静态的定义，容器是镜像运行时的实体。容器为镜像提供了一个标准的和隔离的运行环境，它可以被启动、开始、停止、删除。每个容器都是相互隔离的、保证安全的平台

**2 从镜像容器角度**

可以把容器看做是一个简易版的 Linux 环境（包括root用户权限、进程空间、用户空间和网络空间等）和运行在其中的应用程序。

- 仓库(reposity)：

仓库（Repository）是集中存放镜像文件的场所

仓库分为公开仓库（Public）和私有仓库（Private）两种形式。

最大的公开仓库是 Docker Hub(https://hub.docker.com/)，

存放了数量庞大的镜像供用户下载。国内的公开仓库包括阿里云 、网易云等。

> 环境准备
>
> 1、linux
>
> 2、CentOS 7
>
> - #### 安装步骤
>
> ```shell
> # 1、卸载旧的Docker
> sudo yum remove docker \
>                   docker-client \
>                   docker-client-latest \
>                   docker-common \
>                   docker-latest \
>                   docker-latest-logrotate \
>                   docker-logrotate \
>                   docker-engine
>                   
> # 2、需要的安装包
> sudo yum install -y yum-utils
> 
> #3、设置镜像的仓库
> sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
> 
> #4、更新yum软件包索引
> yum makecache fast
> 
> # 4、安装  ce-社区版  ee-企业版  
> sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
> 
> # 5、启动
> sudo systemctl start docker
> 
> # 6、
> docker version
> 
> # 7、
> sudo docker run hello-world
> 
> ```
>
> <img src=".\asset\images\image-20240404221735768.png" alt="image-20240404221735768" style="zoom:67%;" />
>
> ```shell
> # 8、查看hello-world镜像
> docker imags #查看所有的镜像
> ```
>
> - #### 卸载docker
>
> ```shell
> # 1、卸载依赖
> sudo yum remove docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras
> 
> # 2、删除资源
> # /var/lib/docker   docker的默认工作路径
> sudo rm -rf /var/lib/docker   
> sudo rm -rf /var/lib/containerd
> ```
>
> - #### 阿里云镜像加速
>
> 1、阿里云 容器服务
>
> 2、找到镜像加速地址
>
> 3、配置阿里云docker加速

### Hello-world 流程

<img src=".\asset\images\image-20240404222602664.png" alt="image-20240404222602664" style="zoom:67%;" />

### 底层原理

**Docker 是怎么工作的**

docker是一个 Client-Server结构的系统，Docker的守护进行一直在后台运行，通过Socket从客户端访问，

DockerServer 接收到Docker-Client的指令 就会执行这个命令

![image-20240404223059583](.\asset\images\image-20240404223059583.png)

**为什么比虚拟机快**

1、Docker有着比虚拟机更少的抽象层

![Docker技术入门-DaoCloud道客博客](https://tse3-mm.cn.bing.net/th/id/OIP-C.dCNtai_xEnaILwJJu3bOmQAAAA?rs=1&pid=ImgDetMain)

2、Docker利用的是宿主机的内核  VM使用的是Guest OS

当新建一个容器时,虚拟机需要重新加载一个操作系统内核。引寻、加载操作系统内核返回等比较费时费资源的过程,当新建一个虚拟机时,虚拟机软件需要加载OS,返回新建过程是分钟级别的。

而docker由于直接利用宿主机的操作系统,则省略了返回过程,因此新建一个docker容器只需要几秒钟。





## Docker常用命令

### 帮助命令

```
docker version			# 显示版本信息
docker info				# 显示系统信息，包括镜像和容器的数量
docker 命令 --help	   # 查看 
```

### 镜像命令

**docker images**   列出所有的镜像

```shell
[root@localhost ~]# docker image ls
REPOSITORY    TAG       IMAGE ID       CREATED         SIZE
hello-world   latest    d2c94e258dcb   11 months ago   13.3kB


## 
REPOSITORY 镜像源
TAG			镜像标签
IMAGE ID	镜像ID
CREATED		创建时间
SIZE		大小

  -a, --all             Show all images (default hides intermediate images)  # 列出全部
      --digests         Show digests
  -f, --filter filter   Filter output based on conditions provided
      --format string   Format output using a custom template:
                        'table':            Print output in table format with column headers (default)
                        'table TEMPLATE':   Print output in table format using the given Go template
                        'json':             Print in JSON format
                        'TEMPLATE':         Print output using the given Go template.
                        Refer to https://docs.docker.com/go/formatting/ for more information about formatting output with templates
      --no-trunc        Don't truncate output
  -q, --quiet           Only show image IDs									# 仅显示ID
```

**docker search xxxx** 搜索镜像

```shell
docker search mysql
NAME                            DESCRIPTION                                     STARS     OFFICIAL
mariadb                         MariaDB Server is a high performing open sou…   5714      [OK]
percona                         Percona Server is a fork of the MySQL relati…   627       [OK]
mysql                           MySQL is a widely used, open-source relation…   14982     [OK]
phpmyadmin                      phpMyAdmin - A web interface for MySQL and M…   963       [OK]

# 可选项
Options:
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print search using a Go template
      --limit int       Max number of search results
      --no-trunc        Don't truncate output
      
[root@localhost ~]# docker search mysql -f=STARS=5000     #搜索STARS大于5000的
NAME      DESCRIPTION                                     STARS     OFFICIAL
mysql     MySQL is a widely used, open-source relation…   14982     [OK]
mariadb   MariaDB Server is a high performing open sou…   5714      [OK]
```

**docker pull xxxx  下载镜像**

```shell
[root@localhost ~]# docker pull mysql
Using default tag: latest					##如果不加tag 默认下载最新版的
latest: Pulling from library/mysql
9a5c778f631f: Pull complete 				## docker 分层下载 imag核心  联合文件系统
ccc451c3fb55: Pull complete 
db534de989c8: Pull complete 
c1a1ab6fb3ea: Pull complete 
d18a374d12e6: Pull complete 
2d9f4c3e8c03: Pull complete 
4c79cbebfe62: Pull complete 
b3549fdd6799: Pull complete 
c08846a4ab7a: Pull complete 
084bd453daf0: Pull complete 
Digest: sha256:4552fcc5d3cdb8cdee76ee25cce28bf60b0eb3ce93d25ba3bfff7a66bfdcdee8			##签名
Status: Downloaded newer image for mysql:latest  
docker.io/library/mysql:latest							## 真实地址
#docker pull mysql 等价于 docker pull docker.io/library/mysql:latest

# 指定版本下载
[root@localhost ~]# docker pull mysql:5.7
```

**docker rmi 删除镜像**

```shell
## 删除指定镜像
[root@localhost ~]# docker rmi -f 镜像id

## 先通过 docker images -q获取所有的imageid 接着一个一个递归移除
[root@localhost ~]# docker rmi -f $(docker images -q)

## 删除多个镜像
[root@localhost ~]# docker rmi -f 镜像id 镜像id 镜像id 镜像id
```

### 容器命令

有了镜像才能够创建容器

```shell
[root@localhost ~]# docker pull centos
```

**新建容器并启动**

```shell
docker run [可选参数] image

# 参数说明
--name="Name" 		#为运行的容器起名字 以区分容器
-d					#以后台方式运行
-it					#使用交互方式运行
-p 					#指定容器的端口
		-p ip:主机端口:容器端口
		-p 主机端口:容器端口
		-p 容器端口
		
-P					#随机指定端口

## 启动并进入容器		主机名就是镜像名字
[root@localhost ~]# docker run -it centos /bin/bash
[root@9d3498e8b2e8 /]# ls	
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
## 退出容器
[root@9d3498e8b2e8 /]# exit
exit
```

**列出所有运行中的容器**

```shell
# docker ps
			## 列出所有运行中的容器
-a			## 列出所有运行中的容器 + 历史运行的容器
-n=数字 	   ## 列出最近创建的几个容器 
-q			## 仅显示容器ID

[root@localhost ~]# docker ps				## 列出所有运行中的容器
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
[root@localhost ~]# docker ps -a			## 列出所有曾经运行中的容器
CONTAINER ID   IMAGE          COMMAND       CREATED             STATUS                         PORTS     NAMES
9d3498e8b2e8   centos         "/bin/bash"   2 minutes ago       Exited (0) 45 seconds ago                ecstatic_goldwasser
736a18534ddf   d2c94e258dcb   "/hello"      About an hour ago   Exited (0) About an hour ago             admiring_bohr
```

**退出容器**

```shell
exit		# 直接退出并停止容器
Ctrl+p+q	# 退出容器 但是不关闭
```

**删除容器**

```shell
docker rm 容器id					#删除指定容器 运行中的容器不能直接删除
docker rm -f 容器id				#强制删除指定容器
docker rm -f $(docker ps -aq)	 #和上面的逻辑相同，删除全部容器
docker ps -a -q|xargs docker rm  #使用管道删除容器
```

**启动和停止容器的操作**

```
docker start 容器id		#启动
docker stop 容器id		#停止
docker restart 容器id		#重启
docker kill 容器id		#强制停止
```

### 常用其他命令

**后台启动容器**

```shell
## docker run -d 镜像名
[root@localhost ~]# docker run -d centos
6b2937ec29347ff846bb9aea0113e37ad0cbeae29b368c2d01ced4722f01196a

## docker容器后台运行，必须要有前台进程，不然会停止 docker发现没有前台应用 就会自动停止

```

**查看日志**

```shell
docker logs [Options:] 容器id

Options:
      --details        Show extra details provided to logs
  -f, --follow         Follow log output
      --since string   Show logs since timestamp (e.g. "2013-01-02T13:23:37Z") or relative (e.g. "42m" for 42 minutes)
  -n, --tail string    Number of lines to show from the end of the logs (default "all")
  -t, --timestamps     Show timestamps
      --until string   Show logs before a timestamp (e.g. "2013-01-02T13:23:37Z") or relative (e.g. "42m" for 42 minutes)
      
      
[root@localhost ~]# docker run -d centos /bin/sh -c "while true;do echo test;sleep 1;done"  ## 让centos后台启动并一直打印test
[root@localhost ~]# docker logs -f -t --details eb208932e854								## 查看日志信息
```

**查看容器中的进程信息**

```shell
docker top 容器ID 
```

查看镜像源数据

```shell
docker inspect 容器ID

[root@localhost ~]# docker inspect eb208932e854
[
    {
        "Id": "eb208932e8542bd7b77d0f4b5552c09ecf3db71480eff780864c85f64ebb5c4f",
        "Created": "2024-04-04T15:53:48.667104881Z",
        "Path": "/bin/sh",
        "Args": [
            "-c",
            "while true;do echo test;sleep 1;done"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 11066,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2024-04-04T15:53:48.777314873Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:5d0da3dc976460b72c77d94c8a1ad043720b0416bfc16c52c45d4847e53fadb6",
        "ResolvConfPath": "/var/lib/docker/containers/eb208932e8542bd7b77d0f4b5552c09ecf3db71480eff780864c85f64ebb5c4f/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/eb208932e8542bd7b77d0f4b5552c09ecf3db71480eff780864c85f64ebb5c4f/hostname",
        "HostsPath": "/var/lib/docker/containers/eb208932e8542bd7b77d0f4b5552c09ecf3db71480eff780864c85f64ebb5c4f/hosts",
        "LogPath": "/var/lib/docker/containers/eb208932e8542bd7b77d0f4b5552c09ecf3db71480eff780864c85f64ebb5c4f/eb208932e8542bd7b77d0f4b5552c09ecf3db71480eff780864c85f64ebb5c4f-json.log",
        "Name": "/tender_burnell",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "ConsoleSize": [
                40,
                137
            ],
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": [],
            "BlkioDeviceWriteBps": [],
            "BlkioDeviceReadIOps": [],
            "BlkioDeviceWriteIOps": [],
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": [],
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware",
                "/sys/devices/virtual/powercap"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/72d76ca61be6e85eab7f5881adde04ab7a46427481156c7ccd1df8118e7b5888-init/diff:/var/lib/docker/overlay2/638c3e6ef42affdd2bb01a7b27d0d2d6b7634bf0f7ef715a68ad2e1967cff73d/diff",
                "MergedDir": "/var/lib/docker/overlay2/72d76ca61be6e85eab7f5881adde04ab7a46427481156c7ccd1df8118e7b5888/merged",
                "UpperDir": "/var/lib/docker/overlay2/72d76ca61be6e85eab7f5881adde04ab7a46427481156c7ccd1df8118e7b5888/diff",
                "WorkDir": "/var/lib/docker/overlay2/72d76ca61be6e85eab7f5881adde04ab7a46427481156c7ccd1df8118e7b5888/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "eb208932e854",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "while true;do echo test;sleep 1;done"
            ],
            "Image": "centos",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "org.label-schema.build-date": "20210915",
                "org.label-schema.license": "GPLv2",
                "org.label-schema.name": "CentOS Base Image",
                "org.label-schema.schema-version": "1.0",
                "org.label-schema.vendor": "CentOS"
            }
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "b6514e8d24d26a840b2cc293b324e84acaf2f74efb8f63fac87db7b352717ce5",
            "SandboxKey": "/var/run/docker/netns/b6514e8d24d2",
            "Ports": {},
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "ca0b9b26d156dd952382a35dc9d31da99e1b5978b82b78ff62e6070568a4daf1",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "MacAddress": "02:42:ac:11:00:02",
                    "NetworkID": "5af9980126cf4203692b0798261f68c69961517b23550a2645306b15a11393f7",
                    "EndpointID": "ca0b9b26d156dd952382a35dc9d31da99e1b5978b82b78ff62e6070568a4daf1",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "DriverOpts": null,
                    "DNSNames": null
                }
            }
        }
    }
]
```

**进入当前正在运行的容器**

```shell
#通常容器都是后台运行的  需要进入容器 修改一些配置

# 进入容器后 开一个新的终端，可以在里面操作
docker extc -it 容器id bash/shell

# eg
[root@localhost ~]# docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS     NAMES
eb208932e854   centos    "/bin/sh -c 'while t…"   10 minutes ago   Up 10 minutes             tender_burnell
[root@localhost ~]# docker exec -it eb208932e854
"docker exec" requires at least 2 arguments.
See 'docker exec --help'.

Usage:  docker exec [OPTIONS] CONTAINER COMMAND [ARG...]

Execute a command in a running container
[root@localhost ~]# docker exec -it eb208932e854 /bin/bash
[root@eb208932e854 /]# ps -ef
UID         PID   PPID  C STIME TTY          TIME CMD
root          1      0  0 15:53 ?        00:00:00 /bin/sh -c while true;do echo test;sleep 1;done
root        695      0  0 16:05 pts/0    00:00:00 /bin/bash
root        714      1  0 16:05 ?        00:00:00 /usr/bin/coreutils --coreutils-prog-shebang=sleep /usr/bin/sleep 1
root        715    695  0 16:05 pts/0    00:00:00 ps -ef



## 第二中方法  进入容器正在执行的终端
docker attach 容器id			
```

**从容器内拷贝文件到主机**

容器在数据就在 就能够拷贝出来 和运行不运行没有关系

```shell
docker cp 容器id:容器内路径  目的主机路径

# eg
[root@localhost home]# docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED              STATUS              PORTS     NAMES
1fadf52d86a4   centos    "/bin/bash"   About a minute ago   Up About a minute             priceless_mcnulty
[root@localhost home]# docker attach 1fadf52d86a4
[root@1fadf52d86a4 /]# cd /home
[root@1fadf52d86a4 home]# touch test.java
[root@1fadf52d86a4 home]# ls
test.java
[root@1fadf52d86a4 home]# exit
exit
[root@localhost home]# docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
[root@localhost home]# docker cp 1fadf52d86a4:/home/test.java /home
Successfully copied 1.54kB to /home
[root@localhost home]# ls
test.java  xlwwwww  xlwwwwwtest
[root@localhost home]# 

```

拷贝是一个手动过程，后续能够使用 -v 卷挂载 自动同步



### 命令小结

![img](https://cdn.nlark.com/yuque/0/2022/png/27791237/1658128808416-bc53f29a-c6b0-4b5f-966c-b6cb83a9e66b.png)

```shell
attach    Attach to a running container                 # 当前 shell 下 attach 连接指定运行镜像
build     Build an image from a Dockerfile              # 通过 Dockerfile 定制镜像
commit    Create a new image from a container changes   # 提交当前容器为新的镜像
cp        Copy files/folders from the containers filesystem to the host path   #从容器中拷贝指定文件或者目录到宿主机中
create    Create a new container                        # 创建一个新的容器，同 run，但不启动容器
diff      Inspect changes on a container's filesystem   # 查看 docker 容器变化
events    Get real time events from the server          # 从 docker 服务获取容器实时事件
exec      Run a command in an existing container        # 在已存在的容器上运行命令
export    Stream the contents of a container as a tar archive   # 导出容器的内容流作为一个 tar 归档文件[对应 import ]
history   Show the history of an image                  # 展示一个镜像形成历史
images    List images                                   # 列出系统当前镜像
import    Create a new filesystem image from the contents of a tarball # 从tar包中的内容创建一个新的文件系统映像[对应export]
info      Display system-wide information               # 显示系统相关信息
inspect   Return low-level information on a container   # 查看容器详细信息
kill      Kill a running container                      # kill 指定 docker 容器
load      Load an image from a tar archive              # 从一个 tar 包中加载一个镜像[对应 save]
login     Register or Login to the docker registry server    # 注册或者登陆一个 docker 源服务器
logout    Log out from a Docker registry server          # 从当前 Docker registry 退出
logs      Fetch the logs of a container                 # 输出当前容器日志信息
port      Lookup the public-facing port which is NAT-ed to PRIVATE_PORT    # 查看映射端口对应的容器内部源端口
pause     Pause all processes within a container        # 暂停容器
ps        List containers                               # 列出容器列表
pull      Pull an image or a repository from the docker registry server   # 从docker镜像源服务器拉取指定镜像或者库镜像
push      Push an image or a repository to the docker registry server    # 推送指定镜像或者库镜像至docker源服务器
restart   Restart a running container                   # 重启运行的容器
rm        Remove one or more containers                 # 移除一个或者多个容器
rmi       Remove one or more images       # 移除一个或多个镜像[无容器使用该镜像才可删除，否则需删除相关容器才可继续或 -f 强制删除]
run       Run a command in a new container              # 创建一个新的容器并运行一个命令
save      Save an image to a tar archive                # 保存一个镜像为一个 tar 包[对应 load]
search    Search for an image on the Docker Hub         # 在 docker hub 中搜索镜像
start     Start a stopped containers                    # 启动容器
stop      Stop a running containers                     # 停止容器
tag       Tag an image into a repository                # 给源中镜像打标签
top       Lookup the running processes of a container   # 查看容器中运行的进程信息
unpause   Unpause a paused container                    # 取消暂停容器
version   Show the docker version information           # 查看 docker 版本号
wait      Block until a container stops, then print its exit code   # 截取容器停止时的退出状态值
```





### 实践

- 设置Nginx服务器

```shell
# 1、下载镜像
# 2、
[root@192 ~]# docker images
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
nginx        latest    92b11f67642b   7 weeks ago   187MB
centos       latest    5d0da3dc9764   2 years ago   231MB

# 测试

# -d 后台运行  --name 起别名  
# -p 宿主机端口:容器内部端口 实际上就是将容器的内部端口做了一个映射，能够通过宿主机的3344端口访问内部正在运行的80端口
[root@192 ~]# docker run -d --name nginx01 -p 3344:80 nginx
3fa1e99b1ca4d5a1966bf234a1801a39537f1895068e9de3a02ac76ad9f56541
[root@192 ~]# docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                                   NAMES
3fa1e99b1ca4   nginx     "/docker-entrypoint.…"   3 seconds ago   Up 2 seconds   0.0.0.0:3344->80/tcp, :::3344->80/tcp   nginx01
[root@192 ~]# curl localhost:3344
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>

# 进入docker 容器
[root@192 ~]# docker exec -it 3fa1e99b1ca4 /bin/bash
root@3fa1e99b1ca4:/# whereis nginx
nginx: /usr/sbin/nginx /usr/lib/nginx /etc/nginx /usr/share/nginx
root@3fa1e99b1ca4:/# cd /etc/nginx/
root@3fa1e99b1ca4:/etc/nginx# ls
conf.d  fastcgi_params  mime.types  modules  nginx.conf  scgi_params  uwsgi_params
root@3fa1e99b1ca4:/etc/nginx# 
```

但是对docker内部的文件进行修改需要进入容器，后续能够通过数据卷 在外部完成容器内部的自动修改



- 配置es+kibana

  > es 需要暴露的端口很多
  >
  > es 十分消耗内存
  >
  > es 需要挂载外部文件

  

```shell
# 下载并启动 elasticsearch
[root@192 ~]# docker run -d --name elasticsearch --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2

# 查看cpu状态
docker stats

# 测试 es是否成功
[root@192 ~]# curl localhost:9200
{
  "name" : "9ef1ca0e5092",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "bX_b309bSLCwgFvuJAMxbA",
  "version" : {
    "number" : "7.6.2",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "ef48eb35cf30adf4db14086e8aabd07ef6fb113f",
    "build_date" : "2020-03-26T06:34:37.794943Z",
    "build_snapshot" : false,
    "lucene_version" : "8.4.0",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}

# 增加内存限制  修改配置文件  -e 环境配置修改
[root@192 ~]# docker run -d --name elasticsearch02 --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m-Xms512m"  elasticsearch:7.6.2

```

使用kibana连接elasticsearch

![image-20240405141907105](.\asset\images\image-20240405141907105.png)





## Docker可视化

- portainer

  什么是portainer

  Docker图形化管理工具 提供一个后台面板供我们操作

  

  下载

  ```
  docker run -d -p 8088:9000 --restart=always -v /var/run/docker.sock:/var/run/docker.sock  --privileged=true portainer/portainer
  ```

  访问测试

  登录 http:/ip:8088 就能以可视化的方法操作docker

- Rancher(CI/CD)



## Docker镜像讲解

### 镜像是什么

是一种轻量级、可执行的独立软件包，它包含运行某个软件所需的所有内容，我们把应用程序和配置依赖打包好形成一个可交付的运行环境(包括代码、运行时需要的库、环境变量和配置文件等)，这个打包好的运行环境就是image镜像文件。

只有通过这个镜像文件才能生成Docker容器实例(类似Java中new出来一个对象)。

- 远程仓库下载
- 拷贝
- 自己制作

### Docker 镜像加载原理

> UnionFS（联合文件系统）

UnionFS（联合文件系统）：Union文件系统（UnionFS）是一种分层、轻量级并且高性能的文件系统，它支持对文件系统的修改作为一次提交来一层层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统下(unite several directories into a single virtual filesystem)。Union 文件系统是 Docker 镜像的基础。镜像可以通过分层来进行继承，基于基础镜像（没有父镜像），可以制作各种具体的应用镜像。

特性：一次同时加载多个文件系统，但从外面看起来，只能看到一个文件系统，联合加载会把各层文件系统叠加起来，这样最终的文件系统会包含所有底层的文件和目录

多个镜像有相同的基础层，就不用重复下载



> Docker 镜像加载原理

docker的镜像实际上由一层一层的文件系统组成，这种层级的文件系统UnionFS。

bootfs(boot file system)主要包含bootloader和kernel, bootloader主要是引导加载kernel, Linux刚启动时会加载bootfs文件系统，在Docker镜像的最底层是引导文件系统bootfs。这一层与我们典型的Linux/Unix系统是一样的，包含boot加载器和内核。当boot加载完成之后整个内核就都在内存中了，此时内存的使用权已由bootfs转交给内核，此时系统也会卸载bootfs。

rootfs (root file system) ，在bootfs之上。包含的就是典型 Linux 系统中的 /dev, /proc, /bin, /etc 等标准目录和文件。rootfs就是各种不同的操作系统发行版，比如Ubuntu，Centos等等。 

![image.png](https://cdn.nlark.com/yuque/0/2022/png/27791237/1658139701434-5322591f-b716-4232-8f3c-9420ba706072.png?x-oss-process=image%2Fformat%2Cwebp%2Fresize%2Cw_836%2Climit_0)

对于一个精简的OS，rootfs可以很小，只需要包括最基本的命令、工具和程序库就可以了，因为底层直接用Host的kernel，自己只需要提供 rootfs 就行了。由此可见对于不同的linux发行版, bootfs基本是一致的, rootfs会有差别, 因此不同的发行版可以公用bootfs。



### 分层理解

![image-20240405151942415](.\asset\images\image-20240405151942415.png)

层级的结构 方便多个镜像复用共同的基层



> 特点

Docker镜像都是只读的，当容器启动时，一个新的可写层被加载到镜像的顶部。

这一层就是我们通常所说的容器层，容器之下的都叫镜像层。

通过再次打包经过处理的容器层和基础的镜像层，能够提交新的镜像。



### 提交镜像

commit

```shell
# 提交容器成为一个新的副本
docker commit -m=提交的描述信息 -a="作者" 容器ID 目标镜像名:[Tag]

```

```shell
# 启动一个默认的tomcat 该官方镜像默认webapps文件夹下没有文件
[root@localhost ~]# docker run -it -p 8080:8080 tomcat 

# 拷贝进去基本文件
[root@localhost ~]# docker exec -it 6d8c12031532 /bin/bash
root@6d8c12031532:/usr/local/tomcat# cd webapps
root@6d8c12031532:/usr/local/tomcat/webapps# ls
root@6d8c12031532:/usr/local/tomcat/webapps# cd ..
root@6d8c12031532:/usr/local/tomcat# cp -r webapps.dist/* webapps
root@6d8c12031532:/usr/local/tomcat# cd webapps
root@6d8c12031532:/usr/local/tomcat/webapps# ls
docs  examples  host-manager  manager  ROOT

# 退出镜像后，将该镜像打包提交
[root@localhost ~]# docker commit -a="xlwwwww" -m="add webapps" 6d8c12031532 tomcat02:1.0
```

![image-20240405154911566](.\asset\images\image-20240405154911566.png)



## 容器数据卷

Docker 将应用和环境打包成一个镜像

但是数据都是存储在容器中的，容器删除，数据就会丢失 ------> 要求数据应该可以持久化

并且 对于容器中数据的修改 需要进入容器才能处理 

因此，希望容器之间能够有一个数据共享的技术！

在Docker容器中产生的数据能够同步到本地  =======》将容器内的目录 挂载到宿主机上

![image-20240405163005917](.\asset\images\image-20240405163005917.png)

总结，使用挂载数据卷实现 容器的持久化和同步操作  容器间也是可以进行数据共享的。



### 使用数据卷

#### 直接使用命令挂载  -v

```shell
docker run [options] -v 主机目录地址:容器内目录 容器名字
```

使用 docker inspect 容器ID 查看容器信息 看到容d器映射成功

```shell
[root@localhost ceshi]# docker ps
CONTAINER ID   IMAGE                 COMMAND        CREATED         STATUS         PORTS                                                           NAMES
41b755fe0acb   centos                "/bin/bash"    2 minutes ago   Up 2 minutes                                                                   recursing_blackburn
5a0af8ac2970   portainer/portainer   "/portainer"   2 hours ago     Up 2 hours     8000/tcp, 9443/tcp, 0.0.0.0:8088->9000/tcp, :::8088->9000/tcp   inspiring_shtern
[root@localhost ceshi]# docker inspect 41b755fe0acb
```

![image-20240405163901490](.\asset\images\image-20240405163901490.png)



**两个映射的文件夹的数据是会相互同步的**

![image-20240405165221145](.\asset\images\image-20240405165221145.png)

运行的容器删除后，无论强制与否，挂载的本地数据卷中的内容依旧没有丢失

```
# 删除容器
docker rm -f 容器名/容器ID
```



##### 测试 安装MySQL

mysql所有的数据都保存在data目录下

```shell
# 1、获取镜像
[root@localhost ~]# docker pull mysql

# 2、启动镜像 并挂载文件卷  安装mysql是需要配置密码的
## 官方测试代码  docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag

-d 后台运行 
-p 本地端口:容器端口
-v 数据卷挂载 通过两个 -v 挂载两个目录
--name 起别名

[root@localhost ~]# docker run -d -p 3301:3306  -v /home/mysql/conf:/etc/mysql/conf.d -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=13141567 --name mysql01 mysql


# 3、启动成功 就能远程连接容器内正在运行的Mysql服务器了 ip 是宿主机ip 端口是映射出的端口
mysql -h ip -P port -u root -p123456 
```

容器彻底删除后，同步到宿主机的文件依旧没有丢失

![image-20240405172113846](.\asset\images\image-20240405172113846.png)



#### 具名和匿名挂载

##### 匿名挂载

```shell
# -v 参数后不指定主机地址
docker run -d -p 3344:80 --name nginx01 -v /ect/nginx nginx

# 使用docker volumn  查看卷信息
[root@localhost data]# docker volume --help
Usage:  docker volume COMMAND
Manage volumes
Commands:
  create      Create a volume
  inspect     Display detailed information on one or more volumes
  ls          List volumes
  prune       Remove unused local volumes
  rm          Remove one or more volumes
  
[root@localhost data]# docker volume ls
DRIVER    VOLUME NAME
local     3e6bd92e6faa834e13b33a9c4bd73e00a409332f7afe328f766cf5371024ddf3         ## 这些就是匿名映射的信息 没有写文件外的信息
local     4f9f4872397210f7104c579523a16fbf3d81b10edc47f3185574a6a67fc5fb82
local     5dc3713a2f39f93e501e81b268315f8b93d77470f2226e976983132685717633
local     7a100d87fd344b71115c380057a7f87ae301e430d0a9a820a7aeda9cbc6b531d
local     8e9f09f318b4d5406b474c4b758d613bbf0d5ee35f5d6104a1e61ae2c2c53190
local     31c4a3ac20f3f93ed6d47d250cc45ac2be26e0d194b87937a7d6447a9a1944da
local     31e326c24cb8b01046ec889c11c20233dc3100beceb8663febea570c350d0bbe
local     36928552557be144b88c31ec781be259794463df6dbb270e27ecf97104eeecd3
local     f5d792c1f6c44e0131ae57d8ab6d1aeb24cb2ef5bd721977caf04e4e1dfa317d
local     f932f14d08279865da8b66cff494f7a824e90514b29f643c229f516e9882ed54

## 查看
[root@localhost data]# 
```

##### 具名挂载

-v 宿主机地址:容器内地址

```shell
[root@localhost data]# docker run -d -P --name nginx03 -v juming-nginx:/etc/nginx nginx
8ad3406759eabd17b5294c749ab7cc6a846a601184d76b3c3b00dc9ca7f1da82
[root@localhost data]# docker volume ls
DRIVER    VOLUME NAME
local     juming-nginx

# 因为在这里映射时没指定根目录 只制定了 juming-nginx 文件夹 因此 可以使用 docker volume inspect 查看详细信息
## 查看详细信息
## 所有没指定 根目录的映射 都在/var/lib/docker/volumes文件夹下
[root@localhost data]# docker volume inspect juming-nginx
[
    {
        "CreatedAt": "2024-04-05T17:40:49+08:00",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/juming-nginx/_data",
        "Name": "juming-nginx",
        "Options": null,
        "Scope": "local"
    }
] 
```

<img src=".\asset\images\image-20240405174211284.png" alt="image-20240405174211284" style="zoom:67%;" />

> 进到这个地址可以看到 之前的匿名映射的文件夹也在这里 ，这就是映射的默认文件夹
>
> ![image-20240405175111936](.\asset\images\image-20240405175111936.png)



**因此，通过具名挂载能够高效的拿到数据**



##### 如何区分具名挂载还是匿名挂载，还是指定路径挂载

-v 容器内路径		#匿名挂载

-v 卷名称:容器内路径   #具名挂载

-v 宿主机路径:容器内路径	#指定路径挂载



##### 拓展

```shell
# 通过 -v 容器内路径 ro rw 改变读写权限
ro:readonly     rw: read&write

# 一旦设定容器权限
docker run -d -P --name nginx03 -v juming-nginx:/etc/nginx:ro nginx			##ro 容器没有写权限，只能从宿主机改变
docker run -d -P --name nginx03 -v juming-nginx:/etc/nginx:rw nginx			##rw 可读可写 默认权限
```



#### 使用DockerFile挂载

DockerFile就是Docker镜像的构建文件

commit 是在原有镜像的基础上构建

DockerFile 就是一段命令脚本 镜像是一层一层的 脚本就是一个一个的命令 每个命令都是一层

```shell
FROM centos

VOLUME ["volume01","volume02"]

CMD echo "----end----"
CMD /bin/bash
```

创建镜像

```shell
# 1、创建dockerfile1 文件
# 文件内容如下
FROM centos

VOLUME ["volume01","volume02"]		# 挂载数据卷 这里什么都没写 只写了容器内的目录 一定是匿名挂载

CMD echo "----end----"
CMD /bin/bash

# 2、使用dockerfile 创建镜像
docker build -f dockerfile1 -t kuangshen-centos:1.0 .
```

<img src=".\asset\images\image-20240405181656976.png" alt="image-20240405181656976" style="zoom:67%;" />



```shell
# 3、运行一下
docker run -it f8b442ef8a08 /bin/bash
```

<img src=".\asset\images\image-20240405182116158.png" alt="image-20240405182116158" style="zoom: 80%;" />

```shell
# 这个卷和外部的地址一定是关联的 查看外部挂载路径
[root@localhost ~]# docker ps
CONTAINER ID   IMAGE          COMMAND       CREATED         STATUS         PORTS     NAMES
1f87bcc5e9ad   f8b442ef8a08   "/bin/bash"   4 minutes ago   Up 4 minutes             objective_kowalevski
[root@localhost ~]# docker inspect 1f87bcc5e9ad
```

<img src=".\asset\images\image-20240405182600031.png" alt="image-20240405182600031" style="zoom:67%;" />



#### 数据卷容器

多个容器间同步数据

![image-20240405192935769](.\asset\images\image-20240405192935769.png)



```shell
## 启动三个容器

#[root@localhost ~]# docker run -it --name docker01 centos

#[root@localhost ~]# docker run -it --name docker02 --volumes-from docker01 centos
 
#在docker01中创建的文件同步到了docker02中 两个容器间的数据能够同步 docker01称为数据卷容器
```

<img src=".\asset\images\image-20240405211014628.png" alt="image-20240405211014628" style="zoom:67%;" />

> 同理，还能在docker01的数据卷上继续挂载 之间的数据是互通的
>
> ![image-20240405211721969](.\asset\images\image-20240405211721969.png)
>
> 此时，docker02 docker03都继承于docker01
>
> 那么docker01停止后会发生什么事情
>
> docker01，docker02和docker03中的数据还在，并且docker02和docker03还能够继续互通数据
>
> ![image-20240405213330836](.\asset\images\image-20240405213330836.png)
>
> 各数据卷之间的是拷贝的概念，一个容器删除后，其余的数据库应该是不受影响的



实践：多个mysql实现数据共享

```shell
[root@localhost ~]# docker run -d -p 3301:3306  -v /etc/mysql/conf.d -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=13141567 --name mysql01 mysql


[root@localhost ~]# docker run -d -p 3302:3306  -v /etc/mysql/conf.d -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=13141567 --name mysql02 mysql --volumes--from mysql01

## 这就能够实现了两个容器数据相通
```



结论：

容器之间配置信息的传递，**数据卷容器的生命周期一直持续到没有容器使用为止**

一旦你持久化到了本地，这个时候，本地的数据是不会删除的

**这里有一个疑问，我能不能在将文件挂载到本地的基础上，进一步将其创建为数据卷容器呢，应该是可以的，在第一个文件`-v`到本地后，接着后面的容器 `--volumes-from docker01`**

## DockerFile

### 概述

用来构建docker镜像的文件! 命令参数脚本！

1、编写 dockerfile 文件

2、docker build 构建成为一个镜像

3、docker run 运行镜像

4、docker push 发布镜像

### DockerFile构建过程

#### 基础知识：

1、每个保留关键字(指令)都必须是大写字母

2、执行顺序是从上到下的

3、#表示注释

4、每一个指令都会提交一个新的镜像层，并提交。

![Docker最新版19.03 详细教程_docker 19.03-CSDN博客](https://img-blog.csdnimg.cn/20200630141015221.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM3MjQyNTIw,size_16,color_FFFFFF,t_70)



DockerFile ：构建文件，定义一切的步骤 源代码

Docker Image：通过DockerFile构建生成的文件，最终发布和运行的产品 

Docker Container: 容器就是镜像运行提供服务的



### DockerFile指令



<img src="https://img-blog.csdnimg.cn/img_convert/f23693c9b140f728f56d9785360076f6.png" alt="Dockerfile 详解，看这一篇就够了_dockerfile文件详解-CSDN博客" style="zoom:67%;" />

1、FROM 			基础镜像 ，一切从这里进行构建

2、MAINTAINER  	  镜像作者姓名+邮箱

3、RUN  			docker镜像运行的指令

4、ADD			  加入的文件，不能只有主系统吧，你可能要在这里面安装什么 需要先把文件添加进来

5、WORKDIR		镜像的工作目录

6、VOLUME		  事先指定的挂载的匿名卷，这样在运行时如果用户不指定挂载，其应用也可以正常运行

​					VOLUME指令只是起到了声明了容器中的目录作为匿名卷，但是并没有将匿名卷绑定到宿主机指定目录的功能。

​					主要用于避免用户忘记指定-v的时候导致的数据丢失

>1、如果-v和volume指定了不同的位置，会发生什么事呢？
>会以-v设定的目录为准，其实volume指令的设定的目的就是为了避免用户忘记指定-v的时候导致的数据丢失，那么如果用户指定了-v，自然而然就不需要volume指定的位置了。
>
>2、那么为什么dockerfile中不提供一个能够映射为主机目录:容器目录这样的指令呢？
>
>其实这样的设计是有道理的，如果在dockerfile中指定了主机目录，这样dockerfile就不具备了可移植性了，毕竟每个人所需要映射的目录可能是不同的，那么最好的办法就是把这个权利交给每个运行这个dockerfile的人，所以才会有 run -v 主机目录:容器目录 这样的指令。
>
>[一文说清楚Dockerfile 中VOLUME到底有什么用？ - 技术颜良 - 博客园 (cnblogs.com)](https://www.cnblogs.com/cheyunhua/p/16923339.html)

7、EXPOSE		    指定对外的端口

8、CMD			 指定容器启动要运行的命令 CMD只有最后一个会生效 可被替代

9、ENTRYPOINT  	指定容器启动要运行的命令，能够追加命令

10、ONBUILD	      构建一个被继承 DockerFile 那么这个时候就会运行 ONBUILD的指令。触发指令

11、COPY                     类似ADD 将文件拷贝到镜像中

12、ENV			构建的时候设置环境变量



### 构建自己的CentOS

DockerHub 中 99%的镜像都是从基础镜像 scratch过来的

1、编写配置文件

```shell
FROM centos:7
MAINTAINER xlwwwwww<xlwwwwwj@163.com>

ENV MYPATH /usr/local
WORKDIR $MYPATH

RUN yum -y install vim
RUN yum -y install net-tools

EXPOSE 80

CMD echo $MYPATH
CMD echo “----end----"
CMD /bin/bash
```

2、构建镜像

```shell
-f 文件路径   -t 镜像名:版本号
[root@localhost dockerfile]# docker build -f mydockerfile -t mycentos:1.0 .
```

3、测试运行

```shell
# 运行
[root@localhost dockerfile]# docker run -it mycentos:1.0   

# 查看当前目录 修改为了/usr/local
[root@f9baac0e42fd local]# pwd

# 并且还安装了 vim 和 net-tools 能够查询本机网络信息
```

![image-20240405232557805](.\asset\images\image-20240405232557805.png)

4、docker history 查看镜像变更历史

```shell
[root@localhost dockerfile]# docker history mycentos:1.0
IMAGE          CREATED             CREATED BY                                      SIZE      COMMENT
a69caa4c5527   About an hour ago   CMD ["/bin/sh" "-c" "/bin/bash"]                0B        buildkit.dockerfile.v0
<missing>      About an hour ago   CMD ["/bin/sh" "-c" "echo “----end----\""]      0B        buildkit.dockerfile.v0
<missing>      About an hour ago   CMD ["/bin/sh" "-c" "echo $MYPATH"]             0B        buildkit.dockerfile.v0
<missing>      About an hour ago   EXPOSE map[80/tcp:{}]                           0B        buildkit.dockerfile.v0
<missing>      About an hour ago   RUN /bin/sh -c yum -y install net-tools # bu…   212MB     buildkit.dockerfile.v0
<missing>      About an hour ago   RUN /bin/sh -c yum -y install vim # buildkit    301MB     buildkit.dockerfile.v0
<missing>      About an hour ago   WORKDIR /usr/local                              0B        buildkit.dockerfile.v0
<missing>      About an hour ago   ENV MYPATH=/usr/local                           0B        buildkit.dockerfile.v0
<missing>      About an hour ago   MAINTAINER xlwwwwww<xlwwwwwj@163.com>           0B        buildkit.dockerfile.v0
<missing>      2 years ago         /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B        
<missing>      2 years ago         /bin/sh -c #(nop)  LABEL org.label-schema.sc…   0B        
<missing>      2 years ago         /bin/sh -c #(nop) ADD file:b3ebbe8bd304723d4…   204MB 
```



> CMD 和 ENTRYPOINT区别
>
> ```shell
> CMD			 指定容器启动要运行的命令 CMD只有最后一个会生效 可被替代
> ENTRYPOINT  	指定容器启动要运行的命令，能够追加命令
> ```
>
> - 构建cmdtest
>
> ```shell
> # 构建DockerFile文件
> [root@localhost dockerfile]# vi dockerfile2
> [root@localhost dockerfile]# cat dockerfile2 
> FROM centos
> CMD ["ls","-a"]				#docker run的时候就会执行这个命令
> 
> # 编译  镜像名字为cmdtest:1.0 .
> [root@localhost dockerfile]# docker build -f dockerfile2 -t cmdtest:1.0 .
> 
> [root@localhost dockerfile]# docker run cmdtest:1.0
> .
> ..
> .dockerenv
> bin
> dev
> etc
> home
> lib
> lib64
> lost+found
> media
> mnt
> opt
> proc
> root
> run
> sbin
> srv
> sys
> tmp
> usr
> var
> [root@localhost dockerfile]# 
> 
> # 想追加一个命令 -l 希望执行 ls -a -l 
> # 在CMD情况下  -l命令替换了CMD ["ls","-a"] 导致报错
> [root@localhost dockerfile]# docker run cmdtest:1.0 -l
> docker: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "-l": executable file not found in $PATH: unknown.
> 
> # 想要完整的使用 需要这么写
> [root@localhost dockerfile]# docker run cmdtest:1.0 ls -a -l
> ```
>
> ![image-20240405234238637](.\asset\images\image-20240405234238637.png)
>
> 
>
> - 构建ENTRYPOINT
>
> ```shell
> [root@localhost dockerfile]# vi dockerfile3 
> [root@localhost dockerfile]# cat dockerfile3
> FROM centos
> ENTRYPOINT ["ls","-a"]
> [root@localhost dockerfile]# docker build -f dockerfile3 -t entrypoint_test:1.0 .
> 
> [root@localhost dockerfile]# docker run entrypoint_test:1.0
> .
> ..
> .dockerenv
> bin
> dev
> etc
> home
> lib
> lib64
> lost+found
> media
> mnt
> opt
> proc
> root
> run
> sbin
> srv
> sys
> tmp
> usr
> var
> # ENTRYPOINT 支持在后面追加参数
> [root@localhost dockerfile]# docker run entrypoint_test:1.0 -l
> total 0
> drwxr-xr-x.   1 root root   6 Apr  5 15:49 .
> drwxr-xr-x.   1 root root   6 Apr  5 15:49 ..
> -rwxr-xr-x.   1 root root   0 Apr  5 15:49 .dockerenv
> lrwxrwxrwx.   1 root root   7 Nov  3  2020 bin -> usr/bin
> drwxr-xr-x.   5 root root 340 Apr  5 15:49 dev
> drwxr-xr-x.   1 root root  66 Apr  5 15:49 etc
> drwxr-xr-x.   2 root root   6 Nov  3  2020 home
> lrwxrwxrwx.   1 root root   7 Nov  3  2020 lib -> usr/lib
> lrwxrwxrwx.   1 root root   9 Nov  3  2020 lib64 -> usr/lib64
> drwx------.   2 root root   6 Sep 15  2021 lost+found
> drwxr-xr-x.   2 root root   6 Nov  3  2020 media
> drwxr-xr-x.   2 root root   6 Nov  3  2020 mnt
> drwxr-xr-x.   2 root root   6 Nov  3  2020 opt
> dr-xr-xr-x. 120 root root   0 Apr  5 15:49 proc
> dr-xr-x---.   2 root root 162 Sep 15  2021 root
> drwxr-xr-x.  11 root root 163 Sep 15  2021 run
> lrwxrwxrwx.   1 root root   8 Nov  3  2020 sbin -> usr/sbin
> drwxr-xr-x.   2 root root   6 Nov  3  2020 srv
> dr-xr-xr-x.  13 root root   0 Apr  5 06:59 sys
> drwxrwxrwt.   7 root root 171 Sep 15  2021 tmp
> drwxr-xr-x.  12 root root 144 Sep 15  2021 usr
> drwxr-xr-x.  20 root root 262 Sep 15  2021 var
> ```
>
> <img src=".\asset\images\image-20240405235218144.png" alt="image-20240405235218144" style="zoom:67%;" />



### 构建Tomcat镜像

1、准备镜像文件 tomcat压缩包 又因为Tomcat 依赖于 jdk 因此还需要jdk的压缩包

2、编写Dockerfile文件 命名为Dockerfile  `build`时就不需要 -f 指定名字了 docker会自动寻找这个文件

```shell
[root@localhost tomcat]# vi Dockerfile
[root@localhost tomcat]# cat Dockerfile 
FROM centos:7
MAINTAINER xlwwwww<xlwwwwwj@163.com>

COPY readme.txt /usr/local/readme.txt

ADD jdk-22_linux-x64_bin.tar.gz /usr/local/
ADD apache-tomcat-10.1.20.tar.gz /usr/local/

RUN yum -y install vim

ENV MYPATH /usr/local
WORKDIR $MYPATH

ENV JAVA_HOME /usr/local/jdk-22_linux-x64_bin
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar

ENV CATALINA_HOME /usr/local/apache-tomcat-10.1.20
ENV CATALINA_BASH /usr/local/apache-tomcat-10.1.20
ENV PATH &PATH:&JAVA_HOME/bin:$CATALINA_HOME/lib:CATALINA_HOME/bin

EXPOSE 8080

CMD /usr/local/apache-tomcat-10.1.20/bin/startup.sh && tail -F /usr/local/apache-tomcat-10.1.20/bin/logs
```



3、构建镜像 docker build

```shell
[root@localhost tomcat]# docker build -t diytomcat:1.0 .
```

4、运行  

- 把内部的webapps/test 映射到本地/home/build/tomcat/test 文件夹 方便同步
- 把内部的日志文件logs 映射到本地/home/build/tomcat/tomcatlogs文件夹 

```shell
[root@localhost tomcat]# docker run -d -p 9090:8080 --name diytomcat -v /home/build/tomcat/test:/usr/local/apache-tomcat-10.1.20/webapps/test -v /home/build/tomcat/tomcatlogs:/usr/local/apache-tomcat-10.1.20/logs diytomcat:1.0
```



### 发布镜像

- 发布到dockerhub

  - 注册账号并登录

  - 在服务器上提交自己的镜像

    ```shell
    # 在服务器登录
    docker login -u -p
    
    # 注销
    docker logout
    ```
  
    
  
  - 提交到dockerhub
  
    ```shell
    # docker push 作者名/镜像名[:Tag]  这里的作者名需要和dockerhub的一致 不然会被拒绝
    docker push kuangshen/diytomcat:1.0
    
    # 给镜像增加标签 或者改名 docker tag
    docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG] 改成符合条件的名字上传就行了
    ```
  
- 发布到阿里云 参考官方文档
  
  - 登录阿里云
  
  - 找到容器镜像服务
  
  - 创建命名空间
  
  - 创建容器镜像
  
  - 上传
  
    ```shell
    # 还是 先注册
    docker login
    
    # 改名
    docker tag     
    
    # 上传
    docker push
    ```
  
    
  
## 小结

  <img src="https://img-blog.csdnimg.cn/8cd02de7be46499b9072ce7c42db19ae.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAc2VsZWN0RGVsZQ==,size_20,color_FFFFFF,t_70,g_se,x_16" alt="Docker安装文档_安装docker引擎的官方文档-CSDN博客" style="zoom: 67%;" />

  ```shell
  # 打包
  docker save [OPTIONS] IMAGE [IMAGE...]
  Save one or more images to a tar archive (streamed to STDOUT by default)
  Aliases:
    docker image save, docker save
  Options:
    -o, --output string   Write to a file, instead of STDOUT
    
  # load
  docker load [OPTIONS]
  Load an image from a tar archive or STDIN
  Aliases:
    docker image load, docker load
  Options:
    -i, --input string   Read from tar archive file, instead of STDIN
    -q, --quiet          Suppress the load output
  ```



  

## Docker网络

### 理解Docker0

> ip addr
>
> ![image-20240406104356957](.\asset\images\image-20240406104356957.png)
>
> 运行镜像后后 可以在查看镜像内的IP地址
>
> ```shell
> [root@localhost ~]# docker run -d -P --name tomcat01 tomcat
> 
> # 可以进去执行ip addr
> [root@localhost ~]# docker exec -it tomcat01 /bin/bash
> root@1b9b1fd5b52c:/usr/local/tomcat# ip addr
> 
> # 也可以不进去，直接追加命令
> [root@localhost ~]# docker exec -it tomcat01 ip addr
> ```
>
> ![image-20240406105606072](.\asset\images\image-20240406105606072.png)
>
> 
>
> 思考：可以从外部ping通这个eth0吗
>
> - 宿主机能够ping通镜像ip
>
> ![image-20240406105747698](.\asset\images\image-20240406105747698.png)
>
> 
>
> 每启动一个Docker容器，docker就会分配一个ip  
>
> 只要安装了docker，就会有一个网卡**docker0桥接模式**，使用的技术是**evth-path技术**。
>
> 每启动一个容器，宿主机都会产生一个新的veth开头的ip  
>
> ![image-20240406110637429](.\asset\images\image-20240406110637429.png)

> 这是容器内的ip addr
>
> ![image-20240406111338026](.\asset\images\image-20240406111338026.png)
>
> 运行容器产生的网卡,都是成对出现的   一端连着协议 一端彼此相连
>
> 正因为此特性,evth-path充当一个桥梁,连接各种虚拟网络设备
>
> 不同镜像间也能够ping通
>
> ```shell
> [root@localhost ~]# docker exec -it tomcat01 /bin/bash
> # ping tomcat02的地址
> root@1b9b1fd5b52c:/usr/local/tomcat# ping 172.17.0.3				
> ```
>
> <img src=".\asset\images\image-20240406113420455.png" alt="image-20240406113420455" style="zoom:67%;" />
>
> tomcat01和tomcat02共同连接docker0 
>
> docker会给所有运行的容器指定一个默认的可用IP,所有运行的容器是通过docker0进行间接通信的,并不是直连的
>
> 

Docker 使用的是Linux的桥接,宿主机是一个Docker容器的网桥 docker0.

容器间的通信是通过Veth转发的.

Docker中所有的网络接口都是虚拟的.

只要运行中的容器删除,对应的网桥就没有了

<img src=".\asset\images\image-20240406114007548.png" alt="image-20240406114007548" style="zoom:67%;" />

### --link

在docker中,所有的docker镜像运行后,得到的IP地址都是随机分配的,

若是我们项目重启后,数据库ip地址换掉了,我们希望能够通过名字来进行访问,  这里能够使用 --link 

```shell
[root@localhost ~]# docker run -d -P --name tomcat03 --link tomcat02 tomcat			# 运行tomcat03与tomcat02连接
[root@localhost ~]# docker exec -it tomcat03 /bin/bash								# 进入环境
root@9ae6ae51b3dc:/usr/local/tomcat# apt update && apt install -y iproute2			# tomcat默认不安装网络 这里安装一下
...
root@9ae6ae51b3dc:/usr/local/tomcat# apt install inetutils-ping						# 安装ping
...
root@9ae6ae51b3dc:/usr/local/tomcat# ping tomcat02									# 能够直接通过镜像名访问了 不需要IP地址了
PING tomcat02 (172.17.0.3): 56 data bytes
64 bytes from 172.17.0.3: icmp_seq=0 ttl=64 time=0.127 ms
64 bytes from 172.17.0.3: icmp_seq=1 ttl=64 time=0.137 ms
64 bytes from 172.17.0.3: icmp_seq=2 ttl=64 time=0.224 ms
```

> tomcat03运行时,--link到tomcat02 但是在tomcat2未配置的情况下,是连接不到tomcat03的
>
> **其实 tomcat3的--link 就是在tomcat3本地host中增加了tomcat02的配置**
>
> ![image-20240406122008930](.\asset\images\image-20240406122008930.png)
>
> 
>
> 现在在Docker中 已经不建议使用 --link了 
>
> docker0 问题, 它不直接支持容器名访问  ---->由此引出自定义网络

### 自定义网络

产看所有的docker网络

![image-20240406122333638](.\asset\images\image-20240406122333638.png)

> 网络模式:
>
> bridge : 桥接  docker默认  自己创建也是用桥接模式
>
> none   : 不配置网络
>
> host    : 和宿主机共享网络
>
> container :  容器网络联通

```shell
# 之前的启动命令,实际上是有 --net bridge 这个默认操作的 这个就是我们的docker0
[root@localhost ~]# docker run -d -P --name tomcat01 --net bridge tomcat

# docker0 特点 默认 容器名不能直接访问  --link能够打通两个容器的连接

# 因此, 我们可以自定义一个网络docker network create
# --driver bridge 桥接  --subnet 192.168.0.0/16 子网  --gateway 192.168.0.1默认网关
[root@localhost ~]# docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet

# 查看详细信息
[root@localhost ~]# docker network inspect mynet
[
    {
        "Name": "mynet",
        "Id": "f4ea54f276bb0efb461ae0cf7d217d622a7e4293bff743a26bf00d2a22e767de",
        "Created": "2024-04-06T12:33:30.885348276+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "192.168.0.0/16",
                    "Gateway": "192.168.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]
```

> 
>
> `docker network create [OPTIONS] NETWORK`
>
> Create a network
>
> Options:
>       --attachable           Enable manual container attachment
>       --aux-address map      Auxiliary IPv4 or IPv6 addresses used by Network driver (default map[])
>       --config-from string   The network from which to copy the configuration
>       --config-only          Create a configuration only network
>   -d, --driver string        Driver to manage the Network (default "bridge")
>       --gateway strings      IPv4 or IPv6 Gateway for the master subnet
>       --ingress              Create swarm routing-mesh network
>       --internal             Restrict external access to the network
>       --ip-range strings     Allocate container ip from a sub-range
>       --ipam-driver string   IP Address Management Driver (default "default")
>       --ipam-opt map         Set IPAM driver specific options (default map[])
>       --ipv6                 Enable IPv6 networking
>       --label list           Set metadata on a network
>   -o, --opt map              Set driver specific options (default map[])
>       --scope string         Control the network's scope
>       --subnet strings       Subnet in CIDR format that represents a network segment

将服务启动在自己的网络

```shell
[root@localhost ~]# docker run -d -P --name tomcat-net-01 --net mynet tomcat
cdb9b487a3d1e9fc00c7e4e10e2740539177ada6974dd4f771c7487bd2189d09
[root@localhost ~]# docker run -d -P --name tomcat-net-02 --net mynet tomcat
602e026b99e5c6abc8ddd291a6c34bf4cdcb8c24e5a6b055f638bfc27ba7896e
[root@localhost ~]# docker network inspect mynet
```

> ![image-20240406152710416](.\asset\images\image-20240406152710416.png)

![image-20240406124429728](.\asset\images\image-20240406124429728.png)

![image-20240406152926653](.\asset\images\image-20240406152926653.png) **这里为啥 tomcat01能pingtomcat02 但是tomcat02不能pingtomcat01**   

自定义的网络 docker都帮我们维护好了对应的关系,推荐我们平时这样使用网络

不同的集群使用不同的网络,保证集群运行健康

### 网络联通

如何打通使用不同网络运行的容器

<img src=".\asset\images\image-20240406161637733.png" alt="image-20240406161637733" style="zoom:67%;" />

> docker network connect
>
> Usage:  docker network connect [OPTIONS] NETWORK CONTAINER
>
> Connect a container to a network
>
> Options:
>       --alias strings           Add network-scoped alias for the container
>       --driver-opt strings      driver options for the network
>       --ip string               IPv4 address (e.g., "172.30.100.104")
>       --ip6 string              IPv6 address (e.g., "2001:db8::33")
>       --link list               Add link to another container
>       --link-local-ip strings   Add a link-local address for the container

```shell
[root@localhost ~]# docker network connect mynet tomcat01
[root@localhost ~]# docker network inspect mynet
# 联通后就是直接将 tomcat01放到了 mynet 网络下
# 这就是 一个容器 两个IP
```

<img src=".\asset\images\image-20240406161138676.png" alt="image-20240406161138676" style="zoom:67%;" />

这里 tomat01已经联通了 但是由于tomcat2没有设置，因此还没有联通

·![image-20240406161418790](.\asset\images\image-20240406161418790.png)

因此，可以通过 docker network connect 联通后 ，跨网络操作别的容器。



### 部署Redis集群

![image-20240406161844523](.\asset\images\image-20240406161844523.png)

```shell
#创建网络
[root@localhost ~]# docker network create redis --subnet 172.38.0.0/15

# 编写redis创建脚本
for port in $(seq 1 6);
do
mkdir -p /mydata/redis/node-${port}/conf
touch /mydata/redis/node-${port}/conf/redis.conf
cat << EOF >/mydata/redis/node-${port}/conf/redis.conf
port 6379
bind 0.0.0.0
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.38.0.1${port}
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes
EOF
done

# 启动redis服务器
docker run -p 6371:6379 -p 16371:16379 --name redis-1 -v /mydata/redis/node-1/data:/data -v /mydata/redis/node-1/conf/redis.conf:/etc/redis/redis.conf -d --net redis --ip 172.38.0.11 redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf

docker run -p 6372:6379 -p 16372:16379 --name redis-2 -v /mydata/redis/node-2/data:/data -v /mydata/redis/node-2/conf/redis.conf:/etc/redis/redis.conf -d --net redis --ip 172.38.0.12 redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf

docker run -p 6373:6379 -p 16373:16379 --name redis-3 -v /mydata/redis/node-3/data:/data -v /mydata/redis/node-3/conf/redis.conf:/etc/redis/redis.conf -d --net redis --ip 172.38.0.13 redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf

docker run -p 6374:6379 -p 16374:16379 --name redis-4 -v /mydata/redis/node-4/data:/data -v /mydata/redis/node-4/conf/redis.conf:/etc/redis/redis.conf -d --net redis --ip 172.38.0.14 redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf

docker run -p 6375:6379 -p 16375:16379 --name redis-5 -v /mydata/redis/node-5/data:/data -v /mydata/redis/node-5/conf/redis.conf:/etc/redis/redis.conf -d --net redis --ip 172.38.0.15 redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf

docker run -p 6376:6379 -p 16376:16379 --name redis-6 -v /mydata/redis/node-6/data:/data -v /mydata/redis/node-6/conf/redis.conf:/etc/redis/redis.conf -d --net redis --ip 172.38.0.16 redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf


# 进入一个容器创建集群
[root@localhost conf]# docker exec -it redis-1 /bin/sh
/data # redis-cli --cluster create 172.38.0.11:6379 172.38.0.12:6379 172.38.0.13:6379 172.38.0.14:6379 172.38.0.15:6379 172.38.0.16:6379 --cluster-replicas 1
>>> Performing hash slots allocation on 6 nodes...
...
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.

```

![image-20240406170616799](.\asset\images\image-20240406170616799.png)



## Docker Swarm

### 环境准备

有多个镜像等待部署 一个一个部署吗 很费时间 

- 首先阿里云购买服务器 **1主三从**

![image-20240406215647792](.\asset\images\image-20240406215647792.png)

> CentOS 7.8  自定义密码   默认安全组

- 四台机器安装docker

- 阿里云镜像加速   阿里云 容器服务 配置镜像加速器

  [容器镜像服务 (aliyun.com)](https://cr.console.aliyun.com/cn-huhehaote/instances/mirrors)

  ```shell
  sudo mkdir -p /etc/docker
  sudo tee /etc/docker/daemon.json <<-'EOF'
  {
    "registry-mirrors": ["https://7e63i884.mirror.aliyuncs.com"]
  }
  EOF
  sudo systemctl daemon-reload
  sudo systemctl restart docker
  ```

  ![image-20240406222321891](.\asset\images\image-20240406222321891.png

### 工作模式

节点   管理节点 和 工作节点 管理节点负责管理工作节点 各

Raft    一致性算法

![image-20240406222518495](.\asset\images\image-20240406222518495.png)



### 搭建集群

私网 公网

> ``` shell
> docker swarm init [OPTIONS]
> 
> Initialize a swarm
> 
> Options:
>       --advertise-addr string                  Advertised address (format: "<ip|interface>[:port]")
>       --autolock                               Enable manager autolocking (requiring an unlock key to start a stopped manager)
>       --availability string                    Availability of the node ("active", "pause", "drain") (default "active")
>       --cert-expiry duration                   Validity period for node certificates (ns|us|ms|s|m|h) (default 2160h0m0s)
>       --data-path-addr string                  Address or interface to use for data path traffic (format: "<ip|interface>")
>       --data-path-port uint32                  Port number to use for data path traffic (1024 - 49151). If no value is set or is set to 0, the default port (4789) is used.
>       --default-addr-pool ipNetSlice           default address pool in CIDR format (default [])
>       --default-addr-pool-mask-length uint32   default address pool subnet mask length (default 24)
>       --dispatcher-heartbeat duration          Dispatcher heartbeat period (ns|us|ms|s|m|h) (default 5s)
>       --external-ca external-ca                Specifications of one or more certificate signing endpoints
>       --force-new-cluster                      Force create a new cluster from current state
>       --listen-addr node-addr                  Listen address (format: "<ip|interface>[:port]") (default 0.0.0.0:2377)
>       --max-snapshots uint                     Number of additional Raft snapshots to retain
>       --snapshot-interval uint                 Number of log entries between Raft snapshots (default 10000)
>       --task-history-limit int                 Task history retention limit (default 5)
> ```



- manager 节点初始化 `docker swarm init`

  ``` shell
  [root@iZhp3f9on9yuqop7e7zcn8Z ~]# docker swarm init --advertise-addr 172.20.242.210
  ```

  ![image-20240406223621189](.\asset\images\image-20240406223621189.png)

  上面的提示，加入一个工作节点    docker swarm join --      

  ```shell
  # 获取令牌
  docker swarm join-token manager			创建manager节点令牌
  docker swarm join-token worker			创建worker节点令牌
  ```

  

- 将其余作为工作节点加入manager节点

  - 以worker节点状态加入

  ```shell
  docker swarm join --token SWMTKN-1-5synkvx3t1txat1b3xr4dvbufm9ump3j1rm82hs2jfvx7hjylc-2pnmr0os924vrlegincw5o94i 172.20.242.210:2377
  ```

  这里在manager节点的主机能够 通过 `docker node ls` 查看连接的节点信息

  ![image-20240406230243823](.\asset\images\image-20240406230243823.png)

  同样 能够在相应的主机 使用 `docker swarm leave`命令离开 swarm

  - 以manager状态加入

    ```shell
    # 在manager节点获取令牌
    [root@iZhp3f9on9yuqop7e7zcn8Z ~]# docker swarm join-token manager
    To add a manager to this swarm, run the following command:
    
        docker swarm join --token SWMTKN-1-5synkvx3t1txat1b3xr4dvbufm9ump3j1rm82hs2jfvx7hjylc-303g27kf8aiue38opvualr7u4 172.20.242.210:2377
        
    # 在相应的主机 以manager身份加入
    [root@iZhp3f9on9yuqop7e7zcn9Z ~]# docker swarm join --token SWMTKN-1-5synkvx3t1txat1b3xr4dvbufm9ump3j1rm82hs2jfvx7hjylc-303g27kf8aiue38opvualr7u4 172.20.242.210:2377
    This node joined a swarm as a manager.
    ```

    ![image-20240406230915259](.\asset\images\image-20240406230915259.png)

  

## CI/CD  jenkins
