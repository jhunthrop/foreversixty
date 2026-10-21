package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

func runningAs(t *testing.T, v string) {
	t.Helper()
	was := updater.Version
	updater.Version = v
	t.Cleanup(func() { updater.Version = was })
}

func recorded(t *testing.T, dirs paths.Dirs) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(dirs.Home, "last-version"))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestAFirstInstallIsNotAnUpdate(t *testing.T) {
	runningAs(t, "0.2.11")
	h := newHarness(t)
	if got := h.app.Snapshot().UpdatedTo; got != "" {
		t.Fatalf("updated_to = %q on a first install", got)
	}
	if got := recorded(t, h.dirs); got != "0.2.11\n" {
		t.Fatalf("recorded %q", got)
	}
}

func TestANewVersionSaysSoOnceAndTheNextLaunchDoesNot(t *testing.T) {
	runningAs(t, "0.2.12")
	h := newHarnessSeeded(t, func(d paths.Dirs) {
		if err := os.WriteFile(filepath.Join(d.Home, "last-version"), []byte("0.2.11\n"), 0o600); err != nil {
			t.Fatal(err)
		}
	})
	if got := h.app.Snapshot().UpdatedTo; got != "0.2.12" {
		t.Fatalf("updated_to = %q, want 0.2.12", got)
	}
	if got := recorded(t, h.dirs); got != "0.2.12\n" {
		t.Fatalf("recorded %q", got)
	}
	// The same version again, as the next launch: nothing to say.
	again := newHarnessSeeded(t, func(d paths.Dirs) {
		if err := os.WriteFile(filepath.Join(d.Home, "last-version"), []byte("0.2.12\n"), 0o600); err != nil {
			t.Fatal(err)
		}
	})
	if got := again.app.Snapshot().UpdatedTo; got != "" {
		t.Fatalf("updated_to = %q on the second launch", got)
	}
}

// Releases before the record existed left the tray notice behind; a home
// that has it was run by an older companion.
func TestAHomeThatRanAnOlderCompanionIsToldAboutTheUpdate(t *testing.T) {
	runningAs(t, "0.2.11")
	h := newHarnessSeeded(t, func(d paths.Dirs) {
		if err := os.WriteFile(filepath.Join(d.Home, "tray-notice-seen"), nil, 0o600); err != nil {
			t.Fatal(err)
		}
	})
	if got := h.app.Snapshot().UpdatedTo; got != "0.2.11" {
		t.Fatalf("updated_to = %q", got)
	}
}
