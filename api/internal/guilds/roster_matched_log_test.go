// api/internal/guilds/roster_matched_log_test.go
//
// A member's own log that the guild matcher attributed (reports.matched_guild_id, guild_id null)
// is that member's log on the guild's pages: the roster row and the member profile read its best
// kill, rating and last log. A private one never shows to anyone else, and an attributed log is
// no guild night: reports-mode attendance and log verification keep reading guild_id alone.
package guilds

import (
	"testing"
	"time"
)

// matchedLogWorld is one verified member whose only log is a dungeon boss kill attributed to the
// guild by matched_guild_id, with guild_id null (production's report gxrgvz4yzsi2).
type matchedLogWorld struct {
	h        *httpHarness
	gid, uid int64
	key      string
}

const matchedLogEncounter = 3480 // a leveling dungeon boss: ranked, never a raid night

func newMatchedLogWorld(t *testing.T, visibility string) matchedLogWorld {
	t.Helper()
	h := newHTTPHarness(t)
	w := matchedLogWorld{h: h, key: "us/pvp/pink-shorts"}
	w.gid = seedGuild(t, h.pool, "Censored")
	w.uid = seedUser(t, h.pool, "matched-log@example.com")
	seedCharacter(t, h.pool, w.gid, w.uid, w.key, "member", true)
	seedExport(t, h.pool, w.uid, w.key, "us", "pvp", "Pink Shorts")
	recomputeMembership(t, h.pool, w.gid, w.uid)

	seedL68Night(t, h.pool, w.gid, l68Night{id: "matched-log", owner: w.uid, visibility: visibility,
		at: time.Now().Add(-2 * time.Hour), keys: []string{w.key}, rated: []string{w.key}})
	for _, stmt := range []string{
		`update reports set matched_guild_id = guild_id, guild_id = null, logging_character = $2 where id = $1 and $3 > 0`,
		`update fights set encounter_id = $3 where report_id = $1 and $2 <> ''`,
		`update fight_metrics set encounter_id = $3 where report_id = $1 and player_key = $2`,
	} {
		if _, err := h.pool.Exec(t.Context(), stmt, "matched-log", w.key, matchedLogEncounter); err != nil {
			t.Fatal(err)
		}
	}
	return w
}

func TestRosterRowReadsAMembersMatchedLog(t *testing.T) {
	w := newMatchedLogWorld(t, "public")
	row := rosterRowOf(t, w.h, w.gid, w.uid, w.key)
	if row.BestParse == nil || row.BestParse.ReportID != "matched-log" {
		t.Errorf("best parse = %+v, want the matched log's kill", row.BestParse)
	}
	if row.Rating == nil || row.Rating.Overall != 50 {
		t.Errorf("rating = %+v, want overall 50 from the matched log", row.Rating)
	}
	if row.LastReportAt == nil {
		t.Error("last report at = nil, want the matched log's date")
	}
}

func TestMemberProfileReadsAMembersMatchedLog(t *testing.T) {
	w := newMatchedLogWorld(t, "public")
	profile, err := w.h.store.MemberProfile(t.Context(), w.gid, w.uid, w.key)
	if err != nil {
		t.Fatal(err)
	}
	if profile.Member == nil || profile.Member.BestParse == nil || profile.Member.BestParse.ReportID != "matched-log" {
		t.Fatalf("profile member = %+v, want the matched log's kill as the best parse", profile.Member)
	}
}

func TestAPrivateMatchedLogStaysOffTheGuildPages(t *testing.T) {
	w := newMatchedLogWorld(t, "private")
	viewer := seedUser(t, w.h.pool, "guildmate@example.com")
	seedCharacter(t, w.h.pool, w.gid, viewer, "us/pvp/guildmate", "member", true)
	recomputeMembership(t, w.h.pool, w.gid, viewer)
	row := rosterRowOf(t, w.h, w.gid, viewer, w.key)
	if row.BestParse != nil || row.Rating != nil || row.LastReportAt != nil {
		t.Errorf("row = best %+v rating %+v last %v, want nothing from a private log", row.BestParse, row.Rating, row.LastReportAt)
	}
	profile, err := w.h.store.MemberProfile(t.Context(), w.gid, w.uid, w.key)
	if err != nil {
		t.Fatal(err)
	}
	if profile.Member == nil || profile.Member.BestParse != nil {
		t.Errorf("profile member = %+v, want no best parse from a private log", profile.Member)
	}
}

// An attributed log is the member's, not a guild night: no officer put it in the guild, so the
// reports-mode attendance window and the log verification of guildmates do not count it.
func TestAMatchedLogIsNoGuildNight(t *testing.T) {
	w := newMatchedLogWorld(t, "public")
	row := rosterRowOf(t, w.h, w.gid, w.uid, w.key)
	if row.Attendance.Nights != 0 || row.Attendance.Present != 0 {
		t.Errorf("attendance = %+v, want 0 of 0: a matched log is no guild night", row.Attendance)
	}
	nights, err := logNights(t.Context(), w.h.pool, w.gid, map[string]int64{w.key: w.uid + 1})
	if err != nil {
		t.Fatal(err)
	}
	if nights[w.key] != 0 {
		t.Errorf("log nights = %d, want 0: a matched log verifies nobody", nights[w.key])
	}
}
