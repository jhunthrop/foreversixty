// api/internal/dataaddon/db_test.go
package dataaddon

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// testURL is this package's own copy of the shape every DB-backed package in
// this repo declares (compare api/internal/rating/db_test.go):
// TEST_DATABASE_URL must point at dataaddon_test, not the shared
// forever_test other lanes' tests truncate concurrently, per this lane's
// dispatch ("own throwaway test database").
func testURL(t *testing.T) string {
	t.Helper()
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL not set; see docs/superpowers/plans/2026-09-21-data-addon.md Task 1 Step 1")
	}
	return u
}

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := testURL(t)
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}
