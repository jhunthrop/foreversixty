package main

import "testing"

func TestScoreStatsOnly(t *testing.T) {
	c := candidate{Stats: map[string]float64{"agility": 10, "crit": 5, "stamina": 20}}
	weights := map[string]float64{"agility": 2.0, "crit": 1.5, "hit": 3.0}
	// stamina has no weight entry, so weights["stamina"] is the zero
	// value and contributes nothing - an unweighted stat scores zero
	// rather than panicking on a missing map key.
	got := score(c, "chest", weights, 0, false)
	want := 10*2.0 + 5*1.5
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreWeaponDPSConvertsThroughAttackPowerPerDPS(t *testing.T) {
	c := candidate{DPS: 20, Stats: map[string]float64{"agility": 4}}
	weights := map[string]float64{"ranged_attack_power": 0.5, "agility": 2.0}
	got := score(c, "ranged", weights, 0, false)
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
	got := score(c, "trinket1", weights, 0, false)
	if got != 0 {
		t.Errorf("score = %v, want 0 (trinket1 is not in weaponAPStat)", got)
	}
}

func TestScoreMainHandAndOffHandUseMeleeAttackPower(t *testing.T) {
	c := candidate{DPS: 10}
	weights := map[string]float64{"attack_power": 1.0}
	for _, slot := range []string{"main_hand", "off_hand"} {
		got := score(c, slot, weights, 0, false)
		want := 10 * attackPowerPerDPS * 1.0
		if got != want {
			t.Errorf("score(%q) = %v, want %v", slot, got, want)
		}
	}
}

func TestScoreZeroDPSAddsNothing(t *testing.T) {
	c := candidate{DPS: 0, Stats: map[string]float64{"agility": 1}}
	weights := map[string]float64{"agility": 1, "ranged_attack_power": 100}
	got := score(c, "ranged", weights, 0, false)
	if got != 1 {
		t.Errorf("score = %v, want 1 (no DPS contribution)", got)
	}
}

func TestScoreGenericAttackPowerCountsTowardRangedAttackPowerWeight(t *testing.T) {
	// Blackwater Belt-shaped case: items/<class>.json states a generic AP
	// bonus as one "attack_power" line (no separate "ranged_attack_power"
	// key), but Classic's generic AP raises ranged attack power too - a
	// hunter dps spec's weights (ranged_attack_power, no attack_power entry
	// per data/curated/specs.json) must still value it.
	c := candidate{Stats: map[string]float64{"attack_power": 18}}
	weights := map[string]float64{"ranged_attack_power": 3.0}
	got := score(c, "waist", weights, 0, false)
	want := 18 * 3.0
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreGenericAttackPowerStillCountsTowardMeleeAttackPowerWeight(t *testing.T) {
	c := candidate{Stats: map[string]float64{"attack_power": 18}}
	weights := map[string]float64{"attack_power": 2.0}
	got := score(c, "waist", weights, 0, false)
	want := 18 * 2.0
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreGenericAttackPowerDoesNotDoubleCountWhenASpecWeighsBoth(t *testing.T) {
	// No spec's weight_stats actually lists both today (see
	// genericAttackPowerWeightStats' comment), but the sum still has to be
	// the physically correct one if that ever changes: a single generic AP
	// bonus is one physical quantity, valued once per weight it feeds.
	c := candidate{Stats: map[string]float64{"attack_power": 10}}
	weights := map[string]float64{"attack_power": 1.0, "ranged_attack_power": 4.0}
	got := score(c, "waist", weights, 0, false)
	want := 10 * (1.0 + 4.0)
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreCasterWandDPSFallsBackToReferenceDPSPerPointWhenSpecCastsShoot(t *testing.T) {
	// bis-ranker-integrity-4 lane, caster sweep item 2: a caster spec's
	// weight_stats never carries attack_power/ranged_attack_power
	// (specs.json, confirmed against every one of mage/priest/warlock/
	// shaman-elemental/druid-balance's own rows), so weights[apStat] is
	// the Go zero value here exactly as it is in production - a
	// stat-less wand's DPS must still count for a spec whose rotation
	// actually fires Shoot, converted through referenceDPSPerPoint
	// instead of an AP weight that will never exist for this spec.
	c := candidate{DPS: 10}
	weights := map[string]float64{"spell_power": 2.0}
	referenceDPSPerPoint := 0.5
	got := score(c, "ranged", weights, referenceDPSPerPoint, true)
	want := 10 / referenceDPSPerPoint
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreCasterWandFallbackAddsToStats(t *testing.T) {
	c := candidate{DPS: 10, Stats: map[string]float64{"spell_power": 5}}
	weights := map[string]float64{"spell_power": 2.0}
	referenceDPSPerPoint := 0.5
	got := score(c, "ranged", weights, referenceDPSPerPoint, true)
	want := 5*2.0 + 10/referenceDPSPerPoint
	if got != want {
		t.Errorf("score = %v, want %v", got, want)
	}
}

func TestScoreWandFallbackNeverFiresWithoutCastsShoot(t *testing.T) {
	// bis-ranker-integrity-4 lane, caster sweep item 2: shaman-elemental
	// and druid-balance also carry zero attack_power/ranged_attack_power
	// weight, exactly like every true caster, but their own rotations
	// never cast Shoot (confirmed against data/curated/apl/shaman-
	// elemental.json and druid-balance.json: no spell id 5019 anywhere in
	// either) - a wand's raw DPS must not count for them at all, the
	// same as it never counted before this fallback existed.
	c := candidate{DPS: 10}
	weights := map[string]float64{"spell_power": 2.0}
	got := score(c, "ranged", weights, 0.5, false)
	if got != 0 {
		t.Errorf("score = %v, want 0 (castsShoot is false)", got)
	}
}

func TestScoreMainHandOrOffHandWeaponDPSNeverCountsWithoutAPWeight(t *testing.T) {
	// The real defect the caster sweep found (bis-ranker-integrity-4
	// lane, item 2): Manual Crowd Pummeler (Str/Agi, a melee-haste
	// proc, zero caster stats) was outscoring Staff of Jordan (real
	// spell power) at band 30/40 for shaman-elemental and druid-balance
	// because the old fallback treated ANY zero-AP-weight weapon slot's
	// flat DPS as free score, main_hand/off_hand included. A caster
	// never swings their melee weapon: main_hand/off_hand must score
	// from Stats alone, regardless of castsShoot or how large
	// referenceDPSPerPoint is - only the ranged slot's wand is ever a
	// real damage source for these specs.
	for _, slot := range []string{"main_hand", "off_hand"} {
		for _, castsShoot := range []bool{true, false} {
			c := candidate{DPS: 50, Stats: map[string]float64{"strength": 16, "agility": 5}}
			weights := map[string]float64{"spell_power": 2.0}
			got := score(c, slot, weights, 0.5, castsShoot)
			if got != 0 {
				t.Errorf("score(%q, castsShoot=%v) = %v, want 0 (a caster's melee weapon DPS must never count, strength/agility are unweighted)", slot, castsShoot, got)
			}
		}
	}
}

func TestScoreWeaponFallbackNeverFiresWhenAPWeightIsPositive(t *testing.T) {
	// A melee/hybrid/hunter spec always has a positive apStat weight in
	// production (weightsRequest's own doc: a bare, weapon-only ladder
	// character would otherwise report ErrNoWeights) - the fallback
	// must never ALSO add referenceDPSPerPoint's conversion on top of
	// the real AP-weighted term, which would double-count the same
	// weapon's damage twice.
	c := candidate{DPS: 10}
	weights := map[string]float64{"attack_power": 1.0}
	got := score(c, "main_hand", weights, 0.5, false)
	want := 10 * attackPowerPerDPS * 1.0
	if got != want {
		t.Errorf("score = %v, want %v (fallback must not also fire)", got, want)
	}
}

func TestScoreWeaponFallbackIsNoOpWhenReferenceDPSPerPointIsZero(t *testing.T) {
	// Every existing test (and any caller that has not yet run a
	// weights sim, e.g. a unit test constructing weights by hand) passes
	// 0 - the old behaviour (weapon DPS contributes nothing for a
	// caster) must hold exactly, not divide by zero.
	c := candidate{DPS: 10}
	weights := map[string]float64{"spell_power": 2.0}
	got := score(c, "ranged", weights, 0, true)
	if got != 0 {
		t.Errorf("score = %v, want 0 (no reference DPS/point available)", got)
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
	got := score(c, "ranged", weights, 0, false)
	if got != 0 {
		t.Errorf("score = %v, want 0 for a zero-stat relic", got)
	}
}
