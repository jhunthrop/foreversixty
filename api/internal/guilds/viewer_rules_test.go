// api/internal/guilds/viewer_rules_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// Task L6.12 fix round 2: what the guild centre shows depends on who asks. The Progression
// tab counts the guild's private nights for a verified member only, and every other viewer
// (signed out, another account, an unverified row) gets the public guild page's numbers. The
// Raids tab never lists a flagged report to anyone but a verified member who owns it. Home
// and Loot open to a verified row only.

// seedViewerCounts gives a guild a public Onyxia wipe, a private Onyxia kill and a flagged
// public Onyxia kill, and in Barrow Deeps a public, a private and a flagged trash pull.
func seedViewerCounts(t *testing.T, h *httpHarness) int64 {
	t.Helper()
	gid := seedGuild(t, h.pool, "Viewer Counts")
	now := time.Now()
	for _, n := range []l68Night{
		{id: "vc-public-wipe", visibility: "public", at: now.Add(-96 * time.Hour)},
		{id: "vc-private-kill", visibility: "private", at: now.Add(-72 * time.Hour)},
		{id: "vc-flagged-kill", visibility: "public", at: now.Add(-48 * time.Hour), flagged: "tampered"},
		{id: "vc-public-trash", visibility: "public", at: now.Add(-95 * time.Hour), trash: true},
		{id: "vc-private-trash", visibility: "private", at: now.Add(-71 * time.Hour), trash: true},
		{id: "vc-flagged-trash", visibility: "public", at: now.Add(-47 * time.Hour), trash: true, flagged: "tampered"},
	} {
		seedL68Night(t, h.pool, gid, n)
	}
	ctx := context.Background()
	for _, q := range []string{
		`update fights set kill = false where report_id = 'vc-public-wipe'`,
		`update reports set zone = 'Barrow Deeps' where id like 'vc-%-trash'`,
	} {
		if _, err := h.pool.Exec(ctx, q); err != nil {
			t.Fatal(err)
		}
	}
	return gid
}

// seedViewer adds an account with one row in the guild, verified or not, and recomputes its
// membership the way production does.
func seedViewer(t *testing.T, h *httpHarness, gid int64, name string, verified bool) int64 {
	t.Helper()
	uid := seedUser(t, h.pool, name+"@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/"+name, "member", verified)
	recomputeMembership(t, h.pool, gid, uid)
	return uid
}

func (h *httpHarness) progressionAs(t *testing.T, gid int64, actor auth.Actor) (ProgressionView, string) {
	t.Helper()
	h.actor = actor
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/progression", gid), "")
	cache := res.Header.Get("Cache-Control")
	var view ProgressionView
	h.data(res, &view)
	return view, cache
}

func nightIDs(ep EncounterProgress) []string {
	out := []string{}
	for _, n := range ep.PullsByNight {
		out = append(out, n.ReportID)
	}
	return out
}

func barrowDeeps(view ProgressionView) UnnamedZoneProgress {
	for _, u := range view.Unnamed {
		if u.Zone == zoneBarrowDeeps {
			return u
		}
	}
	return UnnamedZoneProgress{}
}

// Anyone but a verified member reads the public numbers (the public wipe alone, never the
// private kill or the flagged one), in a shareable response; a verified member reads the
// private kill too, never the flagged one, in a private response.
func TestProgressionCountsFollowTheViewer(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedViewerCounts(t, h)
	member := seedViewer(t, h, gid, "vcmember", true)
	unverified := seedViewer(t, h, gid, "vcunverified", false)
	stranger := seedUser(t, h.pool, "vc-stranger@example.com")

	for _, v := range []struct {
		who   string
		actor auth.Actor
	}{
		{"signed out", auth.Actor{}},
		{"another account", auth.Actor{UserID: stranger, Role: "user", Method: "session"}},
		{"an unverified row", auth.Actor{UserID: unverified, Role: "user", Method: "session"}},
	} {
		view, cache := h.progressionAs(t, gid, v.actor)
		ony := view.Encounters[0]
		if ony.Pulls != 1 || ony.Kills != 0 || ony.FirstKillAt != nil || ony.BestKillMS != nil ||
			len(ony.PullsByNight) != 1 || ony.PullsByNight[0].ReportID != "vc-public-wipe" {
			t.Errorf("%s: Onyxia = pulls %d, kills %d, first kill %v, best kill %v, nights %v; "+
				"want the public wipe alone", v.who, ony.Pulls, ony.Kills, ony.FirstKillAt, deref(ony.BestKillMS), nightIDs(ony))
		}
		if view.Tier.Down != 0 {
			t.Errorf("%s: tier down = %d, want 0 (the only kills are private or flagged)", v.who, view.Tier.Down)
		}
		if bd := barrowDeeps(view); bd.Pulls != 1 || bd.Nights != 1 {
			t.Errorf("%s: Barrow Deeps = %+v, want the public night's one pull", v.who, bd)
		}
		if cache != "public, max-age=30, stale-while-revalidate=60" {
			t.Errorf("%s: Cache-Control = %q, want the public window", v.who, cache)
		}
	}

	view, cache := h.progressionAs(t, gid, auth.Actor{UserID: member, Role: "user", Method: "session"})
	ony := view.Encounters[0]
	nights := nightIDs(ony)
	if ony.Pulls != 2 || ony.Kills != 1 || ony.FirstKillAt == nil || len(nights) != 2 ||
		nights[0] != "vc-public-wipe" || nights[1] != "vc-private-kill" {
		t.Errorf("verified member: Onyxia = pulls %d, kills %d, first kill %v, nights %v; "+
			"want the public wipe and the private kill, not the flagged kill", ony.Pulls, ony.Kills, ony.FirstKillAt, nights)
	}
	if view.Tier.Down != 1 {
		t.Errorf("verified member: tier down = %d, want 1 (the private kill)", view.Tier.Down)
	}
	if bd := barrowDeeps(view); bd.Pulls != 2 || bd.Nights != 2 {
		t.Errorf("verified member: Barrow Deeps = %+v, want the public and private nights' pulls", bd)
	}
	if cache != "private, no-cache" {
		t.Errorf("verified member: Cache-Control = %q, want private, no-cache", cache)
	}
}

// raidIDsAs lists the report ids the Raids tab shows actor, and checks each listed row
// carries its own one fight.
func (h *httpHarness) raidIDsAs(t *testing.T, gid int64, actor auth.Actor) map[string]bool {
	t.Helper()
	h.actor = actor
	var page RaidsPage
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), ""), &page)
	out := map[string]bool{}
	for _, row := range page.Rows {
		out[row.ID] = true
		if len(row.Fights) != 1 {
			t.Errorf("row %s fights = %+v, want its own one fight", row.ID, row.Fights)
		}
	}
	return out
}

// The Raids tab never lists a flagged report, row or fights, to a non-member: not signed
// out, not to the account that uploaded it, not to an unverified row. A verified member sees
// their own flagged report (they manage it) and no one else's.
func TestRaidsShowAFlaggedReportOnlyToAMemberWhoOwnsIt(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Viewer Raids")
	officer := seedViewer(t, h, gid, "vrofficer", true)
	member := seedViewer(t, h, gid, "vrmember", true)
	unverified := seedViewer(t, h, gid, "vrunverified", false)
	outsider := seedUser(t, h.pool, "vr-outsider@example.com")
	now := time.Now()
	for i, n := range []l68Night{
		{id: "vr-clean", visibility: "public", owner: officer},
		{id: "vr-guild", visibility: "guild", owner: officer},
		{id: "vr-flagged", visibility: "public", owner: officer, flagged: "tampered"},
		{id: "vr-guild-flagged", visibility: "guild", owner: officer, flagged: "unverified"},
		{id: "vr-member-flagged", visibility: "guild", owner: member, flagged: "tampered"},
		{id: "vr-outsider-flagged", visibility: "public", owner: outsider, flagged: "tampered"},
		{id: "vr-outsider-private", visibility: "private", owner: outsider},
	} {
		n.at = now.Add(-time.Duration(i+1) * time.Hour)
		seedL68Night(t, h.pool, gid, n)
	}

	for _, v := range []struct {
		who   string
		actor auth.Actor
		want  []string
	}{
		{"signed out", auth.Actor{}, []string{"vr-clean"}},
		{"the outsider who uploaded a flagged report", auth.Actor{UserID: outsider, Role: "user", Method: "session"},
			[]string{"vr-clean", "vr-outsider-private"}},
		{"an unverified row", auth.Actor{UserID: unverified, Role: "user", Method: "session"}, []string{"vr-clean"}},
		{"a verified member", auth.Actor{UserID: member, Role: "user", Method: "session"},
			[]string{"vr-clean", "vr-guild", "vr-member-flagged"}},
	} {
		got := h.raidIDsAs(t, gid, v.actor)
		match := len(got) == len(v.want)
		for _, id := range v.want {
			match = match && got[id]
		}
		if !match {
			t.Errorf("%s: raids = %v, want %v", v.who, got, v.want)
		}
	}
}

// Home and Loot open to a verified row only: an account whose only row in the guild is
// unverified gets 403 on both, a verified one 200.
func TestHomeAndLootRequireAVerifiedRow(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Viewer Gates")
	unverified := seedViewer(t, h, gid, "vgunverified", false)
	verified := seedViewer(t, h, gid, "vgverified", true)
	for _, v := range []struct {
		who  string
		user int64
		want int
	}{
		{"an unverified row", unverified, http.StatusForbidden},
		{"a verified row", verified, http.StatusOK},
	} {
		h.actor = auth.Actor{UserID: v.user, Role: "user", Method: "session"}
		for _, tab := range []string{"home", "loot"} {
			res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/%s", gid, tab), "")
			res.Body.Close()
			if res.StatusCode != v.want {
				t.Errorf("%s reading %s = %d, want %d", v.who, tab, res.StatusCode, v.want)
			}
		}
	}
}
