// Package score turns the engine's metrics for one player into the figures
// the site and the ranker publish: a healer's effective healing per second
// and a tank's score. The nightly ranker, the in-process runner and the
// browser's wasm all call this one copy, so the same run produces the same
// number in all three.
package score

import (
	"math"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// EstimateOf lifts the engine's distribution into an api.Estimate whose
// Error is the standard error of the mean over iterations. A missing
// distribution is the zero Estimate.
func EstimateOf(d *proto.DistributionMetrics, iterations int32) api.Estimate {
	if d == nil {
		return api.Estimate{}
	}
	est := api.Estimate{Mean: d.Avg, StdDev: d.Stdev, Min: d.Min, Max: d.Max}
	if iterations > 0 {
		est.Error = d.Stdev / math.Sqrt(float64(iterations))
	}
	return est
}
