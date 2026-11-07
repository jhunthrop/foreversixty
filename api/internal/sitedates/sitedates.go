// Package sitedates holds the site calendar's dates the API reasons about. The site's
// web/src/data/dates.json is the source; a test holds these to it.
package sitedates

import "time"

// FirstRaids is the calendar day the first raids open (dates.json, role raids).
const FirstRaids = "2026-12-09"

var firstRaidsDay = mustDay(FirstRaids)

func mustDay(day string) time.Time {
	t, err := time.Parse(time.DateOnly, day)
	if err != nil {
		panic("sitedates: " + day + " is not a date")
	}
	return t
}

// RaidsOpen reports whether the first raids have opened at now: from the start of their day, UTC.
func RaidsOpen(now time.Time) bool {
	return !now.UTC().Before(firstRaidsDay)
}
