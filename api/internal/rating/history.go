package rating

// The owner's history: a character's Forever Score and parse over time, every rated fight in
// every report the owner can see (public, unlisted, guild and private), never a flagged one.
// It backs GET /v1/characters/{region}/{ruleset}/{name}/history, the contract in
// design/specs/2026-10-24-character-history.md section 10. The public page's score_history
// (rankings.scoreHistory) is a different, public-only, one-point-per-date read and is untouched.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// The closed set of reasons a stretch of the score line is empty.
const (
	GapNoRatedFights = "No rated fights"
	GapReportFlagged = "Report flagged"
	GapScoreNotKept  = "Score not kept"
)

const (
	// GapDays is how many days between two dated score points make a gap: more than this.
	GapDays = 14
	// HistoryMaxFights caps the rated fights the history reads: the newest.
	HistoryMaxFights = 1000

	dateLayout      = "2006-01-02"
	outputComponent = "output"
)

// HistoryPoint is one rated fight on the score line.
type HistoryPoint struct {
	At          time.Time
	Score       int
	Tier        string
	ReportID    string
	FightIndex  int
	Encounter   string
	EncounterID int64
	// Difficulty is DifficultyName of the fight's difficulty id.
	Difficulty string
	// ParsePercentile is the fight's output component percentile, nil when it has none.
	ParsePercentile *float64
	Kill            bool
	// Visibility is the report's: public, guild or private (the query reads an unlisted report as public).
	Visibility string
}

// HistoryParsePoint is one kill's parse on a boss's line.
type HistoryParsePoint struct {
	At              time.Time
	ParsePercentile float64
	ReportID        string
	FightIndex      int
}

// HistoryEncounter is one boss at one difficulty: its kills with a parse, oldest first.
type HistoryEncounter struct {
	Encounter   string
	EncounterID int64
	Difficulty  string
	Points      []HistoryParsePoint
}

// HistoryGap is a stretch with no score point: From and To are the UTC dates (YYYY-MM-DD)
// of the points either side, Reason one of the Gap* constants.
type HistoryGap struct {
	From, To string
	Reason   string
}

// History is the owner's chart data. Every slice is non-nil.
type History struct {
	// AsOf is the UTC date of the newest score point, today when there is none.
	AsOf       string
	Score      []HistoryPoint
	Encounters []HistoryEncounter
	Gaps       []HistoryGap
	// Fights are the rows behind every point, newest first, one per (report, fight).
	Fights []rankings.OwnerFight
}

// ReadOwnerHistory reads the history of the character playerKey. Authorising the caller is the
// handler's job; this reads whatever it is asked for.
func (s *Store) ReadOwnerHistory(ctx context.Context, playerKey string) (History, error) {
	rated, err := s.ratedPoints(ctx, playerKey)
	if err != nil {
		return History{}, err
	}
	names, err := (&rankings.Store{Pool: s.Pool}).EncounterNames(ctx, encounterIDs(rated))
	if err != nil {
		return History{}, err
	}
	score := make([]HistoryPoint, len(rated))
	for i, r := range rated {
		score[i] = r.point(names)
	}
	marks, err := s.stretchMarks(ctx, playerKey, score)
	if err != nil {
		return History{}, err
	}
	fights, err := (&rankings.Store{Pool: s.Pool}).OwnerFights(ctx, playerKey, fightRefs(score))
	if err != nil {
		return History{}, err
	}
	return History{
		AsOf:       asOf(score, time.Now()),
		Score:      score,
		Encounters: encountersOf(score, names),
		Gaps:       gapsFor(scoreDates(score), marks),
		Fights:     fights,
	}, nil
}

// ratedPoint is one rating_scores row with a score snapshot, as read.
type ratedPoint struct {
	at          time.Time
	score       int
	tier        string
	reportID    string
	fightIndex  int
	encounterID int64
	difficulty  int64
	kill        bool
	components  json.RawMessage
	visibility  string
}

func (r ratedPoint) point(names map[int64]string) HistoryPoint {
	return HistoryPoint{
		At: r.at, Score: r.score, Tier: r.tier, ReportID: r.reportID, FightIndex: r.fightIndex,
		Encounter: names[r.encounterID], EncounterID: r.encounterID,
		Difficulty: DifficultyName(r.difficulty), ParsePercentile: outputPercentile(r.components),
		Kill: r.kill, Visibility: r.visibility,
	}
}

// ratedPoints reads the newest HistoryMaxFights rated fights that kept a score snapshot, in
// reports that are not flagged, oldest first.
func (s *Store) ratedPoints(ctx context.Context, playerKey string) ([]ratedPoint, error) {
	rows, err := s.Pool.Query(ctx,
		`select rs.fought_at, rs.score, rs.tier, rs.report_id, rs.fight_index,
		        coalesce(rs.encounter_id, 0), coalesce(rs.difficulty, 0), rs.kill, rs.components,
		        case when r.visibility = 'unlisted' then 'public' else r.visibility end
		 from rating_scores rs join reports r on r.id = rs.report_id
		 where rs.player_key = $1 and rs.score is not null and rs.tier is not null and `+presence.NotFlagged("r")+`
		 order by rs.fought_at desc, rs.report_id desc, rs.fight_index desc limit $2`,
		playerKey, HistoryMaxFights)
	if err != nil {
		return nil, fmt.Errorf("rating: read history %s: %w", playerKey, err)
	}
	defer rows.Close()
	newestFirst := []ratedPoint{}
	for rows.Next() {
		var r ratedPoint
		if err := rows.Scan(&r.at, &r.score, &r.tier, &r.reportID, &r.fightIndex,
			&r.encounterID, &r.difficulty, &r.kill, &r.components, &r.visibility); err != nil {
			return nil, fmt.Errorf("rating: scan history %s: %w", playerKey, err)
		}
		newestFirst = append(newestFirst, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reversed(newestFirst), nil
}

func reversed[T any](xs []T) []T {
	out := make([]T, len(xs))
	for i, x := range xs {
		out[len(xs)-1-i] = x
	}
	return out
}

// outputComponent's percentile out of a stored components array: nil when the component is
// absent, excluded or has none.
func outputPercentile(components json.RawMessage) *float64 {
	var cs []struct {
		Name       string   `json:"name"`
		Percentile *float64 `json:"percentile"`
		Excluded   bool     `json:"excluded"`
	}
	if json.Unmarshal(components, &cs) != nil {
		return nil
	}
	for _, c := range cs {
		if c.Name == outputComponent && !c.Excluded {
			return c.Percentile
		}
	}
	return nil
}

func encounterIDs(rated []ratedPoint) []int64 {
	ids := make([]int64, len(rated))
	for i, r := range rated {
		ids[i] = r.encounterID
	}
	return ids
}

func fightRefs(score []HistoryPoint) []rankings.FightRef {
	refs := make([]rankings.FightRef, len(score))
	for i, p := range score {
		refs[i] = rankings.FightRef{ReportID: p.ReportID, FightIndex: p.FightIndex}
	}
	return refs
}

func asOf(score []HistoryPoint, now time.Time) string {
	if len(score) == 0 {
		return now.UTC().Format(dateLayout)
	}
	return score[len(score)-1].At.UTC().Format(dateLayout)
}

// encountersOf groups the kills that have a parse by (encounter, difficulty): points oldest
// first (score is oldest first), groups newest kill first.
func encountersOf(score []HistoryPoint, names map[int64]string) []HistoryEncounter {
	type groupKey struct {
		id         int64
		difficulty string
	}
	groups := map[groupKey]*HistoryEncounter{}
	order := []groupKey{}
	for _, p := range score {
		if !p.Kill || p.ParsePercentile == nil {
			continue
		}
		k := groupKey{p.EncounterID, p.Difficulty}
		g, ok := groups[k]
		if !ok {
			g = &HistoryEncounter{Encounter: names[p.EncounterID], EncounterID: p.EncounterID, Difficulty: p.Difficulty}
			groups[k] = g
		}
		g.Points = append(g.Points, HistoryParsePoint{
			At: p.At, ParsePercentile: *p.ParsePercentile, ReportID: p.ReportID, FightIndex: p.FightIndex,
		})
		order = moveToEnd(order, k)
	}
	out := make([]HistoryEncounter, 0, len(order))
	for i := len(order) - 1; i >= 0; i-- {
		out = append(out, *groups[order[i]])
	}
	return out
}

// moveToEnd puts k last in order, so the order is by each group's newest kill.
func moveToEnd[T comparable](order []T, k T) []T {
	kept := order[:0]
	for _, o := range order {
		if o != k {
			kept = append(kept, o)
		}
	}
	return append(kept, k)
}

// scoreDates is the distinct UTC dates holding a score point, oldest first.
func scoreDates(score []HistoryPoint) []time.Time {
	dates := []time.Time{}
	for _, p := range score {
		d := utcDate(p.At)
		if len(dates) == 0 || !dates[len(dates)-1].Equal(d) {
			dates = append(dates, d)
		}
	}
	return dates
}

func utcDate(t time.Time) time.Time {
	y, m, d := t.UTC().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// stretchMark is a fight that explains an empty stretch: one in a flagged report, or (not
// Flagged) one whose rating row kept no score snapshot.
type stretchMark struct {
	Date    time.Time
	Flagged bool
}

// stretchMarks reads the marks from the first score point on. A date with no marks is read
// as no rated fights.
func (s *Store) stretchMarks(ctx context.Context, playerKey string, score []HistoryPoint) ([]stretchMark, error) {
	marks := []stretchMark{}
	if len(score) == 0 {
		return marks, nil
	}
	rows, err := s.Pool.Query(ctx,
		`select (rs.fought_at at time zone 'UTC')::date, r.flagged is not null
		 from rating_scores rs join reports r on r.id = rs.report_id
		 where rs.player_key = $1 and rs.fought_at >= $2
		   and (r.flagged is not null or rs.score is null or rs.tier is null)`,
		playerKey, score[0].At)
	if err != nil {
		return nil, fmt.Errorf("rating: read history gaps %s: %w", playerKey, err)
	}
	defer rows.Close()
	for rows.Next() {
		var m stretchMark
		if err := rows.Scan(&m.Date, &m.Flagged); err != nil {
			return nil, fmt.Errorf("rating: scan history gaps %s: %w", playerKey, err)
		}
		marks = append(marks, m)
	}
	return marks, rows.Err()
}

// gapsFor is every stretch of more than GapDays between consecutive dates (UTC midnights,
// oldest first), with the reason the marks strictly between the two dates give: a flagged
// report first, then a fight whose score was not kept, else none. Never nil.
func gapsFor(dates []time.Time, marks []stretchMark) []HistoryGap {
	gaps := []HistoryGap{}
	for i := 1; i < len(dates); i++ {
		from, to := dates[i-1], dates[i]
		if to.Sub(from) <= GapDays*24*time.Hour {
			continue
		}
		gaps = append(gaps, HistoryGap{
			From: from.Format(dateLayout), To: to.Format(dateLayout), Reason: gapReason(from, to, marks),
		})
	}
	return gaps
}

func gapReason(from, to time.Time, marks []stretchMark) string {
	reason := GapNoRatedFights
	for _, m := range marks {
		if !m.Date.After(from) || !m.Date.Before(to) {
			continue
		}
		if m.Flagged {
			return GapReportFlagged
		}
		reason = GapScoreNotKept
	}
	return reason
}

// DifficultyName is the word the API gives a difficulty id: the one-word names of the
// ids a Forever raid is fought on, and "difficulty N" for an id nobody has named yet.
func DifficultyName(id int64) string {
	if name, ok := difficultyNames[id]; ok {
		return name
	}
	return fmt.Sprintf("difficulty %d", id)
}

// difficultyNames is the single source of the difficulty words (the engine's difficultyID).
var difficultyNames = map[int64]string{
	1: "normal", 3: "normal", 4: "normal",
	2: "hard", 5: "hard", 6: "hard",
}
