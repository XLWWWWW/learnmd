#!/usr/bin/env bash
[ "$#" -eq 1 ] && [ -n "$1" ] || { printf 'usage: e10.sh TARGET < INPUT\n' >&2; exit 2; }
target=$1
case $target in /*) ;; *) target=./$target ;; esac
[ ! -L "$target" ] && { [ ! -e "$target" ] || [ -f "$target" ]; } || {
    printf 'target must be a regular file or absent\n' >&2; exit 1;
}
parent=${target%/*}
umask 077
# 临时文件和目标位于同一目录，正常本地文件系统上的 rename 才是原子替换。
tmp=$(mktemp "$parent/.shell-write.XXXXXX") || exit 1
cleanup() { rm -f -- "$tmp"; }
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM
cat > "$tmp" || exit 1
mv -f -- "$tmp" "$target" || exit 1
