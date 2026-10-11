// companion/internal/addon/sync_outbox_test.go
// The outbox and inbox v2 round trip, driven through the real Sync.Poll and the real client
// against the repo's fakeapi server: the outbox goes up with the characters, the site's
// answer is logged and kept for the status page, and the inbox v2 comes back down with the
// ack the fake writes once it has stored the export.
package addon_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
)

var t0 = time.Date(2026, 11, 4, 20, 0, 0, 0, time.UTC)

// seen is when the fake says it stored the export.
var seen = time.Date(2026, 11, 4, 14, 3, 0, 0, time.UTC)

// outboxSaved is a saved-variables file with one character and the outbox the addon leaves at
// logout; requestedAt is the sim request's stamp, so a test can make a new request.
func outboxSaved(requestedAt string, format int) string {
	return `
ForeverSixtyDB = {
	["characters"] = {
		["Obnoxious-Normal"] = {
			["name"] = "Obnoxious",
			["ruleset"] = "normal",
			["region"] = "us",
			["export"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
		},
	},
	["outbox"] = {
		["format"] = ` + string(rune('0'+format)) + `,
		["export"] = {
			["character"] = "Obnoxious-Yell",
			["code"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
			["saved_at"] = "2026-11-04 14:02",
		},
		["sim_request"] = {
			["character"] = "Obnoxious-Yell",
			["spec"] = "warrior-arms",
			["requested_at"] = "` + requestedAt + `",
			["gear_code"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
		},
		["roster_snapshot"] = {
			["guild"] = "OLYMPUS XXVII",
			["members"] = {
				{
					["name"] = "Obnoxious-Yell",
					["level"] = 20,
					["spec"] = "warrior-arms",
					["gear_gap"] = 14,
					["build"] = "Arms leveling",
					["seen_at"] = "2026-11-04 14:01",
				}, -- [1]
				{
					["level"] = 30,
				}, -- [2]
			},
		},
	},
}
`
}

// inboxV2 is an inbox v2 body whose first message is a sim result.
const inboxV2 = `{"format":2,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[
	{"type":"sim_result","character":"us/normal/obnoxious","spec":"warrior-arms","dps":412.6,"error":3.1,
	 "generated_at":"2026-11-04T20:55:00Z","slots":[{"slot":"chest","item_id":6627,"delta":4.2}]}
]}`

type rig struct {
	t   *testing.T
	api addon.API
	// state is the companion's state directory, which outlives a restart.
	state string
	srv   *fakeapi.Server
	sync  *addon.Sync
	path  string
	logs  *syncBuffer
	at    time.Time
}

// syncBuffer is a log sink the test reads while nothing else writes.
type syncBuffer struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (s *syncBuffer) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.Write(p)
}

func (s *syncBuffer) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.String()
}

func newRig(t *testing.T, saved string) *rig {
	t.Helper()
	srv := fakeapi.New()
	t.Cleanup(srv.Close)
	srv.SetClock(func() time.Time { return seen })
	c, err := client.New(client.Options{
		BaseURL: srv.URL,
		Token:   func() string { return fakeapi.Token },
		Retry:   client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
	})
	if err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	r := &rig{t: t, srv: srv, api: c, state: t.TempDir(),
		path: filepath.Join(dir, addon.SavedVariablesName+".lua"), logs: &syncBuffer{}, at: t0}
	r.restart()
	r.save(saved)
	return r
}

// restart is the companion quitting and starting again: a new Sync, the same files.
func (r *rig) restart() {
	r.sync = addon.New(addon.Options{
		Paths:    func() []string { return []string{r.path} },
		API:      r.api,
		Log:      slog.New(slog.NewTextHandler(r.logs, nil)),
		StateDir: r.state,
	})
}

// save writes the saved-variables file as the game does at logout: new content, a new
// modification time.
func (r *rig) save(body string) { r.t.Helper(); r.saveAfter(body, time.Hour) }

// saveAfter is save, d after the last one.
func (r *rig) saveAfter(body string, d time.Duration) {
	r.t.Helper()
	if err := os.WriteFile(r.path, []byte(body), 0o600); err != nil {
		r.t.Fatal(err)
	}
	r.at = r.at.Add(d)
	if err := os.Chtimes(r.path, r.at, r.at); err != nil {
		r.t.Fatal(err)
	}
}

func (r *rig) poll() error {
	r.t.Helper()
	return r.sync.Poll(r.t.Context(), r.at)
}

func (r *rig) posts() []fakeapi.ExportPost { return r.srv.ExportPosts() }

func TestAnOutboxSimRequestGoesUpWithTheCharactersAndTheAckComesBack(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
	r.srv.SetInbox(inboxV2)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}

	posts := r.posts()
	if len(posts) != 1 {
		t.Fatalf("posts = %d, want one", len(posts))
	}
	if len(posts[0].Characters) != 1 || posts[0].Characters[0]["name"] != "Obnoxious" {
		t.Errorf("the characters did not ride with the outbox: %+v", posts[0].Characters)
	}
	var ob struct {
		Format     int `json:"format"`
		SimRequest struct {
			Spec        string `json:"spec"`
			RequestedAt string `json:"requested_at"`
			GearCode    string `json:"gear_code"`
		} `json:"sim_request"`
		RosterSnapshot struct {
			Members []map[string]any `json:"members"`
		} `json:"roster_snapshot"`
	}
	if err := json.Unmarshal(posts[0].Outbox, &ob); err != nil {
		t.Fatalf("the post carried no outbox: %v (%s)", err, posts[0].Outbox)
	}
	if ob.Format != 1 || ob.SimRequest.Spec != "warrior-arms" || ob.SimRequest.RequestedAt != "2026-11-04 14:10" ||
		!strings.HasPrefix(ob.SimRequest.GearCode, "FS1:") || len(ob.RosterSnapshot.Members) != 2 {
		t.Errorf("outbox = %s", posts[0].Outbox)
	}

	// The site's answer, one plain line per part.
	logs := r.logs.String()
	for _, want := range []string{"Export stored", "Sim request queued", "Roster snapshot stored, 1 member dropped"} {
		if !strings.Contains(logs, want) {
			t.Errorf("the log does not say %q:\n%s", want, logs)
		}
	}

	// The inbox came down in the same pass, v2, with the fake's ack for the stored export.
	body, err := os.ReadFile(addon.InboxPath(r.path))
	if err != nil {
		t.Fatalf("the inbox was not written: %v", err)
	}
	ran, ok := addon.RunLua51(t, body, `
local I = assert(ForeverSixtyInbox)
assert(I.format == 2, "format")
assert(#I.messages == 2, "#messages = " .. #I.messages)
assert(I.messages[1].type == "sim_result", "messages[1].type = " .. tostring(I.messages[1].type))
assert(I.messages[1].slots[1].item_id == 6627, "slots")
local ack = I.messages[2]
assert(ack.type == "ack" and ack.character == "us/normal/obnoxious", "ack")
assert(ack.export_seen_at == "2026-11-04T14:03:00Z", "export_seen_at = " .. tostring(ack.export_seen_at))
assert(ack.sim_request_state == "queued", "sim_request_state")
`)
	if ok {
		t.Logf("Lua 5.1 check ran through %q", ran)
	} else if !strings.Contains(string(body), `["type"] = "ack"`) {
		t.Errorf("no ack in the inbox:\n%s", body)
	}

	st := r.sync.Status()
	if !st.SeenAt.Equal(seen) || st.SimRequestState != addon.SimQueued {
		t.Errorf("status = %+v", st)
	}
}

func TestAnUnchangedOutboxIsNotSentAgain(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	// Quiet passes send nothing at all.
	for range 3 {
		if err := r.sync.Poll(t.Context(), r.at.Add(time.Second)); err != nil {
			t.Fatal(err)
		}
	}
	if n := len(r.posts()); n != 1 {
		t.Fatalf("quiet passes posted: %d posts", n)
	}

	// The game rewrites the file with the same outbox: the characters go up, the outbox does not.
	r.save(outboxSaved("2026-11-04 14:10", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts := r.posts()
	if len(posts) != 2 || len(posts[1].Characters) != 1 || len(posts[1].Outbox) != 0 {
		t.Fatalf("second post = %+v", posts[len(posts)-1])
	}

	// A new sim request is a new outbox.
	r.save(outboxSaved("2026-11-04 15:30", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts = r.posts()
	if len(posts) != 3 || !strings.Contains(string(posts[2].Outbox), "2026-11-04 15:30") {
		t.Fatalf("the new request did not go up: %+v", posts[len(posts)-1])
	}
}

func TestAnOutboxOfAnotherFormatIsNotSent(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 2))
	before, err := os.ReadFile(r.path)
	if err != nil {
		t.Fatal(err)
	}
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts := r.posts()
	if len(posts) != 1 || len(posts[0].Characters) != 1 || len(posts[0].Outbox) != 0 {
		t.Fatalf("posts = %+v", posts)
	}
	after, err := os.ReadFile(r.path)
	if err != nil || !bytes.Equal(before, after) {
		t.Error("the saved-variables file was touched")
	}
	if st := r.sync.Status(); !st.SeenAt.IsZero() || st.SimRequestState != "" {
		t.Errorf("status without an outbox = %+v", st)
	}
}

func TestAnOutboxLostToTheNetworkIsSentOnTheNextPass(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
	r.srv.Offline(true)
	if err := r.poll(); err == nil {
		t.Fatal("an offline pass reported no error")
	}
	r.srv.Offline(false)
	if err := r.sync.Poll(t.Context(), r.at.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	posts := r.posts()
	if len(posts) != 1 || !strings.Contains(string(posts[0].Outbox), "sim_request") {
		t.Fatalf("the outbox was not retried: %+v", posts)
	}
}

func TestAnOutboxTheSiteRefusesIsNeverRetried(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
	r.srv.RefuseOutbox(true)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	// The characters still went up, alone, in the same pass.
	posts := r.posts()
	if len(posts) != 1 || len(posts[0].Characters) != 1 || len(posts[0].Outbox) != 0 {
		t.Fatalf("posts = %+v", posts)
	}
	if !strings.Contains(r.logs.String(), "The site could not read the outbox") {
		t.Errorf("the refusal was not logged:\n%s", r.logs.String())
	}
	// The same outbox on the next logout is not sent again.
	r.srv.RefuseOutbox(false)
	r.save(outboxSaved("2026-11-04 14:10", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts = r.posts()
	if len(posts) != 2 || len(posts[1].Outbox) != 0 {
		t.Fatalf("the refused outbox was sent again: %+v", posts[len(posts)-1])
	}
}

func TestASimRequestOutcomeIsLoggedAndKeptForTheStatusPage(t *testing.T) {
	for outcome, want := range map[string]struct{ log, state string }{
		"premium_required": {"Sim request needs a premium account", addon.SimNeedsPremium},
		"duplicate":        {"Sim request already running this hour", addon.SimDuplicate},
		"invalid":          {"Sim request invalid", addon.SimInvalid},
	} {
		t.Run(outcome, func(t *testing.T) {
			r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
			r.srv.SetSimOutcome(outcome)
			if err := r.poll(); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(r.logs.String(), want.log) {
				t.Errorf("the log does not say %q:\n%s", want.log, r.logs.String())
			}
			if st := r.sync.Status(); st.SimRequestState != want.state || !st.SeenAt.Equal(seen) {
				t.Errorf("status = %+v, want %s", st, want.state)
			}
		})
	}
}

// Fix round 1 -------------------------------------------------------------------------------

// noOutbox is outboxSaved with an outbox the companion does not send (format 2).
var noOutbox = outboxSaved("2026-11-04 14:10", 2)

func TestTheAckLandsInTheSamePassAsTheOutbox(t *testing.T) {
	r := newRig(t, noOutbox)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	fetches := r.srv.InboxFetches()
	if fetches != 1 {
		t.Fatalf("the first pass fetched the inbox %d times", fetches)
	}

	// A logout that changes only the characters, inside the ten minutes: the timer holds.
	r.saveAfter(strings.Replace(noOutbox, `["region"] = "us",`, `["region"] = "us", ["level"] = 21,`, 1), time.Minute)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if got := r.srv.InboxFetches(); got != fetches {
		t.Fatalf("a characters-only upload fetched the inbox early: %d fetches", got)
	}

	// A logout with an outbox, still inside the ten minutes: the inbox comes down at once.
	r.saveAfter(outboxSaved("2026-11-04 14:10", 1), time.Minute)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if got := r.srv.InboxFetches(); got != fetches+1 {
		t.Fatalf("the outbox upload did not fetch the inbox in the same pass: %d fetches", got)
	}
	body, err := os.ReadFile(addon.InboxPath(r.path))
	if err != nil || !strings.Contains(string(body), `["export_seen_at"] = "2026-11-04T14:03:00Z"`) {
		t.Fatalf("the ack is not in the inbox (%v):\n%s", err, body)
	}
	if st := r.sync.Status(); !st.SeenAt.Equal(seen) {
		t.Errorf("status = %+v", st)
	}
}

func TestARestartNeverResendsTheOutbox(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if posts := r.posts(); len(posts) != 1 || len(posts[0].Outbox) == 0 {
		t.Fatalf("posts = %+v", posts)
	}

	// The companion restarts over the same state directory and the same saved variables.
	r.restart()
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	for i, p := range r.posts()[1:] {
		if len(p.Outbox) != 0 {
			t.Fatalf("post %d after the restart resent the outbox: %s", i+2, p.Outbox)
		}
	}
	// A new outbox after the restart still goes up.
	r.save(outboxSaved("2026-11-04 15:30", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if posts := r.posts(); !strings.Contains(string(posts[len(posts)-1].Outbox), "15:30") {
		t.Fatalf("the new outbox did not go up: %+v", posts[len(posts)-1])
	}
}

func TestANewRosterIsANewOutbox(t *testing.T) {
	r := newRig(t, outboxSaved("2026-11-04 14:10", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	// Same saved_at, same requested_at, a member's build changed: a different outbox.
	r.save(strings.Replace(outboxSaved("2026-11-04 14:10", 1), "Arms leveling", "Arms raid", 1))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts := r.posts()
	if len(posts) != 2 || !strings.Contains(string(posts[1].Outbox), "Arms raid") {
		t.Fatalf("the changed roster did not go up: %+v", posts[len(posts)-1])
	}
}

// paddedOutbox is outboxSaved with a roster snapshot of n long-named members.
func paddedOutbox(n int) string {
	var pad strings.Builder
	for i := range n {
		fmt.Fprintf(&pad, "\t\t\t\t{ [\"name\"] = \"Member%04d-Yell\", [\"level\"] = 60, [\"spec\"] = \"warrior-arms\", "+
			"[\"gear_gap\"] = 0, [\"build\"] = \"%s\", [\"seen_at\"] = \"2026-11-04 14:01\" },\n", i, strings.Repeat("x", 120))
	}
	return strings.Replace(outboxSaved("2026-11-04 14:10", 1), "[\"members\"] = {\n", "[\"members\"] = {\n"+pad.String(), 1)
}

func TestAnOutboxTooLargeIsDroppedAndTheExportsStillGoUp(t *testing.T) {
	big := paddedOutbox(2000)
	r := newRig(t, big)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	posts := r.posts()
	if len(posts) != 1 || len(posts[0].Characters) != 1 || len(posts[0].Outbox) != 0 {
		t.Fatalf("posts = %+v", posts)
	}
	const line = "Outbox too large to send; the character exports went up alone"
	if n := strings.Count(r.logs.String(), line); n != 1 {
		t.Fatalf("the log says %q %d times:\n%s", line, n, r.logs.String())
	}
	// Marked answered: the next logout with the same outbox does not try it again.
	r.save(big)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if posts := r.posts(); len(posts) != 2 || len(posts[1].Outbox) != 0 {
		t.Fatalf("posts = %+v", posts)
	}
	if n := strings.Count(r.logs.String(), line); n != 1 {
		t.Errorf("the drop was logged %d times", n)
	}
	// Just under the bound still goes up.
	r.save(paddedOutbox(100))
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	if posts := r.posts(); len(posts[len(posts)-1].Outbox) == 0 {
		t.Error("an outbox under the bound was dropped")
	}
}

func TestASkippedInboxMessageIsLoggedOncePerProcess(t *testing.T) {
	r := newRig(t, noOutbox)
	r.srv.SetInbox(`{"format":2,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[
		{"type":"future_nested","character":"us/normal/obnoxious","rows":[1,2]}]}`)
	for i := range 3 {
		if err := r.sync.Poll(t.Context(), r.at.Add(time.Duration(i)*addon.InboxEvery)); err != nil {
			t.Fatal(err)
		}
	}
	if r.srv.InboxFetches() != 3 {
		t.Fatalf("fetches = %d", r.srv.InboxFetches())
	}
	if n := strings.Count(r.logs.String(), "skipped an inbox message"); n != 1 {
		t.Errorf("the skip was logged %d times:\n%s", n, r.logs.String())
	}
}

func TestAnInboxOfAFormatTheCompanionDoesNotKnowIsNotWritten(t *testing.T) {
	r := newRig(t, noOutbox)
	if err := r.poll(); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(addon.InboxPath(r.path))
	if err != nil {
		t.Fatal(err)
	}
	r.srv.SetInbox(`{"format":3,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[
		{"type":"ack","character":"us/normal/obnoxious","export_seen_at":"2026-11-04T14:03:00Z"}]}`)
	for i := 1; i <= 2; i++ {
		if err := r.sync.Poll(t.Context(), r.at.Add(time.Duration(i)*addon.InboxEvery)); err != nil {
			t.Fatal(err)
		}
	}
	after, err := os.ReadFile(addon.InboxPath(r.path))
	if err != nil || !bytes.Equal(before, after) {
		t.Errorf("the inbox was rewritten from a format the companion does not know:\n%s", after)
	}
	if n := strings.Count(r.logs.String(), "inbox format"); n != 1 {
		t.Errorf("the unknown format was logged %d times:\n%s", n, r.logs.String())
	}
	if st := r.sync.Status(); !st.SeenAt.IsZero() {
		t.Errorf("an ack in an unknown format reached the status: %+v", st)
	}
}
