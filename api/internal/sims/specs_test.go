package sims

import (
	"encoding/json"
	"net/http"
	"slices"
	"testing"
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
