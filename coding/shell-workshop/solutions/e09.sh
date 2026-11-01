#!/usr/bin/env bash
[ "$#" -eq 1 ] || { printf 'usage: e09.sh CONFIG\n' >&2; exit 2; }
[ -f "$1" ] && [ -r "$1" ] || { printf 'unreadable config\n' >&2; exit 1; }
export LC_ALL=C
invalid() { printf 'invalid config at line %s\n' "$line_no" >&2; exit 2; }
name_seen=0; port_seen=0; line_no=0
while IFS= read -r line || [ -n "$line" ]; do
    line_no=$((line_no + 1))
    [[ $line != *$'\r'* ]] || invalid
    case $line in ''|'#'*) continue ;; *=*) ;; *) invalid ;; esac
    key=${line%%=*}
    value=${line#*=}
    case $key in
        APP_NAME)
            [ "$name_seen" -eq 0 ] && [ -n "$value" ] || invalid
            app_name=$value; name_seen=1
            ;;
        PORT)
            [ "$port_seen" -eq 0 ] && [[ $value =~ ^[0-9]{1,5}$ ]] || invalid
            port=$((10#$value))
            [ "$port" -ge 1 ] && [ "$port" -le 65535 ] || invalid
            port_seen=1
            ;;
        *) invalid ;;
    esac
done < "$1"
[ "$name_seen" -eq 1 ] && [ "$port_seen" -eq 1 ] || invalid
printf 'name=%s\nport=%s\n' "$app_name" "$port"
