package adapter

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// TestSummarizePetActiveIsTheEnginesUptime pins the Active figure of a pet
// row to the engine's pet_active_seconds_avg, as a share of the engine's
// iteration clock laid on the summary's own. A Demonology warlock
// registers all four demons: the Imp (sacrificed before the pull) and the
// Felhunter read zero, not the whole fight; a pet out half the fight reads
// half; and a pet out for all of it reads the fight exactly, never past it.
func TestSummarizePetActiveIsTheEnginesUptime(t *testing.T) {
	player := oneAction()
	res := resultWith(player, 100)
	player.Pets = []*proto.UnitMetrics{
		{Name: "Imp"},
		{Name: "Succubus", PetActiveSecondsAvg: res.AvgIterationDuration + 0.4},
		{Name: "Felhunter"},
		{Name: "Voidwalker", PetActiveSecondsAvg: res.AvgIterationDuration / 2},
	}
	got, err := Summarize(res, req())
	if err != nil {
		t.Fatal(err)
	}
	half := int64(math.Round(float64(got.DurationMS) / 2))
	want := map[string]int64{"Fury": got.DurationMS, "Imp": 0, "Succubus": got.DurationMS, "Felhunter": 0, "Voidwalker": half}
	if len(got.DamageDone) != len(want) {
		t.Fatalf("DamageDone has %d rows, want %d", len(got.DamageDone), len(want))
	}
	for _, row := range got.DamageDone {
		if row.ActiveMS != want[row.Name] {
			t.Errorf("%s ActiveMS = %d, want %d", row.Name, row.ActiveMS, want[row.Name])
		}
		// Only a pet row says its Active is a measured uptime; the player's is the fight.
		if isPet := row.Name != "Fury"; row.PetUptime != isPet {
			t.Errorf("%s PetUptime = %v, want %v", row.Name, row.PetUptime, isPet)
		}
	}
}
