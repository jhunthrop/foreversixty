// api/internal/dataaddon/flagged_test.go
package dataaddon

import (
	"context"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// Task L6.10: a report flagged tampered or unverified is not evidence of anything a record
// says (contract 5). Every row these tests write carries the "dataaddon-l610-" prefix.
func cleanL610(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'dataaddon-l610-%'`)
		pool.Exec(bg, `delete from fights where report_id like 'dataaddon-l610-%'`)
		pool.Exec(bg, `delete from reports where id like 'dataaddon-l610-%'`)
		pool.Exec(bg, `delete from guild_characters where character_key like 'us/normal/dataaddon-l610-%'`)
		pool.Exec(bg, `delete from guilds where name like 'Dataaddon L610 %'`)
		pool.Exec(bg, `delete from users where battletag = 'DataaddonL610'`)
	})
}

// windowParses never reads a rated fight of a flagged report.
func TestWindowParsesSkipsAFlaggedReport(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	cleanL610(t, pool)
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)
	at := now.Add(-2 * 24 * time.Hour)
	if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into reports (id, visibility, status, created_at, flagged)
		values ('dataaddon-l610-flagged', 'public', 'complete', $1, 'tampered')`, at); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into fights (report_id, fight_index, encounter_id, name, kill, start_ms)
		values ('dataaddon-l610-flagged', 1, 990811, 'Lucifron', true, $1)`, at.UnixMilli()); err != nil {
		t.Fatal(err)
	}
	const forger = "us/normal/dataaddon-l610-forger"
	if _, err := pool.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
		   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
		 values ('dataaddon-l610-flagged', 1, $1, $1, 990811, true, 99, 99, false, '[]', 'test', $2)`, forger, at); err != nil {
		t.Fatal(err)
	}
	got, err := (&Store{Pool: pool}).windowParses(ctx, now.Add(-ratingWindow))
	if err != nil {
		t.Fatal(err)
	}
	if bosses, ok := got[forger]; ok {
		t.Fatalf("a flagged report's fight is in the window: %v", bosses)
	}
}

// guildLine reads one guild record's standing and nights out of the rendered Data.lua.
func guildLine(t *testing.T, lua []byte, name string) (standing, nights int) {
	t.Helper()
	m := regexp.MustCompile(`\["` + regexp.QuoteMeta(name) + `"\] = \{\s*standing = (\d+),\s*nights = (\d+),`).FindSubmatch(lua)
	if m == nil {
		t.Fatalf("no guild record for %q in:\n%s", name, lua)
	}
	standing, _ = strconv.Atoi(string(m[1]))
	nights, _ = strconv.Atoi(string(m[2]))
	return standing, nights
}

// Through the real nightly writer: a guild with three clean nights and one flagged night
// records 3 nights, and its standing counts the clean nights' kills only. The flagged night
// kills two bosses the guild never killed on a clean night, enough to lift it above a rival
// with two kills of its own; ignored, the guild's one clean kill ranks it below that rival.
func TestTheNightlyGuildRecordIgnoresAFlaggedNight(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	cleanL610(t, pool)
	now := time.Now().UTC()

	var user int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('DataaddonL610') returning id`).Scan(&user); err != nil {
		t.Fatal(err)
	}
	guild := func(name, member string) int64 {
		t.Helper()
		var id int64
		if err := pool.QueryRow(ctx, `insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`,
			name).Scan(&id); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx, `insert into guild_characters (guild_id, character_key, user_id, verified_at)
			values ($1, $2, $3, now())`, id, member, user); err != nil {
			t.Fatal(err)
		}
		return id
	}
	forgers := guild("Dataaddon L610 Forgers", "us/normal/dataaddon-l610-forger")
	rivals := guild("Dataaddon L610 Rivals", "us/normal/dataaddon-l610-rival")

	night := func(id string, guildID int64, daysAgo int, flagged string, kills ...int64) {
		t.Helper()
		at := now.Add(time.Duration(-daysAgo) * 24 * time.Hour)
		if _, err := pool.Exec(ctx, `insert into reports (id, guild_id, visibility, status, created_at, flagged)
			values ($1, $2, 'public', 'complete', $3, nullif($4, ''))`, id, guildID, at, flagged); err != nil {
			t.Fatal(err)
		}
		for i, encounter := range kills {
			if _, err := pool.Exec(ctx, `insert into fights (report_id, fight_index, encounter_id, kill, start_ms)
				values ($1, $2, $3, true, $4)`, id, i+1, encounter, at.UnixMilli()); err != nil {
				t.Fatal(err)
			}
		}
	}
	night("dataaddon-l610-f1", forgers, 3, "", 990821)
	night("dataaddon-l610-f2", forgers, 5, "", 990821)
	night("dataaddon-l610-f3", forgers, 7, "", 990821)
	night("dataaddon-l610-forged", forgers, 9, "tampered", 990822, 990823)
	night("dataaddon-l610-r1", rivals, 4, "", 990821, 990822)

	result, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Now: now})
	if err != nil {
		t.Fatal(err)
	}
	lua := result.Rendered["Data.lua"]
	forgersStanding, forgersNights := guildLine(t, lua, "Dataaddon L610 Forgers")
	rivalsStanding, _ := guildLine(t, lua, "Dataaddon L610 Rivals")
	if forgersNights != 3 {
		t.Errorf("nights = %d, want 3 (the flagged night is not a night)", forgersNights)
	}
	if forgersStanding <= rivalsStanding {
		t.Errorf("standing %d against the rivals' %d: the flagged night's two kills were counted (want the guild ranked below)",
			forgersStanding, rivalsStanding)
	}
}
