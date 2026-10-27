// api/internal/guilds/loot_raid.go
//
// Lane G2: a loot award remembers the raid it dropped in. The body may name a raid (a done raid
// of the guild); without one the award goes to the raid being played now, else to no raid until
// a match attaches it. Re-awarding the same drop in a later raid is a new row.
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
)

// AwardRaidError is an award that names a raid that cannot take it. It is an ErrInvalidAward.
type AwardRaidError struct{ RaidID int64 }

func (e *AwardRaidError) Error() string {
	return fmt.Sprintf("guilds: raid %d is not a done Onyxia raid of this guild", e.RaidID)
}

func (e *AwardRaidError) Is(target error) bool { return target == ErrInvalidAward }

// resolveAwardRaid is the raid an award is filed under: the named done raid of the guild
// (an *AwardRaidError when it is not a done raid of the Onyxia zone), else the raid being
// played now - the guild's Onyxia raid, of any status but cancelled, whose match window (an hour
// before the start to six after the end) holds now, the start nearest now winning - else nil, and
// the award joins its raid when that raid's log matches (attachMatchLoot). A previous done raid
// never takes a default award, so a later raid's award never overwrites an earlier one's.
func (s *Store) resolveAwardRaid(ctx context.Context, guildID int64, named *int64, now time.Time) (*int64, error) {
	if named != nil {
		var ok bool
		err := s.Pool.QueryRow(ctx,
			`select exists(select 1 from guild_raids where id = $1 and guild_id = $2 and zone = $3 and status = 'done')`,
			*named, guildID, onyxiaZoneKey).Scan(&ok)
		if err != nil {
			return nil, fmt.Errorf("guilds: award raid: %w", err)
		}
		if !ok {
			return nil, &AwardRaidError{RaidID: *named}
		}
		return named, nil
	}
	var id int64
	err := s.Pool.QueryRow(ctx, `
		select id from guild_raids
		where guild_id = $1 and zone = $2 and status <> 'cancelled'
		  and $3 >= starts_at - make_interval(hours => $4)
		  and $3 <= starts_at + make_interval(mins => duration_minutes) + make_interval(hours => $5)
		order by abs(extract(epoch from (starts_at - $3))), id limit 1`,
		guildID, onyxiaZoneKey, now, matchLeadHours, matchTailHours).Scan(&id)
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
