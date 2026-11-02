package main

import (
	"fmt"
	"log"
	"math"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// bandWeights is one weights sweep of a band's character, after the
// reference-stat retry and the primary-anchor check: what runSpec ranks by,
// publishes and falls back from.
type bandWeights struct {
	request              api.SimRequest
	result               map[string]api.StatWeight
	referenceDPSPerPoint float64
	// reason is referenceMeasurementReason's verdict: "" when the sweep
	// measured the reference stat positive beyond its error.
	reason        string
	lowConfidence bool
	// ladder is the bare ladder sweep when a geared sweep replaced it: the
	// set its weights pick is one more set the band's guard measures
	// (ladderWeightsContender).
	ladder *bandWeights
}

// measureBandWeights sweeps ch's stat weights. A sweep whose reference stat
// is not positive beyond its own error re-runs once at
// weightsRetryIterationsFactor iterations (warlock-destruction band 60's
// repro, bis-ranker-integrity-11: the retry tightens the error but does not
// always flip a negative measurement, so runSpec's lower-band fallback still
// exists). A trusted sweep whose primary-anchor row came back insignificant
// re-measures that row once (primaryStatSignificanceCheck); an untrusted
// band publishes no anchor, so it skips that retry.
func measureBandWeights(runner engineRunner, spec specInfo, band, iterations int, anchorStat string, ch api.CharacterSpec) (bandWeights, error) {
	req := weightsRequest(spec, ch, iterations, 3)
	result, reference, err := runner.RunWeights(req)
	if err != nil {
		return bandWeights{}, fmt.Errorf("band %d weights run: %w", band, err)
	}
	reason := referenceMeasurementReason(spec.ReferenceStat, result, reference)
	if reason != "" {
		retryReq := weightsRequest(spec, ch, iterations*weightsRetryIterationsFactor, 3)
		retryResult, retryReference, retryErr := runner.RunWeights(retryReq)
		if retryErr != nil {
			return bandWeights{}, fmt.Errorf("band %d weights retry run: %w", band, retryErr)
		}
		log.Printf("leveling-bis: %s band %d: sweep at %d iterations/direction was not significant (%s); re-ran at %dx", spec.Spec, band, iterations, reason, weightsRetryIterationsFactor)
		result, reference = retryResult, retryReference
		reason = referenceMeasurementReason(spec.ReferenceStat, result, reference)
	}
	out := bandWeights{request: req, result: result, referenceDPSPerPoint: reference, reason: reason}
	if reason != "" || anchorStat == "" {
		return out, nil
	}
	retried, lowConf, err := primaryStatSignificanceCheck(runner, spec, ch, anchorStat, iterations, 3, result)
	if err != nil {
		return bandWeights{}, fmt.Errorf("band %d primary-stat (%s) weights retry: %w", band, anchorStat, err)
	}
	if lowConf {
		log.Printf("leveling-bis: %s band %d: primary stat %s still not significant after a %dx retry; publishing the anchor anyway (weights_low_confidence)", spec.Spec, band, anchorStat, primaryStatRetryIterationsFactor)
	}
	out.result, out.lowConfidence = retried, lowConf
	return out, nil
}

// firstPassGear is the set a band's weights pick, as the gear a geared
// weights sweep wears.
type firstPassGear func(weights map[string]float64, referenceDPSPerPoint float64) []api.GearSlot

// setMeasure sims one gear set: its score (DPS, healing or tank score, as
// the band's runner reads it) and standard error.
type setMeasure func(gear []api.GearSlot) (score, stdErr float64, err error)

// gearedBandWeights is the weights a band ranks by: a sweep of the ladder
// character (the bare, armed character ladderCharacter builds), then, when
// that sweep is trusted, a second sweep of the same character wearing the
// set those ladder weights pick (firstPass). The geared sweep replaces the
// ladder one when it is itself trusted, unless the set the geared weights
// pick measures worse beyond error than the ladder weights' set (measure):
// a geared sweep reads one set's margins, and a pick that moves far from
// that set (a caster that sheds the Intellect the sweep wore) can land
// where those margins no longer hold.
//
// Round 8 found the bare character misprices every stat a geared one
// carries: its small mana pool overvalues Intellect for every mana-using
// physical spec (Retribution band 60: 6.7 attack power per point bare, 3.4
// on its own set; Survival band 40: 3.3 against 1.4), and without gear's
// hit, crit and attack power it overvalues those too (Fury band 60's hit:
// 38 against 22). The ranker's own picks then lost to the previous published
// sets by 2 to 8% (the incumbent guard kept them).
func gearedBandWeights(runner engineRunner, spec specInfo, band, iterations int, anchorStat string, ladder api.CharacterSpec, firstPass firstPassGear, measure setMeasure) (bandWeights, error) {
	bare, err := measureBandWeights(runner, spec, band, iterations, anchorStat, ladder)
	if err != nil || bare.reason != "" {
		return bare, err
	}
	bareSet := wearOver(ladder.Gear, firstPass(effectiveWeights(bare.result), bare.referenceDPSPerPoint))
	if gearKeyOf(bareSet) == gearKeyOf(ladder.Gear) {
		return bare, nil
	}
	geared := ladder
	geared.Gear = bareSet
	worn, err := measureBandWeights(runner, spec, band, iterations, anchorStat, geared)
	if err != nil {
		return bandWeights{}, err
	}
	if worn.reason != "" {
		log.Printf("leveling-bis: %s band %d: the geared weights sweep was not significant (%s); ranking by the ladder sweep", spec.Spec, band, worn.reason)
		return bare, nil
	}
	wornSet := wearOver(ladder.Gear, firstPass(effectiveWeights(worn.result), worn.referenceDPSPerPoint))
	if gearKeyOf(wornSet) == gearKeyOf(bareSet) {
		return worn, nil
	}
	worn.ladder = &bare
	bareScore, bareErr, err := measure(bareSet)
	if err != nil {
		return bandWeights{}, fmt.Errorf("band %d measuring the ladder weights' set: %w", band, err)
	}
	wornScore, wornErr, err := measure(wornSet)
	if err != nil {
		return bandWeights{}, fmt.Errorf("band %d measuring the geared weights' set: %w", band, err)
	}
	if bareScore-wornScore > math.Hypot(bareErr, wornErr) {
		log.Printf("leveling-bis: %s band %d: the geared weights' set measures %.1f +/- %.1f against the ladder weights' %.1f +/- %.1f; ranking by the ladder sweep", spec.Spec, band, wornScore, wornErr, bareScore, bareErr)
		return bare, nil
	}
	return worn, nil
}

// ladderWeightsContender is the set the ladder sweep's weights pick for a
// faction, wearing the trinkets the band's final set earned (pick() places
// trinkets by score only; the trinket tournament decides them): the guard
// (keepVerifyBest) measures it beside the band's final set, so ranking by a
// geared sweep never publishes a set worse beyond error than ranking by the
// ladder sweep would have. Round 8: Balance band 50's geared weights picked
// a first pass no worse than the ladder's, yet its final set measured 0.8%
// (Alliance) and 1.9% (Horde) below the ladder-weighted run's.
func ladderWeightsContender(ladderPicks, final map[string]slotPick) map[string]slotPick {
	out := clonePicksMap(ladderPicks)
	out["trinket1"] = final["trinket1"]
	out["trinket2"] = final["trinket2"]
	return out
}

// wearOver is set worn over base: every slot set fills, plus base's items
// in the slots set leaves empty (the ladder's weapons beside a first pass
// that found none), in slotOrder. A two-handed main hand from set clears
// base's off hand.
func wearOver(base, set []api.GearSlot) []api.GearSlot {
	bySlot := make(map[string]int, len(base)+len(set))
	for _, g := range base {
		bySlot[g.Slot] = g.ItemID
	}
	setSlots := make(map[string]bool, len(set))
	for _, g := range set {
		bySlot[g.Slot] = g.ItemID
		setSlots[g.Slot] = true
	}
	if setSlots["main_hand"] && !setSlots["off_hand"] {
		delete(bySlot, "off_hand")
	}
	out := make([]api.GearSlot, 0, len(bySlot))
	for _, slot := range slotOrder {
		if id, ok := bySlot[slot]; ok {
			out = append(out, api.GearSlot{Slot: slot, ItemID: id})
		}
	}
	return out
}

// gearKeyOf is a gear list's identity, slot by slot.
func gearKeyOf(gear []api.GearSlot) string {
	key := ""
	for _, g := range gear {
		key += fmt.Sprintf("%s=%d;", g.Slot, g.ItemID)
	}
	return key
}
