package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// specInfo is what this command needs from data/curated/specs.json.
type specInfo struct {
	Spec      string `json:"spec"`
	ClassSlug string `json:"class_slug"`
	SpecSlug  string `json:"spec_slug"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	TreeIndex int    `json:"tree_index"`

	ReferenceStat string `json:"reference_stat"`
}

// dpsRole is the only role a DPS search answers for: a healer's or a
// tank's objective is not damage.
const dpsRole = "dps"

// writtenAPL is data/curated/apl/<spec>.json's state for a spec with
// a real rotation.
const writtenAPL = "written"

func loadSpec(repoRoot, spec string) (specInfo, error) {
	b, err := os.ReadFile(filepath.Join(repoRoot, "data", "curated", "specs.json"))
	if err != nil {
		return specInfo{}, fmt.Errorf("reading curated/specs.json: %w", err)
	}
	var specs []specInfo
	if err := json.Unmarshal(b, &specs); err != nil {
		return specInfo{}, fmt.Errorf("decoding curated/specs.json: %w", err)
	}
	for _, s := range specs {
		if s.Spec == spec {
			return s, nil
		}
	}
	return specInfo{}, fmt.Errorf("no spec %q in curated/specs.json", spec)
}

func aplState(repoRoot, spec string) (string, error) {
	path := filepath.Join(repoRoot, "data", "curated", "apl", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", path, err)
	}
	var f struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return "", fmt.Errorf("decoding %s: %w", path, err)
	}
	return f.State, nil
}

// bisBand is one band of data/builds/<build>/bis/<spec>.json.
type bisBand struct {
	Band    int     `json:"band"`
	Preset  string  `json:"preset"`
	Faction string  `json:"faction"`
	Race    string  `json:"race"`
	Talents string  `json:"talents"`
	SetDPS  float64 `json:"set_dps"`
	Slots   []struct {
		Slot   string `json:"slot"`
		ItemID int    `json:"item_id"`
	} `json:"slots"`
}

// loadBISBand is the committed BiS set for one band and faction - the
// gear every candidate wears, so only the talents vary.
func loadBISBand(buildDir, spec string, level int, faction, preset string) (bisBand, error) {
	path := filepath.Join(buildDir, "bis", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return bisBand{}, fmt.Errorf("reading %s: %w", path, err)
	}
	var f struct {
		Bands []bisBand `json:"bands"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return bisBand{}, fmt.Errorf("decoding %s: %w", path, err)
	}
	for _, band := range f.Bands {
		if band.Band == level && band.Faction == faction && band.Preset == preset {
			return band, nil
		}
	}
	return bisBand{}, fmt.Errorf("%s has no level-%d %s %s band", path, level, faction, preset)
}

func (b bisBand) gear() []api.GearSlot {
	var out []api.GearSlot
	for _, s := range b.Slots {
		if s.ItemID != 0 {
			out = append(out, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID})
		}
	}
	return out
}
