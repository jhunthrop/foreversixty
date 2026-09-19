package request

// The engine's EquipmentSpec.Items is a positional array, not a map:
// slot order is the contract. These names are the planner's, and the
// index is the engine's ItemSlot enum value. A mismatch equips a helm in
// the weapon slot and nothing complains, so slotOrder is checked against
// proto.ItemSlot by the test rather than trusted.
var slotOrder = []string{
	"head",
	"neck",
	"shoulder",
	"back",
	"chest",
	"wrist",
	"hands",
	"waist",
	"legs",
	"feet",
	"finger1",
	"finger2",
	"trinket1",
	"trinket2",
	"main_hand",
	"off_hand",
	"ranged",
}

// SlotCount is how many slots an EquipmentSpec always carries. Unfilled
// slots are present and empty, because the engine indexes the array
// rather than searching it.
var SlotCount = len(slotOrder)

var slotIndex = func() map[string]int {
	m := make(map[string]int, len(slotOrder))
	for i, s := range slotOrder {
		m[s] = i
	}
	return m
}()

// SlotIndex maps a planner slot name to its engine position.
func SlotIndex(name string) (int, bool) {
	i, ok := slotIndex[name]
	return i, ok
}
