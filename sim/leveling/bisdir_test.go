package leveling

import (
	"os"
	"path/filepath"
	"testing"
)

func addBisBuild(t *testing.T, root, build string, files ...string) string {
	t.Helper()
	dir := filepath.Join(root, build)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		return dir
	}
	if err := os.MkdirAll(filepath.Join(dir, "bis"), 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range files {
		if err := os.WriteFile(filepath.Join(dir, "bis", name), []byte("{}"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestBisDirKeepsABuildThatHasItsOwnFiles(t *testing.T) {
	root := t.TempDir()
	addBisBuild(t, root, "1.60.1.70009", "a.json")
	active := addBisBuild(t, root, "1.60.1.70291", "a.json")
	if got, want := BisDir(active), filepath.Join(active, "bis"); got != want {
		t.Fatalf("BisDir = %s, want %s", got, want)
	}
}

func TestBisDirFallsBackToTheNewestRankedBuild(t *testing.T) {
	root := t.TempDir()
	addBisBuild(t, root, "1.60.1.69893", "a.json")
	older := addBisBuild(t, root, "1.60.1.70009", "a.json")
	active := addBisBuild(t, root, "1.60.1.70291", "a.md")
	if got, want := BisDir(active), filepath.Join(filepath.Dir(older), "1.60.1.70009", "bis"); got != want {
		t.Fatalf("BisDir = %s, want %s", got, want)
	}
}

func TestBisDirNeverFallsBackForANonBuildDirectory(t *testing.T) {
	root := t.TempDir()
	addBisBuild(t, root, "1.60.1.70009", "a.json")
	missing := filepath.Join(root, "no-such-build")
	if got, want := BisDir(missing), filepath.Join(missing, "bis"); got != want {
		t.Fatalf("BisDir = %s, want %s", got, want)
	}
}

func TestBisDirWithNoRankedBuildAnswersItsOwnPath(t *testing.T) {
	root := t.TempDir()
	active := addBisBuild(t, root, "1.60.1.70291")
	if got, want := BisDir(active), filepath.Join(active, "bis"); got != want {
		t.Fatalf("BisDir = %s, want %s", got, want)
	}
}
