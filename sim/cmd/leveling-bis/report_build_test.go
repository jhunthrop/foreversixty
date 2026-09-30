package main

import (
	"encoding/json"
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
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10}
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
	pick := &scored{candidate: candidate{ID: 1, Name: "Helm"}, Score: 10}
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

// This lane's brief, item 2 (bis-ranker-integrity, 2026-09-29):
// Sentinel's Medallion (Agility/Stamina) scoring exactly 0 for a caster
// spec must publish empty with empty_reason "no_dps_value", not as a
// "verified" pick - the exact defect every caster spec's own band-20
// neck carried before this fix.
func TestBuildReportPublishesEmptySlotWithNoDPSValueReason(t *testing.T) {
	picks := map[string]slotPick{
		"neck": {Item: &scored{candidate: candidate{ID: 1, Name: "Sentinel's Medallion"}, Score: 0}},
	}
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 40, "alliance", "human", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
	for _, row := range r.Slots {
		if row.ItemID == 0 && row.EmptyReason == "" {
			t.Errorf("slot %s: no item and no EmptyReason - every empty slot must carry a reason", row.Slot)
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
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 1.0, nil)
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
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, referenceDPSPerPoint, nil)
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

	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 0.0444, sw)

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
	got := buildAlternatives(pk, "head", list, map[string]slotPick{"head": pk}, 0.05, sw)
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

// This lane's brief, item 9: swapMargin (verify.go) deliberately
// keeps the scored pick when a runner-up measures higher but not
// enough to clear the noise margin (Beat == false), but the runner-up
// still genuinely measured a few tenths of a DPS higher - the exact
// shape that published a POSITIVE, "verified: true" dps_delta before
// this fix (55 slots across ret/feral/enhancement, the controller's
// own review). The contract: every alternative's dps_delta <= 0 after
// the swap stage, verified or not.
func TestSwapMeasuredDeltaCapsAPositiveNotBeatDeltaAtZero(t *testing.T) {
	// SwapDPS (100.3) is higher than BaselineDPS (100), but not by
	// enough to clear swapMargin (1%) - Beat is correctly false, but
	// the naive SwapDPS-BaselineDPS would be +0.3.
	sw := swapResult{SwapDPS: 100.3, BaselineDPS: 100, Beat: false}
	got := swapMeasuredDelta(sw)
	if got > 0 {
		t.Fatalf("swapMeasuredDelta(%+v) = %v, want <= 0 (a candidate the swap stage did not promote must never publish a positive delta)", sw, got)
	}
	if got != 0 {
		t.Errorf("swapMeasuredDelta(%+v) = %v, want exactly 0 (capped, not the naive +0.3)", sw, got)
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

// End-to-end pin through buildAlternatives: the exact "flips sign
// between factions" shape the controller's review named (a coin-flip
// margin either side of the 1% bar) must publish dps_delta <= 0
// either way, never a positive "verified" win for the side that
// happened to land just inside the margin's other side.
func TestBuildAlternativesNeverPublishesAPositiveVerifiedDeltaWhenNotBeat(t *testing.T) {
	pick := &scored{candidate: candidate{ID: 1, Name: "Knight's Leather Pants"}, Score: 100}
	runnerUp := &scored{candidate: candidate{ID: 2, Name: "Stormshroud Pants"}, Score: 105}
	pk := slotPick{Item: pick, RunnerUp: runnerUp}
	list := []scored{
		{candidate: candidate{ID: 1, Name: "Knight's Leather Pants"}, Score: 100},
		{candidate: candidate{ID: 2, Name: "Stormshroud Pants"}, Score: 105, Source: itemSource{Kind: "crafted", Label: "Tailoring"}},
	}
	sw := &swapResult{Slot: "legs", SwapDPS: 100.4, BaselineDPS: 100, Beat: false}
	got := buildAlternatives(pk, "legs", list, map[string]slotPick{"legs": pk}, 0.05, sw)
	if len(got) != 1 || got[0].ItemID != 2 {
		t.Fatalf("buildAlternatives = %+v, want exactly the runner-up", got)
	}
	if got[0].DPSDelta > 0 {
		t.Errorf("Stormshroud Pants DPSDelta = %v, want <= 0 (Beat was false; the sim did not promote it)", got[0].DPSDelta)
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
	got := buildAlternatives(picks["finger1"], "finger1", list, picks, 1.0, nil)
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
	got := buildAlternatives(slotPick{}, "off_hand", list, map[string]slotPick{}, 1.0, nil)
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
	got := buildAlternatives(slotPick{Item: pick}, "main_hand", list, map[string]slotPick{"main_hand": {Item: pick}}, 1.0, nil)
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

// This lane's brief, item 7: bandReport always documents its own
// Score field's unit, regardless of what the band's picks look like.
func TestBuildReportAlwaysPublishesScoreUnit(t *testing.T) {
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, swaps, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 60, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, nil, 0)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 45.8, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0)
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
	got := buildAlternatives(pk, "main_hand", list, map[string]slotPick{"main_hand": pk}, 0, sw)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, nil, nil, picks, 50, swaps, nil, nil, 0, 0, nil, nil, bySlot, 0.05)
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
