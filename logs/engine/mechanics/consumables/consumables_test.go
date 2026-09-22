// logs/engine/mechanics/consumables/consumables_test.go
package consumables

import "testing"

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
