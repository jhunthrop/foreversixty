package leveling

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

const pointOrderRepoRoot = "../.."

// guideTargets is one of a spec's guide builds against the active
// build's trees, the way sim/cmd/addon-builds reads it.
func guideTargets(t *testing.T, class, specSlug string, read func(string, string, string) (string, [3]string, error)) ([]TalentTree, map[int]int) {
	t.Helper()
	build, err := ReadActiveBuild(pointOrderRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	guideBuild, digits, err := read(pointOrderRepoRoot, class, specSlug)
	if err != nil {
		t.Fatal(err)
	}
	if err := RequireGuideBuildMatchesActive(guideBuild, build); err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(pointOrderRepoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	return trees, GuideTalentTargets(trees, digits)
}

// replay spends order's points onto trees one at a time, failing on any
// point the site planner would refuse (web/src/lib/planner/rules.ts
// validateOrder: tier gate, prerequisite, max rank), and returns the
// positional talent string after each point.
func replay(t *testing.T, trees []TalentTree, order []Point) []string {
	t.Helper()
	digits := make([][]byte, len(trees))
	inTree := make([]int, len(trees))
	for i, tree := range trees {
		digits[i] = []byte(strings.Repeat("0", len(tree.Talents)))
	}
	rankOf := func(tab, id int) int {
		for j, node := range trees[tab].Talents {
			if node.ID == id {
				return int(digits[tab][j] - '0')
			}
		}
		return 0
	}
	var after []string
	for n, p := range order {
		tab := p.Tab - 1
		found := false
		for j, node := range trees[tab].Talents {
			if node.Tier+1 != p.Tier || node.Column+1 != p.Column {
				continue
			}
			found = true
			switch {
			case inTree[tab] < pointsPerTier*node.Tier:
				t.Fatalf("point %d %+v: tier %d with %d points in the tree", n+1, p, node.Tier, inTree[tab])
			case node.PrereqTalentID != 0 && rankOf(tab, node.PrereqTalentID) < node.PrereqRank:
				t.Fatalf("point %d %+v: prerequisite %d below rank %d", n+1, p, node.PrereqTalentID, node.PrereqRank)
			case int(digits[tab][j]-'0') >= node.MaxRank:
				t.Fatalf("point %d %+v: already at max rank %d", n+1, p, node.MaxRank)
			}
			digits[tab][j]++
			inTree[tab]++
		}
		if !found {
			t.Fatalf("point %d %+v names no talent", n+1, p)
		}
		parts := make([]string, len(digits))
		for i, d := range digits {
			parts[i] = string(d)
		}
		after = append(after, strings.Join(parts, "-"))
	}
	return after
}

// requireGuideRanks fails unless s holds exactly the guide's non-zero
// ranks, by talent id.
func requireGuideRanks(t *testing.T, trees []TalentTree, targets map[int]int, s string) {
	t.Helper()
	got, err := TalentRanksFromString(trees, s)
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

func TestPointOrderReplaysTheArmsGuideBuild(t *testing.T) {
	trees, targets := guideTargets(t, "warrior", "arms", GuideBuildTalents)
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
	after := replay(t, trees, order)
	requireGuideRanks(t, trees, targets, after[len(after)-1])
}

// Where the ladder only ever adds a point, the order is its per-level
// diff: after level L's point the build is the ladder's own level-L
// build.
func TestPointOrderIsTheLaddersOwnDiffWhereTheLadderOnlyAdds(t *testing.T) {
	trees, targets := guideTargets(t, "warrior", "arms", GuideBuildTalents)
	order, err := LadderPointOrder(trees, targets, 0)
	if err != nil {
		t.Fatal(err)
	}
	for i, got := range replay(t, trees, order) {
		level := 10 + i
		if want := LadderTalentString(trees, targets, 0, level); got != want {
			t.Fatalf("level %d: order builds %s, ladder %s", level, got, want)
		}
	}
}

// Somewhere among the guide builds the ladder moves a rank between two
// levels (an own-tree spine taking a deep talent's chain; 1.60.1.70658's
// Protection paladin did it at level 40). Wherever it does, the point order
// never takes a point back and still ends on the guide build.
func TestPointOrderNeverRemovesAPointWhereTheLadderMovesOne(t *testing.T) {
	found := 0
	for _, spec := range specs.All {
		trees, targets := guideTargets(t, spec.ClassSlug, spec.SpecSlug, GuideBuildTalents)
		for level := 11; level <= 60; level++ {
			before := LadderTalentString(trees, targets, spec.TreeIndex, level-1)
			at := LadderTalentString(trees, targets, spec.TreeIndex, level)
			if !ladderMovesARank(before, at) {
				continue
			}
			found++
			order, err := LadderPointOrder(trees, targets, spec.TreeIndex)
			if err != nil {
				t.Fatalf("%s: %v", spec.Spec, err)
			}
			if len(order) != 51 {
				t.Fatalf("%s: got %d points, want 51", spec.Spec, len(order))
			}
			after := replay(t, trees, order)
			requireGuideRanks(t, trees, targets, after[len(after)-1])
			break
		}
	}
	if found == 0 {
		t.Skip("no guide build's ladder moves a rank at any level on this build")
	}
}

func ladderMovesARank(before, at string) bool {
	for i := range before {
		if at[i] < before[i] {
			return true
		}
	}
	return false
}

// Every spec's leveling and raid guide builds have a legal 51-point
// order ending on the guide build: what sim/cmd/addon-builds needs to
// write all of them.
func TestPointOrderExistsForEveryGuideBuild(t *testing.T) {
	readers := map[string]func(string, string, string) (string, [3]string, error){
		"leveling": GuideBuildTalents,
		"raid":     GuideRaidBuildTalents,
	}
	for _, spec := range specs.All {
		for kind, read := range readers {
			t.Run(spec.Spec+"/"+kind, func(t *testing.T) {
				trees, targets := guideTargets(t, spec.ClassSlug, spec.SpecSlug, read)
				order, err := LadderPointOrder(trees, targets, spec.TreeIndex)
				if err != nil {
					t.Fatal(err)
				}
				if len(order) != 51 {
					t.Fatalf("got %d points, want 51", len(order))
				}
				after := replay(t, trees, order)
				requireGuideRanks(t, trees, targets, after[len(after)-1])
			})
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
