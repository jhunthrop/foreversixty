package leveling

import (
	"reflect"
	"testing"
)

// These functions had no standalone unit tests before this move: on
// main, ladder.go's talentNode/talentTree/talentsFile/loadTalentTrees/
// guideBuildTalents/guideTalentTargets/ladderTalentString were only
// exercised indirectly through TestRotationLadder in
// sim/request/ladder_test.go (an integration test against real build
// fixtures under data/builds/<build>), which still runs unchanged
// against this package through ladder.go's new import. These tests are
// new, synthetic-fixture unit tests added by lane bis-all so the moved
// package carries its own direct coverage rather than relying solely
// on that integration path.

func TestGuideTalentTargetsMapsByStableID(t *testing.T) {
	guideTrees := []TalentTree{
		{Position: 0, Talents: []TalentNode{{ID: 101, Tier: 0, Column: 0, MaxRank: 5}, {ID: 102, Tier: 1, Column: 0, MaxRank: 5}}},
		{Position: 1, Talents: []TalentNode{{ID: 201, Tier: 0, Column: 0, MaxRank: 1}}},
	}
	got := GuideTalentTargets(guideTrees, [3]string{"32", "1"})
	want := map[int]int{101: 3, 102: 2, 201: 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GuideTalentTargets = %v, want %v", got, want)
	}
}

func TestGuideTalentTargetsShortDigitsLeaveZero(t *testing.T) {
	guideTrees := []TalentTree{
		{Position: 0, Talents: []TalentNode{{ID: 1, MaxRank: 5}, {ID: 2, MaxRank: 5}}},
	}
	got := GuideTalentTargets(guideTrees, [3]string{"4"})
	want := map[int]int{1: 4, 2: 0}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GuideTalentTargets = %v, want %v", got, want)
	}
}

func TestLadderTalentStringSpendsOwnTreeFirstTopRowDown(t *testing.T) {
	activeTrees := []TalentTree{
		{Talents: []TalentNode{{ID: 1, MaxRank: 5}, {ID: 2, MaxRank: 5}}},
		{Talents: []TalentNode{{ID: 3, MaxRank: 5}}},
	}
	targets := map[int]int{1: 5, 2: 5, 3: 5}

	// level 12 -> budget 3, own tree index 1 spent first: tree 1's
	// single talent gets min(3,5)=3, tree 0's two talents get nothing
	// (budget spent) but still print their zero digits - the string is
	// positional, not trimmed.
	got := LadderTalentString(activeTrees, targets, 1, 12)
	want := "00-3"
	if got != want {
		t.Fatalf("LadderTalentString = %q, want %q", got, want)
	}
}

func TestLadderTalentStringBudgetBelowLevel10IsZero(t *testing.T) {
	activeTrees := []TalentTree{{Talents: []TalentNode{{ID: 1, MaxRank: 5}}}}
	targets := map[int]int{1: 5}
	got := LadderTalentString(activeTrees, targets, 0, 5)
	if got != "0" {
		t.Fatalf("LadderTalentString below level 10 = %q, want %q", got, "0")
	}
}

func TestLadderTalentStringCapsAtMaxRank(t *testing.T) {
	activeTrees := []TalentTree{{Talents: []TalentNode{{ID: 1, MaxRank: 2}}}}
	// Guide target exceeds the active build's max rank for this id (a
	// build where the talent was nerfed after the guide was authored):
	// the active build's MaxRank wins.
	targets := map[int]int{1: 5}
	got := LadderTalentString(activeTrees, targets, 0, 20)
	if got != "2" {
		t.Fatalf("LadderTalentString capped = %q, want %q", got, "2")
	}
}

func TestLoadTalentTreesMissingFile(t *testing.T) {
	if _, err := LoadTalentTrees(t.TempDir(), "no-such-build", "hunter"); err == nil {
		t.Fatal("LoadTalentTrees with a missing build directory: want an error, got nil")
	}
}

func TestGuideBuildTalentsMissingGuide(t *testing.T) {
	if _, _, err := GuideBuildTalents(t.TempDir(), "hunter", "marksmanship"); err == nil {
		t.Fatal("GuideBuildTalents with a missing guide: want an error, got nil")
	}
}

func TestParseBuildCode(t *testing.T) {
	build, trees, err := ParseBuildCode("FS1:1.60.1.70009:shaman:dwarf:553231130010305/01/553302:")
	if err != nil || build != "1.60.1.70009" || trees != [3]string{"553231130010305", "01", "553302"} {
		t.Fatalf("got %q %v %v", build, trees, err)
	}
	if _, _, err := ParseBuildCode("553231130010305/01/553302"); err == nil {
		t.Fatal("a bare digit string must be rejected")
	}
}
