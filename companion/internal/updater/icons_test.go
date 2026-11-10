package updater

import (
	"bytes"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// stubNotify stands in for the shell call and counts the calls.
func stubNotify(t *testing.T, err error) *int {
	t.Helper()
	calls := 0
	was := notifyIconChange
	notifyIconChange = func() error { calls++; return err }
	t.Cleanup(func() { notifyIconChange = was })
	return &calls
}

func fakeBundle(t *testing.T, old time.Time) (bundle, exe string) {
	t.Helper()
	bundle = filepath.Join(t.TempDir(), "Forever Sixty.app")
	exe = filepath.Join(bundle, "Contents", "MacOS", "foreversixty-companion")
	if err := os.MkdirAll(filepath.Dir(exe), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(exe, []byte("bin"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(bundle, old, old); err != nil {
		t.Fatal(err)
	}
	return bundle, exe
}

// After a Mac swap the bundle folder is newer than the icon LaunchServices
// cached, and the shell is told once.
func TestRefreshIconsTouchesTheBundleAndTellsTheShell(t *testing.T) {
	old := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Date(2026, 11, 9, 22, 0, 0, 0, time.UTC)
	bundle, exe := fakeBundle(t, old)
	calls := stubNotify(t, nil)
	refreshIcons(exe, slog.New(slog.DiscardHandler), now)
	fi, err := os.Stat(bundle)
	if err != nil {
		t.Fatal(err)
	}
	if !fi.ModTime().Equal(now) {
		t.Errorf("bundle mtime = %v, want %v", fi.ModTime(), now)
	}
	if *calls != 1 {
		t.Errorf("shell told %d times, want 1", *calls)
	}
}

// A bare binary (Windows, Linux) has no bundle to touch; the shell is
// still told.
func TestRefreshIconsForABareBinaryOnlyTellsTheShell(t *testing.T) {
	exe := filepath.Join(t.TempDir(), "foreversixty-companion.exe")
	calls := stubNotify(t, nil)
	refreshIcons(exe, slog.New(slog.DiscardHandler), time.Now())
	if *calls != 1 {
		t.Errorf("shell told %d times, want 1", *calls)
	}
}

// Both steps are best effort: a failure is a log line, never a panic or a
// failed update.
func TestRefreshIconsLogsAndCarriesOnWhenAStepFails(t *testing.T) {
	stubNotify(t, errors.New("shell32 has no SHChangeNotify"))
	var buf bytes.Buffer
	gone := filepath.Join(t.TempDir(), "Gone.app", "Contents", "MacOS", "foreversixty-companion")
	refreshIcons(gone, slog.New(slog.NewTextHandler(&buf, nil)), time.Now())
	out := buf.String()
	if !strings.Contains(out, "Dock may show the old icon") || !strings.Contains(out, "Explorer may show the old icon") {
		t.Errorf("log = %q, want both best-effort lines", out)
	}
}

// The real shell call is safe to make on every platform (a no-op off Windows).
func TestTheShellNotificationRuns(t *testing.T) {
	if err := shellNotifyIconChange(); err != nil {
		t.Errorf("shellNotifyIconChange = %v", err)
	}
}
