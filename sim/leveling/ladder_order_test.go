package leveling

import (
	"os"
	"path/filepath"
	"testing"
)

// ladderOrderSpecs names the guide and own tree of every spec ladder-order.json may list.
var ladderOrderSpecs = map[string]struct {
	class, slug string
	tree        int
}{
	"hunter-beast-mastery": {"hunter", "beast-mastery", 0},
	"mage-fire":            {"mage", "fire", 1},
}

func TestEmbeddedLadderOrderMatchesCurated(t *testing.T) {
	curated, err := os.ReadFile(filepath.Join("..", "..", "data", "curated", "ladder-order.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(curated) != string(ladderOrderJSON) {
		t.Fatal("sim/leveling/ladder-order.json differs from data/curated/ladder-order.json; run `make ladder-first-sync`")
	}
	if _, err := parseLadderFirst(ladderOrderJSON); err != nil {
		t.Fatal(err)
	}
}

// A hand-set order may only move points between checkpoints: at every level a sim runs,
// the spec's ladder with the row is the ladder without it, so no band, golden or published
// number moves. Between checkpoints it stays legal and never takes back a point.
func TestLadderOrderMovesNoCheckpoint(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	for spec := range embeddedLadderOrder().Specs {
		info, ok := ladderOrderSpecs[spec]
		if !ok {
			t.Errorf("%s: not a spec this test knows; add it to ladderOrderSpecs with its tree", spec)
			continue
		}
		trees, err := LoadTalentTrees(repoRoot, active, info.class)
		if err != nil {
			t.Fatal(err)
		}
		_, digits, err := GuideBuildTalents(repoRoot, info.class, info.slug)
		if err != nil {
			t.Fatal(err)
		}
		targets := GuideTalentTargets(trees, digits)
		withoutOrder := withFirst(embeddedLadderFirst().Specs[spec].First)
		for _, level := range ladderCheckpointLevels {
			got := LadderTalentString(trees, targets, info.tree, level, ForSpec(spec))
			want := LadderTalentString(trees, targets, info.tree, level, withoutOrder)
			if got != want {
				t.Errorf("%s level %d: %q, want the checkpoint build %q; a row that moves a checkpoint belongs in ladder-first.json", spec, level, got, want)
			}
		}
		var prev map[int]int
		for level := 10; level <= 60; level++ {
			s := LadderTalentString(trees, targets, info.tree, level, ForSpec(spec))
			ranks, err := TalentRanksFromString(trees, s)
			if err != nil {
				t.Fatal(err)
			}
			for id, rank := range prev {
				if ranks[id] < rank {
					t.Errorf("%s level %d: %q takes back a point level %d had", spec, level, s, level-1)
				}
			}
			prev = ranks
		}
	}
}

// Owner 2026-11-04: Beast Mastery takes Focused Fire as soon as its row opens at 15, ahead
// of Endurance Training's third and fourth ranks.
func TestBeastMasteryTakesFocusedFireAtFifteen(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "hunter")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "hunter", "beast-mastery")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	var focusedFire, enduranceTraining int
	for _, node := range trees[0].Talents {
		switch node.Name {
		case "Focused Fire":
			focusedFire = node.ID
		case "Endurance Training":
			enduranceTraining = node.ID
		}
	}
	ranks, err := TalentRanksFromString(trees, LadderTalentString(trees, targets, 0, 16, ForSpec("hunter-beast-mastery")))
	if err != nil {
		t.Fatal(err)
	}
	if ranks[focusedFire] != 2 || ranks[enduranceTraining] != 0 {
		t.Errorf("level 16: Focused Fire %d, Endurance Training %d; want 2 and 0", ranks[focusedFire], ranks[enduranceTraining])
	}
}
