// Package livesurface serves the mobile apps' home-screen and lock-screen surfaces (spec
// design/specs/2026-10-31-mobile-wave3.md §5.6, contracts W3-C12 to W3-C14):
//
//   - GET /v1/me/widget, the one snapshot every widget, lock-screen widget, Live Activity and
//     ongoing notification draws from;
//   - POST /v1/me/live-activities and DELETE /v1/me/live-activities/{id}, which keep the push
//     tokens of a raid countdown or a server sim's progress;
//   - the push payloads those surfaces are updated with (payloads.go).
//
// The routes are mounted only with APP_WIDGETS=true (config.AppWidgets), so they land dark.
// Every figure comes from a package that already owns it: the raid from guilds (the raid
// page's calendar and sign-ups), the upgrades from the Home answer, the runs from sims.
package livesurface

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Accounts reads the caller's account and characters (auth.Store).
type Accounts interface {
	User(ctx context.Context, id int64) (auth.User, error)
	Characters(ctx context.Context, userID int64) ([]auth.Character, error)
}

// Raids reads the raid half of the snapshot and the raid a countdown follows (guilds.Service).
type Raids interface {
	WidgetRaidFor(ctx context.Context, userID int64, characterKey string, now time.Time) (guilds.WidgetRaid, error)
	LiveRaidEnd(ctx context.Context, userID, raidID int64) (time.Time, bool, error)
}

// Upgrades is the Home answer's top upgrades for a character (lane A3's builder). nil on the
// Service leaves the snapshot's upgrades null.
type Upgrades interface {
	WidgetUpgrades(ctx context.Context, characterKey string) (*WidgetUpgrades, error)
}

// Sims reads the caller's server runs (sims.Store).
type Sims interface {
	RunningServerRuns(ctx context.Context, userID int64) (int, error)
	OwnsServerRun(ctx context.Context, userID int64, id string) (bool, error)
}

// Entitlements answers whether the caller holds a Premium feature (entitlements.Store).
type Entitlements interface {
	Can(ctx context.Context, userID int64, feature entitlements.Feature) (bool, entitlements.Reason, error)
}

// Service serves the live-surface routes.
type Service struct {
	Accounts     Accounts
	Raids        Raids
	Upgrades     Upgrades
	Sims         Sims
	Entitlements Entitlements
	Tokens       *Store
	Log          *slog.Logger
	// Now is the clock; nil reads time.Now.
	Now func() time.Time
	// RegisterLimiter is the per-account registration budget; nil uses the default.
	RegisterLimiter *httpx.KeyLimiter
}

// registrationsPerMinute bounds how often one account may register a live surface.
const registrationsPerMinute = 30

// Mount registers the routes. Each takes a browser session or an app's device token.
func Mount(mux *http.ServeMux, s *Service) {
	if s.RegisterLimiter == nil {
		s.RegisterLimiter = httpx.NewKeyLimiter(registrationsPerMinute, time.Minute)
	}
	mux.HandleFunc("GET /v1/me/widget", auth.RequireSession(s.widget))
	mux.HandleFunc("POST /v1/me/live-activities", auth.RequireSession(s.register))
	mux.HandleFunc("DELETE /v1/me/live-activities/{id}", auth.RequireSession(s.unregister))
}

func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("livesurface", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}
