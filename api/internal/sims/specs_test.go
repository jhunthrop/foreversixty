package sims

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

func TestStateForIsTheDesignsDefinitionOfValidated(t *testing.T) {
	for _, c := range []struct {
		name   string
		gap    float64
		parses int
		want   string
	}{
		{"nothing measured", 0, 0, SpecUnsupported},
		{"close but too few parses", 0.01, 49, SpecInProgress},
		{"enough parses but too wide", 0.06, 200, SpecInProgress},
		{"exactly the parse floor and inside the gap", 0.049, 50, SpecValidated},
		{"comfortably validated", 0.02, 120, SpecValidated},
		{"exactly the gap is not inside it", 0.05, 120, SpecInProgress},
	} {
		if got := StateFor(c.gap, c.parses); got != c.want {
			t.Errorf("%s: StateFor(%v, %d) = %q, want %q", c.name, c.gap, c.parses, got, c.want)
		}
	}
}

func TestEverySpecHasACardBeforeAnythingIsMeasured(t *testing.T) {
	h := newHarness(t)
	var out struct {
		Specs []SpecFidelity `json:"specs"`
	}
	res := h.do(http.MethodGet, "/v1/specs", "", nil)
	if cc := res.Header.Get("Cache-Control"); cc == "" {
		t.Error("the support page should be cacheable")
	}
	want := "public, max-age=300"
	if cc := res.Header.Get("Cache-Control"); cc != want {
		t.Errorf("Cache-Control = %q, want %q", cc, want)
	}
	h.data(res, &out)
	if len(out.Specs) != len(DPSSpecs()) {
		t.Fatalf("%d cards, want one per dps spec (%d)", len(out.Specs), len(DPSSpecs()))
	}
	for _, f := range out.Specs {
		if f.State != SpecUnsupported {
			t.Errorf("%s is %q before anything measured it", f.Spec, f.State)
		}
		if f.MedianGap != nil {
			t.Errorf("%s has a gap before anything measured it", f.Spec)
		}
		if f.UpdatedAt != nil {
			t.Errorf("%s was updated at %v before anything measured it", f.Spec, *f.UpdatedAt)
		}
		if f.WorstActions == nil {
			t.Errorf("%s: worst_actions must be an empty list, never null", f.Spec)
		}
	}
	if !slices.IsSortedFunc(out.Specs, func(a, b SpecFidelity) int {
		if a.Spec < b.Spec {
			return -1
		} else if a.Spec > b.Spec {
			return 1
		}
		return 0
	}) {
		t.Error("the cards are not in a stable order")
	}
}

// TestAnUnmeasuredCardCarriesNoDate is the wire-level half of the
// check above: the year-one date must never reach the page, so the
// field is null in the JSON rather than a zero time.
func TestAnUnmeasuredCardCarriesNoDate(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/v1/specs", "", nil)
	defer res.Body.Close()
	var env struct {
		Data struct {
			Specs []map[string]any `json:"specs"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if len(env.Data.Specs) == 0 {
		t.Fatal("no cards")
	}
	for _, card := range env.Data.Specs {
		at, ok := card["updated_at"]
		if !ok {
			t.Fatalf("%v: the field must be present so the web can read its absence", card["spec"])
		}
		if at != nil {
			t.Errorf("%v: updated_at = %v, want null", card["spec"], at)
		}
	}
}

func TestAMeasuredSpecIsLaidOverItsCard(t *testing.T) {
	h := newHarness(t)
	gap := 0.021
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecValidated, MedianGap: &gap, Parses: 50,
		WorstActions:  []WorstAction{{SpellID: 23881, Name: "Bloodthirst", SimCasts: 52, ActualCasts: 41}},
		EngineVersion: testEngine}); err != nil {
		t.Fatal(err)
	}
	var out struct {
		Specs []SpecFidelity `json:"specs"`
	}
	h.data(h.do(http.MethodGet, "/v1/specs", "", nil), &out)
	if len(out.Specs) != len(DPSSpecs()) {
		t.Fatalf("%d cards, want %d", len(out.Specs), len(DPSSpecs()))
	}
	var fury SpecFidelity
	for _, f := range out.Specs {
		if f.Spec == "warrior-fury" {
			fury = f
		}
	}
	if fury.State != SpecValidated || fury.Parses != 50 {
		t.Fatalf("fury: %+v", fury)
	}
	if fury.MedianGap == nil || *fury.MedianGap != gap {
		t.Fatalf("fury gap: %v", fury.MedianGap)
	}
	if len(fury.WorstActions) != 1 || fury.WorstActions[0].SpellID != 23881 ||
		fury.WorstActions[0].Name != "Bloodthirst" {
		t.Fatalf("worst actions: %+v", fury.WorstActions)
	}
	if fury.EngineVersion != testEngine {
		t.Errorf("engine version %q", fury.EngineVersion)
	}
	// A measured card does have a date, which is what makes the null
	// on an unmeasured one meaningful.
	if fury.UpdatedAt == nil || fury.UpdatedAt.IsZero() {
		t.Errorf("a measured card has no updated_at: %v", fury.UpdatedAt)
	}
}

func TestPutSpecOverwritesTheNightBefore(t *testing.T) {
	h := newHarness(t)
	first := 0.09
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecInProgress, MedianGap: &first, Parses: 50,
		WorstActions: []WorstAction{}, EngineVersion: "an older build"}); err != nil {
		t.Fatal(err)
	}
	second := 0.03
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecValidated, MedianGap: &second, Parses: 50,
		WorstActions: []WorstAction{}, EngineVersion: testEngine}); err != nil {
		t.Fatal(err)
	}
	specs, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	var n int
	for _, f := range specs {
		if f.Spec == "warrior-fury" {
			n++
			if f.State != SpecValidated || f.EngineVersion != testEngine {
				t.Fatalf("after the second night: %+v", f)
			}
		}
	}
	if n != 1 {
		t.Fatalf("%d warrior-fury cards, want exactly one", n)
	}
}

func TestOnlyAValidatedSpecIsScored(t *testing.T) {
	h := newHarness(t)
	if ok, err := h.store.Validated(t.Context(), "warrior-fury"); err != nil || ok {
		t.Fatalf("an unmeasured spec is not validated: ok=%v err=%v", ok, err)
	}
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecInProgress, Parses: 3,
		WorstActions: []WorstAction{}}); err != nil {
		t.Fatal(err)
	}
	if ok, err := h.store.Validated(t.Context(), "warrior-fury"); err != nil || ok {
		t.Fatalf("an in-progress spec is not validated: ok=%v err=%v", ok, err)
	}
	gap := 0.02
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecValidated, MedianGap: &gap, Parses: 60,
		WorstActions: []WorstAction{}}); err != nil {
		t.Fatal(err)
	}
	if ok, err := h.store.Validated(t.Context(), "warrior-fury"); err != nil || !ok {
		t.Fatalf("a validated spec is validated: ok=%v err=%v", ok, err)
	}
}

// TestPutSpecBoundsWorstActionsToTheConstant pins the fix for the
// review's HIGH finding: WorstActionsPerSpec is declared to bound the
// list, so PutSpec has to be the one place that enforces it, rather
// than trusting every future caller to trim its own slice first.
func TestPutSpecBoundsWorstActionsToTheConstant(t *testing.T) {
	h := newHarness(t)
	given := make([]WorstAction, WorstActionsPerSpec+3)
	for i := range given {
		given[i] = WorstAction{SpellID: int64(i + 1), Name: fmt.Sprintf("Ability %d", i+1)}
	}
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", WorstActions: given}); err != nil {
		t.Fatal(err)
	}
	specs, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	var fury SpecFidelity
	for _, f := range specs {
		if f.Spec == "warrior-fury" {
			fury = f
		}
	}
	if len(fury.WorstActions) != WorstActionsPerSpec {
		t.Fatalf("%d worst actions, want the bound (%d)", len(fury.WorstActions), WorstActionsPerSpec)
	}
	for i, a := range fury.WorstActions {
		if a.SpellID != given[i].SpellID {
			t.Fatalf("worst action %d: spell %d, want %d (caller's order not kept)",
				i, a.SpellID, given[i].SpellID)
		}
	}
}

// TestPutSpecComputesStateRatherThanTrustingTheCaller pins the fix
// for the review's MEDIUM finding: StateFor is the design's only
// definition of "validated", so PutSpec has to compute State itself
// from MedianGap and Parses rather than writing whatever the caller
// handed it, which could disagree with its own figures.
func TestPutSpecComputesStateRatherThanTrustingTheCaller(t *testing.T) {
	h := newHarness(t)

	// The caller claims unsupported, but the figures clearly qualify
	// as validated: the computed state wins.
	gap := 0.02
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", State: SpecUnsupported, MedianGap: &gap, Parses: 60,
		WorstActions: []WorstAction{}}); err != nil {
		t.Fatal(err)
	}
	if ok, err := h.store.Validated(t.Context(), "warrior-fury"); err != nil || !ok {
		t.Fatalf("the figures said validated regardless of the claimed state: ok=%v err=%v", ok, err)
	}

	// The caller claims validated, but there is no measured gap at
	// all: an unmeasured gap can never read back validated.
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "mage-fire", State: SpecValidated, Parses: 60,
		WorstActions: []WorstAction{}}); err != nil {
		t.Fatal(err)
	}
	specs, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	var fire SpecFidelity
	for _, f := range specs {
		if f.Spec == "mage-fire" {
			fire = f
		}
	}
	if fire.State != SpecInProgress {
		t.Fatalf("mage-fire: %+v, want in_progress (an unmeasured gap cannot be validated)", fire)
	}
}

func TestEverySpecCardNamesItsReferenceStat(t *testing.T) {
	h := newHarness(t)
	// A measured row and an unmeasured one both carry it: the weights
	// page reads the reference off the card before anything has been
	// simmed.
	gap := 0.02
	if err := h.store.PutSpec(t.Context(), SpecFidelity{
		Spec: "warrior-fury", Parses: 50, MedianGap: &gap, EngineVersion: testEngine,
	}); err != nil {
		t.Fatal(err)
	}
	cards, err := h.store.Specs(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) == 0 {
		t.Fatal("no cards at all")
	}
	for _, c := range cards {
		want := specs.ByKey[c.Spec].ReferenceStat
		if want == "" {
			t.Fatalf("%s has no reference_stat in the generated spec list", c.Spec)
		}
		if c.ReferenceStat != want {
			t.Errorf("%s: reference_stat %q, want %q", c.Spec, c.ReferenceStat, want)
		}
	}
}
