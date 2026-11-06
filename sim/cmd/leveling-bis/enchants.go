package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"sort"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

// The permanent enchants a spec wears on the raid pass at level 60: what a
// fresh 60 can buy or have made in Phase 1 (data/curated/raid-enchants.json).
// The bare pass and the leveling bands stay unenchanted, because a
// levelling character does not re-enchant every few levels.
//
// The plan is keyed by the shape of what sits in a slot, not by the slot
// alone, so a two-hand enchant only ever reaches a two-hander, a shield
// enchant a shield, a weapon enchant a weapon, and nothing a
// held-in-off-hand item.

// Plan keys beyond the planner's own armor and ranged slot names.
const (
	enchantKeyMainHand = "main_hand"
	enchantKeyOffHand  = "off_hand"
	enchantKeyTwoHand  = "two_hand"
	enchantKeyShield   = "shield"
	shieldWeaponType   = "shield"
)

// enchantPlanKeys are every key raid-enchants.json may name, with the slot
// the enchant goes in and a representative item of that shape, so loading
// can check each enchant fits before any sim runs.
var enchantPlanKeys = map[string]struct {
	slot string
	item simdb.Item
}{
	"head":             {"head", simdb.Item{}},
	"neck":             {"neck", simdb.Item{}},
	"shoulder":         {"shoulder", simdb.Item{}},
	"back":             {"back", simdb.Item{}},
	"chest":            {"chest", simdb.Item{}},
	"wrist":            {"wrist", simdb.Item{}},
	"hands":            {"hands", simdb.Item{}},
	"waist":            {"waist", simdb.Item{}},
	"legs":             {"legs", simdb.Item{}},
	"feet":             {"feet", simdb.Item{}},
	"ranged":           {"ranged", simdb.Item{}},
	enchantKeyMainHand: {"main_hand", simdb.Item{HandType: simdb.HandOne}},
	enchantKeyOffHand:  {"off_hand", simdb.Item{HandType: simdb.HandOne}},
	enchantKeyTwoHand:  {"main_hand", simdb.Item{HandType: simdb.HandTwo}},
	enchantKeyShield:   {"off_hand", simdb.Item{WeaponType: shieldWeaponType}},
}

// enchantPlan is one spec's enchant id per plan key. A nil plan wears
// nothing.
type enchantPlan map[string]int

// raidEnchantsPath is where the curated plans live under the repository root.
func raidEnchantsPath(repoRoot string) string {
	return filepath.Join(repoRoot, "data", "curated", "raid-enchants.json")
}

type curatedEnchant struct {
	ID  int    `json:"id"`
	Why string `json:"why"`
}

type raidEnchantsFile struct {
	Profiles map[string]map[string]*curatedEnchant `json:"profiles"`
	Specs    map[string]struct {
		Profile string                     `json:"profile"`
		Slots   map[string]*curatedEnchant `json:"slots"`
	} `json:"specs"`
}

// enchantFits is the check a plan's enchants must pass: simdb.EnchantFits
// over the active build's enchants.json.
type enchantFits func(effectID int, slot string, item simdb.Item, class string) error

// loadRaidEnchants reads spec's plan from the curated file and checks
// every enchant against the active build's enchant table.
func loadRaidEnchants(repoRoot string, spec specInfo) (enchantPlan, error) {
	b, err := os.ReadFile(raidEnchantsPath(repoRoot))
	if err != nil {
		return nil, fmt.Errorf("reading curated/raid-enchants.json: %w", err)
	}
	return parseRaidEnchants(b, spec.Spec, spec.ClassSlug, simdb.EnchantFits)
}

// parseRaidEnchants is loadRaidEnchants over given bytes and fit check.
// A spec the file does not name, a profile it does not define, a key that
// is not a plan key, an enchant id the build does not carry or one that
// does not fit its key's slot and shape is an error: a typo stops the run
// rather than ranking without the enchant.
func parseRaidEnchants(b []byte, spec, class string, fits enchantFits) (enchantPlan, error) {
	var file raidEnchantsFile
	if err := json.Unmarshal(b, &file); err != nil {
		return nil, fmt.Errorf("decoding curated/raid-enchants.json: %w", err)
	}
	entry, ok := file.Specs[spec]
	if !ok {
		return nil, fmt.Errorf("curated/raid-enchants.json names no plan for %s", spec)
	}
	profile, ok := file.Profiles[entry.Profile]
	if !ok {
		return nil, fmt.Errorf("curated/raid-enchants.json: %s names unknown profile %q", spec, entry.Profile)
	}
	merged := map[string]*curatedEnchant{}
	for key, e := range profile {
		merged[key] = e
	}
	for key, e := range entry.Slots {
		merged[key] = e
	}
	plan := enchantPlan{}
	var errs []error
	for _, key := range sortedKeys(merged) {
		e := merged[key]
		if e == nil {
			continue
		}
		shape, ok := enchantPlanKeys[key]
		if !ok {
			errs = append(errs, fmt.Errorf("%s: %q is not a slot key", spec, key))
			continue
		}
		if err := fits(e.ID, shape.slot, shape.item, class); err != nil {
			errs = append(errs, fmt.Errorf("%s %s: enchant %d: %w", spec, key, e.ID, err))
			continue
		}
		plan[key] = e.ID
	}
	if err := errors.Join(errs...); err != nil {
		return nil, fmt.Errorf("curated/raid-enchants.json: %w", err)
	}
	return plan, nil
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// enchantKey is the plan key for item in slot: the slot itself for armor
// and ranged, and by shape for the hands. "" when the slot takes none (a
// ring, a trinket, a held-in-off-hand item).
func enchantKey(slot string, item simdb.Item) string {
	switch slot {
	case "main_hand":
		if item.HandType == simdb.HandTwo {
			return enchantKeyTwoHand
		}
		return enchantKeyMainHand
	case "off_hand":
		if item.WeaponType == shieldWeaponType {
			return enchantKeyShield
		}
		if item.HandType == simdb.HandNone {
			return ""
		}
		return enchantKeyOffHand
	}
	if _, ok := enchantPlanKeys[slot]; ok {
		return slot
	}
	return ""
}

// lookupItem resolves an item id to its simdb shape.
type lookupItem func(id int) (simdb.Item, bool)

// wear is gear with the plan's enchant on every slot that takes one, as a
// new slice. A slot that already names an enchant keeps it; an item the
// build's database does not carry is left alone (the engine unequips it).
func (p enchantPlan) wear(class string, gear []api.GearSlot) []api.GearSlot {
	return p.wearWith(class, gear, simdb.Lookup, simdb.EnchantFits)
}

// wearWith is wear over a given item lookup and fit check. The plan was
// checked against every shape when it loaded, so an enchant that does not
// fit here is a bug in enchantKey, and it panics rather than sim without it.
func (p enchantPlan) wearWith(class string, gear []api.GearSlot, lookup lookupItem, fits enchantFits) []api.GearSlot {
	if len(p) == 0 {
		return gear
	}
	out := make([]api.GearSlot, len(gear))
	for i, g := range gear {
		out[i] = g
		if g.Enchant != 0 || g.ItemID == 0 {
			continue
		}
		item, ok := lookup(g.ItemID)
		if !ok {
			continue
		}
		id := p[enchantKey(g.Slot, item)]
		if id == 0 {
			continue
		}
		if err := fits(id, g.Slot, item, class); err != nil {
			panic(fmt.Sprintf("leveling-bis: raid enchant %d on %s item %d: %v", id, g.Slot, g.ItemID, err))
		}
		out[i].Enchant = id
	}
	return out
}

// withRaidEnchants is ch wearing spec's raid enchants on the raid pass, and
// ch unchanged on every other pass.
func withRaidEnchants(spec specInfo, ch api.CharacterSpec) api.CharacterSpec {
	if spec.Applied == nil {
		return ch
	}
	ch.Gear = spec.RaidEnchants.wear(spec.ClassSlug, ch.Gear)
	return ch
}

// withWeaponEnchantCredit is items with each weapon's planned enchant added
// to its stats on the raid pass, as a new slice with new stat maps. score()
// reads item stats alone, so without it a staff would lose the main hand on
// its bare stats and never be simmed wearing +65 spell power, which is the
// whole case for a two-hander. It is an estimate: which weapon wins is still
// the sims' decision (verify.go). A one-hander takes the main_hand key's
// stats whichever hand it ends up in.
func withWeaponEnchantCredit(items []candidate, spec specInfo) []candidate {
	if spec.Applied == nil {
		return items
	}
	twoHand := simdb.EnchantStats(spec.RaidEnchants[enchantKeyTwoHand])
	oneHand := simdb.EnchantStats(spec.RaidEnchants[enchantKeyMainHand])
	if len(twoHand) == 0 && len(oneHand) == 0 {
		return items
	}
	out := make([]candidate, len(items))
	for i, c := range items {
		out[i] = c
		if c.ClassID != itemClassWeapon {
			continue
		}
		credit := oneHand
		if c.TwoHand {
			credit = twoHand
		}
		if len(credit) == 0 {
			continue
		}
		stats := maps.Clone(c.Stats)
		if stats == nil {
			stats = make(map[string]float64, len(credit))
		}
		for stat, amount := range credit {
			stats[stat] += amount
		}
		out[i].Stats = stats
	}
	return out
}

// withEnchantIDs is a raid pass's report rows with the enchant each
// published item wears, so a page can show it. Other passes are unchanged.
func withEnchantIDs(spec specInfo, rows []slotRow) []slotRow {
	if spec.Applied == nil || len(spec.RaidEnchants) == 0 {
		return rows
	}
	gear := make([]api.GearSlot, 0, len(rows))
	for _, r := range rows {
		gear = append(gear, api.GearSlot{Slot: r.Slot, ItemID: r.ItemID})
	}
	worn := spec.RaidEnchants.wear(spec.ClassSlug, gear)
	out := make([]slotRow, len(rows))
	for i, r := range rows {
		out[i] = r
		out[i].EnchantID = worn[i].Enchant
	}
	return out
}
