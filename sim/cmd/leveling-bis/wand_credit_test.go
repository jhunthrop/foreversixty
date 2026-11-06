package main

import (
	"fmt"
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

// Round 9: a wand's score rests on the measured wand credit, and the verify
// swap only measures the runner-up. Shadow band 40 Alliance (gates open, the
// ranks lane's rank choices) scored Jaina's Firestarter first and Goblin
// Igniter second; verify swapped in the Igniter, and Twisted Nether Wand,
// third on score, was never measured. At 10,000 iterations on two seeds it
// is 3.6 DPS better in that set. The wand slot now measures its top
// wandTournamentSize wands by score in the full set and keeps the best.
func TestTheWandSlotKeepsTheBestMeasuredOfItsTopWands(t *testing.T) {
	wand := func(id int, name string, score float64) scored {
		return scored{candidate: candidate{ID: id, Name: name, ClassID: itemClassWeapon, SubclassID: wandSubclassID, DPS: 30, Slots: []string{"ranged"}}, Score: score, HasSource: true}
	}
	pool := []scored{wand(13064, "Jaina's Firestarter", 40), wand(5253, "Goblin Igniter", 39), wand(249144, "Twisted Nether Wand", 38), wand(1, "Fourth", 37)}
	picks := map[string]slotPick{"ranged": slotPickFrom(pool)}
	fake := &fakeEngine{DPSByGear: map[string]float64{
		"ranged=13064;":  144,
		"ranged=5253;":   148,
		"ranged=249144;": 151,
		"ranged=1;":      160, // outside the top three by score: never measured
	}}
	out, _ := rankWandSlot(fake, specInfo{}, "human", "priest", 40, "", picks, map[string][]scored{"ranged": pool})
	if got := out["ranged"].Item.ID; got != 249144 {
		t.Fatalf("ranged = %d, want Twisted Nether Wand (249144), the best measured of the top three", got)
	}
	if ru := out["ranged"].RunnerUp; ru == nil || ru.ID != 5253 {
		t.Errorf("runner-up = %+v, want the second best measured (Goblin Igniter)", ru)
	}
	if len(fake.Calls) != 3 {
		t.Errorf("sims = %d, want 3", len(fake.Calls))
	}
	if out["ranged"].Item.MeasuredDPS != 151 {
		t.Errorf("MeasuredDPS = %v, want 151", out["ranged"].Item.MeasuredDPS)
	}
}

// The wand's worth moves with the rest of the set (a mana-bound Shadow
// priest wands more with less Intellect), so verify's swaps and set
// completion can leave a wand that is no longer the best. On the rerun of
// Shadow band 40 Alliance the tournament before verify kept Goblin Igniter
// (148.7 in the final set), and in that final set Jaina's Firestarter
// measures 151.2. The final check re-measures the top wands in the final
// set and adopts one ahead by swapMargin.
func TestTheFinalWandCheckReplacesAWandTheFinalSetOutgrew(t *testing.T) {
	wand := func(id int, name string, score float64) scored {
		return scored{candidate: candidate{ID: id, Name: name, ClassID: itemClassWeapon, SubclassID: wandSubclassID, DPS: 30, Slots: []string{"ranged"}}, Score: score, HasSource: true}
	}
	pool := []scored{wand(13064, "Jaina's Firestarter", 40), wand(5253, "Goblin Igniter", 39), wand(249144, "Twisted Nether Wand", 38)}
	goblin := pool[1]
	vb := verifiedBand{
		picks:  map[string]slotPick{"ranged": {Item: &goblin}},
		setDPS: 148.7,
		swaps:  []swapResult{{Slot: "ranged", Beat: true}, {Slot: "head"}},
	}
	fake := &fakeEngine{DPSByGear: map[string]float64{"ranged=13064;": 151.2, "ranged=5253;": 148.7, "ranged=249144;": 150.1}}
	out, _ := finalWandCheck(fake, specInfo{}, "human", "priest", 40, "", vb, map[string][]scored{"ranged": pool})
	if got := out.picks["ranged"].Item.ID; got != 13064 {
		t.Fatalf("ranged = %d, want Jaina's Firestarter (13064)", got)
	}
	if out.setDPS != 151.2 {
		t.Errorf("setDPS = %v, want the winning measurement 151.2", out.setDPS)
	}
	if len(out.swaps) != 1 || out.swaps[0].Slot != "head" {
		t.Errorf("swaps = %+v, want the stale ranged swap dropped and the head swap kept", out.swaps)
	}
	// Ahead by less than swapMargin: the wand stays.
	fake = &fakeEngine{DPSByGear: map[string]float64{"ranged=13064;": 149.0, "ranged=5253;": 148.7, "ranged=249144;": 148.0}}
	if out, _ := finalWandCheck(fake, specInfo{}, "human", "priest", 40, "", vb, map[string][]scored{"ranged": pool}); out.picks["ranged"].Item.ID != 5253 {
		t.Errorf("a 0.2%% lead replaced the wand; want it kept under swapMargin")
	}
}

// Round 24 (items lane, from the BiS lane): the Alliance Destruction raid
// set wore Torch of Light (6 spell power, 11 Spirit) over Brilliant Wand
// (9 spell power, 7 Intellect). A raid warlock never shoots, so neither
// wand earns wand credit and both score on their stats; Torch scored first
// only on that band's Spirit weight (0.394, inside its own noise bar; the
// horde run of the same build measured 0.139). The tournament then took the
// higher of two 300-iteration runs, Torch by 0.07 DPS against errors of
// about 2, and nothing after it moves a wand by less than 1%. On the frozen
// set Brilliant Wand is 0.24 to 0.34% ahead at 10,000 iterations on both
// seeds and factions. A lead inside the combined error at the verify
// harness is now settled at the guard harness.
func TestTheWandTournamentSettlesALeadInsideTheErrorAtTheGuardHarness(t *testing.T) {
	wand := func(id int, name string, score float64) scored {
		return scored{candidate: candidate{ID: id, Name: name, ClassID: itemClassWeapon, SubclassID: wandSubclassID, DPS: 60, Slots: []string{"ranged"}}, Score: score, HasSource: true}
	}
	pool := []scored{wand(279246, "Torch of Light", 10.33), wand(249385, "Brilliant Wand", 10.22), wand(249237, "Greater Eternal Wand", 9)}
	picks := map[string]slotPick{"ranged": slotPickFrom(pool)}
	spec := specInfo{Spec: "warlock-destruction", ClassSlug: "warlock", Role: "dps"}
	verifyRun := map[string]float64{"ranged=279246;": 865.24, "ranged=249385;": 865.17, "ranged=249237;": 863.7}
	guardRun := map[string]float64{"ranged=279246;": 731.25, "ranged=249385;": 733.36, "ranged=249237;": 732.8}
	fake := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if req.Iterations == guardIterations(spec) {
				return guardRun[gearKey(req.Character.Gear)], nil
			}
			return verifyRun[gearKey(req.Character.Gear)], nil
		},
		StdErrFunc: func(req api.SimRequest) float64 {
			if req.Iterations == guardIterations(spec) {
				return 0.78
			}
			return 2.1
		},
	}
	out, _ := rankWandSlot(fake, spec, "gnome", "warlock", 60, "", picks, map[string][]scored{"ranged": pool})
	if got := out["ranged"].Item.ID; got != 249385 {
		t.Fatalf("ranged = %d, want Brilliant Wand (249385): Torch's 0.07 lead was inside the error and the guard harness has Brilliant 2.1 ahead", got)
	}
	if ru := out["ranged"].RunnerUp; ru == nil || ru.ID != 249237 {
		t.Errorf("runner-up = %+v, want Greater Eternal Wand, second at the guard harness", ru)
	}
	if got := out["ranged"].Item.MeasuredDPS; got != 733.36 {
		t.Errorf("MeasuredDPS = %v, want the guard harness reading 733.36", got)
	}
}

// A lead beyond the combined error at the verify harness needs no guard run.
func TestTheWandTournamentRunsNoGuardForALeadBeyondTheError(t *testing.T) {
	wand := func(id int, score float64) scored {
		return scored{candidate: candidate{ID: id, Name: fmt.Sprint(id), ClassID: itemClassWeapon, SubclassID: wandSubclassID, DPS: 60, Slots: []string{"ranged"}}, Score: score, HasSource: true}
	}
	pool := []scored{wand(1, 3), wand(2, 2)}
	picks := map[string]slotPick{"ranged": slotPickFrom(pool)}
	fake := &fakeEngine{DPSByGear: map[string]float64{"ranged=1;": 800, "ranged=2;": 790}, DefaultStdErr: 2}
	out, _ := rankWandSlot(fake, specInfo{Role: "dps"}, "orc", "warlock", 60, "", picks, map[string][]scored{"ranged": pool})
	if out["ranged"].Item.ID != 1 {
		t.Fatalf("ranged = %d, want 1", out["ranged"].Item.ID)
	}
	if len(fake.Calls) != 2 {
		t.Errorf("sims = %d, want 2 (no guard run)", len(fake.Calls))
	}
}

// The final check compares the best wand with the worn one inside the same
// tournament: when a tie went to the guard harness both readings are at the
// guard count, so a 300-iteration set figure on another scale cannot fake a
// 1% lead.
func TestTheFinalWandCheckComparesTheWornWandInTheSameHarness(t *testing.T) {
	wand := func(id int, score float64) scored {
		return scored{candidate: candidate{ID: id, Name: fmt.Sprint(id), ClassID: itemClassWeapon, SubclassID: wandSubclassID, DPS: 60, Slots: []string{"ranged"}}, Score: score, HasSource: true}
	}
	pool := []scored{wand(1, 3), wand(2, 2)}
	worn := pool[1]
	vb := verifiedBand{picks: map[string]slotPick{"ranged": {Item: &worn}}, setDPS: 700}
	spec := specInfo{Role: "dps"}
	fake := &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if req.Iterations == guardIterations(spec) {
				return map[string]float64{"ranged=1;": 733, "ranged=2;": 731}[gearKey(req.Character.Gear)], nil
			}
			return map[string]float64{"ranged=1;": 865, "ranged=2;": 864}[gearKey(req.Character.Gear)], nil
		},
		DefaultStdErr: 2,
	}
	out, _ := finalWandCheck(fake, spec, "orc", "warlock", 60, "", vb, map[string][]scored{"ranged": pool})
	if got := out.picks["ranged"].Item.ID; got != 2 {
		t.Fatalf("ranged = %d, want the worn wand kept: 733 against its own 731 is under 1%%", got)
	}
}
