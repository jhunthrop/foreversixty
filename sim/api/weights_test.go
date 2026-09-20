package api

import (
	"slices"
	"strings"
	"testing"
)

func weightsReq() SimRequest {
	req := runReq()
	req.Weights = &WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit", "hit"},
		Reference: "attack_power",
	}
	return req
}

func TestWeightsValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string
	}{
		{"the happy path", func(*SimRequest) {}, ""},
		{"no stats", func(r *SimRequest) { r.Weights.Stats = nil }, "at least one stat"},
		{"no reference", func(r *SimRequest) { r.Weights.Reference = "" }, "weights.reference"},
		{"a reference that is not weighed", func(r *SimRequest) { r.Weights.Reference = "strength" }, "one of the stats it weighs"},
		{"a repeated stat", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", "crit"}
			r.Weights.Reference = "crit"
		}, "listed twice"},
		{"an empty stat id", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", ""}
			r.Weights.Reference = "crit"
		}, "empty stat id"},
		{"a stat id outside the vocabulary", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", "made_up_stat"}
			r.Weights.Reference = "crit"
		}, "not a known stat id"},
		{"more stats than the vocabulary has", func(r *SimRequest) {
			r.Weights.Stats = append([]string{}, KnownStats...)
			r.Weights.Stats = append(r.Weights.Stats, "one_too_many")
			r.Weights.Reference = "crit"
		}, "the vocabulary only has"},
		{"weights and bulk together", func(r *SimRequest) {
			r.Bulk = &BulkSpec{Mode: KindGear, Precision: PrecisionNormal, Cap: 10,
				Candidates: []Candidate{{Slot: "head", ItemID: 1, Origin: OriginBag}}}
		}, "a request is one kind"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := weightsReq()
			c.edit(&req)
			err := req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// The stat ids themselves are the engine's; sim/request resolves them
// and IDS.md publishes them. KnownStats is a copy of that same list -
// needed to bound a request's size at the envelope's own boundary
// without importing the engine's proto - not a second source of
// truth for it; sim/request's
// TestAPIKnownStatsMatchTheGeneratedVocabulary keeps the two in step.
func TestWeightsKindAndShape(t *testing.T) {
	req := weightsReq()
	if req.Kind() != KindWeights {
		t.Errorf("Kind() = %q, want %q", req.Kind(), KindWeights)
	}
	w := StatWeight{Stat: "crit", Weight: 1.0, Error: 0.04}
	if w.Stat != "crit" || w.Weight != 1.0 || w.Error != 0.04 {
		t.Errorf("StatWeight round trip: %+v", w)
	}
}

// KnownStats is an interface other packages and the generator's own
// cross-check depend on: sorted, and with no id repeated.
func TestKnownStatsIsSortedAndUnique(t *testing.T) {
	if !slices.IsSorted(KnownStats) {
		t.Error("KnownStats is not sorted")
	}
	seen := map[string]bool{}
	for _, id := range KnownStats {
		if seen[id] {
			t.Errorf("KnownStats lists %q twice", id)
		}
		seen[id] = true
	}
}

// TestWeightsIterationsMatchesTheEnginesArithmetic pins
// WeightsIterations against a hand-worked case: the engine's
// buildStatWeightRequests (sim/core/statweight.go) halves Iterations
// once for RNG parity, then runs the baseline at that halved count
// plus one low and one high pass per distinct stat, each also at the
// halved count.
//
// 10,000 iterations, weighing 2 stats (crit is both a stat and the
// reference, so it costs one pass pair, not two):
//
//	half        = 10,000 / 2 = 5,000
//	baseline    = 1 pass  x 5,000 = 5,000
//	crit        = 2 passes x 5,000 = 10,000
//	strength    = 2 passes x 5,000 = 10,000
//	total       = 5,000 + 10,000 + 10,000 = 25,000
func TestWeightsIterationsMatchesTheEnginesArithmetic(t *testing.T) {
	req := runReq()
	req.Iterations = 10000
	req.Weights = &WeightsSpec{Stats: []string{"strength", "crit"}, Reference: "crit"}
	if got, want := WeightsIterations(req), 25000; got != want {
		t.Fatalf("WeightsIterations = %d, want %d", got, want)
	}
}

func TestWeightsIterationsScalesWithIterationsAndStatCount(t *testing.T) {
	for _, c := range []struct {
		name       string
		iterations int
		stats      []string
		want       int
	}{
		{"one stat (the reference alone)", 3000, []string{"crit"}, 1500 * 3},
		{"four stats", 3000, []string{"strength", "agility", "crit", "hit"}, 1500 * 9},
		{"the smallest offered iteration count", 500, []string{"strength", "crit"}, 250 * 5},
		{"no weights block", 10000, nil, 0},
	} {
		t.Run(c.name, func(t *testing.T) {
			req := runReq()
			req.Iterations = c.iterations
			if c.stats != nil {
				req.Weights = &WeightsSpec{Stats: c.stats, Reference: c.stats[0]}
			}
			if got := WeightsIterations(req); got != c.want {
				t.Errorf("WeightsIterations = %d, want %d", got, c.want)
			}
		})
	}
}
