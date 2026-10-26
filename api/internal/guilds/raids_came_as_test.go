package guilds

import "testing"

// The night's Joined leaves out the character a signed main came as by its key, never by its
// name: another realm's raider of the same name still joined.
func TestNightOfDedupesCameAsByKey(t *testing.T) {
	in := "in"
	alt, namesake := "us/hardcore/grim", "us/softcore/grim"
	rows := []AttendanceRow{
		{CharacterKey: "us/hardcore/main", Name: "Main", Signed: &in, Outcome: OutcomePresent,
			CameAs: textPtr("Grim"), CameAsKey: textPtr(alt)},
		{CharacterKey: alt, Name: "Grim", Outcome: OutcomePresent},
		{CharacterKey: namesake, Name: "Grim", Outcome: OutcomePresent},
	}
	if n := nightOf(rows, raidMatchMeta{}); n.Joined != 1 || n.Signed != 1 || n.Came != 1 {
		t.Fatalf("night = %+v, want 1 signed who came and 1 joined", n)
	}
}

// computeAttendance records the came_as character's key beside its name.
func TestCameAsCarriesTheCharactersKey(t *testing.T) {
	const main, alt = "us/hardcore/main", "us/hardcore/grim"
	in := attendanceInput{
		signups: []Signup{{CharacterKey: main, Name: "Main", State: SignupIn, Role: "dps", SignedBy: 7}},
		roster: map[string]raidRosterChar{
			main: {RaidCharacter: RaidCharacter{CharacterKey: main, Name: "Main"}, userID: 7},
			alt:  {RaidCharacter: RaidCharacter{CharacterKey: alt, Name: "Grim"}, userID: 7},
		},
		present:    map[string]logPresence{alt: {Name: "Grim", PullsPresent: 3}},
		pullsTotal: 3,
	}
	for _, r := range computeAttendance(in) {
		if r.CharacterKey != main {
			continue
		}
		if r.CameAs == nil || *r.CameAs != "Grim" || r.CameAsKey == nil || *r.CameAsKey != alt {
			t.Fatalf("came_as %v key %v, want Grim / %s", r.CameAs, r.CameAsKey, alt)
		}
		return
	}
	t.Fatal("no row for the signed main")
}
