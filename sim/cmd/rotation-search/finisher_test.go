package main

import (
	"strings"
	"testing"
)

func TestCastsPerIterationFoldsIDsByNameAndZeroFillsWanted(t *testing.T) {
	names := map[int]string{1: "Rupture", 2: "Rupture", 3: "Eviscerate", 4: "Other"}
	got := castsPerIteration(map[int64]int64{1: 10, 2: 5, 4: 99}, 5, names, []string{"Rupture", "Eviscerate"})
	if got["Rupture"] != 3 || got["Eviscerate"] != 0 || len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestFinisherMarkdownFlagsWinnerWithoutRequiredFinishers(t *testing.T) {
	f := finisherCasts{
		abilities: []string{"Eviscerate", "Slice and Dice"},
		baseline:  castTable{"Eviscerate": 2, "Slice and Dice": 1},
		best:      castTable{"Eviscerate": 0, "Slice and Dice": 1},
	}
	out := f.markdown()
	if !strings.Contains(out, "NOT ADOPTABLE: the winner never casts Eviscerate") || !strings.Contains(out, "| Eviscerate | 2.00 | 0.00 |") {
		t.Fatalf("got:\n%s", out)
	}
	f.best["Eviscerate"] = 1
	if strings.Contains(f.markdown(), "NOT ADOPTABLE") {
		t.Fatal("adoptable winner was flagged")
	}
}

func TestCastsPerIterationNamesTalentSpellsTheLearnedTableLacks(t *testing.T) {
	got := castsPerIteration(map[int64]int64{14177: 4}, 2, map[int]string{}, []string{"Cold Blood"})
	if got["Cold Blood"] != 2 {
		t.Fatalf("got %v", got)
	}
}
