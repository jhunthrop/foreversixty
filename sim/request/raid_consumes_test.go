package request

import (
	"bytes"
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const raidConsumesPath = "raid-consumes.json"

// The /sim page's Raid-buffed consumables are this file (owner ruling, round 5): the
// published level-60 numbers' own set, so the page and the published DPS agree by
// default. A stale copy is a failing test here, not two numbers disagreeing quietly.
func TestRaidConsumesJSONIsCommitted(t *testing.T) {
	want, err := RaidConsumesJSON(loadRealPresets(t))
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(raidConsumesPath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.TrimSpace(got), bytes.TrimSpace(want)) {
		t.Error("sim/request/raid-consumes.json is stale; run `go run ./internal/genraidconsumes` from sim/")
	}
}

// Every spec's row is exactly what the published numbers send: the raid preset resolved
// for the spec with the class kit layered on top (spec-breakdown, leveling-bis).
func TestRaidConsumesAreThePublishedSet(t *testing.T) {
	presets := loadRealPresets(t)
	got, err := RaidConsumes(presets)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(specs.All) {
		t.Fatalf("%d specs, want %d", len(got), len(specs.All))
	}
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatal(err)
		}
		_, want := resolved.Layer(leveling.KitBuffs(spec.Spec, api.MaxLevel), leveling.KitConsumes(spec.Spec, api.MaxLevel))
		if !slices.Equal(got[spec.Spec], want) {
			t.Errorf("%s: %v, want %v", spec.Spec, got[spec.Spec], want)
		}
		if len(got[spec.Spec]) == 0 {
			t.Errorf("%s carries no consumables", spec.Spec)
		}
	}
}

func TestRaidConsumesGiveRoguesTheirPoisons(t *testing.T) {
	got, err := RaidConsumes(loadRealPresets(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, spec := range []string{"rogue-assassination", "rogue-combat", "rogue-subtlety"} {
		for _, id := range []string{"main_hand_imbue:deadly_poison", "off_hand_imbue:instant_poison", "conjured_rogue_thistle_tea"} {
			if !slices.Contains(got[spec], id) {
				t.Errorf("%s: %v has no %s", spec, got[spec], id)
			}
		}
		if slices.Contains(got[spec], "main_hand_imbue:elemental_sharpening_stone") {
			t.Errorf("%s: the stone displaced the main-hand poison: %v", spec, got[spec])
		}
	}
	survival := got["hunter-survival"]
	for _, id := range []string{"major_mana_potion", "juju_might"} {
		if !slices.Contains(survival, id) {
			t.Errorf("hunter-survival: %v has no %s", survival, id)
		}
	}
	if slices.Contains(survival, "mighty_rage_potion") || slices.Contains(survival, "flask_of_the_titans") {
		t.Errorf("hunter-survival carries a consumable its published set does not: %v", survival)
	}
}

// Owner ruling (round 5): one priest gives Power Infusion to one player, so no preset and
// no default carries it.
func TestNoRaidDefaultCarriesPowerInfusion(t *testing.T) {
	presets := loadRealPresets(t)
	consumes, err := RaidConsumes(presets)
	if err != nil {
		t.Fatal(err)
	}
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatal(err)
		}
		buffs, _ := resolved.Layer(leveling.KitBuffs(spec.Spec, api.MaxLevel), leveling.KitConsumes(spec.Spec, api.MaxLevel))
		for _, id := range append(buffs, consumes[spec.Spec]...) {
			if buffBase(id) == "power_infusions" {
				t.Errorf("%s: the raid default carries %s", spec.Spec, id)
			}
		}
	}
	raw, err := os.ReadFile(raidConsumesPath)
	if err != nil {
		t.Fatal(err)
	}
	var file map[string]any
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(raw, []byte("power_infusion")) {
		t.Error("raid-consumes.json names Power Infusion")
	}
}

// A request that names no buffs gives no Power Infusion, and the id a player ticks on /sim
// gives exactly one (owner ruling, round 5: offered as a user-toggled external only).
func TestPowerInfusionOnlyWhenTheRequestNamesIt(t *testing.T) {
	none, err := buffsFor(nil)
	if err != nil {
		t.Fatal(err)
	}
	if none.Individual.GetPowerInfusions() != 0 {
		t.Errorf("a request with no buffs carries %d Power Infusions", none.Individual.GetPowerInfusions())
	}
	ticked, err := buffsFor([]string{"power_infusions"})
	if err != nil {
		t.Fatal(err)
	}
	if ticked.Individual.GetPowerInfusions() != 1 {
		t.Errorf("power_infusions gives %d, want 1", ticked.Individual.GetPowerInfusions())
	}
}
