#!/usr/bin/env bash
set -euo pipefail

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTNAME="$(basename "${BASH_SOURCE[0]}")"

pushd "$DIR/.." >/dev/null
function reset() {
	popd >/dev/null
}
trap reset EXIT

cd cmd/genset
go test -v -covermode=count -coverprofile=coverage.out ./...
$HOME/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
