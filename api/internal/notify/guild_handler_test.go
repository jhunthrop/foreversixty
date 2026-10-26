package notify

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestGuildTestPostsTheGuildEmbed(t *testing.T) {
	g := newGuildAPI(t, officerWithPlan())
	g.saveGuildHook(connectedHook)
	if status, raw, _ := g.call("POST", g.route("/test"), ""); status != http.StatusNoContent {
		t.Fatalf("status %d: %s", status, raw)
	}
	posts := g.poster.Posts()
	if len(posts) != 1 || posts[0].URL != connectedHook {
		t.Fatalf("posts = %+v", posts)
	}
	e := posts[0].Embed
	if e.Title != "Forever Sixty will post raids for Olympus XXVII here." ||
		e.URL != "https://foreversixty.gg"+guildPath+"#raids" || e.Footer != "Olympus XXVII · Forever Sixty" ||
		e.Color != embedColor || e.Description != "" || len(e.Fields) != 0 {
		t.Fatalf("embed = %+v", e)
	}
}

func TestGuildTestFailureDoesNotCountAndSuccessSwitchesBackOn(t *testing.T) {
	ctx := context.Background()
	g := newGuildAPI(t, officerWithPlan())
	g.saveGuildHook(connectedHook)
	for i := 0; i < MaxFailures; i++ {
		if _, err := g.svc.GuildStore.RecordFailure(ctx, g.gid, "boom", true); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := g.svc.Store.Pool.Exec(ctx, `update guild_notification_webhooks set failures = 2 where guild_id = $1`, g.gid); err != nil {
		t.Fatal(err)
	}
	g.poster.PostErr = &DiscordError{Status: http.StatusNotFound, Body: "Unknown Webhook"}
	status, raw, _ := g.call("POST", g.route("/test"), "")
	if status != http.StatusBadGateway || !strings.Contains(raw, "Unknown Webhook") {
		t.Fatalf("failed test: %d %s", status, raw)
	}
	if w, _ := g.guildStored(); w.Failures != 2 {
		t.Fatalf("a manual test counted: failures = %d", w.Failures)
	}
	g.poster.PostErr = nil
	if status, _, _ := g.call("POST", g.route("/test"), ""); status != http.StatusNoContent {
		t.Fatalf("status %d", status)
	}
	if w, _ := g.guildStored(); w.Failures != 0 || w.DisabledReason != nil || w.LastError != nil {
		t.Fatalf("success did not switch back on: %+v", w)
	}
}

func TestGuildTestGates(t *testing.T) {
	member := GuildStanding{Name: "Olympus XXVII", Path: guildPath}
	noPlan, frozen := officerWithPlan(), officerWithPlan()
	noPlan.Plan, frozen.Frozen = false, true
	cases := []struct {
		name     string
		standing GuildStanding
		save     bool
		status   int
		code     string
	}{
		{"member", member, true, http.StatusNotFound, "not_found"},
		{"no plan", noPlan, true, http.StatusForbidden, "guild_plan_required"},
		{"frozen", frozen, true, http.StatusConflict, "claim_contested"},
		{"nothing saved", officerWithPlan(), false, http.StatusNotFound, "not_found"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := newGuildAPI(t, tc.standing)
			if tc.save {
				g.saveGuildHook(connectedHook)
			}
			status, raw, _ := g.call("POST", g.route("/test"), "")
			if status != tc.status || !strings.Contains(raw, `"`+tc.code+`"`) {
				t.Fatalf("status %d: %s", status, raw)
			}
			if len(g.poster.Posts()) != 0 {
				t.Fatal("a refused test posted")
			}
		})
	}
}

func TestGuildRemoveDeletesWithoutThePlan(t *testing.T) {
	standing := officerWithPlan()
	standing.Plan = false
	g := newGuildAPI(t, standing)
	g.saveGuildHook(connectedHook)
	if status, raw, _ := g.call("DELETE", g.route(""), ""); status != http.StatusNoContent {
		t.Fatalf("status %d: %s", status, raw)
	}
	if _, ok := g.guildStored(); ok {
		t.Fatal("the channel survived remove")
	}
	if deletes := g.poster.Deletes(); len(deletes) != 1 || deletes[0] != connectedHook {
		t.Fatalf("Discord deletes = %v", deletes)
	}
	// Removing again is fine, and a Discord-side failure never fails the remove.
	g.saveGuildHook(connectedHook)
	g.poster.DeleteErr = &DiscordError{Status: http.StatusInternalServerError}
	if status, _, _ := g.call("DELETE", g.route(""), ""); status != http.StatusNoContent {
		t.Fatalf("status %d with Discord failing", status)
	}
	if status, _, _ := g.call("DELETE", g.route(""), ""); status != http.StatusNoContent {
		t.Fatalf("status %d removing nothing", status)
	}
}

func TestGuildRemoveIsHiddenFromMembers(t *testing.T) {
	g := newGuildAPI(t, GuildStanding{Name: "Olympus XXVII", Path: guildPath})
	g.saveGuildHook(connectedHook)
	if status, _, _ := g.call("DELETE", g.route(""), ""); status != http.StatusNotFound {
		t.Fatalf("status %d, want 404", status)
	}
	if _, ok := g.guildStored(); !ok {
		t.Fatal("a member removed the channel")
	}
}

func TestGuildTestEmbedEscapesTheName(t *testing.T) {
	e := GuildTestEmbed("https://foreversixty.gg", "  Raid *Nite*  ", guildPath, fixedFinish)
	if e.Title != `Forever Sixty will post raids for Raid \*Nite\* here.` || e.Footer != "Raid *Nite* · Forever Sixty" {
		t.Fatalf("embed = %+v", e)
	}
	if e := GuildTestEmbed("not-a-url", "G", guildPath, fixedFinish); e.URL != "" {
		t.Fatalf("relative base gave a link %q", e.URL)
	}
}
