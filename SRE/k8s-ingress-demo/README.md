# 域名 → Ingress Controller → Service → Pod

这套文件演示同一域名的路径分流：

| 请求 | Kubernetes 后端 | Pod 实际监听 |
| --- | --- | --- |
| `http://demo.example.com/` | `web-svc:80` | `web-*` Pod 的 `8080` |
| `http://demo.example.com/api/users` | `api-svc:80` | `api-*` Pod 的 `8080` |

`demo.example.com` 和文中的 `203.0.113.10` 均为示例，使用时替换。whoami 是回显请求的演示程序；这里的 api 并未实现真正的用户查询业务。

## 1. 先区分配置关系和实际流量

逻辑关系：

```text
Ingress 规则
  └── backend 引用 Service 名称和 Service port
        └── Service selector 匹配 Pod labels
              └── EndpointSlice 记录 Pod IP、解析后的 targetPort、就绪状态
```

Ingress 是规则对象，Ingress Controller 才是读取规则并配置代理/负载均衡器的组件。本例中 Traefik 同时承担控制器和请求代理的职责。有些云控制器则主要负责配置外部负载均衡器，流量不经过控制器 Pod。

EndpointSlice 是 Kubernetes 维护的后端地址清单，不是额外的一台代理；它也不会主动转发请求。本例 Service 有 selector，因此不用手写 EndpointSlice。旧 Endpoints API 从 Kubernetes 1.33 起弃用，观察后端优先用 EndpointSlice。

本例 Traefik 默认直接选择 Pod：

```text
浏览器先向 DNS 查询 demo.example.com 的入口地址

请求流量：
公网入口:80
  → Traefik Controller（由其 LoadBalancer/NodePort Service 暴露）
    → 根据 Host + Path 匹配 Ingress
      → 根据后端 Service 及其 EndpointSlice 得到可用 Pod 地址
        → 某个 PodIP:8080
```

所以，Service 决定逻辑后端，并不意味着所有 Ingress 实现都会连接业务 Service 的 ClusterIP。Traefik 默认直接连接 Pod；若在业务 Service 的 metadata.annotations 中设置 `traefik.ingress.kubernetes.io/service.nativelb: "true"`，则会走 `Traefik → Service ClusterIP:80 → PodIP:8080`，由 Kubernetes 的 Service 数据面选择后端。

## 2. 前提和文件

- 已有可用 Kubernetes 集群，CNI 能让 Controller 访问业务 Pod，相关 NetworkPolicy 允许该流量。
- 使用 kubectl 和 Helm。示例固定 Traefik Chart `42.1.3`，该 Chart 声明要求 Kubernetes `>=1.30`。
- 云上路径假设平台已支持供应公网 LoadBalancer；不同平台可能需要专用 annotations、loadBalancerClass、子网和权限配置。
- 自建集群可选 NodePort 路径；NodePort 本身不会分配公网 IP。
- 这是一套 HTTP 路由学习配置；HTTPS 可按第 7 节切换。

| 文件 | 用途 |
| --- | --- |
| `controller-values.yaml` | Helm values，配置 Controller、IngressClass 和入口 Service |
| `apps.yaml` | Namespace、两个 Deployment、两个 ClusterIP Service |
| `ingress.yaml` | Host + Path 规则 |
| `optional/controller-nodeport-values.yaml` | 自建集群的 Helm 入口覆盖配置 |
| `optional/ingress-https.yaml` | 同一 Ingress 的 HTTPS 替换配置 |

只对 Kubernetes 资源文件执行 kubectl apply；values 文件交给 Helm。不要对整个目录直接执行 `kubectl apply -f .`。

## 3. 云上 LoadBalancer 入口

在本目录执行以下命令。安装会由 Helm 创建 Controller Deployment、Service、ServiceAccount、RBAC 和 IngressClass 等资源：

```bash
helm repo add traefik https://traefik.github.io/charts
helm repo update

helm upgrade --install ingress-demo traefik/traefik \
  --version 42.1.3 \
  --namespace ingress-demo-system --create-namespace \
  -f controller-values.yaml \
  --wait --timeout 5m

kubectl apply -f apps.yaml
kubectl apply -f ingress.yaml

kubectl -n ingress-demo rollout status deployment/web
kubectl -n ingress-demo rollout status deployment/api
kubectl -n ingress-demo-system get service ingress-demo-traefik
```

查看 Controller Service 的 EXTERNAL-IP，可能得到公网 IP，也可能得到云 LB 的域名。只有平台支持并配置了外部 LB，才会供应实际入口；裸集群创建 LoadBalancer Service 后通常会一直 pending。Ingress 对象并不会自行生成公网入口或配置公网 DNS。

在域名服务商处配置：

```text
若入口是公网 IP：
demo.example.com    A        <LB 公网 IPv4>

若入口是 LB 域名：
demo.example.com    CNAME    <云 LB 域名>
```

若使用 IPv6，则配置相应 AAAA 记录并确保 IPv6 入口可达。DNS 只负责将域名解析到入口，不按 `/api` 等路径分流。修改 `Ingress.spec.rules.host` 是请求匹配条件，不会自动创建 DNS 记录。

开放入口的 80；使用 HTTPS 时还需开放 443。LB 到节点或 Pod 的后端端口、健康检查也须按云平台要求放行。公网地址不要填写业务 Service 的 ClusterIP 或业务 Pod IP。

如需查看 Helm 实际生成的完整原生资源 YAML，可以在安装前渲染，不会修改集群：

```bash
helm template ingress-demo traefik/traefik \
  --version 42.1.3 \
  --namespace ingress-demo-system \
  -f controller-values.yaml > rendered-controller.yaml
```

## 4. 自建集群 NodePort 入口

采用此方案时，使用下面的安装命令替换第 3 节的 Helm 安装命令，业务 apps.yaml 和 ingress.yaml 相同：

```bash
helm upgrade --install ingress-demo traefik/traefik \
  --version 42.1.3 \
  --namespace ingress-demo-system --create-namespace \
  -f controller-values.yaml \
  -f optional/controller-nodeport-values.yaml \
  --wait --timeout 5m
```

入口 Service 将使用 NodePort 30080 和 30443。节点地址与 NodePort 必须能从入口侧访问；本例不启用 externalTrafficPolicy: Local，因此不要求被访问的节点本地运行 Controller Pod。

仅验证时，可将域名解析到一台公网可达的节点，访问 `http://demo.example.com:30080/api/users`。要用浏览器默认的 80/443 端口，则在公网 LB、网关或路由器配置四层转发/NAT：

```text
公网入口:80  → 可达的 NodeIP:30080
公网入口:443 → 可达的 NodeIP:30443
```

DNS 此时指向该公网入口。若没有公网 IP、入站映射或其他公网入口，仅修改 DNS 无法让外网进入内网。普通 A/CNAME 记录不携带 NodePort 端口号。多节点生产入口还需要健康检查和入口高可用。

## 5. 验证转发逻辑

DNS 配好后：

```bash
curl http://demo.example.com/
curl http://demo.example.com/api/users
```

DNS 尚未配置、但公网入口已可达时，假设实际入口 IP 是 203.0.113.10，可临时指定解析：

```bash
curl --resolve demo.example.com:80:203.0.113.10 \
  http://demo.example.com/api/users
```

若入口是云 LB 域名，也可保持请求 Host 来验证 HTTP：

```bash
curl -H 'Host: demo.example.com' http://YOUR-LB-HOSTNAME/api/users
```

尚无公网入口时，可从本机验证 Controller 到 Pod 的链路。保持下面命令在一个终端运行，在另一个终端发请求：

```bash
kubectl -n ingress-demo-system port-forward service/ingress-demo-traefik 18080:80
```

```bash
curl -H 'Host: demo.example.com' http://127.0.0.1:18080/
curl -H 'Host: demo.example.com' http://127.0.0.1:18080/api/users
```

port-forward 只验证被转发到的 Controller 和其后端链路，不验证公网 LB、DNS 或其他 Controller 副本。

响应中可以看到 `Hostname: api-...`、`GET /api/users HTTP/1.1` 等信息；访问 `/` 时 Hostname 应以 `web-` 开头。连续发起独立请求可能看到不同 Pod 名称，但不要把严格轮询当成正确性的条件。

查看业务 Pod、Service 和自动生成的后端地址：

```bash
kubectl -n ingress-demo get pods -o wide --show-labels
kubectl -n ingress-demo get services
kubectl -n ingress-demo get endpointslices \
  -l kubernetes.io/service-name=api-svc -o yaml
kubectl -n ingress-demo describe ingress demo
kubectl get ingressclass traefik-demo
kubectl -n ingress-demo-system logs deployment/ingress-demo-traefik --tail=100
```

例如访问 `http://demo.example.com/api/users`：

1. DNS 返回公网入口地址。
2. 请求经入口 LB 或 NodePort 到达 Traefik 的 HTTP 监听端口 8000。
3. Traefik 按 Ingress 的 host 匹配 `demo.example.com`。
4. `/api/users` 同时匹配 `/` 和 `/api`，最长匹配路径 `/api` 优先。
5. Ingress 的 backend 引用同命名空间的 `api-svc`，端口 80。
6. Service 的 selector `app: api` 匹配 api Deployment 创建的 Pod 标签。
7. `targetPort: http` 指向 Pod 中名为 http 的端口，即 8080。应用由 `--port=8080` 真正监听；containerPort 声明本身不会启动监听。
8. EndpointSlice 记录这些 Pod 的 IP、端口 8080 和就绪状态。默认模式下 Traefik 从可用地址中选择一个，直接连接 PodIP:8080。
9. Pod 将响应经 Traefik 和入口返回客户端。

Ingress backend 的 80 对应 Service port，不对应 Controller 的 8000，也不直接对应应用的 8080。

本例不做路径改写，后端收到的路径仍是 `/api/users`。如果应用只实现了 `/users`，应调整应用路由或显式配置 Controller 的路径改写功能。`Prefix` 不是自动去前缀。

严格路径段匹配由 controller-values.yaml 中的 strictprefixmatching 参数启用：`/api`、`/api/`、`/api/users` 走 api，`/apix` 走 `/` 的 web 后端。

## 6. 快速排查

| 现象 | 优先检查 |
| --- | --- |
| 域名无法解析 | 公网 DNS、记录值、TTL |
| 连接超时或拒绝 | 公网入口、监听端口、安全组、防火墙、NAT、LB 健康检查 |
| Controller 返回 404 | Host、Path、ingressClassName、Controller 是否加载规则 |
| 后端返回 404 | 应用实际实现的路径，是否误以为 Prefix 会删前缀 |
| 502/503 | Service selector、EndpointSlice 就绪地址、targetPort、Pod 监听、网络策略；具体状态码取决于 Controller 配置和故障原因 |
| EXTERNAL-IP pending | 是否有可用 LB 实现，云权限、子网与 LB 事件 |

## 7. 可选 HTTPS

准备覆盖实际域名的可信证书和私钥；将 optional/ingress-https.yaml 里的域名改成你的域名。Secret 必须与 Ingress 同命名空间：

```bash
kubectl -n ingress-demo create secret tls demo-tls \
  --cert=/absolute/path/fullchain.pem \
  --key=/absolute/path/privkey.pem \
  --dry-run=client -o yaml | kubectl apply -f -

kubectl apply -f optional/ingress-https.yaml
curl https://demo.example.com/api/users
```

此文件替换同名 Ingress，使其只接受 websecure 入口的 HTTPS 请求；未配置 HTTP 自动跳转。证书在 Controller 终止，Controller 到演示 Pod 仍使用 HTTP。需要公网 443 到 Controller 的网络链路可达。

## 官方资料

- [Kubernetes Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)
- [Kubernetes Service / EndpointSlice](https://kubernetes.io/docs/concepts/services-networking/service/)
- [Traefik Kubernetes Ingress provider](https://doc.traefik.io/traefik/reference/install-configuration/providers/kubernetes/kubernetes-ingress/)
- [Traefik Ingress 路由与 Service 原生负载均衡](https://doc.traefik.io/traefik/reference/routing-configuration/kubernetes/ingress/)
- [Traefik Chart 42.1.3](https://github.com/traefik/traefik-helm-chart/releases/tag/v42.1.3)
- [whoami v1.11.0](https://github.com/traefik/whoami/tree/v1.11.0)
- [社区 ingress-nginx 退役说明](https://kubernetes.io/blog/2026/01/29/ingress-nginx-statement/)

Ingress API 与社区 ingress-nginx Controller 是不同的东西。社区 ingress-nginx 已于 2026 年 3 月退役，停止上游维护；本例使用 Traefik 实现标准 Ingress。Kubernetes 建议新功能和新架构考虑 Gateway API；这不影响用 Ingress 学习上述路由关系。

文件已做本地 YAML 解析和名称、selector、端口关联检查；当前环境没有 kubectl/Helm，未在实际集群执行部署或端到端验证。
