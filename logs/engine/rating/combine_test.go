// logs/engine/rating/combine_test.go
package rating

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/weights"
)

func roleWeights(t *testing.T, role string) weights.RoleWeights {
	t.Helper()
	w, ok := weights.Default().For(role)
	if !ok {
		t.Fatalf("weights.Default().For(%q) reported not-ok", role)
	}
	return w
}

func TestCombineRenormalisesWeightsAcrossExcludedComponents(t *testing.T) {
	w := roleWeights(t, RoleTank) // Output 10, Survival 35, Mechanics 20, Utility 20, Preparation 10, Activity 5
	components := [6]Component{
		{Name: ComponentNameOutput, Score: 55},
		{Name: ComponentNameSurvival, Score: 77.06},
		{Name: ComponentNameMechanics, Excluded: true, Reason: ReasonNoMechanicsTable},
		{Name: ComponentNameUtility, Score: 91},
		{Name: ComponentNamePreparation, Score: 100},
		{Name: ComponentNameActivity, Score: 70},
	}
	overallUncapped, overall, capped, _, _ := combine(&components, w, false, true, DefaultCapThreshold)
	if capped {
		t.Fatal("no cap condition was signalled; capped must be false")
	}
	// spec §1.6's tank worked example, corrected: the original text mis-stated
	// Survival's combined score (76.06, should be 77.06) and the tank weight
	// table's Utility figure in its redistribution step (15, should be 20 per
	// §1.4's own tank row) -- see the spec's own "CORRECTION" note next to this
	// example.
	if want := 80.21; overallUncapped != want {
		t.Fatalf("overallUncapped = %v, want %v (spec §1.6's tank worked example, corrected)", overallUncapped, want)
	}
	if overall != overallUncapped {
		t.Fatalf("overall = %v, want %v (no cap fired)", overall, overallUncapped)
	}
	if components[2].Weight != 0 {
		t.Errorf("an excluded component's Weight must be 0, got %v", components[2].Weight)
	}
	var sumWeights float64
	for _, c := range components {
		sumWeights += c.Weight
	}
	if roundTo1(sumWeights) != 100 {
		t.Errorf("renormalised weights sum to %v, want 100", sumWeights)
	}
}

func TestCombineAppliesTheCatastropheCapWhenEnabled(t *testing.T) {
	w := roleWeights(t, RoleDPS)
	components := [6]Component{
		{Name: ComponentNameOutput, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNameSurvival, Score: 40, Basis: BasisAbsolute},
		{Name: ComponentNameMechanics, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNameUtility, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNamePreparation, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNameActivity, Score: 90, Basis: BasisPercentile},
	}
	overallUncapped, overall, capped, basis, _ := combine(&components, w, true, true, DefaultCapThreshold)
	if !capped {
		t.Fatal("survivalDeathScoreZero was true; capped must be true")
	}
	if overallUncapped <= DefaultCapThreshold {
		t.Fatalf("overallUncapped = %v, fixture is built to exceed the cap", overallUncapped)
	}
	if overall != DefaultCapThreshold {
		t.Fatalf("overall = %v, want the cap threshold %v", overall, DefaultCapThreshold)
	}
	if basis != "mixed" {
		t.Fatalf("basis = %q, want mixed (percentile and absolute both present)", basis)
	}
}

func TestCombineLeavesOverallUncappedWhenCapDisabled(t *testing.T) {
	w := roleWeights(t, RoleDPS)
	components := [6]Component{
		{Name: ComponentNameOutput, Score: 90}, {Name: ComponentNameSurvival, Score: 40},
		{Name: ComponentNameMechanics, Score: 90}, {Name: ComponentNameUtility, Score: 90},
		{Name: ComponentNamePreparation, Score: 90}, {Name: ComponentNameActivity, Score: 90},
	}
	overallUncapped, overall, capped, _, _ := combine(&components, w, true, false, DefaultCapThreshold)
	if !capped {
		t.Fatal("the condition still fired even though the cap is disabled")
	}
	if overall != overallUncapped {
		t.Fatalf("overall = %v, want it to equal overallUncapped (%v) when the cap is disabled", overall, overallUncapped)
	}
}

func TestRosterRowFindsAPlayerAndSaysWhenThereIsNone(t *testing.T) {
	fight := fixtureSummary()
	if _, ok := rosterRow(fight, "does-not-exist"); ok {
		t.Fatal("an unknown guid must report false")
	}
	if _, ok := rosterRow(fight, fixturePlayer); !ok {
		t.Fatal("the fixture's own player must be found")
	}
}

func TestRound2RoundsToTwoDecimalPlaces(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{70.004, 70.0}, {70.006, 70.01}, {-1.006, -1.01}, {0, 0}, {99.994, 99.99},
	}
	for _, c := range cases {
		if got := round2(c.in); got != c.want {
			t.Errorf("round2(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestWeightOfPanicsOnAnUnrecognisedComponentName(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("weightOf(\"not-a-component\") did not panic, want a panic naming the bad component")
		}
		msg, ok := r.(string)
		if !ok || !strings.Contains(msg, "not-a-component") {
			t.Fatalf("panic value = %v, want a message naming the unrecognised component", r)
		}
	}()
	weightOf(roleWeights(t, RoleDPS), "not-a-component")
}

func TestSpecSlugMatchesCuratedSpecsFileFormat(t *testing.T) {
	cases := []struct{ class, spec, want string }{
		{"Warrior", "Protection", "warrior-protection"},
		{"Hunter", "Beast Mastery", "hunter-beast-mastery"},
		{"Priest", "Holy", "priest-holy"},
	}
	for _, c := range cases {
		if got := specSlug(c.class, c.spec); got != c.want {
			t.Errorf("specSlug(%q, %q) = %q, want %q", c.class, c.spec, got, c.want)
		}
	}
}
