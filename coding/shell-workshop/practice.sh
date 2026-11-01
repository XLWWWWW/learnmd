#!/usr/bin/env bash
# 所有路径相对于脚本本身，支持从任意工作目录启动。
project_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd) || exit 1
bash_bin=${BASH_BIN:-/bin/bash}
if [ ! -x "$bash_bin" ]; then printf 'BASH_BIN must name an executable path\n' >&2; exit 2; fi
command_name=${1:-help}
if [ "$#" -gt 0 ]; then shift; fi
case $command_name in
    E0[1-9]|E1[0-6]|all|solution)
        command -v python3 >/dev/null 2>&1 || { printf 'tests require Python 3.8+\n' >&2; exit 1; }
        mode=practice; exercise=$command_name
        if [ "$command_name" = solution ]; then mode=solution; exercise=${1:-all}; fi
        exec python3 "$project_dir/tests/test_workshop.py" --mode "$mode" --exercise "$exercise" --bash "$bash_bin"
        ;;
    demo)
        exec "$bash_bin" "$project_dir/examples/demo.sh" "${1:-basic}"
        ;;
    check)
        failed=0
        while IFS= read -r -d '' file; do
            "$bash_bin" -n "$file" || failed=1
        done < <(find "$project_dir" -path "$project_dir/scratch" -prune -o -type f -name '*.sh' -print0)
        [ "$failed" -eq 0 ] || exit 1
        printf 'Bash syntax: PASS\n'
        if command -v shellcheck >/dev/null 2>&1; then
            find "$project_dir" -path "$project_dir/scratch" -prune -o -type f -name '*.sh' -exec shellcheck --shell=bash {} +
        else
            printf 'ShellCheck: SKIP (not installed; no automatic installation)\n'
        fi
        ;;
    doctor)
        "$bash_bin" --version | head -n 1
        if command -v python3 >/dev/null 2>&1; then python3 --version; else printf 'Python 3 missing (only tests need it)\n'; fi
        for tool in awk sort grep find mktemp tar ln; do command -v "$tool" || exit 1; done
        if command -v shellcheck >/dev/null 2>&1; then shellcheck --version; else printf 'ShellCheck: optional, unavailable\n'; fi
        ;;
    help|-h|--help)
        printf '%s\n' \
            './practice.sh E01            测试你的第 1 题（E01..E16）' \
            './practice.sh all            测试全部练习；初始 TODO 会失败' \
            './practice.sh solution E01   用同一套测试验证第 1 题参考答案' \
            './practice.sh solution       验证全部参考答案' \
            './practice.sh demo basic     运行演示：basic|logs|config|backup' \
            './practice.sh check          语法检查；可用时运行 ShellCheck' \
            './practice.sh doctor         查看本地依赖（不下载安装）'
        ;;
    *) printf 'unknown command: %s; use help\n' "$command_name" >&2; exit 2 ;;
esac
