// api/internal/billing/store_test.go
package billing

import (
	"context"
	"os"
	"testing"

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
		`truncate users, guilds, entitlements, stripe_customers, stripe_events cascade`); err != nil {
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
