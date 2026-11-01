// api/internal/guilds/calendar_handler.go
//
// The calendar link (W3-C17, spec 2026-10-31-mobile-wave3 §F16.4 and §10):
//   - GET /v1/me/calendar-feed makes the account's link on first call and answers its URL once;
//     later calls answer url null with the link's created_at (the token is stored hashed).
//   - POST /v1/me/calendar-feed/reset revokes the live link and answers a new one.
//   - GET /v1/calendar/{token}.ics is public by token: the planned raids of the account's guilds.
//
// The token is 256 random bits (auth.NewSessionID), shown once, stored as its SHA-256, never
// logged (httpx.LogPath redacts the route), and budgeted per live token and per address's misses.
package guilds

import (
	"encoding/hex"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	// CalendarFeedPrefix is the public feed's path before the token.
	CalendarFeedPrefix = "/v1/calendar/"
	calendarFeedSuffix = ".ics"
	// calendarReadsPerTokenHour is the spec's budget: 60 reads an hour per link.
	calendarReadsPerTokenHour = 60
	// calendarMissesPerIPHour bounds the links one address may try that do not exist. Only a
	// miss spends it: Google's and Apple's calendar services fetch many players' real links
	// from a few shared addresses, and a budget on every read would refuse them. The token's
	// 256 bits, not this budget, are what make a link unguessable.
	calendarMissesPerIPHour = 30
	// calendarResetsPerHour bounds an account's resets.
	calendarResetsPerHour = 10
	// calendarFeedMaxAge is how long a calendar app may keep a fetched feed (spec §10).
	calendarFeedMaxAge = 5 * time.Minute
)

// calendarToken is the token's form: 32 bytes, base64url without padding.
var calendarToken = regexp.MustCompile(`^[A-Za-z0-9_-]{43}$`)

// CalendarFeedView is the answer of the two account routes. URL is null when the account's
// link was shown before: it is not stored, so only a reset shows a link again.
type CalendarFeedView struct {
	URL       *string   `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

// mountCalendar registers the three calendar routes.
func mountCalendar(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	mux.HandleFunc("GET /v1/me/calendar-feed", auth.RequireSession(s.getCalendarFeed))
	resets := httpx.RateLimitByKey(calendarResetsPerHour, time.Hour, auth.AccountRateLimitKey)
	mux.Handle("POST /v1/me/calendar-feed/reset", resets(auth.RequireSession(s.resetCalendarFeed)))
	limits := calendarLimits{
		perToken: httpx.NewKeyLimiter(calendarReadsPerTokenHour, time.Hour),
		misses:   httpx.NewKeyLimiter(calendarMissesPerIPHour, time.Hour),
		hops:     trustedProxyHops,
	}
	mux.HandleFunc("GET "+CalendarFeedPrefix+"{file}", s.calendarFeed(limits))
}

// calendarFeedURL is the link a calendar app subscribes to.
func (s *Service) calendarFeedURL(token string) string {
	return s.Store.APIBaseURL + CalendarFeedPrefix + token + calendarFeedSuffix
}

func (s *Service) getCalendarFeed(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	token := auth.NewSessionID()
	feed, created, err := s.Store.CreateCalendarFeed(r.Context(), auth.ActorFrom(r.Context()).UserID, auth.TokenHash(token))
	if err != nil {
		s.fail(w, r, "calendar_feed", err, "could not load your calendar link just now")
		return
	}
	view := CalendarFeedView{CreatedAt: feed.CreatedAt}
	if created {
		url := s.calendarFeedURL(token)
		view.URL = &url
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

func (s *Service) resetCalendarFeed(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	token := auth.NewSessionID()
	feed, err := s.Store.ResetCalendarFeed(r.Context(), auth.ActorFrom(r.Context()).UserID, auth.TokenHash(token))
	if err != nil {
		s.fail(w, r, "calendar_feed_reset", err, "could not reset your calendar link just now")
		return
	}
	url := s.calendarFeedURL(token)
	httpx.WriteOK(w, r, http.StatusOK, CalendarFeedView{URL: &url, CreatedAt: feed.CreatedAt})
}

// writeCalendarNotFound is the one answer for a malformed, unknown or revoked link.
func writeCalendarNotFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such calendar", nil)
}

// calendarLimits are the feed's two budgets. perToken is keyed by a live link's hash only, so
// random tokens can never fill it; misses is keyed by address and spent only by a link that
// does not exist.
type calendarLimits struct {
	perToken, misses *httpx.KeyLimiter
	hops             int
}

func writeCalendarRateLimited(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
}

// calendarMiss answers a link that is malformed, unknown or revoked: the one 404, or 429 once
// the address has spent its misses.
func (l calendarLimits) miss(w http.ResponseWriter, r *http.Request) {
	if !l.misses.Allow(httpx.ClientIP(r, l.hops)) {
		writeCalendarRateLimited(w, r)
		return
	}
	writeCalendarNotFound(w, r)
}

// calendarFeed serves the .ics. Every failure to find a live link answers the same 404 (or the
// same 429 once the address has tried too many), and a live link is budgeted per link.
func (s *Service) calendarFeed(limits calendarLimits) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Referrer-Policy", "no-referrer")
		token, ok := strings.CutSuffix(r.PathValue("file"), calendarFeedSuffix)
		if !ok || !calendarToken.MatchString(token) {
			limits.miss(w, r)
			return
		}
		hash := auth.TokenHash(token)
		userID, found, err := s.Store.CalendarFeedUser(r.Context(), hash)
		if err != nil {
			s.fail(w, r, "calendar_ics", err, "could not load that calendar just now")
			return
		}
		if !found {
			limits.miss(w, r)
			return
		}
		if !limits.perToken.Allow(hex.EncodeToString(hash)) {
			writeCalendarRateLimited(w, r)
			return
		}
		now := time.Now()
		raids, err := s.Store.CalendarRaids(r.Context(), userID, now)
		if err != nil {
			s.fail(w, r, "calendar_ics", err, "could not load that calendar just now")
			return
		}
		httpx.CachePrivateFor(w, calendarFeedMaxAge)
		w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
		w.Header().Set("Content-Disposition", `inline; filename="forever-sixty-raids.ics"`)
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		if r.Method != http.MethodHead {
			_, _ = w.Write([]byte(RenderCalendar(raids, s.PublicBaseURL, now)))
		}
	}
}
