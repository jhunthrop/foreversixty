package sims

import (
	"context"
	"io"
	"log/slog"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

// fakeParses answers with parses a test set up, without a bucket.
type fakeParses struct {
	bySpec map[string][]Parse
	err    error
}

func (f fakeParses) TopParses(_ context.Context, spec, _ string, n int) ([]Parse, error) {
	if f.err != nil {
		return nil, f.err
	}
	rows := f.bySpec[spec]
	if len(rows) > n {
		rows = rows[:n]
	}
	return rows, nil
}

// parsesAt builds parses whose actual DPS is each of the given
// values, all on the same spec.
func parsesAt(spec string, values ...float64) []Parse {
	out := make([]Parse, 0, len(values))
	for i, v := range values {
		out = append(out, Parse{
			ReportID: "rep" + spec, FightIndex: i + 1, PlayerKey: "us/normal/baelgrim",
			Spec: spec, Class: "warrior", ActualDPS: v, DurationSec: 180,
			Combatant: summary.CombatantRow{GUID: "Player-1", Name: "Baelgrim"},
			// Keyed by the row identity, with the log's own name for
			// the card. 23881 is Bloodthirst, which the fixture's sim
			// casts too, so the two sides join.
			ActualCasts: map[int64]CastCount{23881: {Name: "Bloodthirst", Casts: 41}},
		})
	}
	return out
}

func (h *harness) validateDeps(top Parses, engine runner.Runner, build Builder, scores Scores) ValidateDeps {
	return ValidateDeps{
		Store: h.store, Top: top, Engine: engine, Build: build, Scores: scores,
		Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

// measured pulls one spec's card out of the support-page read.
func (h *harness) measured(t *testing.T, spec string) SpecFidelity {
	t.Helper()
	all, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range all {
		if f.Spec == spec {
			return f
		}
	}
	t.Fatalf("%s has no card at all", spec)
	return SpecFidelity{}
}

func TestMedianIsTheMiddleAndDoesNotReorderItsInput(t *testing.T) {
	odd := []float64{0.3, 0.1, 0.2}
	if got := median(odd); got != 0.2 {
		t.Fatalf("median(odd) = %v, want 0.2", got)
	}
	if odd[0] != 0.3 {
		t.Fatalf("median sorted the caller's slice: %v", odd)
	}
	if got := median([]float64{0.4, 0.1, 0.2, 0.3}); math.Abs(got-0.25) > 1e-9 {
		t.Fatalf("median(even) = %v, want 0.25", got)
	}
}

func TestAWellSimmedSpecBecomesValidated(t *testing.T) {
	h := newHarness(t)
	// Fifty parses all at 1020 against a sim at 1000: a 2% gap.
	values := make([]float64, TopParses)
	for i := range values {
		values[i] = 1020
	}
	top := fakeParses{bySpec: map[string][]Parse{"warrior-fury": parsesAt("warrior-fury", values...)}}
	scores := &fakeScores{}
	err := Validate(t.Context(),
		h.validateDeps(top, &runner.Fixture{Mean: 1000}, &alwaysBuilds{}, scores),
		[]string{"warrior-fury"}, "raids-1", testEngine)
	if err != nil {
		t.Fatal(err)
	}

	f := h.measured(t, "warrior-fury")
	if f.State != SpecValidated {
		t.Fatalf("state %q, want %q", f.State, SpecValidated)
	}
	if f.Parses != TopParses {
		t.Errorf("parses %d, want %d", f.Parses, TopParses)
	}
	if f.MedianGap == nil || math.Abs(*f.MedianGap-0.02) > 1e-9 {
		t.Errorf("median gap %v, want 0.02", f.MedianGap)
	}
	if f.EngineVersion != testEngine {
		t.Errorf("engine version %q", f.EngineVersion)
	}
	if len(f.WorstActions) == 0 || f.WorstActions[0].SpellID != 23881 {
		t.Fatalf("worst actions %+v", f.WorstActions)
	}
	if f.WorstActions[0].Name != "Bloodthirst" {
		t.Errorf("the card shows %q, want the log's name", f.WorstActions[0].Name)
	}
	// Per-fight averages, not sums over fifty fights. The fixture
	// casts Bloodthirst 52 times a fight; the parses cast it 41.
	w := f.WorstActions[0]
	if math.Abs(w.SimCasts-52) > 1e-9 || math.Abs(w.ActualCasts-41) > 1e-9 {
		t.Errorf("cast counts %+v, want the per-fight averages 52 and 41", w)
	}
	// And the column the fight-close scorer could not fill is filled.
	if len(scores.calls) != TopParses {
		t.Errorf("%d execution scores written, want %d", len(scores.calls), TopParses)
	}
}

func TestATooWideSpecIsInProgressNotValidated(t *testing.T) {
	h := newHarness(t)
	values := make([]float64, TopParses)
	for i := range values {
		values[i] = 800 // a 20% gap against a sim at 1000
	}
	top := fakeParses{bySpec: map[string][]Parse{"mage-frost": parsesAt("mage-frost", values...)}}
	if err := Validate(t.Context(),
		h.validateDeps(top, &runner.Fixture{Mean: 1000}, &alwaysBuilds{}, &fakeScores{}),
		[]string{"mage-frost"}, "raids-1", testEngine); err != nil {
		t.Fatal(err)
	}
	if got := h.measured(t, "mage-frost").State; got != SpecInProgress {
		t.Fatalf("state %q, want %q", got, SpecInProgress)
	}
}

func TestASpecWithNoParsesStaysUnsupportedWithANullGap(t *testing.T) {
	h := newHarness(t)
	if err := Validate(t.Context(),
		h.validateDeps(fakeParses{}, &runner.Fixture{Mean: 1000}, &alwaysBuilds{}, &fakeScores{}),
		[]string{"rogue-combat"}, "raids-1", testEngine); err != nil {
		t.Fatal(err)
	}
	f := h.measured(t, "rogue-combat")
	if f.State != SpecUnsupported || f.MedianGap != nil {
		t.Fatalf("card %+v, want unsupported with a null gap", f)
	}
}

func TestWithNoBuilderYesterdaysRowStands(t *testing.T) {
	h := newHarness(t)
	gap := 0.02
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecValidated, MedianGap: &gap, Parses: 60,
		WorstActions: []WorstAction{}, EngineVersion: testEngine}); err != nil {
		t.Fatal(err)
	}
	top := fakeParses{bySpec: map[string][]Parse{"warrior-fury": parsesAt("warrior-fury", 1020)}}
	if err := Validate(t.Context(),
		h.validateDeps(top, &runner.Fixture{Mean: 1000}, NoBuilder{}, &fakeScores{}),
		[]string{"warrior-fury"}, "raids-1", testEngine); err != nil {
		t.Fatalf("a missing character source is not a failure: %v", err)
	}
	f := h.measured(t, "warrior-fury")
	if f.State != SpecValidated || f.Parses != 60 {
		t.Fatalf("yesterday's row was overwritten: %+v", f)
	}
}

func TestOneFailingSpecDoesNotStopTheOthers(t *testing.T) {
	h := newHarness(t)
	// A reader that fails for everything: both specs fail, and the
	// error names both.
	err := Validate(t.Context(),
		h.validateDeps(fakeParses{err: errAnyway}, &runner.Fixture{}, &alwaysBuilds{}, &fakeScores{}),
		[]string{"warrior-fury", "mage-frost"}, "raids-1", testEngine)
	if err == nil {
		t.Fatal("expected the job to exit non-zero")
	}

	// A reader that answers for one spec and not the other: both rows
	// are still written, one measured and one unsupported.
	good := fakeParses{bySpec: map[string][]Parse{"warrior-fury": parsesAt("warrior-fury", 1020)}}
	if err := Validate(t.Context(),
		h.validateDeps(good, &runner.Fixture{Mean: 1000}, &alwaysBuilds{}, &fakeScores{}),
		[]string{"warrior-fury", "mage-frost"}, "raids-1", testEngine); err != nil {
		t.Fatal(err)
	}
	if got := h.measured(t, "warrior-fury").Parses; got != 1 {
		t.Fatalf("warrior-fury parses %d, want 1", got)
	}
	if got := h.measured(t, "mage-frost").State; got != SpecUnsupported {
		t.Fatalf("mage-frost state %q", got)
	}
}

func TestWorstActionsAreTheBiggestGapsFirstAndCappedAtFive(t *testing.T) {
	casts := map[int64]*WorstAction{
		1680:  {SpellID: 1680, Name: "Whirlwind", SimCasts: 10, ActualCasts: 9},
		23881: {SpellID: 23881, Name: "Bloodthirst", SimCasts: 52, ActualCasts: 20},
		20647: {SpellID: 20647, Name: "Execute", SimCasts: 8, ActualCasts: 2},
		11567: {SpellID: 11567, Name: "Heroic Strike", SimCasts: 30, ActualCasts: 28},
		12292: {SpellID: 12292, Name: "Death Wish", SimCasts: 2, ActualCasts: 1},
		2687:  {SpellID: 2687, Name: "Bloodrage", SimCasts: 2, ActualCasts: 2},
	}
	got := worstOf(casts, 1)
	if len(got) != WorstActionsPerSpec {
		t.Fatalf("%d rows, want %d", len(got), WorstActionsPerSpec)
	}
	if got[0].Name != "Bloodthirst" || got[1].Name != "Execute" {
		t.Fatalf("order: %q then %q", got[0].Name, got[1].Name)
	}
	// Two equal gaps keep a stable order, so the page does not look
	// like it moved overnight. Ties break on the id, which is stable
	// even when two rows share a name.
	if got := worstOf(map[int64]*WorstAction{
		9: {SpellID: 9, Name: "Bravo", SimCasts: 5, ActualCasts: 4},
		2: {SpellID: 2, Name: "Alpha", SimCasts: 5, ActualCasts: 4},
	}, 1); got[0].SpellID != 2 {
		t.Fatalf("ties are not ordered by row identity: %+v", got)
	}
}

func TestOnlyTheSameRowIdentityIsCompared(t *testing.T) {
	// The sim's row names are the engine's own form; the parse's are
	// the log's. Joining on the name would produce two half-rows.
	casts := map[int64]*WorstAction{}
	foldCasts(casts, summary.Summary{Casts: []summary.CastRow{
		{SpellID: 23881, SpellName: "spell:23881", Succeeded: 52},
		{SpellID: 2_000_000_123, SpellName: "other:melee", Succeeded: 148},
	}}, Parse{ActualCasts: map[int64]CastCount{23881: {Name: "Bloodthirst", Casts: 41}}})
	if len(casts) != 2 {
		t.Fatalf("%d rows, want the joined spell and the sim-only one", len(casts))
	}
	joined := casts[23881]
	if joined.SimCasts != 52 || joined.ActualCasts != 41 {
		t.Fatalf("the join did not happen: %+v", joined)
	}
	if joined.Name != "Bloodthirst" {
		t.Errorf("name %q, want the log's", joined.Name)
	}
	if only := casts[2_000_000_123]; only.ActualCasts != 0 || only.Name != "other:melee" {
		t.Errorf("a sim-only row should keep the engine's name: %+v", only)
	}
}

func TestTheJobMeasuresEveryDPSSpecAndOnlyThose(t *testing.T) {
	h := newHarness(t)
	if err := Validate(t.Context(),
		h.validateDeps(fakeParses{}, &runner.Fixture{}, &alwaysBuilds{}, &fakeScores{}),
		DPSSpecs(), "raids-1", testEngine); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.store.Pool.QueryRow(t.Context(), `select count(*) from sim_specs`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != len(DPSSpecs()) {
		t.Fatalf("%d rows written, want one per dps spec (%d)", n, len(DPSSpecs()))
	}
}
