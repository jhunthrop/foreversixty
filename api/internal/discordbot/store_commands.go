package discordbot

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// commandsLockKey is the advisory lock that keeps two starting instances from registering at
// once ("DSCMDS" as a number).
const commandsLockKey int64 = 0x44534d4453

// WithCommandsLock runs fn holding the registration lock on one connection; the lock is
// released when fn returns, and a failure to take it is an error.
func (s *Store) WithCommandsLock(ctx context.Context, fn func(ctx context.Context) error) error {
	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("discordbot: commands lock connection: %w", err)
	}
	defer conn.Release()
	if _, err := conn.Exec(ctx, `select pg_advisory_lock($1)`, commandsLockKey); err != nil {
		return fmt.Errorf("discordbot: commands lock: %w", err)
	}
	runErr := fn(ctx)
	// The unlock must run even when ctx ended, or the pooled connection keeps the lock.
	if _, err := conn.Exec(context.WithoutCancel(ctx), `select pg_advisory_unlock($1)`, commandsLockKey); err != nil {
		conn.Conn().PgConn().Close(context.WithoutCancel(ctx)) //nolint:errcheck // dropping the connection frees the lock
		return errors.Join(runErr, fmt.Errorf("discordbot: commands unlock: %w", err))
	}
	return runErr
}

// CommandsHash is the hash last registered for the application's scope; "" when none.
func (s *Store) CommandsHash(ctx context.Context, applicationID, scope string) (string, error) {
	var hash string
	err := s.Pool.QueryRow(ctx,
		`select hash from discord_command_sets where application_id = $1 and scope = $2`,
		applicationID, scope).Scan(&hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("discordbot: commands hash: %w", err)
	}
	return hash, nil
}

// SaveCommandsHash records what was registered for the application's scope.
func (s *Store) SaveCommandsHash(ctx context.Context, applicationID, scope, hash string) error {
	_, err := s.Pool.Exec(ctx, `insert into discord_command_sets (application_id, scope, hash)
		values ($1, $2, $3)
		on conflict (application_id, scope) do update set hash = excluded.hash, registered_at = now()`,
		applicationID, scope, hash)
	if err != nil {
		return fmt.Errorf("discordbot: save commands hash: %w", err)
	}
	return nil
}
