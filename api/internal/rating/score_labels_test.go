package rating

import (
	"encoding/json"
	"os"
	"testing"
)

// scoreLabels is the shared fixture the site and the addon read the score's ceilings and tier
// floors from (docs/contracts/fixtures/score-labels.json); this formula is their only home, so
// the fixture must say what the constants say.
type scoreLabels struct {
	Components []struct {
		Key string `json:"key"`
		Max int    `json:"max"`
	} `json:"components"`
	Tiers map[string]struct {
		From int `json:"from"`
	} `json:"tiers"`
}

func TestScoreLabelsFixtureMatchesTheFormula(t *testing.T) {
	raw, err := os.ReadFile("../../../docs/contracts/fixtures/score-labels.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var labels scoreLabels
	if err := json.Unmarshal(raw, &labels); err != nil {
		t.Fatalf("decode fixture: %v", err)
	}
	wantMax := map[string]int{"gear": GearMax, "experience": ExperienceMax, "parses": ParsesMax,
		"attendance": AttendanceMax, "discipline": DisciplineMax}
	if len(labels.Components) != len(wantMax) {
		t.Fatalf("fixture has %d components, the formula %d", len(labels.Components), len(wantMax))
	}
	for _, component := range labels.Components {
		if want, ok := wantMax[component.Key]; !ok || component.Max != want {
			t.Errorf("component %q max = %d, the formula says %d", component.Key, component.Max, want)
		}
	}
	wantFrom := map[string]int{TierCommon: 0, TierUncommon: UncommonFrom, TierRare: RareFrom, TierEpic: EpicFrom,
		TierLegendary: LegendaryFrom}
	for tier, want := range wantFrom {
		if got, ok := labels.Tiers[tier]; !ok || got.From != want {
			t.Errorf("tier %q from = %d, the formula says %d", tier, got.From, want)
		}
	}
}
