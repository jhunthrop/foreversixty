package leveling

import (
	"path/filepath"
	"testing"
)

// The bear rotation casts Lacerate (a Feral talent on the 15-point row, rank 1 learned at 25,
// prerequisite Savage Fury 2) and Primal Bite (the 30-point row, rank 1 learned at 40,
// prerequisite Leader of the Pack). A level 30 druid's 21 points hold Lacerate and a level 40
// druid's 31 points hold Primal Bite, but the plain top-row-down walk spends them on
// Shredding Attacks (a Shred discount, and a Lacerate discount the build cannot use) and on
// Blood Frenzy. The curated reservation takes each as soon as it fits (band 20 takes Savage
// Fury's first rank on the way); bands 50 and 60 already carry both and do not change.
func TestBearTakesLacerateAtThirtyAndPrimalBiteAtForty(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "druid")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "druid", "feral-bear")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	const feralTree = 1

	idOf := func(name string) int {
		for _, node := range trees[feralTree].Talents {
			if node.Name == name {
				return node.ID
			}
		}
		t.Fatalf("no %s in the Feral tree", name)
		return -1
	}
	lacerate, primalBite := idOf("Lacerate"), idOf("Primal Bite")

	ranksAt := func(level int) map[int]int {
		s := LadderTalentString(trees, targets, feralTree, level, ForSpec("druid-feral-bear"))
		ranks, err := TalentRanksFromString(trees, s)
		if err != nil {
			t.Fatal(err)
		}
		return ranks
	}
	if got := ranksAt(30); got[lacerate] != 1 {
		t.Errorf("level 30 bear build has Lacerate %d, want 1", got[lacerate])
	}
	if got := ranksAt(40); got[lacerate] != 1 || got[primalBite] != 1 {
		t.Errorf("level 40 bear build has Lacerate %d and Primal Bite %d, want 1 and 1", got[lacerate], got[primalBite])
	}

	// Level 60 is the guide build's. (Band 50 moved in sim-accuracy round 29: Heart of the
	// Wild and Naturalist come before Natural Reaction, TestBearTakesNaturalistBeforeNaturalReaction.)
	got := LadderTalentString(trees, targets, feralTree, 60, ForSpec("druid-feral-bear"))
	if want := LadderTalentString(trees, targets, feralTree, 60); got != want {
		t.Errorf("level 60 bear build %q, want the plain walk's %q", got, want)
	}
}
