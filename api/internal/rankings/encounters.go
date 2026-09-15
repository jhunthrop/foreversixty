package rankings

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/digest"
)

// ErrNoEncounter is returned when a slug names no encounter this
// deployment has ever seen.
var ErrNoEncounter = fmt.Errorf("rankings: no such encounter")

// Encounter is one encounter as the rankings know it: the id every row
// carries, and the name the fights recorded for it.
//
// There is no encounter table in the schema and the design says there
// will not be one ("Encounters name themselves from ENCOUNTER_START"),
// so the names come from the fights themselves - the most recent name
// seen for each id, which is what a rename would leave behind.
type Encounter struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// encounterQuery is the one place an encounter's display name comes
// from. Both the slug resolver and the name lookup read it.
const encounterQuery = `select distinct on (encounter_id) encounter_id, name
	from fights where encounter_id is not null and name <> ''
	order by encounter_id, start_ms desc`

// Encounters lists every encounter any report has recorded.
func (s *Store) Encounters(ctx context.Context) ([]Encounter, error) {
	rows, err := s.Pool.Query(ctx, encounterQuery)
	if err != nil {
		return nil, fmt.Errorf("rankings: list encounters: %w", err)
	}
	defer rows.Close()
	out := []Encounter{}
	for rows.Next() {
		var e Encounter
		if err := rows.Scan(&e.ID, &e.Name); err != nil {
			return nil, fmt.Errorf("rankings: list encounters: %w", err)
		}
		e.Slug = character.Slug(e.Name)
		out = append(out, e)
	}
	return out, rows.Err()
}

// ResolveEncounter reads the `encounter` parameter, which the contract
// lets a caller write either way: the numeric id the rows carry, or the
// slug the site's own /rankings/<encounter-slug> URL uses. A number is
// taken at face value and costs no query; anything else is matched
// against the encounter names.
func (s *Store) ResolveEncounter(ctx context.Context, v string) (int64, error) {
	if id, err := strconv.ParseInt(v, 10, 64); err == nil {
		if id <= 0 {
			return 0, ErrNoEncounter
		}
		return id, nil
	}
	slug := character.Slug(v)
	if slug == "" {
		return 0, ErrNoEncounter
	}
	all, err := s.Encounters(ctx)
	if err != nil {
		return 0, err
	}
	for _, e := range all {
		if e.Slug == slug {
			return e.ID, nil
		}
	}
	return 0, ErrNoEncounter
}

// EncounterNames maps encounter ids to their display names. Ids nothing
// has recorded a name for are simply absent.
func (s *Store) EncounterNames(ctx context.Context, ids []int64) (map[int64]string, error) {
	out := map[int64]string{}
	if len(ids) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx,
		`select distinct on (encounter_id) encounter_id, name
		 from fights where encounter_id = any($1) and name <> ''
		 order by encounter_id, start_ms desc`, ids)
	if err != nil {
		return nil, fmt.Errorf("rankings: read encounter names: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("rankings: read encounter names: %w", err)
		}
		out[id] = name
	}
	return out, rows.Err()
}

// bracketKey identifies one percentile bracket by every axis a digest
// is keyed on. This is distinct from store.go's own digestKey (a
// {metric, spec} pair scoped to one fight's fixed encounter and
// difficulty, used only to order digest row locks during a write):
// a read spans many encounters, difficulties and phases at once, so it
// needs all five axes to address a bracket.
type bracketKey struct {
	EncounterID int64
	Difficulty  int64
	Spec        string
	Phase       string
	Metric      string
}

// digestsFor reads every digest belonging to a set of encounters in one
// query, so a character page places a hundred fights without a hundred
// round trips.
func (s *Store) digestsFor(ctx context.Context, ids []int64) (map[bracketKey]*digest.Digest, error) {
	out := map[bracketKey]*digest.Digest{}
	if len(ids) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx,
		`select encounter_id, difficulty, spec, phase, metric, digest
		 from percentile_digests where encounter_id = any($1)`, ids)
	if err != nil {
		return nil, fmt.Errorf("rankings: read digests: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			k   bracketKey
			raw []byte
		)
		if err := rows.Scan(&k.EncounterID, &k.Difficulty, &k.Spec, &k.Phase, &k.Metric, &raw); err != nil {
			return nil, fmt.Errorf("rankings: read digests: %w", err)
		}
		d, err := digest.Unmarshal(raw)
		if err != nil {
			return nil, err
		}
		out[k] = d
	}
	return out, rows.Err()
}

// percentileOf places one value in its bracket, or returns nil when the
// bracket has no digest yet.
func percentileOf(digests map[bracketKey]*digest.Digest, k bracketKey, value float64) *float64 {
	d, ok := digests[k]
	if !ok || d.Count() == 0 {
		return nil
	}
	p := d.CDF(value) * 100
	return &p
}
