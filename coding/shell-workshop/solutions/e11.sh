#!/usr/bin/env bash
[ "$#" -eq 2 ] || { printf 'usage: e11.sh SOURCE_DIR NEW_ARCHIVE\n' >&2; exit 2; }
# 本题只接收绝对路径；备份文件必须在源目录之外，且不能覆盖已有文件。
case $1 in /*) ;; *) printf 'source must be absolute\n' >&2; exit 2 ;; esac
case $2 in /*) ;; *) printf 'archive must be absolute\n' >&2; exit 2 ;; esac
source_input=$1
while [[ $source_input == */ && $source_input != / ]]; do source_input=${source_input%/}; done
[ -d "$source_input" ] && [ ! -L "$source_input" ] || { printf 'invalid source directory\n' >&2; exit 1; }
source_dir=$(CDPATH= cd -- "$source_input" && pwd -P) || exit 1
archive=$2
archive_parent=$(CDPATH= cd -- "${archive%/*}" && pwd -P) || exit 1
archive=$archive_parent/${archive##*/}
case $source_dir in /) printf 'cannot archive root\n' >&2; exit 2 ;; esac
case $archive in "$source_dir"/*) printf 'archive must be outside source\n' >&2; exit 2 ;; esac
[ ! -e "$archive" ] && [ ! -L "$archive" ] || { printf 'archive already exists\n' >&2; exit 1; }
umask 077
tmp=$(mktemp "$archive_parent/.shell-backup.XXXXXX") || exit 1
cleanup() { rm -f -- "$tmp"; }
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM
tar -czf "$tmp" -C "$source_dir" . || exit 1
tar -tzf "$tmp" >/dev/null || exit 1
# 硬链接发布：目标已存在则失败，避免检查后被另一进程抢先创建的覆盖竞态。
# 临时文件与归档在同一目录；本题约定父目录可信，目标不会被替换成目录。
ln -- "$tmp" "$archive" || exit 1
