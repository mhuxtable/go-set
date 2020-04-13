#!/usr/bin/env bash
set -euo pipefail

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GENSET_SRC="$DIR/../cmd/genset"

function usage() {
	>&2 printf "Usage: $0 output_directory\n"
	exit 1
}

if [ $# -ne 1 ]; then
	usage
fi

OUTPUT_DIRECTORY="$(realpath "$1")"

pushd "$GENSET_SRC" >/dev/null
function reset() {
	popd >/dev/null
}
trap reset EXIT

if [ ! -d $OUTPUT_DIRECTORY ]; then
	mkdir -p $OUTPUT_DIRECTORY
fi

go generate ./templates
go build -o "$OUTPUT_DIRECTORY/genset" .
