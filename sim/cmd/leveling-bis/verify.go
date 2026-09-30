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
	"math"
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
	// BaselineDPS is the scored set's own DPS the swap was measured against.
	BaselineDPS float64
	Beat        bool
	// SwapStdErr/BaselineStdErr are the two runs' own standard error
	// (api.Estimate.Error, RunPlainDPSWithError's second return) -
	// data-followups-10 lane, 2026-09-30, item 7: read for free off the
	// SAME single sim run RunPlainDPS already made (runPlainDPSEstimate's
	// own doc - no second run, no extra cost), so Significant (below)
	// can tell a genuine, reproducible DPS difference apart from this
	// run's own sampling noise the exact way trinkets.go's
	// trinketGainSignificant/positiveBeyondError already do for a
	// trinket's gain - this file's own swap comparison had no such
	// check at all before this lane (a live repro, hunter-marksmanship
	// band 20: Serpent Gloves/+7 spell power, a stat this spec's own
	// weight_stats never measures at all, so score() ties it exactly
	// against Gloves of the Fang's +4 Strength, also unweighted - kept
	// "winning" the item-level tie-break for three sweeps running, and
	// verify.go's own bare 1%-margin swap test against Gloves of the
	// Fang, same fixed verifySeed every run, kept reporting a "real",
	// sim-measured negative delta for it that never once retested
	// whether that gap was bigger than the run's own noise).
	SwapStdErr     float64
	BaselineStdErr float64
}

// Significant reports whether sw's own SwapDPS and BaselineDPS differ
// by more than their own combined standard error - the same
// combination-in-quadrature rule gainsIndistinguishable
// (faction_trinkets.go) and trinkets.go's own gainStdErr already use,
// applied here to an ordinary gear swap's baseline/runner-up pair
// instead of two trinkets' own gains. Computed on demand, never
// stored, so a swapResult literal built by a test that has not been
// updated to populate SwapStdErr/BaselineStdErr (the Go zero value,
// 0.0) trivially reports "significant" for ANY nonzero delta - this
// method is additive, never a silent behaviour change for a caller
// that predates real error bars. Beat (verifyBand's own flat
// swapMargin bar) and Significant are BOTH required before applySwaps
// promotes a runner-up (a real improvement must be both practically
// and statistically real) and before report.go treats this slot's own
// measured delta as confirmed evidence rather than an ordinary,
// unverified estimate.
func (sw swapResult) Significant() bool {
	combined := math.Sqrt(sw.SwapStdErr*sw.SwapStdErr + sw.BaselineStdErr*sw.BaselineStdErr)
	return math.Abs(sw.SwapDPS-sw.BaselineDPS) > combined
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

// pairSlot is each slot's partner that can never physically hold the
// same item at once: pick() only ever excludes a pair-mate's own
// id/name from the OTHER slot's candidate list going forward (finger1
// decided before finger2), so a slot's own runner-up can still be the
// exact item its pair-mate already wears - finger1's runner-up is
// only ever ranked against finger1's OWN list, which was never told
// what finger2 took. Trying that runner-up alone, without checking
// this, would equip the same ring in both finger slots at once during
// verification: a real, invalid gear list (sim/bulk/expand.go's
// valid() refuses exactly this shape) that would silently double the
// item's stats and misreport the runner-up as stronger than it is.
//
// main_hand/off_hand carries the identical risk for a
// leveling.DualWieldSpecs member: off_hand's own candidate pool is
// main_hand's one-handed weapons merged in (pick.go's own "off_hand"
// case), so the very same physical one-hander can be main_hand's
// runner-up while off_hand already wears it (found in this run's own
// dogfood: hunter-survival's level-35 horde list published Tok'kar's
// Murloc Shanker in BOTH hands after main_hand's runner-up swap
// promoted it, wouldDuplicatePairMate's own new test). A one-hand
// spec (shield or two-hand) never collides here: its off_hand pool
// (shields/held items, or none at all) shares no ids with main_hand's
// weapons.
var pairSlot = map[string]string{
	"finger1":   "finger2",
	"finger2":   "finger1",
	"trinket1":  "trinket2",
	"trinket2":  "trinket1",
	"main_hand": "off_hand",
	"off_hand":  "main_hand",
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
func verifyBand(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick) (baselineDPS float64, swaps []swapResult, verifyErrors []string, err error) {
	baseline := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, buildGear(picks)), verifyIterations, verifySeed)
	var baselineStdErr float64
	baselineDPS, baselineStdErr, err = runner.RunPlainDPSWithError(baseline)
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
		req := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, gear), verifyIterations, verifySeed)
		dps, stdErr, runErr := runner.RunPlainDPSWithError(req)
		if runErr != nil {
			verifyErrors = append(verifyErrors, fmt.Sprintf("%s: runner-up %s (id %d): %v", slot, runnerUp.Name, runnerUp.ID, runErr))
			continue
		}
		sw := swapResult{
			Slot:           slot,
			SwapDPS:        dps,
			BaselineDPS:    baselineDPS,
			SwapStdErr:     stdErr,
			BaselineStdErr: baselineStdErr,
		}
		sw.Beat = beatsByMargin(dps, baselineDPS) && sw.Significant()
		swaps = append(swaps, sw)
	}
	return baselineDPS, swaps, verifyErrors, nil
}

// wouldDuplicatePairMate reports whether promoting itemID/itemName
// into slot would leave it wearing the exact same physical item (by
// id, or by name for a lower/higher-quality reprint) as its own
// finger/trinket pair-mate already does in out -- pick()'s and
// rankTrinketSlot's own pairing rule (excludePaired, topByItemLevel),
// which this file's own swap trial can otherwise defeat: swapSlot
// already drops the pair-mate from the TRIAL GEAR so the sim measures
// one ring/trinket cleanly (this file's own doc above swapSlot), but
// that trial result is "wearing one fewer paired item than the
// baseline had", not "this item in this slot is an improvement" -- a
// slot's own runner-up is only ever ranked against ITS OWN pool at
// the time it was ranked (pick.go/trinkets.go), which for finger1 in
// particular runs before finger2's final item is known, so finger1's
// runner-up can freely be whatever finger2 later becomes. Promoting
// it anyway would equip the one physical item in both slots at once
// -- a real, invalid gear list (bulk/expand.go's valid() refuses
// exactly this shape) and the literal "same trinket twice" defect
// this lane's brief names as a published output the owner caught by
// eye (mage-arcane/mage-frost/priest-shadow/shaman-elemental/warlock-
// destruction's finger or trinket pair, this run).
func wouldDuplicatePairMate(out map[string]slotPick, slot string, itemID int, itemName string) bool {
	mate, ok := pairSlot[slot]
	if !ok || out[mate].Item == nil {
		return false
	}
	return out[mate].Item.ID == itemID || out[mate].Item.Name == itemName
}

// wouldBreakTwoHandInvariant is whether promoting runnerUp into
// off_hand would leave it equipped under a two-handed main hand -
// pick.go's own enforceTwoHandOffHandInvariant rule, re-checked here
// because a main_hand promotion earlier in this SAME applySwaps call
// can make the character two-handed after off_hand's own swap was
// already measured against the old (one-handed) main_hand.
//
// A main_hand promotion TO a two-hander is not refused here (it used
// to be, and that used to be this function's whole job): applySwaps
// itself now empties off_hand as PART of that promotion instead (this
// lane's brief, defect 3) - the promoted swap's own SwapDPS was
// already measured with off_hand correctly dropped (verify.go's own
// swapSlot), so refusing the promotion outright would silently keep a
// weaker dual-wield pick even when the sim had already measured the
// two-hander beating it. The bug this function used to guard against
// (the nightly of 2026-09-29 published Darkwood Staff beside Nightglow
// Concoction for priest-shadow band 20) was promoting a two-hander
// WITHOUT dropping off_hand at all; the fix now is to drop it
// properly, not to refuse the promotion.
func wouldBreakTwoHandInvariant(out map[string]slotPick, slot string) bool {
	if slot == "off_hand" {
		return out["main_hand"].Item != nil && out["main_hand"].Item.TwoHand
	}
	return false
}

// swapMargin is how much a runner-up's measured set DPS must exceed the
// scored pick's before the sim's verdict overrides the score. The
// verify runs are short, so two items within a fraction of a percent
// are a tie the noise decides: the nightly of 2026-09-29 replaced
// Serpent's Shoulders (+5 agility) with Mantle of Honor (+7 intellect,
// +7 spirit) on the Alliance level-20 hunter list on 78.7 vs 78.6 DPS.
// A tie keeps the scored pick.
const swapMargin = 0.01

// beatsByMargin is whether dps beats baseline by more than swapMargin.
func beatsByMargin(dps, baseline float64) bool {
	return dps > baseline*(1+swapMargin)
}

// applySwaps promotes every runner-up that beat its slot's scored pick
// into the pick (the scored pick becomes the row's runner-up, so the
// report can say what was beaten), then measures the resulting set once
// more so the published set DPS is the set's own, not the pre-swap
// baseline. With no swap that beat, picks and setDPS come back as they
// were and no sim runs.
//
// The returned []swapResult is swaps with Beat forced false for any
// slot wouldDuplicatePairMate rejected, so report.go's own swap_note
// (built from this same slice) never claims a promotion that did not
// happen -- see that function's own doc for why silently declining a
// duplicate here but leaving the original swaps slice unchanged would
// print a swap_note for an item the row no longer shows.
func applySwaps(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, swaps []swapResult, setDPS float64) (map[string]slotPick, float64, []swapResult, error) {
	promoted := false
	out := make(map[string]slotPick, len(picks))
	for slot, pk := range picks {
		out[slot] = pk
	}
	adjusted := make([]swapResult, len(swaps))
	for i, sw := range swaps {
		adjusted[i] = sw
		pk, ok := out[sw.Slot]
		if !sw.Beat || !ok || pk.RunnerUp == nil {
			continue
		}
		if wouldDuplicatePairMate(out, sw.Slot, pk.RunnerUp.ID, pk.RunnerUp.Name) || wouldBreakTwoHandInvariant(out, sw.Slot) {
			adjusted[i].Beat = false
			continue
		}
		// This lane's brief, item 7: the promoted item's own MeasuredDPS
		// (scored's own doc) is sw.SwapDPS, the real measurement that
		// just won this comparison - report.go's buildReport reads it to
		// publish sim_dps instead of score()'s stat estimate for this
		// row. promotedItem is a copy (not pk.RunnerUp itself) so this
		// never mutates a scored value some OTHER slot's bySlot entry
		// might still be reading.
		promotedItem := *pk.RunnerUp
		promotedItem.MeasuredDPS = sw.SwapDPS
		out[sw.Slot] = slotPick{Item: &promotedItem, RunnerUp: pk.Item}
		promoted = true
		if sw.Slot == "main_hand" && pk.RunnerUp.TwoHand {
			// Equipping a two-hander physically empties the off hand -
			// the same rule buildGear/swapSlot already apply to the
			// gear list this swap's own SwapDPS was measured against
			// (this file's own doc above). Clearing the SLOT ITSELF
			// here (not just one trial's gear list) is the fix this
			// lane's brief (defect 3) asks for in place of
			// wouldBreakTwoHandInvariant's old outright refusal.
			out["off_hand"] = slotPick{}
		}
	}
	if !promoted {
		return picks, setDPS, adjusted, nil
	}
	final := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, buildGear(out)), verifyIterations, verifySeed)
	dps, err := runner.RunPlainDPS(final)
	if err != nil {
		return nil, 0, nil, err
	}
	return out, dps, adjusted, nil
}
