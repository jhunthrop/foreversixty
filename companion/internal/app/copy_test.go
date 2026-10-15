package app_test

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/app"
)

// The words are the spec's (§4.10, §5.2); a change to them is a change to the copy, made here.
func TestTheAddonCopyIsTheSpecsWords(t *testing.T) {
	for got, want := range map[string]string{
		app.CopySeenBySite:      "Seen by the site",
		app.CopySimQueued:       "Sim queued",
		app.CopySimNeedsPremium: "Sim needs premium",
		app.CopySimDuplicate:    "Sim already running",
		app.CopySimInvalid:      "Sim request invalid",
	} {
		if got != want {
			t.Errorf("copy = %q, want %q", got, want)
		}
	}
}

func TestTheAddonLineIsLocalTimeAndSaysWhatBecameOfTheSim(t *testing.T) {
	local := time.FixedZone("EST", -5*3600)
	now := time.Date(2026, 11, 4, 18, 0, 0, 0, local)
	seenToday := time.Date(2026, 11, 4, 19, 3, 0, 0, time.UTC) // 14:03 in EST
	seenEarlier := time.Date(2026, 11, 3, 19, 3, 0, 0, time.UTC)
	tomorrow := now.Add(24 * time.Hour)

	for _, c := range []struct {
		name  string
		seen  time.Time
		state string
		now   time.Time
		want  string
	}{
		{"no ack", time.Time{}, addon.SimQueued, now, ""},
		{"today", seenToday, "", now, app.CopySeenBySite + " 14:03"},
		{"another day", seenToday, "", tomorrow, app.CopySeenBySite + " 4 Nov 14:03"},
		{"yesterday's ack", seenEarlier, "", now, app.CopySeenBySite + " 3 Nov 14:03"},
		{"queued", seenToday, addon.SimQueued, now, app.CopySeenBySite + " 14:03" + app.CopySeparator + app.CopySimQueued},
		{"premium", seenToday, addon.SimNeedsPremium, now, app.CopySeenBySite + " 14:03" + app.CopySeparator + app.CopySimNeedsPremium},
		{"duplicate", seenToday, addon.SimDuplicate, now, app.CopySeenBySite + " 14:03" + app.CopySeparator + app.CopySimDuplicate},
		{"invalid", seenToday, addon.SimInvalid, now, app.CopySeenBySite + " 14:03" + app.CopySeparator + app.CopySimInvalid},
		{"a state with no words", seenToday, "unknown_character", now, app.CopySeenBySite + " 14:03"},
	} {
		if got := app.AddonLine(c.seen, c.state, c.now); got != c.want {
			t.Errorf("%s: AddonLine = %q, want %q", c.name, got, c.want)
		}
	}
}
