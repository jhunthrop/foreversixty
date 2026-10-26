package notify

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestGuildConnectHidesTheRouteFromNonOfficers(t *testing.T) {
	for name, standing := range map[string]GuildStanding{
		"member":     {Name: "Olympus XXVII", Path: guildPath},
		"non-member": {},
	} {
		t.Run(name, func(t *testing.T) {
			g := newGuildAPI(t, standing)
			if res, _ := g.get(g.route("/connect")); res.StatusCode != http.StatusNotFound {
				t.Fatalf("status %d, want 404", res.StatusCode)
			}
		})
	}
}

func TestGuildConnectWithoutThePlanIs403(t *testing.T) {
	standing := officerWithPlan()
	standing.Plan = false
	g := newGuildAPI(t, standing)
	status, raw, _ := g.call("GET", g.route("/connect"), "")
	if status != http.StatusForbidden || !strings.Contains(raw, `"guild_plan_required"`) ||
		!strings.Contains(raw, guildPlanRequiredMessage) {
		t.Fatalf("status %d: %s", status, raw)
	}
}

func TestGuildConnectWritesAStateAndRedirectsToDiscord(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	state := g.startGuildConnect()
	if len(state) < 40 {
		t.Fatalf("state %q is too short to be 32 random bytes", state)
	}
	var userID, guildID int64
	var ttl float64
	if err := g.svc.Store.Pool.QueryRow(context.Background(),
		`select user_id, guild_id, extract(epoch from expires_at - now()) from notify_connect_states where state = $1`,
		state).Scan(&userID, &guildID, &ttl); err != nil {
		t.Fatal(err)
	}
	if userID != g.uid || guildID != g.gid || ttl < 590 || ttl > 600 {
		t.Fatalf("row = user %d guild %d ttl %.0f", userID, guildID, ttl)
	}
}

func TestGuildCallbackSavesTheChannelAndRedirectsConnected(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	var logs bytes.Buffer
	g.svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
	old := "https://discord.com/api/webhooks/999/OLDtoken"
	g.saveGuildHook(old)
	state := g.startGuildConnect()
	_, loc := g.callback("code="+authCode+"&state="+url.QueryEscape(state), nil)
	g.assertGuildOutcome(loc, outcomeConnected)
	for _, secret := range []string{authCode, connectedToken, "OLDtoken", "webhooks"} {
		if strings.Contains(loc, secret) || strings.Contains(logs.String(), secret) {
			t.Fatalf("%q leaked into the Location or the logs", secret)
		}
	}
	w, ok := g.guildStored()
	if !ok || w.ChannelID != testChan || w.DiscordGuildID == nil || *w.DiscordGuildID != connectedGuild ||
		w.ConnectedBy == nil || *w.ConnectedBy != g.uid {
		t.Fatalf("stored = %+v", w)
	}
	if opened, err := g.svc.openGuild(context.Background(), g.gid); err != nil || opened != connectedHook {
		t.Fatalf("sealed url opens to %q (%v)", opened, err)
	}
	if deletes := g.poster.Deletes(); len(deletes) != 1 || deletes[0] != old {
		t.Fatalf("Discord deletes = %v, want the replaced webhook", deletes)
	}
	if _, ok, _ := g.svc.Store.Get(context.Background(), g.uid, KindDiscord); ok {
		t.Fatal("the guild connect saved the officer's own account webhook")
	}
	// The state is used once.
	_, loc = g.callback("code="+authCode+"&state="+url.QueryEscape(state), nil)
	g.assertOutcome(loc, outcomeFailed)
}

func TestGuildCallbackDeniedAtDiscord(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	state := g.startGuildConnect()
	_, loc := g.callback("error=access_denied&state="+url.QueryEscape(state), nil)
	g.assertGuildOutcome(loc, outcomeDenied)
	if _, ok := g.guildStored(); ok {
		t.Fatal("a denied connect stored a channel")
	}
}

func TestGuildCallbackFailures(t *testing.T) {
	cases := map[string]func(g *guildHarness) string{
		"expired state": func(g *guildHarness) string {
			state := g.startGuildConnect()
			if _, err := g.svc.Store.Pool.Exec(context.Background(),
				`update notify_connect_states set expires_at = now() - interval '1 second' where state = $1`, state); err != nil {
				t.Fatal(err)
			}
			return state
		},
		"another user's state": func(g *guildHarness) string {
			other := seedUser(t, g.svc.Store.Pool)
			if err := g.svc.GuildStore.CreateConnectState(context.Background(), "other-state", other, g.gid, time.Minute); err != nil {
				t.Fatal(err)
			}
			return "other-state"
		},
		"no longer an officer": func(g *guildHarness) string {
			state := g.startGuildConnect()
			g.access.set(GuildStanding{Name: "Olympus XXVII", Path: guildPath})
			return state
		},
		"the plan ended": func(g *guildHarness) string {
			state := g.startGuildConnect()
			lapsed := officerWithPlan()
			lapsed.Plan = false
			g.access.set(lapsed)
			return state
		},
		"the exchange failed": func(g *guildHarness) string {
			g.oauth.Err = ErrExchange
			return g.startGuildConnect()
		},
	}
	for name, setup := range cases {
		t.Run(name, func(t *testing.T) {
			g := newGuildAPI(t, officerWithPlan())
			state := setup(g)
			_, loc := g.callback("code="+authCode+"&state="+url.QueryEscape(state), nil)
			g.assertGuildOutcome(loc, outcomeFailed)
			if _, ok := g.guildStored(); ok {
				t.Fatal("a failed connect stored a channel")
			}
		})
	}
}

func TestGuildCallbackUnknownStateIsTheAccountsFailure(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	_, loc := g.callback("code="+authCode+"&state=forged", nil)
	g.assertOutcome(loc, outcomeFailed)
	if len(g.oauth.Codes()) != 0 {
		t.Fatal("a code was exchanged for an unknown state")
	}
}

func TestGuildRoutesAreOffWithoutConnect(t *testing.T) {
	h := newAPIWith(t, nil, func(s *Service) {
		s.Guilds = &fakeGuildAccess{standing: officerWithPlan(), found: true}
		s.GuildStore = &GuildStore{Pool: s.Store.Pool}
	})
	if h.svc.GuildConnectEnabled() {
		t.Fatal("GuildConnectEnabled with no OAuth client")
	}
	if res, _ := h.get("/v1/guilds/1/notifications/discord/connect"); res.StatusCode != http.StatusNotFound {
		t.Fatalf("connect status %d, want 404", res.StatusCode)
	}
}
