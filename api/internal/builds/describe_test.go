package builds

import (
	"reflect"
	"testing"
)

func TestDescribeCountsPointsPerTreeInClientOrder(t *testing.T) {
	data := fixture(t)
	b := Build{
		ID: "znorjmts", ClassID: 1, RaceID: 1, TreeVersion: "test-1",
		PointOrder: []int{101, 101, 101, 101, 101, 103, 201},
		Title:      "Arms leveling",
	}
	got := Describe(data, b)
	if got.ClassName != "Warrior" || got.RaceName != "Human" || got.ClassColor != "#c69b6d" {
		t.Fatalf("names = %+v", got)
	}
	if !reflect.DeepEqual(got.Split, []int{6, 1}) || got.SplitText != "6/1" {
		t.Fatalf("split = %v %q, want [6 1] and 6/1", got.Split, got.SplitText)
	}
	if got.Level != 16 {
		t.Fatalf("level = %d, want 10 + 7 - 1", got.Level)
	}
	if got.Heading != "Human Warrior" || got.Title != "Arms leveling" {
		t.Fatalf("heading = %q title = %q", got.Heading, got.Title)
	}
}

func TestDescribeFallsBackToTheHeadingWhenThereIsNoTitle(t *testing.T) {
	got := Describe(fixture(t), Build{ClassID: 1, RaceID: 3, TreeVersion: "test-1"})
	if got.Title != "Dwarf Warrior" {
		t.Fatalf("title = %q", got.Title)
	}
	if got.SplitText != "0/0" {
		t.Fatalf("split text = %q, want a zero per tree", got.SplitText)
	}
	if got.Level != 9 {
		t.Fatalf("level = %d, want 9 for a build with no points", got.Level)
	}
}

func TestDescribeSurvivesATreeVersionWithNoData(t *testing.T) {
	got := Describe(fixture(t), Build{ClassID: 1, RaceID: 1, TreeVersion: "gone", PointOrder: []int{101}})
	if got.ClassName != "Class 1" || got.RaceName != "Race 1" {
		t.Fatalf("names = %+v", got)
	}
	if got.SplitText != "0" || got.ClassColor != defaultClassColor {
		t.Fatalf("split = %q color = %q", got.SplitText, got.ClassColor)
	}
	if got.Level != 10 {
		t.Fatalf("level = %d", got.Level)
	}
}
