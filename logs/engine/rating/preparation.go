// logs/engine/rating/preparation.go
package rating

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// scorePreparation implements spec §1.3's Preparation component: a
// completeness score built from the consumable catalogue, percentile-
// within-bracket (or its own value as the absolute standard, RULING R5,
// since it is already a bounded 0-100 score).
//
// What the player brought is read through consumables.AtPull: the seeded
// aura tracks (or an aura applied within AtPullGraceMS of the pull), never
// CombatantRow.Consumables, which production never fills (log coach spec
// §13). With no character sheet nothing reads present; potion casts still
// count, since they need no sheet.
func scorePreparation(fight summary.Summary, player string, bracket Bracket, cat *consumables.RoleCatalogue, src PercentileSource) Component {
	bracket.Component = ComponentNamePreparation
	c := Component{Name: ComponentNamePreparation}
	if cat == nil {
		c.Excluded, c.Reason = true, ReasonNoConsumableCatalogue
		return c
	}
	value, _ := PreparationValue(fight, player, cat)
	return placeBoundedOrSelf(c, src, bracket, value)
}

// PreparationValue is the raw 0-100 completeness scorePreparation places for player; false
// with no catalogue, where the component is excluded instead. Exported so the rating
// rebuild can refold a preparation digest from stored summaries without scoring a card.
func PreparationValue(fight summary.Summary, player string, cat *consumables.RoleCatalogue) (float64, bool) {
	if cat == nil {
		return 0, false
	}
	pull := consumables.AtPull(fight, player, *cat)
	total := carriedWeight(pull, cat) + worldBuffWeight(pull, cat) + potionWeight(pull, cat)
	return clamp(total, 0, 100), true
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
