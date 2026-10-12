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
//
// Item effects (trinkets, potions, bandages, racials, mounts, Hearthstone, professions) are
// handled by the no-class-spell rule above, never by a name here. The curated names are only
// class families, and TestEveryCuratedExclusionNameIsAClassFamily holds that against the
// newest data build: a name that is no class's family could never match anything. The
// warlock's Healthstones and Spellstones are the one consumable that is a class family, so
// they are named below.

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
			"Abolish Poison", "Cure Poison", "Abolish Disease", "Cure Disease", "Purge", "Tranquilizing Shot"},
	},
	{
		category: "interrupt",
		names:    []string{"Pummel", "Shield Bash", "Kick", "Counterspell"},
	},
	{
		category: "taunt and threat",
		names: []string{"Taunt", "Mocking Blow", "Challenging Shout", "Growl", "Challenging Roar",
			"Hand of Reckoning", "Feign Death", "Fade", "Vanish", "Cower", "Feint"},
	},
	{
		category: "resurrection",
		names:    []string{"Resurrection", "Redemption", "Ancestral Spirit", "Rebirth", "Soulstone Resurrection"},
	},
	{
		// Self-buffs, stances, forms, seals (and the judgement a seal turns into), aspects,
		// armors, shields, shouts and totems a rotation does not list.
		category: "self-buff, stance, form, seal, aspect, armor, totem",
		names: []string{"Frost Armor", "Ice Armor", "Mage Armor", "Molten Armor", "Demon Armor",
			"Demon Skin", "Inner Fire", "Lightning Shield", "Water Shield", "Earth Shield", "Righteous Fury",
			"Thorns", "Omen of Clarity", "Battle Shout", "Commanding Shout", "Mana Shield"},
		prefixes: []string{"Seal of ", "Judgement of ", "Aspect of the "},
		suffixes: []string{" Form", " Stance", " Totem", " Aura"},
	},
	{
		// The one racial a class's own spell list carries (the rogue's); every other racial is
		// no class spell.
		category: "racial",
		names:    []string{"Shadowmeld"},
	},
	{
		// Warlock Healthstones and Spellstones: consumables, but warlock families.
		category: "consumable",
		names: []string{"Healthstone", "Minor Healthstone", "Lesser Healthstone", "Greater Healthstone",
			"Major Healthstone", "Spellstone", "Greater Spellstone", "Major Spellstone"},
	},
	{
		category: "movement",
		names: []string{"Sprint", "Blink", "Ghost Wolf", "Slow Fall", "Levitate", "Charge", "Intercept",
			"Feral Charge"},
	},
	{
		// Conjuring, summoning, tracking, pets and the long raid buffs cast before the pull.
		category: "utility",
		names: []string{"Arcane Intellect", "Arcane Brilliance", "Power Word: Fortitude", "Prayer of Fortitude",
			"Divine Spirit", "Prayer of Spirit", "Shadow Protection", "Prayer of Shadow Protection",
			"Mark of the Wild", "Gift of the Wild", "Amplify Magic", "Dampen Magic", "Revive Pet", "Mend Pet",
			"Eyes of the Beast", "Disarm Trap"},
		prefixes: []string{"Blessing of ", "Greater Blessing of ", "Conjure ", "Create ", "Ritual of ",
			"Summon ", "Track "},
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
