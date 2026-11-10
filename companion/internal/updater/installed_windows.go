//go:build windows

// companion/internal/updater/installed_windows.go
// The real registry behind SyncInstalledVersion.
package updater

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

// OpenUserKey opens a key under HKEY_CURRENT_USER to read and set values.
func OpenUserKey(path string) (RegistryKey, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE|registry.SET_VALUE)
	if errors.Is(err, registry.ErrNotExist) {
		return nil, ErrNoRegistryKey
	}
	if err != nil {
		return nil, err
	}
	return userKey{k}, nil
}

type userKey struct{ k registry.Key }

func (u userKey) GetString(name string) (string, error) {
	v, _, err := u.k.GetStringValue(name)
	return v, err
}

func (u userKey) SetString(name, value string) error { return u.k.SetStringValue(name, value) }

func (u userKey) Close() error { return u.k.Close() }
