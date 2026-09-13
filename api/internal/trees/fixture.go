package trees

import (
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
