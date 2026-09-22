// api/internal/dataaddon/golden_test.go
package dataaddon

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// goldenFixturePath is the one file both this test and
// addon/tests/data_addon_spec.lua assert against, so the Go job and the
// Lua reader can never silently drift out of lockstep.
const goldenFixturePath = "../../../addon/tests/fixtures/data_addon_sample.lua"

// TestGoldenDataLuaIsByteExact builds the exact database fixture the
// plan's "Key-format reconciliation" and "Aggregation rule" sections work
// through by hand: Thoradin (every component present, two in-window
// public fights, one out-of-window fight and one private-report fight
// that must both be excluded), Mörk (a non-ASCII name on another
// region/ruleset, only two of six components computable), Hiddenhero (an
// anonymized owner, must not appear at all), and the guild Iron Vanguard
// (two verified members -- one with an apostrophe in its name -- one
// unverified member excluded from the roster, two public reports on
// different UTC dates, and progression 2 killed of 3 attempted
// encounters). One call to Run must reproduce the checked-in fixture
// byte-for-byte.
func TestGoldenDataLuaIsByteExact(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var normalUser, hiddenUser, memberUser int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Normal') returning id`).Scan(&normalUser); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag, anonymize) values ('Hidden', true) returning id`).Scan(&hiddenUser); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Member') returning id`).Scan(&memberUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from users where id in ($1, $2, $3)`, normalUser, hiddenUser, memberUser)
	})

	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values
		 ('us/normal/thoradin', 'us', 'normal', 'Thoradin', $1),
		 ('us/normal/hiddenhero', 'us', 'normal', 'Hiddenhero', $2)`, normalUser, hiddenUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key in ('us/normal/thoradin', 'us/normal/hiddenhero')`)
	})

	var guildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Iron Vanguard') returning id`,
	).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guilds where id = $1`, guildID) })

	verifiedAt := now.Add(-5 * 24 * time.Hour)
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, verified_at) values
		 ($1, 'us/normal/thoradin', $2, $4),
		 ($1, 'us/normal/o''malley', $3, $4),
		 ($1, 'us/normal/notyet', $2, null)`,
		guildID, normalUser, memberUser, verifiedAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guild_characters where guild_id = $1`, guildID) })

	reports := []struct {
		id, visibility string
		guildID        *int64
		createdAt      time.Time
	}{
		{"golden-pub-1", "public", &guildID, now.Add(-2 * 24 * time.Hour)},
		{"golden-pub-2", "public", &guildID, now.Add(-10 * 24 * time.Hour)},
		{"golden-eu-1", "public", nil, now.Add(-3 * 24 * time.Hour)},
		{"golden-private-1", "private", nil, now.Add(-1 * 24 * time.Hour)},
	}
	for _, r := range reports {
		if _, err := pool.Exec(ctx,
			`insert into reports (id, guild_id, visibility, status, created_at) values ($1, $2, $3, 'complete', $4)`,
			r.id, r.guildID, r.visibility, r.createdAt); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from reports where id in ('golden-pub-1', 'golden-pub-2', 'golden-eu-1', 'golden-private-1')`)
	})

	if _, err := pool.Exec(ctx,
		`insert into fights (report_id, fight_index, encounter_id, kill) values
		 ('golden-pub-1', 1, 100, true), ('golden-pub-1', 2, 101, false),
		 ('golden-pub-2', 1, 100, true), ('golden-pub-2', 2, 102, true)`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from fights where report_id in ('golden-pub-1', 'golden-pub-2')`) })

	thoradinA := `[{"name":"output","score":90,"excluded":false},{"name":"survival","score":70,"excluded":false},` +
		`{"name":"mechanics","score":85,"excluded":false},{"name":"utility","score":60,"excluded":false},` +
		`{"name":"preparation","score":95,"excluded":false},{"name":"activity","score":null,"excluded":true}]`
	thoradinB := `[{"name":"output","score":92,"excluded":false},{"name":"survival","score":74,"excluded":false},` +
		`{"name":"mechanics","score":83,"excluded":false},{"name":"utility","score":64,"excluded":false},` +
		`{"name":"preparation","score":93,"excluded":false},{"name":"activity","score":91,"excluded":false}]`
	mork := `[{"name":"output","score":70,"excluded":false},{"name":"survival","score":null,"excluded":true},` +
		`{"name":"mechanics","score":70,"excluded":false},{"name":"utility","score":null,"excluded":true},` +
		`{"name":"preparation","score":null,"excluded":true},{"name":"activity","score":null,"excluded":true}]`

	insertScore := func(reportID string, index int, playerKey string, overall float64, components string, foughtAt time.Time) {
		// rating_scores is partitioned by month on fought_at (migration
		// 0022); ensure the partition exists before inserting into it, the
		// same way rating.Store.RateFight does before its own insert. This
		// fixture spans two calendar months (the in-window fights in
		// September 2026, the deliberately-out-of-window fight about 100
		// days back in June 2026), so this must run per insert, not once.
		if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
			 values ($1, $2, $3, $3, 100, false, $4, $4, false, $5, 'test', $6)`,
			reportID, index, playerKey, overall, components, foughtAt); err != nil {
			t.Fatal(err)
		}
	}
	insertScore("golden-pub-1", 1, "us/normal/thoradin", 80, thoradinA, now.Add(-2*24*time.Hour))
	insertScore("golden-pub-2", 1, "us/normal/thoradin", 82, thoradinB, now.Add(-10*24*time.Hour))
	insertScore("golden-pub-1", 3, "us/normal/thoradin", 999, thoradinB, now.Add(-100*24*time.Hour))   // out of window
	insertScore("golden-private-1", 1, "us/normal/thoradin", 999, thoradinB, now.Add(-1*24*time.Hour)) // private report
	insertScore("golden-pub-1", 4, "us/normal/hiddenhero", 999, thoradinB, now.Add(-2*24*time.Hour))   // anonymized
	insertScore("golden-eu-1", 1, "eu/pvp/mörk", 88, mork, now.Add(-3*24*time.Hour))
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id in
			('golden-pub-1', 'golden-pub-2', 'golden-eu-1', 'golden-private-1')`)
	})

	result, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: now})
	if err != nil {
		t.Fatal(err)
	}
	if result.Characters != 2 || result.Guilds != 1 {
		t.Fatalf("result = %+v, want 2 characters and 1 guild", result)
	}

	want, err := os.ReadFile(goldenFixturePath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(result.Rendered["Data.lua"], want) {
		t.Errorf("Data.lua does not match the checked-in golden fixture.\ngot:\n%s\nwant:\n%s", result.Rendered["Data.lua"], want)
	}
	assertValidLua(t, result.Rendered["Data.lua"])
}

// assertValidLua runs the rendered file through the `lua` interpreter when
// one is on PATH (dispatch: "run it through the lua interpreter if present
// in the test environment"), asserting it parses and defines the global
// with no runtime error. Without one, it falls back to a minimal
// structural check and says so, per the dispatch's own fallback.
func assertValidLua(t *testing.T, data []byte) {
	t.Helper()
	path, err := exec.LookPath("lua")
	if err != nil {
		t.Log("no `lua` interpreter on PATH; falling back to a minimal structural check")
		if !bytes.Contains(data, []byte("ForeverSixtyData = {")) || !bytes.Contains(data, []byte("\n}")) {
			t.Errorf("does not look like a well-formed ForeverSixtyData table: %s", data)
		}
		return
	}
	tmp := filepath.Join(t.TempDir(), "check.lua")
	script := string(data) + "\nassert(type(ForeverSixtyData) == \"table\" and ForeverSixtyData.format == 1)\n"
	if err := os.WriteFile(tmp, []byte(script), 0o600); err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command(path, tmp).CombinedOutput(); err != nil {
		t.Errorf("lua rejected the rendered file: %v\n%s", err, out)
	}
}
