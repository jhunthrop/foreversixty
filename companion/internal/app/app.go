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
	"maps"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/addonupdater"
	"github.com/jhunthrop/foreversixty/companion/internal/browser"
	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/logger"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/picker"
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
	// Addons are the addon updater's options (internal/addonupdater). The
	// key, the API base and the HTTP client default to the Updater's, and a
	// test may set Processes to say which games are running.
	Addons addonupdater.Options
	// Retry is the upload retry policy. Zero means the client's
	// default; tests shrink it so an outage costs milliseconds.
	Retry client.Retry
	// Open shows a link in the system browser. Zero means the real
	// thing (package browser); tests record instead of launching one.
	Open func(url string) error
	// PickFolder is the native folder dialog. Zero means the real one
	// (package picker); tests answer without a dialog.
	PickFolder func() (string, error)
	// ConnectInterval overrides how often a sign-in is polled. Zero
	// means the server's own interval, two seconds.
	ConnectInterval time.Duration
	// HeartbeatEvery overrides how often an open report tells the site
	// the companion is alive. Zero means once a minute, the API's limit;
	// tests shrink it.
	HeartbeatEvery time.Duration
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

	mu  sync.Mutex
	cfg config.Config
	// signedOut is set when the site refused the device token, until the
	// next sign-in succeeds; guarded by mu.
	signedOut *SignedOut
	// failSaid is the last failure each per-tick step logged, so a step that fails
	// every second says so once and not every second (see saidFailure).
	failMu   sync.Mutex
	failSaid map[string]string
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
	// addons keeps the addon and the data addon current in every install;
	// addonsBusy is set while its goroutine is running.
	addons     *addonupdater.Updater
	addonsBusy atomic.Bool

	// bg is the context the sign-in poll runs under; Close ends it and
	// waits for the poll, which must not outlive the app.
	bg   context.Context
	stop context.CancelFunc
	bgWG sync.WaitGroup

	// conn is the sign-in in progress (connect.go), guarded by mu.
	conn connectState
	// added is the install paths that came from a folder the player
	// added, as opposed to one found by looking; guarded by mu.
	added map[string]bool
	// noticeShown is set once FirstClose has been asked in this process.
	noticeShown bool
	// updatedTo is the version when this process is the first launch of
	// it after an update; empty otherwise (version.go).
	updatedTo string
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
	if o.Open == nil {
		o.Open = browser.Open
	}
	if o.PickFolder == nil {
		o.PickFolder = picker.Pick
	}
	if o.Secret == nil {
		o.Secret = secret.ConfigFile{Path: o.Dirs.ConfigFile()}
	}
	tok, err := localToken()
	if err != nil {
		return nil, err
	}
	a := &App{o: o, log: o.Log, now: o.Now, token: tok, cfg: o.Config}
	a.bg, a.stop = context.WithCancel(context.Background())
	// The secret store owns the device token. The in-memory copy is
	// blanked so no later write of this configuration can carry a
	// stale token back over the one on disk; SaveSettings reads the
	// live value out of config.json when it saves.
	a.cfg.DeviceToken = ""
	a.dropOtherGames()
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
		Records:  a.addonRecords,

		HeartbeatEvery: o.HeartbeatEvery,
	})
	if err != nil {
		return nil, err
	}
	a.sync = addon.New(addon.Options{
		Paths: a.savedVariables, API: a.client, Log: a.log, Cache: a.exports,
		StateDir: o.Dirs.State, Now: a.now,
	})
	uo := o.Updater
	uo.Dir, uo.Log = o.Dirs.Update, a.log
	a.updater = updater.New(uo)
	ao := o.Addons
	ao.Dir, ao.Log = filepath.Join(o.Dirs.Update, "addons"), a.log
	if ao.PublicKey == "" {
		ao.PublicKey = uo.PublicKey
		if ao.PublicKey == "" {
			ao.PublicKey = updater.PublicKey
		}
	}
	if ao.APIBase == "" {
		ao.APIBase = uo.APIBase
	}
	if ao.HTTP == nil {
		ao.HTTP = uo.HTTP
	}
	if ao.Now == nil {
		ao.Now = a.now
	}
	ao.Installs, ao.Enabled = a.installPaths, a.addonUpdatesOn
	a.addons = addonupdater.New(ao)
	a.noteVersion()
	return a, nil
}

// Close releases what the app holds.
func (a *App) Close() error {
	a.stop()
	a.bgWG.Wait()
	return a.pipe.Close()
}

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

// dropOtherGames takes out of the stored folders any that is not a Forever
// variant (an old config that remembered _retail_), with one log line each.
// It runs at start, before the installs are found. Callers hold a.mu or are New.
func (a *App) dropOtherGames() {
	keep := make([]string, 0, len(a.cfg.WoWPaths))
	for _, p := range a.cfg.WoWPaths {
		var r *wow.Refusal
		if _, err := wow.Pick(p); errors.As(err, &r) {
			a.log.Info("dropped a stored game folder that is not a Forever variant",
				"component", "app", "path", p, "folder", r.Folder)
			continue
		}
		keep = append(keep, p)
	}
	a.cfg.WoWPaths = keep
}

// detect lists the installs: the configured paths first, then
// anything found in the usual places, and records which are which
// (a.added). Callers hold a.mu, except New.
func (a *App) detect() []wow.Install {
	var out []wow.Install
	seen := map[string]bool{}
	added := map[string]bool{}
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
			added[in.Path] = true
		}
	}
	for _, in := range wow.Detect() {
		add(in)
	}
	a.added = added
	return out
}

// switchLead is how much newer another install's activity must be
// before the watch leaves the one it is on, so two installs touched
// within a moment of each other never make the tail flap.
const switchLead = time.Second

// logsDir is the Logs directory the watcher follows: the install played
// most recently (the newest file under its Logs), the first install on
// a tie or when nothing has been played, or a directory that does not
// exist when there is none, which the watcher handles as "nothing to
// do". An install already being watched keeps the watch unless another
// is newer by at least switchLead. It reads a.installs and a.watching,
// so every caller but New holds a.mu.
func (a *App) logsDir() string {
	if len(a.installs) == 0 {
		return a.o.Dirs.Home
	}
	best := a.installs[0]
	bestAt, bestOK := wow.LastActivity(best)
	var cur string
	var curAt time.Time
	var curOK bool
	for _, in := range a.installs {
		at, ok := wow.LastActivity(in)
		if in.Logs == a.watching {
			cur, curAt, curOK = in.Logs, at, ok
		}
		if ok && (!bestOK || at.After(bestAt)) {
			best, bestAt, bestOK = in, at, ok
		}
	}
	if cur != "" && best.Logs != cur {
		if bestOK && curOK && bestAt.Sub(curAt) < switchLead {
			return cur
		}
	}
	return best.Logs
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

// addonRecords is every character the addon has exported, from every
// install's account files, for working out which of them a combat log
// belongs to. The watched install's own records win a tie. It reads
// through the cache, so a log that is open for hours costs one parse
// per logout.
func (a *App) addonRecords() []logger.Record {
	a.mu.Lock()
	installs := slices.Clone(a.installs)
	watching := a.watching
	a.mu.Unlock()
	var out []logger.Record
	for _, in := range installs {
		paths, err := in.SavedVariables(addon.SavedVariablesName)
		if err != nil {
			continue
		}
		for _, p := range paths {
			found, _, err := a.exports.Exports(p)
			if err != nil {
				continue
			}
			out = append(out, logger.Records(found, in.Logs == watching)...)
		}
	}
	return out
}

// deviceName is what the site lists this device as.
func deviceName() string {
	name, err := os.Hostname()
	if err != nil || name == "" {
		return "Forever Sixty companion"
	}
	return name
}

// platform is the device's platform as the API wants it.
func platform() string { return runtime.GOOS + "/" + runtime.GOARCH }

// Pair exchanges a pairing code for a device token and stores it. This
// is the fallback to Connect, for a machine that cannot open a browser.
// A failure is a *Fault: the page shows the sentence for its code, and
// the cause goes to the log.
func (a *App) Pair(ctx context.Context, code string) error {
	code = strings.TrimSpace(code)
	if code == "" {
		return fault(CodeEmptyCode, nil)
	}
	dev, err := a.client.Claim(ctx, client.Claim{
		Code: code, Name: deviceName(), Platform: platform(),
	})
	if err != nil {
		a.log.Warn("the pairing code was not accepted", "component", "app", "err", err.Error())
		return fault(netCode(err, CodeBadCode), err)
	}
	if err := a.saveDevice(dev.Token, dev.AccountLabel); err != nil {
		return err
	}
	a.cancelConnect()
	a.log.Info("this device is paired", "component", "app", "device_id", dev.DeviceID)
	return nil
}

// saveDevice stores a device token and who it belongs to. The
// config-file backend reads config.json, edits it and writes it back,
// which is the same read-modify-write SaveSettings does: both happen
// under a.mu so neither can lose the other's change. The network call
// that produced the token stays outside it.
func (a *App) saveDevice(token, label string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := a.o.Secret.SetToken(token); err != nil {
		a.log.Error("could not save the device token", "component", "app", "err", err.Error())
		return fault(CodeSave, err)
	}
	a.forgetToken()
	a.signedOut = nil
	// A heartbeat that stopped on the old token starts again for the
	// report still open.
	a.pipe.ResumeHeartbeat()
	if err := a.writeLabel(label); err != nil {
		// The token is saved and the device works; only the "connected
		// as" line is lost.
		a.log.Warn("could not save the account label", "component", "app", "err", err.Error())
	}
	return nil
}

// writeLabel stores the account label in config.json beside whatever
// else is there, and in memory. Callers hold a.mu.
func (a *App) writeLabel(label string) error {
	a.cfg.AccountLabel = label
	onDisk, err := config.Load(a.o.Dirs.ConfigFile())
	if err != nil {
		return err
	}
	onDisk.AccountLabel = label
	return config.Save(a.o.Dirs.ConfigFile(), onDisk)
}

// RevokeTimeout bounds how long Forget waits for the site before it clears
// the token here anyway.
const RevokeTimeout = 10 * time.Second

// Unpair forgets this device. It first asks the site to revoke the device
// (DELETE /v1/devices/self), so the account's device list no longer shows
// it, then clears the token here whether or not that worked. revoked
// reports whether the site confirmed; a token the site already refuses
// counts as confirmed. The reports already uploaded stay on the site.
func (a *App) Unpair(ctx context.Context) (revoked bool, err error) {
	revoked = true
	if a.deviceToken() != "" {
		rctx, cancel := context.WithTimeout(ctx, RevokeTimeout)
		rerr := a.client.RevokeSelf(rctx)
		cancel()
		if rerr != nil {
			revoked = false
			a.log.Warn("could not revoke the device on the site", "component", "app", "err", rerr.Error())
		}
	}
	return revoked, a.clearLocal()
}

// clearLocal drops the token and the account label from this PC.
func (a *App) clearLocal() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	err := a.o.Secret.Clear()
	a.forgetToken()
	a.signedOut = nil
	if err != nil {
		a.log.Error("could not forget the device token", "component", "app", "err", err.Error())
		return fault(CodeSave, err)
	}
	if err := a.writeLabel(""); err != nil {
		a.log.Warn("could not clear the account label", "component", "app", "err", err.Error())
	}
	return nil
}

// Settings is what the settings page can change.
type Settings struct {
	Visibility       string               `json:"visibility"`
	WoWPaths         []string             `json:"wow_paths,omitempty"`
	LoggingCharacter *character.Character `json:"logging_character"`
	// KeepAddonUpdated is the Keep the addon up to date switch; absent
	// leaves it as it is.
	KeepAddonUpdated *bool `json:"keep_addon_updated,omitempty"`
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
// the whole read-modify-write of config.json. A failure is a *Fault.
func (a *App) store(in Settings) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	cfg := a.cfg
	cfg.ReportVisibility = in.Visibility
	if !slices.Contains(config.Visibilities, in.Visibility) {
		return fault(CodeBadRequest, fmt.Errorf("report visibility %q is not one of %v", in.Visibility, config.Visibilities))
	}
	// The settings page no longer sends folders: they are added and
	// removed one at a time. A nil list leaves them as they are.
	if in.WoWPaths != nil {
		cfg.WoWPaths = in.WoWPaths
	}
	cfg.LoggingCharacter = in.LoggingCharacter
	if in.KeepAddonUpdated != nil {
		keep := *in.KeepAddonUpdated
		cfg.KeepAddonUpdated = &keep
	}
	if cfg.LoggingCharacter != nil && !cfg.LoggingCharacter.Valid() {
		return fault(CodeNoCharacter, errors.New("the logging character needs a region, a ruleset and a name"))
	}
	for _, p := range cfg.WoWPaths {
		if _, err := wow.Pick(p); err != nil {
			return fault(CodeNotWoW, err)
		}
	}
	return a.persist(cfg)
}

// persist writes a configuration and puts it to work: the installs are
// found again and the pipeline is told. Callers hold a.mu.
func (a *App) persist(cfg config.Config) error {
	if cfg.WoWPaths == nil {
		cfg.WoWPaths = []string{}
	}
	if err := cfg.Validate(); err != nil {
		return fault(CodeSettings, err)
	}
	// The device token lives in config.json on the fallback backend
	// and is written there by the secret store alone. Whatever is on
	// disk now is the truth; carrying the in-memory blank over it
	// would silently unpair the player.
	onDisk, err := config.Load(a.o.Dirs.ConfigFile())
	if err != nil {
		return fault(CodeSettings, err)
	}
	cfg.DeviceToken = onDisk.DeviceToken
	if err := config.Save(a.o.Dirs.ConfigFile(), cfg); err != nil {
		return fault(CodeSettings, err)
	}
	cfg.DeviceToken = ""
	a.cfg = cfg
	a.installs = a.detect()
	a.pipe.SetConfig(cfg)
	return nil
}

// retarget points the tail at the most recently played install's Logs
// directory when the settings or the player have moved it. The swap waits for the open report to
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
	// AppVersion is the companion's own version, for the footer.
	// EngineVersion is the log engine's, kept for support and not shown.
	AppVersion    string `json:"app_version"`
	EngineVersion string `json:"engine_version"`
	Paired        bool   `json:"paired"`
	// AccountLabel is who the device is connected as, empty when it
	// is not known.
	AccountLabel string `json:"account_label"`
	// DeviceName is what the site lists this device as.
	DeviceName string `json:"device_name"`
	// SignedOut is present from the moment the site refuses this device's
	// token until the next sign-in succeeds. AccountLabel stays empty then;
	// the label the sign-in had is in here.
	SignedOut *SignedOut `json:"signed_out,omitempty"`
	// Headline is the Status page's opening line (headline.go).
	Headline Headline `json:"headline"`
	// Connect is the Battle.net sign-in in progress (connect.go).
	Connect          ConnectView `json:"connect"`
	TokenBackend     string      `json:"token_backend"`
	TokenBackendName string      `json:"token_backend_name"`
	// TokenBackendNote is the muted line under the store: what Credential Manager is,
	// or, on the file fallback, which store was not available.
	TokenBackendNote string `json:"token_backend_note,omitempty"`
	APIBaseURL       string `json:"api_base_url"`
	SiteBaseURL      string `json:"site_base_url"`
	Visibility       string `json:"visibility"`
	// LoggingCharacter is the character picked in Settings to override
	// the log; nil is Automatic. DetectedCharacter is who the combat
	// log says is logging (the last one seen when no report is open),
	// empty before any log has named one.
	LoggingCharacter  *character.Character `json:"logging_character,omitempty"`
	DetectedCharacter string               `json:"detected_character"`
	Installs          []InstallStatus      `json:"installs"`
	// WatchPending is true while the tail has not yet moved to the
	// first install, because a report is open.
	WatchPending bool            `json:"watch_pending"`
	Pipeline     pipeline.Status `json:"pipeline"`
	Reports      []ReportRow     `json:"reports"`
	Characters   []addon.Export  `json:"characters"`
	UpdateStaged bool            `json:"update_staged"`
	// UpdatedTo is the version this process is the first launch of, empty
	// on every launch after it. The footer says so for half a minute.
	UpdatedTo string `json:"updated_to"`
	// AddonSeenAt is when the site last saw an export (the newest
	// ack's export_seen_at across characters), in local time; absent
	// until the site has acked one.
	AddonSeenAt *time.Time `json:"addon_seen_at,omitempty"`
	// SimRequestState is what became of the newest in-game sim
	// request, in the ack's words (addon.SimQueued and the rest).
	SimRequestState string `json:"sim_request_state,omitempty"`
	// AddonLine is the two above in the player's words (AddonLine in
	// copy.go); before any ack it says the addon has not synced.
	AddonLine string `json:"addon_line"`
	// AddonNever is true before the site has acked any export; the row
	// is then muted rather than a fault.
	AddonNever bool `json:"addon_never"`
	// AddonStale is true when the last sync is over a day old, and
	// AddonHelp then says what to do; before any sync it holds the
	// install instruction instead.
	AddonStale bool   `json:"addon_stale"`
	AddonHelp  string `json:"addon_help,omitempty"`
	// AddonSeenTitle is the absolute time, for a tooltip.
	AddonSeenTitle string `json:"addon_seen_title,omitempty"`
	// LastSync is the Last sync row: what the newest character list
	// put on the site, or the site's reason for refusing it.
	LastSync LastSyncView `json:"last_sync"`
	// Credit is who the current report is credited to; none when no
	// report is open.
	Credit Credit `json:"credit"`
	// BuildsChecked is what the last Check for builds now found; absent
	// until the player has pressed it.
	BuildsChecked *BuildsChecked `json:"builds_checked,omitempty"`
	// BuildsWaiting is the Addon row's line while a sent build has not
	// yet shown in the game; absent otherwise.
	BuildsWaiting string `json:"builds_waiting,omitempty"`
	// KeepAddonUpdated is the Keep the addon up to date switch, and
	// AddonUpdate the Addon row's version lines and install buttons.
	KeepAddonUpdated bool            `json:"keep_addon_updated"`
	AddonUpdate      AddonUpdateView `json:"addon_update"`
}

// InstallStatus is one game folder as the status page shows it.
type InstallStatus struct {
	Path string `json:"path"`
	// Flavor is the folder name, "_classic_era_" and the like;
	// FlavorName is the same in the player's words ("Classic Era").
	Flavor     string `json:"flavor"`
	FlavorName string `json:"flavor_name"`
	// Watching is the folder the pipeline is tailing, and Added is
	// true for a folder the player added (as opposed to one found).
	Watching bool `json:"watching"`
	Added    bool `json:"added"`
	// LastPlayed is when the game last wrote under this install's
	// Logs; nil when it never has.
	LastPlayed *time.Time `json:"last_played,omitempty"`
	// Advanced is true, false, or unknown when Config.wtf could not
	// be read — which is the normal state before the first login.
	// AdvancedLogHelp is the instruction line for a switch that is off,
	// and only the watched install carries it: an install the player is
	// not logging is not a fault, so the window shows it "off" in the
	// muted token with no instruction. It moves to another install the
	// moment the player watches that one.
	Advanced        bool   `json:"advanced_logging"`
	AdvancedKnown   bool   `json:"advanced_logging_known"`
	AdvancedLogHelp string `json:"advanced_logging_help,omitempty"`
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
	// Credit is who the report is credited to.
	Credit Credit `json:"credit"`
}

// installStatuses is the installs as the window shows them.
func (a *App) installStatuses() (rows []InstallStatus, watchPending bool) {
	a.mu.Lock()
	installs := slices.Clone(a.installs)
	added := maps.Clone(a.added)
	watching := a.watching
	want := a.logsDir()
	a.mu.Unlock()

	rows = make([]InstallStatus, 0, len(installs))
	for _, in := range installs {
		row := InstallStatus{Path: in.Path, Flavor: in.Flavor, FlavorName: wow.FlavorName(in.Flavor),
			Watching: in.Logs == watching, Added: added[in.Path]}
		if row.Watching {
			row.AdvancedLogHelp = CopyAdvancedRowHelp
		}
		if at, ok := wow.LastActivity(in); ok {
			row.LastPlayed = &at
		}
		if on, err := in.AdvancedLogging(); err == nil {
			row.Advanced, row.AdvancedKnown = on, true
		}
		rows = append(rows, row)
	}
	return rows, want != watching
}

// Headline is the Status page's opening line right now. The tray asks
// for it on a timer, so it reads only what it needs.
func (a *App) Headline() Headline {
	installs, _ := a.installStatuses()
	connecting := a.connectView().State == ConnectWaiting || a.connectView().State == ConnectBlocked
	st := a.pipe.Status()
	a.mu.Lock()
	signedOut := a.signedOut != nil
	a.mu.Unlock()
	return headlineFor(a.paired(signedOut), connecting, signedOut, st.Logging, a.currentCredit(st.ReportKey), installs)
}

// ReportsURL is the player's reports on the site.
func (a *App) ReportsURL() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.cfg.SiteBaseURL + "/logs"
}

// Snapshot builds the UI's view of the world.
func (a *App) Snapshot() Snapshot {
	a.mu.Lock()
	cfg := a.cfg
	a.mu.Unlock()

	a.mu.Lock()
	so := a.signedOut
	a.mu.Unlock()
	paired := a.paired(so != nil)
	s := Snapshot{
		AppVersion:       updater.Version,
		UpdatedTo:        a.updatedTo,
		EngineVersion:    session.Version,
		Paired:           paired,
		SignedOut:        so,
		AccountLabel:     cfg.AccountLabel,
		DeviceName:       deviceName(),
		Connect:          a.connectView(),
		TokenBackend:     string(a.o.Secret.Backend()),
		TokenBackendName: secret.Name(a.o.Secret.Backend(), runtime.GOOS),
		TokenBackendNote: tokenNote(a.o.Secret.Backend()),
		APIBaseURL:       cfg.APIBaseURL,
		SiteBaseURL:      cfg.SiteBaseURL,
		Visibility:       cfg.ReportVisibility,
		LoggingCharacter: cfg.LoggingCharacter,
		KeepAddonUpdated: cfg.AddonUpdates(),
		Pipeline:         a.pipe.Status(),
		Reports:          []ReportRow{},
		Characters:       []addon.Export{},
		Credit:           Credit{State: CreditNone},
	}
	if !paired {
		s.AccountLabel = ""
	}
	s.Installs, s.WatchPending = a.installStatuses()
	connecting := s.Connect.State == ConnectWaiting || s.Connect.State == ConnectBlocked
	s.DetectedCharacter = s.Pipeline.Detected
	reports, err := state.List(a.o.Dirs.State)
	if err != nil {
		a.log.Warn("could not list the reports", "component", "app", "err", err.Error())
	}
	// One report is one row: a state file read twice (a copy caught mid-save,
	// or one report under two names) keeps the newest copy, the one that has
	// the server's id.
	seen := map[string]bool{}
	for i := len(reports) - 1; i >= 0; i-- {
		r := reports[i]
		if seen[r.Key] {
			continue
		}
		seen[r.Key] = true
		row := ReportRow{Key: r.Key, ReportID: r.ReportID, Zone: r.Zone,
			Fights: r.Fights, StartedAt: r.StartedAt, Done: r.Done, Credit: creditOf(r)}
		if r.Key == s.Pipeline.ReportKey {
			s.Credit = row.Credit
		}
		if r.ReportID != "" {
			row.URL = cfg.SiteBaseURL + "/reports/" + r.ReportID
		}
		s.Reports = append(s.Reports, row)
	}
	s.Headline = headlineFor(paired, connecting, so != nil, s.Pipeline.Logging, s.Credit, s.Installs)
	for _, p := range a.savedVariables() {
		// Through the cache: this runs on a two-second poll and the
		// file behind it is megabytes of Lua. The stamp is the sync's
		// business, not the status page's — reading here as often as
		// the window likes costs the sync nothing.
		found, _, err := a.exports.Exports(p)
		if err != nil {
			continue
		}
		s.Characters = append(s.Characters, found...)
	}
	if _, err := os.Stat(updater.PendingPath(a.o.Dirs.Update)); err == nil {
		s.UpdateStaged = true
	}
	now := a.now()
	st := a.sync.Status()
	if c := st.Checked; !c.At.IsZero() {
		s.BuildsChecked = &BuildsChecked{At: c.At.In(now.Location()), New: c.New, For: c.For}
	}
	if st.Waiting > 0 {
		s.BuildsWaiting = CopyBuildsWaiting(st.Waiting, st.WaitingFor)
	}
	s.AddonLine = AddonLine(st.SeenAt, st.SimRequestState, now)
	s.AddonNever = st.SeenAt.IsZero()
	s.AddonUpdate = a.addonUpdateView()
	if s.AddonNever {
		s.AddonHelp = neverSyncedHelp(s.AddonUpdate)
	}
	s.LastSync = a.lastSyncView(st, now)
	if !st.SeenAt.IsZero() {
		local := st.SeenAt.In(now.Location())
		s.AddonSeenAt, s.SimRequestState = &local, st.SimRequestState
		s.AddonSeenTitle = local.Format("Mon 2 Jan 2006, 15:04")
		if s.AddonStale = AddonStale(st.SeenAt, now); s.AddonStale {
			s.AddonHelp = CopyStaleHelp
		}
	}
	return s
}

// Step is one pass of everything on a timer: tail, snapshot, upload,
// addon sync and the update check. It returns no error because the
// companion must keep logging through every one of them failing.
func (a *App) Step(ctx context.Context, now time.Time) {
	tok := a.deviceToken()
	a.retarget()
	// The addon updater needs no sign-in, so it ticks before anything that does.
	a.tickAddons(now)
	if err := a.pipe.Tick(now); a.saidFailure("tail", err) {
		a.log.Error("the tail failed", "component", "app", "err", err.Error())
	}
	// A heartbeat the site refused has stopped itself; the signed-out
	// path decides what that means, exactly as for a refused snapshot.
	if err := a.pipe.HeartbeatRefused(); err != nil && a.signOutIfRefused(err, tok) {
		return
	}
	if err := a.pipe.Live(ctx, now); err != nil {
		if a.signOutIfRefused(err, tok) {
			return
		}
		if client.Unauthorized(err) {
			return // already signed out; the pipeline logged the attempt
		}
		if a.saidFailure("live", err) {
			a.log.Error("a live snapshot failed", "component", "app", "err", err.Error())
		}
	} else {
		a.saidFailure("live", nil)
	}
	a.mu.Lock()
	signedOut := a.signedOut != nil
	a.mu.Unlock()
	if signedOut || a.deviceToken() == "" {
		return // nothing below here works unpaired
	}
	if err := a.pipe.Drain(ctx); err != nil && !errors.Is(err, context.Canceled) {
		if a.signOutIfRefused(err, tok) {
			return
		}
		if a.saidFailure("drain", err) {
			a.log.Warn("uploads are behind", "component", "app", "err", err.Error())
		}
	} else {
		a.saidFailure("drain", nil)
	}
	if err := a.sync.Poll(ctx, now); err != nil {
		if a.signOutIfRefused(err, tok) {
			return
		}
		if a.saidFailure("addonsync", err) {
			a.log.Warn("the addon sync did not complete", "component", "app", "err", err.Error())
		}
	} else {
		a.saidFailure("addonsync", nil)
	}
	if staged, err := a.updater.Poll(ctx, now); err != nil {
		if a.saidFailure("updatecheck", err) {
			a.log.Warn("the update check failed", "component", "app", "err", err.Error())
		}
	} else if a.saidFailure("updatecheck", nil); staged {
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

// SignedOutRevoked is why a device is signed out: the site stopped
// accepting its token, because the device was removed from the account or
// the account was reset.
const SignedOutRevoked = "revoked"

// SignedOut says the site no longer accepts this device's sign-in.
type SignedOut struct {
	Reason string `json:"reason"`
	// AccountLabel is who the device was connected as, for the message.
	AccountLabel string    `json:"account_label"`
	At           time.Time `json:"at"`
}

// paired is true when a device token is held and the site has not
// refused it.
func (a *App) paired(signedOut bool) bool {
	return !signedOut && a.deviceToken() != ""
}

// signOutIfRefused marks the device signed out when err is the site
// refusing the token tok, which was the one in use when the call began.
// A token is either valid or revoked, so the first refusal is final:
// the token is forgotten (the uploads stay queued on disk, to resume
// after the next sign-in) and the cause is logged once. It reports
// whether err was such a refusal.
func (a *App) signOutIfRefused(err error, tok string) bool {
	if err == nil || tok == "" || !client.Unauthorized(err) {
		return false
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.signedOut != nil || a.deviceToken() != tok {
		// Already said, or a new sign-in replaced the token this call
		// used; the refusal belongs to the old one.
		return true
	}
	if !a.connectedBefore() {
		// A token left in the system store by an earlier install, refused on
		// the first call this home ever makes: nothing was lost here, so the
		// card reads as a first run and not as a sign-in that lapsed.
		if cerr := a.o.Secret.Clear(); cerr != nil {
			a.log.Error("could not forget the refused device token", "component", "app", "err", cerr.Error())
		}
		a.forgetToken()
		a.log.Info("a device token left by an earlier install was refused; starting as a first run",
			"component", "app", "err", err.Error())
		return true
	}
	a.signedOut = &SignedOut{Reason: SignedOutRevoked, AccountLabel: a.cfg.AccountLabel, At: a.now()}
	if cerr := a.o.Secret.Clear(); cerr != nil {
		a.log.Error("could not forget the refused device token", "component", "app", "err", cerr.Error())
	}
	a.forgetToken()
	a.log.Warn("the site no longer accepts this device's sign-in; sign in again",
		"component", "app", "err", err.Error())
	return true
}

// tokenNote is the muted line under "Saved in" on the Device page.
func tokenNote(b secret.Backend) string {
	if n := secret.FallbackNote(b, runtime.GOOS); n != "" {
		return n
	}
	return secret.StoreNote(b, runtime.GOOS)
}

// connectedBefore is true when this home has been signed in: it holds who
// the device was connected as, or the state of a report it uploaded. A home
// with neither is a first run, whatever the system store still holds.
// Callers hold a.mu.
func (a *App) connectedBefore() bool {
	if a.cfg.AccountLabel != "" {
		return true
	}
	reports, err := state.List(a.o.Dirs.State)
	return err == nil && len(reports) > 0
}

// saidFailure reports whether a step's failure is news: true the first time
// err is seen for key and whenever its text changes, false while the same
// failure repeats. A nil err (the step worked) forgets the failure, so the
// next one is said again. Step runs every second, and a step that fails
// every second must reach the log once, not once a second.
func (a *App) saidFailure(key string, err error) bool {
	a.failMu.Lock()
	defer a.failMu.Unlock()
	if err == nil {
		delete(a.failSaid, key)
		return false
	}
	if a.failSaid == nil {
		a.failSaid = map[string]string{}
	}
	if a.failSaid[key] == err.Error() {
		return false
	}
	a.failSaid[key] = err.Error()
	return true
}
