#!/usr/bin/env bash
set -euo pipefail

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTNAME="$(basename "${BASH_SOURCE[0]}")"

pushd "$DIR/.." >/dev/null
function reset() {
	popd >/dev/null
}
trap reset EXIT

# Update the compat tests
"$DIR/update_compat_test_helpers.sh"

# Build genset for use later
BIN_DIR="$(realpath ".bin")"
mkdir -p .bin
function cleanup() {
	rm -r .bin
}
trap cleanup EXIT

$DIR/build_genset.sh "$BIN_DIR"

# Update the locally generated sets
PATH="$BIN_DIR:$PATH" go generate ./genericset/
