// api/internal/rating/handler.go
package rating

import (
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// cacheSeconds matches rankings' own edge-cache policy (spec §5.3: "ratings are exactly as
// volatile as a rankings page").
const cacheSeconds = 30

// Service serves the two rating read routes.
type Service struct {
	Store    *Store
	Reports  *reports.Store
	Accounts Accounts
	Log      *slog.Logger
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
	if !visible(r, rep, s.Accounts) {
		// The same 404, whether the report does not exist or the caller may not see it -
		// spec's own "no endpoint leaks the existence of a non-public report" rule.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
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
	players := make([]playerRatingDTO, 0, len(rows))
	for _, cr := range rows {
		anon, err := s.Store.anonymized(r.Context(), cr.PlayerKey)
		if err != nil {
			s.fail(w, r, "fight", err, "could not read that fight's ratings just now")
			return
		}
		if anon {
			continue // spec §5.1: an anonymized player's row is omitted entirely, no placeholder
		}
		players = append(players, toPlayerRatingDTO(cr))
	}
	if rep.Visibility == reports.Public {
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheSeconds))
	} else {
		w.Header().Set("Cache-Control", "private")
	}
	httpx.WriteOK(w, r, http.StatusOK, fightRatingsDTO{
		FightIndex: n, Kill: rows[0].Kill, KillTimeBand: rows[0].KillTimeBand,
		ModelVersion: rows[0].ModelVersion, Players: players,
	})
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
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheSeconds))
	httpx.WriteOK(w, r, http.StatusOK, buildCharacterRatingDTO(playerKey, rows, hasMore))
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
