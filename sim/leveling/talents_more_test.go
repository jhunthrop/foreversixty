package leveling

import "testing"

// repoRootFixture is testdata/reporoot: a minimal stand-in for the
// site repository, carrying just one build's talents/hunter.json and
// one guide's frontmatter (recommendedRaces + the FS1 build line) -
// enough to exercise LoadTalentTrees/GuideBuildTalents's real success
// path, not just their already-tested missing-file error path.
const repoRootFixture = "testdata/reporoot"

func TestLoadTalentTreesSortsByPositionThenTierThenColumn(t *testing.T) {
	trees, err := LoadTalentTrees(repoRootFixture, "testbuild", "hunter")
	if err != nil {
		t.Fatalf("LoadTalentTrees: %v", err)
	}
	if len(trees) != 2 {
		t.Fatalf("LoadTalentTrees = %d trees, want 2", len(trees))
	}
	// Fixture's trees are given position 1 then position 0 - the loader
	// must reorder them so trees[0].Position == 0.
	if trees[0].Position != 0 || trees[1].Position != 1 {
		t.Fatalf("trees not sorted by Position: %+v", trees)
	}
	// The position-1 tree's talents are given out of (tier, column)
	// order (id 20 tier1/col0, id 10 tier0/col1, id 11 tier0/col0) - the
	// loader must sort them to id 11 (t0c0), id 10 (t0c1), id 20 (t1c0).
	tree1 := trees[1]
	if len(tree1.Talents) != 3 {
		t.Fatalf("tree1.Talents = %+v, want 3", tree1.Talents)
	}
	gotOrder := []int{tree1.Talents[0].ID, tree1.Talents[1].ID, tree1.Talents[2].ID}
	wantOrder := []int{11, 10, 20}
	for i := range wantOrder {
		if gotOrder[i] != wantOrder[i] {
			t.Fatalf("tree1 talent order = %v, want %v", gotOrder, wantOrder)
		}
	}
}

func TestLoadTalentTreesMalformedJSON(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, dir+"/data/builds/testbuild/talents/hunter.json", "{not json"); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadTalentTrees(dir, "testbuild", "hunter"); err == nil {
		t.Fatal("LoadTalentTrees with malformed JSON: want an error, got nil")
	}
}

func TestGuideBuildTalentsReadsFS1Line(t *testing.T) {
	build, trees, err := GuideBuildTalents(repoRootFixture, "hunter", "marksmanship")
	if err != nil {
		t.Fatalf("GuideBuildTalents: %v", err)
	}
	if build != "testbuild" {
		t.Errorf("build = %q, want testbuild", build)
	}
	want := [3]string{"2050501", "50", "005"}
	if trees != want {
		t.Errorf("trees = %v, want %v", trees, want)
	}
}

func TestGuideBuildTalentsNoFS1Line(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, dir+"/web/src/content/guides/hunter/no-build.md", "---\ntitle: no build line\n---\n"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := GuideBuildTalents(dir, "hunter", "no-build"); err == nil {
		t.Fatal("GuideBuildTalents with no FS1 line: want an error, got nil")
	}
}

func TestGuideTalentTargetsMoreTreesThanDigitsBreaksEarly(t *testing.T) {
	// GuideTalentTargets must stop at len(treeDigits) rather than index
	// out of range when guideTrees carries more entries than the FS1
	// code's three tree-digit strings (should not happen in real data,
	// but the function's own bounds check is a branch this pins).
	guideTrees := []TalentTree{
		{Talents: []TalentNode{{ID: 1, MaxRank: 5}}},
		{Talents: []TalentNode{{ID: 2, MaxRank: 5}}},
		{Talents: []TalentNode{{ID: 3, MaxRank: 5}}},
		{Talents: []TalentNode{{ID: 4, MaxRank: 5}}}, // a 4th tree - out of range for [3]string
	}
	got := GuideTalentTargets(guideTrees, [3]string{"5", "5", "5"})
	if _, ok := got[4]; ok {
		t.Fatalf("GuideTalentTargets = %v, want id 4 absent (its tree has no digit string)", got)
	}
	if len(got) != 3 {
		t.Fatalf("GuideTalentTargets = %v, want exactly 3 entries", got)
	}
}

func TestLadderTalentStringOwnTreeIndexOutOfRangeIsSkipped(t *testing.T) {
	// An ownTreeIndex that does not name a real tree (defensive: should
	// not happen with real curated data) is simply skipped rather than
	// panicking - the other trees still spend their budget in order.
	activeTrees := []TalentTree{
		{Talents: []TalentNode{{ID: 1, MaxRank: 5}}},
	}
	targets := map[int]int{1: 5}
	got := LadderTalentString(activeTrees, targets, 99, 20)
	if got != "5" {
		t.Fatalf("LadderTalentString with an out-of-range ownTreeIndex = %q, want %q (tree 0 still spent)", got, "5")
	}
}

func TestLadderTalentStringMultipleTreesJoinedWithDash(t *testing.T) {
	activeTrees := []TalentTree{
		{Talents: []TalentNode{{ID: 1, MaxRank: 5}}},
		{Talents: []TalentNode{{ID: 2, MaxRank: 5}}},
		{Talents: []TalentNode{{ID: 3, MaxRank: 5}}},
	}
	targets := map[int]int{1: 5, 2: 5, 3: 5}
	got := LadderTalentString(activeTrees, targets, 0, 60)
	// budget = 51: own tree (0) gets min(5,51)=5, remaining 46; tree 1
	// gets 5, remaining 41; tree 2 gets 5.
	if got != "5-5-5" {
		t.Fatalf("LadderTalentString = %q, want 5-5-5", got)
	}
}
