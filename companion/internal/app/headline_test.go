// companion/internal/app/headline_test.go
package app

import "testing"

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
			got := headlineFor(c.paired, c.connecting, c.signedOut, c.logging, "", c.installs)
			if got.Key != c.want {
				t.Errorf("headline key = %q, want %q", got.Key, c.want)
			}
		})
	}
}
