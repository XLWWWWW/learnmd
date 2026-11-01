#!/usr/bin/env bash
[ "$#" -eq 1 ] || { printf 'usage: e05.sh FILE\n' >&2; exit 2; }
[ -f "$1" ] && [ -r "$1" ] || { printf 'unreadable regular file\n' >&2; exit 1; }
i=0
while IFS= read -r line || [ -n "$line" ]; do
    i=$((i + 1))
    printf '%s:%s\n' "$i" "$line"
done < "$1"
