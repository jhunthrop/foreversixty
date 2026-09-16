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

// Table is one encounter's mechanics.
type Table struct {
	EncounterID int64      `json:"encounter_id"`
	Name        string     `json:"name"`
	Mechanics   []Mechanic `json:"mechanics"`
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
