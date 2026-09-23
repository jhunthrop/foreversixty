// api/internal/bnetbuild/gear.go
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type blizzardEquipment struct {
	EquippedItems []blizzardEquippedItem `json:"equipped_items"`
}

type blizzardEquippedItem struct {
	Item struct {
		ID int `json:"id"`
	} `json:"item"`
	Enchantments []struct {
		EnchantmentID   int `json:"enchantment_id"`
		EnchantmentSlot struct {
			Type string `json:"type"`
		} `json:"enchantment_slot"`
	} `json:"enchantments"`
	Slot struct {
		Type string `json:"type"`
	} `json:"slot"`
	Name string `json:"name"`
}

// bnetSlotToOurs maps Blizzard's equipped-item slot.type to this site's SLOTS names (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2). A slot not in this
// table (SHIRT, TABARD, and anything else the simulator does not model) is skipped and
// named in Report.SkippedSlots.
var bnetSlotToOurs = map[string]string{
	"HEAD": "head", "NECK": "neck", "SHOULDER": "shoulder", "BACK": "back", "CHEST": "chest",
	"WRIST": "wrist", "HANDS": "hands", "WAIST": "waist", "LEGS": "legs", "FEET": "feet",
	"FINGER_1": "finger1", "FINGER_2": "finger2", "TRINKET_1": "trinket1", "TRINKET_2": "trinket2",
	"MAIN_HAND": "main_hand", "OFF_HAND": "off_hand", "RANGED": "ranged",
}

// encodeGear builds the FS1 v1 gear list (spec §2.2, contract
// docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md §10.5:
// "<slot>=<item_id>[:<enchant[:<suffix>]]", in SLOTS order). suffixes may be empty; a miss
// is never reported (plan ruling 4 — no items.json eligibility table is available to
// distinguish "unknown suffix" from "never had one").
func encodeGear(raw json.RawMessage, suffixes SuffixTable) (gear string, skippedSlots, noSuffix []string, err error) {
	var body blizzardEquipment
	if err := json.Unmarshal(raw, &body); err != nil {
		return "", nil, nil, fmt.Errorf("bnetbuild: decode equipment: %w", err)
	}

	bySlot := make(map[string]blizzardEquippedItem, len(body.EquippedItems))
	for _, it := range body.EquippedItems {
		ourSlot, ok := bnetSlotToOurs[it.Slot.Type]
		if !ok {
			skippedSlots = append(skippedSlots, it.Slot.Type)
			continue
		}
		bySlot[ourSlot] = it
	}

	var parts []string
	for _, slot := range SLOTS {
		it, ok := bySlot[slot]
		if !ok {
			continue
		}
		entry := slot + "=" + strconv.Itoa(it.Item.ID)
		enchant := permanentEnchant(it)
		suffix, matched := matchSuffix(it.Name, suffixes)
		switch {
		case enchant != 0 && matched:
			entry += ":" + strconv.Itoa(enchant) + ":" + strconv.Itoa(suffix)
		case enchant != 0:
			entry += ":" + strconv.Itoa(enchant)
		case matched:
			entry += "::" + strconv.Itoa(suffix)
		}
		parts = append(parts, entry)
	}
	return strings.Join(parts, ","), skippedSlots, noSuffix, nil
}

// permanentEnchant returns the item's PERMANENT-slot enchantment_id, or 0 when it has none
// (spec §2.2: "the enchant is the PERMANENT enchantment_id when present").
func permanentEnchant(it blizzardEquippedItem) int {
	for _, e := range it.Enchantments {
		if e.EnchantmentSlot.Type == "PERMANENT" {
			return e.EnchantmentID
		}
	}
	return 0
}

// matchSuffix finds the longest suffix name in the table that name ends with, preceded by a
// space (plan ruling 4: an exact trailing match against real suffixes.json strings, not a
// heuristic on "of").
func matchSuffix(name string, suffixes SuffixTable) (id int, matched bool) {
	bestLen := -1
	for suffixName, suffixID := range suffixes {
		trailer := " " + suffixName
		if strings.HasSuffix(name, trailer) && len(suffixName) > bestLen {
			bestLen = len(suffixName)
			id = suffixID
			matched = true
		}
	}
	return id, matched
}
