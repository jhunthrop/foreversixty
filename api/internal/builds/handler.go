package builds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

const (
	// maxBodyBytes is the contract's 8 KB cap on a save body.
	maxBodyBytes = 8 << 10
	// savesPerHour is the contract's per-IP save budget.
	savesPerHour = 20
	// fetchMaxAge is the contract's cache lifetime for a fetched record.
	fetchMaxAge = 24 * time.Hour
)

// Storer is the part of Store the handlers use, so they can be tested
// without Postgres.
type Storer interface {
	Save(ctx context.Context, b Build, userID *int64) (Build, bool, error)
	Get(ctx context.Context, id string) (Build, error)
	Mine(ctx context.Context, userID int64, page int) (Page, error)
}

// Service serves the JSON build endpoints.
type Service struct {
	Store         Storer
	Data          *trees.Data
	PublicBaseURL string
	Log           *slog.Logger
}

// Mount registers POST /v1/builds, GET /v1/builds/{id}, and
// GET /v1/builds?mine=1. Saves carry their own hourly per-IP limiter on
// top of the router-wide per-minute one.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	limited := httpx.RateLimitPer(savesPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/builds", limited(http.HandlerFunc(s.save)))
	mux.HandleFunc("GET /v1/builds/{id}", s.fetch)
	mux.HandleFunc("GET /v1/builds", auth.RequireSession(s.mine))
}

func (s *Service) save(w http.ResponseWriter, r *http.Request) {
	var in Input
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBodyBytes)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON of at most 8 KB", nil)
		return
	}
	if fields := Validate(s.Data, in); fields != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this build breaks the talent rules", fields)
		return
	}
	b, err := New(in)
	if err != nil {
		s.logger().Error("builds", "id", httpx.RequestIDFrom(r.Context()), "op", "save", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that build", nil)
		return
	}
	var owner *int64
	if a := auth.ActorFrom(r.Context()); a.Signed() {
		id := a.UserID
		owner = &id
	}
	stored, created, err := s.Store.Save(r.Context(), b, owner)
	if err != nil {
		s.logger().Error("builds", "id", httpx.RequestIDFrom(r.Context()), "op", "save", "build", b.ID, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that build", nil)
		return
	}
	status := http.StatusOK
	if created {
		status = http.StatusCreated
	}
	httpx.WriteOK(w, r, status, map[string]string{
		"id":  stored.ID,
		"url": s.PublicBaseURL + "/b/" + stored.ID,
	})
}

func (s *Service) fetch(w http.ResponseWriter, r *http.Request) {
	b, err := s.Store.Get(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such build", nil)
	case err != nil:
		s.logger().Error("builds", "id", httpx.RequestIDFrom(r.Context()), "op", "fetch", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not load that build", nil)
	default:
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", int(fetchMaxAge.Seconds())))
		httpx.WriteOK(w, r, http.StatusOK, b)
	}
}

// mine is the signed-in player's own builds. Like the sim history, the
// parameter is required so the route's meaning is on the URL: there is
// no "everyone's builds" list and inventing one by omission would be a
// surprise.
func (s *Service) mine(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("mine") != "1" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "this list is mine=1 only",
			map[string]string{"mine": "1"})
		return
	}
	page := 1
	if v := r.URL.Query().Get("page"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "page must be 1 or more",
				map[string]string{"page": "a page number from 1"})
			return
		}
		page = n
	}
	out, err := s.Store.Mine(r.Context(), auth.ActorFrom(r.Context()).UserID, page)
	if err != nil {
		s.logger().Error("builds", "id", httpx.RequestIDFrom(r.Context()), "op", "mine", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal",
			"could not read your builds just now", nil)
		return
	}
	// Per-account: never cached at a shared edge.
	w.Header().Set("Cache-Control", "private, no-store")
	httpx.WriteOK(w, r, http.StatusOK, out)
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}
