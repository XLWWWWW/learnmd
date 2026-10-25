# 03 · 练习：预测、实现、设计

先做题，再看提示，最后才看答案。推理题每题建议写下：是否能编译、是否可能阻塞或 panic、输出范围、涉及的保证。片段默认放在 `main` 中，省略 import；未特别注明时没有其他 goroutine。

“主 goroutine 卡住”不总会出现同样的运行时提示：最小程序可能报全局 deadlock，而长时间运行的服务可能只是泄漏某几个 goroutine。

## A · 24 道推理与排错题

### Q01 · 无缓冲的第一次发送

```go
ch := make(chan int)
ch <- 1
fmt.Println(<-ch)
```

能打印 1 吗？写出两种合理修改，并说明它们的同步语义有何不同。

### Q02 · 缓冲有上限

```go
ch := make(chan int, 1)
ch <- 1
ch <- 2
fmt.Println(<-ch)
```

停在哪一行？把容量改成 2 后是否还需要 close 才能接收？

### Q03 · 真正的零值与结束

```go
ch := make(chan int, 2)
ch <- 0
ch <- 8
close(ch)
for i := 0; i < 4; i++ {
    v, ok := <-ch
    fmt.Println(v, ok)
}
```

逐行写出输出。哪一次接收说明 channel 关闭且排空？

### Q04 · range 什么时候结束

```go
ch := make(chan int, 1)
ch <- 9
for n := range ch {
    fmt.Println(n)
}
fmt.Println("finished")
```

能打印 finished 吗？消费者能否在循环里随意 `close(ch)` 来修复？

### Q05 · nil、空、关闭

```go
var ch chan int
select {
case v := <-ch:
    fmt.Println(v)
default:
    fmt.Println("default")
}
```

输出什么？如果 ch 是已关闭的空 channel 呢？`close(ch)` 在原代码中会怎样？

### Q06 · 赋值是否复制了队列

```go
a := make(chan int, 1)
b := a
b <- 5
fmt.Println(<-a)
close(a)
_, ok := <-b
fmt.Println(ok)
```

写出输出，解释 a 和 b 的关系。

### Q07 · 单向权限

```go
func f(out chan<- int) {
    out <- 1
    close(out)
}
func g(in <-chan int) {
    close(in)
}
```

哪部分不能编译？`chan<-` 是否意味着只能发、不能 close？

### Q08 · 竞争消费还是广播

两个 worker 都执行 `for job := range jobs`。生产者向 jobs 发送一次 Job：两个 worker 会各收到一次吗？如果要让十个等待者共同观察“初始化完成”，你会怎样做？

### Q09 · select 的顺序

```go
a := make(chan int, 1)
b := make(chan int, 1)
a <- 1
b <- 2
select {
case <-a:
    fmt.Println("A")
case <-b:
    fmt.Println("B")
}
```

A 是否一定先？如果把第一个 case 改成已经关闭的 `ctx.Done()`，取消是否一定胜出？

### Q10 · 看起来“没有阻塞”

```go
for {
    select {
    case job := <-jobs:
        process(job)
    default:
    }
}
```

jobs 空着但没关闭时，会发生什么？关闭后又有什么问题？写出带取消、能正确处理关闭的版本。

### Q11 · break 跳出了哪里

```go
for {
    select {
    case <-done:
        break
    case n := <-input:
        use(n)
    }
}
```

done 关闭后循环会退出吗？给出两种修复。

### Q12 · 关闭输入的分支

一个 select 监听 a、b。a 先关闭，b 之后才会继续产生数据。如果不检查 `ok`，a 分支会怎样？如何只禁用 a 分支而继续处理 b？什么时候整个循环才能结束？

### Q13 · len 检查

```go
if len(jobs) < cap(jobs) {
    jobs <- job
}
```

两个生产者并发执行，能保证都不阻塞吗？对 channel 调用 len 本身是否就是数据竞争？写出真正不等待的提交逻辑。

### Q14 · 能否先检测关闭再发送

有人写 `if !isClosed(ch) { ch <- value }`。即使 isClosed 的实现正确，为什么仍不安全？用 `sync.Once` 包住 close 就足够了吗？

### Q15 · channel 安全是否等于对象安全

```go
data := []int{1, 2, 3}
ch := make(chan []int, 1)
ch <- data
go func() { fmt.Println((<-ch)[0]) }()
data[0] = 99
```

假定主 goroutine 随后正确等待子 goroutine 结束，这段程序仍有什么问题？给出两种修复，并写出所有权约定。

### Q16 · 不靠 Sleep 的可见性

```go
done := make(chan struct{})
var value int
go func() {
    value = 17
    close(done)
}()
<-done
fmt.Println(value)
```

是否保证打印 17？去掉 `<-done`，换成 Sleep 一秒能否得到同样保证？

### Q17 · 交接与处理完成

```go
jobs := make(chan int)
go func() {
    n := <-jobs
    slowSave(n)
}()
jobs <- 7
fmt.Println("saved")
```

saved 打印时是否保证 slowSave 已经结束？如何让打印表达的含义准确？

### Q18 · 容量改了，证明还成立吗

```go
ch := make(chan struct{}, 1)
var x int
go func() {
    x = 7
    <-ch
}()
ch <- struct{}{}
fmt.Println(x)
```

是否存在数据竞争？把容量改成 0 后，结果有什么变化？请用通信前后的顺序解释，不要用调度速度解释。

### Q19 · 多个发送者都 defer close

三个 worker 共享 out，每个 worker 都写 `defer close(out)`。为什么错？谁应该关闭，关闭前需要证明什么？

### Q20 · 等完再收集

```go
out := make(chan int)
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    out <- 1
}()
wg.Wait()
close(out)
for n := range out {
    fmt.Println(n)
}
```

画出循环等待关系。如何在不猜缓冲容量的情况下修复？

### Q21 · 调用方超时之后

```go
result := make(chan int)
go func() {
    result <- slowCompute()
}()
select {
case n := <-result:
    use(n)
case <-ctx.Done():
    return
}
```

假定 slowCompute 之后会返回，但调用方先超时，子 goroutine 会停在哪里？容量 1 能解决哪些问题，不能解决哪些问题？

### Q22 · 总超时还是空闲超时

```go
for {
    select {
    case _, ok := <-events:
        if !ok { return }
    case <-time.After(time.Second):
        return
    }
}
```

events 每 100ms 来一个事件，这个循环是否会在启动一秒后结束？如何实现整个循环最多运行一秒？

### Q23 · 没被选中的 case 是否会调用函数

```go
ctx, cancel := context.WithCancel(context.Background())
cancel()
var out chan int
makeValue := func() int {
    fmt.Println("computed")
    return 1
}
select {
case out <- makeValue():
case <-ctx.Done():
    fmt.Println("canceled")
}
```

逐行写出输出。若 makeValue 内部执行一个阻塞调用，context 能跳过它吗？

### Q24 · 消费者提前结束

你写了一条 `Generate → Map → Merge → Take(1)` 流水线。Take 正常结束了，上游却还没完成。由谁取消？每一层的哪些操作必须能响应取消？为什么 Take 不应该直接 close 它的输入？

## B · 10 道可运行编程题

编辑 `lab/exercises.go`。每题测试命令中的前缀会同时运行该题的子测试，例如 `go test ./lab -run '^TestE08' -v -timeout 15s`。

所有带 ctx 的函数都要求 ctx 非 nil，所有实际使用的 fn/job 都要求非 nil。取消与通信同时就绪时允许发生少量交付，不要求严格优先级。每题都禁止通过 Sleep 安排正确性、关闭不属于自己的输入，或启动没有退出路径的 goroutine。

### E01 · 汇总一个有限流 ★

实现 `Sum(in <-chan int) int`。读到输入关闭且排空后返回总和；空输入返回 0；真实零值和负数都算数据。不能改用固定次数接收。nil 输入会一直等待，这道基础题不要求取消。

验收：Q03、Q04 能解释；`TestE01` 通过；能说明为什么函数不负责 close(in)。

### E02 · 非阻塞提交 ★

实现 `TrySend(out chan<- int, n int) bool`。能立即发送则 true，否则 false；nil、满缓冲、无人接收的无缓冲 channel 都不能让函数等待。

前置约定：调用期间 out 不会被关闭。不要使用 len 先判断。向已关闭 channel 发送不在该 API 的保护范围内。

验收：`TestE02` 通过；解释为什么 false 只说明“这次没有立即发送”，不能证明永远没有消费者。

### E03 · 可取消的接收 ★★

实现 `Receive(ctx, in) (int, bool, error)`：

| 条件 | 返回 |
|---|---|
| 函数进入前 ctx 已经取消 | `0, false, ctx.Err()` |
| 收到真实值 | `value, true, nil` |
| 输入关闭且读空 | `0, false, nil` |
| 等待期间取消 | `0, false, ctx.Err()` |

nil 输入只禁用输入分支，仍要响应取消。并发到达的数据与取消允许任一胜出。

验收：`TestE03` 通过；不能把“输入关闭”当作“超时”返回。

### E04 · 可取消生产者 ★★

实现 `Generate(ctx, nums ...int) <-chan int`。创建自己的输出，在 goroutine 中按顺序发送；正常完成或取消都关闭输出。空参数直接结束。为简化所有权，返回前复制 nums；调用者返回后可以修改原切片。

验收：`TestE04` 通过，包括调用方修改原切片后输出仍保持原值；输出无人消费时取消也有退出路径。扩展：设计一个不复制、而是明确移交输入所有权的版本，说明调用约束。

### E05 · 可取消转换器 ★★

实现 `Map(ctx, in, fn) <-chan int`。正常情况下每项调用 fn 一次，保持顺序；输入结束后输出结束；输入等待、输出等待都支持取消。

fn 约定为非 nil、短时间完成的计算。函数不能强制中断 fn；不得偷偷另开一个无限等待的 goroutine 来假装实现“中断”。

验收：`TestE05` 通过；用自己的语言解释为何只有发送使用 select、接收写 range 仍不够。

### E06 · 合并多个流 ★★★

实现 `Merge(ctx, inputs ...<-chan int) <-chan int`：每个输入启动一个转发者；不同输入间顺序不限，同一输入内部顺序保留；所有输入结束后关闭输出；取消时停止所有转发者；零输入直接结束。

nil 输入按照正常 nil channel 处理：会等待，直到 ctx 取消。不要把 nil 自动当作“已经关闭”，以免悄悄改变契约。

验收：`TestE06` 通过；明确指出唯一关闭者；证明 WaitGroup 不会与正数 Add 发生不正确的先后关系。

### E07 · 只取前 N 项 ★★

实现 `Take(ctx, in, n) <-chan int`：最多转发 n 项；输入更早结束则提前结束；n<=0 不读取输入，直接关闭输出；收发都支持取消。

函数不得 close(in)，也不自行创建或取消调用者的 context。组织流水线的上层负责在提前停止后通知上游。

验收：`TestE07` 通过；把 Generate、Map、Take 串起来，只取两个，然后取消整个流水线。

### E08 · 固定 worker pool ★★★★

实现 `ParallelMap(ctx, in, workers, fn) (<-chan Result, error)`。允许改变输出顺序。

- workers<=0：立即返回 `nil, ErrInvalidWorkers`，不能启动工作 goroutine。
- 正常运行：最多 workers 个 fn 同时执行；每个输入恰好产生一个 Result，包括出错项。
- 单项失败写入 Result.Err，继续处理其余项；用 Result.Input 关联输入。
- 输入关闭：完成已领取任务，发完结果，所有 worker 退出后关闭输出。
- 取消：各级允许放弃未交付结果并尽快退出；fn 必须配合 ctx，否则无法保证退出时间。
- 固定数量 worker 直接竞争输入，不要为每项任务创建一个新 goroutine。

验收：`TestE08` 通过；真实并发能到达 workers，又不会超过；回调都退出后才能关闭输出。

扩展：加入 Index 保持输入顺序；再思考慢任务导致的结果缓存上限。

### E09 · 第一个成功结果 ★★★★

实现 `FirstSuccess(ctx, jobs...) (int, error)`。每个 job 最多执行一次。

- 已取消的父 context：优先返回 ctx.Err，不启动任务。
- 非取消情况下的空 jobs：返回 ErrNoJobs。
- 第一个成功结果立即返回；某个任务先失败不能阻止其他任务成功。
- 全部失败：用 `errors.Join` 汇总，`errors.Is` 能识别各原始错误。
- 返回前请求取消其余任务；其他任务稍后返回时不能堵在结果发送上。
- 不要求函数返回时所有清理都已结束；job 必须响应传入的子 context。

验收：`TestE09` 通过；解释结果 channel 的容量为什么与任务数相等，以及为什么不必 close。

扩展：设计一个“返回时所有任务都已结束”的变体，并解释代价。再设计并发数有界的变体。

### E10 · 大小或时间触发批量 ★★★★

实现 `Batch(ctx, in, size, ticks) (<-chan []int, error)`：

- size<=0：返回 `nil, ErrInvalidBatchSize`。
- 满 size 立即提交；有未提交数据时收到 tick 也提交；空批次不提交。
- 输入正常关闭：提交最后的非空尾批，再关闭输出。
- 取消：允许丢弃未交付批次，结束输出；阻塞提交也必须能响应取消。
- ticks 为 nil：只按数量/输入关闭刷新。ticks 关闭后，将该分支禁用，继续接收输入。
- 每个已提交切片的内容此后不能被内部改写。

验收：`TestE10` 通过；测试通过手动发送 tick 控制时机，不依赖真实 Sleep。解释周期窗口与首项到达后超时的区别。

扩展：改成“收到第一项后最多等待 D”，用可注入的 timer 接口测试重置与停止。

## C · 3 个工作型综合项目

这些是开放设计题，没有预置实现或自动验收测试。先写 10～20 行设计说明，再实现，并自行补上验收测试。参考设计方向在答案文件中。

### P1 · 有界任务服务与优雅退出

实现一个内存任务服务，使用 4 个 worker、容量 16 的任务队列、独立结果收集者。任务 ID 唯一，模拟处理可以通过测试 gate 控制。

验收要求：

1. 提交支持 context；队列满时等待或取消，不无限创建 goroutine。
2. 运行中可以并发提交；返回 nil 只表示服务已接收任务。
3. `Shutdown` 开始后拒绝新提交；已经确认接收的任务正常排空。
4. 生产者彻底停止发送后才 close(jobs)；worker 都结束后才 close(results)。
5. 结果收集结束后 Shutdown 才能返回成功，重复 Shutdown 有定义。
6. Shutdown 有最大等待时间；超时返回时说明哪些工作可能仍在进行，提供强制取消路径。
7. 测试：满队列、并发提交与关闭、单项失败、排空成功、排空超时，不出现 send-on-closed panic。

设计挑战：不能让 Submit 持有状态锁阻塞在发送上，再让 Shutdown 等同一把锁且 worker 早已停止。需要认真设计入口关闭和在途提交的登记协议。

### P2 · 并发查询与有序输出

输入 100 个 ID，最多 5 个并发查询。输出必须与输入顺序相同。单项错误保留原 ID 与 Index，整个请求最多 2 秒，取消时所有配合取消的任务退出。

验收要求：结果不丢、不重复、顺序正确；最大并发<=5；查询使用 ctx；消费者提前退出后能清理。让第 0 项故意很慢，观察汇总缓存能否增长到接近全部结果。

进阶：把重排窗口限制为 10 项，说明这个限制如何反向影响领取任务速度，以及吞吐与内存的权衡。

### P3 · 可测试的日志批量提交器

每 100 条或每 200ms 刷新，允许在入口过载时丢弃调试级日志，但错误级日志必须返回明确提交结果。关闭时刷新尾批。

验收要求：不复用仍被下游读取的数组；记录丢弃数、队列长度、提交耗时和错误；注入时钟与 writer；writer 失败时有明确重试/放弃策略；停止时等待刷新与写入完成。

设计挑战：不要把“成功进入 channel”当作“成功写入目标”。声明其进程内语义，若要求崩溃不丢，补充持久化设计。

## D · 如何判断自己真的会了

| 能力 | 自测方式 |
|---|---|
| 解释行为 | 不运行就能解释 Q01～Q09 的结果 |
| 发现错误 | Q10～Q24 指出阻塞点、退出责任与修复边界 |
| 独立实现 | E01～E10 不看答案通过相应测试与 -race |
| 解释生命周期 | 为每个 go 语句写出正常结束与取消结束路径 |
| 迁移到工作 | 至少完成 P1，并能说明队列满、下游退出、关闭竞态如何处理 |

每题做完记录三个问题：我刚才错误地假设了什么？哪个测试暴露了它？如果换成真实网络请求或持久化写入，还缺少什么？
