#!/usr/bin/env bash
if [ "$#" -lt 3 ] || [ "$2" != -- ]; then
    printf 'usage: e08.sh --dry-run|--run -- COMMAND [ARG...]\n' >&2
    exit 2
fi
mode=$1
shift 2
case $mode in
    --dry-run)
        # %q 是给人检查参数边界的 Bash 表示，不再交给 eval。
        for arg in "$@"; do printf '%q\n' "$arg"; done
        ;;
    --run) exec "$@" ;;
    *) printf 'unknown mode\n' >&2; exit 2 ;;
esac
