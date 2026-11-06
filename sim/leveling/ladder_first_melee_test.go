package leveling

import (
	"path/filepath"
	"testing"
)

// ladderAt is the ladder build of a class and spec at level, from its guide build.
func ladderAt(t *testing.T, class, spec string, ownTree, level int) string {
	t.Helper()
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, class)
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, class, spec)
	if err != nil {
		t.Fatal(err)
	}
	return LadderTalentString(trees, GuideTalentTargets(trees, digits), ownTree, level, ForSpec(class+"-"+spec))
}

// Sim-accuracy round 26: a leveling Subtlety rogue takes Initiative (a combo point on
// every Garrote, which the rotation casts after each Ghostly Strike) before the dead
// toll points Camouflage and Elusiveness, which only pay for the 15- and 20-point rows:
// two points by 25, the third right after Ghostly Strike. Ghostly Strike still comes at
// 26 and Hemorrhage at 30.
func TestSubtletyTakesInitiativeBeforeTheTollPoints(t *testing.T) {
	const subtletyTree = 2
	const empty = "00000000000000000-0000000000000000000-"
	want := map[int]string{
		20: empty + "035200001000000000",
		25: empty + "035200021300000000",
		26: empty + "035200021301000000",
		30: empty + "035220031301000100",
		40: empty + "035230031321001151",
	}
	for level, w := range want {
		if got := ladderAt(t, "rogue", "subtlety", subtletyTree, level); got != w {
			t.Errorf("level %d Subtlety build %q, want %q", level, got, w)
		}
	}
}

// Sim-accuracy round 26: a leveling Arms warrior takes Improved Overpower before
// Improved Tactical Mastery. Tactical Mastery only keeps rage through a stance change,
// and the Arms rotation changes stance only for Recklessness from 50, so below 50 its
// points do nothing, while Overpower is cast from 12.
func TestArmsTakesImprovedOverpowerBeforeTacticalMastery(t *testing.T) {
	const armsTree = 0
	const rest = "-000000000000000000-000000000000000000"
	want := map[int]string{
		10: "10000000000000000" + rest,
		20: "13320200000000000" + rest,
		25: "13323211000000000" + rest,
		40: "13323213032214001" + rest,
	}
	for level, w := range want {
		if got := ladderAt(t, "warrior", "arms", armsTree, level); got != w {
			t.Errorf("level %d Arms build %q, want %q", level, got, w)
		}
	}
}
