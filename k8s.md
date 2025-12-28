# k8s

ssh://r***@192.168.66.

## Pod 控制器

几种常用的pod部署器及其处理

- ReplicaSet
- Deployment
- DaemonSet
- Statefulset
- Job/Crobjob
- Horizontal Pod Autoscaler

### ReplicationController 和 ReplicaSet

#### ReplicationController

ReplicationController（RC）用来**确保容器应用的副本数始终保持在用户定义的副本数**，即如果有容器异常退出，会自动创建新的Pod来替代；而如果异常多出来的容器也会自动回收；

在新版本的 Kubernetes 中建议使用 ReplicaSet 来取代 ReplicationController **。ReplicaSet 跟 ReplicationController 没有本质的不同，只是名字不一样，并且 ReplicaSet 支持集合式的 selector;** 

> - pod 的 重启策略 多的是指当 pod内部发生错误时，自己启动一个新的pod 这会体现在pod 信息的restart 
>
> - 确保容器应用的副本数始终保持在用户定义的副本数
>
>   指的是这个pod直接不存在了(删除） ，那么RC在检测到 replica数量差异时会重新生成新的pod 
>
>   会导致这里的pod名称改变 因调度规则发生的部署node改变的情况

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: rc-demo
spec:
  replicas: 3        # 期望的 Pod 数量（推荐英文注释：Desired number of pods）
  selector:          # Pod 选择器（推荐英文注释：Selector for matching pods）
    app: rc-demo	 # 判定是本资源的选择器
  template:          # Pod 模板（推荐英文注释：Pod template definition）
    metadata:
      labels:
        app: rc-demo
    spec:
      containers:
      - name: rc-demo-container
        image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
        env:          # 环境变量（推荐英文注释：Environment variables）
        - name: GET_HOSTS_FROM
          value: dns
        - name: zhangsan
          value: "123"
        ports:
        - containerPort: 80  # 容器暴露的端口（推荐英文注释：Container exposed port）
```

> 
>
> <img src=".\asset\images\image-20250420120956764.png" alt="image-20250420120956764" style="zoom:67%;" />
>
> - 归属情况测试
>
>   ```bash
>   [root@k8s-master01 1-ReplicationController]# kubectl get pod --show-labels
>   NAME            READY   STATUS    RESTARTS   AGE   LABELS
>   rc-demo-brnlp   1/1     Running   0          17m   app=rc-demo
>   rc-demo-lwcw5   1/1     Running   0          16m   app=rc-demo
>   rc-demo-z7jhx   1/1     Running   0          32m   app=rc-demo
>                               
>   ## 对某个pod添加label，仍是 RC selector 中  app: rc-demo 的子集 不会导致pod的重新调度
>   [root@k8s-master01 1-ReplicationController]# kubectl label pod rc-demo-brnlp version=1 
>   pod/rc-demo-brnlp labeled
>   [root@k8s-master01 1-ReplicationController]# kubectl get pod --show-labels
>   NAME            READY   STATUS    RESTARTS   AGE   LABELS
>   rc-demo-brnlp   1/1     Running   0          18m   app=rc-demo,version=1
>   rc-demo-lwcw5   1/1     Running   0          16m   app=rc-demo
>   rc-demo-z7jhx   1/1     Running   0          32m   app=rc-demo
>                               
>   ## 修改 app标签后，rc 不认为这是其管理的pod 因此需要重新调度pod 创建了rc-demo-b469r
>   [root@k8s-master01 1-ReplicationController]# kubectl label pod rc-demo-lwcw5  app=test --overwrite
>   pod/rc-demo-lwcw5 labeled
>   [root@k8s-master01 1-ReplicationController]# kubectl get pod --show-labels
>   NAME            READY   STATUS    RESTARTS   AGE   LABELS
>   rc-demo-b469r   1/1     Running   0          2s    app=rc-demo
>   rc-demo-brnlp   1/1     Running   0          19m   app=rc-demo,version=1
>   rc-demo-lwcw5   1/1     Running   0          17m   app=test
>   rc-demo-z7jhx   1/1     Running   0          33m   app=rc-demo
>                               
>   ## 当修改回pod label 时，RC 默认存在时间久的pod携带更多信息，因此delete上一次新创建的pod
>   [root@k8s-master01 1-ReplicationController]# kubectl label pod rc-demo-lwcw5  app=rc-demo --overwrite
>   pod/rc-demo-lwcw5 labeled
>   [root@k8s-master01 1-ReplicationController]# kubectl get pod --show-labels
>   NAME            READY   STATUS    RESTARTS   AGE   LABELS
>   rc-demo-brnlp   1/1     Running   0          20m   app=rc-demo,version=1
>   rc-demo-lwcw5   1/1     Running   0          19m   app=rc-demo
>   rc-demo-z7jhx   1/1     Running   0          35m   app=rc-demo
>   ```
>
>   
>
> - 调整副本数量
>
>   ```shell
>   [root@k8s-master01 1-ReplicationController]# kubectl scale rc rc-demo --replicas=10
>   replicationcontroller/rc-demo scaled
>   [root@k8s-master01 1-ReplicationController]# kubectl get pod --show-labels
>   NAME            READY   STATUS    RESTARTS   AGE   LABELS
>   rc-demo-5n5gt   1/1     Running   0          10s   app=rc-demo
>   rc-demo-8cdbf   1/1     Running   0          10s   app=rc-demo
>   rc-demo-brnlp   1/1     Running   0          28m   app=rc-demo,version=1
>   rc-demo-c5lhs   1/1     Running   0          10s   app=rc-demo
>   rc-demo-c8925   1/1     Running   0          10s   app=rc-demo
>   rc-demo-lqn7l   1/1     Running   0          10s   app=rc-demo
>   rc-demo-lwcw5   1/1     Running   0          27m   app=rc-demo
>   rc-demo-xngb6   1/1     Running   0          10s   app=rc-demo
>   rc-demo-xnhbk   1/1     Running   0          10s   app=rc-demo
>   rc-demo-z7jhx   1/1     Running   0          43m   app=rc-demo
>   ```
>
>   

#### ReplicaSet

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: rs-ml-demo
spec:
  replicas: 3
  selector:
    matchLabels:
      app: rs-ml-demo
  template:
    metadata:
      labels:
        app: rs-ml-demo
        domain: testrs
        version: v1
    spec:
      containers:
      - name: rs-ml-demo-container
        image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
        env:
        - name: GET_HOSTS_FROM
          value: dns
        ports:
        - containerPort: 80
```

> - 创建情况
>
>   ```bash
>   [root@k8s-master01 1-ReplicationController]# kubectl get rs
>   NAME         DESIRED   CURRENT   READY   AGE
>   rs-ml-demo   3         3         3       44s
>   [root@k8s-master01 1-ReplicationController]# kubectl get pods -o wide
>   NAME               READY   STATUS    RESTARTS   AGE   IP              NODE         NOMINATED NODE   READINESS GATES
>   rc-demo-brnlp      1/1     Running   0          47m   10.244.58.198   k8s-node02   <none>           <none>
>   rc-demo-lwcw5      1/1     Running   0          46m   10.244.58.199   k8s-node02   <none>           <none>
>   rc-demo-z7jhx      1/1     Running   0          61m   10.244.58.196   k8s-node02   <none>           <none>
>   rs-ml-demo-2rdgk   1/1     Running   0          75s   10.244.85.208   k8s-node01   <none>           <none>
>   rs-ml-demo-bf79n   1/1     Running   0          75s   10.244.85.207   k8s-node01   <none>           <none>
>   rs-ml-demo-bzzgs   1/1     Running   0          75s   10.244.58.202   k8s-node02   <none>           <none>
>                               
>   ```
>
>   

> - **selector.matchExpressions**
>
>   rs 在标签选择器上，除了可以定义键值对的选择形式，还支持 matchExpressions 字段，可以提供多种选择。
>
>   目前支持的操作包括：
>
>   - In：label 的值在某个列表中
>   - NotIn：label 的值不在某个列表中
>   - Exists：某个 label 存在
>   - DoesNotExist：某个 label 不存在
>
>   **这里的匹配逻辑有默认的父方检测，只有是我创建的 并且符合这些匹配规则的才会算作子集部署pod**，这样的话，能够排除那些仅标签规则符合但是实际上毫无关系的pod
>
>    ```yaml
>    apiVersion: apps/v1
>    kind: ReplicaSet
>    metadata:
>     name: rs-me-exists-demo
>    spec:
>     selector:
>       matchExpressions:
>         - key: app
>           operator: Exists
>     template:
>       metadata:
>         labels:
>           app: spring-k8s
>       spec:
>         containers:
>           - name: rs-me-exists-demo-container
>             image: wangyanglinux/myapp:v1.0
>             ports:
>               - containerPort: 80
>    ```
>
>   
>
>   ```yaml
>   apiVersion: apps/v1
>   kind: ReplicaSet
>   metadata:
>     name: rs-ml-demo
>   spec:
>     replicas: 3
>     selector:
>       matchLabels:		# 这里就是原先 rc 中 仅有的功能
>         app: rs-ml-demo
>       matchExpressions:	# rs中的标签选择表达式
>         - key: environment
>           operator: In
>           values:
>             - production
>             - staging
>         - key: tier
>           operator: Exists
>     template:
>       metadata:
>         labels:
>           app: rs-ml-demo
>           environment: production  # Example label that would match the selector
>           tier: frontend         # Example label that would match the selector
>       spec:
>         containers:
>         - name: rs-ml-demo-container
>           image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
>           env:
>           - name: GET_HOSTS_FROM
>             value: dns
>           ports:
>           - containerPort: 80
>   ```



### Deployment

Deployment 为Pod和 ReplicaSet 提供了一个声明式定义(declarative)方法，用来替代以前的ReplicationController 来方便的管理应用。

> <img src=".\asset\images\image-20250420184238676.png" alt="image-20250420184238676" style="zoom:67%;" />
>
> <img src=".\asset\images\image-20250420184355087.png" alt="image-20250420184355087" style="zoom:67%;" />
>
> <img src=".\asset\images\image-20250420184714259.png" alt="image-20250420184714259" style="zoom:67%;" />
>
> - 总的来说 replace 命令是完全覆盖 **对于缺省值也会直接按照新的覆盖**
>
> - 而 apply 仅会修改不同的部分，**缺省值不会做修改**
> - 通过 kubectl diff -f yaml 能够diff 已经部署和yaml文件(未部署) 区别的地方 
>
> > kubectl  create 、apply 、replace
> >
> > -  create 
> >
> >    创建资源对象 
> >
> >   -f  通过基于文件的创建**，但是如果此文件描述的对象存在，那么那怕文件描述的信息发生了改变，再次提交时也不会应用**
> >
> > -  apply
> >
> >    创建资源对象、修改资源对象
> >
> >   -f  基于文件创建，如果目标对象与文件本身发生改变，那么会根据文件的指定一一修改目标对象的属性（**部分更新**） 
> >
> > -  replace 
> >
> >    创建资源对象、修改资源对象
> >
> >     -f  基于文件创建，如果目标对象与文件本身发生改变，那么会重建此对象（替换）

典型的应用场景包括：

- 定义 Deployment 来创建Pod 和 ReplicaSet 
- 滚动升级和回滚应用
- 扩容和缩容
- 暂停和继续 Deployment 



```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: myapp-deploy    # 这是 deployment的标签
  name: myapp-deploy
spec:
  replicas: 1
  selector:
    matchLabels:		# 同理这里还有 matchExpressions 运算匹配表达式
      app: myapp-deploy
  template:
    metadata:
      labels:
        app: myapp-deploy
    spec:
      containers:
      - image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
        name: myapp
```

> deployment 运行原理
>
> deployment 实际上的功能 是通过掌控 ReplicaSet来完成的
>
> <img src=".\asset\images\image-20250420190959242.png" alt="image-20250420190959242" style="zoom:67%;" />
>
> **实际的优势在于，通过声明式的部署原理，能够在变更原有 rs的基础上 滚动更新，不会立即终止掉原有的服务**，保证用户实际体验
>
> 同时，如若现有的版本出现问题，原本的 rs 配置文件还会保存在 etcd中，方便 **版本回滚**
>
> <img src=".\asset\images\image-20250420191607347.png" alt="image-20250420191607347" style="zoom:67%;" />

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deployment-demo
  labels:
    app: deployment-demo
spec:
  replicas: 5
  selector:
    matchLabels:
      app: deployment-demo
  template:
    metadata:
      labels:
        app: deployment-demo
    spec:
      containers:
      - name: deployment-demo-container
        image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
```

> - `kubectl create -f deployment. yaml -record`
>
>   \#—record参数可以记录命令，我们可以很方便的查看每次revision的变化kubectl scale deployment nginx-deployment -replicas 10
>
> - `kubectl autoscale deployment nginx-deployment -min=10 -max=15 -cpu-percent=80`
>
>   设置最大最小值，能够根据负载情况 自动调整 
>
> - `kubectl set image deployment/nginx-deployment nginx-deployment-container-wangyanglinux/myapp:v2.0`
>
>   动态更新镜像
>
>   > deployment 更新策略
>   >
>   > ```yaml
>   > spec:
>   >   strategy:
>   >     type: RollingUpdate    #  RollingUpdate 逐渐替换原有的pod  RollingUpdate Recreate：会先删除所有现有的 Pods，然后创建新的 Pods，适用于无法并行工作的应用。
>   >     rollingUpdate: 
>   >       maxSurge: 1	# 定义了更新过程中，可以有多少个额外的 Pods 在新增时创建。可以使用百分比
>   >       maxUnavailable: 0 # 更新过程中，可以有多少个 Pods 在任意时刻不可用（被删除） 可以使用百分比
>   > ```
>   >
>   > 
>   >
>   > <img src=".\asset\images\image-20250420200402483.png" alt="image-20250420200402483" style="zoom:67%;" />
>   >
>   > ![image-20250420200657413](.\asset\images\image-20250420200657413.png)
>   >
>   > - **金丝雀部署**
>   >
>   >   就是滚动更新过程中，通过用户测试新版本的pod，如果出现问题，马上回滚
>   >
>   >   ```bash
>   >   kubectl patch deployment deployment-demo-p '{"spec": ("strategy":("rollingUpdate": ("maxSurge": 1, "maxUnavailable": 0))))'
>   >               
>   >   # kubectl rollout pause deploy deployment-demo 滚动暂停
>   >   # 创建一个后立即暂停滚动
>   >   kubectl patch deployment deployment-demo--patch '("spec": {"template": {"spec": {"containers": O[f"name": "deployment-demo-container","image":"wangyanglinux/myapp:v2.0"}' & kubectl rollout pause deploy deployment-demo
>   >               
>   >   # 恢复当前滚动更新
>   >   kubectl rollout resume deploy deployment-demo 
>   >   ```
>   >
>   > 
>
> - kubectl rollout undo deployment/nginx-deployment 
>
>   回滚至**上一个版本** 实际上deployment 就是通过设置控制 replicaSet 的 replica 实现
>
>   > - kubectl rollout status deployments nginx-deployment
>   >
>   >   执行变更命令后查看回滚状态
>   >
>   >   echo$?  返回上一条命令的执行结果
>   >
>   > - kubectl rollout history deployment/nginx-deployment
>   >
>   >   查看变更历史![image-20250420203550804](.\asset\images\image-20250420203550804.png)在执行变更命令时 加 --record 会记录命令
>   >
>   >   但是前一条加后一条不加的情况下，后一条的记录会复制上一条的record
>   >
>   > - kubectl rollout undo deployment/nginx-deployment --to-revision=2
>   >
>   >   指定回滚版本 这个版本是通过 上一条命令获取的
>   >
>   > - kubectl rollout pause deployment/nginx-deployment
>   >
>   >   暂停变更
>
> - 清理策略
>
>    您可以通过设置 `.spec.revisionHistoryLimit` 项来指定 Deployment 最多保留多少 revision 历史记录。默认会保留所有的 revision；如果将该项设置为 0，Deployment 就不允许回退了。

- 变更版本控制问题

  Q: 比如，我将运行镜像的版本从V1 --》 V2 但是在未变更完全的途中 又设置变更为 V3

  官方做法是将所有并存的转换为V3 而不是将V1的全部转换为V2后再转换为V3





### DaemonSet

**DaemonSet 核心功能**

- **确保**：在全部（或指定）Node 上运行一个 Pod 副本
- **动态扩展**：新 Node 加入集群时自动创建 Pod
- **自动回收**：Node 移除时删除对应 Pod
- **级联删除**：删除 DaemonSet 会清理其创建的所有 Pod

------

**DaemonSet 典型用法**

| 场景分类         | 具体示例                                                     |
| :--------------- | :----------------------------------------------------------- |
| **集群存储**     | `glusterd`、`ceph`（每个节点提供分布式存储服务）             |
| **日志收集**     | `fluentd`、`logstash`（采集节点日志并转发到中心系统）        |
| **监控指标**     | `Prometheus Node Exporter`、`collectd`（暴露节点级监控指标） |
| **商业监控代理** | Datadog 代理、New Relic 代理（与云监控平台集成）             |
| **集群状态**     | Ganglia `gmond`（收集节点性能数据）                          |

------

**技术说明**

1. **与 Deployment 区别**

   - DaemonSet 不关注 `replicas` 数量，而是保证每个匹配节点 1 个 Pod
   - 通过 `nodeSelector` 或 `affinity` 指定目标节点（默认全部节点）

2. **配置示例**

   ```yaml
   apiVersion: apps/v1
   kind: DaemonSet
   metadata:
     name: fluentd-logging
   spec:
     selector:
       matchLabels:
         name: fluentd
     template:
       metadata:
         labels:
           name: fluentd
       spec:
         containers:
         - name: fluentd
           image: fluentd:latest
   ```

3. **适用场景**

   - 需要**节点级后台服务**（如存储、网络、监控）
   - 服务与节点**强绑定**（如硬件指标采集）



```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: daemonset-demo
  labels:
    app: daemonset-demo
spec:
  selector:
    matchLabels:
      name: daemonset-demo
  template:
    metadata:
      labels:
        name: daemonset-demo
    spec:
      containers:
      - name: daemonset-demo-container
        image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
```



### Job

**Job 核心功能**

负责管理**批处理任务**，确保一个或多个 Pod 成功执行后终止(当前返回码为0)。

------

**特殊配置说明**

| 字段/属性                         | 说明                                                         |
| :-------------------------------- | :----------------------------------------------------------- |
| **`spec.template`**               | 格式与 Pod 定义完全一致                                      |
| **`restartPolicy`**               | 仅允许 `Never` 或 `OnFailure`（不可用 `Always`）             |
| **单 Pod 默认行为**               | 当单个 Pod 成功运行后，Job 自动完成                          |
| **`.spec.completions`**           | 需成功运行的 Pod 数量（默认值：1）                           |
| **`.spec.parallelism`**           | 允许并行运行的 Pod 数量（默认值：1）说人话就是一次性启动的pod数量 |
| **`.spec.activeDeadlineSeconds`** | Pod 重试的最大时间（秒），超时后 Job 标记为失败并停止重试    |

**使用场景**

- 数据处理任务（如ETL作业）
- 科学计算任务
- 一次性数据库迁移



```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: job-demo
spec:
  template:
    metadata:
      name: job-demo-pod
    spec:
      containers:
      - name: job-demo-container
        image: registry.cn-hangzhou.aliyuncs.com/wangyangshare/share:maqingpythonv1
      restartPolicy: Never  # Job必须设置为Never或OnFailure 这里pod的重启策略虽然是never 但是返回非0时 job会执行自己的重启策略
   
```

> ```bash
> [root@k8s-master01 4-Job]# kubectl get job
> NAME       COMPLETIONS   DURATION   AGE
> job-demo   1/1           3s         3m24s
> [root@k8s-master01 4-Job]# kubectl get pod
> NAME             READY   STATUS      RESTARTS   AGE
> job-demo-zhrrf   0/1     Completed   0          3m26s
> [root@k8s-master01 4-Job]# kubectl logs job-demo-zhrrf
> 3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491298336733624406566430860213949463952247371907021798609437027705392171762931767523846748184676694051320005681271452635608277857713427577896091736371787214684409012249534301465495853710507922796892589235420199561121290219608640344181598136297747713099605187072113499999983729780499510597317328160963185950244594553469083026425223082533446850352619311881710100031378387528865875332083814206171776691473035982534904287554687311595628638823537875937519577818577805321712268066130019278766111959092164201989
> Total time:0.0011260509491s
> ```



```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: rand
spec:
  template:
    spec:
      containers:
      - name: rand
        image: registry.cn-hangzhou.aliyuncs.com/wangyangshare/share:randexitv1
        args: ["--exitcode=1"]
      restartPolicy: Never
```

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: rand
spec:
  completions: 10		# 成功条目数 表示 Job 需要执行多少次才能完成。
  parallelism: 5		# 并行运行数 控制并发执行的 Pod 数量。 例如，设置 parallelism: 3 会同时启动 3 个 Pod 来并行处理任务。
  activeDeadlineSeconds: 300  # 超时时间(秒) 设定 Job 执行的最大时间（以秒为单位）。超过此时间后，如果任务还未完成，将被终止。
  backoffLimit: 3		#  # 最大重试次数		
  template:
    spec:
      containers:
      - name: rand
        image: wangyanglinux/tools:randexitv1
      restartPolicy: Never
      
```







### CronJob

**创建 Job 操作应该是幂等的**

**CronJob 核心功能**

管理基于时间调度的 Job，支持两种模式：

1. **单次运行** - 在指定时间点执行一次
2. **周期性运行** - 按计划重复执行（如每日备份）

------

**版本要求**

- Kubernetes 集群版本 ≥ 1.8

------

**典型应用场景**

| 场景类型     | 具体示例                       |
| :----------- | :----------------------------- |
| **定时任务** | 每天凌晨执行数据库备份         |
| **周期作业** | 每小时发送一次系统状态报告邮件 |
| **批处理**   | 每月1号执行账单生成任务        |

> **1. 核心字段**
>
> | 字段                    | 说明                                                         | 是否必填 | 默认值 |
> | :---------------------- | :----------------------------------------------------------- | :------- | :----- |
> | **`.spec.schedule`**    | Cron 表达式，定义任务执行周期（如 `"0 * * * *"` 分时日月周） | 必填     | 无     |
> | **`.spec.jobTemplate`** | Job 模板，定义具体执行的任务（格式同标准 Job）               | 必填     | 无     |
>
> ------
>
> **2. 高级调度控制**
>
> | 字段                                | 说明                                                         |
> | :---------------------------------- | :----------------------------------------------------------- |
> | **`.spec.startingDeadlineSeconds`** | 允许启动 Job 的最大延迟时间（秒），超时视为失败；不设置则无限制 |
> | **`.spec.concurrencyPolicy`**       | **并发策略：是否允许不同批次的job是否能够重叠运行**<br /> • `Allow`（默认）：允许并发执行 <br />• `Forbid`：禁止并发，跳过新任务 <br />• `Replace`：终止当前任务并启动新任务 |
> | **`.spec.suspend`**                 | 设置为 `true` 暂停后续调度（不影响已运行的 Job)<br />        |
>
> > <img src=".\asset\images\image-20250420223902268.png" alt="image-20250420223902268" style="zoom:67%;" />
>
> ------
>
> **3. 历史记录管理**
>
> | 字段                                   | 说明                                             |          |
> | :------------------------------------- | :----------------------------------------------- | -------- |
> | **`.spec.successfulJobsHistoryLimit`** | 保留成功 Job 的历史记录数量（设置为 `0` 不保留） | 默认 `3` |
> | **`.spec.failedJobsHistoryLimit`**     | 保留失败 Job 的历史记录数量（设置为 `0` 不保留） | 默认 `1` |



```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: cronjob-demo
spec:
  schedule: "*/1 * * * *"
  successfulJobsHistoryLimit: 3  # 保留3个成功记录
  failedJobsHistoryLimit: 1      # 保留1个失败记录  
  jobTemplate:			# job 模板 同时支持原有的 job 的参数设置
    spec:
      completions: 3
      template:			# pod 模板
        spec:
          containers:
          - name: cronjob-demo-container
            image: registry.cn-hangzhou.aliyuncs.com/wangyangshare/share:busybox
            args:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure
```

> ```bash
> [root@k8s-master01 chapter-5]# kubectl get job
> NAME                    COMPLETIONS   DURATION   AGE
> cronjob-demo-29067177   3/3           10s        2m10s
> cronjob-demo-29067178   3/3           10s        70s
> cronjob-demo-29067179   3/3           10s        10s
> 
> [root@k8s-master01 chapter-5]# kubectl get pod
> NAME                          READY   STATUS      RESTARTS   AGE
> cronjob-demo-29067177-f74fg   0/1     Completed   0          2m15s
> cronjob-demo-29067177-fz45f   0/1     Completed   0          2m12s
> cronjob-demo-29067177-wklgt   0/1     Completed   0          2m8s
> cronjob-demo-29067178-6bn2r   0/1     Completed   0          68s
> cronjob-demo-29067178-q4ht8   0/1     Completed   0          75s
> cronjob-demo-29067178-rjxhb   0/1     Completed   0          72s
> cronjob-demo-29067179-49sv6   0/1     Completed   0          12s
> cronjob-demo-29067179-lgvpt   0/1     Completed   0          8s
> cronjob-demo-29067179-lhw5c   0/1     Completed   0          15s
> 
> [root@k8s-master01 chapter-5]# kubectl get cronjob
> NAME           SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
> cronjob-demo   */1 * * * *   False     0        32s             2m55s
> 
> # 目前设置 successfulJobsHistoryLimit: 3  仅保留成功三次的 history
> [root@k8s-master01 chapter-5]# kubectl get job
> NAME                    COMPLETIONS   DURATION   AGE
> cronjob-demo-29067178   3/3           10s        2m29s
> cronjob-demo-29067179   3/3           10s        89s
> cronjob-demo-29067180   3/3           10s        29s
> ```
>
> - 







### 小结

#### 一、Pod 控制器

##### 1. 守护进程类型控制器

| 控制器类型     | 核心功能                                        | 特点                                                         |
| :------------- | :---------------------------------------------- | :----------------------------------------------------------- |
| **RC**         | 保障 Pod 数量始终与期望值一致                   | - 基础副本控制 - 仅支持等值标签选择器                        |
| **RS**         | 在 RC 基础上增强标签选择能力                    | - 支持集合式标签选择器（`matchExpressions`） - 通常被 Deployment 间接使用 |
| **Deployment** | 提供声明式更新和版本控制                        | - 滚动更新/回滚能力 - 通过控制 RS 实现多版本管理 - 典型工作流：Deployment → RS → Pod |
| **DaemonSet**  | 确保每个（或指定）节点运行且仅运行一个 Pod 副本 | - 节点级守护进程 - 新节点加入自动部署 - 典型应用：日志收集、节点监控 |

##### 2. 批处理任务类型控制器

| 控制器类型  | 核心功能                               | 特点                                                         |
| :---------- | :------------------------------------- | :----------------------------------------------------------- |
| **Job**     | 确保批处理任务成功完成（1个或多个Pod） | - 任务式运行 - 支持并行/串行执行 - 成功标准：Pod 退出码为0   |
| **CronJob** | 基于时间调度 Job 执行                  | - 周期/定时任务 - 继承 Job 所有功能 - 典型应用：数据库备份、定期报表生成 |

#### 二、核心原理图示

![deepseek_mermaid_20250420_ac0d3c (1)](C:\Users\xlwwwww\Downloads\deepseek_mermaid_20250420_ac0d3c (1).png)

#### 三、选择建议

1. **长期运行服务** → Deployment
2. **节点级系统服务** → DaemonSet
3. **一次性任务** → Job
4. **定时任务** → CronJob
5. **传统副本控制** → RS/RC (已逐渐被 Deployment 替代)



## Service

Kubernetes `Service` 定义了这样一种抽象：一个 `Pod` 的逻辑分组，一种可以访问它们的策略 —— 通常称为微服务。这一组 `Pod` 能够被 `Service` 访问到，通常是通过 `Label Selector`

> 传统的部署，单个物理机单单部署一个应用或者服务的话，整体的利用率太低了
>
> 而 k8s 能够通过pod的方式部署，并且 service 通过标签的方式抓取pod信息 并支持负载均衡
>
> - 之前曾经提到过，pod在运行之后，如果通过就绪探针验证，那么就会被service抓取到负载均衡集群
>
> <img src=".\asset\images\image-20250426170420049.png" alt="image-20250426170420049" style="zoom:67%;" />
>
> - 这里通过 deployment 管理稳定的pod
>
> - 再通过 service 提供负载均衡的访问
>
>   <img src=".\asset\images\image-20250426170934696.png" alt="image-20250426170934696" style="zoom:67%;" />

在 Kubernetes 集群中，每个 Node 运行一个 `kube-proxy` 进程。 `kube-proxy` 负责为 `Service` 实现了一种 VIP（虚拟 IP）的形式

在 Kubernetes v1.0 版本，代理完全在 userspace。在 Kubernetes v1.1 版本，新增了 iptables 代理，但并不是默认的运行模式。从 Kubernetes v1.2 起，默认就是 iptables 代理。在 Kubernetes v1.8.0-beta.0 中，添加了 ipvs 代理。



### Kubernetes Service 类型总结

| 类型             | 作用                                                         | 访问方式                                                     | 适用场景                                       |
| :--------------- | :----------------------------------------------------------- | :----------------------------------------------------------- | :--------------------------------------------- |
| **ClusterIP**    | 默认类型，分配一个**仅集群内部可访问的虚拟 IP**（VIP）       | 通过 `<ClusterIP>:<Port>` 在集群内部访问                     | 集群内服务间通信（如微服务间调用）             |
| **NodePort**     | 在 ClusterIP 基础上，在每个节点上绑定一个静态端口（默认范围：30000-32767） | 通过 `<任意NodeIP>:<NodePort>` 从集群外部访问（自动路由到 ClusterIP） | 开发测试环境或需要直接通过节点 IP 访问的场景   |
| **LoadBalancer** | 在 NodePort 基础上，通过云厂商创建外部负载均衡器             | 通过云厂商提供的负载均衡器域名/IP 访问（自动转发到 `<NodeIP>:<NodePort>`） | 生产环境需暴露服务到公网（依赖云平台支持）     |
| **ExternalName** | 将集群外部服务映射为集群内部别名（通过 DNS CNAME 记录）      | 通过 `<ServiceName>.<Namespace>.svc.cluster.local` 解析到外部服务的域名 | 集成外部服务（如数据库、第三方 API）到集群内部 |



> - clusterIP
>
> <img src=".\asset\images\image-20250426174418335.png" alt="image-20250426174418335" style="zoom: 50%;" />
>
> - nodeport
>
>   clusterIP提供的是内部访问IP的能力，nodeport 提供了 外部访问的能力
>
>   <img src=".\asset\images\image-20250426174712764.png" alt="image-20250426174712764" style="zoom:50%;" />
>
> 
>
> - LoadBalancer
>
>   ![image-20250426175519373](.\asset\images\image-20250426175519373.png)
>
> - ExternalName 
>
>   dns别名机制
>
>   <img src=".\asset\images\image-20250426175917798.png" alt="image-20250426175917798" style="zoom:50%;" />





### clusterIP

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-clusterip-deploy
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
      release: stable
      svc: clusterip
  template:
    metadata:
      labels:
        app: myapp
        release: stable
        env: test
        svc: clusterip
    spec:
      containers:
      - name: myapp-container
        image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
        imagePullPolicy: IfNotPresent
        ports:
        - name: http
          containerPort: 80
        readinessProbe:
          httpGet:
            path: /index1.html
            port: 80
          initialDelaySeconds: 1
          periodSeconds: 3
```



```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp-clusterip
  namespace: default
spec:
  type: ClusterIP
  selector:
    app: myapp
    release: stable
    svc: clusterip
  ports:
  - name: http
    port: 80   # 负载均衡的端口
    targetPort: 80  # 后端真实端口
    protocol: TCP  # 显式声明协议（可选，默认为TCP）
```

> 创建的 svc 会默认创建域名
>
> `servicename.namespace.svc.dominName.`
>
> ```bash
> [root@k8s-master01 chapter-6]# dig -t -A myapp-clusterip.default.svc.cluster.local. @10.0.0.10
> ;; Warning, ignoring invalid type -A
> 
> ; <<>> DiG 9.16.23-RH <<>> -t -A myapp-clusterip.default.svc.cluster.local. @10.0.0.10
> ;; global options: +cmd
> ;; Got answer:
> ;; WARNING: .local is reserved for Multicast DNS
> ;; You are currently testing what happens when an mDNS query is leaked to DNS
> ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 49121
> ;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
> ;; WARNING: recursion requested but not available
> 
> ;; OPT PSEUDOSECTION:
> ; EDNS: version: 0, flags:; udp: 4096
> ; COOKIE: 1597ae0261c159a0 (echoed)
> ;; QUESTION SECTION:
> ;myapp-clusterip.default.svc.cluster.local. IN A
> 
> ;; ANSWER SECTION:
> myapp-clusterip.default.svc.cluster.local. 30 IN A 10.12.48.187
> 
> ;; Query time: 34 msec
> ;; SERVER: 10.0.0.10#53(10.0.0.10)
> ;; WHEN: Mon Apr 07 23:03:12 CST 2025
> ;; MSG SIZE  rcvd: 139
> ```
>
> 因此 集群内部能够通过 clusterIP能够通过 IP 或者 域名进行访问
>
> ## **`internalTrafficPolicy` 的作用**
>
> 它有两个可选值：
>
> 1. **`Cluster`（默认值）**
>    - 流量可以路由到集群内**任何节点**上的 Pod（即使 Pod 不在客户端所在的节点）。
>    - 适用于大多数场景，确保负载均衡和高可用性。
> 2. **`Local`**
>    - 流量**优先路由到与客户端在同一节点上的 Pod**（如果存在）。
>    - 如果当前节点没有符合条件的 Pod，则请求会被丢弃（返回 `No endpoints available` 错误）。
>    - 适用于优化网络延迟或减少跨节点流量（如高吞吐量服务）。
>
> > ### ClusterIP 的工作原理：
> >
> > 1. **虚拟 IP 地址**：
> >    - 当你创建一个 `ClusterIP` 类型的 Service 时，Kubernetes 会为该服务分配一个虚拟的 IP 地址（通常是 `10.0.0.0/24` 范围内的地址），并将该地址与后端的 Pod（服务的选择器匹配的 Pod）关联起来。
> > 2. **集群内部通信**：
> >    - 集群内的其他 Pod 可以通过这个虚拟的 `ClusterIP` 地址访问服务。无论后端的 Pod 如何变动，虚拟 IP 地址始终是固定的，因此集群内部的其他组件无需关注 Pod 的实际 IP 地址变化。
> > 3. **负载均衡**：
> >    - `ClusterIP` 类型的 Service 会对所有访问该 IP 地址的流量进行负载均衡。即使服务有多个 Pod，Kubernetes 会自动将流量分发到这些 Pod 中。
> > 4. **内部 DNS 名称解析**：
> >    - Kubernetes 还为每个 `ClusterIP` 服务提供了一个 DNS 名称，格式通常是：`<service-name>.<namespace>.svc.cluster.local`。这个 DNS 名称也会解析到对应的 `ClusterIP`，例如：`my-service.default.svc.cluster.local`。这让集群内的应用可以通过 DNS 名称来访问服务，而不必关心 IP 地址。
> > 5. **可以在 Pod 内部使用 `curl` 命令（或其他网络工具）通过 `ClusterIP` 的域名或 IP 地址来访问 Service，从而访问后端的 Pod。**



### NodePort

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-nodeport-deploy
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
      release: stable
      svc: nodeport
  template:
    metadata:
      labels:
        app: myapp
        release: stable
        env: test
        svc: nodeport
    spec:
      containers:
      - name: myapp-container
        image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
        imagePullPolicy: IfNotPresent
        ports:
        - name: http
          containerPort: 80
          protocol: TCP  # 显式声明协议（可选，默认为TCP）
```



```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp-nodeport
  namespace: default
spec:
  type: NodePort
  externalTrafficPolicy: Local	# 流量路由策略
  selector:
    app: myapp
    release: stable
    svc: nodeport
  ports:
  - name: http
    protocol: TCP    # 显式声明协议（可选，默认为TCP）
    port: 80         # Service 监听的端口
    targetPort: 80   # 后端 Pod 的端口
    nodePort: 30010  # 节点暴露的端口（范围：30000-32767）
    
```

**`NodePort` 通过将请求转发到集群中的每个节点来实现外部访问。当你创建一个 `NodePort` 类型的 Service 时，Kubernetes 会在每个节点上打开指定的端口（在此例中是 30010）。这些端口会通过 kube-proxy 进行转发**。

- 当请求到达某个节点的 30001 端口时，kube-proxy 会通过负载均衡将流量转发到 `NodePort` 服务的后端 Pods。
- **这使得集群外部可以通过 `<NodeIP>:<NodePort>` 的形式访问该服务。**

> ### **1. `port`（Service 端口）**
>
> - **作用**：
>   Service 在集群内部的虚拟 IP（ClusterIP）上监听的端口。
>   **其他 Pod 通过 `<ServiceName>:<port>` 访问该服务**。
>
> - **示例**：
>
>   ```yaml
>   ports:
>     - port: 80  # 集群内访问 myapp-nodeport:80
>   ```
>
> - **实际场景**：
>   集群内其他服务访问 `http://myapp-nodeport.default.svc.cluster.local:80`。
>
> ------
>
> ### **2. `targetPort`（Pod 端口）**
>
> - **作用**：
>   Service 将流量转发到后端 Pod 的哪个端口。
>   必须与 Pod 中 `containerPort` 一致（否则流量无法到达容器）。
>
> - **示例**：
>
>   ```yaml
>   # Deployment 中定义的容器端口
>   containers:
>     - ports:
>         - containerPort: 8080  # Pod 实际监听 8080
>
>   # Service 中配置
>   ports:
>     - targetPort: 8080  # 将流量转发到 Pod 的 8080 端口
>   ```
>
> - **常见情况**：
>
>   - 如果 `targetPort` 未指定，默认与 `port` 相同。
>   - 如果 Pod 监听 `80`，但 Service 的 `port` 是 `8080`，则访问 `Service:8080` → `Pod:80`。
>
> ------
>
> ### **3. `nodePort`（节点端口）**
>
> - **作用**：
>   仅适用于 `type: NodePort` 或 `LoadBalancer`。
>   **在每个节点上开放一个静态端口（30000-32767），允许通过 `<NodeIP>:<nodePort>` 从外部访问。**
>
> - **示例**：
>
>   ```yaml
>   ports:
>     - nodePort: 30010  # 外部通过节点 IP:30010 访问
>   ```
>
> - **访问路径**：
>   `外部用户 → <任意节点IP>:30010 → Service:80 → Pod:8080`

> 同样，这里也支持 **`externalTrafficPolicy`** 的设置
>
> #### 1. **`Cluster`（默认模式）**
>
> - **作用**：
>   外部流量可以路由到集群中**任意节点**上的 Pod（即使 Pod 不在接收流量的节点上）。
> - **流量路径**：
>   `客户端 → 节点A（NodePort） → 转发到节点B上的 Pod`
> - **特点**：
>   - **负载均衡更均匀**：流量会均匀分布到所有可用 Pod。
>   - **丢失客户端真实 IP**：流量经过节点间的转发，源 IP 会被替换为节点 IP（需通过 `X-Forwarded-For` 头传递）。
>   - **潜在跨节点开销**：如果 Pod 不在接收流量的节点上，会产生额外的网络跳转。
>
> #### 2. **`Local`（本地模式）**
>
> - **作用**：
>   外部流量**仅路由到接收流量的节点上运行的 Pod**（如果该节点没有符合条件的 Pod，请求会被丢弃）。
> - **流量路径**：
>   `客户端 → 节点A（NodePort） → 仅转发到节点A上的 Pod`
> - **特点**：
>   - **保留客户端真实 IP**：流量不经过其他节点，源 IP 不会被替换。
>   - **负载不均**：如果 Pod 分布不均匀，某些节点可能负载过高。
>   - **需 Pod 与节点分布匹配**：如果节点没有 Pod，返回 `Connection refused`（需配合 Pod 亲和性或 DaemonSet 使用）。
>
> > 同时 因为nodeport方式是支持 clusterIP方式下的 集群内部的访问的，因此，这里的仍然支持上面的 **`internalTrafficPolicy` **策略，不过管控的范围就成了集群内部的访问了



### LoadBalancer

如果按照前面的设置nodeport，每个集群便拥有了外部访问的能力，同时，为了保证外部访问的高可用能力，也能在外部架设一层负载均衡，来保证外部用户对某个地址(这里能够指定外网地址)的稳定访问

<img src=".\asset\images\image-20250426195233205.png" alt="image-20250426195233205" style="zoom: 50%;" />



### ExternalName

**`ExternalName`** 是 Kubernetes 中的一种特殊 Service 类型，它的核心作用是将**集群外部的服务**（如数据库、第三方 API）映射为**集群内部的 DNS 名称**，使得集群内的 Pod 可以通过 Kubernetes 的 DNS 系统直接访问外部服务，而无需硬编码外部地址。

实际上更多的用法，假设你想要在 Kubernetes 集群内创建一个 `Service`，它将请求转发到外部的 `example.com` 网站。

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-external-service  # 定义的服务名	
spec:
  type: ExternalName     # 类型
  externalName: example.com  # 外部url
```



------

**核心作用**

1. **DNS 别名（CNAME 记录）**
   - 将形如 `<service-name>.<namespace>.svc.cluster.local` 的集群内域名，解析为外部服务的真实域名（如 `api.example.com`）。
   - **不创建任何代理或负载均衡**，纯粹是 DNS 层面的重定向。
2. **解耦外部依赖**
   - 允许在 Pod 中通过固定的 Kubernetes Service 名称访问外部服务，避免直接暴露外部地址（方便迁移或切换环境）。
3. **跨命名空间统一访问**
   - 例如，不同命名空间的 Pod 均可通过 `my-db-service.default.svc.cluster.local` 访问外部数据库，而无需关心实际地址。

![image-20250426200637314](.\asset\images\image-20250426200637314.png)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service-1
  namespace: default
spec:
  type: ExternalName
  externalName: www.baidu.com
  # 注意：ExternalName 类型的 Service 不需要 selector 和 ports
```



### Endpoints

<img src=".\asset\images\image-20250426202655056.png" alt="image-20250426202655056"  />

**1. Service 的作用**

- **定义逻辑集合**：
  Service 是 Kubernetes 中定义的一组 Pod 的**逻辑抽象**，提供统一的访问入口（如 ClusterIP、NodePort 等）。
- **访问策略**：
  通过 `Label Selector` 动态选择后端 Pod，并制定流量分发规则（如负载均衡、会话保持）。
- **核心功能**：
  - **服务发现**：集群内其他服务通过 Service 名称（DNS）访问 Pod，无需关心 Pod IP 变化。
  - **负载均衡**：自动将请求分发到多个 Pod。

**2. Endpoints 的作用**

- **实际端点集合**：
  Endpoints 是 Service 背后**具体可用的 Pod 地址列表**（IP:Port），**由 Kubernetes 自动维护。**
- **动态更新**：
  只有满足以下条件的 Pod 才会被加入 Endpoints：
  - 匹配 Service 的 `Label Selector`。
  - **Pod 状态为 `Running` 且通过就绪探针（`readinessProbe`）。**
- **与 Service 的关系**：
  - 每个 Service 默认会**自动创建一个同名的 Endpoints 对象**（除非手动指定 `externalName` 或显式定义 Endpoints）。



**3、关联体系**

- **配置 selector ---> 自动关联体系**

  配置了selector 会自动创建同名的 endpoint

  实际上由内部的endpoint通过拉取符合条件的pod信息，并且将其访问地址进行维护来提供的

  ![image-20250426203018835](.\asset\images\image-20250426203018835.png)

  > 创建service后，会自动创建对应的endpoint负责维护符合条件的pod的访问信息
  >
  > endpoint也是一种资源
  >
  > ![image-20250426204216105](.\asset\images\image-20250426204216105.png)

- **没有配置selector  ---> 手动关联体系**

  没有配置 selector 就不会自动创建同名的 endpoint 资源对象来进行处理

  管理员手动创建一个与之对应的 endpoint 来维护这个关系

  ![image-20250426203208920](.\asset\images\image-20250426203208920.png)

  > 1、创建不包含 selector 的 service
  >
  > ```yaml
  > apiVersion: v1
  > kind: Service
  > metadata:
  > name: nginx-noselector  # 修正拼写错误
  > spec:
  > ports:
  >     - protocol: TCP
  >       port: 6666          # Service 对外暴露的端口
  >       targetPort: 80      # 目标 Pod 的端口（需后续手动绑定 Endpoints）
  >   # 注意：此 Service 无 selector，需手动创建同名 Endpoints
  > ```
  >
  > 2、手动创建 endpoints 名字必须和service相同
  >
  > ```yaml
  > apiVersion: v1
  > kind: Endpoints
  > metadata:
  >   name: nginx-noselector  # 名字必须和service相同
  > subsets:
  >   - addresses:
  >       - ip: 192.168.66.13  # 后端服务的IP地址
  >     ports:
  >       - port: 80           # 后端服务的端口
  >         protocol: TCP      # 显式声明协议（可选，默认为TCP）
  > ```
  >
  > 3、手动在 ip: 192.168.66.13 创建一个容器 看是否能够被service访问
  >
  > ```bash
  > docker run --name nginx1  -p 80:80 -d swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  > ```
  >
  > 4、通过service访问
  >
  > ![image-20250426205909312](.\asset\images\image-20250426205909312.png)
  >
  > 这里经过测试 无论 endpoints 是否为手动创建，在清理掉其对应的service时会被一并删除

### publishNotReadyAddresses

  之前提到过，被service负载均衡维护需要有两个条件

1、selector 选中  ----》 无 selector 实际上会影响 endpoints的处理 这里能够通过手动创建对应的 selector处理

2、pod处于就绪状态

对于第二个条件，实际上是默认项，这里能够通过设置 `publishNotReadyAddresses` 属性设置

> 1、创建pod资源
>
> ```yaml
> apiVersion: v1
> kind: Pod
> metadata:
>   name: readiness-httpget-pod
>   namespace: default
>   labels:
>     app: myapp
>     env: test
> spec:
>   containers:
>   - name: readiness-httpget-container
>     image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
>     imagePullPolicy: IfNotPresent
>     ports:
>     - containerPort: 80  # 显式声明容器端口（可选但推荐）
>     readinessProbe:			# 创建了就绪探针 保证其一定不会通过
>       httpGet:
>         path: /index1.html
>         port: 80
>       initialDelaySeconds: 1
>       periodSeconds: 3
> ```
>
> ![image-20250426211322879](.\asset\images\image-20250426211322879.png)
>
> 2、创建service
>
> ```yaml
> apiVersion: v1
> kind: Service
> metadata:
>   name: myapp
>   labels:
>     app: myapp
> spec:
>   type: ClusterIP
>   selector:
>     app: myapp
>   ports:
>   - name: 80-80
>     protocol: TCP
>     port: 80
>     targetPort: 80
> ```
>
> > 未就绪pod无法被抓取
> >
> > <img src=".\asset\images\image-20250426211516049.png" alt="image-20250426211516049" style="zoom:67%;" />
>
> 3、开启 publishNotReadyAddresses
>
> ```bash
> kubectl patch service myapp -p '{"spec":{"publishNotReadyAddresses": true}}'
> ```
>
> 或者手动添加此属性
>
> > endpoints能够捕捉到未就绪的pod了
> >
> > <img src=".\asset\images\image-20250426211830040.png" alt="image-20250426211830040" style="zoom:67%;" />





## 存储

### 分类

#### 元数据管理

1. **ConfigMap**
   - 用途：存储非敏感的配置数据（明文形式）
   - 特点：可以存储键值对或整个配置文件，便于应用配置与容器镜像分离
2. **Secret**
   - 用途：存储敏感信息（如密码、密钥、令牌等）
   - 特点：数据以 base64 编码（非加密），提供基本的安全保护
   - 注意：虽然编码但仍需配合 RBAC 等安全措施
3. **Downward API**
   - 用途：允许容器在运行时获取自身或集群的信息
   - 特点：可以获取 Pod 名称、IP、标签、注解等信息
   - 使用方式：通过环境变量或卷文件方式暴露给容器

#### 真实数据存储

1. **Volume**
   - 用途：存储临时或持久性数据
   - 特点：
     - 生命周期通常与 Pod 绑定（部分类型除外）
     - 支持多种后端存储（emptyDir、hostPath、云存储等）
     - 适合缓存、临时文件等场景
2. **PersistentVolume (PV)**
   - 用途：提供持久化存储解决方案
   - 特点：
     - 独立于 Pod 生命周期的存储资源
     - 通过 PersistentVolumeClaim (PVC) 申请使用
     - 支持各种存储后端（NFS、iSCSI、云存储等）
     - 适合数据库、日志等需要持久化的场景



### ConfigMap

#### 基本介绍

ConfigMap 是 Kubernetes 在 1.2 版本引入的核心功能，主要用于解耦应用配置与容器镜像，实现配置的集中管理和动态注入。

##### 主要特性

1. **配置来源支持**：
   - 可以存储从配置文件、命令行参数或环境变量中读取的配置信息
   - 提供了统一的配置管理接口
2. **数据格式灵活性**：
   - 支持保存单个键值对属性
   - 可以保存完整的配置文件内容
   - 能够存储 JSON 等二进制对象
3. **注入机制**：
   - 提供了多种向容器注入配置的方式：
     - 挂载为卷（适合配置文件）
     - 设置为环境变量（适合单个参数）
     - 通过kubectl命令行参数传递

##### 使用场景

- 应用配置与镜像分离
- 不同环境(dev/test/prod)使用不同配置
- 热更新配置(部分场景支持不重启pod更新配置)
- 作为轻量级的配置中心使用

##### 注意事项

- 不适合存储敏感数据(应使用Secret)
- 大小限制(etcd限制通常为1MB)
- 更新后需要一定时间传播到各节点

ConfigMap 是 Kubernetes 配置即代码理念的重要实现，极大简化了云原生应用的配置管理工作。



> 如何使得不同服务内部的文件达成一致呢，目前主要由这两种方法
>
> 1、共享
>
> 类似 docker中持久卷的处理方法，将每个pod中的配置文件挂载到外部机器的某个对应盘中，在外部进行文件修改后，所有pod中的配置文件也就完成了修改
>
> 每一次在读取文件的时候，都会发生网络IO
>
> 2、注入(目前configmap的使用方式)
>
> 修改配置文件，自动的将新的配置文件更新掉原有的配置文件，
>
> 一次注入后，多次读取不再消耗网络IO，

> configMap 创建
>
> - --from-file
>
>   要求文件内部必须是一行一对的 k=v，可以在使用的时候注入至pod内部，变成环境变量
>
>   **可以在使用的时候注入至pod内部变成环境变量**
>
>   ```txt
>   # 1.txt
>   name=zhangsan
>   passwd=134
>   ```
>
>   文件内部不是一行一对的 k=v
>
>   ```
>   # 2.txt
>   taday is a happy day
>   ```
>
> 
>
> ## 从文件创建 ConfigMap --from-file
>
> ```bash
> $ kubectl create configmap game-config --from-file=hongfu.file
> ```
>
> **功能**：
>
> - 创建一个名为 `game-config` 的 ConfigMap
> - 内容来自本地文件 `hongfu.file`
>
> **特点**：
>
> - 文件内容会作为整体存储在 ConfigMap 中
> - ConfigMap 中的键名默认为文件名 `hongfu.file`
> - 适合导入完整的配置文件（如 JSON、YAML、properties 等）
> - 文件名为key 文件内容会变为value
>
> **使用场景**：
>
> - 需要将整个配置文件注入到容器中
> - 配置文件较大或结构复杂的情况
> - 保持原文件格式和内容不变
>
> ## 命令二：从字面值创建 ConfigMap --from-literal
>
> ```bash
> $ kubectl create configmap literal-config --from-literal=name=dave --from-literal=password=pass
> ```
>
> **功能**：
>
> - 创建一个名为 `literal-config` 的 ConfigMap
> - 包含两个键值对：
>   - `name=dave`
>   - `password=pass`
>
> **特点**：
>
> - 直接在命令行指定键值对
> - 每个 `--from-literal` 参数对应一个键值对
> - 适合简单的配置参数
>
> **使用场景**：
>
> - 需要设置少量独立配置项
> - 快速测试和调试
> - 不需要完整配置文件的情况

#### 使用方法

- **将当前的configMap注入到容器内部作为env进行调用**

  >```yaml
  >apiVersion: v1
  >kind: ConfigMap
  >    metadata:
  >      name: literal-config
  >  namespace: default
  >data:
  >      name: dave
  >      password: pass
  >---   # 多个配置文件能够通过 --- 分割
  >apiVersion: v1
  >kind: ConfigMap
  >metadata:
  >      name: env-config
  >      namespace: default
  >data:
  >  	log_level: INFO
  >---
  >apiVersion: v1
  >kind: Pod
  >metadata:
  >  name: cm-env-pod
  >spec:
  >  containers:
  >    - name: myapp-container
  >      image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >      command: ["/bin/sh", "-c", "env"]
  >      env:	# 通过指明字段的方式获取env
  >        - name: USERNAME
  >          valueFrom:
  >            configMapKeyRef:
  >              name: literal-config
  >              key: name
  >        - name: PASSWORD
  >          valueFrom:
  >            configMapKeyRef:
  >              name: literal-config
  >              key: password
  >      envFrom:		# 直接获取所有的configMap中的配置作为 env
  >        - configMapRef:
  >            name: env-config
  >  restartPolicy: Never
  >```
  >
  >![image-20250511123431296](.\asset\images\image-20250511123431296.png)

- **将当前的configMap作为pod的启动命令** 

  这里实际上还是配置env参数

  >```
  >apiVersion: v1
  >kind: Pod
  >metadata:
  >  name: cm-command-pod
  >spec:
  >  containers:
  >    - name: myapp-container
  >      image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >      command: ["/bin/sh", "-c", "echo $(USERNAME) $(PASSWORD)"]
  >      env:
  >        - name: USERNAME
  >          valueFrom:
  >            configMapKeyRef:
  >              name: literal-config
  >              key: name
  >        - name: PASSWORD
  >          valueFrom:
  >            configMapKeyRef:
  >              name: literal-config
  >              key: password
  >  restartPolicy: Never
  >```

- **卷挂载**

  > ```yaml
  > apiVersion: v1
  > kind: Pod
  > metadata:
  >   name: cm-volume-pod
  > spec:
  >     containers:
  >     - name: myapp-container
  >       image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >       volumeMounts:  # 将 config-volume 挂载到容器内的 /etc/config 目录
  >         - name: config-volume
  >           mountPath: /etc/config
  >   volumes:	# 将configmap声明为卷 名称为 config-volume
  >     - name: config-volume
  >       configMap:
  >         name: literal-config
  >   restartPolicy: Never
  > ```
  >
  > <img src=".\asset\images\image-20250511131809258.png" alt="image-20250511131809258" style="zoom:67%;" />
  >
  >
  > ![image-20250511132330148](.\asset\images\image-20250511132330148.png)

- **热更新**

> default.conf
>
> ```
> server {
>    listen 80 default_server;
>    server_name example.com www.example.com;
>    location / {
>        root   /usr/share/nginx/html;
>        index index.html index.htm;
>     }
> }
> ```
>
> ```yaml
> apiVersion: apps/v1
> kind: Deployment
> metadata:
>   labels:
>     app: hotupdate-deploy
>   name: hotupdate-deploy
> spec:
>   replicas: 5
>   selector:
>     matchLabels:
>       app: hotupdate-deploy
>   template:
>     metadata:
>       labels:
>         app: hotupdate-deploy
>     spec:
>       containers:
>         - name: nginx
>           image: nginx
>           volumeMounts:  
>             - name: config-volume
>               mountPath: /etc/nginx/conf.d/  # 接着将 config-volume卷挂载到容器内部
>     volumes:
>         - name: config-volume	# 还是同样将configmap 挂载为卷 config-volume
>           configMap:
>             name: default-nginx
> ```
>
> 
>
> 后续通过修改 configmap的方法，就能够更改deployment中部署的所有pod信息
>
> ```bash
> $ kubectl edif configmap default-nginx
> ```
>
> 变更后，配置并不会立即同步到所有pod防止瞬时压力过大，这里k8s会自己评估调度，如果pod有自己检查配置更新并更新的能力，这里的pod就会自己进行更新
>
> > ## ConfigMap 更新特性
> >
> > 1. **默认行为**：
> >    - 更新 ConfigMap 不会自动触发相关 Pod 的滚动更新
> >    - 这是 Kubernetes 的预期设计行为，避免配置变更导致意外的大规模重启
> > 2. **不同引用方式的更新表现**：
> >    - **环境变量方式 (Env)**：
> >      - **不会同步更新（Pod 创建后环境变量即固定）**
> >      - 必须重建 Pod 才能获取新值
> >    - **卷挂载方式 (Volume)**：
> >      - 会自动同步更新（约10秒延迟）
> >      - 更新通过 kubelet 定期检查机制实现
> >      - 文件内容会原子性替换（inotify 可检测到变化）
> >
> > ## 强制触发滚动更新的方法
> >
> > ```bash
> > kubectl patch deployment nginx-test --patch \
> > '{"spec": {"template": {"metadata": {"annotations": {"version/config":"66666666"}}}}}'
> > ```
> >
> > **实现原理**：
> >
> > - 通过修改 Deployment 的模板注解（annotations）
> > - Kubernetes 认为模板已变更，触发滚动更新
> > - 常用版本号或时间戳作为注解值



- 不可改变

> Kubernetes 提供了将 ConfigMap 和 Secret 标记为**不可变（immutable）**的特性，该特性通过 `immutable: true` 字段实现：
>
> ```yaml
> apiVersion: v1
> kind: ConfigMap
> metadata:
>     name: my-config
>   immutable: true  # 关键配置字段
> data:
>     config.json: |
>       { "env": "production" }
> ```
>
> ## 主要优势
>
> 1. **稳定性保障**：
>
>    - 彻底防止意外配置更新导致的应用程序中断
>    - 避免因配置变更引发的级联故障（特别对于大规模集群）
>
> 2. **性能提升**：
>
>    - 禁用对 ConfigMap/Secret 的监视机制
>    - 显著降低 kube-apiserver 的负载（减少 watch 操作）
>    - 对使用上万个 ConfigMap 的集群效果尤为明显
>
> 3. **操作是不可逆的**
>
>    当将其配置为不可修改后，既无法修改其他的相关配置，**也无法将再更新这个字段**，此configmap会变得无法更新



### Secret

#### 概述

**核心定位**

Secret 是 Kubernetes 专门设计用于存储敏感数据的资源对象，主要解决以下问题：

- **安全性问题**：避免敏感信息直接暴露在 Pod 定义或容器镜像中
- **灵活性问题**：实现敏感数据与应用的解耦管理
- **节点级分发**：Secret **仅分发到需要该 Secret 的 Pod 所在节点**
- **内存驻留机制**：通过 tmpfs 内存文件系统挂载，避免落盘
- **内存锁定**：防止交换到磁盘（需节点支持）
- **短生命周期**：kubelet 定期回收未使用的 Secret

**关键特性**

1. **安全存储机制**

- **Base64 编码**：默认采用 Base64 编码（注意：非加密）
- **加密支持**：可配合 [Encryption at Rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/) 实现静态加密
- **细粒度访问控制**：通过 RBAC 限制访问权限

2. **主要数据类型**

| 数据类型    | 典型用例       | 安全建议                   |
| :---------- | :------------- | :------------------------- |
| 用户名/密码 | 数据库凭证     | 定期轮换                   |
| API 令牌    | 服务间通信认证 | 最小权限原则               |
| TLS 证书    | HTTPS 通信     | 设置自动续期机制           |
| SSH 密钥    | 代码仓库访问   | 使用临时凭证最佳3.内置类型 |

**3.内置类型**

Kubernetes 提供了多种预定义的 Secret 类型，通过 `type` 字段指定：

| 类型                                | 用途描述         | 数据字段要求                      | 典型应用场景        |
| :---------------------------------- | :--------------- | :-------------------------------- | :------------------ |
| Opaque (默认)                       | 任意用户定义数据 | 无固定格式要求                    | 数据库凭证、API密钥 |
| kubernetes.io/service-account-token | 服务账户令牌     | 必须包含 `token`、`ca.crt` 等     | Pod身份认证         |
| kubernetes.io/dockercfg             | Docker仓库配置   | 必须包含 `.dockercfg` 文件        | 私有镜像库认证      |
| kubernetes.io/dockerconfigjson      | Docker配置JSON   | 必须包含 `.dockerconfigjson`      | 新版私有镜像库认证  |
| kubernetes.io/basic-auth            | 基础认证凭据     | 必须包含 `username` 和 `password` | HTTP基础认证        |
| kubernetes.io/ssh-auth              | SSH认证密钥      | 必须包含 `ssh-privatekey`         | Git仓库SSH访问      |
| kubernetes.io/tls                   | TLS证书文件      | 必须包含 `tls.crt` 和 `tls.key`   | HTTPS终端           |
| kubernetes.io/bootstrap-token       | 节点引导令牌     | 特定格式的token                   | 节点加入集          |



#### Opaque

1. **Opaque 类型定义**：
   - 是 Kubernetes Secret 的**默认类型**（当未显式指定 `type` 字段时自动生效）
   - 名称 "Opaque" 意为"不透明的"，表示该类型 Secret 的内容结构对 Kubernetes 不透明
   - 设计用于存储**任意用户自定义的键值对数据**
2. **编码规范**：
   - 所有数据值必须进行 **Base64 编码 ** k8s会对secret中的值进行编解码，如若没有经过编码，那么会发生错乱
   - 推荐使用 `echo -n "raw-value" | base64` 生成编码值
   - 避免在 YAML 中直接存储明文
3. **使用方法**

- **通过env使用secret**

> - base64 编码
>
> ![image-20250511171514954](.\asset\images\image-20250511171514954.png)
>
> ```yaml
> apiVersion: v1
> kind: Secret
> metadata:
>   name: mysecret
> type: Opaque
> data:
>   password: MWYyZDFlMmU2N2Rm
>   username: YWRtaW4=
> ```
>
> - 解码
>
>   ```bash
>   echo -n "covered-field" | base64 -d
>   ```
>
> - deployment
>
>   ```yaml
>   apiVersion: apps/v1
>   kind: Deployment
>   metadata:
>     labels:
>       app: opaque-secret-env
>     name: opaque-secret-env-deploy
>   spec:
>     replicas: 5
>     selector:
>       matchLabels:
>         app: op-se-env-pod
>     template:
>       metadata:
>         labels:
>           app: op-se-env-pod
>       spec:
>         containers:
>           - image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
>             name: myapp-container
>             ports:
>               - containerPort: 80
>             env:
>               - name: TEST_USER
>                 valueFrom:	# 值来源
>                   secretKeyRef:
>                     name: mysecret
>                     key: username	# mysecret 中的 key
>               - name: TEST_PASSWORD
>                 valueFrom:
>                   secretKeyRef:
>                     name: mysecret
>                     key: password
>   ```
>
> - 结果
>
>   <img src=".\asset\images\image-20250511172452150.png" alt="image-20250511172452150" style="zoom:67%;" />

- 通过volume使用

  >```yaml
  >apiVersion: v1
  >kind: Pod
  >metadata:
  >  labels:
  >    name: secret-volume
  >  name: secret-volume-pod
  >spec:
  >  volumes:
  >    - name: volumes12    	
  >      secret:
  >        secretName: mysecret  # 将secret声明为卷 名称为 volumes12
  >        defaultMode: 256    # 能够设置默认属性
  >  containers:
  >    - image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >      name: myapp-container
  >      volumeMounts:
  >        - name: volumes12     # 将 volumes12 挂载到pod 内部的 /data 下
  >          mountPath: "/data"
  >```
  >
  >能够指定声明单独的键
  >
  >```yaml
  >apiVersion: v1
  >kind: Pod
  >metadata:
  >  labels:
  >    name: secret-volume
  >  name: secret-volume-pod
  >spec:
  >  volumes:
  >  - name: volumes12
  >  	secret:
  >  	  secretName: mysecret
  >  	  items:
  >  	  - key: username
  >  	    path: my-group/my-username
  >  containers:
  >    - image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >      name: myapp-container
  >      volumeMounts:
  >        - name: volumes12     # 将 volumes12 挂载到pod 内部的 /data 下
  >          mountPath: "/data"
  >```



##### 更新机制

**常规卷挂载更新流程**：

- **kubelet 周期性检查**：默认每 1 分钟同步检查 Secret 更新（可通过 `--sync-frequency` 调整）
- **原子性更新**：当检测到变更时，kubelet 会原子性替换整个卷目录

- **延迟性**：更新最终一致，通常需要 10-60 秒传播到所有节点

**本地缓存行为**：

- 各节点 kubelet 维护独立缓存
- 新 Pod 调度到节点时使用该节点当前缓存值
- 可能导致不同节点上的 Pod 短暂看到不同 Secret 值

**子路径挂载限制**：

- **不会接收更新**：即使 Secret 变更，已挂载的子路径文件保持旧值
- **设计原理**：subPath 视为静态文件引用，不纳入 kubelet 同步范围

> | 特性             | 常规卷挂载     | 子路径挂载              |
> | :--------------- | :------------- | :---------------------- |
> | 挂载范围         | 整个卷内容     | 仅卷中的特定文件/子目录 |
> | 更新机制         | 自动同步更新   | 不接收更新              |
> | 目标路径         | 必须为目录     | 可以是文件或目录        |
> | 是否覆盖现有内容 | 会覆盖整个目录 | 精确替换特定路径        |



##### immutable

不可变 Secret 是 Kubernetes 提供的一种安全强化机制，通过设置 `immutable: true` 字段实现：

**1. 稳定性保障**

- **防意外修改**：彻底杜绝配置变更导致的应用程序中断
- **版本控制**：配合命名规范（如 `-v1`、`-v2` 后缀）实现明确的版本管理
- **审计追踪**：每个版本独立存在，便于问题追溯

**2. 性能优化**

- **API 服务器减负**：
- **典型性能提升**：
  - 减少 40-60% 的 kube-apiserver 内存占用（万级 Secret 场景）
  - 降低 30% 以上的 etcd 存储压力







### downwardAPI

Downward API 是 Kubernetes 提供的一种特殊机制，允许容器**无需直接访问 API Server**即可获取自身及其运行环境的元数据。它通过两种方式暴露数据：

```yaml
env:
  - name: POD_NAME
    valueFrom:
      fieldRef:
        fieldPath: metadata.name
```

或 volume

```yaml
volumeMounts:  # 绑定卷到对应的目录下
  - name: podinfo
    mountPath: /etc/podinfo
volumes:    # 声明卷
  - name: podinfo
    downwardAPI:
      items:
        - path: "labels"
          fieldRef:
            fieldPath: metadata.labels
```

> ## 使用模式对比
>
> | 注入方式 | 优点                          | 缺点               | 适用场景              |
> | :------- | :---------------------------- | :----------------- | :-------------------- |
> | 环境变量 | 使用简单，应用无感知          | 不支持复杂数据结构 | 简单元数据（名称/IP） |
> | 卷挂载   | 支持完整元数据结构 支持热更新 | 需文件读取逻辑     | 标签/注解等复杂数据   |
> | 两者结合 | 灵活性强                      | 配置复杂度高       | 混合需求场景          |

- 通过 ENV

  >```yaml
  >apiVersion: v1
  >kind: Pod
  >metadata:
  >  name: downward-api-env-example
  >spec:
  >  containers:
  >    - name: my-container
  >      image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >      env:
  >        - name: POD_NAME
  >          valueFrom:
  >            fieldRef:
  >              fieldPath: metadata.name
  >        - name: POD_NAMESPACE
  >          valueFrom:
  >            fieldRef:
  >              fieldPath: metadata.namespace
  >        - name: POD_IP
  >          valueFrom:
  >            fieldRef:
  >              fieldPath: status.podIP
  >        - name: CPU_REQUEST
  >          valueFrom:
  >            resourceFieldRef:
  >              resource: requests.cpu
  >        - name: CPU_LIMIT
  >          valueFrom:
  >            resourceFieldRef:
  >              resource: limits.cpu
  >        - name: MEMORY_REQUEST
  >          valueFrom:
  >            resourceFieldRef:
  >              resource: requests.memory
  >        - name: MEMORY_LIMIT
  >          valueFrom:
  >            resourceFieldRef:
  >              resource: limits.memory
  >  restartPolicy: Never
  >```
  >
  >
  >
  ><img src=".\asset\images\image-20250511180652978.png" alt="image-20250511180652978" style="zoom:67%;" />



- volume

  同样使用卷挂载的方式 也能够热更新

  >```yaml
  >apiVersion: v1
  >kind: Pod
  >metadata:
  >  name: downward-api-volume-example
  >spec:
  >  containers:
  >    - name: my-container
  >      image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
  >      resources:
  >        limits:
  >          cpu: "1"
  >          memory: "512Mi"
  >        requests:
  >          cpu: "0.5"
  >          memory: "256Mi"
  >      volumeMounts:
  >        - name: downward-api-volume
  >          mountPath: /etc/podinfo
  >  volumes:		# 声明卷
  >    - name: downward-api-volume
  >      downwardAPI:
  >        items:
  >          - path: "annotations"
  >            fieldRef:
  >              fieldPath: metadata.annotations
  >          - path: "labels"
  >            fieldRef:
  >              fieldPath: metadata.labels
  >          - path: "name"
  >            fieldRef:
  >              fieldPath: metadata.name
  >          - path: "namespace"
  >            fieldRef:
  >              fieldPath: metadata.namespace
  >          - path: "uid"
  >            fieldRef:
  >              fieldPath: metadata.uid
  >          - path: "cpuRequest"
  >            resourceFieldRef:
  >              containerName: my-container
  >              resource: requests.cpu
  >          - path: "memoryRequest"
  >            resourceFieldRef:
  >              containerName: my-container
  >              resource: requests.memory
  >          - path: "cpuLimit"
  >            resourceFieldRef:
  >              containerName: my-container
  >              resource: limits.cpu
  >          - path: "memoryLimit"
  >            resourceFieldRef:
  >              containerName: my-container
  >              resource: limits.memory
  >  restartPolicy: Never
  >```
  >
  >
  >
  ><img src=".\asset\images\image-20250511181402586.png" alt="image-20250511181402586" style="zoom:67%;" />



Downward API 是 Kubernetes 提供的**轻量级元数据注入机制**，主要特点包括：

- **自省性**：**仅能获取当前 Pod 自身的元数据**

- **有限字段**：支持约 20 个预设字段（如 `metadata.name`, `status.podIP` 等）

- **无权限需求**：容器内进程无需 API 访问权限

downwardAPI 仅仅能够获取当前的pod的信息，通过设置权限的方式访问apiServer来获取非自己pod的原信息

> ```yaml
> [root@k8s-master01 apiServer]# cat 1.RBAC.yaml 
> apiVersion: rbac.authorization.k8s.io/v1
> kind: ClusterRoleBinding
> metadata:
>  name: test-api-cluster-admin-binding
> subjects:
> - kind: ServiceAccount
>  name: test-api
>  namespace: default
> roleRef:
>  kind: ClusterRole
>  name: cluster-admin
>  apiGroup: rbac.authorization.k8s.io
> [root@k8s-master01 apiServer]# cat 2.pod.yaml 
> apiVersion: v1
> kind: Pod
> metadata:
>  name: curl
> spec:
>  serviceAccountName: test-api
>  containers:
>  - name: main
>    image: tutum/curl
>    command: ["sleep", "9999"]
> root@curl:/# TOKEN=$( cat /var/run/secrets/kubernetes.io/serviceaccount/token )
> root@curl:/# CAPATH="/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
> root@curl:/# NS=$( cat /var/run/secrets/kubernetes.io/serviceaccount/namespace )
> root@curl:/# curl -H "Authorization: Bearer $TOKEN" --cacert $CAPATH 
> https://kubernetes/api/v1/namespaces/$NS/pods
> ```
>
> 



### volume

> ### **Volume 的作用：**
>
> - **申请存储空间**：通过创建一个 Volume，Kubernetes 允许你为 Pod 提供持久化存储或者临时存储空间，供 Pod 内的容器使用。
> - **挂载存储**：Kubernetes 将创建的 Volume 挂载到 Pod 内部的一个或多个容器上。容器可以通过指定的路径访问这个 Volume。
> - **名称和路径**：Volume 是通过在 Pod 的 YAML 配置文件中指定的名称来标识的，然后该名称与挂载路径关联，容器通过挂载路径来访问 Volume。
>
> ### **具体操作流程：**
>
> 1. **申请存储空间**：首先，你创建一个 Kubernetes Volume 资源（例如 `emptyDir`、`hostPath`、`PersistentVolume` 等），并且将它指定给 Pod 中的容器使用。
> 2. **为 Volume 命名**：你给 Volume 指定一个名称，之后容器通过这个名称来引用这个存储。
> 3. **挂载存储空间**：你将 Volume 挂载到 Pod 内部的一个路径上。容器将通过这个挂载路径访问 Volume 中的存储空间。



**临时性存储问题**

1. **数据易失性**：

   - 容器崩溃重启后，所有写入默认存储层的文件都将丢失
   - 每次重启都恢复到镜像初始状态

2. **多容器隔离**：

   - 同一 Pod 内容器的文件系统相互隔离
   - 缺乏原生共享存储机制

3. **核心功能矩阵**

   | 功能       | 描述                                 |
   | :--------- | :----------------------------------- |
   | 数据持久化 | 超越容器生命周期的数据存储           |
   | 多容器共享 | 同一 Pod 内多个容器可访问相同存储    |
   | 数据安全   | 支持多种持久化后端（如云存储/NFS等） |
   | 动态供给   | 通过 StorageClass 实现按需分配       |

4. **基础类型对比**

   | 类型             | 生命周期                                                | 适用场景                  | 示例配置                    |
   | :--------------- | :------------------------------------------------------ | :------------------------ | :-------------------------- |
   | emptyDir         | 随 Pod 存在                                             | 临时存储/容器间共享       | `emptyDir: {}`              |
   | hostPath         | 与节点生命周期相同，<br />将pod内部与host主机路径相关联 | 开发测试/访问节点系统文件 | `hostPath: { path: /data }` |
   | configMap/Secret | 独立管理                                                | 配置注入                  | 见前文示例                  |
   | persistentVolume | 持久化存储                                              | 数据库/有状态服务         | 关联 PVC 使用               |



#### emptyDir

emptyDir 卷为容器提供了灵活高效的**临时存储方案**，特别适合需要快速暂存数据的场景。正确使用可以显著提升有状态应用的可靠性，同时避免持久化存储的开销。

emptyDir 是 Kubernetes 中最基础的临时存储卷类型，具有以下关键特性：

1. **生命周期绑定**：

   绑定的时pod，对于pod中的容器崩溃 不会影响声明周期

   - 创建时机：Pod 被分配到节点时即时创建
   - 销毁条件：Pod 从节点移除时（非容器崩溃）

2. **存储介质选择**：

   - `medium` 类型：
     - 空值/未指定：节点默认存储（通常是磁盘）
     - `Memory`：使用 tmpfs（内存文件系统）
   - `sizeLimit`：配额限制（超过可能触发驱逐）

##### 使用方式

**disk共享**

> ```yaml
> apiVersion: v1
> kind: Pod
> metadata:
> name: volume-emptydir-disk-pod
> namespace: default
> spec:
>   containers:
>     - name: myapp
>       image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
>       ports:
>         - containerPort: 80
>       volumeMounts:
>         - name: logs-volume
>           mountPath: /usr/local/nginx/logs
>     - name: busybox
>       image: registry.cn-hangzhou.aliyuncs.com/wangyangshare/share:busybox
>       command: ["/bin/sh", "-c", "touch /logs/access.log && tail -f /logs/access.log"]
>       volumeMounts:
>         - name: logs-volume
>           mountPath: /logs
>   volumes:		# 声明 emptyDir 卷
>     - name: logs-volume
>       emptyDir: {}  # 初始化空的 emptyDir 必须加{}
> ```
>
> 这里的逻辑就是 myapp 将声明的卷挂载到 /usr/local/nginx/logs 写入自己的日志
>
> busybox 将声明的卷挂载到 /logs 下，并不断读取和打印其下的access.log 文件 因为这两个的容器是共享这里声明的 logs-volumn的。
>
> > #### **特点：**
> >
> > - **生命周期**：`emptyDir` 是一个临时的、空的目录，挂载到 Pod 中的容器里。这个目录的生命周期与 Pod 相同，Pod 一旦删除，`emptyDir` 中的数据也会丢失。
> > - **存储介质**：默认情况下，`emptyDir` 使用集群节点的本地磁盘空间（通常是磁盘文件系统）。但是，在某些情况下，如果使用了 `emptyDir` 和 `memory` 配合，它的存储可以放在 RAM 中（即 `memory` 类型）。
> >
> > #### **使用场景：**
> >
> > - **临时存储**：适用于 Pod 内部容器之间需要共享数据、或者用于临时存储文件的场景。
> > - **缓存数据**：当容器需要临时缓存数据但不要求持久化时，可以使用 `emptyDir`。
> > - **进程间通信**：多个容器共享 `emptyDir` 时，它可以用于不同容器之间的文件传递和数据共享。
>
> 

**共享内存**

> ```yaml
> apiVersion: v1
> kind: Pod
> metadata:
> name: volume-emptydir-mem
> namespace: default
> spec:
>   containers:
>     - name: myapp
>       image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
>       ports:
>         - containerPort: 80
>       resources:
>         limits:
>           cpu: "1"
>           memory: "1024Mi"
>         requests:
>           cpu: "1"
>           memory: "1024Mi"
>       volumeMounts:
>         - name: mem-volume
>           mountPath: /data
>   volumes:
>     - name: mem-volume
>       emptyDir:
>         medium: Memory		# 共享内存方式	
>         sizeLimit: "500Mi"	# 最大限制
> ```
>
> > #### **特点：**
> >
> > - **存储介质**：当 `emptyDir` 设置为 `medium: Memory` 时，`emptyDir` 的存储会使用节点的 **RAM**（内存）而不是磁盘。
> > - **数据存储**：与普通 `emptyDir` 一样，它也会在 Pod 生命周期内存在，但不同之处在于它会更快，因为它是存储在内存中的，而不是磁盘。
> > - **性能**：`memory` 类型的 Volume 会比磁盘存储更快，适用于需要高性能、高速存取的场景，但它的数据不会持久化，一旦 Pod 重启，数据就会丢失。
> >
> > #### **使用场景：**
> >
> > - **缓存或临时计算数据**：适用于需要快速存取的缓存、计算数据，且不要求持久化存储。
> > - **提高性能**：如果某些应用（如数据库缓存、session 数据）需要更高的 I/O 性能，可以使用 `memory` 类型的存储。
> > - **敏感数据**：如果存储的数据不希望被持久化，可以使用 `memory` 类型，以避免存储泄漏或数据丢失。

#### hostpath

hostPath 卷允许将**节点主机上的文件系统路径**直接挂载到 Pod 中，实现容器与主机间的文件系统互通。这是一种特殊的持久化存储方式，**但数据生命周期与节点而非 Pod 绑定**。

**1. 基础配置**

```yaml
volumes:
  - name: hostpath-volume
    hostPath:
      path: /data/pod-data  # 主机绝对路径
      type: DirectoryOrCreate  # 类型约束
```

**2. 类型(Type)规范**

| 类型值            | 说明                                                       |
| :---------------- | :--------------------------------------------------------- |
|                   | 空字符串, 默认,意味着再挂载 hostPath卷之前不会执行任何检查 |
| Directory         | 必须存在的目录（不存在则报错）                             |
| DirectoryOrCreate | 目录不存在时自动创建（权限755，属主kubelet）               |
| File              | 必须存在的文件（不存在则报错）                             |
| FileOrCreate      | 文件不存在时自动创建（权限644，属主kubelet）               |
| Socket            | 必须存在的 UNIX 套接字                                     |
| CharDevice        | 必须存在的字符设备                                         |
| BlockDevice       | 必须存在的块设备                                           |

**3.注意事项**

**节点间行为不一致性**

**问题本质**：

- hostPath 挂载的是节点本地存储，不同节点的目录结构和内容可能不同
- 导致相同 Pod 模板在不同节点产生不同行为

**资源调度盲区**

**核心问题**：

- Kubernetes 调度器无法感知 hostPath 占用的存储资源
- 可能导致节点磁盘空间耗尽

**权限管理**

向主机本地卷进行读写需要有相对应的权限

![deepseek_mermaid_20250511_2afc96](C:\Users\xlwwwww\Downloads\deepseek_mermaid_20250511_2afc96.png)



> ```yaml
> apiVersion: v1
> kind: Pod
> metadata:
>  name: hostpath-pod
> spec:
>  containers:
>  - image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1
>    name: myapp
>    volumeMounts:
>    - mountPath: /test-pd
>      name: test-volume
>  volumes:
>  - name: test-volume
>    hostPath:
>      path: /data   # directory location on host
>      type: Directory  # this field is optional 类型检测
> ```
>
> 



### PV/PVC

pvc --> 抽象出来的存储需求

pv --> 实际的存储

通过pvc和pv的匹配和绑定策略，将pod的存储需求和实际的存储相分离解耦

#### **1. PersistentVolume (PV)**

**定义**：集群级别的存储资源，由管理员预先配置或通过 StorageClass 动态供给

**关键特性**：

- 生命周期独立于 Pod
- 支持多种存储后端（NFS、iSCSI、云存储等）
- 可设置回收策略（Retain/Delete/Recycle）

#### 2. PersistentVolumeClaim (PVC)

**定义**：用户对存储资源的请求，相当于存储资源的"消费方"

**核心功能**：

- 指定所需存储大小和访问模式
- 自动绑定符合条件的 PV
- 为 Pod 提供持久化存储抽象



#### PV/PVC 通用参数

| 参数                       | 可选值                                   | 说明           |
| :------------------------- | :--------------------------------------- | :------------- |
| accessModes                | ReadWriteOnce/ReadOnlyMany/ReadWriteMany | 定义访问模式   |
| storageClassName           | 字符串                                   | 指定存储类名称 |
| resources.requests.storage | 如 "100Gi"                               | 请求的存储大小 |

#### PV 专有参数

| 参数                          | 说明                                                  |
| :---------------------------- | :---------------------------------------------------- |
| persistentVolumeReclaimPolicy | Retain/Delete/Recycle (动态供给仅支持Delete)          |
| volumeMode                    | Filesystem/Block (决定作为文件系统还是原始块设备暴露) |
| nodeAffinity                  | 限制可使用该PV的节点(本地存储必备)                    |



#### 匹配条件

**1. 容量匹配规则**

- **刚性要求**：PV的容量必须 ≥ PVC请求的容量
- **最佳实践**：尽量保持PV容量与PVC请求一致，避免资源浪费
- **特殊场景**：动态扩容时需确保存储后端支持（CSI Volume Expansion）

**2. 访问模式匹配**

访问模式必须**精确匹配**，不存在降级兼容：

| PVC访问模式                    | 可绑定的PV访问模式 | 典型存储类型           |
| :----------------------------- | :----------------- | :--------------------- |
| ReadWriteOnce(RWO) 单节点读写  | 必须包含RWO        | 本地SSD、AWS EBS       |
| ReadOnlyMany(ROX)  多节点只读  | 必须包含ROX        | NFS、CephFS            |
| ReadWriteMany(RWX)  多节点只写 | 必须包含RWX        | Azure Files、GlusterFS |

**3. StorageClass匹配**

- **显式指定**：PV和PVC的`storageClassName`字段必须完全一致
- **空值处理**：
  - PVC未指定storageClass → 绑定到未设置storageClass的PV
  - PVC设置storageClass="" → 强制绑定到未设置storageClass的PV

#### 回收策略

当绑定的pv&PVC中，pvc被删除后，pv的默认处理规则

| 策略    | 动态供给 | 行为                                      |
| :------ | :------- | :---------------------------------------- |
| Retain  | ❌        | 保留PV和数据，需手动清理，不允许新pvc复用 |
| Delete  | ✅        | 自动删除PV和关联的外部存储资源            |
| Recycle | ❌        | 废弃策略，基本数据擦除(不推荐使用)        |

#### PV/PVC状态

**1. Available (可用状态)**

**特征**：

- PV 已创建但未绑定到任何 PVC
- 可被合适的 PVC 绑定
- 新创建的 PV 初始状态

**2. Bound (已绑定状态)**

**特征：**

- PV 与 PVC 成功绑定
- 绑定关系显示在 PV 的 `spec.claimRef` 字段
- 此时 PV 不能被其他 PVC 绑定

**3. Released (已释放状态)**

**触发条件**：

1. 绑定的 PVC 被删除
2. PV 的回收策略为 `Retain`

**关键行为**：

- 数据保留在存储系统上
- PV 不能被新 PVC 直接绑定，需手动清理
- 需要管理员介入处理：

**4. Failed (失败状态)**

**常见诱因**：

- 动态供给的 PV 在删除时后端存储释放失败
- 手动创建的 PV 存储不可访问
- 回收操作（如 `Recycle` 策略）执行失败



#### PVC保护

pvc保护的目的是确保pod正在使用的PVC不会从系统中移除，因为如果被移除的话可能会导致数据丢失

当pod状态为 `pending` 并且pod 已经分配给节点或 pod 为 `running` 状态时，pvc处于活动状态

当启用 pvc保护功能时，如果用户删除了一个 pod 正在使用的PVC，该PVC不会被立即删除，PVC的删除将被推迟，直到PVC不再被任何pod使用.





> ```bash
> [root@k8s-master01]# yum install -y nfs-utils rpcbind
> [root@k8s-master01 ]# mkdir /nfsdata
> [root@k8s-master01 ]# chmod 666 /nfsdata
> [root@k8s-master01 ]# chown nobody /nfsdata
> [root@k8s-master01 nfsdata]# mkdir {1..10}
> [root@k8s-master01 nfsdata]# ls
> 1  10  2  3  4  5  6  7  8  9
> [root@k8s-master01 nfsdata]# echo "1" > 1/index.html
> [root@k8s-master01 nfsdata]# echo "1" > 2/index.html
> [root@k8s-master01 nfsdata]# echo "2" > 2/index.html
> [root@k8s-master01 nfsdata]# echo "3" > 3/index.html
> [root@k8s-master01 nfsdata]# echo "4" > 4/index.html
> [root@k8s-master01 nfsdata]# echo "5" > 5/index.html
> [root@k8s-master01 nfsdata]# echo "6" > 6/index.html
> [root@k8s-master01 nfsdata]# echo "7" > 7/index.html
> [root@k8s-master01 nfsdata]# echo "8" > 8/index.html
> [root@k8s-master01 nfsdata]# echo "9" > 9/index.html
> [root@k8s-master01 nfsdata]# echo "10" > 10/index.html
> ```
>
> 添加10个
>
> ```bash
> [root@k8s-master01 nfsdata]# cat /etc/exports 
> /nfsdata/1 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/2 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/3 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/4 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/5 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/6 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/7 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/8 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/9 *(rw,no_root_squash,no_all_squash,sync)
> /nfsdata/10 *(rw,no_root_squash,no_all_squash,sync
> ```
>
> 重启
>
> ```bash
> systemctl restart rpcbind
> systemctl restart nfs-server
> ```
>
> 
>
> 部署PV
>
> ```yaml
> apiVersion: v1
> kind: PersistentVolume
> metadata:
> name: nfspv1
> spec:
> capacity:
>   storage: 1Gi
>   accessModes:
>     - ReadWriteOnce
>   persistentVolumeReclaimPolicy: Recycle
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/1
>     server: 10.66.66.11
> ---
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: nfspv2
> spec:
>   capacity:
>     storage: 0.9Gi
>   accessModes:
>     - ReadWriteOnce
>   persistentVolumeReclaimPolicy: Recycle
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/2
>     server: 10.66.66.11
> ---
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: nfspv3
> spec:
>   capacity:
>     storage: 1.2Gi
>   accessModes:
>     - ReadWriteOnce
>   persistentVolumeReclaimPolicy: Recycle
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/3
>     server: 10.66.66.11
> ---
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: nfspv4
> spec:
>   capacity:
>     storage: 1Gi
>   accessModes:
>     - ReadWriteOnce
>   persistentVolumeReclaimPolicy: Recycle
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/4
>     server: 10.66.66.11
> ---
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: nfspv5
> spec:
>   capacity:
>     storage: 1Gi
>   accessModes:
>     - ReadWriteMany
>   persistentVolumeReclaimPolicy: Recycle
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/5
>     server: 10.66.66.11
> ---
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: nfspv6
> spec:
>   capacity:
>     storage: 1Gi
>   accessModes:
>     - ReadWriteOnce
>   persistentVolumeReclaimPolicy: Recycle
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/6
>     server: 10.66.66.11
> ---
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: nfspv7
> spec:
>   capacity:
>     storage: 1Gi
>   accessModes:
>     - ReadWriteOnce
>   persistentVolumeReclaimPolicy: Retain
>   storageClassName: nfs
>   nfs:
>     path: /nfsdata/7
>     server: 10.66.66.11
> ```
>
> 创建service&statefulset
>
> ```yaml
> apiVersion: v1
> kind: Service
> metadata:
>   name: nginx
>   labels:
>     app: nginx
> spec:
>   ports:
>     - port: 80
>       name: web
>   clusterIP: None     #  创建无头服务，用于StatefulSet的DNS解析 每个Pod将获得独立的网络标识：web-0.nginx, web-1.nginx, web-2.nginx
>   selector:
>     app: nginx
> ---
> apiVersion: apps/v1
> kind: StatefulSet
> metadata:
>   name: web
> spec:
>   selector:
>     matchLabels:
>       app: nginx
>   serviceName: "nginx"
>   replicas: 3
>   template:
>     metadata:
>       labels:
>         app: nginx
>     spec:
>       containers:
>         - name: nginx
>           image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1.0
>           ports:
>             - containerPort: 80
>               name: web
>           volumeMounts:
>             - name: www
>               mountPath: /usr/local/nginx/html
>   volumeClaimTemplates:		# pvc 声明
>     - metadata:
>         name: www
>       spec:
>         accessModes: ["ReadWriteOnce"]
>         storageClassName: "nfs"
>         resources:
>           requests:
>             storage: 1Gi
> ```
>
> > #### 一、基础特性
> >
> > #### 1. 有序部署与扩展
> >
> > - **顺序创建**：Pod 按序号顺序创建（如 web-0, web-1, web-2）
> >
> > - **逆序终止**：删除时按相反顺序终止（web-2 先于 web-1）
> >
> > - **滚动更新策略**：
> >
> >   ```yaml
> >   updateStrategy:
> >     type: RollingUpdate
> >     rollingUpdate:
> >       partition: 1 # 控制分阶段更新 仅更新大于该序号的pod实例
> >   ```
> >
> > #### 2. 稳定的网络标识
> >
> > - **DNS格式**：`<pod-name>.<service-name>.<namespace>.svc.cluster.local`
> >
> > - **Headless Service 要求**：
> >
> >   service必须配置为无头服务
> >   
> >   ```yaml
> >   spec:
> >     clusterIP: None # 必须配置为无头服务
> >   ```
> >
> > > ### **kubernetes Pod 的 DNS 访问：**
> > >
> > > 在 Kubernetes 中，**所有 Pod 都能通过 DNS 名称进行访问**，其形式通常是：
> > >
> > > ```
> > > <pod-name>.<workload-name>.<namespace>.svc.cluster.local
> > > ```
> > >
> > > 这里的 `<pod-name>` 是 Pod 的名称，`<workload-name>` 是 Pod 所属于的控制器（例如 `Deployment`、`StatefulSet` 等）的名字，`<namespace>` 是所在的命名空间。
> > >
> > > ### **区别：**
> > >
> > > - **对于一般的控制器（如 Deployment、ReplicaSet）管理的 Pod：**
> > >
> > >   - Pod 的名称是 **动态的**，通常是由 Kubernetes 自动生成的，并且包括一个随机的字符串。每次 Pod 被重建时，Pod 的名称都会变化。
> > >   - 由于 Pod 的名称不稳定，DNS 名称也会随之变化。因此，**Pod 的 DNS 名称不稳定**，在多次重建或调度的情况下，Pod 的 DNS 名称也会改变。
> > >
> > >   **例如**：对于一个 `Deployment`，Pod 的 DNS 名称可能是：
> > >
> > >   ```
> > >   myapp-abcde.default.svc.cluster.local
> > >   ```
> > >
> > >   如果这个 Pod 被重建，新的 Pod 可能被命名为：
> > >
> > >   ```
> > >   myapp-fghij.default.svc.cluster.local
> > >   ```
> > >
> > >   所以，**Pod 名称不稳定**，DNS 名称也不稳定。
> > >
> > > - **对于 StatefulSet 管理的 Pod：**
> > >
> > >   - `StatefulSet` 中的 Pod 名称是 **固定的**，每个 Pod 都会有一个固定的索引（如 `web-0`、`web-1` 等），即使 Pod 被删除并重新创建，它的名称也不会改变。
> > >   - 由于 Pod 的名称是稳定的，`StatefulSet` 中的 Pod 会分配一个 **稳定的 DNS 名称**，这对于需要固定标识（例如数据库、分布式系统）的有状态应用来说非常重要。
> > >
> > >   **例如**：对于一个 `StatefulSet` 名为 `web`，它的 Pod DNS 名称将是：
> > >
> > >   ```
> > >   web-0.web.default.svc.cluster.local
> > >   web-1.web.default.svc.cluster.local
> > >   web-2.web.default.svc.cluster.local
> > >   ```
> > >
> > >   这些 DNS 名称 **保持不变**，即使 Pod 被重启或重新调度，DNS 名称也不会改变。
> > >
> > > ### **总结：**
> > >
> > > - **所有 Pod 都能通过 DNS 访问**，但普通 Pod 的 DNS 名称会随着 Pod 重建而变化。
> > > - **StatefulSet 的 Pod 会有稳定的 DNS 名称**，因为它们的名称（如 `web-0`、`web-1`）是固定的，即使 Pod 重启或迁移到其他节点，DNS 名称也不会改变。
> > > - **Pod 的 DNS 访问**形式为 `<pod-name>.<workload-name>.<namespace>.svc.cluster.local`，其中 `<pod-name>` 是 Pod 的名称，而对于 StatefulSet 管理的 Pod，名称是稳定的，确保 DNS 名称也稳定。
> >
> > #### 3. 持久化存储yaml
> >
> > - **VolumeClaimTemplates**：
> >
> >   ```yaml
> >   volumeClaimTemplates:
> >   - metadata:
> >       name: data
> >     spec:
> >       accessModes: [ "ReadWriteOnce" ]
> >       storageClassName: "ssd"
> >       resources:
> >         requests:
> >           storage: 10Gi
> >   ```
> >
> > - **存储绑定规则**：
> >
> >   - 每个 Pod 获得独立的 PVC（如 data-web-0, data-web-1）
> >   - Pod 重建后仍绑定原有 PVC
> >
> > ### 4. PV PVC 行为
> >
> > > **一、PV 创建方式**：
> > >
> > > 1. **静态创建 PV：**
> > >
> > >    - 在静态绑定模式下，管理员需要事先创建 PV，并且手动配置好存储的详细信息（例如存储的大小、访问模式、存储类等）。当应用程序创建 PVC 时，Kubernetes 会寻找与该 PVC 匹配的 PV，并将它们绑定在一起。
> > >
> > >    **静态 PV 示例：**
> > >
> > >    ```
> > >    apiVersion: v1
> > >    kind: PersistentVolume
> > >    metadata:
> > >      name: pv-storage
> > >    spec:
> > >      capacity:
> > >        storage: 1Gi
> > >      volumeMode: Filesystem
> > >      accessModes:
> > >        - ReadWriteOnce
> > >      persistentVolumeReclaimPolicy: Retain
> > >      storageClassName: standard
> > >      hostPath:
> > >        path: /mnt/data
> > >    ```
> > >
> > >    在此示例中：
> > >
> > >    - `capacity` 定义了存储空间大小。
> > >    - `accessModes` 定义了访问模式，如 `ReadWriteOnce`（只能被单个节点挂载并读写）。
> > >    - `hostPath` 指定了该 PV 的实际存储位置（这里是集群节点上的 `/mnt/data`）。
> > >    - `persistentVolumeReclaimPolicy: Retain` 表示当 PVC 被删除时，PV 仍然保留，必须手动删除。
> > >
> > > 2. **动态创建 PV（使用 StorageClass）：**
> > >
> > >    - 在动态绑定模式下，**StorageClass** 会用来自动创建 PV。当 PVC 被创建时，Kubernetes 会根据 PVC 中的 `storageClassName` 来查找并选择匹配的 StorageClass，然后根据该 StorageClass 自动创建一个 PV。
> > >
> > >    **动态 PV 示例：**
> > >
> > >    ```
> > >    apiVersion: v1
> > >    kind: PersistentVolumeClaim
> > >    metadata:
> > >      name: pvc-storage
> > >    spec:
> > >      accessModes:
> > >        - ReadWriteOnce
> > >      resources:
> > >        requests:
> > >          storage: 1Gi
> > >      storageClassName: standard
> > >    ```
> > >
> > >    **StorageClass 示例：**
> > >
> > >    ```
> > >    apiVersion: storage.k8s.io/v1
> > >    kind: StorageClass
> > >    metadata:
> > >      name: standard
> > >    provisioner: kubernetes.io/gce-pd  # 对于 GCP 来说，自动创建 GCE PersistentDisk
> > >    parameters:
> > >      type: pd-standard
> > >    ```
> > >
> > >    在这个示例中，创建的 PVC 使用了 `storageClassName: standard`，Kubernetes 会根据该 StorageClass 查找并自动创建符合条件的 PV。
> > >
> > > **二、PV 与 PVC 的绑定过程：****
> > >
> > > 1. **PV 创建**：管理员或系统通过静态或动态方式创建 PV。
> > > 2. **PVC 创建**：应用程序或用户创建 PVC，指定所需的存储资源（如大小、访问模式、存储类等）。
> > > 3. **PV 和 PVC 绑定**：Kubernetes 自动或手动将 PVC 与符合要求的 PV 进行绑定。绑定后，PVC 和 PV 会共享同一个存储卷。
> > > 4. **使用 PV**：Pod 通过 PVC 使用存储卷，而无需直接访问 PV。
> > >
> > > **三、关于 StatefulSet 和 PV/PVC：**
> > >
> > > 在 `StatefulSet` 中使用 `VolumeClaimTemplates` 时，Kubernetes 会根据模板自动为每个 Pod 创建对应的 **PVC**。这些 PVC 会绑定到 **PV** 上（如果有可用的 PV），或者通过 **StorageClass** 动态创建 PV。
> > >
> > > - **如果使用动态卷**（通过 StorageClass），`StatefulSet` 中的每个 Pod 会自动与一个新的 PV 绑定。
> > > - **如果使用静态卷**，你需要确保已预先创建的 PV 与每个 PVC 匹配，然后 Kubernetes 会将 PVC 与相应的 PV 绑定。
> > >
> > > **四、pv pvc删除行为**
> > >
> > > 当一个 PVC 被删除时，它不再与任何 Pod 相关联，但 PV 的行为依赖于其 **回收策略**。下面是根据不同回收策略下，PVC 被删除后的数据处理情况：
> > >
> > > **1. ReclaimPolicy: Retain**
> > >
> > > - **数据保留**：当 PVC 被删除后，PV 仍然保留在集群中，数据保持不变。
> > >
> > > - **手动处理**：你必须手动清理数据或删除 PV，直到 PV 被重新配置或删除为止。它不会自动清除存储卷中的数据。
> > >
> > > - **场景**：适用于希望手动干预存储数据或保留数据用于后续用途的情况。
> > >
> > >   **举个例子**：如果你有一个数据库存储卷，当 PVC 被删除时，你可能希望保留该数据，以便稍后手动备份或恢复。
> > >
> > > **2. ReclaimPolicy: Delete**
> > >
> > > - **数据删除**：当 PVC 被删除时，PV 以及关联的存储卷（如云提供商的 EBS 卷、GCE Persistent Disk 等）会自动删除，存储中的数据会被销毁。
> > >
> > > - **自动清理**：对于动态创建的存储卷（例如使用云存储提供商的卷），Kubernetes 会在 PVC 删除时自动销毁 PV 和存储卷。对于静态创建的 PV（如 `hostPath`），如果 PV 被删除，数据也会随之被删除。
> > >
> > > - **场景**：适用于存储数据不需要保留，且 PVC 删除时希望自动回收存储资源的情况。
> > >
> > >   **举个例子**：如果你使用的是 `AWS EBS` 卷并设置了 `Delete` 回收策略，那么当 PVC 删除时，关联的 EBS 卷会被销毁，数据会丢失。
> > >
> > > **3. ReclaimPolicy: Recycle（已废弃）**
> > >
> > > - **数据清空**：这个策略会删除卷中的数据，通常执行 `rm -rf /` 操作清空卷，然后将 PV 状态改为 `Available`，可以再次被绑定到新的 PVC 上。
> > > - **场景**：该策略已废弃，不建议使用，但如果依然存在，可以假设它在 PVC 删除时清空数据并准备重新使用。
> > >
> > > **4**.**PVC 删除后，PV 上的数据会怎么样？**
> > >
> > > - **当 PVC 删除时，具体的数据处理取决于 PV 的 `ReclaimPolicy`**：
> > >   - 如果是 `Retain`，数据保留，PV 不会被删除，必须手动处理。
> > >   - 如果是 `Delete`，数据会被删除，PV 和底层存储卷（如云存储）也会被销毁。
> > >   - `Recycle`（已废弃）会清空数据并准备卷以供再次使用。
> > >
> > > 5. **动态卷和静态卷的区别：**
> > >
> > > 1. **动态 PV**（由 `StorageClass` 自动创建）：
> > >    - 如果 `ReclaimPolicy` 为 `Delete`，那么动态创建的卷会在 PVC 删除后被自动销毁。
> > >    - 如果 `ReclaimPolicy` 为 `Retain`，则存储卷会保持存在，但需要手动清理数据。
> > > 2. **静态 PV**（管理员手动创建）：
> > >    - 如果 `ReclaimPolicy` 为 `Delete`，则管理员需要手动删除 PV，因为静态 PV 不会自动销毁。
> > >    - 如果 `ReclaimPolicy` 为 `Retain`，数据会保留，但需要手动清理。
> >
> > 
> >
> > #### 二、高级特性
> >
> > #### 1. 拓扑约束
> >
> > ```yaml
> > spec:
> >   podManagementPolicy: Parallel # 可改为并行创建
> >   serviceName: "nginx"
> >   replicas: 3
> >   template:
> >     spec:
> >       affinity:		# pod 亲和性
> >         podAntiAffinity:
> >           requiredDuringSchedulingIgnoredDuringExecution:
> >           - labelSelector:
> >               matchExpressions:
> >               - key: app
> >                 operator: In
> >                 values: ["nginx"]
> >             topologyKey: "kubernetes.io/hostname"
> > ```
> >
> > #### 2. 更新策略对比
> >
> > | 策略类型      | 行为特点                        | 适用场景           |
> > | :------------ | :------------------------------ | :----------------- |
> > | RollingUpdate | 按序逐步更新（可设置partition） | 生产环境推荐       |
> > | OnDelete      | 手动删除Pod触发更新             | 需要精确控制更新时 |
> >
> > #### 3. 特殊字段说明
> >
> > - **podManagementPolicy**：
> >   - `OrderedReady`（默认）：顺序创建
> >   - `Parallel`：并行创建（仍保证网络标识唯一性）



>  高级绑定机制
>
> ### 1. 选择器(Selector)匹配
>
> PV可设置标签选择器限制绑定：
>
> ```
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: pv-with-selector
>   labels:
>     tier: ssd
> spec:
>   selector:
>     matchLabels:
>       usage: "production"  # 只绑定有相同标签的PVC
>   ...
> ```
>
> ### 2. 卷模式(volumeMode)匹配
>
> 必须精确匹配文件系统(Filesystem)或原始块设备(Block)：
>
> ```
> # PV配置
> volumeMode: Block
> 
> # 匹配的PVC必须相同
> volumeMode: Block
> ```
>
> ### 3. 节点亲和性限制
>
> 本地存储PV需通过节点亲和性约束：
>
> ```
> apiVersion: v1
> kind: PersistentVolume
> metadata:
>   name: local-pv
> spec:
>   nodeAffinity:
>     required:
>       nodeSelectorTerms:
>       - matchExpressions:
>         - key: kubernetes.io/hostname
>           operator: In
>           values: ["node-01"]
> ```





### Storage Class

1. **基本定义**：StorageClass 是 Kubernetes 中的一种资源对象，用于管理持久卷(Persistent Volumes, PV)的动态供给。
2. **核心功能**：
   - 定义动态供给策略：当应用程序请求持久卷声明(Persistent Volume Claim, PVC)时，Kubernetes 可以根据 StorageClass 自动创建合适的 PV
   - 允许管理员定义多种存储类型：可以为不同的存储后端(如 AWS EBS、Azure Disk、NFS 等)创建不同的 StorageClass
3. **优势**：
   - 使存储管理更加灵活：不同应用可以根据需求选择不同类型的存储
   - 实现存储供给自动化：无需管理员手动创建 PV，系统可以按需自动创建

StorageClass 是现代 Kubernetes 存储管理的核心组件之一，特别适合云环境下的动态存储需求。



## 调度器

### 基本概念

1. **调度器（Scheduler）的核心功能**：

   - 作为独立进程运行，持续监听 Kubernetes API Server
   - 筛选待调度 Pod：识别 `spec.nodeName` 为空的 Pod（即未分配节点的 Pod）
   - 创建绑定（Binding）：决定 Pod 应该运行在哪个节点，并通过 API Server 写入绑定信息

2. **调度过程的四个关键阶段**：

   ![deepseek_mermaid_20250817_080da4](C:\Users\xlwwwww\Downloads\deepseek_mermaid_20250817_080da4.svg)

3. **调度器设计的关键考量因素**：

   | 维度           | 说明               | 技术实现示例                                                 |
   | :------------- | :----------------- | :----------------------------------------------------------- |
   | **公平性**     | 避免节点资源饥饿   | 轮询调度、资源配额平衡                                       |
   | **资源利用率** | 最大化集群资源使用 | 装箱算法（Bin Packing）、碎片整理                            |
   | **效率**       | 快速调度大批量 Pod | 并发调度、优先级队列                                         |
   | **灵活性**     | 支持自定义调度逻辑 | 调度器扩展（Scheduler Extender）、调度框架（Scheduling Framework） |

------

### 深度解析

1. **调度决策流程**：

   - **过滤（Filtering）**：排除不满足 Pod 需求的节点（如资源不足、标签不匹配）

     > #### **1. PodFitsResources**
     >
     > **功能**：检查节点剩余资源是否满足 Pod 请求
     > **实现逻辑**：
     >
     > go
     >
     > ```
     > if node.Allocatable.Memory < pod.Requests.Memory || node.Allocatable.CPU < pod.Requests.CPU {
     >     return false // 节点资源不足
     > }
     > ```
     >
     > **关键点**：
     >
     > - 比较节点可分配资源（`allocatable`）与 Pod 的 `requests`
     > - 涉及 CPU、Memory、GPU 等资源类型
     >
     > ------
     >
     > #### **2. PodFitsHost**
     >
     > **功能**：校验 Pod 指定的节点名称
     > **触发条件**：当 Pod 显式设置 `spec.nodeName` 时
     > **示例**：
     >
     > yaml
     >
     > ```
     > spec:
     >   nodeName: node-01 # 必须调度到指定节点
     > ```
     >
     > ------
     >
     > #### **3. PodFitsHostPorts**
     >
     > **功能**：检查节点端口冲突
     > **规则**：
     >
     > - 节点上已占用的 HostPort 不能与 Pod 申请的端口重复
     > - 仅针对使用 `hostPort` 的 Pod
     >   **冲突示例**：
     >
     > yaml
     >
     > ```
     > # Pod A 已占用 80 端口
     > containers:
     >   - ports:
     >     - hostPort: 80
     > 
     > # Pod B 再申请 80 端口时会触发过滤
     > ```
     >
     > ------
     >
     > #### **4. PodSelectorMatches**
     >
     > **功能**：节点标签匹配校验
     > **应用场景**：
     >
     > - 当 Pod 设置 `nodeSelector` 或 `nodeAffinity` 时
     >   **示例**：
     >
     > yaml
     >
     > ```
     > spec:
     >   nodeSelector:
     >     disktype: ssd # 只选择有 disktype=ssd 标签的节点
     > ```
     >
     > ------
     >
     > #### **5. NoDiskConflict**
     >
     > **功能**：检查存储卷冲突
     > **规则**：
     >
     > - 如果节点已挂载的卷与 Pod 请求的卷冲突（读写模式冲突），则过滤该节点
     >
     > - 例外：所有卷均为只读时允许共存
     >   **冲突类型**：
     >   | 已挂载卷模式 | Pod 请求卷模式 | 是否冲突 |
     >   |--------------|----------------|----------|
     >   | RW | RW | 是 |
     >   | RW | RO | 是 |
     >   | RO | RO | 否 |

   - **评分（Scoring）**：对符合要求的节点评分（如资源空闲率、亲和性权重）

     在预选阶段的没有合适的节点时，pod就会一直的处于pending阶段，不断重试调试，直至所有节点满足条件，如果有多个节点满足条件，就会继续优选，按照优先级大小进行节点排序

   - **绑定（Binding）**：选择最高分节点，通过 `POST /bindings` API 创建绑定

2. **实际挑战与解决方案**：

   - **大规模集群调度**：采用分片调度器（如 Uber 的 Peloton）
   - **特殊调度需求**：通过调度器扩展实现 GPU/异构资源调度
   - **敏感型应用**：使用 Pod 亲和性/反亲和性规则控制拓扑分布

3. **调度器演进**：

   - **Kubernetes 调度框架**：提供插件化架构，支持自定义扩展点

> **注**：现代 Kubernetes 支持多调度器并行工作，可通过 `spec.schedulerName` 为 Pod 指定特定调度器（如 `my-custom-scheduler`），实现租户隔离或特殊业务调度需求。

创建deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  labels:
    app: myapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      schedulerName: my-scheduler  # 使用自定义调度器
      containers:
        - name: myapp
          image: swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/wangyanglinux/myapp:v1.0
```

创建自己的调度器

```bash
#!/bin/bash

SERVER='localhost:8001'

while true; do
    # 获取所有未调度的 Pod（使用 my-scheduler 且未分配节点）
    for PODNAME in $(kubectl --server $SERVER get pods -o json | \
                    jq -r '.items[] |
                          select(.spec.schedulerName == "my-scheduler") |
                          select(.spec.nodeName == null) |
                          .metadata.name')
    do
        # 获取所有可用节点
        NODES=($(kubectl --server $SERVER get nodes -o json | \
               jq -r '.items[].metadata.name'))
        NUMNODES=${#NODES[@]}

        # 随机选择一个节点
        CHOSEN=${NODES[$((RANDOM % NUMNODES))]}

        # 创建绑定请求
        curl --header "Content-Type: application/json" \
             --request POST \
             --data '{
                "apiVersion": "v1",
                "kind": "Binding",
                "metadata": {
                    "name": "'$PODNAME'"
                },
                "target": {
                    "apiVersion": "v1",
                    "kind": "Node",
                    "name": "'$CHOSEN'"
                }
             }' \
             "http://$SERVER/api/v1/namespaces/default/pods/$PODNAME/binding/"

        echo "[$(date)] Assigned $PODNAME to $CHOSEN"
    done

    sleep 1
done
```



### 亲和性

亲和性策略是 Kubernetes 调度系统的核心高级功能，允许精细控制 Pod 的调度位置。主要分为两类：

------

#### **一、节点亲和性 (Node Affinity)**

控制 Pod 与节点的匹配关系，替代传统的 `nodeSelector` 并提供更强大的表达式能力。

**`nodeAntiAffinity`**  **`nodeAffinity`**  都是存在的

**1. 硬性要求 (`requiredDuringSchedulingIgnoredDuringExecution`)**
必须满足的条件，否则 Pod 无法调度：

```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
      - matchExpressions:
        - key: gpu-type
          operator: In
          values: [a100, v100]
        - key: storage-tier
          operator: NotIn
          values: [hdd]
```

**2. 软性偏好 (`preferredDuringSchedulingIgnoredDuringExecution`)**
优先满足的条件，但不强制

```yaml
preferredDuringSchedulingIgnoredDuringExecution:
- weight: 80  # 权重值 (1-100)
  preference:
    matchExpressions:
    - key: zone
      operator: In
      values: [us-east-1a]
```

**操作符支持**：

| 操作符         | 功能描述          | 示例                  |
| :------------- | :---------------- | :-------------------- |
| `In`           | 标签值在集合内    | `values: [ssd, nvme]` |
| `NotIn`        | 标签值不在集合内  | `values: [slow]`      |
| `Exists`       | 标签键存在        | 不指定 `values`       |
| `DoesNotExist` | 标签键不存在      | 不指定 `values`       |
| `Gt`/`Lt`      | 数值比较 (> 或 <) | `values: ["2"]`       |

------

#### **二、Pod 间亲和性 (Inter-Pod Affinity/Anti-Affinity)**

控制 Pod 之间的分布关系，支持拓扑域约束。

**1. Pod 亲和性 (Affinity)**

同样这里的也存在**硬策略**和**软策略**

将 Pod 调度到已有同类 Pod 的节点：

```yaml
podAffinity:
  requiredDuringSchedulingIgnoredDuringExecution:
  - labelSelector:
      matchExpressions:
      - key: app
        operator: In
        values: [cache]
    topologyKey: rack  # 拓扑域定义 
```

**2. Pod 反亲和性 (Anti-Affinity)**
避免 Pod 集中部署：

```yaml
podAntiAffinity:
  preferredDuringSchedulingIgnoredDuringExecution:
  - weight: 100
    podAffinityTerm:
      labelSelector:
        matchExpressions:
        - key: app
          operator: In
          values: [web]
      topologyKey: zone
```

> ### **1. nodeAffinity (节点亲和性)**
>
> | 特性         | 说明                                                         |
> | :----------- | :----------------------------------------------------------- |
> | **匹配标签** | 作用于**节点(Node)的标签**                                   |
> | **操作符**   | `In`, `NotIn`, `Exists`, `DoesNotExist`, `Gt`(大于), `Lt`(小于) |
> | **拓扑域**   | ❌ 不支持                                                     |
> | **调度目标** | 将 Pod 调度到**特定标签的节点**上                            |
> | **典型场景** | • 将 GPU 负载调度到带 `accelerator=gpu` 标签的节点 • 避免调度到标记为 `disk=slow` 的节点 |
>
> ------
>
> ### **2. podAffinity (Pod 亲和性)**
>
> | 特性         | 说明                                                         |
> | :----------- | :----------------------------------------------------------- |
> | **匹配标签** | 作用于**其他 Pod 的标签**                                    |
> | **操作符**   | `In`, `NotIn`, `Exists`, `DoesNotExist`                      |
> | **拓扑域**   | ✅ 支持（通过 `topologyKey` 定义域，如 `kubernetes.io/hostname`） |
> | **调度目标** | 将 Pod 调度到**与指定 Pod 在同一拓扑域**的节点上             |
> | **典型场景** | • 前端 Pod 与后端 Pod 部署在同一可用区（`topologyKey: topology.kubernetes.io/zone`） • 微服务紧密耦合的组件部署在同一主机 |
>
> ------
>
> ### **3. podAntiAffinity (Pod 反亲和性)**
>
> | 特性         | 说明                                                         |
> | :----------- | :----------------------------------------------------------- |
> | **匹配标签** | 作用于**其他 Pod 的标签**                                    |
> | **操作符**   | `In`, `NotIn`, `Exists`, `DoesNotExist`                      |
> | **拓扑域**   | ✅ 支持（通过 `topologyKey` 定义域）                          |
> | **调度目标** | 将 Pod 调度到**与指定 Pod 不在同一拓扑域**的节点上           |
> | **典型场景** | • 高可用：同一服务的多个副本分散在不同节点/可用区 • 避免资源竞争：数据库主从节点分离部署 |

### 污点和容忍

**核心概念**

| **概念**                 | **作用对象** | **目的**                                                     |
| :----------------------- | :----------- | :----------------------------------------------------------- |
| **Taint（污点）**        | 节点 (Node)  | 标记节点特性，**排斥**不能容忍这些污点的 Pod                 |
| **Toleration（容忍度）** | Pod          | 允许 Pod **接受**特定污点，使其可以被调度到有匹配污点的节点上 |

------

**工作机制**

1. **节点打污点**
   - 管理员给节点添加一个或多个污点（例如：`disk=slow:NoSchedule`）
   - 每个污点包含三个部分：`Key=Value:Effect`
2. **Pod 设置容忍度**
   - 在 Pod 配置中声明它能容忍哪些污点（例如：容忍 `disk=slow`）
3. **调度器决策**
   - ❌ 如果 Pod **没有**容忍节点的污点 → Pod **不会被调度**到该节点
   - ✅ 如果 Pod **有**容忍节点的污点 → Pod **可能被调度**到该节点

------

**污点效果（Effect）类型**

| **效果**           | **行为**                                                     |
| :----------------- | :----------------------------------------------------------- |
| `NoSchedule`       | 新 Pod 不会调度到此节点（已运行的 Pod 不受影响）             |
| `PreferNoSchedule` | 尽量避免调度新 Pod 到此节点（非强制）                        |
| `NoExecute`        | 新 Pod 不会调度到此节点，且**已运行但不容忍的 Pod 会被驱逐** |

------

**关键特性**

1. **容忍不是偏好**

   - 容忍度仅表示 Pod **能接受**污点，不表示 Pod **偏好**该节点
   - 调度器仍可能优先选择无污点节点

2. **驱逐机制（NoExecute）**

   bash

   ```
   # 添加会驱逐 Pod 的污点
   kubectl taint nodes node1 key1=value1:NoExecute
   ```

   - 已运行但**不容忍**此污点的 Pod 会被**立即驱逐**
   - 已运行且**容忍**此污点的 Pod 保持运行

3. **系统默认污点**
   Kubernetes 自动添加的污点：

   - `node.kubernetes.io/not-ready`（节点未就绪）
   - `node.kubernetes.io/unreachable`（节点失联）
   - `node.kubernetes.io/disk-pressure`（磁盘压力大）

```bash
kubectl taint nodes [node-name] key=val:effect
kubectl taint nodes [node-name] key=val-   # -表示移除
```



4. 容忍类型

> 在 Kubernetes 中，Pod 的容忍度（Toleration）有以下几种主要类型，主要通过 `operator` 和 `effect` 字段来定义：
>
> ------
>
> ### **1. 按操作符（Operator）分类**
>
> #### (1) **`Equal`（相等匹配）**
>
> - **要求**：Pod 容忍度的 `key`、`value` 和 `effect` 必须与节点污点**完全匹配**
>
> - **示例**：
>
>   ```yaml
>   tolerations:
>   - key: "disk"           # 污点的 key
>     operator: "Equal"     # 操作符
>     value: "ssd"          # 必须等于 ssd
>     effect: "NoSchedule"  # 污点效果
>   ```
>
> #### (2) **`Exists`（存在性匹配）**
>
> - **要求**：只需污点的 `key` 和 `effect` 存在，不检查 `value`（此时不能指定 `value` 字段）
>
> - **示例**：
>
>   yaml
>
>   ```yaml
>   tolerations:
>   - key: "gpu"            # 污点的 key
>     operator: "Exists"    # 操作符（忽略 value）
>     effect: "NoSchedule"  # 污点效果
>   ```
>
> ------
>
> ### **2. 按污点效果（Effect）分类**
>
> #### (1) **`NoSchedule`（禁止调度）**
>
> - **行为**：Pod 可以调度到有匹配污点的节点，但**不容忍的 Pod 不会被调度**
>
> - **特点**：不影响已运行的 Pod
>
> - **示例**：
>
>   ```yaml
>   tolerations:
>   - key: "env"
>     operator: "Equal"
>     value: "prod"
>     effect: "NoSchedule"  # 仅影响新调度
>   ```
>
> #### (2) **`PreferNoSchedule`（尽量避免调度）**
>
> - **行为**：调度器**尽量不**将 Pod 调度到该节点，但非强制
>
> - **特点**：软性约束，优先级低于资源请求等条件
>
> - **示例**：
>
>   ```yaml
>   tolerations:
>   - key: "load"
>     operator: "Exists"
>     effect: "PreferNoSchedule"  # 非强制
>   ```
>
> #### (3) **`NoExecute`（执行驱逐）**
>
> - **行为**：
>
>   - 新 Pod：不容忍则不会被调度
>   - 已运行 Pod：不容忍则会被**驱逐**
>
> - **特殊字段**：`tolerationSeconds`（容忍时间）
>
>   ```yaml
>   tolerations:
>   - key: "maintenance"
>     operator: "Equal"
>     value: "true"
>     effect: "NoExecute"
>     tolerationSeconds: 3600  # 被驱逐前可运行 1 小时
>   ```
>
> ------
>
> ### **3. 特殊容忍类型**
>
> #### (1) **容忍所有污点**
>
> ```yaml
> tolerations:
> - operator: "Exists"  # 容忍所有 key 和 effect
> ```
>
> #### (2) **容忍特定效果的所有污点**
>
> yaml
>
> ```yaml
> tolerations:
> - operator: "Exists"  # 容忍所有 key
>   effect: "NoExecute" # 但仅匹配此效果
> ```
>
> #### (3) **空 key 容忍（Kubernetes 1.28+）**
>
> ```yaml
> tolerations:
> - effect: "NoExecute" # 容忍所有 NoExecute 污点
>   operator: "Exists"   # 无需指定 key
> ```



### 固定节点调度

#### spec.nodeName

直接将pod调度到的指定node节点，会跳过Scheduler的调度策略，该匹配是强制匹配

1. **强制节点调度 (nodeName)**

   - 使用 `spec.template.spec.nodeName: k8s-node01` 直接将所有 Pod 固定调度到指定节点
   - **绕过调度器**：跳过 Kubernetes Scheduler 的正常调度流程
   - **强制匹配**：**无视节点资源状态/污点等限制**，若节点不存在则 Pod 无法启动

2. 配置

   ```yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: nodename-test
   spec:
     replicas: 7
     selector:
       matchLabels:
         app: nodename
     template:
       metadata:
         labels:
           app: nodename
       spec:
         nodeName: k8s-node01  # 节点选择
         containers:
         - name: myweb
           image: wanoyanglinux/myapp:v1.0
           ports:
           - containerPort: 80
   ```



#### spec.nodeSelector

##### 📌 **nodeSelector 核心概念**

| **特性**     | **说明**                                                     |
| :----------- | :----------------------------------------------------------- |
| **作用**     | 通过节点标签(Label)强制将 Pod 调度到特定节点                 |
| **匹配机制** | 使用 Kubernetes 的 `label-selector` 机制                     |
| **约束类型** | **强制约束** - 调度器必须满足条件才会调度                    |
| **优先级**   | 高于亲和性(affinity)规则，低于污点(taint)规则 即还受到污点的影响 |
| **适用场景** | 简单节点分配需求（如专用硬件、特定区域节点）                 |

------

##### 🛠️ **YAML 配置解析**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nodeselect-test    # 部署名称
spec:
  replicas: 2              # 创建2个副本
  selector:                # 标签选择器（管理Pod）
    matchLabels:
      app: nodeselect      # 匹配 app=nodeselect 标签的Pod
  template:                # Pod模板
    metadata:
      labels:
        app: nodeselect    # 给Pod打上 app=nodeselect 标签
    spec:
      nodeSelector:       # ⭐ 节点选择器（关键配置）
        type: nodeselect   # 要求节点有 type=nodeselect 标签
      containers:
      - name: myweb        # 容器名称
        image: wangyanglinux/myapp:v1.0  # 镜像
        ports:
        - containerPort: 80 # 容器端口
```



## k8s鉴权

**认证 -----》 鉴权  -----》 准入**

### 认证

1. **Kubernetes 组件对 API Server 的访问**
   - **访问者**: Kubernetes 的核心控制平面组件和工作节点组件，包括：
     - `kubectl`： (命令行工具) 用户通过它向 API Server 发送指令。
     - `Controller Manager`： (控制器管理器) 负责维护集群的状态（如副本数），需要持续与 API Server 通信。
     - `Scheduler`： (调度器) 负责为 Pod 分派节点，需要从 API Server 获取未调度的 Pod 信息，并回写调度决策。
     - `kubelet`： (节点代理) 运行在每个工作节点上，负责维护节点上的 Pod 生命周期，需要向 API Server 报告节点和 Pod 状态。
     - `kube-proxy`： (网络代理) 负责节点上的网络规则，实现服务访问，需要从 API Server 获取服务信息。
   - **被访问者**: `API Server`，它是 Kubernetes 集群所有 REST 操作的唯一入口，是整个系统的“大脑”，负责验证请求、处理数据、更新存储（etcd）等。所有组件都需要通过它与集群交互。
2. **Kubernetes 管理的 Pod 对容器的访问**
   - **访问者**: `Pod`。Pod 是 Kubernetes 中最小的可部署计算单元，一个 Pod 内可以运行一个或多个容器。
   - **被访问者**: Pod 内部运行的`容器`。
   - **补充说明**: 文本特别指出 `dashboard` (Kubernetes Web UI) 也是以一个 Pod 的形式运行的。这意味着 Dashboard 服务本身也需要访问其自身容器内的进程来提供 Web 界面，同时它作为一个 Pod，也受 Kubernetes 管理（例如被调度、被 kubelet 监管）。



##### 访问方式

1. **本地非安全端口访问 (Local Insecure Access)**
   - **组件**: `Controller Manager` (控制器管理器), `Scheduler` (调度器)
   - **方式**: 直接使用 API Server 的**非安全端口**进行 HTTP 访问。
   - **原因**: 因为这些核心控制平面组件通常与 `API Server` **部署在同一台物理机或虚拟机（Master 节点）上**。它们之间的通信被视为“本地”或“可信”的通信，不走公网。
   - **技术实现**: API Server 通过启动参数 `--insecure-bind-address=127.0.0.1` 来监听本地回环地址（localhost）上的非加密 HTTP 连接。这是一种简化部署的方式，牺牲了本地回路间的加密以换取性能。
2. **基于证书的 HTTPS 双向认证 (Certificate-based HTTPS Mutual Authentication)**
   - **组件**: `kubectl` (命令行工具), `kubelet` (节点代理), `kube-proxy` (网络代理)
   - **方式**: 必须使用 **TLS 证书**进行 **HTTPS 双向认证**。
   - **原因**: 这些组件通常**不**与 API Server 在同一台机器上。
     - `kubectl` 可能运行在任意一台可以访问集群的运维机器上。
     - `kubelet` 和 `kube-proxy` 运行在每个工作节点（Node）上。
   - 它们与 API Server 之间的网络通信需要穿越节点间网络，被认为是**不可信的**。因此必须使用严格的加密和身份验证机制来保证通信的安全性和可靠性。
   - **技术实现**:
     - **HTTPS**: 通信通道进行加密，防止窃听。
     - **双向认证 (mTLS)**: 不仅客户端（如 kubelet）需要验证服务器（API Server）的证书以确保连接到了真正的 API Server，API Server 也需要验证客户端的证书以确认其身份和权限（例如，确认这是某个合法的节点 kubelet）。



##### 证书颁发方式

1. **手动签发 (Manual Issuance)**
   - **方式**: 由集群管理员**手动**使用 Kubernetes 集群的**根证书颁发机构 (Root CA)** 为客户端（如 **`kubelet`**, **`kube-proxy`**, 或者用户使用的 `kubectl`）生成 HTTPS 客户端证书。
   - **流程**: 管理员需要为每个节点或每个用户执行以下操作：
     1. 为该实体生成一个私钥和证书签名请求 (CSR)。
     2. 使用集群的 CA 证书和私钥来签署该 CSR，生成客户端的证书。
     3. 将生成的证书和私钥分发到对应的节点或用户机器上，并配置组件使用这些文件。
   - **特点**: 操作繁琐，不易于大规模集群的管理，尤其是在节点需要动态扩缩容时。
2. **自动签发 (Automatic Issuance - TLS Bootstrapping)**
   - **方式**: 这是一个自动化的流程，通常特指 **`kubelet`** 的初始证书获取过程。
   - **流程**:
     - **首次访问 (Bootstrap)**: 当一个新节点上的 **`kubelet`** 第一次启动时，它还没有合法的客户端证书。此时，它会使用一个预先配置好的**引导令牌 (Bootstrap Token)** 作为身份凭证去访问 API Server。
     - **认证与签发**: API Server 验证该令牌的合法性。验证通过后，`kubelet` 可以发起一个证书签名请求 (CSR)。
     - **审批请求**: （通常由 `Controller Manager` 中的 `csrapproving` 控制器）自动批准这个 CSR 请求。
     - **获取证书**: CSR 被批准后，`kubelet` 从 API Server 获取为其签发的有效客户端证书和私钥，并将其保存在本地。
     - **后续访问**: 从此以后，`kubelet` 就使用这个自动生成的、合法的证书来进行 HTTPS 双向认证，访问 API Server。引导令牌仅在第一次时使用。
   - **特点**: 极大简化了集群节点的扩容工作，实现了节点身份的自动化管理，是生产环境的推荐做法。、



##### **kubeconfig**

kubeconfig 文件是一个 **YAML 或 JSON 格式的配置文件**，其主要作用是**将关于 Kubernetes 集群、用户和上下文（Context）的信息组织在一起，并提供给 `kubectl` 等客户端工具使用**，以便客户端知道如何找到并认证到一个或多个 Kubernetes 集群。

简单来说，它就是 `kubectl` 的“连接指南”和“通行证”。

> ### kubeconfig 文件的核心结构
>
> 一个典型的 kubeconfig 文件包含以下几部分：
>
> 1. **Clusters (集群)**
>
>    - 定义你可以访问的 Kubernetes 集群。
>
>    - 主要信息：**集群的 API Server 地址** 和 **用于验证服务器证书的 CA 证书**。
>
>    - 示例：
>
>      ```yaml
>      clusters:
>      - name: my-production-cluster
>        cluster:
>          server: https://k8s-api.example.com:6443
>          certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0t... (Base64 编码的 CA 证书内容)
>      ```
>
> 2. **Users (用户/凭证)**
>
>    - 定义用于认证到集群的身份凭证。
>
>    - 可以是多种认证方式：**客户端证书**、** bearer token**、**用户名密码**等。
>
>    - 示例（客户端证书方式）：
>
>      ```yaml
>      users:
>      - name: admin-user
>        user:
>          client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0t... (Base64 编码的客户端证书)
>          client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQ... (Base64 编码的客户端私钥)
>      ```
>
> 3. **Contexts (上下文)**
>
>    - **将 Cluster 和 User 捆绑在一起**。一个上下文回答了“*用哪个用户身份（User）去访问哪个集群（Cluster）？*”这个问题。
>
>    - 还可以设置默认的命名空间 (Namespace)。
>
>    - 示例：
>
>      ```yaml
>      contexts:
>      - name: prod-admin@my-production-cluster
>        context:
>          cluster: my-production-cluster
>          user: admin-user
>          namespace: default # (可选) 设置默认命名空间
>      ```
>
> 4. **Current Context (当前上下文)**
>
>    - 指定当前**正在使用**的上下文是哪一个。`kubectl` 的所有操作都会基于这个当前上下文指向的集群和用户。
>
>    - 示例：
>
>      ```yaml
>      current-context: prod-admin@my-production-cluster
>      ```
>
> ------
>
> ### kubeconfig 文件如何工作？
>
> 1. 当你执行 `kubectl get pods` 时，`kubectl` 会默认从 `~/.kube/config` 路径加载 kubeconfig 文件。
> 2. `kubectl` 查看 `current-context` 字段，找到当前使用的上下文（例如 `prod-admin@my-production-cluster`）。
> 3. 根据该上下文的定义，它知道要去 `my-production-cluster` 这个集群，并使用 `admin-user` 这个用户的凭证。
> 4. 它从 `clusters` 列表中找到 `my-production-cluster` 的地址 (`server`) 和 CA 证书。
> 5. 它从 `users` 列表中找到 `admin-user` 的客户端证书和私钥。
> 6. 最后，`kubectl` 使用找到的地址、CA 证书、用户证书和私钥，构建一个安全的 HTTPS 请求发送给 API Server，并进行双向认证。
>
> ------
>
> ### 管理与使用
>
> - **默认位置**: `$HOME/.kube/config`
>
> - **指定其他文件**: 可以使用 `--kubeconfig` 参数指定其他配置文件。
>
>   ```yaml
>   kubectl --kubeconfig=/path/to/another/config get pods
>   ```
>
> - **环境变量**: 设置 `KUBECONFIG` 环境变量可以覆盖默认的查找路径。
>
> - **合并配置**: 你可以有多个 kubeconfig 文件（例如，用于开发、测试、生产不同集群），并通过设置 `KUBECONFIG` 变量将它们合并使用。
>
>   ```yaml
>   export KUBECONFIG=~/.kube/config:~/.kube/config-prod
>   ```
>
> - **切换上下文**: 使用 `kubectl config use-context <context-name>` 命令来切换当前上下文，从而快速在不同集群和用户身份间切换。





#### ServiceAccount

##### 基本概念

1. **问题陈述：Pod 动态性的挑战**
   - Pod 是 Kubernetes 中最小的部署单元，其生命周期是**动态的** - 它们会被频繁地创建、销毁和重启。
   - 如果采用为每个 Pod **手动签发证书** 的方式（就像为静态组件 kubelet 所做的那样），在管理上将是一场噩梦。证书的创建、分发和轮换无法跟上 Pod 快速变化的节奏。
   - 因此，文本指出手动签发证书的方式 **"就不可行了"**。
2. **解决方案：Service Account（服务账户）**
   - Kubernetes 设计了 **Service Account** 机制来专门解决 Pod 的身份认证问题。
   - **Service Account** 是一种 Kubernetes 资源对象，它代表了一个在 Pod 中运行的应用程序的身份，而不是用户（User Account）的身份。
   - 每个命名空间（Namespace）都有一个默认的 Service Account（名为 `default`）。你也可以创建自定义的 Service Account。
3. **工作原理：**
   - **自动挂载**：当 Pod 被创建时，如果没有指定特定的 Service Account，它会自动使用所在命名空间的 `default` Service Account。
   - **凭证注入**：Kubernetes 会自动将该 Service Account 所关联的认证令牌（**Token**）**挂载**到 Pod 内部的一个固定路径（`/var/run/secrets/kubernetes.io/serviceaccount`）。
   - **安全访问**：Pod 内的容器进程可以通过读取这个 Token 文件，在访问 API Server 时将其作为 Bearer Token 进行身份认证。
   - **权限控制**：Service Account 可以与 **Role** 和 **RoleBinding**（或 ClusterRole/ClusterRoleBinding）关联，从而实现基于角色的访问控制（RBAC），精确控制该 Pod 有权在集群中执行哪些操作。

##### 组成

###### 1. Kubernetes Secret 资源

- **定义**：Secret 是 Kubernetes 中用于存储和管理**敏感信息**（如密码、OAuth 令牌、ssh 密钥等）的一种资源对象。
- **两种类型**：
  1. **`service-account-token`**：**专用于 Service Account**。由 Kubernetes 自动创建和管理，用于存储 Pod 访问 API Server 所需的凭证。
  2. **`Opaque`**：**通用类型**，用于存储**用户自定义**的保密数据。用户手动创建，可以将数据以 base64 编码的形式存入，并在 Pod 中以环境变量或卷的形式挂载使用。

###### 2. Service Account Token Secret 的三个核心部分

当一个 Service Account 被创建时，Kubernetes 会自动生成一个对应的 `service-account-token` 类型的 Secret。这个 Secret 包含三个关键文件，它们会被自动挂载到使用该 Service Account 的 Pod 中（默认路径：`/var/run/secrets/kubernetes.io/serviceaccount`）。

1. **`token`**:
   - **本质**：这是一个由 **API Server 的私钥**签名的 **JWT (JSON Web Token)**。
   - **作用**：**客户端认证**。当 Pod 中的应用程序访问 API Server 时，会将此 Token 放在 HTTP 请求的 `Authorization: Bearer <token>` 头中。API Server 使用其公钥来验证 JWT 的签名，以此确认 Pod 的身份（即它属于哪个 Service Account）。
2. **`ca.crt`**:
   - **本质**：Kubernetes 集群的**根证书颁发机构（CA）的证书**。
   - **作用**：**服务器端认证**。Pod 中的应用程序（客户端）使用此 CA 证书来验证它所连接的 API Server 提供的 TLS 证书是否合法，确保不会连接到恶意的中间人服务器。
3. **`namespace`**:
   - **本质**：一个文本文件，内容是该 Pod 所在**命名空间的名称**。
   - **作用**：为 Pod 内的应用程序提供上下文信息。应用程序在初始化时可以通过读取这个文件来知道自己运行在哪个命名空间下，从而可以构造正确的 API 请求（例如，只列出当前命名空间的资源）。

> json web token
>
> - **紧凑且安全**：适合在 HTTP 头部中高效传输。
> - **自包含**：Token 本身包含了所有必要的用户身份信息（声明），避免了服务端频繁查询会话数据库。
> - **适用于分布式系统**：这与 Kubernetes 的微服务架构理念完美契合。API Server 作为“服务提供者”，通过验证 JWT 来确认来自 Pod（“身份提供者”认证过的客户端）的请求。

**默认情况下，每个namespace都会有一个 ServiceAccount 如果 pod在创建时没有指定ServiceAccount，那么就会使用pod所属namespace的ServiceAccount。**

**默认挂载在pod容器内的目录下: `/run/secrets/kubernetes.io/serviceaccount/`**

> SA的创建 只需要 kubectl create CA CAName -n namespace 
>
> 背后涉及到 ca token，namespace都是 k8s自动封装创建的



### 鉴权

#### 鉴权类型

API Server 通过 `--authorization-mode` 启动参数来配置授权策略，可以同时启用多个，此时会按顺序检查。

1. **AlwaysDeny**

   - **行为**：**拒绝所有请求**。
   - **用途**：主要用于**测试**。可以快速确保授权环节是否被正确触发，或者临时阻断所有访问。

2. **AlwaysAllow**

   - **行为**：**允许所有请求**。
   - **用途**：适用于**不需要任何授权检查**的环境，例如小型私有测试集群。**在生产环境中使用极其危险**，因为它意味着任何通过认证的实体都拥有集群的最高权限。

3. **ABAC (基于属性的访问控制)**

   - **原理**：管理员需要在一个策略文件中静态地定义一系列规则。每条规则指定了**哪些属性**（用户、组、资源类型、命名空间、API 组等）在**什么条件下**具有**何种权限**。
   - **示例规则**：“用户 `alice` 在命名空间 `project1` 中对 `pods` 拥有 `get`, `watch`, `list` 权限”。
   - **缺点**：策略文件是静态的，每次修改都需要**重启 API Server**，非常不灵活，难以管理大型集群的复杂权限。

4. **Webhook**

   - **原理**：API Server 不会自己做授权决策，而是将授权决定**外包**给一个外部的 RESTful 服务。当需要授权时，API Server 会向预先配置好的 Webhook 服务器发送一个请求，询问“用户 X 能否对资源 Y 执行 Z 操作？”，并根据外部服务的响应（允许或拒绝）来决策。
   - **用途**：用于集成企业已有的**统一授权系统**（如 Open Policy Agent - OPA）。

5. **RBAC (基于角色的访问控制) - 现行默认规则**

   - **原理**：这是 Kubernetes **当前默认且最主流的授权模式**。其权限管理通过一组 Kubernetes **资源对象**来实现：

     - **Role / ClusterRole**：定义**一组规则**，声明了对哪些资源（如 pods、services）可以执行哪些动词（如 get, create, delete）。
       - `Role` 是命名空间范围的。
       - `ClusterRole` 是集群范围的。
     - **RoleBinding / ClusterRoleBinding**：将 **Role/ClusterRole 绑定到特定的主体**（用户、组或 ServiceAccount），从而授予其相应的权限。
       - `RoleBinding` 在特定命名空间内进行绑定。
       - `ClusterRoleBinding` 在整个集群范围内进行绑定。

   - **优点**：

     - **无需重启**：权限配置通过 API 对象完成，可以动态修改，无需重启 API Server。

     - **意图明确**：权限分配清晰可见，易于理解和审计。

     - **灵活性高**：完美适配 Kubernetes 的资源模型和命名空间概念。

     - 整个RPAC 完全由几个API完成，同其他API对象一样，可以用kubectl和API进行操作

     - 可以在运行时进行调整 无需重启api-server

     - RBAC 引用了的4个顶级的资源对象： Role ClusterRole Rolebingding ClusterRoleBinding 4种资源对象均可以通过kubectl与API操作

       > 1. **左侧 (What - 操作与资源)**
       >
       >    - **Resource（资源）**: 这代表 Kubernetes 中的**操作对象**，例如 Pods、Services、Deployments、Secrets 等。
       >    - **create, get, update (动词)**: 这代表了对资源可以执行的**操作**。这些是 Kubernetes API 的 HTTP 动词（POST, GET, PUT 等）的抽象。其他常见动词还包括 `list`, `delete`, `watch`, `patch` 等。
       >
       >    **这一侧共同定义了一组“权限”**，例如：“拥有 `create`、`get`、`update` `Pods` 的权限”。
       >
       > 2. **右侧 (Who - 主体)**
       >
       >    - **User（用户）**: 代表外部的人类用户或通过证书认证的客户端。
       >    - **Group（组）**: 一组用户的集合。给一个组授权，组内的所有用户都会获得相应的权限。
       >    - **ServiceAccount（服务账户）**: 代表在 Pod 中运行的应用的身份。**这是最常被绑定的主体**，用于控制 Pod 能对集群做什么操作。
       >
       >    **这一侧定义了需要被授予权限的“主体”**。
       >
       > 3. **中间 (How - 授权机制)**
       >
       >    - **Role（角色）**: 它是**规则的集合**。一个 Role 精确地定义了“左侧”的内容，即“可以对哪些资源执行哪些操作”。
       >      - **重要提示**: `Role` 是命名空间（Namespace）范围的资源，它只能授予访问同一命名空间内资源的权限。图片中省略了它的集群版本 `ClusterRole`（用于定义集群范围资源的权限，如 Nodes、PersistentVolumes，或跨所有命名空间的权限）。
       >    - **RoleBinding（角色绑定）**: 它充当了**桥梁**的作用。它将一个 `Role`（定义了权限）与一个或多个“主体”（User, Group, ServiceAccount）**绑定**起来。从而将角色中包含的权限授予这些主体。
       >      - **同样提示**: `RoleBinding` 也有其集群版本 `ClusterRoleBinding`，用于将 `ClusterRole` 绑定到主体。
       >
       > ------
       >
       > ### 如何理解它们之间的关系？
       >
       > 图片中的箭头和布局暗示了以下关系：
       >
       > - **Role 是核心**：它封装了权限规则（Resource + Verbs）。
       > - **RoleBinding 是连接器**：它**引用（References）** 一个 Role，并**绑定（Binds）** 一个或多个主体（Subjects）。
       > - **最终效果**：通过创建 Role 和 RoleBinding，你完成了这样一条授权链：
       >   `(User/Group/ServiceAccount)` <-[被 RoleBinding 绑定]- `(Role)` <-[包含权限]-> `(Resource + Verbs)`

       

     ![image-20250824113936069](.\asset\images\image-20250824113936069.png)

> 绑定类型
>
>  Role -----> RoleBinding namespace下
>
> ClusterRole -----> ClusterRoleBinding 集群下
>
> ClusterRole -----> RoleBinding 降维 作用域被限制在namespace下



#### Role

- **Role 的本质**：它是一个**规则（rules）的集合**，这些规则定义了**对哪些资源可以执行哪些操作**。
- **权限累加原则**：RBAC 的权限是**“白名单”机制，只会叠加，不会减少**。你无法通过一个 Role 来收回某个主体已经拥有的权限。如果一个用户被多个 RoleBinding 绑定了多个 Role，那么他的最终权限是这些 Role 所定义权限的并集。
- **作用范围**：Role 是**命名空间（Namespace）级别**的资源。这个例子中的 Role 只会在 `default` 命名空间中生效。要为集群级别的资源（如 Node、PersistentVolume）或跨所有命名空间授权，需要使用 `ClusterRole`。

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1beta1 
metadata:
    namespace: default
    name: pod-reader
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
```



#### ClusterRole

- **ClusterRole 的本质**：它与 Role 一样，是**规则（rules）的集合**。关键区别在于它的**作用范围是集群级别**的，不受命名空间的限制。

- **三大核心用途**：

  1. **集群级别的资源控制**：授权访问**集群范围**的资源，这些资源不属于任何命名空间。例如：
     - `nodes`
     - `persistentvolumes`
     - `namespaces` 本身
     - `clusterroles`, `clusterrolebindings`
  2. **非资源型 endpoints**：授权访问**不属于任何特定 REST 资源**的 API 路径。例如：
     - `/healthz`：用于健康检查
     - `/api`， `/apis`：用于发现 API
     - `/metrics`：用于提供指标数据
  3. **所有命名空间资源控制**：为**跨所有命名空间**的命名空间资源（如 pods、secrets、configmaps）定义权限。如果你希望一个用户或组件能管理所有命名空间中的 Pod，只需为其绑定一个具有相应权限的 ClusterRole，而无需为每个命名空间重复创建 Role。

  ```yaml
  kind: ClusterRole
  apiVersion: rbac.authorization.k8s.io/v1beta1
  metadata:
      # "namespace" omitted since ClusterRoles are not namespaced
      name: secret-reader
  rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "watch", "list"]
  ```
  
  4. 子资源
  
     - **资源与 API URL**：Kubernetes 中的所有操作都通过 RESTful API 进行，每个资源都有其对应的 URL 路径。例如，Pod 的路径是 `/api/v1/namespaces/{namespace}/pods/{name}`。
     - **子资源（Subresource）**：许多资源拥有**子资源**，它们是隶属于父资源下的特定端点，用于执行某些特定操作。最常见的例子就是：
       - **`pods/log`**：用于获取 Pod 的日志（对应 `kubectl logs` 命令）。
       - **`pods/exec`**：用于进入 Pod 执行命令（对应 `kubectl exec` 命令）。
       - **`pods/status`**：用于获取或更新 Pod 的状态。
       - **`services/proxy`**：用于代理到 Service 等。
     - **授权方式**：在 RBAC 规则的 `resources` 字段中，可以通过 **`<资源>/<子资源>`** 的格式来单独为子资源授权。例如，`pods/log` 表示 Pod 的 log 子资源。
  
     ```yaml
     kind: Role
     apiVersion: rbac.authorization.k8s.io/v1beta1 # 修正了 apiversion 和 vlbeta1 的识别错误
     metadata:
       namespace: default
       name: pod-and-pod-logs-reader
     rules:
     - apiGroups: [""]
       resources: ["pods", "pods/log"] # 修正了资源字段的缩进和内容
       verbs: ["get", "list"]
     ```
  
     



#### Binding

##### RoleBinding + Role

- **RoleBinding 的本质**：它是**权限授予的桥梁**。它通过引用一个定义好的角色（Role 或 ClusterRole），将该角色所包含的**权限授予一个或多个主体**。
- **Subjects (主体列表)**：`subjects` 字段指定了权限要授予给谁。它可以包含多种类型的授权对象：
  - **Users**：外部用户。
  - **Groups**：用户组。
  - **ServiceAccounts**：服务账户（**这是最常见的主体**，用于 Pod 身份）。
- **roleRef (角色引用)**：`roleRef` 字段指定了**哪个角色**的权限将被授予。它就像一张指向 Role 或 ClusterRole 的“指针”。
- **作用范围**：RoleBinding 是**命名空间级别**的资源。这个例子中的 RoleBinding 只在 `default` 命名空间中生效。图片末尾提到，如果要进行集群范围的授权，需要使用 `ClusterRoleBinding`。

- RoleBinding + Role

  ```yaml
  kind: RoleBinding
  apiVersion: rbac.authorization.k8s.io/v1beta1
  metadata:
    name: read-pods
    namespace: default
  subjects:
  - kind: User
    name: jane
    apiGroup: rbac.authorization.k8s.io
  roleRef:
    kind: Role
    name: pod-reader
    apiGroup: rbac.authorization.k8s.io
  ```

##### ClusterRole + ClusterRoleBinding 

- **核心机制**：一个**命名空间级别**的 **RoleBinding** 可以引用一个**集群级别**的 **ClusterRole**。

- **最终效果**：通过这种引用，**ClusterRole 所定义的强大权限将被“限制”在 RoleBinding 所在的命名空间内生效**。

- **巨大优势**：这允许集群管理员实现 **“权限模板”的复用**。

  1. **定义通用模板**：管理员可以事先定义一些通用的、具有广泛权限的 ClusterRole（例如：`secret-reader`, `pod-admin`, `namespace-viewer`）。
  2. **在命名空间中授权**：然后，各个命名空间的管理员可以在自己的命名空间内创建 RoleBinding，通过 `roleRef` 指向这些通用的 ClusterRole，来为特定主体授权。
  3. **结果**：**避免了在每个命名空间中重复创建内容和权限完全相同的 Role**，极大地简化了权限管理和维护工作。

- ```yaml
  kind: RoleBinding
  apiVersion: rbac.authorization.k8s.io/v1beta1
  metadata:
    name: read-secrets
    namespace: dev # This only grants permissions within the "development" namespace.
  subjects:
  - kind: User
    name: dave
    apiGroup: rbac.authorization.k8s.io
  roleRef:
    kind: ClusterRole			# rolebinding 到 clusterrole
    name: secret-reader
    apiGroup: rbac.authorization.k8s.io
  ```

  

##### ClusterRole +  ClusterRoleBinding

- **ClusterRoleBinding 的本质**：它是 **ClusterRole 的绑定器**，其最关键的特征是：**它授予的权限是集群级别的，不受任何命名空间的限制**。

- **作用范围**：通过 ClusterRoleBinding 授予的权限**在整个集群中有效**。这意味着被授权的主体可以访问：

  1. 所有命名空间中的资源（如所有命名空间中的 Pods、Secrets）。
  2. 集群级别的资源（如 Nodes、PersistentVolumes）。
  3. 非资源端点（如 `/healthz`）。

- ```yaml
  kind: ClusterRoleBinding
  apiVersion: rbac.authorization.k8s.io/v1beta1 # 修正了 apiVersion 和 vibetal 的识别错误
  metadata:
    name: read-secrets-global
  subjects:
  - kind: Group
    name: manager
    apiGroup: rbac.authorization.k8s.io
  roleRef:
    kind: ClusterRole
    name: secret-reader
    apiGroup: rbac.authorization.k8s.io
  ```

  

> 实验
>
> 创建证书 ---> 转换为 kubeconfig文件  ---> 创建的namspace  ---> 角色绑定



| 访问的资源                                                   | 使用的角色类型  | 使用的绑定类型     |
| :----------------------------------------------------------- | :-------------- | :----------------- |
| 集群级别的资源 (Nodes、PersistentVolumes、……)                | ClusterRole     | ClusterRoleBinding |
| 非资源型 URL (`/api`, `/healthz`, ……)                        | ClusterRole     | ClusterRoleBinding |
| 在任何命名空间中的资源（和跨所有命名空间的资源）             | ClusterRole     | ClusterRoleBinding |
| 在具体命名空间中的资源（在多个命名空间中重用这个相同的 ClusterRole） | **ClusterRole** | **RoleBinding**    |
| 在具体命名空间中的资源 (Role 必须在每个命名空间中定义好）    | **Role**        | **RoleBinding**    |



| 角色名称            | 描述                                                         | 适用场景                                                     |
| :------------------ | :----------------------------------------------------------- | :----------------------------------------------------------- |
| **`view`**          | 允许对某个命名空间内**大多数资源**的**只读**权限。无法查看 Secrets 或修改任何资源。 | 授予开发者或监控工具，使其能够查看应用状态、日志（因为可以读 Pod），但无法进行任何更改。 |
| **`edit`**          | 允许对某个命名空间内**大多数资源**的**读写**权限（包括创建、删除、更新 Pods, Deployments 等）。**通常仍然无法修改资源访问权限本身**（如 Roles/Bindings）。 | 授予开发者，使其能够在特定命名空间内进行日常开发、调试和部署操作。 |
| **`admin`**         | 允许对某个命名空间内**所有资源**的**完全控制**权限，**包括在该命名空间内创建 Role 和 RoleBinding 的权限**。除了 ResourceQuota 和 命名空间资源本身. | 授予命名空间负责人或团队负责人，使其拥有该命名空间的完全自主权。 |
| **`cluster-admin`** | **超级用户权限**。允许对**集群中所有资源**执行任何操作。**如果通过 ClusterRoleBinding 绑定，则授权整个集群；如果通过 RoleBinding 绑定，则授权单个命名空间内的所有资源**。 | **极其谨慎地使用**。通常仅授予极少数集群管理员。             |





### 准入控制

#### 核心概念

- **是什么**：准入控制器是一段代码，它会在请求通过**认证**和**授权**之后、**持久化到数据库 (etcd) 之前**，拦截对 API Server 的请求。
- **做什么**：它可以**验证**请求的合法性，甚至可以**修改**请求对象。
- **类比**：想象一个公司的采购流程：
  - **认证 (Authentication)**：检查你的工牌，确认你是公司员工。
  - **授权 (Authorization)**：检查你的部门权限，确认你有权申请采购。
  - **准入控制 (Admission Control)**：财务部门（准入控制器）审核你的采购申请单，可能会**拒绝**不合理的申请（Validating），也可能会**修改**你的申请，比如将申请的型号更正为公司协议价型号（Mutating），最后才正式批准执行。

#### 2. 准入控制器的类型

- **变更准入控制器 (Mutating Admission Controllers)**：可以先修改发送到 API Server 的对象。多个变更控制器会按顺序调用。
- **验证准入控制器 (Validating Admission Controllers)**：可以验证发送到 API Server 的对象。所有验证控制器都会调用，并且任何控制器拒绝请求，该请求就会失败。

#### 3. 图片中提到的准入控制器解析

1. **`NamespaceLifecycle`**：
   - **功能**：
     - **确保处于终止状态的命名空间无法创建新对象**
     - 防止在不存在的命名空间中创建对象。
     - 防止删除系统预制 namespace
     - 删除namespace时 连带删除他所有的资源对象
   - **为什么重要**：防止在已删除或无效的命名空间中创建资源，维护命名空间生命周期的完整性。
2. **`LimitRanger`**：
   - **功能**：观察传入的请求，确保它不违反命名空间中 `LimitRange` 对象设置的约束（例如，为 Pod 设置默认的内存请求和限制）。
   - **为什么重要**：实施资源配额，防止单个 Pod 或容器消耗过多资源，实现公平调度和成本控制。
3. **`ServiceAccount`**：
   - **功能**：**自动执行**：如果 Pod 没有指定 ServiceAccount，则自动为其添加默认的 ServiceAccount。这是 ServiceAccount 机制得以实现的基础。
   - **为什么重要**：正如我们之前讨论的，它是 Pod 身份认证的基石。
   - 例如之前提到的 每创建一个新的ns 就会在该ns下自动创建的一个新的名为default的ServiceAccount
4. **`DefaultStorageClass`**：
   - **功能**：如果 PersistentVolumeClaim (PVC) 没有指定任何存储类 (StorageClass)，则自动为其添加默认的 StorageClass。
   - **为什么重要**：简化用户操作，用户创建 PVC 时无需关心后台使用哪个存储系统。
5. **`DefaultTolerationSeconds`**：
   - **功能**：为 Pod 设置默认的容忍时间，用于应对节点不可用（如 `not-ready` 和 `unreachable` 等 taint）。
   - **为什么重要**：在节点出现短暂问题时，避免 Pod 被立即驱逐，提高应用韧性。
6. **`MutatingAdmissionWebhook`** 和 **`ValidatingAdmissionWebhook`**：
   - **功能**：这是两个**极其强大**的控制器。它们本身不直接做决策，而是允许你部署一个外部的 HTTP 回调服务（Webhook Server）。API Server 会在准入阶段调用你定义好的 Webhook 服务。
     - `MutatingAdmissionWebhook`：调用外部服务来**修改**对象。
     - `ValidatingAdmissionWebhook`：调用外部服务来**验证**对象。
   - **为什么重要**：它们极大地扩展了 Kubernetes 的准入控制能力，允许用户实现**自定义的、极其复杂的**准入逻辑，是实现高级策略和集群治理的关键（例如：强制给所有 Pod 打上标签、检查镜像是否来自可信仓库、验证资源配置是否合规等）。
7. **`ResourceQuota`**：
   - **功能**：观察传入的请求，确保它不违反命名空间中 `ResourceQuota` 对象设置的总体配额约束（例如，整个命名空间最多只能创建 10 个 Pod，或总共只能请求 16GiB 内存）。
   - **为什么重要**：在多租户集群中至关重要，用于限制单个命名空间或用户所能消耗的资源总量，防止其耗尽整个集群的资源。



## Helm

- **部署流程繁琐**：即使是一个简单的应用，也需要用户手动按顺序创建多个 Kubernetes 资源清单文件（如 `deployment.yaml`, `service.yaml`, `configmap.yaml`, `ingress.yaml` 等），并多次调用 `kubectl apply -f <file>` 命令。这个过程容易出错且效率低下。
- **管理复杂度高**：对于微服务架构的应用，可能包含几十个甚至上百个服务。每个服务都有上述的一套配置文件。要管理这些文件的版本、不同环境的配置（开发、测试、生产）、应用的升级和回滚，变得极其困难，几乎无法有效进行。
- **缺乏封装和版本控制**：传统的 `kubectl` 直接部署方式，很难将一组相关的 K8s 资源作为一个完整的“应用程序”来整体管理，缺乏版本化、一键安装和卸载的能力。

**总结：`kubectl` 更适合管理独立的、静态的资源，而不是复杂的、完整的应用程序。**

### 基本概念

- **本质**：Helm 是一个**模板渲染引擎 + 包管理工具**。
  - **“可配置，能动态生成”**：这是 Helm 最核心的概念。Helm 允许你为 Kubernetes 资源清单文件创建**模板（Template）**。在模板中，你可以使用变量（如 `{{ .Values.image.tag }}`）。
  - **“动态生成 K8s 资源清单文件”**：用户不需要直接编写最终的 `deployment.yaml`，而是提供一个**值文件（values.yaml）** 来定义模板中变量的具体值。Helm 客户端会根据**模板**和**值文件**，在内存中动态渲染并生成最终的 Kubernetes 资源清单文件。
  - **“调用 Kubectl 自动执行”**：Helm 最终会将这些动态生成的 YAML 文件发送给 Kubernetes API Server（内部调用类似于 `kubectl apply` 的逻辑），完成应用的部署。

------

### 核心概念

为了更深入地理解，需要了解 Helm 的三个核心概念：

1. **Chart**：
   - Helm 的**打包格式**，就像一个 Linux 的 `rpm`/`deb` 包或者一个 Homebrew 的 formula。
   - 它就是一个包含了所有必要文件（模板、默认值文件、依赖信息等）的目录集合，用于定义一个完整的应用程序。
   - 这解决了“**打包**”的问题。
2. **Release**：
   - 在 Kubernetes 集群中运行的 **Chart 的一个实例**。
   - 每次安装或升级一个 Chart，都会创建一个新的 **Release**。这意味着你可以将同一个 Chart（应用）**安装多次**到同一个集群，每次都会生成一个独立的、不干扰的应用实例（例如，为不同用户部署同样的博客系统）。
   - 这解决了“**发布和版本控制**”的问题。
3. **Repository**：
   - 用于存放和共享 Charts 的 HTTP 服务器，类似于 Docker Registry 或 Maven Central。
   - 这解决了“**共享和分发**”的问题。

### 工作流程举例

假设有一个名为 `nginx` 的 Chart，其模板中定义了镜像标签为 `{{ .Values.imageTag }}`。

1. 用户准备一个 `values-prod.yaml` 文件：`imageTag: "1.21.6"`
2. 用户执行命令：`helm install my-nginx-release ./nginx -f values-prod.yaml`
3. Helm 读取 `nginx` Chart 的模板。
4. Helm 将模板中的 `{{ .Values.imageTag }}` 替换为 `1.21.6`，在内存中生成最终的 Deployment YAML（其中镜像是 `nginx:1.21.6`）。
5. Helm 将生成的所有 YAML 资源提交给 Kubernetes API Server。
6. 一个名为 `my-nginx-release` 的应用实例就在集群中运行起来了。



> ![image-20250824134429400](.\asset\images\image-20250824134429400.png)
>
> #### 1. Helm v2 的架构与问题 (C/S 架构)
>
> - **客户端 (Client)**：本地安装的 `helm` 命令行工具。
> - **服务端 (Server)**：部署在 Kubernetes 集群内部的 `tiller` Pod。
> - **工作流程**：
>   1. 用户执行 `helm install ...` 命令。
>   2. 本地的 `helm` 客户端将 Chart 打包发送给集群内的 `tiller`。
>   3. `tiller` 接收到 Chart 后，在集群内部渲染模板，生成 Kubernetes 资源清单文件。
>   4. `tiller` 调用 Kubernetes API，创建这些资源。
> - **核心问题**：
>   - **巨大的安全风险**：`tiller` 需要一个权限非常大的 ServiceAccount（通常是 `cluster-admin`）才能在集群内创建资源。这意味着任何能访问 `tiller` 服务的人（例如，如果其服务端口被意外暴露），就间接获得了整个集群的最高权限。
>   - **复杂的权限管理**：与 Kubernetes 的 **RBAC** 精细权限模型背道而驰。很难实现“某个用户只能通过 Helm 在特定命名空间部署应用”这种需求，因为所有操作都是由同一个高权限的 `tiller` 完成的。
>   - **多租户支持困难**：无法区分到底是哪个用户发起的 `helm` 命令，所有操作都记录为 `tiller` 这个身份。
>
> #### 2. Helm v3 的架构革新 (移除 Tiller)
>
> Helm v3 做出了一个大胆而正确的决定：**彻底移除 `tiller`**。
>
> - **新架构**：从 C/S 架构变为**纯客户端架构**。
> - **新工作流程**：
>   1. 用户执行 `helm install ...` 命令。
>   2. 本地的 `helm` 客户端利用本地的模板引擎渲染 Chart。
>   3. `helm` 客户端**直接使用本地的 `kubeconfig` 文件**（和 `kubectl` 一样）与 Kubernetes API Server 通信，创建渲染后的资源。
>
> #### 3. 移除 Tiller 带来的四大好处
>
> 图片中详细列举了移除 Tiller 的直接好处：
>
> 1. **架构更为简单和灵活**：
>    - **简单**：无需在集群中维护 `tiller` 组件的运行和健康状态。
>    - **灵活**：`helm` 客户端的行为不再受限于集群内 `tiller` 的版本，版本兼容性管理变得更简单。
> 2. **不再需要创建 ServiceAccount**：
>    - 在 Helm v2 中，你必须事先为 `tiller` 创建一个拥有足够权限的 ServiceAccount 并进行绑定，部署流程繁琐。
>    - 在 Helm v3 中，**完全省去了这一步**。Helm 直接继承执行 `helm` 命令的用户的权限（由 `kubeconfig` 文件定义）。这意味着你可以直接使用 Kubernetes 的 **RBAC 系统**来精细控制用户通过 Helm 能做什么、不能做什么，实现了完美的权限集成。
>    - 即使用了当前 .kube/config下的权限
> 3. **更为安全**：
>    - 这是最核心的好处。安全风险从“保护一个集群内的高权限 `tiller` 服务”转变为“保护用户的 `kubeconfig` 文件”，后者是 Kubernetes 的标准安全实践，更易于管理和审计。
>    - 操作者的身份得以明确，方便审计追踪。`helm` 执行的操作就是当前用户的操作，遵守了最小权限原则。
> 4. **不再需要使用 `helm init` 来进行初始化**：
>    - 在 Helm v2 中，使用前必须在集群中执行 `helm init` 来部署 `tiller`。
>    - 在 Helm v3 中，**完全没有初始化步骤**。只要你有一个配置好的 `kubeconfig` 文件，可以立即开始使用 `helm`，就像使用 `kubectl` 一样简单。



> 1. **安装 (Helm v3)**
>
>    - **命令**：
>
>      ```bash
>      curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
>      chmod 700 get_helm.sh
>      ./get_helm.sh
>      ```
>
> 2. **初始化 (添加Chart仓库)**
>
>    - **说明**：安装后需添加Chart仓库来获取可安装的应用程序包。
>
>    - **命令**：
>
>      ```bash
>      helm repo add bitnami https://charts.bitnami.com/bitnami  # 添加Bitnami仓库
>      helm search repo bitnami  # 搜索该仓库中的Charts
>      ```
>
> 3. **安装Chart示例**
>
>    - **命令**：
>
>      ```bash
>      helm repo update                 # 更新仓库信息，获取最新Charts列表
>      helm show values bitnami/apache  # 查看某个Chart（如Apache）的可配置参数
>      helm install bitnami/apache --generate-name  # 安装该Chart并让Helm自动生成发布名称
>      helm show chart bitnami/apache   # 显示Chart的基本信息
>      helm show all bitnami/apache     # 显示Chart的详细信息
>      ```
>
> 4. **卸载一个Release**
>
>    - **命令**：
>
>      ```bash
>      helm uninstall apache-1612624192    # 彻底卸载该发布，删除相关所有Kubernetes资源
>      helm uninstall --keep-history       # 卸载但保留该发布的版本历史记录
>      helm status apache-1612624192       # 查看该发布的状态信息
>      ```

> #### **三大概念 (Chart, Release, Repository)**
>
> 1. **Chart**
>
>    - **定义**：Helm软件包，包含在Kubernetes集群中运行应用所需的所有资源定义模板。类似于Linux系统中的软件包格式（如RPM, DEB）。
>
> 2. **`helm search`: 查找Charts**
>
>    - **命令**：
>
>      ```bash
>      # 用于在 Helm Hub（https://hub.helm.sh）上搜索 Helm charts
>      helm search hub wordpress   # 在Artifact Hub上搜索Charts（需网络）
>      helm search repo wordpress  # 在本地已配置的仓库中搜索Charts
>      ```
>
> 3. **`helm install`: 安装一个Helm包**
>
>    - **命令**：`helm install apache-1612624192 bitnami/apache`
>
>    - **安装顺序**：文档列出了Kubernetes资源的创建顺序（如Namespace -> NetworkPolicy -> ... -> Ingress）。
>
>      > ```text
>      > Namespace > NetworkPolicy > ResourceQuota > LimitRange > 
>      > PodSecurityPolicy > PodDisruptionBudget > ServiceAccount > Secret > 
>      > SecretList > ConfigMap > StorageClass
>      > PersistentVolume > PersistentVolumeClaim > CustomResourceDefinition 
>      > > ClusterRole
>      > ClusterRoleList > ClusterRoleBinding > ClusterRoleBindingList > Role 
>      > > RoleList
>      > RoleBinding > RoleBindingList > Service > DaemonSet > Pod > 
>      > ReplicationController > ReplicaSet > Deployment > 
>      > HorizontalPodAutoscaler > StatefulSet > Job > CronJob > Ingress > 
>      > APIService
>      > ```
>      >
>      > 
>
> 4. **安装前自定义Chart**
>
>    - **步骤**：
>      1. `helm show values bitnami/apache`：查看Chart的可配置选项。
>      2. 创建一个YAML文件（如`values.yaml`）来覆盖默认配置。
>      3. `helm install -f values.yaml bitnami/apache --generate-name`：使用自定义配置安装。
>
> 5. **传递配置数据的两种方式**
>
>    - `--values` / `-f`：使用YAML文件覆盖配置。
>
>    - `--set`：通过命令行直接设置值。
>
>    - 如果同时使用两种方式，则 --set 中的值会被合并到 --values 中，但是 --set 中的值优先级更高。
>
>      在 --set 中覆盖的内容会被被保存在 ConfigMap 中。可以通过 helm get values <release-name> 来
>
>      查看指定 release 中 --set 设置的值。也可以通过运行 helm upgrade 并指定 --reset-values 字段来
>
>      清除 --set 中设置的值
>
>    - **优先级**：`--set` > `-f` > Chart默认的`values.yaml`。
>
>    - **`--set`的格式**：支持简单的键值对、嵌套结构、列表、空值、访问数组元素和转义特殊字符。
>
>      - 选项使用0或多个 name/value 对。最简单的用法类似于： --set name=valu
>
>      - 多个值使用逗号分割，因此 --set a=b,c=d 的 YAML 表示是：
>
>        ```yaml
>        a: b
>        c: d
>        ```
>
>      - 支持更复杂的表达式。例如， --set outer.inner=value 被转换成了：
>
>        ```yaml
>        outer:
>         inner: value
>        ```
>
>      - 列表使用花括号（ {} ）来表示。例如， --set name={a, b, c} 被转换成了：
>
>        ```yaml
>        name:
>         - a
>         - b
>         - c
>        ```
>
>      - 某些 name/key 可以设置为 null 或者空数组，例如 --set name=[],a=null
>
>      - 如果需要在 --set 中使用特殊字符，你可以使用反斜线来进行转义； --set name=value1\,value2 就变成了：
>
>        ```yaml
>        name: "value1,value2"
>        ```
>
>        --set nodeSelector."kubernetes.io/role"=master
>
>        ```yaml
>        nodeSelector:
>         kubernetes.io/role: master
>        ```
>
>      - 示例：`--set name=value`, `--set a=b,c=d`, `--set outer.inner=value`, `--set name={a,b,c}`, `--set servers[0].port=80`, `--set name=value1\,value2`
>
> 6. **更多安装方法**
>
>    Helm可以从多种来源安装：仓库、本地压缩包（`.tgz`）、解压后的目录、完整的URL。
>
>    - chart 的仓库（如上所述）
>    - 本地 chart 压缩包（ helm install foo foo-0.1.1.tgz ）
>    - 解压后的 chart 目录（ helm install foo path/to/foo ）
>    - 完整的 URL（ helm install foo https://example.com/charts/foo-1.2.3.tgz ）
>
> 7. **`helm upgrade` 和 `helm rollback`：升级和回滚**
>
>    - **升级**：`helm upgrade -f new-values.yaml apache-23213213 bitnami/apache`
>
>      **Helm 会尝试执行最小侵入式升级。即它只会更新自上次发布以来发生了更改的内容**
>
>    - **回滚**：`helm rollback apache-23213213 1` （回滚到修订版本1）
>
>      上面这条命令将我们的 apache-23213213 回滚到了它最初的版本。release 版本其实是一个增量修订（revision）。 每当发生了一次安装、升级或回滚操作，revision 的值就会加1。第一次 revision 的值永远是1。
>
>      我们可以使用 helm history [RELEASE] 命令来查看一个特定 release 的修订版本号
>
>    - **查看历史**：`helm history [RELEASE].`
>
> 8. **安装、升级、回滚时的有用选项**
>
>    - `--timeout`：设置等待Kubernetes命令完成的超时时间。
>
>    - `--wait`：等待所有Pod就绪、PVC绑定、Service就绪后才标记发布成功。
>
>      表示必须要等到所有的 Pods 都处于 ready 状态，PVC 都被绑定，Deployments 都至少拥有最小 ready 状态 Pods 个数（Desired减去 maxUnavailable），并且 Services 都具有 IP 地址（如果是LoadBalancer， 则为 Ingress），才会标记该 release 为成功。最长等待时间由 --timeout 值指定。如果达到超时时间，release 将被标记为 FAILED。注意：当 Deployment 的replicas 被设置为1，但其滚动升级策略中的 maxUnavailable 没有被设置为0时，--wait 将返回就绪，因为已经满足了最小 ready Pod 数
>
>    - `--no-hooks`：不运行生命周期钩子。
>
>    - `--recreate-pods`：（已废弃）重建所有Pod。
>
> 9. **`helm uninstall`: 卸载release**
>
>    helm v2 版本中，当一个 release 被删除，会保留一条删除记录。
>
>    而在 Helm 3 中，删除也会移除release 的记录。 如果你想保留删除记录，使用 helm uninstall --keep-history 。
>
>    使用 helm list --uninstalled 只会展示使用了 --keep-history 删除的 release
>
>    - **命令**：`helm uninstall apache-23213213`
>    - **保留历史**：使用`helm uninstall --keep-history`。可使用`helm list --uninstalled`或`helm list --all`查看。
>
> 10. **`helm repo`: 使用仓库**
>
>     - **命令**：
>
>       ```bash
>       helm repo list          # 列出已配置的仓库
>       # helm repo add reponame url
>       helm repo add dev https://example.com/dev-charts  # 添加新仓库
>       helm repo update        # 更新本地仓库缓存
>       helm repo remove reponame       # 移除仓库
>       ```



> #### **创建一个自己的Chart**
>
> 1. **基本模式**
>
>    - **创建模板**：`helm create myapp` （生成标准目录结构）。
>
>      > ```bash
>      > myapp/
>      > ├── charts/		   # 此目录用于存放此 Chart 所依赖的【子Charts】	
>      > ├── Chart.yaml     # 这是 Chart 的元数据文件，类似于一个软件的 package.json 或 setup.py。
>      > ├── templates/     # 这是放置 Kubernetes 资源清单模板的地方。Helm 的核心模板引擎会处理这个目录中的所有文件（除了 _helpers.tpl），将它们与 values.yaml 和用户提供的值合并，渲染成最终的 Kubernetes YAML 清单，然后提交给 API Server。
>      > │   ├── deployment.yaml
>      > │   ├── _helpers.tpl
>      > │   ├── hpa.yaml
>      > │   ├── ingress.yaml
>      > │   ├── service.yaml
>      > │   ├── serviceaccount.yaml
>      > │   └── tests/
>      > │       └── test-connection.yaml
>      > └── values.yaml   # 这是 Chart 的元数据文件，类似于一个软件的 package.json 或 setup.py。
>      > ```
>
>    - **编写模板**：在`templates/`目录下编写Kubernetes资源清单文件（如`service.yaml`, `deployment.yaml`），并使用Go Template语法使其动态化。
>
>    - **发布部署**：`helm install myapp myapp/` （安装自定义的Chart）。
>
> 2. **注入Helm灵魂 (模板化)**
>
>    - 使用`{{ .Values.<path> }}`在模板中引用`values.yaml`中定义的变量。
>
>    - **示例**：
>
>      - `image: {{ .Values.image.repository }}:{{ .Values.image.tag }}`
>      - `replicas: {{ .Values.replicaCount }}`
>      - `type: {{ .Values.service.type | quote }}` （`quote`函数添加引号）
>
>    - **使用条件判断**：
>
>      yaml
>
>      ```
>      {{- if eq .Values.service.type "NodePort" }}
>      nodePort: {{ .Values.service.nodeport }}
>      {{- end }}
>      ```
>
>    - **使用函数**：`{{ now | date "20060102030405" }}` （生成时间戳）。
>
>    - **定义Values**：在`values.yaml`中为模板中的变量设置默认值。

## 镜像仓库

Harbor 的目标是帮助用户迅速搭建一个企业级的 Docker registry 服务。它以 Docker 公司开源的 registry 为基础，额外提供了如下功能：

- 基于角色的访问控制（Role Based Access Control）
- 镜像的删除和空间清理（Image deletion & garbage collection）
- 基于策略的镜像复制（Policy based image replication）
- 镜像的漏洞扫描（Vulnerability Scanning）
- 友好的管理UI（Graphical user portal）
- AD/LDAP集成（LDAP/AD support）
- 审计日志（Audit logging）
- RESTful API
- 部署简单（Easy deployment）

> ![image-20250831113026894](.\asset\images\image-20250831113026894.png)







安装

```yaml
apiVersion: v1
kind: Namespace
metadata:
 annotations:
   scheduler.alpha.kubernetes.io/node-selector: harbor=env
 name: harbor
```





## 日志方案

loki  -- 云原生的日志服务器 可能是目前最具性价比的方案



![image-20250831130946996](.\asset\images\image-20250831130946996.png)



![image-20250831131145781](.\asset\images\image-20250831131145781.png)



> ![image-20250831131333355](.\asset\images\image-20250831131333355.png)
>
> ![image-20250831132405216](.\asset\images\image-20250831132405216.png)
>
> 例如这种 CDN(内容分发) 在每个机房部署一个用于收集存储日志的 L 也好 K 也好 将内容推送至 es，这样虽然l或者k的资源占用比较高 但是也是可以接受的
>
> 但是 对于k8s的架构，需要收集每个node上的pod的xinxi 这样的每个node上都需要部署 一个l或者k，这样的话资源占用就太大了，因此 这里推荐使用云原生 ---- loki



### Loki 是什么？

Loki 是一个**水平可扩展、高可用、多租户的日志聚合系统**，由 Grafana Labs 团队设计并开源。它的设计灵感源自 Prometheus，核心思想是：**只索引日志的元数据（标签），而不索引日志内容本身**，通过这种独特的方式极大地降低了存储和运算开销，使运维更加高效和经济。

您可以将其理解为 Kubernetes 和云原生环境下的 “轻量级 ELK/EFK”（Elasticsearch, Logstash/Kibana, Fluentd）替代方案。

------

### Loki 的核心特点与设计哲学

#### 1. 基于标签的索引（核心创新）

这是 Loki 与 ELK 等传统日志系统最根本的区别。

- **ELK/EFK**：会为日志的**全文**建立索引，这虽然查询功能强大灵活，但需要巨大的存储和计算资源，成本很高。
- **Loki**：**只对日志流的元数据（即标签，Labels）建立索引**，日志内容本身以压缩形式存储。查询时，先通过标签快速缩小范围，再使用类似 `grep` 的方式对压缩的日志内容进行解压和搜索。

**举个例子：**
一条日志来自 `cluster=prod`, `namespace=monitoring`, `pod=prometheus-xyz123`。

- Loki 只会为 `cluster`, `namespace`, `pod` 这些标签建立索引。
- 日志内容 `level=error msg="request failed"` 不会被索引，而是被压缩存储。
- 当您查询 `{cluster="prod", namespace="monitoring"} |= "request failed"` 时，Loki 会先用索引找到所有 `prod` 集群 `monitoring` 命名空间的日志流，然后再在这些流中“grep”搜索内容包含 “request failed” 的日志行。

#### 2. 与 Prometheus 和 Grafana 无缝集成

- **相同的标签体系**：Loki 使用和 Prometheus **完全相同**的标签（Labels）。这意味着您可以通过在 Prometheus 中发现的 Pod 标签，直接去 Loki 里查询它的日志，无需任何转换，体验非常一致。
- **完美的 Grafana 原生支持**：在 Grafana 面板中，您可以同时展示来自 Prometheus 的指标曲线和来自 Loki 的详细日志，实现真正的“可观测性”（Observability），快速从指标异常定位到具体的错误日志。

#### 3. 极高的成本效益

- **存储开销低**：由于不为日志内容建立索引，且日志内容以压缩块形式存储（通常存储在对象存储中，如 AWS S3, Google Cloud Storage, MinIO 等），Loki 的存储成本通常比 ELK 低一个数量级。
- **运维简单**：架构相对简单，组件更少，与云原生生态结合紧密，部署和运维成本也更低。

#### 4. 云原生（Cloud-Native）

Loki 从一开始就是为动态的云原生环境（尤其是 Kubernetes）而设计的。它能够很好地处理容器和微服务带来的海量、短生命周期的日志源。

------

### Loki 的核心架构组件

一个典型的 Loki 集群包含以下几个主要组件：

![image-20250831140535821](.\asset\images\image-20250831140535821.png)

1. **Distributor**：**分发器**，负责接收日志数据（例如从 Fluent Bit, Promtail 等而来），并进行校验和转发。
2. **Ingester**：**摄取器**，接收来自 Distributor 的日志流，在内存中构建数据块并压缩，然后定期刷写到后端存储中。它是内存密集型组件，负责写操作。
3. **Querier**：**查询器**，负责执行日志查询。它从 Ingester（查询最新数据）和后端存储（查询历史数据）中获取数据，进行聚合后返回给用户。
4. **Query Frontend**：**查询前端**（可选，但推荐），位于 Querier 之前，提供查询调度以提高性能，支持并行查询和缓存，以加速重复查询。
5. **Ruler**：**规则器**（可选），负责持续查询并计算生成警报和记录规则（类似于 Prometheus 的警报），结果可以发送给 Alertmanager。
6. **后端存储**：需要配置两种存储。
   - **索引存储**：用于存储标签索引。支持 DynamoDB, Bigtable, Cassandra，或者更简单的 **BoltDB（单机模式）**。
   - **日志存储**：用于存储压缩的日志内容块。通常支持 **Amazon S3, Google Cloud Storage, Azure Blob Storage, MinIO** 等对象存储。

对于测试或轻量级使用，也有 **all-in-one** 的单进程模式，所有组件都运行在一起，非常简单。

------

### 如何将日志送入 Loki？

最常用的日志收集客户端是 **Promtail** 或 **Fluent Bit**。

- **Promtail**：Grafana Labs 为 Loki 开发的专用客户端，是 Kubernetes 环境下的首选。**它能自动发现 Pod，并从 Pod 中提取日志，并自动附加如 `pod_name`, `namespace` 等标签。**
- **Fluent Bit**：一个超轻量级的、高性能的日志收集处理器，社区生态丰富，也官方支持输出到 Loki。

------

### 适用场景

- **Kubernetes 集群日志管理**：这是 Loki 的“主战场”，体验最佳。
- **成本敏感的日志集中管理**：当您觉得 ELK 栈成本过高或运维复杂时。
- **已经使用 Prometheus + Grafana 技术栈的团队**：可以无缝整合，提升可观测性能力。
- **不需要对日志内容进行复杂关联分析和全文搜索的场景**：Loki 更擅长基于标签的快速过滤和简单的内容搜索。

> ![image-20250831133008454](.\asset\images\image-20250831133008454.png)
>
> | 特性               | Loki                                                 | ELK / EFK                                                    |
> | :----------------- | :--------------------------------------------------- | :----------------------------------------------------------- |
> | **存储方式**       | 使用标签索引和压缩日志块，查询效率高，存储成本低     | 依赖 Elasticsearch 建立全文索引，存储和查询开销大            |
> | **数据收集和处理** | 使用 **Promtail**，轻量、高性能，云原生友好          | ELK 使用 Logstash，EFK 使用 Fluentd；Fluentd 更轻量，Logstash 资源消耗大 |
> | **扩展性**         | 水平扩展能力强，尤其适用于云原生和大规模日志场景     | 扩展性强，支持多种数据源，但架构较重                         |
> | **学习曲线**       | 较简单，尤其适合已有 Prometheus/Grafana 技术栈的用户 | 较复杂，组件多，功能丰富，配置和管理难度较大                 |
> | **可视化工具**     | 使用 **Grafana**，与 Prometheus 等可观测数据无缝集成 | 使用 **Kibana**，提供强大的日志搜索和可视化仪表盘            |
>
> ![image-20250831140416027](.\asset\images\image-20250831140416027.png)

## 资源监控方案

prometheus

> **prometheus**
>
> Prometheus 是一个**开源的系统监控和警报工具**，最初是在 SoundCloud 上构建的。Prometheus 于2016年加入 CNCF（云原生计算基金会），成为继 Kubernetes 之后的第二个托管项目。2018 年 8 月 9 日，CNCF 宣布开放源代码监控工具 Prometheus 已从孵化状态进入毕业状态。
>
> **CAdvisor**
>
> Cadvisor 是 Google 开源的一款**用来检测、分析、展示单节点的一个容器性能指标和资源监控的可视化工具**（也可以监控本机，针对单台物理机），监控包括容器的内存使用率、CPU 使用率、网络 IO、磁盘IO 及文件系统使用情况，利用 Linux 的 Cgroup 获取容器及本机的资源使用信息。同时提供了一个 Web 界面用于查看容器的实时运行状态。
>
>  **kube-state-metrics**
>
> - **项目归属：** 属于 `kubernetes` 组织下的一个项目。
> - **核心功能：** 监听 Kubernetes API Server，并生成关于集群中各种**资源对象状态**的指标。
> - **监控内容举例：** Deployment、DaemonSet、Pod、Node 等资源的状态。例如：
>   - Deployment 的期望副本数 (`replicas`) 和可用副本数 (`available_replicas`)。
>   - Pod 的当前状态（是 Running、Pending 还是 Failed）。
>   - Node 是否处于 Ready 状态。
> - **特点：** 它关注的是**元数据和工作状态**，而不是机器或容器的实时性能数据。
>
> **Metrics Server**
>
> - **核心功能：** 从 Kubernetes API Server 收集集群范围的**资源使用率**指标。
> - **监控内容：** 主要是 **CPU** 和**内存**的实际使用量。
> - **核心作用：** 它是 Kubernetes **自动扩缩容 (HPA)** 和 **垂直扩缩容 (VPA)** 的核心组件。`kubectl top` 命令的数据也来源于它。
> - **特点：** 它关注的是**实时的资源消耗**，为集群的自动化决策提供数据支持。
>
> **Grafana**
>
> - **核心定位：** 一个功能强大的**数据可视化与监控仪表板**平台。
> - **主要功能：**
>   1. **多数据源支持：** 其最关键的特性之一。它可以连接并统一展示来自众多数据源（如 Prometheus、MySQL、Elasticsearch、Loki、Graphite 等）的指标数据。
>   2. **可视化展示：** 提供丰富的图表类型（如折线图、仪表盘、热图、地理地图等），让用户能够通过自定义的仪表板（Dashboard）直观地查看和分析监控数据。
>   3. **告警功能：** 内置的告警规则引擎允许用户直接在仪表板上设置告警规则，并在条件触发时通过多种渠道（如 Slack、Email、Webhook 等）发送通知。
> - **特点：** 专注于 **“数据的呈现与洞察”**，界面美观、灵活性强，是监控数据的“前端脸面”。
>
> **Alertmanager**
>
> - **核心定位：** 专用于处理告警的**告警管理中心**（通常与 Prometheus 配套使用）。
> - **主要功能：**
>   1. **接收告警：** 接收来自 Prometheus Server 等客户端发来的告警信息。
>   2. **告警处理：** 对告警进行一系列高级处理，这是它的核心价值：
>      - **分组 (Grouping)：** 将同一类型的多个告警合并为一个通知，避免告警风暴（例如，同一个集群节点下的几十个Pod实例同时宕机，只会收到一条“XX节点Pod实例异常”的通知）。
>      - **抑制 (Inhibition)：** 抑制某些不必要的告警（例如，整个机房断电，那么所有“服务器宕机”的告警应该抑制其他所有关于该服务器上服务的告警）。
>      - **静默 (Silencing)：** 允许在一段时间内静默特定的告警，适用于计划内的维护。
>      - **路由 (Routing) 与分发：** 根据告警的标签（label）将不同类型的告警路由到不同的接收方（如数据库告警发给DBA团队，网络告警发给网络团队）。
>   3. **通知发送：** 支持通过丰富的渠道（如 Email、Slack、PagerDuty、Webhook 等）将处理后的告警通知发送给相关人员。
> - **特点：** 专注于 **“告警的处理与分发”**，是确保告警信息准确、有效、不冗余的“大脑”。
>
> > 一个例子
> >
> > ![image-20250831181921421](.\asset\images\image-20250831181921421.png)



> **目录**
> 一、部署 Prometheus
> 二、监控 ingress-nginx
> 三、创建 node-exporter，监控节点资源
> 四、监控容器资源指标
> 五、监控 ApiServer 指标
> 六、通过 Service 监控服务
> 七、部署 kube-state-metrics，并监控
> 八、部署 grafana 服务
> 九、监控 metrics.server
> 十、alertmanager 部署
>
> ------
>
> ### 一、部署 Prometheus
>
> bash
>
> ```yaml
> $ kubectl create ns kube-ops
> ```
>
> **创建 prometheus-cm.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: ConfigMap
> metadata:
>   name: prometheus-config
>   namespace: kube-ops
> data:
>   prometheus.yml: |
>     global:
>       scrape_interval: 15s # 表示 prometheus 抓取指标数据的频率，默认是 15s
>       scrape_timeout: 15s  # 表示 prometheus 抓取指标数据的超时时间，默认是 15s
>     scrape_configs:
>     - job_name: 'prometheus'
>       static_configs:
>       - targets: ['localhost:9090']
> ```
>
> **创建 prometheus-pvc.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: PersistentVolumeClaim
> metadata:
>   name: prometheus
>   namespace: kube-ops
> spec:
>   storageClassName: nfs-storage
>   accessModes:
>   - ReadWriteMany
>   resources:
>     requests:
>       storage: 10Gi
> ```
>
> **创建 prometheus-rbac.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: ServiceAccount
> metadata:
>   name: prometheus
>   namespace: kube-ops
> ---
> apiVersion: rbac.authorization.k8s.io/v1
> kind: ClusterRole
> metadata:
>   name: prometheus
> rules:
> - apiGroups:
>   - ""
>   resources:
>   - nodes
>   - services
>   - endpoints
>   - pods
>   - nodes/proxy
>   verbs:
>   - get
>   - list
>   - watch
> - apiGroups:
>   - ""
>   resources:
>   - configmaps
>   - nodes/metrics
>   verbs:
>   - get
> - nonResourceURLs:
>   - /metrics
>   verbs:
>   - get
> ---
> apiVersion: rbac.authorization.k8s.io/v1
> kind: ClusterRoleBinding
> metadata:
>   name: prometheus
> roleRef:
>   apiGroup: rbac.authorization.k8s.io
>   kind: ClusterRole
>   name: prometheus
> subjects:
> - kind: ServiceAccount
>   name: prometheus
>   namespace: kube-ops
> ```
>
> **创建 prometheus-deploy.yaml**
>
> yaml
>
> ```yaml
> apiVersion: apps/v1
> kind: Deployment
> metadata:
>   name: prometheus
>   namespace: kube-ops
>   labels:
>     app: prometheus
> spec:
>   selector:
>     matchLabels:
>       app: prometheus
>   template:
>     metadata:
>       labels:
>         app: prometheus
>     spec:
>       serviceAccountName: prometheus
>       containers:
>       - image: prom/prometheus:v2.4.3
>         name: prometheus
>         command:
>         - "/bin/prometheus"
>         args:
>         - "--config.file=/etc/prometheus/prometheus.yml"
>         - "--storage.tsdb.path=/prometheus"
>         - "--storage.tsdb.retention=24h"
>         - "--web.enable-admin-api" # 控制对admin HTTP API的访问，其中包括删除时间序列
>         - "--web.enable-lifecycle" # 支持热更新，直接执行localhost:9090/-/reload立即生效 (注：此参数在后续页面提到，添加于此)
>         ports:
>         - containerPort: 9090
>           protocol: TCP
>           name: http
>         volumeMounts:
>         - mountPath: "/prometheus"
>           subPath: prometheus
>           name: data
>         - mountPath: "/etc/prometheus"
>           name: config-volume
>         resources:
>           requests:
>             cpu: 100m
>             memory: 512Mi
>           limits:
>             cpu: 100m
>             memory: 512Mi
>         securityContext:
>           runAsUser: 0
>       volumes:
>       - name: data
>         persistentVolumeClaim:
>           claimName: prometheus
>       - name: config-volume
>         configMap:
>           name: prometheus-config
> ```
>
> **创建 prometheus-svc.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: Service
> metadata:
>   name: prometheus
>   namespace: kube-ops
>   labels:
>     app: prometheus
> spec:
>   selector:
>     app: prometheus
>   type: NodePort
>   ports:
>   - name: web
>     port: 9090
>     targetPort: http
> ```
>
> ------
>
> ### 二、监控 ingress-nginx
>
> Prometheus 的数据指标是通过一个公开的 HTTP(S) 数据接口获取到的，我们不需要单独安装监控的 agent，只需要暴露一个 metrics 接口，Prometheus 就会定期去拉取数据；对于一些普通的 HTTP 服务，我们完全可以直接重用这个服务，添加一个 /metrics 接口暴露给 Prometheus。
>
> **创建新的配置文件 prome-cm.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: ConfigMap
> metadata:
>   name: prometheus-config
>   namespace: kube-ops
> data:
>   prometheus.yml: |
>     global:
>       scrape_interval: 30s
>       scrape_timeout: 30s
>     scrape_configs:
>     - job_name: 'prometheus'
>       static_configs:
>       - targets: ['localhost:9090']
>     - job_name: 'ingressnginx12'
>       static_configs:
>       - targets: ['192.168.66.12:10254']
>     - job_name: 'ingressnginx13'
>       static_configs:
>       - targets: ['192.168.66.13:10254']
> ```
>
> **重载 prometheus**
>
> bash
>
> ```yaml
> $ kubectl get svc -n kube-ops
> NAME         TYPE       CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
> prometheus   NodePort   10.11.6.226   <none>        9090:30604/TCP   6m19s
> 
> $ curl -X POST "http://10.11.6.226:9090/-/reload"
> ```
>
> **确认是否捕获数据 PQL:**
>
> promql
>
> ```yaml
> sum(nginx_ingress_controller_nginx_process_requests_total{controller_class="k8s.io/ingress-nginx",controller_namespace="ingress-nginx"})
> ```
>
> ------
>
> ### 三、创建 node-exporter，监控节点资源
>
> node_exporter 用于采集服务器节点的各种运行指标，目前 node_exporter 支持几乎所有常见的监控点，比如 conntrack, cpu, diskstats, filesystem, loadavg, meminfo, netstat 等。
>
> **创建 prome-node-exporter.yaml**
>
> yaml
>
> ```yaml
> apiVersion: apps/v1
> kind: DaemonSet
> metadata:
>   name: node-exporter
>   namespace: kube-ops
>   labels:
>     name: node-exporter
> spec:
>   selector:
>     matchLabels:
>       name: node-exporter
>   template:
>     metadata:
>       labels:
>         name: node-exporter
>     spec:
>       hostPID: true
>       hostIPC: true
>       hostNetwork: true
>       containers:
>       - name: node-exporter
>         image: prom/node-exporter:v0.16.0
>         ports:
>         - containerPort: 9100
>         resources:
>           requests:
>             cpu: 0.15
>         securityContext:
>           privileged: true
>         args:
>         - --path.procfs
>         - /host/proc
>         - --path.sysfs
>         - /host/sys
>         - --collector.filesystem.ignored-mount-points
>         - '"^/(sys|proc|dev|host|etc)($$|/)"'
>         volumeMounts:
>         - name: dev
>           mountPath: /host/dev
>         - name: proc
>           mountPath: /host/proc
>         - name: sys
>           mountPath: /host/sys
>         - name: rootfs
>           mountPath: /rootfs
>       tolerations:
>       - key: "node-role.kubernetes.io/control-plane"
>         operator: "Exists"
>         effect: "NoSchedule"
>       volumes:
>       - name: proc
>         hostPath:
>           path: /proc
>       - name: dev
>         hostPath:
>           path: /dev
>       - name: sys
>         hostPath:
>           path: /sys
>       - name: rootfs
>         hostPath:
>           path: /
> ```
>
> **说明：**
> 在 Kubernetes 下，Prometheus 通过与 Kubernetes API 集成，目前主要支持5种服务发现模式，分别是：Node、Service、Pod、Endpoints、Ingress。
>
> 通过指定 `kubernetes_sd_configs` 的模式为 `node`，Prometheus 会自动发现所有 node 节点并作为监控目标。默认访问端口是10250（kubelet 只读接口），但现在指标数据通常由 node-exporter 在端口9100提供。因此需要使用 `relabel_configs` 替换端口。
>
> **修改 prometheus 配置文件 prome-cm.yaml** (添加 nodes 和 kubelet job)
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: ConfigMap
> metadata:
>   name: prometheus-config
>   namespace: kube-ops
> data:
>   prometheus.yml: |
>     global:
>       scrape_interval: 30s
>       scrape_timeout: 30s
>     scrape_configs:
>     - job_name: 'prometheus'
>       static_configs:
>       - targets: ['localhost:9090']
>     - job_name: 'ingressnginx12'
>       static_configs:
>       - targets: ['192.168.66.12:10254']
>     - job_name: 'ingressnginx13'
>       static_configs:
>       - targets: ['192.168.66.13:10254']
>     - job_name: 'kubernetes-nodes' # 监控 node-exporter
>       kubernetes_sd_configs:
>       - role: node
>       relabel_configs:
>       - source_labels: [__address__]
>         regex: '(.*):10250'
>         replacement: '${1}:9100'
>         target_label: __address__
>         action: replace
>       - action: labelmap
>         regex: __meta_kubernetes_node_label_(.+)
>     - job_name: 'kubernetes-kubelet' # 监控 kubelet 本身指标 (K8s 1.11+ 使用 10250 https)
>       kubernetes_sd_configs:
>       - role: node
>       scheme: https
>       tls_config:
>         ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
>         insecure_skip_verify: true
>       bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
>       relabel_configs:
>       - action: labelmap
>         regex: __meta_kubernetes_node_label_(.+)
> ```
>
> ------
>
> ### 四、监控容器资源指标
>
> `cadvisor` 已内置在 `kubelet` 组件中，数据路径为 `/api/v1/nodes/<node>/proxy/metrics/cadvisor`。
>
> **修改 `prometheus` 配置文件，prome-cm.yaml** (添加 cadvisor job)
>
> yaml
>
> ```yaml
>     - job_name: 'kubernetes-cadvisor'
>       kubernetes_sd_configs:
>       - role: node
>       scheme: https
>       tls_config:
>         ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
>       bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
>       relabel_configs:
>       - action: labelmap
>         regex: __meta_kubernetes_node_label_(.+)
>       - target_label: __address__
>         replacement: kubernetes.default.svc:443
>       - source_labels: [__meta_kubernetes_node_name]
>         regex: (.+)
>         target_label: __metrics_path__
>         replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor
> ```
>
> ------
>
> ### 五、监控 ApiServer 指标
>
> **修改 prometheus 配置文件** (添加 apiserver job)
>
> yaml
>
> ```yaml
>     - job_name: 'kubernetes-apiservers'
>       kubernetes_sd_configs:
>       - role: endpoints
>       scheme: https
>       tls_config:
>         ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
>       bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
>       relabel_configs:
>       - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
>         action: keep
>         regex: default;kubernetes;https
> ```
>
> ------
>
> ### 六、通过 Service 监控服务
>
> **修改 prometheus，监控 Service** (添加 service-endpoints job)
>
> yaml
>
> ```yaml
>     - job_name: 'kubernetes-service-endpoints'
>       kubernetes_sd_configs:
>       - role: endpoints
>       relabel_configs:
>       - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_scrape]
>         action: keep
>         regex: true
>       - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_scheme]
>         action: replace
>         target_label: __scheme__
>         regex: (https?)
>       - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_path]
>         action: replace
>         target_label: __metrics_path__
>         regex: (.+)
>       - source_labels: [__address__, __meta_kubernetes_service_annotation_prometheus_io_port]
>         action: replace
>         target_label: __address__
>         regex: ([^:]+)(?::\d+)?;(\d+)
>         replacement: $1:$2
>       - action: labelmap
>         regex: __meta_kubernetes_service_label_(.+)
>       - source_labels: [__meta_kubernetes_namespace]
>         action: replace
>         target_label: kubernetes_namespace
>       - source_labels: [__meta_kubernetes_service_name]
>         action: replace
>         target_label: kubernetes_name
> ```
>
> 在 `relabel_configs` 中过滤了 annotation 有 `prometheus.io/scrape=true` 的 Service。
>
> ------
>
> ### 七、部署 kube-state-metrics，并监控
>
> **部署 kube-state-metrics**
>
> bash
>
> ```yaml
> # 部署kube-state-metrics，地址https://github.com/kubernetes/kube-state-metrics
> $ kubectl apply -f examples/standard
> ```
>
> **创建 kube-state-metrics 服务的 svc 文件，添加注解以便被自动发现**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: Service
> metadata:
>   annotations:
>     prometheus.io/scrape: 'true'
>     prometheus.io/port: "8080"
>   namespace: kube-system
>   labels:
>     app: kube-state-metrics
>   name: kube-state-metrics-exporter
> spec:
>   ports:
>   - name: 80-8080
>     port: 80
>     protocol: TCP
>     targetPort: 8080
>   selector:
>     app.kubernetes.io/name: kube-state-metrics
>   type: ClusterIP
> ```
>
> ------
>
> ### 八、部署 grafana 服务
>
> **创建 grafana 部署文件**
>
> yaml
>
> ```yaml
> apiVersion: apps/v1
> kind: Deployment
> metadata:
>   name: grafana
>   namespace: kube-ops
>   labels:
>     app: grafana
> spec:
>   revisionHistoryLimit: 10
>   selector:
>     matchLabels:
>       app: grafana
>   template:
>     metadata:
>       labels:
>         app: grafana
>     spec:
>       containers:
>       - name: grafana
>         image: grafana/grafana:5.3.4
>         imagePullPolicy: IfNotPresent
>         ports:
>         - containerPort: 3000
>           name: grafana
>         env:
>         - name: GF_SECURITY_ADMIN_USER
>           value: admin
>         - name: GF_SECURITY_ADMIN_PASSWORD
>           value: admin321
>         readinessProbe:
>           failureThreshold: 10
>           httpGet:
>             path: /api/health
>             port: 3000
>             scheme: HTTP
>           initialDelaySeconds: 60
>           periodSeconds: 10
>           successThreshold: 1
>           timeoutSeconds: 30
>         livenessProbe:
>           failureThreshold: 3
>           httpGet:
>             path: /api/health
>             port: 3000
>             scheme: HTTP
>           periodSeconds: 10
>           successThreshold: 1
>           timeoutSeconds: 1
>         resources:
>           limits:
>             cpu: 100m
>             memory: 256Mi
>           requests:
>             cpu: 100m
>             memory: 256Mi
>         volumeMounts:
>         - mountPath: /var/lib/grafana
>           subPath: grafana
>           name: storage
>         securityContext:
>           fsGroup: 472
>           runAsUser: 472
>       volumes:
>       - name: storage
>         persistentVolumeClaim:
>           claimName: grafana
> ```
>
> **创建 grafana 存储 pvc, grafana-volume.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: PersistentVolumeClaim
> metadata:
>   name: grafana
>   namespace: kube-ops
> spec:
>   storageClassName: nfs-client
>   accessModes:
>   - ReadWriteOnce
>   resources:
>     requests:
>       storage: 1Gi
> ```
>
> **创建 grafana svc 文件, grafana-svc.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: Service
> metadata:
>   name: grafana
>   namespace: kube-ops
>   labels:
>     app: grafana
> spec:
>   type: NodePort
>   ports:
>   - port: 3000
>   selector:
>     app: grafana
> ```
>
> **创建 job，调整 grafana 挂载目录权限，grafana-chown-job.yaml**
>
> yaml
>
> ```yaml
> apiVersion: batch/v1
> kind: Job
> metadata:
>   name: grafana-chown
>   namespace: kube-ops
> spec:
>   template:
>     spec:
>       restartPolicy: Never
>       containers:
>       - name: grafana-chown
>         command: ["chown", "-R", "472:472", "/var/lib/grafana"]
>         image: busybox
>         imagePullPolicy: IfNotPresent
>         volumeMounts:
>         - name: storage
>           subPath: grafana
>           mountPath: /var/lib/grafana
>       volumes:
>       - name: storage
>         persistentVolumeClaim:
>           claimName: grafana
> ```
>
> ------
>
> ### 九、监控 metrics.server
>
> 从 Kubernetes v1.8 开始，资源使用情况的监控可以通过 Metrics API 获取，由 Metrics Server 组件提供，用来替换 heapster。
>
> **部署 metrics-server**
>
> bash
>
> ```yaml
> # https://github.com/kubernetes-sigs/metrics-server
> kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
> ```
>
> **(components.yaml 内容较长，主要包含 ServiceAccount, ClusterRole, ClusterRoleBinding, RoleBinding, Service, Deployment, APIService 等资源)**
>
> ------
>
> ### 十、alertmanager 部署
>
> **创建 alertmanager 配置文件，alertmanager-conf.yaml**
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: ConfigMap
> metadata:
>   name: alert-config
>   namespace: kube-ops
> data:
>   config.yml: |-
>     global:
>       # 在没有报警的情况下声明为已解决的时间
>       resolve_timeout: 5m
>       # 配置邮件发送信息
>       smtp_smarthost: 'smtp.163.com:25'
>       smtp_from: 'wangyanglinux@163.com'
>       smtp_auth_username: 'wangyanglinux@163.com'
>       smtp_auth_password: 'APYDOSTDPDUOEEHQ' # 注意：此为示例，应使用有效密码或凭据
>       smtp_hello: '163.com'
>       smtp_require_tls: false
>     # 所有报警信息进入后的根路由，用来设置报警的分发策略
>     route:
>       # 这里的标签列表是收到报警信息后的重新分组标签，例如，接收到的报警信息里面有许多具有 cluster=A 和 alertname=latencyhigh 这样的标签的报警信息将会批量被聚合到一个分组里面
>       group_by: ['alertname', 'cluster']
>       # 当一个新的报警分组被创建后，需要等待至少 group_wait 时间来初始化通知，这种方式可以确保您能有足够的时间为同一分组来获取多个警报，然后一起触发这个报警信息。
>       group_wait: 30s
>       # 当第一个报警发送后，等待'group_interval'时间来发送新的一组报警信息。
>       group_interval: 5m
>       # 如果一个报警信息已经发送成功了，等待'repeat_interval'时间来重新发送他们
>       repeat_interval: 5m
>       # 默认的 receiver：如果一个报警没有被一个 route 匹配，则发送给默认的接收器
>       receiver: default
>       # 上面所有的属性都由所有子路由继承，并且可以在每个子路由上进行覆盖。
>       routes:
>       - receiver: email
>         group_wait: 10s
>         match:
>           team: node
>     receivers:
>     - name: 'default'
>       email_configs:
>       - to: 'wangyanglinux@88.com' # 示例邮箱
>         send_resolved: true
>     - name: 'email'
>       email_configs:
>       - to: 'wangyanglinux@88.com' # 示例邮箱
>         send_resolved: true
> ```
>
> **修改 prometheus 配置文件，添加 alerting 配置和 rule_files** (prometheus-cm.yaml)
>
> yaml
>
> ```yaml
> # ... (之前的 scrape_configs)
> alerting:
>   alertmanagers:
>   - static_configs:
>     - targets: ["localhost:9093"] # 假设alertmanager与prometheus同Pod
> 
> rule_files:
>   - /etc/prometheus/rules.yml
> ```
>
> **修改 prometheus service 文件，暴露 alertmanager 端口** (prometheus-svc.yaml)
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: Service
> metadata:
>   name: prometheus
>   namespace: kube-ops
>   labels:
>     app: prometheus
> spec:
>   selector:
>     app: prometheus
>   type: NodePort
>   ports:
>   - name: web
>     port: 9090
>     targetPort: http
>   - name: alertmanager # 添加此端口
>     port: 9093
>     targetPort: 9093
> ```
>
> **合并 alertmanager 至 prometheus deploy 文件, prometheus-alertmanager-deploy.yaml** (将alertmanager容器添加到Prometheus的Deployment中)
>
> yaml
>
> ```yaml
> apiVersion: apps/v1
> kind: Deployment
> metadata:
>   name: prometheus
>   namespace: kube-ops
>   labels:
>     app: prometheus
> spec:
>   selector:
>     matchLabels:
>       app: prometheus
>   template:
>     metadata:
>       labels:
>         app: prometheus
>     spec:
>       serviceAccountName: prometheus
>       containers:
>       - name: alertmanager # Alertmanager 容器
>         image: prom/alertmanager:v0.15.3
>         imagePullPolicy: IfNotPresent
>         args:
>         - "--config.file=/etc/alertmanager/config.yml"
>         - "--storage.path=/alertmanager/data"
>         ports:
>         - containerPort: 9093
>           name: alertmanager
>         volumeMounts:
>         - mountPath: "/etc/alertmanager"
>           name: alertcfg
>         resources:
>           requests:
>             cpu: 100m
>             memory: 256Mi
>           limits:
>             cpu: 100m
>             memory: 256Mi
>       - name: prometheus # 原有的 Prometheus 容器
>         image: prom/prometheus:v2.4.3
>         # ... (原有 Prometheus 容器配置，确保包含 --web.enable-lifecycle)
>         volumeMounts:
>         # ... (原有 volumeMounts)
>         - mountPath: "/etc/alertmanager" # Prometheus 容器也可能需要访问 alertmanager 配置？通常不需要，这里根据实际情况调整
>           name: alertcfg
>           readOnly: true
>       volumes:
>       - name: data
>         persistentVolumeClaim:
>           claimName: prometheus
>       - name: config-volume
>         configMap:
>           name: prometheus-config
>       - name: alertcfg # Alertmanager 配置 volume
>         configMap:
>           name: alert-config
> ```
>
> **添加报警规则演示** (创建包含 rules.yml 的 ConfigMap)
>
> yaml
>
> ```yaml
> apiVersion: v1
> kind: ConfigMap
> metadata:
>   name: prometheus-config
>   namespace: kube-ops
> data:
>   rules.yml: | # 报警规则文件
>     groups:
>     - name: test-rule
>       rules:
>       - alert: NodeMemoryUsage
>         expr: (node_memory_MemTotal_bytes - (node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes)) / node_memory_MemTotal_bytes * 100 > 20
>         for: 2m
>         labels:
>           team: node
>         annotations:
>           summary: "{{$labels.instance}}: High Memory usage detected"
>           description: "{{$labels.instance}}: Memory usage is above 20% (current value is: {{ $value }})"
>   prometheus.yml: | # 主配置文件
>     global:
>       scrape_interval: 30s
>       scrape_timeout: 30s
>     scrape_configs:
>     # ... (原有的 scrape_configs 内容)
>     alerting:
>       alertmanagers:
>       - static_configs:
>         - targets: ["localhost:9093"]
>     rule_files:
>     - /etc/prometheus/rules.yml
> ```



固定IP地址到pod

NetworkPolicy

