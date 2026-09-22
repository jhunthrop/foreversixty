// logs/engine/mechanics/consumables/consumables.go
// Package consumables holds the curated raid-consumable catalogue by role
// (docs/superpowers/specs/2026-09-21-performance-rating-design.md §3.3),
// for the rating engine's Preparation component. Same embed-parse-validate
// pattern as logs/engine/mechanics and logs/engine/mechanics/utility.
package consumables

import (
	"embed"
	"encoding/json"
	"fmt"
)

// Entry is one consumable: a persistent buff or a combat potion, matched
// against CombatantRow.Consumables or a player's Casts.
type Entry struct {
	SpellID  int64  `json:"spell_id"`
	Name     string `json:"name"`
	Verified string `json:"verified"`
}

// PotionGroup is combat_potion's shape: the entries that count, and the
// shared-cooldown cap on how many uses count toward completeness in one
// fight.
type PotionGroup struct {
	MaxUses int64   `json:"max_uses"`
	Entries []Entry `json:"entries"`
}

// RoleCatalogue is one role's consumable categories and their weights,
// each category's weight earned per the rating engine's Preparation
// component: flask/food/weapon_enchant credit in full if any one listed
// entry is present (a player carries exactly one of each), world_buffs
// credits proportionally to how many of the listed buffs are present
// (several stack at once), combat_potion credits proportionally to uses
// against MaxUses.
type RoleCatalogue struct {
	Weights       map[string]int `json:"weights"`
	Flask         []Entry        `json:"flask"`
	Food          []Entry        `json:"food"`
	WeaponEnchant []Entry        `json:"weapon_enchant"`
	WorldBuffs    []Entry        `json:"world_buffs"`
	CombatPotion  PotionGroup    `json:"combat_potion"`
}

// weightsTotal is what one role's category weights must sum to (spec
// §3.3: "each role's weights sum to 100").
const weightsTotal = 100

func (rc RoleCatalogue) weightSum() int {
	total := 0
	for _, w := range rc.Weights {
		total += w
	}
	return total
}

// Catalogue is the whole consumable table, one row per role.
type Catalogue struct {
	Roles map[string]RoleCatalogue `json:"roles"`
}

//go:embed catalogue.json
var catalogueFile embed.FS

// A malformed embedded catalogue is a build defect, matching
// logs/engine/mechanics's own panic-at-boot pattern.
var defaultCatalogue = mustParseEmbedded()

func mustParseEmbedded() Catalogue {
	data, err := catalogueFile.ReadFile("catalogue.json")
	if err != nil {
		panic(fmt.Errorf("consumables: reading catalogue.json: %w", err))
	}
	c, err := Parse(data)
	if err != nil {
		panic(fmt.Errorf("consumables: catalogue.json: %w", err))
	}
	return c
}

// Parse reads a catalogue and refuses one whose role weights do not sum to
// 100 or whose entries carry no positive spell id or no verified source.
func Parse(data []byte) (Catalogue, error) {
	var c Catalogue
	if err := json.Unmarshal(data, &c); err != nil {
		return Catalogue{}, fmt.Errorf("consumables: %w", err)
	}
	for role, rc := range c.Roles {
		if sum := rc.weightSum(); sum != weightsTotal {
			return Catalogue{}, fmt.Errorf("consumables: %s weights sum to %d, want %d", role, sum, weightsTotal)
		}
		all := make([]Entry, 0, 16)
		all = append(all, rc.Flask...)
		all = append(all, rc.Food...)
		all = append(all, rc.WeaponEnchant...)
		all = append(all, rc.WorldBuffs...)
		all = append(all, rc.CombatPotion.Entries...)
		for _, e := range all {
			if e.SpellID <= 0 {
				return Catalogue{}, fmt.Errorf("consumables: %s: %q has no positive spell_id", role, e.Name)
			}
			if e.Verified == "" {
				return Catalogue{}, fmt.Errorf("consumables: %s: %q has no verified source", role, e.Name)
			}
		}
	}
	return c, nil
}

// Load returns the embedded catalogue.
func Load() Catalogue { return defaultCatalogue }

// For selects a role's catalogue by the same role strings
// summary.RosterRow.Role already uses ("dps", "healer", "tank"). ok is
// false for an unknown role.
func (c Catalogue) For(role string) (RoleCatalogue, bool) {
	rc, ok := c.Roles[role]
	return rc, ok
}
