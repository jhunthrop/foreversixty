package billing

import (
	"net/http"
	"slices"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

func (h *harness) openCheckout(t *testing.T, gid, uid int64) {
	t.Helper()
	if err := h.svc.Store.RecordPendingCheckout(t.Context(), gid, uid, "cs_open", time.Now().Add(30*time.Minute)); err != nil {
		t.Fatal(err)
	}
}

// While a Checkout for the guild is open, the offer it was made against
// stands: no withdraw, decline or re-offer, and no upgrade.
func TestTransfer_RefusedWhileACheckoutIsOpen(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "window", "sub_win", catalog.LookupKeyGuildMonthly)
	officer := h.member(t, gid, "officer", "officer")
	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": officer})
	h.openCheckout(t, gid, officer)

	for name, req := range map[string]struct {
		path string
		user int64
		body any
	}{
		"withdraw": {"/transfer/withdraw", leader, nil},
		"decline":  {"/transfer/decline", officer, nil},
		"re-offer": {"/transfer", leader, map[string]any{"to_user_id": officer}},
		"upgrade":  {"/upgrade", leader, map[string]any{"proration_date": time.Now().Unix()}},
	} {
		status, body := h.post(t, billingPath(gid, req.path), req.user, req.body)
		if status != http.StatusConflict || errorCode(t, body) != "transfer_in_progress" {
			t.Fatalf("%s: status = %d, body = %s", name, status, body)
		}
	}
	if _, _, state := h.transferRow(t, gid); state != transferOffered {
		t.Fatalf("state = %q", state)
	}
	if len(h.gateway.PriceChanges) != 0 {
		t.Fatalf("price changes = %+v", h.gateway.PriceChanges)
	}
}

// A double-submitted upgrade changes the price once: the second finds
// Stripe's item already on the target inside the lock.
func TestUpgrade_DoubleSubmitChangesThePriceOnce(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "double", "sub_double", catalog.LookupKeyGuildMonthly)
	plan, err := h.svc.Entitlements.GuildBilling(t.Context(), gid)
	if err != nil {
		t.Fatal(err)
	}
	target := upgradeTarget{actor: guildActor{guildID: gid, userID: leader}, plan: plan,
		priceID: "price_" + catalog.LookupKeyRaidGuildMonthly}
	for range 2 {
		if err := h.svc.changeToRaidGuild(t.Context(), target, time.Now()); err != nil {
			t.Fatal(err)
		}
	}
	if len(h.gateway.PriceChanges) != 1 {
		t.Fatalf("price changes = %+v", h.gateway.PriceChanges)
	}
}

// An outgoing subscription that already ended inside the Checkout window
// is not canceled (Stripe would refuse); the transfer completes.
func TestWebhook_TransferFromAnEndedSubscriptionCompletes(t *testing.T) {
	for name, outgoing := range map[string]*stripe.Subscription{
		"missing":  nil,
		"canceled": {ID: "sub_ended_old", Status: stripe.SubscriptionStatusCanceled},
	} {
		t.Run(name, func(t *testing.T) {
			h := newHarness(t)
			oldUID := h.seedUser(t, "ended-old@example.com")
			sub, newUID, gid := guildFixture(h, t, "sub_ended_new", catalog.LookupKeyGuildMonthly, 1500)
			h.offerTransfer(t, gid, oldUID, newUID, "sub_ended_old", transferOffered)
			sub.Metadata["intent"] = "transfer"
			sub.Metadata["from_subscription_id"] = "sub_ended_old"
			if outgoing != nil {
				h.registerSubscription(outgoing)
			}
			h.deliverSubscription(t, sub)
			if state, _ := h.svc.Store.TransferState(t.Context(), gid, "sub_ended_old"); state != transferAccepted {
				t.Fatalf("state = %q", state)
			}
			if len(h.gateway.CanceledAtPeriodEnd) != 0 {
				t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
			}
			if got, _, _, _ := h.guildRow(t, gid); got != sub.ID {
				t.Fatalf("row = %q", got)
			}
		})
	}
}

// The review's window: the sender withdraws while the recipient's Checkout
// is open. The completed subscription must not take the live plan's row,
// and the old plan keeps billing alone.
func TestWebhook_WithdrawnTransferNeverTakesTheLiveRow(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "withdrawn", "sub_live", catalog.LookupKeyGuildMonthly)
	officer := h.member(t, gid, "officer", "officer")
	h.offerTransfer(t, gid, leader, officer, "sub_live", transferDeclined)
	newcomer := fixtureSubscription("sub_late", "guild", officer, gid, stripe.SubscriptionStatusTrialing, time.Now().Add(time.Hour))
	newcomer.Created = time.Now().Unix()
	newcomer.Metadata["intent"] = "transfer"
	newcomer.Metadata["from_subscription_id"] = "sub_live"
	h.deliverSubscription(t, newcomer)
	if got, _, _, _ := h.guildRow(t, gid); got != "sub_live" {
		t.Fatalf("row = %q, want the live plan kept", got)
	}
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{"sub_late"}) {
		t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
	}
}

// After a completed transfer the transfer row moves on (a new offer away
// from the incoming subscription); that subscription's later events still
// write, and nothing reads them as a stray.
func TestWebhook_CompletedTransfersLaterEventsStillWrite(t *testing.T) {
	h := newHarness(t)
	oldUID := h.seedUser(t, "later-old@example.com")
	sub, newUID, gid := guildFixture(h, t, "sub_later_new", catalog.LookupKeyGuildMonthly, 1500)
	h.offerTransfer(t, gid, oldUID, newUID, "sub_later_old", transferOffered)
	sub.Metadata["intent"] = "transfer"
	sub.Metadata["from_subscription_id"] = "sub_later_old"
	h.registerSubscription(fixtureSubscription("sub_later_old", "guild", oldUID, gid, stripe.SubscriptionStatusActive,
		time.Now().Add(time.Hour)))
	h.deliverSubscription(t, sub)
	if err := h.svc.Store.OfferTransfer(t.Context(), gid, newUID, oldUID, sub.ID); err != nil {
		t.Fatal(err)
	}
	sub.Items.Data[0].Price = &stripe.Price{ID: "price_raid", LookupKey: catalog.LookupKeyRaidGuildMonthly, UnitAmount: 2500}
	h.deliverSubscription(t, sub)
	if _, tier, _, _ := h.guildRow(t, gid); tier != catalog.TierRaidGuild {
		t.Fatalf("tier = %q", tier)
	}
	if !slices.Equal(h.gateway.CanceledAtPeriodEnd, []string{"sub_later_old"}) {
		t.Fatalf("canceled = %v", h.gateway.CanceledAtPeriodEnd)
	}
	if n := h.anomalyCount(t, "unexpected_transfer", sub.ID); n != 0 {
		t.Fatalf("unexpected_transfer anomalies = %d", n)
	}
}
