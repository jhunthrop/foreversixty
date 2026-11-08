package activity

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// loopLock is the advisory lock one API instance holds while it runs a pass, so the
// instances of a deployment never run the pass twice at once.
const loopLock = 79_000_001

// DefaultInterval is how often the loop runs its pass.
const DefaultInterval = time.Hour

// Loop is the in-process job behind the feed's derived rows and the weekly snapshots: the
// BiS diff, this week's tier grades and spec boards (with their ranking rows) and the purge.
// It runs at startup and then every Interval.
type Loop struct {
	Store *Store
	// Events, when set, receives the weekly tier moves after each tier snapshot; PublicBaseURL
	// roots their link.
	Events        events.Publisher
	PublicBaseURL string
	DataDir       string
	Build         string
	Curated       Curated
	Interval      time.Duration
	Log           *slog.Logger
	Now           func() time.Time
}

// Run passes until ctx ends.
func (l *Loop) Run(ctx context.Context) {
	interval := l.Interval
	if interval <= 0 {
		interval = DefaultInterval
	}
	for {
		l.passLocked(ctx)
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		}
	}
}

func (l *Loop) passLocked(ctx context.Context) {
	conn, err := l.Store.Pool.Acquire(ctx)
	if err != nil {
		l.logger().Error("activity", "op", "loop lock", "err", err)
		return
	}
	defer conn.Release()
	var held bool
	if err := conn.QueryRow(ctx, `select pg_try_advisory_lock($1)`, loopLock).Scan(&held); err != nil || !held {
		return
	}
	defer func() { _, _ = conn.Exec(context.WithoutCancel(ctx), `select pg_advisory_unlock($1)`, loopLock) }()
	if err := l.Pass(ctx); err != nil {
		l.logger().Error("activity", "op", "loop", "err", err)
	}
}

// Pass runs every step once; a failed step is logged and the next one still runs. It
// answers the first error.
func (l *Loop) Pass(ctx context.Context) error {
	now := l.now()
	var first error
	step := func(name string, run func() (int, error)) {
		n, err := run()
		if err != nil {
			l.logger().Error("activity", "op", name, "err", err)
			if first == nil {
				first = fmt.Errorf("%s: %w", name, err)
			}
			return
		}
		l.logger().Info("activity", "op", name, "rows", n)
	}
	if l.Build != "" {
		step("bis diff", func() (int, error) { return l.Store.DiffBis(ctx, l.DataDir, l.Build) })
		step("tier snapshot", func() (int, error) { return l.Store.SnapshotTiers(ctx, l.DataDir, l.Build, WeekOf(now)) })
		step("tier moves", func() (int, error) { return l.Store.PublishTierMoves(ctx, l.Events, l.PublicBaseURL, now) })
	}
	if len(l.Curated.ZoneBands) > 0 {
		step("boards", func() (int, error) { return l.Store.RefreshBoards(ctx, l.Curated.ZoneBands, now) })
	}
	step("purge", func() (int, error) {
		n, err := l.Store.Purge(ctx, now)
		return int(n), err
	})
	return first
}

func (l *Loop) now() time.Time {
	if l.Now != nil {
		return l.Now().UTC()
	}
	return time.Now().UTC()
}

func (l *Loop) logger() *slog.Logger {
	if l.Log != nil {
		return l.Log
	}
	return slog.Default()
}
