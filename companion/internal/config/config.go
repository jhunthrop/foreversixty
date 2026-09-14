// companion/internal/config/config.go
// Package config is config.json: the four settings the interface
// contract names, loaded with defaults and saved atomically.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"slices"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
)

// DefaultAPIBaseURL is where the companion talks to unless the user or
// a test points it somewhere else, and DefaultSiteBaseURL is where a
// finished report can be read.
const (
	DefaultAPIBaseURL  = "https://api.foreversixty.gg"
	DefaultSiteBaseURL = "https://foreversixty.gg"
)

// Visibilities are the values the API accepts for a new report.
var Visibilities = []string{"public", "unlisted", "private", "guild"}

// Config is config.json.
type Config struct {
	APIBaseURL string `json:"api_base_url"`
	// SiteBaseURL is where report links point. It is additive to the
	// four keys the contract names, so a staging build can send its
	// links somewhere else.
	SiteBaseURL string `json:"site_base_url"`
	// DeviceToken is only written when no OS keychain is available; the
	// secret store owns the token otherwise and leaves this empty.
	DeviceToken      string   `json:"device_token"`
	WoWPaths         []string `json:"wow_paths"`
	ReportVisibility string   `json:"report_visibility"`
	// LoggingCharacter is who the reports are attributed to, chosen on
	// the settings page from the characters the addon sync has seen.
	// It is additive to the four keys the contract names and is
	// omitted when the player has not chosen one.
	LoggingCharacter *character.Character `json:"logging_character,omitempty"`
}

// Default is the configuration a fresh install starts from.
func Default() Config {
	return Config{
		APIBaseURL:       DefaultAPIBaseURL,
		SiteBaseURL:      DefaultSiteBaseURL,
		WoWPaths:         []string{},
		ReportVisibility: "public",
	}
}

// baseURL rejects anything that is not an absolute http or https
// URL. site_base_url becomes the href of a link in the local page, so
// a javascript: value in config.json would otherwise be a clickable
// script URL inside the webview, and api_base_url is where the device
// token is sent.
func baseURL(field, raw string) error {
	if raw == "" {
		return fmt.Errorf("%s is empty", field)
	}
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("%s is not a URL: %w", field, err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("%s must be an http or https URL, not %q", field, u.Scheme)
	}
	if u.Host == "" {
		return fmt.Errorf("%s has no host", field)
	}
	return nil
}

// Validate rejects a configuration the rest of the app cannot use.
func (c Config) Validate() error {
	if err := baseURL("api_base_url", c.APIBaseURL); err != nil {
		return err
	}
	if err := baseURL("site_base_url", c.SiteBaseURL); err != nil {
		return err
	}
	if !slices.Contains(Visibilities, c.ReportVisibility) {
		return fmt.Errorf("report_visibility %q is not one of %v", c.ReportVisibility, Visibilities)
	}
	return nil
}

// Load reads config.json. A missing file is not an error: it yields the
// defaults, so a first run needs no installer step.
func Load(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return Default(), nil
	}
	if err != nil {
		return Config{}, fmt.Errorf("read %s: %w", path, err)
	}
	c := Default()
	if err := json.Unmarshal(b, &c); err != nil {
		return Config{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if c.WoWPaths == nil {
		c.WoWPaths = []string{}
	}
	if err := c.Validate(); err != nil {
		return Config{}, fmt.Errorf("%s: %w", path, err)
	}
	return c, nil
}

// Save writes config.json atomically with owner-only permissions, so a
// crash mid-write can never leave an unreadable configuration and the
// fallback token is not world-readable.
func Save(path string, c Config) error {
	if err := c.Validate(); err != nil {
		return err
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')
	tmp, err := os.CreateTemp(filepath.Dir(path), ".config-*.json")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(name, path)
}
