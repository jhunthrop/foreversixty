package subscribe

import (
	"context"
	"os"
	"testing"

	"github.com/PLACEHOLDER/forever/api/internal/db"
)

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
	if _, err := pool.Exec(context.Background(), `truncate subscribers`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func TestStoreUpsertDedupesByNormalizedEmail(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	existing, err := s.Upsert(ctx, "Player@Example.com", "player@example.com", "tok1")
	if err != nil || existing {
		t.Fatalf("first upsert: existing=%v err=%v", existing, err)
	}
	existing, err = s.Upsert(ctx, "Player@Example.com", "player@example.com", "tok2")
	if err != nil || !existing {
		t.Fatalf("second upsert: existing=%v err=%v", existing, err)
	}
}

func TestStoreConfirmAndUnsubscribe(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok"); err != nil {
		t.Fatal(err)
	}
	if found, err := s.Confirm(ctx, "tok"); err != nil || !found {
		t.Fatalf("confirm known: found=%v err=%v", found, err)
	}
	if found, err := s.Confirm(ctx, "missing"); err != nil || found {
		t.Fatalf("confirm unknown: found=%v err=%v", found, err)
	}
	if found, err := s.Unsubscribe(ctx, "tok"); err != nil || !found {
		t.Fatalf("unsubscribe known: found=%v err=%v", found, err)
	}
	if found, err := s.Unsubscribe(ctx, "missing"); err != nil || found {
		t.Fatalf("unsubscribe unknown: found=%v err=%v", found, err)
	}
}
