// api/internal/guilds/raids_bot_reads.go
//
// The reads the Discord bot's sign-up flow (spec 2026-10-27-discord-bot §3.3.2) needs beyond
// SignUp and SignupCandidates: one raid of a guild, the guild's calendar, its name and its
// raid timezone. Each is the store's own read; the bot never reaches the store directly.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// GetRaid is the guild's raid by id; found is false when the raid is not this guild's.
func (s *Service) GetRaid(ctx context.Context, guildID, raidID int64) (Raid, bool, error) {
	return s.Store.GetRaid(ctx, guildID, raidID)
}

// PlannedRaids is the guild's calendar at now (Store.PlannedRaids), soonest first.
func (s *Service) PlannedRaids(ctx context.Context, guildID int64, now time.Time) ([]Raid, error) {
	return s.Store.PlannedRaids(ctx, guildID, now)
}

// NextRaidFor is the guild's next raid as viewerID sees it (Store.NextRaidFor); nil when none.
func (s *Service) NextRaidFor(ctx context.Context, guildID, viewerID int64, headerKey *string, now time.Time) (*NextRaidView, error) {
	return s.Store.NextRaidFor(ctx, guildID, viewerID, headerKey, now)
}

// GuildName is the guild's name as its page shows it.
func (s *Service) GuildName(ctx context.Context, guildID int64) (string, error) {
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return "", fmt.Errorf("guilds: guild name: %w", err)
	}
	return g.Name, nil
}

// GuildRaidTimezone is the guild's IANA raid timezone, "" when never set.
func (s *Service) GuildRaidTimezone(ctx context.Context, guildID int64) (string, error) {
	return s.Store.GuildRaidTimezone(ctx, guildID)
}
