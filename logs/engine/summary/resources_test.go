// logs/engine/summary/resources_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
)

// energize is one power gain: what was gained, what was gained past the cap, and the
// reading the advanced block leaves the bar at.
func energize(sec float64, guid string, powerType, amount, over, current, max int64) event.Event {
	return event.Event{
		Time: at(sec), Kind: event.Energize, Name: "SPELL_ENERGIZE",
		Source: event.Unit{GUID: guid, Flags: 0x512}, Dest: event.Unit{GUID: guid, Flags: 0x512},
		Spell:        event.Spell{ID: 34428, Name: "Victory Rush"},
		Amount:       event.OptInt{V: amount, OK: true},
		OverEnergize: event.OptInt{V: over, OK: true},
		PowerType:    event.OptInt{V: powerType, OK: true},
		MaxPower:     event.OptInt{V: max, OK: true},
		Adv: event.Advanced{OK: true, InfoGUID: guid, PowerType: powerType,
			CurrentPower: current, MaxPower: max},
	}
}

// A rage bar that fills and stays full is the thing this table exists to show: every
// point poured into it after that is a point thrown away, and the seconds it spent full
// are seconds the player had a resource they were not spending.
func TestResourcesRecordTheCapTheTimeAtItAndWhatWasWastedPastIt(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		energize(0, tank, 1, 60, 0, 60, 100),
		energize(1, tank, 1, 40, 25, 100, 100),
		energize(2, tank, 1, 10, 10, 100, 100),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(3),
		Players: []string{tank}}, "test")

	if len(s.Resources) != 1 {
		t.Fatalf("resources = %+v", s.Resources)
	}
	tr := s.Resources[0]
	if tr.Max != 100 {
		t.Errorf("max = %d, want 100 (the largest maximum any line reported)", tr.Max)
	}
	// Two of the three seconds read 100 of 100.
	if tr.AtMaxMS != 2000 {
		t.Errorf("at_max_ms = %d, want 2000", tr.AtMaxMS)
	}
	if tr.Wasted != 35 {
		t.Errorf("wasted = %d, want 35 (25 and 10 past the cap)", tr.Wasted)
	}
	if tr.Gained != 110 {
		t.Errorf("gained = %d, want 110", tr.Gained)
	}
}

// A track the log never gave a maximum for has no cap to draw and no time at one: zero,
// not the peak standing in for a cap nobody reported.
func TestResourcesWithNoReportedMaximumHaveNoCapAndNoTimeAtIt(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	e := event.Event{
		Time: at(0), Kind: event.Damage, Name: "SPELL_DAMAGE",
		Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: event.Unit{GUID: boss, Flags: 0xa48},
		Spell:    event.Spell{ID: 116, Name: "Frostbolt", School: 0x10},
		Amount:   event.OptInt{V: 100, OK: true},
		Overkill: event.OptInt{V: -1, OK: true},
		Adv:      event.Advanced{OK: true, InfoGUID: mage, PowerType: 0, CurrentPower: 500, MaxPower: 0},
	}
	reg.Observe(e)
	a.Add(e)
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(1),
		Players: []string{mage}}, "test")
	for _, tr := range s.Resources {
		if tr.Max != 0 || tr.AtMaxMS != 0 {
			t.Errorf("track %+v: with no reported maximum both max and at_max_ms must be zero", tr)
		}
	}
}
