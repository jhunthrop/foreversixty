package pipeline_test

import (
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
)

// TestStatusIsSafeWhileTheTailRuns pins the companion's one real
// concurrency: the loopback UI answers /api/status from the HTTP
// server's goroutine while the companion's ticker is inside Tick. The
// fields Status reads are the ones Tick mutates, so without the
// pipeline's own lock this is a data race and fails under -race.
func TestStatusIsSafeWhileTheTailRuns(t *testing.T) {
	r := newRig(t, 1<<20)

	stop := make(chan struct{})
	var readers sync.WaitGroup
	for range 2 {
		readers.Add(1)
		go func() {
			defer readers.Done()
			for {
				select {
				case <-stop:
					return
				default:
					_ = r.pipe.Status()
				}
			}
		}()
	}

	r.write(fixture.Header + fixture.Zone)
	for i := range 4 {
		r.write(fixture.Encounter(i) + fixture.Heartbeat(i))
		r.clock = r.clock.Add(time.Second)
		if err := r.pipe.Tick(r.clock); err != nil {
			t.Fatal(err)
		}
	}
	close(stop)
	readers.Wait()

	if s := r.pipe.Status(); !s.Logging || s.Fights == 0 {
		t.Fatalf("status = %+v", s)
	}
}
