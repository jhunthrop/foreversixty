package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// dpsRole is the only role this search answers for: a healer's or a
// tank's objective is not damage.
const dpsRole = "dps"

// writtenAPL is data/curated/apl/<spec>.json's state for a spec with a
// real, curated rotation - the ladder's own writtenAPL constant,
// copied rather than imported because sim/request does not export it.
const writtenAPL = "written"

// specInfo is what this command needs from data/curated/specs.json.
type specInfo struct {
	Spec      string `json:"spec"`
	ClassSlug string `json:"class_slug"`
	SpecSlug  string `json:"spec_slug"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	TreeIndex int    `json:"tree_index"`
}

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

// curatedAPL is the half of data/curated/apl/<spec>.json this command
// reads: the state gate and the rotation this search starts from and
// ultimately proposes a replacement for.
type curatedAPL struct {
	Spec     string          `json:"spec"`
	State    string          `json:"state"`
	Rotation json.RawMessage `json:"rotation"`
}

func loadCuratedAPL(repoRoot, spec string) (curatedAPL, error) {
	path := filepath.Join(repoRoot, "data", "curated", "apl", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return curatedAPL{}, fmt.Errorf("reading %s: %w", path, err)
	}
	var c curatedAPL
	if err := json.Unmarshal(b, &c); err != nil {
		return curatedAPL{}, fmt.Errorf("decoding %s: %w", path, err)
	}
	if c.Spec != spec {
		return curatedAPL{}, fmt.Errorf("%s declares spec %q", path, c.Spec)
	}
	return c, nil
}

// bisBand is one band of data/builds/<build>/bis/<spec>.json - the
// gear and race every sim in this search wears; only the rotation
// varies, the same discipline sim/cmd/talent-search's data.go uses for
// talents (there, gear is held fixed; here, gear AND talents are held
// fixed).
type bisBand struct {
	Band    int    `json:"band"`
	Faction string `json:"faction"`
	Race    string `json:"race"`
	Talents string `json:"talents"`
	Slots   []struct {
		Slot   string `json:"slot"`
		ItemID int    `json:"item_id"`
	} `json:"slots"`
}

func loadBISBand(buildDir, spec string, level int, faction string) (bisBand, error) {
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
		if band.Band == level && band.Faction == faction {
			return band, nil
		}
	}
	return bisBand{}, fmt.Errorf("%s has no level-%d %s band", path, level, faction)
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
