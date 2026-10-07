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
//     bonus on a poisoned target, Venom) assumes them - and Mutilate's
//     "+20% against poisoned targets" counts only a poison that lingers
//     (Deadly or Wound), not Instant, so Assassination carries Deadly
//     Poison on the main hand from level 30, when its first rank is
//     learned (2026-10-07; the rogue curation lane measured the Deadly
//     main hand about 17% above two Instants with the same ranking).
//     The off hand stays Instant.
//   - Shaman-enhancement: ONE weapon imbue, on the main hand only -
//     Rockbiter Weapon from level 1 (its earliest rank is learnable at
//     level 1, well before Windfury Weapon exists), replaced by
//     Windfury Weapon from level 30, when its first rank is learned.
//     Owner-stated rule (Justin, 2026-09-30, bis-ranker-integrity-13's
//     brief): shamans cannot dual wield in WoW Forever, so
//     enhancement's off hand is never a second weapon and never
//     carries a second imbue - the shaman-elemental/-restoration off
//     hand (a shield or a held item) has nothing to imbue either. A
//     bare main-hand-only Enhancement shaman still sims the rotation
//     a real Forever player runs; the spec's identity is Stormstrike
//     plus one imbued weapon, not two.
//
// Flasks, food, oils and potions stay off. The ladder and the leveling
// BiS ranker both read this so their numbers agree.
func KitConsumes(spec string, level int) []string {
	switch {
	case isRogueSpec(spec):
		if level < 20 {
			return nil
		}
		if spec == "rogue-assassination" && level >= deadlyPoisonLevel {
			return []string{"main_hand_imbue:deadly_poison", "off_hand_imbue:instant_poison"}
		}
		return []string{"main_hand_imbue:instant_poison", "off_hand_imbue:instant_poison"}
	case spec == "shaman-enhancement":
		if level < 30 {
			return []string{"main_hand_imbue:rockbiter_weapon"}
		}
		return []string{"main_hand_imbue:windfury_weapon"}
	default:
		return nil
	}
}

// deadlyPoisonLevel is the level Deadly Poison's first rank is learned.
const deadlyPoisonLevel = 30

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
//
// shaman-enhancement is deliberately absent: the owner's own rule (see
// KitConsumes above) is that shamans never dual wield in Forever, so
// enhancement's off hand draws from the same shield/held-item pool
// shaman-elemental already does, not this map's weapon pool. Rogue
// (all three specs) and hunter (all three specs) train dual wield from
// level 20 in Forever, same as warrior-fury; band 20 is the earliest
// band any of them can be offered a second weapon.
var DualWieldSpecs = map[string]bool{
	"rogue-assassination":  true,
	"rogue-combat":         true,
	"rogue-subtlety":       true,
	"warrior-fury":         true,
	"hunter-beast-mastery": true,
	"hunter-marksmanship":  true,
	"hunter-survival":      true,
}
