// api/internal/guilds/raids_validate.go
//
// Request validation for the raid routes. The store validates nothing (the migration's checks
// only backstop), so every bound lives here, answered as 422 invalid with field names.
package guilds

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// Raid form bounds (spec §3.5) and the sign-up note bound (migration 0050).
const (
	raidTitleMax     = 80
	raidNotesMax     = 2000
	raidDurationMin  = 60
	raidDurationMax  = 240
	raidDurationStep = 30
	raidRepeatWeeks  = 3 // R7: "Also plan the next 3 weeks"
	signupNoteMax    = 140
	// maxRaidAhead is how far ahead a raid may start (and its last repeated week may fall).
	maxRaidAhead       = 366 * 24 * time.Hour
	raidStartsAtInPast = "Pick a time after now."
	raidStartsAtTooFar = "Pick a date within the next year."
	controlCharMessage = "Remove the control characters."
)

// raidRequest is the body of POST and PATCH /v1/guilds/{id}/raids[/{raid}]. PATCH ignores
// RepeatWeeks.
type raidRequest struct {
	Zone            string `json:"zone"`
	Title           string `json:"title"`
	StartsAt        string `json:"starts_at"`
	DurationMinutes int    `json:"duration_minutes"`
	Notes           string `json:"notes"`
	Timezone        string `json:"timezone"`
	RepeatWeeks     int    `json:"repeat_weeks"`
	// PostToDiscord is the form's box; nil when the save did not carry it.
	PostToDiscord *bool `json:"post_to_discord"`
	// RemindDayBefore is the form's reminder box; nil reads as the guild's remind default.
	RemindDayBefore *bool `json:"remind_day_before"`
}

// raidPlan is a validated raidRequest: the first raid's input and the guild's location.
type raidPlan struct {
	input    RaidInput
	location *time.Location
	timezone string
	repeat   int
}

// validateRaid checks req against the form's bounds. keepStart is the stored start on an
// edit, which may stay in the past; nil on create. The map is nil when req is valid.
func validateRaid(req raidRequest, now time.Time, keepStart *time.Time, allowRepeat bool) (raidPlan, map[string]string) {
	fields := map[string]string{}
	plan := raidPlan{repeat: req.RepeatWeeks}
	zone, ok := ZoneByKey(req.Zone)
	if !ok {
		fields["zone"] = "Pick a raid."
	}
	title, titleClean := cleanText(strings.TrimSpace(req.Title))
	if !titleClean {
		fields["title"] = controlCharMessage
	}
	if title == "" {
		title = zone.Name
	}
	if utf8.RuneCountInString(title) > raidTitleMax {
		fields["title"] = "Keep the title to 80 characters."
	}
	d := req.DurationMinutes
	if d < raidDurationMin || d > raidDurationMax || d%raidDurationStep != 0 {
		fields["duration_minutes"] = "Pick a length between 1 and 4 hours."
	}
	if utf8.RuneCountInString(req.Notes) > raidNotesMax {
		fields["notes"] = "Keep the notes to 2,000 characters."
	} else if _, notesClean := cleanNotes(req.Notes); !notesClean {
		fields["notes"] = controlCharMessage
	}
	loc, err := loadTimezone(req.Timezone)
	if err != nil {
		fields["timezone"] = "Pick a timezone."
	}
	startsAt, msg := validateStart(req.StartsAt, now, keepStart)
	if msg == "" && !repeatWithinReach(startsAt, req.RepeatWeeks, now, keepStart) {
		msg = raidStartsAtTooFar
	}
	if msg != "" {
		fields["starts_at"] = msg
	}
	if req.RepeatWeeks != 0 && (!allowRepeat || req.RepeatWeeks != raidRepeatWeeks) {
		fields["repeat_weeks"] = "Repeat for 3 weeks or not at all."
	}
	if len(fields) > 0 {
		return raidPlan{}, fields
	}
	plan.location, plan.timezone = loc, req.Timezone
	plan.input = RaidInput{
		Zone: zone.Key, Title: title, StartsAt: startsAt, DurationMinutes: d,
		Notes: req.Notes, PostDefault: ticked(req.PostToDiscord),
	}
	return plan, nil
}

// loadTimezone loads an IANA zone name; an empty name is refused, never read as UTC.
func loadTimezone(name string) (*time.Location, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errEmptyTimezone
	}
	return time.LoadLocation(name)
}

var errEmptyTimezone = errors.New("guilds: empty timezone")

// validateStart parses an RFC 3339 start that must be after now, unless it equals keepStart.
func validateStart(raw string, now time.Time, keepStart *time.Time) (time.Time, string) {
	t, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return time.Time{}, "Pick a date and time."
	}
	if keepStart != nil && t.Equal(*keepStart) {
		return t, ""
	}
	if !t.After(now) {
		return time.Time{}, raidStartsAtInPast
	}
	if t.After(now.Add(maxRaidAhead)) {
		return time.Time{}, raidStartsAtTooFar
	}
	return t, ""
}

// repeatWithinReach reports whether the last of repeat weekly raids after start is still
// within maxRaidAhead; an unmoved edit is left alone.
func repeatWithinReach(start time.Time, repeat int, now time.Time, keepStart *time.Time) bool {
	if keepStart != nil && start.Equal(*keepStart) {
		return true
	}
	return !start.AddDate(0, 0, 7*repeat).After(now.Add(maxRaidAhead))
}

// cleanText returns s and whether it holds no control character at all; titles and
// sign-up notes are one line, and Postgres refuses NUL outright.
func cleanText(s string) (string, bool) { return s, !hasControl(s, false) }

// cleanNotes is cleanText for the multi-line notes, which may hold newlines and tabs.
func cleanNotes(s string) (string, bool) { return s, !hasControl(s, true) }

func hasControl(s string, allowLayout bool) bool {
	for _, r := range s {
		if r < 0x20 && !(allowLayout && (r == '\n' || r == '\t')) {
			return true
		}
	}
	return false
}

// weeklyInputs expands plan into the first raid and plan.repeat weekly ones at the same
// guild wall time: each week's instant is built from the first raid's wall clock in the
// guild's zone, so a daylight-saving change keeps 20:00 guild time (R4).
func weeklyInputs(plan raidPlan) []RaidInput {
	first := plan.input.StartsAt.In(plan.location)
	y, m, d := first.Date()
	out := make([]RaidInput, 0, plan.repeat+1)
	for k := 0; k <= plan.repeat; k++ {
		in := plan.input
		in.StartsAt = time.Date(y, m, d+7*k, first.Hour(), first.Minute(), 0, 0, plan.location)
		out = append(out, in)
	}
	return out
}

// signupRequest is the body of PUT …/raids/{raid}/signups/{character_key}.
type signupRequest struct {
	State string `json:"state"`
	Role  string `json:"role"`
	Note  string `json:"note"`
	// Spec, when set, is a specs.json key saved as the character's raid spec before the sign-up,
	// whose role the sign-up then takes (spec 2026-10-28 §P.10.7).
	Spec string `json:"spec"`
}

// validateSignup checks a sign-up body; role may be empty (the spec's default applies).
func validateSignup(req signupRequest) (signupRequest, map[string]string) {
	fields := map[string]string{}
	if _, ok := raidspec.ByKey(req.Spec); req.Spec != "" && !ok {
		fields["spec"] = "Choose a spec."
	}
	switch req.State {
	case SignupIn, SignupLate, SignupBench, SignupOut:
	default:
		fields["state"] = "Pick In, Late, Bench or Out."
	}
	switch req.Role {
	case "", RoleTank, RoleHealer, RoleDPS:
	default:
		fields["role"] = "Pick tank, healer or DPS."
	}
	req.Note = strings.TrimSpace(req.Note)
	if utf8.RuneCountInString(req.Note) > signupNoteMax {
		fields["note"] = "Keep the note to 140 characters."
	} else if _, clean := cleanText(req.Note); !clean {
		fields["note"] = controlCharMessage
	}
	if len(fields) > 0 {
		return signupRequest{}, fields
	}
	return req, nil
}
