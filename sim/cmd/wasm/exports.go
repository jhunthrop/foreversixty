package main

// The JSON half of the exports.
//
// Everything the browser calls is a JSON string in and a JSON string
// out, and the js.Func wrappers in main.go do nothing but unwrap the
// arguments. The work lives here, with NO build tag, so it is compiled
// and tested on the host: a bug in the planner's boundary would
// otherwise only ever be found in a browser.

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/wowsims/classic/sim/core/proto"
)

// strictDecode refuses a field the type does not carry. A page sending
// something this wasm cannot honour is a version mismatch, and running
// anyway would drop part of it silently.
func strictDecode(s string, into any) error {
	dec := json.NewDecoder(strings.NewReader(s))
	dec.DisallowUnknownFields()
	return dec.Decode(into)
}

// decodeRequest is strictDecode into an api.SimRequest - the one shape
// every export's first argument is.
func decodeRequest(s string) (api.SimRequest, error) {
	var req api.SimRequest
	err := strictDecode(s, &req)
	return req, err
}

// errorJSON wraps a message as the {"error": "..."} shape every export
// returns on failure. It goes through the JSON encoder rather than
// string concatenation, because an error message carries quotes and a
// hand-built string would hand the caller something it cannot parse.
func errorJSON(msg string) string {
	b, err := json.Marshal(struct {
		Error string `json:"error"`
	}{msg})
	if err != nil {
		return `{"error":"the error itself could not be encoded"}`
	}
	return string(b)
}

// planJSON is simPlan's body: a bulk SimRequest in, the first stage
// out.
//
// The page then runs each of stage.requests as a WHOLE, UNSPLIT run,
// in chunks of the pool's width, and hands the results back to
// simRank in the same order. Stage requests are not sharded: a stage
// is already many independent sims, sharding each of them would
// multiply the scheduling for nothing, and a stage that is a queue of
// whole runs is what makes "abort returns what finished" true.
// simSplit and simCombine are for plain runs. No statistics happen in
// TypeScript.
func planJSON(requestJSON string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	stage, err := bulk.Plan(req)
	if err != nil {
		return planErrorJSON(err)
	}
	return encodeOrError(stage)
}

// rankJSON is simRank's body: the request, the stage it answers, and
// one result per request in that stage's order.
//
// It returns {"next": stage} or {"result": SimResult}, never both,
// which is how the page knows whether to loop.
func rankJSON(requestJSON, stageJSON, resultsJSON string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	var stage bulk.StageRequests
	if err := strictDecode(stageJSON, &stage); err != nil {
		return errorJSON("the stage is not valid JSON: " + err.Error())
	}
	var results []api.SimResult
	if err := strictDecode(resultsJSON, &results); err != nil {
		return errorJSON("the results are not valid JSON: " + err.Error())
	}
	next, final, err := bulk.Rank(req, stage, results)
	if err != nil {
		return errorJSON(err.Error())
	}
	out := struct {
		Next   *bulk.StageRequests `json:"next,omitempty"`
		Result *api.SimResult      `json:"result,omitempty"`
	}{Next: next, Result: final}
	return encodeOrError(out)
}

// countJSON is simCount's body: how many combinations this request
// would plan, without RETAINING more than the cap's worth of them and
// without enumerating more than bulk.ExpandWorkBudget of them. The
// page asks on every candidate tick, which is why it is not "call
// simPlan and count the answer": simPlan retains every combination it
// plans, and this does not.
//
// It does build a request per combination as it counts - validity can
// only be asked of a finished gear list - so the honest guarantee is
// the budget, not "allocates nothing". Past the budget the answer is
// a cap refusal quoting an arithmetic upper bound rather than an
// enumerated count; below it, the count is exact and agrees with
// simPlan exactly.
func countJSON(requestJSON string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	n, err := bulk.Count(req)
	if err != nil {
		return planErrorJSON(err)
	}
	return encodeOrError(struct {
		Combinations int `json:"combinations"`
	}{n})
}

// needsMoreJSON is simNeedsMore's body. The decision itself is
// api.NeedsMoreIterations, in Go, because the native binary asks the
// same question and the two lanes must stop at the same precision.
// This export exists so the page can ask it without a copy.
func needsMoreJSON(resultJSON, requestJSON string) string {
	var res api.SimResult
	if err := strictDecode(resultJSON, &res); err != nil {
		return errorJSON("the result is not valid JSON: " + err.Error())
	}
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	return encodeOrError(struct {
		NeedsMore bool `json:"needs_more"`
	}{api.NeedsMoreIterations(res, req)})
}

// fieldError is one validation failure, addressed to the control that
// produced it.
type fieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// validateJSON is simValidate's body: the request drawer runs the
// SAME api.SimRequest.Validate the run will, and shows the failures
// inline. Returning them one per field rather than as one joined
// sentence is what lets the drawer mark the control.
//
// This is why an unparsable body still comes back as
// {"ok": false, "errors": [...]} rather than the bare {"error": "..."}
// shape: simValidate's contract (spec 10.2) is "any request in,
// {ok, errors} out", so a request the drawer cannot even parse is
// still a validation failure, not a thrown exception. The wrapper in
// main.go has its own separate {"error": "..."} case for a caller that
// gets simValidate's own arity wrong - that one is a programming
// mistake in the caller, not a request from a member, and never
// reaches this function.
func validateJSON(requestJSON string) string {
	out := struct {
		OK     bool         `json:"ok"`
		Errors []fieldError `json:"errors"`
	}{Errors: []fieldError{}}

	req, err := decodeRequest(requestJSON)
	if err != nil {
		out.Errors = append(out.Errors, fieldError{Message: "the request is not valid JSON: " + err.Error()})
		return encodeOrError(out)
	}
	// The browser lane's numbers: this is the page asking, and a
	// request the page cannot run is a failure the page must show.
	for _, e := range flatten(req.ValidateLane(api.LaneBrowser)) {
		out.Errors = append(out.Errors, fieldError{Field: leadingField(e), Message: e})
	}
	out.OK = len(out.Errors) == 0
	return encodeOrError(out)
}

// flatten pulls the messages out of an errors.Join tree, one per
// failure, in the order Validate appended them.
func flatten(err error) []string {
	if err == nil {
		return nil
	}
	if joined, ok := err.(interface{ Unwrap() []error }); ok {
		var out []string
		for _, e := range joined.Unwrap() {
			out = append(out, flatten(e)...)
		}
		return out
	}
	return []string{err.Error()}
}

// fieldToken matches the leading path of a validation message -
// "iterations must be...", "bulk.candidates[2] is on..." - so the
// drawer can mark the control that produced it.
var fieldToken = regexp.MustCompile(`^[a-z][a-z0-9_]*(\.[a-z0-9_]+|\[[0-9]+\])*`)

// topLevelFields is every SimRequest field Validate/ValidateLane names,
// bare, at the very start of one of its own messages: "iterations must
// be...", "targets must be...", "weights needs at least...". A nested
// path ("bulk.candidates[2]", "encounter.movement.kind") is trusted on
// its shape alone, but a bare word is not, because ValidateLane's own
// added messages ("the browser lane plans...", "the browser lane
// runs...") and several of Validate's ("a split part's iterations...",
// "a request is one kind...") also start with an ordinary lowercase
// word - "the", "a" - that the regex alone cannot tell from a field
// name. Without this list leadingField reports field "the" or "a" for
// exactly those messages, which marks no control in the drawer and
// never falls back to the top-of-drawer message its own doc comment
// promises.
//
// This list was built by walking every errors.New/fmt.Errorf call
// under sim/api reachable from Validate/ValidateLane and noting which
// ones are bare rather than dotted - not by reasoning from the field
// list, since a field missing here (weights, initially) fails exactly
// the same way a stray "the"/"a" does: silently, with no test to catch
// it. Every other top-level field (source, character, encounter, bulk)
// only ever appears in a dotted message ("source.kind", "bulk.mode",
// ...), so it needs no entry here - the dot already trusts it.
var topLevelFields = map[string]bool{
	"iterations":     true,
	"targets":        true,
	"duration_sec":   true,
	"spec":           true,
	"lane":           true,
	"engine_version": true,
	"execute_ratio":  true,
	"variation":      true,
	"target_error":   true,
	"weights":        true,
}

// leadingField is the message's field path, or "" when the message is
// prose rather than a field name - which the drawer shows at the top
// instead of against a control.
//
// A trailing colon is stripped before matching so a wrapped error -
// "bulk.candidates[2]: origin must be one of ..." from
// fmt.Errorf("bulk.candidates[%d]: %w", i, err) - still reports the
// path in front of the colon, rather than losing its field to the one
// character the wrap added.
func leadingField(msg string) string {
	head, _, _ := strings.Cut(msg, " ")
	head = strings.TrimSuffix(head, ":")
	token := fieldToken.FindString(head)
	if token != head {
		return ""
	}
	if strings.ContainsAny(token, ".[") {
		return token
	}
	if topLevelFields[token] {
		return token
	}
	return ""
}

// planErrorJSON is the error shape for simPlan and simCount.
//
// A cap breach is STRUCTURED - {"error":"cap_exceeded","cap":n,
// "combinations":n} - because the page draws a notice with both
// numbers in it and the API answers 400 cap_exceeded with the same
// two fields. Either parsing them out of a sentence would break the
// first time the sentence was reworded.
func planErrorJSON(err error) string {
	var capped api.ErrCapExceeded
	if errors.As(err, &capped) {
		// api.ErrCapExceeded already carries json:"cap" and
		// json:"combinations" (sim/api/bulk.go); embedding it rather
		// than repeating the two keys here keeps the web's two field
		// names to one source of truth.
		return encodeOrError(struct {
			Error string `json:"error"`
			api.ErrCapExceeded
		}{"cap_exceeded", capped})
	}
	return errorJSON(err.Error())
}

// encodeOrError marshals, or returns the {"error": ...} shape.
func encodeOrError(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}

// stamp fills the two fields every export sets the same way.
// EngineVersion is enginever.Version, never the request's claim: a
// row's provenance is a fact about the binary that produced it.
func stamp(res api.SimResult) api.SimResult {
	res.EngineVersion = enginever.Version
	res.Lane = api.LaneBrowser
	return res
}

// failJSON wraps an error as a SimResult, so every export returns the
// same shape and the worker never has to distinguish a throw from a
// result. The empty summary rather than a zero one: a nil Go slice
// marshals as null, and the page would need a null check per key on
// the path least likely to be exercised.
func failJSON(req api.SimRequest, msg string) string {
	return encodeOrError(stamp(api.SimResult{Request: req, Error: msg, Summary: adapter.EmptySummary()}))
}

// weightsRunner is the engine half of simWeights, supplied by the
// wasm build. The host build has none, so weightsJSON's refusals are
// testable without syscall/js and the engine call is not duplicated.
var weightsRunner func(req api.SimRequest, callbackID string) (api.SimResult, error)

// weightsJSON is simWeights' body: a weights SimRequest in, a
// SimResult with Weights out. Progress is reported through the same
// simProgress global a plain run uses.
func weightsJSON(requestJSON, callbackID string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return failJSON(req, "the request is not valid JSON: "+err.Error())
	}
	if req.Weights == nil {
		return failJSON(req, "simWeights takes a request with a weights block; this one has none")
	}
	if err := req.ValidateLane(api.LaneBrowser); err != nil {
		return failJSON(req, err.Error())
	}
	if weightsRunner == nil {
		return failJSON(req, "this build has no engine; simWeights runs only in the browser")
	}
	res, err := weightsRunner(req, callbackID)
	if err != nil {
		return failJSON(req, err.Error())
	}
	return encodeOrError(stamp(res))
}

// weightsResult turns the engine's raw StatWeightsResult into the
// SimResult simWeights answers with, or an error for the caller to wrap
// as a failure. It touches no js.Value and calls no engine function, so
// - unlike runWeights, which drives core.StatWeightsAsync and so needs
// the wasm build - the accounting below is testable on the host:
//
//   - Abort. The engine reports Stop the same way here as on a plain
//     run: an ErrorOutcome with Type ErrorOutcomeAborted and an EMPTY
//     Message (fork core/sim.go; asserted for weights specifically by
//     core/simsignals/api_test.go's StatWeightsAsync case).
//     adapter.Weights only branches on a NON-EMPTY message, so an abort
//     would otherwise fall through to res.GetDps() == nil and come back
//     as ErrNoWeights - "the result carries no stat weights" - which
//     misdescribes a Stop as a corrupt result. This mirrors simRun's
//     adapter.ResultError type check for RaidSimResult; StatWeightsResult
//     has no ResultError of its own to call.
//   - Iterations. iterationsRun is the caller's running total from the
//     engine's own progress ticks (ProgressMetrics.CompletedIterations),
//     which is the real number of iterations a sweep ran - req.Iterations
//     is only the PER-SIM count, and a sweep runs 2*len(stats)+1 sims. A
//     caller that saw no progress tick before the final message (e.g. a
//     sweep with no stats) passes 0, which falls back to req.Iterations
//     rather than reporting a bare zero.
func weightsResult(req api.SimRequest, engineRes *proto.StatWeightsResult, iterationsRun int) (api.SimResult, error) {
	if iterationsRun == 0 {
		iterationsRun = req.Iterations
	}
	if engineRes.GetError().GetType() == proto.ErrorOutcomeType_ErrorOutcomeAborted {
		return api.SimResult{
			Request:       req,
			Aborted:       true,
			IterationsRun: iterationsRun,
			Summary:       adapter.EmptySummary(),
		}, nil
	}
	weights, err := adapter.Weights(engineRes, req)
	if err != nil {
		return api.SimResult{}, err
	}
	return api.SimResult{
		Request:       req,
		IterationsRun: iterationsRun,
		Summary:       adapter.EmptySummary(),
		Weights:       weights,
	}, nil
}
