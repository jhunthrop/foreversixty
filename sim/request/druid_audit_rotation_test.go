package request

import (
	"encoding/json"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Sim-accuracy round 16 (design/reviews/2026-11-04-sim-audit-round16/druid.md).

const (
	catRipID           = 9896
	catFerociousBiteID = 22829
	// catRipRefreshAtMost is the latest a 5-point Rip may be refreshed: a new
	// Rip replaces the old one and its remaining ticks are lost, so the cat
	// lets it run to its last tick (2 s apart).
	catRipRefreshAtMost = 2 * time.Second
)

// finisherThreshold is the dotRemainingTime constant of the first
// top-level line casting spellID in spec's curated rotation.
func finisherThreshold(t *testing.T, spec string, spellID int32) time.Duration {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatalf("%s: %v", spec, err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != spellID {
			continue
		}
		for _, val := range item.GetAction().GetCondition().GetAnd().GetVals() {
			cmp := val.GetCmp()
			if cmp.GetLhs().GetDotRemainingTime() == nil {
				continue
			}
			d, err := time.ParseDuration(cmp.GetRhs().GetConst().GetVal())
			if err != nil {
				t.Fatalf("%s: line %d: %v", spec, spellID, err)
			}
			return d
		}
	}
	t.Fatalf("%s: no castSpell %d line with a dotRemainingTime condition", spec, spellID)
	return 0
}

// Rip at 5 points is recast only in its last tick: refreshing it with 6 s
// left threw away 1.2 ticks a cast (4.8 of 6 landed) and cost the cat
// about 1.6% at raid 60.
func TestCatLetsRipRunToItsLastTick(t *testing.T) {
	rip := finisherThreshold(t, "druid-feral", catRipID)
	if rip > catRipRefreshAtMost {
		t.Errorf("druid-feral refreshes Rip with %v left, want at most %v", rip, catRipRefreshAtMost)
	}
	if bite := finisherThreshold(t, "druid-feral", catFerociousBiteID); bite < rip {
		t.Errorf("druid-feral Bites with Rip at %v or more, below Rip's own %v refresh", bite, rip)
	}
}

// Wave 2: Balance's rank rows judged level by level (10,000 iterations,
// seeds 1 and 50001). Moonfire rank 2 through the 20s and rank 4 from 36;
// Wrath rank 3 at 22; Insect Swarm rank 2 at 39. Round 20 ends the Wrath
// rank-3 row at 23 and the Moonfire rank-2 row at 28
// (TestBalanceRankRowsTwentyToTwentyNine), so 29 is rank 3 here now; round
// 24 starts the rank-4 row at 35, so the rank-3 check sits at 34.
func TestBalanceRankRowsByLevel(t *testing.T) {
	const (
		moonfire2, moonfire3, moonfire4, moonfire5 = 8924, 8925, 8926, 8927
		wrath3, wrath4                             = 5178, 5179
		insectSwarm1, insectSwarm2                 = 5570, 24974
	)
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Moonfire at 23 is rank 2", 23, moonfire2, moonfire4},
		{"Moonfire at 29 is rank 3", 29, moonfire3, moonfire5},
		{"Moonfire at 34 is rank 3", 34, moonfire3, moonfire4},
		{"Moonfire at 36 is rank 4", 36, moonfire4, moonfire3},
		{"Wrath at 22 is rank 3", 22, wrath3, wrath4},
		{"Wrath at 27 is rank 4", 27, wrath4, wrath3},
		{"Insect Swarm at 38 is rank 1", 38, insectSwarm1, insectSwarm2},
		{"Insect Swarm at 39 is rank 2", 39, insectSwarm2, insectSwarm1},
	} {
		ids := rotationIDs(t, "druid-balance", "druid", c.level, specRankChoices("druid-balance", BarePreset, "night-elf"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// Sim-accuracy round 18 (design/reviews/2026-11-04-sim-audit-round18/druid.md
// R1): Restoration's Healing Touch filler (line 9, the tank under 55%) was
// never judged by level. Rank 3 beats the top rank by 1.2 to 1.5% at 30 to
// 33 and rank 5 by 1.1 to 1.3% at 40 to 44 (10,000 iterations, seeds 1 and
// 50001, both factions). Round 19: the levels between and after gain beyond
// error too, and rank rows take no 1% bar: rank 4 beats the top rank by 0.6
// to 0.9% at 34 to 39, rank 5 by 1.1% at 45 and rank 6 by 0.8 to 0.9% at 46
// to 49; the top rank wins from 50. Round 24's second pass, on the talents
// each level has: rank 6 ties at 48 and the top rank wins at 49. The Nature's
// Swiftness line (4) keeps the top rank.
func TestRestorationHealingTouchFillerRanksByLevel(t *testing.T) {
	const (
		healingTouch3, healingTouch4, healingTouch5 = 5187, 5188, 5189
		healingTouch6, healingTouch7, healingTouch8 = 6778, 8903, 9758
		healingTouch9                               = 9888
	)
	for _, c := range []struct {
		name      string
		level     int
		want, not []int32
	}{
		{"30: rank 3 filler beside the top-rank save", 30, []int32{healingTouch3, healingTouch5}, nil},
		{"33: rank 3 filler beside the top-rank save", 33, []int32{healingTouch3, healingTouch6}, nil},
		{"34: rank 4 filler beside the top-rank save", 34, []int32{healingTouch4, healingTouch6}, []int32{healingTouch3}},
		{"39: rank 4 filler beside the top-rank save", 39, []int32{healingTouch4, healingTouch7}, []int32{healingTouch5}},
		{"40: rank 5 filler beside the top-rank save", 40, []int32{healingTouch5, healingTouch7}, nil},
		{"45: rank 5 filler beside the top-rank save", 45, []int32{healingTouch5, healingTouch8}, []int32{healingTouch6}},
		{"46: rank 6 filler beside the top-rank save", 46, []int32{healingTouch6, healingTouch8}, []int32{healingTouch5}},
		{"48: rank 6 filler beside the top-rank save", 48, []int32{healingTouch6, healingTouch8}, nil},
		{"49: top rank only (round 24 second pass)", 49, []int32{healingTouch8}, []int32{healingTouch6}},
		{"50: top rank only", 50, []int32{healingTouch9}, []int32{healingTouch6}},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, "night-elf"))
		for _, id := range c.want {
			if !ids[id] {
				t.Errorf("%s: want Healing Touch %d cast", c.name, id)
			}
		}
		for _, id := range c.not {
			if ids[id] {
				t.Errorf("%s: want no Healing Touch %d", c.name, id)
			}
		}
	}
}

// Round 18 R2: the party Rejuvenation (authored 8910) is rank 2 at 40 and 41,
// where it beats rank 1 by 1.4 to 2.1% (night elf) and 0.8 to 1.4% (tauren).
// Round 19: rank 2 also wins at 42 (+0.8% night elf, +0.3% tauren, beyond
// error at both seeds; rank rows take no 1% bar); rank 1 again at 39 and 43.
func TestRestorationPartyRejuvenationAtFortyToFortyTwo(t *testing.T) {
	const rejuvenation1, rejuvenation2 = 774, 1058
	for _, c := range []struct {
		level     int
		want, not int32
	}{
		{39, rejuvenation1, 0},
		{40, rejuvenation2, rejuvenation1},
		{41, rejuvenation2, rejuvenation1},
		{42, rejuvenation2, rejuvenation1},
		{43, rejuvenation1, rejuvenation2},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, "night-elf"))
		if !ids[c.want] || (c.not != 0 && ids[c.not]) {
			t.Errorf("level %d: want Rejuvenation %d cast and %d not", c.level, c.want, c.not)
		}
	}
}

// Sim-accuracy round 20: the Healing Touch filler (line 9) judged at 20 to 26,
// below round 18's sweep. Rank 2 (learned 8) beats the rank 4 learned at 20 by
// 0.7% (night elf) and 1.8% (tauren) at 20; a tauren keeps it to 23 (+1.1,
// +0.8, +0.3%), a night elf casts the top rank from 21; both from 24.
// The Nature's Swiftness line casts the top rank at every level, so only rank
// 2's presence or absence is the filler's to assert.
func TestRestorationHealingTouchFillerAtTwenty(t *testing.T) {
	const healingTouch2 = 5186
	for _, c := range []struct {
		race  string
		level int
		rank2 bool
	}{
		{"night-elf", 20, true},
		{"night-elf", 21, false},
		{"tauren", 20, true},
		{"tauren", 23, true},
		{"tauren", 24, false},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, c.race))
		if ids[healingTouch2] != c.rank2 {
			t.Errorf("%s at %d: Healing Touch rank 2 (%d) cast = %v, want %v", c.race, c.level, healingTouch2, ids[healingTouch2], c.rank2)
		}
	}
}

// Sim-accuracy round 20: the party Rejuvenation re-judged at 20 to 39. Rank 2
// beats rank 1 at 27 (+0.4% night elf, +0.1% tauren), 28 (+0.4%, +0.1%) and
// 29 (+0.3%, +0.3%), every seed; rank 1 wins again at 30 (-1.2%). Round 24's
// second pass moved the row's start to 26 (rank 2 +0.1 to 0.2%), with a race
// split at 25 (TestDruidRankRowsSecondPassRound24).
func TestRestorationPartyRejuvenationAtTwentySevenToTwentyNine(t *testing.T) {
	const rejuvenation1, rejuvenation2 = 774, 1058
	// Rank 2 (1058) is also the tank's rank through the 20s and 30s, so only
	// rank 1 (774), which no other line casts, carries a negative check.
	for _, c := range []struct {
		level int
		want  int32
		not   int32
	}{
		{24, rejuvenation1, 0},
		{26, rejuvenation2, rejuvenation1},
		{27, rejuvenation2, rejuvenation1},
		{29, rejuvenation2, rejuvenation1},
		{30, rejuvenation1, 0},
	} {
		for _, race := range []string{"night-elf", "tauren"} {
			ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, race))
			if !ids[c.want] || (c.not != 0 && ids[c.not]) {
				t.Errorf("%s at %d: want party Rejuvenation %d cast and %d not", race, c.level, c.want, c.not)
			}
		}
	}
}

// Sim-accuracy round 19 (design/reviews/2026-11-04-sim-audit-round19/druid.md).
// The tank Rejuvenation (authored 9841) judged level by level at 40-59 under
// the low-rank penalty (10,000 iterations, seeds 1 and 50001, both factions):
// rank 4 wins at 40-43, rank 5 at 44-48, rank 6 at 49-56 (it beats the top
// rank by 1.0 to 2.7% at 53-56, where round 16 left the top rank) and rank 8
// at 57-59.
// Round 20: at 38 and 39, on the band-30 gear those levels wear, rank 2 beats
// rank 4 by 1.7 to 1.8% and 1.0 to 1.3%, so the rank-4 row starts at 40.
func TestRestorationTankRejuvenationRanksByLevel(t *testing.T) {
	rejuvenation := map[int]int32{2: 1058, 4: 2090, 5: 2091, 6: 3627, 8: 9839, 9: 9840}
	for _, c := range []struct {
		level int
		rank  int
	}{
		{38, 2}, {39, 2}, {40, 4}, {43, 4}, {44, 5}, {48, 5}, {49, 6}, {53, 6}, {56, 6}, {57, 8}, {59, 8},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, "night-elf"))
		for rank, id := range rejuvenation {
			// Rank 2 (1058) is also the party line's rank at 40 to 42
			// (TestRestorationPartyRejuvenationAtFortyToFortyTwo), so its
			// absence is only the tank's to assert below 40.
			if rank == 2 && c.level >= 40 {
				continue
			}
			if got, want := ids[id], rank == c.rank; got != want {
				t.Errorf("level %d: Rejuvenation rank %d (%d) cast = %v, want %v", c.level, rank, id, got, want)
			}
		}
	}
}

// Sim-accuracy round 20 (design/reviews/2026-11-04-sim-audit-round20/druid.md
// L1, L2): the band-20 Balance rows re-judged level by level (10,000
// iterations, seeds 1 and 50001, night elf and tauren). Wrath rank 3 (learned
// 14) beats the rank 4 learned at 22 by 1.1% at 22 and ties it at 23; from 24
// rank 4 wins by 1.3% (24), 2.1% (25) and 2.5% (26), so round 16's rank-3 row
// ends at 23, not 26. Moonfire's rank-2 row was re-judged with the new Wrath
// row in: with Wrath rank 4 from 24, rank 2 beats rank 3 by 2.1% at 24 down
// to 0.5% at 28, and rank 3 wins by 0.3 to 0.4% at 29 (full-precision DPS),
// so the rank-2 row ends at 28. Judged alone, with the old Wrath row, rank 3
// had looked ahead at 25 and 26: the two rows move together, so a row is
// judged with its neighbours' rows in.
func TestBalanceRankRowsTwentyToTwentyNine(t *testing.T) {
	const (
		wrath3, wrath4       = 5178, 5179
		moonfire2, moonfire3 = 8924, 8925
	)
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Wrath at 22 is rank 3", 22, wrath3, wrath4},
		{"Wrath at 23 is rank 3", 23, wrath3, wrath4},
		{"Wrath at 24 is rank 4", 24, wrath4, wrath3},
		{"Wrath at 26 is rank 4", 26, wrath4, wrath3},
		{"Wrath at 29 is rank 4", 29, wrath4, wrath3},
		{"Moonfire at 24 is rank 2", 24, moonfire2, moonfire3},
		{"Moonfire at 25 is rank 2", 25, moonfire2, moonfire3},
		{"Moonfire at 26 is rank 3", 26, moonfire3, moonfire2},
		{"Moonfire at 27 is rank 2", 27, moonfire2, moonfire3},
		{"Moonfire at 28 is rank 2", 28, moonfire2, moonfire3},
		{"Moonfire at 29 is rank 3", 29, moonfire3, moonfire2},
		{"Moonfire at 30 is rank 3", 30, moonfire3, moonfire2},
	} {
		for _, race := range []string{"night-elf", "tauren"} {
			ids := rotationIDs(t, "druid-balance", "druid", c.level, specRankChoices("druid-balance", BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s (%s): want %d cast and %d not", c.name, race, c.want, c.not)
			}
		}
	}
}

// Sim-accuracy round 20: the tank Rejuvenation's rank-2 row (20 to 37)
// re-judged at 20, 23, 26, 28, 29, 30, 33, 36 and 39. It holds everywhere but
// 29, where a night elf heals 0.1 to 0.3% more with rank 3 on every one of
// four seeds and a tauren is level, so 29 is a race-keyed split: rank 3 for a
// night elf, rank 2 for every other race.
func TestRestorationTankRejuvenationAtTwentyNine(t *testing.T) {
	const rejuvenation2, rejuvenation3 = 1058, 1430
	for _, c := range []struct {
		race      string
		level     int
		want, not int32
	}{
		{"night-elf", 28, rejuvenation2, rejuvenation3},
		// At 29 the party line casts rank 2 too (its 27 to 29 row), so
		// only rank 3's presence is the tank's to assert.
		{"night-elf", 29, rejuvenation3, 0},
		{"night-elf", 30, rejuvenation2, rejuvenation3},
		{"tauren", 29, rejuvenation2, rejuvenation3},
		{"high-order-skyborne", 29, rejuvenation2, rejuvenation3},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, c.race))
		if !ids[c.want] || (c.not != 0 && ids[c.not]) {
			t.Errorf("%s at %d: want Rejuvenation %d cast and %d not", c.race, c.level, c.want, c.not)
		}
	}
}

// Round 19: Balance at 40-59 judged level by level (10,000 iterations, seeds
// 1 and 50001, night elf and tauren). Wrath rank 6 beats rank 7 by 0.1 to
// 0.7% at 46 to 48 (tie at 49); Insect Swarm rank 3 beats rank 4 by 0.1 to
// 0.6% at 54 to 58 (tie at 53 and 59). Moonfire and Starfire keep the top
// rank at every level checked.
func TestBalanceRankRowsFortyToFiftyNine(t *testing.T) {
	const (
		wrath6, wrath7             = 6780, 8905
		insectSwarm3, insectSwarm4 = 24975, 24976
	)
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Wrath at 45 is rank 6", 45, wrath6, wrath7},
		{"Wrath at 46 is rank 6", 46, wrath6, wrath7},
		{"Wrath at 47 is rank 6", 47, wrath6, wrath7},
		{"Wrath at 49 is rank 7", 49, wrath7, wrath6},
		{"Insect Swarm at 53 is rank 4", 53, insectSwarm4, insectSwarm3},
		{"Insect Swarm at 54 is rank 3", 54, insectSwarm3, insectSwarm4},
		{"Insect Swarm at 56 is rank 3", 56, insectSwarm3, insectSwarm4},
		{"Insect Swarm at 57 is rank 4", 57, insectSwarm4, insectSwarm3},
		{"Insect Swarm at 58 is rank 4", 58, insectSwarm4, insectSwarm3},
	} {
		for _, race := range []string{"night-elf", "tauren"} {
			ids := rotationIDs(t, "druid-balance", "druid", c.level, specRankChoices("druid-balance", BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s (%s): want %d cast and %d not", c.name, race, c.want, c.not)
			}
		}
	}
}

// Sim-accuracy round 24: the rank sweep below 20, judged level by level with
// Wrath and Moonfire together (band-20 gear, 10,000 iterations, seeds 1 and
// 50001, night elf and tauren). Re-judged in round 26 on the corrected
// below-band gear: Wrath rank 1 at 10 to 13 and 15 to 16, rank 2 at 14 and
// 17 to 18 (rank 3 from 19); Moonfire rank 1 at 12 to 13 and 15 to 17, rank 2
// at 14 and 18 to 19. Moonfire rank 4 also wins at 35, a level earlier than
// round 16's row (round24/druid.md).
func TestBalanceRankRowsBelowTwentyAndMoonfireAtThirtyFive(t *testing.T) {
	const (
		wrath1, wrath2, wrath3          = 5176, 5177, 5178
		moonfire1, moonfire2, moonfire3 = 8921, 8924, 8925
		moonfire4                       = 8926
	)
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Wrath at 10 is rank 1", 10, wrath1, wrath2},
		{"Wrath at 12 is rank 1", 12, wrath1, wrath2},
		{"Wrath at 13 is rank 1", 13, wrath1, wrath2},
		{"Wrath at 14 is rank 2", 14, wrath2, wrath3},
		{"Wrath at 18 is rank 2", 18, wrath2, wrath3},
		{"Wrath at 19 is rank 3", 19, wrath3, wrath2},
		{"Moonfire at 15 is rank 1", 15, moonfire1, moonfire2},
		{"Moonfire at 16 is rank 1", 16, moonfire1, moonfire3},
		{"Moonfire at 17 is rank 1", 17, moonfire1, moonfire3},
		{"Moonfire at 19 is rank 2", 19, moonfire2, moonfire3},
		{"Moonfire at 20 is rank 2", 20, moonfire2, moonfire1},
		{"Moonfire at 34 is rank 3", 34, moonfire3, moonfire4},
		{"Moonfire at 35 is rank 4", 35, moonfire4, moonfire3},
	} {
		for _, race := range []string{"night-elf", "tauren"} {
			ids := rotationIDs(t, "druid-balance", "druid", c.level, specRankChoices("druid-balance", BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s (%s): want %d cast and %d not", c.name, race, c.want, c.not)
			}
		}
	}
}

// Sim-accuracy round 24: Healing Touch rank 5 arrives at 26, and a tauren's
// hole filler heals 0.2 to 0.5% more with rank 4 at 26 to 28 (both seeds); a
// night elf ties and casts the top rank.
func TestRestorationTaurenHealingTouchRankFourAtTwentySixToTwentyEight(t *testing.T) {
	const healingTouch4, healingTouch5 = 5188, 5189
	for _, c := range []struct {
		race      string
		level     int
		want, not int32
	}{
		{"tauren", 26, healingTouch4, 0},
		{"tauren", 28, healingTouch4, 0},
		{"night-elf", 27, healingTouch5, healingTouch4},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, c.race))
		if !ids[c.want] || (c.not != 0 && ids[c.not]) {
			t.Errorf("%s at %d: want Healing Touch %d cast and %d not", c.race, c.level, c.want, c.not)
		}
	}
}

// Sim-accuracy round 24: the feral card's caster lines before Cat Form, swept
// below 20 with Wrath and Moonfire together; Moonfire rank 2 at 16 to 19 over
// the rank 3 learned at 16. Round 25 re-judged 10 to 15 with the gear the
// level can hold (round 24's runs were unarmed below 18): Wrath rank 1 only
// to 11, rank 2 at 12 and 13 (to 15 for a Night Elf); Moonfire rank 1 at 10
// to 12 (11 for a Tauren, TestRound25FeralCasterRankRows), rank 2 from 13.
func TestFeralCasterRankRowsBelowTwenty(t *testing.T) {
	const (
		wrath1, wrath2, wrath3          = 5176, 5177, 5178
		moonfire1, moonfire2, moonfire3 = 8921, 8924, 8925
	)
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Wrath at 10 is rank 1", 10, wrath1, wrath2},
		{"Wrath at 13 is rank 2", 13, wrath2, wrath1},
		{"Wrath at 16 is rank 3", 16, wrath3, wrath2},
		{"Moonfire at 10 is rank 1", 10, moonfire1, moonfire2},
		{"Moonfire at 12 is rank 1", 12, moonfire1, moonfire2},
		{"Moonfire at 13 is rank 2", 13, moonfire2, moonfire1},
		{"Moonfire at 14 is rank 2", 14, moonfire2, moonfire1},
		{"Moonfire at 16 is rank 2", 16, moonfire2, moonfire3},
		{"Moonfire at 19 is rank 2", 19, moonfire2, moonfire3},
	} {
		for _, race := range []string{"night-elf", "tauren"} {
			ids := rotationIDs(t, "druid-feral", "druid", c.level, specRankChoices("druid-feral", BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s (%s): want %d cast and %d not", c.name, race, c.want, c.not)
			}
		}
	}
}

// Sim-accuracy round 24: Restoration below 20, the four heals judged
// together (band-20 gear and heal profile, 10,000 iterations, night elf and
// tauren). The fight runs dry, so the cheap ranks win. Re-judged in round 26
// on the corrected below-band gear (night elf and tauren) and in round 27 for
// the Skyborne races on their own: the tank and party Rejuvenation ranks
// below by race, and from 17 to 19 rank 1 of the tank and party
// Rejuvenation, Healing Touch and (from 18) Regrowth.
func TestRestorationRankRowsBelowTwenty(t *testing.T) {
	healingTouch := map[int]int{1: 5185, 2: 5186, 3: 5187, 4: 5188}
	regrowth := map[int]int{1: 8936, 2: 8938}
	tank := linesCasting(t, "druid-restoration", 9841)
	party := linesCasting(t, "druid-restoration", 8910)
	regrowthLine := linesCasting(t, "druid-restoration", 9858)
	// level -> rank, tank then party, per race.
	ranks := map[string][2]map[int]int{
		"night-elf":           {{10: 1, 12: 2, 13: 1, 16: 1, 19: 1}, {10: 1, 12: 2, 13: 2, 15: 1, 16: 1, 17: 1}},
		"tauren":              {{10: 1, 12: 2, 13: 1, 16: 1, 19: 1}, {10: 2, 13: 2, 15: 1, 16: 3, 17: 1}},
		"high-order-skyborne": {{10: 2, 12: 1, 13: 1, 16: 1, 19: 1}, {10: 2, 12: 2, 13: 2, 15: 1, 16: 1, 17: 1}},
		"windshaper-skyborne": {{10: 2, 12: 2, 13: 1, 16: 1, 19: 1}, {10: 2, 13: 1, 14: 1, 15: 1, 16: 1, 17: 1}},
	}
	for race, lines := range ranks {
		for i, line := range []json.RawMessage{tank, party} {
			var cases []rankCase
			for level, rank := range lines[i] {
				cases = append(cases, rankCase{level, BarePreset, rank})
			}
			assertRanks(t, line, "druid-restoration", "druid", race, rejuvenationRank, cases)
		}
	}
	for _, race := range []string{"night-elf", "tauren"} {
		// The Nature's Swiftness line keeps the top rank, so only the filler's
		// own rank is asserted present. Healing Touch at 14 to 16 splits by
		// race (TestDruidRankRowsSecondPassRound24).
		for _, c := range []rankCase{{17, BarePreset, 1}, {19, BarePreset, 1}} {
			ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, race))
			if !ids[int32(healingTouch[c.rank])] {
				t.Errorf("%s at %d: want Healing Touch rank %d cast", race, c.level, c.rank)
			}
		}
		assertRanks(t, regrowthLine, "druid-restoration", "druid", race, regrowth, []rankCase{
			{17, BarePreset, 1}, {18, BarePreset, 1}, {19, BarePreset, 1},
		})
	}
}

// Sim-accuracy round 24, second pass: every druid row judged on band-floor
// talents re-judged with spec-breakdown -level N (the talents the player has
// at that level) and a forced -rank, 10,000 iterations, seeds 1 and 50001,
// night elf and tauren (round24/druid.md, "Second pass"). Moonfire rank 3
// wins at 26 (rank 2 either side); Insect Swarm rank 4 wins at 57 and 58; the
// tank Rejuvenation keeps rank 1 to 22; the party Rejuvenation splits at 25
// (rank 2 for a night elf, rank 1 for a tauren) and takes rank 2 from 26;
// Healing Touch rank 1 at 14 to 16 for a night elf, rank 2 for a tauren; and
// the top Healing Touch at 49.
func TestDruidRankRowsSecondPassRound24(t *testing.T) {
	const moonfire2, moonfire3 = 8924, 8925
	for _, race := range []string{"night-elf", "tauren"} {
		for _, c := range []struct {
			level     int
			want, not int32
		}{{25, moonfire2, moonfire3}, {26, moonfire3, moonfire2}, {27, moonfire2, moonfire3}} {
			ids := rotationIDs(t, "druid-balance", "druid", c.level, specRankChoices("druid-balance", BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("Balance %s at %d: want Moonfire %d cast and %d not", race, c.level, c.want, c.not)
			}
		}
	}
	tank := linesCasting(t, "druid-restoration", 9841)
	party := linesCasting(t, "druid-restoration", 8910)
	for _, race := range []string{"night-elf", "tauren"} {
		assertRanks(t, tank, "druid-restoration", "druid", race, rejuvenationRank, []rankCase{
			{20, BarePreset, 1}, {22, BarePreset, 1}, {23, BarePreset, 2}, {28, BarePreset, 2},
		})
		assertRanks(t, party, "druid-restoration", "druid", race, rejuvenationRank, []rankCase{
			{24, BarePreset, 1}, {26, BarePreset, 2},
		})
	}
	for race, rank := range map[string]int{"night-elf": 2, "high-order-skyborne": 2, "tauren": 1, "windshaper-skyborne": 1} {
		assertRanks(t, party, "druid-restoration", "druid", race, rejuvenationRank, []rankCase{{25, BarePreset, rank}})
	}
	healingTouch := map[int]int32{1: 5185, 2: 5186, 6: 6778, 8: 9758}
	for _, c := range []struct {
		race      string
		level     int
		want, not int
	}{
		// Round 26: rank 2 at 14 for every race (corrected below-band gear).
		{"night-elf", 14, 2, 1}, {"night-elf", 16, 1, 2}, {"high-order-skyborne", 15, 1, 2},
		// Round 27: a Windshaper Skyborne takes rank 1 at 15 and 16, as a night elf does.
		{"tauren", 14, 2, 1}, {"tauren", 16, 2, 1}, {"windshaper-skyborne", 15, 1, 2},
		// The Nature's Swiftness line casts the top rank (8) at 48 too.
		{"night-elf", 48, 6, 0}, {"night-elf", 49, 8, 6}, {"tauren", 49, 8, 6},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, c.race))
		if !ids[healingTouch[c.want]] || (c.not != 0 && ids[healingTouch[c.not]]) {
			t.Errorf("%s at %d: want Healing Touch rank %d cast and rank %d not", c.race, c.level, c.want, c.not)
		}
	}
}
