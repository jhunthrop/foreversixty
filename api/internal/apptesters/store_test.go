package apptesters

import (
	"context"
	"os"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/db"
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
	if _, err := pool.Exec(context.Background(), `truncate app_testers`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func TestStoreAddsOnceAndThenReportsTheSignupAsExisting(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	created, err := s.Add(ctx, "a@example.com", PlatformIOS)
	if err != nil || !created {
		t.Fatalf("first add: created=%v err=%v", created, err)
	}
	created, err = s.Add(ctx, "a@example.com", PlatformIOS)
	if err != nil || created {
		t.Fatalf("second add: created=%v err=%v", created, err)
	}
	created, err = s.Add(ctx, "a@example.com", PlatformAndroid)
	if err != nil || !created {
		t.Fatalf("other platform: created=%v err=%v", created, err)
	}
}

func TestStoreListsNewestFirstWithPlatformAndDate(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Pool.Exec(ctx, `insert into app_testers (email, platform, created_at) values
		('old@example.com', 'android', now() - interval '2 days'),
		('new@example.com', 'ios', now())`); err != nil {
		t.Fatal(err)
	}
	got, err := s.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Email != "new@example.com" || got[0].Platform != PlatformIOS ||
		got[1].Email != "old@example.com" || got[1].Platform != PlatformAndroid {
		t.Fatalf("got %+v", got)
	}
	if got[0].CreatedAt.IsZero() || !got[0].CreatedAt.After(got[1].CreatedAt) {
		t.Fatalf("dates: %+v", got)
	}
}

func TestStoreRefusesAPlatformTheTableDoesNotKnow(t *testing.T) {
	s := testStore(t)
	if _, err := s.Add(context.Background(), "a@example.com", "windows"); err == nil {
		t.Fatal("want the check constraint to refuse windows")
	}
}
