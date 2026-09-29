package main

// Set bonuses (this lane's brief, item 3's second half): "when a pick
// would complete a 2- or 3-piece set the engine implements, run the
// verify with the set together... and keep the set if the verified
// DPS wins." score() sums each item's own stats independently
// (score.go) and has no notion that two items worn together unlock a
// THIRD thing neither carries alone -- a set bonus is invisible to it
// exactly the same way a proc is (rank.go's own doc).

import (
	"fmt"
	"sort"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// setCandidate is one slot's best score()-based candidate that
// belongs to an engine-implemented set.
type setCandidate struct {
	slot string
	item scored
}

// bestSetPieces groups bySlot's own #1-by-score candidate per slot by
// SetID, keeping only sets the engine actually implements
// (engineImplementedSetIDs) -- trinket1/trinket2 are excluded (they
// are ranked by rankTrinketSlot on a completely different axis, item
// level, not score, so "the #1 scored trinket" is not a meaningful
// question) and a slot with zero candidates is skipped.
//
// Two defences mirrored from pick.go/rank.go, needed here for the
// same reason: this function walks bySlot's OWN pools, not pick()'s
// already-defended output.
//
//   - A dual-wielder's main hand never offers a two-hander (see
//     rank.go's rankSlotWithEffects, which excludes the same way):
//     without this, a two-hand set piece could beat this spec's own
//     one-handed pick here and be tried together with an off-hand
//     item trySetCompletion still wears.
//   - finger1 and finger2 (data.go's plannerSlots: a ring's Slots is
//     always both) share the exact same candidate list, so their own
//     #1-scored item is the identical physical ring for both --
//     iterated in slotOrder (not map order, so which slot keeps it is
//     deterministic across runs) and deduplicated per set by item id,
//     rather than counting one ring as two of its own set's pieces
//     and having trySetCompletion equip it in both finger slots at
//     once (the same "same trinket twice" shape the trinket exclusion
//     above already avoids for the other paired slot).
func bestSetPieces(bySlot map[string][]scored, specSlug string) map[int][]setCandidate {
	out := map[int][]setCandidate{}
	seenItem := map[int]map[int]bool{}
	for _, slot := range slotOrder {
		list := bySlot[slot]
		if slot == "trinket1" || slot == "trinket2" || len(list) == 0 {
			continue
		}
		if slot == "main_hand" && leveling.DualWieldSpecs[specSlug] {
			list = excludeTwoHand(list)
			if len(list) == 0 {
				continue
			}
		}
		best := list[0]
		if best.SetID == nil || !setEffectImplemented(*best.SetID) {
			continue
		}
		if seenItem[*best.SetID] == nil {
			seenItem[*best.SetID] = map[int]bool{}
		}
		if seenItem[*best.SetID][best.ID] {
			continue
		}
		seenItem[*best.SetID][best.ID] = true
		out[*best.SetID] = append(out[*best.SetID], setCandidate{slot: slot, item: best})
	}
	return out
}

// setAlreadyFullyEquipped reports whether picks already wears every
// one of cands' items in exactly the slots cands names -- the case
// where trying this set again would just re-verify the current picks
// against themselves.
func setAlreadyFullyEquipped(picks map[string]slotPick, cands []setCandidate) bool {
	for _, c := range cands {
		if picks[c.slot].Item == nil || picks[c.slot].Item.ID != c.item.ID {
			return false
		}
	}
	return true
}

// trySetCompletion looks at every engine-implemented set with 2 or 3
// of its pieces among bySlot's own top-scored candidates, tries
// equipping every one of those pieces together (replacing whatever
// picks currently occupy those slots), and keeps the trial only when
// its own verify sim beats a same-budget verify of the picks it
// started from. Sets are tried independently against the ORIGINAL
// picks (not against each other's trial results): two different sets
// competing for the same slot is not a case this build's leveling
// item pool produces (a set piece belongs to exactly one set), so
// this does not need to arbitrate between them.
//
// Returns the possibly-updated picks, and any per-set verify errors
// (logged by the caller, not fatal -- the same resilience
// rankTrinketSlot/rankSlotWithEffects already give a single bad
// candidate).
func trySetCompletion(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored) (map[string]slotPick, []string) {
	bySet := bestSetPieces(bySlot, spec.Spec)
	setIDs := make([]int, 0, len(bySet))
	for id, cands := range bySet {
		if len(cands) < 2 || setAlreadyFullyEquipped(picks, cands) {
			continue
		}
		setIDs = append(setIDs, id)
	}
	if len(setIDs) == 0 {
		// Nothing to try: no implemented set has 2+ candidate pieces
		// this band's picks do not already wear, so spending a verify
		// run on a baseline nobody will compare against would be pure
		// nightly-budget waste.
		return picks, nil
	}
	sort.Ints(setIDs)

	baselineReq := plainRequest(spec, bandCharacter("set-completion-baseline", race, classSlug, level, talents, buildGear(picks)), trinketRankIterations, verifySeed)
	baselineDPS, err := runner.RunPlainDPS(baselineReq)
	if err != nil {
		return picks, []string{fmt.Sprintf("set completion: baseline verify failed: %v", err)}
	}

	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}
	var notes []string

	for _, setID := range setIDs {
		cands := bySet[setID]

		trial := make(map[string]slotPick, len(picks))
		for k, v := range picks {
			trial[k] = v
		}
		for _, c := range cands {
			item := c.item
			trial[c.slot] = slotPick{Item: &item}
		}

		req := plainRequest(spec, bandCharacter("set-completion", race, classSlug, level, talents, buildGear(trial)), trinketRankIterations, verifySeed)
		dps, err := runner.RunPlainDPS(req)
		if err != nil {
			notes = append(notes, fmt.Sprintf("set %d completion (%d pieces): verify failed: %v", setID, len(cands), err))
			continue
		}
		if beatsByMargin(dps, baselineDPS) {
			notes = append(notes, fmt.Sprintf("set %d completion (%d pieces) beat the independently-scored picks: %.1f vs %.1f set DPS - adopted", setID, len(cands), dps, baselineDPS))
			out = trial
			baselineDPS = dps
		}
	}
	return out, notes
}
