package request

import "testing"

// Sim-accuracy 2026-11-12 (melee-hunters), level by level on the leveling order's
// talents (spec-breakdown -level, bare bands, 10,000 iterations): with Arcane Shot at
// 11% of base mana on every rank, Beast Mastery's rank 1 Serpent Sting beats the
// newest rank by 2.3 to 3.9% at every level from 26 to 60 (Troll and Dwarf), and the
// newest rank wins below 26.
func TestBeastMasterySerpentStingRankOneFrom26(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"at 18 the newest rank", hBM, BarePreset, "troll", 18, 13550, 1978},
		{"at 25 the newest rank", hBM, BarePreset, "dwarf", 25, 13550, 1978},
		{"at 26", hBM, BarePreset, "troll", 26, 1978, 13551},
		{"at 39", hBM, BarePreset, "dwarf", 39, 1978, 13549},
		{"at 40", hBM, BarePreset, "dwarf", 40, 1978, 13552},
		{"at 59", hBM, BarePreset, "troll", 59, 1978, 13555},
		{"at bare 60", hBM, BarePreset, "troll", 60, 1978, 13555},
	})
}

// Sim-accuracy 2026-11-12 (melee-hunters): once Resourcefulness stops cheapening
// Lacerate (440529's mask), Survival runs drier and an older Raptor Strike wins at
// 26 to 29, 40 to 52 and 56 to 59, a rank lower for the Horde races (every hunter
// race at 26 to 29, 31 and 40 to 59; a gain counts at 0.2% on one seed). Rank 1 at
// 31 is gone.
func TestSurvivalRaptorStrikeRowsAfterTheLacerateCost(t *testing.T) {
	const r1, r2, r3, r4, r5, r6, r7, r8 = 2973, 14260, 14261, 14262, 14263, 14264, 14265, 14266
	checkHunterRankCases(t, []hunterRankCase{
		{"Troll at 25 the newest rank", hSV, BarePreset, "troll", 25, r4, r3},
		{"Troll at 26", hSV, BarePreset, "troll", 26, r2, r4},
		{"Dwarf at 26", hSV, BarePreset, "dwarf", 26, r3, r4},
		{"Orc at 28", hSV, BarePreset, "orc", 28, r2, r3},
		{"Human at 28 the newest rank", hSV, BarePreset, "human", 28, r4, r3},
		{"High Order Skyborne at 28", hSV, BarePreset, "high-order-skyborne", 28, r3, r4},
		{"Orc at 29", hSV, BarePreset, "orc", 29, r3, r4},
		{"Night Elf at 29 the newest rank", hSV, BarePreset, "night-elf", 29, r4, r3},
		{"Troll at 31 the newest rank", hSV, BarePreset, "troll", 31, r4, r1},
		{"Orc at 40", hSV, BarePreset, "orc", 40, r4, r6},
		{"Human at 40", hSV, BarePreset, "human", 40, r5, r6},
		{"Dwarf at 45", hSV, BarePreset, "dwarf", 45, r5, r6},
		{"Human at 45 the newest rank", hSV, BarePreset, "human", 45, r6, r5},
		{"Orc at 46", hSV, BarePreset, "orc", 46, r3, r6},
		{"Dwarf at 47", hSV, BarePreset, "dwarf", 47, r4, r6},
		{"Orc at 49", hSV, BarePreset, "orc", 49, r4, r7},
		{"Dwarf at 49", hSV, BarePreset, "dwarf", 49, r5, r7},
		{"Human at 49 the newest rank", hSV, BarePreset, "human", 49, r7, r5},
		{"Windshaper Skyborne at 51 the newest rank", hSV, BarePreset, "windshaper-skyborne", 51, r7, r5},
		{"Tauren at 52", hSV, BarePreset, "tauren", 52, r5, r7},
		{"Orc at 53 the newest rank", hSV, BarePreset, "orc", 53, r7, r5},
		{"Orc at 56", hSV, BarePreset, "orc", 56, r5, r8},
		{"Human at 56 the newest rank", hSV, BarePreset, "human", 56, r8, r5},
		{"Human at 59", hSV, BarePreset, "human", 59, r5, r8},
	})
}

// Sim-accuracy 2026-11-12 (melee-hunters): Marksmanship's Serpent Sting rows, every
// rank at every level 18 to 60 (Troll and Dwarf): rank 1 at 18 and 26 to 32 and at
// bare 60, rank 3 at 42 to 49 and rank 4 at 50 to 59; Rapid Recuperation keeps the
// newest rank ahead at 33 to 41.
func TestMarksmanshipSerpentStingRows(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"at 18", hMM, BarePreset, "dwarf", 18, 1978, 13550},
		{"at 20 the newest rank", hMM, BarePreset, "troll", 20, 13550, 1978},
		{"at 26", hMM, BarePreset, "troll", 26, 1978, 13551},
		{"at 32", hMM, BarePreset, "dwarf", 32, 1978, 13551},
		{"at 33 the newest rank", hMM, BarePreset, "troll", 33, 13551, 1978},
		{"at 41 the newest rank", hMM, BarePreset, "dwarf", 41, 13552, 13550},
		{"at 42", hMM, BarePreset, "troll", 42, 13550, 13553},
		{"at 49", hMM, BarePreset, "dwarf", 49, 13550, 13553},
		{"at 50", hMM, BarePreset, "troll", 50, 13551, 13554},
		{"at 59", hMM, BarePreset, "dwarf", 59, 13551, 13555},
		{"at bare 60", hMM, BarePreset, "troll", 60, 1978, 13555},
	})
}
