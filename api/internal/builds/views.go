package builds

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

const (
	// viewQueue is how many pending views fit before new ones are dropped.
	viewQueue = 1024
	// viewFlushEvery is how often pending counts are written.
	viewFlushEvery = 5 * time.Second
	// viewFlushTimeout bounds a single flush on its own, independently of
	// the server's shutdown deadline: cmd/api/main.go calls Close after
	// srv.Shutdown has returned, so the two budgets run in sequence rather
	// than nested. A shutdown cut short before the drain finishes simply
	// drops the pending window, which the counter's drop policy allows.
	viewFlushTimeout = 5 * time.Second
)

// ViewAdder is the part of Store the counter needs.
type ViewAdder interface {
	AddViews(ctx context.Context, counts map[string]int64) error
}

// Views counts build page views off the request path. Handlers call Record,
// which never blocks and never fails; one goroutine batches the counts and
// writes them on a ticker. Close asks for one last flush, but it is not a
// guarantee: a shutdown cut short before that flush finishes drops the
// pending window, which the counter's drop policy allows.
type Views struct {
	ch    chan string
	done  chan struct{}
	store ViewAdder
	log   *slog.Logger

	// mu guards closed and the send on ch: Record holds it for reading
	// while it sends, and Close takes it for writing before closing the
	// channel, so a request that finishes during shutdown can never send
	// on a closed channel.
	mu     sync.RWMutex
	closed bool
}

// NewViews starts a counter writing through store.
func NewViews(store ViewAdder, log *slog.Logger) *Views {
	return newViews(store, log, viewQueue, viewFlushEvery)
}

func newViews(store ViewAdder, log *slog.Logger, queue int, every time.Duration) *Views {
	if log == nil {
		log = slog.Default()
	}
	v := &Views{
		ch:    make(chan string, queue),
		done:  make(chan struct{}),
		store: store,
		log:   log,
	}
	go v.run(every)
	return v
}

// Record queues one view for id: a view count is never worth delaying a page
// render or failing a shutdown. A full queue drops the view and logs it; a
// counter that has already been closed drops it silently, because a line per
// in-flight request during shutdown would be noise.
func (v *Views) Record(id string) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if v.closed {
		return
	}
	select {
	case v.ch <- id:
	default:
		v.log.Warn("views", "op", "record", "dropped", id)
	}
}

// Close stops the counter and flushes what is pending. It is safe to call
// more than once, and safe to call while requests are still recording.
func (v *Views) Close() {
	v.mu.Lock()
	if !v.closed {
		v.closed = true
		close(v.ch)
	}
	v.mu.Unlock()
	<-v.done
}

func (v *Views) run(every time.Duration) {
	defer close(v.done)
	ticker := time.NewTicker(every)
	defer ticker.Stop()
	pending := map[string]int64{}
	for {
		select {
		case id, ok := <-v.ch:
			if !ok {
				v.flush(pending)
				return
			}
			pending[id]++
		case <-ticker.C:
			v.flush(pending)
		}
	}
}

// flush writes and empties pending. A failed write is logged and its counts
// are dropped: retrying a view count is not worth unbounded memory.
func (v *Views) flush(pending map[string]int64) {
	if len(pending) == 0 {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), viewFlushTimeout)
	defer cancel()
	if err := v.store.AddViews(ctx, pending); err != nil {
		v.log.Error("views", "op", "flush", "builds", len(pending), "err", err)
	}
	clear(pending)
}
