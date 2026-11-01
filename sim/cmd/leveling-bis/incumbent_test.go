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

// incumbentConsidered is the incumbent as the one set keepVerifyBest
// considers, or none when the pools cannot reproduce it.
func incumbentConsidered(gear map[string]int, current verifiedBand, pools map[string][]scored) []consideredSet {
	picks, ok := incumbentPicks(gear, current.picks, pools)
	if !ok {
		return nil
	}
	return []consideredSet{{Label: consideredIncumbent, Picks: picks}}
}

func TestKeepVerifyBestKeepsTheIncumbentThatBeatsTheNewSetBeyondError(t *testing.T) {
	current, pools := incumbentFixture()
	engine := engineFor(500, 520, 2)
	got, note, err := keepVerifyBest(engine, specInfo{Spec: "druid-restoration"}, "tauren", "druid", 60, "", current, incumbentConsidered(incumbentGear, current, pools))
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
		t.Error("a changed slot is sim-decided at the kept set's score; an unchanged one is not")
	}
	if want := []string{"chest", "head"}; !reflect.DeepEqual(note.DifferingSlots, want) {
		t.Errorf("differing slots = %v, want %v", note.DifferingSlots, want)
	}
	if note.Kept != consideredIncumbent || note.KeptScore != 520 || note.NewScore != 500 || note.KeptError != 2 || note.NewError != 2 {
		t.Errorf("note = %+v, want the kept set's label, both scores and errors", note)
	}
}

func TestKeepVerifyBestKeepsTheNewSetWhenTheDifferenceIsInsideError(t *testing.T) {
	current, pools := incumbentFixture()
	// 3 ahead against a combined error of 2.83 would be kept; 2 is not.
	_, note, err := keepVerifyBest(engineFor(500, 502, 2), specInfo{}, "tauren", "druid", 60, "", current, incumbentConsidered(incumbentGear, current, pools))
	if err != nil || note != nil {
		t.Fatalf("note = %+v, err = %v: a gap inside the combined error must keep the new set", note, err)
	}
	_, note, err = keepVerifyBest(engineFor(500, 490, 0.1), specInfo{}, "tauren", "druid", 60, "", current, incumbentConsidered(incumbentGear, current, pools))
	if err != nil || note != nil {
		t.Fatalf("note = %+v, err = %v: a worse incumbent must never be kept", note, err)
	}
}

func TestKeepVerifyBestCostsNothingWhenThereIsNothingToCompare(t *testing.T) {
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
		got, note, err := keepVerifyBest(engine, specInfo{}, "tauren", "druid", 60, "", current, incumbentConsidered(gear, current, pools))
		if err != nil || note != nil || got.setDPS != 500 {
			t.Errorf("%s: got %+v note %+v err %v, want the new band back", name, got, note, err)
		}
		if len(engine.Calls) != 0 {
			t.Errorf("%s: %d sim runs, want none", name, len(engine.Calls))
		}
	}
}

func TestKeepVerifyBestDropsASlotTheIncumbentLeftEmpty(t *testing.T) {
	current, pools := incumbentFixture()
	engine := &fakeEngine{DPSByGear: map[string]float64{gear(1, 2, 3): 500, gear(1, 2): 530}}
	got, note, err := keepVerifyBest(engine, specInfo{}, "tauren", "druid", 60, "", current, incumbentConsidered(map[string]int{"head": 1, "chest": 2}, current, pools))
	if err != nil || note == nil {
		t.Fatalf("note = %+v, err = %v, want the incumbent without legs kept", note, err)
	}
	if got.picks["legs"].Item != nil || !reflect.DeepEqual(note.DifferingSlots, []string{"legs"}) {
		t.Errorf("picks = %+v, differing = %v, want legs empty", got.picks, note.DifferingSlots)
	}
}

// Owner ruling 4 (round 5): the ranker never publishes a set its own check
// scores lower than an alternative it considered. Every set a band passed
// through on its way to the final one is such an alternative, and the best
// of them that beats the final set beyond error is what gets published.
func TestKeepVerifyBestPublishesTheBestConsideredSet(t *testing.T) {
	current, _ := incumbentFixture()
	beforeSwaps := map[string]slotPick{
		"head":  pickOf(gearItem(11, "head")),
		"chest": pickOf(gearItem(2, "chest")),
		"legs":  pickOf(gearItem(3, "legs")),
	}
	beforeSets := map[string]slotPick{
		"head":  pickOf(gearItem(1, "head")),
		"chest": pickOf(gearItem(12, "chest")),
		"legs":  pickOf(gearItem(3, "legs")),
	}
	engine := &fakeEngine{DPSByGear: map[string]float64{
		gear(1, 2, 3):  500,
		gear(11, 2, 3): 530,
		gear(1, 12, 3): 515,
	}, DefaultStdErr: 2}
	got, note, err := keepVerifyBest(engine, specInfo{Spec: "mage-arcane"}, "troll", "mage", 60, "", current, []consideredSet{
		{Label: consideredBeforeVerify, Picks: beforeSwaps},
		{Label: consideredBeforeSets, Picks: beforeSets},
	})
	if err != nil {
		t.Fatal(err)
	}
	if note == nil || note.Kept != consideredBeforeVerify {
		t.Fatalf("note = %+v, want the set before the swaps (530) kept over the final 500", note)
	}
	if got.setDPS != 530 || got.picks["head"].Item.ID != 11 || got.picks["chest"].Item.ID != 2 {
		t.Errorf("kept band = %+v, want head 11 and chest 2 at 530", got.picks)
	}
	if got.picks["head"].RunnerUp == nil || got.picks["head"].RunnerUp.ID != 1 {
		t.Error("the final set's item must be the kept slot's runner-up")
	}
}

// An alternative is measured once however many stages left it unchanged,
// and the final set once.
func TestKeepVerifyBestMeasuresEachDistinctSetOnce(t *testing.T) {
	current, _ := incumbentFixture()
	other := map[string]slotPick{
		"head":  pickOf(gearItem(11, "head")),
		"chest": pickOf(gearItem(2, "chest")),
		"legs":  pickOf(gearItem(3, "legs")),
	}
	engine := &fakeEngine{DPSByGear: map[string]float64{gear(1, 2, 3): 500, gear(11, 2, 3): 490}}
	_, note, err := keepVerifyBest(engine, specInfo{}, "troll", "mage", 60, "", current, []consideredSet{
		{Label: consideredBeforeVerify, Picks: other},
		{Label: consideredBeforeSets, Picks: other},
		{Label: consideredBeforeConfirmed, Picks: current.picks},
	})
	if err != nil || note != nil {
		t.Fatalf("note = %+v, err = %v, want the final set kept", note, err)
	}
	if len(engine.Calls) != 2 {
		t.Errorf("%d runs (%v), want two: the final set and the one distinct alternative", len(engine.Calls), engine.Calls)
	}
}

// The confirmed-stats check holds an item the client has not confirmed out
// of the set on purpose; a considered set that wears one again is not an
// alternative the guard may restore.
func TestKeepVerifyBestNeverRestoresAnUnconfirmedItem(t *testing.T) {
	current, _ := incumbentFixture()
	unconfirmed := gearItem(11, "head")
	unconfirmed.ClientUnconfirmed = true
	withUnconfirmed := map[string]slotPick{
		"head":  pickOf(unconfirmed),
		"chest": pickOf(gearItem(2, "chest")),
		"legs":  pickOf(gearItem(3, "legs")),
	}
	engine := &fakeEngine{DPSByGear: map[string]float64{gear(1, 2, 3): 500, gear(11, 2, 3): 600}}
	_, note, err := keepVerifyBest(engine, specInfo{}, "troll", "mage", 60, "", current, []consideredSet{{Label: consideredBeforeConfirmed, Picks: withUnconfirmed}})
	if err != nil || note != nil {
		t.Fatalf("note = %+v, err = %v, want the unconfirmed item left out", note, err)
	}
	if len(engine.Calls) != 0 {
		t.Errorf("%d runs, want none", len(engine.Calls))
	}
}

// The guard compares sets at the healer harness's effective iterations for
// every role: a healer's runner scales verifyIterations itself, every other
// role asks for the scaled count directly, so a 1 to 2% gap is resolved
// beyond error where the 300-iteration verify run cannot resolve it.
func TestGuardIterationsMatchTheHealerHarnessForEveryRole(t *testing.T) {
	if got := guardIterations(specInfo{Role: healerRole}); got != verifyIterations {
		t.Errorf("healer guard iterations = %d, want %d (the heal runner scales them)", got, verifyIterations)
	}
	if got := guardIterations(specInfo{Role: "dps"}); got != verifyIterations*healIterationScale {
		t.Errorf("dps guard iterations = %d, want %d", got, verifyIterations*healIterationScale)
	}
	current, _ := incumbentFixture()
	other := map[string]slotPick{"head": pickOf(gearItem(11, "head")), "chest": pickOf(gearItem(2, "chest")), "legs": pickOf(gearItem(3, "legs"))}
	var seen []int
	engine := &fakeEngine{DPSFunc: func(req api.SimRequest) (float64, error) {
		seen = append(seen, req.Iterations)
		return 500, nil
	}}
	if _, _, err := keepVerifyBest(engine, specInfo{Role: "tank"}, "orc", "warrior", 60, "", current, []consideredSet{{Label: consideredBeforeVerify, Picks: other}}); err != nil {
		t.Fatal(err)
	}
	for _, n := range seen {
		if n != verifyIterations*healIterationScale {
			t.Errorf("a tank guard run asked for %d iterations, want %d", n, verifyIterations*healIterationScale)
		}
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
