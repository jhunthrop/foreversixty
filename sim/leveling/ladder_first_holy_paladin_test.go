package leveling

import (
	"path/filepath"
	"testing"
)

// The plain top-row-down walk spends Holy points a healer cannot use while live ones wait:
// at 20 the 11th point goes to Purifying Power (Cleanse cost and the Exorcism and Holy
// Wrath cooldowns, none of which the healer casts) rather than Illumination; at 30 the 19th
// goes to Divine Precision (Holy spell hit; heals cannot miss) rather than Infusion of
// Light's second rank (round 15). Round 29 (walks-heal): Reverence before Healing Light,
// Illumination before Divine Favor, Divine Intellect before the dead tolls, Benediction
// and Conviction before Holy Shock, and Precision, Guardian's Favor and Sacred Duty, which
// do nothing for a healer, after everything else.
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
		25: "50320005000100000-00000000000000000-00000000000000000",
		30: "50320015000100000-40000000000000000-00000000000000000",
	}
	for level, w := range want {
		if got := LadderTalentString(trees, targets, holyTree, level, ForSpec("paladin-holy")); got != w {
			t.Errorf("level %d Holy build %q, want %q", level, got, w)
		}
	}
	dead := map[int]string{}
	for _, node := range trees[1].Talents {
		switch node.Name {
		case "Precision", "Guardian's Favor", "Sacred Duty":
			dead[node.ID] = node.Name
		}
	}
	for level := 10; level <= 53; level++ {
		ranks, err := TalentRanksFromString(trees, LadderTalentString(trees, targets, holyTree, level, ForSpec("paladin-holy")))
		if err != nil {
			t.Fatal(err)
		}
		for id, name := range dead {
			if ranks[id] > 0 {
				t.Errorf("level %d holds %s, which does nothing for a healer, before the live points", level, name)
			}
		}
	}
}
