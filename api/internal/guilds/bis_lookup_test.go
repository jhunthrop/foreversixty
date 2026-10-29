package guilds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// BisBuild reads the filesystem once and keeps the answer while DataDir and Trees stay the
// same; a new Trees (or DataDir) works it out again.
func TestBisBuildIsKeptUntilItsInputsChange(t *testing.T) {
	repo := realTreesStore(t)
	latest, ok := repo.Trees.Active()
	if !ok {
		t.Fatal("no build loaded")
	}
	dir := t.TempDir()
	const fallback = "1.0.0.1"
	bisFile := filepath.Join(dir, fallback, "bis", "mage-fire.json")
	for _, d := range []string{filepath.Join(dir, latest.Version), filepath.Dir(bisFile)} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(bisFile, []byte(`{"bands":[]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	s := &Store{DataDir: dir, Trees: repo.Trees}
	if got := s.BisBuild(); got != fallback {
		t.Fatalf("first read = %q, want the fallback %q", got, fallback)
	}
	if err := os.Remove(bisFile); err != nil {
		t.Fatal(err)
	}
	if got := s.BisBuild(); got != fallback {
		t.Fatalf("second read = %q, want the kept %q", got, fallback)
	}
	reloaded, err := trees.Load(repo.DataDir)
	if err != nil {
		t.Fatal(err)
	}
	s.Trees = reloaded
	if got := s.BisBuild(); got != latest.Version {
		t.Fatalf("after a new Trees = %q, want %q", got, latest.Version)
	}
}
