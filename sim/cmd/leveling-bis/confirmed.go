package main

// Preferring confirmed stats inside the noise.
//
// Some items carry stats the client has not confirmed: the id is in the
// client but its ItemSparse stats row is not, so the ranker scores the 1.12
// classic-db itemization (candidate.ClientUnconfirmed). Such an item is never
// excluded, but it may only take a slot from a confirmed-stats item when it
// beats that item by more than the two runs' combined standard error. A gap
// inside the error is a tie the noise decides, and a tie goes to the item
// whose stats we can show.

import (
	"fmt"
	"log"
	"math"
)

// confirmedAlternative is the confirmed-stats item to try against slot's
// unconfirmed pick: the sim-verified runner-up when it is confirmed, else the
// best-scored confirmed candidate the slot's pool offers. Nil when the slot
// has none, or when wearing it would break a pair or two-hand rule.
func confirmedAlternative(picks map[string]slotPick, slot string, pool []scored) *scored {
	if ru := picks[slot].RunnerUp; ru != nil && !ru.ClientUnconfirmed {
		return usableAlternative(picks, slot, ru)
	}
	for i := range pool {
		if pool[i].ClientUnconfirmed || pool[i].ID == picks[slot].Item.ID {
			continue
		}
		if alt := usableAlternative(picks, slot, &pool[i]); alt != nil {
			return alt
		}
	}
	return nil
}

func usableAlternative(picks map[string]slotPick, slot string, alt *scored) *scored {
	if wouldDuplicatePairMate(picks, slot, alt.ID, alt.Name) {
		return nil
	}
	if slot == "off_hand" && picks["main_hand"].Item != nil && picks["main_hand"].Item.TwoHand {
		return nil
	}
	return alt
}

// preferConfirmedStats walks the final picks and, for each slot held by an
// unconfirmed-stats item, measures the set with the confirmed alternative
// worn instead. The alternative is published unless the unconfirmed item
// beats it by more than the combined standard error; the unconfirmed item
// then stays as the slot's runner-up (and in alternatives). It returns the
// number of slots changed. A slot with no confirmed candidate keeps its pick.
func preferConfirmedStats(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, vb verifiedBand, bySlot map[string][]scored) (verifiedBand, int, error) {
	if !hasUnconfirmedPick(vb.picks) {
		return vb, 0, nil
	}
	baseDPS, baseErr, err := measureSet(runner, spec, race, classSlug, level, talents, vb.picks)
	if err != nil {
		return verifiedBand{}, 0, fmt.Errorf("measuring the set before the confirmed-stats check: %w", err)
	}
	picks := copyPicks(vb.picks)
	held := map[string]bool{}
	var errs []string
	for _, slot := range slotOrder {
		cur := picks[slot]
		if cur.Item == nil || !cur.Item.ClientUnconfirmed {
			continue
		}
		alt := confirmedAlternative(picks, slot, bySlot[slot])
		if alt == nil {
			continue
		}
		req := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, swapSlot(picks, slot, alt.ID, alt.TwoHand)), verifyIterations, verifySeed)
		altDPS, altErr, runErr := runner.RunPlainDPSWithError(req)
		if runErr != nil {
			errs = append(errs, fmt.Sprintf("%s: confirmed-stats alternative %s (id %d): %v", slot, alt.Name, alt.ID, runErr))
			continue
		}
		// The doubt here is the item's stats, which the client has not
		// confirmed, not the sims' noise: a lead inside the adoption margin
		// (swapMargin, 1% of the set) is not enough to send a player after
		// it, however many iterations the verify run had. The sims' own
		// error still applies when it is the larger of the two.
		if baseDPS-altDPS > math.Max(math.Hypot(baseErr, altErr), baseDPS*swapMargin) {
			continue
		}
		adopted := *alt
		adopted.MeasuredDPS = altDPS
		picks[slot] = slotPick{Item: &adopted, RunnerUp: cur.Item}
		if slot == "main_hand" && alt.TwoHand {
			picks["off_hand"] = slotPick{}
		}
		baseDPS, baseErr = altDPS, altErr
		held[slot] = true
		log.Printf("leveling-bis: %s: %s keeps %s (stats the client has not confirmed) out: it does not beat the confirmed %s beyond error", spec.Spec, slot, cur.Item.Name, alt.Name)
	}
	if len(held) == 0 {
		vb.errors = append(vb.errors, errs...)
		return vb, 0, nil
	}
	return verifiedBand{
		picks:  dropStaleSetBonuses(picks),
		setDPS: baseDPS,
		swaps:  withoutSlots(vb.swaps, held),
		errors: append(vb.errors, errs...),
	}, len(held), nil
}

func hasUnconfirmedPick(picks map[string]slotPick) bool {
	for _, pk := range picks {
		if pk.Item != nil && pk.Item.ClientUnconfirmed {
			return true
		}
	}
	return false
}

func copyPicks(picks map[string]slotPick) map[string]slotPick {
	out := make(map[string]slotPick, len(picks))
	for slot, pk := range picks {
		out[slot] = pk
	}
	return out
}

// withoutSlots drops the swap results of slots whose pick changed, so no
// swap_note describes a comparison the published row no longer shows.
func withoutSlots(swaps []swapResult, slots map[string]bool) []swapResult {
	var out []swapResult
	for _, sw := range swaps {
		if !slots[sw.Slot] {
			out = append(out, sw)
		}
	}
	return out
}
