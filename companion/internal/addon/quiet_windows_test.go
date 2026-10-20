//go:build windows

package addon

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// A file the game holds open (no delete sharing) cannot be renamed over: the write is
// deferred 30 s, tried again, and lands when the game lets go.
func TestAHeldInboxIsRetriedLater(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	inbox := InboxPath(r.path)
	if err := os.WriteFile(inbox, []byte("held"), 0o644); err != nil {
		t.Fatal(err)
	}
	held, err := os.Open(inbox)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := WriteInbox(inbox, []byte("new")); !errors.Is(err, ErrBusy) {
		t.Fatalf("err = %v, want ErrBusy", err)
	}
	r.poll(t0)
	r.poll(t0.Add(10 * time.Second))
	if b, _ := os.ReadFile(inbox); string(b) != "held" {
		t.Fatalf("the held inbox was replaced: %q", b)
	}
	if n := strings.Count(r.logs.String(), "trying again later"); n != 1 {
		t.Errorf("retry logged %d times inside the 30 s:\n%s", n, r.logs)
	}
	held.Close()
	r.poll(t0.Add(31 * time.Second))
	if b, _ := os.ReadFile(inbox); string(b) == "held" {
		t.Fatal("the inbox never landed after the game let go")
	}
	if _, err := os.Stat(filepath.Join(r.dir, InboxName+".lua.tmp")); !os.IsNotExist(err) {
		t.Errorf("temporary file lingers: %v", err)
	}
}
