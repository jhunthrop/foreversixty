// companion/internal/updater/installed.go
// Apps & Features shows the version the Windows setup wrote into its
// uninstall entry (internal/install, setup.go). The companion updates
// itself without the setup, so each launch brings that entry up to the
// running version: only when the setup made the entry, and only for the
// copy it installed.
package updater

import (
	"errors"
	"path/filepath"
	"strings"
)

// UninstallKey is the setup's Apps & Features entry under HKEY_CURRENT_USER
// (internal/install writes it, this keeps its version current).
const UninstallKey = `Software\Microsoft\Windows\CurrentVersion\Uninstall\ForeverSixty`

// The values the entry carries that this reads and writes.
const (
	valueInstallLocation = "InstallLocation"
	valueDisplayVersion  = "DisplayVersion"
)

// RegistryKey is one open registry key's string values.
type RegistryKey interface {
	GetString(name string) (string, error)
	SetString(name, value string) error
	Close() error
}

// ErrNoRegistryKey says the key does not exist: the companion was not installed by
// the setup.
var ErrNoRegistryKey = errors.New("updater: no such registry key")

// OpenKey opens a key under HKEY_CURRENT_USER for reading and writing, or
// returns ErrNoRegistryKey.
type OpenKey func(path string) (RegistryKey, error)

// SyncInstalledVersion writes version as the setup entry's DisplayVersion
// when the entry exists, names exe's folder as its install location, and
// says another version. It reports whether it wrote.
func SyncInstalledVersion(open OpenKey, exe, version string) (bool, error) {
	key, err := open(UninstallKey)
	if errors.Is(err, ErrNoRegistryKey) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	defer key.Close()
	location, err := key.GetString(valueInstallLocation)
	if err != nil || !sameDir(location, filepath.Dir(exe)) {
		return false, nil // a copy the setup did not install, or an entry without a folder
	}
	if current, err := key.GetString(valueDisplayVersion); err == nil && current == version {
		return false, nil
	}
	if err := key.SetString(valueDisplayVersion, version); err != nil {
		return false, err
	}
	return true, nil
}

// sameDir compares two Windows folders as Windows does: case aside, with or
// without a trailing separator.
func sameDir(a, b string) bool {
	clean := func(p string) string {
		return strings.ToLower(strings.TrimRight(filepath.Clean(p), `\/`))
	}
	return a != "" && clean(a) == clean(b)
}
