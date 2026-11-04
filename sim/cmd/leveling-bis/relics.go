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
	out, notes := rankByMeasuredGain(runner, spec, race, classSlug, level, talents, picks, slot, candidates)
	return preferTouchingRelic(out, slot), notes
}

// preferTouchingRelic swaps a relic slot's winner for its runner-up when the
// winner measured exactly the empty slot (it touches nothing the rotation
// casts) and the runner-up is a near tie (relicNearTie): on a sub-DPS race
// the noisy ordering can crown the inert relic, which would then publish
// empty while a relic that does something sits right behind it.
func preferTouchingRelic(picks map[string]slotPick, slot string) map[string]slotPick {
	sp := picks[slot]
	if sp.Item == nil || sp.RunnerUp == nil || !sp.Item.GainMeasured || sp.Item.MeasuredGainDPS != 0 || !relicNearTie(sp.RunnerUp) {
		return picks
	}
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}
	winner, inert := *sp.RunnerUp, *sp.Item
	sp.Item, sp.RunnerUp = &winner, &inert
	out[slot] = sp
	return out
}

// nearTie is a published relic near tie's measured gain over an empty relic
// slot and the error bar it sits inside (trinketGainSignificanceMultiplier
// times its standard error, the bar a significant gain must clear), so the
// page can show how small it is.
type nearTie struct {
	GainDPS float64 `json:"gain"`
	Error   float64 `json:"error"`
}

// relicNearTie reports whether a modelled relic whose measured gain over an
// empty relic slot is inside its own error still publishes, as a near tie
// (report.go's NearTie). A relic has no stats to lose and every class that
// wears one wears one, so an empty slot is the wrong answer when the relic
// did something: its gain is not exactly zero (the candidate and the empty
// slot run on the same seed, so a relic whose effect reaches no spell the
// rotation casts reproduces the empty slot's DPS to the last digit:
// Enhancement's Totem of Urgency and Burning Totem, round 18 shaman S3) and
// not significantly below zero. trinketGainSignificant's 2x bar on two
// unpaired 100-iteration runs is about 1.4 DPS on a raid set, which no
// Enhancement totem (Totem of the Storm about +0.8, Totem of Thunder +0.24)
// can clear.
func relicNearTie(item *scored) bool {
	if item == nil || !isRelicCandidate(item.candidate) || !item.GainMeasured {
		return false
	}
	gain, bar := item.MeasuredGainDPS, trinketGainSignificanceMultiplier*item.MeasuredGainStdErr
	return gain != 0 && gain > -bar && !trinketGainSignificant(gain, item.MeasuredGainStdErr)
}
