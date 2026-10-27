// api/internal/billing/checkout_launch_test.go
package billing

import (
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

func (h *harness) clockAt(at time.Time) { h.svc.Now = func() time.Time { return at } }

// checkout posts a checkout request and returns the status and the
// session params the gateway received, nil when none was created.
func (h *harness) checkout(t *testing.T, userID int64, body map[string]any) (int, *CheckoutParams) {
	t.Helper()
	before := len(h.gateway.Checkouts)
	w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", userID, body))
	if len(h.gateway.Checkouts) == before {
		return w.Code, nil
	}
	last := h.gateway.Checkouts[len(h.gateway.Checkouts)-1]
	return w.Code, &last
}

func mustCheckout(t *testing.T, h *harness, userID int64, body map[string]any) CheckoutParams {
	t.Helper()
	code, params := h.checkout(t, userID, body)
	if code != http.StatusOK || params == nil {
		t.Fatalf("status = %d, params = %v", code, params)
	}
	return *params
}

func TestPremiumCheckout_FounderKeyWhileOpen(t *testing.T) {
	h := newHarness(t)
	h.clockAt(catalog.FounderWindowCloses().Add(-time.Second))
	p := mustCheckout(t, h, h.seedUser(t, "f-open@example.com"), map[string]any{"plan": "premium", "interval": "yearly"})
	if p.PriceID != "price_premium_founder_yearly" {
		t.Fatalf("price = %s", p.PriceID)
	}
}

func TestPremiumCheckout_ListKeyAfterClose(t *testing.T) {
	h := newHarness(t)
	h.clockAt(catalog.FounderWindowCloses())
	p := mustCheckout(t, h, h.seedUser(t, "f-closed@example.com"), map[string]any{"plan": "premium", "interval": "yearly"})
	if p.PriceID != "price_premium_yearly_launch" {
		t.Fatalf("price = %s", p.PriceID)
	}
}

func TestCheckout_NeverSellsLegacyPremium(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "legacy@example.com")
	for _, at := range []time.Time{catalog.FounderWindowCloses().Add(-time.Hour), catalog.FounderWindowCloses().Add(time.Hour)} {
		h.clockAt(at)
		for _, interval := range []string{"monthly", "yearly"} {
			p := mustCheckout(t, h, uid, map[string]any{"plan": "premium", "interval": interval})
			if p.PriceID == "price_premium_monthly" || p.PriceID == "price_premium_yearly" {
				t.Fatalf("legacy price sold: %s", p.PriceID)
			}
		}
	}
}

func TestCheckout_RejectsUnknownPlanAndComebackOutsidePremium(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "bad-plan@example.com")
	for _, body := range []map[string]any{
		{"plan": "gold", "interval": "monthly"},
		{"plan": "premium", "interval": "weekly"},
		{"plan": "guild", "interval": "monthly", "guild_id": 1, "comeback_id": 1},
	} {
		if code, _ := h.checkout(t, uid, body); code != http.StatusBadRequest {
			t.Fatalf("%v: status = %d", body, code)
		}
	}
}

func TestRaidGuildCheckout_MetadataPlanGuild(t *testing.T) {
	h := newHarness(t)
	h.clockAt(catalog.FounderWindowCloses().Add(-time.Hour))
	gid, officer := h.officerOf(t, "raid-guild")
	p := mustCheckout(t, h, officer, map[string]any{"plan": "raid_guild", "interval": "monthly", "guild_id": gid})
	if p.PriceID != "price_raid_guild_founder_monthly" || p.Metadata["plan"] != "guild" {
		t.Fatalf("params = %+v", p)
	}
}

func TestGuildCheckout_TrialOnce(t *testing.T) {
	h := newHarness(t)
	gid, officer := h.officerOf(t, "trial-once")
	body := map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid}
	p := mustCheckout(t, h, officer, body)
	if p.TrialPeriodDays != catalog.GuildTrialDays || p.TrialEnd != nil {
		t.Fatalf("first checkout trial = %d / %v", p.TrialPeriodDays, p.TrialEnd)
	}
	if _, err := h.pool.Exec(t.Context(), `delete from pending_checkouts`); err != nil {
		t.Fatal(err)
	}
	h.seedGuildTrial(t, gid)
	p = mustCheckout(t, h, officer, body)
	if p.TrialPeriodDays != 0 {
		t.Fatalf("second checkout trial = %d", p.TrialPeriodDays)
	}
}

func TestTransferCheckout_RequiresOffer(t *testing.T) {
	h := newHarness(t)
	gid, officer := h.officerOf(t, "transfer-no-offer")
	h.seedStripeGuildPlan(t, gid, catalog.LookupKeyGuildYearly, time.Now().Add(100*time.Hour))
	code, p := h.checkout(t, officer, map[string]any{"plan": "guild", "interval": "yearly", "guild_id": gid, "intent": "transfer"})
	if code != http.StatusConflict || p != nil {
		t.Fatalf("status = %d, params = %v", code, p)
	}
}

func TestTransferCheckout_KeepsFounderAfterClose(t *testing.T) {
	h := newHarness(t)
	now := catalog.FounderWindowCloses().Add(24 * time.Hour)
	h.clockAt(now)
	gid, officer := h.officerOf(t, "transfer-founder")
	prev := h.seedUser(t, "transfer-founder-prev@example.com")
	h.seedStripeGuildPlan(t, gid, catalog.LookupKeyGuildFounderYearly, now.Add(200*time.Hour))
	h.seedTransferOffer(t, gid, prev, officer)
	p := mustCheckout(t, h, officer, map[string]any{"plan": "guild", "interval": "yearly", "guild_id": gid, "intent": "transfer"})
	if p.PriceID != "price_guild_founder_yearly" || p.TrialPeriodDays != 0 {
		t.Fatalf("params = %+v", p)
	}
	if p.TrialEnd == nil || !p.TrialEnd.Equal(now.Add(200*time.Hour)) {
		t.Fatalf("trial end = %v", p.TrialEnd)
	}
	if p.Metadata["intent"] != "transfer" || p.Metadata["from_subscription_id"] != "sub_old" {
		t.Fatalf("metadata = %+v", p.Metadata)
	}
}

func TestTransferCheckout_TrialEndFloor(t *testing.T) {
	h := newHarness(t)
	now := time.Now().UTC().Truncate(time.Second)
	h.clockAt(now)
	gid, officer := h.officerOf(t, "transfer-floor")
	prev := h.seedUser(t, "transfer-floor-prev@example.com")
	h.seedStripeGuildPlan(t, gid, catalog.LookupKeyGuildFounderMonthly, now.Add(10*time.Hour))
	h.seedTransferOffer(t, gid, prev, officer)
	p := mustCheckout(t, h, officer, map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid, "intent": "transfer"})
	if p.TrialEnd == nil || !p.TrialEnd.Equal(now.Add(49*time.Hour)) {
		t.Fatalf("trial end = %v", p.TrialEnd)
	}
}

func TestComebackCheckout_UsesOldKey(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "comeback@example.com")
	other := h.seedUser(t, "comeback-other@example.com")
	id := h.seedComeback(t, uid, catalog.LookupKeyPremiumFounderYearly, time.Now().Add(24*time.Hour))
	expired := h.seedComeback(t, uid, catalog.LookupKeyPremiumFounderMonthly, time.Now().Add(-time.Hour))

	if code, _ := h.checkout(t, other, map[string]any{"plan": "premium", "interval": "monthly", "comeback_id": id}); code != http.StatusNotFound {
		t.Fatalf("another user's comeback: %d", code)
	}
	if code, _ := h.checkout(t, uid, map[string]any{"plan": "premium", "interval": "monthly", "comeback_id": expired}); code != http.StatusNotFound {
		t.Fatalf("expired comeback: %d", code)
	}
	p := mustCheckout(t, h, uid, map[string]any{"plan": "premium", "interval": "monthly", "comeback_id": id})
	if p.PriceID != "price_premium_founder_yearly" || p.Metadata["comeback_id"] == "" {
		t.Fatalf("params = %+v", p)
	}
}

func TestPortal_FounderConfiguration(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "portal-founder@example.com")
	if err := h.svc.Store.SaveCustomerID(t.Context(), uid, "cus_founder"); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(t.Context(),
		`insert into entitlements (user_id, plan, source, status, stripe_subscription_id, price_lookup_key)
		 values ($1, 'premium', 'stripe', 'active', 'sub_f', 'premium_founder_monthly')`, uid); err != nil {
		t.Fatal(err)
	}
	if w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", uid, nil)); w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := h.gateway.Portals[0].ConfigurationKey; got != "premium_founder" {
		t.Fatalf("configuration = %q", got)
	}
}
