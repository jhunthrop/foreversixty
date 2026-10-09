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

// HealingRun runs req (a healer) against the profile's fake raid and
// returns what the healer did. The profile sets the fight, so the request's
// own encounter length is replaced.
func HealingRun(req api.SimRequest, profile request.HealProfile) (score.HealingResult, error) {
	Register()
	engineReq, err := request.BuildWith(req, buildOptions)
	if err != nil {
		return score.HealingResult{}, fmt.Errorf("building the request: %w", err)
	}
	profile.Attach(engineReq)
	if err := simdb.Attach(engineReq); err != nil {
		return score.HealingResult{}, fmt.Errorf("attaching the item database: %w", err)
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return score.HealingResult{}, errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return score.HealingResult{}, fmt.Errorf("the sim failed: %w", err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		return score.HealingResult{}, fmt.Errorf("reading the healer's metrics: %w", err)
	}
	return score.HealingResultOf(player, res.IterationsDone, engineReq.Encounter.Duration), nil
}
