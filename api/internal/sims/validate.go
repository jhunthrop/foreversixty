package sims

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// ValidateJobCommand is the argument the image dispatches on for the
// nightly validation run.
const ValidateJobCommand = "sim-validate"

// TopParses is how many of a spec's best parses the nightly job sims,
// and ValidateIterations how many iterations each one runs. Section 6
// of the design fixes both.
const (
	TopParses          = 50
	ValidateIterations = defaultIterations
)

// ValidateTimeout bounds one spec's whole pass. Fifty parses at three
// thousand iterations is about two minutes natively, and the design
// budgets the whole nightly job at twenty.
const ValidateTimeout = 10 * time.Minute

// Parse is one ranked fight the validation job sims: the recorded
// character, what they actually did, and how long the fight ran.
type Parse struct {
	ReportID    string
	FightIndex  int
	PlayerKey   string
	Spec        string
	Class       string
	ActualDPS   float64
	DurationSec int
	Combatant   summary.CombatantRow
	// ActualCasts is how many times the player actually landed each
	// ability, from the fight's stored summary, keyed by the summary's
	// row identity. It is what the sim's own cast counts are compared
	// against, and the key is an id rather than a name because the two
	// sides name the same ability differently.
	ActualCasts map[int64]CastCount
}

// CastCount is one ability's count on one side of the comparison, with
// the name to show for it.
type CastCount struct {
	Name  string
	Casts int64
}

// Parses reads the top n parses for one spec in one phase.
// ParseReader is the real implementation.
type Parses interface {
	TopParses(ctx context.Context, spec, phase string, n int) ([]Parse, error)
}

// ValidateDeps is everything `api sim-validate` needs.
type ValidateDeps struct {
	Store  *Store
	Top    Parses
	Engine runner.Runner
	Build  Builder
	Scores Scores
	Log    *slog.Logger
}

func (d ValidateDeps) logger() *slog.Logger {
	return loggerOr(d.Log)
}

// Validate sims the top parses of every spec in specs and writes each
// one's fidelity. One spec failing does not stop the others: the job
// is a nightly measurement, and a spec that cannot be measured keeps
// yesterday's figure rather than taking the whole run down. The job
// still exits non-zero, naming the specs that failed, so the
// schedule's own failure is visible.
func Validate(ctx context.Context, d ValidateDeps, specs []string, phase, engineVersion string) error {
	var failed []string
	for _, spec := range specs {
		if err := validateSpec(ctx, d, spec, phase, engineVersion); err != nil {
			d.logger().Error("sims", "op", "validate", "spec", spec, "err", err)
			failed = append(failed, spec)
		}
	}
	if len(failed) > 0 {
		return fmt.Errorf("sims: validate: %d spec(s) failed: %s", len(failed), strings.Join(failed, ", "))
	}
	return nil
}

// validateSpec measures one spec and writes its row.
func validateSpec(ctx context.Context, d ValidateDeps, spec, phase, engineVersion string) error {
	specCtx, cancel := context.WithTimeout(ctx, ValidateTimeout)
	defer cancel()

	parses, err := d.Top.TopParses(specCtx, spec, phase, TopParses)
	if err != nil {
		return err
	}
	var (
		gaps    []float64
		casts   = map[int64]*WorstAction{}
		skipped int
	)
	for _, p := range parses {
		character, err := d.Build.FightCharacter(spec, p.Class, p.Combatant)
		if errors.Is(err, ErrNoCharacter) {
			// This one parse cannot be measured; the reason is logged
			// and counted, and the rest of the spec's parses still
			// run. The card below is written from whatever could be
			// used, even if that is none of them.
			skipped++
			d.logger().Warn("sims", "op", "validate", "spec", spec,
				"report", p.ReportID, "fight", p.FightIndex, "err", err)
			continue
		}
		if err != nil {
			return err
		}
		res, err := d.Engine.Run(specCtx, simapi.SimRequest{
			EngineVersion: engineVersion, Spec: spec, Iterations: ValidateIterations,
			Source: simapi.CharacterSource{
				Kind: simapi.SourceFight,
				Ref:  fmt.Sprintf("%s:%d", p.ReportID, p.FightIndex),
			},
			Character: character,
			Encounter: withEncounterDefaults(simapi.EncounterSpec{DurationSec: p.DurationSec}),
		}, nil)
		if err != nil {
			return err
		}
		if res.DPS.Mean <= 0 {
			// A sim that produced nothing measures nothing; it is not
			// a hundred per cent gap.
			continue
		}
		gaps = append(gaps, math.Abs(p.ActualDPS-res.DPS.Mean)/res.DPS.Mean)
		foldCasts(casts, res.Summary, p)
		// The nightly pass is also the backstop for the column the
		// fight-close scorer could not fill.
		if err := d.Scores.SetExecutionScore(specCtx, p.ReportID, p.FightIndex,
			p.PlayerKey, p.ActualDPS/res.DPS.Mean); err != nil {
			return err
		}
	}

	if skipped > 0 {
		d.logger().Warn("sims", "op", "validate", "spec", spec,
			"parses", len(gaps), "skipped", skipped)
	}

	f := SpecFidelity{
		Spec: spec, Parses: len(gaps), WorstActions: worstOf(casts, len(gaps)),
		EngineVersion: engineVersion,
	}
	// An unmeasured spec has no gap at all, which is a null column
	// and not a zero. StateFor is given an infinite gap so it can
	// only fall through to unsupported.
	gap := math.Inf(1)
	if len(gaps) > 0 {
		m := median(gaps)
		f.MedianGap, gap = &m, m
	}
	f.State = StateFor(gap, f.Parses)
	return d.Store.PutSpec(specCtx, f)
}

// foldCasts adds one parse's per-ability cast counts, simulated and
// actual, into the running totals, keyed by the summary's row
// identity. worstOf divides both sides by the parse count, so the
// figures the page shows are per-fight averages rather than sums over
// fifty fights.
//
// The sim's own name for a row is the engine's form ("spell:23881"),
// so the parse's name wins wherever there is one: the card should read
// Bloodthirst.
func foldCasts(into map[int64]*WorstAction, sim summary.Summary, p Parse) {
	at := func(id int64, name string) *WorstAction {
		row := into[id]
		if row == nil {
			row = &WorstAction{SpellID: id, Name: name}
			into[id] = row
		}
		return row
	}
	for _, c := range sim.Casts {
		at(c.SpellID, c.SpellName).SimCasts += float64(c.Succeeded)
	}
	for id, actual := range p.ActualCasts {
		row := at(id, actual.Name)
		row.ActualCasts += float64(actual.Casts)
		if actual.Name != "" {
			row.Name = actual.Name
		}
	}
}

// worstOf is the WorstActionsPerSpec abilities whose per-fight cast
// counts the sim and the parses disagree on most, largest gap first.
func worstOf(casts map[int64]*WorstAction, parses int) []WorstAction {
	out := make([]WorstAction, 0, len(casts))
	for _, a := range casts {
		row := *a
		if parses > 0 {
			row.SimCasts /= float64(parses)
			row.ActualCasts /= float64(parses)
		}
		out = append(out, row)
	}
	slices.SortFunc(out, func(a, b WorstAction) int {
		da, db := math.Abs(a.SimCasts-a.ActualCasts), math.Abs(b.SimCasts-b.ActualCasts)
		switch {
		case da > db:
			return -1
		case da < db:
			return 1
		case a.SpellID < b.SpellID:
			// A stable order, so two equal gaps do not swap between
			// nights and make the support page look like it moved.
			// The id breaks the tie, because two rows can share a name.
			return -1
		case a.SpellID > b.SpellID:
			return 1
		default:
			return 0
		}
	})
	if len(out) > WorstActionsPerSpec {
		out = out[:WorstActionsPerSpec]
	}
	return out
}

// median is the middle value of a sample, averaging the two middle
// values of an even one. It sorts a copy: the caller's slice is not
// this function's to reorder.
func median(values []float64) float64 {
	sorted := slices.Clone(values)
	slices.Sort(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2]
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2
}
