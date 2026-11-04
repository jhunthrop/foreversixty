package leveling

import "testing"

// The Arms raid build fights a boss that never swings at it, so a point in
// Deflection (Arms node 105957, parry) or Blood Craze (Fury node 105934, a
// heal after a crit taken) does nothing. Those points reach Spearing Strike
// (Arms node 105949: 75% weapon damage, 150% against Giants and Dragonkin)
// and, through 15 Fury points, Reckless Swings (node 113636: the next 4
// swings generate 5 more rage) and Improved Execute (node 105932: Execute
// costs 5 less rage), all three of which the rotation casts
// (data/curated/apl/warrior-arms.json).
func TestArmsRaidBuildSpendsNoPointOnAnUnhitDefense(t *testing.T) {
	const (
		deflection      = 105957
		bloodCraze      = 105934
		spearingStrike  = 105949
		recklessSwings  = 113636
		improvedExecute = 105932
		mortalStrike    = 105941
	)
	_, targets := guideTargets(t, "warrior", "arms", GuideRaidBuildTalents)
	for name, id := range map[string]int{"Deflection": deflection, "Blood Craze": bloodCraze} {
		if targets[id] != 0 {
			t.Errorf("%s %d, want 0 (the boss never swings at an Arms warrior)", name, targets[id])
		}
	}
	for name, want := range map[string][2]int{
		"Spearing Strike":  {spearingStrike, 1},
		"Reckless Swings":  {recklessSwings, 1},
		"Improved Execute": {improvedExecute, 2},
		"Mortal Strike":    {mortalStrike, 1},
	} {
		if targets[want[0]] != want[1] {
			t.Errorf("%s %d, want %d (the rotation casts it)", name, targets[want[0]], want[1])
		}
	}
}
