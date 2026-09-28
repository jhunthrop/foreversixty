package main

import (
	"sort"

	"github.com/jhunthrop/foreversixty/sim/api"
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

// topByItemLevel returns up to trinketTopN candidates from list, highest
// item level first (ties broken by id, ascending, for a stable order
// across runs), excluding any candidate that is already this slot's
// pair-mate (by id or name - a lower/higher-quality version of "the
// same trinket" - matching excludePaired's own rule in pick.go) so a
// trial here can never ask to equip the same physical trinket twice.
func topByItemLevel(list []scored, excludeID int, excludeName string) []scored {
	candidates := make([]scored, 0, len(list))
	for _, c := range list {
		if c.ID == excludeID || (excludeName != "" && c.Name == excludeName) {
			continue
		}
		candidates = append(candidates, c)
	}
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
func rankTrinketSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, picks map[string]slotPick, bySlot map[string][]scored, slot string) (map[string]slotPick, []string) {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}

	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok && picks[mate].Item != nil {
		mateID, mateName = picks[mate].Item.ID, picks[mate].Item.Name
	}
	candidates := topByItemLevel(bySlot[slot], mateID, mateName)
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
		req := plainRequest(spec, api.CharacterSpec{Name: "trinket-rank", Race: race, Class: classSlug, Level: level, Gear: gear}, trinketRankIterations, verifySeed)
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
	sp := slotPick{Item: &best}
	if len(results) > 1 {
		runnerUp := results[1].item
		sp.RunnerUp = &runnerUp
	}
	out[slot] = sp
	return out, notes
}

func formatTrinketRankError(slot string, c scored, err error) string {
	return slot + ": ranking candidate " + c.Name + " failed: " + err.Error()
}
