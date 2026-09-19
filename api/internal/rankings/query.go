package rankings

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// PerPage is the fixed page size the contract sets.
const PerPage = 100

// Player is who a ranking row belongs to.
type Player struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Class string `json:"class,omitempty"`
	Spec  string `json:"spec,omitempty"`
}

// GuildRef is the guild shown beside a name.
type GuildRef struct {
	Name    string `json:"name"`
	Ruleset string `json:"ruleset"`
	Region  string `json:"region"`
}

// Row is one line of a leaderboard.
type Row struct {
	Rank        int       `json:"rank"`
	Player      Player    `json:"player"`
	Guild       *GuildRef `json:"guild,omitempty"`
	Value       float64   `json:"value"`
	Ilvl        int       `json:"ilvl"`
	Size        int       `json:"size"`
	FoughtAt    time.Time `json:"fought_at"`
	DurationMS  int       `json:"duration_ms"`
	TalentSplit string    `json:"talent_split"`
	Trinkets    []int64   `json:"trinkets"`
	BuffCount   int       `json:"buff_count"`
	ReportID    string    `json:"report_id"`
	FightIndex  int       `json:"fight_index"`
	State       string    `json:"state"`
	// ExecutionScore is the fraction of what this player's gear can
	// do that they actually did. Null means the spec is not validated
	// or the fight predates scoring.
	ExecutionScore *float64 `json:"execution_score"`
}

// Page is a leaderboard page.
type Page struct {
	Rows      []Row     `json:"rows"`
	Total     int       `json:"total"`
	Page      int       `json:"page"`
	PerPage   int       `json:"per_page"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Query is a leaderboard request. Encounter is required; everything
// else narrows.
type Query struct {
	EncounterID int64
	Difficulty  *int64
	Metric      string
	Spec        string
	Class       string
	Phase       string
	Region      string
	Ruleset     string
	Faction     string
	Since       string
	Page        int
}

// metricColumn is the column a metric ranks on.
func metricColumn(metric string) string {
	switch metric {
	case MetricHPS:
		return "m.metric_hps"
	case MetricDamageTaken:
		return "m.damage_taken"
	case MetricExecution:
		return "m.execution_score"
	default:
		return "m.metric_dps"
	}
}

// where builds the shared filter for a leaderboard query.
func (q Query) where(now time.Time) (string, []any, error) {
	clauses := []string{"m.encounter_id = $1", "m.kill", "m.state <> 'removed'"}
	if q.Metric == MetricExecution {
		// A fight with no score is not last on the execution board,
		// it is simply not on it: the spec is not validated, or the
		// fight predates scoring.
		clauses = append(clauses, "m.execution_score is not null")
	}
	args := []any{q.EncounterID}
	add := func(sql string, value any) {
		args = append(args, value)
		clauses = append(clauses, fmt.Sprintf(sql, len(args)))
	}
	if q.Difficulty != nil {
		add("m.difficulty = $%d", *q.Difficulty)
	}
	if q.Spec != "" {
		add("m.spec = $%d", q.Spec)
	}
	if q.Class != "" {
		add("m.class = $%d", q.Class)
	}
	if q.Phase != "" {
		add("m.phase = $%d", q.Phase)
	}
	if q.Faction != "" {
		add("m.faction = $%d", q.Faction)
	}
	if q.Region != "" {
		add("split_part(m.player_key, '/', 1) = $%d", q.Region)
	}
	if q.Ruleset != "" {
		add("split_part(m.player_key, '/', 2) = $%d", q.Ruleset)
	}
	from, err := since(q.Since, now)
	if err != nil {
		return "", nil, err
	}
	if !from.IsZero() {
		add("m.fought_at >= $%d", from)
	}
	return strings.Join(clauses, " and "), args, nil
}

// Rankings answers one leaderboard page.
func (s *Store) Rankings(ctx context.Context, q Query, now time.Time) (Page, error) {
	where, args, err := q.where(now)
	if err != nil {
		return Page{}, err
	}
	if q.Page < 1 {
		q.Page = 1
	}
	page := Page{Rows: []Row{}, Page: q.Page, PerPage: PerPage, UpdatedAt: now.UTC()}
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from fight_metrics m where `+where, args...).Scan(&page.Total); err != nil {
		return Page{}, fmt.Errorf("rankings: count: %w", err)
	}
	column := metricColumn(q.Metric)
	rows, err := s.Pool.Query(ctx,
		`select m.player_key, m.player_name, coalesce(m.class, ''), coalesce(m.spec, ''),
		        `+column+`, coalesce(m.ilvl, 0), coalesce(m.size, 0), m.fought_at,
		        coalesce(m.duration_ms, 0), coalesce(m.talent_split, ''), m.trinkets, m.buff_count,
		        m.report_id, m.fight_index, m.state, m.execution_score, g.name, g.region, g.ruleset
		 from fight_metrics m
		 left join reports r on r.id = m.report_id
		 left join guilds g on g.id = r.guild_id
		 where `+where+`
		 order by `+column+` desc nulls last, m.fought_at desc, m.report_id, m.fight_index, m.player_key
		 limit `+fmt.Sprint(PerPage)+` offset `+fmt.Sprint((q.Page-1)*PerPage), args...)
	if err != nil {
		return Page{}, fmt.Errorf("rankings: query: %w", err)
	}
	defer rows.Close()
	rank := (q.Page-1)*PerPage + 1
	for rows.Next() {
		var (
			r                          Row
			value                      *float64
			guildName, region, ruleset *string
		)
		if err := rows.Scan(&r.Player.Key, &r.Player.Name, &r.Player.Class, &r.Player.Spec,
			&value, &r.Ilvl, &r.Size, &r.FoughtAt, &r.DurationMS, &r.TalentSplit, &r.Trinkets,
			&r.BuffCount, &r.ReportID, &r.FightIndex, &r.State, &r.ExecutionScore,
			&guildName, &region, &ruleset); err != nil {
			return Page{}, fmt.Errorf("rankings: scan: %w", err)
		}
		if value != nil {
			r.Value = *value
		}
		if guildName != nil && region != nil && ruleset != nil {
			r.Guild = &GuildRef{Name: *guildName, Region: *region, Ruleset: *ruleset}
		}
		if r.Trinkets == nil {
			r.Trinkets = []int64{}
		}
		r.Rank = rank
		rank++
		page.Rows = append(page.Rows, r)
	}
	return page, rows.Err()
}

// CharacterBest is a character's best parse on one encounter.
type CharacterBest struct {
	Encounter   string    `json:"encounter"`
	EncounterID int64     `json:"encounter_id"`
	Difficulty  int64     `json:"difficulty"`
	Metric      string    `json:"metric"`
	Value       float64   `json:"value"`
	Percentile  *float64  `json:"percentile,omitempty"`
	Spec        string    `json:"spec,omitempty"`
	FoughtAt    time.Time `json:"fought_at"`
	ReportID    string    `json:"report_id"`
	FightIndex  int       `json:"fight_index"`
}

// CharacterFight is one line of a character's history.
type CharacterFight struct {
	Encounter   string    `json:"encounter"`
	EncounterID int64     `json:"encounter_id"`
	Difficulty  int64     `json:"difficulty"`
	Kill        bool      `json:"kill"`
	Metric      string    `json:"metric"`
	Value       float64   `json:"value"`
	Percentile  *float64  `json:"percentile,omitempty"`
	Spec        string    `json:"spec,omitempty"`
	TalentSplit string    `json:"talent_split,omitempty"`
	DurationMS  int       `json:"duration_ms"`
	FoughtAt    time.Time `json:"fought_at"`
	ReportID    string    `json:"report_id"`
	FightIndex  int       `json:"fight_index"`
	// ExecutionScore is the fraction of what this player's gear can
	// do that they actually did, or null when the fight has no score.
	ExecutionScore *float64 `json:"execution_score"`
}

// BuildSeen is one talent split a character has been seen in.
//
// The contract's shape is { build_id, first_seen }. A planner build id
// needs the order the points were spent in, which no combat log
// records, so the split itself identifies the build until the planner
// can be handed one; the field is named accordingly rather than
// carrying an id that would always be empty.
type BuildSeen struct {
	TalentSplit string    `json:"talent_split"`
	Spec        string    `json:"spec,omitempty"`
	FirstSeen   time.Time `json:"first_seen"`
}

// Character is the body of the character page.
type Character struct {
	Character  CharacterRef     `json:"character"`
	Best       []CharacterBest  `json:"best"`
	History    []CharacterFight `json:"history"`
	BuildsSeen []BuildSeen      `json:"builds_seen"`
}

// CharacterRef names a character.
type CharacterRef struct {
	Key     string `json:"key"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Name    string `json:"name"`
	Class   string `json:"class,omitempty"`
}

// HistoryLimit is how many of a character's fights the page shows.
const HistoryLimit = 100

// Character assembles a character page. It reports false when the
// character has no ranked fights and no row of their own.
func (s *Store) Character(ctx context.Context, region, ruleset, name string) (Character, bool, error) {
	key := strings.ToLower(region) + "/" + strings.ToLower(ruleset) + "/" + strings.ToLower(name)
	out := Character{
		Character:  CharacterRef{Key: key, Region: region, Ruleset: ruleset, Name: name},
		Best:       []CharacterBest{},
		History:    []CharacterFight{},
		BuildsSeen: []BuildSeen{},
	}
	var class *string
	err := s.Pool.QueryRow(ctx,
		`select name, class from characters where key = $1`, key).Scan(&out.Character.Name, &class)
	if err == nil && class != nil {
		out.Character.Class = *class
	} else if err != nil && err != pgx.ErrNoRows {
		return Character{}, false, fmt.Errorf("rankings: read character: %w", err)
	}

	rows, err := s.Pool.Query(ctx,
		`select m.encounter_id, m.difficulty, m.kill, m.role, m.metric_dps, m.metric_hps,
		        m.damage_taken, coalesce(m.spec, ''), coalesce(m.talent_split, ''),
		        coalesce(m.duration_ms, 0), m.fought_at, m.report_id, m.fight_index,
		        coalesce(m.class, ''), coalesce(m.phase, ''), m.execution_score
		 from fight_metrics m
		 where m.player_key = $1 and m.state <> 'removed'
		 order by m.fought_at desc limit $2`, key, HistoryLimit)
	if err != nil {
		return Character{}, false, fmt.Errorf("rankings: read character history: %w", err)
	}
	defer rows.Close()
	best := map[[2]int64]CharacterBest{}
	seen := map[string]BuildSeen{}
	found := out.Character.Class != ""
	// phases are remembered per history row so the percentiles can be
	// read in one go once every row is in.
	phases := []bracketKey{}
	for rows.Next() {
		var (
			f                       CharacterFight
			role, class, at         string
			dps, hps                *float64
			taken                   *int64
			encounterID, difficulty *int64
		)
		if err := rows.Scan(&encounterID, &difficulty, &f.Kill, &role, &dps, &hps, &taken,
			&f.Spec, &f.TalentSplit, &f.DurationMS, &f.FoughtAt, &f.ReportID, &f.FightIndex,
			&class, &at, &f.ExecutionScore); err != nil {
			return Character{}, false, fmt.Errorf("rankings: scan character history: %w", err)
		}
		found = true
		if out.Character.Class == "" {
			out.Character.Class = class
		}
		if encounterID != nil {
			f.EncounterID = *encounterID
		}
		if difficulty != nil {
			f.Difficulty = *difficulty
		}
		f.Metric = metricOf(role)
		f.Value = pickValue(f.Metric, dps, hps, taken)
		out.History = append(out.History, f)
		phases = append(phases, bracketKey{
			EncounterID: f.EncounterID, Difficulty: f.Difficulty,
			Spec: f.Spec, Phase: at, Metric: f.Metric,
		})

		if f.Kill {
			k := [2]int64{f.EncounterID, f.Difficulty}
			if b, ok := best[k]; !ok || f.Value > b.Value {
				best[k] = CharacterBest{
					EncounterID: f.EncounterID, Difficulty: f.Difficulty, Metric: f.Metric,
					Value: f.Value, Spec: f.Spec, FoughtAt: f.FoughtAt,
					ReportID: f.ReportID, FightIndex: f.FightIndex,
					Percentile: nil,
				}
			}
		}
		if f.TalentSplit != "" {
			if b, ok := seen[f.TalentSplit]; !ok || f.FoughtAt.Before(b.FirstSeen) {
				seen[f.TalentSplit] = BuildSeen{TalentSplit: f.TalentSplit, Spec: f.Spec, FirstSeen: f.FoughtAt}
			}
		}
	}
	if err := rows.Err(); err != nil {
		return Character{}, false, err
	}

	// One query for the encounter names and one for the digests, then
	// every row is named and placed in memory.
	ids := make([]int64, 0, len(out.History))
	for _, k := range phases {
		ids = append(ids, k.EncounterID)
	}
	names, err := s.EncounterNames(ctx, ids)
	if err != nil {
		return Character{}, false, err
	}
	digests, err := s.digestsFor(ctx, ids)
	if err != nil {
		return Character{}, false, err
	}
	byBracket := map[[2]int64]bracketKey{}
	for i := range out.History {
		f := &out.History[i]
		f.Encounter = names[f.EncounterID]
		f.Percentile = percentileOf(digests, phases[i], f.Value)
		if f.Kill {
			byBracket[[2]int64{f.EncounterID, f.Difficulty}] = phases[i]
		}
	}
	for k, b := range best {
		b.Encounter = names[b.EncounterID]
		if key, ok := byBracket[k]; ok {
			key.Spec, key.Metric = b.Spec, b.Metric
			b.Percentile = percentileOf(digests, key, b.Value)
		}
		out.Best = append(out.Best, b)
	}
	sortBest(out.Best)
	for _, b := range seen {
		out.BuildsSeen = append(out.BuildsSeen, b)
	}
	sortBuilds(out.BuildsSeen)
	return out, found, nil
}

// pickValue reads the column a metric lives in, treating a missing
// value as zero.
func pickValue(metric string, dps, hps *float64, taken *int64) float64 {
	switch metric {
	case MetricHPS:
		if hps != nil {
			return *hps
		}
	case MetricDamageTaken:
		if taken != nil {
			return float64(*taken)
		}
	default:
		if dps != nil {
			return *dps
		}
	}
	return 0
}
