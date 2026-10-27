// api/internal/billing/checkout_review_test.go
package billing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

func transferBody(gid int64) map[string]any {
	return map[string]any{"plan": "guild", "interval": "yearly", "guild_id": gid, "intent": "transfer"}
}

func TestTransferCheckout_RefusedOffers(t *testing.T) {
	cases := []struct {
		name  string
		offer func(h *harness, gid, from, to, other int64)
	}{
		{"offer to another officer", func(h *harness, gid, from, _, other int64) { h.seedTransferOffer(t, gid, from, other) }},
		{"no offer (withdrawn)", func(*harness, int64, int64, int64, int64) {}},
		{"declined", func(h *harness, gid, from, to, _ int64) {
			h.seedTransferOffer(t, gid, from, to)
			h.setTransferState(t, gid, "declined")
		}},
		{"accepted", func(h *harness, gid, from, to, _ int64) {
			h.seedTransferOffer(t, gid, from, to)
			h.setTransferState(t, gid, "accepted")
		}},
	}
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := newHarness(t)
			gid, officer := h.officerOf(t, fmt.Sprintf("refused-offer-%d", i))
			prev := h.seedUser(t, fmt.Sprintf("refused-prev-%d@example.com", i))
			other := h.seedUser(t, fmt.Sprintf("refused-other-%d@example.com", i))
			h.seedStripeGuildPlan(t, gid, catalog.LookupKeyGuildYearly, time.Now().Add(100*time.Hour))
			tc.offer(h, gid, prev, officer, other)
			code, p := h.checkout(t, officer, transferBody(gid))
			if code != http.StatusConflict || p != nil {
				t.Fatalf("status = %d, params = %v", code, p)
			}
		})
	}
}

func (h *harness) setTransferState(t *testing.T, guildID int64, state string) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(),
		`update plan_transfers set state = $2, decided_at = now() where guild_id = $1`, guildID, state); err != nil {
		t.Fatal(err)
	}
}

func TestTransferCheckout_StaleOfferRefused(t *testing.T) {
	h := newHarness(t)
	gid, officer := h.officerOf(t, "stale-offer")
	prev := h.seedUser(t, "stale-offer-prev@example.com")
	h.seedStripeGuildPlan(t, gid, catalog.LookupKeyGuildYearly, time.Now().Add(100*time.Hour))
	h.seedTransferOffer(t, gid, prev, officer)
	if _, err := h.pool.Exec(t.Context(),
		`update entitlements set stripe_subscription_id = 'sub_newer' where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	if code, p := h.checkout(t, officer, transferBody(gid)); code != http.StatusConflict || p != nil {
		t.Fatalf("status = %d, params = %v", code, p)
	}
}

func TestTransferCheckout_NeedsALivePlan(t *testing.T) {
	h := newHarness(t)
	gid, officer := h.officerOf(t, "lapsed-transfer")
	prev := h.seedUser(t, "lapsed-prev@example.com")
	h.seedStripeGuildPlan(t, gid, catalog.LookupKeyGuildYearly, time.Now().Add(100*time.Hour))
	h.seedTransferOffer(t, gid, prev, officer)
	if _, err := h.pool.Exec(t.Context(),
		`update entitlements set status = 'canceled' where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	if code, p := h.checkout(t, officer, transferBody(gid)); code != http.StatusConflict || p != nil {
		t.Fatalf("status = %d, params = %v", code, p)
	}
}

func TestTransferCheckout_KeylessRowPricesByItsOwnTier(t *testing.T) {
	h := newHarness(t)
	h.clockAt(catalog.FounderWindowCloses().Add(time.Hour))
	gid, officer := h.officerOf(t, "keyless-transfer")
	prev := h.seedUser(t, "keyless-prev@example.com")
	if _, err := h.pool.Exec(t.Context(),
		`insert into entitlements (guild_id, plan, tier, source, status, stripe_subscription_id, current_period_end)
		 values ($1, 'guild', 'raid_guild', 'stripe', 'active', 'sub_old', now() + interval '100 hours')`, gid); err != nil {
		t.Fatal(err)
	}
	h.seedTransferOffer(t, gid, prev, officer)
	p := mustCheckout(t, h, officer, map[string]any{"plan": "guild", "interval": "yearly", "guild_id": gid, "intent": "transfer"})
	if p.PriceID != "price_raid_guild_yearly" {
		t.Fatalf("price = %s", p.PriceID)
	}
}

func TestCheckout_RejectsUnknownIntent(t *testing.T) {
	h := newHarness(t)
	gid, officer := h.officerOf(t, "bad-intent")
	for _, body := range []map[string]any{
		{"plan": "guild", "interval": "monthly", "guild_id": gid, "intent": "transfr"},
		{"plan": "premium", "interval": "monthly", "intent": "transfer"},
	} {
		if code, _ := h.checkout(t, officer, body); code != http.StatusBadRequest {
			t.Fatalf("%v: status = %d", body, code)
		}
	}
}

func TestComebackCheckout_UsedOrUnknownKeyIs404AndNoIntervalNeeded(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "comeback-edge@example.com")
	future := time.Now().Add(24 * time.Hour)
	used := h.seedComeback(t, uid, catalog.LookupKeyPremiumFounderYearly, future)
	if _, err := h.pool.Exec(t.Context(), `update seat_comebacks set used_at = now() where id = $1`, used); err != nil {
		t.Fatal(err)
	}
	guildKey := h.seedComeback(t, uid, catalog.LookupKeyGuildYearly, future)
	unknown := h.seedComeback(t, uid, "mystery_key", future)
	for name, id := range map[string]int64{"used": used, "guild key": guildKey, "unknown key": unknown} {
		if code, _ := h.checkout(t, uid, map[string]any{"plan": "premium", "comeback_id": id}); code != http.StatusNotFound {
			t.Fatalf("%s: status = %d", name, code)
		}
	}
	good := h.seedComeback(t, uid, catalog.LookupKeyPremiumFounderMonthly, future)
	p := mustCheckout(t, h, uid, map[string]any{"plan": "premium", "comeback_id": good})
	if p.PriceID != "price_premium_founder_monthly" {
		t.Fatalf("price = %s", p.PriceID)
	}
}

func TestGuildCheckout_SizeIsASoftCap(t *testing.T) {
	h := newHarness(t)
	gid, officer := h.officerOf(t, "big-guild")
	for i := 0; i < catalog.GuildTierMemberLimit; i++ {
		h.seedGuildMember(t, gid, h.seedUser(t, fmt.Sprintf("big-guild-member-%d@example.com", i)), "member")
	}
	p := mustCheckout(t, h, officer, map[string]any{"plan": "guild", "interval": "monthly", "guild_id": gid})
	if p.Metadata["plan"] != "guild" {
		t.Fatalf("params = %+v", p)
	}
}

func TestPortal_GuildFamilies(t *testing.T) {
	for _, tc := range []struct{ name, key, want string }{
		{"founder row", catalog.LookupKeyRaidGuildFounderYearly, "raid_guild_founder"},
		{"keyless row", "", "raid_guild"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := newHarness(t)
			gid, officer := h.officerOf(t, "portal-"+tc.want)
			if err := h.svc.Store.SaveCustomerID(t.Context(), officer, "cus_portal_guild"); err != nil {
				t.Fatal(err)
			}
			if _, err := h.pool.Exec(t.Context(),
				`insert into entitlements (guild_id, plan, tier, source, status, billing_user_id, price_lookup_key)
				 values ($1, 'guild', 'raid_guild', 'stripe', 'active', $2, nullif($3, ''))`, gid, officer, tc.key); err != nil {
				t.Fatal(err)
			}
			w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/portal", officer, map[string]any{"guild_id": gid}))
			if w.Code != http.StatusOK {
				t.Fatalf("status = %d", w.Code)
			}
			if got := h.gateway.Portals[0].ConfigurationKey; got != tc.want {
				t.Fatalf("configuration = %q, want %q", got, tc.want)
			}
		})
	}
}
