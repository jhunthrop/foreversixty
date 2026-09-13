package db

import (
	"context"
	"os"
	"testing"
)

func testURL(t *testing.T) string {
	t.Helper()
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	return u
}

func TestMigrateCreatesSubscribers(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	if err := Migrate(url); err != nil { // idempotent
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	var n int
	if err := pool.QueryRow(context.Background(), `select count(*) from information_schema.tables where table_name = 'subscribers'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("subscribers table count = %d", n)
	}
}
