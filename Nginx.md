# Nginx 

## 基本概念

### Nginx是什么

Nginx (engine x) 是一个高性能的HTTP和反向代理web服务器，同时也提供了IMAP/POP3/SMTP服务。Nginx是由伊戈尔·赛索耶夫为俄罗斯访问量第二的Rambler.ru站点（俄文：Рамблер）开发的，第一个公开版本0.1.0发布于2004年10月4日。

其将源代码以类BSD许可证的形式发布，因它的稳定性、丰富的功能集、简单的配置文件和低系统资源的消耗而闻名。2011年6月1日，nginx 1.0.4发布。Nginx是一款轻量级的Web 服务器/反向代理服务器及电子邮件（IMAP/POP3）代理服务器，在BSD-like 协议下发行。其特点是占有内存少，并发能力强，事实上nginx的并发能力在同类型的网页服务器中表现较好，中国大陆使用nginx网站用户有：百度、京东、新浪、网易、腾讯、淘宝等。

> **Web服务器**
>
> WEB 服务器也叫网页服务器，英文名叫 Web Server，主要功能是为用户提供网上信息浏览服务。
>
> **HTTP**
>
> HTTP 是超文本传输协议的缩写，是用于从 WEB 服务器传输超文本到本地浏览器的传输协议，也是互联网上应用最为广泛的一种网络协议。HTTP 是一个客户端和服务器端请求和应答的标准，客户端是终端用户，服务端是网站，通过使用 Web 浏览器、网络爬虫或者其他工具，客户端发起一个到服务器上指定端口的 HTTP 请求。
>
> **POP3/SMTP/IMAP**
>
> POP3(Post Offic Protocol 3)邮局协议的第三个版本；
>
> SMTP(Simple Mail Transfer Protocol)简单邮件传输协议；
>
> IMAP(Internet Mail Access Protocol)交互式邮件存取协议；
>
> 通过上述名词的解释，我们可以了解到 Nginx 也可以作为电子邮件代理服务器。
>
> **正向代理**
>
> **<img src=".\asset\images\image-20250907110010709.png" alt="image-20250907110010709" style="zoom:67%;" />**
>
> **反向代理**
>
> <img src=".\asset\images\image-20250907110050323.png" alt="image-20250907110050323" style="zoom:67%;" />
>
> 正向代理的 Proxy 是多对一，反向代理的 Proxy 是一对多。
>
> 
>
> ## 常见服务器对比
>
> 在介绍这一节内容之前，我们先来认识一家公司叫 Netcraft。
>
> > Netcraft 公司于 1994 年底在英国成立，多年来一直致力于互联网市场以及在线安全方面的咨询服务，其中在国际上最具影响力的当属其针对网站服务器、SSL市场所做的客观严谨的分析研究，公司官网每月公布的调研数据（Web Server Survey）已成为当今人们了解全球网站数量以及服务器市场分额情况的主要参考依据，时常被诸如华尔街杂志，英国 BBC，Slashdot 等媒体报道或引用。
>
> 我们先来看一组数据，我们先打开 Nginx 的官方网站 [http://nginx.org/ (opens new window)](http://nginx.org/)，找到 Netcraft 公司公布的数据，对当前主流服务器产品进行介绍。
>
> 上面这张图展示了 2019 年全球主流 Web 服务器的市场情况，其中有 Apache、Microsoft-IIS、google Servers、Nginx、Tomcat 等，而我们在了解新事物的时候，往往习惯通过类比来帮助自己理解事物的概貌。可以看出绿色线的 nginx 在 2019 年已经领现全球了。
>
> 所以下面我们把几种常见的服务器来给大家简单介绍下：
>
> ### [#](https://frxcat.fun/pages/9551ee/#iis)IIS
>
> 全称(Internet Information Services)即互联网信息服务，是由微软公司提供的基于 Windows 系统的互联网基本服务。Windows 作为服务器在稳定性与其他一些性能上都不如类 UNIX 操作系统，因此在需要高性能 Web 服务器的场合下，IIS 可能就会被「冷落」.
>
> ### [#](https://frxcat.fun/pages/9551ee/#tomcat)Tomcat
>
> Tomcat是一个运行 Servlet 和 JSP 的 Web 应用软件，Tomcat 技术先进、性能稳定而且开放源代码，因此深受 Java 爱好者的喜爱并得到了部分软件开发商的认可，成为目前比较流行的 Web 应用服务器。但是 Tomcat 天生是一个重量级的 Web 服务器，对静态文件和高并发的处理比较弱。
>
> ### [#](https://frxcat.fun/pages/9551ee/#apache)Apache
>
> Apache 的发展时期很长，同时也有过一段辉煌的业绩。从上图可以看出大概在 2014 年以前都是市场份额第一的服务器。Apache 有很多优点，如稳定、开源、跨平台等。但是它出现的时间太久了，在它兴起的年代，互联网的产业规模远远不如今天，所以它被设计成一个重量级的、不支持高并发的 Web 服务器。在 Apache 服务器上，如果有数以万计的并发 HTTP 请求同时访问，就会导致服务器上消耗大量能存，操作系统内核对成百上千的 Apache 进程做进程间切换也会消耗大量的 CPU 资源，并导致 HTTP 请求的平均响应速度降低，这些都决定了 Apache 不可能成为高性能的 Web 服务器。这也促使了 Lighttpd 和 Nginx 的出现。
>
> ### [#](https://frxcat.fun/pages/9551ee/#lighttpd)Lighttpd
>
> Lighttpd 是德国的一个开源的 Web 服务器软件，它和 Nginx 一样，都是轻量级、高性能的 Web 服务器，欧美的业界开发者比较钟爱 Lighttpd，而国内的公司更多的青睐 Nginx，同时网上 Nginx 的资源要更丰富些。
>
> ### [#](https://frxcat.fun/pages/9551ee/#其他的服务器)其他的服务器
>
> Google Servers，Weblogic, Webshpere(IBM) ......
>
> Google Servers 是闭源的，市面上很少看到。Weblogic 和 Webshpere(IBM) 使用起来都需要支付一定费用。
>
> 经过各个服务器的对比，种种迹象都表明，Nginx 将以性能为王。这也是我们为什么选择 Nginx 的理由。

### Nginx的优点

### [#](https://frxcat.fun/pages/9551ee/#速度更快、并发更高)速度更快、并发更高

单次请求或者高并发请求的环境下，Nginx 都会比其他 Web 服务器响应的速度更快。一方面在正常情况下，单次请求会得到更快的响应，另一方面，在高峰期(如有数以万计的并发请求)，Nginx 比其他 Web 服务器更快的响应请求。Nginx 之所以有这么高的并发处理能力和这么好的性能原因在于 Nginx 采用了多进程和 I/O 多路复用(epoll)的底层实现。

### [#](https://frxcat.fun/pages/9551ee/#配置简单-扩展性强)配置简单，扩展性强

Nginx 的设计极具扩展性，它本身就是由很多模块组成，这些模块的使用可以通过配置文件的配置来添加。这些模块有官方提供的也有第三方提供的模块，如果需要完全可以开发服务自己业务特性的定制模块

### [#](https://frxcat.fun/pages/9551ee/#高可靠性)高可靠性

Nginx 采用的是多进程模式运行，其中有一个 master 主进程和 N 多个 worker 进程，一个 master 管理多个 worker，worker 进程的数量我们可以手动设置，每个 worker 进程之间都是相互独立提供服务，并且 master 主进程可以在某一个 worker 进程出错时，快速去「拉起」新的 worker 进程提供服务。

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220727/image.3wz8th685620.webp)

### [#](https://frxcat.fun/pages/9551ee/#热部署-核心)热部署(核心)

现在互联网项目都要求以 7 * 24 小时进行服务的提供，针对于这一要求，Nginx 也提供了热部署功能，即可以在 Nginx 不停止的情况下，对 Nginx 进行文件升级、更新配置和更换日志文件等功能。

### [#](https://frxcat.fun/pages/9551ee/#成本低、bsd许可证)成本低、BSD许可证

BSD 是一个开源的许可证，世界上的开源许可证有很多，现在比较流行的有六种分别是 GPL、BSD、MIT、Mozilla、Apache、LGPL。这六种的区别是什么，我们可以通过下面一张图来解释下：

<img src=".\asset\images\image-20250907111115227.png" alt="image-20250907111115227" style="zoom:67%;" />

### [#](https://frxcat.fun/pages/9551ee/#nginx的功能特性及常用功能)Nginx的功能特性及常用功能

Nginx 提供的基本功能服务从大体上归纳为「基本 HTTP 服务」、「高级 HTTP 服务」和「邮件服务」等三大类。

### [#](https://frxcat.fun/pages/9551ee/#基本http服务)基本HTTP服务

Nginx 可以提供基本 HTTP 服务，可以作为 HTTP 代理服务器和反向代理服务器，支持通过缓存加速访问，可以完成简单的负载均衡和容错，支持包过滤功能，支持 SSL 等。

- 处理静态文件、处理索引文件以及支持自动索引
- 提供反向代理服务器，并可以使用缓存加上反向代理，同时完成负载均衡和容错
- 提供对 FastCGI、memcached 等服务的缓存机制，，同时完成负载均衡和容错
- 使用 Nginx 的模块化特性提供过滤器功能。Nginx 基本过滤器包括 gzip 压缩、ranges 支持、chunked 响应、XSLT、SSI 以及图像缩放等。其中针对包含多个 SSI 的页面，经由 FastCGI 或反向代理，SSI 过滤器可以并行处理
- 支持 HTTP 下的安全套接层安全协议 SSL.
- 支持基于加权和依赖的优先权的 HTTP/2

### [#](https://frxcat.fun/pages/9551ee/#高级http服务)高级HTTP服务

- 支持基于名字和 IP 的虚拟主机设置
- 支持 HTTP/1.0 中的 KEEP-Alive 模式和管线(PipeLined)模型连接
- 自定义访问日志格式、带缓存的日志写操作以及快速日志轮转
- 提供 3xx~5xx 错误代码重定向功能
- 支持重写（Rewrite)模块扩展
- 支持重新加载配置以及在线升级时无需中断正在处理的请求
- 支持网络监控
- 支持 FLV 和 MP4 流媒体传输

### [#](https://frxcat.fun/pages/9551ee/#邮件服务)邮件服务

Nginx 提供邮件代理服务也是其基本开发需求之一，主要包含以下特性：

- 支持 IMPA/POP3 代理服务功能
- 支持内部 SMTP 代理服务功能

### [#](https://frxcat.fun/pages/9551ee/#nginx常用的功能模块)Nginx常用的功能模块

- 静态资源部署：核心功能
- Rewrite 地址重写：用到正则表达式
- 反向代理
- 负载均衡：轮询、加权轮询、ip_hash、url_hash、fair 等算法
- Web 缓存
- 环境部署：搭建高可用的环境
- 用户认证模块 ...
- Nginx 的核心组成
  - nginx 二进制可执行文件（启动、关闭、加载 Nginx）
  - nginx.conf 配置文件
  - error.log 错误的日志记录
  - access.log 访问日志记录



## 安装

### 环境准备

- 虚拟机：VMware WorkStation
- Linux 系统：Centos7
- 上传软件工具：MobaXterm、xsheel,SecureCRT 之一
- 网络

#### 确认CentOS的内核

准备一个内核为 2.6 及以上版本的操作系统，因为 linux2.6 及以上内核才支持 epoll,而 Nginx 需要解决高并发压力问题是需要用到 epoll，所以我们需要有这样的版本要求。

我们可以使用 `uname -a` 命令来查询linux的内核版本



```sh
[root@master ~]# uname -a
Linux master 3.10.0-1160.el7.x86_64 #1 SMP Mon Oct 19 16:18:59 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```

#### 确保Centos能联网

```sh
ping www.baidu.com
```

三种网络模式：桥接、主机、NAT，建议使用虚拟机使用 NAT 模式

NAT 模式下，只要主机有网络，则虚拟机也有网络。

#### 确认关闭防火墙

这一项的要求仅针对于那些对 Linux 系统的防火墙设置规则不太清楚的，建议大家把防火墙都关闭掉，因为把防火墙关闭掉，可以省掉后续 Nginx 使用过程中遇到的诸多问题。

关闭的方式有如下两种:

```sh
systemctl stop firewalld      # 关闭运行的防火墙，系统重新启动后，防火墙将重新打开
systemctl disable firewalld   # 永久关闭防火墙，，系统重新启动后，防火墙依然关闭
systemctl status firewalld	 # 查看防火墙状态
```

如果不想关闭防火墙，记得开放 Nginx 的访问端口

#### 确认停用selinux

selinux(security-enhanced linux)，美国安全局对于强制访问控制的实现，在 Linux2.6 内核以后的版本中，selinux 已经成功内核中的一部分。可以说selinux 是 Linux 史上最杰出的新安全子系统之一。虽然有了 selinux，我们的系统会更安全，但是对于使用 Nginx 或者其他的学习过程中，会有很多设置，所以这块建议大家将 selinux 进行关闭。

`sestatus` 查看状态

```sh
[root@master ~]# sestatus
SELinux status:                 disabled
```

如果查看不是 disabled 状态，我们可以通过修改配置文件来进行设置,修改 `SELINUX=disabled`，然后重启下系统即可生效。

```sh
vim /etc/selinux/config
```

#### Nginx环境安装

Nginx 的安装方式有两种分别是：

- 通过 Nginx 源码
  - 通过 Nginx 源码简单安装
  - 通过 Nginx 源码复杂安装
- 通过 yum 安装

如果通过 Nginx 源码安装需要提前准备安装好环境，我相信，任何 C 语言编写的软件都离不开这些环境，安装后所有软件都会收益，往下看。

#### GCC编译器

Nginx 是使用 C 语言编写的程序，因此想要运行 Nginx 就需要安装一个编译工具。GCC 就是一个开源的编译器集合，用于处理各种各样的语言，其中就包含了 C 语言。

使用命令 `yum install -y gcc` 来安装

安装成功后，可以通过 `gcc --version` 来查看 gcc 是否安装成功

```sh
# 安装 gcc
yum install -y gcc

# 查看 gcc 版本
gcc --version
```

#### PCRE

Nginx 在编译过程中需要使用到 PCRE 库（perl Compatible Regular Expressoin 兼容正则表达式库)，因为在 Nginx 的 Rewrite 模块和 Http 核心模块都会使用到 PCRE 正则表达式语法。

可以使用命令 `yum install -y pcre pcre-devel` 来进行安装

安装成功后，可以通过 `rpm -qa pcre pcre-devel` 来查看是否安装成功



```sh
# 安装 PCRE 和 pcre-devel
yum install -y pcre pcre-devel

# 查看 PCRE 是否安装成功
rpm -qa pcre pcre-devel
```

#### zlib

zlib 库提供了开发人员的压缩算法，在 Nginx 的各个模块中需要使用 gzip 压缩，所以我们也需要提前安装其库及源代码 zlib 和 zlib-devel。

可以使用命令 `yum install -y zlib zlib-devel` 来进行安装

安装成功后，可以通过 `rpm -qa zlib zlib-devel` 来查看是否安装成功



```sh
# 安装 zlib 和 zlib-devel
yum install -y zlib zlib-devel

# 查看 zlib 和 zlib-devel 是否安装成功
rpm -qa zlib zlib-devel
```

#### OpenSSL

OpenSSL 是一个开放源代码的软件库包，应用程序可以使用这个包进行安全通信，并且避免被窃听。

SSL：Secure Sockets Layer 安全套接协议的缩写，可以在 Internet 上提供秘密性传输，其目标是保证两个应用间通信的保密性和可靠性。在 Nginx 中，如果服务器需要提供安全网页时就需要用到 OpenSSL 库，所以我们需要对 OpenSSL 的库文件及它的开发安装包进行一个安装。

可以使用命令 `yum install -y openssl openssl-devel` 来进行安装

安装成功后，可以通过 `rpm -qa openssl openssl-devel` 来查看是否安装成功



```sh
# 安装 OpenSSL 库
yum install -y openssl openssl-devel

# 查看 OpenSSL 是否安装成功
rpm -qa openssl openssl-devel
```

#### 全部安装命令

上述命令，一个个来的话比较麻烦，我们也可以通过一条命令来进行安装

`yum install -y gcc pcre pcre-devel zlib zlib-devel openssl openssl-devel` 进行全部安装。



```sh
# 全部安装 GCC、PCRE、zlib、OpenSSL
yum install -y gcc pcre pcre-devel zlib zlib-devel openssl openssl-devel

# 查看安装版本
rpm -qa pcre pcre-devel zlib zlib-devel openssl openssl-devel
```

### Nginx安装

Nginx 的官方下载网站为 [http://nginx.org/en/download.html(opens new window)](http://nginx.org/en/download.html)

#### 源码简单安装

进入官网查找需要下载版本的链接地址，然后使用 wget 命令进行下载，这里下载在 `/opt` 目录下

```sh
cd /opt

wget http://nginx.org/download/nginx-1.21.6.tar.gz
```

建议大家将下载的资源进行包管理，在 `/opt` 目录下创建 `nginx/core` 目录，`-p` 代表允许创建多级目录

```sh
# 创建指定目录
mkdir -p /opt/nginx/core

# 移动到指定目录
mv nginx-1.21.6.tar.gz nginx/core
```

解压缩

```sh
# 进入指定目录
cd /opt/nginx/core

# 解压
tar -xzf nginx-1.21.6.tar.gz
```

1. 进入资源文件中，发现 configure 文件，执行该文件

```sh
# 进入解压出的目录
cd nginx-1.21.6/

# 执行 configure 
./configure --prefix=/usr/local/nginx
```

编译

```sh
make
```

安装

```sh
make install
```

查看 nginx 的安装位置

```sh
whereis nginx
```

```sh
[root@master sbin]# whereis nginx
nginx: /usr/local/nginx
```

启动测试。启动文件在安装目录下的 sbin 命令

```sh
# 进入启动文件的目录
cd /usr/local/nginx/sbin

# 执行 nginx 文件
./nginx
```

访问自己 Linux 的 IP 地址，不需要加端口，如果看到如下图，代表成功

#### yum安装

使用源码进行简单安装，我们会发现安装的过程比较繁琐，需要提前准备 GCC 编译器、PCRE 兼容正则表达式库、zlib 压缩库、OpenSSL 安全通信的软件库包，然后才能进行 Nginx 的安装。下面使用 yum 安装，过程更简单。

安装 yum-utils

```sh
sudo yum  install -y yum-utils
```

添加 yum 源文件

```sh
vim /etc/yum.repos.d/nginx.repo
```

添加如下内容

```sh
[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true

[nginx-mainline]
name=nginx mainline repo
baseurl=http://nginx.org/packages/mainline/centos/$releasever/$basearch/
gpgcheck=1
enabled=0
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true
```

使用 yum 安装 Nginx

```sh
sudo yum install -y nginx
```

查看是否安装成功

```sh
yum list | grep nginx
```

使用 yum 进行安装

```sh
yun install -y nginx
```

查看nginx的安装位置

```sh
whereis nginx
```

启动测试

```sh
# 进入文件
cd /usr/sbin

# 启动 nginx
./nginx
```

#### 源码简单安装和yum安装的差异

这里先介绍一个命令: `./nginx -V`，通过该命令可以查看到所安装 Nginx 的版本及相关配置信息。

简单安装的配置信息查看：

```sh
[root@master sbin]# ./nginx -V
nginx version: nginx/1.21.6
built by gcc 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC)
configure arguments: --prefix=/usr/local/nginx
```

yum 安装的配置信息查看：（这里截取 1.16.1 版本的 图片）

可以看出 yum 安装的参数非常多，如果想源码简单安装也有这些参数，这涉及到源码复杂安装，往下看。

执行 `tar -zxvf nginx-1.16.1.tar.gz` 对下载的资源进行解压缩后，进入压缩后的目录，可以看到如下结构

内容解释：

| 内容名     | 作用                                                         |
| ---------- | ------------------------------------------------------------ |
| auto       | 存放的是编译相关的脚本                                       |
| CHANGES    | 版本变更记录                                                 |
| CHANGES.ru | 俄罗斯文的版本变更记录                                       |
| conf       | Nginx 默认的配置文件                                         |
| configure  | Nginx 软件的自动脚本程序,是一个比较重要的文件，作用如下： 1、检测环境及根据环境检测结果生成 C 代码 2、生成编译代码需要的 Makefile 文件 |
| contrib    | 存放的是几个特殊的脚本文件，其中 README 中对脚本有着详细的说明 |
| html       | 存放的是 Nginx 自带的两个 html 页面，访问 Nginx 的首页和错误页面 |
| LICENSE    | 许可证的相关描述文件                                         |
| man        | Nginx 的 man 手册                                            |
| README     | Nginx 的阅读指南                                             |
| src        | Nginx 的源代码                                               |

#### 源码复杂安装

这种方式和简单的安装配置不同的地方在第一步，通过 `./configure` 来对编译参数进行设置，需要我们手动来指定。那么都有哪些参数可以进行设置，接下来我们进行一个详细的说明。

- PATH：是和路径相关的配置信息
- with：是启动模块，默认是关闭的
- without：是关闭模块，默认是开启的

我们先来认识一些简单的路径配置已经通过这些配置来完成一个简单的编译：

| 指令                  | 作用                              | 默认值                     |
| --------------------- | --------------------------------- | -------------------------- |
| --prefix=PATH         | 指向 Nginx 的安装目录             | /usr/local/nginx           |
| --sbin-path=PATH      | 指向(执行)程序文件(nginx)的路径   | <安装目录>/sbin/nginx      |
| --modules-path=PATH   | 指向 Nginx 动态模块安装目录       | <安装目录>//modules        |
| --conf-path=PATH      | 指向配置文件(nginx.conf)的路径    | <安装目录>/conf/nginx.conf |
| --error-log-path=PATH | 指向错误日志文件的路径            | <安装目录>/logs/error.log  |
| --http-log-path=PATH  | 指向访问日志文件的路径            | <安装目录>/logs/access.log |
| --pid-path=PATH       | 指向 Nginx 启动后进行ID的文件路径 | <安装目录>/logs/nginx.pid  |
| --lock-path=PATH      | 指向 Nginx 锁文件的存放路径       | <安装目录>/logs/nginx.lock |

源码复杂安装命令如下：（先进入解压目录）

```sh
# 进入解压目
cd ~/nginx/core

# 安装命令
./configure --prefix=/usr/local/nginx \
--sbin-path=/usr/local/nginx/sbin/nginx \
--modules-path=/usr/local/nginx/modules \
--conf-path=/usr/local/nginx/conf/nginx.conf \
--error-log-path=/usr/local/nginx/logs/error.log \
--http-log-path=/usr/local/nginx/logs/access.log \
--pid-path=/usr/local/nginx/logs/nginx.pid \
--lock-path=/usr/local/nginx/logs/nginx.lock

# 执行编译和安装
make && make install
```

在使用上述命令之前，需要将之前服务器已经安装的 Nginx 进行卸载。

#### Nginx卸载

简单安装后如果想要卸载，步骤分为三步骤：

步骤一：需要将 Nginx 的进程关闭（要在 安装目录的 sbin 目录下）

```sh
./nginx -s stop
```

步骤二：将安装的 Nginx 进行删除

```sh
rm -rf /usr/local/nginx
```

步骤三：进入解压目录，将安装包之前编译的环境清除掉

```sh
# 进入解压目录
cd /opt/nginx/core

# 清除环境
make clean
```

#### Nginx目录结构分析

在使用 Nginx 之前，我们先对安装好的 Nginx 目录文件进行一个分析，在这块给大家介绍一个工具 tree，通过 tree 我们可以很方面的去查看 Centos 系统上的文件目录结构，当然，如果想使用 tree 工具，就得先通过 `yum install -y tree` 来进行安装，安装成功后，可以通过执行 `tree /usr/local/nginx` (tree 后面跟的是 Nginx 的安装目录)，获取的结果如下：

```sh
[root@master /]# tree /usr/local/nginx
/usr/local/nginx
├── client_body_temp
├── conf
│   ├── fastcgi.conf
│   ├── fastcgi.conf.default
│   ├── fastcgi_params
│   ├── fastcgi_params.default
│   ├── koi-utf
│   ├── koi-win
│   ├── mime.types
│   ├── mime.types.default
│   ├── nginx.conf
│   ├── nginx.conf.default
│   ├── scgi_params
│   ├── scgi_params.default
│   ├── uwsgi_params
│   ├── uwsgi_params.default
│   └── win-utf
├── fastcgi_temp
├── html
│   ├── 50x.html
│   └── index.html
├── logs
│   ├── access.log
│   ├── error.log
│   └── nginx.pid
├── proxy_temp
├── sbin
│   ├── nginx
│   └── nginx.old
├── scgi_temp
└── uwsgi_temp

9 directories, 22 files
```

CGI(Common Gateway Interface)通用网关【接口】，主要解决的问题是从客户端发送一个请求和数据，服务端获取到请求和数据后可以调用调用 CGI【程序】处理及相应结果给客户端的一种标准规范。

| 目录 | 文件名                    | 作用                                                         |
| ---- | ------------------------- | ------------------------------------------------------------ |
| conf |                           | Nginx 所有配置文件目录                                       |
|      | fastcgi.conf              | fastcgi相关配置文件                                          |
|      | fastcgi.conf.default      | fastcgi.conf 的备份文件                                      |
|      | fastcgi_params            | fastcgi 的参数文件                                           |
|      | fastcgi_params.default    | fastcgi 的参数备份文件                                       |
|      | scgi_params               | scgi 的参数文件                                              |
|      | scgi_params.default       | scgi 的参数备份文件                                          |
|      | uwsgi_params              | uwsgi 的参数文件                                             |
|      | uwsgi_params.default      | uwsgi 的参数备份文件                                         |
|      | mime.types                | **记录的是 HTTP 协议中的 Content-Type 的值和文件后缀名的对应关系** |
|      | mime.types.default        | mime.types 的备份文件                                        |
|      | nginx.conf                | **这是 Nginx 的核心配置文件，这个文件非常重要，也是我们即将要学习的重点** |
|      | nginx.conf.default        | nginx.conf 的备份文件                                        |
|      | koi-utf、koi-win、win-utf | 这三个文件都是与编码转换映射相关的配置文件，用来将一种编码转换成另一种编码 |
| html |                           | 存放 Nginx 自带的两个静态的 html 页面                        |
|      | 50x.html                  | 访问失败后的失败页面                                         |
|      | index.html                | 成功访问的默认首页                                           |
| logs |                           | 记录入门的文件，当 Nginx 服务器启动后，这里面会有 access.log error.log 和 nginx.pid 三个文件出现 |
|      | access.log                | **访问日志，每次访问成功都会进行记录**                       |
|      | error.log                 | **错误日志，每次访问失败都会进行记录**                       |
|      | nginx.pid                 | 启动 Nginx 后，系统生成一个进程 PID，这个文件记录这个 PID    |
| sbin |                           | 是存放执行程序文件 nginx                                     |
|      | nginx                     | 用来控制 Nginx 的启动和停止等相关的命令。**注意：该文件名就叫 nginx** |

## nginx基本使用

#### 目录结构

进入Nginx的主目录我们可以看到这些文件夹

```sh
[root@master /]# cd /usr/local/nginx/
[root@master nginx]# ll
总用量 0
drwx------ 2 nobody root   6 7月  27 20:35 client_body_temp
drwxr-xr-x 2 root   root 333 7月  27 20:31 conf
drwx------ 2 nobody root   6 7月  27 20:35 fastcgi_temp
drwxr-xr-x 2 root   root  40 7月  27 20:31 html
drwxr-xr-x 2 root   root  58 7月  27 20:35 logs
drwx------ 2 nobody root   6 7月  27 20:35 proxy_temp
drwxr-xr-x 2 root   root  36 7月  27 20:33 sbin
drwx------ 2 nobody root   6 7月  27 20:35 scgi_temp
drwx------ 2 nobody root   6 7月  27 20:35 uwsgi_temp
```

其中这几个文件夹在刚安装后是没有的，主要用来存放运行过程中的临时文件

```sh
client_body_temp fastcgi_temp proxy_temp scgi_temp
```

- 主要目录说明

| 目录 | 说明                                   |
| ---- | -------------------------------------- |
| conf | 用来存放配置文件相关                   |
| html | 用来存放静态文件的默认目录 html、css等 |
| sbin | nginx的主程序                          |



#### 基本运行原理

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220727/image.5klrdamfdsg0.webp)

#### 启动停止命令

对于 Nginx 的启停在 Linux 系统中也有很多种方式，我们介绍两种方式：

- Nginx 服务的信号控制
- Nginx 的命令行控制

#### 服务信号控制

在了解内容之前，我们首先要考虑一些问题：

问题

```
Nginx 中的 master 和 worker 进程?

Nginx 的工作方式?

如何获取进程的 PID?

信号有哪些?

如何通过信号控制 Nginx 的启停等相关操作?
```

前面在提到 Nginx 的高性能，其实也和它的架构模式有关。Nginx 默认采用的是多进程的方式来工作的，当将 Nginx 启动后，我们通过 `ps -ef | grep nginx` 命令可以查看到如下内容：

```sh
ps -ef | grep nginx
```

```sh
[root@master sbin]# ps -ef| grep nginx
root       3564      1  0 13:54 ?        00:00:00 nginx: master process ./nginx
nobody     3565   3564  0 13:54 ?        00:00:00 nginx: worker process
root       3567   3483  0 13:54 pts/1    00:00:00 grep --color=auto nginx
```

从上图中可以看到，Nginx 后台进程中包含一个 master 进程和多个 worker 进程，master 进程主要用来管理 worker 进程，包含接收外界的信息，并将接收到的信号发送给各个 worker 进程，监控 worker 进程的状态。当 worker 进程出现异常退出后，会自动重新启动新的 worker 进程。而 worker 进程则是专门用来处理用户请求的，各个 worker 进程之间是平等的并且相互独立，处理请求的机会也是一样的。

Nginx 的进程模型，我们可以通过下图来说明下：

<img src=".\asset\images\image-20250907130442439.png" alt="image-20250907130442439" style="zoom:67%;" />

我们现在作为管理员，只需要通过给 master 进程发送信号就可以来控制 Nginx，这个时候我们需要有两个前提条件，一个是要操作的 master 进程，一个是 给 master 进程的信号。

1. 要想操作 Nginx 的 master 进程，就需要获取到 master 进程的进程号 PID。获取方式简单介绍两个：

- 通过 `ps -ef | grep nginx`

```sh
ps -ef | grep nginx
```

- 在讲解 Nginx 的 `./configure` 的配置参数的时候，有一个参数 `--pid-path=PATH`，它的默认值是 `/usr/local/nginx/logs/nginx.pid`，所以可以通过查看该文件来获取 Nginx 的 master 进程 PID

```sh
cat /usr/local/nginx/logs/nginx.pid
```

1. 信号(signal)

| 信号     | 作用                                                         |
| -------- | ------------------------------------------------------------ |
| TERM/INT | 立即关闭整个服务（关闭 Nginx）                               |
| QUIT     | 「优雅」的关闭整个服务（关闭 Nginx）                         |
| HUP      | 重读配置文件并使用服务对新配置项生效（重启 Nginx）           |
| USR1     | 重新打开日志文件，可以用来进行日志切割（重启日志）           |
| USR2     | 平滑升级到最新版的 Nginx                                     |
| WINCH    | 所有子进程不在接收处理新连接，相当于给 Work 进程发送 QUIT 指令 |



调用命令为 `kill -signal PID`

> signal：即为信号；PID 即为获取到的 master 进程 PID

```sh
kill -signal PID
```

```sh
# 格式一：
kill -TERM PID

# 立即关闭当前线程
kill -TERM `cat /usr/local/nginx/logs/nginx.pid`

# 格式一：
kill -INT PID

# 立即关闭当前线程
kill -INT `cat /usr/local/nginx/logs/nginx.pid`
```

- 案例

1. 发送 TERM/INT 信号给 master 进程，会将 Nginx 服务立即关闭。

```sh
[root@master nginx]# cat logs/nginx.pid
3564
[root@master nginx]# kill  -TERM 3564
[root@master nginx]# ps -ef | grep nginx
root       8874   8753  0 15:58 pts/2    00:00:00 grep --color=auto nginx
```

- 发送 QUIT 信号给 master 进程，master 进程会控制所有的 work 进程不再接收新的请求，等所有请求处理完后，在把进程都关闭掉。

```sh
# 优雅 关闭线程
kill -QUIT PID

# 「优雅」关闭当前线程
kill -QUIT `cat /usr/local/nginx/logs/nginx.pid`
```

- 发送 HUP 信号给 master 进程，master 进程会把控制旧的 worker 进程不再接收新的请求，等处理完请求后将旧的 worker 进程关闭掉，然后根据更改Nginx 的配置文件重新启动新的 worker 进程

```sh
# 重启 worker 进程
kill -HUP PID 

# 重启当前 worker 进程
kill -HUP `cat /usr/local/nginx/logs/nginx.pid`
```

- 发送 USR1 信号给 master 进程，告诉 Nginx 重新开启日志文件。如果日志文件被删除了，可以利用此命令重新打开。

```sh
# 重新打开日志文件
kill -USR1 PID

# 重新打开当前 Nginx 的日志文件
kill -USR1 `cat /usr/local/nginx/logs/nginx.pid`
```

- 发送 USR2 信号给 master 进程，告诉 master 进程要平滑升级，

  这个时候，会重新开启对应的 master 进程和 worker 进程，整个系统中将会有两个master 进程，并且新的 master 进程的 PID 会被记录在 `/usr/local/nginx/logs/nginx.pid`，而之前的旧的 master 进程 PID 会被记录在 `/usr/local/nginx/logs/nginx.pid.oldbin` 文件中，

  接着再次发送 QUIT 信号给旧的 master 进程，让其处理完请求后再进行关闭

```sh
# 开启新的进程，但是不删除旧的进程
kill -USR2 PID

# 开启新的进程，但是不删除当前进程
kill -USR2 `cat /usr/local/nginx/logs/nginx.pid`
```

当新进程升级后（完全启动后），再关闭旧的进程，旧进程的 PID 在另一个 `nginx.pid.oldbin` 文件里

```sh
# 关闭旧的线程
kill -QUIT `cat /usr/local/nginx/logs/nginx.pid.oldbin`
```

- 发送 WINCH 信号给 master 进程,让 master 进程控制不让所有的 worker 进程在接收新的请求了，请求处理完后关闭 worker 进程。注意 master 进程不会被关闭掉



```sh
# 停止 worker 进程，但是不停止 master 进程
kill -WINCH PID 

# 停止当前 worker 进程，但是不停止 master 进程
kill -WINCH `cat /usr/local/nginx/logs/nginx.pid`
```

#### 命令行控制

此方式是通过 Nginx 安装目录下的 sbin 下的可执行文件 nginx(文件名) 来进行对 Nginx 状态的控制，我们可以通过 `nginx -h` 来查看都有哪些参数可以用

```sh
[root@master sbin]# ./nginx -h
nginx version: nginx/1.21.6
Usage: nginx [-?hvVtTq] [-s signal] [-p prefix]
             [-e filename] [-c filename] [-g directives]

Options:
  -?,-h         : this help
  -v            : show version and exit
  -V            : show version and configure options then exit
  -t            : test configuration and exit
  -T            : test configuration, dump it and exit
  -q            : suppress non-error messages during configuration testing
  -s signal     : send signal to a master process: stop, quit, reopen, reload
  -p prefix     : set prefix path (default: /usr/local/nginx/)
  -e filename   : set error log file (default: logs/error.log)
  -c filename   : set configuration file (default: conf/nginx.conf)
  -g directives : set global directives out of configuration file
```

| 选项     | 作用                                                         |
| -------- | ------------------------------------------------------------ |
| -? 和 -h | 显示帮助信息                                                 |
| -v       | 打印版本号信息并退出                                         |
| -V       | 打印版本号信息和配置信息并退出                               |
| -t       | 测试 Nginx 的配置文件语法是否正确并退出                      |
| -T       | 测试 Nginx 的配置文件语法是否正确并列出用到的配置文件信息然后退出 |
| -q       | 在配置测试期间过滤掉非错误消息                               |
| -s       | signal 信号，后面的命令和服务信号控制功能类似：<br /> `stop`：快速关闭，类似于 TERM/INT 信号的作用<br /> `quit`：优雅的关闭，类似于 QUIT 信号的作用<br /> `reopen`：重新打开日志文件类似于 USR1 信号的作用 <br />`reload`：重启 Nginx，类似于 HUP 信号的作用 |
| -p       | prefix，指定 Nginx 的默认安装路径，(默认为：/usr/local/nginx/) |
| -c       | filename，指定 Nginx 的配置文件路径,(默认为：conf/nginx.conf) |
| -g       | 用来补充 Nginx 配置文件，向 Nginx 服务指定启动时应用全局的配置 |

> - 案例
>
> 如果觉得每次执行 nginx 指令都必须进入 sbin 目录，则将该指令设置为全局使用。
>
> 1. 两个查看版本命令
>
> ```sh
> # 查看版本指令 1
> [root@master sbin]# nginx -v
> # 返回结果
> nginx version: nginx/1.21.6
> 
> # 查看版本指令 2
> [root@master sbin]# nginx -V
> # 返回结果
> nginx version: nginx/1.21.6
> built by gcc 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC)
> configure arguments: --prefix=/usr/local/nginx
> ```
>
> 1. 测试 Nginx 的配置文件语法
>
> 我们首先要知道配置文件的路径在哪，先执行 `-t` 进行测试
>
> ```sh
> # 测试 Nginx 的配置文件语法
> [root@master sbin]# nginx -t
> 
> #结果 返回 成功
> nginx: the configuration file /usr/local/nginx/conf/nginx.conf syntax is ok
> nginx: configuration file /usr/local/nginx/conf/nginx.conf test is successful
> ```
>
> 由第 6 行代码可以知道测试成功，第 5 行代码告诉我们配置文件的目录，我们去修改配置文件，然后再进行测试
>
> ```sh
> [root@master sbin]# vim /usr/local/nginx/conf/nginx.conf
> 
> ###add
> HelloWorld
> ```
>
> 重新进行测试
>
> ```sh
> # 测试 Nginx 的配置文件语法
> [root@master sbin]# nginx -t
> 
> # 返回结果（失败）
> nginx: [emerg] unknown directive "HelloWorld" in /usr/local/nginx/conf/nginx.conf:3
> nginx: configuration file /usr/local/nginx/conf/nginx.conf test failed
> ```
>
> 由第 6 行代码可以知道，配置文件出错了。验证完了，记得将配置文件改回来。
>
> 1. 指定 Nginx 的默认安装路径
>
> ```sh
> nginx -p /usr/local/nginx/
> ```
>
> 1. 指定 Nginx 的配置文件路径
>
> 先把配置文件拷贝到另一个目录，然后修改拷贝后的配置文件内容
>
> ```sh
> # 拷贝配置文件
> cp /usr/local/nginx/conf/nginx.conf /opt
> 
> # 修改拷贝后的配置文件内容
> vim /opt/nginx.conf
> 
> # add
> HelloWorld
> ```
>
> 测试配置文件的时候，指定拷贝后的配置文件进行测试
>
> ```sh
> # 指定配置文件进行测试
> [root@master nginx]# nginx -tc /opt/nginx.conf
> 
> #返回结果
> nginx: [emerg] unknown directive "HelloWorld" in /opt/nginx.conf:3
> nginx: configuration file /opt/nginx.conf test failed
> ```
>
> 说明指定配置文件目录生效，只是文件内容语法不对。

#### 版本升级和新增模块

如果想对 Nginx 的版本进行更新，或者要应用一些新的模块，最简单的做法就是停止当前的 Nginx 服务，然后开启新的 Nginx 服务。

但是这样会导致在一段时间内，用户是无法访问服务器。为了解决这个问题，我们就需要用到 Nginx 服务器提供的平滑升级功能。

这个也是 Nginx 的一大特点，使用这种方式，就可以使 Nginx 在 7 * 24 小时不间断的提供服务了。接下来我们分析下需求：

> 需求：Nginx 的版本最开始使用的是 Nginx-1.14.2，由于服务升级，需要将 Nginx 的版本升级到 Nginx-1.16.1,要求 Nginx 不能中断提供服务。

为了应对上述的需求，这里我们提供两种解决方案:

- 使用 Nginx 服务信号完成 Nginx 的升级
- 使用 Nginx 安装目录的 make 命令完成升级

版本升级其实就是替换可执行文件 nginx。

##### 环境准备

1. 先准备两个版本的 Nginx 分别是 1.14.2 和 1.16.1
2. 使用 Nginx 源码安装的方式将 1.14.2 版本安装成功并正确访问

```sh
# 解压 1.14.2 版本
tar -xzf nginx-1.14.2.tar.gz
# 进入解压目录
cd nginx-1.14.2/

# 执行配置文件
./configure

# 编译安装
make && make install
```

1. 将 Nginx 1.16.1 版本进行参数配置和编译，不需要进行安装。

```sh
# 解压 1.16.1 版本
tar -xzf nginx-1.16.1.tar.gz
# 进入解压目录
cd nginx-1.16.1/

# 执行配置文件
./configure

# 仅仅编译
make 
```

##### 服务信号进行升级

第一步:将 1.14.2 版本的 sbin 目录下的 nginx 进行备份

> 不是复制一份，是直接修改原来的 nginx。

```sh
# 进入 sbin 目录下
cd /usr/local/nginx/sbin

# 备份为 nginxold 文件
mv nginx nginx.backup
```

第二步：将 Nginx 1.16.1 安装目录编译后的 objs 目录下的 nginx 文件，拷贝到原来 `/usr/local/nginx/sbin` 目录下

如果第一步没有备份，那么将会覆盖 1.14.2 的 nginx 文件

```sh
# 进入 objs 目录
cd ~/nginx/core/nginx-1.16.1/objs

# 拷贝可执行文件到原来的目录
cp nginx /usr/local/nginx/sbin
```

第三步：发送信号 USR2 给 Nginx 的 1.14.2 版本对应的 master 进程

```sh
kill -USR2 `cat /usr/local/logs/nginx.pid`
```

第四步：发送信号 QUIT 给 Nginx 的 1.14.2 版本对应的 master 进程

```sh
kill -QUIT `cat /usr/local/logs/nginx.pid.oldbin`
```

##### 安装目录的make命令完成升级

第一步：将 1.14.2 版本的 sbin 目录下的 nginx 进行备份

> 不是复制一份，是直接修改原来的 nginx。

```sh
# 进入 sbin 目录下
cd /usr/local/nginx/sbin

# 备份为 nginxold 文件
mv nginx nginx.backup
```

第二步：将 Nginx1.16.1 安装目录编译后的 objs 目录下的 nginx 文件，拷贝到原来 `/usr/local/nginx/sbin` 目录下

```sh
# 进入 objs 目录
cd ~/nginx/core/nginx-1.16.1/objs

# 拷贝可执行文件到原来的目录
cp nginx /usr/local/nginx/sbin
```

第三步：进入到安装目录，执行 `make upgrade`

```sh
make upgrade
```

> ![image-20250907133614446](.\asset\images\image-20250907133614446.png)

第四步：查看是否更新成功

```sh
nginx -v
```

在整个过程中，其实 Nginx 是一直对外提供服务的。并且当 Nginx 的服务器启动成功后，我们是可以通过浏览器进行直接访问的，同时我们可以通过更改 html 目录下的页面来修改我们在页面上所看到的内容，那么问题来了，为什么我们要修改 html 目录下的文件，能不能多添加一些页面是 Nginx 的功能更加丰富，还有前面聊到 Nginx 的前端功能又是如何来实现的，这就需要我们对 [Nginx 的核心配置文件](https://frxcat.fun/middleware/Nginx/Nginx_Configuration_file) 进行一个详细的学习。



## nginx.conf

### 配置文件内容

读取 Nginx 自带的 Nginx 配置文件，配置文件内容很多，我们先将其中的注释部分【学习一个技术点就是在 Nginx 的配置文件中可以使用 `#` 来注释】删除掉后，就剩下如下内容：

```nginx
worker_processes  1;   # 使用指令 1 

events {    # 这是 events 块
    worker_connections  1024;
}

http {    # 这是 http 块
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;   # 监听 80 端口
        server_name  localhost;   # 监听请求过来的 IP
        location / {   # 请求的地址是 /，则进入这个配置，访问 idnex.html
            root   html;      # 进入 html 目录找到访问的页面
            index  index.html index.htm;
        }
        # 如果访问的页面是 500 502 503 504，则发送 /50x.html 请求
        error_page   500 502 503 504  /50x.html;  
        location = /50x.html {      # 如果匹配上 /50x.html 请求
            root   html;     # 则进入 html 目录找到 /50x.html
        }
    }
}
```

对上面文件内容的解释，一一对应比较解释：



```nginx
指令名	指令值;  # 全局块，主要设置 Nginx 服务器整体运行的配置指令

# events 块，主要设置 Nginx 服务器与用户的网络连接,这一部分对 Nginx 服务器的性能影响较大
events {	 
    指令名	指令值;
}

# http 块，是 Nginx 服务器配置中的重要部分，代理、缓存、日志记录、第三方模块配置...             
http {		
    指令名	指令值;

    server { 		# server 块，是 Nginx 配置和虚拟主机相关的内容
        指令名	指令值;
        location / {      # location 块，基于 Nginx 服务器接收请求字符串与 location 后面的值进行匹配，对特定请求进行处理
            指令名	指令值;
        }
    }
    ...
}
```

小结

nginx.conf 配置文件中默认有三大块：全局块、events 块、http 块

http 块中可以配置多个 server 块，每个 server 块又可以配置多个 location 块。

#### 全局块

全局块的配置影响 Nginx 的全局设置。如用户权限，启动的进程数等。

##### user指令

1. **user：用于配置运行 Nginx 服务器的 worker 进程的用户和用户组。**

| 语法                | 默认值 | 位置   |
| ------------------- | ------ | ------ |
| user <user> [group] | nobody | 全局块 |

该属性也可以在编译的时候指定，语法如下 `./configure --user=user --group=group`，如果两个地方都进行了设置，最终生效的是配置文件中的配置。

该指令的使用步骤:

1. 进入配置文件添加一个用户信息 『 www 』

```sh
user www
```

测试进行测试配置文件会报错：

```sh
nginx -t
```

```sh
[root@master conf]# nginx -t
nginx: [emerg] getpwnam("www") failed in /usr/local/nginx/conf/nginx.conf:1
nginx: configuration file /usr/local/nginx/conf/nginx.conf test failed
```

原因在于 Linux 系统不存在 www 用户，我们需要创建它。

创建一个用户

```sh
useradd www
```

重启 Nginx 的配置文件

```sh
nginx -s reload

# 查看重启是否生效
ps -ef | grep nginx
```

最后返回的结果由 root 用户改为 www 用户，代表配置成功。

```sh
[root@master conf]# ps -ef | grep nginx
root       8960      1  0 16:13 ?        00:00:00 nginx: master process ./nginx
www       11975   8960  0 20:44 ?        00:00:00 nginx: worker process
root      11978  10615  0 20:44 pts/1    00:00:00 grep --color=auto nginx
```

1. 在 Linux 的 `/root` 下创建一个 html 目录，并且进入 html 目录，创建 index.html 文件

```sh
mkdir /root/html

cd /root/html

vim index.html
```

然后在 `/root/html/index.html` 文件里添加如下内容：

```html
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
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
<p><em>I am WWW</em></p>
</body>
</html>
```

这些内容其实就是在 Nginx 的欢迎页面上多加别 `I am WWW` 内容。

1. 修改 nginx.conf

```nginx
location / {
	# root   html;  # 原始的代码注释掉
	root   /root/html;    # 不再是 html 目录，而是 root 下的 html 目录
	index  index.html index.htm;  # 就是上方创建的 index.html
}
```

1. `nginx -s reload`重新加载后，测试启动访问

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220728/image.6dtnfz58zpo0.webp)

页面会报 403 拒绝访问的错误。

1. 分析原因：因为当前用户（www 用户）没有访问 `/root/html` 目录的权限，这个目录只有 root 才能访问。

那么 www 如何访问我们写的 index.html 页面呢？我们知道，每新建一个用户，`/home` 下都会生成该用户权限的目录。

1. 将文件移动到 `/home/www/html/index.html`

```sh
mv /root/html /home/www
```

1. 记得修改配置文件的资源内容

```nginx
location / {
	# root   html;  # 原始的代码注释掉
	# root   /root/html;   # 这属于 root 权限的页面，注释或者删除掉
	root   /home/www/html;  # 这是 www 用户有权限访问的目录
	index  index.html index.htm;   # 访问了 html 目录，然后访问 index.html 文件
}
```

1. 再次测试启动访问，可以正常访问。

综上所述，使用 user 指令可以指定启动运行工作进程的用户及用户组，这样对于系统的权限访问控制的更加精细，也更加安全。

我们也能理解了配置文件的 `locaotion` 块的基本使用，root 对应着访问目录，index 对应着访问目录下的默认页面。

##### work process指令

1. `master_process` 指令用来指定是否开启 worker 工作进程。

如果为 off，则代表关闭了 worker 进程，这时候启动 Nginx，只有 master 进程启动，没有 worker 进程。默认开启 worker 工作进程。(需要重启nginx服务生效)

| 语法                        | 默认值             | 位置   |
| --------------------------- | ------------------ | ------ |
| master_process <on \| off>; | master_process on; | 全局块 |

`worker_processes` 指令用于配置 Nginx 生成 worker 工作进程的数量，这个是 Nginx 服务器实现并发处理

服务的关键所在。

理论上来说 `workder process` 的值越大，可以支持的并发处理量也越多，但事实上这个值的设定是需要受到来自服务器自身的限制，建议将该值和服务器 CPU 的内核数保存一致。

| 语法                            | 默认值 | 位置   |
| ------------------------------- | ------ | ------ |
| worker_processes <num \| auto>; | 1      | 全局块 |

如果将 `worker_processes` 设置成 2，则会看到如下内容:

```sh
[root@master conf]# ps -ef | grep nginx
root       8960      1  0 16:13 ?        00:00:00 nginx: master process ./nginx
www       12299   8960  0 21:14 ?        00:00:00 nginx: worker process
www       12300   8960  0 21:14 ?        00:00:00 nginx: worker process
root      12302  10615  0 21:14 pts/1    00:00:00 grep --color=auto nginx
```

出现两个 worker 工作进程。

##### 其他指令

- `daemon` 指令设置 Nginx 是否以守护进程的方式启动。on 代表开启守护进程，off 代表关闭守护进程，默认开启。(需要重启nginx生效)

守护式进程是 Linux 后台执行的一种服务进程，特点是 **独立于控制终端，不会随着终端关闭而停止**，也就是后台启动。

| 语法                | 默认值     | 位置   |
| ------------------- | ---------- | ------ |
| daemon <on \| off>; | daemon on; | 全局块 |

- `pid` 指令用来配置 Nginx 当前 master 进程的进程号 ID 存储的文件路径。默认路径是 `/usr/local/nginx/logs/nginx.pid`。

| 语法        | 默认值                          | 位置   |
| ----------- | ------------------------------- | ------ |
| pid <file>; | /usr/local/nginx/logs/nginx.pid | 全局块 |



该属性也可以通过 `./configure --pid-path=PATH` 在编译时来指定。

- `error_log` 指令用来配置 Nginx 的错误日志存放路径。默认路径是 `/usr/local/nginx/logs/error.log`。

| 语法                         | 默认值                          | 位置                           |
| ---------------------------- | ------------------------------- | ------------------------------ |
| error_log <file> [日志级别]; | error_log logs/error.log error; | 全局块、http、server、location |



该属性也可以通过 `./configure --error-log-path=PATH` 在编译时来指定。

其中日志级别的值有『 debug | info | notice | warn | error | crit | alert | emerg 』，翻译过来为「调试 | 信息 | 通知 | 警告 | 错误 | 临界 | 警报 | 紧急」，这块建议大家设置的时候不要设置成 info 以下的等级，因为会带来大量的磁盘 I/O 消耗，影响 Nginx 的性能。

- `include` 指令用来引入其他的配置文件，使 Nginx 的配置更加灵活。

| 语法            | 默认值 | 位置 |
| --------------- | ------ | ---- |
| include <file>; | 无     | any  |



#### events块

##### 指令

- `accept_mutex` 指令用来设置是否开启 Nginx 网络连接序列化。默认开启。

| 语法                      | 默认值           | 位置   |
| ------------------------- | ---------------- | ------ |
| accept_mutex <on \| off>; | accept_mutex on; | events |



这个配置主要可以用来解决常说的「惊群」问题。

大致意思是在某一个时刻，客户端发来一个请求连接，Nginx 后台是以多进程的工作模式，也就是说有多个 worker 进程会被同时唤醒，但是最终只会有一个进程可以获取到连接，如果每次唤醒的进程数目太多，就会影响 Nginx 的整体性能。

如果将上述值设置为 on (开启状态)，将会对多个 Nginx 进程接收连接进行序列号，一个个来唤醒接收，就防止了多个进程对连接的争抢。

如图的小狗，如果只是一块「骨头」出现，则只需要唤醒一个小狗即可（开启 on），如果多个「骨头」如三个同时出现，那么唤醒三个小狗效率更高（此时需要设置 off）



- `multi_accept` 指令用来设置是否开启同时接收多个网络连接。默认开启。

| 语法                      | 默认值            | 位置   |
| ------------------------- | ----------------- | ------ |
| multi_accept <on \| off>; | multi_accept off; | events |



如果 multi_accept 被禁止了，Nginx 的一个工作进程只能同时接受一个新的连接。如果开启，一个工作进程可以同时接受所有的新连接。**建议开启。**



- `worker_connections` 指令用来配置单个 worker 进程最大的连接数。默认 512 个连接数。

| 语法                         | 默认值                  | 位置   |
| ---------------------------- | ----------------------- | ------ |
| worker_connections <number>; | worker_commections 512; | events |



这里的连接数不仅仅包括和前端用户建立的连接数，而是包括所有可能的连接数。另外，number 值不能大于操作系统支持打开的最大文件句柄数量。

- `use` 指令用来设置 Nginx 服务器选择哪种事件驱动来处理网络消息。

| 语法          | 默认值           | 位置   |
| ------------- | ---------------- | ------ |
| use <method>; | 根据操作系统规定 | events |



注意：此处所选择事件处理模型是 Nginx 优化部分的一个重要内容，method 的可选值有『 select | poll | epoll | kqueue 』等，之前在准备 Centos 环境的时候，我们强调过要使用 Linux 内核在 2.6 以上，就是为了能使用 epoll 函数来优化 Nginx。

另外这些值的选择，我们也可以在编译的时候使用 `--with-select_module`、`--without-select_module`、`--with-poll_module`、`--without-poll_module` 来设置是否需要将对应的事件驱动模块编译到 Nginx 的内核。

##### events指令配置模板

打开 Nginx 的配置文件 nginx.conf，添加如下配置



```nginx
events{
	accept_mutex on;    # 开启 Nginx 网络连接序列化
	multi_accept on;    # 开启同时接收多个网络连接
	worker_connections 1024;   # 单个 worker 进程最大的连接数
	use epoll;   # 使用 epoll 函数来优化 Nginx
}
```

启动测试



```sh
# 测试配置是否语法出错
nginx -t
# 重新加载 Nginx
nginx -s reload
```

#### http块

##### 定义MIME-Type

我们都知道浏览器中可以显示的内容有 HTML、XML、GIF 等种类繁多的文件、媒体等资源，浏览器为了区分这些资源，就需要使用 MIME Type。

所以说 MIME Type 是网络资源的媒体类型。Nginx 作为 Web 服务器，也需要能够识别前端请求的资源类型。

在 Nginx 的配置文件中，默认有两行配置：

```sh
include mime.types;      # 引入 mime.types 文件的内容
default_type application/octet-stream;     # 默认的 MIME 类型
```

- `default_type` 指令用来配置 Nginx 响应前端请求默认的 MIME 类型。默认是 text 文本。

| 语法                      | 默认值                   | 位置                   |
| ------------------------- | ------------------------ | ---------------------- |
| default_type <mime-type>; | default_type text/plain; | http、server、location |



在 `default_type` 前面还有一句 `include mime.types`，include 之前我们已经介绍过，相当于把 mime.types 文件中 MIMT 类型与相关类型文件的文件后缀名的对应关系加入到当前的配置文件中。

举例来说明：

有些时候请求某些接口的时候需要返回指定的文本字符串或者 json 字符串，而不是页面，如果逻辑非常简单或者干脆是固定的字符串，那么可以使用 Nginx 快速实现，这样就不用编写程序响应请求了，可以减少服务器资源占用并且响应性能非常快。

如何实现：



```nginx
location /get_text {
    default_type text/html;      # 等价于 text/plain，返回文本类型
    return 200 "<h1>This is nginx's text</h1>";
}
location /get_json{
    default_type application/json;   # 返回 json 字符串类型
    return 200 '{"name": "xu", "age": 21}';
}
```

##### 自定义服务日志

Nginx 中日志的类型分 access.log、error.log。

`access.log` 日志用来记录用户所有的访问请求。

`error.log` 日志记录 Nginx 本身运行时的错误信息，不会记录用户的访问请求。

Nginx 服务器支持对服务日志的格式、大小、输出等进行设置，需要使用到两个指令，分别是 `access_log` 和 `log_format` 指令。

1. `access_log` 指令用来设置用户访问日志的相关属性。

| 语法                                     | 默认值                               | 位置                   |
| ---------------------------------------- | ------------------------------------ | ---------------------- |
| access_log <path> [format[buffer=size]]; | access_log logs/access.log combined; | http、server、location |



format 对应着 `log_format` 的 name，必须保持一致。

1. `log_format` 指令用来指定日志的输出格式。

| 语法                                                         | 默认值                     | 位置 |
| ------------------------------------------------------------ | -------------------------- | ---- |
| log_format <name> [escape=default \| json \| none] <string> ...... ; | log_format combined "..."; | http |



name 对用 `access_log` 的 format，必须保持一致。

> **例子 1：自定义日志路径和输出格式**

- 在 `/usr/local/nginx/logs` 下创建 my.log 文件，该文件作为日志。



```sh
mkdir /usr/local/nginx/logs/my.log
```

- 自定义日志输出格式：`==========>This is My format`
- 在配置文件配置相关指令



```nginx
log_format myformat '=========>This is My format';
access_log logs/my.log myformat;
```

- 重启服务并进行测试



```sh
# 重启 Nginx 服务
nginx -s reload

# 监听日志
tail -f /usr/local/nginx/logs/my.log
```

浏览器访问一次 Nginx 的欢迎页面，回来看日志的输出，结果如图：



> **例子 2：输出内容加上访问机器的信息**

- 进入配置文件，在输出格式上加上 Nginx 的内置参数

```nginx
log_format myformat '=========>This is My format:$http_user_agent';
access_log logs/my.log myformat;
```

- 重启测试

```sh
# 重启 Nginx 服务
nginx -s reload

# 监听日志
tail -f /usr/local/nginx/logs/my.log
```



##### 其他配置指令

- `sendfile`：用来设置 Nginx 服务器是否使用 sendfile 传输文件，该属性可以大大提高 Nginx 处理静态资源的性能。默认关闭，建议开启。

| 语法                  | 默认值        | 位置                   |
| --------------------- | ------------- | ---------------------- |
| sendfile <on \| off>; | sendfile off; | http、server、location |

- `keepalive_timeout`：用来设置长连接的超时时间，默认超时时间是 75 秒。

> **为什么要使用 keepalive？**
>
> 我们都知道 HTTP 是一种无状态协议，客户端向服务端发送一个 TCP 请求，服务端响应完毕后断开连接。
>
> 如何客户端向服务端发送多个请求，每个请求都需要重新创建一次连接，效率相对来说比较多，使用 keepalive 模式，可以告诉服务器端在处理完一个请求s后保持这个 TCP 连接的打开状态，若接收到来自这个客户端的其他请求，服务端就会利用这个未被关闭的连接，而不需要重新创建一个新连接，提升效率，但是这个连接也不能一直保持，这样的话，连接如果过多，也会是服务端的性能下降，这个时候就需要我们进行设置其的超时时间。

| 语法                      | 默认值                 | 位置                   |
| ------------------------- | ---------------------- | ---------------------- |
| keepalive_timeout <time>; | keepalive_timeout 75s; | http、server、location |

- `keepalive_requests`：用来设置一个 keep-alive 连接使用的次数，默认是 100 次。

| 语法                         | 默认值                  | 位置                   |
| ---------------------------- | ----------------------- | ---------------------- |
| keepalive_requests <number>; | keepalive_requests 100; | http、server、location |



##### server块和location块

server 块和 location 块都是我们要重点学习的内容，因为我们后面会对 Nginx 的功能进行详细讲解，所以该内容在[静态资源部署](https://frxcat.fun/middleware/Nginx/Nginx_Static_resource_deployment/)和[静态资源访问](https://frxcat.fun/middleware/Nginx/NginxStatic_resource_access/)进行详细说明。

本次我们这是认识下 Nginx 默认给的 nginx.conf 中的相关内容，以及 server 块与 location 块在使用的时候需要注意的一些内容。



```nginx
server {
	listen       80;     # 监听 80 端口，如果更改端口，则外界访问的时候带上对应的端口号，如 8080
	server_name  localhost;  # 指定可以访问 Nginx 的 IP 地址
	
	location / {      
		root   html;    # 访问资源所对应的目录，这里是 html 目录
		index  index.html index.htm;    # 访问资源所对应目录下的默认页面，优先级递增
	}

	error_page   500 502 503 504 404  /50x.html;  # 访问错误，跳转访问 /50x.html 请求
	location = /50x.html {   # 访问 /50x.html 请求的处理
		root   html;    # 访问资源所对应的目录，这里是 html 目录的 50x.html
	}
}
```

- listen 和 server_name 是我们的 `http://server_name:listen`，如 `http://localhost:80`
- location / 就是访问 `http://server_name:listen/`，里面的配置对应着 `http://server_name:listen/html/index.html`
- 页面产生 500 502 503 504 404，就会发送 `http://server_name:listen/50x.html`
- location = /50x.html 就是 `http://server_name:listen/50x.html`，它会自动访问 `http://server_name:listen/html/50x.html`
- root 代表资源目录指令
- index 代表默认访问网页指令



> 1、nginx.conf文件详解
>
> ### 一、配置文件的核心结构
>
> Nginx 的配置文件采用一种**分层的、基于指令块**的结构。它可以被看作一个树形结构，从主到次，从全局到具体。
>
> 默认的 `nginx.conf` 通常包含以下几个核心块：
>
> nginx
>
> ```yaml
> # 全局块 (Main Context)
> user nginx nginx;
> worker_processes auto;
> error_log /var/log/nginx/error.log warn;
> pid /var/run/nginx.pid;
> 
> # Events 块 (Events Context)
> events {
>     worker_connections 1024;
>     use epoll; # Linux 高效网络模型
> }
> 
> # HTTP 块 (HTTP Context) - 这是配置最丰富的部分
> http {
>     # HTTP 全局配置
>     include /etc/nginx/mime.types;
>     default_type application/octet-stream;
>     access_log /var/log/nginx/access.log combined;
> 
>     # Server 块 (Server Context) - 定义虚拟主机
>     server {
>         # 监听端口和域名
>         listen 80;
>         server_name example.com www.example.com;
> 
>         # Location 块 (Location Context) - 定义URI匹配规则
>         location / {
>             root /usr/share/nginx/html;
>             index index.html index.htm;
>         }
> 
>         location /api/ {
>             proxy_pass http://backend_server;
>         }
>     }
> 
>     # 可以定义多个 Server 块
>     server {
>         listen 80 default_server;
>         server_name _;
>         return 404;
>     }
> 
>     # Upstream 块 (Upstream Context) - 定义上游服务器组（负载均衡）
>     upstream backend_server {
>         server 10.0.1.101:8080 weight=3;
>         server 10.0.1.102:8080;
>         server 10.0.1.103:8080 backup;
>     }
> }
> ```
>
> 
>
> ------
>
> ### 二、各模块详细解析
>
> #### 1. 全局块 (Main Context)
>
> 这是配置文件的最高层级，设置的指令将影响 Nginx 的整体运行。
>
> - `user nginx nginx;`: 指定 Nginx worker 进程运行的用户和用户组。出于安全考虑，不应使用 root。
> - `worker_processes auto;`: 定义 worker 进程的数量。设置为 `auto` 会让 Nginx 自动设置为与 CPU 核心数相同，这是最佳实践。
> - `error_log /var/log/nginx/error.log warn;`: 错误日志文件路径和记录级别（`debug | info | notice | warn | error | crit`）。生产环境通常设为 `warn` 或 `error`。
> - `pid /var/run/nginx.pid;`: 指定存储主进程 PID 的文件位置。
> - `worker_rlimit_nofile 65535;`: 一个 worker 进程能打开的最大文件描述符数。对于高流量站点，需要调高此值（需同时调整系统 `ulimit`）。
>
> #### 2. Events 块
>
> 设置影响 Nginx 如何处理客户端连接的参数。
>
> - `worker_connections 1024;`: **每个 worker 进程**可以同时处理的最大连接数。最大客户端数 = `worker_processes * worker_connections`。
> - `use epoll;`: 指定高效的事件驱动模型。在 Linux 上，`epoll` 是最佳选择，Nginx 通常会自动选择最佳模型，无需手动设置。
> - `multi_accept on;`: 设置一个 worker 进程是否同时接受多个新连接。
>
> #### 3. HTTP 块
>
> 这是配置的**核心**，所有 HTTP 相关的配置都在此模块中。它包含了众多的指令，并且可以包含多个 `server` 块。
>
> **常见 HTTP 全局指令：**
>
> - `include /etc/nginx/mime.types;`: 引入 MIME 类型映射文件，使 Nginx 能正确返回文件类型（如 `text/html`, `image/jpeg`）。
> - `default_type application/octet-stream;`: 默认的 MIME 类型，当无法从文件扩展名判断类型时使用。
> - `access_log /var/log/nginx/access.log combined;`: 定义访问日志的路径和格式。`combined` 是一种预定义的较详细的日志格式。
> - `sendfile on;`: 开启高效文件传输模式，减少文件在用户和内核空间之间的拷贝，提升性能。
> - `tcp_nopush on;`: 与 `sendfile on` 一起使用。它允许在同一个数据包中发送响应头和数据体。
> - `tcp_nodelay on;`: 禁用 Nagle 算法，使得小数据包能够立即发送，降低延迟。对于 keepalive 连接很重要。
> - `keepalive_timeout 65;`: 客户端与服务器 keepalive 连接的超时时间（秒）。
> - `gzip on;`: 开启 Gzip 压缩，有效减少传输数据量。
> - `gzip_types text/plain text/css application/json application/javascript;`: 指定需要压缩的 MIME 类型。
>
> #### 4. Server 块 (虚拟主机)
>
> 在一个 HTTP 块内，**可以定义多个 `server` 块来托管多个网站（虚拟主机）**。Nginx 通过 `listen` 和 `server_name` 指令来决定将客户端的请求路由到哪个 `server` 块。
>
> - `listen 80;`: 监听端口，可以是 `80`, `443 ssl`, `8080` 等。
> - `server_name example.com www.example.com;`: 绑定的域名。可以精确匹配、通配符匹配（`*.example.com`）或正则表达式匹配（以 `~` 开头）。`_` 代表匹配所有未明确指定的域名，通常用于返回错误页面。
> - `root /usr/share/nginx/html;`: 该虚拟主机的根目录。后续 `location` 中的路径会基于此目录查找。
> - `index index.html index.htm;`: 定义默认索引文件。
>
> #### 5. Location 块
>
> 在 `server` 块内，`location` 块用于根据请求的 URI 来匹配和处理请求。**这是最灵活、最强大的部分**。
>
> **语法：** `location [修饰符] pattern { ... }`
>
> - **修饰符：**
>   - **(无)**: 前缀匹配。匹配以指定模式**开头**的 URI。
>   - `=`: 精确匹配。URI 必须与模式完全相等。
>   - `~`: 区分大小写的正则匹配。
>   - `~*`: 不区分大小写的正则匹配。
>   - `^~`: 最佳非正则匹配。如果匹配成功，则不再进行正则匹配。
> - **匹配优先级** (从高到低):
>   1. `=` 精确匹配
>   2. `^~` 最佳前缀匹配
>   3. `~` 或 `~*` 正则匹配 (按配置文件中出现的顺序)
>   4. 普通前缀匹配 (最长前缀匹配)
> - **常见指令：**
>   - `root /path/to/files;`: 定义该 location 的根目录（会覆盖 server 级的 root）。
>   - `alias /path/to/files/;`: 定义路径别名。常用于将 location 路径映射到文件系统的另一个路径。**注意末尾的 `/`**。
>   - `proxy_pass http://upstream_name/;`: 将请求反向代理到指定的上游服务器或 URL。
>   - `try_files $uri $uri/ /index.html;`: 按顺序检查文件是否存在，并返回找到的第一个文件。常用于单页面应用（SPA）或友好的错误处理。例如：先找请求的文件 `$uri`，再找对应的目录 `$uri/`，最后都找不到则返回 `/index.html`。
>   - `return 301 https://$host$request_uri;`: 返回特定状态码和 URL，常用于重定向（如 HTTP 跳转到 HTTPS）。
>
> #### 6. Upstream 块
>
> 定义一组后端服务器，用于实现**负载均衡**。通常与 `location` 中的 `proxy_pass` 指令配合使用。
>
> - `server 10.0.1.101:8080;`: 定义后端服务器地址。
> - **参数：**
>   - `weight=5`: 权重，权重越高的服务器接收越多请求。
>   - `max_fails=3`: 最大失败次数，超过后认为服务器不可用。
>   - `fail_timeout=30s`: 服务器被标记为不可用的时长。
>   - `backup`: 备份服务器，只有当其他所有服务器都宕机后才会被启用。
> - **负载均衡算法：**
>   - 默认：加权轮询 (Round Robin)
>   - `ip_hash`: 根据客户端 IP 进行哈希计算，确保同一客户端的请求总是发往同一台服务器（可用于会话保持）。
>   - `least_conn`: 将请求发送给当前连接数最少的服务器。
>
> 2、设置多个server块的意义
>
> ### 1. 基于不同域名的虚拟主机 (最最常见)
>
> 这是 `server` 块最核心的用途。你可以在同一台服务器、同一个端口（如 80）上，托管无数个不同的网站。
>
> **场景**：一台服务器，一个IP，要同时运行 `example.com` 和 `another-site.com`。
>
> nginx
>
> ```yaml
> http {
>     # 第一个 Server 块：处理 example.com 的请求
>     server {
>         listen 80; # 监听 80 端口
>         server_name example.com www.example.com; # 匹配此域名
> 
>         root /var/www/example.com;
>         index index.html;
>         # ... 其他针对 example.com 的配置
>     }
> 
>     # 第二个 Server 块：处理 another-site.com 的请求
>     server {
>         listen 80; # 同样监听 80 端口
>         server_name another-site.com www.another-site.com; # 但匹配另一个域名
> 
>         root /var/www/another-site.com;
>         index index.html;
>         # ... 其他针对 another-site.com 的配置
>     }
> }
> ```
>
> 
>
> 当用户访问 `http://example.com` 时，Nginx 会根据 `Host` 头匹配到第一个 `server` 块，并返回 `/var/www/example.com/index.html`。
> 当用户访问 `http://another-site.com` 时，则会匹配到第二个 `server` 块。
>
> ### 2. 处理不同端口
>
> 虽然不常见，但 `server` 块确实可以用于监听不同的端口，提供完全不同的服务。
>
> **场景**：在同一台机器上，80 端口提供 Web 服务，8080 端口提供一个管理后台或另一个应用。
>
> nginx
>
> ```yaml
> http {
>     # Server 块 A：提供主 Web 服务
>     server {
>         listen 80;
>         server_name example.com;
>         # ... 主网站配置
>     }
> 
>     # Server 块 B：在另一个端口提供管理界面
>     server {
>         listen 8080;
>         server_name example.com;
> 
>         root /var/www/admin-panel;
>         # ... 管理后台配置
>         # 通常还会包含额外的访问控制，如 auth_basic
>     }
> }
> ```
>
> 
>
> ### 3. 设置默认/兜底服务器
>
> 用于处理那些无法匹配任何已知 `server_name` 的请求（比如有人用IP地址访问，或者解析到本机的未知域名）。这能防止他人将任意域名指向你的服务器IP并占用你的资源。
>
> **场景**：拦截所有未知域名的请求，返回错误或重定向。
>
> nginx
>
> ```yaml
> http {
>     # 明确定义一个默认服务器
>     server {
>         listen 80 default_server; # 关键：default_server 参数
>         server_name _; # 使用下划线匹配所有未定义的域名
> 
>         return 444; # Nginx 特有的非标准状态码，直接关闭连接，不发送任何响应体
>         # 或者返回一个自定义错误页
>         # return 404;
>     }
> 
>     server {
>         listen 80;
>         server_name example.com;
>         # ... 正常配置
>     }
> }
> ```
>
> 
>
> ### 4. HTTP 到 HTTPS 的重定向
>
> 这是一个极其重要的用法。创建一个专门监听 80 端口的 `server` 块，其唯一任务就是将所有HTTP流量重定向到安全的HTTPS版本。
>
> **场景**：强制用户使用加密的 HTTPS 连接。
>
> nginx
>
> ```yaml
> http {
>     # 这个 Server 块只负责重定向
>     server {
>         listen 80;
>         server_name example.com www.example.com;
>         return 301 https://$server_name$request_uri; # 301 永久重定向
>     }
> 
>     # 这个 Server 块处理真正的 HTTPS 流量
>     server {
>         listen 443 ssl;
>         server_name example.com www.example.com;
> 
>         ssl_certificate /path/to/cert.crt;
>         ssl_certificate_key /path/to/private.key;
>         # ... HTTPS 网站的其他配置
>         root /var/www/example.com;
>     }
> }
> ```
>
> 
>
> ### 5. 服务于同一域名的不同子域
>
> 为移动端API、管理后台、静态资源等创建独立的子域，并使用不同的 `server` 块进行配置，逻辑清晰，便于管理。
>
> **场景**：为 `api.example.com` 配置反向代理，为 `static.example.com` 配置高效的静态文件服务。
>
> nginx
>
> ```yaml
> http {
>     # 主站
>     server {
>         listen 80;
>         server_name example.com www.example.com;
>         # ... 主站配置
>     }
> 
>     # API 子域 - 反向代理到后端应用
>     server {
>         listen 80;
>         server_name api.example.com;
> 
>         location / {
>             proxy_pass http://localhost:3000;
>         }
>     }
> 
>     # 静态文件子域 - 设置长缓存
>     server {
>         listen 80;
>         server_name static.example.com;
> 
>         location / {
>             root /var/www/static;
>             expires 1y; # 告诉浏览器缓存1年
>             add_header Cache-Control "public, immutable";
>         }
>     }
> }
> ```



## 基础配置实例

前面我们已经对 Nginx 服务器默认配置文件的结构和涉及的基本指令做了详细的阐述。通过这些指令的合理配置，我们就可以让一台 Nginx 服务器正常工作，并且提供基本的 Web 服务器功能。

接下来我们将通过一个比较完整和最简单的基础配置实例，来巩固下前面所学习的指令及其配置。

### 需求

- 有如下访问：

  `http://192.168.199.27:8081/server1/location1` 访问的是：index_sr1_location1.html `http://192.168.199.27:8081/server1/location2` 访问的是：index_sr1_location2.html `http://192.168.199.27:8082/server2/location1` 访问的是：index_sr2_location1.html `http://192.168.199.27:8082/server2/location2` 访问的是：index_sr2_location2.html

- 如果访问的资源不存在，返回自定义的 404 页面

- 将 /server1 和 /server2 的配置使用不同的配置文件分割，将两个文件文件放到 /home/www/conf.d 目录下，然后在 Nginx 的配置文件使用 include 合并两个文件

- 为 /server1 和 /server2 各自创建一个访问日志文件

### 实现

```sh
# 创建 404 页面
touch /home/www/404.html

# 创建 conf.d 目录
mkdir /home/www/conf

# 创建两个配置文件
touch /home/www/conf/server1.conf
touch /home/www/conf/server2.conf

# 创建 myweb 目录
mkdir /home/www/myweb

# 创建 server1 目录和其子目录以及 index.html 文件
mkdir -p /home/www/myweb/server1/location1
mkdir -p /home/www/myweb/server1/location2

touch /home/www/myweb/server1/location1/index.html
touch /home/www/myweb/server1/location2/index.html

# 创建日志目录和日志文件
mkdir -p /home/www/myweb/server1/logs
touch /home/www/myweb/server1/logs/access.log

# 创建 server2 目录和其子目录以及 index.html 文件
# 和创建 server1 步骤一样，把 1 改为 2 即可
```

准备相关文件，/homw/www 目录如下：

因为 Nginx 自带配置文件的备份，即 nginx.conf.default，所以我们可以直接修改配置文件，但是如果你的配置文件曾经修改过，那么请进行备份。

```sh
cp /usr/local/nginx/conf/nginx.conf /usr/local/nginx/conf/nginx.conf.backup
```

备份后，进入 `/usr/local/nginx/conf/nginx.conf` 配置文件

```sh
vim /usr/local/nginx/conf/nginx.conf
```

先清空文件，然后添加如下内容:

```nginx
user www; # 配置允许运行 Nginx 工作进程的用户和用户组
worker_processes 2;  # 配置运行 Nginx 进程生成的 worker 进程数
error_log logs/error.log;  # 配置 Nginx 服务器运行对错误日志存放的路径
pid logs/nginx.pid;   # 配置 Nginx 服务器允许时记录 Nginx 的 master 进程的 PID 文件路径和名称
daemon on;   # 配置 Nginx 服务是否以守护进程方法启动

events{
	accept_mutex on;   # 设置 Nginx 网络连接序列化,解决惊群
	multi_accept on;   # 设置 Nginx 的 worker 进程是否可以同时接收多个请求
	worker_connections 1024;   # 设置 Nginx 的 worker 进程最大的连接数
	use epoll;   # 设置 Nginx 使用的事件驱动模型
}

http{

	include mime.types;   # 定义 MIME-Type
	default_type application/octet-stream;
	sendfile on;   # 配置允许使用 sendfile 方式运输
	keepalive_timeout 65;   # 配置连接超时时间
	
	# 配置请求处理日志格式
	log_format server1 '===>server1 access log';
	log_format server2 '===>server2 access log';
	
	include /home/www/conf/*.conf;  # 引用其他 conf 文件
}
```

第 25 行代码使用 include 将 service1 和service2 的配置文件进行引用。以后无需修改主配置文件，只需要引入子配置文件即可，主配置文件作为默认值，子配置文件的内容会覆盖和主配置文件相同的内容。

进入 server1.conf 文件



```sh
vim /home/www/conf/server1.conf
```

server1.conf 文件内容：

```nginx
server{
  listen 8081;   # 配置监听端口和主机名称
  server_name localhost;
  access_log /home/www/myweb/server1/logs/access.log server1;   # 配置请求处理日志存放路径
  error_page 404 /404.html;   # 配置错误页面

  location /server1/location1{   # 配置处理 /server1/location1 请求的 location
      root /home/www/myweb;
      index index.html;       # 这是 server1 下的 location1 的 index.html
  }

  location /server1/location2{   # 配置处理 /server1/location2 请求的 location
      root /home/www/myweb;
      index index.html;    # 这是 server1 下的 location2 的 index.html
  }

  location = /404.html {   # 配置错误页面转向
      root /home/www;
      index 404.html;
  }
}
```



server2.conf 文件内容：

```nginx
server{
  listen 8082;   # 配置监听端口和主机名称
  server_name localhost;
  access_log /home/www/myweb/server2/logs/access.log server2;   # 配置请求处理日志存放路径
  error_page 404 /404.html;   # 配置错误页面,对404.html做了定向配置
  
  location /server2/location1{   # 配置处理 /server1/location1 请求的 location
      root /home/www/myweb;
      index index.html;   # 这是 server2 下的 location1 的 index.html
  }
 
  location /server2/location2{   # 配置处理 /server2/location2 请求的 location
      root /home/www/myweb;
      index index.html;    # 这是 server2 下的 location2 的 index.html
  }
  
  location = /404.html {   # 配置错误页面转向
      root /home/www;
      index 404.html;
  }
}
```

server1下面的location1下面的index.html的内容

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
</head>
<body>
	<h1>server1下面的loaction1下面的index.html</h1>
</body>
</html>
```

其他的三个页面把数字改了就可以

404.html内容;

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
</head>
<body>
	<h1>不好意思，程序小哥正在加紧维修中 ...... </h1>
</body>
</html>
```

重启 Nginx ，使得配置文件生效

```sh
nginx -s reload
```

### 测试

如果没有关闭防火墙，记得开放 8081 和 8082 端口。



```sh
# 开放 8081 和 8082 端口
firewall-cmd --zone=public --add-port=8081/tcp --permanent
firewall-cmd --zone=public --add-port=8082/tcp --permanent

# 重启防火墙
firewall-cmd --reload
```



## 操作的问题

经过前面的操作，我们会发现，如果想要启动、关闭或重新加载 Nginx 配置文件，都需要先进入到 Nginx 的安装目录的 sbin 目录，然后使用 Nginx 的二级制可执行文件 nginx 来操作，相对来说操作比较繁琐，这块该如何优化？另外如果我们想把 Nginx 设置成随着服务器启动就自动完成启动操作，又该如何来实现？

这就需要用到接下来我们要讲解的两个知识点：

- Nginx 服务启停配置
- Nginx 全局命令配置

## 服务启停配置

把 Nginx 应用服务设置成为系统服务，方便对 Nginx 服务的启动和停止等相关操作，具体实现步骤:

- 在 `/usr/lib/systemd/system` 目录下创建 nginx.service 文件

```sh
vim /usr/lib/systemd/system/nginx.service
```

文件添加如下内容：

```sh
[Unit]
Description=nginx web service
Documentation=http://nginx.org/en/docs/
After=network.target

[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx
ExecReload=/usr/local/nginx/sbin/nginx -s reload
ExecStop=/usr/local/nginx/sbin/nginx -s stop
PrivateTmp=true

[Install]
WantedBy=default.target
```

注意：可执行文件 nginx 根据自己的路径进行修改，以及 .conf 配置文件和 .pid 文件的路径。这份内容是基于默认安装目录的。

- 添加完成后，如果权限有问题需要进行权限设置，没有则忽略这一步

```sh
chmod 755 /usr/lib/systemd/system/nginx.service
```

- 使用系统命令来操作 Nginx 服务

```sh
# 启动 Nginx
systemctl start nginx

# 停止 Nginx
systemctl stop nginx

# 重启 Nginx
systemctl restart nginx

# 重新加载配置文件
systemctl reload nginx

# 查看 Nginx 状态
systemctl status nginx

# 开机启动
systemctl enable nginx

# 关闭开启启动
systemctl disable nginx
```

## 全局命令配置

前面我们介绍过 Nginx 安装目录下的二级制可执行文件 `nginx` 的很多命令，要想使用这些命令前提是需要进入 sbin 目录下才能使用，很不方便，如何去优化，我们可以将该二进制可执行文件加入到系统的环境变量，这样的话在任何目录都可以使用 nginx 对应的相关命令。具体实现步骤如下:

方法一：

- 修改 `/etc/profile` 文件

```sh
vim /etc/profile

# 在最后一行添加
export PATH=$PATH:/usr/local/nginx/sbin
```

可执行文件 nginx 的路径根据自己的路径修改，这里是默认路径。

- 使之立即生效

```sh
source /etc/profile
```

- 任意位置执行 nginx 命令，测试成功

```sh
[root@master ~]# nginx -v
nginx version: nginx/1.21.6
```

方法二：

- 将可执行文件 nginx 拷贝一份到 /usr/bin 目录下

```sh
cp /usr/local/nginx/sbin/nginx /usr/bin
```

- 任意位置执行 nginx 命令，测试成功

```sh
[root@master ~]# nginx -v
nginx version: nginx/1.21.6
```





## Nginx静态资源

上网去搜索访问资源对于我们来说并不陌生，通过浏览器发送一个 HTTP 请求实现从客户端发送请求到服务器端获取所需要内容后并把内容回显展示在页面的一个过程。这个时候，我们所请求的内容就分为两种类型，一类是静态资源、一类是动态资源。

静态资源即指在服务器端真实存在并且能直接拿来展示的一些文件，比如常见的 html 页面、css 文件、js 文件、图片、视频等资源；

动态资源即指在服务器端真实存在但是要想获取需要经过一定的业务逻辑处理，根据不同的条件展示在页面不同这 一部分内容，比如说报表数据展示、根据当前登录用户展示相关具体数据等资源；

Nginx 处理静态资源的内容，我们需要考虑下面这几个问题：

- 静态资源的配置指令
- 静态资源的配置优化
- 静态资源的压缩配置指令
- 静态资源的缓存处理
- 静态资源的访问控制，包括跨域问题和防盗链问题

### listen指令

该指令是用来配置监听端口。默认监听 80（root 启动 Nginx） 和 8000（非 root 启动 Nginx） 端口。

| 语法                                                         | 默认值                | 位置   |
| ------------------------------------------------------------ | --------------------- | ------ |
| listen <address>[:port] [default_server] ...... ; listen <port> [default_server] ...... ; | listen *:80 \| *:8000 | server |



- `listen` 指令的设置比较灵活，我们通过几个例子来把常用的设置方式熟悉下：



```sh
listen 127.0.0.1:8000; # listen localhost:8000 监听指定的IP和端口
listen 127.0.0.1;	# 监听指定IP的所有端口
listen 8000;	# 监听指定端口上的连接
listen *:8000;	# 监听指定端口上的连接
```

- `default_server` 属性是标识符，用来将此虚拟主机设置成默认主机。

所谓的默认主机指的是如果没有匹配到对应的 `address:port`，则会执行默认的 server。

如果不指定该标识符，又没有匹配到对应的 `address:port` 时，默认使用的是第一个 server，所以第一个 server 要好好设置，建议第一个 server 就加上 `default_server`。



```nginx
server{
	listen 8080;
	server_name 127.0.0.1;
	location / {
		root html;
		index index.html;
	}
}
server{
	listen 8080 default_server;
	server_name localhost;
	default_type text/plain;
	return 444 'This is a error request';
}
```

此时访问 8080 端口，它会访问第二个 server，如果第二个 server 去掉 `default_server` ，则默认访问第一个 server。

### server_name指令

该指令用来设置虚拟主机服务名称。默认为空。

比如 127.0.0.1、localhost、域名[www.baidu.com | www.jd.com]。

| 语法                        | 默认值          | 位置   |
| --------------------------- | --------------- | ------ |
| server_name <name> ...... ; | server_name ""; | server |

- name 可以提供多个中间用空格分隔。

关于 server_name 的配置方式有三种，分别是：

- 精确匹配
- 通配符匹配
- 正则表达式匹配

> **配置方式一：精确匹配**

如：

```nginx
server {
	listen 80;
	server_name www.frx.com www.bing.com;
	...
}
```

此时以 `www.frx.com` 或者 `www.bing.com` 域名进行访问，就会跳转到 Nginx 的欢迎页面，前提是你需要拥有该域名，并且该域名和 Nginx 所在的系统 IP 进行绑定。

所以我可以利用 hosts 文件进行「模拟」域名。

> 补充小知识点
>
> hosts 是一个没有扩展名的系统文件，可以用记事本等工具打开，其作用就是将一些常用的网址域名与其对应的IP地址建立一个关联「数据库」，当用户在浏览器中输入一个需要登录的网址时，系统会首先自动从 hosts 文件中寻找对应的IP地址，一旦找到，系统会立即打开对应网页，如果没有找到，则系统会再将网址提交 DNS 域名解析服务器进行 IP 地址的解析。
>
> hosts 文件不同系统的位置：
>
> - 在 windows 的位置：C:\Windows\System32\drivers\etc
>
> - 在 centos 的位置：/etc/hosts
>
>   因为域名是要收取一定的费用，所以我们可以使用修改 hosts 文件来制作一些虚拟域名来使用。需要修改 `/etc/hosts` 文件来添加
>
>   ```
>   # 进入 hosts 文件
>   vim /etc/hosts
>               
>   # 添加内容
>   127.0.0.1 www.frx.com
>   127.0.0.1 www.bing.com
>   ```

> **配置方式二：使用通配符配置**

server_name 指令支持通配符 *，**但需要注意的是通配符不能出现在域名的中间，只能出现在首段或尾段**，如：

```nginx
server {
	listen 80;
	server_name  *.frx.com	www.frx.*;
	# www.frx.cn abc.frx.cn www.bing.cn www.frx.com
	...
}
```

>  下面的配置就会报错，因为 * 不能出现在域名的中间和与其他字符串联使用
>
> ```
> server {
> 	listen 80;
> 	server_name  www.*.cn www.frx.c*
> 	...
> }
> 
> ```
>
> 

```nginx
server {
	listen 80;
	server_name  .frx.com
	...
}
.frx.com` 相当于 `*..frx.com` + `frx.com
```

> **配置三：使用正则表达式配置**

server_name 指令可以使用正则表达式，**并且使用 `~` 作为正则表达式字符串的开始标记。**

常见的正则表达式：

| 代码  | 说明                                                       |
| ----- | ---------------------------------------------------------- |
| ^     | 匹配搜索字符串开始位置                                     |
| $     | 匹配搜索字符串结束位置                                     |
| .     | 匹配除换行符 \n 之外的任何单个字符                         |
| \     | 转义字符，将下一个字符标记为特殊字符                       |
| [xyz] | 字符集，与任意一个指定字符匹配                             |
| [a-z] | 字符范围，匹配指定范围内的任何字符                         |
| \w    | 与以下任意字符匹配 A-Z a-z 0-9 和下划线,等效于[A-Za-z0-9_] |
| \d    | 数字字符匹配，等效于[0-9]                                  |
| {n}   | 正好匹配 n 次                                              |
| {n,}  | 至少匹配 n 次                                              |
| {n,m} | 匹配至少 n 次至多 m 次                                     |
| *     | 零次或多次，等效于{0,}                                     |
| +     | 一次或多次，等效于{1,}                                     |
| ?     | 零次或一次，等效于{0,1}                                    |

配置如下：

```nginx
server{
	listen 80;
    server_name ~^www\.\w+\.com$;
    default_type text/plain;
    return 200 $1;
}
```

**注意 ~ 后面不能加空格**。

> 同理 这里也能够使用组进行捕获
>
> ```nginx
> server{
> 	listen 80;
>     server_name ~^www\.(\w+)\.com$;
>     default_type text/plain;
>     return 200 $1;
> }
> ```
>
> 括号代表可以在 Ngxin 配置文件内获取其中的值，如上方的 (\w+) 的内容可以用 $1 获取到参数，如果有多个括号，依次使用 $2 $3 ...... 获取。
>
> 比如现在访问 `http://www.frx.com`，则返回 frx 到页面上，因为frx 被放在 $1 处：`return 200 frx`。

这里进行说明：server_name 配置了 localhost 和配置 IP 或者域名的区别

localhost 是「虚拟 IP」，如果不是本机访问，而是外界访问，那么这个就是无效的，但是为什么学习阶段的时候都不改呢，因为当 Nginx 没有匹配到指定的 `server_name`，默认找到第一个 server 块，而 Nginx 默认的第一个 server 块就是 localhost，哪怕你把 localhost 改为其他的，如 hello，它也能访问。只是因为在匹配不到的情况下，默认是第一个 service 块，哪怕它确实和访问的 IP 不匹配。

IP 如果和域名进行了绑定，那么在 `server_name` 中，两者都可以填写，填了域名，最后也是找到 IP，只是因为域名好记住。

知道了区别，在生产环境上，可以给个指定的错误页面。如果匹配不上，则返回友好的提示，如第一个 server 块：



```nginx
server{
	listen 80 default_server;
    server_name _;    # 匹配不上的时候，也可以填写 localhost
    return "<h1>不好意思，匹配不到！";
    # return 403;   # 也可以直接返回错误码
    # 
}

server{
	listen 80;
    server_name www.frxcat.fun;    # 匹配域名访问
    # ......
}
```

上面代码块只允许域名访问，而不允许 IP 访问，避免其他人把未备案的域名解析到自己的服务器 IP。

当然你也可以不设置错误页面。因为 Nginx 匹配不上时，直接返回它的欢迎界面。

> 在 Nginx 配置中，`listen` 和 `server_name` 确实是**双重匹配规则**，两者必须同时满足，请求才会被特定的 `server` 块处理。
>
> ```yaml
> # 处理来自公网的请求（例如面向用户的网站）
> server {
>     listen 203.0.113.10:80;       # 规则1：必须发送到这个公网IP
>     server_name example.com;       # 规则2：Host头必须是example.com
>     root /var/www/public;
> }
> 
> # 处理来自内网的请求（例如管理后台、API）
> server {
>     listen 192.168.1.100:80;       # 规则1：必须发送到这个内网IP
>     server_name internal.example.com; # 规则2：Host头必须是internal.example.com
>     root /var/www/internal;
> }
> ```
>
> > 如果同时设置了 listen指定IP 并且设置域名，这里需要双重匹配 如果在访问时 对应的域名没有转发到对应的IP nginx也并不会转发到对应后端的服务
> >
> > 1. **Nginx 首先根据 `listen` 指令进行第一层过滤**：
> >    - 检查请求的目标 IP 地址和端口
> >    - 只有发送到匹配 IP:端口的请求才会进入候选列表
> > 2. **然后在匹配的 `listen` 组内，根据 `server_name` 进行第二层过滤**：
> >    - 检查请求的 Host 头部
> >    - 选择最匹配的 `server` 块处理请求
>
> 所以这里能够通过设置内网IP设置 实现精确的流量控制能力，允许根据来源IP和请求域名进行精细化路由
>
> > #### 场景1：正确的内网访问
> >
> > - 用户在内网，DNS 将 `internal.example.com` 解析为 `192.168.1.100`
> > - 请求：`GET / HTTP/1.1 Host: internal.example.com` 发送到 `192.168.1.100:80`
> > - **结果**：匹配内网 server 块，因为：
> >   - `listen 192.168.1.100:80` ✅ (IP和端口匹配)
> >   - `server_name internal.example.com` ✅ (域名匹配)
> >
> > #### 场景2：错误的外网访问尝试
> >
> > - 用户在外网，DNS 将 `internal.example.com` 解析为公网 IP `203.0.113.10`
> > - 请求：`GET / HTTP/1.1 Host: internal.example.com` 发送到 `203.0.113.10:80`
> > - **结果**：**不会**匹配内网 server 块，因为：
> >   - `listen 192.168.1.100:80` ❌ (IP不匹配 - 请求发送到203.0.113.10而不是192.168.1.100)
> >   - 即使 `server_name internal.example.com` ✅ (域名匹配)，但第一层过滤已失败
> >
> > #### 场景3：公网访问公网服务
> >
> > - 用户在外网，DNS 将 `example.com` 解析为公网 IP `203.0.113.10`
> > - 请求：`GET / HTTP/1.1 Host: example.com` 发送到 `203.0.113.10:80`
> > - **结果**：匹配公网 server 块，因为：
> >   - `listen 203.0.113.10:80` ✅ (IP和端口匹配)
> >   - `server_name example.com` ✅ (域名匹配)

#### 匹配执行顺序

由于 server_name 指令支持通配符和正则表达式，因此在包含多个虚拟主机的配置文件中，可能会出现一个名称被多个虚拟主机的 server_name 匹配成功，当遇到这种情况，当前的请求交给谁来处理呢？如下：



```nginx
server{
	listen 80;
	server_name ~^www\.\w+\.com$;
	default_type text/plain;
	return 200 'regex_success';
}

server{
	listen 80;
	server_name www.frx.*;
	default_type text/plain;
	return 200 'wildcard_after_success';
}

server{
	listen 80;
	server_name *.frx.com;
	default_type text/plain;
	return 200 'wildcard_before_success';
}

server{
	listen 80;
	server_name www.frx.com;
	default_type text/plain;
	return 200 'exact_success';
}

server{
	listen 80 default_server;
	server_name _;
	default_type text/plain;
	return 444 'default_server not found server';
}
```

访问 `http://www.kele.com` 的优先级：（访问完后请注释掉处理这个请求的 server）

- exact_success（第一个是精确匹配，然后注释掉该 server）
- wildcard_before_success（第二个是开始匹配符，然后注释掉该 server）
- wildcard_after_success（第三个是结尾匹配符，然后注释掉该 server）
- regex_success（第四个是正则表达式，然后注释掉该 server）
- default_server not found server!!（第五个是默认 server，因为前四个都注释了，所以只能走默认的 server）

**结论**

No1：准确匹配 server_name

No2：通配符在开始时匹配 server_name 成功

No3：通配符在结束时匹配 server_name 成功

No4：正则表达式匹配 server_name 成功

No5：被默认的 default_server 处理，如果没有指定默认找第一个 server

### location指令

`location` 指令是用来设置请求的 URI。



```nginx
server{
	listen 80;
	server_name localhost;
	location / {
	
	}
	location /abc {
	
	}
	...
}
```

| 语法                                            | 默认值 | 位置             |
| ----------------------------------------------- | ------ | ---------------- |
| location [ = \| ~ \| ~* \| ^~ \| @ ] <uri>{...} | —      | server、location |



url 变量是待匹配的请求字符串，可以不包含正则表达式，也可以包含正则表达式，那么 Nginx 服务器在搜索匹配 location 的时候，是先使用不包含正则表达式进行匹配，找到一个匹配度最高的一个，然后在通过包含正则表达式的进行匹配，**如果能匹配到直接访问正则表达式的，匹配不到，就使用刚才匹配度最高（前缀最长的）的那个 location 来处理请求**。

- 不带符号，要求必须以指定模式开头，但是不要求精确匹配


```nginx
server {
	listen 80;
	server_name 127.0.0.1;
	location /abc {
		default_type text/plain;
		return 200 "access success";
	}
}

# 以下访问都是正确的
http://192.168.91.200/abc
http://192.168.91.200/abc?p1=kele
http://192.168.91.200/abc/
http://192.168.91.200/abcdef
```

- `=` 是用于不包含正则表达式的 uri，必须与指定的模式精确匹配

```nginx
server {
	listen 80;
	server_name 127.0.0.1;
	location =/abc {
		default_type text/plain;
		return 200 "access success";
	}
}
# 匹配成功
http://192.168.91.200/abc
http://192.168.91.200/abc?p1=TOM
# 匹配失败
http://192.168.91.200/abc/
http://192.168.91.200/abcdef
```

- `~` 是用于表示当前 uri 中包含了正则表达式，并且区分大小写

- `~*` 是用于表示当前 uri 中包含了正则表达式，但是是不区分大小写

```nginx
# 使用正则表达式，区分大小写
server {
	listen 80;
	server_name 127.0.0.1;
	location ~^/abc\w$ {
		default_type text/plain;
		return 200 "access success";
	}
}

# 使用正则表达式，不区分大小写
server {
	listen 80;
	server_name 127.0.0.1;
	location ~*^/abc\w$ {
		default_type text/plain;
		return 200 "access success";
	}
}
```

- `^~` 是用于不包含正则表达式的 uri，功能和不加符号的一致，唯一不同的是，如果请求匹配上了，那么就停止搜索其他模式了。



```nginx
server {
	listen 80;
	server_name 127.0.0.1;
	location ^~/abc {
		default_type text/plain;
		return 200 "access success";
	}
}
```

- `@` 前缀可以用来定义一个命名的 location,该 location 不处理正常的外部请求,一般用来当作标识供内部重定向使用。它们不能嵌套,也不能包含嵌套的 location。



```nginx
location /try {
    try_files $uri $uri/ @name;
}

location /error {
    error_page 404 = @name;
    return 404;
}

location @name {
    return 200 "@name";
}
```

这时访问 `/try` 或者 `/error` 都会返回 `@name`。

### root/alias指令

- `root` 指令是设置请求资源的根目录。默认值是 html。

| 语法         | 默认值     | 位置                   |
| ------------ | ---------- | ---------------------- |
| root <path>; | root html; | http、server、location |



path 是 Nginx 服务器接收到请求以后查找资源的根目录路径。

- `alias` 指令是用来更改 location 的 URI。

| 语法          | 默认值 | 位置     |
| ------------- | ------ | -------- |
| alias <path>; | —      | location |



- path 是修改后的根路径。

**以上两个指令都可以来指定访问资源的路径，那么这两者之间的区别是什么？**

> **举例说明**

1. 在 `/usr/local/nginx/html` 目录下创建一个 images 目录,并在目录下放入一张图片 `mv.png` 图片。

   然后进入配置文件，添加如下内容：



```nginx
location /images {
    root /usr/local/nginx/html;
}
```

访问图片的路径为：`http://192.168.91.200/images/mv.png`

1. 如果把root改为alias



```nginx
location /images {
    alias /usr/local/nginx/html;
}
```

再次访问上述地址，页面会出现 404 的错误，查看错误日志会发现是因为地址不对，所以验证了：

- root 的处理结果是：

  **追加路径**

  **root 路径 + location 路径**，location 路径包括匹配后面的请求，即包括 /mv.png

  `/usr/local/nginx/html/images/mv.png`

- alias 的处理结果是：**使用 alias 路径替换 location 路径**，但是不会替换匹配后面的请求，即不会替换 /mv.png

  最终文件路径 = alias路径 + (URI减去location匹配的部分)

需要在 alias 后面路径改为：



```nginx
location /images {
    alias /usr/local/nginx/html/images;
}
```

1. 如果 location 路径是以 / 结尾,则 alias 也必须是以 / 结尾，root 没有要求。

将上述配置修改为：



```nginx
location /images/ {
    alias /usr/local/nginx/html/images;
}
```

访问就会出问题，查看错误日志还是路径不对，所以需要把 alias 后面加上 /

小结：

- root 的处理结果是: root 路径 + location 路径
- alias 的处理结果是:使用 alias 路径替换 location 路径
- alias 是一个目录别名的定义，root 则是最上层目录的含义
- 如果 location 路径是以 / 结尾,则 alias 也必须是以 / 结尾，root 没有要求
  - **alias 不支持 location 的 =**

这里再多言几句，alias 后指定的资源路径，Nginx 就会去这个路径下找资源，「忽略」location 本身的的请求，仅拼接 location 后面的请求。如果你想去拼接 location 本身，就用绝对路径（包括 location）的alias。

### index指令

`index` 指令是**设置网站的默认首页**。默认是 index.html。

| 语法              | 默认值            | 位置                   |
| ----------------- | ----------------- | ---------------------- |
| index <file> ...; | index index.html; | http、server、location |



`index` 后面可以跟多个设置，如果访问的时候没有指定具体访问的资源，则会从左往右依次进行查找，找到第一个为止。

举例说明：



```nginx
location / {
	root /usr/local/nginx/html;
	index index.html index.htm;
}
```

访问该 location 的时候，可以通过 `http://ip:port/` 访问，地址后面如果不添加任何内容，则默认依次访问 index.html 和 index.htm，找到第一个来进行返回。

### error_page指令

error_page 指令是设置网站的错误页面。

| 语法                                          | 默认值 | 位置                          |
| --------------------------------------------- | ------ | ----------------------------- |
| error_page <code> ...... [=[response]] <uri>; | —      | http、server、location ...... |



code 是响应码。

**当出现对应的响应 code 后，如何来处理？**

> **举例说明**

1. 可以指定具体跳转的地址



```nginx
server {
	error_page 404 http://www.frx.com;
}
```

当页面产生 404 时，自动跳转到 `http://www.frx.com`

1. 可以指定重定向地址

```nginx
server{
	error_page 404 /50x.html;
	error_page 500 502 503 504 /50x.html;
	location =/50x.html {
		root html;
	}
}
```

产生错误页面时，重定向到 /50x.html，然后触发 location，最终访问的是 html 目录下的 50x.html 页面

1. 使用 location 的 @ 符合完成错误信息展示

```nginx
server{
	error_page 404 @jump_to_error;
	location @jump_to_error {
		default_type text/plain;
		return 404 'Not Found Page...';
	}
}
```

可选项 `=[response]` 的作用是用来将相应代码更改为另外一个，如下：

```nginx
server{
	error_page 404 =200 /50x.html;
	location =/50x.html {
		root html;
	}
}
```

这样的话，当返回 404 找不到对应的资源的时候，在浏览器上可以看到，最终返回的状态码是 200 而不是 404，这块需要注意下，编写 error_page 后面的内容，404 后面需要加空格，200 前面不能加空格。

## 静态资源优化配置

Nginx 对静态资源如何进行优化配置。这里从三个属性配置进行优化：



```nginx
sendfile on;
tcp_nopush on;
tcp_nodeplay on;
```

建议三个都开启。如果想知道为什么，请往下看。

### sendﬁle

该指令是用来开启高效的文件传输模式。默认关闭，建议开启。

| 语法                 | 默认值      | 位置                          |
| -------------------- | ----------- | ----------------------------- |
| sendﬁle <on \| off>; | sendﬁle oﬀ; | http、server、location ...... |



请求静态资源的过程：客户端通过网络接口向服务端发送请求，操作系统将这些客户端的请求传递给服务器端应用程序，服务器端应用程序会处理这些请求，请求处理完成以后，操作系统还需要将处理得到的结果通过网络适配器传递回去。

如：

```nginx
server {
	listen 80;
	server_name localhost；
	location / {
		root html;
		index index.html;
	}
}
```

假设在 html 目录下有一个 welcome.html 页面，访问地址：`http://192.168.91.200/welcome.html`。

流程如下：

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220731/image.24hrbukvg0lc.webp)

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220731/image.5l2efa1uqlo0.webp)

### tcp_nopush

该指令必须在 sendfile 打开的状态下才会生效，主要是用来提升网络包的传输「效率」。默认关闭。

| 语法                    | 默认值         | 位置                   |
| ----------------------- | -------------- | ---------------------- |
| tcp_nopush <on \| off>; | tcp_nopush oﬀ; | http、server、location |

#### tcp_nodelay

该指令必须在 keep-alive 连接开启的情况下才生效，来提高网络包传输的「实时性」。默认开启。

| 语法                     | 默认值          | 位置                   |
| ------------------------ | --------------- | ---------------------- |
| tcp_nodelay <on \| off>; | tcp_nodelay on; | http、server、location |

`tcp_nopush` 就像大巴车，等所有旅客占满了座位，才开始发车到景点（客户端），而 `tcp_nodelay`，上来一个旅客，就马上发车到景点客户端）。

#### 优化总结

经过分析，『 tcp_nopush 』和『 tcp_nodelay 』看起来是「互斥的」，那么为什么要将这两个值都打开，这个大家需要知道的是在 Linux2.5.9 以后的版本中两者是可以兼容的，三个指令都开启的好处是，sendfile 可以开启高效的文件传输模式，『 tcp_nopush 』开启可以确保在发送到客户端之前数据包已经充分「填满」，这大大减少了网络开销，并加快了文件发送的速度。然后，当它到达最后一个可能因为没有「填满」而暂停的数据包时，Nginx 会忽略『 tcp_nopush 』参数， 然后，『 tcp_nodelay 』强制套接字发送数据。由此可知，『 tcp_nopush 』可以与『 tcp_nodelay 』一起设置，它比单独配置『 tcp_nodelay 』具有更强的性能。

所以回归开头，我们可以使用如下配置来优化 Nginx 静态资源的处理：

```nginx
# 三个都开启
sendfile on;
tcp_nopush on;
tcp_nodelay on;
```

### 静态资源压缩配置

经过上述内容的优化，我们再次思考一个问题，假如在满足上述优化的前提下，我们传送一个 1M 的数据和一个 10M 的数据那个效率高？答案显而易见，**传输内容小，速度就会快**。那么问题又来了，同样的内容，如果把大小降下来，我们脑袋里面要蹦出一个词就是「压缩」，接下来，我们来学习 Nginx 的静态资源压缩模块。

在 Nginx 的配置文件中可以通过配置 gzip 来对静态资源进行压缩，相关的指令可以配置在 http 块、server 块和 location 块中，Nginx 可以通过对这些指令进行解析和处理：

- `ngx_http_gzip_module` 模块
- `ngx_http_gzip_static_module` 模块
- `ngx_http_gunzip_module` 模块

接下来我们从以下内容进行学习：

- Gzip 各模块支持的配置指令
- Gzip 压缩功能的配置
- Gzip 和 sendfile 的冲突解决
- 浏览器不支持 Gzip 的解决方案

#### Gzip模块配置指令

接下来所学习的指令都来自 `ngx_http_gzip_module` 模块，该模块会在 Nginx 安装的时候内置到 Nginx 的安装环境中，也就是说我们可以直接使用这些指令。

- `gzip` 指令是用于开启或者关闭 Gzip 功能。默认关闭

| 语法              | 默认值    | 位置                          |
| ----------------- | --------- | ----------------------------- |
| gzip <on \| off>; | gzip off; | http、server、location ...... |

注意：只有该指令为打开状态，下面的指令才有效果

```nginx
http{
	gzip on;
}
```

- `gzip_types` 指令可以根据响应页的 MIME 类型选择性地开启 Gzip 压缩功能。默认是 text/html

| 语法                            | 默认值                | 位置                   |
| ------------------------------- | --------------------- | ---------------------- |
| gzip_types <mime-type> ...... ; | gzip_types text/html; | http、server、location |



所选择的值可以从 mime.types 文件中进行查找，也可以使用 * 代表所有。



```nginx
http{
	gzip_types application/javascript;
	# * 代表所有
	gzip_types *
}
```

- `gzip_comp_level` 指令是用于设置 Gzip 压缩程度，级别从 1-9，1 表示要是程度最低，要是效率最高，9 刚好相反，压缩程度最高，但是效率最低、最费时间。默认值是 1

| 语法                     | 默认值             | 位置                   |
| ------------------------ | ------------------ | ---------------------- |
| gzip_comp_level <level>; | gzip_comp_level 1; | http、server、location |



```nginx
http{
	gzip_comp_level 6;
}
```

- `gzip_vary` 指令是用于设置使用 Gzip 进行压缩发送是否携带『Vary:Accept-Encoding』头域的响应头部。

  **主要是告诉接收方，所发送的数据经过了 Gzip 压缩处理。**默认关闭

| 语法                   | 默认值         | 位置                   |
| ---------------------- | -------------- | ---------------------- |
| gzip_vary <on \| off>; | gzip_vary off; | http、server、location |

- `gzip_buffers` 指令是用于处理请求压缩的缓冲区数量和大小

| 语法                          | 默认值                       | 位置                   |
| ----------------------------- | ---------------------------- | ---------------------- |
| gzip_buffers <number> <size>; | gzip_buffers 32 4k \| 16 8k; | http、server、location |



其中 number 是指定 Nginx 服务器向系统申请缓存空间个数，size 指的是每个缓存空间的大小。

主要实现的是申请 number 个每个大小为 size 的内存空间。

这个值的设定一般会和服务器的操作系统有关，所以建议此项不设置，使用默认值即可。



```nginx
gzip_buffers 4 16K;	  # 缓存空间大小
```

- `gzip_disable` 指令是针对不同种类客户端发起的请求，可以选择性地开启和关闭 Gzip 功能

| 语法                          | 默认值 | 位置                   |
| ----------------------------- | ------ | ---------------------- |
| gzip_disable <regex> ...... ; | —      | http、server、location |



**regex 是根据客户端的浏览器标志(user-agent)来设置，支持使用正则表达式。**

指定的浏览器标志不使用 Gzip.该指令一般是用来排除一些明显不支持 Gzip 的浏览器。



```nginx
gzip_disable "MSIE [1-6]\.";
```

- `gzip_http_version` 指令是针对不同的 HTTP 协议版本，可以选择性地开启和关闭 Gzip 功能。默认是 1.1 版本

| 语法                            | 默认值                 | 位置                   |
| ------------------------------- | ---------------------- | ---------------------- |
| gzip_http_version <1.0 \| 1.1>; | gzip_http_version 1.1; | http、server、location |

该指令是指定使用 Gzip 的 HTTP 最低版本，该指令一般采用默认值即可。

- `gzip_min_length` 指令是针对传输数据的大小，可以选择性地开启和关闭 Gzip 功能

| 语法                      | 默认值              | 位置                   |
| ------------------------- | ------------------- | ---------------------- |
| gzip_min_length <length>; | gzip_min_length 20; | http、server、location |

Nignx 计量大小的单位：bytes [字节] / kb [千字节] / M [兆]

例如: 1024 / 10k | K / 10m | M

Gzip 压缩功能对大数据的压缩效果明显，但是如果要压缩的数据比较小的话，可能出现越压缩数据量越大的情况，因此我们需要根据响应内容的大小来决定是否使用 Gzip 功能，响应页面的大小可以通过头信息中的 `Content-Length` 来获取。但是如何使用了 Chunk 编码动态压缩，该指令将被忽略。建议设置为 1K 或以上。

- `gzip_proxied` 指令设置是否对服务端返回的结果进行 Gzip 压缩

| 语法                                                         | 默认值            | 位置                   |
| ------------------------------------------------------------ | ----------------- | ---------------------- |
| gzip_proxied <off \| expired \| no-cache \| no-store \| private \| no_last_modified \| no_etag \| auth \| any>; | gzip_proxied off; | http、server、location |



- off：关闭 Nginx 服务器对后台服务器返回结果的 Gzip 压缩
- expired：如果 header 头中包含 『Expires』头信息，启用压缩
- no-cache：如果 header 头中包含 『Cache-Control:no-cache』头信息，启用压缩
- no-store：如果 header 头中包含 『Cache-Control:no-store』头信息，启用压缩
- private：如果 header 头中包含 『Cache-Control:private』头信息，启用压缩
- no_last_modified：如果 header 头中不包含 『Last-Modified』头信息，启用压缩
- no_etag：如果 header 头中不包含 『ETag』 头信息，启用压缩
- auth：如果 header 头中包含 『Authorization』 头信息，启用压缩
- any：无条件启用压缩

### Gzip压缩功能配置模板



```nginx
gzip on;  			   		 # 开启 Gzip 功能
gzip_types *;		   	 	 # 压缩源文件类型,根据具体的访问资源类型设定
gzip_comp_level 6;	   		 # Gzip 压缩级别
gzip_min_length 1k;          # 进行压缩响应页面的最小长度，content-length
gzip_buffers 4 16K;	         # 缓存空间大小
gzip_http_version 1.1;       # 指定压缩响应所需要的最低 HTTP 请求版本
gzip_vary  on;		         # 往头信息中添加压缩标识
gzip_disable "MSIE [1-6]\."; # 对 IE6 以下的版本都不进行压缩
gzip_proxied  off;           # Nginx 作为反向代理压缩服务端返回数据的条件
```

这些配置在很多地方可能都会用到，所以我们可以将这些内容抽取到一个配置文件中，然后通过 include 指令把配置文件再次加载到 nginx.conf 配置文件中，方法使用。

创建压缩配置文件：`nginx_gzip.conf`，添加如下内容：



```nginx
gzip on;
gzip_types *;
gzip_comp_level 6;
gzip_min_length 1k;
gzip_buffers 4 16K;
gzip_http_version 1.1;
gzip_vary  on;
gzip_disable "MSIE [1-6]\.";
gzip_proxied  off;
```

在 Nginx 核心配置文件 `nginx.conf` 进行引入，添加如下内容：



```nginx
include nginx_gzip.conf
```

### Gzip和sendfile共存问题

前面在讲解 sendfile 的时候，提到过，开启 sendfile 以后，在读取磁盘上的静态资源文件的时候，可以减少拷贝的次数，可以不经过用户进程将静态文件通过网络设备发送出去，但是 Gzip 要想对资源压缩，是需要经过用户进程进行操作的。所以如何解决两个设置的共存问题。

可以使用 `ngx_http_gzip_static_module` 模块的 `gzip_static` 指令来解决。

#### gzip_static指令

`gzip_static` 指令用于在检查与访问资源同名的 .gz 文件时，response 中以 Gzip 相关的 header 返回 .gz 文件的内容。默认关闭。

| 语法                               | 默认值           | 位置                   |
| ---------------------------------- | ---------------- | ---------------------- |
| gzip_static <on \| off \| always>; | gzip_static off; | http、server、location |





```nginx
gzip_static on;
```

在配置文件添加上述命令后，会报一个错误：`unknown directive "gzip_static"`，主要的原因是 Nginx 默认是没有添加 ngx_http_gzip_static_module 模块。如何来添加？

#### Nginx模块添加

1. 查询当前 Nginx 的配置参数，即查看 `configure arguments` 的配置信息，拷贝出来



```sh
nginx -V

# 拷贝 configure arguments 后面的数据
```



```nginx
--prefix=/usr/local/nginx
```

1. 将 Nginx 安装目录下 sbin 目录中的 nginx 二进制文件进行更名备份



```sh
cd /usr/local/nginx/sbin
mv nginx nginx.backup
```

1. 进入 Nginx 的安装目录



```sh
cd /root/nginx/core/nginx-1.21.6
```

1. 执行 make clean 清空之前编译的内容



```sh
make clean
```

1. 使用 configure 来配置参数，添加 `ngx_http_gzip_static_module` 模块，记得加上第1步拷贝的配置信息



```sh
./configure --with-http_gzip_static_module  # 记得添加 configure arguments 后的数据
```

1. 使用 make 命令进行编译



```sh
make
```

1. 将 objs 目录下的 nginx 二进制执行文件移动到 nginx 安装目录下的 sbin 目录中



```sh
mv /opt/nginx/core/nginx-1.21.6/objs/nginx /usr/local/nginx/sbin
```

如果不执行第（2）步进行备份，则该步骤会覆盖原来的 nginx 可执行文件

1. 在源码目录下执行更新命令



```sh
cd /opt/nginx/core/nginx-1.21.6
make upgrade
```

#### gzip_static测试

准备好一个 jquery.js 文件，放在 html 目录下

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220731/image.576ef03x00c0.webp)

1. 直接访问 `http://192.168.91.200/jquery.js`

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220731/image.2b8dhabn0cbo.webp)

1. 使用 Gzip 命令进行压缩



```sh
# 进入 html 目录
cd /usr/local/nginx/html

# 压缩 js 文件
gzip jquery.js
```

1. 再次访问 `http://192.168.91.200/jquery.js`

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220731/image.4ivj1d091pg.webp)

可以看出 `Content-Length` 的大小已经变得非常小。

## 静态资源缓存配置

当浏览器请求 Nginx 服务器的资源后，我们可以让这些资源缓存在浏览器里，这样再一次请求相同的资源时，无需请求 Nginx 服务器，直接从浏览器的缓存里获取，减少 Nginx 服务器的压力。

### 什么是缓存和Web缓存

缓存（cache），原始意义是指访问速度比一般随机存取存储器（RAM）快的一种高速存储器，通常它不像系统主存那样使用 DRAM 技术，而使用昂贵但较快速的 SRAM 技术。缓存的设置是所有现代计算机系统发挥高性能的重要因素之一。

Web 缓存是指一个 Web 资源（如 html 页面，图片，js，数据等）存在于 Web 服务器和客户端（浏览器）之间的副本。缓存会根据进来的请求保存输出内容的副本；当下一个请求来到的时候，如果是相同的 URL，缓存会根据缓存机制决定是直接使用副本响应访问请求，还是向源服务器再次发送请求。比较常见的就是浏览器会缓存访问过网站的网页，当再次访问这个 URL 地址的时候，如果网页没有更新，就不会再次下载网页，而是直接使用本地缓存的网页。只有当网站明确标识资源已经更新，浏览器才会再次下载网页。

### Web缓存的种类

客户端缓存

- 浏览器缓存

服务端缓存

- Nginx
- Redis
- Memcached 等

#### 为什么要用浏览器缓存

- 成本最低的一种缓存实现
- 减少网络带宽消耗
- 降低服务器压力
- 减少网络延迟，加快页面打开速度

#### 浏览器缓存执行流程

HTTP 协议中和页面缓存相关的字段，我们先来认识下：

| header        | 说明                                          |
| ------------- | --------------------------------------------- |
| Expires       | 缓存过期的日期和时间                          |
| Cache-Control | 设置和缓存相关的配置信息                      |
| Last-Modified | 请求资源最后修改时间                          |
| ETag          | 请求变量的实体标签的当前值，比如文件的 MD5 值 |



![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220731/image.209c4krjzgf4.webp)

1. 用户首次通过浏览器发送请求到服务端获取数据，客户端是没有对应的缓存，所以需要发送 request 请求来获取数据；
2. 服务端接收到请求后，获取服务端的数据及服务端缓存的允许后，返回 200 的成功状态码并且在响应头上附上对应资源以及缓存信息；
3. 当用户再次访问相同资源的时候，客户端会在浏览器的缓存目录中查找是否存在响应的缓存文件；
4. 如果没有找到对应的缓存文件，则走第2步；
5. 如果有缓存文件，接下来对缓存文件是否过期进行判断，过期的判断标准是(Expires)；
6. 如果没有过期，则直接从本地缓存中返回数据进行展示；
7. 如果 Expires 过期，接下来需要判断缓存文件是否发生过变化；
8. 判断的标准有两个，一个是 ETag(Entity Tag)，一个是 Last-Modified；
9. 判断结果是未发生变化，则服务端返回 304，直接从缓存文件中获取数据；
10. 如果判断是发生了变化，重新从服务端获取数据，并根据缓存协商(服务端所设置的是否需要进行缓存数据的设置)来进行数据缓存。

### 浏览器缓存相关指令

![image-20250921123233045](.\asset\images\image-20250921123233045.png)

Nginx 需要进行缓存相关设置，就需要用到如下的指令。

#### expires指令

该指令用来控制页面缓存的作用。

**可以通过该指令控制 HTTP 应答中的『Expires』和『Cache-Control』**

| 语法                                                         | 默认值       | 位置                   |
| ------------------------------------------------------------ | ------------ | ---------------------- |
| expires [modified] <time>;<br />expires <epoch \| max \| off>; | expires off; | http、server、location |



- time：可以整数也可以是负数，指定过期时间，单位为 s（秒）。

  如果是负数，Cache-Control 则为 no-cache，

  如果为整数或 0，则 Cache-Control 的值为 max-age=time

- epoch：指定 Expires 的值为『'1 January,1970,00:00:01 GMT'』，即 1970-01-01 00:00:00 ，Cache-Control 的值 no-cache

  因为 Expires（缓存过期时间）是 1970 年，所以不缓存。

- max：指定 Expires 的值为『'31 December2037 23:59:59GMT' 』，即(2037-12-31 23:59:59，Cache-Control 的值为 10 年

  因为 Expires（缓存过期时间）是 2037 年，虽然还有 16 年过期，但是最大只能缓存 10 年。

- off：默认不缓存

> **例子 1**

在配置文件添加如下内容：

```nginx
location ~ .*\.(html|js|css|png|jpg|jpeg|gif)$ {
    # ...
    expires max
    # ...
}
```

发送请求：`http://192.168.91.200/jquery.js`

315360000 折算下来正好是 10 年。

> **其他格式**

```nginx
expires 30s;  # 表示把数据缓存 30 秒

expires 30m;  # 表示把数据缓存 30 分

expires 10h;  # 表示把数据缓存 10 小时

expires 1d;   # 表示把数据缓存 1 天
```

#### add_header指令

add_header 指令是用来添加指定的响应头和响应值。

add_header 是响应体的指令，不是请求时的指令（比如 expires ），并且 add_header 也有和 expires 一样的功能。

如果 expires 和 add_header 同时开启的情况下，则 add_header 优于 expires 生效。

| 语法                                | 默认值 | 位置                              |
| ----------------------------------- | ------ | --------------------------------- |
| add_header <name> <value> [always]; | —      | location > server > http > ...... |



always 可选，代表总是添加。

Cache-Control 作为响应头信息，可以在 Nginx 配置文件设置如下缓存响应指令：



```nginx
add_header Cache-control must-revalidate;
add_header Cache-control no-cache;
add_header Cache-control no-store;
add_header Cache-control no-transform;
add_header Cache-control public;
add_header Cache-control private;
add_header Cache-control proxy-revalidate;
add_header Cache-Control max-age=<seconds>;  # 秒
add_header Cache-control s-maxage=<seconds>; # 秒
```

描述：

| 指令             | 说明                                                         |
| ---------------- | ------------------------------------------------------------ |
| must-revalidate  | 可缓存但必须再向源服务器进行确认                             |
| no-cache         | 数据内容不能被缓存，每次请求都重新访问服务器，若有 max-age，则缓存期间不访问服务器 |
| no-store         | 不缓存请求或响应的任何内容，暂存也不可以(临时文件夹中不能暂存该资源) |
| no-transform     | 代理不可更改媒体类型                                         |
| public           | 可以被任何缓存区缓存，如: 浏览器、服务器、代理服务器等       |
| private（默认）  | 只能在浏览器中缓存，只有在第一次请求的时候才访问服务器，若有 max-age，则缓存期间不访问服务器 |
| proxy-revalidate | 要求中间缓存服务器对缓存的响应有效性再进行确认               |
| max-age=<秒>     | 过期时间，即以秒为单位的缓存时间                             |
| s-maxage=<秒>    | 公共缓存服务器响应的最大 Age 值                              |



值得注意的是：

- 设置了 no-cache 或者 private 后，打开新窗口时会重新访问服务器。若设置 max-age，则缓存期间不访问服务器
- 设置 private 和正数的 max-age 时，后退时候不会访问服务器
- 设置 no-cache 和正数的 max-age 时，后退时会访问服务器

### Nginx服务端缓存

在 Web 缓存的种类，我们提到了 Ngixn 服务端缓存，而上面仅仅介绍了在浏览器进行缓存，而因为 Ngixn 服务端缓存的内容比较多，所以前往 [Nginx - 缓存集成](https://frxcat.fun/middleware/Nginx/Nginx_Static_resource_deployment/) 进行学习。





## Nginx的跨域问题

跨域问题，我们主要从以下方面进行解决：

- 什么情况下会出现跨域问题
- 实例演示跨域问题
- 具体的解决方案是什么

### 同源策略

浏览器的同源策略：是一种约定，是浏览器最核心也是最基本的安全功能，如果浏览器少了同源策略，则浏览器的正常功能可能都会受到影响。

同源：协议、域名(IP)、端口相同即为同源

```sh
http://192.168.200.131/user/1
https://192.168.200.131/user/1
# 不满足同源

http://192.168.200.131/user/1
http://192.168.200.132/user/1
# 不满足同源

http://192.168.200.131/user/1
http://192.168.200.131:8080/user/1
# 不满足同源

http://www.nginx.com/user/1
http://www.nginx.org/user/1
# 不满足同源

http://192.168.200.131/user/1
http://192.168.200.131:8080/user/1
# 不满足同源

http://www.nginx.org:80/user/1
http://www.nginx.org/user/1
# 满足同源
```

### 跨域问题

简单描述下：

有两台服务器分别为 A、B，如果从服务器 A 的页面发送异步请求到服务器 B 获取数据，如果服务器 A 和服务器 B 不满足同源策略，则就会出现跨域问题。

### 跨域案例

出现跨域问题会有什么效果？接下来通过一个需求来给大家演示下：

![image-20250921125431310](.\asset\images\image-20250921125431310.png)

1、Nginx 的 html 目录下新建一个 a.html

```sh
vim /usr/local/nginx/html/a.htm
```

添加如下内容：

```html
<html>
  <head>
        <meta charset="utf-8">
        <title>跨域问题演示</title>
        <script src="jquery.js"></script>
        <script>
            $(function(){
                $("#btn").click(function(){
                        $.get('http://192.168.200.133:8080/getUser',function(data){
                                alert(JSON.stringify(data));
                        });
                });
            });
        </script>
  </head>
  <body>
        <input type="button" value="获取数据" id="btn"/>
  </body>
</html>
```

2、在 nginx.conf 配置如下内容

```sh
vim /usr/local/nginx/conf/nginx.conf
```



```nginx
server{
    listen  8080;
    server_name localhost;
    location /getUser{
        default_type application/json;
        return 200 '{"id":1,"name":"TOM","age":18}';
    }
}
server{
	listen 	80;
	server_name localhost;
	location /{
		root html;
		index index.html;
	}
}
```

3、通过浏览器测试访问

![image](https://cdn.jsdmirror.com//gh/xustudyxu/image-hosting1@master/20220801/image.4md37urjzug0.webp)

### 解决方案

使用 `add_header` 指令，该指令可以用来添加一些头信息。

| 语法                             | 默认值 | 位置                   |
| -------------------------------- | ------ | ---------------------- |
| add_header <name> <value> ...... | —      | http、server、location |



此处用来解决跨域问题，需要添加两个头信息，分别是

- `Access-Control-Allow-Origin`
- `Access-Control-Allow-Methods`

`Access-Control-Allow-Origin`：直译过来是**允许跨域访问的源地址信息**，可以配置多个(多个用逗号分隔)，也可以使用 `*` 代表所有源。

`Access-Control-Allow-Methods`：直译过来是**允许跨域访问的请求方式**，值可以为 GET、POST、PUT、DELETE ......，

可以全部设置，也可以根据需要设置，多个用逗号分隔。

这里设置B server的配置，允许A server能够直接访问

具体配置方式：

```nginx
location /getUser {
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Methods GET,POST,PUT,DELETE;
    default_type application/json;   # return 的格式是 json
    return 200 '{"id":1,"name":"TOM","age":18}';
}
```

## 静态资源防盗链

### 什么是资源盗链

资源盗链指的是此内容不在自己服务器上，而是通过技术手段，绕过别人的限制将别人的内容放到自己页面上最终展示给用户。

以此来盗取大网站的空间和流量。简而言之就是用别人的东西成就自己的网站。

提供两种图片进行演示：

- 京东：`https://img14.360buyimg.com/n7/jfs/t1/101062/37/2153/254169/5dcbd410E6d10ba22/4ddbd212be225fcd.jpg`
- 百度：`https://pics7.baidu.com/feed/cf1b9d16fdfaaf516f7e2011a7cda1e8f11f7a1a.jpeg?token=551979a23a0995e5e5279b8fa1a48b34&s=BD385394D2E963072FD48543030030BB`

我们在 html 目录下准备一个页面 a.html，在页面上利用 img 标签引入这两个图片:

从上面的效果，可以看出来，下面的图片地址添加了防止盗链的功能，京东这边我们可以直接使用其图片。

### 防盗链实现原理

了解防盗链的原理之前，我们得先学习一个 HTTP 的头信息 Referer，

**当浏览器向 Web 服务器发送请求的时候，一般都会带上 Referer，来告诉浏览器该网页是从哪个页面链接过来的。*

**后台服务器可以根据获取到的这个 Referer 信息来判断是否为自己信任的网站地址，如果是则放行继续访问，如果不是则可以返回 403（服务端拒绝访问）的状态信息。**

### 防盗链实现实例

在本地模拟上述的服务器效果图：

Nginx 防盗链的具体实现：

**valid_referers 指令**：Nginx 会通过查看 Referer 自动和 valid_referers 的内容进行匹配，**如果匹配到了就将 `$invalid_referer` 变量置 0，如果没有匹配到，则将 `$invalid_referer` 变量置为 1**，匹配的过程中不区分大小写。

所以我们可以在配置文件判断 `$invalid_referer` 是否等于 1（true），即没有匹配到 ，等于则返回 403。

| 语法                                                         | 默认值 | 位置             |
| ------------------------------------------------------------ | ------ | ---------------- |
| valid_referers <none \| blocked \| server_names \| string> ...... | —      | server、location |



- none：如果 Header 中的 Referer 为空，允许访问
- blocked：在 Header 中的 Referer 不为空，但是该值被防火墙或代理进行伪装过，如不带『 http:// 』 、『 https:// 』等协议头的资源才允许访问。
- server_names：指定具体的域名或者 IP
- string：可以支持正则表达式和 `*` 的字符串。如果是正则表达式，需要以 `~` 开头表示

例如：



```nginx
location ~ *\.(png|jpg|gif){
    valid_referers none blocked www.baidu.com 192.168.91.200;
    
    # valid_referers none blocked *.example.com example.*  www.example.org  ~\.google\.;
    
    if ($invalid_referer){
        return 403;
    }
    root /usr/local/nginx/html;
}
```

上方代码如果没有匹配上 `www.baidu.com` 和 `192.168.91.200`，则 `$invalid_referer` 为 1（true），返回 403，代表不允许获取资源。

Nginx 配置文件支持 if 判断，但是 if 后面必须有空格。

**问题：如果图片有很多，该如何批量进行防盗链？可以针对目录进行防盗链。**

### 针对目录防盗链

假设 html 目录下有一个 images 目录，里面专门放防盗链的图片。

配置如下：



```nginx
location /images {
    valid_referers none blocked www.baidu.com 192.168.199.27;
    
    # valid_referers none blocked *.example.com example.*  www.example.org  ~\.google\.;
    
    if ($invalid_referer){
        return 403;
    }
    root /usr/local/nginx/html;
}
```

只需将 location 的地址改成一个目录，这样我们可以对一个目录下的所有资源进行翻到了操作。

**问题：Referer 的限制比较粗，比如浏览器发送请求时恶意加一个 Referer，上面的方式是无法进行限制的。那么这个问题改如何解决？**

此时我们需要用到 Nginx 的第三方模块 `ngx_http_accesskey_module`，第三方模块如何实现盗链，如何在 Nginx 中使用第三方模块的功能，在后面有讲解。





## Rewrite

Rewrite 是 Nginx 服务器提供的一个重要基本功能，是 Web 服务器产品中几乎必备的功能。主要的作用是用来实现 URI 的重写。

注意

Nginx 服务器的 Rewrite 功能的实现依赖于 PCRE 的支持，因此在编译安装 Nginx 服务器之前，需要安装 PCRE 库。

Nginx 使用的是`ngx_http_rewrite_module` 模块来解析和处理 Rewrite 功能的相关配置。

### 地址重写与地址转发

重写和转发的区别:

- **地址重写浏览器地址会发生变化而地址转发则不变**
- 一次地址重写会产生两次请求而一次地址转发只会产生一次请求
- 地址重写到的页面必须是一个完整的路径而地址转发则不需要
- 地址重写因为是两次请求，所以 request 范围内属性不能传递给新页面，而地址转发因为是一次请求所以可以传递值
- 地址转发速度快于地址重写

### set指令

该指令用来设置一个新的变量。

| 语法                | 默认值 | 位置                 |
| ------------------- | ------ | -------------------- |
| set <$key> <value>; | —      | server、location、if |

- variable：变量的名称，该变量名称要用 `$` 作为变量的第一个字符，且不能与 Nginx 服务器内置的全局变量同名。
- value：变量的值，可以是字符串、其他变量或者变量的组合等。

例如：



```nginx
server {
    listen 8081;
    server_name localhost;
    location /server {
        set $name TOM;
        set $age 18;
        default_type text/plain;
        return 200 $name=$age;
    }
}
```

访问 `https://192.168.200.133:8081:server`，返回结果如图：


> **Rewrite常用全局变量**
>
> | 变量               | 说明                                                         |
> | ------------------ | ------------------------------------------------------------ |
> | $args              | 变量中存放了请求 URL 中的请求指令。比如 `http://192.168.200.133:8080?arg1=value1&args2=value2` 中的『 arg1=value1&arg2=value2 』，功能和 $query_string 一样 |
> | $http_user_agent   | 变量存储的是用户访问服务的代理信息（如果通过浏览器访问，记录的是浏览器的相关版本信息） |
> | $host              | 变量存储的是访问服务器的 server_name 值                      |
> | $document_uri      | 变量存储的是当前访问地址的URI。比如 `http://192.168.200.133/server?id=10&name=zhangsan`中的『 /server 』，功能和 $uri 一样 |
> | $document_root     | 变量存储的是当前请求对应 location 的 root 值，如果未设置，默认指向 Nginx 自带 html 目录所在位置 |
> | $content_length    | 变量存储的是请求头中的 Content-Length 的值                   |
> | $content_type      | 变量存储的是请求头中的 Content-Type 的值                     |
> | $http_cookie       | 变量存储的是客户端的 cookie 信息，可以通过 `add_header Set-Cookie 'cookieName=cookieValue'` 来添加 cookie 数据 |
> | $limit_rate        | 变量中存储的是 Nginx 服务器对网络连接速率的限制，也就是 Nginx 配置中对 limit_rate 指令设置的值，默认是 0，不限制。 |
> | $remote_addr       | 变量中存储的是客户端的 IP 地址                               |
> | $remote_port       | 变量中存储了客户端与服务端建立连接的端口号                   |
> | $remote_user       | 变量中存储了客户端的用户名，需要有认证模块才能获取           |
> | $scheme            | 变量中存储了访问协议                                         |
> | $server_addr       | 变量中存储了服务端的地址                                     |
> | $server_name       | 变量中存储了客户端请求到达的服务器的名称                     |
> | $server_port       | 变量中存储了客户端请求到达服务器的端口号                     |
> | $server_protocol   | 变量中存储了客户端请求协议的版本，比如 『 HTTP/1.1 』        |
> | $request_body_file | 变量中存储了发给后端服务器的本地文件资源的名称               |
> | $request_method    | 变量中存储了客户端的请求方式，比如『 GET 』,『 POST 』等     |
> | $request_filename  | 变量中存储了当前请求的资源文件的路径名                       |
> | $request_uri       | 变量中存储了当前请求的 URI，并且携带请求参数，比如 `http://192.168.200.133/server?id=10&name=zhangsan` 中的 『 /server?id=10&name=zhangsan 』 |



> **例如**

```nginx
server {
    listen 8081;
    server_name localhost;
    location /server {
        root /usr/local/nginx/abc;
        set $name TOM;
        set $age 18;
        default_type text/plain;
        return 200 $name=$age=$args=$http_user_agent=$host=$document_root;
    }
}
```

访问：`http://192.168.200.133:8081/server?username=JERRY&gender=1`

效果如图：

> **可以把访问的信息记录在日志中**
>
> ```nginx
> http{
> 	# ......
> 	log_format main '$remote_addr - $request - $status - $request_uri - $http_user_agent';
>     server {
>         listen 8081;
>         server_name localhost;
>         location /server {
>         	access_log logs/access.log main;
>             root /usr/local/nginx/abc;
>             set $name TOM;
>             set $age 18;
>             default_type text/plain;
>             return 200 $name=$age=$args=$http_user_agent=$host=$document_root;
>         }
>     }
> }
> ```



访问：`http://192.168.200.133:8081/server?username=JERRY&gender=1`

然后查看日志，效果如图：

### if指令

该指令用来支持条件判断，并根据条件判断结果选择不同的 Nginx 配置。

| 语法                   | 默认值 | 位置             |
| ---------------------- | ------ | ---------------- |
| if (condition) { ... } | —      | server、location |

if 和括号之间要有空格，condition 为判定条件，可以支持以下写法：

- 变量名。如果变量名对应的值为空或者是 0，if 都判断为 false，其他条件为 true。

```nginx
if ($param){
	
}
```

- 使用『 = 』和『 != 』比较变量和字符串是否相等，满足条件为 true，不满足为 false

```nginx
if ($request_method = POST){
	return 405;
}
```

注意：POST 和 Java 不太一样的地方是字符串不需要添加引号。

- 使用正则表达式对变量进行匹配，匹配成功返回 true，否则返回 false。变量与正则表达式之间使用『 ~ 』，『 ~* 』，『 !~ 』，『 !~* 』来连接。
  - 『 ~ 』代表匹配正则表达式过程中区分大小写，进行模糊匹配
  - 『 ~* 』代表匹配正则表达式过程中不区分大小写，进行模糊匹配
  - 『 !~ 』和『 !~* 』刚好和上面取相反值，如果匹配上返回 false，匹配不上返回 true，进行模糊匹配

```nginx
if ($http_user_agent ~ MSIE){
	# $http_user_agent 的值中是否包含 MSIE 字符串，如果包含返回 true
}
```

- 判断请求的文件是否存在使用『 -f 』和『 !-f 』
  - 当使用『 -f 』时，如果请求的文件存在返回 true，不存在返回 false。
  - 当使用『 !-f 』时，如果请求文件不存在，但该文件所在目录存在返回 true，文件和目录都不存在返回 false，如果文件存在返回 false。

```nginx
if (-f $request_filename){
	# 判断请求的文件是否存在
}
if (!-f $request_filename){
	# 判断请求的文件是否不存在
}
```

例如：用户访问的页面不存在，则返回一个友好的提示

```nginx
location / {
    root html;
    default_type text/html;
    # 判断请求的文件是否不存在
    if (!-f $request_filename){
        return 200 '<h1>不好意思，文件资源找不到！</h1>';
    }
}
```

- 判断请求的目录是否存在使用『 -d 』和『 !-d 』

  当使用『 -d 』时，如果请求的目录存在，返回 true，如果目录不存在则返回 false。

  当使用『 !-d 』时，如果请求的目录不存在但该目录的上级目录存在则返回 true，该目录和它上级目录都不存在则返回 false，如果请求目录存在也返回false。

### break指令

该指令用于中断当前相同作用域中的其他 Nginx 配置。

与该指令处于同一作用域的 Nginx 配置中，位于它前面的指令配置生效，位于后面的指令配置无效。

**并且break还有另外一个功能就是终止当前的匹配并把当前的URI在本location进行重定向访问处理。**

| 语法   | 默认值 | 位置                 |
| ------ | ------ | -------------------- |
| break; | —      | server、location、if |

例子：

```nginx
location /testbreak {
    default_type text/plain;
    set $username TOM;
	if ($args){
		set $username JERRY;
		break;
		set $username ROSE;
	}
    add_header username $username;
    return 200 $username;
}
```



### return指令

该指令用于完成对请求的处理，直接向客户端返回响应状态代码。在 return 后的所有 Nginx 配置都是无效的。

| 语法                                                     | 默认值 | 位置                 |
| -------------------------------------------------------- | ------ | -------------------- |
| return <code> [text]; return <code> <URL>; return <URL>; | —      | server、location、if |

- code：返回给客户端的 HTTP 状态代理。可以返回的状态代码为 0 ~ 999 的任意 HTTP 状态代理
- text：返回给客户端的响应体内容，支持变量的使用和 JSON 字符串
- URL：跳转给客户端的 URL 地址。

例如：

```nginx
location / {
    default_type text/plain;
    return 200 "欢迎使用 Nginx";
}

location /baidu {
    return 302 https://www.baidu.com;
}
```

此时访问 Nginx，就会在页面看到这句话：欢迎使用 Nginx。

如果访问 `/baidu`，则跳转到 `https://www.baidu.com`。

### rewrite指令

该指令通过正则表达式的使用来改变 URI。可以同时存在一个或者多个指令，按照顺序依次对 URL 进行匹配和处理。

URL 和 URI 的区别：

- URI：统一资源标识符
- URL：统一资源定位符

| 语法                              | 默认值 | 位置                 |
| --------------------------------- | ------ | -------------------- |
| rewrite regex replacement [flag]; | —      | server、location、if |

- regex：用来匹配 URI 的正则表达式

- replacement：匹配成功后，用于替换 URI 中被截取内容的字符串。如果该字符串是以 『 http:// 』或者『 https:// 』开头的，则不会继续向下对URI 进行其他处理，而是直接返回重写后的 URI 给客户端。

  例如：（括号的值会作为 $1 的值）^ 代表匹配输入字符串的起始位置



```nginx
# ......
listen 8081;
location /rewrite {
    rewrite ^/rewrite/url\w*$ https://www.baidu.com;
    rewrite ^/rewrite/(test)\w*$ /$1;   # 如果是 /rewrite/testxxx，则重写 uri 为 test
    rewrite ^/rewrite/(demo)\w*$ /$1;    # 如果是 /rewrite/demoxxx，则重写 uri 为 demo
}
location /test {   # 重写后的 url 如果为 test，触发 location
    default_type text/plain;
    return 200 test_sucess;
}
location /demo {   # 重写后的 url 如果为 demo，触发 location
    default_type text/plain;
    return 200 demo_sucess;
}
```

访问 `http://192.168.200.113/8081/rewrite/urlxxx`，跳转到 `https://www.baidu.com`。

访问 `http://192.168.200.113/8081/rewrite/testxxx`，返回 test_sucess。

访问 `http://192.168.200.113/8081/rewrite/demoxxx`，返回 demo_sucess。

#### flag

**flag：用来设置 Rewrite 对 URI 的处理行为，可选值有如下：**

- `last`：终止继续在本 location 块中处理接收到的后续 URI，并将此处重写的 URl 作为一个新的 URI，使用各 location 块进行处理。

  该标志将重写后的 URI 重写在 server 块中执行，为重写后的 URI 提供了转入到其他 location 块的机会。

  **重写地址后访问其他的 location 块，浏览器地址栏 URL 地址不变**



```nginx
# ......
listen 8081;
location /rewrite {
    rewrite ^/rewrite/(test)\w*$ /$1 last;   # 如果是 /rewrite/testxxx，则重写 url 为 test
    rewrite ^/rewrite/(demo)\w*$ $1 last;    # 如果是 /rewrite/demoxxx，则重写 url 为 demo
}
location /test {   # 重写后的 url 如果为 test，触发 location
    default_type text/plain;
    return 200 test_sucess;
}
location /demo {   # 重写后的 url 如果为 demo，触发 location
    default_type text/plain;
    return 200 demo_sucess;
}
```

访问 `http://192.168.200.113/8081/rewrite/testxxx`，返回 test_sucess。

访问 `http://192.168.200.113/8081/rewrite/demoxxx`，返回 demo_sucess。

单次访问不明显，多次访问，last 只处理第一个。



- `break`：将此处重写的 URl 作为一个新的 URI，在本块中继续进行处理。

  该标志将重写后的地址在当前的 location 块中执行，不会将新的 URI 转向其他的 location 块。

  **仅仅重写地址，不会触发其他 location 块，浏览器地址栏 URL 地址不变**



```nginx
# ......
listen 8081;
location /rewrite {
    rewrite ^/rewrite/(test)\w*$ /$1 break;   # 如果是 /rewrite/testxxx，则重写 url 为 test
    rewrite ^/rewrite/(demo)\w*$ $1 break;    # 如果是 /rewrite/demoxxx，则重写 url 为 demo
    
    # /test 和 /demo 就在当前块进行处理，所以会在当前的 location 块找到如下 html 页面：
    # /usr/local/nginx/html/test/index.html
    # /usr/local/nginx/html/demo/index.html
}
location /test {   # 重写后的 url 如果为 test，触发 location
    default_type text/plain;
    return 200 test_sucess;
}
location /demo {   # 重写后的 url 如果为 demo，触发 location
    default_type text/plain;
    return 200 demo_sucess;
}
```

和 break 指令类似。假设访问的是 /test，则将 /test 放在当前的 location 块进行处理，哪怕第二个 location 块就是处理 /test 的，它也不会去找第二个 location 块，只在当前块进行处理。所以他会请求 `/usr/local/nginx/html/test/index.html`。



- `redirect`：将重写后的 URI 返回给客户端，状态码为 302，指明是临时重定向 URL，

  主要用在 replacement 变量不是以『 http:// 』或者『 https:// 』开头的情况

```nginx
# ......
listen 8081;
location /rewrite {
    rewrite ^/rewrite/(test)\w*$ /$1 redirect;   # 如果是 /rewrite/testxxx，则重写 url 为 test
    rewrite ^/rewrite/(demo)\w*$ $1 redirect;    # 如果是 /rewrite/demoxxx，则重写 url 为 demo
}
location /test {   # 重写后的 url 如果为 test，触发 location
    default_type text/plain;
    return 200 test_sucess;
}
location /demo {   # 重写后的 url 如果为 demo，触发 location
    default_type text/plain;
    return 200 demo_sucess;
}
```

特点是重定向，就是浏览的地址栏会发送改变。如发送请求 `/testxxx`，它会重定向到 `/test`，触发第二个 location 块，浏览的地址栏也会由 `/testxxx` 变成 `/test`。





- `permanent`：将重写后的 URI 返回给客户端，状态码为 301，指明是永久重定向 URL，

  主要用在 replacement 变量不是以『 http:// 』或者『 https:// 』开头的情况

```nginx
# ......
listen 8081;
location /rewrite {
    rewrite ^/rewrite/(test)\w*$ /$1 permanent;   # 如果是 /rewrite/testxxx，则重写 url 为 test
    rewrite ^/rewrite/(demo)\w*$ $1 permanent;    # 如果是 /rewrite/demoxxx，则重写 url 为 demo
}
location /test {   # 重写后的 url 如果为 test，触发 location
    default_type text/plain;
    return 200 test_sucess;
}
location /demo {   # 重写后的 url 如果为 demo，触发 location
    default_type text/plain;
    return 200 demo_sucess;
}
```

和 `redirect` 的区别就是状态码为 301，并且是永久重定向。

> - “`replacement` 变量不是以 『http://』 或者 『https://』 开头”：
>   - 如果 `replacement` **是**以 `http://` 或 `https://` 开头，Nginx 会认为这是一个完整的绝对路径URL，会直接把它作为重定向目标。
>   - 如果 `replacement` **不是**以它们开头，Nginx 会认为这是一个**相对路径**。此时，Nginx 会自动在这个相对路径的前面加上 **`$scheme://$server_name`** 来构建一个完整的 URL。
>     - `$scheme`： 代表请求使用的协议（http 或 https）。
>     - `$server_name`： 代表服务器配置中匹配的域名。
>
> > 假设你的服务器域名为 `www.example.com`。
> >
> > #### 场景：网站改版，旧文章路径改变
> >
> > 旧的文章 URL 结构是：`/old-news/123.html`
> > 新的文章 URL 结构是：`/archives/123`
> >
> > 我们希望所有访问旧链接的用户都被自动临时重定向到新链接。
> >
> > **Nginx 配置如下：**
> >
> > ```nginx
> > server {
> >     listen 80;
> >     server_name www.example.com;
> > 
> >     # 使用 rewrite 指令 with ‘redirect' flag
> >     rewrite ^/old-news/(\d+)\.html$ /archives/$1 redirect;
> >     # 正则表达式 ^/old-news/(\d+)\.html$ 用于匹配旧路径
> >     # replacement /archives/$1 不是以 http(s):// 开头
> >     # flag 是 redirect (302 临时重定向)
> > 
> >     location / {
> >         # ... 其他配置
> >     }
> > }
> > ```
> >
> > #### 访问流程分析：
> >
> > 1. 用户在浏览器中输入：`http://www.example.com/old-news/123.html`
> > 2. Nginx 收到请求，匹配到 `rewrite` 规则。
> > 3. 规则中的 `replacement` 是 `/archives/$1`（`$1` 是正则捕获的数字 `123`），它**不是以 `http://` 或 `https://` 开头**。
> > 4. 根据规则，Nginx 会**自动补全**域名和协议，构建完整的重定向URL：`http://www.example.com/archives/123`
> > 5. Nginx 向客户端（浏览器）返回一个 **302** 响应，响应头中包含：`Location: http://www.example.com/archives/123`
> > 6. 浏览器收到 302 响应，**自动地**向新地址 `http://www.example.com/archives/123` 发起一个新的 GET 请求。
> > 7. 用户最终看到新地址的内容，并且浏览器地址栏的URL也变成了新地址。

#### flag 总结

| 标记符号  | 说明                                                 |
| --------- | ---------------------------------------------------- |
| last      | 本条规则匹配完成后继续向下匹配新的 location URI 规则 |
| break     | 本条规则匹配完成后终止，不在匹配任何规则             |
| redirect  | 返回 302 临时重定向                                  |
| permanent | 返回 301 永久重定向                                  |



- break 与 last 都停止处理后续重写规则，只不过 last 会重新发起新的请求并使用新的请求路由匹配location，但 break 不会。所以当请求 break 时，如匹配成功，则请求成功，返回 200；如果匹配失败，则返回 404
- 服务器配置好 redirect 和 permanent 之后，打开浏览器分别访问这两个请求地址，然后停止 Nginx 服务。这时再访问 redirect 请求会直接报出无法连接的错误。但是 permanent 请求是永久重定向，浏览器会忽略原始地址直接访问永久重定向之后的地址，所以请求仍然成功。（这个验证不能禁用浏览器的缓存，否则即使是 permanent 重定向，浏览器仍然会向原始地址发出请求验证之前的永久重定向是否有效）
- 对于搜索引擎来说，搜索引擎在抓取到 301 永久重定向请求响应内容的同时也会将原始的网址替换为重定向之后的网址，而对于 302 临时重定向请求则仍然会使用原始的网址并且可能会被搜索引擎认为有作弊的嫌疑。所以对于线上正式环境来讲，尽量避免使用 302 跳转
- 如果 replacement 以 「 http:// 」或「 https:// 」或「 $scheme 」开始，处理过程将终止，并将这个重定向直接返回给客户端

#### rewrite_log指令

该指令配置是否开启 URL 重写日志的输出功能，默认关闭。

| 语法                     | 默认值           | 位置                       |
| ------------------------ | ---------------- | -------------------------- |
| rewrite_log <on \| off>; | rewrite_log off; | http、server、location、if |

**开启后，URL 重写的相关日志将以 notice 级别输出到 error_log 指令配置的日志文件汇总**。

```nginx
location /rewrite_log {
    rewrite_log on;    # 开启重写日志
	error_log logs /error.log notice;   # 切换为 notice 模式，因为只支持这个模式
    return 200 '开启了重写日志';
}
```

> ### Rewrite的案例
>
> ### 域名跳转
>
> 问题分析
>
> 先来看一个效果，如果我们想访问京东网站，大家都知道我们可以输入 `www.jd.com`，但是同样的我们也可以输入 `www.360buy.com` 同样也都能访问到京东网站。这个其实是因为京东刚开始的时候域名就是 `www.360buy.com`，后面由于各种原因把自己的域名换成了 `www.jd.com`，虽然说域名改变了，但是对于以前只记住了 `www.360buy.com` 的用户来说，我们如何把这部分用户也迁移到我们新域名的访问上来，针对于这个问题，我们就可以使用 Nginx 中 Rewrite 的域名跳转来解决。
>
> ### 环境准备
>
> - 准备两个域名 www.360buy.com | www.jd.com
>
> ```sh
> vim /etc/hosts
> ```
>
> 添加内容：
>
> ```nginx
> 192.168.200.133 www.360buy.com
> 192.168.200.133 www.jd.com
> ```
>
> - 在 `/usr/local/nginx/html/test` 目录下创建一个访问页面 frx.html
>
> 添加内容：
>
> ```html
> <html>
> 	<title></title>
> 	<body>
> 		<h1>欢迎来到我的网站</h1>
> 	</body>
> </html>
> ```
>
> - 通过 Nginx 实现当访问 `www.frx.com` 访问到 frx.html 页面
>
> 
>
> ```nginx
> server {
> 	listen 80;
> 	server_name www.frx.com;
> 	location / {
> 		root /usr/local/nginx/html/;
> 		index frx.html;
> 	}
> }
> ```
>
> 通过 Rewrite 完成将 `www.360buy.com` 的请求跳转到 `www.jd.com`
>
> ```nginx
> server {
> 	listen 80;
> 	server_name www.360buy.com;
> 	rewrite ^/ http://www.jd.com permanent;   # 永久重定向
> }
> ```
>
> **问题描述：如何在域名跳转的过程中携带请求的 URI？**
>
> 比如 `www.360buy.com?part=显示器` 变成 `www.jd.com?part=显示器`
>
> - 修改配置信息
>
> 
>
> ```nginx
> server {
> 	listen 80;
> 	server_name www.itheima.com;
> 	rewrite ^(.*) http://www.hm.com$1 permanent;
> }
> ```
>
> 括号里是 `www.itheima.com` 后面出现 0 次或 多次不以 \n（换行）结尾的值，该值赋给 $1。
>
> **问题描述：我们除了上述说的只有 www.jd.com、www.360buy.com，其实还有我们也可以通过 www.jingdong.com 来访问，那么如何通过 Rewrite 来实现多个域名的跳转?**
>
> - 添加域名
>
> ```sh
> # 打开 hosts 文件
> vim /etc/hosts
> 
> # 添加域名
> 192.168.200.133 www.jingdong.com
> ```
>
> - 修改配置信息
>
> ```nginx
> server{
> 	listen 80;
> 	server_name www.360buy.com www.jingdong.com;
> 	rewrite ^(.*) http://www.jd.com$1 permanent;
> }
> ```
>
> 多个 server_name 用空格隔开。



> **`rewrite` 正则表达式匹配的是 `$uri`，即当前请求的、规范化后的、不包括协议、域名、端口和查询字符串（`?`后面的参数）的URI部分。**
>
> 一个完整的URL（例如 `https://www.example.com:443/products/index.html?category=books&page=2#review`）可以分解为以下几个部分：
>
> - **`https://`** - 协议 (Scheme)
> - **`www.example.com`** - 域名 (Server Name)
> - **`:443`** - 端口 (Port，HTTPS默认隐藏)
> - **`/products/index.html`** - **URI (这是 `rewrite` 匹配的目标)**
> - **`?category=books&page=2`** - 查询字符串 (Query String，即 `$args`)
> - **`#review`** - 锚点 (Fragment，服务器不接收)
>
> `rewrite` 指令中的正则表达式 `regex` 只会对上面加粗的 **URI** 部分（即 `/products/index.html`）进行匹配。
>
> ### 重要特性与例子
>
> #### 1. 匹配纯路径（最常见的场景）
>
> **配置：**
>
> ```nginx
> server {
>     listen 80;
>     server_name example.com;
> 
>     location / {
>         rewrite ^/old-page$ /new-page permanent;
>     }
> }
> ```
>
> - **访问：** `http://example.com/old-page`
> - **`$uri` 是什么：** `/old-page`
> - **正则 `^/old-page$` 匹配什么：** 它匹配**精确为 `/old-page`** 的URI。
> - **结果：** 用户被 301 重定向到 `http://example.com/new-page`。
>
> #### 2. 使用正则分组捕获动态部分
>
> **配置：**
>
> nginx
>
> ```
> rewrite ^/user/([0-9]+)/profile$ /profile.php?user_id=$1 last;
> ```
>
> - **访问：** `http://example.com/user/12345/profile`
> - **`$uri` 是什么：** `/user/12345/profile`
> - **正则匹配：**
>   - `^/user/` - 匹配开头
>   - `([0-9]+)` - **捕获组**，匹配并捕获一个以上的数字（这里是 `12345`）
>   - `/profile$` - 匹配结尾
> - **结果：** 内部重写（`last`标志）到 `/profile.php?user_id=12345`。这里捕获的 `$1` 被用在了 `replacement` 中。
>
> #### 3. **不匹配查询字符串（Query String）**
>
> 这是一个非常关键且容易混淆的点。**`rewrite` 的 `regex` 默认不匹配 `?` 之后的查询参数。**
>
> 查询参数存储在变量 `$args` 中，如果需要处理，需要单独引用。
>
> **配置：**
>
> nginx
>
> ```nginx
> # 尝试匹配带参数的URL - 这是错误的写法！
> rewrite ^/search/(.*)$ /search.php?q=$1? permanent;
> ```
>
> - **访问：** `http://example.com/search/nginx%20tutorial?sort=date`
> - **`$uri` 是什么：** `/search/nginx%20tutorial` （注意，`?sort=date` 是 `$args`，不属于 `$uri`）
> - **会发生什么：**
>   - 规则会匹配 `/search/nginx%20tutorial`。
>   - 重定向到 `/search.php?q=nginx tutorial?`。（注意，这里多了一个问号，并且丢失了原始的 `sort=date` 参数）
>
> **正确的做法：** 如果要保留原始查询字符串，应该在 `replacement` 后面加上 `$is_args$args`。
>
> ```nginx
> # 正确做法：保留原查询参数
> rewrite ^/search/(.*)$ /search.php?q=$1 permanent;
> # 或者更精确地，使用$is_args$args来显式传递原参数
> rewrite ^/search/(.*)$ /search.php?q=$1$is_args$args permanent;
> ```
>
> - `$is_args`： 如果存在原始查询参数，它的值是 `?`，否则是空字符串。
> - `$args`： 就是原始的查询字符串（例如 `sort=date`）。
> - **结果：** 用户被重定向到 `http://example.com/search.php?q=nginx+tutorial&sort=date`，原始参数得以保留。
>
> 

### 域名镜像

上述案例中，将 `www.360buy.com` 和 `www.jingdong.com` 都能跳转到 `www.jd.com`，那么 `www.jd.com` 我们就可以把它起名叫主域名，其他两个就是我们所说的镜像域名，当然如果我们不想把整个网站做镜像，只想为其中某一个子目录下的资源做镜像，比如用户可以跳到首页 Web下，而管理员跳转到后台 Web，我们可以在 location 块中配置 Rewrite 功能。

比如：

```nginx
server {
	listen 80;
	server_name rewrite.myweb.com;
	location ^~ /user {
		rewrite ^/user(.*) http://www.myweb.com/index$1 last;  # 用户跳到首页
	}
	location ^~ /manage {
		rewrite ^/manage(.*) http://www.myweb.com/manage$1 last;  # 管理员跳到后台
	}
}
```

### 独立域名

一个完整的项目包含多个模块，比如购物网站有商品商品搜索模块、商品详情模块已经购物车模块等，那么我们如何为每一个模块设置独立的域名。

需求：

- `http://search.product.com`：访问商品搜索模块
- `http://item.product.com`：访问商品详情模块
- `http://cart.product.com`：访问商品购物车模块



```nginx
server{
	listen 80;
	server_name search.product.com;
	rewrite ^(.*) http://www.shop.com/search$1 last;
}
server{
	listen 81;
	server_name item.product.com;
	rewrite ^(.*) http://www.shop.com/item$1 last;
}
server{
	listen 82;
	server_name cart.product.com;
	rewrite ^(.*) http://www.shop.com/cart$1 last;
}
```

### 自动加『/』

有时候访问的地址要求后面以 `/` 结尾，那么我们需要解决如果用户忘记输入 `/`，Nginx 就会自动加上 `/`。

通过一个例子来演示问题：

```nginx
server {
	listen	80;
	server_name localhost;
	location / {
		root html;
		index index.html;
	}
}
```

要想访问上述资源，很简单，只需要通过 `http://192.168.200.133` 直接就能访问，地址后面不需要加 /，但是如果将上述的配置修改为如下内容:

```nginx
server {
	listen	80;
	server_name localhost;
	location /frx {
		root html;
		index index.html;
	}
}
```

这个时候，要想访问上述资源，按照上述的访问方式，我们可以通过 `http://192.168.200.133/frx/` 来访问，



如果不加斜杠，Nginx 服务器内部会自动做一个 301 的重定向，重定向的地址会有一个指令叫 `server_name_in_redirect` 来决定重定向的地址：

- 如果该指令为 on

  重定向的地址为：`http://server_name/目录名/`

- 如果该指令为 off

  重定向的地址为：`http://原URL中的域名/目录名/`

所以就拿刚才的地址来说，访问 `http://192.168.200.133/frx` 如果不加斜杠，那么按照上述规则：

- 如果指令 `server_name_in_redirect` 为 on，则 301 重定向地址变为 `http://localhost/frx/`，IP 发生改变，地址出现了问题
- 如果指令 `server_name_in_redirect` 为 off，则 301 重定向地址变为 `http://192.168.200.133/frx/`。这个符合我们的期望

注意 `server_name_in_redirect` 指令在 Nginx 的 0.8.48 版本之前默认都是 on，之后改成了 off，所以现在我们这个版本不需要考虑这个问题，但是如果是 0.8.48 以前的版本并且 server_name_in_redirect 设置为 on，我们如何通过 Rewrite 来解决这个问题？

解决方案

我们可以使用 Rewrite 功能为末尾没有斜杠的 URL 自动添加一个斜杠



```nginx
server {
	listen	80;
	server_name localhost;
	server_name_in_redirect on;
	location /frx {
		if (-d $request_filename){   # 如果请求的资源目录存在
			rewrite ^/(.*)([^/])$ http://$host/$1$2/ permanent; # $2 获取第二个括号的值：/
		}
	}
}
```

$1 是第一个括号的值，$2 是第二个括号的值。

### 合并目录

搜索引擎优化(SEO)是一种利用搜索引擎的搜索规则来提供目的网站的有关搜索引擎内排名的方式。

我们在创建自己的站点时，可以通过很多中方式来有效的提供搜索引擎优化的程度。其中有一项就包含 URL 的目录层级，一般不要超过三层，否则的话不利于搜索引擎的搜索，也给客户端的输入带来了负担，但是将所有的文件放在一个目录下，又会导致文件资源管理混乱，并且访问文件的速度也会随着文件增多而慢下来，这两个问题是相互矛盾的，那么使用 Rewrite 如何解决这些问题呢？

举例，网站中有一个资源文件的访问路径 `/server/11/22/33/44/20.html`，也就是说 20.html 存在于第 5 级目录下，如果想要访问该资源文件，客户端的 URL 地址就要写成 `http://www.web.com/server/11/22/33/44/20.html`，并且在配置文件进行如下配置：



```nginx
server {
	listen 80;
	server_name www.web.com;
	location /server {
		root html;
	}
}
```

但是这个是非常不利于 SEO 搜索引擎优化的，同时客户端也不好记。使用 Rewrite 的正则表达式，我们可以进行如下配置：



```nginx
server {
    listen 80;
    server_name www.web.com;
    location /server {
        rewrite ^/server-([0-9]+)-([0-9]+)-([0-9]+)-([0-9]+)\.html$  /server/$1/$2/$3/$4/$5.html last;
    }
}
```

这样配置后，客户端只需要输入 `http://www.web.com/server-11-22-33-44-20.html` 就可以访问到 20.html 页面了。

这里也充分利用了 Rewrite 指令支持正则表达式的特性。

### 多级域名

当你配置了多级域名，如二级域名 `xxx.frxcat.fun`，并且静态资源目录恰好和二级域名的 `xxx` 可以匹配，则可以使用正则表达式进行匹配，日后，如果又多个 `xxx`，则再创建对应的该目录即可。



```nginx
server {
	listen 80;
	server_name ~^(.+)?.frxcat.fun$;
    index idnex.html;
    if ($host = frxcat.fun){
        rewrite ^(.*)$ https://www.frxcat.fun$2 permanent;
    }
    root /data/html/$1/;
}
```

这样访问 `docs.frxcat.fun`，自动去 `/data/html/docs/` 目录下找到 index.html，如果是 `bing.youngkbt.cn`，则会去 `/data/html/bing/` 目录下找到 idnex.html，以此类推。

if 语句的作用是将 `frxcat.fun` 重定向到 `www.frxcat.fun`，这样既解决了网站的主目录访问，又可以增加 SEO 中对 `www.frxcat.fun` 的域名权重。

### 防盗链

防盗链之前我们已经介绍过了相关的知识，在 Rewrite 中的防盗链和之前将的原理其实都是一样的，只不过通过 Rewrite 可以将防盗链的功能进行完善下，当出现防盗链的情况，我们可以使用 Rewrite 将请求转发到自定义的一张图片和页面，给用户比较好的提示信息。

下面有两个配置实例：

- 根据文件类型实现防盗链配置：



```nginx
server{
	listen 80;
	server_name www.web.com;
	locatin ~* ^.+\.(gif|jpg|png|swf|flv|rar|zip)$ {
		valid_referers none blocked server_names *.web.com; # server_names 后指定具体的域名或者 IP
		if ($invalid_referer){
			rewrite ^/ http://www.web.com/images/forbidden.png;  # 跳转到默认地址
		}
	}
}
```

- 根据目录实现防盗链配置：



```nginx
server{
	listen 80;
	server_name www.web.com;
	location /file {
		root /server/file;  # 资源在 server 目录下的 file 目录里
		valid_referers none blocked server_names *.web.com; # server_names 后指定具体的域名或者 IP
		if ($invalid_referer){
			rewrite ^/ http://www.web.com/images/forbidden.png;  # 跳转到 file 目录下的图片
		}
	}
}
```

## 访问限流

我们构建网站是为了让用户访问它们，我们希望用于合法访问。所以不得不采取一些措施限制滥用访问的用户。这种滥用指的是从同一 IP 每秒到服务器请求的连接数。因为这可能是在同一时间内，世界各地的多台机器上的爬虫机器人多次尝试爬取网站的内容。



```nginx
# 限制用户连接数来预防 DOS 攻击
limit_conn_zone $binary_remote_addr zone=perip:10m;
limit_conn_zone $server_name zone=perserver:10m;
# 限制同一客户端 ip 最大并发连接数
limit_conn perip 2;
# 限制同一server最大并发连接数
limit_conn perserver 20;
# 限制下载速度，根据自身服务器带宽配置
limit_rate 300k; 
```

## 链接超时

长时间占着连接资源不释放，最终会导致请求的堆积，Nginx 处理请求效率大大降低。所以我们对连接的控制都要注意设置超时时间，通过超时机制自动回收资源、避免资源浪费。

```nginx
# 客户端、服务端设置
server_names_hash_bucket_size 128;
server_names_hash_max_size 512;
# 长连接超时配置
keepalive_timeout  65;
client_header_timeout 15s;
client_body_timeout 15s;
send_timeout 60s;

# 代理设置
# 与后端服务器建立连接的超时时间。注意这个一般不能大于 75 秒
proxy_connect_timeout 30s;
proxy_send_timeout 120s;
# 从后端服务器读取响应的超时
proxy_read_timeout 120s;
```

## [#](https://frxcat.fun/middleware/Nginx/NginxStatic_resource_access/#html引入)HTML引入

我们编写 .html 文件的时候，难免需要引入 css 和 js 文件，如果是在本地，那么引入非常简单，直接相对路径即可，但是部署到 Nginx 时，相对路径不再是相对 html 文件的目录，所以生产环境和开发环境的引入格式不一样。

在 Nginx 中的 .html 文件，引入 css 和 js，要加上 `/` 作为开头，`/` 代表 Nginx 的根目录，即配置文件 `location /` 的指定的 root 路径。

比如 Nginx 的配置文件内容如下：



```nginx
server {
	listen 80;
    server_name localhost;
    
    location / {
        root /usr/local/nginx/html; # 静态文件根目录
        index idnex.html;
    }
}
```

有一个 aa.html 在 `/usr/local/nginx/html/test` 目录下，并且 aa.html 引入了 aa.css 和 aa.js，两个静态文件在 aa.html 所在目录的 static 文件夹里。



```sh
/usr/local/nginx/html/test 目录
├── a.html
├── static
│	├── a.css
│	├── a.js
```

在本地环境，我们可以这样写：



```html
<link rel="stylesheet" href="static/aa.css" />
<script src="static/aa.js"></script>
```

但是部署到 Nginx 后，这样写会找不到这两个资源，因为 `/` 触发 `location /`，进入 `/usr/local/nginx/html` 目录，而这两个文件在 `/usr/local/nginx/html/test/static` 目录下，所以我们部署到 Nginx 后，需要修改为：



```html
<link rel="stylesheet" href="/test/static/aa.css" />
<script src="/test/static/aa.js"></script>
```





## 反向代理概述

关于正向代理和反向代理，在 [Nginx - 介绍](https://frxcat.fun/pages/9551ee/#名词解释) 已经通过一张图详细的介绍过了，简而言之就是正向代理代理的对象是客户端，反向代理代理的是服务端，这是两者之间最大的区别。

Nginx 即可以实现正向代理，也可以实现反向代理。

我们先来通过一个小案例演示下 Nginx 正向代理的简单应用。

### 正向代理

先提需求：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.53oygoz19gc0.webp)

**服务端的设置(133)：**

```nginx
http {
    log_format main 'client send request=>clientIp=$remote_addr serverIp=$host';
	server{
		listen 80;
		server_name	localhost;
		access_log logs/access.log main;
		location / {
			root html;
			index index.html index.htm;
		}
	}
}
```

**使用客户端访问服务端：`http://192.168.200.133`，打开日志查看结果**

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.16ifxal7swyk.webp)

**代理服务器设置(146)：**

```nginx
server {
    listen  82;
    resolver 8.8.8.8;   # 设置 DNS 的 IP，用来解析 proxy_pass 中的域名
    location / {
        proxy_pass http://$host$request_uri;   # proxy_pass 实现正向代理
    }
}
```

`proxy_pass` 后面有讲解。

配置代理服务器的 IP(192.168.200.146)和 Nginx 配置监听的端口(82)

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.3yx93qbazra0.webp)

设置完成后，再次通过浏览器访问服务端

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.6litsat09000.webp)

通过对比，上下两次的日志记录，会发现虽然我们是客户端访问服务端，但是使用了代理，那么服务端能看到的只是代理发送过去的请求，这样就使用 Nginx 实现了正向代理的设置。

但是 Nginx 正向代理，在实际的应用中不是特别多，所以我们简单了解下，接下来我们继续学习 Nginx 的反向代理，这是 Nginx 比较重要的一个功能。

> 在 Nginx 中，**`$host` 变量捕获的是客户端请求中的域名部分**。
>
> ### `$host` 变量的具体行为
>
> `$host` 变量会按以下优先级获取域名信息：
>
> 1. 首先，尝试获取 HTTP 请求头中的 `Host` 字段的值。
> 2. 如果 `Host` 头不存在，则使用与服务器建立连接的 IP 地址对应的 `server_name`。
> 3. **重要**：它会**排除请求行中的端口号**。
>
> ### 举例说明
>
> 假设您的 Nginx 配置中有一个 `server` 块，其 `server_name` 为 `example.com` 和 `www.example.com`。
>
> | 客户端请求的 URL                              | `$host` 的值      | 解释                                            |
> | :-------------------------------------------- | :---------------- | :---------------------------------------------- |
> | `http://example.com/page`                     | `example.com`     | 获取 `Host` 头，并去掉端口（默认80端口）        |
> | `https://www.example.com:443/page`            | `www.example.com` | 获取 `Host` 头，并去掉端口（默认443端口）       |
> | `http://example.com:8080/page`                | `example.com`     | 获取 `Host` 头，并**去掉指定的端口号（:8080）** |
> | 直接通过 IP 访问：`http://192.168.1.100/page` | `192.168.1.100`   | `Host` 头是 IP 地址，所以 `$host` 也是 IP 地址  |
>
> ------
>
> ### 与其他相关变量的区别
>
> | 变量               | 示例值                 | 说明                                                         |
> | :----------------- | :--------------------- | :----------------------------------------------------------- |
> | **`$host`**        | `www.example.com`      | **最常用**。包含域名，**不包含**端口号。                     |
> | **`$http_host`**   | `www.example.com:8080` | 直接取自 HTTP 请求头中的 `Host` 字段原样值，**包含端口号**（如果请求中有的话）。 |
> | **`$server_name`** | `example.com`          | 在 Nginx 配置文件中由 `server_name` 指令设置的**服务器名称**，是服务端的配置值，而非客户端的请求值。 |
> | **`$server_addr`** | `192.168.1.100`        | 服务器接受此请求的 IP 地址。                                 |

### 反向代理

Nginx 反向代理模块的指令是由 `ngx_http_proxy_module` 模块进行解析，该模块在安装 Nginx 的时候已经自动加载到 Nginx 中了，接下来我们把反向代理中的常用指令一一介绍下：

- proxy_pass：配置代理的服务器地址
- proxy_set_header：转发给被代理服务器时，设置一些请求头信息
- proxy_redirect：防止客户端可以看到被代理服务器的地址

这里只介绍三个指令，关于反向代理的指令非常多，想要了解更多，请前往 [Nginx 反向代理文档(opens new window)](https://nginx.org/en/docs/stream/ngx_stream_proxy_module.html)

#### proxy_pass

该指令用来设置被代理服务器地址，可以是主机名称、IP 地址加端口号形式，没有默认值。

| 语法              | 默认值 | 位置     |
| ----------------- | ------ | -------- |
| proxy_pass <URL>; | —      | location |

`URL`：为要设置的被代理服务器地址，包含传输协议(`http`、`https://`)、主机名称或 IP 地址加端口号、URI 等要素。

例如：

```nginx
proxy_pass http://www.baidu.com;

# 例子
location /server {
    # 结尾不加斜杠
    proxy_pass http://192.168.200.146; # 
    # 访问的是：http://192.168.200.146/server/index.html

    # 结尾加斜杠
    proxy_pass http://192.168.200.146/;
    # 访问的是：http://192.168.200.146/index.html
}
```

> **实例**

准备两台服务器或者按照 Linux 系统的虚拟机，这里是 `192.168.200.133` 和 `192.168.200.146`，为了方便，我们称前者为服务器 A，后者为服务器 B。

1. 在服务器 A 的 Nginx 配置文件添加如下内容：

```nginx
http {
    # ......
    server {
        listen 8080;
        server_name localhost;
        location / {
            proxy_pass http://192.168.200.146;
        }
    }
    # ......
}
```

当客户端请求服务器 A `http://192.168.200.133`，它会转发给服务器 B，此时的服务器 A 就是一个代理的角色。

访问服务器 A，我们看到 Nginx 的欢迎界面其实是服务器 B 的 Nginx，可以在服务器 B 的 Nginx 欢迎页面添加新的内容：`I am 146`，再次访问服务器 A，效果如图：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.4ai1u3uxs940.webp)

**在编写 proxy_pass 的时候，后面的值要不要加 /?**

这是看情况的。

接下来通过例子来说明提到的问题：

```nginx
server {
	listen 80;
	server_name localhost;
	location / {
        # 下面两个地址加不加斜杠，效果都一样，因为 location 后的 / 会添加在代理地址后面
		proxy_pass http://192.168.200.146;
		proxy_pass http://192.168.200.146/;
	}
}

server{
	listen 80;
	server_name localhost;
	location /server {
        # 下面两个地址必须加斜杠，因为 location 后的 /server 会添加在代理地址后面，第一个将没有 / 结尾
		#proxy_pass http://192.168.200.146;
		proxy_pass http://192.168.200.146/;
	}
}
# 上面的 location：当客户端访问 http://localhost/server/index.html
# 第一个 proxy_pass 就变成了 http://localhost/server/index.html
# 第二个 proxy_pass 就变成了 http://localhost/index.html 效果就不一样了。
```

- 第一个 location（第 4 行代码）：当客户端访问 `http://localhost/index.html`，两个 `proxy_pass` 效果是一样的，因为 location 后的 `/` 会添加在代理地址后面，所以有没有 `/`，效果都一样。

- 第一个 location（第 14 行代码）：当客户端访问 `http://localhost/server/index.html`，这个时候，第一个 proxy_pass 就变成了 `http://192.168.200.146/server/index.html`，第二个 proxy_pass 就变成了 `http://192.168.200.146/index.html` 效果就不一样了

  如果不以 `/` 结尾，则 location 后的 `/server` 会添加在地址后面，所以第一个 proxy_pass 因为没有 `/` 结尾而被加上 `/server`，而第二个自带了 `/` ，所以不会添加 `/server`。

上面的例子仅仅针对：访问任意请求如 `/server` 时，想要代理到其他服务器的首页，则加 `/`，否则你如果真的想访问 `/server` 下的资源，那么不要加 `/`。

所以加了 `/` 后，请求的是服务器根目录下的资源。

#### proxy_set_header

该指令可以更改 Nginx 服务器接收到的客户端请求的请求头信息，然后将新的请求头发送给代理的服务器。默认值是发送代理服务器的地址和 close。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.4bt7qzhbz1i0.webp)

| 语法                              | 默认值                                                       | 位置                   |
| --------------------------------- | ------------------------------------------------------------ | ---------------------- |
| proxy_set_header <field> <value>; | proxy_set_header Host $proxy_host; <br />proxy_set_header Connection close; | http、server、location |



需要注意的是，如果想要看到结果，必须在被代理的服务器上来获取添加的头信息。

> **实例**

被代理服务器：服务器 B `192.168.200.146` 的 Nginx 配置文件内容：

```nginx
server {
    listen  8080;
    server_name localhost;
    default_type text/plain;
    return 200 $http_username;    # 获取代理服务器发送过来的 http 请求头的 username 值
}
```

代理服务器: 服务器 A `192.168.200.133` 的 Nginx 配置文件内容：

```nginx
server {
    listen  8080;
    server_name localhost;
    location /server {           # 访问 /server 触发代理
        proxy_pass http://192.168.200.146:8080/;  # 配置服务器 B 的地址
        proxy_set_header username TOM;  # 发送 key 为 username，value 为 TOM 的请求头给服务器 B
    }
}
```

访问测试

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.398ksbw6rv20.webp)

客户端访问的是服务器 A，服务器 A 会将请求转发给服务器 B，服务器 B 返回打印 TOM 的页面给服务器 A，服务器 A 最后返回给客户端。

#### proxy_redirect

该指令是用来重置头信息中的『 Location 』和『 Refresh 』的值，防止客户端可以看到被代理服务器的地址。

因为客户端看到的返回结果是『 Location 』和『 Refresh 』的值，所以在到达代理服务器的时，将两个值修改掉，防止客户端直接看到被代理服务器的地址。

| 语法                                                         | 默认值                  | 位置                   |
| ------------------------------------------------------------ | ----------------------- | ---------------------- |
| proxy_redirect redirect replacement;<br />proxy_redirect default; <br />proxy_redirect off; | proxy_redirect default; | http、server、location |



> **为什么要用该指令？**

首先说明一下思路：客户端通过代理服务器 A 访问服务器 B 的资源，但是服务器 B 不存在该资源，则会报错。

此时我们不希望它直接返回报错页面给客户端，我们希望服务器 B 返回的是它的欢迎页面。那么如何做呢？

- 首先在服务器 B 进行判断是否存在资源，不存在则返回自己的欢迎页面，即重定向到自己的欢迎页面地址并返回，此时浏览器的地址将会发生改变
- 代理服务器 A 收到服务器 B 的欢迎页面和地址，但是我们不能直接返回给客户端，因为它会暴露服务器 B 的地址，这是重定向的原因

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.a4tptcx6lbk.webp)

- 此时用到 `proxy_redirect` 指令，重置服务器 B 返回过来的『 Location 』和『 Refresh 』值，将两个值改为代理服务器 A 的某个地址
- 因为改为了代理服务器 A 的某个地址，所以代理服务器 A 根据这个地址又去获取理服务器 B 的欢迎页面地址，返回给客户端

很绕，简单总结下：客户端通过 A 找 B 不存在的资源，B 不想返回报错页面，于是重定向到自己的欢迎页面地址并返回给 A，A 收到了页面和地址（正常情况不要接收地址，只接收页面），发现不能暴露 B 的地址，于是修改接收的 B 的地址为自己的某一个地址，这个地址会重新发送请求去获取 B 的欢迎页面地址，然后返回给客户端。

这里要明白 B 返回的是重定向后的欢迎页面，重定向后，浏览器地址栏会变成重定向的地址，所以 A 才会以自己的地址转发获取到 B 的地址，所以最后浏览器显示 A 的地址，看到的却是 B 的欢迎页面。

代码：



服务端 B `192.168.200.146`

```nginx
server {
    listen  8081;
    server_name localhost;
    if (!-f $request_filename){
    	return 302 http://192.168.200.146;   #  2.如果请求的资源不存在，则重定向到服务器 B
    }
}
```



代理服务端 A `192.168.200.133`

```nginx
server {
	listen  8081;
	server_name localhost;
	location / {
		proxy_pass http://192.168.200.146:8081/;  # 1.转发给服务器 B
		proxy_redirect http://192.168.200.146 http://192.168.200.133; # 3.修改服务器 B 的地址
	}
}
# 该 server 去请求服务器 B 的欢迎页面
server {
	listen  80;
	server_name 192.168.200.133;
	location / {
		proxy_pass http://192.168.200.146;  # 4.重新发送请求给服务器 B，获取欢迎页面
	}
}
```

第 6 行代码，当服务器 B 返回的是 `http://192.168.200.146`，为了不让它出现在浏览器的地址栏上，我们需要利用 `proxy_redirect` 将它修改为代理服务器 A 的地址，这个地址会以自己的地址重新访问服务器 B 的欢迎页面，最后返回给客户端。

**该指令的三组选项**

- `proxy_redirect redirect replacement;`

  - redirect：被代理服务器返回的 Location 值
  - replacement：要替换 Location 的值

- `proxy_redirect default;`

  - default：相比较第一组选项，default 仅仅提供了 `redirect` 和 `replacement` 的默认值

    将本范围 location 块的 uri 变量作为 replacement。

    将 proxy_pass 变量作为 redwadairect



```nginx
server {
	listen  8081;
	server_name localhost;
	location /server { 
		proxy_pass http://192.168.200.146:8081/;
		proxy_redirect default;  # redirect 是 proxy_pass 的值：http://192.168.200.146:8081/
        						 # replacement 是 location 后的值：/server
        # 等价于：proxy_redirect http://192.168.200.146:8081/ /server
	}
}
```

- `proxy_redirect off;`

  关闭 proxy_redirect 的功能

#### 反向代理实战

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.59ecpsonzxc0.webp)

服务器 1，2，3 存在两种情况

- 第一种情况: 三台服务器的内容不一样
- 第二种情况: 三台服务器的内容是一样

第一种情况

- 如果服务器 1、服务器 2 和服务器 3 的内容不一样，那我们可以根据用户请求来分发到不同的服务器。

  服务器有限，只能以三个端口模拟三台服务器，实际上是一个 IP 对应一个服务器。

  代理服务器配置文件内容：

```nginx
# 代理服务器
server {
    listen  8082;
    server_name  localhost; 
    location /server1 {    
        proxy_pass http://192.168.200.146:9001/;   # 代理 server1
    }
    location /server2 {
        proxy_pass http://192.168.200.146:9002/;   # 代理 server2
    }
    location /server3 {
        proxy_pass http://192.168.200.146:9003/;   # 代理 server3
    }
}
```

服务器配置文件内容：

```nginx
# 服务器
# server1
server {
    listen  9001;
    server_name  localhost;
    default_type text/html;
    return 200 '<h1>192.168.200.146:9001</h1>'
}
# server2
server {
    listen  9002;
    server_name  localhost;
    default_type text/html;
    return 200 '<h1>192.168.200.146:9002</h1>'
}
# server3
server {
    listen  9003;
    server_name  localhost;
    default_type text/html;
    return 200 '<h1>192.168.200.146:9003</h1>'
}
```

- 如果服务器 1、服务器 2 和服务器 3 的内容是一样的，该如何处理?

  请看负载均衡的相关内容，里面将进行详细的介绍。

#### 斜杠总结

这里将发送 `http://192.168.199.27/frx/xu` 请求。

> ### 两大核心规则
>
> `proxy_pass` 指令的最终转发地址（上游服务器接收到的 URI）由以下公式决定：
>
> ```
> 最终URI = proxy_pass指令中指定的URI + 原始请求URI去除与location匹配的部分
> ```
>
> 关键在于 **`proxy_pass` 指令中是否包含 URI**（即 IP:port 之后的部分）。
>
> #### 规则一：当 `proxy_pass` 不包含 URI（即只有 IP:port）
>
> Nginx 会将 **完整的原始请求路径** 追加到 `proxy_pass` 指定的地址后面。
>
> - **行为**：原封不动地转发路径。
> - **公式**： `最终URI = 原始请求URI`
>
> **案例解析（对应您的不带字符串情况）：**
>
> | 案例 | proxy_pass               | 是否含URI    | 原始请求  | location | 最终转发路径 | 解释                                                         |
> | :--- | :----------------------- | :----------- | :-------- | -------- | :----------- | :----------------------------------------------------------- |
> | 1    | `http://192.168.199.27`  | **否**       | `/frx/xu` | /frx     | `/frx/xu`    | 规则一：原始请求路径 `/frx/xu` 被完整地追加到 `http://192.168.199.27` 之后。 |
> | 2    | `http://192.168.199.27`  | **否**       | `/frx/xu` | /frx/    | `/frx/xu`    | 规则一：同上。**注意**：虽然 location 是 `/frx/`，但因为 proxy_pass 不含 URI，规则一优先，路径依旧完整转发。 |
> | 3    | `http://192.168.199.27/` | **是** (`/`) | `/frx/xu` | /frx     | `//xu`       | **规则二**：因为 proxy_pass 包含了 URI (`/`)，所以它进入了规则二。匹配部分 `/frx` 被移除，剩下 `/xu`，拼接到 `/` 后得到 `//xu`。 |
> | 4    | `http://192.168.199.27/` | **是** (`/`) | `/frx/xu` | /frx/    | `/xu`        | **规则二**：proxy_pass 含 URI (`/`)。匹配部分 `/frx/` 被移除，剩下 `xu`，拼接到 `/` 后得到 `/xu`。 |
>
> #### 规则二：当 `proxy_pass` 包含 URI（即 IP:port/xxx）
>
> Nginx 会将 **原始请求路径中与 `location` 匹配的部分** 替换为 `proxy_pass` 中指定的 URI。
>
> - **行为**：替换路径。
> - **公式**： `最终URI = proxy_pass的URI部分 + (原始请求URI - location的匹配部分)`
>
> **案例解析（对应您的带字符串情况）：**
>
> | 案例 | location | proxy_pass                    | 原始请求  | location | 最终转发路径 | 解释                                                         |
> | :--- | :------- | :---------------------------- | :-------- | -------- | :----------- | :----------------------------------------------------------- |
> | 1    | `/frx`   | `http://192.168.199.27/bing`  | `/frx/xu` | /frx     | `/bing/xu`   | **规则二**：匹配部分 `/frx` 被移除，剩下 `/xu`。拼接到 `bing` 后得到 `/bing/xu`。**注意**：`bing` 和 `/xu` 之间会产生一个 `/`，这是路径的自然拼接。 |
> | 2    | `/frx/`  | `http://192.168.199.27/bing`  | `/frx/xu` | /frx/    | `/bingxu`    | **规则二**：匹配部分 `/frx/` 被**精确移除**，剩下 `xu`（**注意**：因为 location 是 `/frx/`，它匹配的是整个前缀，所以移除后剩下的部分 `xu` 开头没有 `/`）。拼接到 `bing` 后得到 `/bingxu`。 |
> | 3    | `/frx`   | `http://192.168.199.27/bing/` | `/frx/xu` | /frx     | `/bing//xu`  | **规则二**：匹配部分 `/frx` 被移除，剩下 `/xu`。拼接到 `bing/` 后得到 `/bing//xu`（`bing/` + `/xu`）。 |
> | 4    | `/frx/`  | `http://192.168.199.27/bing/` | `/frx/xu` | /frx/    | `/bing/xu`   | **规则二**：匹配部分 `/frx/` 被移除，剩下 `xu`。拼接到 `bing/` 后得到 `/bing/xu`（`bing/` + `xu`）。 |
>
> ------
>
> ### 记忆技巧与最终总结
>
> 1. **先看 `proxy_pass`**：这是最重要的。看 IP:port 后面有没有东西（`/` 也算东西！）。
>    - **没东西**：走规则一，**原样转发**。此时 `location` 后面有没有 `/` 不影响最终结果。
>    - **有东西**：走规则二，**替换路径**。此时需要仔细处理。
> 2. **走规则二时，再看 `location`**：
>    - 如果 `location` 匹配块以 `/` 结尾（如 `location /frx/`），则它会在原始请求中**精确地剥离掉整个匹配串**（包括最后的 `/`）。剩下的部分如果以 `/` 开头，那拼接后就会少一个 `/`；如果不以 `/` 开头，就会和 proxy_pass 的 URI 直接相连（如案例2）。
>    - 如果 `location` 匹配块**不**以 `/` 结尾（如 `location /frx`），则它只会剥离掉匹配的前缀，原始请求剩余的部分（包括可能的 `/`）会保留下来。

## Nginx安全控制

关于 Web 服务器的安全是比较大的一个话题，里面所涉及的内容很多，Nginx 反向代理是如何来提升 Web 服务器的安全呢？

答案是：安全隔离。

### 什么是安全隔离

通过代理分开了客户端到应用程序服务器端的连接，实现了安全措施。在反向代理之前设置防火墙，仅留一个入口供代理服务器访问。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.3vhdv63nf9q0.webp)

### 如何使用SSL对流量加密

翻译成大家能熟悉的说法就是将我们常用的 http 请求转变成 https 请求，那么这两个之间的区别简单的来说两个都是 HTTP 协议，只不过 https 是身披 SSL 外壳的 http.

HTTPS 是一种通过计算机网络进行安全通信的传输协议。它经由 HTTP 进行通信，利用 SSL/TLS 建立全通信，加密数据包，确保数据的安全性。

- SSL(Secure Sockets Layer)安全套接层
- TLS(Transport Layer Security)传输层安全

上述这两个是为网络通信提供安全及数据完整性的一种安全协议，TLS 和 SSL 在传输层和应用层对网络连接进行加密。

总结来说为什么要使用 https:

- http 协议是明文传输数据，存在安全问题，而 https 是加密传输，相当于 http + SSL，并且可以防止流量劫持。

Nginx 默认不支持 https 开头的协议，如果要想使用 SSL，需要满足一个条件即需要添加一个模块 `--with-http_ssl_module`，而该模块在编译的过程中又需要 OpenSSL 的支持，OpenSSL 我们在 [Nginx 安装](https://frxcat.fun/middleware/Nginx/Nginx_install/#nginx环境安装) 已经准备好了。

### 添加SSL支持

完成 `--with-http_ssl_module` 模块的增量添加。

1将原有 `/usr/local/nginx/sbin/nginx` 进行备份

```sh
cd /usr/local/nginx/sbin
mv nginx nginx.backup
```

2查看 `configure arguments` 的配置信息，拷贝出来

```sh
nginx -V

# 拷贝 configure arguments 后面的数据
```

3进入 Nginx 的安装目录，执行 make clean 清空之前编译的内容

```nginx
cd /root/nginx/core/nginx-1.20.2

make clean
```

4使用 configure 来配置参数，添加 `ngx_http_ssl_module` 模块，记得加上第（2）步拷贝的配置信息

```sh
./configure --with-http_ssl_module # 记得添加 configure arguments 后的数据
```

5通过 make 模板进行编译

```sh
make
```

6将 objs 下面的 nginx 可执行文件移动到 `/usr/local/nginx/sbin` 下

```sh
mv /opt/nginx/core/nginx-1.20.2/objs/nginx /usr/local/nginx/sbin
```

7在源码目录（安装包目录）下执行 `make upgrade` 进行升级，这个可以实现不停机添加新模块的功能

```sh
cd /opt/nginx/core/nginx-1.20.2
make upgrade
```

### [#](https://frxcat.fun/middleware/Nginx/Nginx_Reverse_proxy/#ssl相关指令)SSL相关指令

该模块的指令都是通过 `ngx_http_ssl_module` 模块来解析的，我们上一步已经添加了。

这里只介绍常用的几个指令，了解更多指令请前往 [ngx_http_ssl_module 模块文档 (opens new window)](http://nginx.org/en/docs/http/ngx_http_ssl_module.html)。

`ssl` 指令用来在指定的服务器开启 HTTPS，默认关闭。可以使用 listen 443 ssl，这种方式更通用些。

| 语法            | 默认值   | 位置         |
| --------------- | -------- | ------------ |
| ssl <on \|off>; | ssl off; | http、server |

ssl 默认监听的是 443 端口，所以使用下面的指令和 `ssl on` 效果一致，因为下面的指令能突出 sll 的监听端口，所以建议使用它。

```nginx
server{
	listen 443 ssl;
}
```

`ssl_certificate` 指令是为当前这个虚拟主机指定一个带有 PEM 格式证书的证书。

| 语法                    | 默认值 | 位置         |
| ----------------------- | ------ | ------------ |
| ssl_certificate <file>; | —      | http、server |

`ssl_certificate_key` 指令用来指定 PEM secret key 文件的路径

| 语法                         | 默认值 | 位置         |
| ---------------------------- | ------ | ------------ |
| ssl_ceritificate_key <file>; | —      | http、server |

`ssl_session_cache` 指令用来配置用于 SSL 会话的缓存

| 语法                                                         | 默认值                  | 位置         |
| ------------------------------------------------------------ | ----------------------- | ------------ |
| ssl_sesion_cache <off \| none \| [builtin[:size]] [shared:name:size]> | ssl_session_cache none; | http、server |

选项介绍：

- `off`：严格禁止使用会话缓存：Nginx 明确告诉客户端会话不能被重用

- `none`：禁止使用会话缓存，Nginx 告诉客户端会话可以被重用，但实际上并不在缓存中存储会话参数（任性，言语同意用，行为取消用）

- `builtin`：内置 OpenSSL 缓存，仅在一个工作进程中使用。缓存大小在会话中指定。如果未给出大小，则等于 20480 个会话。使用内置缓存可能会导致内存碎片

- `shared`：所有工作进程之间共享缓存，缓存的相关信息用 name 和 size 来指定，同 name 的缓存可用于多个虚拟服务器

  name 是允许缓存的数据名，size 是允许缓存的数据大小，以字节为单位

例如：

```nginx
ssl_session_cache builtin:1000 shared:SSL:10m;  
```

10m 的 m 是兆。

`ssl_session_timeout` 指令用于开启 SSL 会话功能后，设置客户端能够反复使用储存在缓存中的会话参数时间，默认值超时时间是 5 秒

| 语法                        | 默认值                  | 位置         |
| --------------------------- | ----------------------- | ------------ |
| ssl_session_timeout <time>; | ssl_session_timeout 5m; | http、server |

`ssl_ciphers` 指令指出允许的密码，密码指定为 OpenSSL 支持的格式

| 语法                   | 默认值                        | 位置         |
| ---------------------- | ----------------------------- | ------------ |
| ssl_ciphers <ciphers>; | ssl_ciphers HIGH:!aNULL:!MD5; | http、server |

可以直接在 Linux 系统上使用 `openssl ciphers` 查看 OpenSSl 支持的格式

```sh
openssl ciphers
```

`ssl_prefer_server_ciphers` 指令指定是否服务器密码优先客户端密码，默认关闭，建议开启。

| 语法                                   | 默认值                         | 位置         |
| -------------------------------------- | ------------------------------ | ------------ |
| ssl_perfer_server_ciphers <on \| off>; | ssl_perfer_server_ciphers off; | http、server |

### SSL证书生成

**方式一：使用阿里云/腾讯云等第三方服务进行购买免费版**

需要购买域名进行证书的绑定，否则证书无法使用。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.5p2qyl61ozc0.webp)

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.5l7iqyudol00.webp)

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.6ci6w9wu24g0.webp)

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.4k9v2b9n6tm0.webp)

接着在右边弹窗进行域名绑定，填完写域名和个人信息，进入到验证信息

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.11wx9z7jit3k.webp)

点击验证，不成功则去自己的域名解析列表查看，如下，点击添加记录，进行配置，或者已经看到记录类型是 TXT，记录值和上图一样的，则说明成功。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.6qcpqete8ic0.webp)

提交审核后，点击下载

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.28m5753glfy8.webp)

下载 Nginx 服务器的证书

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.6kmq7nhonys0.webp)

下载压缩包进行加压后，得到 .pem 证书和 .key 证书，把两个证书上传到 Linux，进行配置，往下看。

**方式二：使用 OpenSSL 生成证书**

这个方式适用于学习阶段，实际开发阶段使用方式一

先要确认当前系统是否有安装 OpenSSL

```sh
openssl version
```

安装下面的命令进行生成（一步一步来）

创建 `/root/cerr` 目录并进入

```sh
mkdir /root/cert

cd /root/cert
```

指定加密算法和加密方式，生成 `server.key`

```sh
# genrsa 是加密算法，des3 是加密方式，-out 代表输入长度为 2048 的 server.key
openssl genrsa -des3 -out server.key 2048

# 然后它会让你注册 server.key 的密码
# 输入密码 ......
```

根据你注册的 server.key 密码，生成 server.csr 文件，生成后它会让你注册你的基本信息，因为是个人的，所以信息随便填写

```sh
openssl req -new -key server.key -out server.csr

# 这里注册你的基本信息，信息随便填写
```

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.2b4bgtawqb4.webp)

备份 server.key

```sh
cp server.key server.key.org
```

重新生成 server.key 文件，并输入刚才注册的密码

```sh
openssl rsa -in server.key.org -out server.key

# 会让你重新输入注册密码
```

生成 server.crt 文件

```sh
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```

最后使用 `ll` 查看目录下是否生成相应的文件：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.3yl8ne2bkre0.webp)

### SSL实例模板(通用)

Nginx 的置文件添加如下内容：

```nginx
server {
    listen       80;
    # ......
}
server {
    listen       443 ssl;		# 开启 SSL 功能
    server_name  localhost;     # 如果是购买的域名，这里加上该域名

    ssl_certificate      /root/cert/server.cert; # 生成的 cert 或者 pem 证书路径，根据需求修改
    ssl_certificate_key  /root/cert/server.key; # 生成的 key 证书路径，根据需求修改
    ssl_session_timeout 5m; 
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4; # 表示使用的加密套件的类型
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;  # 表示使用的TLS协议的类型
    ssl_prefer_server_ciphers on;

    location / {
        root   html;
        index  index.html index.htm;
    }
}
```

其实这个模板就是 Nginx 解压目录的 conf 目录下的 nginx.conf 内容。

建议

配置 SSL 证书时候，不要修改原来的 server 模块（`listen 80` 模块），应该新建一个 server 模块。

**解决默认 http 问题。**

配置完 SSL 证书还远远不安全。如果直接输入 `www.frx.com`，会默认加上『 http:// 』 而不是『 https:// 』，如何修改为『 https:// 』呢，我们利用 Rewrite 功能

```nginx
server {
    listen 443 ssl;
    server_name www.frx.com;   # 如果是 www.frx.com 发送请求
    
    location / {
        # ......
        rewrite ^(.*)$ https://www.frx.com$1;  # 则改为 https 方式
        # ......
	}
    # ......
}
```

## 反向代理系统调优

反向代理值 Buffer 和 Cache。

Buffer 翻译过来是「缓冲」，Cache 翻译过来是「缓存」。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220802/image.6p7uoj5ab0k0.webp)

客户端通过代理服务器向被代理服务器获取数据后，代理服务器在获取的数据存储在缓存「瓶子」里，客户端再次获取相同资源时，直接从缓存「瓶子」里获取数据，不需要向被代理服务器获取数据，减轻压力。

相同点:

- 两种方式都是用来提供 IO 吞吐效率，都是用来提升 Nginx 代理的性能。

不同点:

- 缓冲主要用来解决不同设备之间数据传递速度不一致导致的性能低的问题，缓冲中的数据一旦此次操作完成后，就可以删除
- 缓存主要是备份，将被代理服务器的数据缓存一份到代理服务器，这样的话，客户端再次获取相同数据的时候，就只需要从代理服务器上获取，效率较高，缓存中的数据可以重复使用，只有满足特定条件才会删除

### Proxy Buffer 相关指令

- `proxy_buffering` 指令用来开启或者关闭代理服务器的缓冲区，默认开启。

  | 语法                         | 默认值              | 位置                   |
  | ---------------------------- | ------------------- | ---------------------- |
  | proxy_buffering <on \| off>; | proxy_buffering on; | http、server、location |

- `proxy_buffers` 指令用来指定单个连接从代理服务器读取响应的缓存区的个数和大小。

  | 语法                           | 默认值                                    | 位置                   |
  | ------------------------------ | ----------------------------------------- | ---------------------- |
  | proxy_buffers <number> <size>; | proxy_buffers 8 4k \| 8K;(与系统平台有关) | http、server、location |

  - number：缓冲区的个数
  - size：每个缓冲区的大小，缓冲区的总大小就是 number * size

- `proxy_buffer_size` 指令用来设置从被代理服务器获取的第一部分响应数据的大小。保持与 proxy_buffers 中的 size 一致即可，当然也可以更小。

  | 语法                      | 默认值                                      | 位置                   |
  | ------------------------- | ------------------------------------------- | ---------------------- |
  | proxy_buffer_size <size>; | proxy_buffer_size 4k \| 8k;(与系统平台有关) | http、server、location |

- `proxy_busy_buffers_size` 指令用来限制同时处于 BUSY 状态的缓冲总大小。

  | 语法                            | 默认值                             | 位置                   |
  | ------------------------------- | ---------------------------------- | ---------------------- |
  | proxy_busy_buffers_size <size>; | proxy_busy_buffers_size 8k \| 16K; | http、server、location |

- `proxy_temp_path` 指令用于当缓冲区存满后，仍未被 Nginx 服务器完全接受，响应数据就会被临时存放在磁盘文件上的该指令设置的文件路径下

  | 语法                    | 默认值                      | 位置                   |
  | ----------------------- | --------------------------- | ---------------------- |
  | proxy_temp_path <path>; | proxy_temp_path proxy_temp; | http、server、location |

  注意 path 最多设置三层。

- `proxy_temp_file_write_size` 指令用来设置磁盘上缓冲文件的大小。

  | 语法                               | 默认值                                | 位置                   |
  | ---------------------------------- | ------------------------------------- | ---------------------- |
  | proxy_temp_file_write_size <size>; | proxy_temp_file_write_size 8K \| 16K; | http、server、location |

### 网站调优模板(通用)

```nginx
proxy_buffering on;
proxy_buffers 4 64k;
proxy_buffer_size 64k;
proxy_busy_buffers_size 128k;
proxy_temp_file_write_size 128k;
```

### 反向代理问题

反向代理有一个严重的问题，那就是反向代理的网站需要的 css、js、png 等静态文件默认是从 Nginx 里获取，显然一个你的服务器的代理了如淘宝、京东等网站，那么淘宝、京东等网站的静态文件是从你的服务器里获取，但是服务器根本没有，那么我们如何让代理的网站以及网站需要的 css、js、png 等静态文件一同获取呢，添加如下配置：



```nginx
server {
    listen  80;
    server_name localhost;
    
    location ~ .* {							# 如果不是代理网站的根路径，请自行修改
        proxy_pass http://127.0.0.1:8081;   # 代理的网站地址
        # 将网站的静态文件也代理过来
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	}
}
```



## 负载均衡

早期的网站流量和业务功能都比较简单，单台服务器足以满足基本的需求，但是随着互联网的发展，业务流量越来越大并且业务逻辑也跟着越来越复杂，单台服务器的性能及单点故障问题就凸显出来了，因此需要多台服务器进行性能的水平扩展及避免单点故障出现。那么如何将不同用户的请求流量分发到不同的服务器上呢？这就需要负载均衡来处理。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.5kgql84ikx80.webp)

### 负载均衡原理及处理流程

系统的扩展可以分为纵向扩展和横向扩展。

- 纵向扩展是从单机的角度出发，通过增加系统的硬件处理能力来提升服务器的处理能力
- 横向扩展是通过添加机器来满足大型网站服务的处理能力

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.1xwd1pm6du68.webp)

如上图，负载均衡涉及到两个重要的角色分别是「应用集群」和「负载均衡器」。

- 应用集群：将同一应用部署到多台机器上，组成处理集群，接收负载均衡设备分发的请求，进行处理并返回响应的数据
- 负载均衡器：将用户访问的请求根据对应的负载均衡算法，分发到集群中的一台服务器进行处理

### 负载均衡作用

- 解决服务器的高并发压力，提高应用程序的处理性能
- 提供故障转移，实现高可用
- 通过添加或减少服务器数量，增强网站的可扩展性
- 在负载均衡器上进行过滤，可以提高系统的安全性

### 负载均衡常用处理方式

先说明，我们常用的是 [四/七层负载均衡](https://frxcat.fun/middleware/Nginx/Nginx_load_balancing/#四-七层负载均衡) 方式，前面两个方式可以了解。

#### 用户手动选择

这种方式比较原始，主要实现的方式就是在网站主页上面提供不同线路、不同服务器链接方式，让用户来选择自己访问的具体服务器，来实现负载均衡。

如下图，用户点击不同的下载方式，就会跳转到不同的下载地址。这是主动式的负载均衡，我们无法控制用户的选择。如果用户全部点击第一个下载方式，那么服务器的压力将非常大。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.2auhoddxb2zo.webp)

#### DNS轮询方式

**DNS：域名系统（服务）协议（DNS）是一种分布式网络目录服务，主要用于域名与 IP 地址的相互转换。**

大多域名注册商都支持对同一个主机名添加多条 A 记录，这就是 DNS 轮询，DNS 服务器将解析请求按照 A 记录的顺序，随机分配到不同的 IP 上，这样就能完成简单的负载均衡。DNS 轮询的成本非常低，在一些不重要的服务器，被经常使用。

如下图：客户端如果想访问服务器集群，首先去 DNS 服务器获取我们曾经在 DNS 服务器保存的「记录表」，这个「记录表」将会把某个服务器的地址返回给客户端，客户端再根据这个地址，访问指定的服务器。这个「记录表」在开始期间需要我们去 DNS 服务器进行添加。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.b1m3qaeuxsk.webp)

「记录表」长什么样，如下图的主机记录 www。这是我为某一个域名添加的 IP 地址，用 2 台服务器来做负载均衡。其中两个记录值都绑定了 `www.nginx521.cn` 地址。(一个域名可以绑定多个IP地址)

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.3drj9szeeus0.webp)

验证:

```sh
ping www.nginx521.cn
```

注意：记得清空本地的 DNS 缓存，如果本地有缓存，无论你怎么 `ping`，都会 `ping` 到缓存的服务器地址，无法负载均衡

```sh
ipconfig/flushdns
```

目前需要 `ping` 一次然后清理一次缓存，才能实现负载均衡的轮询效果。

我们发现使用 DNS 来实现轮询，不需要投入过多的成本，虽然 DNS 轮询成本低廉，但是 DNS 负载均衡存在明显的缺点：

1. 可靠性低

   假设一个域名 DNS 轮询多台服务器，如果其中的一台服务器发生故障，那么所有的访问该服务器的请求将不会有所回应，即使你将该服务器的 IP 从 DNS 中去掉，但是由于各大宽带接入商将众多的 DNS 存放在缓存中，以节省访问时间，导致 DNS 不会实时更新。所以 DNS 轮流上一定程度上解决了负载均衡问题，但是却存在可靠性不高的缺点。

2. 负载均衡不均衡

   DNS 负载均衡采用的是简单的轮询负载算法，不能区分服务器的差异，不能反映服务器的当前运行状态，不能做到为性能好的服务器多分配请求，另外本地计算机也会缓存已经解析的域名到 IP 地址的映射，这也会导致使用该 DNS 服务器的用户在一定时间内访问的是同一台 Web 服务器，从而引发 Web 服务器减的负载不均衡。

   负载不均衡则会导致某几台服务器负荷很低，而另外几台服务器负荷确很高，处理请求的速度慢，配置高的服务器分配到的请求少，而配置低的服务器分配到的请求多。

### 四/七层负载均衡

介绍四/七层负载均衡之前，我们先了解一个概念，OSI(open system interconnection)，叫开放式系统互联模型，这个是由国际标准化组织 ISO 指定的一个不基于具体机型、操作系统或公司的网络体系结构。该模型将网络通信的工作分为七层。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.5g3vvj5xa3g0.webp)

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.6uq265ihol80.webp)

- 应用层：为应用程序提供网络服务。
- 表示层：对数据进行格式化、编码、加密、压缩等操作
- 会话层：建立、维护、管理会话连接
- 传输层：建立、维护、管理端到端的连接，常见的有 TCP/UDP
- 网络层：IP 寻址和路由选择
- 数据链路层：控制网络层与物理层之间的通信
- 物理层：比特流传输

**什么是四层负载均衡**

所谓四层负载均衡指的是 OSI 七层模型中的传输层，主要是基于 IP + PORT 的负载均衡

实现四层负载均衡的方式：

- 硬件：F5 BIG-IP、Radware 等，性能好，成本高、无法扩展
- 软件：LVS、Nginx、Hayproxy 等，性能较好，成本低、可以扩展

**什么是七层负载均衡**

所谓的七层负载均衡指的是在应用层，主要是基于虚拟的 URL 或主机 IP 的负载均衡

实现七层负载均衡的方式：

- 软件：Nginx、Hayproxy 等

**四层和七层负载均衡的区别**

- 四层负载均衡数据包是在底层就进行了分发，而七层负载均衡数据包则在最顶端进行分发，所以四层负载均衡的效率比七层负载均衡的要高（四层比七层少，速度块，效率高，但是可能请求丢失等）
- 四层负载均衡不识别域名，而七层负载均衡识别域名

处理四层和七层负载以外，其实还有二层、三层负载均衡，二层是在数据链路层基于 mac 地址来实现负载均衡，三层是在网络层一般采用虚拟 IP 地址的方式实现负载均衡。

**实际环境采用的模式：四层负载(LVS) + 七层负载(Nginx)**

### 七层负载均衡

Nginx 要实现七层负载均衡需要用到 proxy_pass 代理模块配置。Nginx 默认安装支持这个模块，我们不需要再做任何处理。

**Nginx 的负载均衡是在 Nginx 反向代理的基础上把用户的请求根据指定的算法分发到一组「upstream 虚拟服务池」**。

#### 七层负载均衡指令

##### upstream指令

该指令是用来定义一组服务器，它们可以是监听不同端口的服务器，并且也可以是同时监听 TCP 和 Unix socket 的服务器。

服务器可以指定不同的权重，默认为 1。

| 语法                  | 默认值 | 位置 |
| --------------------- | ------ | ---- |
| upstream <name> {...} | —      | http |

##### server指令

该指令用来指定后端服务器的名称和一些参数，可以使用域名、IP、端口或者 Unix socket。

| 语法                        | 默认值 | 位置     |
| --------------------------- | ------ | -------- |
| server <name> [paramerters] | —      | upstream |

server 后的 name 就是 upstream 后的 name，两者保持一致。

> #### 七层负载均衡指令案例
>
> 准备四台服务器，一台用来做负载均衡器，三台用来接收负载均衡器的请求。
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220803/image.etsfgkwsxaw.webp)
>
> 因为目前只有两台服务器，所以一台用来做负载均衡器，另外一台用来接收负载均衡器的请求。
>
> 服务器设置：这里以三个端口代替三个服务器，在配置文件进行如下配置：
>
> 
>
> ```nginx
> # 服务器 1
> server {
>     listen   9001;
>     server_name localhost;
>     default_type text/html;
>     location /{
>     	return 200 '<h1>192.168.200.146:9001</h1>';
>     }
> }
> # 服务器 2
> server {
>     listen   9002;
>     server_name localhost;
>     default_type text/html;
>     location /{
>     	return 200 '<h1>192.168.200.146:9002</h1>';
>     }
> }
> # 服务器 3
> server {
>     listen   9003;
>     server_name localhost;
>     default_type text/html;
>     location / {
>     	return 200 '<h1>192.168.200.146:9003</h1>';
>     }
> }
> ```
>
> 负载均衡器设置：这是一个 Nginx 代理服务器，让它去负载均衡访问三个服务器，在配置文件进行如下配置：
>
> ```nginx
> upstream backend{
> 	server 192.168.200.146:9091;
> 	server 192.168.200.146:9092;
> 	server 192.168.200.146:9093;
> }
> server {
> 	listen 8083;
> 	server_name localhost;
> 	location / {
> 		proxy_pass http://backend;   # backend 要对应上 upstream 后的值，根据需求修改
> 	}
> }
> ```
>
> 访问负载均衡器的地址，如 `http://192.168.200.133:8083`，它会找到 `proxy_pass` 后的地址，比如上方，它会根据 backend 找到对应的 upstream 里内地址，替换掉 backend，变成：
>
> - proxy_pass `http://192.168.200.146:9091`
> - proxy_pass `http://192.168.200.146:9092`
> - proxy_pass `http://192.168.200.146:9093`
>
> 但是它不会全部访问三个服务器地址，而是根据自己的算法（轮询）选择其中一个服务器地址。

#### 七层负载均衡状态标识

代理服务器在负责均衡调度中的状态有以下几个：

| 状态         | 概述                                |
| ------------ | ----------------------------------- |
| down         | 当前的 server 暂时不参与负载均衡    |
| backup       | 预留的备份服务器                    |
| max_fails    | 允许请求失败的次数                  |
| fail_timeout | 经过 max_fails 失败后，服务暂停时间 |
| max_conns    | 限制最大的接收连接数                |

##### down

`down` 指令将该服务器标记为永久不可用，那么负载均衡器将不参与该服务器的负载均衡。

如下，如果不希望负载均衡器以负载均衡来处理 `192.168.200.146` 服务器：

```nginx
upstream backend{
	server 192.168.200.146:9001 down;
	server 192.168.200.146:9002
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location / {
		proxy_pass http://backend;
	}
}
```

该状态一般会对需要停机维护的服务器进行设置。

##### backup

`backup` 指令将该服务器标记为备份服务器，当主服务器不可用时，才用备份服务器来传递请求。

它不同于 down 指令，down 指令将服务器永久禁止，而 backp 指令仅仅临时禁止，当主服务器不可用后，临时禁止的服务器就会站出来。

```nginx
upstream backend{
	server 192.168.200.146:9001 down;
	server 192.168.200.146:9002 backup;
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

上方代码中 9001 服务器永久禁止，而 9002 服务器是备份服务器，所以 9003 服务器是主服务器。

此时需要将 9003 端口的访问禁止掉，用它来模拟当唯一对外提供访问的服务宕机以后，backup 的备份服务器就能开始对外提供服务。

> 为了测试验证，我们需要使用防火墙来进行拦截。
>
> 介绍一个工具 `firewall-cmd`，该工具是 Linux 提供的专门用来操作 firewall 防火墙的。
>
> 查询防火墙中指定的端口是否开放
>
> ```sh
> firewall-cmd --query-port=9001/tcp
> ```
>
> 开放一个指定的端口
>
> ```sh
> firewall-cmd --permanent --add-port=9002/tcp
> ```
>
> 批量添加开发端口
>
> ```sh
> firewall-cmd --permanent --add-port=9001-9003/tcp
> ```
>
> 如何移除一个指定的端口
>
> ```sh
> firewall-cmd --permanent --remove-port=9003/tcp
> ```
>
> 重新加载
>
> ```sh
> firewall-cmd --reload
> ```
>
> 其中
>
> - `--permanent` 表示设置为持久
> - `--add-port` 表示添加指定端口
> - `--remove-port` 表示移除指定端口
>
> 经过测试，禁用掉 9003 端口后，再次访问负载均衡器，它只会请求 9002 端口的服务器(备份服务器)，而恢复 9003 端口，只会请求 9003 端口的服务器。

##### max_conns

`max_conns` 指令用来限制同时连接到 upstream 负载上的单个服务器的最大连接数。默认为 0，表示不限制，使用该配置可以根据后端服务器处理请求的并发量来进行设置，防止后端服务器被压垮。

| 语法               | 默认值 | 位置     |
| ------------------ | ------ | -------- |
| max_conns=<number> | 0      | upstream |

- number 是大于 0 的数字。

```nginx
upstream backend{
	server 192.168.200.146:9001 down;
	server 192.168.200.146:9002 backup;
	server 192.168.200.146:9003 max_conns=2;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

第 4 行配置标识 9003 端口的服务器最大能被 2 个客户端请求。

##### max_fails和fail_timeout

`max_fails` 指令设置允许请求代理服务器失败的次数，默认为 1。

`fail_timeout` 指令设置经过 max_fails 失败后，服务暂停的时间，默认是 10 秒。

| 语法                | 默认值 | 位置     |
| ------------------- | ------ | -------- |
| max_fails=<number>  | 1      | upstream |
| fail_timeout=<time> | 10 秒  | upstream |

- number 是大于 0 的数字
- time 是时间，单位为秒

```nginx
upstream backend{
	server 192.168.200.133:9001 down;
	server 192.168.200.133:9002 backup;
	server 192.168.200.133:9003 max_fails=3 fail_timeout=15;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

#### 七层负载均衡策略

介绍完 Nginx 负载均衡的相关指令后，我们已经能实现将用户的请求分发到不同的服务器上，那么除了采用默认的分配方式以外，我们还能采用什么样的负载算法？

Nginx 的 upstream 支持如下六种方式的分配算法，分别是:

| 算法名称   | 说明              |
| ---------- | ----------------- |
| 轮询       | 默认方式          |
| weight     | 权重方式          |
| ip_hash    | 依据 IP 分配方式  |
| least_conn | 依据最少连接方式  |
| url_hash   | 依据 URL 分配方式 |
| fair       | 依据响应时间方式  |

##### 轮询

这是 `upstream` 模块负载均衡默认的策略。每个请求会按时间顺序逐个分配到不同的后端服务器。轮询不需要额外的配置。

```nginx
upstream backend{
	server 192.168.200.146:9001;
	server 192.168.200.146:9002;
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

##### weight加权[加权轮询]

`weight` 指令用来设置服务器的权重，默认为 1，权重数据越大，被分配到请求的几率越大；该权重值，主要是针对实际工作环境中不同的后端服务器硬件配置进行调整的，所有此策略比较适合服务器的硬件配置差别比较大的情况。

| 语法            | 默认值 | 位置     |
| --------------- | ------ | -------- |
| weight=<number> | 1      | upstream |

- number 是大于 0 的数字

```nginx
upstream backend{
	server 192.168.200.146:9001 weight=10;
	server 192.168.200.146:9002 weight=5;
	server 192.168.200.146:9003 weight=3;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```



##### ip_hash

当对后端的多台动态应用服务器做负载均衡时，`ip_hash` 指令能够将某个客户端 IP 的请求通过哈希算法定位到同一台后端服务器上。

这样，当来自某一个 IP 的用户在后端 Web 服务器 A 上登录后，在访问该站点的其他 URL，能保证其访问的还是后端 Web 服务器 A

这里是为了解决因为负载均衡导致同一个用户在访问时 后端服务变化导致可能出现重复登录的情况出现，因为当用户IP相同时，后端提供服务的服务器是固定的。

总结：哪个服务器曾经处理过请求，无论在哪里，相同的请求依然让该服务器处理

| 语法     | 默认值 | 位置     |
| -------- | ------ | -------- |
| ip_hash; | —      | upstream |

```nginx
upstream backend{
	ip_hash;
	server 192.168.200.146:9001;
	server 192.168.200.146:9002;
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

需要额外多说一点的是使用 ip_hash 指令无法保证后端服务器的负载均衡，可能导致有些后端服务器接收到的请求多，有些后端服务器接收的请求少，而且设置后端服务器权重等方法将不起作用。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.7a0ysgphfxg0.webp)

##### least_conn

最少连接数，把请求转发给连接数较少的后端服务器。

轮询算法是把请求平均的转发给各个后端，使它们的负载大致相同；但是，有些请求占用的时间很长，会导致其所在的后端负载较高。这种情况下，`least_conn` 这种方式就可以达到更好的负载均衡效果。

```nginx
upstream backend{
	least_conn;
	server 192.168.200.146:9001;
	server 192.168.200.146:9002;
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

此负载均衡策略适合请求处理时间长短不一造成服务器过载的情况。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.5vug3q9whrc0.webp)

##### url_hash

按访问 URL 的 hash 结果来分配请求，使每个 URL 定向到同一个后端服务器，要配合缓存命中来使用。

当出现同一个资源多次请求，可能会到达不同的服务器上，导致不必要的多次下载，缓存命中率不高，以及一些资源时间的浪费时，使用 `url_hash`，可以使得同一个 URL（也就是同一个资源请求）会到达同一台服务器，一旦缓存住了资源，再此收到请求，就可以从缓存中读取。

总结：发送相同的请求时，希望只有一个服务器处理该请求，使用 `uri_hash`。因为 URL 相同，则哈希值(hash)相同，那么哈希值对应的服务器也相同。

```nginx
upstream backend{
	hash &request_uri;  # 通过 &request_uri 获取访问路径 再计算hash做负载均衡
	server 192.168.200.146:9001;
	server 192.168.200.146:9002;
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

如图：文件系统有一个文件，目前只有 web 服务 1 和服务 3 获取了该文件，那么我们想要下载这个文件时，只能找服务 1 或服务 3，这时候就固定一个 URL，该 URL 不允许服务 2 进行处理，那么如何规定哪个服务处理呢？就用到 `url_hash`。

它会根据 URL 计算处哈希值，由哈希值对应服务，所以固定下载文件的 URL，就固定了一个服务处理。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.utb83x535xc.webp)

##### fair

`fair` 指令采用的不是内建负载均衡使用的轮换的均衡算法，而是可以根据页面大小、加载时间长短智能的进行负载均衡。

```nginx
nupstream backend{
	fair;
	server 192.168.200.146:9001;
	server 192.168.200.146:9002;
	server 192.168.200.146:9003;
}
server {
	listen 8083;
	server_name localhost;
	location /{
		proxy_pass http://backend;
	}
}
```

> 
>
>  fair 属于第三方模块实现的负载均衡。需要添加 `nginx-upstream-fair` 模块，如何添加对应的模块：
>
> 1、下载 `nginx-upstream-fair` 模块，下载地址如下：
>
> ```sh
> https://github.com/gnosek/nginx-upstream-fair
> ```
>
> 2、将下载的文件上传到服务器并进行解压缩
>
> ```sh
> # 进入安装包目录
> cd /opt
> 
> # 解压
> unzip nginx-upstream-fair-master.zip
> ```
>
> 3、我的解压目录在 `/opt`，所以第 6 步记得指定好模块的路径。
>
> ```sh
> mv nginx-upstream-fair-master fair
> ```
>
> 4、将原有 `/usr/local/nginx/sbin/nginx` 进行备份
>
> ```sh
> mv /usr/local/nginx/sbin/nginx /usr/local/nginx/sbin/nginx.backup
> ```
>
> 5、查看 `configure arguments` 的配置信息，拷贝出来
>
> ```sh
> nginx -V
> 
> # 拷贝 configure arguments 后面的数据
> ```
>
> 6、进入 Nginx 的安装目录，执行 make clean 清空之前编译的内容
>
> ```sh
> cd /root/nginx/core/nginx-1.21.6
> 
> make clean
> ```
>
> 7、使用 configure 来配置参数，添加模块，记得加上第（4）步拷贝的配置信息
>
> ```sh
> ./configure --add-module=/opt/fair  # 记得添加 configure arguments 后的数据
> ```
>
> 8、通过 make 模板进行编译
>
> ```sh
> make
> ```
>
> 编译可能会出现如下错误，`ngx_http_upstream_srv_conf_t` 结构中缺少 `default_port`
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.vj3xrayy84g.webp)
>
> 解决方案：
>
> 在 Nginx 的源码目录（安装包目录）中 `src/http/ngx_http_upstream.h`，找到 `ngx_http_upstream_srv_conf_s`，在模块中添加添加 `default_port` 属性
>
> ```sh
> vim /opt/nginx/core/nginx-1.21.6/src/http/ngx_http_upstream.h
> ```
>
> 添加内容：
>
> ```sh
> in_port_t	   default_port
> ```
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.1mcquncem2u8.webp)
>
> 然后再进行 make。
>
> 9、将安装目录下的 objs 中的 nginx 拷贝到 sbin 目录
>
> ```sh
> cd /opt/nginx/core/nginx-1.21.6/objs
> cp nginx /usr/local/nginx/sbin
> ```
>
> 10、更新nginx
>
> ```sh
> cd /opt/nginx/core/nginx-1.21.6
> make upgrade
> ```
>



#### 七层负载均衡案例

> #### 案例一：对所有请求实现一般轮询规则的负载均衡
>
> ```nginx
> upstream backend{
> 	server 192.168.200.146:9001;
> 	server 192.168.200.146:9002;
> 	server 192.168.200.146:9003;
> }
> server {
> 	listen 8083;
> 	server_name localhost;
> 	location / {
> 		proxy_pass http://backend;
> 	}
> }
> ```
>
> #### 案例二：对所有请求实现加权轮询规则的负载均衡
>
> ```nginx
> upstream backend{
> 	server 192.168.200.146:9001 weight=7;
> 	server 192.168.200.146:9002 weight=3;
> 	server 192.168.200.146:9003 weight=5;
> }
> server {
> 	listen 8083;
> 	server_name localhost;
> 	location /{
> 		proxy_pass http://backend;
> 	}
> }
> ```
>
> 处理请求概率：9001 端口 > 9003 端口 > 9002 端口
>
> #### 案例三：对特定资源实现负载均衡
>
> ```nginx
> upstream videobackend{
> 	server 192.168.200.146:9001;
> 	server 192.168.200.146:9002;
> }
> upstream filebackend{
> 	server 192.168.200.146:9003;
> 	server 192.168.200.146:9004;
> }
> server {
> 	listen 8084;
> 	server_name localhost;
> 	location /video/ {
> 		proxy_pass http://videobackend;
> 	}
> 	location /file/ {
> 		proxy_pass http://filebackend;
> 	}
> }
> ```
>
> 发送 `/video/` 请求会被 9001 和 9002 端口的服务器处理。
>
> 发送 `/file/` 请求会被 9003 和 9004 端口的服务器处理。
>
> #### 案例四：对不同域名实现负载均衡
>
> ```nginx
> upstream frxbackend{
> 	server 192.168.200.146:9001;
> 	server 192.168.200.146:9002;
> }
> upstream bingbackend{
> 	server 192.168.200.146:9003;
> 	server 192.168.200.146:9004;
> }
> server {
> 	listen	8085;
> 	server_name www.frx.com;
> 	location / {
> 		proxy_pass http://frxbackend;
> 	}
> }
> server {
> 	listen	8086;
> 	server_name www.bing.com;
> 	location / {
> 		proxy_pass http://bingbackend;
> 	}
> }
> ```
>
> `www.frx.com` 地址的请求由 9001 端口和 9002 端口处理。
>
> `www.bing.com` 地址的请求由 9003 端口和 9004 端口处理。
>
> #### 案例五：实现带有URL重写的负载均衡
>
> ```nginx
> upstream backend{
> 	server 192.168.200.146:9001;
> 	server 192.168.200.146:9002;
> 	server 192.168.200.146:9003;
> }
> server {
> 	listen	80;
> 	server_name localhost;
> 	location /file/ {
> 		rewrite ^(/file/.*) /server/$1 last;
> 	}
> 	location /server {
> 		proxy_pass http://backend;
> 	}
> }
> ```
>
> 将 `/file/xxx` 请求重写为 `/server/xxx`，然后触发 `location /server`，实现负载均衡。
>
> 此时被负载均衡的服务器地址也会带有 `/server` 以及后面的参数，如 `192.168.200.146:9001/server/xxx

### 四层负载均衡

Nginx 在 1.9 之后，增加了一个 stream 模块，用来实现四层协议的转发、代理、负载均衡等。stream 模块的用法跟 http 的用法类似，允许我们配置一组 TCP 或者 UDP 等协议的监听，然后通过 proxy_pass 来转发我们的请求，通过 upstream 添加多个后端服务，实现负载均衡。

四层协议负载均衡的实现，一般都会用到 LVS、HAProxy、F5 等，要么很贵要么配置很麻烦，而 Nginx 的配置相对来说更简单，更能快速完成工作。

#### 添加stream模块的支持

Nginx 默认是没有编译这个模块的，需要使用到 stream 模块，那么需要在编译的时候加上 `--with-stream`。

完成添加 `--with-stream` 的实现步骤：

- 将原有 `/usr/local/nginx/sbin/nginx` 进行备份
- 拷贝 `Nginx -V` 的 configure arguments 配置信息
- 在 Nginx 的安装源码进行配置指定对应模块：`./configure --with-stream 加上一步拷贝的configure arguments 配置`
- 通过 make 模板进行编译
- 将 objs 下面的 nginx 移动到 `/usr/local/nginx/sbin` 下
- 在源码目录下执行 `make upgrade` 进行升级，这个可以实现不停机添加新模块的功能

添加模块的详细步骤我已经在 [七层负载均衡策略-fail 指令](https://frxcat.fun/middleware/Nginx/Nginx_load_balancing/)、[静态资源部署-Nginx 模块添加](https://frxcat.fun/middleware/Nginx/Nginx_Static_resource_deployment/#nginx模块添加)、[反向代理-添加ssl支持](https://frxcat.fun/middleware/Nginx/Nginx_Reverse_proxy/#添加ssl支持) 描述过，而你只需要替换模块名字罢了。

#### 四层负载均衡指令

如果不想在 http 模块使用负载均衡，可以在 steam 模块使用。

##### stream指令

该指令提供在其中指定流服务器指令的配置文件上下文。**和 http 模块同级**。

| 语法           | 默认值 | 位置 |
| -------------- | ------ | ---- |
| stream { ... } | —      | main |



```nginx
http {
    server {
        listen 80;
        # ......
    }
}
stream {
    upstream backend{
        server 192.168.200.146:6379;
        server 192.168.200.146:6378;
    }
    server {
        listen 81;
        proxy_pass backend;
    }
}
```

##### upstream指令

该指令和七层负载均衡的 upstream 指令是类似的。

#### 四层负载均衡的案例

准备两台服务器，这里称为 A 和 B。服务器 A 的 IP 为 `192.168.200.146`，服务器 B 的IP 为 `192.168.200.133`，服务器 A 存放 Redis 和 Tomcat，服务器 B 作为负载均衡器，对服务器 A 的端口进行负载均衡处理。

##### 需求分析

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.6yak0csdt880.webp)

##### Redis 配置

准备 Redis 服务器,在服务器 A 上准备两个 Redis，端口分别是 6379、6378。

1、上传 redis 的安装包 `redis-4.0.14.tar.gz`，这里上传目录是 `/opt`

2、将安装包进行解压缩

```sh
tar -zxf redis-4.0.14.tar.gz
```

3、进入redis的安装包

```sh
cd redis-4.0.14
```

4、使用 make 和 install 进行编译和安装，这里的安装路径是 `/usr/local/redis/redis01`

```sh
make PREFIX=/usr/local/redis/redis01 install
```

5、拷贝 redis 配置文件 `redis.conf` 到 `/usr/local/redis/redis01/bin` 目录中，因为安装后，目录并没有 redis.conf

```sh
cp /opt/redis-4.0.14/redis.conf	/usr/local/redis/redis01/bin
```

6、修改 redis.conf 配置文件，注意：不是添加内容，是修改内容，要自己搜索 bind、port 和 daemonize 进行修改

```sh
bind 0.0.0.0   # 允许任意地址访问
port  6379      # redis 的端口
daemonize yes   # 后台启动 redis
```

7、将 redis01 复制一份为 redis02

```sh
cd /usr/local/redis
cp -r redis01 redis02
```

8、将 redis02 文件夹中的 redis.conf 进行修改，注意：不是添加内容，是修改内容，要自己搜索 bind、port 和 daemonize 进行修改

```sh
bind 0.0.0.0   # 允许任意地址访问
port  6378      # redis 的端口
daemonize yes   # 后台启动 redis
```

9、分别启动，即可获取两个 Redis 并查看

```sh
redis-server <配置文件路径>

ps -ef | grep redis
```

使用 Nginx 将请求分发到不同的 Redis 服务器上。

安装 Redis 并验证能启动成功后，在另一台服务器 B `192.168.200.133` 的 Nginx 配置文件添加如下内容：（确保安装了 steam 模块）

```nginx
stream {
    upstream redisbackend{
        server 192.168.200.146:6379;   # 服务器 B 的 6379 端口
        server 192.168.200.146:6378;   # 服务器 B 的 6378 端口
    }
    server {
        listen 81;
        proxy_pass redisbackend;
    }
}
```

此时利用 redis-cli 连接测试

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220804/image.cbee2u83yhs.webp)

服务器 B 通过负载均衡连接到了服务器 A 的 Redis，只是不知道连接的是 6378 还是 6379 端口，可以在 Redis 添加不一样的数据来测试，这里不演示了。

##### Tomcat 配置

准备 Tomcat 服务器 到服务器 A

1、上传 tomcat 的安装包，`apache-tomcat-8.5.56.tar.gz`

2、将安装包进行解压缩

```sh
tar -zxf apache-tomcat-8.5.56.tar.gz
```

3、进入 tomcat 的 bin 目录，启动 tomcat

```sh
cd apache-tomcat-8.5.56/bin
./startup
```

服务器 B 的配置文件 nginx.conf 配置如下：

```nginx
stream {
    upstream redisbackend {
        server 192.168.200.146:6379;    # 服务器 B 的 6379 端口
        server 192.168.200.146:6378;    # 服务器 B 的 6378 端口
    }
    upstream tomcatbackend {
        server 192.168.200.146:8080;   # 服务器 B 的 8080 端口
    }
    server {
        listen  81;
        proxy_pass redisbackend; # redis 的负载均衡
    }
    server {
        listen	82;
        proxy_pass tomcatbackend;  # tomcat 的负载均衡
    }
}
```

访问服务器 B 的地址进行测试：`192.168.200.133:82`。

> ## 核心概念区别
>
> | 特性         | 四层负载均衡 (Layer 4)    | 七层负载均衡 (Layer 7)     |
> | :----------- | :------------------------ | :------------------------- |
> | **工作层级** | 传输层 (TCP/UDP)          | 应用层 (HTTP/HTTPS)        |
> | **决策依据** | IP地址 + 端口号           | URL路径、Header、Cookie等  |
> | **可见内容** | 只能看到原始TCP/UDP数据包 | 可以看到完整的HTTP请求内容 |
> | **性能**     | 更高（处理简单）          | 相对较低（需要解析内容）   |
> | **配置模块** | `stream` 模块             | `http` 模块                |
>
> ## 配置结构差异
>
> ### 四层负载均衡配置
>
> ```nginx
> # 在 nginx.conf 的 main 层级配置
> stream {
>     upstream backend_servers {
>         server 192.168.1.10:3306;  # MySQL
>         server 192.168.1.11:3306;
>         server 192.168.1.12:3306;
>     }
>     
>     server {
>         listen 13306;  # 代理端口
>         proxy_pass backend_servers;
>         proxy_timeout 3s;
>         proxy_connect_timeout 2s;
>     }
> }
> ```
>
> 
>
> ### 七层负载均衡配置
>
> 
>
> ```nginx
> # 在 http 块内配置
> http {
>     upstream web_servers {
>         server 192.168.1.20:8080 weight=3;
>         server 192.168.1.21:8080 weight=2;
>         server 192.168.1.22:8080 weight=1;
>         
>         # 负载均衡算法
>         least_conn;  # 最少连接数
>     }
>     
>     server {
>         listen 80;
>         server_name example.com;
>         
>         location / {
>             proxy_pass http://web_servers;
>             proxy_set_header Host $host;
>             proxy_set_header X-Real-IP $remote_addr;
>         }
>     }
> }
> ```
>
> 
>
> ## 具体应用场景举例
>
> ### 四层负载均衡用例
>
> #### 例1：数据库负载均衡
>
> ```nginx
> stream {
>     # MySQL 读写分离负载均衡
>     upstream mysql_master {
>         server db-master:3306;
>     }
>     
>     upstream mysql_slaves {
>         server db-slave1:3306;
>         server db-slave2:3306;
>     }
>     
>     # 写操作端口
>     server {
>         listen 13306;
>         proxy_pass mysql_master;
>     }
>     
>     # 读操作端口  
>     server {
>         listen 13307;
>         proxy_pass mysql_slaves;
>     }
> }
> ```
>
> **特点**：客户端应用连接到 `13306` 端口进行写操作，连接到 `13307` 端口进行读操作，Nginx 只做 TCP 层面的转发。
>
> #### 例2：SSH 连接负载均衡
>
> ```nginx
> stream {
>     upstream ssh_servers {
>         server 192.168.1.100:22;
>         server 192.168.1.101:22;
>         server 192.168.1.102:22;
>     }
>     
>     server {
>         listen 2222;
>         proxy_pass ssh_servers;
>     }
> }
> ```
>
> **特点**：用户 SSH 连接到 `2222` 端口，Nginx 将连接转发到后端 SSH 服务器之一。
>
> #### 例3：DNS 负载均衡
>
> ```nginx
> stream {
>     upstream dns_servers {
>         server 8.8.8.8:53;
>         server 1.1.1.1:53;
>     }
>     
>     server {
>         listen 5353 udp;  # UDP 协议
>         proxy_pass dns_servers;
>     }
> }
> ```
>
> 
>
> ### 七层负载均衡用例
>
> #### 例1：基于 URL 路径的路由
>
> ```nginx
> http {
>     upstream app_servers {
>         server 192.168.1.30:8001;
>         server 192.168.1.31:8001;
>     }
>     
>     upstream api_servers {
>         server 192.168.1.40:8002;
>         server 192.168.1.41:8002;
>     }
>     
>     upstream static_servers {
>         server 192.168.1.50:8003;
>         server 192.168.1.51:8003;
>     }
>     
>     server {
>         listen 80;
>         
>         # 根据 URL 路径路由到不同后端
>         location /app/ {
>             proxy_pass http://app_servers;
>         }
>         
>         location /api/ {
>             proxy_pass http://api_servers;
>         }
>         
>         location /static/ {
>             proxy_pass http://static_servers;
>         }
>     }
> }
> ```
>
> 
>
> #### 例2：基于 Cookie 的会话保持
>
> ```nginx
> http {
>     upstream backend {
>         sticky cookie srv_id expires=1h domain=.example.com path=/;
>         server 192.168.1.60:8080;
>         server 192.168.1.61:8080;
>         server 192.168.1.62:8080;
>     }
>     
>     server {
>         listen 80;
>         
>         location / {
>             proxy_pass http://backend;
>             # 会话保持，同一用户总是访问同一后端
>         }
>     }
> }
> ```
>
> 
>
> #### 例3：基于 Header 的条件路由
>
> ```nginx
> http {
>     map $http_user_agent $backend {
>         default        web_servers;
>         "~*mobile"     mobile_servers;
>         "~*bot"        bot_servers;
>     }
>     
>     upstream web_servers {
>         server 192.168.1.70:80;
>     }
>     
>     upstream mobile_servers {
>         server 192.168.1.71:80;
>     }
>     
>     upstream bot_servers {
>         server 192.168.1.72:80;
>     }
>     
>     server {
>         listen 80;
>         
>         location / {
>             proxy_pass http://$backend;
>             # 根据 User-Agent 路由到不同后端
>         }
>     }
> }
> ```
>
> 
>
> ## 实际部署建议
>
> ### 混合部署架构
>
> ```
> 客户端
>     ↓
> 四层负载均衡 (L4) - 用于 TCP/UDP 服务
>     ↓
> 七层负载均衡 (L7) - 用于 HTTP/HTTPS 服务  
>     ↓
> 后端服务器集群
> ```
>
> 
>
> ### 配置示例：混合使用
>
> ```nginx
> # 四层：处理数据库、SSH等
> stream {
>     upstream mysql_cluster {
>         server db1:3306;
>         server db2:3306;
>     }
>     
>     server {
>         listen 13306;
>         proxy_pass mysql_cluster;
>     }
> }
> 
> # 七层：处理 Web 应用
> http {
>     upstream web_app {
>         server web1:80;
>         server web2:80;
>     }
>     
>     server {
>         listen 80;
>         location / {
>             proxy_pass http://web_app;
>         }
>     }
> }
> ```
>
> > 这个stream块配置了一个监听在13306端口的TCP（默认）代理。当有客户端连接到这个端口时，Nginx会将连接代理到mysql_cluster定义的后端服务器（db1:3306或db2:3306）。
> >
> > 注意：在四层代理中，没有HTTP协议的概念，因此不会解析HTTP请求头中的Host（server_name）字段。它仅仅是在TCP层进行流量转发。所以，当任何客户端连接到这个13306端口时，无论客户端是谁（无论什么IP，无论什么域名，因为四层根本不管域名），都会被转发到mysql_cluster中的某一台服务器。
> >
> > 而七层代理（http块）则不同，它解析HTTP协议，可以根据HTTP头中的Host字段（即server_name）来区分不同的请求，从而将请求转发到不同的上游（upstream）或做不同的处理。
>
> ## 选择建议
>
> **使用四层负载均衡当：**
>
> - 需要高性能的 TCP/UDP 代理
> - 处理非 HTTP 协议（数据库、SSH、DNS等）
> - 只需要简单的连接分发
>
> **使用七层负载均衡当：**
>
> - 需要基于 HTTP 内容进行路由
> - 需要 SSL 终止、内容压缩等 HTTP 特性
> - 需要基于 Cookie 的会话保持
> - 需要 URL 重写、Header 修改等
>
> ## 总结
>
> - **四层**：简单高效，工作在传输层，适合非 HTTP 服务
> - **七层**：功能丰富，工作在应用层，适合 Web 应用
> - **选择依据**：根据业务需求、协议类型、性能要求来决定
>
> 在实际生产环境中，经常是两者结合使用，发挥各自的优势！



## Nginx缓存集成

之前，我们学习了如何在**浏览器进行缓存**，而本内容学习的是 Nginx。

### 缓存的概念

缓存就是数据交换的缓冲区（称作：Cache），当用户要获取数据的时候，会先从缓存中去查询获取数据，如果缓存中有就会直接返回给用户，如果缓存中没有，则会发请求从服务器重新查询数据，将数据返回给用户的同时将数据放入缓存，下次用户就会直接从缓存中获取数据。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.371wkgon7280.webp)

缓存其实在很多场景中都有用到，比如：

| 场景             | 作用                     |
| ---------------- | ------------------------ |
| 操作系统磁盘缓存 | 减少磁盘机械操作         |
| 数据库缓存       | 减少文件系统的IO操作     |
| 应用程序缓存     | 减少对数据库的查询       |
| Web 服务器缓存   | 减少对应用服务器请求次数 |
| 浏览器缓存       | 减少与后台的交互次数     |

缓存的优点

- 减少数据传输，节省网络流量，加快响应速度，提升用户体验
- 减轻服务器压力
- 提供服务端的高可用性

缓存的缺点

- 数据的不一致
- 增加成本

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.2itt4ge9erq0.webp)

在 [静态资源部署 - 缓存配置](https://frxcat.fun/middleware/Nginx/Nginx_Static_resource_deployment/#静态资源缓存配置) 的时候，我们学习了如何在浏览器进行缓存，而本内容学习的是 Nginx。

Nginx 作为 Web 服务器，Nginx 作为 Web 缓存服务器，它介于客户端和应用服务器之间，当用户通过浏览器访问一个 URL 时，Web 缓存服务器会去应用服务器获取要展示给用户的内容，将内容缓存到自己的服务器上，当下一次请求到来时，如果访问的是同一个 URL，Web 缓存服务器就会直接将之前缓存的内容返回给客户端，而不是向应用服务器再次发送请求。

Web 缓存降低了应用服务器、数据库的负载，减少了网络延迟，提高了用户访问的响应速度，增强了用户的体验。

### Web缓存服务

Nginx 是从 0.7.48 版开始提供缓存功能。Nginx 是基于 Proxy Store 来实现的，

**其原理是把 URL 及相关组合当做 Key，在使用 MD5 算法对 Key 进行哈希化，得到硬盘上对应的哈希目录路径，从而将缓存内容保存在该目录中**。

它可以支持任意 URL 连接，同时也支持 404/301/302 这样的非200 状态码。

Nginx 即可以支持对指定 URL 或者状态码设置过期时间，也可以使用 purge 命令来手动清除指定 URL 的缓存。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.5rf186zoffg0.webp)

### 缓存设置相关指令

Nginx 的 Web 缓存服务主要是使用 `ngx_http_proxy_module` 模块相关指令集来完成，接下来我们把常用的指令来进行介绍下。

[ngx_http_proxy_module 文档地址(opens new window)](http://nginx.org/en/docs/http/ngx_http_proxy_module.html)

#### proxy_cache_path

该指定用于设置缓存文件的存放路径。

| 语法                                                         | 默认值 | 位置 |
| ------------------------------------------------------------ | ------ | ---- |
| proxy_cache_path <path> [levels=number] <keys_zone=zone_name:zone_size> [inactive=time][max_size=size]; | —      | http |



- `path`：缓存路径地址，如：

```sh
/usr/local/proxy_cache
```



- `levels`: 指定该缓存空间 `path` 基础上新建的目录，最多可以设置 3 层，每层取 1 到 2 个字母作为目录名，格式为：

```nginx
levels=num:num:num   # 三个 num 代表三层，每层目录名分别取 num 个字母
levels=num:num       # 两个 num 代表两层，每层目录名分别取 num 个字母
levels=num           # 一个 num 代表一层，每层目录名分别取 num 个字母
```

如：

```nginx
levels=1:2   # 缓存空间有两层目录，第一层目录名取 1 个字母，第二层目录名取 2 个字母
```

字母名从 MD5 加密的值后面往前截取。

举例说明：

```nginx
# 假设 proxy_cache_key 为 kele，通过 MD5 加密以后的值为 27ce47ea65c1381dbe5175f7c77d8a3a
levels=1:2    # 最终的存储路径为 /usr/local/proxy_cache/a/a3，每层截取个数根据 1:2
levels=2:1:2  # 最终的存储路径为 /usr/local/proxy_cache/3a/a/d8，每层截取个数根据 2:1:2
levels=2:2:2  # 最终的存储路径为 /usr/local/proxy_cache/3a/8a/7d，每层截取个数根据 2:2:2
```

还不理解吗？存储路径在 `path` 目录基础上再创建新的目录，新的目录名从加密后的值的后面往前面截取。



- `keys_zone`：用来为这个缓存区设置名称和指定大小，如：

```nginx
keys_zone=kele:200m  # 缓存区的名称是 kele，大小为 200M，1M 大概能存储 8000 个 keys
```



- `inactive`：指定的时间内未访问的缓存数据会从缓存中删除，默认情况下，`inactive` 设置为 10 分钟。如：

```nginx
inactive=1d   # 缓存数据在 1 天内没有被访问就会被删除
```



- `max_size`：设置最大缓存空间，如果缓存空间存满，默认会覆盖缓存时间最长的资源，默认单位为兆。如：

```nginx
max_size=20g    # 最大缓存空间为 20G
```

配置实例:

```nginx
http{
	proxy_cache_path /usr/local/proxy_cache keys_zone=kele:200m levels=1:2:1 inactive=1d max_size=20g;
}
```

此时重启 Nginx 配置文件，发现 `/usr/local` 目录里多出一个目录，名字叫做 proxy_cache。



#### proxy_cache

该指令用来开启或关闭代理缓存，如果是开启则自定使用哪个缓存区来进行缓存。默认关闭。

| 语法                            | 默认值           | 位置                   |
| ------------------------------- | ---------------- | ---------------------- |
| proxy_cache <zone_name \| off>; | proxy_cache off; | http、server、location |



zone_name：指定使用缓存区的名称。

注意

**缓存区的名称必须是 `proxy_cache_path` 里的 `keys_zone` 生成的缓存名。**



#### proxy_cache_key

该指令用来设置 Web 缓存的 key 值，Nginx 会根据 key 值利用 MD5 计算处哈希值并缓存起来，作为缓存目录名的参考。

| 语法                   | 默认值                                  | 位置                   |
| ---------------------- | --------------------------------------- | ---------------------- |
| proxy_cache_key <key>; | proxy_cache_key proxy_host$request_uri; | http、server、location |



如 kele 由 MD5 计算出来是 27ce47ea65c1381dbe5175f7c77d8a3a

在哪计算出来的？ 前往 [MD5 在线加密网站(opens new window)](https://md5jiami.bmcx.com/)



#### proxy_cache_valid

**该指令用来对不同返回状态码的 URL 设置不同的缓存时间。**

| 语法                                     | 默认值 | 位置                   |
| ---------------------------------------- | ------ | ---------------------- |
| proxy_cache_valid [code ...... ] <time>; | —      | http、server、location |



如：

```nginx
proxy_cache_valid 200 302 10m; # 为 200 和 302 的响应 URL 设置 10 分钟缓存时间
proxy_cache_valid 404 1m;      # 为 404 的响应 URL 设置 1 分钟缓存时间
proxy_cache_valid any 1m;      # 对所有响应状态码的URL都设置 1 分钟缓存时间
```



#### proxy_cache_min_uses

该指令用来设置资源被访问多少次后才会被缓存。默认是 1 次。

| 语法                           | 默认值                  | 位置                   |
| ------------------------------ | ----------------------- | ---------------------- |
| proxy_cache_min_uses <number>; | proxy_cache_min_uses 1; | http、server、location |



#### proxy_cache_methods

该指令是设置缓存哪些 HTTP 方法的请求资源。

| 语法                                       | 默认值                        | 位置                   |
| ------------------------------------------ | ----------------------------- | ---------------------- |
| proxy_cache_methods <GET \| HEAD \| POST>; | proxy_cache_methods GET HEAD; | http、server、location |



默认缓存 HTTP 的 GET 和 HEAD 方法的请求资源，不缓存 POST 方法的请求资源。



#### 缓存设置案例

##### 需求分析

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.4ye8ddxwbng0.webp)

##### 步骤实现

**应用服务器的环境准备**

1、在 `192.168.200.146` 服务器 A 上的 tomcat 的 webapps 下面添加一个 js 目录，并在 js 目录中添加一个 jquery.js 文件

2、启动 tomcat

```sh
cd /usr/local/tomcat/bin
./startup.sh
```

3、访问服务器 A 进行测试

```http
http://192.168.200.146:8080/js/jquery.js
```

**Nginx 的环境准备**

1、准备服务器 B 完成 Nginx 的反向代理配置

```nginx
http{
	upstream backend{
		server 192.168.200.146:8080;   # 服务器 A 地址
	}
	server {
		listen       8080;
        server_name  localhost;
        location / {
        	proxy_pass http://backend/js/;
        }
	}
}
```

2、完成 Nginx 缓存配置

```nginx
http{
    proxy_cache_path /usr/local/proxy_cache levels=2:1 keys_zone=bing:200m inactive=1d max_size=20g;
    upstream backend{
        server 192.168.200.146:8080;   # 服务器 A 的地址
    }
    server {
        listen       8080;     			# 监听 8080 端口
        server_name  localhost; 		# 监听 localhost 的IP
        location / {					# 监听包含 / 的请求
            proxy_cache bing;    		# 开启 bing 缓存区，和第 2 行的 keys_zone 对应
            proxy_cache_key kele;  		# 缓存的 key 值，会被 MD5 解析成字符串用于生成缓存的目录
            proxy_cache_min_uses 5; 	# 资源被访问 5 次后才会被缓存
            proxy_cache_valid 200 5d;	# 为 200 响应 URL 设置 5 天缓存时间
            proxy_cache_valid 404 30s;  # 为 404 的响应 URL 设置 30 秒缓存时间
            proxy_cache_valid any 1m;	# 为除了上方的任意响应 URL 设置 1 分钟缓存时间
            add_header nginx-cache "$upstream_cache_status";  # 将缓存的状态放到请求头里
            proxy_pass http://backend/js/;  # 代理 backend，将 /js/ 追加到 backend 模块里的地址后面
        }
    }
}
```

> 这种情况下 该访问下所有的key都是 kele 因为目前资源比较少，所以ok 但是如果资源不同的话 必须使用完整的 proxy_cache_key
>
> 所有请求都用同一个缓存键 `kele`，导致：
>
> - 不同路径的资源会相互覆盖
> - 只有最后一个请求的资源会被缓存
> - 所有用户都得到相同的缓存内容
>
>
> 更完整的配置    
>
> ```nginx
> http {
>     proxy_cache_path /usr/local/proxy_cache levels=2:1 keys_zone=bing:200m inactive=1d max_size=20g;
>     
> 	upstream backend {
>     	server 192.168.200.146:8080;
> 	}
> 
> 	server {
>     	listen 8080;
>     	server_name localhost;
>     
>     	location / {
>         	# 正确的缓存键：包含完整请求信息
>         	proxy_cache_key "$scheme$request_method$host$request_uri";
>         
>         	proxy_cache bing;
>         	proxy_cache_min_uses 5;
>         	proxy_cache_valid 200 5d;
>         	proxy_cache_valid 404 30s;
>         	proxy_cache_valid any 1m;
>         	add_header nginx-cache "$upstream_cache_status";
>         
>         	# 根据需求选择正确的代理方式
>         	proxy_pass http://backend;
>     	}
> 	}
> }
> ```
> 
>
> ## 不同代理方式的对比
>
> ### 方式1：路径完全保留（推荐）
>
> nginx
>
> ```nginx
> location /api/ {
>     proxy_pass http://backend;  # 不带斜杠
> }
> ```
>
> 
>
> - 请求：`/api/users/list`
>
> - 转发：`http://backend/api/users/list`
>
> - ✅ **保持原路径结构**
>
>   
>
> ### 方式2：路径替换
>
> ```nginx
> location /api/ {
>     proxy_pass http://backend/v1/;  # 带斜杠 - 替换 location 匹配部分
> }
> ```
>
> - 请求：`/api/users/list`
>
> - 转发：`http://backend/v1/users/list`
>
> - ⚠️ **路径被修改**：`/api/` 被替换为 `/v1/`
>
>   
>
> ### 方式3：路径重写 + 代理
>
> ```nginx
> location /legacy/ {
>     rewrite ^/legacy/(.*)$ /new/$1 break;  # 先重写路径
>     proxy_pass http://backend;              # 再代理
> }
> ```
>
> - 请求：`/legacy/users/list`
> - 转发：`http://backend/new/users/list`
> - ✅ **灵活控制路径转换**

3、测试是否缓存成功

利用 `$upstream_cache_status` 的值在控制台(F12)查看是否缓存。

第一次访问 `192.168.200.113:8080/jquery.js`，如图：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.3p1azxp6cdy0.webp)

因为第一次访问时，正在缓存，所以返回的请求头 MISS 是没有缓存成功。

第二次访问 `192.168.200.113:8080/jquery.js`，如图：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.73w7kaxo8r00.webp)

HIT 代表成功缓存。

1. 测试 404 缓存时间

测试 404 缓存时间，我们访问 `192.168.200.113:8080/jquery.js111`，它会返回 404 页面，并缓存 404 页面，当我们立即访问正确的 `192.168.200.113:8080/jquery.js`，它依然返回 404 页面，因为 `/jquery.js` 请求目前被缓存为 404，还没到 30 秒过期，等 30 秒后再访问，就成功了。



#### 缓存的删除

这里介绍两种方式：

- 删除对应的缓存目录
- 使用第三方扩展模块

- ##### 删除缓存目录

假设缓存目录是 `/usr/local/proxy_cache/`

```sh
rm -rf /usr/local/proxy_cache/......
```

如果想删除某个缓存目录，就在后面加上目录名。如果想删除整个缓存目录，直接删除 `/usr/local/proxy_cache/` 即可。



- ##### ngx_cache_purge删除

使用第三方扩展模块 `ngx_cache_purge` 进行删除缓存。

1、下载 `ngx_cache_purge` 模块对应的资源包，并上传到服务器的 `/root/nginx/module/` 目录下。

这里的资源包是 ngx_cache_purge-2.3.tar.gz

2、对资源文件进行解压缩

```sh
tar -zxf ngx_cache_purge-2.3.tar.gz
```

3、修改文件夹名称为 `purge`，方便后期配置

```sh
mv ngx_cache_purge-2.3 purge
```

4、查询 Nginx 的配置参数 `configure arguments`，并拷贝出来

```sh
nginx -V
```

5、进入 Nginx 的安装包目录，使用 ./configure 进行参数配置，记得加上 `nginx -V` 查询出来的 `configure arguments` 参数

```sh
./configure --add-module=/root/nginx/module/purge  # 加上之前的 configure arguments 参数
```

6、使用 make 进行编译

```sh
make
```

7、将 Nginx 安装目录的 nginx 二级制可执行文件备份

```sh
mv /usr/local/nginx/sbin/nginx /usr/local/nginx/sbin/nginx.backup
```

8、将编译后的 objs 中的 nginx 拷贝到 nginx 的 sbin 目录下

```sh
cp objs/nginx /usr/local/nginx/sbin
```

9、使用 `make upgrade` 进行升级，记得在安装包目录下执行

```sh
cd /opt/nginx/core/nginx-1.20.2

make upgrade
```

10、在 Nginx 配置文件中进行如下配置

```nginx
server{
    location ~/purge(/.*) {
        proxy_cache_purge bing kele;
    }
}
```



`proxy_cache_purge` 指令

| 语法                            | 默认值 | 位置                   |
| ------------------------------- | ------ | ---------------------- |
| proxy_cache_purge <cache> <key> | -      | http、server、location |

- `cache` 是 `proxy_cache`，详细内容看 [proxy_cache(opens new window)](https://notes.youngkbt.cn/nginx/cache-integration/#proxy-cache)
- `key` 是 `proxy_cache_key`，详细内容看 [proxy_cache_key(opens new window)](https://notes.youngkbt.cn/nginx/cache-integration/#proxy-cache-key)



#### 资源不缓存

前面咱们已经完成了 Nginx 作为 Web 缓存服务器的使用。但是我们得思考一个问题，**不是所有的数据都适合进行缓存**。比如说对于一些经常发生变化的数据。如果进行缓存的话，就很容易出现用户访问到的数据不是服务器真实的数据。所以对于这些资源我们在缓存的过程中就需要进行过滤，不进行缓存。

Nginx 也提供了这块的功能设置，需要使用到如下两个指令：

- proxy_no_cache
- proxy_cache_bypass

##### proxy_no_cache

该指令是用来定义不将数据进行缓存的条件，也就是不缓存指定的数据。

| 语法                             | 默认值 | 位置                   |
| -------------------------------- | ------ | ---------------------- |
| proxy_no_cache <string> ...... ; | —      | http、server、location |



可设置多个 string。

配置实例：

```nginx
proxy_no_cache $cookie_nocache $arg_nocache $arg_comment;
```

##### proxy_cache_bypass

该指令是用来设置不从缓存中获取数据的条件，也就是虽然缓存了指定的资源，但请求过来也不会去获取它，而是去服务器里获取资源。

| 语法                                 | 默认值 | 位置                   |
| ------------------------------------ | ------ | ---------------------- |
| proxy_cache_bypass <string> ...... ; | —      | http、server、location |



可设置多个 string。

配置实例：

```nginx
proxy_cache_bypass $cookie_nocache $arg_nocache $arg_comment;
```

上述两个指令都有一个指定的条件，这个条件可以是多个，**并且多个条件中至少有一个不为空且不等于「0」，则条件满足成立。**

上面给的配置实例是从官方网站获取的，里面使用到了三个变量，分别是 `$cookie_nocache`、`$arg_nocache`、`$arg_comment`

#### 常用不缓存变量

常用不缓存的三个变量分别为：

- `$cookie_nocache`
- `$arg_nocache`
- `$arg_comment`

这三个变量分别代表的含义是:

- `$cookie_nocache`：指的是当前请求的 cookie 中 key 为 nocache 的 value 值
- `$arg_nocache` 和 `$arg_comment`：指的是当前请求的参数中属性名为 nocache 和 comment 对应的属性值

案例演示:

```nginx
log_format params $cookie_nocache | $arg_nocache | $arg_comment；
server {
    listen	8081;
    server_name localhost;
    location / {
        access_log logs/access_params.log params;
        add_header Set-Cookie 'nocache=888';
        root html;
        index index.html;
    }
}
```

访问 `192.168.200.133:8081?nocache=999&comment=777`，然后去日志查看结果，如图所示：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220805/image.695kmtpob840.webp)

以后访问的某一个资源如果不想缓存，在 URL 后面加入三个变量中的任意一个或多个即可，只要它们不为空或 0。

注意

这三个变量推荐作为不缓存资源的条件，但并不是只能作为不缓存资源的条件。

#### 案例模板

设置不缓存资源的配置方案模板：

- 如果访问的是 js 文件，则不会缓存该 js 文件
- 如果 `$nocache` `$cookie_nocache` `$arg_nocache` `$arg_comment` 任意不为空或 0，则访问的资源不进行缓存



```nginx
server {
    listen	8080;
    server_name localhost;
    location / {
        if ($request_uri ~ /.*\.js$){
            set $nocache 1;
        }
        proxy_no_cache $nocache $cookie_nocache $arg_nocache $arg_comment;
        proxy_cache_bypass $nocache $cookie_nocache $arg_nocache $arg_comment;
    }
}
```

为什么不会缓存 js 文件呢，看第 5 - 6 行代码。如果访问的文件是 js 文件，则设置 `$nocache` 为 1，只要它不为 0，则触发第 8 行代码，`proxy_no_cache` 后面的参数只要有一个不为空或 0，则访问的资源不进行缓存。





## Nginx 部署和集群

### Nginx与Tomcat部署

前面已经将 Nginx 的大部分内容进行了讲解，我们都知道了 Nginx 在高并发场景和处理静态资源是非常高性能的，但是在实际项目中除了静态资源还有就是后台业务代码模块，一般后台业务都会被部署在 Tomcat、weblogic 或者是 websphere 等 Web 服务器上。

那么如何使用 Nginx 接收用户的请求并把请求转发到后台 Web 服务器？

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.iachagf6u0g.webp)

步骤分析:

- 在服务器 A 上准备 Tomcat 环境，并在 Tomcat 上部署一个 Web 项目。这步骤在 [环境准备(Tomcat)](https://frxcat.fun/middleware/Nginx/Nginx_Deployment_and_cluster/#环境准备-tomcat)
- 在服务器 B 上准备 Nginx 环境，使用 Nginx 接收请求，并把请求分发到 Tomcat 上。这步骤在 [环境准备(Nginx)](https://frxcat.fun/middleware/Nginx/Nginx_Deployment_and_cluster/#环境准备-nginx)

#### 环境准备(Tomcat)

本次将采用 Tomcat 作为后台 Web 服务器。

- 在 服务器 A 上准备一个 Tomcat
  - Tomcat 官网地址：[https://tomcat.apache.org/(opens new window)](https://tomcat.apache.org/)
- 下载 Tomcat，这里使用的是 `apache-tomcat-9.0.54.tar.gz`
- 进入上传目录，将 Tomcat 进行解压缩

```sh
mkdir /usr/local/tomcat

tar -zxf apache-tomcat-9.0.54.tar.gz -C /usr/local/tomcat
```

准备一个 Web 项目，将其打包为 War 包，这里是 demo.war

- 将写好的 War 包上传到 Tomcat 目录下的 webapps 包下
- 将 Tomcat 进行启动，进入 Tomcat 的 bin 目录下，执行命令：

```sh
./startup.sh
```

启动完成 Tomcat 后，进行访问测试

浏览器访问：

- 静态资源：`http://192.168.200.146:8080/demo/index.html`

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.pln0ovd1h1c.webp)

- 动态资源：`http://192.168.200.146:8080/demo/getAddress`

动态资源可以是端口号，此时的端口是 8080![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.31nc8e1cdee0.webp)

自此，服务器 A 的 Tomcat 部署已经实现。

**demo.war 的内容有什么呢？**

其实你可以自己制作一个 war 包，这里说明一下，demo.war 里有两个图片，和一个 index.html

index.html 文件引用了两个图片：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="js/jquery.min.js"></script>
    <script>
        $(function(){
           $.get('http://192.168.200.133/demo/getAddress',function(data){
               $("#msg").html(data);
           });
        });
    </script>
</head>
<body>
    <img src="images/logo.png"/>
    <h1>Nginx如何将请求转发到后端服务器</h1>
    <h3 id="msg"></h3>
    <img src="images/mv.png"/>
</body>
</html>
```

可以看出，当访问 index.html 时，它会主动去请求 `/getAddress`，这个请求返回端口号，即把 8080 当作动态资源，显示到页面上，如果是 9090 端口访问，则返回 9090 显示页面上。你也可以直接访问 `/getAddress`，直接获取端口号。

#### 环境准备(Nginx)

我们已经在 Tomcat 实现了两个效果，那么现在需要把 Tomcat 的地址放到 Nginx 里，由 Nginx 帮我们代理这个 Tomcat 地址，这样我们访问 Nginx，实际上就是访问 Tomcat。

1、使用 Nginx 的反向代理，将请求转给 Tomcat 进行处理。

```nginx
upstream webservice {
	server 192.168.200.146:8080;  # 服务器 A 的 Tomcat  地址
}
server{
    listen		80;
    server_name localhost;
    location /demo {
    	proxy_pass http://webservice;
    }
}
```

2、启动访问服务器 B，测试是否代理到服务器 A 的 Tomcat，效果如图：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.6taygwj4nt80.webp)

学习到这，可能大家会有一个困惑，明明直接通过 Tomcat 就能访问，为什么还需要多加一个 Nginx，这样不是反而是系统的复杂度变高了么? 那接下来我们从两个方便给大家分析下这个问题，

- 第一个使用 Nginx 实现动静分离
- 第二个使用 Nginx 搭建 Tomcat 的集群

### 动静分离

**什么是动静分离？**

- 动：后台应用程序的业务处理
- 静：网站的静态资源（html，javaScript，css，images 等文件）
- 分离：将两者进行分开部署访问，提供用户进行访问。

**举例说明就是以后所有和静态资源相关的内容都交给 Nginx 来部署访问，非静态内容则交个类似于 Tomcat 的服务器来部署访问。**

**为什么要动静分离？**

前面我们介绍过 Nginx 在处理静态资源的时候，效率是非常高的，而且 Nginx 的并发访问量也是名列前茅，而 Tomcat 则相对比较弱一些，所以把静态资源交给 Nginx 后，可以减轻 Tomcat 服务器的访问压力并提高静态资源的访问速度。

动静分离以后，降低了动态资源和静态资源的耦合度。如动态资源宕机了也不影响静态资源的展示。

**如何实现动静分离?**

实现动静分离的方式很多，比如静态资源可以部署到 CDN、Nginx 等服务器上，动态资源可以部署到 Tomcat、weblogic 或者 websphere 上。这里使用 Nginx + Tomcat 来实现动静分离。

#### 需求分析

如下图，因为 Nginx 处理静态资源性能高，所以我们把静态资源放在 Nginx 服务器上，然后把动态资源放到 Tomcat 服务器上。当访问 Nginx 的静态资源时，Nginx 会去访问 Tocmat 获取动态资源。实现动静分离。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.2v1uen08t3c0.webp)

#### 实现步骤

- 将 demo.war 项目中的静态资源（两个图片）都删除掉，重新打包生成一个 War 包

  这时候 War 包只留下动态资源，而静态资源要放到 Nginx 上。

- 将新的 War 包部署到 Tomcat 中，把之前部署的内容删除掉

  - 进入到 tomcat 的 webapps 目录下，将之前的 demo 目录和 demo.war 包删除掉
  - 将新的 War 包复制到 webapps 下
  - 将 tomcat 启动

- 在 Nginx 所在的服务器 B 上创建如下目录，并将对应的静态资源放入指定的位置

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.3e4j68pre4k0.webp)



```sh
mkdir /usr/local/nginx/html/web/images
mkdir /usr/local/nginx/html/web/js

cp logo.png /usr/local/nginx/html/web/images
cp mv.png /usr/local/nginx/html/web/images
cp jquery.min.js /usr/local/nginx/html/web/js

vim /usr/local/nginx/html/web/index.html
```

其中 index.html 页面的内容如下:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <script src="js/jquery.min.js"></script>
    <script>
        $(function(){
           $.get('http://192.168.200.133/demo/getAddress',function(data){
               $("#msg").html(data);
           });
        });
    </script>
</head>
<body>
    <img src="images/logo.png"/>
    <h1>Nginx如何将请求转发到后端服务器</h1>
    <h3 id="msg"></h3>
    <img src="images/mv.png"/>
</body>
</html>
```

第 9 行代码，它访问的这个地址是服务器 B 的 Nginx，通过这个地址让 Nginx 去获取服务器 A 的 Tomcat 动态资源。

- 在配置文件配置 Nginx 的静态资源与动态资源的访问

```nginx
upstream webservice{
    server 192.168.200.146:8080;  # 服务器 A 的 Tocmat
}
server {
    listen       80;
    server_name  localhost;

    # 动态资源从 Tomcat 获取
    location /demo {     		 # index.html 第 9 行代码触发该 location
        proxy_pass http://webservice;
    }
    # 静态资源从自己身上获取
    location ~/.*\.(png|jpg|gif|js){
        root html/web;
        gzip on;
    }

    location / {
        root   html/web;
        index  index.html index.htm;
    }
}
```

- 启动测试，访问 `http://192.168.200.133`

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.2yip3al9lva0.webp)

假如某个时间点，由于某个原因导致 Tomcat 后的服务器宕机了，我们再次访问 Nginx，会得到如下效果：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.3s0ttch8r4i0.webp)

用户还是能看到页面，只是缺失 Tomcat 的动态资源，这就是前后端耦合度降低的效果，并且整个请求只和后的服务器交互了一次，js 和 images 都直接从 Nginx 服务器里返回，提供了效率，降低了后端服务器的压力。



### 集群搭建

#### Tomcat集群搭建

在使用 Nginx 和 Tomcat 部署项目的时候，我们使用的是一台 Nginx 服务器和一台 Tomcat 服务器，效果图如下：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.4ukvicg3of80.webp)

那么问题来了，如果 Tomcat 的真的宕机了，整个系统就会不完整，所以如何解决上述问题？

一台服务器容易宕机，那就多搭建几台 Tomcat 服务器，这样的话就提升了后的服务器的可用性。这也就是我们常说的集群，搭建 Tomcat 的集群需要用到了 Nginx 的反向代理和赋值均衡的知识，具体如何来实现？

我们先来分析下原理：

用户请求到 Nginx，Nginx 使用负载均衡对三个 Tomcat 服务器进行访问，如果一个 Tomcat 服务器宕机了，那么还有两个 Tomcat 服务器可以使用。

如图：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.664h3o0m8y00.webp)

##### 环境搭建

1. 准备 3 台 Tomcat 服务器，使用端口进行区分（实际环境应该是三台服务器），修改 Tomcat 的 server.xml，将端口修改分别修改为 8080、8180、8280



```sh
cp -r apache-tomcat-9.0.54 tomcat01
cp -r apache-tomcat-9.0.54 tomcat02
cp -r apache-tomcat-9.0.54 tomcat03
```

修改三个 Tocmat 配置文件的端口

```sh
vim tomcat01/conf/server.xml
vim tomcat02/conf/server.xml
vim tomcat03/conf/server.xml
```

修改的内容位置如下：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.76l3jwkc3b40.webp)

1、在 Nginx 对应的配置文件中添加如下内容:

```nginx
upstream webservice{
    server 192.168.200.146:8080;     # tomcat01
    server 192.168.200.146:8180; 	 # tomcat02
    server 192.168.200.146:8280;	 # tomcat03
}
server{
    listen		80;
    server_name localhost;
    location /demo {
        proxy_pass http://webservice;
    }
}
```

2、启动 Tomcat 并访问测试

```http
http://192.168.200.146:8080/demo/getAddress
```

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.2389oc46sdxc.webp)



```http
http://192.168.200.146:8180/demo/getAddress
```

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.1sjv5fphqk8w.webp)



```http
http://192.168.200.146:8280/demo/getAddress
```

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.6uckjob0ptc0.webp)

好了，完成了上述环境的部署，我们已经解决了 Tomcat 的高可用性，一台服务器宕机，还有其他两条对外提供服务，同时也可以实现后台服务器的不间断更新。

**但是新问题出现了，上述环境中，如果是 Nginx 宕机了呢，那么整套系统都将服务对外提供服务了，这个如何解决？**

之前的设置中 nginx虽然有多个的worker 但是服务器宕机后 这里的worker会全部不可用



#### Nginx集群搭建

针对于上面提到的问题，我们来分析下要想解决上述问题，需要面临哪些问题？

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.664h3o0m8y00.webp)

需要两台以上的 Nginx 服务器对外提供服务，这样的话就可以解决其中一台宕机了，另外一台还能对外提供服务，但是如果是两台 Nginx 服务器的话，会有两个 IP 地址，用户该访问哪台服务器，用户怎么知道哪台是好的，哪台是宕机了的？

> #### Keepalived
>
> 使用 Keepalived 来解决，Keepalived 软件由 C 编写的，最初是专为 LVS 负载均衡软件设计的，Keepalived 软件主要是通过 VRRP 协议实现高可用功能。
>
> #### VRRP介绍
>
> VRRP（Virtual Route Redundancy Protocol）协议，翻译过来为虚拟路由冗余协议。
>
> **VRRP 协议将两台或多台路由器设备虚拟成一个设备，对外提供虚拟路由器 IP，而在路由器组内部，如果实际拥有这个对外 IP 的路由器如果工作正常的话就是 MASTER，MASTER 实现针对虚拟路由器IP的各种网络功能。**
>
> 其他设备不拥有该虚拟 IP，状态为 BACKUP，处了接收 MASTER 的 VRRP 状态通告信息以外，不执行对外的网络功能。当主机失效时，BACKUP 将接管原先 MASTER 的网络功能。
>
> 看图分析：VRRP 把两个 Nginx 分成两个路由（VRRP 路由 1 和 VRRP 路由 2），并生成一个 Virtual 路由，用户访问的是 Virtual 路由，该路由会去访问两个 Nginx 生成的 VRRP 路由。那么到底访问谁呢？VRRP 会给两个路由分配角色，一个是 Master（老大），另一个是 Backup（备份），所以访问的是 Master 角色的路由，当 Master 角色路由宕机了，才会找到 Backup 备份路由。
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.zc2md5tuu0w.webp)
>
> 从上面的介绍信息获取到的内容就是 VRRP 是一种协议，那这个协议是用来干什么的？
>
> 1. 选择协议
>
>    VRRP 可以把一个虚拟路由器的责任动态分配到局域网上的 VRRP 路由器中的一台。其中的虚拟路由即 Virtual 路由是由 VRRP 路由群组创建的一个不真实存在的路由，这个虚拟路由也是有对应的 IP 地址。而且 VRRP 路由 1 和 VRRP 路由 2 之间会有竞争选择，通过选择会产生一个 Master 路由和一个 Backup 路由。
>
> 2. 路由容错协议
>
>    Master 路由和 Backup 路由之间会有一个**心跳检测**，Master 会定时告知 Backup 自己的状态，如果在指定的时间内，Backup 没有接收到这个通知内容，Backup 就会替代 Master 成为新的 Master。Master 路由有一个特权就是虚拟路由和后端服务器都是通过 Master 进行数据传递交互的，而备份节点则会直接丢弃这些请求和数据，不做处理，只是去监听 Master 的状态。
>
> 用了 Keepalived 后，解决方案如图下:
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.3n6clkv1blk0.webp)
>
> 看图分析：VIP 是虚拟路由，是专门给用户发送请求。一旦用户发送请求到 VIP，VIP 就会发送给 Master（主）的 Nginx，如果 Master（主）Nginx 宕机了，才会发送给 Backup（备份） Nginx 路由。

##### 环境搭建

环境准备

| VIP IP          | Nginx IP                    | 主机名      | 主/从  |
| --------------- | --------------------------- | ----------- | ------ |
|                 | 192.168.200.133（服务器 A） | keepalived1 | Master |
| 192.168.200.222 |                             |             |        |
|                 | 192.168.200.122（服务器 B） | keepalived2 | Backup |



**确保服务器 A 和服务器 B 的 Nginx 配置保持一致。**

keepalived 的安装步骤如下：

- 步骤1：从官方网站下载 keepalived，官网地址 [https://keepalived.org/(opens new window)](https://keepalived.org/)
- 步骤2：将下载的资源上传到服务器，这里是 `keepalived-2.0.20.tar.gz`
- 步骤3：在 `/opt` 目录下创建 keepalived 目录，方便管理资源

```sh
mkdir /opt/keepalived
```

- 步骤4：将压缩文件进行解压缩，解压缩到指定的目录

```sh
tar -zxf keepalived-2.0.20.tar.gz -C /opt/keepalived
```

- 步骤5：对 keepalived 进行配置，编译和安装

```sh
cd /opt/keepalived/keepalived-2.0.20

./configure --sysconf=/etc --prefix=/usr/local   # 安装到 /usr/local 目录下，可修改

make && make install
```

**两台 Nginx 服务器都要安装 keepalive。**

安装完成后，有两个文件需要我们认识下：

- `/etc/keepalived/keepalived.conf`：keepalived 的系统配置文件，我们主要操作的就是该文件

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.5hd9zg18lp80.webp)

- `/usr/local/sbin` 目录下的 `keepalived`：这是系统配置脚本，用来启动和关闭 keepalived

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.54vzhzf4avw0.webp)

##### Keepalived配置文件介绍

打开 keepalived.conf 配置文件

这里面会分三部：

- 第一部分是 global 全局配置
- 第二部分是 vrrp 相关配置
- 第三部分是 LVS 相关配置。

这里主要是使用 keepalived 实现高可用部署，没有用到 LVS，所以我们重点关注的是前两部分。

打开 keepalived.conf 文件

```sh
vim /etc/keepalived/keepalived.conf
```

文件内容部分介绍如下：

```nginx
# global全局部分
global_defs {
    
   notification_email {  # 通知邮件，当 keepalived 切换 Master 和 Backup 时需要发 email 给具体的邮箱地址
     tom@itcast.cn
     jerry@itcast.cn
   }
   notification_email_from kele@youngkbt.com   # 设置发件人的邮箱信息
   
   smtp_server 192.168.200.1   # 指定 smpt 服务地址
   
   smtp_connect_timeout 30   # 指定 smpt 服务连接超时时间
   
   router_id LVS_DEVEL   # 运行 keepalived 服务器的一个标识，可以用作发送邮件的主题信息
   
   # 默认是不跳过检查。检查收到的 VRRP 通告中的所有地址可能会比较耗时，设置此命令的意思是，如果通告与接收的上一个通告来自相同的 master 路由器，则不执行检查(跳过检查)
   vrrp_skip_check_adv_addr
   
   vrrp_strict    # 严格遵守 VRRP 协议
  
   vrrp_garp_interval 0   # 在一个接口发送的两个免费 ARP 之间的延迟。可以精确到毫秒级。默认是 0
   
   vrrp_gna_interval 0  # 在一个网卡上每组消息之间的延迟时间，默认为 0
}
```

这里需要修改的是 5、6、8 行代码。

VRRP 部分可以包含以下四个子模块：

1. vrrp_script
2. vrrp_sync_group
3. garp_group
4. vrrp_instance

我们会用到第一个（[vrrp_script](https://frxcat.fun/middleware/Nginx/Nginx_Deployment_and_cluster/#vrrp-script)）和第四个（[vrrp_instance](https://frxcat.fun/middleware/Nginx/Nginx_Deployment_and_cluster/#vrrp-instance)）。

- ##### vrrp_instance

`vrrp_instance` 模块内容：

```nginx
# 设置 keepalived 实例的相关信息，VI_1 为 VRRP 实例名称
vrrp_instance VI_1 {
    state MASTER  		  # 有两个值可选 MASTER 主，BACKUP 备
    interface ens33		  # vrrp 实例绑定的接口，用于发送 VRRP 包[当前服务器使用的网卡名称]
    virtual_router_id 51  # 指定 VRRP 实例 ID，范围是 0-255
    priority 100		  # 指定优先级，优先级高的将成为 MASTER
    advert_int 1		  # 指定发送 VRRP 通告的间隔，单位是秒。这里是心跳检查的时间
    authentication {	  # vrrp 之间通信的认证信息
        auth_type PASS	  # 指定认证方式。PASS 简单密码认证(推荐)
        auth_pass 1111	  # 指定认证使用的密码，最多 8 位
    }
    virtual_ipaddress {   # 虚拟 IP 地址设置虚拟 IP 地址，供用户访问使用，可设置多个，一行一个
        192.168.200.222
    }
}
```

`vrrp_instance` 模块中我们修改的是第 3、5、6、7、13 行代码。

第 3 行代码是说明当前 Nginx 服务器的角色是 Master 还是 Backup。分别在服务器 A 和 B 进行角色配置。

第 5 行代码是 VIP 的 ID，如果使用相同的虚拟路由 VIP，请保持 ID 一致。

第 6 行代码是优先级，请让 Master 服务器的优先级大于 Backup 服务器的优先级。如 100 > 90。

第 7 行代码是 Master 和 Backup 之间通信的间隔时间，如果无法通信，说明 Master 已经宕机，则切换为 Backup。

第 13 行代码是用户访问的虚拟 IP 地址，即 VIP，它会发送给 Nginx 服务器。

##### 服务器配置

Keepalived 的具体配置内容如下：

::: tab 服务器A

```nginx
global_defs {
   notification_email {
        tom@itcast.cn
        jerry@itcast.cn
   }
   notification_email_from zhaomin@itcast.cn
   smtp_server 192.168.200.1
   smtp_connect_timeout 30
   router_id keepalived1
   vrrp_skip_check_adv_addr
   vrrp_strict
   vrrp_garp_interval 0
   vrrp_gna_interval 0
}

vrrp_instance VI_1 {
    state MASTER
    interface ens33
    virtual_router_id 51
    priority 100
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1111
    }
    virtual_ipaddress {
        192.168.200.222
    }
}
```

:::

::: tab 服务器B

```nginx
! Configuration File for keepalived

global_defs {
   notification_email {
        tom@itcast.cn
        jerry@itcast.cn
   }
   notification_email_from zhaomin@itcast.cn
   smtp_server 192.168.200.1
   smtp_connect_timeout 30
   router_id keepalived2
   vrrp_skip_check_adv_addr
   vrrp_strict
   vrrp_garp_interval 0
   vrrp_gna_interval 0
}

vrrp_instance VI_1 {
    state BACKUP
    interface ens33
    virtual_router_id 51
    priority 90
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1111
    }
    virtual_ipaddress {
        192.168.200.222
    }
}
```

:::

::::

> ##### 访问测试
>
> 1、启动 keepalived 之前，先使用命令 `ip a`，查看 `192.168.200.133` 和 `192.168.200.122` 这两台服务器的 IP 情况
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.4lh36tkt0kk0.webp)
>
> 2、分别启动两台服务器的 keepalived
>
> ```sh
> cd /usr/local/sbin
> 
> ./keepalived
> ```
>
> 再次通过 `ip a` 查看 IP
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.79pu8lyco700.webp)
>
> 此时发现服务器 A 多出了 `192.168.200.222`，正是配置的虚拟路由 VIP，而服务器 B 并没有，说明服务器 A 是 Master，优先级高于服务器 B。
>
> 3、当把 `192.168.200.133` 服务器 A 上的 keepalived 关闭后，再次查看 IP
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.qwb7iteldn4.webp)
>
> 说明当 Master 服务器 A 宕机后，服务器 B 由 Backup 晋升为 Master。
>
> 通过上述的测试，我们会发现，虚拟 IP(VIP)会在 Master 节点上，当 Master 节点上的 keepalived 出问题以后，因为 Backup 无法收到 Master 发出的 VRRP 状态通过信息，就会直接升为 Master 。VIP 也会「漂移」到新的 Master 。
>
> **上面测试和 Nginx 有什么关系?**
>
> 我们把 `192.168.200.133` 服务器 A 的 keepalived 再次启动下，由于它的优先级高于 `192.168.200.122` 服务器 B，所有它会再次成为 Master，VIP 也会「漂移」过去。
>
> 我们通过浏览器访问：
>
> ```http
> http://192.168.200.222/
> ```
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.6xosmovingw0.webp)
>
> 如果把 `192.168.200.133` 服务器 A 的 keepalived 进程关闭掉
>
> ```sh
> kill keepalived 
> ```
>
> 再次访问相同的地址，效果如图：
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.60wzzgrveak0.webp)
>
> 虽然效果成功实现了，但是此时是我们手动把服务器上的 keepalived 关闭，才让 VIP 进行切换。

**而什么时候关闭 keepalived 呢？**

应该是在 keepalived 所在服务器的 Nginx 出现问题后，把 keepalived 关闭掉，就可以让 VIP 执行另外一台服务器。

**但是现在这所有的操作都是通过手动来完成的**，我们如何能让系统自动判断当前服务器的 Nginx 是否正确启动，如果没有，要能让 VIP 自动进行「漂移」，这个问题该如何解决？往下看。

##### vrrp_script

keepalived 只能做到对网络故障和 keepalived 本身的监控，即当出现网络故障或者 keepalived 本身出现问题时，进行切换。但是这些还不够，我们还需要监控 keepalived 所在服务器上的其他业务，比如 Nginx，如果 Nginx 出现异常了，而 keepalived 却保持正常，是无法完成系统的正常工作的，因此需要根据业务进程的运行状态决定是否需要进行主备切换，这个时候，我们可以通过编写脚本对业务进程进行检测监控。

首先我们要知道 keepalived 的 vrrp_script 的配置模板：

```nginx
vrrp_script 脚本名称
{
    script "脚本位置"
    interval 3 # 执行时间间隔
    weight -20 # 动态调整 vrrp_instance 的优先级
}
```

实现步骤：

- 编写脚本，这里的脚本名是 `ck_nginx.sh`，位置在 `/etc/keepalived` 路径下

```shell
#!/bin/bash
num=`ps -C nginx --no-header | wc -l`  # 查询 Nginx 的进程数

if [ $num -eq 0 ];then       	       # 如果 Nginx 的进程数等于 0
/usr/local/nginx/sbin/nginx			   # 则可执行文件 nginx，启动 Nginx 服务

sleep 2								   # 阻塞 2 秒 

if [ `ps -C nginx --no-header | wc -l` -eq 0 ]; then  # 再次查询 Nginx 的进程数
killall keepalived		# 如果 Nginx 的进程数不等于 0，则杀死 keepalived 进程

fi
fi
```

- ps 命令用于显示当前进程 (process) 的状态。
- -C(command)：指定命令的所有进程
- --no-header：排除标题

命令的效果如图：

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.1ypeqh0nut8g.webp)

代表目前的 num = 3。

这个脚本其实就是判断 Nginx 是否启动还是宕机了，如果没有启动，则重新启动。重新启动后再次查看 Nginx 是否启动成功，如果没有启动，说明 Nginx 宕机了，则杀死 keepalived 进程，这样，另一台服务器的 Nginx 就晋升为 Master。

- 为脚本文件设置权限

```sh
chmod 755 ck_ngi
```

- 将脚本添加到 Master 服务器 A 的 keepalived 的配置文件里

```sh
vim /etc/keepalived/keepalived.conf
```

- 添加如下内容：

```nginx
vrrp_script ck_nginx {
   script "/etc/keepalived/ck_nginx.sh" # 执行脚本的位置
   interval 2		# 执行脚本的周期，秒为单位
   weight -20		# 权重的计算方式
}
vrrp_instance VI_1 {
    state MASTER
    interface ens33
    virtual_router_id 10
    priority 100
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1111
    }
    virtual_ipaddress {
        192.168.200.111
    }
    track_script {
      ck_nginx
    }
}
```

- 如果效果没有出来，可以使用 `tail -f /var/log/messages` 查看日志信息，找对应的错误信息
- 两个 Nginx 启动后，关闭 Master 的 Nginx，通过 `ip a` 查看 Backup 的 Nginx 的IP，是否晋升为 Master

**问题思考**

通常如果 Master 服务死掉后，Backup 会变成 Master，但是当原来的 Master 服务又恢复了，它会和原来的 Backup 会抢占 VIP，这样就会发生两次切换，这对业务繁忙的网站来说是不好的。所以我们要在配置文件加入 nopreempt 非抢占，但是这个参数只能用于 Backup 的服务器，所以我们在用配置的时候，最好 Master 和 Backup 的 state 都设置成 Backup，这样它们只能通过 priority 优先级来竞争。





## Nginx站点和认证

### Nginx制作下载站点

首先我们先要清楚什么是下载站点？

我们先来看一个网站 [http://nginx.org/download/(opens new window)](http://nginx.org/download/)

这个是大家去下载 Nginx 时经常访问的网站，该网站主要就是用来提供用户来下载相关资源的网站，就叫做下载网站。

![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.17v6jxd79kao.webp)

如何制作一个下载站点：

- Nginx 使用的是模块 `ngx_http_autoindex_module` 来实现的（自带），该模块处理以斜杠(『 / 』)结尾的请求，并生成目录列表。
- Nginx 编译的时候会自动加载该模块，但是该模块默认是关闭的，我们需要使用下来指令来完成对应的配置



- ##### autoindex

`autoindex` 指令启用或禁用目录列表的输出

| 语法                   | 默认值         | 位置                   |
| ---------------------- | -------------- | ---------------------- |
| autoindex <on \| off>; | autoindex off; | http、server、location |



- ##### autoindex_exact_size

`autoindex_exact_size` 指令对应 HTLM 格式，指定是否在目录列表展示文件的详细大小。

默认为 on，显示出文件的确切大小，单位是 bytes。

改为 off 后，显示出文件的大概大小，单位是 kB 或者 MB 或者 GB。

| 语法                              | 默认值                   | 位置                   |
| --------------------------------- | ------------------------ | ---------------------- |
| autoindex_exact_size <on \| off>; | autoindex_exact_size on; | http、server、location |



- ##### autoindex_format

`autoindex_format` 指令设置目录列表的格式。

| 语法                                             | 默认值                 | 位置                   |
| ------------------------------------------------ | ---------------------- | ---------------------- |
| autoindex_format <html \| xml \| json \| jsonp>; | autoindex_format html; | http、server、location |

只有当 `autoindex_format` 指令设置为 html 时候，上方的 `autoindex_exact_size` 指令才会起作用。

注意

该指令在 1.7.9 及以后版本中出现。

- ##### autoindex_localtime

`autoindex_localtime` 指令对应 HTML 格式，是否在目录列表上显示时间。

默认为 off，显示的文件时间为 GMT 时间。

改为 on 后，显示的文件时间为文件的服务器时间。

| 语法                            | 默认值                   | 位置                   |
| ------------------------------- | ------------------------ | ---------------------- |
| autoindex_localtime <on \|off>; | autoindex_localtime off; | http、server、location |

> ### 实现案例
>
> 自行准备几个文件或者压缩包，我这里准备了 4 个用过的压缩包。
>
> - 创建一个目录，将压缩包放入其中，我这里创建的路径是 `/opt/download`
>
> ```sh
> mkdir /opt/download
> ```
>
> - 然后把压缩包都放到该目录下
>
> 效果如下:
>
> ```sh
> [root@master download]# pwd
> /opt/download
> [root@master download]# ll
> 总用量 545504
> -rw-r--r--. 1 root root 408587111 3月  12 22:02 hadoop-2.10.1.tar.gz
> -rw-r--r--. 1 root root 145520298 3月  12 21:29 jdk-8u301-linux-x64.tar.gz
> -rw-r--r--  1 root root     25680 4月  27 2017 mysql57-community-release-el7-11.noarch.rpm
> -rw-r--r--  1 root root   4456335 5月   9 19:40 mysql-connector-java-5.1.48.tar.gz
> ```
>
> - 打开 Nginx 的配置文件
>
> ```sh
> vim /usr/local/nginx/conf/nginx.conf
> ```
>
> - 添加配置如下内容：
>
> :::: tabs cache-lifetime="5" :options="{ useUrlFragment: false }"
>
> ::: tab 有注释版
>
> ```nginx
> location /download {
>     root /opt;                # 下载目录所在的路径，location 后面的 /download 拼接到 /opt 后面
>     # 以这些后缀的文件点击后为下载，注释掉则 txt 等文件是在网页打开并查看内容
>     if ($request_filename ~* ^.*?\.(txt|doc|pdf|rar|gz|zip|docx|exe|xlsx|ppt|pptx|conf)$){
> 			  add_header Content-Disposition 'attachment;';
> 		  }
>     autoindex on;			  # 启用目录列表的输出
>     autoindex_exact_size on;  # 在目录列表展示文件的详细大小
>     autoindex_format html;	  # 设置目录列表的格式为 html
>     autoindex_localtime on;   # 目录列表上显示系统时间
> }
> ```
>
> 
>
> root 指令后面必须是下载路径，因为我的下载路径是 `/opt/download`，所以这里填写 `/opt`，而 location 的 `/download` 会自动拼接到后面，形成完整的下载路径。
>
> - 访问 `192.168.91.200/download`
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.3gmq4emy84w0.webp)
>
> - JSON和XML格式(一般不用这两种格式)
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.6n5evv8oo5k0.webp)
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.kvsekaaiu40.webp)

### Nginx用户认证模块

对应系统资源的访问，我们往往需要限制谁能访问，谁不能访问。这块就是我们通常所说的认证部分，认证需要做的就是根据用户输入的用户名和密码来判定用户是否为合法用户，如果是则放行访问，如果不是则拒绝访问。

Nginx 对应用户认证这块是通过 `ngx_http_auth_basic_module` 模块来实现的，它允许通过使用「HTTP基本身份验证」协议验证用户名和密码来限制对资源的访问。默认情况下 Nginx 是已经安装了该模块，如果不需要则使用 `--without-http_auth_basic_module` 删除认证模块。

该模块的指令比较简单。

- ##### auth_basic

`auth_basic` 指令使用「HTTP基本身份验证」协议启用用户名和密码的验证。默认关闭。

| 语法                        | 默认值          | 位置                                 |
| --------------------------- | --------------- | ------------------------------------ |
| auth_basic <string \| off>; | auth_basic off; | http、server、location、limit_except |

开启后，服务端会返回 401，指定的字符串会返回到客户端，给用户以提示信息，但是不同的浏览器对内容的展示不一致。

- ##### auth_basic_user_file

`auth_basic_user_file` 指令指定用户名和密码所在文件，包括所在的路径。

| 语法                   | 默认值 | 位置                                 |
| ---------------------- | ------ | ------------------------------------ |
| auth_basic_user_file ; | —      | http、server、location、limit_except |

指定文件路径，该文件中设置用户名和密码，密码需要进行加密。可以采用工具自动生成。

> ### 实现案例
>
> - 在配置文件 nginx.conf 添加如下内容：
>
> ```nginx
> location /download{
>     # 下载站点知识
>     root /opt;                # 下载目录所在的路径，location 后面的 /download 拼接到 /opt 后面
>     autoindex on;			  # 启用目录列表的输出
>     autoindex_exact_size on;  # 在目录列表展示文件的详细大小
>     autoindex_format html;	  # 设置目录列表的格式为 html
>     autoindex_localtime on;   # 目录列表上显示系统时间
> 
>     # 认证模块知识
>     auth_basic 'please input your auth';   # 启用户名和密码的验证，并在请求头插入数据
>     auth_basic_user_file htpasswd;    # 存用户名和密码的文件路径
> }
> ```
>
> - 我们需要使用 `htpasswd` 工具生成包含用户名和密码的文件
>
> ```sh
> yum install -y httpd-tools
> ```
>
> 该工具基本操作指令如下：
>
> ```sh
> htpasswd -c /usr/local/nginx/conf/htpasswd username   # 创建一个新文件记录用户名和密码，密码后面弹出输入
> htpasswd -b /usr/local/nginx/conf/htpasswd username password   # 在指定文件新增一个用户名和密码
> htpasswd -D /usr/local/nginx/conf/htpasswd username   # 从指定文件删除一个用户信息
> htpasswd -v /usr/local/nginx/conf/htpasswd username   # 验证用户名和密码是否正确
> ```
>
> 根据需求指定生成路径的位置。
>
> - 我们创建一个 frx 的用户名，密码是 123456
>
> ![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.5uqgyc41uuc0.webp)
>
> 可以查看生成的文件内容
>
> ```sh
> [root@master conf]# cat /usr/local/nginx/conf/htpasswd
> frx:$apr1$wDXdSh0O$yR66Agylnta9zupO7cD3k.
> ```
>
> - 浏览器访问 `192.168.91.200/download`![image](https://jsdelivr.codeqihan.com//gh/xustudyxu/image-hosting1@master/20220806/image.1j6bwl4jr2g0.webp)
>
> 上述方式虽然能实现用户名和密码的验证，但是大家也看到了，所有的用户名和密码信息都记录在文件里面，如果用户量过大的话，这种方式就显得有点麻烦了，这时候我们就得通过后台业务代码来进行用户权限的校验了。