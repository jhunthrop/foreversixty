package engagement

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
	if _, err := pool.Exec(context.Background(), `truncate users cascade`); err != nil {
		t.Fatal(err)
	}
	return &Store{Pool: pool}
}

func exec(t *testing.T, s *Store, sql string, args ...any) {
	t.Helper()
	if _, err := s.Pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
}

func TestMarkActiveAddsSurfacesAndMakesABackfilledDayLive(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	exec(t, s, `insert into users (id, battletag) values (5, 'a#1')`)
	exec(t, s, `insert into user_active_days (user_id, day, surfaces, backfilled) values (5, '2026-11-12', 2, true)`)
	for _, m := range []struct {
		day     string
		surface Surface
	}{{"2026-11-13", SurfaceSite}, {"2026-11-13", SurfaceApp}, {"2026-11-13", SurfaceSite}, {"2026-11-12", SurfaceAddon}} {
		if err := s.MarkActive(ctx, 5, date(m.day), m.surface); err != nil {
			t.Fatal(err)
		}
	}
	rows, err := s.Pool.Query(ctx, `select day::text, surfaces, backfilled from user_active_days where user_id = 5 order by day`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	type row struct {
		day        string
		surfaces   int16
		backfilled bool
	}
	var got []row
	for rows.Next() {
		var r row
		if err := rows.Scan(&r.day, &r.surfaces, &r.backfilled); err != nil {
			t.Fatal(err)
		}
		got = append(got, r)
	}
	want := []row{{"2026-11-12", 2, false}, {"2026-11-13", 5, false}}
	if len(got) != 2 || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("rows = %+v, want %+v", got, want)
	}
}

func TestReportCountsActivesNewReturningAndRetentionWithoutInternalAccounts(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	exec(t, s, `insert into users (id, battletag, role, created_at) values
		(78, 'owner#1', 'admin', '2026-10-01T12:00:00Z'),
		(99, 'comped#1', 'user', '2026-10-01T12:00:00Z'),
		(1, 'new#1', 'user', '2026-11-13T08:00:00Z'),
		(2, 'back#1', 'user', '2026-11-02T08:00:00Z'),
		(3, 'gone#1', 'user', '2026-11-02T09:00:00Z')`)
	exec(t, s, `insert into entitlements (user_id, plan, source, status) values (99, 'premium', 'grant', 'active')`)
	exec(t, s, `insert into user_active_days (user_id, day, surfaces, backfilled) values
		(2, '2026-11-02', 1, true), (3, '2026-11-02', 1, true), (2, '2026-11-03', 1, true),
		(2, '2026-11-10', 2, false), (78, '2026-11-13', 1, false), (99, '2026-11-13', 2, false),
		(1, '2026-11-13', 1, false), (2, '2026-11-13', 5, false)`)

	rep, err := s.Report(ctx, time.Date(2026, 11, 13, 15, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if rep.Today != "2026-11-13" || rep.LiveSince == nil || *rep.LiveSince != "2026-11-10" || rep.InternalAccounts != 2 {
		t.Fatalf("header = %s %v %d", rep.Today, rep.LiveSince, rep.InternalAccounts)
	}
	if len(rep.Days) != DaysShown || len(rep.Weeks) != WeeksShown || len(rep.Cohorts) != CohortsShown {
		t.Fatalf("lengths %d %d %d", len(rep.Days), len(rep.Weeks), len(rep.Cohorts))
	}
	if got, want := rep.Days[0], (Period{Start: "2026-11-13", Active: 2, New: 1, Returning: 1, Site: 2, App: 1}); got != want {
		t.Fatalf("today = %+v, want %+v", got, want)
	}
	if got := rep.Days[3]; got.Start != "2026-11-10" || got.Active != 1 || got.Addon != 1 || got.Partial {
		t.Fatalf("the first live day = %+v", got)
	}
	if got := rep.Days[10]; got.Start != "2026-11-03" || got.Active != 1 || !got.Partial {
		t.Fatalf("a backfilled day = %+v", got)
	}
	if got, want := rep.Weeks[0], (Period{Start: "2026-11-09", Active: 2, New: 1, Returning: 1, Site: 2, Addon: 1, App: 1, Partial: true}); got != want {
		t.Fatalf("this week = %+v, want %+v", got, want)
	}
	if got := rep.Weeks[1]; got.Start != "2026-11-02" || got.Active != 2 || got.New != 2 || got.Returning != 0 {
		t.Fatalf("last week = %+v", got)
	}
	last, prev := rep.Cohorts[CohortsShown-1], rep.Cohorts[CohortsShown-2]
	if last.Week != "2026-11-09" || last.Joined != 1 || len(last.Weeks) != 0 {
		t.Fatalf("this week's cohort = %+v", last)
	}
	if prev.Week != "2026-11-02" || prev.Joined != 2 || len(prev.Weeks) != 1 || prev.Weeks[0].Active != 1 {
		t.Fatalf("last week's cohort = %+v", prev)
	}
}
