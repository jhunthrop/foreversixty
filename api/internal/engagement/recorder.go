package engagement

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// markBudget bounds the background write one request starts.
const markBudget = 5 * time.Second

// marker writes one account's surface for one UTC day (Store satisfies it).
type marker interface {
	MarkActive(ctx context.Context, userID int64, day time.Time, surface Surface) error
}

// seenKey is one account on one surface.
type seenKey struct {
	userID  int64
	surface Surface
}

// Recorder is the authenticator's activity hook (auth.ActivityHook). It keeps the last UTC day
// it wrote for each account and surface, so a request costs a map lookup and the database hears
// at most once per account, day and surface per instance. The write runs in the background: the
// request never waits on it.
type Recorder struct {
	Store marker
	Log   *slog.Logger
	// Now is the clock; nil is time.Now.
	Now func() time.Time
	// run starts the background write; nil is a goroutine. Tests run it in place.
	run  func(func())
	seen sync.Map // seenKey -> the UTC day written, as time.DateOnly
}

// Active records that the actor's account was here today on the actor's surface.
func (r *Recorder) Active(a auth.Actor) {
	if !a.Signed() {
		return
	}
	day := r.now().UTC().Truncate(24 * time.Hour)
	stamp := day.Format(time.DateOnly)
	key := seenKey{userID: a.UserID, surface: SurfaceOf(a)}
	if prev, loaded := r.seen.Swap(key, stamp); loaded && prev == stamp {
		return
	}
	r.start(func() { r.mark(key, day, stamp) })
}

// mark writes the day; a failed write forgets it so the next request tries again.
func (r *Recorder) mark(key seenKey, day time.Time, stamp string) {
	ctx, cancel := context.WithTimeout(context.Background(), markBudget)
	defer cancel()
	if err := r.Store.MarkActive(ctx, key.userID, day, key.surface); err != nil {
		r.seen.CompareAndDelete(key, stamp)
		r.logger().Warn("engagement", "op", "mark", "user_id", key.userID, "surface", key.surface, "err", err)
	}
}

func (r *Recorder) start(f func()) {
	if r.run != nil {
		r.run(f)
		return
	}
	go f()
}

func (r *Recorder) now() time.Time {
	if r.Now != nil {
		return r.Now()
	}
	return time.Now()
}

func (r *Recorder) logger() *slog.Logger {
	if r.Log != nil {
		return r.Log
	}
	return slog.Default()
}
