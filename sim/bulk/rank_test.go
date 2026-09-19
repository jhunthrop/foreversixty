package bulk

import (
	"errors"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// resultsFor fabricates one finished result per request, with the DPS
// given. Nothing here runs the engine: staging is arithmetic and is
// tested as arithmetic.
func resultsFor(stage StageRequests, means []float64, stderr float64) []api.SimResult {
	out := make([]api.SimResult, len(stage.Requests))
	for i := range stage.Requests {
		out[i] = api.SimResult{
			EngineVersion: enginever.Version,
			Request:       stage.Requests[i],
			Lane:          api.LaneBrowser,
			IterationsRun: stage.Iterations,
			DPS:           api.Estimate{Mean: means[i], StdDev: stderr * 10, Error: stderr},
		}
	}
	return out
}

// planOf is a stage with n distinguishable candidates: n helms, each a
// real row, so expansion keeps them all in their own slot.
func planOf(t *testing.T, precision string, n int) (api.SimRequest, StageRequests) {
	t.Helper()
	req := base()
	final, _ := api.FinalIterations(precision)
	req.Iterations = final
	req.Bulk = &api.BulkSpec{Mode: api.KindGear, Precision: precision, Cap: api.Caps[api.LaneServer]}
	for i := 0; i < n; i++ {
		req.Bulk.Candidates = append(req.Bulk.Candidates, candidate("head", helmIDs[i]))
	}
	stage, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(stage.Combos) != n {
		t.Fatalf("planned %d combinations, want %d; helmIDs may be stale", len(stage.Combos), n)
	}
	return req, stage
}

func TestRankRefusesResultsThatDoNotMatchTheStage(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	short := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 5)[:3]
	if _, _, err := Rank(req, stage, short); !errors.Is(err, ErrStageMismatch) {
		t.Errorf("Rank = %v, want ErrStageMismatch", err)
	}
}

func TestRankPropagatesAFailedRun(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	results := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 5)
	results[2].Error = "the engine died"
	_, _, err := Rank(req, stage, results)
	if err == nil {
		t.Fatal("Rank ignored a failed run")
	}
}

// Normal is two rungs: 1,000 then 3,000, keeping the top ten plus
// anything still overlapping the tenth.
func TestNormalKeepsTheTopTenPlusTies(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 14)
	// The equipped set, then fourteen combinations 10 DPS apart, so
	// nothing is a tie at an error of 1.
	means := []float64{1000}
	for i := 0; i < 14; i++ {
		means = append(means, float64(1100-10*i))
	}
	next, final, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 2 produced a final result")
	}
	if next.Stage != 2 || next.Iterations != 3000 {
		t.Errorf("next stage is %d at %d iterations", next.Stage, next.Iterations)
	}
	if len(next.Combos) != 10 {
		t.Errorf("kept %d combinations, want the top 10", len(next.Combos))
	}
	if len(next.Requests) != 11 {
		t.Errorf("next stage has %d requests, want 10 plus the equipped set", len(next.Requests))
	}

	// With a large error every candidate overlaps the tenth, so every
	// one survives: the cut keeps the top ten PLUS the ties, and
	// throwing a winner away over noise is what the slack exists to
	// prevent.
	next, _, err = Rank(req, stage, resultsFor(stage, means, 50))
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Combos) != 14 {
		t.Errorf("kept %d combinations at an error of 50, want all 14", len(next.Combos))
	}
}

// Fast is three rungs, and the first keeps a quarter.
func TestFastKeepsAQuarterThenTheTopTen(t *testing.T) {
	req, stage := planOf(t, api.PrecisionFast, 40)
	means := []float64{1000}
	for i := 0; i < 40; i++ {
		means = append(means, float64(1400-10*i))
	}
	if stage.Iterations != 100 {
		t.Fatalf("stage 1 runs %d iterations, want 100", stage.Iterations)
	}
	next, final, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 3 produced a final result")
	}
	if next.Stage != 2 || next.Iterations != 1000 {
		t.Errorf("next stage is %d at %d iterations", next.Stage, next.Iterations)
	}
	if len(next.Combos) != 10 {
		t.Errorf("kept %d of 40, want a quarter", len(next.Combos))
	}
	// Ran (contract A10): stage 1 carried none, so stage 2's history is
	// exactly the stage that was just run - the FIRST stage's own
	// iteration count and combination count, not the second stage's.
	want := []api.Stage{{Iterations: 100, Combos: 40}}
	if !slices.Equal(next.Ran, want) {
		t.Errorf("stage 2's Ran = %+v, want %+v", next.Ran, want)
	}

	means2 := means[:len(next.Requests)]
	third, final, err := Rank(req, *next, resultsFor(*next, means2, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 2 of 3 produced a final result")
	}
	if third.Stage != 3 || third.Iterations != 3000 {
		t.Errorf("third stage is %d at %d iterations", third.Stage, third.Iterations)
	}
	// Threaded again: stage 3's history is stage 1's entry PLUS stage
	// 2's, oldest first - Rank appends onto what it was handed rather
	// than replacing it.
	want = append(want, api.Stage{Iterations: 1000, Combos: len(next.Combos)})
	if !slices.Equal(third.Ran, want) {
		t.Errorf("stage 3's Ran = %+v, want %+v", third.Ran, want)
	}

	// The append above must not have mutated next.Ran through a shared
	// backing array: stage 2's own history is still just stage 1's.
	if len(next.Ran) != 1 {
		t.Errorf("building stage 3's history mutated stage 2's Ran: %+v", next.Ran)
	}
}

// On the ladder's LAST rung Rank stops cutting and finishes instead:
// next is nil and final is not, which is the one decision the cut
// arithmetic above never has to make for itself.
func TestRankFinishesOnTheLastStage(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	means := []float64{1000, 1010, 1020, 1030, 1040}
	next, final, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 2 produced a final result")
	}

	// Stage 2 is the normal ladder's last rung.
	means2 := means[:len(next.Requests)]
	next2, final, err := Rank(req, *next, resultsFor(*next, means2, 1))
	if err != nil {
		t.Fatal(err)
	}
	if next2 != nil {
		t.Error("the last stage produced a next stage instead of finishing")
	}
	if final == nil {
		t.Fatal("the last stage produced no final result")
	}
	if final.Request.Character.Name != req.Character.Name {
		t.Error("the final result does not carry the original request")
	}
	if final.Equipped == nil || final.Equipped.Mean != means[0] {
		t.Errorf("final.Equipped = %v, want a mean of %v", final.Equipped, means[0])
	}
}

// A cut never keeps nothing, however small the fraction.
func TestACutAlwaysKeepsAtLeastOne(t *testing.T) {
	req, stage := planOf(t, api.PrecisionFast, 2)
	next, _, err := Rank(req, stage, resultsFor(stage, []float64{1000, 1100, 900}, 1))
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Combos) < 1 {
		t.Fatal("a quarter of two kept nothing")
	}
}

// The survivors are the BEST ones, in order, and the next stage's
// requests are parallel to them with the equipped set still first.
func TestTheSurvivorsAreTheBestInOrder(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 14)
	means := []float64{1000}
	for i := 0; i < 14; i++ {
		// Ascending, so the LAST candidate is the best and a planner
		// that kept the first ten would be caught.
		means = append(means, float64(1000+10*i))
	}
	next, _, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	best := next.Combos[0].Substitutions[0].ItemID
	if best != helmIDs[13] {
		t.Errorf("the leader is item %d, want %d", best, helmIDs[13])
	}
	if next.Requests[0].Character.Gear[0].ItemID != req.Character.Gear[0].ItemID {
		t.Error("the equipped set is not request 0 of the next stage")
	}
	for i, c := range next.Combos {
		if next.Requests[i+1].Character.Gear[0].ItemID != c.Substitutions[0].ItemID {
			t.Errorf("request %d does not match combo %d", i+1, i)
		}
	}
}
