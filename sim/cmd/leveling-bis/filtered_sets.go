package main

import (
	"fmt"
	"math"
	"slices"
)

// Filtered-set DPS (coordinator scope, 2026-11-05): the page labels a
// filtered set with its own DPS where the ranker measured one. Each band
// re-picks its set under a few fixed filter sets (repickFiltered, the
// same rule the page resolves by_kind with) and measures every distinct
// set once the way set_dps is measured (measureFinalSet), so a filtered
// figure and set_dps compare like with like; "all" is set_dps itself. Any other combination
// of kinds the page builds has no measured DPS: the page shows the
// unfiltered figure, labelled as such.

// filterSet is one published filter combination.
type filterSet struct {
	// Key is the bandReport.FilteredSetDPS key.
	Key   string
	Kinds []string
	// MaxStanding drops a candidate whose source of every allowed kind
	// needs a higher reputation standing; "" for no limit.
	MaxStanding string
	// LevelingOnly sets are measured below level 60 only.
	LevelingOnly bool
	// SixtyOnly sets are measured at level 60 only.
	SixtyOnly bool
}

// publishedFilterSets are the filter keys a file can carry:
//   - all: every kind on (the unfiltered set, under the same harness).
//   - dungeon: dungeon drops only.
//   - leveling_default: the page's leveling default, every kind on but no
//     reputation item above Friendly; bands below 60 only. The default's
//     other half, no craft above the character's likely skill, is not
//     applied: neither loot.json nor the item data carries a recipe's
//     skill, so crafts stay on in this figure.
//   - pre_raid: every kind but raid and PvP (high-rank PvP gear is not
//     pre-raid); band 60 only. The BiS page's Pre-raid toggle
//     (web/src/lib/bis/pre-raid.ts) is the same kinds.
var publishedFilterSets = []filterSet{
	{Key: "all", Kinds: filterKindOrder},
	{Key: "dungeon", Kinds: []string{filterKindDungeon}},
	{Key: "leveling_default", Kinds: filterKindOrder, MaxStanding: "friendly", LevelingOnly: true},
	{Key: "pre_raid", Kinds: kindsWithout(filterKindRaid, filterKindPvP), SixtyOnly: true},
}

// kindsWithout is filterKindOrder less the given kinds.
func kindsWithout(drop ...string) []string {
	out := make([]string, 0, len(filterKindOrder))
	for _, kind := range filterKindOrder {
		if !slices.Contains(drop, kind) {
			out = append(out, kind)
		}
	}
	return out
}

// filteredSetDPS is one filter set's measured set.
type filteredSetDPS struct {
	DPS   float64 `json:"dps"`
	Error float64 `json:"error"`
	// ChangedSlots is how many slots differ from the published set.
	ChangedSlots int `json:"changed_slots"`
}

// allows reports whether c may stand in under the filter set.
func (f filterSet) allows(c scored) bool {
	sources := filterKindSources(c)
	for _, kind := range f.Kinds {
		s, ok := sources[kind]
		if ok && (f.MaxStanding == "" || standingRank[s.Standing] <= standingRank[f.MaxStanding]) {
			return true
		}
	}
	return false
}

// repickFiltered is each slot's item with only candidates keep allows:
// per slot in slotOrder, the first of [pick, ranked pool] that may stand
// in, is allowed and does not clash with its pair slot's filtered item;
// no off hand beside a two-hander. resolveFilteredSlots reproduces it
// from the published by_kind fields for any kind-only filter.
func repickFiltered(spec string, picks map[string]slotPick, pickBySlot map[string][]scored, keep func(scored) bool) map[string]slotPick {
	out := make(map[string]slotPick, len(slotOrder))
	for _, slot := range slotOrder {
		var mate *scored
		if first, ok := pairSecondOf[slot]; ok {
			mate = out[first].Item
		}
		if slot == "off_hand" && mate != nil && mate.TwoHand {
			out[slot] = slotPick{}
			continue
		}
		pk := picks[slot]
		ordered := make([]scored, 0, len(pickBySlot[slot])+1)
		if pk.Item != nil {
			ordered = append(ordered, *pk.Item)
		}
		for _, c := range rankedPool(spec, slot, pickBySlot) {
			if (pk.Item == nil || c.ID != pk.Item.ID) && filterEligible(c, slot) {
				ordered = append(ordered, c)
			}
		}
		out[slot] = slotPick{}
		for _, c := range ordered {
			if (mate != nil && mateOf(mate).excludes(c.candidate)) || !keep(c) {
				continue
			}
			c := c
			out[slot] = slotPick{Item: &c}
			break
		}
	}
	return out
}

// changedSlots counts the slots whose item differs between two sets.
func changedSlots(a, b map[string]slotPick) int {
	n := 0
	for _, slot := range slotOrder {
		if itemIDOf(a[slot]) != itemIDOf(b[slot]) {
			n++
		}
	}
	return n
}

func itemIDOf(pk slotPick) int {
	if pk.Item == nil {
		return 0
	}
	return pk.Item.ID
}

// filteredSetInputs is one band+faction+preset's set and pools.
type filteredSetInputs struct {
	Spec       specInfo
	Race       string
	Band       int
	Talents    string
	Picks      map[string]slotPick
	PickBySlot map[string][]scored
	// SetDPS and SetDPSError are the published set's own measurement.
	SetDPS      float64
	SetDPSError float64
}

// measureFilteredSets measures each published filter set's re-picked set
// once per distinct gear (gearFingerprint); a set equal to the published
// one reuses set_dps. A failed sim leaves its key
// out and is returned as a note for the log.
func measureFilteredSets(runner engineRunner, in filteredSetInputs) (map[string]filteredSetDPS, []string) {
	out := make(map[string]filteredSetDPS, len(publishedFilterSets))
	byGear := map[string]filteredSetDPS{
		gearFingerprint(in.Picks): {DPS: roundTenth(in.SetDPS), Error: roundTenth(in.SetDPSError)},
	}
	var notes []string
	for _, f := range publishedFilterSets {
		if (f.LevelingOnly && in.Band >= 60) || (f.SixtyOnly && in.Band < 60) {
			continue
		}
		set := repickFiltered(in.Spec.Spec, in.Picks, in.PickBySlot, f.allows)
		key := gearFingerprint(set)
		got, seen := byGear[key]
		if !seen {
			dps, stdErr, err := measureFinalSet(runner, in.Spec, in.Race, in.Band, in.Talents, set)
			if err != nil {
				notes = append(notes, fmt.Sprintf("filtered set %s: %v", f.Key, err))
				continue
			}
			got = filteredSetDPS{DPS: roundTenth(dps), Error: roundTenth(stdErr)}
			byGear[key] = got
		}
		got.ChangedSlots = changedSlots(in.Picks, set)
		out[f.Key] = got
	}
	if len(out) == 0 {
		return nil, notes
	}
	return out, notes
}

func roundTenth(x float64) float64 { return math.Round(x*10) / 10 }
