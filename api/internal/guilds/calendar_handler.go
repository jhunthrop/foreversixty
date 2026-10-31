// api/internal/guilds/calendar_handler.go
//
// The calendar link (W3-C17, spec 2026-10-31-mobile-wave3 §F16.4 and §10):
//   - GET /v1/me/calendar-feed makes the account's link on first call and answers its URL once;
//     later calls answer url null with the link's created_at (the token is stored hashed).
//   - POST /v1/me/calendar-feed/reset revokes the live link and answers a new one.
//   - GET /v1/calendar/{token}.ics is public by token: the planned raids of the account's guilds.
//
// The token is 256 random bits (auth.NewSessionID), shown once, stored as its SHA-256, never
// logged (httpx.LogPath redacts the route), and budgeted per token and per address.
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
	// calendarReadsPerIPHour bounds one address across links. Calendar services fetch many
	// players' links from a few addresses, so it is looser than the per-link budget; the
	// token's 256 bits, not this budget, are what make a link unguessable.
	calendarReadsPerIPHour = 600
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
	perIP := httpx.RateLimitPer(calendarReadsPerIPHour, time.Hour, trustedProxyHops)
	perToken := httpx.NewKeyLimiter(calendarReadsPerTokenHour, time.Hour)
	mux.Handle("GET "+CalendarFeedPrefix+"{file}", perIP(s.calendarFeed(perToken)))
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

// calendarFeed serves the .ics. The order matters: the form check and the per-link budget run
// before any database read, and every failure to find a live link answers the same 404.
func (s *Service) calendarFeed(perToken *httpx.KeyLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Referrer-Policy", "no-referrer")
		token, ok := strings.CutSuffix(r.PathValue("file"), calendarFeedSuffix)
		if !ok || !calendarToken.MatchString(token) {
			writeCalendarNotFound(w, r)
			return
		}
		hash := auth.TokenHash(token)
		if !perToken.Allow(hex.EncodeToString(hash)) {
			httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
			return
		}
		userID, found, err := s.Store.CalendarFeedUser(r.Context(), hash)
		if err != nil {
			s.fail(w, r, "calendar_ics", err, "could not load that calendar just now")
			return
		}
		if !found {
			writeCalendarNotFound(w, r)
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
