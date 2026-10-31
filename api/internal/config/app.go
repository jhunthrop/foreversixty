package config

import (
	"fmt"
	"regexp"
	"strings"
)

// AppConfig is the mobile apps' API settings. Enabled (APP_API=true) mounts the app routes;
// off by default, so they land dark. The rest is what GET /v1/app/config answers, kept in
// config so a store ruling, a feature switch or a breaking change needs no app update.
type AppConfig struct {
	Enabled bool
	// WebCheckoutRegions are, per app platform ("ios", "android"), the storefront regions
	// (ISO 3166 alpha-2) where that app may link to web checkout (plan decision L):
	// APP_WEB_CHECKOUT_REGIONS_IOS ("US") and APP_WEB_CHECKOUT_REGIONS_ANDROID ("none"),
	// "none" for no region. A platform with no entry links nowhere.
	WebCheckoutRegions map[string][]string
	// MinVersionIOS and MinVersionAndroid are the oldest app builds the API still serves
	// (APP_MIN_VERSION_IOS, APP_MIN_VERSION_ANDROID; "1.0.0").
	MinVersionIOS, MinVersionAndroid string
	// Features are the app feature switches turned on (APP_FEATURES, a comma list of
	// apps/shared/features.json keys; unset is none, and each app keeps its bundled default).
	Features []string
}

// AppPlatforms are the platforms an X-App header names, each with its checkout variable and
// that variable's default.
var AppPlatforms = []struct {
	Name, RegionsEnv, RegionsDefault string
}{
	{"ios", "APP_WEB_CHECKOUT_REGIONS_IOS", "US"},
	{"android", "APP_WEB_CHECKOUT_REGIONS_ANDROID", noAppRegions},
}

const (
	defaultAppMinVersion = "1.0.0"
	noAppRegions         = "none"
)

var (
	appRegionPattern  = regexp.MustCompile(`^[A-Z]{2}$`)
	appVersionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	appFeaturePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
)

// loadApp reads the app settings; errors name the variable, never the value.
func loadApp(getenv func(string) string) (AppConfig, error) {
	c := AppConfig{
		Enabled:            getenv("APP_API") == "true",
		MinVersionIOS:      envOr(getenv, "APP_MIN_VERSION_IOS", defaultAppMinVersion),
		MinVersionAndroid:  envOr(getenv, "APP_MIN_VERSION_ANDROID", defaultAppMinVersion),
		WebCheckoutRegions: map[string][]string{},
	}
	for _, p := range AppPlatforms {
		regions, err := appRegions(p.RegionsEnv, envOr(getenv, p.RegionsEnv, p.RegionsDefault))
		if err != nil {
			return AppConfig{}, err
		}
		c.WebCheckoutRegions[p.Name] = regions
	}
	for name, v := range map[string]string{"APP_MIN_VERSION_IOS": c.MinVersionIOS, "APP_MIN_VERSION_ANDROID": c.MinVersionAndroid} {
		if !appVersionPattern.MatchString(v) {
			return AppConfig{}, fmt.Errorf("config: %s must be a version like 1.0.0", name)
		}
	}
	features, err := appFeatures(getenv("APP_FEATURES"))
	if err != nil {
		return AppConfig{}, err
	}
	c.Features = features
	return c, nil
}

// appRegions parses one platform's region list: "none" is no region, never nil.
func appRegions(name, raw string) ([]string, error) {
	out := []string{}
	if strings.TrimSpace(raw) == noAppRegions {
		return out, nil
	}
	for _, region := range splitCommaList(raw, "") {
		region = strings.ToUpper(region)
		if !appRegionPattern.MatchString(region) {
			return nil, fmt.Errorf("config: %s must list two-letter region codes, or none", name)
		}
		out = append(out, region)
	}
	return out, nil
}

// appFeatures parses APP_FEATURES; unset or empty is no switch turned on.
func appFeatures(raw string) ([]string, error) {
	out := []string{}
	for _, key := range splitCommaList(raw, "") {
		if !appFeaturePattern.MatchString(key) {
			return nil, fmt.Errorf("config: APP_FEATURES must list lower-case feature keys")
		}
		out = append(out, key)
	}
	return out, nil
}
