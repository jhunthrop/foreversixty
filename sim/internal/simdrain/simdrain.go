// Package simdrain reads the engine's progress channel to the end and
// hands back the run's final result.
//
// It exists as a package because both lanes need the same loop and the
// invariant in it is subtle: sim/cmd/forever-sim reports progress as
// JSON lines on stderr and sim/cmd/wasm calls a JavaScript callback,
// which is the only difference between them. The browser copy is the
// one where getting it wrong is a wedged worker rather than a failed
// job, so the loop is written once, here, and the tests beside it
// cover both lanes.
package simdrain

import "github.com/wowsims/classic/sim/core/proto"

// ToResult reads reporter until it has the engine's final result,
// handing every intermediate tick to tick along the way. tick may be
// nil, when the caller wants no progress reporting.
//
// It does NOT simply break on the first FinalRaidResult: the engine's
// run() sends that message from INSIDE the producer goroutine, then
// keeps running - for a sample request it still has to return up to
// runSim, which replays the median iteration and only THEN sets
// FinalRaidResult.SampleIteration on that same pointer, before closing
// the channel. Breaking on sight of the message used to read the
// struct while that replay was still in flight, which raced
// SampleIteration nil almost every time. Draining to the channel's
// close instead relies on Go's channel-close happens-before: whatever
// the producer did before close(progress), including that mutation,
// is guaranteed visible once range observes the close.
//
// That fix only holds for a SUCCESSFUL run, because the engine's sim
// body is the only path that closes progress on its way out
// (core/sim.go). Two other engine paths send a FinalRaidResult and
// then return WITHOUT ever closing the channel: a failed
// simsignals.RegisterWithId - an empty or duplicate request id -
// inside RunRaidSimAsync/RunRaidSimConcurrentAsync (core/api.go), and
// SimOptions.IsTest, which registers no closing defer at all
// (core/sim.go) - neither lane's requests set IsTest, but nothing
// stops a future caller from being the first to. A blanket drain hangs
// forever on either. Neither ever carries a sample (core/sim.go's
// replay is itself guarded on result.Error == nil), so an error result
// has nothing left worth waiting for: take it and stop rather than
// block on a close that may never come.
func ToResult(reporter chan *proto.ProgressMetrics, tick func(*proto.ProgressMetrics)) *proto.RaidSimResult {
	var engineRes *proto.RaidSimResult
	for p := range reporter {
		if p.FinalRaidResult != nil {
			engineRes = p.FinalRaidResult
			if engineRes.Error != nil {
				break
			}
			continue
		}
		if tick != nil {
			tick(p)
		}
	}
	return engineRes
}
