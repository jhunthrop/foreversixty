package app

import (
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// configMaxAge is how long GET /v1/app/config may be cached: a ruling or a minimum-version
// change reaches every app within five minutes.
const configMaxAge = 5 * time.Minute

// Config is GET /v1/app/config.
type Config struct {
	// WebCheckoutRegions are the storefront regions (ISO 3166 alpha-2) where an app may link
	// to web checkout; [] for none.
	WebCheckoutRegions []string `json:"web_checkout_regions"`
	// MinAppVersion is the oldest build of each app the API still serves.
	MinAppVersion MinAppVersion `json:"min_app_version"`
	// ActiveBuild is the client build the site and the static data serve.
	ActiveBuild string `json:"active_build"`
}

// MinAppVersion is a minimum version per platform.
type MinAppVersion struct {
	IOS     string `json:"ios"`
	Android string `json:"android"`
}

// config answers GET /v1/app/config; 503 while no client build is loaded.
func (s *Service) config(w http.ResponseWriter, r *http.Request) {
	build, ok := s.Data.ActiveBuild()
	if !ok {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable", "no client build is loaded", nil)
		return
	}
	regions := append([]string{}, s.Settings.WebCheckoutRegions...)
	httpx.CachePublic(w, configMaxAge, 0)
	httpx.WriteOK(w, r, http.StatusOK, Config{
		WebCheckoutRegions: regions, ActiveBuild: build,
		MinAppVersion: MinAppVersion{IOS: s.Settings.MinVersionIOS, Android: s.Settings.MinVersionAndroid},
	})
}
