// companion/internal/app/headline_test.go
package app

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/pipeline"
)

// The headline judges the watched install's advanced-logging switch, not every folder on the disk.
func TestHeadlineJudgesTheWatchedInstall(t *testing.T) {
	on := func(watching, added bool) InstallStatus {
		return InstallStatus{Watching: watching, Added: added, Advanced: true, AdvancedKnown: true}
	}
	off := func(watching, added bool) InstallStatus {
		return InstallStatus{Watching: watching, Added: added, AdvancedKnown: true}
	}
	unknown := InstallStatus{Watching: true}

	for _, c := range []struct {
		name       string
		paired     bool
		connecting bool
		signedOut  bool
		logging    bool
		installs   []InstallStatus
		want       string
	}{
		{"watched on, another off", true, false, false, false, []InstallStatus{off(false, false), on(true, false), on(false, false)}, "ready"},
		{"watched on, another off, logging", true, false, false, true, []InstallStatus{off(false, false), on(true, false)}, "logging"},
		{"watched off", true, false, false, false, []InstallStatus{on(false, false), off(true, false)}, "advanced_off"},
		{"watched off, logging", true, false, false, true, []InstallStatus{off(true, false)}, "advanced_off"},
		{"watched unknown, another off", true, false, false, false, []InstallStatus{off(false, false), unknown}, "ready"},
		{"none watched, added off", true, false, false, false, []InstallStatus{on(false, false), off(false, true)}, "advanced_off"},
		{"none watched, added on, found off", true, false, false, false, []InstallStatus{off(false, false), on(false, true)}, "ready"},
		{"none watched, found only off", true, false, false, false, []InstallStatus{off(false, false), off(false, false)}, "ready"},
		{"watched on beats added off", true, false, false, false, []InstallStatus{on(true, false), off(false, true)}, "ready"},
		{"connecting", false, true, false, false, []InstallStatus{off(true, false)}, "connecting"},
		{"signed out", true, false, true, false, []InstallStatus{off(true, false)}, "signed_out"},
		{"not paired", false, false, false, false, []InstallStatus{off(true, false)}, "not_connected"},
		{"no folder", true, false, false, false, nil, "no_folder"},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := headlineFor(c.paired, c.connecting, c.signedOut, c.logging, Credit{State: CreditNone}, c.installs)
			if got.Key != c.want {
				t.Errorf("headline key = %q, want %q", got.Key, c.want)
			}
		})
	}
}

// The Logging headline's sub line says who the report is credited to, that the
// character is not matched yet, or nothing, and the dot is green in all three.
// A name stored in its log form prints with a space, never a hyphen.
func TestTheLoggingSubLineFollowsTheCreditState(t *testing.T) {
	watched := []InstallStatus{{Watching: true, Advanced: true, AdvancedKnown: true}}
	for _, c := range []struct {
		name   string
		credit Credit
		want   string
	}{
		{"credited", Credit{Name: "Bow Jackzon", State: CreditCredited},
			"Uploading your fights as they end · logging as Bow Jackzon"},
		{"credited, stored with the log's hyphen", Credit{Name: "Bow-Jackzon", State: CreditCredited},
			"Uploading your fights as they end · logging as Bow Jackzon"},
		{"a name read, not matched", Credit{Name: "Bow Jackzon", State: CreditUnmatched},
			"Uploading your fights as they end · character not matched yet"},
		{"no name read yet", Credit{State: CreditNone}, "Uploading your fights as they end"},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := headlineFor(true, false, false, true, c.credit, watched)
			if got.Key != "logging" || got.Tone != "on" || got.Title != CopyLoggingTitle {
				t.Errorf("headline = %+v, want the green Logging state", got)
			}
			if got.Sub != c.want {
				t.Errorf("sub = %q, want %q", got.Sub, c.want)
			}
			if strings.Contains(got.Sub, "-J") {
				t.Errorf("sub prints a hyphenated name: %q", got.Sub)
			}
		})
	}
}

// The status fixtures the window's tests read (ui/testdata) are the snapshots the companion
// sends in the three credit states, written from the real headline builder so a change to the
// words fails here before the JavaScript tests read stale ones. UPDATE_FIXTURES=1 rewrites them.
func TestTheStatusFixturesAreWhatTheCompanionSends(t *testing.T) {
	watched := time.Date(2026, 10, 20, 17, 40, 0, 0, time.UTC)
	for _, c := range []struct {
		file   string
		credit Credit
	}{
		{"status-credited.json", Credit{Name: "Bow Jackzon", State: CreditCredited}},
		{"status-not-credited.json", Credit{Name: "Bow Jackzon", State: CreditUnmatched}},
		{"status-no-name.json", Credit{State: CreditNone}},
	} {
		t.Run(c.file, func(t *testing.T) {
			installs := []InstallStatus{
				{Path: `C:\Program Files (x86)\World of Warcraft\_classic_era_`, Flavor: "_classic_era_",
					FlavorName: "Classic Era", Watching: true, Added: true, LastPlayed: &watched,
					Advanced: true, AdvancedKnown: true},
				{Path: `C:\Program Files (x86)\World of Warcraft\_anniversary_`, Flavor: "_anniversary_",
					FlavorName: "Classic Anniversary", AdvancedKnown: true},
			}
			s := Snapshot{
				AppVersion: "0.2.8", EngineVersion: "1.0.0", Paired: true, AccountLabel: "Jhunt",
				DeviceName: "JH-DESKTOP", TokenBackend: "file", TokenBackendName: "A private file in your user folder",
				APIBaseURL: "https://api.foreversixty.gg", SiteBaseURL: "https://foreversixty.gg",
				Visibility: "public", DetectedCharacter: c.credit.Name,
				Installs: installs,
				Pipeline: pipeline.Status{Logging: true, ReportKey: "20261020-174100", ReportID: "r_9f2c",
					Detected: c.credit.Name},
				Reports:    []ReportRow{{Key: "20261020-174100", ReportID: "r_9f2c", Fights: 0, StartedAt: watched, Credit: c.credit}},
				Characters: []addon.Export{},
				Headline:   headlineFor(true, false, false, true, c.credit, installs),
				AddonLine:  "Last synced 1 minute ago", Credit: c.credit,
				LastSync: LastSyncView{State: SyncSent, Count: 2, At: &watched, Title: "Tue 20 Oct 2026, 17:40"},
			}
			got, err := json.MarshalIndent(map[string]any{"ok": true, "data": s, "error": nil}, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			got = append(got, '\n')
			p := filepath.Join("..", "..", "ui", "testdata", c.file)
			if os.Getenv("UPDATE_FIXTURES") != "" {
				if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(p, got, 0o644); err != nil {
					t.Fatal(err)
				}
			}
			want, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			if strings.ReplaceAll(string(want), "\r\n", "\n") != string(got) {
				t.Errorf("%s is stale; run with UPDATE_FIXTURES=1", c.file)
			}
		})
	}
}
