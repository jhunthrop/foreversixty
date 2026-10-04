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
}

// Slot is one band's pick for one gear slot. ItemID <= 0 means the band has no known
// source for this slot (the BiS file's own "nothing to recommend here" shape) - callers
// skip such a slot outright, never treat 0 as an item id.
type Slot struct {
	Slot         string        `json:"slot"`
	ItemID       int           `json:"item_id"`
	ItemName     string        `json:"item_name"`
	Alternatives []Alternative `json:"alternatives"`
}

// Band is one spec's BiS picks for one level band and faction.
type Band struct {
	Slots []Slot
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

// bisFile and bisBand are the on-disk JSON shape - unexported, since LoadBand is the only
// thing that ever needs the whole file.
type bisFile struct {
	Bands []bisBand `json:"bands"`
}

type bisBand struct {
	Band    int    `json:"band"`
	Faction string `json:"faction"`
	Slots   []Slot `json:"slots"`
}

// LoadBand reads fileSlug's band/faction entry out of dataDir/build/bis/<fileSlug>.json
// (dataDir is the API's own TreeDataDir config - one directory per client build, matching
// api/cmd/seedguild/gear.go's own bisDir()). ok is false, with no error, for every "there is
// nothing to compare against" case: an empty fileSlug (FileSlugFor found no catalogue
// entry), a build with no BiS directory at all, a spec with no BiS file (every tank/healer
// spec - the catalogue only ever simmed DPS specs), or a file with no matching band/faction
// entry. A real I/O or parse error on a file that does exist is still returned as an error.
func LoadBand(dataDir, build, fileSlug string, band int, faction string) (Band, bool, error) {
	if fileSlug == "" || dataDir == "" {
		return Band{}, false, nil
	}
	path := filepath.Join(dataDir, build, "bis", fileSlug+".json")
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Band{}, false, nil
		}
		return Band{}, false, fmt.Errorf("bis: read %s: %w", path, err)
	}
	var file bisFile
	if err := json.Unmarshal(raw, &file); err != nil {
		return Band{}, false, fmt.Errorf("bis: parse %s: %w", path, err)
	}
	for _, b := range file.Bands {
		if b.Band == band && b.Faction == faction {
			return Band{Slots: b.Slots}, true, nil
		}
	}
	return Band{}, false, nil
}
