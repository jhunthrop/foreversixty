package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func p(id int, twoHand bool) *scored {
	return &scored{candidate: candidate{ID: id, TwoHand: twoHand}}
}

func TestBuildGearDropsOffHandForATwoHandedPick(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: p(1, true)},
		"off_hand":  {Item: p(2, false)}, // should never happen from pick(), but buildGear defends anyway
		"head":      {Item: p(3, false)},
	}
	gear := buildGear(picks)
	for _, g := range gear {
		if g.Slot == "off_hand" {
			t.Fatalf("gear = %+v, off_hand should be dropped for a two-handed main_hand", gear)
		}
	}
	if len(gear) != 2 {
		t.Fatalf("gear = %+v, want 2 entries (main_hand, head)", gear)
	}
}

func TestBuildGearKeepsOffHandForAOneHandedPick(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: p(1, false)},
		"off_hand":  {Item: p(2, false)},
	}
	gear := buildGear(picks)
	if len(gear) != 2 {
		t.Fatalf("gear = %+v, want both hands", gear)
	}
}

func TestSwapSlotDropsOffHandWhenTheSwappedInMainHandIsTwoHanded(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: p(1, false), RunnerUp: p(9, true)},
		"off_hand":  {Item: p(2, false)},
	}
	gear := swapSlot(picks, "main_hand", 9, true)
	want := map[string]int{"main_hand": 9}
	got := map[string]int{}
	for _, g := range gear {
		got[g.Slot] = g.ItemID
	}
	if len(got) != len(want) || got["main_hand"] != 9 {
		t.Fatalf("swapSlot = %+v, want %+v (off_hand dropped)", got, want)
	}
}

// Real bug this run's own dogfood run hit (this lane's report):
// hunter-survival's level-35 horde main_hand runner-up was the exact
// weapon already worn in off_hand (a dual-wield spec's off_hand pool
// borrows main_hand's one-handers - pick.go's own "off_hand" case) -
// swapSlot must drop off_hand from the trial the same way it already
// drops off_hand for a two-hand swap, or the trial (and the swap
// decision made from it) doubly equips one physical weapon.
func TestSwapSlotDropsOffHandWhenTheSwappedInMainHandMatchesItsCurrentItem(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: p(1, false), RunnerUp: p(9, false)},
		"off_hand":  {Item: p(9, false)},
	}
	gear := swapSlot(picks, "main_hand", 9, false)
	byID := map[string]int{}
	for _, g := range gear {
		byID[g.Slot] = g.ItemID
	}
	if byID["main_hand"] != 9 {
		t.Fatalf("main_hand = %d, want 9", byID["main_hand"])
	}
	if _, ok := byID["off_hand"]; ok {
		t.Fatalf("gear = %+v, off_hand should be dropped (same physical weapon as main_hand's swap)", byID)
	}
}

func TestSwapSlotOnANonWeaponSlotLeavesHandsAlone(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: p(1, false)},
		"off_hand":  {Item: p(2, false)},
		"head":      {Item: p(3, false), RunnerUp: p(4, false)},
	}
	gear := swapSlot(picks, "head", 4, false)
	byID := map[string]int{}
	for _, g := range gear {
		byID[g.Slot] = g.ItemID
	}
	if byID["head"] != 4 {
		t.Fatalf("head = %d, want 4", byID["head"])
	}
	if byID["main_hand"] != 1 || byID["off_hand"] != 2 {
		t.Fatalf("hands changed unexpectedly: %+v", byID)
	}
}

func TestSwapSlotDropsThePairMateWhenTheRunnerUpIsAlreadyWornThere(t *testing.T) {
	// finger1's own runner-up (item 5) happens to be the exact ring
	// finger2 is already wearing - trying it on finger1 must not
	// double-equip it.
	picks := map[string]slotPick{
		"finger1": {Item: p(1, false), RunnerUp: p(5, false)},
		"finger2": {Item: p(5, false)},
	}
	gear := swapSlot(picks, "finger1", 5, false)
	count := 0
	for _, g := range gear {
		if g.ItemID == 5 {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("gear = %+v, want item 5 worn exactly once", gear)
	}
	for _, g := range gear {
		if g.Slot == "finger2" {
			t.Fatalf("gear = %+v, want finger2 dropped rather than double-wearing item 5", gear)
		}
	}
}

func TestSwapSlotKeepsThePairMateWhenTheRunnerUpDiffers(t *testing.T) {
	picks := map[string]slotPick{
		"finger1": {Item: p(1, false), RunnerUp: p(9, false)},
		"finger2": {Item: p(5, false)},
	}
	gear := swapSlot(picks, "finger1", 9, false)
	byID := map[string]int{}
	for _, g := range gear {
		byID[g.Slot] = g.ItemID
	}
	if byID["finger1"] != 9 || byID["finger2"] != 5 {
		t.Fatalf("gear = %+v, want finger1=9 finger2=5", byID)
	}
}

// A sanity check that this file's gear rows are shaped the way
// plainRequest/api.CharacterSpec expects, so a type mismatch would
// fail to compile rather than fail at request-build time.
func TestBuildGearProducesAPIGearSlots(t *testing.T) {
	var _ []api.GearSlot = buildGear(map[string]slotPick{})
}

func TestVerifyBandNoRunnerUpsReturnsBaselineOnly(t *testing.T) {
	picks := map[string]slotPick{"head": {Item: p(1, false)}}
	fake := &fakeEngine{DefaultDPS: 100}
	dps, swaps, verifyErrors, err := verifyBand(fake, specInfo{Spec: "hunter-marksmanship"}, "dwarf", "hunter", 20, "", picks)
	if err != nil {
		t.Fatalf("verifyBand: %v", err)
	}
	if dps != 100 {
		t.Errorf("baseline dps = %v, want 100", dps)
	}
	if len(swaps) != 0 || len(verifyErrors) != 0 {
		t.Fatalf("swaps/verifyErrors = %v/%v, want both empty: no runner-ups", swaps, verifyErrors)
	}
}

func TestVerifyBandRunnerUpBeatsThePick(t *testing.T) {
	picks := map[string]slotPick{
		"head": {Item: p(1, false), RunnerUp: p(2, false)},
	}
	fake := &fakeEngine{
		DefaultDPS: 100, // the baseline (head=1)
		DPSByGear:  map[string]float64{gearKey([]api.GearSlot{{Slot: "head", ItemID: 2}}): 150},
	}
	dps, swaps, verifyErrors, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, "", picks)
	if err != nil {
		t.Fatalf("verifyBand: %v", err)
	}
	if dps != 100 {
		t.Fatalf("baseline dps = %v, want 100", dps)
	}
	if len(swaps) != 1 || swaps[0].Slot != "head" || !swaps[0].Beat || swaps[0].SwapDPS != 150 {
		t.Fatalf("swaps = %+v, want head beaten by 150", swaps)
	}
	if len(verifyErrors) != 0 {
		t.Fatalf("verifyErrors = %v, want none", verifyErrors)
	}
}

func TestVerifyBandRunnerUpLosesToThePick(t *testing.T) {
	picks := map[string]slotPick{
		"head": {Item: p(1, false), RunnerUp: p(2, false)},
	}
	fake := &fakeEngine{
		DefaultDPS: 100,
		DPSByGear:  map[string]float64{gearKey([]api.GearSlot{{Slot: "head", ItemID: 2}}): 50},
	}
	_, swaps, _, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, "", picks)
	if err != nil {
		t.Fatalf("verifyBand: %v", err)
	}
	if len(swaps) != 1 || swaps[0].Beat {
		t.Fatalf("swaps = %+v, want head NOT beaten", swaps)
	}
}

// data-followups-10 lane, 2026-09-30, item 7: a runner-up that
// numerically beats the pick's baseline by more than swapMargin (1%)
// must still NOT be marked Beat when that gap sits inside the two
// runs' own combined standard error - the live repro this pins
// (hunter-marksmanship band 20, Serpent Gloves vs Gloves of the Fang)
// kept reporting a same-fixed-seed "beat by 1.5%" verdict for three
// regen sweeps running, purely from sampling noise neither run's own
// error bar was ever checked against.
func TestVerifyBandRunnerUpNumericallyAheadButWithinNoiseIsNotBeat(t *testing.T) {
	picks := map[string]slotPick{
		"hands": {Item: p(1, false), RunnerUp: p(2, false)},
	}
	fake := &fakeEngine{
		DefaultDPS:    100,
		DefaultStdErr: 1.0,
		DPSByGear:     map[string]float64{gearKey([]api.GearSlot{{Slot: "hands", ItemID: 2}}): 101.2},
		StdErrByGear:  map[string]float64{gearKey([]api.GearSlot{{Slot: "hands", ItemID: 2}}): 1.0},
	}
	_, swaps, _, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, "", picks)
	if err != nil {
		t.Fatalf("verifyBand: %v", err)
	}
	if len(swaps) != 1 {
		t.Fatalf("swaps = %+v, want exactly one", swaps)
	}
	if swaps[0].Significant() {
		t.Fatalf("swaps[0].Significant() = true, want false: 1.5 DPS sits inside the combined error bar (sqrt(1^2+1^2) ~ 1.41)")
	}
	if swaps[0].Beat {
		t.Fatalf("swaps = %+v, want Beat false: a numeric edge inside the combined error bar is not a real win", swaps)
	}
}

func TestVerifyBandBaselineFailurePropagates(t *testing.T) {
	picks := map[string]slotPick{"head": {Item: p(1, false)}}
	fake := &fakeEngine{FailGear: gearKey(buildGear(picks))}
	_, _, _, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, "", picks)
	if err == nil {
		t.Fatal("verifyBand with a failing baseline: want an error, got nil")
	}
}

func TestVerifyBandSwapFailureIsCollectedNotFatal(t *testing.T) {
	// Two slots each carry a runner-up; one runner-up's own sim fails
	// (an engine-side error, per verify.go's own doc). The OTHER slot's
	// swap must still be tried and reported - one bad candidate must
	// not lose the whole band's verification.
	picks := map[string]slotPick{
		"head": {Item: p(1, false), RunnerUp: p(2, false)},
		"neck": {Item: p(3, false), RunnerUp: p(4, false)},
	}
	failingGear := gearKey(swapSlot(picks, "head", 2, false))
	fake := &fakeEngine{
		DefaultDPS: 100,
		FailGear:   failingGear,
		DPSByGear:  map[string]float64{gearKey(swapSlot(picks, "neck", 4, false)): 200},
	}
	_, swaps, verifyErrors, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, "", picks)
	if err != nil {
		t.Fatalf("verifyBand: %v", err)
	}
	if len(verifyErrors) != 1 {
		t.Fatalf("verifyErrors = %v, want exactly 1 (head's failed swap)", verifyErrors)
	}
	if len(swaps) != 1 || swaps[0].Slot != "neck" || !swaps[0].Beat {
		t.Fatalf("swaps = %+v, want neck alone, beaten", swaps)
	}
}

// A runner-up the sim measured ahead of the scored pick becomes the pick,
// the scored pick becomes the row's runner-up, and the set is measured
// once more with the winner in; with no winning swap nothing runs.
func TestApplySwapsPromotesTheMeasuredWinnerAndRemeasuresTheSet(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	better := &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}
	picks := map[string]slotPick{"head": {Item: pick, RunnerUp: better}}
	engine := &fakeEngine{DPSByGear: map[string]float64{gearKey(buildGear(map[string]slotPick{"head": {Item: better}})): 200}, DefaultDPS: 150}
	spec := specInfo{Spec: "hunter-marksmanship", ClassSlug: "hunter"}

	out, dps, gotSwaps, err := applySwaps(engine, spec, "dwarf", "hunter", 30, "", picks, []swapResult{{Slot: "head", SwapDPS: 200, BaselineDPS: 150, Beat: true}}, 150)
	if err != nil {
		t.Fatal(err)
	}
	if out["head"].Item.ID != 2 || out["head"].RunnerUp.ID != 1 {
		t.Fatalf("head = item %d runner-up %d, want the winner (2) promoted over the scored pick (1)", out["head"].Item.ID, out["head"].RunnerUp.ID)
	}
	// This lane's brief, item 7: the promoted item carries the swap's
	// own measured SwapDPS - report.go's buildReport reads this to
	// publish sim_dps instead of score()'s stat estimate for this row.
	if out["head"].Item.MeasuredDPS != 200 {
		t.Errorf("promoted item MeasuredDPS = %v, want 200 (sw.SwapDPS)", out["head"].Item.MeasuredDPS)
	}
	if better.MeasuredDPS != 0 {
		t.Errorf("original RunnerUp scored value MeasuredDPS = %v, want 0: applySwaps must copy, not mutate, the shared candidate", better.MeasuredDPS)
	}
	if dps != 200 {
		t.Fatalf("set DPS after the swap = %v, want the re-measured 200", dps)
	}
	if len(gotSwaps) != 1 || !gotSwaps[0].Beat {
		t.Fatalf("adjusted swaps = %+v, want the one promoted swap still marked Beat", gotSwaps)
	}
	if len(engine.Calls) != 1 {
		t.Fatalf("engine calls = %d, want exactly one re-measure", len(engine.Calls))
	}
	if picks["head"].Item.ID != 1 {
		t.Fatal("applySwaps mutated its input picks")
	}

	same, sameDPS, sameSwaps, err := applySwaps(engine, spec, "dwarf", "hunter", 30, "", picks, []swapResult{{Slot: "head", SwapDPS: 100, BaselineDPS: 150, Beat: false}}, 150)
	if err != nil || sameDPS != 150 || same["head"].Item.ID != 1 || len(engine.Calls) != 1 {
		t.Fatalf("a losing swap must leave picks and DPS alone without a sim: dps=%v item=%d calls=%d err=%v", sameDPS, same["head"].Item.ID, len(engine.Calls), err)
	}
	if len(sameSwaps) != 1 || sameSwaps[0].Beat {
		t.Fatalf("adjusted swaps = %+v, want the losing swap still marked not-beat", sameSwaps)
	}
}

// The literal bug this lane's audit caught by eye: a finger/trinket
// pair-mate's runner-up is ranked before the OTHER half of the pair is
// final (pick.go/trinkets.go), so it can freely equal what the mate
// later becomes; applySwaps must refuse to promote that runner-up
// rather than equip the same physical ring/trinket in both slots.
func TestApplySwapsRefusesToPromoteADuplicateOfItsPairMate(t *testing.T) {
	ring := &scored{candidate: candidate{ID: 5351, Name: "Bounty Hunter's Ring"}}
	otherRing := &scored{candidate: candidate{ID: 3235, Name: "Ring of Scorn"}}
	picks := map[string]slotPick{
		"finger1": {Item: otherRing, RunnerUp: ring},
		// finger2 already settled on the exact item finger1's own
		// (stale) runner-up now points at - decided by a later,
		// independent ranking pass finger1's own runner-up ranking
		// never saw (pick.go's own rule: finger2 excludes finger1's
		// FINAL pick, not the other way around).
		"finger2": {Item: ring},
	}
	engine := &fakeEngine{DefaultDPS: 999} // must never be called: nothing to re-measure once the only swap is refused
	spec := specInfo{Spec: "hunter-marksmanship", ClassSlug: "hunter"}

	out, dps, gotSwaps, err := applySwaps(engine, spec, "dwarf", "hunter", 30, "", picks, []swapResult{{Slot: "finger1", SwapDPS: 200, BaselineDPS: 150, Beat: true}}, 150)
	if err != nil {
		t.Fatal(err)
	}
	if out["finger1"].Item.ID != otherRing.ID {
		t.Fatalf("finger1 = item %d, want the original pick (%d) kept - promoting %d would duplicate finger2", out["finger1"].Item.ID, otherRing.ID, ring.ID)
	}
	if out["finger2"].Item.ID != ring.ID {
		t.Fatalf("finger2 = item %d, want it untouched", out["finger2"].Item.ID)
	}
	if dps != 150 {
		t.Fatalf("set DPS = %v, want the pre-swap baseline (150): nothing was promoted, so nothing should re-measure", dps)
	}
	if len(engine.Calls) != 0 {
		t.Fatalf("engine calls = %d, want 0: a refused promotion must not re-measure the set", len(engine.Calls))
	}
	if len(gotSwaps) != 1 || gotSwaps[0].Beat {
		t.Fatalf("adjusted swaps = %+v, want the refused swap reported as not-beat, so report.go's swap_note does not claim a promotion that did not happen", gotSwaps)
	}
}

// The same refusal, for the main_hand/off_hand pair (pairSlot's own
// newer entry) - the real hunter-survival level-35 horde defect this
// lane's report names: main_hand's runner-up was the exact weapon
// off_hand already wore.
func TestApplySwapsRefusesToPromoteAMainHandDuplicateOfOffHand(t *testing.T) {
	current := &scored{candidate: candidate{ID: 7714, Name: "Hypnotic Blade"}}
	runnerUp := &scored{candidate: candidate{ID: 9680, Name: "Tok'kar's Murloc Shanker"}}
	picks := map[string]slotPick{
		"main_hand": {Item: current, RunnerUp: runnerUp},
		"off_hand":  {Item: runnerUp},
	}
	engine := &fakeEngine{DefaultDPS: 999}
	spec := specInfo{Spec: "hunter-survival", ClassSlug: "hunter"}

	out, dps, gotSwaps, err := applySwaps(engine, spec, "dwarf", "hunter", 35, "", picks, []swapResult{{Slot: "main_hand", SwapDPS: 200, BaselineDPS: 150, Beat: true}}, 150)
	if err != nil {
		t.Fatal(err)
	}
	if out["main_hand"].Item.ID != current.ID {
		t.Fatalf("main_hand = item %d, want the original pick (%d) kept - promoting %d would duplicate off_hand", out["main_hand"].Item.ID, current.ID, runnerUp.ID)
	}
	if dps != 150 || len(engine.Calls) != 0 {
		t.Fatalf("a refused promotion must not re-measure the set: dps=%v calls=%d", dps, len(engine.Calls))
	}
	if len(gotSwaps) != 1 || gotSwaps[0].Beat {
		t.Fatalf("adjusted swaps = %+v, want the refused swap reported as not-beat", gotSwaps)
	}
}

// This lane's brief, defect 3: a two-handed runner-up that beat the
// swap sim (its own SwapDPS already measured with off_hand correctly
// dropped, per swapSlot) is promoted, and applySwaps empties off_hand
// itself as part of that promotion - wouldBreakTwoHandInvariant used
// to refuse this outright instead, silently keeping the weaker
// dual-wield pick even though the sim had already measured the
// two-hander beating it.
func TestApplySwapsPromotesATwoHanderAndEmptiesTheOffHand(t *testing.T) {
	current := &scored{candidate: candidate{ID: 1936, Name: "Goblin Screwdriver"}}
	staff := &scored{candidate: candidate{ID: 3446, Name: "Darkwood Staff", TwoHand: true}}
	held := &scored{candidate: candidate{ID: 3451, Name: "Nightglow Concoction"}}
	picks := map[string]slotPick{
		"main_hand": {Item: current, RunnerUp: staff},
		"off_hand":  {Item: held},
	}
	engine := &fakeEngine{DefaultDPS: 999}
	spec := specInfo{Spec: "priest-shadow", ClassSlug: "priest"}

	out, dps, gotSwaps, err := applySwaps(engine, spec, "undead", "priest", 20, "", picks, []swapResult{{Slot: "main_hand", SwapDPS: 23.4, BaselineDPS: 23.3, Beat: true}}, 23.3)
	if err != nil {
		t.Fatal(err)
	}
	if out["main_hand"].Item == nil || out["main_hand"].Item.ID != staff.ID {
		t.Fatalf("main_hand = %+v, want the promoted two-hander (%d)", out["main_hand"].Item, staff.ID)
	}
	if out["main_hand"].RunnerUp == nil || out["main_hand"].RunnerUp.ID != current.ID {
		t.Fatalf("main_hand runner-up = %+v, want the demoted one-hander (%d)", out["main_hand"].RunnerUp, current.ID)
	}
	if out["off_hand"].Item != nil {
		t.Fatalf("off_hand = %+v, want nil: equipping a two-hander empties it", out["off_hand"].Item)
	}
	if dps != 999 || len(engine.Calls) != 1 {
		t.Fatalf("a promotion must re-measure the set once: dps=%v calls=%d", dps, len(engine.Calls))
	}
	if len(gotSwaps) != 1 || !gotSwaps[0].Beat {
		t.Fatalf("adjusted swaps = %+v, want the promoted swap still marked Beat", gotSwaps)
	}
}

func TestApplySwapsRefusesAnOffHandUnderATwoHandedMainHand(t *testing.T) {
	staff := &scored{candidate: candidate{ID: 3446, Name: "Darkwood Staff", TwoHand: true}}
	held := &scored{candidate: candidate{ID: 3451, Name: "Nightglow Concoction"}}
	picks := map[string]slotPick{
		"main_hand": {Item: staff},
		"off_hand":  {RunnerUp: held},
	}
	engine := &fakeEngine{DefaultDPS: 999}
	spec := specInfo{Spec: "priest-shadow", ClassSlug: "priest"}

	out, _, gotSwaps, err := applySwaps(engine, spec, "undead", "priest", 20, "", picks, []swapResult{{Slot: "off_hand", SwapDPS: 24, BaselineDPS: 23, Beat: true}}, 23)
	if err != nil {
		t.Fatal(err)
	}
	if out["off_hand"].Item != nil || len(engine.Calls) != 0 || gotSwaps[0].Beat {
		t.Fatalf("off_hand = %+v calls=%d swaps=%+v; want no off hand under a two-hander", out["off_hand"].Item, len(engine.Calls), gotSwaps)
	}
}

func TestARunnerUpWithinTheNoiseMarginDoesNotBeatTheScoredPick(t *testing.T) {
	// Mantle of Honor vs Serpent's Shoulders, Alliance level 20, 2026-09-29: 78.7 vs 78.6.
	if beatsByMargin(78.7, 78.6) {
		t.Fatal("0.1 DPS on 78.6 is noise, not a win")
	}
	if !beatsByMargin(80.0, 78.6) {
		t.Fatal("1.8% is a real win")
	}
}
