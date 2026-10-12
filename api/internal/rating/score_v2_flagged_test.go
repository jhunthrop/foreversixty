package rating

// Task L6.10: the Forever Score ignores flagged reports. A report flagged tampered or
// unverified (reports.flagged) is evidence of nothing: no score input reads its rows
// (presence.EvidenceReport, the one report selector the score and the guild side share).

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// flagReport marks a report tampered, the way parse/sample.go's Flag does.
func flagReport(t *testing.T, fx scoreFixture, id string) {
	t.Helper()
	mustExec(t, fx.pool, `update reports set flagged = 'tampered' where id = $1`, id)
}

// A character whose only rated fight is in a flagged report has no score at all.
func TestAFlaggedReportsOnlyRatedFightGivesNoScore(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	mustExec(t, fx.pool, `delete from rating_scores where player_key = $1 and not (report_id = $2 and fight_index = 4)`,
		fx.main, sv2Report+"c")
	if got, err := fx.store.ScoreV2(ctx, fx.main); err != nil || got.Fights != 1 {
		t.Fatalf("before the flag: %+v, %v; want a score on its one fight", got, err)
	}
	flagReport(t, fx, sv2Report+"c")
	got, err := fx.store.ScoreV2(ctx, fx.main)
	if !errors.Is(err, ErrNoLogs) {
		t.Fatalf("the only rated fight is in a flagged report: %+v, err = %v; want ErrNoLogs", got, err)
	}
}

// One clean and one flagged fight score exactly as the clean fight alone: the flagged
// report's rated row (discipline, fights, as_of, attendance) and its ranked rows (a better
// parse, a boss nobody else killed) count for nothing.
func TestAFlaggedFightCountsForNothingInTheScore(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	flagged := sv2Report + "c"
	// The clean fight is the first night's first (the one pulled without a flask); the
	// flagged one is the third night's first, fully prepared.
	mustExec(t, fx.pool, `delete from rating_scores where player_key = $1
		and not ((report_id = $2 and fight_index = 1) or (report_id = $3 and fight_index = 1))`,
		fx.main, sv2Report+"a", flagged)
	// No Battle.net item level: the score's item level is the newest ranked fight's.
	mustExec(t, fx.pool, `update characters set equipped_item_level = null where key = $1`, fx.main)
	at := time.Now().UTC().Truncate(24 * time.Hour).Add(-25 * 24 * time.Hour)
	if err := db.EnsureMetricsPartition(ctx, fx.pool, at); err != nil {
		t.Fatal(err)
	}
	for _, m := range []struct {
		fight     int
		encounter int64
		dps       float64
	}{{1, sv2Encounter, 99}, {7, sv2Encounter + 6, 50}} {
		mustExec(t, fx.pool, `insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
			role, metric_dps, encounter_id, difficulty, kill, phase, fought_at, ilvl)
			values ($1, $2, $3, 'Sv2main', 'Warrior', 'Arms', 'dps', $4, $5, 3, true, '', $6, 77)`,
			flagged, m.fight, fx.main, m.dps, m.encounter, at)
	}

	counted, err := fx.store.ScoreV2(ctx, fx.main)
	if err != nil {
		t.Fatal(err)
	}
	flagReport(t, fx, flagged)
	ignored, err := fx.store.ScoreV2(ctx, fx.main)
	if err != nil {
		t.Fatal(err)
	}
	for _, sql := range []string{
		`delete from rating_scores where report_id = $1`, `delete from fight_metrics where report_id = $1`,
		`delete from fights where report_id = $1`, `delete from reports where id = $1`,
	} {
		mustExec(t, fx.pool, sql, flagged)
	}
	alone, err := fx.store.ScoreV2(ctx, fx.main)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("counted %+v\nignored %+v\nalone   %+v", counted, ignored, alone)

	if !ignored.AsOf.Equal(alone.AsOf) {
		t.Errorf("as_of %v with the flagged fight, want %v as with the clean fight alone", ignored.AsOf, alone.AsOf)
	}
	ignored.AsOf, alone.AsOf = time.Time{}, time.Time{}
	if ignored != alone {
		t.Errorf("with a flagged fight: %+v\nwant as with the clean fight alone: %+v", ignored, alone)
	}
	// The flagged report would have moved every component the score reads from a report.
	if counted.Components.Experience == alone.Components.Experience || counted.Components.Parses == alone.Components.Parses ||
		counted.Components.Attendance == alone.Components.Attendance || counted.Components.Discipline == alone.Components.Discipline ||
		counted.Fights == alone.Fights || counted.ItemLevel == alone.ItemLevel {
		t.Errorf("the fixture's flagged report moves nothing: counted %+v, alone %+v", counted, alone)
	}
}

// The attendance component's nights exclude a flagged report: a guild night that is a
// flagged report is neither a night nor an absence.
func TestTheAttendanceComponentsNightsExcludeAFlaggedReport(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	flagged := sv2Report + "c"
	// The main was not at the third night.
	mustExec(t, fx.pool, `delete from rating_scores where report_id = $1 and player_key = $2`, flagged, fx.main)
	since := time.Now().Add(-AttendanceWindow)
	if nights, attended, err := fx.store.guildNights(ctx, fx.main, since); err != nil || nights != 3 || attended != 2 {
		t.Fatalf("before the flag: %d of %d nights (%v), want 2 of 3", attended, nights, err)
	}
	flagReport(t, fx, flagged)
	nights, attended, err := fx.store.guildNights(ctx, fx.main, since)
	if err != nil {
		t.Fatal(err)
	}
	if nights != 2 || attended != 2 {
		t.Errorf("guild nights = %d of %d, want 2 of 2 (the flagged night is not a night)", attended, nights)
	}
	got, err := fx.store.ScoreV2(ctx, fx.main)
	if err != nil {
		t.Fatal(err)
	}
	if got.Components.Attendance != AttendanceMax || got.Attendance != 1 {
		t.Errorf("attendance %d (%v), want %d (1)", got.Components.Attendance, got.Attendance, AttendanceMax)
	}
}

// The guild contract says log verification and roster attendance count "the same evidence
// the Forever Score trusts". On one seeded guild - three clean nights the main was at, one
// clean night it missed, and one flagged night it was at - the guild roster's attendance and
// the score's attendance read the same nights and the same presence.
func TestTheScoreAndTheGuildRosterReadTheSameNights(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	today := time.Now().UTC().Truncate(24 * time.Hour)
	night := func(id string, daysAgo int, flagged string, present string) {
		t.Helper()
		at := today.Add(time.Duration(-daysAgo)*24*time.Hour + 20*time.Hour)
		mustExec(t, fx.pool, `insert into reports (id, owner_id, guild_id, title, visibility, zone, status, created_at, flagged)
			values ($1, $2, $3, 'Score v2', 'public', 'Molten Core', 'complete', $4, nullif($5, ''))`,
			id, fx.userID, fx.guildID, at, flagged)
		fought := at.Add(15 * time.Minute)
		mustExec(t, fx.pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
			values ($1, 1, $2, 'Score v2 boss', 3, true, $3)`, id, sv2Encounter, fought.UnixMilli())
		if err := db.EnsureRatingsPartition(ctx, fx.pool, fought); err != nil {
			t.Fatal(err)
		}
		mustExec(t, fx.pool, `insert into rating_scores (report_id, fight_index, player_key, overall, components,
			model_version, fought_at, kill) values ($1, 1, $2, 50, '[]', '', $3, true)`, id, present, fought)
	}
	night(sv2Report+"d-flagged", 5, "tampered", fx.main)
	night(sv2Report+"e-missed", 3, "", fx.altA)

	tx, err := fx.pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := guilds.RecomputeMembership(ctx, tx, fx.guildID, &fx.userID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		fx.pool.Exec(context.Background(), `delete from guild_members where guild_id = $1`, fx.guildID)
	})
	roster, err := (&guilds.Store{Pool: fx.pool}).HomeRoster(ctx, fx.guildID, guilds.Guild{}, fx.userID, false, false)
	if err != nil {
		t.Fatal(err)
	}
	var guildSide *guilds.Attendance
	for i := range roster {
		if roster[i].CharacterKey == fx.main {
			guildSide = &roster[i].Attendance
		}
	}
	if guildSide == nil {
		t.Fatalf("%s is not on the guild's roster: %+v", fx.main, roster)
	}
	nights, attended, err := fx.store.guildNights(ctx, fx.main, time.Now().Add(-AttendanceWindow))
	if err != nil {
		t.Fatal(err)
	}
	want := guilds.Attendance{Present: 3, Nights: 4}
	if *guildSide != want {
		t.Errorf("guild roster attendance = %+v, want %+v", *guildSide, want)
	}
	if (guilds.Attendance{Present: attended, Nights: nights}) != *guildSide {
		t.Errorf("the score's attendance reads %d of %d nights, the guild's roster %d of %d: not the same evidence",
			attended, nights, guildSide.Present, guildSide.Nights)
	}
}

// The re-rate job runs the backfill's own path (the stored summary, RateFight) for every
// character with a row in a flagged report, then rolls its score up: the snapshot on its
// newest clean fight drops to the score without the flagged report. The flagged report's own
// rows are left in place, the model version is not bumped, and a second run changes nothing.
func TestRerateFlaggedRollsTheScoreBackDown(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	flagged, newestClean := sv2Report+"c", sv2Report+"b"
	row := summary.RosterRow{GUID: "sv2-g", Name: "Sv2main", Class: "Warrior", Spec: "Arms", Role: "dps",
		ExecutionScore: execScorePtr(0.8)}
	sum := summary.Summary{
		EncounterID: int64(sv2Encounter + 3), Difficulty: 3, DurationMS: 180000, Kill: true,
		Roster: []summary.RosterRow{row}, Combatants: []summary.CombatantRow{preparedCombatant(row.GUID)},
	}
	// The rater rolled the score up onto the newest clean fight while the forged report
	// still counted.
	if err := fx.store.RollUpScores(ctx, RatedFight{ReportID: newestClean, FightIndex: 4, Region: "us",
		Ruleset: "normal", Summary: sum}); err != nil {
		t.Fatal(err)
	}
	snapshot := func() (score int, tier string, computed time.Time) {
		t.Helper()
		if err := fx.pool.QueryRow(ctx, `select coalesce(score, -1), coalesce(tier, ''), computed_at from rating_scores
			where report_id = $1 and fight_index = 4 and player_key = $2`, newestClean, fx.main).
			Scan(&score, &tier, &computed); err != nil {
			t.Fatal(err)
		}
		return score, tier, computed
	}
	before, _, computedBefore := snapshot()
	flagReport(t, fx, flagged)
	want, err := fx.store.ScoreV2(ctx, fx.main)
	if err != nil {
		t.Fatal(err)
	}
	if want.Score == before {
		t.Fatalf("the fixture's flagged report does not move the score (%d)", before)
	}

	body, err := json.Marshal(sum)
	if err != nil {
		t.Fatal(err)
	}
	getter := fakeGetter{objects: map[string][]byte{store.Keys{ReportID: newestClean}.FightSummary(4): body}}
	deps := BackfillDeps{Store: fx.store, Summaries: getter}
	for run := 1; run <= 2; run++ {
		got, err := RerateFlagged(ctx, deps)
		if err != nil {
			t.Fatal(err)
		}
		if got != (RerateCounts{Characters: 1, Fights: 1, Rerated: 1}) {
			t.Errorf("run %d: counts %+v, want 1 character, 1 fight, 1 re-rated, none failed", run, got)
		}
		score, tier, computed := snapshot()
		if score != want.Score || tier != want.Tier {
			t.Errorf("run %d: the newest clean fight's snapshot is %d %q, want %d %q (the score without the flagged report; it was %d)",
				run, score, tier, want.Score, want.Tier, before)
		}
		if !computed.After(computedBefore) {
			t.Errorf("run %d: the fight was not re-rated (computed_at %v, was %v)", run, computed, computedBefore)
		}
	}
	var left int
	var versions []string
	if err := fx.pool.QueryRow(ctx, `select count(*) from rating_scores where report_id = $1`, flagged).Scan(&left); err != nil {
		t.Fatal(err)
	}
	if left != 4 {
		t.Errorf("the flagged report holds %d rated rows, want its 4 left in place", left)
	}
	rows, err := fx.pool.Query(ctx, `select distinct model_version from rating_scores where report_id like 'sv2-%'`)
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var v string
		rows.Scan(&v)
		versions = append(versions, v)
	}
	rows.Close()
	if len(versions) != 1 || versions[0] != ratingengine.DefaultModelVersion {
		t.Errorf("model versions %v, want only %s (no bump)", versions, ratingengine.DefaultModelVersion)
	}
}
