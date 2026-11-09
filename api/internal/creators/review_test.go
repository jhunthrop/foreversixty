package creators

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

// The security review's cases (2026-11-08): no invite for a removed creator, a rebind ends the
// earlier account's creator Premium, binding drops the account's own unsettled payments, and a
// refund that arrives before its payment is still booked.
func TestReviewCases(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.creator("grimbrew", "2026-11-08")
	first := h.user(at("2026-11-01T00:00:00Z"))
	second := h.user(at("2026-11-01T00:00:00Z"))
	if err := h.svc.bind(ctx, c.ID, first); err != nil {
		t.Fatal(err)
	}
	url, err := h.svc.Invite(ctx, c.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.AcceptInvite(ctx, url[strings.LastIndex(url, "=")+1:], second); err != nil {
		t.Fatal(err)
	}
	var status string
	h.pool.QueryRow(ctx, `select status from entitlements where user_id = $1 and plan = 'premium'`, first).Scan(&status)
	if status != "canceled" {
		t.Fatalf("the earlier account's creator Premium = %q, want canceled", status)
	}

	// An account that paid as the creator's player before it was bound to the creator.
	player := h.returning(c, at("2026-11-10T10:00:00Z"), "us-normal-self")
	h.gw.invoices["in_self"] = InvoiceMoney{ChargeID: "ch_self", PaidAt: at("2026-11-11T10:00:00Z"), AmountCents: 500, NetCents: 470}
	h.gw.refunds["ch_self"] = []Refund{{ID: "re_early", AmountCents: 100, Created: at("2026-11-12T00:00:00Z")}}
	if err := h.svc.InvoicePaid(ctx, player, "in_self"); err != nil {
		t.Fatal(err)
	}
	var refunds int
	h.pool.QueryRow(ctx, `select count(*) from creator_refunds where refund_id = 're_early'`).Scan(&refunds)
	if refunds != 1 {
		t.Fatal("a refund already on the charge was not booked with the payment")
	}
	if err := h.svc.bind(ctx, c.ID, player); err != nil {
		t.Fatal(err)
	}
	if st := h.statement(c, "2026-11"); st.TotalCents != 0 || len(st.Lines) != 0 {
		t.Fatalf("the creator's own lines survived the bind: %+v", st)
	}

	if err := h.svc.Remove(ctx, c.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.svc.Invite(ctx, c.ID); !errors.Is(err, ErrNotFound) {
		t.Fatalf("an invite for a removed creator = %v", err)
	}
}

// A trial's $0 invoice is answered by the gateway with no amount, and nothing is recorded.
func TestATrialInvoiceRecordsNothing(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	c := h.creator("grimbrew", "2026-11-08")
	uid := h.returning(c, at("2026-11-10T10:00:00Z"), "us-normal-trial")
	if err := h.svc.CheckoutCompleted(ctx, uid, "", at("2026-11-10T10:00:00Z"), "in_trial_zero"); err != nil {
		t.Fatal(err)
	}
	var rows int
	h.pool.QueryRow(ctx, `select count(*) from creator_revenue`).Scan(&rows)
	if rows != 0 {
		t.Fatalf("a $0 invoice was recorded (%d rows)", rows)
	}
}

// An arrival sent without a click time (an older page) counts as the sign-up's own moment.
func TestALinkArrivalWithoutAClickTimeStillCounts(t *testing.T) {
	h := newHarness(t)
	c := h.creator("grimbrew", "2026-11-08")
	uid := h.user(h.now.Add(-10 * time.Minute))
	if ok, err := h.svc.Store.AttributeLink(context.Background(), uid, c.Slug, h.now); err != nil || !ok {
		t.Fatalf("an arrival sent without a click time = %v %v", ok, err)
	}
}
