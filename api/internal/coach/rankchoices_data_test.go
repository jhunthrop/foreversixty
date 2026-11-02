package coach

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
	"testing"
)

// Every spell data/curated/rank-choices.json names is a trainable family of its class in the
// newest build, under the same name, so F5 reads each choice; a renamed spell would leave
// the coach calling the site's own downrank a mistake again.
func TestRankChoicesJoinTheTrainableFamilies(t *testing.T) {
	body, err := os.ReadFile("../../../data/curated/rank-choices.json")
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Specs map[string][]struct {
			Spell string `json:"spell"`
			Bands []struct {
				FromLevel int `json:"from_level"`
				Rank      int `json:"rank"`
			} `json:"bands"`
		} `json:"specs"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		t.Fatal(err)
	}
	if len(file.Specs) == 0 {
		t.Fatal("rank-choices.json names no spec")
	}
	for spec, spells := range file.Specs {
		class, _, _ := strings.Cut(spec, "-")
		trainable, err := ParseTrainables(buildFile(t, "trainables", class+".json"))
		if err != nil {
			t.Fatal(err)
		}
		for _, sp := range spells {
			if _, ok := trainable[sp.Spell]; !ok {
				t.Errorf("%s: %q is no trainable family of %s", spec, sp.Spell, class)
				continue
			}
			level := sp.Bands[0].FromLevel
			got, err := rankChoicesAt(spec, trainable, level)
			if err != nil {
				t.Fatal(err)
			}
			if !slices.Contains(got[sp.Spell], sp.Bands[0].Rank) {
				t.Errorf("%s %s at %d: chosen %v lacks rank %d", spec, sp.Spell, level, got[sp.Spell], sp.Bands[0].Rank)
			}
		}
	}
}
