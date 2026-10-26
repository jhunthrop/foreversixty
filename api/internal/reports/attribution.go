package reports

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// attributedGuild is the guild a report belongs to by its owner's standing, or null:
//
//  1. the logging character's verified guild_characters row, when the report's owner is that
//     character's account (a client names the logging character, so a character someone else
//     owns never hands its guild a stranger's log);
//  2. else the owner's verified guild, when the owner has exactly one;
//  3. else null - an owner verified in two guilds with an unverified logger is not guessed at.
const attributedGuild = `coalesce(
	(select gc.guild_id from reports r
	   join guild_characters gc on gc.character_key = r.logging_character
	    and gc.user_id = r.owner_id and gc.verified_at is not null
	  where r.id = $1),
	(select min(gc.guild_id) from reports r
	   join guild_characters gc on gc.user_id = r.owner_id and gc.verified_at is not null
	  where r.id = $1
	 having count(distinct gc.guild_id) = 1))`

// AttributeGuild gives a report with no guild the guild its owner's standing names
// (attributedGuild) and returns the report's guild as it now stands. A guild already set, by
// hand or by an earlier attribution, is never overridden.
func (s *Store) AttributeGuild(ctx context.Context, reportID string) (*int64, error) {
	if _, err := s.Pool.Exec(ctx,
		`update reports set guild_id = `+attributedGuild+`
		  where id = $1 and guild_id is null and `+attributedGuild+` is not null`, reportID); err != nil {
		return nil, fmt.Errorf("reports: attribute guild of %s: %w", reportID, err)
	}
	var guildID *int64
	err := s.Pool.QueryRow(ctx, `select guild_id from reports where id = $1`, reportID).Scan(&guildID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("reports: read guild of %s: %w", reportID, err)
	}
	return guildID, nil
}

// SetLogZone records the zone the engine read from the log. An empty zone says nothing and
// writes nothing; the same zone again is a no-op.
func (s *Store) SetLogZone(ctx context.Context, reportID, zone string) error {
	if zone == "" {
		return nil
	}
	if _, err := s.Pool.Exec(ctx,
		`update reports set log_zone = $2 where id = $1 and log_zone is distinct from $2`,
		reportID, zone); err != nil {
		return fmt.Errorf("reports: set log zone of %s: %w", reportID, err)
	}
	return nil
}
