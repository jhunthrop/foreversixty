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
	dps, swaps, verifyErrors, err := verifyBand(fake, specInfo{Spec: "hunter-marksmanship"}, "dwarf", "hunter", 20, picks)
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
	dps, swaps, verifyErrors, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, picks)
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
	_, swaps, _, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, picks)
	if err != nil {
		t.Fatalf("verifyBand: %v", err)
	}
	if len(swaps) != 1 || swaps[0].Beat {
		t.Fatalf("swaps = %+v, want head NOT beaten", swaps)
	}
}

func TestVerifyBandBaselineFailurePropagates(t *testing.T) {
	picks := map[string]slotPick{"head": {Item: p(1, false)}}
	fake := &fakeEngine{FailGear: gearKey(buildGear(picks))}
	_, _, _, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, picks)
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
	_, swaps, verifyErrors, err := verifyBand(fake, specInfo{}, "dwarf", "hunter", 20, picks)
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
