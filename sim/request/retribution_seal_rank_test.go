package request

import (
	"bytes"
	"testing"
)

// Seal of Righteousness rank 7 (20292) and rank 8 (20293). The 70755 client
// gives rank 8's Seal aura a spell damage coefficient of 0 (rank 7: 0.1) and
// every swing proc 0, so Retribution keeps rank 7 at 58 to 60 in both
// presets (sim-accuracy round 18, paladin S1): the rotation authors rank 7.
// Every Seal of Righteousness it casts is a ranked castSpell, never
// castPaladinPrimarySeal (always the top rank learned), so the refresh line
// and the Crusader swap keep rank 7 as well as the twist.
const (
	sealOfRighteousnessRank7 = 20292
	sealOfRighteousnessRank8 = 20293
)

func TestRetributionKeepsSealOfRighteousnessRank7From58(t *testing.T) {
	for _, c := range []struct {
		preset string
		level  int
		want   int32
	}{
		{BarePreset, 57, sealOfRighteousnessRank7},
		{BarePreset, 58, sealOfRighteousnessRank7},
		{BarePreset, 59, sealOfRighteousnessRank7},
		{BarePreset, 60, sealOfRighteousnessRank7},
		{RaidPreset, 60, sealOfRighteousnessRank7},
	} {
		for _, race := range []string{"human", "dwarf", "undead"} {
			ids := rotationIDs(t, "paladin-retribution", "paladin", c.level, specRankChoices("paladin-retribution", c.preset, race))
			if !ids[c.want] || ids[sealOfRighteousnessRank8] {
				t.Errorf("%s %d %s: Seal of Righteousness casts rank 8 (%v), want rank 7 only", c.preset, c.level, race, ids[sealOfRighteousnessRank8])
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
		t.Error("castPaladinPrimarySeal always casts the top rank and skips the Seal of Righteousness rank row; cast the ranked spell")
	}
}
