package discordbot

import (
	"fmt"
	"image/color"
	"net/url"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
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

// lineJoin joins the parts of an item embed's lines.
const lineJoin = " · "

// The stripes besides class and quality: the teal accent on a normal answer (round 5,
// color.forever.accent) and the kill green on a done raid and its report.
var (
	accentColor = brand.RGB(brand.ColorForeverAccent)
	doneGreen   = brand.RGB(brand.ColorOutcomeKill)
)

// classColor is --color-class-{slug} as an embed colour integer; package sheet holds the one
// copy of the class colours. false for a slug that is not a class.
func classColor(slug string) (int, bool) {
	c, ok := sheet.ClassColor(slug)
	return brand.RGB(color.NRGBA(c)), ok
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
	if one := joinPresent(slot, sheet.QualityWord(it.Quality), itemLevelText(it.ItemLevel)); one != "" {
		lines = append(lines, one)
	}
	if two := joinPresent(delta, it.Source); two != "" {
		lines = append(lines, two)
	}
	e.Description = strings.Join(lines, "\n")
	return e
}

// qualityColor is the game's own quality colour (color.rarity.*, never the lightened text
// variants) as an embed colour integer, 0 (none) for an unknown quality.
func qualityColor(q int) int {
	c, ok := sheet.QualityColor(q)
	if !ok {
		return 0
	}
	return brand.RGB(color.NRGBA(c))
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
func signedDPS(v float64) string { return upgrades.SignedFigure(v, "DPS") }

// signedPoints is a Forever Score gain: "+12.4 points".
func signedPoints(v float64) string { return upgrades.SignedFigure(v, pointsWord) }

// closeEmbeds puts footer, with the seal before it, and the time on the last embed alone
// (spec §3.3.0).
func (s *Service) closeEmbeds(embeds []notify.Embed, footer string, at time.Time) {
	if len(embeds) == 0 {
		return
	}
	last := &embeds[len(embeds)-1]
	last.Footer, last.Timestamp = footer, at
	*last = notify.WithSealFooter(*last, s.sealURL())
}

// sealURL is the seal's public URL, the icon every footer carries.
func (s *Service) sealURL() string { return notify.AvatarURL(s.d.PublicBaseURL) }

// sealFooters is m with the seal before every embed footer.
func sealFooters(m Message, sealURL string) Message {
	embeds := make([]notify.Embed, len(m.Embeds))
	for i, e := range m.Embeds {
		embeds[i] = notify.WithSealFooter(e, sealURL)
	}
	m.Embeds = embeds
	return m
}

// slotLabel is the slot as the /bis slot choice names it ("finger1" is "Finger 1").
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
