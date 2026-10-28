package coach

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

var (
	poolNow      = time.Date(2026, 10, 27, 15, 0, 0, 0, time.UTC)
	poolFoughtAt = time.Date(2026, 10, 20, 20, 0, 0, 0, time.UTC)
	onyxia40Fury = PoolKey{EncounterID: 1084, Difficulty: 0, Size: 40, Spec: "Fury", Phase: "p1"}
	poolReviewed = "us/normal/other"
)

func poolTestDB(t *testing.T) *pgxpool.Pool {
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
	ctx := context.Background()
	if err := db.EnsureCoachProfilesPartition(ctx, pool, poolFoughtAt); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `truncate coach_profiles, reports cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// poolRow is one coach_profiles row to seed; zero fields take the Onyxia 40 Fury kill.
type poolRow struct {
	key    PoolKey
	player string
	metric float64
	cpm    float64
	report string
}

func seedPoolRow(t *testing.T, pool *pgxpool.Pool, r poolRow) {
	t.Helper()
	ctx := context.Background()
	if r.report == "" {
		r.report = fmt.Sprintf("rep-%s-%v-%d", r.player, r.metric, time.Now().UnixNano())
	}
	if _, err := pool.Exec(ctx, `insert into reports (id, visibility, status) values ($1, 'public', 'complete')
		on conflict do nothing`, r.report); err != nil {
		t.Fatal(err)
	}
	profile, err := json.Marshal(cpmProfile(r.cpm))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into coach_profiles
		(report_id, fight_index, player_key, class, spec, role, encounter_id, difficulty, size, kill,
		 duration_ms, phase, metric, fought_at, model_version, profile)
		values ($1, $2, $3, 'Warrior', $4, 'dps', $5, $6, $7, true, 100000, $8, $9, $10, 'v1', $11)`,
		r.report, time.Now().Nanosecond(), r.player, r.key.Spec, r.key.EncounterID, r.key.Difficulty, r.key.Size,
		r.key.Phase, r.metric, poolFoughtAt, profile); err != nil {
		t.Fatal(err)
	}
}

// seedPlayers seeds n players p1..pn, one kill each, metric and cpm equal to the index.
func seedPlayers(t *testing.T, pool *pgxpool.Pool, k PoolKey, n int) {
	t.Helper()
	for i := 1; i <= n; i++ {
		seedPoolRow(t, pool, poolRow{key: k, player: fmt.Sprintf("us/normal/p%d", i), metric: float64(i), cpm: float64(i)})
	}
}

func poolAt(pool *pgxpool.Pool) *Pool { return &Pool{DB: pool, Now: func() time.Time { return poolNow }} }

func TestPool_NineteenPlayersIsNoBasis(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 19)
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 19 || ref != nil {
		t.Fatalf("n=%d ref=%v err=%v", n, ref, err)
	}
	in := withPool(furyInput(t), Refs{PoolCount: n})
	_, b, thin := comparison(in, func(p Profile) (float64, bool) { return p.CastsPerMinOf("Bloodthirst") }, nil, false)
	want := "not enough logs yet — 19 Fury kills of Onyxia (40) on the site, the coach needs 20."
	if b != nil || thin != want {
		t.Fatalf("basis %v thin %q; want %q", b, thin, want)
	}
}

func TestPool_TwentyPlayersMedianOfTop25(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 30)
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 30 || ref == nil || ref.N != PoolTopN {
		t.Fatalf("n=%d ref=%+v err=%v", n, ref, err)
	}
	if v, _ := ref.Median.CastsPerMinOf("Bloodthirst"); v != 18 {
		t.Fatalf("median of metrics 6..30 = %v; want 18", v)
	}
}

func TestPool_ExactlyTwentyIsABasis(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 20)
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 20 || ref == nil || ref.N != 20 {
		t.Fatalf("n=%d ref=%+v err=%v", n, ref, err)
	}
}

func TestPool_OneBestKillPerPlayer(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 19)
	for i := 1; i <= 5; i++ {
		seedPoolRow(t, pool, poolRow{key: onyxia40Fury, player: "us/normal/many", metric: float64(i * 10), cpm: float64(i * 10)})
	}
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 20 || ref == nil {
		t.Fatalf("five kills by one player count once: n=%d ref=%v err=%v", n, ref, err)
	}
	// 19 players at 1..19 plus the one at its best (50): median of 20 is (10+11)/2.
	if v, _ := ref.Median.CastsPerMinOf("Bloodthirst"); v != 10.5 {
		t.Fatalf("median = %v; the multi-kill player must count at its best (50), want 10.5", v)
	}
}

func TestPool_ExcludesTheReviewedPlayer(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 19)
	seedPoolRow(t, pool, poolRow{key: onyxia40Fury, player: poolReviewed, metric: 99, cpm: 99})
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 19 || ref != nil {
		t.Fatalf("20 players including the reviewed one: n=%d ref=%v err=%v", n, ref, err)
	}
}

func TestPool_ExcludedInTheCachedTopTakesTheTwentySixth(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 30)
	p := poolAt(pool)
	n, ref, err := p.Reference(context.Background(), onyxia40Fury, "us/normal/p30")
	if err != nil || n != 29 || ref == nil || ref.N != PoolTopN {
		t.Fatalf("n=%d ref=%+v err=%v", n, ref, err)
	}
	// players 5..29 remain in the top 25: median 17.
	if v, _ := ref.Median.CastsPerMinOf("Bloodthirst"); v != 17 {
		t.Fatalf("median = %v; want 17", v)
	}
	n, _, _ = p.Reference(context.Background(), onyxia40Fury, "us/normal/p1")
	if n != 29 {
		t.Fatalf("an excluded player below the top still leaves the count at 29, got %d", n)
	}
}

func TestPool_PhaseSizeAndDifficultyBound(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 19)
	other := []PoolKey{onyxia40Fury, onyxia40Fury, onyxia40Fury, onyxia40Fury}
	other[0].Phase, other[1].Size, other[2].Difficulty, other[3].Spec = "p2", 20, 1, "Arms"
	for i, k := range other {
		seedPoolRow(t, pool, poolRow{key: k, player: fmt.Sprintf("us/normal/o%d", i), metric: 5, cpm: 5})
	}
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 19 || ref != nil {
		t.Fatalf("rows of another phase, size, difficulty or spec leaked in: n=%d ref=%v err=%v", n, ref, err)
	}
}

func TestPool_FlaggedReportsLeaveThePool(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 20)
	if _, err := pool.Exec(context.Background(),
		`update reports set flagged = 'tampered' where id = (select report_id from coach_profiles where player_key = 'us/normal/p3')`); err != nil {
		t.Fatal(err)
	}
	n, ref, err := poolAt(pool).Reference(context.Background(), onyxia40Fury, poolReviewed)
	if err != nil || n != 19 || ref != nil {
		t.Fatalf("a flagged report's player stays out: n=%d ref=%v err=%v", n, ref, err)
	}
}

func TestPool_CachesForAnHour(t *testing.T) {
	pool := poolTestDB(t)
	seedPlayers(t, pool, onyxia40Fury, 30)
	now := poolNow
	p := &Pool{DB: pool, Now: func() time.Time { return now }}
	ctx := context.Background()
	n, _, err := p.Reference(ctx, onyxia40Fury, poolReviewed)
	if err != nil || n != 30 {
		t.Fatalf("n=%d err=%v", n, err)
	}
	if _, err := pool.Exec(ctx, `delete from coach_profiles`); err != nil {
		t.Fatal(err)
	}
	now = poolNow.Add(59 * time.Minute)
	if n, ref, _ := p.Reference(ctx, onyxia40Fury, poolReviewed); n != 30 || ref == nil {
		t.Fatalf("within the TTL the cached pool answers: n=%d", n)
	}
	now = poolNow.Add(time.Hour)
	if n, ref, _ := p.Reference(ctx, onyxia40Fury, poolReviewed); n != 0 || ref != nil {
		t.Fatalf("after the TTL the pool is re-read: n=%d", n)
	}
}
