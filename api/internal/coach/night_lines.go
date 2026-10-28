package coach

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

const (
	// maxNightLines is how many characters' lines NightLinesFor answers.
	maxNightLines = 3
	// warmBudget bounds the background build of packs a reader stopped waiting for.
	warmBudget = 2 * time.Minute
	// coachTab is the report page's Coach tab.
	coachTab = "coach"
)

// ratedFightsSQL is a character's rated boss fights in one report (kills and wipes), in fight
// order: the night the night line reads. An insufficient card is not a rated fight.
const ratedFightsSQL = `
select distinct fight_index from rating_scores
where report_id = $1 and player_key = $2 and not coalesce(insufficient, false)
order by fight_index`

// ownedInReportSQL is the viewer's characters with a rated fight in the report.
const ownedInReportSQL = `
select distinct c.key, c.name from characters c
join rating_scores rs on rs.player_key = c.key
where rs.report_id = $1 and c.user_id = $2
order by c.name, c.key`

// holdersInReportSQL is the accounts owning a character with a rated fight in the report.
const holdersInReportSQL = `
select distinct c.user_id from characters c
join rating_scores rs on rs.player_key = c.key
where rs.report_id = $1 and c.user_id is not null
order by c.user_id`

// CoachLine is one character's night line for Discord (Task 32): Pending when its packs
// were not ready inside the budget (they warm in the background).
type CoachLine struct {
	Character, PlayerKey string
	Line                 NightLine
	Pending              bool
	URL                  string
}

// ReportURL is the report page with the Coach tab open on fight.
func (s *Service) ReportURL(reportID string, fight int) string {
	q := url.Values{"tab": {coachTab}}
	if fight > 0 {
		q.Set("fight", strconv.Itoa(fight))
	}
	return s.PublicBaseURL + "/reports/" + url.PathEscape(reportID) + "?" + q.Encode()
}

// ratedFights is playerKey's rated fight indices in reportID.
func (s *Service) ratedFights(ctx context.Context, reportID, playerKey string) ([]int, error) {
	rows, err := s.Store.DB.Query(ctx, ratedFightsSQL, reportID, playerKey)
	if err != nil {
		return nil, fmt.Errorf("coach: rated fights %s: %w", reportID, err)
	}
	defer rows.Close()
	var out []int
	for rows.Next() {
		var i int
		if err := rows.Scan(&i); err != nil {
			return nil, fmt.Errorf("coach: rated fights %s: %w", reportID, err)
		}
		out = append(out, i)
	}
	return out, rows.Err()
}

// packsWithin is playerKey's packs for fights of rep. A budget of zero waits for every
// pack; past a positive budget it answers pending and the packs keep building in the
// background, detached from ctx, so the next read finds them cached. A fight the player
// is not in is left out; any other failed pack is logged and left out.
func (s *Service) packsWithin(ctx context.Context, rep reports.Report, fights []int, playerKey string, budget time.Duration) ([]NightFight, bool) {
	if budget <= 0 {
		return s.buildNight(ctx, rep, fights, playerKey), false
	}
	warm, cancel := context.WithTimeout(context.WithoutCancel(ctx), warmBudget)
	done := make(chan []NightFight, 1)
	go func() {
		defer cancel()
		done <- s.buildNight(warm, rep, fights, playerKey)
	}()
	timer := time.NewTimer(budget)
	defer timer.Stop()
	select {
	case out := <-done:
		return out, false
	case <-timer.C:
		return nil, true
	case <-ctx.Done():
		return nil, true
	}
}

// buildNight loads every fight's pack in parallel, in fight order.
func (s *Service) buildNight(ctx context.Context, rep reports.Report, fights []int, playerKey string) []NightFight {
	got := make([]*NightFight, len(fights))
	var wg sync.WaitGroup
	for i, index := range fights {
		wg.Go(func() {
			p, err := s.PackFor(ctx, rep, index, playerKey)
			if err != nil {
				s.logPackFailure(rep.ID, index, err)
				return
			}
			got[i] = &NightFight{FightIndex: index, Pack: p}
		})
	}
	wg.Wait()
	out := make([]NightFight, 0, len(fights))
	for _, f := range got {
		if f != nil {
			out = append(out, *f)
		}
	}
	return out
}

func (s *Service) logPackFailure(reportID string, fight int, err error) {
	if errors.Is(err, ErrNoPlayer) || errors.Is(err, ErrNoFight) {
		return
	}
	s.logger().Warn("coach: night pack", slog.String("report", reportID), slog.Int("fight", fight),
		slog.String("error", err.Error()))
}

// nightOf is playerKey's night line in rep, and whether its packs were still building.
func (s *Service) nightOf(ctx context.Context, rep reports.Report, playerKey string, budget time.Duration) (NightLine, bool, bool, error) {
	fights, err := s.ratedFights(ctx, rep.ID, playerKey)
	if err != nil {
		return NightLine{}, false, false, err
	}
	night, pending := s.packsWithin(ctx, rep, fights, playerKey, budget)
	if pending {
		return NightLine{}, false, true, nil
	}
	line, ok := Night(night)
	return line, ok, false, nil
}

// NightLinesFor serves Discord (Task 32): one line per character the viewer owns in the
// report, three at most, and none for a viewer without log_coach. Packs missing past budget
// make that line Pending and warm in the background.
func (s *Service) NightLinesFor(ctx context.Context, userID int64, reportID string, budget time.Duration) ([]CoachLine, error) {
	can, _, err := s.Entitlements.Can(ctx, userID, entitlements.FeatureLogCoach)
	if err != nil || !can {
		return nil, err
	}
	rep, err := s.Reports.Get(ctx, reportID)
	if err != nil {
		return nil, fmt.Errorf("coach: night lines %s: %w", reportID, err)
	}
	owned, err := s.ownedInReport(ctx, userID, reportID)
	if err != nil {
		return nil, err
	}
	out := make([]CoachLine, 0, len(owned))
	for _, c := range owned {
		line, ok, pending, err := s.nightOf(ctx, rep, c.PlayerKey, budget)
		if err != nil {
			return nil, err
		}
		if ok || pending {
			c.Line, c.Pending, c.URL = line, pending, s.ReportURL(reportID, line.FightIndex)
			out = append(out, c)
		}
	}
	return out, nil
}

// ownedInReport is the viewer's characters in the report, maxNightLines at most.
func (s *Service) ownedInReport(ctx context.Context, userID int64, reportID string) ([]CoachLine, error) {
	rows, err := s.Store.DB.Query(ctx, ownedInReportSQL, reportID, userID)
	if err != nil {
		return nil, fmt.Errorf("coach: owned characters in %s: %w", reportID, err)
	}
	defer rows.Close()
	var out []CoachLine
	for rows.Next() && len(out) < maxNightLines {
		var c CoachLine
		if err := rows.Scan(&c.PlayerKey, &c.Character); err != nil {
			return nil, fmt.Errorf("coach: owned characters in %s: %w", reportID, err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// AnyCoachHolder answers notify's raid-night sentence (Task 32): whether any account owning
// a character with a rated fight in the report holds log_coach.
func (s *Service) AnyCoachHolder(ctx context.Context, reportID string) (bool, error) {
	rows, err := s.Store.DB.Query(ctx, holdersInReportSQL, reportID)
	if err != nil {
		return false, fmt.Errorf("coach: holders in %s: %w", reportID, err)
	}
	var users []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return false, fmt.Errorf("coach: holders in %s: %w", reportID, err)
		}
		users = append(users, id)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("coach: holders in %s: %w", reportID, err)
	}
	for _, id := range users {
		if can, _, err := s.Entitlements.Can(ctx, id, entitlements.FeatureLogCoach); err != nil || can {
			return can, err
		}
	}
	return false, nil
}
