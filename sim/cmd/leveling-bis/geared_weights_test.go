package main

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// wearsArmor reports whether a weights character wears anything beyond the
// weapons the bare ladder character carries (ladderCharacter).
func wearsArmor(gear []api.GearSlot) bool {
	for _, g := range gear {
		if !weaponSlots[g.Slot] {
			return true
		}
	}
	return false
}

// TestRunSpecMeasuresWeightsOnTheBandsOwnGear is round 8's ranker finding:
// the weights the band ranks by were measured on the bare ladder character,
// whose tiny mana pool and missing stats misprice every stat a geared
// character carries (Retribution's Intellect 6.7 attack power naked against
// 3.4 on its own band-60 set; Fury's hit 38 against 22). The band's weights
// must come from a sweep of a character wearing the band's first-pass set.
func TestRunSpecMeasuresWeightsOnTheBandsOwnGear(t *testing.T) {
	const nakedAgility, gearedAgility = 1.8, 0.9
	var gearedSweeps int
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsFunc: func(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
			agility := nakedAgility
			if wearsArmor(req.Character.Gear) {
				agility = gearedAgility
				gearedSweeps++
			}
			return map[string]api.StatWeight{
				"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0, Error: 0.01},
				"agility":             {Stat: "agility", Weight: agility, Error: 0.01},
			}, 1.0, nil
		},
	}
	// Band 30: the fixture's helm (quest floor 25) joins the bow there.
	outDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", outDir, "hunter-marksmanship", []int{30}, 5, identityTalentLayout); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	if gearedSweeps == 0 {
		t.Fatal("no weights sweep wore the band's gear; every sweep measured the bare ladder character")
	}
	report := readSpecReportForTest(t, filepath.Join(outDir, "hunter-marksmanship.json"))
	for _, b := range report.Bands {
		for _, w := range b.Weights {
			if w.Stat == "agility" && w.Weight != gearedAgility {
				t.Errorf("band %d %s agility weight = %v, want the geared sweep's %v", b.Band, b.Faction, w.Weight, gearedAgility)
			}
		}
	}
}

// TestRunSpecKeepsTheLadderWeightsWhenTheGearedSweepIsUntrustworthy: a geared
// sweep whose reference stat is not positive beyond its error (even at the
// retry) must not replace the ladder's trusted weights.
func TestRunSpecKeepsTheLadderWeightsWhenTheGearedSweepIsUntrustworthy(t *testing.T) {
	const nakedAgility = 1.8
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsFunc: func(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
			if wearsArmor(req.Character.Gear) {
				return map[string]api.StatWeight{
					"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1, Error: -5},
					"agility":             {Stat: "agility", Weight: 0.1, Error: 0.05},
				}, -0.02, nil
			}
			return map[string]api.StatWeight{
				"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0, Error: 0.01},
				"agility":             {Stat: "agility", Weight: nakedAgility, Error: 0.01},
			}, 1.0, nil
		},
	}
	outDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", outDir, "hunter-marksmanship", []int{30}, 5, identityTalentLayout); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	report := readSpecReportForTest(t, filepath.Join(outDir, "hunter-marksmanship.json"))
	for _, b := range report.Bands {
		if b.WeightsReason != "" {
			t.Errorf("band %d %s weights_reason = %q, want \"\" (the ladder sweep was trusted)", b.Band, b.Faction, b.WeightsReason)
		}
		for _, w := range b.Weights {
			if w.Stat == "agility" && w.Weight != nakedAgility {
				t.Errorf("band %d %s agility weight = %v, want the ladder sweep's %v", b.Band, b.Faction, w.Weight, nakedAgility)
			}
		}
	}
}

// gearedWeightsFixture is a fake engine whose ladder sweep reads agility at
// 1.8 and whose geared sweep reads it at 0.9, and a first pass that picks a
// different set for each.
func gearedWeightsFixture() (*fakeEngine, firstPassGear) {
	fake := &fakeEngine{
		WeightsFunc: func(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
			agility := 1.8
			if wearsArmor(req.Character.Gear) {
				agility = 0.9
			}
			return map[string]api.StatWeight{
				"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0, Error: 0.01},
				"agility":             {Stat: "agility", Weight: agility, Error: 0.01},
			}, 1.0, nil
		},
	}
	firstPass := func(w map[string]float64, _ float64) []api.GearSlot {
		if w["agility"] > 1 {
			return []api.GearSlot{{Slot: "head", ItemID: 1}}
		}
		return []api.GearSlot{{Slot: "head", ItemID: 2}}
	}
	return fake, firstPass
}

func TestGearedBandWeightsHandsBackTheLadderSweepItReplaced(t *testing.T) {
	fake, firstPass := gearedWeightsFixture()
	ladder := api.CharacterSpec{Level: 30, Gear: []api.GearSlot{{Slot: "ranged", ItemID: 9}}}
	got, err := gearedBandWeights(fake, specInfo{Spec: "hunter-marksmanship", ReferenceStat: "ranged_attack_power"}, 30, 5, "", ladder, firstPass)
	if err != nil {
		t.Fatal(err)
	}
	if got.ladder == nil || got.ladder.result["agility"].Weight != 1.8 {
		t.Fatalf("ladder = %+v, want the replaced ladder sweep (agility 1.8) for the guard's contender", got.ladder)
	}
}

// chooseRun: a faction publishes the run whose set measures higher. Both
// sets are simmed at the same seed and iteration count, so their
// difference carries far less noise than either error alone; a tie keeps the
// first run (the geared sweep's).
func TestChooseRunPublishesTheLadderSetWhenItMeasuresBetter(t *testing.T) {
	head := func(id int) map[string]slotPick {
		return map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: id}}}}
	}
	cases := []struct {
		name   string
		ladder float64
		want   int
	}{
		{"the ladder set is better beyond error", 110, 1},
		{"the ladder set is ahead within either error", 100.3, 1},
		{"a tie keeps the geared set", 100, 0},
		{"the ladder set is worse", 90, 0},
	}
	for _, tc := range cases {
		fake := &fakeEngine{DPSByGear: map[string]float64{"head=1;": 100, "head=2;": tc.ladder}, DefaultStdErr: 0.5}
		got, scores, err := chooseRun(fake, specInfo{Spec: "hunter-marksmanship", ClassSlug: "hunter"}, "dwarf", 30, "", []map[string]slotPick{head(1), head(2)})
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("%s: chooseRun = %d (scores %v), want %d", tc.name, got, scores, tc.want)
		}
	}
}
