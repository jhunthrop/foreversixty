package addon

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// quietRig is a sync over one SavedVariables folder whose ForeverSixty.lua was last written
// at a time the test chooses.
type quietRig struct {
	t    *testing.T
	dir  string
	path string
	s    *Sync
	logs *bytes.Buffer
}

func newQuietRig(t *testing.T, savedAt time.Time) *quietRig {
	t.Helper()
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	r := &quietRig{t: t, dir: dir, path: filepath.Join(dir, SavedVariablesName+".lua"), logs: &bytes.Buffer{}}
	r.touch(SavedVariablesName+".lua", savedAt)
	api := &fakeAPI{inbox: Inbox{Builds: []Build{{ID: "b1", Name: "Holy", Code: "FSB1:x"}}}}
	r.s = New(Options{Paths: func() []string { return []string{r.path} }, API: api,
		Log: slog.New(slog.NewTextHandler(r.logs, nil))})
	return r
}

func (r *quietRig) touch(name string, at time.Time) {
	r.t.Helper()
	p := filepath.Join(r.dir, name)
	if _, err := os.Stat(p); err != nil {
		if err := os.WriteFile(p, []byte(savedVariables), 0o600); err != nil {
			r.t.Fatal(err)
		}
	}
	if err := os.Chtimes(p, at, at); err != nil {
		r.t.Fatal(err)
	}
}

func (r *quietRig) poll(at time.Time) {
	r.t.Helper()
	if err := r.s.Poll(r.t.Context(), at); err != nil {
		r.t.Fatal(err)
	}
}

func (r *quietRig) written() bool {
	_, err := os.Stat(InboxPath(r.path))
	return err == nil
}

func TestAQuietFolderGetsTheInboxAtOnce(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	r.poll(t0)
	if !r.written() {
		t.Fatal("a folder nothing touched for an hour did not get the inbox")
	}
	if strings.Contains(r.logs.String(), "deferred_ms") {
		t.Errorf("a write that did not wait logged deferred_ms:\n%s", r.logs)
	}
}

func TestAReloadDefersTheInboxUntilAMinuteAfterTheSavedVariablesChange(t *testing.T) {
	r := newQuietRig(t, t0)
	r.poll(t0.Add(time.Second))
	r.poll(t0.Add(30 * time.Second))
	r.poll(t0.Add(59 * time.Second))
	if r.written() {
		t.Fatal("the inbox was written inside the minute after ForeverSixty.lua changed")
	}
	r.poll(t0.Add(61 * time.Second))
	if !r.written() {
		t.Fatal("the deferred inbox never landed")
	}
	if !strings.Contains(r.logs.String(), "deferred_ms=60000") {
		t.Errorf("the log does not say how long the write waited:\n%s", r.logs)
	}
	if n := strings.Count(r.logs.String(), "wrote the addon inbox"); n != 1 {
		t.Errorf("the coalesced write was logged %d times", n)
	}
}

func TestABurstOfChangesKeepsDeferringTheInbox(t *testing.T) {
	r := newQuietRig(t, t0)
	r.poll(t0.Add(time.Second))
	for _, sec := range []int{50, 100, 150} {
		r.touch(SavedVariablesName+".lua", t0.Add(time.Duration(sec)*time.Second))
		r.poll(t0.Add(time.Duration(sec+30) * time.Second))
		if r.written() {
			t.Fatalf("the inbox landed 30 s after a change at +%ds", sec)
		}
	}
	r.poll(t0.Add(211 * time.Second))
	if !r.written() {
		t.Fatal("the inbox did not land once the folder went quiet")
	}
}

func TestAnyOtherFileInTheFolderDefersTheInboxForTwentySeconds(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	r.touch("Other.lua", t0)
	r.poll(t0.Add(10 * time.Second))
	if r.written() {
		t.Fatal("the inbox landed 10 s after another file in the folder changed")
	}
	r.poll(t0.Add(21 * time.Second))
	if !r.written() {
		t.Fatal("the inbox did not land after 20 s of quiet")
	}
}

func TestTheInboxIsWrittenThroughATemporaryFileThatDoesNotLinger(t *testing.T) {
	path := filepath.Join(t.TempDir(), InboxName+".lua")
	if err := os.WriteFile(path, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}
	if wrote, err := WriteInbox(path, []byte("new")); err != nil || !wrote {
		t.Fatalf("wrote=%v err=%v", wrote, err)
	}
	if b, _ := os.ReadFile(path); string(b) != "new" {
		t.Errorf("inbox = %q", b)
	}
	if _, err := os.Stat(path + ".tmp"); !os.IsNotExist(err) {
		t.Errorf("the temporary file lingers: %v", err)
	}
	entries, _ := os.ReadDir(filepath.Dir(path))
	if len(entries) != 1 {
		t.Errorf("folder holds %d entries, want only the inbox", len(entries))
	}
}
