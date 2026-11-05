package main

import (
	"fmt"
	"math"
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
// Nether Wand, 3.6 DPS better). Wands that measure inside the combined error
// of the leader at the verify harness are measured again at the guard
// harness, and the order there decides (round 24: Torch of Light led
// Brilliant Wand by 0.07 DPS at 300 iterations on Alliance Destruction raid,
// against errors of about 2, and is 0.3% behind it). A ranged pick that is
// not a wand, or a slot with one wand, is left alone. It returns a new map.
func rankWandSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored) (map[string]slotPick, []string) {
	out := clonePicksMap(picks)
	results, notes := rankWands(runner, spec, race, classSlug, level, talents, picks, bySlot)
	if len(results) == 0 {
		return out, notes
	}
	best := results[0].item
	best.MeasuredDPS = results[0].dps
	pick := slotPick{Item: &best}
	if len(results) > 1 {
		runnerUp := results[1].item
		runnerUp.MeasuredDPS = results[1].dps
		pick.RunnerUp = &runnerUp
	}
	out["ranged"] = pick
	if best.ID != picks["ranged"].Item.ID {
		notes = append(notes, fmt.Sprintf("ranged: %s measured best of the top %d wands (%.1f), over the scored %s", best.Name, len(results), results[0].dps, picks["ranged"].Item.Name))
	}
	return out, notes
}

// rankWands is rankWandSlot's measured order of the top wands, best first:
// nil when the set wears no wand or the slot offers fewer than two.
func rankWands(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored) ([]wandReading, []string) {
	if wornWand(picks) == nil {
		return nil, nil
	}
	wands := topWands(bySlot["ranged"])
	if len(wands) < 2 {
		return nil, nil
	}
	results, notes := measureWands(runner, spec, race, classSlug, level, talents, picks, wands, verifyIterations)
	if len(results) == 0 {
		return nil, notes
	}
	results, guardNotes := settleWandTies(runner, spec, race, classSlug, level, talents, picks, results)
	return results, append(notes, guardNotes...)
}

// wandReading is one wand's measured set DPS and that run's standard error.
type wandReading struct {
	item   scored
	dps    float64
	stdErr float64
}

// topWands is the first wandTournamentSize real wands of a score-ordered
// ranged pool.
func topWands(pool []scored) []scored {
	var wands []scored
	for _, c := range pool {
		if c.ClassID == itemClassWeapon && c.SubclassID == wandSubclassID && c.DPS > 0 {
			wands = append(wands, c)
		}
		if len(wands) == wandTournamentSize {
			break
		}
	}
	return wands
}

// measureWands sims the band's set wearing each wand at iterations
// (verifySeed) and returns the readings, best first.
func measureWands(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, wands []scored, iterations int) ([]wandReading, []string) {
	var results []wandReading
	var notes []string
	for _, w := range wands {
		req := plainRequest(spec, bandCharacter("verify", race, classSlug, spec.Spec, level, talents, swapSlot(picks, "ranged", w.ID, false)), iterations, verifySeed)
		dps, stdErr, err := runner.RunPlainDPSWithError(req)
		if err != nil {
			notes = append(notes, fmt.Sprintf("ranged: wand %s (id %d) could not be measured: %v", w.Name, w.ID, err))
			continue
		}
		results = append(results, wandReading{item: w, dps: dps, stdErr: stdErr})
	}
	sort.SliceStable(results, func(i, j int) bool { return results[i].dps > results[j].dps })
	return results, notes
}

// settleWandTies re-measures the leader and every wand inside the combined
// error of it at the guard harness (guardIterations), and puts them first in
// the guard order. Wands outside the error keep their place behind them. A
// guard run that fails leaves the verify order as it was.
func settleWandTies(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, results []wandReading) ([]wandReading, []string) {
	lead := results[0]
	tied := []scored{lead.item}
	var rest []wandReading
	for _, r := range results[1:] {
		if lead.dps-r.dps <= math.Hypot(lead.stdErr, r.stdErr) {
			tied = append(tied, r.item)
		} else {
			rest = append(rest, r)
		}
	}
	if len(tied) < 2 {
		return results, nil
	}
	guarded, notes := measureWands(runner, spec, race, classSlug, level, talents, picks, tied, guardIterations(spec))
	if len(guarded) != len(tied) {
		return results, notes
	}
	return append(guarded, rest...), notes
}

// wornWandReading is the worn wand's DPS in the same tournament the best
// wand was measured in, when the tournament kept it as the runner-up (a tie
// settled at the guard harness has both at the guard count), else the set's
// own fallback figure.
func wornWandReading(pick slotPick, worn *scored, fallback float64) float64 {
	if ru := pick.RunnerUp; ru != nil && ru.ID == worn.ID && ru.MeasuredDPS > 0 {
		return ru.MeasuredDPS
	}
	return fallback
}

// finalWandCheck re-runs the wand tournament on the band's final set, after
// verify and set completion have changed the slots around the wand, and
// adopts a wand that measures ahead of the worn one by swapMargin. The
// band's set DPS becomes that measurement (the same harness verify uses),
// and the ranged slot's verify swap, which no longer describes the pick, is
// dropped.
func finalWandCheck(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, vb verifiedBand, bySlot map[string][]scored) (verifiedBand, []string) {
	current := vb.picks["ranged"].Item
	if current == nil {
		return vb, nil
	}
	ranked, notes := rankWandSlot(runner, spec, race, classSlug, level, talents, vb.picks, bySlot)
	best := ranked["ranged"].Item
	if best == nil || best.ID == current.ID || !beatsByMargin(best.MeasuredDPS, wornWandReading(ranked["ranged"], current, vb.setDPS)) {
		return vb, nil
	}
	out := vb
	out.picks = ranked
	out.setDPS = best.MeasuredDPS
	out.swaps = nil
	for _, sw := range vb.swaps {
		if sw.Slot != "ranged" {
			out.swaps = append(out.swaps, sw)
		}
	}
	return out, append(notes, fmt.Sprintf("ranged: final set wears %s (%.1f) over %s (%.1f)", best.Name, best.MeasuredDPS, current.Name, vb.setDPS))
}
