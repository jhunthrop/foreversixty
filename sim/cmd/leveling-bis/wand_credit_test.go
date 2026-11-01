package main

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func wandCandidate(id int, dps float64) *scored {
	return &scored{candidate: candidate{ID: id, ClassID: itemClassWeapon, SubclassID: wandSubclassID, DPS: dps, Slots: []string{"ranged"}}}
}

func damageAction(spellID int32, damage float64) *proto.ActionMetrics {
	return &proto.ActionMetrics{
		Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: spellID}},
		Targets: []*proto.TargetedActionMetrics{{Damage: damage}},
	}
}

// Round 6, warlock T3: the ranker credited a wand's whole tooltip DPS
// whenever the APL had a Shoot line. A band that never shoots must give a
// wand no credit at all, and one that shoots a quarter of the time a quarter.
func TestWandScorePerDPSIsTheMeasuredShootDamageOverTheWornWand(t *testing.T) {
	cases := []struct {
		name                         string
		shootDPS, wandDPS, reference float64
		want                         float64
	}{
		{"never shoots", 0, 76.8, 0.5, 0},
		{"a quarter of the fight", 19.2, 76.8, 0.5, 0.5},
		{"no wand worn", 10, 0, 0.5, 0},
		{"no reference", 10, 76.8, 0, 0},
	}
	for _, c := range cases {
		if got := wandScorePerDPS(c.shootDPS, c.wandDPS, c.reference); math.Abs(got-c.want) > 1e-12 {
			t.Errorf("%s: wandScorePerDPS = %v, want %v", c.name, got, c.want)
		}
	}
	if got := fullWandScorePerDPS(0.5); got != 2 {
		t.Errorf("fullWandScorePerDPS(0.5) = %v, want 2 (the old whole-fight credit)", got)
	}
}

func TestScoreCreditsAWandOnlyAtTheBandsWandCredit(t *testing.T) {
	torch := candidate{DPS: 76.8, Stats: map[string]float64{"spell_power": 6}}
	brilliant := candidate{DPS: 59.6, Stats: map[string]float64{"spell_power": 12}}
	weights := map[string]float64{"spell_power": 1}
	// A band that shoots all fight: Torch's 17 more DPS wins.
	if score(torch, "ranged", weights, 2) <= score(brilliant, "ranged", weights, 2) {
		t.Error("at full credit the higher-DPS wand should win")
	}
	// A band that never shoots: only the stats count.
	if got := score(torch, "ranged", weights, 0); got != 6 {
		t.Errorf("score at zero credit = %v, want 6 (the stats alone)", got)
	}
	if score(torch, "ranged", weights, 0) >= score(brilliant, "ranged", weights, 0) {
		t.Error("at zero credit the wand with more spell damage should win")
	}
}

func TestShootDPSIsShootsShareOfThePlayersOwnDPS(t *testing.T) {
	player := &proto.UnitMetrics{
		Dps: &proto.DistributionMetrics{Avg: 500},
		Actions: []*proto.ActionMetrics{
			damageAction(11661, 9000),
			damageAction(shootSpellID, 1000),
			{Id: &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 13444}}},
		},
	}
	if got := shootDPS(player); math.Abs(got-50) > 1e-9 {
		t.Errorf("shootDPS = %v, want 50 (a tenth of 500)", got)
	}
	if got := shootDPS(&proto.UnitMetrics{Dps: &proto.DistributionMetrics{Avg: 500}, Actions: []*proto.ActionMetrics{damageAction(11661, 9000)}}); got != 0 {
		t.Errorf("shootDPS with no Shoot row = %v, want 0", got)
	}
	if got := shootDPS(nil); got != 0 {
		t.Errorf("shootDPS(nil) = %v, want 0", got)
	}
}

func TestMeasureWandCreditSimsTheBandsSetAndReadsItsShootDamage(t *testing.T) {
	picks := map[string]slotPick{
		"head":   {Item: p(1, false)},
		"ranged": {Item: wandCandidate(279246, 76.8)},
	}
	fake := &fakeEngine{ShootDPS: 19.2}
	credit, err := measureWandCredit(fake, specInfo{Spec: "warlock-affliction"}, "troll", "warlock", 60, "", picks, 0.5)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(credit-0.5) > 1e-12 {
		t.Errorf("credit = %v, want 0.5 (19.2 / 76.8 / 0.5)", credit)
	}
	want := gearKey([]api.GearSlot{{Slot: "head", ItemID: 1}, {Slot: "ranged", ItemID: 279246}})
	if len(fake.ShootCalls) != 1 || fake.ShootCalls[0] != want {
		t.Errorf("shoot runs = %v, want one run of the band's set %s", fake.ShootCalls, want)
	}
}

func TestMeasureWandCreditIsZeroWithoutAWandAndRunsNothing(t *testing.T) {
	fake := &fakeEngine{ShootDPS: 19.2}
	for name, picks := range map[string]map[string]slotPick{
		"empty ranged": {"head": {Item: p(1, false)}},
		"not a wand":   {"ranged": {Item: &scored{candidate: candidate{ID: 5, ClassID: itemClassWeapon, SubclassID: 2, DPS: 30}}}},
	} {
		credit, err := measureWandCredit(fake, specInfo{}, "troll", "warlock", 60, "", picks, 0.5)
		if err != nil || credit != 0 {
			t.Errorf("%s: credit, err = %v, %v; want 0, nil", name, credit, err)
		}
	}
	if len(fake.ShootCalls) != 0 {
		t.Errorf("shoot runs = %v, want none", fake.ShootCalls)
	}
}
