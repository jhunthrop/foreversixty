// api/internal/guilds/raid_badge_rows.go
//
// The Raid tab's guild rows on the raid-badge read: Guild sims shows the viewer's main's place on
// the guild's sims ladder, and Loot history what the guild's last done raid gave out.
package guilds

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// SimsRanker reads a character's place on its guild's sims ladder (activity.Home, the same
// ladder Home's Guild sims block reads); ok is false when it has none there.
type SimsRanker interface {
	GuildSimsRank(ctx context.Context, guildID int64, characterKey string) (position, of int, ok bool, err error)
}

// SimsRank is the viewer's main's place on the guild's sims ladder: 2 of 7.
type SimsRank struct {
	Position int `json:"position"`
	Of       int `json:"of"`
}

// LastRaidLoot is how many items the guild's most recent done raid gave out.
type LastRaidLoot struct {
	Count  int   `json:"count"`
	RaidID int64 `json:"raid_id"`
}

// fillBadgeRows sets the badge's sims rank and last raid loot.
func (s *Service) fillBadgeRows(ctx context.Context, guildID, viewerID int64, out *RaidBadge) error {
	rank, err := s.simsRankOf(ctx, guildID, viewerID)
	if err != nil {
		return err
	}
	out.SimsRank = rank
	out.LastRaidLoot, err = s.Store.lastRaidLoot(ctx, guildID)
	return err
}

// simsRankOf is viewerID's main's place on the guild's ladder; nil without a ranker, a main
// in the guild, or a place on the ladder.
func (s *Service) simsRankOf(ctx context.Context, guildID, viewerID int64) (*SimsRank, error) {
	if s.SimsRanks == nil {
		return nil, nil
	}
	key, ok, err := s.Store.guildMainOf(ctx, guildID, viewerID)
	if err != nil || !ok {
		return nil, err
	}
	pos, of, ok, err := s.SimsRanks.GuildSimsRank(ctx, guildID, key)
	if err != nil || !ok {
		return nil, err
	}
	return &SimsRank{Position: pos, Of: of}, nil
}

// guildMainOf is userID's main in the guild: the account's marked main when it is verified
// there, else its highest level verified character there (by key on a tie).
func (s *Store) guildMainOf(ctx context.Context, guildID, userID int64) (string, bool, error) {
	var key string
	err := s.Pool.QueryRow(ctx, `
		select gc.character_key
		from guild_characters gc
		join characters c on c.key = gc.character_key and c.user_id = $2
		left join users u on u.id = $2
		where gc.guild_id = $1 and gc.verified_at is not null
		order by (gc.character_key = u.main_character_key) is true desc, c.level desc nulls last, gc.character_key
		limit 1`, guildID, userID).Scan(&key)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("guilds: main of %d: %w", userID, err)
	}
	return key, true, nil
}

// lastRaidLoot counts what the guild's most recent done raid gave out: its awards, plus the
// copies of each recorded drop no award on the raid names. Nil with no done raid.
func (s *Store) lastRaidLoot(ctx context.Context, guildID int64) (*LastRaidLoot, error) {
	var out LastRaidLoot
	err := s.Pool.QueryRow(ctx, `
		with last as (
		  select id from guild_raids where guild_id = $1 and status = 'done' order by starts_at desc, id desc limit 1
		)
		select last.id,
		       (select count(*) from loot_awards a where a.raid_id = last.id)
		       + coalesce((select sum(d.count) from raid_drops d where d.raid_id = last.id
		           and not exists (select 1 from loot_awards a where a.raid_id = last.id and a.item_id = d.item_id)), 0)
		from last`, guildID).Scan(&out.RaidID, &out.Count)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: last raid loot: %w", err)
	}
	return &out, nil
}
