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

func TestMigrateCreatesBuildsWithTheContractColumns(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(),
		`select column_name, data_type, is_nullable from information_schema.columns where table_name = 'builds'`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	got := map[string][2]string{}
	for rows.Next() {
		var name, dataType, nullable string
		if err := rows.Scan(&name, &dataType, &nullable); err != nil {
			t.Fatal(err)
		}
		got[name] = [2]string{dataType, nullable}
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	want := map[string][2]string{
		"id":           {"text", "NO"},
		"class_id":     {"smallint", "NO"},
		"race_id":      {"smallint", "NO"},
		"tree_version": {"text", "NO"},
		"point_order":  {"ARRAY", "NO"},
		"gear":         {"jsonb", "NO"},
		"title":        {"text", "YES"},
		"created_at":   {"timestamp with time zone", "NO"},
		"views":        {"bigint", "NO"},
	}
	if len(got) != len(want) {
		t.Fatalf("columns = %v, want exactly %d columns", got, len(want))
	}
	for name, w := range want {
		if got[name] != w {
			t.Errorf("column %s = %v, want %v", name, got[name], w)
		}
	}
}
