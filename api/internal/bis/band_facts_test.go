package bis

import (
	"os"
	"path/filepath"
	"testing"
)

const factsFixture = `{"bands": [
 {"band": 60, "faction": "alliance", "preset": "raid", "role": "tank", "talents": "05-31-0", "talent_points": 9,
  "weights": [{"stat": "agility", "weight": 2.5}, {"stat": "armor", "weight": 0.1}],
  "slots": [{"slot": "head", "item_id": 3, "item_name": "Raid Helm"}]},
 {"band": 40, "faction": "alliance", "slots": [{"slot": "head", "item_id": 5, "item_name": "Level Helm"}]}
]}`

// A band carries what the home answer reads beside its picks: the preset it was ranked under,
// its role, its talent build and its stat weights.
func TestABandCarriesItsPresetRoleTalentsAndWeights(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "b1", "bis")
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, "warrior-protection.json"), []byte(factsFixture), 0o600); err != nil {
		t.Fatal(err)
	}
	band, found, err := LoadPresetBand(dir, "b1", "warrior-protection", 60, "alliance", RaidPreset)
	if err != nil || !found {
		t.Fatalf("found %v err %v", found, err)
	}
	if band.Band != 60 || band.Preset != RaidPreset || band.Role != "tank" || band.Talents != "05-31-0" || band.TalentPoints != 9 {
		t.Fatalf("band = %+v", band)
	}
	if w := band.WeightsByStat(); len(w) != 2 || w["agility"] != 2.5 || w["armor"] != 0.1 {
		t.Fatalf("weights = %v", w)
	}
	bare, _, _ := LoadPresetBand(dir, "b1", "warrior-protection", 40, "alliance", RaidPreset)
	if bare.Preset != BarePreset || bare.Band != 40 || len(bare.WeightsByStat()) != 0 {
		t.Fatalf("an entry with no preset reads bare: %+v", bare)
	}
}
