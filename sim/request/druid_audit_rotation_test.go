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
