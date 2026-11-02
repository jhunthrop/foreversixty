package request

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Subtlety's Ghostly Strike line says "on cooldown", but Hemorrhage (35
// energy) fired the moment energy reached 35, so energy never pooled to
// Ghostly Strike's 40 and it was cast 1.2 to 1.4 times a fight instead of
// about 9. Hemorrhage now holds back while Ghostly Strike is about to be
// ready, unless energy is high (round 9, physical lane).
func TestSubtletyHemorrhageWaitsForGhostlyStrike(t *testing.T) {
	const hemorrhage = 16511
	curated, err := loadLadderCurated(repoRoot, "rogue-subtlety")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != hemorrhage {
			continue
		}
		cond := protojson.Format(item.GetAction().GetCondition())
		if !strings.Contains(cond, "spellTimeToReady") || !strings.Contains(cond, "14278") {
			t.Errorf("Hemorrhage's condition does not wait for Ghostly Strike (14278): %s", cond)
		}
		return
	}
	t.Fatal("Subtlety casts no Hemorrhage")
}
