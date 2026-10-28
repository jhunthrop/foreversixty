// logs/engine/rating/preparation.go
package rating

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// scorePreparation implements spec §1.3's Preparation component: a
// completeness score built from the consumable catalogue, percentile-
// within-bracket (or its own value as the absolute standard, RULING R5,
// since it is already a bounded 0-100 score). With no character sheet the
// pull is unknown, so the component is excluded rather than placed as an
// empty pull (no preparation digest fills with sheetless zeros).
func scorePreparation(fight summary.Summary, player string, bracket Bracket, cat *consumables.RoleCatalogue, src PercentileSource) Component {
	bracket.Component = ComponentNamePreparation
	c := Component{Name: ComponentNamePreparation}
	if cat == nil {
		c.Excluded, c.Reason = true, ReasonNoConsumableCatalogue
		return c
	}
	value, ok := preparationTotal(fight, player, cat)
	if !ok {
		c.Excluded, c.Reason = true, ReasonNoCharacterSheet
		return c
	}
	return placeBoundedOrSelf(c, src, bracket, value)
}

// preparationTotal is the raw 0-100 completeness for player, read through
// consumables.AtPull (the seeded aura tracks, or an aura applied within
// AtPullGraceMS of the pull; log coach spec §13). false when the log has no
// character sheet for player, or (unreachable on the shipped catalogue,
// whose flask and world_buffs lists are complete) when nothing on it can be
// judged: either way the pull is unknown, not empty.
//
// A complete list's category weighs in full whatever was read. An
// incomplete list can prove presence, never absence, so its category
// counts only what it earned, in the earned and the possible total alike:
// a read adds credit and never a shortfall, and bringing one more thing
// never lowers the score.
func preparationTotal(fight summary.Summary, player string, cat *consumables.RoleCatalogue) (float64, bool) {
	pull := consumables.AtPull(fight, player, *cat)
	if !pull.Known {
		return 0, false
	}
	var earned, possible float64
	for _, category := range weighedCategories {
		credit := categoryCredit(pull, cat, category)
		earned += credit
		if cat.IsComplete(category) {
			possible += float64(cat.Weights[category])
		} else {
			possible += credit
		}
	}
	if possible <= 0 {
		return 0, false
	}
	return clamp(100*earned/possible, 0, 100), true
}

// PreparationValue is the raw 0-100 completeness scorePreparation places for player; false
// with no catalogue or no character sheet, where the component is excluded instead. Exported
// so the rating rebuild can refold a preparation digest from stored summaries without
// scoring a card.
func PreparationValue(fight summary.Summary, player string, cat *consumables.RoleCatalogue) (float64, bool) {
	if cat == nil {
		return 0, false
	}
	return preparationTotal(fight, player, cat)
}

// carriedCategories are the categories a player carries exactly one of, so
// any one listed entry present earns the category's full weight.
var carriedCategories = []string{consumables.CategoryFlask, consumables.CategoryFood, consumables.CategoryWeaponEnchant}

// weighedCategories are every category Preparation weighs.
var weighedCategories = append(append([]string{}, carriedCategories...),
	consumables.CategoryWorldBuffs, consumables.CategoryCombatPotion)

// categoryCredit is the weight the pull earned in one category.
func categoryCredit(pull consumables.Pull, cat *consumables.RoleCatalogue, category string) float64 {
	switch category {
	case consumables.CategoryWorldBuffs:
		return worldBuffWeight(pull, cat)
	case consumables.CategoryCombatPotion:
		return potionWeight(pull, cat)
	}
	if pull.Has(category) {
		return float64(cat.Weights[category])
	}
	return 0
}

// worldBuffWeight credits world_buffs: several stack at once, so credit is
// proportional to how many of the catalogue's listed buffs are present.
func worldBuffWeight(pull consumables.Pull, cat *consumables.RoleCatalogue) float64 {
	if len(cat.WorldBuffs) == 0 {
		return 0
	}
	present := len(pull.Present[consumables.CategoryWorldBuffs])
	return float64(cat.Weights[consumables.CategoryWorldBuffs]) * float64(present) / float64(len(cat.WorldBuffs))
}

// potionWeight credits combat_potion proportional to uses against the
// shared-cooldown cap: the casts in the fight, plus one for a pre-pull
// potion, which shows only as a seeded aura.
func potionWeight(pull consumables.Pull, cat *consumables.RoleCatalogue) float64 {
	limit := cat.CombatPotion.MaxUses
	if limit <= 0 {
		return 0
	}
	uses := min(pull.PotionUses+int64(len(pull.Present[consumables.CategoryCombatPotion])), limit)
	return float64(cat.Weights[consumables.CategoryCombatPotion]) * float64(uses) / float64(limit)
}
