package main

import "testing"

func rangedCandidate(id, requiredLevel, itemLevel int) candidate {
	return candidate{ID: id, RequiredLevel: requiredLevel, ItemLevel: itemLevel, Slots: []string{"ranged"}}
}

func TestLadderWeaponPicksHighestItemLevelWearable(t *testing.T) {
	items := []candidate{
		rangedCandidate(1, 10, 20),
		rangedCandidate(2, 10, 35),
		rangedCandidate(3, 40, 50), // not yet wearable at level 30
		{ID: 4, RequiredLevel: 10, ItemLevel: 999, Slots: []string{"main_hand"}}, // not ranged
	}
	got := ladderWeapon(items, 30)
	if got == nil || got.ID != 2 {
		t.Fatalf("ladderWeapon = %+v, want item 2 (highest item level wearable and ranged)", got)
	}
}

func TestLadderWeaponNoneWearable(t *testing.T) {
	items := []candidate{rangedCandidate(1, 40, 20)}
	if got := ladderWeapon(items, 20); got != nil {
		t.Fatalf("ladderWeapon = %+v, want nil: required_level 40 > level 20", got)
	}
}

func TestLadderWeaponNoRangedItems(t *testing.T) {
	items := []candidate{{ID: 1, RequiredLevel: 10, ItemLevel: 50, Slots: []string{"main_hand"}}}
	if got := ladderWeapon(items, 30); got != nil {
		t.Fatalf("ladderWeapon = %+v, want nil: no ranged candidate", got)
	}
}

func TestLadderCharacterWithWeapon(t *testing.T) {
	w := rangedCandidate(7, 10, 30)
	ch := ladderCharacter("dwarf", "hunter", 20, "0500000", &w)
	if ch.Name != "ladder" || ch.Race != "dwarf" || ch.Class != "hunter" || ch.Level != 20 || ch.Talents != "0500000" {
		t.Fatalf("ladderCharacter = %+v, wrong base fields", ch)
	}
	if len(ch.Gear) != 1 || ch.Gear[0].Slot != "ranged" || ch.Gear[0].ItemID != 7 {
		t.Fatalf("ladderCharacter.Gear = %+v, want one ranged slot with item 7", ch.Gear)
	}
}

func TestLadderCharacterWithoutWeapon(t *testing.T) {
	ch := ladderCharacter("troll", "hunter", 10, "", nil)
	if len(ch.Gear) != 0 {
		t.Fatalf("ladderCharacter.Gear = %+v, want empty: no weapon given", ch.Gear)
	}
}

func TestWeightsRequestCarriesSpecWeightsAndReference(t *testing.T) {
	spec := specInfo{Spec: "hunter-marksmanship", WeightStats: []string{"agility", "crit"}, ReferenceStat: "ranged_attack_power"}
	ch := ladderCharacter("dwarf", "hunter", 20, "", nil)
	req := weightsRequest(spec, ch, 50, 3)
	if req.Spec != "hunter-marksmanship" {
		t.Errorf("req.Spec = %q, want hunter-marksmanship", req.Spec)
	}
	if req.Iterations != 50 || req.RandomSeed != 3 {
		t.Errorf("req.Iterations/RandomSeed = %d/%d, want 50/3", req.Iterations, req.RandomSeed)
	}
	if req.Weights == nil || req.Weights.Reference != "ranged_attack_power" || len(req.Weights.Stats) != 2 {
		t.Fatalf("req.Weights = %+v, want reference ranged_attack_power and 2 stats", req.Weights)
	}
	if req.Character.Name != "ladder" {
		t.Errorf("req.Character = %+v, want the ladder character passed in", req.Character)
	}
}

func TestPlainRequestCarriesGivenCharacterAndIterations(t *testing.T) {
	spec := specInfo{Spec: "mage-fire"}
	ch := ladderCharacter("gnome", "mage", 60, "50", nil)
	req := plainRequest(spec, ch, 300, 7)
	if req.Spec != "mage-fire" || req.Iterations != 300 || req.RandomSeed != 7 {
		t.Fatalf("plainRequest = %+v, wrong fields", req)
	}
	if req.Weights != nil {
		t.Error("plainRequest.Weights is set, want nil for a plain (non-weights) request")
	}
}
