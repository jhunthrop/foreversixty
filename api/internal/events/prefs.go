package events

import (
	"errors"
	"maps"
	"time"
)

// minutesPerDay bounds a quiet-hours minute of day.
const minutesPerDay = 24 * 60

// QuietHours is a daily window, in minutes after midnight in the player's zone. Start after
// End wraps past midnight (22:00 to 07:00).
type QuietHours struct {
	Enabled bool
	Start   int
	End     int
}

// Validate fails a minute outside the day, or an enabled window with no length.
func (q QuietHours) Validate() error {
	for _, m := range []int{q.Start, q.End} {
		if m < 0 || m >= minutesPerDay {
			return errors.New("events: quiet hours are minutes of a day, 0 to 1439")
		}
	}
	if q.Enabled && q.Start == q.End {
		return errors.New("events: quiet hours need a start and an end that differ")
	}
	return nil
}

// contains reports whether minute m falls inside the window.
func (q QuietHours) contains(m int) bool {
	if q.Start < q.End {
		return m >= q.Start && m < q.End
	}
	return m >= q.Start || m < q.End
}

// The settings of an account that never saved any; migration 0060 has the same defaults.
const (
	DefaultTimeZone   = "UTC"
	DefaultQuietStart = 22 * 60
	DefaultQuietEnd   = 7 * 60
)

// DefaultPrefs is the catalogue's defaults, in UTC, with quiet hours off.
func DefaultPrefs() Prefs {
	return Prefs{TimeZone: DefaultTimeZone, Quiet: QuietHours{Start: DefaultQuietStart, End: DefaultQuietEnd}}
}

// PrefKey is one event on one channel.
type PrefKey struct {
	Kind    Kind
	Channel Channel
}

// Prefs is a player's notification settings: the overrides of the catalogue's defaults, the
// quiet hours and the zone they are read in ("" is UTC).
type Prefs struct {
	TimeZone  string
	Quiet     QuietHours
	Overrides map[PrefKey]bool
}

// Allows reports whether kind goes out on channel: the player's override, else the
// catalogue's default. An unknown kind or channel is never allowed.
func (p Prefs) Allows(kind Kind, channel Channel) bool {
	spec, ok := Lookup(kind)
	if !ok || !KnownChannel(channel) {
		return false
	}
	if on, set := p.Overrides[PrefKey{kind, channel}]; set {
		return on
	}
	if spec.OffWhenOn != "" && p.Allows(spec.OffWhenOn, channel) {
		return false
	}
	return spec.DefaultOn
}

// With returns a copy of p with kind on channel set to on.
func (p Prefs) With(kind Kind, channel Channel, on bool) Prefs {
	out := p
	out.Overrides = maps.Clone(p.Overrides)
	if out.Overrides == nil {
		out.Overrides = map[PrefKey]bool{}
	}
	out.Overrides[PrefKey{kind, channel}] = on
	return out
}

// InQuiet reports whether at falls in the player's quiet hours, read in their zone.
func (p Prefs) InQuiet(at time.Time) bool {
	if !p.Quiet.Enabled {
		return false
	}
	local := at.In(p.location())
	return p.Quiet.contains(local.Hour()*60 + local.Minute())
}

// location is the player's zone; an unreadable one falls back to UTC.
func (p Prefs) location() *time.Location {
	if p.TimeZone == "" {
		return time.UTC
	}
	loc, err := time.LoadLocation(p.TimeZone)
	if err != nil {
		return time.UTC
	}
	return loc
}

// ValidTimeZone reports whether name is an IANA zone the server knows.
func ValidTimeZone(name string) bool {
	if name == "" || name == "Local" {
		return false
	}
	_, err := time.LoadLocation(name)
	return err == nil
}
