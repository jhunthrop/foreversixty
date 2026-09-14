package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOfAMissingFileGivesTheDefaults(t *testing.T) {
	got, err := Load(filepath.Join(t.TempDir(), "config.json"))
	if err != nil {
		t.Fatal(err)
	}
	if got.APIBaseURL != DefaultAPIBaseURL || got.SiteBaseURL != DefaultSiteBaseURL ||
		got.ReportVisibility != "public" || len(got.WoWPaths) != 0 {
		t.Fatalf("Load = %+v", got)
	}
}

func TestSaveThenLoadRoundTrips(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	want := Config{
		APIBaseURL:       "http://127.0.0.1:8080",
		SiteBaseURL:      "http://127.0.0.1:4321",
		DeviceToken:      "fsd_abcdef",
		WoWPaths:         []string{"/Applications/World of Warcraft/_classic_era_"},
		ReportVisibility: "unlisted",
	}
	if err := Save(path, want); err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if perm := fi.Mode().Perm(); perm != 0o600 {
		t.Errorf("permissions = %v, want 0600", perm)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if got.APIBaseURL != want.APIBaseURL || got.DeviceToken != want.DeviceToken ||
		got.ReportVisibility != want.ReportVisibility || len(got.WoWPaths) != 1 ||
		got.WoWPaths[0] != want.WoWPaths[0] {
		t.Fatalf("Load = %+v, want %+v", got, want)
	}
}

func TestAnUnknownVisibilityIsRefused(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := Save(path, Config{APIBaseURL: "x", SiteBaseURL: "y", ReportVisibility: "everyone"}); err == nil {
		t.Fatal("Save accepted an unknown visibility")
	}
	if err := os.WriteFile(path, []byte(`{"report_visibility":"everyone"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil {
		t.Fatal("Load accepted an unknown visibility")
	}
}

func TestAnEmptyAPIBaseURLIsRefused(t *testing.T) {
	if err := (Config{SiteBaseURL: "y", ReportVisibility: "public"}).Validate(); err == nil {
		t.Fatal("Validate accepted an empty api_base_url")
	}
	if err := (Config{APIBaseURL: "x", ReportVisibility: "public"}).Validate(); err == nil {
		t.Fatal("Validate accepted an empty site_base_url")
	}
}

func TestABaseURLMustBeHTTPOrHTTPS(t *testing.T) {
	for _, tc := range []struct {
		name string
		api  string
		site string
		ok   bool
	}{
		{"the defaults", DefaultAPIBaseURL, DefaultSiteBaseURL, true},
		{"a local server", "http://127.0.0.1:8080", "http://127.0.0.1:4321", true},
		// site_base_url becomes the href of a link in the local page,
		// so a script URL there is a script URL inside the webview.
		{"a javascript site URL", DefaultAPIBaseURL, "javascript:alert(1)", false},
		{"a javascript API URL", "javascript:alert(1)", DefaultSiteBaseURL, false},
		{"a file URL", DefaultAPIBaseURL, "file:///etc/passwd", false},
		{"a bare host", DefaultAPIBaseURL, "foreversixty.gg", false},
		{"no host", DefaultAPIBaseURL, "https://", false},
		{"an empty site URL", DefaultAPIBaseURL, "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := Default()
			c.APIBaseURL, c.SiteBaseURL = tc.api, tc.site
			if err := c.Validate(); (err == nil) != tc.ok {
				t.Fatalf("Validate() = %v, want ok = %v", err, tc.ok)
			}
		})
	}
}
