package activity

// The weekly tier move (spec 2026-11-08-app-flows F4): after the hourly tier snapshot, a damage
// spec whose grade this week differs from last week's tells the accounts whose main is that spec
// at that band and faction (`Fury Warrior moved to S`). Each grade is told once
// (tier_grade_weeks.notified_grade, migration 0086); a second move inside the week is told
// again. The main's spec, level and faction are read as Home reads a character.

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// tierMoveCopy is the words of the tier move.
var tierMoveCopy = struct{ Title, Up, Down, TiersPath string }{
	Title:     "%s moved to %s",
	Up:        "Up from %s last week, level %s, %s.",
	Down:      "Down from %s last week, level %s, %s.",
	TiersPath: "/tiers",
}

// tierMove is one spec row whose grade changed since last week.
type tierMove struct {
	spec, faction, grade, last string
	band                       int
}

const tierMovesSQL = `select t.spec, t.band, t.faction, t.grade, p.grade
	from tier_grade_weeks t
	join tier_grade_weeks p on p.week = $2 and p.spec = t.spec and p.band = t.band and p.faction = t.faction
	where t.week = $1 and t.grade is not null and p.grade is not null and t.grade <> p.grade
	  and t.notified_grade is distinct from t.grade`

// mainsSQL is every account's main with what resolves its spec, level and faction.
const mainsSQL = `select u.id, c.level, coalesce(c.faction, 'alliance'), c.class, x.export,
	(select cs.spec from character_sims cs where cs.character_key = c.key order by cs.created_at desc limit 1)
	from users u join characters c on c.key = u.main_character_key
	left join addon_exports x on x.character_key = c.key
	where not coalesce(u.anonymize, false)`

// mainGroup is the spec, band and faction a main is told about.
type mainGroup struct {
	spec, faction string
	band          int
}

// PublishTierMoves publishes this week's untold tier moves at now to the specs' mains and marks
// them told; it answers how many moves it published.
func (s *Store) PublishTierMoves(ctx context.Context, pub events.Publisher, publicBaseURL string, now time.Time) (int, error) {
	if pub == nil {
		return 0, nil
	}
	week := WeekOf(now)
	moves, err := s.tierMoves(ctx, week, PreviousWeek(now))
	if err != nil || len(moves) == 0 {
		return 0, err
	}
	mains, err := s.mainsByGroup(ctx)
	if err != nil {
		return 0, err
	}
	for _, m := range moves {
		if users := mains[mainGroup{m.spec, m.faction, m.band}]; len(users) > 0 {
			if err := pub.Publish(ctx, tierMoveEvent(m, week, users, publicBaseURL)); err != nil {
				return 0, fmt.Errorf("activity: publish tier move %s: %w", m.spec, err)
			}
		}
		if _, err := s.Pool.Exec(ctx, `update tier_grade_weeks set notified_grade = grade
			where week = $1 and spec = $2 and band = $3 and faction = $4`, week, m.spec, m.band, m.faction); err != nil {
			return 0, fmt.Errorf("activity: mark tier move %s: %w", m.spec, err)
		}
	}
	return len(moves), nil
}

func (s *Store) tierMoves(ctx context.Context, week, last time.Time) ([]tierMove, error) {
	rows, err := s.Pool.Query(ctx, tierMovesSQL, week, last)
	if err != nil {
		return nil, fmt.Errorf("activity: tier moves: %w", err)
	}
	defer rows.Close()
	var out []tierMove
	for rows.Next() {
		var m tierMove
		if err := rows.Scan(&m.spec, &m.band, &m.faction, &m.grade, &m.last); err != nil {
			return nil, fmt.Errorf("activity: scan tier move: %w", err)
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// mainsByGroup is every account's main grouped by its spec, band and faction.
func (s *Store) mainsByGroup(ctx context.Context) (map[mainGroup][]int64, error) {
	rows, err := s.Pool.Query(ctx, mainsSQL)
	if err != nil {
		return nil, fmt.Errorf("activity: mains: %w", err)
	}
	defer rows.Close()
	out := map[mainGroup][]int64{}
	for rows.Next() {
		var uid int64
		var f characterFacts
		if err := rows.Scan(&uid, &f.level, &f.faction, &f.class, &f.export, &f.simSpec); err != nil {
			return nil, fmt.Errorf("activity: scan main: %w", err)
		}
		spec, level, _ := f.resolve()
		if spec != "" {
			g := mainGroup{spec, f.faction, BandOf(level)}
			out[g] = append(out[g], uid)
		}
	}
	return out, rows.Err()
}

// tierMoveEvent is m told to users.
func tierMoveEvent(m tierMove, week time.Time, users []int64, publicBaseURL string) events.Event {
	body := tierMoveCopy.Down
	if gradeIndex(m.grade) < gradeIndex(m.last) {
		body = tierMoveCopy.Up
	}
	return events.Event{
		Kind: events.TierMoved, Audience: events.Users(users...),
		DedupeKey: fmt.Sprintf("tier.moved:%s:%s:%d:%s:%s", week.Format(time.DateOnly), m.spec, m.band, m.faction, m.grade),
		Title:     fmt.Sprintf(tierMoveCopy.Title, SpecName(m.spec), m.grade),
		Body:      fmt.Sprintf(body, m.last, BandWords(m.band), titleWord(m.faction)),
		Link:      publicBaseURL + tierMoveCopy.TiersPath,
		Data: map[string]string{"spec": m.spec, "band": fmt.Sprint(m.band), "faction": m.faction,
			"grade": m.grade, "last_grade": m.last},
	}
}

// gradeIndex is a letter's place in tierBands, best first; an unknown letter sorts last.
func gradeIndex(letter string) int {
	i := slices.IndexFunc(tierBands, func(b tierBand) bool { return strings.EqualFold(b.letter, letter) })
	if i < 0 {
		return len(tierBands)
	}
	return i
}
