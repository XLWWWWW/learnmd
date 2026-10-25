#!/bin/sh
set -eu

project_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$project_dir"

if command -v go >/dev/null 2>&1; then
    go_cmd=$(command -v go)
elif [ -x /private/tmp/go-chan-toolchain/go/bin/go ]; then
    go_cmd=/private/tmp/go-chan-toolchain/go/bin/go
    export GOCACHE="${TMPDIR:-/tmp}/go-chan-lab-build-cache"
    export GOPATH="${TMPDIR:-/tmp}/go-chan-lab-gopath"
else
    echo '请先安装 Go 1.23 或更高版本：https://go.dev/doc/install' >&2
    exit 1
fi
export GOTOOLCHAIN=local

case "${1:-help}" in
    E0[1-9]|E10)
        exec "$go_cmd" test ./lab -run "^Test$1" -v -timeout 15s
        ;;
    all)
        exec "$go_cmd" test -race ./... -timeout 60s
        ;;
    solution)
        exec "$go_cmd" test -race -tags solution ./... -timeout 60s
        ;;
    demo)
        exec "$go_cmd" run -tags solution ./cmd/demo -scenario "${2:-basic}"
        ;;
    *)
        echo '用法：./practice.sh E01        测试自己的第一题（其他题改编号）'
        echo '      ./practice.sh all        测试自己的全部实现并检查数据竞争'
        echo '      ./practice.sh solution   用相同测试验证参考答案'
        echo '      ./practice.sh demo basic 运行演示：basic|pipeline|workers|first|batch'
        ;;
esac
