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
	"encoding/json"
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
// preset for a character of faction.
func (c rankChoices) forRequest(preset, faction string) rankChoices {
	out := make(rankChoices, len(c))
	for id, rows := range c {
		for _, row := range rows {
			if appliesTo(row.Preset, preset) && appliesTo(row.Faction, faction) {
				out[id] = append(out[id], row)
			}
		}
	}
	return out
}

// appliesTo reports whether a row keyed key ("" for every value) applies to value.
func appliesTo(key, value string) bool { return key == "" || key == value }

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
				if row.Faction != "" && row.Faction != AllianceFaction && row.Faction != HordeFaction {
					return nil, fmt.Errorf("rank choices: %s %d names faction %q; a row's faction is %q or %q", spec, sp.Authored, row.Faction, AllianceFaction, HordeFaction)
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

// overlap reports whether two rows can apply to the same request: their
// presets and their factions both meet.
func overlap(a, b RankChoiceRow) bool {
	return keysMeet(a.Preset, b.Preset) && keysMeet(a.Faction, b.Faction)
}

// keysMeet reports whether two rows' keys ("" for every value) share a value.
func keysMeet(a, b string) bool { return a == "" || b == "" || a == b }

var embeddedRankChoices = sync.OnceValues(func() (map[string]rankChoices, error) {
	return parseRankChoices(rankChoicesJSON)
})

// specRankChoices is spec's rows under preset for faction from the embedded
// table; none when the spec has no row (or the embed is unreadable, which
// TestEmbeddedRankChoicesMatchCurated fails on).
func specRankChoices(spec, preset, faction string) rankChoices {
	all, err := embeddedRankChoices()
	if err != nil {
		return nil
	}
	return all[spec].forRequest(preset, faction)
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
