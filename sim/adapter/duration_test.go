package adapter

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestSumActorTotalsIncludesPets(t *testing.T) {
	actors := []summary.Actor{{Total: 1000}, {Total: 200}, {Total: 50}}
	if got, want := sumActorTotals(actors), int64(1250); got != want {
		t.Errorf("sumActorTotals(player+2 pets) = %d, want %d", got, want)
	}
}

func TestSumActorTotalsOfNoActors(t *testing.T) {
	if got := sumActorTotals(nil); got != 0 {
		t.Errorf("sumActorTotals(nil) = %d, want 0", got)
	}
}

// The formula: the derived duration makes totalDamage / duration equal
// meanDPS exactly (to integer-millisecond precision).
func TestDeriveDurationMSMatchesTheFormula(t *testing.T) {
	got := deriveDurationMS(101805, 566.1262, 179856)
	want := int64(179827) // round(101805 / 566.1262 * 1000)
	if got != want {
		t.Errorf("deriveDurationMS(101805, 566.1262, ...) = %d, want %d", got, want)
	}
	// Recovering a rate from an integer-millisecond duration cannot be
	// exact - see durationTolerance in golden_test.go for the same
	// bound applied to the checked-in fixtures - so this checks the
	// recovered rate is close, not bit-identical.
	gotDPS := float64(101805) / (float64(got) / 1000)
	if diff := gotDPS - 566.1262; diff < -durationTolerance || diff > durationTolerance {
		t.Errorf("total/derived duration = %v, want within %v of meanDPS 566.1262", gotDPS, durationTolerance)
	}
}

func TestDeriveDurationMSFallsBackOnZeroDamage(t *testing.T) {
	if got := deriveDurationMS(0, 500, 180000); got != 180000 {
		t.Errorf("deriveDurationMS(0 damage, ...) = %d, want the fallback 180000", got)
	}
}

func TestDeriveDurationMSFallsBackOnNonPositiveMeanDPS(t *testing.T) {
	if got := deriveDurationMS(50000, 0, 180000); got != 180000 {
		t.Errorf("deriveDurationMS(damage, 0 DPS, ...) = %d, want the fallback 180000", got)
	}
	if got := deriveDurationMS(50000, -12, 180000); got != 180000 {
		t.Errorf("deriveDurationMS(damage, negative DPS, ...) = %d, want the fallback 180000", got)
	}
}

// The fallback itself must never be zero in practice - a zero duration
// turns every per-second figure downstream into a division by zero -
// but deriveDurationMS is not responsible for validating its caller's
// fallback; it just has to hand it back unchanged on the degenerate
// paths, which this and the two tests above pin.
func TestDeriveDurationMSNeverDividesByZero(t *testing.T) {
	for _, tc := range []struct {
		name     string
		damage   int64
		meanDPS  float64
		fallback int64
	}{
		{"zero damage, positive fallback", 0, 500, 180000},
		{"zero DPS, positive fallback", 50000, 0, 180000},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := deriveDurationMS(tc.damage, tc.meanDPS, tc.fallback)
			if got <= 0 {
				t.Fatalf("deriveDurationMS(%d, %v, %d) = %d, want a positive duration", tc.damage, tc.meanDPS, tc.fallback, got)
			}
		})
	}
}
