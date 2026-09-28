package leveling

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFile creates path (and its parent directories) with contents -
// for tests that need a layout the committed testdata/reporoot fixture
// does not carry (e.g. malformed input).
func writeFile(t *testing.T, path, contents string) error {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(contents), 0o644)
}
