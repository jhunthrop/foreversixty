package rankings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// RankedEncounter is one row of GET /v1/rankings/encounters (W3-C6): an encounter with ranked
// kills, the zone the newest of its logs named, its usual raid size and its kill count. zone
// and size are null when no log says.
type RankedEncounter struct {
	ID    int64   `json:"id"`
	Slug  string  `json:"slug"`
	Name  string  `json:"name"`
	Zone  *string `json:"zone"`
	Size  *int    `json:"size"`
	Kills int     `json:"kills"`
}

// rankedEncountersQuery counts the kills the boards rank (Query.where's own rule: a kill whose
// rows are not removed), names each encounter by encounterQuery, and reads the zone from the
// newest report that recorded one.
const rankedEncountersQuery = `
	with k as (
	  select m.encounter_id, count(distinct (m.report_id, m.fight_index))::int as kills,
	         mode() within group (order by m.size) as size
	  from fight_metrics m
	  where m.encounter_id is not null and m.kill and m.state <> 'removed'
	  group by m.encounter_id
	), n as (` + encounterQuery + `)
	select k.encounter_id, n.name, z.log_zone, k.size, k.kills
	from k join n on n.encounter_id = k.encounter_id
	left join lateral (
	  select r.log_zone from fights f join reports r on r.id = f.report_id
	  where f.encounter_id = k.encounter_id and r.log_zone is not null
	  order by f.start_ms desc limit 1
	) z on true
	order by k.kills desc, n.name`

// RankedEncounters lists the encounters that have ranked kills, most kills first.
func (s *Store) RankedEncounters(ctx context.Context) ([]RankedEncounter, error) {
	rows, err := s.Pool.Query(ctx, rankedEncountersQuery)
	if err != nil {
		return nil, fmt.Errorf("rankings: ranked encounters: %w", err)
	}
	defer rows.Close()
	out := []RankedEncounter{}
	for rows.Next() {
		var e RankedEncounter
		if err := rows.Scan(&e.ID, &e.Name, &e.Zone, &e.Size, &e.Kills); err != nil {
			return nil, fmt.Errorf("rankings: ranked encounters: %w", err)
		}
		e.Slug = character.Slug(e.Name)
		out = append(out, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rankings: ranked encounters: %w", err)
	}
	return out, nil
}

// rankedEncounters serves the apps' encounter picker (F11.2).
func (s *Service) rankedEncounters(w http.ResponseWriter, r *http.Request) {
	rows, err := s.Store.RankedEncounters(r.Context())
	if err != nil {
		s.fail(w, r, "ranked encounters", err, "could not read the encounters just now")
		return
	}
	CacheLive(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"encounters": rows})
}
