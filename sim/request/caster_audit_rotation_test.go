package request

import (
	"slices"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// The 2026-10-31 caster theorycraft audit
// (design/reviews/2026-10-31-theorycraft-audit/casters.md) found rotation
// lines that disagreed with the build or with Forever theorycraft. These
// pin the corrected lines so a later rotation search cannot quietly undo
// them without a decision.

// castOrder lists the spell ids of a curated rotation's top-level castSpell
// lines, in priority order.
func castOrder(t *testing.T, spec string) []int32 {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatalf("%s: %v", spec, err)
	}
	var ids []int32
	for _, item := range rot.PriorityList {
		if cast := item.GetAction().GetCastSpell(); cast != nil {
			ids = append(ids, cast.GetSpellId().GetSpellId())
		}
		if channel := item.GetAction().GetChannelSpell(); channel != nil {
			ids = append(ids, channel.GetSpellId().GetSpellId())
		}
	}
	return ids
}

func requireBefore(t *testing.T, spec string, ids []int32, first, second int32) {
	t.Helper()
	i, j := slices.Index(ids, first), slices.Index(ids, second)
	if i < 0 || j < 0 {
		t.Fatalf("%s: want both %d and %d cast, got %v", spec, first, second, ids)
	}
	if i > j {
		t.Errorf("%s: %d must sit above %d, got %v", spec, first, second, ids)
	}
}

// W1, W2: the Affliction build takes Siphon Life, so the rotation keeps it
// up; Bane of Agony, not Bane of Doom, is the maintained curse.
func TestAfflictionKeepsSiphonLifeAndAgony(t *testing.T) {
	const (
		siphonLife = 18881
		agony      = 11713
		doom       = 603
	)
	ids := castOrder(t, "warlock-affliction")
	if !slices.Contains(ids, siphonLife) {
		t.Errorf("Affliction never casts Siphon Life (%d): %v", siphonLife, ids)
	}
	if !slices.Contains(ids, agony) || slices.Contains(ids, doom) {
		t.Errorf("Affliction must keep Bane of Agony (%d), not Bane of Doom (%d): %v", agony, doom, ids)
	}
}

// D2: Demonic Pact is real and modeled. The Demonology raid build takes it,
// and the prepull summons the Imp, sacrifices it, then summons the Succubus
// before Soul Link; Shadow Bolt is the filler under both Shadow bonuses.
func TestDemonologySacrificesTheImpUnderDemonicPact(t *testing.T) {
	const (
		summonImp         = 688
		demonicSacrifice  = 18788
		summonSuccubus    = 712
		soulLink          = 19028
		shadowBolt        = 11661
		searingPain       = 17923
		demonologySpecKey = "warlock-demonology"
	)
	curated, err := loadLadderCurated(repoRoot, demonologySpecKey)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	var prepull []int32
	for _, p := range rot.PrepullActions {
		prepull = append(prepull, p.GetAction().GetCastSpell().GetSpellId().GetSpellId())
	}
	want := []int32{summonImp, demonicSacrifice, summonSuccubus, soulLink}
	if !slices.Equal(prepull, want) {
		t.Errorf("Demonology prepull = %v, want %v", prepull, want)
	}
	ids := castOrder(t, demonologySpecKey)
	if !slices.Contains(ids, shadowBolt) || slices.Contains(ids, searingPain) {
		t.Errorf("Demonology's filler must be Shadow Bolt (%d), not Searing Pain (%d): %v", shadowBolt, searingPain, ids)
	}
}
