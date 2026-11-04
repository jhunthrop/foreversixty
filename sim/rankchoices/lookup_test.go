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
		// A row naming a preset still counts for a log, which does not say which preset.
		{"shaman-elemental", "Lightning Bolt", 60, []int{8}},
		// A race's row counts for any race (round 13): a log's rank is no mistake
		// where the site casts it for one race.
		{"priest-shadow", "Mind Blast", 45, []int{3, 4}},
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
		{"priest-discipline", "Power Word: Shield", 45, nil},
		{"priest-discipline", "Penance", 45, nil},
		// Renew is authored at rank 7 since round 15 wave 2; a row keeps rank 1 to 39.
		{"priest-discipline", "Renew", 25, []int{1}},
		{"paladin-holy", "Flash of Light", 35, []int{1}},
		// Round 16 wave 2: the tank's Flash of Light is rank 2 at 44 to 49.
		{"paladin-holy", "Flash of Light", 45, []int{1, 2}},
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

// Sim-accuracy round 16 wave 2: under the MaxLevel + 12 rule, Renew rank 1
// (MaxLevel 13) loses its bonus from 26, so the priests' tank Renew switches
// rank inside band 30 (level by level, 10000 iterations, seeds 1 and 50001).
func TestPriestRenewRowsSwitchInsideBandThirty(t *testing.T) {
	for _, c := range []struct {
		spec  string
		level int
		want  []int
	}{
		{"priest-holy", 31, []int{1}},
		{"priest-holy", 33, []int{3}},
		{"priest-holy", 39, []int{3}},
		{"priest-discipline", 29, []int{1}},
		{"priest-discipline", 30, []int{2}},
		{"priest-discipline", 34, []int{3}},
	} {
		got, err := ChosenRanks(c.spec, "Renew", c.level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, c.want) {
			t.Errorf("%s Renew at %d: %v, want %v", c.spec, c.level, got, c.want)
		}
	}
}

// Devouring Plague rank 1 (learned at 20) loses its bonus from 38; rank 2
// wins from 52 on both factions (round 16 wave 2).
func TestShadowDevouringPlagueMovesToRankTwoAtFiftyTwo(t *testing.T) {
	for level, want := range map[int][]int{51: {1, 2}, 52: {2}, 59: {2}} {
		got, err := ChosenRanks("priest-shadow", "Devouring Plague", level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("Devouring Plague at %d: %v, want %v", level, got, want)
		}
	}
}
