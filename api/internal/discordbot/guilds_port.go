package discordbot

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// GuildService is what the bot asks of the guilds package. *guilds.Service
// satisfies it; later tasks append methods.
type GuildService interface {
	GuildStanding(ctx context.Context, guildID, userID int64) (notify.GuildStanding, bool, error)
	// SettingsDiscordFor is the settings.discord object a viewer would read.
	SettingsDiscordFor(ctx context.Context, guildID, viewerID int64) (*guilds.SettingsDiscord, error)
}
