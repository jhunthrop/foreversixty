// api/internal/guilds/calendar_ics.go
//
// The calendar link's body (W3-C17): RFC 5545 with one VEVENT per raid. The words match the
// app's own calendar sync (spec 2026-10-31-mobile-wave3 §F16.4): SUMMARY "{zone} · {answer}",
// LOCATION "{guild} · {zone}", DESCRIPTION "Pull {time}. {n} of {size} signed up." and the
// raid page's link. No VALARM: the player's calendar app decides the alert.
package guilds

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// CalendarName is the calendar's name, in the link and in the app's own calendar.
const CalendarName = "Forever Sixty raids"

// calendarRefresh asks a subscribing calendar to fetch again hourly; the server's cache is five
// minutes, and the per-token budget is well above an hourly poll.
const calendarRefresh = "PT1H"

// icsLineOctets is RFC 5545 §3.1's line limit before folding.
const icsLineOctets = 75

// stateWords are the answers as the event title says them.
var stateWords = map[string]string{"in": "In", "late": "Late", "bench": "Bench", "out": "Out"}

// notSignedUp is the title's answer when none of the player's characters answered.
const notSignedUp = "Not signed up"

// CalendarEventSummary is the event title: the zone, then the player's answer.
func CalendarEventSummary(zone, state string) string {
	word, ok := stateWords[state]
	if !ok {
		word = notSignedUp
	}
	return zone + " · " + word
}

// calendarZoneName is the raid's zone as players name it, else its title.
func calendarZoneName(r Raid) string {
	if z, ok := ZoneByKey(r.Zone); ok {
		return z.Name
	}
	return r.Title
}

// calendarDescription is the event's notes: the pull in the guild's time, and the count.
func calendarDescription(c CalendarRaid) string {
	loc, err := time.LoadLocation(c.Timezone)
	if c.Timezone == "" || err != nil {
		loc = time.UTC
	}
	pull := c.StartsAt.In(loc).Format("3:04 pm MST")
	count := fmt.Sprintf("%d signed up.", c.SignedUp)
	if z, ok := ZoneByKey(c.Zone); ok && z.Size != nil {
		count = fmt.Sprintf("%d of %d signed up.", c.SignedUp, *z.Size)
	}
	return "Pull " + pull + ". " + count
}

// RenderCalendar writes the feed for raids. publicBaseURL is the site origin the raid links
// open; stamp is DTSTAMP's fallback for a raid with no update time.
func RenderCalendar(raids []CalendarRaid, publicBaseURL string, stamp time.Time) string {
	var b strings.Builder
	line := func(name, value string) { b.WriteString(foldICSLine(name + ":" + value)) }
	line("BEGIN", "VCALENDAR")
	line("VERSION", "2.0")
	line("PRODID", "-//Forever Sixty//Raids//EN")
	line("CALSCALE", "GREGORIAN")
	line("METHOD", "PUBLISH")
	line("X-WR-CALNAME", escapeICSText(CalendarName))
	line("REFRESH-INTERVAL;VALUE=DURATION", calendarRefresh)
	line("X-PUBLISHED-TTL", calendarRefresh)
	for _, c := range raids {
		zone := calendarZoneName(c.Raid)
		updated := c.UpdatedAt
		if updated.IsZero() {
			updated = stamp
		}
		line("BEGIN", "VEVENT")
		line("UID", fmt.Sprintf("raid-%d@foreversixty.gg", c.ID))
		line("DTSTAMP", icsTime(updated))
		line("LAST-MODIFIED", icsTime(updated))
		line("DTSTART", icsTime(c.StartsAt))
		line("DTEND", icsTime(c.EndsAt()))
		line("SUMMARY", escapeICSText(CalendarEventSummary(zone, c.State)))
		line("LOCATION", escapeICSText(c.Guild.Name+" · "+zone))
		line("DESCRIPTION", escapeICSText(calendarDescription(c)))
		line("URL", publicBaseURL+RaidPagePath(GuildPagePath(c.Guild), c.ID))
		line("STATUS", "CONFIRMED")
		line("TRANSP", "OPAQUE")
		line("END", "VEVENT")
	}
	line("END", "VCALENDAR")
	return b.String()
}

// icsTime is a UTC date-time, RFC 5545 §3.3.5 form 2.
func icsTime(t time.Time) string { return t.UTC().Format("20060102T150405Z") }

// escapeICSText escapes a TEXT value (RFC 5545 §3.3.11): backslash, semicolon, comma and line
// breaks, so a guild or raid name cannot end a property or inject one.
func escapeICSText(s string) string {
	r := strings.NewReplacer(`\`, `\\`, ";", `\;`, ",", `\,`, "\r\n", `\n`, "\n", `\n`, "\r", `\n`)
	return r.Replace(s)
}

// foldICSLine ends a content line with CRLF, folded at 75 octets without splitting a UTF-8
// character (RFC 5545 §3.1): each continuation starts with one space.
func foldICSLine(s string) string {
	var b strings.Builder
	limit := icsLineOctets
	for len(s) > limit {
		cut := limit
		for cut > 0 && !utf8.RuneStart(s[cut]) {
			cut--
		}
		b.WriteString(s[:cut])
		b.WriteString("\r\n ")
		s = s[cut:]
		limit = icsLineOctets - 1
	}
	b.WriteString(s)
	b.WriteString("\r\n")
	return b.String()
}
