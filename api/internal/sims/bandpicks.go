package sims

// The band's picks: what the leveling-bis ladder (sim/cmd/leveling-bis) chose for one spec in
// one level band, as the data addon ships them (data/pipeline/addonbis.py builds Data.lua's
// bis table from the same <TreeDataDir>/<build>/bis/<spec>.json). Nothing here ranks an
// item: the file's order is the ladder's own, and this only reads it through the data
// addon's rules:
//
//   - the bare preset only (level 60 also carries a "raid" entry per faction, which the
//     addon's table leaves out), for the character's own faction;
//   - task L1.4's two-phase walk (addonbis._faction_picks): a candidate whose stats the
//     client has not confirmed (trees.Item.ClientUnconfirmed, items.json's
//     client_unconfirmed) is never the pick. Every slot whose own pick is confirmed keeps
//     it, fixed first; then each other slot, in the slot order, promotes its first
//     confirmed candidate in the ladder's order (the pick, then runners_up, then the
//     alternatives past them) that the band-wide rules allow beside the picks already
//     fixed. A promoted pick carries its own drop chance at its own source when loot.json
//     states one (addonbis._drop_chances), else none;
//   - the band-wide rules (leveling-bis's verify.go pairSlot and confirmed.go
//     usableAlternative): a slot never holds its pair mate's item (finger1/finger2,
//     trinket1/trinket2, main_hand/off_hand), the same id or the same name, unique or
//     not; the off-hand has no pick beside a two-handed main-hand pick, and a two-hander
//     promoted into the main hand empties the off-hand, as confirmed.go does.
//
// A slot the file has no row for (or a row with no item_id: nothing eligible at the band),
// or whose candidates are all unconfirmed, has no pick.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
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

// pairSlot is leveling-bis's verify.go pairSlot: the slots that may never hold the same item
// as each other.
var pairSlot = map[string]string{
	"finger1": "finger2", "finger2": "finger1",
	"trinket1": "trinket2", "trinket2": "trinket1",
	"main_hand": "off_hand", "off_hand": "main_hand",
}

// BandPick is one slot's pick in a band.
type BandPick struct {
	ItemID int
	// Name is the item's name in the build's tables, else the ladder's own item_name.
	Name string
	// Source is the plain name a player knows the item's source by ("Wailing Caverns:
	// Mutanus the Devourer", "Blacksmithing", "The League of Arathor").
	Source string
	// Chance is the drop fraction (0-1]: the ladder's own for the original pick, loot.json's
	// at the candidate's own source for a promoted one; nil when the source states no rate.
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
	chances, err := src.dropChances(build)
	if err != nil {
		return nil, err
	}
	chanceOf := func(id int, source string) *float64 {
		if c, ok := chances[lootKey{id, source}]; ok {
			return &c
		}
		return nil
	}
	for _, e := range entries {
		preset := e.Preset
		if preset == "" {
			preset = BarePreset
		}
		if e.Band == band && e.Faction == faction && preset == BarePreset {
			// The spec's class tables answer for every candidate: an id they do not
			// have counts as confirmed, one-handed and nameless (so it matches its
			// pair mate by id only), as items.json's absence of a flag does for L1.4.
			return walkBand(e.Slots, func(id int) (trees.Item, bool) { return b.Item(classID, id) }, chanceOf), nil
		}
	}
	return nil, fmt.Errorf("%w: %s band %d %s in build %s", ErrNoBestInSlot, spec, band, faction, build)
}

// walkBand applies the data addon's rules to one bare entry's rows (see the file comment).
// item looks an id up in the spec's class tables; chance is a promoted candidate's drop
// chance at its own source.
func walkBand(rows []bisRow, item func(int) (trees.Item, bool), chance func(int, string) *float64) map[string]BandPick {
	candidates := map[string][]bisCandidate{}
	ownChance := map[string]*float64{}
	for _, r := range rows {
		if r.ItemID != nil {
			candidates[r.Slot] = rowCandidates(r)
			ownChance[r.Slot] = r.Chance
		}
	}
	unconfirmed := func(id int) bool {
		it, ok := item(id)
		return id <= 0 || (ok && it.ClientUnconfirmed)
	}
	twoHand := func(id int) bool {
		it, ok := item(id)
		return ok && it.TwoHand
	}
	// same is verify.go's wouldDuplicatePairMate test: the same id, or the same name.
	same := func(a, b int) bool {
		if a == b {
			return true
		}
		ia, okA := item(a)
		ib, okB := item(b)
		return okA && okB && ia.Name != "" && ia.Name == ib.Name
	}
	pickOf := func(slot string, c bisCandidate, promoted bool) BandPick {
		p := BandPick{ItemID: c.ItemID, Name: c.ItemName, Source: c.Source, Promoted: promoted}
		if it, ok := item(c.ItemID); ok && it.Name != "" {
			p.Name = it.Name
		}
		if !promoted {
			p.Chance = ownChance[slot]
		} else if chance != nil {
			p.Chance = chance(c.ItemID, c.Source)
		}
		return p
	}

	// Phase one: every confirmed own pick is fixed.
	final := map[string]BandPick{}
	for _, slot := range GearSlots {
		if cs, ok := candidates[slot]; ok && !unconfirmed(cs[0].ItemID) {
			final[slot] = pickOf(slot, cs[0], false)
		}
	}
	// Phase two: each other slot promotes its first confirmed candidate the rules allow
	// beside the picks already final.
	for _, slot := range GearSlots {
		cs, ok := candidates[slot]
		if _, done := final[slot]; !ok || done {
			continue
		}
		for _, c := range cs[1:] {
			if unconfirmed(c.ItemID) {
				continue
			}
			if mate, ok := final[pairSlot[slot]]; ok && same(mate.ItemID, c.ItemID) {
				continue
			}
			if mh, ok := final["main_hand"]; slot == "off_hand" && ok && twoHand(mh.ItemID) {
				continue
			}
			final[slot] = pickOf(slot, c, true)
			break
		}
	}
	if mh, ok := final["main_hand"]; ok && twoHand(mh.ItemID) {
		delete(final, "off_hand")
	}
	return final
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

// lootKey is one (item id, source label) pair of loot.json.
type lootKey struct {
	itemID int
	source string
}

// dropChances is loadDropChances of build, kept once read; an error is not kept.
func (g *gearData) dropChances(build string) (map[lootKey]float64, error) {
	g.mu.Lock()
	c, ok := g.chances[build]
	g.mu.Unlock()
	if ok {
		return c, nil
	}
	c, err := loadDropChances(g.dir, build)
	if err != nil {
		return nil, err
	}
	g.mu.Lock()
	g.chances[build] = c
	g.mu.Unlock()
	return c, nil
}

// loadDropChances is addonbis._drop_chances: (item id, source label) -> drop chance as a
// fraction, indexed the way leveling-bis's loadLootIndex labels a source (a source's own
// items under its name, a boss's under "<source>: <boss>"; the flat quest source skipped)
// and converted as its dropChance does (percent to fraction, math.Round(chance*1e4)/1e6; a
// chance of 0 is none). The larger chance wins when one pair appears twice. A build with no
// loot.json has no chances.
func loadDropChances(dir, build string) (map[lootKey]float64, error) {
	body, err := os.ReadFile(filepath.Join(dir, build, lootFile))
	if errors.Is(err, os.ErrNotExist) {
		return map[lootKey]float64{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("sims: read loot for %s: %w", build, err)
	}
	var file struct {
		Sources []struct {
			Kind        string             `json:"kind"`
			Name        string             `json:"name"`
			Items       []int              `json:"items"`
			ItemChances map[string]float64 `json:"item_chances"`
			Bosses      []struct {
				Name        string             `json:"name"`
				Items       []int              `json:"items"`
				ItemChances map[string]float64 `json:"item_chances"`
			} `json:"bosses"`
		} `json:"sources"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return nil, fmt.Errorf("sims: decode loot for %s: %w", build, err)
	}
	out := map[lootKey]float64{}
	add := func(id int, label string, percent float64) {
		if percent <= 0 {
			return
		}
		fraction := math.Round(percent*1e4) / 1e6
		if k := (lootKey{id, label}); fraction > out[k] {
			out[k] = fraction
		}
	}
	for _, src := range file.Sources {
		if src.Kind == "quest" {
			continue
		}
		for _, id := range src.Items {
			add(id, src.Name, src.ItemChances[strconv.Itoa(id)])
		}
		for _, boss := range src.Bosses {
			label := src.Name
			if boss.Name != "" {
				label = src.Name + ": " + boss.Name
			}
			for _, id := range boss.Items {
				add(id, label, boss.ItemChances[strconv.Itoa(id)])
			}
		}
	}
	return out, nil
}
