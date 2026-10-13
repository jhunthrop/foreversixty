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
	"sync/atomic"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/combine"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
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
	var aborted bool

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
		// Every step but the first skips the sample, mirroring
		// combine.Split's own part.NoSample = req.NoSample || i != 0:
		// combine.Results keeps only part zero's Sample, so a later
		// step computing one is a full extra iteration paid for and
		// then thrown away, the exact waste Split was written to
		// avoid at the shard level.
		part.NoSample = req.NoSample || len(parts) != 0

		res, err := Execute(part, progress)
		if err != nil {
			if errors.Is(err, adapter.ErrAborted) {
				if len(parts) == 0 {
					// Nothing pooled yet, so the step's own partial
					// result - not a zeroed one - is what "how far it
					// got" means here.
					return res, err
				}
				// What completed is still an answer, and the caller
				// asked for the stop - both are true, so the pooled
				// result is written AND marked below, rather than
				// picking one over the other.
				aborted = true
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
	if aborted {
		// The pooled steps are still the answer - combine.Results
		// already merged only the ones that finished - but the run
		// did not reach its target or its ceiling, and the caller
		// asked for the stop, so the result says so rather than
		// leaving it for a reader to infer from TargetError alone.
		pooled.Aborted = true
		return pooled, adapter.ErrAborted
	}
	return pooled, nil
}

// executeWeights computes stat weights natively, and the baseline DPS beside them: one
// weights run answers with both (the in-game "Sim me" request is one such run, and its
// sim_result reads DPS and error off this result).
//
// It drives the engine's own sweep step by step rather than through core.StatWeightsAsync,
// because StatWeightsAsync keeps its baseline sim to itself and answers only the weights.
// The steps are the engine's exported halves of that same call, in the same order
// (runStatWeights in the wowsims fork's sim/core/statweight.go):
//
//   - core.StatWeightRequests builds the baseline request and one low/high pair per stat,
//     with the halved iteration count and the shared seed that line their random streams up;
//   - each request runs to its result, the baseline first, on the engine's concurrent
//     entry point (the single-threaded one for a test request, as the engine itself picks);
//   - core.StatWeightCompute turns the baseline and the pairs into the weights.
//
// So the weights are the ones StatWeightsAsync would compute, and the baseline's Summary,
// DPS and error are this result's, the same fields a plain run fills. IterationsRun is the
// sweep's total across every sub-sim. Each sub-sim registers its own signal id
// (simsignals refuses a duplicate while the last one is still registered), and an interrupt
// stops whichever one is running; a stopped sub-sim stops the sweep, answered as an aborted
// result with the empty summary and no weights, the way an aborted plain run is.
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
	data := core.StatWeightRequests(engineReq)

	total := data.BaseRequest.SimOptions.Iterations
	sims := int32(1)
	for _, pair := range data.StatSimRequests {
		total += pair.RequestLow.SimOptions.Iterations + pair.RequestHigh.SimOptions.Iterations
		sims += 2
	}
	entry := core.RunRaidSimConcurrentAsync
	if data.BaseRequest.SimOptions.GetIsTest() {
		entry = core.RunRaidSimAsync
	}

	var current atomic.Value
	current.Store("")
	defer onInterruptFunc(func() string { return current.Load().(string) })()

	var enc *json.Encoder
	if progress != nil {
		enc = json.NewEncoder(progress)
	}
	// finished is the iterations of every sub-sim that has ended; running is the one under way.
	var finished, running, completedSims int32
	aborted := func() (api.SimResult, error) {
		return api.SimResult{
			EngineVersion: enginever.Version,
			Request:       req,
			Lane:          api.LaneServer,
			Aborted:       true,
			IterationsRun: int(finished + running),
			DurationMS:    time.Since(start).Milliseconds(),
			Summary:       adapter.EmptySummary(),
		}, adapter.ErrAborted
	}
	// run is one sub-sim of the sweep, to its result. A weights run is many sims, so the sim
	// counts are the honest progress and the iteration count alone would restart per stat:
	// they ride the same three fields a bulk stage uses.
	run := func(r *proto.RaidSimRequest) (*proto.RaidSimResult, error) {
		id := runID()
		current.Store(id)
		reporter := make(chan *proto.ProgressMetrics, 32)
		entry(r, reporter, id)
		running = 0
		res := simdrain.ToResult(reporter, func(p *proto.ProgressMetrics) {
			running = p.CompletedIterations
			if enc != nil {
				_ = enc.Encode(struct {
					Completed   int32   `json:"completed"`
					Total       int32   `json:"total"`
					DPS         float64 `json:"dps"`
					CombosDone  int32   `json:"combos_done"`
					CombosTotal int32   `json:"combos_total"`
				}{finished + running, total, p.Dps, completedSims, sims})
			}
		})
		if res == nil {
			return nil, errors.New("the engine produced no result")
		}
		completedSims++
		finished += res.IterationsDone
		running = 0
		// Switching on the type, not on the message: an abort's ErrorOutcome carries none.
		if err := adapter.ResultError(res); err != nil {
			if errors.Is(err, adapter.ErrAborted) {
				return nil, err
			}
			return nil, fmt.Errorf("the sim failed: %w", err)
		}
		return res, nil
	}

	base, err := run(data.BaseRequest)
	if err != nil {
		if errors.Is(err, adapter.ErrAborted) {
			return aborted()
		}
		return api.SimResult{}, err
	}
	pairs := make([]*proto.StatWeightsStatResultData, 0, len(data.StatSimRequests))
	for _, pair := range data.StatSimRequests {
		low, err := run(pair.RequestLow)
		if err == nil {
			var high *proto.RaidSimResult
			if high, err = run(pair.RequestHigh); err == nil {
				pairs = append(pairs, &proto.StatWeightsStatResultData{
					StatData: pair.StatData, ResultLow: low, ResultHigh: high,
				})
				continue
			}
		}
		if errors.Is(err, adapter.ErrAborted) {
			return aborted()
		}
		return api.SimResult{}, err
	}
	engineRes := core.StatWeightCompute(&proto.StatWeightsCalcRequest{
		BaseResult: base, EpReferenceStat: data.EpReferenceStat, StatSimResults: pairs,
	})
	if engineRes.GetError().GetType() == proto.ErrorOutcomeType_ErrorOutcomeAborted {
		return aborted()
	}
	weights, err := adapter.Weights(engineRes, req)
	if err != nil {
		return api.SimResult{}, err
	}
	sum, err := adapter.Summarize(base, req)
	if err != nil {
		return api.SimResult{}, fmt.Errorf("adapting the baseline: %w", err)
	}
	return api.SimResult{
		EngineVersion: enginever.Version,
		Request:       req,
		Lane:          api.LaneServer,
		IterationsRun: int(finished),
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       sum,
		DPS:           adapter.DPS(base),
		Weights:       weights,
	}, nil
}
