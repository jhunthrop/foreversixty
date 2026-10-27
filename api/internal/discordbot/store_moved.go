package discordbot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// MovedFrom is the start the raid was moved from, as its post shows it; nil when the raid has
// no post row or was never moved.
func (s *Store) MovedFrom(ctx context.Context, raidID int64) (*time.Time, error) {
	var was *time.Time
	err := s.Pool.QueryRow(ctx, `select moved_from from guild_raid_discord_posts where raid_id = $1 and not gone`,
		raidID).Scan(&was)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("discordbot: moved from: %w", err)
	}
	return was, nil
}
