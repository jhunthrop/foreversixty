package main

import "math"

// The trinket inspector (coordinator scope, 2026-11-05): beyond the
// three alternatives, each trinket row publishes the tournament that
// decided it - up to trinketRankingLimit measured candidates with their
// set-DPS gap against the pick, all from one tournament (same harness,
// same iterations, same seed), so the gaps compare with each other.

// trinketRankingLimit bounds slotRow.Ranking.
const trinketRankingLimit = 8

// tournamentEntry is one measured trinket candidate.
type tournamentEntry struct {
	Item scored
	DPS  float64
}

// tournamentOf is a tournament's results, already sorted best first.
func tournamentOf(n int, at func(i int) (scored, float64)) []tournamentEntry {
	out := make([]tournamentEntry, n)
	for i := range out {
		item, dps := at(i)
		out[i] = tournamentEntry{Item: item, DPS: dps}
	}
	return out
}

// rankingRow is one slotRow.Ranking entry.
type rankingRow struct {
	ItemID     int    `json:"item_id"`
	SourceKind string `json:"source_kind,omitempty"`
	// DPSDelta is the candidate's measured set DPS minus the pick's, in
	// the same tournament, to 0.01 DPS: negative when worse. It can be
	// positive: the ranker keeps a modelled trinket over an unmodelled
	// one within the margin, and a trinket whose gain over an empty slot
	// is inside its error gives way (trinkets.go).
	DPSDelta float64 `json:"dps_delta"`
}

// trinketRanking is the published ranking for a trinket slot: nil when
// no tournament decided it, or when a later pass (a swap, a set trial,
// the guard) replaced the tournament's pick, whose gaps would then be
// against an item the tournament never put first.
func trinketRanking(pk slotPick) []rankingRow {
	if pk.Item == nil || len(pk.Tournament) == 0 {
		return nil
	}
	pickDPS, found := 0.0, false
	for _, e := range pk.Tournament {
		if e.Item.ID == pk.Item.ID {
			pickDPS, found = e.DPS, true
			break
		}
	}
	if !found {
		return nil
	}
	out := make([]rankingRow, 0, trinketRankingLimit)
	for _, e := range pk.Tournament {
		if e.Item.ID == pk.Item.ID {
			continue
		}
		if len(out) == trinketRankingLimit {
			break
		}
		d := math.Round((e.DPS-pickDPS)*byKindDeltaPrecision) / byKindDeltaPrecision
		if d == 0 {
			d = 0
		}
		out = append(out, rankingRow{ItemID: e.Item.ID, SourceKind: e.Item.Source.Kind, DPSDelta: d})
	}
	return out
}
