// Package synclog is the character sync history: one character_syncs row per addon export
// taken or Battle.net import or refresh run, and the two facts GET /v1/me derives from it
// (median gap between syncs, last failed Battle.net refresh). The history is the single
// source of truth; nothing here is stored as a computed value.
package synclog

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Sources a sync can come from.
const (
	SourceAddon    = "addon"
	SourceBlizzard = "blizzard"
)

// Outcomes a sync can have. OutcomeRefreshFailed is the machine code GET /v1/me exposes as
// build.sync_error.
const (
	OutcomeOK            = "ok"
	OutcomeRefreshFailed = "bnet_refresh_failed"
)

// minSyncsForMedian is how many successful syncs the median needs: three syncs, two gaps.
const minSyncsForMedian = 3

// recentSyncWindow bounds how many of a character's newest successful syncs the median reads.
const recentSyncWindow = 30

// Execer is the write half of a pgx pool or transaction.
type Execer interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

// Querier is the read half of a pgx pool or transaction.
type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// Summary is what the history says about one character. Each field is nil when the history
// cannot answer it.
type Summary struct {
	MedianGapSec *int
	SyncError    *string
}

// Record appends one sync to a character's history.
func Record(ctx context.Context, q Execer, characterKey, source, outcome string) error {
	if source != SourceAddon && source != SourceBlizzard {
		return fmt.Errorf("synclog: unknown source %q", source)
	}
	if _, err := q.Exec(ctx,
		`insert into character_syncs (character_key, source, outcome) values ($1, $2, $3)`,
		characterKey, source, outcome); err != nil {
		return fmt.Errorf("synclog: record %s: %w", characterKey, err)
	}
	return nil
}

// Summaries reads the history of every key and derives each character's Summary. A key with
// no history has no entry.
func Summaries(ctx context.Context, q Querier, keys []string) (map[string]Summary, error) {
	out := map[string]Summary{}
	if len(keys) == 0 {
		return out, nil
	}
	okTimes, err := recentOkTimes(ctx, q, keys)
	if err != nil {
		return nil, err
	}
	for key, times := range okTimes {
		out[key] = Summary{MedianGapSec: MedianGapSec(times)}
	}
	errs, err := lastRefreshErrors(ctx, q, keys)
	if err != nil {
		return nil, err
	}
	for key, code := range errs {
		s := out[key]
		s.SyncError = &code
		out[key] = s
	}
	return out, nil
}

func recentOkTimes(ctx context.Context, q Querier, keys []string) (map[string][]time.Time, error) {
	rows, err := q.Query(ctx,
		`select character_key, created_at from (
		   select character_key, created_at,
		          row_number() over (partition by character_key order by created_at desc) as n
		   from character_syncs where character_key = any($1) and outcome = $2
		 ) recent where n <= $3`, keys, OutcomeOK, recentSyncWindow)
	if err != nil {
		return nil, fmt.Errorf("synclog: read sync times: %w", err)
	}
	defer rows.Close()
	out := map[string][]time.Time{}
	for rows.Next() {
		var key string
		var at time.Time
		if err := rows.Scan(&key, &at); err != nil {
			return nil, fmt.Errorf("synclog: scan sync times: %w", err)
		}
		out[key] = append(out[key], at)
	}
	return out, rows.Err()
}

// lastRefreshErrors answers, per key, the outcome of the latest Battle.net row when it was a
// failure. Only a later Battle.net row clears it: an addon export says nothing about whether
// Blizzard answered.
func lastRefreshErrors(ctx context.Context, q Querier, keys []string) (map[string]string, error) {
	rows, err := q.Query(ctx,
		`select distinct on (character_key) character_key, outcome from character_syncs
		 where character_key = any($1) and source = $2
		 order by character_key, created_at desc, id desc`, keys, SourceBlizzard)
	if err != nil {
		return nil, fmt.Errorf("synclog: read refresh outcomes: %w", err)
	}
	defer rows.Close()
	out := map[string]string{}
	for rows.Next() {
		var key, outcome string
		if err := rows.Scan(&key, &outcome); err != nil {
			return nil, fmt.Errorf("synclog: scan refresh outcomes: %w", err)
		}
		if outcome != OutcomeOK {
			out[key] = outcome
		}
	}
	return out, rows.Err()
}

// MedianGapSec is the median gap, in whole seconds, between successive syncs (in any input
// order), or nil when there are fewer than three. An even number of gaps takes the mean of
// the middle pair.
func MedianGapSec(syncs []time.Time) *int {
	if len(syncs) < minSyncsForMedian {
		return nil
	}
	sorted := slices.Clone(syncs)
	slices.SortFunc(sorted, func(a, b time.Time) int { return a.Compare(b) })
	gaps := make([]time.Duration, len(sorted)-1)
	for i := range gaps {
		gaps[i] = sorted[i+1].Sub(sorted[i])
	}
	slices.Sort(gaps)
	mid := len(gaps) / 2
	median := gaps[mid]
	if len(gaps)%2 == 0 {
		median = (gaps[mid-1] + gaps[mid]) / 2
	}
	sec := int(median / time.Second)
	return &sec
}
