package request

import "testing"

// TestPickGearItemPrefersASpecsPrimaryStatOverRawItemLevel is harness
// rule 2 (this lane's brief): the old pickGearItem sorted purely on
// item_level, which is what handed a level-60 mage a stat-less
// higher-level weapon over a lower-level one carrying real spell power.
// A caster's own primary stat (specWeaponPrimaryStats, from the spec's
// WeightStats) now outranks item_level outright.
func TestPickGearItemPrefersASpecsPrimaryStatOverRawItemLevel(t *testing.T) {
	items := []buildItem{
		{
			// The higher-level pick under the old, item-level-only rule:
			// no useful stat for a caster at all.
			ID: 1, Slot: "main_hand", ItemLevel: 60,
			WeaponClass: itemClassWeapon, Speed: 2.0, DamageMax: 40,
			Stats: map[string]float64{"stamina": 20},
		},
		{
			// Lower item level, but carries the caster's own primary
			// stat - this is the pick weaponScore should now make.
			ID: 2, Slot: "main_hand", ItemLevel: 40,
			WeaponClass: itemClassWeapon, Speed: 2.0, DamageMax: 30,
			Stats: map[string]float64{"spell_power": 25},
		},
	}
	known := map[int]bool{1: true, 2: true}
	primaryStats := specWeaponPrimaryStats("mage-arcane")
	if len(primaryStats) == 0 {
		t.Fatalf("specWeaponPrimaryStats(mage-arcane) is empty - mage-arcane's WeightStats should carry spell_power")
	}

	got, ok := pickGearItem(items, known, nil, "main_hand", 60, handAny, nil, primaryStats, specWeaponDPSMatters("mage-arcane"))
	if !ok || got.ID != 2 {
		t.Fatalf("pickGearItem(mage-arcane) = (%+v, %v), want item 2 (the spell-power weapon, despite the lower item level)", got, ok)
	}
}

// TestPickGearItemWeightsWeaponDPSForAPhysicalSpec is the other half of
// harness rule 2: a melee/hunter spec's damage scales with the weapon's
// own raw output, so among two candidates that tie on primary stat (here
// neither carries one), the higher-DPS weapon wins even though it has
// the lower item_level - unlike a caster, whose pick (the test above)
// never looks at DPS at all.
func TestPickGearItemWeightsWeaponDPSForAPhysicalSpec(t *testing.T) {
	items := []buildItem{
		{
			ID: 1, Slot: "main_hand", ItemLevel: 50,
			WeaponClass: itemClassWeapon, Speed: 2.0, DamageMax: 40, DPS: 15.0,
		},
		{
			ID: 2, Slot: "main_hand", ItemLevel: 60,
			WeaponClass: itemClassWeapon, Speed: 3.0, DamageMax: 30, DPS: 10.0,
		},
	}
	known := map[int]bool{1: true, 2: true}
	primaryStats := specWeaponPrimaryStats("warrior-arms")
	weightDPS := specWeaponDPSMatters("warrior-arms")
	if !weightDPS {
		t.Fatalf("specWeaponDPSMatters(warrior-arms) = false, want true (a physical spec)")
	}

	got, ok := pickGearItem(items, known, nil, "main_hand", 60, handAny, nil, primaryStats, weightDPS)
	if !ok || got.ID != 1 {
		t.Fatalf("pickGearItem(warrior-arms) = (%+v, %v), want item 1 (the higher-DPS weapon, despite the lower item level)", got, ok)
	}
}
