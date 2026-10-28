// logs/engine/rating/preparation_test.go
package rating

import (
	"fmt"
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
		WorldBuffs:    []consumables.Entry{{SpellID: 22888, Name: "Rallying Cry"}, {SpellID: 24425, Name: "Spirit of Zandalar"}},
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

// A sheet whose every list is incomplete, with nothing listed up, has
// nothing to judge: unknown, never a zero.
func TestPreparation_NothingJudgeableIsNotAZero(t *testing.T) {
	every := testCatalogue(consumables.CategoryFlask, consumables.CategoryFood, consumables.CategoryWeaponEnchant,
		consumables.CategoryWorldBuffs, consumables.CategoryCombatPotion)
	if v, ok := PreparationValue(preparedFight(), fixturePlayer, every); ok {
		t.Fatalf("PreparationValue = %v, true; want no value", v)
	}
}

// potionCast is one successful cast of spellID by the fixture player.
func potionCast(spellID int64) summary.CastRow {
	return summary.CastRow{GUID: fixturePlayer, OwnerGUID: fixturePlayer, SpellID: spellID, Succeeded: 1}
}

// An incomplete category can add credit, never a shortfall: what it earned
// counts in both the earned and the possible total, and what it did not
// earn counts in neither.
func TestPreparation_AnIncompleteCategoryCountsOnlyWhatItEarned(t *testing.T) {
	flask := preparedFight(seededAura(17626))
	checkPreparation(t, "off-list food: 30 of 85", flask, testCatalogue(consumables.CategoryFood), 35.29)
	fed := preparedFight(seededAura(17626), seededAura(1249522))
	checkPreparation(t, "listed food earns: 45 of 100", fed, testCatalogue(consumables.CategoryFood), 45)
	checkPreparation(t, "an unreadable stone: 45 of 85", fed, testCatalogue(consumables.CategoryWeaponEnchant), 52.94)
	checkPreparation(t, "a complete stone list: 45 of 100", fed, testCatalogue(), 45)
	checkPreparation(t, "no use of a partial potion list: 45 of 80", fed, testCatalogue(consumables.CategoryCombatPotion), 56.25)
	quaffed := preparedFight(seededAura(17626), seededAura(1249522))
	quaffed.Casts = []summary.CastRow{potionCast(17528)}
	checkPreparation(t, "one use of a partial potion list: 55 of 90", quaffed, testCatalogue(consumables.CategoryCombatPotion), 61.11)
}

// The spec's worked examples (Fury 70, Holy 45; score_test.go) assume every
// list is complete. The shipped catalogue marks food, weapon_enchant and
// combat_potion incomplete for every role, so the same pulls read higher:
// what those lists did not see is unknown, not missing. These pin the
// shipped values so a catalogue change that moves them is seen.
func TestPreparation_ShippedCatalogueValues(t *testing.T) {
	healer, _ := consumables.Load().For(RoleHealer)
	quaffed := preparedFight(seededAura(17626), seededAura(1249522), seededAura(25122))
	quaffed.Casts = []summary.CastRow{potionCast(17528)}
	for _, c := range []struct {
		name  string
		fight summary.Summary
		want  float64
	}{
		{"dps flask only: 30 of 50", preparedFight(seededAura(17626)), 60},
		{"dps flask, food, oil: 60 of 80", preparedFight(seededAura(17626), seededAura(1249522), seededAura(25122)), 75},
		{"dps flask, food, oil, one potion: 70 of 90", quaffed, 77.78},
	} {
		if got := placedPreparation(t, c.fight); len(got) != 1 || roundTo(got[0], 2) != c.want {
			t.Errorf("%s: placed %v, want [%v]", c.name, got, c.want)
		}
	}
	holy := preparedFight(seededAura(healer.Flask[0].SpellID), seededAura(healer.Food[0].SpellID))
	checkPreparation(t, "Holy's worked example on the shipped catalogue: 45 of 65", holy, &healer, 69.23)
}

// pullItem is one thing a player can bring, added to a fight.
type pullItem struct {
	name string
	add  func(*summary.Summary)
}

// pullItemsPerCategory caps the entries each list contributes to the exhaustive
// subset walk: a category credits any one entry, so two cover "first of its
// kind" and "one more of a kind already there", and the walk stays 2^n over a
// small n now that a flask list carries every flask the build ships (18).
// TestPreparation_EveryEntryCountsLikeTheFirst covers the rest one at a time.
const pullItemsPerCategory = 2

// pullItemCategories are the lists a pull is credited from, in order.
var pullItemCategories = []string{consumables.CategoryFlask, consumables.CategoryFood,
	consumables.CategoryWeaponEnchant, consumables.CategoryWorldBuffs, consumables.CategoryCombatPotion}

// pullItems is what rc can credit: up to pullItemsPerCategory entries of every
// list seeded on the pull, plus two casts of the first potion (the second use).
func pullItems(rc consumables.RoleCatalogue) []pullItem {
	var items []pullItem
	for _, category := range pullItemCategories {
		entries := rc.Entries(category)
		for _, e := range entries[:min(len(entries), pullItemsPerCategory)] {
			id := e.SpellID
			items = append(items, pullItem{category + " " + e.Name,
				func(f *summary.Summary) { f.Auras = append(f.Auras, seededAura(id)) }})
		}
	}
	potion := rc.CombatPotion.Entries[0].SpellID
	for use := 1; use <= 2; use++ {
		items = append(items, pullItem{fmt.Sprintf("potion cast %d", use),
			func(f *summary.Summary) { f.Casts = append(f.Casts, potionCast(potion)) }})
	}
	return items
}

func fightWith(items []pullItem, mask int) summary.Summary {
	fight := preparedFight()
	for i, it := range items {
		if mask&(1<<i) != 0 {
			it.add(&fight)
		}
	}
	return fight
}

// Bringing one more thing never lowers Preparation: exhaustive over every
// subset of the shipped catalogue's entries (plus two potion casts) for each
// role, and on every completeness mix of the test catalogue.
func TestPreparation_BringingMoreNeverLowersTheScore(t *testing.T) {
	cats := map[string]*consumables.RoleCatalogue{}
	for _, role := range []string{RoleDPS, RoleHealer, RoleTank} {
		rc, _ := consumables.Load().For(role)
		cats["shipped "+role] = &rc
	}
	for _, incomplete := range [][]string{nil, {consumables.CategoryFood},
		{consumables.CategoryFood, consumables.CategoryWeaponEnchant, consumables.CategoryCombatPotion}} {
		cats[fmt.Sprintf("test, incomplete %v", incomplete)] = testCatalogue(incomplete...)
	}
	for name, cat := range cats {
		items := pullItems(*cat)
		for mask := 0; mask < 1<<len(items); mask++ {
			before, _ := preparationTotal(fightWith(items, mask), fixturePlayer, cat)
			for i, it := range items {
				if mask&(1<<i) != 0 {
					continue
				}
				after, _ := preparationTotal(fightWith(items, mask|1<<i), fixturePlayer, cat)
				if after < before-1e-9 {
					t.Fatalf("%s: adding %s to mask %b lowered Preparation %.2f -> %.2f", name, it.name, mask, before, after)
				}
			}
		}
	}
}

// Every shipped entry, seeded alone on the pull, scores exactly as its list's
// first entry does: the exhaustive walk above samples two per list.
func TestPreparation_EveryEntryCountsLikeTheFirst(t *testing.T) {
	for _, role := range []string{RoleDPS, RoleHealer, RoleTank} {
		rc, _ := consumables.Load().For(role)
		for _, category := range pullItemCategories {
			entries := rc.Entries(category)
			if len(entries) == 0 {
				continue
			}
			want, _ := preparationTotal(preparedFight(seededAura(entries[0].SpellID)), fixturePlayer, &rc)
			for _, e := range entries[1:] {
				if got, _ := preparationTotal(preparedFight(seededAura(e.SpellID)), fixturePlayer, &rc); got != want {
					t.Errorf("%s %s: %s scores %.2f, the first entry %.2f", role, category, e.Name, got, want)
				}
			}
		}
	}
}

func TestModelVersionMovedForConsumables(t *testing.T) {
	if DefaultModelVersion != "rating-2026-10-28-consumables-at-pull" {
		t.Fatal(DefaultModelVersion)
	}
}
