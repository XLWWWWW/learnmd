#!/usr/bin/env bash
if [ "$#" -lt 3 ] || [[ ! $1 =~ ^[1-4]$ ]] || [ "$3" != -- ]; then
    printf 'usage: e15.sh LIMIT(1..4) EXECUTABLE -- [ITEM...]\n' >&2; exit 2
fi
limit=$1; worker=$2; shift 3
case $worker in /*) ;; *) worker=./$worker ;; esac
[ -f "$worker" ] && [ -x "$worker" ] || { printf 'worker must be an executable file\n' >&2; exit 2; }
pids=(); failed=0
stop_children() {
    code=$1
    trap '' INT TERM
    for pid in "${pids[@]}"; do kill -TERM "$pid" 2>/dev/null || :; done
    for pid in "${pids[@]}"; do wait "$pid" 2>/dev/null || :; done
    exit "$code"
}
trap 'stop_children 130' INT
trap 'stop_children 143' TERM
wait_batch() {
    for pid in "${pids[@]}"; do
        if wait "$pid"; then :; else failed=1; fi
    done
    pids=()
}
# Bash 3.2 没有 wait -n。按批并发满足上限，但慢任务会阻止下一批启动。
for item in "$@"; do
    "$worker" "$item" </dev/null &
    pids+=("$!")
    if [ "${#pids[@]}" -eq "$limit" ]; then wait_batch; fi
done
wait_batch
exit "$failed"
