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
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// The five character sources.
const (
	SourceArmory = "armory"
	SourceAddon  = "addon"
	SourceBuild  = "build"
	SourceFight  = "fight"
	SourceManual = "manual"
)

// Sources is the closed set Validate checks CharacterSource.Kind
// against. Declaring the constants without ever comparing to them let a
// typo through as a stored row nothing could ever join on.
var Sources = []string{SourceArmory, SourceAddon, SourceBuild, SourceFight, SourceManual}

// The encounter profiles the contract names.
//
// ProfilePatchwerk is a stationary target that does nothing but be hit,
// which is exactly what the sim already builds, so it and the empty
// string are the same fight. ProfileEncounterPrefix names a curated
// encounter; nothing simulates one yet - there is no zone-to-biome
// table and no mechanics model - so it is refused at the boundary
// rather than run as a patchwerk under an encounter's name.
const (
	ProfilePatchwerk       = "patchwerk"
	ProfileEncounterPrefix = "encounter:"
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

// StepIterations is how many iterations a target-error run adds at a
// time. It is one number for both lanes: the browser's pool runs a step
// as an ordinary sharded run and the native binary runs it in one go,
// and a step of a different size in each would make the two report
// different iteration counts for the same request.
const StepIterations = 1000

// LaneIterationCeiling bounds a target-error run that never reaches its
// target. Without it a flat distribution would run until the tab was
// closed.
var LaneIterationCeiling = map[string]int{LaneBrowser: 30000, LaneServer: 100000}

// MaxTargets is the settings bar's cap.
const MaxTargets = 10

// GearSlots is the slot vocabulary the envelope validates against. It is
// the planner's own list and the order sim/request builds the engine's
// positional equipment array in, which is why that package holds the
// table and this one names it; sim/request's unexported slotOrder is
// its source and TestGearSlotsMatchTheRequestTable proves the two
// agree.
var GearSlots = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
	"legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
	"main_hand", "off_hand", "ranged",
}

// MinDurationSec and MaxDurationSec bound the fight-length control.
// Twenty seconds is the shortest fight the ramp-up of any rotation
// says anything about; ten minutes is longer than any encounter in
// the game and is what a target dummy run is capped at.
const (
	MinDurationSec = 20
	MaxDurationSec = 600
)

// SimLevel is the only level a sim runs at. It is Forever's level cap
// and it is also the engine's fixed one: sim/core builds every
// character at core.CharacterMaxLevel and its Player protobuf carries
// no level at all, so a level-40 request cannot be answered - it would
// be simulated at 60 and reported without a word. Validating it here
// keeps that refusal at the request rather than at the worker.
const SimLevel = 60

// BossLevel is the level of every target a sim fights: three above the
// player, which is what the attack table's suppression terms are
// derived against. sim/request builds the encounter at this level and
// sim/measure refuses to fit a boss-level attack table from a log that
// never saw one, so the two have to be the same number.
const BossLevel = SimLevel + 3

type SimRequest struct {
	EngineVersion string          `json:"engine_version"`
	Spec          string          `json:"spec"`
	Source        CharacterSource `json:"source"`
	Character     CharacterSpec   `json:"character"`
	Encounter     EncounterSpec   `json:"encounter"`
	Iterations    int             `json:"iterations"`
	RandomSeed    int64           `json:"random_seed"`
	// Bulk turns one character into many sims. A request without it is
	// exactly today's single run; see sim/api/bulk.go.
	Bulk *BulkSpec `json:"bulk,omitempty"`
	// Weights asks for stat weights instead of a DPS run.
	Weights *WeightsSpec `json:"weights,omitempty"`
	// TargetError, when above zero, turns Iterations into a CEILING:
	// the run continues in steps of StepIterations until DPS.Error over
	// DPS.Mean is at or under it, or Iterations is reached. Zero is
	// today's fixed-count run.
	//
	// This is Raidbots' Smart Sim, made visible. The results line says
	// which of the two ended the run, which is why the loop is a loop
	// over whole results rather than a number the engine is handed.
	TargetError float64 `json:"target_error,omitempty"`
	// NoSample says this request's cast log would never be read, so
	// the engine should not give up its concurrent entry point to
	// replay one fight and record it.
	//
	// It is not a user setting and the page never sends it. Three
	// callers do: sim/bulk sets it on every stage request, because a
	// stage is a plain run by shape and has to be able to say it is
	// not the run whose sample the report will show; sim/combine sets
	// it on every part of a split but the first, because Results keeps
	// part zero's sample and the rest would be replayed and thrown
	// away; and a batch caller of forever-sim - the nightly validation
	// job, the execution scorer - sets it through -no-sample, for the
	// same reason on thousands of runs.
	NoSample bool `json:"no_sample,omitempty"`
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
	// Cooldowns is when to use the major cooldowns and potions the
	// rotation would otherwise fire on cooldown.
	Cooldowns []CooldownSpec `json:"cooldowns,omitempty"`
}

// CooldownSpec pins one cooldown's timings.
type CooldownSpec struct {
	// ID is "spell:<id>", "item:<id>", or a consumable id from IDS.md,
	// which the build's consumable table resolves to an item.
	ID string `json:"id"`
	// AtSec are the times to use it, in fight seconds. Each value is
	// one usage; usages past the list happen as soon as the rotation
	// allows. An empty list is "on cooldown", which is the default the
	// engine already has.
	AtSec []float64 `json:"at_sec,omitempty"`
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

	// Style is the page preset's LABEL and nothing more: the fields
	// below carry what it expanded to. Storing the label as well as the
	// fields is what lets a saved sim say "Heavy movement" after the
	// preset's numbers have been retuned, and lets an edited request
	// stop claiming a preset it no longer matches.
	Style string `json:"style,omitempty"`

	// Movement schedules time out of melee or out of casting.
	Movement *Movement `json:"movement,omitempty"`

	// TargetsOverTime OVERRIDES Targets when set: a dungeon pull is a
	// boss and then packs, not a fixed count.
	TargetsOverTime []TargetCount `json:"targets_over_time,omitempty"`

	// TargetLevel is 60 to 63; 0 means BossLevel, which is what every
	// sim fought before this field existed.
	TargetLevel int `json:"target_level,omitempty"`

	// TargetArmor overrides the level's preset. nil (the field absent
	// from the request) means the preset; a non-nil pointer is an
	// explicit choice, INCLUDING one pointing at 0 - an unarmoured
	// target is a real request a player can make, not a second way to
	// spell "unset". A plain int could not tell "never sent" from
	// "sent as 0" apart on decode, which is how a blank settings field
	// and a typed 0 used to reach the engine as the identical request
	// and run identically (2026-09-21 result-page review, Defect 3).
	TargetArmor *int `json:"target_armor,omitempty"`

	// TargetType changes what Hunter and Warlock abilities do. "" maps
	// to the engine's own default, MobTypeHumanoid (sim/request's
	// encounter()), which is what every sim fought before this field
	// existed - keeping that is what stops the pin bump changing every
	// stored spec's number.
	//
	// Contract 10.8 settles the dummy case: a target dummy has no
	// creature type at all, so the "dummy" style sets TargetType to
	// TargetTypeUnknown, which maps to MobTypeUnknown and collects no
	// Hunter/Warlock creature-type bonus. An empty TargetType anywhere
	// else still maps to humanoid.
	TargetType string `json:"target_type,omitempty"`

	// Dummy is the training dummy: no debuffs, no execute window, no
	// armor reduction.
	Dummy bool `json:"dummy,omitempty"`
}

// Movement is a repeating window the player spends away from the target
// or unable to cast.
type Movement struct {
	IntervalSec int    `json:"interval_sec"`
	DurationSec int    `json:"duration_sec"`
	Kind        string `json:"kind"`
}

// The two kinds of movement window. Away is out of melee range with no
// casting; Casting interrupts spells while melee continues.
const (
	MovementAway    = "away"
	MovementCasting = "casting"
)

// MovementKinds is the closed set.
var MovementKinds = []string{MovementAway, MovementCasting}

// TargetCount is one step of a target-count timeline: from AtSec, this
// many targets are alive.
type TargetCount struct {
	AtSec int `json:"at_sec"`
	Count int `json:"count"`
}

// The target levels the settings bar offers.
const (
	MinTargetLevel = SimLevel
	MaxTargetLevel = BossLevel
)

// TargetArmorByLevel is the armor a target of each level carries when
// the request does not override it.
//
// The boss row is the engine's own preset
// (sim/encounters/default_presets.go); the three below it fall
// linearly to the level-60 figure. Contract A8 ratifies these four
// numbers and says a better source replaces the three interior ones.
// They are here rather than in sim/request because the page shows the
// preset beside the override control, and a second copy there would
// drift from the number the sim actually ran.
var TargetArmorByLevel = map[int]int{60: 3300, 61: 3444, 62: 3588, 63: 3731}

// TargetArmorFor resolves an encounter's armor: the override when it is
// set - nil or not, including a pointer at 0 - otherwise the level's
// preset, otherwise the boss's.
func TargetArmorFor(level int, override *int) int {
	if override != nil {
		return *override
	}
	if armor, ok := TargetArmorByLevel[level]; ok {
		return armor
	}
	return TargetArmorByLevel[BossLevel]
}

// TargetTypeUnknown is a target with no creature type, which is what a
// training dummy is and what every sim fought before this field.
const TargetTypeUnknown = "unknown"

// TargetTypes is the closed set, matching the engine's MobType enum.
// sim/request proves the pairing against the enum.
var TargetTypes = []string{
	"beast", "demon", "dragonkin", "elemental", "giant",
	"humanoid", "mechanical", "undead", TargetTypeUnknown,
}

// DefaultEncounter is the settings bar's opening state: a three-minute
// single-target fight with the standard duration variation and execute
// window.
func DefaultEncounter() EncounterSpec {
	return EncounterSpec{DurationSec: 180, Variation: 0.2, Targets: 1, ExecuteRatio: 0.25}
}

// Validate checks everything a malformed client could get wrong, at the
// boundary, before anything reaches the engine. It requires the request
// to name this build's own engine: a request is about to be run, and
// running one build's request against another's talents, item rows and
// spell constants would produce a row stamped with a version that did
// not produce it.
func (r SimRequest) Validate() error {
	return r.validate(true, true)
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
	return r.validate(false, true)
}

// validate is the body of Validate, ValidatePart and ValidateSaved.
// closedSet says whether the iteration count must be one the settings
// bar offers; requireCurrentEngine says whether engine_version must
// name this build - true for anything about to be run, false for a
// result being saved, which named whatever engine actually produced
// it.
func (r SimRequest) validate(closedSet, requireCurrentEngine bool) error {
	var errs []error
	switch {
	case r.EngineVersion == "":
		errs = append(errs, errors.New("engine_version is required"))
	case requireCurrentEngine && r.EngineVersion != enginever.Version:
		// One string identifies the engine build everywhere, and this
		// binary is one build. A request naming another cannot be run
		// here: the talents, the item rows and the spell constants all
		// belong to the pinned sha, so running it anyway would produce
		// a row stamped with a version that did not produce it. A
		// stored result from an older engine is still readable and is
		// labelled stale; re-running it is a new request, restamped by
		// whoever asks.
		errs = append(errs, fmt.Errorf("engine_version is %q, but this build is %q; a result is never produced by an engine other than the one it names", r.EngineVersion, enginever.Version))
	}
	if r.Spec == "" {
		errs = append(errs, errors.New("spec is required"))
	}
	switch {
	case r.Bulk != nil:
		// A bulk request's count is the precision's, checked by
		// BulkSpec.validate against the ladder.
	case !closedSet:
		// A split part comes first, ahead of the target-error case,
		// because a part of a target-error run is both: combine.Split
		// copies the request into each part, TargetError
		// included, since the loop owns the target and the part is
		// only a fixed-count share of one step. So a part's count is
		// neither one of the settings bar's numbers nor a whole
		// number of steps - a 1,000-iteration step split four ways is
		// 250 - and only the outer bound applies.
		if r.Iterations <= 0 || r.Iterations > MaxIterations {
			errs = append(errs, fmt.Errorf("a split part's iterations must be between 1 and %d, got %d", MaxIterations, r.Iterations))
		}
	case r.TargetError > 0:
		// A target-error run's Iterations is a ceiling, not a choice
		// from the settings bar, so the closed set does not apply and
		// neither does MaxIterations. It has to be a whole number of
		// steps, because a step is what the loop adds.
		errs = append(errs, validateCeiling(r.Iterations, largestCeiling())...)
	case !slices.Contains(ValidIterations, r.Iterations):
		errs = append(errs, fmt.Errorf("iterations must be one of %v, got %d", ValidIterations, r.Iterations))
	}
	if r.TargetError < 0 || r.TargetError >= 1 {
		errs = append(errs, fmt.Errorf("target_error is a fraction of the mean, so it must be between 0 and 1, got %v", r.TargetError))
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
	errs = append(errs, validateEncounterAdditions(r.Encounter)...)
	for i, cd := range r.Character.Cooldowns {
		if cd.ID == "" {
			errs = append(errs, fmt.Errorf("character.cooldowns[%d] has no id", i))
		}
		for j, at := range cd.AtSec {
			if at < 0 {
				errs = append(errs, fmt.Errorf("character.cooldowns[%d].at_sec[%d] is %v; a cooldown is used during the fight, and the pre-pull is the rotation's job", i, j, at))
			}
		}
	}
	if !slices.Contains(Sources, r.Source.Kind) {
		errs = append(errs, fmt.Errorf("source.kind must be one of %v, got %q", Sources, r.Source.Kind))
	}
	if err := validateProfile(r.Encounter.Profile); err != nil {
		errs = append(errs, err)
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
	if r.Bulk != nil {
		errs = append(errs, r.Bulk.validate(r.Iterations)...)
	}
	if r.Weights != nil {
		errs = append(errs, r.Weights.validate()...)
	}
	if r.Bulk != nil && r.Weights != nil {
		// Kind() reads Bulk first, so a request carrying both would run
		// as a bulk and silently answer a different question from the
		// one the weights block asked.
		errs = append(errs, errors.New("a request is one kind: it carries bulk or weights, never both"))
	}
	if r.TargetError > 0 && (r.Bulk != nil || r.Weights != nil) {
		// The same reasoning: a bulk request's iteration counts are its
		// precision's ladder and a weights request's are the engine's
		// own, so neither loops towards a target. Accepting the field
		// and then dropping it would report an answer at a precision
		// nobody ran.
		errs = append(errs, errors.New("a request is one kind: a bulk or weights request runs its own iteration counts, so it carries no target_error"))
	}
	return errors.Join(errs...)
}

// validateProfile refuses a profile the sim would otherwise accept and
// then ignore. "" and "patchwerk" are the fight it builds; an
// "encounter:<id>" is not, and nothing else is in the vocabulary.
func validateProfile(profile string) error {
	switch {
	case profile == "" || profile == ProfilePatchwerk:
		return nil
	case strings.HasPrefix(profile, ProfileEncounterPrefix):
		if id := strings.TrimPrefix(profile, ProfileEncounterPrefix); id == "" {
			return fmt.Errorf("encounter.profile %q names no encounter", profile)
		}
		return fmt.Errorf("encounter.profile %q is not simulated yet: there is no encounter table, so the run would be a patchwerk reported under an encounter's name", profile)
	default:
		return fmt.Errorf("encounter.profile must be %q, %q or %q<id>, got %q", "", ProfilePatchwerk, ProfileEncounterPrefix, profile)
	}
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
	// Aborted says the run was stopped on request rather than failing.
	// It is not an error: the user pressed Stop, IterationsRun is
	// whatever completed, and the page says "stopped" rather than
	// "something went wrong". The engine reports an abort as an
	// ErrorOutcome with an EMPTY message, so a lane that only looked at
	// Error reported it as a corrupt result.
	Aborted bool `json:"aborted,omitempty"`
	// Combos is the ranked answer to a bulk request, best first.
	Combos []Combo `json:"combos,omitempty"`
	// Equipped is the base character at the FINAL stage, which is what
	// every Delta is measured against. It is a pointer so that "no
	// baseline" and "a baseline of zero" are different things.
	Equipped *Estimate `json:"equipped,omitempty"`
	// Stages is what the planner actually ran, so the page can say
	// "3 stages, 38 combinations" without recomputing the ladder.
	Stages []Stage `json:"stages,omitempty"`
	// Weights is the answer to a weights request.
	Weights []StatWeight `json:"weights,omitempty"`
	// Sample is one iteration's casts in order - the median-DPS one.
	// It is a table view of the cast log, not a guide, and the page
	// says so.
	Sample []SampleCast `json:"sample,omitempty"`
}

// Combo is one substitution set's result.
type Combo struct {
	Substitutions []Substitution `json:"substitutions"`
	DPS           Estimate       `json:"dps"`
	// Delta is against Equipped, paired at the same stage, WITH ITS OWN
	// ERROR - which is what lets the page say "within error" honestly
	// instead of ranking noise.
	Delta Estimate `json:"delta"`
	// Group is 0 for the leader's within-error group, then 1, 2, and so
	// on. Rows in one group are the same answer and the page ranks them
	// the same.
	Group int `json:"group"`
}

// The four things a combination can substitute.
const (
	SubstitutionItem     = "item"
	SubstitutionTalents  = "talents"
	SubstitutionSet      = "set"
	SubstitutionConsumes = "consumes"
)

// NoConsumablesLabel is a SubstitutionConsumes chip's Name when the
// alternative list it names is empty. An empty list is a real
// candidate to compare against - "no consumables" - not the absence of
// one, so its chip needs a label rather than joining zero ids to "".
const NoConsumablesLabel = "no consumables"

// Substitution is one change from the base character.
type Substitution struct {
	Kind string `json:"kind"`
	// Slot is where an item landed. A ring or a trinket says which of
	// the two slots it took, because that is the answer the player
	// acts on.
	Slot    string `json:"slot,omitempty"`
	ItemID  int    `json:"item_id,omitempty"`
	Enchant int    `json:"enchant,omitempty"`
	Suffix  int    `json:"suffix,omitempty"`
	// Name is the item's, the loadout's or the set's. An ITEM's name
	// is filled too (from simdb, by sim/bulk): the API composes a
	// headline - "+41 DPS from Vis'kag" - at save time and has no item
	// table of its own, and a stored row read back years later still
	// says what it was about.
	Name    string `json:"name,omitempty"`
	Talents string `json:"talents,omitempty"`
	Origin  string `json:"origin,omitempty"`
	// SourceName is copied from the candidate: the human name of where
	// it came from, which is what a Droptimizer headline reads.
	SourceName string `json:"source_name,omitempty"`
	// Consumes is the alternative consumable list this combination
	// ran, for a consumables substitution. Its Name is the same ids
	// joined by ", ", so the API's headline and a chip's tooltip read
	// the one field every other kind of substitution puts its label
	// in (contract 10.8).
	Consumes []string `json:"consumes,omitempty"`
}

// Stage is one rung of the ladder, as run.
type Stage struct {
	Iterations int `json:"iterations"`
	Combos     int `json:"combos"`
}

// SampleCast is one cast of the sample iteration.
type SampleCast struct {
	// AtMS is negative during the pre-pull, which the page renders as
	// its own section.
	AtMS int64 `json:"at_ms"`
	// Action is the summary's ACTION KEY - "spell:23881", "item:13503",
	// "other:attack" - and not a display name. The page resolves the
	// name with resolveActionName, exactly as it does for a cast row,
	// so the two tables name one action one way and a sample row needs
	// no second lookup table. An earlier draft carried a numeric
	// SpellID and a Name; contract A12 removed both.
	Action string `json:"action"`
	Target string `json:"target,omitempty"`
	// Resources are what the player held AFTER the cast: rage, energy,
	// mana, combo_points.
	Resources map[string]int `json:"resources,omitempty"`
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
	// The three below are a bulk run's, and absent for a plain one:
	// stage 1 of 3, 31 of 96 combinations. Stage numbers start at 1.
	Stage       int `json:"stage,omitempty"`
	CombosDone  int `json:"combos_done,omitempty"`
	CombosTotal int `json:"combos_total,omitempty"`
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

// ValidateSaved checks a browser result on its way into POST /v1/sims.
// Saving is not running: the result already exists, produced by
// whatever engine the member's browser was holding, and the contract's
// answer to a build behind the deployment's pin is to store it and
// read it back stale (Stale), never to refuse it - a member who ran a
// sim on a cached bundle from before the last deploy must not lose it
// to a 400. So this checks the request's shape the way Validate does,
// except which engine produced it, and refuses the two things a saved
// row cannot be: a partial run, or a failed one wearing a done result's
// shape. The browser only ever posts a result that finished; an
// aborted one belongs to the page's own history, not the server's, and
// a result carrying an Error is the same story the job's own Finish
// already tells apart - stored as-is it would read back as a confident
// "0 DPS" rather than the failure it is.
func (r SimResult) ValidateSaved() error {
	if r.Aborted {
		return errors.New("an aborted result cannot be saved")
	}
	if r.Error != "" {
		return errors.New("a result carrying an error cannot be saved")
	}
	return r.Request.validate(true, false)
}

// largestCeiling is the biggest target-error ceiling any lane allows.
// Validate uses it because the envelope carries no lane; ValidateLane
// narrows it.
func largestCeiling() int {
	return largestValue(LaneIterationCeiling)
}

// largestValue is the biggest value in a lane-keyed map of ints. Both
// largestCeiling here and the bulk cap check in bulk.go need the
// largest lane's number because Validate has no lane to look up; one
// helper keeps the two loops from drifting apart.
func largestValue(m map[string]int) int {
	out := 0
	for _, v := range m {
		out = max(out, v)
	}
	return out
}

// validateCeiling holds a target-error run's ceiling to whole steps and
// to a lane's limit.
func validateCeiling(iterations, ceiling int) []error {
	var errs []error
	if iterations <= 0 || iterations%StepIterations != 0 {
		errs = append(errs, fmt.Errorf("a target-error run's iterations is a ceiling and must be a positive multiple of %d, got %d", StepIterations, iterations))
	}
	if iterations > ceiling {
		errs = append(errs, fmt.Errorf("a target-error run's iterations is at most %d on this lane, got %d", ceiling, iterations))
	}
	return errs
}

// ValidateLane is Validate plus the two limits that belong to a lane
// rather than to the request: the combination cap and the target-error
// ceiling. The envelope carries no lane - a request is a question, and
// the same question can be asked of either - so Validate applies the
// largest lane's numbers and this applies one lane's. The api handler
// calls it with LaneServer and the page with LaneBrowser.
func (r SimRequest) ValidateLane(lane string) error {
	var errs []error
	if err := r.Validate(); err != nil {
		errs = append(errs, err)
	}
	capped, ok := Caps[lane]
	ceiling := LaneIterationCeiling[lane]
	if !ok {
		return errors.Join(append(errs, fmt.Errorf("lane must be %q or %q, got %q", LaneBrowser, LaneServer, lane))...)
	}
	if r.Bulk != nil && r.Bulk.Cap > capped {
		errs = append(errs, fmt.Errorf("the %s lane plans at most %d combinations, and bulk.cap is %d", lane, capped, r.Bulk.Cap))
	}
	if r.TargetError > 0 && r.Iterations > ceiling {
		errs = append(errs, fmt.Errorf("the %s lane runs at most %d iterations, and iterations is %d", lane, ceiling, r.Iterations))
	}
	return errors.Join(errs...)
}

// NeedsMoreIterations reports whether a target-error run should run
// another step.
//
// It lives here, in Go, because BOTH lanes loop and they must agree.
// The browser runs each step as an ordinary sharded run through
// simSplit/simRun/simCombine and asks this between steps; forever-sim
// runs the step itself and asks the same function. A copy of this
// arithmetic in TypeScript is exactly the drift the module exists to
// prevent - the two would stop at different precisions and the same
// request would report a different error bar depending on where it ran.
//
// A run that failed or was stopped never steps: there is nothing to
// refine, and stepping would turn one bad answer into several. Neither
// does a bulk run: its counts are its precision's ladder, its
// Iterations is the last stage's and not a ceiling, and Validate
// refuses the pairing outright - this is the same ruling stated where
// the loop would otherwise act on it.
func NeedsMoreIterations(res SimResult, req SimRequest) bool {
	if req.TargetError <= 0 || req.Bulk != nil || res.Error != "" || res.Aborted {
		return false
	}
	if res.IterationsRun >= req.Iterations {
		return false
	}
	if res.DPS.Mean <= 0 {
		// No mean yet, so no relative error either. One more step is
		// the only way to learn anything, and the ceiling above is what
		// stops it being forever.
		return true
	}
	return res.DPS.Error/res.DPS.Mean > req.TargetError
}

// NextStepIterations is how many iterations the next step runs, and 0
// when there is no next step. The last step is shortened rather than
// overshooting, so a ceiling of 3,000 is 3,000 and never 3,500.
func NextStepIterations(res SimResult, req SimRequest) int {
	if !NeedsMoreIterations(res, req) {
		return 0
	}
	return min(StepIterations, req.Iterations-res.IterationsRun)
}

// validateEncounterAdditions checks the fields the parity contract added
// to EncounterSpec. They are all optional, so every check is on a value
// the client actually sent.
func validateEncounterAdditions(e EncounterSpec) []error {
	var errs []error
	if m := e.Movement; m != nil {
		if !slices.Contains(MovementKinds, m.Kind) {
			errs = append(errs, fmt.Errorf("encounter.movement.kind must be one of %v, got %q", MovementKinds, m.Kind))
		}
		if m.IntervalSec <= 0 {
			errs = append(errs, fmt.Errorf("encounter.movement.interval_sec must be positive, got %d", m.IntervalSec))
		}
		if m.DurationSec <= 0 {
			errs = append(errs, fmt.Errorf("encounter.movement.duration_sec must be positive, got %d", m.DurationSec))
		}
		if m.IntervalSec > 0 && m.DurationSec >= m.IntervalSec {
			errs = append(errs, fmt.Errorf("encounter.movement.duration_sec (%d) must be shorter than the interval (%d), or the player never stands still", m.DurationSec, m.IntervalSec))
		}
	}
	if len(e.TargetsOverTime) > 0 {
		if e.TargetsOverTime[0].AtSec != 0 {
			errs = append(errs, fmt.Errorf("encounter.targets_over_time starts at 0, not %d; the fight has targets from the pull", e.TargetsOverTime[0].AtSec))
		}
		prev := -1
		for i, tc := range e.TargetsOverTime {
			if tc.AtSec <= prev {
				errs = append(errs, fmt.Errorf("encounter.targets_over_time must be in time order; entry %d is at %d after %d", i, tc.AtSec, prev))
			}
			prev = tc.AtSec
			if tc.Count < 1 || tc.Count > MaxTargets {
				errs = append(errs, fmt.Errorf("encounter.targets_over_time[%d].count must be between 1 and %d, got %d", i, MaxTargets, tc.Count))
			}
		}
	}
	if e.TargetLevel != 0 && (e.TargetLevel < MinTargetLevel || e.TargetLevel > MaxTargetLevel) {
		errs = append(errs, fmt.Errorf("encounter.target_level must be between %d and %d, got %d", MinTargetLevel, MaxTargetLevel, e.TargetLevel))
	}
	if e.TargetArmor != nil && *e.TargetArmor < 0 {
		errs = append(errs, fmt.Errorf("encounter.target_armor must not be negative, got %d; omit the field for the level's preset, or send 0 for no armor", *e.TargetArmor))
	}
	if e.TargetType != "" && !slices.Contains(TargetTypes, e.TargetType) {
		errs = append(errs, fmt.Errorf("encounter.target_type must be one of %v, got %q", TargetTypes, e.TargetType))
	}
	return errs
}
