package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestFormatWeights(t *testing.T) {
	got := formatWeights([]string{"agility", "crit"}, map[string]api.StatWeight{
		"agility": {Stat: "agility", Weight: 1.5, Error: 0.1},
		"crit":    {Stat: "crit", Weight: 0.75, Error: 0.05},
	})
	want := "agility=1.500 ± 0.100, crit=0.750 ± 0.050"
	if got != want {
		t.Errorf("formatWeights = %q, want %q", got, want)
	}
}

func TestFormatWeightsMissingStatDefaultsToZeroAndReadsNotSignificant(t *testing.T) {
	got := formatWeights([]string{"agility", "hit"}, map[string]api.StatWeight{
		"agility": {Stat: "agility", Weight: 1, Error: 0.1},
	})
	want := "agility=1.000 ± 0.100, hit=not significant (0.000 ± 0.000)"
	if got != want {
		t.Errorf("formatWeights = %q, want %q", got, want)
	}
}

func TestFormatWeightsFlagsHighErrorAsNotSignificant(t *testing.T) {
	got := formatWeights([]string{"melee_haste"}, map[string]api.StatWeight{
		"melee_haste": {Stat: "melee_haste", Weight: 14.87, Error: 6.0},
	})
	want := "melee_haste=not significant (14.870 ± 6.000)"
	if got != want {
		t.Errorf("formatWeights = %q, want %q", got, want)
	}
}

func TestTalentPointsSpent(t *testing.T) {
	cases := map[string]int{
		"0500000":  5,
		"5-3-0":    8,
		"":         0,
		"000":      0,
		"10-20-30": 6, // digits summed independently: 1+0+2+0+3+0
	}
	for talents, want := range cases {
		if got := talentPointsSpent(talents); got != want {
			t.Errorf("talentPointsSpent(%q) = %d, want %d", talents, got, want)
		}
	}
}

func TestParseBands(t *testing.T) {
	got, err := parseBands("10,20, 30 ,40")
	if err != nil {
		t.Fatalf("parseBands: %v", err)
	}
	want := []int{10, 20, 30, 40}
	if len(got) != len(want) {
		t.Fatalf("parseBands = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("parseBands = %v, want %v", got, want)
		}
	}
}

func TestParseBandsInvalidNumber(t *testing.T) {
	if _, err := parseBands("10,abc,30"); err == nil {
		t.Fatal("parseBands(10,abc,30): want an error, got nil")
	}
}

func TestParseBandsDefaultFlag(t *testing.T) {
	got, err := parseBands(defaultBandsFlag)
	if err != nil {
		t.Fatalf("parseBands(defaultBandsFlag): %v", err)
	}
	if len(got) != 5 || got[0] != 20 || got[len(got)-1] != 60 {
		t.Fatalf("parseBands(defaultBandsFlag) = %v, want 5 bands from 20 to 60", got)
	}
}

func TestReadActiveBuild(t *testing.T) {
	build, err := readActiveBuild(repoRootFixture)
	if err != nil {
		t.Fatalf("readActiveBuild: %v", err)
	}
	if build != "testbuild" {
		t.Errorf("build = %q, want testbuild", build)
	}
}

func TestReadActiveBuildMissingFile(t *testing.T) {
	if _, err := readActiveBuild(t.TempDir()); err == nil {
		t.Fatal("readActiveBuild on an empty dir: want an error, got nil")
	}
}

func TestReadActiveBuildEmptyBuildField(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "web", "src", "data", "active-build.json"), `{"build":""}`); err != nil {
		t.Fatal(err)
	}
	if _, err := readActiveBuild(dir); err == nil {
		t.Fatal("readActiveBuild with an empty build field: want an error, got nil")
	}
}

func TestReadActiveBuildMalformedJSON(t *testing.T) {
	dir := t.TempDir()
	if err := writeFile(t, filepath.Join(dir, "web", "src", "data", "active-build.json"), `{not json`); err != nil {
		t.Fatal(err)
	}
	if _, err := readActiveBuild(dir); err == nil {
		t.Fatal("readActiveBuild with malformed JSON: want an error, got nil")
	}
}

func TestWriteHeapProfileWritesAFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "heap.prof")
	writeHeapProfile(path)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}
	if info.Size() == 0 {
		t.Error("heap profile file is empty")
	}
}

func TestWriteHeapProfileBadPathLogsRatherThanPanics(t *testing.T) {
	// A path under a directory that cannot exist (a file, not a
	// directory, in the middle of it) makes os.Create fail - this must
	// be logged, not fatal (writeHeapProfile's own doc: "a failure to
	// write one is logged, not fatal"). The test's assertion is simply
	// that this does not panic.
	dir := t.TempDir()
	blocker := filepath.Join(dir, "not-a-directory")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeHeapProfile(filepath.Join(blocker, "heap.prof"))
}

func TestRunSpecWritesReportsForEveryBandAndFaction(t *testing.T) {
	outDir := t.TempDir()
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
	}
	buildDir := buildDirFixture()
	err := runSpec(fake, repoRootFixture, buildDir, "testbuild", outDir, "hunter-marksmanship", []int{20, 30}, 5)
	if err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	jsonPath := filepath.Join(outDir, "hunter-marksmanship.json")
	if _, err := os.Stat(jsonPath); err != nil {
		t.Fatalf("stat %s: %v", jsonPath, err)
	}
	mdPath := filepath.Join(outDir, "hunter-marksmanship.md")
	b, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("reading %s: %v", mdPath, err)
	}
	content := string(b)
	if !strings.Contains(content, "## Alliance") || !strings.Contains(content, "## Horde") {
		t.Errorf("markdown missing both faction sections:\n%s", content)
	}
	if !strings.Contains(content, "Band 20") || !strings.Contains(content, "Band 30") {
		t.Errorf("markdown missing both bands:\n%s", content)
	}
}

func TestRunSpecUnknownSpecErrors(t *testing.T) {
	fake := &fakeEngine{}
	err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", t.TempDir(), "no-such-spec", []int{20}, 5)
	if err == nil {
		t.Fatal("runSpec(no-such-spec): want an error, got nil")
	}
}

func TestRunSpecWeightsFailurePropagates(t *testing.T) {
	fake := &fakeEngine{FailWeights: true}
	err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", t.TempDir(), "hunter-marksmanship", []int{20}, 5)
	if err == nil {
		t.Fatal("runSpec with a failing weights run: want an error, got nil")
	}
}

// digitSum totals a LadderTalentString's digits ("5-5-1" -> 11) - a
// cheap, deterministic stand-in for "how many talent points this
// string spends", used below to make a fake engine's DPS a function of
// Character.Talents without needing the real engine.
func digitSum(talents string) int {
	sum := 0
	for _, r := range talents {
		if r >= '0' && r <= '9' {
			sum += int(r - '0')
		}
	}
	return sum
}

// readSpecReportForTest reads a data/builds/<build>/bis/<spec>.json
// file back into a specReport - the same shape writeSpecReport writes
// (report.go) - for a test to assert on published fields runSpec's own
// return value does not expose.
func readSpecReportForTest(t *testing.T, path string) specReport {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	var out specReport
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("decoding %s: %v", path, err)
	}
	return out
}

// TestRunSpecPublishesDifferentSetDPSForDifferentTalentStrings guards
// this lane's brief, defect 1: a direct sim of the published BM
// level-60 set gave 227 DPS with talents equipped and 201 without, yet
// the published set_dps was 200.15 for BOTH BM and MM at EVERY band -
// the tell that Character.Talents never reached the engine at all
// (character.go's bandCharacter, used by every plain-DPS site, is the
// fix). This drives two specs whose fixture guide builds spend
// different points (marksmanship.md: 5/5/5, beast-mastery.md: 1/1/1)
// through the real pipeline with a fake engine whose DPS is a function
// of Character.Talents, not of gear, and asserts their published
// set_dps differ - the same shape a real regression here would show.
func TestRunSpecPublishesDifferentSetDPSForDifferentTalentStrings(t *testing.T) {
	fake := &fakeEngine{
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
		DPSFunc: func(req api.SimRequest) (float64, error) {
			return 200 + float64(digitSum(req.Character.Talents))*2, nil
		},
	}

	mmDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", mmDir, "hunter-marksmanship", []int{20}, 5); err != nil {
		t.Fatalf("runSpec(hunter-marksmanship): %v", err)
	}
	bmDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", bmDir, "hunter-beast-mastery", []int{20}, 5); err != nil {
		t.Fatalf("runSpec(hunter-beast-mastery): %v", err)
	}

	mm := readSpecReportForTest(t, filepath.Join(mmDir, "hunter-marksmanship.json"))
	bm := readSpecReportForTest(t, filepath.Join(bmDir, "hunter-beast-mastery.json"))
	if len(mm.Bands) == 0 || len(bm.Bands) == 0 {
		t.Fatalf("expected at least one band report each, got mm=%d bm=%d", len(mm.Bands), len(bm.Bands))
	}
	if mm.Bands[0].Talents == bm.Bands[0].Talents {
		t.Fatalf("fixture guide builds must spend different points so this test proves something: both bands got talents %q", mm.Bands[0].Talents)
	}
	if mm.Bands[0].SetDPS == bm.Bands[0].SetDPS {
		t.Fatalf("hunter-marksmanship and hunter-beast-mastery talents differ (%q vs %q) but published set_dps is identical (%.2f) - Character.Talents is not reaching the engine", mm.Bands[0].Talents, bm.Bands[0].Talents, mm.Bands[0].SetDPS)
	}
}

// TestRunSpecEveryPlainDPSRequestCarriesTheBandsTalents is the
// narrower, per-call form of the guard above: every RunPlainDPS
// request runSpec issues (verifyBand's baseline/swaps, applySwaps'
// re-measure, rankTrinketSlot, rankSlotWithEffects, trySetCompletion -
// this lane's brief names all eight CharacterSpec construction sites)
// must carry the band's own non-empty talent string, not the zero
// value.
func TestRunSpecEveryPlainDPSRequestCarriesTheBandsTalents(t *testing.T) {
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
	}
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", t.TempDir(), "hunter-marksmanship", []int{20}, 5); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	if len(fake.TalentsSeen) == 0 {
		t.Fatal("no RunPlainDPS calls were recorded - this test proves nothing")
	}
	for i, talents := range fake.TalentsSeen {
		if talents == "" {
			t.Fatalf("RunPlainDPS call %d (gear %s) carried an empty Character.Talents", i, fake.Calls[i])
		}
	}
}

// TestRunSpecRetriesWeightsSweepAndUsesTheResolvedResult is this
// lane's brief, item 1's guard: a band whose reference stat measured
// "not positive beyond its own error" at the ordinary iteration count
// re-runs the sweep once at weightsRetryIterationsFactor iterations
// before giving up on it. Band 20 here always measures significant
// (one call); band 30's own fake result depends on req.Iterations - the
// ordinary run (5, this test's own weightsIterations) is not
// trustworthy, the retry run (5*weightsRetryIterationsFactor) is - so
// band 30 must publish an EMPTY WeightsReason and its own
// ReferenceDPSPerPoint, not band 20's carried-forward weights.
func TestRunSpecRetriesWeightsSweepAndUsesTheResolvedResult(t *testing.T) {
	const weightsIterations = 5
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsFunc: func(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
			significant := map[string]api.StatWeight{
				"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0, Error: 0.01},
				"agility":             {Stat: "agility", Weight: 1.8, Error: 0.05},
			}
			if req.Character.Level == 20 {
				return significant, 1.0, nil
			}
			// Band 30: the ordinary iteration count reads noisy (a
			// small reference delta well within its own error); the
			// retry count reads clean.
			if req.Iterations == weightsIterations {
				return map[string]api.StatWeight{
					"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1, Error: -5},
					"agility":             {Stat: "agility", Weight: 1.8, Error: 0.05},
				}, -0.02, nil
			}
			return significant, 1.0, nil
		},
	}
	outDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", outDir, "hunter-marksmanship", []int{20, 30}, weightsIterations); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	report := readSpecReportForTest(t, filepath.Join(outDir, "hunter-marksmanship.json"))
	var band30 *bandReport
	for i := range report.Bands {
		if report.Bands[i].Band == 30 {
			band30 = &report.Bands[i]
			break
		}
	}
	if band30 == nil {
		t.Fatal("no band-30 report published")
	}
	if band30.WeightsReason != "" {
		t.Errorf("band 30 weights_reason = %q, want \"\" (the retry resolved it)", band30.WeightsReason)
	}
	if band30.ReferenceDPSPerPoint == nil || *band30.ReferenceDPSPerPoint != 1.0 {
		t.Errorf("band 30 reference_dps_per_point = %v, want 1.0 (the retry's own measurement)", band30.ReferenceDPSPerPoint)
	}
	// Band 20 (significant first try) issues exactly one RunWeights
	// call; band 30 (insignificant first try) issues two, the second
	// at weightsRetryIterationsFactor times the ordinary count.
	want := []int{weightsIterations, weightsIterations, weightsIterations * weightsRetryIterationsFactor}
	if len(fake.WeightsIterationsSeen) != len(want) {
		t.Fatalf("WeightsIterationsSeen = %v, want %v", fake.WeightsIterationsSeen, want)
	}
	for i, w := range want {
		if fake.WeightsIterationsSeen[i] != w {
			t.Errorf("WeightsIterationsSeen[%d] = %d, want %d (%v)", i, fake.WeightsIterationsSeen[i], w, fake.WeightsIterationsSeen)
		}
	}
}

// TestRunSpecFallsBackToNearestLowerBandWeightsWhenRetryStillFails is
// this lane's brief, item 1's fallback half: when even the
// weightsRetryIterationsFactor-iteration retry still measures the
// reference stat "not positive beyond its own error", the band ranks
// and verifies its picks against the nearest LOWER band's own
// significant weights instead of publishing every slot empty
// (warlock-destruction band 60's own repro, this lane's report) - and
// says exactly that in weights_reason.
func TestRunSpecFallsBackToNearestLowerBandWeightsWhenRetryStillFails(t *testing.T) {
	const weightsIterations = 5
	significant := map[string]api.StatWeight{
		"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0, Error: 0.01},
		"agility":             {Stat: "agility", Weight: 1.8, Error: 0.05},
	}
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsFunc: func(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
			if req.Character.Level == 20 {
				return significant, 1.0, nil
			}
			// Band 30 never resolves, at either iteration count.
			return map[string]api.StatWeight{
				"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1, Error: -5},
				"agility":             {Stat: "agility", Weight: 1.8, Error: 0.05},
			}, -0.02, nil
		},
	}
	outDir := t.TempDir()
	if err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", outDir, "hunter-marksmanship", []int{20, 30}, weightsIterations); err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	report := readSpecReportForTest(t, filepath.Join(outDir, "hunter-marksmanship.json"))
	var band30 *bandReport
	for i := range report.Bands {
		if report.Bands[i].Band == 30 {
			band30 = &report.Bands[i]
			break
		}
	}
	if band30 == nil {
		t.Fatal("no band-30 report published")
	}
	if !strings.Contains(band30.WeightsReason, "weights carried from band 20") {
		t.Errorf("band 30 weights_reason = %q, want it to say weights were carried from band 20", band30.WeightsReason)
	}
	if !strings.Contains(band30.WeightsReason, "band 30's own sweep measured") {
		t.Errorf("band 30 weights_reason = %q, want it to name band 30's own (still-untrustworthy) measurement", band30.WeightsReason)
	}
	if band30.ReferenceDPSPerPoint != nil {
		t.Errorf("band 30 reference_dps_per_point = %v, want nil (weights_reason is set)", *band30.ReferenceDPSPerPoint)
	}
	// The whole point of the guard: an empty slot is only ever
	// published when no candidate exists, never because a sweep was
	// noisy - band 30 must score picks the same way band 20 did
	// (carried weights), not publish every slot no_dps_value.
	allEmptyForNoDPSValue := true
	for _, slot := range band30.Slots {
		if slot.EmptyReason != noDPSValueReason {
			allEmptyForNoDPSValue = false
			break
		}
	}
	if allEmptyForNoDPSValue {
		t.Errorf("band 30 published every slot with empty_reason=%q - the fallback weights should have scored real candidates", noDPSValueReason)
	}
}
