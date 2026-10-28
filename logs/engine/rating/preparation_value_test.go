// logs/engine/rating/preparation_value_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
)

func TestPreparationValue_IsTheRawValueScorePreparationPlaces(t *testing.T) {
	rc, _ := consumables.Load().For(RoleDPS)
	fight := preparedFight(seededAura(17626))
	got, ok := PreparationValue(fight, fixturePlayer, &rc)
	if want := placedPreparation(t, fight); !ok || len(want) != 1 || got != want[0] {
		t.Fatalf("PreparationValue = %v, %v; want %v", got, ok, want)
	}
}

func TestPreparationValue_NoCatalogueIsNotOK(t *testing.T) {
	if _, ok := PreparationValue(preparedFight(), fixturePlayer, nil); ok {
		t.Fatal("a nil catalogue must report no value")
	}
}
