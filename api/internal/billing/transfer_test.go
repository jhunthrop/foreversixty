// api/internal/billing/transfer_test.go
package billing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

func (h *harness) transferRow(t *testing.T, guildID int64) (toUserID int64, from, state string) {
	t.Helper()
	if err := h.pool.QueryRow(t.Context(),
		`select to_user_id, from_subscription_id, state from plan_transfers where guild_id = $1`, guildID).
		Scan(&toUserID, &from, &state); err != nil {
		t.Fatal(err)
	}
	return toUserID, from, state
}

func (h *harness) expectNoContent(t *testing.T, path string, userID int64, body any) {
	t.Helper()
	status, resp := h.post(t, path, userID, body)
	if status != http.StatusNoContent || len(resp) != 0 {
		t.Fatalf("%s: status = %d, body = %q", path, status, resp)
	}
}

func TestTransfer_SendWithdrawDecline(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "transfer", "sub_tr", catalog.LookupKeyGuildFounderMonthly)
	officer := h.member(t, gid, "officer", "officer")

	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	if to, from, state := h.transferRow(t, gid); to != officer || from != "sub_tr" || state != transferOffered {
		t.Fatalf("row = %d %q %q", to, from, state)
	}
	read, _ := h.readBilling(t, gid, leader)
	tr := read.Transfer
	if tr == nil || tr.State != transferOffered || tr.ToYou || tr.ToName != fmt.Sprintf("user-%d", officer) ||
		!tr.Founder || tr.Interval == nil || *tr.Interval != "monthly" || tr.AmountCents == nil || *tr.AmountCents != 1200 {
		t.Fatalf("transfer = %+v", tr)
	}
	if len(read.TransferCandidates) != 1 || read.TransferCandidates[0].UserID != officer {
		t.Fatalf("candidates = %+v", read.TransferCandidates)
	}

	h.expectNoContent(t, billingPath(gid, "/transfer/withdraw"), leader, nil)
	if _, _, state := h.transferRow(t, gid); state != transferDeclined {
		t.Fatalf("after withdraw state = %q", state)
	}
	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	h.expectNoContent(t, billingPath(gid, "/transfer/decline"), officer, nil)
	if _, _, state := h.transferRow(t, gid); state != transferDeclined {
		t.Fatalf("after decline state = %q", state)
	}
	if status, _ := h.post(t, billingPath(gid, "/transfer/decline"), officer, nil); status != http.StatusNotFound {
		t.Fatalf("second decline status = %d", status)
	}
	if status, _ := h.post(t, billingPath(gid, "/transfer/withdraw"), leader, nil); status != http.StatusNotFound {
		t.Fatalf("withdraw of nothing status = %d", status)
	}
}

func TestTransfer_TargetMustBeOfficer(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "target", "sub_tg", catalog.LookupKeyGuildMonthly)
	member := h.member(t, gid, "member", "member")
	stranger := h.seedUser(t, "stranger@example.com")
	for _, to := range []int64{member, stranger, leader} {
		status, body := h.post(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": to})
		if status != http.StatusConflict || errorCode(t, body) != "not_member" {
			t.Fatalf("to %d: status = %d, body = %s", to, status, body)
		}
	}
	officer := h.member(t, gid, "officer", "officer")
	other := h.member(t, gid, "other", "officer")
	if status, _ := h.post(t, billingPath(gid, "/transfer"), officer, map[string]any{"to_user_id": other}); status != http.StatusForbidden {
		t.Fatalf("non-subscriber status = %d", status)
	}
	if status, _ := h.post(t, billingPath(gid, "/transfer"), leader, map[string]any{}); status != http.StatusBadRequest {
		t.Fatalf("no target status = %d", status)
	}
}

func TestTransfer_ReceiverSeesToYou(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "receiver", "sub_rc", catalog.LookupKeyGuildMonthly)
	officer := h.member(t, gid, "officer", "officer")
	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	read, raw := h.readBilling(t, gid, officer)
	if read.Transfer == nil || !read.Transfer.ToYou || read.Transfer.FromName != fmt.Sprintf("user-%d", leader) {
		t.Fatalf("transfer = %+v", read.Transfer)
	}
	if string(raw["can"]) != "[]" || string(raw["transfer_candidates"]) != "[]" {
		t.Fatalf("can = %s, candidates = %s", raw["can"], raw["transfer_candidates"])
	}
}

func TestTransfer_NoPlanRefused(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "noplantr")
	officer := h.member(t, gid, "officer", "officer")
	status, body := h.post(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	if status != http.StatusForbidden || errorCode(t, body) != "guild_plan_required" {
		t.Fatalf("status = %d, body = %s", status, body)
	}
}

// An accepted transfer's outgoing subscription runs to its period end; a new
// offer then would overwrite the row retiredByTransfer and reconcile read,
// so it waits until Stripe has ended that subscription, whatever the
// incoming subscription's trial says.
func TestTransfer_AcceptedInLastPeriodRefused(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "inprogress")
	officer := h.member(t, gid, "officer", "officer")
	trialEnded := time.Now().Add(-time.Hour)
	h.seedPlan(t, gid, guildPlan{subID: "sub_new", billedBy: leader, key: catalog.LookupKeyGuildMonthly, cents: 1500,
		trialEnd: &trialEnded})
	h.offerTransfer(t, gid, officer, leader, "sub_prev", transferAccepted)
	outgoing := fixtureSubscription("sub_prev", "guild", officer, gid, stripe.SubscriptionStatusActive, time.Now().Add(time.Hour))
	outgoing.CancelAtPeriodEnd = true
	h.gateway.Subscriptions = map[string]*stripe.Subscription{"sub_prev": outgoing}

	status, body := h.post(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	if status != http.StatusConflict || errorCode(t, body) != "transfer_in_progress" {
		t.Fatalf("status = %d, body = %s", status, body)
	}
	if _, from, state := h.transferRow(t, gid); from != "sub_prev" || state != transferAccepted {
		t.Fatalf("row overwritten: %q %q", from, state)
	}

	outgoing.Status = stripe.SubscriptionStatusCanceled
	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	if _, from, state := h.transferRow(t, gid); from != "sub_new" || state != transferOffered {
		t.Fatalf("row = %q %q", from, state)
	}
}

// The offer's starts_at is the instant the transfer checkout trials to: the
// period end, floored 49 hours ahead.
func TestTransfer_StartsAtAppliesTrialFloor(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.payerGuild(t, "floor")
	officer := h.member(t, gid, "officer", "officer")
	now := time.Now().UTC().Truncate(time.Second)
	h.clockAt(now)
	h.seedPlan(t, gid, guildPlan{subID: "sub_fl", billedBy: leader, key: catalog.LookupKeyGuildMonthly, cents: 1500,
		periodEnd: now.Add(time.Hour)})
	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	read, _ := h.readBilling(t, gid, officer)
	want := now.Add(transferTrialFloor).Format(time.RFC3339)
	if read.Transfer == nil || read.Transfer.StartsAt == nil || *read.Transfer.StartsAt != want {
		t.Fatalf("transfer = %+v, want starts_at %s", read.Transfer, want)
	}
}
