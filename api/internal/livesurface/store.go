package livesurface

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Live surface kinds and platforms (migration live_activity_tokens).
const (
	KindRaid        = "raid"
	KindSim         = "sim"
	PlatformIOS     = "ios"
	PlatformAndroid = "android"
)

// MaxActivePerUser is how many live registrations one account keeps; registering beyond it
// drops the oldest, so a client that never deletes cannot grow the table without bound.
const MaxActivePerUser = 20

// Registration is one live surface the app asked the server to keep current.
type Registration struct {
	UserID           int64
	DeviceID         string // "" for a browser session
	Kind             string
	Platform         string
	ObjectID         string
	PushToken        string
	PushToStartToken string // "" when none
	ExpiresAt        time.Time
}

// Subscriber is a registration as the push senders read it.
type Subscriber struct {
	ID               int64
	UserID           int64
	Platform         string
	PushToken        string
	PushToStartToken string
}

// Store is the live_activity_tokens table.
type Store struct{ Pool *pgxpool.Pool }

// Register keeps reg, answering its id. Registering the same token for the same object again
// refreshes the row rather than adding one. The account's oldest rows past MaxActivePerUser go.
func (s *Store) Register(ctx context.Context, reg Registration) (int64, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `
		insert into live_activity_tokens
		  (user_id, device_id, kind, platform, object_id, push_token, push_to_start_token, expires_at)
		values ($1, nullif($2, ''), $3, $4, $5, $6, nullif($7, ''), $8)
		on conflict (kind, object_id, push_token) do update set
		  user_id = excluded.user_id, device_id = excluded.device_id, platform = excluded.platform,
		  push_to_start_token = excluded.push_to_start_token, expires_at = excluded.expires_at,
		  updated_at = now()
		returning id`,
		reg.UserID, reg.DeviceID, reg.Kind, reg.Platform, reg.ObjectID, reg.PushToken, reg.PushToStartToken,
		reg.ExpiresAt).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("livesurface: register: %w", err)
	}
	if _, err := s.Pool.Exec(ctx, `
		delete from live_activity_tokens where user_id = $1 and id not in (
		  select id from live_activity_tokens where user_id = $1
		  order by updated_at desc, id desc limit $2)`, reg.UserID, MaxActivePerUser); err != nil {
		return 0, fmt.Errorf("livesurface: trim: %w", err)
	}
	return id, nil
}

// Unregister removes the caller's own registration; another account's id is left alone.
func (s *Store) Unregister(ctx context.Context, userID, id int64) error {
	if _, err := s.Pool.Exec(ctx, `delete from live_activity_tokens where id = $1 and user_id = $2`, id, userID); err != nil {
		return fmt.Errorf("livesurface: unregister: %w", err)
	}
	return nil
}

// Subscribers lists the live registrations for one object that have not expired, for the push
// senders.
func (s *Store) Subscribers(ctx context.Context, kind, objectID string, now time.Time) ([]Subscriber, error) {
	rows, err := s.Pool.Query(ctx, `
		select id, user_id, platform, push_token, coalesce(push_to_start_token, '')
		from live_activity_tokens where kind = $1 and object_id = $2 and expires_at > $3
		order by id`, kind, objectID, now)
	if err != nil {
		return nil, fmt.Errorf("livesurface: subscribers: %w", err)
	}
	defer rows.Close()
	out := []Subscriber{}
	for rows.Next() {
		var sub Subscriber
		if err := rows.Scan(&sub.ID, &sub.UserID, &sub.Platform, &sub.PushToken, &sub.PushToStartToken); err != nil {
			return nil, fmt.Errorf("livesurface: subscribers: %w", err)
		}
		out = append(out, sub)
	}
	return out, rows.Err()
}

// Drop removes a registration whose token the push service refused.
func (s *Store) Drop(ctx context.Context, id int64) error {
	if _, err := s.Pool.Exec(ctx, `delete from live_activity_tokens where id = $1`, id); err != nil {
		return fmt.Errorf("livesurface: drop: %w", err)
	}
	return nil
}

// DeleteExpired removes every registration whose object has ended, answering how many went.
func (s *Store) DeleteExpired(ctx context.Context, now time.Time) (int64, error) {
	tag, err := s.Pool.Exec(ctx, `delete from live_activity_tokens where expires_at <= $1`, now)
	if err != nil {
		return 0, fmt.Errorf("livesurface: delete expired: %w", err)
	}
	return tag.RowsAffected(), nil
}
