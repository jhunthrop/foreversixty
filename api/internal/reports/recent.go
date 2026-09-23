package reports

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RecentPerPage is the page size of the public "recent reports" feed.
const RecentPerPage = 10

// recentCursorSep separates the two halves a recent cursor encodes.
const recentCursorSep = "|"

// RecentSummary is one row of GET /v1/reports/recent: enough to link and
// place a public report, and nothing that could name its owner. The
// owner is never read for this feed at all - unlike the report page's
// own Owner field (see view, which reads u.PublicName so an anonymized
// owner's battletag never reaches even that response) - so there is no
// anonymize check to apply here: the row simply carries no owner field.
type RecentSummary struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	FightCount int       `json:"fight_count"`
	KillCount  int       `json:"kill_count"`
	// GuildName is empty when the report has no guild.
	GuildName string `json:"guild_name,omitempty"`
}

// RecentPage is the body of GET /v1/reports/recent.
type RecentPage struct {
	Rows []RecentSummary `json:"rows"`
	// NextCursor is set only when a further page may exist - a full
	// page came back - and is passed back as ?cursor= to read it.
	NextCursor string `json:"next_cursor,omitempty"`
}

// recentCursor is the keyset position GET /v1/reports/recent pages
// from: the (created_at, id) pair of the last row a page returned.
type recentCursor struct {
	CreatedAt time.Time
	ID        string
}

// encodeRecentCursor renders the position after row (createdAt, id) as
// the opaque token the contract hands back as next_cursor.
func encodeRecentCursor(createdAt time.Time, id string) string {
	raw := createdAt.UTC().Format(time.RFC3339Nano) + recentCursorSep + id
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

// decodeRecentCursor is encodeRecentCursor's inverse. It reports false
// for anything that is not a cursor this service produced - a caller
// may not walk the feed by any timestamp and id of their own choosing.
func decodeRecentCursor(s string) (recentCursor, bool) {
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return recentCursor{}, false
	}
	createdAt, id, ok := strings.Cut(string(raw), recentCursorSep)
	if !ok || id == "" {
		return recentCursor{}, false
	}
	t, err := time.Parse(time.RFC3339Nano, createdAt)
	if err != nil {
		return recentCursor{}, false
	}
	return recentCursor{CreatedAt: t, ID: id}, true
}

// recent serves the public, unauthenticated feed of the newest complete
// public reports. It has no route-specific limiter of its own: the
// router-wide per-IP budget (see server.NewRouter) is the same one
// every other public read - a report page, a rankings page - answers
// to.
func (s *Service) recent(w http.ResponseWriter, r *http.Request) {
	var before *recentCursor
	if v := r.URL.Query().Get("cursor"); v != "" {
		c, ok := decodeRecentCursor(v)
		if !ok {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a cursor",
				map[string]string{"cursor": "the next_cursor a previous page returned"})
			return
		}
		before = &c
	}
	rows, err := s.Store.Recent(r.Context(), before, RecentPerPage)
	if err != nil {
		s.fail(w, r, "recent", err, "could not read the recent reports just now")
		return
	}
	page := RecentPage{Rows: rows}
	if len(rows) == RecentPerPage {
		last := rows[len(rows)-1]
		page.NextCursor = encodeRecentCursor(last.CreatedAt, last.ID)
	}
	// Public feed, the Live public boards class (spec §2.2), the same as rankings: the
	// browser and the CDN reuse it across the home page and /logs instead of a fresh
	// query per island per load.
	httpx.CachePublic(w, reportsFeedMaxAge, reportsFeedStale)
	httpx.WriteOK(w, r, http.StatusOK, page)
}
