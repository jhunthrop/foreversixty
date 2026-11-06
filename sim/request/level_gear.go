package request

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// LevelGear fits a band's gear to a level below the band (round 25 platform).
// spec-breakdown's -level 10 to 19 runs the band-20 entry, and every slot whose
// item needs more than the level used to be left empty: rogues fought unarmed
// below 19 and hunters had no bow below 20. It also gated on the client's
// required_level alone, so a quest reward the client lists at 0 (Forsaken
// Greataxe, a level-22 quest's) was held at 10. LevelGear gates on the level
// the ladder and the BiS ranker use (leveling.EffectiveRequiredLevel) and
// refills a weapon, shield or wand slot from the class's client item table
// with the ladder's own pickers.
type LevelGear struct {
	spec   string
	items  []buildItem
	byID   map[int]buildItem
	known  map[int]bool
	floors map[int]int
	// clientLevels is every item's client required_level (items.json), for an
	// id the class table does not carry.
	clientLevels map[int]int
}

// LoadLevelGear reads spec's class item table for build.
func LoadLevelGear(repoRoot, build, spec string) (LevelGear, error) {
	info, ok := specs.ByKey[spec]
	if !ok {
		return LevelGear{}, fmt.Errorf("level gear: unknown spec %q", spec)
	}
	items, err := loadClassItems(repoRoot, build, info.ClassSlug)
	if err != nil {
		return LevelGear{}, err
	}
	known, err := obtainableItemIDs(repoRoot, build)
	if err != nil {
		return LevelGear{}, err
	}
	floors, err := loadRequiredLevelFloors(repoRoot, build, items)
	if err != nil {
		return LevelGear{}, err
	}
	clientLevels, err := loadClientRequiredLevels(repoRoot, build)
	if err != nil {
		return LevelGear{}, err
	}
	byID := make(map[int]buildItem, len(items))
	for _, it := range items {
		byID[it.ID] = it
	}
	return LevelGear{spec: spec, items: items, byID: byID, known: known, floors: floors, clientLevels: clientLevels}, nil
}

func loadClientRequiredLevels(repoRoot, build string) (map[int]int, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "items.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("level gear: reading %s: %w", path, err)
	}
	var rows []struct {
		ID            int `json:"id"`
		RequiredLevel int `json:"required_level"`
	}
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("level gear: parsing %s: %w", path, err)
	}
	out := make(map[int]int, len(rows))
	for _, r := range rows {
		out[r.ID] = r.RequiredLevel
	}
	return out, nil
}

// RequiredLevel is the level a character needs to hold item id: its quest or
// recipe floor when the class table carries it, else the client's own value.
func (g LevelGear) RequiredLevel(id int) int {
	if it, ok := g.byID[id]; ok {
		return leveling.EffectiveRequiredLevel(it.RequiredLevel, g.floors[id])
	}
	return g.clientLevels[id]
}

// Fit is gear as a level-level character holds it. A slot it can wear keeps
// its item. A slot it cannot takes the first of alternatives[slot] (the band's
// own runners-up, best first) it can wear, else the class table's best weapon,
// shield or wand for that slot at the level, else stays empty. A weapon refill
// keeps the dropped item's handedness and, where one exists, its weapon type
// (a dagger stays a dagger for Backstab), and never repeats the main hand's
// item. A two-hand main hand leaves no off hand.
func (g LevelGear) Fit(gear []api.GearSlot, alternatives map[string][]int, level int) []api.GearSlot {
	out := make([]api.GearSlot, 0, len(gear))
	for _, s := range gear {
		if g.RequiredLevel(s.ItemID) <= level {
			out = append(out, s)
			continue
		}
		if id, ok := g.replacement(s, alternatives[s.Slot], out, level); ok {
			out = append(out, api.GearSlot{Slot: s.Slot, ItemID: id})
		}
	}
	return withoutOffHandBesideTwoHand(out, g.byID)
}

func (g LevelGear) replacement(dropped api.GearSlot, alternatives []int, worn []api.GearSlot, level int) (int, bool) {
	taken := slotItemIDs(worn)
	for _, id := range alternatives {
		if !taken[id] && id != 0 && g.RequiredLevel(id) <= level && g.sameHandedness(id, dropped.ItemID) {
			return id, true
		}
	}
	it, ok := g.byID[dropped.ItemID]
	if !ok {
		return 0, false
	}
	pool := withoutItem(g.items, worn)
	switch {
	case it.WeaponClass == itemClassArmor && it.WeaponSubclass == armorSubclassShield:
		pick, found := pickShieldItem(pool, g.known, g.floors, level)
		return pick.ID, found
	case it.WeaponClass != itemClassWeapon:
		return 0, false
	case it.WeaponSubclass == weaponWand:
		pick, found := pickWandItem(pool, g.known, g.floors, level)
		return pick.ID, found
	}
	return g.weapon(pool, dropped.Slot, it, level)
}

// sameHandedness is whether a and b are both two-handers or both not: a
// dual-wielder's main hand never takes a two-hander that would empty the off
// hand. An id the class table lacks matches anything.
func (g LevelGear) sameHandedness(a, b int) bool {
	x, okA := g.byID[a]
	y, okB := g.byID[b]
	return !okA || !okB || x.TwoHand == y.TwoHand
}

// weapon is the best weapon for slot at level of the dropped item's type, else
// of its family (any type of its handedness; any launcher for a launcher).
func (g LevelGear) weapon(pool []buildItem, slot string, dropped buildItem, level int) (int, bool) {
	hand := handOne
	if dropped.TwoHand {
		hand = handTwo
	}
	if slot == "ranged" {
		hand = handAny
	}
	stats, weightDPS := specWeaponPrimaryStats(g.spec), specWeaponDPSMatters(g.spec)
	for _, types := range [][]int{{dropped.WeaponSubclass}, weaponFamily(dropped.WeaponSubclass)} {
		if pick, ok := pickGearItem(pool, g.known, g.floors, slot, level, hand, types, stats, weightDPS); ok {
			return pick.ID, true
		}
	}
	return 0, false
}

// weaponWand is the client's wand subclass (items.json class 2, subclass 19).
const weaponWand = 19

// launcherSubclasses are the ranged weapons Auto Shot fires.
var launcherSubclasses = []int{weaponBow, weaponGun, weaponCrossbow}

// weaponFamily is the fallback type list for a refill: every launcher for a
// launcher, the thrown type for a thrown weapon, else nil (any type the
// handedness allows).
func weaponFamily(subclass int) []int {
	if containsInt(launcherSubclasses, subclass) {
		return launcherSubclasses
	}
	if subclass == weaponThrown {
		return []int{weaponThrown}
	}
	return nil
}

func slotItemIDs(gear []api.GearSlot) map[int]bool {
	out := make(map[int]bool, len(gear))
	for _, s := range gear {
		out[s.ItemID] = true
	}
	return out
}

func withoutOffHandBesideTwoHand(gear []api.GearSlot, byID map[int]buildItem) []api.GearSlot {
	twoHand := false
	for _, s := range gear {
		if s.Slot == "main_hand" && byID[s.ItemID].TwoHand {
			twoHand = true
		}
	}
	if !twoHand {
		return gear
	}
	out := make([]api.GearSlot, 0, len(gear))
	for _, s := range gear {
		if s.Slot != "off_hand" {
			out = append(out, s)
		}
	}
	return out
}
