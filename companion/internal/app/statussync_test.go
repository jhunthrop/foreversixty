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
	if s.Credit.State != app.CreditUnmatched || s.Credit.Name != "Morrowlyn" ||
		s.Credit.Reason != "the log names Morrowlyn, but no addon record carries that realm — "+
			"log out once with the addon loaded" {
		t.Fatalf("unmatched: %+v", s.Credit)
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
	if c := h.app.Snapshot().Credit; c.State != app.CreditCredited || c.Name != "Morrowlyn" || c.Reason != "" {
		t.Errorf("matched: %+v", c)
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
