package main

import (
	"math"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// Source filters on the BiS page (owner MUST, 2026-11-05): filtering is
// not hiding. With a kind off, each slot shows the best pick from the
// kinds still on. A slot row carries, beside its pick, by_kind: per
// filter kind (source_kinds.go), the best-ranked candidate of that kind,
// so the page can re-pick under any combination of kinds without the
// whole pool. resolveFilteredSlots below is the reference resolver;
// web/src/lib/bis/source-filter.ts is its twin.
//
// Rank order is the ranker's own: the pick first (whatever decided it -
// score, a trinket or effect tournament, a set trial, a swap sim), then
// the slot's pool in score order (sortCandidates). byKindRow.Rank is the
// candidate's 1-based position in that pool, so the resolver picks the
// lowest rank among the kinds that are on.
//
// Paired slots. finger1/finger2, trinket1/trinket2 and main_hand/off_hand
// are decided in that order, the second excluding whatever the first
// wears (pairMate: same id, same name, same limit category), exactly as
// pick() does. The first slot of a pair resolves on its own. The second
// slot's by_kind:
//   - keeps every kind, even the pick's own (its pick may be blocked by
//     the first slot's new item);
//   - names in not_with the first slot's possible items (its pick and
//     every by_kind entry) it cannot be worn beside;
//   - carries next, the following candidate of that kind that does not
//     clash with the entry itself, for when the entry is blocked (so only
//     on an entry with a not_with).
// The row's own not_with does the same for the second slot's pick. An
// off hand is never resolved beside a two-handed main hand (two_hand on
// the main hand's row and by_kind entries); a dual-wield spec's band-level
// two-hander-versus-pair decision (twoHandBeatsPair) stays as the ranker
// made it. The one case next cannot cover is a first-slot item that
// clashes with both the entry and its next by different keys (same name
// as one, same limit category as the other); no Forever item does that.

// pairSecondOf maps the second slot of each pair to the first.
var pairSecondOf = map[string]string{
	"finger2":  "finger1",
	"trinket2": "trinket1",
	"off_hand": "main_hand",
}

// byKindDeltaPrecision rounds by_kind DPS gaps to 0.01 DPS: the order
// comes from rank, so finer digits only add bytes.
const byKindDeltaPrecision = 100

// byKindRow is one by_kind entry: the slot's best candidate of one kind.
type byKindRow struct {
	ItemID int `json:"item_id"`
	// Rank is the candidate's 1-based position in the slot's ranked pool.
	Rank int `json:"rank"`
	// SourceKind and Source are the entry's source of THIS kind (an item
	// with several sources names the one the filter lets through).
	// SourceKind is omitted when it is the by_kind key itself, which is
	// every kind but world (world, world_drop or zone).
	SourceKind string `json:"source_kind,omitempty"`
	Source     string `json:"source"`
	// DPSDelta is the gap against the row's pick, in real DPS, negative
	// when worse: alternativeRow.DPSDelta's convention (the score gap
	// times reference_dps_per_point, never positive), or the alternative's
	// own sim-measured delta when the same item is one. Absent on a row
	// with no pick to compare against.
	DPSDelta *float64 `json:"dps_delta,omitempty"`
	// TwoHand marks a two-handed main hand: no off hand beside it.
	TwoHand bool `json:"two_hand,omitempty"`
	// Standing is the reputation standing the source needs ("friendly" ..
	// "exalted"): a rep source's, or an item's own required reputation.
	Standing string `json:"standing,omitempty"`
	// NotWith (second slot of a pair only) is the first slot's possible
	// items this one cannot be worn beside.
	NotWith []int `json:"not_with,omitempty"`
	// Next (second slot of a pair only) is the next candidate of the kind
	// that does not clash with this one.
	Next *byKindRow `json:"next,omitempty"`
}

// sourceFilterInputs is what withSourceKindPicks reads per band+faction.
type sourceFilterInputs struct {
	Spec                 string
	Picks                map[string]slotPick
	PickBySlot           map[string][]scored
	ReferenceDPSPerPoint float64
}

// rankedPool is the pool a slot is decided from, in rank order: pick()'s
// own list, minus the pair-mate exclusion (not_with carries that instead).
func rankedPool(spec, slot string, pickBySlot map[string][]scored) []scored {
	switch slot {
	case "main_hand":
		if leveling.DualWieldSpecs[spec] && !twoHandBeatsPair(pickBySlot) {
			return excludeTwoHand(pickBySlot[slot])
		}
	case "off_hand":
		return offHandCandidates(spec, pickBySlot, nil)
	}
	return pickBySlot[slot]
}

// filterEligible is whether a pool candidate may stand in for the pick:
// a real stat value, as alternatives require, except in a weapon slot,
// where the ranker itself falls back to a low-value weapon over none.
func filterEligible(c scored, slot string) bool {
	return c.Score > 0 || weaponSlots[slot]
}

// withSourceKindPicks fills each row's by_kind and pick kind fields.
// rows are buildReport's, in slotOrder; the result is a new slice.
func withSourceKindPicks(rows []slotRow, in sourceFilterInputs) []slotRow {
	out := make([]slotRow, len(rows))
	choices := make(map[string][]scored, len(pairSecondOf))
	for i, row := range rows {
		pk := in.Picks[row.Slot]
		list := rankedPool(in.Spec, row.Slot, in.PickBySlot)
		firstSlot, isSecond := pairSecondOf[row.Slot]
		var mateChoices []scored
		if isSecond {
			mateChoices = choices[firstSlot]
		}
		filled, slotChoices := sourceKindRow(row, pk, list, isSecond, mateChoices, in.ReferenceDPSPerPoint)
		choices[row.Slot] = slotChoices
		out[i] = filled
	}
	return out
}

// sourceKindRow is one row with its source-filter fields, and every item
// the row can resolve to under some filter (the pick, then each entry).
func sourceKindRow(row slotRow, pk slotPick, list []scored, isSecond bool, mateChoices []scored, refDPSPerPoint float64) (slotRow, []scored) {
	var pickKinds map[string]itemSource
	pickID := 0
	var possible []scored
	if pk.Item != nil {
		pickID = pk.Item.ID
		pickKinds = filterKindSources(*pk.Item)
		if kinds := orderedFilterKinds(pickKinds); len(kinds) > 1 {
			row.Kinds = kinds
		}
		row.TwoHand = row.Slot == "main_hand" && pk.Item.TwoHand
		row.Standing = pk.Item.Source.Standing
		row.KindStandings = kindStandings(pickKinds)
		if isSecond {
			row.NotWith = clashingIDs(pk.Item.candidate, mateChoices)
		}
		possible = append(possible, *pk.Item)
	}
	delta := byKindDelta(pk, row.Alternatives, refDPSPerPoint)
	byKind := make(map[string]byKindRow)
	for _, kind := range filterKindOrder {
		if _, own := pickKinds[kind]; own && !isSecond {
			continue
		}
		at := nextOfKind(list, 0, kind, pickID, row.Slot, pairMate{})
		if at < 0 {
			continue
		}
		entry := newByKindRow(list[at], at, kind, row.Slot, delta)
		if isSecond {
			entry.NotWith = clashingIDs(list[at].candidate, mateChoices)
		}
		if len(entry.NotWith) > 0 {
			if n := nextOfKind(list, at+1, kind, pickID, row.Slot, mateOf(&list[at])); n >= 0 {
				next := newByKindRow(list[n], n, kind, row.Slot, delta)
				next.NotWith = clashingIDs(list[n].candidate, mateChoices)
				entry.Next = &next
			}
		}
		byKind[kind] = entry
		possible = append(possible, list[at])
	}
	if len(byKind) > 0 {
		row.ByKind = byKind
	}
	return row, possible
}

// nextOfKind is the index of the first candidate at or after from that
// belongs to kind, is not the pick, may stand in for it and is not ruled
// out by clash; -1 when none.
func nextOfKind(list []scored, from int, kind string, pickID int, slot string, clash pairMate) int {
	for i := from; i < len(list); i++ {
		c := list[i]
		if c.ID == pickID || !filterEligible(c, slot) || clash.excludes(c.candidate) {
			continue
		}
		if _, ok := filterKindSources(c)[kind]; ok {
			return i
		}
	}
	return -1
}

func newByKindRow(c scored, index int, kind, slot string, delta func(scored) *float64) byKindRow {
	src := filterKindSources(c)[kind]
	return byKindRow{
		ItemID:     c.ID,
		Rank:       index + 1,
		SourceKind: sourceKindUnlessKey(src.Kind, kind),
		Source:     src.Label,
		DPSDelta:   delta(c),
		TwoHand:    slot == "main_hand" && c.TwoHand,
		Standing:   src.Standing,
	}
}

// sourceKindUnlessKey is the entry's source_kind, empty when the by_kind
// key already says it (entrySourceKind restores it).
func sourceKindUnlessKey(sourceKind, key string) string {
	if sourceKind == key {
		return ""
	}
	return sourceKind
}

// entrySourceKind is an entry's source kind: source_kind, else the key.
func entrySourceKind(e byKindRow, key string) string {
	if e.SourceKind != "" {
		return e.SourceKind
	}
	return key
}

// byKindDelta is the DPS gap function for one row: the alternative's own
// published delta when the candidate is one, else the score gap in DPS,
// never positive (buildAlternatives' convention); nil without a pick.
func byKindDelta(pk slotPick, alternatives []alternativeRow, refDPSPerPoint float64) func(scored) *float64 {
	altDelta := make(map[int]float64, len(alternatives))
	for _, a := range alternatives {
		altDelta[a.ItemID] = a.DPSDelta
	}
	return func(c scored) *float64 {
		if pk.Item == nil {
			return nil
		}
		d, ok := altDelta[c.ID]
		if !ok {
			d = math.Min(0, (c.Score-pk.Item.Score)*refDPSPerPoint)
		}
		r := math.Round(d*byKindDeltaPrecision) / byKindDeltaPrecision
		if r == 0 {
			r = 0 // no "-0"
		}
		return &r
	}
}

// clashingIDs is the ids among choices c may not be worn beside.
func clashingIDs(c candidate, choices []scored) []int {
	var out []int
	for i := range choices {
		if mateOf(&choices[i]).excludes(c) && !containsInt(out, choices[i].ID) {
			out = append(out, choices[i].ID)
		}
	}
	return out
}

func containsInt(list []int, v int) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

// filteredSlot is one slot's answer under a source filter.
type filteredSlot struct {
	ItemID  int
	TwoHand bool
	// Source/SourceKind are what the row shows; empty for the pick itself.
	Source     string
	SourceKind string
}

// kindStandings is the standing each filter kind's source needs, for the
// kinds a standing gates; nil when none is.
func kindStandings(sources map[string]itemSource) map[string]string {
	var out map[string]string
	for kind, s := range sources {
		if s.Standing == "" {
			continue
		}
		if out == nil {
			out = make(map[string]string, len(sources))
		}
		out[kind] = s.Standing
	}
	return out
}

// pickFilterKinds is a row's pick's filter kinds: kinds, else its one
// published source_kind's.
func pickFilterKinds(row slotRow) []string {
	if len(row.Kinds) > 0 {
		return row.Kinds
	}
	if k, ok := filterKindOfSource(row.SourceKind); ok {
		return []string{k}
	}
	return nil
}

// resolveFilteredSlots is the reference resolver: each slot's item with
// only the allowed kinds on, nil for "No <slot> from these sources".
func resolveFilteredSlots(rows []slotRow, allowed map[string]bool) map[string]*filteredSlot {
	out := make(map[string]*filteredSlot, len(rows))
	for _, row := range rows {
		var mate *filteredSlot
		firstSlot, isSecond := pairSecondOf[row.Slot]
		if isSecond {
			mate = out[firstSlot]
		}
		out[row.Slot] = resolveFilteredRow(row, allowed, isSecond, mate)
	}
	return out
}

func resolveFilteredRow(row slotRow, allowed map[string]bool, isSecond bool, mate *filteredSlot) *filteredSlot {
	if row.Slot == "off_hand" && mate != nil && mate.TwoHand {
		return nil
	}
	mateID := 0
	if mate != nil {
		mateID = mate.ItemID
	}
	if row.ItemID != 0 && anyAllowed(pickFilterKinds(row), allowed) && !containsInt(row.NotWith, mateID) {
		return &filteredSlot{ItemID: row.ItemID, TwoHand: row.TwoHand}
	}
	var best *byKindRow
	bestKind := ""
	for _, kind := range filterKindOrder {
		entry, ok := row.ByKind[kind]
		if !ok || !allowed[kind] {
			continue
		}
		e := &entry
		if isSecond && containsInt(e.NotWith, mateID) {
			e = e.Next
		}
		if e == nil || (isSecond && containsInt(e.NotWith, mateID)) {
			continue
		}
		if best == nil || e.Rank < best.Rank {
			best, bestKind = e, kind
		}
	}
	if best == nil {
		return nil
	}
	return &filteredSlot{ItemID: best.ItemID, TwoHand: best.TwoHand, Source: best.Source, SourceKind: entrySourceKind(*best, bestKind)}
}

func anyAllowed(kinds []string, allowed map[string]bool) bool {
	for _, k := range kinds {
		if allowed[k] {
			return true
		}
	}
	return false
}
