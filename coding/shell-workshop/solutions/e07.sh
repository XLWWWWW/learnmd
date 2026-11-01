#!/usr/bin/env bash
[ "$#" -eq 1 ] || { printf 'usage: e07.sh DIRECTORY\n' >&2; exit 2; }
root=$1
check_root=$root
while [[ $check_root == */ && $check_root != / ]]; do check_root=${check_root%/}; done
[ -d "$root" ] && [ ! -L "$check_root" ] || { printf 'expected a real directory\n' >&2; exit 1; }
case $root in /*) ;; *) root=./$root ;; esac
# NUL 是记录分隔符。不要把结果装进 Shell 变量。
find "$root" -type f -name '*.log' -print0
