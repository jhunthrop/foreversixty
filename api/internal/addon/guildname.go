package addon

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// maxGuildNameRunes is the game's own guild name limit.
const maxGuildNameRunes = 24

// validateGuildName normalises a decoded guild name (Unicode NFC,
// trimmed) and rejects one that cannot be a real WoW guild name: not
// valid UTF-8 (a hand-crafted percent-encoded sequence can decode to
// bytes that are not - this also rules out surrogates, which are never
// valid UTF-8 on their own), empty after trimming, longer than the
// game's own 24-character limit, or containing a rune from Unicode
// category C (Cc control, Cf format, Co private-use, Cs surrogate, and
// Cn unassigned - Go's unicode.C already covers all five; its doc
// comment reads "the set of Unicode control, special, and unassigned
// code points"). Cf alone covers every exploit named in the second
// security review response (B): zero-width space (U+200B), zero-width
// joiner (U+200D), right-to-left override (U+202E), and the byte-order
// mark (U+FEFF) - none of them can appear in a real WoW guild name, and
// a bidi override in particular can make a guild's displayed name
// misleading about what it actually is.
func validateGuildName(name string) (string, bool) {
	if !utf8.ValidString(name) {
		return "", false
	}
	normalized := strings.TrimSpace(norm.NFC.String(name))
	if normalized == "" || utf8.RuneCountInString(normalized) > maxGuildNameRunes {
		return "", false
	}
	for _, r := range normalized {
		if unicode.Is(unicode.C, r) {
			return "", false
		}
	}
	return normalized, true
}
