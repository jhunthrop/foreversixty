package leveling

import "testing"

// The Fury guide build casts Whirlwind on cooldown and never Intercepts
// (data/curated/apl/warrior-fury.json). Improved Whirlwind (Fury node 105977,
// tier 5, 2 ranks since client 1.60.1.70755, where Raging Blows was: "Reduces
// the cooldown of your Whirlwind ability by 2 sec and increases the damage it
// deals to the closest enemy by 30%") pays on every Whirlwind; Improved
// Intercept (node 105926) does nothing for a spec that walks into a raid boss.
func TestFuryGuideBuildTakesImprovedWhirlwindNotImprovedIntercept(t *testing.T) {
	const (
		improvedWhirlwind = 105977
		improvedIntercept = 105926
	)
	_, targets := guideTargets(t, "warrior", "fury", GuideBuildTalents)
	if targets[improvedWhirlwind] != 2 {
		t.Errorf("Improved Whirlwind %d, want 2 (Whirlwind is cast on cooldown)", targets[improvedWhirlwind])
	}
	if targets[improvedIntercept] != 0 {
		t.Errorf("Improved Intercept %d, want 0 (the rotation never Intercepts)", targets[improvedIntercept])
	}
}

// Fury fights in Berserker Stance and casts Berserker Rage below 80 rage
// (data/curated/apl/warrior-fury.json). Improved Berserker Rage (Fury node
// 105978, tier 4, 2 ranks: "will instantly generate 10 Rage") pays on every
// cast. Lingering Rage (node 110857: rage decays later "after leaving
// combat") and Piercing Howl (node 105935: a 10 yd daze) do nothing inside a
// fight, so the build spends no point on them.
func TestFuryGuideBuildTakesImprovedBerserkerRageNotOutOfFightTalents(t *testing.T) {
	const (
		improvedBerserkerRage = 105978
		lingeringRage         = 110857
		piercingHowl          = 105935
	)
	_, targets := guideTargets(t, "warrior", "fury", GuideBuildTalents)
	if targets[improvedBerserkerRage] != 2 {
		t.Errorf("Improved Berserker Rage %d, want 2 (Berserker Rage is cast for rage)", targets[improvedBerserkerRage])
	}
	for name, id := range map[string]int{"Lingering Rage": lingeringRage, "Piercing Howl": piercingHowl} {
		if targets[id] != 0 {
			t.Errorf("%s %d, want 0 (it does nothing inside a fight)", name, targets[id])
		}
	}
}
