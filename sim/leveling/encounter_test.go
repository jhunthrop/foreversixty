package leveling

import (
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestRankerEncounterIsTheDefaultFightWithNoCreatureType(t *testing.T) {
	got, want := RankerEncounter(), api.DefaultEncounter()
	if got.TargetType != api.TargetTypeUnknown {
		t.Errorf("target type = %q, want %q", got.TargetType, api.TargetTypeUnknown)
	}
	got.TargetType = want.TargetType
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ranker encounter %+v differs from the default %+v beyond its target type", got, want)
	}
}
