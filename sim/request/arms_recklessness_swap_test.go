package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Arms swaps to Berserker Stance in the execute phase only to press
// Recklessness, which a warrior trains at 50. The swap line read
// spellIsReady(Recklessness), and below 50 that value builds to nil and the
// engine drops a nil operand from an AND (newValueAnd), so at bare 30 and
// 40 Arms left Battle Stance for nothing: 2.5% and 3.6% DPS (round 12,
// regression-melee lane; the line is from S 99afd1a6b). spellIsKnown is
// false, not nil, for an unlearned spell, so the line now asks it first.
func TestArmsBerserkerSwapNeedsRecklessnessTrained(t *testing.T) {
	const berserkerStance, recklessness = 2458, 1719
	curated, err := loadLadderCurated(repoRoot, "warrior-arms")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	found := false
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != berserkerStance {
			continue
		}
		found = true
		known := false
		for _, v := range item.GetAction().GetCondition().GetAnd().GetVals() {
			if v.GetSpellIsKnown().GetSpellId().GetSpellId() == recklessness {
				known = true
			}
		}
		if !known {
			t.Errorf("the Berserker Stance swap does not require Recklessness to be known")
		}
	}
	if !found {
		t.Fatal("no Berserker Stance line in the Arms rotation")
	}
}
