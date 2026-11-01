#!/usr/bin/env bash
if [ "$#" -ne 0 ]; then
    printf 'usage: e02.sh (reads APP_ENV and APP_PORT)\n' >&2
    exit 2
fi
printf 'env=%s\nport=%s\n' "${APP_ENV:-dev}" "${APP_PORT:-8080}"
