package main

// TestKnownGoodItemsAreSeenAndRanked is lane rank-guardrails' guardrail
// B: TestPublishedBISFilesPassSanityChecks (published_test.go) can only
// confirm what the pipeline DID publish is internally consistent - it
// has no opinion on whether the pipeline missed something a human
// theorycrafter would call an obvious pick. testdata/known-good.json is
// a small, hand-curated list (this lane's own Classic knowledge, every
// id checked against data/builds/1.60.1.70009/items/<class>.json before
// being committed) of items the Classic community agrees are best or
// near-best for six specs at bands 20, 40 and 60. This test re-derives
// eligibility and source the same way checkPublishedBand does (same
// loaders, same build), then asks the harder question: was this
// specific well-known item actually SEEN by the ranker (sourced, not
// invisible) and actually COMPETITIVE (picked, the runner-up, or within
// 10% of the slot's pick score)?
//
// The three-sim-pass gap this test cannot close: rankTrinketSlot,
// rankSlotWithEffects and trySetCompletion (main.go's runSpec) each run
// a real engine sim to re-rank a slot beyond score()'s plain stat total
// - a unit test cannot afford that. This test recomputes
// buildBandPool/candidatesBySlot/pick locally instead, using the REAL
// committed item/loot data and the REAL stat weights the published
// report's own Weights carries (reconstructWeights) - for every slot
// those three passes do not touch (every slot but trinket1/trinket2,
// and only the few main_hand/off_hand/set slots effect-verification or
// set-completion actually changed), this reproduces the exact score and
// pick the nightly published. A known-good item this gap causes to read
// as "not picked" when the real, sim-verified pipeline did pick it
// would be a false failure on this test's part, not a ranker bug - see
// the lane report for which (if any) of tonight's failures fall in that
// category.
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// knownGoodItem is one testdata/known-good.json entry.
type knownGoodItem struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
	Reason   string `json:"reason"`
}

// knownGoodFile is testdata/known-good.json's whole shape: spec slug ->
// band (a JSON string key - encoding/json requires string map keys) ->
// the items curated for that spec at that band. A faction is not part
// of the schema: this test derives which faction(s) to check an item
// against from the item's OWN candidate.FactionRestriction (real data,
// looked up by id), rather than trusting the fixture to redeclare a
// fact the item file already carries and could drift from.
type knownGoodFile map[string]map[string][]knownGoodItem

const knownGoodPath = "testdata/known-good.json"

func TestKnownGoodItemsAreSeenAndRanked(t *testing.T) {
	// Same gate as published_test.go (this lane's brief): the nightly's
	// own sanity-check step (bis.yml, continue-on-error) sets this, an
	// ordinary `go test` run does not.
	if os.Getenv("FOREVER_BIS_SANITY") == "" {
		t.Skip("set FOREVER_BIS_SANITY=1 to check known-good items against the committed bis files (the nightly does)")
	}

	fixture, err := loadKnownGoodFile(knownGoodPath)
	if err != nil {
		t.Fatalf("reading %s: %v", knownGoodPath, err)
	}
	if len(fixture) == 0 {
		t.Fatalf("%s decoded with zero specs - the fixture is empty or mis-shaped", knownGoodPath)
	}

	activeBuild, err := readActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatalf("readActiveBuild: %v", err)
	}
	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", activeBuild)

	for specName, byBand := range fixture {
		specName, byBand := specName, byBand
		t.Run(specName, func(t *testing.T) {
			checkKnownGoodSpec(t, buildDir, specName, byBand)
		})
	}
}

func loadKnownGoodFile(path string) (knownGoodFile, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var f knownGoodFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	return f, nil
}

// checkKnownGoodSpec loads this spec's real item/loot data once (the
// same loaders main.go's own runSpec and published_test.go's
// checkPublishedSpecFile use), reads the nightly's own published
// bis/<spec>.json for its stat weights and picks, then checks every
// known-good item against every band the fixture lists for it.
func checkKnownGoodSpec(t *testing.T, buildDir, specName string, byBand map[string][]knownGoodItem) {
	t.Helper()

	spec, err := loadSpec(publishedRepoRoot, specName)
	if err != nil {
		t.Fatalf("loadSpec(%q): %v", specName, err)
	}

	items, _, err := loadCandidates(buildDir, spec.ClassSlug)
	if err != nil {
		t.Fatalf("loadCandidates(%s, %s): %v", buildDir, spec.ClassSlug, err)
	}
	itemFactionRestriction := make(map[int]string, len(items))
	for _, c := range items {
		if c.FactionRestriction != "" {
			itemFactionRestriction[c.ID] = c.FactionRestriction
		}
	}
	lootIdx, questFloors, err := loadLootIndex(buildDir, itemFactionRestriction)
	if err != nil {
		t.Fatalf("loadLootIndex: %v", err)
	}
	items = applyEffectiveRequiredLevels(items, lootIdx, questFloors)
	byID := make(map[int]candidate, len(items))
	for _, c := range items {
		byID[c.ID] = c
	}

	report, err := readPublishedReport(buildDir, specName)
	if err != nil {
		// Not every written spec has been ranked into a committed
		// bis/*.json yet (published_test.go's own reasoning for
		// skipping a build with no bis directory at all) - report it,
		// do not fail a curation fixture over a spec the nightly has
		// not reached.
		t.Skipf("no published bis/%s.json to check known-good items against: %v", specName, err)
		return
	}

	for bandStr, wantItems := range byBand {
		band, convErr := strconv.Atoi(bandStr)
		if convErr != nil {
			t.Fatalf("known-good.json: %s's band key %q is not a number", specName, bandStr)
		}
		for _, faction := range []string{"alliance", "horde"} {
			bandRow, ok := findBandReport(report, band, faction)
			if !ok {
				t.Errorf("%s: published %s.json has no band %d %s entry (bandRow contract: every written spec covers 20..60 step 10, both factions)", specName, specName, band, faction)
				continue
			}
			weights := reconstructWeights(bandRow.Weights)
			pool := buildBandPool(items, lootIdx, spec.ClassSlug, band, faction, weights, bandRow.ReferenceDPSPerPoint)
			// pickBySlot mirrors main.go's own runSpec exactly (this
			// lane's brief, item 3): without excludeAbovePvpRankCap, this
			// test's own locally-recomputed "primary.Item" reconstructs a
			// DIFFERENT pick than the one actually published whenever a
			// rank-11+ PvP reward outscores everything else in a slot -
			// found dogfooding this exact fix (warrior-arms band 60's own
			// main_hand: this recomputation picked "High Warlord's
			// Greatsword" while the real, published pick is "Arcanite
			// Champion", a Blacksmithing craft) - checkKnownGoodItem's own
			// score() comparison must be measured against what actually
			// ships, not against a picker this test forgot to update.
			bySlot := candidatesBySlot(pool.Scored)
			pickBySlot := make(map[string][]scored, len(bySlot))
			for slot, list := range bySlot {
				filtered := excludeAbovePvpRankCap(list)
				if weaponSlots[slot] {
					filtered = promoteLowValueWeapon(filtered, spec.WeightStats)
				}
				pickBySlot[slot] = filtered
			}
			picks := pick(spec.Spec, pickBySlot)
			for _, want := range wantItems {
				checkKnownGoodItem(t, specName, spec.ClassSlug, band, faction, want, byID, lootIdx, weights, picks, bandRow.ReferenceDPSPerPoint)
			}
		}
	}
}

func readPublishedReport(buildDir, specName string) (specReport, error) {
	path := filepath.Join(buildDir, "bis", specName+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return specReport{}, err
	}
	var report specReport
	if err := json.Unmarshal(b, &report); err != nil {
		return specReport{}, fmt.Errorf("decoding %s: %w", path, err)
	}
	return report, nil
}

func findBandReport(report specReport, band int, faction string) (bandReport, bool) {
	for _, b := range report.Bands {
		if b.Band == band && b.Faction == faction {
			return b, true
		}
	}
	return bandReport{}, false
}

// reconstructWeights rebuilds score()'s own weights map straight from
// the PUBLISHED report's Weights rows, mirroring effectiveWeights'
// (weights.go) own rule without needing that function's
// map[string]api.StatWeight input shape: an insignificant weight
// (report.go's own isWeightSignificant call, already baked into the
// published row's Insignificant flag) contributes zero, exactly as it
// did when main.go's runSpec produced this same band's real picks.
func reconstructWeights(rows []weightRow) map[string]float64 {
	weights := make(map[string]float64, len(rows))
	for _, w := range rows {
		if w.Insignificant {
			continue
		}
		weights[w.Stat] = w.Weight
	}
	return weights
}

// checkKnownGoodItem is this test's whole point: given one known-good
// item and the band+faction's real eligibility/source facts and a
// locally recomputed pick, report eligible/sourced/picked/runner-up/
// score, and fail on the two conditions this lane's brief names -
// eligible but unsourced (invisible to the list, no matter how good
// the item is), or scoring above the pick by more than the 10% margin
// yet neither picked nor the runner-up (a real ranking gap the sim's
// own greedy pick() should not have made).
func checkKnownGoodItem(t *testing.T, specName, classSlug string, band int, faction string, want knownGoodItem, byID map[int]candidate, lootIdx lootIndex, weights map[string]float64, picks map[string]slotPick, referenceDPSPerPoint float64) {
	t.Helper()

	c, ok := byID[want.ItemID]
	if !ok {
		t.Fatalf("%s band %d: known-good item %d (%s) is not in %s's item file at all - this lane's own curation step should have caught this before committing testdata/known-good.json", specName, band, want.ItemID, want.ItemName, specName)
	}
	if c.Name != want.ItemName {
		t.Errorf("%s band %d: known-good.json names item %d as %q, the item file calls it %q - the fixture has drifted from the real data", specName, band, want.ItemID, want.ItemName, c.Name)
	}

	// A faction-restricted item is only ever checked against the one
	// faction that can equip it - eligible() would correctly fail it
	// for the other faction, which is not the interesting finding this
	// test is looking for.
	if c.FactionRestriction != "" && c.FactionRestriction != faction {
		return
	}

	label := fmt.Sprintf("%s band %d %s: %d %s", specName, band, faction, c.ID, c.Name)

	if !eligible(c, classSlug, band, faction) {
		t.Logf("%s: eligible=false (effective required level %d) - %s", label, c.EffectiveRequiredLevel, want.Reason)
		return
	}

	_, sourced := sourceFor(c.ID, band, faction, c.FactionRestriction, lootIdx)
	if !sourced {
		t.Errorf("%s: eligible=true sourced=false - invisible to every published list until a data lane sources it (%s)", label, want.Reason)
		return
	}

	if len(c.Slots) == 0 {
		t.Errorf("%s: eligible=true sourced=true, but carries no planner slot at all (data.go's plannerSlots) - cannot be picked for anything", label)
		return
	}
	itemScore := score(c, c.Slots[0], weights, referenceDPSPerPoint)

	picked, runnerUp := false, false
	for _, slot := range c.Slots {
		sp := picks[slot]
		if sp.Item != nil && sp.Item.ID == c.ID {
			picked = true
		}
		if sp.RunnerUp != nil && sp.RunnerUp.ID == c.ID {
			runnerUp = true
		}
	}

	primary := picks[c.Slots[0]]
	// primary.Item is the slot's actual published pick (recomputed
	// locally against the same weights - see reconstructWeights' own
	// doc). Two DIFFERENT facts get reported here, and only one of
	// them can fail this test (this lane's brief, read literally):
	//
	//   - withinTenPercent: the known-good item's own score is at
	//     least 90% of the pick's - a near-miss worth a reader's
	//     attention, but never a failure by itself. A community-known
	//     item legitimately losing to a better raid-tier item at band
	//     60 (a real, common, CORRECT outcome - see this lane's
	//     report) would otherwise false-fail this test on every such
	//     item, drowning out the findings that actually matter.
	//   - outscoresPick: the known-good item's own score() is HIGHER
	//     than what got published, or nothing was published for the
	//     slot at all - the one case pick()'s own greedy, best-score-
	//     first rule (pick.go) should make impossible unless something
	//     upstream (a pairing/two-hand/cross-class-set exclusion, or a
	//     genuine pick() bug) kept a legitimately better item out of
	//     contention. THIS is "scores above the pick yet was not
	//     picked" - the brief's own fail condition - and it fails
	//     whether or not the item happens to also be within 10%.
	pickScore, pickDesc := 0.0, "(no pick for this slot)"
	outscoresPick := true
	withinTenPercent := false
	if primary.Item != nil {
		pickScore = primary.Item.Score
		pickDesc = fmt.Sprintf("%.2f (%s, id %d)", pickScore, primary.Item.Name, primary.Item.ID)
		outscoresPick = itemScore > pickScore
		// A published pick's score is never negative (report.go's own
		// checkPublishedBand asserts this of the real data); guard the
		// rare recomputed-locally case anyway so *0.9 cannot flip the
		// comparison's direction for a zero-or-negative pick score.
		if pickScore <= 0 {
			withinTenPercent = itemScore >= pickScore
		} else {
			withinTenPercent = itemScore >= pickScore*0.9
		}
	}

	pass := picked || runnerUp || !outscoresPick
	t.Logf("%s: eligible=true sourced=true picked=%v runner_up=%v within_10pct_of_pick=%v score=%.2f vs slot %s's pick %s -- %s", label, picked, runnerUp, withinTenPercent, itemScore, c.Slots[0], pickDesc, want.Reason)
	if !pass {
		t.Errorf("%s: scored %.2f, ABOVE slot %s's published pick %s, yet was neither picked nor the runner-up - a real ranking gap (a pairing/two-hand/cross-class-set exclusion, or a pick() bug), not just a close call", label, itemScore, c.Slots[0], pickDesc)
	}
}
