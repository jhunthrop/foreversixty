package leveling

import (
	"path/filepath"
	"strings"
	"testing"
)

// specRow is what the soundness test needs of sim/specs (which this package cannot import
// in its own tests without a cycle in the future): the class, the guide file and the tree.
var ladderFirstSpecs = map[string]struct {
	class, slug string
	tree        int
}{
	"druid-restoration":  {"druid", "restoration", 2},
	"paladin-holy":       {"paladin", "holy", 0},
	"paladin-protection": {"paladin", "protection", 1},
	"priest-discipline":  {"priest", "discipline", 0},
	"priest-holy":        {"priest", "holy", 1},
	"shaman-restoration": {"shaman", "restoration", 2},
}

// Every curated reservation is sound on the active build: its spec is one this test knows,
// every talent sits in the spec's own tree with a guide rank at or above the reserved one,
// every level's build is legal, no level removes a point the level before it had, and
// level 60 is the guide build's own walk.
func TestCuratedLadderFirstIsSound(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	for spec, row := range embeddedLadderFirst().Specs {
		info, ok := ladderFirstSpecs[spec]
		if !ok {
			t.Errorf("%s: not a healer this lane's mechanism covers; add it to ladderFirstSpecs with its tree", spec)
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
		own := trees[info.tree]
		byName := map[string]TalentNode{}
		for _, node := range own.Talents {
			byName[node.Name] = node
		}
		for _, r := range row.First {
			node, ok := byName[r.Talent]
			if !ok {
				t.Errorf("%s: %q is not a talent of its own tree (%s)", spec, r.Talent, own.Name)
				continue
			}
			if targets[node.ID] < r.Rank {
				t.Errorf("%s: %q reserved at rank %d, past the guide's %d", spec, r.Talent, r.Rank, targets[node.ID])
			}
		}
		var prev map[int]int
		for level := 10; level <= 60; level++ {
			s := LadderTalentString(trees, targets, info.tree, level, ForSpec(spec))
			ranks, err := TalentRanksFromString(trees, s)
			if err != nil {
				t.Fatal(err)
			}
			for ti, part := range strings.Split(s, "-") {
				d := make([]int, len(part))
				for j, c := range part {
					d[j] = int(c - '0')
				}
				if !treeIsLegal(trees[ti], d) {
					t.Errorf("%s level %d: %q breaks a tier gate or a prerequisite", spec, level, s)
				}
			}
			for id, rank := range prev {
				if ranks[id] < rank {
					t.Errorf("%s level %d: %q takes back a point level %d had", spec, level, s, level-1)
				}
			}
			prev = ranks
		}
		if got, want := LadderTalentString(trees, targets, info.tree, 60, ForSpec(spec)), LadderTalentString(trees, targets, info.tree, 60); got != want {
			t.Errorf("%s at 60: %q, want the guide build's %q", spec, got, want)
		}
	}
}
