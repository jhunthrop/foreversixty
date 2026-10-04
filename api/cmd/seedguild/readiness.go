// api/cmd/seedguild/readiness.go
//
// The readiness variety the coordinator's raid-control-centre scope note asked for: most
// raiders fully BiS-geared, a few with real (not invented) alternative items instead of
// the top BiS pick, two missing one enchant, two with unspent talent points, professions
// for everyone, and a consumables bags= section for the one gear_bags-consent raider.
// Every item id (gear alternatives, enchants, consumables) is a real id read out of this
// repo's own data tables — never invented — matching the seed's own gear rule.
package main

import "sort"

// enchantCatalogue is a small, real subset of data/builds/<build>/enchants.json: one
// enchant id per equippable slot this roster's gear ever occupies, read by hand from
// that file (ids 247/724/856/etc. below) rather than invented. It is deliberately not
// "the BiS enchant" for every spec — a generic, broadly-appropriate pick is enough to
// make the export's gear=slot:enchant grammar real and decodable; the readiness board
// this unlocks (design spec 2026-10-04-guild-page.md §9) is about whether a slot HAS an
// enchant at all, not which one.
var enchantCatalogue = map[string]int{
	"back":      247,  // Enchant Cloak - Minor Agility
	"chest":     41,   // Enchant Chest - Minor Health
	"wrist":     724,  // Enchant Bracer - Lesser Stamina
	"hands":     856,  // Enchant Gloves - Strength
	"legs":      1505, // Lesser Arcanum of Resilience
	"feet":      724,  // Enchant Boots - Lesser Stamina
	"main_hand": 250,  // Enchant Weapon - Minor Striking
	"off_hand":  250,
	"ranged":    32, // Standard Scope
}

// professionsCatalogue assigns each class a flavour-appropriate gathering/crafting pair -
// slugs only (professions= carries no item ids, so there is nothing to invent here).
var professionsCatalogue = map[string][]string{
	"warrior": {"mining", "blacksmithing"},
	"paladin": {"mining", "blacksmithing"},
	"hunter":  {"skinning", "leatherworking"},
	"rogue":   {"herbalism", "alchemy"},
	"priest":  {"tailoring", "enchanting"},
	"shaman":  {"skinning", "leatherworking"},
	"mage":    {"tailoring", "enchanting"},
	"warlock": {"herbalism", "alchemy"},
	"druid":   {"herbalism", "alchemy"},
}

// professionsFor is the class's own profession pair, plus first aid for every raider
// (every one of this roster is "professions set for everyone," per the seed's own rule).
func professionsFor(class string) []string {
	base := professionsCatalogue[class]
	out := make([]string, 0, len(base)+1)
	out = append(out, base...)
	out = append(out, "first-aid")
	return out
}

// bagsConsumables is the real, level-60-raid consumable item ids
// data/builds/<build>/simconsumes.json carries for the gear_bags-consent raider: a flask,
// a potion, and a bandage, in a fixed order.
var bagsConsumables = []int{
	13510,  // Flask of the Titans
	13444,  // Major Mana Potion
	232433, // Dense Runecloth Bandage
}

// bagsFor is c's bags= contents: the shared consumables list when consent is
// "gear_bags" (the one consent level the roster's Guild.svelte, per the guild membership
// design, treats as "sees bags/bank too"), nil otherwise.
func bagsFor(c mockCharacter) []int {
	if c.Consent != "gear_bags" {
		return nil
	}
	return bagsConsumables
}

// gearWithReadiness turns a character's raw BiS slots (gear.go's loadBisSlots) into the
// gearPiece map buildFS1 wants, applying this character's own readiness gaps:
//   - NonBisSlots of its slots (in slot-name order, for determinism) take that slot's
//     first real alternative item instead of the top BiS pick, when one exists.
//   - every enchantable slot it has gets enchantCatalogue's real enchant id, except
//     MissingEnchantSlot, which carries none.
func gearWithReadiness(c mockCharacter, slots []bisSlot) map[string]gearPiece {
	bySlot := make(map[string]bisSlot, len(slots))
	order := make([]string, 0, len(slots))
	for _, s := range slots {
		if s.ItemID <= 0 {
			continue // the BiS table's own "nothing equipped here" sentinel
		}
		bySlot[s.Slot] = s
		order = append(order, s.Slot)
	}
	sort.Strings(order)

	swapped := map[string]bool{}
	remaining := c.NonBisSlots
	for _, slot := range order {
		if remaining <= 0 {
			break
		}
		if alts := bySlot[slot].Alternatives; len(alts) > 0 && alts[0].ItemID > 0 {
			swapped[slot] = true
			remaining--
		}
	}

	gear := make(map[string]gearPiece, len(order))
	for _, slot := range order {
		itemID := bySlot[slot].ItemID
		if swapped[slot] {
			itemID = bySlot[slot].Alternatives[0].ItemID
		}
		piece := gearPiece{ItemID: itemID}
		if enchant, ok := enchantCatalogue[slot]; ok && slot != c.MissingEnchantSlot {
			piece.Enchant = enchant
		}
		gear[slot] = piece
	}
	return gear
}
