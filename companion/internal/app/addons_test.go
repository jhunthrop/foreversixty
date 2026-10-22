package app_test

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	minisign "github.com/jedisct1/go-minisign"

	"github.com/jhunthrop/foreversixty/companion/internal/addonupdater"
	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

// signer mints a minisign key, as the release workflows hold one.
type signer struct {
	sk  minisign.PrivateKey
	pub string
}

func newSigner(t *testing.T) signer {
	t.Helper()
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	var sk minisign.PrivateKey
	sk.SignatureAlgorithm = [2]byte{'E', 'd'}
	copy(sk.SecretKey[:], priv)
	pk := sk.PublicKey()
	raw := append(append(append([]byte{}, pk.SignatureAlgorithm[:]...), pk.KeyId[:]...), pk.PublicKey[:]...)
	return signer{sk: sk, pub: base64.StdEncoding.EncodeToString(raw)}
}

func (s signer) release(t *testing.T, tag, zipName string, files map[string]string) fakeapi.GitHubRelease {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for n, body := range files {
		w, err := zw.Create(n)
		if err != nil {
			t.Fatal(err)
		}
		w.Write([]byte(body))
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	sig, err := s.sk.Sign(buf.Bytes(), minisign.SignOptions{Hashed: true})
	if err != nil {
		t.Fatal(err)
	}
	return fakeapi.GitHubRelease{Tag: tag, Assets: []fakeapi.GitHubAsset{
		{Name: zipName, Body: buf.Bytes()}, {Name: zipName + ".minisig", Body: sig.Encode()},
	}}
}

func (s signer) addon(t *testing.T, version string) fakeapi.GitHubRelease {
	return s.release(t, "addon-v"+version, "ForeverSixty-addon-v"+version+"-forever.zip", map[string]string{
		"ForeverSixty/ForeverSixty.toc":           "## Version: " + version + "\n",
		"ForeverSixty/Core.lua":                   "-- " + version + "\n",
		"ForeverSixtyInbox/ForeverSixtyInbox.toc": "## Version: addon-v" + version + "\n",
	})
}

func (s signer) data(t *testing.T, date string) fakeapi.GitHubRelease {
	return s.release(t, "data-v"+date, "ForeverSixtyData-data-v"+date+"-forever.zip", map[string]string{
		"ForeverSixtyData/ForeverSixtyData.toc": "## Version: data-v" + date + "\n",
		"ForeverSixtyData/Data.lua":             "-- " + date + "\n",
	})
}

// addonWorld is the harness with the addon updater keyed, and a game that
// can be running or closed.
type addonWorld struct {
	*harness
	signer  signer
	mu      sync.Mutex
	running bool
	tick    int
	logs    lockedBuf
}

func newAddonWorld(t *testing.T) *addonWorld {
	t.Helper()
	w := &addonWorld{signer: newSigner(t)}
	w.harness = newHarnessWith(t, nil, func(o *app.Options) {
		o.Log = slog.New(slog.NewTextHandler(&w.logs, &slog.HandlerOptions{Level: slog.LevelDebug}))
		o.Updater = updater.Options{PublicKey: w.signer.pub, APIBase: o.Config.APIBaseURL}
		o.Addons = addonupdater.Options{
			Now: time.Now,
			Processes: func() ([]addonupdater.Proc, error) {
				w.mu.Lock()
				defer w.mu.Unlock()
				if !w.running {
					return nil, nil
				}
				return []addonupdater.Proc{{Name: "WowClassic.exe", Path: filepath.Join(w.harness.game, "WowClassic.exe")}}, nil
			},
		}
	})
	return w
}

func (w *addonWorld) setRunning(v bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.running = v
}

func (w *addonWorld) addons() string { return filepath.Join(w.game, "Interface", "AddOns") }

func (w *addonWorld) put(folder, toc string) {
	w.t.Helper()
	dir := filepath.Join(w.addons(), folder)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		w.t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, folder+".toc"), []byte(toc), 0o644); err != nil {
		w.t.Fatal(err)
	}
}

// lines are the Addon row's version lines as the window gets them.
func (w *addonWorld) lines() []string {
	var out []string
	for _, l := range w.app.Snapshot().AddonUpdate.Lines {
		out = append(out, l.Text)
	}
	return out
}

// step runs the companion's step, a minute later each time, until the Addon
// row says want, since the updater works on its own goroutine.
func (w *addonWorld) stepUntil(want ...string) {
	w.t.Helper()
	for i := 0; i < 300; i++ {
		w.tick++
		w.app.Step(context.Background(), t0.Add(time.Duration(w.tick)*2*time.Minute))
		if strings.Join(w.lines(), "|") == strings.Join(want, "|") {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	w.t.Fatalf("the Addon row never read %q; it reads %q", want, w.lines())
}

func TestTheAddonRowSaysWhatIsInstalled(t *testing.T) {
	w := newAddonWorld(t)
	w.srv.SetReleases() // nothing newer exists
	w.put("ForeverSixty", "## Version: 2.0.15\n")
	w.put("ForeverSixtyData", "## Version: data-v2026.10.21\n")
	got := w.lines()
	if strings.Join(got, "|") != "Addon 2.0.15"+app.CopySeparator+app.CopyUpToDate+"|Data 2026.10.21"+app.CopySeparator+app.CopyUpToDate {
		t.Errorf("lines = %q", got)
	}
	if got[0] != "Addon 2.0.15 · up to date" || got[1] != "Data 2026.10.21 · up to date" {
		t.Errorf("the words moved: %q", got)
	}
	if offers := w.app.Snapshot().AddonUpdate.Offers; len(offers) != 0 {
		t.Errorf("offers with the addon installed = %+v", offers)
	}
}

func TestTheAddonRowOffersTheInstallWhereItIsMissing(t *testing.T) {
	w := newAddonWorld(t)
	w.srv.SetReleases(w.signer.data(t, "2026.10.21"), w.signer.addon(t, "2.0.16"))
	v := w.app.Snapshot().AddonUpdate
	if len(v.Lines) != 1 || v.Lines[0].Text != "Addon not found in Classic beta" {
		t.Fatalf("lines = %+v", v.Lines)
	}
	if len(v.Offers) != 1 || v.Offers[0].Label != "Install the addon into Classic beta" || v.Offers[0].Path != w.game {
		t.Fatalf("offers = %+v", v.Offers)
	}

	// Automatic updates never install where the addon is missing.
	for i := 0; i < 3; i++ {
		w.tick++
		w.app.Step(context.Background(), t0.Add(time.Duration(w.tick)*8*time.Hour))
		time.Sleep(20 * time.Millisecond)
	}
	if _, err := os.Stat(filepath.Join(w.addons(), "ForeverSixty")); err == nil {
		t.Fatal("the addon was installed with no button pressed")
	}

	// The game is running: the button says so, in a sentence, and writes nothing.
	w.setRunning(true)
	code, body := w.call(http.MethodPost, w.url("/api/addon/install"), map[string]string{"path": w.game})
	if code != http.StatusConflict || errCode(body) != app.CodeGameRunning {
		t.Fatalf("install while the game runs = %d %v", code, body)
	}
	if msg := errMessage(body); msg != app.Message(app.CodeGameRunning) || !strings.Contains(msg, "Close World of Warcraft") {
		t.Errorf("message = %q", msg)
	}
	if _, err := os.Stat(filepath.Join(w.addons(), "ForeverSixty")); err == nil {
		t.Fatal("the button wrote into the game's folder while the game ran")
	}

	// Closed: it installs both, and the row reads installed.
	w.setRunning(false)
	code, body = w.call(http.MethodPost, w.url("/api/addon/install"), map[string]string{"path": w.game})
	if code != http.StatusOK {
		t.Fatalf("install = %d %v", code, body)
	}
	want := []string{"Addon 2.0.16 installed · takes effect at your next login", "Data 2026.10.21 installed · takes effect at your next login"}
	if got := w.lines(); strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("after the button: %q", got)
	}
	if offers := w.app.Snapshot().AddonUpdate.Offers; len(offers) != 0 {
		t.Errorf("offers after the install = %+v", offers)
	}
	// A login is seen in the log: back to the first form.
	time.Sleep(20 * time.Millisecond)
	if err := os.WriteFile(filepath.Join(w.game, "Logs", "Client.log"), []byte("login\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := w.lines(); got[0] != "Addon 2.0.16 · up to date" || got[1] != "Data 2026.10.21 · up to date" {
		t.Errorf("after a login: %q", got)
	}

	// An unknown path is not a folder to write into.
	code, body = w.call(http.MethodPost, w.url("/api/addon/install"), map[string]string{"path": t.TempDir()})
	if code != http.StatusBadRequest {
		t.Errorf("install into an unknown folder = %d %v", code, body)
	}
}

func TestAnUpdateWaitsForTheGameThenSaysSo(t *testing.T) {
	w := newAddonWorld(t)
	w.put("ForeverSixty", "## Version: 2.0.15\n")
	w.put("ForeverSixtyData", "## Version: data-v2026.10.20\n")
	w.srv.SetReleases(w.signer.data(t, "2026.10.21"), w.signer.addon(t, "2.0.16"))

	w.setRunning(true)
	w.stepUntil("Addon 2.0.16 ready · installs when the game closes", "Data 2026.10.21 ready · installs when the game closes")
	if v, _ := addonupdater.ReadVersion(w.addons(), "ForeverSixty"); v != "2.0.15" {
		t.Fatalf("the addon changed under a running game: %s", v)
	}

	w.setRunning(false)
	w.stepUntil("Addon 2.0.16 installed · takes effect at your next login", "Data 2026.10.21 installed · takes effect at your next login")
	if v, _ := addonupdater.ReadVersion(w.addons(), "ForeverSixty"); v != "2.0.16" {
		t.Fatalf("addon = %s", v)
	}
}

func TestTheSwitchTurnsTheChecksAndInstallsOff(t *testing.T) {
	w := newAddonWorld(t)
	if !w.app.Snapshot().KeepAddonUpdated {
		t.Fatal("the switch is not on by default")
	}
	w.put("ForeverSixty", "## Version: 2.0.15\n")
	w.srv.SetReleases(w.signer.addon(t, "2.0.16"))

	code, body := w.call(http.MethodPost, w.url("/api/settings"), map[string]any{"visibility": "public", "keep_addon_updated": false})
	if code != http.StatusOK {
		t.Fatalf("settings = %d %v", code, body)
	}
	if w.app.Snapshot().KeepAddonUpdated || w.app.Snapshot().AddonUpdate.Enabled {
		t.Error("the switch did not turn off")
	}
	onDisk, err := config.Load(w.dirs.ConfigFile())
	if err != nil || onDisk.AddonUpdates() {
		t.Errorf("config.json = %+v, %v", onDisk, err)
	}
	for i := 0; i < 4; i++ {
		w.tick++
		w.app.Step(context.Background(), t0.Add(time.Duration(w.tick)*8*time.Hour))
		time.Sleep(20 * time.Millisecond)
	}
	if reads := w.srv.ReleaseListReads(); reads != 0 {
		t.Errorf("the release list was read %d times with the switch off", reads)
	}
	if v, _ := addonupdater.ReadVersion(w.addons(), "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed with the switch off: %s", v)
	}

	// A save that does not mention it leaves it alone; turning it on again takes effect.
	w.call(http.MethodPost, w.url("/api/settings"), map[string]any{"visibility": "private"})
	if w.app.Snapshot().KeepAddonUpdated {
		t.Error("a save without the field turned the switch back on")
	}
	w.call(http.MethodPost, w.url("/api/settings"), map[string]any{"visibility": "private", "keep_addon_updated": true})
	w.stepUntil("Addon 2.0.16 installed · takes effect at your next login")
}

func TestABuildWithoutAKeyOffersNoInstall(t *testing.T) {
	h := newHarness(t) // the test binary has no compiled-in key
	v := h.app.Snapshot().AddonUpdate
	if len(v.Offers) != 0 {
		t.Errorf("offers = %+v; a build that cannot verify a zip must not offer to install one", v.Offers)
	}
	if len(v.Lines) != 1 || v.Lines[0].Text != "Addon not found in Classic beta" {
		t.Errorf("lines = %+v", v.Lines)
	}
}

func errCode(body map[string]any) string {
	e, _ := body["error"].(map[string]any)
	s, _ := e["code"].(string)
	return s
}

func errMessage(body map[string]any) string {
	e, _ := body["error"].(map[string]any)
	s, _ := e["message"].(string)
	return s
}

func TestNoOfferIsMadeForAnotherGamesInstall(t *testing.T) {
	retail := newInstall(t, "_retail_")
	anni := newInstall(t, "_anniversary_")
	w := &addonWorld{signer: newSigner(t)}
	w.harness = newHarnessWith(t, nil, func(o *app.Options) {
		o.Updater = updater.Options{PublicKey: w.signer.pub, APIBase: o.Config.APIBaseURL}
		o.Config.WoWPaths = append(append([]string{}, o.Config.WoWPaths...), retail, anni)
	})
	w.srv.SetReleases()
	w.put("ForeverSixty", "## Version: 2.0.15\n")
	w.put("ForeverSixtyData", "## Version: data-v2026.10.21\n")
	v := w.app.Snapshot().AddonUpdate
	for _, l := range v.Lines {
		if strings.Contains(l.Text, "Retail") || strings.Contains(l.Text, "Anniversary") || strings.Contains(l.Text, "not found") {
			t.Errorf("a line about another game: %q", l.Text)
		}
	}
	if len(v.Offers) != 0 {
		t.Errorf("offers = %+v", v.Offers)
	}
	if len(v.Lines) != 2 {
		t.Errorf("lines = %+v, want the Classic beta's two", v.Lines)
	}
}

func TestTheAddonRowHelpSaysNoCharacterYetWhenTheAddonIsInstalled(t *testing.T) {
	w := newAddonWorld(t)
	w.srv.SetReleases()

	// No install has the addon: the step is to install it.
	s := w.app.Snapshot()
	if !s.AddonNever || s.AddonHelp != app.CopyNeverSyncedHelp {
		t.Errorf("addon not installed: never %v, help %q", s.AddonNever, s.AddonHelp)
	}

	// The addon is installed and no export has arrived: nothing is missing but the first logout.
	w.put("ForeverSixty", "## Version: 2.0.15\n")
	s = w.app.Snapshot()
	want := "No character yet · log out once in game and the addon's save arrives here."
	if !s.AddonNever || s.AddonHelp != want {
		t.Errorf("addon installed, no export: never %v, help %q", s.AddonNever, s.AddonHelp)
	}
	if got := w.lines(); len(got) != 1 || got[0] != "Addon 2.0.15 · up to date" {
		t.Errorf("lines = %q", got)
	}
}

func TestAWaitingUpdateIsLoggedOnceNotEveryMinute(t *testing.T) {
	w := newAddonWorld(t)
	w.put("ForeverSixty", "## Version: 2.0.15\n")
	w.put("ForeverSixtyData", "## Version: data-v2026.10.20\n")
	w.srv.SetReleases(w.signer.data(t, "2026.10.21"), w.signer.addon(t, "2.0.16"))

	w.setRunning(true)
	w.stepUntil("Addon 2.0.16 ready · installs when the game closes", "Data 2026.10.21 ready · installs when the game closes")
	// Many more minutes with the game still up.
	for i := 0; i < 20; i++ {
		w.tick++
		w.app.Step(context.Background(), t0.Add(time.Duration(w.tick)*2*time.Minute))
		time.Sleep(10 * time.Millisecond)
	}
	// One line per family (the addon and the data are staged apart), not one a minute.
	if n := w.logs.count("an addon update is waiting for the game to close"); n != 2 {
		t.Fatalf("the wait was logged %d times over 20 more minutes; want once per family (2)", n)
	}

	w.setRunning(false)
	w.stepUntil("Addon 2.0.16 installed · takes effect at your next login", "Data 2026.10.21 installed · takes effect at your next login")
	if n := w.logs.count("installed the addon 2.0.16"); n != 1 {
		t.Errorf("the install was logged %d times", n)
	}
}
