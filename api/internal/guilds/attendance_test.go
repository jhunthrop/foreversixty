// api/internal/guilds/attendance_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// Task L6.8: the guild's attendance, log verification, moderator evidence and raid presence
// read a character's presence by its key, from the rated and ranked rows (fight_metrics,
// rating_scores) - never from fights.players, which holds the log's GUIDs.

// l68GUID is what engine.go writes into fights.players: a log GUID, never a character key.
const l68GUID = "Player-4395-0A1B2C3D"

// l68Night is one report with one fight: a boss fight (encounter 1084) unless trash, whose
// fights.players holds only the log's GUID, and one fight_metrics row per key in keys - the
// rows the ingest writes for a ranked fight - and one rating_scores row per key in rated.
// owner is the report's owner_id (nil for none); flagged is reports.flagged ("" for none).
type l68Night struct {
	id, visibility string
	owner          any
	at             time.Time
	trash          bool
	keys           []string
	rated          []string
	flagged        string
}

func seedL68Night(t *testing.T, pool *pgxpool.Pool, guildID int64, n l68Night) {
	t.Helper()
	ctx := context.Background()
	if n.visibility == "" {
		n.visibility = "guild"
	}
	if _, err := pool.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at, flagged)
		 values ($1, $2, $3, $4, 'complete', $5, nullif($6, ''))`,
		n.id, n.owner, guildID, n.visibility, n.at, n.flagged); err != nil {
		t.Fatal(err)
	}
	seedL68Fight(t, pool, n.id, !n.trash, n.keys)
	if len(n.rated) > 0 {
		if err := db.EnsureRatingsPartition(ctx, pool, n.at); err != nil {
			t.Fatal(err)
		}
		// rating_scores has no foreign key to reports: the harness's truncate leaves it.
		t.Cleanup(func() {
			pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, n.id)
		})
	}
	for _, key := range n.rated {
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, overall, components, model_version, fought_at, kill)
			 values ($1, 0, $2, 50, '[]', '', $3, true)`, n.id, key, n.at); err != nil {
			t.Fatal(err)
		}
	}
}

// seedL68Fight adds fight 0 to an existing report: a boss fight (encounter 1084) when boss,
// else trash, whose fights.players holds only the log's GUID, and one fight_metrics row per
// key, fought at the report's created_at.
func seedL68Fight(t *testing.T, pool *pgxpool.Pool, reportID string, boss bool, keys []string) {
	t.Helper()
	ctx := context.Background()
	var at time.Time
	if err := pool.QueryRow(ctx, `select created_at from reports where id = $1`, reportID).Scan(&at); err != nil {
		t.Fatal(err)
	}
	var encounter any
	if boss {
		encounter = int64(1084)
	}
	if _, err := pool.Exec(ctx,
		`insert into fights (report_id, fight_index, name, encounter_id, kill, start_ms, players)
		 values ($1, 0, 'Onyxia', $2, true, $3, $4)`,
		reportID, encounter, at.UnixMilli(), []string{l68GUID}); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureMetricsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	for _, key := range keys {
		if _, err := pool.Exec(ctx,
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, role, metric_dps, kill, fought_at)
			 values ($1, 0, $2, 'L68', 'warrior', 'dps', 100, true, $3)`, reportID, key, at); err != nil {
			t.Fatal(err)
		}
	}
}

// rosterRowOf reads the guild home as viewer and returns key's roster row.
func rosterRowOf(t *testing.T, h *httpHarness, guildID, viewer int64, key string) RosterRow {
	t.Helper()
	h.actor = auth.Actor{UserID: viewer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", guildID), "")
	var view HomeView
	h.data(res, &view)
	for _, row := range view.Roster {
		if row.CharacterKey == key {
			return row
		}
	}
	t.Fatalf("%s missing from the roster", key)
	return RosterRow{}
}

// The brief's scenario: one verified member keyed by a beta realm, one rated fight whose
// fights.players lists only GUIDs and whose fight_metrics row carries the key.
func TestRosterAttendanceReadsTheCharacterKeyNotTheLogsGUIDs(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "L68 Guild")
	uid := seedUser(t, h.pool, "l68-member@example.com")
	key := "/Classic Beta PvP/Name"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "pvp", "Name")
	recomputeMembership(t, h.pool, gid, uid)
	seedL68Night(t, h.pool, gid, l68Night{id: "l68-night", at: time.Now().Add(-2 * time.Hour), keys: []string{key}})

	row := rosterRowOf(t, h, gid, uid, key)
	if row.Attendance.Present != 1 || row.Attendance.Nights != 1 {
		t.Errorf("attendance = %+v, want present 1 of 1 night (the fight_metrics row carries the key)", row.Attendance)
	}
}

// A member with no rated fight reads 0 of the guild's nights.
func TestRosterAttendanceIsZeroForAMemberWithNoRatedFight(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "L68 Guild")
	uid := seedUser(t, h.pool, "l68-absent@example.com")
	key := "/Classic Beta PvP/Absent"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "pvp", "Absent")
	recomputeMembership(t, h.pool, gid, uid)
	seedL68Night(t, h.pool, gid, l68Night{id: "l68-night-other", at: time.Now().Add(-2 * time.Hour),
		keys: []string{"/Classic Beta PvP/Someoneelse"}})

	row := rosterRowOf(t, h, gid, uid, key)
	if row.Attendance.Present != 0 || row.Attendance.Nights != 1 {
		t.Errorf("attendance = %+v, want present 0 of 1 night", row.Attendance)
	}
}

// The moderator's evidence counts the nights in other accounts' reports by the character's
// key.
func TestEvidenceCountsNightsInOthersReportsByCharacterKey(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "L68 Guild")
	uid := seedUser(t, pool, "l68-evidence@example.com")
	key := "us/hardcore/l68evidence"
	seedCharacter(t, pool, gid, uid, key, "member", false)
	uploader := seedUser(t, pool, "l68-evidence-uploader@example.com")
	now := time.Now()
	seedL68Night(t, pool, gid, l68Night{id: "l68-ev-1", owner: uploader, at: now.Add(-10 * 24 * time.Hour), keys: []string{key}})
	seedL68Night(t, pool, gid, l68Night{id: "l68-ev-2", owner: uploader, at: now.Add(-3 * 24 * time.Hour), keys: []string{key}})
	// Not counted: the character's own report, and a report the character is not in.
	seedL68Night(t, pool, gid, l68Night{id: "l68-ev-own", owner: uid, at: now.Add(-5 * 24 * time.Hour), keys: []string{key}})
	seedL68Night(t, pool, gid, l68Night{id: "l68-ev-absent", owner: uploader, at: now.Add(-6 * 24 * time.Hour)})

	ev, err := s.evidenceFor(ctx, gid, uid)
	if err != nil {
		t.Fatal(err)
	}
	if ev.NightsInOthersReports != 2 {
		t.Errorf("nights in others' reports = %d, want 2", ev.NightsInOthersReports)
	}
}

// Verification by logs counts non-private reports only: a private report has no ranked rows
// in production, and is never evidence even when one is present.
func TestVerifyByLogsIgnoresPrivateReports(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "L68 Guild")
	uid := seedUser(t, pool, "l68-private@example.com")
	key := "us/hardcore/l68private"
	seedCharacter(t, pool, gid, uid, key, "member", false)
	uploader := seedUser(t, pool, "l68-private-uploader@example.com")
	now := time.Now()
	seedL68Night(t, pool, gid, l68Night{id: "l68-pv-1", owner: uploader, visibility: "private", at: now.Add(-10 * 24 * time.Hour), keys: []string{key}})
	seedL68Night(t, pool, gid, l68Night{id: "l68-pv-2", owner: uploader, visibility: "private", at: now.Add(-3 * 24 * time.Hour), keys: []string{key}})

	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var verified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = $1`, key).Scan(&verified)
	if verified {
		t.Fatal("two private nights must not verify the character")
	}

	seedL68Night(t, pool, gid, l68Night{id: "l68-pv-3", owner: uploader, visibility: "public", at: now.Add(-2 * 24 * time.Hour), keys: []string{key}})
	seedL68Night(t, pool, gid, l68Night{id: "l68-pv-4", owner: uploader, visibility: "guild", at: now.Add(-1 * 24 * time.Hour), keys: []string{key}})
	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = $1`, key).Scan(&verified)
	if !verified {
		t.Fatal("two non-private nights whose fight_metrics rows carry the key should verify the character")
	}
}

// The same scenario with a character key in the form the ingest really builds
// (character.KeyFromUnit: region/ruleset/name).
func TestRosterAttendanceReadsARealCharacterKey(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "L68 Guild")
	uid := seedUser(t, h.pool, "l68-real@example.com")
	key := "us/pvp/l68real"
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	seedExport(t, h.pool, uid, key, "us", "pvp", "L68real")
	recomputeMembership(t, h.pool, gid, uid)
	seedL68Night(t, h.pool, gid, l68Night{id: "l68-real-1", at: time.Now().Add(-3 * time.Hour), keys: []string{key}})
	seedL68Night(t, h.pool, gid, l68Night{id: "l68-real-2", at: time.Now().Add(-2 * time.Hour), rated: []string{key}})
	seedL68Night(t, h.pool, gid, l68Night{id: "l68-real-3", at: time.Now().Add(-1 * time.Hour), keys: []string{"us/pvp/l68other"}})

	row := rosterRowOf(t, h, gid, uid, key)
	if row.Attendance.Present != 2 || row.Attendance.Nights != 3 {
		t.Errorf("attendance = %+v, want present 2 of 3 nights (one ranked row, one rated row, one absence)", row.Attendance)
	}
}

// Fix round 1, item 1: a report flagged as tampered (or unverified) is never evidence. Its
// rated rows verify nobody, count toward nobody's attendance or evidence, and list nobody on
// the raid page; neither does a private report. The member uploaded all three, so the Raids
// tab lists the flagged ones to them (visibilityClause).
func TestAFlaggedReportVerifiesNobodyAndCountsForNobody(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	s := &Store{Pool: h.pool}
	gid := seedGuild(t, h.pool, "L68 Guild")
	member := seedUser(t, h.pool, "l68f-member@example.com")
	memberKey := "us/pvp/l68fmember"
	seedCharacter(t, h.pool, gid, member, memberKey, "member", true)
	seedExport(t, h.pool, member, memberKey, "us", "pvp", "L68fmember")
	recomputeMembership(t, h.pool, gid, member)
	candidate := seedUser(t, h.pool, "l68f-candidate@example.com")
	candidateKey := "us/pvp/l68fcandidate"
	seedCharacter(t, h.pool, gid, candidate, candidateKey, "member", false)

	both := []string{memberKey, candidateKey}
	now := time.Now()
	seedL68Night(t, h.pool, gid, l68Night{id: "l68f-tampered", owner: member, flagged: "tampered",
		at: now.Add(-3 * 24 * time.Hour), rated: both})
	seedL68Night(t, h.pool, gid, l68Night{id: "l68f-unverified", owner: member, flagged: "unverified",
		at: now.Add(-2 * 24 * time.Hour), rated: both, keys: both})
	seedL68Night(t, h.pool, gid, l68Night{id: "l68f-private", owner: member, visibility: "private",
		at: now.Add(-1 * 24 * time.Hour), rated: both, keys: both})

	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var verified bool
	h.pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = $1`, candidateKey).Scan(&verified)
	if verified {
		t.Error("two flagged report nights verified the candidate")
	}
	ev, err := s.evidenceFor(ctx, gid, candidate)
	if err != nil {
		t.Fatal(err)
	}
	if ev.NightsInOthersReports != 0 {
		t.Errorf("nights in others' reports = %d, want 0 (flagged reports are not evidence)", ev.NightsInOthersReports)
	}
	if row := rosterRowOf(t, h, gid, member, memberKey); row.Attendance != (Attendance{}) {
		t.Errorf("attendance = %+v, want 0 of 0 nights (a flagged or private report is not a night)", row.Attendance)
	}

	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/raids", gid), "")
	var page RaidsPage
	h.data(res, &page)
	if len(page.Rows) != 3 {
		t.Fatalf("raids rows = %d, want the 3 nights", len(page.Rows))
	}
	for _, row := range page.Rows {
		if len(row.Present) != 0 || row.Raiders != 0 {
			t.Errorf("raid %s: present %+v, raiders %d; want nobody (flagged or private)", row.ID, row.Present, row.Raiders)
		}
	}
}
