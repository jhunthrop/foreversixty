package leveling

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A small own tree: two tier-0 talents, a tier-1 talent and a tier-2 talent.
func firstTestTrees() []TalentTree {
	return []TalentTree{{Name: "Own", Talents: []TalentNode{
		{ID: 1, Name: "Top A", Tier: 0, Column: 0, MaxRank: 5},
		{ID: 2, Name: "Top B", Tier: 0, Column: 1, MaxRank: 5},
		{ID: 3, Name: "Middle", Tier: 1, Column: 0, MaxRank: 3},
		{ID: 4, Name: "Deep", Tier: 2, Column: 0, MaxRank: 1},
	}}}
}

var firstTestTargets = map[int]int{1: 5, 2: 5, 3: 3, 4: 1}

// Without a reservation the walk is unchanged: top row down.
func TestLadderWithoutReservationIsThePlainWalk(t *testing.T) {
	trees := firstTestTrees()
	for _, level := range []int{19, 20, 30} {
		plain := LadderTalentString(trees, firstTestTargets, 0, level)
		none := LadderTalentString(trees, firstTestTargets, 0, level, withFirst(nil))
		if plain != none {
			t.Errorf("level %d: %q with an empty reservation, %q without", level, none, plain)
		}
	}
	if got := LadderTalentString(trees, firstTestTargets, 0, 20); got != "5510" {
		t.Errorf("plain walk at 20 = %q, want 5510", got)
	}
}

// A reservation is spent first, each entry once the build can legally hold it: at 11
// points Deep (tier 2) needs ten points above it, so Middle 1 and Deep 1 come first and
// the walk fills the rest top row down.
func TestLadderSpendsTheReservationFirst(t *testing.T) {
	trees := firstTestTrees()
	first := []LadderFirstRank{{Talent: "Middle", Rank: 1}, {Talent: "Deep", Rank: 1}}
	cases := map[int]string{
		20: "5411", // 11 points: Top A 5, Top B 4, Middle 1, Deep 1
		19: "5410", // 10 points: Deep cannot be held yet (9 above it), so only Middle is reserved
		14: "5000", // 5 points: Middle (tier 1) cannot be held, so nothing is reserved
		15: "5010", // 6 points: Middle once five sit above it
		60: "5531", // every target met: the reservation changes nothing at the end
	}
	for level, want := range cases {
		if got := LadderTalentString(trees, firstTestTargets, 0, level, withFirst(first)); got != want {
			t.Errorf("level %d: %q, want %q", level, got, want)
		}
	}
}

// A reservation never takes a talent past the guide's target for it.
func TestLadderReservationStaysInsideTheGuideTargets(t *testing.T) {
	trees := firstTestTrees()
	targets := map[int]int{1: 5, 2: 5, 3: 3}
	first := []LadderFirstRank{{Talent: "Deep", Rank: 1}}
	if got := LadderTalentString(trees, targets, 0, 30, withFirst(first)); strings.HasSuffix(got, "1") {
		t.Errorf("level 30 = %q: Deep is reserved though the guide never takes it", got)
	}
}

// The curated table: its copy is current, every spec names a written spec, every talent
// exists in that spec's own tree with a guide target at or above the reserved rank, and
// the ladder it makes is legal at every level, never removes a point from one level to the
// next and ends on the guide's own level-60 build.
func TestEmbeddedLadderFirstMatchesCurated(t *testing.T) {
	curated, err := os.ReadFile(filepath.Join("..", "..", "data", "curated", "ladder-first.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(curated) != string(ladderFirstJSON) {
		t.Fatal("sim/leveling/ladder-first.json differs from data/curated/ladder-first.json; run `make ladder-first-sync`")
	}
	if _, err := parseLadderFirst(ladderFirstJSON); err != nil {
		t.Fatal(err)
	}
}
