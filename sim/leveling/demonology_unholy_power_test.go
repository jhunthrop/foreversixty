package leveling

import (
	"path/filepath"
	"testing"
)

// Sim-accuracy round 14 (warlock lane): every Demonology band summons the Sayaad, and the
// guide never levels with the Imp, so Improved Imp's points did nothing. The leveling build
// puts them in Unholy Power on the same row. Demonology's own tree is the second (index 1);
// Improved Imp is its second digit and Unholy Power its fourth.
func TestDemonologyLevelingBuildTakesUnholyPowerOverImprovedImp(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "warlock")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "warlock", "demonology")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	for _, level := range []int{20, 30, 40, 50, 60} {
		got := LadderTalentString(trees, targets, 1, level, ForSpec("warlock-demonology"))
		demo := got[len("00000000000000000-"):]
		if demo[1] != '0' {
			t.Errorf("level %d: %s takes Improved Imp (%c), want 0", level, got, demo[1])
		}
		if demo[3] != '5' {
			t.Errorf("level %d: %s takes Unholy Power %c, want 5", level, got, demo[3])
		}
	}
}
