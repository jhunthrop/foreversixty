package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func setItem(id int, name string, setID int, slots ...string) scored {
	sid := setID
	return scored{candidate: candidate{ID: id, Name: name, SetID: &sid, Slots: slots}}
}

func TestBestSetPiecesKeepsOnlyEngineImplementedSetsAndSkipsTrinkets(t *testing.T) {
	// 41 is a real entry in setids_generated.go (this lane's own
	// generated table); 999999 is not a real set id at all.
	bySlot := map[string][]scored{
		"head":      {setItem(1, "Implemented Set Head", 41, "head")},
		"chest":     {setItem(2, "Unimplemented Set Chest", 999999, "chest")},
		"trinket1":  {setItem(3, "Set Trinket", 41, "trinket1")}, // excluded: trinkets ranked separately
		"main_hand": {effectItem(4, "No Set Weapon", "", "main_hand")},
	}
	got := bestSetPieces(bySlot, "")
	if len(got) != 1 {
		t.Fatalf("bestSetPieces = %+v, want exactly one implemented set (41)", got)
	}
	pieces, ok := got[41]
	if !ok || len(pieces) != 1 || pieces[0].slot != "head" {
		t.Fatalf("bestSetPieces[41] = %+v, want the one head candidate", pieces)
	}
}

// The real bug this test guards: finger1 and finger2 fan out from the
// exact same candidate list (data.go's plannerSlots - a ring's Slots
// is always both), so their own #1-scored item is the identical
// physical ring in both. Before this dedup, bestSetPieces counted
// that one ring as TWO of its own set's pieces, and
// trySetCompletion's trial equipped it in finger1 AND finger2 at
// once - the literal "same trinket twice" shape (there, a ring
// instead) this lane's audit found by eye elsewhere in tonight's
// output (applySwaps' own new test, verify_test.go).
func TestBestSetPiecesNeverCountsTheSamePhysicalRingTwice(t *testing.T) {
	ring := setItem(50, "Set Ring", 41, "finger1", "finger2")
	bySlot := map[string][]scored{
		"finger1": {ring},
		"finger2": {ring},
		"chest":   {setItem(51, "Set Chest", 41, "chest")},
	}
	got := bestSetPieces(bySlot, "")
	pieces, ok := got[41]
	if !ok {
		t.Fatalf("bestSetPieces = %+v, want set 41", got)
	}
	seenRing := 0
	for _, p := range pieces {
		if p.item.ID == ring.ID {
			seenRing++
		}
	}
	if seenRing != 1 {
		t.Fatalf("set 41's pieces = %+v, want the ring counted exactly once (finger1 and finger2 share one candidate list)", pieces)
	}
	if len(pieces) != 2 {
		t.Fatalf("set 41's pieces = %+v, want the ring once plus the chest piece (2 total)", pieces)
	}
}

// Mirrors rank.go's own dual-wield defense: a two-hander must never be
// offered as a dual-wield spec's main-hand set piece either, or
// trySetCompletion could equip it alongside an off-hand item from a
// completely different, independently-scored trial.
func TestBestSetPiecesExcludesTwoHandMainHandForADualWieldSpec(t *testing.T) {
	twoHander := scored{candidate: candidate{ID: 60, Name: "Two-Hand Set Sword", SetID: intPtr(41), Slots: []string{"main_hand"}, TwoHand: true}}
	bySlot := map[string][]scored{
		"main_hand": {twoHander},
	}
	got := bestSetPieces(bySlot, "hunter-survival")
	if len(got) != 0 {
		t.Fatalf("bestSetPieces = %+v, want no set pieces (the only main_hand candidate is a two-hander, excluded for a dual-wield spec)", got)
	}
}

func intPtr(n int) *int { return &n }

func TestTrySetCompletionAdoptsTheSetWhenItVerifiesHigher(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Score-based Head"}}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Score-based Chest"}}},
	}
	bySlot := map[string][]scored{
		"head":  {setItem(11, "Set Head", 41, "head")},
		"chest": {setItem(21, "Set Chest", 41, "chest")},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "head", ItemID: 10}, {Slot: "chest", ItemID: 20}}): 100, // baseline (independently-scored picks)
			gearKey([]api.GearSlot{{Slot: "head", ItemID: 11}, {Slot: "chest", ItemID: 21}}): 150, // the completed set - higher
		},
	}
	out, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, picks, bySlot)
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1 (the adoption note)", notes)
	}
	if out["head"].Item == nil || out["head"].Item.ID != 11 {
		t.Fatalf("head = %+v, want the set piece (11)", out["head"].Item)
	}
	if out["chest"].Item == nil || out["chest"].Item.ID != 21 {
		t.Fatalf("chest = %+v, want the set piece (21)", out["chest"].Item)
	}
}

func TestTrySetCompletionKeepsTheIndependentlyScoredPicksWhenTheSetDoesNotVerifyHigher(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Score-based Head"}}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Score-based Chest"}}},
	}
	bySlot := map[string][]scored{
		"head":  {setItem(11, "Set Head", 41, "head")},
		"chest": {setItem(21, "Set Chest", 41, "chest")},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "head", ItemID: 10}, {Slot: "chest", ItemID: 20}}): 200, // baseline wins
			gearKey([]api.GearSlot{{Slot: "head", ItemID: 11}, {Slot: "chest", ItemID: 21}}): 150,
		},
	}
	out, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, picks, bySlot)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none (the set did not win)", notes)
	}
	if out["head"].Item.ID != 10 || out["chest"].Item.ID != 20 {
		t.Fatalf("picks changed despite the set losing: %+v / %+v", out["head"].Item, out["chest"].Item)
	}
}

func TestTrySetCompletionSkipsASetAlreadyFullyEquipped(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 11, Name: "Set Head"}}},
		"chest": {Item: &scored{candidate: candidate{ID: 21, Name: "Set Chest"}}},
	}
	bySlot := map[string][]scored{
		"head":  {setItem(11, "Set Head", 41, "head")},
		"chest": {setItem(21, "Set Chest", 41, "chest")},
	}
	fake := &fakeEngine{}
	_, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, picks, bySlot)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if len(fake.Calls) != 0 {
		t.Fatalf("fake engine was called %d times, want 0 (already fully equipped, nothing to try)", len(fake.Calls))
	}
}

func TestTrySetCompletionNoImplementedSetsReturnsUnchanged(t *testing.T) {
	original := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Plain Head"}}}}
	fake := &fakeEngine{}
	out, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, original, map[string][]scored{})
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["head"].Item.ID != 1 {
		t.Fatalf("out = %+v, want unchanged", out)
	}
	if len(fake.Calls) != 0 {
		t.Fatalf("fake engine was called %d times, want 0", len(fake.Calls))
	}
}
