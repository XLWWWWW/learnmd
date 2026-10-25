#!/bin/sh
set -eu

go_reflection_root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
if [ -n "${GO_REFLECTION_GO:-}" ]; then
    go_reflection_bin=$GO_REFLECTION_GO
elif command -v go >/dev/null 2>&1; then
    go_reflection_bin=$(command -v go)
elif [ -x /private/tmp/go-reflection-toolchain/go/bin/go ]; then
    go_reflection_bin=/private/tmp/go-reflection-toolchain/go/bin/go
else
    printf '%s\n' 'Go not found. Install a supported Go version >= 1.22, or set GO_REFLECTION_GO to its absolute path.' >&2
    exit 127
fi

cd "$go_reflection_root"
exec env \
    "GOCACHE=${GOCACHE:-${TMPDIR:-/tmp}/go-reflection-gocache}" \
    "GOPATH=${GOPATH:-${TMPDIR:-/tmp}/go-reflection-gopath}" \
    GOTOOLCHAIN=local \
    "$go_reflection_bin" "$@"
