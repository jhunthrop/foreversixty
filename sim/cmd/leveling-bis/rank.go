package main

// Making an engine-implemented effect count in the ranking (this
// lane's brief, item 3): score() (score.go) is a weighted sum of an
// item's Stats plus its weapon DPS -- it has no notion of a proc at
// all, so a candidate whose real value is an on-hit/on-use effect
// scores at whatever its flat stats alone are worth, exactly the
// "dishonest tie-break" trinkets.go's own doc already called out for
// trinkets specifically. This file generalises that fix to every
// slot: any candidate whose effect_text the engine actually
// implements (effectids_generated.go) is ranked by a real verify sim
// instead, exactly like a trinket; one whose effect the engine does
// NOT implement keeps its score()-based rank and is flagged
// effect_unmodelled in the report (report.go) so the page can say
// "proc not simulated" on it rather than silently pretending the
// stat total is the whole story.

import (
	"sort"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// hasImplementedEffect reports whether c carries an on-hit/on-use/proc
// effect the engine actually implements. A candidate with no
// effect_text at all (most gear) is never "implemented" in this
// sense -- it has nothing for a verify pass to value beyond what
// score() already sees in its Stats.
func hasImplementedEffect(c candidate) bool {
	return c.EffectText != "" && effectImplemented(c.ID)
}

// slotsNeedingEffectVerification returns every slot, other than
// trinket1/trinket2 (already ranked unconditionally by
// rankTrinketSlot regardless of effects -- a trinket has no scorable
// stats to fall back on at all, see trinkets.go's own doc), whose
// candidate pool contains at least one item with an engine-
// implemented effect. A slot with no such candidate is left to
// score() entirely: nothing here could move its ranking, so running
// a verify sim on it would only spend nightly budget for no reason.
func slotsNeedingEffectVerification(bySlot map[string][]scored) []string {
	var out []string
	for slot, list := range bySlot {
		if slot == "trinket1" || slot == "trinket2" {
			continue
		}
		for _, c := range list {
			if hasImplementedEffect(c.candidate) {
				out = append(out, slot)
				break
			}
		}
	}
	sort.Strings(out)
	return out
}

// effectRankTopN bounds how many implemented-effect candidates
// rankSlotWithEffects verifies per slot, beyond the slot's current
// pick -- the same budget reasoning as trinketTopN (trinkets.go).
const effectRankTopN = 4

// rankSlotWithEffects replaces picks[slot] with the engine-verified
// best of: the slot's current score()-based pick, plus up to
// effectRankTopN candidates in its pool that carry an engine-
// implemented effect (best score() first, excluding this slot's own
// pair-mate the same way rankTrinketSlot does). This is the fair
// comparison the brief asks for between "scores best on raw stats"
// and "the engine can actually simulate this one's proc, which
// score() alone cannot see at all". A slot whose own pick already IS
// the only implemented-effect candidate is left unchanged (nothing to
// compare it against). Mirrors rankTrinketSlot's own error handling:
// a candidate whose verify sim fails is skipped and logged, not fatal
// to the slot.
func rankSlotWithEffects(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored, slot string) (map[string]slotPick, []string) {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}

	current := picks[slot]
	if current.Item == nil {
		return out, nil
	}

	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok && picks[mate].Item != nil {
		mateID, mateName = picks[mate].Item.ID, picks[mate].Item.Name
	}

	pool := bySlot[slot]
	// A dual-wielder's main hand is never a two-hander (pick.go's own
	// "main_hand" case, excludeTwoHand): a two-hand candidate's own
	// implemented proc can still out-measure the current one-handed
	// pick in a plain DPS sim below, which for a spec whose kit assumes
	// two weapons (leveling.DualWieldSpecs) would silently swap the
	// character onto a two-hander with no comparison against what it
	// actually forfeits (a whole second weapon's worth of stats/imbue,
	// buildGear's own off_hand-drop rule - the exact bug this excludes:
	// hunter-survival's level-35 main hand picking Frost Tiger Blade
	// [two-hand, a frost-bolt proc] over its dual-wield set, silently
	// carrying an off_hand pick the sim never actually wore).
	if slot == "main_hand" && leveling.DualWieldSpecs[spec.Spec] {
		pool = excludeTwoHand(pool)
	}
	candidates := []scored{*current.Item}
	seen := map[int]bool{current.Item.ID: true}
	for _, c := range pool {
		if len(candidates) >= 1+effectRankTopN {
			break
		}
		if seen[c.ID] || c.ID == mateID || (mateName != "" && c.Name == mateName) {
			continue
		}
		if !hasImplementedEffect(c.candidate) {
			continue
		}
		candidates = append(candidates, c)
		seen[c.ID] = true
	}
	if len(candidates) < 2 {
		return out, nil
	}

	type measured struct {
		item scored
		dps  float64
	}
	var results []measured
	var notes []string
	for _, c := range candidates {
		gear := swapSlot(picks, slot, c.ID, c.TwoHand)
		req := plainRequest(spec, bandCharacter("effect-rank", race, classSlug, spec.Spec, level, talents, gear), trinketRankIterations, verifySeed)
		dps, err := runner.RunPlainDPS(req)
		if err != nil {
			notes = append(notes, formatTrinketRankError(slot, c, err))
			continue
		}
		results = append(results, measured{item: c, dps: dps})
	}
	if len(results) < 2 {
		return out, notes
	}
	sort.SliceStable(results, func(i, j int) bool { return results[i].dps > results[j].dps })

	// This lane's brief, item 1: rankTrinketSlot's own beatsByMargin
	// rule (verify.go's swapMargin, 1%) applies here too, not just to
	// trinkets. Without it, a zero/low-stat proc item (e.g. Shortsword
	// of Vengeance, stats: {}, a world-drop proc-only sword) can win
	// this tournament over the pool's real stat-scored best purely off
	// this short verify sim's own noise floor - the fifth wow-player
	// sweep caught it flipping by faction on an identical candidate
	// pool (warlock-affliction band 60 main_hand: the sword beat Staff
	// of Dar'Orahil, +11 int/+10 hit, on Horde but lost to it on
	// Alliance). current is always the pool's stat-scored best entering
	// this tournament (picks[slot] came from pick()'s score-first
	// choice before this function ran, and every other candidate here
	// was appended "best score() first, excluding this slot's own
	// pair-mate" - this function's own doc above), so it plays the
	// same role a modelled trinket plays in rankTrinketSlot: it keeps
	// the slot unless some other candidate's measured DPS clears
	// beatsByMargin over it.
	winner := 0
	for i, r := range results {
		if r.item.ID == current.Item.ID {
			winner = i
			break
		}
	}
	// winner is now the stat-scored best's own index (0 when it also
	// measured highest, or when it is missing from results entirely -
	// its own sim failed, see the error-skip loop above - in which case
	// the top measured candidate is the only fair fallback). Only when
	// results[0] is some other candidate AND actually clears the noise
	// floor over the stat-scored best does that other candidate get to
	// keep winning.
	if winner > 0 && beatsByMargin(results[0].dps, results[winner].dps) {
		winner = 0
	}
	runnerUpIdx := -1
	for i := range results {
		if i != winner {
			runnerUpIdx = i
			break
		}
	}

	best := results[winner].item
	// This lane's brief, item 7: this tournament ran a real sim (this
	// function's own precondition, len(candidates) >= 2), so the slot's
	// decision is sim-based even when the plain score()-based pick
	// happened to win its own fair fight - MeasuredDPS (scored's own
	// doc) is what tells buildReport (report.go) to publish this row's
	// real DPS instead of score()'s stat estimate.
	best.MeasuredDPS = results[winner].dps
	sp := slotPick{Item: &best}
	if runnerUpIdx >= 0 {
		runnerUp := results[runnerUpIdx].item
		runnerUp.MeasuredDPS = results[runnerUpIdx].dps
		sp.RunnerUp = &runnerUp
	}
	out[slot] = sp
	return out, notes
}
