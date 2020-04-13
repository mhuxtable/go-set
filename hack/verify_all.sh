#!/usr/bin/env bash
set -euo pipefail

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTNAME="$(basename "${BASH_SOURCE[0]}")"

pushd "$DIR/.." >/dev/null
function reset() {
	popd >/dev/null
}
trap reset EXIT

DIFF="$(git status --porcelain --untracked-files=no)"
printf "%s\n" "$DIFF"
if [ ! -z "$DIFF" ]; then
	>&2 echo "Untracked changes found before updates. This script will produce meaningless results."
	exit 1
fi

"$DIR/update_all.sh"

DIFF="$(git status --porcelain --untracked-files=no)"
printf "%s\n" "$DIFF"
if [ ! -z "$DIFF" ]; then
	>&2 echo "Changes found in generated files. Run hack/update_all.sh to rectify."
	exit 2
fi
