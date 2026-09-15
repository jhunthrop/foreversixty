package engine

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

func TestFixtureParsesOneKilledEncounter(t *testing.T) {
	f, err := NewFixture("rpt")
	if err != nil {
		t.Fatal(err)
	}
	if f.Fight.EncounterID != 9001 || !f.Fight.Kill {
		t.Fatalf("fight = %+v, want encounter 9001 killed", f.Fight)
	}
	if got := len(f.Fight.Players); got != 3 {
		t.Fatalf("players = %d, want 3", got)
	}
	if f.Health.ParseErrors != 0 || len(f.Health.UnknownEvents) != 0 {
		t.Fatalf("health = %+v, want a clean parse", f.Health)
	}
	if !f.Health.AdvancedLogging {
		t.Fatal("advanced logging should be detected in the fixture header")
	}
}

// fixtureHeader is the header a companion posts alongside the Parquet:
// the fields the Parquet round trip drops.
func fixtureHeader() Header {
	return Header{EncounterID: 9001, Name: "Warden Kelthas", Difficulty: 8, Size: 5, Kill: true}
}

func TestRebuildFromParquetMatchesTheOriginalSummary(t *testing.T) {
	f, err := NewFixture("rpt")
	if err != nil {
		t.Fatal(err)
	}
	events, err := parquet.Unmarshal(f.Parquet)
	if err != nil {
		t.Fatal(err)
	}
	got, sum, err := Rebuild(7, fixtureHeader(), events)
	if err != nil {
		t.Fatal(err)
	}
	if got.Index != 7 {
		t.Fatalf("index = %d, want the route's 7", got.Index)
	}
	if len(got.Players) != len(f.Fight.Players) {
		t.Fatalf("players = %v, want %v", got.Players, f.Fight.Players)
	}
	if sum.DurationMS != f.Summary.DurationMS {
		t.Fatalf("duration = %d, want %d", sum.DurationMS, f.Summary.DurationMS)
	}
	if len(sum.Roster) != len(f.Summary.Roster) {
		t.Fatalf("roster = %d rows, want %d", len(sum.Roster), len(f.Summary.Roster))
	}
	for i, r := range sum.Roster {
		w := f.Summary.Roster[i]
		if r.GUID != w.GUID || r.DamageDone != w.DamageDone || r.HealingDone != w.HealingDone ||
			r.DamageTaken != w.DamageTaken || r.ActiveMS != w.ActiveMS || r.Role != w.Role {
			t.Fatalf("roster[%d] = %+v, want %+v", i, r, w)
		}
	}
}

func TestRebuildRejectsAnEmptyBundle(t *testing.T) {
	if _, _, err := Rebuild(0, fixtureHeader(), nil); err != ErrNoEvents {
		t.Fatalf("err = %v, want ErrNoEvents", err)
	}
}

func TestRebuildCountsPlayerDeathsAndNPCKills(t *testing.T) {
	f, err := NewFixture("rpt")
	if err != nil {
		t.Fatal(err)
	}
	events := append([]event.Event{}, f.Events...)
	last := events[len(events)-1]
	events = append(events, event.Event{
		Time: last.Time, Kind: event.Death,
		Dest: event.Unit{GUID: "Player-4184-000000A1", Name: "Baelgrim-Nightslayer"},
	})
	got, _, err := Rebuild(1, fixtureHeader(), events)
	if err != nil {
		t.Fatal(err)
	}
	if got.Deaths != 1 {
		t.Fatalf("deaths = %d, want 1", got.Deaths)
	}
}

func TestSummaryOptionsCarryTheEngineSpecTables(t *testing.T) {
	o := SummaryOptions(units.NewRegistry(UnitOptions()))
	if o.Registry == nil {
		t.Fatal("registry must be set")
	}
	if len(o.SpecNames) == 0 {
		t.Fatal("spec names must come from the engine's table")
	}
	if o.Bucket == 0 || o.ActiveGap == 0 {
		t.Fatalf("options = %+v, want the engine defaults filled in", o)
	}
}

func TestLayoutByNameFindsTheEnginesRows(t *testing.T) {
	if _, ok := LayoutByName("retail-v16"); !ok {
		t.Fatal("the retail row should be registered")
	}
	if _, ok := LayoutByName("no-such-layout"); ok {
		t.Fatal("an unknown name must not resolve")
	}
}

func TestSessionOptionsForPinsTheLayoutWhenItIsKnown(t *testing.T) {
	o := SessionOptionsFor("r", FixtureBase, true, "retail-v16")
	if o.Layout.Name != "retail-v16" || !o.KeepEvents {
		t.Fatalf("options = %+v", o)
	}
	if o = SessionOptionsFor("r", FixtureBase, false, ""); o.Layout.Name != "" {
		t.Fatalf("an empty name should leave the session to infer: %+v", o.Layout)
	}
}

func TestNameFromEventsTakesTheBossTheDamageWentInto(t *testing.T) {
	f, err := NewFixture("rpt")
	if err != nil {
		t.Fatal(err)
	}
	if got := NameFromEvents(Header{Name: "Warden Kelthas"}, f.Summary); got != "Warden Kelthas" {
		t.Fatalf("a header that carries a name wins: %q", got)
	}
	if got := NameFromEvents(Header{}, f.Summary); got != "Trash" {
		t.Fatalf("a fight with no encounter is trash: %q", got)
	}
	got := NameFromEvents(Header{EncounterID: 9001}, f.Summary)
	if got != "Hollow Sentinel" {
		t.Fatalf("name = %q, want the hostile unit that took the damage", got)
	}
	if got := NameFromEvents(Header{EncounterID: 9001}, summary.Summary{}); got != "Encounter 9001" {
		t.Fatalf("with nothing to go on the id is the name: %q", got)
	}
}

func TestTheFixtureLogIsAvailableWhole(t *testing.T) {
	if len(FixtureLog()) == 0 {
		t.Fatal("the fixture log should not be empty")
	}
	// Nine: the header, six spell lines and the tank's one swing, which a real log writes
	// twice (SWING_DAMAGE, then SWING_DAMAGE_LANDED with the target's health).
	if FixtureLines() != 9 {
		t.Fatalf("lines = %d, want the fixture's nine", FixtureLines())
	}
}
