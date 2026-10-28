// logs/engine/mechanics/consumables/consumables_test.go
package consumables

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoadHasAllThreeRolesWithWeightsSummingToOneHundred(t *testing.T) {
	c := Load()
	for _, role := range []string{"dps", "healer", "tank"} {
		rc, ok := c.For(role)
		if !ok {
			t.Fatalf("no catalogue for role %q", role)
		}
		if rc.weightSum() != 100 {
			t.Errorf("%s weights sum to %d, want 100", role, rc.weightSum())
		}
	}
	if _, ok := c.For("not-a-role"); ok {
		t.Fatal("an unknown role must report false")
	}
}

func TestHealerHasNoWeaponEnchantWeight(t *testing.T) {
	rc, ok := Load().For("healer")
	if !ok {
		t.Fatal("no healer catalogue")
	}
	if rc.Weights["weapon_enchant"] != 0 || len(rc.WeaponEnchant) != 0 {
		t.Fatalf("healer weapon_enchant = weight %d, entries %v, want 0/[]", rc.Weights["weapon_enchant"], rc.WeaponEnchant)
	}
}

func TestParseRejectsWeightsThatDoNotSumToOneHundred(t *testing.T) {
	bad := []byte(`{"roles": {"dps": {"weights": {"flask": 50}, "flask": [], "food": [], "weapon_enchant": [], "world_buffs": [], "combat_potion": {"max_uses": 1, "entries": []}}}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("weights that do not sum to 100 must be refused")
	}
}

func TestParseRejectsAnEntryWithNoVerifiedSource(t *testing.T) {
	bad := []byte(`{"roles": {"dps": {"weights": {"flask": 100}, "flask": [{"spell_id": 1, "name": "X"}], "food": [], "weapon_enchant": [], "world_buffs": [], "combat_potion": {"max_uses": 1, "entries": []}}}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an entry with no verified source must be refused")
	}
}

func TestParseRejectsAnEntryWithNoSpellID(t *testing.T) {
	bad := []byte(`{"roles": {"dps": {"weights": {"flask": 100}, "flask": [{"spell_id": 0, "name": "X", "verified": "y"}], "food": [], "weapon_enchant": [], "world_buffs": [], "combat_potion": {"max_uses": 1, "entries": []}}}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an entry with no positive spell_id must be refused")
	}
}

// validRole is a role body that Parse accepts, for the raid_buffs and
// complete refusals below to differ from it in one field only.
const validRole = `"dps": {"weights": {"flask": 100}, "flask": [], "food": [], "weapon_enchant": [], "world_buffs": [], "combat_potion": {"max_uses": 1, "entries": []}%s}`

func catalogueWith(roleExtra, top string) []byte {
	return []byte(`{"roles": {` + fmt.Sprintf(validRole, roleExtra) + `}` + top + `}`)
}

func TestParse_AcceptsTheValidBase(t *testing.T) {
	good := catalogueWith(`, "complete": {"flask": true}`,
		`, "raid_buffs": [{"name": "Arcane Intellect", "class": "Mage", "spell_ids": [1459], "verified": "x"}]`)
	if _, err := Parse(good); err != nil {
		t.Fatalf("valid catalogue refused: %v", err)
	}
}

func TestParse_RefusesUnstampedRaidBuff(t *testing.T) {
	bad := catalogueWith("", `, "raid_buffs": [{"name": "Arcane Intellect", "class": "Mage", "spell_ids": [1459], "verified": ""}]`)
	_, err := Parse(bad)
	if err == nil || !strings.Contains(err.Error(), "Arcane Intellect") {
		t.Fatalf("unstamped raid buff: err = %v, want an error naming it", err)
	}
}

func TestParse_RefusesMalformedRaidBuffs(t *testing.T) {
	cases := map[string]string{
		"no class":     `{"name": "B", "class": "", "spell_ids": [1], "verified": "x"}`,
		"no spell ids": `{"name": "B", "class": "Mage", "spell_ids": [], "verified": "x"}`,
		"zero id":      `{"name": "B", "class": "Mage", "spell_ids": [1, 0], "verified": "x"}`,
		"negative id":  `{"name": "B", "class": "Mage", "spell_ids": [-4], "verified": "x"}`,
	}
	for name, buff := range cases {
		if _, err := Parse(catalogueWith("", `, "raid_buffs": [`+buff+`]`)); err == nil {
			t.Errorf("%s: raid buff must be refused", name)
		}
	}
}

func TestParse_RefusesUnknownCompleteKey(t *testing.T) {
	if _, err := Parse(catalogueWith(`, "complete": {"elixir": true}`, "")); err == nil {
		t.Fatal(`"complete": {"elixir": true} must be refused`)
	}
}

func TestCatalogue_FoodAndWeaponEnchantNotComplete(t *testing.T) {
	c := Load()
	for role, rc := range c.Roles {
		if rc.IsComplete(CategoryFood) || rc.IsComplete(CategoryWeaponEnchant) {
			t.Errorf("%s: food/weapon_enchant must stay incomplete until confirmed", role)
		}
		for _, cat := range []string{CategoryFlask, CategoryCombatPotion, CategoryWorldBuffs} {
			if role == "dps" && cat == CategoryCombatPotion {
				continue // only Mighty Rage is listed: casters, hunters and rogues drink others
			}
			if !rc.IsComplete(cat) {
				t.Errorf("%s: %s should start complete", role, cat)
			}
		}
	}
	if c.Roles["dps"].IsComplete(CategoryCombatPotion) {
		t.Fatal("dps combat_potion must stay incomplete until it lists the mana and other potions")
	}
	if c.AllRoles().IsComplete(CategoryFood) {
		t.Fatal("AllRoles must not call food complete")
	}
}

func TestCatalogue_RaidBuffsCoverTheSevenFamilies(t *testing.T) {
	want := map[string]string{
		"Power Word: Fortitude": "Priest", "Arcane Intellect": "Mage", "Mark of the Wild": "Druid",
		"Blessing of Might": "Paladin", "Blessing of Wisdom": "Paladin", "Blessing of Kings": "Paladin",
		"Battle Shout": "Warrior",
	}
	got := map[string]string{}
	for _, b := range Load().RaidBuffs {
		got[b.Name] = b.Class
	}
	if len(got) != len(want) {
		t.Fatalf("raid buffs = %v, want %v", got, want)
	}
	for name, class := range want {
		if got[name] != class {
			t.Errorf("%s: class %q, want %q", name, got[name], class)
		}
	}
}

func TestAllRoles_MergesEveryRolesEntriesOnce(t *testing.T) {
	all := Load().AllRoles()
	ids := map[int64]int{}
	for _, e := range all.Flask {
		ids[e.SpellID]++
	}
	// Titans (dps, tank), Supreme Power (dps), Distilled Wisdom and Chromatic Resistance (healer, tank)
	for _, id := range []int64{17626, 17627, 17628, 17629} {
		if ids[id] != 1 {
			t.Errorf("flask %d appears %d times in AllRoles, want 1", id, ids[id])
		}
	}
	potions := map[int64]bool{}
	for _, e := range all.CombatPotion.Entries {
		potions[e.SpellID] = true
	}
	if !potions[17528] || !potions[17531] {
		t.Fatalf("AllRoles potions = %v, want rage and mana potions", all.CombatPotion.Entries)
	}
}
