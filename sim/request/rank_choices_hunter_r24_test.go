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

// Round 24 (hunter lane): the round-22 Arcane Shot rows were judged at each
// band floor with the band floor's talents, so a hunter above the floor ran
// without the mana talents he has by then. Re-judged level by level on the
// talents the leveling order gives at the level (spec-breakdown -level,
// bare bands, 10,000 iterations, seeds 1 and 50001, Troll and Dwarf), a
// Marksmanship hunter's best Arcane Shot from 33 is the newest-but-one or
// the one before it, not rank 1 or rank 3.
func TestRound24MarksmanshipArcaneShotRankRows(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"at 32 keeps rank 3", hMM, BarePreset, "troll", 32, 14282, 14283},
		{"at 33", hMM, BarePreset, "troll", 33, 14283, 14282},
		{"at 35", hMM, BarePreset, "dwarf", 35, 14283, 14284},
		{"at 36", hMM, BarePreset, "troll", 36, 14284, 14283},
		{"at 41", hMM, BarePreset, "dwarf", 41, 14284, 3044},
		{"at 42", hMM, BarePreset, "troll", 42, 14283, 14284},
		{"at 47", hMM, BarePreset, "dwarf", 47, 14283, 14284},
		{"at 48", hMM, BarePreset, "troll", 48, 14284, 3044},
		{"at 59", hMM, BarePreset, "dwarf", 59, 14284, 3044},
		{"at bare 60", hMM, BarePreset, "troll", 60, 14283, 3044},
		{"raid 60 keeps the top rank", hMM, RaidPreset, "troll", 60, 14287, 14283},
	})
}

// Round 24: the same re-judgement for Beast Mastery. Rank 1 now starts at 20
// (it ran from 26), rank 3 wins at 34 to 39 inside the rank 1 stretch, and
// rank 4 beats rank 3 at 54 to 57 (at 53 for a Troll only; a Dwarf ties).
func TestRound24BeastMasteryArcaneShotRankRows(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"at 20", hBM, BarePreset, "troll", 20, 3044, 14282},
		{"at 25", hBM, BarePreset, "dwarf", 25, 3044, 14282},
		{"at 33", hBM, BarePreset, "troll", 33, 3044, 14282},
		{"at 34", hBM, BarePreset, "dwarf", 34, 14282, 3044},
		{"at 39", hBM, BarePreset, "troll", 39, 14282, 3044},
		{"at 40", hBM, BarePreset, "dwarf", 40, 3044, 14282},
		{"at 52", hBM, BarePreset, "troll", 52, 14282, 14283},
		{"Troll at 53", hBM, BarePreset, "troll", 53, 14283, 14282},
		{"Dwarf at 53", hBM, BarePreset, "dwarf", 53, 14282, 14283},
		{"at 54", hBM, BarePreset, "dwarf", 54, 14283, 14282},
		{"at 57", hBM, BarePreset, "troll", 57, 14283, 14282},
		{"at 58", hBM, BarePreset, "dwarf", 58, 14282, 14283},
		{"at bare 60", hBM, BarePreset, "troll", 60, 14282, 14287},
	})
}

// Round 24: Beast Mastery's Serpent Sting rank 1 row now starts at 20, and
// rank 2 beats it at 38 and 39 next to Arcane Shot rank 3.
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
