package engagement

import "time"

// How far back the owner's page looks.
const (
	DaysShown    = 28 // daily rows, today included
	WeeksShown   = 12 // weekly rows, this week included
	CohortsShown = 8  // signup weeks in the retention table, this week included
)

const day = 24 * time.Hour

// Report is the owner's /admin/activity: every figure excludes the internal accounts.
type Report struct {
	Today string `json:"today"`
	// LiveSince is the first UTC day the middleware recorded; days before it hold only what the
	// backfill could reconstruct, a floor. Null until the first live write.
	LiveSince        *string  `json:"live_since"`
	InternalAccounts int      `json:"internal_accounts"`
	Days             []Period `json:"days"`
	Weeks            []Period `json:"weeks"`
	Cohorts          []Cohort `json:"cohorts"`
}

// Period is one day or one week (Monday to Sunday, UTC), newest first in the report.
type Period struct {
	Start string `json:"start"`
	// Active is every account seen; New the ones that signed up in the period; Returning the rest.
	Active    int `json:"active"`
	New       int `json:"new"`
	Returning int `json:"returning"`
	// Site, Addon and App count the accounts seen on each surface; one account can be on several.
	Site    int  `json:"site"`
	Addon   int  `json:"addon"`
	App     int  `json:"app"`
	Partial bool `json:"partial"`
}

// Cohort is the accounts that signed up in one week and how many came back in each week after.
type Cohort struct {
	Week   string `json:"week"`
	Joined int    `json:"joined"`
	// Weeks[k] is week N+k+1, up to and including this week.
	Weeks []CohortWeek `json:"weeks"`
}

// CohortWeek is one later week of a cohort: how many of its accounts were active in it.
type CohortWeek struct {
	Active  int  `json:"active"`
	Partial bool `json:"partial"`
}

// periodRow is one aggregated row as the store reads it.
type periodRow struct {
	start                           time.Time
	active, fresh, site, addon, app int
}

// cohortCell is one (signup week, weeks later) count as the store reads it; offset 0 is the
// cohort's size.
type cohortCell struct {
	week   time.Time
	offset int
	count  int
}

// dayOf is t's UTC day at midnight.
func dayOf(t time.Time) time.Time { return t.UTC().Truncate(day) }

// weekOf is the Monday (UTC) starting t's week.
func weekOf(t time.Time) time.Time {
	d := dayOf(t)
	return d.AddDate(0, 0, -((int(d.Weekday()) + 6) % 7))
}

// partialBefore reports whether a period starting at start holds days before liveSince.
func partialBefore(start time.Time, liveSince *time.Time) bool {
	return liveSince == nil || start.Before(*liveSince)
}

// periods lays rows out newest first, one per step back from newest, n of them, a missing
// period as zeros. A period is partial when it begins before liveSince.
func periods(rows []periodRow, newest time.Time, n int, step func(time.Time, int) time.Time, liveSince *time.Time) []Period {
	byStart := make(map[time.Time]periodRow, len(rows))
	for _, r := range rows {
		byStart[r.start.UTC()] = r
	}
	out := make([]Period, n)
	for i := range n {
		start := step(newest, -i)
		r := byStart[start]
		out[i] = Period{
			Start: start.Format(time.DateOnly), Active: r.active, New: r.fresh, Returning: r.active - r.fresh,
			Site: r.site, Addon: r.addon, App: r.app, Partial: partialBefore(start, liveSince),
		}
	}
	return out
}

func stepDays(t time.Time, n int) time.Time  { return t.AddDate(0, 0, n) }
func stepWeeks(t time.Time, n int) time.Time { return t.AddDate(0, 0, 7*n) }

// cohorts lays the signup weeks out oldest first, from first through thisWeek. Every cohort
// carries one cell per week elapsed since it signed up; a cell is partial when its week begins
// before liveSince.
func cohorts(cells []cohortCell, first, thisWeek time.Time, liveSince *time.Time) []Cohort {
	counts := map[time.Time]map[int]int{}
	for _, c := range cells {
		w := c.week.UTC()
		if counts[w] == nil {
			counts[w] = map[int]int{}
		}
		counts[w][c.offset] = c.count
	}
	var out []Cohort
	for w := first; !w.After(thisWeek); w = stepWeeks(w, 1) {
		elapsed := int(thisWeek.Sub(w) / (7 * day))
		weeks := make([]CohortWeek, elapsed)
		for k := range weeks {
			weeks[k] = CohortWeek{Active: counts[w][k+1], Partial: partialBefore(stepWeeks(w, k+1), liveSince)}
		}
		out = append(out, Cohort{Week: w.Format(time.DateOnly), Joined: counts[w][0], Weeks: weeks})
	}
	return out
}
