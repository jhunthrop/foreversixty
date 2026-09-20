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
// produces, sized by len(env.AllUnits) (targets ++ raid units). In every
// case actually measured (baseline.md's D4 table, and the AFTER numbers
// in task-4-report.md) that phantom entry's damage is exactly zero;
// phantomDamage lets a test override that to exercise the reconciliation
// path for the case that has never been observed (see
// TestPerTargetReconcilesNonzeroDamageOnANonTargetIndex).
func actionAcrossAllUnits(spellID int32, targetCount int, phantomDamage float64) *proto.ActionMetrics {
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
	// last real target.
	targets = append(targets, &proto.TargetedActionMetrics{
		UnitIndex: int32(targetCount),
		Casts:     0,
		Hits:      0,
		Damage:    phantomDamage,
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
				Actions:   []*proto.ActionMetrics{actionAcrossAllUnits(11584, tc.targetCount, 0)},
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
				Actions:   []*proto.ActionMetrics{actionAcrossAllUnits(11584, targetCount, 0)},
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

// TestPerTargetFallsBackToOneUnknownRowWhenEncounterMetricsIsMissing
// pins the fallback's own shape: a result carrying no EncounterMetrics
// (an old or malformed result - a real completed sim always sets it,
// sim/core/sim.go:391 calls sim.Encounter.GetMetricsProto()
// unconditionally) must NOT silently reproduce the phantom-row bug this
// task removes. resultWith and oneAction() build exactly such a result:
// one action dealing 1000 damage to UnitIndex 1 with no EncounterMetrics
// at all, so nothing here can tell that index apart from a raid
// member's. Rather than falling back to a row per recorded index - the
// pre-fix shape - the actor's whole Total collapses into one row that
// says the breakdown is unknown, which is unambiguously distinguishable
// from a real target row.
func TestPerTargetFallsBackToOneUnknownRowWhenEncounterMetricsIsMissing(t *testing.T) {
	u := oneAction()
	res := resultWith(u, 100)
	if res.EncounterMetrics != nil {
		t.Fatal("test fixture assumption broken: resultWith now sets EncounterMetrics")
	}

	got, err := Summarize(res, req())
	if err != nil {
		t.Fatal(err)
	}
	a := got.DamageDone[0]
	if len(a.Targets) != 1 {
		t.Fatalf("Targets = %+v, want exactly one fallback row, not one per recorded index", a.Targets)
	}
	row := a.Targets[0]
	if row.GUID != "sim-target-unknown" || row.Name != "Unknown Target" {
		t.Errorf(`fallback row = %+v, want GUID "sim-target-unknown", Name "Unknown Target"`, row)
	}
	if row.Total != a.Total {
		t.Errorf("fallback row Total = %d, want the actor's whole Total %d", row.Total, a.Total)
	}
}

// TestPerTargetNoFallbackRowWhenActorDealtNoDamage is the fallback's
// other edge, alongside the row-per-index case above: an actor with
// nothing to attribute (Total == 0) gets no row at all, fallback or
// otherwise. Task 3's invariant, sum(Targets) == Total, holds trivially
// at 0 == 0 - an "Unknown Target" row carrying zero would be exactly the
// kind of always-zero phantom row this task removes.
func TestPerTargetNoFallbackRowWhenActorDealtNoDamage(t *testing.T) {
	u := &proto.UnitMetrics{Name: "Idle", UnitIndex: 0, Dps: &proto.DistributionMetrics{}}
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	a := got.DamageDone[0]
	if len(a.Targets) != 0 {
		t.Errorf("Targets = %+v, want none for an actor with zero Total", a.Targets)
	}
}

// TestPerTargetReconcilesNonzeroDamageOnANonTargetIndex pins the
// proportional-redistribution case apportion's own comment describes
// but no other fixture exercises: three real targets plus one
// non-target index that - contrary to every case actually measured -
// carries NONZERO recorded damage. The row list still bounds to the
// three real targets (the non-target index gets no row of its own, and
// Total is nonzero so the single-row fallback above does not apply
// either - idx is non-empty here), and the actor's whole Total,
// phantom damage included, is apportioned across the real rows by
// their own proportions rather than lost.
func TestPerTargetReconcilesNonzeroDamageOnANonTargetIndex(t *testing.T) {
	const targetCount = 3
	u := &proto.UnitMetrics{
		Name:      "Sim",
		UnitIndex: int32(targetCount),
		Dps:       &proto.DistributionMetrics{Avg: 500, Stdev: 10, Max: 600, Min: 400},
		Actions:   []*proto.ActionMetrics{actionAcrossAllUnits(11584, targetCount, 500)},
	}
	res := resultWith(u, 100)
	res.EncounterMetrics = &proto.EncounterMetrics{Targets: engineTargets(targetCount)}

	got, err := Summarize(res, req())
	if err != nil {
		t.Fatal(err)
	}
	a := got.DamageDone[0]
	if len(a.Targets) != targetCount {
		t.Fatalf("row count = %d, want %d; the non-target index must not get its own row", len(a.Targets), targetCount)
	}
	var sum int64
	for _, row := range a.Targets {
		sum += row.Total
	}
	if sum != a.Total {
		t.Errorf("sum(Targets[].Total) = %d, want Total = %d; the non-target index's damage must be redistributed, not lost", sum, a.Total)
	}
}
