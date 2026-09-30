package main

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// trinketTestStat and trinketTestWeights (bis-ranker-integrity-8,
// 2026-09-30): trinket() sets this stat on every synthetic trinket it
// builds, and every rankTrinketSlot/trinketShortlist test below passes
// trinketTestWeights as the weights map, so these plain synthetic
// trinkets clear trinketShortlist's own qualifying gate
// (hasPositivelyWeightedStat) exactly the way a real trinket's own
// stats would - without changing what any of these tests actually
// assert (the DPS tournament, swap-margin and immutability behaviour
// downstream of trinketShortlist, not the gate itself, which has its
// own dedicated tests below).
const trinketTestStat = "test_stat"

var trinketTestWeights = map[string]float64{trinketTestStat: 1}

func trinket(id int, name string, itemLevel int) scored {
	return scored{candidate: candidate{ID: id, Name: name, ItemLevel: itemLevel, Slots: []string{"trinket1", "trinket2"}, Stats: map[string]float64{trinketTestStat: 1}}}
}

// trinketWithEffect is trinket() plus an EffectText, for the hybrid
// sweep's own trinket-margin tests below. id must be a real engine-
// implemented-effect id (effectids_generated.go) for a "modelled"
// trinket, or any id NOT in that map (e.g. one obviously fake, like
// 999999) for an "unmodelled" one - trinketEffectUnmodelled (trinkets.go)
// reads exactly that distinction.
func trinketWithEffect(id int, name string, itemLevel int, effectText string) scored {
	s := trinket(id, name, itemLevel)
	s.EffectText = effectText
	return s
}

func TestTopByItemLevelOrdersHighestFirstAndBoundsToTopN(t *testing.T) {
	list := []scored{
		trinket(1, "A", 10),
		trinket(2, "B", 30),
		trinket(3, "C", 20),
		trinket(4, "D", 30), // ties item level 30 with B, breaks on id ascending
		trinket(5, "E", 5),
		trinket(6, "F", 40),
	}
	got := topByItemLevel(list, 0, "")
	if len(got) != trinketTopN {
		t.Fatalf("topByItemLevel returned %d, want trinketTopN (%d)", len(got), trinketTopN)
	}
	wantOrder := []int{6, 2, 4, 3, 1} // 40, 30(id2), 30(id4), 20, 10 - item 5 (ilvl 5) is bumped
	for i, id := range wantOrder {
		if got[i].ID != id {
			t.Fatalf("topByItemLevel[%d].ID = %d, want %d (full: %+v)", i, got[i].ID, id, got)
		}
	}
}

// trinketScored is trinket() plus an explicit Score, for the score-axis
// half of trinketShortlist's own tests.
func trinketScored(id int, name string, itemLevel int, sc float64) scored {
	s := trinket(id, name, itemLevel)
	s.Score = sc
	return s
}

// trinketWithStat is a bare candidate (no dummy trinketTestStat, unlike
// trinket()) carrying exactly one named stat amount - trinketShortlist's
// own gating tests below need to control precisely which stat a
// candidate carries, since trinket()'s own dummy stat would otherwise
// trivially qualify every candidate under trinketTestWeights.
func trinketWithStat(id int, name string, itemLevel int, stat string, amount float64) scored {
	return scored{candidate: candidate{ID: id, Name: name, ItemLevel: itemLevel, Slots: []string{"trinket1", "trinket2"}, Stats: map[string]float64{stat: amount}}}
}

func TestHasPositivelyWeightedStat(t *testing.T) {
	weights := map[string]float64{"hit": 1, "crit": 0, "stamina": -1}
	cases := []struct {
		name string
		c    candidate
		want bool
	}{
		{"a positive amount of a positively-weighted stat", candidate{Stats: map[string]float64{"hit": 9}}, true},
		{"a positive amount of a zero-weighted stat", candidate{Stats: map[string]float64{"crit": 14}}, false},
		{"a positive amount of a negatively-weighted stat", candidate{Stats: map[string]float64{"stamina": 20}}, false},
		{"a stat this spec's weights do not mention at all", candidate{Stats: map[string]float64{"spirit": 5}}, false},
		{"no stats at all", candidate{}, false},
		{"a zero amount of a positively-weighted stat", candidate{Stats: map[string]float64{"hit": 0}}, false},
	}
	for _, tc := range cases {
		if got := hasPositivelyWeightedStat(tc.c, weights); got != tc.want {
			t.Errorf("%s: hasPositivelyWeightedStat(%+v) = %v, want %v", tc.name, tc.c, got, tc.want)
		}
	}
}

func TestHasUseEffect(t *testing.T) {
	cases := []struct {
		name string
		text string
		want bool
	}{
		{"a real Use: effect", "Use: +150 Attack Power, +2% Hit.", true},
		{"a passive Equip/proc effect", "Equip: Restores health over time.", false},
		{"no effect at all", "", false},
	}
	for _, tc := range cases {
		if got := hasUseEffect(candidate{EffectText: tc.text}); got != tc.want {
			t.Errorf("%s: hasUseEffect(%q) = %v, want %v", tc.name, tc.text, got, tc.want)
		}
	}
}

// This is the regression fix's own repro (bis-ranker-integrity-8,
// 2026-09-30): hunter-marksmanship Alliance band 60's Frozen Heart of
// the Mountain shape - a real, positively-weighted hit-rating trinket
// whose score() total (already run through the rating-to-percent
// conversion, score.go) sits far below several higher-item-level,
// higher-score trinkets in the same pool. Neither the old top-5-by-
// item-level nor top-5-by-score bucket would ever have reached it;
// hasPositivelyWeightedStat must include it directly, regardless of
// either axis.
func TestTrinketShortlistRatingTrinketWithLowPostConversionScoreStillQualifies(t *testing.T) {
	weights := map[string]float64{"hit": 1}
	list := []scored{
		// Five higher-item-level, higher-score candidates that carry no
		// weighted stat at all (e.g. pure stamina/spirit trinkets) -
		// exactly the shape that used to fill both top-5 buckets first.
		trinketWithStat(2, "ItemLevel 90, unweighted stats", 90, "stamina", 40),
		trinketWithStat(3, "ItemLevel 85, unweighted stats", 85, "stamina", 35),
		trinketWithStat(4, "ItemLevel 83, unweighted stats", 83, "stamina", 30),
		trinketWithStat(5, "ItemLevel 80, unweighted stats", 80, "stamina", 25),
		trinketWithStat(6, "ItemLevel 78, unweighted stats", 78, "stamina", 20),
		// Frozen Heart of the Mountain's own shape: low item level, a
		// small rating-derived hit percentage (post-conversion), but a
		// genuinely positive weighted stat.
		trinketWithStat(249469, "Frozen Heart of the Mountain", 55, "hit", 0.9),
	}
	got := trinketShortlist(list, 0, "", weights)
	found := false
	for _, c := range got {
		if c.ID == 249469 {
			found = true
		}
	}
	if !found {
		t.Fatalf("trinketShortlist = %+v, want Frozen Heart of the Mountain (249469) included on its own weighted hit stat, off both the item-level and score axes", got)
	}
}

// The mirror case: a candidate whose only stats are NOT positively
// weighted by this spec, with no engine-implemented effect and no Use
// effect either, is genuinely worth nothing here - excluded, no matter
// how high its item level or how many stat points it carries.
func TestTrinketShortlistExcludesACandidateOffEveryQualifyingAxis(t *testing.T) {
	weights := map[string]float64{"hit": 1}
	list := []scored{
		trinketWithStat(2, "Frozen Heart of the Mountain", 55, "hit", 0.9),
		trinketWithStat(999, "High ItemLevel, Unweighted Stats", 90, "stamina", 999),
	}
	got := trinketShortlist(list, 0, "", weights)
	for _, c := range got {
		if c.ID == 999 {
			t.Fatalf("trinketShortlist = %+v, want item 999 excluded: no weighted stat, no implemented effect, no Use effect", got)
		}
	}
}

// No arbitrary top-5 cap applies to the qualifying set as long as it
// stays at or under trinketShortlistBound - this lane's brief: the OLD
// top-5-by-score cap, applied before any relevance check, is exactly
// what dropped Frozen Heart of the Mountain in the first place.
func TestTrinketShortlistIncludesEveryQualifyingCandidateBelowTheBound(t *testing.T) {
	weights := map[string]float64{"hit": 1}
	var list []scored
	for i := 1; i <= 10; i++ {
		list = append(list, trinketWithStat(i, "Qualifying", 50+i, "hit", float64(i)))
	}
	got := trinketShortlist(list, 0, "", weights)
	if len(got) != 10 {
		t.Fatalf("trinketShortlist returned %d candidates, want all 10 qualifying candidates (well under trinketShortlistBound, no cap should apply)", len(got))
	}
}

// This lane's brief (bis-ranker-integrity-6), item 2: Hand of Justice's
// own repro shape - a real, engine-implemented-effect trinket
// (effectids_generated.go) with an empty stat block and an item level
// below five OTHER real trinkets in the pool still had no way to reach
// a single sim before this lane, even though the engine can measure
// its real effect. trinketShortlist must include it anyway - weights
// is empty here so none of the plain candidates qualify by stat,
// isolating the effect-based qualification this test is about.
func TestTrinketShortlistIncludesAnImplementedEffectTrinketOffBothAxes(t *testing.T) {
	list := []scored{
		trinketScored(2, "ItemLevel 90, Best Score", 90, 100),
		trinketScored(3, "ItemLevel 85", 85, 90),
		trinketScored(4, "ItemLevel 83", 83, 80),
		trinketScored(5, "ItemLevel 80", 80, 70),
		trinketScored(6, "ItemLevel 78", 78, 60),
		// Hand of Justice's own shape: lowest item level in the pool,
		// no weighted stats, but its effect IS implemented.
		trinketWithEffect(modelledEffectItemID, "Hand of Justice", 58, "1% chance on Melee hit to gain 1 extra attack."),
	}
	got := trinketShortlist(list, 0, "", map[string]float64{})
	found := false
	for _, c := range got {
		if c.ID == modelledEffectItemID {
			found = true
		}
	}
	if !found {
		t.Fatalf("trinketShortlist = %+v, want the implemented-effect trinket (id %d) included despite losing both the item-level and score axes", got, modelledEffectItemID)
	}
}

// The mirror case: a trinket with NO implemented effect (an ordinary
// candidate, or one whose real proc the engine does not simulate -
// Blackhand's Breadth's own crit-chance proc, not in
// effectids_generated.go) and no weighted stat or Use effect gets no
// force-include.
func TestTrinketShortlistDoesNotForceIncludeAnUnimplementedEffectTrinket(t *testing.T) {
	list := []scored{
		trinketScored(2, "ItemLevel 90, Best Score", 90, 100),
		trinketScored(3, "ItemLevel 85", 85, 90),
		trinketScored(4, "ItemLevel 83", 83, 80),
		trinketScored(5, "ItemLevel 80", 80, 70),
		trinketScored(6, "ItemLevel 78", 78, 60),
		// Blackhand's Breadth's own shape: lowest item level, no
		// weighted stats, a real effect_text the engine does not
		// implement and which is not a Use: effect either (999999 is
		// not a real effectids_generated.go id).
		trinketWithEffect(999999, "Blackhand's Breadth", 63, "Improves your chance to get a critical strike with melee attacks by 2%."),
	}
	got := trinketShortlist(list, 0, "", map[string]float64{})
	for _, c := range got {
		if c.ID == 999999 {
			t.Fatalf("trinketShortlist = %+v, want the unimplemented-effect trinket excluded (off every qualifying axis)", got)
		}
	}
}

// An on-USE effect the engine does not implement still qualifies -
// this lane's brief's third condition, independent of hasImplementedEffect.
func TestTrinketShortlistIncludesAnUnimplementedUseEffectTrinket(t *testing.T) {
	list := []scored{
		trinketScored(2, "ItemLevel 90, Best Score", 90, 100),
		trinketScored(3, "ItemLevel 85", 85, 90),
		trinketScored(4, "ItemLevel 83", 83, 80),
		trinketScored(5, "ItemLevel 80", 80, 70),
		trinketScored(6, "ItemLevel 78", 78, 60),
		// A real on-use trinket (client tooltip prefix "Use:") whose own
		// effect the engine does not implement (999999 is not a real
		// effectids_generated.go id) - still a click a player can make.
		trinketWithEffect(999999, "Some On-Use Trinket", 58, "Use: +150 Attack Power for 15 sec."),
	}
	got := trinketShortlist(list, 0, "", map[string]float64{})
	found := false
	for _, c := range got {
		if c.ID == 999999 {
			found = true
		}
	}
	if !found {
		t.Fatalf("trinketShortlist = %+v, want the unimplemented Use: effect trinket included", got)
	}
}

// trinketQualifyingScored is trinketWithStat's own "hit" shape (always
// qualifies via hasPositivelyWeightedStat against weights {"hit": 1})
// plus an explicit Score, so a test list can be built already
// best-score-first (candidatesBySlot's own contract, which
// trinketShortlist's own score-trim relies on).
func trinketQualifyingScored(id int, name string, itemLevel int, sc float64) scored {
	s := trinketWithStat(id, name, itemLevel, "hit", 1)
	s.Score = sc
	return s
}

// Once the qualifying set exceeds trinketShortlistBound, this function
// trims by score down toward the bound, but still keeps the qualifying
// set's own top-by-item-level candidates even if their score fell
// outside that cut.
func TestTrinketShortlistTrimsToTheBoundByScoreButKeepsTopItemLevel(t *testing.T) {
	weights := map[string]float64{"hit": 1}
	// 30 qualifying candidates (over trinketShortlistBound, 24), already
	// best-score-first (list's own contract): score and item level both
	// descend together, id 100 highest.
	var list []scored
	for i := 0; i < 30; i++ {
		list = append(list, trinketQualifyingScored(100+i, "Qualifying", 130-i, float64(30-i)))
	}
	// One more candidate, appended last (lowest score in the pool, so it
	// sits outside the top trinketShortlistBound by score) but the
	// single HIGHEST item level of the whole pool - it must survive the
	// trim via topByItemLevel.
	list = append(list, trinketQualifyingScored(999, "Low Score, Highest ItemLevel", 500, -1))

	got := trinketShortlist(list, 0, "", weights)
	foundHighItemLevel := false
	for _, c := range got {
		if c.ID == 999 {
			foundHighItemLevel = true
		}
	}
	if !foundHighItemLevel {
		t.Fatalf("trinketShortlist trimmed away item 999 (the pool's own highest item level), want it kept via topByItemLevel")
	}
	// The worst-score, worst-item-level of the 30 "ordinary" candidates
	// (id 129: score 1, item level 101) is outside both the top-24-by-
	// score cut and topByItemLevel's own top trinketTopN (5) - trimmed.
	for _, c := range got {
		if c.ID == 129 {
			t.Fatalf("trinketShortlist = %+v, want the pool's own worst-score, worst-item-level ordinary candidate trimmed", got)
		}
	}
}

func TestTrinketShortlistExcludesPairMateByIDAndName(t *testing.T) {
	list := []scored{trinket(1, "Same Name", 10), trinket(2, "Same Name", 20), trinket(3, "Other", 5)}
	got := trinketShortlist(list, 1, "", trinketTestWeights)
	for _, c := range got {
		if c.ID == 1 {
			t.Fatalf("trinketShortlist still carries excluded id 1: %+v", got)
		}
	}
	got2 := trinketShortlist(list, 0, "Same Name", trinketTestWeights)
	for _, c := range got2 {
		if c.Name == "Same Name" {
			t.Fatalf("trinketShortlist still carries an item sharing the excluded name: %+v", got2)
		}
	}
}

func TestTopByItemLevelExcludesPairMateByIDAndName(t *testing.T) {
	list := []scored{trinket(1, "Same Name", 10), trinket(2, "Same Name", 20), trinket(3, "Other", 5)}
	got := topByItemLevel(list, 1, "")
	for _, c := range got {
		if c.ID == 1 {
			t.Fatalf("topByItemLevel still carries excluded id 1: %+v", got)
		}
	}
	got2 := topByItemLevel(list, 0, "Same Name")
	for _, c := range got2 {
		if c.Name == "Same Name" {
			t.Fatalf("topByItemLevel still carries an item sharing the excluded name: %+v", got2)
		}
	}
}

func TestFormatTrinketRankError(t *testing.T) {
	testErr := errors.New("engine exploded")
	msg := formatTrinketRankError("trinket1", trinket(9, "Bad Trinket", 10), testErr)
	if !strings.Contains(msg, "trinket1") || !strings.Contains(msg, "Bad Trinket") || !strings.Contains(msg, testErr.Error()) {
		t.Errorf("formatTrinketRankError = %q, missing an expected substring", msg)
	}
}

func TestRankTrinketSlotPicksTheHighestMeasuredDPS(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder"}}},
	}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Low DPS Trinket", 30), trinket(3, "High DPS Trinket", 20)},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}): 100,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 3}}): 200,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 3 {
		t.Fatalf("trinket1 pick = %+v, want item 3 (higher measured dps)", out["trinket1"].Item)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != 2 {
		t.Fatalf("trinket1 runner-up = %+v, want item 2", out["trinket1"].RunnerUp)
	}
	// picks is untouched (immutability rule): a fresh map came back.
	if picks["trinket1"].Item.ID != 1 {
		t.Fatalf("the input picks map was mutated: %+v", picks["trinket1"])
	}
	// This lane's brief, item 7: the winner and runner-up each carry
	// their OWN measured DPS from this tournament, not each other's and
	// not zero - report.go's buildReport reads this to publish sim_dps
	// instead of score()'s stat estimate.
	if out["trinket1"].Item.MeasuredDPS != 200 {
		t.Errorf("trinket1 pick MeasuredDPS = %v, want 200", out["trinket1"].Item.MeasuredDPS)
	}
	if out["trinket1"].RunnerUp.MeasuredDPS != 100 {
		t.Errorf("trinket1 runner-up MeasuredDPS = %v, want 100", out["trinket1"].RunnerUp.MeasuredDPS)
	}
}

func TestRankTrinketSlotSkipsAFailingCandidateAndKeepsGoing(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Fails To Sim", 30), trinket(3, "Sims Fine", 20)},
	}
	fake := &fakeEngine{
		FailGear:  gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}),
		DPSByGear: map[string]float64{gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 3}}): 50},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1 (item 2 failed)", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 3 {
		t.Fatalf("trinket1 pick = %+v, want item 3 (the only one that sims)", out["trinket1"].Item)
	}
	if out["trinket1"].RunnerUp != nil {
		t.Fatalf("trinket1 runner-up = %+v, want nil: only one candidate sim succeeded", out["trinket1"].RunnerUp)
	}
}

func TestRankTrinketSlotAllCandidatesFailLeavesPicksUnchanged(t *testing.T) {
	original := &scored{candidate: candidate{ID: 1, Name: "Original Pick"}}
	picks := map[string]slotPick{"trinket1": {Item: original}}
	bySlot := map[string][]scored{"trinket1": {trinket(2, "Fails", 30)}}
	fake := &fakeEngine{FailGear: gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}})}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1", notes)
	}
	if out["trinket1"].Item != original {
		t.Fatalf("trinket1 = %+v, want the original pick unchanged", out["trinket1"].Item)
	}
}

func TestRankTrinketSlotNoCandidatesReturnsUnchanged(t *testing.T) {
	original := &scored{candidate: candidate{ID: 1, Name: "Original Pick"}}
	picks := map[string]slotPick{"trinket1": {Item: original}}
	fake := &fakeEngine{}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, map[string][]scored{}, "trinket1", trinketTestWeights)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["trinket1"].Item != original {
		t.Fatalf("trinket1 = %+v, want unchanged", out["trinket1"].Item)
	}
}

// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 2:
// rankTrinketSlot's own no-trinket baseline sim (swapSlot's itemID-0
// shape) is what lets report.go's zero-value gate tell a genuinely
// worthless trinket apart from a real one - MeasuredDPS alone (the
// whole SET's own absolute DPS) can never do this, since it is always
// positive regardless of whether the trinket itself contributes
// anything at all.
func TestRankTrinketSlotComputesGainAgainstANoTrinketBaseline(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder"}}},
	}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Low DPS Trinket", 30), trinket(3, "High DPS Trinket", 20)},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}): 100,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 3}}): 200,
			// The baseline call swaps the slot's own item id to 0, which
			// swapSlot's own doc says drops the slot from the gear list
			// entirely - an empty gear key.
			gearKey(nil): 190,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if !out["trinket1"].Item.GainMeasured {
		t.Fatalf("trinket1 pick GainMeasured = false, want true")
	}
	if out["trinket1"].Item.MeasuredGainDPS != 10 {
		t.Errorf("trinket1 pick MeasuredGainDPS = %v, want 10 (200 measured - 190 baseline)", out["trinket1"].Item.MeasuredGainDPS)
	}
	if !out["trinket1"].RunnerUp.GainMeasured || out["trinket1"].RunnerUp.MeasuredGainDPS != -90 {
		t.Errorf("trinket1 runner-up gain = measured=%v dps=%v, want measured=true dps=-90 (100 - 190)", out["trinket1"].RunnerUp.GainMeasured, out["trinket1"].RunnerUp.MeasuredGainDPS)
	}
}

// A baseline sim failure costs the gain check, not the ranking itself -
// tenet 8: never claim a gain this command could not actually measure.
func TestRankTrinketSlotLeavesGainUnmeasuredWhenTheBaselineSimFails(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Only Candidate", 30)},
	}
	fake := &fakeEngine{
		// The no-trinket baseline call's own gear list is empty
		// (swapSlot's itemID-0 shape) - fakeEngine's FailGear sentinel
		// cannot target an empty fingerprint (it treats "" as "unset"),
		// so DPSFunc singles out the baseline call by its empty gear.
		DPSFunc: func(req api.SimRequest) (float64, error) {
			if len(req.Character.Gear) == 0 {
				return 0, errors.New("fakeEngine: forced baseline failure")
			}
			return 100, nil
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 1 {
		t.Fatalf("notes = %v, want exactly 1 (the baseline failure)", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 2 {
		t.Fatalf("trinket1 pick = %+v, want item 2 (the ranking itself is unaffected)", out["trinket1"].Item)
	}
	if out["trinket1"].Item.GainMeasured {
		t.Errorf("trinket1 pick GainMeasured = true, want false: the baseline sim failed")
	}
}

// modelledEffectItemID is a real engine-implemented-effect id
// (effectids_generated.go) that is ALSO known to the real embedded
// simdb (rank_test.go's own 754, "the real Shortsword of Vengeance
// id" - the same item rank_test.go already uses for exactly this
// reason). trinketEffectUnmodelled now asks effectVerifiedInSim
// (rank.go), which requires both: an id in effectids_generated.go's
// map but absent from simdb (Hand of Justice 11815's own defect,
// this lane's brief) reads as unmodelled precisely because a real
// sim of it never actually wears the item - picking a ONLY-in-map id
// for this constant (the previous 647) would fail that second check
// and read as unmodelled here too, which is not what these tests are
// about.
const modelledEffectItemID = 754

// Hybrid sweep, bis-ranker-integrity-4 lane, item 4: Serenity Field (an
// unmodelled Spirit self-buff) beat a real combat trinket at band 60
// ret/enhancement by a margin smaller than the tournament's own noise -
// a modelled trinket must keep the slot over a currently-leading
// unmodelled one unless the unmodelled one clears beatsByMargin's own
// 1% bar (verify.go).
func TestRankTrinketSlotModelledTrinketKeepsSlotWhenUnmodelledWinIsWithinMargin(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(modelledEffectItemID, "Real Combat Trinket", 60, "On use: deals damage"),
			trinketWithEffect(999999, "Serenity Field", 60, "Equip: restores spirit over time"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: modelledEffectItemID}}): 200,
			// 201 is inside beatsByMargin's 1% bar over 200 (200*1.01 = 202)
			// - not enough to unseat the modelled trinket.
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999999}}): 201,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 60, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != modelledEffectItemID {
		t.Fatalf("trinket1 pick = %+v, want the modelled trinket (%d) despite the unmodelled one measuring higher raw dps", out["trinket1"].Item, modelledEffectItemID)
	}
	if out["trinket1"].Item.MeasuredDPS != 200 {
		t.Errorf("trinket1 pick MeasuredDPS = %v, want 200 (its own measured dps, not the unmodelled one's)", out["trinket1"].Item.MeasuredDPS)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != 999999 {
		t.Fatalf("trinket1 runner-up = %+v, want the unmodelled trinket (999999): it measured the higher raw dps even though it did not win", out["trinket1"].RunnerUp)
	}
	if out["trinket1"].RunnerUp.MeasuredDPS != 201 {
		t.Errorf("trinket1 runner-up MeasuredDPS = %v, want 201", out["trinket1"].RunnerUp.MeasuredDPS)
	}
}

// The mirror case: once the unmodelled trinket's own measured dps
// clears the noise floor, it is real enough to win outright.
func TestRankTrinketSlotUnmodelledTrinketWinsWhenItClearsTheMargin(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(modelledEffectItemID, "Real Combat Trinket", 60, "On use: deals damage"),
			trinketWithEffect(999999, "Serenity Field", 60, "Equip: restores spirit over time"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: modelledEffectItemID}}): 200,
			// 205 clears 200*1.01 = 202: a real, not noise-level, lead.
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999999}}): 205,
		},
	}
	out, _ := rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 60, "", picks, bySlot, "trinket1", trinketTestWeights)
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 999999 {
		t.Fatalf("trinket1 pick = %+v, want the unmodelled trinket (999999): it cleared the margin", out["trinket1"].Item)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != modelledEffectItemID {
		t.Fatalf("trinket1 runner-up = %+v, want the modelled trinket (%d)", out["trinket1"].RunnerUp, modelledEffectItemID)
	}
}

// With no modelled candidate in the pool at all, the margin rule has
// nothing to defer to - the highest measured dps wins exactly like
// before this lane's own fix, even though both candidates are
// unmodelled.
func TestRankTrinketSlotPicksHighestWhenEveryCandidateIsUnmodelled(t *testing.T) {
	picks := map[string]slotPick{}
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(999998, "Unmodelled A", 60, "Equip: does something unmodelled"),
			trinketWithEffect(999999, "Unmodelled B", 60, "Equip: does something else unmodelled"),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999998}}): 100,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 999999}}): 100.5,
		},
	}
	out, _ := rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 60, "", picks, bySlot, "trinket1", trinketTestWeights)
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != 999999 {
		t.Fatalf("trinket1 pick = %+v, want item 999999 (higher measured dps, no modelled alternative to defer to)", out["trinket1"].Item)
	}
}

// This lane's brief (bis-ranker-integrity-6), item 3: hunter-beast-
// mastery band 60 Alliance's own repro - the two REAL top candidates
// (Thunderbrew's Boot Flask, +8 Spirit; Frozen Heart of the Mountain,
// +9 Hit) both measure a genuine, positive, stat-driven gain, but both
// ALSO carry an unrelated effect_text the engine does not implement.
// The OLD version of this margin check walked past both of them
// looking for the first MODELLED candidate however far down the list
// (here, a third item whose own implemented effect is a pure
// self-heal - zero DPS relevance, measuring exactly the no-trinket
// baseline) and crowned that objectively worse, zero-gain candidate
// instead. The margin check must only ever compare the top TWO
// candidates: when the immediate runner-up is ALSO unmodelled, there
// is no real "modelled safety net" to defer to, so the plain highest
// measured dps wins, exactly as it would if a third, weaker modelled
// candidate were not in the pool at all.
func TestRankTrinketSlotDoesNotSkipPastMultipleUnmodelledCandidatesForAWeakModelledOne(t *testing.T) {
	picks := map[string]slotPick{}
	const (
		thunderbrewID  = 744    // real Blackrock Depths trinket, +8 Spirit, unimplemented use-effect
		dormantHeartID = 249473 // real trinket, healer use-effect the engine leaves unregistered (era-item-effects-3)
	)
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(thunderbrewID, "Thunderbrew's Boot Flask", 44, "Deals 75 Fire damage... Gets you quite drunk too!"),
			trinketWithEffect(dormantHeartID, "Dormant Heart of the Mountain", 55, "Your casts of Greater Heal, Healing Touch, Healing Wave, or Holy Light..."),
			trinketWithEffect(modelledEffectItemID, "Darkmoon Card: Heroism", 66, "Sometimes heals bearer of 150 damage when damaging an enemy in melee."),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: thunderbrewID}}):        200.77,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: dormantHeartID}}):       200.47,
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: modelledEffectItemID}}): 200.00,
			// The no-trinket baseline (swapSlot's own itemID-0 gear):
			// 200.00, matching the modelled candidate's own value -
			// its implemented effect (a self-heal) contributes nothing
			// to DPS, exactly like Darkmoon Card: Heroism's own real
			// measurement in this lane's brief repro.
			gearKey([]api.GearSlot{}): 200.00,
		},
	}
	out, _ := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 60, "", picks, bySlot, "trinket1", trinketTestWeights)
	if out["trinket1"].Item == nil || out["trinket1"].Item.ID != thunderbrewID {
		t.Fatalf("trinket1 pick = %+v, want Thunderbrew's Boot Flask (%d): the raw highest measured dps, with no genuinely-modelled candidate immediately behind it to defer to", out["trinket1"].Item, thunderbrewID)
	}
	if out["trinket1"].Item.MeasuredGainDPS < trinketZeroGainThresholdDPS {
		t.Errorf("trinket1 pick MeasuredGainDPS = %v, want >= %v (a real, positive, stat-driven gain, not the zero-gain modelled candidate)", out["trinket1"].Item.MeasuredGainDPS, trinketZeroGainThresholdDPS)
	}
	if out["trinket1"].RunnerUp == nil || out["trinket1"].RunnerUp.ID != dormantHeartID {
		t.Fatalf("trinket1 runner-up = %+v, want Frozen Heart of the Mountain (%d): the next-highest real measurement", out["trinket1"].RunnerUp, dormantHeartID)
	}
}

// This lane's brief (bis-ranker-integrity-6), item 5: the end-to-end
// shape main.go's own trinket loop runs - trinket1's rankTrinketSlot
// call, THEN trinket2's, sharing one picks map - reproducing
// shaman-elemental band 40 / druid-balance band 40's own repro: every
// trinket in this pool ties at score 0 (trinkets carry no scorable
// stats - trinkets.go's own package doc), so pick()'s own tiebreak
// (lowest item id) would have landed a real, valuable candidate
// (Ankh of Life, id 1713, LOWER than every other candidate here on
// purpose) on trinket2 as a bare placeholder before either slot's own
// real tournament ever ran. Without clearTrinketPlaceholders, trinket1's
// OWN rankTrinketSlot call (running first) would read that placeholder
// as an already-decided pair-mate and wrongly exclude Ankh of Life from
// its own shortlist - trinket1's only OTHER candidate is a genuine
// zero-value stat-stick, so the slot would empty outright even though
// Ankh of Life clearly deserves ONE of the two trinket slots. With the
// placeholders cleared first, trinket1's own shortlist sees Ankh of
// Life fairly, picks it (the only real value in the pool), and
// trinket2 - now correctly excluding Ankh of Life as trinket1's REAL,
// final pick - is left with only the zero-value stat-stick and empties
// honestly instead.
func TestTrinketLoopDoesNotLetTrinket1ExcludeTrinket2sStalePlaceholder(t *testing.T) {
	const (
		ankhOfLifeID = 1713 // lowest id in the pool - pick()'s own tiebreak would land it on trinket2 first
		zeroValueID  = 21565
	)
	bySlot := map[string][]scored{
		"trinket1": {
			trinketWithEffect(ankhOfLifeID, "Ankh of Life", 45, "Reincarnates the user."),
			trinket(zeroValueID, "Rune of Perfection", 45),
		},
		"trinket2": {
			trinketWithEffect(ankhOfLifeID, "Ankh of Life", 45, "Reincarnates the user."),
			trinket(zeroValueID, "Rune of Perfection", 45),
		},
	}
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: ankhOfLifeID}}): 105,
			gearKey([]api.GearSlot{{Slot: "trinket2", ItemID: ankhOfLifeID}}): 105,
			// Rune of Perfection and the no-trinket baseline all tie at
			// 100 - a genuine zero-value stat-stick, matching this
			// lane's own repro exactly.
		},
		DefaultDPS: 100,
	}

	// pick()'s own placeholder, exactly as main.go's runSpec builds it
	// before the trinket loop runs: lowest item id wins the tiebreak
	// (every trinket ties score 0), landing Ankh of Life on trinket2.
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: zeroValueID, Name: "Rune of Perfection"}}},
		"trinket2": {Item: &scored{candidate: candidate{ID: ankhOfLifeID, Name: "Ankh of Life"}}},
	}
	picks = clearTrinketPlaceholders(picks)

	var notes []string
	for _, slot := range []string{"trinket1", "trinket2"} {
		var slotNotes []string
		picks, slotNotes = rankTrinketSlot(fake, specInfo{}, "dwarf", "paladin", 40, "", picks, bySlot, slot, trinketTestWeights)
		notes = append(notes, slotNotes...)
	}
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	if picks["trinket1"].Item == nil || picks["trinket1"].Item.ID != ankhOfLifeID {
		t.Fatalf("trinket1 = %+v, want Ankh of Life (%d): the only real value in the pool, fairly considered", picks["trinket1"].Item, ankhOfLifeID)
	}
	// rankTrinketSlot itself never empties a slot (report.go's own
	// trinketLowGain gate does that later, from MeasuredGainDPS) - the
	// contract this test protects is narrower: trinket2's own real
	// tournament must correctly exclude Ankh of Life as trinket1's REAL
	// pick (not the stale placeholder) and be left with only the
	// zero-value stat-stick, whose own measured gain is genuinely 0 -
	// exactly what later lets buildReport empty it honestly.
	if picks["trinket2"].Item == nil || picks["trinket2"].Item.ID != zeroValueID {
		t.Fatalf("trinket2 = %+v, want the remaining zero-value candidate (%d)", picks["trinket2"].Item, zeroValueID)
	}
	if !picks["trinket2"].Item.GainMeasured || picks["trinket2"].Item.MeasuredGainDPS >= trinketZeroGainThresholdDPS {
		t.Fatalf("trinket2 pick GainMeasured/MeasuredGainDPS = %v/%v, want a measured, genuinely-zero gain", picks["trinket2"].Item.GainMeasured, picks["trinket2"].Item.MeasuredGainDPS)
	}
}

func TestTrinketEffectUnmodelled(t *testing.T) {
	cases := []struct {
		name string
		c    candidate
		want bool
	}{
		{"no effect at all is not unmodelled (pure stats)", candidate{ID: 1}, false},
		{"a real, implemented effect is modelled", candidate{ID: modelledEffectItemID, EffectText: "On use: deals damage"}, false},
		{"an effect the engine does not implement is unmodelled", candidate{ID: 999999, EffectText: "Equip: restores spirit"}, true},
	}
	for _, tc := range cases {
		if got := trinketEffectUnmodelled(tc.c); got != tc.want {
			t.Errorf("%s: trinketEffectUnmodelled(%+v) = %v, want %v", tc.name, tc.c, got, tc.want)
		}
	}
}

// TestRankTrinketSlotComputesMeasuredGainStdErrInQuadrature is this
// lane's brief, item 2: MeasuredGainStdErr must combine the winning
// candidate's own run error and the no-trinket baseline's own run
// error the ordinary way for a difference of two independent means -
// sqrt(a^2 + b^2) - not just copy one or the other, so report.go's
// trinketGainSignificant gate has a real combined error to check
// MeasuredGainDPS against.
func TestRankTrinketSlotComputesMeasuredGainStdErrInQuadrature(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder"}}},
	}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Only Candidate", 20)},
	}
	winnerGear := gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}})
	baselineGear := gearKey(nil) // swapSlot's own itemID-0 shape: no other slot is set in picks, so the baseline gear list is empty.
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			winnerGear:   210,
			baselineGear: 200,
		},
		StdErrByGear: map[string]float64{
			winnerGear:   0.6,
			baselineGear: 0.5,
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	item := out["trinket1"].Item
	if item == nil || item.ID != 2 {
		t.Fatalf("trinket1 pick = %+v, want item 2", item)
	}
	if !item.GainMeasured {
		t.Fatal("GainMeasured = false, want true")
	}
	if diff := item.MeasuredGainDPS - 10; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("MeasuredGainDPS = %v, want 10 (210-200)", item.MeasuredGainDPS)
	}
	want := math.Sqrt(0.6*0.6 + 0.5*0.5)
	if diff := item.MeasuredGainStdErr - want; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("MeasuredGainStdErr = %v, want %v (sqrt(0.6^2+0.5^2), combined in quadrature)", item.MeasuredGainStdErr, want)
	}
}

// This lane's brief (bis-ranker-integrity-12), item 1: trinketAdaptiveGain
// is the mechanism fix - a gain sitting close to trinketGainSignificant's
// own bar can swing across it purely from noise, so the winning
// candidate's own gain gets re-measured at escalating iteration counts
// until its error is comfortably below it. TestTrinketAdaptiveGain*
// below exercise the function directly (no rankTrinketSlot/shortlist
// scaffolding needed); TestRankTrinketSlotAdaptiveGain* below that
// confirm the real call site wires it in.
func TestTrinketAdaptiveGainSkipsWhenAlreadyPrecise(t *testing.T) {
	fake := &fakeEngine{}
	picks := map[string]slotPick{}
	gain, stdErr, err := trinketAdaptiveGain(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, "trinket1", 2, 10, 2)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if gain != 10 || stdErr != 2 {
		t.Fatalf("trinketAdaptiveGain = (%v, %v), want unchanged (10, 2): already comfortably below gain/4 (2.5)", gain, stdErr)
	}
	if len(fake.Calls) != 0 {
		t.Fatalf("Calls = %v, want none - an already-precise gain must never re-sim at all", fake.Calls)
	}
}

func TestTrinketAdaptiveGainSkipsWhenGainIsNotPositive(t *testing.T) {
	fake := &fakeEngine{}
	picks := map[string]slotPick{}
	gain, stdErr, err := trinketAdaptiveGain(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, "trinket1", 2, -3, 5)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if gain != -3 || stdErr != 5 {
		t.Fatalf("trinketAdaptiveGain = (%v, %v), want unchanged (-3, 5): a non-positive gain is never a noise-near-the-bar case", gain, stdErr)
	}
	if len(fake.Calls) != 0 {
		t.Fatalf("Calls = %v, want none", fake.Calls)
	}
}

// TestTrinketAdaptiveGainEscalatesUntilPrecise is this lane's own
// repro shape: the flat trinketRankIterations (100) pass leaves the
// gain's error too close to the gain itself, so the loop doubles the
// iteration count (100 -> 200 -> 400) until the combined error clears
// gain/4, re-measuring both the candidate and the no-trinket baseline
// each round (gainStdErr's own quadrature rule).
func TestTrinketAdaptiveGainEscalatesUntilPrecise(t *testing.T) {
	winnerGear := gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}})
	baselineGear := gearKey(nil)
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			winnerGear:   210,
			baselineGear: 200,
		},
		StdErrFunc: func(req api.SimRequest) float64 {
			switch req.Iterations {
			case 200:
				return 2.0
			case 400:
				return 0.4
			default:
				t.Fatalf("unexpected iteration count %d requested", req.Iterations)
				return 0
			}
		},
	}
	picks := map[string]slotPick{}
	// Entry point: the 100-iteration ordering pass's own measurement -
	// gain 10, stdErr 3.5 (>= 10/4 = 2.5, so escalation must start).
	gain, stdErr, err := trinketAdaptiveGain(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, "trinket1", 2, 10, 3.5)
	if err != nil {
		t.Fatalf("err = %v, want nil", err)
	}
	if gain != 10 {
		t.Errorf("gain = %v, want 10 (210-200, unchanged across rounds)", gain)
	}
	wantStdErr := math.Sqrt(0.4*0.4 + 0.4*0.4)
	if diff := stdErr - wantStdErr; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("stdErr = %v, want %v (sqrt(0.4^2+0.4^2) at the 400-iteration round, where it first clears gain/4)", stdErr, wantStdErr)
	}
	if len(fake.Calls) != 4 {
		t.Fatalf("Calls = %v, want exactly 4 (candidate+baseline at 200, then again at 400 - the 200 round did not clear gain/4)", fake.Calls)
	}
}

// TestTrinketAdaptiveGainStopsAtCeiling is a trinket whose own error
// never shrinks enough (weights.go's trinketGainSignificanceMultiplier
// doc: a genuinely zero-relevance trinket's gain does not shrink
// toward zero as iterations rise either) - the loop must still
// terminate at trinketGainAdaptiveIterationCeiling rather than escalate
// forever, and return its last measurement (not significant, but not
// this function's job to force significance) without an error.
func TestTrinketAdaptiveGainStopsAtCeiling(t *testing.T) {
	var maxIterationsSeen int
	fake := &fakeEngine{
		DPSByGear: map[string]float64{
			gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}): 205,
			gearKey(nil): 200,
		},
		StdErrFunc: func(req api.SimRequest) float64 {
			if req.Iterations > maxIterationsSeen {
				maxIterationsSeen = req.Iterations
			}
			return 4.0 // never shrinks below gain/4 (5/4 = 1.25)
		},
	}
	picks := map[string]slotPick{}
	gain, stdErr, err := trinketAdaptiveGain(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, "trinket1", 2, 5, 4.0)
	if err != nil {
		t.Fatalf("err = %v, want nil (a ceiling is not a failure)", err)
	}
	if gain != 5 {
		t.Errorf("gain = %v, want 5 (205-200)", gain)
	}
	wantStdErr := math.Sqrt(4.0*4.0 + 4.0*4.0)
	if diff := stdErr - wantStdErr; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("stdErr = %v, want %v (last measurement at the ceiling)", stdErr, wantStdErr)
	}
	if maxIterationsSeen != trinketGainAdaptiveIterationCeiling {
		t.Errorf("max iterations requested = %d, want exactly the ceiling (%d), never beyond it", maxIterationsSeen, trinketGainAdaptiveIterationCeiling)
	}
}

// TestTrinketAdaptiveGainReturnsOriginalMeasurementOnRerunFailure: a
// re-run that fails (an engine-side error, not a scoring one) must not
// discard the ordering pass's own valid measurement - the caller
// (rankTrinketSlot) falls back to it and logs a note instead.
func TestTrinketAdaptiveGainReturnsOriginalMeasurementOnRerunFailure(t *testing.T) {
	fake := &fakeEngine{
		FailGear: gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}}),
	}
	picks := map[string]slotPick{}
	gain, stdErr, err := trinketAdaptiveGain(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, "trinket1", 2, 10, 3.5)
	if err == nil {
		t.Fatal("err = nil, want the forced engine failure")
	}
	if gain != 10 || stdErr != 3.5 {
		t.Errorf("(gain, stdErr) = (%v, %v), want the original (10, 3.5) unchanged on a failed re-run", gain, stdErr)
	}
}

// TestRankTrinketSlotAdaptiveGainRefinesTheWinnersOwnGain is the real
// call site (rankTrinketSlot's own tail): the ordering pass measures a
// noisy near-the-bar gain at trinketRankIterations, and the adaptive
// refinement above narrows it before report.go's trinketGainSignificant
// ever sees it.
func TestRankTrinketSlotAdaptiveGainRefinesTheWinnersOwnGain(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder"}}},
	}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Only Candidate", 20)},
	}
	winnerGear := gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}})
	baselineGear := gearKey(nil)
	fake := &fakeEngine{
		DPSByGear: map[string]float64{winnerGear: 210, baselineGear: 200},
		StdErrFunc: func(req api.SimRequest) float64 {
			if req.Iterations <= trinketRankIterations {
				return 3.5 // the ordering pass's own noisy measurement
			}
			return 0.1 // any escalated round is precise
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 0 {
		t.Fatalf("notes = %v, want none", notes)
	}
	item := out["trinket1"].Item
	if item == nil || item.ID != 2 {
		t.Fatalf("trinket1 pick = %+v, want item 2", item)
	}
	wantStdErr := math.Sqrt(0.1*0.1 + 0.1*0.1)
	if diff := item.MeasuredGainStdErr - wantStdErr; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("MeasuredGainStdErr = %v, want %v - the ordering pass's own noisy 3.5 must not survive into the published gain", item.MeasuredGainStdErr, wantStdErr)
	}
	if !trinketGainSignificant(item.MeasuredGainDPS, item.MeasuredGainStdErr) {
		t.Errorf("trinketGainSignificant(%v, %v) = false, want true once adaptively refined", item.MeasuredGainDPS, item.MeasuredGainStdErr)
	}
}

// TestRankTrinketSlotAdaptiveGainFailureKeepsOriginalMeasurementAndNotes
// confirms the graceful fallback: an adaptive re-run failure never
// costs the slot its ordering-pass measurement, only adds a note.
func TestRankTrinketSlotAdaptiveGainFailureKeepsOriginalMeasurementAndNotes(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Placeholder"}}},
	}
	bySlot := map[string][]scored{
		"trinket1": {trinket(2, "Only Candidate", 20)},
	}
	winnerGear := gearKey([]api.GearSlot{{Slot: "trinket1", ItemID: 2}})
	baselineGear := gearKey(nil)
	fake := &fakeEngine{
		DPSByGear:    map[string]float64{winnerGear: 210, baselineGear: 200},
		StdErrByGear: map[string]float64{winnerGear: 3.5, baselineGear: 0},
		DPSFunc: func(req api.SimRequest) (float64, error) {
			key := gearKey(req.Character.Gear)
			if req.Iterations > trinketRankIterations {
				return 0, errors.New("fakeEngine: forced adaptive-round failure")
			}
			if key == winnerGear {
				return 210, nil
			}
			return 200, nil
		},
	}
	out, notes := rankTrinketSlot(fake, specInfo{}, "dwarf", "hunter", 20, "", picks, bySlot, "trinket1", trinketTestWeights)
	if len(notes) != 1 || !strings.Contains(notes[0], "adaptive gain re-measurement") {
		t.Fatalf("notes = %v, want exactly one naming the failed adaptive re-measurement", notes)
	}
	item := out["trinket1"].Item
	if !item.GainMeasured || item.MeasuredGainDPS != 10 || item.MeasuredGainStdErr != 3.5 {
		t.Fatalf("item gain = (measured=%v dps=%v stdErr=%v), want the ordering pass's own (true, 10, 3.5) kept on a failed re-run", item.GainMeasured, item.MeasuredGainDPS, item.MeasuredGainStdErr)
	}
}
