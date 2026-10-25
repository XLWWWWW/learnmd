# 验证记录

日期：2026-10-25（Asia/Shanghai）。

## 运行环境

- macOS / arm64。
- 官方 Go 工具链：`go1.27.3 darwin/arm64`。
- 工具链仅解压于 `/private/tmp/go-chan-toolchain/go`，未写入系统安装位置。
- 下载文件 SHA-256：`dd3908cf20959d798c1e06ebe5b72213f0505ce01373b40ce3110975a6cd99c2`，与[官方发布元数据](https://go.dev/dl/?mode=json)一致。
- 模块声明最低 Go 1.23；本次实际运行使用上面的 Go 1.27.3，没有额外在 Go 1.23 上执行兼容性测试。

## 已完成的检查

| 检查 | 结果 |
|---|---|
| 参考实现 `go test -tags solution ./... -timeout 30s` | 通过 |
| 参考实现 `go test -race -tags solution ./... -timeout 60s` | 通过 |
| 补充输入复制测试后 `./practice.sh solution` | 通过，包含数据竞争检测 |
| 默认练习版本仅编译 `go test ./... -run '^$'` | 通过 |
| 参考实现 `go vet -tags solution ./...` | 通过 |
| `gofmt` | 已格式化全部 Go 文件 |
| 五个演示 basic / pipeline / workers / first / batch | 均正常退出，输出符合约定 |
| 文档内部文件链接与代码围栏 | 通过 |
| 题目数量 | 24 道推理题、10 道编程题、3 道综合设计题 |
| 编程测试 | 19 个测试函数，另含 E03 子测试 |

## 边界

默认练习实现保留 TODO，运行它的测试会失败，这是练习入口的预期行为。`-tags solution` 只验证参考实现，不会验证学习者的改动。

综合设计题 P1～P3 提供题目、验收条件与设计方向，没有预置完整实现。文档中的模式片段有些依赖业务函数，用于讲解；可直接运行的完整入口是 `cmd/demo`。

测试和 race detector 不证明任意调用方、任意回调、所有调度下都绝无泄漏或业务错误。必须同时遵守函数约定，并审查关闭所有权及各个阻塞点。
