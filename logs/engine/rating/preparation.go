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
func scorePreparation(fight summary.Summary, player string, bracket Bracket, cat *consumables.RoleCatalogue, src PercentileSource) Component {
	bracket.Component = ComponentNamePreparation
	c := Component{Name: ComponentNamePreparation}
	if cat == nil {
		c.Excluded, c.Reason = true, ReasonNoConsumableCatalogue
		return c
	}

	combatant, _ := combatantFor(fight, player)
	var total float64

	// flask/food/weapon_enchant: a player carries exactly one of each, so
	// any one listed entry present earns the category's full weight.
	if hasAnyConsumable(combatant, cat.Flask) {
		total += float64(cat.Weights["flask"])
	}
	if hasAnyConsumable(combatant, cat.Food) {
		total += float64(cat.Weights["food"])
	}
	if len(cat.WeaponEnchant) > 0 && hasAnyConsumable(combatant, cat.WeaponEnchant) {
		total += float64(cat.Weights["weapon_enchant"])
	}
	// world_buffs: several stack at once, so credit is proportional to how
	// many of the catalogue's listed buffs are present.
	if len(cat.WorldBuffs) > 0 {
		present := countPresentConsumables(combatant, cat.WorldBuffs)
		total += float64(cat.Weights["world_buffs"]) * float64(present) / float64(len(cat.WorldBuffs))
	}
	// combat_potion: credit proportional to uses against the shared-
	// cooldown cap.
	if cat.CombatPotion.MaxUses > 0 {
		uses := potionUses(fight, player, cat.CombatPotion.Entries)
		if uses > cat.CombatPotion.MaxUses {
			uses = cat.CombatPotion.MaxUses
		}
		total += float64(cat.Weights["combat_potion"]) * float64(uses) / float64(cat.CombatPotion.MaxUses)
	}

	return placeBoundedOrSelf(c, src, bracket, clamp(total, 0, 100))
}

func combatantFor(fight summary.Summary, player string) (summary.CombatantRow, bool) {
	for _, cr := range fight.Combatants {
		if cr.GUID == player {
			return cr, true
		}
	}
	return summary.CombatantRow{}, false
}

func hasAnyConsumable(combatant summary.CombatantRow, entries []consumables.Entry) bool {
	return countPresentConsumables(combatant, entries) > 0
}

func countPresentConsumables(combatant summary.CombatantRow, entries []consumables.Entry) int {
	count := 0
	for _, e := range entries {
		for _, c := range combatant.Consumables {
			if c.SpellID == e.SpellID {
				count++
				break
			}
		}
	}
	return count
}

func potionUses(fight summary.Summary, player string, entries []consumables.Entry) int64 {
	var uses int64
	for _, cr := range fight.Casts {
		if cr.OwnerGUID != player {
			continue
		}
		for _, e := range entries {
			if cr.SpellID == e.SpellID {
				uses += cr.Succeeded
			}
		}
	}
	return uses
}
