package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// raidBuffedDPS is what the fake engine reports for a request carrying
// the fixture raid preset's debuff, so a raid entry's set DPS differs
// from the bare one's.
const (
	bareFixtureDPS = 500.0
	raidFixtureDPS = 650.0
)

func presetFake() *fakeEngine {
	return &fakeEngine{
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if slices.Contains(req.Character.Buffs, "sunder_armor") {
				return raidFixtureDPS, nil
			}
			return bareFixtureDPS, nil
		},
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
	}
}

func runPresetSpec(t *testing.T, fake *fakeEngine, bands []int) specReport {
	t.Helper()
	outDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", outDir, "hunter-marksmanship", bands, 5, identityTalentLayout); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(outDir, "hunter-marksmanship.json"))
	if err != nil {
		t.Fatal(err)
	}
	var got specReport
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	return got
}

func entriesWith(report specReport, band int, preset string) []bandReport {
	var out []bandReport
	for _, b := range report.Bands {
		if b.Band == band && b.Preset == preset {
			out = append(out, b)
		}
	}
	return out
}

func TestEveryBandIsTaggedAndOnlyLevelSixtyGetsARaidEntry(t *testing.T) {
	report := runPresetSpec(t, presetFake(), []int{20, 60})
	for _, b := range report.Bands {
		if b.Preset == "" {
			t.Errorf("band %d %s has no preset tag", b.Band, b.Faction)
		}
	}
	if got := len(entriesWith(report, 20, presetRaid)); got != 0 {
		t.Errorf("band 20 has %d raid entries, want none", got)
	}
	if got := len(entriesWith(report, 20, presetBare)); got != 2 {
		t.Errorf("band 20 has %d bare entries, want one per faction", got)
	}
	for _, preset := range []string{presetBare, presetRaid} {
		entries := entriesWith(report, 60, preset)
		if len(entries) != 2 {
			t.Fatalf("band 60 has %d %s entries, want one per faction", len(entries), preset)
		}
		factions := []string{entries[0].Faction, entries[1].Faction}
		slices.Sort(factions)
		if !slices.Equal(factions, []string{"alliance", "horde"}) {
			t.Errorf("band 60 %s factions = %v", preset, factions)
		}
	}
}

func TestRaidEntryIsMeasuredUnderTheRaidPresetAndBareIsNot(t *testing.T) {
	report := runPresetSpec(t, presetFake(), []int{60})
	for _, b := range entriesWith(report, 60, presetBare) {
		if b.SetDPS != bareFixtureDPS {
			t.Errorf("bare %s set_dps = %v, want %v", b.Faction, b.SetDPS, bareFixtureDPS)
		}
	}
	for _, b := range entriesWith(report, 60, presetRaid) {
		if b.SetDPS != raidFixtureDPS {
			t.Errorf("raid %s set_dps = %v, want %v", b.Faction, b.SetDPS, raidFixtureDPS)
		}
		if len(b.Slots) == 0 || len(b.Weights) == 0 || b.Talents == "" || b.Race == "" {
			t.Errorf("raid %s entry is not the same shape as the bare one: %+v", b.Faction, b)
		}
	}
}

func TestRaidWeightsComeFromTheBuffedCharacter(t *testing.T) {
	fake := presetFake()
	runPresetSpec(t, fake, []int{60})
	raidSweeps, bareSweeps := 0, 0
	for _, buffs := range fake.WeightsBuffsSeen {
		if slices.Contains(buffs, "battle_shout") && slices.Contains(buffs, "sunder_armor") {
			raidSweeps++
		} else if !slices.Contains(buffs, "battle_shout") && !slices.Contains(buffs, "sunder_armor") {
			bareSweeps++
		}
	}
	// A sweep the fake finds insignificant is retried, in both passes alike.
	if raidSweeps == 0 || raidSweeps != bareSweeps {
		t.Errorf("weights sweeps: %d raid and %d bare, want the same positive number (%v)", raidSweeps, bareSweeps, fake.WeightsBuffsSeen)
	}
}

func TestBareRequestsNeverCarryThePreset(t *testing.T) {
	fake := presetFake()
	runPresetSpec(t, fake, []int{20, 30})
	for i, buffs := range fake.BuffsSeen {
		if slices.Contains(buffs, "sunder_armor") || slices.Contains(buffs, "battle_shout") {
			t.Fatalf("call %d of a run without band 60 carried preset buffs: %v", i, buffs)
		}
		if slices.Contains(fake.ConsumesSeen[i], "elixir_of_the_mongoose") {
			t.Fatalf("call %d of a run without band 60 carried preset consumables", i)
		}
	}
}

func TestRaidPassCarriesThePresetConsumables(t *testing.T) {
	fake := presetFake()
	runPresetSpec(t, fake, []int{60})
	raid := 0
	for i, buffs := range fake.BuffsSeen {
		if slices.Contains(buffs, "sunder_armor") {
			raid++
			if !slices.Contains(fake.ConsumesSeen[i], "elixir_of_the_mongoose") {
				t.Fatalf("raid call %d lacks the preset consumable: %v", i, fake.ConsumesSeen[i])
			}
		}
	}
	if raid == 0 {
		t.Fatal("no raid-buffed call was recorded")
	}
}

func TestSpecReportPublishesThePresetsObject(t *testing.T) {
	report := runPresetSpec(t, presetFake(), []int{20})
	raid, ok := report.Presets[presetRaid]
	if !ok {
		t.Fatalf("presets = %v, want a raid entry", report.Presets)
	}
	if raid.Label != "Raid-ready, Phase 1" || raid.Notes == "" {
		t.Errorf("raid label/notes = %q / %q", raid.Label, raid.Notes)
	}
	if len(raid.Buffs) != 1 || raid.Buffs[0].ID != "battle_shout" || raid.Buffs[0].Label != "Battle Shout" {
		t.Errorf("raid buffs = %+v", raid.Buffs)
	}
	if len(raid.Debuffs) != 1 || raid.Debuffs[0].ID != "sunder_armor" {
		t.Errorf("raid debuffs = %+v", raid.Debuffs)
	}
	if len(raid.Consumes) != 1 || raid.Consumes[0].ID != "elixir_of_the_mongoose" {
		t.Errorf("raid consumes = %+v", raid.Consumes)
	}
}

func TestRaidPassDoesNotDisturbTheNextBandsDiffOrWeights(t *testing.T) {
	withRaid := runPresetSpec(t, presetFake(), []int{50, 60})
	bare := entriesWith(withRaid, 60, presetBare)
	raid := entriesWith(withRaid, 60, presetRaid)
	if len(bare) != 2 || len(raid) != 2 {
		t.Fatalf("got %d bare and %d raid band-60 entries", len(bare), len(raid))
	}
	for i := range bare {
		if !slices.Equal(bare[i].NewAtBand, raid[i].NewAtBand) {
			t.Errorf("%s: raid new_at_band %v differs from bare %v (both diff against band 50)", bare[i].Faction, raid[i].NewAtBand, bare[i].NewAtBand)
		}
	}
}
