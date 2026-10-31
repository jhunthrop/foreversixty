package upgrades

// Slot is one of the character frame's 17 gear slots: the BiS data's key and its words.
type Slot struct {
	Key, Label string
}

// Slots are the 17 gear slots in character-frame order.
var Slots = []Slot{
	{"head", "Head"}, {"neck", "Neck"}, {"shoulder", "Shoulder"}, {"back", "Back"},
	{"chest", "Chest"}, {"wrist", "Wrist"}, {"hands", "Hands"}, {"waist", "Waist"},
	{"legs", "Legs"}, {"feet", "Feet"}, {"finger1", "Finger 1"}, {"finger2", "Finger 2"},
	{"trinket1", "Trinket 1"}, {"trinket2", "Trinket 2"}, {"main_hand", "Main hand"},
	{"off_hand", "Off hand"}, {"ranged", "Ranged"},
}

// weaponSlots are the slots whose gain no stat-weight score measures: a weapon's score bakes
// in its damage, which an item's stat line does not carry.
var weaponSlots = map[string]bool{"main_hand": true, "off_hand": true, "ranged": true}

// SlotLabel is a slot key's words ("finger1" is "Finger 1"); an unknown key is itself.
func SlotLabel(key string) string {
	for _, s := range Slots {
		if s.Key == key {
			return s.Label
		}
	}
	return key
}
