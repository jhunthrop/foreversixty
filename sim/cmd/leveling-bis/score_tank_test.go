package main

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	simscore "github.com/jhunthrop/foreversixty/sim/score"
)

func tankRun(health, dtps, tps, tmi float64) simscore.TankRunResult {
	return simscore.TankRunResult{
		Health: health,
		DTPS:   api.Estimate{Mean: dtps, Error: dtps * 0.01},
		TPS:    api.Estimate{Mean: tps, Error: tps * 0.02},
		TMI:    api.Estimate{Mean: tmi, Error: 0.5},
		DPS:    api.Estimate{Mean: 123},
	}
}

func TestEffectiveHealthIsHealthOverTheShareOfRawDamageTaken(t *testing.T) {
	f, err := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	if err != nil {
		t.Fatal(err)
	}
	if f.EffectiveHealth != 20000 {
		t.Errorf("effective health = %v, want 10000 health at half the raw damage = 20000", f.EffectiveHealth)
	}
	if f.DPS != 123 {
		t.Errorf("own damage = %v, want it carried through", f.DPS)
	}
}

func TestATankTheBossNeverHitsIsRefused(t *testing.T) {
	if _, err := simscore.NewTankFigures(tankRun(10000, 0, 400, 30), 1000); err == nil {
		t.Fatal("a run with no damage taken must fail, not score infinite effective health")
	}
}

func TestMitigationIsAOneForOneMultiplier(t *testing.T) {
	a, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	b, _ := simscore.NewTankFigures(tankRun(10000, 450, 400, 30), 1000)
	if got := b.Score / a.Score; math.Abs(got-500.0/450.0) > 1e-9 {
		t.Errorf("10%% less damage taken scaled the score by %v, want %v", got, 500.0/450.0)
	}
}

func TestThreatEntersAtAQuarterPower(t *testing.T) {
	a, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	b, _ := simscore.NewTankFigures(tankRun(10000, 500, 800, 30), 1000)
	if got, want := b.Score/a.Score, math.Pow(2, simscore.TankThreatExponent); math.Abs(got-want) > 1e-9 {
		t.Errorf("doubling threat scaled the score by %v, want %v", got, want)
	}
}

func TestEachTMIPointCostsItsShareOfScore(t *testing.T) {
	a, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	b, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 40), 1000)
	if got, want := b.Score/a.Score, math.Exp(-simscore.TankRiskPerTMI*10); math.Abs(got-want) > 1e-9 {
		t.Errorf("ten more TMI scaled the score by %v, want %v", got, want)
	}
}

func TestTheScoreErrorSumsTheThreeTermsErrors(t *testing.T) {
	f, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	want := f.Score * (0.01 + simscore.TankThreatExponent*0.02 + simscore.TankRiskPerTMI*0.5)
	if math.Abs(f.ScoreError-want) > 1e-9 {
		t.Errorf("score error = %v, want %v", f.ScoreError, want)
	}
}

func TestStaminaWeighsOneAndTheOthersAreRatiosOfTheScoreSlope(t *testing.T) {
	base, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	sweeps := map[string]statSweep{
		// 10 health a point, nothing else moves.
		"stamina": {health: 10},
		// -0.5 damage taken per second a point of armor, no health.
		"armor": {dtps: -0.05},
		// A stat that only costs: more threat-free damage taken.
		"hit": {dtps: 0.5},
	}
	rows, perPoint, err := tankWeightsFromSweeps([]string{"stamina", "armor", "hit"}, "stamina", base, 10000, sweeps)
	if err != nil {
		t.Fatal(err)
	}
	byStat := map[string]api.StatWeight{}
	for _, r := range rows {
		byStat[r.Stat] = r
	}
	if byStat["stamina"].Weight != 1 {
		t.Errorf("reference weight = %v, want exactly 1", byStat["stamina"].Weight)
	}
	if want := base.Score * 10 / 10000; math.Abs(perPoint-want) > 1e-9 {
		t.Errorf("score per point of stamina = %v, want %v", perPoint, want)
	}
	// armor: 0.05/500 = 1e-4 relative per point; stamina 10/10000 = 1e-3.
	if got := byStat["armor"].Weight; math.Abs(got-0.1) > 1e-9 {
		t.Errorf("armor weight = %v, want 0.1 of a stamina", got)
	}
	if byStat["hit"].Weight >= 0 {
		t.Errorf("a stat that raises damage taken must weigh negative, got %v", byStat["hit"].Weight)
	}
}

func TestAReferenceThatDoesNotHelpCannotNormalise(t *testing.T) {
	base, _ := simscore.NewTankFigures(tankRun(10000, 500, 400, 30), 1000)
	if _, _, err := tankWeightsFromSweeps([]string{"stamina"}, "stamina", base, 10000, map[string]statSweep{"stamina": {}}); err == nil {
		t.Fatal("a reference with no effect must be refused")
	}
}

func TestTheTankRunnerOnlyTakesATankSpec(t *testing.T) {
	if _, ok := roleRunner(realEngine{}, specInfo{Role: "dps"}).(tankEngine); ok {
		t.Error("a dps spec must keep the damage runner")
	}
	if _, ok := roleRunner(realEngine{}, specInfo{Role: roleTank}).(tankEngine); !ok {
		t.Error("a tank spec must get the tank runner")
	}
}

func TestTheWeightsCharacterIsDressedInTheMostArmoredNearBestPiece(t *testing.T) {
	items := []candidate{
		{ID: 1, Slots: []string{"head"}, Quality: 3, ItemLevel: 60, Armor: 500, Stats: map[string]float64{"stamina": 10}},
		{ID: 2, Slots: []string{"head"}, Quality: 3, ItemLevel: 58, Armor: 900, Stats: map[string]float64{"stamina": 5}},
		{ID: 3, Slots: []string{"head"}, Quality: 3, ItemLevel: 40, Armor: 2000},
		{ID: 4, Slots: []string{"head"}, Quality: 1, ItemLevel: 60, Armor: 3000},
		{ID: 5, Slots: []string{"off_hand"}, Quality: 3, ItemLevel: 60, Armor: 0},
		{ID: 6, Slots: []string{"off_hand"}, Quality: 3, ItemLevel: 60, Armor: 2000},
	}
	ch := tankLadderCharacter(specInfo{Role: roleTank}, items, 60, api.CharacterSpec{Gear: []api.GearSlot{{Slot: "main_hand", ItemID: 99}}})
	got := map[string]int{}
	for _, g := range ch.Gear {
		got[g.Slot] = g.ItemID
	}
	if got["head"] != 2 {
		t.Errorf("head = %d, want 2 (most armor within the item level window; 3 is too low, 4 too poor)", got["head"])
	}
	if got["off_hand"] != 6 {
		t.Errorf("off hand = %d, want the shield 6", got["off_hand"])
	}
	if got["main_hand"] != 99 {
		t.Error("the ladder's weapon must stay")
	}
}

func TestADamageSpecsWeightsCharacterIsUntouched(t *testing.T) {
	in := api.CharacterSpec{Gear: []api.GearSlot{{Slot: "main_hand", ItemID: 99}}}
	out := tankLadderCharacter(specInfo{Role: "dps"}, []candidate{{ID: 1, Slots: []string{"head"}, Quality: 3}}, 60, in)
	if len(out.Gear) != 1 {
		t.Errorf("gear = %v, want it unchanged", out.Gear)
	}
}

func TestAnItemsArmorScoresAtTheArmorWeightAndNowhereElse(t *testing.T) {
	c := candidate{Armor: 100, Stats: map[string]float64{"stamina": 10}}
	if got := score(c, "head", map[string]float64{"stamina": 1, "armor": 0.1}, 0, false); math.Abs(got-20) > 1e-9 {
		t.Errorf("score = %v, want 10 stamina + 100 armor at 0.1 = 20", got)
	}
	if got := score(c, "head", map[string]float64{"stamina": 1}, 0, false); got != 10 {
		t.Errorf("score = %v, want armor to add nothing without a weight", got)
	}
}
