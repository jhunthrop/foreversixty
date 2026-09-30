package main

import (
	"sort"
)

// Ranking trinkets by engine-verified DPS instead of score().
//
// score() is the weighted sum of an item's STAT block; a trinket's real
// value is almost always its proc or on-use effect, which carries no
// stats at all (see score.go's own doc). Every trinket therefore scores
// at or near 0 under the normal pipeline, which the prototype's report
// flagged as dishonest: pick()'s tie-break (item id, ascending) would
// otherwise decide trinket1/trinket2 by which id happens to be lowest,
// not which trinket is actually strongest.
//
// This lane's brief: "rank trinket candidates by the verify step
// (engine, top-N by item level) instead of the stat score." trinketTopN
// candidates per trinket slot, ordered by item level (a real, if rough,
// proxy for "how good a trinket this tier of content dropped" when
// stats cannot rank them), are each equipped in turn alongside the rest
// of this band's picks and run through a real DPS sim; the one that
// measures highest becomes the pick, the next becomes the runner-up -
// so trinket1/trinket2's Verified flag in the final report means the
// same thing it means for every other slot (a real sim, not a stat
// formula, decided it).
//
// Item-level-only selection has its own blind spot (bis-ranker-integrity
// lane, 2026-09-29): a trinket that IS mostly plain stats (Neltharion's
// Tear, spell_power 44 + hit 20 - no proc at all) can sit at a lower
// item level than several proc/on-use trinkets from the same or a later
// raid tier, so it never got a single sim and its own, much higher,
// score() total showed up only as an unverified "alternatives" entry
// the published pick had never actually been measured against (exactly
// the dishonest tie-break this file's own opening paragraph already
// named, just moved from "which id is lowest" to "which axis got
// asked"). trinketShortlist (below) fixes this by adding the top
// trinketTopN candidates BY SCORE to the real-sim pool alongside the
// item-level ones, so the best plain stat-stick and the best proc
// candidates are always tested against EACH OTHER, not just against
// their own axis.
const (
	// trinketTopN bounds how many of a slot's item-level-ranked
	// candidates get their own sim: bounded because this runs once per
	// trinket slot per band per faction per spec, and the nightly
	// budget (this lane's brief) has to fit every written spec.
	trinketTopN = 5
	// trinketRankIterations is fewer than verifyIterations (300): this
	// pass only has to ORDER a handful of candidates relative to each
	// other, not settle a single close call the way verifyBand's final
	// swap check does, so it spends less of the nightly budget per
	// candidate. Same fixed verifySeed as every other sim in this
	// command, so a report is reproducible.
	trinketRankIterations = 100
)

// excludePairMate drops a slot's own pair-mate (by id, or by name - a
// lower/higher-quality version of "the same trinket" - matching
// excludePaired's own rule in pick.go) from list, so neither
// topByItemLevel nor topByScore below can ever ask to equip the same
// physical trinket twice.
func excludePairMate(list []scored, excludeID int, excludeName string) []scored {
	candidates := make([]scored, 0, len(list))
	for _, c := range list {
		if c.ID == excludeID || (excludeName != "" && c.Name == excludeName) {
			continue
		}
		candidates = append(candidates, c)
	}
	return candidates
}

// topByItemLevel returns up to trinketTopN candidates from list, highest
// item level first (ties broken by id, ascending, for a stable order
// across runs). list is assumed already pair-mate-filtered
// (excludePairMate).
func topByItemLevel(list []scored, excludeID int, excludeName string) []scored {
	candidates := excludePairMate(list, excludeID, excludeName)
	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].ItemLevel != candidates[j].ItemLevel {
			return candidates[i].ItemLevel > candidates[j].ItemLevel
		}
		return candidates[i].ID < candidates[j].ID
	})
	if len(candidates) > trinketTopN {
		candidates = candidates[:trinketTopN]
	}
	return candidates
}

// trinketShortlist is rankTrinketSlot's real candidate pool: the union
// of topByItemLevel's top trinketTopN (highest item level - a rough
// proxy for "how good a trinket this tier of content dropped", the
// only ranking this file had before this lane) and the top trinketTopN
// BY SCORE (list is already best-score-first - candidatesBySlot's own
// contract - so this is simply its own head, pair-mate-filtered),
// deduplicated by item id.
//
// This lane's brief (bis-ranker-integrity, 2026-09-29), item 1's own
// mechanism finding: a plain-stat trinket that out-scores everything
// else in the pool (Neltharion's Tear, spell_power 44 + hit 20, against
// mage-fire's Naxxramas picks) can sit at a LOWER item level than five
// higher-ilvl on-use/proc trinkets from the same or a later raid tier,
// so topByItemLevel alone never gave it a single sim - the real
// tournament below crowned whichever proc trinket won among candidates
// that never had to face the best plain stat-stick in the pool at all.
// bySlot[slot]'s own score() ranking cannot crown the WINNER by itself
// either (score() has no notion of a proc - trinkets.go's own package
// doc), but it is exactly the list this file needs to make sure the
// best stat-based candidate gets a fair, real-sim shot alongside the
// best item-level-based ones, so whichever axis actually wins the real
// sim becomes the pick, and the loser becomes an honestly-Verified
// alternative instead of an unverified score() estimate the pick never
// actually faced (report.go's buildAlternatives, this lane's other
// fix).
func trinketShortlist(list []scored, excludeID int, excludeName string) []scored {
	filtered := excludePairMate(list, excludeID, excludeName)
	byItemLevel := topByItemLevel(filtered, 0, "")
	byScore := filtered
	if len(byScore) > trinketTopN {
		byScore = byScore[:trinketTopN]
	}
	seen := make(map[int]bool, len(byItemLevel)+len(byScore))
	out := make([]scored, 0, len(byItemLevel)+len(byScore))
	for _, group := range [][]scored{byItemLevel, byScore} {
		for _, c := range group {
			if seen[c.ID] {
				continue
			}
			seen[c.ID] = true
			out = append(out, c)
		}
	}
	return out
}

// rankTrinketSlot replaces picks[slot] with the engine-verified best of
// its top-item-level candidates (and the runner-up with the second
// best), leaving every other slot's pick untouched. It returns a new
// map rather than mutating picks (this command's own immutability
// rule - see the other pick.go/band.go functions, which all return new
// values instead of editing their input).
//
// A candidate whose own sim fails is skipped and logged, the same
// resilience verify.go's own per-slot swap loop uses (one bad
// candidate should not lose the whole slot's ranking); if every
// candidate fails, picks is returned unchanged (the slot keeps
// whatever score()-based pick it already had - almost always the
// lowest-id eligible trinket, which the report's swap_note on that row
// still leaves honestly unverified) and the caller is told so via the
// returned notes slice, one line per skipped candidate.
func rankTrinketSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored, slot string) (map[string]slotPick, []string) {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}

	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok && picks[mate].Item != nil {
		mateID, mateName = picks[mate].Item.ID, picks[mate].Item.Name
	}
	candidates := trinketShortlist(bySlot[slot], mateID, mateName)
	if len(candidates) == 0 {
		return out, nil
	}

	type measured struct {
		item scored
		dps  float64
	}
	var results []measured
	var notes []string
	for _, c := range candidates {
		gear := swapSlot(picks, slot, c.ID, false)
		req := plainRequest(spec, bandCharacter("trinket-rank", race, classSlug, level, talents, gear), trinketRankIterations, verifySeed)
		dps, err := runner.RunPlainDPS(req)
		if err != nil {
			notes = append(notes, formatTrinketRankError(slot, c, err))
			continue
		}
		results = append(results, measured{item: c, dps: dps})
	}
	if len(results) == 0 {
		return out, notes
	}
	sort.SliceStable(results, func(i, j int) bool { return results[i].dps > results[j].dps })

	best := results[0].item
	// This lane's brief, item 7: best/runnerUp's own MeasuredDPS
	// (scored's own doc) records the real, per-candidate full-set DPS
	// this very tournament measured - a trinket has no scorable stats
	// at all (this file's own package doc), so score()'s Score field
	// stays whatever near-zero value it always was; buildReport reads
	// MeasuredDPS, not Score, to decide what this row publishes.
	best.MeasuredDPS = results[0].dps
	sp := slotPick{Item: &best}
	if len(results) > 1 {
		runnerUp := results[1].item
		runnerUp.MeasuredDPS = results[1].dps
		sp.RunnerUp = &runnerUp
	}
	out[slot] = sp
	return out, notes
}

func formatTrinketRankError(slot string, c scored, err error) string {
	return slot + ": ranking candidate " + c.Name + " failed: " + err.Error()
}
