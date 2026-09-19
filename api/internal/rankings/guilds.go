package rankings

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// sortBest orders a character's bests by encounter then difficulty, so
// the page is stable between loads.
func sortBest(rows []CharacterBest) {
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].EncounterID != rows[j].EncounterID {
			return rows[i].EncounterID < rows[j].EncounterID
		}
		return rows[i].Difficulty < rows[j].Difficulty
	})
}

// sortBuilds orders the builds a character has been seen in, newest
// first seen last.
func sortBuilds(rows []BuildSeen) {
	sort.Slice(rows, func(i, j int) bool { return rows[i].FirstSeen.Before(rows[j].FirstSeen) })
}

// The guild leaderboard kinds.
const (
	KindSpeed     = "speed"
	KindExecution = "execution"
	KindProgress  = "progress"
)

// ValidKind reports whether k is a guild leaderboard kind.
func ValidKind(k string) bool {
	return k == KindSpeed || k == KindExecution || k == KindProgress
}

// GuildRow is one line of a guild leaderboard.
type GuildRow struct {
	Rank        int       `json:"rank"`
	Guild       GuildRef  `json:"guild"`
	Value       float64   `json:"value"`
	EncounterID int64     `json:"encounter_id,omitempty"`
	Kills       int       `json:"kills,omitempty"`
	FoughtAt    time.Time `json:"fought_at"`
	ReportID    string    `json:"report_id,omitempty"`
	FightIndex  int       `json:"fight_index,omitempty"`
}

// GuildRankings answers one guild leaderboard. Speed ranks the fastest
// kill, execution the kill with the fewest deaths, and progress the
// guilds that have killed the most encounters, earliest first.
func (s *Store) GuildRankings(ctx context.Context, encounterID int64, kind, at string) ([]GuildRow, error) {
	var (
		query string
		args  []any
	)
	phaseClause := ""
	if at != "" {
		phaseClause = " and m.phase = $3"
	}
	switch kind {
	case KindSpeed:
		query = `select g.name, g.region, g.ruleset, min(m.duration_ms)::float8, min(m.fought_at)
			 from fight_metrics m
			 join reports r on r.id = m.report_id
			 join guilds g on g.id = r.guild_id
			 where m.encounter_id = $1 and m.kill and m.state <> 'removed'` + phaseClause + `
			 group by g.id, g.name, g.region, g.ruleset
			 order by 4 asc limit $2`
	case KindExecution:
		query = `select g.name, g.region, g.ruleset, min(deaths)::float8, min(fought_at) from (
			   select r.guild_id as guild_id, sum(m.deaths) as deaths, min(m.fought_at) as fought_at
			   from fight_metrics m
			   join reports r on r.id = m.report_id
			   where m.encounter_id = $1 and m.kill and m.state <> 'removed'` + phaseClause + `
			     and r.guild_id is not null
			   group by r.guild_id, m.report_id, m.fight_index
			 ) k join guilds g on g.id = k.guild_id
			 group by g.id, g.name, g.region, g.ruleset
			 order by 4 asc limit $2`
	case KindProgress:
		query = `select g.name, g.region, g.ruleset, count(distinct m.encounter_id)::float8, min(m.fought_at)
			 from fight_metrics m
			 join reports r on r.id = m.report_id
			 join guilds g on g.id = r.guild_id
			 where m.kill and m.state <> 'removed' and ($1 = 0 or m.encounter_id = $1)` + phaseClause + `
			 group by g.id, g.name, g.region, g.ruleset
			 order by 4 desc, 5 asc limit $2`
	default:
		return nil, fmt.Errorf("rankings: %q is not a guild leaderboard kind", kind)
	}
	args = []any{encounterID, PerPage}
	if at != "" {
		args = append(args, at)
	}
	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("rankings: guild leaderboard: %w", err)
	}
	defer rows.Close()
	out := []GuildRow{}
	rank := 1
	for rows.Next() {
		var r GuildRow
		if err := rows.Scan(&r.Guild.Name, &r.Guild.Region, &r.Guild.Ruleset, &r.Value, &r.FoughtAt); err != nil {
			return nil, fmt.Errorf("rankings: scan guild row: %w", err)
		}
		r.Rank, r.EncounterID = rank, encounterID
		if kind == KindProgress {
			r.EncounterID, r.Kills = 0, int(r.Value)
		}
		rank++
		out = append(out, r)
	}
	return out, rows.Err()
}

// Progression is one encounter's standing for a guild.
type Progression struct {
	Encounter   string     `json:"encounter"`
	EncounterID int64      `json:"encounter_id"`
	Difficulty  int64      `json:"difficulty"`
	Kills       int        `json:"kills"`
	PullCount   int        `json:"pull_count"`
	FirstKillAt *time.Time `json:"first_kill_at,omitempty"`
}

// RosterBest is one member's best parse for the guild.
type RosterBest struct {
	Player      Player    `json:"player"`
	Encounter   string    `json:"encounter"`
	EncounterID int64     `json:"encounter_id"`
	Metric      string    `json:"metric"`
	Value       float64   `json:"value"`
	FoughtAt    time.Time `json:"fought_at"`
	// ExecutionScore is the fraction of what this raider's gear can
	// do that they actually did on that parse, or null when it has
	// none.
	ExecutionScore *float64 `json:"execution_score"`
}

// GuildReport is one of a guild's reports, newest first.
type GuildReport struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Zone      string    `json:"zone"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Guild is the body of the guild page.
type Guild struct {
	Guild       GuildIdentity `json:"guild"`
	Progression []Progression `json:"progression"`
	RosterBest  []RosterBest  `json:"roster_best"`
	Reports     []GuildReport `json:"reports"`
}

// GuildIdentity names a guild.
type GuildIdentity struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
}

// GuildReportLimit is how many recent reports the guild page lists.
const GuildReportLimit = 25

// Guild assembles a guild page. It reports false when no such guild is
// registered.
func (s *Store) Guild(ctx context.Context, region, ruleset, name string) (Guild, bool, error) {
	out := Guild{Progression: []Progression{}, RosterBest: []RosterBest{}, Reports: []GuildReport{}}
	err := s.Pool.QueryRow(ctx,
		`select id, name, region, ruleset from guilds
		 where region = $1 and ruleset = $2 and lower(name) = lower($3)`,
		strings.ToLower(region), strings.ToLower(ruleset), name).
		Scan(&out.Guild.ID, &out.Guild.Name, &out.Guild.Region, &out.Guild.Ruleset)
	if err == pgx.ErrNoRows {
		return Guild{}, false, nil
	}
	if err != nil {
		return Guild{}, false, fmt.Errorf("rankings: read guild: %w", err)
	}

	prog, err := s.Pool.Query(ctx,
		`select f.encounter_id, coalesce(f.difficulty, 0),
		        count(*) filter (where f.kill), count(*), min(f.start_ms) filter (where f.kill)
		 from fights f join reports r on r.id = f.report_id
		 where r.guild_id = $1 and f.encounter_id is not null and r.visibility <> 'private'
		 group by f.encounter_id, f.difficulty order by f.encounter_id`, out.Guild.ID)
	if err != nil {
		return Guild{}, false, fmt.Errorf("rankings: read progression: %w", err)
	}
	defer prog.Close()
	for prog.Next() {
		var (
			p           Progression
			firstKillMS *int64
		)
		if err := prog.Scan(&p.EncounterID, &p.Difficulty, &p.Kills, &p.PullCount, &firstKillMS); err != nil {
			return Guild{}, false, fmt.Errorf("rankings: scan progression: %w", err)
		}
		if firstKillMS != nil {
			at := time.UnixMilli(*firstKillMS).UTC()
			p.FirstKillAt = &at
		}
		out.Progression = append(out.Progression, p)
	}
	if err := prog.Err(); err != nil {
		return Guild{}, false, err
	}

	best, err := s.Pool.Query(ctx,
		`select distinct on (m.player_key, m.encounter_id)
		        m.player_key, m.player_name, coalesce(m.class, ''), coalesce(m.spec, ''),
		        m.encounter_id, m.role, m.metric_dps, m.metric_hps, m.damage_taken, m.fought_at,
		        m.execution_score
		 from fight_metrics m join reports r on r.id = m.report_id
		 where r.guild_id = $1 and m.kill and m.state <> 'removed'
		 order by m.player_key, m.encounter_id, greatest(coalesce(m.metric_dps, 0),
		          coalesce(m.metric_hps, 0), coalesce(m.damage_taken, 0)) desc`, out.Guild.ID)
	if err != nil {
		return Guild{}, false, fmt.Errorf("rankings: read roster bests: %w", err)
	}
	defer best.Close()
	for best.Next() {
		var (
			r           RosterBest
			role        string
			dps, hps    *float64
			taken       *int64
			encounterID *int64
		)
		if err := best.Scan(&r.Player.Key, &r.Player.Name, &r.Player.Class, &r.Player.Spec,
			&encounterID, &role, &dps, &hps, &taken, &r.FoughtAt, &r.ExecutionScore); err != nil {
			return Guild{}, false, fmt.Errorf("rankings: scan roster best: %w", err)
		}
		if encounterID != nil {
			r.EncounterID = *encounterID
		}
		r.Metric = metricOf(role)
		r.Value = pickValue(r.Metric, dps, hps, taken)
		out.RosterBest = append(out.RosterBest, r)
	}
	if err := best.Err(); err != nil {
		return Guild{}, false, err
	}

	// One query names every encounter both lists mention.
	ids := make([]int64, 0, len(out.Progression)+len(out.RosterBest))
	for _, p := range out.Progression {
		ids = append(ids, p.EncounterID)
	}
	for _, r := range out.RosterBest {
		ids = append(ids, r.EncounterID)
	}
	names, err := s.EncounterNames(ctx, ids)
	if err != nil {
		return Guild{}, false, err
	}
	for i := range out.Progression {
		out.Progression[i].Encounter = names[out.Progression[i].EncounterID]
	}
	for i := range out.RosterBest {
		out.RosterBest[i].Encounter = names[out.RosterBest[i].EncounterID]
	}

	// This list is report identities - id, title, zone - served to
	// anyone who loads the guild page, so it takes the stricter rule
	// than the derived aggregates above: "guild" means the guild may
	// read it, and reports.mayView refuses the body to everyone else.
	// Listing it here would tell the world the report exists and what
	// it is called. The progression counts stay on reports.Ranked,
	// which is the right rule for an aggregate.
	reps, err := s.Pool.Query(ctx,
		`select id, title, zone, status, created_at from reports
		 where guild_id = $1 and visibility in ('public', 'unlisted')
		 order by created_at desc limit $2`, out.Guild.ID, GuildReportLimit)
	if err != nil {
		return Guild{}, false, fmt.Errorf("rankings: read guild reports: %w", err)
	}
	defer reps.Close()
	for reps.Next() {
		var r GuildReport
		if err := reps.Scan(&r.ID, &r.Title, &r.Zone, &r.Status, &r.CreatedAt); err != nil {
			return Guild{}, false, fmt.Errorf("rankings: scan guild report: %w", err)
		}
		out.Reports = append(out.Reports, r)
	}
	return out, true, reps.Err()
}
