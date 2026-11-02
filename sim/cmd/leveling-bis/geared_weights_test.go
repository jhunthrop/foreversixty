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

func scoreByHead(scores map[int]float64) setMeasure {
	return func(gear []api.GearSlot) (float64, float64, error) {
		for _, g := range gear {
			if g.Slot == "head" {
				return scores[g.ItemID], 0.5, nil
			}
		}
		return 0, 0.5, nil
	}
}

func TestGearedBandWeightsKeepsTheLadderSweepWhenItsSetMeasuresBetter(t *testing.T) {
	fake, firstPass := gearedWeightsFixture()
	ladder := api.CharacterSpec{Level: 30, Gear: []api.GearSlot{{Slot: "ranged", ItemID: 9}}}
	got, err := gearedBandWeights(fake, specInfo{Spec: "hunter-marksmanship", ReferenceStat: "ranged_attack_power"}, 30, 5, "", ladder, firstPass, scoreByHead(map[int]float64{1: 110, 2: 100}))
	if err != nil {
		t.Fatal(err)
	}
	if w := got.result["agility"].Weight; w != 1.8 {
		t.Errorf("agility = %v, want the ladder sweep's 1.8 (its set measures 110 against 100)", w)
	}
}

func TestGearedBandWeightsTakesTheGearedSweepWhenItsSetIsNoWorse(t *testing.T) {
	fake, firstPass := gearedWeightsFixture()
	ladder := api.CharacterSpec{Level: 30, Gear: []api.GearSlot{{Slot: "ranged", ItemID: 9}}}
	got, err := gearedBandWeights(fake, specInfo{Spec: "hunter-marksmanship", ReferenceStat: "ranged_attack_power"}, 30, 5, "", ladder, firstPass, scoreByHead(map[int]float64{1: 100.2, 2: 100}))
	if err != nil {
		t.Fatal(err)
	}
	if w := got.result["agility"].Weight; w != 0.9 {
		t.Errorf("agility = %v, want the geared sweep's 0.9 (the ladder set is ahead by less than its error)", w)
	}
}

func TestGearedBandWeightsHandsBackTheLadderSweepItReplaced(t *testing.T) {
	fake, firstPass := gearedWeightsFixture()
	ladder := api.CharacterSpec{Level: 30, Gear: []api.GearSlot{{Slot: "ranged", ItemID: 9}}}
	got, err := gearedBandWeights(fake, specInfo{Spec: "hunter-marksmanship", ReferenceStat: "ranged_attack_power"}, 30, 5, "", ladder, firstPass, scoreByHead(map[int]float64{1: 100, 2: 100}))
	if err != nil {
		t.Fatal(err)
	}
	if got.ladder == nil || got.ladder.result["agility"].Weight != 1.8 {
		t.Fatalf("ladder = %+v, want the replaced ladder sweep (agility 1.8) for the guard's contender", got.ladder)
	}
}

func TestLadderWeightsContenderWearsTheFinalSetsTrinkets(t *testing.T) {
	item := func(id int) *scored { return &scored{candidate: candidate{ID: id}} }
	ladderPicks := map[string]slotPick{"head": {Item: item(1)}, "trinket1": {Item: item(90)}}
	final := map[string]slotPick{"head": {Item: item(2)}, "trinket1": {Item: item(91)}, "trinket2": {Item: item(92)}}
	got := ladderWeightsContender(ladderPicks, final)
	if got["head"].Item.ID != 1 || got["trinket1"].Item.ID != 91 || got["trinket2"].Item.ID != 92 {
		t.Errorf("contender = head %d, trinkets %d/%d; want the ladder head 1 and the final trinkets 91/92", got["head"].Item.ID, got["trinket1"].Item.ID, got["trinket2"].Item.ID)
	}
	if ladderPicks["trinket1"].Item.ID != 90 {
		t.Error("ladderWeightsContender changed its input")
	}
}

// The guard's ladder-weights contender goes through the verify swaps the
// band's own set does: a runner-up that measures better is promoted.
func TestLadderWeightsSetRunsTheVerifySwaps(t *testing.T) {
	one := scored{candidate: candidate{ID: 1, Name: "One", Slots: []string{"head"}}, Score: 2}
	two := scored{candidate: candidate{ID: 2, Name: "Two", Slots: []string{"head"}}, Score: 1}
	pass := bandPicks{
		picks:      map[string]slotPick{"head": {Item: &one, RunnerUp: &two}},
		pickBySlot: map[string][]scored{"head": {one, two}},
	}
	fake := &fakeEngine{DefaultDPS: 500, DPSByGear: map[string]float64{"head=2;": 600}, DefaultStdErr: 0.1}
	got, err := ladderWeightsSet(fake, specInfo{Spec: "hunter-marksmanship", ClassSlug: "hunter"}, "dwarf", 30, "", pass, map[string]slotPick{}, setCatalog{})
	if err != nil {
		t.Fatal(err)
	}
	if got["head"].Item == nil || got["head"].Item.ID != 2 {
		t.Fatalf("head = %+v, want the runner-up (id 2) the verify swap promotes", got["head"].Item)
	}
}
