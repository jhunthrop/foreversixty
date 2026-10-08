package main

// Keeping the incumbent.
//
// A healer's published set must never be worse than the set the site
// already published for the same band, preset and faction, when both are
// measured under the harness the new set was ranked with. The ranker picks
// slot by slot and the guarded score (score_heal.go) is a noisy, nonlinear
// number, so a long chain of locally better picks can still end below a set
// it could have kept. After the last pass the previous published set is
// read back from the committed report, measured under the final harness
// next to the new set, and kept when it beats the new set beyond the two
// runs' combined standard error.

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

// incumbentKey names one published set: a band, a preset, a faction.
type incumbentKey struct {
	Band    int
	Preset  string
	Faction string
}

// incumbentSets is the previous report's gear, slot to item id, per key.
type incumbentSets map[incumbentKey]map[string]int

// loadIncumbentSets reads the previous published report of one spec. A
// missing file is no incumbent (the first run of a spec), not an error.
func loadIncumbentSets(path string) (incumbentSets, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return incumbentSets{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading the previous report %s: %w", path, err)
	}
	var report struct {
		Bands []struct {
			Band    int    `json:"band"`
			Preset  string `json:"preset"`
			Faction string `json:"faction"`
			Slots   []struct {
				Slot   string `json:"slot"`
				ItemID int    `json:"item_id"`
			} `json:"slots"`
		} `json:"bands"`
	}
	if err := json.Unmarshal(raw, &report); err != nil {
		return nil, fmt.Errorf("decoding the previous report %s: %w", path, err)
	}
	out := make(incumbentSets, len(report.Bands))
	for _, b := range report.Bands {
		gear := make(map[string]int, len(b.Slots))
		for _, s := range b.Slots {
			if s.ItemID != 0 {
				gear[s.Slot] = s.ItemID
			}
		}
		out[incumbentKey{Band: b.Band, Preset: b.Preset, Faction: b.Faction}] = gear
	}
	return out, nil
}

// keptIncumbent is the published note of a band that kept its previous set:
// both sets' guarded scores under the final harness, with their errors, and
// the slots where they differ.
type keptIncumbent struct {
	IncumbentScore float64  `json:"incumbent_score"`
	IncumbentError float64  `json:"incumbent_error"`
	NewScore       float64  `json:"new_score"`
	NewError       float64  `json:"new_error"`
	DifferingSlots []string `json:"differing_slots"`
}

// incumbentPicks rebuilds the incumbent as picks on the current candidate
// pools. A slot whose item the current pool does not offer (renamed away,
// no longer eligible) makes the incumbent unreproducible: ok is false and
// the new set stands, as does a band with no incumbent at all. A slot where the incumbent wears what the new set
// wears keeps the new set's own pick; a slot where it differs takes the
// incumbent's item with the new set's item as its runner-up.
func incumbentPicks(gear map[string]int, current map[string]slotPick, pools map[string][]scored) (picks map[string]slotPick, differing []string, ok bool) {
	if len(gear) == 0 {
		return nil, nil, false
	}
	picks = make(map[string]slotPick, len(current))
	for slot, id := range gear {
		if cur := current[slot]; cur.Item != nil && cur.Item.ID == id {
			picks[slot] = cur
			continue
		}
		item, found := findScored(pools[slot], id)
		if !found {
			return nil, nil, false
		}
		picks[slot] = slotPick{Item: &item, RunnerUp: current[slot].Item}
		differing = append(differing, slot)
	}
	for slot, cur := range current {
		if _, worn := gear[slot]; !worn && cur.Item != nil {
			differing = append(differing, slot)
		}
	}
	sort.Strings(differing)
	return picks, differing, true
}

// findScored is the candidate with the given item id.
func findScored(pool []scored, id int) (scored, bool) {
	for _, c := range pool {
		if c.ID == id {
			return c, true
		}
	}
	return scored{}, false
}

// keepIncumbent measures the new band and the incumbent set under the final
// harness and returns the incumbent's band when it beats the new set beyond
// error, the new band otherwise. The note is nil unless the incumbent was
// kept. An incumbent identical to the new set, or one the pools cannot
// reproduce, costs no run.
func keepIncumbent(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, current verifiedBand, gear map[string]int, pools map[string][]scored) (verifiedBand, *keptIncumbent, error) {
	picks, differing, ok := incumbentPicks(gear, current.picks, pools)
	if !ok || len(differing) == 0 {
		return current, nil, nil
	}
	newScore, newErr, err := measureSet(runner, spec, race, classSlug, level, talents, current.picks)
	if err != nil {
		return verifiedBand{}, nil, fmt.Errorf("measuring the new set: %w", err)
	}
	incScore, incErr, err := measureSet(runner, spec, race, classSlug, level, talents, picks)
	if err != nil {
		return verifiedBand{}, nil, fmt.Errorf("measuring the incumbent set: %w", err)
	}
	if incScore-newScore <= math.Hypot(newErr, incErr) {
		return current, nil, nil
	}
	for _, slot := range differing {
		if pk, worn := picks[slot]; worn {
			measured := *pk.Item
			measured.MeasuredDPS = incScore
			pk.Item = &measured
			picks[slot] = pk
		}
	}
	note := &keptIncumbent{IncumbentScore: incScore, IncumbentError: incErr, NewScore: newScore, NewError: newErr, DifferingSlots: differing}
	return verifiedBand{picks: dropStaleSetBonuses(picks), setDPS: incScore, errors: current.errors}, note, nil
}

// measureSet is a set's score and standard error under the verification
// harness: the verification seed and iterations, the band's talents.
func measureSet(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick) (score, stdErr float64, err error) {
	ch := bandCharacter("verify", race, classSlug, spec.Spec, level, talents, buildGear(picks))
	return runner.RunPlainDPSWithError(plainRequest(spec, ch, verifyIterations, verifySeed))
}
