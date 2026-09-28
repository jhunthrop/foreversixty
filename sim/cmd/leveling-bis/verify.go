package main

// Verification: running the chosen set, and each slot's runner-up
// swapped in alone, as plain DPS sims - and comparing.
//
// This lane's brief asks for "sim/bulk Top Gear (mode gear)... 300
// iterations". Top Gear's gear mode (sim/bulk's gearCombinations,
// see expand.go) builds the FULL CROSS PRODUCT of every slot's
// candidates at once - the right shape when the page is asking "try
// these few candidates in these few slots, in every combination",
// because set-bonus and stat-breakpoint interactions between slots
// are exactly what it is for. This prototype's verification question
// is narrower - one runner-up per slot, checked against the picked
// set one slot at a time - and running it through gearCombinations
// would still enumerate 2^(number of slots with a runner-up) full
// sims (every slot could independently keep the pick or take the
// runner-up), which for hunter-marksmanship's dozen-plus slots is
// thousands of 300-iteration sims, not the "keep it under ~5 minutes"
// this lane's brief asks for. So this file runs the one-swap-at-a-
// time shape directly: sim/bulk/expand.go's OWN singleCombinations
// (unexported, "one substitution at a time... every placement on its
// own") is the closest thing sim/bulk has to this, but it is reached
// only through Bulk.Mode != "gear" (drops/talents), a different
// tool's request shape. Mirrored here rather than routed through
// sim/bulk at all - documented as this lane's report says to.
//
// The two-hand/off-hand validity rule this file still has to honour
// (a runner-up main-hand two-hander cannot be tried beside a kept
// off-hand item) is handled by dropping off_hand from that one
// trial, which is what equipping the two-hander would actually do;
// pick() already guarantees the reverse case (an off_hand runner-up
// existing while main_hand is two-handed) never arises, because it
// leaves off_hand's own slotPick empty in that case.

import (
	"fmt"
	"sort"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// verifyIterations is this file's fixed 300, per the brief.
const verifyIterations = 300

// verifySeed is one seed for every verification run, so a report is
// reproducible.
const verifySeed = 7

// swapResult is one slot's verification: the runner-up's own
// full-set DPS (with just that slot swapped from the pick) and
// whether it beat the baseline.
type swapResult struct {
	Slot    string
	SwapDPS float64
	Beat    bool
}

// buildGear turns a pick map into the engine's gear list, dropping
// off_hand when main_hand is two-handed (bulk/expand.go's valid():
// "no two-hander beside an off-hand" - stated here as "do not equip
// one" rather than "refuse the combination", since this function
// only ever builds gear this command chose itself).
func buildGear(picks map[string]slotPick) []api.GearSlot {
	twoHand := picks["main_hand"].Item != nil && picks["main_hand"].Item.TwoHand
	var out []api.GearSlot
	for _, slot := range slotOrder {
		if slot == "off_hand" && twoHand {
			continue
		}
		if picks[slot].Item == nil {
			continue
		}
		out = append(out, api.GearSlot{Slot: slot, ItemID: picks[slot].Item.ID})
	}
	return out
}

// pairSlot is each ring/trinket slot's partner: pick() only ever
// excludes a pair-mate's own id/name from the OTHER slot's candidate
// list going forward (finger1 decided before finger2), so a slot's
// own runner-up can still be the exact item its pair-mate already
// wears - finger1's runner-up is only ever ranked against finger1's
// OWN list, which was never told what finger2 took. Trying that
// runner-up alone, without checking this, would equip the same ring
// in both finger slots at once during verification: a real, invalid
// gear list (sim/bulk/expand.go's valid() refuses exactly this shape)
// that would silently double the item's stats and misreport the
// runner-up as stronger than it is.
var pairSlot = map[string]string{
	"finger1":  "finger2",
	"finger2":  "finger1",
	"trinket1": "trinket2",
	"trinket2": "trinket1",
}

// swapSlot rebuilds the gear list with one slot's item replaced by
// itemID, applying the same two-hand/off-hand rule buildGear does. A
// main_hand swap checks the ITEM BEING SWAPPED IN for two-handedness
// (the runner-up may be two-handed even when the pick was not); every
// other slot's swap checks the pick's own main_hand, unchanged. If
// the swapped-in item is also worn in this slot's pair-mate (see
// pairSlot), the pair-mate is dropped from the trial rather than
// worn twice.
func swapSlot(picks map[string]slotPick, slot string, itemID int, itemIsTwoHand bool) []api.GearSlot {
	twoHand := picks["main_hand"].Item != nil && picks["main_hand"].Item.TwoHand
	if slot == "main_hand" {
		twoHand = itemIsTwoHand
	}
	dropPairMate := false
	if mate, ok := pairSlot[slot]; ok && picks[mate].Item != nil && picks[mate].Item.ID == itemID {
		dropPairMate = true
	}
	var out []api.GearSlot
	for _, s := range slotOrder {
		if s == "off_hand" && twoHand {
			continue
		}
		if dropPairMate && s == pairSlot[slot] {
			continue
		}
		id := 0
		switch {
		case s == slot:
			id = itemID
		case picks[s].Item != nil:
			id = picks[s].Item.ID
		}
		if id == 0 {
			continue
		}
		out = append(out, api.GearSlot{Slot: s, ItemID: id})
	}
	return out
}

// verifyBand runs the baseline set once, then every slot with a
// runner-up once more (with just that slot swapped), and reports
// which runner-ups beat the pick.
//
// A single slot's swap sim can fail for reasons that have nothing to
// do with which item scored higher - this lane's run hit two: an
// item-set bonus registered for the wrong class, and two picked
// items' proc auras colliding ("Aura X already registered!", the
// engine's own item_effects registering the same aura twice when two
// equipped items grant it). Both are engine-side content gaps a
// prototype ranking command cannot fix, and sim/core recovers the
// panic itself and reports it through RaidSimResult.Error rather than
// crashing the process (adapter.ResultError turns that into the
// ordinary Go error runPlainDPS returns) - so this function treats a
// per-slot swap error as "could not verify this one", logs it, and
// keeps going, rather than losing the whole band's report over one
// bad candidate. A baseline failure is different: with no baseline
// DPS there is nothing to compare a swap against, so that error
// still propagates and the caller reports the band as unverified.
func verifyBand(spec specInfo, race, classSlug string, level int, picks map[string]slotPick) (baselineDPS float64, swaps []swapResult, verifyErrors []string, err error) {
	baseline := plainRequest(spec, api.CharacterSpec{Name: "verify", Race: race, Class: classSlug, Level: level, Gear: buildGear(picks)}, verifyIterations, verifySeed)
	baselineDPS, err = runPlainDPS(baseline)
	if err != nil {
		return 0, nil, nil, err
	}

	var slots []string
	for slot := range picks {
		if picks[slot].RunnerUp != nil {
			slots = append(slots, slot)
		}
	}
	sort.Strings(slots)

	for _, slot := range slots {
		runnerUp := picks[slot].RunnerUp
		gear := swapSlot(picks, slot, runnerUp.ID, runnerUp.TwoHand)
		req := plainRequest(spec, api.CharacterSpec{Name: "verify", Race: race, Class: classSlug, Level: level, Gear: gear}, verifyIterations, verifySeed)
		dps, runErr := runPlainDPS(req)
		if runErr != nil {
			verifyErrors = append(verifyErrors, fmt.Sprintf("%s: runner-up %s (id %d): %v", slot, runnerUp.Name, runnerUp.ID, runErr))
			continue
		}
		swaps = append(swaps, swapResult{Slot: slot, SwapDPS: dps, Beat: dps > baselineDPS})
	}
	return baselineDPS, swaps, verifyErrors, nil
}
