package leveling

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// Sim-accuracy round 16 (warrior lane): no character can take a talent point back without a
// respec, so every spec's ladder only ever adds points from one level to the next. The
// ladder walks each level on its own; this holds every spec's guide build, with its own
// curated reservations (ForSpec), to that at every level from 10 to 60. Protection
// Warrior's moved a point at 41 before the walk kept the level below's ranks, in both
// its build and its raid build (Balance's raid build did the same at 42), so round 17
// holds the raid build's ladder to the same rule.
func TestEveryLadderOnlyAddsPoints(t *testing.T) {
	reads := map[string]func(string, string, string) (string, [3]string, error){
		"build":     GuideBuildTalents,
		"raidBuild": GuideRaidBuildTalents,
	}
	for _, spec := range specs.All {
		for name, read := range reads {
			t.Run(spec.Spec+"/"+name, func(t *testing.T) {
				trees, targets := guideTargets(t, spec.ClassSlug, spec.SpecSlug, read)
				prev := LadderTalentString(trees, targets, spec.TreeIndex, 10, ForSpec(spec.Spec))
				for level := 11; level <= 60; level++ {
					at := LadderTalentString(trees, targets, spec.TreeIndex, level, ForSpec(spec.Spec))
					if dropped := droppedTalents(t, trees, prev, at); len(dropped) > 0 {
						t.Errorf("level %d: %s takes back %v from level %d's %s", level, at, dropped, level-1, prev)
					}
					prev = at
				}
			})
		}
	}
}

// droppedTalents names every talent holding fewer ranks in at than in before.
func droppedTalents(t *testing.T, trees []TalentTree, before, at string) []string {
	t.Helper()
	was, err := TalentRanksFromString(trees, before)
	if err != nil {
		t.Fatal(err)
	}
	now, err := TalentRanksFromString(trees, at)
	if err != nil {
		t.Fatal(err)
	}
	var dropped []string
	for _, tree := range trees {
		for _, node := range tree.Talents {
			if now[node.ID] < was[node.ID] {
				dropped = append(dropped, node.Name)
			}
		}
	}
	return dropped
}
