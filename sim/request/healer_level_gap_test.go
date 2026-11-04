package request

import (
	"encoding/json"
	"slices"
	"strings"
	"testing"
)

// Sim-accuracy round 15 wave 2: Forever's level-gap rule (a rank keeps its
// whole spell power coefficient until the caster is 17 levels past its learn
// level, then 5% less a level; engine core.LevelPenalty) makes the rank-1 and
// rank-2 heals the lists cast at 60 heal with none of the healer's healing
// power. Each list now authors a rank that keeps it, and rows keep the low
// rank only at the bands where it still wins on the band's own fight (seeds
// 1 and 50001, design/reviews/2026-11-04-sim-audit-round15/priest.md, wave 2).

// linesCasting is the spec's list cut down to the top-level lines that cast
// one of ids, as a rotation RewriteRotationRanks reads.
func linesCasting(t *testing.T, spec string, ids ...int) json.RawMessage {
	t.Helper()
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/"+spec+".json", &list)
	// Every entry keeps its place (the others empty), so a rank choice scoped
	// to lines still finds its line.
	keep := make([]json.RawMessage, len(list.Rotation.PriorityList))
	found := false
	for index, line := range list.Rotation.PriorityList {
		keep[index] = json.RawMessage(`{}`)
		var entry struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
			} `json:"action"`
		}
		if err := json.Unmarshal(line, &entry); err != nil {
			t.Fatal(err)
		}
		if entry.Action.CastSpell == nil {
			continue
		}
		if slices.Contains(ids, entry.Action.CastSpell.SpellID.SpellID) {
			keep[index], found = line, true
		}
	}
	if !found {
		t.Fatalf("%s casts none of %v", spec, ids)
	}
	out, err := json.Marshal(map[string]any{"type": "TypeAPL", "priorityList": keep})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

type rankCase struct {
	level  int
	preset string
	rank   int
}

func assertRanks(t *testing.T, rotation json.RawMessage, spec, class, race string, family map[int]int, cases []rankCase) {
	t.Helper()
	for _, c := range cases {
		assertOnlyRank(t, rotation, spec, class, c.preset, race, c.level, family, c.rank)
	}
}

var (
	renewRanks       = map[int]int{1: 139, 2: 6074, 3: 6075, 4: 6076, 5: 6077, 6: 6078, 7: 10927, 8: 10928, 9: 10929}
	healRanks        = map[int]int{1: 2054, 2: 2055, 3: 6063, 4: 6064}
	lesserHealRanks  = map[int]int{1: 2050, 2: 2052, 3: 2053}
	greaterHealRanks = map[int]int{1: 2060, 2: 10963, 3: 10964, 4: 10965}
	rejuvenationRank = map[int]int{1: 774, 2: 1058, 3: 1430, 4: 2090, 5: 2091, 6: 3627, 7: 8910, 8: 9839, 9: 9840, 10: 9841}
	flashOfLightRank = map[int]int{1: 19750, 2: 19939, 3: 19940, 4: 19941, 5: 19942, 6: 19943}
	holyLightRanks   = map[int]int{1: 635, 2: 639, 3: 647, 4: 1026, 5: 1042, 6: 3472, 7: 10328, 8: 10329}
	mindBlastRanks   = map[int]int{1: 8092, 2: 8102, 3: 8103, 4: 8104, 5: 8105, 6: 8106, 7: 10945, 8: 10946, 9: 10947}
)

// Both priests' tank Renew is rank 7 (10927, learned at 44, whole bonus at
// 60); rank 1 still wins to 29 (Holy to 32) and rank 3 at 39 (round 16
// priest wave 2: Holy and Discipline split their Renew rows level by level).
// Round 18: at 40-59 a cheaper rank with its whole bonus wins too: rank 4 at
// 40, then Holy rank 5 at 50 and Discipline the list's rank 7.
func TestPriestsRenewKeepsItsHealingBonus(t *testing.T) {
	atFifty := map[string]int{"priest-holy": 5, "priest-discipline": 7}
	for _, spec := range []string{"priest-holy", "priest-discipline"} {
		renew := linesCasting(t, spec, 139, 10927)
		if strings.Contains(string(renew), `"spellId":139`) && !strings.Contains(string(renew), `"spellId":10927`) {
			t.Errorf("%s still authors Renew rank 1", spec)
		}
		assertRanks(t, renew, spec, "priest", "human", renewRanks, []rankCase{
			{20, BarePreset, 1}, {29, BarePreset, 1}, {39, BarePreset, 3}, {40, BarePreset, 4}, {50, BarePreset, atFifty[spec]},
			{60, BarePreset, 7}, {60, RaidPreset, 7},
		})
	}
}

// Holy's filler is Greater Heal at its top trainer rank (10965), ahead of
// each Heal line; Heal stays behind it at its top rank (6064) for the levels
// before Greater Heal (learned at 40), where it casts the highest rank known.
// Round 18: Greater Heal is rank 1 from 46 to 57, then the highest rank known.
func TestHolyFillsWithGreaterHealAheadOfHeal(t *testing.T) {
	heal := linesCasting(t, "priest-holy", 2054, 6064)
	greater := linesCasting(t, "priest-holy", 10965)
	assertRanks(t, heal, "priest-holy", "priest", "troll", healRanks, []rankCase{
		{16, BarePreset, 1}, {22, BarePreset, 2}, {28, BarePreset, 3}, {34, BarePreset, 4}, {50, BarePreset, 4},
		{60, BarePreset, 4}, {60, RaidPreset, 4},
	})
	assertRanks(t, greater, "priest-holy", "priest", "troll", greaterHealRanks, []rankCase{
		{40, BarePreset, 1}, {50, BarePreset, 1}, {58, BarePreset, 4}, {60, BarePreset, 4}, {60, RaidPreset, 4},
	})
	var h, g []any
	_ = json.Unmarshal(heal, &struct{ PriorityList *[]any }{&h})
	_ = json.Unmarshal(greater, &struct{ PriorityList *[]any }{&g})
	if len(h) != len(g) {
		t.Errorf("Holy has %d Heal lines and %d Greater Heal lines, want one Greater Heal line ahead of each", len(h), len(g))
	}
}

// Discipline's party filler is Greater Heal rank 1 (2060), ahead of each
// Lesser Heal line, which falls back to rank 3 (2053) before 40.
func TestDisciplineFillsThePartyWithGreaterHealRankOne(t *testing.T) {
	lesser := linesCasting(t, "priest-discipline", 2050, 2053)
	assertRanks(t, lesser, "priest-discipline", "priest", "undead", lesserHealRanks, []rankCase{
		{4, BarePreset, 2}, {10, BarePreset, 3}, {50, BarePreset, 3}, {60, BarePreset, 3}, {60, RaidPreset, 3},
	})
	party := linesCasting(t, "priest-discipline", 2060)
	assertRanks(t, party, "priest-discipline", "priest", "undead", greaterHealRanks, []rankCase{
		{40, BarePreset, 1}, {60, BarePreset, 1}, {60, RaidPreset, 1},
	})
}

// The druid's party Rejuvenation is rank 7 (8910) in a raid; on every bare
// fight, bare 60 included, rank 1 still wins (the fight runs dry), so rows
// keep it. The tank's rank-10 line follows round 16's level-by-level sweep
// (design/reviews/2026-11-04-sim-audit-round16/druid.md, wave 2): rank 2 at
// 20 to 37, rank 4 at 38 to 49, rank 6 at 50 to 52, the top rank at 53 to 57,
// rank 8 at 58 and 59, and the top rank at 60.
func TestRestorationDruidRejuvenationKeepsItsHealingBonusInARaid(t *testing.T) {
	party := linesCasting(t, "druid-restoration", 774, 8910)
	assertRanks(t, party, "druid-restoration", "druid", "tauren", rejuvenationRank, []rankCase{
		{20, BarePreset, 1}, {59, BarePreset, 1}, {60, BarePreset, 1}, {60, RaidPreset, 7},
	})
	tank := linesCasting(t, "druid-restoration", 9841)
	assertRanks(t, tank, "druid-restoration", "druid", "tauren", rejuvenationRank, []rankCase{
		{20, BarePreset, 2}, {30, BarePreset, 2}, {37, BarePreset, 2}, {38, BarePreset, 4}, {49, BarePreset, 4},
		{50, BarePreset, 6}, {52, BarePreset, 6}, {53, BarePreset, 9}, {58, BarePreset, 8}, {59, BarePreset, 8},
		{60, BarePreset, 10}, {60, RaidPreset, 10},
	})
}

// The shaman's tank Healing Wave is the top rank (10396) at 60: rank 2
// keeps none of the healing bonus from level 40. Below 60 the shaman lane's
// rows pick the rank by level (TestRestorationShamanFillerHealingWaveByBand).
func TestRestorationShamanTankHealingWaveIsTheTopRank(t *testing.T) {
	filler, _ := restorationHealingWaveLines(t)
	assertRanks(t, filler, "shaman-restoration", "shaman", "tauren", healingWaveRanks, []rankCase{
		{12, BarePreset, 3}, {24, BarePreset, 2}, {32, BarePreset, 3}, {40, BarePreset, 4}, {50, BarePreset, 6},
		{60, BarePreset, 9}, {60, RaidPreset, 9},
	})
}

// Holy paladin: the filler Flash of Light is rank 5 (19942) in a raid. Round
// 16 wave 2 swept the rows level by level on the MaxLevel + 12 rule: the
// tank's Flash of Light keeps rank 1 to 43, rank 2 at 44 to 49, rank 4 at 50
// to 53 and rank 5, as authored, from 54; Holy Light casts rank 1 to 24, rank 2 at 25
// to 29, its top rank at 30 to 46, rank 1 at 47 to 49 and rank 2 at 50 to 59.
func TestHolyPaladinFlashOfLightAndHolyLightKeepTheirHealingBonus(t *testing.T) {
	filler := linesCasting(t, "paladin-holy", 19750, 19942)
	assertRanks(t, filler, "paladin-holy", "paladin", "human", flashOfLightRank, []rankCase{
		{20, BarePreset, 1}, {39, BarePreset, 1}, {40, BarePreset, 1}, {43, BarePreset, 1}, {44, BarePreset, 2}, {49, BarePreset, 2},
		{50, BarePreset, 4}, {53, BarePreset, 4}, {54, BarePreset, 5}, {59, BarePreset, 5},
		{60, BarePreset, 5}, {60, RaidPreset, 5},
	})
	light := linesCasting(t, "paladin-holy", 10329)
	assertRanks(t, light, "paladin-holy", "paladin", "human", holyLightRanks, []rankCase{
		{20, BarePreset, 1}, {24, BarePreset, 1}, {25, BarePreset, 2}, {29, BarePreset, 2}, {30, BarePreset, 5},
		{40, BarePreset, 6}, {46, BarePreset, 7}, {47, BarePreset, 1}, {49, BarePreset, 1}, {50, BarePreset, 2}, {59, BarePreset, 2}, {60, BarePreset, 8}, {60, RaidPreset, 8},
	})
}

// Shadow's Mind Blast at 50 to 59 casts the highest rank learned: ranks 2 and
// 3 keep 15% and 45% of spell power at 50, and the rows lost 3.3% there.
func TestShadowMindBlastAtFiftyIsTheHighestRankLearned(t *testing.T) {
	blast := linesCasting(t, "priest-shadow", 10947)
	for _, race := range []string{"gnome", "undead", "human", "dwarf"} {
		assertRanks(t, blast, "priest-shadow", "priest", race, mindBlastRanks, []rankCase{
			{50, BarePreset, 7}, {52, BarePreset, 8}, {58, BarePreset, 9},
		})
	}
}

var powerWordShieldRanks = map[int]int{1: 17, 2: 592, 3: 600, 4: 3747, 5: 6065, 6: 6066, 7: 10898, 8: 10899, 9: 10900, 10: 10901}

// Discipline shields at the top rank learned. Round 14 held rank 4 (3747)
// because a full shield was thought to cost a warrior tank his rage; Forever
// removed that (Blizzard's 30 September priest and warrior deep dive: "Rage
// gain also ignores all damage absorption effects, so Power Word: Shield will
// no longer decrease Rage gained"), and under the level-gap rule rank 4 keeps
// 5% of its bonus at 60: rank 10 is +1.0 to +1.4% at raid 60, both seeds, and
// no bare band loses.
func TestDisciplineShieldsAtTheTopRank(t *testing.T) {
	shield := linesCasting(t, "priest-discipline", 3747, 10901)
	assertRanks(t, shield, "priest-discipline", "priest", "human", powerWordShieldRanks, []rankCase{
		{20, BarePreset, 3}, {30, BarePreset, 5}, {40, BarePreset, 6}, {45, BarePreset, 7}, {50, BarePreset, 8},
		{60, BarePreset, 10}, {60, RaidPreset, 10},
	})
}
