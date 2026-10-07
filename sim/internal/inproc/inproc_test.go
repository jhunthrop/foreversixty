package inproc

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestRegisterIsIdempotent(t *testing.T) {
	Register()
	Register()
}

func TestNextRunIDIsUniquePerCall(t *testing.T) {
	a, b := nextRunID(), nextRunID()
	if a == b || a == "" || b == "" {
		t.Fatalf("nextRunID returned %q then %q, want two distinct non-empty ids", a, b)
	}
}

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
	got := healingResultOf(player, 100, 300)

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
