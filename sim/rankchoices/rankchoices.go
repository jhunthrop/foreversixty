// Package rankchoices holds the per-band rank-choice table (data/curated/rank-choices.json,
// copied here by `make rank-choices-sync`) with no engine import, so code outside the sim,
// such as the log coach in api/, can read it without linking the engine. sim/request
// rewrites rotations from the same bytes.
package rankchoices

import _ "embed"

// JSON is the embedded copy of data/curated/rank-choices.json; sim/request's
// TestEmbeddedRankChoicesMatchCurated fails when it is stale.
//
//go:embed rank-choices.json
var JSON []byte

// Row is one band's choice: from FromLevel through ToLevel the spell casts rank Rank.
// Outside every row the default rule stands. A row naming a Preset (sim/request's
// BarePreset or RaidPreset) applies only to a request simmed under it; one with none
// applies to both.
type Row struct {
	FromLevel int    `json:"from_level"`
	ToLevel   int    `json:"to_level"`
	Rank      int    `json:"rank"`
	Preset    string `json:"preset,omitempty"`
}

// Spell is one authored spell id's rows, as the curated file writes them.
type Spell struct {
	// Spell is the spellranks.json name of Authored's spell, for a reader.
	Spell    string `json:"spell"`
	Authored int32  `json:"authored"`
	Bands    []Row  `json:"bands"`
}

// File is the table's shape.
type File struct {
	Specs map[string][]Spell `json:"specs"`
}
