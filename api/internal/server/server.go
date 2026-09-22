package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/site"
	"github.com/jhunthrop/foreversixty/api/internal/subscribe"
)

// routerRequestsPerMinute is the router-wide per-IP request budget.
const routerRequestsPerMinute = 120

type Deps struct {
	Version       string
	Log           *slog.Logger
	AllowedOrigin string
	Subscribe     *subscribe.Service
	Builds        *builds.Service
	Site          *site.Deps
	// Auth resolves the request's identity and enforces CSRF. Without
	// it none of the Phase 3 routes are mounted, because every one of
	// them is about who is asking.
	Auth     *auth.Authenticator
	Accounts *auth.Service
	Reports  *reports.Service
	Ingest   *reports.Ingest
	Uploads  *reports.Uploads
	Rankings *rankings.Service
	Addon    *addon.Service
	Guilds   *guilds.Service
	Sims     *sims.Service
	Rating   *rating.Service

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
	// The phase boundaries, for any client that has to bucket a date the
	// same way the rankings do. Four fixed instants that change only with
	// a deploy, so an hour at the edge is safe.
	mux.HandleFunc("GET /v1/phases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=3600")
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{"phases": phase.Boundaries})
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
	if d.Accounts != nil {
		auth.Mount(mux, d.Accounts, d.TrustedProxyHops)
	}
	if d.Reports != nil {
		reports.Mount(mux, d.Reports)
	}
	if d.Ingest != nil {
		reports.MountIngest(mux, d.Ingest)
	}
	if d.Uploads != nil {
		reports.MountUploads(mux, d.Uploads)
	}
	if d.Rankings != nil {
		rankings.Mount(mux, d.Rankings)
	}
	if d.Addon != nil {
		addon.Mount(mux, d.Addon)
	}
	if d.Guilds != nil {
		guilds.Mount(mux, d.Guilds, d.TrustedProxyHops)
	}
	if d.Sims != nil {
		sims.Mount(mux, d.Sims)
	}
	if d.Rating != nil {
		rating.Mount(mux, d.Rating)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
	})
	mws := []func(http.Handler) http.Handler{
		httpx.RequestID(), httpx.Recover(d.Log), httpx.Logger(d.Log), httpx.CORS(d.AllowedOrigin),
		httpx.RateLimitExcept(routerRequestsPerMinute, d.TrustedProxyHops, sharedRead),
	}
	if d.Auth != nil {
		mws = append(mws, d.Auth.Middleware)
	}
	return httpx.Chain(mux, mws...)
}

// sharedRead reports whether r is a read of the shared, cached surface -
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
//
// Phase 3 adds the report card under /reports/ on the same terms: it is
// fetched by unfurlers, not by people, it is a rendered image with no
// private content, and it is cached for five minutes.
func sharedRead(r *http.Request) bool {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return false
	}
	return strings.HasPrefix(r.URL.Path, "/b/") || strings.HasPrefix(r.URL.Path, "/reports/")
}
