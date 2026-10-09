package combine

import (
	"errors"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/score"
)

func healerPart(iters int, mean, stddev, hpm float64) api.SimResult {
	r := req(iters, 0)
	r.Spec, r.RoleMetrics = "priest-holy", true
	return api.SimResult{
		Request: r, IterationsRun: iters, DPS: api.Estimate{Mean: 1, Max: 2},
		Healing: &api.HealingResult{
			EffectiveHPS: api.Estimate{Mean: mean, StdDev: stddev, Max: mean * 2, Min: mean / 2},
			HPS:          api.Estimate{Mean: mean * 1.5},
			ManaLastsSec: mean, HPM: hpm,
		},
	}
}

func tankPart(iters int, dtps, tps, tmi, death float64) api.SimResult {
	r := req(iters, 0)
	r.Spec, r.RoleMetrics = "warrior-protection", true
	return api.SimResult{
		Request: r, IterationsRun: iters, DPS: api.Estimate{Mean: 1, Max: 2},
		Tank: &api.TankResult{
			DTPS: api.Estimate{Mean: dtps, StdDev: dtps / 10}, TPS: api.Estimate{Mean: tps, StdDev: tps / 10},
			TMI: api.Estimate{Mean: tmi, StdDev: 5}, Health: 5000, ChanceOfDeath: death,
		},
	}
}

func TestPooledHealingWeightsByIterationsAndKeepsHealingPerManaARatioOfTotals(t *testing.T) {
	a, b := healerPart(100, 200, 20, 2), healerPart(300, 400, 40, 4)
	got, err := Results([]api.SimResult{a, b})
	if err != nil {
		t.Fatal(err)
	}
	h := got.Healing
	if want := (200*100 + 400*300) / 400.0; math.Abs(h.EffectiveHPS.Mean-want) > 1e-9 {
		t.Errorf("effective hps = %v, want %v", h.EffectiveHPS.Mean, want)
	}
	// healing 20000 + 120000 over mana 10000 + 30000.
	if want := 140000.0 / 40000.0; math.Abs(h.HPM-want) > 1e-9 {
		t.Errorf("hpm = %v, want %v", h.HPM, want)
	}
	if h.EffectiveHPS.Error <= 0 || h.EffectiveHPS.Min != 100 || h.EffectiveHPS.Max != 800 {
		t.Errorf("pooled estimate lost its spread: %+v", h.EffectiveHPS)
	}
	if want := (200*100 + 400*300) / 400.0; math.Abs(h.ManaLastsSec-want) > 1e-9 {
		t.Errorf("mana lasts = %v, want %v", h.ManaLastsSec, want)
	}
}

func TestPooledTankIsScoredFromThePooledTermsNotAveragedFromPartScores(t *testing.T) {
	a, b := tankPart(100, 400, 300, 20, 0.1), tankPart(300, 800, 100, 60, 0)
	got, err := Results([]api.SimResult{a, b})
	if err != nil {
		t.Fatal(err)
	}
	tank := got.Tank
	rawDPS, err := score.BossRawDPS(60)
	if err != nil {
		t.Fatal(err)
	}
	dtps, tps, tmi := (400*100+800*300)/400.0, (300*100+100*300)/400.0, (20*100+60*300)/400.0
	want := score.TankScoreOf(5000*rawDPS/dtps, tps, tmi)
	if math.Abs(tank.Score.Mean-want) > 1e-9*want {
		t.Errorf("score = %v, want the score of the pooled terms %v", tank.Score.Mean, want)
	}
	partScores := (score.TankScoreOf(5000*rawDPS/400, 300, 20)*100 + score.TankScoreOf(5000*rawDPS/800, 100, 60)*300) / 400
	if math.Abs(tank.Score.Mean-partScores) < 1e-6*want {
		t.Error("the score equals the average of the part scores; it must be nonlinear in the pooled terms")
	}
	if math.Abs(tank.ChanceOfDeath-0.025) > 1e-12 || tank.Score.Error <= 0 {
		t.Errorf("chance of death %v, score error %v", tank.ChanceOfDeath, tank.Score.Error)
	}
}

func TestPartsThatDisagreeOnCarryingARoleBlockAreMixed(t *testing.T) {
	a, b := healerPart(100, 200, 20, 2), healerPart(100, 200, 20, 2)
	b.Healing = nil
	if _, err := Results([]api.SimResult{a, b}); !errors.Is(err, ErrMixedParts) {
		t.Fatalf("err = %v, want ErrMixedParts", err)
	}
}

func TestAPlainRunKeepsNoRoleBlocks(t *testing.T) {
	a := api.SimResult{Request: req(100, 0), IterationsRun: 100, DPS: api.Estimate{Mean: 5, Max: 6}}
	got, err := Results([]api.SimResult{a})
	if err != nil || got.Healing != nil || got.Tank != nil {
		t.Fatalf("a plain run gained role blocks: %+v %v", got, err)
	}
}
