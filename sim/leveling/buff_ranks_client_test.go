package leveling

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/spellconst"
)

// The engine's rank tables for Arcane Intellect, Blessing of Might, Mark
// of the Wild and Battle Shout (fork sim/core/buff_ranks.go) are the
// client's amounts, not the vanilla AQ-era values wowsims shipped. This
// test pins every row against the client spellconst the data pipeline
// emits for the build below (the committed copy of the SpellEffect and
// SpellLevels rows), so a new build moves this test instead of silently
// moving a number.
const clientSpellconstDir = "../../data/builds/1.60.1.70009/spellconst/"

// SpellEffect aura codes and misc values the buffs use.
const (
	auraModResistance  = 22 // misc 1 = armor, a school mask otherwise
	auraModStat        = 29 // misc -1 = all stats, 3 = Intellect
	auraModAttackPower = 99
	// Heart of the Lion's stat effect, on both of its spells.
	auraModTotalStatPercentage = 137
	// Blessing of Wisdom's single effect is a periodic energize (aura 24)
	// of mana every five seconds.
	auraPeriodicEnergize = 24
	// Trueshot Aura's single effect is ranged attack power (aura 124, misc
	// 0); the engine grants the same amount to melee as the client's area
	// aura does.
	auraModRangedAttackPower = 124

	miscArmor     = 1
	miscAllStats  = -1
	miscIntellect = 3
	// perLevelTolerance absorbs the float32 rounding of the client's
	// points-per-level column (0.6 is stored as 0.60000002384).
	perLevelTolerance = 0.001
)

// resistanceMasks are the engine's five resistance schools as the
// client's school masks: fire, nature, frost, shadow, arcane.
var resistanceMasks = []int32{4, 8, 16, 32, 64}

func loadClass(t *testing.T, class string) spellconst.Class {
	t.Helper()
	loaded, err := spellconst.Load(clientSpellconstDir + class + ".json")
	if err != nil {
		t.Fatalf("loading the client's %s spellconst: %v", class, err)
	}
	return loaded
}

// findEffect returns the spell's effect with the aura and misc value; the
// zero Effect (amount 0) means the rank states no such effect.
func findEffect(spell spellconst.Spell, aura, misc int32) spellconst.Effect {
	for _, effect := range spell.Effects {
		if effect.Aura == aura && effect.MiscValue == misc {
			return effect
		}
	}
	return spellconst.Effect{}
}

func assertRanks(t *testing.T, label string, client spellconst.Class, ranks core.BuffRanks, aura, misc int32, checkMaxLevel bool) {
	t.Helper()
	for _, rank := range ranks {
		spell, ok := client.ByID(rank.SpellID)
		if !ok {
			t.Errorf("%s: spell %d is not in the client table", label, rank.SpellID)
			continue
		}
		if spell.SpellLevel != rank.Level {
			t.Errorf("%s %d: engine learns at %d, client at %d", label, rank.SpellID, rank.Level, spell.SpellLevel)
		}
		effect := findEffect(spell, aura, misc)
		if effect.Amount != rank.Amount {
			t.Errorf("%s %d: engine amount %v, client %v", label, rank.SpellID, rank.Amount, effect.Amount)
		}
		if math.Abs(effect.PointsPerLevel-rank.PerLevel) > perLevelTolerance {
			t.Errorf("%s %d: engine points per level %v, client %v", label, rank.SpellID, rank.PerLevel, effect.PointsPerLevel)
		}
		if checkMaxLevel && spell.MaxLevel != rank.MaxLevel {
			t.Errorf("%s %d: engine max level %d, client %d", label, rank.SpellID, rank.MaxLevel, spell.MaxLevel)
		}
	}
}

func TestArcaneIntellectRanksMatchTheClient(t *testing.T) {
	assertRanks(t, "Arcane Intellect", loadClass(t, "mage"), core.ArcaneIntellectRanks, auraModStat, miscIntellect, false)
}

func TestBlessingOfMightRanksMatchTheClient(t *testing.T) {
	assertRanks(t, "Blessing of Might", loadClass(t, "paladin"), core.BlessingOfMightRanks, auraModAttackPower, 0, false)
}

func TestBlessingOfWisdomRanksMatchTheClient(t *testing.T) {
	assertRanks(t, "Blessing of Wisdom", loadClass(t, "paladin"), core.BlessingOfWisdomRanks, auraPeriodicEnergize, 0, false)
}

func TestMarkOfTheWildRanksMatchTheClient(t *testing.T) {
	client := loadClass(t, "druid")
	assertRanks(t, "Mark of the Wild armor", client, core.MarkOfTheWildArmorRanks, auraModResistance, miscArmor, false)
	assertRanks(t, "Mark of the Wild stats", client, core.MarkOfTheWildStatRanks, auraModStat, miscAllStats, false)
	for _, mask := range resistanceMasks {
		assertRanks(t, "Mark of the Wild resistance", client, core.MarkOfTheWildResistRanks, auraModResistance, mask, false)
	}
}

func TestBattleShoutRanksMatchTheClient(t *testing.T) {
	assertRanks(t, "Battle Shout", loadClass(t, "warrior"), core.BattleShoutRankTable, auraModAttackPower, 0, true)
}

func TestTrueshotAuraRanksMatchTheClient(t *testing.T) {
	assertRanks(t, "Trueshot Aura", loadClass(t, "hunter"), core.TrueshotAuraRanks, auraModRangedAttackPower, 0, false)
}

func TestDevotionAuraRanksMatchTheClient(t *testing.T) {
	assertRanks(t, "Devotion Aura", loadClass(t, "paladin"), core.DevotionAuraRanks, auraModResistance, miscArmor, false)
}

// Heart of the Lion's area buff (409583) states melee (aura 99) and ranged
// (aura 124) attack power with misc -1, 40 base points and 4 per level,
// capped at the spell's level 60; the hunter's own spell (409580) and the
// area buff both state aura 137 (mod total stat percentage) at +10 on all
// stats.
func TestHeartOfTheLionMatchesTheClient(t *testing.T) {
	client := loadClass(t, "hunter")
	assertRanks(t, "Heart of the Lion melee attack power", client, core.HeartOfTheLionRanks, auraModAttackPower, miscAllStats, true)
	assertRanks(t, "Heart of the Lion ranged attack power", client, core.HeartOfTheLionRanks, auraModRangedAttackPower, miscAllStats, true)

	for _, id := range []int32{core.HeartOfTheLionSpellID, core.HeartOfTheLionRanks[0].SpellID} {
		spell, ok := client.ByID(id)
		if !ok {
			t.Fatalf("Heart of the Lion %d is not in the client table", id)
		}
		stats := findEffect(spell, auraModTotalStatPercentage, miscAllStats)
		if want := core.HeartOfTheLionStatMultiplier; 1+stats.Amount/100 != want {
			t.Errorf("Heart of the Lion %d: client stat percentage %v, engine multiplier %v", id, stats.Amount, want)
		}
	}
}

// Flametongue Totem (fork sim/core/flametongue_totem.go): each rank's cast
// spell is learned at the rank's level, and the proc spell it triggers
// states the rank's amount as the base points of a dummy effect (effect 3,
// no aura), which the engine multiplies by weapon speed over 100.
func TestFlametongueTotemRanksMatchTheClient(t *testing.T) {
	client := loadClass(t, "shaman")
	for i, rank := range core.FlametongueTotemRanks {
		cast, ok := client.ByID(rank.SpellID)
		if !ok {
			t.Fatalf("Flametongue Totem %d is not in the client table", rank.SpellID)
		}
		if cast.SpellLevel != rank.Level {
			t.Errorf("Flametongue Totem %d: engine learns at %d, client at %d", rank.SpellID, rank.Level, cast.SpellLevel)
		}
		procID := core.FlametongueTotemProcSpellIDs[i]
		proc, ok := client.ByID(procID)
		if !ok {
			t.Fatalf("Flametongue Totem proc %d is not in the client table", procID)
		}
		if amount := findEffect(proc, 0, 0).Amount; amount != rank.Amount {
			t.Errorf("Flametongue Totem proc %d: engine amount %v, client %v", procID, rank.Amount, amount)
		}
	}
}

// Judgement of Light: each rank's judgement aura is learned at the rank's
// level and the heal spell it names (20267 / 20341 / 20342 / 20343, effect
// 10) states the amount. Judgement of Wisdom's mana spells (20268 / 20352 /
// 20353, effect 30) sit outside the spellconst extract, so only the
// judgement auras' levels are pinned here; the fork's judgement_ranks_test
// quotes the energize rows.
func TestJudgementRanksMatchTheClient(t *testing.T) {
	client := loadClass(t, "paladin")
	lightHeals := map[int32]int32{20185: 20267, 20344: 20341, 20345: 20342, 20346: 20343}
	for _, rank := range core.JudgementOfLightRanks {
		aura, ok := client.ByID(rank.SpellID)
		if !ok || aura.SpellLevel != rank.Level {
			t.Errorf("Judgement of Light %d: client level %d (found %v), engine %d", rank.SpellID, aura.SpellLevel, ok, rank.Level)
		}
		heal, ok := client.ByID(lightHeals[rank.SpellID])
		if !ok {
			t.Fatalf("Judgement of Light heal %d is not in the client table", lightHeals[rank.SpellID])
		}
		if amount := findEffect(heal, 0, 0).Amount; amount != rank.Amount {
			t.Errorf("Judgement of Light heal %d: engine amount %v, client %v", heal.ID, rank.Amount, amount)
		}
	}
	for _, rank := range core.JudgementOfWisdomRanks {
		aura, ok := client.ByID(rank.SpellID)
		if !ok || aura.SpellLevel != rank.Level {
			t.Errorf("Judgement of Wisdom %d: client level %d (found %v), engine %d", rank.SpellID, aura.SpellLevel, ok, rank.Level)
		}
	}
}
