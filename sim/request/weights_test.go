package request

import (
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func weights() api.SimRequest {
	req := fury()
	req.Weights = &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit", "hit"},
		Reference: "attack_power",
	}
	return req
}

// The weights request is the SAME player, buffs, encounter and options
// as a DPS run - that is the whole point of reusing BuildWith - plus
// the stats to weigh.
func TestBuildWeightsIsTheSameRunPlusStats(t *testing.T) {
	got, err := BuildWeights(weights(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	run, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if got.Player.GetName() != run.Raid.Parties[0].Players[0].GetName() {
		t.Error("the weights request carries a different player")
	}
	if got.Player.GetRotation() == nil {
		t.Error("the weights request carries no rotation")
	}
	if got.Player.GetDatabase() != nil {
		t.Error("BuildWeights attached a database; sim/internal/simdb.AttachWeights does that, once, at the caller")
	}
	if got.Encounter.GetDuration() != run.Encounter.GetDuration() {
		t.Error("the weights request fights a different encounter")
	}
	if got.SimOptions.GetIterations() != int32(fury().Iterations) {
		t.Errorf("iterations = %d", got.SimOptions.GetIterations())
	}
	if got.RaidBuffs == nil || got.PartyBuffs == nil || got.Debuffs == nil {
		t.Error("the weights request lost the buffs")
	}
	if got.Tanks == nil {
		t.Error("tanks is nil; the engine indexes it")
	}
	want := []proto.Stat{proto.Stat_StatAgility, proto.Stat_StatAttackPower, proto.Stat_StatCrit, proto.Stat_StatHit}
	if len(got.StatsToWeigh) != len(want) {
		t.Fatalf("stats_to_weigh = %v", got.StatsToWeigh)
	}
	for i, w := range want {
		if got.StatsToWeigh[i] != w {
			t.Errorf("stats_to_weigh[%d] = %v, want %v", i, got.StatsToWeigh[i], w)
		}
	}
	if got.EpReferenceStat != proto.Stat_StatAttackPower {
		t.Errorf("ep_reference_stat = %v", got.EpReferenceStat)
	}
}

func TestBuildWeightsRefusals(t *testing.T) {
	if _, err := BuildWeights(fury(), Options{}); !errors.Is(err, ErrNotWeights) {
		t.Error("BuildWeights accepted a plain run")
	}
	req := weights()
	req.Weights.Stats = []string{"haste", "attack_power"}
	req.Weights.Reference = "attack_power"
	_, err := BuildWeights(req, Options{})
	if !errors.Is(err, ErrUnknownStat) {
		t.Errorf("BuildWeights = %v, want ErrUnknownStat for a stat the engine does not carry", err)
	}
}
