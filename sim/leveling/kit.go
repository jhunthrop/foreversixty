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
//     Poison on both weapons, and from level 30, when Deadly Poison's
//     first rank is learned, Deadly Poison on the main hand and Instant
//     Poison on the off hand. Mutilate's "+20% against poisoned targets"
//     counts only a poison that lingers (Deadly or Wound), not Instant,
//     so Assassination needs a Deadly hand (2026-10-07; the rogue
//     curation lane measured the Deadly main hand about 17% above two
//     Instants). The rogue-parity lane measured the four arrangements for
//     all three specs (design/reviews/2026-10-07-rogue-parity.md): one
//     Deadly beats two Instants by 4 to 10 percent, two Deadlys lose
//     most of that, and the main hand is the stronger Deadly hand for
//     Assassination and Combat at every band and for Subtlety below 60
//     (at 60 Subtlety's mirror is 0.9 percent ahead; one rule is kept).
//     Windfury Totem is a raid-buff aura, not an imbue, so a rogue's
//     poisons never compete with it.
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
		if level >= deadlyPoisonLevel {
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

// Buff lines of the self-buff kit, in the request vocabulary
// (sim/request/buffs.go).
const (
	arcaneBrillianceBuff = "arcane_brilliance"
	giftOfTheWildBuff    = "gift_of_the_wild"
	blessingOfMightBuff  = "blessing_of_might"
	huntersMarkBuff      = "hunters_mark"
)

// blessingOfMightLevel is the level Blessing of Might's first rank is
// learned. Arcane Intellect and Mark of the Wild are rank 1 at level 1.
const blessingOfMightLevel = 4

// huntersMarkLevel is the level Hunter's Mark's first rank is learned
// (client spell 1130).
const huntersMarkLevel = 6

// KitBuffs is the buff counterpart of KitConsumes: the self-cast buffs a
// class carries on every real run. A leveling character is "bare" - no
// raid buffs - but the engine models three class spells as raid buffs
// rather than castable spells, so a mage, a druid and a paladin would
// otherwise sim without their own Arcane Intellect, Mark of the Wild
// and Blessing of Might, a character no real player runs (a Retribution
// number without Blessing of Might above all). Like KitConsumes it
// takes the full spec slug and decides by spec, and the ladder, the
// leveling BiS ranker (its stat weights included), rotation-search and
// talent-search all read it so their numbers agree. The engine applies
// the highest rank the character's level can learn.
//
//   - Any mage spec, from level 1: Arcane Intellect (arcane_brilliance
//     is the engine's field for it; Improved Arcane Intellect does not
//     exist in Forever's Arcane tree, so there is no improved form).
//
//   - Any druid spec, from level 1: Mark of the Wild (gift_of_the_wild).
//     The ":improved" form is carried by a spec only when its guide build
//     takes Improved Mark of the Wild. Forever's Restoration tree has no
//     such talent (it became a baseline passive, of a rank the client
//     tables do not state) and neither the Feral nor the Balance build
//     can take it, so every druid spec carries the plain buff.
//
//   - Any paladin spec, from level 4 (Blessing of Might's first rank):
//     Blessing of Might. Forever's Holy tree has no Improved Blessing of
//     Might, so no ":improved" form.
//
//   - hunter-survival, from level 6 (Hunter's Mark's first rank): Hunter's
//     Mark, a target debuff the engine takes as a request buff. Expose
//     Prey (2026-10-07, design/reviews/2026-10-07-mongoose-bite.md) opens
//     Mongoose Bite's window only against a marked target, and a
//     Survival hunter marks his target before anything else.
//
// Priest Power Word: Fortitude is stamina and Divine Spirit a
// Discipline talent, so priests carry nothing; Inner Fire (armor only)
// and Lightning Shield (reactive) are not modelled as buffs.
func KitBuffs(spec string, level int) []string {
	switch {
	case isClassSpec(spec, "mage"):
		return []string{arcaneBrillianceBuff}
	case isClassSpec(spec, "druid"):
		return []string{giftOfTheWildBuff}
	case isClassSpec(spec, "paladin"):
		if level < blessingOfMightLevel {
			return nil
		}
		return []string{blessingOfMightBuff}
	case spec == "hunter-survival":
		if level < huntersMarkLevel {
			return nil
		}
		return []string{huntersMarkBuff}
	default:
		return nil
	}
}

// isRogueSpec matches the bare class ("rogue", used where a caller has
// no spec to name) as well as any of its three specs.
func isRogueSpec(spec string) bool { return isClassSpec(spec, "rogue") }

// isClassSpec matches the bare class slug as well as any of its specs.
func isClassSpec(spec, class string) bool {
	return spec == class || strings.HasPrefix(spec, class+"-")
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

// NoWeaponImbueSpecs fight in a shapeshift form, where a weapon imbue
// (a stone, an oil, a poison) does nothing: the form's claws swing, not
// the weapon. The preset resolver drops every imbue for them so a stone's
// crit never reaches a cat.
var NoWeaponImbueSpecs = map[string]bool{
	"druid-feral": true,
}
