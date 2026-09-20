//go:build js && wasm

// Command wasm is the browser half of the sim. It exports ten
// functions - simRun, simSplit, simCombine, simAbort, simPlan,
// simRank, simCount, simNeedsMore, simValidate and simWeights - all
// taking and returning JSON strings, plus one string global,
// simEngineVersion, which is the engine sha this module was built
// from. The page reads that global rather than being told the sha by
// the server, so a request can never name an engine the wasm it is
// running in is not.
//
// simSplit and simCombine are for plain runs only: a bulk stage's
// requests run whole and unsplit, one worker-pool slot per request, so
// that "abort returns what finished" stays true and no request's
// statistics are computed twice.
//
// It exists in this repository rather than in the engine because
// sim/request and sim/adapter are linked in here: the browser gets a
// finished SimResult with its summary.Summary already built by the same
// Go code the server runs. The engine's own thirteen js.Global().Set
// entrypoints are an implementation detail behind these ten and the web
// must not call them.
//
// The active build's item database is embedded through
// sim/internal/simdb rather than taken from the engine's --tags=with_db
// table, which is vanilla's and which Forever re-itemises out from
// under. It costs 0.13 MB gzipped, so the browser can equip a real
// Forever gear set inside the 4 MB budget.
package main

import (
	"encoding/json"
	"errors"
	"syscall/js"
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
	"github.com/wowsims/classic/sim/core/simsignals"
)

func main() {
	engine.RegisterAll()
	core.SetRunningInWasm()

	js.Global().Set("simRun", js.FuncOf(simRun))
	js.Global().Set("simSplit", js.FuncOf(simSplit))
	js.Global().Set("simCombine", js.FuncOf(simCombine))
	js.Global().Set("simAbort", js.FuncOf(simAbort))
	js.Global().Set("simPlan", js.FuncOf(simPlan))
	js.Global().Set("simRank", js.FuncOf(simRank))
	js.Global().Set("simCount", js.FuncOf(simCount))
	js.Global().Set("simNeedsMore", js.FuncOf(simNeedsMore))
	js.Global().Set("simValidate", js.FuncOf(simValidate))
	js.Global().Set("simWeights", js.FuncOf(simWeights))
	// The pin is compiled in, not injected: a plain `go build ./cmd/wasm`
	// used to produce "dev" and stamp it on real rows.
	js.Global().Set("simEngineVersion", js.ValueOf(enginever.Version))

	// The host page defines wasmready and is told the moment every
	// export exists, so it never races them.
	js.Global().Call("wasmready")
	select {}
}

// simRun(requestJSON, callbackId) runs one request to completion and
// returns SimResult JSON. Progress is reported by calling the global
// simProgress(callbackId, progressJSON), where progressJSON is an
// api.Progress - Pick<SimResult, 'iterations_run' | 'dps'> - so the page
// reads a partial result with the accessors it already has. JSON in,
// JSON out, like everything else here.
func simRun(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return failJSON(api.SimRequest{}, "simRun takes (requestJSON, callbackId)")
	}
	req, err := decodeRequest(args[0].String())
	if err != nil {
		return failJSON(req, "the request is not valid JSON: "+err.Error())
	}
	callbackID := args[1].String()

	// Every request the browser runs is a part: the worker pool calls
	// simSplit first, even for one worker, and a part's iteration count
	// is not one of api.ValidIterations by construction - 3,000 over
	// four workers is 750. OpenIterations relaxes that check and
	// nothing else. The whole request was validated before it was split,
	// by the page and by the api lane.
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: req.NoSample})
	if err != nil {
		return failJSON(req, err.Error())
	}
	// Forever's own item rows, from the active build, embedded at build
	// time: the browser has no protobuf and cannot send them.
	if err := simdb.Attach(engineReq); err != nil {
		return failJSON(req, err.Error())
	}

	start := time.Now()
	reporter := make(chan *proto.ProgressMetrics, 32)
	// Threading does not work in wasm, so this is the serial entrypoint.
	// Parallelism comes from the worker pool: the page calls simSplit and
	// gives each worker one part.
	core.RunRaidSimAsync(engineReq, reporter, callbackID)

	// The drain loop is sim/internal/simdrain's, shared with
	// sim/cmd/forever-sim and tested there: draining to the channel's
	// close (rather than breaking the instant a FinalRaidResult
	// arrives) is what makes a sample request's SampleIteration
	// mutation guaranteed-visible, and stopping on an error result is
	// what keeps a worker from wedging on a close that never comes.
	// ToResult carries the whole argument; this lane supplies only the
	// progress sink.
	engineRes := simdrain.ToResult(reporter, func(p *proto.ProgressMetrics) {
		cb := js.Global().Get("simProgress")
		if cb.Type() != js.TypeFunction {
			return
		}
		b, err := json.Marshal(api.Progress{
			IterationsRun: int(p.CompletedIterations),
			DPS:           api.Estimate{Mean: p.Dps},
		})
		if err != nil {
			return
		}
		cb.Invoke(callbackID, string(b))
	})
	if engineRes == nil {
		return failJSON(req, "the engine produced no result")
	}
	// An abort's ErrorOutcome carries no message, so this switches on
	// the type. Testing the message alone let Stop through as a result
	// with zero iterations, which the adapter then called corrupt.
	if err := adapter.ResultError(engineRes); err != nil {
		if errors.Is(err, adapter.ErrAborted) {
			// Not a failure - the user pressed Stop - so it carries no
			// error message and the page renders it as a run that
			// ended early rather than as something that went wrong.
			return encodeOrError(stamp(api.SimResult{
				Request:       req,
				Aborted:       true,
				IterationsRun: int(engineRes.IterationsDone),
				Summary:       adapter.EmptySummary(),
			}))
		}
		return failJSON(req, err.Error())
	}

	sum, err := adapter.Summarize(engineRes, req)
	if err != nil {
		return failJSON(req, err.Error())
	}
	return encodeOrError(stamp(api.SimResult{
		Request:       req,
		DPS:           adapter.DPS(engineRes),
		IterationsRun: int(engineRes.IterationsDone),
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       sum,
		Sample:        adapter.Sample(engineRes),
	}))
}

// simSplit(requestJSON, n) returns a JSON array of n request JSONs, one
// per worker, with the seeds already offset.
func simSplit(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return errorJSON("simSplit takes (requestJSON, n)")
	}
	req, err := decodeRequest(args[0].String())
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	parts, err := combine.Split(req, args[1].Int())
	if err != nil {
		return errorJSON(err.Error())
	}
	b, err := json.Marshal(parts)
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}

// simCombine(resultsJSON) folds a JSON array of partial SimResults into
// one SimResult JSON.
func simCombine(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorJSON("simCombine takes (resultsJSON)")
	}
	var parts []api.SimResult
	if err := json.Unmarshal([]byte(args[0].String()), &parts); err != nil {
		return errorJSON("the results are not valid JSON: " + err.Error())
	}
	out, err := combine.Results(parts)
	if err != nil {
		return errorJSON(err.Error())
	}
	b, err := json.Marshal(out)
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}

// simAbort(callbackId) stops a run started with the same id and returns
// {"aborted": true|false}, where false means no run is registered under
// that id. It returns JSON rather than a bare boolean so that a wrong
// call - which used to come back as the same `false` as "no such run" -
// is distinguishable, and so that every export has one shape.
func simAbort(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorJSON("simAbort takes (callbackId)")
	}
	b, err := json.Marshal(struct {
		Aborted bool `json:"aborted"`
	}{simsignals.AbortById(args[0].String())})
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}

// simPlan(requestJSON) returns the first stage of a bulk run:
// {"stage":1,"iterations":100,"requests":[...],"combos":[...]}.
func simPlan(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorJSON("simPlan takes (requestJSON)")
	}
	return planJSON(args[0].String())
}

// simRank(requestJSON, stageJSON, resultsJSON) scores a finished stage
// and returns {"next": stage} or {"result": SimResult}.
func simRank(_ js.Value, args []js.Value) any {
	if len(args) < 3 {
		return errorJSON("simRank takes (requestJSON, stageJSON, resultsJSON)")
	}
	return rankJSON(args[0].String(), args[1].String(), args[2].String())
}

// simCount(requestJSON) returns {"combinations": n} without building
// a single request, so the page can show the count on every tick.
func simCount(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorJSON("simCount takes (requestJSON)")
	}
	return countJSON(args[0].String())
}

// simNeedsMore(resultJSON, requestJSON) returns {"needs_more": bool}:
// the Smart Sim decision, asked of Go so both lanes stop at the same
// precision.
func simNeedsMore(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return errorJSON("simNeedsMore takes (resultJSON, requestJSON)")
	}
	return needsMoreJSON(args[0].String(), args[1].String())
}

// simValidate(requestJSON) returns {"ok": bool, "errors": [...]}: the
// same Validate the run applies, per field, for the request drawer. A
// bad request body still comes back this way (see validateJSON); only
// this wrapper's own arity check below returns the bare
// {"error": "..."} shape, for a caller bug rather than a member's
// request.
func simValidate(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorJSON("simValidate takes (requestJSON)")
	}
	return validateJSON(args[0].String())
}

func init() {
	weightsRunner = runWeights
}

// runWeights is the engine half of simWeights.
//
// The engine runs the same character twice per stat, a little above
// and a little below, so this is a raid sim request with a stat list
// - which is exactly what request.BuildWeights builds. Progress is
// the engine's own per-sim ticks, reported through the simProgress
// global a plain run already uses, so the page's progress handling is
// unchanged.
//
// This does NOT go through sim/internal/simdrain.ToResult, unlike
// simRun. That helper drains to the channel's close because a
// SUCCESSFUL raid sim closes progress on its way out (core/sim.go) and
// the sample-iteration mutation after the final message is only
// visible once that close is observed. core.StatWeightsAsync has no
// such path: on every outcome - success, sim error, or a failed
// signals registration - it sends exactly one
// ProgressMetrics.FinalWeightResult and returns without ever closing
// the channel (core/api.go, core/statweight.go). Draining to close
// here would hang forever on the common case, not just the two rare
// paths ToResult's own doc comment carves out. Taking the first
// FinalWeightResult and stopping is therefore correct, not a shortcut
// - and whatever res.Error turns out to mean (an abort, or a real
// failure) is weightsResult's decision below, not this loop's; nothing
// here needs to inspect it first.
func runWeights(req api.SimRequest, callbackID string) (api.SimResult, error) {
	start := time.Now()
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return api.SimResult{}, err
	}
	// Forever's own item rows: a weights run equips the character the
	// same way a DPS run does.
	if err := simdb.AttachWeights(engineReq); err != nil {
		return api.SimResult{}, err
	}

	reporter := make(chan *proto.ProgressMetrics, 32)
	core.StatWeightsAsync(engineReq, reporter, callbackID)

	var engineRes *proto.StatWeightsResult
	// The engine's own running total across the WHOLE sweep, not the
	// per-sim count: see weightsResult's doc comment on why
	// req.Iterations is the wrong number to report.
	var iterationsRun int
	for p := range reporter {
		if p.FinalWeightResult != nil {
			engineRes = p.FinalWeightResult
			break
		}
		iterationsRun = int(p.CompletedIterations)
		if cb := js.Global().Get("simProgress"); cb.Type() == js.TypeFunction {
			if b, err := json.Marshal(api.Progress{
				IterationsRun: int(p.CompletedIterations),
				DPS:           api.Estimate{Mean: p.Dps},
				// A weights run is many sims, so the page shows the
				// same "n of m" line a bulk stage does rather than a
				// bare iteration count that restarts per stat.
				CombosDone:  int(p.CompletedSims),
				CombosTotal: int(p.TotalSims),
			}); err == nil {
				cb.Invoke(callbackID, string(b))
			}
		}
	}
	if engineRes == nil {
		return api.SimResult{}, errors.New("the engine produced no weights")
	}
	res, err := weightsResult(req, engineRes, iterationsRun)
	if err != nil {
		return api.SimResult{}, err
	}
	// Not set on an abort: stopped() never carried one either, since
	// there is nothing about the elapsed wall time worth reporting for
	// a run that did not finish.
	if !res.Aborted {
		res.DurationMS = time.Since(start).Milliseconds()
	}
	return res, nil
}

// simWeights(requestJSON, callbackId) computes stat weights and
// returns a SimResult with Weights filled. Progress is reported the
// way simRun reports it.
func simWeights(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return errorJSON("simWeights takes (requestJSON, callbackId)")
	}
	return weightsJSON(args[0].String(), args[1].String())
}
