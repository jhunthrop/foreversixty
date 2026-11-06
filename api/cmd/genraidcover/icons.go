package main

import (
	"fmt"
	"path/filepath"
	"strconv"
)

// iconIndex is spell id -> icon stem from a build's data, best source first: spellicons.json
// (the client's SpellMisc join the site draws from), then the guide tooltips (spelltips/), then
// the talent trees (a talent's own icon for its rank spells).
type iconIndex []map[int]string

func (ix iconIndex) of(id int) string {
	for _, m := range ix {
		if stem := m[id]; stem != "" {
			return stem
		}
	}
	return ""
}

func loadIcons(buildDir string) (iconIndex, error) {
	spellIcons, err := spellIconsFile(filepath.Join(buildDir, "spellicons.json"))
	if err != nil {
		return nil, err
	}
	tips, err := spelltipIcons(filepath.Join(buildDir, "spelltips"))
	if err != nil {
		return nil, err
	}
	talents, err := talentIcons(filepath.Join(buildDir, "talents"))
	if err != nil {
		return nil, err
	}
	return iconIndex{spellIcons, tips, talents}, nil
}

func spellIconsFile(path string) (map[int]string, error) {
	var raw map[string]string
	if err := readJSON(path, &raw); err != nil {
		return nil, err
	}
	out := make(map[int]string, len(raw))
	for k, stem := range raw {
		id, err := strconv.Atoi(k)
		if err != nil {
			return nil, fmt.Errorf("%s: key %q: %w", path, k, err)
		}
		out[id] = stem
	}
	return out, nil
}

func spelltipIcons(dir string) (map[int]string, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, err
	}
	out := map[int]string{}
	for _, path := range paths {
		var file struct {
			Spells map[string]struct {
				Icon string `json:"icon"`
			} `json:"spells"`
		}
		if err := readJSON(path, &file); err != nil {
			return nil, err
		}
		for k, tip := range file.Spells {
			if id, err := strconv.Atoi(k); err == nil && tip.Icon != "" {
				out[id] = tip.Icon
			}
		}
	}
	return out, nil
}

func talentIcons(dir string) (map[int]string, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, err
	}
	out := map[int]string{}
	for _, path := range paths {
		var file struct {
			Trees []struct {
				Talents []struct {
					Icon  string `json:"icon"`
					Ranks []struct {
						SpellID int `json:"spell_id"`
					} `json:"ranks"`
				} `json:"talents"`
			} `json:"trees"`
		}
		if err := readJSON(path, &file); err != nil {
			return nil, err
		}
		for _, tree := range file.Trees {
			for _, t := range tree.Talents {
				for _, r := range t.Ranks {
					out[r.SpellID] = t.Icon
				}
			}
		}
	}
	return out, nil
}
