// api/internal/bis/gain.go
//
// The gear-gap gain rule, ported from web/src/lib/home/upgrades.ts exactly as the contract
// states it (docs/contracts/2026-10-04-guild-centre-api.md's "Gear gap" paragraph): pick ==
// worn -> none; worn in the band's alternatives -> gain = -dps_delta; otherwise not
// sim-checked. This is deliberately the contract's own three-branch shorthand, not
// upgrades.ts's full rule - that TypeScript module also falls back to a stat-weight score
// diff (scoreItem) for a non-weapon slot with no listed alternative, which needs the item
// stat table and scorer this package does not have. The contract names only the three
// branches below, so that is exactly what this port does; see
// api/internal/guilds/CONTROL_CENTRE.md for this named as a deviation from upgrades.ts's
// fuller rule, not from the contract.
package bis

// SlotVerdict is one slot's gear-gap verdict against a band's pick.
type SlotVerdict struct {
	// Upgrade is true whenever the worn item is not the pick and not a zero-delta tie
	// with it - the slot counts toward a readiness row's upgrades/not_sim_checked and a
	// loot candidate's own ranking.
	Upgrade bool
	// GainDps is the pick's measured advantage over the worn item, only when the worn
	// item is a listed alternative with a nonzero delta. Nil otherwise (including when
	// Upgrade is true but NotSimChecked is also true).
	GainDps *float64
	// NotSimChecked is true when Upgrade is true but no listed alternative measured the
	// worn item at all - the worn item is unknown to the band's own ranking.
	NotSimChecked bool
}

// VerdictFor is pure: no DB, no file I/O, so it is trivially table-tested against hand-built
// bands. wornItemID/hasWorn describe what (if anything) the character has equipped in
// pick's own slot.
func VerdictFor(pick Slot, wornItemID int, hasWorn bool) SlotVerdict {
	if hasWorn && wornItemID == pick.ItemID {
		return SlotVerdict{}
	}
	if hasWorn {
		for _, alt := range pick.Alternatives {
			if alt.ItemID != wornItemID {
				continue
			}
			if alt.DpsDelta == 0 {
				// A real, measured tie: no upgrade, same as wearing the pick itself.
				return SlotVerdict{}
			}
			gain := -alt.DpsDelta
			return SlotVerdict{Upgrade: true, GainDps: &gain}
		}
	}
	return SlotVerdict{Upgrade: true, NotSimChecked: true}
}

// GearGap is the readiness board's own gear_gap object: how many of a band's slots are an
// upgrade, the DPS sum of the ones with a known gain, and how many have none.
type GearGap struct {
	Upgrades      int
	GainDps       float64
	NotSimChecked int
}

// GapFor sums VerdictFor over every slot band names a real pick for (ItemID > 0 - the BiS
// table's own "nothing to recommend here" sentinel is skipped, never treated as an upgrade
// target). gear is slot -> worn item id, fs1.Decoded.Gear's own shape.
func GapFor(band Band, gear map[string]int) GearGap {
	var out GearGap
	for _, slot := range band.Slots {
		if slot.ItemID <= 0 {
			continue
		}
		wornID, hasWorn := gear[slot.Slot]
		v := VerdictFor(slot, wornID, hasWorn)
		if !v.Upgrade {
			continue
		}
		out.Upgrades++
		switch {
		case v.NotSimChecked:
			out.NotSimChecked++
		case v.GainDps != nil:
			out.GainDps += *v.GainDps
		}
	}
	return out
}
