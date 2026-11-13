package engagement

import (
	"testing"
	"time"
)

func date(s string) time.Time {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestWeekOfIsTheMondayInUTC(t *testing.T) {
	for in, want := range map[string]string{
		"2026-11-13": "2026-11-09", // Friday
		"2026-11-09": "2026-11-09", // Monday
		"2026-11-15": "2026-11-09", // Sunday
	} {
		if got := weekOf(date(in).Add(23 * time.Hour)).Format(time.DateOnly); got != want {
			t.Errorf("weekOf(%s) = %s, want %s", in, got, want)
		}
	}
}

func TestPeriodsFillMissingDaysAndMarkThoseBeforeTheLiveRecord(t *testing.T) {
	live := date("2026-11-12")
	rows := []periodRow{{start: date("2026-11-13"), active: 5, fresh: 2, site: 4, addon: 2, app: 1}}
	got := periods(rows, date("2026-11-13"), 3, stepDays, &live)
	if len(got) != 3 || got[0].Start != "2026-11-13" || got[2].Start != "2026-11-11" {
		t.Fatalf("periods = %+v", got)
	}
	if got[0] != (Period{Start: "2026-11-13", Active: 5, New: 2, Returning: 3, Site: 4, Addon: 2, App: 1}) {
		t.Fatalf("today = %+v", got[0])
	}
	if got[1].Active != 0 || got[1].Partial || !got[2].Partial {
		t.Fatalf("a quiet live day and a backfilled one: %+v %+v", got[1], got[2])
	}
	if all := periods(nil, date("2026-11-13"), 1, stepDays, nil); !all[0].Partial {
		t.Fatal("with no live record yet every day is partial")
	}
}

func TestCohortsCarryOneCellPerElapsedWeek(t *testing.T) {
	live := date("2026-11-04") // a Wednesday: its week is partial, the next is whole
	cells := []cohortCell{
		{week: date("2026-10-26"), offset: 0, count: 10},
		{week: date("2026-10-26"), offset: 1, count: 6},
		{week: date("2026-10-26"), offset: 2, count: 4},
		{week: date("2026-11-09"), offset: 0, count: 3},
	}
	got := cohorts(cells, date("2026-10-26"), date("2026-11-09"), &live)
	if len(got) != 3 {
		t.Fatalf("cohorts = %+v", got)
	}
	first := got[0]
	if first.Week != "2026-10-26" || first.Joined != 10 || len(first.Weeks) != 2 {
		t.Fatalf("first cohort = %+v", first)
	}
	if first.Weeks[0] != (CohortWeek{Active: 6, Partial: true}) || first.Weeks[1] != (CohortWeek{Active: 4}) {
		t.Fatalf("first cohort's weeks = %+v", first.Weeks)
	}
	if got[1].Joined != 0 || len(got[1].Weeks) != 1 || got[2].Joined != 3 || len(got[2].Weeks) != 0 {
		t.Fatalf("later cohorts = %+v", got[1:])
	}
}
