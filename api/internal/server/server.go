package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/site"
	"github.com/jhunthrop/foreversixty/api/internal/subscribe"
)

// routerRequestsPerMinute is the router-wide per-IP request budget.
const routerRequestsPerMinute = 120

type Deps struct {
	Version          string
	Log              *slog.Logger
	AllowedOrigin    string
	Subscribe        *subscribe.Service
	Builds           *builds.Service
	Site             *site.Deps
	TrustedProxyHops int
}

func NewRouter(d Deps) http.Handler {
	if d.Log == nil {
		d.Log = slog.Default()
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"version": d.Version})
	})
	if d.Subscribe != nil {
		subscribe.Mount(mux, d.Subscribe, d.Log, d.TrustedProxyHops)
	}
	if d.Builds != nil {
		builds.Mount(mux, d.Builds, d.TrustedProxyHops)
	}
	if d.Site != nil {
		site.Mount(mux, *d.Site)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
	})
	return httpx.Chain(mux, httpx.RequestID(), httpx.Recover(d.Log), httpx.Logger(d.Log), httpx.CORS(d.AllowedOrigin), httpx.RateLimitExcept(routerRequestsPerMinute, d.TrustedProxyHops, sharedBuildRead))
}

// sharedBuildRead reports whether r is a read of the shared build surface -
// GET (or HEAD, which net/http serves from the same handlers) under /b/,
// i.e. the build page and its preview card. Those two routes, and only
// those, skip the router-wide per-IP limiter.
//
// Why: foreversixty.gg/b/* is proxied here by the site's Cloudflare Worker,
// so with TRUSTED_PROXY_HOPS=1 the address this service resolves as the
// client is the Worker's egress IP, not the visitor's. Every visitor coming
// through one Cloudflare colo would otherwise share a single
// routerRequestsPerMinute bucket - roughly two requests a second for the
// whole shared-build surface, which is the traffic the feature exists to
// serve. Both routes are read-only, content-hash-addressed and CDN-cached,
// so the limiter buys little there.
//
// What it costs: the build pages and cards become an unauthenticated read
// surface with no per-IP ceiling and so can be scraped, bounded only by
// Cloud Run's max instance count. That is an accepted trade. Writes are
// untouched: the planner calls POST /v1/builds directly against this
// service rather than through the Worker, and that route keeps both this
// limiter and its own 20-saves-per-IP-per-hour one.
func sharedBuildRead(r *http.Request) bool {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return false
	}
	return strings.HasPrefix(r.URL.Path, "/b/")
}
