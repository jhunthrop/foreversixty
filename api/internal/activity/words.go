package activity

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The level bands the BiS lists and the boards are cut at.
const (
	firstBand = 20
	topBand   = 60
	bandWidth = 10
)

// BandOf is the band a level reads against: 20 for anything below 30 (no band below 20 is
// ranked), then each ten, 60 at the cap.
func BandOf(level int) int {
	if level < firstBand {
		return firstBand
	}
	return min(topBand, level/bandWidth*bandWidth)
}

// ValidBand reports whether b is one of the bands.
func ValidBand(b int) bool {
	return b >= firstBand && b <= topBand && b%bandWidth == 0
}

// BandWords is how a sentence names a band: "20 to 29", "60".
func BandWords(band int) string {
	if band >= topBand {
		return "60"
	}
	return fmt.Sprintf("%d to %d", band, band+bandWidth-1)
}

// LevelWords is how a sentence places a character's sim: its band ("20 to 29", "60"), or for a
// character below the first band, which plays against that band's list without being in it,
// its own level ("level 13").
func LevelWords(level, band int) string {
	if level > 0 && level < firstBand {
		return fmt.Sprintf("level %d", level)
	}
	return BandWords(band)
}

// SpecName is the spec's player-facing name, the catalogue's full name: "Fury Warrior",
// "Feral Druid (Bear)"; "" for an unknown key.
func SpecName(key string) string {
	return specs.ByKey[key].FullName
}

// SpecRole is the spec's role ("dps", "tank", "healer"); "dps" for an unknown key.
func SpecRole(key string) string {
	if s, ok := specs.ByKey[key]; ok && s.Role != "" {
		return s.Role
	}
	return "dps"
}

// SpecKey turns a log's class and spec ("Hunter", "BeastMastery" or "Beast Mastery") into
// the site's key ("hunter-beast-mastery"); ok is false when no spec matches.
func SpecKey(class, spec string) (string, bool) {
	key := strings.ToLower(class) + "-" + slugWords(spec)
	_, ok := specs.ByKey[key]
	return key, ok
}

func slugWords(s string) string {
	var b strings.Builder
	for i, r := range strings.TrimSpace(s) {
		switch {
		case r == ' ' || r == '_' || r == '-':
			b.WriteByte('-')
		case unicode.IsUpper(r) && i > 0 && !strings.HasSuffix(b.String(), "-"):
			b.WriteByte('-')
			b.WriteRune(unicode.ToLower(r))
		default:
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func titleWord(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
