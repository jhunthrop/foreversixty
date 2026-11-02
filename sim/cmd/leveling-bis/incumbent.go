package main

// Keeping the verify-best set.
//
// Owner ruling 4 (round 5): a published set must never be worse than a set
// the ranker considered for the same band, preset and faction, when both
// are measured under one harness. The ranker picks slot by slot on noisy
// short sims, so a chain of locally better picks can end below the set it
// started from, or below the set the site already published. After the
// last pass every such set (the previous published set read back from the
// committed report, and each stage the band passed through) is measured
// under the guard harness next to the final set, and the best of them is
// kept when it beats the final set beyond the two runs' combined standard
// error. Every role runs the guard; until round 5 only healers compared
// against the published set, and DPS and tank bands published sets their
// own check scored lower (Arcane, Enhancement, Destruction, the tanks).

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
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

// consideredSet is one set a band passed through on its way to the set it
// publishes, or the set the site published before.
type consideredSet struct {
	Label string
	Picks map[string]slotPick
}

// The labels keepVerifyBest publishes for the set it kept.
const (
	consideredIncumbent       = "the previous published set"
	consideredBeforeVerify    = "the set before the verify swaps"
	consideredBeforeSets      = "the set before set completion"
	consideredBeforeConfirmed = "the set before the confirmed-stats check"
)

// keptAlternative is the published note of a band whose final set lost to a
// set it considered: which set was kept, both scores under the guard harness
// with their errors, and the slots where they differ.
type keptAlternative struct {
	Kept           string   `json:"kept"`
	KeptScore      float64  `json:"kept_score"`
	KeptError      float64  `json:"kept_error"`
	NewScore       float64  `json:"new_score"`
	NewError       float64  `json:"new_error"`
	DifferingSlots []string `json:"differing_slots"`
}

// incumbentPicks rebuilds the incumbent as picks on the current candidate
// pools. A slot whose item the current pool does not offer (renamed away,
// no longer eligible) makes the incumbent unreproducible: ok is false and
// the new set stands, as does a band with no incumbent at all. A slot where
// the incumbent wears what the new set wears keeps the new set's own pick.
func incumbentPicks(gear map[string]int, current map[string]slotPick, pools map[string][]scored) (picks map[string]slotPick, ok bool) {
	if len(gear) == 0 {
		return nil, false
	}
	picks = make(map[string]slotPick, len(current))
	for slot, id := range gear {
		if cur := current[slot]; cur.Item != nil && cur.Item.ID == id {
			picks[slot] = cur
			continue
		}
		item, found := findScored(pools[slot], id)
		if !found {
			return nil, false
		}
		picks[slot] = slotPick{Item: &item}
	}
	return picks, true
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

// differingSlots is every slot where a and b wear different items (or one
// wears nothing), sorted.
func differingSlots(a, b map[string]slotPick) []string {
	seen := map[string]bool{}
	var out []string
	for _, m := range []map[string]slotPick{a, b} {
		for slot := range m {
			if seen[slot] {
				continue
			}
			seen[slot] = true
			if pickedID(a[slot]) != pickedID(b[slot]) {
				out = append(out, slot)
			}
		}
	}
	sort.Strings(out)
	return out
}

func pickedID(pk slotPick) int {
	if pk.Item == nil {
		return 0
	}
	return pk.Item.ID
}

// wearsUnconfirmedIn reports whether picks wears an item the client has not
// confirmed in any of slots: preferConfirmedStats holds such an item out on
// purpose, so the guard never restores one.
func wearsUnconfirmedIn(picks map[string]slotPick, slots []string) bool {
	for _, slot := range slots {
		if pk := picks[slot]; pk.Item != nil && pk.Item.ClientUnconfirmed {
			return true
		}
	}
	return false
}

// guardIterations is the iterations a guard run asks for. Every role is
// compared at the healer harness's effective count: the heal runner scales
// verifyIterations by healIterationScale itself, so a healer asks for
// verifyIterations and every other role for the scaled count. At 300
// iterations a set's error is about 1% of its DPS, too wide to resolve the
// 1 to 2% losses round 5 found.
func guardIterations(spec specInfo) int {
	if isHealer(spec) {
		return verifyIterations
	}
	return verifyIterations * healIterationScale
}

// keepVerifyBest is owner ruling 4's guard: the ranker never publishes a set
// its own check scores lower than an alternative it considered. The final
// band and every distinct considered set are measured under the guard
// harness; the best considered set is published instead when it beats the
// final one beyond the two runs' combined standard error. A set identical
// to the final one, a repeat of a set already measured, or one wearing an
// item the client has not confirmed costs no run. The note is nil unless
// another set was kept.
func keepVerifyBest(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, current verifiedBand, considered []consideredSet) (verifiedBand, *keptAlternative, error) {
	type contender struct {
		set       consideredSet
		differing []string
	}
	var contenders []contender
	measuredKeys := map[string]bool{}
	for _, c := range considered {
		differing := differingSlots(c.Picks, current.picks)
		if len(differing) == 0 || wearsUnconfirmedIn(c.Picks, differing) {
			continue
		}
		key := gearFingerprint(c.Picks)
		if measuredKeys[key] {
			continue
		}
		measuredKeys[key] = true
		contenders = append(contenders, contender{set: c, differing: differing})
	}
	if len(contenders) == 0 {
		return current, nil, nil
	}
	newScore, newErr, err := measureSet(runner, spec, race, classSlug, level, talents, current.picks)
	if err != nil {
		return verifiedBand{}, nil, fmt.Errorf("measuring the final set: %w", err)
	}
	var best *contender
	var bestScore, bestErr float64
	for i := range contenders {
		score, stdErr, err := measureSet(runner, spec, race, classSlug, level, talents, contenders[i].set.Picks)
		if err != nil {
			return verifiedBand{}, nil, fmt.Errorf("measuring %s: %w", contenders[i].set.Label, err)
		}
		if best == nil || score > bestScore {
			best, bestScore, bestErr = &contenders[i], score, stdErr
		}
	}
	if bestScore-newScore <= math.Hypot(newErr, bestErr) {
		return current, nil, nil
	}
	picks := copyPicks(best.set.Picks)
	for _, slot := range best.differing {
		pk, worn := picks[slot]
		if !worn || pk.Item == nil {
			continue
		}
		measured := *pk.Item
		measured.MeasuredDPS = bestScore
		pk.Item = &measured
		if replaced := current.picks[slot].Item; replaced != nil {
			pk.RunnerUp = replaced
		}
		picks[slot] = pk
	}
	note := &keptAlternative{Kept: best.set.Label, KeptScore: bestScore, KeptError: bestErr, NewScore: newScore, NewError: newErr, DifferingSlots: best.differing}
	return verifiedBand{picks: dropStaleSetBonuses(picks), setDPS: bestScore, errors: current.errors}, note, nil
}

// gearFingerprint is a set's equipped gear as a comparable string.
func gearFingerprint(picks map[string]slotPick) string {
	var b strings.Builder
	for _, g := range buildGear(picks) {
		fmt.Fprintf(&b, "%s=%d;", g.Slot, g.ItemID)
	}
	return b.String()
}

// measureSet is a set's score and standard error under the guard harness:
// the verification seed, guardIterations, the band's talents.
func measureSet(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick) (score, stdErr float64, err error) {
	ch := bandCharacter("verify", race, classSlug, spec.Spec, level, talents, buildGear(picks))
	return runner.RunPlainDPSWithError(plainRequest(spec, ch, guardIterations(spec), verifySeed))
}
