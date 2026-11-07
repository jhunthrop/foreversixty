package activity

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Curated is what the feed reads from data/curated: which level band each instance's boards
// count for, and which zones are raids.
type Curated struct {
	// ZoneBands maps a log zone name to its band: the band of its level range's midpoint
	// (instance-levels.json).
	ZoneBands map[string]int
	// RaidZones are the first tier's raids by log zone and by name (raids/first-tier.json).
	RaidZones map[string]bool
}

// LoadCurated reads the two files under dir (the repository's data/curated; /curated in
// the image).
func LoadCurated(dir string) (Curated, error) {
	var levels struct {
		Zones map[string]struct {
			Name     string `json:"name"`
			LevelMin int    `json:"level_min"`
			LevelMax int    `json:"level_max"`
		} `json:"zones"`
	}
	if err := readJSON(filepath.Join(dir, "instance-levels.json"), &levels); err != nil {
		return Curated{}, err
	}
	var raids struct {
		Zones []struct {
			Name    string  `json:"name"`
			LogZone *string `json:"log_zone"`
		} `json:"zones"`
	}
	if err := readJSON(filepath.Join(dir, "raids", "first-tier.json"), &raids); err != nil {
		return Curated{}, err
	}
	c := Curated{ZoneBands: map[string]int{}, RaidZones: map[string]bool{}}
	for _, z := range levels.Zones {
		c.ZoneBands[z.Name] = BandOf((z.LevelMin + z.LevelMax) / 2)
	}
	for _, z := range raids.Zones {
		c.RaidZones[z.Name] = true
		if z.LogZone != nil {
			c.RaidZones[*z.LogZone] = true
		}
	}
	return c, nil
}

func readJSON(path string, into any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("activity: read %s: %w", path, err)
	}
	if err := json.Unmarshal(raw, into); err != nil {
		return fmt.Errorf("activity: parse %s: %w", path, err)
	}
	return nil
}
