package subscribe

import (
	"context"
	"os"
	"testing"
	"time"

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
	if _, err := pool.Exec(context.Background(), `truncate subscribers`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func TestStoreUpsertDedupesByNormalizedEmail(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	first, err := s.Upsert(ctx, "Player@Example.com", "player@example.com", "tok1", "utok1")
	if err != nil || first.Existing || first.Token != "" || first.Confirmed {
		t.Fatalf("first upsert: %+v err=%v", first, err)
	}
	second, err := s.Upsert(ctx, "Player@Example.com", "player@example.com", "tok2", "utok2")
	if err != nil || !second.Existing || second.Token != "tok1" || second.Confirmed {
		t.Fatalf("second upsert: %+v err=%v", second, err)
	}
}

func TestStoreUpsertReportsConfirmedOnExistingRow(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok", "utok"); err != nil {
		t.Fatal(err)
	}
	if found, err := s.Confirm(ctx, "tok"); err != nil || !found {
		t.Fatalf("confirm: found=%v err=%v", found, err)
	}
	res, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok2", "utok2")
	if err != nil || !res.Existing || res.Token != "tok" || !res.Confirmed || res.Unsubscribed {
		t.Fatalf("upsert after confirm: %+v err=%v", res, err)
	}
}

func TestStoreUpsertReportsUnsubscribedOnExistingRow(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok", "utok"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Confirm(ctx, "tok"); err != nil {
		t.Fatal(err)
	}
	if found, err := s.Unsubscribe(ctx, "utok"); err != nil || !found {
		t.Fatalf("unsubscribe: found=%v err=%v", found, err)
	}
	res, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok2", "utok2")
	if err != nil || !res.Existing || !res.Confirmed || !res.Unsubscribed {
		t.Fatalf("upsert after unsubscribe: %+v err=%v", res, err)
	}
}

func TestStoreClaimConfirmationSendSetsTimestampOnFirstClaimOnly(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok", "utok"); err != nil {
		t.Fatal(err)
	}

	fresh, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok2", "utok2")
	if err != nil {
		t.Fatal(err)
	}
	if fresh.ConfirmationSentAt != nil {
		t.Fatalf("confirmation_sent_at should start nil, got %v", fresh.ConfirmationSentAt)
	}

	before := time.Now().Add(-time.Second)
	claimed, err := s.ClaimConfirmationSend(ctx, "a@example.com", 15*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if !claimed {
		t.Fatal("expected the first claim on a never-sent address to succeed")
	}

	// Two immediate claims for the same address: the second must lose the
	// race against the cooldown the first just started.
	claimed, err = s.ClaimConfirmationSend(ctx, "a@example.com", 15*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if claimed {
		t.Fatal("expected an immediate second claim to fail (still within cooldown)")
	}

	after, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok3", "utok3")
	if err != nil {
		t.Fatal(err)
	}
	if after.ConfirmationSentAt == nil {
		t.Fatal("expected confirmation_sent_at to be set after the successful claim")
	}
	if after.ConfirmationSentAt.Before(before) {
		t.Fatalf("confirmation_sent_at = %v, want at/after %v", after.ConfirmationSentAt, before)
	}
}

func TestStoreClaimConfirmationSendSucceedsAgainAfterCooldownElapses(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok", "utok"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Pool.Exec(ctx, `update subscribers set confirmation_sent_at = now() - interval '20 minutes' where email_normalized = $1`, "a@example.com"); err != nil {
		t.Fatal(err)
	}

	claimed, err := s.ClaimConfirmationSend(ctx, "a@example.com", 15*time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if !claimed {
		t.Fatal("expected a claim to succeed once the previous send is older than the cooldown")
	}
}

func TestStoreClaimConfirmationSendIsIndependentPerAddress(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok-a", "utok-a"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Upsert(ctx, "b@example.com", "b@example.com", "tok-b", "utok-b"); err != nil {
		t.Fatal(err)
	}
	if claimed, err := s.ClaimConfirmationSend(ctx, "a@example.com", 15*time.Minute); err != nil || !claimed {
		t.Fatalf("claim for a@example.com: claimed=%v err=%v", claimed, err)
	}
	if claimed, err := s.ClaimConfirmationSend(ctx, "b@example.com", 15*time.Minute); err != nil || !claimed {
		t.Fatalf("claim for b@example.com should be unaffected by a's cooldown: claimed=%v err=%v", claimed, err)
	}
}

func TestStoreConfirmAndUnsubscribe(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok", "utok"); err != nil {
		t.Fatal(err)
	}
	if found, err := s.Confirm(ctx, "tok"); err != nil || !found {
		t.Fatalf("confirm known: found=%v err=%v", found, err)
	}
	if found, err := s.Confirm(ctx, "missing"); err != nil || found {
		t.Fatalf("confirm unknown: found=%v err=%v", found, err)
	}
	if found, err := s.Unsubscribe(ctx, "utok"); err != nil || !found {
		t.Fatalf("unsubscribe known: found=%v err=%v", found, err)
	}
	if found, err := s.Unsubscribe(ctx, "missing"); err != nil || found {
		t.Fatalf("unsubscribe unknown: found=%v err=%v", found, err)
	}
	// unsubscribing via the confirm token must not work: tokens are separate.
	if found, err := s.Unsubscribe(ctx, "tok"); err != nil || found {
		t.Fatalf("unsubscribe with confirm token: found=%v err=%v", found, err)
	}
}

func TestStoreConfirmClearsUnsubscribed(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok", "utok"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Confirm(ctx, "tok"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Unsubscribe(ctx, "utok"); err != nil {
		t.Fatal(err)
	}
	res, err := s.Upsert(ctx, "a@example.com", "a@example.com", "tok2", "utok2")
	if err != nil || !res.Unsubscribed {
		t.Fatalf("expected unsubscribed before re-confirm: %+v err=%v", res, err)
	}
	if found, err := s.Confirm(ctx, "tok"); err != nil || !found {
		t.Fatalf("re-confirm: found=%v err=%v", found, err)
	}
	res, err = s.Upsert(ctx, "a@example.com", "a@example.com", "tok3", "utok3")
	if err != nil || res.Unsubscribed {
		t.Fatalf("expected unsubscribed cleared by Confirm: %+v err=%v", res, err)
	}
}
