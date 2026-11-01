package upgrades

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// parityDir holds one fixture character over a real band and item rows of build 1.60.1.70658
// (input.json) and the site's own answer for it (site.json), which
// web/src/lib/home/upgrades-parity.test.ts writes and pins from the TypeScript side.
const parityDir = "testdata/parity"

type parityRow struct {
	Slot          string   `json:"slot"`
	PickID        int      `json:"pick_id"`
	WornID        int      `json:"worn_id"`
	WornUnknown   bool     `json:"worn_unknown"`
	Gain          *float64 `json:"gain"`
	NotSimChecked bool     `json:"not_sim_checked"`
}

type parityBest struct {
	Slot     string `json:"slot"`
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

type parityAnswer struct {
	Upgrades           []parityRow  `json:"upgrades"`
	AlreadyBest        []parityBest `json:"already_best"`
	TotalGain          float64      `json:"total_gain"`
	NotSimCheckedCount int          `json:"not_sim_checked_count"`
	Unit               Unit         `json:"unit"`
}

// parityInput is the fixture character: a band entry, the worn gear and the item rows.
type parityInput struct {
	Band  json.RawMessage `json:"band"`
	Gear  map[string]int  `json:"gear"`
	Items []trees.Item    `json:"items"`
}

func readParity(t *testing.T) (bis.Band, map[string]int, Items, parityAnswer) {
	t.Helper()
	var in parityInput
	readJSON(t, "input.json", &in)
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "b", "bis"), 0o755); err != nil {
		t.Fatal(err)
	}
	file := []byte(`{"bands": [` + string(in.Band) + `]}`)
	if err := os.WriteFile(filepath.Join(dir, "b", "bis", "hunter-beast-mastery.json"), file, 0o600); err != nil {
		t.Fatal(err)
	}
	band, found, err := bis.LoadPresetBand(dir, "b", "hunter-beast-mastery", 60, "alliance", bis.RaidPreset)
	if err != nil || !found {
		t.Fatalf("band: %v %v", found, err)
	}
	byID := map[int]trees.Item{}
	for _, it := range in.Items {
		byID[it.ID] = it
	}
	var site parityAnswer
	readJSON(t, "site.json", &site)
	return band, in.Gear, func(id int) (trees.Item, bool) { it, ok := byID[id]; return it, ok }, site
}

func readJSON(t *testing.T, name string, v any) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(parityDir, name))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, v); err != nil {
		t.Fatalf("%s: %v", name, err)
	}
}

// TestCompareMatchesTheSite: the same upgrades, gains, already-best slots, total and
// not-sim-checked count as web/src/lib/home/upgrades.ts on the fixture character.
func TestCompareMatchesTheSite(t *testing.T) {
	band, gear, items, site := readParity(t)
	got := Compare(band, gear, items)
	round := func(v float64) float64 { return math.Round(v*1e6) / 1e6 }
	var rows []parityRow
	for _, u := range got.Upgrades {
		r := parityRow{Slot: u.Slot, PickID: u.Pick.ItemID, WornID: u.WornID, WornUnknown: u.WornID > 0 && !u.WornKnown,
			NotSimChecked: u.NotSimChecked}
		if u.Gain != nil {
			r.Gain = ptr(round(*u.Gain))
		}
		rows = append(rows, r)
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Slot < rows[j].Slot })
	var best []parityBest
	for _, b := range got.AlreadyBest {
		best = append(best, parityBest{Slot: b.Slot, ItemID: b.ItemID, ItemName: b.ItemName})
	}
	answer := parityAnswer{Upgrades: rows, AlreadyBest: best, TotalGain: round(got.TotalGain),
		NotSimCheckedCount: got.NotSimChecked, Unit: got.Unit}
	if !reflect.DeepEqual(answer, site) {
		gotJSON, _ := json.MarshalIndent(answer, "", " ")
		wantJSON, _ := json.MarshalIndent(site, "", " ")
		t.Fatalf("the API's answer differs from the site's\n got: %s\nwant: %s", gotJSON, wantJSON)
	}
	// The fixture reaches every branch: an unknown worn item, a weapon with no listed
	// alternative, an empty slot with and without a figure, a stat-weight loss and a tie.
	if len(site.Upgrades) < 10 || len(site.AlreadyBest) < 4 || site.NotSimCheckedCount != 1 {
		t.Fatalf("the fixture lost a branch: %+v", site)
	}
}
