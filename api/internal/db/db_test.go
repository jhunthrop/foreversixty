package db

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
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

	migrateTo(t, url, 4)
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
