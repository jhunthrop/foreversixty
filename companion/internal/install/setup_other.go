//go:build !windows

// companion/internal/install/setup_other.go
// Only Windows has the setup: the Mac moves itself into Applications
// (offer.go) and Linux runs the binary where it is.
package install

import "errors"

// SetupSupported reports whether this platform has the setup.
const SetupSupported = false

// NewWinSystem is nil off Windows; nothing calls it there.
func NewWinSystem() WinSystem { return nil }

// WindowsFolders has no answer off Windows.
func WindowsFolders() (Folders, error) {
	return Folders{}, errors.New("install: the setup is Windows only")
}
