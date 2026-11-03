package leveling

import "testing"

// The Fury guide build casts Whirlwind on cooldown and never Intercepts
// (data/curated/apl/warrior-fury.json). Raging Blows (Fury node 105977,
// tier 3: "Reduces the Rage cost of your Cleave and Whirlwind abilities by
// 3") pays on every Whirlwind; Improved Intercept (node 105926, tier 4:
// "Reduces the cooldown of your Intercept ability by 5 sec") does nothing
// for a spec that walks into a raid boss. Moving the point keeps every tier
// gate: 25 Fury points sit below tier 4 instead of 24.
func TestFuryGuideBuildTakesRagingBlowsNotImprovedIntercept(t *testing.T) {
	const (
		ragingBlows       = 105977
		improvedIntercept = 105926
	)
	for _, read := range []struct {
		name string
		fn   func(string, string, string) (string, [3]string, error)
	}{{"build", GuideBuildTalents}, {"raidBuild", GuideRaidBuildTalents}} {
		_, targets := guideTargets(t, "warrior", "fury", read.fn)
		if targets[ragingBlows] != 1 {
			t.Errorf("%s: Raging Blows %d, want 1 (Whirlwind is cast on cooldown)", read.name, targets[ragingBlows])
		}
		if targets[improvedIntercept] != 0 {
			t.Errorf("%s: Improved Intercept %d, want 0 (the rotation never Intercepts)", read.name, targets[improvedIntercept])
		}
	}
}
