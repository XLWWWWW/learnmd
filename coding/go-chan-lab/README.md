# Go channel：从理解到工作中可靠使用

是的，`chan` 是 Go 并发编程的核心特性之一。它把**传递数据**和**协调执行**结合在一起。真正掌握它的标准，是你能说明：谁发送、谁接收、哪里可能阻塞、谁关闭、下游提前退出后谁负责停止上游。

这套资料面向已经会 Go 基础语法、刚开始系统学习并发的开发者。所有例子均为本项目编写；关键语义已对照 Go 官方资料核对。代码只使用标准库，模块最低版本为 Go 1.23。

## 从哪里开始

| 顺序 | 内容 | 你将得到什么 |
|---|---|---|
| 1 | [核心机制](docs/01-channel.md) | 创建、阻塞、缓冲、关闭、nil、select、内存可见性 |
| 2 | [工作场景与模式](docs/02-patterns.md) | 12 类场景、代码、选择理由和适用边界 |
| 3 | [练习题](docs/03-exercises.md) | 24 道推理题、10 道可运行编程题、3 个综合项目 |
| 4 | [分层提示](docs/04-hints.md) | 卡住时先看提示，再继续自己写 |
| 5 | [答案与解析](docs/05-answers.md) | 推理题答案、实现思路、取消与关闭的证明 |
| 6 | [工作检查表](docs/06-checklist.md) | 代码审查、排障和迁移到业务代码的注意点 |

## 先运行一个例子

当前电脑的 PATH 中没有 Go。本项目提供了便捷脚本：优先使用已安装的 Go，也可使用资料生成时放在临时目录的官方工具链。

```bash
cd /Users/xlwwwww/learnmd/coding/go-chan-lab
./practice.sh demo basic
./practice.sh E01       # 开始第一题，未实现时会提示 TODO
./practice.sh solution  # 对照：参考答案运行相同测试并检查数据竞争
```

正式安装 Go 后也可以直接使用标准命令：

```bash
cd /Users/xlwwwww/learnmd/coding/go-chan-lab
go version
go run -tags solution ./cmd/demo -scenario basic
```

输出：

```text
value=0 ok=true
value=7 ok=true
value=0 ok=false
```

这三个输出区分了“收到真实的零值”和“channel 关闭且已经读空”。其他演示：

```bash
go run -tags solution ./cmd/demo -scenario pipeline
go run -tags solution ./cmd/demo -scenario workers
go run -tags solution ./cmd/demo -scenario first
go run -tags solution ./cmd/demo -scenario batch
```

`pipeline` 打印 `1`、`4`；`workers` 返回五个任务的结果，其中任务 3 故意失败，结果顺序不保证；`first` 返回 `42 <nil>`；`batch` 展示满批与尾批刷新，定时触发也可能改变分组边界。

## 边写边测

待实现代码位于 `lab/exercises.go`。先不要阅读 `lab/solutions.go`。

```bash
# 只做第一题；其他题的 TODO 不影响这一题。
go test ./lab -run '^TestE01' -v -timeout 15s

# 做第二题时改成 TestE02，后续同理。
go test ./lab -run '^TestE02' -v -timeout 15s

# 做完全部题后，验证你自己的实现。
go test ./... -timeout 30s
go test -race ./... -timeout 60s

# 需要对照答案时：使用同一套测试运行参考实现。
go test -tags solution ./... -timeout 30s
go test -race -tags solution ./... -timeout 60s
```

默认构建选中练习版本，所以尚未完成时遇到 `panic: TODO E01` 等提示是预期行为。`-tags solution` 会改为选择参考实现，不会测试你写的版本。两者的函数签名与测试完全相同。请保留文件顶部的构建标签。

若尚未安装 Go，可按[官方安装说明](https://go.dev/doc/install)安装，然后运行上面的命令。资料生成时额外下载了一份官方工具链到临时目录做验证，没有修改系统安装；临时目录不适合作为长期开发环境。

## 建议学习节奏：10 次，每次 45～75 分钟

| 次数 | 阅读与练习 | 当次目标 |
|---|---|---|
| 1 | 核心机制 1～5；Q01～Q06；E01 | 准确预测收发与关闭行为 |
| 2 | 核心机制 6～9；Q07～Q12；E02 | 理解 nil、单向 channel、select |
| 3 | 核心机制 10～13；Q13～Q18；E03 | 分清通信安全与共享对象安全 |
| 4 | 场景 1～4；E04、E05 | 写出能取消的生产者与转换器 |
| 5 | 场景 5；E06、E07 | 掌握多路合并与提前停止 |
| 6 | 场景 6～7；E08 | 限制并发、协调关闭、表达单项错误 |
| 7 | 场景 8～9；E09 | 处理首个成功、失败汇总、超时 |
| 8 | 场景 10～11；E10 | 处理批量、时钟、切片所有权 |
| 9 | Q19～Q24；综合项目 P1 | 发现隐藏的泄漏和死锁 |
| 10 | P2/P3 任选；工作检查表 | 为真实服务写出可解释的退出协议 |

每题都按“先预测 → 写出阻塞点 → 实现 → 跑测试 → 解释退出过程 → 改一个条件再验证”的顺序完成。隔两天不看答案重写 E06、E08，是很有效的自测。

## 资料中的约定

- 文档标为“片段”的代码需要放入自己的函数或补全 import；可直接运行的入口统一在 `cmd/demo`。
- 除 E01 和 E02 的特定约定外，并发阶段使用 `context.Context` 取消。Context 必须非 nil；回调函数必须非 nil，且不能随意 panic。
- 取消是协作式的。它不强杀 goroutine，也不保证取消之后绝对不会再交付一个已经就绪的值。
- 输出顺序只在明确说明时保证；并发结果不要凭打印顺序判断正确性。
- 测试使用握手、计数和 channel 关闭验证行为，超时只用来防止测试挂死。通过 `-race` 不代表对所有执行路径的数学证明。

验证记录见 [VALIDATION.md](VALIDATION.md)。
