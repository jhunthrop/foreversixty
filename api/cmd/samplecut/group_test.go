package main

import (
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

func TestGroupOfKeepsThePlayersInTheLoggersGroup(t *testing.T) {
	us := []units.Unit{
		{GUID: "Player-1-0002", Kind: units.KindPlayer, Flags: 0x512},
		{GUID: "Player-1-0001", Kind: units.KindPlayer, Flags: 0x511},
		{GUID: "Player-1-0003", Kind: units.KindPlayer, Flags: 0x514},
		{GUID: "Player-1-0004", Kind: units.KindPlayer, Flags: 0x518},
		{GUID: "Pet-0-1-1-1-1-0001", Kind: units.KindPet, Flags: 0x1111},
		{GUID: "Creature-0-1-1-1-1-0001", Kind: units.KindCreature, Flags: 0xa48},
	}
	got := groupOf(us)
	want := []string{"Player-1-0001", "Player-1-0002", "Player-1-0003"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("groupOf = %v, want %v", got, want)
	}
}

func TestGroupOfIsEmptyNotNull(t *testing.T) {
	if got := groupOf(nil); got == nil || len(got) != 0 {
		t.Errorf("groupOf(nil) = %#v, want an empty slice", got)
	}
}
