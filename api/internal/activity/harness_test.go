package activity

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

var testNow = time.Date(2026, 11, 7, 18, 0, 0, 0, time.UTC)

func testStore(t *testing.T) *Store {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(), `truncate users, guilds, reports, fight_metrics, characters,
		activity_events, bis_pick_state, tier_grade_weeks, spec_board_weeks, character_sims, sims cascade`); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureMetricsPartitions(context.Background(), pool, testNow.AddDate(0, -1, 0), 3); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func exec(t *testing.T, pool *pgxpool.Pool, sql string, args ...any) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
}

func user(t *testing.T, s *Store, tag string, anonymize bool) int64 {
	t.Helper()
	var id int64
	if err := s.Pool.QueryRow(context.Background(),
		`insert into users (battletag, anonymize) values ($1, $2) returning id`, tag, anonymize).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func char(t *testing.T, s *Store, key, name, class string, level int, userID int64) {
	t.Helper()
	exec(t, s.Pool, `insert into characters (key, region, ruleset, name, class, level, user_id) values ($1, 'us', 'normal', $2, $3, $4, $5)`,
		key, name, class, level, userID)
}

func guild(t *testing.T, s *Store, name string) int64 {
	t.Helper()
	var id int64
	if err := s.Pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`, name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func member(t *testing.T, s *Store, guildID int64, key string, userID int64) {
	t.Helper()
	exec(t, s.Pool, `insert into guild_characters (guild_id, character_key, user_id, verified_at) values ($1, $2, $3, now())`,
		guildID, key, userID)
	exec(t, s.Pool, `insert into guild_members (guild_id, user_id, verified_at) values ($1, $2, now()) on conflict do nothing`,
		guildID, userID)
}

func ev(kind, dedupe string, at time.Time) Event {
	return Event{Kind: kind, OccurredAt: at, TagBase: TagLogs, Segments: []Segment{Text(dedupe)}, DedupeKey: dedupe}
}

func mustAppend(t *testing.T, s *Store, e Event) {
	t.Helper()
	if _, err := s.Append(context.Background(), e); err != nil {
		t.Fatal(err)
	}
}
