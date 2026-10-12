package rating

// The casts rotation adherence does not count as ability casts (adherence.go), in one place.
// A cast of an ability the band lists always counts first, whatever its category here (a
// shout, a totem or a seal the rotation asks for is part of the rotation). Otherwise a cast is
// set aside when:
//
//   - its spell is in the engine's consumables catalogue
//     (logs/engine/mechanics/consumables/catalogue.json, the catalogue score_v2.go's
//     discipline reads), every category, every role; or
//   - its spell is no rank of any of the class's own spells (sims.Rotation.Family, from the
//     data build's spellranks.json): an item or trinket on-use, a racial, a potion, a bandage,
//     an item mount or a profession. This is the item, racial and consumable category, and it
//     needs no list of names; or
//   - its spellranks.json family name (the data's own ability name, never the log's, so a
//     localized log matches the same) is one the curated categories below name, exactly or
//     by a prefix or suffix ("Seal of ", " Totem").
//
// Overlaps worth knowing: "Shadow Protection" is the priest buff (utility below) and also the
// effect name of the Shadow Protection Potion; the potion is no priest spell, so the item rule
// sets it aside, and the name here only ever matches the priest's own family.

import (
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// adherenceExclusion is one category of class spells that are not ability casts, matched on
// the spellranks.json family name.
type adherenceExclusion struct {
	category string
	names    []string
	prefixes []string
	suffixes []string
}

var adherenceExclusions = []adherenceExclusion{
	{
		// Ranged white hits a class trains (melee swings and wands are no class spell).
		category: "auto-attack",
		names:    []string{"Auto Shot", "Shoot", "Throw"},
	},
	{
		category: "dispel and cleanse",
		names: []string{"Dispel Magic", "Cleanse", "Purify", "Remove Curse", "Remove Lesser Curse",
			"Abolish Poison", "Cure Poison", "Abolish Disease", "Cure Disease", "Purge"},
	},
	{
		category: "interrupt",
		names:    []string{"Pummel", "Shield Bash", "Kick", "Counterspell", "Silence"},
	},
	{
		category: "taunt and threat",
		names: []string{"Taunt", "Mocking Blow", "Challenging Shout", "Growl", "Challenging Roar",
			"Hand of Reckoning", "Righteous Defense", "Feign Death", "Fade", "Vanish", "Cower", "Feint",
			"Soulshatter"},
	},
	{
		category: "resurrection",
		names:    []string{"Resurrection", "Redemption", "Ancestral Spirit", "Rebirth", "Soulstone Resurrection"},
	},
	{
		// Self-buffs, stances, forms, seals (and the judgement a seal turns into), aspects,
		// armors, shields, shouts and totems a rotation does not list.
		category: "self-buff, stance, form, seal, aspect, armor, totem",
		names: []string{"Shadowform", "Frost Armor", "Ice Armor", "Mage Armor", "Molten Armor", "Demon Armor",
			"Demon Skin", "Inner Fire", "Lightning Shield", "Water Shield", "Earth Shield", "Righteous Fury",
			"Thorns", "Omen of Clarity", "Battle Shout", "Commanding Shout"},
		prefixes: []string{"Seal of ", "Judgement of ", "Aspect of the "},
		suffixes: []string{" Form", " Stance", " Totem", " Aura"},
	},
	{
		// Racials a class's own spell list happens to carry (the rest are no class spell).
		category: "racial",
		names: []string{"Blood Fury", "Berserking", "Stoneform", "Perception", "War Stomp", "Shadowmeld",
			"Will of the Forsaken", "Escape Artist", "Cannibalize", "Find Treasure"},
	},
	{
		category: "mount",
		names:    []string{"Summon Warhorse", "Summon Charger", "Summon Felsteed", "Summon Dreadsteed"},
	},
	{
		category: "movement",
		names: []string{"Sprint", "Dash", "Blink", "Ghost Wolf", "Slow Fall", "Levitate", "Charge", "Intercept",
			"Feral Charge"},
	},
	{
		// Travel, gathering, conjuring, summoning, pets and the long raid buffs cast before
		// the pull.
		category: "utility",
		names: []string{"Hearthstone", "Fishing", "Arcane Intellect", "Arcane Brilliance",
			"Power Word: Fortitude", "Prayer of Fortitude", "Divine Spirit", "Prayer of Spirit",
			"Shadow Protection", "Prayer of Shadow Protection", "Mark of the Wild", "Gift of the Wild",
			"Detect Magic", "Detect Invisibility", "Detect Lesser Invisibility", "Unending Breath",
			"Water Walking", "Water Breathing", "Amplify Magic", "Dampen Magic", "Call Pet", "Revive Pet",
			"Dismiss Pet", "Feed Pet", "Mend Pet", "Tame Beast", "Eyes of the Beast", "Eye of Kilrogg",
			"Far Sight", "Pick Lock", "Disarm Trap", "Sense Undead", "Sense Demons"},
		prefixes: []string{"Blessing of ", "Greater Blessing of ", "Conjure ", "Create ", "Ritual of ",
			"Summon ", "Track ", "Find "},
	},
}

// adherenceCatalogueIDs is every entry of every category of the consumables catalogue,
// across all roles.
var adherenceCatalogueIDs = catalogueSpellIDs()

// adherenceExcludedNames is every exact family name adherenceExclusions lists.
var adherenceExcludedNames = func() map[string]bool {
	names := map[string]bool{}
	for _, e := range adherenceExclusions {
		for _, n := range e.names {
			names[n] = true
		}
	}
	return names
}()

func catalogueSpellIDs() map[int64]bool {
	ids := map[int64]bool{}
	for _, rc := range consumables.Load().Roles {
		for _, list := range [][]consumables.Entry{rc.Flask, rc.Food, rc.WeaponEnchant, rc.WorldBuffs, rc.CombatPotion.Entries} {
			for _, e := range list {
				ids[e.SpellID] = true
			}
		}
	}
	return ids
}

// excludedFromAdherence reports whether c, a cast of no ability rot lists, is not an ability
// cast.
func excludedFromAdherence(c summary.CastRow, rot sims.Rotation) bool {
	if adherenceCatalogueIDs[c.SpellID] {
		return true
	}
	family, ok := rot.Family(c.SpellID)
	if !ok {
		return true // an item, trinket, racial or consumable effect: no class spell
	}
	return excludedFamily(family)
}

// excludedFamily reports whether a class spell family is one adherenceExclusions names.
func excludedFamily(family string) bool {
	if adherenceExcludedNames[family] {
		return true
	}
	for _, e := range adherenceExclusions {
		for _, p := range e.prefixes {
			if strings.HasPrefix(family, p) {
				return true
			}
		}
		for _, s := range e.suffixes {
			if strings.HasSuffix(family, s) {
				return true
			}
		}
	}
	return false
}
