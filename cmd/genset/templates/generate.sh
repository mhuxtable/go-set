#!/usr/bin/env bash
set -euo pipefail
set -x

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTNAME="$(basename "${BASH_SOURCE[0]}")"

cat <<EOF >$DIR/values.go
package templates

//go:generate ./$SCRIPTNAME

var (
EOF

TEMPLATES=

function emit_file() {
	local x="$1"

	# Not tolerant to non-ASCII template filenames. Okay for now...
	tpl="$(<"$x")"
	tpl=${tpl//\`/"\` + \"\`\" + \`"}

	cat <<EOF >>$DIR/values.go

	// source: ${x}
	tpl_$(basename "${x%.*}") = \`${tpl}\`

EOF
}

emit_file set.tpl
emit_file set_test.tpl

cat <<EOF >>$DIR/values.go
)
EOF

go fmt $DIR/
>&2 echo ">>> Generated $DIR/values.go"
