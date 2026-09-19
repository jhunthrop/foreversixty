package request

// The fight styles.
//
// A style is a page preset: a name a player picks, and the encounter
// fields it stands for. The expansion lives here, in Go, because the
// page applies it too and a second copy of the table in TypeScript
// would drift the first time a preset was retuned - the page would
// offer "Heavy movement" and the sim would run last month's numbers.
// So this map is the source, `go run ./internal/genstyles` renders it
// to styles.json, and the web lane's test compares its own table
// against that file.
//
// A style sets only the fields it owns. Fight length, variation, target
// level, armor and type are the player's and survive a style change,
// which is what makes the style control a preset rather than a reset.

import (
	"encoding/json"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// style is one preset's expansion, as a delta over the base encounter.
type style struct {
	Targets         int               `json:"targets"`
	ExecuteRatio    float64           `json:"execute_ratio"`
	Movement        *api.Movement     `json:"movement,omitempty"`
	TargetsOverTime []api.TargetCount `json:"targets_over_time,omitempty"`
	Dummy           bool              `json:"dummy,omitempty"`
}

// StyleIDs is the vocabulary, in the order the page lists it.
var StyleIDs = []string{
	"patchwerk",
	"execute",
	"light-movement",
	"heavy-movement",
	"cleave-2",
	"cleave-3",
	"cleave-5",
	"dungeon",
	"dummy",
}

// The movement windows the two movement styles use. Five seconds is
// one global cooldown plus a step; the interval is what makes light
// and heavy different.
const (
	movementWindowSec   = 5
	lightMovementEvery  = 45
	heavyMovementEvery  = 20
	patchwerkExecute    = 0.25
	executeHeavyExecute = 0.35
)

// styles is the contract's section 1.6 table.
var styles = map[string]style{
	"patchwerk": {Targets: 1, ExecuteRatio: patchwerkExecute},
	"execute":   {Targets: 1, ExecuteRatio: executeHeavyExecute},
	"light-movement": {Targets: 1, ExecuteRatio: patchwerkExecute,
		Movement: &api.Movement{IntervalSec: lightMovementEvery, DurationSec: movementWindowSec, Kind: api.MovementAway}},
	"heavy-movement": {Targets: 1, ExecuteRatio: patchwerkExecute,
		Movement: &api.Movement{IntervalSec: heavyMovementEvery, DurationSec: movementWindowSec, Kind: api.MovementAway}},
	"cleave-2": {Targets: 2, ExecuteRatio: patchwerkExecute},
	"cleave-3": {Targets: 3, ExecuteRatio: patchwerkExecute},
	"cleave-5": {Targets: 5, ExecuteRatio: patchwerkExecute},
	// A boss, then packs of three and five, then back down: the shape
	// of a dungeon pull. No execute window, because a pack dies from
	// full health to nothing rather than sliding down a health bar.
	"dungeon": {Targets: 1, ExecuteRatio: 0, TargetsOverTime: []api.TargetCount{
		{AtSec: 0, Count: 1},
		{AtSec: 40, Count: 3},
		{AtSec: 80, Count: 5},
		{AtSec: 130, Count: 3},
		{AtSec: 160, Count: 1},
	}},
	// A training dummy has no debuffs, no execute and no armor
	// reduction; the engine's target_dummy flag is what turns all
	// three off at once.
	"dummy": {Targets: 1, ExecuteRatio: 0, Dummy: true},
}

// ExpandStyle applies a style to an encounter, keeping every field the
// style does not own. The label rides along in Style so a saved sim can
// say which preset produced it.
func ExpandStyle(id string, base api.EncounterSpec) (api.EncounterSpec, bool) {
	s, ok := styles[id]
	if !ok {
		return base, false
	}
	out := base
	out.Style = id
	out.Targets = s.Targets
	out.ExecuteRatio = s.ExecuteRatio
	out.Dummy = s.Dummy
	out.Movement = nil
	if s.Movement != nil {
		// A copy: the map's pointer is shared by every caller, and a
		// caller that edited the encounter it got back would retune the
		// preset for the whole process.
		m := *s.Movement
		out.Movement = &m
	}
	out.TargetsOverTime = slices.Clone(s.TargetsOverTime)
	return out, true
}

// StylesJSON renders the table as the web lane's fixture: an array in
// StyleIDs order, each entry the style id and its expansion.
func StylesJSON() ([]byte, error) {
	type row struct {
		ID string `json:"id"`
		style
	}
	out := make([]row, 0, len(StyleIDs))
	for _, id := range StyleIDs {
		out = append(out, row{ID: id, style: styles[id]})
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(b, '\n'), nil
}
