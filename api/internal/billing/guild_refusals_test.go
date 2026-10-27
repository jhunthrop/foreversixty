// api/internal/billing/guild_refusals_test.go
package billing

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// refusalFixture is a Guild-tier Stripe plan with its leader (the billing
// user), a second officer the plan is offered to, and a third officer.
type refusalFixture struct {
	h                       *harness
	gid, leader, to, others int64
}

func newRefusalFixture(t *testing.T) refusalFixture {
	t.Helper()
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "refusals", "sub_rf", catalog.LookupKeyGuildMonthly)
	f := refusalFixture{h: h, gid: gid, leader: leader,
		to: h.member(t, gid, "to", "officer"), others: h.member(t, gid, "other", "officer")}
	h.expectNoContent(t, billingPath(gid, "/transfer"), leader, map[string]any{"to_user_id": f.to})
	return f
}

func (f refusalFixture) freshProration() map[string]any {
	return map[string]any{"proration_date": time.Now().Unix()}
}

func TestGuildBillingRoutes_Refusals(t *testing.T) {
	cases := []struct {
		name   string
		setup  func(f refusalFixture) (method, tail string, userID int64, body any)
		raw    string // a literal request body, sent instead of body
		status int
		code   string
	}{
		{name: "decline by an officer the offer is not to", status: http.StatusNotFound, code: "no_transfer_offer",
			setup: func(f refusalFixture) (string, string, int64, any) {
				return http.MethodPost, "/transfer/decline", f.others, nil
			}},
		{name: "withdraw by a non-subscriber", status: http.StatusForbidden, code: "forbidden",
			setup: func(f refusalFixture) (string, string, int64, any) {
				return http.MethodPost, "/transfer/withdraw", f.to, nil
			}},
		{name: "transfer by a frozen claimant", status: http.StatusConflict, code: "claim_contested",
			setup: func(f refusalFixture) (string, string, int64, any) {
				f.h.svc.Guilds = fakeFrozenClaimant{frozenUserID: f.leader}
				return http.MethodPost, "/transfer", f.leader, map[string]any{"to_user_id": f.others}
			}},
		{name: "upgrade replay", status: http.StatusConflict, code: "not_upgradable",
			setup: func(f refusalFixture) (string, string, int64, any) {
				if status, body := f.h.post(t, billingPath(f.gid, "/upgrade"), f.leader, f.freshProration()); status != http.StatusOK {
					t.Fatalf("first upgrade: %d %s", status, body)
				}
				return http.MethodPost, "/upgrade", f.leader, f.freshProration()
			}},
		{name: "upgrade applied but the row write failed", status: http.StatusInternalServerError, code: "internal",
			setup: func(f refusalFixture) (string, string, int64, any) {
				f.h.gateway.Subscriptions = nil
				return http.MethodPost, "/upgrade", f.leader, f.freshProration()
			}},
		{name: "upgrade error Stripe may have applied", status: http.StatusBadGateway, code: "upgrade_unconfirmed",
			setup: func(f refusalFixture) (string, string, int64, any) {
				f.h.gateway.ChangePriceErr = errors.New("read tcp: i/o timeout")
				return http.MethodPost, "/upgrade", f.leader, f.freshProration()
			}},
		{name: "preview gateway error", status: http.StatusConflict, code: "upgrade_failed",
			setup: func(f refusalFixture) (string, string, int64, any) {
				f.h.gateway.Subscriptions = nil
				return http.MethodGet, "/upgrade-preview", f.leader, nil
			}},
		{name: "no seats wiring", status: http.StatusServiceUnavailable, code: "billing_unavailable",
			setup: func(f refusalFixture) (string, string, int64, any) {
				f.h.svc.Seats = nil
				return http.MethodGet, "", f.leader, nil
			}},
		{name: "upgrade with a body that is not JSON", raw: "{not json", status: http.StatusBadRequest, code: "invalid",
			setup: func(f refusalFixture) (string, string, int64, any) {
				return http.MethodPost, "/upgrade", f.leader, nil
			}},
		{name: "upgrade with a bad body from a non-officer", raw: "{not json", status: http.StatusNotFound, code: "not_found",
			setup: func(f refusalFixture) (string, string, int64, any) {
				return http.MethodPost, "/upgrade", f.h.seedUser(t, "outsider@example.com"), nil
			}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := newRefusalFixture(t)
			method, tail, uid, body := tc.setup(f)
			r := f.h.sessionRequest(t, method, billingPath(f.gid, tail), uid, body)
			if tc.raw != "" {
				r = httptest.NewRequest(method, billingPath(f.gid, tail), strings.NewReader(tc.raw))
				r = r.WithContext(auth.WithActor(r.Context(), auth.Actor{UserID: uid, Role: "member", Method: "session"}))
			}
			w := f.h.do(t, r)
			if w.Code != tc.status || errorCode(t, w.Body.Bytes()) != tc.code {
				t.Fatalf("status = %d, body = %s; want %d %s", w.Code, w.Body.String(), tc.status, tc.code)
			}
		})
	}
}

func TestTransfer_FrozenReceiverCanDecline(t *testing.T) {
	f := newRefusalFixture(t)
	f.h.svc.Guilds = fakeFrozenClaimant{frozenUserID: f.to}
	f.h.expectNoContent(t, billingPath(f.gid, "/transfer/decline"), f.to, nil)
}

// A demoted billing user still pays (as on the seat routes): the read,
// the upgrade and the withdraw admit a verified member who is the subscriber.
func TestGuildBilling_DemotedBillingUserAdmitted(t *testing.T) {
	f := newRefusalFixture(t)
	if _, err := f.h.pool.Exec(t.Context(), `update guild_members set rank = 'member' where guild_id = $1 and user_id = $2`,
		f.gid, f.leader); err != nil {
		t.Fatal(err)
	}
	read, _ := f.h.readBilling(t, f.gid, f.leader)
	if read.Billing == nil || !read.Billing.YouAreBillingContact {
		t.Fatalf("read = %+v", read)
	}
	f.h.expectNoContent(t, billingPath(f.gid, "/transfer/withdraw"), f.leader, nil)
	if status, body := f.h.post(t, billingPath(f.gid, "/upgrade"), f.leader, f.freshProration()); status != http.StatusOK {
		t.Fatalf("upgrade: %d %s", status, body)
	}
}

// After an upgrade Stripe applied but our row missed, the preview brings the
// row up to date and refuses to preview a second move.
func TestUpgradePreview_AlreadyOnTargetUpsertsTier(t *testing.T) {
	h := newHarness(t)
	gid, leader := h.stripeGuildPlan(t, "halfdone", "sub_hd", catalog.LookupKeyGuildFounderYearly)
	target := catalog.LookupKeyRaidGuildFounderYearly
	h.gateway.Subscriptions["sub_hd"].Items.Data[0].Price = &stripe.Price{
		ID: "price_" + target, LookupKey: target, UnitAmount: 20000}
	w := h.do(t, h.sessionRequest(t, http.MethodGet, billingPath(gid, "/upgrade-preview"), leader, nil))
	if w.Code != http.StatusConflict || errorCode(t, w.Body.Bytes()) != "not_upgradable" {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if _, tier, key, _ := h.guildRow(t, gid); tier != catalog.TierRaidGuild || key != target {
		t.Fatalf("row tier = %q, key = %q", tier, key)
	}
	if len(h.gateway.Previews) != 0 {
		t.Fatalf("previews = %+v", h.gateway.Previews)
	}
}
