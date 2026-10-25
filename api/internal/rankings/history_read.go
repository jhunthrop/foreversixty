package rankings

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/digest"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// historyRow is one ranked fight as the history reads it: the page's row, the class it was
// fought as, the visibility of its report and the percentile bracket its value is placed in.
type historyRow struct {
	fight      CharacterFight
	class      string
	visibility string
	bracket    bracketKey
}

// historyRead is a set of history rows with the encounter names and digests that named and
// placed them, so a caller that needs more placements (the bests) reuses the same reads.
type historyRead struct {
	rows    []historyRow
	names   map[int64]string
	digests map[bracketKey]*digest.Digest
}

// readHistory reads a character's ranked fights, newest first, up to limit, from the reports
// that scope selects (a SQL condition on the reports row aliased r; its placeholders start at
// $3 and take scopeArgs). Every row is named and placed in memory after one query for the
// names and one for the digests, so a hundred fights cost no hundred round trips.
func (s *Store) readHistory(ctx context.Context, key, scope string, limit int, scopeArgs ...any) (historyRead, error) {
	rows, err := s.Pool.Query(ctx,
		`select m.encounter_id, m.difficulty, m.kill, m.role, m.metric_dps, m.metric_hps,
		        m.damage_taken, coalesce(m.spec, ''), coalesce(m.talent_split, ''),
		        coalesce(m.duration_ms, 0), m.fought_at, m.report_id, m.fight_index,
		        coalesce(m.class, ''), coalesce(m.phase, ''), m.execution_score, r.visibility
		 from fight_metrics m join reports r on r.id = m.report_id
		 where m.player_key = $1 and m.state <> 'removed' and `+scope+`
		 order by m.fought_at desc, m.report_id desc, m.fight_index desc limit $2`,
		append([]any{key, limit}, scopeArgs...)...)
	if err != nil {
		return historyRead{}, fmt.Errorf("rankings: read character history: %w", err)
	}
	defer rows.Close()
	var out historyRead
	for rows.Next() {
		h, err := scanHistoryRow(rows)
		if err != nil {
			return historyRead{}, err
		}
		out.rows = append(out.rows, h)
	}
	if err := rows.Err(); err != nil {
		return historyRead{}, err
	}
	return out, s.placeHistory(ctx, &out)
}

// scanHistoryRow reads one readHistory row and gives it its metric and value.
func scanHistoryRow(rows pgx.Rows) (historyRow, error) {
	var (
		h                       historyRow
		role, phase             string
		dps, hps                *float64
		taken                   *int64
		encounterID, difficulty *int64
	)
	f := &h.fight
	if err := rows.Scan(&encounterID, &difficulty, &f.Kill, &role, &dps, &hps, &taken,
		&f.Spec, &f.TalentSplit, &f.DurationMS, &f.FoughtAt, &f.ReportID, &f.FightIndex,
		&h.class, &phase, &f.ExecutionScore, &h.visibility); err != nil {
		return historyRow{}, fmt.Errorf("rankings: scan character history: %w", err)
	}
	if encounterID != nil {
		f.EncounterID = *encounterID
	}
	if difficulty != nil {
		f.Difficulty = *difficulty
	}
	f.Metric = metricOf(role)
	f.Value = pickValue(f.Metric, dps, hps, taken)
	h.bracket = bracketKey{
		EncounterID: f.EncounterID, Difficulty: f.Difficulty,
		Spec: f.Spec, Phase: phase, Metric: f.Metric,
	}
	return h, nil
}

// placeHistory names every row's encounter and places its value in its bracket.
func (s *Store) placeHistory(ctx context.Context, read *historyRead) error {
	ids := make([]int64, 0, len(read.rows))
	for _, h := range read.rows {
		ids = append(ids, h.fight.EncounterID)
	}
	var err error
	if read.names, err = s.EncounterNames(ctx, ids); err != nil {
		return err
	}
	if read.digests, err = s.digestsFor(ctx, ids); err != nil {
		return err
	}
	for i := range read.rows {
		h := &read.rows[i]
		h.fight.Encounter = read.names[h.fight.EncounterID]
		h.fight.Percentile = percentileOf(read.digests, h.bracket, h.fight.Value)
	}
	return nil
}

// FightRef names one fight of a report.
type FightRef struct {
	ReportID   string
	FightIndex int
}

// OwnerFight is a character's history row with the visibility of the report it is in.
type OwnerFight struct {
	CharacterFight
	Visibility string `json:"visibility"`
}

// OwnerFights is the character's history rows for exactly the fights in refs, newest first,
// in any report that is not flagged: the owner sees their private and guild reports, so the
// public page's visibility rule is the one thing dropped. A fight without a ranked row is
// absent. Never nil.
func (s *Store) OwnerFights(ctx context.Context, key string, refs []FightRef) ([]OwnerFight, error) {
	out := []OwnerFight{}
	if len(refs) == 0 {
		return out, nil
	}
	reports := make([]string, len(refs))
	indexes := make([]int32, len(refs))
	for i, r := range refs {
		reports[i], indexes[i] = r.ReportID, int32(r.FightIndex)
	}
	read, err := s.readHistory(ctx, key,
		presence.NotFlagged("r")+` and (m.report_id, m.fight_index) in (select * from unnest($3::text[], $4::int[]))`,
		len(refs), reports, indexes)
	if err != nil {
		return nil, err
	}
	for _, h := range read.rows {
		out = append(out, OwnerFight{CharacterFight: h.fight, Visibility: h.visibility})
	}
	return out, nil
}
