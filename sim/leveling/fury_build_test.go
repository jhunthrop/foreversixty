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
