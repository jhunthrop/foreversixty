package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const logCoachVersion = 55

var logCoachTables = []string{
	"coach_profiles", "coach_sim_refs", "coach_packs", "coach_reviews",
	"coach_review_leases", "llm_spend", "coach_account_day", "coach_review_samples",
}

func logCoachPool(t *testing.T) (*pgxpool.Pool, string) {
	t.Helper()
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool, url
}

func TestLogCoachMigration_UpDownUp(t *testing.T) {
	pool, url := logCoachPool(t)
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	check := func(stage string, want int) {
		for _, name := range logCoachTables {
			if n := tableCount(t, pool, name); n != want {
				t.Errorf("%s: table %s count = %d, want %d", stage, name, n, want)
			}
		}
	}
	check("after up", 1)
	migrateTo(t, url, logCoachVersion-1)
	check("after down", 0)
	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	check("after second up", 1)
}

func insertCoachProfile(ctx context.Context, pool *pgxpool.Pool, at time.Time) error {
	_, err := pool.Exec(ctx,
		`insert into coach_profiles (report_id, fight_index, player_key, class, spec, role,
		   encounter_id, difficulty, size, kill, duration_ms, phase, metric, fought_at, model_version, profile)
		 values ('r1', 0, 'p1', 'mage', 'fire', 'dps', 1, 1, 40, true, 1000, 'all', 1.5, $1, 'v1', '{}')`, at)
	return err
}

func TestLogCoachMigration_ProfilesArePartitioned(t *testing.T) {
	pool, _ := logCoachPool(t)
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate coach_profiles`); err != nil {
		t.Fatal(err)
	}
	// The migration creates no partition, so a far-future month has none until
	// EnsureCoachProfilesPartition makes it.
	far := time.Date(2041, 6, 15, 0, 0, 0, 0, time.UTC)
	if err := insertCoachProfile(ctx, pool, far); err == nil {
		t.Fatal("insert succeeded with no partition holding the row")
	}
	if err := EnsureCoachProfilesPartition(ctx, pool, far); err != nil {
		t.Fatal(err)
	}
	if err := insertCoachProfile(ctx, pool, far); err != nil {
		t.Fatalf("insert after EnsureCoachProfilesPartition: %v", err)
	}
	name := partitionName(CoachProfilesTable, far)
	var inTree bool
	if err := pool.QueryRow(ctx,
		`select exists (select 1 from pg_partition_tree('coach_profiles') where relid::text = $1 and isleaf)`,
		name).Scan(&inTree); err != nil {
		t.Fatal(err)
	}
	if !inTree {
		t.Errorf("%s is not a partition of coach_profiles", name)
	}
}

func TestLogCoachMigration_ReviewsCascadeWithTheReport(t *testing.T) {
	pool, _ := logCoachPool(t)
	ctx := context.Background()
	stmts := []string{
		`insert into reports (id, visibility, status) values ('coach-cascade', 'private', 'complete')`,
		`insert into coach_packs (report_id, fight_index, player_key, pack_version, pack)
		   values ('coach-cascade', 0, 'p', 'pv', '{}')`,
		`insert into coach_reviews (report_id, fight_index, player_key, pack_version, model_version, persona_version, review, source)
		   values ('coach-cascade', 0, 'p', 'pv', 'mv', 'ps', '{}', 'rules')`,
		`insert into coach_review_leases (report_id, fight_index, player_key, pack_version, model_version, persona_version, leased_until)
		   values ('coach-cascade', 0, 'p', 'pv', 'mv', 'ps', now())`,
		`insert into coach_review_samples (day, report_id, fight_index, player_key, character_name, spec, source, validation, pack, review)
		   values (current_date, 'coach-cascade', 0, 'p', 'Name', 'fire', 'rules', '[]', '{}', '{}')`,
	}
	for _, s := range stmts {
		if _, err := pool.Exec(ctx, s); err != nil {
			t.Fatalf("%s: %v", s, err)
		}
	}
	if _, err := pool.Exec(ctx, `delete from reports where id = 'coach-cascade'`); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"coach_packs", "coach_reviews", "coach_review_leases", "coach_review_samples"} {
		var n int
		if err := pool.QueryRow(ctx, `select count(*) from `+name+` where report_id = 'coach-cascade'`).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("%s kept %d rows after the report was deleted", name, n)
		}
	}
}

func TestPartitionJobSweepsCoachProfiles(t *testing.T) {
	pool, _ := logCoachPool(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	job := &PartitionJob{
		Pool: pool, Ahead: 1, Every: time.Hour, Tables: []string{"fight_metrics", CoachProfilesTable},
		Now: func() time.Time { return time.Date(2043, 3, 10, 0, 0, 0, 0, time.UTC) },
	}
	if err := job.Run(ctx); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"coach_profiles_2043_03", "coach_profiles_2043_04"} {
		if n := tableCount(t, pool, want); n != 1 {
			t.Errorf("partition %s missing", want)
		}
	}
}
