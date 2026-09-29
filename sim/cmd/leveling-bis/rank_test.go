package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func effectItem(id int, name, effectText string, slots ...string) scored {
	return scored{candidate: candidate{ID: id, Name: name, EffectText: effectText, Slots: slots}}
}

func TestHasImplementedEffect(t *testing.T) {
	// 3854 (Frost Tiger Blade) is a real entry in
	// effectids_generated.go (this lane's own fork work); 424242 is not
	// a real item id at all and can never appear there.
	implemented := effectItem(3854, "Frost Tiger Blade", "Launches a bolt of frost.")
	unimplemented := effectItem(424242, "Nobody's Trinket", "Does something nobody coded.")
	noEffect := effectItem(1, "Plain Helm", "")

	if !hasImplementedEffect(implemented.candidate) {
		t.Error("item 3854 has an engine-implemented effect and effect_text; hasImplementedEffect = false, want true")
	}
	if hasImplementedEffect(unimplemented.candidate) {
		t.Error("item 424242 is not a real item; hasImplementedEffect = true, want false")
	}
	if hasImplementedEffect(noEffect.candidate) {
		t.Error("an item with no effect_text at all; hasImplementedEffect = true, want false")
	}
}

func TestSlotsNeedingEffectVerificationSkipsTrinketsAndPlainSlots(t *testing.T) {
	bySlot := map[string][]scored{
		"trinket1": {effectItem(3854, "Frost Tiger Blade", "text", "trinket1")}, // implemented, but trinket - always excluded
		"head":     {effectItem(1, "Plain Helm", "", "head")},                   // no effect at all
		"main_hand": {
			effectItem(2, "Plain Sword", "", "main_hand"),
			effectItem(3854, "Frost Tiger Blade", "text", "main_hand"), // implemented, real slot
		},
		"waist": {effectItem(424242, "Nobody's Belt", "text", "waist")}, // effect, but not engine-implemented
	}
	got := slotsNeedingEffectVerification(bySlot)
	want := []string{"main_hand"}
	if len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("slotsNeedingEffectVerification = %v, want %v", got, want)
	}
}

func TestRankSlotWithEffectsPicksTheHighestMeasuredDPSAmongImplementedCandidates(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "Plain Sword (score() pick)"}}},
	}
	bySlot := map[string][]scored{
		"main_hand": {
			effectItem(3854, "Frost Tiger Blade", "text", "main_hand"),
			effectItem(7959, "Blight", "text", "main_hand"),
			effectItem(424242, "Nobody's Sword", "text", "main_hand"), // not implemented, excluded
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 1}}):    100, // the original score() pick
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 3854}}): 150,
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 7959}}): 200, // highest - should win
		},
	}
	out, notes := rankSlotWithEffects(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "main_hand")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["main_hand"].Item == nil || out["main_hand"].Item.ID != 7959 {
		t.Fatalf("main_hand pick = %+v, want item 7959 (Blight, highest measured DPS)", out["main_hand"].Item)
	}
	// The unimplemented candidate (424242) must never have been simmed.
	for _, call := range fake.Calls {
		if call == gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 424242}}) {
			t.Fatalf("rankSlotWithEffects simmed an unimplemented-effect candidate: %v", fake.Calls)
		}
	}
	// picks is untouched (immutability rule).
	if picks["main_hand"].Item.ID != 1 {
		t.Fatalf("the input picks map was mutated: %+v", picks["main_hand"])
	}
}

// The real bug this test guards, found in tonight's published output:
// hunter-survival (a leveling.DualWieldSpecs member - its off hand
// holds a real weapon) picked Frost Tiger Blade, a TWO-HAND weapon
// whose frost-bolt proc IS engine-implemented, as its level-35 main
// hand - beating the dual-wield pick on measured DPS alone, with no
// term anywhere in this function for what a two-hander costs a
// dual-wielder (an entire second weapon, buildGear's own off_hand-
// drop rule) - while leaving the off_hand slot's own one-hander pick
// untouched, publishing a two-hand main hand next to an off-hand item
// the character was never actually simmed wearing at once. A
// dual-wielder's main hand must never be offered a two-hander here,
// full stop, the same rule pick.go's own "main_hand" case already
// enforces for the plain score()-based pick.
func TestRankSlotWithEffectsExcludesTwoHandersForADualWieldSpec(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "Black Menace (one-hand, dual-wield pick)"}}},
	}
	bySlot := map[string][]scored{
		"main_hand": {
			{candidate: candidate{ID: 3854, Name: "Frost Tiger Blade", EffectText: "text", Slots: []string{"main_hand"}, TwoHand: true}},
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			// Frost Tiger Blade would win on raw measured DPS alone if it
			// were ever simmed - it must not be.
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 3854}}): 9999,
		},
		DefaultDPS: 100,
	}
	out, notes := rankSlotWithEffects(fake, specInfo{Spec: "hunter-survival"}, "dwarf", "hunter", 35, "", picks, bySlot, "main_hand")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none (nothing left to compare the pick against)", notes)
	}
	if out["main_hand"].Item.ID != 1 {
		t.Fatalf("main_hand = %+v, want the dual-wield pick (1) kept - the two-hander must never be offered", out["main_hand"].Item)
	}
	for _, call := range fake.Calls {
		if call == gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 3854}}) {
			t.Fatalf("rankSlotWithEffects simmed a two-hander for a dual-wield spec: %v", fake.Calls)
		}
	}
}

func TestRankSlotWithEffectsLeavesTheSlotAloneWhenOnlyOneCandidateQualifies(t *testing.T) {
	original := &scored{candidate: candidate{ID: 3854, Name: "Frost Tiger Blade", EffectText: "text"}}
	picks := map[string]slotPick{"main_hand": {Item: original}}
	bySlot := map[string][]scored{
		"main_hand": {
			{candidate: candidate{ID: 3854, Name: "Frost Tiger Blade", EffectText: "text", Slots: []string{"main_hand"}}},
			effectItem(1, "Plain Sword", "", "main_hand"), // no effect, never a candidate here
		},
	}
	fake := &fakeEngine{}
	out, notes := rankSlotWithEffects(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "main_hand")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none (nothing to compare the pick against)", notes)
	}
	if out["main_hand"].Item != original {
		t.Fatalf("main_hand = %+v, want unchanged (only the pick itself qualifies)", out["main_hand"].Item)
	}
	if len(fake.Calls) != 0 {
		t.Fatalf("fake engine was called %d times, want 0 (nothing to verify)", len(fake.Calls))
	}
}
