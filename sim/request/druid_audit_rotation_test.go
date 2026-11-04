package request

import (
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
// Wrath rank 3 at 22 to 26; Insect Swarm rank 2 at 39.
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
		{"Moonfire at 29 is rank 2", 29, moonfire2, moonfire5},
		{"Moonfire at 35 is rank 3", 35, moonfire3, moonfire4},
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
// to 49; the top rank wins from 50. The Nature's Swiftness line (4) keeps the
// top rank.
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
		{"49: rank 6 filler beside the top-rank save", 49, []int32{healingTouch6, healingTouch8}, nil},
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

// Sim-accuracy round 19 (design/reviews/2026-11-04-sim-audit-round19/druid.md).
// The tank Rejuvenation (authored 9841) judged level by level at 40-59 under
// the low-rank penalty (10,000 iterations, seeds 1 and 50001, both factions):
// rank 4 wins at 40-43, rank 5 at 44-48, rank 6 at 49-56 (it beats the top
// rank by 1.0 to 2.7% at 53-56, where round 16 left the top rank) and rank 8
// at 57-59.
func TestRestorationTankRejuvenationRanksByLevel(t *testing.T) {
	rejuvenation := map[int]int32{4: 2090, 5: 2091, 6: 3627, 8: 9839, 9: 9840}
	for _, c := range []struct {
		level int
		rank  int
	}{
		{43, 4}, {44, 5}, {48, 5}, {49, 6}, {53, 6}, {56, 6}, {57, 8}, {59, 8},
	} {
		ids := rotationIDs(t, "druid-restoration", "druid", c.level, specRankChoices("druid-restoration", BarePreset, "night-elf"))
		for rank, id := range rejuvenation {
			if got, want := ids[id], rank == c.rank; got != want {
				t.Errorf("level %d: Rejuvenation rank %d (%d) cast = %v, want %v", c.level, rank, id, got, want)
			}
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
		{"Wrath at 48 is rank 6", 48, wrath6, wrath7},
		{"Wrath at 49 is rank 7", 49, wrath7, wrath6},
		{"Insect Swarm at 53 is rank 4", 53, insectSwarm4, insectSwarm3},
		{"Insect Swarm at 54 is rank 3", 54, insectSwarm3, insectSwarm4},
		{"Insect Swarm at 58 is rank 3", 58, insectSwarm3, insectSwarm4},
	} {
		for _, race := range []string{"night-elf", "tauren"} {
			ids := rotationIDs(t, "druid-balance", "druid", c.level, specRankChoices("druid-balance", BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s (%s): want %d cast and %d not", c.name, race, c.want, c.not)
			}
		}
	}
}
