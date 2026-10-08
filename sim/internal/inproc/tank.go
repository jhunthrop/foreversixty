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
	"github.com/wowsims/classic/sim/core/stats"
)

// Estimate is one measured quantity: its mean over the run's iterations
// and the standard error of that mean.
type Estimate struct {
	Mean  float64
	Error float64
}

// TankRunResult is what a tank fight reports: what the tank took, how
// close it came to dying, the threat it made, its own damage, and the
// hit points that all of it is measured against.
type TankRunResult struct {
	DPS           Estimate
	DTPS          Estimate
	TPS           Estimate
	TMI           Estimate
	ChanceOfDeath float64
	// Health is the character's maximum health, from the engine's own
	// stat computation for the same request.
	Health float64
}

// TankRun runs req as a tank fight (sim/request puts a tank spec in
// front of the curated boss and healers) at exactly req.Iterations and
// reads the tank's metrics back from the engine.
func TankRun(req api.SimRequest) (TankRunResult, error) {
	Register()
	engineReq, err := request.BuildWith(req, buildOptions)
	if err != nil {
		return TankRunResult{}, fmt.Errorf("building the request: %w", err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		return TankRunResult{}, fmt.Errorf("attaching the item database: %w", err)
	}
	health, err := MaxHealth(engineReq.GetRaid(), engineReq.GetEncounter())
	if err != nil {
		return TankRunResult{}, err
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return TankRunResult{}, errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return TankRunResult{}, fmt.Errorf("the sim failed: %w", err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		return TankRunResult{}, fmt.Errorf("reading the player's metrics: %w", err)
	}
	n := float64(res.IterationsDone)
	return TankRunResult{
		DPS:           tankEstimateOf(player.GetDps(), n),
		DTPS:          tankEstimateOf(player.GetDtps(), n),
		TPS:           tankEstimateOf(player.GetThreat(), n),
		TMI:           tankEstimateOf(player.GetTmi(), n),
		ChanceOfDeath: player.GetChanceOfDeath(),
		Health:        health,
	}, nil
}

func tankEstimateOf(d *proto.DistributionMetrics, iterations float64) Estimate {
	if d == nil {
		return Estimate{}
	}
	e := Estimate{Mean: d.Avg}
	if iterations > 0 {
		e.Error = d.Stdev / math.Sqrt(iterations)
	}
	return e
}

// MaxHealth is the first player's maximum health as the engine computes
// it for raid: gear, talents, buffs and consumables included.
func MaxHealth(raid *proto.Raid, encounter *proto.Encounter) (float64, error) {
	result := core.ComputeStats(&proto.ComputeStatsRequest{Raid: raid, Encounter: encounter})
	for _, party := range result.GetRaidStats().GetParties() {
		for _, player := range party.GetPlayers() {
			final := player.GetFinalStats().GetStats()
			if int(stats.Health) < len(final) && final[stats.Health] > 0 {
				return final[stats.Health], nil
			}
		}
	}
	return 0, errors.New("the engine computed no maximum health for the character")
}
