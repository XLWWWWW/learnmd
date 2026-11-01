#!/usr/bin/env bash
[ "$#" -eq 2 ] || { printf 'usage: e16.sh LITERAL FILE\n' >&2; exit 2; }
[ -f "$2" ] && [ -r "$2" ] || { printf 'unreadable file\n' >&2; exit 1; }
export LC_ALL=C
# -F 按字面匹配；-e 允许模式以 - 开头；重定向不把路径解释为选项。
grep -F -e "$1" < "$2" | sort -u
statuses=("${PIPESTATUS[@]}")
# grep 1 表示没有匹配，是本题合法的空结果；2 及以上才是错误。
if [ "${statuses[0]}" -gt 1 ] || [ "${statuses[1]}" -ne 0 ]; then
    printf 'filter pipeline failed\n' >&2; exit 1
fi
exit 0
