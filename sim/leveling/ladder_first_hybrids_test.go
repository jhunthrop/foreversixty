package leveling

import "testing"

// Sim-accuracy round 28 (hybrids): a leveling cat takes Sharpened Claws at 23 and 24,
// ahead of Savage Fury, and Leader of the Pack, Blood Frenzy, Natural Instinct and Rend and
// Tear at 31 to 38, ahead of Predatory Discipline, which only changes Ferocious Bite (the
// cat Bites little before raid gear) and comes at 39 to open Shifting Power. Shredding
// Attacks keeps 20 to 22: its first point does nothing alone, but it is the price of a
// cheap Shred at 21. Feral Charge keeps 25 for its Cat Form leap behind the target. Levels
// 20, 25 to 30 (except 26) and 40 do not change.
func TestFeralTakesSharpenedClawsAndTheCritTalentsEarlier(t *testing.T) {
	const feralTree = 1
	const empty = "0000000000000000"
	want := map[int]string{
		19: empty + "-50212000000000000000-" + empty,
		20: empty + "-50212010000000000000-" + empty,
		22: empty + "-50212030000000000000-" + empty,
		23: empty + "-50212030010000000000-" + empty,
		24: empty + "-50212030020000000000-" + empty,
		25: empty + "-50212030120000000000-" + empty,
		26: empty + "-50212031120000000000-" + empty,
		27: empty + "-50212032120000000000-" + empty,
		30: empty + "-50212032123000000000-" + empty,
		31: empty + "-50212032123000100000-" + empty,
		32: empty + "-50212032123001100000-" + empty,
		34: empty + "-50212032123001120000-" + empty,
		35: empty + "-50212032123002120000-" + empty,
		38: empty + "-50212032123002123000-" + empty,
		39: empty + "-50212032123012123000-" + empty,
		40: empty + "-50212032123012123001-" + empty,
	}
	for level, w := range want {
		if got := ladderAt(t, "druid", "feral", feralTree, level); got != w {
			t.Errorf("level %d Feral build %q, want %q", level, got, w)
		}
	}
}

// Sim-accuracy round 28 (hybrids): a leveling Survival hunter takes Predator's Edge at 25
// to 29 and Resourcefulness, Expose Prey, Survivalist's Discipline and Lightning Reflexes
// after Strider Kick, before Deterrence and Counterattack. Deterrence is never pressed on a
// target that hits someone else, and Counterattack needs a parry, so neither does anything
// in the fight; they come at 38 and 39, where they pay the toll for Lacerating Strikes.
func TestSurvivalTakesPredatorsEdgeBeforeDeterrence(t *testing.T) {
	const survivalTree = 2
	const empty = "0000000000000000-0000000000000000-"
	want := map[int]string{
		24: empty + "500230230000000000",
		25: empty + "500230230010000000",
		29: empty + "500230230050000000",
		30: empty + "500230230050000100",
		32: empty + "500230230050200100",
		34: empty + "500230230050220100",
		36: empty + "500230230050222100",
		37: empty + "500230230050222110",
		38: empty + "500230231050222110",
		39: empty + "500230231051222110",
		40: empty + "500230231051222111",
	}
	for level, w := range want {
		if got := ladderAt(t, "hunter", "survival", survivalTree, level); got != w {
			t.Errorf("level %d Survival build %q, want %q", level, got, w)
		}
	}
}

// Sim-accuracy round 28 (hybrids): a leveling Protection paladin takes Shield
// Specialization at 21 and 22, before Sacred Duty's second point, and Swift Judgement at
// 25, the first level the 15-point row opens, before Toughness's last four points. Sacred
// Duty only shortens cooldowns the tank seldom needs at that level and Toughness only adds
// armor, while Shield Specialization's blocks return mana and Swift Judgement is a second
// Judgement. Level 20 keeps Templar's Bulwark and level 30 does not change.
func TestProtectionTakesSwiftJudgementBeforeToughness(t *testing.T) {
	const protectionTree = 1
	const empty = "00000000000000000"
	want := map[int]string{
		20: empty + "-15301100000000000-" + empty,
		21: empty + "-15301101000000000-" + empty,
		22: empty + "-15301102000000000-" + empty,
		23: empty + "-15302102000000000-" + empty,
		24: empty + "-15302102100000000-" + empty,
		25: empty + "-15302102110000000-" + empty,
		29: empty + "-15302102510000000-" + empty,
		30: empty + "-15302102510001000-" + empty,
	}
	for level, w := range want {
		if got := ladderAt(t, "paladin", "protection", protectionTree, level); got != w {
			t.Errorf("level %d Protection build %q, want %q", level, got, w)
		}
	}
}
