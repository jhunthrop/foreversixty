package config

import (
	"fmt"
	"regexp"
	"strings"
)

// AppConfig is the mobile apps' API settings. Enabled (APP_API=true) mounts the app routes;
// off by default, so they land dark. The rest is what GET /v1/app/config answers, kept in
// config so a store ruling or a breaking change needs no app update.
type AppConfig struct {
	Enabled bool
	// WebCheckoutRegions are the storefront regions (ISO 3166 alpha-2) where an app may link
	// to web checkout (APP_WEB_CHECKOUT_REGIONS, "US"; "none" for no region).
	WebCheckoutRegions []string
	// MinVersionIOS and MinVersionAndroid are the oldest app builds the API still serves
	// (APP_MIN_VERSION_IOS, APP_MIN_VERSION_ANDROID; "1.0.0").
	MinVersionIOS, MinVersionAndroid string
}

const (
	defaultAppCheckoutRegions = "US"
	defaultAppMinVersion      = "1.0.0"
	noAppRegions              = "none"
)

var (
	appRegionPattern  = regexp.MustCompile(`^[A-Z]{2}$`)
	appVersionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
)

// loadApp reads the app settings; errors name the variable, never the value.
func loadApp(getenv func(string) string) (AppConfig, error) {
	c := AppConfig{
		Enabled:           getenv("APP_API") == "true",
		MinVersionIOS:     envOr(getenv, "APP_MIN_VERSION_IOS", defaultAppMinVersion),
		MinVersionAndroid: envOr(getenv, "APP_MIN_VERSION_ANDROID", defaultAppMinVersion),
	}
	if raw := strings.TrimSpace(getenv("APP_WEB_CHECKOUT_REGIONS")); raw != noAppRegions {
		for _, region := range splitCommaList(raw, defaultAppCheckoutRegions) {
			region = strings.ToUpper(region)
			if !appRegionPattern.MatchString(region) {
				return AppConfig{}, fmt.Errorf("config: APP_WEB_CHECKOUT_REGIONS must list two-letter region codes")
			}
			c.WebCheckoutRegions = append(c.WebCheckoutRegions, region)
		}
	}
	for name, v := range map[string]string{"APP_MIN_VERSION_IOS": c.MinVersionIOS, "APP_MIN_VERSION_ANDROID": c.MinVersionAndroid} {
		if !appVersionPattern.MatchString(v) {
			return AppConfig{}, fmt.Errorf("config: %s must be a version like 1.0.0", name)
		}
	}
	return c, nil
}
