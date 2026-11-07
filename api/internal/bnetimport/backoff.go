// api/internal/bnetimport/backoff.go
package bnetimport

import (
	"sync"
	"time"
)

// The 429 backoff: the first rate limit pauses this process's automatic refreshes for
// backoffFirst, each further one in a row doubles the pause up to backoffMax, and a batch
// that finishes without one resets it. Battle.net's quota is per client, shared by every
// sign-in, visit and loop in the process, so one pause covers them all.
const (
	backoffFirst = 30 * time.Second
	backoffMax   = 15 * time.Minute
)

// rateBackoff is the process's pause after a 429. The zero value is not paused.
type rateBackoff struct {
	mu    sync.Mutex
	until time.Time
	step  time.Duration
}

// hit records a 429 at now and answers the pause it starts.
func (b *rateBackoff) hit(now time.Time) time.Duration {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.step = min(max(b.step*2, backoffFirst), backoffMax)
	b.until = now.Add(b.step)
	return b.step
}

// reset ends the run of 429s; a pause already started still runs out.
func (b *rateBackoff) reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.step = 0
}

// blocked reports whether now falls inside the pause.
func (b *rateBackoff) blocked(now time.Time) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return now.Before(b.until)
}
