package sims

import (
	"reflect"
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

func TestAnAddonGearCodeBecomesTheEnginesCharacter(t *testing.T) {
	got, err := characterFromCode("Obnoxious",
		"FS1:1.60.1.70291:warrior:human:371000000/0a/0:main_hand=23171:1900,chest=6627|level=24|professions=mining")
	if err != nil {
		t.Fatal(err)
	}
	want := simapi.CharacterSpec{
		Name: "Obnoxious", Race: "human", Class: "warrior", Level: 24,
		// One decimal digit per talent, trees dash-joined, trailing zeros
		// trimmed; a base-36 digit past 9 is clamped the way the web's
		// talentsString clamps it.
		Talents: "371-09-",
		Gear: []simapi.GearSlot{
			{Slot: "chest", ItemID: 6627},
			{Slot: "main_hand", ItemID: 23171, Enchant: 1900},
		},
		Buffs: []string{}, Consumes: []string{}, Profession: []string{"mining"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("character = %+v\nwant        %+v", got, want)
	}
}

func TestAGearCodeWithNoLevelIsSimmedAtItsTalentLevel(t *testing.T) {
	got, err := characterFromCode("Obnoxious", "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627")
	if err != nil {
		t.Fatal(err)
	}
	// Eleven points spent: a level-20 character, as the planner reads it.
	if got.Level != 20 {
		t.Fatalf("level = %d, want 20", got.Level)
	}
}

func TestAGearCodeThatIsNotOneCannotBeSimmed(t *testing.T) {
	if _, err := characterFromCode("x", "hello"); err == nil {
		t.Fatal("want an error")
	}
}
