// Package leveling: the per-class kit every simulated leveling character
// carries besides gear and talents.
package leveling

import "strings"

// KitConsumes is the exception to "a leveling character carries no
// consumables": a handful of specs have a consumable so core to how
// they play that leaving it off sims a rotation no real player runs.
// It takes the full spec slug ("rogue-assassination",
// "shaman-enhancement"), not the bare class, because the kit is a
// spec property: shaman-elemental and shaman-restoration are not
// weapon-imbue specs the way shaman-enhancement is, even though all
// three share the shaman class.
//
//   - Any rogue spec, from level 20 (the poison quest): Instant
//     Poison on both weapons. Assassination's whole design (Mutilate's
//     bonus on a poisoned target, Venom) assumes them.
//   - Shaman-enhancement: Rockbiter Weapon on both weapons from level
//     1 (its earliest rank is learnable at level 1, well before
//     Windfury Weapon exists); Windfury Weapon replaces it on the main
//     hand from level 30, when its first rank is learned, since a
//     dual-Windfury enhancement shaman shares one internal cooldown
//     across both weapons in this engine (sim/shaman/windfury_weapon.go)
//     and gains nothing keeping Rockbiter off the off hand. The spec's
//     entire identity is built on weapon imbues the way Stormstrike
//     itself is; a bare-weapon Enhancement shaman sims a rotation no
//     real player runs, the same failure mode a poison-less rogue was.
//
// Flasks, food, oils and potions stay off. The ladder and the leveling
// BiS ranker both read this so their numbers agree.
func KitConsumes(spec string, level int) []string {
	switch {
	case isRogueSpec(spec):
		if level < 20 {
			return nil
		}
		return []string{"main_hand_imbue:instant_poison", "off_hand_imbue:instant_poison"}
	case spec == "shaman-enhancement":
		if level < 30 {
			return []string{"main_hand_imbue:rockbiter_weapon", "off_hand_imbue:rockbiter_weapon"}
		}
		return []string{"main_hand_imbue:windfury_weapon", "off_hand_imbue:rockbiter_weapon"}
	default:
		return nil
	}
}

// isRogueSpec matches the bare class ("rogue", used where a caller has
// no spec to name) as well as any of its three specs.
func isRogueSpec(spec string) bool {
	return spec == "rogue" || strings.HasPrefix(spec, "rogue-")
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
