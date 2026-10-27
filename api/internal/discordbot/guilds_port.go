package discordbot

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// GuildService is what the bot asks of the guilds package. *guilds.Service
// satisfies it; later tasks append methods.
type GuildService interface {
	GuildStanding(ctx context.Context, guildID, userID int64) (notify.GuildStanding, bool, error)
}
