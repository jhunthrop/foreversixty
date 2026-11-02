package request

import (
	"strings"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Beast Mastery refreshes Serpent Sting only once less than one tick (3 s)
// is left. A refresh drops the sting's remaining ticks, and refreshing
// under 5 s clipped a tick a cast; 1.5 s measured +0.2% raid and +0.1 to
// +1.0% at the bare bands (round 9, physical lane).
func TestBeastMasteryRefreshesSerpentStingWithoutClippingATick(t *testing.T) {
	const serpentSting = 13555
	curated, err := loadLadderCurated(repoRoot, "hunter-beast-mastery")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	found := false
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != serpentSting {
			continue
		}
		found = true
		for _, v := range item.GetAction().GetCondition().GetAnd().GetVals() {
			cmp := v.GetCmp()
			if cmp.GetLhs().GetDotRemainingTime() == nil {
				continue
			}
			d, err := time.ParseDuration(strings.TrimSpace(cmp.GetRhs().GetConst().GetVal()))
			if err != nil {
				t.Fatal(err)
			}
			if d >= 3*time.Second {
				t.Errorf("Serpent Sting refreshes with %v left, at or above a tick's 3 s", d)
			}
		}
	}
	if !found {
		t.Fatal("Beast Mastery casts no Serpent Sting")
	}
}
