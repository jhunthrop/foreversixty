package request

import (
	"errors"
	"slices"
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
	// BuildWeights multiplies the request's own Iterations by
	// api.WeightsIterationsFactor (see its doc): sim/core/statweight.go's
	// WeightsStdev is a population standard deviation that iteration
	// count alone cannot shrink, and adapter.Weights' sqrt(N)
	// conversion needs the engine to actually run at that multiplied
	// count.
	if want := int32(fury().Iterations) * int32(api.WeightsIterationsFactor); got.SimOptions.GetIterations() != want {
		t.Errorf("iterations = %d, want %d", got.SimOptions.GetIterations(), want)
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

// A weights sweep is one sub-sim per stat per direction, and no
// StatWeightsResult consumer ever reads a cast log - so BuildWeights
// must never ask the engine for one, regardless of what the caller's
// own Options say. Passing NoSampleIteration: false here is the
// point: if BuildWeights just forwarded the caller's opt unchanged, a
// weights request would inherit a plain run's sample by default and
// every sub-sim would pay for a replay nothing reads.
func TestBuildWeightsNeverAsksForASample(t *testing.T) {
	got, err := BuildWeights(weights(), Options{NoSampleIteration: false})
	if err != nil {
		t.Fatal(err)
	}
	if got.SimOptions.GetSampleIteration() {
		t.Error("a weights request asked for a sample iteration; no consumer reads a stat sweep's cast log")
	}
}

// TestBuildWeightsIterationsMatchesTheCostEstimate pins the thing
// that breaks quietly if BuildWeights and api.WeightsIterations ever
// disagree about the multiplied count: the server would either
// refuse a run it could afford or accept one it cannot. What
// BuildWeights actually sets the engine's SimOptions.Iterations to,
// halved back out for RNG parity and expanded by the baseline-plus-
// two-passes-per-stat shape, must equal what WeightsIterations
// costs the same request at.
func TestBuildWeightsIterationsMatchesTheCostEstimate(t *testing.T) {
	req := weights()
	got, err := BuildWeights(req, Options{})
	if err != nil {
		t.Fatal(err)
	}
	engineIterations := int(got.SimOptions.GetIterations())
	distinctStats := len(req.Weights.Stats)
	totalRun := (engineIterations / 2) * (1 + 2*distinctStats)
	if want := api.WeightsIterations(req); totalRun != want {
		t.Errorf("BuildWeights implies %d total iterations run, api.WeightsIterations costs %d",
			totalRun, want)
	}
}

func TestBuildWeightsRefusals(t *testing.T) {
	if _, err := BuildWeights(fury(), Options{}); !errors.Is(err, ErrNotWeights) {
		t.Error("BuildWeights accepted a plain run")
	}
	// "haste" is refused before BuildWeights' own ParseStat loop ever
	// runs: BuildWith validates the request first, and
	// WeightsSpec.validate now bounds Stats to api.KnownStats -
	// TestAPIKnownStatsMatchTheGeneratedVocabulary below pins that
	// copy against this package's own KnownStats(), the same
	// vocabulary ParseStat resolves against. ParseStat's ErrUnknownStat
	// stays as the loop's own defense - reachable only if that copy
	// ever drifts from the generated list - rather than being removed
	// as dead code, since the two lists are independently maintained.
	req := weights()
	req.Weights.Stats = []string{"haste", "attack_power"}
	req.Weights.Reference = "attack_power"
	if _, err := BuildWeights(req, Options{}); err == nil {
		t.Error("BuildWeights accepted a stat id outside the vocabulary")
	}
}

// sim/api may not import the engine's proto - no protobuf crosses a
// lane boundary - so api.KnownStats is a copy of KnownStats rather
// than a call to it, kept only so WeightsSpec.validate can bound a
// request's size at the envelope's own boundary. This is the one
// place both lists are in scope together, and it is what keeps the
// copy from drifting the day the engine's Stat enum gains or loses a
// value: the same day TestKnownStatsAreTheEnginesEnum would fail if
// this package's own list moved, this test fails if the copy did not
// move with it.
func TestAPIKnownStatsMatchTheGeneratedVocabulary(t *testing.T) {
	want := KnownStats()
	got := slices.Clone(api.KnownStats)
	slices.Sort(got)
	if !slices.Equal(got, want) {
		for _, id := range got {
			if !slices.Contains(want, id) {
				t.Errorf("api.KnownStats lists %q, which the generated vocabulary does not", id)
			}
		}
		for _, id := range want {
			if !slices.Contains(got, id) {
				t.Errorf("the generated vocabulary lists %q and api.KnownStats does not", id)
			}
		}
	}
}
