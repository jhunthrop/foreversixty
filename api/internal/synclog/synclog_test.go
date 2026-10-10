package synclog

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func at(hours ...int) []time.Time {
	base := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	out := make([]time.Time, len(hours))
	for i, h := range hours {
		out[i] = base.Add(time.Duration(h) * time.Hour)
	}
	return out
}

func TestMedianGapSecIsNilBelowThreeSyncs(t *testing.T) {
	for _, n := range []int{0, 1, 2} {
		if got := MedianGapSec(at(make([]int, n)...)); got != nil {
			t.Fatalf("%d syncs: median = %d, want nil", n, *got)
		}
	}
}

func TestMedianGapSecOddGapCount(t *testing.T) {
	// gaps 1h, 10h, 2h -> sorted 1, 2, 10 -> median 2h
	got := MedianGapSec(at(0, 1, 11, 13))
	if got == nil || *got != 2*3600 {
		t.Fatalf("median = %v, want 7200", got)
	}
}

func TestMedianGapSecEvenGapCountAveragesTheMiddlePair(t *testing.T) {
	// three syncs -> gaps 1h, 3h -> median 2h
	got := MedianGapSec(at(0, 1, 4))
	if got == nil || *got != 2*3600 {
		t.Fatalf("median = %v, want 7200", got)
	}
}

func TestMedianGapSecIgnoresInputOrder(t *testing.T) {
	got := MedianGapSec(at(4, 0, 1))
	if got == nil || *got != 2*3600 {
		t.Fatalf("median = %v, want 7200", got)
	}
}

func testPool(t *testing.T) *pgxpool.Pool {
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
	if _, err := pool.Exec(context.Background(), `truncate users, characters cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func seedCharacter(t *testing.T, pool *pgxpool.Pool, key string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name) values ($1, 'us', 'pvp', 'Kiloz')`, key); err != nil {
		t.Fatal(err)
	}
}

func insertAt(t *testing.T, pool *pgxpool.Pool, key, source, outcome string, when time.Time) {
	t.Helper()
	if _, err := pool.Exec(context.Background(),
		`insert into character_syncs (character_key, source, outcome, created_at) values ($1, $2, $3, $4)`,
		key, source, outcome, when); err != nil {
		t.Fatal(err)
	}
}

func TestSummariesReadsMedianFromOkRowsAndErrorFromLatestBlizzardRow(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	seedCharacter(t, pool, "us/pvp/kiloz")
	seedCharacter(t, pool, "us/pvp/quiet")
	seedCharacter(t, pool, "us/pvp/healed")
	times := at(0, 1, 4)
	for _, w := range times {
		insertAt(t, pool, "us/pvp/kiloz", "addon", OutcomeOK, w)
	}
	insertAt(t, pool, "us/pvp/kiloz", "blizzard", OutcomeRefreshFailed, times[2].Add(time.Hour))
	insertAt(t, pool, "us/pvp/quiet", "addon", OutcomeOK, times[0])
	insertAt(t, pool, "us/pvp/healed", "blizzard", OutcomeRefreshFailed, times[0])
	insertAt(t, pool, "us/pvp/healed", "blizzard", OutcomeOK, times[1])

	got, err := Summaries(ctx, pool, []string{"us/pvp/kiloz", "us/pvp/quiet", "us/pvp/healed", "us/pvp/none"})
	if err != nil {
		t.Fatal(err)
	}
	kiloz := got["us/pvp/kiloz"]
	if kiloz.MedianGapSec == nil || *kiloz.MedianGapSec != 2*3600 {
		t.Fatalf("kiloz median = %v, want 7200 (the failed row must not count as a sync)", kiloz.MedianGapSec)
	}
	if kiloz.SyncError == nil || *kiloz.SyncError != OutcomeRefreshFailed {
		t.Fatalf("kiloz sync_error = %v, want %s", kiloz.SyncError, OutcomeRefreshFailed)
	}
	if q := got["us/pvp/quiet"]; q.MedianGapSec != nil || q.SyncError != nil {
		t.Fatalf("quiet = %+v, want both nil", q)
	}
	if h := got["us/pvp/healed"]; h.SyncError != nil {
		t.Fatalf("healed sync_error = %v, want nil after a later successful refresh", *h.SyncError)
	}
	if _, ok := got["us/pvp/none"]; ok {
		t.Fatal("a character with no history must have no entry")
	}
}

func TestSummariesAnAddonSyncDoesNotClearARefreshError(t *testing.T) {
	pool := testPool(t)
	seedCharacter(t, pool, "us/pvp/kiloz")
	times := at(0, 1)
	insertAt(t, pool, "us/pvp/kiloz", "blizzard", OutcomeRefreshFailed, times[0])
	insertAt(t, pool, "us/pvp/kiloz", "addon", OutcomeOK, times[1])
	got, err := Summaries(context.Background(), pool, []string{"us/pvp/kiloz"})
	if err != nil {
		t.Fatal(err)
	}
	if e := got["us/pvp/kiloz"].SyncError; e == nil || *e != OutcomeRefreshFailed {
		t.Fatalf("sync_error = %v, want it kept: only a Battle.net refresh answers a Battle.net refresh", e)
	}
}

func TestRecordWritesOneRow(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	seedCharacter(t, pool, "us/pvp/kiloz")
	if err := Record(ctx, pool, "us/pvp/kiloz", SourceAddon, OutcomeOK); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from character_syncs where character_key = 'us/pvp/kiloz'`).Scan(&n); err != nil || n != 1 {
		t.Fatalf("rows = %d, err = %v, want 1", n, err)
	}
}

func TestRecordRejectsAnUnknownSource(t *testing.T) {
	if err := Record(context.Background(), nil, "us/pvp/kiloz", "carrier-pigeon", OutcomeOK); err == nil {
		t.Fatal("want an error for an unknown source")
	}
}
