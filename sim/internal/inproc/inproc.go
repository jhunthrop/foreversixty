// Package inproc runs a plain (non-weights, non-bulk) DPS sim in this
// process, the same pipeline sim/cmd/forever-sim's execute() runs:
// request.BuildWith, simdb.Attach, the engine's concurrent entry point,
// simdrain.ToResult, adapter.DPS. sim/cmd/leveling-bis and
// sim/cmd/talent-search both need it; this is the one copy.
package inproc

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// buildOptions is PlainDPS's own request.Options, named so
// PlainDPSWithRotation can build an identical request and then swap
// in a different rotation.
var buildOptions = request.Options{OpenIterations: true, NoSampleIteration: true}

var registerOnce sync.Once

// Register registers every spec's agent factory exactly once per
// process. The engine's own guard is an unsynchronised package bool,
// so this must run before the first request is built.
func Register() { registerOnce.Do(engine.RegisterAll) }

var runCounter atomic.Int64

func nextRunID() string {
	return fmt.Sprintf("inproc-%d-%d", os.Getpid(), runCounter.Add(1))
}

// PlainDPS runs req at exactly req.Iterations (OpenIterations) with no
// sample iteration and returns the mean DPS and its standard error.
func PlainDPS(req api.SimRequest) (api.Estimate, error) {
	return PlainDPSWithRotation(req, nil)
}

// PlainDPSWithRotation is PlainDPS with the built player's rotation
// replaced by rotation when it is non-nil. The request still names
// req.Spec so the player gets the right agent and default options
// (totems, shout, and the rest of applySpec's per-spec setup) - only
// the priority list changes, which is what lets sim/cmd/rotation-search
// sim a mutated APL without ever touching the curated or embedded
// rotation files on disk. A nil rotation is exactly PlainDPS.
func PlainDPSWithRotation(req api.SimRequest, rotation *proto.APLRotation) (api.Estimate, error) {
	Register()
	engineReq, err := request.BuildWith(req, buildOptions)
	if err != nil {
		return api.Estimate{}, fmt.Errorf("building the request: %w", err)
	}
	if rotation != nil {
		for _, party := range engineReq.GetRaid().GetParties() {
			for _, player := range party.GetPlayers() {
				player.Rotation = rotation
			}
		}
	}
	if err := simdb.Attach(engineReq); err != nil {
		return api.Estimate{}, fmt.Errorf("attaching the item database: %w", err)
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return api.Estimate{}, errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return api.Estimate{}, fmt.Errorf("the sim failed: %w", err)
	}
	return adapter.DPS(res), nil
}
