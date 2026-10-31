package db

import (
	"context"
	"testing"
)

const fightCountsVersion = 59

// fightCountColumns are the columns 0059 adds to fights, all nullable: null is a fight indexed
// before engine 0.6.0, which the API answers with no counts at all.
var fightCountColumns = []string{"mob_deaths", "group_kills", "group_deaths", "group_players"}

func TestMigrationAddsTheFightCountsNullable(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	nullable := func() map[string]bool {
		t.Helper()
		rows, err := pool.Query(ctx, `select column_name, is_nullable = 'YES' from information_schema.columns
			where table_name = 'fights' and column_name = any($1)`, fightCountColumns)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()
		out := map[string]bool{}
		for rows.Next() {
			var name string
			var null bool
			if err := rows.Scan(&name, &null); err != nil {
				t.Fatal(err)
			}
			out[name] = null
		}
		return out
	}

	migrateTo(t, url, fightCountsVersion-1)
	if got := nullable(); len(got) != 0 {
		t.Fatalf("before 0059 fights already has %v", got)
	}
	migrateTo(t, url, fightCountsVersion)
	got := nullable()
	for _, c := range fightCountColumns {
		if null, ok := got[c]; !ok || !null {
			t.Errorf("after 0059 column %s present=%v nullable=%v, want a nullable column", c, ok, null)
		}
	}
	migrateTo(t, url, fightCountsVersion-1)
	if got := nullable(); len(got) != 0 {
		t.Errorf("0059 down left %v behind", got)
	}
}
