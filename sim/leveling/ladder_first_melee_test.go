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
		10: "00100000000000000" + rest, // Improved Rend first since round 27
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

// Sim-accuracy round 27: a leveling Subtlety rogue takes Opportunity (+4% a point to
// Backstab, Garrote and Ambush) at 10 to 14, before Master of Deception, which only
// helps a rogue sneak past mobs and does nothing in a fight. Master of Deception follows
// at 15 to 17, so level 18 onward does not change.
func TestSubtletyTakesOpportunityBeforeMasterOfDeception(t *testing.T) {
	const subtletyTree = 2
	const empty = "00000000000000000-0000000000000000000-"
	want := map[int]string{
		10: empty + "001000000000000000",
		12: empty + "003000000000000000",
		14: empty + "005000000000000000",
		15: empty + "015000000000000000",
		17: empty + "035000000000000000",
		18: empty + "035100000000000000",
		20: empty + "035200001000000000",
	}
	for level, w := range want {
		if got := ladderAt(t, "rogue", "subtlety", subtletyTree, level); got != w {
			t.Errorf("level %d Subtlety build %q, want %q", level, got, w)
		}
	}
}

// Sim-accuracy round 27: a leveling Arms warrior takes Improved Rend (+12% a point to
// Rend, which the rotation keeps up from 4) at 10 to 12, then Improved Heroic Strike at
// 13, before Deflection, whose parry does nothing while the target is hitting someone
// else. Deflection's first point at 14 pays the row toll for Improved Overpower at 15,
// and its other two follow at 17 and 18, so level 20 does not change.
func TestArmsTakesImprovedRendBeforeDeflection(t *testing.T) {
	const armsTree = 0
	const rest = "-000000000000000000-000000000000000000"
	want := map[int]string{
		10: "00100000000000000" + rest,
		12: "00300000000000000" + rest,
		13: "10300000000000000" + rest,
		14: "11300000000000000" + rest,
		15: "11300100000000000" + rest,
		16: "11300200000000000" + rest,
		18: "13300200000000000" + rest,
		20: "13320200000000000" + rest,
	}
	for level, w := range want {
		if got := ladderAt(t, "warrior", "arms", armsTree, level); got != w {
			t.Errorf("level %d Arms build %q, want %q", level, got, w)
		}
	}
}
