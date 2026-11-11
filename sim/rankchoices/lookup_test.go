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
		{"paladin-retribution", "Consecration", 52, []int{1}},
		// Round 24 second pass: rank 3 at 53 to 59 for every race (Holy Conduit 2/2).
		{"paladin-retribution", "Consecration", 53, []int{3}},
		{"paladin-retribution", "Consecration", 59, []int{3}},
		{"paladin-retribution", "Consecration", 25, nil},
		{"druid-balance", "Moonfire", 30, []int{3}},
		// A row naming a preset still counts for a log, which does not say which preset.
		{"shaman-elemental", "Lightning Bolt", 60, []int{8}},
		// A race's row counts for any race (round 13): a log's rank is no mistake
		// where the site casts it for one race.
		// Round 24: rank 4 for most races at 42, rank 5 for a Gnome or Undead.
		{"priest-shadow", "Mind Blast", 42, []int{4, 5}},
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
		// Round 20: the tank's Flash of Light is rank 3 at 50 to 54 and rank 4 at 55 to 59 and bare 60.
		{"paladin-holy", "Flash of Light", 52, []int{3}},
		// Round 22: the raid-60 party row adds rank 5, so 60 carries both.
		{"paladin-holy", "Flash of Light", 60, []int{4, 5}},
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
		// Round 24 (on the talents at each level): rank 2 and 3 take over race by race.
		{"priest-holy", 30, []int{1, 2}}, // round 30: rank 2 for the Alliance races, rank 1 for a Troll or Undead
		{"priest-holy", 31, []int{2}},
		{"priest-holy", 33, []int{2}},
		{"priest-holy", 39, []int{3}},
		{"priest-discipline", 29, []int{1}},
		{"priest-discipline", 30, []int{2}},
		{"priest-discipline", 34, []int{2}}, // round 24: rank 2 holds through 34
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
		{"priest-holy", 50, []int{5}},
		{"priest-holy", 51, []int{5, 6}}, // round 19: rank 6 for an Undead
		{"priest-holy", 52, []int{5, 6}}, // round 19: rank 5 for a Gnome, Dwarf or Night Elf
		{"priest-holy", 53, []int{6}},
		{"priest-holy", 58, []int{6}}, // round 24: a Human casts the default rank 7 at 58
		{"priest-holy", 59, nil},      // round 24: rank 7 (the default) for every race
		{"priest-discipline", 40, []int{4}},
		{"priest-discipline", 43, []int{4}},
		{"priest-discipline", 44, []int{5}}, // round 24: rank 5 for a Dwarf, Gnome, Night Elf or Undead
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
// race but a Human, who keeps rank 1 there (round 19).
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
		{"Mind Blast", 40, []int{4, 5}}, // round 24: rank 5 for a Human, Gnome or Undead
		{"Mind Blast", 43, []int{4}},
		{"Mind Blast", 44, nil},
		{"Mind Blast", 49, nil},
		{"Devouring Plague", 40, []int{1}},
		{"Devouring Plague", 43, []int{1}},
		{"Devouring Plague", 44, []int{2}},
		{"Devouring Plague", 48, []int{2}},
		{"Devouring Plague", 49, []int{2, 3}}, // round 24: a Human casts rank 3 at 49
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
// Round 24: rank 1 holds at 58 and 59 too, when rank 4 arrives (+0.2% to +1.6%).
func TestPriestGreaterHealRankOneBelowSixty(t *testing.T) {
	for _, spec := range []string{"priest-holy", "priest-discipline"} {
		for level, want := range map[int][]int{45: nil, 46: {1}, 57: {1}, 58: {1}, 59: {1}, 60: nil} {
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

// Sim-accuracy round 19: Restoration shaman's Chain Heal line casts rank 1
// from 46 to 59. It costs 260 mana to rank 2's 315 and rank 3's 405, and a
// leveling healer is mana-bound, so the cheaper heal heals more over the
// fight (+0.3% at 46 to +1.2% at 59, 10,000 iterations, seeds 1 and 50001,
// both factions). Below 46 rank 1 is the only rank; at 60 the default
// stands for a raid (the bare-60 row is preset-scoped).
func TestRestorationShamanChainHealRankOneWhileLeveling(t *testing.T) {
	for level, want := range map[int][]int{45: nil, 46: {1}, 53: {1}, 59: {1}} {
		got, err := ChosenRanks("shaman-restoration", "Chain Heal", level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("Chain Heal at %d: %v, want %v", level, got, want)
		}
	}
}

// Sim-accuracy round 19: with Chain Heal rank 1 through 59, Restoration's
// tank Healing Wave steps up to rank 7 at 56 (+0.4% to +0.7% over rank 6,
// both factions, seeds 1 and 50001), and Elemental's filler Lightning Bolt
// casts rank 3 at 26 and 27, where rank 5, just learned, costs more than
// the fight can pay (+3.7% to +4.9%). Round 24 re-judged the bolt on the
// talents a player has at each level: rank 3 wins from 20 to 33 (at 25 by
// 0.3% to 4.6% over rank 4, at 28 and 29 by 2.1% to 2.9% over rank 4 and
// 5.4% to 6.7% over the top rank 5), and rank 4 from 34.
func TestRoundNineteenShamanRowsByLevel(t *testing.T) {
	for _, tc := range []struct {
		spec, spell string
		level       int
		want        []int
	}{
		{"shaman-restoration", "Healing Wave", 55, []int{6}},
		{"shaman-restoration", "Healing Wave", 56, []int{7}},
		{"shaman-restoration", "Healing Wave", 58, []int{7}},
		{"shaman-elemental", "Lightning Bolt", 25, []int{3}},
		{"shaman-elemental", "Lightning Bolt", 26, []int{3}},
		{"shaman-elemental", "Lightning Bolt", 27, []int{3}},
		{"shaman-elemental", "Lightning Bolt", 28, []int{3}},
		{"shaman-elemental", "Lightning Bolt", 29, []int{3}},
		{"shaman-elemental", "Lightning Bolt", 34, []int{4}},
	} {
		got, err := ChosenRanks(tc.spec, tc.spell, tc.level)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, tc.want) {
			t.Errorf("%s %s at %d: %v, want %v", tc.spec, tc.spell, tc.level, got, tc.want)
		}
	}
}
