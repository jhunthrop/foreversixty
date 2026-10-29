package discordbot

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// GuildService is what the bot asks of the guilds package. *guilds.Service
// satisfies it; later tasks append methods.
type GuildService interface {
	OfficerReads
	GuildStanding(ctx context.Context, guildID, userID int64) (notify.GuildStanding, bool, error)
	// SettingsDiscordFor is the settings.discord object a viewer would read.
	SettingsDiscordFor(ctx context.Context, guildID, viewerID int64) (*guilds.SettingsDiscord, error)
	// RaidPostFacts is what a raid's channel post shows (the drainer's PostFacts).
	RaidPostFacts(ctx context.Context, raidID int64) (guilds.RaidPostFacts, bool, error)
	// SignUp is the site's own sign-up, with its per-account limit.
	SignUp(ctx context.Context, cmd guilds.SignupCommand) (guilds.SignupResult, error)
	// SaveRaidSpec stores the spec the actor's own character brings to the guild's raids.
	SaveRaidSpec(ctx context.Context, cmd guilds.RaidSpecCommand) (guilds.RaidSpecResult, error)
	// SignupCandidates is the account's characters in the guild, its main first.
	SignupCandidates(ctx context.Context, guildID, userID int64) ([]guilds.SignupCandidate, error)
	// PlannedRaids is the guild's calendar at now, soonest first.
	PlannedRaids(ctx context.Context, guildID int64, now time.Time) ([]guilds.Raid, error)
	// GetRaid is one raid of the guild; found is false for another guild's raid.
	GetRaid(ctx context.Context, guildID, raidID int64) (guilds.Raid, bool, error)
	// MemberRaidView is the raid as a verified member sees it; found is false for anyone else.
	MemberRaidView(ctx context.Context, guildID, viewerID, raidID int64) (guilds.RaidView, bool, error)
	// NextRaidFor is the guild's next raid with the viewer's own sign-up; nil when none.
	NextRaidFor(ctx context.Context, guildID, viewerID int64, headerKey *string, now time.Time) (*guilds.NextRaidView, error)
	GuildName(ctx context.Context, guildID int64) (string, error)
	// GuildRaidTimezone is the guild's IANA raid zone, "" when never set.
	GuildRaidTimezone(ctx context.Context, guildID int64) (string, error)
	// BotGuildFacts is the guild's name, region, ruleset and faction.
	BotGuildFacts(ctx context.Context, guildID int64) (guilds.BotGuild, error)
	// PublicRosterNames is the guild's verified characters with a public page, by name.
	PublicRosterNames(ctx context.Context, guildID int64, typed string, limit int) ([]string, error)
	// CharacterReadinessFacts is the readiness board's word and sync time for one character.
	CharacterReadinessFacts(ctx context.Context, guildID int64, characterKey string) (guilds.ReadinessFacts, error)
	// ActiveBuild is the client build the site serves; BuildItem one item of its tables.
	ActiveBuild() (string, bool)
	BuildItem(version, classSlug string, itemID int) (trees.Item, bool)
	// BisDataDir is the root the BiS files are read under; BisBuild the build they are read
	// from, the site BiS page's own rule.
	BisDataDir() string
	BisBuild() string
	// RaceFaction is a race name's faction ("Night Elf" is alliance); "" when unknown.
	RaceFaction(raceName string) string
}

// CharacterCards is the public character page as a read: rankings.Service.CharacterCard.
type CharacterCards interface {
	CharacterCard(ctx context.Context, region, ruleset, name string) (rankings.Character, bool, error)
}
