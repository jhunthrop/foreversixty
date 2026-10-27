package seats

import (
	"context"
	"fmt"
	"time"
)

// Officer is a verified officer or leader of a guild, named by the character
// that stands for the account there.
type Officer struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

// Officers returns guildID's verified officers and leaders, by name.
func (s *Store) Officers(ctx context.Context, guildID int64) ([]Officer, error) {
	rows, err := s.Pool.Query(ctx, `
		select m.user_id, coalesce(ch.name, '')
		from guild_members m
		left join lateral `+bestCharacterLateral+` on true
		where m.guild_id = $1 and m.verified_at is not null and m.rank in ('officer', 'leader')
		order by lower(coalesce(ch.name, '')), m.user_id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("seats: officers of %d: %w", guildID, err)
	}
	defer rows.Close()
	out := []Officer{}
	for rows.Next() {
		var o Officer
		if err := rows.Scan(&o.UserID, &o.Name); err != nil {
			return nil, fmt.Errorf("seats: scan officer: %w", err)
		}
		out = append(out, o)
	}
	return out, rows.Err()
}

// MemberName is the name of the character that stands for userID in guildID,
// falling back to the account's main anywhere; "" when it has no character.
func (s *Store) MemberName(ctx context.Context, guildID, userID int64) (string, error) {
	ch, err := characterOf(ctx, s.Pool, guildID, userID)
	if err != nil {
		return "", err
	}
	return ch.name, nil
}

// BillingReader is what the guild billing read needs from seats: the store's
// reads, and the per-guild sweep it runs before reading.
type BillingReader struct {
	*Store
	Sweeper *Service
}

// SweepGuild releases the guild's seats that no longer stand.
func (r BillingReader) SweepGuild(ctx context.Context, guildID int64, now time.Time) error {
	return r.Sweeper.SweepGuild(ctx, guildID, now)
}
