package request

import (
	"bytes"
	"encoding/json"
	"slices"
	"strconv"
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
	holyShockRanks   = map[int]int{1: 1311606, 2: 20473, 3: 20929, 4: 20930}
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
// Round 18: Greater Heal is rank 1 from 46 to 57; round 24: to 59 as well, the
// highest rank known only at 60.
func TestHolyFillsWithGreaterHealAheadOfHeal(t *testing.T) {
	heal := linesCasting(t, "priest-holy", 2054, 6064)
	greater := linesCasting(t, "priest-holy", 10965)
	assertRanks(t, heal, "priest-holy", "priest", "troll", healRanks, []rankCase{
		{16, BarePreset, 1}, {22, BarePreset, 1}, {28, BarePreset, 3}, {34, BarePreset, 1}, {38, BarePreset, 2}, {39, BarePreset, 2}, {40, BarePreset, 4}, {50, BarePreset, 4},
		{60, BarePreset, 4}, {60, RaidPreset, 4},
	})
	assertRanks(t, greater, "priest-holy", "priest", "troll", greaterHealRanks, []rankCase{
		{40, BarePreset, 1}, {50, BarePreset, 1}, {58, BarePreset, 1}, {59, BarePreset, 1}, {60, BarePreset, 4}, {60, RaidPreset, 4},
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
// (design/reviews/2026-11-04-sim-audit-round16/druid.md, wave 2) and round
// 19's (round19/druid.md) and round 20's (round20/druid.md, rank 2 to 39 on
// band-30 gear): rank 2 at 23 to 39 (rank 1 to 22, round 24), rank 4 at 40 to 43, rank 5 at 44 to 48,
// rank 6 at 49 to 55, rank 9 at 56 to 58 and rank 8 at 59 (round 30), and the top rank at 60.
func TestRestorationDruidRejuvenationKeepsItsHealingBonusInARaid(t *testing.T) {
	party := linesCasting(t, "druid-restoration", 774, 8910)
	assertRanks(t, party, "druid-restoration", "druid", "tauren", rejuvenationRank, []rankCase{
		{20, BarePreset, 1}, {59, BarePreset, 1}, {60, BarePreset, 1}, {60, RaidPreset, 7},
	})
	tank := linesCasting(t, "druid-restoration", 9841)
	assertRanks(t, tank, "druid-restoration", "druid", "tauren", rejuvenationRank, []rankCase{
		{23, BarePreset, 2}, {30, BarePreset, 2}, {37, BarePreset, 2}, {39, BarePreset, 2}, {40, BarePreset, 4}, {43, BarePreset, 4},
		{44, BarePreset, 5}, {48, BarePreset, 5}, {49, BarePreset, 6}, {55, BarePreset, 6}, {56, BarePreset, 9}, {57, BarePreset, 9},
		{58, BarePreset, 9}, {59, BarePreset, 8},
		{60, BarePreset, 10}, {60, RaidPreset, 10},
	})
}

// The shaman's tank Healing Wave is the top rank (10396) at 60: rank 2
// keeps none of the healing bonus from level 40. Below 60 the shaman lane's
// rows pick the rank by level (TestRestorationShamanFillerHealingWaveByBand).
func TestRestorationShamanTankHealingWaveIsTheTopRank(t *testing.T) {
	filler, _ := restorationHealingWaveLines(t)
	assertRanks(t, filler, "shaman-restoration", "shaman", "tauren", healingWaveRanks, []rankCase{
		{12, BarePreset, 3}, {24, BarePreset, 3}, {32, BarePreset, 3}, {40, BarePreset, 5}, {50, BarePreset, 6},
		{60, BarePreset, 9}, {60, RaidPreset, 9},
	})
}

// Holy paladin: the filler Flash of Light is rank 5 (19942) in a raid. Round
// 16 wave 2 swept the rows level by level on the MaxLevel + 12 rule; round 19
// re-judged 40 to 49; round 20 re-judged 43 to 60 after Oath of Light began
// adding its whole amount under the rank's level rule (10,000 iterations,
// seeds 1 and 50001, both factions). The tank's Flash of Light keeps rank 1
// to 39, rank 2 at 40 to 49 (round 29: the new leveling order's Conviction and
// Divine Intellect at 40 to 42 make rank 2 win there too), rank 3 at 50 to 54, rank 4 at 55 to 59 and at bare 60 (rank 5, as authored, in a
// raid). Holy Light casts rank 1 to 23, its top rank at 24, rank 1 at 25,
// rank 2 at 26 to 29 (round 24, on the talents at each level), its top rank
// at 30 to 39, rank 4 at 40 to 42, its top
// rank at 43 to 46, rank 4 at 47 to 49 and its top rank from 50: the cheap
// ranks round 16 kept at 47 to 59 leaned on an Oath bonus the level rule
// now takes away. Sim-accuracy 2026-11-12 (Holy paladin lane), with
// Concentration Aura applied and Illumination's casting regen gone: the
// tank's Flash of Light is rank 1 to 42 and rank 2 from 43, rank 3 to 53
// for a Human and rank 4 from 54; Holy Light is rank 1 to 28 and rank 2
// at 29; the group's Flash of Light is rank 1 at 40 and 41 and, for a
// Human, rank 2 from 42.
func TestHolyPaladinFlashOfLightAndHolyLightKeepTheirHealingBonus(t *testing.T) {
	filler := linesCasting(t, "paladin-holy", 19750, 19942)
	assertRanks(t, filler, "paladin-holy", "paladin", "human", flashOfLightRank, []rankCase{
		{20, BarePreset, 1}, {39, BarePreset, 1}, {40, BarePreset, 1}, {42, BarePreset, 1}, {43, BarePreset, 2}, {48, BarePreset, 2}, {49, BarePreset, 2},
		{50, BarePreset, 3}, {53, BarePreset, 3}, {54, BarePreset, 4}, {55, BarePreset, 4}, {59, BarePreset, 4},
		{60, BarePreset, 4}, {60, RaidPreset, 5},
	})
	light := linesCasting(t, "paladin-holy", 10329)
	assertRanks(t, light, "paladin-holy", "paladin", "human", holyLightRanks, []rankCase{
		{20, BarePreset, 1}, {23, BarePreset, 1}, {24, BarePreset, 1}, {25, BarePreset, 1}, {26, BarePreset, 1}, {28, BarePreset, 1}, {29, BarePreset, 2}, {30, BarePreset, 5},
		{39, BarePreset, 6}, {40, BarePreset, 4}, {42, BarePreset, 4}, {43, BarePreset, 6}, {46, BarePreset, 7}, {47, BarePreset, 4}, {49, BarePreset, 4}, {50, BarePreset, 7}, {59, BarePreset, 8}, {60, BarePreset, 8}, {60, RaidPreset, 8},
	})
}

// Holy paladin, round 19: the group's Flash of Light (authored rank 4, 19941)
// is rank 2 at 44 to 46, where rank 1 (MaxLevel 25) has lost most of its
// bonus and rank 2 (MaxLevel 31) still keeps 85 to 95% of it. Round 20: rank
// 2 to 49 as well (+1.0 to +1.3% over rank 1 at 47 to 49), then the top rank.
// Round 24, on the talents at each level: rank 2 at 40 to 42 too (+0.1 to
// +0.3% over rank 1), rank 1 at 43.
// Holy Shock casts rank 1 at 40 to 49: it keeps its whole bonus to 46 for 160
// mana against rank 2's 225, and +0.3 to +0.6% at 47 to 49; from 50 the ranks
// tie and the top rank is cast. Round 22: rank 5 in a raid at 60.
func TestHolyPaladinGroupFlashOfLightAndHolyShockBelowFifty(t *testing.T) {
	group := linesCasting(t, "paladin-holy", 19941)
	assertRanks(t, group, "paladin-holy", "paladin", "human", flashOfLightRank, []rankCase{
		{39, BarePreset, 1}, {40, BarePreset, 1}, {41, BarePreset, 1}, {42, BarePreset, 2}, {43, BarePreset, 2}, {44, BarePreset, 2}, {46, BarePreset, 2}, {47, BarePreset, 2}, {49, BarePreset, 2},
		{50, BarePreset, 4}, {60, RaidPreset, 5},
	})
	shock := linesCasting(t, "paladin-holy", 20930)
	assertRanks(t, shock, "paladin-holy", "paladin", "human", holyShockRanks, []rankCase{
		{30, BarePreset, 1}, {39, BarePreset, 1}, {40, BarePreset, 1}, {46, BarePreset, 1}, {47, BarePreset, 1}, {49, BarePreset, 1},
		{50, BarePreset, 3}, {56, BarePreset, 4}, {60, RaidPreset, 4},
	})
}

// Shadow's Mind Blast at 50 to 59 casts the highest rank learned: ranks 2 and
// 3 keep 15% and 45% of spell power at 50, and the rows lost 3.3% there.
// Sim-accuracy 2026-11-12 (shadow), on the wand's first-shot windup: a Dwarf
// casts rank 6 at 50 (whole bonus, cheaper, ahead of rank 7) and the top rank
// from 51; a Gnome, Human and Undead cast the highest rank learned.
func TestShadowMindBlastAtFiftyIsTheHighestRankLearned(t *testing.T) {
	blast := linesCasting(t, "priest-shadow", 10947)
	for race, at := range map[string][3]int{
		"gnome": {7, 7, 8}, "human": {7, 7, 8}, "undead": {7, 7, 8}, "dwarf": {6, 7, 8},
	} {
		assertRanks(t, blast, "priest-shadow", "priest", race, mindBlastRanks, []rankCase{
			{50, BarePreset, at[0]}, {51, BarePreset, at[1]}, {52, BarePreset, at[2]}, {58, BarePreset, 9},
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
// no bare band loses. Round 23: a Human keeps rank 5 at 40 (+0.3%, its first
// level on band-40 gear), the one level here where a race row holds the older
// rank (rank_choices_priest_r23_test.go).
func TestDisciplineShieldsAtTheTopRank(t *testing.T) {
	shield := linesCasting(t, "priest-discipline", 3747, 10901)
	assertRanks(t, shield, "priest-discipline", "priest", "human", powerWordShieldRanks, []rankCase{
		{20, BarePreset, 3}, {30, BarePreset, 5}, {40, BarePreset, 5}, {41, BarePreset, 6}, {45, BarePreset, 7}, {50, BarePreset, 8},
		{60, BarePreset, 10}, {60, RaidPreset, 10},
	})
}

// Sim-accuracy round 18: when Holy's top-rank Greater Heal is out of reach
// (it would overheal or the mana is short), Greater Heal rank 1 (2060) comes
// before Heal rank 4. At 60 Heal rank 4 keeps 55% of the healing bonus and
// rank 1 85%: +0.5% to +1.0% effective healing at raid 60 and +0.2% to +0.4%
// at bare 60, seeds 1 and 50001, both factions. Each Heal line has a rank-1
// Greater Heal line with its condition right before it; before 40 those lines
// drop out and Heal is the filler as before.
func TestHolySteppingDownCastsGreaterHealRankOneBeforeHeal(t *testing.T) {
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/priest-holy.json", &list)
	castOf := func(line json.RawMessage) (int, json.RawMessage) {
		var entry struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
					Target json.RawMessage `json:"target"`
				} `json:"castSpell"`
			} `json:"action"`
			Condition json.RawMessage `json:"condition"`
		}
		if err := json.Unmarshal(line, &entry); err != nil {
			t.Fatal(err)
		}
		if entry.Action.CastSpell == nil {
			return 0, nil
		}
		return entry.Action.CastSpell.SpellID.SpellID, append(entry.Action.CastSpell.Target, entry.Condition...)
	}
	heals := 0
	for index, line := range list.Rotation.PriorityList {
		id, key := castOf(line)
		if id != 6064 {
			continue
		}
		heals++
		match := false
		for _, earlier := range list.Rotation.PriorityList[:index] {
			if eid, ekey := castOf(earlier); eid == 2060 && string(ekey) == string(key) {
				match = true
			}
		}
		if !match {
			t.Errorf("Heal line %d has no Greater Heal rank 1 line with its target and condition ahead of it", index)
		}
	}
	if heals == 0 {
		t.Fatal("Holy casts no Heal rank 4")
	}
	stepDown := linesCasting(t, "priest-holy", 2060)
	assertRanks(t, stepDown, "priest-holy", "priest", "troll", greaterHealRanks, []rankCase{
		{40, BarePreset, 1}, {60, BarePreset, 1}, {60, RaidPreset, 1},
	})
	if out, err := RewriteRotationRanks(stepDown, "priest-holy", "priest", BarePreset, "troll", 30); err != nil {
		t.Fatal(err)
	} else if strings.Contains(string(out), "2060") {
		t.Error("a level 30 priest casts Greater Heal rank 1")
	}
}

// Sim-accuracy round 19: Discipline's tank filler is Greater Heal (10965),
// learned at 40, so below 40 the tank had no filler at all (Penance, Renew
// and the shield only, with Lesser Heal and Flash Heal on the party). A Heal
// line behind each tank Greater Heal line, with its target and condition,
// is +7.4% / +9.1% effective healing at bare 20 and +2.7% / +2.1% at bare 30
// (Alliance / Horde, seeds 1 and 50001); from 40 the Greater Heal line ahead
// of it takes the cast and nothing moves.
func TestDisciplineHealsTheTankWithHealBeforeGreaterHeal(t *testing.T) {
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/priest-discipline.json", &list)
	type cast struct {
		Action struct {
			CastSpell *struct {
				SpellID struct {
					SpellID int `json:"spellId"`
				} `json:"spellId"`
				Target json.RawMessage `json:"target"`
			} `json:"castSpell"`
			Condition json.RawMessage `json:"condition"`
		} `json:"action"`
	}
	parse := func(line json.RawMessage) (int, string) {
		var c cast
		if err := json.Unmarshal(line, &c); err != nil {
			t.Fatal(err)
		}
		if c.Action.CastSpell == nil {
			return 0, ""
		}
		condition := c.Action.Condition
		if c.Action.CastSpell.SpellID.SpellID == 6064 {
			// The Heal line's condition is the Greater Heal line's, and
			// Greater Heal rank 1 not yet learned: at 60 a Heal rank 4
			// fallback on the tank costs 0.8% to 1.1% at raid 60.
			var gated struct {
				And struct {
					Vals []json.RawMessage `json:"vals"`
				} `json:"and"`
			}
			if err := json.Unmarshal(c.Action.Condition, &gated); err != nil || len(gated.And.Vals) != 2 {
				return c.Action.CastSpell.SpellID.SpellID, "ungated"
			}
			if compactJSON(t, gated.And.Vals[1]) != `{"not":{"val":{"spellIsKnown":{"spellId":{"spellId":2060,"rank":1}}}}}` {
				return c.Action.CastSpell.SpellID.SpellID, "ungated"
			}
			condition = gated.And.Vals[0]
		}
		return c.Action.CastSpell.SpellID.SpellID, compactJSON(t, c.Action.CastSpell.Target) + compactJSON(t, condition)
	}
	lines := list.Rotation.PriorityList
	tankFillers := 0
	for index, line := range lines {
		id, key := parse(line)
		if id != 10965 {
			continue
		}
		tankFillers++
		if index+1 >= len(lines) {
			t.Fatalf("tank Greater Heal line %d is the last line, with no Heal behind it", index)
		}
		if nid, nkey := parse(lines[index+1]); nid != 6064 || nkey != key {
			t.Errorf("tank Greater Heal line %d is not followed by a Heal rank 4 line with its target and condition, gated on Greater Heal being unlearned", index)
		}
	}
	if tankFillers == 0 {
		t.Fatal("Discipline has no tank Greater Heal line")
	}
	heal := linesCasting(t, "priest-discipline", 6064)
	assertRanks(t, heal, "priest-discipline", "priest", "human", healRanks, []rankCase{
		{20, BarePreset, 1}, {30, BarePreset, 1}, {36, BarePreset, 3}, {60, RaidPreset, 4},
	})
}

// compactJSON is raw with its insignificant whitespace removed.
func compactJSON(t *testing.T, raw json.RawMessage) string {
	t.Helper()
	var out bytes.Buffer
	if len(raw) == 0 {
		return ""
	}
	if err := json.Compact(&out, raw); err != nil {
		t.Fatal(err)
	}
	return out.String()
}

// Sim-accuracy round 19: from 40 Discipline casts no Lesser Heal rank 3
// (2053, learned at 10). Its line fires when Greater Heal rank 1 is out of
// mana reach, and by 47 the level-gap rule has taken its whole bonus: each
// cast heals for little and breaks the five-second rule. Without it the bar
// regenerates for the next Greater Heal: +7.7% / +6.4% at bare 40, +13% /
// +11% at 43, +21% / +17% at 49 and +45% / +38% at 53 (Alliance / Horde,
// seeds 1 and 50001). A Heal rank 4 line in its place is cast 0.1 times a
// fight (it needs the mana Greater Heal rank 1 lacked) and adds nothing.
// Round 20 took the gate's upper bound away: at 60 the rank keeps none of
// its bonus either (client MaxLevel 15, whole bonus to 27), the level-gap
// rule has no top-rank exemption (HealRankAdvisor measured Lesser Heal rank
// 3 penalised in game at 30), and Blizzard's own examples of the rule are
// at 60. Before 40 Lesser Heal stays.
// scopedBelowForty reports whether one of an and's terms is currentLevel < N
// with N at most 40: round 20's line reads < 40, round 22's tank lines < 16.
func scopedBelowForty(t *testing.T, vals []json.RawMessage) bool {
	t.Helper()
	for _, v := range vals {
		var c struct {
			Cmp struct {
				Op  string `json:"op"`
				Lhs struct {
					CurrentLevel *struct{} `json:"currentLevel"`
				} `json:"lhs"`
				Rhs struct {
					Const struct {
						Val string `json:"val"`
					} `json:"const"`
				} `json:"rhs"`
			} `json:"cmp"`
		}
		if json.Unmarshal(v, &c) != nil || c.Cmp.Op != "OpLt" || c.Cmp.Lhs.CurrentLevel == nil {
			continue
		}
		if n, err := strconv.Atoi(c.Cmp.Rhs.Const.Val); err == nil && n <= 40 {
			return true
		}
	}
	return false
}

func TestDisciplineSkipsLesserHealFromForty(t *testing.T) {
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/priest-discipline.json", &list)
	lessers := 0
	for index, raw := range list.Rotation.PriorityList {
		var l struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
				Condition json.RawMessage `json:"condition"`
			} `json:"action"`
		}
		if err := json.Unmarshal(raw, &l); err != nil {
			t.Fatal(err)
		}
		if l.Action.CastSpell == nil || l.Action.CastSpell.SpellID.SpellID != 2053 {
			continue
		}
		lessers++
		var and struct {
			And struct {
				Vals []json.RawMessage `json:"vals"`
			} `json:"and"`
		}
		if err := json.Unmarshal(l.Action.Condition, &and); err != nil || !scopedBelowForty(t, and.And.Vals) {
			t.Errorf("Lesser Heal line %d is not scoped below level 40", index)
		}
	}
	if lessers == 0 {
		t.Fatal("Discipline casts no Lesser Heal rank 3")
	}
}
