package request

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// 2026-09-28 quest-levels lane: pickGearItem/pickShieldItem/pickWandItem
// must gate on leveling.EffectiveRequiredLevel(RequiredLevel,
// floors[id]), not RequiredLevel alone -- a quest reward or crafted
// item's own required_level is 0 in the client far more often than
// not. These are synthetic-fixture unit tests, unlike TestRotationLadder
// above (a real-build integration test), so this lane's fix is proven
// directly rather than only through whichever real items happen to be
// quest rewards on the committed build.

func writeLootJSONFixture(t *testing.T, dir string, quests map[string]int, craftedIDs []int) {
	t.Helper()
	type questEntry struct {
		MinLevel int `json:"min_level"`
	}
	doc := map[string]any{
		"sources": []map[string]any{
			{"kind": "crafted", "items": craftedIDs},
		},
		"quests": map[string][]questEntry{},
	}
	quests_ := doc["quests"].(map[string][]questEntry)
	for id, minLevel := range quests {
		quests_[id] = []questEntry{{MinLevel: minLevel}}
	}
	b, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	buildDir := filepath.Join(dir, "data", "builds", "testbuild")
	if err := os.MkdirAll(buildDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(buildDir, "loot.json"), b, 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadRequiredLevelFloorsRaisesAQuestRewardAndFallsBackForACraftedItem(t *testing.T) {
	dir := t.TempDir()
	// Item 274271 (Deadhead Blade shape): required_level 0, but its
	// quest's own min_level is 20 -- this lane's fix means the ladder
	// never equips it on a level-10 character.
	writeLootJSONFixture(t, dir, map[string]int{"274271": 20}, []int{999})
	items := []buildItem{
		{ID: 274271, RequiredLevel: 0, ItemLevel: 58},
		{ID: 999, RequiredLevel: 0, ItemLevel: 44}, // crafted, no recipe skill level in loot.json
		{ID: 1, RequiredLevel: 30, ItemLevel: 40},  // an ordinary drop: untouched
	}
	floors, err := loadRequiredLevelFloors(dir, "testbuild", items)
	if err != nil {
		t.Fatalf("loadRequiredLevelFloors: %v", err)
	}
	if floors[274271] != 20 {
		t.Errorf("floors[274271] = %d, want 20 (the quest's min_level)", floors[274271])
	}
	// leveling.ItemLevelProxyRequiredLevel(44) = min(60, 39) = 39.
	if floors[999] != 39 {
		t.Errorf("floors[999] = %d, want 39 (item_level_proxy fallback)", floors[999])
	}
	if _, ok := floors[1]; ok {
		t.Errorf("floors[1] = %d, want absent (not a quest or crafted item)", floors[1])
	}
}

func TestPickGearItemRefusesADeadheadBladeShapedQuestRewardBelowItsQuestLevel(t *testing.T) {
	items := []buildItem{
		{
			ID: 274271, Slot: "main_hand", RequiredLevel: 0, ItemLevel: 58,
			WeaponClass: itemClassWeapon, WeaponSubclass: weaponDagger, Speed: 1.8, DamageMax: 30,
		},
		{
			ID: 100, Slot: "main_hand", RequiredLevel: 8, ItemLevel: 12,
			WeaponClass: itemClassWeapon, WeaponSubclass: weaponDagger, Speed: 1.6, DamageMax: 10,
		},
	}
	known := map[int]bool{274271: true, 100: true}
	floors := map[int]int{274271: 20} // the quest's min_level

	// A level-10 character: the quest-level-60-shaped item is refused
	// even though its own RequiredLevel is 0, and the pick falls back
	// to the level-appropriate item instead of hanging with no pick at
	// all.
	got, ok := pickGearItem(items, known, floors, "main_hand", 10, handAny, nil, nil, false)
	if !ok || got.ID != 100 {
		t.Fatalf("pickGearItem at level 10 = (%+v, %v), want item 100 (Deadhead Blade refused)", got, ok)
	}

	// At level 20 (the quest's own min_level), the item becomes pickable
	// and, being the higher item level (neither candidate carries any
	// stat, so weaponScore falls all the way through to its itemLevel
	// tiebreak), wins.
	got, ok = pickGearItem(items, known, floors, "main_hand", 20, handAny, nil, nil, false)
	if !ok || got.ID != 274271 {
		t.Fatalf("pickGearItem at level 20 = (%+v, %v), want item 274271", got, ok)
	}

	// With no floors map at all (nil), the old required_level-only
	// behaviour is exactly what leveling.EffectiveRequiredLevel(0, 0)
	// reduces to: eligible at any level. Documented here so a future
	// change to the zero-value behaviour is a deliberate, visible one.
	got, ok = pickGearItem(items, known, nil, "main_hand", 1, handAny, nil, nil, false)
	if !ok || got.ID != 274271 {
		t.Fatalf("pickGearItem with nil floors at level 1 = (%+v, %v), want item 274271 (undefended without the floors map)", got, ok)
	}
}
