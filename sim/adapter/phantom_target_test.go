package adapter

import (
	"fmt"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// engineTargets builds the EncounterMetrics a real result carries: one
// UnitMetrics per encounter target, UnitIndex 0..n-1 and Name the
// engine's own Unit.Label ("Target 1", "Target 2", ... - ONE-indexed;
// see target.go's NewTarget and GetMetricsProto in the wowsims-forever
// fork).
func engineTargets(n int) []*proto.UnitMetrics {
	out := make([]*proto.UnitMetrics, n)
	for i := 0; i < n; i++ {
		out[i] = &proto.UnitMetrics{UnitIndex: int32(i), Name: fmt.Sprintf("Target %d", i+1)}
	}
	return out
}

// actionAcrossAllUnits builds one ActionMetrics whose Targets slice
// spans every unit index 0..n (inclusive): 0..n-1 are the encounter's
// own targets, each dealt 1000*(i+1) damage, and index n is the
// player's OWN unit index - the shape sim/core/spell.go:453 actually
// produces, sized by len(env.AllUnits) (targets ++ raid units), always
// zero for a unit that never damages itself. This is the exact fixture
// the dps and tank reviews' phantom rows came from: one row too many,
// always at zero.
func actionAcrossAllUnits(spellID int32, targetCount int) *proto.ActionMetrics {
	targets := make([]*proto.TargetedActionMetrics, 0, targetCount+1)
	for i := 0; i < targetCount; i++ {
		targets = append(targets, &proto.TargetedActionMetrics{
			UnitIndex: int32(i),
			Casts:     10,
			Hits:      10,
			Damage:    float64(1000 * (i + 1)),
		})
	}
	// The phantom entry: the player's own unit index, one past the
	// last real target, always zero.
	targets = append(targets, &proto.TargetedActionMetrics{
		UnitIndex: int32(targetCount),
		Casts:     0,
		Hits:      0,
		Damage:    0,
	})
	return &proto.ActionMetrics{
		Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: spellID}},
		IsMelee: true,
		Targets: targets,
	}
}

// TestPerTargetRowsAreBoundedToEncounterTargets is the row-count guard
// from the brief: a result whose player metrics carry a unit index at or
// above the encounter's own target count must produce NO row for it, for
// both a single-target fight and a five-target one.
func TestPerTargetRowsAreBoundedToEncounterTargets(t *testing.T) {
	tests := []struct {
		name        string
		targetCount int
	}{
		{"single target", 1},
		{"five targets", 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &proto.UnitMetrics{
				Name:      "Sim",
				UnitIndex: int32(tc.targetCount), // one past the last real target
				Dps:       &proto.DistributionMetrics{Avg: 500, Stdev: 10, Max: 600, Min: 400},
				Actions:   []*proto.ActionMetrics{actionAcrossAllUnits(11584, tc.targetCount)},
			}
			res := resultWith(u, 100)
			res.EncounterMetrics = &proto.EncounterMetrics{Targets: engineTargets(tc.targetCount)}

			got, err := Summarize(res, req())
			if err != nil {
				t.Fatal(err)
			}
			a := got.DamageDone[0]
			if len(a.Targets) != tc.targetCount {
				t.Fatalf("row count = %d, want %d (the encounter's own target count): %+v", len(a.Targets), tc.targetCount, a.Targets)
			}
			for i, row := range a.Targets {
				wantName := fmt.Sprintf("Target %d", i+1)
				if row.Name != wantName {
					t.Errorf("Targets[%d].Name = %q, want %q (the engine's own target name, not the adapter's index-based one)", i, row.Name, wantName)
				}
				if row.Total == 0 {
					t.Errorf("Targets[%d] (%s) is zero; every real target in this fixture was dealt nonzero damage - a zero row here means a real target's index was dropped, not the phantom one", i, row.Name)
				}
			}
		})
	}
}

// TestPerTargetTotalsSumToActorTotalAfterFiltering is the conservation
// guard: dropping the phantom (non-target) unit index from the per-
// target rows must not change what the row totals sum to - Task 3's
// invariant, sum(Targets[].Total) == Total, survives the filtering in
// this task.
func TestPerTargetTotalsSumToActorTotalAfterFiltering(t *testing.T) {
	for _, targetCount := range []int{1, 3, 5} {
		t.Run(fmt.Sprintf("%d targets", targetCount), func(t *testing.T) {
			u := &proto.UnitMetrics{
				Name:      "Sim",
				UnitIndex: int32(targetCount),
				Dps:       &proto.DistributionMetrics{Avg: 500, Stdev: 10, Max: 600, Min: 400},
				Actions:   []*proto.ActionMetrics{actionAcrossAllUnits(11584, targetCount)},
			}
			res := resultWith(u, 100)
			res.EncounterMetrics = &proto.EncounterMetrics{Targets: engineTargets(targetCount)}

			got, err := Summarize(res, req())
			if err != nil {
				t.Fatal(err)
			}
			a := got.DamageDone[0]
			var sum int64
			for _, row := range a.Targets {
				sum += row.Total
			}
			if sum != a.Total {
				t.Errorf("sum(Targets[].Total) = %d, want Total = %d", sum, a.Total)
			}
		})
	}
}
