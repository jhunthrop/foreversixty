package main

import (
	"fmt"
	"testing"
)

func src(kind, label string) itemSource { return itemSource{Kind: kind, Label: label} }

// fixtureItem is a pool candidate with its sources by filter kind; the
// first source is the published one.
func fixtureItem(id int, name string, slots []string, score float64, sources ...itemSource) scored {
	kinds := make(map[string]itemSource, len(sources))
	for _, s := range sources {
		k, _ := filterKindOfSource(s.Kind)
		kinds[k] = s
	}
	return scored{
		candidate:   candidate{ID: id, Name: name, Slots: slots},
		Score:       score,
		Source:      sources[0],
		HasSource:   true,
		KindSources: kinds,
	}
}

func weapon(s scored, twoHand, mainHandOnly bool) scored {
	s.ClassID = itemClassWeapon
	s.TwoHand = twoHand
	s.MainHandOnly = mainHandOnly
	return s
}

func limitCategory(s scored, category int) scored {
	s.LimitCategory = category
	return s
}

var (
	fixtureHead    = []string{"head"}
	fixtureFinger  = []string{"finger1", "finger2"}
	fixtureTrinket = []string{"trinket1", "trinket2"}
	fixtureMain    = []string{"main_hand"}
	fixtureOff     = []string{"off_hand"}
)

// sourceFilterFixturePool is a band's pool with every pairing hazard: a
// quest reward a dungeon also drops, a rep item at Revered, two rings in
// one limit category, two rings sharing a name, a two-hander atop
// one-handers and shields, and an ineligible zero-score head.
func sourceFilterFixturePool(dualWield bool) []scored {
	pool := []scored{
		fixtureItem(101, "Dungeon Helm", fixtureHead, 10, src("dungeon", "Deadmines: Smite")),
		fixtureItem(102, "Quest Helm", fixtureHead, 8, src("quest", "Helm Quest"), src("dungeon", "Wailing Caverns: Mutanus")),
		fixtureItem(103, "Crafted Helm", fixtureHead, 7, src("crafted", "Blacksmithing")),
		fixtureItem(104, "Rep Helm", fixtureHead, 5, itemSource{Kind: "rep", Label: "Stormwind", Standing: "honored"}),
		fixtureItem(105, "Dead Helm", fixtureHead, 0, src("world", "Some Mob")),
		fixtureItem(201, "Ring A", fixtureFinger, 10, src("dungeon", "Scarlet Monastery")),
		fixtureItem(202, "Quest Ring", fixtureFinger, 9, src("quest", "Ring Quest")),
		fixtureItem(203, "Dual Ring", fixtureFinger, 8, src("dungeon", "Razorfen Kraul"), src("quest", "Other Quest")),
		limitCategory(fixtureItem(204, "Seal One", fixtureFinger, 7, src("crafted", "Jewelcrafting")), 50),
		limitCategory(fixtureItem(205, "Seal Two", fixtureFinger, 6, src("pvp", "PvP reward")), 50),
		fixtureItem(206, "Ring A", fixtureFinger, 5, src("world_drop", "World drop")),
		fixtureItem(301, "Raid Trinket", fixtureTrinket, 10, src("raid", "Molten Core")),
		fixtureItem(302, "Rep Trinket", fixtureTrinket, 9, itemSource{Kind: "rep", Label: "Argent Dawn", Standing: "revered"}),
		fixtureItem(303, "Small Raid Trinket", fixtureTrinket, 4, src("raid", "Onyxia's Lair")),
		fixtureItem(304, "Vendor Trinket", fixtureTrinket, 3, src("vendor", "Some Vendor")),
	}
	if dualWield {
		return append(pool,
			weapon(fixtureItem(401, "Big Axe", fixtureMain, 15, src("dungeon", "Uldaman")), true, false),
			weapon(fixtureItem(402, "Dungeon Sword", fixtureMain, 10, src("dungeon", "Gnomeregan")), false, false),
			weapon(fixtureItem(403, "Vendor Sword", fixtureMain, 9, src("vendor", "Weaponsmith")), false, false),
			weapon(fixtureItem(404, "Main Only", fixtureMain, 8, src("crafted", "Blacksmithing")), false, true),
			weapon(fixtureItem(405, "Rep Dagger", fixtureOff, 6, src("rep", "Darnassus")), false, false),
		)
	}
	return append(pool,
		weapon(fixtureItem(401, "Big Axe", fixtureMain, 20, src("dungeon", "Uldaman")), true, false),
		weapon(fixtureItem(402, "Quest Sword", fixtureMain, 12, src("quest", "Sword Quest")), false, false),
		weapon(fixtureItem(403, "Crafted Sword", fixtureMain, 9, src("crafted", "Blacksmithing")), false, false),
		fixtureItem(501, "Vendor Shield", fixtureOff, 5, src("vendor", "Armorer")),
		fixtureItem(502, "Dungeon Shield", fixtureOff, 4, src("dungeon", "Stockades")),
	)
}

// fixtureRows is buildReport's skeleton for picks: the pick, its source.
func fixtureRows(picks map[string]slotPick) []slotRow {
	rows := make([]slotRow, 0, len(slotOrder))
	for _, slot := range slotOrder {
		row := slotRow{Slot: slot}
		if pk := picks[slot]; pk.Item != nil {
			row.ItemID, row.ItemName = pk.Item.ID, pk.Item.Name
			row.Source, row.SourceKind = pk.Item.Source.Label, pk.Item.Source.Kind
		}
		rows = append(rows, row)
	}
	return rows
}

// bruteFilteredSlots is the definition the published fields must
// reproduce: repickFiltered over the whole pool, kinds only.
func bruteFilteredSlots(spec string, picks map[string]slotPick, pickBySlot map[string][]scored, allowed map[string]bool) map[string]int {
	keep := func(c scored) bool { return anyAllowed(orderedFilterKinds(filterKindSources(c)), allowed) }
	set := repickFiltered(spec, picks, pickBySlot, keep)
	out := make(map[string]int, len(set))
	for slot, pk := range set {
		if id := itemIDOf(pk); id != 0 {
			out[slot] = id
		}
	}
	return out
}

// allKindSubsets is every combination of filter kinds, on or off.
func allKindSubsets() []map[string]bool {
	n := len(filterKindOrder)
	out := make([]map[string]bool, 0, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		allowed := make(map[string]bool, n)
		for i, k := range filterKindOrder {
			if mask&(1<<i) != 0 {
				allowed[k] = true
			}
		}
		out = append(out, allowed)
	}
	return out
}

func checkFilteredPicksMatchBruteForce(t *testing.T, spec string, picks map[string]slotPick, bySlot map[string][]scored) {
	t.Helper()
	rows := withSourceKindPicks(fixtureRows(picks), sourceFilterInputs{Spec: spec, Picks: picks, PickBySlot: bySlot, ReferenceDPSPerPoint: 0.5})
	for _, allowed := range allKindSubsets() {
		want := bruteFilteredSlots(spec, picks, bySlot, allowed)
		got := resolveFilteredSlots(rows, allowed)
		for _, slot := range slotOrder {
			gotID := 0
			if got[slot] != nil {
				gotID = got[slot].ItemID
			}
			if gotID != want[slot] {
				t.Errorf("%s %v: %s resolved to %d, want %d (the top-ranked candidate of the allowed kinds)", spec, fmt.Sprint(allowed), slot, gotID, want[slot])
			}
		}
	}
}

// Owner MUST 2026-11-05: for every slot and every combination of kinds,
// the published pick plus by_kind give exactly the top-ranked candidate
// of the allowed kinds, pairs and two-handers included.
func TestFilteredPickIsTheTopRankedCandidateOfTheAllowedKinds(t *testing.T) {
	for _, spec := range []string{"warrior-arms", "rogue-combat"} {
		bySlot := candidatesBySlot(sourceFilterFixturePool(spec == "rogue-combat"))
		checkFilteredPicksMatchBruteForce(t, spec, pick(spec, bySlot), bySlot)
	}
}

// A pick a sim promoted over the score order (a swap, a tournament) still
// comes first: rank order is the ranker's own decision, then the pool.
func TestFilteredPickKeepsASimPromotedPickFirst(t *testing.T) {
	bySlot := candidatesBySlot(sourceFilterFixturePool(false))
	picks := pick("warrior-arms", bySlot)
	crafted := bySlot["head"][2]
	picks["head"] = slotPick{Item: &crafted}
	ring := bySlot["finger2"][4] // Seal Two, pvp, ranked under Dual Ring
	picks["finger2"] = slotPick{Item: &ring}
	checkFilteredPicksMatchBruteForce(t, "warrior-arms", picks, bySlot)
}

func TestByKindShape(t *testing.T) {
	bySlot := candidatesBySlot(sourceFilterFixturePool(false))
	picks := pick("warrior-arms", bySlot)
	rows := withSourceKindPicks(fixtureRows(picks), sourceFilterInputs{Spec: "warrior-arms", Picks: picks, PickBySlot: bySlot, ReferenceDPSPerPoint: 0.5})
	bySlotRow := make(map[string]slotRow, len(rows))
	for _, r := range rows {
		bySlotRow[r.Slot] = r
	}

	head := bySlotRow["head"]
	if _, ok := head.ByKind["dungeon"]; ok {
		t.Error("head by_kind repeats the pick's own kind (dungeon)")
	}
	quest := head.ByKind["quest"]
	if quest.ItemID != 102 || quest.Rank != 2 || quest.SourceKind != "" || quest.Source != "Helm Quest" {
		t.Errorf("head by_kind.quest = %+v, want 102 at rank 2 from Helm Quest, source_kind left to the key", quest)
	}
	if quest.DPSDelta == nil || *quest.DPSDelta != -1 {
		t.Errorf("head by_kind.quest dps_delta = %v, want -1 ((8-10) x 0.5)", quest.DPSDelta)
	}
	if rep := head.ByKind["rep"]; rep.Standing != "honored" {
		t.Errorf("head by_kind.rep standing = %q, want honored", rep.Standing)
	}
	if _, ok := head.ByKind["world"]; ok {
		t.Error("a zero-score head stood in for the pick")
	}

	f2 := bySlotRow["finger2"]
	if f2.ItemID != 202 {
		t.Fatalf("finger2 pick = %d, want 202", f2.ItemID)
	}
	if _, ok := f2.ByKind["quest"]; !ok {
		t.Error("the second ring's by_kind dropped its pick's own kind; a moved first ring can block the pick")
	}
	dungeon := f2.ByKind["dungeon"]
	// finger1 can resolve to 201 (its pick) or 206 (its world entry): both
	// are named Ring A, so either blocks either.
	if dungeon.ItemID != 201 || fmt.Sprint(dungeon.NotWith) != "[201 206]" || dungeon.Next == nil || dungeon.Next.ItemID != 203 {
		t.Errorf("finger2 by_kind.dungeon = %+v, want Ring A not with either Ring A, next Dual Ring", dungeon)
	}
	if seal := f2.ByKind["pvp"]; fmt.Sprint(seal.NotWith) != "[204 205]" {
		t.Errorf("finger2 by_kind.pvp = %+v, want Seal Two not with either seal (limit category 50)", seal)
	}

	if main := bySlotRow["main_hand"]; !main.TwoHand {
		t.Error("the two-handed main hand pick is not marked two_hand")
	}
	off := bySlotRow["off_hand"]
	if off.ItemID != 0 || off.ByKind["vendor"].ItemID != 501 {
		t.Errorf("off_hand under a two-hander = %+v, want no pick and the shields by kind", off)
	}
	if off.ByKind["vendor"].DPSDelta != nil {
		t.Error("an off hand with no pick published a DPS gap against nothing")
	}
}
