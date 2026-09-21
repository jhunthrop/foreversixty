// logs/engine/mechanics/utility/utility.go
// Package utility holds the curated per-spec table of raid buffs, debuffs,
// cooldowns and informational talent modifiers a spec is expected to bring
// to a raid, for the performance rating engine's Utility component
// (docs/superpowers/specs/2026-09-21-performance-rating-design.md §3.2).
// Files ship inside the engine the same way logs/engine/mechanics ships its
// own tables: embedded and validated at build time, a panic on a malformed
// file at process start rather than a runtime surprise.
package utility

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

// Kind is what an owned entry means for scoring.
type Kind string

const (
	// Buff is a self/raid buff the spec applies to allies, scored by
	// uptime share.
	Buff Kind = "buff"
	// Debuff is an enemy debuff the spec applies, scored by uptime share.
	Debuff Kind = "debuff"
	// Cooldown is a burst/utility cooldown whose use, not uptime, is what
	// should be tracked (e.g. Innervate, Power Infusion).
	Cooldown Kind = "cooldown"
	// TalentModifier is informational only, never independently scored as
	// an aura -- most specs have a handful, recorded for completeness so a
	// reviewer does not wonder why it is missing from the card.
	TalentModifier Kind = "talent-modifier"
)

var kinds = map[Kind]bool{Buff: true, Debuff: true, Cooldown: true, TalentModifier: true}

// Entry is one ability a spec owns.
type Entry struct {
	SpellID int64  `json:"spell_id"`
	Name    string `json:"name"`
	Kind    Kind   `json:"kind"`
	// Target is "enemy" or "ally", where meaningful; empty for a
	// talent-modifier or cooldown entry it does not apply to.
	Target string `json:"target,omitempty"`
	// Verified is the spell id's provenance: "<build>/spells.json#<id>",
	// the same convention logs/engine/mechanics/encounters.json's own
	// note describes.
	Verified string `json:"verified"`
	Note     string `json:"note,omitempty"`
}

// Table is one spec's owned utility.
type Table struct {
	Spec  string  `json:"spec"`
	Owned []Entry `json:"owned"`
}

//go:embed tables/*.json
var tables embed.FS

var _ = mustParseAll()

// mustParseAll parses every embedded table and panics on the first one
// that is not the format or whose file name does not match its spec.
func mustParseAll() int {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		panic(fmt.Errorf("utility: reading the embedded tables: %w", err))
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := "tables/" + entry.Name()
		data, err := tables.ReadFile(path)
		if err != nil {
			panic(fmt.Errorf("utility: reading %s: %w", path, err))
		}
		t, err := Parse(data)
		if err != nil {
			panic(fmt.Errorf("utility: %s: %w", path, err))
		}
		if want := t.Spec + ".json"; entry.Name() != want {
			panic(fmt.Errorf("utility: %s holds spec %q, so it must be named %s", path, t.Spec, want))
		}
	}
	return len(entries)
}

// Parse reads a table and refuses one that is not the format.
func Parse(data []byte) (Table, error) {
	var t Table
	if err := json.Unmarshal(data, &t); err != nil {
		return Table{}, fmt.Errorf("utility: %w", err)
	}
	if strings.TrimSpace(t.Spec) == "" {
		return Table{}, fmt.Errorf("utility: spec is required")
	}
	seen := make(map[int64]bool, len(t.Owned))
	for i, e := range t.Owned {
		if e.SpellID <= 0 {
			return Table{}, fmt.Errorf("owned[%d]: spell_id must be positive", i)
		}
		if !kinds[e.Kind] {
			return Table{}, fmt.Errorf("owned[%d]: kind %q is not buff, debuff, cooldown or talent-modifier", i, e.Kind)
		}
		if strings.TrimSpace(e.Verified) == "" {
			return Table{}, fmt.Errorf("owned[%d]: verified provenance is required", i)
		}
		if seen[e.SpellID] {
			return Table{}, fmt.Errorf("owned[%d]: spell_id %d is listed twice", i, e.SpellID)
		}
		seen[e.SpellID] = true
	}
	return t, nil
}

// Load returns the embedded table for a spec slug (data/curated/specs.json's
// "spec" field, e.g. "warrior-protection"), and false when there is none.
func Load(specSlug string) (Table, bool) {
	data, err := tables.ReadFile("tables/" + specSlug + ".json")
	if err != nil {
		return Table{}, false
	}
	t, err := Parse(data)
	if err != nil {
		// An embedded table that does not parse is a build defect, not a
		// runtime case.
		panic(err)
	}
	return t, true
}
