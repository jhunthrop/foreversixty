package guilds

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

// TestRaidRouteHidesFromNonMembers is R1: the raid route answers 404, never 403, to anyone
// who is not a verified member, and GET /raids carries the planned section only for members.
func TestRaidRouteHidesFromNonMembers(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	stranger := seedUser(t, w.h.pool, "stranger@example.com")
	unverified := w.seedRaider(t, "unverified@example.com", "member", raidChar{"us/hardcore/newbie", "Newbie", "hunter", false})

	viewers := []struct {
		name string
		uid  int64
	}{
		{"signed-out visitor", 0},
		{"signed-in stranger", stranger},
		{"unverified member", unverified},
	}
	for _, v := range viewers {
		t.Run(v.name, func(t *testing.T) {
			w.as(v.uid)
			expectError(t, w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""), http.StatusNotFound, "not_found")
			page := rawData(t, w.h.do(http.MethodGet, w.path("/raids"), ""))
			for _, key := range []string{"planned", "zones", "guild_timezone"} {
				if _, ok := page[key]; ok {
					t.Fatalf("GET /raids carries %q for a non-member", key)
				}
			}
			if _, ok := page["rows"]; !ok {
				t.Fatal("GET /raids lost its rows")
			}
		})
	}

	t.Run("verified member", func(t *testing.T) {
		w.as(w.member)
		page := rawData(t, w.h.do(http.MethodGet, w.path("/raids"), ""))
		for _, key := range []string{"planned", "zones", "guild_timezone"} {
			if _, ok := page[key]; !ok {
				t.Fatalf("GET /raids lacks %q for a member", key)
			}
		}
		if got := string(page["guild_timezone"]); got != `"America/Chicago"` {
			t.Fatalf("guild_timezone = %s", got)
		}
		if !strings.Contains(string(page["zones"]), `"key":"raid:onyxias-lair","name":"Onyxia's Lair","size":40`) {
			t.Fatalf("zones = %s", page["zones"])
		}
		if got := w.getRaid(t, raid.ID); got.ID != raid.ID {
			t.Fatalf("member GET = %+v", got)
		}
	})

	t.Run("another guild's raid id", func(t *testing.T) {
		other := seedGuild(t, w.h.pool, "Elsewhere")
		w.as(w.member)
		res := w.h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids/%d", other, raid.ID), "")
		expectError(t, res, http.StatusNotFound, "not_found")
	})
}

// TestPlannedListIsEmptyNotAbsentForAMember: a member of a guild with no raids reads
// planned: [] and guild_timezone: null.
func TestPlannedListIsEmptyNotAbsentForAMember(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member)
	page := rawData(t, w.h.do(http.MethodGet, w.path("/raids"), ""))
	if string(page["planned"]) != "[]" || string(page["guild_timezone"]) != "null" {
		t.Fatalf("planned = %s, guild_timezone = %s", page["planned"], page["guild_timezone"])
	}
}

// TestCreateRaidRepeatKeepsGuildWallTime is R4/R7: four raids at 20:00 Chicago across the
// 14 March 2027 DST start.
func TestCreateRaidRepeatKeepsGuildWallTime(t *testing.T) {
	w := newRaidWorld(t, true)
	raids := w.createRaids(t, raidBody("2027-03-08T02:00:00Z", 3))
	want := []string{"2027-03-08T02:00:00Z", "2027-03-15T01:00:00Z", "2027-03-22T01:00:00Z", "2027-03-29T01:00:00Z"}
	if len(raids) != len(want) {
		t.Fatalf("created %d raids, want %d", len(raids), len(want))
	}
	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		t.Fatal(err)
	}
	for i, r := range raids {
		if got := r.StartsAt.UTC().Format(time.RFC3339); got != want[i] {
			t.Fatalf("raid %d starts %s, want %s", i, got, want[i])
		}
		if wall := r.StartsAt.In(chicago).Format("15:04"); wall != "20:00" {
			t.Fatalf("raid %d wall time %s, want 20:00", i, wall)
		}
		if r.Title != "Onyxia's Lair" || r.ZoneName != "Onyxia's Lair" || r.Size == nil || *r.Size != 40 {
			t.Fatalf("raid %d = %+v", i, r)
		}
		if r.NotesHTML != "<p>Bring <strong>flasks</strong></p>\n" {
			t.Fatalf("notes_html = %q", r.NotesHTML)
		}
		if r.Notes != "Bring **flasks**" {
			t.Fatalf("notes = %q, want the raw markdown", r.Notes)
		}
	}
}

// The 14 March raid is at 01:00Z: 2027-03-07 + 7 days in Chicago crosses into CDT.
func TestCreateRaidRepeatAcrossDSTIsOneHourEarlierInUTC(t *testing.T) {
	w := newRaidWorld(t, true)
	raids := w.createRaids(t, raidBody("2027-03-08T02:00:00Z", 3))
	if d := raids[1].StartsAt.Sub(raids[0].StartsAt); d != 7*24*time.Hour-time.Hour {
		t.Fatalf("gap across DST = %s, want 167h", d)
	}
}

func TestCreateRaidRefusesTheWholeRepeatPastTheCap(t *testing.T) {
	w := newRaidWorld(t, true)
	for i := 0; i < RaidCap-3; i++ {
		w.createRaid(t)
	}
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids"), raidBody(futureStart(), 3)), http.StatusUnprocessableEntity, "raid_cap")
	w.as(w.member)
	var page struct {
		Planned []RaidView `json:"planned"`
	}
	w.h.data(w.h.do(http.MethodGet, w.path("/raids"), ""), &page)
	if len(page.Planned) != RaidCap-3 {
		t.Fatalf("planned = %d, want %d (nothing created)", len(page.Planned), RaidCap-3)
	}
}

func TestCreateRaidValidation(t *testing.T) {
	w := newRaidWorld(t, true)
	start := futureStart()
	past := time.Now().Add(-time.Hour).UTC().Format(time.RFC3339)
	body := func(zone, title, startsAt, tz, notes string, duration, repeat int) string {
		return fmt.Sprintf(`{"zone":%q,"title":%q,"starts_at":%q,"duration_minutes":%d,"notes":%q,"timezone":%q,"repeat_weeks":%d}`,
			zone, title, startsAt, duration, notes, tz, repeat)
	}
	cases := []struct {
		name, body, field string
	}{
		{"unknown zone", body("raid:molten-core", "", start, "America/Chicago", "", 180, 0), "zone"},
		{"long title", body(raidZoneOny, strings.Repeat("x", 81), start, "America/Chicago", "", 180, 0), "title"},
		{"duration off the step", body(raidZoneOny, "", start, "America/Chicago", "", 75, 0), "duration_minutes"},
		{"duration too long", body(raidZoneOny, "", start, "America/Chicago", "", 270, 0), "duration_minutes"},
		{"start in the past", body(raidZoneOny, "", past, "America/Chicago", "", 180, 0), "starts_at"},
		{"start not RFC 3339", body(raidZoneOny, "", "Thursday 8pm", "America/Chicago", "", 180, 0), "starts_at"},
		{"unknown timezone", body(raidZoneOny, "", start, "Mars/Olympus", "", 180, 0), "timezone"},
		{"no timezone", body(raidZoneOny, "", start, "", "", 180, 0), "timezone"},
		{"long notes", body(raidZoneOny, "", start, "America/Chicago", strings.Repeat("n", 2001), 180, 0), "notes"},
		{"repeat of two", body(raidZoneOny, "", start, "America/Chicago", "", 180, 2), "repeat_weeks"},
	}
	w.as(w.officer)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := expectError(t, w.h.do(http.MethodPost, w.path("/raids"), tc.body), http.StatusUnprocessableEntity, "invalid")
			if _, ok := e.Fields[tc.field]; !ok {
				t.Fatalf("fields = %v, want %q", e.Fields, tc.field)
			}
		})
	}
	t.Run("past start message", func(t *testing.T) {
		e := expectError(t, w.h.do(http.MethodPost, w.path("/raids"), cases[4].body), http.StatusUnprocessableEntity, "invalid")
		if e.Fields["starts_at"] != "Pick a time after now." {
			t.Fatalf("starts_at = %q", e.Fields["starts_at"])
		}
	})
}

// TestOfficerRaidWritesNeedThePlan: members are refused first (forbidden), an officer without
// the plan gets guild_plan_required on every write, and still signs up their own character.
func TestOfficerRaidWritesNeedThePlan(t *testing.T) {
	const planMessage = "Planning and running raids are part of the guild plan."
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if _, err := w.h.pool.Exec(t.Context(), `delete from entitlements where guild_id = $1`, w.gid); err != nil {
		t.Fatal(err)
	}
	patch := raidBody(futureStart(), 0)
	writes := []struct{ method, path, body string }{
		{http.MethodPost, w.path("/raids"), raidBody(futureStart(), 0)},
		{http.MethodPatch, w.path("/raids/%d", raid.ID), patch},
		{http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""},
		{http.MethodPost, w.path("/raids/%d/unlock", raid.ID), ""},
		{http.MethodPost, w.path("/raids/%d/cancel", raid.ID), ""},
	}
	for _, wr := range writes {
		t.Run(wr.method+" "+wr.path, func(t *testing.T) {
			w.as(w.member)
			expectError(t, w.h.do(wr.method, wr.path, wr.body), http.StatusForbidden, "forbidden")
			w.as(w.officer)
			e := expectError(t, w.h.do(wr.method, wr.path, wr.body), http.StatusForbidden, "guild_plan_required")
			if e.Message != planMessage {
				t.Fatalf("message = %q", e.Message)
			}
		})
	}
	t.Run("own sign-up still saves", func(t *testing.T) {
		w.as(w.officer)
		res := w.putSignup(raid.ID, officerKey, `{"state":"in"}`)
		if res.StatusCode != http.StatusOK {
			t.Fatalf("own PUT = %d: %s", res.StatusCode, readBody(res))
		}
		res.Body.Close()
	})
	t.Run("someone else's sign-up needs the plan", func(t *testing.T) {
		w.as(w.officer)
		expectError(t, w.putSignup(raid.ID, member2Key, `{"state":"in"}`), http.StatusForbidden, "guild_plan_required")
	})
}

func TestRaidStatusTransitions(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	steps := []struct {
		action, wantStatus string
		wantCode           int
	}{
		{"unlock", "", http.StatusConflict},
		{"lock", RaidLocked, http.StatusOK},
		{"lock", "", http.StatusConflict},
		{"unlock", RaidPlanned, http.StatusOK},
		{"cancel", RaidCancelled, http.StatusOK},
		{"lock", "", http.StatusConflict},
		{"cancel", "", http.StatusConflict},
	}
	for i, st := range steps {
		res := w.h.do(http.MethodPost, w.path("/raids/%d/%s", raid.ID, st.action), "")
		if st.wantCode != http.StatusOK {
			expectError(t, res, st.wantCode, "conflict")
			continue
		}
		var v RaidView
		w.h.data(res, &v)
		if v.Status != st.wantStatus {
			t.Fatalf("step %d %s: status %q, want %q", i, st.action, v.Status, st.wantStatus)
		}
	}
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID+999), ""), http.StatusNotFound, "not_found")
}

func TestPatchRaidRewritesAndSetsTimezone(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	body := fmt.Sprintf(`{"zone":"raid:hyjal-summit","title":"Hyjal push","starts_at":%q,"duration_minutes":240,`+
		`"notes":"","timezone":"America/New_York","post_to_discord":true}`, raid.StartsAt.UTC().Format(time.RFC3339))
	var v RaidView
	w.h.data(w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), body), &v)
	if v.Zone != "raid:hyjal-summit" || v.Title != "Hyjal push" || v.DurationMinutes != 240 || v.Size != nil ||
		!v.Discord.PostDefault || v.Discord.Connected || v.GuildTimezone == nil || *v.GuildTimezone != "America/New_York" {
		t.Fatalf("patched = %+v", v)
	}
	// An unchanged past start is fine on an edit; a new past start is not.
	if _, err := w.h.pool.Exec(t.Context(), `update guild_raids set starts_at = now() - interval '1 hour' where id = $1`, raid.ID); err != nil {
		t.Fatal(err)
	}
	stored := w.getRaid(t, raid.ID).StartsAt.UTC().Format(time.RFC3339Nano)
	same := fmt.Sprintf(`{"zone":%q,"title":"Late notes","starts_at":%q,"duration_minutes":180,"notes":"x","timezone":"America/New_York"}`,
		raidZoneOny, stored)
	w.h.data(w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), same), &v)
	if v.Title != "Late notes" {
		t.Fatalf("title = %q", v.Title)
	}
	moved := strings.Replace(same, stored, time.Now().Add(-2*time.Hour).UTC().Format(time.RFC3339), 1)
	expectError(t, w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), moved), http.StatusUnprocessableEntity, "invalid")
}

func TestOfficerRaidWritesFreezeWhileContested(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	if _, err := w.h.pool.Exec(t.Context(), `update guilds set claim_contested_at = now() where id = $1`, w.gid); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	writes := []struct{ method, path, body string }{
		{http.MethodPost, w.path("/raids"), raidBody(futureStart(), 0)},
		{http.MethodPatch, w.path("/raids/%d", raid.ID), raidBody(futureStart(), 0)},
		{http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""},
		{http.MethodPost, w.path("/raids/%d/cancel", raid.ID), ""},
	}
	for _, wr := range writes {
		expectError(t, w.h.do(wr.method, wr.path, wr.body), http.StatusConflict, "claim_contested")
	}
}

func TestPatchRefusesACancelledRaid(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/raids/%d/cancel", raid.ID), "")
	res.Body.Close()
	expectError(t, w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), raidBody(futureStart(), 0)), http.StatusConflict, "conflict")
}
