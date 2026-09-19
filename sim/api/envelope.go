// Package api holds the request and result envelopes both compute lanes
// speak. They are JSON all the way down.
//
// No protobuf crosses a lane boundary. sim/request turns a SimRequest
// into the engine's RaidSimRequest and sim/adapter turns the result back
// into a summary.Summary; both are Go and both run inside the browser's
// wasm as well as on the server, so neither the web nor the API handler
// ever encodes or decodes a protobuf. An earlier draft of the contract
// carried a Raw []byte field holding the engine request; it was removed
// because it would have forced a protobuf toolchain into the front end
// and a second copy of the adapter's mapping table in TypeScript.
//
// The JSON field names here are authoritative and are mirrored verbatim in
// web/src/lib/sim/types.ts.
package api

import (
	"errors"
	"fmt"
	"slices"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The five character sources.
const (
	SourceArmory = "armory"
	SourceAddon  = "addon"
	SourceBuild  = "build"
	SourceFight  = "fight"
	SourceManual = "manual"
)

// The two compute lanes.
const (
	LaneBrowser = "browser"
	LaneServer  = "server"
)

// ValidIterations is the closed set the UI offers: a planner-inline
// estimate, the default, and the precision toggle.
var ValidIterations = []int{500, 3000, 10000}

// MaxIterations is the largest run anything may ask for. It is the top of
// ValidIterations rather than a second copy of the number, and it bounds
// a split part too: a part is a share of a whole request, so it can never
// legitimately exceed one.
var MaxIterations = slices.Max(ValidIterations)

// MaxTargets is the settings bar's cap.
const MaxTargets = 10

// MinDurationSec and MaxDurationSec bound the fight-length control.
const (
	MinDurationSec = 60
	MaxDurationSec = 480
)

// SimLevel is the only level a sim runs at. It is Forever's level cap
// and it is also the engine's fixed one: sim/core builds every
// character at core.CharacterMaxLevel and its Player protobuf carries
// no level at all, so a level-40 request cannot be answered - it would
// be simulated at 60 and reported without a word. Validating it here
// keeps that refusal at the request rather than at the worker.
const SimLevel = 60

// MaxLevel is Forever's level cap, which is the same number.
const MaxLevel = SimLevel

type SimRequest struct {
	EngineVersion string          `json:"engine_version"`
	Spec          string          `json:"spec"`
	Source        CharacterSource `json:"source"`
	Character     CharacterSpec   `json:"character"`
	Encounter     EncounterSpec   `json:"encounter"`
	Iterations    int             `json:"iterations"`
	RandomSeed    int64           `json:"random_seed"`
}

// CharacterSpec is everything the engine needs about the player, in JSON.
// sim/request turns it into the engine's RaidSimRequest; nothing outside
// Go ever touches a protobuf.
type CharacterSpec struct {
	Name  string `json:"name"`
	Race  string `json:"race"`
	Class string `json:"class"`
	Level int    `json:"level"`
	// Talents is the engine's own talent string, positional against the
	// class's tree sizes, e.g. "01102123133-12312312-".
	Talents    string     `json:"talents"`
	Gear       []GearSlot `json:"gear"`
	Buffs      []string   `json:"buffs"`
	Consumes   []string   `json:"consumes"`
	Profession []string   `json:"professions,omitempty"`
}

// GearSlot is one equipped item. Slot names are the planner's.
type GearSlot struct {
	Slot    string `json:"slot"`
	ItemID  int    `json:"item_id"`
	Enchant int    `json:"enchant,omitempty"`
	Suffix  int    `json:"suffix,omitempty"`
}

type CharacterSource struct {
	Kind       string `json:"kind"`
	Ref        string `json:"ref"`
	CapturedAt string `json:"captured_at"`
}

type EncounterSpec struct {
	DurationSec  int     `json:"duration_sec"`
	Variation    float64 `json:"variation"`
	Targets      int     `json:"targets"`
	ExecuteRatio float64 `json:"execute_ratio"`
	Profile      string  `json:"profile"`
}

// DefaultEncounter is the settings bar's opening state: a three-minute
// single-target fight with the standard duration variation and execute
// window.
func DefaultEncounter() EncounterSpec {
	return EncounterSpec{DurationSec: 180, Variation: 0.2, Targets: 1, ExecuteRatio: 0.25}
}

// Validate checks everything a malformed client could get wrong, at the
// boundary, before anything reaches the engine.
func (r SimRequest) Validate() error {
	return r.validate(true)
}

// ValidatePart is Validate for one worker's share of a split run.
//
// It checks everything Validate checks except the closed iteration set:
// combine.Split divides a request that has already passed Validate, and
// 3,000 iterations over four workers is 750 - a number the settings bar
// never offers and never should. The count still has to be positive and
// no larger than the largest run the UI can ask for, so a part that lost
// its iterations, or one hand-rolled to ask for a million, is refused
// here rather than at the worker.
func (r SimRequest) ValidatePart() error {
	return r.validate(false)
}

// validate is the body of both. closedSet says whether the iteration
// count must be one the settings bar offers.
func (r SimRequest) validate(closedSet bool) error {
	var errs []error
	if r.EngineVersion == "" {
		errs = append(errs, errors.New("engine_version is required"))
	}
	if r.Spec == "" {
		errs = append(errs, errors.New("spec is required"))
	}
	switch {
	case closedSet && !slices.Contains(ValidIterations, r.Iterations):
		errs = append(errs, fmt.Errorf("iterations must be one of %v, got %d", ValidIterations, r.Iterations))
	case !closedSet && (r.Iterations <= 0 || r.Iterations > MaxIterations):
		errs = append(errs, fmt.Errorf("a split part's iterations must be between 1 and %d, got %d", MaxIterations, r.Iterations))
	}
	if r.Encounter.DurationSec < MinDurationSec || r.Encounter.DurationSec > MaxDurationSec {
		errs = append(errs, fmt.Errorf("duration_sec must be between %d and %d, got %d", MinDurationSec, MaxDurationSec, r.Encounter.DurationSec))
	}
	if r.Encounter.Targets < 1 || r.Encounter.Targets > MaxTargets {
		errs = append(errs, fmt.Errorf("targets must be between 1 and %d, got %d", MaxTargets, r.Encounter.Targets))
	}
	if r.Encounter.Variation < 0 || r.Encounter.Variation > 1 {
		errs = append(errs, fmt.Errorf("variation must be between 0 and 1, got %v", r.Encounter.Variation))
	}
	if r.Encounter.ExecuteRatio < 0 || r.Encounter.ExecuteRatio > 1 {
		errs = append(errs, fmt.Errorf("execute_ratio must be between 0 and 1, got %v", r.Encounter.ExecuteRatio))
	}
	if r.Character.Class == "" {
		errs = append(errs, errors.New("character.class is required"))
	}
	if r.Character.Race == "" {
		errs = append(errs, errors.New("character.race is required"))
	}
	if r.Character.Level != SimLevel {
		errs = append(errs, fmt.Errorf("character.level must be %d, got %d; the engine simulates no other level", SimLevel, r.Character.Level))
	}
	return errors.Join(errs...)
}

type SimResult struct {
	SimID         string `json:"sim_id,omitempty"`
	EngineVersion string `json:"engine_version"`
	// Request is stored whole. There is nothing to strip: the envelope
	// carries no protobuf, so a stored row can be re-run by handing it
	// straight back to sim/request.
	Request       SimRequest      `json:"request"`
	Lane          string          `json:"lane"`
	DPS           Estimate        `json:"dps"`
	IterationsRun int             `json:"iterations_run"`
	DurationMS    int64           `json:"duration_ms"`
	Summary       summary.Summary `json:"summary"`
	Error         string          `json:"error,omitempty"`
}

// Progress is what the browser's simRun progress callback carries:
// Pick<SimResult, 'iterations_run' | 'dps'>, so the page reads a partial
// result with the same two accessors it reads the finished one with.
//
// Only DPS.Mean is populated while a run is in flight - it is the
// running mean the engine reports per batch - and the rest of the
// Estimate stays zero until the run completes and a whole SimResult
// replaces it. The callback's first argument is still the callback id
// the page passed to simRun.
type Progress struct {
	IterationsRun int      `json:"iterations_run"`
	DPS           Estimate `json:"dps"`
}

type Estimate struct {
	Mean   float64 `json:"mean"`
	StdDev float64 `json:"stddev"`
	Error  float64 `json:"error"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
}

// Stale reports whether this result came from an engine build other than
// the current one. A stale result is still readable and is labelled in the
// UI; it is never silently re-run.
func (r SimResult) Stale(current string) bool {
	return r.EngineVersion != current
}
