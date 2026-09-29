package main

import (
	"errors"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func trinket(id int, name string, itemLevel int) scored {
	return scored{candidate: candidate{ID: id, Name: name, ItemLevel: itemLevel, Slots: []string{"trinket1", "trinket2"}}}
}

func TestTopByItemLevelOrdersHighestFirstAndBoundsToTopN(t *testing.T) {
	list := []scored{
		trinket(1, "A", 10),
		trinket(2, "B", 30),
		trinket(3, "C", 20),
		trinket(4, "D", 30), // ties item level 30 with B, breaks on id ascending
		trinket(5, "E", 5),
		trinket(6, "F", 40),
	}
	got := topByItemLevel(list, 0, "")
	if len(got) != trinketTopN {
		t.Fatalf("topByItemLevel returned %d, want trinketTopN (%d)", len(got), trinketTopN)
	}
	wantOrder := []int{6, 2, 4, 3, 1} // 40, 30(id2), 30(id4), 20, 10 - item 5 (ilvl 5) is bumped
	for i, id := range wantOrder {
		if got[i].ID != id {
			t.Fatalf("topByItemLevel[%d].ID = %d, want %d (full: %+v)", i, got[i].ID, id, got)
		}
	}
}

// trinketScored is trinket() plus an explicit Score, for the score-axis
// half of trinketShortlist's own test.
func trinketScored(id int, name string, itemLevel int, sc float64) scored {
	s := trinket(id, name, itemLevel)
	s.Score = sc
	return s
}

// This lane's brief (bis-ranker-integrity, 2026-09-29): a candidate
// that is far ahead on SCORE but far behind on item level must still
// enter the real-sim pool, not just the item-level winners -
// Neltharion's Tear (mage-fire's own dogfood case) is exactly this
// shape: low item level, by far the best score() in the pool.
func TestTrinketShortlistUnionsTopByItemLevelAndTopByScore(t *testing.T) {
	// list is already score-sorted, matching candidatesBySlot's own
	// contract (trinketShortlist's doc): item 100 is the best SCORE by
	// far, but its item level (10) would never make topByItemLevel's own
	// top trinketTopN (5) against items 2-6, all higher item level.
	list := []scored{
		trinketScored(100, "Best Score, Low ItemLevel", 10, 999),
		trinketScored(2, "ItemLevel 90", 90, 5),
		trinketScored(3, "ItemLevel 85", 85, 4),
		trinketScored(4, "ItemLevel 83", 83, 3),
		trinketScored(5, "ItemLevel 80", 80, 2),
		trinketScored(6, "ItemLevel 78", 78, 1),
		trinketScored(7, "ItemLevel 70", 70, 0),
	}
	got := trinketShortlist(list, 0, "")
	found := false
	for _, c := range got {
		if c.ID == 100 {
			found = true
		}
	}
	if !found {
		t.Fatalf("trinketShortlist = %+v, want item 100 (best score, never top-item-level) included", got)
	}
	// Every item-level-top-trinketTopN candidate is still present too -
	// this is a union, not a replacement.
	for _, id := range []int{2, 3, 4, 5, 6} {
		hasIt := false
		for _, c := range got {
			if c.ID == id {
				hasIt = true
			}
		}
		if !hasIt {
			t.Errorf("trinketShortlist = %+v, want item %d (top-item-level) still included", got, id)
		}
	}
	// item 7 (item level 70, worst score) makes neither top-N: excluded.
	for _, c := range got {
		if c.ID == 7 {
			t.Errorf("trinketShortlist = %+v, want item 7 excluded (bottom of both axes)", got)
		}
	}
}

func TestTrinketShortlistExcludesPairMateByIDAndName(t *testing.T) {
	list := []scored{trinket(1, "Same Name", 10), trinket(2, "Same Name", 20), trinket(3, "Other", 5)}
	got := trinketShortlist(list, 1, "")
	for _, c := range got {
		if c.ID == 1 {
			t.Fatalf("trinketShortlist still carries excluded id 1: %+v", got)
		}
	}
	got2 := trinketShortlist(list, 0, "Same Name")
	for _, c := range got2 {
		if c.Name == "Same Name" {
			t.Fatalf("trinketShortlist still carries an item sharing the excluded name: %+v", got2)
		}
	}
}

func TestTopByItemLevelExcludesPairMateByIDAndName(t *testing.T) {
	list := []scored{trinket(1, "Same Name", 10), trinket(2, "Same Name", 20), trinket(3, "Other", 5)}
	got := topByItemLevel(list, 1, "")
	for _, c := range got {
		if c.ID == 1 {
			t.Fatalf("topByItemLevel still carries excluded id 1: %+v", got)
		}
	}
	got2 := topByItemLevel(list, 0, "Same Name")
	for _, c := range got2 {
		if c.Name == "Same Name" {
			t.Fatalf("topByItemLevel still carries an item sharing the excluded name: %+v", got2)
		}
	}
}

func TestFormatTrinketRankError(t *testing.T) {
	testErr := errors.New("engine exploded")
	msg := formatTrinketRankError("trinket1", trinket(9, "Bad Trinket", 10), testErr)
	if !strings.Contains(msg, "trinket1") || !strings.Contains(msg, "Bad Trinket") || !strings.Contains(msg, testErr.Error()) {
		t.Errorf("formatTrinketRankError = %q, missing an expected substring", msg)
	}
}

func TestRankTrinketSlotPicksTheHighestMeasuredDPS(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder"}}},
	}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Low DPS Trinket", 30), trinket(3, "High DPS Trinket", 20)},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}): 100,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 3}}): 200,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 3 {
		t.Fatalf("trinket1 pick = %+v, want item 3 (higher measured dps)", out["trinket1"].Item)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != 2 {
		t.Fatalf("trinket1 runner-up = %+v, want item 2", out["trinket1"].RunnerUp)
	}
	// picks is untouched (immutability rule): a fresh map came back.
	if picks["trinket1"].Item.ID != 1 {
		t.Fatalf("the input picks map was mutated: %+v", picks["trinket1"])
	}
}

func TestRankTrinketSlotSkipsAFailingCandidateAndKeepsGoing(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Fails To Sim", 30), trinket(3, "Sims Fine", 20)},
	}
	fake := &fakeEngine{
		FailGear:  gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}),
		DPSByGear: map[string]float64{gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 3}}): 50},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1")
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1 (item 2 failed)", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 3 {
		t.Fatalf("trinket1 pick = %+v, want item 3 (the only one that sims)", out["trinket1"].Item)
	}
	if out["trinket1"].RunnerUp != nil {
		t.Fatalf("trinket1 runner-up = %+v, want nil: only one candidate sim succeeded", out["trinket1"].RunnerUp)
	}
}

func TestRankTrinketSlotAllCandidatesFailLeavesPicksUnchanged(t *testing.T) {
	original := &scored{candidate: candidate{ID: 1, Name: "Original Pick"}}
	picks := map[string]slotPick{"trinket1": {Item: original}}
	bySlot := map[string][]scored{"trinket1": {trinket(2, "Fails", 30)}}
	fake := &fakeEngine{FailGear: gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}})}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1")
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1", notes)
	}
	if out["trinket1"].Item != original {
		t.Fatalf("trinket1 = %+v, want the original pick unchanged", out["trinket1"].Item)
	}
}

func TestRankTrinketSlotNoCandidatesReturnsUnchanged(t *testing.T) {
	original := &scored{candidate: candidate{ID: 1, Name: "Original Pick"}}
	picks := map[string]slotPick{"trinket1": {Item: original}}
	fake := &fakeEngine{}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, map[string][]scored{}, "trinket1")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["trinket1"].Item != original {
		t.Fatalf("trinket1 = %+v, want unchanged", out["trinket1"].Item)
	}
}
