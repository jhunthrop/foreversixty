package activity

import "time"

// The weekly reset the snapshots are cut at: Tuesday 15:00 UTC.
const (
	resetWeekday = time.Tuesday
	resetHour    = 15
)

// WeekOf is the reset week t falls in, named by the date of its reset.
func WeekOf(t time.Time) time.Time {
	t = t.UTC()
	reset := time.Date(t.Year(), t.Month(), t.Day(), resetHour, 0, 0, 0, time.UTC)
	back := (int(t.Weekday()) - int(resetWeekday) + 7) % 7
	reset = reset.AddDate(0, 0, -back)
	if reset.After(t) {
		reset = reset.AddDate(0, 0, -7)
	}
	return time.Date(reset.Year(), reset.Month(), reset.Day(), 0, 0, 0, 0, time.UTC)
}

// PreviousWeek is the reset week before the one t falls in.
func PreviousWeek(t time.Time) time.Time { return WeekOf(t).AddDate(0, 0, -7) }
