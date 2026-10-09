package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/request"
	simscore "github.com/jhunthrop/foreversixty/sim/score"
)

func gearItem(id int, slot string) scored {
	return scored{candidate: candidate{ID: id, Name: fmt.Sprintf("Item %d", id), Slots: []string{slot}}}
}

func pickOf(item scored) slotPick { return slotPick{Item: &item} }

// incumbentFixture is a new set (head 1, chest 2, legs 3) and a pool that
// also offers head 11 and chest 12, the incumbent's items.
func incumbentFixture() (verifiedBand, map[string][]scored) {
	current := verifiedBand{picks: map[string]slotPick{
		"head":  pickOf(gearItem(1, "head")),
		"chest": pickOf(gearItem(2, "chest")),
		"legs":  pickOf(gearItem(3, "legs")),
	}, setDPS: 500}
	pools := map[string][]scored{
		"head":  {gearItem(1, "head"), gearItem(11, "head")},
		"chest": {gearItem(2, "chest"), gearItem(12, "chest")},
		"legs":  {gearItem(3, "legs")},
	}
	return current, pools
}

var incumbentGear = map[string]int{"head": 11, "chest": 12, "legs": 3}

// engineFor scores the new set at newScore and the incumbent set at
// incumbentScore, both with the same error.
func engineFor(newScore, incumbentScore, stdErr float64) *fakeEngine {
	return &fakeEngine{
		DPSByGear:     map[string]float64{gear(1, 2, 3): newScore, gear(11, 12, 3): incumbentScore},
		DefaultStdErr: stdErr,
	}
}

func TestKeepIncumbentKeepsTheOldSetThatBeatsTheNewOneBeyondError(t *testing.T) {
	current, pools := incumbentFixture()
	engine := engineFor(500, 520, 2)
	got, note, err := keepIncumbent(engine, specInfo{Spec: "druid-restoration"}, "tauren", "druid", 60, "", current, incumbentGear, pools)
	if err != nil {
		t.Fatal(err)
	}
	if note == nil {
		t.Fatal("the incumbent beat the new set by 20 against a combined error of 2.8 and was not kept")
	}
	if got.setDPS != 520 || got.picks["head"].Item.ID != 11 || got.picks["chest"].Item.ID != 12 || got.picks["legs"].Item.ID != 3 {
		t.Errorf("kept band = %+v, want the incumbent's items at 520", got)
	}
	if got.picks["head"].RunnerUp == nil || got.picks["head"].RunnerUp.ID != 1 {
		t.Error("the new pick must be the kept slot's runner-up")
	}
	if got.picks["head"].Item.MeasuredDPS != 520 || got.picks["legs"].Item.MeasuredDPS != 0 {
		t.Error("a changed slot is sim-decided at the incumbent's score; an unchanged one is not")
	}
	if want := []string{"chest", "head"}; !reflect.DeepEqual(note.DifferingSlots, want) {
		t.Errorf("differing slots = %v, want %v", note.DifferingSlots, want)
	}
	if note.IncumbentScore != 520 || note.NewScore != 500 || note.IncumbentError != 2 || note.NewError != 2 {
		t.Errorf("note = %+v, want both scores and errors", note)
	}
}

func TestKeepIncumbentKeepsTheNewSetWhenTheDifferenceIsInsideError(t *testing.T) {
	current, pools := incumbentFixture()
	// 3 ahead against a combined error of 2.83 would be kept; 2 is not.
	_, note, err := keepIncumbent(engineFor(500, 502, 2), specInfo{}, "tauren", "druid", 60, "", current, incumbentGear, pools)
	if err != nil || note != nil {
		t.Fatalf("note = %+v, err = %v: a gap inside the combined error must keep the new set", note, err)
	}
	_, note, err = keepIncumbent(engineFor(500, 490, 0.1), specInfo{}, "tauren", "druid", 60, "", current, incumbentGear, pools)
	if err != nil || note != nil {
		t.Fatalf("note = %+v, err = %v: a worse incumbent must never be kept", note, err)
	}
}

func TestKeepIncumbentCostsNothingWhenThereIsNothingToCompare(t *testing.T) {
	current, pools := incumbentFixture()
	cases := map[string]map[string]int{
		"the same set":                             {"head": 1, "chest": 2, "legs": 3},
		"no incumbent":                             nil,
		"an item the pool lacks":                   {"head": 99, "chest": 2, "legs": 3},
		"a slot the pool lacks":                    {"head": 1, "chest": 2, "legs": 3, "wrist": 5},
		"an item of another slot":                  {"head": 12, "chest": 2, "legs": 3},
		"a set with the new set's items elsewhere": {"head": 1, "chest": 2, "legs": 11},
	}
	for name, gear := range cases {
		engine := engineFor(500, 600, 0)
		got, note, err := keepIncumbent(engine, specInfo{}, "tauren", "druid", 60, "", current, gear, pools)
		if err != nil || note != nil || got.setDPS != 500 {
			t.Errorf("%s: got %+v note %+v err %v, want the new band back", name, got, note, err)
		}
		if len(engine.Calls) != 0 {
			t.Errorf("%s: %d sim runs, want none", name, len(engine.Calls))
		}
	}
}

func TestKeepIncumbentDropsASlotTheIncumbentLeftEmpty(t *testing.T) {
	current, pools := incumbentFixture()
	engine := &fakeEngine{DPSByGear: map[string]float64{gear(1, 2, 3): 500, gear(1, 2): 530}}
	got, note, err := keepIncumbent(engine, specInfo{}, "tauren", "druid", 60, "", current, map[string]int{"head": 1, "chest": 2}, pools)
	if err != nil || note == nil {
		t.Fatalf("note = %+v, err = %v, want the incumbent without legs kept", note, err)
	}
	if got.picks["legs"].Item != nil || !reflect.DeepEqual(note.DifferingSlots, []string{"legs"}) {
		t.Errorf("picks = %+v, differing = %v, want legs empty", got.picks, note.DifferingSlots)
	}
}

func TestLoadIncumbentSetsReadsTheCommittedReport(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "druid-restoration.json")
	body := `{"bands":[{"band":60,"preset":"raid","faction":"horde","slots":[{"slot":"head","item_id":11},{"slot":"neck"}]}]}`
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := loadIncumbentSets(path)
	if err != nil {
		t.Fatal(err)
	}
	want := incumbentSets{{Band: 60, Preset: "raid", Faction: "horde"}: {"head": 11}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("incumbents = %v, want %v", got, want)
	}
	if none, err := loadIncumbentSets(filepath.Join(dir, "absent.json")); err != nil || len(none) != 0 {
		t.Errorf("a missing report = %v, %v, want no incumbents and no error", none, err)
	}
	if err := os.WriteFile(path, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := loadIncumbentSets(path); err == nil {
		t.Error("a corrupt report must fail loudly, not read as no incumbent")
	}
}

// noisyBackend answers like the engine: the healing's error shrinks with
// the square root of the iterations asked for.
type noisyBackend struct{ fakeHealing }

func (b *noisyBackend) Run(req api.SimRequest, _ request.HealProfile) (simscore.HealingResult, error) {
	b.requests = append(b.requests, req)
	return simscore.HealingResult{
		Effective:    api.Estimate{Mean: 600, Error: 14 / math.Sqrt(float64(req.Iterations))},
		Raw:          api.Estimate{Mean: 620},
		ManaLastsSec: 280,
	}, nil
}

// The score's error at the healer verification harness is well below the
// adoption margin, and the iterations asked for are scaled.
func TestHealerVerificationErrorIsBelowTheAdoptionMargin(t *testing.T) {
	backend := &noisyBackend{}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}
	score, stdErr, err := engine.RunPlainDPSWithError(api.SimRequest{Iterations: verifyIterations})
	if err != nil {
		t.Fatal(err)
	}
	if relative := stdErr / score; relative >= swapMargin/2 {
		t.Errorf("error at the verification harness is %.2f%% of the score, want under half the %.0f%% margin", 100*relative, 100*swapMargin)
	}
	if got := backend.requests[0].Iterations; got != verifyIterations*healIterationScale {
		t.Errorf("ran %d iterations, want %d", got, verifyIterations*healIterationScale)
	}
}
