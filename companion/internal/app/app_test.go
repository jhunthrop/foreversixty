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
	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
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

	// What the browser and the folder dialog were asked, and what they
	// answer.
	mu      sync.Mutex
	opened  []string
	openErr error
	pick    string
	pickErr error
}

func (h *harness) open(url string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.opened = append(h.opened, url)
	return h.openErr
}

func (h *harness) openedLinks() []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]string(nil), h.opened...)
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
	game := filepath.Join(t.TempDir(), "_classic_beta_")
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

func newHarness(t *testing.T) *harness { return newHarnessWith(t, nil, nil) }

// newHarnessSeeded builds the harness, calling seed with the resolved
// directories before the app is assembled, so a test can leave behind
// what a long-running install would have left behind.
func newHarnessSeeded(t *testing.T, seed func(paths.Dirs)) *harness {
	return newHarnessWith(t, seed, nil)
}

// newHarnessWith also lets a test adjust the app's options before it is built.
func newHarnessWith(t *testing.T, seed func(paths.Dirs), tweak func(*app.Options)) *harness {
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
	h := &harness{t: t, srv: srv, dirs: dirs, secret: store, game: game,
		log: filepath.Join(game, "Logs", "WoWCombatLog.txt")}
	opts := app.Options{
		Dirs: dirs, Config: cfg,
		Secret: store,
		Retry:  client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
		Now:    func() time.Time { return t0 },
		// Never a real browser or a real dialog from a test.
		Open:            h.open,
		PickFolder:      func() (string, error) { return h.pick, h.pickErr },
		ConnectInterval: 2 * time.Millisecond,
	}
	if tweak != nil {
		tweak(&opts)
	}
	a, err := app.New(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { a.Close() })
	h.app = a
	return h
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
	ui, err := h.app.Serve()
	if err != nil {
		h.t.Fatal(err)
	}
	h.t.Cleanup(func() { ui.Close() })
	// http://127.0.0.1:PORT/<token>/ → /<token>/<path>
	cut := strings.SplitN(strings.TrimPrefix(ui.URL, "http://"), "/", 2)
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
	for _, path := range []string{"/", "/app.js", "/copy.js", "/format.js", "/styles.css"} {
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
	if !strings.Contains(w.Body.String(), "<title>Forever Sixty</title>") {
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

// The logging character is read from the combat log: the snapshot
// names it, the Status headline says so, and a report is credited to
// the addon's record of that character without anyone picking it.
func TestTheLogNamesTheCharacterTheReportIsCreditedTo(t *testing.T) {
	h := newHarness(t)
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	sv := filepath.Join(h.game, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	if err := os.WriteFile(sv, []byte(`ForeverSixtyDB = { ["characters"] = { `+
		`["Morrowlyn-Hardcore"] = { ["name"] = "Morrowlyn", ["realm"] = "Hardcore", `+
		`["region"] = "us", ["export"] = "FS1:1.60.1.70291:paladin:human:0/0/0:head=1" } } }`), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0)
	if s := h.app.Snapshot(); s.DetectedCharacter != "" || s.LoggingCharacter != nil {
		t.Fatalf("before any log: detected %q, chosen %+v", s.DetectedCharacter, s.LoggingCharacter)
	}
	if err := os.WriteFile(h.log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))

	s := h.app.Snapshot()
	if s.DetectedCharacter != "Morrowlyn" {
		t.Errorf("detected_character = %q", s.DetectedCharacter)
	}
	if s.LoggingCharacter != nil {
		t.Errorf("the chosen character is %+v; Automatic must stay unset", s.LoggingCharacter)
	}
	if s.Headline.Key != "logging" ||
		s.Headline.Sub != "Uploading your fights as they end · logging as Morrowlyn" {
		t.Errorf("headline = %+v", s.Headline)
	}
	reps := h.srv.Reports()
	if len(reps) != 1 {
		t.Fatalf("%d reports", len(reps))
	}
	for _, rep := range reps {
		want := character.Character{Region: "us", Ruleset: "hardcore", Name: "Morrowlyn"}
		if rep.Created == nil || *rep.Created != want {
			t.Errorf("the report was created for %+v, want %+v", rep.Created, want)
		}
	}
}

func TestPickingAndClearingTheLoggingCharacterRoundTrips(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	pick := &character.Character{Region: "us", Ruleset: "hardcore", Name: "Morrowlyn"}
	if code, body := h.call(http.MethodPost, base+"/api/settings", app.Settings{
		Visibility: "public", LoggingCharacter: pick}); code != http.StatusOK {
		t.Fatalf("saving a pick = %d %v", code, body)
	}
	if s := h.app.Snapshot(); s.LoggingCharacter == nil || *s.LoggingCharacter != *pick {
		t.Fatalf("after a pick: %+v", s.LoggingCharacter)
	}
	// Automatic is null on the wire.
	if code, body := h.call(http.MethodPost, base+"/api/settings",
		map[string]any{"visibility": "public", "logging_character": nil}); code != http.StatusOK {
		t.Fatalf("saving Automatic = %d %v", code, body)
	}
	if s := h.app.Snapshot(); s.LoggingCharacter != nil {
		t.Errorf("after Automatic the snapshot still has %+v", s.LoggingCharacter)
	}
	stored, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	if stored.LoggingCharacter != nil {
		t.Errorf("config.json still holds %+v", stored.LoggingCharacter)
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

func TestTheAddressTheCompanionLogsCarriesNoToken(t *testing.T) {
	h := newHarness(t)
	ui, err := h.app.Serve()
	if err != nil {
		t.Fatal(err)
	}
	token := strings.Trim(strings.TrimPrefix(ui.URL, "http://"+ui.Addr), "/")
	if token == "" {
		t.Fatal("the UI URL carries no session token")
	}
	if strings.Contains(ui.Addr, token) {
		t.Errorf("the logged address %q carries the session token", ui.Addr)
	}
	resp, err := http.Get(ui.URL + "api/status")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("the served UI answered %d", resp.StatusCode)
	}
	if err := ui.Close(); err != nil {
		t.Fatalf("closing the UI server: %v", err)
	}
	if _, err := http.Get(ui.URL + "api/status"); err == nil {
		t.Error("the UI server still answers after Close")
	}
}

// TestTheCharactersOnDiskAtLaunchStillReachTheSite holds the addon
// sync to its one job while the status page shares its parse cache.
// The window polls /api/status every two seconds and the tick runs
// every second, so a poll routinely reads a SavedVariables file
// before the sync does; that must not cost the upload.
func TestTheCharactersOnDiskAtLaunchStillReachTheSite(t *testing.T) {
	h := newHarness(t)
	sv := filepath.Join(h.game, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	writeExport := func(name string, at time.Time) {
		t.Helper()
		if err := os.WriteFile(sv, []byte(`ForeverSixtyDB = { ["characters"] = { `+
			`["`+name+`-Hardcore"] = { ["name"] = "`+name+`", ["realm"] = "Hardcore", `+
			`["region"] = "us", ["export"] = "FS1:x" } } }`), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := os.Chtimes(sv, at, at); err != nil {
			t.Fatal(err)
		}
	}
	writeExport("Morrowlyn", t0)
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}

	// The window is open before the first tick: its poll reads the
	// file the sync has not looked at yet.
	if n := len(h.app.Snapshot().Characters); n != 1 {
		t.Fatalf("the status page shows %d characters", n)
	}
	h.app.Step(t.Context(), t0)
	if got := h.srv.Exports(); len(got) != 1 || got[0]["name"] != "Morrowlyn" {
		t.Fatalf("the characters present at launch never reached the site: %+v", got)
	}

	// The player logs out and the game rewrites the file. The status
	// poll gets there first again; the logout must still upload.
	writeExport("Thessaly", t0.Add(time.Hour))
	h.app.Snapshot()
	h.app.Step(t.Context(), t0.Add(time.Hour))
	if got := h.srv.Exports(); len(got) != 2 || got[1]["name"] != "Thessaly" {
		t.Fatalf("a logout's exports never reached the site: %+v", got)
	}

	// A pass with nothing new does not upload the same file again.
	h.app.Step(t.Context(), t0.Add(time.Hour+time.Second))
	if got := h.srv.Exports(); len(got) != 2 {
		t.Fatalf("a quiet pass uploaded again: %+v", got)
	}
}

// TestANewFirstGameFolderWaitsForTheOpenReportAndThenMoves covers the
// refuse-then-retry half of retarget: the save cannot move the tail
// while a report is open, and the step after the report closes does
// it without anyone asking again.
func TestANewFirstGameFolderWaitsForTheOpenReportAndThenMoves(t *testing.T) {
	h := newHarness(t)
	h.app.Step(t.Context(), t0) // prime the tail over the empty folder
	if err := os.WriteFile(h.log, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), t0.Add(time.Second))
	open := h.app.Snapshot().Pipeline
	if !open.Logging || open.LogPath != h.log {
		t.Fatalf("no report is open on the first folder: %+v", open)
	}

	// Mid-raid, the player puts another game folder in front of the
	// one being logged.
	second := newGame(t)
	secondLog := filepath.Join(second, "Logs", "WoWCombatLog.txt")
	if code, body := h.call(http.MethodPost, h.url("/api/settings"), app.Settings{
		Visibility: "public", WoWPaths: []string{second, h.game},
	}); code != http.StatusOK {
		t.Fatalf("settings = %d %v", code, body)
	}
	// The player launches the other install: it is now the one played last.
	played(t, second, time.Now().Add(time.Hour))

	// The tail has not moved: the rest of the night still lands in
	// the open report.
	f, err := os.OpenFile(h.log, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString(fixture.Encounter(1) + fixture.Heartbeat(1)); err != nil {
		t.Fatal(err)
	}
	f.Close()
	h.app.Step(t.Context(), t0.Add(2*time.Second))
	grown := h.app.Snapshot().Pipeline
	if !grown.Logging || grown.LogPath != h.log {
		t.Fatalf("the save moved the tail out from under the open report: %+v", grown)
	}
	if grown.Offset <= open.Offset {
		t.Fatalf("the open report stopped reading: offset %d, was %d", grown.Offset, open.Offset)
	}

	// The night ends and the report completes; the tail moves on the
	// next step with no second save.
	quiet := t0.Add(2*time.Second + watch.CompleteIdle + time.Minute)
	h.app.Step(t.Context(), quiet)
	if s := h.app.Snapshot().Pipeline; s.Logging {
		t.Fatalf("the report did not close: %+v", s)
	}
	h.app.Step(t.Context(), quiet.Add(time.Second)) // the retry that moves the tail
	if err := os.WriteFile(secondLog, []byte(aNight()), 0o600); err != nil {
		t.Fatal(err)
	}
	h.app.Step(t.Context(), quiet.Add(2*time.Second))
	if got := h.app.Snapshot().Pipeline.LogPath; got != secondLog {
		t.Fatalf("the tail is on %q, not the newly added folder's log %q", got, secondLog)
	}
}

// TestTheStatusPageShowsWhenTheSiteSawTheExport is the tray line end to end: the outbox goes
// up on a step, the fake acks it, and the snapshot carries the ack's time and the sim's fate.
func TestTheStatusPageShowsWhenTheSiteSawTheExport(t *testing.T) {
	h := newHarness(t)
	seen := time.Date(2026, 12, 9, 14, 3, 0, 0, time.UTC)
	h.srv.SetClock(func() time.Time { return seen })
	sv := filepath.Join(h.game, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	if err := os.WriteFile(sv, []byte(`ForeverSixtyDB = {
		["characters"] = { ["Obnoxious-Normal"] = { ["name"] = "Obnoxious", ["ruleset"] = "normal",
			["region"] = "us", ["export"] = "FS1:x" } },
		["outbox"] = { ["format"] = 1,
			["export"] = { ["character"] = "Obnoxious-Yell", ["code"] = "FS1:x", ["saved_at"] = "2026-12-09 14:02" },
			["sim_request"] = { ["character"] = "Obnoxious-Yell", ["spec"] = "warrior-arms",
				["requested_at"] = "2026-12-09 14:02", ["gear_code"] = "FS1:x" } } }`), 0o600); err != nil {
		t.Fatal(err)
	}
	if s := h.app.Snapshot(); s.AddonLine != app.CopyNeverSynced || s.AddonSeenAt != nil {
		t.Fatalf("a line before any ack: %+v", s)
	}
	// The state is short; what to do about it is the muted line under it.
	if s := h.app.Snapshot(); !s.AddonNever || s.AddonHelp != app.CopyNeverSyncedHelp || s.AddonStale {
		t.Fatalf("never-synced help: never=%v stale=%v help=%q", s.AddonNever, s.AddonStale, s.AddonHelp)
	}
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	h.app.Step(t.Context(), t0)

	s := h.app.Snapshot()
	if s.AddonSeenAt == nil || !s.AddonSeenAt.Equal(seen) || s.SimRequestState != addon.SimQueued {
		t.Fatalf("snapshot = seen %v, sim %q", s.AddonSeenAt, s.SimRequestState)
	}
	if want := app.CopyLastSynced + " 14:03" + app.CopySeparator + app.CopySimQueued; s.AddonLine != want {
		t.Errorf("line = %q, want %q", s.AddonLine, want)
	}
	_, body := h.call(http.MethodGet, h.url("/api/status"), nil)
	if data := body["data"].(map[string]any); data["addon_line"] != s.AddonLine {
		t.Errorf("the status JSON carries %v", data["addon_line"])
	}
}

func TestForgetRevokesTheDeviceOnTheSiteThenClearsTheToken(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	h.call(http.MethodPost, base+"/api/pair", map[string]string{"code": fakeapi.PairCode})
	code, body := h.call(http.MethodPost, base+"/api/unpair", map[string]string{})
	data, _ := body["data"].(map[string]any)
	if code != http.StatusOK || data["paired"] != false || data["revoked"] != true {
		t.Fatalf("unpair = %d %v", code, body)
	}
	if n := h.srv.SelfRevokes(); n != 1 {
		t.Fatalf("the site was asked to revoke %d times, want 1", n)
	}
	stored, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		t.Fatal(err)
	}
	if stored.DeviceToken != "" {
		t.Fatalf("the token is still stored: %q", stored.DeviceToken)
	}
}

func TestForgetClearsLocallyAndSaysSoWhenTheSiteCannotBeReached(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	h.call(http.MethodPost, base+"/api/pair", map[string]string{"code": fakeapi.PairCode})
	h.srv.FailNext("DELETE /v1/devices/self", 1)
	code, body := h.call(http.MethodPost, base+"/api/unpair", map[string]string{})
	data, _ := body["data"].(map[string]any)
	if code != http.StatusOK || data["paired"] != false || data["revoked"] != false {
		t.Fatalf("unpair = %d %v", code, body)
	}
	if n := h.srv.SelfRevokes(); n != 0 {
		t.Fatalf("a failed revoke was counted: %d", n)
	}
}

func TestForgetCountsADeadTokenAsRevoked(t *testing.T) {
	h := newHarness(t)
	base := h.url("")
	h.call(http.MethodPost, base+"/api/pair", map[string]string{"code": fakeapi.PairCode})
	// The site already revoked it (from Account > Devices): the token is 401.
	h.srv.RevokeToken()
	code, body := h.call(http.MethodPost, base+"/api/unpair", map[string]string{})
	data, _ := body["data"].(map[string]any)
	if code != http.StatusOK || data["paired"] != false || data["revoked"] != true {
		t.Fatalf("unpair = %d %v", code, body)
	}
}

// A state file listed twice under two names (the copy of a save in flight) is
// still one report, and so one row in Reports.
func TestAReportListedTwiceIsOneRow(t *testing.T) {
	h := newHarnessSeeded(t, func(dirs paths.Dirs) {
		r := state.Report{Key: "20261201-200000-bbbbbbbb", StartedAt: t0.Add(-time.Hour), Fights: 3}
		if err := state.Save(dirs.State, r); err != nil {
			t.Fatal(err)
		}
		// The same report under a name that sorts after it, with the server's id.
		b := []byte(`{"key":"20261201-200000-bbbbbbbb","report_id":"rpt000000001","started_at":"` +
			r.StartedAt.Format(time.RFC3339) + `","fights":3}`)
		if err := os.WriteFile(filepath.Join(dirs.State, "20261201-200000-bbbbbbbb-copy.json"), b, 0o600); err != nil {
			t.Fatal(err)
		}
	})
	rows := h.app.Snapshot().Reports
	if len(rows) != 1 {
		t.Fatalf("Reports has %d rows for one report: %+v", len(rows), rows)
	}
	if rows[0].ReportID != "rpt000000001" {
		t.Errorf("the row kept the older copy: %+v", rows[0])
	}
}

func TestTheInboxGoesToTheForeverInstallAndNeverToAnotherGame(t *testing.T) {
	retail := newInstall(t, "_retail_")
	anni := newInstall(t, "_anniversary_")
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Config.WoWPaths = append(append([]string{}, o.Config.WoWPaths...), retail, anni)
	})
	if code, _ := h.call(http.MethodPost, h.url("/api/pair"),
		map[string]string{"code": fakeapi.PairCode}); code != http.StatusOK {
		t.Fatal("pairing failed")
	}
	h.app.Step(t.Context(), t0)
	h.app.Step(t.Context(), t0.Add(5*time.Minute))
	sv := func(game string) string {
		return filepath.Join(game, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	}
	if _, err := os.Stat(addon.InboxPath(sv(h.game))); err != nil {
		t.Errorf("the Classic beta got no inbox: %v", err)
	}
	for _, other := range []string{retail, anni} {
		if _, err := os.Stat(addon.InboxPath(sv(other))); err == nil {
			t.Errorf("an inbox was written under %s", other)
		}
	}
	if n := len(h.app.Snapshot().Installs); n != 1 {
		t.Errorf("%d installs listed, want the Classic beta alone", n)
	}
}
