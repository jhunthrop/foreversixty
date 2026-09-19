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

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
)

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
// would plan, without allocating a single request. The page asks on
// every candidate tick, which is why it is not "call simPlan and
// count the answer".
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
// drawer can mark the control that produced it. A message that does
// not start with one gets an empty field and is shown at the top.
var fieldToken = regexp.MustCompile(`^[a-z][a-z0-9_]*(\.[a-z0-9_]+|\[[0-9]+\])*`)

// leadingField is the message's field path, or "".
func leadingField(msg string) string {
	head, _, _ := strings.Cut(msg, " ")
	if fieldToken.FindString(head) != head {
		return ""
	}
	return head
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
		return encodeOrError(struct {
			Error        string `json:"error"`
			Cap          int    `json:"cap"`
			Combinations int    `json:"combinations"`
		}{"cap_exceeded", capped.Cap, capped.Combinations})
	}
	return errorJSON(err.Error())
}

// strictDecode refuses a field the type does not carry, the way
// decodeRequest does: a page sending something this wasm cannot honour
// is a version mismatch, and running anyway would drop part of it.
func strictDecode(s string, into any) error {
	dec := json.NewDecoder(strings.NewReader(s))
	dec.DisallowUnknownFields()
	return dec.Decode(into)
}

// encodeOrError marshals, or returns the {"error": ...} shape.
func encodeOrError(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}
