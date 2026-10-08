package inproc

import (
	"errors"
	"fmt"
	"math"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// HealingResult is what a healer did against a heal profile's fake raid.
type HealingResult struct {
	// Effective is healing that landed on a health bar (or soaked damage,
	// for a shield) per second; Raw includes the overheal.
	Effective, Raw api.Estimate
	// ManaLastsSec is when the healer first ran out of mana, averaged over
	// the iterations; an iteration that never did contributes the time it
	// would have, projected from its mana left (the engine's own
	// convention), so a value past the fight's length means mana to spare.
	ManaLastsSec float64
	// ManaLastsError is the standard error of ManaLastsSec: the spread of
	// the per-iteration time to empty over the square root of the
	// iterations. Zero when the engine reports no spread.
	ManaLastsError float64
	// ManaSpent and EffectiveHealed are totals over every iteration, so
	// their ratio is healing per mana without a per-iteration rounding.
	ManaSpent       float64
	EffectiveHealed float64
}

// HealingPerMana is effective healing per point of mana spent.
func (r HealingResult) HealingPerMana() float64 {
	if r.ManaSpent <= 0 {
		return 0
	}
	return r.EffectiveHealed / r.ManaSpent
}

// OverhealShare is the share of raw healing that overhealed, 0 to 1.
func (r HealingResult) OverhealShare() float64 {
	if r.Raw.Mean <= 0 {
		return 0
	}
	return math.Max(0, 1-r.Effective.Mean/r.Raw.Mean)
}

// HealingRun runs req (a healer) against the profile's fake raid and
// returns what the healer did. The profile sets the fight, so the request's
// own encounter length is replaced.
func HealingRun(req api.SimRequest, profile request.HealProfile) (HealingResult, error) {
	Register()
	engineReq, err := request.BuildWith(req, buildOptions)
	if err != nil {
		return HealingResult{}, fmt.Errorf("building the request: %w", err)
	}
	profile.Attach(engineReq)
	if err := simdb.Attach(engineReq); err != nil {
		return HealingResult{}, fmt.Errorf("attaching the item database: %w", err)
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return HealingResult{}, errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return HealingResult{}, fmt.Errorf("the sim failed: %w", err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		return HealingResult{}, fmt.Errorf("reading the healer's metrics: %w", err)
	}
	return healingResultOf(player, res.IterationsDone, engineReq.Encounter.Duration), nil
}

func healingResultOf(player *proto.UnitMetrics, iterations int32, durationSec float64) HealingResult {
	out := HealingResult{
		Effective:      estimateOf(player.EffectiveHps, iterations),
		Raw:            estimateOf(player.Hps, iterations),
		ManaLastsSec:   player.GetTto().GetAvg(),
		ManaLastsError: estimateOf(player.GetTto(), iterations).Error,
		ManaSpent:      manaSpent(player.Resources),
	}
	out.EffectiveHealed = out.Effective.Mean * durationSec * float64(iterations)
	return out
}

func estimateOf(d *proto.DistributionMetrics, iterations int32) api.Estimate {
	if d == nil {
		return api.Estimate{}
	}
	est := api.Estimate{Mean: d.Avg, StdDev: d.Stdev, Min: d.Min, Max: d.Max}
	if iterations > 0 {
		est.Error = d.Stdev / math.Sqrt(float64(iterations))
	}
	return est
}

// manaSpent is every point of mana the healer paid out over the run: the
// engine records a spend as a negative gain on the action that cost it.
func manaSpent(resources []*proto.ResourceMetrics) float64 {
	spent := 0.0
	for _, r := range resources {
		if r.Type == proto.ResourceType_ResourceTypeMana && r.Gain < 0 {
			spent -= r.Gain
		}
	}
	return spent
}
