#!/usr/bin/env bash
if [ "$#" -lt 3 ] || [ -z "$1" ] || [ "$2" != -- ]; then
    printf 'usage: e14.sh LOCK_DIR -- COMMAND [ARG...]\n' >&2; exit 2
fi
lock=$1; shift 2
# mkdir 的一次创建充当互斥操作，不先 test 再创建。
if ! mkdir -- "$lock" 2>/dev/null; then
    printf 'cannot acquire lock\n' >&2; exit 75
fi
child=''
cleanup() { rmdir -- "$lock" || printf 'lock directory was changed\n' >&2; }
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
"$@" <&0 &
child=$!
if wait "$child"; then rc=0; else rc=$?; fi
child=''
exit "$rc"
