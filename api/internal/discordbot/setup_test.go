package discordbot

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// An install started from the setup checklist (?return=setup) comes back to the checklist.
func TestInstallFromSetupReturnsToTheChecklist(t *testing.T) {
	h := newInstallHarness(t)
	rec := h.do(http.MethodGet, h.path(h.gid, "/install?return=setup"), "")
	if rec.Code != http.StatusFound {
		t.Fatalf("start = %d: %s", rec.Code, rec.Body.String())
	}
	ctx := context.Background()
	var state, kind string
	if err := h.svc.d.Store.Pool.QueryRow(ctx,
		`select state, kind from notify_connect_states where guild_id = $1`, h.gid).Scan(&state, &kind); err != nil {
		t.Fatal(err)
	}
	if kind != notify.StateKindBotInstallSetup {
		t.Fatalf("kind = %q", kind)
	}
	cs, ok, err := h.svc.d.Notify.GuildStore.TakeConnectState(ctx, state)
	if err != nil || !ok {
		t.Fatalf("take: %v %v", ok, err)
	}
	req := httptest.NewRequest(http.MethodGet, "/v1/notifications/discord/callback?code=c0de", nil)
	req = req.WithContext(auth.WithActor(req.Context(), h.officerActor()))
	out := httptest.NewRecorder()
	h.svc.installCallback(botSetupOutcomeParam)(out, req, cs)
	want := installPublic + olympusPath + "/setup?discord_bot=installed#discord"
	if loc := out.Header().Get("Location"); loc != want {
		t.Fatalf("Location = %q, want %q", loc, want)
	}
}

// The recruiting channel saves from the server's postable channels, clears with "", and the
// settings view carries it.
func TestRecruitChannelSaveAndClear(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	route := h.path(h.gid, "/recruit-channel")
	if rec := h.do(http.MethodPut, route, `{"channel_id":"c2"}`); rec.Code != http.StatusOK {
		t.Fatalf("save = %d: %s", rec.Code, rec.Body.String())
	}
	in, _ := h.installFor(h.gid)
	if in.RecruitChannelID == nil || *in.RecruitChannelID != "c2" || *in.RecruitChannelName != "general" {
		t.Fatalf("install = %+v", in)
	}
	if rec := h.do(http.MethodPut, route, `{"channel_id":"locked"}`); rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("locked channel = %d", rec.Code)
	}
	if rec := h.do(http.MethodPut, route, `{"channel_id":""}`); rec.Code != http.StatusOK {
		t.Fatalf("clear = %d: %s", rec.Code, rec.Body.String())
	}
	if in, _ := h.installFor(h.gid); in.RecruitChannelID != nil {
		t.Fatalf("cleared install = %+v", in)
	}
	h.setStanding(func(s *notify.GuildStanding) { s.Officer = false })
	if rec := h.do(http.MethodPut, route, `{"channel_id":"c2"}`); rec.Code != http.StatusNotFound {
		t.Fatalf("member = %d", rec.Code)
	}
}

// A new application posts in the recruiting channel; without one nothing is sent.
func TestPostApplicationUsesTheRecruitingChannel(t *testing.T) {
	h := newInstallHarness(t)
	h.callback(h.gid, "code=c", h.officerActor())
	ctx := context.Background()
	level, ilvl := 60, 61
	post := guilds.ApplicationPost{GuildName: "Olympus XXVII", Character: "Thornhide", SpecClass: "Restoration Druid",
		Message: "Two tiers of heals.", Availability: "Tue and Thu", ReviewURL: "https://site.test/guild/x/applications",
		Level: &level, ItemLevel: &ilvl}
	if word, err := h.svc.PostApplication(ctx, h.gid, post); err != nil || word != guilds.OutcomeNoBot {
		t.Fatalf("no channel: %q %v", word, err)
	}
	name, id := "recruitment", "c2"
	if err := h.svc.d.Store.SetRecruitChannel(ctx, h.gid, &id, &name); err != nil {
		t.Fatal(err)
	}
	if word, err := h.svc.PostApplication(ctx, h.gid, post); err != nil || word != guilds.OutcomeSent {
		t.Fatalf("post: %q %v", word, err)
	}
	var sent *RESTCall
	for _, c := range h.rest.Calls() {
		if c.Method == "CreateMessage" && c.ChannelID == "c2" {
			sent = &c
		}
	}
	if sent == nil {
		t.Fatalf("calls = %+v", h.rest.Calls())
	}
	e := sent.Message.Embeds[0]
	if e.Title != "Application · Thornhide" || e.Author != "Restoration Druid" || e.Description != post.Message ||
		len(e.Fields) != 3 || e.Footer != "Olympus XXVII · Forever Sixty" {
		t.Fatalf("embed = %+v", e)
	}
	if b := sent.Message.Components[0][0]; b.URL != post.ReviewURL || !strings.Contains(b.Label, "Review") {
		t.Fatalf("button = %+v", b)
	}
}

func TestClipRunes(t *testing.T) {
	if got := clipRunes("abcdef", 4); got != "abc…" {
		t.Fatalf("clip = %q", got)
	}
	if got := clipRunes("abc", 4); got != "abc" {
		t.Fatalf("short = %q", got)
	}
}
