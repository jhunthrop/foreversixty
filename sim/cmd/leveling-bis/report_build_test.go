package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestTitleCase(t *testing.T) {
	cases := map[string]string{"horde": "Horde", "alliance": "Alliance", "": "", "a": "A"}
	for in, want := range cases {
		if got := titleCase(in); got != want {
			t.Errorf("titleCase(%q) = %q, want %q", in, got, want)
		}
	}
}

func reportSpec() specInfo {
	return specInfo{Spec: "hunter-marksmanship", ReferenceStat: "ranged_attack_power", WeightStats: []string{"ranged_attack_power", "agility"}}
}

func TestBuildReportFirstBandHasNoPreviousSoEveryPickIsNew(t *testing.T) {
	picks := map[string]slotPick{
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10, HasSource: true, Source: itemSource{Kind: "quest", Label: "A Quest"}}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "0500000", 5, map[string]api.StatWeight{"agility": {Stat: "agility", Weight: 1.5, Error: 0.1}}, reportSpec().WeightStats, picks, 500, nil, nil, nil, 1.2, 3.4, nil, nil, nil, 0, "")
	if r.Band != 20 || r.Faction != "horde" || r.Race != "troll" || r.TalentPoints != 5 {
		t.Fatalf("buildReport base fields wrong: %+v", r)
	}
	if r.SetDPS != 500 || r.WeightsRunSeconds != 1.2 || r.VerifyRunSeconds != 3.4 {
		t.Fatalf("buildReport timing/dps fields wrong: %+v", r)
	}
	if len(r.Slots) != len(slotOrder) {
		t.Fatalf("len(r.Slots) = %d, want %d (one row per slotOrder entry)", len(r.Slots), len(slotOrder))
	}
	var headRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			headRow = s
		}
	}
	if headRow.ItemID != 1 || headRow.ItemName != "Helm" || !headRow.Verified {
		t.Fatalf("head row = %+v, want item 1 verified (no swap result recorded)", headRow)
	}
	if headRow.Source != "A Quest" || headRow.SourceKind != "quest" {
		t.Fatalf("head row source = %q/%q, want A Quest/quest", headRow.Source, headRow.SourceKind)
	}
	if len(r.NewAtBand) != 1 || !strings.Contains(r.NewAtBand[0], "Helm") {
		t.Fatalf("NewAtBand = %v, want Helm listed (no previous band)", r.NewAtBand)
	}
}

func TestBuildReportUnchangedFromPreviousBandIsNotNew(t *testing.T) {
	// Score: 10 (not the zero-value default) so this row clears the
	// zero-value gate above and actually publishes with an ItemID -
	// new_at_band is computed from the final, gated rows (this lane's
	// brief, item 7), not the raw pick, so a zero-scoring item here
	// would be emptied before new_at_band ever saw it.
	item1 := &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10}
	picks := map[string]slotPick{"head": {Item: item1}}
	previous := map[string]slotPick{"head": {Item: item1}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, previous, 0, 0, nil, nil, nil, 0, "")
	if len(r.NewAtBand) != 0 {
		t.Fatalf("NewAtBand = %v, want empty: item 1 unchanged from the previous band", r.NewAtBand)
	}
}

func TestBuildReportChangedFromPreviousBandIsNew(t *testing.T) {
	picks := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 2, Name: "Better Helm"}, Score: 10}}}
	previous := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10}}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, previous, 0, 0, nil, nil, nil, 0, "")
	if len(r.NewAtBand) != 1 || !strings.Contains(r.NewAtBand[0], "Better Helm") {
		t.Fatalf("NewAtBand = %v, want Better Helm listed", r.NewAtBand)
	}
}

// This lane's brief (bis-ranker-integrity-6), item 7: the mage-fire
// band 20 Alliance repro - new_at_band named "neck: Sentinel's
// Medallion", "trinket1: Rune of Perfection" and "trinket2: Rune of
// Duty" while every one of those slots' own row published empty
// (score() valued each at exactly 0 with no effect_text, the same
// zero-value gate TestBuildReportPublishesEmptySlotWithNoDPSValueReason
// exercises). new_at_band must never name a slot its own row does not
// actually publish a pick for.
func TestBuildReportNewAtBandNeverNamesAnEmptiedSlot(t *testing.T) {
	picks := map[string]slotPick{
		// A real, kept pick (Score != 0): should be named.
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Real Helm"}, Score: 10}},
		// Zero-scoring, no effect_text, no trinket/weapon exemption -
		// the zero-value gate empties this row entirely.
		"neck":     {Item: &scored{candidate: candidate{ID: 2, Name: "Sentinel's Medallion"}, Score: 0}},
		"trinket1": {Item: &scored{candidate: candidate{ID: 3, Name: "Rune of Perfection"}, Score: 0}},
		"trinket2": {Item: &scored{candidate: candidate{ID: 4, Name: "Rune of Duty"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "alliance", "human", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	for _, slot := range []string{"neck", "trinket1", "trinket2"} {
		if byslot[slot].ItemID != 0 {
			t.Fatalf("test setup wrong: %s row = %+v, want emptied by the zero-value gate", slot, byslot[slot])
		}
	}
	for _, entry := range r.NewAtBand {
		slot, _, _ := strings.Cut(entry, ":")
		row, ok := byslot[slot]
		if !ok || row.ItemID == 0 {
			t.Errorf("new_at_band entry %q names a slot with no published pick - row = %+v", entry, row)
		}
	}
	if len(r.NewAtBand) != 1 || !strings.Contains(r.NewAtBand[0], "Real Helm") {
		t.Fatalf("NewAtBand = %v, want only head: Real Helm", r.NewAtBand)
	}
}

func TestBuildReportSwapBeatenRowIsTheWinnerVerifiedWithNote(t *testing.T) {
	// applySwaps has already promoted the runner-up into Item and demoted
	// the scored pick to RunnerUp before buildReport sees the picks.
	winner := &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}
	beaten := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	picks := map[string]slotPick{"head": {Item: winner, RunnerUp: beaten}}
	swaps := []swapResult{{Slot: "head", SwapDPS: 200, BaselineDPS: 150, Beat: true}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 200, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var headRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			headRow = s
		}
	}
	if headRow.ItemID != 2 || headRow.ItemName != "Better Helm" {
		t.Fatalf("head row = %d %q, want the measured winner Better Helm (2)", headRow.ItemID, headRow.ItemName)
	}
	if !headRow.Verified {
		t.Fatal("head row Verified = false, want true: the winner was measured by the very swap run")
	}
	if !strings.Contains(headRow.SwapNote, "Helm (id 1)") || !strings.Contains(headRow.SwapNote, "200.0") || !strings.Contains(headRow.SwapNote, "150.0") {
		t.Fatalf("head row SwapNote = %q, want it to name the beaten pick and both set DPS figures", headRow.SwapNote)
	}
}

// This lane's brief, item 5: when this row's own SwapDPS really is the
// band's own final SetDPS (the promotion was the last word - the
// exact shape TestBuildReportSwapBeatenRowIsTheWinnerVerifiedWithNote
// pins), DPSDelta is still published alongside the trustworthy "vs"
// pair - a consumer that only reads DPSDelta gets the same real number
// either way.
func TestBuildReportSwapBeatPublishesDPSDelta(t *testing.T) {
	winner := &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}
	beaten := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	picks := map[string]slotPick{"head": {Item: winner, RunnerUp: beaten}}
	swaps := []swapResult{{Slot: "head", SwapDPS: 200, BaselineDPS: 150, Beat: true}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 200, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var headRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			headRow = s
		}
	}
	if headRow.DPSDelta == nil || *headRow.DPSDelta != 50 {
		t.Fatalf("head row DPSDelta = %v, want 50 (200 - 150)", headRow.DPSDelta)
	}
}

// This lane's brief, item 5's own named example: paladin-retribution 60
// Alliance published a header Set DPS of 182.9 while a promoted row's
// own swap comparison read "180.8 vs 176.4" - a pair that does not
// involve 182.9 at all, because this slot's promotion was measured
// BEFORE some other slot's own promotion changed the band's real
// final total. When the row's own SwapDPS does not match the band's
// final SetDPS, SwapNote must publish the delta ALONE - never a "vs"
// pair the header cannot corroborate.
func TestBuildReportSwapBeatOmitsUncorroboratedAbsoluteNumbers(t *testing.T) {
	winner := &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}
	beaten := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	picks := map[string]slotPick{"head": {Item: winner, RunnerUp: beaten}}
	swaps := []swapResult{{Slot: "head", SwapDPS: 180.8, BaselineDPS: 176.4, Beat: true}}
	// The band's own final SetDPS (182.9), measured AFTER other slots'
	// own promotions this same run - it does not match this row's own
	// 180.8, so that number is stale, not the finished set's own total.
	r := buildReport(reportSpec(), 60, "alliance", "human", "", 0, nil, nil, picks, 182.9, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var headRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			headRow = s
		}
	}
	if strings.Contains(headRow.SwapNote, "180.8") || strings.Contains(headRow.SwapNote, "176.4") {
		t.Fatalf("head row SwapNote = %q, want no uncorroborated absolute numbers (180.8/176.4 do not match the band's own final SetDPS 182.9)", headRow.SwapNote)
	}
	if !strings.Contains(headRow.SwapNote, "+4.4 DPS") {
		t.Fatalf("head row SwapNote = %q, want the real delta (+4.4 DPS, 180.8 - 176.4) published alone", headRow.SwapNote)
	}
	if headRow.DPSDelta == nil {
		t.Fatal("head row DPSDelta is nil, want the real delta published even when the absolute numbers are suppressed")
	}
	if diff := *headRow.DPSDelta - 4.4; diff > 1e-9 || diff < -1e-9 {
		t.Fatalf("head row DPSDelta = %v, want 4.4 (180.8 - 176.4)", *headRow.DPSDelta)
	}
}

// The mirror image for a "kept the pick" (not promoted) row: its own
// BaselineDPS must match the band's own final SetDPS before the raw
// "vs" pair is trustworthy enough to publish.
func TestBuildReportKeptPickOmitsUncorroboratedAbsoluteNumbers(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Butcher's Cleaver"}, Score: 236.9}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Diamond Hammer"}, Score: 232.8}
	// off_hand already wears the runner-up's own item - buildAlternatives
	// excludes it as main_hand's pair-mate (this lane's own doc,
	// TestBuildReportAddsSwapNoteWhenVerifiedButNoVisibleEvidence),
	// which is what keeps this row's Alternatives empty of real
	// evidence and lets the "confirmed by the sim" fallback fire below.
	offHandPick := &scored{candidate: candidate{ID: 2, Name: "Diamond Hammer"}, Score: 232.8}
	picks := map[string]slotPick{
		"main_hand": {Item: pick, RunnerUp: runnerUp},
		"off_hand":  {Item: offHandPick},
	}
	bySlot := map[string][]scored{
		"main_hand": {
			{candidate: candidate{ID: 1, Name: "Butcher's Cleaver"}, Score: 236.9},
			{candidate: candidate{ID: 2, Name: "Diamond Hammer"}, Score: 232.8},
		},
	}
	swaps := []swapResult{{Slot: "main_hand", SwapDPS: 174.2, BaselineDPS: 178.4, Beat: false}}
	// The band's final SetDPS (182.9) does not match this row's own
	// pre-promotion baseline (178.4) - some OTHER slot's own promotion
	// this band raised the real total past what this slot's own swap
	// pass ever measured.
	r := buildReport(reportSpec(), 60, "alliance", "human", "", 0, nil, nil, picks, 182.9, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			row = s
		}
	}
	if strings.Contains(row.SwapNote, "178.4") || strings.Contains(row.SwapNote, "174.2") {
		t.Fatalf("main_hand row SwapNote = %q, want no uncorroborated absolute numbers", row.SwapNote)
	}
	if !strings.Contains(row.SwapNote, "+4.2 DPS") {
		t.Fatalf("main_hand row SwapNote = %q, want the real delta (+4.2 DPS, 178.4 - 174.2) published alone", row.SwapNote)
	}
}

// This lane's brief, item 6: druid-feral band 50 Alliance trinket2
// published its own pick, Frozen Heart of the Mountain, at sim_dps
// 124.4 (rankTrinketSlot's own tournament number, measured before the
// rest of the band's picks were final) while its own verified
// alternative, Smoking Heart of the Mountain, carried sim_dps 154.8 -
// a published pick must never look weaker than a verified alternative
// sitting on its very own row. verifyBand's own later swap pass
// (verify.go) re-measured this exact pick against the actually-
// finished gear; buildReport must publish THAT number, not the
// earlier tournament's stale one.
func TestBuildReportRefreshesStaleSimDPSFromTheLaterSwapBaseline(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 249469, Name: "Frozen Heart of the Mountain"}, MeasuredDPS: 124.4}
	runnerUp := &scored{candidate: candidate{ID: 11811, Name: "Smoking Heart of the Mountain"}, MeasuredDPS: 154.8}
	picks := map[string]slotPick{"trinket2": {Item: pick, RunnerUp: runnerUp}}
	bySlot := map[string][]scored{
		"trinket2": {
			{candidate: candidate{ID: 249469, Name: "Frozen Heart of the Mountain"}},
			{candidate: candidate{ID: 11811, Name: "Smoking Heart of the Mountain"}},
		},
	}
	// The later verifyBand swap pass re-measured this exact pick
	// (BaselineDPS) against the finished gear; the runner-up did not
	// beat it (Beat: false), matching the real druid-feral row's own
	// dps_delta of -0.97 on the alternative.
	swaps := []swapResult{{Slot: "trinket2", SwapDPS: 154.8, BaselineDPS: 155.8, Beat: false}}
	r := buildReport(reportSpec(), 50, "alliance", "night_elf", "", 0, nil, nil, picks, 155.8, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket2" {
			row = s
		}
	}
	if row.SimDPS != 155.8 {
		t.Fatalf("trinket2 row.SimDPS = %v, want 155.8 (the later, closer-to-finished baseline), not the stale 124.4 tournament snapshot", row.SimDPS)
	}
	for _, a := range row.Alternatives {
		if a.Verified && a.SimDPS > row.SimDPS {
			t.Fatalf("verified alternative %+v carries sim_dps higher than its own pick's row.SimDPS %v - the exact contradiction this lane's brief item 6 fixes", a, row.SimDPS)
		}
	}
}

func TestBuildReportSwapLostKeepsSlotVerified(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Worse Helm"}}
	picks := map[string]slotPick{"head": {Item: pick, RunnerUp: runnerUp}}
	swaps := []swapResult{{Slot: "head", SwapDPS: 50, Beat: false}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 150, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var headRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			headRow = s
		}
	}
	if !headRow.Verified || headRow.SwapNote != "" {
		t.Fatalf("head row = %+v, want verified with no swap note", headRow)
	}
}

func TestBuildReportVerifyErrorMarksSlotUnconfirmed(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Other Helm"}}
	picks := map[string]slotPick{"head": {Item: pick, RunnerUp: runnerUp}}
	verifyErrors := []string{"head: runner-up Other Helm (id 2): the engine reported an error: boom"}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 150, nil, nil, nil, 0, 0, verifyErrors, nil, nil, 0, "")
	var headRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			headRow = s
		}
	}
	if headRow.Verified {
		t.Fatal("head row Verified = true, want false: its swap sim errored")
	}
	if !strings.Contains(headRow.SwapNote, "engine-side error") {
		t.Fatalf("head row SwapNote = %q, want the engine-side-error explanation", headRow.SwapNote)
	}
	if r.VerifyErrors[0] != verifyErrors[0] {
		t.Fatalf("r.VerifyErrors = %v, want it carried through unchanged", r.VerifyErrors)
	}
}

func TestBuildReportNoSourceCountAndSample(t *testing.T) {
	var noSource []candidate
	for i := 1; i <= noSourceSampleSize+5; i++ {
		noSource = append(noSource, candidate{ID: i, Name: "Unsourced"})
	}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, noSource, nil, 0, 0, nil, nil, nil, 0, "")
	if r.NoSourceCount != len(noSource) {
		t.Fatalf("NoSourceCount = %d, want %d", r.NoSourceCount, len(noSource))
	}
	if len(r.NoSourceSample) != noSourceSampleSize {
		t.Fatalf("len(NoSourceSample) = %d, want %d (bounded)", len(r.NoSourceSample), noSourceSampleSize)
	}
}

func TestBuildReportWeightsFollowOrder(t *testing.T) {
	weights := map[string]api.StatWeight{
		"agility":             {Stat: "agility", Weight: 2.5, Error: 0.1},
		"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0, Error: 0.05},
	}
	order := []string{"ranged_attack_power", "agility"}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	if len(r.Weights) != 2 || r.Weights[0].Stat != "ranged_attack_power" || r.Weights[1].Stat != "agility" {
		t.Fatalf("r.Weights = %+v, want order preserved", r.Weights)
	}
	if r.Weights[0].Weight != 1.0 || r.Weights[1].Weight != 2.5 {
		t.Fatalf("r.Weights = %+v, wrong values", r.Weights)
	}
}

func TestBuildReportFlagsInsignificantWeights(t *testing.T) {
	// The owner's own repro (2026-09-28, hunter-marksmanship band 20):
	// melee_haste read 14.87 with an error wide enough that report.go's
	// own significanceErrorFraction (25%) rejects it, while agility's
	// tight error passes.
	weights := map[string]api.StatWeight{
		"agility":     {Stat: "agility", Weight: 2.05, Error: 0.1},
		"melee_haste": {Stat: "melee_haste", Weight: 14.87, Error: 6.0},
	}
	order := []string{"agility", "melee_haste"}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byStat := map[string]weightRow{}
	for _, w := range r.Weights {
		byStat[w.Stat] = w
	}
	if byStat["agility"].Insignificant {
		t.Errorf("agility (error 0.1 on weight 2.05, %.0f%% below threshold) reported insignificant", significanceErrorFraction*100)
	}
	if !byStat["melee_haste"].Insignificant {
		t.Errorf("melee_haste (error 6.0 on weight 14.87, over threshold) reported significant")
	}
	if byStat["melee_haste"].Error != 6.0 {
		t.Errorf("melee_haste.Error = %v, want 6.0 carried through unchanged", byStat["melee_haste"].Error)
	}
}

// This lane's brief (bis-ranker-integrity-6), item 1: warrior-arms/
// fury band 40-60's own repro -- the reference stat (attack_power,
// weight always 1.00 by construction) published Insignificant: true
// because its error happened to be wide enough to fail
// isWeightSignificant's generic 25%-of-value bar, even though
// referenceMeasurementReason (weights.go) already trusted this exact
// band's reference measurement (weightsReason == ""). The reference
// row IS that trusted measurement, so it must publish significant
// regardless of its own error bar, while an ordinary non-reference
// stat with the same wide error still gets flagged.
func TestBuildReportReferenceRowIsSignificantWhenBandIsTrusted(t *testing.T) {
	weights := map[string]api.StatWeight{
		// The reference stat's own row: Weight is always 1.0 by
		// construction, and its Error here (0.43) fails the 25% bar
		// isWeightSignificant applies to an ordinary stat.
		"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.00, Error: 0.43},
		// An ordinary stat with the identical error shape must still
		// be flagged -- this test only exempts the reference row.
		"agility": {Stat: "agility", Weight: 1.00, Error: 0.43},
	}
	order := []string{"ranged_attack_power", "agility"}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0.5, "")
	byStat := map[string]weightRow{}
	for _, w := range r.Weights {
		byStat[w.Stat] = w
	}
	if byStat["ranged_attack_power"].Insignificant {
		t.Errorf("reference stat ranged_attack_power reported insignificant while the band's own reference measurement (weightsReason == \"\") is trusted")
	}
	if !byStat["agility"].Insignificant {
		t.Errorf("agility (same error shape, not the reference stat) reported significant - only the reference row should be exempted")
	}
}

func TestBuildReportFlagsEffectUnmodelledOnAnUnimplementedProcButNotAnImplementedOne(t *testing.T) {
	picks := map[string]slotPick{
		// 424242 is not a real item id: its effect can never be
		// engine-implemented (rank.go's hasImplementedEffect).
		"main_hand": {Item: &scored{candidate: candidate{ID: 424242, Name: "Unmodelled Sword", EffectText: "Does something nobody coded."}}},
		// 3854 (Frost Tiger Blade) is real, from this lane's own
		// effectids_generated.go.
		"ranged": {Item: &scored{candidate: candidate{ID: 3854, Name: "Frost Tiger Blade", EffectText: "Launches a bolt of frost."}}},
		// A plain item with no effect_text at all must never be flagged.
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Plain Helm"}}},
	}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	if !byslot["main_hand"].EffectUnmodelled {
		t.Error("main_hand (unimplemented proc) EffectUnmodelled = false, want true")
	}
	if byslot["ranged"].EffectUnmodelled {
		t.Error("ranged (implemented proc) EffectUnmodelled = true, want false")
	}
	if byslot["head"].EffectUnmodelled {
		t.Error("head (no effect_text at all) EffectUnmodelled = true, want false")
	}
}

func TestBuildReportFlagsEffectUnmodelledOnARelicTheEngineDoesNotImplement(t *testing.T) {
	// night-relic-exempt: a relic (libram/idol/totem) carries zero
	// stats, so score() alone can never distinguish two candidates in
	// the ranged slot -- the engine-verified effect ranking (rank.go)
	// is the only thing that can, and only for a relic whose spell id
	// is in effectids_generated.go. 30101 is not a real item id, so its
	// effect can never be engine-implemented: the report must flag it
	// effect_unmodelled rather than silently reporting score 0 as if
	// that were a real answer about which relic is best.
	picks := map[string]slotPick{
		"ranged": {Item: &scored{candidate: candidate{
			ID: 30101, Name: "Libram of Effect Only", ClassID: armorClassID,
			SubclassID: armorSublibramID, Stats: map[string]float64{},
			EffectText: "Reduces the cast time of Holy Light.",
		}}},
	}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	if !byslot["ranged"].EffectUnmodelled {
		t.Error("ranged (relic with an unimplemented effect) EffectUnmodelled = false, want true")
	}
}

// This lane's brief, defect 4: a slot's tied alternatives (pick.go's
// own Ties field) must reach the published row, so the page can show
// "or Blackwater Cutlass" instead of implying the lowest-id winner was
// uniquely best.
func TestBuildReportCarriesTiedAlternatives(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {
			Item: &scored{candidate: candidate{ID: 1, Name: "Rusty Sword"}, Score: 8.28},
			Ties: []scored{
				{candidate: candidate{ID: 2, Name: "Blackwater Cutlass"}, Score: 8.28},
				{candidate: candidate{ID: 3, Name: "Bent Blade"}, Score: 8.28},
			},
		},
		"head": {Item: &scored{candidate: candidate{ID: 4, Name: "Plain Helm"}, Score: 12}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	got := byslot["main_hand"].Ties
	want := []tieAlternative{{ItemID: 2, ItemName: "Blackwater Cutlass"}, {ItemID: 3, ItemName: "Bent Blade"}}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("main_hand Ties = %+v, want %+v", got, want)
	}
	if len(byslot["head"].Ties) != 0 {
		t.Fatalf("head Ties = %+v, want none: no tied alternative was recorded", byslot["head"].Ties)
	}
}

// TestBuildReportDropsATieARealSimMeasuredAsALoss is this lane's brief
// (bis-ranker-integrity-14), item 3: hunter-beast-mastery/hunter-
// marksmanship band 20 hands published Serpent Gloves as TYING Gloves
// of the Fang (an exact score() match) while the row's own
// Alternatives entry for Gloves of the Fang showed dps_delta -0.08 -
// buildAlternatives' own swap-override publishing the REAL,
// sim-measured result of verify.go's swap pass against pk.RunnerUp
// (Gloves of the Fang is both the score tie AND the one candidate a
// real sim actually tested). A score tie and a real, above-noise
// measured loss cannot both be true on the same row: the tie must be
// dropped once real evidence contradicts it (reconcileTies, report.go).
func TestBuildReportDropsATieARealSimMeasuredAsALoss(t *testing.T) {
	glovesOfTheFang := scored{candidate: candidate{ID: 101, Name: "Gloves of the Fang"}, Score: 10}
	picks := map[string]slotPick{
		"hands": {
			Item:     &scored{candidate: candidate{ID: 100, Name: "Serpent Gloves"}, Score: 10},
			RunnerUp: &glovesOfTheFang,
			Ties:     []scored{glovesOfTheFang},
		},
	}
	// BaselineDPS (the kept pick's own measured DPS) minus SwapDPS
	// (Gloves of the Fang's own measured DPS) is a real -0.08, well
	// past swapDeltaNoiseFloorDPS (0.05) - not a measurement artifact.
	swaps := []swapResult{{Slot: "hands", BaselineDPS: 100.08, SwapDPS: 100.00, Beat: false}}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 100.08, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	hands := byslot["hands"]
	if len(hands.Ties) != 0 {
		t.Fatalf("hands.Ties = %+v, want none: a real sim measured a -0.08 DPS loss for the only tie, contradicting the score tie", hands.Ties)
	}
	var fang *alternativeRow
	for i := range hands.Alternatives {
		if hands.Alternatives[i].ItemID == 101 {
			fang = &hands.Alternatives[i]
		}
	}
	if fang == nil {
		t.Fatal("hands.Alternatives has no row for Gloves of the Fang (id 101)")
	}
	if !fang.Verified || math.Abs(fang.DPSDelta-(-0.08)) > 1e-9 {
		t.Fatalf("Gloves of the Fang alternative = %+v, want Verified=true DPSDelta=-0.08 (the real swap measurement)", *fang)
	}
}

// TestBuildReportKeepsATieTheSimNeverContradicted is
// TestBuildReportDropsATieARealSimMeasuredAsALoss's own control case:
// a genuine tie with no real-sim evidence against it (no swap ran for
// this slot at all) must still publish as a tie - reconcileTies only
// drops a tie a real measurement actually contradicts, never a tie it
// simply has no opinion about.
func TestBuildReportKeepsATieTheSimNeverContradicted(t *testing.T) {
	picks := map[string]slotPick{
		"hands": {
			Item: &scored{candidate: candidate{ID: 100, Name: "Serpent Gloves"}, Score: 10},
			Ties: []scored{{candidate: candidate{ID: 101, Name: "Gloves of the Fang"}, Score: 10}},
		},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var hands slotRow
	for _, s := range r.Slots {
		if s.Slot == "hands" {
			hands = s
		}
	}
	want := []tieAlternative{{ItemID: 101, ItemName: "Gloves of the Fang"}}
	if len(hands.Ties) != 1 || hands.Ties[0] != want[0] {
		t.Fatalf("hands.Ties = %+v, want %+v (no sim ran, so the score tie stands)", hands.Ties, want)
	}
}

// This lane's brief, item 2 (bis-ranker-integrity, 2026-09-29):
// Sentinel's Medallion (Agility/Stamina) scoring exactly 0 for a caster
// spec must publish empty with empty_reason "no_dps_value", not as a
// "verified" pick - the exact defect every caster spec's own band-20
// neck carried before this fix.
func TestBuildReportPublishesEmptySlotWithNoDPSValueReason(t *testing.T) {
	picks := map[string]slotPick{
		"neck": {Item: &scored{candidate: candidate{ID: 1, Name: "Sentinel's Medallion"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var neckRow slotRow
	for _, s := range r.Slots {
		if s.Slot == "neck" {
			neckRow = s
		}
	}
	if neckRow.ItemID != 0 || neckRow.ItemName != "" {
		t.Fatalf("neck row = %+v, want empty (no item published)", neckRow)
	}
	if neckRow.EmptyReason != noDPSValueReason {
		t.Fatalf("neck row EmptyReason = %q, want %q", neckRow.EmptyReason, noDPSValueReason)
	}
}

// This lane's brief, item 1: a weapon slot (band.go's weaponSlots) is
// NEVER emptied by the score-zero gate, unlike every other slot -
// instead the row stands, published, with LowValue set so the page
// can say so honestly instead of implying a genuine BiS win.
func TestBuildReportFlagsLowValueOnAZeroScoringWeaponInsteadOfEmptying(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "Notched Shortsword"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			row = s
		}
	}
	if row.ItemID != 1 || row.EmptyReason != "" {
		t.Fatalf("main_hand row = %+v, want item 1 published (never emptied), no EmptyReason", row)
	}
	if !row.LowValue {
		t.Fatalf("main_hand row = %+v, want LowValue true (score() found this weapon worth exactly 0)", row)
	}
}

// A weapon slot that scores something real (not 0) must never carry
// LowValue - it is a genuine, measured pick, not a fallback.
func TestBuildReportDoesNotFlagLowValueOnARealScoringWeapon(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "A Real Sword"}, Score: 12.5}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			row = s
		}
	}
	if row.LowValue {
		t.Fatalf("main_hand row = %+v, want LowValue false for a genuinely-scoring weapon", row)
	}
}

// This lane's brief, item 8: every empty slot carries an
// empty_reason. off_hand with no pick at all, because main_hand
// equipped a two-hander, must publish "two_hand_equipped" - before
// this fix it published a bare {"slot": "off_hand"} with no reason
// (hunter-marksmanship band 40 alliance's own named case).
func TestBuildReportEmptyOffHandUnderTwoHanderCarriesTwoHandEquippedReason(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "A Greatsword", TwoHand: true}, Score: 10}},
		"off_hand":  {},
	}
	r := buildReport(reportSpec(), 40, "alliance", "human", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "off_hand" {
			row = s
		}
	}
	if row.EmptyReason != twoHandEquippedReason {
		t.Fatalf("off_hand row EmptyReason = %q, want %q", row.EmptyReason, twoHandEquippedReason)
	}
}

// An empty slot with no two-hander in play at all (nothing was ever
// sourced for it under this spec's own picking pool) must still carry
// SOME empty_reason, never a bare row - this lane's brief, item 8's
// own general contract test.
func TestBuildReportEmptySlotWithNoTwoHanderCarriesNoSourcedItemReason(t *testing.T) {
	picks := map[string]slotPick{
		"waist": {},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "waist" {
			row = s
		}
	}
	if row.EmptyReason != noSourcedItemReason {
		t.Fatalf("waist row EmptyReason = %q, want %q", row.EmptyReason, noSourcedItemReason)
	}
}

// Contract test, this lane's brief item 8, literally: every row in a
// built report either publishes an item or carries a non-empty
// EmptyReason - never both empty (a bare "nothing here" the page
// cannot explain).
func TestBuildReportContractEveryRowHasAnItemOrAnEmptyReason(t *testing.T) {
	picks := map[string]slotPick{
		"main_hand": {Item: &scored{candidate: candidate{ID: 1, Name: "A Greatsword", TwoHand: true}, Score: 10}},
		"off_hand":  {},
		"neck":      {Item: &scored{candidate: candidate{ID: 2, Name: "Sentinel's Medallion"}, Score: 0}},
		"waist":     {},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	for _, row := range r.Slots {
		if row.ItemID == 0 && row.EmptyReason == "" {
			t.Errorf("slot %s: no item and no EmptyReason - every empty slot must carry a reason", row.Slot)
		}
	}
}

// This lane's brief, item 6: a whole-report contract, not just the one
// druid-feral repro - no row's own published pick may carry a lower
// sim_dps than one of its own verified alternatives, whatever slot or
// tournament decided it. Two slots here: one already fixed by the
// stale-SimDPS refresh (trinket2, this lane's own dogfood number), one
// a plain swap-promoted row (head) where the invariant already held
// before this lane touched anything - both must hold at once.
func TestBuildReportContractNoVerifiedAlternativeOutranksItsOwnPick(t *testing.T) {
	trinketPick := &scored{candidate: candidate{ID: 249469, Name: "Frozen Heart of the Mountain"}, MeasuredDPS: 124.4}
	trinketRunnerUp := &scored{candidate: candidate{ID: 11811, Name: "Smoking Heart of the Mountain"}, MeasuredDPS: 154.8}
	// applySwaps always sets a promoted item's own MeasuredDPS to
	// sw.SwapDPS before buildReport ever sees it (verify.go) - set here
	// to match what the real pipeline hands buildReport.
	headWinner := &scored{candidate: candidate{ID: 20, Name: "Better Helm"}, MeasuredDPS: 200}
	headBeaten := &scored{candidate: candidate{ID: 21, Name: "Helm"}}
	picks := map[string]slotPick{
		"trinket2": {Item: trinketPick, RunnerUp: trinketRunnerUp},
		"head":     {Item: headWinner, RunnerUp: headBeaten},
	}
	bySlot := map[string][]scored{
		"trinket2": {
			{candidate: candidate{ID: 249469, Name: "Frozen Heart of the Mountain"}},
			{candidate: candidate{ID: 11811, Name: "Smoking Heart of the Mountain"}},
		},
	}
	swaps := []swapResult{
		{Slot: "trinket2", SwapDPS: 154.8, BaselineDPS: 155.8, Beat: false},
		{Slot: "head", SwapDPS: 200, BaselineDPS: 150, Beat: true},
	}
	r := buildReport(reportSpec(), 50, "alliance", "night_elf", "", 0, nil, nil, picks, 155.8, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0, "")
	for _, row := range r.Slots {
		pickDPS := row.SimDPS
		for _, a := range row.Alternatives {
			if a.Verified && a.SimDPS > pickDPS {
				t.Errorf("slot %s: verified alternative %+v carries sim_dps %v, higher than its own pick's row.SimDPS %v", row.Slot, a, a.SimDPS, pickDPS)
			}
		}
	}
}

// A trinket slot's own Score is near-meaningless (score() cannot value a
// proc/on-use effect at all - trinkets.go's own doc): rankTrinketSlot's
// real-sim decision is authoritative regardless of what score() says, so
// a trinket at Score 0 must NOT be emptied.
// A trinket with a real effect (score() cannot value it at all - a
// proc/on-use effect carries no scorable stats) stays exempt from the
// score-zero gate even at Score 0 - bis-ranker-integrity-2 lane, item
// 4's own kept half: EffectText is what makes this trinket different
// from a bare stat-stick, not merely being a trinket.
func TestBuildReportKeepsATrinketWithAnEffectAtScoreZero(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "A Proc Trinket", EffectText: "Chance on hit: deals damage"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 1 || row.EmptyReason != "" {
		t.Fatalf("trinket1 row = %+v, want item 1 published with no EmptyReason (an effect_text trinket is exempt from the score-zero gate)", row)
	}
}

// This lane's brief, item 4: Rune of Perfection (+6 spell penetration,
// +4 stamina) and Rune of Duty (pure resistance) carry no effect_text
// and no stat this fixture's spec weighs - a bare stat-stick trinket
// with nothing score() or a real sim could ever value is exactly as
// zero-value as Sentinel's/Scout's Medallion (the original neck-slot
// finding), and must empty the same way, not stay published as
// "verified" BiS just because it is a trinket.
func TestBuildReportEmptiesATrinketWithNoEffectAndNoValueAtScoreZero(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "Rune of Perfection", Stats: map[string]float64{"spell_penetration": 6, "stamina": 4}}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 0 || row.EmptyReason != noDPSValueReason {
		t.Fatalf("trinket1 row = %+v, want empty with EmptyReason %q (no effect_text, no valued stat)", row, noDPSValueReason)
	}
}

// A candidate whose effect the engine does NOT implement (EffectUnmodelled)
// might have real, unmeasured value score() cannot see either - it must
// not be emptied just because its plain stat score happens to be 0.
func TestBuildReportKeepsAnEffectUnmodelledPickAtScoreZero(t *testing.T) {
	picks := map[string]slotPick{
		"back": {Item: &scored{candidate: candidate{ID: 1, Name: "Mystery Cloak", EffectText: "Does something unimplemented"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "back" {
			row = s
		}
	}
	if row.ItemID != 1 || row.EmptyReason != "" {
		t.Fatalf("back row = %+v, want item 1 published with no EmptyReason (effect_unmodelled is exempt)", row)
	}
}

// This lane's brief, item 1 (owner defect A, 2026-09-29): a slot's
// alternatives must carry the next-best sourced candidates after the
// pick, so a player who cannot get the pick's own source has a real
// fallback instead of an empty page - warrior-arms horde band 20's
// main_hand published Forsaken Greataxe with no record that Smite's
// Mighty Hammer (a Deadmines drop) was ever considered.
//
// Owner review, tenet 8 (2026-09-29): the final list is sorted by
// dps_delta, descending, not "ties then score order" - the two agree
// whenever nothing outscores the pick (the overwhelmingly common case:
// bySlot's own list is best-score-first, so the pick IS the top entry
// and every tie sits at the exact same score right behind it).
// Hammerbone here DOES outscore this test's pick with no swap
// correction applied (this fixture deliberately leaves sw nil), but
// this lane's brief (bis-ranker-integrity, 2026-09-29) caps exactly
// this case at 0 rather than publishing the raw positive estimate as a
// fact: nothing has actually simmed Hammerbone against the pick (see
// TestBuildAlternativesSwapBeatOverridesTheDemotedRunnerUpsDelta below
// for the same numbers WITH that real-sim correction, where Hammerbone's
// delta is real and allowed to be negative), so its unverified estimate
// ties with 0 - a real tie (Tied Greataxe) and an unverified "maybe
// better" both read as "at best equal to the pick" - and the two tie-
// break by item id, ascending, same as every other tie in this file.
func TestBuildAlternativesSortsByDPSDeltaDescendingTiesIncluded(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Forsaken Greataxe"}, Score: 302.9}
	pk := slotPick{
		Item: pick,
		// A tie: scores identically to the pick, dps_delta exactly 0.
		Ties: []scored{
			{candidate: candidate{ID: 5, Name: "Tied Greataxe"}, Score: 302.9, Source: itemSource{Kind: "quest", Label: "Quests"}},
		},
	}
	// bySlot's own score-sorted pool for this slot: best first, per
	// candidatesBySlot's contract.
	list := []scored{
		{candidate: candidate{ID: 6, Name: "Hammerbone"}, Score: 306.05, Source: itemSource{Kind: "quest", Label: "Quests"}},
		{candidate: candidate{ID: 1, Name: "Forsaken Greataxe"}, Score: 302.9, Source: itemSource{Kind: "quest", Label: "Quests"}}, // the pick itself
		{candidate: candidate{ID: 7230, Name: "Smite's Mighty Hammer"}, Score: 297.82, Source: itemSource{Kind: "dungeon", Label: "The Deadmines: Mr. Smite"}},
		{candidate: candidate{ID: 8, Name: "Living Root"}, Score: 296.94, Source: itemSource{Kind: "dungeon", Label: "Wailing Caverns: Verdan the Everliving"}},
		{candidate: candidate{ID: 9, Name: "Too Far Down"}, Score: 100, Source: itemSource{Kind: "quest", Label: "Quests"}},
	}
	// referenceDPSPerPoint 1.0 here so DPSDelta reads as the raw
	// score-unit gap - the dedicated DPS-conversion contract test below
	// (TestBuildAlternativesConvertsScoreDeltaToRealDPS) is where a
	// non-trivial referenceDPSPerPoint is exercised; this test is about
	// ordering and exclusion, not arithmetic.
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 1.0, nil, 0)
	want := []alternativeRow{
		// Tied Greataxe (id 5) and Hammerbone (id 6) both read DPSDelta 0
		// (an unverified positive estimate is capped, not published),
		// so they tie-break by item id ascending.
		{ItemID: 5, ItemName: "Tied Greataxe", SourceKind: "quest", Source: "Quests", DPSDelta: 0},
		{ItemID: 6, ItemName: "Hammerbone", SourceKind: "quest", Source: "Quests", DPSDelta: 0},
		{ItemID: 7230, ItemName: "Smite's Mighty Hammer", SourceKind: "dungeon", Source: "The Deadmines: Mr. Smite", DPSDelta: 297.82 - 302.9},
	}
	if len(got) != len(want) {
		t.Fatalf("buildAlternatives = %+v (%d entries), want %d", got, len(got), len(want))
	}
	for i := range want {
		if got[i].ItemID != want[i].ItemID {
			t.Errorf("alternative[%d].ItemID = %d, want %d (%+v)", i, got[i].ItemID, want[i].ItemID, got[i])
			continue
		}
		if diff := got[i].DPSDelta - want[i].DPSDelta; diff > 1e-9 || diff < -1e-9 {
			t.Errorf("alternative[%d] (%s) DPSDelta = %v, want %v", i, got[i].ItemName, got[i].DPSDelta, want[i].DPSDelta)
		}
		if got[i].Source != want[i].Source || got[i].SourceKind != want[i].SourceKind {
			t.Errorf("alternative[%d] (%s) source = %q/%q, want %q/%q", i, got[i].ItemName, got[i].Source, got[i].SourceKind, want[i].Source, want[i].SourceKind)
		}
	}
}

// Owner review, tenet 8 (2026-09-29): dps_delta must be real DPS, not
// a bare score-unit number with nothing saying so. Contract: a 5-point
// score gap at this band's own reference_dps_per_point of 0.0444
// converts to 0.22 DPS (5 * 0.0444 = 0.222, rounds to 0.22) - the
// exact figure the owner's own repro named.
func TestBuildAlternativesConvertsScoreDeltaToRealDPS(t *testing.T) {
	const referenceDPSPerPoint = 0.0444
	pick := &scored{candidate: candidate{ID: 1, Name: "Pick"}, Score: 302.91}
	pk := slotPick{Item: pick}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Pick"}, Score: 302.91},
		{candidate: candidate{ID: 2, Name: "Five Points Back"}, Score: 302.91 - 5},
	}
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, referenceDPSPerPoint, nil, 0)
	if len(got) != 1 {
		t.Fatalf("buildAlternatives = %+v, want exactly 1 alternative", got)
	}
	wantDPSDelta := -5 * referenceDPSPerPoint
	if diff := got[0].DPSDelta - wantDPSDelta; diff > 1e-9 || diff < -1e-9 {
		t.Fatalf("DPSDelta = %v, want %v (score_delta * reference_dps_per_point)", got[0].DPSDelta, wantDPSDelta)
	}
	rounded := math.Round(math.Abs(got[0].DPSDelta)*100) / 100
	if rounded != 0.22 {
		t.Fatalf("|dps_delta| rounded to 2dp = %v, want 0.22 (the owner's own contract number)", rounded)
	}
}

// Owner review, tenet 8: a runner-up verify.go's own swap pass actually
// simmed against the pick, but which LOST that sim, must publish the
// real measured delta (negative) rather than the score estimate that
// made it look better than the winning pick - the exact warrior-arms
// horde band 20 defect (Hammerbone scored 3.14 points above Forsaken
// Greataxe, but measured 29.9 vs 32.2 DPS and lost).
func TestBuildAlternativesSwapBeatOverridesTheDemotedRunnerUpsDelta(t *testing.T) {
	// Forsaken Greataxe: the promoted pick (post-applySwaps, Item is the
	// former runner-up; RunnerUp is the demoted, higher-scoring original
	// pick - buildReport's own doc for why Ties/RunnerUp still name the
	// physical items this way after a promotion).
	promoted := &scored{candidate: candidate{ID: 251533, Name: "Forsaken Greataxe"}, Score: 302.91}
	demoted := &scored{candidate: candidate{ID: 270018, Name: "Hammerbone"}, Score: 306.05}
	pk := slotPick{Item: promoted, RunnerUp: demoted}
	list := []scored{
		{candidate: candidate{ID: 270018, Name: "Hammerbone"}, Score: 306.05, Source: itemSource{Kind: "quest", Label: "Quests"}},
		{candidate: candidate{ID: 251533, Name: "Forsaken Greataxe"}, Score: 302.91, Source: itemSource{Kind: "quest", Label: "Quests"}}, // the pick itself
		{candidate: candidate{ID: 7230, Name: "Smite's Mighty Hammer"}, Score: 297.82, Source: itemSource{Kind: "dungeon", Label: "The Deadmines: Mr. Smite"}},
	}
	// The real swap sim: Forsaken (SwapDPS, the runner-up-at-verify-time
	// that got promoted) measured 32.2, Hammerbone (BaselineDPS, the
	// then-current pick) measured 29.9 - the owner's own numbers.
	sw := &swapResult{Slot: "main_hand", SwapDPS: 32.2, BaselineDPS: 29.9, Beat: true}

	// setDPS matches sw.BaselineDPS (29.9, Hammerbone's own measured
	// value as the demoted former pick) - bis-ranker-integrity-7
	// lane, item 2's own hygiene guard on alternativeRow.SimDPS.
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 0.0444, sw, 29.9)

	var hammerbone, smite *alternativeRow
	for i := range got {
		switch got[i].ItemID {
		case 270018:
			hammerbone = &got[i]
		case 7230:
			smite = &got[i]
		}
	}
	if hammerbone == nil {
		t.Fatalf("buildAlternatives = %+v, want Hammerbone (the demoted runner-up) among the alternatives", got)
	}
	wantDelta := 29.9 - 32.2
	if diff := hammerbone.DPSDelta - wantDelta; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("Hammerbone DPSDelta = %v, want %v (BaselineDPS - SwapDPS, the real measured loss)", hammerbone.DPSDelta, wantDelta)
	}
	if hammerbone.DPSDelta >= 0 {
		t.Errorf("Hammerbone DPSDelta = %v, want negative: it LOST the real sim and must never look better than the pick", hammerbone.DPSDelta)
	}
	if !hammerbone.Verified {
		t.Error("Hammerbone Verified = false, want true: this row's delta came from a real swap sim, not a score estimate")
	}
	// This lane's brief, item 7: the one row a real sim actually
	// measured publishes THAT measurement (SimDPS), not score()'s
	// estimate - Hammerbone is the demoted former pick, so its own
	// measured value is sw.BaselineDPS (29.9). alternativeRow carries no
	// score at all any more (bis-ranker-integrity-3), so there is
	// nothing left to clear.
	if hammerbone.SimDPS != 29.9 {
		t.Errorf("Hammerbone SimDPS = %v, want 29.9 (sw.BaselineDPS, its own measured value as the demoted former pick)", hammerbone.SimDPS)
	}
	// Ordering: after the swap correction, Hammerbone's real delta
	// (-2.3) is worse than Smite's Mighty Hammer's plain score estimate
	// (-5.09 score points ~ -0.226 DPS at this reference_dps_per_point),
	// so Smite's Mighty Hammer must now rank ahead of it.
	if smite == nil {
		t.Fatal("buildAlternatives did not include Smite's Mighty Hammer at all")
	}
	if got[0].ItemID != smite.ItemID {
		t.Errorf("alternatives[0] = %+v, want Smite's Mighty Hammer first (dps_delta descending after the swap correction)", got[0])
	}
}

// The symmetric !Beat case: the runner-up verify.go simmed did NOT
// beat the pick, so no promotion happened, but the sim still measured
// a real delta for it that should override the score estimate the
// same way.
func TestBuildAlternativesSwapNotBeatOverridesTheRunnerUpsDelta(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Pick"}, Score: 100}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Runner Up"}, Score: 95}
	pk := slotPick{Item: pick, RunnerUp: runnerUp}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Pick"}, Score: 100},
		{candidate: candidate{ID: 2, Name: "Runner Up"}, Score: 95, Source: itemSource{Kind: "quest", Label: "Quests"}},
	}
	sw := &swapResult{Slot: "head", SwapDPS: 40, BaselineDPS: 50, Beat: false}
	// setDPS matches sw.SwapDPS (40) - bis-ranker-integrity-7 lane, item
	// 2's own hygiene guard.
	got := buildAlternatives(pk, "head", list, map[string]slotPick{"head": pk}, 0.05, sw, 40)
	if len(got) != 1 || got[0].ItemID != 2 {
		t.Fatalf("buildAlternatives = %+v, want exactly the runner-up", got)
	}
	if !got[0].Verified {
		t.Error("Verified = false, want true")
	}
	wantDelta := 40.0 - 50.0
	if got[0].DPSDelta != wantDelta {
		t.Errorf("DPSDelta = %v, want %v (SwapDPS - BaselineDPS)", got[0].DPSDelta, wantDelta)
	}
	// This lane's brief, item 7: the runner-up here was never promoted
	// (!Beat), so its own measured value is sw.SwapDPS (40) - the
	// number the sim actually measured FOR IT, not the baseline.
	if got[0].SimDPS != 40 {
		t.Errorf("SimDPS = %v, want 40 (sw.SwapDPS, this row's own measured value)", got[0].SimDPS)
	}
}

// TestSwapAlternativeMeasuredDPS pins swapAlternativeMeasuredDPS's own
// two branches directly - this lane's brief, item 7.
func TestSwapAlternativeMeasuredDPS(t *testing.T) {
	beat := swapResult{SwapDPS: 32.2, BaselineDPS: 29.9, Beat: true}
	if got := swapAlternativeMeasuredDPS(beat); got != 29.9 {
		t.Errorf("swapAlternativeMeasuredDPS(%+v) = %v, want 29.9 (BaselineDPS - the demoted former pick's own measured value)", beat, got)
	}
	notBeat := swapResult{SwapDPS: 40, BaselineDPS: 50, Beat: false}
	if got := swapAlternativeMeasuredDPS(notBeat); got != 40 {
		t.Errorf("swapAlternativeMeasuredDPS(%+v) = %v, want 40 (SwapDPS - the still-just-a-runner-up's own measured value)", notBeat, got)
	}
}

// Controller direction, bis-ranker-integrity-9, 2026-09-30 (the ninth
// wow-player sweep's caster finding, day3/player-review-24/casters.md
// finding 1), SUPERSEDING bis-ranker-integrity-8's own direction below:
// mage-arcane band 60's own Weakness Analyzer published dps_delta
// +2.34 next to the KEPT pick Talisman of Ascendance - bis-ranker-
// integrity-8 (2026-09-30) had just fixed swapMeasuredDelta to publish
// exactly this kind of real, well-above-noise positive delta rather
// than force it to 0, reasoning that a genuine few-tenths (or, it
// turns out, several) DPS gap should not be hidden. But a player reads
// ANY positive number next to an alternative as "the site chose the
// worse item", which is never true once !sw.Beat kept the pick on
// purpose (beatsByMargin already decided this candidate did not clear
// swapMargin). The fix: every positive delta in the !sw.Beat branch is
// now an honest tie (0), regardless of size - only a genuine loss
// (delta <= 0) can ever publish a nonzero number from this branch.
func TestSwapMeasuredDeltaNeverPublishesAPositiveNotBeatDelta(t *testing.T) {
	// SwapDPS (100.3) is higher than BaselineDPS (100), but not by
	// enough to clear swapMargin (1%) - Beat is correctly false. The
	// real measured delta (+0.3) clears swapDeltaNoiseFloorDPS (0.05),
	// but a positive delta in this branch is never published as
	// anything other than a tie.
	sw := swapResult{SwapDPS: 100.3, BaselineDPS: 100, Beat: false}
	if got := swapMeasuredDelta(sw); got != 0 {
		t.Errorf("swapMeasuredDelta(%+v) = %v, want exactly 0: a runner-up that did not clear swapMargin is a tie, never a positive number", sw, got)
	}
}

// The floor itself: a delta genuinely too small to trust (here, well
// under swapDeltaNoiseFloorDPS in either direction) still reads as an
// honest, exact tie - the floor exists precisely for this case, not
// for "did not clear swapMargin".
func TestSwapMeasuredDeltaFloorsAGenuinelyTinyDeltaToZero(t *testing.T) {
	cases := []swapResult{
		{SwapDPS: 100.02, BaselineDPS: 100, Beat: false}, // +0.02, under the floor
		{SwapDPS: 99.98, BaselineDPS: 100, Beat: false},  // -0.02, under the floor
	}
	for _, sw := range cases {
		if got := swapMeasuredDelta(sw); got != 0 {
			t.Errorf("swapMeasuredDelta(%+v) = %v, want exactly 0 (below swapDeltaNoiseFloorDPS)", sw, got)
		}
	}
}

// Controller direction's own explicit ask: "a tournament loser inside
// the margin publishes its negative delta." A runner-up that measured
// genuinely LOWER than the pick, by an amount well inside swapMargin's
// own 1% promotion band (so Beat is false regardless), still publishes
// that real, signed, negative number - this direction was never
// clamped by the old code either (only the POSITIVE direction was),
// and stays true under the new floor-based rule as long as the loss
// itself clears swapDeltaNoiseFloorDPS.
func TestSwapMeasuredDeltaTournamentLoserInsideMarginPublishesItsNegativeDelta(t *testing.T) {
	// -0.3 DPS is well inside swapMargin's 1% band around 100 (99..101)
	// but far outside swapDeltaNoiseFloorDPS (0.05) - a real, if small,
	// measured loss.
	sw := swapResult{SwapDPS: 99.7, BaselineDPS: 100, Beat: false}
	got := swapMeasuredDelta(sw)
	want := -0.3
	if diff := got - want; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("swapMeasuredDelta(%+v) = %v, want %v (the real measured loss, published as-is)", sw, got, want)
	}
}

// The Beat branch was already correct (a promoted item's own delta,
// from the DEMOTED item's row, is always negative by construction),
// pinned directly so a future edit cannot regress the sign.
func TestSwapMeasuredDeltaBeatBranchIsNegative(t *testing.T) {
	sw := swapResult{SwapDPS: 110, BaselineDPS: 100, Beat: true}
	got := swapMeasuredDelta(sw)
	want := 100.0 - 110.0
	if got != want {
		t.Errorf("swapMeasuredDelta(%+v) = %v, want %v", sw, got, want)
	}
	if got >= 0 {
		t.Errorf("swapMeasuredDelta(%+v) = %v, want negative: the demoted item genuinely lost", sw, got)
	}
}

// A genuine loss in the !Beat branch (SwapDPS below BaselineDPS) is
// unaffected by the cap - only the positive direction is clamped.
func TestSwapMeasuredDeltaNotBeatGenuineLossIsUnchanged(t *testing.T) {
	sw := swapResult{SwapDPS: 40, BaselineDPS: 50, Beat: false}
	got := swapMeasuredDelta(sw)
	want := 40.0 - 50.0
	if got != want {
		t.Errorf("swapMeasuredDelta(%+v) = %v, want %v (unchanged - only positive deltas are capped)", sw, got, want)
	}
}

// End-to-end pin through buildAlternatives, updated for the
// controller's bis-ranker-integrity-9 direction (2026-09-30),
// SUPERSEDING bis-ranker-integrity-8's own direction (see
// TestSwapMeasuredDeltaNeverPublishesAPositiveNotBeatDelta's own doc
// above for the full history): a runner-up that measured higher but
// did not clear swapMargin's own 1% promotion bar (Beat == false)
// publishes dps_delta 0 (a tie), never the real +0.4 - a positive
// number here reads as "the site chose the worse item", which is
// never true once the pick was kept. Verified stays true regardless:
// this row still came from a real swap sim.
func TestBuildAlternativesPublishesATieWhenNotBeatEvenAboveTheNoiseFloor(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Knight's Leather Pants"}, Score: 100}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Stormshroud Pants"}, Score: 105}
	pk := slotPick{Item: pick, RunnerUp: runnerUp}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Knight's Leather Pants"}, Score: 100},
		{candidate: candidate{ID: 2, Name: "Stormshroud Pants"}, Score: 105, Source: itemSource{Kind: "crafted", Label: "Tailoring"}},
	}
	sw := &swapResult{Slot: "legs", SwapDPS: 100.4, BaselineDPS: 100, Beat: false}
	got := buildAlternatives(pk, "legs", list, map[string]slotPick{"legs": pk}, 0.05, sw, 0)
	if len(got) != 1 || got[0].ItemID != 2 {
		t.Fatalf("buildAlternatives = %+v, want exactly the runner-up", got)
	}
	if got[0].DPSDelta != 0 {
		t.Errorf("Stormshroud Pants DPSDelta = %v, want exactly 0: Beat was false (did not clear swapMargin), so this is a tie, never a positive number", got[0].DPSDelta)
	}
	if !got[0].Verified {
		t.Error("Verified = false, want true: this row still came from a real swap sim")
	}
}

// A finger/trinket/main_hand+off_hand pair-mate is never offered as a
// fallback for the OTHER half of its own pair: the same physical item
// already worn there (verify.go's own pairSlot doc).
func TestBuildAlternativesExcludesThePairMate(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Ring A"}, Score: 10}
	mate := &scored{candidate: candidate{ID: 2, Name: "Ring B"}, Score: 9}
	picks := map[string]slotPick{
		"finger1": {Item: pick},
		"finger2": {Item: mate},
	}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Ring A"}, Score: 10},
		{candidate: candidate{ID: 2, Name: "Ring B"}, Score: 9}, // finger2's own pick - must not appear
		{candidate: candidate{ID: 3, Name: "Ring C"}, Score: 8},
	}
	got := buildAlternatives(picks["finger1"], "finger1", list, picks, 1.0, nil, 0)
	if len(got) != 1 || got[0].ItemID != 3 {
		t.Fatalf("finger1 alternatives = %+v, want only Ring C (finger2's own pick excluded)", got)
	}
}

// TestBuildAlternativesExcludesAnotherIDOfThePicksOwnName is this
// lane's brief (bis-ranker-integrity-12), item 2: Sergeant Major's
// Cape has three real ids (16315 req 25, 16336 req 40, 16337 req 55,
// all rank 9 Alliance) - the same physical reward at three points in a
// player's own PvP progression, but three DIFFERENT item ids. ranker-
// 11 made alternatives unique among THEMSELVES by name
// (dedupeAlternatives), but never checked a candidate's name against
// the PICK's own name, so a lower-req-level id of the identical cape
// reached this row as its own "alternative" (ten such pick-vs-
// alternative pairs across the hybrid specs, this lane's own dogfood:
// the page showed a pick and an alternative both reading "Sergeant
// Major's Cape - PvP rank 9 - Sergeant Major - Alliance").
func TestBuildAlternativesExcludesAnotherIDOfThePicksOwnName(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 16337, Name: "Sergeant Major's Cape"}, Score: 10}
	pk := slotPick{Item: pick}
	list := []scored{
		{candidate: candidate{ID: 16337, Name: "Sergeant Major's Cape"}, Score: 10},
		{candidate: candidate{ID: 16315, Name: "Sergeant Major's Cape"}, Score: 10}, // a lower-req-level id of the SAME cape
		{candidate: candidate{ID: 3, Name: "A Different Back Item"}, Score: 9},
	}
	got := buildAlternatives(pk, "back", list, map[string]slotPick{"back": pk}, 1.0, nil, 0)
	for _, a := range got {
		if a.ItemID == 16315 {
			t.Fatalf("buildAlternatives = %+v, want no alternative sharing the pick's own name (16315 is another id of the picked 16337 - Sergeant Major's Cape)", got)
		}
	}
	if len(got) != 1 || got[0].ItemID != 3 {
		t.Fatalf("buildAlternatives = %+v, want only the genuinely different item (id 3)", got)
	}
}

// The same guard must also hold for the swap-tested runner-up's own
// force-include step, which bypasses the ordinary tie/list loop's
// excluded() check entirely - a runner-up sharing the pick's own name
// must never be force-included either, whatever verify.go's swap pass
// measured for it.
func TestBuildAlternativesExcludesASwapTestedRunnerUpSharingThePicksOwnName(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 16337, Name: "Sergeant Major's Cape"}, Score: 10}
	runnerUp := &scored{candidate: candidate{ID: 16315, Name: "Sergeant Major's Cape"}, Score: 8}
	pk := slotPick{Item: pick, RunnerUp: runnerUp}
	sw := &swapResult{Slot: "back", SwapDPS: 45.0, BaselineDPS: 48.0, Beat: false}
	got := buildAlternatives(pk, "back", nil, map[string]slotPick{"back": pk}, 0, sw, 48.0)
	for _, a := range got {
		if a.ItemID == 16315 {
			t.Fatalf("buildAlternatives = %+v, want the swap-tested runner-up excluded too: it shares the pick's own name", got)
		}
	}
}

// This lane's brief, item 2: hunter-beast-mastery/marksmanship band
// 40/50 main_hand published caster-stat weapons (spell power/
// intellect staves) as alternatives next to a ranged spec's real
// pick - those staves score() exactly 0 for a hunter (none of the
// spec's positively weighted stats appear on them) and were never run
// through any tournament, so nothing about them is a real alternative
// a player could act on. Score 0 and untested must be excluded.
func TestBuildAlternativesExcludesAZeroScoreNeverSimmedCandidate(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Real hunter weapon"}, Score: 50}
	pk := slotPick{Item: pick}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Real hunter weapon"}, Score: 50},
		{candidate: candidate{ID: 2, Name: "Caster Staff of Nothing For You"}, Score: 0, Source: itemSource{Kind: "quest", Label: "Quests"}},
	}
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 1.0, nil, 0)
	if len(got) != 0 {
		t.Fatalf("buildAlternatives = %+v, want none - the only other candidate scores 0 and was never simmed", got)
	}
}

// The same zero-score candidate must still be force-included when it
// really is the one thing a tournament simmed (verify.go's own swap
// pass) - "or a sim-measured delta from a tournament" is an OR, not an
// AND, in this lane's brief item 2. This guards against a future fix
// tightening realAlternative in a way that also blocks the
// swap-tested runner-up's own force-include step.
func TestBuildAlternativesForceIncludesAZeroScoreButSwapTestedRunnerUp(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Pick"}, Score: 50}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Zero Score But Tested"}, Score: 0}
	pk := slotPick{Item: pick, RunnerUp: runnerUp}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Pick"}, Score: 50},
		{candidate: candidate{ID: 2, Name: "Zero Score But Tested"}, Score: 0},
	}
	sw := &swapResult{Slot: "main_hand", SwapDPS: 38.0, BaselineDPS: 45.8, Beat: false}
	// setDPS matches sw.SwapDPS (38.0, the tested runner-up's own
	// measured value here, since !sw.Beat) - bis-ranker-integrity-7
	// lane, item 2's own hygiene guard.
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 0, sw, 38.0)
	var found *alternativeRow
	for i := range got {
		if got[i].ItemID == 2 {
			found = &got[i]
		}
	}
	if found == nil {
		t.Fatalf("buildAlternatives = %+v, want the swap-tested runner-up (id 2) force-included despite scoring 0", got)
	}
	if !found.Verified || found.SimDPS != 38.0 {
		t.Errorf("swap-tested runner-up = %+v, want Verified true and SimDPS 38.0", found)
	}
}

// A slot with no pick (an empty row) and a two-handed main_hand's
// off_hand (which pick.go's enforceTwoHandOffHandInvariant always
// leaves nil) both carry no alternatives - there is no pick to offer
// fallbacks for.
func TestBuildAlternativesNilForAnUnfilledSlot(t *testing.T) {
	list := []scored{{candidate: candidate{ID: 1, Name: "Anything"}, Score: 5}}
	got := buildAlternatives(slotPick{}, "off_hand", list, map[string]slotPick{}, 1.0, nil, 0)
	if got != nil {
		t.Fatalf("buildAlternatives on an unfilled slot = %+v, want nil", got)
	}
}

// This lane's brief: a two-handed main_hand pick's own alternatives
// still offer one-handers mixed in with two-handers (candidatesBySlot's
// own pool, unfiltered by pick()'s per-slot two-hand exclusion) - the
// same list a dual-wield spec's excludeTwoHand call narrows only for
// ITS OWN pick, never for bySlot itself.
func TestBuildAlternativesMainHandMixesOneAndTwoHanders(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Greataxe", TwoHand: true}, Score: 20}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Greataxe", TwoHand: true}, Score: 20},
		{candidate: candidate{ID: 2, Name: "Rusty Sword", TwoHand: false}, Score: 15},
	}
	got := buildAlternatives(slotPick{Item: pick}, "main_hand", list, map[string]slotPick{"main_hand": {Item: pick}}, 1.0, nil, 0)
	if len(got) != 1 || got[0].ItemID != 2 {
		t.Fatalf("main_hand alternatives = %+v, want the one-hander Rusty Sword offered alongside the two-handed pick", got)
	}
}

// End-to-end through buildReport: bySlot flows into row.Alternatives
// for every filled slot, and stays empty for a slot with no pick.
func TestBuildReportWiresAlternativesFromBySlot(t *testing.T) {
	picks := map[string]slotPick{
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Plain Helm"}, Score: 10}},
	}
	bySlot := map[string][]scored{
		"head": {
			{candidate: candidate{ID: 1, Name: "Plain Helm"}, Score: 10},
			{candidate: candidate{ID: 2, Name: "Runner Up Helm"}, Score: 8, Source: itemSource{Kind: "quest", Label: "Quests"}},
		},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, bySlot, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	if len(byslot["head"].Alternatives) != 1 || byslot["head"].Alternatives[0].ItemID != 2 {
		t.Fatalf("head Alternatives = %+v, want Runner Up Helm", byslot["head"].Alternatives)
	}
	if len(byslot["neck"].Alternatives) != 0 {
		t.Fatalf("neck (no pick) Alternatives = %+v, want none", byslot["neck"].Alternatives)
	}
}

// This lane's brief, item 2: reference_dps_per_point flows from
// buildReport's own parameter straight onto the published band.
func TestBuildReportPublishesReferenceDPSPerPoint(t *testing.T) {
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0.0714, "")
	if r.ReferenceDPSPerPoint == nil || *r.ReferenceDPSPerPoint != 0.0714 {
		t.Fatalf("ReferenceDPSPerPoint = %v, want a pointer to 0.0714", r.ReferenceDPSPerPoint)
	}
}

// This lane's brief, item 2: a non-empty weightsReason forces every
// weight row Insignificant and omits ReferenceDPSPerPoint entirely
// (nil, not a misleading 0) - warlock-demonology band 60's own defect
// (reference_dps_per_point -0.189, every weight sign-flipped and none
// marked insignificant on its own noise bar).
func TestBuildReportUntrustworthyReferenceOmitsItAndForcesInsignificant(t *testing.T) {
	weights := map[string]api.StatWeight{
		"spell_power": {Stat: "spell_power", Weight: 1, Error: -0.62},
		"intellect":   {Stat: "intellect", Weight: 4.55, Error: -0.74},
	}
	order := []string{"spell_power", "intellect"}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, -0.189, "reference stat spell_power measured -0.1890 DPS per point - not positive beyond its own error")
	if r.ReferenceDPSPerPoint != nil {
		t.Fatalf("ReferenceDPSPerPoint = %v, want nil (omitted)", *r.ReferenceDPSPerPoint)
	}
	if r.WeightsReason == "" {
		t.Fatal("WeightsReason = \"\", want the reason buildReport was given")
	}
	for _, w := range r.Weights {
		if !w.Insignificant {
			t.Errorf("%s: Insignificant = false, want true - a non-empty weightsReason must force every row", w.Stat)
		}
	}
}

// This lane's brief, item 7: bandReport always documents its own
// Score field's unit, regardless of what the band's picks look like.
func TestBuildReportAlwaysPublishesScoreUnit(t *testing.T) {
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	if r.ScoreUnit != scoreUnitReferenceStatPoints {
		t.Fatalf("ScoreUnit = %q, want %q", r.ScoreUnit, scoreUnitReferenceStatPoints)
	}
}

// A sim-decided pick (MeasuredDPS > 0 - trinkets.go/rank.go/sets.go's
// own tournaments all set this) publishes SimDPS and omits Score -
// this lane's brief, item 7: "one number per row", never both units on
// the same row.
func TestBuildReportSimDecidedPickPublishesSimDPSNotScore(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "A Real Trinket"}, Score: 42, MeasuredDPS: 301.5}},
	}
	// setDPS matches this pick's own MeasuredDPS (bis-ranker-integrity-7
	// lane, item 2: a sim-decided SimDPS is only trustworthy when it
	// still equals the band's finished set_dps - see buildReport's own
	// new hygiene guard) - this test is about the simDecided/Score
	// convention, not about that guard, so setDPS is set to describe
	// the same finished set this pick's own MeasuredDPS does.
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 301.5, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.Score != 0 {
		t.Errorf("Score = %v, want 0 (omitted): this row was sim-decided", row.Score)
	}
	if row.SimDPS != 301.5 {
		t.Errorf("SimDPS = %v, want 301.5", row.SimDPS)
	}
	if row.ItemID != 1 {
		t.Fatalf("row = %+v, want item 1 published (a sim-decided trinket is never emptied)", row)
	}
}

// A score()-decided pick (MeasuredDPS == 0, the overwhelming majority
// of gear) publishes Score exactly as before and never sets SimDPS.
func TestBuildReportScoreDecidedPickPublishesScoreNotSimDPS(t *testing.T) {
	picks := map[string]slotPick{
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "A Real Helm"}, Score: 42}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			row = s
		}
	}
	if row.Score != 42 {
		t.Errorf("Score = %v, want 42", row.Score)
	}
	if row.SimDPS != 0 {
		t.Errorf("SimDPS = %v, want 0 (this row was score-decided)", row.SimDPS)
	}
}

// A swap-promoted pick (verify.go's applySwaps) is sim-decided too:
// its own MeasuredDPS is the promotion's own measured SwapDPS, and its
// row publishes SimDPS, not Score, alongside the existing SwapNote.
func TestBuildReportSwapPromotedPickPublishesSimDPS(t *testing.T) {
	promoted := &scored{candidate: candidate{ID: 1, Name: "The Winner"}, Score: 302.91, MeasuredDPS: 32.2}
	demoted := &scored{candidate: candidate{ID: 2, Name: "The Loser"}, Score: 306.05}
	picks := map[string]slotPick{
		"main_hand": {Item: promoted, RunnerUp: demoted},
	}
	swaps := []swapResult{{Slot: "main_hand", SwapDPS: 32.2, BaselineDPS: 29.9, Beat: true}}
	// setDPS matches this slot's own SwapDPS: this band's promotion was
	// the only one, so applySwaps' own final re-measurement is exactly
	// this trial's own number (bis-ranker-integrity-7 lane, item 2's own
	// hygiene guard - see TestBuildReportSwapPromotedPickOmitsSimDPSWhenAnotherSlotAlsoPromoted
	// for the mismatched case).
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 32.2, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			row = s
		}
	}
	if row.Score != 0 {
		t.Errorf("Score = %v, want 0 (omitted): a swap-promoted pick is sim-decided", row.Score)
	}
	if row.SimDPS != 32.2 {
		t.Errorf("SimDPS = %v, want 32.2 (the promotion's own measured SwapDPS)", row.SimDPS)
	}
	if row.SwapNote == "" {
		t.Error("SwapNote is empty, want the existing promotion note still present")
	}
}

// bis-ranker-integrity-7 lane, item 2: the seventh wow-player sweep's
// hybrids report still found a promoted row's own sim_dps NOT matching
// its band's set_dps - this is the exact code path. applySwaps
// promotes every winning swap independently against the ONE shared
// pre-swap baseline (its own doc): main_hand's own SwapDPS (32.2) was
// measured with ONLY main_hand swapped, but off_hand ALSO promoted this
// same band, so the band's real finished-set total (40.0, applySwaps'
// own final re-measurement with BOTH promotions applied) is not 32.2 at
// all. main_hand's own row must omit sim_dps - its SwapNote/DPSDelta
// (both true of that one single-slot trial regardless) stay published.
func TestBuildReportSwapPromotedPickOmitsSimDPSWhenAnotherSlotAlsoPromoted(t *testing.T) {
	promoted := &scored{candidate: candidate{ID: 1, Name: "The Winner"}, Score: 302.91, MeasuredDPS: 32.2}
	demoted := &scored{candidate: candidate{ID: 2, Name: "The Loser"}, Score: 306.05}
	picks := map[string]slotPick{
		"main_hand": {Item: promoted, RunnerUp: demoted},
	}
	swaps := []swapResult{{Slot: "main_hand", SwapDPS: 32.2, BaselineDPS: 29.9, Beat: true}}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 40.0, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			row = s
		}
	}
	if row.SimDPS != 0 {
		t.Errorf("SimDPS = %v, want 0 (omitted): 32.2 (this slot's own single-swap trial) does not match set_dps 40.0 (another slot also promoted)", row.SimDPS)
	}
	if row.SwapNote == "" {
		t.Error("SwapNote is empty, want the promotion note still present even though sim_dps is withheld")
	}
	if row.DPSDelta == nil {
		t.Fatal("DPSDelta is nil, want 2.3 (32.2 - 29.9, this row's own real measured gain, true regardless of what else promoted)")
	}
	if diff := *row.DPSDelta - 2.3; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("DPSDelta = %v, want 2.3 (32.2 - 29.9, this row's own real measured gain, true regardless of what else promoted)", *row.DPSDelta)
	}
}

func TestWriteSpecReportWritesReadableJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "hunter-marksmanship.json")
	reports := []bandReport{{Spec: "hunter-marksmanship", Band: 20, Faction: "horde"}}
	if err := writeSpecReport(path, "hunter-marksmanship", "testbuild", reports); err != nil {
		t.Fatalf("writeSpecReport: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading written file: %v", err)
	}
	var decoded specReport
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("decoding written file: %v", err)
	}
	if decoded.Spec != "hunter-marksmanship" || decoded.Build != "testbuild" {
		t.Fatalf("decoded = %+v, wrong spec/build", decoded)
	}
	if decoded.EngineVersion == "" || decoded.GeneratedAt == "" {
		t.Fatalf("decoded = %+v, want EngineVersion/GeneratedAt populated", decoded)
	}
	if len(decoded.Bands) != 1 || decoded.Bands[0].Band != 20 {
		t.Fatalf("decoded.Bands = %+v, want the one band passed in", decoded.Bands)
	}
}

func TestWriteMarkdownRendersFactionsSortedWithTablesAndNotes(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "hunter-marksmanship.md")
	spec := specInfo{Spec: "hunter-marksmanship", Name: "Marksmanship", ReferenceStat: "ranged_attack_power"}
	reports := []bandReport{
		{
			Band: 20, Faction: "horde", Race: "troll", Talents: "0500000",
			Weights: []weightRow{{Stat: "ranged_attack_power", Weight: 1.0}},
			Slots: []slotRow{
				{
					Slot: "head", ItemID: 1, ItemName: "Helm", Score: 10, Verified: true, Source: "A Quest", SourceKind: "quest",
					Ties: []tieAlternative{{ItemID: 5, ItemName: "Tied Cap"}},
					Alternatives: []alternativeRow{
						{ItemID: 6, ItemName: "Runner Up Helm", SourceKind: "dungeon", Source: "Some Dungeon", DPSDelta: -2},
						{ItemID: 7, ItemName: "Sim-Verified Cap", SourceKind: "quest", Source: "Quests", DPSDelta: -1.5, Verified: true},
					},
				},
				{Slot: "neck"}, // unpicked slot renders as "-"
			},
			SetDPS: 250.5, NoSourceCount: 2, NoSourceSample: []string{"3 Ghost Item"},
			NewAtBand: []string{"head: Helm"},
		},
		{
			Band: 20, Faction: "alliance", Race: "dwarf", Talents: "0500000",
			Slots: []slotRow{{Slot: "head"}},
		},
	}
	if err := writeMarkdown(path, spec, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading written file: %v", err)
	}
	content := string(b)

	allianceIdx := strings.Index(content, "## Alliance")
	hordeIdx := strings.Index(content, "## Horde")
	if allianceIdx == -1 || hordeIdx == -1 || allianceIdx > hordeIdx {
		t.Fatalf("factions not sorted alphabetically (Alliance before Horde) in:\n%s", content)
	}
	if !strings.Contains(content, "Helm (1)") {
		t.Error("markdown missing the picked item's name/id")
	}
	if !strings.Contains(content, "Helm (1) (or Tied Cap (5))") {
		t.Error("markdown missing the tied alternative (this lane's brief, defect 4)")
	}
	if !strings.Contains(content, "A Quest [quest]") {
		t.Error("markdown missing the item's source label/kind")
	}
	if !strings.Contains(content, "New at 20:") && !strings.Contains(content, "head: Helm") {
		t.Error("markdown missing the New at band line")
	}
	if !strings.Contains(content, "New at this band:** nothing changed") {
		t.Error("markdown missing the no-change fallback line for the alliance band")
	}
	if !strings.Contains(content, "No-known-source sample") {
		t.Error("markdown missing the no-known-source sample line")
	}
	if !strings.Contains(content, "Runner Up Helm (6, -2.00 DPS) [dungeon]") {
		t.Error("markdown missing the head slot's own Alternatives column entry (this lane's brief, item 1)")
	}
	if !strings.Contains(content, "Sim-Verified Cap (7, -1.50 DPS, sim-verified) [quest]") {
		t.Error("markdown missing the sim-verified note on a swap-corrected alternative (owner review, tenet 8)")
	}
}

// TestWriteMarkdownCapsTiedAlternativesAtFiveWithAndNMore is this
// lane's brief (bis-ranker-integrity-14), item 4: priest-shadow band
// 20 main_hand published 107 (alliance) / 117 (horde) tied items in
// one table cell (day3/player-review-33/casters.md finding 8) - the
// live site never reads `ties` at all and caps Alternatives at 3, so
// this .md-only rendering is the only place the wall of names ever
// reached a reader. Capped at mdTieDisplayLimit (5) plus a summary of
// the rest.
func TestWriteMarkdownCapsTiedAlternativesAtFiveWithAndNMore(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "spec.md")
	var ties []tieAlternative
	for i := 0; i < 8; i++ {
		ties = append(ties, tieAlternative{ItemID: 100 + i, ItemName: fmt.Sprintf("Tied Item %d", i)})
	}
	reports := []bandReport{{
		Band: 20, Faction: "horde", Race: "troll",
		Slots: []slotRow{{Slot: "main_hand", ItemID: 1, ItemName: "Evocator's Blade", Score: 10, Verified: true, Ties: ties}},
	}}
	if err := writeMarkdown(path, specInfo{Spec: "priest-shadow", Name: "Shadow", ReferenceStat: "spell_power"}, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	content := string(b)
	for i := 0; i < 5; i++ {
		if !strings.Contains(content, fmt.Sprintf("Tied Item %d (%d)", i, 100+i)) {
			t.Errorf("markdown missing tied item %d, want the first 5 named in full:\n%s", i, content)
		}
	}
	for i := 5; i < 8; i++ {
		if strings.Contains(content, fmt.Sprintf("Tied Item %d (%d)", i, 100+i)) {
			t.Errorf("markdown names tied item %d, want only the first 5 named, the rest summarized:\n%s", i, content)
		}
	}
	if !strings.Contains(content, "and 3 more") {
		t.Errorf("markdown missing the \"and 3 more\" summary for the remaining ties:\n%s", content)
	}
}

// TestWriteMarkdownPrintsSimDecidedInsteadOfAFabricatedZeroScore is
// this lane's brief (bis-ranker-integrity-14), item 4: 46 occurrences
// across 8 caster specs (day3/player-review-33/casters.md finding 9) -
// a pick whose Score was deliberately zeroed because a real sim
// tournament decided it (buildReport's own doc), not a genuinely
// zero-value LowValue weapon, but whose own SimDPS/DPSDelta were ALSO
// both withheld, fell through to the raw reference-points branch and
// printed the literal "0.0 spell_power points (0.00 DPS)" - read by a
// player as "this item does nothing", the exact misreading the
// SimDPS/DPSDelta cases already exist to prevent for every other state
// this row could be in.
func TestWriteMarkdownPrintsSimDecidedInsteadOfAFabricatedZeroScore(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "spec.md")
	refDPSPerPoint := 0.5
	reports := []bandReport{{
		Band: 30, Faction: "alliance", Race: "night-elf",
		ReferenceDPSPerPoint: &refDPSPerPoint,
		Slots:                []slotRow{{Slot: "main_hand", ItemID: 9604, ItemName: "Mechanic's Pipehammer", Score: 0, Verified: true}},
	}}
	if err := writeMarkdown(path, specInfo{Spec: "druid-balance", Name: "Balance", ReferenceStat: "spell_power"}, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	content := string(b)
	if strings.Contains(content, "0.0 spell_power points (0.00 DPS)") {
		t.Errorf("markdown still prints the fabricated zero score:\n%s", content)
	}
	if !strings.Contains(content, "sim-decided") {
		t.Errorf("markdown missing the \"sim-decided\" label for a zero-score, non-LowValue pick:\n%s", content)
	}
}

// TestWriteMarkdownStillPrintsAPlainZeroScoreForALowValueWeapon is the
// above test's own control: a genuinely zero-scoring weapon pick
// (LowValue true, band.go's weaponSlots - score() found no weighted
// stat on it at all, not a sim tournament) is a different, honest
// claim ("this weapon carries none of this spec's weighted stats") and
// must not be relabelled "sim-decided".
func TestWriteMarkdownStillPrintsAPlainZeroScoreForALowValueWeapon(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "spec.md")
	refDPSPerPoint := 0.5
	reports := []bandReport{{
		Band: 30, Faction: "alliance", Race: "night-elf",
		ReferenceDPSPerPoint: &refDPSPerPoint,
		Slots:                []slotRow{{Slot: "main_hand", ItemID: 1, ItemName: "Plain Stick", Score: 0, Verified: true, LowValue: true}},
	}}
	if err := writeMarkdown(path, specInfo{Spec: "druid-balance", Name: "Balance", ReferenceStat: "spell_power"}, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	content := string(b)
	if strings.Contains(content, "sim-decided") {
		t.Errorf("markdown labels a genuinely zero-scoring LowValue weapon \"sim-decided\":\n%s", content)
	}
	if !strings.Contains(content, "0.0 spell_power points (0.00 DPS)") {
		t.Errorf("markdown missing the plain zero-score line for a genuinely zero-scoring LowValue weapon:\n%s", content)
	}
}

// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 2: a
// trinket rankTrinketSlot's own baseline-relative gain (GainMeasured/
// MeasuredGainDPS, trinkets.go) measured BELOW trinketZeroGainThresholdDPS
// must empty exactly like a bare score()-zero trinket does - simDecided
// (MeasuredDPS > 0) alone can never gate this, since MeasuredDPS is the
// whole SET's own absolute DPS, always positive regardless of whether
// the trinket itself contributes anything (warrior-arms band 20's own
// Rune of Perfection, this lane's dogfood: +6 spell penetration/+4
// stamina, no effect_text, real measured gain over an empty trinket
// slot of exactly 0.0).
func TestBuildReportEmptiesATrinketWithLowMeasuredGain(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 21566, Name: "Rune of Perfection", Stats: map[string]float64{"spell_penetration": 6, "stamina": 4}}, MeasuredDPS: 73.65, MeasuredGainDPS: 0, GainMeasured: true}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 0 || row.EmptyReason != noDPSValueReason {
		t.Fatalf("trinket1 row = %+v, want empty with EmptyReason %q (real measured gain 0.0, below the noise floor)", row, noDPSValueReason)
	}
}

// The symmetric case: a real measured gain that clears the noise floor
// must publish, exactly as a genuinely-valuable trinket always has.
func TestBuildReportKeepsATrinketWithMeasuredGainAboveThreshold(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "A Real Trinket"}, MeasuredDPS: 250, MeasuredGainDPS: 12.5, GainMeasured: true}},
	}
	// setDPS matches (bis-ranker-integrity-7 lane, item 2's own hygiene
	// guard, same reasoning as TestBuildReportSimDecidedPickPublishesSimDPSNotScore).
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 250, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 1 || row.EmptyReason != "" {
		t.Fatalf("trinket1 row = %+v, want item 1 published with no EmptyReason (a real 12.5 DPS gain clears the noise floor)", row)
	}
	if row.SimDPS != 250 {
		t.Errorf("SimDPS = %v, want 250", row.SimDPS)
	}
}

// A trinket rankTrinketSlot's own baseline sim failed to measure at all
// (GainMeasured false) must NOT be treated as zero-value - tenet 8:
// never publish a claim ("this trinket is worthless") this command
// could not actually check. It still publishes via the pre-existing
// simDecided exemption (this file's own
// TestBuildReportSimDecidedPickPublishesSimDPSNotScore, unchanged).
func TestBuildReportKeepsATrinketWhoseGainWasNeverMeasured(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 1, Name: "A Real Trinket"}, MeasuredDPS: 250}},
	}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 1 || row.EmptyReason != "" {
		t.Fatalf("trinket1 row = %+v, want item 1 published with no EmptyReason (an unmeasured gain is not evidence of zero value)", row)
	}
}

// TestBuildReportTrinketLowGainKeepsAlternatives is this lane's brief
// (bis-ranker-integrity-12), item 1: the OLD trinketLowGain case
// replaced the whole row with a bare slotRow, discarding
// row.Alternatives (built earlier in buildReport, before this gate
// ever runs) along with the now-hidden pick - shaman-elemental band 50
// Horde's own repro, both trinket slots publishing empty with
// literally zero alternatives even though a real shortlist was simmed;
// paladin-retribution band 50 Alliance's own trinket2 losing Fire Ruby
// entirely as a fallback while Horde's own (not gated) trinket2 row for
// the identical comparison still showed it at -0.14 DPS. A trinket
// whose own gain does not clear the noise floor is still a real,
// checked fact about this band's OTHER shortlisted trinkets, and must
// survive the pick being hidden.
func TestBuildReportTrinketLowGainKeepsAlternatives(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Weak Trinket"}, MeasuredDPS: 100, MeasuredGainDPS: 0.1, MeasuredGainStdErr: 1.0, GainMeasured: true}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Real Runner Up"}}
	picks := map[string]slotPick{"trinket1": {Item: pick, RunnerUp: runnerUp}}
	swaps := []swapResult{{Slot: "trinket1", SwapDPS: 98.0, BaselineDPS: 100.0, Beat: false}}
	r := buildReport(reportSpec(), 50, "horde", "troll", "", 0, nil, nil, picks, 100.0, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 0 || row.EmptyReason != noDPSValueReason {
		t.Fatalf("trinket1 row = %+v, want empty with EmptyReason %q (gain 0.1 does not clear 2x its own stdErr 1.0)", row, noDPSValueReason)
	}
	if len(row.Alternatives) == 0 {
		t.Fatal("trinket1 Alternatives is empty, want the real, swap-verified runner-up (id 2) to survive the pick being hidden")
	}
	found := false
	for _, a := range row.Alternatives {
		if a.ItemID == 2 && a.Verified {
			found = true
		}
	}
	if !found {
		t.Errorf("trinket1 Alternatives = %+v, want a Verified entry for item 2", row.Alternatives)
	}
}

// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 2: a relic
// (libram/idol/totem) whose one real selling point - its engraved
// effect - the engine does not implement at all must empty with its
// own reason instead of falling back to score()'s plain stat total,
// which was never designed to value it either way (druid-balance's own
// Idol of the Huntress: an "Improved Swipe" - a Feral rune - engrave
// with no Balance-relevant stats, picked at 4 of 5 bands because it was
// simply never emptied).
func TestBuildReportEmptiesARelicWithAnUnmodelledEffect(t *testing.T) {
	picks := map[string]slotPick{
		"ranged": {Item: &scored{candidate: candidate{
			ID: 227444, Name: "Idol of the Huntress", ClassID: armorClassID,
			SubclassID: armorSubidolID, Stats: map[string]float64{},
			EffectText: "Engrave your cloak with the Improved Swipe rune.",
		}}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "ranged" {
			row = s
		}
	}
	if row.ItemID != 0 || row.EmptyReason != effectNotModelledReason {
		t.Fatalf("ranged row = %+v, want empty with EmptyReason %q", row, effectNotModelledReason)
	}
	if !row.EffectUnmodelled {
		t.Error("ranged row EffectUnmodelled = false, want true (still named, even though the row is empty)")
	}
}

// An ordinary armor piece (not a relic subclass) carrying an
// unmodelled effect keeps the pre-existing exemption - only a relic's
// one-real-selling-point situation gets the new, stricter treatment.
func TestBuildReportDoesNotEmptyANonRelicWithAnUnmodelledEffect(t *testing.T) {
	picks := map[string]slotPick{
		"back": {Item: &scored{candidate: candidate{ID: 1, Name: "Mystery Cloak", EffectText: "Does something unimplemented"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "back" {
			row = s
		}
	}
	if row.ItemID != 1 || row.EmptyReason != "" {
		t.Fatalf("back row = %+v, want item 1 published (not a relic subclass, so the pre-existing EffectUnmodelled exemption still applies)", row)
	}
}

// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 3: a real
// sim DID run for this slot (verifyBand's own swap pass) even though it
// did not promote anything, but the ONE candidate it actually tested is
// not visible anywhere in Alternatives - here because it is ALSO the
// slot's own pair-mate's item (buildAlternatives correctly refuses to
// offer the same physical weapon back as a fallback for the other hand
// - shaman-enhancement's own dogfood, exactly reproduced: main_hand's
// real runner-up, Diamond Hammer, is off_hand's own pick). "Verified:
// true" must still carry real evidence somewhere on the row (three
// sweeps running: "verified" with nothing anywhere backing it up).
func TestBuildReportAddsSwapNoteWhenVerifiedButNoVisibleEvidence(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Butcher's Cleaver"}, Score: 236.9}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Diamond Hammer"}, Score: 232.8}
	offHandPick := &scored{candidate: candidate{ID: 2, Name: "Diamond Hammer"}, Score: 232.8}
	picks := map[string]slotPick{
		"main_hand": {Item: pick, RunnerUp: runnerUp},
		"off_hand":  {Item: offHandPick},
	}
	bySlot := map[string][]scored{
		"main_hand": {
			{candidate: candidate{ID: 1, Name: "Butcher's Cleaver"}, Score: 236.9},
			{candidate: candidate{ID: 3, Name: "Smite's Mighty Hammer"}, Score: 297.8},
			{candidate: candidate{ID: 4, Name: "Forsaken Greataxe"}, Score: 290.0},
			{candidate: candidate{ID: 2, Name: "Diamond Hammer"}, Score: 232.8},
		},
	}
	swaps := []swapResult{{Slot: "main_hand", SwapDPS: 38.0, BaselineDPS: 45.8, Beat: false}}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 45.8, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			row = s
		}
	}
	if !row.Verified {
		t.Fatalf("main_hand row = %+v, want Verified true (the pick beat the real swap test)", row)
	}
	for _, a := range row.Alternatives {
		if a.ItemID == 2 {
			t.Fatalf("main_hand Alternatives = %+v, want Diamond Hammer (id 2) excluded: it is off_hand's own pick, not a real main_hand fallback", row.Alternatives)
		}
	}
	if row.SwapNote == "" {
		t.Fatal("main_hand row.SwapNote is empty, want the real sim numbers named (Verified: true with no visible evidence anywhere else on the row)")
	}
	if !strings.Contains(row.SwapNote, "Diamond Hammer") || !strings.Contains(row.SwapNote, "45.8") || !strings.Contains(row.SwapNote, "38.0") {
		t.Errorf("main_hand row.SwapNote = %q, want it to name Diamond Hammer and the real 45.8/38.0 set DPS", row.SwapNote)
	}
}

// buildAlternatives itself: a real swap-tested runner-up that ranks
// below the top alternativesLimit candidates by raw score() (and is
// NOT excluded as a pair-mate) is force-included anyway, evicting the
// weakest untested filler - this lane's brief item 3's other half: the
// one candidate this command actually measured must never be silently
// dropped just because untested candidates happened to score higher.
func TestBuildAlternativesForceIncludesALowScoringButActuallyTestedRunnerUp(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Pick"}, Score: 236.9}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Weak But Tested"}, Score: 100}
	pk := slotPick{Item: pick, RunnerUp: runnerUp}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Pick"}, Score: 236.9},
		{candidate: candidate{ID: 3, Name: "Untested A"}, Score: 297.8},
		{candidate: candidate{ID: 4, Name: "Untested B"}, Score: 290.0},
		{candidate: candidate{ID: 5, Name: "Untested C"}, Score: 280.0},
		{candidate: candidate{ID: 2, Name: "Weak But Tested"}, Score: 100},
	}
	sw := &swapResult{Slot: "main_hand", SwapDPS: 38.0, BaselineDPS: 45.8, Beat: false}
	// setDPS matches sw.SwapDPS (38.0) - bis-ranker-integrity-7 lane,
	// item 2's own hygiene guard.
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 0, sw, 38.0)
	if len(got) > alternativesLimit {
		t.Fatalf("buildAlternatives = %+v, want at most %d entries", got, alternativesLimit)
	}
	var found *alternativeRow
	for i := range got {
		if got[i].ItemID == 2 {
			found = &got[i]
		}
	}
	if found == nil {
		t.Fatalf("buildAlternatives = %+v, want the tested runner-up (id 2) force-included even though it scores lowest", got)
	}
	if !found.Verified || found.SimDPS != 38.0 {
		t.Errorf("tested runner-up = %+v, want Verified true and SimDPS 38.0 (sw.SwapDPS, this row's own measured value)", found)
	}
}

// The fallback SwapNote above must stay silent when the tested
// runner-up already shows up as a Verified alternative - this lane's
// brief item 3 fixes a row with NO evidence, not double-documents one
// that already has it.
func TestBuildReportSkipsFallbackSwapNoteWhenAlternativeAlreadyShowsIt(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Pick"}, Score: 100}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Runner Up"}, Score: 95}
	picks := map[string]slotPick{"head": {Item: pick, RunnerUp: runnerUp}}
	bySlot := map[string][]scored{
		"head": {
			{candidate: candidate{ID: 1, Name: "Pick"}, Score: 100},
			{candidate: candidate{ID: 2, Name: "Runner Up"}, Score: 95},
		},
	}
	swaps := []swapResult{{Slot: "head", SwapDPS: 40, BaselineDPS: 50, Beat: false}}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 50, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0.05, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "head" {
			row = s
		}
	}
	if row.SwapNote != "" {
		t.Fatalf("head row.SwapNote = %q, want empty: the runner-up is already a Verified alternative, no fallback note needed", row.SwapNote)
	}
	found := false
	for _, a := range row.Alternatives {
		if a.ItemID == 2 && a.Verified {
			found = true
		}
	}
	if !found {
		t.Fatal("head row.Alternatives does not show the verified runner-up - test setup is wrong")
	}
}

// Owner review, tenet 8: a sim-decided row's Score is deliberately
// zeroed (buildReport's own doc), so the markdown table must never
// print a bare "0.0" beside "Verified: yes" - it reads as "this item
// does nothing" when the row actually carries a real measured number
// (hybrids sweep, item 5: Dawn's Edge/Ebon Hand/Annihilator all showed
// this). The real SimDPS is printed instead.
// This lane's brief, item 3: the fifth wow-player sweep found the
// published "Score" column mixing unlabeled bare numbers (e.g. "7.0")
// with "sim-verified (N DPS)" in the very same column, with no unit
// named anywhere on the bare rows - a player has no way to tell "7.0"
// means "7 reference-stat points" rather than some kind of DPS. The
// column header must name the band's own reference stat, and a
// score()-decided row must show the real DPS conversion beside the
// raw points number whenever this band's reference_dps_per_point is
// trustworthy (bandReport.ReferenceDPSPerPoint non-nil).
func TestWriteMarkdownLabelsTheScoreColumnAndConvertsToDPS(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "warlock-affliction.md")
	spec := specInfo{Spec: "warlock-affliction", Name: "Affliction", ReferenceStat: "shadow_power"}
	refDPSPerPoint := 0.0444
	reports := []bandReport{
		{
			Band: 50, Faction: "horde", Race: "orc",
			ReferenceDPSPerPoint: &refDPSPerPoint,
			Slots: []slotRow{
				{Slot: "head", ItemID: 1, ItemName: "Helm", Score: 10, Verified: true},
			},
		},
	}
	if err := writeMarkdown(path, spec, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading written file: %v", err)
	}
	content := string(b)
	if !strings.Contains(content, "Score (shadow_power points)") {
		t.Errorf("markdown header does not name the score column's own unit:\n%s", content)
	}
	if !strings.Contains(content, "10.0 shadow_power points (0.44 DPS)") {
		t.Errorf("markdown does not show the DPS conversion beside the raw points number:\n%s", content)
	}
	if strings.Contains(content, "| 10.0 |") {
		t.Errorf("markdown still prints a bare, unitless score number:\n%s", content)
	}
}

// The DPS conversion must not be printed when this band's own weights
// sweep was not trustworthy (ReferenceDPSPerPoint nil, weights.go's
// referenceMeasurementReason) - the raw points number still needs its
// unit named, though, so it is never bare either.
func TestWriteMarkdownLabelsTheScoreColumnWithNoDPSConversionWhenUntrustworthy(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "warlock-demonology.md")
	spec := specInfo{Spec: "warlock-demonology", Name: "Demonology", ReferenceStat: "shadow_power"}
	reports := []bandReport{
		{
			Band: 60, Faction: "alliance", Race: "human",
			WeightsReason: "reference measurement was negative",
			Slots: []slotRow{
				{Slot: "head", ItemID: 1, ItemName: "Helm", Score: 10, Verified: true},
			},
		},
	}
	if err := writeMarkdown(path, spec, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading written file: %v", err)
	}
	content := string(b)
	if !strings.Contains(content, "10.0 shadow_power points") {
		t.Errorf("markdown does not label the unlabeled points number:\n%s", content)
	}
	if strings.Contains(content, "DPS)") {
		t.Errorf("markdown printed a DPS conversion with no trustworthy reference_dps_per_point:\n%s", content)
	}
}

func TestWriteMarkdownPrintsSimVerifiedDPSNotZeroScore(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "shaman-enhancement.md")
	spec := specInfo{Spec: "shaman-enhancement", Name: "Enhancement", ReferenceStat: "attack_power"}
	reports := []bandReport{
		{
			Band: 60, Faction: "horde", Race: "orc",
			Slots: []slotRow{
				{Slot: "main_hand", ItemID: 1, ItemName: "Annihilator", Score: 0, SimDPS: 165.0, Verified: true},
			},
		},
	}
	if err := writeMarkdown(path, spec, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading written file: %v", err)
	}
	content := string(b)
	if strings.Contains(content, "| 0.0 | yes |") {
		t.Fatalf("markdown still prints a bare 0.0 score beside Verified: yes:\n%s", content)
	}
	if !strings.Contains(content, "sim-verified (165.0 DPS)") {
		t.Errorf("markdown missing the sim-verified DPS note in the Score column:\n%s", content)
	}
}

// This lane's brief (bis-ranker-integrity-6), item 6: the
// shaman-elemental band 60 Horde repro - eight unrelated slots all
// publishing the identical sim_dps (verifyBand's one shared,
// pre-promotion baseline sim), none of them equal to the band's own
// final set_dps because some OTHER slot's own swap promoted and
// applySwaps re-measured the whole set to a new total. A row publishes
// an absolute sim_dps only when it still equals the finished set's own
// set_dps (finishedSetEpsilon); every unpromoted slot sharing the
// stale baseline must omit sim_dps instead of all repeating one
// now-wrong number.
func TestBuildReportOmitsSimDPSWhenTheSharedSwapBaselineNoLongerMatchesSetDPS(t *testing.T) {
	picks := map[string]slotPick{
		"head":  {Item: &scored{candidate: candidate{ID: 1, Name: "Helm of the Elements"}, MeasuredDPS: 300.0}, RunnerUp: &scored{candidate: candidate{ID: 11, Name: "Runner Up Helm"}}},
		"chest": {Item: &scored{candidate: candidate{ID: 2, Name: "Robe of the Elements"}, MeasuredDPS: 300.0}, RunnerUp: &scored{candidate: candidate{ID: 12, Name: "Runner Up Robe"}}},
	}
	// Every slot's own swap trial shares the identical, stale
	// pre-promotion baseline (300.0) - verifyBand's own doc: one
	// baseline sim, run once for the whole band. Neither swap beat it
	// (Beat: false), but some OTHER slot elsewhere in the real band
	// promoted, so the band's own final SetDPS (below, 310.5) no
	// longer matches this shared number.
	swaps := []swapResult{
		{Slot: "head", SwapDPS: 290.0, BaselineDPS: 300.0, Beat: false},
		{Slot: "chest", SwapDPS: 288.0, BaselineDPS: 300.0, Beat: false},
	}
	r := buildReport(reportSpec(), 60, "horde", "orc", "", 0, nil, nil, picks, 310.5, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	for _, slot := range []string{"head", "chest"} {
		row := byslot[slot]
		if row.SimDPS != 0 {
			t.Errorf("%s row.SimDPS = %v, want 0 (omitted): the shared baseline 300.0 does not match set_dps 310.5", slot, row.SimDPS)
		}
		// buildAlternatives' own swap-override force-include (this
		// lane's brief item 3, an earlier lane) always surfaces the one
		// candidate verify.go actually simmed as a Verified alternative
		// with its own real DPSDelta/SimDPS, regardless of whether this
		// row's own top-level DPSDelta got set - that IS the "dps_delta
		// already on the row" this lane's brief refers to, and it must
		// survive this row's own sim_dps being withheld.
		if len(row.Alternatives) != 1 || !row.Alternatives[0].Verified {
			t.Fatalf("%s row.Alternatives = %+v, want the one verified, sim-tested runner-up still present", slot, row.Alternatives)
		}
	}
}

// The markdown table must not fall back to printing "0.0 <ref> points
// (0.00 DPS)" for a row in exactly the state
// TestBuildReportOmitsSimDPSWhenTheSharedSwapBaselineNoLongerMatchesSetDPS
// leaves behind (Score deliberately zeroed by the simDecided
// convention, SimDPS withheld for staleness) - it must show the row's
// own dps_delta instead, the same way it already shows a promoted
// swap's real numbers.
func TestWriteMarkdownShowsDPSDeltaWhenSimDPSWasWithheldForStaleness(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "shaman-elemental.md")
	spec := specInfo{Spec: "shaman-elemental", Name: "Elemental", ReferenceStat: "spell_power"}
	delta := 10.0
	reports := []bandReport{
		{
			Band: 60, Faction: "horde", Race: "orc",
			Slots: []slotRow{
				{Slot: "head", ItemID: 1, ItemName: "Helm of the Elements", Score: 0, SimDPS: 0, DPSDelta: &delta, Verified: true},
			},
		},
	}
	if err := writeMarkdown(path, spec, reports); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading written file: %v", err)
	}
	content := string(b)
	if strings.Contains(content, "0.0 spell_power points (0.00 DPS)") {
		t.Fatalf("markdown still prints the misleading bare-zero score/DPS conversion:\n%s", content)
	}
	if !strings.Contains(content, "+10.0 DPS vs the runner-up") {
		t.Errorf("markdown missing the withheld row's own dps_delta:\n%s", content)
	}
}

// bis-ranker-integrity-7 lane, item 2: the contract this whole file's
// scattered sim_dps-hygiene tests each pin one slice of - "every
// published sim_dps equals its band's own set_dps, within
// finishedSetEpsilon" - checked over one whole report covering every
// way a row can reach buildReport with a nonzero MeasuredDPS:
//   - trinket1: sim-decided (rankTrinketSlot/trySetCompletion's own
//     shape - no RunnerUp at all, so verifyBand never tries it and
//     neither swap switch case in buildReport ever sees it) whose own
//     MeasuredDPS happens to still match the finished set.
//   - trinket2: the identical shape, but stale - this is exactly the
//     path that used to publish unconditionally before this lane's fix
//     (trySetCompletion/rankTrinketSlot/rankSlotWithEffects measure a
//     slot's own DPS against whatever the OTHER slots were at that
//     moment in greedy fill, not the finished set).
//   - main_hand: a swap promotion (verify.go's applySwaps) whose own
//     single-slot SwapDPS happens to equal the finished set (the only
//     promotion this band).
//   - off_hand: a swap promotion whose own SwapDPS does NOT match -
//     the shape a second promotion in the same band produces
//     (applySwaps promotes every winning swap independently against
//     one shared pre-swap baseline, so no individual trial's own
//     SwapDPS reflects two promotions at once).
//   - head: an ordinary score()-decided pick (never simDecided at all).
func TestBuildReportContractEverySimDPSMatchesSetDPS(t *testing.T) {
	const finishedSetDPS = 500.0
	picks := map[string]slotPick{
		"head":     {Item: &scored{candidate: candidate{ID: 1, Name: "Plain Helm"}, Score: 42}},
		"trinket1": {Item: &scored{candidate: candidate{ID: 2, Name: "Fresh Trinket"}, MeasuredDPS: finishedSetDPS}},
		"trinket2": {Item: &scored{candidate: candidate{ID: 3, Name: "Stale Trinket"}, MeasuredDPS: 480.0}},
		"main_hand": {
			Item:     &scored{candidate: candidate{ID: 4, Name: "Sole Promotion"}, MeasuredDPS: finishedSetDPS},
			RunnerUp: &scored{candidate: candidate{ID: 5, Name: "Demoted"}},
		},
		"off_hand": {
			Item:     &scored{candidate: candidate{ID: 6, Name: "Co-Promoted"}, MeasuredDPS: 470.0},
			RunnerUp: &scored{candidate: candidate{ID: 7, Name: "Also Demoted"}},
		},
	}
	swaps := []swapResult{
		{Slot: "main_hand", SwapDPS: finishedSetDPS, BaselineDPS: 450.0, Beat: true},
		{Slot: "off_hand", SwapDPS: 470.0, BaselineDPS: 450.0, Beat: true},
	}
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, finishedSetDPS, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")

	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
		// The contract itself: whatever published this row's sim_dps,
		// it is only ever trustworthy when it still equals the band's
		// finished set_dps.
		if s.SimDPS != 0 && math.Abs(s.SimDPS-r.SetDPS) > finishedSetEpsilon {
			t.Errorf("slot %s: sim_dps %v does not match set_dps %v (epsilon %v) - a stale measurement was published as fact", s.Slot, s.SimDPS, r.SetDPS, finishedSetEpsilon)
		}
	}
	if byslot["trinket1"].SimDPS != finishedSetDPS {
		t.Errorf("trinket1 sim_dps = %v, want %v (a fresh, matching measurement must still publish)", byslot["trinket1"].SimDPS, finishedSetDPS)
	}
	if byslot["trinket2"].SimDPS != 0 {
		t.Errorf("trinket2 sim_dps = %v, want 0 (omitted): 480.0 is stale against set_dps %v", byslot["trinket2"].SimDPS, finishedSetDPS)
	}
	if byslot["main_hand"].SimDPS != finishedSetDPS {
		t.Errorf("main_hand sim_dps = %v, want %v (the sole promotion's own SwapDPS matches)", byslot["main_hand"].SimDPS, finishedSetDPS)
	}
	if byslot["off_hand"].SimDPS != 0 {
		t.Errorf("off_hand sim_dps = %v, want 0 (omitted): a second promotion (main_hand) in the same band left this slot's own single-swap SwapDPS (470.0) stale against set_dps %v", byslot["off_hand"].SimDPS, finishedSetDPS)
	}
	// Both withheld rows must still carry their own real, independently-
	// true evidence (SwapNote/DPSDelta) - sim_dps hygiene withholds the
	// one absolute number that no longer describes the finished set, not
	// the whole row's evidence.
	if byslot["off_hand"].SwapNote == "" {
		t.Error("off_hand SwapNote is empty, want the promotion note still present despite the withheld sim_dps")
	}
}

// TestBuildReportRefusesEvidenceForAnItemNotInSimDB is this lane's
// brief (bis-ranker-integrity-9), item 1's own contract test (item
// 2): a candidate carrying NotInSimDB (data.go's own doc - set from
// this build's real simdb.Known off the client export the effect-
// procs lane's own Hand of Justice 11815 regression names) must never
// publish Verified/SimDPS/DPSDelta, whatever MeasuredDPS or a swap
// promotion claims - simdb.Attach/AttachWeights' own UnequipUnknown
// would have silently stripped it from every character before any sim
// this command runs ever built one, so nothing here ever actually
// measured it. The row's own SimStatus names why, and the pick itself
// stays published, score-decided (Score is never zeroed to make room
// for the bogus SimDPS) - never emptied, since "the pick stays
// score-decided" (this lane's brief) is not "the pick is withheld". A
// band whose final gear includes such an item publishes
// SetDPSPartial, not a withheld set_dps (bandReport.SetDPSPartial's
// own doc explains why omitting the number outright would throw away
// a true measurement of every other slot). A second, ordinary
// score()-decided row (head) proves the flag is scoped to the one
// slot that actually carries it.
func TestBuildReportRefusesEvidenceForAnItemNotInSimDB(t *testing.T) {
	picks := map[string]slotPick{
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Known Helm"}, Score: 10}},
		"neck": {
			Item:     &scored{candidate: candidate{ID: 999999, Name: "Stripped Amulet", NotInSimDB: true}, Score: 3, MeasuredDPS: 250.0},
			RunnerUp: &scored{candidate: candidate{ID: 998, Name: "Demoted Amulet"}},
		},
	}
	swaps := []swapResult{
		{Slot: "neck", SwapDPS: 250.0, BaselineDPS: 200.0, Beat: true},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 250.0, swaps, nil, nil, 0, 0, nil, nil, nil, 0, "")

	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}

	neckRow := byslot["neck"]
	if neckRow.SimStatus != notInSimReason {
		t.Fatalf("neck row SimStatus = %q, want %q", neckRow.SimStatus, notInSimReason)
	}
	if neckRow.Verified {
		t.Error("neck row Verified = true, want false: this item's own sim never actually ran")
	}
	if neckRow.SimDPS != 0 {
		t.Errorf("neck row SimDPS = %v, want 0: MeasuredDPS must never publish for an item simdb.Known rejects", neckRow.SimDPS)
	}
	if neckRow.DPSDelta != nil {
		t.Errorf("neck row DPSDelta = %v, want nil", *neckRow.DPSDelta)
	}
	if neckRow.ItemID != 999999 || neckRow.Score != 3 {
		t.Errorf("neck row = %+v, want the pick published score-decided (item_id 999999, score 3 - never zeroed, never emptied)", neckRow)
	}
	if !r.SetDPSPartial {
		t.Error("SetDPSPartial = false, want true: the band's final gear includes an item not in this build's simdb")
	}

	headRow := byslot["head"]
	if headRow.SimStatus != "" || !headRow.Verified {
		t.Errorf("head row = %+v, want an ordinary verified score-decided row, untouched by neck's own flag", headRow)
	}
}

// TestBuildReportAlternativesAreUniqueByItemID is a whole-report
// contract test - controller direction, 2026-09-30 (lane bis-ranker-
// integrity-9's own addition, priest-shadow band 60 Alliance
// main_hand's own repro on nightly 2cd94a6b): Grand Marshal's Stave
// reached this slot's own candidate pool as two rows sharing the
// identical item id, one Kind "pvp" (loot.json's own
// "pvp:rank-18:alliance" source) and one Kind "vendor"
// (vendorInheritsPvpRankGate, data.go's own doc - the quartermaster
// row that inherited that same rank) - published twice with an
// identical dps_delta of 0 before this fix. No slot in a whole report
// may ever list the same item id twice in its own Alternatives, and
// the surviving row must be the "pvp" one (betterAlternative's own
// tie-break, report.go).
func TestBuildReportAlternativesAreUniqueByItemID(t *testing.T) {
	pvpVersion := scored{candidate: candidate{ID: 500, Name: "Grand Marshal's Stave"}, Score: 90, HasSource: true, Source: itemSource{Kind: "pvp", Label: "Alliance PvP Rank 18"}}
	vendorVersion := scored{candidate: candidate{ID: 500, Name: "Grand Marshal's Stave"}, Score: 90, HasSource: true, Source: itemSource{Kind: "vendor", Label: "Quartermaster"}}
	picks := map[string]slotPick{
		"main_hand": {
			Item: &scored{candidate: candidate{ID: 1, Name: "Butcher's Cleaver"}, Score: 200},
			// vendorVersion listed FIRST deliberately: a tie-break that
			// merely kept "whichever was seen first" (the old
			// seen-map-only guard) would publish the vendor row here,
			// the wrong survivor per the controller's own rule -
			// betterAlternative (report.go) must actively prefer the
			// "pvp" row regardless of which order the two reached this
			// list.
			Ties: []scored{vendorVersion, pvpVersion},
		},
	}
	bySlot := map[string][]scored{
		"main_hand": {*picks["main_hand"].Item, vendorVersion, pvpVersion},
	}
	r := buildReport(reportSpec(), 60, "alliance", "human", "", 0, nil, nil, picks, 500, nil, nil, nil, 0, 0, nil, nil, bySlot, 0.05, "")

	for _, s := range r.Slots {
		seen := map[int]bool{}
		for _, a := range s.Alternatives {
			if seen[a.ItemID] {
				t.Fatalf("slot %s lists item id %d twice in its own alternatives: %+v", s.Slot, a.ItemID, s.Alternatives)
			}
			seen[a.ItemID] = true
		}
	}

	var mainHand slotRow
	for _, s := range r.Slots {
		if s.Slot == "main_hand" {
			mainHand = s
		}
	}
	var found bool
	for _, a := range mainHand.Alternatives {
		if a.ItemID != 500 {
			continue
		}
		found = true
		if a.SourceKind != "pvp" {
			t.Errorf("Grand Marshal's Stave alternative source_kind = %q, want %q (the pvp row wins the tie)", a.SourceKind, "pvp")
		}
	}
	if !found {
		t.Fatal("Grand Marshal's Stave (id 500) missing from main_hand's alternatives entirely")
	}
}

// TestBuildReportNeverPublishesAPositiveDPSDeltaOnAnAlternative is a
// whole-report contract test - controller direction, 2026-09-30
// (ninth wow-player sweep, day3/player-review-24/casters.md finding
// 1): mage-arcane band 60's own Weakness Analyzer published dps_delta
// +2.34 next to the KEPT pick Talisman of Ascendance (verified: true)
// - a real, well-above-noise measured gain (verify.go's own swap sim)
// that never cleared swapMargin, so the pick was kept on purpose, yet
// the alternative row still showed a positive number a player reads as
// "the site chose the worse item". Rule (swapMeasuredDelta, report.go):
// a runner-up that measured higher but stayed inside the margin
// publishes dps_delta 0 (a tie), Verified true; a runner-up that
// measured higher BY MORE than the margin is never left as an
// alternative at all - applySwaps (verify.go) already promotes it into
// the pick before buildReport ever runs. No alternative in a whole
// report may ever carry a positive dps_delta.
func TestBuildReportNeverPublishesAPositiveDPSDeltaOnAnAlternative(t *testing.T) {
	const baselineDPS = 300.0
	// +2.34 DPS: real and well above swapDeltaNoiseFloorDPS (0.05), but
	// under swapMargin's own 1% bar (baselineDPS*1.01 = 303.0), so
	// beatsByMargin is false and verify.go kept the pick - the exact
	// shape of the caster sweep's own repro.
	const swapDPS = 302.34
	picks := map[string]slotPick{
		"trinket2": {
			Item:     &scored{candidate: candidate{ID: 1, Name: "Talisman of Ascendance"}, Score: 50},
			RunnerUp: &scored{candidate: candidate{ID: 2, Name: "Weakness Analyzer"}, Score: 10},
		},
	}
	swaps := []swapResult{
		{Slot: "trinket2", SwapDPS: swapDPS, BaselineDPS: baselineDPS, Beat: false},
	}
	r := buildReport(reportSpec(), 60, "alliance", "human", "", 0, nil, nil, picks, baselineDPS, swaps, nil, nil, 0, 0, nil, nil, nil, 0.05, "")

	var found bool
	for _, s := range r.Slots {
		for _, a := range s.Alternatives {
			if a.DPSDelta > 0 {
				t.Errorf("slot %s alternative %q (id %d) published dps_delta %v, want <= 0: a runner-up that did not clear swapMargin is a tie, never a positive number", s.Slot, a.ItemName, a.ItemID, a.DPSDelta)
			}
			if a.ItemID == 2 {
				found = true
				if a.DPSDelta != 0 {
					t.Errorf("Weakness Analyzer's own alternative row dps_delta = %v, want 0 (a tie: it measured higher but never cleared swapMargin)", a.DPSDelta)
				}
				if !a.Verified {
					t.Error("Weakness Analyzer's own alternative row Verified = false, want true: a real sim did compare it to the pick")
				}
			}
		}
	}
	if !found {
		t.Fatal("Weakness Analyzer (id 2) missing from trinket2's own alternatives entirely")
	}
}

// TestBuildReportEmptiesATrinketWhoseGainDoesNotClearTwiceItsError is
// bis-ranker-integrity-11's own repro at the buildReport level: a
// trinket whose MeasuredGainDPS is comfortably positive (well above
// the old flat trinketZeroGainThresholdDPS floor) but does not clear
// trinketGainSignificanceMultiplier times its own MeasuredGainStdErr
// (Fire Ruby's own hunter-beast-mastery band 50 Alliance numbers -
// gain 0.773, stdErr 0.7125) must still publish empty, not verified -
// the exact case a bare positive-DPS-gain check let through before
// this fix.
func TestBuildReportEmptiesATrinketWhoseGainDoesNotClearTwiceItsError(t *testing.T) {
	picks := map[string]slotPick{
		"trinket1": {Item: &scored{candidate: candidate{ID: 20036, Name: "Fire Ruby"}, MeasuredDPS: 233.6139, MeasuredGainDPS: 0.773, MeasuredGainStdErr: 0.7125, GainMeasured: true}},
	}
	r := buildReport(reportSpec(), 50, "alliance", "dwarf", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket1" {
			row = s
		}
	}
	if row.ItemID != 0 || row.EmptyReason != noDPSValueReason {
		t.Fatalf("trinket1 row = %+v, want empty with EmptyReason %q (gain 0.773 does not clear 2x its own stdErr 0.7125)", row, noDPSValueReason)
	}
}

// TestBuildReportKeepsATrinketWhoseGainClearsTwiceItsError is the
// symmetric real-signal case: Frozen Heart of the Mountain's own
// rogue-assassination band 50 Horde numbers (gain 3.03, stdErr 0.98)
// comfortably clear 2x and must still publish, exactly as this lane's
// brief requires ("Frozen Heart of the Mountain at band 50... stays
// legitimate").
func TestBuildReportKeepsATrinketWhoseGainClearsTwiceItsError(t *testing.T) {
	picks := map[string]slotPick{
		"trinket2": {Item: &scored{candidate: candidate{ID: 249469, Name: "Frozen Heart of the Mountain"}, MeasuredDPS: 123.3857, MeasuredGainDPS: 3.03, MeasuredGainStdErr: 0.98, GainMeasured: true}},
	}
	r := buildReport(reportSpec(), 50, "horde", "troll", "", 0, nil, nil, picks, 123.3857, nil, nil, nil, 0, 0, nil, nil, nil, 0, "")
	var row slotRow
	for _, s := range r.Slots {
		if s.Slot == "trinket2" {
			row = s
		}
	}
	if row.ItemID != 249469 || row.EmptyReason != "" {
		t.Fatalf("trinket2 row = %+v, want item 249469 published with no EmptyReason (gain 3.03 clears 2x its own stdErr 0.98)", row)
	}
}
