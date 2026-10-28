// api/internal/guilds/subscribe_facts_test.go
package guilds

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// withStepSources wires the billing and plan reads the subscribe step needs, as main does.
func withStepSources(h *httpHarness) *httpHarness {
	h.store.Billing = &billing.Store{Pool: h.pool}
	h.store.Plans = &entitlements.Store{Pool: h.pool}
	return h
}

// subscribeOf is guildID's subscribe view for viewer, failing when it is missing.
func subscribeOf(t *testing.T, h *httpHarness, viewer, guildID int64) auth.GuildSubscribeView {
	t.Helper()
	got, err := h.store.GuildStanding(context.Background(), viewer, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	sub := got[guildID].Subscribe
	if sub == nil {
		t.Fatalf("guild %d: no subscribe view in %+v", guildID, got[guildID])
	}
	return *sub
}

func linkBattleNet(t *testing.T, h *httpHarness, userID int64) {
	t.Helper()
	setGuild(t, h, `update users set bnet_sub = 'bnet-' || id::text where id = $1`, userID)
}

func TestGuildStanding_SubscribeFromBilling(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid := seedGuild(t, h.pool, "Holders")
	leader := seedRanked(t, h, gid, "Holdmaster", "leader", true)
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, gid, leader)

	sub := subscribeOf(t, h, leader, gid)
	if sub.Step != string(StepReady) || sub.TrialUsed == nil || *sub.TrialUsed {
		t.Fatalf("holder before any trial = %+v, want ready with trial_used false", sub)
	}

	setGuild(t, h, `insert into pending_checkouts (guild_id, user_id, stripe_session_id, expires_at)
		values ($1, $2, 'cs_open', now() + interval '1 hour')`, gid, leader)
	if sub := subscribeOf(t, h, leader, gid); sub.Step != string(StepCheckoutOpen) || sub.TrialUsed != nil {
		t.Fatalf("with a pending checkout = %+v, want checkout_open and no trial_used", sub)
	}

	setGuild(t, h, `delete from pending_checkouts where guild_id = $1`, gid)
	setGuild(t, h, `insert into guild_trials (guild_id, stripe_subscription_id) values ($1, 'sub_old')`, gid)
	if sub := subscribeOf(t, h, leader, gid); sub.Step != string(StepReady) || sub.TrialUsed == nil || !*sub.TrialUsed {
		t.Fatalf("with a used trial = %+v, want ready with trial_used true", sub)
	}

	setGuild(t, h, guildTierSQL, gid)
	if sub := subscribeOf(t, h, leader, gid); sub.Step != string(StepOnPlan) {
		t.Fatalf("on a live plan = %+v, want on_plan", sub)
	}
}

func TestGuildStanding_SubscribeRateLimited(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid := seedGuild(t, h.pool, "Waiting")
	leader := seedRanked(t, h, gid, "Patientone", "leader", false)
	linkBattleNet(t, h, leader)

	// A contest attempt never counts against the claim budget.
	setGuild(t, h, `insert into guild_claim_attempts (user_id, attempted_at, kind)
		values ($1, now() - interval '2 days', 'contest')`, leader)
	if sub := subscribeOf(t, h, leader, gid); sub.Step != string(StepClaim) {
		t.Fatalf("after only a contest attempt = %+v, want claim", sub)
	}

	var attempted time.Time
	if err := h.pool.QueryRow(context.Background(), `insert into guild_claim_attempts (user_id, attempted_at, kind)
		values ($1, now() - interval '10 days', 'claim') returning attempted_at`, leader).Scan(&attempted); err != nil {
		t.Fatal(err)
	}
	sub := subscribeOf(t, h, leader, gid)
	want := attempted.Add(30 * 24 * time.Hour).UTC().Format(time.RFC3339)
	if sub.Step != string(StepClaimBlocked) || sub.Reason == nil || *sub.Reason != BlockedRateLimited ||
		sub.AvailableAt == nil || *sub.AvailableAt != want || sub.OtherGuildName != nil {
		t.Fatalf("claimed 10 days ago = %+v (available_at %v), want claim_blocked rate_limited at %s",
			sub, deref(sub.AvailableAt), want)
	}
}

func TestGuildStanding_SubscribeHoldsOther(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	held := seedGuild(t, h.pool, "First Home")
	gid := seedGuild(t, h.pool, "Second Home")
	officer := seedRanked(t, h, gid, "Twohomes", "officer", false)
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, held, officer)

	sub := subscribeOf(t, h, officer, gid)
	if sub.Step != string(StepClaimBlocked) || sub.Reason == nil || *sub.Reason != BlockedHoldsOther ||
		sub.OtherGuildName == nil || *sub.OtherGuildName != "First Home" || sub.AvailableAt != nil {
		t.Fatalf("officer holding another claim = %+v, want claim_blocked holds_other First Home", sub)
	}
}

func TestGuildStanding_OwnerBothGuilds(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	olympus := seedGuild(t, h.pool, "OLYMPUS XXVII")
	second := seedGuild(t, h.pool, "Second Guild")
	owner := seedRanked(t, h, olympus, "Hunthrop", "member", false)
	seedCharacter(t, h.pool, second, owner, "us/hardcore/Hunthropalt", "member", false)
	syncMembership(t, h.pool, second, owner)
	linkBattleNet(t, h, owner)

	for _, gid := range []int64{olympus, second} {
		sub := subscribeOf(t, h, owner, gid)
		if sub.Step != string(StepAskLeader) || sub.TrialUsed != nil || sub.AvailableAt != nil ||
			sub.Reason != nil || sub.OtherGuildName != nil {
			t.Errorf("guild %d: owner = %+v, want a bare ask_leader", gid, sub)
		}
	}
}

func TestGuildStanding_NilBillingFactsLeavesSubscribeOut(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Unwired")
	viewer := seedRanked(t, h, gid, "Plainmember", "member", false)

	got, err := h.store.GuildStanding(context.Background(), viewer, []int64{gid})
	if err != nil {
		t.Fatal(err)
	}
	if got[gid].Subscribe != nil {
		t.Fatalf("Subscribe = %+v with no billing source, want nil", got[gid].Subscribe)
	}
}

// meGuild is one guilds[] entry of GET /v1/me, its plan-backed objects left raw.
type meGuild struct {
	ID        int64           `json:"id"`
	Subscribe json.RawMessage `json:"subscribe"`
	Covered   json.RawMessage `json:"covered"`
	Plan      json.RawMessage `json:"plan"` // absent (nil) unless the viewer is a verified officer of a planned guild
}

// meGuilds serves GET /v1/me for viewer through the real handler, with the guild store as its
// Standing and capability source, and returns its guilds[] entries by guild id.
func meGuilds(t *testing.T, h *httpHarness, viewer int64) map[int64]meGuild {
	t.Helper()
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	svc := &auth.Service{
		Store: &auth.Store{Pool: h.pool}, Entitlements: &entitlements.Store{Pool: h.pool},
		GuildCapabilities: h.store, Standing: h.store, Log: quiet,
	}
	mux := http.NewServeMux()
	auth.Mount(mux, svc, 0)
	actor := auth.Actor{UserID: viewer, Role: "user", Method: "session"}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), actor)))
	}))
	t.Cleanup(srv.Close)
	res, err := http.Get(srv.URL + "/v1/me")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	var env struct {
		Data struct {
			Guilds []meGuild `json:"guilds"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	out := map[int64]meGuild{}
	for _, g := range env.Data.Guilds {
		out[g.ID] = g
	}
	return out
}

func TestMe_SubscribeOwnerAndReadyGuild(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	olympus := seedGuild(t, h.pool, "OLYMPUS XXVII")
	second := seedGuild(t, h.pool, "Second Guild")
	owner := seedRanked(t, h, olympus, "Hunthrop", "member", false)
	seedCharacter(t, h.pool, second, owner, "us/hardcore/Hunthropalt", "member", false)
	syncMembership(t, h.pool, second, owner)

	askLeader := `{"step":"ask_leader","trial_used":null,"available_at":null,"reason":null,"other_guild_name":null}`
	got := meGuilds(t, h, owner)
	for _, gid := range []int64{olympus, second} {
		g := got[gid]
		if string(g.Subscribe) != askLeader || string(g.Covered) != "null" || g.Plan != nil {
			t.Errorf("owner's guild %d = subscribe %s covered %s plan %s, want %s, null and no plan",
				gid, g.Subscribe, g.Covered, g.Plan, askLeader)
		}
	}

	ready := seedGuild(t, h.pool, "Ready Guild")
	holder := seedRanked(t, h, ready, "Readymaster", "leader", true)
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, ready, holder)
	want := `{"step":"ready","trial_used":false,"available_at":null,"reason":null,"other_guild_name":null}`
	if g := meGuilds(t, h, holder)[ready]; string(g.Subscribe) != want || string(g.Covered) != "null" {
		t.Fatalf("holder's ready guild = subscribe %s covered %s, want %s and null", g.Subscribe, g.Covered, want)
	}
}

// planPeriodEnd is a fixed, whole-second period end, so the raw JSON below is a literal.
const planPeriodEnd = "2027-01-06T00:00:00Z"

// seedPaidGuild is a claimed guild on an active Stripe Guild plan paid by its verified
// leader (battletag Payer#1234, character Payer), and a verified officer and member of it.
func seedPaidGuild(t *testing.T, h *httpHarness) (gid, payer, officer, member int64) {
	t.Helper()
	gid = seedGuild(t, h.pool, "Paid Guild")
	payer = seedRanked(t, h, gid, "Payer", "leader", true)
	officer = seedRanked(t, h, gid, "Officerone", "officer", true)
	member = seedRanked(t, h, gid, "Memberone", "member", true)
	setGuild(t, h, `update users set battletag = 'Payer#1234' where id = $1`, payer)
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, gid, payer)
	end, err := time.Parse(time.RFC3339, planPeriodEnd)
	if err != nil {
		t.Fatal(err)
	}
	seedStripePlan(t, h, gid, "guild", "active", &end, false, &payer)
	return gid, payer, officer, member
}

// TestMe_CoveredOnAPaidGuild: through the real handler, a verified member sees the plan that
// covers the guild and who pays it, and gets no officer-only plan object.
func TestMe_CoveredOnAPaidGuild(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid, _, _, member := seedPaidGuild(t, h)
	g := meGuilds(t, h, member)[gid]
	wantCovered := `{"tier":"guild","state":"active","ends_at":"` + planPeriodEnd + `","billed_by":"Payer",` +
		`"you_pay":false,"you_hold_seat":false,"seats_given":0,"seats_total":1}`
	if string(g.Covered) != wantCovered {
		t.Errorf("covered = %s, want %s", g.Covered, wantCovered)
	}
	wantSub := `{"step":"on_plan","trial_used":null,"available_at":null,"reason":null,"other_guild_name":null}`
	if string(g.Subscribe) != wantSub {
		t.Errorf("subscribe = %s, want %s", g.Subscribe, wantSub)
	}
	if g.Plan != nil {
		t.Errorf("plan = %s for a member, want none", g.Plan)
	}
}

// TestGuildStanding_PlanShapeUnchanged: the officer-only plan object keeps its shape beside
// covered, for a verified officer who does not pay.
func TestGuildStanding_PlanShapeUnchanged(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid, _, officer, _ := seedPaidGuild(t, h)
	want := `{"scope":"plan","tier":"guild","you_hold_seat":false,"status":"active",` +
		`"current_period_end":"` + planPeriodEnd + `","cancel_at_period_end":false,` +
		`"billed_by":"Payer#1234","you_are_billing_contact":false}`
	if g := meGuilds(t, h, officer)[gid]; string(g.Plan) != want {
		t.Fatalf("plan = %s, want %s", g.Plan, want)
	}
}
