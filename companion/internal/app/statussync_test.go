package app_test

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

func (h *harness) pairAndPrime() {
	h.t.Helper()
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		h.t.Fatal("pairing failed")
	}
	h.app.Step(h.t.Context(), t0)
}

func (h *harness) sv() string {
	return filepath.Join(h.game, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
}

func (h *harness) writeCharacters(n int) {
	h.t.Helper()
	var b strings.Builder
	b.WriteString(`ForeverSixtyDB = { ["characters"] = { `)
	for i := range n {
		name := "Alt" + string(rune('A'+i))
		b.WriteString(`["` + name + `-Hardcore"] = { ["name"] = "` + name + `", ["realm"] = "Hardcore", ` +
			`["region"] = "us", ["export"] = "FS1:1.60.1.70291:paladin:human:0/0/0:head=1" }, `)
	}
	b.WriteString(`} }`)
	if err := os.WriteFile(h.sv(), []byte(b.String()), 0o600); err != nil {
		h.t.Fatal(err)
	}
}

func TestStatusSaysWhatTheSiteHoldsAfterTheFirstSync(t *testing.T) {
	h := newHarness(t)
	// No addon folder at all: the install line on the Addon row says it.
	if s := h.app.Snapshot(); s.LastSync.State != app.SyncAbsent {
		t.Fatalf("with no addon: %+v", s.LastSync)
	}
	// The addon is installed and nothing has gone up.
	if err := os.MkdirAll(filepath.Join(h.game, "Interface", "AddOns", "ForeverSixty"), 0o755); err != nil {
		t.Fatal(err)
	}
	s := h.app.Snapshot()
	if s.LastSync.State != app.SyncNone || s.LastSync.Line != app.CopyNoCharactersSent {
		t.Fatalf("addon installed, nothing sent: %+v", s.LastSync)
	}
	h.pairAndPrime()
	h.writeCharacters(3)
	h.app.Step(t.Context(), t0.Add(time.Second))
	s = h.app.Snapshot()
	if s.LastSync.State != app.SyncSent || s.LastSync.Count != 3 || s.LastSync.At == nil || !s.LastSync.At.Equal(t0) ||
		s.LastSync.Title != "Wed 9 Dec 2026, 20:00" {
		t.Fatalf("after the upload: %+v", s.LastSync)
	}
	_, body := h.call(http.MethodGet, h.url("/api/status"), nil)
	ls, _ := body["data"].(map[string]any)["last_sync"].(map[string]any)
	if ls["state"] != "sent" || ls["count"] != float64(3) {
		t.Errorf("the status JSON carries %v", ls)
	}
}

func TestTheLastSyncSurvivesARestart(t *testing.T) {
	h := newHarness(t)
	h.pairAndPrime()
	h.writeCharacters(2)
	h.app.Step(t.Context(), t0.Add(time.Second))
	h.app.Close()

	// Same state directory, a new process, and the file unchanged: nothing is sent again,
	// yet the page still knows what the site holds.
	cfg, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	again, err := app.New(app.Options{Dirs: h.dirs, Config: cfg, Secret: h.secret,
		Now: func() time.Time { return t0.Add(2 * time.Minute) }})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { again.Close() })
	if s := again.Snapshot(); s.LastSync.State != app.SyncSent || s.LastSync.Count != 2 {
		t.Errorf("after a restart: %+v", s.LastSync)
	}
}

func TestTheSitesRefusalOfTheListKeepsItsOwnLine(t *testing.T) {
	h := newHarness(t)
	h.pairAndPrime()
	h.srv.RefuseCharacters(true)
	h.writeCharacters(2)
	h.app.Step(t.Context(), t0.Add(time.Second))
	s := h.app.Snapshot()
	want := "The site refused the last character list: that export cannot be stored"
	if s.LastSync.State != app.SyncRefused || s.LastSync.Line != want {
		t.Fatalf("after a refusal: %+v", s.LastSync)
	}
}

func TestAReportSaysWhoItIsCreditedTo(t *testing.T) {
	h := newHarness(t)
	h.pairAndPrime()
	if s := h.app.Snapshot(); s.Credit.State != app.CreditNone {
		t.Fatalf("with no report: %+v", s.Credit)
	}
	// The log names Morrowlyn and no addon record carries the realm.
	if err := os.WriteFile(h.log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	s := h.app.Snapshot()
	if s.Credit != (app.Credit{Name: "Morrowlyn", State: app.CreditUnmatched}) {
		t.Fatalf("unmatched: %+v", s.Credit)
	}
	// The window reads the two-state copy; the name-versus-realm split is the log's
	// (pipeline's logger_test checks that line).
	if s.Headline.Tone != "on" || s.Headline.Sub != "Uploading your fights as they end · character not matched yet" {
		t.Errorf("headline = %+v", s.Headline)
	}
	if len(s.Reports) != 1 || s.Reports[0].Credit != s.Credit {
		t.Errorf("the Reports row differs from the Current report: %+v", s.Reports)
	}

	// The player logs out with the addon loaded; the next step matches the record.
	if err := os.WriteFile(h.sv(), []byte(`ForeverSixtyDB = { ["characters"] = { `+
		`["Morrowlyn-Hardcore"] = { ["name"] = "Morrowlyn", ["realm"] = "Hardcore", `+
		`["region"] = "us", ["export"] = "FS1:1.60.1.70291:paladin:human:0/0/0:head=1" } } }`), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(2*time.Second))
	h.app.Step(t.Context(), t0.Add(3*time.Second))
	s = h.app.Snapshot()
	if s.Credit != (app.Credit{Name: "Morrowlyn", State: app.CreditCredited}) {
		t.Errorf("matched: %+v", s.Credit)
	}
	if s.Headline.Tone != "on" || s.Headline.Sub != "Uploading your fights as they end · logging as Morrowlyn" {
		t.Errorf("headline = %+v", s.Headline)
	}
}

func TestAReportWithNobodyNamedIsNotCredited(t *testing.T) {
	h := newHarness(t)
	h.pairAndPrime()
	// A report closed before the log ever named anyone.
	if err := state.Save(h.dirs.State, state.Report{Key: state.NewKey(t0), Auto: true,
		StartedAt: t0, Done: true}); err != nil {
		t.Fatal(err)
	}
	s := h.app.Snapshot()
	if len(s.Reports) != 1 || s.Reports[0].Credit != (app.Credit{State: app.CreditNone}) {
		t.Errorf("a log that names no one: %+v", s.Reports)
	}
}

func TestAChosenCharacterIsTheCredit(t *testing.T) {
	h := newHarness(t)
	h.pairAndPrime()
	if code, body := h.call(http.MethodPost, h.url("/api/settings"), map[string]any{
		"visibility":        "public",
		"logging_character": map[string]string{"region": "us", "ruleset": "hardcore", "name": "Bow Jackzon"},
	}); code != http.StatusOK {
		t.Fatalf("settings = %d %v", code, body)
	}
	if err := os.WriteFile(h.log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	if c := h.app.Snapshot().Credit; c.State != app.CreditCredited || c.Name != "Bow Jackzon" {
		t.Errorf("chosen: %+v", c)
	}
}

// A report an older companion stored under "Bow-Jackzon" reads as
// "Bow Jackzon" in the headline and the credit row once the update
// starts, with the reason that names the right character.
func TestAReportStoredUnderTheOldTwoWordNameShowsTheSpacedName(t *testing.T) {
	h := newHarnessSeeded(t, func(d paths.Dirs) {
		if err := state.Save(d.State, state.Report{Key: state.NewKey(t0), Auto: true,
			DetectedName: "Bow-Jackzon", DetectedRealm: "ClassicBetaPvP",
			StartedAt: t0, Closed: true, Done: true}); err != nil {
			t.Fatal(err)
		}
	})
	h.pairAndPrime()
	s := h.app.Snapshot()
	if s.DetectedCharacter != "Bow Jackzon" {
		t.Errorf("detected_character = %q", s.DetectedCharacter)
	}
	if len(s.Reports) != 1 {
		t.Fatalf("%d reports", len(s.Reports))
	}
	c := s.Reports[0].Credit
	if c != (app.Credit{Name: "Bow Jackzon", State: app.CreditUnmatched}) {
		t.Errorf("credit = %+v", c)
	}
}

// The Check for builds now button: one inbox fetch, and the Addon row's second line has what it found.
func TestCheckForBuildsNowFetchesOnceAndSaysWhatItFound(t *testing.T) {
	h := newHarness(t)
	if code, body := h.call(http.MethodPost, h.url("/api/builds/check"), nil); code != http.StatusBadRequest ||
		errorCode(body) != "not_paired" {
		t.Fatalf("before pairing = %d %v", code, body)
	}
	h.pairAndPrime()
	before := h.srv.InboxFetches()
	if s := h.app.Snapshot(); s.BuildsChecked != nil {
		t.Fatalf("checked before anyone pressed the button: %+v", s.BuildsChecked)
	}

	h.srv.SetInbox(`{"format":2,"generated_at":"2026-12-09T20:00:00Z","builds":[` +
		`{"id":"b1","name":"Beast Mastery build","character":"us/pvp/bow-jackzon","code":"FSB1:x"}],"messages":[]}`)
	code, body := h.call(http.MethodPost, h.url("/api/builds/check"), nil)
	if code != http.StatusOK {
		t.Fatalf("check = %d %v", code, body)
	}
	if n := h.srv.InboxFetches() - before; n != 1 {
		t.Errorf("the button fetched %d times, want 1", n)
	}
	got := body["data"].(map[string]any)["builds_checked"].(map[string]any)
	if got["new_builds"] != float64(1) || got["for"].([]any)[0] != "Bow Jackzon" {
		t.Errorf("builds_checked = %v", got)
	}

	// A second press with the same inbox finds nothing new.
	_, body = h.call(http.MethodPost, h.url("/api/builds/check"), nil)
	got = body["data"].(map[string]any)["builds_checked"].(map[string]any)
	if got["new_builds"] != float64(0) {
		t.Errorf("an unchanged inbox reported %v new", got["new_builds"])
	}
}
