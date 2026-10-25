# 01 · 把 channel 的行为真正想清楚

## 1. channel 解决什么问题

`goroutine` 负责并发执行；`channel` 让这些执行单元交换有类型的数据，并在通信点协调进度。`go f()` 启动 goroutine，创建 channel 本身不会启动任何执行单元。

例如：HTTP 请求进入后，把任务交给后台 worker，再把处理结果送回来。这里有三个问题：任务怎么传递？处理能力不足时怎么办？请求取消后怎么办？channel 可以参与解决这三个问题，但仍需要明确的生命周期设计。

并发表示任务在结构上可以交错推进；并行表示任务在同一时刻执行。用了 channel 不保证更快，也不自动保证并行。

## 2. 最小语法：创建、发送、接收

```go
ch := make(chan string) // 无缓冲
go func() {
    ch <- "hello"     // 发送
}()
message := <-ch       // 接收
fmt.Println(message)
```

这里不能在同一个 goroutine 中先 `ch <- "hello"` 再 `<-ch`：第一步就会等待接收者，程序走不到第二步。

```go
var a chan int        // nil，尚未创建可用 channel
b := make(chan int)   // 可用的无缓冲 channel
c := make(chan int, 3)// 可用的有缓冲 channel，容量为 3
```

channel 的元素类型可以是结构体、指针、切片、接口，甚至另一个 channel。优先使用能表达业务含义的类型，例如 `chan Job`、`chan Result`，少用需要到处断言的 `chan any`。

赋值和传参会复制 channel 值，两个值仍引用同一个通信对象：

```go
a := make(chan int, 1)
b := a
b <- 7
fmt.Println(<-a) // 7
fmt.Println(a == b) // true
```

## 3. 无缓冲：一次会合

`make(chan T)` 的容量为 0。发送方需要等到接收方参与，接收方需要等到有发送方参与。

```text
发送方                 接收方
  ch <- task  ────────>  task := <-ch
  继续执行               继续执行
```

关键区别：**发送完成只代表交接完成，不代表接收方的后续处理完成。**

```go
go func() {
    job := <-jobs
    save(job) // 可能还需要 5 秒
}()
jobs <- job
// 这里不能据此认为 save 已完成。
```

如果业务要等“保存成功”，必须额外返回确认结果，或等待代表完成的信号。

同样，发送之后打印“已发送”和接收之后打印“已接收”，两个打印的先后顺序不由这次交接决定。

## 4. 有缓冲：有限队列

```go
ch := make(chan int, 2)
ch <- 10 // 不需要已有接收者
ch <- 20 // 不需要已有接收者
// ch <- 30 // 缓冲已满，会等待
fmt.Println(<-ch) // 10
ch <- 30         // 有一个空位，可以发送
```

空队列的接收者等待；满队列的发送者等待。缓冲能吸收短暂突发，却不能解决“长期生产速度高于消费速度”。容量为 100 的队列面对每秒多积压 50 个任务，大约两秒就会填满。

这种等待可以把压力传回上游，称为背压。你仍需决定满了以后是等待、超时、拒绝，还是允许丢弃。把所有发送都包进新的 goroutine，会把压力转变成不断增长的 goroutine 数量。

容量通常不能动态调整。选容量要考虑每项保留的内存、可接受排队时间、突发规模、worker 的处理能力，不能只凭“越大越好”。元素中的指针还可能间接保留大块内存。

## 5. 一张必须会解释的行为表

下表讨论直接收发；在 `select` 中，不能立即完成的操作不会成为已就绪的通信分支。

| channel 状态 | 发送 | 接收 | close |
|---|---|---|---|
| nil | 永久等待 | 永久等待 | panic |
| 无缓冲、打开 | 等待配对接收 | 等待配对发送 | 成功一次 |
| 有缓冲、打开 | 有空位则完成，否则等待 | 有值则完成，否则等待 | 成功一次 |
| 已关闭、尚有缓冲值 | panic | 先取出缓冲值，`ok=true` | panic |
| 已关闭、已经排空 | panic | 立刻返回零值，`ok=false` | panic |

`len(ch)` 是当前排队的元素数，`cap(ch)` 是容量。nil channel 的这两个结果都是 0；无缓冲 channel 的 `len` 也是 0，不代表没人等着收发。

这些是语言规则，可查阅 [Go 规范中的 channel 定义](https://go.dev/ref/spec#Channel_types)、[发送](https://go.dev/ref/spec#Send_statements)、[接收](https://go.dev/ref/spec#Receive_operator)和[关闭](https://go.dev/ref/spec#Close)。

## 6. close 是“不再发送”的声明

```go
ch := make(chan int, 2)
ch <- 0
ch <- 9
close(ch)

a, okA := <-ch // 0, true：真实数据
b, okB := <-ch // 9, true：缓冲里的数据
c, okC := <-ch // 0, false：关闭且排空
```

`ok` 的含义不是“接收时 channel 是否打开”，而是“这次是否接收到实际发送的值”。关闭后读缓冲中的值，`ok` 仍为 true。

`close` 不会清空缓冲，不会释放仍在使用的数据，也不等待消费者处理完成。它还会让阻塞的接收者有机会继续；在已经没有缓冲数据时，这些接收返回零值。已经阻塞的发送者会因关闭而发生 panic。

工程上通常让**拥有发送生命周期的一方**关闭 channel：

```go
go func() {
    defer close(out)
    // 该 goroutine 是 out 的唯一发送者。
    for _, n := range values {
        out <- n
    }
}()
```

这不是语言强制的“只有发送者身份才能 close”。语言不记录 goroutine 的身份；只要类型允许，接收方也能调用 close，但若还有发送者就可能触发 panic。

多个发送者共用输出时，让协调者等待它们全部退出，再关闭一次：

```go
wg.Add(workerCount) // 必须在启动等待之前完成正数 Add
for i := 0; i < workerCount; i++ {
    go func() {
        defer wg.Done()
        // 发送到 out，所有退出路径都经过 Done。
    }()
}
go func() {
    wg.Wait()
    close(out)
}()
```

`sync.Once` 可以避免重复执行 `close`，但不能防止另一个 goroutine 同时发送。必须先有“不会再发送”的协议。

不需要对每个 channel 都 close。固定发送一次、接收一次，或由 context 管理长期任务时，可以不 close。close 的主要用途是表达协议结束；可达性与垃圾回收不是靠 close 决定的。

## 7. range：一直读到关闭且排空

```go
for n := range ch {
    fmt.Println(n)
}
```

它会读完关闭前入队的值再退出。队列暂时空了不会退出；nil channel 会一直等待。不关闭输入又没有其他退出条件，是典型的卡住原因。

需要支持取消时，不能只写 `for range` 再在循环体里检查 context：如果卡在下一次接收，就进不了循环体。应将接收本身放入 `select`。

## 8. 单向 channel：把权限写进 API

```go
func produce(out chan<- int) { // 可以发送，也可以 close；不能接收
    out <- 1
    close(out)
}

func consume(in <-chan int) { // 可以接收；不能发送，也不能 close
    for n := range in {
        fmt.Println(n)
    }
}
```

一个双向 `chan int` 可以传给上述两个函数。单向类型限制调用方能执行的操作，不会创建第二条 channel，也不是深拷贝。

常见 API 返回 `<-chan Result`，表示调用者只消费结果。需要关闭它的责任保留在实现内部。

## 9. select：等待任意一个通信条件

```go
select {
case value, ok := <-input:
    // 收到数据或发现输入已经结束。
    _ = value
    _ = ok
case output <- result:
    // 成功交付一个结果。
case <-ctx.Done():
    return
}
```

每次只执行一个分支。多个分支同时可执行时，语言规定进行均匀伪随机选择；不是按代码顺序挑选，也不承诺某个分支在有限次循环内一定被选中。没有就绪分支且没有 default 时等待。

```go
select {
case queue <- job:
    return nil
default:
    return ErrQueueFull
}
```

`default` 让这次选择不等待。它适合明确的“立即尝试”策略。放在无限循环中而没有其他等待，常会导致空转占用 CPU。

`select` 并不保护向已关闭的 channel 发送：这种发送会 panic，不能靠 default 绕开。

### nil 可以动态禁用分支

```go
for a != nil || b != nil {
    select {
    case v, ok := <-a:
        if !ok {
            a = nil
            continue
        }
        handle(v)
    case v, ok := <-b:
        if !ok {
            b = nil
            continue
        }
        handle(v)
    }
}
```

这个片段假定 `a`、`b` 最终都会关闭，`handle` 能完成。生产代码如需取消，还应加 `ctx.Done()` 分支。已关闭的 channel 会持续就绪，不能一直留在 select 里反复读零值。

### case 里的表达式也有执行时机

```go
select {
case out <- expensiveCalculation():
case <-ctx.Done():
}
```

进入 select 时就会求发送表达式的值，即使最后选中取消分支。所有通信分支的 channel 表达式和发送右侧表达式都会按规则求值一次；它们本身若很慢或阻塞，会延迟进入选择。因此不要把不可取消的耗时调用藏在 case 中。[官方 select 语义](https://go.dev/ref/spec#Select_statements)

## 10. FIFO 不等于并发结果有序

一个发送者按 `1、2、3` 顺序发送，一个接收者会按这个顺序接收。

多个发送者并发发送时，不能从源代码排列推断它们的交错顺序；多个 worker 收任务后，完成顺序也可能不同。即使接收交接有先后，后续打印也可能反过来。

如果业务需要结果保持输入顺序，给任务加 `Index`，在汇总端按索引排序或用一个有界窗口重排。代价是额外内存，以及慢任务可能阻塞后面的交付。

同一个数据 channel 上的多个接收者是竞争消费，每个发送的值只匹配一个接收；并不会每人收到一份。广播业务事件需要显式分发机制。关闭一个信号 channel 则可以被多个等待者共同观察。

## 11. channel 传的是值，对象不一定独立

```go
type Job struct {
    ID   int
    Data []byte
}

jobs <- Job{ID: 1, Data: payload}
// 如果马上修改 payload，消费者读 Data 时可能发生数据竞争。
```

结构体被按值复制，但 slice 只复制描述信息，底层数组仍共享；map、指针指向的数据也不会自动深拷贝。

三种常见办法：发送后移交修改权，原发送者不再碰对象；发送前复制需要独立的数据；或对共享访问使用 Mutex 等同步手段。约定要涵盖接收之后的行为，不能只考虑发送那一刻。

channel 自身支持多个 goroutine 并发收发，并不意味着通过它传递的对象自动线程安全。

## 12. channel 也是内存同步点

```go
done := make(chan struct{})
var config string
go func() {
    config = "ready"
    close(done)
}()
<-done
fmt.Println(config) // 此处可以观察到关闭之前的写入
```

对应的发送与接收、以及关闭与观察到关闭的接收，建立了内存模型规定的同步关系。不要把 `time.Sleep`、`len(ch)` 或“看起来它应该已经跑完了”当成同步。

但缓冲发送返回并不能证明消费者已经执行。无缓冲 channel 还有相反方向的交接同步保证，不能随意把它改成容量 1 后仍保留同样推理。工作中优先画清“写入 → 通信 → 读取”的顺序，而不是凭时序猜测。[Go 内存模型](https://go.dev/ref/mem#Channel_communication)

## 13. 取消请求与完成确认是两件事

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

done := make(chan struct{})
go func() {
    defer close(done)
    for {
        select {
        case <-ctx.Done():
            return
        case job, ok := <-jobs:
            if !ok {
                return
            }
            process(ctx, job)
        }
    }
}()

cancel() // 请求停止
<-done  // 确认该 worker 已经退出
```

`process` 本身也需要支持取消或保证能在可接受的时间内返回。context 无法中断任意函数。`cancel()` 返回，不等于所有使用它的 goroutine 都已完成清理。[context 文档](https://pkg.go.dev/context)

当数据分支与取消分支同时就绪，select 可能选数据。循环前检查 `ctx.Err()` 可以减少取消后的额外工作，但检查与发送之间仍不是原子的。需要“取消返回后绝无新任务启动”这样的强约束时，要额外设计入口关闭、状态锁与完成屏障。

## 14. 不要用这些办法猜状态

| 看起来方便的写法 | 为什么靠不住 | 更合适的处理 |
|---|---|---|
| `if len(ch) < cap(ch) { ch <- x }` | 判断后别的发送者可能抢走空位 | `select` 加 default，或带取消的阻塞发送 |
| `if len(ch) == 0 { /* finished */ }` | 未来还可能有新值；也可能有在途处理 | close 协议 + 完成确认 |
| 先 `isClosed(ch)` 再发送 | 检查和发送之间可能被关闭 | 通过所有权与等待协调关闭 |
| 用接收来检测是否关闭 | 可能消费掉正常数据，改变业务 | 在正常消费路径用 `v, ok` |
| 向 done 发送一个值以“通知所有人” | 通常只会交给一个接收者 | 唯一负责人 `close(done)` |
| `recover` 包住所有 send/close | 掩盖了错误的生命周期协议 | 修复关闭责任与发送退出条件 |
| 给队列加很大缓冲 | 只是推迟满载，增加排队与内存 | 限制生产、超时、拒绝、扩容消费能力 |

## 15. 什么时候该换别的工具

| 需求 | 常用起点 |
|---|---|
| 在阶段之间传递任务、结果、流式数据 | channel |
| 保护一个共享 map 或一组关联状态 | Mutex / RWMutex |
| 等待若干 goroutine 完成 | WaitGroup |
| 传播请求取消、截止时间 | context |
| 简单独立计数器 | atomic，或 Mutex |
| 跨进程、持久化、重试、确认投递 | 消息系统 / 数据库队列 |

它们常常组合使用。例如 worker 用 channel 收任务，用 WaitGroup 确认退出，用 context 取消 I/O，用 Mutex 更新共享缓存。不要把所有同步需求都改写成 channel。[sync 标准库](https://pkg.go.dev/sync)
