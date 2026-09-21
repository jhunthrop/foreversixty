package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PartitionJob keeps every table in Tables partitioned ahead of today. Run it at startup;
// Run does the first pass itself and sweeps in the background until the context is
// cancelled.
type PartitionJob struct {
	Pool  *pgxpool.Pool
	Log   *slog.Logger
	Every time.Duration
	Ahead int
	// Tables is which partitioned tables to sweep. Empty means {"fight_metrics"}, so
	// every caller and test that predates the ratings lane keeps its exact old behaviour
	// with no change.
	Tables []string
	// Now is the clock, so a test can drive the job from a fixed date.
	Now func() time.Time
}

func (j *PartitionJob) tables() []string {
	if len(j.Tables) == 0 {
		return []string{"fight_metrics"}
	}
	return j.Tables
}

// Run creates every table's partitions once and then again on every tick, until ctx is
// cancelled. The first pass's error is returned so a startup can fail loudly on a database
// that cannot be partitioned at all; later failures are logged, because by then the
// service is serving traffic.
func (j *PartitionJob) Run(ctx context.Context) error {
	every, ahead, now := j.Every, j.Ahead, j.Now
	if every <= 0 {
		every = PartitionSweepEvery
	}
	if ahead <= 0 {
		ahead = PartitionMonthsAhead
	}
	if now == nil {
		now = time.Now
	}
	log := j.Log
	if log == nil {
		log = slog.Default()
	}
	tables := j.tables()
	if err := sweepOnce(ctx, j.Pool, tables, now(), ahead); err != nil {
		return err
	}
	go func() {
		t := time.NewTicker(every)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				if err := sweepOnce(ctx, j.Pool, tables, now(), ahead); err != nil {
					log.Error("partitions", "err", err)
				}
			}
		}
	}()
	return nil
}

func sweepOnce(ctx context.Context, pool *pgxpool.Pool, tables []string, at time.Time, ahead int) error {
	for _, table := range tables {
		if err := EnsurePartitions(ctx, pool, table, at.UTC(), ahead); err != nil {
			return err
		}
	}
	return nil
}
