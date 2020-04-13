package main

import "go/token"

// IsIdentifier determines whether ident is a valid Go identifier. This is more
// complicated than it needs to be, but token.IsIdentifier was only added in
// go1.13.
func IsIdentifier(ident string) bool {
	return token.Lookup(ident) == token.IDENT
}
