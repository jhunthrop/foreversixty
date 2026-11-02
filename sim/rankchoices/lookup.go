package rankchoices

// A read of the per-band rank choices for code outside the rotation: the log coach
// (api/internal/coach F5) must not call a lower rank a mistake where the site's own
// rotation casts it. sim/request rewrites rotations from the same embedded table.

import (
	"encoding/json"
	"fmt"
	"slices"
	"sync"
)

// rankChoicesByName is spec -> spellranks.json spell name -> that spell's rows.
var rankChoicesByName = sync.OnceValues(func() (map[string]map[string][]Row, error) {
	var file File
	if err := json.Unmarshal(JSON, &file); err != nil {
		return nil, fmt.Errorf("rank choices: the table is not readable: %w", err)
	}
	out := make(map[string]map[string][]Row, len(file.Specs))
	for spec, spells := range file.Specs {
		byName := make(map[string][]Row, len(spells))
		for _, sp := range spells {
			byName[sp.Spell] = append(byName[sp.Spell], sp.Bands...)
		}
		out[spec] = byName
	}
	return out, nil
})

// ChosenRanks is every rank data/curated/rank-choices.json has spec's rotation cast of the
// spell named name (its spellranks.json name) at level, under either preset, ascending and
// without repeats: a log does not say which preset the player raided under. Nil means no
// row holds level and the default rule (the highest rank learned) stands.
func ChosenRanks(spec, name string, level int) ([]int, error) {
	table, err := rankChoicesByName()
	if err != nil {
		return nil, err
	}
	var out []int
	for _, row := range table[spec][name] {
		if row.FromLevel <= level && level <= row.ToLevel && !slices.Contains(out, row.Rank) {
			out = append(out, row.Rank)
		}
	}
	slices.Sort(out)
	return out, nil
}
