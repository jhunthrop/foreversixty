package rating

// The casts rotation adherence does not count as ability casts (adherence.go), in one place:
// every consumable in the engine's consumables catalogue
// (logs/engine/mechanics/consumables/catalogue.json, the catalogue score_v2.go's discipline
// reads), and the small curated list below, by category. An ability the rotation lists is
// always counted, whatever its category here (a shout or a form the rotation asks for is part
// of the rotation).
//
// Entries match a cast row by spell id or by the log's own spell name; names cover every rank
// of a family and every item that casts the same effect.

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// adherenceExclusion is one category of casts that are not ability casts.
type adherenceExclusion struct {
	category string
	ids      []int64
	names    []string
}

var adherenceExclusions = []adherenceExclusion{
	{
		// Melee and ranged white hits: the swing timer, not a choice.
		category: "auto-attack",
		ids:      []int64{6603, 75, 5019, 3018, 2480, 7918, 7919, 2764},
		names:    []string{"Attack", "Auto Shot", "Shoot", "Shoot Bow", "Shoot Gun", "Shoot Crossbow", "Throw"},
	},
	{
		// Potions, runes, healthstones and bandages beyond the catalogue's own entries.
		category: "consumable",
		names: []string{"Healing Potion", "Restore Mana", "Mana Potion", "Rejuvenation Potion", "Demonic Rune",
			"Dark Rune", "Healthstone", "Minor Healthstone", "Lesser Healthstone", "Greater Healthstone",
			"Major Healthstone", "First Aid", "Free Action", "Living Free Action", "Invulnerability",
			"Restoration", "Fire Protection", "Frost Protection", "Nature Protection", "Shadow Protection",
			"Arcane Protection", "Holy Protection"},
	},
	{
		// Class mounts; an item mount in a boss fight is rare enough to leave unlisted.
		category: "mount",
		ids:      []int64{13819, 23214, 5784, 23161},
		names:    []string{"Summon Warhorse", "Summon Charger", "Summon Felsteed", "Summon Dreadsteed"},
	},
	{
		category: "movement",
		names: []string{"Sprint", "Dash", "Blink", "Aspect of the Cheetah", "Aspect of the Pack", "Travel Form",
			"Aquatic Form", "Ghost Wolf", "Slow Fall", "Levitate"},
	},
	{
		// Out-of-rotation utility: travel, gathering, conjuring and the long raid buffs cast
		// once before the pull.
		category: "utility",
		names: []string{"Hearthstone", "Fishing", "Find Herbs", "Find Minerals", "Find Treasure",
			"Conjure Water", "Conjure Food", "Ritual of Summoning", "Arcane Intellect", "Arcane Brilliance",
			"Power Word: Fortitude", "Prayer of Fortitude", "Divine Spirit", "Prayer of Spirit",
			"Prayer of Shadow Protection", "Mark of the Wild", "Gift of the Wild", "Blessing of Kings",
			"Blessing of Might", "Blessing of Wisdom", "Blessing of Salvation", "Blessing of Light",
			"Blessing of Sanctuary", "Greater Blessing of Kings", "Greater Blessing of Might",
			"Greater Blessing of Wisdom", "Greater Blessing of Salvation", "Greater Blessing of Light",
			"Greater Blessing of Sanctuary"},
	},
}

// adherenceExcludedIDs and adherenceExcludedNames are adherenceExclusions and every entry of
// every category of the consumables catalogue, across all roles.
var adherenceExcludedIDs, adherenceExcludedNames = adherenceExclusionSets()

func adherenceExclusionSets() (map[int64]bool, map[string]bool) {
	ids, names := map[int64]bool{}, map[string]bool{}
	for _, e := range adherenceExclusions {
		for _, id := range e.ids {
			ids[id] = true
		}
		for _, n := range e.names {
			names[n] = true
		}
	}
	for _, rc := range consumables.Load().Roles {
		for _, list := range [][]consumables.Entry{rc.Flask, rc.Food, rc.WeaponEnchant, rc.WorldBuffs, rc.CombatPotion.Entries} {
			for _, e := range list {
				ids[e.SpellID] = true
				if e.Name != "" {
					names[e.Name] = true
				}
			}
		}
	}
	return ids, names
}

// excludedFromAdherence reports whether c is not an ability cast.
func excludedFromAdherence(c summary.CastRow) bool {
	return adherenceExcludedIDs[c.SpellID] || adherenceExcludedNames[c.SpellName]
}
