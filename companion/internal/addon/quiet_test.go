package addon

import (
	"bytes"
	"context"
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
	api := &countingAPI{fakeAPI: fakeAPI{inbox: Inbox{Builds: []Build{{ID: "b1", Name: "Holy", Code: "FSB1:x"}}}}}
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

func TestAReloadDefersTheInboxUntilAFewSecondsAfterTheSavedVariablesChange(t *testing.T) {
	r := newQuietRig(t, t0)
	r.poll(t0.Add(time.Second))
	r.poll(t0.Add(2 * time.Second))
	r.poll(t0.Add(2*time.Second + 900*time.Millisecond))
	if r.written() {
		t.Fatal("the inbox was written inside the three seconds after ForeverSixty.lua changed")
	}
	r.poll(t0.Add(4 * time.Second))
	if !r.written() {
		t.Fatal("the deferred inbox never landed")
	}
	if !strings.Contains(r.logs.String(), "deferred_ms=3000") {
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
		r.poll(t0.Add(time.Duration(sec+2) * time.Second))
		if r.written() {
			t.Fatalf("the inbox landed 2 s after a change at +%ds", sec)
		}
	}
	r.poll(t0.Add(154 * time.Second))
	if !r.written() {
		t.Fatal("the inbox did not land once the folder went quiet")
	}
}

func TestAnyOtherFileInTheFolderDefersTheInboxForThreeSeconds(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	r.touch("Other.lua", t0)
	r.poll(t0.Add(2 * time.Second))
	if r.written() {
		t.Fatal("the inbox landed 2 s after another file in the folder changed")
	}
	r.poll(t0.Add(4 * time.Second))
	if !r.written() {
		t.Fatal("the inbox did not land after 3 s of quiet")
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

// A logout's export is followed by one inbox fetch at the first quiet pass, so the file the next
// login reads is fresh, and the fetch's own inbox write does not ask for another.
func TestALogoutExportIsFollowedByOneInboxFetchWhenTheFolderGoesQuiet(t *testing.T) {
	r := newQuietRig(t, t0)
	api := r.s.o.API.(*countingAPI)
	r.poll(t0.Add(time.Second)) // the export goes up; the first ever pass also fetches
	if api.fetches != 1 {
		t.Fatalf("first pass fetched %d times", api.fetches)
	}
	r.poll(t0.Add(2 * time.Second)) // not quiet, timer not due
	if api.fetches != 1 {
		t.Fatalf("fetched inside the quiet period: %d", api.fetches)
	}
	r.poll(t0.Add(QuietSaved + time.Second)) // quiet: the logout fetch, and the write lands
	if api.fetches != 2 || !r.written() {
		t.Fatalf("fetches = %d, written = %v after the folder went quiet", api.fetches, r.written())
	}
	for _, sec := range []int{70, 90, 200, 400} { // the inbox write is not an export: no loop
		r.poll(t0.Add(time.Duration(sec) * time.Second))
	}
	if api.fetches != 2 {
		t.Errorf("fetches = %d, want exactly the one extra per logout", api.fetches)
	}
	// The next logout is one more.
	r.touch(SavedVariablesName+".lua", t0.Add(500*time.Second))
	r.poll(t0.Add(501 * time.Second))
	r.poll(t0.Add(500*time.Second + QuietSaved + time.Second))
	if api.fetches != 3 {
		t.Errorf("fetches = %d after a second logout, want 3", api.fetches)
	}
}

// Check for builds now fetches at once, whatever the timer says, and reports what was new.
func TestCheckNowFetchesAtOnceAndReportsWhatIsNew(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	api := r.s.o.API.(*countingAPI)
	r.poll(t0)
	got, err := r.s.CheckNow(t.Context(), t0.Add(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if api.fetches != 2 || got.New != 0 {
		t.Fatalf("fetches = %d, found %+v: the build was already rendered", api.fetches, got)
	}
	api.inbox.Builds = append(api.inbox.Builds, Build{ID: "b2", Name: "Beast Mastery build",
		Character: "us/pvp/bow-jackzon", Code: "FSB1:y"})
	got, err = r.s.CheckNow(t.Context(), t0.Add(2*time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if got.New != 1 || len(got.For) != 1 || got.For[0] != "Bow Jackzon" {
		t.Errorf("found %+v", got)
	}
	if r.s.Status().Checked.New != 1 {
		t.Errorf("status = %+v", r.s.Status().Checked)
	}
}

func TestACharacterKeyReadsAsTheNameThePlayerKnows(t *testing.T) {
	for key, want := range map[string]string{
		"us/pvp/bow-jackzon": "Bow Jackzon", "us/normal/obnoxious": "Obnoxious", "": "", "us/pvp/": "",
	} {
		if got := characterName(key); got != want {
			t.Errorf("characterName(%q) = %q, want %q", key, got, want)
		}
	}
}

// countingAPI is the fixed-inbox API that counts the inbox fetches.
type countingAPI struct {
	fakeAPI
	fetches int
}

func (c *countingAPI) AddonInbox(ctx context.Context) (Inbox, error) {
	c.fetches++
	return c.fakeAPI.AddonInbox(ctx)
}

// The walk of 22 Oct: a build sent in game is in the file the client rewrote at the logout
// within ten seconds of the burst, so a player who logs out and straight back in finds it.
func TestABuildSentInGameIsInTheInboxSecondsAfterTheLogoutBurst(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	api := r.s.o.API.(*countingAPI)
	r.poll(t0) // in game: the build is written
	if !r.written() {
		t.Fatal("the build was not written while the player was in game")
	}
	// The client rewrites every SavedVariables file and the inbox from memory (no builds).
	logout := t0.Add(time.Minute)
	if err := os.WriteFile(InboxPath(r.path), []byte("ForeverSixtyInbox = {builds = {}}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	r.touch("Other.lua", logout)
	r.touch(SavedVariablesName+".lua", logout.Add(time.Second)) // the burst ends one second later
	end := logout.Add(time.Second)
	var landed time.Time
	for sec := 1; sec <= 12; sec++ {
		at := end.Add(time.Duration(sec) * time.Second)
		r.poll(at)
		if b, _ := os.ReadFile(InboxPath(r.path)); strings.Contains(string(b), `"b1"`) {
			landed = at
			break
		}
	}
	if landed.IsZero() {
		t.Fatal("the inbox never came back after the logout")
	}
	if d := landed.Sub(end); d > 10*time.Second {
		t.Errorf("the build was back %v after the burst, want within 10 s", d)
	}
	if api.fetches != 2 {
		t.Errorf("fetches = %d, want the first pass and the one logout fetch", api.fetches)
	}
}

// A write never races a burst in progress: files that keep changing, even a second apart, hold it.
func TestTheInboxNeverLandsInsideABurstOfWrites(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	r.poll(t0)
	if err := os.Remove(InboxPath(r.path)); err != nil {
		t.Fatal(err)
	}
	r.s.pending[r.path] = &pendingWrite{since: t0}
	for sec := 1; sec <= 20; sec++ { // a file changes every other second
		at := t0.Add(time.Duration(sec) * time.Second)
		if sec%2 == 1 {
			r.touch("Other.lua", at)
		}
		r.poll(at)
		if r.written() {
			t.Fatalf("the inbox landed at +%ds inside a burst", sec)
		}
	}
	r.poll(t0.Add(24 * time.Second))
	if !r.written() {
		t.Fatal("the inbox did not land once the burst was over")
	}
}

// The busy-rename safety stays what it was: the retry is still half a minute, and the copy
// still waits for several held renames.
func TestTheHeldFileSafetyIsUnchanged(t *testing.T) {
	if BusyRetry != 30*time.Second || BusyTries != 5 {
		t.Errorf("BusyRetry = %v, BusyTries = %d", BusyRetry, BusyTries)
	}
	if QuietSaved > 5*time.Second || QuietFolder > 5*time.Second {
		t.Errorf("the quiet windows are %v and %v: the write is meant to follow the logout by seconds", QuietSaved, QuietFolder)
	}
}

// The Status card's waiting line: a build the companion found after its first look is waiting
// until a logout leaves the client's own copy of the inbox with the build in it.
func TestABuildIsWaitingUntilALogoutShowsTheAddonHeldIt(t *testing.T) {
	r := newQuietRig(t, t0.Add(-time.Hour))
	api := r.s.o.API.(*countingAPI)
	api.inbox = Inbox{}
	r.poll(t0) // the first look: an empty inbox, nothing waiting
	if st := r.s.Status(); st.Waiting != 0 {
		t.Fatalf("waiting = %d before any build was sent", st.Waiting)
	}
	api.inbox = Inbox{Builds: []Build{{ID: "b1", Name: "Holy", Character: "us/pvp/bow-jackzon", Code: "FSB1:x"}}}
	if _, err := r.s.CheckNow(t.Context(), t0.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	if st := r.s.Status(); st.Waiting != 1 || len(st.WaitingFor) != 1 || st.WaitingFor[0] != "Bow Jackzon" {
		t.Fatalf("status = %+v, want one build waiting for Bow Jackzon", st)
	}
	// A logout in which the addon never loaded it: the client writes the inbox without it.
	logout := t0.Add(time.Minute)
	if err := os.WriteFile(InboxPath(r.path), []byte("ForeverSixtyInbox = {builds = {}}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	r.touch(SavedVariablesName+".lua", logout)
	r.s.fresh[r.path] = true
	r.poll(logout.Add(5 * time.Second))
	if st := r.s.Status(); st.Waiting != 1 {
		t.Fatalf("waiting = %d after a logout that did not hold the build", st.Waiting)
	}
	// The next logout, after the addon loaded the file with the build: its dump has the id.
	logout2 := t0.Add(10 * time.Minute)
	if err := os.WriteFile(InboxPath(r.path), []byte("ForeverSixtyInbox = {\n[\"builds\"] = {{[\"id\"] = \"b1\"}},\n}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	r.touch(SavedVariablesName+".lua", logout2)
	r.s.fresh[r.path] = true
	r.poll(logout2.Add(5 * time.Second))
	if st := r.s.Status(); st.Waiting != 0 {
		t.Errorf("waiting = %d after a logout that held the build", st.Waiting)
	}
}
