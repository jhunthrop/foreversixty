package rating

// Task L6.10: the Forever Score ignores flagged reports. A report flagged tampered or
// unverified (reports.flagged) is evidence of nothing: no score input reads its rows
// (presence.EvidenceReport, the one report selector the score and the guild side share).

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
		// The main's two alts are linked to its account; neither has a rated fight to roll onto.
		wantCounts := RerateCounts{Flagged: RerateGroup{Characters: 1, Fights: 1, Rerated: 1},
			Linked: RerateGroup{Characters: 2}}
		if got != wantCounts {
			t.Errorf("run %d: counts %+v, want %+v", run, got, wantCounts)
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

// Task L6.12: the re-rate reaches everyone whose score the flagged night moved, not only the
// characters in it. On one flagged guild night: the main was there (its fights, as_of,
// discipline); a guild-mate who missed it had it counted as an absence (attendance's
// denominator); and the main's alt, never in that guild, had the night's one unique kill
// counted (experience, across the account's linked characters). Each gets its newest clean
// fight re-rated and its score rolled back onto it; a second run changes nothing.
func TestRerateFlaggedReachesTheNightsGuildmatesAndTheAccountsAlts(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	flagged, mateKey, altReport := sv2Report+"c", "us/normal/sv2mate", sv2Report+"alt"
	t.Cleanup(func() {
		bg := context.Background()
		fx.pool.Exec(bg, `delete from guild_characters where character_key = $1`, mateKey)
		fx.pool.Exec(bg, `delete from characters where key = $1`, mateKey)
		fx.pool.Exec(bg, `delete from users where email = 'scorev2-mate@example.com'`)
	})
	var mateUser int64
	if err := fx.pool.QueryRow(ctx, `insert into users (email) values ('scorev2-mate@example.com')
		on conflict (email) do update set email = excluded.email returning id`).Scan(&mateUser); err != nil {
		t.Fatal(err)
	}
	mustExec(t, fx.pool, `insert into characters (key, region, ruleset, name, class, user_id)
		values ($1, 'us', 'normal', 'Sv2mate', 'Warrior', $2) on conflict (key) do update set user_id = excluded.user_id`,
		mateKey, mateUser)
	mustExec(t, fx.pool, `insert into guild_characters (guild_id, character_key, user_id, verified_at)
		values ($1, $2, $3, now()) on conflict (guild_id, character_key) do update set verified_at = now()`,
		fx.guildID, mateKey, mateUser)

	rosterOf := func(name, guid string) summary.Summary {
		row := summary.RosterRow{GUID: guid, Name: name, Class: "Warrior", Spec: "Arms", Role: "dps",
			ExecutionScore: execScorePtr(0.8)}
		return summary.Summary{EncounterID: sv2Encounter, Difficulty: 3, DurationMS: 180000, Kill: true,
			Roster: []summary.RosterRow{row}, Combatants: []summary.CombatantRow{preparedCombatant(guid)}}
	}
	rate := func(reportID string, index int, foughtAt time.Time, sum summary.Summary) RatedFight {
		t.Helper()
		mustExec(t, fx.pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
			values ($1, $2, $3, 'Score v2 boss', 3, true, $4)`, reportID, index, sum.EncounterID, foughtAt.UnixMilli())
		f := RatedFight{ReportID: reportID, FightIndex: index, Region: "us", Ruleset: "normal",
			FoughtAt: foughtAt, EncounterID: sum.EncounterID, Summary: sum}
		if err := fx.store.RateFight(ctx, f); err != nil {
			t.Fatal(err)
		}
		return f
	}
	nightOf := func(reportID string) time.Time {
		t.Helper()
		var at time.Time
		if err := fx.pool.QueryRow(ctx, `select created_at from reports where id = $1`, reportID).Scan(&at); err != nil {
			t.Fatal(err)
		}
		return at
	}

	// The guild-mate was at the first two nights and missed the flagged third.
	mateSum := rosterOf("Sv2mate", "sv2-mate")
	rate(sv2Report+"a", 5, nightOf(sv2Report+"a").Add(80*time.Minute), mateSum)
	mateFight := rate(sv2Report+"b", 5, nightOf(sv2Report+"b").Add(80*time.Minute), mateSum)
	// The alt raided once, outside the guild.
	mustExec(t, fx.pool, `insert into reports (id, owner_id, title, visibility, zone, status, created_at)
		values ($1, $2, 'Alt night', 'public', 'Molten Core', 'complete', $3)`,
		altReport, fx.userID, nightOf(sv2Report+"b").Add(24*time.Hour))
	altSum := rosterOf("Sv2alta", "sv2-alta")
	altFight := rate(altReport, 1, nightOf(altReport).Add(15*time.Minute), altSum)
	// The flagged night holds a kill nobody on the account made anywhere else.
	unique := nightOf(flagged).Add(2 * time.Hour)
	if err := db.EnsureRatingsPartition(ctx, fx.pool, unique); err != nil {
		t.Fatal(err)
	}
	mustExec(t, fx.pool, `insert into rating_scores (report_id, fight_index, player_key, overall, components,
		model_version, fought_at, kill, encounter_id) values ($1, 7, $2, 50, '[]', $3, $4, true, $5)`,
		flagged, fx.main, ratingengine.DefaultModelVersion, unique, sv2Encounter+6)

	// The rater rolled both scores up while the flagged night still counted.
	snapshot := func(f RatedFight, key string) (score int, computed time.Time) {
		t.Helper()
		if err := fx.pool.QueryRow(ctx, `select coalesce(score, -1), computed_at from rating_scores
			where report_id = $1 and fight_index = $2 and player_key = $3`, f.ReportID, f.FightIndex, key).
			Scan(&score, &computed); err != nil {
			t.Fatal(err)
		}
		return score, computed
	}
	type member struct {
		name, key      string
		fight          RatedFight
		before, want   int
		computedBefore time.Time
	}
	members := []*member{{name: "guild-mate", key: mateKey, fight: mateFight}, {name: "alt", key: fx.altA, fight: altFight}}
	for _, m := range members {
		if err := fx.store.RollUpScores(ctx, m.fight); err != nil {
			t.Fatal(err)
		}
		m.before, m.computedBefore = snapshot(m.fight, m.key)
	}
	flagReport(t, fx, flagged)
	for _, m := range members {
		sc, err := fx.store.ScoreV2(ctx, m.key)
		if err != nil {
			t.Fatal(err)
		}
		m.want = sc.Score
		if m.want == m.before {
			t.Fatalf("the fixture's flagged night does not move the %s's score (%d)", m.name, m.before)
		}
	}

	objects := map[string][]byte{}
	for _, f := range []RatedFight{mateFight, altFight} {
		body, err := json.Marshal(f.Summary)
		if err != nil {
			t.Fatal(err)
		}
		objects[store.Keys{ReportID: f.ReportID}.FightSummary(f.FightIndex)] = body
	}
	mainSum := rosterOf("Sv2main", "sv2-g")
	mainSum.EncounterID = sv2Encounter + 3 // the newest clean fight's own boss
	mainBody, err := json.Marshal(mainSum)
	if err != nil {
		t.Fatal(err)
	}
	objects[store.Keys{ReportID: sv2Report + "b"}.FightSummary(4)] = mainBody
	deps := BackfillDeps{Store: fx.store, Summaries: fakeGetter{objects: objects}}
	want := RerateCounts{
		Flagged:    RerateGroup{Characters: 1, Fights: 1, Rerated: 1},
		Guildmates: RerateGroup{Characters: 1, Fights: 1, Rerated: 1},
		// The main's other alt has no rated fight, so no score to roll up.
		Linked: RerateGroup{Characters: 2, Fights: 1, Rerated: 1},
	}
	for run := 1; run <= 2; run++ {
		got, err := RerateFlagged(ctx, deps)
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Errorf("run %d: counts %+v, want %+v", run, got, want)
		}
		for _, m := range members {
			score, computed := snapshot(m.fight, m.key)
			if score != m.want {
				t.Errorf("run %d: the %s's newest clean fight's snapshot is %d, want %d (the score without the flagged night; it was %d)",
					run, m.name, score, m.want, m.before)
			}
			if !computed.After(m.computedBefore) {
				t.Errorf("run %d: the %s's fight was not re-rated (computed_at %v, was %v)", run, m.name, computed, m.computedBefore)
			}
		}
	}
}

// The guild-mates of a flagged night are the guild's verified members (the membership
// guildNights counts nights for), with or without a rated fight, and anyone with a rated fight
// in the guild's reports inside the attendance window; a raider whose only fight in the guild
// is older than the window is not one.
func TestRerateFlaggedGuildmatesAreTheMembersAndTheWindowsRaiders(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	const member, old = "us/normal/sv2member", "us/normal/sv2old"
	t.Cleanup(func() {
		bg := context.Background()
		fx.pool.Exec(bg, `delete from guild_characters where character_key = $1`, member)
		fx.pool.Exec(bg, `delete from rating_scores where player_key = $1`, old)
		fx.pool.Exec(bg, `delete from users where email = 'scorev2-member@example.com'`)
	})
	var memberUser int64
	if err := fx.pool.QueryRow(ctx, `insert into users (email) values ('scorev2-member@example.com')
		on conflict (email) do update set email = excluded.email returning id`).Scan(&memberUser); err != nil {
		t.Fatal(err)
	}
	// A verified member with no rated fight anywhere.
	mustExec(t, fx.pool, `insert into guild_characters (guild_id, character_key, user_id, verified_at)
		values ($1, $2, $3, now()) on conflict (guild_id, character_key) do update set verified_at = now()`,
		fx.guildID, member, memberUser)
	// A raider of the guild's, no member, whose only fight was 120 days ago.
	oldNight := time.Now().UTC().Add(-120 * 24 * time.Hour)
	mustExec(t, fx.pool, `insert into reports (id, owner_id, guild_id, title, visibility, zone, status, created_at)
		values ($1, $2, $3, 'Old night', 'public', 'Molten Core', 'complete', $4)`,
		sv2Report+"old", fx.userID, fx.guildID, oldNight)
	if err := db.EnsureRatingsPartition(ctx, fx.pool, oldNight); err != nil {
		t.Fatal(err)
	}
	mustExec(t, fx.pool, `insert into rating_scores (report_id, fight_index, player_key, overall, components,
		model_version, fought_at, kill) values ($1, 1, $2, 50, '[]', '', $3, true)`, sv2Report+"old", old, oldNight)
	flagReport(t, fx, sv2Report+"c")

	guildmates := func(since time.Time) map[string]reratePick {
		t.Helper()
		picks, err := fx.store.flaggedRerates(ctx, since)
		if err != nil {
			t.Fatal(err)
		}
		out := map[string]reratePick{}
		for _, p := range picks {
			if p.group == rerateGuildmate {
				out[p.key] = p
			}
		}
		return out
	}
	got := guildmates(time.Now().Add(-AttendanceWindow))
	if p, ok := got[member]; !ok || p.ReportID != "" {
		t.Errorf("the verified member with no rated fight: %+v (in the set %v), want a guild-mate with no fight to re-rate", p, ok)
	}
	if _, ok := got[old]; ok {
		t.Errorf("the raider whose only fight is 120 days old is a guild-mate: %+v", got)
	}
	if len(got) != 1 {
		t.Errorf("guild-mates = %+v, want only %s", got, member)
	}
	// With the window reaching back past the old night, the old raider is a guild-mate too.
	got = guildmates(oldNight.Add(-time.Hour))
	if _, ok := got[old]; !ok || len(got) != 2 {
		t.Errorf("guild-mates with the window past the old night = %+v, want %s and %s", got, member, old)
	}
}

// An invite row (the synthetic account: key an invite link writes, verified at once) is no
// character, so it is never a guild-mate of a flagged night: there is no score to re-rate.
func TestRerateFlaggedGuildmatesLeaveOutAnInviteRow(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	var invitee int64
	if err := fx.pool.QueryRow(ctx, `insert into users (email) values ('scorev2-invitee@example.com')
		on conflict (email) do update set email = excluded.email returning id`).Scan(&invitee); err != nil {
		t.Fatal(err)
	}
	invite := fmt.Sprintf("account:%d", invitee)
	t.Cleanup(func() {
		bg := context.Background()
		fx.pool.Exec(bg, `delete from guild_characters where character_key = $1`, invite)
		fx.pool.Exec(bg, `delete from users where email = 'scorev2-invitee@example.com'`)
	})
	mustExec(t, fx.pool, `insert into guild_characters (guild_id, character_key, user_id, rank, source, verified_at, verified_by)
		values ($1, $2, $3, 'member', 'invite', now(), 'invite')`, fx.guildID, invite, invitee)
	flagReport(t, fx, sv2Report+"c")
	picks, err := fx.store.flaggedRerates(ctx, time.Now().Add(-AttendanceWindow))
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range picks {
		if p.key == invite {
			t.Errorf("the invite row %s is picked for a re-rate: %+v", invite, p)
		}
	}
}
