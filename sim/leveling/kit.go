// Package leveling: the per-class kit every simulated leveling character
// carries besides gear and talents.
package leveling

// KitConsumes is the one exception to "a leveling character carries no
// consumables": a rogue's poisons are class kit, taught by the level-20
// poison quest, and Assassination's whole design (Mutilate's bonus on a
// poisoned target, Venom) assumes them; a poison-less rogue sims a
// rotation no real rogue runs. Instant Poison on both weapons from 20.
// Flasks, food, oils and potions stay off. The ladder and the leveling
// BiS ranker both read this so their numbers agree.
func KitConsumes(class string, level int) []string {
	if class != "rogue" || level < 20 {
		return nil
	}
	return []string{"main_hand_imbue:instant_poison", "off_hand_imbue:instant_poison"}
}

// DualWieldSpecs names the specs whose off hand holds a WEAPON: for them
// a held-in-off-hand item or a shield in the off-hand candidate list is a
// ranking error (the level-20 assassination list once wore Grayson's
// Torch and swung one hand). Every other spec's off hand is a shield or a
// held item, and weapons are not offered there.
var DualWieldSpecs = map[string]bool{
	"rogue-assassination":  true,
	"rogue-combat":         true,
	"rogue-subtlety":       true,
	"warrior-fury":         true,
	"hunter-beast-mastery": true,
	"hunter-marksmanship":  true,
	"hunter-survival":      true,
	"shaman-enhancement":   true,
}
