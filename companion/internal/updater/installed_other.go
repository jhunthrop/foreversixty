//go:build !windows

// companion/internal/updater/installed_other.go
// Only Windows has the setup's uninstall entry.
package updater

// OpenUserKey finds no key off Windows.
func OpenUserKey(string) (RegistryKey, error) { return nil, ErrNoRegistryKey }
