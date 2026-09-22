// api/internal/billing/store_test.go
package billing

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

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
	if _, err := pool.Exec(context.Background(),
		`truncate users, guilds, entitlements, stripe_customers, stripe_events,
		 pending_checkouts, entitlement_anomalies cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func seedUser(t *testing.T, pool *pgxpool.Pool, email string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`, email).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func seedGuild(t *testing.T, pool *pgxpool.Pool, name string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1) returning id`,
		name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func TestCustomerIDSavesAndReads(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	uid := seedUser(t, pool, "customer-id@example.com")
	if _, ok, err := s.CustomerID(context.Background(), uid); err != nil || ok {
		t.Fatalf("before save: %v, %v", ok, err)
	}
	if err := s.SaveCustomerID(context.Background(), uid, "cus_abc"); err != nil {
		t.Fatal(err)
	}
	id, ok, err := s.CustomerID(context.Background(), uid)
	if err != nil || !ok || id != "cus_abc" {
		t.Fatalf("after save: %v, %v, %v", id, ok, err)
	}
	// Re-saving (a reused customer, or a retried checkout) updates in place.
	if err := s.SaveCustomerID(context.Background(), uid, "cus_def"); err != nil {
		t.Fatal(err)
	}
	if id, _, _ := s.CustomerID(context.Background(), uid); id != "cus_def" {
		t.Fatalf("after re-save: %v", id)
	}
}

func TestRecordEventOnceIsIdempotent(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	proceed, err := s.RecordEventOnce(context.Background(), "evt_1", "checkout.session.completed", []byte(`{}`))
	if err != nil || !proceed {
		t.Fatalf("first insert: %v, %v", proceed, err)
	}
	if err := s.MarkEventProcessed(context.Background(), "evt_1"); err != nil {
		t.Fatal(err)
	}
	proceed, err = s.RecordEventOnce(context.Background(), "evt_1", "checkout.session.completed", []byte(`{}`))
	if err != nil || proceed {
		t.Fatalf("redelivery of a processed event: %v, %v, want proceed=false", proceed, err)
	}
	var processedAt *string
	if err := pool.QueryRow(context.Background(),
		`select processed_at::text from stripe_events where id = 'evt_1'`).Scan(&processedAt); err != nil {
		t.Fatal(err)
	}
	if processedAt == nil {
		t.Fatal("processed_at should be set after MarkEventProcessed")
	}
}

// TestWithGuildLockSerializesTwoCallers is the advisory-lock half of the
// security review's double-billing fix (2026-09-21): two concurrent
// callers for the same guild id must never run fn at the same time —
// this is what makes checkGuildCheckout's entitlement read through
// Checkout Session creation race-free against a second officer's
// concurrent checkout.
func TestWithGuildLockSerializesTwoCallers(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	var inFlight int32
	var maxInFlight int32
	fn := func(ctx context.Context) error {
		n := atomic.AddInt32(&inFlight, 1)
		for {
			max := atomic.LoadInt32(&maxInFlight)
			if n <= max || atomic.CompareAndSwapInt32(&maxInFlight, max, n) {
				break
			}
		}
		time.Sleep(50 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)
		return nil
	}
	var wg sync.WaitGroup
	errs := make([]error, 2)
	wg.Add(2)
	go func() { defer wg.Done(); errs[0] = s.WithGuildLock(context.Background(), 42, fn) }()
	go func() { defer wg.Done(); errs[1] = s.WithGuildLock(context.Background(), 42, fn) }()
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	if got := atomic.LoadInt32(&maxInFlight); got != 1 {
		t.Fatalf("max concurrent callers inside the lock = %d, want 1", got)
	}
}

// TestWithGuildLockDoesNotSerializeDifferentGuilds confirms the lock is
// scoped per guild id, not a single global lock that would serialize
// every guild's checkout against every other's.
func TestWithGuildLockDoesNotSerializeDifferentGuilds(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	release := make(chan struct{})
	started := make(chan struct{})
	go func() {
		_ = s.WithGuildLock(context.Background(), 100, func(ctx context.Context) error {
			close(started)
			<-release
			return nil
		})
	}()
	<-started
	done := make(chan struct{})
	go func() {
		_ = s.WithGuildLock(context.Background(), 200, func(ctx context.Context) error { return nil })
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("a different guild id must not be blocked by guild 100's held lock")
	}
	close(release)
}

func TestPendingCheckoutRecordExistsDeleteRoundTrip(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "pending-checkout")
	uid := seedUser(t, pool, "pending-checkout@example.com")

	if ok, err := s.PendingGuildCheckout(context.Background(), gid); err != nil || ok {
		t.Fatalf("before any record: %v, %v", ok, err)
	}
	future := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	if err := s.RecordPendingCheckout(context.Background(), gid, uid, "cs_pending_1", future); err != nil {
		t.Fatal(err)
	}
	if ok, err := s.PendingGuildCheckout(context.Background(), gid); err != nil || !ok {
		t.Fatalf("after record: %v, %v", ok, err)
	}
	if err := s.DeletePendingCheckout(context.Background(), "cs_pending_1"); err != nil {
		t.Fatal(err)
	}
	if ok, err := s.PendingGuildCheckout(context.Background(), gid); err != nil || ok {
		t.Fatalf("after delete: %v, %v", ok, err)
	}
}

func TestPendingCheckoutExpiredRowDoesNotCountAsPending(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "pending-expired")
	uid := seedUser(t, pool, "pending-expired@example.com")
	past := time.Now().Add(-time.Hour).UTC().Truncate(time.Second)
	if err := s.RecordPendingCheckout(context.Background(), gid, uid, "cs_expired_1", past); err != nil {
		t.Fatal(err)
	}
	if ok, err := s.PendingGuildCheckout(context.Background(), gid); err != nil || ok {
		t.Fatalf("expired row: %v, %v, want ok=false", ok, err)
	}
}

func TestSweepExpiredPendingCheckoutsRemovesOnlyExpiredRows(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	gid := seedGuild(t, pool, "sweep-guild")
	uid := seedUser(t, pool, "sweep@example.com")
	past := time.Now().Add(-time.Hour).UTC().Truncate(time.Second)
	future := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	if err := s.RecordPendingCheckout(context.Background(), gid, uid, "cs_sweep_expired", past); err != nil {
		t.Fatal(err)
	}
	if err := s.RecordPendingCheckout(context.Background(), gid, uid, "cs_sweep_live", future); err != nil {
		t.Fatal(err)
	}
	n, err := s.SweepExpiredPendingCheckouts(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("swept = %d, want 1", n)
	}
	var remaining int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from pending_checkouts where guild_id = $1`, gid).Scan(&remaining); err != nil {
		t.Fatal(err)
	}
	if remaining != 1 {
		t.Fatalf("remaining rows = %d, want 1 (the live one)", remaining)
	}
}

func TestRecordEventOnceRetriesAnEventRecordedButNeverProcessed(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	proceed, err := s.RecordEventOnce(context.Background(), "evt_stuck", "checkout.session.completed", []byte(`{}`))
	if err != nil || !proceed {
		t.Fatalf("first insert: %v, %v", proceed, err)
	}
	// Simulate a prior attempt that recorded the event but crashed
	// before calling MarkEventProcessed — processed_at stays null.
	proceed, err = s.RecordEventOnce(context.Background(), "evt_stuck", "checkout.session.completed", []byte(`{}`))
	if err != nil || !proceed {
		t.Fatalf("retry of a never-processed event: %v, %v, want proceed=true", proceed, err)
	}
}
