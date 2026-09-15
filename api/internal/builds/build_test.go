package builds

import "testing"

func TestCanonicalJSONSortsKeysAndExcludesTitle(t *testing.T) {
	in := Input{
		ClassID:     2,
		RaceID:      3,
		TreeVersion: "1.15.9.69722",
		PointOrder:  []int{1451, 1451, 1451, 1451, 1451, 1452},
		Gear:        map[string]int{"head": 12640, "finger1": 19325},
		Title:       "Holy leveling, 10 to 30",
	}
	got, err := CanonicalJSON(in)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"class_id":2,"gear":{"finger1":19325,"head":12640},"point_order":[1451,1451,1451,1451,1451,1452],"race_id":3,"tree_version":"1.15.9.69722"}`
	if string(got) != want {
		t.Fatalf("canonical json =\n%s\nwant\n%s", got, want)
	}
}

func TestIDMatchesTheFixedVectors(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   Input
		want string
	}{
		{
			name: "gear and points",
			in: Input{ClassID: 2, RaceID: 3, TreeVersion: "1.15.9.69722",
				PointOrder: []int{1451, 1451, 1451, 1451, 1451, 1452},
				Gear:       map[string]int{"head": 12640, "finger1": 19325}},
			want: "dbsmoe6j",
		},
		{
			name: "empty build",
			in:   Input{ClassID: 2, RaceID: 3, TreeVersion: "1.15.9.69722"},
			want: "jotol22p",
		},
		{
			name: "fixture build",
			in: Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
				PointOrder: []int{101, 101, 101, 101, 101, 201}},
			want: "znorjmts",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ID(tc.in)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Fatalf("id = %q, want %q", got, tc.want)
			}
			if len(got) != 8 {
				t.Fatalf("id %q is %d characters, want 8", got, len(got))
			}
		})
	}
}

func TestIDIgnoresTheTitleAndTheGearKeyOrder(t *testing.T) {
	base := Input{ClassID: 2, RaceID: 3, TreeVersion: "1.15.9.69722"}
	titled := base
	titled.Title = "different title"
	a, err := ID(base)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ID(titled)
	if err != nil {
		t.Fatal(err)
	}
	if a != b {
		t.Fatalf("title changed the id: %q vs %q", a, b)
	}

	first, err := ID(Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Gear: map[string]int{"head": 1, "feet": 2}})
	if err != nil {
		t.Fatal(err)
	}
	second, err := ID(Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1", Gear: map[string]int{"feet": 2, "head": 1}})
	if err != nil {
		t.Fatal(err)
	}
	if first != second {
		t.Fatalf("gear key order changed the id: %q vs %q", first, second)
	}
}

func TestNormalizeFillsEmptyCollectionsAndTrimsTheTitle(t *testing.T) {
	got := Input{Title: "  Holy leveling  "}.Normalize()
	if got.PointOrder == nil || len(got.PointOrder) != 0 {
		t.Fatalf("point_order = %v, want an empty slice (nil marshals as null)", got.PointOrder)
	}
	if got.Gear == nil || len(got.Gear) != 0 {
		t.Fatalf("gear = %v, want an empty map (nil marshals as null)", got.Gear)
	}
	if got.Title != "Holy leveling" {
		t.Fatalf("title = %q", got.Title)
	}
}

func TestNewBuildsTheRecordWithItsID(t *testing.T) {
	got, err := New(Input{ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 101, 101, 101, 201}, Title: " Arms leveling "})
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != "znorjmts" || got.Title != "Arms leveling" || got.ClassID != 1 || len(got.PointOrder) != 6 {
		t.Fatalf("build = %+v", got)
	}
	if got.Gear == nil {
		t.Fatal("gear must be an empty map, not nil")
	}
}

func TestSlotsAndItemSlot(t *testing.T) {
	if len(Slots) != 17 {
		t.Fatalf("Slots has %d entries, want 17", len(Slots))
	}
	for slot, want := range map[string]string{
		"finger1": "finger", "finger2": "finger",
		"trinket1": "trinket", "trinket2": "trinket",
		"head": "head", "main_hand": "main_hand",
	} {
		if got := ItemSlot(slot); got != want {
			t.Errorf("ItemSlot(%q) = %q, want %q", slot, got, want)
		}
	}
}

func TestValidIDIsTheShapeIDProduces(t *testing.T) {
	id, err := ID(Input{ClassID: 1, RaceID: 1, TreeVersion: "1.15.0", PointOrder: []int{1}})
	if err != nil {
		t.Fatal(err)
	}
	if !ValidID(id) {
		t.Fatalf("ID() produced %q, which ValidID rejects", id)
	}
	for _, s := range []string{
		"", "abcdefg", "abcdefghi", "abcdef01", "ABCDEFGH", "abcd-efg", "abcdef g",
	} {
		if ValidID(s) {
			t.Errorf("ValidID(%q) = true, want false", s)
		}
	}
}
