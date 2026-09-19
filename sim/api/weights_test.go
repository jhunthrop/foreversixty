package api

import (
	"strings"
	"testing"
)

func weightsReq() SimRequest {
	req := runReq()
	req.Weights = &WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit", "hit"},
		Reference: "attack_power",
	}
	return req
}

func TestWeightsValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string
	}{
		{"the happy path", func(*SimRequest) {}, ""},
		{"no stats", func(r *SimRequest) { r.Weights.Stats = nil }, "at least one stat"},
		{"no reference", func(r *SimRequest) { r.Weights.Reference = "" }, "weights.reference"},
		{"a reference that is not weighed", func(r *SimRequest) { r.Weights.Reference = "strength" }, "one of the stats it weighs"},
		{"a repeated stat", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", "crit"}
			r.Weights.Reference = "crit"
		}, "listed twice"},
		{"an empty stat id", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", ""}
			r.Weights.Reference = "crit"
		}, "empty stat id"},
		{"weights and bulk together", func(r *SimRequest) {
			r.Bulk = &BulkSpec{Mode: KindGear, Precision: PrecisionNormal, Cap: 10,
				Candidates: []Candidate{{Slot: "head", ItemID: 1, Origin: OriginBag}}}
		}, "a request is one kind"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := weightsReq()
			c.edit(&req)
			err := req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// The stat ids themselves are the engine's, so the envelope does not
// carry a copy of the list; sim/request resolves them and IDS.md
// publishes them. What the envelope owns is the shape.
func TestWeightsKindAndShape(t *testing.T) {
	req := weightsReq()
	if req.Kind() != KindWeights {
		t.Errorf("Kind() = %q, want %q", req.Kind(), KindWeights)
	}
	w := StatWeight{Stat: "crit", Weight: 1.0, Error: 0.04}
	if w.Stat != "crit" || w.Weight != 1.0 || w.Error != 0.04 {
		t.Errorf("StatWeight round trip: %+v", w)
	}
}
