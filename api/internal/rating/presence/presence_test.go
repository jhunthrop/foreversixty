// api/internal/rating/presence/presence_test.go
package presence

import (
	"context"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// Task L6.8: presence is read by character key from the rated and ranked rows, never from
// fights.players (the log's GUIDs). Every row this test writes carries the "l68p-" prefix.
func TestInReadsRatedAndRankedRowsByCharacterKey(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	t.Cleanup(func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id like 'l68p-%'`)
		pool.Exec(bg, `delete from fight_metrics where report_id like 'l68p-%'`)
		pool.Exec(bg, `delete from reports where id like 'l68p-%'`)
	})
	at := time.Date(2026, 10, 9, 23, 30, 0, 0, time.UTC)
	if err := db.EnsureMetricsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"l68p-a", "l68p-b", "l68p-c"} {
		mustExec(t, pool, `insert into reports (id, visibility, status, created_at) values ($1, 'guild', 'complete', $2)`, id, at)
		// fights.players holds the log's GUIDs, as engine.go writes it.
		mustExec(t, pool, `insert into fights (report_id, fight_index, encounter_id, kill, players)
			values ($1, 0, 1084, true, $2)`, id, []string{"Player-4395-0A1B2C3D"})
	}
	const ranked, rated, removed = "/Classic Beta PvP/L68ranked", "/Classic Beta PvP/L68rated", "/Classic Beta PvP/L68removed"
	mustExec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, fought_at) values ('l68p-a', 0, $1, $2)`, ranked, at)
	mustExec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, fought_at) values ('l68p-c', 0, $1, $2)`, ranked, at)
	mustExec(t, pool, `insert into rating_scores (report_id, fight_index, player_key, overall, components, model_version, fought_at, kill)
		values ('l68p-b', 0, $1, 50, '[]', 'l68', $2, true)`, rated, at)
	mustExec(t, pool, `insert into fight_metrics (report_id, fight_index, player_key, fought_at, state)
		values ('l68p-a', 0, $1, $2, 'removed')`, removed, at)

	got, err := In(ctx, pool, []string{"l68p-a", "l68p-b"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	var seen []string
	for _, p := range got {
		seen = append(seen, p.Key+" "+p.ReportID+" "+p.Night.Format("2006-01-02"))
	}
	sort.Strings(seen)
	want := []string{ranked + " l68p-a 2026-10-09", rated + " l68p-b 2026-10-09"}
	if len(seen) != len(want) || seen[0] != want[0] || seen[1] != want[1] {
		t.Errorf("In(a, b) = %v, want %v (a removed row and a report outside the set are absent, and no GUID is a key)", seen, want)
	}

	only, err := In(ctx, pool, []string{"l68p-a", "l68p-b", "l68p-c"}, []string{ranked})
	if err != nil {
		t.Fatal(err)
	}
	if len(only) != 2 || only[0].Key != ranked || only[1].Key != ranked {
		t.Errorf("In restricted to %s = %+v, want its two reports", ranked, only)
	}

	none, err := In(ctx, pool, nil, nil)
	if err != nil || len(none) != 0 {
		t.Errorf("In over no reports = %+v, %v; want nothing", none, err)
	}
}

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func mustExec(t *testing.T, pool *pgxpool.Pool, sql string, args ...any) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
}
