package leveling

import "testing"

// checkLadder holds a spec's ladder build to want at each listed level.
func checkLadder(t *testing.T, class, spec string, ownTree int, want map[int]string) {
	t.Helper()
	for level, w := range want {
		if got := ladderAt(t, class, spec, ownTree, level); got != w {
			t.Errorf("level %d %s-%s build %q, want %q", level, class, spec, got, w)
		}
	}
}

// Sim-accuracy round 28 (casters): a leveling Destruction warlock spends no point on
// Destructive Reach (range) or Molten Skin (damage taken) while a damage talent can take
// it. Bane 10 to 14, Aftermath 15 to 17, Cataclysm 18 and 19, Shadowburn 20, Ruin 21 to
// 24, Agonizing Flames 25 to 27 and Ruin 28. Conflagrate eats Immolate and loses damage
// until Fire and Brimstone backs it, so Bane of Havoc (the Incinerate toll) comes at 30,
// Conflagrate at 31 with Fire and Brimstone 32 to 34, and Shadow and Flame 35 to 38, five
// levels sooner; the second Destructive Reach point and Molten Skin wait for 39 and 41 to
// 43. Level 44 is the guide build's.
func TestDestructionTakesBaneBeforeDestructiveReach(t *testing.T) {
	const destructionTree = 2
	const empty = "00000000000000000-0000000000000000000-"
	checkLadder(t, "warlock", "destruction", destructionTree, map[int]string{
		10: empty + "0010000000000000",
		14: empty + "0050000000000000",
		17: empty + "0050030000000000",
		19: empty + "0050230000000000",
		20: empty + "0050230100000000",
		24: empty + "0050234100000000",
		25: empty + "0050234101000000",
		28: empty + "0050235103000000",
		29: empty + "1050235103000000",
		30: empty + "1050235103001000",
		31: empty + "1050235103101000",
		34: empty + "1050235103101300",
		35: empty + "1050235103101310",
		38: empty + "1050235103101340",
		40: empty + "2050235103101341",
		43: empty + "2053235103101341",
		44: empty + "2053235103101351",
	})
}

// Sim-accuracy round 28 (casters): a leveling Demonology warlock takes Unholy Power (pet
// damage) at 10 to 14 and Fel Vitality (mana) at 15 to 17, before Improved Health Funnel
// and Demonic Embrace, which do nothing in a fight the demon is not dying in; the walk
// rejoins the old one at 25. Demonic Knowledge comes at 31 to 33, before Decimation's
// second point, which does nothing before Soul Fire at 48.
func TestDemonologyTakesUnholyPowerBeforeHealthFunnel(t *testing.T) {
	const demonologyTree = 1
	const pre, post = "00000000000000000-", "-0000000000000000"
	checkLadder(t, "warlock", "demonology", demonologyTree, map[int]string{
		10: pre + "0001000000000000000" + post,
		14: pre + "0005000000000000000" + post,
		17: pre + "0005003000000000000" + post,
		18: pre + "0005003100000000000" + post,
		20: pre + "2005003100000000000" + post,
		24: pre + "2035103100000000000" + post,
		25: pre + "2035113100000000000" + post,
		30: pre + "2035113101210001000" + post,
		31: pre + "2035113101210001100" + post,
		33: pre + "2035113101210001300" + post,
		34: pre + "2035113101220001300" + post,
	})
}

// Sim-accuracy round 28 (casters): a leveling Frost mage takes Precision at 15 to 17,
// before Frost Warding, which only adds armor. Ice Shards 18 and 19 open the row for
// Frost Channeling at 20 to 22 (the mana that keeps Frostbolt going); Ice Shards ends at
// 25. Frostbite 26 lets Shatter (27 and 28) work on a frozen mob, Ice Lance at 29 opens
// Fingers of Frost at 30 and 31, six levels sooner, and Shatter ends at 32. Frost Warding,
// Permafrost, Ice Block and Cold Snap fill 33 to 37, and Ice Barrier still comes at 40.
func TestFrostTakesPrecisionBeforeFrostWarding(t *testing.T) {
	const frostTree = 2
	const empty = "000000000000000000-00000000000000000-"
	checkLadder(t, "mage", "frost", frostTree, map[int]string{
		14: empty + "050000000000000000",
		15: empty + "051000000000000000",
		17: empty + "053000000000000000",
		19: empty + "053200000000000000",
		20: empty + "053200010000000000",
		22: empty + "053200030000000000",
		25: empty + "053500030000000000",
		26: empty + "053510030000000000",
		29: empty + "053510030100200000",
		30: empty + "053510030100200100",
		32: empty + "053510030100300200",
		37: empty + "253510130101301200",
		40: empty + "253510130101301221",
	})
}
