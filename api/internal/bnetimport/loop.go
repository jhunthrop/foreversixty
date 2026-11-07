// api/internal/bnetimport/loop.go
package bnetimport

import (
	"context"
	"time"
)

// refreshLoopLock is the advisory lock one API instance holds while it runs a refresh pass,
// so the instances of a deployment never refresh the same characters twice at once. The
// activity loop holds 79_000_001.
const refreshLoopLock = 79_000_002

// DefaultRefreshInterval is how often the loop looks for due characters. With the 3 h
// active window, an active account's characters are refreshed every 3 to 4 hours.
const DefaultRefreshInterval = time.Hour

// RefreshLoop is the API's in-process Battle.net refresh: at startup and then every
// Interval, one instance re-syncs every due character (staleBnetCharacters' windows). The
// bnet-refresh job runs the same RefreshDue by hand or from Cloud Scheduler.
type RefreshLoop struct {
	Service  *Service
	Interval time.Duration
}

// Run passes until ctx ends.
func (l *RefreshLoop) Run(ctx context.Context) {
	interval := l.Interval
	if interval <= 0 {
		interval = DefaultRefreshInterval
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

func (l *RefreshLoop) passLocked(ctx context.Context) {
	log := l.Service.logger()
	conn, err := l.Service.Pool.Acquire(ctx)
	if err != nil {
		log.Error("bnetimport", "op", "refresh_loop", "err", err)
		return
	}
	defer conn.Release()
	var held bool
	if err := conn.QueryRow(ctx, `select pg_try_advisory_lock($1)`, refreshLoopLock).Scan(&held); err != nil || !held {
		return
	}
	defer func() { _, _ = conn.Exec(context.WithoutCancel(ctx), `select pg_advisory_unlock($1)`, refreshLoopLock) }()
	result, err := l.Service.RefreshDue(ctx)
	if err != nil {
		log.Error("bnetimport", "op", "refresh_loop", "err", err)
		return
	}
	log.Info("bnetimport", "op", "refresh_loop", "considered", result.Considered, "refreshed", result.Refreshed,
		"skipped", result.Skipped, "rate_limited", result.RateLimited)
}
