package request

import (
	"slices"
	"testing"
)

// The round 9 hybrid audit (design/reviews/2026-11-02-sim-audit-round9/hybrids.md)
// pins the rotation lines it corrected.

// N-WFT: Windfury Weapon on the main hand "disables any benefit you
// personally receive from Windfury Totem" (16362), and the raid preset
// already carries a Windfury Totem, so Enhancement's own Air totem is Grace
// of Air.
func TestEnhancementDropsGraceOfAirNotItsOwnWindfuryTotem(t *testing.T) {
	const (
		windfuryTotem = 10614
		graceOfAir    = 10627
	)
	ids := castOrder(t, "shaman-enhancement")
	if slices.Contains(ids, windfuryTotem) {
		t.Errorf("Enhancement casts its own Windfury Totem (%d), which its Windfury Weapon disables: %v", windfuryTotem, ids)
	}
	if !slices.Contains(ids, graceOfAir) {
		t.Errorf("Enhancement never drops Grace of Air Totem (%d): %v", graceOfAir, ids)
	}
}
