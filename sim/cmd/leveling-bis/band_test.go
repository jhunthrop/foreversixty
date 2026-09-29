package main

import "testing"

func TestSourceForNoSource(t *testing.T) {
	idx := lootIndex{}
	if _, ok := sourceFor(1, 30, "alliance", "", idx); ok {
		t.Fatal("sourceFor with an empty index: want ok=false")
	}
}

func TestSourceForPicksHighestPriorityKind(t *testing.T) {
	// sourceKindPriority: quest, dungeon, crafted, rep, pvp, world, raid
	// - dungeon must win over world even though world was inserted
	// first, because the priority order (not insertion order) decides.
	idx := lootIndex{
		1: {
			{Kind: "world", Label: "World Vendor"},
			{Kind: "dungeon", Label: "A Dungeon"},
		},
	}
	src, ok := sourceFor(1, 30, "alliance", "", idx)
	if !ok || src.Kind != "dungeon" || src.Label != "A Dungeon" {
		t.Fatalf("sourceFor = %+v, %v, want dungeon/A Dungeon", src, ok)
	}
}

func TestSourceForRaidExcludedBelow60(t *testing.T) {
	idx := lootIndex{1: {{Kind: "raid", Label: "Molten Core"}}}
	if _, ok := sourceFor(1, 59, "alliance", "", idx); ok {
		t.Fatal("sourceFor at level 59 with only a raid source: want ok=false")
	}
	src, ok := sourceFor(1, 60, "alliance", "", idx)
	if !ok || src.Kind != "raid" {
		t.Fatalf("sourceFor at level 60 = %+v, %v, want the raid source", src, ok)
	}
}

func TestSourceForFallsBackWhenNoNonRaidKindPresent(t *testing.T) {
	idx := lootIndex{1: {{Kind: "raid", Label: "Molten Core"}}}
	if _, ok := sourceFor(1, 30, "alliance", "", idx); ok {
		t.Fatal("sourceFor with only a below-60-excluded raid source: want ok=false, not falling through to it anyway")
	}
}

func TestWeaponWithNoDPS(t *testing.T) {
	cases := []struct {
		name string
		c    candidate
		want bool
	}{
		{"a ranged weapon with dps 0 is flagged", candidate{DPS: 0, Slots: []string{"ranged"}}, true},
		{"a main_hand weapon with dps 0 is flagged", candidate{DPS: 0, Slots: []string{"main_hand"}}, true},
		{"a ranged weapon with real dps is not flagged", candidate{DPS: 10, Slots: []string{"ranged"}}, false},
		{"a non-weapon slot with dps 0 is not flagged", candidate{DPS: 0, Slots: []string{"head"}}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := weaponWithNoDPS(tc.c); got != tc.want {
				t.Errorf("weaponWithNoDPS(%+v) = %v, want %v", tc.c, got, tc.want)
			}
		})
	}
}

func TestCrossClassSetItem(t *testing.T) {
	rogueSet := 204 // setclass_generated.go: 204 -> "rogue"
	unknownSet := 999999999
	cases := []struct {
		name string
		c    candidate
		cls  string
		want bool
	}{
		{"no set id at all is never cross-class", candidate{SetID: nil}, "hunter", false},
		{"a rogue set worn by a hunter is cross-class", candidate{SetID: &rogueSet}, "hunter", true},
		{"a rogue set worn by a rogue is not cross-class", candidate{SetID: &rogueSet}, "rogue", false},
		{"a set id absent from setNativeClass is treated as unrestricted", candidate{SetID: &unknownSet}, "hunter", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := crossClassSetItem(tc.c, tc.cls); got != tc.want {
				t.Errorf("crossClassSetItem(%+v, %q) = %v, want %v", tc.c, tc.cls, got, tc.want)
			}
		})
	}
}

func TestBuildBandPoolSeparatesEligibleSourcedCrossClassAndUnsourced(t *testing.T) {
	rogueSet := 204
	items := []candidate{
		// Eligible, sourced, scorable: goes to Scored.
		{ID: 1, Name: "Sourced Helm", RequiredLevel: 10, EffectiveRequiredLevel: 10, Stats: map[string]float64{"agility": 1}, Slots: []string{"head"}},
		// Eligible but no loot.json source: goes to NoSource.
		{ID: 2, Name: "Mystery Cloak", RequiredLevel: 10, EffectiveRequiredLevel: 10, Slots: []string{"back"}},
		// Not eligible at all (required level too high): excluded
		// entirely, appears in neither bucket.
		{ID: 3, Name: "Too High Level", RequiredLevel: 90, EffectiveRequiredLevel: 90, Slots: []string{"waist"}},
		// A rogue-set item on a hunter pool: goes to CrossClassSet, not
		// Scored, even though it has a loot source.
		{ID: 4, Name: "Bonescythe Leggings", RequiredLevel: 10, EffectiveRequiredLevel: 10, SetID: &rogueSet, Slots: []string{"legs"}},
		// A weapon with no dps: still scored, but also collected into
		// NoDPSWeapon.
		{ID: 5, Name: "Blunt Bow", RequiredLevel: 10, EffectiveRequiredLevel: 10, DPS: 0, Slots: []string{"ranged"}},
	}
	idx := lootIndex{
		1: {{Kind: "quest", Label: "A Quest"}},
		4: {{Kind: "quest", Label: "A Quest"}},
		5: {{Kind: "quest", Label: "A Quest"}},
	}
	pool := buildBandPool(items, idx, "hunter", 20, "horde", map[string]float64{"agility": 2})

	if len(pool.Scored) != 2 {
		t.Fatalf("pool.Scored = %+v, want 2 (helm + bow)", pool.Scored)
	}
	if len(pool.NoSource) != 1 || pool.NoSource[0].ID != 2 {
		t.Fatalf("pool.NoSource = %+v, want just item 2", pool.NoSource)
	}
	if len(pool.CrossClassSet) != 1 || pool.CrossClassSet[0].ID != 4 {
		t.Fatalf("pool.CrossClassSet = %+v, want just item 4", pool.CrossClassSet)
	}
	if len(pool.NoDPSWeapon) != 1 || pool.NoDPSWeapon[0].ID != 5 {
		t.Fatalf("pool.NoDPSWeapon = %+v, want just item 5", pool.NoDPSWeapon)
	}
	for _, s := range pool.Scored {
		if s.ID == 1 && s.Score != 2 {
			t.Errorf("scored helm score = %v, want 2 (1 agility * weight 2)", s.Score)
		}
	}
}

func TestBuildBandPoolSkipsItemsWithNoSlots(t *testing.T) {
	// An item that resolved to zero planner slots (should not happen in
	// practice, but buildBandPool defends explicitly against it rather
	// than letting score() panic on Slots[0]) is skipped entirely, not
	// added to any bucket.
	items := []candidate{{ID: 1, RequiredLevel: 1, Slots: nil}}
	idx := lootIndex{1: {{Kind: "quest", Label: "A Quest"}}}
	pool := buildBandPool(items, idx, "hunter", 20, "horde", nil)
	if len(pool.Scored) != 0 || len(pool.NoSource) != 0 {
		t.Fatalf("pool = %+v, want everything empty for a slotless item", pool)
	}
}

// A battleground reputation's reward is obtainable only by its own side,
// and only at friendly/honored below 60: Outrunner's Bow (Warsong
// Outriders, revered) must never head an ALLIANCE level-20 list.
func TestSourceForGatesReputationBySideAndStanding(t *testing.T) {
	idx := lootIndex{20438: {{Kind: "rep", Label: "Warsong Outriders", Side: "horde", Standing: "revered"}}}
	if _, ok := sourceFor(20438, 20, "alliance", "", idx); ok {
		t.Fatal("a Horde reputation reward was offered to an alliance character")
	}
	if _, ok := sourceFor(20438, 20, "horde", "", idx); ok {
		t.Fatal("a revered reward was offered at level 20")
	}
	if _, ok := sourceFor(20438, 60, "horde", "", idx); !ok {
		t.Fatal("a revered reward must be obtainable by its own side at 60")
	}
	idx[7731] = []itemSource{{Kind: "rep", Label: "Silverwing Sentinels", Side: "alliance", Standing: "honored"}}
	if _, ok := sourceFor(7731, 20, "alliance", "", idx); !ok {
		t.Fatal("an honored reward of the character's own side is obtainable while leveling")
	}
}

// A dungeon inside the other side's capital is no source for this side:
// Ragefire Chasm's cloak falls through to a later source kind or to no
// source at all rather than heading an Alliance list.
func TestSourceForSkipsAFactionExclusiveDungeonForTheOtherSide(t *testing.T) {
	idx := lootIndex{14149: {{Kind: "dungeon", Label: "Ragefire Chasm: Taragaman the Hungerer", Side: "horde"}}}
	if _, ok := sourceFor(14149, 20, "alliance", "", idx); ok {
		t.Fatal("an Orgrimmar dungeon drop was offered to an alliance character")
	}
	if src, ok := sourceFor(14149, 20, "horde", "", idx); !ok || src.Kind != "dungeon" {
		t.Fatalf("horde source = %+v, %v; want the dungeon", src, ok)
	}
}

// Scout's Medallion (item 20442, horde_only) and Sentinel's Medallion
// (item 20444, alliance_only) are both real Forever items whose mined
// rep source has the WSG faction backwards versus the item's own
// client-stated restriction (wowhead: Scout's is Horde, sold by Kelm
// Hargunth in the Barrens; Sentinel's is Alliance, sold by Illiyana
// Moonblaze in Ashenvale - loot.json's rep source nonetheless lists
// Scout's under Silverwing Sentinels/alliance and Sentinel's under
// Warsong Outriders/horde). Without the item's own restriction
// breaking the tie, each Medallion is unobtainable by EITHER faction:
// its own side rejects it on the mined Side mismatch, and the other
// side never reaches sourceFor at all (eligible.go's own faction
// filter already excludes it there). This emptied hunter's level-20
// neck slot for both factions at once.
func TestSourceForTrustsItemFactionRestrictionOverAMinedRepSideMismatch(t *testing.T) {
	idx := lootIndex{
		20442: {{Kind: "rep", Label: "Silverwing Sentinels", Side: "alliance", Standing: "honored"}},
		20444: {{Kind: "rep", Label: "Warsong Outriders", Side: "horde", Standing: "honored"}},
	}
	if _, ok := sourceFor(20442, 20, "horde", "horde", idx); !ok {
		t.Fatal("Scout's Medallion (horde_only) must be obtainable by a horde character despite the mined rep source's alliance Side")
	}
	// An alliance query never reaches sourceFor for a horde_only item in
	// production - eligible.go's own FactionRestriction check excludes
	// it first - so sourceFor is not the layer responsible for that
	// gate; this override only widens what a MATCHING faction can see.
	if _, ok := sourceFor(20444, 20, "alliance", "alliance", idx); !ok {
		t.Fatal("Sentinel's Medallion (alliance_only) must be obtainable by an alliance character despite the mined rep source's horde Side")
	}
	// An unrestricted item still obeys the mined Side: the override only
	// applies when the item's OWN restriction names this faction.
	idx[1] = []itemSource{{Kind: "rep", Label: "Warsong Outriders", Side: "horde", Standing: "honored"}}
	if _, ok := sourceFor(1, 20, "alliance", "", idx); ok {
		t.Fatal("an unrestricted item's mined rep Side must still gate a mismatched faction")
	}
}
