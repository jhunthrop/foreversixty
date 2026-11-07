package activity

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The anonymous feed's cache: the "live board" class every public feed uses.
const (
	feedMaxAge    = 30 * time.Second
	feedStale     = 300 * time.Second
	feedCacheSize = 512
)

// GuildLister lists an account's guilds; auth.Store satisfies it.
type GuildLister interface {
	Guilds(ctx context.Context, userID int64) ([]auth.Guild, error)
}

// Service serves GET /v1/feed and GET /v1/home.
type Service struct {
	Store  *Store
	Guilds GuildLister
	// Home answers GET /v1/home; nil leaves the route unmounted.
	Home *Home
	Log  *slog.Logger
	Now  func() time.Time

	cache *pageCache
}

// Mount registers the routes. GET /v1/feed is public: the router-wide per-IP budget limits
// it like every other public read.
func Mount(mux *http.ServeMux, s *Service) {
	s.cache = newPageCache(feedMaxAge, feedCacheSize)
	mux.HandleFunc("GET /v1/feed", s.feed)
	if s.Home != nil {
		mux.HandleFunc("GET /v1/home", auth.RequireSession(s.home))
	}
}

func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now().UTC()
	}
	return time.Now().UTC()
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) feed(w http.ResponseWriter, r *http.Request) {
	q, ok := parseFeedQuery(w, r)
	if !ok {
		return
	}
	q.Now = s.now()
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() {
		s.anonymousFeed(w, r, q)
		return
	}
	ids, err := s.verifiedGuilds(r.Context(), actor.UserID)
	if err != nil {
		s.fail(w, r, "guilds", err)
		return
	}
	q.GuildIDs = ids
	page, err := s.Store.Read(r.Context(), q)
	if err != nil {
		s.readFailed(w, r, err)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, page)
}

func (s *Service) anonymousFeed(w http.ResponseWriter, r *http.Request, q Query) {
	if q.Filter == FilterGuild {
		httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized", "sign in to read your guild's feed", nil)
		return
	}
	q.GuildIDs = []int64{}
	key := r.URL.RawQuery
	page, hit := s.cache.get(key, q.Now)
	if !hit {
		var err error
		if page, err = s.Store.Read(r.Context(), q); err != nil {
			s.readFailed(w, r, err)
			return
		}
		s.cache.put(key, page, q.Now)
	}
	httpx.CachePublic(w, feedMaxAge, feedStale)
	w.Header().Set("Vary", "Cookie, Authorization")
	httpx.WriteOK(w, r, http.StatusOK, page)
}

func (s *Service) verifiedGuilds(ctx context.Context, userID int64) ([]int64, error) {
	ids := []int64{}
	if s.Guilds == nil {
		return ids, nil
	}
	gs, err := s.Guilds.Guilds(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, g := range gs {
		if g.Verified {
			ids = append(ids, g.ID)
		}
	}
	return ids, nil
}

// parseFeedQuery validates the query string, answering 400 itself on a bad value.
func parseFeedQuery(w http.ResponseWriter, r *http.Request) (Query, bool) {
	v := r.URL.Query()
	q := Query{Filter: v.Get("filter"), Spec: strings.ToLower(v.Get("spec")), Cursor: v.Get("cursor")}
	if q.Filter == "" {
		q.Filter = FilterAll
	}
	bad := func(field, msg string) (Query, bool) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msg, map[string]string{field: msg})
		return Query{}, false
	}
	switch q.Filter {
	case FilterAll, FilterSpec, FilterGuild, FilterLogs:
	default:
		return bad("filter", "filter must be all, spec, guild or logs")
	}
	if _, ok := specs.ByKey[q.Spec]; q.Spec != "" && !ok {
		return bad("spec", "spec must be a spec key like warrior-fury")
	}
	if b := v.Get("band"); b != "" {
		n, err := strconv.Atoi(b)
		if err != nil || !ValidBand(n) {
			return bad("band", "band must be 20, 30, 40, 50 or 60")
		}
		q.Band = n
	}
	if c := strings.ToLower(v.Get("character")); c != "" {
		if !character.ValidKey(c) {
			return bad("character", "character must be a key like us/normal/thalrin")
		}
		q.Character = c
	}
	if q.Filter == FilterSpec && (q.Spec == "" || q.Band == 0) {
		return bad("spec", "the spec filter needs spec and band")
	}
	if l := v.Get("limit"); l != "" {
		n, err := strconv.Atoi(l)
		if err != nil || n < 1 || n > MaxLimit {
			return bad("limit", "limit must be 1 to 30")
		}
		q.Limit = n
	}
	return q, true
}

func (s *Service) readFailed(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, ErrBadCursor) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "cursor is not one this feed wrote", map[string]string{"cursor": "invalid"})
		return
	}
	s.fail(w, r, "read", err)
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logger().Error("activity", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read the feed just now", nil)
}
