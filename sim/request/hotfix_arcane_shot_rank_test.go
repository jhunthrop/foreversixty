package request

import "testing"

// Blizzard's 10 November hotfix: "Arcane Shot mana cost changed to 11% of base
// Mana for all ranks of Arcane Shot" (client 1.60.1.70888 SpellPower). Every
// rank costs the same and a lower rank hits for less, so no rank row holds:
// every hunter casts the newest Arcane Shot it has learned
// (design/reviews/2026-11-11-sim-accuracy/hotfixes-other-classes.md).
func TestArcaneShotCastsItsNewestRankAtEveryLevel(t *testing.T) {
	cases := []hunterRankCase{}
	for _, spec := range []string{hBM, hMM} {
		for _, race := range []string{"troll", "dwarf"} {
			cases = append(cases,
				hunterRankCase{"at 14", spec, BarePreset, race, 14, 14281, 3044},
				hunterRankCase{"at 20", spec, BarePreset, race, 20, 14282, 3044},
				hunterRankCase{"at 34", spec, BarePreset, race, 34, 14283, 14282},
				hunterRankCase{"at 40", spec, BarePreset, race, 40, 14284, 3044},
				hunterRankCase{"at 50", spec, BarePreset, race, 50, 14285, 14282},
				hunterRankCase{"at 57", spec, BarePreset, race, 57, 14286, 14283},
				hunterRankCase{"at bare 60", spec, BarePreset, race, 60, 14287, 14282},
				hunterRankCase{"at raid 60", spec, RaidPreset, race, 60, 14287, 14282},
			)
		}
	}
	checkHunterRankCases(t, cases)
}
