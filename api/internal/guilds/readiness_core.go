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

// raidLevel is the level every readiness check is measured for: the gear, enchant,
// consumable and talent needs of a level-60 raid. A character below it is leveling and is not
// graded at all (week-one spec §3.5).
const raidLevel = 60

// firstTalentLevel is the level the first talent point arrives at; one more comes every level
// after it, up to talentPointsAtLevel60.
const firstTalentLevel = 10

// talentPointsAtLevel60 is the most a level-60 character can have spent - the "export
// carries no level section at all" fallback maxTalentPoints uses below.
const talentPointsAtLevel60 = raidLevel - firstTalentLevel + 1

// gearUnitDPS is a gear gap's unit for a damage spec, the only unit any BiS file measures a
// gain in today (healer bands carry stat points, tank bands Forever Score).
const gearUnitDPS = "dps"

// maxTalentPoints is the most points a character at level could have spent by now: one
// at firstTalentLevel and one more every level after it, up to talentPointsAtLevel60, none
// below firstTalentLevel. With no level on file the character is assumed level 60
// (talentPointsAtLevel60) and assumedLevel60 tells the caller so, since the officer's
// nudge_text must say so rather than present a guess as fact.
func maxTalentPoints(level *int) (points int, assumedLevel60 bool) {
	if level == nil {
		return talentPointsAtLevel60, true
	}
	points = *level - firstTalentLevel + 1
	switch {
	case points < 0:
		return 0, false
	case points > talentPointsAtLevel60:
		return talentPointsAtLevel60, false
	}
	return points, false
}

// leveling reports whether a known level is below raidLevel: such a character is not graded.
func leveling(level *int) bool { return level != nil && *level < raidLevel }

// readinessInput is everything one character's readiness reads, already loaded: the
// character's consent, its spec's level-60 BiS band (bis.LoadBand; HasBand false without one),
// its export's gear, enchants and bags (fs1.Decode; HasBagsSection tells an empty bags section
// from none), its talent points spent (HasTalents false when the export carried none), its
// level (nil when neither the export nor the addon roster carries one) and its own and the
// roster's item levels.
type readinessInput struct {
	Consent                    string
	Band                       bis.Band
	HasBand                    bool
	Gear, Enchants             map[string]int
	Bags                       []int
	HasBagsSection             bool
	TalentPointsSpent          int
	HasTalents                 bool
	Level                      *int
	ItemLevel, MedianItemLevel *int
}

// readinessInputOf is a roster row's readiness input against band and the roster median.
func readinessInputOf(row RosterRow, band bis.Band, hasBand bool, median *int) readinessInput {
	return readinessInput{
		Consent: row.Consent, Band: band, HasBand: hasBand,
		Gear: row.gear, Enchants: row.enchants, Bags: row.bags, HasBagsSection: row.hasBagsSection,
		TalentPointsSpent: row.talentPointsSpent, HasTalents: row.hasTalents,
		Level: row.Level, ItemLevel: row.ItemLevel, MedianItemLevel: median,
	}
}

// readinessOf is one roster row's readiness, its band read from the row's own spec.
func (s *Store) readinessOf(row RosterRow, median *int) CharacterReadiness {
	band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
	return computeReadiness(readinessInputOf(row, band, hasBand, median))
}

// CharacterReadiness is one character's full readiness computation, built from
// already-loaded inputs (computeReadiness, below) - no DB or file access of its own, so it
// is trivially unit-tested.
type CharacterReadiness struct {
	// Leveling is true for a character whose known level is below raidLevel: nothing below
	// is graded and FailingCount is 0 (week-one spec §3.5).
	Leveling    bool
	Level       *int
	GearChecked bool
	GearGap     bis.GearGap
	// GearUnit is what GearGap.GainDps counts: gearUnitDPS for a damage band, "" for a healer
	// or tank band, whose gap is never printed as DPS.
	GearUnit            string
	EnchantsChecked     bool
	MissingEnchantSlots []string
	// ConsumablesState is "stocked" | "short" | "unknown".
	ConsumablesState    string
	TalentPointsUnspent int
	// TalentLevelAssumed is true exactly when TalentPointsUnspent was computed against
	// talentPointsAtLevel60 because no level was on file - the officer-only nudge_text's own
	// cue to say so rather than present a guess as fact.
	TalentLevelAssumed bool
	// TalentsRead is true when the export carried a talents section at all, so a zero
	// TalentPointsUnspent is a real zero and not the absence of a reading.
	TalentsRead    bool
	ItemLevel      *int
	ItemLevelDelta *int
}

// hasGearConsent reports whether consent unlocks the gear/gear_bags-gated checks
// (readiness's own floor for Gear gap and Enchants, design spec §4.E).
func hasGearConsent(consent string) bool { return consent == "gear" || consent == "gear_bags" }

// computeReadiness is pure. A leveling character (a known level under raidLevel) gets its
// level and item level and nothing graded: level-60 raid needs are not a level 52's to chase.
//
// The enchant check is deliberately not gated on anything the BiS band names: no BiS file
// in this repo carries a per-slot recommended enchant (checked against every band in
// data/builds/<build>/bis/ (Store.BisBuild; last checked on 1.60.1.70630) - none of the
// seventeen inventory slots an alternatives-bearing file lists ever carries an "enchant"
// key), matching design spec §9's own "no curated recommended-enchant-per-spec table exists
// yet" - so this implements the design spec's narrower EXISTS claim instead: whether an
// enchantable, equipped slot carries any enchant at all, independent of which one. Named as a
// deviation from the contract's own "while the band's BiS lists an enchant for it" clause in
// CONTROL_CENTRE.md.
func computeReadiness(in readinessInput) CharacterReadiness {
	cr := CharacterReadiness{Level: in.Level, ItemLevel: in.ItemLevel, ConsumablesState: consumablesUnknown}
	if leveling(in.Level) {
		cr.Leveling = true
		return cr
	}

	if hasGearConsent(in.Consent) && in.HasBand {
		cr.GearChecked = true
		cr.GearGap = bis.GapFor(in.Band, in.Gear)
		if in.Band.Role == "" {
			cr.GearUnit = gearUnitDPS
		}
	}
	if hasGearConsent(in.Consent) {
		cr.EnchantsChecked = true
		cr.MissingEnchantSlots = missingEnchants(in.Gear, in.Enchants)
	}
	cr.ConsumablesState = consumablesState(in.Consent, in.Bags, in.HasBagsSection)

	if in.HasTalents {
		cr.TalentsRead = true
		maxPoints, assumedLevel60 := maxTalentPoints(in.Level)
		if unspent := maxPoints - in.TalentPointsSpent; unspent > 0 {
			cr.TalentPointsUnspent = unspent
		}
		cr.TalentLevelAssumed = assumedLevel60
	}

	if in.ItemLevel != nil && in.MedianItemLevel != nil {
		delta := *in.ItemLevel - *in.MedianItemLevel
		cr.ItemLevelDelta = &delta
	}
	return cr
}

// missingEnchants is every enchantable slot with an item equipped and no enchant on it.
func missingEnchants(gear, enchants map[string]int) []string {
	var out []string
	for _, slot := range enchantableSlots {
		if _, hasItem := gear[slot]; !hasItem {
			continue // nothing equipped there at all - not an enchant gap
		}
		if _, enchanted := enchants[slot]; !enchanted {
			out = append(out, slot)
		}
	}
	return out
}

// The consumables states a readiness row reads.
const (
	consumablesUnknown = "unknown"
	consumablesStocked = "stocked"
	consumablesShort   = "short"
)

// consumablesState is the bags rule: "unknown" without the bags share, "stocked" with
// anything in the bags, else "short". The readiness board and the raid report's Before pill
// both read it.
func consumablesState(consent string, bags []int, hasBagsSection bool) string {
	switch {
	case consent != "gear_bags" || !hasBagsSection:
		return consumablesUnknown
	case len(bags) > 0:
		return consumablesStocked
	default:
		return consumablesShort
	}
}

// plural is "" for 1, "s" otherwise - the only pluralisation this package's copy needs.
func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

// gearFailText is the gear gap as a phrase: the upgrade count, then the DPS it is worth only
// for a damage spec with a sim-checked gain (an unsimmed upgrade has no number to print, and a
// healer's or tank's gap is never a DPS figure).
func (c CharacterReadiness) gearFailText() string {
	text := fmt.Sprintf("%d gear upgrade%s waiting", c.GearGap.Upgrades, plural(c.GearGap.Upgrades))
	if c.GearUnit == gearUnitDPS && c.GearGap.GainDps >= minPrintedGain {
		text += fmt.Sprintf(" (+%.1f DPS)", c.GearGap.GainDps)
	}
	return text
}

// minPrintedGain is the smallest gain that prints as more than +0.0 at one decimal.
const minPrintedGain = 0.05

// failText renders every failing check as a short phrase, worst-first (gear, then
// enchants, then consumables, then talent points - design spec §4.A.2's own example
// sentence order) - the shared wording between the standing line's "needs before next
// raid" and the readiness board's own officer nudge_text, so the two surfaces never phrase
// the same gap two different ways.
func (c CharacterReadiness) failText() []string {
	out := []string{}
	if c.GearChecked && c.GearGap.Upgrades > 0 {
		out = append(out, c.gearFailText())
	}
	if c.EnchantsChecked && len(c.MissingEnchantSlots) > 0 {
		labels := make([]string, len(c.MissingEnchantSlots))
		for i, slot := range c.MissingEnchantSlots {
			labels[i] = enchantableSlotLabels[slot]
		}
		out = append(out, "no enchant: "+strings.Join(labels, ", "))
	}
	if c.ConsumablesState == consumablesShort {
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
