package api

// Stat weights.
//
// The engine already computes them, so this is a request shape and a
// result shape, not arithmetic. The caution belongs to the page: a
// weight is a linear guess at a non-linear thing, and simming the
// actual items is the better answer. The envelope's job is to make the
// request honest - a reference stat that is not among the stats being
// weighed would normalise against a number nobody computed.

import (
	"errors"
	"fmt"
)

// WeightsSpec asks for stat weights instead of a DPS number.
type WeightsSpec struct {
	// Stats are IDS.md stat ids: the engine's Stat enum names in lower
	// snake case with the prefix stripped, "attack_power", "crit",
	// "spell_haste". sim/request resolves them; an id the engine has no
	// stat for is refused there rather than weighed as nothing.
	Stats []string `json:"stats"`
	// Reference is the stat normalised to exactly 1.0. It is
	// REQUIRED (contract A7): its per-spec default lives in
	// data/curated/specs.json, reaches Go as the generated
	// sim/specs struct's ReferenceStat, and is served to the page on
	// GET /v1/specs. The envelope does not read it - defaulting here
	// would make the envelope a second source of truth for a fact
	// the data lane owns, and a request that did not say which stat
	// it normalised against would be unreproducible.
	Reference string `json:"reference"`
}

// StatWeight is one stat's answer. Reference's Weight is exactly 1.
type StatWeight struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
	Error  float64 `json:"error"`
}

func (w *WeightsSpec) validate() []error {
	var errs []error
	if len(w.Stats) == 0 {
		errs = append(errs, errors.New("weights needs at least one stat to weigh"))
	}
	seen := map[string]bool{}
	for _, s := range w.Stats {
		if s == "" {
			errs = append(errs, errors.New("weights.stats carries an empty stat id"))
			continue
		}
		if seen[s] {
			errs = append(errs, fmt.Errorf("weights.stats has %q listed twice", s))
		}
		seen[s] = true
	}
	switch {
	case w.Reference == "":
		errs = append(errs, errors.New("weights.reference is required; the spec's default is in data/curated/specs.json"))
	case !seen[w.Reference]:
		errs = append(errs, fmt.Errorf("weights.reference is %q, which is not one of the stats it weighs (%v); the reference is normalised to 1 and there would be nothing to normalise", w.Reference, w.Stats))
	}
	return errs
}
