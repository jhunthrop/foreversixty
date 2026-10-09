package score

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestHealingResultReadsTheHealersMetrics(t *testing.T) {
	player := &proto.UnitMetrics{
		EffectiveHps: &proto.DistributionMetrics{Avg: 300, Stdev: 30},
		Hps:          &proto.DistributionMetrics{Avg: 400, Stdev: 40},
		Tto:          &proto.DistributionMetrics{Avg: 275},
		Resources: []*proto.ResourceMetrics{
			{Type: proto.ResourceType_ResourceTypeMana, Gain: -9000},
			{Type: proto.ResourceType_ResourceTypeMana, Gain: -3000},
			{Type: proto.ResourceType_ResourceTypeMana, Gain: 500}, // regeneration is not spending
			{Type: proto.ResourceType_ResourceTypeHealth, Gain: -7000},
		},
	}
	got := HealingResultOf(player, 100, 300)

	if got.Effective.Mean != 300 || got.Raw.Mean != 400 {
		t.Errorf("effective/raw = %v/%v, want 300/400", got.Effective.Mean, got.Raw.Mean)
	}
	if want := 30 / math.Sqrt(100); math.Abs(got.Effective.Error-want) > 1e-9 {
		t.Errorf("effective error = %v, want %v", got.Effective.Error, want)
	}
	if got.OverhealShare() != 0.25 {
		t.Errorf("overheal = %v, want 0.25", got.OverhealShare())
	}
	if got.ManaLastsSec != 275 {
		t.Errorf("mana lasts = %v, want 275", got.ManaLastsSec)
	}
	if got.ManaSpent != 12000 {
		t.Errorf("mana spent = %v, want 12000 (health and regeneration excluded)", got.ManaSpent)
	}
	// 300 effective healing a second for 300 seconds over 100 iterations.
	if want := 300.0 * 300 * 100 / 12000; math.Abs(got.HealingPerMana()-want) > 1e-9 {
		t.Errorf("healing per mana = %v, want %v", got.HealingPerMana(), want)
	}
}

func TestHealingResultWithNothingSpentHasNoHealingPerMana(t *testing.T) {
	if got := (HealingResult{EffectiveHealed: 5000}).HealingPerMana(); got != 0 {
		t.Errorf("healing per mana = %v with no mana spent, want 0", got)
	}
}

func TestMaxHealthFailsWithoutAPlayer(t *testing.T) {
	if _, err := MaxHealth(&proto.Raid{}, &proto.Encounter{}); err == nil {
		t.Fatal("an empty raid has no maximum health")
	}
}

func TestTankEstimateOfScalesStdevByRootN(t *testing.T) {
	if got := EstimateOf(nil, 10); got != (api.Estimate{}) {
		t.Errorf("nil metrics = %+v", got)
	}
	got := EstimateOf(&proto.DistributionMetrics{Avg: 50, Stdev: 20}, 100)
	if got.Mean != 50 || math.Abs(got.Error-2) > 1e-9 {
		t.Errorf("got %+v, want mean 50 error 2", got)
	}
	if got := EstimateOf(&proto.DistributionMetrics{Avg: 50, Stdev: 20}, 0); got.Error != 0 {
		t.Errorf("zero iterations must not divide: %+v", got)
	}
}

func TestOverhealShareHandlesNoRawHealingAndEffectiveAboveRaw(t *testing.T) {
	if got := (HealingResult{}).OverhealShare(); got != 0 {
		t.Errorf("no healing: %v", got)
	}
	over := HealingResult{Effective: api.Estimate{Mean: 120}, Raw: api.Estimate{Mean: 100}}
	if got := over.OverhealShare(); got != 0 {
		t.Errorf("effective above raw must clamp at 0, got %v", got)
	}
}

func TestEstimateOfCarriesTheDistributionAndTheStandardError(t *testing.T) {
	if got := EstimateOf(nil, 5); got != (api.Estimate{}) {
		t.Errorf("nil = %+v", got)
	}
	got := EstimateOf(&proto.DistributionMetrics{Avg: 10, Stdev: 3, Min: 1, Max: 20}, 9)
	if got.Mean != 10 || got.StdDev != 3 || got.Min != 1 || got.Max != 20 || math.Abs(got.Error-1) > 1e-9 {
		t.Errorf("got %+v", got)
	}
	if EstimateOf(&proto.DistributionMetrics{Stdev: 3}, 0).Error != 0 {
		t.Error("zero iterations must leave the error at 0")
	}
}
