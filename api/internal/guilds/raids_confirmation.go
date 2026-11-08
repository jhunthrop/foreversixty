// api/internal/guilds/raids_confirmation.go
//
// The sign-up PUT's confirmation: the line the apps post in place of a raid notification after
// a player answered from it without opening the app (spec 2026-11-08-app-flows F1). It is
// composed here, after the write, so it can hold what only the answer knows: a Late's minutes
// and the arrival time in the guild's raid time zone.
package guilds

import (
	"fmt"
	"time"
)

// SignupConfirmation is the notification the app posts after a background answer.
type SignupConfirmation struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// confirmationTimeLayout is a time as the raid notifications word it ("8:30 pm CST").
const confirmationTimeLayout = "3:04 pm MST"

// confirmationCopy is the words of the confirmation.
var confirmationCopy = struct {
	Title, LateTitle, Pull, Arriving, NameOnly, Moved string
	States                                            map[string]string
}{
	Title:     "You're %s for %s",
	LateTitle: "You're Late (%d min) for %s",
	Pull:      "%s, pull at %s.",
	Arriving:  "%s, arriving %s.",
	NameOnly:  "%s.",
	Moved:     " %s moved to Bench.",
	States: map[string]string{
		SignupIn: "In", SignupLate: "Late", SignupBench: "on the Bench", SignupOut: "Out",
	},
}

// signupConfirmation words the answer state (with a Late's minutes) for name on raid; moved
// is the character rule R3 moved to Bench, nil when none.
func signupConfirmation(raid RaidView, name, state string, lateMinutes *int16, moved *string) SignupConfirmation {
	tz := ""
	if raid.GuildTimezone != nil {
		tz = *raid.GuildTimezone
	}
	pull := raid.StartsAt.In(raidLocation(tz))
	c := SignupConfirmation{Title: fmt.Sprintf(confirmationCopy.Title, confirmationCopy.States[state], raid.ZoneName)}
	switch {
	case state == SignupLate && lateMinutes != nil:
		c.Title = fmt.Sprintf(confirmationCopy.LateTitle, *lateMinutes, raid.ZoneName)
		arrival := pull.Add(time.Duration(*lateMinutes) * time.Minute)
		c.Body = fmt.Sprintf(confirmationCopy.Arriving, name, arrival.Format(confirmationTimeLayout))
	case isAttending(state):
		c.Body = fmt.Sprintf(confirmationCopy.Pull, name, pull.Format(confirmationTimeLayout))
	default:
		c.Body = fmt.Sprintf(confirmationCopy.NameOnly, name)
	}
	if moved != nil {
		c.Body += fmt.Sprintf(confirmationCopy.Moved, *moved)
	}
	return c
}

// raidLocation is the guild's raid time zone, UTC when it set none or names an unknown one.
func raidLocation(tz string) *time.Location {
	if tz == "" {
		return time.UTC
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.UTC
	}
	return loc
}
