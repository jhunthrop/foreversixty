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
// bytes that are not), empty after trimming, longer than the game's own
// 24-character limit, or containing a control character. Part of the
// 2026-09-21 security review response (spec §3.3's amendment) - a
// malformed or oversized name must never reach a database write
// unvalidated.
func validateGuildName(name string) (string, bool) {
	if !utf8.ValidString(name) {
		return "", false
	}
	normalized := strings.TrimSpace(norm.NFC.String(name))
	if normalized == "" || utf8.RuneCountInString(normalized) > maxGuildNameRunes {
		return "", false
	}
	for _, r := range normalized {
		if unicode.IsControl(r) {
			return "", false
		}
	}
	return normalized, true
}
