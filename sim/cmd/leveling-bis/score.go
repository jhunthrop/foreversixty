package main

// attackPowerPerDPS is the engine's DefaultAttackPowerPerDPS
// (wowsims-forever's sim/core/constants.go, pinned at 14.0): a
// weapon's damage-per-second converts to the attack-power stat it is
// worth at this fixed rate, for both melee and ranged weapons (the
// engine carries one constant, not a separate ranged one - see
// sim/core/attack.go, which threads the same AttackPowerPerDPS
// through the main hand, off hand and ranged weapon options).
//
// score() uses it rather than a weights-mode-reported weapon weight
// because this engine build's stat-weights result (sim/adapter.Weights)
// reports only the named stats a WeightsSpec asked for - there is no
// separate "weapon dps" pseudo-stat in the result the way
// PseudoStatRangedDps exists as an engine-internal stat but is never
// surfaced through sim/api.StatWeight. This lane's brief allows either
// source; this one is the one that exists.
const attackPowerPerDPS = 14.0

// weaponAPStat is which weight_stats id a weapon's DPS converts into,
// by which planner slot it occupies: a ranged weapon's DPS is worth
// ranged attack power, a main/off hand weapon's is worth (melee)
// attack power. Slots with no weapon meaning (rings, armor) are not
// listed; score() only consults this for a slot in the map.
var weaponAPStat = map[string]string{
	"ranged":    "ranged_attack_power",
	"main_hand": "attack_power",
	"off_hand":  "attack_power",
}

// genericAttackPowerWeightStats is which weight ids a flat "attack_power"
// entry in a candidate's Stats counts toward. Classic Era's generic
// "+X Attack Power" raises a character's melee AND ranged attack power at
// once - data/builds/<build>/items/<class>.json states it as one
// "attack_power" line (data/pipeline/normalize/gear.py never synthesises a
// second "ranged_attack_power" line there, so tooltips do not show a
// duplicate), so the ranker has to apply it to both weights itself rather
// than finding a "ranged_attack_power" key on the item.
//
// This mirrors data/pipeline/simdb/statmap.py's stat_array, which does the
// same for the engine's own simdb.bin (the actual DPS the sim reports for a
// hunter already counts this AP as ranged attack power; the ranker's
// weighted-sum estimate has to agree or it ranks gear the sim itself would
// not).
var genericAttackPowerWeightStats = []string{"attack_power", "ranged_attack_power"}

// statWeight is the weight an item's stat amount is multiplied by. Every
// stat but "attack_power" uses its own entry in weights, unweighted stats
// scoring zero. "attack_power" is the one exception
// (genericAttackPowerWeightStats): its weight is the sum of the
// spec's attack_power AND ranged_attack_power weights, since a candidate's
// Stats map never carries both keys for the same generic AP bonus (see
// genericAttackPowerWeightStats' own comment) - a melee spec's
// ranged_attack_power weight and a hunter dps spec's attack_power weight
// are each the zero value in practice (data/curated/specs.json's
// weight_stats lists at most one of the two per spec), so this never
// double-counts a spec that actually cares about both.
func statWeight(stat string, weights map[string]float64) float64 {
	if stat != "attack_power" {
		return weights[stat]
	}
	total := 0.0
	for _, id := range genericAttackPowerWeightStats {
		total += weights[id]
	}
	return total
}

// score is the weighted sum of an item's resolved stats against a
// spec's stat weights, plus its weapon DPS (if any, converted through
// attackPowerPerDPS into the attack-power stat the slot it is
// equipped in cares about).
//
// weights is stat id -> normalised weight (sim/adapter.Weights'
// output: the reference stat weighs exactly 1). slot is the planner
// slot this score is being computed for - the same item can score
// differently in main_hand versus off_hand only in that its DPS
// converts through a different weaponAPStat entry, which for this
// engine build never actually differs (both are "attack_power"); the
// parameter exists so a future ranged-vs-melee split in the engine's
// own AttackPowerPerDPS does not require changing this function's
// signature.
//
// referenceDPSPerPoint is runWeights' own second return (simrun.go,
// bandReport.ReferenceDPSPerPoint's own source) - this lane's brief,
// item 1: "does a wand's/weapon's damage count for casters the way
// weaponScore counts it for melee?" It did not. weaponAPStat's
// conversion above only ever lands on "attack_power" or
// "ranged_attack_power" - the two weight_stats entries every hunter/
// melee/hybrid spec's own data/curated/specs.json row carries, and
// NO caster spec (mage/priest/warlock/shaman-elemental/druid-balance)
// carries EITHER (confirmed against specs.json directly: their own
// weight_stats lists spell_power/intellect/crit/hit/spell_haste/
// spell_penetration/<school>_power only), so weights[apStat] is
// always the Go zero value for them and this term was always exactly
// 0 - not a data gap on any one item, a structural one in score()
// itself: a caster's weapon slot had NO path for its damage to count
// at all, confirmed by wowsims-forever's own sim/core/wand.go/
// sim/mage/shoot.go doc comments ("wand damage in Classic never
// scales off AP... a mage with a wand equipped dealt zero wand damage
// until [Shoot was modeled]") - wand/melee weapon damage for a caster
// is flat, unscaled output, never an attack-power-derived quantity,
// so no weight_stats entry could ever give it one even if added.
//
// The fix, narrowed by the fourth wow-player sweep's own caster-sweep
// finding (bis-ranker-integrity-4, 2026-09-30): a caster never swings
// their main_hand/off_hand weapon, so its flat DPS must never convert
// into score at all there, regardless of referenceDPSPerPoint - the
// broader "any zero-AP-weight weapon slot" fallback this lane's own
// prior round shipped let a purely physical stick (Manual Crowd
// Pummeler: Str/Agi/a melee-haste proc, zero caster stats) outscore a
// real caster weapon (Staff of Jordan: spell power) at band 30/40 for
// shaman-elemental and druid-balance, because BOTH specs have zero
// attack_power weight and the old code could not tell "a caster's
// unswung melee weapon" from "a caster's wand, actually fired by the
// rotation" apart. The one real case a caster's weapon DPS is live
// damage is its ranged slot's wand, autoshot through the Shoot spell
// (id 5019) - and only for a spec whose OWN rotation casts it
// (castsShoot, weapon_requirements.go's aplRotationCastsShoot; mage/
// priest-shadow/warlock all do, shaman-elemental/druid-balance do
// not, matching every caster APL this build has curated). Dividing by
// referenceDPSPerPoint converts that wand DPS into this band's own
// score unit (reference-stat points) directly, the same unit
// conversion buildAlternatives (report.go) already applies in the
// other direction (score * referenceDPSPerPoint -> DPS). A
// main_hand/off_hand weapon with no AP weight at all (every caster,
// and any melee/hybrid item whose AP weight the run measured as
// exactly zero or insignificant - effectiveWeights already zeroed a
// non-positive one before this function ever sees it) instead
// contributes nothing from DPS, falling through to whatever its own
// Stats total, and pick.go's promoteLowValueWeapon takes over from
// there when that total is exactly 0 (its own doc: "the best by item
// level among sourced candidates carrying any of the spec's weight
// stats"). referenceDPSPerPoint <= 0 or castsShoot == false (every
// test that does not pass both) leaves the ranged branch a no-op too,
// matching the pre-fallback behaviour exactly.
// deadStatCount counts how many of c's own positive-amount stats this
// spec's weights do not value at all (statWeight <= 0) -
// bis-ranker-integrity-12 lane, item 5: an exact score() tie between
// two candidates is not evidence they are equally good gear for this
// spec - one may carry only stats the spec's own weights actually use
// while the other pads the identical total with a stat the spec never
// weighs (a rogue's Spell Power, a caster's Strength before this build
// weighed either) - so candidatesBySlot's own tie-break reads this
// count before falling back to item level.
func deadStatCount(c candidate, weights map[string]float64) int {
	count := 0
	for stat, amount := range c.Stats {
		if amount > 0 && statWeight(stat, weights) <= 0 {
			count++
		}
	}
	return count
}

func score(c candidate, slot string, weights map[string]float64, referenceDPSPerPoint float64, castsShoot bool) float64 {
	total := 0.0
	for stat, amount := range c.Stats {
		total += amount * statWeight(stat, weights)
	}
	if c.DPS > 0 {
		if apStat, ok := weaponAPStat[slot]; ok {
			if apWeight := weights[apStat]; apWeight > 0 {
				total += c.DPS * attackPowerPerDPS * apWeight
			} else if slot == "ranged" && castsShoot && referenceDPSPerPoint > 0 {
				total += c.DPS / referenceDPSPerPoint
			}
		}
	}
	return total
}
