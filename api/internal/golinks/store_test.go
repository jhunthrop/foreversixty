package golinks

import (
	"context"
	"errors"
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
	ctx := context.Background()
	// The starter links stay (they are the migration's); only what a test added goes.
	if _, err := pool.Exec(ctx, `truncate users, go_link_clicks cascade`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `delete from go_links where slug like 'test-%'`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func newUser(t *testing.T, s *Store, email, age string) int64 {
	t.Helper()
	var id int64
	if err := s.Pool.QueryRow(context.Background(),
		`insert into users (email, created_at) values ($1, now() - $2::interval) returning id`, email, age).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func TestStoreResolvesAStarterLinkAndNotAnUnknownOne(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	target, err := s.Target(ctx, "reddit-sim")
	if err != nil || target != "/sim" {
		t.Fatalf("Target(reddit-sim) = %q, %v", target, err)
	}
	if _, err := s.Target(ctx, "test-nothing"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Target(unknown) err = %v, want ErrNotFound", err)
	}
}

func TestStoreCountsOneRowPerLinkPerDayWithBotsApart(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	for _, bot := range []bool{false, false, true} {
		found, err := s.CountClick(ctx, "reddit", bot)
		if err != nil || !found {
			t.Fatalf("CountClick: found=%v err=%v", found, err)
		}
	}
	found, err := s.CountClick(ctx, "test-nothing", false)
	if err != nil || found {
		t.Fatalf("CountClick(unknown): found=%v err=%v", found, err)
	}
	var rows, clicks, bots int64
	if err := s.Pool.QueryRow(ctx, `select count(*), sum(clicks), sum(bot_clicks) from go_link_clicks`).Scan(&rows, &clicks, &bots); err != nil {
		t.Fatal(err)
	}
	if rows != 1 || clicks != 2 || bots != 1 {
		t.Fatalf("rows=%d clicks=%d bots=%d, want 1, 2, 1", rows, clicks, bots)
	}
}

func TestStoreRecordsAnArrivalOnceForANewAccountOnly(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	fresh := newUser(t, s, "new@example.com", "5 minutes")
	old := newUser(t, s, "old@example.com", "3 days")

	if ok, err := s.RecordArrival(ctx, fresh, "test-nothing"); err != nil || ok {
		t.Fatalf("an unknown slug: ok=%v err=%v", ok, err)
	}
	if ok, err := s.RecordArrival(ctx, fresh, "reddit"); err != nil || !ok {
		t.Fatalf("first arrival: ok=%v err=%v", ok, err)
	}
	if ok, err := s.RecordArrival(ctx, fresh, "discord"); err != nil || ok {
		t.Fatalf("a second arrival overwrote the first: ok=%v err=%v", ok, err)
	}
	if ok, err := s.RecordArrival(ctx, old, "reddit"); err != nil || ok {
		t.Fatalf("a returning account was attributed: ok=%v err=%v", ok, err)
	}
	var slug string
	if err := s.Pool.QueryRow(ctx, `select first_touch_slug from users where id = $1`, fresh).Scan(&slug); err != nil || slug != "reddit" {
		t.Fatalf("first_touch_slug = %q (%v), want reddit", slug, err)
	}
}

func TestStoreCreatesALinkOnce(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	link, err := s.Create(ctx, Link{Slug: "test-new", Target: "/guides", Label: "New"})
	if err != nil || link.Slug != "test-new" || link.CreatedAt.IsZero() {
		t.Fatalf("Create = %+v, %v", link, err)
	}
	if _, err := s.Create(ctx, Link{Slug: "test-new", Target: "/", Label: "Again"}); !errors.Is(err, ErrExists) {
		t.Fatalf("a second create err = %v, want ErrExists", err)
	}
}

func TestStoreFunnelFollowsEachLinksAccounts(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	if _, err := s.Create(ctx, Link{Slug: "test-funnel", Target: "/", Label: "Funnel"}); err != nil {
		t.Fatal(err)
	}
	for _, bot := range []bool{false, false, false, true} {
		if _, err := s.CountClick(ctx, "test-funnel", bot); err != nil {
			t.Fatal(err)
		}
	}
	plain := newUser(t, s, "a@example.com", "1 minute")
	linked := newUser(t, s, "b@example.com", "1 minute")
	paying := newUser(t, s, "c@example.com", "1 minute")
	for _, id := range []int64{plain, linked, paying} {
		if _, err := s.RecordArrival(ctx, id, "test-funnel"); err != nil {
			t.Fatal(err)
		}
	}
	for i, id := range []int64{linked, paying} {
		if _, err := s.Pool.Exec(ctx, `insert into characters (key, region, ruleset, name, user_id)
			values ($1, 'us', 'normal', 'Name', $2)`, []string{"us/normal/a", "us/normal/b"}[i], id); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := s.Pool.Exec(ctx, `insert into entitlements (user_id, plan, source, status)
		values ($1, 'premium', 'grant', 'active')`, paying); err != nil {
		t.Fatal(err)
	}

	rows, err := s.Funnel(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var got *FunnelRow
	var reddit *FunnelRow
	for i := range rows {
		switch rows[i].Slug {
		case "test-funnel":
			got = &rows[i]
		case "reddit":
			reddit = &rows[i]
		}
	}
	if got == nil || reddit == nil {
		t.Fatalf("funnel is missing a link: %+v", rows)
	}
	want := FunnelRow{Link: got.Link, Clicks: 3, BotClicks: 1, SignIns: 3, Linked: 2, Premium: 1}
	if *got != want {
		t.Fatalf("funnel row = %+v, want %+v", *got, want)
	}
	if reddit.Clicks != 0 || reddit.SignIns != 0 {
		t.Fatalf("an untouched link has figures: %+v", *reddit)
	}
}
