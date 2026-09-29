package main

import (
	"encoding/json"
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
		"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Helm"}, HasSource: true, Source: itemSource{Kind: "quest", Label: "A Quest"}}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "0500000", 5, map[string]api.StatWeight{"agility": {Stat: "agility", Weight: 1.5, Error: 0.1}}, reportSpec().WeightStats, picks, 500, nil, nil, nil, 1.2, 3.4, nil, nil, nil, 0)
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
	item1 := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	picks := map[string]slotPick{"head": {Item: item1}}
	previous := map[string]slotPick{"head": {Item: item1}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, previous, 0, 0, nil, nil, nil, 0)
	if len(r.NewAtBand) != 0 {
		t.Fatalf("NewAtBand = %v, want empty: item 1 unchanged from the previous band", r.NewAtBand)
	}
}

func TestBuildReportChangedFromPreviousBandIsNew(t *testing.T) {
	picks := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}}}
	previous := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Helm"}}}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, previous, 0, 0, nil, nil, nil, 0)
	if len(r.NewAtBand) != 1 || !strings.Contains(r.NewAtBand[0], "Better Helm") {
		t.Fatalf("NewAtBand = %v, want Better Helm listed", r.NewAtBand)
	}
}

func TestBuildReportSwapBeatenRowIsTheWinnerVerifiedWithNote(t *testing.T) {
	// applySwaps has already promoted the runner-up into Item and demoted
	// the scored pick to RunnerUp before buildReport sees the picks.
	winner := &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}
	beaten := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	picks := map[string]slotPick{"head": {Item: winner, RunnerUp: beaten}}
	swaps := []swapResult{{Slot: "head", SwapDPS: 200, BaselineDPS: 150, Beat: true}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 200, swaps, nil, nil, 0, 0, nil, nil, nil, 0)
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

func TestBuildReportSwapLostKeepsSlotVerified(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Worse Helm"}}
	picks := map[string]slotPick{"head": {Item: pick, RunnerUp: runnerUp}}
	swaps := []swapResult{{Slot: "head", SwapDPS: 50, Beat: false}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 150, swaps, nil, nil, 0, 0, nil, nil, nil, 0)
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
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Other Helm"}}
	picks := map[string]slotPick{"head": {Item: pick, RunnerUp: runnerUp}}
	verifyErrors := []string{"head: runner-up Other Helm (id 2): the engine reported an error: boom"}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 150, nil, nil, nil, 0, 0, verifyErrors, nil, nil, 0)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, noSource, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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

// This lane's brief, item 1 (owner defect A, 2026-09-29): a slot's
// alternatives must carry the next-best sourced candidates after the
// pick, so a player who cannot get the pick's own source has a real
// fallback instead of an empty page - warrior-arms horde band 20's
// main_hand published Forsaken Greataxe with no record that Smite's
// Mighty Hammer (a Deadmines drop) was ever considered.
func TestBuildAlternativesRanksTiesFirstThenNextBestSourcedByScore(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Forsaken Greataxe"}, Score: 302.9}
	pk := slotPick{
		Item: pick,
		// A tie (this lane's brief addendum, wow-player review):
		// scores identically to the pick, so it must rank ahead of
		// every lower-scoring alternative even though bySlot's own
		// score order does not itself distinguish the two.
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
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk})
	want := []alternativeRow{
		{ItemID: 5, ItemName: "Tied Greataxe", Score: 302.9, SourceKind: "quest", Source: "Quests", DPSDelta: 0},
		{ItemID: 6, ItemName: "Hammerbone", Score: 306.05, SourceKind: "quest", Source: "Quests", DPSDelta: 306.05 - 302.9},
		{ItemID: 7230, ItemName: "Smite's Mighty Hammer", Score: 297.82, SourceKind: "dungeon", Source: "The Deadmines: Mr. Smite", DPSDelta: 297.82 - 302.9},
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
	got := buildAlternatives(picks["finger1"], "finger1", list, picks)
	if len(got) != 1 || got[0].ItemID != 3 {
		t.Fatalf("finger1 alternatives = %+v, want only Ring C (finger2's own pick excluded)", got)
	}
}

// A slot with no pick (an empty row) and a two-handed main_hand's
// off_hand (which pick.go's enforceTwoHandOffHandInvariant always
// leaves nil) both carry no alternatives - there is no pick to offer
// fallbacks for.
func TestBuildAlternativesNilForAnUnfilledSlot(t *testing.T) {
	list := []scored{{candidate: candidate{ID: 1, Name: "Anything"}, Score: 5}}
	got := buildAlternatives(slotPick{}, "off_hand", list, map[string]slotPick{})
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
	got := buildAlternatives(slotPick{Item: pick}, "main_hand", list, map[string]slotPick{"main_hand": {Item: pick}})
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, bySlot, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0.0714)
	if r.ReferenceDPSPerPoint != 0.0714 {
		t.Fatalf("ReferenceDPSPerPoint = %v, want 0.0714", r.ReferenceDPSPerPoint)
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
					Ties:         []tieAlternative{{ItemID: 5, ItemName: "Tied Cap"}},
					Alternatives: []alternativeRow{{ItemID: 6, ItemName: "Runner Up Helm", Score: 8, SourceKind: "dungeon", Source: "Some Dungeon", DPSDelta: -2}},
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
	if !strings.Contains(content, "Runner Up Helm (6, -2.0) [dungeon]") {
		t.Error("markdown missing the head slot's own Alternatives column entry (this lane's brief, item 1)")
	}
}
