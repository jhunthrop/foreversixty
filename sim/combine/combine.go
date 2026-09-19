// Package combine splits a run across workers and puts the pieces back
// together. The browser's worker pool and the server lane both use it,
// so the arithmetic is written once.
package combine

import (
	"errors"
	"fmt"
	"math"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
)

// Split divides a request into n parts by iteration count.
//
// Each part's RandomSeed is offset by the iterations of every part
// before it. That is what the engine's own SplitSimRequestForConcurrency
// does, and the reason is that the engine increments its seed once per
// iteration: without the offset, four workers would run the same four
// thousand rolls and a split run would not match a serial one. A paired
// comparison depends on it.
//
// A part's iteration count is not one of api.ValidIterations - 3,000
// over four workers is 750 - so a part is built with
// request.Options{SplitPart: true}, which validates everything but the
// closed set the settings bar offers.
func Split(req api.SimRequest, n int) ([]api.SimRequest, error) {
	if n <= 0 {
		return nil, fmt.Errorf("combine: split count must be positive, got %d", n)
	}
	if req.Iterations <= 0 {
		return nil, fmt.Errorf("combine: iterations must be positive, got %d", req.Iterations)
	}
	if n > req.Iterations {
		n = req.Iterations
	}

	per := req.Iterations / n
	out := make([]api.SimRequest, n)
	seed := req.RandomSeed
	for i := 0; i < n; i++ {
		part := req
		part.Iterations = per
		if i == 0 {
			// The remainder goes to the first part, as the engine does
			// it; spreading it would lose iterations to truncation.
			part.Iterations += req.Iterations % n
		}
		part.RandomSeed = seed
		seed += int64(part.Iterations)
		out[i] = part
	}
	return out, nil
}

// Results combines partial results into one.
func Results(parts []api.SimResult) (api.SimResult, error) {
	if len(parts) == 0 {
		return api.SimResult{}, errors.New("combine: no results")
	}
	var total int
	for i, p := range parts {
		if p.Error != "" {
			return api.SimResult{}, fmt.Errorf("combine: part %d failed: %s", i, p.Error)
		}
		if p.IterationsRun <= 0 {
			return api.SimResult{}, fmt.Errorf("combine: part %d ran no iterations", i)
		}
		total += p.IterationsRun
	}

	out := parts[0]
	out.IterationsRun = total

	// Pooled mean: weight each part by the iterations behind it.
	var mean float64
	for _, p := range parts {
		mean += p.DPS.Mean * float64(p.IterationsRun)
	}
	mean /= float64(total)

	// Pooled variance is the within-part variance plus the spread
	// between the part means. Dropping the second term would report a
	// tighter error than a serial run, and the sim page would lie about
	// its own precision.
	var pooled float64
	for _, p := range parts {
		w := float64(p.IterationsRun)
		d := p.DPS.Mean - mean
		pooled += w * (p.DPS.StdDev*p.DPS.StdDev + d*d)
	}
	pooled /= float64(total)

	out.DPS = api.Estimate{
		Mean:   mean,
		StdDev: math.Sqrt(pooled),
		Error:  math.Sqrt(pooled) / math.Sqrt(float64(total)),
		Min:    parts[0].DPS.Min,
		Max:    parts[0].DPS.Max,
	}
	for _, p := range parts {
		out.DPS.Min = math.Min(out.DPS.Min, p.DPS.Min)
		out.DPS.Max = math.Max(out.DPS.Max, p.DPS.Max)
	}

	out.DurationMS = 0
	for _, p := range parts {
		if p.DurationMS > out.DurationMS {
			// Wall clock of a parallel run is the slowest part, not the
			// sum: the parts ran at the same time.
			out.DurationMS = p.DurationMS
		}
	}

	out.Summary = weightSummaries(parts, total)
	out.Request.Iterations = total
	return out, nil
}

// weightSummaries averages the per-fight summaries by iteration share.
// Each part's summary is already a per-fight average (sim/adapter divides
// by IterationsDone), so combining them is a weighted mean of like
// quantities rather than a re-sum.
func weightSummaries(parts []api.SimResult, total int) summary.Summary {
	out := parts[0].Summary
	if len(parts) == 1 {
		return out
	}
	// Damage totals and ability rows are the only fields a viewer reads
	// as a number; auras, casts and resources are shares and averages
	// that the largest part already represents within sampling error.
	// Weighting the damage table is the part worth doing exactly.
	scale := func(a *summary.Actor, w float64) {
		a.Total = int64(float64(a.Total) * w)
		a.Effective = int64(float64(a.Effective) * w)
		for i := range a.Abilities {
			a.Abilities[i].Total = int64(float64(a.Abilities[i].Total) * w)
			a.Abilities[i].Effective = int64(float64(a.Abilities[i].Effective) * w)
		}
	}
	merged := make([]summary.Actor, len(out.DamageDone))
	copy(merged, out.DamageDone)
	for i := range merged {
		scale(&merged[i], float64(parts[0].IterationsRun)/float64(total))
	}
	for _, p := range parts[1:] {
		w := float64(p.IterationsRun) / float64(total)
		for i := range merged {
			if i >= len(p.Summary.DamageDone) {
				break
			}
			src := p.Summary.DamageDone[i]
			merged[i].Total += int64(float64(src.Total) * w)
			merged[i].Effective += int64(float64(src.Effective) * w)
			for j := range merged[i].Abilities {
				if j >= len(src.Abilities) {
					break
				}
				merged[i].Abilities[j].Total += int64(float64(src.Abilities[j].Total) * w)
				merged[i].Abilities[j].Effective += int64(float64(src.Abilities[j].Effective) * w)
			}
		}
	}
	out.DamageDone = merged
	return out
}
