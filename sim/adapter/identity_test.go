package adapter

import (
	"fmt"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/wowsims/classic/sim/core/proto"
)

func spellID(id int32, tag, rank int32) *proto.ActionID {
	return &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: id}, Tag: tag, Rank: rank}
}

// A plain spell keeps its own id, because that is what the web resolves
// a name with. Everything else is derived, and no derived id can be
// mistaken for a client spell id.
func TestActionNameIdentifiesEveryKind(t *testing.T) {
	cases := []struct {
		name string
		id   *proto.ActionID
		want string
		raw  bool // the row id is the client id itself
	}{
		{"plain spell", spellID(25286, 0, 0), "spell:25286", true},
		{"tagged spell", spellID(25286, 1, 0), "spell:25286/1", false},
		{"ranked spell", spellID(11597, 0, 5), "spell:11597+r5", false},
		{"item", &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 5513}}, "item:5513", false},
		{"other", &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionAttack}, Tag: 2}, "other:attack/2", false},
		{"nil", nil, "unknown", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, name := ActionName(tc.id)
			if name != tc.want {
				t.Errorf("name = %q, want %q", name, tc.want)
			}
			if tc.raw {
				if got != 25286 {
					t.Errorf("row id = %d, want the spell's own id", got)
				}
				return
			}
			if got < syntheticBase {
				t.Errorf("row id = %d; a derived id must sit in the reserved space at or above %d so it cannot be read as a spell id", got, syntheticBase)
			}
		})
	}
}

// The three engine actions a white swing produces share one spell id
// and differ only by tag. They must be three rows, not one key three
// times, and none of them may take a real spell's id.
func TestDerivedIDsAreInjective(t *testing.T) {
	seen := map[int64]string{}
	for kind := kindSpell; kind <= kindUnknown; kind++ {
		for _, raw := range []int64{0, 1, 7, 25286, 1_271_953, kindStride - 1} {
			for _, tag := range []int32{0, 1, 2, 3, 99} {
				for _, rank := range []int32{0, 1, 5, rankStride - 1} {
					id := derivedID(kind, raw, tag, rank)
					at := fmt.Sprintf("kind=%d raw=%d tag=%d rank=%d", kind, raw, tag, rank)
					if prev, ok := seen[id]; ok {
						t.Fatalf("id %d is produced by both %s and %s", id, prev, at)
					}
					seen[id] = at
					if id < syntheticBase {
						t.Fatalf("%s derives %d, below the reserved space", at, id)
					}
				}
			}
		}
	}
}

// An engine result carries one ActionMetrics per tag, and each must
// reach the report as its own row.
func TestSummarizeSeparatesTaggedActions(t *testing.T) {
	u := oneAction()
	for _, tag := range []int32{1, 2} {
		u.Actions = append(u.Actions, &proto.ActionMetrics{
			Id:      spellID(23894, tag, 0),
			IsMelee: true,
			Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 100, Hits: 100, Damage: 1000}},
		})
	}
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	ids := map[int64]bool{}
	for _, ab := range got.DamageDone[0].Abilities {
		if ids[ab.SpellID] {
			t.Fatalf("two ability rows share id %d", ab.SpellID)
		}
		ids[ab.SpellID] = true
	}
	if len(got.DamageDone[0].Abilities) != 3 {
		t.Fatalf("abilities = %d, want 3 (one untagged and two tagged)", len(got.DamageDone[0].Abilities))
	}
}

// The engine registers some auras twice and reports each registration
// separately. A real fight's summary has one track per aura, so the two
// are added together rather than given a difference they do not have.
func TestSummarizeFoldsRepeatedRows(t *testing.T) {
	u := oneAction()
	u.Auras = append(u.Auras, &proto.AuraMetrics{
		Id:               spellID(12966, 0, 0),
		UptimeSecondsAvg: 10,
		ProcsAvg:         2,
	})
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Auras) != 1 {
		t.Fatalf("Auras = %d, want 1; two metrics for one aura are one track", len(got.Auras))
	}
	if got.Auras[0].UptimeMS != 150400 { // 140400 + 10000
		t.Errorf("UptimeMS = %d, want 150400 (the two uptimes added)", got.Auras[0].UptimeMS)
	}
	if got.Auras[0].Applications != 34 { // 32 + 2
		t.Errorf("Applications = %d, want 34", got.Auras[0].Applications)
	}
}

// The duplicate check is the backstop behind the derivation: a summary
// that somehow carried two rows on one key would break the report's
// keyed blocks, so it never leaves the adapter.
func TestCheckRowIdentityCatchesADuplicate(t *testing.T) {
	dup := summary.Summary{
		DamageDone: []summary.Actor{{
			GUID:      playerGUID,
			Abilities: []summary.Ability{{SpellID: 7}, {SpellID: 7}},
		}},
	}
	if err := checkRowIdentity(dup); err == nil {
		t.Fatal("two ability rows on one key passed the check")
	}
	ok := summary.Summary{
		DamageDone: []summary.Actor{{
			GUID:      playerGUID,
			Abilities: []summary.Ability{{SpellID: 7}, {SpellID: 8}},
		}},
	}
	if err := checkRowIdentity(ok); err != nil {
		t.Fatalf("a summary with distinct rows was rejected: %v", err)
	}
}
