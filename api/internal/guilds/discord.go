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

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// DiscordChannels is the guild channel as notify keeps it; nil means the deployment has
// no guild channel (connect off).
type DiscordChannels interface {
	GuildConnectEnabled() bool
	GuildChannel(ctx context.Context, guildID int64) (notify.GuildWebhook, bool, error)
	// SetGuildPostDefault remembers the raid form's "Post to Discord" box.
	SetGuildPostDefault(ctx context.Context, guildID int64, on bool) error
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

// settingsDiscord is settings.discord for the request's viewer: nil unless the channel
// is on and the viewer is a verified officer or leader (a moderator sees none).
func (s *Service) settingsDiscord(r *http.Request, guildID int64) (*SettingsDiscord, error) {
	if !s.discordOn() {
		return nil, nil
	}
	officer, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil || !officer {
		return nil, err
	}
	w, connected, err := s.Discord.GuildChannel(r.Context(), guildID)
	if err != nil {
		return nil, err
	}
	out := &SettingsDiscord{Available: true, Connected: connected}
	if connected {
		out.ChannelID, out.DiscordGuildID = &w.ChannelID, w.DiscordGuildID
		out.DisabledReason, out.LastError = w.DisabledReason, w.LastError
		out.Failures, out.PostDefault = w.Failures, w.PostDefault
	}
	return out, nil
}

// raidDiscord is every raid view's discord block: the guild's channel, not the raid's.
func (s *Service) raidDiscord(ctx context.Context, guildID int64) (RaidDiscord, error) {
	if !s.discordOn() {
		return RaidDiscord{}, nil
	}
	w, connected, err := s.Discord.GuildChannel(ctx, guildID)
	if err != nil {
		return RaidDiscord{}, err
	}
	return RaidDiscord{Available: true, Connected: connected, PostDefault: connected && w.PostDefault}, nil
}

// raidViews is Store.RaidViews with the guild's discord block on every view.
func (s *Service) raidViews(ctx context.Context, guildID, viewerID int64, raids []Raid) ([]RaidView, error) {
	views, err := s.Store.RaidViews(ctx, guildID, viewerID, raids)
	if err != nil {
		return nil, err
	}
	discord, err := s.raidDiscord(ctx, guildID)
	if err != nil {
		return nil, err
	}
	out := make([]RaidView, len(views))
	for i, v := range views {
		v.Discord = discord
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
