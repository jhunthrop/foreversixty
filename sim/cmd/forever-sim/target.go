package main

// Smart Sim and stat weights, natively.
//
// The stepping decision is api.NeedsMoreIterations, which the browser
// asks too - the loop is written twice because the two lanes drive it
// differently (a worker pool there, a for loop here), but the QUESTION
// is asked once, in Go, so the two stop at the same precision.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/combine"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// ExecuteToTarget is Execute, plus the target-error loop.
//
// A request with no target is one run, unchanged: this is the default
// path and it must stay exactly today's behaviour.
//
// A request with one runs steps of api.StepIterations and pools them
// with combine.Results - the same pooling the browser's worker pool
// uses for a sharded run, so a stepped run's error bar is computed the
// same way whichever lane produced it. Each step's seed is offset by
// the iterations before it, as combine.Split does, so the stream of
// a stepped run matches a single run of the same total.
func ExecuteToTarget(req api.SimRequest, progress io.Writer) (api.SimResult, error) {
	if req.TargetError <= 0 {
		return Execute(req, progress)
	}
	start := time.Now()
	var parts []api.SimResult
	pooled := api.SimResult{Request: req}
	seed := req.RandomSeed

	for {
		step := api.NextStepIterations(pooled, req)
		if step == 0 {
			break
		}
		part := req
		part.TargetError = 0
		part.Iterations = step
		part.RandomSeed = seed
		seed += int64(step)

		res, err := Execute(part, progress)
		if err != nil {
			if errors.Is(err, adapter.ErrAborted) && len(parts) > 0 {
				// What completed is still an answer, and the caller
				// asked for the stop.
				break
			}
			return api.SimResult{}, err
		}
		parts = append(parts, res)

		pooled, err = combine.Results(parts)
		if err != nil {
			return api.SimResult{}, fmt.Errorf("pooling the steps: %w", err)
		}
		// The pooled result's request is part zero's, which carries no
		// target; the run's own request is the one with the target,
		// and it is what the loop and the stored row must both see.
		pooled.Request = req
	}
	if len(parts) == 0 {
		return api.SimResult{}, errors.New("a target-error run made no steps; iterations must be at least one step")
	}
	pooled.EngineVersion = enginever.Version
	pooled.Lane = api.LaneServer
	pooled.DurationMS = time.Since(start).Milliseconds()
	return pooled, nil
}

// executeWeights computes stat weights natively, over the engine's own
// StatWeightsAsync - the same call the browser makes.
//
// This does NOT go through sim/internal/simdrain.ToResult, unlike
// execute. That helper drains to the channel's close because a
// SUCCESSFUL raid sim closes progress on its way out; core.StatWeightsAsync
// has no such path - on every outcome it sends exactly one
// ProgressMetrics.FinalWeightResult and returns without ever closing the
// channel. Draining to close here would hang on every successful weights
// run. Taking the first FinalWeightResult and stopping, as
// sim/cmd/wasm's runWeights does, is therefore correct rather than a
// shortcut.
func executeWeights(req api.SimRequest, progress io.Writer) (api.SimResult, error) {
	registerOnce.Do(engine.RegisterAll)
	start := time.Now()

	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return api.SimResult{}, fmt.Errorf("%w: %v", errBadInput, err)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return api.SimResult{}, err
	}

	reporter := make(chan *proto.ProgressMetrics, 32)
	id := runID()
	defer onInterrupt(id)()
	core.StatWeightsAsync(engineReq, reporter, id)

	var enc *json.Encoder
	if progress != nil {
		enc = json.NewEncoder(progress)
	}
	var engineRes *proto.StatWeightsResult
	// The engine's own running total across the WHOLE sweep, not the
	// per-sim count: iterations_run must report the sweep's total
	// (2*len(stats)+1 sub-sims), never the per-sim req.Iterations.
	var iterationsRun int
	for p := range reporter {
		if p.FinalWeightResult != nil {
			engineRes = p.FinalWeightResult
			break
		}
		iterationsRun = int(p.CompletedIterations)
		if enc != nil {
			// A weights run is many sims, so the sim counts are the
			// honest progress and the iteration count alone would
			// restart per stat. They ride the same three fields a
			// bulk stage uses.
			_ = enc.Encode(struct {
				Completed   int32   `json:"completed"`
				Total       int32   `json:"total"`
				DPS         float64 `json:"dps"`
				CombosDone  int32   `json:"combos_done"`
				CombosTotal int32   `json:"combos_total"`
			}{p.CompletedIterations, p.TotalIterations, p.Dps, p.CompletedSims, p.TotalSims})
		}
	}
	if engineRes == nil {
		return api.SimResult{}, errors.New("the engine produced no weights")
	}
	if iterationsRun == 0 {
		iterationsRun = req.Iterations
	}

	// Switching on the type, not on the message: an aborted weights run's
	// ErrorOutcome carries none, so a message check would misreport Stop
	// as a corrupt result (adapter.Weights only branches on a non-empty
	// message). This mirrors execute's own adapter.ResultError type
	// check for a plain run; StatWeightsResult has no ResultError of its
	// own to call. See sim/cmd/wasm's weightsResult, which this follows.
	if engineRes.GetError().GetType() == proto.ErrorOutcomeType_ErrorOutcomeAborted {
		return api.SimResult{
			EngineVersion: enginever.Version,
			Request:       req,
			Lane:          api.LaneServer,
			Aborted:       true,
			IterationsRun: iterationsRun,
			DurationMS:    time.Since(start).Milliseconds(),
			Summary:       adapter.EmptySummary(),
		}, nil
	}

	weights, err := adapter.Weights(engineRes, req)
	if err != nil {
		return api.SimResult{}, err
	}
	return api.SimResult{
		EngineVersion: enginever.Version,
		Request:       req,
		Lane:          api.LaneServer,
		IterationsRun: iterationsRun,
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       adapter.EmptySummary(),
		Weights:       weights,
	}, nil
}
