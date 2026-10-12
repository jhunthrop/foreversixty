package sims

// The band's picks: what the leveling-bis ladder (sim/cmd/leveling-bis) chose for one spec in
// one level band, as the data addon ships them (data/pipeline/addonbis.py builds Data.lua's
// bis table from the same <TreeDataDir>/<build>/bis/<spec>.json). Nothing here ranks an
// item: the file's order is the ladder's own, and this only reads it through the data
// addon's rules:
//
//   - the bare preset only (level 60 also carries a "raid" entry per faction, which the
//     addon's table leaves out), for the character's own faction;
//   - task L1.4's walk: a candidate whose stats the client has not confirmed
//     (trees.Item.ClientUnconfirmed, items.json's client_unconfirmed) is never the pick;
//     the next confirmed candidate in the ladder's order takes its place (the pick, then
//     runners_up, then the alternatives past them), and a promoted candidate carries no
//     drop chance, because the row's chance was the original pick's;
//   - the band-wide rules L1.4's review added, so a promotion never breaks the set: the
//     second slot of a pair (finger2, trinket2) never takes the unique item its mate
//     (finger1, trinket1) already picked, and the off-hand has no pick beside a
//     two-handed main-hand pick.
//
// A slot the file has no row for (or a row with no item_id: nothing eligible at the band),
// or whose candidates are all unconfirmed, has no pick.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// BarePreset is leveling-bis's preset for a character with no raid buffs, the one the data
// addon's picks come from.
const BarePreset = "bare"

// GearSlots are the 17 slots an FS1 export's gear names, in the export's own order
// (Export.INVENTORY_SLOTS, bnetbuild.SLOTS): the order the band's rules are applied in, so a
// main hand is picked before its off-hand and finger1 before finger2.
var GearSlots = [...]string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
	"finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
}

// pairMates maps the second slot of a pair onto the first.
var pairMates = map[string]string{"finger2": "finger1", "trinket2": "trinket1"}

// BandPick is one slot's pick in a band.
type BandPick struct {
	ItemID int
	// Name is the item's name in the build's tables, else the ladder's own item_name.
	Name string
	// Source is the plain name a player knows the item's source by ("Wailing Caverns:
	// Mutanus the Devourer", "Blacksmithing", "The League of Arathor").
	Source string
	// Chance is the drop fraction (0-1] the ladder states for the original pick; nil for a
	// source with no stated rate and for a promoted candidate.
	Chance *float64
	// Promoted is true when an unconfirmed candidate gave way to this one.
	Promoted bool
}

// bisCandidate is one runner-up or alternative of a bis row.
type bisCandidate struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
	Source   string `json:"source"`
}

// bisRow is one slot of one (band, faction, preset) entry of bis/<spec>.json.
type bisRow struct {
	Slot         string          `json:"slot"`
	ItemID       *int            `json:"item_id"`
	ItemName     string          `json:"item_name"`
	Source       string          `json:"source"`
	Chance       *float64        `json:"chance"`
	RunnersUp    *[]bisCandidate `json:"runners_up"`
	Alternatives []bisCandidate  `json:"alternatives"`
}

// bisEntry is one (band, faction, preset) entry of bis/<spec>.json.
type bisEntry struct {
	Band    int      `json:"band"`
	Faction string   `json:"faction"`
	Preset  string   `json:"preset"`
	Slots   []bisRow `json:"slots"`
}

// BandPicks is spec's pick per slot for band and faction in build (the bare preset, through
// the walk and band-wide rules above), keyed by slot; a slot with no pick has no key.
// ErrNoBestInSlot when the build has no bis file for spec or the file has no bare entry for
// band and faction; ErrGearUnscorable when no gear data is loaded or the build or the
// spec's class is unknown.
func BandPicks(ctx context.Context, build, spec string, band int, faction string) (map[string]BandPick, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	src := currentGearData()
	if src == nil {
		return nil, fmt.Errorf("%w: no gear data is loaded", ErrGearUnscorable)
	}
	b, ok := src.trees.Build(build)
	if !ok {
		return nil, fmt.Errorf("%w: no data for client build %q", ErrGearUnscorable, build)
	}
	classSlug, _, _ := strings.Cut(spec, "-")
	classID, ok := classIDBySlug(b, classSlug)
	if !ok {
		return nil, fmt.Errorf("%w: no class %q in build %s", ErrGearUnscorable, classSlug, build)
	}
	entries, err := src.bisEntries(build, spec)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		preset := e.Preset
		if preset == "" {
			preset = BarePreset
		}
		if e.Band == band && e.Faction == faction && preset == BarePreset {
			return walkBand(e.Slots, func(id int) (trees.Item, bool) { return b.Item(classID, id) }), nil
		}
	}
	return nil, fmt.Errorf("%w: %s band %d %s in build %s", ErrNoBestInSlot, spec, band, faction, build)
}

// walkBand applies the data addon's rules to one bare entry's rows (see the file comment).
// item looks an id up in the spec's class tables; an id they do not have counts as
// confirmed, not unique and one-handed, as items.json's absence of a flag does for L1.4.
func walkBand(rows []bisRow, item func(int) (trees.Item, bool)) map[string]BandPick {
	bySlot := make(map[string]bisRow, len(rows))
	for _, r := range rows {
		if r.ItemID != nil {
			bySlot[r.Slot] = r
		}
	}
	picks := map[string]BandPick{}
	for _, slot := range GearSlots {
		row, ok := bySlot[slot]
		if !ok {
			continue
		}
		if slot == "off_hand" {
			if mh, ok := picks["main_hand"]; ok {
				if it, ok := item(mh.ItemID); ok && it.TwoHand {
					continue
				}
			}
		}
		mate, hasMate := picks[pairMates[slot]]
		for i, c := range rowCandidates(row) {
			it, known := item(c.ItemID)
			if c.ItemID <= 0 || (known && it.ClientUnconfirmed) {
				continue
			}
			if hasMate && mate.ItemID == c.ItemID && known && it.Unique {
				continue
			}
			pick := BandPick{ItemID: c.ItemID, Name: c.ItemName, Source: c.Source, Promoted: i > 0}
			if known && it.Name != "" {
				pick.Name = it.Name
			}
			if i == 0 {
				pick.Chance = row.Chance
			}
			picks[slot] = pick
			break
		}
	}
	return picks
}

// rowCandidates is the row's candidates in the ladder's order: the pick, then runners_up,
// then the alternatives past them (a file written before runners_up existed: the
// alternatives), as addonbis._runners_up reads them.
func rowCandidates(row bisRow) []bisCandidate {
	out := []bisCandidate{{ItemID: *row.ItemID, ItemName: row.ItemName, Source: row.Source}}
	if row.RunnersUp == nil {
		return append(out, row.Alternatives...)
	}
	seen := map[int]bool{}
	for _, r := range *row.RunnersUp {
		seen[r.ItemID] = true
		// A runner-up carries no item_name; the alternatives row for the same id does.
		for _, a := range row.Alternatives {
			if a.ItemID == r.ItemID && r.ItemName == "" {
				r.ItemName = a.ItemName
			}
		}
		out = append(out, r)
	}
	for _, a := range row.Alternatives {
		if !seen[a.ItemID] {
			out = append(out, a)
		}
	}
	return out
}

// bisEntries is loadBisEntries of build and spec, kept once read; an error is not kept.
func (g *gearData) bisEntries(build, spec string) ([]bisEntry, error) {
	g.mu.Lock()
	entries, ok := g.bisRows[build][spec]
	g.mu.Unlock()
	if ok {
		return entries, nil
	}
	entries, err := loadBisEntries(g.dir, build, spec)
	if err != nil {
		return nil, err
	}
	g.mu.Lock()
	if g.bisRows[build] == nil {
		g.bisRows[build] = map[string][]bisEntry{}
	}
	g.bisRows[build][spec] = entries
	g.mu.Unlock()
	return entries, nil
}

// loadBisEntries reads <dir>/<build>/bis/<spec>.json's entries with their rows. A missing
// file is ErrNoBestInSlot.
func loadBisEntries(dir, build, spec string) ([]bisEntry, error) {
	body, err := os.ReadFile(filepath.Join(dir, build, "bis", spec+".json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%w: %s in build %s", ErrNoBestInSlot, spec, build)
	}
	if err != nil {
		return nil, fmt.Errorf("sims: read best-in-slot for %s/%s: %w", build, spec, err)
	}
	var file struct {
		Bands []bisEntry `json:"bands"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return nil, fmt.Errorf("sims: decode best-in-slot for %s/%s: %w", build, spec, err)
	}
	return file.Bands, nil
}

// ItemGearScore is Gear.score of one item for spec in build: the per-item term GearScore
// sums over a character's gear and FullGearScore over a best-in-slot set (statScore of the
// item's stat line in the spec's class tables, under the spec's weights). ok is false when
// the class tables do not have the item. ErrGearUnscorable when no gear data is loaded,
// the build or class is unknown, or the spec has no weights.
func ItemGearScore(ctx context.Context, build, spec string, itemID int) (float64, bool, error) {
	if err := ctx.Err(); err != nil {
		return 0, false, err
	}
	src := currentGearData()
	if src == nil {
		return 0, false, fmt.Errorf("%w: no gear data is loaded", ErrGearUnscorable)
	}
	b, classID, weights, err := src.specTables(build, spec)
	if err != nil {
		return 0, false, err
	}
	item, ok := b.Item(classID, itemID)
	if !ok {
		return 0, false, nil
	}
	return statScore(item.Stats, weights), true, nil
}

// specTables is the build, the spec's class id in it and the spec's weights: what scoring
// one item for spec needs.
func (g *gearData) specTables(build, spec string) (*trees.Build, int, map[string]float64, error) {
	b, ok := g.trees.Build(build)
	if !ok {
		return nil, 0, nil, fmt.Errorf("%w: no data for client build %q", ErrGearUnscorable, build)
	}
	classSlug, _, _ := strings.Cut(spec, "-")
	classID, ok := classIDBySlug(b, classSlug)
	if !ok {
		return nil, 0, nil, fmt.Errorf("%w: no class %q in build %s", ErrGearUnscorable, classSlug, build)
	}
	weights, err := g.specWeights(build, spec)
	if err != nil {
		return nil, 0, nil, err
	}
	return b, classID, weights, nil
}

// GearItem is an item in build's tables for the class with slug classSlug; ok is false when
// no gear data is loaded or the build, class or item is unknown.
func GearItem(build, classSlug string, itemID int) (trees.Item, bool) {
	src := currentGearData()
	if src == nil {
		return trees.Item{}, false
	}
	b, ok := src.trees.Build(build)
	if !ok {
		return trees.Item{}, false
	}
	classID, ok := classIDBySlug(b, classSlug)
	if !ok {
		return trees.Item{}, false
	}
	return b.Item(classID, itemID)
}

// RaceFaction is the faction ("alliance", "horde") of the race with slug raceSlug in build;
// "" when no gear data is loaded or the build or race is unknown.
func RaceFaction(build, raceSlug string) string {
	src := currentGearData()
	if src == nil || raceSlug == "" {
		return ""
	}
	b, ok := src.trees.Build(build)
	if !ok {
		return ""
	}
	for _, r := range b.Races() {
		if r.Slug == raceSlug {
			return r.Faction
		}
	}
	return ""
}
