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

// TestPickGearItemPrefersDPSOverRawStatPointsForAPhysicalSpec pins the
// regression this lane's own first cut of weaponScore hit: paladin
// item 21134 (Dark Edge of Insanity - 35 strength, 19 agility, 86.6 DPS)
// against item 12784 (Arcanite Reaper - 62 attack_power, 53.8 DPS),
// shaped directly from data/builds/1.60.1.70009/items/paladin.json.
// Summing raw stat points (35+19=54 against 62) picks the Reaper - the
// wrong weapon, since strength and agility do not convert to attack
// power 1:1 and the comparison ignores the Edge's far higher DPS
// entirely. A physical spec's score must rank DPS first so this stays
// the Edge, matching what the item-level-only rule (correctly, by
// accident) picked before this lane's fix.
func TestPickGearItemPrefersDPSOverRawStatPointsForAPhysicalSpec(t *testing.T) {
	items := []buildItem{
		{
			ID: 12784, Slot: "main_hand", ItemLevel: 63,
			WeaponClass: itemClassWeapon, Speed: 3.8, DamageMax: 256, DPS: 53.82,
			Stats: map[string]float64{"stamina": 13, "attack_power": 62},
		},
		{
			ID: 21134, Slot: "main_hand", ItemLevel: 84,
			WeaponClass: itemClassWeapon, Speed: 3.5, DamageMax: 364, DPS: 86.57,
			Stats: map[string]float64{"strength": 35, "agility": 19, "stamina": 25},
		},
	}
	known := map[int]bool{12784: true, 21134: true}
	primaryStats := specWeaponPrimaryStats("paladin-retribution")
	weightDPS := specWeaponDPSMatters("paladin-retribution")

	got, ok := pickGearItem(items, known, nil, "main_hand", 60, handAny, nil, primaryStats, weightDPS)
	if !ok || got.ID != 21134 {
		t.Fatalf("pickGearItem(paladin-retribution) = (%+v, %v), want item 21134 (Dark Edge of Insanity, the far higher-DPS weapon)", got, ok)
	}
}
