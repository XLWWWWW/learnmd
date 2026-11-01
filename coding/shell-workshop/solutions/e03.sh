#!/usr/bin/env bash
export LC_ALL=C
if [ "$#" -ne 1 ] || [[ ! $1 =~ ^[0-9]{1,3}$ ]]; then
    printf 'expected one decimal integer in 0..255\n' >&2
    exit 2
fi
# 先限制字符和长度，再做算术，避免八进制、表达式注入和溢出。
n=$((10#$1))
if [ "$n" -gt 255 ]; then
    printf 'number exceeds 255\n' >&2
    exit 2
fi
printf '%s\n' "$((n * (n + 1) / 2))"
