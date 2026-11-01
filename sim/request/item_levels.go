package request

// The level a character must be to use a consumable, from the client's item table.
//
// data/builds/<build>/items.json carries every ItemSparse row with its RequiredLevel. A
// preset consumable id names an engine enum value ("elixir_of_the_mongoose",
// "food_grilled_squid", "conjured_rogue_thistle_tea"), and the engine names its values
// after the items, so the join is by name: the id's value, or a tail of it past an enum
// prefix such as "food_" or "conjured_rogue_", compared with spaces and punctuation
// dropped ("elixir_of_firepower" is the client's "Elixir of Fire Power"). Where the client
// carries two items of one name (Forever added level-1 copies of Grilled Squid, Tender
// Wolf Steak and Elemental Sharpening Stone with no known source), the original, lowest
// item id counts. An id that names no item, such as a shaman's Windfury Weapon, has no
// gate here: the class kit gates its own spells.

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ItemLevels maps a compact item name onto the RequiredLevel of the lowest item id that
// carries it.
type ItemLevels map[string]itemLevel

type itemLevel struct {
	id    int64
	level int
}

// LoadItemLevelsFile reads a build's items.json.
func LoadItemLevelsFile(path string) (ItemLevels, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("request: reading the item table: %w", err)
	}
	var rows []struct {
		ID            int64  `json:"id"`
		Name          string `json:"name"`
		RequiredLevel int    `json:"required_level"`
	}
	if err := json.Unmarshal(raw, &rows); err != nil {
		return nil, fmt.Errorf("request: decoding %s: %w", path, err)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("request: %s carries no items", path)
	}
	levels := make(ItemLevels, len(rows))
	for _, row := range rows {
		key := compactName(row.Name)
		if prior, ok := levels[key]; ok && prior.id < row.ID {
			continue
		}
		levels[key] = itemLevel{id: row.ID, level: row.RequiredLevel}
	}
	return levels, nil
}

// RequiredLevel is the level the consumable id's item needs, and whether the id names an
// item at all.
func (l ItemLevels) RequiredLevel(id string) (int, bool) {
	_, value := splitConsumeID(id)
	tokens := strings.Split(value, "_")
	for start := range tokens {
		if item, ok := l[compactName(strings.Join(tokens[start:], " "))]; ok {
			return item.level, true
		}
	}
	return 0, false
}

// Usable says whether a character of level can use the consumable id.
func (l ItemLevels) Usable(id string, level int) bool {
	need, ok := l.RequiredLevel(id)
	return !ok || need <= level
}

// compactName is a name's letters and digits, lowercased: "Elixir of Fire Power" and
// "elixir_of_firepower" both read elixiroffirepower.
func compactName(name string) string {
	var b strings.Builder
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
