package leveling

import "testing"

// A tree shaped like Protection paladin's: a wide middle (tiers 0-3 and a
// tier-4 sideways talent), then a prerequisite step (id 6, tier 4) and a
// capstone (id 7, tier 5) that needs 25 points above it.
func spineTree() ([]TalentTree, map[int]int) {
	trees := []TalentTree{{Talents: []TalentNode{
		{ID: 1, Tier: 0, Column: 0, MaxRank: 5},
		{ID: 2, Tier: 1, Column: 0, MaxRank: 5},
		{ID: 3, Tier: 2, Column: 0, MaxRank: 5},
		{ID: 4, Tier: 3, Column: 0, MaxRank: 5},
		{ID: 5, Tier: 3, Column: 1, MaxRank: 5},
		{ID: 6, Tier: 4, Column: 0, MaxRank: 1},
		{ID: 8, Tier: 4, Column: 1, MaxRank: 5},
		{ID: 7, Tier: 5, Column: 0, MaxRank: 1, PrereqTalentID: 6, PrereqRank: 1},
	}}}
	targets := map[int]int{1: 5, 2: 5, 3: 5, 4: 5, 5: 5, 6: 1, 8: 5, 7: 1}
	return trees, targets
}

func ladderDigits(t *testing.T, trees []TalentTree, targets map[int]int, level int) string {
	t.Helper()
	return LadderTalentString(trees, targets, 0, level)
}

func TestTheCapstoneIsTakenOnceTheBudgetCarriesItWithoutDroppingATalent(t *testing.T) {
	trees, targets := spineTree()
	// Level 38 = 29 points: 25 above the capstone, the prerequisite and the
	// capstone itself is 27. The plain walk spends them on rows 0-3, the
	// prerequisite and three ranks of the sideways talent (id 8); the spine
	// pays for the capstone out of that sideways talent's ranks.
	got := ladderDigits(t, trees, targets, 38)
	if want := "55555121"; got != want {
		t.Fatalf("level 38 = %q, want %q", got, want)
	}
}

func TestTheCapstoneIsLeftOutWhenItWouldCostATalent(t *testing.T) {
	trees, targets := spineTree()
	// Level 36 = 27 points: the capstone fits, but only by dropping the
	// sideways talent the plain walk gave its one rank.
	got := ladderDigits(t, trees, targets, 36)
	if want := "55555110"; got != want {
		t.Fatalf("level 36 = %q, want %q", got, want)
	}
}

func TestTheWalkKeepsTheTierGate(t *testing.T) {
	trees, targets := spineTree()
	for level := 10; level <= 60; level++ {
		ranks, err := TalentRanksFromString(trees, ladderDigits(t, trees, targets, level))
		if err != nil {
			t.Fatal(err)
		}
		for _, node := range trees[0].Talents {
			if ranks[node.ID] == 0 {
				continue
			}
			above := 0
			for _, other := range trees[0].Talents {
				if other.Tier < node.Tier {
					above += ranks[other.ID]
				}
			}
			if above < pointsPerTier*node.Tier {
				t.Fatalf("level %d: talent %d (tier %d) has %d points above it, needs %d", level, node.ID, node.Tier, above, pointsPerTier*node.Tier)
			}
		}
	}
}

func TestACapstoneOutOfReachLeavesThePlainWalk(t *testing.T) {
	trees, targets := spineTree()
	// Level 33 = 24 points: the capstone needs 26.
	if got, want := ladderDigits(t, trees, targets, 33), "55554000"; got != want {
		t.Fatalf("level 33 = %q, want %q", got, want)
	}
}

func TestABandThatAlreadyReachesTheCapstoneIsUnchanged(t *testing.T) {
	trees, targets := spineTree()
	if got, want := ladderDigits(t, trees, targets, 60), "55555151"; got != want {
		t.Fatalf("level 60 = %q, want %q", got, want)
	}
}

func TestNoGuideDepthMeansNoSpine(t *testing.T) {
	trees := []TalentTree{{Talents: []TalentNode{{ID: 1, Tier: 0, MaxRank: 5}, {ID: 2, Tier: 0, Column: 1, MaxRank: 5}}}}
	if got, want := LadderTalentString(trees, map[int]int{1: 5, 2: 5}, 0, 12), "30"; got != want {
		t.Fatalf("tier-0 tree = %q, want %q", got, want)
	}
}

const holyShieldTalentID = 105628

// Protection paladin's ladder: Iron Creed (tier 5) sits above Holy Shield
// (tier 6, needs 30 points above), and the plain walk spent band 40's 31
// points on every row before it.
func TestProtectionPaladinLadderKeepsHolyShieldWhereThePointsAllowIt(t *testing.T) {
	const repoRoot = "../.."
	build, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	guideBuild, digits, err := GuideBuildTalents(repoRoot, "paladin", "protection")
	if err != nil {
		t.Fatal(err)
	}
	guideTrees, err := LoadTalentTrees(repoRoot, guideBuild, "paladin")
	if err != nil {
		t.Fatal(err)
	}
	activeTrees, err := LoadTalentTrees(repoRoot, build, "paladin")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(guideTrees, digits)
	const protectionTree = 1
	wantHolyShield := map[int]bool{38: false, 40: true, 50: true, 60: true}
	for level, want := range wantHolyShield {
		ranks, err := TalentRanksFromString(activeTrees, LadderTalentString(activeTrees, targets, protectionTree, level))
		if err != nil {
			t.Fatal(err)
		}
		if got := ranks[holyShieldTalentID] > 0; got != want {
			t.Errorf("level %d: Holy Shield taken = %v, want %v", level, got, want)
		}
	}
}
