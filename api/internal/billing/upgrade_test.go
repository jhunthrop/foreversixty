// api/internal/billing/upgrade_test.go
package billing

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// stripeGuildPlan seeds a payer guild on lookupKey and registers its
// subscription with the fake gateway, as the webhook would have stored it.
func (h *harness) stripeGuildPlan(t *testing.T, name, subID, lookupKey string) (guildID, leaderID int64) {
	t.Helper()
	guildID, leaderID = h.payerGuild(t, name)
	price, _ := catalog.Lookup(lookupKey)
	periodEnd := time.Now().Add(20 * 24 * time.Hour)
	h.seedPlan(t, guildID, guildPlan{subID: subID, billedBy: leaderID, tier: price.Offer.Tier(), key: lookupKey,
		cents: price.AmountCents, periodEnd: periodEnd})
	sub := fixtureSubscription(subID, "guild", leaderID, guildID, stripe.SubscriptionStatusActive, periodEnd)
	sub.Items.Data[0].Price = &stripe.Price{ID: "price_" + lookupKey, LookupKey: lookupKey, UnitAmount: price.AmountCents}
	h.gateway.Subscriptions = map[string]*stripe.Subscription{subID: sub}
	return guildID, leaderID
}

func (h *harness) post(t *testing.T, path string, userID int64, body any) (int, []byte) {
	t.Helper()
	w := h.do(t, h.sessionRequest(t, http.MethodPost, path, userID, body))
	return w.Code, w.Body.Bytes()
}

func errorCode(t *testing.T, body []byte) string {
	t.Helper()
	var env struct {
		Error *struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &env); err != nil || env.Error == nil {
		t.Fatalf("no error envelope: %s", body)
	}
	return env.Error.Code
}

func TestUpgradePreview_GuildOnly(t *testing.T) {
	h := newHarness(t)
	renews := time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
	priced := time.Date(2026, 11, 2, 10, 0, 0, 0, time.UTC)
	h.gateway.NextPreview = PriceChangePreview{AmountDueCents: 600, NextAmountCents: 20000, PeriodEnd: renews, ProrationDate: priced}
	gid, leader := h.stripeGuildPlan(t, "preview", "sub_pv", catalog.LookupKeyGuildFounderYearly)

	w := h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(gid, "/upgrade-preview"), leader, nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var env struct {
		Data UpgradePreview `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	want := UpgradePreview{AmountDueCents: 600, NextAmountCents: 20000, Interval: "yearly",
		RenewsAt: renews.Format(time.RFC3339), ProrationDate: priced.Unix()}
	if env.Data != want {
		t.Fatalf("preview = %+v, want %+v", env.Data, want)
	}
	if len(h.gateway.Previews) != 1 || h.gateway.Previews[0].PriceID != "price_"+catalog.LookupKeyRaidGuildFounderYearly {
		t.Fatalf("previews = %+v", h.gateway.Previews)
	}

	officer := h.member(t, gid, "officer", "officer")
	w = h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(gid, "/upgrade-preview"), officer, nil))
	if w.Code != http.StatusForbidden {
		t.Fatalf("non-payer status = %d", w.Code)
	}

	raid, raidLeader := h.stripeGuildPlan(t, "raid", "sub_raid", catalog.LookupKeyRaidGuildMonthly)
	w = h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(raid, "/upgrade-preview"), raidLeader, nil))
	if w.Code != http.StatusConflict || errorCode(t, w.Body.Bytes()) != "not_upgradable" {
		t.Fatalf("raid guild: status = %d, body = %s", w.Code, w.Body.String())
	}

	granted, grantLeader := h.payerGuild(t, "granted")
	h.seedGrantPlan(t, granted)
	w = h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(granted, "/upgrade-preview"), grantLeader, nil))
	if w.Code != http.StatusConflict || errorCode(t, w.Body.Bytes()) != "not_upgradable" {
		t.Fatalf("grant: status = %d, body = %s", w.Code, w.Body.String())
	}

	none, noneLeader := h.payerGuild(t, "none")
	w = h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(none, "/upgrade-preview"), noneLeader, nil))
	if w.Code != http.StatusForbidden || errorCode(t, w.Body.Bytes()) != "guild_plan_required" {
		t.Fatalf("no plan: status = %d, body = %s", w.Code, w.Body.String())
	}
}

func TestUpgrade_FounderLandsOnFounderRaidGuild(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "upgrade", "sub_up", catalog.LookupKeyGuildFounderYearly)
	priced := time.Now().Add(-2 * time.Minute).Truncate(time.Second)

	status, body := h.post(t, billingPath(gid, "/upgrade"), leader, map[string]any{"proration_date": priced.Unix()})
	if status != http.StatusOK {
		t.Fatalf("status = %d, body = %s", status, body)
	}
	if len(h.gateway.PriceChanges) != 1 {
		t.Fatalf("price changes = %+v", h.gateway.PriceChanges)
	}
	change := h.gateway.PriceChanges[0]
	if change.SubscriptionID != "sub_up" || change.PriceID != "price_"+catalog.LookupKeyRaidGuildFounderYearly ||
		!change.ProrationDate.Equal(priced) {
		t.Fatalf("change = %+v", change)
	}
	read, _ := decodeBillingRead(t, body)
	if read.Billing == nil || read.Billing.Tier != catalog.TierRaidGuild || !read.Billing.Founder ||
		len(read.Seats) != catalog.SeatsRaidGuild {
		t.Fatalf("read = %+v", read)
	}
	if _, tier, key, _ := h.guildRow(t, gid); tier != catalog.TierRaidGuild || key != catalog.LookupKeyRaidGuildFounderYearly {
		t.Fatalf("row tier = %q, key = %q", tier, key)
	}
}

func TestUpgrade_PreviewExpired(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "expired", "sub_ex", catalog.LookupKeyGuildMonthly)
	for name, body := range map[string]any{
		"missing": map[string]any{},
		"stale":   map[string]any{"proration_date": time.Now().Add(-11 * time.Minute).Unix()},
		"future":  map[string]any{"proration_date": time.Now().Add(11 * time.Minute).Unix()},
	} {
		status, resp := h.post(t, billingPath(gid, "/upgrade"), leader, body)
		if status != http.StatusConflict || errorCode(t, resp) != "preview_expired" {
			t.Fatalf("%s: status = %d, body = %s", name, status, resp)
		}
	}
	if len(h.gateway.PriceChanges) != 0 {
		t.Fatalf("price changes = %+v", h.gateway.PriceChanges)
	}
}

func TestUpgrade_DeclinedLeavesGuild(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "declined", "sub_dc", catalog.LookupKeyGuildMonthly)
	h.gateway.ChangePriceErr = errors.New("card_declined")
	status, body := h.post(t, billingPath(gid, "/upgrade"), leader, map[string]any{"proration_date": time.Now().Unix()})
	if status != http.StatusConflict || errorCode(t, body) != "upgrade_failed" {
		t.Fatalf("status = %d, body = %s", status, body)
	}
	if _, tier, key, cents := h.guildRow(t, gid); tier != catalog.TierGuild || key != catalog.LookupKeyGuildMonthly || cents != 1500 {
		t.Fatalf("row tier = %q, key = %q, cents = %d", tier, key, cents)
	}
}

func TestUpgrade_FrozenClaimantRefused(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "frozenup", "sub_fz", catalog.LookupKeyGuildMonthly)
	h.svc.Guilds = fakeFrozenClaimant{frozenUserID: leader}
	status, body := h.post(t, billingPath(gid, "/upgrade"), leader, map[string]any{"proration_date": time.Now().Unix()})
	if status != http.StatusConflict || errorCode(t, body) != "claim_contested" {
		t.Fatalf("status = %d, body = %s", status, body)
	}
}
