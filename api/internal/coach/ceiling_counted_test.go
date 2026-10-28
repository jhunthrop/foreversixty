package coach

import (
	"strings"
	"testing"
)

func TestCoreCasts_CeilingIsOverCountedTime(t *testing.T) {
	in := furyFightDeadAt(t, 60_000, fixtureBloodthirstID, []int64{0, 6000}, repeat(500, 60), 30_000)
	in.Refs = noPool(7)
	f := coreCasts(in, CountedTime(in)).Finding
	if f == nil {
		t.Fatal("a player dead at 30 s with two casts should have a finding")
	}
	if !strings.Contains(f.Detail, "ceiling for Bloodthirst is 6 casts") {
		t.Errorf("detail %q", f.Detail)
	}
	assertNumbers(t, f.Facts, "6")
}
