package app

import (
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// configMaxAge is how long GET /v1/app/config may be cached: a ruling or a minimum-version
// change reaches every app within five minutes.
const configMaxAge = 5 * time.Minute

// AppHeader names the app build that sends a request: "<platform>/<version> (<build>)".
const AppHeader = "X-App"

// Config is GET /v1/app/config.
type Config struct {
	// WebCheckoutRegions are the storefront regions (ISO 3166 alpha-2) where the calling app
	// may link to web checkout; [] for none.
	WebCheckoutRegions []string `json:"web_checkout_regions"`
	// MinAppVersion is the oldest build of each app the API still serves.
	MinAppVersion MinAppVersion `json:"min_app_version"`
	// ActiveBuild is the client build the site and the static data serve.
	ActiveBuild string `json:"active_build"`
	// Features are the feature switches by key; {} when none is turned on.
	Features map[string]bool `json:"features"`
}

// MinAppVersion is a minimum version per platform.
type MinAppVersion struct {
	IOS     string `json:"ios"`
	Android string `json:"android"`
}

// AppPlatform is the platform an X-App header names ("ios/1.2.0 (42)" is "ios"), or "" for a
// missing or malformed header.
func AppPlatform(header string) string {
	platform, rest, ok := strings.Cut(strings.TrimSpace(header), "/")
	if !ok || rest == "" {
		return ""
	}
	return strings.ToLower(platform)
}

// regionsFor is the checkout regions of one platform, never nil; an unknown platform has none.
func (s Settings) regionsFor(platform string) []string {
	return append([]string{}, s.WebCheckoutRegions[platform]...)
}

// featureMap is the switches turned on, as the answer's map.
func (s Settings) featureMap() map[string]bool {
	out := make(map[string]bool, len(s.Features))
	for _, key := range s.Features {
		out[key] = true
	}
	return out
}

// config answers GET /v1/app/config; 503 while no client build is loaded. The checkout
// regions depend on the calling platform, so the answer varies on X-App, and it is private:
// the CDN in front of the API does not key on X-App and would hand one platform's answer
// to the other.
func (s *Service) config(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Vary", AppHeader)
	build, ok := s.Data.ActiveBuild()
	if !ok {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable", "no client build is loaded", nil)
		return
	}
	httpx.CachePrivateFor(w, configMaxAge)
	httpx.WriteOK(w, r, http.StatusOK, Config{
		WebCheckoutRegions: s.Settings.regionsFor(AppPlatform(r.Header.Get(AppHeader))),
		ActiveBuild:        build,
		MinAppVersion:      MinAppVersion{IOS: s.Settings.MinVersionIOS, Android: s.Settings.MinVersionAndroid},
		Features:           s.Settings.featureMap(),
	})
}
