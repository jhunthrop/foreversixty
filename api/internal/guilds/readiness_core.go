// api/internal/guilds/readiness_core.go
//
// The per-character readiness checks the Readiness board (contract GET
// .../guilds/{id}/readiness, docs/contracts/2026-10-04-guild-centre-api.md) and the roster
// standing line's own "needs before next raid" (that same contract's home endpoint,
// standing.needs_before_next_raid) both read from - one computation, read by two surfaces,
// so they can never disagree about what a character still needs before the next raid.
package guilds

import (
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

// enchantableSlots are the four slots design spec §4.E's own readiness board checks
// ("Weapon/Chest/Cloak/Boots") - back is Codec.lua's and Export.lua's own slot name for
// Cloak.
var enchantableSlots = []string{"main_hand", "chest", "back", "feet"}

var enchantableSlotLabels = map[string]string{
	"main_hand": "Weapon", "chest": "Chest", "back": "Cloak", "feet": "Boots",
}

// talentPointsAtLevel60 is the most a level-60 character can have spent
// (api/internal/spec.MaxPoints, duplicated as a small constant rather than an import: every
// raider this page concerns is assumed level 60, design spec §4.E's own assumption, and this
// package has no other reason to depend on api/internal/spec).
const talentPointsAtLevel60 = 51

// CharacterReadiness is one character's full readiness computation, built from
// already-loaded inputs (computeReadiness, below) - no DB or file access of its own, so it
// is trivially unit-tested.
type CharacterReadiness struct {
	GearChecked         bool
	GearGap             bis.GearGap
	EnchantsChecked     bool
	MissingEnchantSlots []string
	// ConsumablesState is "stocked" | "short" | "unknown".
	ConsumablesState    string
	TalentPointsUnspent int
	ItemLevel           *int
	ItemLevelDelta      *int
}

// hasGearConsent reports whether consent unlocks the gear/gear_bags-gated checks
// (readiness's own floor for Gear gap and Enchants, design spec §4.E).
func hasGearConsent(consent string) bool { return consent == "gear" || consent == "gear_bags" }

// computeReadiness is pure. band/hasBand is the character's spec's BiS level-60 band
// (bis.LoadBand); gear/enchants come from fs1.Decode of the character's export; bags/
// hasBagsSection likewise (hasBagsSection distinguishes "the export carries an empty bags=
// section" is not actually distinguishable from "no bags= section at all" in fs1's own
// decode - both read as an empty slice - so this parameter is really just "did fs1.Decode
// succeed at all," passed through for clarity at the call site); talentPointsSpent is the
// three-tree digit sum fs1.Talents.Points already totals; itemLevel/medianItemLevel are the
// character's own and the roster's figures.
//
// The enchant check is deliberately not gated on anything the BiS band names: no BiS file
// in this repo carries a per-slot recommended enchant (checked against every band in
// data/builds/1.60.1.70009/bis/ - none of the seventeen inventory slots an alternatives-
// bearing file lists ever carries an "enchant" key), matching design spec §9's own "no
// curated recommended-enchant-per-spec table exists yet" - so this implements the design
// spec's narrower EXISTS claim instead: whether an enchantable, equipped slot carries any
// enchant at all, independent of which one. Named as a deviation from the contract's own
// "while the band's BiS lists an enchant for it" clause in CONTROL_CENTRE.md.
func computeReadiness(
	consent string,
	band bis.Band, hasBand bool,
	gear map[string]int, enchants map[string]int,
	bags []int, hasBagsSection bool,
	talentPointsSpent int, hasTalents bool,
	itemLevel, medianItemLevel *int,
) CharacterReadiness {
	var cr CharacterReadiness

	if hasGearConsent(consent) && hasBand {
		cr.GearChecked = true
		cr.GearGap = bis.GapFor(band, gear)
	}

	if hasGearConsent(consent) {
		cr.EnchantsChecked = true
		for _, slot := range enchantableSlots {
			if _, hasItem := gear[slot]; !hasItem {
				continue // nothing equipped there at all - not an enchant gap
			}
			if _, enchanted := enchants[slot]; !enchanted {
				cr.MissingEnchantSlots = append(cr.MissingEnchantSlots, slot)
			}
		}
	}

	switch {
	case consent != "gear_bags" || !hasBagsSection:
		cr.ConsumablesState = "unknown"
	case len(bags) > 0:
		cr.ConsumablesState = "stocked"
	default:
		cr.ConsumablesState = "short"
	}

	if hasTalents {
		if unspent := talentPointsAtLevel60 - talentPointsSpent; unspent > 0 {
			cr.TalentPointsUnspent = unspent
		}
	}

	cr.ItemLevel = itemLevel
	if itemLevel != nil && medianItemLevel != nil {
		delta := *itemLevel - *medianItemLevel
		cr.ItemLevelDelta = &delta
	}
	return cr
}

// plural is "" for 1, "s" otherwise - the only pluralisation this package's copy needs.
func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

// failText renders every failing check as a short phrase, worst-first (gear, then
// enchants, then consumables, then talent points - design spec §4.A.2's own example
// sentence order) - the shared wording between the standing line's "needs before next
// raid" and the readiness board's own officer nudge_text, so the two surfaces never phrase
// the same gap two different ways.
func (c CharacterReadiness) failText() []string {
	out := []string{}
	if c.GearChecked && c.GearGap.Upgrades > 0 {
		out = append(out, fmt.Sprintf("%d gear upgrade%s waiting (+%.1f DPS)",
			c.GearGap.Upgrades, plural(c.GearGap.Upgrades), c.GearGap.GainDps))
	}
	if c.EnchantsChecked && len(c.MissingEnchantSlots) > 0 {
		labels := make([]string, len(c.MissingEnchantSlots))
		for i, slot := range c.MissingEnchantSlots {
			labels[i] = enchantableSlotLabels[slot]
		}
		out = append(out, "no enchant: "+strings.Join(labels, ", "))
	}
	if c.ConsumablesState == "short" {
		out = append(out, "consumables short")
	}
	if c.TalentPointsUnspent > 0 {
		out = append(out, fmt.Sprintf("%d unspent talent point%s", c.TalentPointsUnspent, plural(c.TalentPointsUnspent)))
	}
	return out
}

// FailingCount is how many checks this character fails - the readiness row's own "failing"
// field and readinessScore's own failCount term (design spec §4.E).
func (c CharacterReadiness) FailingCount() int { return len(c.failText()) }

// Needs is failText capped at max - the standing line's own "max 3" (contract). max <= 0
// means uncapped.
func (c CharacterReadiness) Needs(max int) []string {
	f := c.failText()
	if max > 0 && len(f) > max {
		f = f[:max]
	}
	return f
}
