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

// meGuilds serves GET /v1/me for viewer with the guild store as its Standing source and
// returns the raw guilds[] subscribe objects by guild id.
func meGuilds(t *testing.T, h *httpHarness, viewer int64) map[int64]json.RawMessage {
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
			Guilds []struct {
				ID        int64           `json:"id"`
				Subscribe json.RawMessage `json:"subscribe"`
			} `json:"guilds"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	out := map[int64]json.RawMessage{}
	for _, g := range env.Data.Guilds {
		out[g.ID] = g.Subscribe
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
		if string(got[gid]) != askLeader {
			t.Errorf("owner's guild %d subscribe = %s, want %s", gid, got[gid], askLeader)
		}
	}

	ready := seedGuild(t, h.pool, "Ready Guild")
	holder := seedRanked(t, h, ready, "Readymaster", "leader", true)
	setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, ready, holder)
	want := `{"step":"ready","trial_used":false,"available_at":null,"reason":null,"other_guild_name":null}`
	if sub := meGuilds(t, h, holder)[ready]; string(sub) != want {
		t.Fatalf("holder's ready guild subscribe = %s, want %s", sub, want)
	}
}
