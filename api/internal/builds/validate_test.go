package builds

import (
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func fixture(t *testing.T) *trees.Data {
	t.Helper()
	d, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	return d
}

// Every rule gets a passing case and a failing case against the two-tree
// fixture class: Warrior (1) with Arms (101 Improved Heroic Strike max 5,
// 102 Deflection max 5, 103 Improved Rend tier 1 needing 5 in 101, 104 Deep
// Wounds tier 2) and Fury (201 Booming Voice, 202 Cruelty tier 1).
func TestValidate(t *testing.T) {
	data := fixture(t)
	longTitle := "0123456789012345678901234567890123456789012345678901234567890" // 61 characters

	for _, tc := range []struct {
		name string
		in   Input
		want map[string]string
	}{
		{
			name: "rule 0 pass: a known tree version",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1"},
		},
		{
			name: "rule 0 fail: an unknown tree version",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "nope"},
			want: map[string]string{"tree_version": "No talent data for tree version nope"},
		},
		{
			name: "rule 1 pass: a legal race and class pair",
			in:   Input{ClassID: 2, RaceID: 3, TreeVersion: "test-1"},
		},
		{
			name: "rule 1 fail: unknown class",
			in:   Input{ClassID: 9, RaceID: 1, TreeVersion: "test-1"},
			want: map[string]string{"class_id": "Unknown class 9"},
		},
		{
			name: "rule 1 fail: unknown race",
			in:   Input{ClassID: 1, RaceID: 7, TreeVersion: "test-1"},
			want: map[string]string{"race_id": "Unknown race 7"},
		},
		{
			name: "rule 1 fail: illegal combination",
			in:   Input{ClassID: 2, RaceID: 1, TreeVersion: "test-1"},
			want: map[string]string{"race_id": "Human Paladin is not a legal combination"},
		},
		{
			name: "rule 2 pass: talents of the class",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{101, 201}},
		},
		{
			name: "rule 2 fail: a talent of another class",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: []int{301}},
			want: map[string]string{"point_order[0]": "Talent 301 is not a Warrior talent"},
		},
		{
			name: "rule 3 pass: tier 1 after five points in the same tree",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{201, 201, 201, 201, 201, 202}},
		},
		{
			name: "rule 3 fail: tier 1 with four points in the tree",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{201, 201, 201, 201, 202}},
			want: map[string]string{"point_order[4]": "Tier 1 of Fury needs 5 points in Fury first"},
		},
		{
			name: "rule 4 pass: prerequisite filled first",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{101, 101, 101, 101, 101, 103}},
		},
		{
			name: "rule 4 fail: prerequisite short by one rank",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{101, 101, 101, 101, 102, 103}},
			want: map[string]string{"point_order[5]": "Improved Rend needs 5 points in Improved Heroic Strike first"},
		},
		{
			name: "rule 5 pass: exactly max rank",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{102, 102, 102, 102, 102}},
		},
		{
			name: "rule 5 fail: one rank over max",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{102, 102, 102, 102, 102, 102}},
			want: map[string]string{"point_order[5]": "Deflection has only 5 ranks"},
		},
		{
			name: "rule 6 pass: gear in the right slots",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				Gear: map[string]int{"head": 12640, "finger1": 19325, "finger2": 19325, "trinket1": 12930}},
		},
		{
			name: "rule 6 fail: unknown slot",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Gear: map[string]int{"helm": 12640}},
			want: map[string]string{"gear.helm": "Unknown gear slot helm"},
		},
		{
			name: "rule 6 fail: item the class cannot use",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Gear: map[string]int{"head": 999}},
			want: map[string]string{"gear.head": "Item 999 is not available to Warrior"},
		},
		{
			name: "rule 6 fail: item in the wrong slot",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Gear: map[string]int{"feet": 12640}},
			want: map[string]string{"gear.feet": "Lionheart Helm cannot go in the feet slot"},
		},
		{
			name: "rule 6 fail: the same unique trinket twice",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				Gear: map[string]int{"trinket1": 12930, "trinket2": 12930}},
			want: map[string]string{"gear.trinket2": "Only one Briarwood Reed can be equipped"},
		},
		{
			name: "title pass: sixty characters",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Title: longTitle[:60]},
		},
		{
			name: "title fail: sixty-one characters",
			in:   Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Title: longTitle},
			want: map[string]string{"title": "Title is at most 60 characters"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := Validate(data, tc.in)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("fields = %#v, want %#v", got, tc.want)
			}
		})
	}
}

// The 51-point cap gets its own test: no order of 52 points exists in the
// two-tree fixture that breaks only that rule (the fixture has 26 points of
// talents in total), so the table's exact-map comparison cannot express it.
func TestValidateRejectsMoreThanFiftyOnePoints(t *testing.T) {
	order := make([]int, MaxPoints+1)
	for i := range order {
		order[i] = 101
	}
	got := Validate(fixture(t), Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", PointOrder: order})
	if got["point_order"] != "A build has at most 51 points" {
		t.Fatalf("point_order = %q, want the 51-point cap message (all fields: %#v)", got["point_order"], got)
	}
}

func TestValidateAcceptsAFullyLoadedBuild(t *testing.T) {
	got := Validate(fixture(t), Input{
		ClassID: 1, RaceID: 3, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 101, 101, 101, 103, 201},
		Gear:       map[string]int{"head": 12640, "finger1": 19325, "trinket1": 12930},
		Title:      "Arms leveling",
	})
	if got != nil {
		t.Fatalf("fields = %#v, want nil", got)
	}
}
