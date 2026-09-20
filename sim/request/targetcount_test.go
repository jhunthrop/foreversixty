package request

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// TestTargetCountByFightStyle pins the built-target count each fight
// style maps to (contract 10.3): a fixed style sends exactly the
// targets it names, and a timeline style sends ONE target plus the
// timeline itself, relying on the engine to pad the pool to the
// timeline's maximum. sim/adapter's own tests (TestPhantomTargetRows*)
// are the other half of contract 10.3 - that the engine pads it, and to
// what size the summary must then bound its per-target rows to.
func TestTargetCountByFightStyle(t *testing.T) {
	tests := []struct {
		style string
		want  int
	}{
		{"patchwerk", 1},
		{"execute", 1},
		{"light-movement", 1},
		{"heavy-movement", 1},
		{"cleave-2", 2},
		{"cleave-3", 3},
		{"cleave-5", 5},
		// dungeon overrides Targets with a timeline; targetCount must
		// send the timeline's SITE-side sentinel (one target), not the
		// timeline's peak, or the pool would be built twice - once
		// here, once by the engine's own padding - and the two would
		// disagree the day the padding rule changes.
		{"dungeon", 1},
		{"dummy", 1},
	}
	for _, tc := range tests {
		t.Run(tc.style, func(t *testing.T) {
			e, ok := ExpandStyle(tc.style, api.DefaultEncounter())
			if !ok {
				t.Fatalf("style %q not found", tc.style)
			}
			if got := targetCount(e); got != tc.want {
				t.Errorf("targetCount(%q) = %d, want %d", tc.style, got, tc.want)
			}
		})
	}
}

// TestDungeonTimelinePeaksAtFive pins the dungeon style's own timeline,
// which is what the engine pads its target pool to (contract 10.3):
// sim/adapter's per-target row count for a dungeon sim is bounded by
// this number, not by targetCount's 1.
func TestDungeonTimelinePeaksAtFive(t *testing.T) {
	e, ok := ExpandStyle("dungeon", api.DefaultEncounter())
	if !ok {
		t.Fatal(`style "dungeon" not found`)
	}
	if len(e.TargetsOverTime) == 0 {
		t.Fatal("dungeon style carries no targets_over_time timeline")
	}
	var max int
	for _, step := range e.TargetsOverTime {
		if step.Count > max {
			max = step.Count
		}
	}
	if max != 5 {
		t.Errorf("dungeon timeline peaks at %d targets, want 5", max)
	}
}
