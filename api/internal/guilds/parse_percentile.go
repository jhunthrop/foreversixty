// api/internal/guilds/parse_percentile.go
//
// Where a parse stands among every ranked kill of its bracket, 0 to 100: the figure the
// rankings show (api/internal/rankings: percentileOf over percentile_digests). The raid
// report's parses and the roster's best_parse both read it. This package cannot import
// rankings (rankings imports guilds), so it reads the same table with the same five bracket
// axes and places the value with the same digest.Placement.
package guilds

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/digest"
)

// parseBracket is one percentile_digests row's key: the axes rankings.Store writes a kill's
// digest under (fight_metrics' encounter_id, difficulty, spec and phase, plus the metric).
type parseBracket struct {
	EncounterID int64
	Difficulty  int64
	Spec        string
	Phase       string
	Metric      string
}

// bracketColumns selects a fight_metrics row `m`'s difficulty, spec and phase the way
// rankings keys its digests (a missing axis is 0 or empty); scan them with scanBracket.
const bracketColumns = `coalesce(m.difficulty, 0), coalesce(m.spec, ''), coalesce(m.phase, '')`

// bracketOf completes a bracket read through bracketColumns with the row's encounter and the
// metric its role is judged on.
func bracketOf(b parseBracket, encounterID int64, metric string) parseBracket {
	b.EncounterID, b.Metric = encounterID, metric
	return b
}

// parseDigests reads every digest the brackets name in one query, however many rows share
// them. A bracket with no digest is simply absent from the map.
func (s *Store) parseDigests(ctx context.Context, brackets []parseBracket) (map[parseBracket]*digest.Digest, error) {
	out := map[parseBracket]*digest.Digest{}
	if len(brackets) == 0 {
		return out, nil
	}
	seen := map[parseBracket]bool{}
	var encounters, difficulties []int64
	var specs, phases, metrics []string
	for _, b := range brackets {
		if seen[b] {
			continue
		}
		seen[b] = true
		encounters, difficulties = append(encounters, b.EncounterID), append(difficulties, b.Difficulty)
		specs, phases, metrics = append(specs, b.Spec), append(phases, b.Phase), append(metrics, b.Metric)
	}
	rows, err := s.Pool.Query(ctx, `
		select d.encounter_id, d.difficulty, d.spec, d.phase, d.metric, d.digest
		from percentile_digests d
		join unnest($1::bigint[], $2::bigint[], $3::text[], $4::text[], $5::text[]) as k(e, df, sp, ph, mt)
		  on d.encounter_id = k.e and d.difficulty = k.df and d.spec = k.sp and d.phase = k.ph and d.metric = k.mt`,
		encounters, difficulties, specs, phases, metrics)
	if err != nil {
		return nil, fmt.Errorf("guilds: parse digests: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			k   parseBracket
			raw []byte
		)
		if err := rows.Scan(&k.EncounterID, &k.Difficulty, &k.Spec, &k.Phase, &k.Metric, &raw); err != nil {
			return nil, fmt.Errorf("guilds: parse digests: %w", err)
		}
		d, err := digest.Unmarshal(raw)
		if err != nil {
			return nil, fmt.Errorf("guilds: parse digests: %w", err)
		}
		out[k] = d
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: parse digests: %w", err)
	}
	return out, nil
}

// placeParse is value's percentile in its bracket, nil when the bracket has no digest or an
// empty one (rankings' percentileOf, the same rule).
func placeParse(digests map[parseBracket]*digest.Digest, k parseBracket, value float64) *float64 {
	d, ok := digests[k]
	if !ok || d.Count() == 0 {
		return nil
	}
	p := d.Placement(value) * 100
	return &p
}
