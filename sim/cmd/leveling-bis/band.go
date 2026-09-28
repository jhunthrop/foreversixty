package main

import "sort"

// sourceFor answers loot.json's own question for one item at one
// band: does it have a source usable at this level, and if so which.
//
// Raid sources are excluded below level 60 (this lane's brief: "a
// leveling list is about what a leveling character can get" - the
// leveling-bis design doc's own words). An item with more than one
// source (rare in today's data; loot.json's kinds barely overlap)
// takes the first non-raid one in a fixed kind order, so the choice
// is deterministic across runs rather than dependent on loot.json's
// row order.
//
// TODO(bis-data): "no usable source" is reported as "no known
// source" full stop. The design doc's "zone drop, flagged lucky"
// case needs a zone-drop kind loot.json does not carry yet (see
// data.go's loadLootIndex doc) - once it does, this function should
// return a source with a Lucky marker instead of ok=false for one,
// and the caller (bandPool, below) should keep it out of the BiS
// pick but list it as a "lucky" candidate the way the design doc
// describes, rather than dropping it silently as this prototype does.
var sourceKindPriority = []string{"quest", "dungeon", "crafted", "rep", "pvp", "world", "raid"}

func sourceFor(id, level int, idx lootIndex) (itemSource, bool) {
	srcs := idx[id]
	if len(srcs) == 0 {
		return itemSource{}, false
	}
	byKind := make(map[string]itemSource, len(srcs))
	for _, s := range srcs {
		if _, seen := byKind[s.Kind]; !seen {
			byKind[s.Kind] = s
		}
	}
	for _, kind := range sourceKindPriority {
		s, ok := byKind[kind]
		if !ok {
			continue
		}
		if kind == "raid" && level < 60 {
			continue
		}
		return s, true
	}
	return itemSource{}, false
}

// bandPool is everything candidatesBySlot/pick need for one band and
// faction: every eligible, sourced item, scored - plus the ones that
// were eligible but had no usable source, kept only so the report can
// say how many and which.
type bandPool struct {
	Scored   []scored
	NoSource []candidate
}

// buildBandPool applies eligible(), the source rule, and score(), in
// that order, to the full candidate list for one band and faction. It
// does not pick: candidatesBySlot/pick (pick.go) do that from
// Scored.
func buildBandPool(items []candidate, idx lootIndex, classSlug string, level int, faction string, weights map[string]float64) bandPool {
	var out bandPool
	for _, c := range items {
		if !eligible(c, classSlug, level, faction) {
			continue
		}
		if c.SetID != nil {
			// Excluded, not scored as ineligible: this works around a
			// real engine crash this lane's run hit. At band 60 a
			// picked item carried a set_id whose set-bonus effect is
			// registered for a class other than this character's
			// (wowsims-forever's item_sets_pve.go type-asserts the
			// wearer as a specific class's Agent interface, e.g.
			// rogue.RogueAgent) and the mismatch panics IN A
			// GOROUTINE THIS PROCESS CANNOT RECOVER FROM
			// (core.runSimConcurrent's own goroutine - see the lane
			// report for the exact panic and stack). score() also
			// never models set bonuses (it only sums a single item's
			// own stats), so a set item's score already understates
			// its real value; skipping the whole set-item class here
			// is one exclusion that is honest about both problems at
			// once, not just a crash workaround. See the report's
			// "what I would change" for the real fix: either the
			// engine's cross-class set-bonus registration, or a
			// probe run wrapped so a panic in that specific goroutine
			// cannot take the whole ranking down.
			continue
		}
		src, ok := sourceFor(c.ID, level, idx)
		if !ok {
			out.NoSource = append(out.NoSource, c)
			continue
		}
		if len(c.Slots) == 0 {
			continue
		}
		out.Scored = append(out.Scored, scored{
			candidate: c,
			Score:     score(c, c.Slots[0], weights),
			Source:    src,
			HasSource: true,
		})
	}
	sort.Slice(out.NoSource, func(i, j int) bool { return out.NoSource[i].ID < out.NoSource[j].ID })
	return out
}
