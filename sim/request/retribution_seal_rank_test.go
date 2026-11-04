package request

import (
	"bytes"
	"testing"
)

// Seal of Righteousness rank 7 (20292) and rank 8 (20293). Every rank takes
// the same share of spell damage, the one the 70755 client's tooltip states
// (0.0765 with a one-hander, 0.12 with a two-hander: SpellDescriptionVariables
// 1023, written in 70425 with the "almost twice" fix), so Retribution casts
// rank 8 from 58 in both presets (sim-accuracy round 20, paladin S1; round
// 18 had read rank 8's aura column as 0 and kept rank 7). Every Seal of
// Righteousness it casts is a ranked castSpell, never castPaladinPrimarySeal,
// so the refresh line, the Crusader swap and the twist render on the card.
const (
	sealOfRighteousnessRank7 = 20292
	sealOfRighteousnessRank8 = 20293
)

func TestRetributionCastsSealOfRighteousnessRank8From58(t *testing.T) {
	for _, c := range []struct {
		preset string
		level  int
		want   int32
		not    int32
	}{
		{BarePreset, 57, sealOfRighteousnessRank7, sealOfRighteousnessRank8},
		{BarePreset, 58, sealOfRighteousnessRank8, sealOfRighteousnessRank7},
		{BarePreset, 59, sealOfRighteousnessRank8, sealOfRighteousnessRank7},
		{BarePreset, 60, sealOfRighteousnessRank8, sealOfRighteousnessRank7},
		{RaidPreset, 60, sealOfRighteousnessRank8, sealOfRighteousnessRank7},
	} {
		for _, race := range []string{"human", "dwarf", "undead"} {
			ids := rotationIDs(t, "paladin-retribution", "paladin", c.level, specRankChoices("paladin-retribution", c.preset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s %d %s: Seal of Righteousness casts %d (%v), want %d only", c.preset, c.level, race, c.not, ids[c.not], c.want)
			}
		}
	}
}

func TestRetributionCastsNoPrimarySealAction(t *testing.T) {
	raw, err := aplFS.ReadFile("apl/paladin-retribution.apl.json")
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(raw, []byte(`"castPaladinPrimarySeal"`)) {
		t.Error("castPaladinPrimarySeal does not render on the card; cast the ranked spell")
	}
}
