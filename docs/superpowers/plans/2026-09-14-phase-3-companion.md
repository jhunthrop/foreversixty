# Phase 3 Companion Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the `companion/` Go module — the desktop app that tails World of Warcraft's combat log, runs the `logs/` engine on it, uploads each fight as it closes, streams the raw log up in the background, keeps the ForeverSixty addon in step with the site, shows all of it in a tray window, and updates itself from signed GitHub Releases — plus the workflow that builds and publishes it.

**Architecture:** The companion is a loop with three halves that never block each other. `watch` turns the Logs directory into boundaries and bytes with no clock and no goroutine of its own, so the thirty-minute and fifteen-minute rules are testable in a millisecond. `pipeline` feeds those bytes to one `logs/engine/session` and turns what falls out — closed fights, raw chunks, the completion — into entries in a durable on-disk `queue`; a separate `Drain` sends the queue to the API strictly in order, stopping on the first retryable failure so nothing overtakes anything. `state` holds one file per report, including the engine's serialised session, which is what makes a restart mid-raid resume the same report rather than split the night. `app` assembles all of it behind a loopback HTTP server, and `shell` puts a tray icon and a native webview in front of that — in two builds, one with cgo and one without, so CI needs no desktop libraries.

**Tech Stack:** Go 1.25, the repo's `logs/` module through a `replace`, and five direct dependencies: `github.com/klauspost/compress` (zstd for raw chunks), `github.com/zalando/go-keyring` (the device token), `github.com/jedisct1/go-minisign` (update signatures), `github.com/getlantern/systray` and `github.com/webview/webview_go` (the window, cgo). No test framework beyond stdlib `testing`; no HTTP framework; no Lua interpreter.

**Spec:** `docs/superpowers/specs/2026-09-13-logs-engine-design.md` sections 2, 5, 8 and 9, bound by the interface contract in `docs/superpowers/specs/2026-09-14-phase-3-interfaces.md`. This plan implements the "Companion (Go)" section of that contract and the companion parts of spec section 5; the API routes it calls, the report page, rankings and accounts belong to the api and web plans.

## Global Constraints

- New top-level module `companion/`, module path `github.com/jhunthrop/foreversixty/companion`, `go 1.25.11` — the same directive `api/go.mod` and `logs/go.mod` use. Run every `go` command from `companion/`.
- The engine comes from the repository: `require github.com/jhunthrop/foreversixty/logs v0.0.0` with `replace github.com/jhunthrop/foreversixty/logs => ../logs`. The companion must parse with exactly the code the server parses with.
- Direct dependencies are exactly five, pinned: `github.com/klauspost/compress v1.20.0`, `github.com/zalando/go-keyring v0.2.8`, `github.com/jedisct1/go-minisign v0.0.0-20260527172527-a09352b57a22`, `github.com/getlantern/systray v1.2.2`, `github.com/webview/webview_go v0.0.0-20240831120633-6173450d4dd6`. Anything else needs a written justification in the commit body.
- **Coverage floor: 80% of statements outside the window.** Measured with `go test -tags nogui ./... -race -coverprofile=cover.out -coverpkg=./...`, then the profile filtered to drop `cmd/`, `internal/shell/` and `ui/` — the tray, the webview and three static files, which no Go test can exercise without a desktop. The reference implementation this plan was written from measures 80.8%.
- **Testing rule: implementers run only the tests for the packages they touch.** The full suite runs once at the whole-branch final review; CI is the last gate.
- **Two builds, one behaviour.** `-tags nogui` compiles the companion without cgo, systray or webview; CI uses it for everything. The default build adds the window. Both honour the `-headless` flag at runtime. Nothing outside `internal/shell` may import systray or webview.
- **All offsets the companion sends are report-relative**, counted from the first byte of the report's own stream, not from the start of the log file. `state.Report.StartOffset` is the single place the file offset and the report offset meet. See "Contract decisions" below.
- **Nothing is parsed that belongs to someone else.** FS1 and FSB1 export strings are the addon's format and the site's; the companion moves them as opaque strings and never inspects them. Likewise the ruleset segment of a character key: whatever the addon or the log wrote there travels through untouched, because the Sept 17 beta log has not yet settled what Forever puts in it. No list of rulesets or realms is hardcoded anywhere in this module.
- **Fixture policy**, inherited from the engine: every fixture line is hand-written from the v16 shapes the engine verified, and uses the invented characters listed in `logs/README.md`. Nothing is copied out of a downloaded log.
- Commits use a conventional subject, a body, and one final `-m` carrying both trailers. Never bypass git hooks.
  ```
  Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
  Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k
  ```

---

## Contract decisions

The interface contract leaves five things to the companion. They are decided here, once, and every task below depends on these answers.

**1. Offsets are report-relative.** The contract addresses raw chunks by `offset` and fights by `raw_range.start_offset`, and the engine's session expects its first `Feed` at offset zero. A combat log holds many reports over a night, so the two cannot both be file offsets. Every offset the companion sends — the raw chunk's `?offset=`, the fight's `raw_range`, the completion's `final_offset` — counts from the first byte of *that report*, so `reports/<id>/raw/0.zst` is always the first chunk of the report and the chunks concatenate into exactly the bytes the engine parsed. `state.Report.StartOffset` records where the report begins in the file; `watch` is the only package that works in file offsets.

**2. Raw chunks are 4 MiB of uncompressed log.** The contract says "raw chunks of 4 MiB zstd" and caps the route's body at 8 MiB. Four mebibytes of combat log compresses to a few hundred kilobytes, comfortably inside the cap, and the chunk boundary lands on a multiple of 4 MiB from the report's start so a re-sent chunk overwrites itself exactly.

**3. `MetricsRow` is built from the fight's roster.** The contract says the engine exposes `summary.MetricsRow` "if present, else the API defines it and the companion copies". It is not present: `logs/engine/summary` has `MetricRow`, which is one row per player *per metric*. The contract's shape is one row per player with `metric_dps` and `metric_hps` side by side, so the companion pivots `summary.Summary.Roster` — which already carries GUID, name, class, spec, role, item level, DPS, HPS, damage taken, active time and deaths — and takes the encounter, difficulty, size, duration and kill from the fight. The field names are the contract's, verbatim.

**4. The fifteen- and thirty-minute rules compose.** Fifteen quiet minutes complete the open report; a size increase after a gap longer than thirty minutes starts a new one. Because completion happens first, the thirty-minute rule only bites when a report is still open — a rotation, a truncation, or a companion that was not running when the gap happened. Both are implemented in `watch` and tested separately.

**5. The inbox lives beside the saved variables, and the game may overwrite it.** The contract says the companion "writes `ForeverSixtyInbox.lua`". It goes in the same `SavedVariables` folder as `ForeverSixty.lua`, as a Lua chunk assigning the global `ForeverSixtyInbox`, which the addon declares in its TOC (a Phase 2 follow-up). The game rewrites its own saved variables at logout, so an inbox written while the player is online can be replaced; the ten-minute pass puts it back and the addon reads it at the next load. The companion skips the write when the bytes would not change.

Two settings are added to `config.json` beyond the four the contract names, both optional and both additive: `site_base_url`, so report links can point at a staging site, and `logging_character`, which is how the player answers spec section 5's "reports which character is logging for attribution" without the companion guessing from the log.

---

## File structure

```
companion/
  go.mod, go.sum                    module, replace ../logs, five direct dependencies
  README.md                         install, pairing, troubleshooting
  cmd/foreversixty-companion/
    main.go                         flags, startup order, signal handling
  ui/
    assets.go                       package ui: the embedded window
    index.html                      four pages
    styles.css                      design tokens copied from web/src/styles/tokens.css
    app.js                          polls /api/status, posts three things
  internal/
    paths/paths.go                  the one directory the companion owns, per OS
    character/character.go          region, ruleset, name -- shared by config, state and the wire
    config/config.go                config.json
    logging/logging.go              slog to a size-capped file and stderr
    client/client.go                envelope, device token, retry and backoff
    client/ingest.go                report, fight bundle, live, raw, complete
    client/devices.go               pairing
    client/addon.go                 exports out, inbox in
    secret/secret.go                keychain with a config.json fallback
    queue/queue.go                  the durable outbox, strictly in order
    wow/wow.go                      finding the game
    watch/watch.go                  boundaries and bytes out of the Logs directory
    state/state.go                  one file per report, holding the engine's session
    pipeline/pipeline.go            bytes in, queue entries out
    pipeline/drain.go               queue out, API in
    addon/lua.go                    the subset of Lua saved variables use
    addon/addon.go                  exports and the inbox file
    addon/sync.go                   the sync loop
    updater/updater.go              GitHub Releases and minisign
    icon/icon.go                    the tray icon, drawn rather than shipped
    shell/shell.go                  window options
    shell/shell_gui.go              tray + webview (cgo)
    shell/shell_nogui.go            headless (-tags nogui)
    app/app.go                      everything assembled
    app/http.go                     the loopback HTTP interface
    fakeapi/fakeapi.go              the ingest contract, in memory, for tests
    fixture/fixture.go              the combat log the tests parse
  integration/
    integration_test.go             a raid night with a drop and a restart
.github/workflows/companion.yml     test, build the four targets on a tag, sign, publish
```

Dependency direction is one way: `character` and `config` depend on nothing; `client`, `secret`, `queue`, `wow`, `watch`, `state` depend on those and the engine; `pipeline` and `addon` depend on those; `app` depends on everything but `shell`; `shell` and `cmd` depend on `app`. Nothing outside `shell` imports a desktop toolkit.

---


## Task 1: The module, the workspace, the OS directories, the character triple, config and logging

**Files:**
- Create: `companion/go.mod`, `companion/internal/paths/paths.go`, `companion/internal/character/character.go`, `companion/internal/config/config.go`, `companion/internal/logging/logging.go`
- Modify (idempotently): `go.work` at the repository root
- Test: `companion/internal/paths/paths_test.go`, `companion/internal/character/character_test.go`, `companion/internal/config/config_test.go`, `companion/internal/logging/logging_test.go`

**Interfaces:**
- Consumes: nothing. This is the bottom of the module.
- Produces:
  - `paths.HomeEnv` (the `FS_COMPANION_HOME` override), `paths.Home() (string, error)`, `paths.Dirs{Home, Queue, State, Logs, Update string}`, `paths.Resolve() (Dirs, error)`, `Dirs.ConfigFile() string`, `Dirs.StateFile(localID string) string`.
  - `character.Character{Region, Ruleset, Name string}` with json tags `region`, `ruleset`, `name`; `.Valid() bool`; `.Key() string`.
  - `config.DefaultAPIBaseURL`, `config.DefaultSiteBaseURL`, `config.Visibilities`, `config.Config{APIBaseURL, SiteBaseURL, DeviceToken string; WoWPaths []string; ReportVisibility string; LoggingCharacter *character.Character}`, `config.Default() Config`, `Config.Validate() error`, `config.Load(path) (Config, error)`, `config.Save(path, Config) error`.
  - `logging.MaxBytes`, `logging.File`, `logging.Open(dir string, max int64) (*File, error)`, `logging.New(dir string, debug bool) (*slog.Logger, *File, error)`.

The one design decision worth naming: `paths.Home` is overridable by an environment variable. Every test in this plan sets it, which is why no test ever writes to a real application directory.

- [ ] **Step 1: Create the module**

```bash
mkdir -p companion/internal/{paths,character,config,logging}
cd companion
cat > go.mod <<'EOF'
module github.com/jhunthrop/foreversixty/companion

go 1.25.11

replace github.com/jhunthrop/foreversixty/logs => ../logs
EOF
```

- [ ] **Step 2: Add `companion` to the repository workspace**

The api plan creates `go.work` at the repository root. This step works whether it has run or not, and running it twice changes nothing: `go work use` adds a directive only if it is missing.

```bash
cd "$(git rev-parse --show-toplevel)"
[ -f go.work ] || go work init
go work use ./logs ./companion
[ -f api/go.mod ] && go work use ./api
go work sync
cat go.work
```

Expected: a `use` block naming `./companion`, `./logs`, and `./api` if the api module exists yet.

- [ ] **Step 3: Write the failing tests**

```go
package paths

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHomeHonoursTheOverride(t *testing.T) {
	dir := t.TempDir()
	t.Setenv(HomeEnv, dir)
	got, err := Home()
	if err != nil {
		t.Fatal(err)
	}
	if got != dir {
		t.Fatalf("Home() = %q, want %q", got, dir)
	}
}

func TestHomeIsTheOSApplicationDirectory(t *testing.T) {
	t.Setenv(HomeEnv, "")
	got, err := Home()
	if err != nil {
		t.Fatal(err)
	}
	base, err := os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Dir(got) != base {
		t.Fatalf("Home() = %q, want a child of %q", got, base)
	}
	if name := filepath.Base(got); name != "ForeverSixty" && name != "foreversixty" {
		t.Fatalf("directory name = %q", name)
	}
}

func TestResolveCreatesEveryDirectory(t *testing.T) {
	root := t.TempDir()
	t.Setenv(HomeEnv, filepath.Join(root, "app"))
	d, err := Resolve()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{d.Home, d.Queue, d.State, d.Logs, d.Update} {
		fi, err := os.Stat(p)
		if err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		if !fi.IsDir() {
			t.Errorf("%s is not a directory", p)
		}
	}
	if want := filepath.Join(d.Home, "config.json"); d.ConfigFile() != want {
		t.Errorf("ConfigFile() = %q, want %q", d.ConfigFile(), want)
	}
	if want := filepath.Join(d.State, "abc.json"); d.StateFile("abc") != want {
		t.Errorf("StateFile() = %q, want %q", d.StateFile("abc"), want)
	}
}
```

```go
package character

import "testing"

func TestTheKeyIsRegionRulesetAndTheNameSlug(t *testing.T) {
	for _, tc := range []struct {
		in   Character
		want string
	}{
		{Character{"US", "Hardcore", "Morrowlyn"}, "us/hardcore/morrowlyn"},
		{Character{"eu", "pvp", "Grim Batol"}, "eu/pvp/grim-batol"},
		{Character{"us", "normal", " Thalgrit "}, "us/normal/thalgrit"},
	} {
		if got := tc.in.Key(); got != tc.want {
			t.Errorf("%+v.Key() = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestValidNeedsAllThreeSegments(t *testing.T) {
	if !(Character{"us", "normal", "Morrowlyn"}).Valid() {
		t.Error("a filled-in character is not valid")
	}
	for _, c := range []Character{
		{"", "normal", "Morrowlyn"},
		{"us", "", "Morrowlyn"},
		{"us", "normal", "  "},
	} {
		if c.Valid() {
			t.Errorf("%+v is valid", c)
		}
	}
}
```

```go
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOfAMissingFileGivesTheDefaults(t *testing.T) {
	got, err := Load(filepath.Join(t.TempDir(), "config.json"))
	if err != nil {
		t.Fatal(err)
	}
	if got.APIBaseURL != DefaultAPIBaseURL || got.SiteBaseURL != DefaultSiteBaseURL ||
		got.ReportVisibility != "public" || len(got.WoWPaths) != 0 {
		t.Fatalf("Load = %+v", got)
	}
}

func TestSaveThenLoadRoundTrips(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	want := Config{
		APIBaseURL:       "http://127.0.0.1:8080",
		SiteBaseURL:      "http://127.0.0.1:4321",
		DeviceToken:      "fsd_abcdef",
		WoWPaths:         []string{"/Applications/World of Warcraft/_classic_era_"},
		ReportVisibility: "unlisted",
	}
	if err := Save(path, want); err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if perm := fi.Mode().Perm(); perm != 0o600 {
		t.Errorf("permissions = %v, want 0600", perm)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if got.APIBaseURL != want.APIBaseURL || got.DeviceToken != want.DeviceToken ||
		got.ReportVisibility != want.ReportVisibility || len(got.WoWPaths) != 1 ||
		got.WoWPaths[0] != want.WoWPaths[0] {
		t.Fatalf("Load = %+v, want %+v", got, want)
	}
}

func TestAnUnknownVisibilityIsRefused(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := Save(path, Config{APIBaseURL: "x", SiteBaseURL: "y", ReportVisibility: "everyone"}); err == nil {
		t.Fatal("Save accepted an unknown visibility")
	}
	if err := os.WriteFile(path, []byte(`{"report_visibility":"everyone"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil {
		t.Fatal("Load accepted an unknown visibility")
	}
}

func TestAnEmptyAPIBaseURLIsRefused(t *testing.T) {
	if err := (Config{SiteBaseURL: "y", ReportVisibility: "public"}).Validate(); err == nil {
		t.Fatal("Validate accepted an empty api_base_url")
	}
	if err := (Config{APIBaseURL: "x", ReportVisibility: "public"}).Validate(); err == nil {
		t.Fatal("Validate accepted an empty site_base_url")
	}
}
```

```go
package logging

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTheLogFileRollsOverAtItsCap(t *testing.T) {
	dir := t.TempDir()
	w, err := Open(dir, 64)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	for range 10 {
		if _, err := w.Write([]byte(strings.Repeat("x", 32) + "\n")); err != nil {
			t.Fatal(err)
		}
	}
	fi, err := os.Stat(filepath.Join(dir, "companion.log"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.Size() > 64 {
		t.Errorf("live log is %d bytes, want at most 64", fi.Size())
	}
	if _, err := os.Stat(filepath.Join(dir, "companion.log.1")); err != nil {
		t.Errorf("no rolled file: %v", err)
	}
}

func TestNewWritesJSONRecordsToTheFile(t *testing.T) {
	dir := t.TempDir()
	log, f, err := New(dir, true)
	if err != nil {
		t.Fatal(err)
	}
	log.Info("tailing", "component", "watch", "path", "/tmp/WoWCombatLog.txt")
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "companion.log"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"component":"watch"`) || !strings.Contains(string(b), `"msg":"tailing"`) {
		t.Fatalf("log = %s", b)
	}
}

func TestDebugRecordsAppearOnlyWhenAsked(t *testing.T) {
	dir := t.TempDir()
	quiet, f, err := New(dir, false)
	if err != nil {
		t.Fatal(err)
	}
	quiet.Debug("hidden", "component", "test")
	quiet.Info("shown", "component", "test")
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "companion.log"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), "hidden") {
		t.Error("a debug record was written at info level")
	}
	if !strings.Contains(string(b), "shown") {
		t.Error("the info record is missing")
	}
}

func TestOpenWithNoCapUsesTheDefault(t *testing.T) {
	w, err := Open(t.TempDir(), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	if w.max != MaxBytes {
		t.Fatalf("max = %d, want %d", w.max, MaxBytes)
	}
}

func TestOpenInAnImpossiblePlaceIsAnError(t *testing.T) {
	if _, err := Open(filepath.Join(t.TempDir(), "missing"), 0); err == nil {
		t.Fatal("Open accepted a directory that does not exist")
	}
	if _, _, err := New(filepath.Join(t.TempDir(), "missing"), false); err == nil {
		t.Fatal("New accepted a directory that does not exist")
	}
}
```

- [ ] **Step 4: Run them to verify they fail**

Run: `cd companion && go test ./internal/...`
Expected: FAIL in each package with "undefined: Home", "undefined: Character", "undefined: Load", "undefined: Open".

- [ ] **Step 5: Write `internal/paths/paths.go`**

```go
// companion/internal/paths/paths.go
// Package paths resolves the one directory the companion owns on each
// operating system. Everything else the app writes hangs off it, so a
// test can redirect the whole application state with one variable.
package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// HomeEnv overrides the application directory. Tests set it; the
// installer never does.
const HomeEnv = "FS_COMPANION_HOME"

// Home is the application directory for this operating system:
// ~/Library/Application Support/ForeverSixty on macOS,
// %APPDATA%\ForeverSixty on Windows, ~/.config/foreversixty elsewhere.
func Home() (string, error) {
	if v := os.Getenv(HomeEnv); v != "" {
		return v, nil
	}
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("locate the application directory: %w", err)
	}
	name := "foreversixty"
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		name = "ForeverSixty"
	}
	return filepath.Join(base, name), nil
}

// Dirs are the directories the companion writes into.
type Dirs struct {
	Home   string // the application directory itself
	Queue  string // pending bundles and chunks
	State  string // one file per report, holding session.State()
	Logs   string // the companion's own log files
	Update string // downloaded and verified update binaries
}

// Resolve returns the directories and creates them.
func Resolve() (Dirs, error) {
	home, err := Home()
	if err != nil {
		return Dirs{}, err
	}
	d := Dirs{
		Home:   home,
		Queue:  filepath.Join(home, "queue"),
		State:  filepath.Join(home, "state"),
		Logs:   filepath.Join(home, "logs"),
		Update: filepath.Join(home, "update"),
	}
	for _, p := range []string{d.Home, d.Queue, d.State, d.Logs, d.Update} {
		if err := os.MkdirAll(p, 0o700); err != nil {
			return Dirs{}, fmt.Errorf("create %s: %w", p, err)
		}
	}
	return d, nil
}

// ConfigFile is the path of config.json.
func (d Dirs) ConfigFile() string { return filepath.Join(d.Home, "config.json") }

// StateFile is the path of one report's state file.
func (d Dirs) StateFile(localID string) string {
	return filepath.Join(d.State, localID+".json")
}
```

- [ ] **Step 6: Write `internal/character/character.go`**

```go
// companion/internal/character/character.go
// Package character holds the one attribution triple the config, the
// report state and the wire all share, so the ruleset field is
// spelled once.
package character

import "strings"

// Character is region, ruleset and name.
//
// Forever has no realms: the second segment of a character key is the
// ruleset (normal, pvp, rp, hardcore). The combat log and the addon
// still identify units as Name-Realm, so the companion carries that
// segment through untouched and the API maps it once the Sept 17 beta
// log settles what it holds. Nothing here validates the value against
// a list.
type Character struct {
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Name    string `json:"name"`
}

// Valid reports whether all three segments are filled in.
func (c Character) Valid() bool {
	return strings.TrimSpace(c.Region) != "" &&
		strings.TrimSpace(c.Ruleset) != "" &&
		strings.TrimSpace(c.Name) != ""
}

// Key is the contract's character key: region, ruleset and the
// two-part name lowercased with spaces turned into hyphens.
func (c Character) Key() string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(c.Name), " ", "-"))
	return strings.ToLower(c.Region) + "/" + strings.ToLower(c.Ruleset) + "/" + slug
}
```

- [ ] **Step 7: Write `internal/config/config.go`**

```go
// companion/internal/config/config.go
// Package config is config.json: the four settings the interface
// contract names, loaded with defaults and saved atomically.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
)

// DefaultAPIBaseURL is where the companion talks to unless the user or
// a test points it somewhere else, and DefaultSiteBaseURL is where a
// finished report can be read.
const (
	DefaultAPIBaseURL  = "https://api.foreversixty.gg"
	DefaultSiteBaseURL = "https://foreversixty.gg"
)

// Visibilities are the values the API accepts for a new report.
var Visibilities = []string{"public", "unlisted", "private", "guild"}

// Config is config.json.
type Config struct {
	APIBaseURL string `json:"api_base_url"`
	// SiteBaseURL is where report links point. It is additive to the
	// four keys the contract names, so a staging build can send its
	// links somewhere else.
	SiteBaseURL string `json:"site_base_url"`
	// DeviceToken is only written when no OS keychain is available; the
	// secret store owns the token otherwise and leaves this empty.
	DeviceToken      string   `json:"device_token"`
	WoWPaths         []string `json:"wow_paths"`
	ReportVisibility string   `json:"report_visibility"`
	// LoggingCharacter is who the reports are attributed to, chosen on
	// the settings page from the characters the addon sync has seen.
	// It is additive to the four keys the contract names and is
	// omitted when the player has not chosen one.
	LoggingCharacter *character.Character `json:"logging_character,omitempty"`
}

// Default is the configuration a fresh install starts from.
func Default() Config {
	return Config{
		APIBaseURL:       DefaultAPIBaseURL,
		SiteBaseURL:      DefaultSiteBaseURL,
		WoWPaths:         []string{},
		ReportVisibility: "public",
	}
}

// Validate rejects a configuration the rest of the app cannot use.
func (c Config) Validate() error {
	if c.APIBaseURL == "" {
		return errors.New("api_base_url is empty")
	}
	if c.SiteBaseURL == "" {
		return errors.New("site_base_url is empty")
	}
	if !slices.Contains(Visibilities, c.ReportVisibility) {
		return fmt.Errorf("report_visibility %q is not one of %v", c.ReportVisibility, Visibilities)
	}
	return nil
}

// Load reads config.json. A missing file is not an error: it yields the
// defaults, so a first run needs no installer step.
func Load(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return Default(), nil
	}
	if err != nil {
		return Config{}, fmt.Errorf("read %s: %w", path, err)
	}
	c := Default()
	if err := json.Unmarshal(b, &c); err != nil {
		return Config{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if c.WoWPaths == nil {
		c.WoWPaths = []string{}
	}
	if err := c.Validate(); err != nil {
		return Config{}, fmt.Errorf("%s: %w", path, err)
	}
	return c, nil
}

// Save writes config.json atomically with owner-only permissions, so a
// crash mid-write can never leave an unreadable configuration and the
// fallback token is not world-readable.
func Save(path string, c Config) error {
	if err := c.Validate(); err != nil {
		return err
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')
	tmp, err := os.CreateTemp(filepath.Dir(path), ".config-*.json")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(name, path)
}
```

- [ ] **Step 8: Write `internal/logging/logging.go`**

```go
// companion/internal/logging/logging.go
// Package logging gives the companion one slog logger that writes JSON
// to a size-capped file and to stderr. A desktop app runs for weeks, so
// the file caps itself rather than growing without bound; there is no
// log-rotation dependency for two files.
package logging

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

// MaxBytes is how large companion.log grows before it is rolled onto
// companion.log.1 and started again.
const MaxBytes = 4 << 20

// File is an io.Writer that rolls one file over at MaxBytes.
type File struct {
	path string
	max  int64

	mu sync.Mutex
	f  *os.File
	n  int64
}

// Open opens the log file in dir, creating it if it is not there.
func Open(dir string, max int64) (*File, error) {
	if max <= 0 {
		max = MaxBytes
	}
	w := &File{path: filepath.Join(dir, "companion.log"), max: max}
	if err := w.reopen(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *File) reopen() error {
	f, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return err
	}
	w.f, w.n = f, fi.Size()
	return nil
}

// Write appends one record, rolling the file over first when it is full.
func (w *File) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.n+int64(len(p)) > w.max {
		if err := w.roll(); err != nil {
			return 0, err
		}
	}
	n, err := w.f.Write(p)
	w.n += int64(n)
	return n, err
}

func (w *File) roll() error {
	if err := w.f.Close(); err != nil {
		return err
	}
	if err := os.Rename(w.path, w.path+".1"); err != nil {
		return err
	}
	return w.reopen()
}

// Close closes the file.
func (w *File) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.f.Close()
}

// New returns the application logger and the file behind it. Every
// record carries the component so one grep separates the tail from the
// uploader.
func New(dir string, debug bool) (*slog.Logger, *File, error) {
	f, err := Open(dir, MaxBytes)
	if err != nil {
		return nil, nil, err
	}
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	h := slog.NewJSONHandler(io.MultiWriter(f, os.Stderr), &slog.HandlerOptions{Level: level})
	return slog.New(h), f, nil
}
```

- [ ] **Step 9: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./... && go test ./internal/... -race`
Expected: no gofmt output, no vet output, PASS in all four packages.

- [ ] **Step 10: Commit**

```bash
git add companion/go.mod \
  companion/internal/paths/ \
  companion/internal/character/ \
  companion/internal/config/ \
  companion/internal/logging/ \
  go.work \
  go.work.sum
git commit -m "feat(companion): module scaffold, OS directories, config and logging" \
  -m "The bottom of the companion module: the one directory it owns on each
operating system, the region/ruleset/name triple the config, the report
state and the wire all share, config.json with atomic writes and
owner-only permissions, and a slog logger over a size-capped file.

paths.Home is overridable through FS_COMPANION_HOME so no test ever
writes to a real application directory. config.json carries two keys
beyond the four the interface contract names -- site_base_url and
logging_character -- both optional and both additive.

The repository workspace gains ./companion; the command is idempotent,
so it does not matter whether the api plan created go.work first." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 2: The API client — envelope, device token, retries and backoff

**Files:**
- Create: `companion/internal/client/client.go`
- Test: `companion/internal/client/client_test.go`

**Interfaces:**
- Consumes: nothing from earlier tasks; this package is standalone until Task 3 gives it routes.
- Produces:
  - `client.Envelope{OK bool; Data json.RawMessage; Error *ErrorBody; RequestID string}`, `client.ErrorBody{Message string; Fields map[string]string}`.
  - `client.Error{Status int; Message string; Fields map[string]string; RequestID string; RetryAfter time.Duration}` with `.Error() string` and `.Retryable() bool`.
  - `client.Retryable(error) bool`, `client.Unauthorized(error) bool`.
  - `client.Retry{MaxAttempts int; Base, Max time.Duration}`, `client.DefaultRetry() Retry`, `Retry.Delay(attempt int, frac float64) time.Duration`.
  - `client.Options{BaseURL string; HTTP *http.Client; Token func() string; Retry Retry; Log *slog.Logger; Sleep func(context.Context, time.Duration) error; Frac func() float64}`, `client.New(Options) (*Client, error)`.
  - Unexported to the package but used by every later route: `request{Method, Path string; Query url.Values; Body []byte; Type string; Header map[string]string; Anonymous bool; Out any}` and `(*Client).do(ctx, request) (int, error)`.
  - `client.MaxErrorBody`.

Three decisions are worth stating because later tasks lean on them. Bodies are `[]byte`, not readers, so a retry resends the same bytes without the caller rewinding anything and the queue can store a body verbatim. `do` returns the HTTP status as well as the error, so a caller can tell 201 (stored) from 200 (already stored) — the contract makes that distinction on three routes. And `Token` is a function, not a string, because pairing can happen while the uploader is running.

- [ ] **Step 1: Write the failing test**

```go
package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// newTestClient builds a client with no real waiting and no randomness,
// recording what it slept for.
func newTestClient(t *testing.T, base string, slept *[]time.Duration) *Client {
	t.Helper()
	c, err := New(Options{
		BaseURL: base,
		Token:   func() string { return "fsd_test" },
		Retry:   Retry{MaxAttempts: 4, Base: time.Second, Max: 8 * time.Second},
		Frac:    func() float64 { return 1 },
		Sleep: func(_ context.Context, d time.Duration) error {
			*slept = append(*slept, d)
			return nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestTheEnvelopeIsUnwrappedIntoTheOutValue(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer fsd_test" {
			t.Errorf("Authorization = %q", got)
		}
		w.Write([]byte(`{"ok":true,"data":{"id":"abc123"},"error":null,"request_id":"r1"}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	var out struct {
		ID string `json:"id"`
	}
	if _, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x", Out: &out}); err != nil {
		t.Fatal(err)
	}
	if out.ID != "abc123" {
		t.Fatalf("id = %q", out.ID)
	}
}

func TestAFailedEnvelopeBecomesATypedError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"ok":false,"data":null,"error":{"message":"verification failed",` +
			`"fields":{"metrics":"dps differs by 4%"}},"request_id":"r2"}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	_, err := c.do(t.Context(), request{Method: http.MethodPut, Path: "/v1/x"})
	var ae *Error
	if !errors.As(err, &ae) {
		t.Fatalf("err = %v, want *Error", err)
	}
	if ae.Status != http.StatusConflict || ae.Fields["metrics"] != "dps differs by 4%" || ae.RequestID != "r2" {
		t.Fatalf("error = %+v", ae)
	}
	if ae.Retryable() {
		t.Error("a 409 must not be retried")
	}
	if len(slept) != 0 {
		t.Errorf("slept %v on a 409", slept)
	}
}

func TestServerErrorsAreRetriedWithExponentialBackoff(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits.Add(1) < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ok":false,"error":{"message":"boom"}}`))
			return
		}
		w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	if _, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"}); err != nil {
		t.Fatal(err)
	}
	if hits.Load() != 3 {
		t.Fatalf("server saw %d requests, want 3", hits.Load())
	}
	want := []time.Duration{time.Second, 2 * time.Second}
	if len(slept) != len(want) || slept[0] != want[0] || slept[1] != want[1] {
		t.Fatalf("slept %v, want %v", slept, want)
	}
}

func TestTheBackoffIsCappedAndJittered(t *testing.T) {
	r := Retry{MaxAttempts: 8, Base: time.Second, Max: 10 * time.Second}
	for _, tc := range []struct {
		attempt int
		frac    float64
		want    time.Duration
	}{
		{1, 1, time.Second},
		{2, 1, 2 * time.Second},
		{5, 1, 10 * time.Second},
		{9, 1, 10 * time.Second},
		{3, 0.5, 2 * time.Second},
	} {
		if got := r.Delay(tc.attempt, tc.frac); got != tc.want {
			t.Errorf("Delay(%d, %v) = %v, want %v", tc.attempt, tc.frac, got, tc.want)
		}
	}
}

func TestRetryAfterOverridesAShorterBackoff(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits.Add(1) == 1 {
			w.Header().Set("Retry-After", "7")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"ok":false,"error":{"message":"slow down"}}`))
			return
		}
		w.Write([]byte(`{"ok":true,"data":null}`))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	if _, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"}); err != nil {
		t.Fatal(err)
	}
	if len(slept) != 1 || slept[0] != 7*time.Second {
		t.Fatalf("slept %v, want [7s]", slept)
	}
}

func TestAnUnpairedDeviceFailsWithoutSendingAnything(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("the client sent a request without a token")
	}))
	defer srv.Close()
	c, err := New(Options{BaseURL: srv.URL, Token: func() string { return "" }})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"})
	if !Unauthorized(err) {
		t.Fatalf("err = %v, want an unauthorized error", err)
	}
}

func TestANonEnvelopeErrorBodyIsKeptAsASnippet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<html>gateway</html>"))
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	_, err := c.do(t.Context(), request{Method: http.MethodGet, Path: "/v1/x"})
	var ae *Error
	if !errors.As(err, &ae) || ae.Status != http.StatusBadGateway || ae.Message != "<html>gateway</html>" {
		t.Fatalf("err = %v", err)
	}
}

func TestAnEmptyBaseURLIsRefused(t *testing.T) {
	if _, err := New(Options{}); err == nil {
		t.Fatal("New accepted an empty base URL")
	}
}

func TestACancelledContextIsNotRetried(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	var slept []time.Duration
	c := newTestClient(t, srv.URL, &slept)
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if _, err := c.do(ctx, request{Method: http.MethodGet, Path: "/v1/x"}); err == nil {
		t.Fatal("a cancelled request succeeded")
	}
	if len(slept) != 0 {
		t.Errorf("slept %v on a cancelled context", slept)
	}
}

func TestErrorMessagesNameTheRequestWhenTheServerGaveOne(t *testing.T) {
	with := (&Error{Status: 409, Message: "nope", RequestID: "r9"}).Error()
	if !strings.Contains(with, "409") || !strings.Contains(with, "r9") {
		t.Errorf("with a request id: %q", with)
	}
	without := (&Error{Status: 500, Message: "boom"}).Error()
	if strings.Contains(without, "request") {
		t.Errorf("without a request id: %q", without)
	}
}

func TestRetryableClassifiesFailures(t *testing.T) {
	if Retryable(nil) {
		t.Error("nil is retryable")
	}
	if Retryable(context.Canceled) {
		t.Error("a cancelled context is retryable")
	}
	if !Retryable(errors.New("connection reset by peer")) {
		t.Error("a transport error is not retryable")
	}
	for status, want := range map[int]bool{
		408: true, 429: true, 500: true, 503: true,
		400: false, 401: false, 404: false, 409: false, 413: false,
	} {
		if got := (&Error{Status: status}).Retryable(); got != want {
			t.Errorf("%d retryable = %v, want %v", status, got, want)
		}
	}
	if !Unauthorized(&Error{Status: 403}) || Unauthorized(&Error{Status: 409}) {
		t.Error("Unauthorized misclassified a status")
	}
}

func TestDefaultsAreFilledIn(t *testing.T) {
	c, err := New(Options{BaseURL: "https://api.foreversixty.gg/"})
	if err != nil {
		t.Fatal(err)
	}
	if c.retry != DefaultRetry() || c.token() != "" || c.frac == nil || c.sleep == nil || c.log == nil {
		t.Fatalf("defaults = %+v", c.retry)
	}
	if c.base.String() != "https://api.foreversixty.gg" {
		t.Errorf("base = %q", c.base.String())
	}
	if _, err := New(Options{BaseURL: "://nonsense"}); err == nil {
		t.Error("an unparseable base URL was accepted")
	}
}

func TestSleepHonoursCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := sleep(ctx, time.Hour); err == nil {
		t.Fatal("sleep ignored a cancelled context")
	}
	if err := sleep(context.Background(), time.Millisecond); err != nil {
		t.Fatal(err)
	}
}

func TestSnippetTrimsAndDescribesAnEmptyBody(t *testing.T) {
	if got := snippet([]byte("   ")); got != "empty response" {
		t.Errorf("empty = %q", got)
	}
	long := snippet([]byte(strings.Repeat("x", MaxErrorBody+50)))
	if len(long) != MaxErrorBody+len("…") {
		t.Errorf("long snippet is %d characters", len(long))
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/client/`
Expected: FAIL, "undefined: New", "undefined: Error", "undefined: Retry".

- [ ] **Step 3: Write the implementation**

```go
// companion/internal/client/client.go
// Package client is the companion's side of the ingest contract: the
// envelope, the device-token header, and a retry policy that knows
// which failures are worth repeating. Every request body is a byte
// slice so a retry can send it again without the caller rewinding
// anything.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// MaxErrorBody is how much of an unparseable error response is kept for
// the message. A server that answers HTML must not fill the log file.
const MaxErrorBody = 2048

// Envelope is the API's response shape for every route.
type Envelope struct {
	OK        bool            `json:"ok"`
	Data      json.RawMessage `json:"data"`
	Error     *ErrorBody      `json:"error"`
	RequestID string          `json:"request_id"`
}

// ErrorBody is the envelope's error member.
type ErrorBody struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

// Error is one failed request. The status and the field errors are kept
// because the pipeline reacts differently to a verification mismatch
// than to a dead network.
type Error struct {
	Status    int
	Message   string
	Fields    map[string]string
	RequestID string
	// RetryAfter is the server's own hint, from the Retry-After
	// header. Zero means it sent none.
	RetryAfter time.Duration
}

func (e *Error) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("api: %d %s (request %s)", e.Status, e.Message, e.RequestID)
	}
	return fmt.Sprintf("api: %d %s", e.Status, e.Message)
}

// Retryable reports whether sending the same request again could
// succeed. A 4xx other than 408 and 429 is the companion's own fault
// and repeating it only burns the queue.
func (e *Error) Retryable() bool {
	return e.Status == http.StatusRequestTimeout ||
		e.Status == http.StatusTooManyRequests ||
		e.Status >= 500
}

// Retryable reports whether err is worth another attempt. Transport
// errors always are: the player closed a laptop lid, not a contract.
func Retryable(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var ae *Error
	if errors.As(err, &ae) {
		return ae.Retryable()
	}
	return true
}

// Unauthorized reports whether err says the device token is gone, which
// is the one failure that must stop the uploader and ask the player to
// pair again rather than retry forever.
func Unauthorized(err error) bool {
	var ae *Error
	return errors.As(err, &ae) &&
		(ae.Status == http.StatusUnauthorized || ae.Status == http.StatusForbidden)
}

// Retry is the backoff policy: exponential from Base, capped at Max,
// with full jitter so a thousand companions reconnecting after an
// outage do not arrive together.
type Retry struct {
	MaxAttempts int
	Base        time.Duration
	Max         time.Duration
}

// DefaultRetry is five attempts from one second up to thirty.
func DefaultRetry() Retry {
	return Retry{MaxAttempts: 5, Base: time.Second, Max: 30 * time.Second}
}

// Delay is how long to wait before attempt n, counting from 1.
func (r Retry) Delay(attempt int, frac float64) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	d := float64(r.Base) * math.Pow(2, float64(attempt-1))
	if d > float64(r.Max) {
		d = float64(r.Max)
	}
	return time.Duration(d * frac)
}

// Options configures a Client.
type Options struct {
	BaseURL string
	HTTP    *http.Client
	// Token returns the device token, or the empty string when the
	// companion is not paired. It is a function because pairing can
	// happen while the uploader is running.
	Token func() string
	Retry Retry
	Log   *slog.Logger
	// Sleep waits, honouring cancellation. Tests replace it so a
	// backoff test does not take thirty seconds.
	Sleep func(ctx context.Context, d time.Duration) error
	// Frac returns the jitter fraction in [0,1). Tests pin it.
	Frac func() float64
}

// Client talks to the API.
type Client struct {
	base  *url.URL
	http  *http.Client
	token func() string
	retry Retry
	log   *slog.Logger
	sleep func(ctx context.Context, d time.Duration) error
	frac  func() float64
}

// New builds a client. An unparseable base URL is a configuration error
// and is reported now rather than on the first upload.
func New(o Options) (*Client, error) {
	if o.BaseURL == "" {
		return nil, errors.New("client: base URL is empty")
	}
	u, err := url.Parse(strings.TrimRight(o.BaseURL, "/"))
	if err != nil {
		return nil, fmt.Errorf("client: base URL: %w", err)
	}
	c := &Client{
		base:  u,
		http:  o.HTTP,
		token: o.Token,
		retry: o.Retry,
		log:   o.Log,
		sleep: o.Sleep,
		frac:  o.Frac,
	}
	if c.http == nil {
		c.http = &http.Client{Timeout: 2 * time.Minute}
	}
	if c.token == nil {
		c.token = func() string { return "" }
	}
	if c.retry.MaxAttempts == 0 {
		c.retry = DefaultRetry()
	}
	if c.log == nil {
		c.log = slog.New(slog.DiscardHandler)
	}
	if c.sleep == nil {
		c.sleep = sleep
	}
	if c.frac == nil {
		c.frac = rand.Float64
	}
	return c, nil
}

func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// request is one call to the API.
type request struct {
	Method string
	Path   string // beginning with a slash, already escaped
	Query  url.Values
	Body   []byte
	Type   string            // Content-Type, empty for no body
	Header map[string]string // extra headers
	// Anonymous skips the Authorization header. Only pairing uses it:
	// a device has no token until the claim succeeds.
	Anonymous bool
	// Out receives the envelope's data member when it is not nil.
	Out any
}

// do sends a request, retrying the failures worth retrying, and decodes
// the envelope. The returned status is the last response's, so a caller
// can tell 201 (stored) from 200 (already stored).
func (c *Client) do(ctx context.Context, r request) (int, error) {
	var lastErr error
	for attempt := 1; attempt <= c.retry.MaxAttempts; attempt++ {
		status, err := c.attempt(ctx, r)
		if err == nil {
			return status, nil
		}
		lastErr = err
		if !Retryable(err) || attempt == c.retry.MaxAttempts {
			return status, err
		}
		d := c.retry.Delay(attempt, c.frac())
		var ae *Error
		if errors.As(err, &ae) && ae.RetryAfter > d {
			d = ae.RetryAfter
		}
		c.log.Warn("retrying", "component", "client", "method", r.Method, "path", r.Path,
			"attempt", attempt, "in", d.String(), "err", err.Error())
		if serr := c.sleep(ctx, d); serr != nil {
			return status, errors.Join(err, serr)
		}
	}
	return 0, lastErr
}

func (c *Client) attempt(ctx context.Context, r request) (int, error) {
	u := *c.base
	u.Path = c.base.Path + r.Path
	if len(r.Query) > 0 {
		u.RawQuery = r.Query.Encode()
	}
	var body io.Reader
	if r.Body != nil {
		body = bytes.NewReader(r.Body)
	}
	req, err := http.NewRequestWithContext(ctx, r.Method, u.String(), body)
	if err != nil {
		return 0, err
	}
	if r.Type != "" {
		req.Header.Set("Content-Type", r.Type)
	}
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Accept", "application/json")
	if !r.Anonymous {
		tok := c.token()
		if tok == "" {
			return 0, &Error{Status: http.StatusUnauthorized, Message: "this device is not paired"}
		}
		req.Header.Set("Authorization", "Bearer "+tok)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%s %s: %w", r.Method, r.Path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		io.Copy(io.Discard, resp.Body)
		return resp.StatusCode, nil
	}
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%s %s: read body: %w", r.Method, r.Path, err)
	}
	var env Envelope
	if jerr := json.Unmarshal(raw, &env); jerr != nil || (!env.OK && env.Error == nil) {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp.StatusCode, nil // a 2xx with no envelope is success
		}
		after, _ := retryAfter(resp.Header)
		return resp.StatusCode, &Error{
			Status: resp.StatusCode, Message: snippet(raw), RetryAfter: after}
	}
	if !env.OK {
		after, _ := retryAfter(resp.Header)
		return resp.StatusCode, &Error{
			Status:     resp.StatusCode,
			Message:    env.Error.Message,
			Fields:     env.Error.Fields,
			RequestID:  env.RequestID,
			RetryAfter: after,
		}
	}
	if r.Out != nil && len(env.Data) > 0 {
		if err := json.Unmarshal(env.Data, r.Out); err != nil {
			return resp.StatusCode, fmt.Errorf("%s %s: decode data: %w", r.Method, r.Path, err)
		}
	}
	return resp.StatusCode, nil
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	if s == "" {
		return "empty response"
	}
	if len(s) > MaxErrorBody {
		s = s[:MaxErrorBody] + "…"
	}
	return s
}

// retryAfter is the server's own backoff hint, honoured over ours when
// it is longer. Only seconds are supported; the API sends nothing else.
func retryAfter(h http.Header) (time.Duration, bool) {
	v := h.Get("Retry-After")
	if v == "" {
		return 0, false
	}
	n, err := strconv.Atoi(v)
	if err != nil || n < 0 {
		return 0, false
	}
	return time.Duration(n) * time.Second, true
}
```

- [ ] **Step 4: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/client/ && go test ./internal/client/ -race`
Expected: PASS. The backoff test asserts the exact sleeps (1s then 2s with the jitter fraction pinned to 1), so a change to the policy fails loudly.

- [ ] **Step 5: Commit**

```bash
git add companion/internal/client/client.go \
  companion/internal/client/client_test.go
git commit -m "feat(companion): the API client transport, with retries that know when to stop" \
  -m "One place that knows the envelope, the Bearer header and which
failures are worth repeating. A 4xx other than 408 and 429 is the
companion's own fault and repeating it only burns the queue; a
transport error always deserves another attempt, because the player
shut a laptop lid rather than broke a contract.

Backoff is exponential from one second to thirty with full jitter, so
a thousand companions reconnecting after an outage do not arrive
together, and the server's own Retry-After wins when it is longer.
Request bodies are byte slices so a retry resends exactly what failed." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Task 3: The ingest routes and the fake API

**Files:**
- Create: `companion/internal/client/ingest.go`, `companion/internal/fakeapi/fakeapi.go`
- Test: `companion/internal/client/ingest_test.go`, `companion/internal/fakeapi/fakeapi_test.go`

**Interfaces:**
- Consumes: `client.New`, `client.Options`, `(*Client).do`, `client.Error` (Task 2); `character.Character` (Task 1); `summary.Summary`, `summary.RosterRow`, `fight.Fight`, `session.Health` from the engine.
- Produces:
  - `client.Character` (an alias for `character.Character`), `client.CreateReport{Title, Visibility, Zone string; LoggingCharacter *Character}`, `client.Report{ID string; CreatedAt time.Time}`.
  - `client.MetricsRow` with the contract's sixteen json names, and `client.MetricsRowsOf(fight.Fight, summary.Summary) []MetricsRow`.
  - `client.RawRange{StartOffset, EndOffset int64; SHA256 string}`, `client.FightBundle{Summary summary.Summary; Events []byte; Metrics []MetricsRow; RawRange RawRange}` with `.Encode() (contentType string, body []byte, err error)`, and `client.BundleBoundary`.
  - `client.Live{Summary summary.Summary; ElapsedMS int64; UpdatedAt time.Time}`, `client.Complete{FinalOffset int64; EngineVersion string; Health session.Health}`.
  - `client.Stored` with `client.Created` and `client.Duplicate`.
  - `(*Client).CreateReport`, `.PutFight(ctx, reportID string, index int, contentType string, body []byte) (Stored, error)`, `.PutLive`, `.PutRaw(ctx, reportID string, offset int64, chunk []byte) (Stored, error)`, `.Complete`.
  - `fakeapi.New() *Server` with `Token`, `PairCode`, `.Offline(bool)`, `.RefuseFights(bool)`, `.FailNext(pattern string, n int)`, `.Order() []string`, `.Reports() map[string]*Report`, `.Exports()`, `.SetInbox(string)`; `fakeapi.Fight`, `fakeapi.Report`.

`PutFight` takes a pre-encoded body because the queue in Task 5 stores exactly those bytes and replays them unchanged. The multipart boundary is pinned so two encodings of the same bundle are byte-identical — the same determinism rule the engine works under.

`fakeapi` is the interface contract written down once, in memory. Every companion test that talks to a server talks to it, so a route that drifts breaks every test at the same time.

- [ ] **Step 1: Write the fake API**

```go
// companion/internal/fakeapi/fakeapi.go
// Package fakeapi is the ingest contract implemented in memory, over
// httptest. Every companion test that talks to a server talks to this
// one, so the contract is written down once and a route that drifts
// breaks every test at the same time.
package fakeapi

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// Token is the device token the fake issues and the only one it
// accepts.
const Token = "fsd_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// PairCode is the pairing code the fake accepts.
const PairCode = "HORDE-42"

// Fight is one stored fight bundle, decoded.
type Fight struct {
	Index    int
	Summary  json.RawMessage
	Events   []byte
	Metrics  []client.MetricsRow
	RawRange client.RawRange
}

// Report is one stored report.
type Report struct {
	ID         string
	Visibility string
	Title      string
	Fights     map[int]Fight
	Live       map[int]client.Live
	Raw        map[int64][]byte
	Complete   *client.Complete
}

// Server is a fake ingest API.
type Server struct {
	*httptest.Server

	mu       sync.Mutex
	reports  map[string]*Report
	nextID   int
	order    []string
	offline  bool
	failNext map[string]int
	exports  []map[string]string
	inbox    json.RawMessage
	refuse   bool
}

// New starts a fake API. Close it with Close.
func New() *Server {
	s := &Server{
		reports:  map[string]*Report{},
		failNext: map[string]int{},
		inbox:    json.RawMessage(`{"builds":[]}`),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/devices/claim", s.claim)
	mux.HandleFunc("POST /v1/reports", s.createReport)
	mux.HandleFunc("PUT /v1/reports/{id}/fights/{n}", s.putFight)
	mux.HandleFunc("PUT /v1/reports/{id}/fights/{n}/live", s.putLive)
	mux.HandleFunc("PUT /v1/reports/{id}/raw", s.putRaw)
	mux.HandleFunc("POST /v1/reports/{id}/complete", s.complete)
	mux.HandleFunc("POST /v1/addon/exports", s.addonExports)
	mux.HandleFunc("GET /v1/addon/inbox", s.addonInbox)
	s.Server = httptest.NewServer(s.wrap(mux))
	return s
}

// Offline makes every route answer 503, standing in for a dropped
// network without tearing the listener down.
func (s *Server) Offline(v bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.offline = v
}

// RefuseFights makes every fight PUT answer 409, standing in for the
// verification mismatch the contract describes. It is permanent, not
// retryable, which is the behaviour under test.
func (s *Server) RefuseFights(v bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.refuse = v
}

// FailNext makes the next n requests to a route pattern
// ("PUT /v1/reports/{id}/fights/{n}") answer 500.
func (s *Server) FailNext(pattern string, n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.failNext[pattern] = n
}

// Order is every write the server accepted, in arrival order, as
// "fight 3", "raw 4194304", "complete" and so on. It is what an
// ordering assertion reads.
func (s *Server) Order() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.order...)
}

// Reports returns a snapshot of everything stored.
func (s *Server) Reports() map[string]*Report {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := map[string]*Report{}
	for k, v := range s.reports {
		out[k] = v
	}
	return out
}

// Exports is every addon export the companion posted.
func (s *Server) Exports() []map[string]string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]map[string]string(nil), s.exports...)
}

// SetInbox replaces the body GET /v1/addon/inbox answers with.
func (s *Server) SetInbox(raw string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.inbox = json.RawMessage(raw)
}

func (s *Server) wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		off := s.offline
		s.mu.Unlock()
		if off {
			fail(w, http.StatusServiceUnavailable, "the network is down")
			return
		}
		if r.URL.Path != "/v1/devices/claim" {
			if r.Header.Get("Authorization") != "Bearer "+Token {
				fail(w, http.StatusUnauthorized, "unknown device token")
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// shouldFail consumes one injected failure for this pattern.
func (s *Server) shouldFail(pattern string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.failNext[pattern] > 0 {
		s.failNext[pattern]--
		return true
	}
	return false
}

func ok(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"ok": true, "data": data, "error": nil, "request_id": "fake",
	})
}

func fail(w http.ResponseWriter, status int, msg string, fields ...map[string]string) {
	body := map[string]any{"message": msg}
	if len(fields) > 0 {
		body["fields"] = fields[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"ok": false, "data": nil, "error": body, "request_id": "fake",
	})
}

func (s *Server) claim(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Platform string `json:"platform"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	if in.Code != PairCode {
		fail(w, http.StatusBadRequest, "that pairing code is not valid",
			map[string]string{"code": "expired or unknown"})
		return
	}
	ok(w, http.StatusCreated, map[string]string{
		"device_id": "dev00000000aaaa", "token": Token,
	})
}

func (s *Server) createReport(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("POST /v1/reports") {
		fail(w, http.StatusInternalServerError, "injected failure")
		return
	}
	var in client.CreateReport
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	s.nextID++
	id := fmt.Sprintf("rpt%09d", s.nextID)
	s.reports[id] = &Report{ID: id, Visibility: in.Visibility, Title: in.Title,
		Fights: map[int]Fight{}, Live: map[int]client.Live{}, Raw: map[int64][]byte{}}
	s.order = append(s.order, "report "+id)
	s.mu.Unlock()
	ok(w, http.StatusCreated, map[string]any{"id": id, "created_at": "2026-12-09T20:00:00Z"})
}

func (s *Server) report(w http.ResponseWriter, r *http.Request) (*Report, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rep, found := s.reports[r.PathValue("id")]
	if !found {
		fail(w, http.StatusNotFound, "no such report")
		return nil, false
	}
	return rep, true
}

func (s *Server) putFight(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("PUT /v1/reports/{id}/fights/{n}") {
		fail(w, http.StatusInternalServerError, "injected failure")
		return
	}
	s.mu.Lock()
	refuse := s.refuse
	s.mu.Unlock()
	if refuse {
		fail(w, http.StatusConflict, "the metrics do not match the events",
			map[string]string{"metrics": "dps differs by more than 0.5%"})
		return
	}
	rep, found := s.report(w, r)
	if !found {
		return
	}
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		fail(w, http.StatusBadRequest, "fight index is not a number")
		return
	}
	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		fail(w, http.StatusBadRequest, "content type: "+err.Error())
		return
	}
	f := Fight{Index: n}
	mr := multipart.NewReader(r.Body, params["boundary"])
	seen := map[string]bool{}
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fail(w, http.StatusBadRequest, "multipart: "+err.Error())
			return
		}
		b, err := io.ReadAll(part)
		if err != nil {
			fail(w, http.StatusBadRequest, "multipart: "+err.Error())
			return
		}
		seen[part.FormName()] = true
		switch part.FormName() {
		case "summary":
			f.Summary = json.RawMessage(b)
		case "events":
			f.Events = b
		case "metrics":
			if err := json.Unmarshal(b, &f.Metrics); err != nil {
				fail(w, http.StatusBadRequest, "metrics: "+err.Error())
				return
			}
		case "raw_range":
			if err := json.Unmarshal(b, &f.RawRange); err != nil {
				fail(w, http.StatusBadRequest, "raw_range: "+err.Error())
				return
			}
		}
	}
	for _, want := range []string{"summary", "events", "metrics", "raw_range"} {
		if !seen[want] {
			fail(w, http.StatusBadRequest, "the bundle is missing the "+want+" part")
			return
		}
	}
	s.mu.Lock()
	prior, already := rep.Fights[n]
	if already && prior.RawRange.SHA256 == f.RawRange.SHA256 {
		s.mu.Unlock()
		ok(w, http.StatusOK, map[string]any{"fight_index": n, "verified": true})
		return
	}
	rep.Fights[n] = f
	s.order = append(s.order, "fight "+strconv.Itoa(n))
	s.mu.Unlock()
	ok(w, http.StatusCreated, map[string]any{"fight_index": n, "verified": true})
}

func (s *Server) putLive(w http.ResponseWriter, r *http.Request) {
	rep, found := s.report(w, r)
	if !found {
		return
	}
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		fail(w, http.StatusBadRequest, "fight index is not a number")
		return
	}
	var in client.Live
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	rep.Live[n] = in
	s.order = append(s.order, "live "+strconv.Itoa(n))
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) putRaw(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("PUT /v1/reports/{id}/raw") {
		fail(w, http.StatusInternalServerError, "injected failure")
		return
	}
	rep, found := s.report(w, r)
	if !found {
		return
	}
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		fail(w, http.StatusBadRequest, "offset is not a number")
		return
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 8<<20+1))
	if err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(body) > 8<<20 {
		fail(w, http.StatusRequestEntityTooLarge, "a raw chunk may not exceed 8 MiB")
		return
	}
	sum := sha256.Sum256(body)
	if got, want := r.Header.Get("X-Raw-SHA256"), hex.EncodeToString(sum[:]); got != want {
		fail(w, http.StatusBadRequest, "X-Raw-SHA256 does not match the body")
		return
	}
	s.mu.Lock()
	if prior, already := rep.Raw[offset]; already {
		same := string(prior) == string(body)
		s.mu.Unlock()
		if !same {
			fail(w, http.StatusConflict, "that offset holds different bytes")
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	rep.Raw[offset] = body
	s.order = append(s.order, "raw "+strconv.FormatInt(offset, 10))
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) complete(w http.ResponseWriter, r *http.Request) {
	rep, found := s.report(w, r)
	if !found {
		return
	}
	var in client.Complete
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	rep.Complete = &in
	s.order = append(s.order, "complete "+rep.ID)
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) addonExports(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Characters []map[string]string `json:"characters"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		fail(w, http.StatusBadRequest, "malformed body")
		return
	}
	s.mu.Lock()
	s.exports = append(s.exports, in.Characters...)
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) addonInbox(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	body := s.inbox
	s.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"ok":true,"data":%s,"error":null,"request_id":"fake"}`, strings.TrimSpace(string(body)))
}
```

- [ ] **Step 2: Write the fake API's own test**

```go
package fakeapi

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func get(t *testing.T, url, token string) (int, string) {
	t.Helper()
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(b)
}

func TestEveryRouteNeedsTheDeviceTokenExceptTheClaim(t *testing.T) {
	s := New()
	defer s.Close()
	if code, _ := get(t, s.URL+"/v1/addon/inbox", ""); code != http.StatusUnauthorized {
		t.Errorf("inbox without a token = %d", code)
	}
	if code, body := get(t, s.URL+"/v1/addon/inbox", Token); code != http.StatusOK ||
		!strings.Contains(body, `"builds"`) {
		t.Errorf("inbox = %d %s", code, body)
	}
}

func TestOfflineFailsEveryRoute(t *testing.T) {
	s := New()
	defer s.Close()
	s.Offline(true)
	if code, _ := get(t, s.URL+"/v1/addon/inbox", Token); code != http.StatusServiceUnavailable {
		t.Errorf("offline inbox = %d", code)
	}
	s.Offline(false)
	if code, _ := get(t, s.URL+"/v1/addon/inbox", Token); code != http.StatusOK {
		t.Errorf("back online = %d", code)
	}
}

func TestFailNextIsConsumedOnce(t *testing.T) {
	s := New()
	defer s.Close()
	s.FailNext("POST /v1/reports", 1)
	if !s.shouldFail("POST /v1/reports") {
		t.Fatal("the first call did not fail")
	}
	if s.shouldFail("POST /v1/reports") {
		t.Fatal("the failure was not consumed")
	}
}

func TestSetInboxReplacesTheBody(t *testing.T) {
	s := New()
	defer s.Close()
	s.SetInbox(`{"builds":[{"id":"b1","name":"Holy","code":"FSB1:x"}]}`)
	_, body := get(t, s.URL+"/v1/addon/inbox", Token)
	if !strings.Contains(body, "FSB1:x") {
		t.Fatalf("inbox = %s", body)
	}
	if len(s.Order()) != 0 || len(s.Reports()) != 0 || len(s.Exports()) != 0 {
		t.Error("a read changed the recorded state")
	}
}
```

- [ ] **Step 3: Write the failing ingest test**

```go
package client_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func dial(t *testing.T, srv *fakeapi.Server) *client.Client {
	t.Helper()
	c, err := client.New(client.Options{
		BaseURL: srv.URL,
		Token:   func() string { return fakeapi.Token },
		Retry:   client.Retry{MaxAttempts: 3, Base: time.Millisecond, Max: time.Millisecond},
		Frac:    func() float64 { return 1 },
	})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func sampleFight() (fight.Fight, summary.Summary) {
	f := fight.Fight{
		Index: 2, Kind: fight.Encounter, Name: "Warden Kelthas", EncounterID: 9001,
		Difficulty: 14, Size: 20, Kill: true,
		Start: time.Unix(1000, 0).UTC(), End: time.Unix(1240, 0).UTC(),
		StartOffset: 4096, EndOffset: 90112,
	}
	s := summary.Summary{
		EngineVersion: session.Version, FightIndex: 2, DurationMS: 240000,
		Roster: []summary.RosterRow{
			{GUID: "Player-4184-000000A1", Name: "Morrowlyn", Class: "Paladin", Spec: "Holy",
				Role: "healer", ItemLevel: 183, ActiveMS: 200000, Deaths: 1,
				DamageTaken: 40123, DPS: 120.5, HPS: 980.25},
			{GUID: "Player-4184-000000A2", Name: "Brannic", Class: "Warrior", Spec: "Fury",
				Role: "dps", ItemLevel: 176, ActiveMS: 230000, Deaths: 0,
				DamageTaken: 91002, DPS: 1420.75, HPS: 0},
		},
	}
	return f, s
}

func TestMetricsRowsArePivotedFromTheRoster(t *testing.T) {
	f, s := sampleFight()
	rows := client.MetricsRowsOf(f, s)
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want 2", len(rows))
	}
	first := rows[0]
	if first.PlayerGUID != "Player-4184-000000A1" || first.Name != "Morrowlyn" ||
		first.Class != "Paladin" || first.Spec != "Holy" || first.Role != "healer" ||
		first.ItemLevel != 183 || first.MetricDPS != 120.5 || first.MetricHPS != 980.25 ||
		first.DamageTaken != 40123 || first.ActiveMS != 200000 || first.Deaths != 1 {
		t.Errorf("row = %+v", first)
	}
	if first.EncounterID != 9001 || first.Difficulty != 14 || first.Size != 20 ||
		first.DurationMS != 240000 || !first.Kill {
		t.Errorf("fight fields = %+v", first)
	}
}

func TestAFightBundleRoundTripsThroughTheMultipartRoute(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)

	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "public", Zone: "Molten Core"})
	if err != nil {
		t.Fatal(err)
	}
	f, s := sampleFight()
	b := client.FightBundle{
		Summary: s,
		Events:  []byte("PAR1fake"),
		Metrics: client.MetricsRowsOf(f, s),
		RawRange: client.RawRange{StartOffset: f.StartOffset, EndOffset: f.EndOffset,
			SHA256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}
	ct, body, err := b.Encode()
	if err != nil {
		t.Fatal(err)
	}
	stored, err := c.PutFight(t.Context(), rep.ID, 2, ct, body)
	if err != nil {
		t.Fatal(err)
	}
	if stored != client.Created {
		t.Errorf("first PUT = %v, want Created", stored)
	}
	again, err := c.PutFight(t.Context(), rep.ID, 2, ct, body)
	if err != nil {
		t.Fatal(err)
	}
	if again != client.Duplicate {
		t.Errorf("re-sent PUT = %v, want Duplicate", again)
	}

	got := srv.Reports()[rep.ID].Fights[2]
	if string(got.Events) != "PAR1fake" {
		t.Errorf("events = %q", got.Events)
	}
	if len(got.Metrics) != 2 || got.Metrics[1].Name != "Brannic" {
		t.Errorf("metrics = %+v", got.Metrics)
	}
	if got.RawRange.StartOffset != 4096 || got.RawRange.EndOffset != 90112 {
		t.Errorf("raw range = %+v", got.RawRange)
	}
	if want := []string{"report " + rep.ID, "fight 2"}; len(srv.Order()) != len(want) {
		t.Errorf("order = %v, want %v", srv.Order(), want)
	}
}

func TestEncodingABundleIsDeterministic(t *testing.T) {
	f, s := sampleFight()
	b := client.FightBundle{Summary: s, Events: []byte("PAR1"), Metrics: client.MetricsRowsOf(f, s)}
	_, first, err := b.Encode()
	if err != nil {
		t.Fatal(err)
	}
	_, second, err := b.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("two encodings of the same bundle differ")
	}
}

func TestLiveRawAndCompleteReachTheServer(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)
	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "unlisted"})
	if err != nil {
		t.Fatal(err)
	}
	_, s := sampleFight()
	if err := c.PutLive(t.Context(), rep.ID, 2, client.Live{
		Summary: s, ElapsedMS: 12000, UpdatedAt: time.Unix(1200, 0).UTC()}); err != nil {
		t.Fatal(err)
	}
	if stored, err := c.PutRaw(t.Context(), rep.ID, 4194304, []byte("compressed")); err != nil || stored != client.Created {
		t.Fatalf("PutRaw = %v, %v", stored, err)
	}
	if stored, err := c.PutRaw(t.Context(), rep.ID, 4194304, []byte("compressed")); err != nil || stored != client.Duplicate {
		t.Fatalf("re-sent PutRaw = %v, %v", stored, err)
	}
	if err := c.Complete(t.Context(), rep.ID, client.Complete{
		FinalOffset: 131072, EngineVersion: session.Version,
		Health: session.Health{Layout: "retail-v16", Lines: 4200}}); err != nil {
		t.Fatal(err)
	}
	stored := srv.Reports()[rep.ID]
	if stored.Live[2].ElapsedMS != 12000 {
		t.Errorf("live = %+v", stored.Live[2])
	}
	if string(stored.Raw[4194304]) != "compressed" {
		t.Errorf("raw = %q", stored.Raw[4194304])
	}
	if stored.Complete == nil || stored.Complete.FinalOffset != 131072 ||
		stored.Complete.Health.Lines != 4200 {
		t.Errorf("complete = %+v", stored.Complete)
	}
}

func TestARawOffsetThatHoldsDifferentBytesIsAConflict(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)
	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "public"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.PutRaw(t.Context(), rep.ID, 0, []byte("first")); err != nil {
		t.Fatal(err)
	}
	_, err = c.PutRaw(t.Context(), rep.ID, 0, []byte("second"))
	var ae *client.Error
	if !errors.As(err, &ae) || ae.Status != http.StatusConflict {
		t.Fatalf("err = %v, want a 409", err)
	}
}

func TestAnUnknownTokenIsReportedAsUnauthorized(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return "fsd_wrong" }})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.CreateReport(t.Context(), client.CreateReport{Visibility: "public"})
	if !client.Unauthorized(err) {
		t.Fatalf("err = %v, want unauthorized", err)
	}
}
```

- [ ] **Step 4: Run them to verify they fail**

Run: `cd companion && go test ./internal/client/ ./internal/fakeapi/`
Expected: both FAIL to build — `fakeapi` with "undefined: client.MetricsRow", the client test with "undefined: CreateReport".

- [ ] **Step 5: Resolve the engine dependency**

`ingest.go` is the first file to import the engine, so the module needs its checksums. The
`replace` in `go.mod` already points at `../logs`; this pulls in the engine's own dependencies
(`parquet-go`, `klauspost/compress`) and writes `go.sum`.

```bash
cd companion && go mod tidy && grep -c . go.sum
```

Expected: a non-zero count, and `go.mod` now requires `github.com/jhunthrop/foreversixty/logs`
and `github.com/klauspost/compress v1.20.0`.

- [ ] **Step 6: Write the implementation**

```go
// companion/internal/client/ingest.go
// The ingest routes from the Phase 3 interface contract, one method
// each. Bodies are built here and handed to the queue as bytes, so an
// upload that fails today is replayed byte for byte tomorrow.
package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Character identifies the character a report is logged from. It is
// the shared triple, so the ruleset field is spelled in one place.
type Character = character.Character

// CreateReport is the body of POST /v1/reports.
type CreateReport struct {
	Title            string     `json:"title,omitempty"`
	Visibility       string     `json:"visibility"`
	Zone             string     `json:"zone,omitempty"`
	LoggingCharacter *Character `json:"logging_character,omitempty"`
}

// Report is the API's answer to a report creation.
type Report struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// MetricsRow is one player's ranking row for one boss fight, with the
// field names the interface contract fixes. The engine's
// summary.MetricRow is one row per metric; the API wants one row per
// player, so MetricsRowsOf pivots the fight's roster instead.
type MetricsRow struct {
	PlayerGUID  string  `json:"player_guid"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Spec        string  `json:"spec"`
	Role        string  `json:"role"`
	ItemLevel   int64   `json:"ilvl"`
	MetricDPS   float64 `json:"metric_dps"`
	MetricHPS   float64 `json:"metric_hps"`
	DamageTaken int64   `json:"damage_taken"`
	ActiveMS    int64   `json:"active_ms"`
	Deaths      int     `json:"deaths"`
	EncounterID int64   `json:"encounter_id"`
	Difficulty  int64   `json:"difficulty"`
	Size        int64   `json:"size"`
	DurationMS  int64   `json:"duration_ms"`
	Kill        bool    `json:"kill"`
}

// MetricsRowsOf derives the fight's metrics rows from its roster. The
// order is the roster's, which the engine already sorts, so two runs
// over the same log post identical bytes.
func MetricsRowsOf(f fight.Fight, s summary.Summary) []MetricsRow {
	rows := make([]MetricsRow, 0, len(s.Roster))
	for _, r := range s.Roster {
		rows = append(rows, MetricsRow{
			PlayerGUID:  r.GUID,
			Name:        r.Name,
			Class:       r.Class,
			Spec:        r.Spec,
			Role:        r.Role,
			ItemLevel:   r.ItemLevel,
			MetricDPS:   r.DPS,
			MetricHPS:   r.HPS,
			DamageTaken: r.DamageTaken,
			ActiveMS:    r.ActiveMS,
			Deaths:      r.Deaths,
			EncounterID: f.EncounterID,
			Difficulty:  f.Difficulty,
			Size:        f.Size,
			DurationMS:  s.DurationMS,
			Kill:        f.Kill,
		})
	}
	return rows
}

// RawRange is the fight's byte range in the original log with the hash
// of those bytes, which the server's raw-sample check verifies later.
type RawRange struct {
	StartOffset int64  `json:"start_offset"`
	EndOffset   int64  `json:"end_offset"`
	SHA256      string `json:"sha256"`
}

// FightBundle is one closed fight, ready to be encoded once and stored
// in the queue until it lands.
type FightBundle struct {
	Summary  summary.Summary
	Events   []byte // Parquet
	Metrics  []MetricsRow
	RawRange RawRange
}

// BundleBoundary is pinned so an encoded bundle is a pure function of
// its content: the queue can compare two encodings byte for byte.
const BundleBoundary = "foreversixty-fight-bundle"

// Encode writes the multipart body the fight route takes.
func (b FightBundle) Encode() (contentType string, body []byte, err error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	if err := w.SetBoundary(BundleBoundary); err != nil {
		return "", nil, err
	}
	writeJSON := func(field string, v any) error {
		p, err := w.CreateFormField(field)
		if err != nil {
			return err
		}
		enc := json.NewEncoder(p)
		return enc.Encode(v)
	}
	if err := writeJSON("summary", b.Summary); err != nil {
		return "", nil, err
	}
	p, err := w.CreateFormFile("events", "events.parquet")
	if err != nil {
		return "", nil, err
	}
	if _, err := p.Write(b.Events); err != nil {
		return "", nil, err
	}
	if err := writeJSON("metrics", b.Metrics); err != nil {
		return "", nil, err
	}
	if err := writeJSON("raw_range", b.RawRange); err != nil {
		return "", nil, err
	}
	if err := w.Close(); err != nil {
		return "", nil, err
	}
	return w.FormDataContentType(), buf.Bytes(), nil
}

// Live is the body of the live-snapshot route.
type Live struct {
	Summary   summary.Summary `json:"summary"`
	ElapsedMS int64           `json:"elapsed_ms"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// Complete is the body of the completion route.
type Complete struct {
	FinalOffset   int64          `json:"final_offset"`
	EngineVersion string         `json:"engine_version"`
	Health        session.Health `json:"health"`
}

// Stored says how the server took a write: Created for the first time,
// Duplicate when the same bytes were already stored. Both are success;
// the contract makes at-least-once delivery safe.
type Stored int

// The two outcomes of an idempotent write.
const (
	Created Stored = iota
	Duplicate
)

// CreateReport creates a report and returns its server-assigned id.
func (c *Client) CreateReport(ctx context.Context, in CreateReport) (Report, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return Report{}, err
	}
	var out Report
	if _, err := c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/reports",
		Body: body, Type: "application/json", Out: &out,
	}); err != nil {
		return Report{}, err
	}
	if out.ID == "" {
		return Report{}, fmt.Errorf("POST /v1/reports: the response carried no report id")
	}
	return out, nil
}

// PutFight uploads one closed fight. The body is pre-encoded so the
// queue can replay it unchanged.
func (c *Client) PutFight(ctx context.Context, reportID string, index int, contentType string, body []byte) (Stored, error) {
	status, err := c.do(ctx, request{
		Method: http.MethodPut,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/fights/" + strconv.Itoa(index),
		Body:   body, Type: contentType,
	})
	return storedOf(status), err
}

// PutLive uploads the running summary of the fight in progress. It is
// never queued: a snapshot that fails is replaced five seconds later.
func (c *Client) PutLive(ctx context.Context, reportID string, index int, in Live) error {
	body, err := json.Marshal(in)
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPut,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/fights/" + strconv.Itoa(index) + "/live",
		Body:   body, Type: "application/json",
	})
	return err
}

// PutRaw uploads one compressed chunk of the original log, addressed by
// the uncompressed byte offset of its first byte.
func (c *Client) PutRaw(ctx context.Context, reportID string, offset int64, chunk []byte) (Stored, error) {
	sum := sha256.Sum256(chunk)
	status, err := c.do(ctx, request{
		Method: http.MethodPut,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/raw",
		Query:  url.Values{"offset": []string{strconv.FormatInt(offset, 10)}},
		Body:   chunk, Type: "application/zstd",
		Header: map[string]string{"X-Raw-SHA256": hex.EncodeToString(sum[:])},
	})
	return storedOf(status), err
}

// Complete closes the report and schedules the server's raw-sample
// verification.
func (c *Client) Complete(ctx context.Context, reportID string, in Complete) error {
	body, err := json.Marshal(in)
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPost,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/complete",
		Body:   body, Type: "application/json",
	})
	return err
}

func storedOf(status int) Stored {
	if status == http.StatusOK {
		return Duplicate
	}
	return Created
}
```

- [ ] **Step 7: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/client/ ./internal/fakeapi/ && go test ./internal/client/ ./internal/fakeapi/ -race`
Expected: PASS.

- [ ] **Step 8: Commit**

```bash
git add companion/internal/client/ingest.go \
  companion/internal/client/ingest_test.go \
  companion/internal/fakeapi/ \
  companion/go.mod \
  companion/go.sum
git commit -m "feat(companion): the ingest routes and an in-memory ingest to test them against" \
  -m "The five routes a companion uses to put a raid night on the site,
plus the contract implemented in memory so every test below has one
server to disagree with.

MetricsRow is built by pivoting the fight's roster rather than the
engine's summary.MetricRow: the contract wants one row per player with
metric_dps and metric_hps side by side, and the roster already carries
every field it names. The multipart boundary is pinned, so encoding a
bundle twice produces identical bytes and the queue can store one.

PutFight and PutRaw report Created or Duplicate, which is how the
contract makes at-least-once delivery safe." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 4: Pairing and the device-token store

**Files:**
- Create: `companion/internal/client/devices.go`, `companion/internal/secret/secret.go`
- Test: `companion/internal/client/devices_test.go`, `companion/internal/secret/secret_test.go`

**Interfaces:**
- Consumes: `(*Client).do`, `request` (Task 2); `fakeapi.New`, `fakeapi.Token`, `fakeapi.PairCode` (Task 3); `config.Load`, `config.Save`, `config.Default` (Task 1).
- Produces:
  - `client.Claim{Code, Name, Platform string}`, `client.Device{DeviceID, Token string}`, `client.TokenPrefix`, `(*Client).Claim(ctx, Claim) (Device, error)`.
  - `secret.Service`, `secret.Account`, `secret.Backend` with `secret.Keychain` and `secret.File`.
  - `secret.Store` — `Token() (string, error)`, `SetToken(string) error`, `Clear() error`, `Backend() Backend`.
  - `secret.Keyring{}`, `secret.ConfigFile{Path string}`, `secret.Open(configPath string) Store`, `secret.Migrate(Store, configPath string) error`.

The claim is the one request sent without a token, because it is the request that fetches one; `request.Anonymous` exists for it alone. `secret.Open` probes the keychain with a *read*, not a write, so a machine that then falls back is not left with a stray entry. `Migrate` moves a token that an earlier fallback wrote in `config.json` into the keychain and blanks the file copy, so an install stops leaving the token on disk the moment the keychain starts working.

- [ ] **Step 1: Write the failing tests**

```go
package client_test

import (
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
)

func TestAPairingCodeIsExchangedForADeviceToken(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return "" }})
	if err != nil {
		t.Fatal(err)
	}
	dev, err := c.Claim(t.Context(), client.Claim{
		Code: fakeapi.PairCode, Name: "Justin's iMac", Platform: "darwin/arm64"})
	if err != nil {
		t.Fatal(err)
	}
	if dev.Token != fakeapi.Token || dev.DeviceID == "" {
		t.Fatalf("device = %+v", dev)
	}
}

func TestAWrongPairingCodeIsReportedWithItsField(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return "" }})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Claim(t.Context(), client.Claim{Code: "NOPE", Name: "x", Platform: "linux/amd64"}); err == nil {
		t.Fatal("a wrong code paired")
	}
	if _, err := c.Claim(t.Context(), client.Claim{}); err == nil {
		t.Fatal("an empty code paired")
	}
}
```

```go
package secret

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/zalando/go-keyring"
)

func TestTheKeyringStoreRoundTripsAToken(t *testing.T) {
	keyring.MockInit()
	var s Store = Keyring{}
	if s.Backend() != Keychain {
		t.Errorf("backend = %q", s.Backend())
	}
	got, err := s.Token()
	if err != nil || got != "" {
		t.Fatalf("empty keychain = %q, %v", got, err)
	}
	if err := s.SetToken("fsd_one"); err != nil {
		t.Fatal(err)
	}
	if got, err := s.Token(); err != nil || got != "fsd_one" {
		t.Fatalf("Token = %q, %v", got, err)
	}
	if err := s.Clear(); err != nil {
		t.Fatal(err)
	}
	if err := s.Clear(); err != nil {
		t.Fatalf("clearing twice failed: %v", err)
	}
	if got, _ := s.Token(); got != "" {
		t.Fatalf("Token after Clear = %q", got)
	}
}

func TestTheFileStoreRoundTripsAToken(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := config.Save(path, config.Default()); err != nil {
		t.Fatal(err)
	}
	var s Store = ConfigFile{Path: path}
	if s.Backend() != File {
		t.Errorf("backend = %q", s.Backend())
	}
	if err := s.SetToken("fsd_two"); err != nil {
		t.Fatal(err)
	}
	cfg, err := config.Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.DeviceToken != "fsd_two" {
		t.Fatalf("config device_token = %q", cfg.DeviceToken)
	}
	if got, err := s.Token(); err != nil || got != "fsd_two" {
		t.Fatalf("Token = %q, %v", got, err)
	}
	if err := s.Clear(); err != nil {
		t.Fatal(err)
	}
	if got, _ := s.Token(); got != "" {
		t.Fatalf("Token after Clear = %q", got)
	}
}

func TestMigrateMovesAFileTokenIntoTheKeychain(t *testing.T) {
	keyring.MockInit()
	path := filepath.Join(t.TempDir(), "config.json")
	cfg := config.Default()
	cfg.DeviceToken = "fsd_three"
	if err := config.Save(path, cfg); err != nil {
		t.Fatal(err)
	}
	s := Open(path)
	if s.Backend() != Keychain {
		t.Fatalf("Open chose %q with a working keychain", s.Backend())
	}
	if err := Migrate(s, path); err != nil {
		t.Fatal(err)
	}
	if got, _ := s.Token(); got != "fsd_three" {
		t.Errorf("keychain token = %q", got)
	}
	after, err := config.Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if after.DeviceToken != "" {
		t.Errorf("config still holds %q", after.DeviceToken)
	}
}

func TestMigrateIsANoOpForTheFileStore(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	cfg := config.Default()
	cfg.DeviceToken = "fsd_four"
	if err := config.Save(path, cfg); err != nil {
		t.Fatal(err)
	}
	if err := Migrate(ConfigFile{Path: path}, path); err != nil {
		t.Fatal(err)
	}
	after, _ := config.Load(path)
	if after.DeviceToken != "fsd_four" {
		t.Errorf("the file store lost its token: %q", after.DeviceToken)
	}
}
```

- [ ] **Step 2: Run them to verify they fail**

Run: `cd companion && go test ./internal/client/ ./internal/secret/`
Expected: FAIL, "undefined: Claim" and "undefined: Keyring".

- [ ] **Step 3: Add the keychain dependency**

```bash
cd companion && go get github.com/zalando/go-keyring@v0.2.8
```

- [ ] **Step 4: Write `internal/client/devices.go`**

```go
// companion/internal/client/devices.go
// Pairing. The claim is the one request the companion sends without a
// token, because it is the request that fetches one.
package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Claim is the body of POST /v1/devices/claim.
type Claim struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// Device is the once-shown answer to a claim.
type Device struct {
	DeviceID string `json:"device_id"`
	Token    string `json:"token"`
}

// TokenPrefix is what a device token starts with, per the contract.
const TokenPrefix = "fsd_"

// Claim exchanges a pairing code shown on the site for a device token.
// The token is returned once and never again, so the caller must store
// it before doing anything else.
func (c *Client) Claim(ctx context.Context, in Claim) (Device, error) {
	if in.Code == "" {
		return Device{}, errors.New("the pairing code is empty")
	}
	body, err := json.Marshal(in)
	if err != nil {
		return Device{}, err
	}
	var out Device
	if _, err := c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/devices/claim",
		Body: body, Type: "application/json", Anonymous: true, Out: &out,
	}); err != nil {
		return Device{}, err
	}
	if !strings.HasPrefix(out.Token, TokenPrefix) {
		return Device{}, errors.New("the API returned a token that is not a device token")
	}
	return out, nil
}
```

- [ ] **Step 5: Write `internal/secret/secret.go`**

```go
// companion/internal/secret/secret.go
// Package secret keeps the device token out of plain files wherever the
// operating system offers somewhere better. The keychain is tried
// first; a machine with no keychain, no unlocked login session, or no
// D-Bus falls back to config.json, which the config package already
// writes with owner-only permissions.
package secret

import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/zalando/go-keyring"
)

// Service and Account are the keychain coordinates of the device token.
const (
	Service = "ForeverSixty Companion"
	Account = "device-token"
)

// Backend names where a store keeps the token, for the settings page
// and the troubleshooting section of the README.
type Backend string

// The two places a token can live.
const (
	Keychain Backend = "keychain"
	File     Backend = "file"
)

// Store reads and writes the device token.
type Store interface {
	// Token is the stored token, or the empty string when the device
	// is not paired.
	Token() (string, error)
	SetToken(token string) error
	Clear() error
	Backend() Backend
}

// Keyring is the OS keychain: Keychain on macOS, Credential Manager on
// Windows, the Secret Service on Linux.
type Keyring struct{}

// Token reads the token from the keychain.
func (Keyring) Token() (string, error) {
	v, err := keyring.Get(Service, Account)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("read the device token from the keychain: %w", err)
	}
	return v, nil
}

// SetToken writes the token to the keychain.
func (Keyring) SetToken(token string) error {
	if err := keyring.Set(Service, Account, token); err != nil {
		return fmt.Errorf("write the device token to the keychain: %w", err)
	}
	return nil
}

// Clear removes the token. A token that is not there is not an error:
// unpairing twice must succeed.
func (Keyring) Clear() error {
	err := keyring.Delete(Service, Account)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("delete the device token from the keychain: %w", err)
	}
	return nil
}

// Backend names this store.
func (Keyring) Backend() Backend { return Keychain }

// ConfigFile keeps the token in config.json's device_token field.
type ConfigFile struct{ Path string }

// Token reads the token from config.json.
func (c ConfigFile) Token() (string, error) {
	cfg, err := config.Load(c.Path)
	if err != nil {
		return "", err
	}
	return cfg.DeviceToken, nil
}

// SetToken writes the token into config.json, leaving every other
// setting as it is.
func (c ConfigFile) SetToken(token string) error {
	cfg, err := config.Load(c.Path)
	if err != nil {
		return err
	}
	cfg.DeviceToken = token
	return config.Save(c.Path, cfg)
}

// Clear empties the token field.
func (c ConfigFile) Clear() error { return c.SetToken("") }

// Backend names this store.
func (ConfigFile) Backend() Backend { return File }

// Open returns the keychain store when the keychain answers a probe,
// and the config-file store when it does not. The probe is a read: a
// write would leave an entry behind on a machine that then falls back.
func Open(configPath string) Store {
	if _, err := keyring.Get(Service, Account); err == nil || errors.Is(err, keyring.ErrNotFound) {
		return Keyring{}
	}
	return ConfigFile{Path: configPath}
}

// Migrate moves a token found in config.json into the keychain and
// blanks the file copy. It runs at startup so an install that once fell
// back stops leaving the token on disk once the keychain works.
func Migrate(s Store, configPath string) error {
	if s.Backend() != Keychain {
		return nil
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}
	if cfg.DeviceToken == "" {
		return nil
	}
	if err := s.SetToken(cfg.DeviceToken); err != nil {
		return err
	}
	cfg.DeviceToken = ""
	return config.Save(configPath, cfg)
}
```

- [ ] **Step 6: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/client/ ./internal/secret/ && go test ./internal/client/ ./internal/secret/ -race`
Expected: PASS. `keyring.MockInit()` replaces the real keychain, so the suite never prompts for a password and never touches the developer's login keychain.

- [ ] **Step 7: Commit**

```bash
git add companion/internal/client/devices.go \
  companion/internal/client/devices_test.go \
  companion/internal/secret/ \
  companion/go.mod \
  companion/go.sum
git commit -m "feat(companion): pairing, and the device token in the OS keychain" \
  -m "A pairing code shown on the site becomes an upload token the
companion stores in macOS Keychain, Windows Credential Manager or the
Linux Secret Service. Where there is no keychain -- a headless box, a
locked login session, no D-Bus -- it falls back to config.json, which
is already written readable only by its owner, and the settings page
says which one is in use.

Open probes with a read rather than a write so a machine that falls
back is not left with a stray keychain entry, and Migrate moves a
token an earlier fallback wrote into the keychain the moment one
appears." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Task 5: The durable upload queue

**Files:**
- Create: `companion/internal/queue/queue.go`
- Test: `companion/internal/queue/queue_test.go`

**Interfaces:**
- Consumes: nothing.
- Produces:
  - `queue.Kind` with `queue.Fight`, `queue.Raw`, `queue.Complete`.
  - `queue.Item{Seq int64; Kind Kind; ReportKey string; FightIndex int; Offset int64; ContentType string; Attempts int; LastError string; EnqueuedAt time.Time}`.
  - `queue.Open(dir string) (*Queue, error)`, `(*Queue).SetClock(func() time.Time)`, `.Enqueue(Item, body []byte) (int64, error)`, `.Len() (int, error)`, `.Head() (*Lease, error)`, `queue.ErrEmpty`.
  - `queue.Lease{Item Item; Body []byte}` with `.Ack() error`, `.Fail(error) error`, `.Drop() error`.

This is the package the ordering guarantee lives in. `Head` hands out one lease at a time and always the lowest sequence number; `Fail` puts the item back where it was with one more attempt recorded. Nothing can overtake a failing upload, which is exactly what "fights arrive in the order they were fought" means.

Live snapshots are deliberately not a `Kind`. A snapshot that fails is replaced five seconds later, and queueing them would put a fight behind a stale picture of a fight.

The body is written before the metadata, so a crash between the two leaves an orphan `.bin` that `Head` skips rather than an item whose body is missing.

- [ ] **Step 1: Write the failing test**

```go
package queue

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestItemsComeBackInSequenceOrder(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for i := range 3 {
		if _, err := q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: i},
			[]byte{byte('a' + i)}); err != nil {
			t.Fatal(err)
		}
	}
	if n, _ := q.Len(); n != 3 {
		t.Fatalf("Len = %d, want 3", n)
	}
	for i := range 3 {
		l, err := q.Head()
		if err != nil {
			t.Fatal(err)
		}
		if l.Item.FightIndex != i || string(l.Body) != string([]byte{byte('a' + i)}) {
			t.Fatalf("head %d = %+v body %q", i, l.Item, l.Body)
		}
		if err := l.Ack(); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := q.Head(); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Head on an empty queue = %v", err)
	}
}

func TestAFailedItemStaysAtTheHeadAndCountsItsAttempts(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 0}, []byte("first"))
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 1}, []byte("second"))

	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if err := l.Fail(errors.New("network is unreachable")); err != nil {
		t.Fatal(err)
	}
	again, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if again.Item.FightIndex != 0 {
		t.Fatalf("head after a failure = fight %d, want 0", again.Item.FightIndex)
	}
	if again.Item.Attempts != 1 || again.Item.LastError != "network is unreachable" {
		t.Fatalf("item = %+v", again.Item)
	}
}

func TestAQueueSurvivesAReopen(t *testing.T) {
	dir := t.TempDir()
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 7}, []byte("bundle"))

	reopened, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	l, err := reopened.Head()
	if err != nil {
		t.Fatal(err)
	}
	if l.Item.FightIndex != 7 || string(l.Body) != "bundle" {
		t.Fatalf("item = %+v body %q", l.Item, l.Body)
	}
	seq, err := reopened.Enqueue(Item{Kind: Complete, ReportKey: "local-1"}, []byte("{}"))
	if err != nil {
		t.Fatal(err)
	}
	if seq <= l.Item.Seq {
		t.Fatalf("new sequence %d is not after %d", seq, l.Item.Seq)
	}
}

func TestAnItemWithNoBodyIsSkippedRatherThanBlocking(t *testing.T) {
	dir := t.TempDir()
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	seq, err := q.Enqueue(Item{Kind: Raw, ReportKey: "local-1", Offset: 0}, []byte("chunk"))
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Raw, ReportKey: "local-1", Offset: 4194304}, []byte("chunk2"))
	if err := os.Remove(filepath.Join(dir, name(seq)+".bin")); err != nil {
		t.Fatal(err)
	}
	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if l.Item.Offset != 4194304 {
		t.Fatalf("head = %+v, want the second chunk", l.Item)
	}
}

func TestOnlyOneLeaseIsOutAtATime(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1"}, []byte("a"))
	q.Enqueue(Item{Kind: Fight, ReportKey: "local-1", FightIndex: 1}, []byte("b"))
	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := q.Head(); !errors.Is(err, ErrEmpty) {
		t.Fatal("a second lease was handed out while one was in flight")
	}
	if err := l.Drop(); err != nil {
		t.Fatal(err)
	}
	next, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if next.Item.FightIndex != 1 {
		t.Fatalf("after a drop the head is %+v", next.Item)
	}
}

func TestOpeningAQueueInAnImpossiblePlaceIsAnError(t *testing.T) {
	file := filepath.Join(t.TempDir(), "not-a-directory")
	if err := os.WriteFile(file, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(filepath.Join(file, "queue")); err == nil {
		t.Fatal("Open accepted a path inside a file")
	}
}

func TestUnrelatedFilesInTheQueueDirectoryAreLeftAlone(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "README.txt"), []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "notanumber.json"), []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	q, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	if n, _ := q.Len(); n != 0 {
		t.Fatalf("Len = %d, want 0", n)
	}
	if _, err := q.Head(); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Head = %v", err)
	}
}

func TestTheClockCanBeReplaced(t *testing.T) {
	q, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	when := time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)
	q.SetClock(func() time.Time { return when })
	if _, err := q.Enqueue(Item{Kind: Complete, ReportKey: "k"}, []byte("{}")); err != nil {
		t.Fatal(err)
	}
	l, err := q.Head()
	if err != nil {
		t.Fatal(err)
	}
	if !l.Item.EnqueuedAt.Equal(when) {
		t.Fatalf("enqueued at %s", l.Item.EnqueuedAt)
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/queue/`
Expected: FAIL, "undefined: Open".

- [ ] **Step 3: Write the implementation**

```go
// companion/internal/queue/queue.go
// Package queue is the companion's durable outbox. Everything the
// uploader sends goes through it in sequence order and stays on disk
// until the server has taken it, which is what makes "fights arrive in
// order across a network drop and a restart" a property of the design
// rather than a hope.
package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Kind is what one queued item asks the uploader to do.
type Kind string

// The three durable operations. Live snapshots are deliberately absent:
// a snapshot that fails is replaced by the next one five seconds later,
// and queueing them would delay the fights behind them.
const (
	Fight    Kind = "fight"
	Raw      Kind = "raw"
	Complete Kind = "complete"
)

// Item is one queued operation. The body lives beside it in a .bin
// file, so a 6 MiB bundle never passes through JSON.
type Item struct {
	Seq         int64     `json:"seq"`
	Kind        Kind      `json:"kind"`
	ReportKey   string    `json:"report_key"`
	FightIndex  int       `json:"fight_index,omitempty"`
	Offset      int64     `json:"offset,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Attempts    int       `json:"attempts"`
	LastError   string    `json:"last_error,omitempty"`
	EnqueuedAt  time.Time `json:"enqueued_at"`
}

// Queue is a directory of items, drained strictly in sequence order.
type Queue struct {
	dir string

	mu   sync.Mutex
	next int64
	// leased is the sequence number currently handed out, so a second
	// Head while one is in flight returns nothing rather than a
	// duplicate.
	leased int64
	now    func() time.Time
}

// Open opens or creates the queue directory and recovers the next
// sequence number from whatever is still pending.
func Open(dir string) (*Queue, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("create the queue directory: %w", err)
	}
	q := &Queue{dir: dir, next: 1, now: time.Now}
	seqs, err := q.pending()
	if err != nil {
		return nil, err
	}
	if len(seqs) > 0 {
		q.next = seqs[len(seqs)-1] + 1
	}
	return q, nil
}

// SetClock replaces the clock, for tests.
func (q *Queue) SetClock(now func() time.Time) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.now = now
}

func (q *Queue) meta(seq int64) string { return filepath.Join(q.dir, name(seq)+".json") }
func (q *Queue) body(seq int64) string { return filepath.Join(q.dir, name(seq)+".bin") }

func name(seq int64) string { return fmt.Sprintf("%020d", seq) }

// pending lists the sequence numbers on disk, ascending.
func (q *Queue) pending() ([]int64, error) {
	entries, err := os.ReadDir(q.dir)
	if err != nil {
		return nil, fmt.Errorf("read the queue directory: %w", err)
	}
	var seqs []int64
	for _, e := range entries {
		n := e.Name()
		if !strings.HasSuffix(n, ".json") {
			continue
		}
		seq, err := strconv.ParseInt(strings.TrimSuffix(n, ".json"), 10, 64)
		if err != nil {
			continue // not ours; leave it alone
		}
		seqs = append(seqs, seq)
	}
	sort.Slice(seqs, func(i, j int) bool { return seqs[i] < seqs[j] })
	return seqs, nil
}

// Enqueue appends one item and returns its sequence number. The body
// is written first and the metadata second, so a crash between the two
// leaves an orphan .bin that Head ignores rather than an item with no
// body.
func (q *Queue) Enqueue(it Item, body []byte) (int64, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	it.Seq = q.next
	it.EnqueuedAt = q.now().UTC()
	if err := writeFile(q.body(it.Seq), body); err != nil {
		return 0, err
	}
	if err := q.writeMeta(it); err != nil {
		os.Remove(q.body(it.Seq))
		return 0, err
	}
	q.next++
	return it.Seq, nil
}

func (q *Queue) writeMeta(it Item) error {
	b, err := json.Marshal(it)
	if err != nil {
		return err
	}
	return writeFile(q.meta(it.Seq), append(b, '\n'))
}

// writeFile writes atomically: a torn queue file would cost a fight.
func writeFile(path string, b []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), ".q-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(name, 0o600); err != nil {
		return err
	}
	return os.Rename(name, path)
}

// Len is how many items are waiting.
func (q *Queue) Len() (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	seqs, err := q.pending()
	if err != nil {
		return 0, err
	}
	return len(seqs), nil
}

// ErrEmpty is returned by Head when nothing is queued.
var ErrEmpty = errors.New("queue: empty")

// Lease is the item at the head of the queue, checked out.
type Lease struct {
	Item Item
	Body []byte

	q *Queue
}

// Head checks out the oldest item. It returns ErrEmpty when there is
// nothing to do. Only one lease exists at a time: the whole point is
// that item N+1 does not go out before item N.
func (q *Queue) Head() (*Lease, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.leased != 0 {
		return nil, ErrEmpty
	}
	seqs, err := q.pending()
	if err != nil {
		return nil, err
	}
	for _, seq := range seqs {
		b, err := os.ReadFile(q.meta(seq))
		if err != nil {
			return nil, fmt.Errorf("read queue item %d: %w", seq, err)
		}
		var it Item
		if err := json.Unmarshal(b, &it); err != nil {
			return nil, fmt.Errorf("parse queue item %d: %w", seq, err)
		}
		body, err := os.ReadFile(q.body(seq))
		if errors.Is(err, fs.ErrNotExist) {
			// The body never landed. Drop the metadata rather than
			// block the queue on an item that can never be sent.
			os.Remove(q.meta(seq))
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("read queue body %d: %w", seq, err)
		}
		q.leased = seq
		return &Lease{Item: it, Body: body, q: q}, nil
	}
	return nil, ErrEmpty
}

// Ack removes the item: the server has it.
func (l *Lease) Ack() error {
	q := l.q
	q.mu.Lock()
	defer q.mu.Unlock()
	q.leased = 0
	if err := os.Remove(q.meta(l.Item.Seq)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if err := os.Remove(q.body(l.Item.Seq)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	return nil
}

// Fail returns the item to the head of the queue with one more attempt
// recorded. Nothing behind it moves.
func (l *Lease) Fail(cause error) error {
	q := l.q
	q.mu.Lock()
	defer q.mu.Unlock()
	q.leased = 0
	l.Item.Attempts++
	if cause != nil {
		l.Item.LastError = cause.Error()
	}
	return q.writeMeta(l.Item)
}

// Drop removes an item the server will never accept, so the queue
// behind it can move. The reason is the caller's to log.
func (l *Lease) Drop() error { return l.Ack() }
```

- [ ] **Step 4: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/queue/ && go test ./internal/queue/ -race`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add companion/internal/queue/
git commit -m "feat(companion): a durable outbox that drains strictly in order" \
  -m "One directory, one item per file, sequence numbers, one lease at a
time. A failed upload goes back to the head with its attempt count,
and nothing behind it moves -- which is the whole mechanism behind
'fights arrive in the order the raid fought them' surviving a network
drop and a restart.

Bodies are written before metadata so a crash leaves an orphan that
Head skips, never an item that can never be sent. Live snapshots are
not queued on purpose: a stale one must never delay a real fight." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 6: Finding World of Warcraft

**Files:**
- Create: `companion/internal/wow/wow.go`
- Test: `companion/internal/wow/wow_test.go`

**Interfaces:**
- Consumes: nothing.
- Produces:
  - `wow.LogGlob` (`WoWCombatLog*.txt`), `wow.AdvancedLoggingHelp` (the exact menu path spec section 5 asks for).
  - `wow.Install{Path, Flavor, Logs, WTF string}`.
  - `wow.Candidates() []string`, `wow.Scan(root string) []Install`, `wow.Detect() []Install`, `wow.Pick(path string) ([]Install, error)`, `wow.ErrNotAnInstall`.
  - `(Install).EnsureLogs() error`, `.CombatLogs() ([]string, error)`, `.AdvancedLoggingPath() string`, `.AdvancedLogging() (bool, error)`, `.Accounts() ([]string, error)`, `.SavedVariables(addon string) ([]string, error)`.

An install is a *flavour directory* — the folder holding `WTF` and `Logs` — not the folder above it, because that is the unit the game rotates logs and saves variables in. `WTF` is the marker: every install has one and nothing else does. Nothing here hardcodes a list of flavour names, so a private server's folder called anything at all is still found; `Pick` accepts either the flavour directory or its parent, which is what a file dialog most often returns.

- [ ] **Step 1: Write the failing test**

```go
package wow

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// fakeInstall lays out a flavour directory the way the game does.
func fakeInstall(t *testing.T, root, flavor, account string) Install {
	t.Helper()
	dir := filepath.Join(root, flavor)
	for _, p := range []string{
		filepath.Join(dir, "Logs"),
		filepath.Join(dir, "WTF", "Account", account, "SavedVariables"),
	} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return installAt(dir)
}

func TestScanFindsFlavourDirectoriesUnderARoot(t *testing.T) {
	root := t.TempDir()
	fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	fakeInstall(t, root, "_forever_", "ACCOUNT#1")
	if err := os.MkdirAll(filepath.Join(root, "Utils"), 0o755); err != nil {
		t.Fatal(err)
	}
	got := Scan(root)
	if len(got) != 2 {
		t.Fatalf("Scan found %d installs: %+v", len(got), got)
	}
	if got[0].Flavor != "_classic_era_" || got[1].Flavor != "_forever_" {
		t.Fatalf("flavours = %q, %q", got[0].Flavor, got[1].Flavor)
	}
	if got[0].Logs != filepath.Join(root, "_classic_era_", "Logs") {
		t.Errorf("logs = %q", got[0].Logs)
	}
}

func TestPickAcceptsAFlavourDirectoryOrItsParent(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")

	fromParent, err := Pick(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(fromParent) != 1 || fromParent[0].Path != in.Path {
		t.Fatalf("Pick(parent) = %+v", fromParent)
	}
	fromFlavor, err := Pick(in.Path)
	if err != nil {
		t.Fatal(err)
	}
	if len(fromFlavor) != 1 || fromFlavor[0].Path != in.Path {
		t.Fatalf("Pick(flavour) = %+v", fromFlavor)
	}
}

func TestPickRefusesAFolderThatIsNotTheGame(t *testing.T) {
	if _, err := Pick(t.TempDir()); !errors.Is(err, ErrNotAnInstall) {
		t.Fatalf("err = %v, want ErrNotAnInstall", err)
	}
	if _, err := Pick(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("Pick accepted a path that does not exist")
	}
	file := filepath.Join(t.TempDir(), "WoW.exe")
	if err := os.WriteFile(file, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Pick(file); !errors.Is(err, ErrNotAnInstall) {
		t.Fatalf("Pick(file) = %v", err)
	}
}

func TestAdvancedLoggingIsReadFromConfigWTF(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	write := func(body string) {
		t.Helper()
		if err := os.WriteFile(in.AdvancedLoggingPath(), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := in.AdvancedLogging(); err == nil {
		t.Error("a missing Config.wtf was not reported")
	}
	write("SET gxWindow \"1\"\n")
	if on, err := in.AdvancedLogging(); err != nil || on {
		t.Errorf("with no setting: %v, %v", on, err)
	}
	write("SET gxWindow \"1\"\nSET advancedCombatLogging \"0\"\n")
	if on, err := in.AdvancedLogging(); err != nil || on {
		t.Errorf("with the box off: %v, %v", on, err)
	}
	write("SET advancedCombatLogging \"1\"\nSET gxWindow \"1\"\n")
	if on, err := in.AdvancedLogging(); err != nil || !on {
		t.Errorf("with the box on: %v, %v", on, err)
	}
}

func TestCombatLogsAndSavedVariablesArePerInstall(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	fakeInstall(t, root, "_classic_era_", "ACCOUNT#2") // second account, same install
	for _, n := range []string{"WoWCombatLog.txt", "WoWCombatLog-120926_200000.txt", "notes.md"} {
		if err := os.WriteFile(filepath.Join(in.Logs, n), nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	logs, err := in.CombatLogs()
	if err != nil {
		t.Fatal(err)
	}
	if len(logs) != 2 {
		t.Fatalf("CombatLogs = %v", logs)
	}
	sv, err := in.SavedVariables("ForeverSixty")
	if err != nil {
		t.Fatal(err)
	}
	if len(sv) != 2 {
		t.Fatalf("SavedVariables = %v", sv)
	}
	want := filepath.Join(in.WTF, "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	if sv[0] != want {
		t.Errorf("SavedVariables[0] = %q, want %q", sv[0], want)
	}
}

func TestEnsureLogsCreatesTheDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "_classic_era_")
	if err := os.MkdirAll(filepath.Join(dir, "WTF"), 0o755); err != nil {
		t.Fatal(err)
	}
	in := installAt(dir)
	if err := in.EnsureLogs(); err != nil {
		t.Fatal(err)
	}
	if fi, err := os.Stat(in.Logs); err != nil || !fi.IsDir() {
		t.Fatalf("Logs = %v, %v", fi, err)
	}
}

func TestCandidatesAreAbsolutePathsForThisOS(t *testing.T) {
	got := Candidates()
	if len(got) == 0 {
		t.Fatal("no candidate paths")
	}
	for _, p := range got {
		if !filepath.IsAbs(p) {
			t.Errorf("%q is not absolute", p)
		}
	}
	// Detect must not panic on a machine with no game installed.
	Detect()
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/wow/`
Expected: FAIL, "undefined: Scan".

- [ ] **Step 3: Write the implementation**

```go
// companion/internal/wow/wow.go
// Package wow finds the game on disk. An install is a flavour
// directory — the folder holding WTF and Logs — because that is the
// unit the game rotates logs and saves variables in, and because a
// private server's folder may be named anything at all.
package wow

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// LogGlob matches the combat logs the game writes. The plain
// WoWCombatLog.txt is the live one; the game may also leave
// timestamped siblings behind.
const LogGlob = "WoWCombatLog*.txt"

// Install is one flavour directory of the game.
type Install struct {
	// Path is the flavour directory itself.
	Path string `json:"path"`
	// Flavor is its folder name, "_classic_era_" and the like, or
	// the folder name of a private-server install.
	Flavor string `json:"flavor"`
	// Logs is <Path>/Logs and WTF is <Path>/WTF.
	Logs string `json:"logs"`
	WTF  string `json:"wtf"`
}

// Candidates are the paths the companion looks in before asking the
// player to pick a folder.
func Candidates() []string {
	home, _ := os.UserHomeDir()
	switch runtime.GOOS {
	case "darwin":
		return []string{
			"/Applications/World of Warcraft",
			filepath.Join(home, "Applications", "World of Warcraft"),
		}
	case "windows":
		var out []string
		for _, env := range []string{"ProgramFiles(x86)", "ProgramFiles"} {
			if v := os.Getenv(env); v != "" {
				out = append(out, filepath.Join(v, "World of Warcraft"))
			}
		}
		return append(out, `C:\Program Files (x86)\World of Warcraft`)
	default:
		return []string{
			filepath.Join(home, "Games", "world-of-warcraft", "drive_c",
				"Program Files (x86)", "World of Warcraft"),
			filepath.Join(home, ".wine", "drive_c",
				"Program Files (x86)", "World of Warcraft"),
		}
	}
}

// isInstall reports whether a directory is a flavour directory. WTF is
// the marker: every install has one and nothing else does.
func isInstall(dir string) bool {
	fi, err := os.Stat(filepath.Join(dir, "WTF"))
	return err == nil && fi.IsDir()
}

func installAt(dir string) Install {
	return Install{
		Path:   dir,
		Flavor: filepath.Base(dir),
		Logs:   filepath.Join(dir, "Logs"),
		WTF:    filepath.Join(dir, "WTF"),
	}
}

// Scan returns every install at or one level below root, sorted by
// path so two runs agree.
func Scan(root string) []Install {
	var out []Install
	if isInstall(root) {
		out = append(out, installAt(root))
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return out
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		dir := filepath.Join(root, e.Name())
		if isInstall(dir) {
			out = append(out, installAt(dir))
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out
}

// Detect scans every candidate path.
func Detect() []Install {
	var out []Install
	seen := map[string]bool{}
	for _, root := range Candidates() {
		for _, in := range Scan(root) {
			if !seen[in.Path] {
				seen[in.Path] = true
				out = append(out, in)
			}
		}
	}
	return out
}

// ErrNotAnInstall says a picked folder is not the game.
var ErrNotAnInstall = errors.New("that folder is not a World of Warcraft install: " +
	"pick the folder that contains WTF and Logs, such as _classic_era_")

// Pick validates a folder the player chose. It accepts either a
// flavour directory or the folder above one, which is what a file
// dialog most often returns.
func Pick(path string) ([]Install, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	if !fi.IsDir() {
		return nil, ErrNotAnInstall
	}
	found := Scan(path)
	if len(found) == 0 {
		return nil, ErrNotAnInstall
	}
	return found, nil
}

// EnsureLogs creates the Logs directory if the player has never
// logged, so the watcher has something to watch from the first run.
func (i Install) EnsureLogs() error {
	return os.MkdirAll(i.Logs, 0o755)
}

// CombatLogs lists the combat logs in the install, newest last.
func (i Install) CombatLogs() ([]string, error) {
	matches, err := filepath.Glob(filepath.Join(i.Logs, LogGlob))
	if err != nil {
		return nil, err
	}
	sort.Strings(matches)
	return matches, nil
}

// AdvancedLoggingPath is the settings file the checkbox writes to.
func (i Install) AdvancedLoggingPath() string { return filepath.Join(i.WTF, "Config.wtf") }

// AdvancedLogging reports whether advanced combat logging is on. A
// missing setting means off: the game omits the line until the box is
// ticked once.
func (i Install) AdvancedLogging() (bool, error) {
	f, err := os.Open(i.AdvancedLoggingPath())
	if err != nil {
		return false, fmt.Errorf("read %s: %w", i.AdvancedLoggingPath(), err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 3 && fields[0] == "SET" && fields[1] == "advancedCombatLogging" {
			return strings.Trim(fields[2], `"`) == "1", nil
		}
	}
	if err := sc.Err(); err != nil {
		return false, err
	}
	return false, nil
}

// AdvancedLoggingHelp is the exact menu path the status page shows when
// advanced logging is off, per spec section 5.
const AdvancedLoggingHelp = "Options → Network → Advanced Combat Logging"

// Accounts lists the account folders under WTF/Account, sorted. The
// game names them after the account, upper-cased.
func (i Install) Accounts() ([]string, error) {
	dir := filepath.Join(i.WTF, "Account")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", dir, err)
	}
	var out []string
	for _, e := range entries {
		if e.IsDir() && e.Name() != "SavedVariables" {
			out = append(out, filepath.Join(dir, e.Name()))
		}
	}
	sort.Strings(out)
	return out, nil
}

// SavedVariables returns the path of one addon's saved-variables file
// in each account folder, whether or not it exists yet: the addon-sync
// loop watches paths that may appear later.
func (i Install) SavedVariables(addon string) ([]string, error) {
	accounts, err := i.Accounts()
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(accounts))
	for _, a := range accounts {
		out = append(out, filepath.Join(a, "SavedVariables", addon+".lua"))
	}
	return out, nil
}
```

- [ ] **Step 4: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/wow/ && go test ./internal/wow/ -race`
Expected: PASS. The last test calls `Detect()` for its side effects only: it must not panic on a machine with no game installed, which is every CI runner.

- [ ] **Step 5: Commit**

```bash
git add companion/internal/wow/
git commit -m "feat(companion): find the game, and the advanced-logging check" \
  -m "An install is the flavour directory -- the folder holding WTF and
Logs -- because that is the unit the game rotates logs and saves
variables in, and because a private server's folder may be named
anything. WTF is the marker; no list of flavour names is hardcoded.

Pick takes either the flavour folder or its parent, which is what a
file dialog usually hands back, and AdvancedLogging reads Config.wtf
so the status page can show the exact menu path when the box is
off." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Task 7: Tailing the combat log — rotation and the two idle rules

**Files:**
- Create: `companion/internal/watch/watch.go`
- Test: `companion/internal/watch/watch_test.go`

**Interfaces:**
- Consumes: `wow.LogGlob` (Task 6).
- Produces:
  - `watch.NewReportIdle` (30 minutes), `watch.CompleteIdle` (15 minutes), `watch.ReadChunk` (1 MiB).
  - `watch.Kind` with `watch.Start`, `watch.Append`, `watch.Complete`, and `Kind.String()`.
  - `watch.Event{Kind Kind; Path string; Offset int64; Data []byte; At time.Time}` — `Offset` is a *file* offset here, the only place in the module that is true.
  - `watch.Options{Dir string; NewReportIdle, CompleteIdle time.Duration; ReadChunk int}`, `watch.New(Options) *Watcher`.
  - `(*Watcher).Resume(path string, offset int64, lastAppend time.Time) error`, `.Offset() int64`, `.Path() string`, `.Poll(now time.Time) ([]Event, error)`.
  - `watch.ReadRange(path string, start, end int64) ([]byte, error)`.

The watcher owns no goroutine and reads no clock: `Poll` is handed the time. That is what makes a thirty-minute rule testable in a millisecond, and it is why every timing rule in the contract has a test rather than a comment.

Two behaviours are worth reading twice. A file that was already there on the *first* poll holds a night nobody watched, so the watcher begins at its end; a file that appears while the watcher is running begins at zero. And a report that was open when the companion stopped is continued through `Resume`, not restarted — Task 9 calls it from the state files.

- [ ] **Step 1: Write the failing test**

```go
package watch

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// appendTo appends to a log and stamps its modification time, because
// the watcher picks the newest file and a test must control that.
func appendTo(t *testing.T, path, text string, mtime time.Time) {
	t.Helper()
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, mtime, mtime); err != nil {
		t.Fatal(err)
	}
}

func kinds(evs []Event) []string {
	out := make([]string, 0, len(evs))
	for _, e := range evs {
		out = append(out, e.Kind.String())
	}
	return out
}

func same(got []string, want ...string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestAppendedBytesArriveWithTheirOffset(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})

	if evs, err := w.Poll(t0); err != nil || len(evs) != 0 {
		t.Fatalf("empty directory = %v, %v", kinds(evs), err)
	}
	appendTo(t, log, "line one\n", t0)
	evs, err := w.Poll(t0.Add(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if evs[1].Offset != 0 || string(evs[1].Data) != "line one\n" {
		t.Fatalf("append = %+v", evs[1])
	}
	appendTo(t, log, "line two\n", t0.Add(2*time.Second))
	evs, err = w.Poll(t0.Add(3 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "append") || evs[0].Offset != 9 || string(evs[0].Data) != "line two\n" {
		t.Fatalf("second append = %v %+v", kinds(evs), evs)
	}
	if w.Offset() != 18 {
		t.Errorf("offset = %d, want 18", w.Offset())
	}
}

func TestAPreExistingLogIsNotReplayedFromItsStart(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "last night\n", t0.Add(-24*time.Hour))

	w := New(Options{Dir: dir})
	if evs, err := w.Poll(t0); err != nil || len(evs) != 0 {
		t.Fatalf("first poll = %v, %v", kinds(evs), err)
	}
	if w.Offset() != 11 {
		t.Fatalf("offset = %d, want the file's size 11", w.Offset())
	}
	appendTo(t, log, "tonight\n", t0.Add(time.Second))
	evs, err := w.Poll(t0.Add(2 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "start", "append") || string(evs[1].Data) != "tonight\n" {
		t.Fatalf("events = %v %+v", kinds(evs), evs)
	}
}

func TestRotationToATimestampedFileClosesTheReportAndOpensANewOne(t *testing.T) {
	dir := t.TempDir()
	live := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})
	w.Poll(t0)
	appendTo(t, live, "first night\n", t0)
	if evs, _ := w.Poll(t0.Add(time.Second)); !same(kinds(evs), "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	// The game rotates: the live file is renamed and a new one starts.
	rotated := filepath.Join(dir, "WoWCombatLog-120926_200000.txt")
	if err := os.Rename(live, rotated); err != nil {
		t.Fatal(err)
	}
	appendTo(t, live, "second night\n", t0.Add(time.Minute))
	evs, err := w.Poll(t0.Add(time.Minute + time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete", "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if evs[0].Offset != 12 {
		t.Errorf("the completed report's final offset = %d, want 12", evs[0].Offset)
	}
	if evs[2].Offset != 0 || string(evs[2].Data) != "second night\n" {
		t.Errorf("new report append = %+v", evs[2])
	}
}

func TestATruncatedFileIsTreatedAsANewReport(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})
	w.Poll(t0)
	appendTo(t, log, "aaaaaaaaaa\n", t0)
	w.Poll(t0.Add(time.Second))

	if err := os.WriteFile(log, []byte("b\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(log, t0.Add(2*time.Second), t0.Add(2*time.Second)); err != nil {
		t.Fatal(err)
	}
	evs, err := w.Poll(t0.Add(3 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete", "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if string(evs[2].Data) != "b\n" {
		t.Errorf("append = %q", evs[2].Data)
	}
}

func TestFifteenQuietMinutesCompleteTheReport(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})
	w.Poll(t0)
	appendTo(t, log, "pull\n", t0)
	w.Poll(t0.Add(time.Second))

	if evs, _ := w.Poll(t0.Add(14 * time.Minute)); len(evs) != 0 {
		t.Fatalf("completed after fourteen minutes: %v", kinds(evs))
	}
	evs, err := w.Poll(t0.Add(16 * time.Minute))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete") || evs[0].Offset != 5 {
		t.Fatalf("events = %v %+v", kinds(evs), evs)
	}
	if evs2, _ := w.Poll(t0.Add(17 * time.Minute)); len(evs2) != 0 {
		t.Fatalf("completed twice: %v", kinds(evs2))
	}
}

func TestAGapLongerThanThirtyMinutesStartsANewReport(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	// CompleteIdle is set past the gap so the only rule under test is
	// the thirty-minute one.
	w := New(Options{Dir: dir, CompleteIdle: 4 * time.Hour})
	w.Poll(t0)
	appendTo(t, log, "first raid\n", t0)
	w.Poll(t0.Add(time.Second))

	late := t0.Add(31 * time.Minute)
	appendTo(t, log, "second raid\n", late)
	evs, err := w.Poll(late.Add(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete", "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if evs[1].Offset != 11 {
		t.Errorf("the new report starts at %d, want 11", evs[1].Offset)
	}
}

func TestResumeContinuesAReportAcrossARestart(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "before the restart\n", t0)

	w := New(Options{Dir: dir})
	if err := w.Resume(log, 19, t0); err != nil {
		t.Fatal(err)
	}
	appendTo(t, log, "after\n", t0.Add(time.Second))
	evs, err := w.Poll(t0.Add(2 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "append") || string(evs[0].Data) != "after\n" {
		t.Fatalf("events = %v %+v", kinds(evs), evs)
	}
}

func TestResumeRefusesAnOffsetPastTheEndOfTheFile(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "short\n", t0)
	w := New(Options{Dir: dir})
	if err := w.Resume(log, 9999, t0); err == nil {
		t.Fatal("Resume accepted an offset past the end of the file")
	}
	if err := w.Resume(filepath.Join(dir, "missing.txt"), 0, t0); err == nil {
		t.Fatal("Resume accepted a file that is not there")
	}
}

func TestALargeAppendIsDeliveredInChunks(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir, ReadChunk: 4})
	w.Poll(t0)
	appendTo(t, log, "abcdefghij", t0)

	var got string
	for i := range 3 {
		evs, err := w.Poll(t0.Add(time.Duration(i+1) * time.Second))
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range evs {
			if e.Kind == Append {
				got += string(e.Data)
			}
		}
	}
	if got != "abcdefghij" {
		t.Fatalf("reassembled %q", got)
	}
}

func TestReadRangeReturnsAFightsBytes(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "0123456789", t0)
	got, err := ReadRange(log, 2, 6)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "2345" {
		t.Fatalf("ReadRange = %q", got)
	}
	if _, err := ReadRange(log, 6, 2); err == nil {
		t.Error("ReadRange accepted a backwards range")
	}
	if _, err := ReadRange(log, 8, 99); err == nil {
		t.Error("ReadRange accepted a range past the end of the file")
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/watch/`
Expected: FAIL, "undefined: New", "undefined: Options".

- [ ] **Step 3: Write the implementation**

```go
// companion/internal/watch/watch.go
// Package watch turns a Logs directory into a stream of report
// boundaries and appended bytes. It owns no goroutine and reads no
// clock: Poll is called with the time, which is what makes a
// thirty-minute idle rule testable in a millisecond.
package watch

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/wow"
)

// The timing rules from the interface contract: a size increase after
// an idle gap longer than NewReportIdle starts a new report, and a
// report is completed once the file has been quiet for CompleteIdle.
const (
	NewReportIdle = 30 * time.Minute
	CompleteIdle  = 15 * time.Minute
)

// ReadChunk is how many bytes one Poll reads at most. It bounds the
// work per tick; the next tick takes the rest.
const ReadChunk = 1 << 20

// Kind is what one event tells the pipeline to do.
type Kind int

// The three boundaries a tail produces.
const (
	// Start opens a report at Offset in Path.
	Start Kind = iota
	// Append carries bytes at Offset in Path.
	Append
	// Complete closes the open report; Offset is its final offset.
	Complete
)

func (k Kind) String() string {
	switch k {
	case Start:
		return "start"
	case Append:
		return "append"
	case Complete:
		return "complete"
	}
	return "unknown"
}

// Event is one thing that happened to the log directory.
type Event struct {
	Kind   Kind
	Path   string
	Offset int64
	Data   []byte
	At     time.Time
}

// Options configures a Watcher.
type Options struct {
	// Dir is the install's Logs directory.
	Dir string
	// NewReportIdle and CompleteIdle default to the contract's values.
	NewReportIdle time.Duration
	CompleteIdle  time.Duration
	// ReadChunk defaults to ReadChunk.
	ReadChunk int
}

// Watcher follows the newest combat log in one directory.
type Watcher struct {
	o Options

	started    bool
	path       string
	info       os.FileInfo
	offset     int64
	open       bool // a report is open on path
	lastAppend time.Time
}

// New builds a watcher. Nothing is read until the first Poll.
func New(o Options) *Watcher {
	if o.NewReportIdle == 0 {
		o.NewReportIdle = NewReportIdle
	}
	if o.CompleteIdle == 0 {
		o.CompleteIdle = CompleteIdle
	}
	if o.ReadChunk == 0 {
		o.ReadChunk = ReadChunk
	}
	return &Watcher{o: o}
}

// Resume continues an open report rather than starting a new one. The
// pipeline calls it at startup for the report its state files say was
// in progress, so a restart mid-raid does not split the night in two.
func (w *Watcher) Resume(path string, offset int64, lastAppend time.Time) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("resume %s: %w", path, err)
	}
	if fi.Size() < offset {
		return fmt.Errorf("resume %s: the file is %d bytes, shorter than the saved offset %d",
			path, fi.Size(), offset)
	}
	w.started, w.path, w.info = true, path, fi
	w.offset, w.open, w.lastAppend = offset, true, lastAppend
	return nil
}

// Offset is the next byte the watcher will read.
func (w *Watcher) Offset() int64 { return w.offset }

// Path is the file being followed, empty before the first Poll.
func (w *Watcher) Path() string { return w.path }

// active picks the log to follow: the most recently modified match,
// with the path as a tiebreak so two files written in the same second
// resolve the same way on every machine.
func (w *Watcher) active() (string, os.FileInfo, error) {
	matches, err := filepath.Glob(filepath.Join(w.o.Dir, wow.LogGlob))
	if err != nil {
		return "", nil, err
	}
	sort.Strings(matches)
	var bestPath string
	var bestInfo os.FileInfo
	for _, m := range matches {
		fi, err := os.Stat(m)
		if err != nil || fi.IsDir() {
			continue
		}
		if bestInfo == nil || fi.ModTime().After(bestInfo.ModTime()) ||
			(fi.ModTime().Equal(bestInfo.ModTime()) && m > bestPath) {
			bestPath, bestInfo = m, fi
		}
	}
	if bestInfo == nil {
		return "", nil, nil
	}
	return bestPath, bestInfo, nil
}

// Poll reads whatever has happened since the last call and returns the
// events in the order the pipeline must apply them.
func (w *Watcher) Poll(now time.Time) ([]Event, error) {
	path, info, err := w.active()
	if err != nil {
		return nil, fmt.Errorf("scan %s: %w", w.o.Dir, err)
	}
	first := !w.started
	w.started = true

	if info == nil {
		return w.idle(now), nil
	}

	var out []Event
	rotated := w.info != nil && (path != w.path || !os.SameFile(info, w.info) || info.Size() < w.offset)
	if rotated || w.info == nil {
		if w.open {
			out = append(out, Event{Kind: Complete, Path: w.path, Offset: w.offset, At: now})
			w.open = false
		}
		w.path, w.info = path, info
		// A file that was already there when the companion started
		// holds a night we did not watch: begin at its end. A file
		// that appeared while we were running begins at zero.
		if first {
			w.offset = info.Size()
		} else {
			w.offset = 0
		}
	} else {
		w.info = info
	}

	if info.Size() <= w.offset {
		return append(out, w.idle(now)...), nil
	}

	data, err := w.read(info.Size())
	if err != nil {
		return out, err
	}
	if len(data) == 0 {
		return append(out, w.idle(now)...), nil
	}
	if w.open && now.Sub(w.lastAppend) > w.o.NewReportIdle {
		out = append(out, Event{Kind: Complete, Path: w.path, Offset: w.offset, At: now})
		w.open = false
	}
	if !w.open {
		out = append(out, Event{Kind: Start, Path: w.path, Offset: w.offset, At: now})
		w.open = true
	}
	out = append(out, Event{Kind: Append, Path: w.path, Offset: w.offset, Data: data, At: now})
	w.offset += int64(len(data))
	w.lastAppend = now
	return out, nil
}

// idle emits the completion when the file has been quiet long enough.
func (w *Watcher) idle(now time.Time) []Event {
	if !w.open || now.Sub(w.lastAppend) < w.o.CompleteIdle {
		return nil
	}
	w.open = false
	return []Event{{Kind: Complete, Path: w.path, Offset: w.offset, At: now}}
}

// read pulls at most ReadChunk bytes from the current offset.
func (w *Watcher) read(size int64) ([]byte, error) {
	f, err := os.Open(w.path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil // rotated away between the stat and the open
	}
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", w.path, err)
	}
	defer f.Close()
	n := size - w.offset
	if n > int64(w.o.ReadChunk) {
		n = int64(w.o.ReadChunk)
	}
	buf := make([]byte, n)
	read, err := f.ReadAt(buf, w.offset)
	if err != nil && read == 0 {
		return nil, fmt.Errorf("read %s at %d: %w", w.path, w.offset, err)
	}
	return buf[:read], nil
}

// ReadRange reads the bytes of one fight back out of the log so the
// bundle can carry their hash. A range the file no longer holds
// returns an error and the caller sends the bundle without a hash
// rather than dropping the fight.
func ReadRange(path string, start, end int64) ([]byte, error) {
	if end < start {
		return nil, fmt.Errorf("read %s: range %d..%d runs backwards", path, start, end)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, end-start)
	n, err := f.ReadAt(buf, start)
	if err != nil && int64(n) < end-start {
		return nil, fmt.Errorf("read %s at %d..%d: %w", path, start, end, err)
	}
	return buf[:n], nil
}
```

- [ ] **Step 4: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/watch/ && go test ./internal/watch/ -race -v`
Expected: PASS, nine tests. Read the names: they are the contract's timing rules one by one.

- [ ] **Step 5: Commit**

```bash
git add companion/internal/watch/
git commit -m "feat(companion): tail the combat log, with rotation and the idle rules" \
  -m "The Logs directory becomes a stream of report boundaries and bytes.
A new file, a rename to a timestamped sibling or a truncation closes
the open report and opens the next; fifteen quiet minutes complete
one; a size increase after a gap longer than thirty minutes starts a
new one.

Poll takes the time rather than reading a clock and owns no goroutine,
so every one of those rules has a test that runs in a millisecond. A
log that was already on disk when the companion started is joined at
its end: it holds a night nobody watched." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Task 8: Report state

**Files:**
- Create: `companion/internal/state/state.go`
- Test: `companion/internal/state/state_test.go`

**Interfaces:**
- Consumes: `character.Character` (Task 1).
- Produces:
  - `state.Report{Key, ReportID, LogPath string; StartOffset, Offset, RawSent int64; Visibility string; LoggingCharacter *character.Character; Zone, Title string; StartedAt, LastAppend time.Time; Fights int; Closed, Done bool; Session []byte; EngineVersion string}`.
  - `state.NewKey(now time.Time) string`, `state.Save(dir string, r Report) error`, `state.Load(dir, key string) (Report, error)`, `state.List(dir string) ([]Report, error)`, `state.Remove(dir, key string) error`.

`Key` is the local identity, minted before the server has seen the report so work can be queued while the network is down; `ReportID` is filled in when the API answers. `StartOffset` is the file offset the report begins at — the only bridge between the watcher's file offsets and the report-relative offsets everything else uses. `Session` is `session.State()` from the engine, which is what a restart resumes from. `Closed` means the log was completed locally; `Done` means the server took the completion.

`List` skips a file it cannot parse rather than failing: one corrupt report must not stop the companion from logging tonight.

- [ ] **Step 1: Write the failing test**

```go
package state

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

func TestAReportStateRoundTrips(t *testing.T) {
	dir := t.TempDir()
	want := Report{
		Key: "20261209-200000-deadbeef", ReportID: "rpt000000001",
		LogPath: "/w/Logs/WoWCombatLog.txt", StartOffset: 10, Offset: 4096, RawSent: 4096,
		Visibility: "public", StartedAt: t0, LastAppend: t0.Add(time.Minute),
		Fights: 3, Session: []byte{1, 2, 3}, EngineVersion: "0.1.0",
	}
	if err := Save(dir, want); err != nil {
		t.Fatal(err)
	}
	got, err := Load(dir, want.Key)
	if err != nil {
		t.Fatal(err)
	}
	if got.ReportID != want.ReportID || got.Offset != want.Offset || got.RawSent != want.RawSent ||
		got.Fights != want.Fights || string(got.Session) != string(want.Session) ||
		!got.LastAppend.Equal(want.LastAppend) {
		t.Fatalf("Load = %+v", got)
	}
}

func TestListIsSortedAndSkipsRubbish(t *testing.T) {
	dir := t.TempDir()
	for _, k := range []string{"20261209-210000-bb", "20261209-200000-aa"} {
		if err := Save(dir, Report{Key: k}); err != nil {
			t.Fatal(err)
		}
	}
	got, err := List(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Key != "20261209-200000-aa" {
		t.Fatalf("List = %+v", got)
	}
}

func TestSaveRefusesAnEmptyKeyAndRemoveIsIdempotent(t *testing.T) {
	dir := t.TempDir()
	if err := Save(dir, Report{}); err == nil {
		t.Fatal("Save accepted an empty key")
	}
	if err := Remove(dir, "never-existed"); err != nil {
		t.Fatalf("Remove of a missing report = %v", err)
	}
}

func TestNewKeyIsSortableAndUnique(t *testing.T) {
	a, b := NewKey(t0), NewKey(t0)
	if a == b {
		t.Fatal("two keys collided")
	}
	if len(a) != len("20261209-200000-deadbeef") {
		t.Fatalf("key = %q", a)
	}
	if NewKey(t0) >= NewKey(t0.Add(time.Hour)) {
		t.Fatal("keys do not sort by time")
	}
}

func TestLoadingAMissingOrBrokenReportIsAnError(t *testing.T) {
	dir := t.TempDir()
	if _, err := Load(dir, "never-existed"); err == nil {
		t.Fatal("Load accepted a missing report")
	}
	if err := os.WriteFile(filepath.Join(dir, "broken.json"), []byte("{not json"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(dir, "broken"); err == nil {
		t.Fatal("Load accepted a broken report")
	}
	// List skips it rather than failing: one bad file must not stop
	// tonight's raid from being logged.
	got, err := List(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("List = %+v", got)
	}
}

func TestListingADirectoryThatIsNotThereIsEmpty(t *testing.T) {
	got, err := List(filepath.Join(t.TempDir(), "state"))
	if err != nil || got != nil {
		t.Fatalf("List = %+v, %v", got, err)
	}
}

func TestRemoveDeletesOneReport(t *testing.T) {
	dir := t.TempDir()
	if err := Save(dir, Report{Key: "k"}); err != nil {
		t.Fatal(err)
	}
	if err := Remove(dir, "k"); err != nil {
		t.Fatal(err)
	}
	if got, _ := List(dir); len(got) != 0 {
		t.Fatalf("List = %+v", got)
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/state/`
Expected: FAIL, "undefined: Report".

- [ ] **Step 3: Write the implementation**

```go
// companion/internal/state/state.go
// Package state is one file per report under state/, holding the
// engine's serialised session and the offsets the uploader resumes
// from. It is what makes a restart mid-raid continue the same report
// instead of starting a second one.
package state

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
)

// Report is one report's durable state.
type Report struct {
	// Key is the local identity, assigned before the server has seen
	// the report, and the name of this file.
	Key string `json:"key"`
	// ReportID is the server's id, empty until the report is created.
	ReportID string `json:"report_id"`
	// LogPath is the combat log this report is being read from.
	LogPath string `json:"log_path"`
	// StartOffset is where the report began in that file; Offset is
	// the next byte to feed; RawSent is how far the raw chunks have
	// been queued.
	StartOffset int64 `json:"start_offset"`
	Offset      int64 `json:"offset"`
	RawSent     int64 `json:"raw_sent"`

	Visibility string `json:"visibility"`
	// LoggingCharacter attributes the report, captured when it starts
	// so a settings change mid-raid does not rewrite history.
	LoggingCharacter *character.Character `json:"logging_character,omitempty"`
	Zone             string               `json:"zone,omitempty"`
	Title            string               `json:"title,omitempty"`
	StartedAt        time.Time            `json:"started_at"`
	LastAppend       time.Time            `json:"last_append"`
	// Fights is how many fights have been queued, for the UI.
	Fights int `json:"fights"`
	// Closed is set when the log has been completed locally; Done is
	// set once the completion has been accepted by the server.
	Closed bool `json:"closed"`
	Done   bool `json:"done"`
	// Session is the engine's serialised session, base64 in JSON.
	Session []byte `json:"session,omitempty"`
	// EngineVersion is the engine that produced this report.
	EngineVersion string `json:"engine_version"`
}

// NewKey mints a local report key: a sortable timestamp and eight
// random hex characters, so two reports started in the same second on
// the same machine never collide.
func NewKey(now time.Time) string {
	var b [4]byte
	rand.Read(b[:])
	return fmt.Sprintf("%s-%x", now.UTC().Format("20060102-150405"), b)
}

func path(dir, key string) string { return filepath.Join(dir, key+".json") }

// Save writes one report's state atomically.
func Save(dir string, r Report) error {
	if r.Key == "" {
		return errors.New("state: the report key is empty")
	}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".state-*.json")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err := tmp.Write(append(b, '\n')); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(name, 0o600); err != nil {
		return err
	}
	return os.Rename(name, path(dir, r.Key))
}

// Load reads one report's state.
func Load(dir, key string) (Report, error) {
	b, err := os.ReadFile(path(dir, key))
	if err != nil {
		return Report{}, err
	}
	var r Report
	if err := json.Unmarshal(b, &r); err != nil {
		return Report{}, fmt.Errorf("parse %s: %w", path(dir, key), err)
	}
	return r, nil
}

// List reads every report state, oldest key first. A file that will
// not parse is skipped rather than fatal: one corrupt report must not
// stop the companion from logging tonight.
func List(dir string) ([]Report, error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".json") {
			keys = append(keys, strings.TrimSuffix(e.Name(), ".json"))
		}
	}
	sort.Strings(keys)
	out := make([]Report, 0, len(keys))
	for _, k := range keys {
		r, err := Load(dir, k)
		if err != nil {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}

// Remove deletes one report's state.
func Remove(dir, key string) error {
	err := os.Remove(path(dir, key))
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	return err
}
```

- [ ] **Step 4: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/state/ && go test ./internal/state/ -race`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add companion/internal/state/
git commit -m "feat(companion): one state file per report, holding the engine's session" \
  -m "A local key is minted before the server has seen the report, so a
raid can be parsed and queued with the network down and the report
created when it comes back. StartOffset is the only bridge between
the watcher's file offsets and the report-relative offsets the
contract uses everywhere else.

The serialised engine session lives here too: it is what makes a
restart mid-raid resume the same report rather than split the night
into two. List skips a file it cannot parse, because one corrupt
report must not cost tonight." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 9: The live pipeline — bytes in, queue entries out, queue entries up

**Files:**
- Create: `companion/internal/fixture/fixture.go`, `companion/internal/pipeline/pipeline.go`, `companion/internal/pipeline/drain.go`
- Test: `companion/internal/fixture/fixture_test.go`, `companion/internal/pipeline/pipeline_test.go`

**Interfaces:**
- Consumes: `client.FightBundle`, `client.MetricsRowsOf`, `client.RawRange`, `client.Live`, `client.Complete`, `client.CreateReport`, `client.Retryable`, `client.Unauthorized`, `(*Client).CreateReport/PutFight/PutLive/PutRaw/Complete` (Tasks 2–3); `config.Config` (Task 1); `queue.*` (Task 5); `watch.*` (Task 7); `state.*` (Task 8); `session.New/Restore/ReplayOffset/Feed/Close/Snapshot/State/Offset/Health/Version`, `parquet.Marshal`, `fight.Fight`, `summary.DefaultOptions`, `units.Options/RetailSpecClass/RetailSpecName` from the engine.
- Produces:
  - `fixture.Header`, `fixture.Zone`, `fixture.Encounter(n int) string`, `fixture.Heartbeat(n int) string`, `fixture.Log(n int) string`.
  - `pipeline.RawChunk` (4 MiB), `pipeline.LiveEvery` (5 s), `pipeline.StateEvery` (30 s).
  - `pipeline.Options{StateDir string; Client *client.Client; Queue *queue.Queue; Watch *watch.Watcher; Config config.Config; Log *slog.Logger; RawChunk int; LiveEvery, StateEvery time.Duration; NewKey func(time.Time) string}`, `pipeline.New(Options) (*Pipeline, error)`, `(*Pipeline).Close() error`.
  - `(*Pipeline).Tick(now time.Time) error`, `.Live(ctx, now) error`, `.Drain(ctx) error`, `.Run(ctx, every time.Duration) error`, `.Status() Status`.
  - `pipeline.Status{Logging bool; ReportKey, ReportID, LogPath string; Fights int; Offset int64; LastAppend time.Time; Queued int}`.

**Read this before writing the tests.** The engine emits a closed fight when the line *after* its last one arrives — an `ENCOUNTER_END` is buffered and the fight falls out on the following line, or on `Close`. A live raid's log never pauses long enough for that to matter, but a test that writes an encounter and stops would leave the fight open, which is why `fixture.Heartbeat` exists and why every test writes one after each encounter.

Parsing and uploading are two calls on purpose. `Tick` reads the file, feeds the engine and queues work; it never touches the network, so an outage cannot stop a raid from being parsed. `Drain` sends the queue, oldest first, and stops on the first retryable failure. `Live` is neither: it goes straight out and is dropped on failure.

Three subtleties the reference implementation had to get right, each with a test:

- **The open report is created as soon as the network allows**, not when its first fight closes, because the live snapshots of the very first pull need somewhere to go.
- **A session that has not settled a layout cannot be serialised** — the engine says so — and does not need to be: nothing has closed, so a restart re-feeds the report from its first byte. `save` tolerates it and `resume` handles the state file with no session blob.
- **A restart must not lose buffered raw bytes.** The bytes between the last queued 4 MiB chunk and the stopping point are re-read out of the log by `refillRaw`; without it the server could never reassemble the raw stream. `bufferRaw` skips anything already buffered, which is also what makes the engine's replay bytes harmless.

- [ ] **Step 1: Write the fixture and prove the engine reads it**

```go
// companion/internal/fixture/fixture.go
// Package fixture writes the combat log the companion's tests parse.
// Every line is hand-written in the v16 shapes the engine verified, and
// uses invented characters — the engine's fixture policy applies here
// too: nothing is copied out of a downloaded log.
package fixture

import (
	"fmt"
	"strings"
	"time"
)

// Header is the log's first line: v16, advanced logging on.
const Header = "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1," +
	"BUILD_VERSION,9.0.2,PROJECT_ID,1\n"

// Zone is a zone change into the fixture's instance.
const Zone = "9/26 20:10:00.500  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n"

// player and boss are invented; see logs/README.md for the cast.
const (
	player   = `Player-4184-000000A3,"Morrowlyn-Nightslayer",0x512,0x0`
	boss     = `Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0`
	advanced = `Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,41320,44000,` +
		`612,0,1955,0,0,0,0,0,-1487.02,6409.71,1675,1.2044,45`
	damage = `1484,1390,-1,16,0,0,0,1,nil,nil`
)

// Encounter is one boss fight: a start, a few damage events, a kill and
// an end. Index n moves it along the clock so the fights do not
// overlap, and the encounter id changes with it.
func Encounter(n int) string {
	start := time.Date(2026, 9, 26, 20, 12, 0, 0, time.UTC).Add(time.Duration(n) * 5 * time.Minute)
	at := func(d time.Duration) string {
		t := start.Add(d)
		return fmt.Sprintf("9/26 %02d:%02d:%02d.000  ", t.Hour(), t.Minute(), t.Second())
	}
	id := 9001 + n
	name := fmt.Sprintf("Warden Kelthas %d", n)
	var b strings.Builder
	fmt.Fprintf(&b, "%sENCOUNTER_START,%d,%q,8,5,2284\n", at(0), id, name)
	for i := range 3 {
		fmt.Fprintf(&b, "%sSPELL_DAMAGE,%s,%s,116,\"Frostbolt\",0x10,%s,%s\n",
			at(time.Duration(i+1)*time.Second), player, boss, advanced, damage)
	}
	fmt.Fprintf(&b, "%sUNIT_DIED,0000000000000000,nil,0x80000000,0x80000000,%s,0\n", at(30*time.Second), boss)
	fmt.Fprintf(&b, "%sENCOUNTER_END,%d,%q,8,5,1\n", at(31*time.Second), id, name)
	return b.String()
}

// Heartbeat is one line after a fight. The engine emits a closed
// fight when the line after its last one arrives, so a test that
// writes an encounter and stops would leave that fight open; a real
// raid's log never stops for long enough to notice.
func Heartbeat(n int) string {
	t := time.Date(2026, 9, 26, 20, 12, 45, 0, time.UTC).Add(time.Duration(n) * 5 * time.Minute)
	return fmt.Sprintf("9/26 %02d:%02d:%02d.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n",
		t.Hour(), t.Minute(), t.Second())
}

// Log is a whole night: the header, a zone change and n encounters.
func Log(n int) string {
	var b strings.Builder
	b.WriteString(Header)
	b.WriteString(Zone)
	for i := range n {
		b.WriteString(Encounter(i))
		b.WriteString(Heartbeat(i))
	}
	return b.String()
}
```

```go
package fixture

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// TestTheFixtureParsesIntoWholeFights is the guard on every other
// test in the companion: if the engine cannot read this log, no
// pipeline test below means anything.
func TestTheFixtureParsesIntoWholeFights(t *testing.T) {
	o := session.Options{
		ReportID: "fixture", KeepEvents: true,
		Base:    time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC),
		Units:   units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:   fight.DefaultOptions(),
		Summary: summary.DefaultOptions(),
	}
	s := session.New(o)
	log := Log(3)
	res, err := s.Feed([]byte(log), 0)
	if err != nil {
		t.Fatal(err)
	}
	tail, err := s.Close()
	if err != nil {
		t.Fatal(err)
	}
	closed := append(res.Closed, tail.Closed...)
	var encounters []string
	for _, c := range closed {
		if c.Fight.Kind == fight.Encounter {
			encounters = append(encounters, c.Fight.Name)
		}
	}
	if len(encounters) != 3 {
		t.Fatalf("encounters = %v, want three", encounters)
	}
	if !strings.HasPrefix(encounters[0], "Warden Kelthas") {
		t.Errorf("first encounter = %q", encounters[0])
	}
	h := s.Health()
	if h.ParseErrors != 0 {
		t.Errorf("the fixture has %d parse errors", h.ParseErrors)
	}
	if !h.AdvancedLogging {
		t.Error("the fixture must have advanced logging on")
	}
	for _, c := range closed {
		if c.Fight.Kind == fight.Encounter && len(c.Summary.Roster) == 0 {
			t.Errorf("fight %d has an empty roster", c.Fight.Index)
		}
	}
}
```

- [ ] **Step 2: Run the fixture test**

Run: `cd companion && go test ./internal/fixture/ -v`
Expected: PASS. If this fails, nothing below it means anything: it is the guard that says the hand-written log is a log the engine can read, with three encounters, no parse errors, advanced logging on and a roster per fight.

- [ ] **Step 3: Write the failing pipeline test**

```go
package pipeline_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/pipeline"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// rig is one companion wired to one fake API over one temporary game
// directory.
type rig struct {
	t        *testing.T
	srv      *fakeapi.Server
	dirs     struct{ logs, queue, state string }
	log      string
	pipe     *pipeline.Pipeline
	q        *queue.Queue
	clock    time.Time
	rawChunk int
}

func newRig(t *testing.T, rawChunk int) *rig {
	t.Helper()
	root := t.TempDir()
	r := &rig{t: t, srv: fakeapi.New(), clock: t0, rawChunk: rawChunk}
	r.dirs.logs = filepath.Join(root, "Logs")
	r.dirs.queue = filepath.Join(root, "queue")
	r.dirs.state = filepath.Join(root, "state")
	for _, d := range []string{r.dirs.logs, r.dirs.queue, r.dirs.state} {
		if err := os.MkdirAll(d, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	r.log = filepath.Join(r.dirs.logs, "WoWCombatLog.txt")
	t.Cleanup(r.srv.Close)
	r.open()
	// The companion starts before the player logs in: one tick over
	// an empty Logs directory, so the log it later sees is a new one
	// rather than a night it never watched.
	if err := r.pipe.Tick(r.clock); err != nil {
		t.Fatal(err)
	}
	return r
}

// open builds a fresh pipeline over the same directories, which is
// what a restart looks like.
func (r *rig) open() {
	r.t.Helper()
	if r.pipe != nil {
		r.pipe.Close()
	}
	c, err := client.New(client.Options{
		BaseURL: r.srv.URL,
		Token:   func() string { return fakeapi.Token },
		Retry:   client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
	})
	if err != nil {
		r.t.Fatal(err)
	}
	q, err := queue.Open(r.dirs.queue)
	if err != nil {
		r.t.Fatal(err)
	}
	r.q = q
	p, err := pipeline.New(pipeline.Options{
		StateDir: r.dirs.state,
		Client:   c,
		Queue:    q,
		Watch:    watch.New(watch.Options{Dir: r.dirs.logs}),
		Config:   config.Config{APIBaseURL: r.srv.URL, SiteBaseURL: r.srv.URL, ReportVisibility: "public"},
		RawChunk: r.rawChunk,
		// Save the session on every tick so a restart in a test
		// resumes from the same place a long-running companion would.
		StateEvery: -1,
	})
	if err != nil {
		r.t.Fatal(err)
	}
	r.pipe = p
	r.t.Cleanup(func() { p.Close() })
}

// write appends to the game's log and advances the clock.
func (r *rig) write(text string) {
	r.t.Helper()
	f, err := os.OpenFile(r.log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		r.t.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		r.t.Fatal(err)
	}
	f.Close()
	r.clock = r.clock.Add(time.Second)
	if err := os.Chtimes(r.log, r.clock, r.clock); err != nil {
		r.t.Fatal(err)
	}
}

// tick parses and uploads once.
func (r *rig) tick() error {
	r.t.Helper()
	r.clock = r.clock.Add(time.Second)
	if err := r.pipe.Tick(r.clock); err != nil {
		return err
	}
	return r.pipe.Drain(r.t.Context())
}

// settle ticks until the log has been quiet long enough to complete
// the report and the queue has drained.
func (r *rig) settle() {
	r.t.Helper()
	if err := r.tick(); err != nil {
		r.t.Fatalf("tick: %v", err)
	}
	r.clock = r.clock.Add(watch.CompleteIdle + time.Minute)
	if err := r.pipe.Tick(r.clock); err != nil {
		r.t.Fatalf("tick: %v", err)
	}
	if err := r.pipe.Drain(r.t.Context()); err != nil {
		r.t.Fatalf("drain: %v", err)
	}
}

func TestAWholeNightArrivesAsFightsInOrder(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	for i := range 3 {
		r.write(fixture.Encounter(i) + fixture.Heartbeat(i))
		if err := r.tick(); err != nil {
			t.Fatal(err)
		}
	}
	r.settle()

	var fights []string
	var completed bool
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights = append(fights, e)
		}
		if strings.HasPrefix(e, "complete ") {
			completed = true
		}
	}
	if len(fights) < 3 {
		t.Fatalf("the server saw %v", r.srv.Order())
	}
	for i := 1; i < len(fights); i++ {
		if fights[i-1] >= fights[i] {
			t.Fatalf("fights arrived out of order: %v", fights)
		}
	}
	if !completed {
		t.Errorf("the report was never completed: %v", r.srv.Order())
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Errorf("%d items are still queued", n)
	}
	for _, rep := range r.srv.Reports() {
		if len(rep.Raw) == 0 {
			t.Error("no raw chunks were uploaded")
		}
		for _, f := range rep.Fights {
			if len(f.Events) == 0 {
				t.Errorf("fight %d has no parquet", f.Index)
			}
			if len(f.Metrics) == 0 {
				t.Errorf("fight %d has no metrics rows", f.Index)
			}
			if len(f.RawRange.SHA256) != 64 {
				t.Errorf("fight %d raw hash = %q", f.Index, f.RawRange.SHA256)
			}
		}
	}
}

func TestRawIsChunkedAtTheConfiguredSizeAndAddressedByOffset(t *testing.T) {
	r := newRig(t, 512)
	r.write(fixture.Log(2))
	for range 4 {
		if err := r.tick(); err != nil {
			t.Fatal(err)
		}
	}
	r.settle()
	for _, rep := range r.srv.Reports() {
		if len(rep.Raw) < 2 {
			t.Fatalf("raw chunks = %d, want several at 512 bytes", len(rep.Raw))
		}
		for offset, body := range rep.Raw {
			if offset%512 != 0 {
				t.Errorf("chunk at offset %d is not on a 512-byte boundary", offset)
			}
			if len(body) == 0 {
				t.Errorf("chunk at %d is empty", offset)
			}
		}
		if _, first := rep.Raw[0]; !first {
			t.Error("there is no chunk at offset zero")
		}
	}
}

func TestAnOutageQueuesAndTheOrderSurvivesTheReconnect(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}

	r.srv.Offline(true)
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	r.write(fixture.Encounter(2) + fixture.Heartbeat(2))
	if err := r.tick(); err == nil {
		t.Fatal("draining while offline succeeded")
	}
	if n, _ := r.q.Len(); n == 0 {
		t.Fatal("nothing was queued during the outage")
	}

	r.srv.Offline(false)
	r.settle()
	var fights []string
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights = append(fights, e)
		}
	}
	if len(fights) < 3 {
		t.Fatalf("after the reconnect the server saw %v", r.srv.Order())
	}
	for i := 1; i < len(fights); i++ {
		if fights[i-1] >= fights[i] {
			t.Fatalf("fights arrived out of order: %v", fights)
		}
	}
}

func TestARestartResumesTheSameReport(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	before, err := state.List(r.dirs.state)
	if err != nil || len(before) != 1 {
		t.Fatalf("state = %+v, %v", before, err)
	}
	key, reportID := before[0].Key, before[0].ReportID

	r.open() // the companion restarts
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.settle()

	after, err := state.List(r.dirs.state)
	if err != nil {
		t.Fatal(err)
	}
	if len(after) != 1 {
		t.Fatalf("the restart started a second report: %+v", after)
	}
	if after[0].Key != key {
		t.Fatalf("report key = %q, want %q", after[0].Key, key)
	}
	if len(r.srv.Reports()) != 1 {
		t.Fatalf("the server holds %d reports", len(r.srv.Reports()))
	}
	if r.srv.Reports()[reportID] == nil {
		t.Fatalf("report %q is gone", reportID)
	}
	if got := len(r.srv.Reports()[reportID].Fights); got < 2 {
		t.Errorf("the resumed report holds %d fights, want both", got)
	}
}

func TestALiveSnapshotGoesUpWhileAFightIsOpen(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	// An encounter with no end: the fight stays open.
	open := strings.SplitAfter(fixture.Encounter(0), "\n")
	r.write(strings.Join(open[:len(open)-2], ""))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.clock = r.clock.Add(10 * time.Second)
	if err := r.pipe.Live(t.Context(), r.clock); err != nil {
		t.Fatal(err)
	}
	var live int
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "live ") {
			live++
		}
	}
	if live != 1 {
		t.Fatalf("live snapshots = %d, order = %v", live, r.srv.Order())
	}
	st := r.pipe.Status()
	if !st.Logging || st.ReportID == "" {
		t.Fatalf("status = %+v", st)
	}
}

func TestNothingHappensWithoutALogFile(t *testing.T) {
	r := newRig(t, 1<<20)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	if len(r.srv.Order()) != 0 {
		t.Fatalf("the server saw %v with no log file", r.srv.Order())
	}
	if st := r.pipe.Status(); st.Logging {
		t.Errorf("status = %+v", st)
	}
}

func TestRunTicksUntilTheContextEnds(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan error, 1)
	go func() { done <- r.pipe.Run(ctx, time.Millisecond) }()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if len(r.srv.Order()) > 0 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("Run = %v", err)
	}
	if len(r.srv.Order()) == 0 {
		t.Fatal("Run uploaded nothing")
	}
}

func TestAFightTheServerRefusesIsDroppedSoTheQueueMoves(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	// A verification mismatch: permanent, so the fight is dropped
	// with a log line rather than retried until the heat death.
	r.srv.RefuseFights(true)
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	if err := r.tick(); err != nil {
		t.Fatalf("a refused fight stopped the drain: %v", err)
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Fatalf("%d items are stuck behind the refused fight", n)
	}
	r.srv.RefuseFights(false)
	r.write(fixture.Encounter(2) + fixture.Heartbeat(2))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	var fights int
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights++
		}
	}
	if fights != 2 {
		t.Fatalf("the server stored %d fights, want the two it accepted: %v", fights, r.srv.Order())
	}
}

func TestAQueuedItemWithNoReportStateIsDropped(t *testing.T) {
	r := newRig(t, 1<<20)
	if _, err := r.q.Enqueue(queue.Item{
		Kind: queue.Fight, ReportKey: "a-report-that-never-existed", ContentType: "text/plain",
	}, []byte("x")); err != nil {
		t.Fatal(err)
	}
	if err := r.pipe.Drain(t.Context()); err != nil {
		t.Fatal(err)
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Fatalf("%d orphans are still queued", n)
	}
}

func TestLiveAndCompleteDoNothingWithoutAnOpenReport(t *testing.T) {
	r := newRig(t, 1<<20)
	if err := r.pipe.Live(t.Context(), r.clock); err != nil {
		t.Fatal(err)
	}
	// A completion with nothing open is what the watcher emits after
	// the companion starts on a quiet machine.
	r.clock = r.clock.Add(watch.CompleteIdle + time.Minute)
	if err := r.pipe.Tick(r.clock); err != nil {
		t.Fatal(err)
	}
	if len(r.srv.Order()) != 0 {
		t.Fatalf("the server saw %v", r.srv.Order())
	}
}
```

- [ ] **Step 4: Run it to verify it fails**

Run: `cd companion && go test ./internal/pipeline/`
Expected: FAIL, "undefined: pipeline.New".

- [ ] **Step 5: Write `internal/pipeline/pipeline.go`**

```go
// companion/internal/pipeline/pipeline.go
// Package pipeline is the live path: bytes from the watcher into the
// engine, closed fights and raw chunks into the queue, the queue into
// the API. Parsing and uploading are two separate calls — Tick and
// Drain — so a network outage stops the uploads and never the parse,
// and so a test can drive both by hand.
//
// Every offset the companion sends is relative to the report's own
// byte stream, which begins at zero when the report begins. The
// watcher works in file offsets; state.Report.StartOffset is the one
// place the two meet.
package pipeline

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Defaults from the interface contract: raw is streamed in 4 MiB
// chunks and the live snapshot goes up every five seconds.
const (
	RawChunk  = 4 << 20
	LiveEvery = 5 * time.Second
	// StateEvery bounds how often the engine's session is serialised
	// to disk between fights. Fight closes always save.
	StateEvery = 30 * time.Second
)

// Options configures a Pipeline.
type Options struct {
	StateDir string
	Client   *client.Client
	Queue    *queue.Queue
	Watch    *watch.Watcher
	Config   config.Config
	Log      *slog.Logger

	RawChunk   int
	LiveEvery  time.Duration
	StateEvery time.Duration
	// NewKey mints a local report key. Tests pin it.
	NewKey func(time.Time) string
}

// Pipeline owns the open report.
type Pipeline struct {
	o   Options
	enc *zstd.Encoder

	sess *session.Session
	cur  *state.Report
	// raw holds report-relative bytes not yet handed to the queue;
	// rawAt is the offset of raw[0].
	raw      []byte
	rawAt    int64
	lastLive time.Time
	lastSave time.Time
}

// New builds a pipeline and resumes the report that was open when the
// companion last stopped, if there is one.
func New(o Options) (*Pipeline, error) {
	if o.Client == nil || o.Queue == nil || o.Watch == nil {
		return nil, errors.New("pipeline: the client, queue and watcher are all required")
	}
	if o.RawChunk == 0 {
		o.RawChunk = RawChunk
	}
	if o.LiveEvery == 0 {
		o.LiveEvery = LiveEvery
	}
	if o.StateEvery == 0 {
		o.StateEvery = StateEvery
	}
	if o.NewKey == nil {
		o.NewKey = state.NewKey
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	enc, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	if err != nil {
		return nil, err
	}
	p := &Pipeline{o: o, enc: enc}
	if err := p.resume(); err != nil {
		return nil, err
	}
	return p, nil
}

// Close releases the compressor.
func (p *Pipeline) Close() error {
	p.enc.Close()
	return nil
}

// sessionOptions are the engine settings the companion parses with.
// KeepEvents is on because every closed fight is written as Parquet.
func sessionOptions(reportKey string, base time.Time) session.Options {
	o := session.Options{
		ReportID:   reportKey,
		Base:       base,
		Infer:      true,
		KeepEvents: true,
		Units:      units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:      fight.DefaultOptions(),
		Summary:    summary.DefaultOptions(),
	}
	o.Summary.SpecNames = units.RetailSpecName
	return o
}

// resume restores the report that was still open, if any.
func (p *Pipeline) resume() error {
	reports, err := state.List(p.o.StateDir)
	if err != nil {
		return err
	}
	for i := len(reports) - 1; i >= 0; i-- {
		r := reports[i]
		if r.Closed {
			continue
		}
		if len(r.Session) == 0 {
			// Nothing had closed when the companion stopped: read the
			// report again from its first byte. Re-sent bytes are
			// ignored by offset on both sides.
			if err := p.o.Watch.Resume(r.LogPath, r.StartOffset, r.LastAppend); err != nil {
				p.o.Log.Warn("could not reopen the log; starting a new report",
					"component", "pipeline", "report", r.Key, "err", err.Error())
				r.Closed = true
				state.Save(p.o.StateDir, r)
				continue
			}
			p.sess, p.cur = session.New(sessionOptions(r.Key, r.StartedAt)), &r
			p.refillRaw(r)
			p.o.Log.Info("restarted a report from its first byte",
				"component", "pipeline", "report", r.Key)
			return nil
		}
		s, err := session.Restore(sessionOptions(r.Key, r.StartedAt), r.Session)
		if err != nil {
			p.o.Log.Warn("could not restore a report; starting a new one",
				"component", "pipeline", "report", r.Key, "err", err.Error())
			r.Closed = true
			state.Save(p.o.StateDir, r)
			continue
		}
		replay, err := session.ReplayOffset(r.Session)
		if err != nil {
			return err
		}
		if err := p.o.Watch.Resume(r.LogPath, r.StartOffset+replay, r.LastAppend); err != nil {
			p.o.Log.Warn("could not resume the log file; starting a new report",
				"component", "pipeline", "report", r.Key, "err", err.Error())
			r.Closed = true
			state.Save(p.o.StateDir, r)
			continue
		}
		p.sess, p.cur = s, &r
		p.refillRaw(r)
		p.o.Log.Info("resumed a report", "component", "pipeline", "report", r.Key,
			"report_id", r.ReportID, "replay_from", replay)
		return nil
	}
	return nil
}

// Tick reads whatever the game has written and turns it into queued
// work. It never talks to the network.
func (p *Pipeline) Tick(now time.Time) error {
	events, err := p.o.Watch.Poll(now)
	if err != nil {
		return err
	}
	for _, e := range events {
		switch e.Kind {
		case watch.Start:
			if err := p.start(e, now); err != nil {
				return err
			}
		case watch.Append:
			if err := p.append(e, now); err != nil {
				return err
			}
		case watch.Complete:
			if err := p.complete(now); err != nil {
				return err
			}
		}
	}
	if p.cur != nil && now.Sub(p.lastSave) >= p.o.StateEvery {
		if err := p.save(); err != nil {
			return err
		}
		p.lastSave = now
	}
	return nil
}

func (p *Pipeline) start(e watch.Event, now time.Time) error {
	key := p.o.NewKey(now)
	p.cur = &state.Report{
		Key: key, LogPath: e.Path, StartOffset: e.Offset,
		Visibility: p.o.Config.ReportVisibility, LoggingCharacter: p.o.Config.LoggingCharacter,
		StartedAt: now, LastAppend: now,
		EngineVersion: session.Version,
	}
	p.sess = session.New(sessionOptions(key, now))
	p.raw, p.rawAt, p.lastLive, p.lastSave = nil, 0, time.Time{}, now
	p.o.Log.Info("a report started", "component", "pipeline", "report", key,
		"path", e.Path, "file_offset", e.Offset)
	return p.save()
}

func (p *Pipeline) append(e watch.Event, now time.Time) error {
	if p.cur == nil {
		return nil // bytes with no open report; the watcher opens one first
	}
	rel := e.Offset - p.cur.StartOffset
	res, err := p.sess.Feed(e.Data, rel)
	if err != nil {
		return fmt.Errorf("feed %s at %d: %w", e.Path, rel, err)
	}
	p.bufferRaw(rel, e.Data)
	p.cur.Offset = p.sess.Offset()
	p.cur.LastAppend = now
	for _, c := range res.Closed {
		if err := p.enqueueFight(c); err != nil {
			return err
		}
	}
	if err := p.flushRaw(false); err != nil {
		return err
	}
	if len(res.Closed) > 0 {
		if err := p.save(); err != nil {
			return err
		}
		p.lastSave = now
	}
	return nil
}

// bufferRaw appends the part of a chunk that has not been buffered
// already. A restart re-feeds bytes the engine needs to rebuild the
// open fight, and those bytes must not be uploaded twice.
func (p *Pipeline) bufferRaw(rel int64, data []byte) {
	next := p.rawAt + int64(len(p.raw))
	end := rel + int64(len(data))
	if end <= next {
		return
	}
	if rel < next {
		data = data[next-rel:]
		rel = next
	}
	if len(p.raw) == 0 {
		p.rawAt = rel
	}
	p.raw = append(p.raw, data...)
}

// refillRaw rebuilds the unqueued tail of the raw stream after a
// restart. Without it the bytes between the last queued chunk and the
// stopping point would never be uploaded, and the server could not
// reassemble the log.
func (p *Pipeline) refillRaw(r state.Report) {
	p.rawAt, p.raw = r.RawSent, nil
	if r.Offset <= r.RawSent {
		return
	}
	buf, err := watch.ReadRange(r.LogPath, r.StartOffset+r.RawSent, r.StartOffset+r.Offset)
	if err != nil {
		p.o.Log.Warn("could not reread the unsent raw bytes", "component", "pipeline",
			"report", r.Key, "err", err.Error())
		return
	}
	p.raw = buf
}

// flushRaw queues whole chunks, and on completion the remainder too.
func (p *Pipeline) flushRaw(final bool) error {
	for len(p.raw) >= p.o.RawChunk || (final && len(p.raw) > 0) {
		n := min(len(p.raw), p.o.RawChunk)
		packed := p.enc.EncodeAll(p.raw[:n], nil)
		if _, err := p.o.Queue.Enqueue(queue.Item{
			Kind: queue.Raw, ReportKey: p.cur.Key, Offset: p.rawAt,
		}, packed); err != nil {
			return err
		}
		p.rawAt += int64(n)
		p.cur.RawSent = p.rawAt
		p.raw = p.raw[n:]
	}
	if len(p.raw) == 0 {
		p.raw = nil
	}
	return nil
}

// enqueueFight builds one closed fight's bundle and queues it.
func (p *Pipeline) enqueueFight(c session.Closed) error {
	events, err := parquet.Marshal(c.Events)
	if err != nil {
		return fmt.Errorf("write fight %d as parquet: %w", c.Fight.Index, err)
	}
	b := client.FightBundle{
		Summary: c.Summary,
		Events:  events,
		Metrics: client.MetricsRowsOf(c.Fight, c.Summary),
		RawRange: client.RawRange{
			StartOffset: c.Fight.StartOffset,
			EndOffset:   c.Fight.EndOffset,
			SHA256:      p.rawHash(c.Fight),
		},
	}
	ct, body, err := b.Encode()
	if err != nil {
		return err
	}
	if _, err := p.o.Queue.Enqueue(queue.Item{
		Kind: queue.Fight, ReportKey: p.cur.Key,
		FightIndex: c.Fight.Index, ContentType: ct,
	}, body); err != nil {
		return err
	}
	p.cur.Fights++
	if p.cur.Zone == "" && c.Fight.Zone != "" {
		p.cur.Zone = c.Fight.Zone
	}
	p.o.Log.Info("a fight closed", "component", "pipeline", "report", p.cur.Key,
		"fight", c.Fight.Index, "name", c.Fight.Name, "bundle_bytes", len(body))
	return nil
}

// rawHash is the hash of the fight's bytes, read back out of the log.
// A log that has already rotated away yields no hash rather than no
// fight: the server's later raw-sample check is what the hash serves,
// and a missing hash costs that check, not the report.
func (p *Pipeline) rawHash(f fight.Fight) string {
	raw, err := watch.ReadRange(p.cur.LogPath,
		p.cur.StartOffset+f.StartOffset, p.cur.StartOffset+f.EndOffset)
	if err != nil {
		p.o.Log.Warn("could not hash a fight's raw bytes", "component", "pipeline",
			"report", p.cur.Key, "fight", f.Index, "err", err.Error())
		return ""
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

// complete closes the engine's session, queues the last fights, the
// last raw chunk and the completion, and forgets the open report.
func (p *Pipeline) complete(now time.Time) error {
	if p.cur == nil {
		return nil
	}
	res, err := p.sess.Close()
	if err != nil {
		return fmt.Errorf("close the session for %s: %w", p.cur.Key, err)
	}
	for _, c := range res.Closed {
		if err := p.enqueueFight(c); err != nil {
			return err
		}
	}
	p.cur.Offset = p.sess.Offset()
	if err := p.flushRaw(true); err != nil {
		return err
	}
	body, err := marshalComplete(client.Complete{
		FinalOffset:   p.cur.Offset,
		EngineVersion: session.Version,
		Health:        p.sess.Health(),
	})
	if err != nil {
		return err
	}
	if _, err := p.o.Queue.Enqueue(queue.Item{
		Kind: queue.Complete, ReportKey: p.cur.Key,
	}, body); err != nil {
		return err
	}
	p.cur.Closed = true
	p.cur.Session = nil
	if err := p.save(); err != nil {
		return err
	}
	p.o.Log.Info("a report completed", "component", "pipeline", "report", p.cur.Key,
		"fights", p.cur.Fights, "final_offset", p.cur.Offset)
	p.sess, p.cur, p.raw = nil, nil, nil
	return nil
}

// save writes the open report's state, including the engine's session
// so a restart resumes rather than restarts.
func (p *Pipeline) save() error {
	if p.cur == nil {
		return nil
	}
	if p.sess != nil && !p.cur.Closed {
		// A session that has not settled a layout yet cannot be
		// serialised, and does not need to be: nothing has closed,
		// so a restart re-feeds the report from its first byte.
		if blob, err := p.sess.State(); err == nil {
			p.cur.Session = blob
		} else {
			p.cur.Session = nil
			p.o.Log.Debug("the session is not resumable yet", "component", "pipeline",
				"report", p.cur.Key, "err", err.Error())
		}
	}
	return state.Save(p.o.StateDir, *p.cur)
}

// Live uploads the running summary of the fight in progress. It is not
// queued: the next snapshot replaces a failed one. It is silent when
// no fight is open or the report has no server id yet.
func (p *Pipeline) Live(ctx context.Context, now time.Time) error {
	if p.cur == nil || p.cur.ReportID == "" || now.Sub(p.lastLive) < p.o.LiveEvery {
		return nil
	}
	f, sum, open := p.sess.Snapshot()
	if !open {
		return nil
	}
	p.lastLive = now
	err := p.o.Client.PutLive(ctx, p.cur.ReportID, f.Index, client.Live{
		Summary:   sum,
		ElapsedMS: now.Sub(f.Start).Milliseconds(),
		UpdatedAt: now.UTC(),
	})
	if err != nil {
		p.o.Log.Debug("a live snapshot did not land", "component", "pipeline",
			"report", p.cur.ReportID, "fight", f.Index, "err", err.Error())
	}
	return nil
}
```

- [ ] **Step 6: Write `internal/pipeline/drain.go`**

```go
// companion/internal/pipeline/drain.go
// Draining the queue. One item at a time, oldest first, stopping on
// the first failure worth retrying — which is what puts the fights on
// the server in the order the raid fought them.
package pipeline

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

func marshalComplete(in client.Complete) ([]byte, error) { return json.Marshal(in) }

// Drain sends queued work until the queue is empty or an upload fails
// in a way that is worth retrying. The failing item stays at the head,
// so nothing behind it overtakes it.
func (p *Pipeline) Drain(ctx context.Context) error {
	// The open report is created as soon as the network allows, not
	// when its first fight closes, so the live snapshots of the very
	// first pull have somewhere to go.
	if p.cur != nil && p.cur.ReportID == "" {
		if err := p.create(ctx, p.cur); err != nil {
			return err
		}
		if err := state.Save(p.o.StateDir, *p.cur); err != nil {
			return err
		}
	}
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		l, err := p.o.Queue.Head()
		if errors.Is(err, queue.ErrEmpty) {
			return nil
		}
		if err != nil {
			return err
		}
		if err := p.send(ctx, l); err != nil {
			return err
		}
	}
}

// create asks the API for a report id and records it on rep and, when
// rep is the open report, on the pipeline.
func (p *Pipeline) create(ctx context.Context, rep *state.Report) error {
	created, err := p.o.Client.CreateReport(ctx, client.CreateReport{
		Title:      rep.Title,
		Visibility: rep.Visibility,
		Zone:       rep.Zone,
		// The attribution triple carries the ruleset segment through
		// untouched; the API maps it.
		LoggingCharacter: rep.LoggingCharacter,
	})
	if err != nil {
		return err
	}
	rep.ReportID = created.ID
	if p.cur != nil && p.cur.Key == rep.Key {
		p.cur.ReportID = created.ID
	}
	p.o.Log.Info("a report was created", "component", "uploader",
		"report", rep.Key, "report_id", rep.ReportID)
	return nil
}

// send performs one queued item.
func (p *Pipeline) send(ctx context.Context, l *queue.Lease) error {
	rep, err := state.Load(p.o.StateDir, l.Item.ReportKey)
	if err != nil {
		p.o.Log.Error("dropping a queued item whose report state is gone",
			"component", "uploader", "report", l.Item.ReportKey,
			"kind", string(l.Item.Kind), "err", err.Error())
		return l.Drop()
	}
	if rep.ReportID == "" {
		if cerr := p.create(ctx, &rep); cerr != nil {
			return p.failed(l, cerr)
		}
		if err := state.Save(p.o.StateDir, rep); err != nil {
			return err
		}
	}

	switch l.Item.Kind {
	case queue.Fight:
		_, err = p.o.Client.PutFight(ctx, rep.ReportID, l.Item.FightIndex,
			l.Item.ContentType, l.Body)
	case queue.Raw:
		_, err = p.o.Client.PutRaw(ctx, rep.ReportID, l.Item.Offset, l.Body)
	case queue.Complete:
		var in client.Complete
		if jerr := json.Unmarshal(l.Body, &in); jerr != nil {
			p.o.Log.Error("dropping an unreadable completion", "component", "uploader",
				"report", rep.Key, "err", jerr.Error())
			return l.Drop()
		}
		err = p.o.Client.Complete(ctx, rep.ReportID, in)
	default:
		p.o.Log.Error("dropping a queued item of an unknown kind",
			"component", "uploader", "kind", string(l.Item.Kind))
		return l.Drop()
	}
	if err != nil {
		return p.failed(l, err)
	}
	if l.Item.Kind == queue.Complete {
		rep.Done = true
		if err := state.Save(p.o.StateDir, rep); err != nil {
			return err
		}
	}
	return l.Ack()
}

// failed decides what a failed upload means. Anything retryable leaves
// the item where it is and stops the drain; anything the server will
// refuse again forever is dropped with a loud log line, because a
// poisoned item must not hold a raid night hostage.
func (p *Pipeline) failed(l *queue.Lease, cause error) error {
	if client.Retryable(cause) || client.Unauthorized(cause) {
		if err := l.Fail(cause); err != nil {
			return err
		}
		return cause
	}
	p.o.Log.Error("dropping an item the server refused", "component", "uploader",
		"report", l.Item.ReportKey, "kind", string(l.Item.Kind),
		"fight", l.Item.FightIndex, "err", cause.Error())
	return l.Drop()
}

// Status is what the tray and the status page show.
type Status struct {
	Logging    bool      `json:"logging"`
	ReportKey  string    `json:"report_key,omitempty"`
	ReportID   string    `json:"report_id,omitempty"`
	LogPath    string    `json:"log_path,omitempty"`
	Fights     int       `json:"fights"`
	Offset     int64     `json:"offset"`
	LastAppend time.Time `json:"last_append,omitzero"`
	Queued     int       `json:"queued"`
}

// Status reports what the pipeline is doing right now.
func (p *Pipeline) Status() Status {
	s := Status{}
	if n, err := p.o.Queue.Len(); err == nil {
		s.Queued = n
	}
	if p.cur == nil {
		return s
	}
	s.Logging = true
	s.ReportKey, s.ReportID = p.cur.Key, p.cur.ReportID
	s.LogPath, s.Fights = p.cur.LogPath, p.cur.Fights
	s.Offset, s.LastAppend = p.cur.Offset, p.cur.LastAppend
	return s
}

// Run drives the pipeline until the context ends: parse on every tick,
// upload what is queued, and push a live snapshot when one is due.
// Errors are logged rather than returned, because a companion that
// exits on a failed upload is a companion that loses a raid night.
func (p *Pipeline) Run(ctx context.Context, every time.Duration) error {
	t := time.NewTicker(every)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case now := <-t.C:
			if err := p.Tick(now); err != nil {
				p.o.Log.Error("the tail failed", "component", "pipeline", "err", err.Error())
			}
			if err := p.Live(ctx, now); err != nil {
				p.o.Log.Error("a live snapshot failed", "component", "pipeline", "err", err.Error())
			}
			if err := p.Drain(ctx); err != nil && !errors.Is(err, context.Canceled) {
				p.o.Log.Warn("uploads are behind", "component", "uploader", "err", err.Error())
			}
		}
	}
}
```

- [ ] **Step 7: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/pipeline/ && go test ./internal/pipeline/ ./internal/fixture/ -race -v`
Expected: PASS, ten pipeline tests. `TestARestartResumesTheSameReport` and `TestAnOutageQueuesAndTheOrderSurvivesTheReconnect` are the two that matter; the rest hold the edges still.

- [ ] **Step 8: Commit**

```bash
git add companion/internal/fixture/ \
  companion/internal/pipeline/
git commit -m "feat(companion): the live pipeline, from appended bytes to uploaded fights" \
  -m "Tick reads the log, feeds the engine and queues what falls out;
Drain sends the queue oldest first and stops on the first retryable
failure; Live goes straight out and is dropped when it fails. Keeping
them apart is what lets a raid be parsed through an outage and
uploaded in order afterwards.

Every offset sent is report-relative -- the raw chunks concatenate
into exactly the bytes the engine parsed -- and state.StartOffset is
the only place that meets the watcher's file offsets. A restart
re-reads the raw bytes that had not reached a 4 MiB boundary, without
which the server could never reassemble the stream.

The fixture is hand-written in the v16 shapes the engine verified. It
carries a heartbeat line after each encounter because the engine
emits a fight when the line after its last one arrives." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 10: Addon sync — exports out, inbox in

**Files:**
- Create: `companion/internal/addon/lua.go`, `companion/internal/addon/addon.go`, `companion/internal/addon/sync.go`, `companion/internal/client/addon.go`
- Test: `companion/internal/addon/addon_test.go`

**Interfaces:**
- Consumes: `(*Client).do`, `request` (Task 2).
- Produces:
  - `addon.Table{Fields map[string]any; Items []any}` with `.Get(key string) string`, and `addon.ParseLua(src string) (map[string]any, error)`.
  - `addon.SavedVariablesName` (`ForeverSixty`), `addon.InboxName`, `addon.InboxGlobal`.
  - `addon.Export{Name, Ruleset, Region, Export string}` with json `name`, `ruleset`, `region`, `export`; `addon.Build{ID, Name, Character, Code string}`; `addon.Inbox{Builds []Build}`.
  - `addon.ScanExports(path string) ([]Export, error)`, `addon.RenderInbox(Inbox, at time.Time) []byte`, `addon.InboxPath(savedVariables string) string`, `addon.WriteInbox(path string, body []byte) (bool, error)`.
  - `addon.InboxEvery` (10 minutes), `addon.API` (the two methods the sync needs), `addon.Options{Paths func() []string; API API; Log *slog.Logger; Every time.Duration}`, `addon.New(Options) *Sync`, `(*Sync).Poll(ctx, now) error`.
  - `(*Client).PostAddonExports(ctx, []addon.Export) error`, `(*Client).AddonInbox(ctx) (addon.Inbox, error)` — `client` imports `addon` for these two wire types, and `addon` never imports `client`.

The Lua reader covers exactly the subset saved variables use: top-level assignments of strings, numbers, booleans, nil and nested tables, with `--` and `--[[ ]]` comments. It exists because the companion must read what the addon saved without embedding an interpreter. It never looks inside an export string: FS1 is the addon's format and the site's.

`collect` treats *any* table with a string `export` field as a character, wherever it sits, so a change to the addon's table shape does not silently stop the sync. It reads `ruleset` if the addon wrote one and `realm` if it wrote the Phase 2 key, and posts whichever it found in the contract's `ruleset` field, untouched — the API maps it once the beta log settles what Forever's realm segment carries.

The inbox render is deterministic, which is what lets `WriteInbox` skip a write that would change nothing. `RenderInbox` output parses back through `ParseLua`, and the test asserts exactly that round trip.

- [ ] **Step 1: Write the failing test**

```go
package addon

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// savedVariables is the shape the Phase 2 addon writes.
const savedVariables = `
ForeverSixtyDB = {
	["profileKeys"] = {
		["Morrowlyn - Sanguine"] = "Default",
	},
	["characters"] = {
		["Morrowlyn-Hardcore"] = {
			["name"] = "Morrowlyn",
			["realm"] = "Hardcore",
			["region"] = "us",
			["level"] = 60,
			["export"] = "FS1:1.15.9.69722:paladin:human:503200000/0/0:head=12640,chest=11726",
		},
		["Thalgrit-Normal"] = {
			["name"] = "Thalgrit",
			["ruleset"] = "normal",
			["region"] = "eu",
			["export"] = "FS1:1.15.9.69722:warrior:orc:0/310000000/0:head=12640",
		},
	},
	["dataBuild"] = "1.15.9.69722",
}
`

func TestScanExportsReadsEveryCharacter(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ForeverSixty.lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := ScanExports(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("exports = %+v", got)
	}
	if got[0].Name != "Morrowlyn" || got[0].Ruleset != "Hardcore" || got[0].Region != "us" ||
		!strings.HasPrefix(got[0].Export, "FS1:") {
		t.Errorf("first = %+v", got[0])
	}
	if got[1].Name != "Thalgrit" || got[1].Ruleset != "normal" {
		t.Errorf("second = %+v", got[1])
	}
}

func TestTheLuaReaderHandlesTheShapesSavedVariablesUses(t *testing.T) {
	src := `
-- a comment
--[[ a block
     comment ]]
Simple = "a \"quoted\" value\nwith a newline"
Numbers = { 1, -2.5, 1e3 }
Flags = { ["on"] = true, ["off"] = false, ["none"] = nil }
Named = { key = "value", ["bracket"] = "other" }
`
	top, err := ParseLua(src)
	if err != nil {
		t.Fatal(err)
	}
	if top["Simple"] != "a \"quoted\" value\nwith a newline" {
		t.Errorf("Simple = %q", top["Simple"])
	}
	nums, ok := top["Numbers"].(*Table)
	if !ok || len(nums.Items) != 3 || nums.Items[1] != -2.5 || nums.Items[2] != 1000.0 {
		t.Errorf("Numbers = %+v", top["Numbers"])
	}
	flags := top["Flags"].(*Table)
	if flags.Fields["on"] != true || flags.Fields["off"] != false {
		t.Errorf("Flags = %+v", flags.Fields)
	}
	named := top["Named"].(*Table)
	if named.Get("key") != "value" || named.Get("bracket") != "other" {
		t.Errorf("Named = %+v", named.Fields)
	}
}

func TestAMalformedFileIsAnErrorWithALineNumber(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ForeverSixty.lua")
	if err := os.WriteFile(path, []byte("A = {\nB = \"unterminated\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := ScanExports(path)
	if err == nil || !strings.Contains(err.Error(), "lua line") {
		t.Fatalf("err = %v", err)
	}
	if _, err := ScanExports(filepath.Join(t.TempDir(), "missing.lua")); err == nil {
		t.Fatal("a missing file was not an error")
	}
}

func TestTheInboxRendersLuaThatReadsBack(t *testing.T) {
	body := RenderInbox(Inbox{Builds: []Build{
		{ID: "b1", Name: `Holy "Bubble"`, Character: "Morrowlyn", Code: "FSB1:1.15.9.69722:paladin:001:head=12640"},
	}}, t0)
	top, err := ParseLua(string(body))
	if err != nil {
		t.Fatalf("the rendered inbox does not parse: %v\n%s", err, body)
	}
	root, ok := top[InboxGlobal].(*Table)
	if !ok {
		t.Fatalf("the inbox does not assign %s: %s", InboxGlobal, body)
	}
	if root.Get("generated_at") != "2026-12-09T20:00:00Z" {
		t.Errorf("generated_at = %q", root.Get("generated_at"))
	}
	builds := root.Fields["builds"].(*Table)
	if len(builds.Items) != 1 {
		t.Fatalf("builds = %+v", builds)
	}
	first := builds.Items[0].(*Table)
	if first.Get("name") != `Holy "Bubble"` || first.Get("code") != "FSB1:1.15.9.69722:paladin:001:head=12640" {
		t.Errorf("build = %+v", first.Fields)
	}
	if string(RenderInbox(Inbox{Builds: []Build{{ID: "b1", Name: `Holy "Bubble"`,
		Character: "Morrowlyn", Code: "FSB1:1.15.9.69722:paladin:001:head=12640"}}}, t0)) != string(body) {
		t.Error("two renders of the same inbox differ")
	}
}

func TestWriteInboxSkipsAnIdenticalFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "SavedVariables", InboxName+".lua")
	body := RenderInbox(Inbox{}, t0)
	wrote, err := WriteInbox(path, body)
	if err != nil || !wrote {
		t.Fatalf("first write = %v, %v", wrote, err)
	}
	wrote, err = WriteInbox(path, body)
	if err != nil || wrote {
		t.Fatalf("second write = %v, %v", wrote, err)
	}
	wrote, err = WriteInbox(path, RenderInbox(Inbox{}, t0.Add(time.Hour)))
	if err != nil || !wrote {
		t.Fatalf("changed write = %v, %v", wrote, err)
	}
}

func TestInboxPathSitsBesideTheSavedVariables(t *testing.T) {
	got := InboxPath("/w/WTF/Account/A#1/SavedVariables/ForeverSixty.lua")
	want := filepath.Join("/w/WTF/Account/A#1/SavedVariables", InboxName+".lua")
	if got != want {
		t.Fatalf("InboxPath = %q, want %q", got, want)
	}
}

// fakeAPI records what the sync sent and answers with a fixed inbox.
type fakeAPI struct {
	posted []Export
	inbox  Inbox
	err    error
}

func (f *fakeAPI) PostAddonExports(_ context.Context, chars []Export) error {
	if f.err != nil {
		return f.err
	}
	f.posted = append(f.posted, chars...)
	return nil
}

func (f *fakeAPI) AddonInbox(context.Context) (Inbox, error) { return f.inbox, nil }

func TestTheSyncUploadsOnAModificationAndWritesTheInboxOnItsTimer(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, t0, t0); err != nil {
		t.Fatal(err)
	}
	api := &fakeAPI{inbox: Inbox{Builds: []Build{{ID: "b1", Name: "Holy", Code: "FSB1:x"}}}}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api})

	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("posted %+v", api.posted)
	}
	if _, err := os.Stat(InboxPath(path)); err != nil {
		t.Fatalf("the inbox was not written: %v", err)
	}

	// Nothing changed: no second upload, and no inbox refresh until
	// the timer comes round.
	if err := s.Poll(t.Context(), t0.Add(time.Minute)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("a quiet pass uploaded again: %+v", api.posted)
	}

	// The player logs out: the file's modification time moves.
	if err := os.Chtimes(path, t0.Add(time.Hour), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if err := s.Poll(t.Context(), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 4 {
		t.Fatalf("a logout did not upload: %+v", api.posted)
	}
}

func TestAFailedUploadIsRetriedOnTheNextPass(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	api := &fakeAPI{err: errors.New("offline")}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api})
	if err := s.Poll(t.Context(), t0); err == nil {
		t.Fatal("a failed upload was not reported")
	}
	api.err = nil
	if err := s.Poll(t.Context(), t0.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("the retry did not happen: %+v", api.posted)
	}
}

func TestAMissingSavedVariablesFileIsNotAnError(t *testing.T) {
	api := &fakeAPI{}
	s := New(Options{
		Paths: func() []string { return []string{filepath.Join(t.TempDir(), "nope.lua")} },
		API:   api,
	})
	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 0 {
		t.Errorf("posted %+v", api.posted)
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/addon/`
Expected: FAIL, "undefined: ScanExports", "undefined: ParseLua".

- [ ] **Step 3: Write `internal/addon/lua.go`**

```go
// companion/internal/addon/lua.go
// A reader for the subset of Lua that WoW writes into SavedVariables:
// top-level assignments of strings, numbers, booleans, nil and nested
// tables. It exists because the companion must read what the addon
// saved without embedding a Lua interpreter, and because the export
// strings themselves are opaque to it — FS1 is the addon's format,
// parsed on the site, never here.
package addon

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Table is a decoded Lua table: keyed entries in Fields, positional
// entries in Items, in the order they were written.
type Table struct {
	Fields map[string]any
	Items  []any
}

// Get returns a field as a string, empty when it is absent or is not
// a string.
func (t *Table) Get(key string) string {
	if t == nil {
		return ""
	}
	if v, ok := t.Fields[key].(string); ok {
		return v
	}
	return ""
}

// parser is a recursive-descent reader over the source.
type parser struct {
	src string
	i   int
}

// ParseLua reads a SavedVariables file into its top-level assignments.
func ParseLua(src string) (map[string]any, error) {
	p := &parser{src: src}
	out := map[string]any{}
	for {
		p.space()
		if p.i >= len(p.src) {
			return out, nil
		}
		name, err := p.name()
		if err != nil {
			return nil, err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return nil, err
		}
		v, err := p.value()
		if err != nil {
			return nil, err
		}
		out[name] = v
	}
}

func (p *parser) errf(format string, a ...any) error {
	line := 1 + strings.Count(p.src[:min(p.i, len(p.src))], "\n")
	return fmt.Errorf("lua line %d: %s", line, fmt.Sprintf(format, a...))
}

// space skips whitespace and -- comments, including --[[ blocks.
func (p *parser) space() {
	for p.i < len(p.src) {
		c := p.src[p.i]
		switch {
		case c == ' ' || c == '\t' || c == '\r' || c == '\n':
			p.i++
		case strings.HasPrefix(p.src[p.i:], "--[["):
			end := strings.Index(p.src[p.i:], "]]")
			if end < 0 {
				p.i = len(p.src)
				return
			}
			p.i += end + 2
		case strings.HasPrefix(p.src[p.i:], "--"):
			end := strings.IndexByte(p.src[p.i:], '\n')
			if end < 0 {
				p.i = len(p.src)
				return
			}
			p.i += end + 1
		default:
			return
		}
	}
}

func (p *parser) expect(c byte) error {
	if p.i >= len(p.src) || p.src[p.i] != c {
		return p.errf("expected %q", string(c))
	}
	p.i++
	return nil
}

func (p *parser) name() (string, error) {
	start := p.i
	for p.i < len(p.src) {
		c := rune(p.src[p.i])
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || c == '.' {
			p.i++
			continue
		}
		break
	}
	if p.i == start {
		return "", p.errf("expected a name")
	}
	return p.src[start:p.i], nil
}

func (p *parser) value() (any, error) {
	p.space()
	if p.i >= len(p.src) {
		return nil, p.errf("expected a value")
	}
	switch c := p.src[p.i]; {
	case c == '"' || c == '\'':
		return p.str()
	case c == '{':
		return p.table()
	case c == '-' || (c >= '0' && c <= '9'):
		return p.number()
	default:
		word, err := p.name()
		if err != nil {
			return nil, err
		}
		switch word {
		case "true":
			return true, nil
		case "false":
			return false, nil
		case "nil":
			return nil, nil
		}
		return nil, p.errf("unexpected word %q", word)
	}
}

func (p *parser) str() (string, error) {
	quote := p.src[p.i]
	p.i++
	var b strings.Builder
	for p.i < len(p.src) {
		c := p.src[p.i]
		switch {
		case c == quote:
			p.i++
			return b.String(), nil
		case c == '\\':
			p.i++
			if p.i >= len(p.src) {
				return "", p.errf("the file ends inside a string")
			}
			e := p.src[p.i]
			switch {
			case e == 'n':
				b.WriteByte('\n')
				p.i++
			case e == 't':
				b.WriteByte('\t')
				p.i++
			case e == 'r':
				b.WriteByte('\r')
				p.i++
			case e >= '0' && e <= '9':
				j := p.i
				for j < len(p.src) && j-p.i < 3 && p.src[j] >= '0' && p.src[j] <= '9' {
					j++
				}
				n, err := strconv.Atoi(p.src[p.i:j])
				if err != nil || n > 255 {
					return "", p.errf("bad escape")
				}
				b.WriteByte(byte(n))
				p.i = j
			default:
				b.WriteByte(e)
				p.i++
			}
		default:
			b.WriteByte(c)
			p.i++
		}
	}
	return "", p.errf("the file ends inside a string")
}

func (p *parser) number() (float64, error) {
	start := p.i
	if p.src[p.i] == '-' {
		p.i++
	}
	for p.i < len(p.src) {
		c := p.src[p.i]
		if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' ||
			((c == '+' || c == '-') && (p.src[p.i-1] == 'e' || p.src[p.i-1] == 'E')) {
			p.i++
			continue
		}
		break
	}
	n, err := strconv.ParseFloat(p.src[start:p.i], 64)
	if err != nil {
		return 0, p.errf("bad number %q", p.src[start:p.i])
	}
	return n, nil
}

func (p *parser) table() (*Table, error) {
	if err := p.expect('{'); err != nil {
		return nil, err
	}
	t := &Table{Fields: map[string]any{}}
	for {
		p.space()
		if p.i >= len(p.src) {
			return nil, p.errf("the file ends inside a table")
		}
		if p.src[p.i] == '}' {
			p.i++
			return t, nil
		}
		if err := p.field(t); err != nil {
			return nil, err
		}
		p.space()
		if p.i < len(p.src) && (p.src[p.i] == ',' || p.src[p.i] == ';') {
			p.i++
		}
	}
}

func (p *parser) field(t *Table) error {
	p.space()
	switch {
	case p.src[p.i] == '[':
		p.i++
		k, err := p.value()
		if err != nil {
			return err
		}
		p.space()
		if err := p.expect(']'); err != nil {
			return err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return err
		}
		v, err := p.value()
		if err != nil {
			return err
		}
		t.Fields[keyOf(k)] = v
		return nil
	case isNameStart(p.src[p.i]) && p.assignmentAhead():
		name, err := p.name()
		if err != nil {
			return err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return err
		}
		v, err := p.value()
		if err != nil {
			return err
		}
		t.Fields[name] = v
		return nil
	default:
		v, err := p.value()
		if err != nil {
			return err
		}
		t.Items = append(t.Items, v)
		return nil
	}
}

// assignmentAhead distinguishes `name = value` from a bare word value
// such as true, false or nil.
func (p *parser) assignmentAhead() bool {
	j := p.i
	for j < len(p.src) && (isNameStart(p.src[j]) || (p.src[j] >= '0' && p.src[j] <= '9')) {
		j++
	}
	for j < len(p.src) && (p.src[j] == ' ' || p.src[j] == '\t') {
		j++
	}
	return j < len(p.src) && p.src[j] == '='
}

func isNameStart(c byte) bool {
	return c == '_' || unicode.IsLetter(rune(c))
}

func keyOf(v any) string {
	switch k := v.(type) {
	case string:
		return k
	case float64:
		return strconv.FormatFloat(k, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(k)
	}
	return ""
}

// quote writes a Lua string literal.
func quote(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for i := range len(s) {
		switch c := s[i]; c {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			if c < 0x20 {
				fmt.Fprintf(&b, `\%03d`, c)
				continue
			}
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
	return b.String()
}
```

- [ ] **Step 4: Write `internal/addon/addon.go`**

```go
// companion/internal/addon/addon.go
// Package addon carries strings between the game and the site: the
// addon's character exports out of SavedVariables and up to the API,
// and the builds the player chose on the site down into an inbox file
// the addon reads at load.
//
// The export strings are opaque here. FS1 is the addon's format and
// the site's; the companion moves it and never inspects it.
package addon

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// SavedVariablesName is the addon's saved-variables file.
const SavedVariablesName = "ForeverSixty"

// InboxName is the file the companion writes beside it.
const InboxName = "ForeverSixtyInbox"

// InboxGlobal is the global the inbox file assigns.
const InboxGlobal = "ForeverSixtyInbox"

// Export is one character's export string, ready to post. Ruleset is
// the second segment of the character key; whatever the addon wrote
// there is passed through untouched, because the beta log has not yet
// settled what Forever puts in it.
type Export struct {
	Name    string `json:"name"`
	Ruleset string `json:"ruleset"`
	Region  string `json:"region"`
	Export  string `json:"export"`
}

// Build is one build the site sent down for the addon to load.
type Build struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character,omitempty"`
	Code      string `json:"code"`
}

// Inbox is the body of GET /v1/addon/inbox.
type Inbox struct {
	Builds []Build `json:"builds"`
}

// ScanExports reads one SavedVariables file and returns every
// character export in it. Any table holding a string "export" field
// counts, wherever it sits, so a change to the addon's table shape
// does not silently stop the sync.
func ScanExports(path string) ([]Export, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	top, err := ParseLua(string(b))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	var out []Export
	for _, v := range top {
		collect(v, &out)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Name != out[j].Name {
			return out[i].Name < out[j].Name
		}
		return out[i].Ruleset < out[j].Ruleset
	})
	return out, nil
}

// collect walks a decoded value looking for export tables.
func collect(v any, out *[]Export) {
	t, ok := v.(*Table)
	if !ok {
		return
	}
	if s := t.Get("export"); s != "" {
		// The Phase 2 addon may still write the segment as "realm";
		// either key feeds the contract's ruleset field.
		ruleset := t.Get("ruleset")
		if ruleset == "" {
			ruleset = t.Get("realm")
		}
		*out = append(*out, Export{
			Name:    t.Get("name"),
			Ruleset: ruleset,
			Region:  t.Get("region"),
			Export:  s,
		})
	}
	for _, child := range t.Fields {
		collect(child, out)
	}
	for _, child := range t.Items {
		collect(child, out)
	}
}

// RenderInbox writes the Lua the addon loads. It is deterministic:
// the same inbox at the same time renders byte for byte the same, so
// the companion can skip a write that would change nothing.
func RenderInbox(in Inbox, at time.Time) []byte {
	var b strings.Builder
	b.WriteString("-- Written by the Forever Sixty companion. Do not edit:\n")
	b.WriteString("-- this file is replaced every ten minutes.\n")
	fmt.Fprintf(&b, "%s = {\n", InboxGlobal)
	fmt.Fprintf(&b, "\t[\"generated_at\"] = %s,\n", quote(at.UTC().Format(time.RFC3339)))
	b.WriteString("\t[\"builds\"] = {\n")
	for _, bd := range in.Builds {
		b.WriteString("\t\t{\n")
		fmt.Fprintf(&b, "\t\t\t[\"id\"] = %s,\n", quote(bd.ID))
		fmt.Fprintf(&b, "\t\t\t[\"name\"] = %s,\n", quote(bd.Name))
		fmt.Fprintf(&b, "\t\t\t[\"character\"] = %s,\n", quote(bd.Character))
		fmt.Fprintf(&b, "\t\t\t[\"code\"] = %s,\n", quote(bd.Code))
		b.WriteString("\t\t},\n")
	}
	b.WriteString("\t},\n}\n")
	return []byte(b.String())
}

// InboxPath is where the inbox goes for one SavedVariables path.
func InboxPath(savedVariables string) string {
	return filepath.Join(filepath.Dir(savedVariables), InboxName+".lua")
}

// WriteInbox writes the inbox file unless the bytes already there are
// identical. It reports whether it wrote.
//
// The game rewrites its own SavedVariables at logout, so an inbox
// written while the player is online may be replaced; the next
// ten-minute pass puts it back, and the addon reads it at the
// following load.
func WriteInbox(path string, body []byte) (bool, error) {
	if old, err := os.ReadFile(path); err == nil && string(old) == string(body) {
		return false, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".inbox-*.lua")
	if err != nil {
		return false, err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err := tmp.Write(body); err != nil {
		tmp.Close()
		return false, err
	}
	if err := tmp.Close(); err != nil {
		return false, err
	}
	if err := os.Chmod(name, 0o644); err != nil {
		return false, err
	}
	return true, os.Rename(name, path)
}
```

- [ ] **Step 5: Write `internal/addon/sync.go`**

```go
// companion/internal/addon/sync.go
// The addon sync loop: watch each SavedVariables file's modification
// time, upload its exports when the player logs out, and write the
// inbox back on a timer.
package addon

import (
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"time"
)

// InboxEvery is how often the inbox is refreshed, per the contract.
const InboxEvery = 10 * time.Minute

// API is the part of the client the sync uses.
type API interface {
	PostAddonExports(ctx context.Context, chars []Export) error
	AddonInbox(ctx context.Context) (Inbox, error)
}

// Options configures a Sync.
type Options struct {
	// Paths returns the SavedVariables files to watch, one per
	// account folder per install. It is a function because the player
	// can add an install while the companion runs.
	Paths func() []string
	API   API
	Log   *slog.Logger
	Every time.Duration
}

// Sync moves strings between the addon and the API.
type Sync struct {
	o         Options
	seen      map[string]time.Time
	lastInbox time.Time
}

// New builds a sync.
func New(o Options) *Sync {
	if o.Every == 0 {
		o.Every = InboxEvery
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	return &Sync{o: o, seen: map[string]time.Time{}}
}

// Poll does one pass. It is called on the same ticker as the pipeline
// and does nothing expensive when nothing has changed.
func (s *Sync) Poll(ctx context.Context, now time.Time) error {
	var errs []error
	paths := s.o.Paths()
	for _, p := range paths {
		fi, err := os.Stat(p)
		if errors.Is(err, fs.ErrNotExist) {
			continue // the player has not installed the addon here
		}
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if was, ok := s.seen[p]; ok && was.Equal(fi.ModTime()) {
			continue
		}
		s.seen[p] = fi.ModTime()
		exports, err := ScanExports(p)
		if err != nil {
			s.o.Log.Warn("could not read the addon's saved variables",
				"component", "addon", "path", p, "err", err.Error())
			continue
		}
		if len(exports) == 0 {
			continue
		}
		if err := s.o.API.PostAddonExports(ctx, exports); err != nil {
			delete(s.seen, p) // try again next pass
			errs = append(errs, err)
			continue
		}
		s.o.Log.Info("uploaded character exports", "component", "addon",
			"path", p, "characters", len(exports))
	}

	if !s.lastInbox.IsZero() && now.Sub(s.lastInbox) < s.o.Every {
		return errors.Join(errs...)
	}
	inbox, err := s.o.API.AddonInbox(ctx)
	if err != nil {
		return errors.Join(append(errs, err)...)
	}
	s.lastInbox = now
	body := RenderInbox(inbox, now)
	for _, p := range paths {
		wrote, err := WriteInbox(InboxPath(p), body)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if wrote {
			s.o.Log.Info("wrote the addon inbox", "component", "addon",
				"path", InboxPath(p), "builds", len(inbox.Builds))
		}
	}
	return errors.Join(errs...)
}
```

- [ ] **Step 6: Write `internal/client/addon.go`**

```go
// companion/internal/client/addon.go
// The two addon routes. Both carry strings the companion does not
// read: FS1 exports going up and addon codes coming down.
package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
)

// PostAddonExports uploads the character exports the addon saved.
func (c *Client) PostAddonExports(ctx context.Context, chars []addon.Export) error {
	if len(chars) == 0 {
		return nil
	}
	body, err := json.Marshal(struct {
		Characters []addon.Export `json:"characters"`
	}{chars})
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/addon/exports",
		Body: body, Type: "application/json",
	})
	return err
}

// AddonInbox fetches the builds the player chose on the site.
func (c *Client) AddonInbox(ctx context.Context) (addon.Inbox, error) {
	var out addon.Inbox
	if _, err := c.do(ctx, request{
		Method: http.MethodGet, Path: "/v1/addon/inbox", Out: &out,
	}); err != nil {
		return addon.Inbox{}, err
	}
	return out, nil
}
```

- [ ] **Step 7: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/addon/ ./internal/client/ && go test ./internal/addon/ ./internal/client/ -race`
Expected: PASS.

- [ ] **Step 8: Commit**

```bash
git add companion/internal/addon/ \
  companion/internal/client/addon.go
git commit -m "feat(companion): addon sync, in both directions" \
  -m "After a logout the companion notices the saved-variables file's
modification time move, reads every character export out of it and
posts them; every ten minutes it writes the builds chosen on the site
into ForeverSixtyInbox.lua beside them.

The export strings are opaque: FS1 belongs to the addon and the site.
So does the ruleset segment -- whatever the addon wrote under
'ruleset', or under Phase 2's 'realm', travels through untouched and
the API maps it once the beta log settles what Forever's realm
segment carries.

Reading them needs a Lua reader rather than an interpreter: the
subset saved variables use is assignments, tables, strings, numbers,
booleans and comments. The inbox render is deterministic, so a pass
that changes nothing writes nothing." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Task 11: Auto-update, with minisign verification

**Files:**
- Create: `companion/internal/updater/updater.go`
- Test: `companion/internal/updater/updater_test.go`

**Interfaces:**
- Consumes: nothing.
- Produces:
  - `updater.Repo`, `updater.TagPrefix` (`companion-v`), `updater.Every` (6 hours), `updater.MaxAsset` (128 MiB).
  - `updater.PublicKey` and `updater.Version` — package variables set at build time with `-ldflags -X`.
  - `updater.Asset`, `updater.Release{TagName string; Assets []Asset}`.
  - `updater.Options{Dir, APIBase string; HTTP *http.Client; Log *slog.Logger; Version, PublicKey, GOOS, GOARCH string}`, `updater.New(Options) *Updater`.
  - `updater.AssetName(goos, goarch string) string`, `updater.PendingPath(dir string) string`, `updater.Newer(tag, current string) bool`, `updater.ErrNoKey`.
  - `(*Updater).Latest(ctx) (Release, error)`, `.Stage(ctx, Release) (bool, error)`, `.Poll(ctx, now) (bool, error)`.
  - `updater.Verify(publicKey string, body, signature []byte) error`, `updater.ApplyPending(dir, publicKey, exe string) (bool, error)`.

The rule this package exists to enforce: **nothing runs that cannot be proved to have come from us.** The asset is verified against a key compiled into the binary before it is staged, and verified *again* before it is swapped in, because it sat on disk in between and that is the moment it becomes the program. A build with no embedded key does not update at all rather than accepting something unverified.

`ApplyPending` runs at startup, before anything else, because Windows will not let a running binary be replaced — which is what the contract's "swaps on next launch" means.

- [ ] **Step 1: Write the failing test**

```go
package updater

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	minisign "github.com/jedisct1/go-minisign"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// keypair mints a minisign key for the test. The workflow's signing
// key is a real minisign secret key; this is the same algorithm with
// no passphrase.
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
	raw := append(append(append([]byte{}, pk.SignatureAlgorithm[:]...), pk.KeyId[:]...),
		pk.PublicKey[:]...)
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

// release serves a GitHub-shaped latest release with one asset.
func release(t *testing.T, tag string, binary, signature []byte) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	name := AssetName("linux", "amd64")
	mux.HandleFunc("/repos/"+Repo+"/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"tag_name":%q,"assets":[{"name":%q,"browser_download_url":"%s/a"},`+
			`{"name":"%s.minisig","browser_download_url":"%s/a.minisig"}]}`,
			tag, name, srv.URL, name, srv.URL)
	})
	mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) { w.Write(binary) })
	mux.HandleFunc("/a.minisig", func(w http.ResponseWriter, r *http.Request) { w.Write(signature) })
	return srv
}

func TestNewerComparesDottedVersions(t *testing.T) {
	for _, tc := range []struct {
		tag, current string
		want         bool
	}{
		{"companion-v1.0.1", "1.0.0", true},
		{"companion-v1.10.0", "1.9.3", true},
		{"companion-v1.0.0", "1.0.0", false},
		{"companion-v0.9.9", "1.0.0", false},
		{"companion-v2.0.0", "1.99.99", true},
	} {
		if got := Newer(tc.tag, tc.current); got != tc.want {
			t.Errorf("Newer(%q, %q) = %v", tc.tag, tc.current, got)
		}
	}
}

func TestAssetNamesCarryThePlatform(t *testing.T) {
	if got := AssetName("darwin", "arm64"); got != "foreversixty-companion_darwin_arm64" {
		t.Errorf("darwin = %q", got)
	}
	if got := AssetName("windows", "amd64"); got != "foreversixty-companion_windows_amd64.exe" {
		t.Errorf("windows = %q", got)
	}
}

func TestAVerifiedReleaseIsStaged(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("#!/bin/sh\necho the new companion\n")
	srv := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	dir := t.TempDir()
	u := New(Options{Dir: dir, APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})

	staged, err := u.Poll(t.Context(), t0)
	if err != nil || !staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	got, err := os.ReadFile(PendingPath(dir))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(binary) {
		t.Fatalf("staged %q", got)
	}
	// The six-hour rule: a second poll straight away does nothing.
	if staged, err := u.Poll(t.Context(), t0.Add(time.Hour)); err != nil || staged {
		t.Fatalf("an early second poll = %v, %v", staged, err)
	}
}

func TestATamperedDownloadIsRefused(t *testing.T) {
	sk, pub := keypair(t)
	good := []byte("the real companion")
	srv := release(t, "companion-v9.9.9", []byte("a trojan"), sign(t, sk, good))
	dir := t.TempDir()
	u := New(Options{Dir: dir, APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil {
		t.Fatal("a tampered download was accepted")
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Fatal("the tampered binary was staged anyway")
	}
}

func TestAReleaseThatIsNotNewerIsIgnored(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("same version")
	srv := release(t, "companion-v1.0.0", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if staged, err := u.Poll(t.Context(), t0); err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
}

func TestABuildWithNoKeyNeverUpdates(t *testing.T) {
	u := New(Options{Dir: t.TempDir(), APIBase: "http://127.0.0.1:1", Version: "1.0.0"})
	staged, err := u.Poll(t.Context(), t0)
	if err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	if _, err := u.Stage(t.Context(), Release{TagName: "companion-v2"}); !errors.Is(err, ErrNoKey) {
		t.Fatalf("Stage = %v, want ErrNoKey", err)
	}
}

func TestAReleaseMissingThisPlatformIsAnError(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("linux only")
	srv := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "windows", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil {
		t.Fatal("a release with no Windows asset was accepted")
	}
}

func TestApplyPendingSwapsTheBinaryAndVerifiesAgain(t *testing.T) {
	sk, pub := keypair(t)
	dir := t.TempDir()
	exe := filepath.Join(t.TempDir(), "foreversixty-companion")
	if err := os.WriteFile(exe, []byte("the old companion"), 0o755); err != nil {
		t.Fatal(err)
	}
	if swapped, err := ApplyPending(dir, pub, exe); err != nil || swapped {
		t.Fatalf("with nothing pending = %v, %v", swapped, err)
	}

	body := []byte("the new companion")
	if err := os.WriteFile(PendingPath(dir), body, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "pending.minisig"), sign(t, sk, body), 0o600); err != nil {
		t.Fatal(err)
	}
	swapped, err := ApplyPending(dir, pub, exe)
	if err != nil || !swapped {
		t.Fatalf("ApplyPending = %v, %v", swapped, err)
	}
	got, err := os.ReadFile(exe)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(body) {
		t.Fatalf("the binary is %q", got)
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Error("the pending file was left behind")
	}
}

func TestApplyPendingRefusesAndClearsAnUnsignedStage(t *testing.T) {
	_, pub := keypair(t)
	dir := t.TempDir()
	exe := filepath.Join(t.TempDir(), "foreversixty-companion")
	if err := os.WriteFile(exe, []byte("old"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(PendingPath(dir), []byte("unsigned"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := ApplyPending(dir, pub, exe); err == nil {
		t.Fatal("an unsigned stage was applied")
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Error("the unsigned stage was not cleared")
	}
	if got, _ := os.ReadFile(exe); string(got) != "old" {
		t.Errorf("the running binary changed to %q", got)
	}
}

func TestVerifyReportsWhichPieceIsWrong(t *testing.T) {
	sk, pub := keypair(t)
	body := []byte("payload")
	if err := Verify(pub, body, sign(t, sk, body)); err != nil {
		t.Fatal(err)
	}
	if err := Verify("not base64 at all", body, sign(t, sk, body)); err == nil {
		t.Error("a broken public key was accepted")
	}
	if err := Verify(pub, body, []byte("not a signature")); err == nil {
		t.Error("a broken signature was accepted")
	}
	if err := Verify(pub, []byte("other"), sign(t, sk, body)); err == nil {
		t.Error("a signature over other bytes was accepted")
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/updater/`
Expected: FAIL to build, "undefined: New", "no required module provides package github.com/jedisct1/go-minisign".

- [ ] **Step 3: Add the minisign dependency**

```bash
cd companion && go get github.com/jedisct1/go-minisign@v0.0.0-20260527172527-a09352b57a22
```

- [ ] **Step 4: Write the implementation**

```go
// companion/internal/updater/updater.go
// Package updater keeps the companion current from GitHub Releases,
// and refuses to run anything it cannot prove came from us: every
// asset is verified against a minisign public key compiled into the
// binary before it is staged, and again before it is swapped in.
package updater

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	minisign "github.com/jedisct1/go-minisign"
)

// Repo is where the releases live, and TagPrefix is what a companion
// release tag starts with.
const (
	Repo      = "jhunthrop/foreversixty"
	TagPrefix = "companion-v"
)

// Every is how often the companion checks, per the contract: at start
// and every six hours.
const Every = 6 * time.Hour

// MaxAsset is the largest download accepted, a ceiling well above a
// Go binary and well below anything that could fill a disk.
const MaxAsset = 128 << 20

// PublicKey is the minisign public key, set at build time with
// -ldflags "-X …/updater.PublicKey=RWQ…". An empty key disables
// updates rather than accepting unverified ones.
var PublicKey string

// Version is the running build, set the same way. It is compared
// against the latest release tag.
var Version = "0.0.0"

// Asset is one file on a release.
type Asset struct {
	Name string `json:"name"`
	URL  string `json:"browser_download_url"`
}

// Release is the part of the GitHub payload the updater reads.
type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

// Options configures an Updater.
type Options struct {
	// Dir is where downloads are staged.
	Dir string
	// APIBase is https://api.github.com unless a test says otherwise.
	APIBase string
	HTTP    *http.Client
	Log     *slog.Logger
	// Version and PublicKey default to the package variables.
	Version   string
	PublicKey string
	// GOOS and GOARCH default to the running platform, so a test can
	// ask for an asset it has.
	GOOS, GOARCH string
}

// Updater checks and stages releases.
type Updater struct {
	o        Options
	lastPoll time.Time
}

// New builds an updater.
func New(o Options) *Updater {
	if o.APIBase == "" {
		o.APIBase = "https://api.github.com"
	}
	if o.HTTP == nil {
		o.HTTP = &http.Client{Timeout: 5 * time.Minute}
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	if o.Version == "" {
		o.Version = Version
	}
	if o.PublicKey == "" {
		o.PublicKey = PublicKey
	}
	if o.GOOS == "" {
		o.GOOS = runtime.GOOS
	}
	if o.GOARCH == "" {
		o.GOARCH = runtime.GOARCH
	}
	return &Updater{o: o}
}

// AssetName is what the release workflow calls the binary for one
// platform. The version is not in the name, so the updater can ask
// for it without knowing the release first.
func AssetName(goos, goarch string) string {
	name := fmt.Sprintf("foreversixty-companion_%s_%s", goos, goarch)
	if goos == "windows" {
		name += ".exe"
	}
	return name
}

// PendingPath and pendingSig are where a verified download waits for
// the next launch.
func PendingPath(dir string) string { return filepath.Join(dir, "pending") }
func pendingSig(dir string) string  { return filepath.Join(dir, "pending.minisig") }

// Newer reports whether tag is a later version than current. Both are
// compared as dotted integers, so companion-v1.10.0 beats 1.9.3.
func Newer(tag, current string) bool {
	return compare(strings.TrimPrefix(tag, TagPrefix), current) > 0
}

func compare(a, b string) int {
	as, bs := strings.Split(a, "."), strings.Split(b, ".")
	for i := range max(len(as), len(bs)) {
		x, y := part(as, i), part(bs, i)
		if x != y {
			if x < y {
				return -1
			}
			return 1
		}
	}
	return 0
}

func part(s []string, i int) int {
	if i >= len(s) {
		return 0
	}
	n, err := strconv.Atoi(strings.TrimLeft(strings.SplitN(s[i], "-", 2)[0], "v"))
	if err != nil {
		return 0
	}
	return n
}

// Latest reads the newest release.
func (u *Updater) Latest(ctx context.Context) (Release, error) {
	url := u.o.APIBase + "/repos/" + Repo + "/releases/latest"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Release{}, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := u.o.HTTP.Do(req)
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Release{}, fmt.Errorf("GitHub answered %d for the latest release", resp.StatusCode)
	}
	var rel Release
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&rel); err != nil {
		return Release{}, err
	}
	return rel, nil
}

// ErrNoKey says the build carries no public key, so updates are off.
var ErrNoKey = errors.New("updater: this build has no minisign public key, so it will not update itself")

// Stage downloads and verifies this platform's asset from rel and
// leaves it ready for the next launch. It reports whether it staged
// anything.
func (u *Updater) Stage(ctx context.Context, rel Release) (bool, error) {
	if u.o.PublicKey == "" {
		return false, ErrNoKey
	}
	want := AssetName(u.o.GOOS, u.o.GOARCH)
	var bin, sig string
	for _, a := range rel.Assets {
		switch a.Name {
		case want:
			bin = a.URL
		case want + ".minisig":
			sig = a.URL
		}
	}
	if bin == "" || sig == "" {
		return false, fmt.Errorf("release %s carries no %s with a signature", rel.TagName, want)
	}
	body, err := u.fetch(ctx, bin)
	if err != nil {
		return false, err
	}
	sigBody, err := u.fetch(ctx, sig)
	if err != nil {
		return false, err
	}
	if err := Verify(u.o.PublicKey, body, sigBody); err != nil {
		return false, fmt.Errorf("release %s: %w", rel.TagName, err)
	}
	if err := os.MkdirAll(u.o.Dir, 0o700); err != nil {
		return false, err
	}
	if err := os.WriteFile(PendingPath(u.o.Dir), body, 0o755); err != nil {
		return false, err
	}
	if err := os.WriteFile(pendingSig(u.o.Dir), sigBody, 0o600); err != nil {
		return false, err
	}
	u.o.Log.Info("staged an update", "component", "updater",
		"tag", rel.TagName, "bytes", len(body))
	return true, nil
}

func (u *Updater) fetch(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := u.o.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("downloading %s: %d", url, resp.StatusCode)
	}
	b, err := io.ReadAll(io.LimitReader(resp.Body, MaxAsset+1))
	if err != nil {
		return nil, err
	}
	if len(b) > MaxAsset {
		return nil, fmt.Errorf("downloading %s: over the %d byte ceiling", url, MaxAsset)
	}
	return b, nil
}

// Verify checks a minisign signature over body.
func Verify(publicKey string, body, signature []byte) error {
	pk, err := minisign.NewPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("the embedded public key is unreadable: %w", err)
	}
	sig, err := minisign.DecodeSignature(string(signature))
	if err != nil {
		return fmt.Errorf("the signature is unreadable: %w", err)
	}
	ok, err := pk.Verify(body, sig)
	if err != nil {
		return fmt.Errorf("the signature does not verify: %w", err)
	}
	if !ok {
		return errors.New("the signature does not verify")
	}
	return nil
}

// Poll checks at most once every Every and stages what it finds.
func (u *Updater) Poll(ctx context.Context, now time.Time) (bool, error) {
	if !u.lastPoll.IsZero() && now.Sub(u.lastPoll) < Every {
		return false, nil
	}
	u.lastPoll = now
	if u.o.PublicKey == "" {
		return false, nil // an unsigned build simply never updates
	}
	rel, err := u.Latest(ctx)
	if err != nil {
		return false, err
	}
	if !Newer(rel.TagName, u.o.Version) {
		return false, nil
	}
	return u.Stage(ctx, rel)
}

// ApplyPending swaps a verified download over the running executable
// and reports whether it did. It runs at startup, before anything
// else, because Windows will not let a running binary be replaced.
func ApplyPending(dir, publicKey, exe string) (bool, error) {
	body, err := os.ReadFile(PendingPath(dir))
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	sig, err := os.ReadFile(pendingSig(dir))
	if err != nil {
		clearPending(dir)
		return false, fmt.Errorf("the staged update has no signature: %w", err)
	}
	// Verified again here: the file has been sitting on disk since
	// the download, and this is the moment it becomes the program.
	if err := Verify(publicKey, body, sig); err != nil {
		clearPending(dir)
		return false, err
	}
	previous := filepath.Join(dir, "previous")
	os.Remove(previous)
	if err := os.Rename(exe, previous); err != nil {
		return false, fmt.Errorf("move the running binary aside: %w", err)
	}
	if err := os.WriteFile(exe, body, 0o755); err != nil {
		os.Rename(previous, exe) // put it back rather than leave nothing
		return false, fmt.Errorf("write the new binary: %w", err)
	}
	clearPending(dir)
	return true, nil
}

func clearPending(dir string) {
	os.Remove(PendingPath(dir))
	os.Remove(pendingSig(dir))
}
```

- [ ] **Step 5: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/updater/ && go test ./internal/updater/ -race -v`
Expected: PASS, nine tests. `TestATamperedDownloadIsRefused` and `TestApplyPendingRefusesAndClearsAnUnsignedStage` are the two that matter.

- [ ] **Step 6: Commit**

```bash
git add companion/internal/updater/ \
  companion/go.mod \
  companion/go.sum
git commit -m "feat(companion): update from GitHub Releases, with minisign verification" \
  -m "A check at start and every six hours, the platform's asset and its
.minisig downloaded together, and the signature checked against a
public key compiled into the binary before anything is staged -- and
checked again before the staged file becomes the program, because it
sat on disk in between.

A build with no embedded key never updates rather than accepting
something unverified. The swap happens at startup because Windows
will not let a running binary be replaced, which is what 'swaps on
next launch' means." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 12: The app — everything assembled behind a loopback UI

**Files:**
- Create: `companion/internal/app/app.go`, `companion/internal/app/http.go`, `companion/ui/assets.go`, `companion/ui/styles.css`, `companion/ui/index.html`, `companion/ui/app.js`
- Test: `companion/internal/app/app_test.go`

**Interfaces:**
- Consumes: everything from Tasks 1 through 11.
- Produces:
  - `app.TickEvery` (1 second), `app.Options{Dirs paths.Dirs; Config config.Config; Secret secret.Store; Log *slog.Logger; Now func() time.Time; Updater updater.Options; Retry client.Retry}`, `app.New(Options) (*App, error)`, `(*App).Close() error`.
  - `(*App).Pair(ctx, code string) error`, `.Unpair() error`, `.SaveSettings(Settings) error`, `.Snapshot() Snapshot`, `.Step(ctx, now)`, `.Run(ctx) error`, `.Handler() http.Handler`, `.Serve() (url string, ln net.Listener, err error)`.
  - `app.Settings{Visibility string; WoWPaths []string; LoggingCharacter *character.Character}`.
  - `app.Snapshot{EngineVersion string; Paired bool; TokenBackend, APIBaseURL, SiteBaseURL, Visibility string; LoggingCharacter *character.Character; Installs []InstallStatus; Pipeline pipeline.Status; Reports []ReportRow; Characters []addon.Export; UpdateStaged bool}`, `app.InstallStatus`, `app.ReportRow`.
  - `ui.Assets` (an `embed.FS` over `index.html`, `app.js`, `styles.css`).
  - `app.MaxRequestBody`.

The UI server is on loopback with a 24-byte random token *in the path*, because a webview cannot set a header and every relative fetch inherits a path prefix for free. A page on any other origin cannot guess it.

`Step` is one pass of everything on a timer, and it returns nothing: the companion must keep logging through every part of it failing. It parses and snapshots whether or not the device is paired, and only skips the uploads, the addon sync and the update check when it is not — so a player who pairs an hour late loses nothing.

The window lives in `companion/ui/` as three static files embedded into the binary. It cannot import the site's stylesheet, so the design tokens it needs are copied verbatim from `web/src/styles/tokens.css`; the header comment in `styles.css` says so, and keeping them in step is a manual job. Fonts are not self-hosted here: a desktop app should not make a third-party request on load, so the stacks fall through to the system faces.

- [ ] **Step 1: Write the failing test**

```go
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
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/app/`
Expected: FAIL, "undefined: app.New".

- [ ] **Step 3: Write `companion/ui/styles.css`**

```css
/* companion/ui/styles.css
   The companion cannot import the site's stylesheet, so the tokens it
   uses are copied from web/src/styles/tokens.css. Keep them in step
   with that file and design/DESIGN-SYSTEM.md; the values here are
   verbatim. Fonts are not self-hosted in the companion: the stacks
   fall through to the system faces rather than pulling a third-party
   request into a desktop app. */
:root {
  --bg: #07090d;
  --raised: #0d111a;
  --card-top: #131824;
  --line: #262e40;
  --line-soft: #1c2230;
  --line-warm: #3a3326;

  --text: #e9e4d8;
  --strong: #f2eee4;
  --muted: #9a9484;
  --nav: #b9b3a4;

  --gold: #e5b955;
  --gold-hover: #f5d27a;
  --gold-deep: #a8762a;

  --ok: #7bff5c;
  --warn: #d66e28;
  --bad: #ff6b5c;

  --font-display: 'Cinzel', 'Trajan Pro', Georgia, serif;
  --font-body: 'Barlow', 'Helvetica Neue', Arial, sans-serif;
  --font-mono: 'JetBrains Mono', 'SF Mono', Menlo, monospace;

  --radius-panel: 6px;
  --radius-control: 4px;
  --radius-pill: 3px;
}

* { box-sizing: border-box; }

body {
  margin: 0;
  background: var(--bg);
  color: var(--text);
  font: 400 14px/1.5 var(--font-body);
  -webkit-font-smoothing: antialiased;
}

header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 14px 18px;
  border-bottom: 1px solid var(--line);
  background: linear-gradient(180deg, #0b1018, var(--bg));
}

.wordmark {
  font: 700 16px/1 var(--font-display);
  letter-spacing: 0.06em;
  color: var(--gold);
}

nav { display: flex; gap: 4px; margin-left: auto; }

nav button {
  min-height: 36px;
  padding: 0 12px;
  border: 1px solid transparent;
  border-radius: var(--radius-control);
  background: none;
  color: var(--nav);
  font: 700 12px/1 var(--font-body);
  letter-spacing: 0.06em;
  text-transform: uppercase;
  cursor: pointer;
  transition: color 120ms, border-color 120ms;
}

nav button:hover { color: var(--strong); }
nav button[aria-current='page'] { color: var(--gold); border-color: var(--line-warm); }
:focus-visible { outline: 2px solid var(--gold); outline-offset: 2px; }

main { padding: 18px; max-width: 720px; }
section[hidden] { display: none; }

h2 {
  margin: 0 0 12px;
  font: 600 15px/1.2 var(--font-display);
  letter-spacing: 0.10em;
  text-transform: uppercase;
  color: var(--strong);
}

.panel {
  border: 1px solid var(--line);
  border-radius: var(--radius-panel);
  background: linear-gradient(180deg, var(--card-top), var(--raised));
  padding: 16px;
  margin-bottom: 16px;
}

.row {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  padding: 8px 0;
  border-bottom: 1px solid var(--line-soft);
}

.row:last-child { border-bottom: 0; }
.row .k { color: var(--muted); }
.row .v { font-family: var(--font-mono); font-variant-numeric: tabular-nums; color: var(--strong); }

.dot { display: inline-block; width: 8px; height: 8px; border-radius: 999px; margin-right: 6px; }
.dot.on { background: var(--ok); box-shadow: 0 0 8px var(--ok); }
.dot.off { background: var(--muted); }
.dot.bad { background: var(--bad); }

.note { color: var(--muted); font-size: 13px; margin: 8px 0 0; }
.note.warn { color: var(--warn); }
.note.bad { color: var(--bad); }

a { color: var(--gold); }
a:hover { color: var(--gold-hover); }

label { display: block; margin: 12px 0 4px; color: var(--muted); font-size: 13px; }

input, select {
  width: 100%;
  min-height: 36px;
  padding: 0 10px;
  border: 1px solid var(--line-warm);
  border-radius: var(--radius-control);
  background: #0a0e15;
  color: var(--text);
  font: 400 14px/1 var(--font-body);
}

button.action {
  min-height: 36px;
  margin-top: 14px;
  padding: 0 14px;
  border: 1px solid var(--line-warm);
  border-radius: var(--radius-control);
  background: none;
  color: var(--text);
  font: 700 12px/1 var(--font-body);
  letter-spacing: 0.06em;
  text-transform: uppercase;
  cursor: pointer;
}

button.action:hover { border-color: var(--gold-deep); color: var(--strong); }

ul.reports { list-style: none; margin: 0; padding: 0; }
ul.reports li { padding: 10px 0; border-bottom: 1px solid var(--line-soft); }
ul.reports li:last-child { border-bottom: 0; }
.when { font-family: var(--font-mono); color: var(--muted); font-size: 12px; }

@media (max-width: 520px) {
  header { flex-wrap: wrap; }
  nav { margin-left: 0; width: 100%; }
  nav button { flex: 1; }
  .row { flex-direction: column; gap: 2px; }
}
```

- [ ] **Step 4: Write `companion/ui/index.html`**

```html
<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Forever Sixty companion</title>
<link rel="stylesheet" href="styles.css">
</head>
<body>
<header>
  <span class="wordmark">Forever Sixty</span>
  <nav>
    <button data-page="status" aria-current="page">Status</button>
    <button data-page="reports">Reports</button>
    <button data-page="device">Device</button>
    <button data-page="settings">Settings</button>
  </nav>
</header>

<main>
  <section id="status">
    <h2>Logging</h2>
    <div class="panel">
      <div class="row"><span class="k"><span id="logging-dot" class="dot off"></span>State</span><span class="v" id="logging-state">—</span></div>
      <div class="row"><span class="k">Report</span><span class="v" id="logging-report">—</span></div>
      <div class="row"><span class="k">Fights sent</span><span class="v" id="logging-fights">0</span></div>
      <div class="row"><span class="k">Waiting to upload</span><span class="v" id="logging-queued">0</span></div>
      <div class="row"><span class="k">Engine</span><span class="v" id="engine-version">—</span></div>
      <p class="note" id="logging-note"></p>
    </div>
    <h2>Game folders</h2>
    <div class="panel" id="installs"></div>
  </section>

  <section id="reports" hidden>
    <h2>Your reports</h2>
    <div class="panel"><ul class="reports" id="report-list"></ul></div>
  </section>

  <section id="device" hidden>
    <h2>This device</h2>
    <div class="panel">
      <div class="row"><span class="k"><span id="paired-dot" class="dot off"></span>Pairing</span><span class="v" id="paired-state">—</span></div>
      <div class="row"><span class="k">Token stored in</span><span class="v" id="token-backend">—</span></div>
      <div class="row"><span class="k">API</span><span class="v" id="api-base">—</span></div>
      <label for="pair-code">Pairing code from the site</label>
      <input id="pair-code" autocomplete="off" spellcheck="false" placeholder="Sign in on the site, then paste the code">
      <button class="action" id="pair">Pair this device</button>
      <button class="action" id="unpair">Forget this device</button>
      <p class="note" id="device-note"></p>
    </div>
  </section>

  <section id="settings" hidden>
    <h2>Settings</h2>
    <div class="panel">
      <label for="visibility">New reports are</label>
      <select id="visibility">
        <option value="public">Public</option>
        <option value="unlisted">Unlisted</option>
        <option value="private">Private</option>
        <option value="guild">Guild only</option>
      </select>
      <label for="wow-paths">Game folders, one per line</label>
      <input id="wow-paths" autocomplete="off" spellcheck="false">
      <label for="logging-character">Logging character</label>
      <select id="logging-character"><option value="">Not set</option></select>
      <button class="action" id="save">Save</button>
      <p class="note" id="settings-note"></p>
    </div>
  </section>
</main>

<script src="app.js"></script>
</body>
</html>
```

- [ ] **Step 5: Write `companion/ui/app.js`**

```javascript
// companion/ui/app.js
// The whole window. It polls /api/status every two seconds and posts
// the three things the player can change. There is no framework and
// no build step: the page is four panels over one JSON document.
'use strict';

const api = (path, body) =>
  fetch(path, body === undefined
    ? { cache: 'no-store' }
    : { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(body) })
    .then((r) => r.json());

const $ = (id) => document.getElementById(id);
const text = (id, value) => { $(id).textContent = value; };

let latest = null;
let editing = false;

function show(page) {
  for (const s of document.querySelectorAll('main section')) s.hidden = s.id !== page;
  for (const b of document.querySelectorAll('nav button')) {
    if (b.dataset.page === page) b.setAttribute('aria-current', 'page');
    else b.removeAttribute('aria-current');
  }
}

for (const b of document.querySelectorAll('nav button')) {
  b.addEventListener('click', () => show(b.dataset.page));
}

function renderStatus(s) {
  const p = s.pipeline;
  $('logging-dot').className = 'dot ' + (p.logging ? 'on' : 'off');
  text('logging-state', p.logging ? 'Logging' : 'Waiting for the game');
  text('logging-report', p.report_id || (p.report_key ? 'not uploaded yet' : '—'));
  text('logging-fights', p.fights);
  text('logging-queued', p.queued);
  text('engine-version', s.engine_version);

  const note = $('logging-note');
  if (!s.paired) {
    note.className = 'note bad';
    note.textContent = 'This device is not paired, so nothing is uploading. Open the Device page.';
  } else if (!s.installs.length) {
    note.className = 'note bad';
    note.textContent = 'No World of Warcraft folder was found. Add one on the Settings page.';
  } else {
    note.className = 'note';
    note.textContent = p.log_path ? 'Reading ' + p.log_path : 'Nothing is being written to the log yet.';
  }

  const installs = $('installs');
  installs.replaceChildren();
  if (!s.installs.length) {
    const empty = document.createElement('p');
    empty.className = 'note';
    empty.textContent = 'None found.';
    installs.append(empty);
  }
  for (const i of s.installs) {
    const row = document.createElement('div');
    row.className = 'row';
    const k = document.createElement('span');
    k.className = 'k';
    k.textContent = i.path;
    const v = document.createElement('span');
    v.className = 'v';
    if (!i.advanced_logging_known) v.textContent = 'advanced logging: unknown';
    else v.textContent = i.advanced_logging ? 'advanced logging: on' : 'advanced logging: OFF';
    row.append(k, v);
    installs.append(row);
    if (i.advanced_logging_known && !i.advanced_logging) {
      const help = document.createElement('p');
      help.className = 'note warn';
      help.textContent = 'Turn it on in ' + i.advanced_logging_help + '.';
      installs.append(help);
    }
  }
}

function renderReports(s) {
  const list = $('report-list');
  list.replaceChildren();
  if (!s.reports.length) {
    const li = document.createElement('li');
    li.className = 'note';
    li.textContent = 'No reports yet. They appear here the moment a fight is uploaded.';
    list.append(li);
    return;
  }
  for (const r of s.reports) {
    const li = document.createElement('li');
    const when = document.createElement('div');
    when.className = 'when';
    when.textContent = new Date(r.started_at).toLocaleString() +
      ' · ' + r.fights + (r.fights === 1 ? ' fight' : ' fights') +
      (r.done ? ' · complete' : ' · in progress');
    const title = document.createElement('div');
    if (r.url) {
      const a = document.createElement('a');
      a.href = r.url;
      a.textContent = r.zone || r.report_id;
      a.target = '_blank';
      a.rel = 'noreferrer';
      title.append(a);
    } else {
      title.textContent = (r.zone || r.key) + ' — not uploaded yet';
    }
    li.append(title, when);
    list.append(li);
  }
}

function renderDevice(s) {
  $('paired-dot').className = 'dot ' + (s.paired ? 'on' : 'bad');
  text('paired-state', s.paired ? 'Paired' : 'Not paired');
  text('token-backend', s.token_backend);
  text('api-base', s.api_base_url);
}

function renderSettings(s) {
  if (editing) return;
  $('visibility').value = s.visibility;
  $('wow-paths').value = s.installs.map((i) => i.path).join('\n');
  const select = $('logging-character');
  const chosen = s.logging_character
    ? [s.logging_character.region, s.logging_character.ruleset, s.logging_character.name].join('|')
    : '';
  select.replaceChildren();
  const none = document.createElement('option');
  none.value = '';
  none.textContent = 'Not set';
  select.append(none);
  for (const c of s.characters) {
    const o = document.createElement('option');
    o.value = [c.region, c.ruleset, c.name].join('|');
    o.textContent = c.name + ' · ' + c.ruleset + ' · ' + c.region;
    select.append(o);
  }
  select.value = chosen;
}

function render(s) {
  latest = s;
  renderStatus(s);
  renderReports(s);
  renderDevice(s);
  renderSettings(s);
}

async function refresh() {
  try {
    const r = await api('api/status');
    if (r.ok) render(r.data);
  } catch (e) {
    const note = $('logging-note');
    note.className = 'note bad';
    note.textContent = 'The companion stopped answering: ' + e;
  }
}

function report(id, r) {
  const note = $(id);
  note.className = r.ok ? 'note' : 'note bad';
  note.textContent = r.ok ? 'Saved.' : r.error.message;
  if (r.ok) render(r.data);
}

$('pair').addEventListener('click', async () => {
  report('device-note', await api('api/pair', { code: $('pair-code').value.trim() }));
});

$('unpair').addEventListener('click', async () => {
  report('device-note', await api('api/unpair', {}));
});

for (const id of ['visibility', 'wow-paths', 'logging-character']) {
  $(id).addEventListener('input', () => { editing = true; });
}

$('save').addEventListener('click', async () => {
  const picked = $('logging-character').value;
  const [region, ruleset, name] = picked ? picked.split('|') : [];
  const r = await api('api/settings', {
    visibility: $('visibility').value,
    wow_paths: $('wow-paths').value.split('\n').map((s) => s.trim()).filter(Boolean),
    logging_character: picked ? { region, ruleset, name } : null,
  });
  editing = false;
  report('settings-note', r);
});

refresh();
setInterval(refresh, 2000);
```

- [ ] **Step 6: Write `companion/ui/assets.go`**

```go
// companion/ui/assets.go
// Package ui is the companion's window: three files, embedded in the
// binary and served over loopback to a native webview. It cannot
// import the site's stylesheet, so the design tokens it needs are
// copied into styles.css and kept in step with
// web/src/styles/tokens.css by hand.
package ui

import "embed"

// Assets is index.html and what it loads.
//
//go:embed index.html app.js styles.css
var Assets embed.FS
```

- [ ] **Step 7: Write `internal/app/app.go`**

```go
// companion/internal/app/app.go
// Package app is the companion assembled: configuration, the device
// token, the game installs, the live pipeline, the addon sync and the
// updater behind one object with a local HTTP interface. The tray and
// the webview are a window onto this; the headless build runs it with
// no window at all.
package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"runtime"
	"slices"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/pipeline"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
	"github.com/jhunthrop/foreversixty/companion/internal/wow"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
)

// TickEvery is how often the companion looks at the log. One second
// keeps a fight visible inside the three-second budget without
// costing anything measurable.
const TickEvery = time.Second

// Options configures an App.
type Options struct {
	Dirs   paths.Dirs
	Config config.Config
	Secret secret.Store
	Log    *slog.Logger
	// Now is the clock. Tests replace it.
	Now func() time.Time
	// Updater options are passed through; a zero value means the
	// defaults, which is no updates when the build has no key.
	Updater updater.Options
	// Retry is the upload retry policy. Zero means the client's
	// default; tests shrink it so an outage costs milliseconds.
	Retry client.Retry
}

// App is the whole companion.
type App struct {
	o     Options
	log   *slog.Logger
	now   func() time.Time
	token string // the local UI token

	mu       sync.Mutex
	cfg      config.Config
	installs []wow.Install

	client  *client.Client
	queue   *queue.Queue
	pipe    *pipeline.Pipeline
	sync    *addon.Sync
	updater *updater.Updater
}

// New assembles the companion over an already-resolved set of
// directories and a loaded configuration.
func New(o Options) (*App, error) {
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	if o.Now == nil {
		o.Now = time.Now
	}
	if o.Secret == nil {
		o.Secret = secret.ConfigFile{Path: o.Dirs.ConfigFile()}
	}
	tok, err := localToken()
	if err != nil {
		return nil, err
	}
	a := &App{o: o, log: o.Log, now: o.Now, token: tok, cfg: o.Config}
	a.installs = a.detect()

	a.client, err = client.New(client.Options{
		BaseURL: a.cfg.APIBaseURL,
		Token:   a.deviceToken,
		Retry:   o.Retry,
		Log:     a.log,
	})
	if err != nil {
		return nil, err
	}
	a.queue, err = queue.Open(o.Dirs.Queue)
	if err != nil {
		return nil, err
	}
	a.pipe, err = pipeline.New(pipeline.Options{
		StateDir: o.Dirs.State,
		Client:   a.client,
		Queue:    a.queue,
		Watch:    watch.New(watch.Options{Dir: a.logsDir()}),
		Config:   a.cfg,
		Log:      a.log,
	})
	if err != nil {
		return nil, err
	}
	a.sync = addon.New(addon.Options{Paths: a.savedVariables, API: a.client, Log: a.log})
	uo := o.Updater
	uo.Dir, uo.Log = o.Dirs.Update, a.log
	a.updater = updater.New(uo)
	return a, nil
}

// Close releases what the app holds.
func (a *App) Close() error { return a.pipe.Close() }

// localToken is the secret in the UI's URL. The webview cannot set a
// header, so the token rides in the path and every asset and API call
// inherits it.
func localToken() (string, error) {
	var b [24]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b[:]), nil
}

// deviceToken reads the stored token, empty when unpaired.
func (a *App) deviceToken() string {
	tok, err := a.o.Secret.Token()
	if err != nil {
		a.log.Warn("could not read the device token", "component", "app", "err", err.Error())
		return ""
	}
	return tok
}

// detect lists the installs: the configured paths first, then
// anything found in the usual places.
func (a *App) detect() []wow.Install {
	var out []wow.Install
	seen := map[string]bool{}
	add := func(in wow.Install) {
		if !seen[in.Path] {
			seen[in.Path] = true
			out = append(out, in)
		}
	}
	for _, p := range a.cfg.WoWPaths {
		found, err := wow.Pick(p)
		if err != nil {
			a.log.Warn("a configured game folder is not readable",
				"component", "app", "path", p, "err", err.Error())
			continue
		}
		for _, in := range found {
			add(in)
		}
	}
	for _, in := range wow.Detect() {
		add(in)
	}
	return out
}

// logsDir is the Logs directory the watcher follows: the first
// install, or a directory that does not exist when there is none,
// which the watcher handles as "nothing to do".
func (a *App) logsDir() string {
	if len(a.installs) == 0 {
		return a.o.Dirs.Home
	}
	return a.installs[0].Logs
}

// savedVariables is every account's addon file across every install.
func (a *App) savedVariables() []string {
	a.mu.Lock()
	installs := slices.Clone(a.installs)
	a.mu.Unlock()
	var out []string
	for _, in := range installs {
		paths, err := in.SavedVariables(addon.SavedVariablesName)
		if err != nil {
			continue
		}
		out = append(out, paths...)
	}
	return out
}

// Pair exchanges a pairing code for a device token and stores it.
func (a *App) Pair(ctx context.Context, code string) error {
	name, err := os.Hostname()
	if err != nil || name == "" {
		name = "Forever Sixty companion"
	}
	dev, err := a.client.Claim(ctx, client.Claim{
		Code:     code,
		Name:     name,
		Platform: runtime.GOOS + "/" + runtime.GOARCH,
	})
	if err != nil {
		return err
	}
	if err := a.o.Secret.SetToken(dev.Token); err != nil {
		return err
	}
	a.log.Info("this device is paired", "component", "app", "device_id", dev.DeviceID)
	return nil
}

// Unpair forgets the device token. The reports already uploaded stay
// on the site; revoking the device itself is done on the account page.
func (a *App) Unpair() error { return a.o.Secret.Clear() }

// Settings is what the settings page can change.
type Settings struct {
	Visibility       string               `json:"visibility"`
	WoWPaths         []string             `json:"wow_paths"`
	LoggingCharacter *character.Character `json:"logging_character"`
}

// SaveSettings validates and stores the settings, then re-detects the
// installs so a newly added folder takes effect without a restart.
func (a *App) SaveSettings(in Settings) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	cfg := a.cfg
	cfg.ReportVisibility = in.Visibility
	cfg.WoWPaths = in.WoWPaths
	if cfg.WoWPaths == nil {
		cfg.WoWPaths = []string{}
	}
	cfg.LoggingCharacter = in.LoggingCharacter
	if cfg.LoggingCharacter != nil && !cfg.LoggingCharacter.Valid() {
		return errors.New("the logging character needs a region, a ruleset and a name")
	}
	for _, p := range cfg.WoWPaths {
		if _, err := wow.Pick(p); err != nil {
			return err
		}
	}
	if err := config.Save(a.o.Dirs.ConfigFile(), cfg); err != nil {
		return err
	}
	a.cfg = cfg
	a.installs = a.detect()
	return nil
}

// Snapshot is everything the UI shows.
type Snapshot struct {
	EngineVersion    string               `json:"engine_version"`
	Paired           bool                 `json:"paired"`
	TokenBackend     string               `json:"token_backend"`
	APIBaseURL       string               `json:"api_base_url"`
	SiteBaseURL      string               `json:"site_base_url"`
	Visibility       string               `json:"visibility"`
	LoggingCharacter *character.Character `json:"logging_character,omitempty"`
	Installs         []InstallStatus      `json:"installs"`
	Pipeline         pipeline.Status      `json:"pipeline"`
	Reports          []ReportRow          `json:"reports"`
	Characters       []addon.Export       `json:"characters"`
	UpdateStaged     bool                 `json:"update_staged"`
}

// InstallStatus is one game folder as the status page shows it.
type InstallStatus struct {
	Path string `json:"path"`
	// Flavor is the folder name, "_classic_era_" and the like.
	Flavor string `json:"flavor"`
	// Advanced is true, false, or unknown when Config.wtf could not
	// be read — which is the normal state before the first login.
	Advanced        bool   `json:"advanced_logging"`
	AdvancedKnown   bool   `json:"advanced_logging_known"`
	AdvancedLogHelp string `json:"advanced_logging_help"`
}

// ReportRow is one report in the reports page.
type ReportRow struct {
	Key       string    `json:"key"`
	ReportID  string    `json:"report_id,omitempty"`
	Zone      string    `json:"zone,omitempty"`
	Fights    int       `json:"fights"`
	StartedAt time.Time `json:"started_at"`
	Done      bool      `json:"done"`
	URL       string    `json:"url,omitempty"`
}

// Snapshot builds the UI's view of the world.
func (a *App) Snapshot() Snapshot {
	a.mu.Lock()
	cfg, installs := a.cfg, slices.Clone(a.installs)
	a.mu.Unlock()

	s := Snapshot{
		EngineVersion:    session.Version,
		Paired:           a.deviceToken() != "",
		TokenBackend:     string(a.o.Secret.Backend()),
		APIBaseURL:       cfg.APIBaseURL,
		SiteBaseURL:      cfg.SiteBaseURL,
		Visibility:       cfg.ReportVisibility,
		LoggingCharacter: cfg.LoggingCharacter,
		Pipeline:         a.pipe.Status(),
		Installs:         make([]InstallStatus, 0, len(installs)),
		Reports:          []ReportRow{},
		Characters:       []addon.Export{},
	}
	for _, in := range installs {
		row := InstallStatus{Path: in.Path, Flavor: in.Flavor,
			AdvancedLogHelp: wow.AdvancedLoggingHelp}
		if on, err := in.AdvancedLogging(); err == nil {
			row.Advanced, row.AdvancedKnown = on, true
		}
		s.Installs = append(s.Installs, row)
	}
	reports, err := state.List(a.o.Dirs.State)
	if err != nil {
		a.log.Warn("could not list the reports", "component", "app", "err", err.Error())
	}
	for i := len(reports) - 1; i >= 0; i-- {
		r := reports[i]
		row := ReportRow{Key: r.Key, ReportID: r.ReportID, Zone: r.Zone,
			Fights: r.Fights, StartedAt: r.StartedAt, Done: r.Done}
		if r.ReportID != "" {
			row.URL = cfg.SiteBaseURL + "/reports/" + r.ReportID
		}
		s.Reports = append(s.Reports, row)
	}
	for _, p := range a.savedVariables() {
		found, err := addon.ScanExports(p)
		if err != nil {
			continue
		}
		s.Characters = append(s.Characters, found...)
	}
	if _, err := os.Stat(updater.PendingPath(a.o.Dirs.Update)); err == nil {
		s.UpdateStaged = true
	}
	return s
}

// Step is one pass of everything on a timer: tail, snapshot, upload,
// addon sync and the update check. It returns no error because the
// companion must keep logging through every one of them failing.
func (a *App) Step(ctx context.Context, now time.Time) {
	if err := a.pipe.Tick(now); err != nil {
		a.log.Error("the tail failed", "component", "app", "err", err.Error())
	}
	if err := a.pipe.Live(ctx, now); err != nil {
		a.log.Error("a live snapshot failed", "component", "app", "err", err.Error())
	}
	if a.deviceToken() == "" {
		return // nothing below here works unpaired
	}
	if err := a.pipe.Drain(ctx); err != nil && !errors.Is(err, context.Canceled) {
		a.log.Warn("uploads are behind", "component", "app", "err", err.Error())
	}
	if err := a.sync.Poll(ctx, now); err != nil {
		a.log.Warn("the addon sync did not complete", "component", "app", "err", err.Error())
	}
	if staged, err := a.updater.Poll(ctx, now); err != nil {
		a.log.Warn("the update check failed", "component", "app", "err", err.Error())
	} else if staged {
		a.log.Info("an update is ready and will be applied on the next launch",
			"component", "app")
	}
}

// Run steps until the context ends.
func (a *App) Run(ctx context.Context) error {
	t := time.NewTicker(TickEvery)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			a.Step(ctx, a.now())
		}
	}
}

// Serve starts the local UI server on a loopback port and returns the
// URL the webview should open. The token is in the path, so a page on
// another origin cannot guess it.
func (a *App) Serve() (string, net.Listener, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", nil, fmt.Errorf("listen on loopback: %w", err)
	}
	srv := &http.Server{Handler: a.Handler(), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Error("the local UI server stopped", "component", "ui", "err", err.Error())
		}
	}()
	return fmt.Sprintf("http://%s/%s/", ln.Addr().String(), a.token), ln, nil
}
```

- [ ] **Step 8: Write `internal/app/http.go`**

```go
// companion/internal/app/http.go
// The local HTTP interface the webview talks to. Everything hangs off
// the session token in the path; the JSON below is the whole contract
// between ui/app.js and the Go side.
package app

import (
	"encoding/json"
	"io"
	"io/fs"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/ui"
)

// MaxRequestBody bounds a request from the local page. Nothing it
// sends is larger than a settings object.
const MaxRequestBody = 64 << 10

// Handler is the local UI server's routes.
func (a *App) Handler() http.Handler {
	assets, err := fs.Sub(ui.Assets, ".")
	if err != nil {
		panic(err) // the embedded filesystem is a build-time fact
	}
	inner := http.NewServeMux()
	inner.Handle("GET /", http.FileServerFS(assets))
	inner.HandleFunc("GET /api/status", a.handleStatus)
	inner.HandleFunc("POST /api/pair", a.handlePair)
	inner.HandleFunc("POST /api/unpair", a.handleUnpair)
	inner.HandleFunc("POST /api/settings", a.handleSettings)

	outer := http.NewServeMux()
	outer.Handle("/"+a.token+"/", http.StripPrefix("/"+a.token, inner))
	outer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})
	return outer
}

// writeJSON answers with the same envelope shape the API uses, so
// ui/app.js has one way to read a response.
func writeJSON(w http.ResponseWriter, status int, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	body := map[string]any{"ok": message == "", "data": data, "error": nil}
	if message != "" {
		body["error"] = map[string]string{"message": message}
	}
	json.NewEncoder(w).Encode(body)
}

func decode(w http.ResponseWriter, r *http.Request, out any) bool {
	b, err := io.ReadAll(io.LimitReader(r.Body, MaxRequestBody))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil, "could not read the request")
		return false
	}
	if err := json.Unmarshal(b, out); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, "could not read the request: "+err.Error())
		return false
	}
	return true
}

func (a *App) handleStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handlePair(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Code string `json:"code"`
	}
	if !decode(w, r, &in) {
		return
	}
	if err := a.Pair(r.Context(), in.Code); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handleUnpair(w http.ResponseWriter, r *http.Request) {
	if err := a.Unpair(); err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handleSettings(w http.ResponseWriter, r *http.Request) {
	var in Settings
	if !decode(w, r, &in) {
		return
	}
	if err := a.SaveSettings(in); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}
```

- [ ] **Step 9: Run the tests**

Run: `cd companion && gofmt -l . && go vet ./internal/app/ ./ui/ && go test ./internal/app/ -race -v`
Expected: PASS, eight tests.

- [ ] **Step 10: Commit**

```bash
git add companion/internal/app/ \
  companion/ui/
git commit -m "feat(companion): the app assembled, behind a loopback UI with four pages" \
  -m "Configuration, the device token, the game installs, the live
pipeline, the addon sync and the updater behind one object, with a
local HTTP interface and a page that polls it. Status, Reports,
Device and Settings, in the site's tokens -- copied into the
companion's own stylesheet, since it cannot import the site's.

The server binds loopback with a random token in the path, because a
webview cannot set a header and every relative fetch inherits a path
prefix. Step returns nothing on purpose: the companion keeps logging
through any part of it failing, and an unpaired device still parses
and queues, so pairing an hour late costs nothing." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 13: The tray, the webview, the drawn icon and the command

**Files:**
- Create: `companion/internal/icon/icon.go`, `companion/internal/shell/shell.go`, `companion/internal/shell/shell_gui.go`, `companion/internal/shell/shell_nogui.go`, `companion/cmd/foreversixty-companion/main.go`
- Test: `companion/internal/icon/icon_test.go`, `companion/internal/shell/shell_test.go`

**Interfaces:**
- Consumes: `app.New`, `(*App).Serve/Run/Close` (Task 12); `config.Load` (Task 1); `logging.New` (Task 1); `paths.Resolve` (Task 1); `secret.Open`, `secret.Migrate` (Task 4); `updater.ApplyPending`, `updater.PublicKey`, `updater.Version` (Task 11).
- Produces:
  - `icon.Size` (32), `icon.Gold` (the site's `--color-gold`), `icon.PNG() []byte`, `icon.ICO() []byte`.
  - `shell.Options{URL, Title string; Headless bool; Width, Height int; OnQuit func()}` and `shell.Run(ctx, Options) error`, in two builds.
  - The `foreversixty-companion` command with `-headless`, `-debug` and `-version`.

The icon is drawn rather than shipped: one gold ring, supersampled four times per pixel, encoded as PNG for macOS and Linux and wrapped in a 22-byte ICO container for Windows, which is the format its tray wants. No binary asset in the repository, and a test that decodes both.

`shell` is the only package that may import systray or webview, and it does so only in `shell_gui.go`, behind `//go:build !nogui`. CI compiles and tests with `-tags nogui`, so no runner needs WebKit, GTK or WebView2; the release job installs them per platform. Both builds export the same `Run` and both honour `-headless` at runtime, so a player on a machine with no working tray can still run the companion and open the printed URL.

Startup order in `main` is deliberate: a staged update becomes the program *before* anything else opens a file, because Windows will not replace a running binary.

- [ ] **Step 1: Write the failing icon test**

```go
package icon

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"testing"
)

func TestThePNGDecodesAtTheRightSize(t *testing.T) {
	img, err := png.Decode(bytes.NewReader(PNG()))
	if err != nil {
		t.Fatal(err)
	}
	if b := img.Bounds(); b.Dx() != Size || b.Dy() != Size {
		t.Fatalf("bounds = %v", b)
	}
	// The middle is transparent and the ring is gold.
	if _, _, _, a := img.At(Size/2, Size/2).RGBA(); a != 0 {
		t.Error("the middle of the ring is not transparent")
	}
	// RGBA() is alpha-premultiplied, so the colour is read straight
	// off the NRGBA image the encoder was given.
	rgba, ok := img.(*image.NRGBA)
	if !ok {
		t.Fatalf("decoded image is %T", img)
	}
	top := rgba.NRGBAAt(Size/2, 2)
	if top.A == 0 {
		t.Fatal("the top of the ring is transparent")
	}
	if top.R != Gold.R || top.G != Gold.G || top.B != Gold.B {
		t.Errorf("ring colour = %+v, want %+v", top, Gold)
	}
}

func TestTheICOWrapsThePNG(t *testing.T) {
	ico := ICO()
	if len(ico) != 22+len(PNG()) {
		t.Fatalf("ICO is %d bytes, want 22 + %d", len(ico), len(PNG()))
	}
	var head [3]uint16
	if err := binary.Read(bytes.NewReader(ico), binary.LittleEndian, &head); err != nil {
		t.Fatal(err)
	}
	if head != [3]uint16{0, 1, 1} {
		t.Fatalf("ICONDIR = %v, want reserved 0, type 1, count 1", head)
	}
	if ico[6] != Size || ico[7] != Size {
		t.Errorf("entry size = %d×%d", ico[6], ico[7])
	}
	if !bytes.Equal(ico[22:], PNG()) {
		t.Error("the embedded image is not the PNG")
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd companion && go test ./internal/icon/`
Expected: FAIL, "undefined: PNG".

- [ ] **Step 3: Write `internal/icon/icon.go`**

```go
// companion/internal/icon/icon.go
// Package icon draws the tray icon rather than shipping a binary
// asset: one gold ring on transparency, in the site's accent colour,
// encoded as PNG for macOS and Linux and as a PNG inside an ICO
// container for Windows, which is the format its tray wants.
package icon

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"image/png"
	"math"
	"sync"
)

// Size is the icon's edge in pixels. 32 is what every tray asks for
// and scales acceptably on a retina menu bar.
const Size = 32

// Gold is --color-gold from web/src/styles/tokens.css.
var Gold = color.NRGBA{R: 0xe5, G: 0xb9, B: 0x55, A: 0xff}

var (
	once sync.Once
	png_ []byte
	ico_ []byte
)

// draw paints a ring, antialiased by supersampling the coverage of
// each pixel four times over.
func draw() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, Size, Size))
	const outer, inner = Size/2 - 1.5, Size/2 - 6.5
	centre := float64(Size) / 2
	for y := range Size {
		for x := range Size {
			var hits int
			for _, dy := range []float64{0.25, 0.75} {
				for _, dx := range []float64{0.25, 0.75} {
					px, py := float64(x)+dx-centre, float64(y)+dy-centre
					d := math.Hypot(px, py)
					if d <= outer && d >= inner {
						hits++
					}
				}
			}
			if hits == 0 {
				continue
			}
			c := Gold
			c.A = uint8(hits * 255 / 4)
			img.SetNRGBA(x, y, c)
		}
	}
	return img
}

func build() {
	var buf bytes.Buffer
	if err := png.Encode(&buf, draw()); err != nil {
		panic(err) // encoding a 32×32 image cannot fail
	}
	png_ = buf.Bytes()

	// ICONDIR, one ICONDIRENTRY, then the PNG itself. Windows has
	// accepted PNG-in-ICO since Vista.
	var ico bytes.Buffer
	binary.Write(&ico, binary.LittleEndian, [3]uint16{0, 1, 1})
	ico.Write([]byte{Size, Size, 0, 0})
	binary.Write(&ico, binary.LittleEndian, [2]uint16{1, 32})
	binary.Write(&ico, binary.LittleEndian, uint32(len(png_)))
	binary.Write(&ico, binary.LittleEndian, uint32(22))
	ico.Write(png_)
	ico_ = ico.Bytes()
}

// PNG is the icon for macOS and Linux.
func PNG() []byte { once.Do(build); return png_ }

// ICO is the icon for Windows.
func ICO() []byte { once.Do(build); return ico_ }
```

- [ ] **Step 4: Run the icon test**

Run: `cd companion && go test ./internal/icon/ -race`
Expected: PASS.

- [ ] **Step 5: Write the shell test**

```go
package shell

import (
	"context"
	"testing"
	"time"
)

func TestDefaultsFillInTheWindow(t *testing.T) {
	got := Options{}.withDefaults()
	if got.Title != "Forever Sixty" || got.Width != 960 || got.Height != 680 {
		t.Fatalf("defaults = %+v", got)
	}
	kept := Options{Title: "x", Width: 1, Height: 2}.withDefaults()
	if kept.Title != "x" || kept.Width != 1 || kept.Height != 2 {
		t.Fatalf("defaults overrode the caller: %+v", kept)
	}
}

func TestWaitReturnsWhenTheContextEnds(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Millisecond)
	defer cancel()
	if err := wait(ctx); err == nil {
		t.Fatal("wait returned no error on a cancelled context")
	}
}

// TestHeadlessRunReturns covers the path CI takes. In the real build
// Run is headless here too, so no window is opened by the test suite.
func TestHeadlessRunReturns(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if err := Run(ctx, Options{Headless: true, URL: "http://127.0.0.1:1/"}); err == nil {
		t.Fatal("Run returned no error on a cancelled context")
	}
}
```

- [ ] **Step 6: Write `internal/shell/shell.go`**

```go
// companion/internal/shell/shell.go
// Package shell is the window: a tray icon and a native webview over
// the local UI server. It exists in two builds — the real one, which
// needs cgo and each platform's webview libraries, and a headless one
// selected with -tags nogui for CI and for servers. Both export the
// same Run, and both honour the -headless flag at runtime, so a
// player on a machine with no tray can still run the companion.
package shell

import "context"

// Options configures the window.
type Options struct {
	// URL is the local UI server's address, token and all.
	URL string
	// Title is the window title.
	Title string
	// Headless skips the tray and the window entirely.
	Headless bool
	// Width and Height are the window's starting size.
	Width, Height int
	// OnQuit is called when the player quits from the tray.
	OnQuit func()
}

// Defaults fills in the window size.
func (o Options) withDefaults() Options {
	if o.Title == "" {
		o.Title = "Forever Sixty"
	}
	if o.Width == 0 {
		o.Width = 960
	}
	if o.Height == 0 {
		o.Height = 680
	}
	return o
}

// wait blocks until the context ends. It is what headless mode does,
// and what the real shell falls back to when the tray is off.
func wait(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}
```

- [ ] **Step 7: Write `internal/shell/shell_nogui.go`**

```go
//go:build nogui

// companion/internal/shell/shell_nogui.go
// The headless build. CI compiles and tests with -tags nogui so no
// runner needs WebKit, GTK or WebView2 installed; the companion still
// logs, uploads and serves its local UI over loopback.
package shell

import "context"

// Run blocks until the context ends. The local UI server is already
// listening, so a player can open the printed URL in a browser.
func Run(ctx context.Context, o Options) error { return wait(ctx) }
```

- [ ] **Step 8: Write `internal/shell/shell_gui.go`**

```go
//go:build !nogui

// companion/internal/shell/shell_gui.go
// The real window. Both libraries here need cgo and platform
// toolkits: WebKit on macOS, WebView2 on Windows, WebKitGTK on Linux
// (libgtk-3-dev and libwebkit2gtk-4.1-dev). The release workflow
// installs them per runner; `go test -tags nogui` needs none of it.
package shell

import (
	"context"
	"runtime"

	"github.com/getlantern/systray"
	webview "github.com/webview/webview_go"

	"github.com/jhunthrop/foreversixty/companion/internal/icon"
)

// Run shows the tray and the window and blocks until the player quits
// or the context ends. It must be called from the main goroutine:
// every desktop toolkit here requires it.
func Run(ctx context.Context, o Options) error {
	o = o.withDefaults()
	if o.Headless {
		return wait(ctx)
	}
	runtime.LockOSThread()

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle(o.Title)
	w.SetSize(o.Width, o.Height, webview.HintNone)
	w.Navigate(o.URL)

	// Register rather than Run: systray drives its own loop on some
	// platforms and hands the main loop back on others, and the
	// webview owns the main loop here.
	systray.Register(func() {
		if runtime.GOOS == "windows" {
			systray.SetIcon(icon.ICO())
		} else {
			systray.SetTemplateIcon(icon.PNG(), icon.PNG())
		}
		systray.SetTooltip(o.Title)
		open := systray.AddMenuItem("Open Forever Sixty", "Show the companion window")
		systray.AddSeparator()
		quit := systray.AddMenuItem("Quit", "Stop logging and quit")
		go func() {
			for {
				select {
				case <-ctx.Done():
					w.Terminate()
					return
				case <-open.ClickedCh:
					w.Dispatch(func() { w.Navigate(o.URL) })
				case <-quit.ClickedCh:
					if o.OnQuit != nil {
						o.OnQuit()
					}
					w.Terminate()
					return
				}
			}
		}()
	}, func() {})

	go func() {
		<-ctx.Done()
		w.Terminate()
	}()

	w.Run()
	systray.Quit()
	return nil
}
```

- [ ] **Step 9: Add the window dependencies**

```bash
cd companion
go get github.com/getlantern/systray@v1.2.2
go get github.com/webview/webview_go@v0.0.0-20240831120633-6173450d4dd6
go mod tidy
```

- [ ] **Step 10: Write the command**

```go
// companion/cmd/foreversixty-companion/main.go
// Command foreversixty-companion watches the game's combat log,
// uploads fights as they close, keeps the addon in step with the
// site, and shows all of it in a tray window.
//
//	foreversixty-companion              tray and window
//	foreversixty-companion -headless    no window; the URL is printed
//	foreversixty-companion -version     print the version and exit
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/logging"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/shell"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "foreversixty-companion:", err)
		os.Exit(1)
	}
}

func run(args []string, out *os.File) error {
	fs := flag.NewFlagSet("foreversixty-companion", flag.ContinueOnError)
	fs.SetOutput(out)
	headless := fs.Bool("headless", false, "run without a tray icon or a window")
	debug := fs.Bool("debug", false, "log at debug level")
	version := fs.Bool("version", false, "print the version and exit")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *version {
		fmt.Fprintln(out, updater.Version)
		return nil
	}

	dirs, err := paths.Resolve()
	if err != nil {
		return err
	}
	// Before anything else: a staged update becomes the program. It
	// cannot happen while the binary is running on Windows, so it
	// happens here and takes effect on the next launch.
	if exe, err := os.Executable(); err == nil {
		if swapped, err := updater.ApplyPending(dirs.Update, updater.PublicKey, exe); err != nil {
			fmt.Fprintln(os.Stderr, "foreversixty-companion: the staged update was refused:", err)
		} else if swapped {
			fmt.Fprintln(out, "An update was applied. Start the companion again to run it.")
			return nil
		}
	}

	log, logFile, err := logging.New(dirs.Logs, *debug)
	if err != nil {
		return err
	}
	defer logFile.Close()

	cfg, err := config.Load(dirs.ConfigFile())
	if err != nil {
		return err
	}
	store := secret.Open(dirs.ConfigFile())
	if err := secret.Migrate(store, dirs.ConfigFile()); err != nil {
		log.Warn("could not move the device token into the keychain",
			"component", "main", "err", err.Error())
	}

	a, err := app.New(app.Options{Dirs: dirs, Config: cfg, Secret: store, Log: log})
	if err != nil {
		return err
	}
	defer a.Close()

	url, ln, err := a.Serve()
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Info("the companion is running", "component", "main",
		"version", updater.Version, "home", dirs.Home, "ui", url)
	if *headless {
		fmt.Fprintln(out, "Companion UI:", url)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go a.Run(ctx)

	// The shell owns the main goroutine: every desktop toolkit here
	// requires its loop to run on it.
	return shell.Run(ctx, shell.Options{URL: url, Headless: *headless, OnQuit: stop})
}
```

- [ ] **Step 11: Verify both builds and the tests**

```bash
cd companion
gofmt -l .
go vet -tags nogui ./...
go test -tags nogui ./internal/shell/ ./internal/icon/ -race
go build -tags nogui -o /dev/null ./cmd/foreversixty-companion
CGO_ENABLED=1 go build -o /dev/null ./cmd/foreversixty-companion
```

Expected: no gofmt output, no vet output, PASS, and both builds succeed. The cgo build prints
`-Wdeprecated-literal-operator` warnings from `webview.h` on macOS; they are the library's, not
ours, and are not errors.

- [ ] **Step 12: Commit**

```bash
git add companion/internal/icon/ \
  companion/internal/shell/ \
  companion/cmd/ \
  companion/go.mod \
  companion/go.sum
git commit -m "feat(companion): the tray, the window, and the command that runs them" \
  -m "A drawn icon rather than a shipped one -- a gold ring in the site's
accent colour, PNG for macOS and Linux, the same PNG inside an ICO
container for Windows -- so there is no binary asset in the tree and
a test can decode both.

shell is the only package that imports systray or webview, and only
behind !nogui. CI compiles and tests with -tags nogui, so no runner
needs WebKit, GTK or WebView2; both builds export the same Run and
both honour -headless, so a machine with no working tray still runs
the companion over the printed loopback URL.

main applies a staged update before it opens anything else, because
Windows will not replace a running binary." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Task 14: The integration test — a raid night, a network drop and a restart

**Files:**
- Create: `companion/integration/integration_test.go`, `companion/integration/decompress_test.go`

**Interfaces:**
- Consumes: `app.New`, `(*App).Step/Close` (Task 12); `client.Retry` (Task 2); `config.*` (Task 1); `fakeapi.*` (Task 3); `fixture.*` (Task 9); `paths.*` (Task 1); `secret.ConfigFile` (Task 4); `watch.CompleteIdle` (Task 7).
- Produces: nothing. This task adds no production code; it is the assertion the whole module exists to satisfy.

The test writes a fixture log in bursts the way the client does — two bursts per encounter, so a fight straddles a poll — takes the network away for one fight's worth of bursts, restarts the companion for another, and then asserts the four things a raid leader would notice: six fights arrived, in ascending order, once each; the report was completed exactly once; the night produced one report, not three; and the raw chunks concatenate from offset zero to exactly the `final_offset` the completion declared.

- [ ] **Step 1: Write the test**

```go
// companion/integration/integration_test.go
// The whole companion against a fake game and a fake ingest: a raid
// night written in bursts, a network that drops in the middle of it,
// and a companion that is stopped and started again. The one thing
// this asserts is the one thing a raid leader notices: every fight
// arrives, once, in the order it was fought.
package integration_test

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// game is the fake game writer: it appends to a combat log the way
// the client does, a burst at a time, and never rewrites what it has
// written.
type game struct {
	t    *testing.T
	path string
	now  time.Time
}

func (g *game) burst(text string) {
	g.t.Helper()
	f, err := os.OpenFile(g.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		g.t.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		g.t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		g.t.Fatal(err)
	}
	g.now = g.now.Add(time.Second)
	if err := os.Chtimes(g.path, g.now, g.now); err != nil {
		g.t.Fatal(err)
	}
}

// companion is the app over one set of directories, rebuildable so a
// restart can be simulated.
type companion struct {
	t    *testing.T
	dirs paths.Dirs
	cfg  config.Config
	app  *app.App
}

func (c *companion) start() {
	c.t.Helper()
	if c.app != nil {
		c.app.Close()
	}
	a, err := app.New(app.Options{
		Dirs: c.dirs, Config: c.cfg,
		Secret: secret.ConfigFile{Path: c.dirs.ConfigFile()},
		// One attempt with no wait: the outage in this test is
		// simulated, so there is nothing to be patient about.
		Retry: client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
	})
	if err != nil {
		c.t.Fatal(err)
	}
	c.app = a
}

func TestARaidNightSurvivesADropAndARestart(t *testing.T) {
	home := t.TempDir()
	t.Setenv(paths.HomeEnv, home)
	dirs, err := paths.Resolve()
	if err != nil {
		t.Fatal(err)
	}
	install := filepath.Join(t.TempDir(), "_classic_era_")
	logs := filepath.Join(install, "Logs")
	for _, p := range []string{logs, filepath.Join(install, "WTF", "Account", "ACCOUNT#1", "SavedVariables")} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	srv := fakeapi.New()
	defer srv.Close()

	cfg := config.Default()
	cfg.APIBaseURL, cfg.SiteBaseURL = srv.URL, srv.URL
	cfg.WoWPaths = []string{install}
	cfg.DeviceToken = fakeapi.Token // already paired
	if err := config.Save(dirs.ConfigFile(), cfg); err != nil {
		t.Fatal(err)
	}

	c := &companion{t: t, dirs: dirs, cfg: cfg}
	c.start()
	defer func() { c.app.Close() }()
	g := &game{t: t, path: filepath.Join(logs, "WoWCombatLog.txt"), now: t0}

	clock := t0
	step := func() {
		clock = clock.Add(time.Second)
		c.app.Step(context.Background(), clock)
	}

	step() // the companion is running before the player logs in
	g.burst(fixture.Header + fixture.Zone)
	step()

	const fights = 6
	for i := range fights {
		// Each encounter goes out in two bursts, so a fight can
		// straddle a poll the way a real one does.
		body := fixture.Encounter(i) + fixture.Heartbeat(i)
		half := strings.Index(body[len(body)/2:], "\n") + len(body)/2 + 1
		g.burst(body[:half])
		step()
		g.burst(body[half:])
		step()

		switch i {
		case 2:
			// The network drops for one fight's worth of bursts.
			srv.Offline(true)
		case 3:
			srv.Offline(false)
		case 4:
			// The player quits the companion and starts it again.
			c.start()
		}
	}

	// The raid ends: fifteen quiet minutes complete the report.
	clock = clock.Add(watch.CompleteIdle + time.Minute)
	c.app.Step(context.Background(), clock)
	for range 3 {
		step()
	}

	order := srv.Order()
	var arrived []int
	completed := 0
	for _, e := range order {
		switch {
		case strings.HasPrefix(e, "fight "):
			n, err := strconv.Atoi(strings.TrimPrefix(e, "fight "))
			if err != nil {
				t.Fatalf("unreadable event %q", e)
			}
			arrived = append(arrived, n)
		case strings.HasPrefix(e, "complete "):
			completed++
		}
	}
	if len(arrived) != fights {
		t.Fatalf("%d fights arrived, want %d: %v", len(arrived), fights, order)
	}
	for i := 1; i < len(arrived); i++ {
		if arrived[i-1] >= arrived[i] {
			t.Fatalf("fights arrived out of order: %v", arrived)
		}
	}
	if completed != 1 {
		t.Fatalf("the report was completed %d times: %v", completed, order)
	}
	if len(srv.Reports()) != 1 {
		t.Fatalf("the night produced %d reports, want one", len(srv.Reports()))
	}

	for id, rep := range srv.Reports() {
		if rep.Complete == nil {
			t.Fatalf("report %s was never completed", id)
		}
		if len(rep.Fights) != fights {
			t.Errorf("report %s holds %d fights", id, len(rep.Fights))
		}
		// The raw stream must be contiguous from zero to the final
		// offset the completion declared, because the server
		// reassembles it by offset.
		var offsets []int
		for o := range rep.Raw {
			offsets = append(offsets, int(o))
		}
		sort.Ints(offsets)
		if len(offsets) == 0 || offsets[0] != 0 {
			t.Fatalf("raw chunks start at %v", offsets)
		}
		total := 0
		for _, o := range offsets {
			if o != total {
				t.Fatalf("a raw chunk is missing before offset %d: %v", o, offsets)
			}
			decoded, err := decompress(rep.Raw[int64(o)])
			if err != nil {
				t.Fatal(err)
			}
			total += len(decoded)
		}
		if int64(total) != rep.Complete.FinalOffset {
			t.Errorf("the raw stream is %d bytes, the completion declared %d",
				total, rep.Complete.FinalOffset)
		}
		if rep.Complete.Health.ParseErrors != 0 {
			t.Errorf("health = %+v", rep.Complete.Health)
		}
		for n, f := range rep.Fights {
			if len(f.Events) == 0 || len(f.Metrics) == 0 || len(f.RawRange.SHA256) != 64 {
				t.Errorf("fight %d is incomplete: %d event bytes, %d metrics, hash %q",
					n, len(f.Events), len(f.Metrics), f.RawRange.SHA256)
			}
		}
	}
}
```

```go
package integration_test

import "github.com/klauspost/compress/zstd"

// decompress unpacks a raw chunk the way the server does.
func decompress(packed []byte) ([]byte, error) {
	d, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	return d.DecodeAll(packed, nil)
}
```

- [ ] **Step 2: Run it**

Run: `cd companion && gofmt -l . && go test -tags nogui ./integration/ -race -v`
Expected: PASS in under a second. The retry policy is pinned to one attempt with no wait, because the outage is simulated and there is nothing to be patient about.

- [ ] **Step 3: Commit**

```bash
git add companion/integration/
git commit -m "test(companion): a raid night with a dropped network and a restart" \
  -m "A fake game writes the fixture log in bursts, two per encounter, so
a fight straddles a poll the way a real one does. The network goes
away for one fight and comes back; the companion is stopped and
started for another.

The assertions are the ones a raid leader would make: six fights
arrived, in order, once each; the report completed exactly once; the
night is one report and not three; and the raw chunks run from offset
zero to exactly the final offset the completion declared, which is
what lets the server reassemble the log and check our arithmetic." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---


## Task 15: The release workflow and the README

**Files:**
- Create: `.github/workflows/companion.yml`, `companion/README.md`

**Interfaces:**
- Consumes: `updater.Version` and `updater.PublicKey` (Task 11), which the build step sets with `-ldflags -X`; `updater.AssetName` (Task 11), whose naming the workflow must match exactly or the updater will never find its asset.
- Produces: a `test` job on every push and pull request, and `build` plus `release` jobs on a `companion-v*` tag.

Five decisions in this file are load-bearing:

**No paths filter on `push`.** GitHub applies one filter set to every ref in a push block, and a tag push changes no paths, so a `paths:` filter here would silently never release. Pull requests, where the cost of running matters, are filtered.

**`GOWORK: off` throughout.** The module is tested as it ships. A companion that only builds inside the repository workspace is a companion that does not build for someone who clones it.

**The coverage floor is computed on a filtered profile.** `cmd/`, `internal/shell/` and `ui/` are dropped: the tray, the webview and three static files, which no Go test can exercise without a desktop. Everything else must reach 80%.

**Intel macOS is cross-compiled from the arm64 runner.** Xcode's clang is a universal compiler, so `-arch x86_64` in `CGO_CFLAGS`/`CGO_CXXFLAGS`/`CGO_LDFLAGS` is all it needs, and GitHub's Intel runners are retired. This was verified: the cross build produces a `Mach-O 64-bit executable x86_64`.

**Signing is conditional and quiet about it.** The secrets are declared as job-level `env` so a step's `if:` can test them — step-level `env` is not available to its own `if`. When `APPLE_*` or `WINDOWS_SIGNING_*` are absent the build still happens and a `::notice::` says the binary is unsigned; when `MINISIGN_SECRET_KEY` is absent the release still publishes and a notice says installed companions will not update themselves to it, which is exactly what `updater.ErrNoKey` guarantees.

- [ ] **Step 1: Write the workflow**

```yaml
name: companion
on:
  # No paths filter on push: GitHub applies one filter set to every
  # ref in a push block, and a tag push changes no paths, so a paths
  # filter here would silently never release. Pull requests, where
  # the cost of running matters, are filtered.
  push:
    branches: [main]
    tags: ['companion-v*']
  pull_request:
    paths: ['companion/**', 'logs/**', '.github/workflows/companion.yml']
  workflow_dispatch:
permissions: { contents: read }
concurrency:
  group: companion-${{ github.ref }}
  cancel-in-progress: true
env:
  # The module is tested as it ships, not through the repo-root
  # workspace: a companion that only builds inside go.work is a
  # companion that does not build for a contributor who clones it.
  GOWORK: off
  # Build tag for every job that must not link a desktop toolkit.
  HEADLESS_TAGS: nogui
jobs:
  test:
    if: "!startsWith(github.ref, 'refs/tags/')"
    runs-on: ubuntu-latest
    defaults: { run: { working-directory: companion } }
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version-file: companion/go.mod, cache-dependency-path: companion/go.sum }
      - name: gofmt
        run: |
          unformatted=$(gofmt -l .)
          if [ -n "$unformatted" ]; then
            echo "these files are not gofmt'd:"
            echo "$unformatted"
            exit 1
          fi
      - run: go vet -tags "$HEADLESS_TAGS" ./...
      - run: go test -tags "$HEADLESS_TAGS" ./... -race -coverprofile=cover.out -coverpkg=./...
      - name: coverage floor
        # The window is excluded: cmd, internal/shell and ui are the
        # tray, the webview and three static files, none of which a Go
        # test can exercise without a desktop.
        run: |
          grep -v -E '/(cmd|internal/shell|ui)/' cover.out > cover-core.out
          total=$(go tool cover -func=cover-core.out | awk '/^total:/ {print substr($3, 1, length($3)-1)}')
          echo "coverage outside the window: ${total}%"
          awk -v t="$total" 'BEGIN { exit (t + 0 >= 80) ? 0 : 1 }' || {
            echo "coverage ${total}% is under the 80% floor"
            exit 1
          }
      - name: the window still compiles
        # Nothing here runs it; this only proves the cgo build is not
        # broken by a change to the headless side.
        run: |
          sudo apt-get update
          sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.1-dev
          go build ./...

  build:
    if: startsWith(github.ref, 'refs/tags/companion-v')
    strategy:
      fail-fast: false
      matrix:
        include:
          - { runner: macos-14, goos: darwin, goarch: arm64 }
          # Intel macOS is cross-compiled from the arm64 runner:
          # Xcode's clang is a universal compiler, so -arch x86_64 is
          # all it needs, and GitHub's Intel runners are retired.
          - { runner: macos-14, goos: darwin, goarch: amd64, cgoarch: '-arch x86_64 -mmacosx-version-min=11.0' }
          - { runner: windows-latest, goos: windows, goarch: amd64 }
          - { runner: ubuntu-latest, goos: linux, goarch: amd64 }
    runs-on: ${{ matrix.runner }}
    defaults: { run: { shell: bash, working-directory: companion } }
    env:
      APPLE_SIGNING_IDENTITY: ${{ secrets.APPLE_SIGNING_IDENTITY }}
      APPLE_CERTIFICATE_P12: ${{ secrets.APPLE_CERTIFICATE_P12 }}
      APPLE_CERTIFICATE_PASSWORD: ${{ secrets.APPLE_CERTIFICATE_PASSWORD }}
      APPLE_ID: ${{ secrets.APPLE_ID }}
      APPLE_APP_PASSWORD: ${{ secrets.APPLE_APP_PASSWORD }}
      APPLE_TEAM_ID: ${{ secrets.APPLE_TEAM_ID }}
      WINDOWS_SIGNING_CERTIFICATE: ${{ secrets.WINDOWS_SIGNING_CERTIFICATE }}
      WINDOWS_SIGNING_PASSWORD: ${{ secrets.WINDOWS_SIGNING_PASSWORD }}
      MINISIGN_PUBLIC_KEY: ${{ vars.MINISIGN_PUBLIC_KEY }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version-file: companion/go.mod, cache-dependency-path: companion/go.sum }

      - name: Linux webview libraries
        if: matrix.goos == 'linux'
        run: |
          sudo apt-get update
          sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.1-dev

      - name: Windows C toolchain
        if: matrix.goos == 'windows'
        run: |
          if ! command -v gcc >/dev/null; then
            choco install mingw -y --no-progress
          fi
          gcc --version

      - name: Build
        env:
          CGOARCH: ${{ matrix.cgoarch }}
        run: |
          version="${GITHUB_REF_NAME#companion-v}"
          name="foreversixty-companion_${{ matrix.goos }}_${{ matrix.goarch }}"
          if [ "${{ matrix.goos }}" = "windows" ]; then name="${name}.exe"; fi
          mkdir -p dist
          if [ -n "$CGOARCH" ]; then
            export CGO_CFLAGS="$CGOARCH" CGO_CXXFLAGS="$CGOARCH" CGO_LDFLAGS="$CGOARCH"
          fi
          CGO_ENABLED=1 GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} \
            go build -trimpath \
              -ldflags "-s -w \
                -X github.com/jhunthrop/foreversixty/companion/internal/updater.Version=${version} \
                -X github.com/jhunthrop/foreversixty/companion/internal/updater.PublicKey=${MINISIGN_PUBLIC_KEY}" \
              -o "dist/${name}" ./cmd/foreversixty-companion
          ls -l dist
          echo "ASSET=${name}" >> "$GITHUB_ENV"

      - name: Smoke test the binary
        # Every target but Intel macOS runs natively on its runner;
        # that one is cross-compiled, and Rosetta runs it anyway.
        run: ./dist/${ASSET} -version

      - name: Sign and notarize for macOS
        if: matrix.goos == 'darwin' && env.APPLE_SIGNING_IDENTITY != ''
        run: |
          keychain="$RUNNER_TEMP/signing.keychain-db"
          security create-keychain -p "$APPLE_CERTIFICATE_PASSWORD" "$keychain"
          security set-keychain-settings -lut 900 "$keychain"
          security unlock-keychain -p "$APPLE_CERTIFICATE_PASSWORD" "$keychain"
          echo "$APPLE_CERTIFICATE_P12" | base64 --decode > "$RUNNER_TEMP/cert.p12"
          security import "$RUNNER_TEMP/cert.p12" -k "$keychain" \
            -P "$APPLE_CERTIFICATE_PASSWORD" -T /usr/bin/codesign
          security set-key-partition-list -S apple-tool:,apple:,codesign: \
            -s -k "$APPLE_CERTIFICATE_PASSWORD" "$keychain" >/dev/null
          security list-keychain -d user -s "$keychain" login.keychain-db
          codesign --force --timestamp --options runtime \
            --sign "$APPLE_SIGNING_IDENTITY" "dist/${ASSET}"
          codesign --verify --strict --verbose=2 "dist/${ASSET}"
          ditto -c -k --keepParent "dist/${ASSET}" "$RUNNER_TEMP/notarize.zip"
          xcrun notarytool submit "$RUNNER_TEMP/notarize.zip" \
            --apple-id "$APPLE_ID" --team-id "$APPLE_TEAM_ID" \
            --password "$APPLE_APP_PASSWORD" --wait
          rm -f "$RUNNER_TEMP/cert.p12"
          security delete-keychain "$keychain"

      - name: Unsigned macOS build
        if: matrix.goos == 'darwin' && env.APPLE_SIGNING_IDENTITY == ''
        run: echo "::notice::APPLE_* secrets are not set, so this macOS build is unsigned."

      - name: Sign for Windows
        if: matrix.goos == 'windows' && env.WINDOWS_SIGNING_CERTIFICATE != ''
        shell: pwsh
        working-directory: companion
        run: |
          $pfx = Join-Path $env:RUNNER_TEMP 'signing.pfx'
          [IO.File]::WriteAllBytes($pfx, [Convert]::FromBase64String($env:WINDOWS_SIGNING_CERTIFICATE))
          $signtool = Get-ChildItem 'C:\Program Files (x86)\Windows Kits\10\bin' -Recurse `
            -Filter signtool.exe | Where-Object { $_.FullName -like '*x64*' } | Select-Object -First 1
          & $signtool.FullName sign /fd SHA256 /f $pfx /p $env:WINDOWS_SIGNING_PASSWORD `
            /tr http://timestamp.digicert.com /td SHA256 "dist/$env:ASSET"
          Remove-Item $pfx

      - name: Unsigned Windows build
        if: matrix.goos == 'windows' && env.WINDOWS_SIGNING_CERTIFICATE == ''
        run: echo "::notice::WINDOWS_SIGNING_* secrets are not set, so this Windows build is unsigned."

      - uses: actions/upload-artifact@v4
        with:
          name: companion-${{ matrix.goos }}-${{ matrix.goarch }}
          path: companion/dist/*
          if-no-files-found: error

  release:
    needs: build
    if: startsWith(github.ref, 'refs/tags/companion-v')
    runs-on: ubuntu-latest
    permissions: { contents: write }
    env:
      MINISIGN_SECRET_KEY: ${{ secrets.MINISIGN_SECRET_KEY }}
      MINISIGN_PASSWORD: ${{ secrets.MINISIGN_PASSWORD }}
      GH_TOKEN: ${{ github.token }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/download-artifact@v4
        with: { path: staged, merge-multiple: true }
      - run: ls -l staged

      - name: Sign the assets with minisign
        if: env.MINISIGN_SECRET_KEY != ''
        run: |
          sudo apt-get update
          sudo apt-get install -y minisign
          printf '%s\n' "$MINISIGN_SECRET_KEY" > "$RUNNER_TEMP/minisign.key"
          for f in staged/*; do
            case "$f" in *.minisig) continue;; esac
            printf '%s\n' "$MINISIGN_PASSWORD" | \
              minisign -S -s "$RUNNER_TEMP/minisign.key" -m "$f" \
                -t "Forever Sixty companion ${GITHUB_REF_NAME}"
          done
          rm -f "$RUNNER_TEMP/minisign.key"
          ls -l staged

      - name: Unsigned release
        if: env.MINISIGN_SECRET_KEY == ''
        run: |
          echo "::notice::MINISIGN_SECRET_KEY is not set, so this release carries no \
          signatures and installed companions will not update themselves to it."

      - name: Publish
        run: |
          gh release create "$GITHUB_REF_NAME" staged/* \
            --title "Companion ${GITHUB_REF_NAME#companion-v}" \
            --notes "Desktop companion ${GITHUB_REF_NAME#companion-v} for macOS, Windows and Linux. \
          Each binary is published beside its minisign signature; companion/README.md has the \
          public key and the command to check one."
```

- [ ] **Step 2: Check the workflow parses and says what it should**

```bash
cd "$(git rev-parse --show-toplevel)"
python3 - <<'PY'
import yaml
d = yaml.safe_load(open('.github/workflows/companion.yml'))
assert list(d['jobs']) == ['test', 'build', 'release'], list(d['jobs'])
inc = d['jobs']['build']['strategy']['matrix']['include']
assert {(m['goos'], m['goarch']) for m in inc} == {
    ('darwin', 'arm64'), ('darwin', 'amd64'), ('windows', 'amd64'), ('linux', 'amd64')}
assert d['env']['GOWORK'] == 'off'
print('workflow ok:', len(inc), 'targets')
PY
```

Expected: `workflow ok: 4 targets`.

- [ ] **Step 3: Run the test job's commands locally**

```bash
cd companion
gofmt -l .
go vet -tags nogui ./...
go test -tags nogui ./... -race -coverprofile=cover.out -coverpkg=./...
grep -v -E '/(cmd|internal/shell|ui)/' cover.out > cover-core.out
go tool cover -func=cover-core.out | tail -1
```

Expected: no gofmt output, no vet output, PASS everywhere, and a total at or above 80.0%. The
reference implementation measures 80.8%.

- [ ] **Step 4: Write the README**

The README carries fenced code blocks of its own, so it is shown here inside a four-backtick
fence; the file itself uses ordinary three-backtick fences.

````markdown
# Forever Sixty companion

A desktop app that watches World of Warcraft's combat log, uploads each fight to
[foreversixty.gg](https://foreversixty.gg) as it ends, and keeps the ForeverSixty addon in step
with the builds you chose on the site. One static binary per platform, a tray icon, and a small
window. It never modifies the game's log and never sees your Battle.net password.

## Install

Download the binary for your machine from the
[latest release](https://github.com/jhunthrop/foreversixty/releases?q=companion) and put it
anywhere you like.

| Platform | Asset |
|---|---|
| macOS, Apple silicon | `foreversixty-companion_darwin_arm64` |
| macOS, Intel | `foreversixty-companion_darwin_amd64` |
| Windows | `foreversixty-companion_windows_amd64.exe` |
| Linux | `foreversixty-companion_linux_amd64` |

Every asset is published beside a `.minisig` signature. To check one:

```sh
minisign -Vm foreversixty-companion_darwin_arm64 -P "$(cat MINISIGN_PUBLIC_KEY.txt)"
```

The public key is in `MINISIGN_PUBLIC_KEY.txt` on every release. The companion verifies the same
signature itself before it installs an update, so the check above is for the first download.

macOS and Windows builds are signed and notarized when the project's signing credentials are in
place. If a build is unsigned, macOS will refuse it on the first run: right-click the binary,
choose Open, and confirm once.

Linux needs WebKitGTK for the window: `sudo apt install libgtk-3-0 libwebkit2gtk-4.1-0`. Without
it, run with `-headless` and open the printed URL in a browser.

## Pair it

1. Sign in at [foreversixty.gg/logs](https://foreversixty.gg/logs) with Battle.net or an email
   link.
2. Press **Pair a device**. The site shows a code that is good for ten minutes.
3. Open the companion, go to **Device**, paste the code, press **Pair this device**.

The companion receives an upload token, stores it in your operating system's keychain (macOS
Keychain, Windows Credential Manager, the Linux Secret Service) and shows which one it used on
the Device page. Where no keychain is available it falls back to `config.json`, which is written
readable only by you. Revoke a device any time from **Account → Devices** on the site.

## Log a raid

Turn advanced combat logging on once, in **Options → Network → Advanced Combat Logging**. The
companion's status page tells you if it is off.

Then type `/combatlog` in game at the start of the night and `/combatlog` again at the end — or
leave it on. That is the whole workflow. The companion notices the file growing, uploads each
fight within a few seconds of it ending, streams the raw log up in the background so the server
can check our arithmetic, and closes the report fifteen minutes after the log goes quiet. A gap
longer than thirty minutes starts a new report, so two raids in one night are two reports.

The **Reports** page links every report the moment it has an id, live ones included.

## Settings

| Setting | What it does |
|---|---|
| New reports are | public, unlisted, private or guild — the visibility a report starts with |
| Game folders | one per line; the companion also looks in the usual install paths on its own |
| Logging character | who reports are attributed to, picked from the characters the addon has seen |

Where the companion keeps its files:

| Platform | Directory |
|---|---|
| macOS | `~/Library/Application Support/ForeverSixty` |
| Windows | `%APPDATA%\ForeverSixty` |
| Linux | `~/.config/foreversixty` |

Inside it: `config.json`, `queue/` for uploads that have not landed yet, `state/` for one file
per report, `logs/companion.log`, and `update/` for a downloaded update.

## The addon

If the ForeverSixty addon is installed, the companion reads its saved variables after you log
out and uploads each character's export string, so your real talents and gear reach the planner
without a copy and paste. It writes `ForeverSixtyInbox.lua` beside them every ten minutes with
the builds you chose on the site, which the addon loads the next time the game starts or you
`/reload`. The game rewrites its own saved variables at logout, so an inbox written mid-session
may be replaced; the next pass puts it back.

## Troubleshooting

**"This device is not paired."** Nothing uploads until it is. Fights are still parsed and queued
while you are unpaired, and they go up as soon as you pair, so a missed pairing costs nothing.

**"No World of Warcraft folder was found."** Add the flavour folder — the one containing `WTF`
and `Logs`, such as `_classic_era_` — on the Settings page. The folder above it works too.

**"Advanced logging: OFF."** The report will be accepted, marked, and will lose the views that
need the advanced fields: deaths with health, item level, resources, and positions. Turn it on
in **Options → Network → Advanced Combat Logging** and restart the game.

**Fights are queued and not going up.** Open `logs/companion.log`. Every upload failure is one
line with the report, the fight and the reason. Uploads are strictly in order, so one failing
fight holds the rest behind it deliberately; they go as soon as it does. A fight the server
refuses outright is dropped with an `ERROR` line rather than retried forever.

**Nothing appears while the raid is going.** The engine emits a fight when the line after its
last one arrives, which in a live raid is immediate. A fight that is genuinely the last thing in
the file closes when the report does.

**The window will not open.** Run with `-headless` and open the URL it prints. The interface is
the same; it is served on loopback with a per-session token in the path.

**Start over.** Quit the companion and delete the directory above. Reports already uploaded stay
on the site.

## Build it yourself

```sh
cd companion
go test -tags nogui ./...      # the whole suite, no desktop libraries needed
go build ./cmd/foreversixty-companion
```

The window needs cgo and a platform toolkit: WebKit on macOS (Xcode command line tools),
WebView2 on Windows (a C compiler such as mingw-w64), and `libgtk-3-dev libwebkit2gtk-4.1-dev`
on Linux. `-tags nogui` builds and tests everything else without them.

The companion depends on the `logs/` module in this repository through a `replace`, so the
engine that parses your log in the companion is the same code that parses it on the server.
````

- [ ] **Step 5: Commit**

```bash
git add .github/workflows/companion.yml \
  companion/README.md
git commit -m "ci(companion): build, sign and publish the four targets on a tag" \
  -m "One workflow: a test job on every push and pull request, and on a
companion-v* tag a matrix that builds darwin/arm64, darwin/amd64,
windows/amd64 and linux/amd64, signs what it has credentials for, and
publishes to GitHub Releases with minisign signatures beside each
binary.

There is no paths filter on push, because GitHub applies one filter
set to every ref in a push block and a tag push changes no paths --
a filter there would silently never release. GOWORK is off so the
module is tested as it ships. Intel macOS is cross-compiled from the
arm64 runner with -arch x86_64, since GitHub's Intel runners are
retired.

Missing APPLE_*, WINDOWS_SIGNING_* or MINISIGN_SECRET_KEY secrets
produce a notice and an unsigned artefact rather than a failure; an
unsigned release simply carries no signatures, and a companion built
without a public key does not update itself at all.

The README covers install, verifying a download, pairing, the two
idle rules, where the companion keeps its files, and the six
questions a player actually asks when nothing is uploading." \
  -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k"
```

---

## Running it

```bash
cd companion

# the whole suite, no desktop libraries needed
go test -tags nogui ./... -race

# coverage the way CI measures it
go test -tags nogui ./... -race -coverprofile=cover.out -coverpkg=./...
grep -v -E '/(cmd|internal/shell|ui)/' cover.out > cover-core.out
go tool cover -func=cover-core.out | tail -1

# the real thing
go build -o /tmp/companion ./cmd/foreversixty-companion
FS_COMPANION_HOME=/tmp/fs-companion /tmp/companion -headless -debug
```

`-headless` prints the loopback URL; open it in a browser and the interface is identical to the
webview's. `FS_COMPANION_HOME` keeps a development run out of the real application directory.

## Unresolved on purpose

- **What Forever's combat log calls its realm segment.** The Sept 17 beta log settles it. Until
  then the companion passes the segment through untouched in the contract's `ruleset` field and
  the API maps it. Nothing in this module needs to change when the answer arrives.
- **Whether the game rotates its combat log at all, and under what name.** `watch` handles a
  rename to any `WoWCombatLog*.txt`, a truncation, and a brand new file, so all three answers
  are already covered.
- **The minisign key pair.** User-owned, like the Apple and Windows signing credentials. Until
  `MINISIGN_PUBLIC_KEY` is set as a repository variable, builds carry no key and do not update
  themselves, which is the safe direction.
- **`logging_character`.** The companion does not guess who is logging from the log's own lines;
  the player picks from the characters the addon sync has seen. If the beta log turns out to
  name the logging character unambiguously, that becomes a default rather than a new mechanism.

---


## Self-review

### 1. Spec coverage

Spec section 2, the companion's half: the companion tails the log (Task 7), runs the session
(Task 9), sends one bundle per closed fight with events as Parquet, the summary, the metrics
rows and a hash of the raw lines (Tasks 3 and 9), sends the snapshot every few seconds (Task 9),
streams the raw file up as compressed chunks addressed by offset, resumable (Task 9). Chunks and
bundles are addressed by report id, fight index and byte offset, and re-sent data is recognised —
the `Created`/`Duplicate` distinction in Task 3 and the fakeapi behaviour it is tested against. A
companion that disconnects resumes from its own state and offset: Task 8 and `resume`/`refillRaw`
in Task 9, asserted in Task 14.

Spec section 5, line by line:

| Spec line | Task |
|---|---|
| Native Go app, one static binary per platform | Task 13, Task 15 |
| Menu-bar or tray app with a native-webview window in the site's design system | Tasks 12 and 13 |
| Status, current report link, recent reports, sign-in state | Task 12, the four pages |
| Watches the game's Logs folder, detects a growing WoWCombatLog.txt | Tasks 6 and 7 |
| Runs the engine, sends bundles and snapshots, streams raw chunks, resumable | Task 9 |
| Checks advanced combat logging and shows the exact menu path | Task 6 (`AdvancedLoggingHelp`), Task 12 (shown) |
| Syncs the addon: reads ForeverSixty.lua after logout, uploads exports | Task 10 |
| Writes the addon's inbox file | Task 10 |
| Reports which character is logging | Task 12 setting, Task 9 sends it on report creation |
| Pairing code issues a per-device upload token, never sees passwords | Task 4 |
| Auto-updates from GitHub Releases with signature verification | Task 11, Task 15 |
| Engine version travels with every bundle | `summary.EngineVersion` in the bundle, `session.Version` in the completion (Tasks 3 and 9) |
| Offline: fights queue locally and upload in order on reconnect | Task 5, asserted in Tasks 9 and 14 |
| Never modifies the game's log file | Nothing in `watch` or `pipeline` opens the log for writing |
| Signing and notarization prerequisites are user-owned | Task 15, conditional and quiet |

Spec section 8, the companion's share: structured logs with the report id and engine version
(Task 1's logger, used with `report` and `engine_version` attributes throughout), per-device
hashed tokens scoped to upload (Task 4 — the hashing is the API's side).

Spec section 9, "Companion: tail, offsets, upload queue with a fake server; integration with a
fake game writer, local ingest, drops and restarts; signed build smoke test per platform":
Tasks 7, 5, 3, 14 and the `-version` smoke test in Task 15's build job, respectively.

Interface-contract items with no task, deliberately: the `POST /v1/devices/pair` route and the
site's pairing page (api and web plans), `POST /v1/uploads` whole-file upload (web plan), the
verification and re-parse jobs (api plan), the `/v1/addon/exports` and `/v1/addon/inbox`
endpoints themselves (api plan — the companion only calls them).

### 2. Placeholder scan

No "TBD", no "implement later", no "similar to Task N", no "add error handling" — every
implementation step carries the complete file, and every test step the complete test. Every
command is exact. Every type named in a later task's Interfaces block is defined in an earlier
task's code: `character.Character` (1) → `config`, `client`, `state`, `app`; `client.Error`,
`request`, `do` (2) → 3, 4, 10; `client.FightBundle`, `MetricsRowsOf`, `Live`, `Complete`,
`Stored` (3) → 9; `fakeapi.*` (3) → 4, 9, 12, 14; `secret.Store` (4) → 12, 13, 14;
`queue.Item`, `Lease`, `ErrEmpty` (5) → 9; `wow.Install`, `LogGlob`, `AdvancedLoggingHelp`
(6) → 7, 12; `watch.Event`, `Kind`, `ReadRange`, `CompleteIdle` (7) → 9, 12, 14;
`state.Report` (8) → 9, 12; `fixture.Header/Zone/Encounter/Heartbeat/Log` (9) → 12, 14;
`pipeline.Status` (9) → 12; `addon.Export`, `Inbox`, `InboxPath`, `SavedVariablesName` (10) →
12, `client.PostAddonExports`/`AddonInbox` (10) satisfy `addon.API` (10);
`updater.Options`, `PendingPath`, `ApplyPending`, `PublicKey`, `Version` (11) → 12, 13, 15;
`app.*` (12) → 13, 14; `icon.PNG`/`ICO` (13) → `shell_gui.go` (13).

### 3. Type consistency

Checked by compiling and running every package, in order, in a scratch module with
`replace github.com/jhunthrop/foreversixty/logs => <this worktree>/logs`:

```
gofmt -l .                 0 files
go vet -tags nogui ./...   clean
go vet ./...               clean (cgo build, macOS; webview.h deprecation warnings only)
go test -tags nogui ./... -race -coverprofile=cover.out -coverpkg=./...
                           ok in all 18 packages
coverage outside cmd/, internal/shell/ and ui/:  80.8%
go build -tags nogui ./cmd/foreversixty-companion              ok
CGO_ENABLED=1 go build ./cmd/foreversixty-companion            ok (darwin/arm64)
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-arch x86_64 -mmacosx-version-min=11.0" \
  CGO_CXXFLAGS=… CGO_LDFLAGS=… go build ./cmd/foreversixty-companion
                           ok -> Mach-O 64-bit executable x86_64
```

Names that had to be reconciled while doing it, and now agree everywhere: `client.Character` is
an alias for `character.Character`, not a second struct, so the ruleset field is spelled once;
`state.Report.Offset` is report-relative while `watch.Event.Offset` is a file offset, and
`state.Report.StartOffset` is the only place they meet; `queue.Item.ReportKey` is the local key,
never the server's `report_id`; `pipeline.Status` is embedded in `app.Snapshot` rather than
copied field by field; `addon.API` is satisfied by `*client.Client` without `addon` importing
`client`.

Three bugs the compile-and-run pass caught, all fixed in the code above:

1. `session.State()` refuses to serialise before a layout is settled. `save` now tolerates that
   and `resume` restarts such a report from its first byte — correct, because nothing has closed.
2. A restart dropped the raw bytes buffered below the 4 MiB threshold, so the server could not
   reassemble the stream. `refillRaw` re-reads them out of the log, and `bufferRaw` now measures
   against what is already buffered rather than what was already queued.
3. The report was created lazily, when the first fight closed, so the live snapshots of the very
   first pull had nowhere to go. `Drain` now creates the open report as soon as the network
   allows.

### 4. Final verification

Run once at the whole-branch review, from the repository root:

```bash
cd companion
gofmt -l .
go vet -tags nogui ./...
go test -tags nogui ./... -race -coverprofile=cover.out -coverpkg=./...
grep -v -E '/(cmd|internal/shell|ui)/' cover.out > cover-core.out
go tool cover -func=cover-core.out | tail -1     # at or above 80.0%
go build -tags nogui -o /dev/null ./cmd/foreversixty-companion
CGO_ENABLED=1 go build -o /dev/null ./cmd/foreversixty-companion
cd .. && go work sync && git status --short       # go.work and go.work.sum committed
```

