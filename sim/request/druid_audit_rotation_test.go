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
