package leveling

import (
	"os"
	"path/filepath"
	"testing"
)

func writeActiveBuild(t *testing.T, body string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "web", "src", "data", "active-build.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestReadActiveBuild(t *testing.T) {
	build, err := ReadActiveBuild(writeActiveBuild(t, `{"build":"testbuild"}`))
	if err != nil {
		t.Fatalf("ReadActiveBuild: %v", err)
	}
	if build != "testbuild" {
		t.Errorf("build = %q, want testbuild", build)
	}
}

func TestReadActiveBuildRefusesBadFiles(t *testing.T) {
	for name, dir := range map[string]string{
		"missing file":      t.TempDir(),
		"empty build field": writeActiveBuild(t, `{"build":""}`),
		"malformed JSON":    writeActiveBuild(t, `{not json`),
	} {
		if _, err := ReadActiveBuild(dir); err == nil {
			t.Errorf("%s: want an error, got nil", name)
		}
	}
}
