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
func score(c candidate, slot string, weights map[string]float64) float64 {
	total := 0.0
	for stat, amount := range c.Stats {
		total += amount * weights[stat]
	}
	if c.DPS > 0 {
		if apStat, ok := weaponAPStat[slot]; ok {
			total += c.DPS * attackPowerPerDPS * weights[apStat]
		}
	}
	return total
}
