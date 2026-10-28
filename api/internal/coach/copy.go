package coach

import (
	"cmp"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strings"
	"testing"
)

// Every player-facing pack string the spec shows, word for word
// (docs/superpowers/specs/2026-10-27-log-coach-design.md §5, §6, §11), as format strings
// with named fields filled by Fill. The experience spec's §8.1 strings follow
// (design/specs/2026-10-28-log-coach.md); every string is here.
// verbatim below lists each product-spec string with the spec's own example values, and
// spec_test.go proves every filled example is quoted in the spec; experienceCopy lists the
// §8.1 strings by their table key, and experience_spec_test.go pins each to its row.
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

	DetailCoreCasts = "{lost} of a possible {possible} went unused while you had the {resource}. Longest gaps {gap1} and {gap2}."
	DetailNoFlask   = "You pulled with {items}, and no flask."
	DetailDeathHit  = "{Spell} hit for {amount} at {t}."
	DeathNoPotion   = "No healing potion or Healthstone used in the 10 s before."
	DetailHealRank  = "{pct}% of what {Spell} (Rank {r}) healed was overheal, over {n} casts."
	DetailUptime    = "{Aura} was up for {n} s of {duration}."
	DetailLowRank   = "Rank {next} is trainable at your level. Rank {r} was cast {n} times."

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

// The experience spec's §8.1 pack copy, word for word, by its table key.
const (
	TitleIdle    = "Idle for {n} seconds"
	DetailIdle   = "Idle outside downtime for {n} s, {pct}% of the time you were up. Longest: {windows}."
	BasisIdleTop = "Top {n} {Spec} {classPlural} on {Boss} ({size}): {ref}% active. You: {you}%."
	BasisIdle    = "From your log. Downtime and time spent dead are left out."

	TitleAdherence        = "{pct}% of casts on the rotation"
	DetailAdherence       = "Most-cast spells outside the written rotation: {spells}."
	BasisAdherenceTop     = "Top {n} {Spec} {classPlural} on {Boss} ({size}): {ref}% on the rotation. You: {you}%."
	BasisAdherenceWritten = "Measured against the written {Spec} rotation. Under {floor}% is flagged."

	TitleNoFood           = "No food at the pull"
	DetailNoFood          = "You pulled with {items} but no food."
	DetailNoFlaskBare     = "You pulled with no flask."
	DetailNoFoodBare      = "You pulled with no food."
	DetailAlsoMissing     = "Also missing: {items}."
	TitleNoWeaponEnchant  = "{Slot} has no enchant"
	DetailNoWeaponEnchant = "{Slot} had no permanent enchant at the pull."
	TitleNoEnchants       = "{n} items have no enchant"
	DetailNoEnchants      = "No permanent enchant on {slots}."
	BasisGear             = "From the gear in your log."
	TitlePotions          = "Combat potions: {n} unused"
	DetailPotions         = "You used {used} combat potions. The fight had room for {max}."
	DetailPotionsOne      = "You used {used} combat potion. The fight had room for {max}."
	TitleRaidBuff         = "No {Buff} at the pull"
	TitleRaidBuffs        = "{n} raid buffs missing"
	DetailRaidBuff        = "{Buff} was not on you at the pull. A {class} in the raid provides it."
	DetailRaidBuffs       = "Not on you at the pull: {buffs}. The raid has a class that provides each."
	BasisRaidBuff         = "From the character sheet in your log, against the classes in this raid."

	DetailCoreCastsOne   = "{lost} of a possible {possible} went unused while you had the {resource}. Longest gap {gap1}."
	DetailCoreMore       = "Also lost: {more}."
	DetailOutOfMana      = "Your mana was under 5% for {n} s in all, first from {t}."
	DetailOutOfManaDeath = "Your mana was under 5% for {n} s in all, first from {t}. You died at {death}."
	DetailCapped         = "Your {resource} sat at its cap for {n} s outside {downtime}, longest at {span}."
	DetailCappedPlain    = "Your {resource} sat at its cap for {n} s, longest at {span}."

	NotReviewableTrash = "The coach reviews boss pulls only."
	NotReviewableShort = "This pull is under 30 seconds, too short to show a rotation."
	CannotSpecRules    = "Core casts, uptimes and the rotation: the coach doesn't know the {Spec} rotation yet."

	CannotLevelRanks = "Spell ranks: this log doesn't tell us your level."
	CannotLevelSim   = "Our guide's simulated rotation: this log doesn't tell us your level."

	BasisFromLog             = "From your log."
	DetailDiedWithReadySince = "{Killer} killed you at {t}. {Spell} had been ready since {since}."
	DetailAlsoReady          = "Also ready: {spells}."
	KillerMelee              = "{Mob}'s melee"
	KillerPlayerMelee        = "another player's melee"
	KillerPet                = "a pet"

	TitleDiedNotUsed           = "Died at {t}, {Spell} not used"
	DetailDiedNotUsed          = "{Killer} killed you at {t}. {Spell} was not used this pull."
	DetailDiedNotUsedAvoidable = "{Killer} killed you at {t}, and this boss lets you avoid it. {Spell} was not used this pull."

	StrengthCasts     = "{Spell}: {you} casts a minute. Top {n}: {ref}."
	StrengthCastsSim  = "{Spell}: {you} casts a minute. Our guide's rotation, simulated: {ref}."
	StrengthAdherence = "{pct}% of your casts were on the written rotation. Top {n}: {ref}%."
	StrengthOverheal  = "{Spell} (Rank {r}): {pct}% overheal. Top {n}: {ref}%."

	LinkDeathInGapSentence = "The {Spell} gap at {gap} includes your death at {t}."
	LinkNoFlaskMana        = "You pulled without a flask and your mana ran out from {t}."
	LinkPotionDeath        = "You died at {t} with combat potions unused."

	// The words the rules were drafted with, fixed by §8.1 once the rules existed.
	PoolLacksMeasure             = "Top {Spec} logs on {Boss} ({size}) don't record this."
	TitleNoWeaponOil             = "No weapon oil or stone"
	DetailNoWeaponOil            = "You pulled with no weapon oil or stone."
	CannotWeaponEnchant          = "Weapon oils and stones: the coach can't read them from your log yet."
	StrengthConsumables          = "Every consumable the coach can check was up at the pull."
	DetailDiedWithReadyAvoidable = "{Killer} killed you at {t}, and this boss lets you avoid it. {Spell} had been ready since {since}."
	CannotDefensives             = "Defensive cooldowns: the coach doesn't know your spec's list yet."
	CannotCoreList               = "Core casts: the coach doesn't know your spec's list yet."
	CannotAuraList               = "Uptimes: the coach doesn't know your spec's list yet."
	CannotTalentAura             = "{Aura} uptime: it comes from a talent the log can't show, and it was never up this pull."
	CannotOthersDebuff           = "{Aura} uptime: another player applied it this pull, not you."
	CannotDefensiveLevel         = "Defensive cooldowns you never cast: this log doesn't tell us your level."
	CannotDeathKiller            = "Your death at {t}: the log doesn't say what killed you."
	basisUptimeValue             = "{ref}% up. You: {you}%."
	detailUptimeMore             = "Also low: {others}."
	uptimeMoreItem               = "{Aura} {pct}%"
	coreMoreEntry                = "{Spell} {n}"
	adherenceEntry               = "{Spell} ({n})"
	prepMissingFlask             = "a flask"
	prepMissingFood              = "food"
	prepMissingOil               = "a weapon oil or stone"
	prepMissingEnchant           = "a {Slot} enchant"
	prepMissingPotions           = "combat potions"
	prepMissingBlessing          = "a blessing"
	prepIncomplete               = "{Category}: the coach's {list} list isn't complete yet."
	coreUnknownResource          = "resource"
	blessingWord                 = "blessing"

	StemUptime    = "{Aura} uptime"
	StemIdle      = "Idle time"
	StemCapped    = "{Resource} at the cap"
	StemOutOfMana = "Running out of mana"
	StemLowRank   = "{Spell} below your best rank"
	StemHealRank  = "{Spell} overheal"
	StemNoFlask   = "No flask at the pull"
	StemNoFood    = "No food at the pull"
	StemEnchants  = "Missing enchants"
	StemPotions   = "Unused combat potions"
	StemRaidBuffs = "Missing raid buffs"
	StemDeath     = "Early deaths"
	StemAdherence = "Casts off the rotation"
)

// The §8.1 rows that are lists, in the order the table gives them: the in-game slot names
// by gear index, and the resource nouns by power id. experienceLists pins each.
var (
	enchantSlotNames = map[int]string{4: "Chest", 7: "Feet", 14: "Back", mainHandSlot: "Main Hand"}
	experienceLists  = map[string]func() []string{
		"enchantSlotNames": func() []string { return orderedValues(enchantSlotNames) },
		"powerNames":       func() []string { return orderedValues(PowerNames) },
	}
)

// orderedValues is m's values in key order.
func orderedValues[K cmp.Ordered](m map[K]string) []string {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	out := make([]string, len(keys))
	for i, k := range keys {
		out[i] = m[k]
	}
	return out
}

// experienceCopy is every §8.1 string by the key in the table's first column.
var experienceCopy = map[string]string{
	"TitleIdle": TitleIdle, "DetailIdle": DetailIdle, "BasisIdleTop": BasisIdleTop, "BasisIdle": BasisIdle,
	"TitleAdherence": TitleAdherence, "DetailAdherence": DetailAdherence,
	"BasisAdherenceTop": BasisAdherenceTop, "BasisAdherenceWritten": BasisAdherenceWritten,
	"TitleNoFood": TitleNoFood, "DetailNoFood": DetailNoFood, "DetailNoFlaskBare": DetailNoFlaskBare,
	"DetailNoFoodBare": DetailNoFoodBare, "DetailAlsoMissing": DetailAlsoMissing,
	"TitleNoWeaponEnchant": TitleNoWeaponEnchant, "DetailNoWeaponEnchant": DetailNoWeaponEnchant,
	"TitleNoEnchants": TitleNoEnchants, "DetailNoEnchants": DetailNoEnchants, "BasisGear": BasisGear,
	"TitlePotions": TitlePotions, "DetailPotions": DetailPotions,
	"TitleRaidBuff": TitleRaidBuff, "TitleRaidBuffs": TitleRaidBuffs, "DetailRaidBuff": DetailRaidBuff,
	"DetailRaidBuffs": DetailRaidBuffs, "BasisRaidBuff": BasisRaidBuff,
	"DetailCoreCastsOne": DetailCoreCastsOne, "DetailCoreCasts": DetailCoreCasts, "DetailCoreMore": DetailCoreMore,
	"DetailOutOfMana": DetailOutOfMana, "DetailCappedPlain": DetailCappedPlain,
	"NotReviewableTrash": NotReviewableTrash, "NotReviewableShort": NotReviewableShort, "CannotSpecRules": CannotSpecRules,
	"CannotLevelRanks": CannotLevelRanks, "CannotLevelSim": CannotLevelSim,
	"BasisFromLog": BasisFromLog, "DetailDiedWithReadySince": DetailDiedWithReadySince,
	"DetailAlsoReady": DetailAlsoReady, "KillerMelee": KillerMelee,
	"KillerPlayerMelee": KillerPlayerMelee, "KillerPet": KillerPet, "TitleDiedNotUsed": TitleDiedNotUsed,
	"DetailDiedNotUsed": DetailDiedNotUsed, "DetailDiedNotUsedAvoidable": DetailDiedNotUsedAvoidable,
	"StrengthCasts": StrengthCasts, "StrengthCastsSim": StrengthCastsSim,
	"StrengthAdherence": StrengthAdherence, "StrengthOverheal": StrengthOverheal,
	"LinkDeathInGap": LinkDeathInGapSentence, "LinkNoFlaskMana": LinkNoFlaskMana, "LinkPotionDeath": LinkPotionDeath,
	"StemUptime": StemUptime, "StemIdle": StemIdle, "StemCapped": StemCapped, "StemOutOfMana": StemOutOfMana,
	"StemLowRank": StemLowRank, "StemHealRank": StemHealRank, "StemNoFlask": StemNoFlask,
	"StemNoFood": StemNoFood, "StemEnchants": StemEnchants, "StemPotions": StemPotions,
	"StemRaidBuffs": StemRaidBuffs, "StemDeath": StemDeath, "StemAdherence": StemAdherence,
	"DetailCapped": DetailCapped, "BasisCapOwnBar": BasisCapOwnBar, "DetailOutOfManaDeath": DetailOutOfManaDeath,
	"DetailPotionsOne":   DetailPotionsOne,
	"pool_lacks_measure": PoolLacksMeasure, "title_no_weapon_oil": TitleNoWeaponOil,
	"detail_no_weapon_oil": DetailNoWeaponOil, "cannot_weapon_enchant": CannotWeaponEnchant,
	"strength_consumables": StrengthConsumables, "detail_died_with_ready_avoidable": DetailDiedWithReadyAvoidable,
	"cannot_defensives": CannotDefensives, "cannot_defensive_level": CannotDefensiveLevel,
	"CannotCoreList": CannotCoreList, "CannotAuraList": CannotAuraList,
	"CannotTalentAura": CannotTalentAura, "CannotOthersDebuff": CannotOthersDebuff,
	"cannot_death_killer": CannotDeathKiller,
	"basisUptimeValue":    basisUptimeValue, "detailUptimeMore": detailUptimeMore, "uptimeMoreItem": uptimeMoreItem,
	"coreMoreEntry": coreMoreEntry, "adherenceEntry": adherenceEntry,
	"prepMissingFlask": prepMissingFlask, "prepMissingFood": prepMissingFood, "prepMissingOil": prepMissingOil,
	"prepMissingEnchant": prepMissingEnchant, "prepMissingPotions": prepMissingPotions,
	"prepMissingBlessing": prepMissingBlessing, "prepIncomplete": prepIncomplete,
	"coreUnknownResource": coreUnknownResource, "blessingWord": blessingWord,
}

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
	"DetailCoreCasts":      {DetailCoreCasts, map[string]string{"lost": "9", "possible": "40", "resource": "rage", "gap1": "1:12–1:25", "gap2": "2:40–2:51"}},
	"DetailNoFlask":        {DetailNoFlask, map[string]string{"items": "Grilled Squid and an Elemental Sharpening Stone"}},
	"DetailDeathHit":       {DetailDeathHit, map[string]string{"Spell": "Flame Breath", "amount": "2,140", "t": "3:20"}},
	"DeathNoPotion":        {DeathNoPotion, nil},
	"DetailHealRank":       {DetailHealRank, map[string]string{"pct": "44", "Spell": "Greater Heal", "r": "4", "n": "23"}},
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

// joinAnd is "a", "a and b", "a, b and c": the copy's list form.
func joinAnd(items []string) string {
	if len(items) < 2 {
		return strings.Join(items, "")
	}
	return strings.Join(items[:len(items)-1], ", ") + " and " + items[len(items)-1]
}

// The panel header the fight route sends as persona.label (§5.1). These are web copy, the
// experience spec's §8.2 keys titleModel, titleStarter and titleRules, not §8.1 pack copy,
// so they stay out of the §8.1 registry; a one-way test pins them to the §8.2 rows. The web
// ignores the label on an officer view.
const (
	PersonaLabelModel   = "Your {Spec} coach"
	PersonaLabelStarter = "Your {Spec} coach · early version"
	PersonaLabelRules   = "Your {Spec} review"
)

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
