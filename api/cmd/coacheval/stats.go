package main

import (
	"math"
	"slices"
	"sort"
)

// ModelStats is the table row for one model over the fights that reached it.
type ModelStats struct {
	Model                       string
	Fights, Skipped             int
	PassFirst, PassRetry        int
	Fallback, ModelErrors       int
	MeanIn, MeanOut             float64
	P90In, P90Out               float64
	TotalCostUSD, MeanCostUSD   float64
	CostEstimated               bool
	MeanLatencyMS, P90LatencyMS float64
	ProblemKinds                map[string]int
}

// Passed is the fights whose answer validated, first try or after the retry.
func (s ModelStats) Passed() int { return s.PassFirst + s.PassRetry }

// PassRate is Passed over the fights that reached the model; zero for none.
func (s ModelStats) PassRate() float64 { return ratio(s.Passed(), s.Fights) }

// FirstTryRate is PassFirst over the same.
func (s ModelStats) FirstTryRate() float64 { return ratio(s.PassFirst, s.Fights) }

// FallbackRate counts both validation fallbacks and model errors: what the player gets is the rules.
func (s ModelStats) FallbackRate() float64 { return ratio(s.Fallback+s.ModelErrors, s.Fights) }

func ratio(n, d int) float64 {
	if d == 0 {
		return 0
	}
	return float64(n) / float64(d)
}

// statsByModel aggregates results per model, in the order models first appear.
func statsByModel(results []Result) []ModelStats {
	var order []string
	groups := map[string][]Result{}
	for _, r := range results {
		if _, ok := groups[r.Model]; !ok {
			order = append(order, r.Model)
		}
		groups[r.Model] = append(groups[r.Model], r)
	}
	out := make([]ModelStats, 0, len(order))
	for _, m := range order {
		out = append(out, statsOf(m, groups[m]))
	}
	return out
}

func statsOf(model string, rs []Result) ModelStats {
	s := ModelStats{Model: model, ProblemKinds: map[string]int{}}
	var in, out, lat, cost []float64
	for _, r := range rs {
		if r.Outcome == OutcomeSkipped {
			s.Skipped++
			continue
		}
		s.Fights++
		countOutcome(&s, r)
		in, out = append(in, float64(r.TokensIn)), append(out, float64(r.TokensOut))
		lat, cost = append(lat, float64(r.LatencyMS)), append(cost, r.CostUSD)
		s.CostEstimated = s.CostEstimated || r.CostEstimated
		for _, p := range r.Problems {
			s.ProblemKinds[p.Kind]++
		}
	}
	s.MeanIn, s.MeanOut, s.P90In, s.P90Out = mean(in), mean(out), p90(in), p90(out)
	s.MeanLatencyMS, s.P90LatencyMS = mean(lat), p90(lat)
	s.TotalCostUSD, s.MeanCostUSD = sum(cost), mean(cost)
	return s
}

func countOutcome(s *ModelStats, r Result) {
	switch r.Outcome {
	case OutcomePassFirst:
		s.PassFirst++
	case OutcomePassRetry:
		s.PassRetry++
	case OutcomeFallback:
		s.Fallback++
	case OutcomeModelError:
		s.ModelErrors++
	}
}

func sum(xs []float64) float64 {
	t := 0.0
	for _, x := range xs {
		t += x
	}
	return t
}

func mean(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	return sum(xs) / float64(len(xs))
}

// p90 is the nearest-rank 90th percentile.
func p90(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	s := slices.Clone(xs)
	sort.Float64s(s)
	return s[int(math.Ceil(0.9*float64(len(s))))-1]
}
