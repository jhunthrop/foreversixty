package coach

import (
	"fmt"
	"log/slog"
	"regexp"
	"testing"
)

// Every player-facing pack string the spec shows, word for word
// (docs/superpowers/specs/2026-10-27-log-coach-design.md §5, §6, §11), as format strings
// with named fields filled by Fill. Strings the spec does not show are in copy_drafts.go.
// verbatim below lists each one with the spec's own example values, and spec_test.go
// proves every filled example is quoted in the spec.
const (
	NotReviewableSpec = "We couldn't tell this player's spec from the log."
	CannotTalents     = "Talents and race: Forever logs don't record them."
	CannotRange       = "Range and movement: the log doesn't say where you stood."
	CoreNoCost        = "Couldn't check whether you had the rage."
	SharedDebuff      = "Your share of {Aura}: {n} {classPlural} applied it, and the log merges them."
	FoodIncomplete    = "Food: the coach's food list isn't complete yet."
	NoSheet           = "Consumables and enchants: your log has no character sheet for this pull."
	ThinPool          = "not enough logs yet — {n} {Spec} kills of {Boss} ({size}) on the site, the coach needs 20."

	// BasisTop's template differs from §11's `Top {n} {Spec} {class} on {Boss} ({size}): …`
	// in field names only; the pin is §5.3's filled form.
	BasisTop       = "Top {n} {Spec} {classPlural} on {Boss} ({size}): "
	BasisTopHere   = "Top {n} {Spec} {classPlural} here: "
	BasisSimLabel  = "The written {Spec} rotation, simulated at level {band}: "
	BasisCooldown  = "Measured against {Spell}'s {n}-second cooldown and your own {resource} bar."
	BasisPerMin    = "{ref} a minute. You: {you}."
	BasisSameRank  = "{pct}% on the same rank."
	BasisOwnBar    = "Measured against your own {resource} bar."
	BasisCapOwnBar = "Measured against your own {resource} bar, {downtime} left out."
	BasisSheet     = "From the character sheet in your log, against the site's consumable list for {roles}."
	BasisAvoidable = "From your log. {Spell} is on this boss's list of damage you can avoid."
	BasisDeathCD   = "From your log and {Spell}'s cooldown."
	BasisRanks     = "The spell ranks a level {level} {class} can train."
	FromYourLog    = "From your log"
	LowerRanks     = "Lower ranks cost less mana."

	RulesOpening = "Here is what your log shows for {Boss}."
	RulesNothing = "Nothing to fix on this pull that the log can show."

	TitleCoreCasts     = "{Spell}: {n} casts lost"
	TitleNoFlask       = "No flask at the pull"
	TitleCapped        = "{Resource} full for {n} seconds"
	TitleDiedTo        = "Died at {t} to {Spell}"
	TitleDiedWithReady = "Died at {t} with {Spell} ready"
	TitleOutOfMana     = "Out of mana from {t}"
	TitleHealRank      = "{Spell} (Rank {r}): {pct}% overheal"
	TitleUptime        = "{Aura} up {pct}% of the fight"
	TitleLowRank       = "{Spell}: {n} casts of Rank {r}"

	DetailCoreCasts      = "{actual} of a possible {possible} with {resource} in your bar. Longest gaps {gap1} and {gap2}."
	DetailNoFlask        = "You pulled with {items}, and no flask."
	DetailCapped         = "Your {resource} sat at its cap for {n} s outside the {downtime}, longest at {span}."
	DetailDeathHit       = "{Spell} hit for {amount} at {t}."
	DeathNoPotion        = "No healing potion or Healthstone used in the 10 s before."
	DetailOutOfManaDeath = "Your mana stayed under 5% from {t} until you died at {death}, {n} s in all."
	DetailHealRank       = "{pct}% of what {Spell} (Rank {r}) healed was overheal, over {n} casts."
	DetailDiedWithReady  = "{Killer} killed you at {t}. {Spell} had been ready since the pull."
	DetailUptime         = "{Aura} was up for {n} s of {duration}."
	DetailLowRank        = "Rank {next} is trainable at your level. Rank {r} was cast {n} times."

	StrengthAlive     = "Alive the whole fight."
	StrengthUptime    = "{Aura} up {pct}% of the fight. Top {n}: {ref}%."
	StrengthUptimeOwn = "{Aura} up {pct}% of the fight."
	StrengthStacks    = "{Aura} at {n} stacks from {t} to the end."
	StrengthAtPull    = "{Item} at the pull."

	LinkGapCap           = "The {Spell} gap at {gap} overlaps the {resource} cap at {cap}."
	LinkOutOfManaDeath   = "You died while out of mana."
	LinkOverhealMostCast = "{Spell} (Rank {r}) is your most-cast heal with the most overheal."

	NightCoreCasts = "{Spell} casts lost"
)

// verbatimEntry is one spec-shown string and the fields that fill it into the form the
// spec quotes. No fields: the spec quotes the format itself (a constant, or a template
// the spec shows only as a template: ThinPool, BasisSimLabel, RulesOpening).
type verbatimEntry struct {
	format  string
	example map[string]string
}

// verbatim is every copy.go string the spec shows, with §5.3's own example values where
// the spec shows the filled form. Those values are the spec's invented examples, used
// only to prove the wording.
var verbatim = map[string]verbatimEntry{
	"NotReviewableSpec":    {NotReviewableSpec, nil},
	"CannotTalents":        {CannotTalents, nil},
	"CannotRange":          {CannotRange, nil},
	"CoreNoCost":           {CoreNoCost, nil},
	"SharedDebuff":         {SharedDebuff, map[string]string{"Aura": "Sunder Armor", "n": "3", "classPlural": "warriors"}},
	"FoodIncomplete":       {FoodIncomplete, nil},
	"NoSheet":              {NoSheet, nil},
	"ThinPool":             {ThinPool, nil},
	"BasisTop":             {BasisTop, map[string]string{"n": "25", "Spec": "Fury", "classPlural": "warriors", "Boss": "Onyxia", "size": "40"}},
	"BasisTopHere":         {BasisTopHere, map[string]string{"n": "25", "Spec": "Fury", "classPlural": "warriors"}},
	"BasisSimLabel":        {BasisSimLabel, nil},
	"BasisCooldown":        {BasisCooldown, map[string]string{"Spell": "Bloodthirst", "n": "6", "resource": "rage"}},
	"BasisPerMin":          {BasisPerMin, map[string]string{"ref": "9.1", "you": "8.0"}},
	"BasisSameRank":        {BasisSameRank, map[string]string{"pct": "22"}},
	"BasisOwnBar":          {BasisOwnBar, map[string]string{"resource": "mana"}},
	"BasisCapOwnBar":       {BasisCapOwnBar, map[string]string{"resource": "rage", "downtime": "air phase"}},
	"BasisSheet":           {BasisSheet, map[string]string{"roles": "damage dealers"}},
	"BasisAvoidable":       {BasisAvoidable, map[string]string{"Spell": "Flame Breath"}},
	"BasisDeathCD":         {BasisDeathCD, map[string]string{"Spell": "Shield Wall"}},
	"BasisRanks":           {BasisRanks, map[string]string{"level": "60", "class": "warrior"}},
	"FromYourLog":          {FromYourLog, nil},
	"LowerRanks":           {LowerRanks, nil},
	"RulesOpening":         {RulesOpening, nil},
	"RulesNothing":         {RulesNothing, nil},
	"TitleCoreCasts":       {TitleCoreCasts, map[string]string{"Spell": "Bloodthirst", "n": "9"}},
	"TitleNoFlask":         {TitleNoFlask, nil},
	"TitleCapped":          {TitleCapped, map[string]string{"Resource": "Rage", "n": "14"}},
	"TitleDiedTo":          {TitleDiedTo, map[string]string{"t": "3:20", "Spell": "Flame Breath"}},
	"TitleDiedWithReady":   {TitleDiedWithReady, map[string]string{"t": "2:50", "Spell": "Shield Wall"}},
	"TitleOutOfMana":       {TitleOutOfMana, map[string]string{"t": "3:05"}},
	"TitleHealRank":        {TitleHealRank, map[string]string{"Spell": "Greater Heal", "r": "4", "pct": "44"}},
	"TitleUptime":          {TitleUptime, map[string]string{"Aura": "Shield Block", "pct": "18"}},
	"TitleLowRank":         {TitleLowRank, map[string]string{"Spell": "Sunder Armor", "n": "22", "r": "4"}},
	"DetailCoreCasts":      {DetailCoreCasts, map[string]string{"actual": "31", "possible": "40", "resource": "rage", "gap1": "1:12–1:25", "gap2": "2:40–2:51"}},
	"DetailNoFlask":        {DetailNoFlask, map[string]string{"items": "Grilled Squid and an Elemental Sharpening Stone"}},
	"DetailCapped":         {DetailCapped, map[string]string{"resource": "rage", "n": "14", "downtime": "air phase", "span": "2:41–2:50"}},
	"DetailDeathHit":       {DetailDeathHit, map[string]string{"Spell": "Flame Breath", "amount": "2,140", "t": "3:20"}},
	"DeathNoPotion":        {DeathNoPotion, nil},
	"DetailOutOfManaDeath": {DetailOutOfManaDeath, map[string]string{"t": "3:05", "death": "3:20", "n": "15"}},
	"DetailHealRank":       {DetailHealRank, map[string]string{"pct": "44", "Spell": "Greater Heal", "r": "4", "n": "23"}},
	"DetailDiedWithReady":  {DetailDiedWithReady, map[string]string{"Killer": "Onyxia's melee", "t": "2:50", "Spell": "Shield Wall"}},
	"DetailUptime":         {DetailUptime, map[string]string{"Aura": "Shield Block", "n": "42", "duration": "3:52"}},
	"DetailLowRank":        {DetailLowRank, map[string]string{"next": "5", "r": "4", "n": "22"}},
	"StrengthAlive":        {StrengthAlive, nil},
	"StrengthUptime":       {StrengthUptime, map[string]string{"Aura": "Battle Shout", "pct": "100", "n": "25", "ref": "97"}},
	"StrengthUptimeOwn":    {StrengthUptimeOwn, map[string]string{"Aura": "Inner Fire", "pct": "100"}},
	"StrengthStacks":       {StrengthStacks, map[string]string{"Aura": "Sunder Armor", "n": "5", "t": "0:21"}},
	"StrengthAtPull":       {StrengthAtPull, map[string]string{"Item": "Flask of Distilled Wisdom"}},
	"LinkGapCap":           {LinkGapCap, map[string]string{"Spell": "Bloodthirst", "gap": "2:40–2:51", "resource": "rage", "cap": "2:41–2:50"}},
	"LinkOutOfManaDeath":   {LinkOutOfManaDeath, nil},
	"LinkOverhealMostCast": {LinkOverhealMostCast, map[string]string{"Spell": "Greater Heal", "r": "4"}},
	"NightCoreCasts":       {NightCoreCasts, map[string]string{"Spell": "Bloodthirst"}},
}

// SpecVerbatim is every spec-shown copy string by constant name, in the form the spec
// quotes it: filled with the spec's example values, or the raw template.
func SpecVerbatim() map[string]string {
	out := make(map[string]string, len(verbatim))
	for name, e := range verbatim {
		if e.example == nil {
			out[name] = e.format
			continue
		}
		out[name] = Fill(e.format, e.example)
	}
	return out
}

var placeholder = regexp.MustCompile(`\{([A-Za-z0-9]+)\}`)

// Fill replaces every {name} in format with fields[name]. A field the format names but
// fields lacks is a defect: it panics under go test and is logged in production, where
// the placeholder is left in place for the validator to reject.
func Fill(format string, fields map[string]string) string {
	return placeholder.ReplaceAllStringFunc(format, func(m string) string {
		name := m[1 : len(m)-1]
		if v, ok := fields[name]; ok {
			return v
		}
		if testing.Testing() {
			panic(fmt.Sprintf("coach.Fill: %q has no field %q", format, name))
		}
		slog.Default().Error("coach", "op", "fill", "format", format, "missing", name)
		return m
	})
}
