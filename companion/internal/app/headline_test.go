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
// character is not matched yet, or that the name comes with the first fight; the dot is
// green, except amber when the log is not credited to anyone yet.
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
			"Uploading your fights as they end. Log out once in game with the addon loaded and it is credited to you."},
		{"no name read yet", Credit{State: CreditNone}, "Uploading your fights as they end. Your name is read from the first fight."},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := headlineFor(true, false, false, true, c.credit, watched)
			wantTitle := CopyLoggingTitle
			if c.credit.State == CreditUnmatched {
				wantTitle = "Logging · not credited to you yet"
			}
			wantTone := "on"
			if c.credit.State == CreditUnmatched {
				wantTone = "warn"
			}
			if got.Key != "logging" || got.Tone != wantTone || got.Title != wantTitle {
				t.Errorf("headline = %+v, want the %s Logging state", got, wantTone)
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
// sends, one per state the window draws, written from the real headline builder so a change to
// the words fails here before the JavaScript tests read stale ones. UPDATE_FIXTURES=1 rewrites
// them. The times are fixed around 17:41 on 20 Oct 2026; the capture rig shifts them to now.
func TestTheStatusFixturesAreWhatTheCompanionSends(t *testing.T) {
	watched := time.Date(2026, 10, 20, 17, 40, 0, 0, time.UTC)
	credited := Credit{Name: "Bow Jackzon", State: CreditCredited}
	unmatched := Credit{Name: "Bow Jackzon", State: CreditUnmatched}
	none := Credit{State: CreditNone}
	at := func(h, m, s int) *time.Time {
		v := time.Date(2026, 10, 20, h, m, s, 0, time.UTC)
		return &v
	}

	// base is a paired companion that is logging a report credited to credit.
	base := func(credit Credit) Snapshot {
		installs := []InstallStatus{
			{Path: `C:\Program Files (x86)\World of Warcraft\_classic_era_`, Flavor: "_classic_era_",
				FlavorName: "Classic Era", Watching: true, Added: true, LastPlayed: &watched,
				Advanced: true, AdvancedKnown: true},
			{Path: `C:\Program Files (x86)\World of Warcraft\_anniversary_`, Flavor: "_anniversary_",
				FlavorName: "Classic Anniversary", AdvancedKnown: true},
		}
		s := Snapshot{
			AppVersion: "0.2.13", EngineVersion: "1.0.0", Paired: true, AccountLabel: "Jhunt",
			DeviceName: "JH-DESKTOP", TokenBackend: "keychain", TokenBackendName: "Windows Credential Manager",
			TokenBackendNote: "A private store in your user profile.",
			APIBaseURL:       "https://api.foreversixty.gg", SiteBaseURL: "https://foreversixty.gg",
			Visibility: "public", DetectedCharacter: credit.Name,
			Installs: installs,
			Pipeline: pipeline.Status{Logging: true, ReportKey: "20261020-174100", ReportID: "r_9f2c",
				Detected: credit.Name},
			Reports:    []ReportRow{{Key: "20261020-174100", ReportID: "r_9f2c", Fights: 0, StartedAt: watched, Credit: credit}},
			Characters: []addon.Export{},
			AddonLine:  "Last synced 1 minute ago", Credit: credit,
			LastSync: LastSyncView{State: SyncSent, Count: 2, At: &watched, Title: "Tue 20 Oct 2026, 17:40"},
		}
		s.Headline = headlineFor(true, false, false, true, credit, installs)
		return s
	}
	// headline recomputes the headline from what a case changed.
	headline := func(s *Snapshot) {
		connecting := s.Connect.State == ConnectWaiting || s.Connect.State == ConnectBlocked
		s.Headline = headlineFor(s.Paired, connecting, s.SignedOut != nil, s.Pipeline.Logging, s.Credit, s.Installs)
	}
	// unpaired is a companion that has never signed in.
	unpaired := func() Snapshot {
		s := base(none)
		s.Paired, s.AccountLabel, s.Pipeline, s.Reports, s.Credit = false, "", pipeline.Status{}, []ReportRow{}, none
		s.AddonLine, s.AddonNever, s.AddonHelp = CopyNeverSynced, true, CopyNeverSyncedHelp
		s.LastSync = LastSyncView{State: SyncAbsent}
		headline(&s)
		return s
	}
	waiting := func(expires *time.Time) Snapshot {
		s := unpaired()
		s.Connect = ConnectView{State: ConnectWaiting, UserCode: "KQ4-7TM", URL: "https://foreversixty.gg/connect/KQ4-7TM", Expires: expires}
		headline(&s)
		return s
	}
	// rows is a mixed list: open, credited, not credited and never uploaded.
	rows := func() []ReportRow {
		zones := []string{"Molten Core", "Blackwing Lair", "Onyxia's Lair", "Zul'Gurub", "Ruins of Ahn'Qiraj", "Temple of Ahn'Qiraj", "Naxxramas", "Scholomance", "Stratholme", "Dire Maul"}
		out := []ReportRow{}
		for i, z := range zones {
			r := ReportRow{Key: "k" + z, ReportID: "r_" + z, Zone: z, Fights: 3 + i, StartedAt: watched.Add(-time.Duration(i) * 26 * time.Hour),
				Done: true, URL: "https://foreversixty.gg/reports/r_" + z, Credit: credited}
			switch i {
			case 0:
				r.Done = false
			case 1:
				r.Credit = unmatched
			case 3:
				r.Credit, r.URL, r.ReportID = unmatched, "", ""
			case 5:
				r.Zone = ""
			}
			out = append(out, r)
		}
		return out
	}

	for _, c := range []struct {
		file  string
		build func() Snapshot
	}{
		{"status-credited.json", func() Snapshot { return base(credited) }},
		{"status-not-credited.json", func() Snapshot { return base(unmatched) }},
		{"status-no-name.json", func() Snapshot { return base(none) }},
		{"status-ready.json", func() Snapshot {
			s := base(none)
			s.Pipeline = pipeline.Status{}
			s.Credit = none
			headline(&s)
			return s
		}},
		// The first seconds of a report: three fights closed, none on the site yet.
		{"status-uploading.json", func() Snapshot {
			s := base(unmatched)
			s.Pipeline.Fights, s.Pipeline.Queued, s.Pipeline.QueuedFights = 3, 5, 3
			return s
		}},
		// Some fights up, two still waiting: the two rows count different fights.
		{"status-uploading-some.json", func() Snapshot {
			s := base(credited)
			s.Pipeline.Fights, s.Pipeline.Queued, s.Pipeline.QueuedFights = 5, 4, 2
			return s
		}},
		{"status-advanced-off.json", func() Snapshot {
			s := base(none)
			s.Installs[0].Advanced = false
			s.Installs[0].AdvancedLogHelp = CopyAdvancedRowHelp
			s.Pipeline = pipeline.Status{}
			headline(&s)
			return s
		}},
		{"status-no-folder.json", func() Snapshot {
			s := base(none)
			s.Installs, s.Pipeline = []InstallStatus{}, pipeline.Status{}
			headline(&s)
			return s
		}},
		{"status-not-connected.json", unpaired},
		{"status-signed-out.json", func() Snapshot {
			s := unpaired()
			s.SignedOut = &SignedOut{Reason: "device_revoked", AccountLabel: "Jhunt", At: watched}
			headline(&s)
			return s
		}},
		{"status-connecting.json", func() Snapshot { return waiting(at(17, 50, 57)) }},
		{"status-addon-never.json", func() Snapshot {
			s := base(credited)
			s.AddonLine, s.AddonNever, s.AddonHelp = CopyNeverSynced, true, CopyNeverSyncedHelp
			s.LastSync = LastSyncView{State: SyncNone, Line: CopyNoCharactersSent}
			return s
		}},
		{"status-addon-stale.json", func() Snapshot {
			s := base(credited)
			s.AddonLine, s.AddonStale, s.AddonHelp = "Last synced 3 days ago", true, CopyStaleHelp
			return s
		}},
		{"reports-rows.json", func() Snapshot {
			s := base(credited)
			s.Reports = rows()
			return s
		}},
		{"reports-empty.json", func() Snapshot {
			s := base(none)
			s.Reports, s.Pipeline = []ReportRow{}, pipeline.Status{}
			headline(&s)
			return s
		}},
		{"device-waiting.json", func() Snapshot { return waiting(at(17, 49, 0)) }},
		{"device-waiting-last-minute.json", func() Snapshot { return waiting(at(17, 41, 40)) }},
		{"device-waiting-trouble.json", func() Snapshot {
			s := waiting(at(17, 49, 0))
			s.Connect.Trouble = true
			return s
		}},
		{"device-blocked.json", func() Snapshot {
			s := waiting(at(17, 49, 0))
			s.Connect.State = ConnectBlocked
			headline(&s)
			return s
		}},
		{"device-failed.json", func() Snapshot {
			s := unpaired()
			s.Connect = ConnectView{State: ConnectFailed, Reason: "expired", Message: "The sign-in link expired before you finished. Try again."}
			return s
		}},
		{"update-staged.json", func() Snapshot {
			s := base(credited)
			s.UpdateStaged = true
			return s
		}},
		{"update-arrived.json", func() Snapshot {
			s := base(credited)
			s.UpdatedTo = "0.2.13"
			return s
		}},
	} {
		t.Run(c.file, func(t *testing.T) {
			s := c.build()
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
