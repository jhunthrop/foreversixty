package coach

import (
	"context"
	"errors"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// A call that failed after attempts the provider may have billed still carries their
// worst-case cost: the spend cap counts it and the stored row says so.
func TestGenerate_AFailedCallsBilledCostIsRecorded(t *testing.T) {
	f := newGenFixture(t)
	ctx := context.Background()
	fake := &llm.Fake{
		Responses: []llm.Response{{CostUSD: 0.004, CostEstimated: true}},
		Errs:      []error{errors.New("timed out")},
	}
	r, err := f.gen(t, fake).Generate(ctx, t28Key(), examplePack(), fixturePersona(t), f.user)
	if err != nil || r.Source != SourceRules || r.Reason != reasonModelError || r.CostUSD != 0.004 || !r.CostEstimated {
		t.Fatalf("%+v %v", r, err)
	}
	var usd float64
	var calls, estimated int
	if err := f.pool.QueryRow(ctx, `select usd, calls, estimated_calls from llm_spend`).Scan(&usd, &calls, &estimated); err != nil ||
		math.Abs(usd-0.004*estimatedCostMargin) > 1e-9 || calls != 1 || estimated != 1 {
		t.Fatalf("llm_spend %v %d %d %v", usd, calls, estimated, err)
	}
}

// A failure that cost nothing (a refused request) records no spend.
func TestGenerate_AFailedCallWithNoCostRecordsNothing(t *testing.T) {
	f := newGenFixture(t)
	ctx := context.Background()
	fake := &llm.Fake{Errs: []error{errors.New("503")}}
	if _, err := f.gen(t, fake).Generate(ctx, t28Key(), examplePack(), fixturePersona(t), f.user); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := f.pool.QueryRow(ctx, `select count(*) from llm_spend`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("llm_spend rows %d %v", n, err)
	}
}
