package main

// The engine glue: turning an api.SimRequest into a DPS number or a
// set of stat weights, in-process, the same way sim/cmd/forever-sim's
// execute() and executeWeights() do (that package is `package main`
// and cannot be imported, so the pipeline - request.BuildWith/
// BuildWeights, simdb.Attach/AttachWeights, the engine's own run
// entry point, simdrain.ToResult, adapter.DPS/adapter.Weights - is
// repeated here rather than reused; see this lane's report for why
// shelling out to a built forever-sim binary was rejected in favour
// of this in-process path).
//
// The command never uses the sample-iteration entry point
// (core.RunRaidSimAsync): nothing here reads a cast log, only the
// mean DPS a run reports, so every sim uses the concurrent entry
// point (core.RunRaidSimConcurrentAsync), the same fast path
// sim/bulk's own stage requests take (NoSampleIteration is always
// set for the same reason - see sim/request.BuildWeights's own
// comment on it).

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

var registerEngineOnce sync.Once

// registerEngine registers every spec's agent factory exactly once
// per process. The engine's own guard is an unsynchronised package
// bool (sim/cmd/forever-sim's own registerOnce comment), so this must
// run before the first request is built, not per-request.
func registerEngine() { registerEngineOnce.Do(engine.RegisterAll) }

var runCounter atomic.Int64

func nextRunID() string {
	return fmt.Sprintf("leveling-bis-%d-%d", os.Getpid(), runCounter.Add(1))
}

// runPlainDPS runs req as an ordinary (non-bulk, non-weights) sim and
// returns its mean DPS. req.Iterations is used as given - this
// command sets it explicitly per call site rather than relying on
// api.SimRequest.Validate's closed set (request.Options.OpenIterations
// below is what lets an arbitrary count through).
func runPlainDPS(req api.SimRequest) (float64, error) {
	registerEngine()
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		return 0, fmt.Errorf("building the request: %w", err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		return 0, fmt.Errorf("attaching the item database: %w", err)
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if res == nil {
		return 0, errors.New("the engine produced no result")
	}
	if err := adapter.ResultError(res); err != nil {
		return 0, fmt.Errorf("the sim failed: %w", err)
	}
	return adapter.DPS(res).Mean, nil
}

// runWeights runs req (which must carry a Weights block) and returns
// the normalised stat weights, keyed by stat id.
func runWeights(req api.SimRequest) (map[string]api.StatWeight, error) {
	registerEngine()
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return nil, fmt.Errorf("building the weights request: %w", err)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return nil, fmt.Errorf("attaching the item database: %w", err)
	}
	res := core.StatWeights(engineReq)
	weights, err := adapter.Weights(res, req)
	if err != nil {
		return nil, fmt.Errorf("reading the engine's weights: %w", err)
	}
	out := make(map[string]api.StatWeight, len(weights))
	for _, w := range weights {
		out[w.Stat] = w
	}
	return out, nil
}
