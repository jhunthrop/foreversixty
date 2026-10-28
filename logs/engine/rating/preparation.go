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
// character sheet for player. A carried category whose list is incomplete
// and reads nothing is unknown, not missing: its weight leaves the total
// and the rest is read out of what is left (unknownWeight).
func preparationTotal(fight summary.Summary, player string, cat *consumables.RoleCatalogue) (float64, bool) {
	pull := consumables.AtPull(fight, player, *cat)
	if !pull.Known {
		return 0, false
	}
	possible := totalWeight(cat) - unknownWeight(pull, cat)
	if possible <= 0 {
		return 0, true
	}
	earned := carriedWeight(pull, cat) + worldBuffWeight(pull, cat) + potionWeight(pull, cat)
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

func carriedWeight(pull consumables.Pull, cat *consumables.RoleCatalogue) float64 {
	var total float64
	for _, category := range carriedCategories {
		if pull.Has(category) {
			total += float64(cat.Weights[category])
		}
	}
	return total
}

// weighedCategories are every category Preparation weighs.
var weighedCategories = append(append([]string{}, carriedCategories...),
	consumables.CategoryWorldBuffs, consumables.CategoryCombatPotion)

// unknownWeight is the weight of every category the catalogue cannot judge
// absent: its list is incomplete and nothing on it was read. The food list
// names one food per role, the dps potion list is partial, and the
// weapon_enchant entries are the stone and oil application spells, which
// never show as auras (a temporary enchant rides on the weapon item, and
// AtPull does not read it), so each stays unknown unless a listed entry is
// seen.
func unknownWeight(pull consumables.Pull, cat *consumables.RoleCatalogue) float64 {
	var total float64
	for _, category := range weighedCategories {
		if !cat.IsComplete(category) && !readAny(pull, category) {
			total += float64(cat.Weights[category])
		}
	}
	return total
}

// readAny is whether the pull read anything of the category: a listed
// entry up, or for combat_potion a use.
func readAny(pull consumables.Pull, category string) bool {
	if category == consumables.CategoryCombatPotion && pull.PotionUses > 0 {
		return true
	}
	return pull.Has(category)
}

func totalWeight(cat *consumables.RoleCatalogue) float64 {
	var total float64
	for _, w := range cat.Weights {
		total += float64(w)
	}
	return total
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
