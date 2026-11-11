package request

import "testing"

// hunterRankCase is one level of a rank row: the spell id the rotation must cast
// and a neighbouring rank it must not.
type hunterRankCase struct {
	name, spec, preset, race string
	level                    int
	want, not                int32
}

func checkHunterRankCases(t *testing.T, cases []hunterRankCase) {
	t.Helper()
	for _, c := range cases {
		ids := rotationIDs(t, c.spec, "hunter", c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

const (
	hMM = "hunter-marksmanship"
	hBM = "hunter-beast-mastery"
	hSV = "hunter-survival"
)

// Round 24: Beast Mastery's Serpent Sting rank 1 row now starts at 20, and
// rank 2 beats it at 38 and 39.
func TestRound24BeastMasterySerpentStingRankRows(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"at 20", hBM, BarePreset, "troll", 20, 1978, 13550},
		{"at 37", hBM, BarePreset, "dwarf", 37, 1978, 13549},
		{"at 38", hBM, BarePreset, "troll", 38, 13549, 1978},
		{"at 39", hBM, BarePreset, "dwarf", 39, 13549, 13552},
		{"at 40 keeps the default", hBM, BarePreset, "troll", 40, 13552, 13549},
	})
}

// Round 24: Survival with only the first point of Resourcefulness casts Raptor
// Strike rank 1; with the second point in the newest rank wins by over 2%.
// Round 28 moved Resourcefulness to 31 and 32 (ahead of Deterrence and
// Counterattack), so the rank 1 level is now 31 (+0.7 to +1.1% over rank 4,
// every race, seeds 1 and 50001) and 33 keeps rank 5 (rank 1 loses 2.2 to
// 2.4%). At 40 a Troll gains about 0.2% with rank 5 over rank 6; a Dwarf ties.
func TestRound24SurvivalRaptorStrikeRankRows(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"at 31", hSV, BarePreset, "troll", 31, 2973, 14262},
		{"Dwarf at 31", hSV, BarePreset, "dwarf", 31, 2973, 14262},
		{"at 32 keeps the default", hSV, BarePreset, "troll", 32, 14263, 2973},
		{"at 33 keeps the default", hSV, BarePreset, "troll", 33, 14263, 2973},
		{"at 34 keeps the default", hSV, BarePreset, "dwarf", 34, 14263, 2973},
		// Round 28: with Predator's Edge in place of Deterrence an Orc at 27 gains
		// 0.2% with rank 2 over rank 3 at seeds 1 and 50001.
		{"Orc at 27", hSV, BarePreset, "orc", 27, 14260, 14261},
		{"Dwarf at 27 keeps rank 3", hSV, BarePreset, "dwarf", 27, 14261, 14260},
		{"Troll at 40", hSV, BarePreset, "troll", 40, 14263, 14264},
		{"Dwarf at 40 keeps the default", hSV, BarePreset, "dwarf", 40, 14264, 14263},
		{"Troll at 41 keeps the default", hSV, BarePreset, "troll", 41, 14264, 14263},
	})
}
