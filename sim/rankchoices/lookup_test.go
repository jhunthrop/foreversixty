package rankchoices

import (
	"encoding/json"
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
		{"priest-shadow", "Mind Blast", 42, []int{4}},
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
		// Round 16 wave 2: the tank's Flash of Light is rank 2 at 44 to 48;
		// round 19: the group's is rank 2 at 44 to 46 as well.
		{"paladin-holy", "Flash of Light", 45, []int{2}},
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

// Sim-accuracy round 18: from 40 the tank's Renew was the list's rank 7 (or
// the highest learned) at every level, but a mana-bound healer at 40-59 heals
// more with a cheaper rank that still keeps its whole bonus: rank 4 to 44,
// rank 5 from 45, and for Holy rank 6 from 52 (level by level, 10,000
// iterations, seeds 1 and 50001, both factions).
func TestPriestRenewRowsBelowSixty(t *testing.T) {
	for _, c := range []struct {
		spec  string
		level int
		want  []int
	}{
		{"priest-holy", 40, []int{4}},
		{"priest-holy", 44, []int{4}},
		{"priest-holy", 45, []int{5}},
		{"priest-holy", 51, []int{5}},
		{"priest-holy", 52, []int{6}},
		{"priest-holy", 59, []int{6}},
		{"priest-discipline", 40, []int{4}},
		{"priest-discipline", 44, []int{4}},
		{"priest-discipline", 45, []int{5}},
		{"priest-discipline", 49, []int{5}},
		{"priest-discipline", 50, nil},
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
// wins from 52 on both factions (round 16 wave 2), and at 50 and 51 for every
// race since round 18.
func TestShadowDevouringPlagueMovesToRankTwoAtFiftyTwo(t *testing.T) {
	for level, want := range map[int][]int{51: {2}, 52: {2}, 59: {2}} {
		got, err := ChosenRanks("priest-shadow", "Devouring Plague", level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("Devouring Plague at %d: %v, want %v", level, got, want)
		}
	}
}

// Sim-accuracy round 18: judged level by level under the MaxLevel + 12 rule,
// Mind Blast rank 4 beats the top rank only to 43, and rank 3 never beats
// rank 4 for a Dwarf or a Night Elf. From 44 every race casts the top Mind
// Blast with Devouring Plague rank 2 (rank 1, learned at 20, has lost too
// much spell power by then).
func TestShadowRowsSwitchAtFortyFour(t *testing.T) {
	for _, c := range []struct {
		spell string
		level int
		want  []int
	}{
		{"Mind Blast", 40, []int{4}},
		{"Mind Blast", 43, []int{4}},
		{"Mind Blast", 44, nil},
		{"Mind Blast", 49, nil},
		{"Devouring Plague", 40, []int{1}},
		{"Devouring Plague", 43, []int{1}},
		{"Devouring Plague", 44, []int{2}},
		{"Devouring Plague", 49, []int{2}},
	} {
		got, err := ChosenRanks("priest-shadow", c.spell, c.level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, c.want) {
			t.Errorf("%s at %d: %v, want %v", c.spell, c.level, got, c.want)
		}
	}
}

// A band may carry its own lines and pre-pull scope (round 17); the coach still counts
// every band's rank at a level as a choice, whatever line it is scoped to.
func TestABandCarriesItsOwnLines(t *testing.T) {
	raw := []byte(`{"specs":{"x":[{"spell":"Holy Light","authored":7,"lines":[2],"bands":[
		{"from_level":20,"to_level":24,"rank":1},
		{"from_level":50,"to_level":59,"rank":2,"lines":[4],"prepull":[0]}]}]}}`)
	var file File
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	bands := file.Specs["x"][0].Bands
	if bands[0].Lines != nil || !slices.Equal(bands[1].Lines, []int{4}) || !slices.Equal(bands[1].Prepull, []int{0}) {
		t.Fatalf("bands = %+v, want the second band's own lines [4] and prepull [0], the first none", bands)
	}
	table, err := byName(raw)
	if err != nil {
		t.Fatal(err)
	}
	for level, want := range map[int][]int{22: {1}, 55: {2}, 40: nil} {
		if got := ranksAt(table, "x", "Holy Light", level); !slices.Equal(got, want) {
			t.Errorf("level %d = %v, want %v", level, got, want)
		}
	}
}

// Sim-accuracy round 18: the top-rank Greater Heal lines (10965) cast rank 1
// at 46-57, where a mana-bound healer heals more with it than with rank 2 or
// 3 (+0.2% to +1.2%, 10,000 iterations, seeds 1 and 50001, both factions).
func TestPriestGreaterHealRankOneBelowFiftyEight(t *testing.T) {
	for _, spec := range []string{"priest-holy", "priest-discipline"} {
		for level, want := range map[int][]int{45: nil, 46: {1}, 57: {1}, 58: nil} {
			got, err := ChosenRanks(spec, "Greater Heal", level)
			if err != nil {
				t.Fatal(err)
			}
			if !slices.Equal(got, want) {
				t.Errorf("%s Greater Heal at %d: %v, want %v", spec, level, got, want)
			}
		}
	}
}
