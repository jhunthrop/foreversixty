package main

import (
	"strings"
	"testing"
)

// Round 7 (hybrids lane): Balance druid, Elemental shaman and Shadow priest
// published band 60 with Hand of Edward the Odd (a one-hand mace) in the main
// hand and nothing in the off hand. The one-hander replaced a two-handed
// staff, and nothing filled the off hand the staff had emptied, so a held
// orb or tome worth about 4% was never worn or measured.

func offHandFillFixture() (staff, mace, orb, tome, dagger scored) {
	staff = scored{candidate: candidate{ID: 11750, Name: "Kindling Stave", TwoHand: true, ClassID: itemClassWeapon}, Score: 50}
	mace = effectItem(2243, "Hand of Edward the Odd", "Equip: next spell is instant.", "main_hand")
	mace.ClassID = itemClassWeapon
	mace.Score = 30
	orb = scored{candidate: candidate{ID: 9001, Name: "Test Orb", ClassID: armorClassID}, Score: 12}
	tome = scored{candidate: candidate{ID: 9002, Name: "Test Tome", ClassID: armorClassID}, Score: 8}
	dagger = scored{candidate: candidate{ID: 9003, Name: "Test Dagger", ClassID: itemClassWeapon}, Score: 20}
	return
}

func TestFillEmptyOffHandGivesAOneHandedMainHandTheBestHeldItem(t *testing.T) {
	_, mace, orb, tome, _ := offHandFillFixture()
	bySlot := map[string][]scored{"off_hand": {orb, tome}}
	picks := map[string]slotPick{"main_hand": {Item: &mace}}

	got := fillEmptyOffHand("shaman-elemental", bySlot, picks)
	if got["off_hand"].Item == nil || got["off_hand"].Item.ID != orb.ID {
		t.Fatalf("off_hand = %+v, want the best-scored held item (%d)", got["off_hand"].Item, orb.ID)
	}
	if got["off_hand"].RunnerUp == nil || got["off_hand"].RunnerUp.ID != tome.ID {
		t.Fatalf("off_hand runner-up = %+v, want %d", got["off_hand"].RunnerUp, tome.ID)
	}
	if picks["off_hand"].Item != nil {
		t.Fatal("fillEmptyOffHand mutated its input")
	}
}

func TestFillEmptyOffHandLeavesATwoHanderAndAFilledOffHandAlone(t *testing.T) {
	staff, mace, orb, tome, _ := offHandFillFixture()
	bySlot := map[string][]scored{"off_hand": {orb, tome}}

	two := fillEmptyOffHand("shaman-elemental", bySlot, map[string]slotPick{"main_hand": {Item: &staff}})
	if two["off_hand"].Item != nil {
		t.Fatalf("off_hand under a two-hander = %+v, want empty", two["off_hand"].Item)
	}
	kept := fillEmptyOffHand("shaman-elemental", bySlot, map[string]slotPick{"main_hand": {Item: &mace}, "off_hand": {Item: &tome}})
	if kept["off_hand"].Item.ID != tome.ID {
		t.Fatalf("off_hand = %d, want the existing pick (%d) kept", kept["off_hand"].Item.ID, tome.ID)
	}
}

func TestFillEmptyOffHandNeverGivesAShamanAWeapon(t *testing.T) {
	_, mace, _, _, dagger := offHandFillFixture()
	bySlot := map[string][]scored{"off_hand": {dagger}, "main_hand": {mace, dagger}}
	got := fillEmptyOffHand("shaman-elemental", bySlot, map[string]slotPick{"main_hand": {Item: &mace}})
	if got["off_hand"].Item != nil {
		t.Fatalf("off_hand = %+v, want empty: shamans never dual wield", got["off_hand"].Item)
	}
}

func TestFillEmptyOffHandGivesADualWielderAnOffHandWeapon(t *testing.T) {
	_, mace, orb, _, dagger := offHandFillFixture()
	bySlot := map[string][]scored{"off_hand": {orb}, "main_hand": {mace, dagger}}
	got := fillEmptyOffHand("rogue-combat", bySlot, map[string]slotPick{"main_hand": {Item: &mace}})
	if got["off_hand"].Item == nil || got["off_hand"].Item.ID != dagger.ID {
		t.Fatalf("off_hand = %+v, want the one-handed weapon (%d), never the main hand itself or a held item", got["off_hand"].Item, dagger.ID)
	}
}

func TestApplySwapsFillsTheOffHandWhenAOneHanderReplacesATwoHander(t *testing.T) {
	staff, mace, orb, tome, _ := offHandFillFixture()
	bySlot := map[string][]scored{"off_hand": {orb, tome}}
	picks := map[string]slotPick{"main_hand": {Item: &staff, RunnerUp: &mace}}
	engine := &fakeEngine{DefaultDPS: 500}
	spec := specInfo{Spec: "shaman-elemental", ClassSlug: "shaman"}

	out, dps, _, err := applySwaps(engine, spec, "troll", "shaman", 60, "", picks, bySlot, []swapResult{{Slot: "main_hand", SwapDPS: 480, BaselineDPS: 470, Beat: true}}, 470)
	if err != nil {
		t.Fatal(err)
	}
	if out["main_hand"].Item.ID != mace.ID {
		t.Fatalf("main_hand = %d, want the promoted one-hander", out["main_hand"].Item.ID)
	}
	if out["off_hand"].Item == nil || out["off_hand"].Item.ID != orb.ID {
		t.Fatalf("off_hand = %+v, want the best held item (%d) beside the promoted one-hander", out["off_hand"].Item, orb.ID)
	}
	if dps != 500 || len(engine.Calls) != 1 || !strings.Contains(engine.Calls[0], "off_hand=9001;") {
		t.Fatalf("the final set must be measured with the off hand: dps=%v calls=%v", dps, engine.Calls)
	}
}

func TestVerifyBandMeasuresAOneHandRunnerUpWithItsOffHand(t *testing.T) {
	staff, mace, orb, tome, _ := offHandFillFixture()
	bySlot := map[string][]scored{"off_hand": {orb, tome}}
	picks := map[string]slotPick{"main_hand": {Item: &staff, RunnerUp: &mace}}
	engine := &fakeEngine{DefaultDPS: 500}

	if _, _, _, err := verifyBand(engine, specInfo{Spec: "shaman-elemental"}, "troll", "shaman", 60, "", picks, bySlot); err != nil {
		t.Fatal(err)
	}
	want := "main_hand=2243;off_hand=9001;"
	for _, c := range engine.Calls {
		if c == want {
			return
		}
	}
	t.Fatalf("calls = %v, want the one-hand runner-up tried beside its off hand (%s)", engine.Calls, want)
}

func TestRankSlotWithEffectsMeasuresAOneHanderWithItsOffHand(t *testing.T) {
	staff, mace, orb, tome, _ := offHandFillFixture()
	bySlot := map[string][]scored{"main_hand": {staff, mace}, "off_hand": {orb, tome}}
	picks := map[string]slotPick{"main_hand": {Item: &staff}}
	engine := &fakeEngine{DefaultDPS: 500}

	rankSlotWithEffects(engine, specInfo{Spec: "shaman-elemental"}, "troll", "shaman", 60, "", picks, bySlot, "main_hand")
	want := "main_hand=2243;off_hand=9001;"
	for _, c := range engine.Calls {
		if c == want {
			return
		}
	}
	t.Fatalf("calls = %v, want the one-hand candidate measured beside its off hand (%s)", engine.Calls, want)
}
