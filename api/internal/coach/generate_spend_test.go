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

// An account whose model budget is spent gets the rule-only review, stored with the reason,
// retryable once its cap lifts, and no call is made.
func TestGenerate_AccountSpendCapStoresTheRuleOnlyReview(t *testing.T) {
	f := newGenFixture(t)
	ctx := context.Background()
	fake := &llm.Fake{}
	g := f.gen(t, fake)
	if err := g.Caps.RecordSpend(ctx, f.user, llm.Response{CostUSD: g.Caps.Account.DailyUSD}); err != nil {
		t.Fatal(err)
	}
	r, err := g.Generate(ctx, t28Key(), examplePack(), fixturePersona(t), f.user)
	if err != nil || r.Source != SourceRules || r.Reason != string(RefusalAccountCap) || !r.Retryable || len(fake.Calls) != 0 {
		t.Fatalf("%+v %v", r, err)
	}
	stored, ok, err := f.store.Review(ctx, t28Key())
	if err != nil || !ok || stored.Reason != string(RefusalAccountCap) {
		t.Fatalf("stored %+v %v %v", stored, ok, err)
	}
}

// A call's cost is charged to the account that opened the review.
func TestGenerate_ChargesTheOpeningAccount(t *testing.T) {
	f := newGenFixture(t)
	ctx := context.Background()
	fake := &llm.Fake{Responses: []llm.Response{{CostUSD: 0.0005}}, Errs: []error{errors.New("timed out")}}
	if _, err := f.gen(t, fake).Generate(ctx, t28Key(), examplePack(), fixturePersona(t), f.user); err != nil {
		t.Fatal(err)
	}
	var usd float64
	if err := f.pool.QueryRow(ctx, `select usd::float8 from coach_account_day where user_id = $1`, f.user).Scan(&usd); err != nil ||
		math.Abs(usd-0.0005) > 1e-9 {
		t.Fatalf("account usd %v %v", usd, err)
	}
}
