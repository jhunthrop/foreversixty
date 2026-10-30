package discordbot

import (
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
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

// classColor is --color-class-{slug} as an embed colour integer; package sheet holds the one
// copy of the class colours. false for a slug that is not a class.
func classColor(slug string) (int, bool) {
	c, ok := sheet.ClassColor(slug)
	return int(c.R)<<16 | int(c.G)<<8 | int(c.B), ok
}

// classColorOf is the class's colour, 0 (none) for a slug that is not a class.
func classColorOf(slug string) int {
	c, _ := classColor(slug)
	return c
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

// itemLinkParam is the query value that keeps each item embed's link its own.
const itemLinkParam = "item"

// itemLink is page with key in its query: Discord folds consecutive embeds that share a URL
// into one card, so items linking the same page showed as one. key must be unique among the
// message's embeds.
func itemLink(page, key string) string {
	u, err := url.Parse(page)
	if err != nil {
		return page
	}
	q := u.Query()
	q.Set(itemLinkParam, key)
	u.RawQuery = q.Encode()
	return u.String()
}
