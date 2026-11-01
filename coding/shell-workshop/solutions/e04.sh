#!/usr/bin/env bash
usage() { printf 'usage: e04.sh --name NAME [--times 1..5] [--]\n' >&2; exit 2; }
name_seen=0
times_seen=0
name=''
times=1
while [ "$#" -gt 0 ]; do
    case $1 in
        --name)
            [ "$#" -ge 2 ] && [ "$name_seen" -eq 0 ] && [ -n "$2" ] || usage
            name=$2; name_seen=1; shift 2
            ;;
        --times)
            [ "$#" -ge 2 ] && [ "$times_seen" -eq 0 ] || usage
            case $2 in [1-5]) times=$2 ;; *) usage ;; esac
            times_seen=1; shift 2
            ;;
        --) shift; [ "$#" -eq 0 ] || usage; break ;;
        *) usage ;;
    esac
done
[ "$name_seen" -eq 1 ] || usage
for ((i=0; i<times; i++)); do
    printf 'hello=%s\n' "$name"
done
