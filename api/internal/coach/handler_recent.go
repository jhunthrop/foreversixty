package coach

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// recentBudget bounds the pack builds one recent-line answer waits for.
const recentBudget = 3 * time.Second

const accountCharactersSQL = `select key from characters where user_id = $1 order by key`

// firstDTO is the last raid's night line.
type firstDTO struct {
	Stem string `json:"stem"`
	On   int    `json:"on"`
	Of   int    `json:"of"`
}

// lastRaidDTO is the account's last raid (§10, the account line).
type lastRaidDTO struct {
	Zone         string    `json:"zone"`
	FoughtAt     time.Time `json:"fought_at"`
	ReportID     string    `json:"report_id"`
	FindingCount int       `json:"finding_count"`
	First        *firstDTO `json:"first"`
	URL          string    `json:"url"`
}

// recentDTO is GET /v1/me/coach/recent's body.
type recentDTO struct {
	LastRaid *lastRaidDTO  `json:"last_raid"`
	Pending  bool          `json:"pending"`
	Reviews  []historyItem `json:"reviews"`
}

// lastNight is one character's last raid: its report and rated fights.
type lastNight struct {
	key, reportID, zone string
	foughtAt, latest    time.Time
	fights              []int
}

// recentLine is GET /v1/me/coach/recent: the account's last raid line for everyone signed
// in, and the newest 20 stored reviews for log_coach (a free account gets []).
func (s *Service) recentLine(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	ctx := r.Context()
	userID := auth.ActorFrom(ctx).UserID
	keys, err := s.accountCharacters(ctx, userID)
	if err != nil {
		s.linesFail(w, r, "recent characters", err)
		return
	}
	body := recentDTO{Reviews: []historyItem{}}
	if body.LastRaid, body.Pending, err = s.lastRaid(ctx, keys); err != nil {
		s.linesFail(w, r, "recent last raid", err)
		return
	}
	if body.Reviews, err = s.recentReviews(ctx, userID, keys); err != nil {
		s.linesFail(w, r, "recent reviews", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, body)
}

func (s *Service) recentReviews(ctx context.Context, userID int64, keys []string) ([]historyItem, error) {
	can, err := s.canCoach(ctx, userID)
	if err != nil || !can || len(keys) == 0 {
		return []historyItem{}, err
	}
	return s.storedReviews(ctx, keys, nil, historyPageSize)
}

func (s *Service) accountCharacters(ctx context.Context, userID int64) ([]string, error) {
	rows, err := s.Store.DB.Query(ctx, accountCharactersSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("coach: account characters: %w", err)
	}
	defer rows.Close()
	var keys []string
	for rows.Next() {
		var k string
		if err := rows.Scan(&k); err != nil {
			return nil, fmt.Errorf("coach: account characters: %w", err)
		}
		keys = append(keys, k)
	}
	return keys, rows.Err()
}

// lastRaid is the newest last raid of any of keys within addon.LastRaidWindow, with its
// night line; nil when there is none. pending: the packs were still building at the budget.
func (s *Service) lastRaid(ctx context.Context, keys []string) (*lastRaidDTO, bool, error) {
	if len(keys) == 0 {
		return nil, false, nil
	}
	fights, err := (&addon.Store{Pool: s.Store.DB, Log: s.Log}).LastRaidFights(ctx, keys, s.linesNow().Add(-addon.LastRaidWindow))
	if err != nil {
		return nil, false, fmt.Errorf("coach: last raids: %w", err)
	}
	night, ok := newestNight(fights)
	if !ok {
		return nil, false, nil
	}
	rep, err := s.Reports.Get(ctx, night.reportID)
	if err != nil {
		return nil, false, fmt.Errorf("coach: last raid report %s: %w", night.reportID, err)
	}
	packs, pending := s.packsWithin(ctx, rep, night.fights, night.key, recentBudget)
	out := &lastRaidDTO{Zone: night.zone, FoughtAt: night.foughtAt, ReportID: night.reportID,
		URL: s.ReportURL(night.reportID, 0)}
	if line, ok := Night(packs); ok {
		out.FindingCount = line.FindingCount
		out.First = &firstDTO{Stem: line.Stem, On: line.On, Of: line.Of}
		out.URL = s.ReportURL(night.reportID, line.FightIndex)
	}
	return out, pending, nil
}

// newestNight groups the fights (oldest first per character) by character and keeps the
// character whose last fight is the newest.
func newestNight(fights []addon.LastRaidFight) (lastNight, bool) {
	var best *lastNight
	var cur *lastNight
	for _, f := range fights {
		if cur == nil || cur.key != f.CharacterKey {
			cur = &lastNight{key: f.CharacterKey, reportID: f.ReportID, zone: f.Zone, foughtAt: f.FoughtAt}
		}
		cur.fights = append(cur.fights, f.FightIndex)
		cur.latest = f.FoughtAt
		if best == nil || cur.latest.After(best.latest) {
			best = cur
		}
	}
	if best == nil {
		return lastNight{}, false
	}
	return *best, true
}

func (s *Service) linesNow() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}
