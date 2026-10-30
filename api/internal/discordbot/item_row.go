package discordbot

import (
	"context"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/sheet"
)

// itemRow is one item of the gear panel (spec 2026-10-29 §A.3): the item, its slot word, the
// gain when there is one, and on /char what it replaces.
type itemRow struct {
	item ItemFacts
	// slot is the slot key ("finger1"); its words are slotLabel's.
	slot string
	// gain is line 2's figure, already worded ("+7.4 DPS", panelCopy.NotModelled), or "".
	gain string
	// replaces is /char's line 4 (replacesLine), or "".
	replaces string
	// noSource drops line 3's source (the text guard, §A.6).
	noSource bool
}

// itemRowText is a row's Text Display: the dot and the bold name, the slot and the gain, the
// item's facts as subtext, and on /char what it replaces.
func (s *Service) itemRowText(ctx context.Context, r itemRow) string {
	lines := []string{s.qualityDot(ctx, r.item.Quality) + "**" + memberText(r.item.Name) + "**",
		joinPresent(slotLabel(r.slot), r.gain)}
	source := r.item.Source
	if r.noSource {
		source = ""
	}
	if facts := joinPresent(qualityWord(r.item.Quality), itemLevelText(r.item.ItemLevel), source); facts != "" {
		lines = append(lines, subtext(facts))
	}
	if r.replaces != "" {
		lines = append(lines, r.replaces)
	}
	return strings.Join(lines, "\n")
}

// itemSection is a row as a Section with the item's tile as its thumbnail: 3 components.
func (s *Service) itemSection(ctx context.Context, build string, r itemRow) Node {
	tile := thumbnailNode(sheet.TileURL(s.d.APIBaseURL, build, r.item.ID), tileAlt(r.item, slotLabel(r.slot)))
	return sectionNode(s.itemRowText(ctx, r), tile)
}

// tileAlt is "{Item name}, {Quality} {slot} item", without the quality when it is unknown.
func tileAlt(it ItemFacts, slot string) string {
	slot = strings.ToLower(slot)
	if q := qualityWord(it.Quality); q != "" {
		return fill(panelCopy.TileAlt, "{Item name}", it.Name, "{Quality}", q, "{slot}", slot)
	}
	return fill(panelCopy.TileAltNoQuality, "{Item name}", it.Name, "{slot}", slot)
}

// crestAlt is a class crest's alt text.
func crestAlt(class string) string { return fill(panelCopy.CrestAlt, "{Class}", classLabel(class)) }

// crestURL is a class's ringed HD crest on the site.
func (s *Service) crestURL(class string) string {
	return s.d.PublicBaseURL + "/icons/hd/crests/" + class + ".webp"
}

// crestThumb is the header's crest accessory; nil for a class the site has no crest for.
func (s *Service) crestThumb(class string) *Node {
	if _, known := classColor(class); !known {
		return nil
	}
	return thumbnailNode(s.crestURL(class), crestAlt(class))
}

// qualityDot is the quality's application emoji and a space (spec §A.4.1), or "" when the
// quality is unknown or its emoji is not loaded: the quality word on line 3 always carries it.
func (s *Service) qualityDot(ctx context.Context, quality int) string {
	name := sheet.QualityEmojiName(quality)
	if name == "" {
		return ""
	}
	id := s.emoji.Marks(ctx)[name]
	if id == "" {
		return ""
	}
	return fmt.Sprintf("<:%s:%s> ", name, id)
}

// replacesLine is /char's line 4: "Replaces {worn name}, item level {n}", without the level
// when it is unknown, "" without a name.
func replacesLine(name string, level int) string {
	if name == "" {
		return ""
	}
	if level <= 0 {
		return fill(panelCopy.ReplacesNoLevel, "{worn name}", memberText(name))
	}
	return fill(panelCopy.Replaces, "{worn name}", memberText(name), "{n}", fmt.Sprint(level))
}
