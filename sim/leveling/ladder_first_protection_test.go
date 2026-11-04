package leveling

import (
	"path/filepath"
	"testing"
)

// Holy Shield sits on the Protection tree's 21-point row (tier 4) and the rotation casts it
// from rank 1 (20925, level 30). A level 30 paladin has 21 points, enough for the 20 points
// above it plus Holy Shield, but the plain top-row-down walk spends the 21st on Improved
// Righteous Fury. The curated reservation takes Holy Shield (and Shieldsworn, so band 40
// keeps the capstone the spine gave it) as soon as each fits. Band 20's Templar's Bulwark
// (owner ruling, round 14) is TestProtectionTakesTemplarsBulwarkAtTwenty's; the bands above
// keep the plain walk.
func TestProtectionTakesHolyShieldAtThirty(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "paladin")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "paladin", "protection")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	const protectionTree = 1

	holyShield := -1
	for _, node := range trees[protectionTree].Talents {
		if node.Name == "Holy Shield" {
			holyShield = node.ID
		}
	}
	if holyShield < 0 {
		t.Fatal("no Holy Shield in the Protection tree")
	}

	at30 := LadderTalentString(trees, targets, protectionTree, 30, ForSpec("paladin-protection"))
	ranks, err := TalentRanksFromString(trees, at30)
	if err != nil {
		t.Fatal(err)
	}
	if ranks[holyShield] != 1 {
		t.Errorf("level 30 Protection build %q has no Holy Shield", at30)
	}

	for _, level := range []int{40, 50, 60} {
		got := LadderTalentString(trees, targets, protectionTree, level, ForSpec("paladin-protection"))
		want := LadderTalentString(trees, targets, protectionTree, level)
		if got != want {
			t.Errorf("level %d Protection build %q, want the plain walk's %q", level, got, want)
		}
	}
}
