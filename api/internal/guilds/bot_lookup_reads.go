// api/internal/guilds/bot_lookup_reads.go
//
// The reads the Discord bot's lookups (/char, /bis; spec 2026-10-27-discord-bot §3.3.3) need
// of a guild and of the build data the guild pages already read.
package guilds

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// BotGuild is what the bot's lookups need of a guild: its name, where its characters live
// and its faction ("" when it has never resolved).
type BotGuild struct {
	Name, Region, Ruleset, Faction string
}

// BotGuildFacts is the guild's BotGuild.
func (s *Service) BotGuildFacts(ctx context.Context, guildID int64) (BotGuild, error) {
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return BotGuild{}, fmt.Errorf("guilds: bot guild: %w", err)
	}
	out := BotGuild{Name: g.Name, Region: g.Region, Ruleset: g.Ruleset}
	if g.Faction != nil {
		out.Faction = *g.Faction
	}
	return out, nil
}

// PublicRosterNames is the names of the guild's verified characters whose account is not
// anonymised (each has a public character page), containing typed, alphabetical, at most limit.
// The anonymise rule is rankings' (rankings/character_extras.go Store.anonymized: the owning
// account's users.anonymize), restated here as a join; change both together.
func (s *Service) PublicRosterNames(ctx context.Context, guildID int64, typed string, limit int) ([]string, error) {
	rows, err := s.Store.Pool.Query(ctx, `
		select c.name from guild_characters gc
		join characters c on c.key = gc.character_key
		left join users u on u.id = c.user_id
		where gc.guild_id = $1 and gc.verified_at is not null and c.name <> ''
		  and coalesce(u.anonymize, false) = false
		  and strpos(lower(c.name), lower($2)) > 0
		order by lower(c.name) limit $3`, guildID, strings.TrimSpace(typed), limit)
	if err != nil {
		return nil, fmt.Errorf("guilds: roster names: %w", err)
	}
	names, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, fmt.Errorf("guilds: roster names: %w", err)
	}
	return names, nil
}

// ActiveBuild is the active client build (trees.Data.Active), the one the site serves.
func (s *Service) ActiveBuild() (string, bool) {
	if s.Store.Trees == nil {
		return "", false
	}
	b, ok := s.Store.Trees.Active()
	if !ok {
		return "", false
	}
	return b.Version, true
}

// BuildItem is Store.ItemIn: one item of a build's per-class table.
func (s *Service) BuildItem(version, classSlug string, itemID int) (trees.Item, bool) {
	return s.Store.ItemIn(version, classSlug, itemID)
}

// RaceFaction is a race's faction by its display name ("Night Elf", the form the character
// page's race carries), read from the BiS build's races.json (BisBuild, as factionForRace
// reads it); "" when unknown or neutral.
func (s *Service) RaceFaction(raceName string) string {
	if raceName == "" || s.Store.Trees == nil {
		return ""
	}
	b, ok := s.Store.Trees.Build(s.Store.BisBuild())
	if !ok {
		return ""
	}
	for _, r := range b.Races() {
		if strings.EqualFold(r.Name, raceName) && (r.Faction == FactionAlliance || r.Faction == FactionHorde) {
			return r.Faction
		}
	}
	return ""
}

// BisDataDir is the root the BiS files are read under (data/builds).
func (s *Service) BisDataDir() string { return s.Store.DataDir }
