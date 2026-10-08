package main

// Set-aware completion. score() sums each item's own stats independently
// (score.go) and has no notion that pieces worn together unlock a third
// thing neither carries alone -- a set bonus is invisible to it exactly
// the same way a proc is (rank.go's own doc). The per-slot pass therefore
// can only wear a set piece when that piece is the top-scored candidate of
// its own slot, so a bonus worth more than the stat gap between a set
// piece and the best loose piece never enters a pick.
//
// trySetCompletion closes that gap with a bounded search. For every set
// the engine implements, it builds "wear the set's best piece in each of
// its slots, up to the next bonus threshold, keep the per-slot picks
// elsewhere", runs it through the same verification harness the per-slot
// picks use (verifyIterations, verifySeed) and adopts it only when it
// beats the picks it started from by swapMargin and beyond the combined
// sim error. A short screening run (trinketRankIterations) comes first and
// only a trial that clears half the margin earns the full-length run, which
// keeps the cost near a third of running every trial at full length.

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// setBonusTier is one bonus of a set: what it grants at that many worn pieces.
type setBonusTier struct {
	Pieces      int    `json:"pieces"`
	Description string `json:"description"`
}

// setInfo is a set's display name and its bonus tiers, ascending by pieces.
type setInfo struct {
	Name    string         `json:"name"`
	Bonuses []setBonusTier `json:"bonuses"`
}

// setCatalog is a build's sets.json keyed by set id.
type setCatalog map[int]setInfo

// setBonusNote is the published reason a piece is worn: the set bonus it
// completes. slotRow.SetBonus carries it, omitted for every other row.
type setBonusNote struct {
	Set    string `json:"set"`
	Pieces int    `json:"pieces"`
	Bonus  string `json:"bonus"`
}

// loadSetCatalog reads <buildDir>/sets.json.
func loadSetCatalog(buildDir string) (setCatalog, error) {
	raw, err := os.ReadFile(filepath.Join(buildDir, "sets.json"))
	if err != nil {
		return nil, fmt.Errorf("reading sets.json: %w", err)
	}
	var rows []struct {
		ID int `json:"id"`
		setInfo
	}
	if err := json.Unmarshal(raw, &rows); err != nil {
		return nil, fmt.Errorf("parsing sets.json: %w", err)
	}
	out := make(setCatalog, len(rows))
	for _, r := range rows {
		info := r.setInfo
		sort.Slice(info.Bonuses, func(i, j int) bool { return info.Bonuses[i].Pieces < info.Bonuses[j].Pieces })
		out[r.ID] = info
	}
	return out, nil
}

// bonusAt is the tier granted at exactly pieces worn pieces.
func (s setInfo) bonusAt(pieces int) (setBonusTier, bool) {
	for _, b := range s.Bonuses {
		if b.Pieces == pieces {
			return b, true
		}
	}
	return setBonusTier{}, false
}

// bestSetPiecePerSlot is, for each non-trinket slot, the best-scored
// candidate belonging to setID (bySlot lists are score-ordered). Trinkets
// are ranked by rankTrinketSlot on a different axis (item level), so they
// never take part. Two defences mirror pick.go/rank.go:
//
//   - A dual-wielder's main hand never offers a two-hander.
//   - finger1/finger2 share one candidate list, so a ring already offered
//     for one slot is skipped in the next: one physical ring is one piece.
func bestSetPiecePerSlot(bySlot map[string][]scored, specSlug string, setID int) map[string]scored {
	out := map[string]scored{}
	used := map[int]bool{}
	for _, slot := range slotOrder {
		if slot == "trinket1" || slot == "trinket2" {
			continue
		}
		list := bySlot[slot]
		if slot == "main_hand" && leveling.DualWieldSpecs[specSlug] {
			list = excludeTwoHand(list)
		}
		for _, cand := range list {
			if cand.SetID == nil || *cand.SetID != setID || used[cand.ID] {
				continue
			}
			used[cand.ID] = true
			out[slot] = cand
			break
		}
	}
	return out
}

// wornSetPieces counts the picks that belong to setID.
func wornSetPieces(picks map[string]slotPick, setID int) int {
	n := 0
	for _, pk := range picks {
		if pk.Item != nil && pk.Item.SetID != nil && *pk.Item.SetID == setID {
			n++
		}
	}
	return n
}

// setFillSlot is one slot a set trial would change and what it costs in score.
type setFillSlot struct {
	slot string
	item scored
	loss float64
}

// setFillSlots lists the slots where wearing the set's piece replaces a
// pick that is not already a piece of that set, cheapest score loss first
// (slot order breaks ties). A slot with no pick is left alone, and so is a
// piece the pair-mate slot already wears.
func setFillSlots(picks map[string]slotPick, pieces map[string]scored, setID int) []setFillSlot {
	var out []setFillSlot
	for _, slot := range slotOrder {
		piece, ok := pieces[slot]
		current := picks[slot].Item
		if !ok || current == nil || current.ID == piece.ID {
			continue
		}
		if current.SetID != nil && *current.SetID == setID {
			continue
		}
		if mate, paired := pairSlot[slot]; paired && picks[mate].Item != nil && picks[mate].Item.ID == piece.ID {
			continue
		}
		out = append(out, setFillSlot{slot: slot, item: piece, loss: current.Score - piece.Score})
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].loss < out[j].loss })
	return out
}

// setTrial is the candidate gear for one threshold of one set.
type setTrial struct {
	picks   map[string]slotPick
	changed []string
}

// buildSetTrial wears the cheapest need pieces of fill on top of picks.
func buildSetTrial(picks map[string]slotPick, fill []setFillSlot, need int) setTrial {
	trial := setTrial{picks: clonePicksMap(picks)}
	for _, f := range fill[:need] {
		item := f.item
		// The per-slot pick this piece replaces stays the slot's runner-up, so
		// verifyBand still measures it against the set piece: a later swap that
		// beats the set (by the sim, with the rest of the set worn) wins.
		trial.picks[f.slot] = slotPick{Item: &item, RunnerUp: picks[f.slot].Item}
		trial.changed = append(trial.changed, f.slot)
	}
	return trial
}

// setThresholds is every bonus threshold above worn, ascending.
func setThresholds(info setInfo, worn int) []int {
	var out []int
	for _, b := range info.Bonuses {
		if b.Pieces > worn {
			out = append(out, b.Pieces)
		}
	}
	return out
}

// beatsBeyondError is whether dps beats baseline by swapMargin and by more
// than the two runs' combined standard error.
func beatsBeyondError(dps, dpsErr, baseline, baselineErr float64) bool {
	return beatsByMargin(dps, baseline) && dps-baseline > math.Hypot(dpsErr, baselineErr)
}

// setScreenMargin is how far a screening run must beat its baseline to earn a
// full-length confirmation: half of swapMargin, so a trial the short run
// already finds losing or tied never costs the 300-iteration run.
const setScreenMargin = swapMargin / 2

// setReading is one sim run's mean and standard error.
type setReading struct{ dps, stdErr float64 }

// setCompletion carries one band's trial context through the set loop.
type setCompletion struct {
	runner    engineRunner
	spec      specInfo
	race      string
	classSlug string
	level     int
	talents   string
	picks     map[string]slotPick
	catalog   setCatalog
	bySlot    map[string][]scored
	// screenBase and confirmBase are the working picks measured at screening
	// and verification length, nil until a trial first needs them and reset
	// whenever an adoption changes the picks.
	screenBase  *setReading
	confirmBase *setReading
	extraRuns   int
	notes       []string
}

func (c *setCompletion) measure(label string, picks map[string]slotPick, iterations int) (setReading, error) {
	req := plainRequest(c.spec, bandCharacter(label, c.race, c.classSlug, c.spec.Spec, c.level, c.talents, buildGear(picks)), iterations, verifySeed)
	c.extraRuns++
	dps, stdErr, err := c.runner.RunPlainDPSWithError(req)
	return setReading{dps: dps, stdErr: stdErr}, err
}

// baseline measures the working picks once per length and caches the reading.
func (c *setCompletion) baseline(cached **setReading, iterations int) (setReading, error) {
	if *cached == nil {
		reading, err := c.measure("set-completion-baseline", c.picks, iterations)
		if err != nil {
			return setReading{}, err
		}
		*cached = &reading
	}
	return **cached, nil
}

// screens is whether a short run of trial leaves it worth confirming.
func (c *setCompletion) screens(trial setTrial) (bool, error) {
	base, err := c.baseline(&c.screenBase, trinketRankIterations)
	if err != nil {
		return false, err
	}
	reading, err := c.measure("set-completion-screen", trial.picks, trinketRankIterations)
	if err != nil {
		return false, err
	}
	return reading.dps > base.dps*(1+setScreenMargin), nil
}

// confirms is whether a verification-length run of trial beats the working
// picks by swapMargin and beyond the combined sim error, with its reading.
func (c *setCompletion) confirms(trial setTrial) (bool, setReading, error) {
	base, err := c.baseline(&c.confirmBase, verifyIterations)
	if err != nil {
		return false, setReading{}, err
	}
	reading, err := c.measure("set-completion", trial.picks, verifyIterations)
	if err != nil {
		return false, setReading{}, err
	}
	return beatsBeyondError(reading.dps, reading.stdErr, base.dps, base.stdErr), reading, nil
}

// trySet walks one set's reachable thresholds, adopting each trial that wins.
func (c *setCompletion) trySet(setID int) {
	info := c.catalog[setID]
	pieces := bestSetPiecePerSlot(c.bySlot, c.spec.Spec, setID)
	for _, threshold := range setThresholds(info, wornSetPieces(c.picks, setID)) {
		fill := setFillSlots(c.picks, pieces, setID)
		need := threshold - wornSetPieces(c.picks, setID)
		if need <= 0 || need > len(fill) {
			continue
		}
		trial := buildSetTrial(c.picks, fill, need)
		won, reading, err := c.runTrial(trial)
		if err != nil {
			c.notes = append(c.notes, fmt.Sprintf("set %d (%s) %d-piece trial: verify failed: %v", setID, info.Name, threshold, err))
			continue
		}
		if won {
			c.notes = append(c.notes, fmt.Sprintf("set %d (%s) %d-piece beat the per-slot picks: %.1f vs %.1f - adopted", setID, info.Name, threshold, reading.dps, c.confirmBase.dps))
			c.adopt(info, threshold, trial, reading)
		}
	}
}

// runTrial screens trial, then confirms it when the screen passes.
func (c *setCompletion) runTrial(trial setTrial) (bool, setReading, error) {
	ok, err := c.screens(trial)
	if err != nil || !ok {
		return false, setReading{}, err
	}
	return c.confirms(trial)
}

// adopt makes trial the working picks and tags its changed slots with the
// bonus they were chosen for.
func (c *setCompletion) adopt(info setInfo, threshold int, trial setTrial, won setReading) {
	bonus, _ := info.bonusAt(threshold)
	note := setBonusNote{Set: info.Name, Pieces: threshold, Bonus: bonus.Description}
	for _, slot := range trial.changed {
		pk := trial.picks[slot]
		// Every adopted piece is sim-decided: its MeasuredDPS is the winning
		// trial's own full-set figure (scored's own doc), so buildReport
		// publishes it instead of score()'s stat estimate.
		pk.Item.MeasuredDPS = won.dps
		pk.SetBonus = &note
		trial.picks[slot] = pk
	}
	c.picks = trial.picks
	c.screenBase = nil
	c.confirmBase = &won
}

// candidateSetIDs is every implemented set with at least two distinct slots
// offering one of its pieces, ascending.
func candidateSetIDs(bySlot map[string][]scored, specSlug string, catalog setCatalog) []int {
	seen := map[int]bool{}
	for _, list := range bySlot {
		for _, cand := range list {
			if cand.SetID != nil && setEffectImplemented(*cand.SetID) {
				seen[*cand.SetID] = true
			}
		}
	}
	var ids []int
	for id := range seen {
		if _, known := catalog[id]; known && len(bestSetPiecePerSlot(bySlot, specSlug, id)) >= 2 {
			ids = append(ids, id)
		}
	}
	sort.Ints(ids)
	return ids
}

// dropStaleSetBonuses clears a set-bonus note whose bonus a later adoption or
// swap broke by replacing one of that set's pieces.
func dropStaleSetBonuses(picks map[string]slotPick) map[string]slotPick {
	out := clonePicksMap(picks)
	for slot, pk := range picks {
		if pk.SetBonus == nil || pk.Item == nil || pk.Item.SetID == nil {
			continue
		}
		if wornSetPieces(picks, *pk.Item.SetID) < pk.SetBonus.Pieces {
			pk.SetBonus = nil
			out[slot] = pk
		}
	}
	return out
}

// trySetCompletion tries every implemented set with two or more pieces in
// the band's pool against the per-slot picks, one reachable bonus threshold
// at a time, and returns the possibly-updated picks plus log notes. Sets
// are tried in id order, each against the picks the earlier ones left; the
// last note reports the extra sim runs spent. Verify errors are logged, not
// fatal, like every other single-candidate failure in this command.
func trySetCompletion(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored, catalog setCatalog) (map[string]slotPick, []string) {
	c := &setCompletion{
		runner: runner, spec: spec, race: race, classSlug: classSlug, level: level, talents: talents,
		picks: picks, catalog: catalog, bySlot: bySlot,
	}
	started := time.Now()
	for _, id := range candidateSetIDs(bySlot, spec.Spec, catalog) {
		c.trySet(id)
	}
	if c.extraRuns == 0 {
		return picks, nil
	}
	c.notes = append(c.notes, fmt.Sprintf("set completion: %d extra sim runs in %.1fs", c.extraRuns, time.Since(started).Seconds()))
	return dropStaleSetBonuses(c.picks), c.notes
}

// verifiedBand is a band's picks after the verification pass: the set DPS
// the sim measured for them, the runner-ups it tested (swaps) and the slots
// it could not test (errors).
type verifiedBand struct {
	picks  map[string]slotPick
	setDPS float64
	swaps  []swapResult
	errors []string
}

// verifyAndSwap measures picks, tests each slot's runner-up and promotes
// every one the sim measured ahead of the pick: a runner-up the sim measured
// ahead of the scored pick IS the pick, so the published row, the set DPS
// and the next band's diff all name the item a player should wear.
func verifyAndSwap(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick) (verifiedBand, error) {
	setDPS, swaps, verifyErrors, err := verifyBand(runner, spec, race, classSlug, level, talents, picks)
	if err != nil {
		return verifiedBand{}, fmt.Errorf("verify run (baseline): %w", err)
	}
	picks, setDPS, swaps, err = applySwaps(runner, spec, race, classSlug, level, talents, picks, swaps, setDPS)
	if err != nil {
		return verifiedBand{}, fmt.Errorf("verify run (after swaps): %w", err)
	}
	return verifiedBand{picks: picks, setDPS: setDPS, swaps: swaps, errors: verifyErrors}, nil
}

// hasSetBonus reports whether any pick carries a set-bonus note.
func hasSetBonus(picks map[string]slotPick) bool {
	for _, pk := range picks {
		if pk.SetBonus != nil {
			return true
		}
	}
	return false
}

// completeSets runs trySetCompletion against a verified band and, when it
// adopts anything, verifies and swaps the completed picks again (an adopted
// slot keeps its replaced pick as runner-up, so the sim can still undo it).
// The completed band is kept only when its set DPS beats the band it started
// from, so set completion can never publish a lower set than the per-slot
// picks would have.
func completeSets(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, before verifiedBand, bySlot map[string][]scored, catalog setCatalog) (verifiedBand, []string, error) {
	completed, notes := trySetCompletion(runner, spec, race, classSlug, level, talents, before.picks, bySlot, catalog)
	if !hasSetBonus(completed) {
		return before, notes, nil
	}
	after, err := verifyAndSwap(runner, spec, race, classSlug, level, talents, enforceTwoHandOffHandInvariant(completed))
	if err != nil {
		return verifiedBand{}, nil, err
	}
	if after.setDPS <= before.setDPS {
		notes = append(notes, fmt.Sprintf("set completion: the completed set verified at %.1f, not above the per-slot set's %.1f - reverted", after.setDPS, before.setDPS))
		return before, notes, nil
	}
	after.picks = dropStaleSetBonuses(after.picks)
	return after, notes, nil
}
