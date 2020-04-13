package main

import (
	"go/token"
	"unicode"
)

// IsIdentifier determines whether ident is a valid Go identifier. This is more
// complicated than it needs to be, but token.IsIdentifier was only added in
// go1.13. This implementation is mostly stolen from the Go source code.
//
// IsIdentifier reports whether name is a Go identifier, that is, a non-empty
// string made up of letters, digits, and underscores, where the first character
// is not a digit. Keywords are not identifiers.
func IsIdentifier(name string) bool {
	for i, c := range name {
		if !unicode.IsLetter(c) && c != '_' && (i == 0 || !unicode.IsDigit(c)) {
			return false
		}
	}

	return name != "" && token.Lookup(name) == token.IDENT
}
