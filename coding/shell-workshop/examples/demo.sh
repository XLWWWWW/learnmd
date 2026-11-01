#!/usr/bin/env bash
root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd) || exit 1
bash_bin=${BASH_BIN:-/bin/bash}
case ${1:-basic} in
    basic)
        "$bash_bin" "$root/solutions/e01.sh" 'hello shell' '' '*.log'
        ;;
    logs)
        printf 'service\trequests\t5xx\n'
        "$bash_bin" "$root/solutions/e06.sh" "$root/fixtures/access.log"
        ;;
    config)
        "$bash_bin" "$root/solutions/e09.sh" "$root/fixtures/app.conf"
        ;;
    backup)
        tmp=$(mktemp -d "${TMPDIR:-/tmp}/shell-demo.XXXXXX") || exit 1
        trap 'rm -rf -- "$tmp"' EXIT
        trap 'exit 130' INT
        trap 'exit 143' TERM
        mkdir "$tmp/source" "$tmp/restore" || exit 1
        printf 'hello backup\n' > "$tmp/source/hello world.txt" || exit 1
        "$bash_bin" "$root/solutions/e11.sh" "$tmp/source" "$tmp/backup.tar.gz" || exit 1
        tar -xzf "$tmp/backup.tar.gz" -C "$tmp/restore" || exit 1
        cmp "$tmp/source/hello world.txt" "$tmp/restore/hello world.txt" || exit 1
        printf 'backup + restore + compare: PASS\n'
        ;;
    *) printf 'unknown demo\n' >&2; exit 2 ;;
esac
