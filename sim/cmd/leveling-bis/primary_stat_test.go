package main

import "testing"

// TestPrimaryStatCoversEverySpec is this lane's brief (ranker-weights-
// anchor), item 1's own test requirement: every spec in the real,
// published data/curated/specs.json must have a primaryStatBySpec
// entry (primary_stat.go), and that entry must be one of the three
// canonical primary stats - never a reference stat, a secondary
// (crit/hit/haste/...) or a typo. Reads publishedRepoRoot (the real
// catalogue), not a test fixture, the same way TestRogueSpecsWeigh
// Strength (data_test.go) does - a fixture pre-dates this table and a
// future spec added only to the real file would silently escape this
// check if it covered the fixture instead.
func TestPrimaryStatCoversEverySpec(t *testing.T) {
	specs, err := loadAllSpecs(publishedRepoRoot)
	if err != nil {
		t.Fatalf("loadAllSpecs: %v", err)
	}
	if len(specs) == 0 {
		t.Fatal("loadAllSpecs returned no specs - this test proves nothing")
	}
	validPrimary := map[string]bool{"strength": true, "agility": true, "intellect": true}
	for _, s := range specs {
		primary, ok := primaryStatBySpec[s.Spec]
		if !ok {
			t.Errorf("%s: no entry in primaryStatBySpec", s.Spec)
			continue
		}
		if !validPrimary[primary] {
			t.Errorf("%s: primaryStatBySpec = %q, want one of strength/agility/intellect", s.Spec, primary)
		}
		if _, ok := primaryAnchorStat(s); !ok {
			t.Errorf("%s: primaryAnchorStat found no entry at all for primary stat %q", s.Spec, primary)
		}
	}
}

// TestPrimaryAnchorStatUsesIntellectDirectlyWhenPresent is the
// ordinary case for every Intellect-primary spec in today's committed
// data/curated/specs.json: every one of them already carries a
// literal "intellect" row in its own weight_stats, so
// primaryAnchorStat never needs the spell_power/healing_power
// fallback in practice (intellectFallbackRows' own doc).
func TestPrimaryAnchorStatUsesIntellectDirectlyWhenPresent(t *testing.T) {
	spec := specInfo{Spec: "mage-frost", WeightStats: []string{"spell_power", "intellect", "crit"}}
	stat, ok := primaryAnchorStat(spec)
	if !ok || stat != "intellect" {
		t.Errorf("primaryAnchorStat(%+v) = (%q, %v), want (\"intellect\", true)", spec, stat, ok)
	}
}

// TestPrimaryAnchorStatFallsBackToSpellPowerWithoutAnIntellectRow is
// this lane's brief, item 1's own parenthetical: a caster spec whose
// own weight_stats never carries a literal "intellect" row anchors on
// spell_power instead - not exercised by any spec in today's real
// data (see the test above), but the fallback this function's own doc
// promises for the day a weight_stats edit drops intellect.
func TestPrimaryAnchorStatFallsBackToSpellPowerWithoutAnIntellectRow(t *testing.T) {
	spec := specInfo{Spec: "mage-frost", WeightStats: []string{"spell_power", "crit"}}
	stat, ok := primaryAnchorStat(spec)
	if !ok || stat != "spell_power" {
		t.Errorf("primaryAnchorStat(%+v) = (%q, %v), want (\"spell_power\", true)", spec, stat, ok)
	}
}

// TestPrimaryAnchorStatFallsBackToHealingPowerBeforeSpellPower checks
// intellectFallbackRows' own order: healing_power is tried first (a
// healer's own weight_stats leads with it today), ahead of
// spell_power, when a spec's weight_stats carries neither intellect
// nor spell_power but does carry healing_power.
func TestPrimaryAnchorStatFallsBackToHealingPowerBeforeSpellPower(t *testing.T) {
	spec := specInfo{Spec: "priest-holy", WeightStats: []string{"healing_power", "spirit", "mp5"}}
	stat, ok := primaryAnchorStat(spec)
	if !ok || stat != "healing_power" {
		t.Errorf("primaryAnchorStat(%+v) = (%q, %v), want (\"healing_power\", true)", spec, stat, ok)
	}
}

// TestPrimaryAnchorStatReturnsFalseForAnUnknownSpec covers the one
// real failure mode: a spec slug primaryStatBySpec has never heard of
// (unreachable in production given TestPrimaryStatCoversEverySpec,
// but the function must still answer honestly rather than panic or
// invent an anchor).
func TestPrimaryAnchorStatReturnsFalseForAnUnknownSpec(t *testing.T) {
	spec := specInfo{Spec: "no-such-spec", WeightStats: []string{"intellect"}}
	if _, ok := primaryAnchorStat(spec); ok {
		t.Error("primaryAnchorStat(unknown spec) = (_, true), want false")
	}
}

func TestContainsStat(t *testing.T) {
	stats := []string{"attack_power", "strength", "crit"}
	if !containsStat(stats, "strength") {
		t.Error("containsStat(..., \"strength\") = false, want true")
	}
	if containsStat(stats, "agility") {
		t.Error("containsStat(..., \"agility\") = true, want false")
	}
}
