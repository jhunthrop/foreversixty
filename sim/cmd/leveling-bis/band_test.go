package main

import "testing"

func TestSourceForNoSource(t *testing.T) {
	idx := lootIndex{}
	if _, ok := sourceFor(1, 30, "alliance", "", idx); ok {
		t.Fatal("sourceFor with an empty index: want ok=false")
	}
}

func TestSourceForPicksHighestPriorityKind(t *testing.T) {
	// quest, then boss (dungeon/raid), then rep, vendor, crafted,
	// world_drop, pvp, world - dungeon (a boss) must win over world even
	// though world was inserted first, because the priority order (not
	// insertion order) decides.
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

func TestSourceForPicksVendorOverWorldAndFallsBackToItAlone(t *testing.T) {
	// vendor sits right after quest, rep and boss in sourceKindPriority
	// (added by the 2026-09-28 night-bis-sources lane) and beats world/etc
	// -- but NOT dungeon/raid any more (wowhead-world-drops lane,
	// 2026-09-29: see TestSourceForPicksABossOverVendorRepAndCrafted
	// below for that case, which this test used to cover before boss
	// outranked vendor).
	idx := lootIndex{
		1: {
			{Kind: "world", Label: "World Vendor"},
			{Kind: "vendor", Label: "A Vendor"},
		},
	}
	src, ok := sourceFor(1, 30, "alliance", "", idx)
	if !ok || src.Kind != "vendor" || src.Label != "A Vendor" {
		t.Fatalf("sourceFor = %+v, %v, want vendor/A Vendor", src, ok)
	}

	// A vendor-only item (nothing else in the index) must still resolve -
	// this was the bug: vendor fell through every case and reported no
	// known source even though loot.json carried it.
	vendorOnly := lootIndex{2: {{Kind: "vendor", Label: "A Vendor"}}}
	src, ok = sourceFor(2, 30, "alliance", "", vendorOnly)
	if !ok || src.Kind != "vendor" {
		t.Fatalf("sourceFor with only a vendor source = %+v, %v, want ok=true, kind=vendor", src, ok)
	}
}

func TestSourceForPicksWorldDropOverPvpAndWorldButNotOverCrafted(t *testing.T) {
	// sourceKindPriority: ... crafted, world_drop, pvp, world --
	// world_drop is obtainable (auction house) at any level in its own
	// range, so it beats pvp/world (and a raid source, excluded below
	// 60 regardless), but a source naming an exact place (crafted) still
	// wins.
	idx := lootIndex{
		1: {
			{Kind: "raid", Label: "A Raid"},
			{Kind: "world", Label: "A World Mob"},
			{Kind: "pvp", Label: "A PvP Rank"},
			{Kind: "world_drop", Label: "World drop"},
		},
	}
	src, ok := sourceFor(1, 30, "alliance", "", idx)
	if !ok || src.Kind != "world_drop" || src.Label != "World drop" {
		t.Fatalf("sourceFor = %+v, %v, want world_drop/World drop", src, ok)
	}

	idx2 := lootIndex{
		2: {
			{Kind: "world_drop", Label: "World drop"},
			{Kind: "crafted", Label: "A Crafted Item"},
		},
	}
	src, ok = sourceFor(2, 30, "alliance", "", idx2)
	if !ok || src.Kind != "crafted" {
		t.Fatalf("sourceFor = %+v, %v, want crafted to win over world_drop", src, ok)
	}

	// A world_drop-only item (nothing else in the index) must still
	// resolve -- it is obtainable (buy it off the auction house), not
	// unsourced.
	worldDropOnly := lootIndex{3: {{Kind: "world_drop", Label: "World drop"}}}
	src, ok = sourceFor(3, 30, "alliance", "", worldDropOnly)
	if !ok || src.Kind != "world_drop" {
		t.Fatalf("sourceFor with only a world_drop source = %+v, %v, want ok=true, kind=world_drop", src, ok)
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

// This lane's brief (bis-ranker-integrity, 2026-09-29), item 3: a
// source data/curated/loot/forever-raid-phases.json curated a content
// phase for (Opens non-empty) is refused at EVERY band, including 60 -
// the real defect this guards against: all nine caster specs' band-60
// main_hand published Atiesh (a Naxxramas legendary quest reward) with
// every trinket slot's "verified" pick also a Naxxramas/BWL/AQ trinket,
// because sourceFor's OWN level<60 raid check let every one of them
// through once the character hit the leveling list's top band.
func TestSourceForRefusesAnOpensTaggedSourceEvenAtLevel60(t *testing.T) {
	idx := lootIndex{1: {{Kind: "raid", Label: "Naxxramas", Opens: "later"}}}
	if _, ok := sourceFor(1, 60, "alliance", "", idx); ok {
		t.Fatal("sourceFor at level 60 with only an Opens:\"later\" raid source: want ok=false")
	}
}

// "raids-1" (Onyxia's Lair, per forever-raid-phases.json's own notes:
// "the first tier opens on 9 December", over a month after launch) is
// refused the same way "later" is - any non-empty Opens value means
// "not open on launch day", not only the literal string "later".
func TestSourceForRefusesRaidsOneTaggedSourceToo(t *testing.T) {
	idx := lootIndex{1: {{Kind: "raid", Label: "Onyxia's Lair", Opens: "raids-1"}}}
	if _, ok := sourceFor(1, 60, "alliance", "", idx); ok {
		t.Fatal("sourceFor with an Opens:\"raids-1\" source: want ok=false")
	}
}

// A non-raid, launch-open source kind (this lane's brief: "dungeons,
// quests, rep, crafted, PvP, world") carries no Opens tag at all and
// must stay obtainable exactly as before - the gate is Opens-based, not
// a blanket new restriction on every source.
func TestSourceForOpensGateDoesNotAffectUntaggedSources(t *testing.T) {
	idx := lootIndex{1: {{Kind: "dungeon", Label: "Blackrock Depths: Some Boss"}}}
	if _, ok := sourceFor(1, 20, "alliance", "", idx); !ok {
		t.Fatal("sourceFor for an untagged dungeon source: want ok=true, unaffected by the Opens gate")
	}
}

// The gate falls back to a later-priority, launch-open source when one
// exists alongside the phase-gated raid drop - an item with both a
// dungeon and a Naxxramas source (rare, but the same shape sourceFor
// already handles for rep/vendor/crafted overlaps) still gets its
// obtainable, non-raid source rather than reporting no source at all.
func TestSourceForFallsBackToANonRaidSourceWhenTheRaidOneIsPhaseGated(t *testing.T) {
	idx := lootIndex{1: {
		{Kind: "raid", Label: "Naxxramas", Opens: "later"},
		{Kind: "crafted", Label: "Blacksmithing"},
	}}
	src, ok := sourceFor(1, 60, "alliance", "", idx)
	if !ok || src.Kind != "crafted" {
		t.Fatalf("sourceFor = %+v, %v, want the crafted source (the raid one is phase-gated)", src, ok)
	}
}

// This lane's brief (bis-ranker-integrity-6), item 4: loot.json's own
// "rep"-kind sources never carry an Opens value at all (lootSource.Opens'
// own doc, data.go), so a reputation faction that is really raid-era
// content - Cenarion Circle (faction 609), earnable only through the
// Ahn'Qiraj War Effort, the same patch this build already gates its
// raid:ahnqiraj source "later" for - would otherwise read as launch-day
// obtainable. loadLootIndex (data.go) closes that gap by applying
// repFactionRaidPhaseOpens[factionID] whenever the pipeline itself is
// silent (firstNonEmpty(src.Opens, repFactionRaidPhaseOpens[factionID])),
// which is already covered by TestLoadLootIndexGatesCenarionCircleRepToLaterPhase
// (data_test.go). What that test does NOT show is that the ranker's own
// gate - sourceObtainable/sourceFor, exactly what band.go's sourceFor
// doc calls "what a launch-day character can get" - actually treats the
// resulting Opens:"later" the same way it treats a raid source: this
// test builds the exact itemSource loadLootIndex produces for
// Earthstrike (item 21180, Cenarion Circle exalted reward, the repro
// this lane's brief names) and confirms sourceFor refuses it even at
// band 60, a Cenarion Circle character could never be exalted before
// launch.
//
// repFactionRaidPhaseOpens (data.go) gates exactly one reputation
// faction id today: 609 (Cenarion Circle).
func TestSourceForRefusesEarthstrikesCenarionCircleRepEvenAtLevel60(t *testing.T) {
	const earthstrikeItemID = 21180
	idx := lootIndex{earthstrikeItemID: {{
		Kind:     "rep",
		Label:    "Cenarion Circle",
		Standing: "exalted",
		Opens:    repFactionRaidPhaseOpens[609],
	}}}
	if _, ok := sourceFor(earthstrikeItemID, 60, "alliance", "", idx); ok {
		t.Fatal("sourceFor(21180 Earthstrike, level 60) with the Cenarion Circle rep source's Opens gate applied: want ok=false")
	}
}

// wowhead-world-drops lane, 2026-09-29: the priority reorder's own
// point - a real dungeon boss now outranks rep/vendor/crafted, so the
// report says "kill this boss" rather than "buy this off a vendor" (or
// "at this reputation") whenever both exist for the same item. See
// band.go's sourceFor doc for why this reverses the 2026-09-28
// night-bis-sources lane's own order.
func TestSourceForPicksABossOverVendorRepAndCrafted(t *testing.T) {
	idx := lootIndex{
		1: {
			{Kind: "crafted", Label: "A Crafted Item"},
			{Kind: "vendor", Label: "A Vendor"},
			{Kind: "rep", Label: "A Reputation"},
			{Kind: "dungeon", Label: "A Dungeon: A Boss"},
		},
	}
	src, ok := sourceFor(1, 30, "alliance", "", idx)
	if !ok || src.Kind != "dungeon" || src.Label != "A Dungeon: A Boss" {
		t.Fatalf("sourceFor = %+v, %v, want the dungeon boss to win over vendor/rep/crafted", src, ok)
	}
}

// Two boss sources naming the same item (a dungeon trash mob padding a
// zone bucket, and the zone's own real boss with a stated chance) pick
// the higher-chance one, not whichever loot.json's array lists first --
// tenet 7's "never show an arbitrary trash mob when a boss ... exists",
// at the tie-break level within the boss tier itself.
func TestSourceForPicksTheHighestChanceBossAmongSeveral(t *testing.T) {
	idx := lootIndex{
		1: {
			{Kind: "dungeon", Label: "The Deadmines: trash", Chance: 0},
			{Kind: "dungeon", Label: "The Deadmines: Mr. Smite", Chance: 20},
		},
	}
	src, ok := sourceFor(1, 30, "alliance", "", idx)
	if !ok || src.Label != "The Deadmines: Mr. Smite" {
		t.Fatalf("sourceFor = %+v, %v, want the higher-chance boss (Mr. Smite)", src, ok)
	}

	// Order must not matter - the higher-chance entry wins regardless of
	// which one the array lists first.
	reversed := lootIndex{
		2: {
			{Kind: "dungeon", Label: "The Deadmines: Mr. Smite", Chance: 20},
			{Kind: "dungeon", Label: "The Deadmines: trash", Chance: 0},
		},
	}
	src, ok = sourceFor(2, 30, "alliance", "", reversed)
	if !ok || src.Label != "The Deadmines: Mr. Smite" {
		t.Fatalf("sourceFor (reversed order) = %+v, %v, want the higher-chance boss (Mr. Smite)", src, ok)
	}
}

// A raid boss with a higher chance than a dungeon boss for the same
// item still loses to the dungeon one below level 60 - the raid<60
// exclusion applies before the chance comparison, not after.
func TestSourceForPicksTheBestObtainableBossAcrossDungeonAndRaid(t *testing.T) {
	idx := lootIndex{
		1: {
			{Kind: "raid", Label: "Molten Core: Ragnaros", Chance: 50},
			{Kind: "dungeon", Label: "The Deadmines: Mr. Smite", Chance: 20},
		},
	}
	src, ok := sourceFor(1, 30, "alliance", "", idx)
	if !ok || src.Kind != "dungeon" {
		t.Fatalf("sourceFor at 30 = %+v, %v, want the dungeon boss (raid excluded below 60)", src, ok)
	}
	src, ok = sourceFor(1, 60, "alliance", "", idx)
	if !ok || src.Kind != "raid" {
		t.Fatalf("sourceFor at 60 = %+v, %v, want the raid boss (higher chance, both obtainable)", src, ok)
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
	pool := buildBandPool(items, idx, "hunter", 20, "horde", map[string]float64{"agility": 2}, 0, false)

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

// TestBuildBandPoolSetsDeadStatCount is this lane's brief
// (bis-ranker-integrity-12), item 5: buildBandPool is the one place
// deadStatCount actually runs (score.go's own function; candidatesBySlot's
// tie-break just reads the field it sets), so this pins that wiring
// directly rather than only through score.go's own unit test.
func TestBuildBandPoolSetsDeadStatCount(t *testing.T) {
	items := []candidate{
		{ID: 1, Name: "All Real Stats", RequiredLevel: 10, EffectiveRequiredLevel: 10, Stats: map[string]float64{"agility": 6}, Slots: []string{"hands"}},
		{ID: 2, Name: "One Dead Stat", RequiredLevel: 10, EffectiveRequiredLevel: 10, Stats: map[string]float64{"agility": 6, "spell_power": 7}, Slots: []string{"hands"}},
	}
	idx := lootIndex{1: {{Kind: "quest", Label: "A Quest"}}, 2: {{Kind: "quest", Label: "A Quest"}}}
	pool := buildBandPool(items, idx, "rogue", 20, "horde", map[string]float64{"agility": 1}, 0, false)
	for _, s := range pool.Scored {
		switch s.ID {
		case 1:
			if s.DeadStatCount != 0 {
				t.Errorf("item 1 DeadStatCount = %d, want 0", s.DeadStatCount)
			}
		case 2:
			if s.DeadStatCount != 1 {
				t.Errorf("item 2 DeadStatCount = %d, want 1 (spell_power is unweighted here)", s.DeadStatCount)
			}
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
	pool := buildBandPool(items, idx, "hunter", 20, "horde", nil, 0, false)
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

// pvp-faction lane, 2026-09-29: loot.json's own pvp:rank-N:alliance/
// pvp:rank-N:horde sources each carry the split's own Faction, which
// data.go's loadLootIndex now copies onto the itemSource's Side (this
// lane's own defect: an Alliance-titled rank reward such as
// Knight-Lieutenant's Pauldrons must never reach a Horde character's
// list, and vice versa).
func TestSourceForGatesAPvpRankItemToItsOwnFaction(t *testing.T) {
	idx := lootIndex{
		16338: {{Kind: "pvp", Label: "Rank 11 (Alliance)", Side: "alliance", Rank: 11}},
	}
	if _, ok := sourceFor(16338, 60, "horde", "", idx); ok {
		t.Fatal("an Alliance-only pvp rank reward was offered to a horde character")
	}
	src, ok := sourceFor(16338, 60, "alliance", "", idx)
	if !ok || src.Kind != "pvp" || src.Rank != 11 {
		t.Fatalf("alliance source = %+v, %v; want the pvp rank-11 source", src, ok)
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

// This lane's brief, defect 2: sourceKindPriority now lists rep ahead
// of vendor, so when an item's index carries both (the quartermaster's
// own vendor row, gated by vendorInheritsRepStandingGate in data.go),
// the published label reads the reputation the player actually has to
// earn - "Silverwing Sentinels (honored)" - rather than the generic
// "vendor" the old order preferred.
func TestSourceForPrefersRepOverVendorForTheSameItem(t *testing.T) {
	idx := lootIndex{
		20444: {
			{Kind: "rep", Label: "Silverwing Sentinels", Side: "alliance", Standing: "honored"},
			{Kind: "vendor", Label: "Illiyana Moonblaze", Side: "alliance", Standing: "honored"},
		},
	}
	src, ok := sourceFor(20444, 20, "alliance", "", idx)
	if !ok || src.Kind != "rep" || src.Label != "Silverwing Sentinels" {
		t.Fatalf("sourceFor = %+v, %v, want the rep source (Silverwing Sentinels), not vendor", src, ok)
	}
}

// The defect itself: a vendor row gated to the same standing as its
// matching rep row must obey that gate exactly like the rep row does -
// Outrunner's Bow (Warsong Outriders revered) must not resolve as an
// ordinary level-agnostic vendor purchase at level 20.
func TestSourceForGatesAVendorRowThatInheritedARepStanding(t *testing.T) {
	idx := lootIndex{
		19562: {
			{Kind: "rep", Label: "Warsong Outriders", Side: "horde", Standing: "revered"},
			{Kind: "vendor", Label: "Kelm Hargunth", Side: "horde", Standing: "revered"},
		},
	}
	if _, ok := sourceFor(19562, 20, "horde", "", idx); ok {
		t.Fatal("a vendor row gated to revered must not resolve for a level-20 character (bypassing the reputation gate)")
	}
	src, ok := sourceFor(19562, 60, "horde", "", idx)
	if !ok || src.Kind != "rep" {
		t.Fatalf("sourceFor at 60 = %+v, %v, want the rep source obtainable at 60", src, ok)
	}
}

// A vendor source with no matching rep source in the index (an
// ordinary gold vendor) must remain ungated - vendorInheritsRepStandingGate
// only touches a vendor row that shares an id with a rep row.
func TestSourceForOrdinaryVendorWithNoMatchingRepStaysUngated(t *testing.T) {
	idx := lootIndex{1: {{Kind: "vendor", Label: "A Gold Vendor"}}}
	src, ok := sourceFor(1, 1, "alliance", "", idx)
	if !ok || src.Kind != "vendor" {
		t.Fatalf("sourceFor = %+v, %v, want the ungated vendor source obtainable at any level", src, ok)
	}
}

// TestBuildBandPoolGatesEveryLegendaryRegardlessOfSourceKind is this
// lane's brief, item 2: a quality-5 item is gated "later" no matter
// which loot.json source kind names it - unlike the old
// raidLockedQuestOpens hand list (data.go), which needed a manual
// quest-id entry per legendary, this is one rule. Four different
// source kinds are asserted here (quest - Atiesh's own real shape;
// raid, which was already gated by Opens before this lane and must
// stay gated; and vendor/crafted, which loot.json does not carry a
// legendary under today but which the rule must still catch, since a
// per-kind allowlist would silently miss one) to prove the exclusion
// is keyed on Quality, not on Kind.
func TestBuildBandPoolGatesEveryLegendaryRegardlessOfSourceKind(t *testing.T) {
	items := []candidate{
		{ID: 22589, Name: "Atiesh, Greatstaff of the Guardian", Quality: legendaryQuality, RequiredLevel: 60, EffectiveRequiredLevel: 60, Slots: []string{"main_hand"}, Stats: map[string]float64{"spell_power": 100}},
		{ID: 19019, Name: "Thunderfury, Blessed Blade of the Windseeker", Quality: legendaryQuality, RequiredLevel: 60, EffectiveRequiredLevel: 60, Slots: []string{"main_hand"}, Stats: map[string]float64{"agility": 100}},
		{ID: 99001, Name: "Hypothetical Legendary Trinket", Quality: legendaryQuality, RequiredLevel: 60, EffectiveRequiredLevel: 60, Slots: []string{"trinket1", "trinket2"}},
		{ID: 810, Name: "Hammer of the Northern Wind", Quality: 4, RequiredLevel: 49, EffectiveRequiredLevel: 49, Slots: []string{"main_hand"}, Stats: map[string]float64{"agility": 10}},
	}
	idx := lootIndex{
		22589: {{Kind: "quest", Label: "Atiesh, Greatstaff of the Guardian"}},
		19019: {{Kind: "raid", Label: "Molten Core: Ragnaros", Opens: "later"}},
		99001: {{Kind: "vendor", Label: "A Vendor Nobody Should Ever Reach"}},
		810:   {{Kind: "world_drop", Label: "World drop"}},
	}
	weights := map[string]float64{"spell_power": 1, "agility": 1}
	pool := buildBandPool(items, idx, "warrior", 60, "alliance", weights, 0, false)

	scoredIDs := make(map[int]bool, len(pool.Scored))
	for _, s := range pool.Scored {
		scoredIDs[s.ID] = true
	}
	for _, legendaryID := range []int{22589, 19019, 99001} {
		if scoredIDs[legendaryID] {
			t.Errorf("pool.Scored contains legendary item %d, want it excluded regardless of source kind", legendaryID)
		}
	}
	if !scoredIDs[810] {
		t.Error("pool.Scored is missing the ordinary (non-legendary) weapon 810 - the rule must not over-exclude")
	}

	var foundLegendaryInNoSource int
	for _, c := range pool.NoSource {
		if c.Quality == legendaryQuality {
			foundLegendaryInNoSource++
		}
	}
	if foundLegendaryInNoSource != 3 {
		t.Errorf("pool.NoSource carries %d legendaries, want 3 (accounted for, same as a raid-gated item)", foundLegendaryInNoSource)
	}
}

// TestPvpRankExceedsCap pins band.go's own rule (this lane's brief,
// item 3): Rank above pvpRankCap is capped regardless of Kind - a
// vendor row vendorInheritsPvpRankGate (data.go) copied a pvp source's
// Rank onto must be capped exactly like the pvp source itself
// (Captain O'Neal's own vendor row for Grand Marshal's Stave, Rank 18 -
// the real bug this lane's own dogfood run found: "vendor" outranks
// "pvp" in sourceKindPriority, so sourceFor always picked the
// Rank-less vendor row over the pvp one, and a Kind == "pvp" check
// here would have silently done nothing for every one of them).
func TestPvpRankExceedsCap(t *testing.T) {
	cases := []struct {
		name string
		src  itemSource
		want bool
	}{
		{"pvp rank 10 (the cap itself) does not exceed it", itemSource{Kind: "pvp", Rank: 10}, false},
		{"pvp rank 11 exceeds it", itemSource{Kind: "pvp", Rank: 11}, true},
		{"pvp rank 18 exceeds it", itemSource{Kind: "pvp", Rank: 18}, true},
		{"a non-pvp source with no inherited rank is never capped", itemSource{Kind: "quest", Rank: 0}, false},
		{"a vendor row that inherited a pvp rank IS capped, same as the pvp source itself", itemSource{Kind: "vendor", Rank: 18}, true},
	}
	for _, tc := range cases {
		if got := pvpRankExceedsCap(tc.src); got != tc.want {
			t.Errorf("%s: pvpRankExceedsCap(%+v) = %v, want %v", tc.name, tc.src, got, tc.want)
		}
	}
}

// TestPvpSourceLabel is the fourth wow-player sweep's own item 2: never
// the bare bucket name or the bare quartermaster name - "PvP rank N ·
// Title · Faction", the same wording web/src/lib/bis/copy.ts's own
// pvpSourceLabel produces from the identical rank/title/faction facts.
func TestPvpSourceLabel(t *testing.T) {
	cases := []struct {
		name           string
		rank           int
		title, faction string
		want           string
	}{
		{"a known rank/title", 9, "Master Sergeant", "alliance", "PvP rank 9 · Master Sergeant · Alliance"},
		{"a horde rank/title", 18, "High Warlord", "horde", "PvP rank 18 · High Warlord · Horde"},
		{"no title (rank outside the known ladder) falls back to rank and faction", 4, "", "alliance", "PvP rank 4 · Alliance"},
	}
	for _, tc := range cases {
		if got := pvpSourceLabel(tc.rank, tc.title, tc.faction); got != tc.want {
			t.Errorf("%s: pvpSourceLabel(%d, %q, %q) = %q, want %q", tc.name, tc.rank, tc.title, tc.faction, got, tc.want)
		}
	}
}
