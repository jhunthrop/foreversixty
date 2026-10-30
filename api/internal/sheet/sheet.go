// Package sheet draws the Discord character sheet and the item tiles (spec
// design/specs/2026-10-29-discord-gear-panel.md, Part B, §A.4 and Part C): a 1200 x 1688 PNG
// laid out like the in-game character frame, posted on Share in channel, and the 72 px icon
// tiles the panel's thumbnails point at.
//
// It draws directly onto an image with golang.org/x/image, at scale 2 over the spec's 600 x
// 844 logical box, with the site's faces (fonts/, Barlow 400 and 600, Cinzel 700, OFL) and
// art (assets/, the HD crests, faction marks and seal, copied byte for byte from web/public).
// Item icons come from an IconSource: the active build's icon folder in production.
package sheet

import (
	"errors"
	"fmt"
	"strconv"
)

// The two sheets.
const (
	KindBis  = "bis"
	KindChar = "char"
)

// Slot is one gear slot's cell.
//
//   - Slot is the BiS data's slot key ("head", "finger1", "main_hand").
//   - Name, Quality and Icon are the item's; an Icon of "" or a Quality outside 0 to 5 is an
//     item unknown to the build: the question mark in an empty-coloured ring, the name in the
//     text colour.
//   - Meta is the /bis source line in full; the renderer finds its short form (ShortSource).
//     A /char cell's meta is "item level {ItemLevel}" and Meta is ignored.
//   - Gain is the /char gold line, already worded: "+9.3 DPS", "+4.0 points",
//     CopyNotModelled; on an empty /char slot, CopyUpgrade or CopyNotModelled (or "" when the
//     band has no pick).
//   - Empty is a slot with no item: "{Slot} · {Gain}" with Gain in gold, else "{Slot} · Empty"
//     on /char and "{Slot} · No pick" on /bis.
//   - Every meta line starts with the slot's label ("Ring 1 · Vendor: Anachronos"), falling back
//     to the label alone, so every cell names its slot.
type Slot struct {
	Slot      string
	ItemID    int
	Name      string
	Quality   int
	Icon      string
	ItemLevel int
	Meta      string
	Gain      string
	Empty     bool
}

// Sheet is the whole model of one sheet (spec §B.3).
//
//   - Class is the class slug ("hunter"): the crest and the title colour. Spec is its display
//     name ("Beast Mastery"). Faction is "alliance" or "horde".
//   - Kicker and Title are drawn as given: /bis "Best in slot" and "{Spec} {Class}"; /char
//     "Level {n} {Spec} {Class} · {Guild}" and the character's name.
//   - The centre column: /bis draws "Level {Level}" over Subtitle (the range, "Phase 1 raid");
//     /char draws "ITEM LEVEL" over Subtitle (the average, "63.5").
//   - Footer is the bottom line: /bis "Updated {d Mon yyyy}", /char the gold legend. The shared
//     /char sheet carries no sync date.
//   - Band is the band label ("level 60 Phase 1 raid"), for the /bis alt text; Name is the
//     character (the /char title) and Guild its guild, for the /char alt text.
//   - TwoHand leaves the off-hand cell out: the main hand fills both hands.
type Sheet struct {
	Kind                 string
	Class, Spec, Faction string
	Title, Kicker        string
	Subtitle             string
	Footer               string
	Level                int
	Band                 string
	Guild                string
	Slots                []Slot
	TwoHand              bool
}

// Render draws the sheet as a 1200 x 1688 PNG. It is safe to call concurrently.
func Render(s Sheet, icons IconSource) ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if icons == nil {
		return nil, errors.New("sheet: no icon source")
	}
	r, err := newRenderer(icons)
	if err != nil {
		return nil, err
	}
	img, err := r.draw(s)
	if err != nil {
		return nil, err
	}
	return encodePNG(img)
}

func (s Sheet) validate() error {
	if s.Kind != KindBis && s.Kind != KindChar {
		return fmt.Errorf("sheet: kind %q is neither %q nor %q", s.Kind, KindBis, KindChar)
	}
	if _, ok := classColors[s.Class]; !ok {
		return fmt.Errorf("sheet: class %q is not a class slug", s.Class)
	}
	if s.Faction != "alliance" && s.Faction != "horde" {
		return fmt.Errorf("sheet: faction %q is neither alliance nor horde", s.Faction)
	}
	seen := map[string]bool{}
	for _, sl := range s.Slots {
		if SlotLabel(sl.Slot) == "" {
			return fmt.Errorf("sheet: slot %q is not a gear slot", sl.Slot)
		}
		if seen[sl.Slot] {
			return fmt.Errorf("sheet: slot %q given twice", sl.Slot)
		}
		seen[sl.Slot] = true
	}
	return nil
}

// slot is the cell for key: the given one, or an empty slot when the model has none.
func (s Sheet) slot(key string) Slot {
	for _, sl := range s.Slots {
		if sl.Slot == key {
			return sl
		}
	}
	return Slot{Slot: key, Empty: true}
}

// known reports whether the cell's item is one the build knows (spec §B.3, "item unknown").
func (sl Slot) known() bool { return sl.Icon != "" && knownQuality(sl.Quality) }

// itemLevelMeta is a /char cell's meta line.
func itemLevelMeta(n int) string {
	if n <= 0 {
		return ""
	}
	return metaILvl + strconv.Itoa(n)
}
