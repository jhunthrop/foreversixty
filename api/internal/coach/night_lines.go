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

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

const (
	// maxNightLines is how many characters' lines NightLinesFor answers.
	maxNightLines = 3
	// maxPackBuilds bounds how many of one night's packs build at once.
	maxPackBuilds = 4
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
	slots := make(chan struct{}, maxPackBuilds)
	var wg sync.WaitGroup
	for i, index := range fights {
		wg.Go(func() {
			slots <- struct{}{}
			defer func() { <-slots }()
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
// report, three at most, and none for a viewer without log_coach or for a report the night
// route would refuse (flagged, or one the viewer may not view). The characters share one
// deadline and are read in parallel; a character not finished by it is Pending (its packs
// warm in the background), so the lines that are ready always post.
func (s *Service) NightLinesFor(ctx context.Context, userID int64, reportID string, budget time.Duration) ([]CoachLine, error) {
	can, err := s.canCoach(ctx, userID)
	if err != nil || !can {
		return nil, err
	}
	rep, err := s.Reports.Get(ctx, reportID)
	if errors.Is(err, reports.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("coach: night lines %s: %w", reportID, err)
	}
	if rep.Flagged != nil || !reports.MayView(ctx, auth.Actor{UserID: userID, Role: "user"}, rep, s.Accounts) {
		return nil, nil
	}
	owned, err := s.ownedInReport(ctx, userID, reportID)
	if err != nil {
		return nil, err
	}
	return s.linesBy(ctx, rep, owned, time.Now().Add(budget)), nil
}

// linesBy reads every character's line in parallel against one deadline, in owned's order,
// keeping the characters with a line or still pending, maxNightLines at most.
func (s *Service) linesBy(ctx context.Context, rep reports.Report, owned []CoachLine, deadline time.Time) []CoachLine {
	got := make([]*CoachLine, len(owned))
	var wg sync.WaitGroup
	for i, c := range owned {
		wg.Go(func() { got[i] = s.lineBy(ctx, rep, c, deadline) })
	}
	wg.Wait()
	out := make([]CoachLine, 0, maxNightLines)
	for _, c := range got {
		if c != nil && len(out) < maxNightLines {
			out = append(out, *c)
		}
	}
	return out
}

// lineBy is one character's line by deadline: nil when it has none. A read that fails once
// the deadline or the request has passed is Pending; any other failure is logged and drops
// only that character.
func (s *Service) lineBy(ctx context.Context, rep reports.Report, c CoachLine, deadline time.Time) *CoachLine {
	line, ok, pending, err := s.nightOf(ctx, rep, c.PlayerKey, max(time.Until(deadline), time.Nanosecond))
	if err != nil {
		if ctx.Err() == nil && time.Now().Before(deadline) {
			s.logger().Warn("coach: night line", slog.String("report", rep.ID), slog.String("error", err.Error()))
			return nil
		}
		pending = true
	}
	if !ok && !pending {
		return nil
	}
	c.Line, c.Pending, c.URL = line, pending, s.ReportURL(rep.ID, line.FightIndex)
	return &c
}

// ownedInReport is the viewer's characters in the report, by name.
func (s *Service) ownedInReport(ctx context.Context, userID int64, reportID string) ([]CoachLine, error) {
	rows, err := s.Store.DB.Query(ctx, ownedInReportSQL, reportID, userID)
	if err != nil {
		return nil, fmt.Errorf("coach: owned characters in %s: %w", reportID, err)
	}
	defer rows.Close()
	var out []CoachLine
	for rows.Next() {
		var c CoachLine
		if err := rows.Scan(&c.PlayerKey, &c.Character); err != nil {
			return nil, fmt.Errorf("coach: owned characters in %s: %w", reportID, err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// AnyCoachHolder answers notify's raid-night sentence (Task 32): whether any account owning
// a character with a rated fight in the report holds log_coach. A flagged or private report
// has no coach sentence: its coach routes refuse it.
func (s *Service) AnyCoachHolder(ctx context.Context, reportID string) (bool, error) {
	rep, err := s.Reports.Get(ctx, reportID)
	if errors.Is(err, reports.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("coach: holders in %s: %w", reportID, err)
	}
	if rep.Flagged != nil || rep.Visibility == reports.Private {
		return false, nil
	}
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
