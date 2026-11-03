package request

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// Every warrior keeps Battle Shout up. The bare preset below 60 carries no
// raid Battle Shout, so a rotation that never shouts sims a warrior with no
// shout at all: the Protection list had none, and its band numbers lost 2.5
// to 3.7 percent threat to it (round 11, auras lane). The raid preset's
// shout is the same aura, so at 60 nothing moves.
func TestEveryWarriorRotationShouts(t *testing.T) {
	const battleShout = 11551
	for _, spec := range []string{"warrior-arms", "warrior-fury", "warrior-protection"} {
		curated, err := loadLadderCurated(repoRoot, spec)
		if err != nil {
			t.Fatal(err)
		}
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
			t.Fatal(err)
		}
		prepull, upkeep := false, false
		for _, action := range rot.PrepullActions {
			if action.GetAction().GetCastSpell().GetSpellId().GetSpellId() == battleShout {
				prepull = true
			}
		}
		for _, item := range rot.PriorityList {
			if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() == battleShout {
				upkeep = true
			}
		}
		if !prepull || !upkeep {
			t.Errorf("%s: Battle Shout prepull %v, upkeep line %v; want both", spec, prepull, upkeep)
		}
	}
}

// The pre-pull shout is skipped when the aura is already up. The raid
// preset carries a raid warrior's Battle Shout (the same aura, and in
// Forever it reaches the whole raid), so an unconditional pre-pull cast
// spent 10 of the opening Bloodrage's rage on nothing: Fury raid 60 about
// -0.2% (round 12, regression-melee lane; the landing's Bloodrage at -2.5 s
// is what first gave the old -1 s shout the rage to cast). The bare bands
// have no raid shout, so they still shout before the pull.
func TestWarriorPrepullShoutSkipsAnAuraAlreadyUp(t *testing.T) {
	const battleShout = 11551
	for _, spec := range []string{"warrior-arms", "warrior-fury", "warrior-protection"} {
		curated, err := loadLadderCurated(repoRoot, spec)
		if err != nil {
			t.Fatal(err)
		}
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
			t.Fatal(err)
		}
		for _, action := range rot.PrepullActions {
			if action.GetAction().GetCastSpell().GetSpellId().GetSpellId() != battleShout {
				continue
			}
			gate := action.GetAction().GetCondition().GetNot().GetVal().GetAuraIsActive()
			if gate.GetAuraId().GetSpellId() != battleShout {
				t.Errorf("%s: the pre-pull Battle Shout is not gated on its aura being down", spec)
			}
		}
	}
}
