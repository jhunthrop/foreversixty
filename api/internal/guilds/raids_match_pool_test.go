package guilds

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// oneConnStore is the harness's store on a pool of one connection: a match that asked the pool
// for a second connection while its transaction held the first would wait forever.
func oneConnStore(t *testing.T, w raidWorld) *Store {
	t.Helper()
	cfg, err := pgxpool.ParseConfig(os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	cfg.MaxConns = 1
	pool, err := pgxpool.NewWithConfig(t.Context(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	s := *w.h.store
	s.Pool = pool
	return &s
}

func TestAMatchHoldsOneConnection(t *testing.T) {
	w, raid := attendanceNight(t)
	s := oneConnStore(t, w)
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()
	if _, err := s.SetMatch(ctx, w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatalf("hand match on a one-connection pool: %v", err)
	}
	if rows := w.attendance(t, raid.ID); len(rows) == 0 {
		t.Fatal("the hand match computed no attendance")
	}
	if _, err := s.ClearMatch(ctx, w.gid, raid.ID, w.officer); err != nil {
		t.Fatal(err)
	}
	w.exec(t, `delete from guild_raid_rejected_reports`)
	if _, ok, err := s.AutoMatch(ctx, "night"); err != nil || !ok {
		t.Fatalf("auto match on a one-connection pool = %v %v", ok, err)
	}
}
