package sheet

import (
	"image/color"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

// The sheet's and the tiles' inks, every one a site token (brand.Color*, generated from
// design/tokens/tokens.json): the round-5 web values.
var (
	colorText      = opaque(brand.ColorBaseText)         // words on the band and the card
	colorMuted     = opaque(brand.ColorForeverMutedCard) // the card's muted ink
	colorGold      = opaque(brand.ColorBaseGold)         // a status the player acts on
	colorGain      = opaque(brand.ColorForeverGain)      // a modelled gain
	colorCream     = opaque(brand.ColorForeverCream)     // titles without a class
	colorFooter    = opaque(brand.ColorForeverCream2)    // the line on the ground under the card
	colorAlliance  = opaque(brand.ColorFactionAlliance)
	colorHorde     = opaque(brand.ColorFactionHorde)
	colorEmptyFill = opaque(brand.ColorBaseBg)   // the empty socket's fill
	colorEmptyRing = opaque(brand.ColorBaseLine) // the empty and unknown ring
	colorBlack     = color.RGBA{A: 0xff}
)

// qualityRing is the game's quality colours (color.rarity.*), by quality number 0 to 5.
var qualityRing = [...]color.RGBA{
	opaque(brand.ColorRarityPoor), opaque(brand.ColorRarityCommon), opaque(brand.ColorRarityUncommon),
	opaque(brand.ColorRarityRare), opaque(brand.ColorRarityEpic), opaque(brand.ColorRarityLegendary),
}

// qualityText is the colour a name is drawn in: the ring colour, except rare and epic, which use
// the rarity text tokens for contrast on the dark card.
var qualityText = [...]color.RGBA{
	qualityRing[0], qualityRing[1], qualityRing[2],
	opaque(brand.ColorRarityRareText), opaque(brand.ColorRarityEpicText), qualityRing[5],
}

// qualityWords are the quality names, by quality number; the emoji names use them lowercased.
var qualityWords = [...]string{"Poor", "Common", "Uncommon", "Rare", "Epic", "Legendary"}

// QualityCount is how many item qualities there are, 0 (Poor) to QualityCount-1 (Legendary).
const QualityCount = len(qualityWords)

// classColors are color.class.*, keyed by class slug.
var classColors = map[string]color.RGBA{
	"warrior": opaque(brand.ColorClassWarrior), "paladin": opaque(brand.ColorClassPaladin),
	"hunter": opaque(brand.ColorClassHunter), "rogue": opaque(brand.ColorClassRogue),
	"priest": opaque(brand.ColorClassPriest), "shaman": opaque(brand.ColorClassShaman),
	"mage": opaque(brand.ColorClassMage), "warlock": opaque(brand.ColorClassWarlock),
	"druid": opaque(brand.ColorClassDruid),
}

// opaque is an opaque token as color.RGBA (the same bytes as its color.NRGBA).
func opaque(c color.NRGBA) color.RGBA { return color.RGBA{R: c.R, G: c.G, B: c.B, A: 0xff} }

// knownQuality reports whether q is one of the game's six qualities.
func knownQuality(q int) bool { return q >= 0 && q < len(qualityRing) }

// QualityColor is the game's colour for quality q, and false for an unknown quality.
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

// ClassColor is color.class.{slug}, and false for a slug that is not a class.
func ClassColor(slug string) (color.RGBA, bool) {
	c, ok := classColors[slug]
	return c, ok
}
