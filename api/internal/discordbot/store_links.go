package discordbot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ErrLinkConflict is returned when the Discord account is already linked to another account.
var ErrLinkConflict = errors.New("discordbot: discord account already linked to another account")

// Link is an account's one Discord identity.
type Link struct {
	UserID        int64
	DiscordUserID string
	DisplayName   string
	LinkedAt      time.Time
}

// SaveLink links userID to the Discord account, replacing the account's earlier link.
// A Discord account linked to a different user is ErrLinkConflict.
func (s *Store) SaveLink(ctx context.Context, userID int64, discordUserID, displayName string) error {
	_, err := s.Pool.Exec(ctx, `insert into discord_links (user_id, discord_user_id, display_name)
		values ($1, $2, $3)
		on conflict (user_id) do update set discord_user_id = excluded.discord_user_id,
		  display_name = excluded.display_name, linked_at = now()`, userID, discordUserID, displayName)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
		return ErrLinkConflict
	}
	if err != nil {
		return fmt.Errorf("discordbot: save link: %w", err)
	}
	return nil
}

// LinkByUser returns the account's link; ok is false when none exists.
func (s *Store) LinkByUser(ctx context.Context, userID int64) (Link, bool, error) {
	l := Link{UserID: userID}
	err := s.Pool.QueryRow(ctx, `select discord_user_id, display_name, linked_at from discord_links
		where user_id = $1`, userID).Scan(&l.DiscordUserID, &l.DisplayName, &l.LinkedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Link{}, false, nil
	}
	if err != nil {
		return Link{}, false, fmt.Errorf("discordbot: link by user: %w", err)
	}
	return l, true, nil
}

// UserByDiscord returns the account linked to the Discord user; ok is false when none is.
func (s *Store) UserByDiscord(ctx context.Context, discordUserID string) (int64, bool, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `select user_id from discord_links where discord_user_id = $1`, discordUserID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("discordbot: user by discord: %w", err)
	}
	return id, true, nil
}

// DeleteLink removes the account's link; an absent one is not an error.
func (s *Store) DeleteLink(ctx context.Context, userID int64) error {
	return s.exec(ctx, `delete from discord_links where user_id = $1`, userID)
}

// LinkedMembers counts the guild's verified accounts and how many of them are linked.
func (s *Store) LinkedMembers(ctx context.Context, guildID int64) (linked, members int, err error) {
	err = s.Pool.QueryRow(ctx, `select count(l.user_id), count(*)
		from guild_members m left join discord_links l on l.user_id = m.user_id
		where m.guild_id = $1 and m.verified_at is not null`, guildID).Scan(&linked, &members)
	if err != nil {
		return 0, 0, fmt.Errorf("discordbot: linked members: %w", err)
	}
	return linked, members, nil
}

// SuggestedGuild names a guild where the account is a verified member and the bot has a
// channel and is still in the server; nil when there is none.
func (s *Store) SuggestedGuild(ctx context.Context, userID int64) (*string, error) {
	var name string
	err := s.Pool.QueryRow(ctx, `select g.name from guild_members m
		join guilds g on g.id = m.guild_id
		join guild_discord_installs i on i.guild_id = m.guild_id
		where m.user_id = $1 and m.verified_at is not null
		  and i.channel_id is not null and not i.removed_in_discord
		order by g.name limit 1`, userID).Scan(&name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("discordbot: suggested guild: %w", err)
	}
	return &name, nil
}
