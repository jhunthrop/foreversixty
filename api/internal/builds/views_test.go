package builds

import (
	"context"
	"io"
	"log/slog"
	"maps"
	"reflect"
	"sync"
	"testing"
	"time"
)

type fakeAdder struct {
	mu     sync.Mutex
	counts map[string]int64
	calls  int
}

func newFakeAdder() *fakeAdder { return &fakeAdder{counts: map[string]int64{}} }

func (f *fakeAdder) AddViews(_ context.Context, counts map[string]int64) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls++
	for id, n := range counts {
		f.counts[id] += n
	}
	return nil
}

func (f *fakeAdder) snapshot() map[string]int64 {
	f.mu.Lock()
	defer f.mu.Unlock()
	return maps.Clone(f.counts)
}

func quietLogger() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }

func TestViewsBatchesEveryRecordAndFlushesOnClose(t *testing.T) {
	adder := newFakeAdder()
	v := newViews(adder, quietLogger(), 64, time.Hour)
	for _, id := range []string{"znorjmts", "znorjmts", "znorjmts", "jotol22p"} {
		v.Record(id)
	}
	v.Close()

	want := map[string]int64{"znorjmts": 3, "jotol22p": 1}
	if got := adder.snapshot(); !reflect.DeepEqual(got, want) {
		t.Fatalf("counts = %v, want %v", got, want)
	}
	if adder.calls != 1 {
		t.Fatalf("AddViews calls = %d, want one batched write", adder.calls)
	}
}

func TestViewsFlushesOnItsInterval(t *testing.T) {
	adder := newFakeAdder()
	v := newViews(adder, quietLogger(), 64, 5*time.Millisecond)
	defer v.Close()
	v.Record("znorjmts")

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if adder.snapshot()["znorjmts"] == 1 {
			return
		}
		time.Sleep(2 * time.Millisecond)
	}
	t.Fatal("the ticker never flushed the pending view")
}

func TestViewsCloseIsIdempotent(t *testing.T) {
	v := newViews(newFakeAdder(), quietLogger(), 4, time.Hour)
	v.Close()
	v.Close()
}

func TestViewsRecordAfterCloseIsSafe(t *testing.T) {
	// A request can finish rendering while shutdown is closing the counter;
	// that Record must be dropped, not panic on a closed channel.
	v := newViews(newFakeAdder(), quietLogger(), 4, time.Hour)
	v.Close()
	v.Record("znorjmts")
}

func TestViewsRecordDropsWhenTheQueueIsFull(t *testing.T) {
	// No goroutine drains this counter, so the buffer fills: the second
	// record must be dropped rather than block the caller.
	v := &Views{ch: make(chan string, 1), done: make(chan struct{}), log: quietLogger()}
	v.Record("znorjmts")
	v.Record("jotol22p")
	if len(v.ch) != 1 {
		t.Fatalf("queued %d records, want 1 with the second dropped", len(v.ch))
	}
}
