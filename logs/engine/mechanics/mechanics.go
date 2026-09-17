// logs/engine/mechanics/mechanics.go
// Package mechanics holds the curated per-encounter tables that say which of a
// boss's abilities are avoidable, which casts should be interrupted and which
// debuffs should be dispelled. The log only says an ability hit someone; a
// person decided that standing in it was a mistake. The tables ship inside the
// engine so the parse job has them; the drafting tool in cmd/forever-logs
// proposes one from a log, and a curator corrects it before it lands here.
package mechanics

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
)

// Kind is what failing a mechanic means.
type Kind string

const (
	// Avoidable damage should not have been taken.
	Avoidable Kind = "avoidable"
	// Unavoidable damage is the fight's, listed so it is not mistaken for a gap.
	Unavoidable Kind = "unavoidable"
	// Interrupt casts should have been stopped.
	Interrupt Kind = "interrupt"
	// Dispel debuffs should have been removed.
	Dispel Kind = "dispel"
)

var kinds = map[Kind]bool{Avoidable: true, Unavoidable: true, Interrupt: true, Dispel: true}

// Mechanic is one ability the table classifies.
type Mechanic struct {
	SpellID int64  `json:"spell_id"`
	Name    string `json:"name"`
	Kind    Kind   `json:"kind"`
	Note    string `json:"note,omitempty"`
	// Effects, for an interrupt or dispel: the spell ids of what the cast
	// does when it goes through -- the damage tick and the heal a channel
	// carries under their own ids -- so the row can say what it cost. Empty
	// means the cast's own id.
	Effects []int64 `json:"effects,omitempty"`
	// Role, for an avoidable ability that one role is meant to take: "tank"
	// for a frontal cleave the tank faces away from the raid. A hit on anyone
	// else is then that role's problem as much as the victim's.
	Role string `json:"role,omitempty"`
}

// EffectIDs is what a mechanic's cost is read from: its effects, or its own id.
func (m Mechanic) EffectIDs() []int64 {
	if len(m.Effects) > 0 {
		return m.Effects
	}
	return []int64{m.SpellID}
}

// The trigger kinds a phase can start on. cast_start is here alongside the
// spec's three because a boss's phase-opening channel begins the phase when
// the cast begins -- and a channel a raid interrupts never lands at all, so
// keying its phase on the success would be keying it on the raid failing.
const (
	OnCastStart   = "cast_start"
	OnCastSuccess = "cast_success"
	OnAuraApplied = "aura_applied"
	OnAuraRemoved = "aura_removed"
)

var phaseOn = map[string]bool{
	OnCastStart: true, OnCastSuccess: true, OnAuraApplied: true, OnAuraRemoved: true,
}

// PhaseStart is what begins a phase: an enemy's cast or aura change, or the
// boss's own health passing a percentage. Exactly one of the two forms is set.
type PhaseStart struct {
	// SpellID with On: the spell whose cast or aura change starts the phase,
	// on any enemy.
	SpellID int64  `json:"spell_id,omitempty"`
	On      string `json:"on,omitempty"`
	// HealthPct starts the phase the first time the boss's own health is at or
	// below this percentage, read from the advanced block on its lines.
	HealthPct float64 `json:"health_pct,omitempty"`
}

// Phase is one named stretch of an encounter, and what begins it. Phase 1
// starts at the pull and needs no entry.
type Phase struct {
	Name   string     `json:"name"`
	Starts PhaseStart `json:"starts"`
}

// Table is one encounter's mechanics, and the phases it is fought in.
type Table struct {
	EncounterID int64      `json:"encounter_id"`
	Name        string     `json:"name"`
	Mechanics   []Mechanic `json:"mechanics"`
	// Phases are the stretches the encounter is fought in, in the order they
	// happen. Empty for an encounter nobody has curated phases for.
	Phases []Phase `json:"phases,omitempty"`
}

//go:embed tables/*.json
var tables embed.FS

// A malformed embedded table is a build defect, so it fails the process at
// start rather than the first time a fight of that encounter happens to open.
var _ = mustParseAll()

// mustParseAll parses every embedded table and panics on the first one that is
// not the format or whose file name does not match its encounter id. Load
// reads a table by that file name, so a mismatch would hide the table.
func mustParseAll() int {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		panic(fmt.Errorf("mechanics: reading the embedded tables: %w", err))
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := "tables/" + entry.Name()
		data, err := tables.ReadFile(path)
		if err != nil {
			panic(fmt.Errorf("mechanics: reading %s: %w", path, err))
		}
		t, err := Parse(data)
		if err != nil {
			panic(fmt.Errorf("mechanics: %s: %w", path, err))
		}
		if want := strconv.FormatInt(t.EncounterID, 10) + ".json"; entry.Name() != want {
			panic(fmt.Errorf("mechanics: %s holds encounter %d, so it must be named %s", path, t.EncounterID, want))
		}
	}
	return len(entries)
}

// Parse reads a table and refuses one that is not the format.
func Parse(data []byte) (Table, error) {
	var t Table
	if err := json.Unmarshal(data, &t); err != nil {
		return Table{}, fmt.Errorf("mechanics: %w", err)
	}
	if t.EncounterID <= 0 {
		return Table{}, fmt.Errorf("mechanics: encounter_id must be positive")
	}
	seen := make(map[int64]bool, len(t.Mechanics))
	for i, m := range t.Mechanics {
		if m.SpellID <= 0 {
			return Table{}, fmt.Errorf("mechanics[%d]: spell_id must be positive", i)
		}
		if !kinds[m.Kind] {
			return Table{}, fmt.Errorf("mechanics[%d]: kind %q is not avoidable, unavoidable, interrupt or dispel", i, m.Kind)
		}
		// One spell, one classification: a second row for a spell id would be
		// silently ignored by Lookup, so the table is wrong rather than lenient.
		if seen[m.SpellID] {
			return Table{}, fmt.Errorf("mechanics[%d]: spell_id %d is listed twice", i, m.SpellID)
		}
		seen[m.SpellID] = true
	}
	for i, p := range t.Phases {
		if p.Name == "" {
			return Table{}, fmt.Errorf("phases[%d]: name is required", i)
		}
		switch {
		case p.Starts.SpellID > 0:
			if !phaseOn[p.Starts.On] {
				return Table{}, fmt.Errorf(
					"phases[%d]: on %q is not cast_start, cast_success, aura_applied or aura_removed", i, p.Starts.On)
			}
			if p.Starts.HealthPct != 0 {
				return Table{}, fmt.Errorf(
					"phases[%d]: a phase starts on a spell or on a health percentage, not both", i)
			}
		case p.Starts.HealthPct > 0 && p.Starts.HealthPct <= 100:
			if p.Starts.On != "" {
				return Table{}, fmt.Errorf("phases[%d]: on belongs to a spell trigger; a health trigger takes none", i)
			}
		default:
			return Table{}, fmt.Errorf(
				"phases[%d]: starts must name a spell_id with on, or a health_pct above 0 and at most 100", i)
		}
	}
	return t, nil
}

// Load returns the embedded table for an encounter, and false when there is none.
func Load(encounterID int64) (Table, bool) {
	data, err := tables.ReadFile("tables/" + strconv.FormatInt(encounterID, 10) + ".json")
	if err != nil {
		return Table{}, false
	}
	t, err := Parse(data)
	if err != nil {
		// An embedded table that does not parse is a build defect, not a runtime case.
		panic(err)
	}
	return t, true
}

// Lookup finds a mechanic by spell id.
func (t Table) Lookup(spellID int64) (Mechanic, bool) {
	for _, m := range t.Mechanics {
		if m.SpellID == spellID {
			return m, true
		}
	}
	return Mechanic{}, false
}
