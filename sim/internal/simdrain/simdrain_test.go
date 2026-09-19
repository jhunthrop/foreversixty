package simdrain

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
)

// Two engine paths send a FinalRaidResult and then return WITHOUT ever
// closing the channel (see ToResult's own comment): a failed
// simsignals.RegisterWithId, and SimOptions.IsTest. Neither is
// reachable through a real request today - IDs are always fresh and
// IsTest is always false - so this drives ToResult directly through a
// channel the test controls and deliberately never closes, which is
// exactly what would hang forever on a drain that did not stop for an
// error result. A timeout is the backstop in case a regression brings
// the hang back.
//
// Both lanes are covered by this: sim/cmd/forever-sim and
// sim/cmd/wasm each call this one function, differing only in the tick
// they pass.
func TestToResultStopsOnAnErrorResultEvenIfTheChannelNeverCloses(t *testing.T) {
	reporter := make(chan *proto.ProgressMetrics, 1)
	want := &proto.RaidSimResult{Error: &proto.ErrorOutcome{Message: "could not register for signals"}}
	reporter <- &proto.ProgressMetrics{FinalRaidResult: want}
	// No close(reporter): the point of this test.

	done := make(chan *proto.RaidSimResult, 1)
	go func() { done <- ToResult(reporter, nil) }()

	select {
	case got := <-done:
		if got != want {
			t.Errorf("ToResult returned %+v, want the error result", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("ToResult hung on an error result from a channel that never closed")
	}
}

// The successful-run counterpart: draining still waits for the close
// (and so still sees a mutation the producer makes after sending the
// message) when the result carries no error.
func TestToResultWaitsForCloseOnASuccessfulResult(t *testing.T) {
	reporter := make(chan *proto.ProgressMetrics, 2)
	result := &proto.RaidSimResult{IterationsDone: 500}
	reporter <- &proto.ProgressMetrics{FinalRaidResult: result}

	done := make(chan *proto.RaidSimResult, 1)
	go func() {
		done <- ToResult(reporter, nil)
	}()

	select {
	case <-done:
		t.Fatal("ToResult returned before the channel closed; a producer's post-send mutation would not be visible")
	case <-time.After(50 * time.Millisecond):
		// Still waiting, as it should be.
	}

	// The producer's "post-send mutation" - the real bug this fixed was
	// FinalRaidResult.SampleIteration getting attached to the same
	// pointer after the message was already sent. This only has
	// meaning because ToResult has not returned yet - see the case
	// above.
	result.SampleIteration = &proto.SampleIteration{Dps: 42}
	close(reporter)

	select {
	case got := <-done:
		if got != result || got.GetSampleIteration().GetDps() != 42 {
			t.Errorf("ToResult returned %+v, want the mutated result", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("ToResult did not return after the channel closed")
	}
}

// Every message that is not the final result is a progress tick, and
// the caller's sink is what turns it into JSON lines or a JavaScript
// callback. A nil tick is the no-progress case both lanes use.
func TestToResultHandsEveryTickToTheSink(t *testing.T) {
	reporter := make(chan *proto.ProgressMetrics, 3)
	reporter <- &proto.ProgressMetrics{CompletedIterations: 100, TotalIterations: 500, Dps: 900}
	reporter <- &proto.ProgressMetrics{CompletedIterations: 300, TotalIterations: 500, Dps: 950}
	reporter <- &proto.ProgressMetrics{FinalRaidResult: &proto.RaidSimResult{IterationsDone: 500}}
	close(reporter)

	var seen []int32
	res := ToResult(reporter, func(p *proto.ProgressMetrics) { seen = append(seen, p.CompletedIterations) })
	if res.GetIterationsDone() != 500 {
		t.Errorf("ToResult returned %+v, want the final result", res)
	}
	if len(seen) != 2 || seen[0] != 100 || seen[1] != 300 {
		t.Errorf("the sink saw %v, want the two ticks [100 300] and not the final message", seen)
	}
}
