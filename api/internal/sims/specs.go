package sims

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// The states a spec can be in on the support page.
const (
	SpecValidated   = "validated"
	SpecInProgress  = "in_progress"
	SpecUnsupported = "unsupported"
)

// ValidatedGap is the median DPS gap under which a spec counts as
// validated, and ValidatedParses the number of parses it has to be
// measured over. Section 6 of the design fixes both.
const (
	ValidatedGap    = 0.05
	ValidatedParses = 50
)

// WorstActionsPerSpec is how many disagreeing abilities the support
// page lists per spec.
const WorstActionsPerSpec = 5

// specsMaxAge is how long the support page may be cached, in seconds:
// the figures change once a night.
const specsMaxAge = 300

// WorstAction is one ability whose cast count the sim and the parses
// disagree on. Both figures are per-fight averages.
//
// SpellID is the summary's row identity, which is what the two sides
// are matched on: a sim's CastRow.SpellName is the engine's own
// "spell:23881" form, a log's is "Bloodthirst", so a name join would
// match nothing. Name is for the card, and is the log's name when a
// parse cast the ability and the sim's form when only the sim did.
type WorstAction struct {
	SpellID     int64   `json:"spell_id"`
	Name        string  `json:"name"`
	SimCasts    float64 `json:"sim_casts"`
	ActualCasts float64 `json:"actual_casts"`
}

// SpecFidelity is one card on the support page.
type SpecFidelity struct {
	Spec          string        `json:"spec"`
	State         string        `json:"state"`
	MedianGap     *float64      `json:"median_gap"`
	Parses        int           `json:"parses"`
	WorstActions  []WorstAction `json:"worst_actions"`
	EngineVersion string        `json:"engine_version"`
	// UpdatedAt is when the nightly job last measured this spec, and
	// nil for a card nothing has measured yet. It is a pointer for the
	// same reason MedianGap is: a zero time.Time marshals as
	// "0001-01-01T00:00:00Z", which reads as a date rather than as an
	// absence, and the support page would render it as one.
	UpdatedAt *time.Time `json:"updated_at"`
}

// StateFor is the state a measurement puts a spec in: validated when
// the median gap is under five per cent over at least fifty parses,
// in progress once anything at all has been measured, unsupported
// otherwise. This function is the design's sentence; nothing else
// decides what "validated" means.
func StateFor(medianGap float64, parses int) string {
	switch {
	case parses >= ValidatedParses && medianGap < ValidatedGap:
		return SpecValidated
	case parses > 0:
		return SpecInProgress
	default:
		return SpecUnsupported
	}
}

// Specs is one card per spec the simulator models, with whatever has
// been measured laid over it. The card list is the data lane's
// generated one, so a spec nobody has simmed yet still has a card
// that says so - which is what design section 4.4 asks for.
func (s *Store) Specs(ctx context.Context) ([]SpecFidelity, error) {
	byspec := map[string]SpecFidelity{}
	for _, spec := range DPSSpecs() {
		// No MedianGap and no UpdatedAt: nothing has measured this
		// spec, and both fields marshal as null to say so.
		byspec[spec] = SpecFidelity{
			Spec: spec, State: SpecUnsupported, WorstActions: []WorstAction{},
		}
	}
	rows, err := s.Pool.Query(ctx,
		`select spec, state, median_gap, parses, worst_actions,
		        coalesce(engine_version, ''), updated_at
		 from sim_specs`)
	if err != nil {
		return nil, fmt.Errorf("sims: read specs: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			f     SpecFidelity
			worst []byte
		)
		if err := rows.Scan(&f.Spec, &f.State, &f.MedianGap, &f.Parses, &worst,
			&f.EngineVersion, &f.UpdatedAt); err != nil {
			return nil, fmt.Errorf("sims: scan spec: %w", err)
		}
		f.WorstActions = []WorstAction{}
		if err := json.Unmarshal(worst, &f.WorstActions); err != nil {
			return nil, fmt.Errorf("sims: decode worst actions for %s: %w", f.Spec, err)
		}
		// A measured row for a spec that is no longer in the list is
		// still shown: a rename should be visible, not silent.
		byspec[f.Spec] = f
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	out := make([]SpecFidelity, 0, len(byspec))
	for _, f := range byspec {
		out = append(out, f)
	}
	slices.SortFunc(out, func(a, b SpecFidelity) int {
		switch {
		case a.Spec < b.Spec:
			return -1
		case a.Spec > b.Spec:
			return 1
		default:
			return 0
		}
	})
	return out, nil
}

// PutSpec writes one spec's measurement. The nightly validation job
// is the only caller.
func (s *Store) PutSpec(ctx context.Context, f SpecFidelity) error {
	if f.WorstActions == nil {
		f.WorstActions = []WorstAction{}
	}
	worst, err := json.Marshal(f.WorstActions)
	if err != nil {
		return fmt.Errorf("sims: encode worst actions for %s: %w", f.Spec, err)
	}
	_, err = s.Pool.Exec(ctx,
		`insert into sim_specs (spec, state, median_gap, parses, worst_actions,
		   engine_version, updated_at)
		 values ($1, $2, $3, $4, $5, $6, now())
		 on conflict (spec) do update set state = excluded.state,
		   median_gap = excluded.median_gap, parses = excluded.parses,
		   worst_actions = excluded.worst_actions,
		   engine_version = excluded.engine_version, updated_at = now()`,
		f.Spec, f.State, f.MedianGap, f.Parses, worst, nullIfEmpty(f.EngineVersion))
	if err != nil {
		return fmt.Errorf("sims: write spec %s: %w", f.Spec, err)
	}
	return nil
}

// Validated reports whether one spec's execution scores may be
// computed and shown. The fight-close scorer asks before it queues,
// and an unmeasured spec is simply not validated.
func (s *Store) Validated(ctx context.Context, spec string) (bool, error) {
	var state string
	err := s.Pool.QueryRow(ctx, `select state from sim_specs where spec = $1`, spec).Scan(&state)
	if isNoRows(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("sims: read spec %s: %w", spec, err)
	}
	return state == SpecValidated, nil
}

// nullIfEmpty keeps an unmeasured spec's engine_version null rather
// than an empty string, which is what the column's nullability means.
func nullIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (s *Service) specs(w http.ResponseWriter, r *http.Request) {
	out, err := s.Store.Specs(r.Context())
	if err != nil {
		s.fail(w, r, "specs", err, "could not read the spec list just now")
		return
	}
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(specsMaxAge))
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"specs": out})
}
