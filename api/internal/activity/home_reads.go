package activity

import (
	"context"
	"fmt"
	"sort"
	"time"
)

// movers shown in the block.
const moversShown = 3

// guildSimsShown is the ladder's top rows; the viewer's own row is pinned below them.
const guildSimsShown = 5

func (s *Store) latestImportSim(ctx context.Context, key string) (*ImportSim, bool, error) {
	var state, id string
	var dps float64
	var at *time.Time
	err := s.Pool.QueryRow(ctx,
		`select sm.id, sm.state, sm.dps_mean, sm.finished_at from character_sims cs join sims sm on sm.id = cs.sim_id
		 where cs.character_key = $1 and cs.kind = 'import' order by cs.created_at desc limit 1`, key).Scan(&id, &state, &dps, &at)
	if isNoRows(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("activity: import sim %s: %w", key, err)
	}
	sim := &ImportSim{State: SimPending, SimID: &id}
	switch state {
	case "done":
		sim.State, sim.DPS, sim.SimmedAt = SimDone, &dps, at
	case "error":
		sim.State = SimFailed
	}
	return sim, true, nil
}

func (s *Store) specWeek(ctx context.Context, spec string, band int, faction string, now time.Time) (*SpecWeek, error) {
	w := SpecWeek{Spec: spec, Band: band, Faction: faction}
	err := s.Pool.QueryRow(ctx,
		`select t.role, t.grade, t.place, t.of_n,
		        (select p.grade from tier_grade_weeks p where p.week = $5 and p.spec = t.spec and p.band = t.band and p.faction = t.faction)
		 from tier_grade_weeks t where t.week = $1 and t.spec = $2 and t.band = $3 and t.faction = $4`,
		WeekOf(now), spec, band, faction, PreviousWeek(now)).Scan(&w.Role, &w.Grade, &w.Place, &w.Of, &w.LastWeekGrade)
	if isNoRows(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("activity: spec week %s: %w", spec, err)
	}
	return &w, nil
}

func (s *Store) boardWeek(ctx context.Context, spec string, band int, week time.Time) ([]MoverRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select rank, name, class, character_key from spec_board_weeks where week = $1 and spec = $2 and band = $3 order by rank`,
		week, spec, band)
	if err != nil {
		return nil, fmt.Errorf("activity: board %s: %w", spec, err)
	}
	defer rows.Close()
	out := []MoverRow{}
	for rows.Next() {
		var r MoverRow
		if err := rows.Scan(&r.Rank, &r.Name, &r.Class, &r.Key); err != nil {
			return nil, fmt.Errorf("activity: scan board: %w", err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// movers is the board's biggest climbs since last week, or, before there is a last week,
// its top 3 with no movement.
func (s *Store) movers(ctx context.Context, spec string, band int, now time.Time) (*Movers, error) {
	current, err := s.boardWeek(ctx, spec, band, WeekOf(now))
	if err != nil || len(current) == 0 {
		return nil, err
	}
	prev, err := s.boardWeek(ctx, spec, band, PreviousWeek(now))
	if err != nil {
		return nil, err
	}
	m := &Movers{Spec: spec, Band: band, Baseline: len(prev) > 0, Rows: []MoverRow{}}
	if !m.Baseline {
		m.Rows = current[:min(moversShown, len(current))]
		return m, nil
	}
	was := map[string]int{}
	for _, r := range prev {
		was[r.Key] = r.Rank
	}
	for _, r := range current {
		if before, ok := was[r.Key]; ok && before > r.Rank {
			d := before - r.Rank
			r.Delta = &d
			m.Rows = append(m.Rows, r)
		}
	}
	sort.SliceStable(m.Rows, func(i, j int) bool { return *m.Rows[i].Delta > *m.Rows[j].Delta })
	m.Rows = m.Rows[:min(moversShown, len(m.Rows))]
	return m, nil
}

// guildSims is the ladder of the character's guild for its role and band, under the preset
// of the character's own latest sim there (bare when it has none).
func (s *Store) guildSims(ctx context.Context, c homeCharacter, band int, now time.Time) (*GuildSims, error) {
	g := GuildSims{Role: SpecRole(c.spec), Band: band, Preset: barePreset}
	err := s.Pool.QueryRow(ctx,
		`select gc.guild_id, g.name, coalesce((select cs.preset from character_sims cs where cs.character_key = $1
		   and cs.role = $2 and cs.band = $3 order by cs.created_at desc limit 1), $4)
		 from guild_characters gc join guilds g on g.id = gc.guild_id
		 where gc.character_key = $1 and gc.verified_at is not null`, c.key, g.Role, band, barePreset).
		Scan(&g.GuildID, &g.GuildName, &g.Preset)
	if isNoRows(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("activity: guild of %s: %w", c.key, err)
	}
	ladder, err := s.GuildLadder(ctx, LadderGroup{g.GuildID, g.Role, band, g.Preset}, now)
	if err != nil || len(ladder) < ladderMin {
		return nil, err
	}
	_, g.Unit = RoleWords(g.Role)
	g.Of = len(ladder)
	for i := range ladder {
		ladder[i].You = ladder[i].Key == c.key
		if ladder[i].You {
			you := ladder[i]
			g.You = &you
		}
	}
	g.Rows = ladder[:min(guildSimsShown, len(ladder))]
	return &g, nil
}

// guildDay is the character's guild membership and whether its guild raids today: a raid (a
// raid zone, planned or locked, not yet over) starting on today's date in the guild's raid
// time zone (UTC when it has none or names an unknown zone).
func (s *Store) guildDay(ctx context.Context, key string, now time.Time) (inGuild, raidDay bool, err error) {
	var tz string
	var next *time.Time
	err = s.Pool.QueryRow(ctx,
		`select coalesce(g.raid_timezone, ''),
		        (select r.starts_at from guild_raids r where r.guild_id = gc.guild_id and r.status in ('planned', 'locked')
		           and starts_with(r.zone, 'raid:') and r.starts_at + make_interval(mins => r.duration_minutes) > $2
		         order by r.starts_at limit 1)
		 from guild_characters gc join guilds g on g.id = gc.guild_id
		 where gc.character_key = $1 and gc.verified_at is not null
		 order by gc.guild_id limit 1`, key, now).Scan(&tz, &next)
	if isNoRows(err) {
		return false, false, nil
	}
	if err != nil {
		return false, false, fmt.Errorf("activity: guild day of %s: %w", key, err)
	}
	return true, next != nil && sameDay(*next, now, tz), nil
}

// sameDay reports whether a and b fall on one calendar date in the named zone.
func sameDay(a, b time.Time, tz string) bool {
	loc, err := time.LoadLocation(tz)
	if err != nil || tz == "" {
		loc = time.UTC
	}
	return a.In(loc).Format(time.DateOnly) == b.In(loc).Format(time.DateOnly)
}
