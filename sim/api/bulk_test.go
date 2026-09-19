package api

import (
	"errors"
	"strings"
	"testing"
)

// gear is a minimal valid Top Gear request. Every case below starts
// from it and breaks one thing, so a failure names the rule it broke.
func gear() SimRequest {
	req := runReq()
	req.Iterations = 3000
	req.Bulk = &BulkSpec{
		Mode:       KindGear,
		Precision:  PrecisionNormal,
		Cap:        Caps[LaneBrowser],
		Candidates: []Candidate{{Slot: "head", ItemID: 16963, Origin: OriginBag}},
	}
	return req
}

func TestKindIsDerivedFromTheRequest(t *testing.T) {
	cases := []struct {
		name string
		req  SimRequest
		want string
	}{
		{"a plain run", runReq(), KindRun},
		{"top gear", gear(), KindGear},
		{"droptimizer", func() SimRequest {
			r := gear()
			r.Bulk.Mode = KindDrops
			return r
		}(), KindDrops},
		{"talent compare", func() SimRequest {
			r := gear()
			r.Bulk.Mode = KindTalents
			return r
		}(), KindTalents},
		// The weights kind is Task 4's; its case lives in
		// TestWeightsKindAndShape, beside the type that makes it
		// possible.
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.req.Kind(); got != c.want {
				t.Errorf("Kind() = %q, want %q", got, c.want)
			}
		})
	}
}

func TestBulkValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string // a substring of the error; "" means the request is legal
	}{
		{"the happy path", func(*SimRequest) {}, ""},
		{"an unknown mode", func(r *SimRequest) { r.Bulk.Mode = "everything" }, "bulk.mode"},
		{"an unknown precision", func(r *SimRequest) { r.Bulk.Precision = "exact" }, "bulk.precision"},
		{"a cap over the largest lane's", func(r *SimRequest) { r.Bulk.Cap = Caps[LaneServer] + 1 }, "bulk.cap"},
		{"a cap of zero", func(r *SimRequest) { r.Bulk.Cap = 0 }, "bulk.cap"},
		{"a candidate on a locked slot", func(r *SimRequest) { r.Bulk.Locked = []string{"head"} }, "locked"},
		{"an unknown locked slot", func(r *SimRequest) { r.Bulk.Locked = []string{"tabard"} }, "bulk.locked"},
		{"an origin outside the vocabulary", func(r *SimRequest) { r.Bulk.Candidates[0].Origin = "wishful" }, "origin"},
		{"a drop origin with no source", func(r *SimRequest) { r.Bulk.Candidates[0].Origin = "drop:" }, "origin"},
		{"a set origin with a name", func(r *SimRequest) { r.Bulk.Candidates[0].Origin = "set:my AQ set" }, ""},
		{"a candidate with no item", func(r *SimRequest) { r.Bulk.Candidates[0].ItemID = 0 }, "item_id"},
		{"an unknown candidate slot", func(r *SimRequest) { r.Bulk.Candidates[0].Slot = "tabard" }, "bulk.candidates"},
		{"gear with nothing to try", func(r *SimRequest) { r.Bulk.Candidates = nil }, "at least one"},
		{"talents mode with a loadout", func(r *SimRequest) {
			r.Bulk.Mode = KindTalents
			r.Bulk.Candidates = nil
			r.Bulk.Talents = []TalentLoadout{{Name: "Deep Fury", Talents: "30305001302-05050005525010051"}}
		}, ""},
		{"talents mode with no loadout", func(r *SimRequest) {
			r.Bulk.Mode = KindTalents
			r.Bulk.Candidates = nil
		}, "at least one talent loadout"},
		{"talents mode carrying candidates", func(r *SimRequest) {
			r.Bulk.Mode = KindTalents
			r.Bulk.Talents = []TalentLoadout{{Name: "Deep Fury", Talents: "30305001302-05050005525010051"}}
		}, "no candidates"},
		{"a loadout with no name", func(r *SimRequest) {
			r.Bulk.Talents = []TalentLoadout{{Talents: "30305001302-05050005525010051"}}
		}, "talent loadout"},
		{"drops mode from a boss", func(r *SimRequest) {
			r.Bulk.Mode = KindDrops
			r.Bulk.Candidates[0].Origin = "drop:raid:mc:lucifron"
		}, ""},
		{"drops mode from a bag", func(r *SimRequest) { r.Bulk.Mode = KindDrops }, "every candidate"},
		{"a set with no gear", func(r *SimRequest) { r.Bulk.Sets = []GearSet{{Name: "PvP"}} }, "gear set"},
		{"alternative consumable lists", func(r *SimRequest) {
			r.Bulk.Consumables = [][]string{{"flask_of_the_titans"}, {"elixir_of_the_mongoose", "juju_power"}, {}}
		}, ""},
		{"a consumable list with an empty id", func(r *SimRequest) {
			r.Bulk.Consumables = [][]string{{""}}
		}, "empty consumable id"},
		{"consumable lists in drops mode", func(r *SimRequest) {
			r.Bulk.Mode = KindDrops
			r.Bulk.Candidates[0].Origin = "drop:raid:mc:lucifron"
			r.Bulk.Consumables = [][]string{{"flask_of_the_titans"}}
		}, "gear-mode dimension"},
		{"a gear set in drops mode", func(r *SimRequest) {
			r.Bulk.Mode = KindDrops
			r.Bulk.Candidates[0].Origin = "drop:raid:mc:lucifron"
			r.Bulk.Sets = []GearSet{{Name: "my AQ set", Gear: []GearSlot{{Slot: "head", ItemID: 16963}}}}
		}, "gear-mode dimension"},
		{"a candidate naming its source", func(r *SimRequest) {
			r.Bulk.Mode = KindDrops
			r.Bulk.Candidates[0].Origin = "drop:raid:mc:lucifron"
			r.Bulk.Candidates[0].SourceName = "Lucifron"
		}, ""},
		{"iterations that are not the precision's final count", func(r *SimRequest) { r.Iterations = 500 }, "iterations must be 3000"},
		{"high precision at ten thousand", func(r *SimRequest) {
			r.Bulk.Precision = PrecisionHigh
			r.Iterations = 10000
		}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := gear()
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

// The ladders are the fork's fast_mode made explicit, and both lanes
// read them, so their shape is pinned rather than trusted: one cut per
// gap between stages, and a final count that is a number the settings
// bar offers.
func TestTheLaddersAreWellFormed(t *testing.T) {
	want := map[string][]int{
		PrecisionFast:   {100, 1000, 3000},
		PrecisionNormal: {1000, 3000},
		PrecisionHigh:   {1000, 10000},
	}
	for _, p := range Precisions {
		l, ok := Ladders[p]
		if !ok {
			t.Fatalf("no ladder for precision %q", p)
		}
		if len(l.Cuts) != len(l.Iterations)-1 {
			t.Errorf("%s: %d stages and %d cuts; there is one cut per gap", p, len(l.Iterations), len(l.Cuts))
		}
		if got := l.Iterations; !equalInts(got, want[p]) {
			t.Errorf("%s ladder is %v, the contract says %v", p, got, want[p])
		}
		final, ok := FinalIterations(p)
		if !ok || final != l.Iterations[len(l.Iterations)-1] {
			t.Errorf("FinalIterations(%q) = %d, %v", p, final, ok)
		}
	}
	if _, ok := FinalIterations("exact"); ok {
		t.Error("FinalIterations accepted a precision that is not in the vocabulary")
	}
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCapExceededCarriesBothNumbers(t *testing.T) {
	err := error(ErrCapExceeded{Cap: 400, Combinations: 1280})
	var capped ErrCapExceeded
	if !errors.As(err, &capped) {
		t.Fatal("ErrCapExceeded does not match errors.As")
	}
	if capped.Cap != 400 || capped.Combinations != 1280 {
		t.Errorf("got %+v", capped)
	}
	for _, want := range []string{"400", "1280"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("%q does not name %s", err, want)
		}
	}
}
