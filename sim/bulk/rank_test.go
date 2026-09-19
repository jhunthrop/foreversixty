package bulk

import (
	"errors"
	"math"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// resultsFor fabricates one finished result per request, with the DPS
// given. Nothing here runs the engine: staging is arithmetic and is
// tested as arithmetic.
func resultsFor(stage StageRequests, means []float64, stderr float64) []api.SimResult {
	errs := make([]float64, len(means))
	for i := range errs {
		errs[i] = stderr
	}
	return resultsForVaried(stage, means, errs)
}

// resultsForVaried is resultsFor with a PER-RESULT standard error,
// for cases a uniform error would hide: applyCut's overlap check and
// score's Delta.Error are both built from each result's own Error, and
// a fixture where every result shares one stderr cannot tell "used the
// right error" apart from "used any error, correctly or not" for a
// candidate that isn't itself.
func resultsForVaried(stage StageRequests, means, errs []float64) []api.SimResult {
	out := make([]api.SimResult, len(stage.Requests))
	for i := range stage.Requests {
		out[i] = api.SimResult{
			EngineVersion: enginever.Version,
			Request:       stage.Requests[i],
			Lane:          api.LaneBrowser,
			IterationsRun: stage.Iterations,
			DPS:           api.Estimate{Mean: means[i], StdDev: errs[i] * 10, Error: errs[i]},
			// A real run's Summary always carries the engine version,
			// and FightIndex is varied PER RESULT (0 for the equipped
			// set, i for the rest): finalResult copies Requests[0]'s
			// whole result (Task 17), and a fixture that stamped every
			// result's Summary identically could not tell "copied the
			// equipped result's" apart from "copied any result's".
			Summary: summary.Summary{EngineVersion: enginever.Version, FightIndex: i},
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
	// The fast ladder's second cut is Top: 10 over exactly 10 already-kept
	// combinations, so keep >= len(ranked) and applyCut's early return is
	// what fires here - worth asserting on its own, since a stage that
	// only ever shrinks would still look right by accident if this count
	// were never checked.
	if len(third.Combos) != 10 {
		t.Errorf("stage 2's cut kept %d of an already-cut 10, want all 10", len(third.Combos))
	}
	// Threaded again: stage 3's history is stage 1's entry PLUS stage
	// 2's, oldest first - Rank appends onto what it was handed rather
	// than replacing it.
	want = append(want, api.Stage{Iterations: 1000, Combos: len(next.Combos)})
	if !slices.Equal(third.Ran, want) {
		t.Errorf("stage 3's Ran = %+v, want %+v", third.Ran, want)
	}

	// Whether Rank's OWN clone of stage.Ran (rank.go, ahead of the
	// append that builds the next stage's history) protects against
	// aliasing is covered directly by
	// TestRankClonesStageRanBeforeAppending below, with a
	// stage.Ran built to actually have spare capacity to expose it.
	// Mutating next.Ran here after the fact would not do that: every
	// StageRequests.Ran leaving this package already came out of
	// stageRequests's own slices.Clone (plan.go), which always
	// returns a slice at exactly len == cap, so the append that
	// builds the FOLLOWING stage's history reallocates regardless of
	// whether rank.go's own clone is present - there is no public
	// path on which removing it changes next.Ran or third.Ran at all.
}

// rank.go's own clone of stage.Ran only has anything to protect
// against if stage.Ran itself carries spare capacity - never true
// along the public path, since stageRequests always hands one back at
// exactly len == cap (see the test above). This builds that
// impossible-in-practice case by hand, directly on the backing array,
// so the defensive clone is pinned by something that would actually
// fail if it were deleted: without it, Rank's append would write the
// next stage's entry straight into the caller's own array at the
// index one past stage.Ran's length, corrupting whatever ELSE views
// that same backing array - which a plain "is next.Ran's own value
// still correct" check can never observe, because everything Rank
// hands back is itself freshly cloned downstream regardless.
func TestRankClonesStageRanBeforeAppending(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	results := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 1)

	// backing has one live entry (what stage.Ran reports) and one
	// spare slot that only a shared-array view can see.
	backing := make([]api.Stage, 2)
	backing[0] = api.Stage{Iterations: 999, Combos: 999}
	sentinel := api.Stage{Iterations: -1, Combos: -1}
	backing[1] = sentinel
	stage.Ran = backing[:1:2] // len 1, cap 2: room for exactly one in-place append.

	if _, _, err := Rank(req, stage, results); err != nil {
		t.Fatal(err)
	}
	if full := backing[:2]; full[1] != sentinel {
		t.Errorf("Rank wrote into stage.Ran's spare capacity instead of cloning it first: backing[1] = %+v, want the untouched sentinel %+v", full[1], sentinel)
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
//
// A SINGLE candidate is what actually exercises the max(keep, 1) floor:
// round(1 * 0.25) = 0, and without the floor applyCut would index
// ranked[-1] computing the bar. Two candidates (round(2*0.25) =
// round(0.5) = 1) would pass this test on math.Round's own rounding
// alone, without ever reaching the floor.
func TestACutAlwaysKeepsAtLeastOne(t *testing.T) {
	req, stage := planOf(t, api.PrecisionFast, 1)
	next, _, err := Rank(req, stage, resultsFor(stage, []float64{1000, 1100}, 1))
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Combos) != 1 {
		t.Fatalf("kept %d of 1, want 1", len(next.Combos))
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

// applyCut's overlap check is against a FIXED bar, checked
// independently for every candidate past the cut - not a contiguous
// run starting there. Error varies candidate to candidate, so a
// tight, non-overlapping candidate can rank ABOVE a wide, overlapping
// one: the cut must still find the wide one past the tight one, and
// must still drop the tight one even though something below it
// survives.
func TestCutFindsAnOverlapPastANonOverlappingCandidate(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 12) // Normal's cut is Top: 10, SlackSE: 2.
	means := []float64{5000}                         // equipped: clearly separate, never itself ranked.
	errs := []float64{1}
	// Nine clear leaders, unambiguously inside the cut.
	for i := 0; i < 9; i++ {
		means = append(means, float64(2000-10*i))
		errs = append(errs, 1)
	}
	// #10 (helmIDs[9]): the cut's own boundary. bar = 1000 - 2*1 = 998.
	means = append(means, 1000)
	errs = append(errs, 1)
	// #11 (helmIDs[10]): ranked just below the boundary with a TIGHT
	// error - 996 + 2*0.5 = 997 < 998, so it does NOT overlap and must
	// be cut.
	means = append(means, 996)
	errs = append(errs, 0.5)
	// #12 (helmIDs[11]): ranked lower still, but a WIDE error -
	// 995 + 2*10 = 1015 >= 998, so it DOES overlap and must survive,
	// even though #11 immediately above it did not.
	means = append(means, 995)
	errs = append(errs, 10)

	next, _, err := Rank(req, stage, resultsForVaried(stage, means, errs))
	if err != nil {
		t.Fatal(err)
	}
	// The top 10, minus none of them, plus #12 - eleven total. A cut
	// that stops scanning at the first failure (#11) would keep only
	// 10; a cut that keeps a contiguous run through the last PASSING
	// index would wrongly keep #11 too and report 12.
	if len(next.Combos) != 11 {
		t.Fatalf("kept %d combinations, want 11 (the top 10 plus #12, skipping #11)", len(next.Combos))
	}
	cut, survived := helmIDs[10], helmIDs[11]
	var sawSurvivor bool
	for _, c := range next.Combos {
		id := c.Substitutions[0].ItemID
		if id == cut {
			t.Errorf("kept item %d (#11, a tight non-overlapping candidate) that should have been cut", cut)
		}
		if id == survived {
			sawSurvivor = true
		}
	}
	if !sawSurvivor {
		t.Errorf("dropped item %d (#12, a wide overlapping candidate ranked below a cut one) that should have survived", survived)
	}
}

// score's Delta is the one piece of genuinely novel arithmetic this
// task introduces, and Task 17's final result is built entirely from
// it, so it gets its own direct assertions rather than being trusted
// through Rank's cut-count checks alone. Errors vary per candidate so
// a swapped mean/error pair, or a delta built from the WRONG
// candidate's error, would be caught.
func TestScoreComputesThePairedQuadratureDelta(t *testing.T) {
	_, stage := planOf(t, api.PrecisionNormal, 3)
	means := []float64{1000, 990, 1010, 1005}
	errs := []float64{2, 3, 4, 1}
	ranked, equipped, err := score(stage, resultsForVaried(stage, means, errs))
	if err != nil {
		t.Fatal(err)
	}
	if equipped.Mean != 1000 || equipped.Error != 2 {
		t.Fatalf("equipped = %+v, want mean 1000, error 2", equipped)
	}
	want := map[int]struct{ mean, errv, stddev float64 }{
		helmIDs[0]: {990 - 1000, math.Hypot(3, 2), math.Hypot(30, 20)},
		helmIDs[1]: {1010 - 1000, math.Hypot(4, 2), math.Hypot(40, 20)},
		helmIDs[2]: {1005 - 1000, math.Hypot(1, 2), math.Hypot(10, 20)},
	}
	if len(ranked) != len(want) {
		t.Fatalf("score returned %d combinations, want %d", len(ranked), len(want))
	}
	for _, s := range ranked {
		id := s.Combo.Substitutions[0].ItemID
		w, ok := want[id]
		if !ok {
			t.Fatalf("unexpected combination for item %d", id)
		}
		if s.Delta.Mean != w.mean {
			t.Errorf("item %d: Delta.Mean = %v, want %v", id, s.Delta.Mean, w.mean)
		}
		if math.Abs(s.Delta.Error-w.errv) > 1e-9 {
			t.Errorf("item %d: Delta.Error = %v, want %v", id, s.Delta.Error, w.errv)
		}
		if math.Abs(s.Delta.StdDev-w.stddev) > 1e-9 {
			t.Errorf("item %d: Delta.StdDev = %v, want %v", id, s.Delta.StdDev, w.stddev)
		}
	}
}

// The equipped set is not always the winner. Every other fixture in
// this file puts a candidate above it; this one puts the baseline
// above every candidate and checks that neither score nor Rank treats
// that as a special case - the deltas go genuinely negative (never
// clamped to zero) and the cut and next-stage plumbing proceed exactly
// as they would if something had won.
func TestEquippedBeingTheBestResultStillRanksAndCuts(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 14)
	means := []float64{2000} // the equipped set beats every candidate here.
	for i := 0; i < 14; i++ {
		means = append(means, float64(1000+10*i)) // best candidate still trails it.
	}
	results := resultsFor(stage, means, 1)

	ranked, equipped, err := score(stage, results)
	if err != nil {
		t.Fatal(err)
	}
	if equipped.Mean != 2000 {
		t.Fatalf("equipped.Mean = %v, want 2000", equipped.Mean)
	}
	for _, s := range ranked {
		want := s.DPS.Mean - equipped.Mean
		if s.Delta.Mean != want {
			t.Errorf("item %d: Delta.Mean = %v, want %v", s.Combo.Substitutions[0].ItemID, s.Delta.Mean, want)
		}
		if s.Delta.Mean >= 0 {
			t.Errorf("item %d: Delta.Mean = %v, want negative - every candidate trails the equipped set here", s.Combo.Substitutions[0].ItemID, s.Delta.Mean)
		}
	}

	next, final, err := Rank(req, stage, results)
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 2 produced a final result")
	}
	if len(next.Combos) != 10 {
		t.Errorf("kept %d combinations, want the top 10", len(next.Combos))
	}
}

// Results the right SIZE but the wrong ORDER pass every other check
// silently and misattribute every DPS to the wrong combination -
// exactly what ErrStageMismatch's own doc comment says never happens.
// Stage requests run through a worker pool with no ordering guarantee
// of their own (spec 10.2), so this is the only defence.
func TestRankRefusesResultsOutOfOrder(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	results := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 5)
	results[1], results[2] = results[2], results[1]
	if _, _, err := Rank(req, stage, results); !errors.Is(err, ErrStageMismatch) {
		t.Errorf("Rank = %v, want ErrStageMismatch for misordered results", err)
	}
}

// A non-finite DPS is refused rather than silently degrading the rest
// of the pipeline: a NaN mean sorts as tied with everything (breaking
// score's ordering) and a NaN bar makes applyCut keep everyone, which
// is exactly the kind of silent, wrong answer this package's package
// doc says never to produce.
func TestScoreRefusesNonFiniteDPS(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	results := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 5)
	results[2].DPS.Mean = math.NaN()
	if _, _, err := Rank(req, stage, results); err == nil {
		t.Fatal("Rank accepted a NaN DPS mean")
	}
}

// High is two rungs (1,000 then 10,000) with a cut of Top: 20, neither
// of which any other test in this file reaches.
func TestHighKeepsTopTwenty(t *testing.T) {
	req, stage := planOf(t, api.PrecisionHigh, 30)
	means := []float64{1000}
	for i := 0; i < 30; i++ {
		means = append(means, float64(1400-10*i))
	}
	if stage.Iterations != 1000 {
		t.Fatalf("stage 1 runs %d iterations, want 1000", stage.Iterations)
	}
	next, final, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 2 produced a final result")
	}
	if next.Stage != 2 || next.Iterations != 10000 {
		t.Errorf("next stage is %d at %d iterations", next.Stage, next.Iterations)
	}
	if len(next.Combos) != 20 {
		t.Errorf("kept %d of 30, want the top 20", len(next.Combos))
	}
}

// The last rung produces a SimResult, not another stage.
func TestTheLastRungProducesTheResult(t *testing.T) {
	req, first := planOf(t, api.PrecisionNormal, 4)
	means := []float64{1000, 1100, 1050, 1020, 990}
	next, final, err := Rank(req, first, resultsFor(first, means, 2))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil || next == nil {
		t.Fatal("the first of two rungs finished the run")
	}

	last := *next
	lastMeans := means[:len(last.Requests)]
	next, final, err = Rank(req, last, resultsFor(last, lastMeans, 2))
	if err != nil {
		t.Fatal(err)
	}
	if next != nil || final == nil {
		t.Fatal("the last rung did not finish the run")
	}
	if final.Equipped == nil || final.Equipped.Mean != 1000 {
		t.Errorf("equipped = %+v, want the baseline 1000", final.Equipped)
	}
	if final.DPS.Mean != 1000 {
		t.Errorf("the result's own DPS is %v; a bulk result's headline number is the equipped set's", final.DPS.Mean)
	}
	if final.IterationsRun != last.Iterations {
		t.Errorf("iterations_run = %d, want the last stage's %d", final.IterationsRun, last.Iterations)
	}
	if final.Request.Bulk == nil {
		t.Error("the result does not carry the request that produced it")
	}
	if len(final.Combos) != len(last.Combos) {
		t.Fatalf("%d combos in the result and %d in the last stage", len(final.Combos), len(last.Combos))
	}
	if final.Combos[0].DPS.Mean != 1100 || final.Combos[0].Delta.Mean != 100 {
		t.Errorf("the leader is %+v", final.Combos[0])
	}
	if len(final.Combos[0].Substitutions) == 0 {
		t.Error("the leader carries no substitution chips")
	}
	// Every stage the ladder ran, in order, with what it ran.
	if len(final.Stages) != 2 {
		t.Fatalf("stages = %+v, want two", final.Stages)
	}
	if final.Stages[0].Iterations != 1000 || final.Stages[0].Combos != 4 {
		t.Errorf("stage 1 = %+v", final.Stages[0])
	}
	if final.Stages[1].Iterations != 3000 || final.Stages[1].Combos != len(last.Combos) {
		t.Errorf("stage 2 = %+v", final.Stages[1])
	}
	// The summary is the EQUIPPED set's, not any candidate's: a bulk
	// report renders the baseline character's breakdown beside the
	// ranking. FightIndex is stamped per result (see resultsForVaried),
	// so this tells "copied results[0]'s Summary" apart from "copied
	// some Summary" - a check on EngineVersion alone cannot, since the
	// fixture gives every result the same one.
	if final.Summary.EngineVersion == "" {
		t.Error("the result carries no summary")
	}
	if final.Summary.FightIndex != 0 {
		t.Errorf("Summary.FightIndex = %d, want 0 (the equipped set's, not a candidate's)", final.Summary.FightIndex)
	}
}

// "Within error" is an overlap of delta intervals with the group's
// leader, and the page ranks a group the same. Getting this wrong
// makes noise look like a decision.
func TestWithinErrorGroups(t *testing.T) {
	cases := []struct {
		name   string
		means  []float64 // the equipped set first
		stderr float64
		groups []int // per combination, best first
	}{
		{
			name:   "three clearly separated",
			means:  []float64{1000, 1100, 1050, 1010},
			stderr: 1,
			groups: []int{0, 1, 2},
		},
		{
			name:   "all one answer",
			means:  []float64{1000, 1100, 1099, 1098},
			stderr: 40,
			groups: []int{0, 0, 0},
		},
		{
			name:   "two tied, then one apart",
			means:  []float64{1000, 1100, 1098, 1000},
			stderr: 2,
			groups: []int{0, 0, 1},
		},
		{
			name:   "a single combination is its own group",
			means:  []float64{1000, 1100},
			stderr: 2,
			groups: []int{0},
		},
		// Leader-anchored grouping (rank.go's group, comparing every
		// row to its GROUP'S leader) and predecessor-chaining (comparing
		// each row only to its immediate neighbour) agree on every case
		// above - none of them actually exercises the reason rank.go's
		// doc comment gives for anchoring on the leader instead of
		// chaining. This one does: 1100 overlaps 1095 which overlaps
		// 1090, so a chain would walk all three into one band even
		// though 1100 and 1090 (a real 10-DPS-over-error-2 gap) do not
		// overlap each other. Leader-anchored gives [0,0,1]; chaining
		// gives [0,0,0], silently reporting a 10-DPS gap as a tie.
		{
			name:   "a chain of overlapping neighbours is not one tie",
			means:  []float64{1000, 1100, 1095, 1090},
			stderr: 2,
			groups: []int{0, 0, 1},
		},
		// The same drift, sharper: six candidates 4 DPS apart with an
		// error of 2 chain end to end (each overlaps its neighbour),
		// but the group is capped at two members before the interval
		// stops reaching back to 1120. Leader-anchored gives
		// [0,0,1,1,2,2]; chaining would walk all six into [0,0,0,0,0,0].
		{
			name:   "a longer chain still breaks into bands",
			means:  []float64{1000, 1120, 1116, 1112, 1108, 1104, 1100},
			stderr: 2,
			groups: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req, first := planOf(t, api.PrecisionNormal, len(c.means)-1)
			next, _, err := Rank(req, first, resultsFor(first, c.means, c.stderr))
			if err != nil {
				t.Fatal(err)
			}
			last := *next
			_, final, err := Rank(req, last, resultsFor(last, c.means[:len(last.Requests)], c.stderr))
			if err != nil {
				t.Fatal(err)
			}
			if len(final.Combos) != len(c.groups) {
				t.Fatalf("%d combos, want %d", len(final.Combos), len(c.groups))
			}
			for i, want := range c.groups {
				if final.Combos[i].Group != want {
					t.Errorf("combo %d (%v DPS) is group %d, want %d",
						i, final.Combos[i].DPS.Mean, final.Combos[i].Group, want)
				}
			}
		})
	}
}

// Three rungs, so the result's stage list is three entries and each
// says what that rung actually ran.
func TestAThreeRungLadderReportsEveryStage(t *testing.T) {
	req, first := planOf(t, api.PrecisionFast, 40)
	means := []float64{1000}
	for i := 0; i < 40; i++ {
		means = append(means, float64(1400-10*i))
	}
	second, _, err := Rank(req, first, resultsFor(first, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	third, _, err := Rank(req, *second, resultsFor(*second, means[:len(second.Requests)], 1))
	if err != nil {
		t.Fatal(err)
	}
	_, final, err := Rank(req, *third, resultsFor(*third, means[:len(third.Requests)], 1))
	if err != nil {
		t.Fatal(err)
	}
	if final == nil {
		t.Fatal("the third rung did not finish the run")
	}
	want := []api.Stage{
		{Iterations: 100, Combos: 40},
		{Iterations: 1000, Combos: len(second.Combos)},
		{Iterations: 3000, Combos: len(third.Combos)},
	}
	if len(final.Stages) != 3 {
		t.Fatalf("stages = %+v", final.Stages)
	}
	for i, w := range want {
		if final.Stages[i] != w {
			t.Errorf("stage %d = %+v, want %+v", i+1, final.Stages[i], w)
		}
	}
}
