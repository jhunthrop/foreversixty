package notify

import (
	"strings"
	"testing"
)

// The post reads the comp target into every role heading and lists the buffs nobody brings and
// those too few bring, so an officer counts nothing by hand.
func TestRaidPostCarriesTheTargetAndTheCover(t *testing.T) {
	in := onyxiaThursday()
	in.Roster = []RosterLine{
		{Name: "Ironvale", Class: "warrior", Spec: "Protection", Role: "tank", State: "in"},
		{Name: "Aurelith", Class: "priest", Spec: "Holy", Role: "healer", State: "in"},
	}
	in.Target = &RoleTarget{Tanks: 2, Healers: 8, Melee: 10, Ranged: 20}
	in.Missing, in.Short = []string{"Mark of the Wild", "Trueshot Aura"}, []string{"Blessing of Kings"}
	e := RaidPostEmbed(RaidPost{RaidPosted: in, State: PostOpen}, fixedFinish)

	heads := map[string]bool{}
	var buffs string
	for _, f := range e.Fields {
		heads[f.Name] = true
		if f.Name == raidCopy.BuffsName {
			buffs = f.Value
		}
	}
	for _, want := range []string{"Tanks · 1 of 2", "Healers · 1 of 8", "Melee · 0 of 10", "Ranged · 0 of 20"} {
		if !heads[want] {
			t.Errorf("no heading %q in %v", want, heads)
		}
	}
	if buffs != "Missing: Mark of the Wild, Trueshot Aura\nNeed more: Blessing of Kings" {
		t.Fatalf("buffs = %q", buffs)
	}
}

// With no target the headings keep their bare counts, and a covered raid has no Buffs field.
func TestRaidPostWithoutTargetOrGaps(t *testing.T) {
	in := onyxiaThursday()
	in.Roster = []RosterLine{{Name: "Ironvale", Class: "warrior", Spec: "Protection", Role: "tank", State: "in"}}
	e := RaidPostEmbed(RaidPost{RaidPosted: in, State: PostOpen}, fixedFinish)
	for _, f := range e.Fields {
		if f.Name == raidCopy.BuffsName || strings.Contains(f.Name, " of ") {
			t.Fatalf("field %q", f.Name)
		}
	}
}

// A group past its target says how far over, so three tanks against two never reads "3 of 2".
func TestRaidPostSaysOverTarget(t *testing.T) {
	in := onyxiaThursday()
	for _, n := range []string{"Ironvale", "Stonebrow", "Kraggor"} {
		in.Roster = append(in.Roster, RosterLine{Name: n, Class: "warrior", Spec: "Protection", Role: "tank", State: "in"})
	}
	in.Target = &RoleTarget{Tanks: 2, Healers: 8, Melee: 10, Ranged: 20}
	e := RaidPostEmbed(RaidPost{RaidPosted: in, State: PostOpen}, fixedFinish)
	for _, f := range e.Fields {
		if strings.HasPrefix(f.Name, "Tanks") && f.Name != "Tanks · 3 of 2, 1 over" {
			t.Fatalf("heading %q", f.Name)
		}
	}
}
