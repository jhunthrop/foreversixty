// api/internal/rating/handler.go
package rating

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// cacheSeconds matches rankings' own edge-cache policy (spec §5.3: "ratings are exactly as
// volatile as a rankings page").
const cacheSeconds = 30

// liveBoardStale is the Live public boards class's
// stale-while-revalidate window (spec §2.2), matching rankings'.
const liveBoardStale = 300 * time.Second

// Service serves the two rating read routes.
type Service struct {
	Store    *Store
	Reports  *reports.Store
	Accounts Accounts
	// Owners and Entitlements back GET .../history; without both the route is not mounted.
	Owners       Owners
	Entitlements Entitlementer
	Log          *slog.Logger
}

// Owners answers whether an account owns a character. auth.Store satisfies it.
type Owners interface {
	OwnsCharacter(ctx context.Context, userID int64, key string) (bool, error)
}

// Entitlementer answers whether an account may use a feature. entitlements.Store satisfies it.
type Entitlementer interface {
	Can(ctx context.Context, userID int64, feature entitlements.Feature) (bool, entitlements.Reason, error)
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("rating", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// Mount registers the rating read routes.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/reports/{id}/fights/{n}/ratings", s.fightRatings)
	mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}/rating", s.characterRating)
	if s.Owners != nil && s.Entitlements != nil {
		mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}/history", auth.RequireSession(s.characterHistory))
	}
}

func (s *Service) fightRatings(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil || n < 1 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the fight index must be a number, counted from one", nil)
		return
	}
	rep, err := s.Reports.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, reports.ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "report", err, "could not read that report just now")
		return
	}
	if !reports.MayView(r.Context(), auth.ActorFrom(r.Context()), rep, s.Accounts) {
		// The same 404, whether the report does not exist or the caller may not see it -
		// spec's own "no endpoint leaks the existence of a non-public report" rule.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if rep.Flagged != nil {
		// A flagged report (tampered or unverified) is evidence of nothing, so its fight cards
		// are not shown either (presence.EvidenceReport's rule). Its rows stay stored until
		// moderation withdraws them. Asked only once the caller may see the report at all, so
		// it says nothing about a report the caller could not open.
		httpx.WriteError(w, r, http.StatusNotFound, "flagged",
			"that report is flagged, so its fights have no ratings to show", nil)
		return
	}
	rows, ok, err := s.Store.ReadFightRatings(r.Context(), rep.ID, n)
	if err != nil {
		s.fail(w, r, "fight", err, "could not read that fight's ratings just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no ratings for that fight", nil)
		return
	}
	players, err := s.visiblePlayers(r.Context(), rows)
	if err != nil {
		s.fail(w, r, "fight", err, "could not read that fight's ratings just now")
		return
	}
	if rep.Visibility == reports.Public {
		httpx.CachePublic(w, cacheSeconds*time.Second, liveBoardStale)
	} else {
		httpx.CachePrivate(w)
	}
	httpx.WriteOK(w, r, http.StatusOK, fightRatingsDTO{
		FightIndex: n, Kill: rows[0].Kill, KillTimeBand: rows[0].KillTimeBand,
		ModelVersion: rows[0].ModelVersion, Players: players,
	})
}

// visiblePlayers turns a fight's stored rows into the DTOs a caller may actually see,
// dropping any row whose owning account has asked to be anonymized (spec §5.1: an
// anonymized player's row is omitted entirely, never masked or replaced with a
// placeholder).
func (s *Service) visiblePlayers(ctx context.Context, rows []CardRow) ([]playerRatingDTO, error) {
	players := make([]playerRatingDTO, 0, len(rows))
	for _, cr := range rows {
		anon, err := s.Store.anonymized(ctx, cr.PlayerKey)
		if err != nil {
			return nil, fmt.Errorf("rating: check anonymize %s: %w", cr.PlayerKey, err)
		}
		if anon {
			continue
		}
		players = append(players, toPlayerRatingDTO(cr))
	}
	return players, nil
}

func (s *Service) characterRating(w http.ResponseWriter, r *http.Request) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	playerKey := character.Key(region, ruleset, r.PathValue("name"))
	anon, err := s.Store.anonymized(r.Context(), playerKey)
	if err != nil {
		s.fail(w, r, "character", err, "could not read that character's rating just now")
		return
	}
	if anon {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	limit, before := parseTrendQuery(r)
	rows, hasMore, err := s.Store.ReadCharacterRating(r.Context(), playerKey, limit, before)
	if err != nil {
		s.fail(w, r, "character", err, "could not read that character's rating just now")
		return
	}
	if len(rows) == 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no rated fights for that character", nil)
		return
	}
	httpx.CachePublic(w, cacheSeconds*time.Second, liveBoardStale)
	httpx.WriteOK(w, r, http.StatusOK, buildCharacterRatingDTO(playerKey, rows, hasMore))
}

var _ rankings.Scorer = (*Service)(nil)

// ForeverScore is the character page's Forever Score (rankings.Scorer). A character with
// no logs (ErrNoLogs) and one whose account asked to be anonymous both read as no score -
// every field null - never as a zero, matching the rating endpoints' own anonymize rule.
func (s *Service) ForeverScore(ctx context.Context, key string) (rankings.ForeverScore, error) {
	anon, err := s.Store.anonymized(ctx, key)
	if err != nil {
		return rankings.ForeverScore{}, err
	}
	if anon {
		return rankings.ForeverScore{}, nil
	}
	sc, err := s.Store.ScoreV2(ctx, key)
	if errors.Is(err, ErrNoLogs) {
		return rankings.ForeverScore{}, nil
	}
	if err != nil {
		return rankings.ForeverScore{}, err
	}
	return toForeverScoreDTO(sc), nil
}

func parseTrendQuery(r *http.Request) (limit int, before *cursorPos) {
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	if c, ok := decodeTrendCursor(r.URL.Query().Get("cursor")); ok {
		before = &c
	}
	return limit, before
}

const trendCursorSep = "|"

func encodeTrendCursor(c cursorPos) string {
	raw := c.FoughtAt.UTC().Format(time.RFC3339Nano) + trendCursorSep + c.ReportID +
		trendCursorSep + strconv.Itoa(c.FightIndex)
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

// decodeTrendCursor is encodeTrendCursor's inverse. It reports false for anything this
// service did not itself produce - a caller may not walk the trend by any timestamp,
// report and fight index of their own choosing, matching reports/recent.go's own
// decodeRecentCursor contract exactly.
func decodeTrendCursor(s string) (cursorPos, bool) {
	if s == "" {
		return cursorPos{}, false
	}
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return cursorPos{}, false
	}
	parts := strings.SplitN(string(raw), trendCursorSep, 3)
	if len(parts) != 3 {
		return cursorPos{}, false
	}
	t, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil || parts[1] == "" {
		return cursorPos{}, false
	}
	n, err := strconv.Atoi(parts[2])
	if err != nil {
		return cursorPos{}, false
	}
	return cursorPos{FoughtAt: t, ReportID: parts[1], FightIndex: n}, true
}

// characterHistory serves the owner's score and parse history (design/specs/2026-10-24-character-history.md
// section 10). Anyone but the owner, and an anonymised character, read as no such character.
func (s *Service) characterHistory(w http.ResponseWriter, r *http.Request) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	playerKey := character.Key(region, ruleset, r.PathValue("name"))
	actor := auth.ActorFrom(r.Context())
	owns, err := s.Owners.OwnsCharacter(r.Context(), actor.UserID, playerKey)
	if err != nil {
		s.fail(w, r, "history", err, "could not read that character's history just now")
		return
	}
	anon := false
	if owns {
		if anon, err = s.Store.anonymized(r.Context(), playerKey); err != nil {
			s.fail(w, r, "history", err, "could not read that character's history just now")
			return
		}
	}
	if !owns || anon {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	allowed, _, err := s.Entitlements.Can(r.Context(), actor.UserID, entitlements.FeatureHistory)
	if err != nil {
		s.fail(w, r, "history", err, "could not check your account just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusPaymentRequired, "premium_required",
			"character history is a premium feature", nil)
		return
	}
	h, err := s.Store.ReadOwnerHistory(r.Context(), playerKey)
	if err != nil {
		s.fail(w, r, "history", err, "could not read that character's history just now")
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, toHistoryDTO(h))
}
