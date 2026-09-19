package combine

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
)

func req(iters int, seed int64) api.SimRequest {
	return api.SimRequest{
		EngineVersion: "7779ebb", Spec: "warrior-fury",
		Character:  api.CharacterSpec{Name: "T", Race: "orc", Class: "warrior", Level: 60},
		Encounter:  api.DefaultEncounter(),
		Iterations: iters, RandomSeed: seed,
	}
}

// Splitting must divide the iterations exactly and offset each part's
// seed by the iterations of the parts before it, so four workers produce
// the same stream a serial run would.
func TestSplitDividesIterationsAndOffsetsSeeds(t *testing.T) {
	parts, err := Split(req(3000, 100), 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(parts) != 4 {
		t.Fatalf("got %d parts, want 4", len(parts))
	}
	var total int
	seed := int64(100)
	for i, p := range parts {
		total += p.Iterations
		if p.RandomSeed != seed {
			t.Errorf("part %d has seed %d, want %d", i, p.RandomSeed, seed)
		}
		seed += int64(p.Iterations)
	}
	if total != 3000 {
		t.Errorf("parts sum to %d iterations, want 3000", total)
	}
}

// A remainder goes to the first part, which is what the engine's own
// splitter does; anything else loses iterations.
func TestSplitHandlesARemainder(t *testing.T) {
	parts, err := Split(req(3000, 0), 7)
	if err != nil {
		t.Fatal(err)
	}
	var total int
	for _, p := range parts {
		total += p.Iterations
		if p.Iterations == 0 {
			t.Error("a part got zero iterations")
		}
	}
	if total != 3000 {
		t.Errorf("parts sum to %d, want 3000", total)
	}
}

func TestSplitNeverExceedsTheIterationCount(t *testing.T) {
	parts, err := Split(req(500, 0), 64)
	if err != nil {
		t.Fatal(err)
	}
	if len(parts) > 500 {
		t.Errorf("got %d parts for 500 iterations", len(parts))
	}
}

func TestSplitRejectsNonsense(t *testing.T) {
	if _, err := Split(req(3000, 0), 0); err == nil {
		t.Error("a zero split was accepted")
	}
	if _, err := Split(req(3000, 0), -1); err == nil {
		t.Error("a negative split was accepted")
	}
}

// Combining must pool, not average averages: two parts of 1,000 and
// 2,000 iterations weigh differently.
func TestResultsPoolsTheMean(t *testing.T) {
	parts := []api.SimResult{
		{IterationsRun: 1000, DPS: api.Estimate{Mean: 1000, StdDev: 0, Min: 900, Max: 1100}},
		{IterationsRun: 2000, DPS: api.Estimate{Mean: 1300, StdDev: 0, Min: 800, Max: 1500}},
	}
	got, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	if got.IterationsRun != 3000 {
		t.Errorf("IterationsRun = %d, want 3000", got.IterationsRun)
	}
	want := (1000*1000.0 + 2000*1300.0) / 3000.0
	if math.Abs(got.DPS.Mean-want) > 1e-9 {
		t.Errorf("Mean = %v, want the iteration-weighted %v", got.DPS.Mean, want)
	}
	if got.DPS.Min != 800 {
		t.Errorf("Min = %v, want 800", got.DPS.Min)
	}
	if got.DPS.Max != 1500 {
		t.Errorf("Max = %v, want 1500", got.DPS.Max)
	}
}

// The pooled standard deviation must account for the spread *between*
// parts as well as within them, or a split run reports a tighter error
// than a serial one and the page lies about its precision.
func TestResultsPoolsTheStdDev(t *testing.T) {
	parts := []api.SimResult{
		{IterationsRun: 1000, DPS: api.Estimate{Mean: 1000, StdDev: 100}},
		{IterationsRun: 1000, DPS: api.Estimate{Mean: 1200, StdDev: 100}},
	}
	got, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	// Within-group variance 10000, between-group variance 10000, so the
	// pooled standard deviation is sqrt(20000).
	want := math.Sqrt(20000)
	if math.Abs(got.DPS.StdDev-want) > 1e-6 {
		t.Errorf("StdDev = %v, want %v; a split run must not report a tighter spread than a serial one",
			got.DPS.StdDev, want)
	}
	wantErr := want / math.Sqrt(2000)
	if math.Abs(got.DPS.Error-wantErr) > 1e-6 {
		t.Errorf("Error = %v, want %v", got.DPS.Error, wantErr)
	}
}

func TestResultsRejectsAnEmptyOrFailedSet(t *testing.T) {
	if _, err := Results(nil); err == nil {
		t.Error("an empty set was combined")
	}
	bad := []api.SimResult{{IterationsRun: 100}, {IterationsRun: 100, Error: "boom"}}
	if _, err := Results(bad); err == nil {
		t.Error("a set containing a failed part was combined")
	}
}

// A part's summary is already a per-fight average, so combining two of
// them is a weighted mean of like quantities. Averaging them evenly when
// one part ran twice the iterations would misreport the damage table
// beside a correctly pooled DPS.
func TestResultsWeightsTheDamageTableByIterationShare(t *testing.T) {
	part := func(iters int, total int64) api.SimResult {
		return api.SimResult{
			IterationsRun: iters,
			DPS:           api.Estimate{Mean: float64(total)},
			Summary: summary.Summary{DamageDone: []summary.Actor{{
				GUID: "sim-player", Total: total, Effective: total,
				Abilities: []summary.Ability{
					{SpellID: 1, Total: total / 2, Effective: total / 2},
					{SpellID: 2, Total: total / 2, Effective: total / 2},
				},
			}}},
		}
	}
	got, err := Results([]api.SimResult{part(1000, 300), part(2000, 600)})
	if err != nil {
		t.Fatal(err)
	}
	// 300 * 1/3 + 600 * 2/3 = 500.
	if len(got.Summary.DamageDone) != 1 {
		t.Fatalf("DamageDone has %d actors, want 1", len(got.Summary.DamageDone))
	}
	a := got.Summary.DamageDone[0]
	if a.Total != 500 {
		t.Errorf("actor Total = %d, want the iteration-weighted 500", a.Total)
	}
	if a.Effective != 500 {
		t.Errorf("actor Effective = %d, want 500", a.Effective)
	}
	for _, ab := range a.Abilities {
		if ab.Total != 250 {
			t.Errorf("ability %d Total = %d, want 250", ab.SpellID, ab.Total)
		}
	}
	// The combined request says how many iterations stand behind it.
	if got.Request.Iterations != 3000 {
		t.Errorf("Request.Iterations = %d, want 3000", got.Request.Iterations)
	}
}

// One part is the whole run: its summary is returned untouched rather
// than scaled by a weight of one, which would round every row.
func TestResultsWithOnePartKeepsItsSummary(t *testing.T) {
	only := api.SimResult{
		IterationsRun: 500, DurationMS: 42,
		DPS:     api.Estimate{Mean: 1000, StdDev: 50, Min: 900, Max: 1100},
		Summary: summary.Summary{DamageDone: []summary.Actor{{GUID: "sim-player", Total: 7}}},
	}
	got, err := Results([]api.SimResult{only})
	if err != nil {
		t.Fatal(err)
	}
	if got.Summary.DamageDone[0].Total != 7 {
		t.Errorf("Total = %d, want the single part's 7 unscaled", got.Summary.DamageDone[0].Total)
	}
	if got.DPS.Mean != 1000 || got.DPS.StdDev != 50 {
		t.Errorf("DPS = %+v, want the single part's own", got.DPS)
	}
}

// The wall clock of a parallel run is the slowest part, not the sum: the
// parts ran at the same time.
func TestResultsTakesTheSlowestPartsWallClock(t *testing.T) {
	got, err := Results([]api.SimResult{
		{IterationsRun: 100, DurationMS: 90},
		{IterationsRun: 100, DurationMS: 250},
		{IterationsRun: 100, DurationMS: 120},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.DurationMS != 250 {
		t.Errorf("DurationMS = %d, want the slowest part's 250", got.DurationMS)
	}
}
