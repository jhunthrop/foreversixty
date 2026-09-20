package request

// The stat weights request.
//
// The engine computes weights by running the same character twice per
// stat, once with a little more of it and once with a little less, so
// a weights request IS a raid sim request with a list of stats
// attached. Building it by taking BuildWith's output apart is what
// keeps the two from diverging: a buff, a consumable or a spec option
// that reaches a DPS run reaches the weights run by construction,
// rather than by somebody remembering to copy the line.

import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

var (
	// ErrNotWeights is returned when a request with no weights block is
	// handed to BuildWeights.
	ErrNotWeights = errors.New("request: the request carries no weights block")
	// ErrUnknownStat is returned for a stat id the engine has no Stat
	// value for. Weighing it as nothing would report a weight of zero
	// for a stat the player asked about.
	ErrUnknownStat = errors.New("request: unknown stat")
)

// BuildWeights turns a validated weights SimRequest into the engine's
// StatWeightsRequest.
func BuildWeights(req api.SimRequest, opt Options) (*proto.StatWeightsRequest, error) {
	if req.Weights == nil {
		return nil, ErrNotWeights
	}
	// A weights sweep is one sub-sim per stat per direction, and no
	// StatWeightsResult consumer ever reads a cast log: forcing
	// NoSampleIteration here, regardless of what the caller's own opt
	// asked for, is what keeps every one of those sub-sims from paying
	// for a median-iteration replay nothing downstream looks at (the
	// cost contract 10.3 exists to avoid).
	opt.NoSampleIteration = true
	run, err := BuildWith(req, opt)
	if err != nil {
		return nil, err
	}
	// See api.WeightsIterationsFactor: the engine's own stat-weights
	// stdev is a population standard deviation that iteration count
	// cannot shrink by itself, only the sample size it is later
	// divided by can - sim/adapter.Weights does that division using
	// this same factor, so the two must never disagree about how many
	// iterations a weight was actually built from.
	run.SimOptions.Iterations *= int32(api.WeightsIterationsFactor)
	stats := make([]proto.Stat, 0, len(req.Weights.Stats))
	for _, id := range req.Weights.Stats {
		s, ok := ParseStat(id)
		if !ok {
			return nil, fmt.Errorf("%w: %q; the ids are in sim/request/IDS.md", ErrUnknownStat, id)
		}
		stats = append(stats, s)
	}
	reference, ok := ParseStat(req.Weights.Reference)
	if !ok {
		return nil, fmt.Errorf("%w: reference %q", ErrUnknownStat, req.Weights.Reference)
	}
	party := run.Raid.Parties[0]
	return &proto.StatWeightsRequest{
		Player:     party.Players[0],
		RaidBuffs:  run.Raid.Buffs,
		PartyBuffs: party.Buffs,
		Debuffs:    run.Raid.Debuffs,
		Encounter:  run.Encounter,
		SimOptions: run.SimOptions,
		// The engine indexes this slice; a nil one is a panic rather
		// than a raid with no tank.
		Tanks:           []*proto.UnitReference{},
		StatsToWeigh:    stats,
		EpReferenceStat: reference,
	}, nil
}
