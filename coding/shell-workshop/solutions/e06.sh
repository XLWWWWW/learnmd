#!/usr/bin/env bash
[ "$#" -eq 1 ] || { printf 'usage: e06.sh LOG\n' >&2; exit 2; }
[ -f "$1" ] && [ -r "$1" ] || { printf 'unreadable log\n' >&2; exit 1; }
export LC_ALL=C
# awk 先完整校验、聚合，END 才输出，避免坏行产生半份报表。
awk '
    /^[[:space:]]*($|#)/ { next }
    NF != 4 || $2 !~ /^(INFO|WARN|ERROR)$/ ||
    $3 !~ /^[a-z][a-z0-9_-]*$/ || $4 !~ /^[1-5][0-9][0-9]$/ {
        bad=1; next
    }
    { count[$3]++; if ($4 + 0 >= 500) errors[$3]++ }
    END {
        if (bad) exit 2
        for (service in count)
            printf "%s\t%d\t%d\n", service, count[service], errors[service]+0
    }
' < "$1" | sort
statuses=("${PIPESTATUS[@]}")
if [ "${statuses[0]}" -ne 0 ] || [ "${statuses[1]}" -ne 0 ]; then
    printf 'log parsing or sorting failed\n' >&2
    exit 1
fi
