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

	// Stage the verified bytes as a sibling of exe first, so the swap
	// below is two same-directory renames rather than a write in
	// place: Windows will not let a running binary be replaced, but
	// it will let one be renamed aside and a new file renamed into
	// its place, and that target-path-is-free rename is effectively
	// instant, unlike a full write.
	tmp, err := os.CreateTemp(filepath.Dir(exe), filepath.Base(exe)+".update-*")
	if err != nil {
		return false, fmt.Errorf("stage the new binary next to the running one: %w", err)
	}
	tmpPath := tmp.Name()
	moved := false
	defer func() {
		if !moved {
			os.Remove(tmpPath) // leave nothing behind for the next launch to trip over
		}
	}()
	if _, err := tmp.Write(body); err != nil {
		tmp.Close()
		return false, fmt.Errorf("write the new binary: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return false, fmt.Errorf("write the new binary: %w", err)
	}
	if err := os.Chmod(tmpPath, 0o755); err != nil {
		return false, fmt.Errorf("make the new binary executable: %w", err)
	}

	previous := filepath.Join(dir, "previous")
	os.Remove(previous)
	if err := os.Rename(exe, previous); err != nil {
		return false, fmt.Errorf("move the running binary aside: %w", err)
	}
	if err := os.Rename(tmpPath, exe); err != nil {
		// The target is free now, so this should be rare, but leaving
		// nothing at exe is worse than a stale binary: restore it, and
		// report a restore failure rather than swallow it.
		if restoreErr := os.Rename(previous, exe); restoreErr != nil {
			return false, fmt.Errorf("swap in the new binary: %w; restoring the previous binary also failed: %w", err, restoreErr)
		}
		return false, fmt.Errorf("swap in the new binary: %w", err)
	}
	moved = true
	clearPending(dir)
	return true, nil
}

func clearPending(dir string) {
	os.Remove(PendingPath(dir))
	os.Remove(pendingSig(dir))
}
