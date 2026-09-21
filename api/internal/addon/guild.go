package addon

import (
	"net/url"
	"strconv"
	"strings"
)

// ParseFS1Guild reads only the `guild=` section of an FS1 v2 export string
// (docs/superpowers/specs/2026-09-21-guild-membership-design.md §1.2), ignoring
// everything else. It exists because the API otherwise never parses an export's
// contents — this is the one field that must cross that line, and it crosses it
// minimally: it is not a port of Codec.decodeFS1/fs1.ts's decodeFS1 and validates
// nothing about any other section.
//
// ok is false both when the export carries no guild= section at all and when it
// carries one that cannot be read (a bad percent-encoding, a missing colon, or a
// non-digit rank). The client-side codec refuses a whole code over a malformed
// *known* section; this narrower server-side reader does not have that luxury —
// the export is otherwise opaque here — so a malformed guild section is simply
// read as "no guild," exactly like an absent one. PutExports acts on ok the same
// way either way: no guild_characters row for this character.
func ParseFS1Guild(export string) (name string, rankIndex int, ok bool) {
	for _, section := range strings.Split(export, "|") {
		payload, isGuild := strings.CutPrefix(section, "guild=")
		if !isGuild {
			continue
		}
		encodedName, rankStr, hasColon := strings.Cut(payload, ":")
		if !hasColon || encodedName == "" || !isDigits(rankStr) {
			return "", 0, false
		}
		decoded, err := url.PathUnescape(encodedName)
		if err != nil || decoded == "" {
			return "", 0, false
		}
		n, err := strconv.Atoi(rankStr)
		if err != nil {
			return "", 0, false
		}
		return decoded, n, true
	}
	return "", 0, false
}

// isDigits reports whether s is one or more ASCII digits — the same grammar
// fs1.ts's /^\d+$/ and Codec.lua's isDigits check for a section's numeric half.
func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
