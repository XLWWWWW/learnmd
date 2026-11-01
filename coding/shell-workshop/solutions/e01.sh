#!/usr/bin/env bash
# E01：每个参数保持为一个参数，包括空串、通配符和空格。
printf 'count=%s\n' "$#"
i=1
for arg in "$@"; do
    printf '%s=<%s>\n' "$i" "$arg"
    i=$((i + 1))
done
