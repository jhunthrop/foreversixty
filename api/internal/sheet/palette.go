package sheet

import (
	"image/color"
	"strings"
)

// The site's tokens (web/src/styles/tokens.css) the sheet and the tiles draw with. No colour
// here is new: each names the token it copies.
var (
	colorSurface   = rgb(0x0d0f15) // --surface-card's base
	colorLineWarm  = rgb(0x3a3326) // --color-line-warm
	colorText      = rgb(0xe9e4d8) // --color-text
	colorStrong    = rgb(0xf2eee4) // --color-strong
	colorMuted     = rgb(0x9a9484) // --color-muted
	colorGold      = rgb(0xe5b955) // --color-gold
	colorAlliance  = rgb(0x6fb1ff) // --color-alliance
	colorHorde     = rgb(0xff6b5c) // --color-horde
	colorEmptyFill = rgb(0x0d111a) // the empty socket's fill (spec §B.3)
	colorEmptyRing = rgb(0x262e40) // the empty and unknown ring (spec §A.4, §B.3)
	colorBlack     = rgb(0x000000)
)

// qualityRing is --color-rarity-*, the game's quality colours, by quality number 0 to 5.
var qualityRing = [...]color.RGBA{
	rgb(0x9d9d9d), rgb(0xffffff), rgb(0x1eff00), rgb(0x0070dd), rgb(0xa335ee), rgb(0xff8000),
}

// qualityText is the colour a name is drawn in: the ring colour, except rare and epic, which use
// --color-rarity-rare-text and --color-rarity-epic-text for contrast on the dark surface.
var qualityText = [...]color.RGBA{
	rgb(0x9d9d9d), rgb(0xffffff), rgb(0x1eff00), rgb(0x3d94f0), rgb(0xb866f5), rgb(0xff8000),
}

// qualityWords are the quality names, by quality number; the emoji names use them lowercased.
var qualityWords = [...]string{"Poor", "Common", "Uncommon", "Rare", "Epic", "Legendary"}

// classColors are --color-class-*, keyed by class slug.
var classColors = map[string]color.RGBA{
	"warrior": rgb(0xc69b6d), "paladin": rgb(0xf48cba), "hunter": rgb(0xaad372), "rogue": rgb(0xfff468),
	"priest": rgb(0xffffff), "shaman": rgb(0x3f8fe0), "mage": rgb(0x3fc7eb), "warlock": rgb(0x8788ee),
	"druid": rgb(0xff7c0a),
}

func rgb(v uint32) color.RGBA {
	return color.RGBA{R: uint8(v >> 16), G: uint8(v >> 8), B: uint8(v), A: 0xff}
}

// knownQuality reports whether q is one of the game's six qualities.
func knownQuality(q int) bool { return q >= 0 && q < len(qualityRing) }

// QualityColor is the game's colour for quality q (--color-rarity-*), and false for an unknown
// quality.
func QualityColor(q int) (color.RGBA, bool) {
	if !knownQuality(q) {
		return color.RGBA{}, false
	}
	return qualityRing[q], true
}

// QualityWord is "Epic" for 4, and "" for an unknown quality.
func QualityWord(q int) string {
	if !knownQuality(q) {
		return ""
	}
	return qualityWords[q]
}

// QualityEmojiName is the quality dot's application-emoji name (spec §A.4.1): fs_q_epic for 4,
// and "" for an unknown quality, which draws no dot.
func QualityEmojiName(q int) string {
	if !knownQuality(q) {
		return ""
	}
	return qualityEmojiPrefix + strings.ToLower(qualityWords[q])
}

const qualityEmojiPrefix = "fs_q_"

// ClassColor is --color-class-{slug}, and false for a slug that is not a class.
func ClassColor(slug string) (color.RGBA, bool) {
	c, ok := classColors[slug]
	return c, ok
}
