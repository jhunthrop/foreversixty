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

// trinketWithEffect is trinket() plus an EffectText, for the hybrid
// sweep's own trinket-margin tests below. id must be a real engine-
// implemented-effect id (effectids_generated.go) for a "modelled"
// trinket, or any id NOT in that map (e.g. one obviously fake, like
// 999999) for an "unmodelled" one - trinketEffectUnmodelled (trinkets.go)
// reads exactly that distinction.
func trinketWithEffect(id int, name string, itemLevel int, effectText string) scored {
	s := trinket(id, name, itemLevel)
	s.EffectText = effectText
	return s
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
	// This lane's brief, item 7: the winner and runner-up each carry
	// their OWN measured DPS from this tournament, not each other's and
	// not zero - report.go's buildReport reads this to publish sim_dps
	// instead of score()'s stat estimate.
	if out["trinket1"].Item.MeasuredDPS != 200 {
		t.Errorf("trinket1 pick MeasuredDPS = %v, want 200", out["trinket1"].Item.MeasuredDPS)
	}
	if out["trinket1"].RunnerUp.MeasuredDPS != 100 {
		t.Errorf("trinket1 runner-up MeasuredDPS = %v, want 100", out["trinket1"].RunnerUp.MeasuredDPS)
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

// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 2:
// rankTrinketSlot's own no-trinket baseline sim (swapSlot's itemID-0
// shape) is what lets report.go's zero-value gate tell a genuinely
// worthless trinket apart from a real one - MeasuredDPS alone (the
// whole SET's own absolute DPS) can never do this, since it is always
// positive regardless of whether the trinket itself contributes
// anything at all.
func TestRankTrinketSlotComputesGainAgainstANoTrinketBaseline(t *testing.T) {
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
			// The baseline call swaps the slot's own item id to 0, which
			// swapSlot's own doc says drops the slot from the gear list
			// entirely - an empty gear key.
			gearKey(nil): 190,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if !out["trinket1"].Item.GainMeasured {
		t.Fatalf("trinket1 pick GainMeasured = false, want true")
	}
	if out["trinket1"].Item.MeasuredGainDPS != 10 {
		t.Errorf("trinket1 pick MeasuredGainDPS = %v, want 10 (200 measured - 190 baseline)", out["trinket1"].Item.MeasuredGainDPS)
	}
	if !out["trinket1"].RunnerUp.GainMeasured || out["trinket1"].RunnerUp.MeasuredGainDPS != -90 {
		t.Errorf("trinket1 runner-up gain = measured=%v dps=%v, want measured=true dps=-90 (100 - 190)", out["trinket1"].RunnerUp.GainMeasured, out["trinket1"].RunnerUp.MeasuredGainDPS)
	}
}

// A baseline sim failure costs the gain check, not the ranking itself -
// tenet 8: never claim a gain this command could not actually measure.
func TestRankTrinketSlotLeavesGainUnmeasuredWhenTheBaselineSimFails(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Only Candidate", 30)},
	}
	fake := &fakeEngine{
		// The no-trinket baseline call's own gear list is empty
		// (swapSlot's itemID-0 shape) - fakeEngine's FailGear sentinel
		// cannot target an empty fingerprint (it treats "" as "unset"),
		// so DPSFunc singles out the baseline call by its empty gear.
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if len(req.Character.Gear) == 0 {
				return 0, errors.New("fakeEngine: forced baseline failure")
			}
			return 100, nil
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1")
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1 (the baseline failure)", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 2 {
		t.Fatalf("trinket1 pick = %+v, want item 2 (the ranking itself is unaffected)", out["trinket1"].Item)
	}
	if out["trinket1"].Item.GainMeasured {
		t.Errorf("trinket1 pick GainMeasured = true, want false: the baseline sim failed")
	}
}

// modelledEffectItemID is a real engine-implemented-effect id
// (effectids_generated.go) - any id in that map makes hasImplementedEffect
// true once EffectText is also set, so trinketEffectUnmodelled reads it
// as "modelled" the same way a real proc trinket would be.
const modelledEffectItemID = 647

// Hybrid sweep, bis-ranker-integrity-4 lane, item 4: Serenity Field (an
// unmodelled Spirit self-buff) beat a real combat trinket at band 60
// ret/enhancement by a margin smaller than the tournament's own noise -
// a modelled trinket must keep the slot over a currently-leading
// unmodelled one unless the unmodelled one clears beatsByMargin's own
// 1% bar (verify.go).
func TestRankTrinketSlotModelledTrinketKeepsSlotWhenUnmodelledWinIsWithinMargin(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(modelledEffectItemID, "Real Combat Trinket", 60, "On use: deals damage"),
			trinketWithEffect(999999, "Serenity Field", 60, "Equip: restores spirit over time"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: modelledEffectItemID}}): 200,
			// 201 is inside beatsByMargin's 1% bar over 200 (200*1.01 = 202)
			// - not enough to unseat the modelled trinket.
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999999}}): 201,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 60, "", picks, bySlot, "trinket1")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != modelledEffectItemID {
		t.Fatalf("trinket1 pick = %+v, want the modelled trinket (%d) despite the unmodelled one measuring higher raw dps", out["trinket1"].Item, modelledEffectItemID)
	}
	if out["trinket1"].Item.MeasuredDPS != 200 {
		t.Errorf("trinket1 pick MeasuredDPS = %v, want 200 (its own measured dps, not the unmodelled one's)", out["trinket1"].Item.MeasuredDPS)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != 999999 {
		t.Fatalf("trinket1 runner-up = %+v, want the unmodelled trinket (999999): it measured the higher raw dps even though it did not win", out["trinket1"].RunnerUp)
	}
	if out["trinket1"].RunnerUp.MeasuredDPS != 201 {
		t.Errorf("trinket1 runner-up MeasuredDPS = %v, want 201", out["trinket1"].RunnerUp.MeasuredDPS)
	}
}

// The mirror case: once the unmodelled trinket's own measured dps
// clears the noise floor, it is real enough to win outright.
func TestRankTrinketSlotUnmodelledTrinketWinsWhenItClearsTheMargin(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(modelledEffectItemID, "Real Combat Trinket", 60, "On use: deals damage"),
			trinketWithEffect(999999, "Serenity Field", 60, "Equip: restores spirit over time"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: modelledEffectItemID}}): 200,
			// 205 clears 200*1.01 = 202: a real, not noise-level, lead.
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999999}}): 205,
		},
	}
	out, _ := rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 60, "", picks, bySlot, "trinket1")
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 999999 {
		t.Fatalf("trinket1 pick = %+v, want the unmodelled trinket (999999): it cleared the margin", out["trinket1"].Item)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != modelledEffectItemID {
		t.Fatalf("trinket1 runner-up = %+v, want the modelled trinket (%d)", out["trinket1"].RunnerUp, modelledEffectItemID)
	}
}

// With no modelled candidate in the pool at all, the margin rule has
// nothing to defer to - the highest measured dps wins exactly like
// before this lane's own fix, even though both candidates are
// unmodelled.
func TestRankTrinketSlotPicksHighestWhenEveryCandidateIsUnmodelled(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(999998, "Unmodelled A", 60, "Equip: does something unmodelled"),
			trinketWithEffect(999999, "Unmodelled B", 60, "Equip: does something else unmodelled"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999998}}): 100,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999999}}): 100.5,
		},
	}
	out, _ := rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 60, "", picks, bySlot, "trinket1")
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 999999 {
		t.Fatalf("trinket1 pick = %+v, want item 999999 (higher measured dps, no modelled alternative to defer to)", out["trinket1"].Item)
	}
}

func TestTrinketEffectUnmodelled(t *testing.T) {
	cases := []struct {
		name string
		c    candidate
		want bool
	}{
		{"no effect at all is not unmodelled (pure stats)", candidate{ID: 1}, false},
		{"a real, implemented effect is modelled", candidate{ID: modelledEffectItemID, EffectText: "On use: deals damage"}, false},
		{"an effect the engine does not implement is unmodelled", candidate{ID: 999999, EffectText: "Equip: restores spirit"}, true},
	}
	for _, tc := range cases {
		if got := trinketEffectUnmodelled(tc.c); got != tc.want {
			t.Errorf("%s: trinketEffectUnmodelled(%+v) = %v, want %v", tc.name, tc.c, got, tc.want)
		}
	}
}
