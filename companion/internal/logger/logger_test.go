package logger

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

func player(guid, name string, flags uint32) units.Unit {
	return units.Unit{GUID: guid, Name: name, Kind: units.KindPlayer, Flags: flags}
}

func TestTheLoggerIsTheOnePlayerWhoseFlagsSayMine(t *testing.T) {
	logger := player("Player-1-A", "Bow Jackzon-Classic Beta PvP-", 0x511)
	logger.Class, logger.ClassSource = "Warrior", "combatant_info"
	got, ok := Detect([]units.Unit{
		player("Player-1-B", "Gwaihir-ClassicBetaPvP-", 0x518), // outsider
		player("Player-1-C", "Lily-ClassicBetaPvP-", 0x512),    // party
		player("Player-1-D", "Nk-ClassicBetaPvP-", 0x514),      // raid
		logger,
	})
	if !ok {
		t.Fatal("no logger found")
	}
	if got.Name != "Bow Jackzon" || got.Realm != "ClassicBetaPvP" || got.Class != "Warrior" {
		t.Errorf("Detect = %+v", got)
	}
}

func TestNobodyIsTheLoggerUntilSomeoneIsMine(t *testing.T) {
	if got, ok := Detect([]units.Unit{player("P-1", "Lily-X-", 0x512), player("P-2", "Nk-X-", 0x514)}); ok {
		t.Errorf("Detect = %+v with no MINE player", got)
	}
	if _, ok := Detect(nil); ok {
		t.Error("Detect found a logger in an empty registry")
	}
}

func TestAPetIsNeverTheLogger(t *testing.T) {
	pet := units.Unit{GUID: "Pet-0-1", Name: "Rex", Kind: units.KindPet, Flags: 0x1111}
	owned := player("Player-1-X", "Ghost-Realm-", 0x511)
	owned.OwnerGUID = "Player-1-A"
	if got, ok := Detect([]units.Unit{pet, owned}); ok {
		t.Errorf("Detect = %+v from a pet and an owned unit", got)
	}
	me := player("Player-1-A", "Xuen-ClassicBetaPvP-", 0x511)
	if got, ok := Detect([]units.Unit{pet, owned, me}); !ok || got.Name != "Xuen" {
		t.Errorf("Detect = %+v, %v", got, ok)
	}
}

func TestTheNewestMineWinsWhenTheLogSpansTwoCharacters(t *testing.T) {
	old := player("Player-1-A", "Xuen-ClassicBetaPvP-", 0x511)
	old.LastSeen = at(1)
	now := player("Player-1-B", "Cobalt-ClassicBetaPvP2-", 0x511)
	now.LastSeen = at(2)
	if got, _ := Detect([]units.Unit{now, old}); got.Name != "Cobalt" {
		t.Errorf("Detect = %+v, want the newer", got)
	}
}

func TestTheUnitNameSplitsIntoNameAndRealm(t *testing.T) {
	for _, tc := range []struct{ in, name, realm string }{
		{"Xuen-ClassicBetaPvP-", "Xuen", "ClassicBetaPvP"},
		{"Cobalt-ClassicBetaPvP2-", "Cobalt", "ClassicBetaPvP2"},
		{"Bow Jackzon", "Bow Jackzon", ""},
		{"Bow Jackzon-", "Bow Jackzon", ""},
		{"Bow Jackzon-Classic Beta PvP", "Bow Jackzon", "ClassicBetaPvP"},
		{"Thrall-Ashbringer", "Thrall", "Ashbringer"},
		{"  Thrall-Ashbringer-  ", "Thrall", "Ashbringer"},
		{"", "", ""},
		{"Bow-Jackzon-ClassicBetaPvP-", "Bow Jackzon", "ClassicBetaPvP"},
		{"Mörk-ClassicBetaPvP", "Mörk", "ClassicBetaPvP"},
		{"Xuen", "Xuen", ""},
		{"На-Бутылку-ClassicBetaPvP2-", "На Бутылку", "ClassicBetaPvP2"},
	} {
		name, realm := Split(tc.in)
		if name != tc.name || realm != tc.realm {
			t.Errorf("Split(%q) = %q, %q; want %q, %q", tc.in, name, realm, tc.name, tc.realm)
		}
	}
}

func rec(name, realm, region, class string, preferred bool) Record {
	return Record{Name: name, Realm: realm, Region: region, Class: class, Preferred: preferred}
}

func TestOneMatchingRecordIsTheCharacter(t *testing.T) {
	got, ok := Resolve(Detected{Name: "Bow Jackzon", Realm: "ClassicBetaPvP", Class: "Warrior"}, []Record{
		rec("Morrowlyn", "Hardcore", "us", "mage", false),
		rec("bow jackzon", "Classic Beta PvP", "US", "warrior", false),
	})
	want := character.Character{Region: "us", Ruleset: "pvp", Name: "bow jackzon"}
	if !ok || got != want {
		t.Errorf("Resolve = %+v, %v; want %+v", got, ok, want)
	}
}

func TestTheRecordSpellsTheNameTheWayTheAddonSavedIt(t *testing.T) {
	got, ok := Resolve(Detected{Name: "Bow Jackzon"}, []Record{rec("Bow Jackzon", "Hardcore", "eu", "", false)})
	want := character.Character{Region: "eu", Ruleset: "hardcore", Name: "Bow Jackzon"}
	if !ok || got != want {
		t.Errorf("Resolve = %+v, %v; want %+v", got, ok, want)
	}
}

func TestNoRecordMeansNoCharacter(t *testing.T) {
	if got, ok := Resolve(Detected{Name: "Bow"}, []Record{rec("Morrowlyn", "pvp", "us", "mage", false)}); ok {
		t.Errorf("Resolve = %+v", got)
	}
	if _, ok := Resolve(Detected{Name: "Bow"}, nil); ok {
		t.Error("Resolve found a character with no records")
	}
	// A record with no region cannot name a character the API accepts.
	if got, ok := Resolve(Detected{Name: "Bow"}, []Record{rec("Bow", "pvp", "", "", false)}); ok {
		t.Errorf("Resolve = %+v from a record with no region", got)
	}
}

func TestTwoInstallsWithTheSameNameAreToldApartByClass(t *testing.T) {
	recs := []Record{
		rec("Bow", "ClassicBetaPvP", "us", "mage", false),
		rec("Bow", "ClassicBetaPvP", "us", "warrior", false),
	}
	got, ok := Resolve(Detected{Name: "Bow", Class: "Warrior"}, recs)
	if !ok || got.Name != "Bow" {
		t.Fatalf("Resolve = %+v, %v", got, ok)
	}
	if _, ok := Resolve(Detected{Name: "Bow"}, recs); ok {
		t.Error("two records of different classes resolved with no class to choose by")
	}
	if _, ok := Resolve(Detected{Name: "Bow", Class: "Priest"}, recs); ok {
		t.Error("a class neither record has resolved")
	}
	// A class with a space and the addon's spelling of it are the same class.
	if _, ok := Resolve(Detected{Name: "Bow", Class: "Death Knight"},
		[]Record{rec("Bow", "pvp", "us", "deathknight", false)}); !ok {
		t.Error("Death Knight did not match deathknight")
	}
}

func TestTheRealmAndTheWatchedInstallBreakTies(t *testing.T) {
	recs := []Record{
		rec("Bow", "Classic Beta PvP", "us", "", false),
		rec("Bow", "Hardcore", "eu", "", false),
	}
	got, ok := Resolve(Detected{Name: "Bow", Realm: "ClassicBetaPvP"}, recs)
	if !ok || got.Region != "us" || got.Ruleset != "pvp" {
		t.Errorf("by realm: Resolve = %+v, %v", got, ok)
	}
	if _, ok := Resolve(Detected{Name: "Bow"}, recs); ok {
		t.Error("two records and no realm resolved")
	}
	recs[1].Preferred = true
	got, ok = Resolve(Detected{Name: "Bow"}, recs)
	if !ok || got.Region != "eu" {
		t.Errorf("by watched install: Resolve = %+v, %v", got, ok)
	}
	// A realm no record has narrows nothing, so it does not veto.
	if got, ok := Resolve(Detected{Name: "Bow", Realm: "Elsewhere"},
		[]Record{rec("Bow", "Hardcore", "us", "", false)}); !ok || got.Ruleset != "hardcore" {
		t.Errorf("unknown realm: Resolve = %+v, %v", got, ok)
	}
}

func TestRecordsCarryTheClassOutOfTheExportString(t *testing.T) {
	got := Records([]addon.Export{
		{Name: "Bow", Ruleset: "Classic Beta PvP", Region: "us", Export: "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627"},
		{Name: "Odd", Ruleset: "pvp", Region: "us", Export: "garbage"},
	}, true)
	if len(got) != 2 {
		t.Fatalf("Records = %+v", got)
	}
	if got[0] != (Record{Name: "Bow", Realm: "Classic Beta PvP", Region: "us", Class: "warrior", Preferred: true}) {
		t.Errorf("first = %+v", got[0])
	}
	if got[1].Class != "" {
		t.Errorf("a malformed export gave class %q", got[1].Class)
	}
}

func at(n int) time.Time { return time.Date(2026, 10, 19, 20, 0, n, 0, time.UTC) }

// The owner's real log (19 Oct 2026) names the logger "Bow-ClassicBetaPvP-" while the addon
// saved "Bow Jackzon": the log and the WTF folder carry only the first word of a two-word
// Forever name. A one-word log name matches the record whose first word it is; a two-word log
// name never matches by its first word alone.
func TestAOneWordLogNameMatchesTheRecordsFirstWord(t *testing.T) {
	records := []Record{
		rec("Bow Jackzon", "Classic Beta PvP", "us", "hunter", true),
		rec("Booty Twerk", "Classic Beta PvE 2", "us", "paladin", false),
	}
	got, ok := Resolve(Detected{Name: "Bow", Realm: "ClassicBetaPvP"}, records)
	want := character.Character{Region: "us", Ruleset: "pvp", Name: "Bow Jackzon"}
	if !ok || got != want {
		t.Errorf("Resolve(Bow) = %+v, %v; want %+v", got, ok, want)
	}
	if _, ok := Resolve(Detected{Name: "Bow Smith", Realm: "ClassicBetaPvP"}, records); ok {
		t.Errorf("a two-word log name must not match by its first word")
	}
	twins := []Record{rec("Bow Jackzon", "Classic Beta PvP", "us", "", false), rec("Bow Arrow", "Classic Beta PvP", "us", "", false)}
	if _, ok := Resolve(Detected{Name: "Bow"}, twins); ok {
		t.Errorf("two records sharing the first word on the same realm are no answer")
	}
}

// The beta client's record shape against each shape the log has had: the full two-word name
// (current client), the first word alone (the September log), and the spaced spelling.
func TestABetaRecordResolvesFromEveryLogShape(t *testing.T) {
	records := []Record{
		rec("Bow Jackzon", "Classic Beta PvP", "us", "hunter", false),
		rec("Finnias Falcon", "Classic Beta PvP", "us", "mage", false),
	}
	want := character.Character{Region: "us", Ruleset: "pvp", Name: "Bow Jackzon"}
	for _, unit := range []string{"Bow-Jackzon-ClassicBetaPvP-", "Bow Jackzon-ClassicBetaPvP-", "Bow-ClassicBetaPvP-"} {
		name, realm := Split(unit)
		got, ok := Resolve(Detected{Name: name, Realm: realm}, records)
		if !ok || got != want {
			t.Errorf("Resolve(%q) = %+v, %v; want %+v", unit, got, ok, want)
		}
	}
}

func TestTheUnmatchedReasonSaysWhetherTheNameOrTheRealmIsMissing(t *testing.T) {
	bow := Detected{Name: "Bow Jackzon", Realm: "ClassicBetaPvP"}
	for _, tc := range []struct {
		desc    string
		d       Detected
		records []Record
		want    Reason
		text    string
	}{
		{"no records", bow, nil, NoRecord,
			"the log names Bow Jackzon, but the addon has not exported that character — log out once with the addon loaded on it"},
		{"another character", bow, []Record{rec("Finnias Falcon", "Classic Beta PvP", "us", "", false)}, NoRecord,
			"the log names Bow Jackzon, but the addon has not exported that character — log out once with the addon loaded on it"},
		{"the first word is not the name", bow, []Record{rec("Bow Smith", "Classic Beta PvP", "us", "", false)}, NoRecord,
			"the log names Bow Jackzon, but the addon has not exported that character — log out once with the addon loaded on it"},
		{"two records, neither on the realm", bow, []Record{
			rec("Bow Jackzon", "Hardcore", "us", "", false), rec("Bow Jackzon", "Hardcore", "eu", "", false)}, NoMatch,
			"the log names Bow Jackzon, but no addon record carries that realm — log out once with the addon loaded"},
		{"a record with no region", bow, []Record{rec("Bow Jackzon", "Classic Beta PvP", "", "", false)}, NoMatch,
			"the log names Bow Jackzon, but no addon record carries that realm — log out once with the addon loaded"},
		{"a match", bow, []Record{rec("Bow Jackzon", "Classic Beta PvP", "us", "", false)}, Matched, ""},
	} {
		_, why := Judge(tc.d, tc.records)
		if why != tc.want || why.Unmatched(tc.d.Name) != tc.text {
			t.Errorf("%s: Judge = %v, %q; want %v, %q", tc.desc, why, why.Unmatched(tc.d.Name), tc.want, tc.text)
		}
	}
}
