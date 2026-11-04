package leveling

import (
	"path/filepath"
	"testing"
)

// The plain top-row-down walk spends Holy points a healer cannot use while live ones wait:
// at 20 the 11th point goes to Purifying Power (Cleanse cost and the Exorcism and Holy
// Wrath cooldowns, none of which the healer casts) rather than Illumination; at 30 the 19th
// goes to Divine Precision (Holy spell hit; heals cannot miss) rather than Infusion of
// Light's second rank. The curated reservation takes Illumination, Oath of Light, Divine
// Favor and Infusion of Light ahead of them, and leaves the bands from 40 on alone.
func TestHolyPaladinBandsSpendNoDeadPoints(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "paladin")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "paladin", "holy")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	const holyTree = 0

	want := map[int]string{
		20: "50320001000000000-00000000000000000-00000000000000000",
		25: "50320014000100000-00000000000000000-00000000000000000",
		30: "50320215002100000-00000000000000000-00000000000000000",
	}
	for level, w := range want {
		if got := LadderTalentString(trees, targets, holyTree, level, ForSpec("paladin-holy")); got != w {
			t.Errorf("level %d Holy build %q, want %q", level, got, w)
		}
	}
	for _, level := range []int{35, 40, 50, 60} {
		got := LadderTalentString(trees, targets, holyTree, level, ForSpec("paladin-holy"))
		plain := LadderTalentString(trees, targets, holyTree, level)
		if got != plain {
			t.Errorf("level %d Holy build %q, want the plain walk's %q", level, got, plain)
		}
	}
}
