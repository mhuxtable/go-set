#!/usr/bin/env bash
set -euo pipefail

DIR="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTNAME="$(basename "${BASH_SOURCE[0]}")"

TEMPLATES=
for x in $DIR/*.tpl; do
	# Not tolerant to non-ASCII template filenames. Okay for now...
	tpl="$(<"$x")"
	tpl=${tpl//\`/"\` + \"\`\" + \`"}

	TEMPLATE=$(printf "\ttpl_%s = \`%s\`" "$(basename "${x%.*}")" "$tpl")
	TEMPLATES+="$TEMPLATE"$'\n\n'
done

cat <<EOF >$DIR/values.go
package templates

//go:generate ./$SCRIPTNAME

var (
$TEMPLATES
)
EOF

go fmt $DIR/

>&2 echo ">>> Generated $DIR/values.go"
