// api/internal/directory/parse.go
package directory

import (
	"encoding/json"
	"fmt"
)

// EquippedItem is one slot of a directory character's equipment summary: what is worn, never
// the full Blizzard payload (stats, binding, durability stay with Blizzard).
type EquippedItem struct {
	Slot    string `json:"slot"`
	ItemID  int64  `json:"item_id"`
	Name    string `json:"name"`
	Quality string `json:"quality,omitempty"`
	Level   int    `json:"level,omitempty"`
}

type equipmentResponse struct {
	EquippedItems []struct {
		Item struct {
			ID int64 `json:"id"`
		} `json:"item"`
		Slot struct {
			Type string `json:"type"`
		} `json:"slot"`
		Quality struct {
			Type string `json:"type"`
		} `json:"quality"`
		Name  string `json:"name"`
		Level struct {
			Value int `json:"value"`
		} `json:"level"`
	} `json:"equipped_items"`
}

// summarizeEquipment is the equipment summary of a /equipment answer.
func summarizeEquipment(raw []byte) ([]EquippedItem, error) {
	var res equipmentResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, fmt.Errorf("directory: equipment: %w", err)
	}
	out := make([]EquippedItem, 0, len(res.EquippedItems))
	for _, it := range res.EquippedItems {
		if it.Item.ID == 0 {
			continue
		}
		out = append(out, EquippedItem{
			Slot: it.Slot.Type, ItemID: it.Item.ID, Name: it.Name, Quality: it.Quality.Type, Level: it.Level.Value,
		})
	}
	return out, nil
}

type specializationsResponse struct {
	Groups []struct {
		IsActive bool `json:"is_active"`
		Specs    []struct {
			Name        string `json:"specialization_name"`
			SpentPoints int    `json:"spent_points"`
		} `json:"specializations"`
	} `json:"specialization_groups"`
}

// activeSpec is the talent tree the character's active group spent the most points in
// ("Arms"); "" when no point is spent or no group is active.
func activeSpec(raw []byte) (string, error) {
	var res specializationsResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		return "", fmt.Errorf("directory: specializations: %w", err)
	}
	best, most := "", 0
	for _, g := range res.Groups {
		if !g.IsActive {
			continue
		}
		for _, sp := range g.Specs {
			if sp.SpentPoints > most {
				best, most = sp.Name, sp.SpentPoints
			}
		}
	}
	return best, nil
}

// Profession is one profession and its skill, as the directory keeps it.
type Profession struct {
	Name  string `json:"name"`
	Skill int    `json:"skill"`
	Max   int    `json:"max,omitempty"`
}

type professionEntry struct {
	Profession struct {
		Name string `json:"name"`
	} `json:"profession"`
	SkillPoints    int `json:"skill_points"`
	MaxSkillPoints int `json:"max_skill_points"`
	Tiers          []struct {
		SkillPoints    int `json:"skill_points"`
		MaxSkillPoints int `json:"max_skill_points"`
	} `json:"tiers"`
}

type professionsResponse struct {
	Primaries   []professionEntry `json:"primaries"`
	Secondaries []professionEntry `json:"secondaries"`
}

// summarizeProfessions is every primary then secondary profession with its highest skill.
func summarizeProfessions(raw []byte) ([]Profession, error) {
	var res professionsResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, fmt.Errorf("directory: professions: %w", err)
	}
	out := []Profession{}
	for _, p := range append(res.Primaries, res.Secondaries...) {
		if p.Profession.Name == "" {
			continue
		}
		prof := Profession{Name: p.Profession.Name, Skill: p.SkillPoints, Max: p.MaxSkillPoints}
		for _, t := range p.Tiers {
			if t.SkillPoints > prof.Skill {
				prof.Skill, prof.Max = t.SkillPoints, t.MaxSkillPoints
			}
		}
		out = append(out, prof)
	}
	return out, nil
}
