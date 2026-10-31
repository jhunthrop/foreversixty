package engine

import (
	"slices"
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

func TestRebuildCountsPlayerDeaths(t *testing.T) {
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

// An uploaded fight is counted the way the whole-file parse counts it: a
// group kill logged as PARTY_KILL and UNIT_DIED is one mob death and one of
// the group's, and a mob an outsider finished is a death but not the group's.
func TestRebuildCountsEachMobDeathOnceAndTheGroupsKills(t *testing.T) {
	f, err := NewFixture("rpt")
	if err != nil {
		t.Fatal(err)
	}
	const (
		mob      = "Creature-0-2085-2284-7855-169753-0000BB0001"
		other    = "Creature-0-2085-2284-7855-169753-0000BB0002"
		logger   = "Player-4184-000000A1"
		stranger = "Player-4184-000000F9"
	)
	before, _, err := Rebuild(1, fixtureHeader(), f.Events)
	if err != nil {
		t.Fatal(err)
	}
	at := f.Events[len(f.Events)-1].Time
	hostile := event.Unit{GUID: mob, Flags: 0xa48}
	events := append(append([]event.Event{}, f.Events...),
		event.Event{Time: at, Kind: event.PartyKill, Name: "PARTY_KILL",
			Source: event.Unit{GUID: logger, Flags: 0x511}, Dest: hostile},
		event.Event{Time: at, Kind: event.Death, Name: "UNIT_DIED", Dest: hostile},
		event.Event{Time: at, Kind: event.Damage, Name: "SWING_DAMAGE",
			Source: event.Unit{GUID: stranger, Flags: 0x518}, Dest: event.Unit{GUID: other, Flags: 0xa48}},
		event.Event{Time: at, Kind: event.Death, Name: "UNIT_DIED", Dest: event.Unit{GUID: other, Flags: 0xa48}},
	)
	got, _, err := Rebuild(1, fixtureHeader(), events)
	if err != nil {
		t.Fatal(err)
	}
	if d := got.MobDeaths - before.MobDeaths; d != 2 {
		t.Errorf("mob deaths rose by %d, want 2", d)
	}
	if d := got.GroupKills - before.GroupKills; d != 1 {
		t.Errorf("group kills rose by %d, want 1", d)
	}
	if !slices.Contains(got.Group, logger) || slices.Contains(got.Group, stranger) {
		t.Errorf("group = %v, want the logger and not the stranger", got.Group)
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

// The class and spec the API reads off the spells a player cast, for a log with no
// COMBATANT_INFO (the sample report's beta night, any real report without an addon export).
func TestUnitOptionsReadAClassAndSpecOffTheSpellsCast(t *testing.T) {
	cast := func(r *units.Registry, guid string, ids ...int64) units.Unit {
		for _, id := range ids {
			r.Observe(event.Event{Kind: event.CastSuccess,
				Source: event.Unit{GUID: guid, Name: "Test-Realm", Flags: 0x511},
				Spell:  event.Spell{ID: id}})
		}
		u, _ := r.Get(guid)
		return u
	}
	o := UnitOptions()
	// Rend and Heroic Strike are a warrior's, and neither names a spec.
	w := cast(units.NewRegistry(o), "Player-4184-00000001", 772, 78)
	if w.Class != "Warrior" || w.ClassSource != "inferred" || w.Spec != "" {
		t.Fatalf("warrior = %+v, want Warrior, inferred, no spec", w)
	}
	// Moonfire is a druid's; Swiftmend is a Restoration talent.
	d := cast(units.NewRegistry(o), "Player-4184-00000002", 8921, 18562)
	if d.Class != "Druid" || d.Spec != "Restoration" {
		t.Fatalf("druid = %+v, want Druid, Restoration", d)
	}
	// An item's spell, a profession and a wand are no class's.
	n := cast(units.NewRegistry(o), "Player-4184-00000003", 2881, 8613, 5019, 0)
	if n.Class != "" || n.Spec != "" {
		t.Fatalf("unknown spells gave %+v, want no class and no spec", n)
	}
	// COMBATANT_INFO still wins over the spells.
	r := units.NewRegistry(o)
	cast(r, "Player-4184-00000004", 8921)
	r.Observe(event.Event{Kind: event.CombatantInfo, Combatant: &event.Combatant{GUID: "Player-4184-00000004", SpecID: 73}})
	u, _ := r.Get("Player-4184-00000004")
	if u.Class != "Warrior" || u.ClassSource != "combatant_info" {
		t.Fatalf("combatant info lost to the spells: %+v", u)
	}
}

func TestEverySpellIdIsOneClassesOnly(t *testing.T) {
	if len(ClassBySpell()) < 3000 || len(SpecBySpell()) < 300 {
		t.Fatalf("tables are thin: %d class spells, %d spec spells", len(ClassBySpell()), len(SpecBySpell()))
	}
	for id := range SpecBySpell() {
		if _, ok := ClassBySpell()[id]; !ok {
			t.Fatalf("spec spell %d has no class", id)
		}
	}
}
