// logs/internal/sample/sample.go
// Package sample fetches the real retail combat log the benchmark and the
// conformance tests run against. The file is 77 MB and its repository is
// AGPL-3.0 licensed, so it is never committed: it is downloaded once into
// a git-ignored cache and reused.
package sample

import (
	"context"
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
const URL = "https://raw.githubusercontent.com/rp4rk/WoWP/main/WoWCombatLog.txt"

// Name is the file's name inside the cache.
const Name = "wowp-retail-v16.txt"

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

// Fetch returns the path to the cached sample, downloading it if needed.
// It returns an error wrapping ErrUnavailable when the machine is offline
// and nothing is cached.
func Fetch(ctx context.Context) (string, error) {
	if p := os.Getenv(EnvOverride); p != "" {
		if _, err := os.Stat(p); err != nil {
			return "", fmt.Errorf("%w: %s is set to %q: %v", ErrUnavailable, EnvOverride, p, err)
		}
		return p, nil
	}
	path := filepath.Join(CacheDir(), Name)
	if fi, err := os.Stat(path); err == nil && fi.Size() > 0 {
		return path, nil
	}
	if err := os.MkdirAll(CacheDir(), 0o755); err != nil {
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
	if err := os.Rename(tmp, path); err != nil {
		return "", fmt.Errorf("%w: rename: %v", ErrUnavailable, err)
	}
	return path, nil
}
