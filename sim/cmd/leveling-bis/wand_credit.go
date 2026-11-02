package main

import (
	"fmt"
	"sort"

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

// wandTournamentSize is how many wands, best score first, rankWandSlot
// measures in the full set.
const wandTournamentSize = 3

// rankWandSlot measures the top wandTournamentSize wands by score, each worn
// in the band's set, and keeps the best measured as the ranged pick (the
// next best its runner-up). A wand's score rests on the measured wand credit,
// and verify only swaps in the runner-up, so a wand the score put third was
// never measured (round 9: Shadow band 40 kept Goblin Igniter over Twisted
// Nether Wand, 3.6 DPS better). A ranged pick that is not a wand, or a slot
// with one wand, is left alone. It returns a new map.
func rankWandSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored) (map[string]slotPick, []string) {
	out := clonePicksMap(picks)
	if wornWand(picks) == nil {
		return out, nil
	}
	var wands []scored
	for _, c := range bySlot["ranged"] {
		if c.ClassID == itemClassWeapon && c.SubclassID == wandSubclassID && c.DPS > 0 {
			wands = append(wands, c)
		}
		if len(wands) == wandTournamentSize {
			break
		}
	}
	if len(wands) < 2 {
		return out, nil
	}
	type measured struct {
		item scored
		dps  float64
	}
	var results []measured
	var notes []string
	for _, w := range wands {
		req := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, swapSlot(picks, "ranged", w.ID, false)), verifyIterations, verifySeed)
		dps, err := runner.RunPlainDPS(req)
		if err != nil {
			notes = append(notes, fmt.Sprintf("ranged: wand %s (id %d) could not be measured: %v", w.Name, w.ID, err))
			continue
		}
		results = append(results, measured{item: w, dps: dps})
	}
	if len(results) == 0 {
		return out, notes
	}
	sort.SliceStable(results, func(i, j int) bool { return results[i].dps > results[j].dps })
	best := results[0].item
	best.MeasuredDPS = results[0].dps
	pick := slotPick{Item: &best}
	if len(results) > 1 {
		runnerUp := results[1].item
		pick.RunnerUp = &runnerUp
	}
	out["ranged"] = pick
	if best.ID != picks["ranged"].Item.ID {
		notes = append(notes, fmt.Sprintf("ranged: %s measured best of the top %d wands (%.1f), over the scored %s", best.Name, len(results), results[0].dps, picks["ranged"].Item.Name))
	}
	return out, notes
}
