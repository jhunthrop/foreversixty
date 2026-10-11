package sims

// What each piece of a raid's loot is worth to one character (docs/contracts/addon-v2.md
// contract 3, the loot_values message; spec §4.8). The value is a stat-line score delta,
// the same Gear.score arithmetic GearScore uses (gearscore.go), of the item against what
// the character's export wears in its slot, turned into DPS with the data lane's DPS per
// reference-stat point (data/builds/<build>/dps-per-point.json) when the spec has one.
//
// Both data files are read from the build directory SetGearData was given, once per build,
// and kept: the inbox asks every ten minutes and the files never change under a running API.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// The units a loot value or an upgrade delta is measured in (contract 3).
const (
	UnitDPS   = "dps"
	UnitScore = "score"
)

const (
	dpsPerPointFile = "dps-per-point.json"
	lootFile        = "loot.json"
)

// LootValue is one item's worth to the character.
type LootValue struct {
	// Value is the change to the character's gear in Unit, rounded to a tenth.
	Value float64
	Unit  string
	// Planned is true when the item is the spec's level-60 best-in-slot pick.
	Planned bool
}

// LootValueSet is a raid's priced loot: the raid's own name and every item the character
// can wear, by item id.
type LootValueSet struct {
	Raid  string
	Items map[int]LootValue
}

// raidLoot is one raid's loot table, every boss's items together.
type raidLoot struct {
	name  string
	items []int
}

// DPSPerPoint is the DPS (healing per second for a healer) one point of spec's reference
// stat is worth in build, from the build's dps-per-point.json; build "" is the newest
// loaded build. ok is false when the file has no entry for the spec (the tank specs), or
// no gear data is wired.
func DPSPerPoint(build, spec string) (float64, bool) {
	src := currentGearData()
	if src == nil {
		return 0, false
	}
	if build == "" {
		b, ok := src.trees.Latest()
		if !ok {
			return 0, false
		}
		build = b.Version
	}
	src.mu.Lock()
	defer src.mu.Unlock()
	bySpec, ok := src.dpsPerPoint[build]
	if !ok {
		bySpec = loadDPSPerPoint(src.dir, build)
		src.dpsPerPoint[build] = bySpec
	}
	v, ok := bySpec[spec]
	return v, ok && v > 0
}

// loadDPSPerPoint reads <dir>/<build>/dps-per-point.json; a missing or unreadable file is
// no entries, so every spec falls back to score units.
func loadDPSPerPoint(dir, build string) map[string]float64 {
	out := map[string]float64{}
	body, err := os.ReadFile(filepath.Join(dir, build, dpsPerPointFile))
	if err != nil {
		return out
	}
	var file map[string]struct {
		DPS float64 `json:"dps"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return out
	}
	for spec, e := range file {
		out[spec] = e.DPS
	}
	return out
}

// raidKey is how a raid's name is matched: lower case, letters and digits only, so the
// zone a report names ("Onyxia's Lair", "onyxias lair") finds loot.json's source.
func raidKey(name string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// raidLootFor finds the raid source of build's loot.json whose name matches zone.
func (g *gearData) raidLootFor(build, zone string) (raidLoot, bool, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	raids, ok := g.raids[build]
	if !ok {
		var err error
		if raids, err = loadRaidLoot(g.dir, build); err != nil {
			return raidLoot{}, false, err
		}
		g.raids[build] = raids
	}
	r, ok := raids[raidKey(zone)]
	return r, ok, nil
}

// loadRaidLoot reads the raid sources of <dir>/<build>/loot.json, keyed by raidKey of the
// raid's name, each with its bosses' items merged and sorted.
func loadRaidLoot(dir, build string) (map[string]raidLoot, error) {
	body, err := os.ReadFile(filepath.Join(dir, build, lootFile))
	if errors.Is(err, os.ErrNotExist) {
		return map[string]raidLoot{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("sims: read loot for %s: %w", build, err)
	}
	var file struct {
		Sources []struct {
			Kind   string `json:"kind"`
			Name   string `json:"name"`
			Bosses []struct {
				Items []int `json:"items"`
			} `json:"bosses"`
		} `json:"sources"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return nil, fmt.Errorf("sims: decode loot for %s: %w", build, err)
	}
	out := map[string]raidLoot{}
	for _, src := range file.Sources {
		if src.Kind != "raid" {
			continue
		}
		seen := map[int]bool{}
		r := raidLoot{name: src.Name}
		for _, boss := range src.Bosses {
			for _, id := range boss.Items {
				if id > 0 && !seen[id] {
					seen[id] = true
					r.items = append(r.items, id)
				}
			}
		}
		sort.Ints(r.items)
		out[raidKey(src.Name)] = r
	}
	return out, nil
}

// exportSlots are the export's slots an item of a given slot could replace: either ring or
// trinket, and for anything else the slot of the same name.
func exportSlots(itemSlot string) []string {
	switch itemSlot {
	case "finger":
		return []string{"finger1", "finger2"}
	case "trinket":
		return []string{"trinket1", "trinket2"}
	}
	return []string{itemSlot}
}

// LootValues prices every item on the loot tables of the raid named zone for the
// character whose FS1 export is gearCode, under spec: an item's value is its stat-line
// score less the score of what the export wears in that slot (the weaker ring or trinket;
// for a two-handed weapon, main and off hand together), under runWeights when the
// character has a weights run of its own and the build's shipped weights for the spec
// otherwise. It is in DPS when DPSPerPoint knows the spec (score times DPS per point) and
// in score units when it does not. An item the character's class item pool has no stat
// line for - one it cannot use, a quest item, a recipe - is skipped. ok is false when no
// raid in the build's loot.json is named zone.
//
// Ruling 9 of the L6.3 brief names this LootValues(raidID, characterKey); the caller
// already holds the character's export and spec, so it passes them rather than this
// package reading them back.
func LootValues(ctx context.Context, zone, gearCode, spec string, runWeights map[string]float64) (LootValueSet, bool, error) {
	if err := ctx.Err(); err != nil {
		return LootValueSet{}, false, err
	}
	src := currentGearData()
	if src == nil {
		return LootValueSet{}, false, fmt.Errorf("%w: no gear data is loaded", ErrGearUnscorable)
	}
	decoded, ok := fs1.Decode(gearCode)
	if !ok {
		return LootValueSet{}, false, fmt.Errorf("%w: not an FS1 code", ErrGearUnscorable)
	}
	build := decoded.DataBuild
	b, ok := src.trees.Build(build)
	if !ok {
		return LootValueSet{}, false, fmt.Errorf("%w: no data for client build %q", ErrGearUnscorable, build)
	}
	classID, ok := classIDBySlug(b, decoded.ClassSlug)
	if !ok {
		return LootValueSet{}, false, fmt.Errorf("%w: no class %q in build %s", ErrGearUnscorable, decoded.ClassSlug, build)
	}
	raid, ok, err := src.raidLootFor(build, zone)
	if err != nil || !ok {
		return LootValueSet{}, false, err
	}
	weights := runWeights
	if len(weights) == 0 {
		if weights, err = src.specWeights(build, spec); err != nil {
			return LootValueSet{}, false, err
		}
	}
	unit, scale := UnitScore, 1.0
	if dpp, ok := DPSPerPoint(build, spec); ok {
		unit, scale = UnitDPS, dpp
	}
	planned, err := plannedItems(src.dir, build, spec)
	if err != nil {
		return LootValueSet{}, false, err
	}

	worn := func(slot string) float64 {
		item, ok := b.Item(classID, decoded.Gear[slot])
		if !ok {
			return 0
		}
		return statScore(item.Stats, weights)
	}
	out := LootValueSet{Raid: raid.name, Items: map[int]LootValue{}}
	for _, id := range raid.items {
		item, ok := b.Item(classID, id)
		if !ok || item.Slot == "" || len(item.Stats) == 0 {
			continue
		}
		replaced := 0.0
		if item.TwoHand {
			replaced = worn("main_hand") + worn("off_hand")
		} else {
			for i, slot := range exportSlots(item.Slot) {
				if s := worn(slot); i == 0 || s < replaced {
					replaced = s
				}
			}
		}
		delta := (statScore(item.Stats, weights) - replaced) * scale
		out.Items[id] = LootValue{Value: math.Round(delta*10) / 10, Unit: unit, Planned: planned[id]}
	}
	return out, true, nil
}

// plannedItems is every item a level-60 best-in-slot set of spec's (bis/<spec>.json, one
// set per preset and faction) picks: what contract 3's planned means. A spec with no file
// plans nothing.
func plannedItems(dir, build, spec string) (map[int]bool, error) {
	planned := map[int]bool{}
	sets, err := loadBestInSlot(dir, build, spec)
	if errors.Is(err, ErrNoBestInSlot) {
		return planned, nil
	}
	if err != nil {
		return nil, err
	}
	for _, set := range sets {
		for _, id := range set {
			planned[id] = true
		}
	}
	return planned, nil
}
