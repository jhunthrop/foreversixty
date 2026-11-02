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
		{"shaman-elemental", "Lightning Bolt", 60, []int{4}},
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
