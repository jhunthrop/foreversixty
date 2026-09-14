package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PartitionMonthsAhead is how many months of fight_metrics partitions the
// service keeps in front of today. Three is a quarter of slack: the
// partition job runs daily, so two whole runs can fail before an insert
// has nowhere to land.
const PartitionMonthsAhead = 3

// PartitionSweepEvery is how often the background partition job runs.
const PartitionSweepEvery = 24 * time.Hour

// MetricsPartition is the table name for the month containing at.
func MetricsPartition(at time.Time) string {
	return fmt.Sprintf("fight_metrics_%s", at.UTC().Format("2006_01"))
}

// EnsureMetricsPartition creates the monthly partition holding at, if it
// is not there already. Partitions are created one month at a time and
// never dropped: the design keeps rankings rows forever.
func EnsureMetricsPartition(ctx context.Context, pool *pgxpool.Pool, at time.Time) error {
	start := time.Date(at.UTC().Year(), at.UTC().Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	stmt := fmt.Sprintf(
		`create table if not exists %s partition of fight_metrics for values from ('%s') to ('%s')`,
		MetricsPartition(start), start.Format(time.RFC3339), end.Format(time.RFC3339))
	if _, err := pool.Exec(ctx, stmt); err != nil {
		return fmt.Errorf("db: create partition %s: %w", MetricsPartition(start), err)
	}
	return nil
}

// EnsureMetricsPartitions creates the partition for the month containing
// from and the months months after it. Call it at startup and from the
// partition job; it is idempotent.
func EnsureMetricsPartitions(ctx context.Context, pool *pgxpool.Pool, from time.Time, months int) error {
	for i := 0; i <= months; i++ {
		if err := EnsureMetricsPartition(ctx, pool, from.AddDate(0, i, 0)); err != nil {
			return err
		}
	}
	return nil
}
