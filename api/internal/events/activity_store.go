package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// The Activity inbox (migration 0084): Store implements Inbox and the Activity routes' reads.

var _ Inbox = (*Store)(nil)

// ActivityRetention is how long an Activity row is kept.
const ActivityRetention = 90 * 24 * time.Hour

// ActivityRow is one inbox row as the routes show it.
type ActivityRow struct {
	ID         int64     `json:"id"`
	Kind       Kind      `json:"kind"`
	Group      Group     `json:"group"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Link       string    `json:"link"`
	OccurredAt time.Time `json:"occurred_at"`
	Read       bool      `json:"read"`
	// Icon is the row's icon (IconOf): its producer's, else its group's; never empty.
	Icon string `json:"icon"`
	// Segments, Crest and Figure are the event's (Event); absent on a row without them.
	Segments []Segment `json:"segments,omitempty"`
	Crest    *string   `json:"crest,omitempty"`
	Figure   *float64  `json:"figure,omitempty"`
}

// ActivityQuery is one page request. An empty Group is every group; Before 0 is the newest
// page.
type ActivityQuery struct {
	Group  Group
	Before int64
	Limit  int
}

// Record keeps e once in each account's inbox; a republished event (the same dedupe key) is
// not added again.
func (s *Store) Record(ctx context.Context, userIDs []int64, e Event) error {
	if len(userIDs) == 0 {
		return nil
	}
	segments, err := segmentsJSON(e.Segments)
	if err != nil {
		return fmt.Errorf("events: record activity: %w", err)
	}
	_, err = s.Pool.Exec(ctx, `insert into user_activity (user_id, kind, title, body, link, dedupe_key, icon, segments, crest, figure)
		select uid, $2, $3, $4, $5, $6, $7, $8, nullif($9, ''), $10 from unnest($1::bigint[]) as uid
		on conflict (user_id, dedupe_key) do nothing`,
		userIDs, string(e.Kind), e.Title, e.activityBody(), e.Link, e.DedupeKey, IconOf(e.Kind, e.Icon), segments, e.Crest, e.Figure)
	if err != nil {
		return fmt.Errorf("events: record activity: %w", err)
	}
	return nil
}

// segmentsJSON is the segments as stored: nil (SQL null) for none.
func segmentsJSON(segs []Segment) ([]byte, error) {
	if len(segs) == 0 {
		return nil, nil
	}
	return json.Marshal(segs)
}

// Activity is one page of the account's inbox, newest first, and the cursor of the next page
// (nil on the last).
func (s *Store) Activity(ctx context.Context, userID int64, q ActivityQuery) ([]ActivityRow, *int64, error) {
	kinds := []string{}
	for _, k := range KindsOf(q.Group) {
		kinds = append(kinds, string(k))
	}
	rows, err := s.Pool.Query(ctx, `select id, kind, title, body, link, occurred_at, read_at is not null, coalesce(icon, ''),
		       segments, crest, figure
		from user_activity
		where user_id = $1 and (not $2 or kind = any($3)) and ($4::bigint = 0 or id < $4::bigint)
		order by id desc limit $5`, userID, q.Group != "", kinds, q.Before, q.Limit+1)
	if err != nil {
		return nil, nil, fmt.Errorf("events: read activity: %w", err)
	}
	out, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (ActivityRow, error) {
		var a ActivityRow
		var kind string
		var segments []byte
		if err := row.Scan(&a.ID, &kind, &a.Title, &a.Body, &a.Link, &a.OccurredAt, &a.Read, &a.Icon,
			&segments, &a.Crest, &a.Figure); err != nil {
			return a, err
		}
		a.Kind, a.Group, a.Icon = Kind(kind), GroupOf(Kind(kind)), IconOf(Kind(kind), a.Icon)
		if len(segments) > 0 {
			if err := json.Unmarshal(segments, &a.Segments); err != nil {
				return a, fmt.Errorf("segments of row %d: %w", a.ID, err)
			}
		}
		return a, nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("events: read activity: %w", err)
	}
	if len(out) <= q.Limit {
		return out, nil, nil
	}
	out = out[:q.Limit]
	next := out[len(out)-1].ID
	return out, &next, nil
}

// UnreadActivity is the account's whole unread count.
func (s *Store) UnreadActivity(ctx context.Context, userID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `select count(*) from user_activity where user_id = $1 and read_at is null`,
		userID).Scan(&n); err != nil {
		return 0, fmt.Errorf("events: count unread activity: %w", err)
	}
	return n, nil
}

// MarkActivityRead marks every row of the account with id <= through read, and answers the
// unread count left.
func (s *Store) MarkActivityRead(ctx context.Context, userID, through int64) (int, error) {
	if _, err := s.Pool.Exec(ctx, `update user_activity set read_at = now()
		where user_id = $1 and id <= $2 and read_at is null`, userID, through); err != nil {
		return 0, fmt.Errorf("events: mark activity read: %w", err)
	}
	return s.UnreadActivity(ctx, userID)
}

// PruneActivity forgets rows older than the retention, at most limit rows per call.
func (s *Store) PruneActivity(ctx context.Context, now time.Time, limit int) (int64, error) {
	tag, err := s.Pool.Exec(ctx, `delete from user_activity where id in (
		select id from user_activity where occurred_at < $1 limit $2)`, now.Add(-ActivityRetention), limit)
	if err != nil {
		return 0, fmt.Errorf("events: prune activity: %w", err)
	}
	return tag.RowsAffected(), nil
}
