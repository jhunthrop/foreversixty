package main

import (
	"fmt"

	"github.com/wowsims/classic/sim/core/proto"
)

// wandSubclassID is the client's ItemSubclassWeapon id for a wand.
const wandSubclassID = 19

// wornWand is the wand picks wears in its ranged slot, or nil.
func wornWand(picks map[string]slotPick) *scored {
	w := picks["ranged"].Item
	if w == nil || w.ClassID != itemClassWeapon || w.SubclassID != wandSubclassID || w.DPS <= 0 {
		return nil
	}
	return w
}

// shootDPS is the Shoot damage per second in a player's metrics: the
// player's own DPS times Shoot's share of the player's own damage. Pets are
// separate units and are in neither figure.
func shootDPS(player *proto.UnitMetrics) float64 {
	if player == nil {
		return 0
	}
	var shoot, total float64
	for _, a := range player.GetActions() {
		var damage float64
		for _, t := range a.GetTargets() {
			damage += t.GetDamage()
		}
		total += damage
		if a.GetId().GetSpellId() == shootSpellID {
			shoot += damage
		}
	}
	if total <= 0 {
		return 0
	}
	return player.GetDps().GetAvg() * shoot / total
}

// measureWandCredit sims the band's set as picked (the verification
// harness: verifyIterations, verifySeed) and returns wandScorePerDPS for the
// wand it wears. A set with no wand runs nothing and returns 0.
func measureWandCredit(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, referenceDPSPerPoint float64) (float64, error) {
	wand := wornWand(picks)
	if wand == nil || referenceDPSPerPoint <= 0 {
		return 0, nil
	}
	req := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, buildGear(picks)), verifyIterations, verifySeed)
	shoot, err := runner.RunShootDPS(req)
	if err != nil {
		return 0, fmt.Errorf("measuring the Shoot damage of the band's set: %w", err)
	}
	return wandScorePerDPS(shoot, wand.DPS, referenceDPSPerPoint), nil
}
