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
		{"Beast Mastery at 40 keeps the default", "hunter-beast-mastery", 40, 13552, 1978},
	} {
		ids := rotationIDs(t, c.spec, "hunter", c.level, specRankChoices(c.spec, BarePreset, "night-elf"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
