package app_test

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/picker"
)

// newInstall makes a second flavour folder the picker can return.
func newInstall(t *testing.T, flavor string) string {
	t.Helper()
	dir := filepath.Join(t.TempDir(), flavor)
	for _, p := range []string{filepath.Join(dir, "Logs"), filepath.Join(dir, "WTF", "Account", "ACCOUNT#1", "SavedVariables")} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func (h *harness) storedPaths() []string {
	h.t.Helper()
	c, err := config.Load(h.dirs.ConfigFile())
	if err != nil {
		h.t.Fatal(err)
	}
	return c.WoWPaths
}

func errorCode(body map[string]any) string {
	e, _ := body["error"].(map[string]any)
	c, _ := e["code"].(string)
	return c
}

func TestAPickedFolderIsValidatedAddedAndStoredByItsOwnPath(t *testing.T) {
	for flavor, name := range map[string]string{"_forever_": "WoW Forever"} {
		t.Run(flavor, func(t *testing.T) {
			h := newHarness(t)
			picked := newInstall(t, flavor)
			h.pick = picked
			code, body := h.call(http.MethodPost, h.url("/api/folders/pick"), nil)
			if code != http.StatusOK || body["ok"] != true {
				t.Fatalf("pick = %d %v", code, body)
			}
			if got := h.storedPaths(); len(got) != 2 || got[1] != picked {
				t.Fatalf("wow_paths = %v", got)
			}
			var added *app.InstallStatus
			for _, in := range h.app.Snapshot().Installs {
				if in.Path == picked {
					added = &in
				}
			}
			if added == nil || !added.Added || added.FlavorName != name || added.Watching {
				t.Fatalf("the new install = %+v", added)
			}
		})
	}
}

func TestAFolderAboveTheInstallsAddsTheForeverOnesOnly(t *testing.T) {
	h := newHarness(t)
	root := t.TempDir()
	// The ended beta is another game: picking the folder above it adds the live
	// game alone, so beta characters never reach the site.
	for _, f := range []string{"_classic_beta_", "_forever_", "_retail_", "_anniversary_"} {
		for _, p := range []string{filepath.Join(root, f, "WTF"), filepath.Join(root, f, "Logs")} {
			if err := os.MkdirAll(p, 0o755); err != nil {
				t.Fatal(err)
			}
		}
	}
	h.pick = root
	if code, _ := h.call(http.MethodPost, h.url("/api/folders/pick"), nil); code != http.StatusOK {
		t.Fatal("pick failed")
	}
	got := h.storedPaths()
	if len(got) != 2 || got[1] != filepath.Join(root, "_forever_") {
		t.Fatalf("wow_paths = %v, want the game folder, then _forever_ by its own path", got)
	}
	for _, in := range h.app.Snapshot().Installs {
		if in.Flavor != "_forever_" {
			t.Errorf("an install of %s is listed", in.Flavor)
		}
	}
}

func TestARetailFolderIsRefusedInWordsThatSayWhy(t *testing.T) {
	h := newHarness(t)
	for pick, want := range map[string]string{
		newInstall(t, "_retail_"):       "That is Retail. Forever Sixty runs on WoW Forever; pick its _forever_ folder.",
		newInstall(t, "_classic_beta_"): "That is Classic beta. Forever Sixty runs on WoW Forever; pick its _forever_ folder.",
		newInstall(t, "_private_srv_"):  "Forever Sixty runs on WoW Forever; pick its _forever_ folder.",
	} {
		h.pick = pick
		code, body := h.call(http.MethodPost, h.url("/api/folders/pick"), nil)
		if code != http.StatusBadRequest || errorCode(body) != "not_wow_folder" {
			t.Fatalf("pick = %d %v", code, body)
		}
		if msg := body["error"].(map[string]any)["message"]; msg != want {
			t.Errorf("message = %q, want %q", msg, want)
		}
	}
	if got := h.storedPaths(); len(got) != 1 {
		t.Fatalf("wow_paths = %v, a refused folder was stored", got)
	}
}

func TestAStoredFolderOfAnotherGameIsDroppedAtLoad(t *testing.T) {
	retail := newInstall(t, "_retail_")
	h := newHarnessWith(t, nil, func(o *app.Options) {
		o.Config.WoWPaths = append(append([]string{}, o.Config.WoWPaths...), retail)
	})
	for _, in := range h.app.Snapshot().Installs {
		if in.Path == retail {
			t.Fatalf("the remembered Retail folder is still listed: %+v", in)
		}
	}
}

func TestAFolderThatIsNotTheGameIsRefusedWithTheDesignersSentence(t *testing.T) {
	h := newHarness(t)
	h.pick = t.TempDir()
	code, body := h.call(http.MethodPost, h.url("/api/folders/pick"), nil)
	if code != http.StatusBadRequest || errorCode(body) != "not_wow_folder" {
		t.Fatalf("pick = %d %v", code, body)
	}
	if msg := body["error"].(map[string]any)["message"]; msg !=
		"That folder is not a World of Warcraft folder. Forever Sixty runs on WoW Forever; pick its _forever_ folder." {
		t.Errorf("message = %q", msg)
	}
	if got := h.storedPaths(); len(got) != 1 {
		t.Fatalf("wow_paths = %v, a refused folder was stored", got)
	}
}

func TestAFolderAlreadyInTheListIsSaidToBe(t *testing.T) {
	h := newHarness(t)
	h.pick = h.game
	code, body := h.call(http.MethodPost, h.url("/api/folders/pick"), nil)
	if code != http.StatusConflict || errorCode(body) != "known_folder" {
		t.Fatalf("pick = %d %v", code, body)
	}
}

func TestACancelledDialogChangesNothing(t *testing.T) {
	h := newHarness(t)
	h.pickErr = picker.ErrCancelled
	code, body := h.call(http.MethodPost, h.url("/api/folders/pick"), nil)
	if code != http.StatusOK || body["ok"] != true {
		t.Fatalf("cancelled pick = %d %v", code, body)
	}
	h.pickErr = picker.ErrUnavailable
	code, body = h.call(http.MethodPost, h.url("/api/folders/pick"), nil)
	if code != http.StatusServiceUnavailable || errorCode(body) != "no_picker" {
		t.Fatalf("no dialog = %d %v", code, body)
	}
	if got := h.storedPaths(); len(got) != 1 {
		t.Fatalf("wow_paths = %v", got)
	}
}

func TestOnlyAFolderThePlayerAddedCanBeRemoved(t *testing.T) {
	h := newHarness(t)
	added := newInstall(t, "_forever_")
	h.pick = added
	h.call(http.MethodPost, h.url("/api/folders/pick"), nil)

	if code, body := h.call(http.MethodPost, h.url("/api/folders/remove"), map[string]string{"path": added}); code != http.StatusOK || body["ok"] != true {
		t.Fatalf("remove = %d %v", code, body)
	}
	if got := h.storedPaths(); len(got) != 1 || got[0] != h.game {
		t.Fatalf("wow_paths = %v", got)
	}
	code, body := h.call(http.MethodPost, h.url("/api/folders/remove"), map[string]string{"path": "/not/a/folder"})
	if code != http.StatusConflict || errorCode(body) != "not_added" {
		t.Fatalf("remove of a stranger = %d %v", code, body)
	}
	if code, _ := h.call(http.MethodPost, h.url("/api/folders/remove"), "nope"); code != http.StatusBadRequest {
		t.Errorf("rubbish = %d", code)
	}
}

// The settings page no longer posts folders. A save without them must not wipe what was added.
func TestSavingSettingsWithoutFoldersKeepsTheOnesTheyAdded(t *testing.T) {
	h := newHarness(t)
	code, _ := h.call(http.MethodPost, h.url("/api/settings"), map[string]any{"visibility": "unlisted"})
	if code != http.StatusOK {
		t.Fatalf("save = %d", code)
	}
	if got := h.storedPaths(); len(got) != 1 || got[0] != h.game {
		t.Fatalf("wow_paths = %v", got)
	}
	code, _ = h.call(http.MethodPost, h.url("/api/settings"), map[string]any{"visibility": "unlisted", "wow_paths": []string{}})
	if code != http.StatusOK {
		t.Fatalf("save with an empty list = %d", code)
	}
	if got := h.storedPaths(); len(got) != 0 {
		t.Fatalf("an explicit empty list was ignored: %v", got)
	}
}

func TestTheSnapshotCarriesTheNewFieldsAndNeverAnErrorsText(t *testing.T) {
	h := newHarness(t)
	s := h.app.Snapshot()
	if s.AppVersion == "" || s.TokenBackendName != "A private file in your user folder" || s.TokenBackendNote == "" {
		t.Errorf("version %q, backend %q", s.AppVersion, s.TokenBackendName)
	}
	if len(s.Installs) != 1 {
		t.Fatalf("installs = %+v", s.Installs)
	}
	in := s.Installs[0]
	if in.FlavorName != "WoW Forever" || !in.Watching || !in.Added || in.Path != h.game {
		t.Errorf("install = %+v", in)
	}
	if in.AdvancedLogHelp != app.CopyAdvancedRowHelp {
		t.Errorf("help = %q", in.AdvancedLogHelp)
	}
	if s.Connect.State != "idle" {
		t.Errorf("connect = %+v", s.Connect)
	}
}

func TestOpenGoesToTheBrowserForHTTPSAndRefusesTheRest(t *testing.T) {
	h := newHarness(t)
	code, _ := h.call(http.MethodPost, h.url("/api/open"), map[string]string{"url": "https://foreversixty.gg/logs"})
	if code != http.StatusOK || len(h.openedLinks()) != 1 {
		t.Fatalf("open = %d, links %v", code, h.openedLinks())
	}
	for _, bad := range []string{"file:///C:/Windows/System32/cmd.exe", "javascript:alert(1)", "http://example.com/", "", "ftp://x/"} {
		code, body := h.call(http.MethodPost, h.url("/api/open"), map[string]string{"url": bad})
		if code != http.StatusBadRequest || errorCode(body) != "bad_url" {
			t.Errorf("open %q = %d %v", bad, code, body)
		}
	}
	if len(h.openedLinks()) != 1 {
		t.Fatalf("a refused link reached the browser: %v", h.openedLinks())
	}
}

func TestTheHeadlineTakesTheFirstMatch(t *testing.T) {
	h := newHarness(t)
	if k := h.app.Snapshot().Headline; k.Key != "not_connected" || k.Title != "Not connected" ||
		k.Action == nil || k.Action.Page != "device" || k.Tone != "bad" {
		t.Fatalf("unpaired = %+v", k)
	}
	h.call(http.MethodPost, h.url("/api/pair"), map[string]string{"code": "HORDE-42"})
	if k := h.app.Snapshot().Headline; k.Key != "ready" || k.Title != "Ready" || k.Action != nil {
		t.Fatalf("paired = %+v", k)
	}
	// Advanced logging off outranks ready.
	if err := os.WriteFile(filepath.Join(h.game, "WTF", "Config.wtf"), []byte("SET advancedCombatLogging \"0\"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if k := h.app.Snapshot().Headline; k.Key != "advanced_off" || k.Tone != "warn" || k.Title != "Advanced logging is off" ||
		k.Sub != "Reports will miss deaths with health, item level, resources and positions. "+
			"Turn it on in Options > System > Network > Advanced Combat Logging, then relog." {
		t.Fatalf("advanced off = %+v", k)
	}
	if err := os.WriteFile(filepath.Join(h.game, "WTF", "Config.wtf"), []byte("SET advancedCombatLogging \"1\"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	// No folder outranks everything below it.
	h.call(http.MethodPost, h.url("/api/settings"), map[string]any{"visibility": "public", "wow_paths": []string{}})
	if k := h.app.Snapshot().Headline; k.Key != "no_folder" || k.Action == nil || k.Action.Page != "settings" {
		t.Fatalf("no folder = %+v", k)
	}
}

func TestTheTrayAsksForTheSameHeadlineTheWindowShows(t *testing.T) {
	h := newHarness(t)
	if got, want := h.app.Headline(), h.app.Snapshot().Headline; got.Key != want.Key || got.Title != want.Title {
		t.Fatalf("tray headline %+v differs from the window's %+v", got, want)
	}
	if got := h.app.ReportsURL(); got != "https://foreversixty.gg/logs" {
		t.Errorf("reports url = %q", got)
	}
}

func TestTheFirstCloseIsTheOnlyOneThatIsTold(t *testing.T) {
	h := newHarness(t)
	if !h.app.FirstClose() {
		t.Fatal("the first close was not the first")
	}
	if h.app.FirstClose() {
		t.Fatal("the second close was told again")
	}
	if _, err := os.Stat(filepath.Join(h.dirs.Home, "tray-notice-seen")); err != nil {
		t.Fatalf("the marker is not on disk: %v", err)
	}
}
