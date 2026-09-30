package main

import (
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
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

// TestEffectVerifiedInSimCatchesAnItemSimdbSilentlyStrips is the
// regression for this lane's own defect: Hand of Justice (11815) is a
// real entry in effectids_generated.go - the engine genuinely
// registers a proc for it (sim/common's HandOfJustice, wowsims-forever)
// - but this build's embedded simdb has no row for it at all (its
// ItemSparse row is missing from this build's client export; see
// sim/internal/simdb's own TestKnownHandOfJusticeRegression). Before
// this lane, hasImplementedEffect alone decided "was this verified",
// so a real sim of Hand of Justice - which simdb.Attach's own
// UnequipUnknown silently strips before the sim ever runs - measured
// bit-identical to no trinket at all and read as a genuine,
// zero-value verification instead of what it actually was: an item
// the sim never wore. effectVerifiedInSim is the fix: it additionally
// requires the id survive simdb.Known, so this exact case reports
// false, the same "not actually verified" answer report.go and
// trinkets.go's own tie-break logic now both ask for.
func TestEffectVerifiedInSimCatchesAnItemSimdbSilentlyStrips(t *testing.T) {
	handOfJustice := effectItem(11815, "Hand of Justice", "Equip: Chance on hit to gain an extra attack.")
	implemented := effectItem(3854, "Frost Tiger Blade", "Launches a bolt of frost.")

	// Any engine-implemented id this build's simdb does not carry must read
	// as unverified; which ids those are moves with the catalogue (Hand of
	// Justice itself is Known since the classic-db supplement), so find one.
	var absent int
	for id := range engineImplementedEffectItemIDs {
		if !simdb.Known(int32(id)) {
			absent = id
			break
		}
	}
	if absent != 0 {
		stripped := effectItem(absent, "absent from simdb", "Equip: some effect.")
		if effectVerifiedInSim(stripped.candidate) {
			t.Errorf("item %d has no simdb row in this build; effectVerifiedInSim = true, want false", absent)
		}
	}
	if !effectVerifiedInSim(implemented.candidate) {
		t.Error("Frost Tiger Blade (3854) is both engine-implemented and known to simdb; effectVerifiedInSim = false, want true")
	}
	// hasImplementedEffect's own, looser question is unaffected: Hand
	// of Justice still gates a real sim (trinketShortlist,
	// slotsNeedingEffectVerification) even though that sim cannot be
	// trusted to have verified its effect.
	if !hasImplementedEffect(handOfJustice.candidate) {
		t.Error("Hand of Justice (11815) is a real effectids_generated.go entry; hasImplementedEffect = false, want true")
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
	// This lane's brief, item 7: the winner carries its own measured
	// DPS from this tournament - report.go's buildReport reads this to
	// publish sim_dps instead of score()'s stat estimate for this slot.
	if out["main_hand"].Item.MeasuredDPS != 200 {
		t.Errorf("main_hand pick MeasuredDPS = %v, want 200", out["main_hand"].Item.MeasuredDPS)
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

// This lane's brief, item 1: the fifth wow-player sweep caught
// warlock-affliction band 60 Horde main_hand publishing Shortsword of
// Vengeance (stats: {}, a proc-only world drop) over Staff of
// Dar'Orahil (+11 int, +10 hit) as a 0.0-DPS "tie", while Alliance,
// the identical pool, published the staff with the sword losing by
// -2.6 DPS - a coin flip decided by this tournament's own measurement
// noise, not a real difference. A candidate whose stat score is lower
// (or zero, like the sword's) must not win the slot unless its
// measured DPS clears beatsByMargin's own 1% bar over the pool's
// stat-scored best (current) - here a 0.5% lead is noise, so the
// stat-scored pick must keep the slot and the proc item becomes its
// alternative.
func TestRankSlotWithEffectsKeepsTheStatScoredPickWithinNoiseMargin(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "Staff of Dar'Orahil (stat pick)"}}},
	}
	bySlot := map[string][]scored{
		// 754 is the real Shortsword of Vengeance id (engineImplemented-
		// EffectItemIDs, effectids_generated.go) - the exact zero-stat
		// proc weapon the brief names.
		"main_hand": {
			effectItem(754, "Shortsword of Vengeance", "text", "main_hand"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 1}}):   100.0,
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 754}}): 100.5, // +0.5%, inside the 1% margin
		},
	}
	out, notes := rankSlotWithEffects(fake, specInfo{}, "human", "warlock", 60, "", picks, bySlot, "main_hand")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["main_hand"].Item == nil || out["main_hand"].Item.ID != 1 {
		t.Fatalf("main_hand pick = %+v, want the stat-scored pick (1) kept - the proc item's lead is inside the noise margin", out["main_hand"].Item)
	}
	if out["main_hand"].RunnerUp == nil || out["main_hand"].RunnerUp.ID != 754 {
		t.Fatalf("runner-up = %+v, want the proc item (754) recorded as the alternative with its measured delta", out["main_hand"].RunnerUp)
	}
}

// The mirror image of the noise-margin test above: a proc item that
// really does clear the 1% bar over the stat-scored best still wins -
// this rule only stops noise-level flips, not a genuine DPS gain.
func TestRankSlotWithEffectsLetsAProcItemWinWhenItClearsTheMargin(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "Stat pick"}}},
	}
	bySlot := map[string][]scored{
		"main_hand": {
			effectItem(754, "Real proc weapon", "text", "main_hand"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 1}}):   100.0,
			gearKey([]api.GearSlot{{Slot: "main_hand", ItemID: 754}}): 102.0, // +2%, clears the 1% margin
		},
	}
	out, notes := rankSlotWithEffects(fake, specInfo{}, "human", "warlock", 60, "", picks, bySlot, "main_hand")
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["main_hand"].Item == nil || out["main_hand"].Item.ID != 754 {
		t.Fatalf("main_hand pick = %+v, want the proc item (754) - it cleared the noise margin", out["main_hand"].Item)
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
