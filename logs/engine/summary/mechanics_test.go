// logs/engine/summary/mechanics_test.go
package summary

import (
	"testing"

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
