package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/statid"
	"github.com/wowsims/classic/sim/core/proto"
)

// registerEngine and nextRunID are the two simrun.go helpers that do
// not themselves reach into the engine's sim loop (registerEngine only
// registers agent factories - sync.Once guarded, safe to call
// repeatedly; nextRunID is pure string formatting), so they are safe
// to exercise directly without triggering the real, slow engine path
// this lane's brief forbids in a test.

func TestRegisterEngineIsIdempotent(t *testing.T) {
	// Calling this more than once (as every other test in this package
	// that reaches runSpec/verifyBand/rankTrinketSlot already does,
	// indirectly, through the real registerEngineOnce guard those
	// functions used to call before this lane's engineRunner refactor)
	// must not panic or double-register anything - sync.Once's whole
	// contract.
	registerEngine()
	registerEngine()
}

func TestNextRunIDIsUniquePerCall(t *testing.T) {
	a := nextRunID()
	b := nextRunID()
	if a == b {
		t.Fatalf("nextRunID returned the same id twice: %q", a)
	}
	if a == "" || b == "" {
		t.Fatal("nextRunID returned an empty id")
	}
}

// A compile-time check that realEngine actually satisfies engineRunner
// (this lane's whole reason for existing: runSpec/verifyBand/
// rankTrinketSlot take engineRunner, and main's run() constructs a
// realEngine{} as the only production implementation).
var _ engineRunner = realEngine{}

// This lane's brief, item 2: referenceStatRawWeight reads the raw
// (un-normalised) DPS-per-point straight off a StatWeightsResult -
// exercised directly here (no real engine sim needed - the function
// only reads a hand-built proto message) rather than only indirectly
// through runWeights, which this package's own brief forbids exercising
// with the real, slow engine in a test.
func TestReferenceStatRawWeightReadsTheReferenceStatsOwnRawEntry(t *testing.T) {
	reference, ok := statid.Parse("attack_power")
	if !ok {
		t.Fatal("statid.Parse(\"attack_power\") = false, want a known stat id")
	}
	raw := make([]float64, int(reference)+1)
	raw[reference] = 14.2
	res := &proto.StatWeightsResult{
		Dps: &proto.StatWeightValues{Weights: &proto.UnitStats{Stats: raw}},
	}
	req := api.SimRequest{Weights: &api.WeightsSpec{Stats: []string{"attack_power"}, Reference: "attack_power"}}

	got, err := referenceStatRawWeight(res, req)
	if err != nil {
		t.Fatalf("referenceStatRawWeight: %v", err)
	}
	if got != 14.2 {
		t.Fatalf("referenceStatRawWeight = %v, want 14.2", got)
	}
}

func TestReferenceStatRawWeightErrorsOnAnUnknownReferenceStat(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{Stats: []string{"attack_power"}, Reference: "not_a_real_stat"}}
	if _, err := referenceStatRawWeight(&proto.StatWeightsResult{}, req); err == nil {
		t.Fatal("referenceStatRawWeight with an unknown reference stat = nil error, want one")
	}
}

func TestReferenceStatRawWeightErrorsWithNoWeightsBlock(t *testing.T) {
	if _, err := referenceStatRawWeight(&proto.StatWeightsResult{}, api.SimRequest{}); err == nil {
		t.Fatal("referenceStatRawWeight with no Weights block = nil error, want one")
	}
}
