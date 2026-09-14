package app_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
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
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

type harness struct {
	t    *testing.T
	app  *app.App
	srv  *fakeapi.Server
	dirs paths.Dirs
	game string // the flavour directory
	log  string
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	home := t.TempDir()
	t.Setenv(paths.HomeEnv, home)
	dirs, err := paths.Resolve()
	if err != nil {
		t.Fatal(err)
	}
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
	srv := fakeapi.New()
	t.Cleanup(srv.Close)

	cfg := config.Default()
	cfg.APIBaseURL, cfg.SiteBaseURL = srv.URL, "https://foreversixty.gg"
	cfg.WoWPaths = []string{game}
	if err := config.Save(dirs.ConfigFile(), cfg); err != nil {
		t.Fatal(err)
	}
	a, err := app.New(app.Options{
		Dirs: dirs, Config: cfg,
		Secret: secret.ConfigFile{Path: dirs.ConfigFile()},
		Retry:  client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
		Now:    func() time.Time { return t0 },
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { a.Close() })
	return &harness{t: t, app: a, srv: srv, dirs: dirs, game: game,
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
