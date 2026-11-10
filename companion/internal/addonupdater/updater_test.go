package addonupdater

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	minisign "github.com/jedisct1/go-minisign"

	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

var t0 = time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)

// ---- fixtures ---------------------------------------------------------

func keypair(t *testing.T) (minisign.PrivateKey, string) {
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
	return sk, base64.StdEncoding.EncodeToString(raw)
}

func sign(t *testing.T, sk minisign.PrivateKey, body []byte) []byte {
	t.Helper()
	sig, err := sk.Sign(body, minisign.SignOptions{Hashed: true})
	if err != nil {
		t.Fatal(err)
	}
	return sig.Encode()
}

func mkZip(t *testing.T, files map[string]string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	names := make([]string, 0, len(files))
	for n := range files {
		names = append(names, n)
	}
	slices.Sort(names)
	for _, n := range names {
		w, err := zw.Create(n)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write([]byte(files[n])); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

// addonZip is what addon-release.yml ships: ForeverSixty and ForeverSixtyInbox at the top.
func addonZip(t *testing.T, version string) []byte {
	return mkZip(t, map[string]string{
		"ForeverSixty/ForeverSixty.toc":           "## Interface: 16001\n## Version: " + version + "\n\nCore.lua\n",
		"ForeverSixty/Core.lua":                   "-- core " + version + "\n",
		"ForeverSixtyInbox/ForeverSixtyInbox.toc": "## Version: addon-v" + version + "\n\nForeverSixtyInbox.lua\n",
		"ForeverSixtyInbox/ForeverSixtyInbox.lua": "-- inbox\n",
	})
}

// dataZip is what addon-data-release.yml ships; the packager stamps the tag into the TOC.
func dataZip(t *testing.T, date string) []byte {
	return mkZip(t, map[string]string{
		"ForeverSixtyData/ForeverSixtyData.toc": "## Version: data-v" + date + "\n\nData.lua\n",
		"ForeverSixtyData/Data.lua":             "-- data " + date + "\n",
	})
}

// release is one release in the fake feed.
type release struct {
	tag     string
	zipName string
	zip     []byte
	sig     []byte // nil: the manifest names no .minisig
	// zipURL and zipSHA, when set, replace what the manifest says for the zip.
	zipURL, zipSHA string
}

// feed serves one manifest per family, naming the newest release published
// of it, and the downloads behind them. A family with nothing published
// answers 404, as the real feed does before its first release.
type feed struct {
	*httptest.Server
	mu       sync.Mutex
	releases []release
	reads    int
	// status, when set, is what every manifest read answers.
	status int
}

func newFeed(t *testing.T) *feed {
	t.Helper()
	f := &feed{}
	mux := http.NewServeMux()
	f.Server = httptest.NewServer(mux)
	t.Cleanup(f.Close)
	mux.HandleFunc("GET /dl/{family}/latest.json", func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		defer f.mu.Unlock()
		f.reads++
		if f.status != 0 {
			http.Error(w, "feed down", f.status)
			return
		}
		family := r.PathValue("family")
		i, ok := f.newest(family)
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(f.manifest(family, i))
	})
	mux.HandleFunc("GET /dl/{family}/{i}/{kind}", func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		defer f.mu.Unlock()
		var i int
		if _, err := fmt.Sscan(r.PathValue("i"), &i); err != nil || i < 0 || i >= len(f.releases) {
			http.NotFound(w, r)
			return
		}
		if r.PathValue("kind") == "zip" {
			w.Write(f.releases[i].zip)
		} else {
			w.Write(f.releases[i].sig)
		}
	})
	return f
}

// base is what an updater is pointed at.
func (f *feed) base() string { return f.URL + "/dl" }

// newest is the index of the newest release of family.
func (f *feed) newest(family string) (int, bool) {
	for i, rel := range f.releases {
		if strings.HasPrefix(rel.tag, family+"-v") {
			return i, true
		}
	}
	return 0, false
}

func (f *feed) manifest(family string, i int) updater.Release {
	rel := f.releases[i]
	at := fmt.Sprintf("%s/%s/%d/", f.base(), family, i)
	zip := updater.Asset{Name: rel.zipName, URL: at + "zip", SHA256: digest(rel.zip)}
	if rel.zipURL != "" {
		zip.URL = rel.zipURL
	}
	if rel.zipSHA != "" {
		zip.SHA256 = rel.zipSHA
	}
	out := updater.Release{Schema: updater.ManifestSchema, Family: family, TagName: rel.tag, Assets: []updater.Asset{zip}}
	if rel.sig != nil {
		out.Assets = append(out.Assets, updater.Asset{Name: rel.zipName + ".minisig", URL: at + "sig", SHA256: digest(rel.sig)})
	}
	return out
}

func digest(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func (f *feed) publish(rel ...release) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.releases = append(rel, f.releases...) // newest first
}

func signedAddon(t *testing.T, sk minisign.PrivateKey, version string) release {
	z := addonZip(t, version)
	return release{tag: "addon-v" + version, zipName: "ForeverSixty-addon-v" + version + "-forever.zip", zip: z, sig: sign(t, sk, z)}
}

func signedData(t *testing.T, sk minisign.PrivateKey, date string) release {
	z := dataZip(t, date)
	return release{tag: "data-v" + date, zipName: "ForeverSixtyData-data-v" + date + "-forever.zip", zip: z, sig: sign(t, sk, z)}
}

// game is a fixture world: a home for the updater and installs under a scratch root.
type game struct {
	t       *testing.T
	home    string
	root    string
	running map[string]bool // install -> client running
	procErr error
	log     *bytes.Buffer
}

func newGame(t *testing.T) *game {
	t.Helper()
	return &game{t: t, home: t.TempDir(), root: t.TempDir(), running: map[string]bool{}, log: &bytes.Buffer{}}
}

// install makes a flavour folder with an account's saved variables and, when
// addonVer is not empty, the addon (and data when dataVer is not empty) plus
// a stranger's addon beside them.
func (g *game) install(name, addonVer, dataVer string) string {
	g.t.Helper()
	in := filepath.Join(g.root, name)
	g.write(filepath.Join(in, "WTF", "Account", "ACC", "SavedVariables", "ForeverSixty.lua"), "ForeverSixtyDB = { keep = 'me' }\n")
	g.write(filepath.Join(in, "Interface", "AddOns", "Stranger", "Stranger.toc"), "## Version: 9.9.9\n")
	g.write(filepath.Join(in, "Interface", "AddOns", "Stranger", "Stranger.lua"), "-- not ours\n")
	add := AddonsDir(in)
	if addonVer != "" {
		g.write(filepath.Join(add, "ForeverSixty", "ForeverSixty.toc"), "## Version: "+addonVer+"\nCore.lua\n")
		g.write(filepath.Join(add, "ForeverSixty", "Core.lua"), "-- old core "+addonVer+"\n")
		g.write(filepath.Join(add, "ForeverSixtyInbox", "ForeverSixtyInbox.toc"), "## Version: addon-v"+addonVer+"\n")
	}
	if dataVer != "" {
		g.write(filepath.Join(add, "ForeverSixtyData", "ForeverSixtyData.toc"), "## Version: data-v"+dataVer+"\n")
		g.write(filepath.Join(add, "ForeverSixtyData", "Data.lua"), "-- old data\n")
	}
	return in
}

func (g *game) write(p, body string) {
	g.t.Helper()
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		g.t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		g.t.Fatal(err)
	}
}

func (g *game) read(p string) string {
	b, err := os.ReadFile(p)
	if err != nil {
		return "<" + err.Error() + ">"
	}
	return string(b)
}

func (g *game) updater(fd *feed, key string, installs ...string) *Updater {
	return New(Options{
		Dir: filepath.Join(g.home, "addons"), FeedBase: fd.base(), PublicKey: key,
		Log:      slog.New(slog.NewTextHandler(g.log, &slog.HandlerOptions{Level: slog.LevelDebug})),
		Installs: func() []string { return installs },
		Processes: func() ([]Proc, error) {
			if g.procErr != nil {
				return nil, g.procErr
			}
			var out []Proc
			for in, up := range g.running {
				if up {
					out = append(out, Proc{Name: "WowClassic.exe", Path: filepath.Join(in, "WowClassic.exe")})
				}
			}
			return out, nil
		},
		Now: func() time.Time { return t0 },
	})
}

func (g *game) version(install, folder string) string {
	v, _ := ReadVersion(AddonsDir(install), folder)
	return v
}

// untouched asserts the stranger's addon and the saved variables are as the fixture made them.
func (g *game) untouched(install string) {
	g.t.Helper()
	if got := g.read(filepath.Join(AddonsDir(install), "Stranger", "Stranger.lua")); got != "-- not ours\n" {
		g.t.Errorf("a stranger's addon changed: %q", got)
	}
	if got := g.read(filepath.Join(install, "WTF", "Account", "ACC", "SavedVariables", "ForeverSixty.lua")); got != "ForeverSixtyDB = { keep = 'me' }\n" {
		g.t.Errorf("saved variables changed: %q", got)
	}
	entries, _ := os.ReadDir(AddonsDir(install))
	for _, e := range entries {
		switch e.Name() {
		case "Stranger", "ForeverSixty", "ForeverSixtyInbox", "ForeverSixtyData":
		default:
			g.t.Errorf("an unexpected entry appeared under AddOns: %s", e.Name())
		}
	}
}

// ---- TOC versions and comparison ---------------------------------------

func TestReadVersionBothFamilies(t *testing.T) {
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "2026.10.21")
	if v, ok := ReadVersion(AddonsDir(in), "ForeverSixty"); !ok || v != "2.0.15" {
		t.Errorf("addon = %q, %v", v, ok)
	}
	// The packager stamps the tag; the version is the date.
	if v, ok := ReadVersion(AddonsDir(in), "ForeverSixtyData"); !ok || v != "2026.10.21" {
		t.Errorf("data = %q, %v", v, ok)
	}
	if v, ok := ReadVersion(AddonsDir(in), "ForeverSixtyInbox"); !ok || v != "2.0.15" {
		t.Errorf("inbox = %q, %v", v, ok)
	}
	if _, ok := ReadVersion(AddonsDir(in), "Nothing"); ok {
		t.Error("a folder that is not there was found")
	}
	// A BOM, odd spacing and case, and a TOC with no version line.
	g.write(filepath.Join(AddonsDir(in), "ForeverSixty", "ForeverSixty.toc"), "\xef\xbb\xbf##   version :  v2.1.0 \r\nCore.lua\r\n")
	if v, _ := ReadVersion(AddonsDir(in), "ForeverSixty"); v != "2.1.0" {
		t.Errorf("tolerant read = %q", v)
	}
	g.write(filepath.Join(AddonsDir(in), "ForeverSixty", "ForeverSixty.toc"), "## Title: x\n")
	if v, ok := ReadVersion(AddonsDir(in), "ForeverSixty"); !ok || v != "" {
		t.Errorf("no version line = %q, %v", v, ok)
	}
}

func TestNewerComparesSemverAndDates(t *testing.T) {
	for _, tc := range []struct {
		f        Family
		cand, in string
		want     bool
	}{
		{Addon, "2.0.16", "2.0.15", true},
		{Addon, "2.0.15", "2.0.15", false},
		{Addon, "2.0.9", "2.0.15", false},
		{Addon, "2.1.0", "2.0.99", true},
		{Addon, "addon-v10.0.0", "9.9.9", true},
		{Addon, "2.0.16", "", true},
		{Addon, "", "2.0.15", false},
		{Data, "2026.10.22", "2026.10.21", true},
		{Data, "data-v2026.10.21", "2026.10.21", false},
		{Data, "2026.9.30", "2026.10.1", false},
		{Data, "2027.1.1", "2026.12.31", true},
		{Data, "2026.10.22", "", true},
		{Data, "2026.10.22", "garbage", true},
		{Data, "garbage", "2026.10.21", false},
	} {
		if got := tc.f.Newer(tc.cand, tc.in); got != tc.want {
			t.Errorf("%s.Newer(%q, %q) = %v, want %v", tc.f.Key, tc.cand, tc.in, got, tc.want)
		}
	}
}

// ---- is the game running -----------------------------------------------

func TestGameRunningGate(t *testing.T) {
	era := filepath.Join(string(filepath.Separator)+"games", "World of Warcraft", "_classic_era_")
	anni := filepath.Join(string(filepath.Separator)+"games", "World of Warcraft", "_anniversary_")
	for _, tc := range []struct {
		name  string
		procs []Proc
		want  bool
	}{
		{"nothing", nil, false},
		{"unrelated programs", []Proc{{Name: "chrome.exe", Path: "/x/chrome.exe"}, {Name: "Battle.net.exe", Path: "/x/Battle.net.exe"}}, false},
		{"this install's client", []Proc{{Name: "WowClassic.exe", Path: filepath.Join(era, "WowClassic.exe")}}, true},
		{"the other install's client", []Proc{{Name: "WowClassic.exe", Path: filepath.Join(anni, "WowClassic.exe")}}, false},
		{"a client with a path the system would not tell", []Proc{{Name: "WowB.exe"}}, true},
		{"a client beside the flavour folders", []Proc{{Name: "Wow.exe", Path: filepath.Join(filepath.Dir(era), "Wow.exe")}}, true},
		{"a client elsewhere", []Proc{{Name: "Wow.exe", Path: "/other/Wow.exe"}}, false},
		{"this install's Mac client", []Proc{{Name: "World of Warcraft", Path: filepath.Join(era, "World of Warcraft Beta.app", "Contents", "MacOS", "World of Warcraft")}}, true},
		{"the Mac launcher beside the flavour folders", []Proc{{Name: "World of Warcraft Launcher", Path: filepath.Join(filepath.Dir(era), "World of Warcraft Launcher.app", "Contents", "MacOS", "World of Warcraft Launcher")}}, false},
	} {
		if got := GameRunning(era, tc.procs); got != tc.want {
			t.Errorf("%s: GameRunning = %v, want %v", tc.name, got, tc.want)
		}
	}
	// The live game's _forever_: its client counts whether its executable is
	// the plain "World of Warcraft" (as in the beta's .app) or names Forever.
	live := filepath.Join(string(filepath.Separator)+"games", "World of Warcraft", "_forever_")
	for _, name := range []string{"World of Warcraft", "World of Warcraft Forever"} {
		procs := []Proc{{Name: name, Path: filepath.Join(live, name+".app", "Contents", "MacOS", name)}}
		if !GameRunning(live, procs) {
			t.Errorf("the live Mac client %q is not seen as running", name)
		}
		if GameRunning(era, procs) {
			t.Errorf("the live Mac client %q counts for another install", name)
		}
	}
	if !GameRunning(live, []Proc{{Name: "Wow.exe", Path: filepath.Join(live, "Wow.exe")}}) {
		t.Error("the live Windows client is not seen as running")
	}
	for name, want := range map[string]bool{"World of Warcraft Forever": true, "World of Warcraft Forever.app": false, "Wow.exe": true, "WowB.exe": true, "WowT.exe": true, "WowClassic.exe": true, "wowclassic.EXE": true,
		"Wow-64.exe": true, "WowError.exe": true, "Showow.exe": false, "Wow.txt": false, "BlizzardError.exe": false,
		"World of Warcraft": true, "World of Warcraft Classic": true, "World of Warcraft Classic Beta": true,
		"World of Warcraft Launcher": false, "World of Warcraft Beta.app": false, "Battle.net": false} {
		if IsClientName(name) != want {
			t.Errorf("IsClientName(%q) = %v", name, !want)
		}
	}
}

func TestUpdatesWaitForTheGameToClose(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)

	g.running[in] = true
	u.Tick(context.Background(), t0)
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Fatalf("the addon changed while the game ran: %s", v)
	}
	if got := g.read(filepath.Join(AddonsDir(in), "ForeverSixty", "Core.lua")); got != "-- old core 2.0.15\n" {
		t.Errorf("a file changed while the game ran: %q", got)
	}
	// The release waits in the companion's home, not in AddOns.
	if _, ok := u.readStaged(Addon); !ok {
		t.Fatal("the release was not kept for later")
	}
	if left := leftovers(AddonsDir(in)); len(left) != 0 {
		t.Errorf("sibling folders appeared while the game ran: %v", left)
	}
	view := u.View(in, time.Time{})
	if view.Addon.Ready != "2.0.16" || view.Addon.Version != "2.0.15" {
		t.Errorf("view while waiting = %+v", view.Addon)
	}

	// A minute later the game is still up: nothing happens. Then it closes.
	u.Tick(context.Background(), t0.Add(2*time.Minute))
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Fatalf("the addon changed while the game ran: %s", v)
	}
	g.running[in] = false
	u.Tick(context.Background(), t0.Add(3*time.Minute))
	if v := g.version(in, "ForeverSixty"); v != "2.0.16" {
		t.Fatalf("the addon was not installed once the game closed: %s", v)
	}
	if _, ok := u.readStaged(Addon); ok {
		t.Error("the staged release was not cleared after it went in")
	}
	g.untouched(in)
}

func TestAProcessListThatFailsCountsAsRunning(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	g.procErr = errors.New("access denied")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	u.Tick(context.Background(), t0.Add(5*time.Minute))
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed with no way to tell whether the game was up: %s", v)
	}
}

// ---- the schedule and the end-to-end update ------------------------------

func TestTickInstallsBothFamiliesIntoEveryInstall(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedData(t, sk, "2026.10.21"), signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	era := g.install("_classic_era_", "2.0.15", "2026.10.20")
	anni := g.install("_anniversary_", "2.0.16", "2026.10.21") // already current
	bare := g.install("_classic_beta_", "", "")                // no addon: never automatic
	u := g.updater(fd, key, era, anni, bare)

	u.Tick(context.Background(), t0)
	if v := g.version(era, "ForeverSixty"); v != "2.0.16" {
		t.Errorf("era addon = %s", v)
	}
	if v := g.version(era, "ForeverSixtyInbox"); v != "2.0.16" {
		t.Errorf("era inbox = %s", v)
	}
	if v := g.version(era, "ForeverSixtyData"); v != "2026.10.21" {
		t.Errorf("era data = %s", v)
	}
	if got := g.read(filepath.Join(AddonsDir(era), "ForeverSixty", "Core.lua")); got != "-- core 2.0.16\n" {
		t.Errorf("the new files are not in: %q", got)
	}
	if _, err := os.Stat(AddonsDir(bare) + "/ForeverSixty"); err == nil {
		t.Error("the addon was installed where it did not exist, with no button pressed")
	}
	if _, err := os.Stat(AddonsDir(bare) + "/ForeverSixtyData"); err == nil {
		t.Error("the data addon was installed where there was no addon")
	}
	if left := leftovers(AddonsDir(era)); len(left) != 0 {
		t.Errorf("leftovers: %v", left)
	}
	for _, in := range []string{era, anni, bare} {
		g.untouched(in)
	}
	// A staged copy is not kept once nothing needs it.
	if _, ok := u.readStaged(Addon); ok {
		t.Error("a release nobody needs is still staged")
	}
	// The six-hour rule: no second list within the interval.
	before := fd.reads
	u.Tick(context.Background(), t0.Add(time.Hour))
	if fd.reads != before {
		t.Error("the feed was read again inside six hours")
	}
	u.Tick(context.Background(), t0.Add(7*time.Hour))
	if fd.reads == before {
		t.Error("the feed was not read after six hours")
	}
}

func TestDataAddonIsInstalledBesideAnExistingAddon(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedData(t, sk, "2026.10.21"), signedAddon(t, sk, "2.0.15"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if v := g.version(in, "ForeverSixtyData"); v != "2026.10.21" {
		t.Errorf("data = %q; it should come with the next update when the addon is already there", v)
	}
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("an addon that was current changed: %s", v)
	}
	g.untouched(in)
}

func TestSwitchOffDisablesChecksAndInstalls(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	u.o.Enabled = func() bool { return false }
	u.Tick(context.Background(), t0)
	if fd.reads != 0 {
		t.Error("the feed was read with the switch off")
	}
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed with the switch off: %s", v)
	}
	// The button still works on press.
	bare := g.install("_forever_", "", "")
	u.o.Installs = func() []string { return []string{in, bare} }
	if err := u.Install(context.Background(), bare); err != nil {
		t.Fatalf("Install with the switch off: %v", err)
	}
	if v := g.version(bare, "ForeverSixty"); v != "2.0.16" {
		t.Errorf("the button installed %q", v)
	}
}

func TestNoKeyInstallsNothing(t *testing.T) {
	sk, _ := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, "", in)
	u.Tick(context.Background(), t0)
	if fd.reads != 0 {
		t.Error("an unkeyed build listed releases")
	}
	if err := u.Install(context.Background(), in); !errors.Is(err, ErrNoKey) {
		t.Errorf("Install without a key = %v", err)
	}
}

func TestInstallButtonPutsTheAddonWhereItIsMissing(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedData(t, sk, "2026.10.21"), signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	bare := g.install("_classic_era_", "", "")
	u := g.updater(fd, key, bare)

	// The game is running there: refused, nothing written.
	g.running[bare] = true
	if err := u.Install(context.Background(), bare); !errors.Is(err, ErrGameRunning) {
		t.Fatalf("Install while the game runs = %v", err)
	}
	if _, err := os.Stat(AddonsDir(bare) + "/ForeverSixty"); err == nil {
		t.Fatal("the addon went in while the game ran")
	}

	g.running[bare] = false
	if err := u.Install(context.Background(), bare); err != nil {
		t.Fatal(err)
	}
	if g.version(bare, "ForeverSixty") != "2.0.16" || g.version(bare, "ForeverSixtyInbox") != "2.0.16" || g.version(bare, "ForeverSixtyData") != "2026.10.21" {
		t.Errorf("after the button: addon %s inbox %s data %s", g.version(bare, "ForeverSixty"),
			g.version(bare, "ForeverSixtyInbox"), g.version(bare, "ForeverSixtyData"))
	}
	g.untouched(bare)

	// Right after, the row says installed; a login in the game ends that.
	v := u.View(bare, time.Time{})
	if !v.Addon.Fresh || !v.Data.Fresh || v.Addon.Version != "2.0.16" {
		t.Errorf("view right after = %+v", v)
	}
	v = u.View(bare, t0.Add(-time.Minute))
	if !v.Addon.Fresh {
		t.Error("a log older than the install ended the takes-effect line")
	}
	v = u.View(bare, t0.Add(time.Minute))
	if v.Addon.Fresh || v.Data.Fresh {
		t.Errorf("a login after the install did not end the takes-effect line: %+v", v)
	}

	// No release at all.
	empty := newFeed(t)
	u2 := g.updater(empty, key, bare)
	if err := u2.Install(context.Background(), bare); !errors.Is(err, ErrNoRelease) {
		t.Errorf("Install with no release = %v", err)
	}
}

func TestInstalledRecordSurvivesARestart(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	bare := g.install("_classic_era_", "", "")
	u := g.updater(fd, key, bare)
	if err := u.Install(context.Background(), bare); err != nil {
		t.Fatal(err)
	}
	again := g.updater(fd, key, bare)
	if v := again.View(bare, time.Time{}); !v.Addon.Fresh {
		t.Errorf("a new process forgot the install: %+v", v)
	}
}

// ---- signatures ---------------------------------------------------------

func TestSignatureRejection(t *testing.T) {
	sk, key := keypair(t)
	other, _ := keypair(t)
	for name, mk := range map[string]func() release{
		"a signature from another key": func() release {
			r := signedAddon(t, other, "2.0.16")
			return r
		},
		"a signature over other bytes": func() release {
			r := signedAddon(t, sk, "2.0.16")
			r.zip = addonZip(t, "2.0.17")
			return r
		},
		"a signature that is not one": func() release {
			r := signedAddon(t, sk, "2.0.16")
			r.sig = []byte("not a signature")
			return r
		},
	} {
		t.Run(name, func(t *testing.T) {
			fd := newFeed(t)
			fd.publish(mk())
			g := newGame(t)
			in := g.install("_classic_era_", "2.0.15", "")
			u := g.updater(fd, key, in)
			u.Tick(context.Background(), t0)
			if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
				t.Errorf("installed a zip that does not verify: %s", v)
			}
			if _, ok := u.readStaged(Addon); ok {
				t.Error("staged a zip that does not verify")
			}
			if !strings.Contains(g.log.String(), "not taken") {
				t.Errorf("the rejection was not logged:\n%s", g.log)
			}
			if err := u.Install(context.Background(), in); err == nil {
				t.Error("the Install button took a zip that does not verify")
			}
			if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
				t.Errorf("the button installed a zip that does not verify: %s", v)
			}
			g.untouched(in)
		})
	}
}

func TestAZipTamperedAfterStagingIsNotInstalled(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	g.running[in] = true
	u.Tick(context.Background(), t0)
	// Someone swaps the staged file while the game runs.
	if err := os.WriteFile(u.zipPath(Addon), addonZip(t, "2.0.16"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(u.zipPath(Addon), append(addonZip(t, "2.0.16"), 'x'), 0o600); err != nil {
		t.Fatal(err)
	}
	g.running[in] = false
	u.Tick(context.Background(), t0.Add(2*time.Minute))
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed a staged zip that no longer verifies: %s", v)
	}
	if _, ok := u.readStaged(Addon); ok {
		t.Error("the tampered release was kept")
	}
}

// ---- what the feed may and may not name ---------------------------------

func TestAManifestWithoutASignatureIsNoCandidate(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	unsigned := signedAddon(t, sk, "2.0.16")
	unsigned.sig = nil
	fd.publish(unsigned)
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed an unsigned release: %s", v)
	}
	if !strings.Contains(g.log.String(), "the update check found no signed addon release") {
		t.Errorf("the missing signature was not said:\n%s", g.log)
	}
	if strings.Contains(g.log.String(), "check failed") {
		t.Errorf("a manifest without a signature was an error:\n%s", g.log)
	}
	g.untouched(in)
}

func TestNoManifestYetIsNoUpdateAndNoError(t *testing.T) {
	_, key := keypair(t)
	fd := newFeed(t) // every family answers 404
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	out := g.log.String()
	if strings.Contains(out, "check failed") {
		t.Errorf("a 404 manifest was an error:\n%s", out)
	}
	for _, f := range Families {
		if !strings.Contains(out, "the update check found no signed "+f.Key+" release") {
			t.Errorf("the %s family's absence was not said:\n%s", f.Key, out)
		}
	}
	if err := u.Install(context.Background(), in); !errors.Is(err, ErrNoRelease) {
		t.Errorf("Install with no manifest = %v", err)
	}
	g.untouched(in)
}

func TestAManifestTheFeedCannotServeIsAnError(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	fd.status = http.StatusInternalServerError
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	g.updater(fd, key, in).Tick(context.Background(), t0)
	if !strings.Contains(g.log.String(), "release feed answered 500 for the addon manifest") {
		t.Errorf("the failure was not logged:\n%s", g.log)
	}
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("addon = %s", v)
	}
}

func TestAZipOffTheFeedsOriginIsRefused(t *testing.T) {
	sk, key := keypair(t)
	hits := 0
	elsewhere := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		w.Write(addonZip(t, "2.0.16"))
	}))
	t.Cleanup(elsewhere.Close)
	rel := signedAddon(t, sk, "2.0.16")
	rel.zipURL = elsewhere.URL + "/ForeverSixty.zip"
	fd := newFeed(t)
	fd.publish(rel)
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if err := u.Install(context.Background(), in); err == nil || !strings.Contains(err.Error(), "off the feed's origin") {
		t.Errorf("Install = %v", err)
	}
	if hits != 0 {
		t.Errorf("the other host was asked %d times", hits)
	}
	if !strings.Contains(g.log.String(), "off the feed's origin") {
		t.Errorf("the refusal was not logged:\n%s", g.log)
	}
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed a zip from another host: %s", v)
	}
	g.untouched(in)
}

func TestAZipThatDoesNotMatchItsDigestIsRefused(t *testing.T) {
	sk, key := keypair(t)
	rel := signedAddon(t, sk, "2.0.16")
	rel.zipSHA = digest([]byte("another zip"))
	fd := newFeed(t)
	fd.publish(rel)
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if _, ok := u.readStaged(Addon); ok {
		t.Error("staged a zip that missed its digest")
	}
	if !strings.Contains(g.log.String(), "does not match the manifest's sha256") {
		t.Errorf("the mismatch was not logged:\n%s", g.log)
	}
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("installed a zip that missed its digest: %s", v)
	}
	g.untouched(in)
}

// ---- the replace, and what it may touch ---------------------------------

func TestReplaceRollsBackWhenTheMoveFails(t *testing.T) {
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	src := t.TempDir()
	if err := extract(addonZip(t, "2.0.16"), src, Addon.Folders); err != nil {
		t.Fatal(err)
	}
	addons := AddonsDir(in)
	log := slog.New(slog.NewTextHandler(g.log, nil))

	for name, failOn := range map[string]string{
		"the first folder's move in":  "ForeverSixty",
		"the second folder's move in": "ForeverSixtyInbox",
	} {
		t.Run(name, func(t *testing.T) {
			real := rename
			defer func() { rename = real }()
			rename = func(from, to string) error {
				if strings.HasSuffix(from, incomingSuffix) && strings.HasSuffix(strings.TrimSuffix(from, incomingSuffix), failOn) {
					return errors.New("the system would not move it")
				}
				return real(from, to)
			}
			err := replace(addons, src, Addon.Folders, "2.0.16", log)
			if err == nil {
				t.Fatal("a failing move was reported as success")
			}
			if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
				t.Errorf("addon = %q after the rollback, want the old 2.0.15", v)
			}
			if got := g.read(filepath.Join(addons, "ForeverSixty", "Core.lua")); got != "-- old core 2.0.15\n" {
				t.Errorf("the old files are not back: %q", got)
			}
			if v := g.version(in, "ForeverSixtyInbox"); v != "2.0.15" {
				t.Errorf("inbox = %q after the rollback", v)
			}
			if left := leftovers(addons); len(left) != 0 {
				t.Errorf("leftovers after a rollback: %v", left)
			}
			g.untouched(in)
		})
	}
}

func TestReplaceRollsBackWhenTheNewTocReadsBackWrong(t *testing.T) {
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	src := t.TempDir()
	if err := extract(addonZip(t, "2.0.16"), src, Addon.Folders); err != nil {
		t.Fatal(err)
	}
	err := replace(AddonsDir(in), src, Addon.Folders, "2.0.99", slog.New(slog.DiscardHandler))
	if err == nil || !strings.Contains(err.Error(), "put back") {
		t.Fatalf("a wrong read-back = %v", err)
	}
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("addon = %q", v)
	}
	if left := leftovers(AddonsDir(in)); len(left) != 0 {
		t.Errorf("leftovers: %v", left)
	}
}

func TestAFailedUpdateIsLoggedAndLeftAloneForAWhile(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	u := g.updater(fd, key, in)
	real := rename
	defer func() { rename = real }()
	fails := 0
	rename = func(from, to string) error {
		if strings.HasSuffix(from, incomingSuffix) {
			fails++
			return errors.New("the system would not move it")
		}
		return real(from, to)
	}
	u.Tick(context.Background(), t0)
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("addon = %q", v)
	}
	if !strings.Contains(g.log.String(), "could not be installed") || !strings.Contains(g.log.String(), "left as it was") {
		t.Errorf("the rollback was not logged:\n%s", g.log)
	}
	if view := u.View(in, time.Time{}); view.Addon.Ready != "" || view.Addon.Version != "2.0.15" {
		t.Errorf("a failed install still reads as ready: %+v", view.Addon)
	}
	first := fails
	u.Tick(context.Background(), t0.Add(2*time.Minute))
	if fails != first {
		t.Errorf("a failed install was retried inside the hour (%d tries)", fails)
	}
	g.untouched(in)
}

func TestAStaleSiblingFromACrashIsCleared(t *testing.T) {
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	addons := AddonsDir(in)
	// A crash between the two renames: the old folder aside, nothing in its place.
	if err := os.Rename(filepath.Join(addons, "ForeverSixty"), filepath.Join(addons, "ForeverSixty"+previousSuffix)); err != nil {
		t.Fatal(err)
	}
	g.write(filepath.Join(addons, "ForeverSixtyInbox"+incomingSuffix, "junk.lua"), "x")
	src := t.TempDir()
	if err := extract(addonZip(t, "2.0.16"), src, Addon.Folders); err != nil {
		t.Fatal(err)
	}
	if err := replace(addons, src, Addon.Folders, "2.0.16", slog.New(slog.DiscardHandler)); err != nil {
		t.Fatal(err)
	}
	if g.version(in, "ForeverSixty") != "2.0.16" {
		t.Error("not installed")
	}
	if left := leftovers(addons); len(left) != 0 {
		t.Errorf("leftovers: %v", left)
	}
}

func TestOnlyTheNamedFoldersAreEverWritten(t *testing.T) {
	sk, key := keypair(t)
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "2026.10.20")

	// A zip that carries a folder the family does not ship, one that climbs
	// out, a file at the top, a link, a lone folder: each refused whole.
	bad := map[string]map[string]string{
		"an extra folder": {
			"ForeverSixty/ForeverSixty.toc": "## Version: 2.0.16\n", "ForeverSixtyInbox/ForeverSixtyInbox.toc": "## Version: 2.0.16\n",
			"Stranger/Stranger.lua": "-- overwritten\n",
		},
		"a path that climbs out": {
			"ForeverSixty/ForeverSixty.toc": "## Version: 2.0.16\n", "ForeverSixtyInbox/ForeverSixtyInbox.toc": "## Version: 2.0.16\n",
			"ForeverSixty/../../WTF/evil.lua": "x",
		},
		"a file at the top": {
			"ForeverSixty/ForeverSixty.toc": "## Version: 2.0.16\n", "ForeverSixtyInbox/ForeverSixtyInbox.toc": "## Version: 2.0.16\n",
			"README.txt": "x",
		},
		"a backslash path": {
			"ForeverSixty/ForeverSixty.toc": "## Version: 2.0.16\n", "ForeverSixtyInbox/ForeverSixtyInbox.toc": "## Version: 2.0.16\n",
			`ForeverSixty\..\..\x.lua`: "x",
		},
		"a missing inbox": {
			"ForeverSixty/ForeverSixty.toc": "## Version: 2.0.16\n",
		},
	}
	for name, files := range bad {
		t.Run(name, func(t *testing.T) {
			z := mkZip(t, files)
			fd := newFeed(t)
			fd.publish(release{tag: "addon-v2.0.16", zipName: "ForeverSixty-addon-v2.0.16-forever.zip", zip: z, sig: sign(t, sk, z)})
			u := g.updater(fd, key, in)
			u.Tick(context.Background(), t0)
			if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
				t.Errorf("a zip with %s was installed (%s)", name, v)
			}
			if _, err := os.Stat(filepath.Join(in, "WTF", "evil.lua")); err == nil {
				t.Error("a file was written outside AddOns")
			}
			g.untouched(in)
		})
	}

	// A good run: the data addon swapped, the stranger and saved variables as they were.
	fd := newFeed(t)
	fd.publish(signedData(t, sk, "2026.10.21"), signedAddon(t, sk, "2.0.16"))
	g.updater(fd, key, in).Tick(context.Background(), t0)
	if g.version(in, "ForeverSixty") != "2.0.16" || g.version(in, "ForeverSixtyData") != "2026.10.21" {
		t.Fatal("the good run did not install")
	}
	g.untouched(in)
}

func TestAZipWhoseTocDisagreesWithItsTagIsRefused(t *testing.T) {
	sk, key := keypair(t)
	z := addonZip(t, "2.0.15") // the author forgot to bump the TOC
	fd := newFeed(t)
	fd.publish(release{tag: "addon-v2.0.16", zipName: "ForeverSixty-addon-v2.0.16-forever.zip", zip: z, sig: sign(t, sk, z)})
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.14", "")
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if v := g.version(in, "ForeverSixty"); v != "2.0.14" {
		t.Errorf("installed a zip whose TOC says %s under a 2.0.16 tag", v)
	}
	u.Tick(context.Background(), t0.Add(7*time.Hour))
	if g.version(in, "ForeverSixty") != "2.0.14" {
		t.Error("it was installed on the second look")
	}
}

func TestADeveloperCopyIsLeftAlone(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	fd.publish(signedAddon(t, sk, "2.0.16"))
	g := newGame(t)
	in := g.install("_classic_era_", "2.0.15", "")
	// A checkout of the source: a .git inside the folder.
	g.write(filepath.Join(AddonsDir(in), "ForeverSixty", ".git", "HEAD"), "ref: refs/heads/main\n")
	g.updater(fd, key, in).Tick(context.Background(), t0)
	if v := g.version(in, "ForeverSixty"); v != "2.0.15" {
		t.Errorf("replaced a source checkout: %s", v)
	}
	// A link to somewhere else.
	g2 := newGame(t)
	in2 := g2.install("_classic_era_", "", "")
	target := filepath.Join(g2.root, "elsewhere", "ForeverSixty")
	g2.write(filepath.Join(target, "ForeverSixty.toc"), "## Version: 1.0.0\n")
	if err := os.Symlink(target, filepath.Join(AddonsDir(in2), "ForeverSixty")); err != nil {
		t.Skipf("no symlinks here: %v", err)
	}
	g2.updater(fd, key, in2).Tick(context.Background(), t0)
	if got := g2.read(filepath.Join(target, "ForeverSixty.toc")); got != "## Version: 1.0.0\n" {
		t.Errorf("wrote through a link: %q", got)
	}
}

func TestWineClientPathsAreMapped(t *testing.T) {
	for _, tc := range []struct {
		arg  string
		want Proc
		ok   bool
	}{
		{`Z:\home\me\Games\wow\_classic_era_\WowClassic.exe`, Proc{Name: "WowClassic.exe", Path: "/home/me/Games/wow/_classic_era_/WowClassic.exe"}, true},
		{`C:\Program Files (x86)\World of Warcraft\_classic_era_\WowClassic.exe`, Proc{Name: "WowClassic.exe"}, true},
		{`WowClassic.exe`, Proc{Name: "WowClassic.exe"}, true},
		{`/usr/bin/wine64`, Proc{}, false},
	} {
		got := Proc{}
		ok := false
		if IsClientName(tc.arg[strings.LastIndexAny(tc.arg, `\/`)+1:]) {
			got, ok = wineArg(tc.arg)
		}
		if ok != tc.ok || got != tc.want {
			t.Errorf("wineArg(%q) = %+v, %v", tc.arg, got, ok)
		}
	}
}

func TestDefaultsAreSane(t *testing.T) {
	u := New(Options{})
	if u.o.Every != 6*time.Hour || u.o.FeedBase != "https://foreversixty.gg/dl" || u.feed.Base != u.o.FeedBase {
		t.Errorf("defaults = %+v", u.o)
	}
	u.Tick(context.Background(), t0) // no key: nothing, no panic
}

// ---- the newest signed release goes in, never an older one --------------

// stageWhileRunning stages 2.0.16 over a 2.0.15 folder while the game runs.
func stageWhileRunning(t *testing.T, sk minisign.PrivateKey, fd *feed, g *game) string {
	t.Helper()
	fd.publish(signedAddon(t, sk, "2.0.16"))
	in := g.install("_forever_", "2.0.15", "")
	g.running[in] = true
	return in
}

func TestTheCloseReReadsTheFeedAndInstallsTheNewest(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	g := newGame(t)
	in := stageWhileRunning(t, sk, fd, g)
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if s, ok := u.readStaged(Addon); !ok || s.Version != "2.0.16" {
		t.Fatalf("staged = %+v, %v", s, ok)
	}

	fd.publish(signedAddon(t, sk, "2.0.17"))
	g.running[in] = false
	u.Tick(context.Background(), t0.Add(2*time.Minute))
	if v := g.version(in, "ForeverSixty"); v != "2.0.17" {
		t.Fatalf("installed %s, want the newer 2.0.17", v)
	}
	if _, ok := u.readStaged(Addon); ok {
		t.Error("the staged release was not cleared")
	}
	if _, err := os.Stat(u.zipPath(Addon)); err == nil {
		t.Error("the older staged zip was kept")
	}
	if !strings.Contains(g.log.String(), "replaces the staged 2.0.16") {
		t.Errorf("the log does not say the newer release took over:\n%s", g.log)
	}
	g.untouched(in)
}

func TestAStagedReleaseOlderThanTheFolderIsDiscarded(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	g := newGame(t)
	in := stageWhileRunning(t, sk, fd, g)
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)

	// The folder moves on to 2.0.17 by another road while the game runs.
	g.write(filepath.Join(AddonsDir(in), "ForeverSixty", "ForeverSixty.toc"), "## Version: 2.0.17\nCore.lua\n")
	g.write(filepath.Join(AddonsDir(in), "ForeverSixty", "Core.lua"), "-- core 2.0.17\n")
	if view := u.View(in, time.Time{}); view.Addon.Ready != "" {
		t.Errorf("the card still says %q is waiting", view.Addon.Ready)
	}
	g.running[in] = false
	u.Tick(context.Background(), t0.Add(2*time.Minute))
	if v := g.version(in, "ForeverSixty"); v != "2.0.17" {
		t.Fatalf("the folder is %s, a downgrade happened", v)
	}
	if got := g.read(filepath.Join(AddonsDir(in), "ForeverSixty", "Core.lua")); got != "-- core 2.0.17\n" {
		t.Errorf("a file was overwritten: %q", got)
	}
	if _, ok := u.readStaged(Addon); ok {
		t.Error("the staged release was kept")
	}
	if want := "addon 2.0.16 is not newer than the installed 2.0.17; nothing to install"; strings.Count(g.log.String(), want) != 1 {
		t.Errorf("want exactly one line %q in:\n%s", want, g.log)
	}
}

func TestAFeedFailureFallsBackToTheStagedReleaseOnlyIfNewer(t *testing.T) {
	for _, tc := range []struct {
		folder, want string
	}{{"2.0.15", "2.0.16"}, {"2.0.17", "2.0.17"}} {
		t.Run(tc.folder, func(t *testing.T) {
			sk, key := keypair(t)
			fd := newFeed(t)
			g := newGame(t)
			in := stageWhileRunning(t, sk, fd, g)
			u := g.updater(fd, key, in)
			u.Tick(context.Background(), t0)
			g.write(filepath.Join(AddonsDir(in), "ForeverSixty", "ForeverSixty.toc"), "## Version: "+tc.folder+"\nCore.lua\n")

			fd.Close() // the feed cannot be read at the moment of the close
			g.running[in] = false
			u.Tick(context.Background(), t0.Add(2*time.Minute))
			if v := g.version(in, "ForeverSixty"); v != tc.want {
				t.Errorf("folder %s ended at %s, want %s", tc.folder, v, tc.want)
			}
			if _, ok := u.readStaged(Addon); ok {
				t.Error("the staged release is still waiting")
			}
		})
	}
}

func TestTheStagedLineNamesTheVersionAndFollowsANewerRelease(t *testing.T) {
	sk, key := keypair(t)
	fd := newFeed(t)
	g := newGame(t)
	in := stageWhileRunning(t, sk, fd, g)
	u := g.updater(fd, key, in)
	u.Tick(context.Background(), t0)
	if r := u.View(in, time.Time{}).Addon.Ready; r != "2.0.16" {
		t.Fatalf("ready = %q", r)
	}
	fd.publish(signedAddon(t, sk, "2.0.18"))
	u.Tick(context.Background(), t0.Add(7*time.Hour)) // the six-hour check, game still up
	if r := u.View(in, time.Time{}).Addon.Ready; r != "2.0.18" {
		t.Errorf("ready after the next check = %q, want 2.0.18", r)
	}
}

func TestVersionComparisonIsNumericPerComponent(t *testing.T) {
	for _, tc := range []struct {
		cand, in string
		want     bool
	}{
		{"2.0.18", "2.0.9", true},
		{"2.0.9", "2.0.18", false},
		{"2.0.16", "2.0.17", false},
		{"2.0.17", "2.0.17", false},
		{"2.10.0", "2.9.9", true},
	} {
		if got := Addon.Newer(tc.cand, tc.in); got != tc.want {
			t.Errorf("Newer(%q, %q) = %v, want %v", tc.cand, tc.in, got, tc.want)
		}
	}
}
