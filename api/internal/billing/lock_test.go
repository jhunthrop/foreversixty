package billing

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// smallPool is a pool capped at two connections: the shape CI's four-core
// runner has once two webhook deliveries are in flight, and the shape a
// busy production instance has under a burst. CI hung for ten minutes
// here: each delivery held a pooled connection for its event lock, then a
// second for its guild lock, then waited for a third to run the
// transaction, and neither could ever get it.
func smallPool(t *testing.T, max int32) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		t.Fatal(err)
	}
	cfg.MaxConns = max
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// Two callers each take an event lock, then a guild lock inside it, then
// run a query inside that: three connections' worth of work each, against
// a pool of two. Session-level advisory locks must not hold pooled
// connections, or this cannot complete.
func TestNestedAdvisoryLocksDoNotExhaustTheConnectionPool(t *testing.T) {
	pool := smallPool(t, 2)
	s := &Store{Pool: pool}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	errs := make([]error, 2)
	for i := range errs {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errs[i] = s.WithEventLock(ctx, "evt_pool_"+string(rune('a'+i)), func(ctx context.Context) error {
				return s.WithGuildLock(ctx, 424242, func(ctx context.Context) error {
					var one int
					return pool.QueryRow(ctx, `select 1`).Scan(&one)
				})
			})
		}(i)
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-ctx.Done():
		t.Fatal("deadlocked: the locks held every pooled connection")
	}
	for i, err := range errs {
		if err != nil {
			t.Errorf("caller %d: %v", i, err)
		}
	}
}
