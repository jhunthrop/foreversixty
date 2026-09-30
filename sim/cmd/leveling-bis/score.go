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
// The fix: when the slot's own AP-based weight is zero (every caster,
// and any melee/hybrid item whose AP weight the run measured as
// exactly zero or insignificant - effectiveWeights already zeroed a
// non-positive one before this function ever sees it), fall back to
// treating the weapon's flat DPS as flat DPS: dividing it by
// referenceDPSPerPoint converts it into this band's own score unit
// (reference-stat points) directly, the same unit conversion
// buildAlternatives (report.go) already applies in the other
// direction (score * referenceDPSPerPoint -> DPS) - a caster's plain
// damage output is exactly as real a DPS contribution as a melee
// spec's attack power, it simply never had a weight_stats entry to
// convert through. referenceDPSPerPoint <= 0 (every test that does
// not pass a real one) leaves this a no-op, matching the old
// behaviour exactly.
func score(c candidate, slot string, weights map[string]float64, referenceDPSPerPoint float64) float64 {
	total := 0.0
	for stat, amount := range c.Stats {
		total += amount * statWeight(stat, weights)
	}
	if c.DPS > 0 {
		if apStat, ok := weaponAPStat[slot]; ok {
			if apWeight := weights[apStat]; apWeight > 0 {
				total += c.DPS * attackPowerPerDPS * apWeight
			} else if referenceDPSPerPoint > 0 {
				total += c.DPS / referenceDPSPerPoint
			}
		}
	}
	return total
}
