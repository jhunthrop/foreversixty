package main

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

func aura(a int) spellconstEffect { return spellconstEffect{Effect: effectApplyAura, Aura: a} }

func TestContributesDamageReadsDamageResourcesAndProcs(t *testing.T) {
	cases := []struct {
		name    string
		effects []spellconstEffect
		want    bool
	}{
		{"school damage", []spellconstEffect{{Effect: effectSchoolDamage}}, true},
		{"periodic damage", []spellconstEffect{aura(auraPeriodicDamage)}, true},
		{"a haste buff", []spellconstEffect{aura(auraModHaste)}, true},
		{"a proc", []spellconstEffect{aura(auraProcTrigger)}, true},
		{"a triggered spell", []spellconstEffect{{Effect: effectTriggerSpell}}, true},
		{"health funnel: own health paid into a heal", []spellconstEffect{aura(auraPeriodicHeal), aura(auraPeriodicDamage)}, false},
		{"a taunt", []spellconstEffect{{Effect: 63}}, false},
		{"an armor buff", []spellconstEffect{{Effect: effectApplyAura, Aura: auraModArmor, Amount: 210}}, false},
		{"an armor debuff on the target", []spellconstEffect{{Effect: effectApplyAura, Aura: auraModArmor, Amount: -90}}, true},
		{"no effects", nil, false},
	}
	for _, tc := range cases {
		if got := contributesDamage(tc.effects); got != tc.want {
			t.Errorf("%s: contributesDamage = %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestSurvivalOnlyNeedsModifierRowsAndOnlySafeNamedSpells(t *testing.T) {
	mod := []spellconstEffect{aura(auraAddPctModifier)}
	spells := map[string]spellconstSpell{
		"1":  {Name: "Demon Armor", Effects: []spellconstEffect{aura(22)}},
		"2":  {Name: "Demon Skin", Effects: []spellconstEffect{aura(22)}},
		"3":  {Name: "Searing Pain", Effects: []spellconstEffect{{Effect: effectSchoolDamage}}},
		"4":  {Name: "Armor", Effects: []spellconstEffect{aura(22)}},
		"10": {Name: "Demonic Aegis", Effects: mod},
		"11": {Name: "Demonic Brand", Effects: mod},
		"12": {Name: "Fel Armor Talent", Effects: []spellconstEffect{aura(auraModDamagePercent)}},
		"13": {Name: "Demonic Pact", Effects: mod},
	}
	talents := map[int]tooltipTalent{
		1: {name: "Demonic Aegis", spellIDs: []int{10}, tooltip: "Increases the effectiveness of your Demon Skin and Demon Armor spells by 30%."},
		2: {name: "Demonic Brand", spellIDs: []int{11}, tooltip: "Your Searing Pain brands the target."},
		3: {name: "Fel Armor Talent", spellIDs: []int{12}, tooltip: "Improves your Demon Armor and all damage you deal."},
		4: {name: "Unlisted", spellIDs: []int{99}, tooltip: "Improves your Demon Armor."},
		5: {name: "Demonic Pact", spellIDs: []int{13}, tooltip: "Your Demonic Aegis effect is not cancelled."},
		6: {name: "Molten Skin", spellIDs: []int{10}, tooltip: "Reduces all damage taken by 10%."},
	}
	rule := newSurvivalRule(spells, talents)
	want := map[int]bool{1: true}
	for id, tt := range talents {
		if got := rule.survivalOnly(tt); got != want[id] {
			t.Errorf("%s: survivalOnly = %v, want %v", tt.name, got, want[id])
		}
	}
	if got := rule.damage.namedSpells("your Demon Armor spell", ""); len(got) != 1 || got[0] != "Demon Armor" {
		t.Errorf("a longer name must win over the shorter one inside it: got %v", got)
	}
}

// The warlock filler the 2026-10-28 report named is survival-only in the
// active build's own data, and the talents that do (or may) add damage
// are not.
func TestWarlockSurvivalOnlyTalentsFromTheActiveBuild(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	build, err := leveling.ReadActiveBuild(root)
	if err != nil {
		t.Fatal(err)
	}
	buildDir := filepath.Join(root, "data", "builds", build)
	got, err := survivalOnlyTalents(buildDir, "warlock")
	if err != nil {
		t.Fatal(err)
	}
	talents, err := loadTooltipTalents(buildDir, "warlock")
	if err != nil {
		t.Fatal(err)
	}
	byName := map[string]int{}
	for id, tt := range talents {
		byName[tt.name] = id
	}
	for name, want := range map[string]bool{
		"Improved Health Funnel": true,
		"Demonic Aegis":          true,
		"Improved Voidwalker":    true,
		"Demonic Brand":          false,
		"Improved Sayaad":        false,
		"Improved Imp":           false,
		"Molten Skin":            false,
		"Grim Reach":             false,
		"Demonic Pact":           false,
		"Bane of Havoc":          false,
	} {
		id, ok := byName[name]
		if !ok {
			t.Fatalf("the active build has no warlock talent %q", name)
		}
		if got[id] != want {
			t.Errorf("%s: survival-only = %v, want %v (tooltip %q)", name, got[id], want, talents[id].tooltip)
		}
	}
}
