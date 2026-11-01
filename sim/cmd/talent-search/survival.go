package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

// A talent is survival-only when the client's own data says it can add
// no damage. The test reads two files of the build:
//
//   - The talent's own spellconst rows (one per rank) are nothing but
//     spell modifiers (aura 107 or 108). It grants no stat, proc or
//     damage of its own, only changes to other spells.
//   - Its top-rank tooltip (talents/<class>.json) names at least one of
//     the class's spells, and none of the spells it names contributes to
//     damage in spellconst (contributesDamage). A tooltip that names
//     another talent, or no spell at all, cannot be judged and fails.
//
// For warlocks: Improved Health Funnel (Health Funnel), Demonic Aegis
// (Demon Skin, Demon Armor) and Improved Voidwalker (Torment, Consume
// Shadows, Sacrifice, Suffering) are survival-only. Demonic Brand, which
// names Searing Pain, is not; nor is Molten Skin, which names no spell.
//
// A talent whose rows are only proc triggers is survival-only on the same
// tooltip test when the tooltip also names no gain: Improved Concussive
// Shot, a stun on Concussive Shot (a snare), which Survival takes only to
// open the Marksmanship tree's second row.
//
// The adoption rule keeps every unmodeled guide talent because the
// engine may simply be missing its damage. A survival-only talent has
// no damage to miss, so a candidate may re-spend it.

// spellconstEffect is the part of one spellconst effect the test reads.
type spellconstEffect struct {
	Effect int     `json:"effect"`
	Aura   int     `json:"aura"`
	Amount float64 `json:"amount"`
}

// spellconstSpell is one spellconst/<class>.json row.
type spellconstSpell struct {
	Name    string             `json:"name"`
	Effects []spellconstEffect `json:"effects"`
}

// The client's SpellEffect numbers this test reads.
const (
	effectDummy                 = 3
	effectSchoolDamage          = 2
	effectHealthLeech           = 9
	effectHeal                  = 10
	effectWeaponDamage          = 17
	effectSummon                = 28
	effectEnergize              = 30
	effectWeaponDamageNoSchool  = 31
	effectSummonPet             = 56
	effectWeaponPercentDamage   = 58
	effectTriggerSpell          = 64
	effectScriptEffect          = 77
	effectSummonObjectSlot1     = 87
	effectSummonObjectSlot2     = 88
	effectSummonObjectSlot3     = 89
	effectSummonObjectSlot4     = 90
	effectNormalizedWeaponDmg   = 121
	effectTriggerSpellWithValue = 142
	effectDamageFromMaxHealthPC = 165
	effectApplyAura             = 6
)

// The client's aura numbers this test reads.
const (
	auraDummy             = 4
	auraPeriodicDamage    = 3
	auraPeriodicHeal      = 8
	auraModDamageDone     = 13
	auraModResistance     = 22
	auraModArmor          = 674 // Forever's armor aura (Demon Armor, Sunder Armor)
	auraPeriodicTrigger   = 23
	auraPeriodicEnergize  = 24
	auraProcTrigger       = 42
	auraModCritPercent    = 52
	auraPeriodicLeech     = 53
	auraModSpellCrit      = 57
	auraPeriodicManaLeech = 64
	auraModCastingSpeed   = 65
	auraModSchoolCrit     = 71
	auraModDamagePercent  = 79
	auraPeriodicDamagePct = 89
	auraModAttackPower    = 99
	auraModRangedAP       = 124
	auraModMeleeHaste     = 138
	auraModRangedHaste    = 140
	auraModHaste          = 192
	auraAddFlatModifier   = 107
	auraAddPctModifier    = 108
	auraAddTargetTrigger  = 109
)

// damageEffects are effects that deal damage, grant a resource, or hand
// off to something whose damage this file cannot see (a script, a
// triggered spell, a summoned pet or totem).
var damageEffects = map[int]bool{
	effectDummy: true, effectSchoolDamage: true, effectHealthLeech: true, effectWeaponDamage: true,
	effectSummon: true, effectEnergize: true, effectWeaponDamageNoSchool: true, effectSummonPet: true,
	effectWeaponPercentDamage: true, effectTriggerSpell: true, effectScriptEffect: true,
	effectSummonObjectSlot1: true, effectSummonObjectSlot2: true, effectSummonObjectSlot3: true,
	effectSummonObjectSlot4: true, effectNormalizedWeaponDmg: true, effectTriggerSpellWithValue: true,
	effectDamageFromMaxHealthPC: true,
}

// damageAuras are auras that deal damage, raise damage, crit, haste or
// attack power, grant a resource, or trigger something unseen.
var damageAuras = map[int]bool{
	auraDummy: true, auraModDamageDone: true, auraPeriodicTrigger: true, auraPeriodicEnergize: true,
	auraProcTrigger: true, auraAddTargetTrigger: true,
	auraModCritPercent: true, auraPeriodicLeech: true, auraModSpellCrit: true, auraPeriodicManaLeech: true,
	auraModCastingSpeed: true, auraModSchoolCrit: true, auraModDamagePercent: true, auraPeriodicDamagePct: true,
	auraModAttackPower: true, auraModRangedAP: true, auraModMeleeHaste: true, auraModRangedHaste: true,
	auraModHaste: true,
}

// contributesDamage reports whether a spell's effects can add to
// damage. A periodic-damage aura on a spell that also heals is Health
// Funnel's shape (the caster's own health paid into the demon), not
// damage on a target.
func contributesDamage(effects []spellconstEffect) bool {
	heals := false
	for _, e := range effects {
		if e.Effect == effectHeal || (e.Effect == effectApplyAura && e.Aura == auraPeriodicHeal) {
			heals = true
		}
	}
	for _, e := range effects {
		if damageEffects[e.Effect] {
			return true
		}
		if e.Effect != effectApplyAura {
			continue
		}
		if damageAuras[e.Aura] || (e.Aura == auraPeriodicDamage && !heals) {
			return true
		}
		// A negative armor or resistance aura is a debuff on the target
		// (Sunder Armor), which raises damage taken.
		if (e.Aura == auraModResistance || e.Aura == auraModArmor) && e.Amount < 0 {
			return true
		}
	}
	return false
}

// onlyProcTriggers reports whether a talent's own row is nothing but proc
// triggers (Improved Concussive Shot's stun on Concussive Shot). What the
// proc does is read from the tooltip (gainWords).
func onlyProcTriggers(effects []spellconstEffect) bool {
	if len(effects) == 0 {
		return false
	}
	for _, e := range effects {
		if e.Effect != effectApplyAura || e.Aura != auraProcTrigger {
			return false
		}
	}
	return true
}

// gainWords are what a proc's tooltip says when it can raise damage or a
// resource; a crowd-control proc ("a 20% chance to stun the target for 3
// sec") says none of them.
var gainWords = regexp.MustCompile(`(?i)\b(damage|attack|critical|haste|speed|power|mana|energy|rage|focus|health|heal\w*|armor|resist\w*|hit|strength|agility|intellect|spirit|stamina|threat)\b`)

// onlyModifiers reports whether a talent's own row grants nothing but
// spell modifiers.
func onlyModifiers(effects []spellconstEffect) bool {
	if len(effects) == 0 {
		return false
	}
	for _, e := range effects {
		if e.Effect != effectApplyAura || (e.Aura != auraAddFlatModifier && e.Aura != auraAddPctModifier) {
			return false
		}
	}
	return true
}

// spellDamage is, per name a tooltip can mention, whether anything of
// that name can add to damage. Every talent name of the class counts as
// damage: a tooltip naming another talent ("Your Demonic Sacrifice
// effect...") improves something spellconst cannot vouch for.
type spellDamage map[string]bool

func newSpellDamage(spells map[string]spellconstSpell, talentNames []string) spellDamage {
	out := spellDamage{}
	for _, s := range spells {
		if s.Name != "" {
			out[s.Name] = out[s.Name] || contributesDamage(s.Effects)
		}
	}
	for _, name := range talentNames {
		out[name] = true
	}
	return out
}

// namedSpells is the spells and talents a tooltip names, longest names
// first so "Demon Armor" is not also read as a shorter name inside it
// and "Bane of Havoc" not as "Bane". The talent's own name is consumed
// first and never counted.
func (d spellDamage) namedSpells(tooltip, self string) []string {
	names := make([]string, 0, len(d))
	for name := range d {
		if name != self {
			names = append(names, name)
		}
	}
	sort.Slice(names, func(i, j int) bool {
		if len(names[i]) != len(names[j]) {
			return len(names[i]) > len(names[j])
		}
		return names[i] < names[j]
	})
	rest := tooltip
	if self != "" {
		rest = wholeWord(self).ReplaceAllString(rest, " ")
	}
	var out []string
	for _, name := range names {
		re := wholeWord(name)
		if re.MatchString(rest) {
			out = append(out, name)
			rest = re.ReplaceAllString(rest, " ")
		}
	}
	sort.Strings(out)
	return out
}

func wholeWord(name string) *regexp.Regexp {
	return regexp.MustCompile(`\b` + regexp.QuoteMeta(name) + `\b`)
}

// namesOnlySafeSpells is the tooltip half of the rule.
func (d spellDamage) namesOnlySafeSpells(tooltip, self string) bool {
	named := d.namedSpells(tooltip, self)
	if len(named) == 0 {
		return false
	}
	for _, name := range named {
		if d[name] {
			return false
		}
	}
	return true
}

// tooltipTalent is one talent of talents/<class>.json: its name, the
// spell ids of its ranks and its top rank's tooltip.
type tooltipTalent struct {
	name     string
	spellIDs []int
	tooltip  string
}

func loadTooltipTalents(buildDir, class string) (map[int]tooltipTalent, error) {
	path := filepath.Join(buildDir, "talents", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var f struct {
		Trees []struct {
			Talents []struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Ranks []struct {
					SpellID     int    `json:"spell_id"`
					Description string `json:"description"`
				} `json:"ranks"`
			} `json:"talents"`
		} `json:"trees"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	out := map[int]tooltipTalent{}
	for _, tree := range f.Trees {
		for _, n := range tree.Talents {
			if len(n.Ranks) == 0 {
				continue
			}
			t := tooltipTalent{name: n.Name, tooltip: n.Ranks[len(n.Ranks)-1].Description}
			for _, r := range n.Ranks {
				t.spellIDs = append(t.spellIDs, r.SpellID)
			}
			out[n.ID] = t
		}
	}
	return out, nil
}

func loadSpellconst(buildDir, class string) (map[string]spellconstSpell, error) {
	path := filepath.Join(buildDir, "spellconst", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var f struct {
		Spells map[string]spellconstSpell `json:"spells"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	return f.Spells, nil
}

// survivalRule is the rule above over one class's data in one build.
type survivalRule struct {
	spells map[string]spellconstSpell
	damage spellDamage
}

func newSurvivalRule(spells map[string]spellconstSpell, talents map[int]tooltipTalent) survivalRule {
	names := make([]string, 0, len(talents))
	for _, t := range talents {
		names = append(names, t.name)
	}
	return survivalRule{spells: spells, damage: newSpellDamage(spells, names)}
}

// survivalOnly applies both halves of the rule to one talent. A rank
// whose row spellconst does not carry fails: nothing vouches for it. Every
// rank is either all spell modifiers or all proc triggers; a proc talent's
// tooltip must also name no gain (gainWords).
func (r survivalRule) survivalOnly(t tooltipTalent) bool {
	if len(t.spellIDs) == 0 {
		return false
	}
	procs := false
	for _, id := range t.spellIDs {
		row, ok := r.spells[strconv.Itoa(id)]
		if !ok {
			return false
		}
		switch {
		case onlyModifiers(row.Effects):
		case onlyProcTriggers(row.Effects):
			procs = true
		default:
			return false
		}
	}
	if procs && gainWords.MatchString(t.tooltip) {
		return false
	}
	return r.damage.namesOnlySafeSpells(t.tooltip, t.name)
}

// survivalOnlyTalents is the class's survival-only talent ids in one
// client build.
func survivalOnlyTalents(buildDir, class string) (map[int]bool, error) {
	talents, err := loadTooltipTalents(buildDir, class)
	if err != nil {
		return nil, err
	}
	spells, err := loadSpellconst(buildDir, class)
	if err != nil {
		return nil, err
	}
	rule := newSurvivalRule(spells, talents)
	out := map[int]bool{}
	for id, t := range talents {
		if rule.survivalOnly(t) {
			out[id] = true
		}
	}
	return out, nil
}
