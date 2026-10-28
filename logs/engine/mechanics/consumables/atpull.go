// logs/engine/mechanics/consumables/atpull.go
package consumables

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Pull is what one player brought into the pull, read from the seeded aura tracks.
type Pull struct {
	Known      bool               // the log has a character sheet (a combatant row with gear) for the player
	Present    map[string][]Entry // category -> catalogue entries seeded on the player at 0 ms
	PotionUses int64              // successful casts of combat_potion entries, owner-folded
}

// Has reports whether any entry of the category was up on the player at the pull.
func (p Pull) Has(category string) bool { return len(p.Present[category]) > 0 }

// Seeded reports whether guid carried spellID at the pull: an aura track on
// guid with a segment starting at 0 ms. summary seeds those tracks from
// COMBATANT_INFO, which the log writes at the pull with no APPLIED line.
func Seeded(fight summary.Summary, guid string, spellID int64) bool {
	for i := range fight.Auras {
		tr := &fight.Auras[i]
		if tr.TargetGUID != guid || tr.SpellID != spellID {
			continue
		}
		for _, seg := range tr.Segments {
			if seg.StartMS == 0 {
				return true
			}
		}
	}
	return false
}

// AtPull reads what the player brought into the pull against rc. With no
// character sheet the log cannot say what was up, so Present stays empty
// and Known false: a reader must call that unknown, never missing.
// PotionUses counts casts, which need no sheet.
func AtPull(fight summary.Summary, guid string, rc RoleCatalogue) Pull {
	p := Pull{PotionUses: potionUses(fight, guid, rc.CombatPotion.Entries)}
	if !hasSheet(fight, guid) {
		return p
	}
	p.Known = true
	p.Present = map[string][]Entry{}
	for _, cat := range categories {
		for _, e := range rc.Entries(cat) {
			if Seeded(fight, guid, e.SpellID) {
				p.Present[cat] = append(p.Present[cat], e)
			}
		}
	}
	return p
}

// hasSheet is the rating's existing rule (api/internal/rating/last_raid.go):
// a combatant row with no gear is no character sheet worth reading.
func hasSheet(fight summary.Summary, guid string) bool {
	for i := range fight.Combatants {
		if fight.Combatants[i].GUID == guid {
			return len(fight.Combatants[i].Gear) > 0
		}
	}
	return false
}

// potionUses sums successful casts of the listed potions by guid and by
// anything guid owns (CastRow.OwnerGUID is the caster's own GUID for a
// player).
func potionUses(fight summary.Summary, guid string, potions []Entry) int64 {
	ids := make(map[int64]bool, len(potions))
	for _, e := range potions {
		ids[e.SpellID] = true
	}
	var n int64
	for _, c := range fight.Casts {
		if c.OwnerGUID == guid && ids[c.SpellID] {
			n += c.Succeeded
		}
	}
	return n
}

// RaidBuffState is one raid buff on one player at the pull.
type RaidBuffState struct {
	Buff             RaidBuff
	Present          bool // any rank seeded on the player at 0 ms
	ProviderInRoster bool // a roster row of Buff.Class in this fight
}

// RaidBuffsAtPull reads each buff on guid at the pull, and whether the
// fight's roster held a class that could have provided it.
func RaidBuffsAtPull(fight summary.Summary, guid string, buffs []RaidBuff) []RaidBuffState {
	classes := map[string]bool{}
	for _, r := range fight.Roster {
		if r.Class != "" {
			classes[r.Class] = true
		}
	}
	out := make([]RaidBuffState, 0, len(buffs))
	for _, b := range buffs {
		out = append(out, RaidBuffState{Buff: b, Present: anySeeded(fight, guid, b.SpellIDs), ProviderInRoster: classes[b.Class]})
	}
	return out
}

func anySeeded(fight summary.Summary, guid string, ids []int64) bool {
	for _, id := range ids {
		if Seeded(fight, guid, id) {
			return true
		}
	}
	return false
}

// AllRoles is every role's entries merged, for the discipline readers that
// ask "did the player bring any flask" without a role. Each spell id
// appears once, a category is complete only if every role calls it
// complete, MaxUses is the largest any role allows, and Weights is nil:
// the merge has no Preparation weighting.
func (c Catalogue) AllRoles() RoleCatalogue {
	roles := make([]string, 0, len(c.Roles))
	for r := range c.Roles {
		roles = append(roles, r)
	}
	sort.Strings(roles)
	all := RoleCatalogue{Complete: map[string]bool{}}
	for _, cat := range categories {
		all.Complete[cat] = len(roles) > 0
	}
	for _, r := range roles {
		rc := c.Roles[r]
		for _, cat := range categories {
			all.Complete[cat] = all.Complete[cat] && rc.IsComplete(cat)
		}
		all.Flask = mergeEntries(all.Flask, rc.Flask)
		all.Food = mergeEntries(all.Food, rc.Food)
		all.WeaponEnchant = mergeEntries(all.WeaponEnchant, rc.WeaponEnchant)
		all.WorldBuffs = mergeEntries(all.WorldBuffs, rc.WorldBuffs)
		all.CombatPotion.Entries = mergeEntries(all.CombatPotion.Entries, rc.CombatPotion.Entries)
		all.CombatPotion.MaxUses = max(all.CombatPotion.MaxUses, rc.CombatPotion.MaxUses)
	}
	return all
}

// mergeEntries returns a new slice of into plus each of more whose spell
// id into does not already hold.
func mergeEntries(into, more []Entry) []Entry {
	out := append([]Entry(nil), into...)
	seen := make(map[int64]bool, len(out))
	for _, e := range out {
		seen[e.SpellID] = true
	}
	for _, e := range more {
		if !seen[e.SpellID] {
			seen[e.SpellID] = true
			out = append(out, e)
		}
	}
	return out
}
