package activity

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// boardSQL ranks every spec's characters on every band: each character's best kill per
// encounter on the band's instances (public, unflagged reports, rows not removed, accounts
// not anonymised), its percentile among its spec on that encounter, then the average of
// those percentiles. $1 and $2 are the zone names and their bands.
var boardSQL = `with zb(zone, band) as (select * from unnest($1::text[], $2::int[])),
best as (
  select distinct on (m.player_key, m.encounter_id) m.player_key, m.player_name, m.class, m.spec, zb.band,
         m.encounter_id, case when m.role = 'healer' then m.metric_hps else m.metric_dps end as val
  from fight_metrics m
  join reports r on r.id = m.report_id
  join zb on zb.zone = coalesce(nullif(r.log_zone, ''), r.zone)
  left join characters c on c.key = m.player_key
  left join users u on u.id = c.user_id
  where m.kill and m.state <> 'removed' and ` + presence.EvidenceReport("r") + `
    and m.spec is not null and m.class is not null and not coalesce(u.anonymize, false)
    and (case when m.role = 'healer' then m.metric_hps else m.metric_dps end) > 0
  order by m.player_key, m.encounter_id, val desc
),
pct as (
  select b.*, 100.0 * (count(*) over g - rank() over (g order by b.val desc) + 1) / count(*) over g as p
  from best b window g as (partition by lower(b.class), lower(b.spec), b.band, b.encounter_id)
)
select lower(class), spec, band, player_key, max(player_name), avg(p), count(*)
from pct group by lower(class), spec, band, player_key`

// BoardRow is one character's place on a spec's band.
type BoardRow struct {
	Spec, Key, Name, Class string
	Band, Rank             int
	Score                  float64
	kills                  int
}

type boardID struct {
	spec string
	band int
}

// computeBoards reads every board as it stands now.
func (s *Store) computeBoards(ctx context.Context, zoneBands map[string]int) (map[boardID][]BoardRow, error) {
	zones, bands := make([]string, 0, len(zoneBands)), make([]int32, 0, len(zoneBands))
	for z, b := range zoneBands {
		zones, bands = append(zones, z), append(bands, int32(b))
	}
	rows, err := s.Pool.Query(ctx, boardSQL, zones, bands)
	if err != nil {
		return nil, fmt.Errorf("activity: boards: %w", err)
	}
	defer rows.Close()
	out := map[boardID][]BoardRow{}
	for rows.Next() {
		var r BoardRow
		var spec string
		if err := rows.Scan(&r.Class, &spec, &r.Band, &r.Key, &r.Name, &r.Score, &r.kills); err != nil {
			return nil, fmt.Errorf("activity: scan boards: %w", err)
		}
		key, ok := SpecKey(r.Class, spec)
		if !ok {
			continue
		}
		r.Spec = key
		id := boardID{key, r.Band}
		out[id] = append(out[id], r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for id, b := range out {
		sort.SliceStable(b, func(i, j int) bool {
			if b[i].Score != b[j].Score {
				return b[i].Score > b[j].Score
			}
			if b[i].kills != b[j].kills {
				return b[i].kills > b[j].kills
			}
			return b[i].Key < b[j].Key
		})
		for i := range b {
			b[i].Rank = i + 1
		}
		out[id] = b
	}
	return out, nil
}

// boardOf reads one stored week's ranks of every board: character key to rank.
func (s *Store) storedRanks(ctx context.Context, week time.Time) (map[boardID]map[string]int, error) {
	rows, err := s.Pool.Query(ctx, `select spec, band, character_key, rank from spec_board_weeks where week = $1`, week)
	if err != nil {
		return nil, fmt.Errorf("activity: stored boards: %w", err)
	}
	defer rows.Close()
	out := map[boardID]map[string]int{}
	for rows.Next() {
		var id boardID
		var key string
		var rank int
		if err := rows.Scan(&id.spec, &id.band, &key, &rank); err != nil {
			return nil, fmt.Errorf("activity: scan stored boards: %w", err)
		}
		if out[id] == nil {
			out[id] = map[string]int{}
		}
		out[id][key] = rank
	}
	return out, rows.Err()
}

// replaceWeek stores boards as week's rows, replacing what the week held.
func (s *Store) replaceWeek(ctx context.Context, week time.Time, boards map[boardID][]BoardRow) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("activity: replace boards: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `delete from spec_board_weeks where week = $1`, week); err != nil {
		return fmt.Errorf("activity: clear boards: %w", err)
	}
	var specsCol, keys, names, classes []string
	var bands, ranks []int32
	var scores []float64
	for _, b := range boards {
		for _, r := range b {
			specsCol, keys, names, classes = append(specsCol, r.Spec), append(keys, r.Key), append(names, r.Name), append(classes, r.Class)
			bands, ranks, scores = append(bands, int32(r.Band)), append(ranks, int32(r.Rank)), append(scores, r.Score)
		}
	}
	if _, err := tx.Exec(ctx, `insert into spec_board_weeks (week, spec, band, character_key, name, class, rank, score)
		select $1, * from unnest($2::text[], $3::int[], $4::text[], $5::text[], $6::text[], $7::int[], $8::float8[])`,
		week, specsCol, bands, keys, names, classes, ranks, scores); err != nil {
		return fmt.Errorf("activity: write boards: %w", err)
	}
	return tx.Commit(ctx)
}
