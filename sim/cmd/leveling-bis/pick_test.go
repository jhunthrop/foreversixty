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
