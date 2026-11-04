package rankchoices

import (
	"slices"
	"testing"
)

// The coach reads the same table the rotation casts from, so a rank the site chooses is
// never called a mistake (round 10, platform).
func TestChosenRanksReadsTheCuratedRowsAtALevel(t *testing.T) {
	for _, c := range []struct {
		spec, spell string
		level       int
		want        []int
	}{
		{"paladin-retribution", "Consecration", 35, []int{1}},
		{"paladin-retribution", "Consecration", 59, []int{1}},
		{"paladin-retribution", "Consecration", 25, nil},
		{"druid-balance", "Moonfire", 30, []int{3}},
		// Level 60's rows name a preset; a log does not say which, so either counts.
		{"shaman-elemental", "Lightning Bolt", 60, []int{4, 7}},
		// A race's row counts for any race (round 13): a log's rank is no mistake
		// where the site casts it for one race.
		{"priest-shadow", "Mind Blast", 55, []int{2, 3}},
		{"warrior-fury", "Bloodthirst", 60, nil},
		{"paladin-retribution", "Seal of Righteousness", 40, nil},
	} {
		got, err := ChosenRanks(c.spec, c.spell, c.level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, c.want) {
			t.Errorf("%s %s at %d: %v, want %v", c.spec, c.spell, c.level, got, c.want)
		}
	}
}

// Sim-accuracy round 12 (healprofile): the sub-60 healer rows re-measured on the logged
// tank-and-group split, each beyond error at seeds 1 and 50001 for both factions
// (design/reviews/2026-11-03-sim-audit-round12/healprofile.md).
func TestRoundTwelveHealerRowsBelowSixty(t *testing.T) {
	for _, c := range []struct {
		spec, spell string
		level       int
		want        []int
	}{
		{"priest-discipline", "Power Word: Shield", 25, nil},
		{"priest-discipline", "Power Word: Shield", 35, nil},
		{"priest-discipline", "Power Word: Shield", 45, []int{6}},
		{"priest-discipline", "Penance", 45, nil},
		{"priest-discipline", "Renew", 25, []int{1}},
		{"paladin-holy", "Flash of Light", 35, []int{1}},
		{"paladin-holy", "Flash of Light", 45, []int{1, 3}},
	} {
		got, err := ChosenRanks(c.spec, c.spell, c.level)
		if err != nil {
			t.Fatal(err)
		}
		slices.Sort(got)
		if !slices.Equal(got, c.want) {
			t.Errorf("%s %s at %d: %v, want %v", c.spec, c.spell, c.level, got, c.want)
		}
	}
}
