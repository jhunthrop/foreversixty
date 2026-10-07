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
