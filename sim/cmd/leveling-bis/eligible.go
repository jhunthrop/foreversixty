package main

import (
	"fmt"
)

// Armor subclass ids, from the client's ItemSubclassArmor enum as
// items.json's subclass_id carries it (class_id 4 is armor). Confirmed
// against data/builds/1.60.1.70009/items.json: every hunter-equippable
// armor row is 0 (misc: rings, cloaks with no armor value on some
// builds), 1 (cloth), 2 (leather) or 3 (mail) - never 4 (plate), which
// the per-class file's own "equippable" filter already excludes. Mail
// is the only one this lane's rule gates further, by level.
const (
	armorClassID      = 4
	armorSubclothID   = 1
	armorSubleatherID = 2
	armorSubmailID    = 3
	armorSubplateID   = 4
)

// armorAvailableLevel is the level a class's armor proficiency for a
// subclass opens, keyed by "<class>:<subclass id>". Only entries a
// class does NOT have from level 1 need to be listed; an absent entry
// means "no gate", which is correct for cloth and leather on every
// class this file has an entry for.
//
// Hunter: leather until 40, mail from 40 (this lane's brief, restated
// from the leveling-bis design doc's own proficiency table). Plate
// never appears in hunter's per-class item file at all (see the
// constant block above), so it needs no entry here.
var armorAvailableLevel = map[string]int{
	"hunter:3": 40, // mail
}

// eligible reports whether a leveling character of class/level/faction
// could equip this candidate at all, per this lane's brief:
//
//   - EffectiveRequiredLevel <= level (2026-09-28 quest-levels lane:
//     NOT c.RequiredLevel directly - a quest reward or crafted item's
//     own required_level is very often 0 in the client, gated instead
//     by the quest's or recipe's own level; applyEffectiveRequiredLevels
//     (data.go) resolves the real gate once per run, from loot.json's
//     quests map, before eligible() ever runs).
//   - class allowed: NOT checked here - candidates are already sourced
//     from data/builds/<build>/items/<class>.json, the pipeline's own
//     "per-class equippable items" file (data/README.md), so a
//     candidate this function ever sees already passed the class mask.
//   - armor proficiency by level, from armorAvailableLevel above.
//   - faction restriction: the item's own faction_restriction (from
//     items.json, already resolved per curated fact, per the design
//     doc's "the item's faction restriction carries the quest's
//     side") must be empty or match.
//
// Source-based exclusions (raid items below 60, no-known-source items
// unless a zone drop) are NOT this function's job: they need
// lootIndex, which is a per-band, not a per-item, question of "is
// this item's source usable at this band" as much as "does this item
// exist" - see band.go's candidatesForBand, which calls this function
// per item and then applies the source rule once per band.
func eligible(c candidate, classSlug string, level int, faction string) bool {
	if c.EffectiveRequiredLevel > level {
		return false
	}
	if c.ClassID == armorClassID {
		key := fmt.Sprintf("%s:%d", classSlug, c.SubclassID)
		if opens, gated := armorAvailableLevel[key]; gated && level < opens {
			return false
		}
	}
	if c.FactionRestriction != "" && c.FactionRestriction != faction {
		return false
	}
	return true
}
