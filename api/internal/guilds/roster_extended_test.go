// api/internal/guilds/roster_extended_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// TestHomeSummaryCountsRaidersAndPending checks the Overview summary's own raider/pending
// split (contract's SummaryView) against a guild with one verified and one unverified
// member.
func TestHomeSummaryCountsRaidersAndPending(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "summary-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/summaryofficer", "officer", true)
	seedExport(t, h.pool, officer, "us/hardcore/summaryofficer", "us", "hardcore", "summaryofficer")
	recomputeMembership(t, h.pool, gid, officer)

	pendingUser := seedUser(t, h.pool, "summary-pending@example.com")
	seedCharacter(t, h.pool, gid, pendingUser, "us/hardcore/summarypending", "member", false)
	seedExport(t, h.pool, pendingUser, "us/hardcore/summarypending", "us", "hardcore", "summarypending")
	recomputeMembership(t, h.pool, gid, pendingUser)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)

	if view.Summary.Raiders != 1 {
		t.Errorf("raiders = %d, want 1", view.Summary.Raiders)
	}
	if view.Summary.WaitingForApproval != 1 {
		t.Errorf("waiting_for_approval = %d, want 1", view.Summary.WaitingForApproval)
	}
	if view.Viewer.Role != "officer" {
		t.Errorf("viewer.role = %q, want officer", view.Viewer.Role)
	}
	if view.Viewer.CharacterKey == nil || *view.Viewer.CharacterKey != "us/hardcore/summaryofficer" {
		t.Errorf("viewer.character_key = %v, want us/hardcore/summaryofficer", view.Viewer.CharacterKey)
	}
}

// TestHomePendingIsOfficerOnly checks that the officer-only pending queue lists exactly
// the unverified rows, and reads empty for a plain member.
func TestHomePendingIsOfficerOnly(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "pending-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/pendingofficer", "officer", true)
	seedExport(t, h.pool, officer, "us/hardcore/pendingofficer", "us", "hardcore", "pendingofficer")
	recomputeMembership(t, h.pool, gid, officer)

	member := seedUser(t, h.pool, "pending-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/pendingmember", "member", true)
	seedExport(t, h.pool, member, "us/hardcore/pendingmember", "us", "hardcore", "pendingmember")
	recomputeMembership(t, h.pool, gid, member)

	unverified := seedUser(t, h.pool, "pending-unverified@example.com")
	seedCharacter(t, h.pool, gid, unverified, "us/hardcore/pendingunverified", "member", false)
	seedExport(t, h.pool, unverified, "us/hardcore/pendingunverified", "us", "hardcore", "pendingunverified")
	recomputeMembership(t, h.pool, gid, unverified)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	if len(view.Pending) != 1 || view.Pending[0].CharacterKey != "us/hardcore/pendingunverified" {
		t.Fatalf("officer pending = %+v, want exactly the unverified row", view.Pending)
	}

	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	h.data(res, &view)
	if len(view.Pending) != 0 {
		t.Fatalf("member pending = %+v, want empty", view.Pending)
	}
}

// TestHomeAttendanceCountsPresenceOverLast8Reports seeds 9 reports (so the window drops
// the oldest) and checks a character's attendance present/nights.
func TestHomeAttendanceCountsPresenceOverLast8Reports(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "attendance@example.com")
	key := "us/hardcore/attendance"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "hardcore", "attendance")
	recomputeMembership(t, h.pool, gid, uid)

	// 9 reports, newest first by created_at offset; the character attends every one
	// except the very oldest (which the 8-report window should drop anyway) and the
	// 3rd-newest (a real absence inside the window).
	for i := 0; i < 9; i++ {
		reportID := fmt.Sprintf("attendance-report-%d", i)
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, guild_id, visibility, status, created_at)
			 values ($1, $2, 'guild', 'complete', now() - ($3 || ' hours')::interval)`,
			reportID, gid, fmt.Sprint(i)); err != nil {
			t.Fatal(err)
		}
		players := []string{key}
		if i == 2 { // the 3rd-newest report: this character is absent
			players = []string{}
		}
		if _, err := h.pool.Exec(ctx,
			`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`,
			reportID, players); err != nil {
			t.Fatal(err)
		}
	}

	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	var row *RosterRow
	for i := range view.Roster {
		if view.Roster[i].CharacterKey == key {
			row = &view.Roster[i]
		}
	}
	if row == nil {
		t.Fatal("character missing from roster")
	}
	if row.Attendance.Nights != 8 {
		t.Errorf("nights = %d, want 8 (the window size)", row.Attendance.Nights)
	}
	if row.Attendance.Present != 7 {
		t.Errorf("present = %d, want 7 (8 in-window reports minus the one absence)", row.Attendance.Present)
	}
}

// TestHomeBestParseAndRatingReadTheNewestKillFight checks best_parse and rating are
// populated from fight_metrics/rating_scores for a character with a logged kill.
func TestHomeBestParseAndRatingReadTheNewestKillFight(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "parse-rating@example.com")
	key := "us/hardcore/parserating"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "hardcore", "parserating")
	recomputeMembership(t, h.pool, gid, uid)

	if err := db.EnsureMetricsPartitions(ctx, h.pool, time.Now(), 0); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureRatingsPartition(ctx, h.pool, time.Now()); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status, created_at) values ('parse-report', $1, 'guild', 'complete', now())`,
		gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, kill, players) values ('parse-report', 0, 'Onyxia', true, $1)`,
		[]string{key}); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, role, metric_dps, kill, fought_at)
		 values ('parse-report', 0, $1, 'dps', 123.4, true, now())`, key); err != nil {
		t.Fatal(err)
	}
	components := `[{"name":"output","score":70},{"name":"survival","score":60}]`
	if _, err := h.pool.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, overall, components, model_version, fought_at, kill)
		 values ('parse-report', 0, $1, 65, $2, 'test', now(), true)`, key, components); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	var row *RosterRow
	for i := range view.Roster {
		if view.Roster[i].CharacterKey == key {
			row = &view.Roster[i]
		}
	}
	if row == nil {
		t.Fatal("character missing from roster")
	}
	if row.BestParse == nil || row.BestParse.Metric != "dps" || row.BestParse.Value != 123.4 || row.BestParse.Encounter != "Onyxia" {
		t.Errorf("bestParse = %+v, want dps 123.4 on Onyxia", row.BestParse)
	}
	if row.Rating == nil || row.Rating.Overall != 65 || row.Rating.Output != 70 || row.Rating.Survival != 60 {
		t.Errorf("rating = %+v, want overall 65, output 70, survival 60", row.Rating)
	}
}

// TestHomeStandingRanksAmongVerifiedSameSpecGearConsentPeers checks the member/officer-only
// standing line: two verified hunters with gear consent, one higher item level, and a third
// character of a different spec that must not be counted as a peer.
func TestHomeStandingRanksAmongVerifiedSameSpecGearConsentPeers(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	if err := db.EnsureMetricsPartitions(ctx, h.pool, time.Now(), 0); err != nil {
		t.Fatal(err)
	}

	seedStandingCharacter := func(email, key, class, spec string, ilvl int) int64 {
		uid := seedUser(t, h.pool, email)
		seedCharacter(t, h.pool, gid, uid, key, "member", true)
		seedExport(t, h.pool, uid, key, "us", "hardcore", key)
		recomputeMembership(t, h.pool, gid, uid)
		if _, err := h.pool.Exec(ctx,
			`insert into fight_metrics (report_id, fight_index, player_key, class, spec, ilvl, fought_at)
			 values ($1, 0, $2, $3, $4, $5, now())`, "standing-"+key, key, class, spec, ilvl); err != nil {
			t.Fatal(err)
		}
		return uid
	}

	self := seedStandingCharacter("standing-self@example.com", "us/hardcore/standingself", "Hunter", "Marksmanship", 60)
	seedStandingCharacter("standing-peer@example.com", "us/hardcore/standingpeer", "Hunter", "Marksmanship", 66)
	seedStandingCharacter("standing-other@example.com", "us/hardcore/standingother", "Mage", "Frost", 70)

	h.actor = auth.Actor{UserID: self, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)

	if view.Standing == nil {
		t.Fatal("standing = nil, want a derivable standing line")
	}
	if view.Standing.SameSpecCount != 2 {
		t.Errorf("same_spec_count = %d, want 2 (self + the other hunter, not the mage)", view.Standing.SameSpecCount)
	}
	if view.Standing.RankByItemLevel != 2 {
		t.Errorf("rank_by_item_level = %d, want 2 (behind the higher-ilvl peer)", view.Standing.RankByItemLevel)
	}
	if view.Standing.ItemLevel != 60 {
		t.Errorf("item_level = %d, want 60", view.Standing.ItemLevel)
	}
}
