// companion/internal/addonupdater/updater.go
package addonupdater

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
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

// Every is how often the releases are listed (the companion's own
// six-hour check), and ApplyEvery how often a release that is waiting for
// the game to close is tried again.
const (
	Every      = updater.Every
	ApplyEvery = time.Minute
)

// maxPages bounds how far down the release list a check reads. The
// repository publishes a data release every night, so an addon release can
// sit a few months down; ten pages of thirty is most of a year.
const maxPages = 10

// retryAfter is how long an install that failed is left alone, so a folder
// the system will not let go of is not copied again every minute.
const retryAfter = time.Hour

// The reasons an install can be refused, for the window to word.
var (
	// ErrNoKey says the build carries no public key, so nothing can be verified.
	ErrNoKey = errors.New("addonupdater: this build has no minisign public key, so it installs no addon")
	// ErrGameRunning says the install's game is running.
	ErrGameRunning = errors.New("addonupdater: the game is running")
	// ErrDownload says the release could not be fetched or did not verify.
	ErrDownload = errors.New("addonupdater: the release could not be downloaded or verified")
	// ErrNoRelease says no signed release of the addon was found.
	ErrNoRelease = errors.New("addonupdater: no signed addon release was found")
)

// Options configures an Updater.
type Options struct {
	// Dir is where downloads wait, in the companion's home.
	Dir string
	// APIBase is https://api.github.com unless a test says otherwise.
	APIBase string
	// Repo is the repository whose releases are read; the companion's own
	// unless a test says otherwise.
	Repo string
	HTTP *http.Client
	Log  *slog.Logger
	// PublicKey is the minisign key a zip must verify against. Empty
	// disables the updater: it installs nothing it cannot verify.
	PublicKey string
	// Installs lists the game installs (flavour folders) to keep current.
	Installs func() []string
	// Enabled says whether automatic checks and installs are on. The
	// Install button works either way.
	Enabled func() bool
	// Processes lists what is running; the default is the system's list.
	Processes func() ([]Proc, error)
	Now       func() time.Time
	// Every and ApplyEvery override the defaults.
	Every, ApplyEvery time.Duration
}

// Updater keeps the addon and the data addon current.
type Updater struct {
	o Options

	// run serialises everything that downloads or writes.
	run sync.Mutex

	mu       sync.Mutex // guards the fields below
	lastPoll time.Time
	lastLook time.Time
	backoff  map[string]time.Time
	rejected map[string]bool
	// waitSaid holds the installs whose "waiting for the game to close" line
	// is already in the log, so the once-a-minute look says it once.
	waitSaid map[string]bool
	// listFailed is true while the process list cannot be read and the log
	// has said so.
	listFailed bool
	recs       records
}

// New builds an updater.
func New(o Options) *Updater {
	if o.APIBase == "" {
		o.APIBase = "https://api.github.com"
	}
	if o.Repo == "" {
		o.Repo = updater.Repo
	}
	if o.HTTP == nil {
		o.HTTP = &http.Client{Timeout: 5 * time.Minute}
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	if o.Installs == nil {
		o.Installs = func() []string { return nil }
	}
	if o.Enabled == nil {
		o.Enabled = func() bool { return true }
	}
	if o.Processes == nil {
		o.Processes = ListProcesses
	}
	if o.Now == nil {
		o.Now = time.Now
	}
	if o.Every <= 0 {
		o.Every = Every
	}
	if o.ApplyEvery <= 0 {
		o.ApplyEvery = ApplyEvery
	}
	u := &Updater{o: o, backoff: map[string]time.Time{}, rejected: map[string]bool{}, waitSaid: map[string]bool{}}
	u.recs = loadRecords(o.Dir)
	return u
}

// candidate is a release of one family with a zip and its signature.
type candidate struct {
	family  Family
	tag     string
	version string
	zipURL  string
	sigURL  string
}

// staged is what waits in the companion's home for a game to close.
type staged struct {
	Tag     string `json:"tag"`
	Version string `json:"version"`
}

func (u *Updater) zipPath(f Family) string  { return filepath.Join(u.o.Dir, f.Key+".zip") }
func (u *Updater) sigPath(f Family) string  { return u.zipPath(f) + ".minisig" }
func (u *Updater) metaPath(f Family) string { return filepath.Join(u.o.Dir, f.Key+".json") }

// readStaged is the release waiting for family f, if one is.
func (u *Updater) readStaged(f Family) (staged, bool) {
	b, err := os.ReadFile(u.metaPath(f))
	if err != nil {
		return staged{}, false
	}
	var s staged
	if json.Unmarshal(b, &s) != nil || s.Version == "" {
		return staged{}, false
	}
	return s, true
}

func (u *Updater) clearStaged(f Family) {
	os.Remove(u.metaPath(f))
	os.Remove(u.zipPath(f))
	os.Remove(u.sigPath(f))
}

// Keyed reports whether this build carries a public key, which is what
// lets it install anything at all.
func (u *Updater) Keyed() bool { return u.o.PublicKey != "" }

// Tick is one pass of the schedule: the release list is read at the first
// call and every six hours after, and anything waiting for a game to close
// is tried at that moment and then once a minute. It does nothing when the
// switch is off or the build carries no key. Failures are logged, never
// returned: the companion keeps logging through all of them.
func (u *Updater) Tick(ctx context.Context, now time.Time) {
	poll, apply := u.due(now)
	if !apply {
		return
	}
	u.run.Lock()
	defer u.run.Unlock()
	if poll {
		if err := u.refresh(ctx); err != nil {
			u.o.Log.Warn("the addon update check failed", "component", "addonupdater", "err", err.Error())
		}
	}
	u.apply(ctx)
}

// Due reports whether Tick has work at now, so a caller that ticks every
// second can leave the goroutine and the disk alone the rest of the time.
func (u *Updater) Due(now time.Time) bool {
	if u.o.PublicKey == "" || !u.o.Enabled() {
		return false
	}
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.lastPoll.IsZero() || now.Sub(u.lastPoll) >= u.o.Every {
		return true
	}
	return now.Sub(u.lastLook) >= u.o.ApplyEvery
}

// due decides what a tick does and records that it looked: the release list
// at the first tick and every six hours, and in between, once a minute, the
// release waiting for a game to close.
func (u *Updater) due(now time.Time) (poll, apply bool) {
	if u.o.PublicKey == "" || !u.o.Enabled() {
		return false, false
	}
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.lastPoll.IsZero() || now.Sub(u.lastPoll) >= u.o.Every {
		u.lastPoll, u.lastLook = now, now
		return true, true
	}
	if now.Sub(u.lastLook) < u.o.ApplyEvery {
		return false, false
	}
	u.lastLook = now
	return false, u.anyStaged()
}

func (u *Updater) anyStaged() bool {
	for _, f := range Families {
		if _, ok := u.readStaged(f); ok {
			return true
		}
	}
	return false
}

// eligible reports whether the automatic updates may touch an install's
// family: only where the main addon already exists, and not where that
// folder is a developer's link or checkout.
func eligible(f Family, install string) bool {
	addons := AddonsDir(install)
	if _, found := ReadVersion(addons, Addon.Main()); !found {
		return false
	}
	return !isDevCopy(filepath.Join(addons, Addon.Main())) && !isDevCopy(filepath.Join(addons, f.Main()))
}

// needs lists the installs that family f's version is newer than: the main
// addon where it is older, the data addon where it is older or missing.
func (u *Updater) needs(f Family, version string) []string {
	var out []string
	for _, in := range u.o.Installs() {
		if !eligible(f, in) {
			continue
		}
		have, found := ReadVersion(AddonsDir(in), f.Main())
		if !found || f.Newer(version, have) {
			out = append(out, in)
		}
	}
	return out
}

// refresh lists the releases and stages the newest signed release of each
// family that some install needs.
func (u *Updater) refresh(ctx context.Context) error {
	cands, err := u.candidates(ctx, Families...)
	if err != nil {
		return err
	}
	for _, f := range Families {
		c, ok := cands[f.Key]
		if !ok {
			u.o.Log.Info("the update check found no signed "+f.Key+" release", "component", "addonupdater")
			continue
		}
		if len(u.needs(f, c.version)) == 0 {
			u.o.Log.Info("the "+f.Key+" is up to date everywhere",
				"component", "addonupdater", "newest", c.version)
			u.clearStaged(f)
			continue
		}
		if u.isRejected(c.tag) {
			continue
		}
		if err := u.stage(ctx, c); err != nil {
			u.o.Log.Warn("the "+f.Key+" release "+c.tag+" was not taken",
				"component", "addonupdater", "err", err.Error())
		}
	}
	return nil
}

func (u *Updater) isRejected(tag string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.rejected[tag]
}

func (u *Updater) reject(tag string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.rejected[tag] = true
}

// candidates reads the release list, newest first, until it has found a
// release with a signed zip for every family asked for, or run out of pages.
// A release without a signature is not a candidate: it is skipped, and an
// older signed one is taken instead.
func (u *Updater) candidates(ctx context.Context, fams ...Family) (map[string]candidate, error) {
	out := map[string]candidate{}
	skipped := 0
	// One line per check, not one per tag: a repository with two dozen
	// releases from before signing would otherwise fill the log every time.
	defer func() {
		if skipped == 0 {
			return
		}
		newest := "none"
		if len(out) > 0 {
			tags := make([]string, 0, len(out))
			for _, c := range out {
				tags = append(tags, c.tag)
			}
			sort.Strings(tags)
			newest = strings.Join(tags, ", ")
		}
		noun := "releases"
		if skipped == 1 {
			noun = "release"
		}
		u.o.Log.Info(fmt.Sprintf("the update check skipped %d unsigned %s; newest signed: %s", skipped, noun, newest),
			"component", "addonupdater")
	}()
	for page := 1; page <= maxPages; page++ {
		rels, err := u.list(ctx, page)
		if err != nil {
			return out, err
		}
		for _, rel := range rels {
			if rel.Draft || rel.Prerelease {
				continue
			}
			for _, f := range fams {
				if _, done := out[f.Key]; done || !strings.HasPrefix(rel.TagName, f.TagPrefix) {
					continue
				}
				if c, ok := pick(f, rel); ok {
					out[f.Key] = c
				} else {
					skipped++
				}
			}
		}
		if len(out) == len(fams) || len(rels) < updater.PerPage {
			break
		}
	}
	return out, nil
}

// pick finds the family's zip on a release and the signature beside it.
func pick(f Family, rel updater.Release) (candidate, bool) {
	var zipName, zipURL string
	for _, a := range rel.Assets {
		if strings.HasPrefix(a.Name, f.Main()+"-") && strings.HasSuffix(a.Name, ".zip") {
			zipName, zipURL = a.Name, a.URL
			break
		}
	}
	if zipName == "" {
		return candidate{}, false
	}
	for _, a := range rel.Assets {
		if a.Name == zipName+".minisig" {
			return candidate{family: f, tag: rel.TagName, version: f.VersionOfTag(rel.TagName),
				zipURL: zipURL, sigURL: a.URL}, true
		}
	}
	return candidate{}, false
}

func (u *Updater) list(ctx context.Context, page int) ([]updater.Release, error) {
	url := fmt.Sprintf("%s/repos/%s/releases?per_page=%d&page=%d", u.o.APIBase, u.o.Repo, updater.PerPage, page)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := u.o.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub answered %d for the release list", resp.StatusCode)
	}
	var rels []updater.Release
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4<<20)).Decode(&rels); err != nil {
		return nil, err
	}
	return rels, nil
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
	b, err := io.ReadAll(io.LimitReader(resp.Body, updater.MaxAsset+1))
	if err != nil {
		return nil, err
	}
	if len(b) > updater.MaxAsset {
		return nil, fmt.Errorf("downloading %s: over the %d byte ceiling", url, updater.MaxAsset)
	}
	return b, nil
}

// stage downloads a candidate, verifies it against the public key, proves
// it unpacks and carries the version its tag names, and leaves it in the
// companion's home. A candidate already waiting is kept as it is.
func (u *Updater) stage(ctx context.Context, c candidate) error {
	if u.o.PublicKey == "" {
		return ErrNoKey
	}
	if have, ok := u.readStaged(c.family); ok && have.Tag == c.tag {
		if _, err := os.Stat(u.zipPath(c.family)); err == nil {
			return nil
		}
	}
	body, err := u.fetch(ctx, c.zipURL)
	if err != nil {
		return err
	}
	sig, err := u.fetch(ctx, c.sigURL)
	if err != nil {
		return err
	}
	if err := updater.Verify(u.o.PublicKey, body, sig); err != nil {
		u.reject(c.tag)
		return fmt.Errorf("%s: %w", c.tag, err)
	}
	if err := u.prove(body, c); err != nil {
		u.reject(c.tag)
		return fmt.Errorf("%s: %w", c.tag, err)
	}
	if err := os.MkdirAll(u.o.Dir, 0o700); err != nil {
		return err
	}
	u.clearStaged(c.family)
	if err := os.WriteFile(u.zipPath(c.family), body, 0o600); err != nil {
		return err
	}
	if err := os.WriteFile(u.sigPath(c.family), sig, 0o600); err != nil {
		return err
	}
	meta, _ := json.Marshal(staged{Tag: c.tag, Version: c.version})
	if err := os.WriteFile(u.metaPath(c.family), meta, 0o600); err != nil {
		return err
	}
	u.o.Log.Info("an addon release is ready to install", "component", "addonupdater",
		"tag", c.tag, "bytes", len(body))
	return nil
}

// prove unpacks a zip in a temporary folder and checks that its main TOC
// says the version its release tag names.
func (u *Updater) prove(body []byte, c candidate) error {
	root, done, err := u.unpack(body, c.family)
	if err != nil {
		return err
	}
	defer done()
	got, found := ReadVersion(root, c.family.Main())
	if !found || got != c.version {
		return fmt.Errorf("the zip's %s.toc says version %q, not the %q its tag names", c.family.Main(), got, c.version)
	}
	return nil
}

// unpack extracts a zip into a fresh folder in the companion's home. done
// removes it.
func (u *Updater) unpack(body []byte, f Family) (root string, done func(), err error) {
	if err := os.MkdirAll(u.o.Dir, 0o700); err != nil {
		return "", nil, err
	}
	root, err = os.MkdirTemp(u.o.Dir, "unpack-"+f.Key+"-")
	if err != nil {
		return "", nil, err
	}
	done = func() { os.RemoveAll(root) }
	if err := extract(body, root, f.Folders); err != nil {
		done()
		return "", nil, err
	}
	return root, done, nil
}

// running reports whether the install's game is up, asking the system
// afresh. A process list that cannot be read counts as running.
func (u *Updater) running(install string) bool {
	procs, err := u.o.Processes()
	if err != nil {
		u.mu.Lock()
		said := u.listFailed
		u.listFailed = true
		u.mu.Unlock()
		if !said {
			u.o.Log.Warn("could not list the running programs, so the game is assumed to be running",
				"component", "addonupdater", "err", err.Error())
		}
		return true
	}
	u.mu.Lock()
	u.listFailed = false
	u.mu.Unlock()
	return GameRunning(install, procs)
}

// apply installs what is waiting into every install that needs it and whose
// game is closed. What cannot go in yet stays where it is.
func (u *Updater) apply(ctx context.Context) {
	for _, f := range Families {
		meta, ok := u.readStaged(f)
		if !ok {
			continue
		}
		need := u.needs(f, meta.Version)
		if len(need) == 0 {
			u.clearStaged(f)
			continue
		}
		var ready []string
		for _, in := range need {
			switch {
			case u.backedOff(in, meta.Tag):
			case u.running(in):
				u.sayWaiting(in, meta.Version)
			default:
				u.clearWaiting(in, meta.Version)
				ready = append(ready, in)
			}
		}
		if len(ready) == 0 {
			continue
		}
		body, err := u.verifiedZip(f)
		if err != nil {
			u.o.Log.Warn("the waiting addon release was dropped", "component", "addonupdater", "err", err.Error())
			u.clearStaged(f)
			continue
		}
		root, done, err := u.unpack(body, f)
		if err != nil {
			u.o.Log.Warn("the waiting addon release was dropped", "component", "addonupdater", "err", err.Error())
			u.clearStaged(f)
			continue
		}
		for _, in := range ready {
			if ctx.Err() != nil {
				break
			}
			// The game may have started since the list was read.
			if u.running(in) {
				continue
			}
			u.install(f, in, root, meta)
		}
		done()
		if len(u.needs(f, meta.Version)) == 0 {
			u.clearStaged(f)
		}
	}
}

func (u *Updater) backedOff(install, tag string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.o.Now().Before(u.backoff[install+"|"+tag])
}

// verifiedZip reads the waiting zip back from disk and checks it against
// the key again: the file has been sitting in the home since the download,
// and this is the moment it becomes part of the game.
func (u *Updater) verifiedZip(f Family) ([]byte, error) {
	if u.o.PublicKey == "" {
		return nil, ErrNoKey
	}
	body, err := os.ReadFile(u.zipPath(f))
	if err != nil {
		return nil, err
	}
	sig, err := os.ReadFile(u.sigPath(f))
	if err != nil {
		return nil, fmt.Errorf("the waiting release has no signature: %w", err)
	}
	if err := updater.Verify(u.o.PublicKey, body, sig); err != nil {
		return nil, err
	}
	return body, nil
}

// install swaps a family into one install and records it. A failure puts
// the previous folders back, is logged, and leaves the install alone for an
// hour.
func (u *Updater) install(f Family, install, root string, meta staged) {
	err := replace(AddonsDir(install), root, f.Folders, meta.Version, u.o.Log)
	if err != nil {
		u.mu.Lock()
		u.backoff[install+"|"+meta.Tag] = u.o.Now().Add(retryAfter)
		u.mu.Unlock()
		if errors.Is(err, ErrDevCopy) {
			u.o.Log.Info("an addon folder is a link or a source checkout, so it was not updated",
				"component", "addonupdater", "install", install, "err", err.Error())
			return
		}
		u.o.Log.Warn(fmt.Sprintf("the %s update to %s could not be installed in %s; the folder was left as it was",
			f.Key, meta.Version, install), "component", "addonupdater", "err", err.Error())
		return
	}
	u.recordInstalled(install, f, meta.Version)
	u.o.Log.Info(fmt.Sprintf("installed the %s %s in %s", f.Key, meta.Version, install),
		"component", "addonupdater", "install", install)
}

// Install is the button: it puts the addon, and the data addon beside it,
// into one install that may not have them. It works whatever the switch
// says, and it refuses while that install's game is running.
func (u *Updater) Install(ctx context.Context, install string) error {
	if u.o.PublicKey == "" {
		return ErrNoKey
	}
	u.run.Lock()
	defer u.run.Unlock()
	if u.running(install) {
		return ErrGameRunning
	}
	cands, err := u.candidates(ctx, Families...)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDownload, err)
	}
	if _, ok := cands[Addon.Key]; !ok {
		return ErrNoRelease
	}
	for _, f := range Families {
		c, ok := cands[f.Key]
		if !ok {
			continue
		}
		if err := u.installOne(ctx, install, c); err != nil {
			u.o.Log.Warn("the "+f.Key+" could not be installed", "component", "addonupdater",
				"install", install, "err", err.Error())
			// The addon is what the player asked for; the data addon follows
			// with the next check if it could not come now.
			if f.Key == Addon.Key {
				return err
			}
		}
	}
	u.settle()
	return nil
}

// installOne fetches, verifies and installs one candidate into one install.
func (u *Updater) installOne(ctx context.Context, install string, c candidate) error {
	if err := u.stage(ctx, c); err != nil {
		return fmt.Errorf("%w: %w", ErrDownload, err)
	}
	body, err := u.verifiedZip(c.family)
	if err != nil {
		return err
	}
	root, done, err := u.unpack(body, c.family)
	if err != nil {
		return err
	}
	defer done()
	if err := replace(AddonsDir(install), root, c.family.Folders, c.version, u.o.Log); err != nil {
		return err
	}
	u.recordInstalled(install, c.family, c.version)
	u.o.Log.Info(fmt.Sprintf("installed the %s %s in %s", c.family.Key, c.version, install),
		"component", "addonupdater", "install", install)
	return nil
}

// settle drops a waiting release that no install needs any more.
func (u *Updater) settle() {
	for _, f := range Families {
		if meta, ok := u.readStaged(f); ok && len(u.needs(f, meta.Version)) == 0 {
			u.clearStaged(f)
		}
	}
}

// ---- what the Status page reads -------------------------------------

// FamilyView is one family in one install.
type FamilyView struct {
	// Present is true when the family's main folder has a TOC.
	Present bool
	// Version is what the TOC says, empty when it says nothing.
	Version string
	// Ready is the version waiting to go in, empty when none is.
	Ready string
	// Fresh is true when the companion installed Version and no login has
	// been seen since.
	Fresh bool
}

// InstallView is both families in one install.
type InstallView struct {
	Addon, Data FamilyView
}

// View reads an install. lastActivity is when the game last wrote under the
// install's Logs (the zero time when it never has): a login after the
// companion's own install is what ends "takes effect at your next login".
func (u *Updater) View(install string, lastActivity time.Time) InstallView {
	var out InstallView
	for _, f := range Families {
		v := FamilyView{}
		v.Version, v.Present = ReadVersion(AddonsDir(install), f.Main())
		// A release whose install just failed is not "ready": it will not be tried
		// again for an hour, and the log says why.
		if meta, ok := u.readStaged(f); ok && eligible(f, install) && !u.backedOff(install, meta.Tag) &&
			(!v.Present || f.Newer(meta.Version, v.Version)) {
			v.Ready = meta.Version
		}
		if rec, ok := u.recorded(install, f); ok && v.Present && rec.Version == v.Version &&
			!lastActivity.After(rec.At) {
			v.Fresh = true
		}
		if f.Key == Addon.Key {
			out.Addon = v
		} else {
			out.Data = v
		}
	}
	return out
}

// ---- what the companion remembers it installed ----------------------

// record is one family installed in one install by the companion.
type record struct {
	Version string    `json:"version"`
	At      time.Time `json:"at"`
}

// records is install path -> family key -> record.
type records map[string]map[string]record

const recordsFile = "installed.json"

func loadRecords(dir string) records {
	r := records{}
	if dir == "" {
		return r
	}
	b, err := os.ReadFile(filepath.Join(dir, recordsFile))
	if err == nil {
		_ = json.Unmarshal(b, &r)
	}
	if r == nil {
		r = records{}
	}
	return r
}

func (u *Updater) recorded(install string, f Family) (record, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	rec, ok := u.recs[install][f.Key]
	return rec, ok
}

func (u *Updater) recordInstalled(install string, f Family, version string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.recs[install] == nil {
		u.recs[install] = map[string]record{}
	}
	u.recs[install][f.Key] = record{Version: version, At: u.o.Now()}
	b, err := json.Marshal(u.recs)
	if err == nil {
		err = os.MkdirAll(u.o.Dir, 0o700)
	}
	if err == nil {
		err = os.WriteFile(filepath.Join(u.o.Dir, recordsFile), b, 0o600)
	}
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		u.o.Log.Warn("could not remember the install", "component", "addonupdater", "err", err.Error())
	}
}

// sayWaiting logs that a staged update is waiting for an install's game to
// close, the first time it is seen so: the look repeats every minute and the
// line must not.
func (u *Updater) sayWaiting(install, version string) {
	key := install + "|" + version
	u.mu.Lock()
	said := u.waitSaid[key]
	u.waitSaid[key] = true
	u.mu.Unlock()
	if !said {
		u.o.Log.Info("an addon update is waiting for the game to close",
			"component", "addonupdater", "install", install, "version", version)
	}
}

// clearWaiting forgets that the wait was logged once the game is closed, so
// a later wait for the same update is said again.
func (u *Updater) clearWaiting(install, version string) {
	u.mu.Lock()
	delete(u.waitSaid, install+"|"+version)
	u.mu.Unlock()
}
