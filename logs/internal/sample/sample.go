// logs/internal/sample/sample.go
// Package sample fetches the real retail combat log the benchmark and the
// conformance tests run against. The file is 77 MB and its repository is
// AGPL-3.0 licensed, so it is never committed: it is downloaded once into
// a git-ignored cache and reused.
package sample

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// URL is the retail sample: COMBAT_LOG_VERSION 16, build 9.0.2, advanced
// logging on, 105 COMBATANT_INFO lines, eighteen encounters.
//
// It is a var, not a const, only so tests in this package can point it at
// an httptest.Server; production code must never assign to it.
var URL = "https://raw.githubusercontent.com/rp4rk/WoWP/main/WoWCombatLog.txt"

// Name is the file's name inside the cache.
const Name = "wowp-retail-v16.txt"

// SHA256 and Size pin the exact bytes every measurement in this module
// describes: the 272,367-event log the retail layout row was verified
// against, the benchmark's 77 MB, and the "0 parse errors" conformance
// evidence. URL points at a third-party repository that can edit or
// replace the file at any time, so a download that does not match these
// is refused rather than silently measured.
const (
	SHA256 = "72b3ee25ac51b0e08c2b250e71171ec4c22ab6069df961e946031105c1cfa2bf"
	Size   = 76979002
)

// pinned is what verify checks a file against. Like URL it is a var only
// so the tests in this package can substitute a small stand-in file;
// production code must never assign to it.
var pinned = struct {
	sha  string
	size int64
}{SHA256, Size}

// verify reads path and reports whether it is the pinned sample. It is
// applied to a fresh download before the rename and to a cache hit, since
// a cached copy can be stale from before an upstream change.
func verify(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%w: open %s: %v", ErrUnavailable, path, err)
	}
	defer f.Close()
	h := sha256.New()
	n, err := io.Copy(h, f)
	if err != nil {
		return fmt.Errorf("%w: read %s: %v", ErrUnavailable, path, err)
	}
	got := hex.EncodeToString(h.Sum(nil))
	if n != pinned.size || got != pinned.sha {
		return fmt.Errorf("%w: %s is not the pinned sample: got sha256 %s (%d bytes), "+
			"want %s (%d bytes). %s may have been changed upstream; delete the cached "+
			"copy to re-download, or point %s at a known-good local file",
			ErrUnavailable, path, got, n, pinned.sha, pinned.size, URL, EnvOverride)
	}
	return nil
}

// EnvOverride points at a local copy instead of downloading.
const EnvOverride = "FOREVER_LOGS_SAMPLE"

// ErrUnavailable is returned when the sample is neither cached nor
// reachable. Callers skip rather than fail: the suite must pass offline.
var ErrUnavailable = errors.New("sample: the retail sample is not available")

// CacheDir is logs/testdata/cache, resolved from this file's own path so it
// works whatever directory a test runs in.
func CacheDir() string {
	_, self, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join("testdata", "cache")
	}
	// self is <module>/internal/sample/sample.go.
	return filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(self))), "testdata", "cache")
}

// cacheDir is what Fetch actually writes into. It defaults to CacheDir, but
// tests in this package redirect it to a t.TempDir() so exercising the
// download path never touches the real testdata/cache or the real sample
// cached there.
var cacheDir = CacheDir

// Fetch returns the path to the cached sample, downloading it if needed,
// and verifies it against the pinned digest. It returns an error wrapping
// ErrUnavailable when the machine is offline and nothing is cached, and
// when what is there is not the pinned file.
//
// A file supplied through EnvOverride is deliberately not checked: the
// override exists so a caller can point the benchmark at their own log,
// and hashing it would defeat that.
func Fetch(ctx context.Context) (string, error) {
	if p := os.Getenv(EnvOverride); p != "" {
		fi, err := os.Stat(p)
		if err != nil {
			return "", fmt.Errorf("%w: %s is set to %q: %v", ErrUnavailable, EnvOverride, p, err)
		}
		if !fi.Mode().IsRegular() {
			return "", fmt.Errorf("%w: %s is set to %q, which is not a regular file", ErrUnavailable, EnvOverride, p)
		}
		return p, nil
	}
	path := filepath.Join(cacheDir(), Name)
	if fi, err := os.Stat(path); err == nil && fi.Size() > 0 {
		if err := verify(path); err != nil {
			return "", err
		}
		return path, nil
	}
	if err := os.MkdirAll(cacheDir(), 0o755); err != nil {
		return "", fmt.Errorf("%w: create cache: %v", ErrUnavailable, err)
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return "", fmt.Errorf("%w: build request: %v", ErrUnavailable, err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrUnavailable, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: %s returned %s", ErrUnavailable, URL, resp.Status)
	}
	tmp := path + ".partial"
	f, err := os.Create(tmp)
	if err != nil {
		return "", fmt.Errorf("%w: create %s: %v", ErrUnavailable, tmp, err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		os.Remove(tmp)
		return "", fmt.Errorf("%w: download: %v", ErrUnavailable, err)
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return "", fmt.Errorf("%w: close %s: %v", ErrUnavailable, tmp, err)
	}
	if err := verify(tmp); err != nil {
		os.Remove(tmp)
		return "", err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return "", fmt.Errorf("%w: rename: %v", ErrUnavailable, err)
	}
	return path, nil
}
