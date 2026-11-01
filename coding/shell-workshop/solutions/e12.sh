#!/usr/bin/env bash
if [ "$#" -lt 4 ] || [[ ! $1 =~ ^[1-5]$ ]] || [[ ! $2 =~ ^[0-2]$ ]] || [ "$3" != -- ]; then
    printf 'usage: e12.sh ATTEMPTS(1..5) DELAY(0..2) -- COMMAND [ARG...]\n' >&2
    exit 2
fi
attempts=$1; delay=$2; shift 3
for ((i=1; i<=attempts; i++)); do
    if "$@"; then
        exit 0
    else
        rc=$?
    fi
    [ "$i" -eq "$attempts" ] && exit "$rc"
    printf 'retry: attempt %s failed (rc=%s)\n' "$i" "$rc" >&2
    sleep "$delay" || exit 1
done
