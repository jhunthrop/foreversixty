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
	got := bestSetPieces(bySlot)
	if len(got) != 1 {
		t.Fatalf("bestSetPieces = %+v, want exactly one implemented set (41)", got)
	}
	pieces, ok := got[41]
	if !ok || len(pieces) != 1 || pieces[0].slot != "head" {
		t.Fatalf("bestSetPieces[41] = %+v, want the one head candidate", pieces)
	}
}

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
