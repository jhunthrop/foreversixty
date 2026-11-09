// api/internal/billing/trial_test.go
package billing

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

var errMailDown = errors.New("mail down")

// One Premium trial per account: the first checkout trials 7 days, an account that had its trial
// or already paid for Premium checks out at the full price, and a code on it is 409.
func TestPremiumTrialOncePerAccount(t *testing.T) {
	h := newHarness(t)
	withCreators(h)
	first := h.seedUser(t, "first@example.com")
	if p := mustCheckout(t, h, first, map[string]any{"plan": "premium", "interval": "yearly"}); p.TrialPeriodDays != catalog.PremiumTrialDays {
		t.Fatalf("a first Premium trialed %d days, want 7", p.TrialPeriodDays)
	}
	if err := h.svc.Store.RecordPremiumTrial(t.Context(), first, "sub_had_trial"); err != nil {
		t.Fatal(err)
	}
	if p := mustCheckout(t, h, first, map[string]any{"plan": "premium", "interval": "yearly"}); p.TrialPeriodDays != 0 {
		t.Fatalf("a second Premium trialed %d days", p.TrialPeriodDays)
	}
	if code, _ := h.checkout(t, first, map[string]any{"plan": "premium", "interval": "yearly", "creator_code": "GRIMBREW"}); code != http.StatusConflict {
		t.Fatalf("a code after the trial = %d, want 409", code)
	}
	paid := h.seedUser(t, "paid-before@example.com")
	if _, err := h.pool.Exec(t.Context(), `insert into entitlements (user_id, plan, source, status, billing_user_id,
		stripe_subscription_id, current_period_end) values ($1, 'premium', 'stripe', 'canceled', $1, 'sub_old', now())`, paid); err != nil {
		t.Fatal(err)
	}
	if p := mustCheckout(t, h, paid, map[string]any{"plan": "premium", "interval": "monthly"}); p.TrialPeriodDays != 0 {
		t.Fatalf("an account that paid before trialed %d days", p.TrialPeriodDays)
	}
}

// A guild's trial does not spend the officer's Premium trial.
func TestGuildTrialLeavesThePremiumTrial(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "officer-trial@example.com")
	g := h.seedGuild(t, "Trialists", true)
	if err := h.svc.Store.RecordGuildTrial(t.Context(), g, "sub_guild_trial"); err != nil {
		t.Fatal(err)
	}
	if p := mustCheckout(t, h, uid, map[string]any{"plan": "premium", "interval": "monthly"}); p.TrialPeriodDays != catalog.PremiumTrialDays {
		t.Fatalf("Premium after a guild trial trialed %d days", p.TrialPeriodDays)
	}
}

// Founder price through the trial (ruling 0B.2): a trial that ends after the founder window closes
// is sold at the founder price fixed when the checkout started, which every invoice then charges.
func TestTrialStartedInTheFounderWindowKeepsTheFounderPrice(t *testing.T) {
	h := newHarness(t)
	start := catalog.FounderWindowCloses().Add(-2 * 24 * time.Hour)
	h.clockAt(start)
	p := mustCheckout(t, h, h.seedUser(t, "founder-trial@example.com"), map[string]any{"plan": "premium", "interval": "monthly"})
	end := start.Add(time.Duration(p.TrialPeriodDays) * 24 * time.Hour)
	if !end.After(catalog.FounderWindowCloses()) {
		t.Fatalf("the trial ends %s, before the window closes", end)
	}
	if p.PriceID != "price_"+catalog.LookupKeyPremiumFounderMonthly {
		t.Fatalf("price %s, want the founder price", p.PriceID)
	}
	founder, _ := catalog.Lookup(catalog.LookupKeyPremiumFounderMonthly)
	list, _ := catalog.Lookup(catalog.CheckoutLookupKey(catalog.OfferPremium, catalog.Monthly, end))
	if founder.AmountCents >= list.AmountCents || list.Founder {
		t.Fatalf("founder %d vs list %d at the first charge", founder.AmountCents, list.AmountCents)
	}
}

func termsOf(t *testing.T, h *harness, uid int64, body map[string]any) (int, TermsView) {
	t.Helper()
	w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout/terms", uid, body))
	var v TermsView
	if w.Code == http.StatusOK {
		raw, err := json.Marshal(decodeEnvelope(t, w).Data)
		if err == nil {
			err = json.Unmarshal(raw, &v)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	return w.Code, v
}

// The chooser's terms are the checkout's: 7 days, a code's 30, an unknown code's 7, and none once
// the trial is spent, with the end date and the price on offer now.
func TestTermsAnswerTheChooser(t *testing.T) {
	h := newHarness(t)
	withCreators(h)
	now := time.Date(2026, 11, 8, 10, 0, 0, 0, time.UTC)
	h.clockAt(now)
	uid := h.seedUser(t, "chooser@example.com")
	cases := []struct {
		code     string
		days     int64
		ends     string
		codeWord string
	}{
		{"", 7, "2026-11-15T10:00:00Z", ""},
		{"grimbrew", 30, "2026-12-08T10:00:00Z", codeApplied},
		{"NOPE", 7, "2026-11-15T10:00:00Z", codeUnknown},
	}
	for _, c := range cases {
		status, v := termsOf(t, h, uid, map[string]any{"plan": "premium", "interval": "yearly", "creator_code": c.code})
		if status != http.StatusOK || v.TrialDays != c.days || v.EndsAt == nil || *v.EndsAt != c.ends || v.PriceCents != 4000 || !v.Founder ||
			(c.codeWord == "") != (v.Code == nil) || (v.Code != nil && *v.Code != c.codeWord) {
			t.Fatalf("code %q: %d %+v", c.code, status, v)
		}
	}
	if err := h.svc.Store.RecordPremiumTrial(t.Context(), uid, "sub_spent"); err != nil {
		t.Fatal(err)
	}
	_, v := termsOf(t, h, uid, map[string]any{"plan": "premium", "interval": "monthly", "creator_code": "GRIMBREW"})
	if v.TrialDays != 0 || v.EndsAt != nil || v.Code == nil || *v.Code != codeTrialUsed || v.PriceCents != 500 {
		t.Fatalf("spent: %+v", v)
	}
	if status, _ := termsOf(t, h, uid, map[string]any{"plan": "guild", "interval": "yearly"}); status != http.StatusBadRequest {
		t.Fatalf("a guild plan without a guild = %d", status)
	}
	g := h.seedGuild(t, "Chooser guild", true)
	if _, v := termsOf(t, h, uid, map[string]any{"plan": "guild", "interval": "yearly", "guild_id": g}); v.TrialDays != catalog.GuildTrialDays {
		t.Fatalf("guild terms %+v", v)
	}
}
