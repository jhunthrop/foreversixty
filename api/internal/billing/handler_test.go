// api/internal/billing/handler_test.go
package billing

import (
	"context"
	"net/http"
	"testing"
)

// fakeFrozenClaimant lets a test control checkGuildCheckout's
// frozen-claimant branch without needing a real guilds.Store — the one
// piece of guild-claim-contest state this package reads but does not
// own (Ruling A/B).
type fakeFrozenClaimant struct {
	frozenUserID int64 // FrozenClaimant returns true only for this exact userID
}

func (f fakeFrozenClaimant) FrozenClaimant(_ context.Context, _, userID int64) (bool, error) {
	return userID == f.frozenUserID, nil
}

func TestCheckoutPersonalPremiumCreatesACustomerAndSession(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "checkout-premium@example.com")
	r := h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", uid,
		map[string]string{"plan": "premium", "interval": "monthly"})
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	env := decodeEnvelope(t, w)
	data := env.Data.(map[string]any)
	if data["checkout_url"] == "" {
		t.Fatal("checkout_url missing")
	}
	if len(h.gateway.Checkouts) != 1 || h.gateway.Checkouts[0].PriceID != "price_premium_monthly" {
		t.Fatalf("checkouts = %+v", h.gateway.Checkouts)
	}
	if h.gateway.Checkouts[0].Metadata["plan"] != "premium" {
		t.Fatalf("metadata = %+v", h.gateway.Checkouts[0].Metadata)
	}
	if len(h.gateway.Customers) != 1 {
		t.Fatalf("a new customer should have been created, got %d", len(h.gateway.Customers))
	}

	// A second checkout for the same account reuses the Customer.
	r2 := h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", uid,
		map[string]string{"plan": "premium", "interval": "yearly"})
	if w2 := h.do(t, r2); w2.Code != http.StatusOK {
		t.Fatalf("second checkout: %d, %s", w2.Code, w2.Body.String())
	}
	if len(h.gateway.Customers) != 1 {
		t.Fatalf("the customer must be reused, not recreated: %d", len(h.gateway.Customers))
	}
}

func TestCheckoutRejectsAnInvalidPlanOrInterval(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "checkout-invalid@example.com")
	for _, body := range []map[string]string{
		{"plan": "gold", "interval": "monthly"},
		{"plan": "premium", "interval": "daily"},
	} {
		r := h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", uid, body)
		if w := h.do(t, r); w.Code != http.StatusBadRequest {
			t.Fatalf("%+v: status = %d", body, w.Code)
		}
	}
}

func TestCheckoutGuildRequiresGuildIDAndVerifiedOfficerOfAClaimedGuild(t *testing.T) {
	h := newHarness(t)

	// Unclaimed guild: 403.
	unclaimed := h.seedGuild(t, "unclaimed", false)
	officer := h.seedUser(t, "guild-officer@example.com")
	h.seedGuildMember(t, unclaimed, officer, "officer")
	r := h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", officer,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": unclaimed})
	if w := h.do(t, r); w.Code != http.StatusForbidden {
		t.Fatalf("unclaimed guild: status = %d, body = %s", w.Code, w.Body.String())
	}

	// Claimed guild, caller not an officer: 403.
	claimed := h.seedGuild(t, "claimed", true)
	member := h.seedUser(t, "guild-member@example.com")
	h.seedGuildMember(t, claimed, member, "member")
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", member,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": claimed})
	if w := h.do(t, r); w.Code != http.StatusForbidden {
		t.Fatalf("non-officer: status = %d, body = %s", w.Code, w.Body.String())
	}

	// Claimed guild, verified officer: 200, guild_id in metadata.
	h.seedGuildMember(t, claimed, officer, "officer")
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", officer,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": claimed})
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("verified officer: status = %d, body = %s", w.Code, w.Body.String())
	}
	last := h.gateway.Checkouts[len(h.gateway.Checkouts)-1]
	if last.Metadata["plan"] != "guild" || last.Metadata["guild_id"] == "" {
		t.Fatalf("metadata = %+v", last.Metadata)
	}

	// Missing guild_id: 400.
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", officer,
		map[string]any{"plan": "guild", "interval": "monthly"})
	if w := h.do(t, r); w.Code != http.StatusBadRequest {
		t.Fatalf("missing guild_id: status = %d", w.Code)
	}
}

func TestCheckoutGuildConflictsWhenAlreadyOnThePlanUnlessTransfer(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "already-plan", true)
	officer := h.seedUser(t, "already-plan-officer@example.com")
	h.seedGuildMember(t, gid, officer, "leader")
	if _, err := h.pool.Exec(t.Context(),
		`insert into entitlements (guild_id, plan, source, status) values ($1, 'guild', 'stripe', 'active')`, gid); err != nil {
		t.Fatal(err)
	}
	r := h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", officer,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid})
	w := h.do(t, r)
	if w.Code != http.StatusConflict {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	env := decodeEnvelope(t, w)
	if env.Error.Fields["portal_hint"] != "1" {
		t.Fatalf("fields = %+v", env.Error.Fields)
	}

	// intent: transfer bypasses the conflict.
	newOfficer := h.seedUser(t, "already-plan-newofficer@example.com")
	h.seedGuildMember(t, gid, newOfficer, "officer")
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", newOfficer,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid, "intent": "transfer"})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("transfer: status = %d, body = %s", w.Code, w.Body.String())
	}
}

func TestCheckoutGuildRefusesTheDisputedClaimantButNotOtherOfficers(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "frozen-claim", true)
	disputedClaimant := h.seedUser(t, "frozen-claimant@example.com")
	otherOfficer := h.seedUser(t, "frozen-other-officer@example.com")
	h.seedGuildMember(t, gid, disputedClaimant, "leader")
	h.seedGuildMember(t, gid, otherOfficer, "officer")
	h.svc.Guilds = fakeFrozenClaimant{frozenUserID: disputedClaimant}

	r := h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", disputedClaimant,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid})
	w := h.do(t, r)
	if w.Code != http.StatusForbidden {
		t.Fatalf("disputed claimant: status = %d, body = %s", w.Code, w.Body.String())
	}
	env := decodeEnvelope(t, w)
	if env.Error.Code != "forbidden" {
		t.Fatalf("disputed claimant: error code = %q", env.Error.Code)
	}

	// A different verified officer of the same guild is unaffected by
	// the contest — only the disputed claimant is frozen.
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", otherOfficer,
		map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid})
	w = h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("other officer: status = %d, body = %s", w.Code, w.Body.String())
	}
}

func TestPortalReadsTheCallersOwnCustomerAndFailsWithNone(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "portal@example.com")
	r := h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", uid, nil)
	if w := h.do(t, r); w.Code != http.StatusNotFound {
		t.Fatalf("no customer yet: status = %d", w.Code)
	}
	if err := h.svc.Store.SaveCustomerID(t.Context(), uid, "cus_portal_1"); err != nil {
		t.Fatal(err)
	}
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", uid, nil)
	w := h.do(t, r)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if len(h.gateway.Portals) != 1 || h.gateway.Portals[0].CustomerID != "cus_portal_1" {
		t.Fatalf("portals = %+v", h.gateway.Portals)
	}
}

func TestPortalForAGuildRequiresBeingTheBillingContact(t *testing.T) {
	h := newHarness(t)
	gid := h.seedGuild(t, "portal-guild", true)
	billingOfficer := h.seedUser(t, "portal-billing@example.com")
	otherOfficer := h.seedUser(t, "portal-other@example.com")
	h.seedGuildMember(t, gid, billingOfficer, "leader")
	h.seedGuildMember(t, gid, otherOfficer, "officer")
	if err := h.svc.Store.SaveCustomerID(t.Context(), billingOfficer, "cus_guild_billing"); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(t.Context(),
		`insert into entitlements (guild_id, plan, source, status, billing_user_id)
		 values ($1, 'guild', 'stripe', 'active', $2)`, gid, billingOfficer); err != nil {
		t.Fatal(err)
	}
	r := h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", otherOfficer,
		map[string]any{"guild_id": gid})
	if w := h.do(t, r); w.Code != http.StatusNotFound {
		t.Fatalf("non-billing-contact: status = %d", w.Code)
	}
	r = h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", billingOfficer,
		map[string]any{"guild_id": gid})
	if w := h.do(t, r); w.Code != http.StatusOK {
		t.Fatalf("billing contact: status = %d", w.Code)
	}
}

func TestBillingRoutesAnswer503WhenStripeIsNotConfigured(t *testing.T) {
	h := newHarness(t)
	h.svc.Gateway = nil
	uid := h.seedUser(t, "unconfigured@example.com")
	for _, req := range []*http.Request{
		h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", uid, map[string]string{"plan": "premium", "interval": "monthly"}),
		h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", uid, nil),
	} {
		w := h.do(t, req)
		if w.Code != http.StatusServiceUnavailable {
			t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
		}
		env := decodeEnvelope(t, w)
		if env.Error.Code != "billing_unavailable" {
			t.Fatalf("code = %q", env.Error.Code)
		}
	}
}
