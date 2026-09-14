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

	// tokMu guards the cached device token. It is its own lock
	// because the token is read on every tick and on every HTTP
	// poll, and must not queue behind a settings save.
	tokMu    sync.Mutex
	tokKnown bool
	tok      string

	mu       sync.Mutex
	cfg      config.Config
	installs []wow.Install
	// watching is the Logs directory the tail is pointed at. The
	// settings page can move it; retarget does the move.
	watching string

	client  *client.Client
	queue   *queue.Queue
	pipe    *pipeline.Pipeline
	sync    *addon.Sync
	exports *addon.Cache
	updater *updater.Updater
}

// ReportRetention is how long a finished report's state file is kept.
// Nothing else ever deleted one, and every status poll reads and
// parses all of them, so without a bound both the Reports page and
// the cost of showing it grow for the life of the install. Six months
// is a little over a raid tier: the season a player is actually in
// stays complete, and what falls off is on the site anyway, which is
// where the Reports page links.
const ReportRetention = 180 * 24 * time.Hour

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
	// The secret store owns the device token. The in-memory copy is
	// blanked so no later write of this configuration can carry a
	// stale token back over the one on disk; SaveSettings reads the
	// live value out of config.json when it saves.
	a.cfg.DeviceToken = ""
	a.installs = a.detect()
	a.watching = a.logsDir()
	a.exports = &addon.Cache{}
	a.prune(a.now())

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
		Watch:    watch.New(watch.Options{Dir: a.watching}),
		Config:   a.cfg,
		Log:      a.log,
	})
	if err != nil {
		return nil, err
	}
	a.sync = addon.New(addon.Options{
		Paths: a.savedVariables, API: a.client, Log: a.log, Cache: a.exports,
	})
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

// deviceToken reads the stored token, empty when unpaired. The
// answer is remembered: the status page asks every two seconds and
// every tick asks again, while on Linux the keychain is a D-Bus round
// trip. Pair and Unpair are the only things that change it, and both
// forget it. A read that failed is not remembered, so a locked
// keychain that is unlocked a moment later works on the next ask.
func (a *App) deviceToken() string {
	a.tokMu.Lock()
	defer a.tokMu.Unlock()
	if a.tokKnown {
		return a.tok
	}
	tok, err := a.o.Secret.Token()
	if err != nil {
		a.log.Warn("could not read the device token", "component", "app", "err", err.Error())
		return ""
	}
	a.tok, a.tokKnown = tok, true
	return tok
}

// forgetToken drops the cached token, so the next read goes to the
// store again.
func (a *App) forgetToken() {
	a.tokMu.Lock()
	defer a.tokMu.Unlock()
	a.tok, a.tokKnown = "", false
}

// prune deletes the state of reports that finished longer ago than
// ReportRetention. It runs once, at startup: the files are small and
// the cost is in their number, which grows by a handful a week.
func (a *App) prune(now time.Time) {
	reports, err := state.List(a.o.Dirs.State)
	if err != nil {
		a.log.Warn("could not list the reports to prune",
			"component", "app", "err", err.Error())
		return
	}
	cutoff := now.Add(-ReportRetention)
	removed := 0
	for _, r := range reports {
		if !r.Done || !r.StartedAt.Before(cutoff) {
			continue
		}
		if err := state.Remove(a.o.Dirs.State, r.Key); err != nil {
			a.log.Warn("could not remove an old report's state", "component", "app",
				"report", r.Key, "err", err.Error())
			continue
		}
		removed++
	}
	if removed > 0 {
		a.log.Info("forgot reports older than the retention", "component", "app",
			"reports", removed, "retention", ReportRetention.String())
	}
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
// which the watcher handles as "nothing to do". It reads a.installs,
// so every caller but New holds a.mu.
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
	// The config-file backend reads config.json, edits it and writes
	// it back, which is the same read-modify-write SaveSettings does:
	// both happen under a.mu so neither can lose the other's change.
	// The network call above stays outside it.
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := a.o.Secret.SetToken(dev.Token); err != nil {
		return err
	}
	a.forgetToken()
	a.log.Info("this device is paired", "component", "app", "device_id", dev.DeviceID)
	return nil
}

// Unpair forgets the device token. The reports already uploaded stay
// on the site; revoking the device itself is done on the account page.
func (a *App) Unpair() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	err := a.o.Secret.Clear()
	a.forgetToken()
	return err
}

// Settings is what the settings page can change.
type Settings struct {
	Visibility       string               `json:"visibility"`
	WoWPaths         []string             `json:"wow_paths"`
	LoggingCharacter *character.Character `json:"logging_character"`
}

// SaveSettings validates and stores the settings, then puts them to
// work without a restart: the visibility and the logging character
// are handed to the pipeline, which reads them when the next report
// opens, and a newly added game folder re-detects the installs and
// moves the tail — as soon as the report in progress closes, because
// a watcher carries a file and an offset and cannot be moved out from
// under an open report.
func (a *App) SaveSettings(in Settings) error {
	if err := a.store(in); err != nil {
		return err
	}
	a.retarget()
	return nil
}

// store validates the settings and writes them, holding the lock for
// the whole read-modify-write of config.json.
func (a *App) store(in Settings) error {
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
	// The device token lives in config.json on the fallback backend
	// and is written there by the secret store alone. Whatever is on
	// disk now is the truth; carrying the in-memory blank over it
	// would silently unpair the player.
	onDisk, err := config.Load(a.o.Dirs.ConfigFile())
	if err != nil {
		return err
	}
	cfg.DeviceToken = onDisk.DeviceToken
	if err := config.Save(a.o.Dirs.ConfigFile(), cfg); err != nil {
		return err
	}
	cfg.DeviceToken = ""
	a.cfg = cfg
	a.installs = a.detect()
	a.pipe.SetConfig(cfg)
	return nil
}

// retarget points the tail at the first install's Logs directory when
// the settings have moved it. The swap waits for the open report to
// close, so it is attempted again on every step rather than only on
// the save that asked for it.
func (a *App) retarget() {
	a.mu.Lock()
	want := a.logsDir()
	same := want == a.watching
	a.mu.Unlock()
	if same {
		return
	}
	if !a.pipe.SetWatch(watch.New(watch.Options{Dir: want})) {
		return // a report is open; the next step tries again
	}
	a.mu.Lock()
	a.watching = want
	a.mu.Unlock()
	a.log.Info("the tail moved to a new game folder", "component", "app", "dir", want)
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
		// Through the cache: this runs on a two-second poll and the
		// file behind it is megabytes of Lua.
		found, _, err := a.exports.Exports(p)
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
	a.retarget()
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

// UI is the local server the window talks to.
type UI struct {
	// URL is what the webview opens: the loopback address with the
	// session token in the path.
	URL string
	// Addr is the address without the token, which is what may be
	// logged. The token is the whole authentication for /api/pair,
	// /api/unpair and /api/settings, and companion.log is a file the
	// README tells players to open and paste.
	Addr string

	srv *http.Server
}

// Close stops the server and releases the port.
func (u *UI) Close() error {
	if err := u.srv.Close(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Serve starts the local UI server on a loopback port. The token is
// in the path, so a page on another origin cannot guess it.
func (a *App) Serve() (*UI, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("listen on loopback: %w", err)
	}
	u := &UI{
		Addr: ln.Addr().String(),
		srv:  &http.Server{Handler: a.Handler(), ReadHeaderTimeout: 5 * time.Second},
	}
	u.URL = fmt.Sprintf("http://%s/%s/", u.Addr, a.token)
	go func() {
		// Close above is the only way this server stops, so anything
		// but ErrServerClosed is a real failure worth an Error line;
		// an ordinary quit produces none.
		if err := u.srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Error("the local UI server stopped", "component", "ui", "err", err.Error())
		}
	}()
	return u, nil
}
