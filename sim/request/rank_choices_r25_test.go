package request

import "testing"

// Round 25 (platform): below 20 the band-20 entry used to run hunters with no
// bow below 20, and Serpent Sting and Arcane Shot never fired below 18 (their
// gate on Multi-Shot's cooldown dropped out with Multi-Shot unlearned), so no
// hunter row below 20 could be judged. Judged level by level with a bow in
// hand (spec-breakdown -level, bare band 20, 10,000 iterations, seeds 1 and
// 50001, Troll and Dwarf): rank 1 Serpent Sting wins at 18 and 19. (Its Arcane
// Shot rows are gone since the 10 November hotfix: hotfix_arcane_shot_rank_test.go.)
func TestRound25LowLevelHunterRankRows(t *testing.T) {
	checkHunterRankCases(t, []hunterRankCase{
		{"MM Serpent Sting at 18", hMM, BarePreset, "dwarf", 18, 1978, 13550},
	})
}

// Round 25 (platform): the Feral caster rows below 20 were judged unarmed and
// in quest gear the level could not have yet. Re-judged on the gear the level
// holds (same method, Tauren and Night Elf): Wrath rank 2 beats rank 1 at 12
// and 13, and for a Night Elf the rank 3 learned at 14 through 15 (a Tauren
// takes rank 3); Moonfire rank 2 wins from 13 with it, and at 11 on a Night Elf.
func TestRound25FeralCasterRankRows(t *testing.T) {
	const feral = "druid-feral"
	cases := []struct {
		name, race string
		level      int
		want, not  int32
	}{
		{"Wrath at 11 keeps rank 1", "tauren", 11, 5176, 5177},
		{"Wrath at 12", "tauren", 12, 5177, 5176},
		{"Wrath at 15 on a Night Elf", "night-elf", 15, 5177, 5178},
		{"Wrath at 14 on a Tauren takes rank 3", "tauren", 14, 5178, 5177},
		{"Wrath at 16 takes rank 3", "tauren", 16, 5178, 5177},
		{"Moonfire at 10 keeps rank 1", "night-elf", 10, 8921, 8924},
		{"Moonfire at 11 on a Tauren keeps rank 1", "tauren", 11, 8921, 8924},
		{"Moonfire at 11 on a Night Elf", "night-elf", 11, 8924, 8921},
		{"Moonfire at 12 keeps rank 1", "night-elf", 12, 8921, 8924},
		{"Moonfire at 13", "tauren", 13, 8924, 8921},
		{"Moonfire at 18 keeps rank 2", "tauren", 18, 8924, 8925},
	}
	for _, c := range cases {
		ids := rotationIDs(t, feral, "druid", c.level, specRankChoices(feral, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
