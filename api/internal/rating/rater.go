// api/internal/rating/rater.go
package rating

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// RateQueueDepth is how many fights may wait to be rated. Beyond it a fight-close drops
// the rating rather than blocking the ingest; the nightly backfill job picks it up on its
// next sweep (RateFight's rewrite-detection makes catching up later exactly as correct as
// catching up immediately).
const RateQueueDepth = 512

// RateDeps is what the rater needs.
type RateDeps struct {
	Store *Store
	Log   *slog.Logger
}

func (d RateDeps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

// Rater runs ratings out of band, so the companion's fight-close call returns at once. It
// is sims.Scorer's exact shape, for sims.Scorer's exact reason: the consumer has to outlive
// the handlers that feed it, so Close - not a cancelled context - is the only thing that
// stops it, and the shutdown order is Shutdown then Close.
type Rater struct {
	Deps  RateDeps
	queue chan reports.RatedFight
	done  chan struct{}
	once  sync.Once

	// mu guards closed against a Schedule that races Close, exactly as
	// sims.Scorer's does: without it a Schedule racing Close is a send on a closed
	// channel, which is a panic rather than a drop.
	mu     sync.RWMutex
	closed bool
}

var _ reports.Rater = (*Rater)(nil)

// NewRater returns a rater with an empty queue.
func NewRater(d RateDeps) *Rater {
	return &Rater{Deps: d, queue: make(chan reports.RatedFight, RateQueueDepth), done: make(chan struct{})}
}

// Schedule queues one fight. It never blocks: a full queue drops the fight (logged), and a
// Schedule racing Close is a no-op rather than a panic.
func (s *Rater) Schedule(f reports.RatedFight) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.closed {
		return
	}
	select {
	case s.queue <- f:
	default:
		s.Deps.logger().Error("rating", "op", "schedule", "report", f.ReportID, "fight", f.FightIndex,
			"err", "the rating queue is full; the backfill job will catch this fight up later")
	}
}

// Run drains the queue until Close is called. It never returns an error: every failure to
// rate one fight is logged and the loop continues, the same best-effort contract i.rate
// itself promises the ingest.
func (s *Rater) Run(ctx context.Context) {
	for {
		select {
		case f, ok := <-s.queue:
			if !ok {
				return
			}
			if err := s.Deps.Store.RateFight(ctx, RatedFight{
				ReportID: f.ReportID, FightIndex: f.FightIndex, Region: f.Region, Ruleset: f.Ruleset,
				FoughtAt: f.FoughtAt, EncounterID: f.EncounterID, Summary: f.Summary,
			}); err != nil {
				s.Deps.logger().Error("rating", "op", "rate", "report", f.ReportID, "fight", f.FightIndex, "err", err)
			}
		case <-s.done:
			return
		}
	}
}

// Close stops Run and makes every later Schedule a no-op. Idempotent.
func (s *Rater) Close() {
	s.once.Do(func() {
		s.mu.Lock()
		s.closed = true
		s.mu.Unlock()
		close(s.done)
	})
}
