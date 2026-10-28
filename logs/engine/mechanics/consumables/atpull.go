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

// MissedAtPull is whether the pull proves the category absent: the log has
// the player's character sheet, rc's list for the category is complete, and
// nothing on it was read (no entry up; for combat_potion, no use either).
// An incomplete list can prove presence, never absence, and no sheet proves
// nothing: unknown is never missing. Every "missing" reader asks this.
func (rc RoleCatalogue) MissedAtPull(p Pull, category string) bool {
	if !p.Known || !rc.IsComplete(category) || p.Has(category) {
		return false
	}
	return category != CategoryCombatPotion || p.PotionUses == 0
}

// AtPullGraceMS is how long into the pull an aura APPLIED during the fight
// still counts as brought into it: a flask drunk or a Battle Shout cast on
// the pull (ruling 2026-10-27, log coach spec §13). Combat potions take no
// grace: a potion drunk in the fight is a use (F6's rule), not preparation.
const AtPullGraceMS = 10_000

// Seeded reports whether guid's own COMBATANT_INFO listed spellID: the
// summary's explicit flag, so a snapshot written a few ms after
// ENCOUNTER_START still reads as the pull.
func Seeded(fight summary.Summary, guid string, spellID int64) bool {
	for i := range fight.Auras {
		tr := &fight.Auras[i]
		if tr.TargetGUID == guid && tr.SpellID == spellID && tr.Seeded {
			return true
		}
	}
	return false
}

// upAtPull is Seeded, or an aura on guid with a segment starting within
// AtPullGraceMS of the pull.
func upAtPull(fight summary.Summary, guid string, spellID int64) bool {
	for i := range fight.Auras {
		tr := &fight.Auras[i]
		if tr.TargetGUID != guid || tr.SpellID != spellID {
			continue
		}
		if tr.Seeded {
			return true
		}
		for _, seg := range tr.Segments {
			if seg.StartMS <= AtPullGraceMS {
				return true
			}
		}
	}
	return false
}

// AtPull reads what the player brought into the pull against rc: an entry
// is present when seeded, or (outside combat_potion) applied within
// AtPullGraceMS. With no character sheet the log cannot say what was up, so
// Present stays empty and Known false: a reader must call that unknown,
// never missing. PotionUses counts casts, which need no sheet.
func AtPull(fight summary.Summary, guid string, rc RoleCatalogue) Pull {
	p := Pull{PotionUses: potionUses(fight, guid, rc.CombatPotion.Entries)}
	if !hasSheet(fight, guid) {
		return p
	}
	p.Known = true
	p.Present = map[string][]Entry{}
	for _, cat := range categories {
		up := upAtPull
		if cat == CategoryCombatPotion {
			up = Seeded
		}
		for _, e := range rc.Entries(cat) {
			if up(fight, guid, e.SpellID) {
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
	Known            bool // the player has a character sheet; false means Present says nothing
	Present          bool // any rank up at the pull (seeded, or applied within AtPullGraceMS); false when !Known
	ProviderInRoster bool // a roster row of Buff.Class in this fight other than the player's own
}

// RaidBuffsAtPull reads each buff on guid at the pull, and whether someone
// else in the fight's roster was of a class that could have provided it (a
// buff only the player provides is their own uptime, not a raid buff). With no
// character sheet every state is Known false and Present false: unknown,
// never missing.
func RaidBuffsAtPull(fight summary.Summary, guid string, buffs []RaidBuff) []RaidBuffState {
	classes := map[string]bool{}
	for _, r := range fight.Roster {
		if r.Class != "" && r.GUID != guid {
			classes[r.Class] = true
		}
	}
	known := hasSheet(fight, guid)
	out := make([]RaidBuffState, 0, len(buffs))
	for _, b := range buffs {
		out = append(out, RaidBuffState{Buff: b, Known: known,
			Present:          known && anyUpAtPull(fight, guid, b.SpellIDs),
			ProviderInRoster: classes[b.Class]})
	}
	return out
}

func anyUpAtPull(fight summary.Summary, guid string, ids []int64) bool {
	for _, id := range ids {
		if upAtPull(fight, guid, id) {
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
