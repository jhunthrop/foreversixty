package app_test

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/app"
)

// The words are the designer's (companion review round 1, findings 1 and 7); a change to them is a
// change to the copy, made here.
func TestTheAddonCopyIsTheDesignersWords(t *testing.T) {
	for got, want := range map[string]string{
		app.CopyLastSynced:      "Last synced",
		app.CopyNeverSynced:     "Addon not synced yet. Install ForeverSixty, then log out once.",
		app.CopyStaleHelp:       "Log out of the game with the addon installed to sync.",
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

func TestTheAddonLineSaysHowLongAgoInLocalTimeAndWhatBecameOfTheSim(t *testing.T) {
	local := time.FixedZone("EST", -5*3600)
	now := time.Date(2026, 11, 4, 18, 0, 0, 0, local)
	at := func(d time.Duration) time.Time { return now.Add(-d) }
	seenToday := time.Date(2026, 11, 4, 19, 3, 0, 0, time.UTC) // 14:03 in EST

	for _, c := range []struct {
		name  string
		seen  time.Time
		state string
		want  string
	}{
		{"no ack", time.Time{}, addon.SimQueued, app.CopyNeverSynced},
		{"just now", at(20 * time.Second), "", "Last synced less than a minute ago"},
		{"a minute", at(70 * time.Second), "", "Last synced 1 minute ago"},
		{"minutes", at(4 * time.Minute), "", "Last synced 4 minutes ago"},
		{"under an hour", at(59 * time.Minute), "", "Last synced 59 minutes ago"},
		{"later today", seenToday, "", "Last synced 14:03"},
		{"yesterday", time.Date(2026, 11, 3, 19, 3, 0, 0, time.UTC), "", "Last synced yesterday"},
		{"a few days", at(3*24*time.Hour + time.Hour), "", "Last synced 3 days ago"},
		{"eighteen days", at(18 * 24 * time.Hour), "", "Last synced 18 days ago"},
		{"queued", seenToday, addon.SimQueued, "Last synced 14:03" + app.CopySeparator + app.CopySimQueued},
		{"premium", seenToday, addon.SimNeedsPremium, "Last synced 14:03" + app.CopySeparator + app.CopySimNeedsPremium},
		{"duplicate", seenToday, addon.SimDuplicate, "Last synced 14:03" + app.CopySeparator + app.CopySimDuplicate},
		{"invalid", seenToday, addon.SimInvalid, "Last synced 14:03" + app.CopySeparator + app.CopySimInvalid},
		{"a state with no words", seenToday, "unknown_character", "Last synced 14:03"},
	} {
		if got := app.AddonLine(c.seen, c.state, now); got != c.want {
			t.Errorf("%s: AddonLine = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestOnlyASyncOlderThanADayIsStale(t *testing.T) {
	now := time.Date(2026, 11, 4, 18, 0, 0, 0, time.UTC)
	for name, c := range map[string]struct {
		seen time.Time
		want bool
	}{
		"never":       {time.Time{}, false},
		"an hour":     {now.Add(-time.Hour), false},
		"a day":       {now.Add(-24 * time.Hour), false},
		"a day plus":  {now.Add(-24*time.Hour - time.Minute), true},
		"a fortnight": {now.Add(-14 * 24 * time.Hour), true},
	} {
		if got := app.AddonStale(c.seen, now); got != c.want {
			t.Errorf("%s: stale = %v, want %v", name, got, c.want)
		}
	}
}

// Every failure the window can be told has a sentence, and none of them leaks a cause.
func TestEveryFailureCodeHasASentenceThatCarriesNoCause(t *testing.T) {
	for _, code := range []string{
		app.CodeBadCode, app.CodeEmptyCode, app.CodeOffline, app.CodeServer, app.CodeSave,
		app.CodeNotWoW, app.CodeKnownFolder, app.CodeNoPicker, app.CodeNoCharacter,
		app.CodeSettings, app.CodeConnected, app.CodeNoConnect, app.CodeBadURL,
		app.CodeNoBrowser, app.CodeBadRequest, app.CodeInternal, app.CodeNotAdded,
	} {
		msg := app.Message(code)
		if msg == "" || !strings.HasSuffix(msg, ".") {
			t.Errorf("%s: message %q is not a sentence", code, msg)
		}
		if strings.Contains(msg, "api:") || strings.Contains(msg, "Error") || strings.Contains(msg, "TypeError") {
			t.Errorf("%s: message %q looks like an error's text", code, msg)
		}
	}
	if app.Message("no-such-code") != app.Message(app.CodeInternal) {
		t.Error("an unknown code does not fall back to the internal sentence")
	}
	if got := app.Message(app.CodeBadCode); got !=
		"That code is not valid. A code works once and expires after 10 minutes. Get a new one on the site." {
		t.Errorf("bad code = %q", got)
	}
	if got := app.Message(app.CodeNotWoW); got !=
		"That folder is not a World of Warcraft folder. Choose the one that contains WTF and Logs, such as _classic_era_." {
		t.Errorf("not wow = %q", got)
	}
}
