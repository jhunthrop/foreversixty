// api/internal/guilds/raids_report_parses.go
//
// The report's parses: per raider per boss kill, the figure every guild surface reads
// (parse_source.go: DPS, HPS for a healer) and the fight's Forever rating when one was computed.
// Best first within each kill. A member's report keeps their own rows (ownOnly).
package guilds

import (
	"context"
	"fmt"
	"sort"
)

// ParseRow is one raider's parse on one boss kill. Rating is the fight's overall rating, null
// when none was computed (an insufficient or not-yet-rated fight).
type ParseRow struct {
	CharacterKey string   `json:"character_key"`
	Name         string   `json:"name"`
	Class        string   `json:"class"`
	Role         string   `json:"role"`
	EncounterID  int      `json:"encounter_id"`
	Encounter    string   `json:"encounter"`
	FightIndex   int      `json:"fight_index"`
	Metric       string   `json:"metric"`
	Value        float64  `json:"value"`
	Rating       *float64 `json:"rating"`
	own          bool
}

func (s *Store) reportParses(ctx context.Context, vc raidViewContext, reportID string) ([]ParseRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select m.player_key, coalesce(m.player_name, ''), coalesce(lower(m.class), ''), coalesce(m.role, ''),
		       m.encounter_id, coalesce(f.name, ''), m.fight_index, m.metric_dps::float8, m.metric_hps::float8,
		       (select rs.overall::float8 from rating_scores rs
		         where rs.report_id = m.report_id and rs.fight_index = m.fight_index and rs.player_key = m.player_key
		           and not coalesce(rs.insufficient, false) limit 1)
		from fight_metrics m
		join fights f on f.report_id = m.report_id and f.fight_index = m.fight_index
		where m.report_id = $1 and `+bossKillMetrics, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report parses: %w", err)
	}
	defer rows.Close()
	out := []ParseRow{}
	for rows.Next() {
		var p ParseRow
		var dps, hps *float64
		if err := rows.Scan(&p.CharacterKey, &p.Name, &p.Class, &p.Role, &p.EncounterID, &p.Encounter, &p.FightIndex,
			&dps, &hps, &p.Rating); err != nil {
			return nil, fmt.Errorf("guilds: report parses: %w", err)
		}
		p.Metric = roleMetric(p.Role)
		v := dps
		if p.Metric == "hps" {
			v = hps
		}
		if v != nil {
			p.Value = *v
		}
		p.own = vc.ownsKey(p.CharacterKey)
		out = append(out, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: report parses: %w", err)
	}
	sortParses(out)
	return out, nil
}

// sortParses orders the rows by kill, then healers apart from damage, then best first.
func sortParses(out []ParseRow) {
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.FightIndex != b.FightIndex {
			return a.FightIndex < b.FightIndex
		}
		if a.Metric != b.Metric {
			return a.Metric < b.Metric // dps before hps
		}
		if a.Value != b.Value {
			return a.Value > b.Value
		}
		return a.Name < b.Name
	})
}

// topParses is the n best damage parses of the night across its kills, one per raider.
func topParses(rows []ParseRow, n int) []ParseRow {
	best := map[string]ParseRow{}
	for _, r := range rows {
		if r.Metric != "dps" {
			continue
		}
		if cur, ok := best[r.CharacterKey]; !ok || r.Value > cur.Value {
			best[r.CharacterKey] = r
		}
	}
	out := make([]ParseRow, 0, len(best))
	for _, r := range best {
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Value != out[j].Value {
			return out[i].Value > out[j].Value
		}
		return out[i].Name < out[j].Name
	})
	if len(out) > n {
		out = out[:n]
	}
	return out
}
