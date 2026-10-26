// api/internal/guilds/loot_raid.go
//
// Lane G2: a loot award remembers the raid it dropped in. The body may name a raid (a done raid
// of the guild); without one the award goes to the guild's most recent done Onyxia raid of the
// last week, else to no raid. Re-awarding the same drop in a later raid is a new row.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	// onyxiaZoneKey is the curated zone every award's encounter belongs to.
	onyxiaZoneKey = "raid:onyxias-lair"
	// awardDefaultRaidWindow is how far back a done raid still takes an award with no raid named.
	awardDefaultRaidWindow = 7 * 24 * time.Hour
)

// resolveAwardRaid is the raid an award is filed under: the named done raid of the guild
// (ErrInvalidAward when it is not one), else the default, else nil.
func (s *Store) resolveAwardRaid(ctx context.Context, guildID int64, named *int64, now time.Time) (*int64, error) {
	if named != nil {
		var ok bool
		err := s.Pool.QueryRow(ctx,
			`select exists(select 1 from guild_raids where id = $1 and guild_id = $2 and status = 'done')`,
			*named, guildID).Scan(&ok)
		if err != nil {
			return nil, fmt.Errorf("guilds: award raid: %w", err)
		}
		if !ok {
			return nil, ErrInvalidAward
		}
		return named, nil
	}
	var id int64
	err := s.Pool.QueryRow(ctx, `
		select id from guild_raids
		where guild_id = $1 and zone = $2 and status = 'done' and starts_at >= $3 and starts_at <= $4
		order by starts_at desc, id desc limit 1`,
		guildID, onyxiaZoneKey, now.Add(-awardDefaultRaidWindow), now).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: award raid: %w", err)
	}
	return &id, nil
}

// upsertAward writes the award, replacing the earlier one for the same drop in the same raid
// (or, with no raid, the same drop with no raid).
func (s *Store) upsertAward(ctx context.Context, guildID, encounterID int64, itemID int, characterKey string, awardedBy int64, raidID *int64) (int64, error) {
	target := `(guild_id, encounter_id, item_id) where raid_id is null`
	if raidID != nil {
		target = `(guild_id, raid_id, encounter_id, item_id)`
	}
	var id int64
	err := s.Pool.QueryRow(ctx, `
		insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, raid_id)
		values ($1, $2, $3, $4, $5, $6)
		on conflict `+target+` do update set
		  character_key = excluded.character_key, awarded_by = excluded.awarded_by,
		  awarded_at = now(), report_id = null
		returning id`, guildID, encounterID, itemID, characterKey, awardedBy, raidID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("guilds: create loot award: %w", err)
	}
	return id, nil
}
