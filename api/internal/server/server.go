package server

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/livesurface"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/site"
	"github.com/jhunthrop/foreversixty/api/internal/subscribe"
)

// routerRequestsPerMinute is the router-wide per-IP request budget.
const routerRequestsPerMinute = 120

// deviceRequestsPerMinute is each device token's own budget, so a companion's
// traffic is never counted against the browser on the same IP.
// deadTokenRequestsPerMinute is what a token already answered 401 is cut to.
const (
	deviceRequestsPerMinute    = 120
	deadTokenRequestsPerMinute = 10
)

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
	Notify   *notify.Service
	Rating   *rating.Service
	Billing  *billing.Service
	// Seats serves the Premium seat routes; nil leaves them unmounted.
	Seats *seats.Service
	// DiscordBot serves the signed interactions endpoint; nil leaves it unmounted.
	DiscordBot *discordbot.Service
	// Coach serves the log coach's routes; nil leaves them unmounted.
	Coach *coach.Service
	// Tiles serves the Discord panel's item tiles; nil leaves the route unmounted.
	Tiles *sheet.Tiles
	// LiveSurface serves the apps' widget snapshot and live registrations; nil (APP_WIDGETS
	// off) leaves them unmounted.
	LiveSurface *livesurface.Service

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
		httpx.CachePublic(w, time.Hour, 0)
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
		addon.Mount(mux, d.Addon, d.TrustedProxyHops)
	}
	if d.Guilds != nil {
		guilds.Mount(mux, d.Guilds, d.TrustedProxyHops)
	}
	if d.Sims != nil {
		sims.Mount(mux, d.Sims)
	}
	if d.Notify != nil {
		notify.Mount(mux, d.Notify)
	}
	if d.Rating != nil {
		rating.Mount(mux, d.Rating)
	}
	if d.Billing != nil {
		billing.Mount(mux, d.Billing, d.TrustedProxyHops)
	}
	if d.Seats != nil {
		seats.Mount(mux, d.Seats, d.TrustedProxyHops)
	}
	if d.DiscordBot != nil {
		discordbot.Mount(mux, d.DiscordBot, d.TrustedProxyHops)
	}
	if d.Coach != nil {
		coach.Mount(mux, d.Coach)
		coach.MountLines(mux, d.Coach)
	}
	if d.Tiles != nil {
		d.Tiles.Mount(mux)
	}
	if d.LiveSurface != nil {
		livesurface.Mount(mux, d.LiveSurface)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
	})
	mws := []func(http.Handler) http.Handler{
		httpx.RequestID(), httpx.Recover(d.Log), httpx.Logger(d.Log), httpx.CORS(d.AllowedOrigin),
		httpx.RateLimitByCredential(routerRequestsPerMinute, deviceRequestsPerMinute, deadTokenRequestsPerMinute, d.TrustedProxyHops, routerLimitSkip(d.DiscordBot != nil, d.Tiles != nil)),
	}
	if d.Auth != nil {
		mws = append(mws, d.Auth.Middleware)
	}
	return httpx.Chain(mux, mws...)
}

// routerLimitSkip is which requests skip the router-wide limiter: the shared
// reads, and Discord's interactions while the bot is mounted. Discord posts
// every interaction from a few egress IPs, so one IP bucket would 429 a busy
// bot; the endpoint's Ed25519 signature check (cheap, before any work) is what
// guards it. With the bot off the path is a 404 like any other, and limited.
//
// The Discord panel's item tiles skip it too, while mounted: Discord's media proxy fetches every
// thumbnail from a few IPs, so one bucket would turn a busy channel's icons into broken
// squares. What bounds them instead: the route answers only a loaded build and a numeric id,
// every item the build lacks shares one cached question-mark tile (so random ids cannot churn
// the LRU or force renders), and the drawn tiles are immutable and cached by the proxy.
func routerLimitSkip(botMounted, tilesMounted bool) func(*http.Request) bool {
	return func(r *http.Request) bool {
		return sharedRead(r) ||
			(botMounted && r.Method == http.MethodPost && r.URL.Path == discordbot.InteractionsPath) ||
			(tilesMounted && (r.Method == http.MethodGet || r.Method == http.MethodHead) && strings.HasPrefix(r.URL.Path, sheet.TilePath))
	}
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
