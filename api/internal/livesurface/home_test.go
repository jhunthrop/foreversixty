package livesurface

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/app"
)

type fakeHomes struct {
	home  app.Home
	asked []string
}

func (f *fakeHomes) HomeOf(_ context.Context, region, ruleset, name string) (app.Home, bool, error) {
	f.asked = []string{region, ruleset, name}
	return f.home, f.home.State != "", nil
}

func TestHomeUpgradesKeepTheHomeAnswersRowsAndTotals(t *testing.T) {
	homes := &fakeHomes{home: app.Home{State: app.StateOK, Unit: ptr("dps"), UpgradeCount: 8, TotalGain: ptr(18.3),
		Upgrades: []app.Upgrade{{Slot: "waist", SlotLabel: "Waist", Gain: ptr(5.3),
			Pick: app.Pick{Item: app.Item{ID: 228190, Name: ptr("Knowledge of the Timbermaw"), Quality: ptr(4), Icon: ptr("inv_belt_09"), ItemLevel: ptr(65)}}},
			{Slot: "trinket1", SlotLabel: "Trinket 1", Pick: app.Pick{Item: app.Item{ID: 7}}}}}}
	got, err := HomeUpgrades{Homes: homes}.WidgetUpgrades(t.Context(), boltzKey)
	if err != nil || got == nil {
		t.Fatalf("upgrades = %v, %v", got, err)
	}
	if homes.asked[0] != "us" || homes.asked[1] != "pvp" || homes.asked[2] != "lightningbolltz" {
		t.Fatalf("asked %v", homes.asked)
	}
	if got.Count != 8 || *got.TotalGain != 18.3 || got.Unit != "dps" || got.Sentence != nil || len(got.Rows) != 2 {
		t.Fatalf("upgrades = %+v", got)
	}
	if r := got.Rows[0]; r.Pick.Name != "Knowledge of the Timbermaw" || r.Pick.Quality != 4 || *r.Gain != 5.3 {
		t.Fatalf("row = %+v", r)
	}
	if r := got.Rows[1]; r.Gain != nil || r.Pick.ID != 7 {
		t.Fatalf("unmodelled row = %+v", r)
	}
}

func TestHomeUpgradesWordEveryEmptyState(t *testing.T) {
	cases := map[string]string{
		app.StateNoGear: sentenceNoGear, app.StateUnder20: sentenceUnder20, app.StateNoSpec: sentenceNoSpec,
		app.StateNoList: sentenceNoList, app.StateOK: sentenceAllBest,
	}
	for state, want := range cases {
		got, _ := HomeUpgrades{Homes: &fakeHomes{home: app.Home{State: state}}}.WidgetUpgrades(t.Context(), boltzKey)
		if got == nil || got.Sentence == nil || *got.Sentence != want || got.Rows == nil {
			t.Errorf("%s: %+v", state, got)
		}
	}
	if got, err := (HomeUpgrades{Homes: &fakeHomes{}}).WidgetUpgrades(t.Context(), boltzKey); got != nil || err != nil {
		t.Fatalf("a character with no page: %v, %v", got, err)
	}
}
