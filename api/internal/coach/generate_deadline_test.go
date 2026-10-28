package coach

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// stallingClient answers nothing until the call's context ends, then fails with its error,
// carrying the worst-case cost the real client charges for a timed-out request.
type stallingClient struct{ cost float64 }

func (c stallingClient) Complete(ctx context.Context, _ llm.Request) (llm.Response, error) {
	<-ctx.Done()
	return llm.Response{CostUSD: c.cost, CostEstimated: c.cost > 0}, ctx.Err()
}

// The lease must outlive the whole generation, or a second view queues a second job.
func TestReviewDeadline_IsInsideTheLease(t *testing.T) {
	if ReviewDeadline <= 0 || ReviewDeadline >= ReviewLeaseTTL {
		t.Fatalf("deadline %v, lease %v", ReviewDeadline, ReviewLeaseTTL)
	}
}

// A generation past its deadline stores the rule-only review, final like a second
// validation failure, and still records what the timed-out call may have cost.
func TestGenerate_PastItsDeadlineStoresTheRulesReviewFinal(t *testing.T) {
	f := newGenFixture(t)
	// Everything slow (the name dictionary, the persona) is built before the clock starts.
	g, pack, persona := f.gen(t, stallingClient{cost: 0.002}), examplePack(), fixturePersona(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := g.Generate(ctx, t28Key(), pack, persona, f.user)
	if err != nil || r.Source != SourceRules || r.Reason != reasonTimeout || r.Retryable {
		t.Fatalf("%+v %v", r, err)
	}
	stored, ok, err := f.store.Review(context.Background(), t28Key())
	if err != nil || !ok || stored.Reason != reasonTimeout || stored.Retryable || !stored.final() {
		t.Fatalf("stored after the deadline: %+v %v %v", stored, ok, err)
	}
	var calls int
	if err := f.pool.QueryRow(context.Background(), `select calls from llm_spend`).Scan(&calls); err != nil || calls != 1 {
		t.Fatalf("spend after the deadline: %d %v", calls, err)
	}
}

// The review job runs its generation under the deadline and stores the fallback.
func TestRunReviewJob_AStalledGenerationStoresTheRulesReview(t *testing.T) {
	f := newT30(t)
	f.svc.Generator.LLM = stallingClient{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []string{t30Report, "1", f.furyKey, strconv.FormatInt(f.premium, 10), "holder-stall"}
	if err := RunReviewJob(ctx, f.svc, args); err != nil {
		t.Fatal(err)
	}
	stored, ok, err := f.svc.Store.Review(context.Background(), f.key(t))
	if err != nil || !ok || stored.Source != SourceRules || stored.Reason != reasonTimeout || stored.Retryable {
		t.Fatalf("%+v %v %v", stored, ok, err)
	}
}
