//go:build !js

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// sim/cmd/wasm is a js/wasm-only package: syscall/js does not build on a
// host platform, so main.go's js.Func wrappers cannot be exercised here.
// The JSON bodies they wrap - planJSON, rankJSON, countJSON,
// needsMoreJSON, validateJSON and the rest of exports.go - carry no
// build tag and are tested below; only the js.Value unwrapping in
// main.go is not, and that unwrapping is a one-line args[i].String()
// call per wrapper. The four original exports' behaviour is covered
// three further ways:
//
//   - sim/combine's tests cover simSplit and simCombine, which are thin
//     wrappers over Split and Results;
//   - sim/cmd/forever-sim's tests cover the request -> engine -> adapter
//     pipeline that simRun runs, against the same code;
//   - api's tests cover the progress payload's shape, which is
//     api.Progress because it is a contract with the web rather than a
//     detail of this file;
//   - the CI smoke test in .github/workflows/sim.yml instantiates the
//     built wasm under node and asserts simRun, simSplit, simCombine
//     and simAbort exist, that simRun returns a SimResult with a
//     summary, and that the progress callback was called with that
//     payload. It does not yet check the five bulk exports this file
//     tests; that is Task 25's file to extend.
func TestWasmIsCoveredElsewhere(t *testing.T) {
	t.Log("see the comment above: combine, forever-sim, and the CI smoke test")
}

// The two bulk exports are thin: the wrappers unwrap js.Value and the
// functions below do the work, so the work is testable on the host
// and the wrappers carry nothing that could be wrong.
func TestPlanJSONRoundTrip(t *testing.T) {
	req := bulkFixtureRequest(t)
	out := planJSON(req)
	var stage bulk.StageRequests
	if err := json.Unmarshal([]byte(out), &stage); err != nil {
		t.Fatalf("simPlan returned %s", out)
	}
	if stage.Stage != 1 || len(stage.Requests) != len(stage.Combos)+1 {
		t.Errorf("stage = %+v", stage)
	}
}

func TestPlanJSONReportsAnError(t *testing.T) {
	out := planJSON("not json")
	var e struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(out), &e); err != nil || e.Error == "" {
		t.Errorf("simPlan of rubbish returned %s", out)
	}
}

func TestRankJSONReturnsTheNextStageThenTheResult(t *testing.T) {
	reqJSON := bulkFixtureRequest(t)
	stageJSON := planJSON(reqJSON)

	var stage bulk.StageRequests
	if err := json.Unmarshal([]byte(stageJSON), &stage); err != nil {
		t.Fatal(err)
	}
	resultsJSON := fakeResultsJSON(t, stage)

	out := rankJSON(reqJSON, stageJSON, resultsJSON)
	var first struct {
		Next   *bulk.StageRequests `json:"next"`
		Result *api.SimResult      `json:"result"`
	}
	if err := json.Unmarshal([]byte(out), &first); err != nil {
		t.Fatalf("simRank returned %s", out)
	}
	if first.Next == nil || first.Result != nil {
		t.Fatalf("the first rung returned %s", out)
	}
	if len(first.Next.Ran) == 0 {
		t.Errorf("next.ran is empty; the ladder's history should thread across the wasm boundary (contract A10)")
	}

	nextJSON, err := json.Marshal(first.Next)
	if err != nil {
		t.Fatal(err)
	}
	out = rankJSON(reqJSON, string(nextJSON), fakeResultsJSON(t, *first.Next))
	var last struct {
		Next   *bulk.StageRequests `json:"next"`
		Result *api.SimResult      `json:"result"`
	}
	if err := json.Unmarshal([]byte(out), &last); err != nil {
		t.Fatalf("simRank returned %s", out)
	}
	if last.Next != nil || last.Result == nil {
		t.Fatalf("the last rung returned %s", out)
	}
	if len(last.Result.Combos) == 0 || last.Result.Equipped == nil {
		t.Errorf("the finished result is %+v", last.Result)
	}
	if len(last.Result.Stages) != 2 {
		t.Errorf("result.stages = %+v, want 2 entries for the normal ladder's two rungs", last.Result.Stages)
	}
}

// The count the page shows as candidates are ticked is the planner's
// own, so a cap notice and a refused run are the same number.
func TestCountJSON(t *testing.T) {
	var got struct {
		Combinations int `json:"combinations"`
	}
	if err := json.Unmarshal([]byte(countJSON(bulkFixtureRequest(t))), &got); err != nil {
		t.Fatal(err)
	}
	if got.Combinations != 3 {
		t.Errorf("combinations = %d, want 3 for two candidates in two slots", got.Combinations)
	}
}

// A cap breach is a STRUCTURED error, not a sentence: the page draws a
// notice with both numbers in it and the API answers 400 cap_exceeded
// with the same two fields, so neither parses prose.
func TestCapExceededIsStructured(t *testing.T) {
	var req api.SimRequest
	if err := json.Unmarshal([]byte(bulkFixtureRequest(t)), &req); err != nil {
		t.Fatal(err)
	}
	req.Bulk.Cap = 1
	body, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	for _, out := range []string{planJSON(string(body)), countJSON(string(body))} {
		var got struct {
			Error        string `json:"error"`
			Cap          int    `json:"cap"`
			Combinations int    `json:"combinations"`
		}
		if err := json.Unmarshal([]byte(out), &got); err != nil {
			t.Fatalf("returned %s", out)
		}
		if got.Error != "cap_exceeded" || got.Cap != 1 || got.Combinations <= 1 {
			t.Errorf("cap breach returned %s", out)
		}
	}
}

// The Smart Sim decision crosses the boundary so the page never
// reimplements it.
func TestNeedsMoreJSON(t *testing.T) {
	reqJSON := bulkFixtureRequest(t)
	var req api.SimRequest
	if err := json.Unmarshal([]byte(reqJSON), &req); err != nil {
		t.Fatal(err)
	}
	req.Bulk = nil
	req.TargetError = 0.005
	req.Iterations = 30000
	body, _ := json.Marshal(req)

	for _, c := range []struct {
		name string
		res  api.SimResult
		want bool
	}{
		{"still noisy", api.SimResult{IterationsRun: 2000, DPS: api.Estimate{Mean: 1000, Error: 20}}, true},
		{"precise enough", api.SimResult{IterationsRun: 2000, DPS: api.Estimate{Mean: 1000, Error: 4}}, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			resJSON, _ := json.Marshal(c.res)
			var got struct {
				NeedsMore bool `json:"needs_more"`
			}
			out := needsMoreJSON(string(resJSON), string(body))
			if err := json.Unmarshal([]byte(out), &got); err != nil {
				t.Fatalf("returned %s", out)
			}
			if got.NeedsMore != c.want {
				t.Errorf("needs_more = %v, want %v", got.NeedsMore, c.want)
			}
		})
	}
}

// The request drawer validates with the SAME Validate the run uses,
// and shows the errors inline, so it needs them one per field rather
// than as one joined sentence.
func TestValidateJSON(t *testing.T) {
	var ok struct {
		OK     bool `json:"ok"`
		Errors []struct {
			Field   string `json:"field"`
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.Unmarshal([]byte(validateJSON(plainFixtureRequest(t))), &ok); err != nil {
		t.Fatal(err)
	}
	if !ok.OK || len(ok.Errors) != 0 {
		t.Errorf("a valid request came back %+v", ok)
	}

	var req api.SimRequest
	if err := json.Unmarshal([]byte(plainFixtureRequest(t)), &req); err != nil {
		t.Fatal(err)
	}
	req.Iterations = 7
	req.Encounter.Targets = 99
	body, _ := json.Marshal(req)
	out := validateJSON(string(body))
	if err := json.Unmarshal([]byte(out), &ok); err != nil {
		t.Fatalf("returned %s", out)
	}
	if ok.OK || len(ok.Errors) < 2 {
		t.Fatalf("two broken fields came back %s", out)
	}
	fields := map[string]bool{}
	for _, e := range ok.Errors {
		fields[e.Field] = true
		if e.Message == "" {
			t.Errorf("an error with no message: %+v", e)
		}
	}
	for _, want := range []string{"iterations", "targets"} {
		if !fields[want] {
			t.Errorf("no error names %q: %s", want, out)
		}
	}
}

// Rubbish is an error shape, not a throw, on every export.
func TestValidateJSONOfRubbish(t *testing.T) {
	var got struct {
		OK     bool `json:"ok"`
		Errors []struct {
			Field   string `json:"field"`
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.Unmarshal([]byte(validateJSON("{")), &got); err != nil {
		t.Fatal(err)
	}
	if got.OK || len(got.Errors) != 1 {
		t.Errorf("rubbish came back %+v", got)
	}
}

// A bare lowercase word at the front of a message is not automatically
// a field: ValidateLane's own added messages, and several of
// Validate's, start with ordinary prose ("the", "a") that the regex
// alone cannot distinguish from a real field name. The first four cases
// are measured straight off ValidateLane(LaneBrowser)'s real output for
// a request over both lane caps and a bulk-plus-target-error request.
func TestLeadingField(t *testing.T) {
	for _, c := range []struct {
		msg  string
		want string
	}{
		{"the browser lane plans at most 400 combinations, and bulk.cap is 500", ""},
		{"the browser lane runs at most 30000 iterations, and iterations is 10000000", ""},
		{"a target-error run's iterations is at most 100000 on this lane, got 10000000", ""},
		{"a request is one kind: a bulk or weights request runs its own iteration counts, so it carries no target_error", ""},
		{"iterations must be one of [500 3000 10000], got 7", "iterations"},
		{"targets must be between 1 and 10, got 99", "targets"},
		{"bulk.candidates[2] is on \"head\", which is locked", "bulk.candidates[2]"},
		// weights is a real top-level SimRequest field reachable
		// through ValidateLane, and, unlike every other top-level
		// field, api/weights.go's own emptiness check produces it
		// bare rather than dotted.
		{"weights needs at least one stat to weigh", "weights"},
		// bulk.go wraps validateOrigin's error as
		// fmt.Errorf("bulk.candidates[%d]: %w", i, err); the trailing
		// colon must not cost the wrapped message its field.
		{`bulk.candidates[2]: origin must be one of [equipped bag bank search], "drop:"<id> or "set:"<name>, got "junk"`, "bulk.candidates[2]"},
	} {
		if got := leadingField(c.msg); got != c.want {
			t.Errorf("leadingField(%q) = %q, want %q", c.msg, got, c.want)
		}
	}
}

func TestRankJSONReportsAMismatch(t *testing.T) {
	reqJSON := bulkFixtureRequest(t)
	out := rankJSON(reqJSON, planJSON(reqJSON), `[]`)
	var e struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(out), &e); err != nil || e.Error == "" {
		t.Errorf("simRank of no results returned %s", out)
	}
}

// plainFixtureRequest is a plain (non-bulk) request over the checked-in
// warrior fixture, as JSON, valid against LaneBrowser. simValidate's
// happy path needs a request that is not bulk, since bulkFixtureRequest
// exists for the bulk exports.
func plainFixtureRequest(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "adapter", "testdata", "warrior-fury.request.json"))
	if err != nil {
		t.Fatal(err)
	}
	var req api.SimRequest
	if err := json.Unmarshal(b, &req); err != nil {
		t.Fatal(err)
	}
	req.EngineVersion = enginever.Version
	out, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

// bulkFixtureRequest is a two-candidate Top Gear over the checked-in
// warrior fixture, as JSON.
func bulkFixtureRequest(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "adapter", "testdata", "warrior-fury.request.json"))
	if err != nil {
		t.Fatal(err)
	}
	var req api.SimRequest
	if err := json.Unmarshal(b, &req); err != nil {
		t.Fatal(err)
	}
	req.EngineVersion = enginever.Version
	req.Iterations = 3000
	head := req.Character.Gear[0]
	req.Bulk = &api.BulkSpec{
		Mode:      api.KindGear,
		Precision: api.PrecisionNormal,
		Cap:       api.Caps[api.LaneBrowser],
		Candidates: []api.Candidate{
			{Slot: "head", ItemID: head.ItemID, Origin: api.OriginBag},
			{Slot: "neck", ItemID: req.Character.Gear[1].ItemID, Origin: api.OriginBag},
		},
	}
	out, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

// fakeResultsJSON answers a stage with plausible numbers, so the wasm
// boundary can be tested without running the engine on the host.
func fakeResultsJSON(t *testing.T, stage bulk.StageRequests) string {
	t.Helper()
	out := make([]api.SimResult, len(stage.Requests))
	for i := range stage.Requests {
		out[i] = api.SimResult{
			EngineVersion: enginever.Version,
			Request:       stage.Requests[i],
			Lane:          api.LaneBrowser,
			IterationsRun: stage.Iterations,
			DPS:           api.Estimate{Mean: 1000 + float64(i)*10, StdDev: 200, Error: 4},
			Summary:       adapter.EmptySummary(),
		}
	}
	b, err := json.Marshal(out)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
