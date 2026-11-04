package request

import (
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Round 15 mage (guide-accuracy round 1 SIM-2): Arcane Power and Presence
// of Mind go out after the tank's first Sunders, not on the pull. The
// rotation's own note says "On a fresh pull, give the tank a Sunder or two
// first"; the guide's opener blasts twice first. Measured at a tie.
func TestArcaneHoldsArcanePowerForTheTanksFirstSunders(t *testing.T) {
	const (
		arcanePower    = 12042
		presenceOfMind = 12043
		minDelay       = 3 * time.Second
	)
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	seen := map[int32]bool{}
	for _, item := range rot.PriorityList {
		id := item.GetAction().GetCastSpell().GetSpellId().GetSpellId()
		if id != arcanePower && id != presenceOfMind {
			continue
		}
		seen[id] = true
		cmp := item.GetAction().GetCondition().GetCmp()
		if cmp == nil || cmp.GetOp() != proto.APLValueCompare_OpGe || cmp.GetLhs().GetCurrentTime() == nil {
			t.Errorf("spell %d: want a currentTime >= condition, got %v", id, item.GetAction().GetCondition())
			continue
		}
		d, err := time.ParseDuration(cmp.GetRhs().GetConst().GetVal())
		if err != nil || d < minDelay {
			t.Errorf("spell %d waits %q, want at least %v", id, cmp.GetRhs().GetConst().GetVal(), minDelay)
		}
	}
	if !seen[arcanePower] || !seen[presenceOfMind] {
		t.Errorf("Arcane must cast Arcane Power and Presence of Mind, saw %v", seen)
	}
}
