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
		"guilds", "guild_members", "characters", "uploads", "reports",
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

// TestMigration0014BackfillsOlderSimHeadlines pins the backfill in
// 0014 against the Go function it has to agree with. Every row saved
// before that migration is necessarily a plain run, so its headline is
// exactly what sims.Headline composes for that branch — the rounding
// rule (halves away from zero) and the thousands grouping included.
//
// The rows are inserted while the schema is at 0013, so they are
// genuinely in the pre-0014 shape rather than post-0014 rows with the
// column blanked.
func TestMigration0014BackfillsOlderSimHeadlines(t *testing.T) {
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
		id    string
		mean  float64
		state string
	}{
		{"mig14-zero", 0, "done"},             // a zero is "0 DPS", not ""
		{"mig14-sub1k", 942.4, "done"},        // under a thousand: no comma
		{"mig14-halfdown", 1000.5, "done"},    // halves go away from zero,
		{"mig14-halfup", 1001.5, "done"},      // ... never to even
		{"mig14-grouped", 1234567.89, "done"}, // three-digit groups
		{"mig14-queued", 1500, "queued"},      // no result yet: no headline
		{"mig14-errored", 1500, "error"},      // no result at all: no headline
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

	migrateTo(t, url, 13)
	for _, c := range cases {
		body, err := json.Marshal(results[c.id])
		if err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(context.Background(),
			`insert into sims (id, spec, engine_version, lane, dps_mean, dps_error,
			   iterations, result, state)
			 values ($1, 'warrior-fury', '0.0.0-test', 'browser', $2, 0, 1000, $3, $4)`,
			c.id, c.mean, body, c.state); err != nil {
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
		want := ""
		if c.state == "done" {
			want = sims.Headline(results[c.id])
		}
		if headline != want {
			t.Errorf("%s (%v, %s): headline = %q, want %q", c.id, c.mean, c.state, headline, want)
		}
	}
}
