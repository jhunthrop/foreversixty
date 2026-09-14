// Package textx is the one place client-supplied text is cut down to
// the size a column will hold.
package textx

import (
	"strings"
	"unicode/utf8"
)

// Trim bounds a client-supplied string to at most max bytes, backing
// off from the cut point to the nearest rune boundary, and trims
// surrounding space.
//
// Byte-slicing UTF-8 blindly can split a multi-byte rune, and Postgres
// refuses the result with "invalid byte sequence for encoding UTF8" -
// a 500 on every attempt for anyone whose name or filename is both
// long and not ASCII. Every bound on text headed for a text column
// goes through here so there is only one rule and only one place it
// can be got wrong.
func Trim(s string, max int) string {
	s = strings.TrimSpace(s)
	if len(s) <= max {
		return s
	}
	s = s[:max]
	for !utf8.ValidString(s) {
		s = s[:len(s)-1]
	}
	return s
}
