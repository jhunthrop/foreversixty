package inproc

import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/score"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// TankRun runs req as a tank fight (sim/request puts a tank spec in
// front of the curated boss and healers) at exactly req.Iterations and
// reads the tank's metrics back from the engine.
func TankRun(req api.SimRequest) (score.TankRunResult, error) {
	Register()
	engineReq, err := request.BuildWith(req, buildOptions)
	if err != nil {
		return score.TankRunResult{}, fmt.Errorf("building the request: %w", err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		return score.TankRunResult{}, fmt.Errorf("attaching the item database: %w", err)
	}
	health, err := score.MaxHealth(engineReq.GetRaid(), engineReq.GetEncounter())
	if err != nil {
		return score.TankRunResult{}, err
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return score.TankRunResult{}, errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return score.TankRunResult{}, fmt.Errorf("the sim failed: %w", err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		return score.TankRunResult{}, fmt.Errorf("reading the player's metrics: %w", err)
	}
	return score.TankRunOf(player, res.IterationsDone, health), nil
}
