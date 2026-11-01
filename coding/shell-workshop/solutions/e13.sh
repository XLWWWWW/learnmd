#!/usr/bin/env bash
if [ "$#" -lt 2 ] || [ "$1" != -- ]; then
    printf 'usage: e13.sh -- COMMAND [ARG...]\n' >&2; exit 2
fi
shift
umask 077
tmp=$(mktemp -d "${TMPDIR:-/tmp}/shell-workspace.XXXXXX") || exit 1
child=''
cleanup() {
    # 清理的对象只有 mktemp 刚创建的目录，不接收用户指定的删除目标。
    rm -rf -- "$tmp"
}
on_signal() {
    code=$1
    trap '' INT TERM
    if [ -n "$child" ]; then
        kill -TERM "$child" 2>/dev/null || :
        wait "$child" 2>/dev/null || :
    fi
    exit "$code"
}
trap cleanup EXIT
trap 'on_signal 130' INT
trap 'on_signal 143' TERM
# & 默认可能使子进程 stdin 指向 /dev/null，所以显式保留输入。
WORK_DIR=$tmp "$@" <&0 &
child=$!
if wait "$child"; then rc=0; else rc=$?; fi
child=''
exit "$rc"
