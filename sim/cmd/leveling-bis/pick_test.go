package main

import "testing"

func item(id int, name string, score float64, slots ...string) scored {
	return scored{candidate: candidate{ID: id, Name: name, Slots: slots}, Score: score}
}

func TestCandidatesBySlotExpandsAliasesAndSortsByScoreDesc(t *testing.T) {
	pool := []scored{
		item(1, "Low Ring", 5, "finger1", "finger2"),
		item(2, "High Ring", 10, "finger1", "finger2"),
		item(3, "Helm", 20, "head"),
	}
	bySlot := candidatesBySlot(pool)
	if len(bySlot["finger1"]) != 2 || len(bySlot["finger2"]) != 2 {
		t.Fatalf("finger alias did not expand to both numbered slots: %+v", bySlot)
	}
	if bySlot["finger1"][0].ID != 2 {
		t.Errorf("finger1[0] = %d, want the higher-scored ring (2)", bySlot["finger1"][0].ID)
	}
	if len(bySlot["head"]) != 1 || bySlot["head"][0].ID != 3 {
		t.Errorf("head slot = %+v, want just the helm", bySlot["head"])
	}
}

func TestPickChoosesBestPerSlot(t *testing.T) {
	bySlot := candidatesBySlot([]scored{
		item(1, "Bad Helm", 5, "head"),
		item(2, "Good Helm", 10, "head"),
	})
	result := pick("", bySlot)
	if result["head"].Item == nil || result["head"].Item.ID != 2 {
		t.Fatalf("head pick = %+v, want item 2", result["head"].Item)
	}
	if result["head"].RunnerUp == nil || result["head"].RunnerUp.ID != 1 {
		t.Fatalf("head runner-up = %+v, want item 1", result["head"].RunnerUp)
	}
}

func TestPickLeavesASlotEmptyWithNoCandidates(t *testing.T) {
	result := pick("", candidatesBySlot(nil))
	if result["head"].Item != nil {
		t.Errorf("head pick = %+v, want nil", result["head"].Item)
	}
}

func TestPickFinger2ExcludesFinger1sItem(t *testing.T) {
	bySlot := candidatesBySlot([]scored{
		item(1, "Only Ring", 10, "finger1", "finger2"),
	})
	result := pick("", bySlot)
	if result["finger1"].Item == nil || result["finger1"].Item.ID != 1 {
		t.Fatalf("finger1 = %+v, want item 1", result["finger1"].Item)
	}
	if result["finger2"].Item != nil {
		t.Fatalf("finger2 = %+v, want nil: the only ring is already worn on finger1", result["finger2"].Item)
	}
}

func TestPickFinger2ExcludesSameNameDifferentQuality(t *testing.T) {
	// Two rows sharing a name are the same ring at two qualities
	// (sim/bulk/expand.go's valid(), mirrored here): picking one for
	// finger1 must not let the other fill finger2.
	bySlot := candidatesBySlot([]scored{
		item(1, "Ring of Fate", 10, "finger1", "finger2"),
		item(2, "Ring of Fate", 9, "finger1", "finger2"),
	})
	result := pick("", bySlot)
	if result["finger1"].Item.ID != 1 {
		t.Fatalf("finger1 = %+v, want item 1 (higher score)", result["finger1"].Item)
	}
	if result["finger2"].Item != nil {
		t.Fatalf("finger2 = %+v, want nil: item 2 shares item 1's name", result["finger2"].Item)
	}
}

func TestPickFinger2PicksADifferentRingWhenOneExists(t *testing.T) {
	bySlot := candidatesBySlot([]scored{
		item(1, "Ring A", 10, "finger1", "finger2"),
		item(2, "Ring B", 8, "finger1", "finger2"),
	})
	result := pick("", bySlot)
	if result["finger1"].Item.ID != 1 {
		t.Fatalf("finger1 = %+v, want item 1", result["finger1"].Item)
	}
	if result["finger2"].Item == nil || result["finger2"].Item.ID != 2 {
		t.Fatalf("finger2 = %+v, want item 2", result["finger2"].Item)
	}
}

func TestPickTrinketPairMirrorsFingerRule(t *testing.T) {
	bySlot := candidatesBySlot([]scored{
		item(1, "Only Trinket", 10, "trinket1", "trinket2"),
	})
	result := pick("", bySlot)
	if result["trinket1"].Item == nil || result["trinket1"].Item.ID != 1 {
		t.Fatalf("trinket1 = %+v, want item 1", result["trinket1"].Item)
	}
	if result["trinket2"].Item != nil {
		t.Fatalf("trinket2 = %+v, want nil", result["trinket2"].Item)
	}
}

func TestPickTwoHandedMainHandLeavesOffHandEmpty(t *testing.T) {
	pool := []scored{
		{candidate: candidate{ID: 1, Name: "Great Axe", Slots: []string{"main_hand"}, TwoHand: true}, Score: 20},
		{candidate: candidate{ID: 2, Name: "Off-hand Blade", Slots: []string{"off_hand"}}, Score: 15},
	}
	result := pick("", candidatesBySlot(pool))
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 1 {
		t.Fatalf("main_hand = %+v, want item 1", result["main_hand"].Item)
	}
	if result["off_hand"].Item != nil {
		t.Fatalf("off_hand = %+v, want nil: main_hand is two-handed", result["off_hand"].Item)
	}
}

// A dual-wield spec's main hand never holds a two-hander either, even
// when one outscores every one-hander: score() converts a weapon's raw
// DPS to attack power per slot with no term for the off hand a
// two-hander forfeits, so a two-hander routinely wins this comparison
// even though a real dual-wielder loses an entire second weapon (and,
// for shaman-enhancement, its off-hand imbue) by wearing one. This is
// the bug behind shaman-enhancement's level-20 list picking Smite's
// Mighty Hammer (item 7230, two-hand) for main hand and leaving
// off_hand permanently empty.
func TestPickExcludesTwoHandFromADualWieldersMainHand(t *testing.T) {
	hammer := scored{candidate: candidate{ID: 7230, Name: "Smite's Mighty Hammer", Slots: []string{"main_hand"}, TwoHand: true, ClassID: itemClassWeapon}, Score: 300}
	axe := scored{candidate: candidate{ID: 2, Name: "One-Hand Axe", Slots: []string{"main_hand", "off_hand"}, ClassID: itemClassWeapon}, Score: 20}
	dagger := scored{candidate: candidate{ID: 3, Name: "One-Hand Dagger", Slots: []string{"main_hand", "off_hand"}, ClassID: itemClassWeapon}, Score: 15}
	bySlot := candidatesBySlot([]scored{hammer, axe, dagger})

	result := pick("shaman-enhancement", bySlot)
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 2 {
		t.Fatalf("shaman-enhancement main_hand = %+v, want the one-hand axe (2), not the two-hand hammer despite its higher score", result["main_hand"].Item)
	}
	if result["off_hand"].Item == nil || result["off_hand"].Item.ID != 3 {
		t.Fatalf("shaman-enhancement off_hand = %+v, want the one-hand dagger (3), the next best one-hander", result["off_hand"].Item)
	}

	// A spec that is not a dual-wielder still takes the higher-scoring
	// two-hander: this rule is specific to DualWieldSpecs.
	notDualWield := pick("shaman-elemental", candidatesBySlot([]scored{hammer, axe, dagger}))
	if notDualWield["main_hand"].Item == nil || notDualWield["main_hand"].Item.ID != 7230 {
		t.Fatalf("shaman-elemental main_hand = %+v, want the two-hand hammer (elemental is not a dual-wielder)", notDualWield["main_hand"].Item)
	}
}

func TestPickOneHandedMainHandStillFillsOffHand(t *testing.T) {
	pool := []scored{
		{candidate: candidate{ID: 1, Name: "Dagger", Slots: []string{"main_hand", "off_hand"}}, Score: 20},
		{candidate: candidate{ID: 2, Name: "Sword", Slots: []string{"main_hand", "off_hand"}}, Score: 15},
	}
	result := pick("", candidatesBySlot(pool))
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 1 {
		t.Fatalf("main_hand = %+v, want item 1", result["main_hand"].Item)
	}
	if result["off_hand"].Item == nil || result["off_hand"].Item.ID != 2 {
		t.Fatalf("off_hand = %+v, want item 2 (the next best one-hander)", result["off_hand"].Item)
	}
}

// TestPickOffersMainHandOneHandersToADualWielderSOffHand pins the real
// bug audit-rogue found 2026-09-28: a real one-handed weapon's own
// classItem row carries only "main_hand" as its .Slot -
// data.go's plannerSlots fans finger/trinket into their numbered pair
// but has no alias that fans a one-hander into "off_hand" too - so
// bySlot["off_hand"] never held a weapon at all and every dual-wield
// spec's off hand came back permanently unpicked at every level band.
// The tests above (TestPickTwoHandedMainHandLeavesOffHandEmpty aside)
// hand-construct candidates whose Slots already lists BOTH hands,
// which never exercised this real, single-slot shape.
func TestPickOffersMainHandOneHandersToADualWielderSOffHand(t *testing.T) {
	pool := []scored{
		{candidate: candidate{ID: 1, Name: "Dagger", ClassID: itemClassWeapon, Slots: []string{"main_hand"}}, Score: 20},
		{candidate: candidate{ID: 2, Name: "Sword", ClassID: itemClassWeapon, Slots: []string{"main_hand"}}, Score: 15},
	}
	result := pick("rogue-assassination", candidatesBySlot(pool))
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 1 {
		t.Fatalf("main_hand = %+v, want item 1", result["main_hand"].Item)
	}
	if result["off_hand"].Item == nil || result["off_hand"].Item.ID != 2 {
		t.Fatalf("off_hand = %+v, want item 2 (the next best one-hander), even though its own .Slots names only \"main_hand\"", result["off_hand"].Item)
	}
}

// TestPickDoesNotOfferMainHandOneHandersToANonDualWieldersOffHand is the
// other side of the fix above: a spec absent from DualWieldSpecs must
// not suddenly grow a weapon in its off hand just because one main_hand
// candidate exists.
func TestPickDoesNotOfferMainHandOneHandersToANonDualWieldersOffHand(t *testing.T) {
	pool := []scored{
		{candidate: candidate{ID: 1, Name: "Sword", ClassID: itemClassWeapon, Slots: []string{"main_hand"}}, Score: 20},
	}
	result := pick("mage-fire", candidatesBySlot(pool))
	if result["off_hand"].Item != nil {
		t.Fatalf("off_hand = %+v, want nil: mage-fire does not dual-wield", result["off_hand"].Item)
	}
}

// TestPickExcludesTwoHandersFromTheMergedOffHandPool covers the merge
// helper's own TwoHand filter: a two-hander sitting among the OTHER
// main_hand candidates (not the one actually chosen for main hand) must
// still never reach the off-hand pool.
func TestPickExcludesTwoHandersFromTheMergedOffHandPool(t *testing.T) {
	pool := []scored{
		{candidate: candidate{ID: 1, Name: "Dagger", ClassID: itemClassWeapon, Slots: []string{"main_hand"}}, Score: 20},
		{candidate: candidate{ID: 2, Name: "Great Axe", ClassID: itemClassWeapon, TwoHand: true, Slots: []string{"main_hand"}}, Score: 15},
	}
	result := pick("rogue-assassination", candidatesBySlot(pool))
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 1 {
		t.Fatalf("main_hand = %+v, want item 1 (the dagger)", result["main_hand"].Item)
	}
	if result["off_hand"].Item != nil {
		t.Fatalf("off_hand = %+v, want nil: the only other main_hand candidate is a two-hander", result["off_hand"].Item)
	}
}

// A dual-wield spec's off hand never holds a held item or a shield, even
// when one scores above every one-hander (Grayson's Torch's spirit once
// beat every level-20 dagger on the assassination list).
func TestPickKeepsAHeldItemOutOfADualWieldersOffHand(t *testing.T) {
	torch := scored{candidate: candidate{ID: 1172, Name: "Grayson's Torch", ClassID: armorClassID}, Score: 50}
	dagger := scored{candidate: candidate{ID: 2567, Name: "Dagger", ClassID: itemClassWeapon}, Score: 10}
	bySlot := map[string][]scored{"off_hand": {torch, dagger}}
	if got := pick("rogue-assassination", bySlot)["off_hand"].Item; got == nil || got.ID != 2567 {
		t.Fatalf("assassination off hand = %v, want the dagger", got)
	}
	if got := pick("shaman-elemental", bySlot)["off_hand"].Item; got == nil || got.ID != 1172 {
		t.Fatalf("elemental off hand = %v, want the held item (it does not dual-wield)", got)
	}
}

// The real bug this lane's report names on priest-shadow and every
// warlock spec: a LATER pass (rankSlotWithEffects, trySetCompletion)
// swaps main_hand onto a two-hander after pick() already gave off_hand
// its own, now-stale item for the ORIGINAL one-handed main_hand.
func TestEnforceTwoHandOffHandInvariantClearsAStaleOffHand(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, TwoHand: true}}},
		"off_hand":  {Item: &scored{candidate: candidate{ID: 2}}},
		"head":      {Item: &scored{candidate: candidate{ID: 3}}},
	}
	out := enforceTwoHandOffHandInvariant(picks)
	if out["off_hand"].Item != nil {
		t.Fatalf("off_hand = %+v, want cleared (main_hand is two-handed)", out["off_hand"].Item)
	}
	if out["head"].Item == nil || out["head"].Item.ID != 3 {
		t.Fatalf("head = %+v, want untouched", out["head"].Item)
	}
	if picks["off_hand"].Item == nil {
		t.Fatal("enforceTwoHandOffHandInvariant mutated its input picks")
	}
}

func TestEnforceTwoHandOffHandInvariantLeavesAOneHandedMainHandAlone(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, TwoHand: false}}},
		"off_hand":  {Item: &scored{candidate: candidate{ID: 2}}},
	}
	out := enforceTwoHandOffHandInvariant(picks)
	if out["off_hand"].Item == nil || out["off_hand"].Item.ID != 2 {
		t.Fatalf("off_hand = %+v, want untouched (main_hand is one-handed)", out["off_hand"].Item)
	}
}
