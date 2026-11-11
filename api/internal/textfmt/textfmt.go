// Package textfmt is the one place the API words a number, a duration, a day or a clock for a
// reader: Discord embeds and replies, push notifications, the activity feed, link previews and
// the calendar feed. It follows the site's house style (web/src/lib/dates.ts and
// web/src/lib/format), and its tests hold it to the examples every surface shares,
// apps/shared/formats.json (owner rulings 2026-11-10: "Sep", grouped numbers, one spoken
// duration, the 12-hour clock where no reader's locale is known).
//
// Where the reader's own clock matters and Discord renders the message, write Discord's own
// timestamp markup (<t:unix:t>) instead: each reader's client shows it in their own locale.
package textfmt

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// The layouts, in Go's reference time. Go's month and weekday abbreviations are already the
// house three letters ("Sep", "Thu").
const (
	dayLayout          = "Mon 2 Jan"
	longDayLayout      = "Monday 2 January"
	dayMonthYearLayout = "2 Jan 2006"
	hourLayout         = "3"
	minuteLayout       = "3:04"
	secondsLayout      = "3:04:05"
	zoneLayout         = "MST"
)

// figureWholeFrom is where a sim figure drops its decimal: 99.5 stays "99.5", 1234.4 is "1,234".
const figureWholeFrom = 100

// Grouped is v rounded to a whole number, a comma every three digits: "1,234", "-12,500".
func Grouped(v float64) string {
	n := int64(math.Round(v))
	digits := strconv.FormatInt(n, 10)
	sign := ""
	if n < 0 {
		sign, digits = "-", digits[1:]
	}
	var b strings.Builder
	for i, r := range digits {
		if i > 0 && (len(digits)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(r)
	}
	return sign + b.String()
}

// Figure is a sim figure (DPS, threat, points) as the feed and the embeds write it: grouped
// whole numbers from 100, one decimal below ("1,234", "87.5"). scale picks the precision, so a
// lead over a rival is written at the precision of the figure it belongs to; pass v for both
// when there is no other figure.
func Figure(v, scale float64) string {
	if math.Abs(scale) >= figureWholeFrom {
		return Grouped(v)
	}
	return strconv.FormatFloat(v, 'f', 1, 64)
}

// spokenSecondsUnder is where the spoken duration drops its seconds: a 41-minute night is "41 min".
const spokenSecondsUnder = 10 * time.Minute

// Duration is the one spoken duration: "30 s", "1 min 30 s", "41 min", "2 h", "2 h 41 min".
// Seconds show under ten minutes; from there it rounds to the minute. Negative reads as zero.
func Duration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	if d < spokenSecondsUnder {
		total := int(math.Round(d.Seconds()))
		if total < 60 {
			return fmt.Sprintf("%d s", total)
		}
		return joinUnits(fmt.Sprintf("%d min", total/60), total%60, "s")
	}
	total := int(math.Round(d.Minutes()))
	if total < 60 {
		return fmt.Sprintf("%d min", total)
	}
	return joinUnits(fmt.Sprintf("%d h", total/60), total%60, "min")
}

// FightClock is a length on the game's clock, a kill or a pull: "1:42", "12:05", "1:02:09";
// under a minute to the tenth ("0:04.2") unless it is a whole second ("0:45").
func FightClock(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	tenths := int(math.Round(float64(d) / float64(100*time.Millisecond)))
	if tenths < 600 && tenths%10 != 0 {
		return fmt.Sprintf("0:%02d.%d", tenths/10, tenths%10)
	}
	total := int(math.Round(d.Seconds()))
	if total >= 3600 {
		return fmt.Sprintf("%d:%02d:%02d", total/3600, total%3600/60, total%60)
	}
	return fmt.Sprintf("%d:%02d", total/60, total%60)
}

func joinUnits(lead string, n int, unit string) string {
	if n == 0 {
		return lead
	}
	return fmt.Sprintf("%s %d %s", lead, n, unit)
}

// Location is the IANA zone name, UTC when the name is empty or unknown: a guild that set no
// raid zone reads its raids in UTC.
func Location(name string) *time.Location {
	if name == "" {
		return time.UTC
	}
	loc, err := time.LoadLocation(name)
	if err != nil {
		return time.UTC
	}
	return loc
}

// RaidDay is t's day in loc: "Thu 10 Dec".
func RaidDay(t time.Time, loc *time.Location) string { return t.In(loc).Format(dayLayout) }

// WeekdayDayMonth is RaidDay with the year after it when t falls in another year than now (both
// read in loc): "Wed 9 Dec", "Wed 6 Jan 2027".
func WeekdayDayMonth(t time.Time, loc *time.Location, now time.Time) string {
	day := RaidDay(t, loc)
	if year := t.In(loc).Year(); year != now.In(loc).Year() {
		return day + " " + strconv.Itoa(year)
	}
	return day
}

// LongDay is t's day in loc in full: "Thursday 10 December".
func LongDay(t time.Time, loc *time.Location) string { return t.In(loc).Format(longDayLayout) }

// DayMonthYear is t's date in loc with its year: "9 Dec 2026".
func DayMonthYear(t time.Time, loc *time.Location) string {
	return t.In(loc).Format(dayMonthYearLayout)
}

// Clock is t's time in loc on the 12-hour clock, lower case, the minutes only off the hour:
// "8 pm", "8:30 pm". It is the clock for a reader whose locale is not known; a Discord message
// writes <t:unix:t> instead, which each reader's client renders in their own.
func Clock(t time.Time, loc *time.Location) string {
	at := t.In(loc)
	layout := minuteLayout
	if at.Minute() == 0 {
		layout = hourLayout
	}
	half := "am"
	if at.Hour() >= 12 {
		half = "pm"
	}
	return at.Format(layout) + " " + half
}

// ClockSeconds is a log's event time on the 12-hour clock, the second part of the record:
// "11:49:25 pm", "8:00:00 pm".
func ClockSeconds(t time.Time, loc *time.Location) string {
	at := t.In(loc)
	half := "am"
	if at.Hour() >= 12 {
		half = "pm"
	}
	return at.Format(secondsLayout) + " " + half
}

// RaidClock is Clock with the zone's abbreviation after it, so a fixed guild time names its
// zone: "8 pm CST", "8:30 pm UTC".
func RaidClock(t time.Time, loc *time.Location) string {
	return Clock(t, loc) + " " + t.In(loc).Format(zoneLayout)
}

// RaidWhen is the day and the clock with its zone: "Thu 10 Dec, 8 pm CST".
func RaidWhen(t time.Time, loc *time.Location) string {
	return RaidDay(t, loc) + ", " + RaidClock(t, loc)
}
