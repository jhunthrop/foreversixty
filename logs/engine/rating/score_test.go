// logs/engine/rating/score_test.go
// The three worked examples from docs/superpowers/specs/2026-09-21-
// performance-rating-design.md §1.6, reproduced as tests. Each uses a fake
// PercentileSource keyed only by bracket.Component (ignoring the raw value
// and every other bracket field, matching how §1.6 itself just states "60th
// pct -> 60" without deriving it from a real distribution) so the test
// isolates the arithmetic §1.6 actually claims: the percentile-vs-absolute
// routing, the death-penalty and completeness formulas that run without
// going through the fake at all, and combine()'s weighted mean and cap.
package rating

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const shazzrahEncounterID = 667

func f64(v float64) *float64 { return &v }

// TestWorkedExampleDPSWarriorFury reproduces spec §1.6's first worked
// example: one avoidable death at 140,000ms of a 180,000ms kill.
func TestWorkedExampleDPSWarriorFury(t *testing.T) {
	const player = "Player-Fury"
	fight := summary.Summary{
		DurationMS: 180000, EncounterID: shazzrahEncounterID, Kill: true,
		Roster: []summary.RosterRow{
			{GUID: player, Class: "Warrior", Spec: "Fury", Role: RoleDPS,
				ExecutionScore: f64(0.97), ActivityPct: 91, ActiveMS: 163800},
		},
		Deaths: []summary.Death{
			{GUID: player, AtMS: 140000, KillingBlow: &summary.DamageRef{SpellID: 19712, SpellName: "Arcane Explosion"}},
		},
		Interrupts: []summary.ExchangeRow{
			{Kind: "interrupt", SourceGUID: player, SpellID: 1715, ExtraSpellID: 19714, ExtraSpellName: "Curse Cast", Count: 2},
		},
		Mechanics: summary.MechanicsBlock{TableFound: true, Rows: []summary.MechanicRow{
			{SpellID: 19714, Name: "Curse Cast", Kind: mechanics.Interrupt, Casts: 3, Damage: 3000},
		}},
		Auras: []summary.AuraTrack{
			{TargetGUID: "Boss", SpellID: 1160, Name: "Demoralizing Shout", Type: "DEBUFF", UptimeMS: 158400, Appliers: []string{player}},
			{TargetGUID: "Boss", SpellID: 7386, Name: "Sunder Armor", Type: "DEBUFF", UptimeMS: 165600, Appliers: []string{player}},
		},
		Combatants: []summary.CombatantRow{
			{GUID: player, Consumables: []summary.AuraRef{{SpellID: 17626}, {SpellID: 1249522}, {SpellID: 25122}}},
		},
		Casts: []summary.CastRow{
			{GUID: player, OwnerGUID: player, SpellID: 17528, Succeeded: 1},
		},
	}

	table := &mechanics.Table{EncounterID: shazzrahEncounterID, Name: "Shazzrah", Mechanics: []mechanics.Mechanic{
		{SpellID: 19712, Name: "Arcane Explosion", Kind: mechanics.Avoidable},
		{SpellID: 19714, Name: "Curse Cast", Kind: mechanics.Interrupt},
	}}
	utilTable := &utility.Table{Spec: "warrior-fury", Owned: []utility.Entry{
		{SpellID: 1160, Name: "Demoralizing Shout", Kind: utility.Debuff, Target: "enemy", Verified: "test"},
		{SpellID: 7386, Name: "Sunder Armor", Kind: utility.Debuff, Target: "enemy", Verified: "test"},
	}}
	cat := &consumables.RoleCatalogue{
		Weights:       map[string]int{"flask": 30, "food": 15, "weapon_enchant": 15, "world_buffs": 20, "combat_potion": 20},
		Flask:         []consumables.Entry{{SpellID: 17626, Name: "Flask of the Titans", Verified: "test"}},
		Food:          []consumables.Entry{{SpellID: 1249522, Name: "Grilled Squid", Verified: "test"}},
		WeaponEnchant: []consumables.Entry{{SpellID: 25122, Name: "Brilliant Wizard Oil", Verified: "test"}},
		WorldBuffs:    []consumables.Entry{{SpellID: 22888, Name: "Rallying Cry", Verified: "test"}},
		CombatPotion:  consumables.PotionGroup{MaxUses: 2, Entries: []consumables.Entry{{SpellID: 17528, Name: "Mighty Rage Potion", Verified: "test"}}},
	}

	src := &fakePercentiles{
		band: BandTypical, bandOK: true,
		placements: map[string]fakePlacement{
			ComponentNameOutput:           {pct: 0.71, n: 100, ok: true},
			ComponentSurvivalAvoidableHit: {pct: 0.40, n: 100, ok: true}, // AvoidableHitScore = 100-40 = 60
			ComponentMechanicsInterrupt:   {pct: 0.55, n: 100, ok: true},
			ComponentNameUtility:          {pct: 0.62, n: 100, ok: true},
			ComponentNamePreparation:      {pct: 0.95, n: 100, ok: true},
			ComponentNameActivity:         {pct: 0.80, n: 100, ok: true},
		},
	}

	card := Score(fight, player, CuratedTables{Mechanics: table, Utility: utilTable, Consumables: cat}, src, nil, DefaultModelInfo())

	checkComponent(t, card, ComponentNameOutput, 71)
	checkComponent(t, card, ComponentNameMechanics, 55)
	checkComponent(t, card, ComponentNameUtility, 62)
	checkComponent(t, card, ComponentNamePreparation, 95)
	checkComponent(t, card, ComponentNameActivity, 80)

	survival := componentByName(card, ComponentNameSurvival)
	if want := 75.9; roundTo1(survival.Score) != want {
		t.Errorf("Survival.Score = %v, want ~%v (spec: 0.65*84.4 + 0.35*60)", survival.Score, want)
	}

	if card.OverallCapped {
		t.Fatal("no death floored DeathScore to zero; the cap must not fire")
	}
	if want := 70.0; math.Round(card.Overall) != want {
		t.Fatalf("round(Overall) = %v (Overall=%v), want %v (spec §1.6's DPS worked example, \"-> 70\")",
			math.Round(card.Overall), card.Overall, want)
	}
	if card.Coverage != 1 || card.Insufficient {
		t.Fatalf("Coverage=%v Insufficient=%v, want 1/false (every component scored, spec §1.5's coverage note)",
			card.Coverage, card.Insufficient)
	}
}

// TestWorkedExampleHealerPriestHoly reproduces spec §1.6's second worked
// example: no deaths, one dropped dispellable debuff, no validated healer
// sim (Output falls back to metric_hps's own percentile directly).
func TestWorkedExampleHealerPriestHoly(t *testing.T) {
	const player = "Player-Holy"
	fight := summary.Summary{
		DurationMS: 180000, EncounterID: shazzrahEncounterID, Kill: true,
		Roster: []summary.RosterRow{
			{GUID: player, Class: "Priest", Spec: "Holy", Role: RoleHealer, HPS: 400, ActivityPct: 78, ActiveMS: 140400},
		},
		Dispels: []summary.ExchangeRow{
			{Kind: "dispel", SourceGUID: player, SpellID: 527, ExtraSpellID: 19713, ExtraSpellName: "Gehennas' Curse", Count: 1},
		},
		Mechanics: summary.MechanicsBlock{TableFound: true, Rows: []summary.MechanicRow{
			{SpellID: 19713, Name: "Gehennas' Curse", Kind: mechanics.Dispel, Dispelled: 1, Healed: 5000},
		}},
		Auras: []summary.AuraTrack{
			{TargetGUID: "Boss", SpellID: 1243, Name: "Power Word: Fortitude", Type: "BUFF", UptimeMS: 172800, Appliers: []string{player}},
		},
		Combatants: []summary.CombatantRow{
			{GUID: player, Consumables: []summary.AuraRef{{SpellID: 17627}, {SpellID: 1249513}}},
		},
	}

	table := &mechanics.Table{EncounterID: shazzrahEncounterID, Name: "Shazzrah", Mechanics: []mechanics.Mechanic{
		{SpellID: 19713, Name: "Gehennas' Curse", Kind: mechanics.Dispel},
	}}
	utilTable := &utility.Table{Spec: "priest-holy", Owned: []utility.Entry{
		{SpellID: 1243, Name: "Power Word: Fortitude", Kind: utility.Buff, Target: "ally", Verified: "test"},
	}}
	cat := &consumables.RoleCatalogue{
		Weights:       map[string]int{"flask": 30, "food": 15, "weapon_enchant": 0, "world_buffs": 20, "combat_potion": 35},
		Flask:         []consumables.Entry{{SpellID: 17627, Name: "Flask of Distilled Wisdom", Verified: "test"}},
		Food:          []consumables.Entry{{SpellID: 1249513, Name: "Nightfin Soup", Verified: "test"}},
		WeaponEnchant: nil,
		WorldBuffs:    []consumables.Entry{{SpellID: 22888, Name: "Rallying Cry", Verified: "test"}},
		CombatPotion:  consumables.PotionGroup{MaxUses: 2, Entries: []consumables.Entry{{SpellID: 17531, Name: "Major Mana Potion", Verified: "test"}}},
	}

	src := &fakePercentiles{
		band: BandTypical, bandOK: true,
		placements: map[string]fakePlacement{
			ComponentNameOutput:           {pct: 0.41, n: 100, ok: true},
			ComponentSurvivalAvoidableHit: {pct: 0.30, n: 100, ok: true}, // AvoidableHitScore = 100-30 = 70
			ComponentMechanicsDispel:      {pct: 0.78, n: 100, ok: true},
			ComponentNameUtility:          {pct: 0.85, n: 100, ok: true},
			ComponentNamePreparation:      {pct: 1.00, n: 100, ok: true},
			ComponentNameActivity:         {pct: 0.66, n: 100, ok: true},
		},
	}

	card := Score(fight, player, CuratedTables{Mechanics: table, Utility: utilTable, Consumables: cat}, src, nil, DefaultModelInfo())

	checkComponent(t, card, ComponentNameOutput, 41)
	checkComponent(t, card, ComponentNameUtility, 85)
	checkComponent(t, card, ComponentNamePreparation, 100)
	checkComponent(t, card, ComponentNameActivity, 66)

	mech := componentByName(card, ComponentNameMechanics)
	if want := 78.0; roundTo1(mech.Score) != want {
		t.Errorf("Mechanics.Score = %v, want %v (only the dispel sub-part is scored: no interrupt mechanic on this table)", mech.Score, want)
	}

	survival := componentByName(card, ComponentNameSurvival)
	if want := 89.5; roundTo1(survival.Score) != want {
		t.Errorf("Survival.Score = %v, want %v (spec: 0.65*100 + 0.35*70)", survival.Score, want)
	}

	if want := 70.0; math.Round(card.Overall) != want {
		t.Fatalf("round(Overall) = %v (Overall=%v), want %v (spec §1.6's healer worked example, \"-> 70\")",
			math.Round(card.Overall), card.Overall, want)
	}
	if card.Coverage != 1 || card.Insufficient {
		t.Fatalf("Coverage=%v Insufficient=%v, want 1/false (every component scored, spec §1.5's coverage note)",
			card.Coverage, card.Insufficient)
	}
}

// TestWorkedExampleTankWarriorProtection reproduces spec §1.6's third
// worked example: one unavoidable death near the very end plus two
// avoidable hits assigned to the other tank, and a fight where Mechanics
// has no data at all (taunted correctly, nothing to interrupt or dispel)
// and is excluded, its weight redistributed across the other five per
// §1.1. The spec's own text originally mis-stated this example's Survival
// combined score and its Utility weight in the redistribution step; both
// are corrected in the spec (see the "CORRECTION" note next to §1.6's
// tank example) and reproduced here.
func TestWorkedExampleTankWarriorProtection(t *testing.T) {
	const player = "Player-Prot"
	const otherTank = "Player-OtherTank"
	fight := summary.Summary{
		DurationMS: 180000, EncounterID: shazzrahEncounterID, Kill: true,
		Roster: []summary.RosterRow{
			{GUID: player, Class: "Warrior", Spec: "Protection", Role: RoleTank, DPS: 200, ActivityPct: 85, ActiveMS: 153000},
			{GUID: otherTank, Class: "Warrior", Spec: "Protection", Role: RoleTank},
		},
		Deaths: []summary.Death{
			{GUID: player, AtMS: 175000, KillingBlow: &summary.DamageRef{SpellID: 19999, SpellName: "Vicious Headbutt"}},
		},
		// Two avoidable hits, earlier in the fight, from a Role-tagged
		// ("tank") cleave -- assigned to the other tank below, so per the
		// whole-branch review's ruling this counts in full against player,
		// not excused: the table alone cannot say whose job a Role-tagged
		// mechanic was, only an Assignment can.
		Mechanics: summary.MechanicsBlock{TableFound: true, Rows: []summary.MechanicRow{
			{SpellID: 19998, Name: "Cleave", Kind: mechanics.Avoidable, Role: "tank",
				Players: []summary.MechanicHit{{GUID: player, Hits: 2, Damage: 4000, FirstMS: 60000, LastMS: 65000}}},
		}},
	}
	assignments := []Assignment{
		{PlayerKey: otherTank, Job: "cleave duty", FromMS: 50000, ToMS: 70000},
	}

	// No interrupt- or dispel-kind mechanic on this table at all: Mechanics
	// is excluded outright (§1.1's "no interrupt/dispel rows" trigger).
	table := &mechanics.Table{EncounterID: shazzrahEncounterID, Name: "Shazzrah", Mechanics: []mechanics.Mechanic{
		{SpellID: 19999, Name: "Vicious Headbutt", Kind: mechanics.Unavoidable},
		{SpellID: 19998, Name: "Cleave", Kind: mechanics.Avoidable, Role: "tank"},
	}}
	utilTable := &utility.Table{Spec: "warrior-protection", Owned: []utility.Entry{
		{SpellID: 7386, Name: "Sunder Armor", Kind: utility.Debuff, Target: "enemy", Verified: "test"},
	}}
	cat := &consumables.RoleCatalogue{
		Weights:       map[string]int{"flask": 30, "food": 15, "weapon_enchant": 15, "world_buffs": 20, "combat_potion": 20},
		Flask:         []consumables.Entry{{SpellID: 17626, Name: "Flask of the Titans", Verified: "test"}},
		Food:          []consumables.Entry{{SpellID: 1248395, Name: "Tender Wolf Steak", Verified: "test"}},
		WeaponEnchant: []consumables.Entry{{SpellID: 22756, Name: "Elemental Sharpening Stone", Verified: "test"}},
		WorldBuffs:    []consumables.Entry{{SpellID: 22888, Name: "Rallying Cry", Verified: "test"}},
		CombatPotion:  consumables.PotionGroup{MaxUses: 2, Entries: []consumables.Entry{{SpellID: 17528, Name: "Mighty Rage Potion", Verified: "test"}}},
	}

	src := &fakePercentiles{
		band: BandTypical, bandOK: true,
		placements: map[string]fakePlacement{
			ComponentNameOutput:           {pct: 0.55, n: 100, ok: true},
			ComponentSurvivalAvoidableHit: {pct: 0.65, n: 100, ok: true}, // AvoidableHitScore = 100-65 = 35
			ComponentNameUtility:          {pct: 0.91, n: 100, ok: true},
			ComponentNamePreparation:      {pct: 1.00, n: 100, ok: true},
			ComponentNameActivity:         {pct: 0.70, n: 100, ok: true},
		},
	}

	card := Score(fight, player, CuratedTables{Mechanics: table, Utility: utilTable, Consumables: cat}, src, assignments, DefaultModelInfo())

	// Proves the fix, not just the display number: the fake's canned 65th
	// percentile answer for the avoidable-hit sub-part would pass unchanged
	// whether the real underlying value was the excused 0 or the fully-
	// counted 22.22 -- so assert on what avoidableDamagePerSecond actually
	// computed and handed to Placement, not only on the resulting score.
	got := src.recorded[ComponentSurvivalAvoidableHit]
	if len(got) != 1 {
		t.Fatalf("Placement(survival_avoidable_hit) called %d times, want 1: %v", len(got), got)
	}
	if want := 4000.0 / 180.0; got[0] != want {
		t.Fatalf("avoidable-hit value placed = %v, want %v (4000 damage over 180s, counted in full: "+
			"the assignment names the OTHER tank for this window, not player)", got[0], want)
	}

	mech := componentByName(card, ComponentNameMechanics)
	if !mech.Excluded {
		t.Fatalf("Mechanics.Excluded = false, want true (no interrupt or dispel mechanic on this table); score=%v", mech.Score)
	}

	survival := componentByName(card, ComponentNameSurvival)
	if want := 77.06; roundTo(survival.Score, 2) != want {
		t.Errorf("Survival.Score = %v, want %v (corrected: 0.65*99.7 + 0.35*35)", survival.Score, want)
	}
	if survival.Score == 0 {
		t.Fatal("DeathScore must not be zero for an unavoidable death near the end of the fight")
	}

	if card.OverallCapped {
		t.Fatal("DeathScore is nonzero (99.7); the cap must not fire")
	}
	if want := 80.0; math.Round(card.Overall) != want {
		t.Fatalf("round(Overall) = %v (Overall=%v, OverallUncapped=%v), want %v (spec §1.6's tank worked example, corrected)",
			math.Round(card.Overall), card.Overall, card.OverallUncapped, want)
	}
	// Mechanics (weight 20 of the tank's 100) is the only excluded
	// component: coverage 0.8, well above MinCoverage -- unaffected by
	// spec §1.5's coverage ruling, per the ruling's own note that this
	// worked example and the healer's are unaffected.
	if want := 0.8; card.Coverage != want || card.Insufficient {
		t.Fatalf("Coverage=%v Insufficient=%v, want %v/false", card.Coverage, card.Insufficient, want)
	}
}

func checkComponent(t *testing.T, card Card, name string, want float64) {
	t.Helper()
	c := componentByName(card, name)
	if c.Excluded {
		t.Fatalf("%s is excluded (%s), want scored", name, c.Reason)
	}
	if roundTo1(c.Score) != want {
		t.Errorf("%s.Score = %v, want %v", name, c.Score, want)
	}
}

func componentByName(card Card, name string) Component {
	for _, c := range card.Components {
		if c.Name == name {
			return c
		}
	}
	return Component{}
}
