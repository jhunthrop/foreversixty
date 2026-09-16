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
}

// Table is one encounter's mechanics.
type Table struct {
	EncounterID int64      `json:"encounter_id"`
	Name        string     `json:"name"`
	Mechanics   []Mechanic `json:"mechanics"`
}

//go:embed tables/*.json
var tables embed.FS

// Parse reads a table and refuses one that is not the format.
func Parse(data []byte) (Table, error) {
	var t Table
	if err := json.Unmarshal(data, &t); err != nil {
		return Table{}, fmt.Errorf("mechanics: %w", err)
	}
	if t.EncounterID <= 0 {
		return Table{}, fmt.Errorf("mechanics: encounter_id must be positive")
	}
	for i, m := range t.Mechanics {
		if m.SpellID <= 0 {
			return Table{}, fmt.Errorf("mechanics[%d]: spell_id must be positive", i)
		}
		if !kinds[m.Kind] {
			return Table{}, fmt.Errorf("mechanics[%d]: kind %q is not avoidable, unavoidable, interrupt or dispel", i, m.Kind)
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
