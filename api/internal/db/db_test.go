package db

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

func testURL(t *testing.T) string {
	t.Helper()
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	return u
}

func TestMigrateCreatesSubscribers(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	if err := Migrate(url); err != nil { // idempotent
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(), `select count(*) from information_schema.tables where table_name = 'subscribers'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("subscribers table count = %d", n)
	}
}

func TestMigrateCreatesBuildsWithTheContractColumns(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(),
		`select column_name, data_type, is_nullable from information_schema.columns where table_name = 'builds'`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	got := map[string][2]string{}
	for rows.Next() {
		var name, dataType, nullable string
		if err := rows.Scan(&name, &dataType, &nullable); err != nil {
			t.Fatal(err)
		}
		got[name] = [2]string{dataType, nullable}
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	want := map[string][2]string{
		"id":           {"text", "NO"},
		"class_id":     {"smallint", "NO"},
		"race_id":      {"smallint", "NO"},
		"tree_version": {"text", "NO"},
		"point_order":  {"ARRAY", "NO"},
		"gear":         {"jsonb", "NO"},
		"title":        {"text", "YES"},
		"created_at":   {"timestamp with time zone", "NO"},
		"views":        {"bigint", "NO"},
		"user_id":      {"bigint", "YES"},
	}
	if len(got) != len(want) {
		t.Fatalf("columns = %v, want exactly %d columns", got, len(want))
	}
	for name, w := range want {
		if got[name] != w {
			t.Errorf("column %s = %v, want %v", name, got[name], w)
		}
	}
}

func TestMigrateCreatesEveryPhase3Table(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	for _, name := range []string{
		"users", "sessions", "login_tokens", "devices", "pairing_codes",
		"guilds", "guild_members", "guild_characters", "characters", "uploads", "reports",
		"fights", "raw_chunks", "fight_metrics", "percentile_digests",
		"moderation", "addon_exports", "addon_inbox",
	} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.tables where table_name = $1`, name).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("table %s: count = %d, want 1", name, n)
		}
	}
}

// tableCount is how the migration tests below ask whether a table is there.
func tableCount(t *testing.T, pool *pgxpool.Pool, name string) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from information_schema.tables where table_name = $1`, name).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// migrateTo runs the migrations to an exact version, which is how a test can
// take 0005 back down without disturbing the migrations under it.
func migrateTo(t *testing.T, url string, version uint) {
	t.Helper()
	src, err := iofs.New(migrations, "migrations")
	if err != nil {
		t.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", src, url)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	if err := m.Migrate(version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		t.Fatal(err)
	}
}

func TestMigration0005DownReversesUp(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	// Whatever this test asserts, the shared database is left at the latest version.
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	// A partition exists, so the down migration has to take the children with it.
	if err := EnsureMetricsPartition(context.Background(), pool,
		time.Date(2027, 3, 1, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}

	if n := indexCount(t, pool, "fights_encounter_idx"); n != 1 {
		t.Error("encounter name resolution has no index to run on")
	}

	migrateTo(t, url, 4)
	if n := indexCount(t, pool, "fights_encounter_idx"); n != 0 {
		t.Error("fights_encounter_idx survived the down migration")
	}
	for _, gone := range []string{
		"users", "sessions", "login_tokens", "devices", "pairing_codes",
		"guilds", "guild_members", "characters", "uploads", "reports",
		"fights", "raw_chunks", "fight_metrics", "fight_metrics_2027_03",
		"percentile_digests", "moderation", "addon_exports", "addon_inbox",
	} {
		if n := tableCount(t, pool, gone); n != 0 {
			t.Errorf("table %s survived the down migration", gone)
		}
	}
	// 0005 must not reverse anything an earlier migration created.
	if n := tableCount(t, pool, "builds"); n != 1 {
		t.Error("the down migration took a table from an earlier migration")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	for _, back := range []string{"users", "fights", "fight_metrics", "addon_inbox"} {
		if n := tableCount(t, pool, back); n != 1 {
			t.Errorf("table %s did not come back", back)
		}
	}
	if n := indexCount(t, pool, "fights_encounter_idx"); n != 1 {
		t.Error("fights_encounter_idx did not come back")
	}
}

func indexCount(t *testing.T, pool *pgxpool.Pool, name string) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from pg_indexes where indexname = $1`, name).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestFightMetricsIsPartitionedByMonth(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()

	var strategy string
	if err := pool.QueryRow(ctx,
		`select partstrat from pg_partitioned_table p
		 join pg_class c on c.oid = p.partrelid where c.relname = 'fight_metrics'`).Scan(&strategy); err != nil {
		t.Fatalf("fight_metrics is not partitioned: %v", err)
	}
	if strategy != "r" {
		t.Fatalf("partition strategy = %q, want range", strategy)
	}

	first := time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC)
	if err := EnsureMetricsPartitions(ctx, pool, first, 2); err != nil {
		t.Fatal(err)
	}
	// Idempotent: a second sweep over the same months is a no-op.
	if err := EnsureMetricsPartitions(ctx, pool, first, 2); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"fight_metrics_2026_12", "fight_metrics_2027_01", "fight_metrics_2027_02"} {
		var n int
		if err := pool.QueryRow(ctx,
			`select count(*) from information_schema.tables where table_name = $1`, want).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("partition %s missing", want)
		}
	}

	if _, err := pool.Exec(ctx, `delete from fight_metrics`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into fight_metrics (report_id, fight_index, player_key, metric_dps, fought_at)
		 values ('r1', 1, 'us/nightslayer/baelgrim', 1234.5, $1)`, first); err != nil {
		t.Fatal(err)
	}
	var landed string
	if err := pool.QueryRow(ctx,
		`select tableoid::regclass::text from fight_metrics where report_id = 'r1'`).Scan(&landed); err != nil {
		t.Fatal(err)
	}
	if landed != "fight_metrics_2026_12" {
		t.Fatalf("row landed in %s, want fight_metrics_2026_12", landed)
	}
	if _, err := pool.Exec(ctx, `delete from fight_metrics`); err != nil {
		t.Fatal(err)
	}
}

func TestPartitionJobCreatesTheMonthsAhead(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job := &PartitionJob{Pool: pool, Ahead: 1, Every: time.Hour,
		Now: func() time.Time { return time.Date(2027, 6, 15, 0, 0, 0, 0, time.UTC) }}
	if err := job.Run(ctx); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"fight_metrics_2027_06", "fight_metrics_2027_07"} {
		var n int
		if err := pool.QueryRow(ctx,
			`select count(*) from information_schema.tables where table_name = $1`, want).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("partition %s missing", want)
		}
	}
}

func TestPartitionWorkCarriesTheErrorWhenTheDatabaseIsGone(t *testing.T) {
	url := testURL(t)
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	pool.Close()
	ctx := context.Background()
	if err := EnsureMetricsPartition(ctx, pool, time.Now()); err == nil {
		t.Fatal("a closed pool must report itself")
	}
	if err := EnsureMetricsPartitions(ctx, pool, time.Now(), 2); err == nil {
		t.Fatal("a closed pool must report itself")
	}
	job := &PartitionJob{Pool: pool}
	if err := job.Run(ctx); err == nil {
		t.Fatal("the job's first pass must fail loudly")
	}
}

func TestMetricsPartitionNamesTheMonth(t *testing.T) {
	if got := MetricsPartition(time.Date(2026, 12, 31, 23, 0, 0, 0, time.UTC)); got != "fight_metrics_2026_12" {
		t.Fatalf("partition = %q", got)
	}
}

func TestMigrateCreatesTheSimulatorTables(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(t.Context(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	for table, columns := range map[string][]string{
		"sims": {"id", "user_id", "spec", "engine_version", "lane", "dps_mean",
			"dps_error", "iterations", "title", "result", "state", "created_at",
			// 0014's five: which tool produced the row, the line the
			// history list shows, and a bulk run's stage counters.
			"kind", "headline", "stage", "combos_done", "combos_total"},
		"sim_specs": {"spec", "state", "median_gap", "parses", "worst_actions",
			"engine_version", "updated_at"},
		"users":         {"premium"},
		"fight_metrics": {"execution_score"},
	} {
		for _, c := range columns {
			var n int
			if err := pool.QueryRow(t.Context(),
				`select count(*) from information_schema.columns
				 where table_name = $1 and column_name = $2`, table, c).Scan(&n); err != nil {
				t.Fatal(err)
			}
			if n != 1 {
				t.Errorf("%s.%s is missing", table, c)
			}
		}
	}

	// The indexes the filtered list queries were written against: without
	// them the kind filter and the owner filter are sequential scans.
	for _, index := range []string{
		"fight_metrics_execution_idx", "sims_user_kind_idx", "builds_user_idx",
	} {
		var n int
		if err := pool.QueryRow(t.Context(),
			`select count(*) from pg_indexes where indexname = $1`, index).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("%s is missing", index)
		}
	}

	// premium defaults to false: nothing in the code ever turns it on.
	if _, err := pool.Exec(t.Context(),
		`insert into users (email) values ('premium-default@example.com')
		 on conflict (email) do nothing`); err != nil {
		t.Fatal(err)
	}
	var premium bool
	if err := pool.QueryRow(t.Context(),
		`select premium from users where email = 'premium-default@example.com'`).
		Scan(&premium); err != nil {
		t.Fatal(err)
	}
	if premium {
		t.Error("a new account must not be premium")
	}
}

// TestMigration0016BackfillsOlderSimHeadlines pins the backfill in 0016
// against the Go function it has to agree with. Every row saved before
// 0014 introduced kind/headline is necessarily a plain run, so its
// headline is exactly what sims.Headline composes for that branch — the
// rounding rule (halves away from zero) and the thousands grouping
// included.
//
// The backfill lives in 0016 rather than in 0014 itself because 0014
// already shipped and ran against production; golang-migrate never
// re-runs an applied migration, so the fix had to be a follow-up
// migration instead of an edit to one already applied.
//
// The rows are inserted while the schema is at 0014/0015 (kind and
// headline exist, defaulted), so they are genuinely in the pre-0016
// shape rather than post-0016 rows with the column blanked.
func TestMigration0016BackfillsOlderSimHeadlines(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	// Whatever this test asserts, the shared database is left at the latest version.
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	cases := []struct {
		id     string
		mean   float64
		state  string
		preset string // non-"": headline is already set to this before the backfill runs, and must survive untouched
	}{
		{id: "mig16-zero", mean: 0, state: "done"},             // a zero is "0 DPS", not ""
		{id: "mig16-sub1k", mean: 942.4, state: "done"},        // under a thousand: no comma
		{id: "mig16-halfdown", mean: 1000.5, state: "done"},    // halves go away from zero,
		{id: "mig16-halfup", mean: 1001.5, state: "done"},      // ... never to even
		{id: "mig16-grouped", mean: 1234567.89, state: "done"}, // three-digit groups
		{id: "mig16-queued", mean: 1500, state: "queued"},      // no result yet: no headline
		{id: "mig16-errored", mean: 1500, state: "error"},      // no result at all: no headline
		// Idempotence: a done row that already carries a headline must not
		// be overwritten, whether the backfill runs once or is re-run.
		{id: "mig16-preset", mean: 1500, state: "done", preset: "999 DPS (preset, do not overwrite)"},
	}
	results := map[string]simapi.SimResult{}
	for _, c := range cases {
		results[c.id] = simapi.SimResult{
			SimID: c.id, EngineVersion: "0.0.0-test", Lane: simapi.LaneBrowser,
			Request: simapi.SimRequest{Spec: "warrior-fury"},
			DPS:     simapi.Estimate{Mean: c.mean},
		}
	}
	t.Cleanup(func() {
		p, err := Connect(context.Background(), url)
		if err != nil {
			t.Errorf("cleaning up the backfill rows: %v", err)
			return
		}
		defer p.Close()
		for _, c := range cases {
			if _, err := p.Exec(context.Background(), `delete from sims where id = $1`, c.id); err != nil {
				t.Errorf("cleaning up %s: %v", c.id, err)
			}
		}
	})

	migrateTo(t, url, 15)
	for _, c := range cases {
		body, err := json.Marshal(results[c.id])
		if err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(context.Background(),
			`insert into sims (id, spec, engine_version, lane, dps_mean, dps_error,
			   iterations, result, state, headline)
			 values ($1, 'warrior-fury', '0.0.0-test', 'browser', $2, 0, 1000, $3, $4, $5)`,
			c.id, c.mean, body, c.state, c.preset); err != nil {
			t.Fatal(err)
		}
	}

	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		var kind, headline string
		if err := pool.QueryRow(context.Background(),
			`select kind, headline from sims where id = $1`, c.id).Scan(&kind, &headline); err != nil {
			t.Fatal(err)
		}
		if kind != simapi.KindRun {
			t.Errorf("%s: kind = %q, want %q", c.id, kind, simapi.KindRun)
		}
		want := c.preset
		if want == "" && c.state == "done" {
			want = sims.Headline(results[c.id])
		}
		if headline != want {
			t.Errorf("%s (%v, %s): headline = %q, want %q", c.id, c.mean, c.state, headline, want)
		}
	}
}

func TestMigration0018DownReversesUp(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	if n := tableCount(t, pool, "guild_characters"); n != 1 {
		t.Fatal("guild_characters should exist at the latest migration")
	}
	for _, col := range []string{"consent", "verified_at"} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = 'guild_members' and column_name = $1`,
			col).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("guild_members.%s should exist at the latest migration", col)
		}
	}

	migrateTo(t, url, 17)
	if n := tableCount(t, pool, "guild_characters"); n != 0 {
		t.Error("guild_characters survived the down migration")
	}
	for _, col := range []string{"consent", "verified_at"} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = 'guild_members' and column_name = $1`,
			col).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("guild_members.%s survived the down migration", col)
		}
	}
	for _, col := range []string{"officer_max_rank_index", "claim_pending_by", "claim_requested_at", "invite_token_hash", "invite_token_rotated_at"} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = 'guilds' and column_name = $1`,
			col).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("guilds.%s survived the down migration", col)
		}
	}
	// 0018 must not reverse anything an earlier migration created.
	if n := tableCount(t, pool, "guild_members"); n != 1 {
		t.Error("the down migration took a table from an earlier migration")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := tableCount(t, pool, "guild_characters"); n != 1 {
		t.Error("guild_characters did not come back")
	}
}

func TestGuildCharactersCharacterKeyIsUniqueAcrossGuilds(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users, guilds cascade`); err != nil {
		t.Fatal(err)
	}
	var uid int64
	if err := pool.QueryRow(ctx,
		`insert into users (email) values ('unique-key@example.com') returning id`).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	var g1, g2 int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'One') returning id`).Scan(&g1); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Two') returning id`).Scan(&g2); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id) values ($1, 'us/normal/baelgrim', $2)`,
		g1, uid); err != nil {
		t.Fatal(err)
	}
	_, err = pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id) values ($1, 'us/normal/baelgrim', $2)`,
		g2, uid)
	if err == nil {
		t.Fatal("the same character_key under a second guild should violate the unique index")
	}
}

func TestMigration0018SecurityHardeningColumnsRoundTrip(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	for _, check := range []struct{ table, column string }{
		{"guild_characters", "verified_by"},
		{"guilds", "claim_contested_at"},
		{"guilds", "claim_contested_by"},
	} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = $1 and column_name = $2`,
			check.table, check.column).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("%s.%s should exist at the latest migration", check.table, check.column)
		}
	}
	if n := tableCount(t, pool, "guild_claim_attempts"); n != 1 {
		t.Fatal("guild_claim_attempts should exist at the latest migration")
	}
	if n := indexCount(t, pool, "guilds_region_ruleset_lower_name_idx"); n != 1 {
		t.Fatal("guilds_region_ruleset_lower_name_idx should exist at the latest migration")
	}

	migrateTo(t, url, 17)
	for _, check := range []struct{ table, column string }{
		{"guild_characters", "verified_by"},
		{"guilds", "claim_contested_at"},
		{"guilds", "claim_contested_by"},
	} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = $1 and column_name = $2`,
			check.table, check.column).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("%s.%s survived the down migration", check.table, check.column)
		}
	}
	if n := tableCount(t, pool, "guild_claim_attempts"); n != 0 {
		t.Error("guild_claim_attempts survived the down migration")
	}
	if n := indexCount(t, pool, "guilds_region_ruleset_lower_name_idx"); n != 0 {
		t.Error("guilds_region_ruleset_lower_name_idx survived the down migration")
	}
	// 0018's own base table must survive its own down migration's new
	// prefix - only the newly-appended block should be reversed here.
	if n := tableCount(t, pool, "guild_characters"); n != 0 {
		t.Error("guild_characters itself should still be gone at version 17 (this assertion documents the base table's own down migration still runs)")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := tableCount(t, pool, "guild_claim_attempts"); n != 1 {
		t.Error("guild_claim_attempts did not come back")
	}
}

func TestGuildsAreCaseInsensitiveByRegionAndRuleset(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate guilds cascade`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Iron Vanguard')`); err != nil {
		t.Fatal(err)
	}
	_, err = pool.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'IRON VANGUARD')`)
	if err == nil {
		t.Fatal("a case-variant name in the same region/ruleset should violate the case-insensitive unique index")
	}
	// A different ruleset, or a genuinely different name, is unaffected.
	if _, err := pool.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Iron Vanguard')`); err != nil {
		t.Fatalf("a different ruleset should not collide: %v", err)
	}
}
