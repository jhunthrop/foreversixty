package combine

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

func req(iters int, seed int64) api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version, Spec: "warrior-fury",
		Source:     api.CharacterSource{Kind: api.SourceManual},
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

// A worker pool asking for more parts than there are iterations is the
// case the clamp exists for, and nothing exercised it: the old version
// split 500 iterations 64 ways, where 64 < 500 and the clamp is never
// taken, so the assertion could not fire.
func TestSplitNeverExceedsTheIterationCount(t *testing.T) {
	parts, err := Split(req(8, 0), 64)
	if err != nil {
		t.Fatal(err)
	}
	if len(parts) != 8 {
		t.Fatalf("got %d parts for 8 iterations, want 8: a part with no iterations is a worker with nothing to do", len(parts))
	}
	var total int
	for i, p := range parts {
		if p.Iterations != 1 {
			t.Errorf("part %d has %d iterations, want 1", i, p.Iterations)
		}
		total += p.Iterations
	}
	if total != 8 {
		t.Errorf("the parts sum to %d, want 8", total)
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

// actorWith builds one part: an actor with the abilities given, in the
// order given, which is what the adapter's damage-descending sort
// produces and what differs between parts.
func actorWith(iters int, guid string, abilities ...summary.Ability) api.SimResult {
	return api.SimResult{
		IterationsRun: iters,
		Summary: summary.Summary{DamageDone: []summary.Actor{{
			GUID: guid, Name: "Sim", Abilities: abilities,
		}}},
	}
}

func ab(id int64, total int64) summary.Ability {
	return summary.Ability{SpellID: id, Name: "spell", Total: total, Effective: total}
}

func totals(t *testing.T, res api.SimResult) map[int64]int64 {
	t.Helper()
	if len(res.Summary.DamageDone) == 0 {
		t.Fatal("no actors in the combined summary")
	}
	out := map[int64]int64{}
	for _, a := range res.Summary.DamageDone[0].Abilities {
		out[a.SpellID] = a.Total
	}
	return out
}

// The adapter sorts each part's abilities by damage, so two parts of one
// split run order near-ties differently. Merging by slice position then
// adds one spell's damage to another's with nothing to show for it, so
// the merge is keyed on the row's identity instead.
func TestResultsMergesRowsByIdentityNotPosition(t *testing.T) {
	parts := []api.SimResult{
		actorWith(1000, "sim-player", ab(111, 600), ab(222, 400)),
		actorWith(1000, "sim-player", ab(222, 700), ab(111, 300)),
	}
	got, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	// Half of each part: 111 is (600+300)/2, 222 is (400+700)/2. Merging
	// by index would report 650 and 350.
	want := map[int64]int64{111: 450, 222: 550}
	for id, w := range want {
		if got := totals(t, got)[id]; got != w {
			t.Errorf("spell %d = %d, want %d", id, got, w)
		}
	}
	// And the table comes out in the adapter's own order, damage first.
	rows := got.Summary.DamageDone[0].Abilities
	if len(rows) != 2 || rows[0].SpellID != 222 || rows[1].SpellID != 111 {
		t.Errorf("rows = %v, want 222 then 111 by damage descending", rows)
	}
}

// A proc that fires in one part and not another is a row one part does
// not have. It must be added at its own identity, not folded into
// whatever happened to sit at its index.
func TestResultsAddsARowOnlyALaterPartHas(t *testing.T) {
	parts := []api.SimResult{
		actorWith(1000, "sim-player", ab(111, 600)),
		actorWith(1000, "sim-player", ab(999, 900), ab(111, 500)),
	}
	// A pet only the later part saw is the same problem one level up.
	parts[1].Summary.DamageDone = append(parts[1].Summary.DamageDone, summary.Actor{
		GUID: "sim-player-pet-0", Name: "Pet", Total: 200,
		Abilities: []summary.Ability{ab(777, 200)},
	})

	got, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	rows := totals(t, got)
	if len(rows) != 2 {
		t.Errorf("the combined table has %d rows, want 2: %v", len(rows), rows)
	}
	if rows[111] != 550 {
		t.Errorf("spell 111 = %d, want (600+500)/2 = 550", rows[111])
	}
	if rows[999] != 450 {
		t.Errorf("spell 999 = %d, want 900/2 = 450; the proc row was folded away", rows[999])
	}
	if len(got.Summary.DamageDone) != 2 {
		t.Fatalf("%d actors, want the player and the pet only the later part saw",
			len(got.Summary.DamageDone))
	}
	pet := got.Summary.DamageDone[1]
	if pet.GUID != "sim-player-pet-0" || pet.Total != 100 {
		t.Errorf("pet row = %+v, want sim-player-pet-0 at 200/2 = 100", pet)
	}
}

// Results must not write into what it was given: the same parts combined
// twice must give the same answer, and a caller that keeps its parts
// must still have them.
func TestResultsDoesNotTouchItsInput(t *testing.T) {
	parts := []api.SimResult{
		actorWith(1000, "sim-player", ab(111, 600), ab(222, 400)),
		actorWith(1000, "sim-player", ab(222, 700), ab(111, 300)),
	}
	parts[0].Summary.DamageDone[0].Total = 1000
	parts[0].Summary.DamageDone[0].Targets = []summary.Pair{{GUID: "t1", Total: 1000}}
	parts[0].Summary.DamageDone[0].Series = []int64{10, 20}
	parts[0].Summary.DamageDone[0].Abilities[0].Misses = map[string]int64{"MISS": 4}

	first, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := totals(t, second), totals(t, first); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("a second Results gave %v, want the same %v", got, want)
	}

	src := parts[0].Summary.DamageDone[0]
	if src.Total != 1000 {
		t.Errorf("parts[0] actor Total = %d, want its own 1000", src.Total)
	}
	if src.Abilities[0].SpellID != 111 || src.Abilities[0].Total != 600 {
		t.Errorf("parts[0] first ability = %+v, want spell 111 at 600", src.Abilities[0])
	}
	if len(src.Targets) != 1 || src.Targets[0].Total != 1000 {
		t.Errorf("parts[0] targets = %v, want its own", src.Targets)
	}
	if len(src.Series) != 2 || src.Series[0] != 10 || src.Series[1] != 20 {
		t.Errorf("parts[0] series = %v, want its own", src.Series)
	}
	if src.Abilities[0].Misses["MISS"] != 4 {
		t.Errorf("parts[0] misses = %v, want its own", src.Abilities[0].Misses)
	}
}

// full is a whole part: the fields Results reads to decide the parts
// belong to one run, plus a one-row damage table to merge.
func full(seed int64, iters int, mut func(*api.SimResult)) api.SimResult {
	r := api.SimResult{
		EngineVersion: enginever.Version,
		Request:       req(iters, seed),
		Lane:          api.LaneBrowser,
		IterationsRun: iters,
		DPS:           api.Estimate{Mean: 1000, StdDev: 50, Min: 900, Max: 1100},
		Summary: summary.Summary{DamageDone: []summary.Actor{{
			GUID: "sim-player", Total: 900, Effective: 900,
			Abilities: []summary.Ability{
				{SpellID: 1, Total: 500, Effective: 500, Hits: 5, Crits: 1},
				{SpellID: 2, Total: 400, Effective: 400, Hits: 3, Crits: 1},
			},
		}}},
	}
	if mut != nil {
		mut(&r)
	}
	return r
}

// Results validated only "non-empty, no error, iterations > 0" and then
// adopted part zero's request, spec and engine version wholesale. Two
// different specs, or a mix of engine versions, pooled into one
// authoritative-looking DPS and nothing downstream could tell.
func TestResultsRefusesPartsFromDifferentRuns(t *testing.T) {
	cases := []struct {
		name string
		mut  func(*api.SimResult)
	}{
		{"another engine", func(r *api.SimResult) { r.EngineVersion = "deadbee" }},
		{"another spec", func(r *api.SimResult) { r.Request.Spec = "mage-frost" }},
		{"another character", func(r *api.SimResult) { r.Request.Character.Race = "troll" }},
		{"another encounter", func(r *api.SimResult) { r.Request.Encounter.DurationSec = 300 }},
		{"different gear", func(r *api.SimResult) {
			r.Request.Character.Gear = []api.GearSlot{{Slot: "head", ItemID: 1}}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Results([]api.SimResult{full(0, 750, nil), full(750, 750, tc.mut)})
			if !errors.Is(err, ErrMixedParts) {
				t.Errorf("Results returned %v, want ErrMixedParts", err)
			}
		})
	}

	// The two fields a split IS allowed to vary must still combine.
	if _, err := Results([]api.SimResult{full(0, 750, nil), full(750, 250, nil)}); err != nil {
		t.Errorf("Results refused two honest parts of one run: %v", err)
	}

	// A part that was stopped is not a share of a finished run either:
	// pooling it reports a full run's precision over a partial sample.
	_, err := Results([]api.SimResult{full(0, 750, nil), full(750, 750, func(r *api.SimResult) { r.Aborted = true })})
	if err == nil {
		t.Error("Results pooled a stopped part")
	}
}

// The adapter guarantees an actor's total is the sum of its ability
// rows; the report reads both. Weighing the two independently broke
// that by a rounding error per row, and truncating biased every row low.
func TestCombinedActorTotalsAreTheSumOfTheirRows(t *testing.T) {
	// Four parts with a total that does not divide evenly by the
	// weights, so the rounding is doing real work.
	parts := []api.SimResult{
		full(0, 751, nil), full(751, 750, nil), full(1501, 750, nil), full(2251, 749, nil),
	}
	out, err := Results(parts)
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range out.Summary.DamageDone {
		var sum, effective int64
		for _, ab := range a.Abilities {
			sum += ab.Total
			effective += ab.Effective
		}
		if a.Total != sum {
			t.Errorf("actor %q totals %d but its rows sum to %d", a.GUID, a.Total, sum)
		}
		if a.Effective != effective {
			t.Errorf("actor %q effective is %d but its rows sum to %d", a.GUID, a.Effective, effective)
		}
	}
	// Each part reported 900 per fight, so the weighted mean is 900:
	// a truncating weigh gave 897 or worse.
	if got := out.Summary.DamageDone[0].Total; got != 900 {
		t.Errorf("the combined actor total is %d, want 900; the weights sum to one, so nothing should be lost", got)
	}
}

// Min and Max are extremes over the parts. A part that reported no
// distribution at all - Max zero - used to pull the run's minimum to
// zero, a figure no iteration produced.
func TestResultsIgnoresAPartWithNoDistribution(t *testing.T) {
	out, err := Results([]api.SimResult{
		full(0, 750, nil),
		full(750, 750, func(r *api.SimResult) { r.DPS = api.Estimate{Mean: 1000} }),
	})
	if err != nil {
		t.Fatal(err)
	}
	if out.DPS.Min != 900 || out.DPS.Max != 1100 {
		t.Errorf("DPS range = [%v, %v], want [900, 1100]", out.DPS.Min, out.DPS.Max)
	}
}
