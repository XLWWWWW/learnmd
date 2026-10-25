# 02 · 工作中常用的 12 类场景

下面的代码是聚焦关键路径的片段。使用 `lab` 的片段对应可运行参考实现，演示命令见 README。每个并发结构都应先写出它的正常结束路径和提前退出路径。

## 1. 异步任务返回一次结果

适用：启动一次计算，同时做其他准备，稍后取结果。

```go
type Result struct {
    Value string
    Err   error
}

result := make(chan Result, 1)
go func() {
    value, err := load(ctx) // load 必须能响应 ctx 或在有限时间内返回
    result <- Result{Value: value, Err: err}
}()

select {
case r := <-result:
    use(r.Value, r.Err)
case <-ctx.Done():
    return ctx.Err()
}
```

容量 1 对应“最多发送一次”的已知上界：即使接收方先离开，发送也能完成。这里不依赖 close 判断完成，所以可以不关闭 result。

这不能终止卡在 `load` 里的任务。如果只是同步调用一次 I/O，通常直接调用支持 context 的 API 更简单；不要为了“看起来异步”给每个调用套 goroutine。

若成功发送后还要继续做清理，接收结果也不代表清理结束；需要另一个 done 或 WaitGroup。

## 2. 有界任务队列与背压

适用：批量导入、发通知、图片处理等，希望任务入口受处理能力约束。

```go
jobs := make(chan Job, 64)

// 策略 A：等待，直到有空位或请求超时。
select {
case jobs <- job:
    return nil
case <-ctx.Done():
    return ctx.Err()
}

// 策略 B：立即拒绝；适合入口允许重试或明确反馈繁忙的业务。
select {
case jobs <- job:
    return nil
default:
    return ErrBusy
}
```

这两个片段是假定调用期间 jobs 不会被关闭的替代策略。共享队列的关闭必须由入口生命周期协议协调，不能让某个请求处理函数自行关闭。

若每个 worker 只持有一个任务，workers 为 W、队列容量为 B，则队列与处理中大致最多容纳 B+W 项；这个界限没有覆盖调用者手里尚未提交的任务、其他内部队列和结果积压。

允许丢弃的遥测数据可以使用策略 B 并记录丢弃数。订单、支付等不能静默丢失的任务，应考虑持久化、重试和幂等；进程内 channel 不提供这些保障。

## 3. 一次完成信号与广播

适用：初始化完成后唤醒多个依赖者，或通知一组 goroutine 停止。

```go
ready := make(chan struct{})

// 唯一负责初始化的一方：
go func() {
    initialize()
    close(ready)
}()

// 可以有多个等待者：
select {
case <-ready:
    useInitializedState()
case <-ctx.Done():
    return ctx.Err()
}
```

`struct{}` 表达“没有载荷，只关心事件”。close 后，后来到达的等待者也能立即观察到信号。

这是一次性广播，不能重新打开 channel。重复发生的业务事件，例如配置每次更新都通知订阅者，需要每个订阅者的队列、版本机制或其他显式广播设计。多个消费者共享一个数据 channel 不会各收到一份。

## 4. 超时与取消：同时照顾等待方和工作方

适用：请求取消、服务超时、数据库或 HTTP 调用的截止时间。

```go
ctx, cancel := context.WithTimeout(parent, 800*time.Millisecond)
defer cancel()

value, ok, err := lab.Receive(ctx, results)
if err != nil {
    return err
}
if !ok {
    return ErrNoResult
}
consume(value)
```

只让接收方超时退出还不够。上游发送也应允许被同一个 context 唤醒：

```go
select {
case results <- value:
case <-ctx.Done():
    return
}
```

如果上游还在等待输入，接收输入也需要同样的取消处理。如果它正在执行网络 I/O，要把 ctx 传给实际 I/O API。`ctx.Done()` 是取消通知，`ctx.Err()` 区分主动取消与截止时间到达。[context 文档](https://pkg.go.dev/context)

如果值与超时同时就绪，需要预先约定“允许成功或超时任一结果”，或另外设计更强的截止时间判定。单纯把取消 case 写在最前面不会提高优先级。

## 5. 流水线、fan-out、fan-in 与提前停止

适用：读取记录 → 解析 → 过滤 → 转换 → 入库，阶段间通过流式数据衔接。

```text
                  ┌── Map(worker 1) ──┐
Generate ─────────┤                   ├── Merge ── Take ── 消费
                  └── Map(worker 2) ──┘
```

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

input := lab.Generate(ctx, 1, 2, 3, 4, 5)
square := func(n int) int { return n * n }
a := lab.Map(ctx, input, square)
b := lab.Map(ctx, input, square) // 同一输入被两个 worker 竞争消费：fan-out
merged := lab.Merge(ctx, a, b)   // 多个输出汇入一个输出：fan-in

for n := range lab.Take(ctx, merged, 2) {
    fmt.Println(n)
}
cancel() // 只取两个，必须让其余上游停止
```

只有一个 Map 时保留输入顺序；两个 Map 竞争消费并合并后，整体输出顺序不保证。Merge 的参考实现保证每条独立输入内部的相对顺序。

`Take` 只负责停止自己的输出，无法推断是否应该取消整个上下游网络。因此由组织流水线的调用者 cancel。不要让 Take 关闭它不拥有的输入。

每个阶段关闭自己的输出；多发送者输出由协调者关闭。下游提前停止时，各级阻塞的收、发都必须有退出路径。缓冲只有在“剩余发送次数有明确上界”时才能独立解决退出后的发送问题。[官方流水线与取消说明](https://go.dev/blog/pipelines)

## 6. 固定 worker pool：限制正在工作的数量

适用：并发调用外部接口、批量查库、处理文件。任务很多，但同时只允许 W 个回调执行。

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()

input := lab.Generate(ctx, 1, 2, 3, 4, 5)
out, err := lab.ParallelMap(ctx, input, 3,
    func(ctx context.Context, n int) (int, error) {
        return fetchAndCalculate(ctx, n)
    })
if err != nil {
    return err
}

for r := range out {
    if r.Err != nil {
        recordFailure(r.Input, r.Err)
        continue
    }
    save(r.Input, r.Value)
}
```

这里采用“单项失败继续处理”：每项产生一个 Result，出错项也有 Input，便于定位、重试和统计。输入正常关闭时，已经取出的任务继续处理，所有 worker 完成后输出才关闭。

若要任一任务失败就整体停止，应由任务组或上层协调者传播首次错误并取消其余工作；这与当前参考实现的契约不同。不要顺手改成遇错 return 而不通知生产者。

如果收集方看到某项失败后要提前 return，必须 cancel，否则其他 worker 可能堵在结果发送。

固定 worker pool 控制的是回调并发数。若回调内部又无限启动 goroutine，这个限制就失去业务意义。W 要结合下游容量、连接池大小和任务性质确定。

## 7. 用缓冲 channel 做信号量

适用：已有遍历逻辑，只想限制少量独立任务的并发，不需要完整的任务池。

```go
func RunLimited(ctx context.Context, jobs []func(context.Context), limit int) error {
    if limit <= 0 {
        return errors.New("limit must be positive")
    }
    sem := make(chan struct{}, limit)
    var wg sync.WaitGroup
    wait := func() error {
        wg.Wait() // 先等已启动任务结束，再读取最终取消状态
        return ctx.Err()
    }

    for _, job := range jobs {
        if err := ctx.Err(); err != nil {
            return wait()
        }
        select {
        case sem <- struct{}{}:
        case <-ctx.Done():
            return wait()
        }
        wg.Add(1)
        go func(run func(context.Context)) {
            defer wg.Done()
            defer func() { <-sem }()
            run(ctx)
        }(job)
    }
    return wait()
}
```

先获取名额，再启动 goroutine，才能限制为工作创建的 goroutine 数量。若在每个新 goroutine 内部获取名额，大量 goroutine 仍然会堆积，只限制了实际进入临界工作区的数量。

示例假定 job 非 nil 且会完成或配合取消；不收集任务级错误。正常结束后不用 close sem，名额的获取和释放才是协议。[内存模型中的信号量用法](https://go.dev/ref/mem#Channel_communication)

信号量限制“同时多少个”，不限制“每秒多少次”。三个名额可能在一秒内被复用几千次。

## 8. 多副本查询，取第一个成功

适用：只读请求可以同时访问若干副本，希望一个成功就返回。

```go
value, err := lab.FirstSuccess(ctx,
    func(ctx context.Context) (int, error) { return queryA(ctx) },
    func(ctx context.Context) (int, error) { return queryB(ctx) },
    func(ctx context.Context) (int, error) { return queryC(ctx) },
)
```

注意这里是“第一个成功”，不是“第一个结束”。最快返回的可能是错误；只有全部失败才应返回聚合错误。空任务列表返回 `ErrNoJobs`。

参考实现的每个任务最多发送一次，结果 channel 的容量就是任务数，因此调用者提前返回后，其他结果也不会永远堵在发送处。无需 close：接收者知道最多接收多少次，也可以提前结束。

函数返回会取消子 context，但不等待所有失败方清理完成。每个回调仍须配合取消。若业务要求返回时所有回调均已结束，要另外加入完成等待，并接受慢任务清理可能增加返回延迟的事实。

这种模式增加请求成本。带副作用的写操作通常不能直接竞速；必须先考虑幂等与重复提交。并发启动任务数也应有合理上限。

## 9. Timer、Ticker 与简单的节奏控制

适用：一次等待超时、周期刷新、心跳、简单的请求间隔控制。

```go
timer := time.NewTimer(500 * time.Millisecond)
defer timer.Stop()
select {
case value, ok := <-input:
    if !ok {
        return ErrInputClosed
    }
    consume(value)
case <-timer.C:
    return ErrTimeout
case <-ctx.Done():
    return ctx.Err()
}
```

一次等待也可以用 `time.After`。若循环每轮都创建它，会有重复分配，而且超时起点每次重置：

```go
// 这是“多久没有新事件”的空闲超时，不是整个任务的总超时。
for {
    select {
    case _, ok := <-events:
        if !ok {
            return
        }
    case <-time.After(time.Second):
        return
    }
}
```

要限制总耗时，在循环外创建 timer 或使用一次 WithTimeout。频繁循环可用可重置的 timer，但要按所支持 Go 版本的 Stop/Reset 规则处理。

```go
ticker := time.NewTicker(100 * time.Millisecond)
defer ticker.Stop()
for {
    select {
    case <-ctx.Done():
        return
    case _, ok := <-jobs:
        if !ok {
            return
        }
    }
    // 此处已经领取任务，但等到下一个 tick 才开始执行。
    select {
    case <-ctx.Done():
        return
    case <-ticker.C:
        callRemote(ctx)
    }
}
```

这是一个简单的顺序节奏控制示例：第一次也等待 tick；处理慢时 tick 可能被调整或丢弃。它不是精确实时调度，也不是完整的令牌桶。需要突发额度、严格业务配额时应使用相应算法或成熟实现。

`Ticker.Stop()`、`Timer.Stop()` 不会关闭它们的 C。不能靠 `for range ticker.C` 自动退出。

版本提醒：Go 1.23 起计时器的回收与 channel 行为有变化，旧的“time.After 一定持续泄漏直到触发”不能当作通用结论。新计时器语义还受主模块 go 版本与 `GODEBUG=asynctimerchan` 设置影响；本项目不依赖复杂 Reset 逻辑。[time 文档](https://pkg.go.dev/time)、[官方 Go 1.23 timer 说明](https://go.dev/wiki/Go123Timer)

## 10. 按数量或时间批量提交

适用：日志写入、批量入库、指标上报。满足任一条件就刷新：达到 N 条；周期 tick 到达；输入正常关闭。

```go
ticker := time.NewTicker(200 * time.Millisecond)
defer ticker.Stop()

batches, err := lab.Batch(ctx, input, 100, ticker.C)
if err != nil {
    return err
}
for values := range batches {
    if err := writeBatch(ctx, values); err != nil {
        cancel() // 假定该调用者持有这条流水线的 cancel
        return err
    }
}
```

本实现的时间窗口从 ticker 启动开始计算，**不是从第一项到达开始计时**。如果第一个值在下一 tick 前 1ms 到达，它可能很快被刷新。若需求是“第一项到达后最多等 200ms”，应在从空批变为非空时启动 timer，提交后停止或重置。

还要明确取消策略。本练习选择：正常关闭刷新尾批；取消允许丢弃未交付的批次。不能据此实现“已接收数据绝不能丢”的业务，后者需要持久化、完成确认或独立的排空阶段。

发送 `[]int` 后不能立刻用 `batch = batch[:0]` 复用底层数组；消费者可能还在读。参考实现交出该数组后分配新的批次。若要做内存复用，必须添加归还缓冲的协议。

E10 注入 tick channel，测试可主动发 tick，避免靠 Sleep 猜测时钟。ticks 关闭后设为 nil，禁用该分支。

## 11. 优雅退出：停止接单、排空、确认完成

适用：服务重启、部署、接收终止信号后完成已接收的工作。

推荐先画出责任链：

```text
停止接收新请求
    ↓
通知并等待所有生产者不再发送 jobs
    ↓
唯一协调者 close(jobs)
    ↓
workers 读完 jobs，完成结果发送后退出
    ↓
确认 workers 全部退出，close(results)
    ↓
收集者读完 results，确认落库或记录结果完成
```

结果收集者必须在 worker 工作期间并发消费。如果主 goroutine 先等 workers，再开始读无缓冲 results，worker 可能一直堵在发送，形成循环等待。

也不能拿同一个已取消的 context 一边取消所有 worker，一边要求它们继续把队列排空。常见设计是分开“停止接单”和“强制停止”，排空阶段再使用独立的截止时间。

当排空超时，明确未处理任务如何恢复。仅有内存队列无法保证进程退出后的任务不丢。这个场景的综合题见 P1。

## 12. 由一个 goroutine 独占状态，通过请求与响应访问

适用：状态机、连接管理、需要按顺序处理的协调逻辑。

```go
type Lookup struct {
    Key   string
    Reply chan LookupResult
}
type LookupResult struct {
    Value string
    Found bool
}

func LookupValue(ctx context.Context, requests chan<- Lookup,
    key string) (LookupResult, error) {
    reply := make(chan LookupResult, 1)
    request := Lookup{Key: key, Reply: reply}
    select {
    case requests <- request:
    case <-ctx.Done():
        return LookupResult{}, ctx.Err()
    }
    select {
    case r := <-reply:
        return r, nil
    case <-ctx.Done():
        return LookupResult{}, ctx.Err()
    }
}
```

状态持有者在自己的循环里访问 map：

```go
for {
    select {
    case <-serviceCtx.Done():
        return
    case r, ok := <-requests:
        if !ok {
            return
        }
        value, found := state[r.Key]
        r.Reply <- LookupResult{Value: value, Found: found}
    }
}
```

这里内部请求协议要求 Reply 非 nil、容量至少 1、只发送一次，且调用方不关闭它。因此请求方先超时后，状态持有者也不会堵在回复。requests 的关闭仍需协调所有提交者，服务结束时的请求语义也应定义。

这个模式把共享状态串行化，能自然表达处理顺序；代价是所有操作经过同一循环。若只是保护一个普通 map，Mutex 往往更清晰。状态循环中执行慢 I/O 还会阻塞所有请求，需要另做任务拆分。

## 把场景迁移到工作代码前，写下这五行

```text
数据：任务 / 结果 / 事件 / 信号？是否允许丢失？
所有权：谁创建、谁发送、谁接收、谁负责关闭？
容量：队列最多几项？同时最多几个任务？结果积压在哪里？
结束：正常完成怎样退出？下游提前返回怎样退出？
错误：单项继续、整体失败，还是只取一个成功？谁等待清理完成？
```

能把这五行写清楚，再选 channel、Mutex、WaitGroup 的组合，会比先写一堆 goroutine 再补退出逻辑更容易维护。
