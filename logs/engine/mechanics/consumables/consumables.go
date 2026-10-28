// logs/engine/mechanics/consumables/consumables.go
// Package consumables holds the curated raid-consumable catalogue by role
// (docs/superpowers/specs/2026-09-21-performance-rating-design.md §3.3),
// for the rating engine's Preparation component, and the raid buffs with
// the class that provides each. AtPull (atpull.go) reads both against a
// fight's seeded aura tracks (docs/superpowers/specs/2026-10-27-log-coach-design.md
// §13). Same embed-parse-validate pattern as logs/engine/mechanics and
// logs/engine/mechanics/utility.
package consumables

import (
	"embed"
	"encoding/json"
	"fmt"
)

// The five consumable categories, the keys of RoleCatalogue.Weights and
// RoleCatalogue.Complete and of Pull.Present.
const (
	CategoryFlask         = "flask"
	CategoryFood          = "food"
	CategoryWeaponEnchant = "weapon_enchant"
	CategoryWorldBuffs    = "world_buffs"
	CategoryCombatPotion  = "combat_potion"
)

// categories is every category in a fixed order.
var categories = []string{CategoryFlask, CategoryFood, CategoryWeaponEnchant, CategoryWorldBuffs, CategoryCombatPotion}

func isCategory(s string) bool {
	for _, c := range categories {
		if c == s {
			return true
		}
	}
	return false
}

// Entry is one consumable: a persistent buff or a combat potion, matched
// against a player's seeded aura tracks or Casts.
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
// against MaxUses. Complete says, per category, whether the list is every
// entry a player could carry: only a complete list may call a player
// missing that category.
type RoleCatalogue struct {
	Weights       map[string]int  `json:"weights"`
	Complete      map[string]bool `json:"complete"`
	Flask         []Entry         `json:"flask"`
	Food          []Entry         `json:"food"`
	WeaponEnchant []Entry         `json:"weapon_enchant"`
	WorldBuffs    []Entry         `json:"world_buffs"`
	CombatPotion  PotionGroup     `json:"combat_potion"`
}

// IsComplete reports whether the category's list is complete. An unknown
// or unlisted category is not.
func (rc RoleCatalogue) IsComplete(category string) bool { return rc.Complete[category] }

// Entries is the category's listed entries, nil for an unknown category.
func (rc RoleCatalogue) Entries(category string) []Entry {
	switch category {
	case CategoryFlask:
		return rc.Flask
	case CategoryFood:
		return rc.Food
	case CategoryWeaponEnchant:
		return rc.WeaponEnchant
	case CategoryWorldBuffs:
		return rc.WorldBuffs
	case CategoryCombatPotion:
		return rc.CombatPotion.Entries
	}
	return nil
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

// RaidBuff is one raid buff and the class that provides it; SpellIDs is every rank.
type RaidBuff struct {
	Name     string  `json:"name"`
	Class    string  `json:"class"` // the log's class name, as summary.RosterRow.Class ("Priest")
	SpellIDs []int64 `json:"spell_ids"`
	Verified string  `json:"verified"`
}

// Catalogue is the whole consumable table, one row per role, and the raid
// buffs.
type Catalogue struct {
	Roles     map[string]RoleCatalogue `json:"roles"`
	RaidBuffs []RaidBuff               `json:"raid_buffs"`
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
// 100, whose entries carry no positive spell id or no verified source,
// whose complete flags name an unknown category, or whose raid buffs carry
// no class, no positive spell ids or no verified source.
func Parse(data []byte) (Catalogue, error) {
	var c Catalogue
	if err := json.Unmarshal(data, &c); err != nil {
		return Catalogue{}, fmt.Errorf("consumables: %w", err)
	}
	for role, rc := range c.Roles {
		if err := checkRole(role, rc); err != nil {
			return Catalogue{}, err
		}
	}
	for _, b := range c.RaidBuffs {
		if err := checkRaidBuff(b); err != nil {
			return Catalogue{}, err
		}
	}
	return c, nil
}

func checkRole(role string, rc RoleCatalogue) error {
	if sum := rc.weightSum(); sum != weightsTotal {
		return fmt.Errorf("consumables: %s weights sum to %d, want %d", role, sum, weightsTotal)
	}
	for key := range rc.Complete {
		if !isCategory(key) {
			return fmt.Errorf("consumables: %s: complete names unknown category %q", role, key)
		}
	}
	for _, cat := range categories {
		for _, e := range rc.Entries(cat) {
			if e.SpellID <= 0 {
				return fmt.Errorf("consumables: %s: %q has no positive spell_id", role, e.Name)
			}
			if e.Verified == "" {
				return fmt.Errorf("consumables: %s: %q has no verified source", role, e.Name)
			}
		}
	}
	return nil
}

func checkRaidBuff(b RaidBuff) error {
	if b.Class == "" {
		return fmt.Errorf("consumables: raid buff %q has no class", b.Name)
	}
	if len(b.SpellIDs) == 0 {
		return fmt.Errorf("consumables: raid buff %q has no spell_ids", b.Name)
	}
	for _, id := range b.SpellIDs {
		if id <= 0 {
			return fmt.Errorf("consumables: raid buff %q has a non-positive spell id %d", b.Name, id)
		}
	}
	if b.Verified == "" {
		return fmt.Errorf("consumables: raid buff %q has no verified source", b.Name)
	}
	return nil
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
