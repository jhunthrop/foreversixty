package main

import "testing"

func TestScoreStatsOnly(t *testing.T) {
	c := candidate{Stats: map[string]float64{"agility": 10, "crit": 5, "stamina": 20}}
	weights := map[string]float64{"agility": 2.0, "crit": 1.5, "hit": 3.0}
	// stamina has no weight entry, so weights["stamina"] is the zero
	// value and contributes nothing - an unweighted stat scores zero
	// rather than panicking on a missing map key.
	got := score(c, "chest", weights)
	want := 10*2.0 + 5*1.5
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreWeaponDPSConvertsThroughAttackPowerPerDPS(t *testing.T) {
	c := candidate{DPS: 20, Stats: map[string]float64{"agility": 4}}
	weights := map[string]float64{"ranged_attack_power": 0.5, "agility": 2.0}
	got := score(c, "ranged", weights)
	want := 4*2.0 + 20*attackPowerPerDPS*0.5
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreWeaponDPSIgnoredInANonWeaponSlot(t *testing.T) {
	// A non-weapon slot (e.g. a trinket that happens to carry a DPS
	// field of 0 in practice, but this pins the general case) never
	// consults weaponAPStat.
	c := candidate{DPS: 20}
	weights := map[string]float64{"attack_power": 1.0, "ranged_attack_power": 1.0}
	got := score(c, "trinket1", weights)
	if got != 0 {
		t.Errorf("score = %v, want 0 (trinket1 is not in weaponAPStat)", got)
	}
}

func TestScoreMainHandAndOffHandUseMeleeAttackPower(t *testing.T) {
	c := candidate{DPS: 10}
	weights := map[string]float64{"attack_power": 1.0}
	for _, slot := range []string{"main_hand", "off_hand"} {
		got := score(c, slot, weights)
		want := 10 * attackPowerPerDPS * 1.0
		if got != want {
			t.Errorf("score(%q) = %v, want %v", slot, got, want)
		}
	}
}

func TestScoreZeroDPSAddsNothing(t *testing.T) {
	c := candidate{DPS: 0, Stats: map[string]float64{"agility": 1}}
	weights := map[string]float64{"agility": 1, "ranged_attack_power": 100}
	got := score(c, "ranged", weights)
	if got != 1 {
		t.Errorf("score = %v, want 1 (no DPS contribution)", got)
	}
}

func TestScoreARelicWithNoStatsAndNoDPSIsZero(t *testing.T) {
	// A relic (libram/idol/totem) carries no armour, no flat stat and no
	// DPS -- its whole value is an on-equip spell effect score() has no
	// term for at all (night-relic-exempt: relics were dropped entirely
	// before this, so score() never saw one). It must score exactly 0,
	// not panic on a nil/empty Stats map, so only rank.go's engine-
	// verified effect ranking (or, absent that, the effect_unmodelled
	// flag) can ever distinguish one relic from another.
	c := candidate{ClassID: armorClassID, SubclassID: armorSublibramID, Stats: map[string]float64{}}
	weights := map[string]float64{"agility": 2.0, "ranged_attack_power": 100}
	got := score(c, "ranged", weights)
	if got != 0 {
		t.Errorf("score = %v, want 0 for a zero-stat relic", got)
	}
}
