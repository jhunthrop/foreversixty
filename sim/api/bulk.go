package api

// The bulk block: one request shape for every combination tool.
//
// Top Gear, Droptimizer and talent compare ask the same question - a
// base character, a set of substitutions, a ranked answer - so they are
// one kind of request with one planner behind it (sim/bulk) rather than
// three. Mode says which of the three asked; it is also the request's
// Kind, which is why the three mode strings ARE the three kind strings
// and not a parallel vocabulary that could drift from them.

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// The five kinds of request. A kind is derived from the request, never
// sent: a client that could name its own kind could name one the
// request's own shape contradicts, and every row keyed on it would be
// wrong in a way nothing could detect afterwards.
const (
	KindRun     = "run"
	KindGear    = "gear"
	KindTalents = "talents"
	KindDrops   = "drops"
	KindWeights = "weights"
)

// BulkModes is the closed set BulkSpec.Mode is checked against. They are
// the three bulk kinds, in the order the tools were built.
var BulkModes = []string{KindGear, KindTalents, KindDrops}

// Kind reports which tool produced this request.
//
// Task 4 adds the weights arm; until then a request with no bulk
// block is a run.
func (r SimRequest) Kind() string {
	if r.Bulk != nil {
		return r.Bulk.Mode
	}
	return KindRun
}

// BulkSpec is everything that turns one character into many sims.
type BulkSpec struct {
	Mode       string          `json:"mode"`
	Candidates []Candidate     `json:"candidates"`
	Talents    []TalentLoadout `json:"talents,omitempty"`
	Sets       []GearSet       `json:"sets,omitempty"`
	// Consumables are alternative consumable lists tried as candidates
	// in gear mode: each inner list REPLACES CharacterSpec.Consumes for
	// that combination, so "flask or elixirs" is a dimension of the
	// product rather than a setting the player has to run twice.
	Consumables [][]string `json:"consumables,omitempty"`
	// Locked names slots that are never substituted, whatever the
	// candidate list says. It is the page's "keep what I have here".
	Locked    []string `json:"locked,omitempty"`
	Precision string   `json:"precision"`
	// Cap is the lane's combination cap, echoed into the request so a
	// saved row says what bounded it. The planner refuses an expansion
	// above it; it never silently trims.
	Cap int `json:"cap"`
}

// Candidate is one item the planner may substitute in.
type Candidate struct {
	// Slot is the planner slot name, or "" for "wherever it fits",
	// which is how a ring, a trinket or a one-hander is offered.
	Slot   string `json:"slot"`
	ItemID int    `json:"item_id"`
	// Enchant of 0 inherits the equipped item's enchant for the slot it
	// lands in, where that enchant fits; it does not mean "none".
	Enchant int    `json:"enchant,omitempty"`
	Suffix  int    `json:"suffix,omitempty"`
	Origin  string `json:"origin"`
	// SourceName is the human name of the place this candidate came
	// from - "Lucifron", "my AQ set". The page fills it from
	// loot.json when it builds a drops request, the substitution
	// carries it through, and the API's headline reads it, so nothing
	// downstream has to join an origin id back to a name.
	SourceName string `json:"source_name,omitempty"`
}

// TalentLoadout is one named build to try.
type TalentLoadout struct {
	Name    string `json:"name"`
	Talents string `json:"talents"`
}

// GearSet is a whole-gear alternative: one candidate that replaces every
// slot at once. It is what Raidbots' legacy Gear Compare was for.
type GearSet struct {
	Name string     `json:"name"`
	Gear []GearSlot `json:"gear"`
}

// Where a candidate came from. The page shows it, history reads it, and
// Droptimizer groups its results by the source id a drop origin names.
const (
	OriginEquipped   = "equipped"
	OriginBag        = "bag"
	OriginBank       = "bank"
	OriginSearch     = "search"
	OriginDropPrefix = "drop:"
	OriginSetPrefix  = "set:"
)

// plainOrigins are the origins that are a whole word.
var plainOrigins = []string{OriginEquipped, OriginBag, OriginBank, OriginSearch}

// The three precisions. They differ in how many stages the planner runs
// and how many iterations the last one gets.
const (
	PrecisionFast   = "fast"
	PrecisionNormal = "normal"
	PrecisionHigh   = "high"
)

// Precisions is the closed set, weakest first.
var Precisions = []string{PrecisionFast, PrecisionNormal, PrecisionHigh}

// Cut is what survives one stage.
//
// Fraction and Top are alternatives: exactly one is non-zero. SlackSE
// widens either by keeping anything whose interval still overlaps the
// last survivor's, which is what stops a stage of 100 iterations from
// throwing away the eventual winner over noise.
type Cut struct {
	Fraction float64 `json:"fraction,omitempty"`
	Top      int     `json:"top,omitempty"`
	SlackSE  float64 `json:"slack_se"`
}

// Ladder is a precision's staging plan: the iteration count of each
// stage and the cut between each pair of them.
type Ladder struct {
	Iterations []int `json:"iterations"`
	Cuts       []Cut `json:"cuts"`
}

// Ladders is the fork's fast_mode ladder made explicit, so the browser
// and the server agree on it rather than each guessing. The equipped set
// runs in every stage, so every delta is paired.
var Ladders = map[string]Ladder{
	PrecisionFast: {
		Iterations: []int{100, 1000, 3000},
		Cuts:       []Cut{{Fraction: 0.25, SlackSE: 2}, {Top: 10, SlackSE: 2}},
	},
	PrecisionNormal: {
		Iterations: []int{1000, 3000},
		Cuts:       []Cut{{Top: 10, SlackSE: 2}},
	},
	PrecisionHigh: {
		Iterations: []int{1000, 10000},
		Cuts:       []Cut{{Top: 20, SlackSE: 2}},
	},
}

// FinalIterations is the iteration count a precision's last stage runs
// at, which is what a bulk request's Iterations field must say.
func FinalIterations(precision string) (int, bool) {
	l, ok := Ladders[precision]
	if !ok || len(l.Iterations) == 0 {
		return 0, false
	}
	return l.Iterations[len(l.Iterations)-1], true
}

// Caps is how many combinations each lane will plan.
//
// The browser's number is 400 because 400 combinations at 100
// iterations is about a minute on a mid laptop with eight workers.
//
// The server's is 5,000, not the 20,000 an earlier draft carried:
// sim/measure benchmarks the native engine at
// measure.NativeIterationsPerCPUSecond iterations per CPU-second, and
// a 20,000-combination fast run does not finish inside the Cloud Run
// job's 15-minute timeout at that rate. The timeout stays; the cap
// moved. The api lane's "too_large" estimate is built from the same
// constant - NativeIterationsPerCPUSecond x 4 CPUs x 840 seconds - so
// the cap and the estimate cannot disagree.
var Caps = map[string]int{LaneBrowser: 400, LaneServer: 5000}

// ErrCapExceeded is returned when an expansion is larger than the lane
// allows. It carries both numbers because the page and the API both say
// them out loud - "31,200 combinations, and this lane plans 5,000" -
// rather than trimming the list and reporting a ranking of a subset
// nobody chose.
type ErrCapExceeded struct {
	Cap          int `json:"cap"`
	Combinations int `json:"combinations"`
}

func (e ErrCapExceeded) Error() string {
	return fmt.Sprintf("the plan is %d combinations, and this lane plans at most %d", e.Combinations, e.Cap)
}

// validateBulk checks everything about the bulk block that a malformed
// client could get wrong, against the largest lane's cap; ValidateLane
// narrows it to one lane.
func (b *BulkSpec) validate(iterations int) []error {
	var errs []error
	if !slices.Contains(BulkModes, b.Mode) {
		errs = append(errs, fmt.Errorf("bulk.mode must be one of %v, got %q", BulkModes, b.Mode))
	}
	if !slices.Contains(Precisions, b.Precision) {
		errs = append(errs, fmt.Errorf("bulk.precision must be one of %v, got %q", Precisions, b.Precision))
	}
	if final, ok := FinalIterations(b.Precision); ok && iterations != final {
		errs = append(errs, fmt.Errorf("a %s bulk request runs its last stage at %d, so iterations must be %d, got %d", b.Precision, final, final, iterations))
	}
	largest := 0
	for _, c := range Caps {
		largest = max(largest, c)
	}
	if b.Cap <= 0 || b.Cap > largest {
		errs = append(errs, fmt.Errorf("bulk.cap must be between 1 and %d, got %d", largest, b.Cap))
	}

	locked := map[string]bool{}
	for _, slot := range b.Locked {
		if !slices.Contains(GearSlots, slot) {
			errs = append(errs, fmt.Errorf("bulk.locked names %q, which is not a gear slot", slot))
			continue
		}
		locked[slot] = true
	}

	for i, c := range b.Candidates {
		if c.ItemID <= 0 {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d] has no item_id", i))
		}
		if c.Slot != "" && !slices.Contains(GearSlots, c.Slot) {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d] names slot %q, which is not a gear slot", i, c.Slot))
		}
		if c.Slot != "" && locked[c.Slot] {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d] is on %q, which is locked", i, c.Slot))
		}
		if err := validateOrigin(c.Origin); err != nil {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d]: %w", i, err))
		}
	}

	for i, l := range b.Talents {
		if l.Name == "" || l.Talents == "" {
			errs = append(errs, fmt.Errorf("bulk.talents[%d] is a talent loadout with no name or no talents", i))
		}
	}
	for i, s := range b.Sets {
		if s.Name == "" || len(s.Gear) == 0 {
			errs = append(errs, fmt.Errorf("bulk.sets[%d] is a gear set with no name or no gear", i))
		}
	}
	for i, list := range b.Consumables {
		// An empty inner list is legal and means "no consumables",
		// which is a real thing to compare against. A list of empty
		// ids is not.
		for j, id := range list {
			if id == "" {
				errs = append(errs, fmt.Errorf("bulk.consumables[%d][%d] is an empty consumable id", i, j))
			}
		}
	}
	if len(b.Consumables) > 0 && b.Mode != KindGear {
		errs = append(errs, fmt.Errorf("bulk.consumables is a gear-mode dimension; mode is %q", b.Mode))
	}

	switch b.Mode {
	case KindGear:
		if len(b.Candidates) == 0 && len(b.Talents) == 0 && len(b.Sets) == 0 && len(b.Consumables) == 0 {
			errs = append(errs, errors.New("a gear request needs at least one candidate, talent loadout, set or consumable list"))
		}
	case KindTalents:
		if len(b.Talents) == 0 {
			errs = append(errs, errors.New("a talents request needs at least one talent loadout"))
		}
		if len(b.Candidates) != 0 || len(b.Sets) != 0 || len(b.Consumables) != 0 {
			errs = append(errs, errors.New("a talents request carries no candidates, sets or consumable lists; everything but the talents is locked"))
		}
	case KindDrops:
		for i, c := range b.Candidates {
			if !strings.HasPrefix(c.Origin, OriginDropPrefix) {
				errs = append(errs, fmt.Errorf("a drops request needs every candidate to come from a source; bulk.candidates[%d] is %q", i, c.Origin))
			}
		}
		if len(b.Candidates) == 0 {
			errs = append(errs, errors.New("a drops request needs at least one candidate"))
		}
	}
	return errs
}

// validateOrigin holds an origin to the vocabulary. A prefixed origin
// must name something after the colon: "drop:" alone would group a
// Droptimizer result under a source that does not exist.
func validateOrigin(origin string) error {
	if slices.Contains(plainOrigins, origin) {
		return nil
	}
	for _, prefix := range []string{OriginDropPrefix, OriginSetPrefix} {
		if rest, ok := strings.CutPrefix(origin, prefix); ok {
			if rest == "" {
				return fmt.Errorf("origin %q names nothing after the prefix", origin)
			}
			return nil
		}
	}
	return fmt.Errorf("origin must be one of %v, %q<id> or %q<name>, got %q",
		plainOrigins, OriginDropPrefix, OriginSetPrefix, origin)
}
