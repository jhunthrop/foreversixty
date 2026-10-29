// api/internal/guilds/discord.go
//
// The guild's Discord channel as this package sees it (design/specs/2026-10-26-guild-raids.md
// §3.8): the standing notify's channel routes ask for, and the channel state the settings
// payload and the raid view carry. notify owns the channel itself.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// DiscordChannels is the guild channel as notify keeps it; nil means the deployment has
// no guild channel (connect off).
type DiscordChannels interface {
	GuildConnectEnabled() bool
	GuildChannel(ctx context.Context, guildID int64) (notify.GuildWebhook, bool, error)
	// SetGuildDiscordDefaults remembers the raid form's two boxes, "Post to Discord" and
	// "Remind the day before".
	SetGuildDiscordDefaults(ctx context.Context, guildID int64, post, remind bool) error
}

// SettingsDiscord is settings.discord: present for a verified officer when the deployment
// has connect on, absent otherwise.
type SettingsDiscord struct {
	Available      bool    `json:"available"`
	Connected      bool    `json:"connected"`
	ChannelID      *string `json:"channel_id"`
	DiscordGuildID *string `json:"discord_guild_id"`
	DisabledReason *string `json:"disabled_reason"`
	LastError      *string `json:"last_error"`
	Failures       int     `json:"failures"`
	PostDefault    bool    `json:"post_default"`
	// BotAvailable is whether this deployment runs the Forever Sixty bot.
	BotAvailable bool `json:"bot_available"`
	// Bot is the bot installed in this guild's server; null when none is.
	Bot *SettingsBot `json:"bot"`
	// BotRefusal is read once: why the officer's last install was refused.
	BotRefusal *BotRefusal `json:"bot_refusal"`
}

// SettingsBot is settings.discord.bot.
type SettingsBot struct {
	ServerID         string  `json:"server_id"`
	ServerName       string  `json:"server_name"`
	ChannelID        *string `json:"channel_id"`
	ChannelName      *string `json:"channel_name"`
	LinkedMembers    int     `json:"linked_members"`
	Members          int     `json:"members"`
	Failures         int     `json:"failures"`
	LastError        *string `json:"last_error"`
	RemovedInDiscord bool    `json:"removed_in_discord"`
	// ChannelUnusable is null while the bot can post in its channel.
	ChannelUnusable *ChannelUnusable `json:"channel_unusable"`
}

// ChannelUnusable is a channel the bot is in the server for but cannot post in: Reason
// "deleted" or "permissions"; Missing the lacking subset of View Channel, Send Messages and
// Embed Links, in that order, empty when deleted.
type ChannelUnusable struct {
	Reason  string   `json:"reason"`
	Missing []string `json:"missing"`
}

// BotRefusal names the guild that already holds the server the officer tried to install into.
type BotRefusal struct {
	GuildName string `json:"guild_name"`
}

// DiscordBot is the Forever Sixty bot as the settings read sees it; nil on Service means
// the deployment has no bot.
type DiscordBot interface {
	BotAvailable() bool
	// SettingsBot is the guild's install and the viewer's one-time install refusal.
	SettingsBot(ctx context.Context, guildID, viewerID int64) (*SettingsBot, *BotRefusal, error)
	// Destination is whether the bot posts the guild's raids (installed with a channel, not
	// removed in Discord, the channel usable), the install's remembered boxes whenever there
	// is an install, and the channel name when active; the zero value without an install.
	Destination(ctx context.Context, guildID int64) (BotDestination, error)
	// SetDiscordDefaults remembers the raid form's two boxes on the guild's install.
	SetDiscordDefaults(ctx context.Context, guildID int64, post, remind bool) error
}

// BotDestination is the bot as the guild's raid destination.
type BotDestination struct {
	Active, PostDefault, RemindDefault bool
	ChannelName                        string
}

// discordOn reports whether the guild channel exists in this deployment.
func (s *Service) discordOn() bool {
	return s.Discord != nil && s.Discord.GuildConnectEnabled()
}

// GuildPagePath is the guild page's site path, the web's guildHref: region, ruleset and
// the name slug, the slug escaped as one path segment.
func GuildPagePath(g Guild) string {
	return "/guild/" + strings.ToLower(g.Region) + "/" + strings.ToLower(g.Ruleset) + "/" +
		url.PathEscape(character.Slug(g.Name))
}

// GuildStanding implements notify.GuildAccess with this package's own rules: a verified
// officer or leader (verifiedOfficerOrLeader's), the guild plan (requireGuildPlan's,
// asked for an officer only) and the claim freeze (freezeCheck's, no moderator bypass).
func (s *Service) GuildStanding(ctx context.Context, guildID, userID int64) (notify.GuildStanding, bool, error) {
	g, err := s.Store.getGuild(ctx, guildID)
	if errors.Is(err, ErrNotFound) {
		return notify.GuildStanding{}, false, nil
	}
	if err != nil {
		return notify.GuildStanding{}, false, fmt.Errorf("guilds: standing %d: load guild: %w", guildID, err)
	}
	rank, verified, err := s.Accounts.GuildRank(ctx, guildID, userID)
	if err != nil {
		return notify.GuildStanding{}, false, fmt.Errorf("guilds: standing %d: rank: %w", guildID, err)
	}
	standing := notify.GuildStanding{
		Officer: verified && (rank == "officer" || rank == "leader"),
		Frozen:  claimState(g, time.Now()).Frozen,
		Name:    g.Name, Path: GuildPagePath(g),
	}
	if standing.Officer && s.Entitlements != nil {
		standing.Plan, _, err = s.Entitlements.CanForGuild(ctx, guildID, userID, entitlements.FeatureOfficerViews)
		if err != nil {
			return notify.GuildStanding{}, false, fmt.Errorf("guilds: standing %d: plan: %w", guildID, err)
		}
	}
	return standing, true, nil
}

// botOn reports whether the deployment runs the bot.
func (s *Service) botOn() bool { return s.Bot != nil && s.Bot.BotAvailable() }

// settingsDiscord is settings.discord for the request's viewer.
func (s *Service) settingsDiscord(r *http.Request, guildID int64) (*SettingsDiscord, error) {
	return s.SettingsDiscordFor(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
}

// SettingsDiscordFor is settings.discord for viewerID: nil unless the webhook channel or
// the bot is on and the viewer is a verified officer or leader (a moderator sees none).
func (s *Service) SettingsDiscordFor(ctx context.Context, guildID, viewerID int64) (*SettingsDiscord, error) {
	if !s.discordOn() && !s.botOn() {
		return nil, nil
	}
	officer, err := s.officerOrLeaderOf(ctx, guildID, viewerID)
	if err != nil || !officer {
		return nil, err
	}
	out := &SettingsDiscord{Available: s.discordOn(), BotAvailable: s.botOn()}
	if out.Available {
		w, connected, err := s.Discord.GuildChannel(ctx, guildID)
		if err != nil {
			return nil, err
		}
		out.Connected = connected
		if connected {
			out.ChannelID, out.DiscordGuildID = &w.ChannelID, w.DiscordGuildID
			out.DisabledReason, out.LastError = w.DisabledReason, w.LastError
			out.Failures, out.PostDefault = w.Failures, w.PostDefault
		}
	}
	if out.BotAvailable {
		if out.Bot, out.BotRefusal, err = s.Bot.SettingsBot(ctx, guildID, viewerID); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// raidDiscord is every raid view's discord block: the guild's Discord destination, not the
// raid's. An active bot is one (the router posts through it), else the webhook channel.
func (s *Service) raidDiscord(ctx context.Context, guildID int64) (RaidDiscord, error) {
	if !s.discordOn() && !s.botOn() {
		return RaidDiscord{}, nil
	}
	if s.botOn() {
		dest, err := s.Bot.Destination(ctx, guildID)
		if err != nil {
			return RaidDiscord{}, err
		}
		if dest.Active {
			out := RaidDiscord{Available: true, Connected: true, PostDefault: dest.PostDefault, RemindDefault: dest.RemindDefault}
			if dest.ChannelName != "" {
				out.ChannelName = &dest.ChannelName
			}
			return out, nil
		}
	}
	if !s.discordOn() {
		return RaidDiscord{Available: true}, nil
	}
	w, connected, err := s.Discord.GuildChannel(ctx, guildID)
	if err != nil {
		return RaidDiscord{}, err
	}
	return RaidDiscord{Available: true, Connected: connected, PostDefault: connected && w.PostDefault,
		RemindDefault: connected && w.RemindDefault}, nil
}

// raidViews is Store.RaidViews with the guild's discord block and each raid's posting
// schedule on every view.
func (s *Service) raidViews(ctx context.Context, guildID, viewerID int64, raids []Raid) ([]RaidView, error) {
	views, err := s.Store.RaidViews(ctx, guildID, viewerID, raids)
	if err != nil {
		return nil, err
	}
	discord, err := s.raidDiscord(ctx, guildID)
	if err != nil {
		return nil, err
	}
	if err := s.scopeAttendance(ctx, guildID, viewerID, views); err != nil {
		return nil, err
	}
	schedules, err := s.Store.SchedulesOf(ctx, raidIDsOf(views))
	if err != nil {
		return nil, err
	}
	out := make([]RaidView, len(views))
	for i, v := range views {
		v.Discord = discord
		v.DiscordPost = discordPostOf(schedules, v.ID)
		out[i] = v
	}
	return out, nil
}

// raidViewOf is raidViews for one raid.
func (s *Service) raidViewOf(ctx context.Context, guildID, viewerID int64, raid Raid) (RaidView, error) {
	views, err := s.raidViews(ctx, guildID, viewerID, []Raid{raid})
	if err != nil {
		return RaidView{}, err
	}
	return views[0], nil
}
