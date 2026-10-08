package request

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const realPresetsPath = "../../data/curated/presets.json"

func loadRealPresets(t *testing.T) Presets {
	t.Helper()
	presets, err := LoadPresets(realPresetsPath)
	if err != nil {
		t.Fatalf("LoadPresets(%s): %v", realPresetsPath, err)
	}
	return presets
}

func entryIDs(entries []PresetEntry) []string {
	out := make([]string, len(entries))
	for i, e := range entries {
		out[i] = e.ID
	}
	return out
}

func buffAndDebuffIDs(p ResolvedPreset) []string {
	return append(entryIDs(p.Buffs), entryIDs(p.Debuffs)...)
}

func writePresets(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "presets.json")
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestEveryRealPresetEntryResolvesForEverySpec(t *testing.T) {
	presets := loadRealPresets(t)
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatalf("Resolve(raid, %s): %v", spec.Spec, err)
		}
		if _, err := buffsFor(buffAndDebuffIDs(resolved)); err != nil {
			t.Errorf("%s: the preset's buffs are not valid: %v", spec.Spec, err)
		}
		if _, err := consumes(entryIDs(resolved.Consumes), nil); err != nil {
			t.Errorf("%s: the preset's consumables are not valid: %v", spec.Spec, err)
		}
	}
}

func TestRaidPresetBuffsAreTheSameForEverySpec(t *testing.T) {
	presets := loadRealPresets(t)
	var first []string
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatal(err)
		}
		got := buffAndDebuffIDs(resolved)
		if first == nil {
			first = got
		} else if !slices.Equal(first, got) {
			t.Errorf("%s: buffs differ from the first spec's:\n%v\n%v", spec.Spec, got, first)
		}
	}
}

func TestRaidPresetNamesTheRequiredBuffsAndExcludesTheForbiddenOnes(t *testing.T) {
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["warrior-arms"])
	if err != nil {
		t.Fatal(err)
	}
	have := buffAndDebuffIDs(resolved)
	for _, want := range []string{
		"arcane_brilliance", "gift_of_the_wild", "power_word_fortitude", "divine_spirit",
		"blessing_of_might", "blessing_of_wisdom", "battle_shout", "trueshot_aura",
		"leader_of_the_pack", "sanctity_aura", "strength_of_earth_totem", "grace_of_air_totem",
		"mana_spring_totem", "windfury_totem", "curse_of_elements", "sunder_armor", "faerie_fire",
		"judgement_of_wisdom", "judgement_of_light", "hunters_mark", "curse_of_recklessness",
	} {
		if !slices.Contains(have, want) {
			t.Errorf("the raid preset lacks %q", want)
		}
	}
	for _, banned := range []string{"shadow_weaving", "blessing_of_kings"} {
		if slices.Contains(have, banned) {
			t.Errorf("the raid preset carries %q", banned)
		}
	}
}

func TestRaidPresetConsumesFollowTheRole(t *testing.T) {
	presets := loadRealPresets(t)
	cases := []struct {
		spec    string
		want    []string
		wantNot []string
	}{
		{"mage-fire", []string{"flask_of_supreme_power", "greater_arcane_elixir", "elixir_of_firepower", "food_runn_tum_tuber_surprise", "main_hand_imbue:brilliant_wizard_oil", "major_mana_potion", "conjured_demonic_rune"}, []string{"elixir_of_the_mongoose"}},
		{"warlock-affliction", []string{"elixir_of_shadow_power", "flask_of_supreme_power"}, []string{"elixir_of_firepower"}},
		{"warrior-fury", []string{"elixir_of_the_mongoose", "juju_power", "r_o_i_d_s", "food_grilled_squid", "juju_might", "mighty_rage_potion", "main_hand_imbue:elemental_sharpening_stone", "off_hand_imbue:dense_sharpening_stone"}, []string{"major_mana_potion", "flask_of_supreme_power"}},
		{"warrior-arms", []string{"mighty_rage_potion", "main_hand_imbue:elemental_sharpening_stone"}, []string{"off_hand_imbue:dense_sharpening_stone"}},
		{"paladin-retribution", []string{"elixir_of_the_mongoose", "major_mana_potion"}, []string{"mighty_rage_potion", "conjured_demonic_rune"}},
		{"hunter-marksmanship", []string{"elixir_of_the_mongoose", "major_mana_potion"}, []string{"off_hand_imbue:dense_sharpening_stone", "flask_of_supreme_power"}},
		{"rogue-combat", []string{"elixir_of_the_mongoose"}, []string{"major_mana_potion", "mighty_rage_potion"}},
	}
	for _, c := range cases {
		resolved, err := presets.Resolve(RaidPreset, specs.ByKey[c.spec])
		if err != nil {
			t.Fatalf("%s: %v", c.spec, err)
		}
		got := entryIDs(resolved.Consumes)
		for _, w := range c.want {
			if !slices.Contains(got, w) {
				t.Errorf("%s: consumes %v lack %q", c.spec, got, w)
			}
		}
		for _, w := range c.wantNot {
			if slices.Contains(got, w) {
				t.Errorf("%s: consumes %v carry %q", c.spec, got, w)
			}
		}
	}
}

// In the client Windfury Totem is an aura on the player, not a weapon
// enchant, so the preset carries it as a buff and never as an imbue: a
// rogue keeps both poisons and still receives it, and a dual wielder has
// both hands free for stones.
func TestRaidPresetCarriesWindfuryTotemAsABuffNotAnImbue(t *testing.T) {
	presets := loadRealPresets(t)
	for _, spec := range specs.All {
		resolved, err := presets.Resolve(RaidPreset, spec)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Contains(buffAndDebuffIDs(resolved), "windfury_totem") {
			t.Errorf("%s: the raid preset lacks the Windfury Totem buff", spec.Spec)
		}
		if slices.Contains(entryIDs(resolved.Consumes), "main_hand_imbue:windfury") {
			t.Errorf("%s: the raid preset still spells the totem as a main-hand imbue", spec.Spec)
		}
	}
}

func TestRaidPresetLeavesARogueBothPoisonsAndTheTotem(t *testing.T) {
	presets := loadRealPresets(t)
	for _, name := range []string{"rogue-assassination", "rogue-combat", "rogue-subtlety"} {
		resolved, err := presets.Resolve(RaidPreset, specs.ByKey[name])
		if err != nil {
			t.Fatal(err)
		}
		buffs, consumeIDs := resolved.Layer(nil, leveling.KitConsumes(name, 60))
		if !slices.Contains(buffs, "windfury_totem") {
			t.Errorf("%s: no Windfury Totem in %v", name, buffs)
		}
		for _, hand := range []string{"main_hand_imbue:", "off_hand_imbue:"} {
			held := 0
			for _, id := range consumeIDs {
				if strings.HasPrefix(id, hand) && strings.Contains(id, "poison") {
					held++
				}
				if strings.HasPrefix(id, hand) && strings.Contains(id, "sharpening_stone") {
					t.Errorf("%s: a stone displaced a poison: %v", name, consumeIDs)
				}
			}
			if held != 1 {
				t.Errorf("%s: %s holds %d poisons in %v, want 1", name, hand, held, consumeIDs)
			}
		}
	}
}

func TestRaidPresetGivesDualWieldersAStoneInEachHand(t *testing.T) {
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["warrior-fury"])
	if err != nil {
		t.Fatal(err)
	}
	got := entryIDs(resolved.Consumes)
	for _, want := range []string{"main_hand_imbue:elemental_sharpening_stone", "off_hand_imbue:dense_sharpening_stone"} {
		if !slices.Contains(got, want) {
			t.Errorf("warrior-fury: consumes %v lack %q", got, want)
		}
	}
}

func TestResolveRejectsAnUnknownPresetName(t *testing.T) {
	_, err := loadRealPresets(t).Resolve("bare", specs.ByKey["mage-fire"])
	if !errors.Is(err, ErrUnknownPreset) {
		t.Fatalf("err = %v, want ErrUnknownPreset", err)
	}
}

func TestLoadPresetsRejectsBadIDs(t *testing.T) {
	const template = `{"presets":{"raid":{"label":"x","notes":"n","buffs":[%s],"debuffs":[%s],"consumes":{"groups":{"caster":{"reference_stats":["spell_power"],"ids":[%s]}}}}}}`
	entry := func(id string) string { return `{"id":"` + id + `","label":"L","reason":"r"}` }
	cases := map[string]struct {
		body string
		want error
	}{
		"unknown buff":    {fmt.Sprintf(template, entry("no_such_buff"), "", ""), ErrUnknownBuff},
		"unknown debuff":  {fmt.Sprintf(template, "", entry("no_such_debuff"), ""), ErrUnknownBuff},
		"buff as debuff":  {fmt.Sprintf(template, "", entry("battle_shout"), ""), ErrPresetMisfiled},
		"debuff as buff":  {fmt.Sprintf(template, entry("sunder_armor"), "", ""), ErrPresetMisfiled},
		"unknown consume": {fmt.Sprintf(template, "", "", entry("elixir_of_nothing")), ErrUnknownConsume},
		"world buff":      {fmt.Sprintf(template, entry("songflower_serenade"), "", ""), ErrUnknownBuff},
	}
	for name, c := range cases {
		if _, err := LoadPresets(writePresets(t, c.body)); !errors.Is(err, c.want) {
			t.Errorf("%s: err = %v, want %v", name, err, c.want)
		}
	}
}

func TestLoadPresetsRejectsAnEntryWithoutALabel(t *testing.T) {
	body := `{"presets":{"raid":{"label":"x","buffs":[{"id":"battle_shout","label":""}],"debuffs":[],"consumes":{"groups":{}}}}}`
	if _, err := LoadPresets(writePresets(t, body)); err == nil {
		t.Fatal("want an error for an empty label")
	}
}

func TestLoadPresetsMissingFile(t *testing.T) {
	if _, err := LoadPresets(filepath.Join(t.TempDir(), "absent.json")); err == nil {
		t.Fatal("want an error for a missing file")
	}
}

func TestLayerKeepsTheKitOnTopOfThePreset(t *testing.T) {
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["shaman-enhancement"])
	if err != nil {
		t.Fatal(err)
	}
	buffs, consumeIDs := resolved.Layer([]string{"blessing_of_might"}, []string{"main_hand_imbue:windfury_weapon"})
	if got := countOf(buffs, "blessing_of_might"); got != 1 {
		t.Errorf("blessing_of_might appears %d times in %v, want once", got, buffs)
	}
	if !slices.Contains(buffs, "windfury_totem") {
		t.Errorf("the raid's Windfury Totem is missing from %v", buffs)
	}
	if !slices.Contains(consumeIDs, "main_hand_imbue:windfury_weapon") {
		t.Errorf("the kit's imbue is missing: %v", consumeIDs)
	}
	if !slices.Contains(consumeIDs, "elixir_of_the_mongoose") {
		t.Errorf("the preset's other consumables are missing: %v", consumeIDs)
	}
}

func TestLayerKitImprovedFormWinsOverThePresetPlainForm(t *testing.T) {
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["druid-feral"])
	if err != nil {
		t.Fatal(err)
	}
	buffs, _ := resolved.Layer([]string{"gift_of_the_wild:improved"}, nil)
	if slices.Contains(buffs, "gift_of_the_wild") || !slices.Contains(buffs, "gift_of_the_wild:improved") {
		t.Errorf("buffs = %v, want only the kit's improved form", buffs)
	}
}

func TestLayerWithNoKitIsThePresetItself(t *testing.T) {
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["mage-frost"])
	if err != nil {
		t.Fatal(err)
	}
	buffs, consumeIDs := resolved.Layer(nil, nil)
	if len(buffs) != len(resolved.Buffs)+len(resolved.Debuffs) || len(consumeIDs) != len(resolved.Consumes) {
		t.Errorf("layer dropped entries: %d buffs, %d consumes", len(buffs), len(consumeIDs))
	}
}

func countOf(list []string, id string) int {
	n := 0
	for _, s := range list {
		if s == id {
			n++
		}
	}
	return n
}

func TestResolveFromFileBareIsNil(t *testing.T) {
	got, err := ResolveFromFile(realPresetsPath, BarePreset, specs.Spec{Spec: "shaman-elemental", ClassSlug: "shaman", ReferenceStat: "spell_power"})
	if err != nil || got != nil {
		t.Fatalf("bare = %v, %v; want nil, nil", got, err)
	}
}

func TestResolveFromFileRaidCarriesTheRaidBuffs(t *testing.T) {
	got, err := ResolveFromFile(realPresetsPath, RaidPreset, specs.Spec{Spec: "shaman-elemental", ClassSlug: "shaman", ReferenceStat: "spell_power"})
	if err != nil || got == nil || len(got.Buffs) == 0 {
		t.Fatalf("raid = %v, %v; want a preset with buffs", got, err)
	}
}

func TestResolveFromFileUnknownPresetFails(t *testing.T) {
	if _, err := ResolveFromFile(realPresetsPath, "nope", specs.Spec{Spec: "shaman-elemental", ClassSlug: "shaman", ReferenceStat: "spell_power"}); !errors.Is(err, ErrUnknownPreset) {
		t.Fatalf("err = %v; want ErrUnknownPreset", err)
	}
}

func TestRaidPresetGivesAShapeshifterNoImbue(t *testing.T) {
	resolved, err := loadRealPresets(t).Resolve(RaidPreset, specs.ByKey["druid-feral"])
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range entryIDs(resolved.Consumes) {
		if strings.Contains(id, "_imbue:") {
			t.Errorf("druid-feral carries %q; a cat's claws take no stone", id)
		}
	}
	if !slices.Contains(entryIDs(resolved.Consumes), "elixir_of_the_mongoose") {
		t.Error("druid-feral lost its elixirs with the imbues")
	}
}
