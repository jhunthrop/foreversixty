package bulk

import (
	"encoding/json"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestPlanIsTheFirstRungOfTheLadder(t *testing.T) {
	cases := []struct {
		precision  string
		iterations int
	}{
		{api.PrecisionFast, 100},
		{api.PrecisionNormal, 1000},
		{api.PrecisionHigh, 1000},
	}
	for _, c := range cases {
		t.Run(c.precision, func(t *testing.T) {
			req := withBulk(api.KindGear, candidate("head", itemHelm))
			req.Bulk.Precision = c.precision
			final, _ := api.FinalIterations(c.precision)
			req.Iterations = final

			stage, err := Plan(req)
			if err != nil {
				t.Fatal(err)
			}
			if stage.Stage != 1 {
				t.Errorf("stage = %d, want 1", stage.Stage)
			}
			if stage.Iterations != c.iterations {
				t.Errorf("iterations = %d, want %d", stage.Iterations, c.iterations)
			}
			if len(stage.Requests) != len(stage.Combos)+1 {
				t.Fatalf("%d requests and %d combos; the equipped set is the extra one",
					len(stage.Requests), len(stage.Combos))
			}
			for i, r := range stage.Requests {
				if r.Iterations != c.iterations {
					t.Errorf("request %d asks for %d iterations", i, r.Iterations)
				}
				if r.Bulk != nil {
					t.Errorf("request %d still carries a bulk block", i)
				}
				if err := r.ValidatePart(); err != nil {
					t.Errorf("request %d is not runnable: %v", i, err)
				}
			}
		})
	}
}

// Requests[0] is always the equipped set: it is the baseline every
// delta is measured against, and the page and the binary both index it
// by position rather than searching for it.
//
// itemHelm2 - a second, real, unrestricted head item in this build -
// stands in for the brief's itemHelm+1: itemHelm has no +1 neighbour
// in this build (it resolves to Invulnerable Mail, a chest item), and
// a chest candidate under a "head" slot request would be skipped for
// a slot mismatch, expanding to nothing and making Plan fail before
// this test ever gets to assert what it means to.
func TestTheEquippedSetIsRequestZero(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm2))
	stage, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	equipped := stage.Requests[0]
	if len(equipped.Character.Gear) != len(req.Character.Gear) {
		t.Fatalf("request 0 has %d items, the character has %d", len(equipped.Character.Gear), len(req.Character.Gear))
	}
	for i, g := range req.Character.Gear {
		if equipped.Character.Gear[i] != g {
			t.Errorf("request 0 slot %d is %+v, the character wears %+v", i, equipped.Character.Gear[i], g)
		}
	}
	if equipped.Character.Talents != req.Character.Talents {
		t.Error("request 0 does not carry the character's own talents")
	}
}

// Every request in a stage shares one seed, so the combinations and
// the equipped set see the same rolls and the delta between them is a
// paired comparison rather than two independent samples. Stages differ,
// so a combination that got a lucky stream at 100 iterations does not
// keep it at 1,000.
func TestSeedsArePairedWithinAStageAndDifferBetweenThem(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm), candidate("finger2", itemRing+1))
	first, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	seed := first.Requests[0].RandomSeed
	for i, r := range first.Requests {
		if r.RandomSeed != seed {
			t.Errorf("request %d has seed %d, request 0 has %d; a stage is paired", i, r.RandomSeed, seed)
		}
	}
	second := stageRequests(req, 2, 1000, first.Combos)
	if second.Requests[0].RandomSeed == seed {
		t.Error("stage 2 reuses stage 1's seed")
	}
	// Deterministic: planning twice produces the same seeds, so a
	// re-run of a saved request reproduces the run.
	again, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	if again.Requests[0].RandomSeed != seed {
		t.Error("two plans of one request disagree about the seed")
	}
}

// The stage object crosses the wasm boundary as JSON, so its names are
// the contract's and are pinned here.
func TestStageRequestsJSONNames(t *testing.T) {
	stage, err := Plan(withBulk(api.KindGear, candidate("head", itemHelm)))
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(stage)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"stage", "iterations", "requests", "combos"} {
		if _, ok := got[key]; !ok {
			t.Errorf("a stage does not carry %q", key)
		}
	}
	var back StageRequests
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Stage != stage.Stage || len(back.Requests) != len(stage.Requests) || len(back.Combos) != len(stage.Combos) {
		t.Errorf("round trip lost something: %+v", back)
	}
}

func TestPlanRefusesAPlainRun(t *testing.T) {
	if _, err := Plan(base()); err == nil {
		t.Error("Plan accepted a request with no bulk block")
	}
}

// Every candidate is for a class the character is not, so nothing
// survives eligibility. A plan of zero combinations is a ranking of
// nothing, and saying so beats running the equipped set alone and
// calling it a Top Gear.
//
// itemClassLocked, offered at "main_hand" where it actually fits,
// stands in for the brief's itemHelm at "head": itemHelm carries no
// class restriction in this build, so a mage would still be offered
// it and the expansion would not be empty. itemClassLocked is the
// fixture already kept for exactly this case (see
// TestCandidatesTheCharacterCannotEquipAreSkipped's "another class's
// item").
func TestPlanRefusesAnExpansionWithNothingInIt(t *testing.T) {
	req := withBulk(api.KindGear, candidate("main_hand", itemClassLocked))
	req.Character.Class = "mage"
	req.Spec = "mage-frost"
	if _, err := Plan(req); err == nil {
		t.Error("Plan accepted an expansion with no combinations")
	}
}

// Task 19's contract (10.3): the sim module sets NoSample on every
// stage request and never on a plain run's own shape, because a stage
// sim's cast log is never read. The equipped baseline is the one
// request in a stage that most resembles an ordinary run and is the
// easiest to forget - so it is checked on its own, not just folded
// into the loop below.
func TestEveryStageRequestSetsNoSample(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm), candidate("finger2", itemRing+1))
	stage, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	if !stage.Requests[0].NoSample {
		t.Error("the equipped baseline does not set NoSample")
	}
	for i, r := range stage.Requests {
		if !r.NoSample {
			t.Errorf("request %d does not set NoSample", i)
		}
	}
}
