package llm

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// spendReq is 300 prompt characters (100 tokens at the worst-case ratio) and a 500-token
// ceiling; at prices 1 and 2 per million its worst case is (100*1 + 500*2) / 1e6.
var spendReq = Request{Messages: []Message{{Role: "user", Content: strings.Repeat("x", 296)}}, MaxTokens: 500}

const spendWorstCase = (100*1 + 500*2) / 1e6

func spendClient(t *testing.T, url string, timeout time.Duration) *HTTPClient {
	t.Helper()
	c, _ := mustNew(t, Config{BaseURL: url, Model: "m", APIKey: testSecret, PriceInPerMTok: 1, PriceOutPerMTok: 2, Timeout: timeout})
	return c
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }

func hang(_ http.ResponseWriter, r *http.Request) {
	select {
	case <-r.Context().Done():
	case <-time.After(500 * time.Millisecond):
	}
}

// A timed-out attempt may have been billed: each one charges the worst case.
func TestComplete_TimeoutsChargeTheWorstCaseEach(t *testing.T) {
	srv := gateway(t, hang)
	res, err := spendClient(t, srv.URL, 50*time.Millisecond).Complete(context.Background(), spendReq)
	if !errors.Is(err, context.DeadlineExceeded) || !res.CostEstimated || !near(res.CostUSD, 3*spendWorstCase) {
		t.Fatalf("%+v %v", res, err)
	}
}

// Earlier billed attempts are added to the answer that finally came back.
func TestComplete_ATimeoutBeforeASuccessIsCharged(t *testing.T) {
	var n atomic.Int32
	srv := gateway(t, func(w http.ResponseWriter, r *http.Request) {
		if n.Add(1) == 1 {
			hang(w, r)
			return
		}
		writeCompletion(w, "{}", 10, 20, 0.1, "g")
	})
	res, err := spendClient(t, srv.URL, 50*time.Millisecond).Complete(context.Background(), spendReq)
	if err != nil || res.Attempts != 2 || !res.CostEstimated || !near(res.CostUSD, 0.1+spendWorstCase) ||
		res.Usage != (Usage{PromptTokens: 10, CompletionTokens: 20}) {
		t.Fatalf("%+v %v", res, err)
	}
}

// A 200 the client cannot read (a broken envelope, a body cut at the read cap) was billed.
func TestComplete_AnUnreadable200ChargesTheWorstCase(t *testing.T) {
	for name, body := range map[string]string{
		"broken envelope": `{"choices": [`,
		"no choices":      `{"id":"g","choices":[]}`,
		"cut at the cap":  `{"id":"g","choices":[{"message":{"content":"` + strings.Repeat("y", maxBodyRead) + `"}}]}`,
	} {
		srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, body) })
		res, err := spendClient(t, srv.URL, 0).Complete(context.Background(), spendReq)
		if !errors.Is(err, errMalformed) || !res.CostEstimated || !near(res.CostUSD, spendWorstCase) {
			t.Fatalf("%s: %+v %v", name, res, err)
		}
	}
}

// A refused request (an error status) is not billed.
func TestComplete_AnErrorStatusChargesNothing(t *testing.T) {
	srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusServiceUnavailable) })
	res, err := spendClient(t, srv.URL, 0).Complete(context.Background(), spendReq)
	if err == nil || res.CostUSD != 0 || res.CostEstimated {
		t.Fatalf("%+v %v", res, err)
	}
}

// An all-zero usage block is no usage at all: the worst case, never $0.
func TestComplete_AnAllZeroUsageBlockChargesTheWorstCase(t *testing.T) {
	for name, usage := range map[string]string{
		"empty":         `{}`,
		"zero tokens":   `{"prompt_tokens":0,"completion_tokens":0}`,
		"zero and cost": `{"prompt_tokens":0,"completion_tokens":0,"cost":0}`,
	} {
		srv := gateway(t, func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, `{"id":"g","choices":[{"message":{"content":"{}"}}],"usage":%s}`, usage)
		})
		res, err := spendClient(t, srv.URL, 0).Complete(context.Background(), spendReq)
		if err != nil || !res.CostEstimated || !near(res.CostUSD, spendWorstCase) ||
			res.Usage != (Usage{PromptTokens: 100, CompletionTokens: 500}) {
			t.Fatalf("%s: %+v %v", name, res, err)
		}
	}
}
