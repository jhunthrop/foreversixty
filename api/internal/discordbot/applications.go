package discordbot

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// GrantGuildRole implements guilds.ApplicationBot: it gives roleID, in the guild's own Discord
// server, to the account's linked Discord user. No install (or the bot removed in Discord) is
// OutcomeNoBot, no linked account OutcomeNotLinked, and Discord refusing (the user is not in the
// server, the role is above the bot's, the bot lacks Manage Roles) OutcomeDiscordRefuse.
func (s *Service) GrantGuildRole(ctx context.Context, guildID, userID int64, roleID string) (string, error) {
	in, found, err := s.d.Store.InstallFor(ctx, guildID)
	if err != nil {
		return "", err
	}
	if !found || in.RemovedInDiscord {
		return guilds.OutcomeNoBot, nil
	}
	link, linked, err := s.d.Store.LinkByUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if !linked {
		return guilds.OutcomeNotLinked, nil
	}
	if err := s.d.REST.AddMemberRole(ctx, in.ServerID, link.DiscordUserID, roleID); err != nil {
		return refusedOr(err)
	}
	return guilds.OutcomeGranted, nil
}

// DirectMessage implements guilds.ApplicationBot: it DMs text to the account's linked Discord
// user. The caller sends only what the player opted in to.
func (s *Service) DirectMessage(ctx context.Context, userID int64, text string) (string, error) {
	link, linked, err := s.d.Store.LinkByUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if !linked {
		return guilds.OutcomeNotLinked, nil
	}
	channelID, err := s.d.REST.OpenDM(ctx, link.DiscordUserID)
	if err != nil {
		return refusedOr(err)
	}
	if _, err := s.d.REST.CreateMessage(ctx, channelID, Message{Content: text}); err != nil {
		return refusedOr(err)
	}
	return guilds.OutcomeSent, nil
}

// refusedOr reads Discord's own answer as a refusal and anything else (a timeout, a transport
// failure) as an error.
func refusedOr(err error) (string, error) {
	var api *APIError
	if errors.As(err, &api) {
		return guilds.OutcomeDiscordRefuse, nil
	}
	return "", fmt.Errorf("discordbot: %w", err)
}

var _ guilds.ApplicationBot = (*Service)(nil)
