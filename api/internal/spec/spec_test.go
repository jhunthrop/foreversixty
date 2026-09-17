package spec

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func testInferrer(t *testing.T) *Inferrer {
	t.Helper()
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, ok := data.Latest()
	if !ok {
		t.Fatal("the fixture build should load")
	}
	return New(b)
}

func TestSplitCountsTalentIdsPerTree(t *testing.T) {
	i := testInferrer(t)
	// Four points in Arms (101 twice, 102, 103) and one in Fury (201).
	split := i.Split("Warrior", []int64{101, 101, 102, 103, 201})
	if len(split) != 2 || split[0] != 4 || split[1] != 1 {
		t.Fatalf("split = %v, want [4 1]", split)
	}
	if got := SplitString(split); got != "4/1" {
		t.Fatalf("SplitString = %q", got)
	}
	if got := i.Name("Warrior", split); got != "Arms" {
		t.Fatalf("name = %q, want the dominant tree", got)
	}
}

func TestSplitReadsATalentsFieldThatIsAlreadyTotals(t *testing.T) {
	i := testInferrer(t)
	split := i.Split("warrior", []int64{31, 20})
	if len(split) != 2 || split[0] != 31 || split[1] != 20 {
		t.Fatalf("split = %v, want [31 20]", split)
	}
	if got := i.Name("Warrior", split); got != "Arms" {
		t.Fatalf("name = %q", got)
	}
}

func TestUnknownTalentsAndClassesAreIgnored(t *testing.T) {
	i := testInferrer(t)
	if got := i.Split("Warrior", []int64{999999, 101}); len(got) != 2 || got[0] != 1 {
		t.Fatalf("split = %v, want the unknown id ignored", got)
	}
	if got := i.Split("Necromancer", []int64{101}); got != nil {
		t.Fatalf("split = %v, want nothing for an unknown class", got)
	}
	if got := i.Name("Necromancer", []int{1, 2}); got != "" {
		t.Fatalf("name = %q, want empty for an unknown class", got)
	}
	if got := i.Name("Warrior", []int{1}); got != "" {
		t.Fatalf("name = %q, want empty when the split does not match the trees", got)
	}
	if got := i.Name("Warrior", []int{0, 0}); got != "" {
		t.Fatalf("name = %q, want empty when no points were spent", got)
	}
}

func TestOfGivesTheNameAndTheSplitTogether(t *testing.T) {
	i := testInferrer(t)
	name, split := i.Of("Warrior", []int64{201, 202, 202})
	if name != "Fury" || split != "0/3" {
		t.Fatalf("Of = %q, %q", name, split)
	}
}

func TestANilInferrerAnswersEmpty(t *testing.T) {
	var i *Inferrer = New(nil)
	if _, ok := i.ClassID("Warrior"); ok {
		t.Fatal("a nil inferrer knows no classes")
	}
	if got := i.Split("Warrior", []int64{101}); got != nil {
		t.Fatalf("split = %v", got)
	}
	name, split := i.Of("Warrior", []int64{101})
	if name != "" || split != "" {
		t.Fatalf("Of = %q, %q", name, split)
	}
}

func TestSplitStringOfNothing(t *testing.T) {
	if got := SplitString(nil); got != "" {
		t.Fatalf("SplitString(nil) = %q", got)
	}
}

func TestTotalsAreOnlyReadWhenTheyLookLikeTotals(t *testing.T) {
	for name, in := range map[string][]int64{
		"too many":    {1, 2, 3},
		"over budget": {40, 40},
		"a talent id": {1451, 1452},
		"all zero":    {0, 0},
	} {
		if _, ok := asTotals(in, 2); ok {
			t.Errorf("%s should not read as totals", name)
		}
	}
}

func TestSplitReadsTalentsWrittenAsSpellIDs(t *testing.T) {
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := data.Build("test-1")
	i := New(b)
	// 12281 is talent 101's spell, in Arms; 20000 is talent 201's, in Fury.
	got := i.Split("Warrior", []int64{12281, 12281, 20000})
	want := []int{2, 1}
	if len(got) != len(want) {
		t.Fatalf("Split = %v, want %v", got, want)
	}
	for n := range want {
		if got[n] != want[n] {
			t.Fatalf("Split = %v, want %v", got, want)
		}
	}
}
