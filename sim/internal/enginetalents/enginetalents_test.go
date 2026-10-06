package enginetalents

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/paladin"
)

// paladinTrees is a slice of the active build's paladin trees: enough
// talents, by their client names, to land in all three engine trees.
func paladinTrees() []leveling.TalentTree {
	return []leveling.TalentTree{
		{Position: 0, Talents: []leveling.TalentNode{{ID: 105639, Name: "Divine Strength", MaxRank: 5}}},
		{Position: 1, Talents: []leveling.TalentNode{{ID: 105630, Name: "Toughness", MaxRank: 5}}},
		{Position: 2, Talents: []leveling.TalentNode{
			{ID: 110882, Name: "Champion of the Light", Tier: 5, Column: 1, MaxRank: 3},
			{ID: 105692, Name: "Twist of Light", Tier: 6, Column: 1, MaxRank: 1},
		}},
	}
}

func engineDir(t *testing.T) string {
	t.Helper()
	dir, err := SourceDir(".")
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func paladinLayout(t *testing.T) Layout {
	t.Helper()
	l, err := ForClass(engineDir(t), "paladin")
	if err != nil {
		t.Fatal(err)
	}
	return l
}

// A renamed talent keeps its node id, and the id is what is matched:
// druid node 104949 is the engine's mangle field, Primal Bite in
// 1.60.1.70009.
func TestEncodeMatchesByIDNotName(t *testing.T) {
	l, err := ForClass(engineDir(t), "druid")
	if err != nil {
		t.Fatal(err)
	}
	f, err := l.FieldFor(1, leveling.TalentNode{ID: 104949, Name: "Primal Bite"})
	if err != nil {
		t.Fatal(err)
	}
	if f.Name != "mangle" {
		t.Fatalf("node 104949 -> %q, want mangle", f.Name)
	}
}

// The string Encode writes must read back, through the engine's own
// FillTalentsProto, as exactly the ranks it was given - whatever
// position each field happens to sit at in this engine's message.
func TestEncodeRoundTripsThroughTheEngine(t *testing.T) {
	l := paladinLayout(t)
	s, err := l.Encode(paladinTrees(), map[int]int{105639: 5, 105630: 3, 110882: 2, 105692: 1})
	if err != nil {
		t.Fatal(err)
	}
	got := &proto.PaladinTalents{}
	core.FillTalentsProto(got.ProtoReflect(), s, paladin.TalentTreeSizes)
	if got.DivineStrength != 5 || got.Toughness != 3 || got.ChampionOfTheLight != 2 || !got.TwistOfLight {
		t.Fatalf("engine read %q as %+v", s, got)
	}
}

func TestEncodeRefusesTwoPointsInABoolField(t *testing.T) {
	l := paladinLayout(t)
	if _, err := l.Encode(paladinTrees(), map[int]int{105692: 2}); err == nil {
		t.Fatal("2 points in a one-rank (bool) field encoded without error")
	}
}

func TestEncodeRefusesATalentTheEngineCannotSee(t *testing.T) {
	l := paladinLayout(t)
	trees := paladinTrees()
	trees[0].Talents = append(trees[0].Talents, leveling.TalentNode{ID: 1, Name: "Divine Strength", MaxRank: 1})
	_, err := l.Encode(trees, map[int]int{1: 1})
	if err == nil || !strings.Contains(err.Error(), "no engine field") {
		t.Fatalf("err = %v, want a no-engine-field error", err)
	}
}

func TestEncodeRefusesAnIDOutsideTheTrees(t *testing.T) {
	l := paladinLayout(t)
	if _, err := l.Encode(paladinTrees(), map[int]int{999: 1}); err == nil {
		t.Fatal("an id outside the trees encoded without error")
	}
}

func TestForClassRejectsAnUnknownClass(t *testing.T) {
	if _, err := ForClass(engineDir(t), "monk"); err == nil {
		t.Fatal("ForClass(monk) returned no error")
	}
}

func TestGoName(t *testing.T) {
	if got := (Field{Name: "one_handed_weapon_specialization"}).GoName(); got != "OneHandedWeaponSpecialization" {
		t.Fatalf("GoName = %q", got)
	}
}
