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
	"github.com/jhunthrop/foreversixty/sim/internal/statid"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// engineRunner is what runSpec, verifyBand and rankTrinketSlot need
// from the engine: a plain DPS run and a stat-weights run. It exists
// so those functions take an explicit dependency (this codebase's own
// rule: explicit dependencies over globals) rather than calling
// runPlainDPS/runWeights directly - a test can then inject a fake that
// returns canned results instead of spawning the real engine, which
// for a weights or verify run is the whole reason
// sim/cmd/leveling-bis's own tests could not exercise runSpec/
// verifyBand/rankTrinketSlot before this type existed. realEngine
// below is the only production implementation.
type engineRunner interface {
	RunPlainDPS(req api.SimRequest) (float64, error)
	// RunPlainDPSWithError is RunPlainDPS plus the same run's own
	// standard error (bis-ranker-integrity-11's brief, item 2) -
	// trinkets.go's own tournament and no-trinket baseline are its only
	// callers, so a trinket's real DPS gain can be checked "positive
	// beyond its own error" the same way weights.go's
	// referenceMeasurementReason already checks a band's reference
	// stat (item 1), instead of against a flat, noise-blind DPS
	// threshold.
	RunPlainDPSWithError(req api.SimRequest) (mean, stdErr float64, err error)
	// RunWeights' second return is referenceDPSPerPoint (this lane's
	// brief, item 2) - see runWeights' own doc for why it has to be
	// read out of the weights run's raw result rather than the
	// normalised map[string]api.StatWeight alone.
	RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error)
}

// realEngine is the engineRunner backed by the actual wowsims-classic
// engine (runPlainDPS/runWeights below) - the only production
// implementation; main's run() is the only place that constructs one.
type realEngine struct{}

func (realEngine) RunPlainDPS(req api.SimRequest) (float64, error) { return runPlainDPS(req) }

func (realEngine) RunPlainDPSWithError(req api.SimRequest) (float64, float64, error) {
	return runPlainDPSWithError(req)
}

func (realEngine) RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	return runWeights(req)
}

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
	est, err := runPlainDPSEstimate(req)
	if err != nil {
		return 0, err
	}
	return est.Mean, nil
}

// runPlainDPSWithError is runPlainDPS plus the same run's own standard
// error (api.Estimate.Error - adapter.DPS's own Stdev/sqrt(iterations)
// - this lane's brief, item 2: "a [trinket] pick must have... a
// tournament delta positive beyond its error", the same "positive
// beyond its own error" bar weights.go's referenceMeasurementReason
// already applies to a band's reference stat measurement (item 1).
// trinkets.go is the only caller - the per-candidate tournament loop
// and the no-trinket baseline both already run exactly this sim, so
// this reads an error the engine had already computed rather than
// spending a second sim run to get it.
func runPlainDPSWithError(req api.SimRequest) (mean, stdErr float64, err error) {
	est, err := runPlainDPSEstimate(req)
	if err != nil {
		return 0, 0, err
	}
	return est.Mean, est.Error, nil
}

// runPlainDPSEstimate is runPlainDPS/runPlainDPSWithError's shared
// engine call - one sim, read back as api.Estimate (Mean and Error
// both, so a caller needing either never duplicates the run).
func runPlainDPSEstimate(req api.SimRequest) (api.Estimate, error) {
	registerEngine()
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		return api.Estimate{}, fmt.Errorf("building the request: %w", err)
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

// runWeights runs req (which must carry a Weights block) and returns
// the normalised stat weights, keyed by stat id, plus
// referenceDPSPerPoint: the measured, UN-normalised DPS this run found
// for one point of req.Weights.Reference (this lane's brief, item 2 -
// "reference_dps_per_point... the raw weight the ratios are normalised
// by"). adapter.Weights computes exactly this number as its own
// unexported "scale" (sim/adapter/adapter.go's own doc: "dividing is
// one arithmetic, in one place") to divide every OTHER weight by, but
// never returns it - the reference stat's own entry in its output is
// always exactly 1.0 by construction (Weight/scale), so the raw figure
// would otherwise be lost the moment adapter.Weights returns.
// referenceStatRawWeight (below) reads it straight off this same res,
// duplicating adapter.go's own small lookup rather than changing that
// function's signature - this file's own package doc already commits
// to that pattern ("the pipeline... is repeated here rather than
// reused").
func runWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	registerEngine()
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return nil, 0, fmt.Errorf("building the weights request: %w", err)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return nil, 0, fmt.Errorf("attaching the item database: %w", err)
	}
	res := core.StatWeights(engineReq)
	weights, err := adapter.Weights(res, req)
	if err != nil {
		return nil, 0, fmt.Errorf("reading the engine's weights: %w", err)
	}
	out := make(map[string]api.StatWeight, len(weights))
	for _, w := range weights {
		out[w.Stat] = w
	}
	referenceDPSPerPoint, err := referenceStatRawWeight(res, req)
	if err != nil {
		return nil, 0, fmt.Errorf("reading the reference stat's raw weight: %w", err)
	}
	return out, referenceDPSPerPoint, nil
}

// referenceStatRawWeight is the engine's own raw (un-normalised) DPS
// delta for one point of req.Weights.Reference - see runWeights' own
// doc for why this is read here instead of coming back from
// adapter.Weights. Validated the same way adapter.Weights validates
// its own reference lookup (statid.Parse); a request that reached this
// point already passed adapter.Weights' own reference checks
// (ErrNoWeights on a bad or zero-weighing reference), so an error here
// in production would mean the two lookups disagree, not that the
// request was actually invalid.
func referenceStatRawWeight(res *proto.StatWeightsResult, req api.SimRequest) (float64, error) {
	if req.Weights == nil {
		return 0, fmt.Errorf("the request carries no weights block")
	}
	reference, ok := statid.Parse(req.Weights.Reference)
	if !ok {
		return 0, fmt.Errorf("reference stat %q is not a known stat id", req.Weights.Reference)
	}
	raw := res.GetDps().GetWeights().GetStats()
	if int(reference) >= len(raw) {
		return 0, nil
	}
	return raw[reference], nil
}
