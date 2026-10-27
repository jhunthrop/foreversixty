package discordbot

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// ItemFacts is what an item embed shows. Name, Quality, ItemLevel and the icon come from the
// build's item table; Source from the BiS data, in the game's words. Icon is the absolute
// thumbnail URL ({PUBLIC}/data/{build}/icons/{icon}.webp), empty when the item is unknown to
// the build; Quality is -1 and ItemLevel 0 then too, and the line leaves them out.
type ItemFacts struct {
	ID        int
	Name      string
	Quality   int
	ItemLevel int
	Icon      string
	Source    string
	URL       string
}

// minusSign is U+2212, the SimC convention for a negative DPS figure.
const minusSign = "−"

// lineJoin joins the parts of an item embed's lines.
const lineJoin = " · "

// The game's own quality colours (--color-rarity-*, never the lightened -text variants),
// by quality number, as embed colour integers.
var qualityColors = []int{0x9d9d9d, 0xffffff, 0x1eff00, 0x0070dd, 0xa335ee, 0xff8000}

// classColors are --color-class-*, as embed colour integers.
var classColors = map[string]int{
	"warrior": 0xc69b6d, "paladin": 0xf48cba, "hunter": 0xaad372, "rogue": 0xfff468,
	"priest": 0xffffff, "shaman": 0x3f8fe0, "mage": 0x3fc7eb, "warlock": 0x8788ee, "druid": 0xff7c0a,
}

// ItemEmbed is the one item grammar of /bis, /char and /report (spec §3.3.3): the name
// linking to the site, the quality's colour, the icon, and two lines: the slot (when slot is
// not empty), quality word and item level over the delta (when there is one) and the source.
// author, when set, is the line above the title.
func ItemEmbed(it ItemFacts, slot string, delta string, author string) notify.Embed {
	e := notify.Embed{
		Title: memberText(it.Name), URL: it.URL, Thumbnail: it.Icon, Author: author,
		Color: qualityColor(it.Quality),
	}
	lines := []string{}
	if one := joinPresent(slot, qualityWord(it.Quality), itemLevelText(it.ItemLevel)); one != "" {
		lines = append(lines, one)
	}
	if two := joinPresent(delta, it.Source); two != "" {
		lines = append(lines, two)
	}
	e.Description = strings.Join(lines, "\n")
	return e
}

func qualityColor(q int) int {
	if q < 0 || q >= len(qualityColors) {
		return 0
	}
	return qualityColors[q]
}

func qualityWord(q int) string {
	if q < 0 || q >= len(qualityWords) {
		return ""
	}
	return qualityWords[q]
}

func itemLevelText(level int) string {
	if level <= 0 {
		return ""
	}
	parts := strings.Split(lookupCopy.ItemLineOne, lineJoin) // {Slot}, {Quality}, item level {n}
	return fill(parts[len(parts)-1], "{n}", fmt.Sprint(level))
}

func joinPresent(parts ...string) string {
	kept := parts[:0:0]
	for _, p := range parts {
		if p != "" {
			kept = append(kept, p)
		}
	}
	return strings.Join(kept, lineJoin)
}

// signedDPS is a DPS gain or loss, one decimal, with its sign: "+9.2 DPS", "−15.0 DPS".
func signedDPS(v float64) string { return signedFigure(v, "DPS") }

// signedPoints is a Forever Score gain: "+12.4 points".
func signedPoints(v float64) string { return signedFigure(v, pointsWord) }

func signedFigure(v float64, unit string) string {
	v = math.Round(v*10) / 10
	sign := "+"
	if v < 0 {
		sign, v = minusSign, -v
	}
	return fmt.Sprintf("%s%.1f %s", sign, v, unit)
}

// closeEmbeds puts footer and the time on the last embed alone (spec §3.3.0).
func closeEmbeds(embeds []notify.Embed, footer string, at time.Time) {
	if len(embeds) == 0 {
		return
	}
	last := &embeds[len(embeds)-1]
	last.Footer, last.Timestamp = footer, at
}

// slotLabel is the slot as the /bis slot choice names it ("finger1" is "Ring 1").
func slotLabel(key string) string {
	for _, c := range slotChoices {
		if c.Value == key {
			return c.Name
		}
	}
	return classLabel(strings.ReplaceAll(key, "_", " "))
}

// classLabel capitalises a slug's first letter ("mage" is "Mage").
func classLabel(slug string) string {
	if slug == "" {
		return ""
	}
	return strings.ToUpper(slug[:1]) + slug[1:]
}
