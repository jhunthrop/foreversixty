// companion/internal/secret/secret.go
// Package secret keeps the device token out of plain files wherever the
// operating system offers somewhere better. The keychain is tried
// first; a machine with no keychain, no unlocked login session, or no
// D-Bus falls back to config.json, which the config package already
// writes with owner-only permissions.
package secret

import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/zalando/go-keyring"
)

// Service and Account are the keychain coordinates of the device token.
const (
	Service = "ForeverSixty Companion"
	Account = "device-token"
)

// Backend names where a store keeps the token, for the settings page
// and the troubleshooting section of the README.
type Backend string

// The two places a token can live.
const (
	Keychain Backend = "keychain"
	File     Backend = "file"
)

// Store reads and writes the device token.
type Store interface {
	// Token is the stored token, or the empty string when the device
	// is not paired.
	Token() (string, error)
	SetToken(token string) error
	Clear() error
	Backend() Backend
}

// Keyring is the OS keychain: Keychain on macOS, Credential Manager on
// Windows, the Secret Service on Linux.
type Keyring struct{}

// Token reads the token from the keychain.
func (Keyring) Token() (string, error) {
	v, err := keyring.Get(Service, Account)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("read the device token from the keychain: %w", err)
	}
	return v, nil
}

// SetToken writes the token to the keychain.
func (Keyring) SetToken(token string) error {
	if err := keyring.Set(Service, Account, token); err != nil {
		return fmt.Errorf("write the device token to the keychain: %w", err)
	}
	return nil
}

// Clear removes the token. A token that is not there is not an error:
// unpairing twice must succeed.
func (Keyring) Clear() error {
	err := keyring.Delete(Service, Account)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("delete the device token from the keychain: %w", err)
	}
	return nil
}

// Backend names this store.
func (Keyring) Backend() Backend { return Keychain }

// ConfigFile keeps the token in config.json's device_token field.
type ConfigFile struct{ Path string }

// Token reads the token from config.json.
func (c ConfigFile) Token() (string, error) {
	cfg, err := config.Load(c.Path)
	if err != nil {
		return "", err
	}
	return cfg.DeviceToken, nil
}

// SetToken writes the token into config.json, leaving every other
// setting as it is.
func (c ConfigFile) SetToken(token string) error {
	cfg, err := config.Load(c.Path)
	if err != nil {
		return err
	}
	cfg.DeviceToken = token
	return config.Save(c.Path, cfg)
}

// Clear empties the token field.
func (c ConfigFile) Clear() error { return c.SetToken("") }

// Backend names this store.
func (ConfigFile) Backend() Backend { return File }

// Open returns the keychain store when the keychain answers a probe,
// and the config-file store when it does not. The probe is a read: a
// write would leave an entry behind on a machine that then falls back.
func Open(configPath string) Store {
	if _, err := keyring.Get(Service, Account); err == nil || errors.Is(err, keyring.ErrNotFound) {
		return Keyring{}
	}
	return ConfigFile{Path: configPath}
}

// Migrate moves a token found in config.json into the keychain and
// blanks the file copy. It runs at startup so an install that once fell
// back stops leaving the token on disk once the keychain works. It
// never overwrites a token the keychain already holds: once pairing has
// written a fresh token straight into the keychain, a stale
// config.json left over from an earlier fallback episode must not
// clobber it.
func Migrate(s Store, configPath string) error {
	if s.Backend() != Keychain {
		return nil
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}
	if cfg.DeviceToken == "" {
		return nil
	}
	current, err := s.Token()
	if err != nil {
		return err
	}
	if current == "" {
		if err := s.SetToken(cfg.DeviceToken); err != nil {
			return err
		}
	}
	// Either the token was just migrated, or the keychain already held
	// one and the config.json copy is stale. Either way, a plaintext
	// token left on disk once the keychain is the live backend is the
	// exact leak this package exists to close, so the file copy is
	// blanked regardless of which branch above ran.
	cfg.DeviceToken = ""
	return config.Save(configPath, cfg)
}
