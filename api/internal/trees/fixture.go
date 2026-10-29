package trees

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// FixtureDir is the testdata build directory shipped with this package. It
// resolves from this file's own compiled-in path, so tests in other packages
// can use it without knowing their working directory. It is test support
// only; nothing in the running service calls it.
func FixtureDir() string {
	_, self, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(self), "testdata", "builds")
}

// LoadFixture loads the two-tree fixture build "test-1".
func LoadFixture() (*Data, error) { return Load(FixtureDir()) }

// LinkBuilds makes dst a builds directory holding the named builds of src without copying
// their megabytes: each build is a real directory (Load lists directories) whose entries
// are symlinks to src's. Test support only, for proving the active-build rule against the
// real data with a chosen ActiveBuildFile (WriteActiveBuildFile).
func LinkBuilds(dst, src string, builds ...string) error {
	absSrc, err := filepath.Abs(src)
	if err != nil {
		return fmt.Errorf("trees: link builds: %w", err)
	}
	for _, build := range builds {
		entries, err := os.ReadDir(filepath.Join(absSrc, build))
		if err != nil {
			return fmt.Errorf("trees: link builds: %w", err)
		}
		if err := os.MkdirAll(filepath.Join(dst, build), 0o755); err != nil {
			return fmt.Errorf("trees: link builds: %w", err)
		}
		for _, e := range entries {
			if err := os.Symlink(filepath.Join(absSrc, build, e.Name()), filepath.Join(dst, build, e.Name())); err != nil {
				return fmt.Errorf("trees: link builds: %w", err)
			}
		}
	}
	return nil
}

// WriteActiveBuildFile writes dir/active-build.json naming build, in the shape
// tools/activate-build.sh writes web/src/data/active-build.json. Test support only.
func WriteActiveBuildFile(dir, build string) error {
	body := fmt.Sprintf("{\n  \"build\": %q\n}\n", build)
	if err := os.WriteFile(filepath.Join(dir, ActiveBuildFile), []byte(body), 0o644); err != nil {
		return fmt.Errorf("trees: write active build file: %w", err)
	}
	return nil
}
