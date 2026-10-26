package guilds

import "testing"

// Review Focus 5: after the switch, nights are per character, so two candidates with the same
// present but different nights sort by ratio, not by raw present.
func TestCandidateSortByAttendanceRatioThenPresent(t *testing.T) {
	cands := []LootCandidate{
		{CharacterKey: "5-of-8", GainDps: f64(5), Attendance: Attendance{Present: 5, Nights: 8}},
		{CharacterKey: "5-of-5", GainDps: f64(5), Attendance: Attendance{Present: 5, Nights: 5}},
		{CharacterKey: "6-of-8", GainDps: f64(5), Attendance: Attendance{Present: 6, Nights: 8}},
		{CharacterKey: "3-of-3", GainDps: f64(5), Attendance: Attendance{Present: 3, Nights: 3}},
		{CharacterKey: "0-of-0", GainDps: f64(5), Attendance: Attendance{}},
	}
	sortLootCandidates(cands)
	want := []string{"5-of-5", "3-of-3", "6-of-8", "5-of-8", "0-of-0"}
	if got := candidateKeys(cands); !equalStrings(got, want) {
		t.Fatalf("sorted = %v, want %v (ratio, then present)", got, want)
	}
}

// Each candidate carries its own roster row's nights, not the last row's.
func TestLootCandidatesCarryTheirOwnNights(t *testing.T) {
	s := realTreesStore(t)
	roster := []RosterRow{
		{CharacterKey: "a", Verified: true, className: "warrior", specStr: "Fury", faction: "alliance",
			Attendance: Attendance{Present: 2, Nights: 3}},
		{CharacterKey: "b", Verified: true, className: "warrior", specStr: "Fury", faction: "alliance",
			Attendance: Attendance{Present: 6, Nights: 8}},
	}
	cands := s.lootCandidates(roster, lootItem{ItemID: 16963, Slot: "head"})
	if len(cands) != len(roster) {
		t.Fatalf("candidates = %+v, want one per roster row", cands)
	}
	for _, c := range cands {
		for _, row := range roster {
			if row.CharacterKey == c.CharacterKey && c.Attendance != row.Attendance {
				t.Errorf("%s attendance = %+v, want its own %+v", c.CharacterKey, c.Attendance, row.Attendance)
			}
		}
	}
}
