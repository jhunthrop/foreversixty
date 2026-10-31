package push

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Token is one registered push token.
type Token struct {
	Token      string
	Platform   Platform
	UserID     int64
	DeviceID   string
	AppVersion string
}

// Store is the push_tokens table (migration 0062).
type Store struct {
	Pool *pgxpool.Pool
}

// Register saves t for its device: a token already saved moves to this device and account,
// and the device's earlier tokens go (an app install holds one token at a time). Nothing is
// saved for a device that is revoked.
func (s *Store) Register(ctx context.Context, t Token) error {
	return pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `delete from push_tokens where device_id = $1 and token <> $2`, t.DeviceID, t.Token); err != nil {
			return fmt.Errorf("push: replace device token: %w", err)
		}
		_, err := tx.Exec(ctx, `insert into push_tokens (token, platform, user_id, device_id, app_version)
			select $1, $2, $3, d.id, $5 from devices d where d.id = $4 and d.user_id = $3 and d.revoked_at is null
			on conflict (token) do update set platform = excluded.platform, user_id = excluded.user_id,
			  device_id = excluded.device_id, app_version = excluded.app_version, updated_at = now()`,
			t.Token, string(t.Platform), t.UserID, t.DeviceID, t.AppVersion)
		if err != nil {
			return fmt.Errorf("push: register token: %w", err)
		}
		return nil
	})
}

// Unregister deletes token when the device registered it; false when it had not.
func (s *Store) Unregister(ctx context.Context, deviceID, token string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `delete from push_tokens where token = $1 and device_id = $2`, token, deviceID)
	if err != nil {
		return false, fmt.Errorf("push: unregister token: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}

// TokensOf is every token of the account on a live device.
func (s *Store) TokensOf(ctx context.Context, userID int64) ([]Token, error) {
	rows, err := s.Pool.Query(ctx, `select t.token, t.platform, t.user_id, t.device_id, t.app_version
		from push_tokens t join devices d on d.id = t.device_id
		where t.user_id = $1 and d.revoked_at is null order by t.created_at`, userID)
	if err != nil {
		return nil, fmt.Errorf("push: tokens of user: %w", err)
	}
	out, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Token, error) {
		var t Token
		var platform string
		err := row.Scan(&t.Token, &platform, &t.UserID, &t.DeviceID, &t.AppVersion)
		t.Platform = Platform(platform)
		return t, err
	})
	if err != nil {
		return nil, fmt.Errorf("push: tokens of user: %w", err)
	}
	return out, nil
}

// Delete drops a token the push service called invalid.
func (s *Store) Delete(ctx context.Context, token string) error {
	if _, err := s.Pool.Exec(ctx, `delete from push_tokens where token = $1`, token); err != nil {
		return fmt.Errorf("push: delete token: %w", err)
	}
	return nil
}
