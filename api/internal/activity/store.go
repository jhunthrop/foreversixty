package activity

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Retention is how long a row is kept; the feed itself reads FeedWindow back.
const (
	Retention  = 30 * 24 * time.Hour
	FeedWindow = 7 * 24 * time.Hour
)

// Store is the activity_events table and the snapshot tables beside it.
type Store struct {
	Pool *pgxpool.Pool
}

type execer interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

const insertEvent = `insert into activity_events (kind, occurred_at, spec, band, guild_id, character_key,
	icon, crest, quality, href, date_only, segments, dedupe_key)
	values ($1, $2, nullif($3, ''), nullif($4, 0), nullif($5, 0), nullif($6, ''), $7, $8, $9, $10, $11, $12, $13)`

func eventArgs(e Event) ([]any, error) {
	if e.DedupeKey == "" || KindTag[e.Kind] == "" {
		return nil, fmt.Errorf("activity: event %q needs a known kind and a dedupe key", e.DedupeKey)
	}
	segments, err := json.Marshal(e.Segments)
	if err != nil {
		return nil, fmt.Errorf("activity: encode %s: %w", e.DedupeKey, err)
	}
	return []any{e.Kind, e.OccurredAt.UTC(), e.Spec, e.Band, e.GuildID, e.CharacterKey, e.Icon,
		e.Crest, e.Quality, e.Href, e.DateOnly, segments, e.DedupeKey}, nil
}

// Append writes e once: a second event with the same dedupe key is dropped. It reports
// whether the row was new.
func (s *Store) Append(ctx context.Context, e Event) (bool, error) {
	return appendOn(ctx, s.Pool, e)
}

func appendOn(ctx context.Context, db execer, e Event) (bool, error) {
	args, err := eventArgs(e)
	if err != nil {
		return false, err
	}
	tag, err := db.Exec(ctx, insertEvent+` on conflict (dedupe_key) do nothing`, args...)
	if err != nil {
		return false, fmt.Errorf("activity: append %s: %w", e.DedupeKey, err)
	}
	return tag.RowsAffected() == 1, nil
}

// Upsert writes e, or rewrites the sentence of the row with its dedupe key: a guild raid's
// counts as members sign up (keepTime: the row keeps its post time), a member's second sim
// of the day (the row moves to the new time).
func (s *Store) Upsert(ctx context.Context, e Event, keepTime bool) error {
	args, err := eventArgs(e)
	if err != nil {
		return err
	}
	at := "excluded.occurred_at"
	if keepTime {
		at = "activity_events.occurred_at"
	}
	_, err = s.Pool.Exec(ctx, insertEvent+` on conflict (dedupe_key) do update set
		segments = excluded.segments, icon = excluded.icon, crest = excluded.crest, href = excluded.href,
		quality = excluded.quality, occurred_at = `+at, args...)
	if err != nil {
		return fmt.Errorf("activity: upsert %s: %w", e.DedupeKey, err)
	}
	return nil
}

// Exists reports whether a row with the dedupe key is stored.
func (s *Store) Exists(ctx context.Context, dedupeKey string) (bool, error) {
	var found bool
	err := s.Pool.QueryRow(ctx, `select exists (select 1 from activity_events where dedupe_key = $1)`, dedupeKey).Scan(&found)
	if err != nil {
		return false, fmt.Errorf("activity: exists %s: %w", dedupeKey, err)
	}
	return found, nil
}

// Purge deletes the rows older than Retention.
func (s *Store) Purge(ctx context.Context, now time.Time) (int64, error) {
	tag, err := s.Pool.Exec(ctx, `delete from activity_events where occurred_at < $1`, now.Add(-Retention))
	if err != nil {
		return 0, fmt.Errorf("activity: purge: %w", err)
	}
	return tag.RowsAffected(), nil
}
