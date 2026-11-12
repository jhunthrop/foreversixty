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
// (attributedGuild) in reports.matched_guild_id, and returns the report's guild for matching
// as it now stands: coalesce(guild_id, matched_guild_id) (reportguild.Of). It never writes
// guild_id, which only the officer PATCH sets because it carries edit rights, the guild
// leaderboards and character verification. matched_guild_id is read through reportguild by the
// guild matcher, a member's own guild figures (guilds.memberLogIs: roster best, rating, last
// log) and the guild's member-facing lists (the /logs guild tab, ForGuild; the guild page's
// recent logs and Raids tab; the recent feed's guild name), nothing else. A guild already set,
// by hand or by an earlier attribution, is never overridden.
func (s *Store) AttributeGuild(ctx context.Context, reportID string) (*int64, error) {
	if _, err := s.Pool.Exec(ctx,
		`update reports set matched_guild_id = `+attributedGuild+`
		  where id = $1 and guild_id is null and matched_guild_id is null and `+attributedGuild+` is not null`, reportID); err != nil {
		return nil, fmt.Errorf("reports: attribute guild of %s: %w", reportID, err)
	}
	var guildID *int64
	err := s.Pool.QueryRow(ctx, `select coalesce(guild_id, matched_guild_id) from reports where id = $1`, reportID).Scan(&guildID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("reports: read guild of %s: %w", reportID, err)
	}
	return guildID, nil
}

// CreditLogger names key as the report's logging character when it names none yet: a
// whole-file upload's writer, read from the log. A character already named is never replaced.
func (s *Store) CreditLogger(ctx context.Context, reportID, key string) error {
	if _, err := s.Pool.Exec(ctx,
		`update reports set logging_character = $2 where id = $1 and logging_character is null`,
		reportID, key); err != nil {
		return fmt.Errorf("reports: credit logger of %s: %w", reportID, err)
	}
	return nil
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
