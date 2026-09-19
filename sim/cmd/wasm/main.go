//go:build js && wasm

// Command wasm is the browser half of the sim. It exports exactly four
// functions, all taking and returning JSON strings.
//
// It exists in this repository rather than in the engine because
// sim/request and sim/adapter are linked in here: the browser gets a
// finished SimResult with its summary.Summary already built by the same
// Go code the server runs. The engine's own thirteen js.Global().Set
// entrypoints are an implementation detail behind these four and the web
// must not call them.
//
// This build carries no item database. The engine's --tags=with_db
// embeds a 4.9 MB table, which measures 3.79 MB gzipped here against a
// 4 MB budget, so the browser is meant to be handed the handful of items
// a request actually equips rather than all of them - the same
// arrangement request.Options.Consumables already has for the build's
// consumable table. Until that lands a browser sim of a geared character
// comes back with the engine's "No item with id"; sim/cmd/forever-sim,
// which has the budget for it, is built with the tag.
package main

import (
	"encoding/json"
	"syscall/js"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/combine"
	"github.com/jhunthrop/foreversixty/sim/request"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// Version is set at build time to the pinned engine sha.
var Version = "dev"

func main() {
	engine.RegisterAll()
	core.SetRunningInWasm()

	js.Global().Set("simRun", js.FuncOf(simRun))
	js.Global().Set("simSplit", js.FuncOf(simSplit))
	js.Global().Set("simCombine", js.FuncOf(simCombine))
	js.Global().Set("simAbort", js.FuncOf(simAbort))
	js.Global().Set("simEngineVersion", js.ValueOf(Version))

	// The host page defines wasmready and is told the moment the four
	// exports exist, so it never races them.
	js.Global().Call("wasmready")
	select {}
}

// fail wraps an error as a SimResult, so every export returns the same
// shape and the worker never has to distinguish a throw from a result.
func fail(req api.SimRequest, msg string) string {
	b, _ := json.Marshal(api.SimResult{
		EngineVersion: req.EngineVersion,
		Request:       req,
		Lane:          api.LaneBrowser,
		Error:         msg,
	})
	return string(b)
}

// simRun(requestJSON, callbackId) runs one request to completion and
// returns SimResult JSON. Progress is reported by calling the global
// simProgress(callbackId, completed, total, dps).
func simRun(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return fail(api.SimRequest{}, "simRun takes (requestJSON, callbackId)")
	}
	var req api.SimRequest
	if err := json.Unmarshal([]byte(args[0].String()), &req); err != nil {
		return fail(req, "the request is not valid JSON: "+err.Error())
	}
	callbackID := args[1].String()

	// Every request the browser runs is a part: the worker pool calls
	// simSplit first, even for one worker, and a part's iteration count
	// is not one of api.ValidIterations by construction - 3,000 over
	// four workers is 750. SplitPart relaxes that check and nothing
	// else. The whole request was validated before it was split, by the
	// page and by the api lane.
	engineReq, err := request.BuildWith(req, request.Options{SplitPart: true})
	if err != nil {
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
			engineRes = p.FinalRaidResult
			break
		}
		if cb := js.Global().Get("simProgress"); cb.Type() == js.TypeFunction {
			cb.Invoke(callbackID, int(p.CompletedIterations), int(p.TotalIterations), p.Dps)
		}
	}
	if engineRes == nil {
		return fail(req, "the engine produced no result")
	}
	if engineRes.Error != nil && engineRes.Error.Message != "" {
		return fail(req, engineRes.Error.Message)
	}

	sum, err := adapter.Summarize(engineRes, req)
	if err != nil {
		return fail(req, err.Error())
	}
	b, err := json.Marshal(api.SimResult{
		EngineVersion: req.EngineVersion,
		Request:       req,
		Lane:          api.LaneBrowser,
		DPS:           adapter.DPS(engineRes),
		IterationsRun: int(engineRes.IterationsDone),
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       sum,
	})
	if err != nil {
		return fail(req, err.Error())
	}
	return string(b)
}

// simSplit(requestJSON, n) returns a JSON array of n request JSONs, one
// per worker, with the seeds already offset.
func simSplit(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return `{"error":"simSplit takes (requestJSON, n)"}`
	}
	var req api.SimRequest
	if err := json.Unmarshal([]byte(args[0].String()), &req); err != nil {
		return `{"error":"the request is not valid JSON"}`
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
		return `{"error":"simCombine takes (resultsJSON)"}`
	}
	var parts []api.SimResult
	if err := json.Unmarshal([]byte(args[0].String()), &parts); err != nil {
		return `{"error":"the results are not valid JSON"}`
	}
	out, err := combine.Results(parts)
	if err != nil {
		return fail(api.SimRequest{}, err.Error())
	}
	b, err := json.Marshal(out)
	if err != nil {
		return fail(api.SimRequest{}, err.Error())
	}
	return string(b)
}

// simAbort(callbackId) stops a run started with the same id.
func simAbort(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return false
	}
	return simsignals.AbortById(args[0].String())
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
