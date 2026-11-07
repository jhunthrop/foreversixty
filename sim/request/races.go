package request

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type raceClassCombo struct {
	RaceID  int `json:"race_id"`
	ClassID int `json:"class_id"`
}

type raceOrClassRow struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Faction string `json:"faction"`
}

func readClientTable(buildDir, name string, into any) error {
	path := filepath.Join(buildDir, name)
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}
	if err := json.Unmarshal(b, into); err != nil {
		return fmt.Errorf("decoding %s: %w", path, err)
	}
	return nil
}

// LegalRaces is every race classSlug can be in faction, in races.json order,
// read from the build's client race-class combinations (combos.json,
// Forever's added pairs included). It is the one source of which races a
// class can be: the BiS ranker's race sweep and spec-breakdown's -race guard
// both read it.
func LegalRaces(buildDir, classSlug, faction string) ([]string, error) {
	var races, classes []raceOrClassRow
	var combos []raceClassCombo
	for name, into := range map[string]any{"races.json": &races, "classes.json": &classes, "combos.json": &combos} {
		if err := readClientTable(buildDir, name, into); err != nil {
			return nil, err
		}
	}
	classID := 0
	for _, c := range classes {
		if c.Slug == classSlug {
			classID = c.ID
		}
	}
	if classID == 0 {
		return nil, fmt.Errorf("class %q is not in %s/classes.json", classSlug, buildDir)
	}
	allowed := make(map[int]bool)
	for _, c := range combos {
		if c.ClassID == classID {
			allowed[c.RaceID] = true
		}
	}
	var out []string
	for _, r := range races {
		if allowed[r.ID] && r.Faction == faction {
			out = append(out, r.Slug)
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no %s race can be a %s in %s/combos.json", faction, classSlug, buildDir)
	}
	return out, nil
}
