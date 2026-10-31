package app

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func TestStoreReadsTheExport(t *testing.T) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	pool, err := db.Connect(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	var owner int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('app-home@example.com')
		on conflict (email) do update set email = excluded.email returning id`).Scan(&owner); err != nil {
		t.Fatal(err)
	}
	at := time.Date(2026, 10, 29, 21, 15, 0, 0, time.UTC)
	if _, err := pool.Exec(ctx, `delete from addon_exports where character_key = $1`, homeKey); err != nil {
		t.Fatal(err)
	}
	s := &Store{Pool: pool}
	if _, found, err := s.Export(ctx, homeKey); err != nil || found {
		t.Fatalf("no export: found %v err %v", found, err)
	}
	if _, err := pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, 'us', 'normal', 'velaren', 'FS1:x', 'addon', $3, $3)`, homeKey, owner, at); err != nil {
		t.Fatal(err)
	}
	e, found, err := s.Export(ctx, homeKey)
	if err != nil || !found || e.Code != "FS1:x" || e.Source != "addon" || !e.CapturedAt.Equal(at) {
		t.Fatalf("export = %+v found %v err %v", e, found, err)
	}
}
