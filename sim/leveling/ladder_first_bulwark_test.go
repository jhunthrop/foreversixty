package leveling

import (
	"path/filepath"
	"testing"
)

// Owner ruling (sim-accuracy round 14): Protection at band 20 takes Templar's Bulwark
// (the 10-point row; no level gate since 70707 hotfix 113139) over Sacred Duty's second
// point. Sacred Duty only shortens Divine Shield, Divine Protection and Bulwark, none of
// which the band-20 rotation leans on, while Bulwark is cast; chance of death on the tank
// fight falls from about 77% to 66% for about 2% less threat. This is a stated exception
// to the threat metric for this band only: every other band keeps its build.
func TestProtectionTakesTemplarsBulwarkAtTwenty(t *testing.T) {
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

	want := map[int]string{
		19: "00000000000000000-15301000000000000-00000000000000000",
		20: "00000000000000000-15301100000000000-00000000000000000",
		25: "00000000000000000-15302102110000000-00000000000000000", // Swift Judgement first since round 28
		30: "00000000000000000-15302102510001000-00000000000000000",
		35: "00000000000000000-15302102513201000-00000000000000000",
		40: "00000000000000000-15302102513501011-00000000000000000",
		50: "05003000000000000-15302102513501031-00000000000000000",
		60: "05005000000000000-15302102513501031-50030000000000000",
	}
	for level, w := range want {
		got := LadderTalentString(trees, targets, protectionTree, level, ForSpec("paladin-protection"))
		if got != w {
			t.Errorf("level %d Protection build %q, want %q", level, got, w)
		}
	}
}
