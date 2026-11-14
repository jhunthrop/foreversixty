// api/internal/directory/loop.go
package directory

import (
	"context"
	"time"
)

// JobCommand is the Cloud Run job that runs one pass by hand or from Cloud Scheduler; the
// API's own Loop runs the same Pass every hour.
const JobCommand = "directory-refresh"

// loopLock is the advisory lock one API instance holds while it runs a pass, so the
// instances of a deployment never read the same characters twice at once. The activity loop
// holds 79_000_001 and the Battle.net refresh 79_000_002.
const loopLock = 79_000_003

const (
	// DefaultInterval is how often the loop starts a pass; a pass runs at most this long.
	DefaultInterval = time.Hour
	// startDelay keeps the first pass off a fresh instance's startup.
	startDelay = 5 * time.Minute
)

// Loop runs a pass every Interval in the API process.
type Loop struct {
	Service  *Service
	Interval time.Duration
	// StartDelay overrides startDelay (tests).
	StartDelay time.Duration
}

// Run passes until ctx ends.
func (l *Loop) Run(ctx context.Context) {
	interval := l.Interval
	if interval <= 0 {
		interval = DefaultInterval
	}
	delay := l.StartDelay
	if delay <= 0 {
		delay = startDelay
	}
	select {
	case <-ctx.Done():
		return
	case <-time.After(delay):
	}
	for {
		started := time.Now()
		l.passLocked(ctx, interval)
		select {
		case <-ctx.Done():
			return
		case <-time.After(max(0, interval-time.Since(started))):
		}
	}
}

func (l *Loop) passLocked(ctx context.Context, interval time.Duration) {
	log := l.Service.logger()
	conn, err := l.Service.Pool.Acquire(ctx)
	if err != nil {
		log.Error("directory", "op", "loop", "err", err)
		return
	}
	defer conn.Release()
	var held bool
	if err := conn.QueryRow(ctx, `select pg_try_advisory_lock($1)`, loopLock).Scan(&held); err != nil || !held {
		return
	}
	defer func() { _, _ = conn.Exec(context.WithoutCancel(ctx), `select pg_advisory_unlock($1)`, loopLock) }()
	passCtx, cancel := context.WithTimeout(ctx, interval)
	defer cancel()
	res, err := l.Service.Pass(passCtx)
	if err != nil {
		log.Error("directory", "op", "pass", "err", err)
	}
	LogPass(l.Service, res)
}

// LogPass writes one pass's counts.
func LogPass(s *Service, res PassResult) {
	s.logger().Info("directory", "op", "pass", "rosters", res.Rosters, "roster_members", res.RosterMembers,
		"profiles", res.Profiles, "private", res.Private, "deleted", res.Deleted, "misses", res.Misses,
		"failed", res.Failed, "realm_rulesets", res.RealmRulesets, "expired", res.Expired, "calls", res.Calls, "rate_limited", res.RateLimited)
}
