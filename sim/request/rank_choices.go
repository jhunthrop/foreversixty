package request

// Per-band rank choices: which rank of a ranked spell a spec's rotation
// casts over a level band below 60 when the default rule (the highest rank
// learned, never above the authored one) is not what a player would cast.
// Owner ruling 3, sim-accuracy round 9: a mana-bound fight downranks.
//
// data/curated/rank-choices.json is the single source; `make
// rank-choices-sync` copies it here, and TestEmbeddedRankChoicesMatchCurated
// fails when the copy is stale. The table is parsed once per process, so the
// rewrite adds a map lookup per ranked id and nothing else to a /sim run.

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/api"
)

//go:embed rank-choices.json
var rankChoicesJSON []byte

// RankChoiceRow is one band's choice: from FromLevel through ToLevel the
// spell casts rank Rank. Outside every row the default rule stands. A row
// naming a Preset (BarePreset or RaidPreset) applies only to a request
// simmed under it; one with none applies to both.
type RankChoiceRow struct {
	FromLevel int    `json:"from_level"`
	ToLevel   int    `json:"to_level"`
	Rank      int    `json:"rank"`
	Preset    string `json:"preset,omitempty"`
}

// rankChoiceSpell is one authored spell id's rows, as the curated file
// writes them.
type rankChoiceSpell struct {
	// Spell is the spellranks.json name of Authored's spell, for a reader.
	Spell    string          `json:"spell"`
	Authored int32           `json:"authored"`
	Bands    []RankChoiceRow `json:"bands"`
}

type rankChoicesFile struct {
	Specs map[string][]rankChoiceSpell `json:"specs"`
}

// rankChoices is one spec's rows by authored id, each list ascending by
// FromLevel.
type rankChoices map[int32][]RankChoiceRow

// forPreset is c with only the rows that apply under preset.
func (c rankChoices) forPreset(preset string) rankChoices {
	out := make(rankChoices, len(c))
	for id, rows := range c {
		for _, row := range rows {
			if row.Preset == "" || row.Preset == preset {
				out[id] = append(out[id], row)
			}
		}
	}
	return out
}

// rankAt is the rank the row holding level chooses for id, if any.
func (c rankChoices) rankAt(id int32, level int) (int, bool) {
	for _, row := range c[id] {
		if row.FromLevel <= level && level <= row.ToLevel {
			return row.Rank, true
		}
	}
	return 0, false
}

// parseRankChoices decodes the curated table, sorting each spell's rows and
// refusing a duplicate authored id, an empty row or two rows that overlap.
func parseRankChoices(b []byte) (map[string]rankChoices, error) {
	var file rankChoicesFile
	if err := json.Unmarshal(b, &file); err != nil {
		return nil, fmt.Errorf("rank choices: the table is not readable: %w", err)
	}
	out := make(map[string]rankChoices, len(file.Specs))
	for spec, spells := range file.Specs {
		bySpell := rankChoices{}
		for _, sp := range spells {
			if _, dup := bySpell[sp.Authored]; dup {
				return nil, fmt.Errorf("rank choices: %s names %d twice", spec, sp.Authored)
			}
			rows := slices.Clone(sp.Bands)
			slices.SortFunc(rows, func(a, b RankChoiceRow) int { return a.FromLevel - b.FromLevel })
			for i, row := range rows {
				if row.Preset != "" && row.Preset != BarePreset && row.Preset != RaidPreset {
					return nil, fmt.Errorf("%w: rank choices: %s %d names preset %q", ErrUnknownPreset, spec, sp.Authored, row.Preset)
				}
				if row.ToLevel < row.FromLevel {
					return nil, fmt.Errorf("rank choices: %s %d has a row from level %d to %d", spec, sp.Authored, row.FromLevel, row.ToLevel)
				}
				if i > 0 && row.FromLevel <= rows[i-1].ToLevel && overlapPresets(row, rows[i-1]) {
					return nil, fmt.Errorf("rank choices: %s %d has two rows at level %d", spec, sp.Authored, row.FromLevel)
				}
			}
			bySpell[sp.Authored] = rows
		}
		out[spec] = bySpell
	}
	return out, nil
}

// overlapPresets reports whether two rows can apply to the same request.
func overlapPresets(a, b RankChoiceRow) bool {
	return a.Preset == "" || b.Preset == "" || a.Preset == b.Preset
}

var embeddedRankChoices = sync.OnceValues(func() (map[string]rankChoices, error) {
	return parseRankChoices(rankChoicesJSON)
})

// specRankChoices is spec's rows under preset from the embedded table; none
// when the spec has no row (or the embed is unreadable, which
// TestEmbeddedRankChoicesMatchCurated fails on).
func specRankChoices(spec, preset string) rankChoices {
	all, err := embeddedRankChoices()
	if err != nil {
		return nil
	}
	return all[spec].forPreset(preset)
}

// rankPreset is the preset a request's rank choices follow: the one it
// names, else the one its level is published under (PublishedPreset), which
// is also what the /sim page defaults to.
func rankPreset(req api.SimRequest) (string, error) {
	switch req.Preset {
	case "":
		return PublishedPreset(req.Character.Level), nil
	case BarePreset, RaidPreset:
		return req.Preset, nil
	}
	return "", fmt.Errorf("%w: %q; a request's preset is %q or %q", ErrUnknownPreset, req.Preset, BarePreset, RaidPreset)
}
