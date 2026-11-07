package activity

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// bisFile is the part of data/builds/<build>/bis/<spec>.json the feed and the tier
// snapshots read.
type bisFile struct {
	Spec        string    `json:"spec"`
	GeneratedAt time.Time `json:"generated_at"`
	Bands       []bisBand `json:"bands"`
}

type bisBand struct {
	Band    int     `json:"band"`
	Preset  string  `json:"preset"`
	Faction string  `json:"faction"`
	Role    string  `json:"role"`
	SetDPS  float64 `json:"set_dps"`
	Metrics *struct {
		DTPS float64 `json:"dtps"`
	} `json:"metrics"`
	Slots []bisSlot `json:"slots"`
}

type bisSlot struct {
	Slot     string `json:"slot"`
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

const barePreset = "bare"

func (b bisBand) preset() string {
	if b.Preset == "" {
		return barePreset
	}
	return b.Preset
}

func (b bisBand) role() string {
	if b.Role == "" {
		return "dps"
	}
	return b.Role
}

// loadBisFiles reads every spec's BiS file of one build, in spec order. A spec with no file
// is skipped; a file that does not parse is an error.
func loadBisFiles(dataDir, build string) ([]bisFile, error) {
	var out []bisFile
	for _, s := range specs.All {
		path := filepath.Join(dataDir, build, "bis", s.Spec+".json")
		raw, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("activity: read %s: %w", path, err)
		}
		var f bisFile
		if err := json.Unmarshal(raw, &f); err != nil {
			return nil, fmt.Errorf("activity: parse %s: %w", path, err)
		}
		f.Spec = s.Spec
		out = append(out, f)
	}
	return out, nil
}

// itemFacts are the icon and quality an item row shows.
type itemFacts struct {
	Icon    string `json:"icon"`
	Quality int    `json:"quality"`
}

// itemTable reads data/builds/<build>/items/<class>.json lazily, per class.
type itemTable struct {
	dataDir, build string
	byClass        map[string]map[int]itemFacts
}

func (t *itemTable) lookup(class string, id int) (itemFacts, bool) {
	if t.byClass == nil {
		t.byClass = map[string]map[int]itemFacts{}
	}
	items, ok := t.byClass[class]
	if !ok {
		items = map[int]itemFacts{}
		raw, err := os.ReadFile(filepath.Join(t.dataDir, t.build, "items", class+".json"))
		if err == nil {
			var f struct {
				Items []struct {
					ID int `json:"id"`
					itemFacts
				} `json:"items"`
			}
			if json.Unmarshal(raw, &f) == nil {
				for _, it := range f.Items {
					items[it.ID] = it.itemFacts
				}
			}
		}
		t.byClass[class] = items
	}
	it, found := items[id]
	return it, found
}

// slotWords is a slot as a sentence says it: "chest", "main hand", "finger".
func slotWords(slot string) string {
	s := strings.TrimRight(strings.ReplaceAll(slot, "_", " "), "12")
	return strings.TrimSpace(s)
}

// sortedBandKeys orders band entries by band, then preset (bare first), then faction.
func sortedBands(bands []bisBand) []bisBand {
	out := append([]bisBand(nil), bands...)
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Band != b.Band {
			return a.Band < b.Band
		}
		if a.preset() != b.preset() {
			return a.preset() == barePreset
		}
		return a.Faction < b.Faction
	})
	return out
}
