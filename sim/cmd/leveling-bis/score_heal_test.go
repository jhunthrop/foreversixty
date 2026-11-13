package main

import (
	"encoding/json"
	"math"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/request"
	simscore "github.com/jhunthrop/foreversixty/sim/score"
)

// testHealProfile is the curated profile, read from the repository the
// way the command reads it.
func testHealProfile(t *testing.T) request.HealProfile {
	t.Helper()
	profile, err := request.LoadHealProfile(filepath.Join("..", "..", "..", "data", "curated", healProfileFile))
	if err != nil {
		t.Fatal(err)
	}
	return profile
}

// testRaidHealProfile is the curated raid-shaped fight.
func testRaidHealProfile(t *testing.T) request.RaidHealProfile {
	t.Helper()
	fight, err := request.LoadRaidHealProfile(filepath.Join("..", "..", "..", "data", "curated", request.RaidHealProfileFile))
	if err != nil {
		t.Fatal(err)
	}
	return fight
}

// fakeHealing answers every healing run with a canned result per gear
// signature and records the requests it was asked. A run on the raid-shaped
// fight answers raidResult and is recorded in raidRequests instead.
type fakeHealing struct {
	result       simscore.HealingResult
	raidResult   simscore.HealingResult
	byGear       func(api.SimRequest) simscore.HealingResult
	requests     []api.SimRequest
	raidRequests []api.SimRequest
	profile      request.HealProfile
}

func (f *fakeHealing) Run(req api.SimRequest, fight request.HealFight) (simscore.HealingResult, error) {
	if _, raid := fight.(request.RaidHealProfile); raid {
		f.raidRequests = append(f.raidRequests, req)
		return f.raidResult, nil
	}
	f.requests = append(f.requests, req)
	f.profile = fight.(request.HealProfile)
	if f.byGear != nil {
		return f.byGear(req), nil
	}
	return f.result, nil
}

func (f *fakeHealing) Weights(req api.SimRequest, profile request.HealProfile) (map[string]api.StatWeight, float64, error) {
	f.requests = append(f.requests, req)
	f.profile = profile
	return map[string]api.StatWeight{"healing_power": {Stat: "healing_power", Weight: 1}}, 0.5, nil
}

func healingResult(effective, raw, manaLasts float64) simscore.HealingResult {
	return simscore.HealingResult{
		Effective:       api.Estimate{Mean: effective, Error: effective / 50},
		Raw:             api.Estimate{Mean: raw},
		ManaLastsSec:    manaLasts,
		ManaSpent:       10000,
		EffectiveHealed: 4 * effective * 300,
	}
}

// Mana is not a second term: a set that runs dry is ranked on the effective
// healing per second it averaged over the whole fight, which already pays
// for the seconds it did not heal.
func TestASetIsRankedOnItsEffectiveHealingWhetherOrNotItRanDry(t *testing.T) {
	dry := healingResult(520, 600, 150)    // higher average, runs dry at half the fight
	steady := healingResult(380, 450, 300) // lower, lasts
	for name, result := range map[string]simscore.HealingResult{"dry": dry, "steady": steady} {
		engine := healEngine{backend: &fakeHealing{result: result}, profile: testHealProfile(t)}
		score, _, err := engine.RunPlainDPSWithError(api.SimRequest{})
		if err != nil {
			t.Fatal(err)
		}
		if score != result.Effective.Mean {
			t.Errorf("%s: score = %v, want its effective healing %v", name, score, result.Effective.Mean)
		}
	}
}

func TestHealEngineRunsEveryRequestOverTheProfilesFight(t *testing.T) {
	backend := &fakeHealing{result: healingResult(400, 500, 300)}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}

	req := api.SimRequest{Encounter: api.EncounterSpec{DurationSec: 180, Variation: 0.2}}
	score, stdErr, err := engine.RunPlainDPSWithError(req)
	if err != nil {
		t.Fatal(err)
	}
	got := backend.requests[0].Encounter
	if got.DurationSec != engine.profile.DurationSec || got.Variation != 0 {
		t.Errorf("fight = %d s +/- %v, want the profile's %d s exactly", got.DurationSec, got.Variation, engine.profile.DurationSec)
	}
	if score != 400 || math.Abs(stdErr-8) > 1e-9 {
		t.Errorf("score = %v +/- %v, want 400 +/- 8 for a set that lasts", score, stdErr)
	}
	if plain, _ := engine.RunPlainDPS(req); plain != score {
		t.Errorf("RunPlainDPS = %v, want the same %v", plain, score)
	}
}

func TestHealEngineReturnsTheHealingsOwnErrorAndScalesIterations(t *testing.T) {
	backend := &fakeHealing{result: healingResult(500, 600, 150)}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}
	score, stdErr, err := engine.RunPlainDPSWithError(api.SimRequest{Iterations: 300})
	if err != nil {
		t.Fatal(err)
	}
	if score != 500 || math.Abs(stdErr-10) > 1e-9 {
		t.Errorf("score = %v +/- %v, want 500 +/- 10 whatever the mana did", score, stdErr)
	}
	if got := backend.requests[0].Iterations; got != 300*healIterationScale {
		t.Errorf("ran %d iterations, want %d", got, 300*healIterationScale)
	}
}

func TestHealEngineWeightsGoThroughTheBackendOverTheProfilesFight(t *testing.T) {
	backend := &fakeHealing{}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}
	weights, perPoint, err := engine.RunWeights(api.SimRequest{Encounter: api.EncounterSpec{DurationSec: 60}})
	if err != nil {
		t.Fatal(err)
	}
	if weights["healing_power"].Weight != 1 || perPoint != 0.5 {
		t.Errorf("weights = %+v per point %v, want the backend's", weights, perPoint)
	}
	if backend.requests[0].Encounter.DurationSec != engine.profile.DurationSec {
		t.Errorf("the sweep ran over %d s, want the profile's %d s", backend.requests[0].Encounter.DurationSec, engine.profile.DurationSec)
	}
}

func TestHealEngineHasNoHitProfile(t *testing.T) {
	profile, err := healEngine{}.HitProfileFor(api.SimRequest{})
	if err != nil || profile.Physical {
		t.Errorf("hit profile = %+v, %v; a healer has no miss table", profile, err)
	}
}

func TestMetricsForPublishesTheContractsFigures(t *testing.T) {
	result := healingResult(400, 500, 5000)
	result.ManaSpent = 20000
	result.EffectiveHealed = 400 * 300 * 10
	metrics := metricsFor(result)

	if metrics.HPS != 400 || metrics.RawHPS != 500 {
		t.Errorf("hps/raw = %v/%v, want 400/500", metrics.HPS, metrics.RawHPS)
	}
	// The run's own standard error rides with the figure, so /tiers/healer never orders two
	// healers by less than the measurement can tell (owner ruling, round 5).
	if metrics.HPSError != result.Effective.Error || metrics.HPSError == 0 {
		t.Errorf("hps_error = %v, want the run's %v", metrics.HPSError, result.Effective.Error)
	}
	if math.Abs(metrics.OverhealPct-0.2) > 1e-9 {
		t.Errorf("overheal = %v, want 0.2", metrics.OverhealPct)
	}
	if metrics.ManaLastsSec != maxManaLastsSec {
		t.Errorf("mana lasts = %v, want it capped at %v", metrics.ManaLastsSec, maxManaLastsSec)
	}
	if want := 400.0 * 300 * 10 / 20000; math.Abs(metrics.HPM-want) > 1e-9 {
		t.Errorf("hpm = %v, want %v", metrics.HPM, want)
	}
}

func TestWithHealerEngineLeavesADamageSpecAlone(t *testing.T) {
	var damage engineRunner = &fakeEngine{}
	got, profile, err := withHealerEngine(damage, "../../..", specInfo{Spec: "mage-fire", Role: "dps"})
	if err != nil || profile != nil || got != damage {
		t.Errorf("got %v, %v, %v; a damage spec keeps its runner and loads no profile", got, profile, err)
	}
}

func TestWithHealerEngineLoadsTheProfileForAHealer(t *testing.T) {
	got, profile, err := withHealerEngine(&fakeEngine{}, "../../..", specInfo{Spec: "priest-holy", Role: "healer"})
	if err != nil {
		t.Fatal(err)
	}
	engine, ok := got.(healEngine)
	if !ok || profile == nil || profile.ID == "" {
		t.Fatalf("got %T with profile %v, want a healEngine carrying the curated profile", got, profile)
	}
	if engine.raidFight.ID == "" || engine.raidFight.ID == profile.ID {
		t.Errorf("raid fight = %q, want the curated raid-shaped fight beside %q", engine.raidFight.ID, profile.ID)
	}
}

// The raid-shaped fight picks nothing: a band 60 set runs it once more for the metrics, and
// the BiS figures (hps, set_dps) stay the BiS profile's (owner ruling, 2026-11-13).
func TestAttachHealerFieldsRunsTheRaidFightAtBand60(t *testing.T) {
	raid := healingResult(120, 200, 3600).WithManaPool(9000)
	backend := &fakeHealing{result: healingResult(800, 820, 200), raidResult: raid}
	engine := healEngine{backend: backend, profile: testHealProfile(t), raidFight: testRaidHealProfile(t)}
	var report bandReport
	spec := specInfo{Spec: "priest-holy", ClassSlug: "priest", Role: "healer"}
	if err := attachHealerFields(&report, engine, spec, "human", "priest", raidFightBand, "", map[string]slotPick{}); err != nil {
		t.Fatal(err)
	}
	metrics := report.Metrics.(*healingMetrics)
	if report.SetDPS != 800 || metrics.HPS != 800 {
		t.Errorf("set_dps = %v, hps = %v, want the BiS profile's 800", report.SetDPS, metrics.HPS)
	}
	got := metrics.RaidFight
	if got == nil {
		t.Fatal("a band 60 healer published no raid_fight block")
	}
	want := raidFightMetrics{
		Profile: engine.raidFight.ID, DurationSec: engine.raidFight.DurationSec,
		HPS: 120, HPSError: raid.Effective.Error, RawHPS: 200, OverhealPct: 0.4,
		ManaLeft: raid.ManaLeft, ManaLeftPct: raid.ManaLeftShare(),
	}
	if math.Abs(got.OverhealPct-want.OverhealPct) > 1e-9 {
		t.Errorf("overheal = %v, want %v", got.OverhealPct, want.OverhealPct)
	}
	got.OverhealPct = want.OverhealPct
	if *got != want {
		t.Errorf("raid_fight = %+v, want %+v", *got, want)
	}
	if len(backend.raidRequests) != 1 {
		t.Fatalf("ran the raid fight %d times, want once", len(backend.raidRequests))
	}
	req := backend.raidRequests[0]
	if req.Encounter.DurationSec != engine.raidFight.DurationSec || req.Encounter.Variation != 0 || req.Iterations != healMetricsIterations {
		t.Errorf("raid fight request = %d s +/- %v at %d iterations, want %d s exactly at %d", req.Encounter.DurationSec, req.Encounter.Variation, req.Iterations, engine.raidFight.DurationSec, healMetricsIterations)
	}
	raw, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Metrics map[string]json.RawMessage `json:"metrics"`
	}
	if err := json.Unmarshal(raw, &file); err != nil || file.Metrics["raid_fight"] == nil {
		t.Errorf("the band entry's metrics carry no raid_fight key: %s", raw)
	}
}

func TestAttachHealerFieldsSkipsTheRaidFightBelow60(t *testing.T) {
	backend := &fakeHealing{result: healingResult(300, 320, 200)}
	engine := healEngine{backend: backend, profile: testHealProfile(t), raidFight: testRaidHealProfile(t)}
	var report bandReport
	spec := specInfo{Spec: "priest-holy", ClassSlug: "priest", Role: "healer"}
	if err := attachHealerFields(&report, engine, spec, "human", "priest", 50, "", map[string]slotPick{}); err != nil {
		t.Fatal(err)
	}
	if report.Metrics.(*healingMetrics).RaidFight != nil || len(backend.raidRequests) != 0 {
		t.Error("a band 50 set ran the level 60 raid fight")
	}
}

func TestWithHealerEngineFailsWithoutAProfile(t *testing.T) {
	if _, _, err := withHealerEngine(&fakeEngine{}, t.TempDir(), specInfo{Spec: "priest-holy", Role: "healer"}); err == nil {
		t.Error("a healer with no curated profile ranked without one")
	}
}

func TestAttachHealerFieldsMeasuresTheFinalSetAndSetsSetDPS(t *testing.T) {
	backend := &fakeHealing{result: healingResult(410, 500, 300)}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}
	var report bandReport
	spec := specInfo{Spec: "priest-holy", ClassSlug: "priest", Role: "healer"}

	if err := attachHealerFields(&report, engine, spec, "human", "priest", 60, "", map[string]slotPick{}); err != nil {
		t.Fatal(err)
	}
	metrics, isHealing := report.Metrics.(*healingMetrics)
	if report.Role != "healer" || report.Profile != engine.profile.ID || !isHealing || metrics == nil {
		t.Fatalf("report = %+v, want role, profile and healing metrics", report)
	}
	if report.SetDPS != 410 || metrics.HPS != 410 {
		t.Errorf("set_dps = %v, hps = %v, want both 410", report.SetDPS, metrics.HPS)
	}
	if backend.requests[0].Iterations != healMetricsIterations {
		t.Errorf("metrics ran %d iterations, want %d", backend.requests[0].Iterations, healMetricsIterations)
	}
}

func TestAttachHealerFieldsLeavesADamageReportAlone(t *testing.T) {
	report := bandReport{SetDPS: 123}
	if err := attachHealerFields(&report, &fakeEngine{}, specInfo{Spec: "mage-fire"}, "gnome", "mage", 60, "", nil); err != nil {
		t.Fatal(err)
	}
	if report.SetDPS != 123 || report.Role != "" || report.Metrics != nil {
		t.Errorf("a damage report changed: %+v", report)
	}
}

func TestOnlyAHealersReportCarriesHealerKeys(t *testing.T) {
	damage, err := json.Marshal(bandReport{Spec: "mage-fire"})
	if err != nil {
		t.Fatal(err)
	}
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(damage, &keys); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"role", "profile", "metrics"} {
		if _, present := keys[key]; present {
			t.Errorf("a damage spec's entry carries %q", key)
		}
	}

	metrics := healingMetrics{HPS: 1}
	healer, err := json.Marshal(bandReport{Spec: "priest-holy", Role: "healer", Profile: "p", Metrics: &metrics})
	if err != nil {
		t.Fatal(err)
	}
	keys = map[string]json.RawMessage{}
	if err := json.Unmarshal(healer, &keys); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"role", "profile", "metrics", "set_dps", "talents"} {
		if _, present := keys[key]; !present {
			t.Errorf("a healer's entry lacks %q", key)
		}
	}
}

func TestAHealersItemHealingIsWeighedByTheHealingPowerRow(t *testing.T) {
	weights := map[string]float64{"healing_power": 1, "intellect": 0.4, "spell_power": 0}
	if got := statWeight("healing", weights); got != 1 {
		t.Errorf("an item's healing weighs %v, want the healing_power row's 1", got)
	}
	if got := statWeight("intellect", weights); got != 0.4 {
		t.Errorf("intellect weighs %v, want 0.4", got)
	}
}

// The tuned profile must not set an order on its own (owner ruling, round 5): the metrics
// carry the final set's effective HPS and error at each tuning bound as well, measured
// under the profile scaled to that bound.
func TestAttachHealerFieldsMeasuresTheSetAtBothTuningBounds(t *testing.T) {
	profile := testHealProfile(t)
	var scales []float64
	backend := &fakeHealing{}
	backend.byGear = func(api.SimRequest) simscore.HealingResult {
		scale := backend.profile.Tuning.Scale
		scales = append(scales, scale)
		return healingResult(400*scale, 500, 300)
	}
	engine := healEngine{backend: backend, profile: profile}
	var report bandReport
	spec := specInfo{Spec: "priest-holy", ClassSlug: "priest", Role: "healer"}
	if err := attachHealerFields(&report, engine, spec, "human", "priest", 60, "", map[string]slotPick{}); err != nil {
		t.Fatal(err)
	}
	metrics := report.Metrics.(*healingMetrics)
	if len(metrics.AtBounds) != len(profile.Tuning.Bounds) {
		t.Fatalf("hps_at_bounds = %+v, want one row per bound %v", metrics.AtBounds, profile.Tuning.Bounds)
	}
	for i, bound := range profile.Tuning.Bounds {
		row := metrics.AtBounds[i]
		if row.Scale != bound || math.Abs(row.HPS-400*bound) > 1e-9 || row.HPSError == 0 {
			t.Errorf("bound %v: %+v, want hps %v with its error", bound, row, 400*bound)
		}
	}
	if metrics.HPS != 400*profile.Tuning.Scale {
		t.Errorf("hps = %v, want the tuned profile's %v", metrics.HPS, 400*profile.Tuning.Scale)
	}
	if engine.profile.Tank.HitDamage != profile.Tank.HitDamage {
		t.Error("measuring at the bounds changed the engine's profile")
	}
}

// Round 10: a healer's ceiling re-measure (trinketGainAdaptiveIterationCeiling,
// 2000) scaled by healIterationScale asked for 14000 iterations, which the
// request layer refuses (api.MaxIterations). Every healer trinket and relic
// whose quick gain was not yet significant then published empty: Holy
// paladin's Libram of Infusion at 60. The scaled count stops at the cap.
func TestHealEngineNeverAsksForMoreThanTheRequestLayerAllows(t *testing.T) {
	backend := &fakeHealing{result: healingResult(500, 600, 150)}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}
	if _, _, err := engine.RunPlainDPSWithError(api.SimRequest{Iterations: trinketGainAdaptiveIterationCeiling}); err != nil {
		t.Fatal(err)
	}
	if got := backend.requests[0].Iterations; got != api.MaxIterations {
		t.Errorf("ran %d iterations, want the cap %d", got, api.MaxIterations)
	}
}
