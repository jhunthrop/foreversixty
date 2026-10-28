// logs/engine/rating/preparation_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// preparedFight is fixtureSummary with a character sheet for the fixture
// player and the given aura tracks on them.
func preparedFight(auras ...summary.AuraTrack) summary.Summary {
	fight := fixtureSummary()
	fight.Combatants = []summary.CombatantRow{{GUID: fixturePlayer, Gear: []event.Item{{ID: 1}}}}
	fight.Auras = auras
	return fight
}

// seededOn is an aura the player's own COMBATANT_INFO listed, up the whole
// 180-second fixture fight.
func seededOn(guid string, spellID int64) summary.AuraTrack {
	return summary.AuraTrack{TargetGUID: guid, SpellID: spellID, Seeded: true,
		Segments: []summary.Segment{{StartMS: 0, EndMS: 180000}}}
}

func seededAura(spellID int64) summary.AuraTrack { return seededOn(fixturePlayer, spellID) }

// checkPlaced asserts the one raw value component placed through src.
func checkPlaced(t *testing.T, src *fakePercentiles, component string, want float64) {
	t.Helper()
	if got := src.recorded[component]; len(got) != 1 || roundTo1(got[0]) != want {
		t.Errorf("%s placed %v, want [%v]", component, got, want)
	}
}

// everyCategoryComplete marks every list complete, as the spec's worked
// examples assume: anything not up is missing.
func everyCategoryComplete() map[string]bool {
	return map[string]bool{
		consumables.CategoryFlask: true, consumables.CategoryFood: true, consumables.CategoryWeaponEnchant: true,
		consumables.CategoryWorldBuffs: true, consumables.CategoryCombatPotion: true,
	}
}

// testCatalogue is the spec's dps weights over one entry per category,
// with the named categories marked incomplete (every other one complete).
func testCatalogue(incomplete ...string) *consumables.RoleCatalogue {
	complete := everyCategoryComplete()
	for _, c := range incomplete {
		complete[c] = false
	}
	return &consumables.RoleCatalogue{
		Weights:       map[string]int{"flask": 30, "food": 15, "weapon_enchant": 15, "world_buffs": 20, "combat_potion": 20},
		Complete:      complete,
		Flask:         []consumables.Entry{{SpellID: 17626, Name: "Flask of the Titans"}},
		Food:          []consumables.Entry{{SpellID: 1249522, Name: "Grilled Squid"}},
		WeaponEnchant: []consumables.Entry{{SpellID: 25122, Name: "Brilliant Wizard Oil"}},
		WorldBuffs:    []consumables.Entry{{SpellID: 22888, Name: "Rallying Cry"}},
		CombatPotion:  consumables.PotionGroup{MaxUses: 2, Entries: []consumables.Entry{{SpellID: 17528, Name: "Mighty Rage Potion"}}},
	}
}

// placedWith is what scorePreparation placed for cat against fight, and the
// component it returned.
func placedWith(t *testing.T, fight summary.Summary, cat *consumables.RoleCatalogue) ([]float64, Component) {
	t.Helper()
	src := &fakePercentiles{}
	c := scorePreparation(fight, fixturePlayer, Bracket{}, cat, src)
	return src.recorded[ComponentNamePreparation], c
}

// placedPreparation is what scorePreparation placed for the embedded dps
// catalogue against fight.
func placedPreparation(t *testing.T, fight summary.Summary) []float64 {
	t.Helper()
	rc, ok := consumables.Load().For(RoleDPS)
	if !ok {
		t.Fatal("no dps catalogue")
	}
	got, _ := placedWith(t, fight, &rc)
	return got
}

func checkPreparation(t *testing.T, name string, fight summary.Summary, cat *consumables.RoleCatalogue, want float64) {
	t.Helper()
	if got, _ := placedWith(t, fight, cat); len(got) != 1 || roundTo(got[0], 2) != want {
		t.Errorf("%s: placed %v, want [%v]", name, got, want)
	}
}

func TestPreparation_SeededFlaskEarnsFlaskWeight(t *testing.T) {
	checkPreparation(t, "a seeded flask, nothing else", preparedFight(seededAura(17626)), testCatalogue(), 30)
}

func TestPreparation_FlaskDrunkAfterTheGraceEarnsNothing(t *testing.T) {
	late := summary.AuraTrack{TargetGUID: fixturePlayer, SpellID: 17626,
		Segments: []summary.Segment{{StartMS: consumables.AtPullGraceMS + 1, EndMS: 180000}}}
	checkPreparation(t, "a flask drunk mid-fight", preparedFight(late), testCatalogue(), 0)
}

// With no character sheet the pull is unknown, never empty: the component
// is excluded and nothing is placed, so no digest fills with zeros.
func TestPreparation_NoSheetIsExcluded(t *testing.T) {
	fight := preparedFight(seededAura(17626))
	fight.Combatants = nil
	got, c := placedWith(t, fight, testCatalogue())
	if !c.Excluded || c.Reason != ReasonNoCharacterSheet || len(got) != 0 {
		t.Fatalf("component %+v placed %v, want excluded (%s) and nothing placed", c, got, ReasonNoCharacterSheet)
	}
	if _, ok := PreparationValue(fight, fixturePlayer, testCatalogue()); ok {
		t.Fatal("PreparationValue reported a value with no character sheet: the rebuild would fold a zero")
	}
}

func TestPreparation_PrePullPotionCountsAsOneUse(t *testing.T) {
	checkPreparation(t, "a seeded potion", preparedFight(seededAura(17528)), testCatalogue(), 10)
}

// An incomplete category can prove presence, never absence: absent, its
// weight drops out of the total and the rest is read out of what is left.
func TestPreparation_AnIncompleteCategoryAbsentIsUnknown(t *testing.T) {
	flask := preparedFight(seededAura(17626))
	checkPreparation(t, "off-list food: 30 of 85", flask, testCatalogue(consumables.CategoryFood), 35.29)
	fed := preparedFight(seededAura(17626), seededAura(1249522))
	checkPreparation(t, "listed food still earns: 45 of 100", fed, testCatalogue(consumables.CategoryFood), 45)
	checkPreparation(t, "an unreadable stone: 45 of 85", fed, testCatalogue(consumables.CategoryWeaponEnchant), 52.94)
	checkPreparation(t, "a complete stone list: 45 of 100", fed, testCatalogue(), 45)
	checkPreparation(t, "no use of a partial potion list: 45 of 80", fed, testCatalogue(consumables.CategoryCombatPotion), 56.25)
	quaffed := preparedFight(seededAura(17626), seededAura(1249522))
	quaffed.Casts = []summary.CastRow{{GUID: fixturePlayer, OwnerGUID: fixturePlayer, SpellID: 17528, Succeeded: 1}}
	checkPreparation(t, "one use of a partial potion list: 55 of 100", quaffed, testCatalogue(consumables.CategoryCombatPotion), 55)
}

func TestModelVersionMovedForConsumables(t *testing.T) {
	if DefaultModelVersion != "rating-2026-10-28-consumables-at-pull" {
		t.Fatal(DefaultModelVersion)
	}
}
