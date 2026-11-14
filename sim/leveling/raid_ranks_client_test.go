package leveling

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
)

// Round 8 (auras lane): the raid buffs and debuffs the engine held at one
// value for every level now follow rank tables (fork sim/core/buff_ranks.go
// and raid_debuff_ranks.go). Every row is pinned here against the 1.60.1.70707
// spellconst, the client's SpellEffect and SpellLevels with the hotfix push.

// SpellEffect aura codes the round-8 rows use.
const (
	auraModArmorPercentOrFlat  = 674 // misc 1: armor, negative on a debuff
	auraModDamagePercentTaken  = 87
	auraModRangedAPAttacker    = 127
	auraModDamageTakenFlat     = 14
	miscStamina                = 2
	miscSpirit                 = 4
	miscHoly                   = 2
	miscAllMagic               = 126
	windfuryTotemFirstCastRank = 1
)

func TestPowerWordFortitudeAndDivineSpiritMatchTheClient(t *testing.T) {
	client := loadClassFrom(t, trueshotAuraSpellconstDir, "priest")
	assertRanks(t, "Power Word: Fortitude", client, core.PowerWordFortitudeRanks, auraModStat, miscStamina, false)
	assertRanks(t, "Divine Spirit", client, core.DivineSpiritRanks, auraModStat, miscSpirit, false)
}

func TestBloodPactMatchesTheClient(t *testing.T) {
	assertRanks(t, "Blood Pact", loadClassFrom(t, trueshotAuraSpellconstDir, "warlock"), core.BloodPactRanks, auraModStat, miscStamina, true)
}

// windfuryStrikeSpellconstDir is the first build whose spellconst carries
// push 113887's Windfury Totem: the rank's spell is a weapon strike.
const windfuryStrikeSpellconstDir = "../../data/builds/1.60.1.70938/spellconst/"

// effectWeaponDamage is SpellEffect.Effect 58, a weapon damage strike.
const effectWeaponDamage = 58

// Windfury Totem: the cast spells 8512, 10613, 10614 are learned at the
// engine's levels, and since push 113887 the rank's spell (8516, 10608,
// 10610) is the main-hand weapon strike the engine casts on a proc, no
// longer an attack power buff. The strike's attack power is the totem aura's
// dummy, which the class files do not carry; the fork's client test reads it.
func TestWindfuryTotemRanksMatchTheClient(t *testing.T) {
	client := loadClassFrom(t, windfuryStrikeSpellconstDir, "shaman")
	casts := []int32{8512, 10613, 10614}
	for i, id := range casts {
		rank := i + windfuryTotemFirstCastRank
		cast, ok := client.ByID(id)
		if !ok || cast.SpellLevel != core.WindfuryTotemLevel[rank] {
			t.Errorf("Windfury Totem %d: client level %d (found %v), engine %d", id, cast.SpellLevel, ok, core.WindfuryTotemLevel[rank])
		}
		strike, ok := client.ByID(core.WindfuryStrikeSpellId[rank])
		if !ok {
			t.Fatalf("Windfury Totem strike %d is not in the client table", core.WindfuryStrikeSpellId[rank])
		}
		if len(strike.Effects) != 1 || strike.Effects[0].Effect != effectWeaponDamage {
			t.Errorf("Windfury Totem strike %d: client effects %+v, want one weapon damage effect (58)", strike.ID, strike.Effects)
		}
	}
}

// assertReductionRanks is assertRanks for a debuff whose client amount is
// the negative of the engine's (armor or resistance removed).
func assertReductionRanks(t *testing.T, label, class string, ranks core.BuffRanks, aura, misc int32) {
	t.Helper()
	client := loadClassFrom(t, trueshotAuraSpellconstDir, class)
	for _, rank := range ranks {
		spell, ok := client.ByID(rank.SpellID)
		if !ok {
			t.Errorf("%s: spell %d is not in the client table", label, rank.SpellID)
			continue
		}
		if spell.SpellLevel != rank.Level {
			t.Errorf("%s %d: engine learns at %d, client at %d", label, rank.SpellID, rank.Level, spell.SpellLevel)
		}
		if amount := -findEffect(spell, aura, misc).Amount; amount != rank.Amount {
			t.Errorf("%s %d: engine removes %v, client %v", label, rank.SpellID, rank.Amount, amount)
		}
	}
}

func TestRaidDebuffRanksMatchTheClient(t *testing.T) {
	assertReductionRanks(t, "Sunder Armor", "warrior", core.SunderArmorRanks, auraModArmorPercentOrFlat, miscArmor)
	assertReductionRanks(t, "Faerie Fire", "druid", core.FaerieFireRanks, auraModArmorPercentOrFlat, miscArmor)
	assertReductionRanks(t, "Curse of Recklessness", "warlock", core.CurseOfRecklessnessRanks, auraModArmorPercentOrFlat, miscArmor)
	assertReductionRanks(t, "Curse of the Elements resistance", "warlock", core.CurseOfTheElementsResistanceRanks, auraModResistance, miscAllMagic)

	client := loadClassFrom(t, trueshotAuraSpellconstDir, "warlock")
	assertRanks(t, "Curse of the Elements damage", client, core.CurseOfTheElementsDamageRanks, auraModDamagePercentTaken, miscAllMagic, false)
	assertRanks(t, "Hunter's Mark", loadClassFrom(t, trueshotAuraSpellconstDir, "hunter"), core.HuntersMarkRanks, auraModRangedAPAttacker, 0, false)
	assertRanks(t, "Judgement of the Crusader", loadClassFrom(t, trueshotAuraSpellconstDir, "paladin"), core.JudgementOfTheCrusaderRanks, auraModDamageTakenFlat, miscHoly, false)
}

// Forever's Curse of Recklessness gives the target no attack power: its
// ranks' only effects are the armor reduction and aura 92 (no fleeing).
func TestCurseOfRecklessnessHasNoAttackPowerInTheClient(t *testing.T) {
	client := loadClassFrom(t, trueshotAuraSpellconstDir, "warlock")
	for _, rank := range core.CurseOfRecklessnessRanks {
		spell, _ := client.ByID(rank.SpellID)
		if ap := findEffect(spell, auraModAttackPower, 0); ap.Amount != 0 {
			t.Errorf("Curse of Recklessness %d states %v attack power", rank.SpellID, ap.Amount)
		}
	}
}
