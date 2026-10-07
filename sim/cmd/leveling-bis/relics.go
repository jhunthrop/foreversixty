package main

// Ranking relics (librams, idols, totems) by engine-verified DPS.
//
// A relic carries no stats; its whole value is an equip effect. score()
// cannot see one, so the old flow let pick() choose an arbitrary relic by
// id, found the engine could not simulate it, and published the slot empty
// with effect_not_modelled. Now that the engine models the DPS relics
// (sim/<class>/items.go, registered through core.NewEquipModItemEffect or a
// hand-written effect), the relic slot is ranked the way a trinket is: each
// modelled relic is equipped in turn, the highest measured DPS wins, and its
// gain over an empty relic slot is measured at escalating precision so
// report.go can refuse a relic whose simulated gain is noise.
//
// A relic whose effect the engine does not model never enters the
// tournament - measuring it would only measure an empty slot - and a slot
// with no modelled relic at all is left to report.go's effect_not_modelled.

// relicShortlist is every relic in pool whose effect a real sim of this
// build exercises (effectVerifiedInSim), best-known first as pool orders them.
func relicShortlist(pool []scored) []scored {
	var shortlist []scored
	for _, c := range pool {
		if isRelicCandidate(c.candidate) && effectVerifiedInSim(c.candidate) {
			shortlist = append(shortlist, c)
		}
	}
	return shortlist
}

// rankRelicSlot replaces picks[slot] with the best simulated modelled relic.
// A slot with no modelled relic is returned unchanged.
func rankRelicSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored, slot string) (map[string]slotPick, []string) {
	candidates := relicShortlist(bySlot[slot])
	if len(candidates) == 0 {
		return picks, nil
	}
	return rankByMeasuredGain(runner, spec, race, classSlug, level, talents, picks, slot, candidates)
}
