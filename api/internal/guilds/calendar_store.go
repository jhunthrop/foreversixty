// api/internal/guilds/calendar_store.go
//
// The calendar link's data layer (migration 0063, W3-C17): one live feed per account, keyed
// by the SHA-256 of its token, and the raids a feed lists. The raw token never reaches this
// file: callers hash it first (auth.TokenHash).
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// calendarFeedPast is how far back a feed keeps a raid: a calendar keeps last week's raids
// and the player sees what they answered.
const calendarFeedPast = 30 * 24 * time.Hour

// calendarFeedMaxRaids bounds one feed's events; a guild holds at most RaidCap open raids, so
// this is a ceiling for a player in many guilds, never a page.
const calendarFeedMaxRaids = 200

// calendarFeedLockClass is the first key of the advisory lock a reset takes per account.
const calendarFeedLockClass = 6301

// CalendarFeed is an account's live feed row.
type CalendarFeed struct {
	UserID    int64
	CreatedAt time.Time
}

// CreateCalendarFeed inserts a live feed with tokenHash for userID unless one is live already.
// created is false (and the feed is the live one) when the account had a feed: its token was
// shown when it was made and is not stored, so the caller cannot show it again.
func (s *Store) CreateCalendarFeed(ctx context.Context, userID int64, tokenHash []byte) (feed CalendarFeed, created bool, err error) {
	feed = CalendarFeed{UserID: userID}
	err = s.Pool.QueryRow(ctx, `
		insert into calendar_feeds (user_id, token_hash) values ($1, $2)
		on conflict (user_id) where revoked_at is null do nothing
		returning created_at`, userID, tokenHash).Scan(&feed.CreatedAt)
	if err == nil {
		return feed, true, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return CalendarFeed{}, false, fmt.Errorf("guilds: create calendar feed: %w", err)
	}
	err = s.Pool.QueryRow(ctx,
		`select created_at from calendar_feeds where user_id = $1 and revoked_at is null`, userID).Scan(&feed.CreatedAt)
	if err != nil {
		return CalendarFeed{}, false, fmt.Errorf("guilds: read calendar feed: %w", err)
	}
	return feed, false, nil
}

// ResetCalendarFeed revokes userID's live feed, if any, and makes a new one with tokenHash, in
// one transaction: the old link stops working the moment the new one exists.
func (s *Store) ResetCalendarFeed(ctx context.Context, userID int64, tokenHash []byte) (CalendarFeed, error) {
	feed := CalendarFeed{UserID: userID}
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		// One reset per account at a time: two concurrent resets would otherwise both revoke
		// the old row and the second insert would hit the one-live-link index.
		if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock($1, $2)`, calendarFeedLockClass, userID); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx,
			`update calendar_feeds set revoked_at = now() where user_id = $1 and revoked_at is null`, userID); err != nil {
			return err
		}
		return tx.QueryRow(ctx,
			`insert into calendar_feeds (user_id, token_hash) values ($1, $2) returning created_at`,
			userID, tokenHash).Scan(&feed.CreatedAt)
	})
	if err != nil {
		return CalendarFeed{}, fmt.Errorf("guilds: reset calendar feed: %w", err)
	}
	return feed, nil
}

// CalendarFeedUser is the account behind a live feed's token hash; ok is false for an unknown
// or revoked token.
func (s *Store) CalendarFeedUser(ctx context.Context, tokenHash []byte) (userID int64, ok bool, err error) {
	err = s.Pool.QueryRow(ctx,
		`select user_id from calendar_feeds where token_hash = $1 and revoked_at is null`, tokenHash).Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("guilds: calendar feed user: %w", err)
	}
	return userID, true, nil
}

// CalendarRaid is one raid as a feed lists it.
type CalendarRaid struct {
	Raid
	Guild    Guild
	Timezone string
	// SignedUp counts the raid's In and Late answers.
	SignedUp int
	// State is the account's own answer (its best one when two of its characters answered:
	// In, then Late, Bench, Out), empty when none of its characters did.
	State string
}

// CalendarRaids lists the raids of every guild where userID is a verified member, as of now:
// planned, locked and done raids that started within calendarFeedPast or start later.
// Cancelled raids are left out, so a subscribed calendar drops them on its next fetch.
// Membership is read here, on every fetch: a player who leaves a guild stops seeing its raids.
func (s *Store) CalendarRaids(ctx context.Context, userID int64, now time.Time) ([]CalendarRaid, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.id, r.guild_id, r.zone, r.title, r.starts_at, r.duration_minutes, r.status, r.updated_at,
		       g.name, g.region, g.ruleset, coalesce(g.raid_timezone, ''),
		       (select count(*) from guild_raid_signups s where s.raid_id = r.id and s.state in ('in', 'late')),
		       coalesce((select s.state from guild_raid_signups s where s.raid_id = r.id and s.signed_by = $1
		                 order by array_position(array['in','late','bench','out'], s.state) limit 1), '')
		from guild_raids r
		join guild_members gm on gm.guild_id = r.guild_id and gm.user_id = $1 and gm.verified_at is not null
		join guilds g on g.id = r.guild_id
		where r.status in ('planned', 'locked', 'done') and r.starts_at > $2
		order by r.starts_at, r.id
		limit $3`, userID, now.Add(-calendarFeedPast), calendarFeedMaxRaids)
	if err != nil {
		return nil, fmt.Errorf("guilds: calendar raids: %w", err)
	}
	defer rows.Close()
	out := []CalendarRaid{}
	for rows.Next() {
		var c CalendarRaid
		if err := rows.Scan(&c.ID, &c.GuildID, &c.Zone, &c.Title, &c.StartsAt, &c.DurationMinutes, &c.Status,
			&c.UpdatedAt, &c.Guild.Name, &c.Guild.Region, &c.Guild.Ruleset, &c.Timezone, &c.SignedUp, &c.State); err != nil {
			return nil, fmt.Errorf("guilds: calendar raids: %w", err)
		}
		c.Guild.ID = c.GuildID
		out = append(out, c)
	}
	return out, rows.Err()
}
