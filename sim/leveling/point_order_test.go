package leveling

import (
	"strings"
	"testing"
)

// armsTargets is the Arms guide's level-60 build against the active
// build's warrior trees, the way sim/cmd/addon-builds reads it.
func armsTargets(t *testing.T) ([]TalentTree, map[int]int) {
	t.Helper()
	const repoRoot = "../.."
	build, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	guideBuild, digits, err := GuideBuildTalents(repoRoot, "warrior", "arms")
	if err != nil {
		t.Fatal(err)
	}
	if err := RequireGuideBuildMatchesActive(guideBuild, build); err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, build, "warrior")
	if err != nil {
		t.Fatal(err)
	}
	return trees, GuideTalentTargets(trees, digits)
}

// replay spends order's points onto trees and writes them as the
// positional talent string TalentRanksFromString reads.
func replay(t *testing.T, trees []TalentTree, order []Point) string {
	t.Helper()
	digits := make([][]byte, len(trees))
	for i, tree := range trees {
		digits[i] = []byte(strings.Repeat("0", len(tree.Talents)))
	}
	for n, p := range order {
		found := false
		for j, node := range trees[p.Tab-1].Talents {
			if node.Tier+1 == p.Tier && node.Column+1 == p.Column {
				digits[p.Tab-1][j]++
				found = true
			}
		}
		if !found {
			t.Fatalf("point %d %+v names no talent", n+1, p)
		}
	}
	parts := make([]string, len(digits))
	for i, d := range digits {
		parts[i] = string(d)
	}
	return strings.Join(parts, "-")
}

func TestPointOrderReplaysTheArmsGuideBuild(t *testing.T) {
	trees, targets := armsTargets(t)
	const armsTree = 0
	order, err := LadderPointOrder(trees, targets, armsTree)
	if err != nil {
		t.Fatal(err)
	}
	if len(order) != 51 {
		t.Fatalf("got %d points, want 51", len(order))
	}
	if order[0].Tab != 1 || order[0].Tier != 1 {
		t.Fatalf("point 1 = %+v, want tab 1 tier 1", order[0])
	}
	got, err := TalentRanksFromString(trees, replay(t, trees, order))
	if err != nil {
		t.Fatal(err)
	}
	want := make(map[int]int)
	for id, rank := range targets {
		if rank > 0 {
			want[id] = rank
		}
	}
	if len(got) != len(want) {
		t.Fatalf("replayed %d talents, the guide has %d: got %v, want %v", len(got), len(want), got, want)
	}
	for id, rank := range want {
		if got[id] != rank {
			t.Errorf("talent %d: replayed rank %d, guide rank %d", id, got[id], rank)
		}
	}
}

func TestPointOrderRefusesABuildThatLeavesALevelUnspent(t *testing.T) {
	// One talent of five ranks: level 15 has nothing left to spend.
	trees := []TalentTree{{Talents: []TalentNode{{ID: 1, Tier: 0, Column: 0, MaxRank: 5}}}}
	_, err := LadderPointOrder(trees, map[int]int{1: 5}, 0)
	if err == nil || !strings.Contains(err.Error(), "level 15") {
		t.Fatalf("err = %v, want one naming level 15", err)
	}
}
