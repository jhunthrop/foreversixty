package builds

import (
	"context"
	"errors"
	"reflect"
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// fakeSimLookup is a SimLookup test double: WithSimDPS's tests never
// need a real sims.Store, only the three answers ForBuild can give.
type fakeSimLookup struct {
	res simapi.SimResult
	ok  bool
	err error
}

func (f fakeSimLookup) ForBuild(context.Context, string) (simapi.SimResult, bool, error) {
	return f.res, f.ok, f.err
}

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

func TestWithSimDPSAddsTheLineWhenADoneSimExists(t *testing.T) {
	sims := fakeSimLookup{ok: true, res: simapi.SimResult{DPS: simapi.Estimate{Mean: 12345.4, Error: 107.3}}}
	got, err := Description{}.WithSimDPS(context.Background(), "znorjmts", sims)
	if err != nil {
		t.Fatal(err)
	}
	if got.DPSLine != "12,345 DPS ±210" {
		t.Fatalf("dps line = %q", got.DPSLine)
	}
}

func TestWithSimDPSFormatsALargeMeanWithEveryComma(t *testing.T) {
	sims := fakeSimLookup{ok: true, res: simapi.SimResult{DPS: simapi.Estimate{Mean: 1234567, Error: 500}}}
	got, err := Description{}.WithSimDPS(context.Background(), "znorjmts", sims)
	if err != nil {
		t.Fatal(err)
	}
	if got.DPSLine != "1,234,567 DPS ±980" {
		t.Fatalf("dps line = %q", got.DPSLine)
	}
}

func TestWithSimDPSLeavesTheLineEmptyWhenTheBuildHasNoSim(t *testing.T) {
	got, err := Description{}.WithSimDPS(context.Background(), "znorjmts", fakeSimLookup{ok: false})
	if err != nil {
		t.Fatal(err)
	}
	if got.DPSLine != "" {
		t.Fatalf("dps line = %q, want empty when no sim was ever run", got.DPSLine)
	}
}

func TestWithSimDPSReturnsALookupErrorAndLeavesTheLineEmpty(t *testing.T) {
	got, err := Description{}.WithSimDPS(context.Background(), "znorjmts",
		fakeSimLookup{err: errors.New("database is down")})
	if err == nil {
		t.Fatal("want an error from a failing lookup")
	}
	if got.DPSLine != "" {
		t.Fatalf("dps line = %q, want empty on a lookup error - the card must still render", got.DPSLine)
	}
}

func TestWithSimDPSIsANoOpWithoutALookup(t *testing.T) {
	d := Description{Title: "Arms leveling"}
	got, err := d.WithSimDPS(context.Background(), "znorjmts", nil)
	if err != nil {
		t.Fatal(err)
	}
	if got.DPSLine != "" || got.Title != "Arms leveling" {
		t.Fatalf("got = %+v, want d unchanged with a nil lookup", got)
	}
}
