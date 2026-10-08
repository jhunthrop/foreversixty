package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// trees is a small class: tier 0 holds a and b, tier 1 holds c (which
// needs a rank in a), tier 2 holds d.
func trees() []leveling.TalentTree {
	return []leveling.TalentTree{{
		Name: "Only",
		Talents: []leveling.TalentNode{
			{ID: 1, Name: "Alpha", Tier: 0, Column: 0, MaxRank: 5},
			{ID: 2, Name: "Beta", Tier: 0, Column: 1, MaxRank: 5},
			{ID: 3, Name: "Gamma", Tier: 1, Column: 0, MaxRank: 5, PrereqTalentID: 1, PrereqRank: 1},
			{ID: 4, Name: "Delta", Tier: 2, Column: 0, MaxRank: 5},
		},
	}}
}

func TestRanksByIDIsCaseInsensitiveAndRefusesUnknownNames(t *testing.T) {
	got, err := ranksByID(trees(), map[string]int{"alpha": 2, "DELTA": 1})
	if err != nil {
		t.Fatal(err)
	}
	if got[1] != 2 || got[4] != 1 || len(got) != 2 {
		t.Errorf("ranks = %v, want {1:2 4:1}", got)
	}
	if _, err := ranksByID(trees(), map[string]int{"Omega": 1}); err == nil {
		t.Error("an unknown talent name resolved")
	}
}

func TestLegalHoldsABuildToTheTreeRules(t *testing.T) {
	cases := []struct {
		name    string
		ranks   map[int]int
		wantErr string
	}{
		{"legal", map[int]int{1: 5, 2: 5, 3: 1, 4: 1}, ""},
		{"too many points", map[int]int{1: 5, 2: 5, 3: 5, 4: 5}, "spends"},
		{"rank above the maximum", map[int]int{1: 6, 2: 5, 3: 5, 4: 5}, "of 5 points"},
		{"tier not unlocked", map[int]int{1: 3, 3: 1, 4: 1}, "needs 5 points above it"},
		{"prerequisite missing", map[int]int{2: 5, 3: 3, 4: 5}, "needs 1 points in Alpha"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := legalFor(trees(), c.ranks, sum(c.ranks))
			if c.name == "too many points" {
				err = legalFor(trees(), c.ranks, 11)
			}
			switch {
			case c.wantErr == "" && err != nil:
				t.Errorf("a legal build was refused: %v", err)
			case c.wantErr != "" && (err == nil || !strings.Contains(err.Error(), c.wantErr)):
				t.Errorf("err = %v, want one containing %q", err, c.wantErr)
			}
		})
	}
}

func sum(ranks map[int]int) int {
	n := 0
	for _, r := range ranks {
		n += r
	}
	return n
}

func TestSiteStringIsOneDigitPerTalentInTreeOrder(t *testing.T) {
	if got := siteString(trees(), map[int]int{1: 5, 3: 2, 4: 1}); got != "5021" {
		t.Errorf("siteString = %q, want 5021", got)
	}
}
