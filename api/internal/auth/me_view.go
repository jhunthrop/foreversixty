package auth

import (
	"context"
	"fmt"
)

// NotificationsBlock is /v1/me's notifications object, keyed by channel so a later channel
// adds a key and changes no field. Discord is nil when nothing is saved.
type NotificationsBlock struct {
	// Connect is true when the one-click Discord connect is configured.
	Connect bool          `json:"connect"`
	Discord *DiscordBlock `json:"discord"`
}

// NotificationsView is the part of the notify service the account read needs. It is an
// interface because notify imports this package.
type NotificationsView interface {
	// DiscordBlock answers the member's saved channel, or nil when none is saved. entitled
	// is their notifications entitlement, which decides whether a kept webhook reads as lapsed.
	DiscordBlock(ctx context.Context, userID int64, entitled bool) (*DiscordBlock, error)
	// ConnectEnabled reports whether the one-click connect is configured.
	ConnectEnabled() bool
}

// Me is the /v1/me body: who you are, what you own, and what you're
// entitled to (spec §1.4).
type Me struct {
	User         User             `json:"user"`
	Characters   []Character      `json:"characters"`
	Guilds       []Guild          `json:"guilds"`
	Entitlements EntitlementsView `json:"entitlements"`
	// BnetImportedAt is when the Battle.net import last ran for this
	// account, omitted when it never has (spec §6).
	BnetImportedAt *string `json:"bnet_imported_at,omitempty"`
	// MainCharacterKey is the account's chosen main, omitted until chosen.
	MainCharacterKey *string `json:"main_character_key,omitempty"`
	// FinishedSims are the caller's unseen finished server runs, newest first, at most three.
	// Always present; empty unless the caller holds the server_sims entitlement.
	FinishedSims []FinishedSim `json:"finished_sims"`
	// ServerSims is omitted when no sims store is configured.
	ServerSims *ServerSimsView `json:"server_sims,omitempty"`
	// Notifications is omitted when the deployment has notifications off.
	Notifications *NotificationsBlock `json:"notifications,omitempty"`
	// DiscordLinkAvailable is true when the Discord bot's account link is configured.
	DiscordLinkAvailable bool `json:"discord_link_available"`
	// DiscordLink is the linked Discord account, nil while unlinked.
	DiscordLink *DiscordLinkBlock `json:"discord_link"`
	// DiscordLinkSuggested names a guild whose bot a link would unlock; only while unlinked.
	DiscordLinkSuggested *string `json:"discord_link_suggested"`
}

// buildNotifications resolves the notifications half of /v1/me; nil when the feature is off.
func (s *Service) buildNotifications(ctx context.Context, userID int64, entitled bool) (*NotificationsBlock, error) {
	if s.Notifications == nil {
		return nil, nil
	}
	discord, err := s.Notifications.DiscordBlock(ctx, userID, entitled)
	if err != nil {
		return nil, fmt.Errorf("auth: me: notifications: %w", err)
	}
	return &NotificationsBlock{Connect: s.Notifications.ConnectEnabled(), Discord: discord}, nil
}

// DiscordLinkBlock is the Discord account an account is linked to.
type DiscordLinkBlock struct {
	DisplayName string `json:"display_name"`
}

// DiscordLinkView is the part of the Discord bot the account read needs. It is an interface
// because the bot's package imports this one.
type DiscordLinkView interface {
	// LinkAvailable reports whether the account link is configured.
	LinkAvailable() bool
	// DiscordLink answers the account's link, or nil when unlinked, and, only while unlinked,
	// the name of a guild whose bot a link would unlock.
	DiscordLink(ctx context.Context, userID int64) (link *DiscordLinkBlock, suggested *string, err error)
}

// discordLinkFields is the three link fields of /v1/me.
type discordLinkFields struct {
	available bool
	block     *DiscordLinkBlock
	suggested *string
}

// buildDiscordLink resolves the link half of /v1/me; all zero when the bot is off.
func (s *Service) buildDiscordLink(ctx context.Context, userID int64) (discordLinkFields, error) {
	if s.DiscordLinks == nil || !s.DiscordLinks.LinkAvailable() {
		return discordLinkFields{}, nil
	}
	block, suggested, err := s.DiscordLinks.DiscordLink(ctx, userID)
	if err != nil {
		return discordLinkFields{}, fmt.Errorf("auth: me: discord link: %w", err)
	}
	return discordLinkFields{available: true, block: block, suggested: suggested}, nil
}
