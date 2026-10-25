# Go 反射：理解、练习与工作中的应用

这套材料面向已经会写 Go 基础语法、struct、interface、指针和 slice，希望能读懂通用库并写出可靠反射代码的开发者。

反射是 Go 很重要的基础设施能力。日常业务开发更重要的是判断何时需要它：**需要在运行时理解未知类型的结构时，反射很有用；已经知道类型和行为时，优先使用直接代码、接口或泛型。**

## 从哪里开始

1. 阅读 [01_反射详解.md](01_反射详解.md)，建立原理与工作场景的联系。
2. 打开 [02_练习题册.md](02_练习题册.md)，先做 Q01—Q03 和 E01。
3. 在 `practice/exercises.go` 中实现函数，每完成一题就运行对应测试。
4. 卡住时查看 [03_提示与解析.md](03_提示与解析.md) 中折叠的提示，再看答案。
5. 完成基础练习后，做 E11 行映射、E12 性能实验和 P01—P03 项目设计。

共 **33 道主任务：12 道输出推演、6 道修复题、11 道有自动测试的实现题、1 道性能实验、3 道综合项目**。项目是进阶挑战，提供验收标准与参考设计，不附完整产品实现。

## 运行方式

需要 **Go 1.22 或以上版本**，没有第三方依赖。`go 1.22` 是本模块采用的语言版本下限；安装工具链时使用官方当前支持的稳定版本即可，见 [Go 安装说明](https://go.dev/doc/install)。

在本目录运行：

当前工作区已准备临时工具链，**可以把下面每条命令开头的 `go` 替换为 `sh run.sh` 直接运行**。例如：

```sh
sh run.sh test -v ./concepts
sh run.sh test ./practice -run '^TestE01' -v
```

脚本优先使用本机的 go，其次使用本次放在临时目录的工具链；临时目录被系统清理后，安装 Go 或通过 `GO_REFLECTION_GO` 指定工具链的绝对路径即可继续使用。脚本不会下载或安装软件。

```sh
# 先验证原理示例：这些不依赖你有没有完成习题。
go test -v ./concepts

# 开始练习：只测 E01，随后把编号换成 E02、E03……
go test ./practice -run '^TestE01' -v

# 验证你写好的全部实现。
go test ./...

# 运行参考答案；solution 标签会排除 exercises.go，使用 solutions.go。
go test -tags solution ./...
go test -tags solution -race ./...

# 运行实际场景演示。
go run -tags solution ./cmd/demo

# E12：只跑性能实验，输出耗时和分配。
go test -tags solution ./practice -run '^$' -bench E12 -benchmem -count=3
```

**初始状态下普通 `go test ./...` 会在 TODO 处失败，这是留给你实现的部分。** `go test -tags solution ./...` 才是完整参考答案的验证方式。不要带 `solution` 标签来判断自己的实现是否正确。

## 文件说明

| 文件 | 用途 |
|---|---|
| `01_反射详解.md` | 核心概念、使用场景、工作边界与 API 速查 |
| `02_练习题册.md` | 题目、明确输入输出约定、验收条件 |
| `03_提示与解析.md` | 分层提示、输出解析、实现思路、项目参考设计 |
| `practice/exercises.go` | 你要编辑的 11 个函数或方法 |
| `practice/exercises_test.go` | 正常输入与边界测试；也是可执行的需求说明 |
| `practice/solutions.go` | 完整参考实现，只在 `solution` 标签下参与编译 |
| `practice/benchmark_test.go` | E12 性能实验 |
| `concepts/examples_test.go` | 12 个带预期输出的可执行原理示例 |
| `cmd/demo/main.go` | 默认值、校验、map、行映射和动态调用的组合演示 |
| `run.sh` | 选择可用工具链并在本模块中执行 Go 命令 |
| `04_验证记录.md` | 本次实际执行的验证结果与限制 |

这些组件是刻意限定范围的教学实现。练习中的 `map`、`default`、`validate`、`db` 标签是本套材料约定的协议，Go 不会自动解释这些标签。尤其要先读清 `omitempty`、`required`、类型转换与失败时是否写入的约定。
