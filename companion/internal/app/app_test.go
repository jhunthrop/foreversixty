package app_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

type harness struct {
	t      *testing.T
	app    *app.App
	srv    *fakeapi.Server
	dirs   paths.Dirs
	secret *countingSecret
	game   string // the flavour directory
	log    string
}

// countingSecret is the token store with a tally, so a test can hold
// the two-second poll's cost to account.
type countingSecret struct {
	secret.Store
	mu    sync.Mutex
	reads int
}

func (c *countingSecret) Token() (string, error) {
	c.mu.Lock()
	c.reads++
	c.mu.Unlock()
	return c.Store.Token()
}

func (c *countingSecret) count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.reads
}

// newGame builds a flavour directory wow.Pick accepts, with a Logs
// folder to tail and one account's SavedVariables.
func newGame(t *testing.T) string {
	t.Helper()
	game := filepath.Join(t.TempDir(), "_classic_era_")
	for _, p := range []string{
		filepath.Join(game, "Logs"),
		filepath.Join(game, "WTF", "Account", "ACCOUNT#1", "SavedVariables"),
	} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(game, "WTF", "Config.wtf"),
		[]byte("SET advancedCombatLogging \"1\"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return game
}

func newHarness(t *testing.T) *harness { return newHarnessSeeded(t, nil) }

// newHarnessSeeded builds the harness, calling seed with the resolved
// directories before the app is assembled, so a test can leave behind
// what a long-running install would have left behind.
func newHarnessSeeded(t *testing.T, seed func(paths.Dirs)) *harness {
	t.Helper()
	home := t.TempDir()
	t.Setenv(paths.HomeEnv, home)
	dirs, err := paths.Resolve()
	if err != nil {
		t.Fatal(err)
	}
	game := newGame(t)
	srv := fakeapi.New()
	t.Cleanup(srv.Close)

	cfg := config.Default()
	cfg.APIBaseURL, cfg.SiteBaseURL = srv.URL, "https://foreversixty.gg"
	cfg.WoWPaths = []string{game}
	if err := config.Save(dirs.ConfigFile(), cfg); err != nil {
		t.Fatal(err)
	}
	if seed != nil {
		seed(dirs)
	}
	store := &countingSecret{Store: secret.ConfigFile{Path: dirs.ConfigFile()}}
	a, err := app.New(app.Options{
		Dirs: dirs, Config: cfg,
		Secret: store,
		Retry:  client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
		Now:    func() time.Time { return t0 },
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { a.Close() })
	return &harness{t: t, app: a, srv: srv, dirs: dirs, secret: store, game: game,
		log: filepath.Join(game, "Logs", "WoWCombatLog.txt")}
}

// call performs one request against the local UI server.
func (h *harness) call(method, path string, body any) (int, map[string]any) {
	h.t.Helper()
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			h.t.Fatal(err)
		}
		r = strings.NewReader(string(b))
	}
	req := httptest.NewRequest(method, path, r)
	w := httptest.NewRecorder()
	h.app.Handler().ServeHTTP(w, req)
	var out map[string]any
	if w.Body.Len() > 0 && strings.HasPrefix(w.Header().Get("Content-Type"), "application/json") {
		if err := json.Unmarshal(w.Body.Bytes(), &out); err != nil {
			h.t.Fatalf("response is not JSON: %s", w.Body.String())
		}
	}
	return w.Code, out
}

// url is a path under the session token, read out of a served URL.
func (h *harness) url(path string) string {
	h.t.Helper()
	full, ln, err := h.app.Serve()
	if err != nil {
		h.t.Fatal(err)
	}
	h.t.Cleanup(func() { ln.Close() })
	// http://127.0.0.1:PORT/<token>/ → /<token>/<path>
	cut := strings.SplitN(strings.TrimPrefix(full, "http://"), "/", 2)
	return "/" + strings.TrimSuffix(cut[1], "/") + path
}

func TestTheUIIsOnlyReachableUnderTheSessionToken(t *testing.T) {
	h := newHarness(t)
	if code, _ := h.call(http.MethodGet, "/api/status", nil); code != http.StatusNotFound {
		t.Errorf("status without the token = %d", code)
	}
	if code, _ := h.call(http.MethodGet, "/some-other-token/api/status", nil); code != http.StatusNotFound {
		t.Errorf("status with a wrong token = %d", code)
	}
	if code, body := h.call(http.MethodGet, h.url("/api/status"), nil); code != http.StatusOK ||
		body["ok"] != true {
		t.Fatalf("status = %d %v", code, body)
	}
}

func TestTheIndexPageAndItsAssetsAreServed(t *testing.T) {
	h := newHarness(t)
	for _, path := range []string{"/", "/app.js", "/styles.css"} {
		req := httptest.NewRequest(http.MethodGet, h.url(path), nil)
		w := httptest.NewRecorder()
		h.app.Handler().ServeHTTP(w, req)
		if w.Code != http.StatusOK || w.Body.Len() == 0 {
			t.Errorf("%s = %d, %d bytes", path, w.Code, w.Body.Len())
		}
	}
	req := httptest.NewRequest(http.MethodGet, h.url("/"), nil)
	w := httptest.NewRecorder()
	h.app.Handler().ServeHTTP(w, req)
	if !strings.Contains(w.Body.String(), "Forever Sixty companion") {
		t.Error("the index page is not the companion's")
	}
}

func TestPairingStoresTheTokenAndUnpairingForgetsIt(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	_, body := h.call(http.MethodGet, base+"/api/status", nil)
	if body["data"].(map[string]any)["paired"] != false {
		t.Fatal("a fresh install reports itself paired")
	}
	code, body := h.call(http.MethodPost, base+"/api/pair", map[string]string{"code": fakeapi.PairCode})
	if code != http.StatusOK || body["ok"] != true {
		t.Fatalf("pair = %d %v", code, body)
	}
	if body["data"].(map[string]any)["paired"] != true {
		t.Fatal("pairing did not take")
	}
	stored, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	if stored.DeviceToken != fakeapi.Token {
		t.Fatalf("stored token = %q", stored.DeviceToken)
	}
	code, body = h.call(http.MethodPost, base+"/api/unpair", map[string]string{})
	if code != http.StatusOK || body["data"].(map[string]any)["paired"] != false {
		t.Fatalf("unpair = %d %v", code, body)
	}
}

func TestAWrongPairingCodeIsReportedToThePage(t *testing.T) {
	h := newHarness(t)
	code, body := h.call(http.MethodPost, h.url("/api/pair"), map[string]string{"code": "NOPE"})
	if code != http.StatusBadRequest || body["ok"] != false {
		t.Fatalf("pair = %d %v", code, body)
	}
	if msg := body["error"].(map[string]any)["message"].(string); msg == "" {
		t.Error("the failure carried no message")
	}
}

func TestSettingsAreSavedAndRefused(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	code, body := h.call(http.MethodPost, base+"/api/settings", app.Settings{
		Visibility: "unlisted", WoWPaths: []string{h.game},
	})
	if code != http.StatusOK || body["ok"] != true {
		t.Fatalf("settings = %d %v", code, body)
	}
	stored, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	if stored.ReportVisibility != "unlisted" {
		t.Fatalf("visibility = %q", stored.ReportVisibility)
	}
	if code, _ := h.call(http.MethodPost, base+"/api/settings", app.Settings{
		Visibility: "everyone", WoWPaths: []string{h.game}}); code != http.StatusBadRequest {
		t.Errorf("an unknown visibility was accepted: %d", code)
	}
	if code, _ := h.call(http.MethodPost, base+"/api/settings", app.Settings{
		Visibility: "public", WoWPaths: []string{t.TempDir()}}); code != http.StatusBadRequest {
		t.Errorf("a folder that is not the game was accepted: %d", code)
	}
	if code, _ := h.call(http.MethodPost, base+"/api/settings", map[string]any{
		"visibility": "public", "logging_character": map[string]string{"name": "Morrowlyn"},
	}); code != http.StatusBadRequest {
		t.Errorf("a half-filled character was accepted: %d", code)
	}
	if code, _ := h.call(http.MethodPost, base+"/api/settings", "not an object"); code != http.StatusBadRequest {
		t.Errorf("rubbish was accepted: %d", code)
	}
}

func TestTheSnapshotShowsInstallsCharactersAndAdvancedLogging(t *testing.T) {
	h := newHarness(t)
	sv := filepath.Join(h.game, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	if err := os.WriteFile(sv, []byte(`ForeverSixtyDB = { ["characters"] = { `+
		`["Morrowlyn-Hardcore"] = { ["name"] = "Morrowlyn", ["realm"] = "Hardcore", `+
		`["region"] = "us", ["export"] = "FS1:x" } } }`), 0o600); err != nil {
		t.Fatal(err)
	}
	s := h.app.Snapshot()
	if len(s.Installs) != 1 || s.Installs[0].Path != h.game {
		t.Fatalf("installs = %+v", s.Installs)
	}
	if !s.Installs[0].AdvancedKnown || !s.Installs[0].Advanced {
		t.Errorf("advanced logging = %+v", s.Installs[0])
	}
	if len(s.Characters) != 1 || s.Characters[0].Ruleset != "Hardcore" {
		t.Fatalf("characters = %+v", s.Characters)
	}
	if s.Pipeline.Logging {
		t.Error("nothing has been logged yet")
	}
}

func TestAStepLogsARaidAndListsTheReport(t *testing.T) {
	h := newHarness(t)
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	h.app.Step(t.Context(), t0) // prime the watcher over an empty folder

	if err := os.WriteFile(h.log,
		[]byte(fixture.Header+fixture.Zone+fixture.Encounter(0)+fixture.Heartbeat(0)), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))

	s := h.app.Snapshot()
	if !s.Pipeline.Logging || s.Pipeline.ReportID == "" {
		t.Fatalf("pipeline = %+v", s.Pipeline)
	}
	if len(s.Reports) != 1 || s.Reports[0].URL == "" {
		t.Fatalf("reports = %+v", s.Reports)
	}
	if !strings.HasPrefix(s.Reports[0].URL, "https://foreversixty.gg/reports/") {
		t.Errorf("report URL = %q", s.Reports[0].URL)
	}
	var fights int
	for _, e := range h.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights++
		}
	}
	if fights == 0 {
		t.Fatalf("no fight reached the server: %v", h.srv.Order())
	}
	// The addon inbox is written on the same step.
	inbox := addon.InboxPath(filepath.Join(h.game, "WTF", "Account", "ACCOUNT#1",
		"SavedVariables", "ForeverSixty.lua"))
	if _, err := os.Stat(inbox); err != nil {
		t.Errorf("the inbox was not written: %v", err)
	}
}

func TestAnUnpairedCompanionStillParsesButUploadsNothing(t *testing.T) {
	h := newHarness(t)
	h.app.Step(t.Context(), t0)
	if err := os.WriteFile(h.log,
		[]byte(fixture.Header+fixture.Zone+fixture.Encounter(0)+fixture.Heartbeat(0)), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	if len(h.srv.Order()) != 0 {
		t.Fatalf("an unpaired companion uploaded %v", h.srv.Order())
	}
	s := h.app.Snapshot()
	if !s.Pipeline.Logging {
		t.Error("an unpaired companion stopped parsing")
	}
	if s.Pipeline.Queued == 0 {
		t.Error("nothing was queued for when the device is paired")
	}
}

// A raid night's worth of log, enough to open a report and close one
// fight.
func aNight() string {
	return fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0)
}

func TestSavingSettingsKeepsTheDeviceTokenAndAppliesWithoutARestart(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	if code, _ := h.call(http.MethodPost, base+"/api/pair",
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	second := newGame(t)
	code, body := h.call(http.MethodPost, base+"/api/settings", app.Settings{
		Visibility: "private", WoWPaths: []string{h.game, second},
	})
	if code != http.StatusOK || body["ok"] != true {
		t.Fatalf("settings = %d %v", code, body)
	}

	// The save must not carry a stale blank over the token the
	// pairing wrote into the same file.
	stored, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	if stored.DeviceToken != fakeapi.Token {
		t.Fatalf("saving the settings unpaired the device: token = %q", stored.DeviceToken)
	}
	s := h.app.Snapshot()
	if !s.Paired {
		t.Error("the companion reports itself unpaired after a settings save")
	}
	if len(s.Installs) != 2 {
		t.Fatalf("the second game folder was not picked up: %+v", s.Installs)
	}

	// The new visibility reaches the next report with no restart.
	h.app.Step(t.Context(), t0)
	if err := os.WriteFile(h.log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	reports := h.srv.Reports()
	if len(reports) != 1 {
		t.Fatalf("the server holds %d reports", len(reports))
	}
	for _, r := range reports {
		if r.Visibility != "private" {
			t.Errorf("the report the save asked to be private is %q", r.Visibility)
		}
	}
}

func TestANewFirstGameFolderMovesTheTailOnTheNextStep(t *testing.T) {
	h := newHarness(t)
	second := newGame(t)
	if code, body := h.call(http.MethodPost, h.url("/api/settings"), app.Settings{
		Visibility: "public", WoWPaths: []string{second, h.game},
	}); code != http.StatusOK {
		t.Fatalf("settings = %d %v", code, body)
	}
	// One step to prime the tail over the folder that is now first.
	h.app.Step(t.Context(), t0)
	log := filepath.Join(second, "Logs", "WoWCombatLog.txt")
	if err := os.WriteFile(log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	if got := h.app.Snapshot().Pipeline.LogPath; got != log {
		t.Fatalf("the tail is on %q, not the newly added folder's log %q", got, log)
	}
}

func TestTheStatusPollDoesNotAskTheTokenStoreEveryTime(t *testing.T) {
	h := newHarness(t)
	h.app.Snapshot()
	settled := h.secret.count()
	for range 5 {
		h.app.Snapshot()
		h.app.Step(t.Context(), t0)
	}
	if got := h.secret.count(); got != settled {
		t.Errorf("ten more passes cost %d token reads", got-settled)
	}
	// Pairing must be visible immediately even so.
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	if !h.app.Snapshot().Paired {
		t.Fatal("the cached token outlived the pairing that replaced it")
	}
	if code, _ := h.call(http.MethodPost, h.url("/api/unpair"), map[string]string{}); code != http.StatusOK {
		t.Fatal("unpairing failed")
	}
	if h.app.Snapshot().Paired {
		t.Fatal("the cached token outlived the unpairing that cleared it")
	}
}

func TestReportsOlderThanTheRetentionAreForgottenAtStartup(t *testing.T) {
	old := t0.Add(-app.ReportRetention - 24*time.Hour)
	keys := struct{ stale, recent, unsent string }{
		"20200101-200000-aaaaaaaa", "20261201-200000-bbbbbbbb", "20200102-200000-cccccccc",
	}
	h := newHarnessSeeded(t, func(dirs paths.Dirs) {
		for _, r := range []state.Report{
			{Key: keys.stale, StartedAt: old, Closed: true, Done: true},
			{Key: keys.recent, StartedAt: t0.Add(-time.Hour), Closed: true, Done: true},
			// Old, but the server never accepted its completion: the
			// queue may still be holding work for it.
			{Key: keys.unsent, StartedAt: old, Closed: true},
		} {
			if err := state.Save(dirs.State, r); err != nil {
				t.Fatal(err)
			}
		}
	})
	left, err := state.List(h.dirs.State)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, r := range left {
		got = append(got, r.Key)
	}
	if len(got) != 2 || got[0] != keys.unsent || got[1] != keys.recent {
		t.Fatalf("state/ holds %v; the finished report from %s should be gone", got, old)
	}
	if n := len(h.app.Snapshot().Reports); n != 2 {
		t.Errorf("the Reports page shows %d reports", n)
	}
}
