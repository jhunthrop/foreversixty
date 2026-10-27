package discordbot

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// GuildService is what the bot asks of the guilds package. *guilds.Service
// satisfies it; later tasks append methods.
type GuildService interface {
	GuildStanding(ctx context.Context, guildID, userID int64) (notify.GuildStanding, bool, error)
	// SettingsDiscordFor is the settings.discord object a viewer would read.
	SettingsDiscordFor(ctx context.Context, guildID, viewerID int64) (*guilds.SettingsDiscord, error)
	// RaidPostFacts is what a raid's channel post shows (the drainer's PostFacts).
	RaidPostFacts(ctx context.Context, raidID int64) (guilds.RaidPostFacts, bool, error)
	// SignUp is the site's own sign-up, with its per-account limit.
	SignUp(ctx context.Context, cmd guilds.SignupCommand) (guilds.SignupResult, error)
	// SignupCandidates is the account's characters in the guild, its main first.
	SignupCandidates(ctx context.Context, guildID, userID int64) ([]guilds.SignupCandidate, error)
	// PlannedRaids is the guild's calendar at now, soonest first.
	PlannedRaids(ctx context.Context, guildID int64, now time.Time) ([]guilds.Raid, error)
	// GetRaid is one raid of the guild; found is false for another guild's raid.
	GetRaid(ctx context.Context, guildID, raidID int64) (guilds.Raid, bool, error)
	// NextRaidFor is the guild's next raid with the viewer's own sign-up; nil when none.
	NextRaidFor(ctx context.Context, guildID, viewerID int64, headerKey *string, now time.Time) (*guilds.NextRaidView, error)
	GuildName(ctx context.Context, guildID int64) (string, error)
	// GuildRaidTimezone is the guild's IANA raid zone, "" when never set.
	GuildRaidTimezone(ctx context.Context, guildID int64) (string, error)
}
