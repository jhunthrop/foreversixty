//go:build js && wasm

// Command wasm is the browser half of the sim. It exports exactly four
// functions, all taking and returning JSON strings, plus one string
// global, simEngineVersion, which is the engine sha this module was
// built from. The page reads that global rather than being told the sha
// by the server, so a request can never name an engine the wasm it is
// running in is not.
//
// It exists in this repository rather than in the engine because
// sim/request and sim/adapter are linked in here: the browser gets a
// finished SimResult with its summary.Summary already built by the same
// Go code the server runs. The engine's own thirteen js.Global().Set
// entrypoints are an implementation detail behind these four and the web
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
	"strings"
	"syscall/js"
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
	"github.com/wowsims/classic/sim/core/simsignals"
)

func main() {
	engine.RegisterAll()
	core.SetRunningInWasm()

	js.Global().Set("simRun", js.FuncOf(simRun))
	js.Global().Set("simSplit", js.FuncOf(simSplit))
	js.Global().Set("simCombine", js.FuncOf(simCombine))
	js.Global().Set("simAbort", js.FuncOf(simAbort))
	// The pin is compiled in, not injected: a plain `go build ./cmd/wasm`
	// used to produce "dev" and stamp it on real rows.
	js.Global().Set("simEngineVersion", js.ValueOf(enginever.Version))

	// The host page defines wasmready and is told the moment the four
	// exports exist, so it never races them.
	js.Global().Call("wasmready")
	select {}
}

// fail wraps an error as a SimResult, so every export returns the same
// shape and the worker never has to distinguish a throw from a result.
func fail(req api.SimRequest, msg string) string {
	return result(api.SimResult{Request: req, Error: msg, Summary: adapter.EmptySummary()})
}

// stopped wraps an abort. It is not a failure - the user pressed Stop -
// so it carries no error message and the page renders it as a run that
// ended early rather than as something that went wrong.
func stopped(req api.SimRequest, iterations int) string {
	return result(api.SimResult{Request: req, Aborted: true, IterationsRun: iterations, Summary: adapter.EmptySummary()})
}

// Both carry adapter.EmptySummary() rather than a zero summary: a nil
// Go slice marshals as null, so every export returns the same shape and
// the page never has to null-check sixteen keys on the paths it is least
// likely to have exercised.
//
// result stamps the two fields every export must fill the same way and
// encodes. EngineVersion is enginever.Version, never the request's
// claim: the row's provenance is a fact about the binary that produced
// it, and api.SimResult.Stale can only fire if it is.
func result(res api.SimResult) string {
	res.EngineVersion = enginever.Version
	res.Lane = api.LaneBrowser
	b, err := json.Marshal(res)
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}

// decodeRequest parses one request strictly. A field the envelope does
// not carry is a client sending something this build cannot honour -
// a profession list to an older wasm, say - and running anyway would
// drop it silently.
func decodeRequest(s string) (api.SimRequest, error) {
	var req api.SimRequest
	dec := json.NewDecoder(strings.NewReader(s))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		return req, err
	}
	return req, nil
}

// simRun(requestJSON, callbackId) runs one request to completion and
// returns SimResult JSON. Progress is reported by calling the global
// simProgress(callbackId, progressJSON), where progressJSON is an
// api.Progress - Pick<SimResult, 'iterations_run' | 'dps'> - so the page
// reads a partial result with the accessors it already has. JSON in,
// JSON out, like everything else here.
func simRun(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return fail(api.SimRequest{}, "simRun takes (requestJSON, callbackId)")
	}
	req, err := decodeRequest(args[0].String())
	if err != nil {
		return fail(req, "the request is not valid JSON: "+err.Error())
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
		return fail(req, err.Error())
	}
	// Forever's own item rows, from the active build, embedded at build
	// time: the browser has no protobuf and cannot send them.
	if err := simdb.Attach(engineReq); err != nil {
		return fail(req, err.Error())
	}

	start := time.Now()
	reporter := make(chan *proto.ProgressMetrics, 32)
	// Threading does not work in wasm, so this is the serial entrypoint.
	// Parallelism comes from the worker pool: the page calls simSplit and
	// gives each worker one part.
	core.RunRaidSimAsync(engineReq, reporter, callbackID)

	var engineRes *proto.RaidSimResult
	for p := range reporter {
		if p.FinalRaidResult != nil {
			// Not a bare break: see the matching comment in
			// sim/cmd/forever-sim's execute for the full reasoning.
			// Short version: draining to the channel's close (instead
			// of breaking the instant a FinalRaidResult arrives) is
			// what makes a sample request's SampleIteration mutation
			// guaranteed-visible, but two engine paths send a
			// FinalRaidResult and then return WITHOUT ever closing the
			// channel - a failed simsignals.RegisterWithId, and
			// SimOptions.IsTest (which this package's requests never
			// set). Neither carries a sample, so an error result is
			// still safe to take immediately rather than block
			// forever waiting for a close that will not come.
			engineRes = p.FinalRaidResult
			if engineRes.Error != nil {
				break
			}
			continue
		}
		if cb := js.Global().Get("simProgress"); cb.Type() == js.TypeFunction {
			if b, err := json.Marshal(api.Progress{
				IterationsRun: int(p.CompletedIterations),
				DPS:           api.Estimate{Mean: p.Dps},
			}); err == nil {
				cb.Invoke(callbackID, string(b))
			}
		}
	}
	if engineRes == nil {
		return fail(req, "the engine produced no result")
	}
	// An abort's ErrorOutcome carries no message, so this switches on
	// the type. Testing the message alone let Stop through as a result
	// with zero iterations, which the adapter then called corrupt.
	if err := adapter.ResultError(engineRes); err != nil {
		if errors.Is(err, adapter.ErrAborted) {
			return stopped(req, int(engineRes.IterationsDone))
		}
		return fail(req, err.Error())
	}

	sum, err := adapter.Summarize(engineRes, req)
	if err != nil {
		return fail(req, err.Error())
	}
	return result(api.SimResult{
		Request:       req,
		DPS:           adapter.DPS(engineRes),
		IterationsRun: int(engineRes.IterationsDone),
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       sum,
		Sample:        adapter.Sample(engineRes),
	})
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
// is distinguishable, and so that all four exports have one shape.
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

// errorJSON wraps a message as the {"error": "..."} shape simSplit's
// caller reads. It goes through the JSON encoder rather than string
// concatenation, because an engine error message carries quotes and a
// hand-built string would hand the worker something it cannot parse.
func errorJSON(msg string) string {
	b, err := json.Marshal(struct {
		Error string `json:"error"`
	}{msg})
	if err != nil {
		return `{"error":"the error itself could not be encoded"}`
	}
	return string(b)
}
