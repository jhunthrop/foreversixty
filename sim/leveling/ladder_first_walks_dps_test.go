package leveling

import "testing"

// Sim-accuracy round 29 (walks-dps): Berserker Rage joins the rotation only at 36, so a leveling Fury warrior takes
// Flurry at 35 to 37 and Improved Whirlwind at 38 and 39 before Improved Berserker Rage's second
// point, which waits for 41. In Arms, the second point of Improved Charge (no Charge in a fight)
// waits behind Anger Management and Deep Wounds as the toll Impale needs.
func TestFuryTakesFlurryBeforeImprovedBerserkerRage(t *testing.T) {
	checkLadder(t, "warrior", "fury", 1, map[int]string{
		35: "00000000000000000-050530025121100010-000000000000000000",
		38: "00000000000000000-050530025121100130-000000000000000000",
		40: "00000000000000000-050530025121100231-000000000000000000",
		50: "30312000000000000-050530025122100231-000000000000000000",
		55: "30313013000000000-050530025122100231-000000000000000000",
	})
}

// Sim-accuracy 2026-11-12 (Protection warrior lane): Shield Specialization at 10 to 14, Improved Revenge at 15 to 17
// and Improved Bloodrage at 18 and 19 (threat +3 to +15% at 11 to 18 once Bloodrage costs health and rage from damage
// taken ignores armor). Round 29 (walks-dps): a leveling Protection warrior takes Improved Revenge before 20, Defiance at
// 20 to 22 and Master of Defense at 23 and 24 before Anticipation. Round 30: on the dungeon-table
// boss a bare tank's chance of death is 0 at every level with either order, so under the tank
// rule Improved Sunder Armor (25 to 27), Improved Thunder Clap (28 to 30) and Focused Rage (31,
// 32) come before Anticipation (33 to 35, 37, 38) and Last Stand (36), threat +1 to +11% at 25 to
// 37. Concussion Blow 39, Shield Slam 40 and Bastion are unchanged.
func TestProtectionWarriorTakesDefianceBeforeAnticipation(t *testing.T) {
	checkLadder(t, "warrior", "protection", 2, map[int]string{
		14: "00000000000000000-000000000000000000-050000000000000000",
		17: "00000000000000000-000000000000000000-050030000000000000",
		19: "00000000000000000-000000000000000000-250030000000000000",
		20: "00000000000000000-000000000000000000-250030000100000000",
		25: "00000000000000000-000000000000000000-250030020310000000",
		28: "00000000000000000-000000000000000000-250031020330000000",
		30: "00000000000000000-000000000000000000-250033020330000000",
		32: "00000000000000000-000000000000000000-250033020330000200",
		33: "00000000000000000-000000000000000000-250133020330000200",
		36: "00000000000000000-000000000000000000-250333120330000200",
		38: "00000000000000000-000000000000000000-250533120330000200",
		40: "00000000000000000-000000000000000000-250533120330001201",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling Arcane mage takes Missile Barrage at 21 and 22, Arcane Impact at 23 to 25
// and Arcane Meditation at 26 to 28 before Arcane Resilience (armor) and Arcane Geometry
// (range), which pay the rows at 29 and 39; Presence of Mind comes at 30. From 41 Precision
// comes first, ahead of Incineration and Improved Frostbolt. Sim-accuracy 2026-11-12 (arcane):
// Arcane Wrath at 44, right after Precision, for the fourth rank of Incineration (+6.8% at 45,
// +2.7 to +2.9% at 50, +0.4 to +1.0% at 60; design/reviews/2026-11-12-sim-accuracy/arcane.md).
func TestArcaneTakesItsDamageTalentsBeforeArmor(t *testing.T) {
	checkLadder(t, "mage", "arcane", 0, map[int]string{
		21: "140005100100000000-00000000000000000-000000000000000000",
		25: "140005203100000000-00000000000000000-000000000000000000",
		30: "141005203100301000-00000000000000000-000000000000000000",
		35: "141005203100301500-00000000000000000-000000000000000000",
		38: "141005203100301530-00000000000000000-000000000000000000",
		45: "141005213100311531-10000000000000000-003000000000000000",
		50: "141005213100311531-30000000000000000-033000000000000000",
		51: "141005213100311531-30000000000000000-043000000000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): after Combustion a leveling Fire mage takes Precision at 41 to 43, Wand
// Specialization at 44 and 45, Improved Channeling (the Arcane row toll), Arcane Concentration
// and Arcane Blast at 54, then Missile Barrage; Accelerant 5 and Burning Soul, which do nothing
// for the Scorch filler or a target that never hits you, come last.
func TestFireTakesPrecisionAfterCombustion(t *testing.T) {
	checkLadder(t, "mage", "fire", 1, map[int]string{
		41: "000000000000000000-54500001230133031-001000000000000000",
		45: "200000000000000000-54500001230133031-003000000000000000",
		50: "230002000000000000-54500001230133031-003000000000000000",
		55: "230005100100000000-54500001230133031-003000000000000000",
		59: "230005200100000000-55500201230133031-003000000000000000",
	})
}

// Sim-accuracy 2026-11-12 (casters): after Ice Barrier at 40 a leveling Frost mage spends its Arcane
// points in the round 29 order (Wand Specialization, Improved Channeling 3, Arcane Concentration,
// Missile Barrage, Arcane Impact to 55), then Arcane Meditation at 56 to 58 and the last two
// Improved Channeling points at 59 and 60, where the old build took Incineration; the ladder's own
// rule would have spent the two toll points at 46 and 47 ahead of Arcane Concentration (-1.9 to
// -2.4% at 48 to 53). Unchanged to 55, +0.5% at 58, +2.2 to +2.3% at 60
// (design/reviews/2026-11-12-sim-accuracy/casters.md).
func TestFrostTakesArcaneMeditationAfterArcaneImpact(t *testing.T) {
	checkLadder(t, "mage", "frost", 2, map[int]string{
		47: "230002000000000000-00000000000000000-253510130101301221",
		55: "230005203000000000-00000000000000000-253510130101301221",
		58: "230005203000300000-00000000000000000-253510130101301221",
		60: "250005203000300000-00000000000000000-253510130101301221",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling Affliction warlock takes Amplify Curse at 22, Pandemic at 23 to 25 and
// Malevolence at 26 and 27 before Grim Reach (range) and Soul Harvest (a Drain Soul kill); Soul Siphon pays
// the row at 31 to 33 so Shadow Mastery is full at 39 and Soul Siphon is ready for Wrack at 40.
// Sim-accuracy 2026-11-12 (warlocks pass 3): from 42 the order takes Improved Bane of Agony, the last three
// Malevolence, Bane, Nightfall, Aftermath and two Unholy Power, and no longer Improved Drains (it stopped
// reaching Wrack with the 70707 text and the rotation channels no drain) or the Demonology toll points.
func TestAfflictionTakesPandemicBeforeTheDrainTalents(t *testing.T) {
	checkLadder(t, "warlock", "affliction", 0, map[int]string{
		22: "25500000100000000-0000000000000000000-0000000000000000",
		25: "25500000130000000-0000000000000000000-0000000000000000",
		30: "25520000132001000-0000000000000000000-0000000000000000",
		35: "25521000132001310-0000000000000000000-0000000000000000",
		38: "25521000132001340-0000000000000000000-0000000000000000",
		40: "25521000132001351-0000000000000000000-0000000000000000",
		45: "25522020134001351-0000000000000000000-0000000000000000",
		50: "25522020135001351-0000000000000000000-0040000000000000",
		55: "25522020135201351-0000000000000000000-0050020000000000",
		59: "25522020135201351-0001000000000000000-0050050000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): Beast Mastery takes Unleashed Fury at 20 to 24 and Ferocity at 25 to 29
// before Endurance Training's fourth point and Bestial Swiftness, which add pet health and run
// speed; Bestial Discipline comes at 30 and 31, and Bestial Swiftness and Intimidation, the toll
// for Bestial Wrath, at 32 and 33.
func TestBeastMasteryTakesUnleashedFuryBeforeThePetHealthPoints(t *testing.T) {
	checkLadder(t, "hunter", "beast-mastery", 0, map[int]string{
		20: "5320000100000000-0000000000000000-000000000000000000",
		25: "5320000501000000-0000000000000000-000000000000000000",
		30: "5320000505000100-0000000000000000-000000000000000000",
		33: "5320001505001200-0000000000000000-000000000000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): Marksmanship takes Careful Aim at 15 to 19 before Efficiency, and from 41
// Focused Fire at 46 and 47 before the Endurance Training points that only pay its row.
func TestMarksmanshipTakesCarefulAimBeforeEfficiency(t *testing.T) {
	checkLadder(t, "hunter", "marksmanship", 1, map[int]string{
		15: "0000000000000000-0050010000000000-000000000000000000",
		20: "0000000000000000-0050150000000000-000000000000000000",
		49: "5220000000000000-0050552001521050-000000000000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): after Moonkin Form a leveling Balance druid takes Heart of the Wild at 42
// to 46, Naturalist (2% damage a point) at 47 to 51 and Reflection at 52 to 54, before Nature's
// Focus, which only saves casts from pushback on a target that never hits you.
func TestBalanceTakesNaturalistBeforeNaturesFocus(t *testing.T) {
	checkLadder(t, "druid", "balance", 0, map[int]string{
		42: "5232220115101151-00000000000000000000-0100000000000000",
		45: "5232220115101151-00000000000000000000-0400000000000000",
		50: "5232220115101151-00000000000000000000-0540000000000000",
		55: "5232220115101151-00000000000000000000-1553000000000000",
		58: "5232220115101151-00000000000000000000-4553000000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling bear pays the 10-point row with Thick Hide and one point of
// Feral Swiftness, so Feral Instinct (Swipe, which a single target never sees) waits for 58 to
// 60; Sharpened Claws comes at 22 and 23 and Rend and Tear from 36. Feral Swiftness's second
// point at 39 and Natural Reaction before Heart of the Wild and Naturalist hold the tank fight's
// chance of death within 5 points of the old order at every level (the round 29 tank rule).
func TestBearPaysItsRowWithThickHide(t *testing.T) {
	checkLadder(t, "druid", "feral-bear", 1, map[int]string{
		16: "0000000000000000-51000100000000000000-0000000000000000",
		20: "0000000000000000-51100301000000000000-0000000000000000",
		25: "0000000000000000-51100312020100000000-0000000000000000",
		28: "0000000000000000-51100332020101000000-0000000000000000",
		29: "0000000000000000-51100332020102000000-0000000000000000",
		30: "0000000000000000-51100332020102100000-0000000000000000",
		35: "0000000000000000-51100332023102120000-0000000000000000",
		38: "0000000000000000-51100332023102123000-0000000000000000",
		39: "0000000000000000-51200332023102123000-0000000000000000",
		40: "0000000000000000-51200332023102123010-0000000000000000",
		45: "0000000000000000-51200332023102125310-0000000000000000",
		50: "0000000000000000-51200332023102125510-0300000000000000",
		55: "0000000000000000-51200332023102125510-0530000000000000",
		59: "0000000000000000-51220332023102125510-0550000000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling Retribution paladin takes one point of Vindication (it only lowers
// the target's attack power), then Seal of Command, Sacred Arbiter and Two-Handed Weapon
// Specialization at 24 to 28 and Vengeance at 30 to 32; Vindication's last two points and
// Repentance pay the rows at 29, 33 and 34.
func TestRetributionTakesSealOfCommandBeforeVindication(t *testing.T) {
	checkLadder(t, "paladin", "retribution", 2, map[int]string{
		24: "00000000000000000-00000000000000000-05233101000000000",
		25: "00000000000000000-00000000000000000-05233101001000000",
		30: "00000000000000000-00000000000000000-05233201001300100",
		33: "00000000000000000-00000000000000000-05233301001300300",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling Shadow priest takes Shadow Weaving at 31 to 33, before Blackout (a
// stun proc) and Vampiric Embrace (a global the rotation spends on healing), which pay the row
// for Darkness at 34 to 36. In Discipline, Mental Strength comes before Silent Resolve's threat
// point, which waits for 60.
func TestShadowTakesShadowWeavingBeforeBlackout(t *testing.T) {
	checkLadder(t, "priest", "shadow", 2, map[int]string{
		31: "020000000000000000-00000000000000000-50542001200100000",
		35: "020000000000000000-00000000000000000-52542001200300000",
		50: "025003000000000000-00000000000000000-52542001201300051",
		55: "025005030000000000-00000000000000000-52542001201300051",
		59: "025005031030000000-00000000000000000-52542001201300051",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling Elemental shaman takes Elemental Alacrity at 15 to 17 before
// Concussion, Call of Thunder at 30 and Lightning Overload at 31 to 33 before Trial by Fire's
// second point (Fire Nova and Magma Totem are pack spells), and Tidal Focus's spell hit at 41 to
// 45 before Thundering Strikes. Sim-accuracy 2026-11-12 (shamans at 60): Mindfulness 2 at 46 and 47 ahead of
// Thundering Strikes (+10% at 47), then Thundering Strikes, Ancestral Knowledge, Shamanistic Focus.
func TestElementalTakesAlacrityAndCallOfThunderEarly(t *testing.T) {
	checkLadder(t, "shaman", "elemental", 0, map[int]string{
		15: "5000100000000000-000000000000000000-0000000000000000",
		20: "5300300000000000-000000000000000000-0000000000000000",
		30: "5500331210001000-000000000000000000-0000000000000000",
		45: "5500331220001351-000000000000000000-0500000000000000",
		47: "5500331220001351-000000000000000000-0520000000000000",
		49: "5500331220001351-020000000000000000-0520000000000000",
		58: "5500331220001351-055000010000000000-0520000000000000",
	})
}

// Sim-accuracy round 29 (walks-dps): a leveling Enhancement shaman takes Mental Dexterity at 15 to 17 before
// Ancestral Knowledge, Mental Quickness at 33 and 34 and Maelstrom Weapon at 35 to 39 before
// Improved Ghost Wolf (a travel talent), so the instant Lightning Bolt arrives at 39, and Tidal
// Focus's hit at 44 to 48 after one point of Convection. Sim-accuracy 2026-11-12 (shamans at 60): Improved
// Stormstrike at 49 and 50 (+5% at 50), Mindfulness 51 to 53, Convection 54 to 57, Elemental Devastation last.
func TestEnhancementTakesMentalDexterityAndMaelstromEarly(t *testing.T) {
	checkLadder(t, "shaman", "enhancement", 1, map[int]string{
		15: "0000000000000000-050010000000000000-0000000000000000",
		20: "0000000000000000-053030000000000000-0000000000000000",
		35: "0000000000000000-055030310051002010-0000000000000000",
		38: "0000000000000000-055030310051002040-0000000000000000",
		40: "0000000000000000-055030310051002051-0000000000000000",
		45: "1000000000000000-055030310051202051-0200000000000000",
		50: "1000000000000000-055030310051202251-0500000000000000",
		55: "3000000000000000-055030310051202251-0530000000000000",
		57: "5000000000000000-055030310051202251-0530000000000000",
		60: "5030000000000000-055030310051202251-0530000000000000",
	})
}
