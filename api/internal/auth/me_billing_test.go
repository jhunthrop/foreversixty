package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// fakeSeats answers SeatFactsSource from a scripted value.
type fakeSeats struct{ facts SeatFacts }

func (f *fakeSeats) SeatFactsForUser(context.Context, int64, time.Time) (SeatFacts, error) {
	return f.facts, nil
}

// meRaw is /v1/me decoded loosely so a test can tell an omitted key from a null one.
func (h *harness) meRaw(t *testing.T, uid int64) map[string]json.RawMessage {
	t.Helper()
	var body map[string]json.RawMessage
	h.decode(t, h.sessionRequest(t, http.MethodGet, "/v1/me", uid), &body)
	return body
}

func field(t *testing.T, raw json.RawMessage) map[string]any {
	t.Helper()
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatalf("decode %s: %v", raw, err)
	}
	return out
}

func (h *harness) billingOf(t *testing.T, uid int64) map[string]any {
	t.Helper()
	ent := field(t, h.meRaw(t, uid)["entitlements"])
	b, _ := ent["billing"].(map[string]any)
	return b
}

func (h *harness) seedStripePremium(t *testing.T, uid int64, key string, cents int64, end time.Time) {
	t.Helper()
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlements.StripeUpsert{
		UserID: &uid, Plan: entitlements.PlanPremium, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_" + key, BillingUserID: uid, Actor: "t",
		PriceLookupKey: key, UnitAmountCents: cents,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestMe_BillingFounderFacts(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "founder@example.com")
	h.seedStripePremium(t, uid, "premium_founder_yearly", 4000, time.Now().Add(24*time.Hour))
	b := h.billingOf(t, uid)
	if b["founder"] != true || b["interval"] != "yearly" || b["amount_cents"] != float64(4000) {
		t.Fatalf("billing = %+v, want founder yearly 4000", b)
	}
}

func TestMe_BillingListPriceFromCatalogWhenRowHasNoAmount(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "list@example.com")
	h.seedStripePremium(t, uid, "premium_monthly_launch", 0, time.Now().Add(24*time.Hour))
	b := h.billingOf(t, uid)
	if b["founder"] != false || b["interval"] != "monthly" || b["amount_cents"] != float64(600) {
		t.Fatalf("billing = %+v, want list monthly 600", b)
	}
}

func TestMe_OwnerGrantStillPremium(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "owner@example.com")
	if err := h.svc.Entitlements.Grant(t.Context(), entitlements.Subject{UserID: &uid},
		entitlements.PlanPremium, nil, nil, "owner", "t", nil); err != nil {
		t.Fatal(err)
	}
	var body struct {
		Entitlements struct {
			ServerSims bool           `json:"server_sims"`
			Billing    map[string]any `json:"billing"`
		} `json:"entitlements"`
	}
	h.decode(t, h.sessionRequest(t, http.MethodGet, "/v1/me", uid), &body)
	b := body.Entitlements.Billing
	if !body.Entitlements.ServerSims || b["plan"] != "premium" {
		t.Fatalf("entitlements = %+v", body.Entitlements)
	}
	if b["founder"] != false || b["interval"] != nil || b["amount_cents"] != nil || b["current_period_end"] != nil {
		t.Fatalf("grant billing = %+v, want null interval and amount", b)
	}
}

func (h *harness) entitlementsOf(t *testing.T, uid int64) map[string]any {
	t.Helper()
	return field(t, h.meRaw(t, uid)["entitlements"])
}

func (h *harness) grantPremium(t *testing.T, uid int64, end *time.Time) {
	t.Helper()
	if err := h.svc.Entitlements.Grant(t.Context(), entitlements.Subject{UserID: &uid},
		entitlements.PlanPremium, end, nil, "owner", "t", nil); err != nil {
		t.Fatal(err)
	}
}

func TestMe_GrantWithoutStripeRow(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "opengrant@example.com")
	h.grantPremium(t, uid, nil)
	ent := h.entitlementsOf(t, uid)
	b, _ := ent["billing"].(map[string]any)
	if b["source"] != "grant" || b["current_period_end"] != nil || ent["personal_portal"] != false {
		t.Fatalf("entitlements = %+v, want grant, null end, no portal", ent)
	}
}

func TestMe_GrantUntilDate(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "datedgrant@example.com")
	end := time.Now().Add(48 * time.Hour).UTC().Truncate(time.Second)
	h.grantPremium(t, uid, &end)
	ent := h.entitlementsOf(t, uid)
	b, _ := ent["billing"].(map[string]any)
	if b["source"] != "grant" || b["current_period_end"] != end.Format(time.RFC3339) || ent["personal_portal"] != false {
		t.Fatalf("entitlements = %+v, want grant ending %s", ent, end.Format(time.RFC3339))
	}
}

func TestMe_StripeSourceAndPortal(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "live@example.com")
	h.seedStripePremium(t, uid, "premium_monthly_launch", 600, time.Now().Add(24*time.Hour))
	ent := h.entitlementsOf(t, uid)
	b, _ := ent["billing"].(map[string]any)
	if b["source"] != "stripe" || ent["personal_portal"] != true {
		t.Fatalf("entitlements = %+v, want stripe with portal", ent)
	}
}

func TestMe_CanceledStripeRowKeepsPortal(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "canceled@example.com")
	h.seedStripePremium(t, uid, "premium_monthly_launch", 600, time.Now().Add(24*time.Hour))
	if _, err := h.svc.Store.Pool.Exec(t.Context(),
		`update entitlements set status = 'canceled' where user_id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	ent := h.entitlementsOf(t, uid)
	b, _ := ent["billing"].(map[string]any)
	if b["source"] != "stripe" || b["status"] != "canceled" || ent["personal_portal"] != true {
		t.Fatalf("entitlements = %+v, want canceled stripe row with portal", ent)
	}
}

func TestMe_SeatKeysOmittedWithoutFacts(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "noseat@example.com")
	h.svc.Seats = &fakeSeats{}
	raw := h.meRaw(t, uid)
	for _, k := range []string{"seat", "seat_offer", "ended_seat"} {
		if _, there := raw[k]; there {
			t.Fatalf("%s must be omitted, got %s", k, raw[k])
		}
	}
}

func TestMe_SeatHolder(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "holder@example.com")
	since := time.Date(2027, 1, 2, 3, 4, 5, 0, time.UTC)
	starts := since.Add(48 * time.Hour)
	h.svc.Seats = &fakeSeats{facts: SeatFacts{Seat: &HeldSeat{
		GuildName: "Olympus XXVII", GivenBy: "Kraggor", Since: since, StartsAt: &starts}}}
	seat := field(t, h.meRaw(t, uid)["seat"])
	if seat["guild_name"] != "Olympus XXVII" || seat["given_by"] != "Kraggor" ||
		seat["since"] != "2027-01-02T03:04:05Z" || seat["starts_at"] != "2027-01-04T03:04:05Z" {
		t.Fatalf("seat = %+v", seat)
	}
}

func TestMe_SeatOfferCarriesOwnEnd(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "offeree@example.com")
	end := time.Now().Add(10 * 24 * time.Hour).UTC().Truncate(time.Second)
	h.seedStripePremium(t, uid, "premium_founder_monthly", 500, end)
	h.svc.Seats = &fakeSeats{facts: SeatFacts{Offer: &SeatOffer{
		GuildID: 7, GuildName: "Olympus XXVII", FromName: "Kraggor", OfferedAt: end.Add(-time.Hour)}}}
	offer := field(t, h.meRaw(t, uid)["seat_offer"])
	if offer["guild_id"] != float64(7) || offer["from_name"] != "Kraggor" ||
		offer["own_ends_at"] != end.Format(time.RFC3339) ||
		offer["own_amount_cents"] != float64(500) || offer["own_interval"] != "monthly" {
		t.Fatalf("offer = %+v", offer)
	}
}

func TestMe_SeatOfferWithoutOwnPlanHasNullOwnFacts(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "offeree2@example.com")
	h.svc.Seats = &fakeSeats{facts: SeatFacts{Offer: &SeatOffer{GuildID: 7, OfferedAt: time.Now()}}}
	offer := field(t, h.meRaw(t, uid)["seat_offer"])
	if v, there := offer["own_ends_at"]; !there || v != nil || offer["own_amount_cents"] != nil || offer["own_interval"] != nil {
		t.Fatalf("offer = %+v, want null own facts", offer)
	}
}

func TestMe_EndedSeatWithin30Days(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "ended@example.com")
	ended := time.Now().Add(-24 * time.Hour).UTC().Truncate(time.Second)
	h.svc.Seats = &fakeSeats{facts: SeatFacts{Ended: &EndedSeat{
		GuildName: "Olympus XXVII", EndedAt: ended, ComebackID: 42, PriceLookupKey: "premium_founder_yearly"}}}
	got := field(t, h.meRaw(t, uid)["ended_seat"])
	if got["guild_name"] != "Olympus XXVII" || got["ended_at"] != ended.Format(time.RFC3339) ||
		got["resubscribe_url"] != "/premium/checkout?plan=premium&interval=yearly&comeback=42" ||
		got["amount_cents"] != float64(4000) || got["interval"] != "yearly" {
		t.Fatalf("ended seat = %+v", got)
	}
}

func TestMe_EndedSeatUnknownKeyLinksToPlans(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "ended2@example.com")
	h.svc.Seats = &fakeSeats{facts: SeatFacts{Ended: &EndedSeat{
		GuildName: "G", EndedAt: time.Now(), ComebackID: 9, PriceLookupKey: "gone"}}}
	got := field(t, h.meRaw(t, uid)["ended_seat"])
	if got["resubscribe_url"] != "/premium" ||
		got["amount_cents"] != nil || got["interval"] != nil {
		t.Fatalf("ended seat = %+v", got)
	}
}

// failingSeats errors on every read.
type failingSeats struct{}

func (failingSeats) SeatFactsForUser(context.Context, int64, time.Time) (SeatFacts, error) {
	return SeatFacts{}, errors.New("seats down")
}

func TestMe_SeatFactsErrorOmitsSeatKeys(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "seatsdown@example.com")
	h.svc.Seats = failingSeats{}
	raw := h.meRaw(t, uid)
	if _, there := raw["user"]; !there {
		t.Fatal("/v1/me must still answer")
	}
	for _, k := range []string{"seat", "seat_offer", "ended_seat"} {
		if _, there := raw[k]; there {
			t.Fatalf("%s must be omitted, got %s", k, raw[k])
		}
	}
}

func TestMe_EndedSeatGoneAfter30Days(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "expired@example.com")
	// The source owns the 30-day window: a comeback past it yields no fact, so no key.
	h.svc.Seats = &fakeSeats{}
	if _, there := h.meRaw(t, uid)["ended_seat"]; there {
		t.Fatal("ended_seat must be omitted once the source has none")
	}
}

func (h *harness) seedGuildPlan(t *testing.T, gid, billing int64, tier, key string) {
	t.Helper()
	end := time.Now().Add(24 * time.Hour)
	if _, err := h.svc.Entitlements.UpsertStripe(t.Context(), entitlements.StripeUpsert{
		GuildID: &gid, Plan: entitlements.PlanGuild, Status: "active", CurrentPeriodEnd: &end,
		StripeSubscriptionID: "sub_g" + key, BillingUserID: billing, Actor: "t",
		Tier: tier, PriceLookupKey: key, UnitAmountCents: 12000,
	}); err != nil {
		t.Fatal(err)
	}
}

type meGuilds struct {
	Guilds []struct {
		ID            int64 `json:"id"`
		VerifiedCount int   `json:"verified_count"`
		Plan          *struct {
			Tier        string `json:"tier"`
			YouHoldSeat bool   `json:"you_hold_seat"`
		} `json:"plan"`
	} `json:"guilds"`
}

func TestMe_GuildPlanTierAndSeat(t *testing.T) {
	h := newHarness(t)
	officer := h.seedUser(t, "officer@example.com")
	gid := h.seedGuild(t, "Olympus XXVII")
	h.seedGuildMember(t, gid, officer, "officer", true)
	h.seedGuildPlan(t, gid, officer, entitlements.TierRaidGuild, "raid_guild_founder_yearly")
	read := func() meGuilds {
		var body meGuilds
		h.decode(t, h.sessionRequest(t, http.MethodGet, "/v1/me", officer), &body)
		return body
	}
	got := read()
	if p := got.Guilds[0].Plan; p == nil || p.Tier != "raid_guild" || p.YouHoldSeat {
		t.Fatalf("plan = %+v, want raid_guild without a seat", p)
	}
	h.svc.Seats = &fakeSeats{facts: SeatFacts{HeldGuildIDs: []int64{gid}}}
	if p := read().Guilds[0].Plan; p == nil || !p.YouHoldSeat {
		t.Fatalf("plan = %+v, want you_hold_seat", p)
	}
}

func TestMe_GuildVerifiedCount(t *testing.T) {
	h := newHarness(t)
	leader := h.seedUser(t, "leader@example.com")
	gid := h.seedGuild(t, "Olympus XXVII")
	h.seedGuildMember(t, gid, leader, "leader", true)
	h.seedGuildMember(t, gid, h.seedUser(t, "a@example.com"), "member", true)
	h.seedGuildMember(t, gid, h.seedUser(t, "b@example.com"), "member", false)
	var body meGuilds
	h.decode(t, h.sessionRequest(t, http.MethodGet, "/v1/me", leader), &body)
	if body.Guilds[0].VerifiedCount != 2 {
		t.Fatalf("verified_count = %d, want 2", body.Guilds[0].VerifiedCount)
	}
}
