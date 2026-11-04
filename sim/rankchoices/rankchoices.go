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
// applies to both. A row naming Races (sim/request race slugs) applies only to a
// character of one of them; one with none applies to every race. Races are keyed only
// where the races' best ranks differ beyond error (owner ruling 3, sim-accuracy round 13),
// and a race named by no row of a band casts the default rule there.
type Row struct {
	FromLevel int      `json:"from_level"`
	ToLevel   int      `json:"to_level"`
	Rank      int      `json:"rank"`
	Preset    string   `json:"preset,omitempty"`
	Races     []string `json:"races,omitempty"`
	// Note is the row's evidence for a reader. A race row says what each race
	// measured and which racial drives the split.
	Note string `json:"note,omitempty"`
	// Lines and Prepull are the Spell's, copied onto each row by sim/request's
	// parse; the file writes them once per spell.
	Lines   []int `json:"-"`
	Prepull []int `json:"-"`
}

// Spell is one authored spell id's rows, as the curated file writes them.
type Spell struct {
	// Spell is the spellranks.json name of Authored's spell, for a reader.
	Spell    string `json:"spell"`
	Authored int32  `json:"authored"`
	// Lines and Prepull, when either is present, scope the rows: they apply
	// only to those 0-based priority-list entries and pre-pull actions, and
	// every other cast of Authored keeps the default rule. Restoration's
	// Nature's Swiftness save and its tank filler cast the same Healing Wave id,
	// and only the filler downranks; Elemental's opening bolt follows its
	// filler, not its top-rank line.
	Lines   []int `json:"lines,omitempty"`
	Prepull []int `json:"prepull,omitempty"`
	Bands   []Row `json:"bands"`
}

// File is the table's shape.
type File struct {
	// Notes is the table's own prose for a reader.
	Notes string             `json:"notes"`
	Specs map[string][]Spell `json:"specs"`
}
