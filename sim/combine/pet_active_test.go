package combine

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
)

// demoShard is one shard of a Demonology run: the player, the Succubus out
// for `succubusShare` of the shard's fight, and an Imp never out.
func demoShard(durationMS int64, succubusShare float64) api.SimResult {
	row := func(guid string, total int64, activeMS int64, pet bool) summary.Actor {
		return summary.Actor{
			GUID: guid, Total: total, Effective: total, ActiveMS: activeMS, PetUptime: pet,
			Abilities: []summary.Ability{{SpellID: 1, Total: total, Effective: total}},
		}
	}
	return api.SimResult{
		IterationsRun: 1500,
		DPS:           api.Estimate{Mean: 1000},
		Summary: summary.Summary{
			DurationMS: durationMS,
			DamageDone: []summary.Actor{
				row("sim-player", durationMS/2, durationMS, false),
				row("sim-player-pet-0", 0, 0, true),
				row("sim-player-pet-1", durationMS/2, int64(math.Round(succubusShare*float64(durationMS))), true),
			},
		},
	}
}

// TestResultsKeepsEachPetsMeasuredShareOfTheFight pins a pet's Active
// through a combine: the player is active the whole fight, a pet keeps the
// iteration-weighted share of the fight it was summoned for on the
// combined clock, and a pet never out stays at zero. Restating every row
// to the fight's length is how a demon never summoned read 100%.
func TestResultsKeepsEachPetsMeasuredShareOfTheFight(t *testing.T) {
	got, err := Results([]api.SimResult{demoShard(180000, 1), demoShard(170000, 0.5)})
	if err != nil {
		t.Fatal(err)
	}
	d := got.Summary.DurationMS
	want := map[string]int64{
		"sim-player":       d,
		"sim-player-pet-0": 0,
		"sim-player-pet-1": int64(math.Round(0.75 * float64(d))),
	}
	for _, row := range got.Summary.DamageDone {
		if row.ActiveMS != want[row.GUID] {
			t.Errorf("%s ActiveMS = %d, want %d", row.GUID, row.ActiveMS, want[row.GUID])
		}
		if isPet := row.GUID != "sim-player"; row.PetUptime != isPet {
			t.Errorf("%s PetUptime = %v, want %v", row.GUID, row.PetUptime, isPet)
		}
	}
}
