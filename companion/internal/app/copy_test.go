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
		app.CopyNeverSynced:     "Not synced yet",
		app.CopyNeverSyncedHelp: "Install ForeverSixty, then log out once.",
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
		app.CodeGameRunning, app.CodeAddonDownload, app.CodeAddonInstall,
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
		"That folder is not a World of Warcraft folder. Forever Sixty runs on the Classic beta; pick its _classic_beta_ folder." {
		t.Errorf("not wow = %q", got)
	}
}

func TestTheSignedOutCopyIsTheDesignersWords(t *testing.T) {
	if app.CopySignedOutTitle != "Signed out by the site" ||
		app.CopySignedOutSub != "This computer was removed from your account, or the account was reset. "+
			"Nothing uploads until you sign in again." {
		t.Errorf("copy = %q / %q", app.CopySignedOutTitle, app.CopySignedOutSub)
	}
}

// The Addon row's version lines are the owner's words, one line per state, and
// never a bare number or an internal term.
func TestTheAddonVersionLinesAreTheRulingsWords(t *testing.T) {
	for got, want := range map[string]string{
		app.CopyVersionLine(app.CopyAddonName, "2.0.15", app.AddonCurrent):    "Addon 2.0.15 · up to date",
		app.CopyVersionLine(app.CopyAddonName, "2.0.16", app.AddonInstalled):  "Addon 2.0.16 installed · takes effect at your next login",
		app.CopyVersionLine(app.CopyAddonName, "2.0.16", app.AddonReady):      "Addon 2.0.16 ready · installs when the game closes",
		app.CopyVersionLine(app.CopyDataName, "2026.10.21", app.AddonCurrent): "Data 2026.10.21 · up to date",
		app.CopyAddonNotFound("Classic Era"):                                  "Addon not found in Classic Era",
		app.CopyInstallAddon("Classic Era"):                                   "Install the addon into Classic Era",
		app.CopyForInstall("Classic Era", "Addon 2.0.15 · up to date"):        "Classic Era: Addon 2.0.15 · up to date",
	} {
		if got != want {
			t.Errorf("copy = %q, want %q", got, want)
		}
		for _, banned := range []string{"manifest", "asset", "stage", "Stage"} {
			if strings.Contains(got, banned) {
				t.Errorf("%q says %q, which is not a player's word", got, banned)
			}
		}
	}
	for _, code := range []string{app.CodeGameRunning, app.CodeAddonDownload, app.CodeAddonInstall} {
		for _, banned := range []string{"manifest", "asset", "stage", "signature", "minisign"} {
			if strings.Contains(strings.ToLower(app.Message(code)), banned) {
				t.Errorf("%s: %q says %q", code, app.Message(code), banned)
			}
		}
	}
}
