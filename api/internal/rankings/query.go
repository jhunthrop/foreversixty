package rankings

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
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
	// Deduped is true when each row is a different character's best
	// kill, so Total counts characters and not kills.
	Deduped bool `json:"deduped"`
	// CharacterRank is where the requested character stands on this
	// board, absent when none was asked for or it has no ranked kill.
	CharacterRank *CharacterRank `json:"character_rank,omitempty"`
}

// CharacterRank is one character's place on a board: its best kill, with
// its true rank, whichever page that falls on.
type CharacterRank struct {
	Rank int `json:"rank"`
	Row  Row `json:"row"`
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
	// Distinct is DistinctCharacter to rank each character's best kill
	// only, or empty to rank every kill.
	Distinct string
	// Character is a player key (region/ruleset/name) to place on the
	// board; it needs Distinct, or its rank would not match any row.
	Character string
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
		// A spec is one name however it was written: the addon and a log export write
		// "BeastMastery", the talent trees and the site say "Beast Mastery". Compared with the
		// spaces gone, a board filtered on either spelling holds the same fights.
		add("replace(lower(m.spec), ' ', '') = replace(lower($%d), ' ', '')", q.Spec)
	}
	if q.Class != "" {
		add("lower(m.class) = lower($%d)", q.Class)
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

// DistinctCharacter is the `distinct` value that ranks characters, not kills.
const DistinctCharacter = "character"

// rankedCTE ranks the filtered kills. When q.Distinct is set each character
// keeps only its best kill, so the rank is a place among characters. The
// metric is exposed as `val`, and `rk` is the 1-based place.
func rankedCTE(q Query, where, column string) string {
	distinctOn, lead := "", ""
	if q.Distinct == DistinctCharacter {
		distinctOn, lead = "distinct on (m.player_key)", "m.player_key, "
	}
	return `with cand as (
		   select ` + distinctOn + ` m.player_key, m.player_name, coalesce(m.class, '') as class,
		          coalesce(m.spec, '') as spec, ` + column + ` as val, coalesce(m.ilvl, 0) as ilvl,
		          coalesce(m.size, 0) as size, m.fought_at, coalesce(m.duration_ms, 0) as duration_ms,
		          coalesce(m.talent_split, '') as talent_split, m.trinkets, m.buff_count,
		          m.report_id, m.fight_index, m.state, m.execution_score
		   from fight_metrics m where ` + where + `
		   order by ` + lead + column + ` desc nulls last, m.fought_at desc, m.report_id, m.fight_index, m.player_key
		 ), ranked as (
		   select cand.*, row_number() over (order by val desc nulls last, fought_at desc,
		          report_id, fight_index, player_key) as rk from cand
		 )
		 select c.rk, c.player_key, c.player_name, c.class, c.spec, c.val, c.ilvl, c.size, c.fought_at,
		        c.duration_ms, c.talent_split, c.trinkets, c.buff_count, c.report_id, c.fight_index,
		        c.state, c.execution_score, g.name, g.region, g.ruleset
		 from ranked c
		 left join reports r on r.id = c.report_id
		 left join guilds g on g.id = r.guild_id
		 `
}

// scanRanked reads rows shaped by rankedCTE.
func scanRanked(rows pgx.Rows) ([]Row, error) {
	out := []Row{}
	for rows.Next() {
		var (
			r                          Row
			value                      *float64
			guildName, region, ruleset *string
		)
		if err := rows.Scan(&r.Rank, &r.Player.Key, &r.Player.Name, &r.Player.Class, &r.Player.Spec,
			&value, &r.Ilvl, &r.Size, &r.FoughtAt, &r.DurationMS, &r.TalentSplit, &r.Trinkets,
			&r.BuffCount, &r.ReportID, &r.FightIndex, &r.State, &r.ExecutionScore,
			&guildName, &region, &ruleset); err != nil {
			return nil, fmt.Errorf("rankings: scan: %w", err)
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
		out = append(out, r)
	}
	return out, rows.Err()
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
	deduped := q.Distinct == DistinctCharacter
	page := Page{Rows: []Row{}, Page: q.Page, PerPage: PerPage, UpdatedAt: now.UTC(), Deduped: deduped}
	count := `select count(*) from fight_metrics m where ` + where
	if deduped {
		count = `select count(distinct m.player_key) from fight_metrics m where ` + where
	}
	if err := s.Pool.QueryRow(ctx, count, args...).Scan(&page.Total); err != nil {
		return Page{}, fmt.Errorf("rankings: count: %w", err)
	}
	cte := rankedCTE(q, where, metricColumn(q.Metric))
	offset := (q.Page - 1) * PerPage
	rows, err := s.Pool.Query(ctx, cte+fmt.Sprintf(`where c.rk > %d and c.rk <= %d order by c.rk`, offset, offset+PerPage), args...)
	if err != nil {
		return Page{}, fmt.Errorf("rankings: query: %w", err)
	}
	defer rows.Close()
	if page.Rows, err = scanRanked(rows); err != nil {
		return Page{}, err
	}
	rows.Close()
	if q.Character != "" {
		if page.CharacterRank, err = s.characterRank(ctx, cte, args, q.Character); err != nil {
			return Page{}, err
		}
	}
	return page, nil
}

// characterRank places one character on the filtered board: its best kill
// and the rank that kill holds, or nil when it has no ranked kill in the
// cohort. cte must be a distinct-character ranking.
func (s *Store) characterRank(ctx context.Context, cte string, args []any, key string) (*CharacterRank, error) {
	rows, err := s.Pool.Query(ctx, cte+fmt.Sprintf(`where c.player_key = $%d`, len(args)+1), append(args, key)...)
	if err != nil {
		return nil, fmt.Errorf("rankings: character rank: %w", err)
	}
	defer rows.Close()
	found, err := scanRanked(rows)
	if err != nil || len(found) == 0 {
		return nil, err
	}
	return &CharacterRank{Rank: found[0].Rank, Row: found[0]}, nil
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
	ForeverScore
	// ScoreHistory is the Forever Score's history (Store.scoreHistory), oldest first; empty,
	// never null, when no rated fight carries a snapshot.
	ScoreHistory []ScorePoint `json:"score_history"`
	// Gear is the gear panel (rating.Service.ScoreAndGear): the 17 export slots in the
	// export's order, empty, never null, when the character has no export. GearGapTotal is
	// GearGapTotal of its rows: null when no row has a gap.
	Gear         []GearSlot `json:"gear"`
	GearGapTotal *float64   `json:"gear_gap_total"`
	// Supporter is true when the character's owning account shows the supporter mark
	// (Service.supporterOf); false for an unowned character.
	Supporter bool `json:"supporter"`
	// ownerID is characters.user_id, read to resolve Supporter and never marshaled.
	ownerID *int64
	// GearSpec is the spec the gear picks are for and GearSpecSource where it came from
	// ("fight", "talents" or "default"); null when the panel has no spec.
	GearSpec       *string `json:"gear_spec"`
	GearSpecSource *string `json:"gear_spec_source"`
}

// ForeverScore is the character page's Forever Score (api/internal/rating's ScoreV2, the
// formula's only home). Every field is null for a character with no rated fight: "no logs
// yet" is never shown as a zero score.
type ForeverScore struct {
	Score      *int             `json:"score"`
	Tier       *string          `json:"tier"`
	Components *ScoreComponents `json:"components"`
	// AsOf is the newest rated fight's date, YYYY-MM-DD.
	AsOf       *string          `json:"as_of"`
	Experience *ScoreExperience `json:"experience"`
	// ComponentCaps are each component's ceiling, so a bar is drawn as value over cap without
	// the reader knowing the formula; null with the score.
	ComponentCaps *ScoreComponents `json:"component_caps"`
}

// ScoreComponents are the score's five parts; they sum to the score.
type ScoreComponents struct {
	Gear       int `json:"gear"`
	Experience int `json:"experience"`
	Parses     int `json:"parses"`
	Attendance int `json:"attendance"`
	Discipline int `json:"discipline"`
}

// ScoreExperience is what the experience component counts, across the account's linked
// characters.
type ScoreExperience struct {
	Dungeons  int `json:"dungeons"`
	RaidKills int `json:"raid_kills"`
	Alts      int `json:"alts"`
}

// CharacterRef names a character.
type CharacterRef struct {
	Key     string `json:"key"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Name    string `json:"name"`
	Class   string `json:"class,omitempty"`
	// AvatarURL and RenderURL are Blizzard's own render-CDN images for
	// this character (.superpowers/account-visual-brief.md §A4), omitted
	// when the Battle.net import has not captured them.
	AvatarURL string `json:"avatar_url,omitempty"`
	RenderURL string `json:"render_url,omitempty"`
	// Guild, Level, Race and Realm say who the character is in game (Store.identity); each
	// is null when no source knows it.
	Guild *CharacterGuild `json:"guild"`
	Level *int            `json:"level"`
	Race  *string         `json:"race"`
	Realm *string         `json:"realm"`
}

// HistoryLimit is how many of a character's fights the page shows.
const HistoryLimit = 100

// Character assembles the public character page: CharacterFights, then who the character
// is (identity) and its score history (scoreHistory). It reports false when the character
// has no ranked fights and no row of their own, and also when its account is anonymised
// (the rule GET .../rating already applies), so an anonymised character reads exactly as
// one nobody has logged.
func (s *Store) Character(ctx context.Context, region, ruleset, name string) (Character, bool, error) {
	key := strings.ToLower(region) + "/" + strings.ToLower(ruleset) + "/" + strings.ToLower(name)
	anon, err := s.anonymized(ctx, key)
	if err != nil || anon {
		return Character{}, false, err
	}
	out, found, err := s.CharacterFights(ctx, region, ruleset, name)
	if err != nil || !found {
		return out, found, err
	}
	if err := s.identity(ctx, &out.Character); err != nil {
		return Character{}, false, err
	}
	if out.ScoreHistory, err = s.scoreHistory(ctx, key); err != nil {
		return Character{}, false, err
	}
	return out, true, nil
}

// CharacterFights is the fights half of a character page: its name and class, bests,
// history and builds seen, with no anonymise check and no identity or score history. The
// Forever Score's parses component reads its bests (rating.Store.meanBestParse). Its rows
// come from reports that are evidence (presence.EvidenceReport): a flagged report's rows,
// which the moderation flow withdraws, are never shown or scored meanwhile. It
// reports false when the character has no ranked fights and no row of their own.
func (s *Store) CharacterFights(ctx context.Context, region, ruleset, name string) (Character, bool, error) {
	key := strings.ToLower(region) + "/" + strings.ToLower(ruleset) + "/" + strings.ToLower(name)
	out := Character{
		Character:    CharacterRef{Key: key, Region: region, Ruleset: ruleset, Name: name},
		Best:         []CharacterBest{},
		History:      []CharacterFight{},
		BuildsSeen:   []BuildSeen{},
		ScoreHistory: []ScorePoint{},
		Gear:         []GearSlot{},
	}
	var class *string
	err := s.Pool.QueryRow(ctx,
		`select name, class, coalesce(avatar_url, ''), coalesce(render_url, ''), user_id from characters where key = $1`, key).
		Scan(&out.Character.Name, &class, &out.Character.AvatarURL, &out.Character.RenderURL, &out.ownerID)
	if err == nil && class != nil {
		out.Character.Class = *class
	} else if err != nil && err != pgx.ErrNoRows {
		return Character{}, false, fmt.Errorf("rankings: read character: %w", err)
	}

	read, err := s.readHistory(ctx, key, presence.EvidenceReport("r"), HistoryLimit)
	if err != nil {
		return Character{}, false, err
	}
	best := map[[2]int64]CharacterBest{}
	seen := map[string]BuildSeen{}
	found := out.Character.Class != ""
	byBracket := map[[2]int64]bracketKey{}
	for _, h := range read.rows {
		f := h.fight
		found = true
		if out.Character.Class == "" {
			out.Character.Class = h.class
		}
		out.History = append(out.History, f)
		if f.Kill {
			k := [2]int64{f.EncounterID, f.Difficulty}
			byBracket[k] = h.bracket
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
	for k, b := range best {
		b.Encounter = read.names[b.EncounterID]
		if key, ok := byBracket[k]; ok {
			key.Spec, key.Metric = b.Spec, b.Metric
			b.Percentile = percentileOf(read.digests, key, b.Value)
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
