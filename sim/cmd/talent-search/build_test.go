package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// testTrees is a small two-tree class:
//
//	tree 0: A (tier 0, 5 ranks), B (tier 0, 5), C (tier 1, 3, needs 5 above),
//	        D (tier 2, 1, needs 10 above and 3 points in C)
//	tree 1: E (tier 0, 5), F (tier 1, 2)
func testTrees() talentTrees {
	return newTalentTrees([]leveling.TalentTree{
		{Name: "One", Position: 0, Talents: []leveling.TalentNode{
			{ID: 1, Name: "A", Tier: 0, Column: 0, MaxRank: 5},
			{ID: 2, Name: "B", Tier: 0, Column: 1, MaxRank: 5},
			{ID: 3, Name: "C", Tier: 1, Column: 0, MaxRank: 3},
			{ID: 4, Name: "D", Tier: 2, Column: 0, MaxRank: 1, PrereqTalentID: 3, PrereqRank: 3},
		}},
		{Name: "Two", Position: 1, Talents: []leveling.TalentNode{
			{ID: 5, Name: "E", Tier: 0, Column: 0, MaxRank: 5},
			{ID: 6, Name: "F", Tier: 1, Column: 0, MaxRank: 2},
		}},
	})
}

func TestLegal(t *testing.T) {
	tr := testTrees()
	for _, c := range []struct {
		name  string
		b     build
		total int
		want  string // "" = legal, else a substring of the error
	}{
		{"tier 0 only", build{1: 5, 5: 2}, 7, ""},
		{"tier 1 after 5 above", build{1: 5, 3: 1}, 6, ""},
		{"tier 1 too early", build{1: 4, 3: 1}, 5, "needs 5 points above"},
		{"points in the other tree do not unlock this one", build{1: 4, 5: 5, 3: 1}, 10, "needs 5 points above"},
		{"tier 2 with its prereq", build{1: 5, 2: 2, 3: 3, 4: 1}, 11, ""},
		{"tier 2 missing its prereq", build{1: 5, 2: 5, 3: 2, 4: 1}, 13, "needs 3 points in C"},
		{"tier 2 gate counts only lower tiers", build{1: 5, 3: 3, 4: 1}, 9, "needs 10 points above"},
		{"over max rank", build{1: 6}, 6, "6 of 5"},
		{"wrong total", build{1: 5}, 6, "spends 5 points, want 6"},
		{"unknown talent", build{99: 1}, 1, "not in this class"},
	} {
		t.Run(c.name, func(t *testing.T) {
			err := tr.legal(c.b, c.total)
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("legal = %v, want nil", err)
			case c.want != "" && (err == nil || !strings.Contains(err.Error(), c.want)):
				t.Fatalf("legal = %v, want an error containing %q", err, c.want)
			}
		})
	}
}

func TestWithNeverChangesTheOriginal(t *testing.T) {
	b := build{1: 2}
	nb := b.with(1, 1).with(2, 1)
	if b[1] != 2 || len(b) != 1 {
		t.Fatalf("with changed the original: %v", b)
	}
	if nb[1] != 3 || nb[2] != 1 {
		t.Fatalf("with = %v", nb)
	}
	if _, ok := nb.with(2, -1)[2]; ok {
		t.Fatal("a talent taken back to zero stayed in the build")
	}
}

func TestDecodeActiveAndFS1RoundTrip(t *testing.T) {
	tr := testTrees()
	b, err := tr.decodeActive("5030-2")
	if err != nil {
		t.Fatal(err)
	}
	if b[1] != 5 || b[3] != 3 || b[5] != 2 || b.points() != 10 {
		t.Fatalf("decodeActive = %v", b)
	}
	if got := tr.fs1("9.9", "x", "human", b); got != "FS1:9.9:x:human:503/2:" {
		t.Fatalf("fs1 = %q", got)
	}
	if got := tr.fs1("9.9", "x", "human", build{5: 1}); got != "FS1:9.9:x:human:0/1:" {
		t.Fatalf("fs1 of an empty first tree = %q", got)
	}
	for _, bad := range []string{"5", "50300-0", "5x-0"} {
		if _, err := tr.decodeActive(bad); err == nil {
			t.Errorf("decodeActive(%q) returned no error", bad)
		}
	}
}

func TestDistanceCountsMovedPoints(t *testing.T) {
	if d := distance(build{1: 5, 2: 1}, build{1: 4, 2: 1, 5: 1}); d != 1 {
		t.Fatalf("one point moved: distance = %d, want 1", d)
	}
	if d := distance(build{1: 5}, build{5: 5}); d != 5 {
		t.Fatalf("five points moved: distance = %d, want 5", d)
	}
}

func TestDiffNames(t *testing.T) {
	added, removed := testTrees().diffNames(build{1: 5, 5: 2}, build{1: 3, 2: 2, 5: 2})
	if strings.Join(added, ",") != "+2 B" || strings.Join(removed, ",") != "-2 A" {
		t.Fatalf("diffNames = %v / %v", added, removed)
	}
}
