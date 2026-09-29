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
	r := buildReport(reportSpec(), 20, "horde", "troll", "0500000", 5, map[string]api.StatWeight{"agility": {Stat: "agility", Weight: 1.5, Error: 0.1}}, reportSpec().WeightStats, picks, 500, nil, nil, nil, 1.2, 3.4, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, previous, 0, 0, nil)
	if len(r.NewAtBand) != 0 {
		t.Fatalf("NewAtBand = %v, want empty: item 1 unchanged from the previous band", r.NewAtBand)
	}
}

func TestBuildReportChangedFromPreviousBandIsNew(t *testing.T) {
	picks := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 2, Name: "Better Helm"}}}}
	previous := map[string]slotPick{"head": {Item: &scored{candidate: candidate{ID: 1, Name: "Helm"}}}}
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, previous, 0, 0, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 200, swaps, nil, nil, 0, 0, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 150, swaps, nil, nil, 0, 0, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 150, nil, nil, nil, 0, 0, verifyErrors)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, map[string]slotPick{}, 0, nil, noSource, nil, 0, 0, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil)
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
	r := buildReport(reportSpec(), 20, "horde", "troll", "", 0, weights, order, map[string]slotPick{}, 0, nil, nil, nil, 0, 0, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil)
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
	r := buildReport(reportSpec(), 30, "horde", "troll", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil)
	byslot := map[string]slotRow{}
	for _, s := range r.Slots {
		byslot[s.Slot] = s
	}
	if !byslot["ranged"].EffectUnmodelled {
		t.Error("ranged (relic with an unimplemented effect) EffectUnmodelled = false, want true")
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
				{Slot: "head", ItemID: 1, ItemName: "Helm", Score: 10, Verified: true, Source: "A Quest", SourceKind: "quest"},
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
}
