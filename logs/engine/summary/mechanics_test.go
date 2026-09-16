// logs/engine/summary/mechanics_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// The fixtureEvents fixture: the boss (guid `boss`) hits the tank with Anima
// Lash 334660 twice, the second killing them; the mage casts Frostbolt; see
// summary_test.go.
func TestMechanicsAttributeAvoidableHitsAndTheKill(t *testing.T) {
	o, reg := opts(t)
	o.Mechanics = &mechanics.Table{EncounterID: 9001, Name: "Warden Kelthas", Mechanics: []mechanics.Mechanic{
		{SpellID: 334660, Name: "Anima Lash", Kind: mechanics.Avoidable},
	}}
	a := New(o)
	a.Start(at(0))
	for _, e := range fixtureEvents() {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fixtureFight(), "test")
	if !s.Mechanics.TableFound || len(s.Mechanics.Rows) != 1 {
		t.Fatalf("mechanics = %+v", s.Mechanics)
	}
	row := s.Mechanics.Rows[0]
	if row.SpellID != 334660 || row.Kind != mechanics.Avoidable || len(row.Players) != 1 {
		t.Fatalf("row = %+v", row)
	}
	hit := row.Players[0]
	if hit.GUID != tank || hit.Hits != 2 || !hit.Killed {
		t.Fatalf("tank's hit = %+v, want two hits and the kill", hit)
	}
}

func TestMechanicsWithoutATableSaySo(t *testing.T) {
	_, _, s := build(t)
	if s.Mechanics.TableFound || len(s.Mechanics.Rows) != 0 {
		t.Fatalf("mechanics without a table = %+v", s.Mechanics)
	}
}

func mechanicRowFor(t *testing.T, rows []MechanicRow, spellID int64) MechanicRow {
	t.Helper()
	for _, r := range rows {
		if r.SpellID == spellID {
			return r
		}
	}
	t.Fatalf("no row for spell %d in %+v", spellID, rows)
	return MechanicRow{}
}

// The fixtureEvents fixture: at sec 12 the tank interrupts (Pummel 6552) the
// boss's cast of Anima Surge (334653) - the boss never logs a cast start for
// it, so Casts must be floored up to Stopped rather than left at zero. At
// sec 13 the healer dispels Wrack Soul (321038) off the mage, who was never
// given the debuff by an AURA_APPLIED line, so Applied stays zero while
// Dispelled counts the one exchange.
func TestMechanicsInterruptAndDispelRows(t *testing.T) {
	o, reg := opts(t)
	o.Mechanics = &mechanics.Table{EncounterID: 9001, Name: "Warden Kelthas", Mechanics: []mechanics.Mechanic{
		{SpellID: 334653, Name: "Anima Surge", Kind: mechanics.Interrupt},
		{SpellID: 321038, Name: "Wrack Soul", Kind: mechanics.Dispel},
	}}
	a := New(o)
	a.Start(at(0))
	for _, e := range fixtureEvents() {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fixtureFight(), "test")
	if !s.Mechanics.TableFound || len(s.Mechanics.Rows) != 2 {
		t.Fatalf("mechanics = %+v", s.Mechanics)
	}
	interrupt := mechanicRowFor(t, s.Mechanics.Rows, 334653)
	if interrupt.Kind != mechanics.Interrupt || interrupt.Stopped != 1 || interrupt.Casts != 1 {
		t.Fatalf("interrupt row = %+v, want Stopped 1 and Casts floored up to 1", interrupt)
	}
	dispel := mechanicRowFor(t, s.Mechanics.Rows, 321038)
	if dispel.Kind != mechanics.Dispel || dispel.Dispelled != 1 || dispel.Applied != 0 {
		t.Fatalf("dispel row = %+v, want Dispelled 1 and Applied 0 (no AURA_APPLIED in the fixture)", dispel)
	}
}

// TestMechanicsIgnoreAHitFromThePlayersOwnSide: the same spell id arriving
// from a friendly unit is not the boss's mechanic. Only the boss's two Anima
// Lashes on the tank may be counted; the mage's friendly-fire copy on the
// healer must leave the healer off the row entirely.
func TestMechanicsIgnoreAHitFromThePlayersOwnSide(t *testing.T) {
	o, reg := opts(t)
	o.Mechanics = &mechanics.Table{EncounterID: 9001, Name: "Warden Kelthas", Mechanics: []mechanics.Mechanic{
		{SpellID: 334660, Name: "Anima Lash", Kind: mechanics.Avoidable},
	}}
	a := New(o)
	a.Start(at(0))
	friendlyFire := event.Event{
		Time: at(17), Kind: event.Damage, Name: "SPELL_DAMAGE",
		Source:   event.Unit{GUID: mage, Flags: 0x512},
		Dest:     event.Unit{GUID: healer, Flags: 0x512},
		Spell:    event.Spell{ID: 334660, Name: "Anima Lash", School: 0x10},
		Amount:   event.OptInt{V: 700, OK: true},
		Overkill: event.OptInt{V: -1, OK: true},
	}
	for _, e := range append(fixtureEvents(), friendlyFire) {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fixtureFight(), "test")
	row := mechanicRowFor(t, s.Mechanics.Rows, 334660)
	if len(row.Players) != 1 || row.Players[0].GUID != tank {
		t.Fatalf("players = %+v, want only the tank: the mage's friendly copy is not the boss's mechanic", row.Players)
	}
	if row.Players[0].Hits != 2 {
		t.Errorf("tank hits = %d, want the boss's 2", row.Players[0].Hits)
	}
}
