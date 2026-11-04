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
	"druid-feral-bear":     {"druid", "feral-bear", 1},
	"druid-restoration":    {"druid", "restoration", 2},
	"hunter-beast-mastery": {"hunter", "beast-mastery", 0},
	"hunter-survival":      {"hunter", "survival", 2},
	"mage-fire":            {"mage", "fire", 1},
	"paladin-holy":         {"paladin", "holy", 0},
	"paladin-protection":   {"paladin", "protection", 1},
	"priest-discipline":    {"priest", "discipline", 0},
	"priest-holy":          {"priest", "holy", 1},
	"rogue-assassination":  {"rogue", "assassination", 0},
	"rogue-subtlety":       {"rogue", "subtlety", 2},
	"shaman-restoration":   {"shaman", "restoration", 2},
	"warlock-affliction":   {"warlock", "affliction", 0},
	"warlock-demonology":   {"warlock", "demonology", 1},
	"warrior-arms":         {"warrior", "arms", 0},
	"warrior-fury":         {"warrior", "fury", 1},
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
			t.Errorf("%s: not a spec this test knows; add it to ladderFirstSpecs with its tree", spec)
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
		chain := pathToDeepestTarget(own, targets)
		capstone := chain[len(chain)-1]
		var prev map[int]int
		for level := 10; level <= 60; level++ {
			s := LadderTalentString(trees, targets, info.tree, level, ForSpec(spec))
			ranks, err := TalentRanksFromString(trees, s)
			if err != nil {
				t.Fatal(err)
			}
			// A reservation replaces the walk's spine, so at a published band it must not lose
			// the capstone the spine would have given. (Between bands, Discipline at 35 and
			// Restoration shaman at 46 to 49 still do: sim-accuracy round 13, walk lane.)
			plain, err := TalentRanksFromString(trees, LadderTalentString(trees, targets, info.tree, level))
			if err != nil {
				t.Fatal(err)
			}
			if level%10 == 0 && plain[capstone.ID] > 0 && ranks[capstone.ID] == 0 {
				t.Errorf("%s level %d: %q drops %s, which the plain walk's spine holds", spec, level, s, capstone.Name)
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

// Sim-accuracy round 12 (healprofile): Restoration shaman holds Water Shield, which the
// rotation casts, at band 20, ahead of Tidal Focus's fourth rank
// (design/reviews/2026-11-03-sim-audit-round12/healprofile.md). Round 13 (walk) holds Mana
// Tide Totem next, which moves band 30's Tidal Mastery point onto it.
func TestRestorationShamanHoldsWaterShieldAtTwenty(t *testing.T) {
	repoRoot := filepath.Join("..", "..")
	active, err := ReadActiveBuild(repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	trees, err := LoadTalentTrees(repoRoot, active, "shaman")
	if err != nil {
		t.Fatal(err)
	}
	_, digits, err := GuideBuildTalents(repoRoot, "shaman", "restoration")
	if err != nil {
		t.Fatal(err)
	}
	targets := GuideTalentTargets(trees, digits)
	for level, want := range map[int]string{
		20: "0000000000000000-000000000000000000-5320000010000000",
		30: "0000000000000000-000000000000000000-5530003310010000",
		40: "0000000000000000-000000000000000000-5530003315510000",
	} {
		if got := LadderTalentString(trees, targets, 2, level, ForSpec("shaman-restoration")); got != want {
			t.Errorf("level %d: %s, want %s", level, got, want)
		}
	}
}

// Sim-accuracy round 15 (hunter, owner 2026-11-04): Beast Mastery levels with Focused Fire
// (2% more damage for the hunter and the pet) as soon as its row opens at 15, ahead of the
// last Endurance Training points (pet health and armor). The ladder bands hold the same
// points either way, so this is an order-only row: real leveling play decides a tie.
func TestBeastMasteryTakesFocusedFireBeforeEnduranceTraining(t *testing.T) {
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
	byName := map[string]int{}
	for _, node := range trees[0].Talents {
		byName[node.Name] = node.ID
	}
	for level, want := range map[int][2]int{15: {1, 0}, 16: {2, 0}, 17: {2, 1}, 20: {2, 4}} {
		ranks, err := TalentRanksFromString(trees, LadderTalentString(trees, targets, 0, level, ForSpec("hunter-beast-mastery")))
		if err != nil {
			t.Fatal(err)
		}
		if got := [2]int{ranks[byName["Focused Fire"]], ranks[byName["Endurance Training"]]}; got != want {
			t.Errorf("level %d: Focused Fire, Endurance Training = %v, want %v", level, got, want)
		}
	}
}
