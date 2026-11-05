package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func confirmedFixture() (map[string]slotPick, map[string][]scored) {
	unconfirmed := scored{candidate: candidate{ID: 12936, Name: "Battleborn Armbraces", Slots: []string{"wrist"}, ClientUnconfirmed: true}, Score: 10}
	confirmed := scored{candidate: candidate{ID: 2000, Name: "Forest Stalker's Bracers", Slots: []string{"wrist"}}, Score: 9}
	picks := map[string]slotPick{"wrist": {Item: &unconfirmed, RunnerUp: &confirmed}}
	return picks, map[string][]scored{"wrist": {unconfirmed, confirmed}}
}

func wristGear(id int) string {
	return gearKey([]api.GearSlot{{Slot: "wrist", ItemID: id}})
}

func runConfirmed(t *testing.T, unconfirmedDPS, confirmedDPS, stdErr float64) (verifiedBand, int) {
	t.Helper()
	picks, bySlot := confirmedFixture()
	return runPicks(t, picks, bySlot, unconfirmedDPS, confirmedDPS, stdErr)
}

func runPicks(t *testing.T, picks map[string]slotPick, bySlot map[string][]scored, unconfirmedDPS, confirmedDPS, stdErr float64) (verifiedBand, int) {
	t.Helper()
	fake := &fakeEngine{DefaultStdErr: stdErr, DPSByGear: map[string]float64{wristGear(12936): unconfirmedDPS, wristGear(2000): confirmedDPS}}
	vb, held, err := preferConfirmedStats(fake, specInfo{}, "dwarf", "warrior", 60, "", verifiedBand{picks: picks, setDPS: unconfirmedDPS}, bySlot)
	if err != nil {
		t.Fatal(err)
	}
	return vb, held
}

func TestUnconfirmedInsideTheErrorLosesToTheConfirmedItem(t *testing.T) {
	vb, held := runConfirmed(t, 104, 100, 3) // gap 4 < hypot(3,3)=4.24
	if held != 1 || vb.picks["wrist"].Item.ID != 2000 {
		t.Fatalf("held=%d wrist=%+v, want the confirmed item and one held back", held, vb.picks["wrist"].Item)
	}
	if vb.picks["wrist"].RunnerUp.ID != 12936 || vb.setDPS != 100 {
		t.Fatalf("runner-up %+v, setDPS %v: want the unconfirmed item kept as runner-up at the measured 100", vb.picks["wrist"].RunnerUp, vb.setDPS)
	}
}

func TestUnconfirmedInsideTheAdoptionMarginLosesHoweverSmallTheError(t *testing.T) {
	// 877 vs 873 with an error of 1: outside the sims' noise, inside the
	// 1% adoption margin, which is the doubt a 1.12-stats item carries.
	vb, held := runConfirmed(t, 877, 873, 1)
	if held != 1 || vb.picks["wrist"].Item.ClientUnconfirmed {
		t.Fatalf("held %d, pick %+v, want the confirmed item", held, vb.picks["wrist"].Item)
	}
}

func TestUnconfirmedBeyondTheErrorWins(t *testing.T) {
	vb, held := runConfirmed(t, 110, 100, 3)
	if held != 0 || vb.picks["wrist"].Item.ID != 12936 {
		t.Fatalf("held=%d wrist=%+v, want the unconfirmed item to keep the slot", held, vb.picks["wrist"].Item)
	}
}

func TestSlotWithOnlyUnconfirmedCandidatesStillPublishesOne(t *testing.T) {
	only := scored{candidate: candidate{ID: 12936, Name: "Battleborn Armbraces", Slots: []string{"wrist"}, ClientUnconfirmed: true}, Score: 10}
	picks := map[string]slotPick{"wrist": {Item: &only}}
	fake := &fakeEngine{DefaultDPS: 100}
	vb, held, err := preferConfirmedStats(fake, specInfo{}, "dwarf", "warrior", 60, "", verifiedBand{picks: picks, setDPS: 100}, map[string][]scored{"wrist": {only}})
	if err != nil || held != 0 || vb.picks["wrist"].Item.ID != 12936 {
		t.Fatalf("err=%v held=%d wrist=%+v, want the lone unconfirmed item published", err, held, vb.picks["wrist"].Item)
	}
}

// A real Forever item still on 1.12 numbers (the client's Item table names it,
// its ItemSparse row has not reached us) is held to the same bar as an
// unconfirmed one: Hand of Justice must beat the client-stated alternative by
// more than the margin to keep the slot.
func TestClassicDBStatsInsideTheMarginLoseToTheClientsOwnStats(t *testing.T) {
	picks, bySlot := classicDBWristPick()
	vb, held := runPicks(t, picks, bySlot, 877, 873, 1)
	if held != 1 || vb.picks["wrist"].Item.ID != 2000 {
		t.Fatalf("held=%d wrist=%+v, want the client-stated item inside the margin", held, vb.picks["wrist"].Item)
	}
	if ru := vb.picks["wrist"].RunnerUp; ru == nil || ru.ID != 12936 {
		t.Fatalf("runner-up %+v, want the classic-db item kept", ru)
	}
}

func TestClassicDBStatsBeyondTheMarginKeepTheSlot(t *testing.T) {
	picks, bySlot := classicDBWristPick()
	vb, held := runPicks(t, picks, bySlot, 910, 873, 1)
	if held != 0 || vb.picks["wrist"].Item.ID != 12936 {
		t.Fatalf("held=%d wrist=%+v, want the classic-db item to keep the slot", held, vb.picks["wrist"].Item)
	}
}

func classicDBWristPick() (map[string]slotPick, map[string][]scored) {
	classic := scored{candidate: candidate{ID: 12936, Name: "Battleborn Armbraces", Slots: []string{"wrist"}, StatsFromClassicDB: true}, Score: 10}
	client := scored{candidate: candidate{ID: 2000, Name: "Client Bracers", Slots: []string{"wrist"}}, Score: 9}
	return map[string]slotPick{"wrist": {Item: &classic, RunnerUp: &client}}, map[string][]scored{"wrist": {classic, client}}
}
