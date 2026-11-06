// api/internal/discordbot/ping.go
//
// A raid's "Ping on Discord" (guilds.RaidPinger, spec 2026-11-05-web-layout §7.1, D16): one
// message, a reply under the raid's live post when it has one (else in the bot's channel),
// that pings only the non-signers it names. Never a DM.
package discordbot

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

var _ guilds.RaidPinger = (*RaidRouter)(nil)

// PingRaid posts p for the guild; false, with no error, when the bot cannot post for it.
func (r *RaidRouter) PingRaid(ctx context.Context, guildID int64, p guilds.RaidPing) (bool, error) {
	channel, ok := r.activeChannel(ctx, guildID)
	if !ok {
		return false, nil
	}
	msg := Message{Content: p.Content, MentionUsers: p.MentionIDs}
	live, found, err := r.Bot.d.Store.LivePost(ctx, p.RaidID)
	if err != nil {
		return false, err
	}
	if found {
		channel, msg.ReplyTo = live.ChannelID, live.MessageID
	}
	if _, err := r.Bot.d.REST.CreateMessage(ctx, channel, msg); err != nil {
		return false, fmt.Errorf("discordbot: raid ping: %s", describe(err))
	}
	return true, nil
}

// LivePostRef is where a raid's post is up: its channel and message.
type LivePostRef struct {
	ChannelID, MessageID string
}

// LivePost is the raid's post when it is up (created and not gone); found is false otherwise.
func (s *Store) LivePost(ctx context.Context, raidID int64) (LivePostRef, bool, error) {
	var ref LivePostRef
	err := s.Pool.QueryRow(ctx, `select channel_id, message_id from guild_raid_discord_posts
		where raid_id = $1 and not gone and message_id is not null`, raidID).Scan(&ref.ChannelID, &ref.MessageID)
	if errors.Is(err, pgx.ErrNoRows) {
		return LivePostRef{}, false, nil
	}
	if err != nil {
		return LivePostRef{}, false, fmt.Errorf("discordbot: live post: %w", err)
	}
	return ref, true, nil
}
