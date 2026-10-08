package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

const testSetID = 41 // a real entry in setids_generated.go

func setItem(id int, name string, setID int, slots ...string) scored {
	sid := setID
	return scored{candidate: candidate{ID: id, Name: name, SetID: &sid, Slots: slots}}
}

func withScore(item scored, score float64) scored {
	item.Score = score
	return item
}

func intPtr(n int) *int { return &n }

func testCatalog() setCatalog {
	return setCatalog{testSetID: {Name: "Test Regalia", Bonuses: []setBonusTier{
		{Pieces: 2, Description: "two piece"},
		{Pieces: 3, Description: "three piece"},
	}}}
}

func gear(pairs ...int) string {
	var g []api.GearSlot
	slots := []string{"head", "chest", "legs"}
	for i, id := range pairs {
		g = append(g, api.GearSlot{Slot: slots[i], ItemID: id})
	}
	return gearKey(g)
}

func TestBestSetPiecePerSlotSkipsTrinketsAndOtherSets(t *testing.T) {
	bySlot := map[string][]scored{
		"head":     {setItem(1, "Other Set Head", 999999, "head"), setItem(2, "Set Head", testSetID, "head")},
		"trinket1": {setItem(3, "Set Trinket", testSetID, "trinket1")},
	}
	got := bestSetPiecePerSlot(bySlot, "", testSetID)
	if len(got) != 1 || got["head"].ID != 2 {
		t.Fatalf("pieces = %+v, want only the set head (2)", got)
	}
}

// finger1 and finger2 share one candidate list: one physical ring must not
// count as two pieces, and the second finger takes the next distinct ring.
func TestBestSetPiecePerSlotNeverCountsTheSameRingTwice(t *testing.T) {
	ringA := setItem(50, "Set Ring A", testSetID, "finger1", "finger2")
	ringB := setItem(52, "Set Ring B", testSetID, "finger1", "finger2")
	bySlot := map[string][]scored{
		"finger1": {ringA, ringB},
		"finger2": {ringA, ringB},
	}
	got := bestSetPiecePerSlot(bySlot, "", testSetID)
	if got["finger1"].ID != 50 || got["finger2"].ID != 52 {
		t.Fatalf("pieces = %+v, want ring A on finger1 and ring B on finger2", got)
	}
	single := bestSetPiecePerSlot(map[string][]scored{"finger1": {ringA}, "finger2": {ringA}}, "", testSetID)
	if len(single) != 1 {
		t.Fatalf("pieces = %+v, want the lone ring once", single)
	}
}

func TestBestSetPiecePerSlotExcludesTwoHandMainHandForADualWieldSpec(t *testing.T) {
	twoHander := scored{candidate: candidate{ID: 60, Name: "Two-Hand Set Sword", SetID: intPtr(testSetID), Slots: []string{"main_hand"}, TwoHand: true}}
	got := bestSetPiecePerSlot(map[string][]scored{"main_hand": {twoHander}}, "hunter-survival", testSetID)
	if len(got) != 0 {
		t.Fatalf("pieces = %+v, want none (the only main_hand piece is a two-hander)", got)
	}
}

// Two mediocre pieces of a set with a strong 2-piece bonus beat two better
// loose pieces: the per-slot pass picked the loose ones (higher score), the
// set trial is adopted and the changed slots say why.
func TestTrySetCompletionAdoptsAStrongTwoPieceBonus(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Loose Head"}, Score: 9}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Loose Chest"}, Score: 9}},
		"legs":  {Item: &scored{candidate: candidate{ID: 30, Name: "Loose Legs"}, Score: 5}},
	}
	bySlot := map[string][]scored{
		"head":  {withScore(setItem(11, "Set Head", testSetID, "head"), 6)},
		"chest": {withScore(setItem(21, "Set Chest", testSetID, "chest"), 7)},
	}
	fake := &fakeEngine{DPSByGear: map[string]float64{
		gear(10, 20, 30): 100,
		gear(11, 21, 30): 112, // the pair of set pieces wins on the bonus
	}}
	out, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, testCatalog())
	for _, slot := range []string{"head", "chest"} {
		pk := out[slot]
		if pk.Item.ID != map[string]int{"head": 11, "chest": 21}[slot] {
			t.Fatalf("%s = %+v, want the set piece", slot, pk.Item)
		}
		if pk.Item.MeasuredDPS != 112 {
			t.Errorf("%s MeasuredDPS = %v, want 112", slot, pk.Item.MeasuredDPS)
		}
		want := setBonusNote{Set: "Test Regalia", Pieces: 2, Bonus: "two piece"}
		if pk.SetBonus == nil || *pk.SetBonus != want {
			t.Errorf("%s SetBonus = %+v, want %+v", slot, pk.SetBonus, want)
		}
	}
	if out["head"].RunnerUp == nil || out["head"].RunnerUp.ID != 10 {
		t.Errorf("head runner-up = %+v, want the replaced per-slot pick (10), so verifyBand can still swap it back", out["head"].RunnerUp)
	}
	if out["legs"].Item.ID != 30 || out["legs"].SetBonus != nil {
		t.Errorf("legs = %+v, want the untouched per-slot pick without a set note", out["legs"])
	}
	if len(fake.Calls) != 4 || !strings.Contains(strings.Join(notes, "|"), "4 extra sim runs in") {
		t.Errorf("calls = %d, notes = %v, want a screen baseline and trial, then a confirm baseline and trial", len(fake.Calls), notes)
	}
}

func TestTrySetCompletionChangesNothingWhenTheBonusDoesNotPay(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Loose Head"}, Score: 9}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Loose Chest"}, Score: 9}},
	}
	bySlot := map[string][]scored{
		"head":  {withScore(setItem(11, "Set Head", testSetID, "head"), 6)},
		"chest": {withScore(setItem(21, "Set Chest", testSetID, "chest"), 7)},
	}
	fake := &fakeEngine{DPSByGear: map[string]float64{gear(10, 20): 100, gear(11, 21): 98}}
	out, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, testCatalog())
	if out["head"].Item.ID != 10 || out["chest"].Item.ID != 20 || out["head"].SetBonus != nil {
		t.Fatalf("picks changed although the set lost: %+v / %+v", out["head"], out["chest"])
	}
	for _, n := range notes {
		if strings.Contains(n, "adopted") {
			t.Fatalf("notes = %v, want no adoption", notes)
		}
	}
}

func TestTrySetCompletionNeedsMoreThanTheSimError(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Loose Head"}, Score: 9}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Loose Chest"}, Score: 9}},
	}
	bySlot := map[string][]scored{
		"head":  {setItem(11, "Set Head", testSetID, "head")},
		"chest": {setItem(21, "Set Chest", testSetID, "chest")},
	}
	// +2%: clears the margin but sits inside two 2.0 standard errors.
	fake := &fakeEngine{DefaultStdErr: 2, DPSByGear: map[string]float64{gear(10, 20): 100, gear(11, 21): 102}}
	out, _ := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, testCatalog())
	if out["head"].Item.ID != 10 {
		t.Fatalf("head = %+v, want the per-slot pick (the gain is inside the sim error)", out["head"].Item)
	}
}

// A weak 2-piece bonus loses, the 3-piece trial then wins against the
// unchanged picks: the ladder keeps climbing past a losing threshold.
// A trial the short screening run already finds losing never costs the
// full-length run: one screening baseline, one screening trial, nothing else.
func TestTrySetCompletionStopsAtTheScreenWhenTheTrialLoses(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Loose Head"}, Score: 9}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Loose Chest"}, Score: 9}},
	}
	bySlot := map[string][]scored{
		"head":  {setItem(11, "Set Head", testSetID, "head")},
		"chest": {setItem(21, "Set Chest", testSetID, "chest")},
	}
	fake := &fakeEngine{DPSByGear: map[string]float64{gear(10, 20): 100, gear(11, 21): 99}}
	_, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, testCatalog())
	if len(fake.Calls) != 2 {
		t.Fatalf("calls = %d, want 2 (screening baseline and trial only)", len(fake.Calls))
	}
	if !strings.Contains(strings.Join(notes, "|"), "2 extra sim runs in") {
		t.Errorf("notes = %v, want the run count", notes)
	}
}

func TestTrySetCompletionClimbsPastALosingThreshold(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Loose Head"}, Score: 9}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Loose Chest"}, Score: 9}},
		"legs":  {Item: &scored{candidate: candidate{ID: 30, Name: "Loose Legs"}, Score: 9}},
	}
	bySlot := map[string][]scored{
		"head":  {withScore(setItem(11, "Set Head", testSetID, "head"), 8)},
		"chest": {withScore(setItem(21, "Set Chest", testSetID, "chest"), 7)},
		"legs":  {withScore(setItem(31, "Set Legs", testSetID, "legs"), 1)},
	}
	fake := &fakeEngine{DPSByGear: map[string]float64{
		gear(10, 20, 30): 100,
		gear(11, 21, 30): 99,
		gear(11, 21, 31): 120,
	}}
	out, _ := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, testCatalog())
	if out["legs"].Item.ID != 31 || out["legs"].SetBonus == nil || out["legs"].SetBonus.Pieces != 3 {
		t.Fatalf("legs = %+v, want the set legs adopted for the 3-piece bonus", out["legs"])
	}
}

func TestTrySetCompletionSkipsASetAlreadyFullyEquipped(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: setItem(11, "Set Head", testSetID, "head").candidate}},
		"chest": {Item: &scored{candidate: setItem(21, "Set Chest", testSetID, "chest").candidate}},
	}
	bySlot := map[string][]scored{
		"head":  {setItem(11, "Set Head", testSetID, "head")},
		"chest": {setItem(21, "Set Chest", testSetID, "chest")},
	}
	fake := &fakeEngine{}
	_, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, testCatalog())
	if len(notes) != 0 || len(fake.Calls) != 0 {
		t.Fatalf("notes = %v, calls = %d, want none (no threshold above 2 is reachable with two pieces)", notes, len(fake.Calls))
	}
}

func TestTrySetCompletionNoImplementedSetsReturnsUnchanged(t *testing.T) {
	original := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Plain Head"}}}}
	fake := &fakeEngine{}
	out, notes := trySetCompletion(fake, specInfo{}, "dwarf", "hunter", 60, "", original, map[string][]scored{}, testCatalog())
	if len(notes) != 0 || out["head"].Item.ID != 1 || len(fake.Calls) != 0 {
		t.Fatalf("out = %+v, notes = %v, calls = %d, want unchanged and no runs", out, notes, len(fake.Calls))
	}
}

func TestSlotRowSetBonusShape(t *testing.T) {
	with, err := json.Marshal(slotRow{Slot: "head", ItemID: 11, SetBonus: &setBonusNote{Set: "Test Regalia", Pieces: 2, Bonus: "two piece"}})
	if err != nil {
		t.Fatal(err)
	}
	if want := `"set_bonus":{"set":"Test Regalia","pieces":2,"bonus":"two piece"}`; !strings.Contains(string(with), want) {
		t.Errorf("row json = %s, want it to contain %s", with, want)
	}
	without, err := json.Marshal(slotRow{Slot: "head", ItemID: 11})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(without), "set_bonus") {
		t.Errorf("row json = %s, want no set_bonus key", without)
	}
}

func TestLoadSetCatalogSortsBonuses(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, dir+"/sets.json", `[{"id":7,"name":"S","item_ids":[1],"bonuses":[{"pieces":4,"description":"b"},{"pieces":2,"description":"a"}]}]`); err != nil {
		t.Fatal(err)
	}
	cat, err := loadSetCatalog(dir)
	if err != nil {
		t.Fatal(err)
	}
	if got := cat[7]; got.Name != "S" || got.Bonuses[0].Pieces != 2 || got.Bonuses[1].Pieces != 4 {
		t.Fatalf("catalog[7] = %+v", got)
	}
}

func completionFixture() (verifiedBand, map[string][]scored) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 10, Name: "Loose Head"}, Score: 9}},
		"chest": {Item: &scored{candidate: candidate{ID: 20, Name: "Loose Chest"}, Score: 9}},
	}
	bySlot := map[string][]scored{
		"head":  {withScore(setItem(11, "Set Head", testSetID, "head"), 6)},
		"chest": {withScore(setItem(21, "Set Chest", testSetID, "chest"), 7)},
	}
	return verifiedBand{picks: picks, setDPS: 100}, bySlot
}

// The completed set is verified again and published when it beats the
// per-slot set the page would otherwise show.
func TestCompleteSetsKeepsTheCompletedSetWhenItVerifiesHigher(t *testing.T) {
	before, bySlot := completionFixture()
	fake := &fakeEngine{DPSByGear: map[string]float64{gear(10, 20): 100, gear(11, 21): 112}}
	got, _, err := completeSets(fake, specInfo{}, "dwarf", "hunter", 60, "", before, bySlot, testCatalog())
	if err != nil {
		t.Fatal(err)
	}
	if got.setDPS != 112 || got.picks["head"].Item.ID != 11 || got.picks["head"].SetBonus == nil {
		t.Fatalf("band = %+v, want the completed set at 112 with its set note", got)
	}
}

// A completed set whose second verification comes in at or below the
// per-slot set is dropped: set completion can never publish a lower set.
func TestCompleteSetsRevertsWhenTheCompletedSetVerifiesLower(t *testing.T) {
	before, bySlot := completionFixture()
	seen := 0
	fake := &fakeEngine{DPSFunc: func(req api.SimRequest) (float64, error) {
		switch gearKey(req.Character.Gear) {
		case gear(10, 20):
			return 100, nil
		case gear(11, 21):
			seen++
			if seen <= 2 { // the screen and the confirmation
				return 112, nil
			}
			return 90, nil // the second verification
		}
		return 0, nil
	}}
	got, notes, err := completeSets(fake, specInfo{}, "dwarf", "hunter", 60, "", before, bySlot, testCatalog())
	if err != nil {
		t.Fatal(err)
	}
	if got.setDPS != 100 || got.picks["head"].Item.ID != 10 || hasSetBonus(got.picks) {
		t.Fatalf("band = %+v, want the per-slot set back", got)
	}
	if !strings.Contains(strings.Join(notes, "|"), "reverted") {
		t.Errorf("notes = %v, want the revert noted", notes)
	}
}
