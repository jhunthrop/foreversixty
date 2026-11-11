package request

import "testing"

// rankRowCase is one level of a rank row: the ids the rewritten rotation must
// cast and the ids it must not.
type rankRowCase struct {
	name, race string
	level      int
	want, not  []int32
}

func checkRankRows(t *testing.T, spec, class string, cases []rankRowCase) {
	t.Helper()
	for _, c := range cases {
		ids := rotationIDs(t, spec, class, c.level, specRankChoices(spec, BarePreset, c.race))
		for _, id := range c.want {
			if !ids[id] {
				t.Errorf("%s %s: %d not cast", spec, c.name, id)
			}
		}
		for _, id := range c.not {
			if ids[id] {
				t.Errorf("%s %s: %d cast", spec, c.name, id)
			}
		}
	}
}

// Round 26 (ranks): the caster and healer rows below 20 were judged on
// below-band gear with empty armor slots and the other faction's quest
// rewards. Re-judged on the corrected gear (spec-breakdown -level N -rank,
// band 20, 10,000 iterations, seeds 1 and 50001, both factions).
func TestRound26BalanceRankRows(t *testing.T) {
	const wr1, wr2, wr3, mf1, mf2, mf3 = 5176, 5177, 5178, 8921, 8924, 8925
	checkRankRows(t, "druid-balance", "druid", []rankRowCase{
		{"11 keeps Wrath 1, Moonfire 2", "tauren", 11, []int32{wr1, mf2}, []int32{wr2, mf1}},
		{"12 Moonfire 1", "night-elf", 12, []int32{wr1, mf1}, []int32{mf2}},
		{"13 Wrath 1, Moonfire 1", "tauren", 13, []int32{wr1, mf1}, []int32{wr2, mf2}},
		{"14 keeps Wrath 2, Moonfire 2", "night-elf", 14, []int32{wr2, mf2}, []int32{wr1, wr3, mf1}},
		{"16 Wrath 1, Moonfire 1", "night-elf", 16, []int32{wr1, mf1}, []int32{wr2, mf2, mf3}},
		{"17 keeps Wrath 2, Moonfire 1", "tauren", 17, []int32{wr2, mf1}, []int32{wr1, mf2}},
		{"18 Wrath 2, Moonfire 2", "tauren", 18, []int32{wr2, mf2}, []int32{wr3, mf1, mf3}},
		{"19 Wrath 3, Moonfire 2", "night-elf", 19, []int32{wr3, mf2}, []int32{wr2, mf1, mf3}},
	})
}

func TestRound26RestorationDruidRankRows(t *testing.T) {
	const rj1, rj2, rj3, ht1, ht2 = 774, 1058, 1430, 5185, 5186
	checkRankRows(t, "druid-restoration", "druid", []rankRowCase{
		{"10 tauren tank rank 1, party rank 2", "tauren", 10, []int32{rj1, rj2, ht2}, []int32{ht1}},
		{"10 night elf rank 1 of everything", "night-elf", 10, []int32{rj1, ht1}, []int32{rj2}},
		{"11 keeps rank 2 Rejuvenation", "night-elf", 11, []int32{rj2}, []int32{rj1}},
		{"12 keeps rank 2 Rejuvenation", "tauren", 12, []int32{rj2}, []int32{rj1}},
		{"13 tank rank 1, party rank 2 (tauren)", "tauren", 13, []int32{rj1, rj2}, nil},
		{"13 tank rank 1, party rank 2 (night elf)", "night-elf", 13, []int32{rj1, rj2}, nil},
		{"14 Healing Touch 2", "night-elf", 14, []int32{rj1, rj2, ht2}, []int32{ht1}},
		{"15 Rejuvenation 1 everywhere (tauren)", "tauren", 15, []int32{rj1, ht2}, []int32{rj2, ht1}},
		{"15 night elf Healing Touch 1", "night-elf", 15, []int32{rj1, ht1}, []int32{rj2, ht2}},
		{"16 tauren tank rank 1, party rank 3", "tauren", 16, []int32{rj1, rj3, ht2}, []int32{rj2}},
		{"16 night elf rank 1", "night-elf", 16, []int32{rj1, ht1}, []int32{rj3}},
	})
}

func TestRound26ProtectionHolyStrikeTakesRankTwo(t *testing.T) {
	checkRankRows(t, "paladin-protection", "paladin", []rankRowCase{
		{"dwarf 12", "dwarf", 12, []int32{678}, []int32{679}},
		{"undead 14", "undead", 14, []int32{678}, []int32{679}},
		{"undead 15", "undead", 15, []int32{678}, []int32{679}},
	})
}

func TestRound26HolyPriestRenewAtSixteen(t *testing.T) {
	const rn1, rn2 = 139, 6074
	checkRankRows(t, "priest-holy", "priest", []rankRowCase{
		{"15 keeps rank 1", "human", 15, []int32{rn1}, []int32{rn2}},
		{"16 rank 1 (round 30)", "undead", 16, []int32{rn1}, []int32{rn2}},
		{"17 rank 1", "human", 17, []int32{rn1}, []int32{rn2}},
	})
}

func TestRound26DisciplineLesserHealAtSeventeen(t *testing.T) {
	const lh2, lh3 = 2052, 2053
	checkRankRows(t, "priest-discipline", "priest", []rankRowCase{
		{"16 keeps rank 3", "human", 16, []int32{lh3}, []int32{lh2}},
		{"17 rank 2", "undead", 17, []int32{lh2}, []int32{lh3}},
		{"18 rank 3", "human", 18, []int32{lh3}, []int32{lh2}},
	})
}

func TestRound26DisciplineShieldRankRows(t *testing.T) {
	const pw1, pw2, pw3 = 17, 592, 600
	checkRankRows(t, "priest-discipline", "priest", []rankRowCase{
		{"12 rank 1", "undead", 12, []int32{pw1}, []int32{pw2}},
		{"13 human keeps rank 2", "human", 13, []int32{pw2}, []int32{pw1}},
		{"13 troll keeps rank 2", "troll", 13, []int32{pw2}, []int32{pw1}},
		{"14 rank 1", "troll", 14, []int32{pw1}, []int32{pw2}},
		{"15 dwarf rank 1", "dwarf", 15, []int32{pw1}, []int32{pw2}},
		{"15 undead keeps rank 2", "undead", 15, []int32{pw2}, []int32{pw1}},
		{"16 rank 1", "human", 16, []int32{pw1}, []int32{pw2}},
		{"17 keeps rank 2", "human", 17, []int32{pw2}, []int32{pw1}},
		{"18 human rank 1", "human", 18, []int32{pw1}, []int32{pw3}},
		{"19 night elf keeps rank 3", "night-elf", 19, []int32{pw3}, []int32{pw1}},
	})
}
