package parquet

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

var base = time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC)

// fixtureEvents decodes the shared v16 fixture, so the Parquet tests cover
// every event shape rather than a hand-built subset.
func fixtureEvents(t *testing.T) []event.Event {
	t.Helper()
	text, err := os.ReadFile("../event/testdata/v16.log")
	if err != nil {
		t.Fatal(err)
	}
	d := event.NewDecoder(layout.RetailV16(), base)
	var out []event.Event
	l := lexer.New()
	emit := func(ln lexer.Line) error {
		out = append(out, d.Decode(ln))
		return nil
	}
	if err := l.Feed(text, 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestWriteIsByteIdenticalAcrossRuns(t *testing.T) {
	evs := fixtureEvents(t)
	first, err := Marshal(evs)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Marshal(evs)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("two writes of the same events differ: %d and %d bytes", len(first), len(second))
	}
	if len(first) == 0 {
		t.Fatal("empty file")
	}
}

func TestRoundTripKeepsEveryTypedField(t *testing.T) {
	evs := fixtureEvents(t)
	b, err := Marshal(evs)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(evs) {
		t.Fatalf("read %d rows, wrote %d", len(got), len(evs))
	}
	byLine := map[int64]event.Event{}
	for _, e := range got {
		byLine[e.Line] = e
	}
	for _, want := range evs {
		g, ok := byLine[want.Line]
		if !ok {
			t.Fatalf("line %d missing after the round trip", want.Line)
		}
		if g.Name != want.Name || g.Kind != want.Kind {
			t.Errorf("line %d: %s/%s, want %s/%s", want.Line, g.Name, g.Kind, want.Name, want.Kind)
		}
		if !g.Time.Equal(want.Time) {
			t.Errorf("line %d time = %s, want %s", want.Line, g.Time, want.Time)
		}
		if g.Amount != want.Amount || g.Overkill != want.Overkill || g.Critical != want.Critical {
			t.Errorf("line %d numbers: %+v %+v %+v", want.Line, g.Amount, g.Overkill, g.Critical)
		}
		if g.Adv != want.Adv {
			t.Errorf("line %d advanced block differs:\n got %+v\nwant %+v", want.Line, g.Adv, want.Adv)
		}
		if g.Source != want.Source || g.Dest != want.Dest || g.Spell != want.Spell {
			t.Errorf("line %d units or spell differ", want.Line)
		}
		if g.ExtraUnit != want.ExtraUnit || g.ExtraSpell != want.ExtraSpell {
			t.Errorf("line %d extra unit/spell differ:\n got %+v/%+v\nwant %+v/%+v",
				want.Line, g.ExtraUnit, g.ExtraSpell, want.ExtraUnit, want.ExtraSpell)
		}
	}
}

func TestNullsSurviveTheRoundTrip(t *testing.T) {
	e := event.Event{Time: base, Line: 1, Name: "SPELL_CAST_START", Kind: event.CastStart}
	b, err := Marshal([]event.Event{e})
	if err != nil {
		t.Fatal(err)
	}
	got, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("rows = %d", len(got))
	}
	if got[0].Amount.OK || got[0].Critical.OK || got[0].Adv.OK {
		t.Fatalf("absent fields came back present: %+v", got[0])
	}
}

func TestRowsComeBackSortedByTimeThenLine(t *testing.T) {
	evs := []event.Event{
		{Time: base.Add(2 * time.Second), Line: 9, Name: "B", Kind: event.Damage},
		{Time: base.Add(time.Second), Line: 4, Name: "A", Kind: event.Damage},
		{Time: base.Add(2 * time.Second), Line: 7, Name: "C", Kind: event.Damage},
	}
	b, err := Marshal(evs)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"A", "C", "B"}
	for i, w := range want {
		if got[i].Name != w {
			t.Fatalf("row %d = %s, want %s (order %v)", i, got[i].Name, w, want)
		}
	}
}

func TestAnEmptyFightWritesAReadableFile(t *testing.T) {
	b, err := Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("rows = %d, want 0", len(got))
	}
}

func TestSummaryRecomputedFromTheEventsFileMatches(t *testing.T) {
	evs := fixtureEvents(t)
	b, err := Marshal(evs)
	if err != nil {
		t.Fatal(err)
	}
	back, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	sum := func(list []event.Event) summary.Summary {
		reg := units.NewRegistry(units.Options{ClassBySpec: units.RetailSpecClass})
		o := summary.DefaultOptions()
		o.Registry = reg
		o.SpecNames = units.RetailSpecName
		a := summary.New(o)
		a.Start(list[0].Time)
		for _, e := range list {
			reg.Observe(e)
			a.Add(e)
		}
		return a.Snapshot(fightFixture(), "test")
	}
	want, got := recomputable(sum(evs)), recomputable(sum(back))
	if !equalJSON(t, want, got) {
		t.Fatal("the summary recomputed from the events file differs from the streamed one")
	}
}

// recomputable strips the parts of a summary that COMBATANT_INFO supplies.
// Gear, talents, spec and item level are structured fields that live in
// summary.json and report.json, not in the events file, so they cannot come
// back from a Parquet round trip and are not part of the property.
func recomputable(s summary.Summary) summary.Summary {
	s.Combatants = nil
	for i := range s.Roster {
		s.Roster[i].SpecID, s.Roster[i].Spec = 0, ""
		s.Roster[i].ItemLevel = 0
		s.Roster[i].Class, s.Roster[i].ClassSource = "", ""
	}
	for i := range s.DamageDone {
		s.DamageDone[i].Class = ""
	}
	for i := range s.DamageTaken {
		s.DamageTaken[i].Class = ""
	}
	for i := range s.Healing {
		s.Healing[i].Class = ""
	}
	for i := range s.HealingTaken {
		s.HealingTaken[i].Class = ""
	}
	for i := range s.Deaths {
		s.Deaths[i].Class = ""
	}
	return s
}

func fightFixture() fight.Fight {
	return fight.Fight{
		Index: 1, Kind: fight.Encounter, EncounterID: 9001, Name: "Warden Kelthas",
		Difficulty: 8, Size: 5, Kill: true,
		Players: []string{
			"Player-4184-000000A1", "Player-4184-000000A2",
			"Player-4184-000000A3", "Player-4184-000000A4",
		},
	}
}

func equalJSON(t *testing.T, a, b summary.Summary) bool {
	t.Helper()
	ja, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	jb, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	return bytes.Equal(ja, jb)
}
