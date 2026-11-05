package request

import "testing"

// Round 22 (hunter lane, per-level sweep): since round 21 (A2) Aimed Shot no
// longer waits for an auto shot, so it takes the shared cooldown far more
// often and a Marksmanship hunter at 50 and 51 spends his mana on rank 4
// (20902, 210 mana) and runs dry for about 90 s of 180. Rank 1 (19434, 75
// mana) beats it by 4.8 to 5.2% for every hunter race at both levels
// (spec-breakdown -level, bare band 50); rounds 17 and 18 measured a tie
// there while Aimed Shot still waited for the auto shot. The rank 1 row now
// runs unbroken from 28 to 59.
func TestRound22MarksmanshipAimedShotRankOneAt50And51(t *testing.T) {
	for _, c := range []struct {
		name      string
		race      string
		level     int
		want, not int32
	}{
		{"Orc Marksmanship Aimed Shot at 50", "orc", 50, 19434, 20902},
		{"Orc Marksmanship Aimed Shot at 51", "orc", 51, 19434, 20902},
		{"Night Elf Marksmanship Aimed Shot at 50", "night-elf", 50, 19434, 20902},
		{"Night Elf Marksmanship Aimed Shot at 51", "night-elf", 51, 19434, 20902},
	} {
		ids := rotationIDs(t, "hunter-marksmanship", "hunter", c.level, specRankChoices("hunter-marksmanship", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Round 22 (hunter lane, every rank at 20 to 60, bare bands): Arcane Shot carries 15% of ranged
// attack power on every rank (Blizzard, 3 November), so a low rank keeps most of its hit for a
// fraction of the mana. Round 20 assumed the share could not flip a rank and never swept.
func TestRound22HunterArcaneShotRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec, preset string
		level              int
		want, not          int32
	}{
		{"Marksmanship at 20", "hunter-marksmanship", BarePreset, 20, 3044, 14282},
		{"Marksmanship at 29", "hunter-marksmanship", BarePreset, 29, 3044, 14283},
		{"Marksmanship at 30", "hunter-marksmanship", BarePreset, 30, 14282, 14283},
		{"Marksmanship at 39", "hunter-marksmanship", BarePreset, 39, 14282, 14284},
		{"Marksmanship at 40", "hunter-marksmanship", BarePreset, 40, 3044, 14284},
		{"Marksmanship at 59", "hunter-marksmanship", BarePreset, 59, 3044, 14286},
		{"Marksmanship at bare 60", "hunter-marksmanship", BarePreset, 60, 3044, 14287},
		{"Marksmanship raid 60 keeps the top rank", "hunter-marksmanship", RaidPreset, 60, 14287, 3044},
		{"Beast Mastery at 25 keeps the default", "hunter-beast-mastery", BarePreset, 25, 14282, 3044},
		{"Beast Mastery at 26", "hunter-beast-mastery", BarePreset, 26, 3044, 14282},
		{"Beast Mastery at 49", "hunter-beast-mastery", BarePreset, 49, 3044, 14285},
		{"Beast Mastery at 50", "hunter-beast-mastery", BarePreset, 50, 14282, 14285},
		{"Beast Mastery at 59", "hunter-beast-mastery", BarePreset, 59, 14282, 14286},
		{"Beast Mastery at bare 60", "hunter-beast-mastery", BarePreset, 60, 14282, 14287},
		{"Beast Mastery raid 60 keeps the top rank", "hunter-beast-mastery", RaidPreset, 60, 14287, 14282},
	} {
		ids := rotationIDs(t, c.spec, "hunter", c.level, specRankChoices(c.spec, c.preset, "orc"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Round 22: the same sweep for Serpent Sting (5% of ranged attack power a tick on every rank):
// rank 1 wins at 20 to 29 for Marksmanship and at 26 to 39 for Beast Mastery, the newest rank
// after that.
func TestRound22HunterSerpentStingRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec string
		level      int
		want, not  int32
	}{
		{"Marksmanship at 20", "hunter-marksmanship", 20, 1978, 13550},
		{"Marksmanship at 29", "hunter-marksmanship", 29, 1978, 13551},
		{"Marksmanship at 30 keeps the default", "hunter-marksmanship", 30, 13551, 1978},
		{"Beast Mastery at 26", "hunter-beast-mastery", 26, 1978, 13551},
		{"Beast Mastery at 39", "hunter-beast-mastery", 39, 1978, 13552},
		{"Beast Mastery at 40 keeps the default", "hunter-beast-mastery", 40, 13552, 1978},
	} {
		ids := rotationIDs(t, c.spec, "hunter", c.level, specRankChoices(c.spec, BarePreset, "night-elf"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
