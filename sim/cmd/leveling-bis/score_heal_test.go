package main

import (
	"encoding/json"
	"math"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/request"
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

// fakeHealing answers every healing run with a canned result per gear
// signature and records the requests it was asked.
type fakeHealing struct {
	result   inproc.HealingResult
	byGear   func(api.SimRequest) inproc.HealingResult
	requests []api.SimRequest
	profile  request.HealProfile
}

func (f *fakeHealing) Run(req api.SimRequest, profile request.HealProfile) (inproc.HealingResult, error) {
	f.requests = append(f.requests, req)
	f.profile = profile
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

func healingResult(effective, raw, manaLasts float64) inproc.HealingResult {
	return inproc.HealingResult{
		Effective:       api.Estimate{Mean: effective, Error: effective / 50},
		Raw:             api.Estimate{Mean: raw},
		ManaLastsSec:    manaLasts,
		ManaSpent:       10000,
		EffectiveHealed: 4 * effective * 300,
	}
}

func TestLastingFactorIsTheSquaredShareOfTheFightManaLasted(t *testing.T) {
	fight := 300 * time.Second
	cases := []struct {
		name  string
		lasts float64
		want  float64
	}{
		{"lasts the fight", 300, 1},
		{"mana to spare", 900, 1},
		{"four fifths of it", 240, 0.64},
		{"half of it", 150, 0.25},
		{"empties at once", 0, 0},
		{"negative is clamped", -5, 0},
	}
	for _, c := range cases {
		if got := lastingFactor(c.lasts, fight); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("%s: lastingFactor(%v) = %v, want %v", c.name, c.lasts, got, c.want)
		}
	}
	if got := lastingFactor(10, 0); got != 1 {
		t.Errorf("a fight of no length scores %v, want 1", got)
	}
}

// The guard: a set whose mana empties before the fight ends ranks below a
// set that lasts, even when its healing while it lasted was higher.
func TestASetThatEmptiesRanksBelowOneThatLasts(t *testing.T) {
	fight := 300 * time.Second
	burst := healingResult(520, 600, 150)  // higher average, runs dry at half the fight
	steady := healingResult(380, 450, 300) // lower, lasts
	if healingScore(burst, fight) >= healingScore(steady, fight) {
		t.Errorf("the burst set scores %v, the lasting set %v: the one that empties must rank below",
			healingScore(burst, fight), healingScore(steady, fight))
	}
	// And when both last, plain healing decides.
	if healingScore(healingResult(420, 500, 400), fight) <= healingScore(steady, fight) {
		t.Error("two sets that both last must rank by effective healing per second")
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

func TestHealEngineScalesTheErrorByTheGuard(t *testing.T) {
	backend := &fakeHealing{result: healingResult(500, 600, 150)}
	engine := healEngine{backend: backend, profile: testHealProfile(t)}
	score, stdErr, err := engine.RunPlainDPSWithError(api.SimRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(score-125) > 1e-9 || math.Abs(stdErr-2.5) > 1e-9 {
		t.Errorf("score = %v +/- %v, want 125 +/- 2.5 (500 +/- 10 at a quarter)", score, stdErr)
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
	if _, ok := got.(healEngine); !ok || profile == nil || profile.ID == "" {
		t.Errorf("got %T with profile %v, want a healEngine carrying the curated profile", got, profile)
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
