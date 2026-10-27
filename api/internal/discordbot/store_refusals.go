package discordbot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// SaveRefusal remembers, for ttl, the name of the guild holding the server a user was
// refused, so the next page load can say so once.
func (s *Store) SaveRefusal(ctx context.Context, userID, guildID int64, heldByName string, ttl time.Duration) error {
	return s.exec(ctx, `insert into guild_discord_refusals (user_id, guild_id, held_by_name, expires_at)
		values ($1, $2, $3, now() + make_interval(secs => $4))
		on conflict (user_id, guild_id) do update set held_by_name = excluded.held_by_name,
		  expires_at = excluded.expires_at`, userID, guildID, heldByName, ttl.Seconds())
}

// TakeRefusal deletes and returns the unexpired refusal's guild name; nil when none.
func (s *Store) TakeRefusal(ctx context.Context, userID, guildID int64) (*string, error) {
	var name string
	err := s.Pool.QueryRow(ctx, `with taken as (
		  delete from guild_discord_refusals where user_id = $1 and guild_id = $2 returning held_by_name, expires_at
		) select held_by_name from taken where expires_at > now()`, userID, guildID).Scan(&name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("discordbot: take refusal: %w", err)
	}
	return &name, nil
}
