package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PartitionMonthsAhead is how many months of partitions the service keeps in front of
// today. Three is a quarter of slack: the partition job runs daily, so two whole runs can
// fail before an insert has nowhere to land.
const PartitionMonthsAhead = 3

// PartitionSweepEvery is how often the background partition job runs.
const PartitionSweepEvery = 24 * time.Hour

// RatingsTable is the rating_scores partitioned table (migration 0022), so a caller
// outside this package never hand-spells it.
const RatingsTable = "rating_scores"

// partitionName is the table name for the month containing at, for any range-partitioned
// table keyed the same way fight_metrics and rating_scores both are.
func partitionName(table string, at time.Time) string {
	return fmt.Sprintf("%s_%s", table, at.UTC().Format("2006_01"))
}

// MetricsPartition is the table name for the month containing at, for fight_metrics.
// Exported for the existing tests and callers that already name it directly.
func MetricsPartition(at time.Time) string {
	return partitionName("fight_metrics", at)
}

// EnsurePartition creates the monthly partition of table holding at, if it is not there
// already. Partitions are created one month at a time and never dropped. table is always a
// fixed, code-provided constant (fight_metrics or RatingsTable today), never
// request-derived, so building the statement by Sprintf carries no injection risk - the
// same trade this function's fight_metrics-only predecessor already made.
func EnsurePartition(ctx context.Context, pool *pgxpool.Pool, table string, at time.Time) error {
	start := time.Date(at.UTC().Year(), at.UTC().Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	stmt := fmt.Sprintf(
		`create table if not exists %s partition of %s for values from ('%s') to ('%s')`,
		partitionName(table, start), table, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if _, err := pool.Exec(ctx, stmt); err != nil {
		return fmt.Errorf("db: create partition %s: %w", partitionName(table, start), err)
	}
	return nil
}

// EnsureMetricsPartition is EnsurePartition pinned to fight_metrics, kept as the name every
// existing call site already uses.
func EnsureMetricsPartition(ctx context.Context, pool *pgxpool.Pool, at time.Time) error {
	return EnsurePartition(ctx, pool, "fight_metrics", at)
}

// EnsureRatingsPartition is EnsurePartition pinned to rating_scores.
func EnsureRatingsPartition(ctx context.Context, pool *pgxpool.Pool, at time.Time) error {
	return EnsurePartition(ctx, pool, RatingsTable, at)
}

// EnsurePartitions creates the partition of table for the month containing from and the
// months months after it. Idempotent.
func EnsurePartitions(ctx context.Context, pool *pgxpool.Pool, table string, from time.Time, months int) error {
	for i := 0; i <= months; i++ {
		if err := EnsurePartition(ctx, pool, table, from.AddDate(0, i, 0)); err != nil {
			return err
		}
	}
	return nil
}

// EnsureMetricsPartitions is EnsurePartitions pinned to fight_metrics, kept as the name
// every existing call site already uses.
func EnsureMetricsPartitions(ctx context.Context, pool *pgxpool.Pool, from time.Time, months int) error {
	return EnsurePartitions(ctx, pool, "fight_metrics", from, months)
}
