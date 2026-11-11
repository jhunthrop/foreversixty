package notify

import (
	"strings"
)

// markdownSpecials are escaped wherever they appear in member text.
const markdownSpecials = "\\*_~`|>#[]"

// lineStartSpecials are escaped only at the start of a line.
const lineStartSpecials = "-"

// escapeMarkdown makes member text render literally in Discord.
func escapeMarkdown(s string) string {
	var b strings.Builder
	lineStart := true
	for _, r := range s {
		if strings.ContainsRune(markdownSpecials, r) || (lineStart && strings.ContainsRune(lineStartSpecials, r)) {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
		lineStart = r == '\n'
	}
	return b.String()
}

// collapse folds every run of whitespace into one space, so member text is
// one line.
func collapse(s string) string { return strings.Join(strings.Fields(s), " ") }

// escapedWithin cuts raw to rawMax runes, escapes it, and shortens it until
// the escaped text fits outMax runes, so no escape is ever cut in half.
func escapedWithin(raw string, rawMax, outMax int) string {
	runes := []rune(raw)
	if len(runes) > rawMax {
		runes = runes[:rawMax]
	}
	for {
		escaped := escapeMarkdown(string(runes))
		if len([]rune(escaped)) <= outMax || len(runes) == 0 {
			return escaped
		}
		runes = runes[:len(runes)-1]
	}
}

// EscapeWithin is escapedWithin for other packages that render member text
// into Discord messages.
func EscapeWithin(raw string, rawMax, outMax int) string { return escapedWithin(raw, rawMax, outMax) }

// Collapse is collapse for other packages: member text as one line.
func Collapse(s string) string { return collapse(s) }
