package main

import (
	"fmt"
	"log"

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
	// band is ranked under it too (band_pipeline.go chooseRun).
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

// gearedBandWeights is the weights a band ranks by: a sweep of the ladder
// character (the bare, armed character ladderCharacter builds), then, when
// that sweep is trusted, a second sweep of the same character wearing the
// set those ladder weights pick (firstPass). The geared sweep is returned
// when it is itself trusted, carrying the ladder sweep it replaced
// (bandWeights.ladder): runSpec ranks the band under both and publishes the
// better verified set (band_pipeline.go chooseRun), because a geared sweep
// reads the margins of one set and the picks can move away from it.
//
// Round 8 found the bare character misprices every stat a geared one
// carries: its small mana pool overvalues Intellect for every mana-using
// physical spec (Retribution band 60: 6.7 attack power per point bare, 3.4
// on its own set; Survival band 40: 3.3 against 1.4), and without gear's
// hit, crit and attack power it overvalues those too (Fury band 60's hit:
// 38 against 22). The ranker's own picks then lost to the previous published
// sets by 2 to 8% (the incumbent guard kept them).
func gearedBandWeights(runner engineRunner, spec specInfo, band, iterations int, anchorStat string, ladder api.CharacterSpec, firstPass firstPassGear) (bandWeights, error) {
	bare, err := measureBandWeights(runner, spec, band, iterations, anchorStat, ladder)
	if err != nil || bare.reason != "" {
		return bare, err
	}
	gear := wearOver(ladder.Gear, firstPass(effectiveWeights(bare.result), bare.referenceDPSPerPoint))
	if gearKeyOf(gear) == gearKeyOf(ladder.Gear) {
		return bare, nil
	}
	geared := ladder
	geared.Gear = gear
	worn, err := measureBandWeights(runner, spec, band, iterations, anchorStat, geared)
	if err != nil {
		return bandWeights{}, err
	}
	if worn.reason != "" {
		log.Printf("leveling-bis: %s band %d: the geared weights sweep was not significant (%s); ranking by the ladder sweep", spec.Spec, band, worn.reason)
		return bare, nil
	}
	worn.ladder = &bare
	return worn, nil
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
