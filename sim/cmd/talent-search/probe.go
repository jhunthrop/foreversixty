package main

import (
	"fmt"
	"math"
	"sort"
)

// estimate is one sim's mean DPS and its standard error.
type estimate struct {
	Mean, Err float64
}

// dpsFunc sims one build at a given iteration count. Every call in one
// search uses the same seed, so two builds' difference is a paired
// comparison (common random numbers).
type dpsFunc func(b build, iterations int) (estimate, error)

// combinedErr is the error of a difference of two independent
// estimates - the bar a delta has to clear to count. Paired seeds make
// the true error of the difference smaller than this, so it is the
// conservative side.
func combinedErr(a, b estimate) float64 { return math.Hypot(a.Err, b.Err) }

// credit is what one talent is worth in the base build's context, as
// the engine measured it: removing it from the base if the base takes
// it, adding it at max rank otherwise (legality is not checked - the
// engine applies whatever ranks it is given, which is what makes a
// probe of a tier-6 talent possible from any build).
type credit struct {
	PerPoint float64 // DPS per point
	Diff     float64 // the probe's whole DPS difference
	Err      float64 // the difference's combined error
	// Damage is the engine crediting the talent with DPS: Diff above
	// minDamageShare of the base build's DPS. Every probe shares the
	// base run's seed, so a talent the engine gives no damage
	// reproduces the base run exactly (Diff 0); Err, the unpaired
	// combined error, is far wider than a paired probe's real noise
	// and would misfile small real gains (hit, a few % strength) as
	// nothing. Classification only steers which candidates are built;
	// every verdict is decided by the evaluation's own error bars.
	Damage bool
}

// minDamageShare is the smallest probe difference, as a share of the
// base build's DPS, that counts as the engine crediting damage.
const minDamageShare = 0.001

// probeCredits measures every talent's credit around base.
func probeCredits(t talentTrees, base build, run dpsFunc, iterations int) (map[int]credit, error) {
	baseEst, err := run(base, iterations)
	if err != nil {
		return nil, fmt.Errorf("probing the base build: %w", err)
	}
	ids := make([]int, 0, len(t.byID))
	for id := range t.byID {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	out := make(map[int]credit, len(ids))
	for _, id := range ids {
		c, err := probeOne(t, base, baseEst, id, run, iterations)
		if err != nil {
			return nil, fmt.Errorf("probing %s: %w", t.byID[id].Name, err)
		}
		out[id] = c
	}
	return out, nil
}

func probeOne(t talentTrees, base build, baseEst estimate, id int, run dpsFunc, iterations int) (credit, error) {
	ranks := base[id]
	probe := base.with(id, -ranks)
	sign := 1.0
	if ranks == 0 {
		ranks = t.byID[id].MaxRank
		probe = base.with(id, ranks)
		sign = -1
	}
	est, err := run(probe, iterations)
	if err != nil {
		return credit{}, err
	}
	diff := sign * (baseEst.Mean - est.Mean)
	e := combinedErr(baseEst, est)
	return credit{PerPoint: diff / float64(ranks), Diff: diff, Err: e, Damage: diff > minDamageShare*baseEst.Mean}, nil
}

// estimateGain is the credit-table estimate of b over base: the sum of
// per-point credits of the points that moved.
func estimateGain(base, b build, credits map[int]credit) float64 {
	g := 0.0
	for id, r := range b {
		g += float64(r-base[id]) * credits[id].PerPoint
	}
	for id, r := range base {
		if _, ok := b[id]; !ok {
			g -= float64(r) * credits[id].PerPoint
		}
	}
	return g
}
