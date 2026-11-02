package request

import "testing"

// Rows re-measured at sim-accuracy round 11 (design/reviews/2026-11-03-sim-audit-round11/ranks.md):
// each case is the id a spec's published rotation casts for one ranked spell at one level
// and preset, and the id it must no longer cast there.
func TestReMeasuredRankChoicesCastTheWinningRank(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset string
		level                     int
		want, not                 int32
	}{
		// Lava Burst rank 2 (1238299) beats rank 1 (408490) from 50, beyond error at both seeds.
		{"Lava Burst at 50", "shaman-elemental", "shaman", BarePreset, 50, 1238299, 408490},
		{"Lava Burst at 59", "shaman-elemental", "shaman", BarePreset, 59, 1238299, 408490},
		{"Lava Burst at 45 keeps rank 1", "shaman-elemental", "shaman", BarePreset, 45, 408490, 1238299},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
