// logs/internal/sample/sample_test.go
package sample

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
