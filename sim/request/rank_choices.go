package request

// Per-band rank choices: which rank of a ranked spell a spec's rotation
// casts over a level band below 60 when the default rule (the highest rank
// learned, never above the authored one) is not what a player would cast.
// Owner ruling 3, sim-accuracy round 9: a mana-bound fight downranks.
//
// data/curated/rank-choices.json is the single source; `make
// rank-choices-sync` copies it to sim/rankchoices, which embeds it, and TestEmbeddedRankChoicesMatchCurated
// fails when the copy is stale. The table is parsed once per process, so the
// rewrite adds a map lookup per ranked id and nothing else to a /sim run.

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/rankchoices"
)

// rankChoicesJSON is the table's embedded bytes (sim/rankchoices).
var rankChoicesJSON = rankchoices.JSON

// RankChoiceRow is one band's choice; see rankchoices.Row.
type RankChoiceRow = rankchoices.Row

type rankChoiceSpell = rankchoices.Spell

type rankChoicesFile = rankchoices.File

// rankChoices is one spec's rows by authored id, each list ascending by
// FromLevel.
type rankChoices map[int32][]RankChoiceRow

// forRequest is c with only the rows that apply to a request simmed under
// preset for a character of race (a slug of races).
func (c rankChoices) forRequest(preset, race string) rankChoices {
	out := make(rankChoices, len(c))
	for id, rows := range c {
		for _, row := range rows {
			if appliesTo(row.Preset, preset) && racesInclude(row.Races, race) {
				out[id] = append(out[id], row)
			}
		}
	}
	return out
}

// appliesTo reports whether a row keyed key ("" for every value) applies to value.
func appliesTo(key, value string) bool { return key == "" || key == value }

// racesInclude reports whether a row keyed to races (none for every race)
// applies to race.
func racesInclude(races []string, race string) bool {
	return len(races) == 0 || slices.Contains(races, race)
}

// noLine is the line of a cast outside the priority list (a pre-pull action):
// no line-scoped row applies to it.
const noLine = -1

// rankAt is the rank the row holding level chooses for id on a line its rows
// do not scope, if any.
func (c rankChoices) rankAt(id int32, level int) (int, bool) {
	return c.rankAtLine(id, level, noLine)
}

// rankAtLine is the rank the row holding level chooses for id cast on
// priority-list entry line, if any: a row scoped to lines applies only there.
func (c rankChoices) rankAtLine(id int32, level, line int) (int, bool) {
	for _, row := range c[id] {
		if row.FromLevel <= level && level <= row.ToLevel && (len(row.Lines) == 0 || slices.Contains(row.Lines, line)) {
			return row.Rank, true
		}
	}
	return 0, false
}

// scopedLines is every line a row for id is scoped to, ascending.
func (c rankChoices) scopedLines(id int32) []int {
	var out []int
	for _, row := range c[id] {
		for _, line := range row.Lines {
			if !slices.Contains(out, line) {
				out = append(out, line)
			}
		}
	}
	slices.Sort(out)
	return out
}

// checkLines refuses a spell's line scope unless it is absent or names
// distinct non-negative priority-list entries.
func checkLines(lines []int) error {
	if lines != nil && len(lines) == 0 {
		return errors.New("lines name no line; leave the key out for every line")
	}
	for i, line := range lines {
		if line < 0 {
			return fmt.Errorf("line %d is not a priority-list entry", line)
		}
		if slices.Contains(lines[:i], line) {
			return fmt.Errorf("line %d is named twice", line)
		}
	}
	return nil
}

// parseRankChoices decodes the curated table, sorting each spell's rows and
// refusing a duplicate authored id, an empty row or two rows that overlap.
func parseRankChoices(b []byte) (map[string]rankChoices, error) {
	var file rankChoicesFile
	dec := json.NewDecoder(bytes.NewReader(b))
	// Strict: a key the table does not define (the deleted "faction" key, a
	// typo of "races") fails here rather than silently applying a row to
	// every race.
	dec.DisallowUnknownFields()
	if err := dec.Decode(&file); err != nil {
		return nil, fmt.Errorf("rank choices: the table is not readable: %w", err)
	}
	out := make(map[string]rankChoices, len(file.Specs))
	for spec, spells := range file.Specs {
		bySpell := rankChoices{}
		for _, sp := range spells {
			if _, dup := bySpell[sp.Authored]; dup {
				return nil, fmt.Errorf("rank choices: %s names %d twice", spec, sp.Authored)
			}
			if err := checkLines(sp.Lines); err != nil {
				return nil, fmt.Errorf("rank choices: %s %d: %w", spec, sp.Authored, err)
			}
			rows := slices.Clone(sp.Bands)
			for i := range rows {
				rows[i].Lines = sp.Lines
			}
			slices.SortFunc(rows, func(a, b RankChoiceRow) int { return a.FromLevel - b.FromLevel })
			for i, row := range rows {
				if row.Preset != "" && row.Preset != BarePreset && row.Preset != RaidPreset {
					return nil, fmt.Errorf("%w: rank choices: %s %d names preset %q", ErrUnknownPreset, spec, sp.Authored, row.Preset)
				}
				if err := checkRowRaces(row.Races); err != nil {
					return nil, fmt.Errorf("rank choices: %s %d from level %d: %w", spec, sp.Authored, row.FromLevel, err)
				}
				if row.ToLevel < row.FromLevel {
					return nil, fmt.Errorf("rank choices: %s %d has a row from level %d to %d", spec, sp.Authored, row.FromLevel, row.ToLevel)
				}
				for _, prev := range rows[:i] {
					if row.FromLevel <= prev.ToLevel && overlap(row, prev) {
						return nil, fmt.Errorf("rank choices: %s %d has two rows at level %d", spec, sp.Authored, row.FromLevel)
					}
				}
			}
			bySpell[sp.Authored] = rows
		}
		out[spec] = bySpell
	}
	return out, nil
}

// checkRowRaces refuses a row's race key unless it is absent (every race) or
// names known race slugs, each once.
func checkRowRaces(rowRaces []string) error {
	if rowRaces != nil && len(rowRaces) == 0 {
		return errors.New("a row's races name no race; leave the key out for every race")
	}
	for i, race := range rowRaces {
		if _, ok := races[race]; !ok {
			return fmt.Errorf("%w: %q", ErrUnknownRace, race)
		}
		if slices.Contains(rowRaces[:i], race) {
			return fmt.Errorf("a row names %q twice", race)
		}
	}
	return nil
}

// overlap reports whether two rows can apply to the same request: their
// presets and their races both meet.
func overlap(a, b RankChoiceRow) bool {
	return keysMeet(a.Preset, b.Preset) && racesMeet(a.Races, b.Races)
}

// keysMeet reports whether two rows' keys ("" for every value) share a value.
func keysMeet(a, b string) bool { return a == "" || b == "" || a == b }

// racesMeet reports whether two rows' race keys (none for every race) share a race.
func racesMeet(a, b []string) bool {
	if len(a) == 0 || len(b) == 0 {
		return true
	}
	return slices.ContainsFunc(a, func(race string) bool { return slices.Contains(b, race) })
}

var embeddedRankChoices = sync.OnceValues(func() (map[string]rankChoices, error) {
	return parseRankChoices(rankChoicesJSON)
})

// specRankChoices is spec's rows under preset for race from the embedded
// table; none when the spec has no row (or the embed is unreadable, which
// TestEmbeddedRankChoicesMatchCurated fails on).
func specRankChoices(spec, preset, race string) rankChoices {
	all, err := embeddedRankChoices()
	if err != nil {
		return nil
	}
	return all[spec].forRequest(preset, race)
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
