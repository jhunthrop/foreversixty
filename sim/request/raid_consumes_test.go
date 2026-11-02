package request

import (
	"bytes"
	"encoding/json"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const raidConsumesPath = "raid-consumes.json"

// publishedConsumes is the raid setup at level: the preset resolved for the spec, cut to
// the items the level can use, with the class kit at that level layered on top.
func publishedConsumes(t *testing.T, presets Presets, levels ItemLevels, spec specs.Spec, level int) []string {
	t.Helper()
	resolved, err := presets.Resolve(RaidPreset, spec)
	if err != nil {
		t.Fatal(err)
	}
	_, want := resolved.UsableAt(level, levels).Layer(leveling.KitBuffs(spec.Spec, level), leveling.KitConsumes(spec.Spec, level))
	return want
}

// The /sim page's Raid-buffed consumables are this file (owner ruling, round 5): the
// published level-60 numbers' own set, so the page and the published DPS agree by
// default. A stale copy is a failing test here, not two numbers disagreeing quietly.
func TestRaidConsumesJSONIsCommitted(t *testing.T) {
	want, err := RaidConsumesJSON(loadRealPresets(t), loadRealItemLevels(t))
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

// Every spec's set at every level is exactly what the published numbers send at that level:
// the raid preset resolved for the spec with the class kit of that level layered on top.
func TestRaidConsumesAreThePublishedSetAtEveryLevel(t *testing.T) {
	presets := loadRealPresets(t)
	levels := loadRealItemLevels(t)
	got, err := RaidConsumes(presets, levels)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(specs.All) {
		t.Fatalf("%d specs, want %d", len(got), len(specs.All))
	}
	for _, spec := range specs.All {
		rows := got[spec.Spec]
		if len(rows) == 0 || rows[0].FromLevel != 1 {
			t.Fatalf("%s: rows %+v do not start at level 1", spec.Spec, rows)
		}
		for level := 1; level <= api.MaxLevel; level++ {
			if have, want := ConsumesAtLevel(rows, level), publishedConsumes(t, presets, levels, spec, level); !slices.Equal(have, want) {
				t.Errorf("%s level %d: %v, want %v", spec.Spec, level, have, want)
			}
		}
	}
}

// The rows carry the buff half the published numbers send too: the raid preset's buffs and
// debuffs with the class kit's buffs layered first, so a Protection paladin's Raid-buffed run
// keeps its Devotion Aura as the ranker's does (ledger, round 6: the page sent the raid list
// alone).
func TestRaidSetsCarryThePublishedBuffsAtEveryLevel(t *testing.T) {
	presets := loadRealPresets(t)
	levels := loadRealItemLevels(t)
	got, err := RaidConsumes(presets, levels)
	if err != nil {
		t.Fatal(err)
	}
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatal(err)
		}
		for level := 1; level <= api.MaxLevel; level++ {
			want, _ := resolved.UsableAt(level, levels).Layer(leveling.KitBuffs(spec.Spec, level), leveling.KitConsumes(spec.Spec, level))
			if have := BuffsAtLevel(got[spec.Spec], level); !slices.Equal(have, want) {
				t.Errorf("%s level %d: buffs %v, want %v", spec.Spec, level, have, want)
			}
		}
	}
	if at := BuffsAtLevel(got["paladin-protection"], api.MaxLevel); !slices.Contains(at, "devotion_aura") {
		t.Errorf("paladin-protection at 60: %v has no Devotion Aura", at)
	}
}

// The kit's level gates reach the rows: a rogue's Instant Poisons from 20, Deadly in the
// main hand from 30; an Enhancement shaman's Rockbiter Weapon until 30, then Windfury.
func TestRaidConsumesFollowTheKitsLevelGates(t *testing.T) {
	got, err := RaidConsumes(loadRealPresets(t), loadRealItemLevels(t))
	if err != nil {
		t.Fatal(err)
	}
	combat := got["rogue-combat"]
	if at := ConsumesAtLevel(combat, 25); !slices.Contains(at, "main_hand_imbue:instant_poison") || slices.Contains(at, "main_hand_imbue:deadly_poison") {
		t.Errorf("rogue-combat at 25: %v, want Instant Poison in the main hand", at)
	}
	if at := ConsumesAtLevel(combat, 19); slices.ContainsFunc(at, func(id string) bool { return strings.HasSuffix(id, "_poison") }) {
		t.Errorf("rogue-combat at 19 carries a poison: %v", at)
	}
	enh := got["shaman-enhancement"]
	if at := ConsumesAtLevel(enh, 29); !slices.Contains(at, "main_hand_imbue:rockbiter_weapon") {
		t.Errorf("shaman-enhancement at 29: %v, want Rockbiter Weapon", at)
	}
	if at := ConsumesAtLevel(enh, 30); !slices.Contains(at, "main_hand_imbue:windfury_weapon") {
		t.Errorf("shaman-enhancement at 30: %v, want Windfury Weapon", at)
	}
}

func TestRaidConsumesGiveRoguesTheirPoisons(t *testing.T) {
	all, err := RaidConsumes(loadRealPresets(t), loadRealItemLevels(t))
	if err != nil {
		t.Fatal(err)
	}
	got := map[string][]string{}
	for spec, rows := range all {
		got[spec] = ConsumesAtLevel(rows, api.MaxLevel)
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
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatal(err)
		}
		buffs, consumes := resolved.Layer(leveling.KitBuffs(spec.Spec, api.MaxLevel), leveling.KitConsumes(spec.Spec, api.MaxLevel))
		for _, id := range append(buffs, consumes...) {
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
