package coach

import (
	"strings"
	"testing"
)

func TestCeilingCasts(t *testing.T) {
	for _, c := range []struct {
		duration, cooldown int64
		want               int
	}{
		{60_000, 6_000, 11}, {59_999, 6_000, 10}, {0, 6_000, 0}, {60_000, 0, 0},
	} {
		if got := CeilingCasts(c.duration, c.cooldown); got != c.want {
			t.Errorf("CeilingCasts(%d, %d) = %d, want %d", c.duration, c.cooldown, got, c.want)
		}
	}
}

func TestCoreCasts_CeilingWhenNoComparison(t *testing.T) {
	in := furyFight(t, 60_000, fixtureBloodthirstID, []int64{0, 6000, 12000, 18000, 24000}, repeat(500, 60))
	in.Refs = noPool(7)
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Comparison != nil {
		t.Fatalf("finding %+v", f)
	}
	if !strings.Contains(f.Detail, "Spell data's ceiling for Bloodthirst is 11 casts in this fight.") {
		t.Errorf("detail %q", f.Detail)
	}
	assertNumbers(t, f.Facts, "11")
}

func TestCoreCasts_NoCeilingWhenTheLogsCompare(t *testing.T) {
	in := furyFight(t, 60_000, fixtureBloodthirstID, []int64{0, 6000, 12000, 18000, 24000}, repeat(500, 60))
	in.Refs = simCPM("Bloodthirst", 9, 60)
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil || f.Comparison == nil || strings.Contains(f.Detail, "ceiling") {
		t.Fatalf("finding %+v", f)
	}
}
