// api/internal/bnetbuild/tables.go
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// TalentTable is one class's talent data within a loaded trees.Data build (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2: "loaded once per class
// from data/builds/<build>/talents/<class>.json"). trees.Load already reads that file for
// every class at process startup; this wraps it rather than re-reading it.
type TalentTable struct {
	Build   *trees.Build
	ClassID int
}

// EnchantTable is the set of known permanent-enchant ids (data/builds/<build>/
// enchants.json), keyed by the same id Blizzard's equipment response calls
// enchantment_id (spec §2.2).
type EnchantTable map[int]bool

type enchantRow struct {
	ID int `json:"id"`
}

// LoadEnchantTable reads enchants.json's flat row list into a lookup set. Unknown JSON
// fields are ignored on purpose — the file carries name/icon/slots/stats this package has
// no use for.
func LoadEnchantTable(path string) (EnchantTable, error) {
	var rows []enchantRow
	if err := readJSON(path, &rows); err != nil {
		return nil, err
	}
	out := make(EnchantTable, len(rows))
	for _, r := range rows {
		out[r.ID] = true
	}
	return out, nil
}

// SuffixTable maps a random suffix's display name, exactly as it appears appended to an
// item's base name (e.g. "of the Falcon"), to its id (data/builds/<build>/suffixes.json).
type SuffixTable map[string]int

type suffixRow struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LoadSuffixTable reads suffixes.json into a name-to-id lookup.
func LoadSuffixTable(path string) (SuffixTable, error) {
	var rows []suffixRow
	if err := readJSON(path, &rows); err != nil {
		return nil, err
	}
	out := make(SuffixTable, len(rows))
	for _, r := range rows {
		out[r.Name] = r.ID
	}
	return out, nil
}

// RaceTable maps Blizzard's race.name (title case, e.g. "Night Elf") to the site's own race
// slug (spec §2.2).
type RaceTable map[string]string

// RaceTableFrom builds a RaceTable from a loaded build's races, rather than re-parsing
// races.json (trees.Load already did; trees.Build.Races exposes it).
func RaceTableFrom(b *trees.Build) RaceTable {
	races := b.Races()
	out := make(RaceTable, len(races))
	for _, r := range races {
		out[r.Name] = r.Slug
	}
	return out
}

// Tables is every data table Encode needs, loaded once per process from the site's active
// client build (spec §2.3: "Data tables are loaded once per process from
// data/builds/<active build>/").
type Tables struct {
	Build    string
	Trees    *trees.Build
	Enchants EnchantTable
	Suffixes SuffixTable
	Races    RaceTable
}

// LoadTables resolves the active build (trees.Data's newest client build, the same rule
// dataaddon's job already uses) and loads its enchant, suffix and race tables. ok is false,
// with a zero Tables, when data holds no build at all — the caller (bnetimport, cmd/api)
// logs and runs with Battle.net-sourced builds disabled, same as when TreeDataDir is empty.
func LoadTables(dir string, data *trees.Data) (Tables, bool, error) {
	active, ok := data.Latest()
	if !ok {
		return Tables{}, false, nil
	}
	buildDir := filepath.Join(dir, active.Version)
	enchants, err := LoadEnchantTable(filepath.Join(buildDir, "enchants.json"))
	if err != nil {
		return Tables{}, false, fmt.Errorf("bnetbuild: load tables %s: %w", active.Version, err)
	}
	suffixes, err := LoadSuffixTable(filepath.Join(buildDir, "suffixes.json"))
	if err != nil {
		return Tables{}, false, fmt.Errorf("bnetbuild: load tables %s: %w", active.Version, err)
	}
	return Tables{
		Build: active.Version, Trees: active, Enchants: enchants, Suffixes: suffixes,
		Races: RaceTableFrom(active),
	}, true, nil
}

func readJSON(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("bnetbuild: open %s: %w", path, err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(v); err != nil {
		return fmt.Errorf("bnetbuild: decode %s: %w", path, err)
	}
	return nil
}
