// logs/engine/summary/phases_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// phaseFight is the fight the phase tests snapshot: sixty seconds of Warden Kelthas.
func phaseFight() fight.Fight {
	return fight.Fight{Index: 1, Kind: fight.Encounter, EncounterID: 9001, Name: "Warden Kelthas",
		Start: at(0), End: at(60), Players: []string{tank, mage}}
}

func withPhases(t *testing.T, phases []mechanics.Phase) (Options, *units.Registry) {
	t.Helper()
	o, reg := opts(t)
	o.Mechanics = &mechanics.Table{EncounterID: 9001, Name: "Warden Kelthas", Phases: phases}
	return o, reg
}

// The cast the table names opens the phase, and the phase before it runs up to that
// instant: the bands on the chart and the presets in the strip are these spans.
func TestPhasesRunFromEachTriggerToTheNext(t *testing.T) {
	o, reg := withPhases(t, []mechanics.Phase{
		{Name: "Phase 2", Starts: mechanics.PhaseStart{SpellID: 334653, On: mechanics.OnCastStart}},
	})
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(1, mage, boss, 116, "Frostbolt", 100, -1),
		{Time: at(14), Kind: event.CastStart, Name: "SPELL_CAST_START",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Spell: event.Spell{ID: 334653, Name: "Anima Surge"}},
		// A second cast of the same spell is not a second phase.
		{Time: at(30), Kind: event.CastStart, Name: "SPELL_CAST_START",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Spell: event.Spell{ID: 334653, Name: "Anima Surge"}},
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(phaseFight(), "test")
	want := []Phase{{Name: "Phase 1", StartMS: 0, EndMS: 14000}, {Name: "Phase 2", StartMS: 14000, EndMS: 60000}}
	if len(s.Phases) != len(want) {
		t.Fatalf("phases = %+v, want %+v", s.Phases, want)
	}
	for i, p := range want {
		if s.Phases[i] != p {
			t.Errorf("phase %d = %+v, want %+v", i, s.Phases[i], p)
		}
	}
}

// Each trigger kind, one fight each, so a curator can trust all four.
func TestEveryTriggerKindOpensItsPhase(t *testing.T) {
	cases := []struct {
		name  string
		start mechanics.PhaseStart
		event event.Event
	}{
		{"cast success", mechanics.PhaseStart{SpellID: 334653, On: mechanics.OnCastSuccess},
			cast(20, boss, tank, 334653, "Anima Surge")},
		{"aura applied", mechanics.PhaseStart{SpellID: 321038, On: mechanics.OnAuraApplied},
			event.Event{Time: at(20), Kind: event.AuraApplied, Name: "SPELL_AURA_APPLIED",
				Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: mage, Flags: 0x512},
				Spell: event.Spell{ID: 321038, Name: "Wrack Soul"}, AuraType: "DEBUFF"}},
		{"aura removed", mechanics.PhaseStart{SpellID: 321038, On: mechanics.OnAuraRemoved},
			event.Event{Time: at(20), Kind: event.AuraRemoved, Name: "SPELL_AURA_REMOVED",
				Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: mage, Flags: 0x512},
				Spell: event.Spell{ID: 321038, Name: "Wrack Soul"}, AuraType: "DEBUFF"}},
		{"boss health", mechanics.PhaseStart{HealthPct: 50},
			event.Event{Time: at(20), Kind: event.Damage, Name: "SPELL_DAMAGE",
				Source: event.Unit{GUID: mage, Flags: 0x512},
				Dest:   event.Unit{GUID: boss, Name: "Warden Kelthas", Flags: 0xa48},
				Spell:  event.Spell{ID: 116, Name: "Frostbolt", School: 0x10},
				Amount: event.OptInt{V: 10, OK: true}, Overkill: event.OptInt{V: -1, OK: true},
				Adv: event.Advanced{OK: true, InfoGUID: boss, CurrentHP: 400000, MaxHP: 1000000}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o, reg := withPhases(t, []mechanics.Phase{{Name: "Phase 2", Starts: c.start}})
			a := New(o)
			a.Start(at(0))
			// The boss must be named before the health trigger can recognise it.
			naming := dmg(1, mage, boss, 116, "Frostbolt", 1, -1)
			naming.Dest.Name = "Warden Kelthas"
			for _, e := range []event.Event{naming, c.event} {
				reg.Observe(e)
				a.Add(e)
			}
			s := a.Snapshot(phaseFight(), "test")
			if len(s.Phases) != 2 || s.Phases[1].Name != "Phase 2" || s.Phases[1].StartMS != 20000 {
				t.Fatalf("phases = %+v", s.Phases)
			}
		})
	}
}

// A pull that never reached the trigger has no phases to read, and says so with an empty
// list rather than one band called Phase 1 across the whole pull.
func TestAFightThatNeverLeftTheOpeningHasNoPhases(t *testing.T) {
	o, reg := withPhases(t, []mechanics.Phase{
		{Name: "Phase 2", Starts: mechanics.PhaseStart{SpellID: 334653, On: mechanics.OnCastStart}},
	})
	a := New(o)
	a.Start(at(0))
	e := dmg(1, mage, boss, 116, "Frostbolt", 100, -1)
	reg.Observe(e)
	a.Add(e)
	s := a.Snapshot(phaseFight(), "test")
	if len(s.Phases) != 0 {
		t.Fatalf("phases = %+v, want none", s.Phases)
	}
}

// An encounter with no table at all, which is most of them, emits an empty list.
func TestAnEncounterWithNoTableHasAnEmptyPhaseList(t *testing.T) {
	_, _, s := build(t)
	if s.Phases == nil {
		t.Fatal("phases must be an empty array, never null")
	}
	if len(s.Phases) != 0 {
		t.Fatalf("phases = %+v, want none", s.Phases)
	}
}
