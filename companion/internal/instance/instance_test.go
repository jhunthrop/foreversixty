package instance_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/instance"
)

func TestASecondClaimOnTheSameHomeIsRefusedWhileTheFirstLives(t *testing.T) {
	home := t.TempDir()
	first, err := instance.Acquire(home)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := instance.Acquire(home); !errors.Is(err, instance.ErrRunning) {
		t.Fatalf("second Acquire = %v; want ErrRunning", err)
	}
	if err := first.Release(); err != nil {
		t.Fatal(err)
	}
	again, err := instance.Acquire(home)
	if err != nil {
		t.Fatalf("Acquire after Release: %v", err)
	}
	again.Release()
}

// A lock file a dead process left behind holds nothing: the operating
// system dropped the claim with the process.
func TestALockFileLeftByADeadProcessIsTakenOver(t *testing.T) {
	home := t.TempDir()
	if err := os.WriteFile(filepath.Join(home, instance.LockFile), []byte("999999\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	l, err := instance.Acquire(home)
	if err != nil {
		t.Fatalf("Acquire over a stale lock: %v", err)
	}
	defer l.Release()
	b, _ := os.ReadFile(filepath.Join(home, instance.LockFile))
	if strings.TrimSpace(string(b)) == "999999" {
		t.Error("the stale pid was left in the lock file")
	}
}

func TestThePublishedURLIsPrivateAndGoneAfterRelease(t *testing.T) {
	home := t.TempDir()
	l, err := instance.Acquire(home)
	if err != nil {
		t.Fatal(err)
	}
	if got := instance.RunningURL(home); got != "" {
		t.Fatalf("URL before publish = %q", got)
	}
	if err := l.PublishURL("http://127.0.0.1:5000/tok/"); err != nil {
		t.Fatal(err)
	}
	if got := instance.RunningURL(home); got != "http://127.0.0.1:5000/tok/" {
		t.Fatalf("URL = %q", got)
	}
	if fi, err := os.Stat(filepath.Join(home, instance.URLFile)); err != nil || fi.Mode().Perm()&0o077 != 0 && os.PathSeparator == '/' {
		t.Errorf("ui-url mode = %v, %v", fi, err)
	}
	l.Release()
	if got := instance.RunningURL(home); got != "" {
		t.Errorf("URL after release = %q", got)
	}
}
