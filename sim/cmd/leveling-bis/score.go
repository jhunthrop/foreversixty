package main

import (
	"math"
	"sort"
)

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
	if stat == itemHealingStat {
		return weights[healingWeightStat]
	}
	if stat != "attack_power" {
		return weights[stat]
	}
	total := 0.0
	for _, id := range genericAttackPowerWeightStats {
		total += weights[id]
	}
	return total
}

// wandScorePerDPS is the score a caster's ranged slot earns per point of a
// wand's tooltip DPS. A wand is only worth the damage the rotation actually
// shoots with it, so the credit is measured, not assumed: shootDPS is the
// Shoot damage per second a sim of the band's own set dealt (shootDPS,
// wand_credit.go), wandDPS the tooltip DPS of the wand that set wore. Their
// ratio is the band's Shoot time share times everything that scales a shot
// (hit, crit, partial resists, Wand Specialization), so another wand's DPS
// times it is the damage that wand would add. Dividing by
// referenceDPSPerPoint turns that DPS into reference-stat points, the unit
// every other stat scores in. A warlock raid set that never shoots measures
// 0 and its wand scores on its stats alone (round 6, Torch of Light).
//
// A caster never swings a main or off hand weapon, so score() applies this
// to the ranged slot only, and only for a spec whose rotation casts Shoot
// (main.go passes 0 otherwise).
func wandScorePerDPS(shootDPS, wandDPS, referenceDPSPerPoint float64) float64 {
	if shootDPS <= 0 || wandDPS <= 0 || referenceDPSPerPoint <= 0 {
		return 0
	}
	return shootDPS / wandDPS / referenceDPSPerPoint
}

// fullWandScorePerDPS is the credit of a wand fired the whole fight at its
// tooltip DPS. It is the first pass only: it makes the band's first set wear
// the hardest-hitting wand, so the sim that measures wandScorePerDPS has a
// wand to shoot.
func fullWandScorePerDPS(referenceDPSPerPoint float64) float64 {
	return wandScorePerDPS(1, 1, referenceDPSPerPoint)
}

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

// score is the weighted sum of an item's resolved stats against a spec's
// stat weights, plus its weapon DPS: through attackPowerPerDPS for a slot
// whose attack-power weight is positive, or, for a caster's ranged wand,
// through wandScorePerDPS (wandCredit; 0 for every spec that never shoots).
func score(c candidate, slot string, weights map[string]float64, wandCredit float64) float64 {
	// Summed in sorted stat order (task L1.7 fix round 3): a Go map ranges
	// in a random order, and a float sum taken in another order can differ
	// by an ulp, so two items with the same stats did not always tie.
	stats := make([]string, 0, len(c.Stats))
	for stat := range c.Stats {
		stats = append(stats, stat)
	}
	sort.Strings(stats)
	total := 0.0
	for _, stat := range stats {
		total += c.Stats[stat] * statWeight(stat, weights)
	}
	// Armor is an item property, not a stat line; only a tank's weights
	// carry an "armor" entry, so this adds nothing for any other spec.
	total += float64(c.Armor) * weights[armorWeightStat]
	if c.DPS > 0 {
		if apStat, ok := weaponAPStat[slot]; ok {
			if apWeight := weights[apStat]; apWeight > 0 {
				total += c.DPS * attackPowerPerDPS * apWeight
			} else if slot == "ranged" && wandCredit > 0 {
				total += c.DPS * wandCredit
			}
		}
	}
	return total
}

// scorePrecision is the grid a stored score and a published dps delta are
// rounded to (task L1.7 fix round 3): a billionth of a point, far below any
// difference a stat point makes and far above float summation noise.
const scorePrecision = 1e9

// roundScore rounds x to scorePrecision, so sums that are equal in exact
// arithmetic but built from different stat sets (0.1+0.2 against 0.3) are
// one value, and the exact comparisons in sortCandidates, pick's Ties loop,
// mergeByScore and the alternatives' sort fall to their tie-breaks. A
// negative zero comes back as 0, so the JSON never prints "-0".
func roundScore(x float64) float64 {
	r := math.Round(x*scorePrecision) / scorePrecision
	if r == 0 {
		return 0
	}
	return r
}
