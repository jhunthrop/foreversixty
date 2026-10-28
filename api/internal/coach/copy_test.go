package coach

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestSpecVerbatimCopyIsInTheSpec(t *testing.T) {
	got := SpecVerbatim()
	if len(got) < 40 {
		t.Fatalf("SpecVerbatim has %d entries; every spec-shown string belongs in it", len(got))
	}
	for name, s := range got {
		if !specQuoted(t, s) {
			t.Errorf("%s: %q is not quoted in the spec", name, s)
		}
	}
}

func TestSpecVerbatimFillsTheSpecsFilledForms(t *testing.T) {
	cases := map[string]string{
		"BasisTop":       "Top 25 Fury warriors on Onyxia (40): ",
		"TitleCoreCasts": "Bloodthirst: 9 casts lost",
		"TitleHealRank":  "Greater Heal (Rank 4): 44% overheal",
		"SharedDebuff":   "Your share of Sunder Armor: 3 warriors applied it, and the log merges them.",
		"ThinPool":       ThinPool, // the spec shows only the template
		"BasisSimLabel":  BasisSimLabel,
	}
	got := SpecVerbatim()
	for name, want := range cases {
		if got[name] != want {
			t.Errorf("%s = %q, want %q", name, got[name], want)
		}
	}
}

// sampleFields fills any title with plausible values: a two-word name and one number.
var sampleFields = map[string]string{
	"Spell": "Greater Heal", "Aura": "Shield Block", "Resource": "Rage", "Item": "Thunderfury",
	"Buff": "Arcane Intellect", "Slot": "Main Hand", "n": "14", "r": "4", "pct": "44", "t": "3:20",
}

var (
	rankNumber = regexp.MustCompile(`Rank \d+`)
	number     = regexp.MustCompile(`\d+(?:[.,:]\d+)*`)
)

// §5.2: a title says what happened in under eight words, with at most one number (a
// spell's rank is part of its name, not a number).
func TestTitlesShape(t *testing.T) {
	titles := map[string]string{}
	for name, e := range verbatim {
		if strings.HasPrefix(name, "Title") {
			titles[name] = e.format
		}
	}
	for key, s := range experienceCopy {
		if strings.HasPrefix(key, "Title") {
			titles[key] = s
		}
	}
	if len(titles) < 17 { // a floor, not a count: the §8 table owns the exact set
		t.Fatalf("found %d titles, want every Title constant and title_ draft", len(titles))
	}
	for name, format := range titles {
		filled := Fill(format, sampleFields)
		if words := len(strings.Fields(filled)); words >= 8 {
			t.Errorf("%s %q has %d words, want under eight", name, filled, words)
		}
		if nums := number.FindAllString(rankNumber.ReplaceAllString(filled, ""), -1); len(nums) > 1 {
			t.Errorf("%s %q has numbers %q, want at most one", name, filled, nums)
		}
	}
}

func TestFill(t *testing.T) {
	got := Fill(TitleCoreCasts, map[string]string{"Spell": "Bloodthirst", "n": "9"})
	if got != "Bloodthirst: 9 casts lost" {
		t.Fatalf("Fill = %q", got)
	}
	defer func() {
		if r := recover(); r == nil || !strings.Contains(r.(string), `"n"`) {
			t.Fatalf("a missing field must panic under test naming it, got %v", r)
		}
	}()
	Fill(TitleCoreCasts, map[string]string{"Spell": "Bloodthirst"})
}

// The copy's literal numbers are the rules' constants, so a threshold change cannot leave
// the copy stating the old value.
func TestCopyNumbersMatchTheRules(t *testing.T) {
	checks := []struct{ name, text, want string }{
		{"ThinPool", ThinPool, "the coach needs " + strconv.Itoa(PoolMinPlayers) + "."},
		{"DeathNoPotion", DeathNoPotion, "in the " + strconv.Itoa(DeathPotionWindowMS/1000) + " s before."},
		{"DetailOutOfManaDeath", DetailOutOfManaDeath, "under " + PctNumber(ManaLowPct) + "%"},
		{"DetailOutOfMana", DetailOutOfMana, "under " + PctNumber(ManaLowPct) + "%"},
		{"NotReviewableShort", NotReviewableShort, "under " + strconv.Itoa(MinFightMS/1000) + " seconds"},
	}
	for _, c := range checks {
		if !strings.Contains(c.text, c.want) {
			t.Errorf("%s %q does not say %q", c.name, c.text, c.want)
		}
	}
}

// Healer copy never says "heal more" (§20, the tone list).
func TestCopyNeverSaysHealMore(t *testing.T) {
	all := []string{}
	for _, s := range SpecVerbatim() {
		all = append(all, s)
	}
	for _, s := range experienceCopy {
		all = append(all, s)
	}
	for _, s := range all {
		if strings.Contains(strings.ToLower(s), "heal more") {
			t.Errorf("%q says heal more", s)
		}
	}
}
