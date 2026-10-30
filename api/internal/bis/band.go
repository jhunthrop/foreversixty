// api/internal/bis/band.go
package bis

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Alternative is one of a slot's runner-up items, carrying the ranker's own sim-verified
// gap against the pick (data/builds/<build>/bis/<spec>.json's own bands[].slots[].
// alternatives[]).
type Alternative struct {
	ItemID   int     `json:"item_id"`
	ItemName string  `json:"item_name"`
	DpsDelta float64 `json:"dps_delta"`
	// Source is where the item comes from, in the game's words ("PvP rank 14 · Champion ·
	// Horde", "Blacksmithing").
	Source string `json:"source"`
	// SourceKind is the source's kind ("vendor", "rep", "quest", "crafted", "dungeon", ...),
	// from which the site's BiS page resolves the full source line (loot.json).
	SourceKind string `json:"source_kind"`
	// EffectUnmodelled marks an item whose proc, on-use or set effect the sim leaves out, so
	// DpsDelta misses what makes the item good.
	EffectUnmodelled bool `json:"effect_unmodelled"`
}

// Slot is one band's pick for one gear slot. ItemID <= 0 means the band has no known
// source for this slot (the BiS file's own "nothing to recommend here" shape) - callers
// skip such a slot outright, never treat 0 as an item id.
type Slot struct {
	Slot         string        `json:"slot"`
	ItemID       int           `json:"item_id"`
	ItemName     string        `json:"item_name"`
	Source       string        `json:"source"`
	SourceKind   string        `json:"source_kind"`
	Alternatives []Alternative `json:"alternatives"`
	// Score is the pick's stat-weight score in the band's ScoreUnit; nil on a sim-decided row
	// (an on-use trinket, a weapon) that carries sim_dps instead.
	Score *float64 `json:"score"`
	// EffectUnmodelled marks a pick whose proc, on-use or set effect the sim leaves out: no
	// figure it carries includes the effect that makes it the pick.
	EffectUnmodelled bool `json:"effect_unmodelled"`
	// EmptyReason says why a slot has no pick ("two_hand_equipped": the main hand is a
	// two-hander); "" otherwise.
	EmptyReason string `json:"empty_reason"`
}

// EmptyTwoHand is the off hand's EmptyReason when the pick in the main hand is two-handed.
const EmptyTwoHand = "two_hand_equipped"

// TankScoreUnit is the ScoreUnit of a tank band, whose figures are Forever Score points.
const TankScoreUnit = "tank_score"

// Band is one spec's BiS picks for one level band and faction.
type Band struct {
	Slots []Slot
	// ScoreUnit is what the band's scores and deltas count: "reference_stat_points" (DPS) or
	// TankScoreUnit.
	ScoreUnit string
	// ReferenceDpsPerPoint turns a Score into DPS; nil when the band states none.
	ReferenceDpsPerPoint *float64
	// GeneratedAt is when the file was ranked (RFC 3339), "" when it does not say.
	GeneratedAt string
}

// BySlot indexes Slots by their own slot name, for a caller comparing one character's gear
// against the whole band.
func (b Band) BySlot() map[string]Slot {
	out := make(map[string]Slot, len(b.Slots))
	for _, s := range b.Slots {
		out[s.Slot] = s
	}
	return out
}

// bisFile and bisBand are the on-disk JSON shape - unexported, since the loaders are the
// only things that ever need the whole file.
type bisFile struct {
	GeneratedAt string    `json:"generated_at"`
	Bands       []bisBand `json:"bands"`
}

type bisBand struct {
	Band                 int      `json:"band"`
	Faction              string   `json:"faction"`
	Preset               string   `json:"preset"`
	Slots                []Slot   `json:"slots"`
	ScoreUnit            string   `json:"score_unit"`
	ReferenceDpsPerPoint *float64 `json:"reference_dps_per_point"`
}

// band is the entry as the loaders return it.
func (b bisBand) band() Band {
	return Band{Slots: b.Slots, ScoreUnit: b.ScoreUnit, ReferenceDpsPerPoint: b.ReferenceDpsPerPoint}
}

// The presets a band entry is ranked under; an entry with no preset is bare.
const (
	BarePreset = "bare"
	RaidPreset = "raid"
)

func (b bisBand) preset() string {
	if b.Preset == "" {
		return BarePreset
	}
	return b.Preset
}

// LoadBand reads fileSlug's band/faction entry out of dataDir/build/bis/<fileSlug>.json
// (dataDir is the API's own TreeDataDir config - one directory per client build, matching
// api/cmd/seedguild/gear.go's own bisDir()). ok is false, with no error, for every "there is
// nothing to compare against" case: an empty fileSlug (FileSlugFor found no catalogue
// entry), a build with no BiS directory at all, a spec with no BiS file (every tank/healer
// spec - the catalogue only ever simmed DPS specs), or a file with no matching band/faction
// entry. A real I/O or parse error on a file that does exist is still returned as an error.
func LoadBand(dataDir, build, fileSlug string, band int, faction string) (Band, bool, error) {
	file, ok, err := readFile(dataDir, build, fileSlug)
	if err != nil || !ok {
		return Band{}, false, err
	}
	for _, b := range file.Bands {
		if b.Band == band && b.Faction == faction {
			return b.band(), true, nil
		}
	}
	return Band{}, false, nil
}

// LoadPresetBand is LoadBand under a preset, the site's own selection
// (web/src/lib/bis/presets.ts selectBand): the entry ranked under preset, else the bare one,
// else the first entry for the band and faction. ok is false as LoadBand's.
func LoadPresetBand(dataDir, build, fileSlug string, band int, faction, preset string) (Band, bool, error) {
	file, ok, err := readFile(dataDir, build, fileSlug)
	if err != nil || !ok {
		return Band{}, false, err
	}
	stamp := func(b Band) Band {
		b.GeneratedAt = file.GeneratedAt
		return b
	}
	var candidates []bisBand
	for _, b := range file.Bands {
		if b.Band == band && b.Faction == faction {
			candidates = append(candidates, b)
		}
	}
	for _, want := range []string{preset, BarePreset} {
		for _, b := range candidates {
			if b.preset() == want {
				return stamp(b.band()), true, nil
			}
		}
	}
	if len(candidates) > 0 {
		return stamp(candidates[0].band()), true, nil
	}
	return Band{}, false, nil
}

// readFile reads dataDir/build/bis/<fileSlug>.json; ok is false, with no error, when there is
// no such file (or no slug or directory to look in).
func readFile(dataDir, build, fileSlug string) (bisFile, bool, error) {
	if fileSlug == "" || dataDir == "" {
		return bisFile{}, false, nil
	}
	path := filepath.Join(dataDir, build, "bis", fileSlug+".json")
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return bisFile{}, false, nil
		}
		return bisFile{}, false, fmt.Errorf("bis: read %s: %w", path, err)
	}
	var file bisFile
	if err := json.Unmarshal(raw, &file); err != nil {
		return bisFile{}, false, fmt.Errorf("bis: parse %s: %w", path, err)
	}
	return file, true, nil
}
