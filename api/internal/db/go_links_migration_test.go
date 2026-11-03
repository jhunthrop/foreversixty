package db

import (
	"context"
	"testing"
)

const goLinksVersion = 66

// 0066 adds the tracked links with their starter set, the daily counters and the account's
// first-touch slug, makes the owner the admin, and its down removes all three.
func TestMigrationAddsGoLinks(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	migrateTo(t, url, goLinksVersion-1)
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users cascade`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into users (id, battletag) values (78, 'hunthrop#1894'), (79, 'someone#1')`); err != nil {
		t.Fatal(err)
	}
	migrateTo(t, url, goLinksVersion)

	var roles []string
	rows, err := pool.Query(ctx, `select role from users order by id`)
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			t.Fatal(err)
		}
		roles = append(roles, role)
	}
	rows.Close()
	if len(roles) != 2 || roles[0] != "admin" || roles[1] != "user" {
		t.Fatalf("roles after 0066 = %v, want [admin user]", roles)
	}

	var starters int
	if err := pool.QueryRow(ctx, `select count(*) from go_links where slug in
		('reddit','reddit-classicwow','reddit-wowforever','reddit-wowclassic','reddit-guides','reddit-sim',
		 'discord','youtube','curseforge','wago','twitter')`).Scan(&starters); err != nil {
		t.Fatal(err)
	}
	if starters != 11 {
		t.Fatalf("starter links = %d, want 11", starters)
	}
	var guidesTarget string
	if err := pool.QueryRow(ctx, `select target from go_links where slug = 'reddit-guides'`).Scan(&guidesTarget); err != nil || guidesTarget != "/guides" {
		t.Fatalf("reddit-guides target = %q (%v), want /guides", guidesTarget, err)
	}

	for _, target := range []string{"//evil.example", "https://evil.example", `/\evil.example`, "evil", "/a b", ""} {
		if _, err := pool.Exec(ctx, `insert into go_links (slug, target, label) values ('bad', $1, 'Bad')`, target); err == nil {
			t.Errorf("target %q was accepted", target)
		}
	}
	for _, slug := range []string{"Reddit", "-x", "a_b", ""} {
		if _, err := pool.Exec(ctx, `insert into go_links (slug, target, label) values ($1, '/', 'Bad')`, slug); err == nil {
			t.Errorf("slug %q was accepted", slug)
		}
	}
	if _, err := pool.Exec(ctx, `insert into go_link_clicks (slug, day, clicks) values ('reddit', current_date, 1)`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `update users set first_touch_slug = 'reddit' where id = 79`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `update users set first_touch_slug = 'nope' where id = 79`); err == nil {
		t.Fatal("a first-touch slug with no link was accepted")
	}

	migrateTo(t, url, goLinksVersion-1)
	var tables int
	if err := pool.QueryRow(ctx, `select count(*) from information_schema.tables
		where table_name in ('go_links', 'go_link_clicks')`).Scan(&tables); err != nil {
		t.Fatal(err)
	}
	var column int
	if err := pool.QueryRow(ctx, `select count(*) from information_schema.columns
		where table_name = 'users' and column_name = 'first_touch_slug'`).Scan(&column); err != nil {
		t.Fatal(err)
	}
	if tables != 0 || column != 0 {
		t.Fatalf("after the down: %d tables, %d column, want none", tables, column)
	}
	if _, err := pool.Exec(ctx, `truncate users cascade`); err != nil {
		t.Fatal(err)
	}
}
