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

// bis-ranker-integrity-7 lane, item 1: candidatesBySlot's own doc -
// an exact score() tie breaks on item level (descending) before item
// id, the mage-arcane band 60 neck repro's own shape (Medallion of the
// Dawn, item level 60, versus Jewel of Kajaro, item level 65, tied by
// this build's own weighted score) rebuilt with plain fixtures so the
// rule itself is pinned independent of any one band's real numbers.
func TestCandidatesBySlotBreaksAnExactScoreTieByHigherItemLevel(t *testing.T) {
	lowerIlvl := item(1, "Lower Ilvl But Lower Id", 10, "neck")
	lowerIlvl.ItemLevel = 60
	higherIlvl := item(2, "Higher Ilvl But Higher Id", 10, "neck")
	higherIlvl.ItemLevel = 65
	bySlot := candidatesBySlot([]scored{lowerIlvl, higherIlvl})
	if bySlot["neck"][0].ID != 2 {
		t.Fatalf("neck[0] = %+v, want item 2 (item level 65 beats item level 60 on an exact score tie, not the lower id)", bySlot["neck"][0])
	}
	if bySlot["neck"][1].ID != 1 {
		t.Fatalf("neck[1] = %+v, want item 1 (the lower-ilvl tied candidate) second", bySlot["neck"][1])
	}
}

// TestCandidatesBySlotBreaksAnExactScoreTieByFewerDeadStatsBeforeItemLevel
// is this lane's brief (bis-ranker-integrity-12), item 5: rogue band
// 20's own repro, all three rogue specs - Serpent Gloves (+6 Agility,
// +7 Spell Power, item level 23) and Gloves of the Fang (+6 Agility,
// +4 Strength, item level 19) scored identically before Strength was
// ever weighted for a rogue (data/curated/specs.json, this lane's own
// second fix) - Spell Power is dead for every rogue spec, and the OLD
// tie-break (item level alone) published the item padded with the
// bigger dead stat, item level 23, as though a higher irrelevant
// number were a real reason to prefer it. Fewer dead stats must win
// the tie even against a LOWER item level.
func TestCandidatesBySlotBreaksAnExactScoreTieByFewerDeadStatsBeforeItemLevel(t *testing.T) {
	moreDeadStats := item(1, "Padded With A Dead Stat", 10, "hands")
	moreDeadStats.ItemLevel = 23
	moreDeadStats.DeadStatCount = 1
	fewerDeadStats := item(2, "All Real Stats", 10, "hands")
	fewerDeadStats.ItemLevel = 19
	fewerDeadStats.DeadStatCount = 0
	bySlot := candidatesBySlot([]scored{moreDeadStats, fewerDeadStats})
	if bySlot["hands"][0].ID != 2 {
		t.Fatalf("hands[0] = %+v, want item 2 (0 dead stats beats item 1's 1 dead stat, despite item 1's higher item level)", bySlot["hands"][0])
	}
	if bySlot["hands"][1].ID != 1 {
		t.Fatalf("hands[1] = %+v, want item 1 second", bySlot["hands"][1])
	}
}

// The tie-break's own final fallback: when item level ALSO ties, id
// ascending still decides, purely for a stable, deterministic sort -
// never observed in this build's real data (this lane's brief), but
// candidatesBySlot must not panic or reorder nondeterministically.
func TestCandidatesBySlotBreaksAFullTieByLowerID(t *testing.T) {
	a := item(2, "B", 10, "neck")
	a.ItemLevel = 60
	b := item(1, "A", 10, "neck")
	b.ItemLevel = 60
	bySlot := candidatesBySlot([]scored{a, b})
	if bySlot["neck"][0].ID != 1 {
		t.Fatalf("neck[0] = %+v, want item 1 (lower id, the final fallback when score AND item level both tie)", bySlot["neck"][0])
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

// A dual-wield spec's main hand does not hold a two-hander that scores
// below the full one-hand PAIR it would replace (twoHandBeatsPair,
// pick.go): score() converts a weapon's raw DPS to attack power per
// slot with no term for the off hand a two-hander forfeits, so a
// two-hander can outscore a SINGLE one-hander even though a real
// dual-wielder loses an entire second weapon (and, for shaman-
// enhancement, its off-hand imbue) by wearing one - comparing against
// the full pair's combined score (main + its own off-hand partner)
// closes that gap. This is the bug behind shaman-enhancement's level-20
// list picking Smite's Mighty Hammer (item 7230, two-hand, scored
// higher than the axe ALONE but lower than the axe+dagger pair) for
// main hand and leaving off_hand permanently empty.
func TestPickExcludesTwoHandFromADualWieldersMainHand(t *testing.T) {
	hammer := scored{candidate: candidate{ID: 7230, Name: "Smite's Mighty Hammer", Slots: []string{"main_hand"}, TwoHand: true, ClassID: itemClassWeapon}, Score: 30}
	axe := scored{candidate: candidate{ID: 2, Name: "One-Hand Axe", Slots: []string{"main_hand", "off_hand"}, ClassID: itemClassWeapon}, Score: 20}
	dagger := scored{candidate: candidate{ID: 3, Name: "One-Hand Dagger", Slots: []string{"main_hand", "off_hand"}, ClassID: itemClassWeapon}, Score: 15}
	bySlot := candidatesBySlot([]scored{hammer, axe, dagger})

	result := pick("shaman-enhancement", bySlot)
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 2 {
		t.Fatalf("shaman-enhancement main_hand = %+v, want the one-hand axe (2): the hammer (30) beats the axe alone (20) but loses to the axe+dagger pair (35)", result["main_hand"].Item)
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

// The other side of the fix (this lane's brief, defect 3, the owner's
// level-20 hunter review): a hunter's melee weapon is a stat stick, not
// its damage source, so its weight run zeros out melee attack_power
// entirely (character.go's weightsRequest doc) and score() is not
// inflating the two-hander's score with an unfairly high dps-derived
// term the way it does for a real melee dual-wielder like
// TestPickExcludesTwoHandFromADualWieldersMainHand above. When the
// two-hander's own score legitimately beats the full pair it would
// replace, hunter-beast-mastery/-marksmanship (both in DualWieldSpecs)
// take it - the actual owner-reported shape: Impaling Harpoon (+9
// agility, scored 18.6) beat Goblin Screwdriver+Poniard (8.3 each,
// 16.6 combined).
func TestPickTakesATwoHanderThatBeatsTheFullPairForAStatStickWeaponSpec(t *testing.T) {
	harpoon := scored{candidate: candidate{ID: 100, Name: "Impaling Harpoon", Slots: []string{"main_hand"}, TwoHand: true, ClassID: itemClassWeapon}, Score: 18.6}
	screwdriver := scored{candidate: candidate{ID: 101, Name: "Goblin Screwdriver", Slots: []string{"main_hand", "off_hand"}, ClassID: itemClassWeapon}, Score: 8.3}
	poniard := scored{candidate: candidate{ID: 102, Name: "Poniard", Slots: []string{"main_hand", "off_hand"}, ClassID: itemClassWeapon}, Score: 8.3}
	bySlot := candidatesBySlot([]scored{harpoon, screwdriver, poniard})

	result := pick("hunter-beast-mastery", bySlot)
	if result["main_hand"].Item == nil || result["main_hand"].Item.ID != 100 {
		t.Fatalf("hunter-beast-mastery main_hand = %+v, want the two-hand Impaling Harpoon (100): 18.6 beats the 8.3+8.3 pair", result["main_hand"].Item)
	}
	if result["off_hand"].Item != nil {
		t.Fatalf("hunter-beast-mastery off_hand = %+v, want nil: the main hand is two-handed", result["off_hand"].Item)
	}
	// The pair's own best one-hander is still the runner-up - the swap
	// pass (verify.go) is what would settle a real close call with an
	// actual sim, same as every other pick() decision.
	if result["main_hand"].RunnerUp == nil || result["main_hand"].RunnerUp.ID != 101 {
		t.Fatalf("hunter-beast-mastery main_hand runner-up = %+v, want the Goblin Screwdriver (101)", result["main_hand"].RunnerUp)
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

// This lane's brief (bis-ranker-integrity-6), item 5: clearTrinketPlaceholders
// must reset BOTH trinket slots (not just one) and leave every other
// slot's own pick() decision untouched.
func TestClearTrinketPlaceholdersResetsBothTrinketSlotsOnly(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder A"}}},
		"trinket2": {Item: &scored{candidate: candidate{ID: 2, Name: "Placeholder B"}}},
		"head":     {Item: &scored{candidate: candidate{ID: 3, Name: "Real Helm"}}},
	}
	out := clearTrinketPlaceholders(picks)
	if out["trinket1"].Item != nil {
		t.Fatalf("trinket1 = %+v, want cleared", out["trinket1"].Item)
	}
	if out["trinket2"].Item != nil {
		t.Fatalf("trinket2 = %+v, want cleared", out["trinket2"].Item)
	}
	if out["head"].Item == nil || out["head"].Item.ID != 3 {
		t.Fatalf("head = %+v, want untouched", out["head"].Item)
	}
	if picks["trinket1"].Item == nil || picks["trinket2"].Item == nil {
		t.Fatal("clearTrinketPlaceholders mutated its input picks")
	}
}

// This lane's brief, defect 4: three one-handers scoring identically
// (8.28 apiece) must all be recorded, not just the lowest-id winner -
// pick()'s own tie-break silently discarded the other two equally-good
// items before this.
func TestPickRecordsEveryEquallyScoredAlternativeAsATie(t *testing.T) {
	pool := []scored{
		item(3, "Bent Blade", 8.28, "head"),
		item(1, "Rusty Sword", 8.28, "head"),
		item(2, "Blackwater Cutlass", 8.28, "head"),
		item(9, "Plain Cap", 5, "head"),
	}
	result := pick("", candidatesBySlot(pool))
	head := result["head"]
	if head.Item == nil || head.Item.ID != 1 {
		t.Fatalf("head pick = %+v, want item 1 (lowest id among the 8.28 tie)", head.Item)
	}
	if len(head.Ties) != 2 {
		t.Fatalf("head ties = %+v, want 2 (items 2 and 3, the other 8.28-scored candidates)", head.Ties)
	}
	gotIDs := map[int]bool{}
	for _, tie := range head.Ties {
		gotIDs[tie.ID] = true
	}
	if !gotIDs[2] || !gotIDs[3] {
		t.Fatalf("head ties = %+v, want ids 2 and 3", head.Ties)
	}
	if gotIDs[9] {
		t.Fatalf("head ties = %+v, want the lower-scored item (9) excluded", head.Ties)
	}
}

// A slot with no tie (the runner-up scores strictly lower) must report
// none.
func TestPickRecordsNoTiesWhenTheRunnerUpScoresLower(t *testing.T) {
	pool := []scored{
		item(1, "Good Helm", 10, "head"),
		item(2, "Bad Helm", 5, "head"),
	}
	result := pick("", candidatesBySlot(pool))
	if len(result["head"].Ties) != 0 {
		t.Fatalf("head ties = %+v, want none: no candidate matched the winner's score", result["head"].Ties)
	}
}

// TestPromoteLowValueWeaponPrefersAWeightStatThenItemLevel is this
// lane's brief, item 1's second half: when every candidate in a
// weapon slot scored 0 (score.go's own caster-DPS fallback still
// leaves this true for a genuinely 0-DPS, 0-relevant-stat weapon),
// the fallback pick is the highest item level among candidates
// carrying ANY of the spec's weight-stat keys, never an arbitrary
// lowest-id tie-break.
func TestPromoteLowValueWeaponPrefersAWeightStatThenItemLevel(t *testing.T) {
	low := item(1, "Notched Shortsword", 0, "main_hand")
	low.ItemLevel = 20
	plain := item(2, "Plain Blade", 0, "main_hand")
	plain.ItemLevel = 30
	withStat := item(3, "Apprentice's Wand", 0, "main_hand")
	withStat.ItemLevel = 15
	withStat.Stats = map[string]float64{"intellect": 3}

	list := []scored{plain, low, withStat} // deliberately not pre-sorted by id
	got := promoteLowValueWeapon(list, []string{"intellect", "spirit"})
	if len(got) == 0 || got[0].ID != 3 {
		t.Fatalf("promoteLowValueWeapon()[0] = %+v, want item 3 (the only one carrying a weight stat, intellect), got order %+v", got[0], got)
	}
	// Among the remaining two (neither carries a weight stat), the
	// higher item level (Plain Blade, 30) comes next, not the lower id.
	if len(got) < 2 || got[1].ID != 2 {
		t.Fatalf("promoteLowValueWeapon()[1] = %+v, want item 2 (higher item level among the stat-less pair)", got[1])
	}
}

// A slot where the top score is already positive is untouched: at
// least one real, scoreable candidate exists, and pick()'s own
// greedy rule already ranks it correctly.
func TestPromoteLowValueWeaponIsNoOpWhenTopScoreIsPositive(t *testing.T) {
	list := []scored{item(1, "Real Weapon", 5, "main_hand"), item(2, "Worthless Weapon", 0, "main_hand")}
	got := promoteLowValueWeapon(list, []string{"intellect"})
	if got[0].ID != 1 {
		t.Fatalf("promoteLowValueWeapon reordered a slot with a genuine positive-scoring candidate: %+v", got)
	}
}

// TestExcludeAbovePvpRankCapDropsOnlyRankAboveTheCap is this lane's
// brief, item 3: a rank-11+ PvP source is dropped from the PICKING
// pool, a rank-10-or-below one is kept, and a non-pvp source is
// unaffected regardless of its own (irrelevant) Rank field.
func TestExcludeAbovePvpRankCapDropsOnlyRankAboveTheCap(t *testing.T) {
	rank10 := item(1, "Rank 10 Sword", 10, "main_hand")
	rank10.HasSource, rank10.Source = true, itemSource{Kind: "pvp", Rank: 10}
	rank18 := item(2, "Grand Marshal's Sword", 20, "main_hand")
	rank18.HasSource, rank18.Source = true, itemSource{Kind: "pvp", Rank: 18}
	questSword := item(3, "Quest Sword", 5, "main_hand")
	questSword.HasSource, questSword.Source = true, itemSource{Kind: "quest"}

	got := excludeAbovePvpRankCap([]scored{rank18, rank10, questSword})
	ids := map[int]bool{}
	for _, c := range got {
		ids[c.ID] = true
	}
	if ids[2] {
		t.Errorf("excludeAbovePvpRankCap kept rank 18: %+v", got)
	}
	if !ids[1] || !ids[3] {
		t.Errorf("excludeAbovePvpRankCap dropped a candidate it should have kept: %+v", got)
	}
}

// If filtering would leave a slot with nothing at all, the original,
// unfiltered list comes back - this lane's brief, item 1's own rule
// ("never empty a weapon slot") outranks the rank cap when a rank-11+
// item is genuinely the only sourced candidate.
func TestExcludeAbovePvpRankCapNeverEmptiesTheOnlyCandidate(t *testing.T) {
	onlyOption := item(1, "Grand Marshal's Sword", 20, "main_hand")
	onlyOption.HasSource, onlyOption.Source = true, itemSource{Kind: "pvp", Rank: 18}

	got := excludeAbovePvpRankCap([]scored{onlyOption})
	if len(got) != 1 || got[0].ID != 1 {
		t.Fatalf("excludeAbovePvpRankCap emptied the only candidate: %+v", got)
	}
}
