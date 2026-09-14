// logs/internal/sample/sample_test.go
package sample

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// withTempCache redirects Fetch's cache location to a t.TempDir() so tests
// that exercise the download path never touch the real testdata/cache (and
// the real sample cached there).
func withTempCache(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	old := cacheDir
	cacheDir = func() string { return dir }
	t.Cleanup(func() { cacheDir = old })
	return dir
}

// withTestServer points URL at an httptest.Server and restores it after the
// test, so tests never make a real network request.
func withTestServer(t *testing.T, handler http.HandlerFunc) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	old := URL
	URL = server.URL
	t.Cleanup(func() { URL = old })
}

func TestTheOverrideEnvironmentVariableWins(t *testing.T) {
	local := filepath.Join(t.TempDir(), "local.txt")
	if err := os.WriteFile(local, []byte("9/26 20:10:00.000  ZONE_CHANGE,1,\"Z\",0\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv(EnvOverride, local)
	got, err := Fetch(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if got != local {
		t.Fatalf("path = %q, want %q", got, local)
	}
}

func TestAnOverridePointingAtNothingIsUnavailable(t *testing.T) {
	t.Setenv(EnvOverride, filepath.Join(t.TempDir(), "absent.txt"))
	_, err := Fetch(t.Context())
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
}

func TestCacheDirSitsUnderTheModulesTestdata(t *testing.T) {
	got := CacheDir()
	if !strings.HasSuffix(filepath.ToSlash(got), "testdata/cache") {
		t.Fatalf("cache dir = %q", got)
	}
}

func TestAnOverridePointingAtADirectoryIsRejected(t *testing.T) {
	t.Setenv(EnvOverride, t.TempDir())
	_, err := Fetch(t.Context())
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
}

func TestACacheHitMakesNoRequest(t *testing.T) {
	dir := withTempCache(t)
	requested := false
	withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		requested = true
		w.WriteHeader(http.StatusOK)
	})
	cached := filepath.Join(dir, Name)
	if err := os.WriteFile(cached, []byte("already here"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := Fetch(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if got != cached {
		t.Fatalf("path = %q, want %q", got, cached)
	}
	if requested {
		t.Fatal("Fetch made a request although the sample was already cached")
	}
}

func TestASuccessfulDownloadLandsTheFileAtomically(t *testing.T) {
	dir := withTempCache(t)
	const body = "9/26 20:10:00.000  ZONE_CHANGE,1,\"Z\",0\n"
	withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, body)
	})
	got, err := Fetch(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(dir, Name)
	if got != want {
		t.Fatalf("path = %q, want %q", got, want)
	}
	b, err := os.ReadFile(want)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != body {
		t.Fatalf("content = %q, want %q", b, body)
	}
	if _, err := os.Stat(want + ".partial"); !os.IsNotExist(err) {
		t.Fatalf(".partial file survived a successful download: %v", err)
	}
}

func TestANon200ResponseLeavesNoFile(t *testing.T) {
	dir := withTempCache(t)
	withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	_, err := Fetch(t.Context())
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
	if _, err := os.Stat(filepath.Join(dir, Name)); !os.IsNotExist(err) {
		t.Fatalf("cache file exists after a non-200 response: %v", err)
	}
}

func TestATruncatedDownloadLeavesNoPartialFile(t *testing.T) {
	dir := withTempCache(t)
	withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1000")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "only a little")
		hj, ok := w.(http.Hijacker)
		if !ok {
			t.Fatal("ResponseWriter does not support hijacking")
		}
		conn, _, err := hj.Hijack()
		if err != nil {
			t.Fatal(err)
		}
		conn.Close()
	})
	_, err := Fetch(t.Context())
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
	want := filepath.Join(dir, Name)
	if _, err := os.Stat(want); !os.IsNotExist(err) {
		t.Fatalf("cache file exists after a truncated download: %v", err)
	}
	if _, err := os.Stat(want + ".partial"); !os.IsNotExist(err) {
		t.Fatalf(".partial file survived a truncated download: %v", err)
	}
}

func TestACacheDirectoryThatCannotBeCreatedIsUnavailable(t *testing.T) {
	dir := t.TempDir()
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	old := cacheDir
	cacheDir = func() string { return filepath.Join(blocker, "cache") }
	t.Cleanup(func() { cacheDir = old })
	requested := false
	withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		requested = true
		w.WriteHeader(http.StatusOK)
	})
	_, err := Fetch(t.Context())
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("err = %v, want ErrUnavailable", err)
	}
	if requested {
		t.Fatal("Fetch made a request although the cache directory could not be created")
	}
}
