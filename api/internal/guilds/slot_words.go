// api/internal/guilds/slot_words.go
//
// A slot in the client's own tooltip words ("Main Hand", "Chest", "Back", "Feet"), the words the
// site prints. clientSlotWords mirrors web/src/lib/items/slot-words.ts SLOT_WORDS exactly
// (slot_words_test.go holds the two together), so the bot and the site never name one slot two
// ways (guild round 3: /readiness said "Weapon, Chest, Cloak, Boots" where the site said "Main
// Hand, Chest and 2 more slots").
package guilds

// clientSlotWords is Item.InventoryType -> the tooltip's slot word (slot-words.ts SLOT_WORDS).
var clientSlotWords = map[int]string{
	1: "Head", 2: "Neck", 3: "Shoulder", 4: "Shirt", 5: "Chest", 6: "Waist", 7: "Legs", 8: "Feet",
	9: "Wrist", 10: "Hands", 11: "Finger", 12: "Trinket", 13: "One-Hand", 14: "Off Hand", 15: "Ranged",
	16: "Back", 17: "Two-Hand", 19: "Tabard", 20: "Chest", 21: "Main Hand", 22: "Off Hand",
	23: "Held In Off-hand", 25: "Thrown", 26: "Ranged", 28: "Relic",
}

// enchantableSlotInventoryTypes is each enchantable slot's client inventory type, the one whose
// word the slot prints: the main hand's INVTYPE_WEAPONMAINHAND, the chest's, the cloak's and the
// feet's.
var enchantableSlotInventoryTypes = map[string]int{"main_hand": 21, "chest": 5, "back": 16, "feet": 8}

// enchantableSlotLabels is each enchantable slot in the client's words.
var enchantableSlotLabels = func() map[string]string {
	out := make(map[string]string, len(enchantableSlotInventoryTypes))
	for slot, invType := range enchantableSlotInventoryTypes {
		out[slot] = clientSlotWords[invType]
	}
	return out
}()
