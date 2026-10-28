package billing

import (
	"net/http"
	"testing"
)

func TestPremiumCheckout_ConflictCarriesSource(t *testing.T) {
	cases := []struct {
		source string
		insert string
	}{
		{"stripe", `insert into entitlements (user_id, plan, source, status, stripe_subscription_id)
			values ($1, 'premium', 'stripe', 'active', 'sub_source_stripe')`},
		{"grant", `insert into entitlements (user_id, plan, source, status)
			values ($1, 'premium', 'grant', 'active')`},
	}
	for _, tc := range cases {
		t.Run(tc.source, func(t *testing.T) {
			h := newHarness(t)
			uid := h.seedUser(t, "source-"+tc.source+"@example.com")
			if _, err := h.pool.Exec(t.Context(), tc.insert, uid); err != nil {
				t.Fatal(err)
			}
			w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", uid,
				map[string]string{"plan": "premium", "interval": "monthly"}))
			if w.Code != http.StatusConflict {
				t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
			}
			fields := decodeEnvelope(t, w).Error.Fields
			if fields["source"] != tc.source || fields["portal_hint"] != "1" {
				t.Fatalf("fields = %+v", fields)
			}
		})
	}
}

func TestPremiumCheckout_SeatOnlyIsNotRefused(t *testing.T) {
	h := newHarness(t)
	gid, _ := h.payerGuild(t, "seat-only")
	uid := h.member(t, gid, "seat-holder", "member")
	w := h.do(t, h.sessionRequest(t, http.MethodPost, "/v1/billing/checkout", uid,
		map[string]string{"plan": "premium", "interval": "monthly"}))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
}
