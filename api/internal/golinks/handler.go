package golinks

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// PublicPrefix is the two routes the site's Worker calls for every /go/ click. All of that
// traffic arrives from Cloudflare's egress addresses, so the router-wide per-IP limiter would
// throttle a busy link's whole colo as one visitor; these skip it (server.routerLimitSkip)
// and the click route is budgeted per link instead.
const PublicPrefix = "/v1/golinks/"

// clicksPerLinkPerMinute bounds what one link can be counted in a minute, whoever calls: a
// launch-day thread is far below it, and a script hammering the route cannot inflate a link
// past it or turn it into a write load.
const clicksPerLinkPerMinute = 600

// resolveMaxAge is how long a link's target may be cached; a new link is never cached as
// missing for longer, because the Worker caches its own answer for the same minute.
const resolveMaxAge = time.Minute

const maxBody = 1 << 10

type links interface {
	Target(ctx context.Context, slug string) (string, error)
	CountClick(ctx context.Context, slug string, bot bool) (bool, error)
	RecordArrival(ctx context.Context, userID int64, slug string) (bool, error)
	Create(ctx context.Context, l Link) (Link, error)
	Funnel(ctx context.Context) ([]FunnelRow, error)
}

// Service serves the tracked links' routes.
type Service struct {
	Links links
	Log   *slog.Logger
	// Arrivals hears each recorded first link with its click time (the creator program); nil hears none.
	Arrivals ArrivalHook
}

// ArrivalHook hears an account's recorded first link and when the click was.
type ArrivalHook interface {
	LinkArrival(ctx context.Context, userID int64, slug string, clickedAt time.Time) error
}

// ClickedAt is when the kept click happened: the arrival's own `at` (ms since the epoch) when it is
// a moment inside the last ClickWindow, else now. A click time is never in the future.
func ClickedAt(atMillis int64, now time.Time) time.Time {
	if atMillis <= 0 {
		return now
	}
	at := time.UnixMilli(atMillis).UTC()
	if at.After(now) || now.Sub(at) > ClickWindow {
		return now
	}
	return at
}

// Mount registers the links' routes: the Worker's resolve and click, the site's arrival, and
// the owner's list and create.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET "+PublicPrefix+"{slug}", s.resolve)
	perLink := httpx.RateLimitByKey(clicksPerLinkPerMinute, time.Minute, func(r *http.Request) (string, bool) {
		return r.PathValue("slug"), true
	})
	mux.Handle("POST "+PublicPrefix+"{slug}/clicks", perLink(http.HandlerFunc(s.click)))
	mux.HandleFunc("POST /v1/me/arrival", auth.RequireSession(s.arrival))
	mux.HandleFunc("GET /v1/admin/golinks", auth.RequireAdmin(s.list))
	mux.HandleFunc("POST /v1/admin/golinks", auth.RequireAdmin(s.create))
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logger().Error("golinks", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read the links just now", nil)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no link with that name", nil)
}

func (s *Service) resolve(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if !ValidSlug(slug) {
		notFound(w, r)
		return
	}
	target, err := s.Links.Target(r.Context(), slug)
	if errors.Is(err, ErrNotFound) {
		httpx.CachePublic(w, resolveMaxAge, 0)
		notFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, "resolve", err)
		return
	}
	httpx.CachePublic(w, resolveMaxAge, 0)
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"slug": slug, "target": target})
}

// decode reads an optional JSON body: an empty one leaves v as it is.
func decode(w http.ResponseWriter, r *http.Request, v any) error {
	err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(v)
	if errors.Is(err, io.EOF) {
		return nil
	}
	return err
}

func (s *Service) click(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if !ValidSlug(slug) {
		notFound(w, r)
		return
	}
	var in struct {
		Bot bool `json:"bot"`
	}
	if err := decode(w, r, &in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON like {\"bot\": false}", nil)
		return
	}
	found, err := s.Links.CountClick(r.Context(), slug, in.Bot)
	if err != nil {
		s.fail(w, r, "click", err)
		return
	}
	if !found {
		notFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) arrival(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Slug string `json:"slug"`
		At   int64  `json:"at"`
	}
	if err := decode(w, r, &in); err != nil || !ValidSlug(in.Slug) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a link's slug", nil)
		return
	}
	userID := auth.ActorFrom(r.Context()).UserID
	recorded, err := s.Links.RecordArrival(r.Context(), userID, in.Slug)
	if err != nil {
		s.fail(w, r, "arrival", err)
		return
	}
	if recorded && s.Arrivals != nil {
		if err := s.Arrivals.LinkArrival(r.Context(), userID, in.Slug, ClickedAt(in.At, time.Now().UTC())); err != nil {
			s.fail(w, r, "arrival_creator", err)
			return
		}
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]bool{"recorded": recorded})
}

func (s *Service) list(w http.ResponseWriter, r *http.Request) {
	rows, err := s.Links.Funnel(r.Context())
	if err != nil {
		s.fail(w, r, "funnel", err)
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"links": rows})
}

// linkFields is every reason a new link is refused, by field.
func linkFields(l Link) map[string]string {
	fields := map[string]string{}
	if !ValidSlug(l.Slug) {
		fields["slug"] = "lowercase letters, digits and dashes, up to 40"
	}
	if !ValidTarget(l.Target) {
		fields["target"] = "a path on this site, like /guides"
	}
	if !ValidLabel(l.Label) {
		fields["label"] = "a name for the table, up to 80 characters"
	}
	return fields
}

func (s *Service) create(w http.ResponseWriter, r *http.Request) {
	var in Link
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a slug, a target and a label", nil)
		return
	}
	l := Link{Slug: strings.TrimSpace(in.Slug), Target: strings.TrimSpace(in.Target), Label: strings.TrimSpace(in.Label)}
	if fields := linkFields(l); len(fields) > 0 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that link cannot be saved", fields)
		return
	}
	created, err := s.Links.Create(r.Context(), l)
	if errors.Is(err, ErrExists) {
		httpx.WriteError(w, r, http.StatusConflict, "exists", "a link already has that slug",
			map[string]string{"slug": "already taken"})
		return
	}
	if err != nil {
		s.fail(w, r, "create", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, created)
}
