package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"
)

// fakePinger records every ping; posted false or err make it refuse.
type fakePinger struct {
	pings  []RaidPing
	posted bool
	err    error
}

func (f *fakePinger) PingRaid(_ context.Context, _ int64, p RaidPing) (bool, error) {
	f.pings = append(f.pings, p)
	return f.posted, f.err
}

const (
	officerDiscord = "100000000000000001"
	member2Discord = "100000000000000002"
)

// pingWorld is a raid world with a planned raid, Discord linked for the officer and member2
// (the member has none), and a fake pinger that posts.
func pingWorld(t *testing.T) (raidWorld, RaidView, *fakePinger) {
	t.Helper()
	w := newRaidWorld(t, true)
	for uid, id := range map[int64]string{w.officer: officerDiscord, w.member2: member2Discord} {
		if _, err := w.h.pool.Exec(t.Context(),
			`insert into discord_links (user_id, discord_user_id, display_name) values ($1, $2, 'x')`, uid, id); err != nil {
			t.Fatal(err)
		}
	}
	pinger := &fakePinger{posted: true}
	w.h.svc.RaidPings = pinger
	w.h.svc.PublicBaseURL = "https://foreversixty.gg"
	return w, w.createRaid(t), pinger
}

func TestNotSignedRaidersAreAnOfficersOnly(t *testing.T) {
	w, raid, _ := pingWorld(t)
	w.as(w.officer)
	v := w.getRaid(t, raid.ID)
	if v.NotSignedRaiders == nil {
		t.Fatal("officer has no not_signed_raiders")
	}
	got := map[string]bool{}
	for _, r := range *v.NotSignedRaiders {
		got[r.Name] = r.DiscordLinked
		if !strings.HasPrefix(r.AccountKey, accountKeyPrefix) {
			t.Errorf("account_key = %q", r.AccountKey)
		}
	}
	// One row per account: the member's Kraggor stands for Sunderfel too (no main, first by name).
	want := map[string]bool{"Thornhide": true, "Kraggor": false, "Mirelle": true}
	if len(got) != len(want) {
		t.Fatalf("not_signed_raiders = %+v", *v.NotSignedRaiders)
	}
	for name, linked := range want {
		if l, ok := got[name]; !ok || l != linked {
			t.Errorf("%s: linked %v (present %v), want %v", name, l, ok, linked)
		}
	}
	if v.Ping == nil || v.Ping.LastPingedAt != nil {
		t.Fatalf("ping = %+v", v.Ping)
	}

	w.as(w.member)
	data := rawData(t, w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""))
	for _, key := range []string{"not_signed_raiders", "ping"} {
		if _, ok := data[key]; ok {
			t.Errorf("member reads %s", key)
		}
	}
	if _, ok := data["not_signed"]; !ok {
		t.Error("member lost not_signed")
	}
}

func TestPingMentionsLinkedNonSignersOnceEverySixHours(t *testing.T) {
	w, raid, pinger := pingWorld(t)
	w.as(w.member)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), http.StatusForbidden, "forbidden")
	if len(pinger.pings) != 0 {
		t.Fatal("a member's ping posted")
	}

	w.as(w.officer)
	var res PingResult
	w.h.data(w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), &res)
	if !res.Posted || res.Mentioned != 2 || res.NoDiscord != 1 || res.NoDiscordNames[0] != "Kraggor" || res.NextAllowedAt == nil {
		t.Fatalf("ping = %+v", res)
	}
	if len(pinger.pings) != 1 {
		t.Fatalf("pings = %d", len(pinger.pings))
	}
	p := pinger.pings[0]
	if len(p.MentionIDs) != 2 || !strings.Contains(p.Content, "<@"+member2Discord+">") ||
		!strings.Contains(p.Content, "<t:") || !strings.Contains(p.Content, "https://foreversixty.gg/guild/") {
		t.Fatalf("ping message = %+v", p)
	}

	e := expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), http.StatusTooManyRequests, "ping_throttled")
	next, err := time.Parse(time.RFC3339, e.Fields["next_allowed_at"])
	if err != nil || time.Until(next) < 5*time.Hour {
		t.Fatalf("next_allowed_at = %q (%v)", e.Fields["next_allowed_at"], err)
	}
	if len(pinger.pings) != 1 {
		t.Fatal("the throttled ping posted")
	}
	if v := w.getRaid(t, raid.ID); v.Ping == nil || v.Ping.LastPingedAt == nil {
		t.Fatalf("raid ping state = %+v", v.Ping)
	}
}

func TestPingThatDidNotLandGivesTheThrottleBack(t *testing.T) {
	w, raid, pinger := pingWorld(t)
	w.as(w.officer)
	pinger.posted, pinger.err = false, errors.New("discord 500")
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), http.StatusBadGateway, "discord_refused")
	pinger.err = nil
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), http.StatusConflict, "discord_not_connected")
	pinger.posted = true
	var res PingResult
	w.h.data(w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), &res)
	if !res.Posted || len(pinger.pings) != 3 {
		t.Fatalf("ping after refusals = %+v, pings %d", res, len(pinger.pings))
	}
}

func TestPingWithoutTheBotOrAnyoneToMention(t *testing.T) {
	w, raid, _ := pingWorld(t)
	w.as(w.officer)
	w.h.svc.RaidPings = nil
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), http.StatusConflict, "discord_not_connected")

	pinger := &fakePinger{posted: true}
	w.h.svc.RaidPings = pinger
	if _, err := w.h.pool.Exec(t.Context(), `delete from discord_links`); err != nil {
		t.Fatal(err)
	}
	var res PingResult
	w.h.data(w.h.do(http.MethodPost, w.path("/raids/%d/ping", raid.ID), ""), &res)
	if res.Posted || res.Mentioned != 0 || res.NoDiscord != 3 || len(pinger.pings) != 0 {
		t.Fatalf("nobody to mention = %+v, pings %d", res, len(pinger.pings))
	}
	// Nothing posted, so the throttle was not spent.
	if v := w.getRaid(t, raid.ID); v.Ping == nil || v.Ping.LastPingedAt != nil {
		t.Fatalf("ping state = %+v", v.Ping)
	}
}

func TestPingContentAndTargets(t *testing.T) {
	start := time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC)
	got := pingContent("Onyxia's Lair", start, "https://x/raids/1", []string{"1", "2"})
	want := "Still need your answer for **Onyxia's Lair** on <t:1796868000:F> (<t:1796868000:R>): <@1> <@2>\nSign up here: https://x/raids/1"
	if got != want {
		t.Fatalf("content = %q", got)
	}
	if strings.Contains(got, "—") {
		t.Fatal("em dash in the ping")
	}
	res := splitPingTargets([]NotSignedRaider{
		{Name: "A", DiscordLinked: true, discordUserID: "100000000000000001"},
		{Name: "B", DiscordLinked: true, discordUserID: "@everyone"},
		{Name: "C"},
	})
	if res.Mentioned != 1 || res.NoDiscord != 2 {
		t.Fatalf("targets = %+v", res)
	}
	raw, _ := json.Marshal(res)
	if strings.Contains(string(raw), "1000000000") {
		t.Fatalf("the answer leaks a Discord id: %s", raw)
	}
}
