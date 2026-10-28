package coach

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// historyPageSize is how many stored reviews one history page (and the recent line) carries.
const historyPageSize = 20

// historySQL is the newest stored review of each fight of the characters in $1, newest
// first, past the keyset cursor ($2 generated_at, $3 report, $4 fight; a null $2 is the
// first page), with the pack's spec, boss and kill.
const historySQL = `
with newest as (
  select distinct on (r.report_id, r.fight_index, r.player_key)
         r.report_id, r.fight_index, r.player_key, r.pack_version, r.source, r.generated_at,
         r.review->'fixFirst'->>'title' as title
    from coach_reviews r
   where r.player_key = any($1)
   order by r.report_id, r.fight_index, r.player_key, r.generated_at desc
)
select n.report_id, n.fight_index, n.player_key, n.source, n.generated_at, coalesce(n.title, ''),
       coalesce(p.pack->'context'->>'spec', ''), coalesce(p.pack->'context'->>'boss', ''),
       coalesce((p.pack->'context'->>'kill')::boolean, true)
  from newest n
  left join coach_packs p on p.report_id = n.report_id and p.fight_index = n.fight_index
   and p.player_key = n.player_key and p.pack_version = n.pack_version
 where $2::timestamptz is null or (n.generated_at, n.report_id, n.fight_index) < ($2, $3::text, $4::int)
 order by n.generated_at desc, n.report_id desc, n.fight_index desc
 limit $5`

// historyItem is one stored review in the history and the recent line (§10.3).
type historyItem struct {
	ReportID    string    `json:"report_id"`
	FightIndex  int       `json:"fight_index"`
	PlayerKey   string    `json:"player_key"`
	Spec        string    `json:"spec"`
	Boss        string    `json:"boss"`
	Kill        bool      `json:"kill"`
	Source      string    `json:"source"`
	GeneratedAt time.Time `json:"generated_at"`
	FixFirst    *fixTitle `json:"fix_first"`
	cursor      historyKey
}

// fixTitle is a history item's fix-first: its title only.
type fixTitle struct {
	Title string `json:"title"`
}

// historyKey is the keyset cursor: (generated_at, report_id, fight_index).
type historyKey struct {
	At         time.Time
	ReportID   string
	FightIndex int
}

// historyPage is the history route's body.
type historyPage struct {
	Reviews    []historyItem `json:"reviews"`
	NextCursor string        `json:"next_cursor"`
}

// storedReviews is a page of keys' stored reviews after cursor (nil: the first page).
func (s *Service) storedReviews(ctx context.Context, keys []string, after *historyKey, limit int) ([]historyItem, error) {
	var at *time.Time
	var reportID string
	var fight int
	if after != nil {
		at, reportID, fight = &after.At, after.ReportID, after.FightIndex
	}
	rows, err := s.Store.DB.Query(ctx, historySQL, keys, at, reportID, fight, limit)
	if err != nil {
		return nil, fmt.Errorf("coach: stored reviews: %w", err)
	}
	defer rows.Close()
	out := []historyItem{}
	for rows.Next() {
		var it historyItem
		var title string
		if err := rows.Scan(&it.ReportID, &it.FightIndex, &it.PlayerKey, &it.Source, &it.GeneratedAt,
			&title, &it.Spec, &it.Boss, &it.Kill); err != nil {
			return nil, fmt.Errorf("coach: stored reviews: %w", err)
		}
		if title != "" {
			it.FixFirst = &fixTitle{Title: title}
		}
		it.cursor = historyKey{At: it.GeneratedAt, ReportID: it.ReportID, FightIndex: it.FightIndex}
		out = append(out, it)
	}
	return out, rows.Err()
}

// encodeCursor and decodeCursor carry a historyKey as an opaque URL-safe string.
func encodeCursor(k historyKey) string {
	raw := k.At.UTC().Format(time.RFC3339Nano) + "|" + strconv.Itoa(k.FightIndex) + "|" + k.ReportID
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

var errBadCursor = errors.New("coach: bad history cursor")

func decodeCursor(s string) (*historyKey, error) {
	if s == "" {
		return nil, nil
	}
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, errBadCursor
	}
	parts := strings.SplitN(string(raw), "|", 3)
	if len(parts) != 3 {
		return nil, errBadCursor
	}
	at, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return nil, errBadCursor
	}
	fight, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errBadCursor
	}
	return &historyKey{At: at, FightIndex: fight, ReportID: parts[2]}, nil
}

// history is GET /v1/characters/{region}/{ruleset}/{name}/coach/history?cursor=: the owner
// only (404 otherwise, the body of an unknown character), log_coach only (403 no_plan).
func (s *Service) history(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	key, ok := s.ownedCharacter(w, r)
	if !ok || !s.hasCoach(w, r) {
		return
	}
	after, err := decodeCursor(r.URL.Query().Get("cursor"))
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, codeInvalid, "that cursor is not one this route gave", nil)
		return
	}
	items, err := s.storedReviews(r.Context(), []string{key}, after, historyPageSize+1)
	if err != nil {
		s.linesFail(w, r, "history", err)
		return
	}
	page := historyPage{Reviews: items}
	if len(items) > historyPageSize {
		page.Reviews = items[:historyPageSize]
		page.NextCursor = encodeCursor(page.Reviews[historyPageSize-1].cursor)
	}
	httpx.WriteOK(w, r, http.StatusOK, page)
}

// ownedCharacter is the path's character key when the viewer owns it; else 404.
func (s *Service) ownedCharacter(w http.ResponseWriter, r *http.Request) (string, bool) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	owns := false
	key := ""
	if character.ValidRegion(region) && character.ValidRuleset(ruleset) {
		key = character.Key(region, ruleset, r.PathValue("name"))
		var err error
		if owns, err = s.Owners.OwnsCharacter(r.Context(), auth.ActorFrom(r.Context()).UserID, key); err != nil {
			s.linesFail(w, r, "history owner", err)
			return "", false
		}
	}
	if !owns {
		httpx.WriteError(w, r, http.StatusNotFound, codeNotFound, "no such character", nil)
	}
	return key, owns
}

// hasCoach answers 403 no_plan unless the viewer holds log_coach.
func (s *Service) hasCoach(w http.ResponseWriter, r *http.Request) bool {
	can, err := s.canCoach(r.Context(), auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.linesFail(w, r, "entitlement", err)
		return false
	}
	if !can {
		httpx.WriteError(w, r, http.StatusForbidden, codeNoPlan, "the coach is a Premium feature", nil)
	}
	return can
}

func (s *Service) canCoach(ctx context.Context, userID int64) (bool, error) {
	can, _, err := s.Entitlements.Can(ctx, userID, entitlements.FeatureLogCoach)
	if err != nil {
		return false, fmt.Errorf("coach: entitlement: %w", err)
	}
	return can, nil
}
