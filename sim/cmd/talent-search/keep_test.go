package main

import "testing"

func TestResolveKeepNamesTalentsAndRejectsUnknownOnes(t *testing.T) {
	tr := testTrees()
	got, err := resolveKeep(tr, []string{"B", "E"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || !got[2] || !got[5] {
		t.Fatalf("resolveKeep = %v, want talents 2 and 5", got)
	}
	if _, err := resolveKeep(tr, []string{"Nope"}); err == nil {
		t.Fatal("an unknown talent name: want an error, got nil")
	}
}

func TestSplitKeepTrimsAndSkipsEmpty(t *testing.T) {
	got := splitKeep(" B, E ,,")
	if len(got) != 2 || got[0] != "B" || got[1] != "E" {
		t.Fatalf("splitKeep = %q", got)
	}
}

func TestSwapsNeverTakeAPointFromAKeptTalent(t *testing.T) {
	tr := testTrees()
	guide := build{1: 3, 2: 2, 5: 3}
	for _, c := range swaps(tr, guide, testCredits(), map[int]bool{5: true}, "") {
		if c.Build[5] < guide[5] {
			t.Errorf("%s: took a point from the kept talent E: %v", c.Label, c.Build)
		}
	}
}

func TestGenerateEveryCandidateHoldsTheKeptTalentsIncludingDeepArchetypes(t *testing.T) {
	tr := testTrees()
	guide := build{1: 5, 2: 2, 5: 3}
	modeled := map[int]bool{1: true, 2: true, 3: true, 4: true, 6: true}
	keep := map[int]bool{5: true, 2: true}
	pool := generate(tr, guide, testCredits(), modeled, keep, guide.points(), 2)
	if len(pool) == 0 {
		t.Fatal("no candidates generated")
	}
	deep := false
	for _, c := range pool {
		if !holdsKept(c.Build, guide, keep) {
			t.Errorf("%s: %v drops a kept talent", c.Label, c.Build)
		}
		if err := tr.legal(c.Build, guide.points()); err != nil {
			t.Errorf("%s: %v", c.Label, err)
		}
		deep = deep || c.Label == "deep One"
	}
	if !deep {
		t.Error("the deep One archetype did not survive the keep list")
	}
}
