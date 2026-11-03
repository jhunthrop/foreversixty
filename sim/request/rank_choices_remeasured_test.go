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
		// Since Swiftmend leaves the effect up (fork 9badd243f), Rejuvenation rank 10 (9841)
		// loses to the rotation's own rank 4 (2090) at bare 60.
		{"Rejuvenation at bare 60", "druid-restoration", "druid", BarePreset, 60, 2090, 9841},
		// Once the published Undead casts Dark Sacrifice when mana is low (casters lane,
		// db8fc067e), Devouring Plague rank 1 (2944) loses 4.4% at bare 60 to the top
		// rank (19280); Shadow Word: Death rank 2 still wins there.
		{"Devouring Plague at bare 60", "priest-shadow", "priest", BarePreset, 60, 19280, 2944},
		{"Shadow Word: Death at bare 60 keeps rank 2", "priest-shadow", "priest", BarePreset, 60, 1309633, 1309636},
		// Without the level penalty on heals (healers lane, 074edfadf), Healing Wave rank 2
		// (332) beats rank 3 (547) by 5.6% at bare 60, and the fight still runs dry.
		{"Healing Wave at bare 60", "shaman-restoration", "shaman", BarePreset, 60, 332, 547},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, HordeFaction))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
