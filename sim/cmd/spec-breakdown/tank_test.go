package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestReplaceGearSwapsAndAddsSlots(t *testing.T) {
	gear := []api.GearSlot{{Slot: "head", ItemID: 1}, {Slot: "off_hand", ItemID: 2}}
	got, err := replaceGear(gear, "off_hand:20688,ranged:7")
	if err != nil {
		t.Fatal(err)
	}
	want := []api.GearSlot{{Slot: "head", ItemID: 1}, {Slot: "off_hand", ItemID: 20688}, {Slot: "ranged", ItemID: 7}}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("slot %d = %v, want %v", i, got[i], want[i])
		}
	}
	if gear[1].ItemID != 2 {
		t.Errorf("the published entry's slice was modified: %v", gear)
	}
}

func TestReplaceGearRefusesAMalformedPair(t *testing.T) {
	for _, list := range []string{"off_hand", "off_hand:", ":5", "off_hand:abc"} {
		if _, err := replaceGear(nil, list); err == nil {
			t.Errorf("%q was accepted", list)
		}
	}
}

func TestBossSwingsTalliesEveryOutcomeOfTheBossSwing(t *testing.T) {
	res := &proto.RaidSimResult{EncounterMetrics: &proto.EncounterMetrics{Targets: []*proto.UnitMetrics{{
		Actions: []*proto.ActionMetrics{{Targets: []*proto.TargetedActionMetrics{{
			Hits: 10, Crits: 5, Misses: 10, Dodges: 30, Parries: 20, Blocks: 20, Crushes: 5,
			Damage: 10000, BlockDamage: 2500,
		}}}},
	}}}}
	got := bossSwings(res, 1)
	if !near(got.Swings, 100) {
		t.Errorf("swings = %v, want 100", got.Swings)
	}
	for name, pair := range map[string][2]float64{
		"miss": {got.Miss, 10}, "dodge": {got.Dodge, 30}, "parry": {got.Parry, 20},
		"block": {got.Block, 20}, "crit": {got.Crit, 5}, "crush": {got.Crush, 5}, "hit": {got.Hit, 10},
	} {
		if !near(pair[0], pair[1]) {
			t.Errorf("%s = %v%%, want %v%%", name, pair[0], pair[1])
		}
	}
	if !near(got.DamagePerSwing, 100) {
		t.Errorf("damage per swing = %v, want 100", got.DamagePerSwing)
	}
}
