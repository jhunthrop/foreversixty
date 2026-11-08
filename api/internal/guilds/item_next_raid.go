// api/internal/guilds/item_next_raid.go
//
// The item sheet's "Onyxia's Lair, Wed 9 Dec 8:00 pm": the guild's next raid in the zone whose
// table boss drops the item (spec 2026-11-08-app-flows F5).
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// ItemNextRaid is the guild's next raid in an item's zone.
type ItemNextRaid struct {
	RaidID   int64     `json:"raid_id"`
	ZoneKey  string    `json:"zone_key"`
	StartsAt time.Time `json:"starts_at"`
}

// itemNextRaid is the guild's next planned or locked raid in the zone that drops itemID, not
// over at now, the soonest first; nil for an item no table raid drops or with none planned.
func (s *Store) itemNextRaid(ctx context.Context, guildID int64, itemID int, now time.Time) (*ItemNextRaid, error) {
	zone := bossZone(dropEncounter(itemID))
	if zone == "" {
		return nil, nil
	}
	out := ItemNextRaid{ZoneKey: zone}
	err := s.Pool.QueryRow(ctx, `
		select id, starts_at from guild_raids
		where guild_id = $1 and zone = $2 and status in ('planned', 'locked')
		  and starts_at + make_interval(mins => duration_minutes) > $3
		order by starts_at, id limit 1`, guildID, zone, now).Scan(&out.RaidID, &out.StartsAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: item next raid: %w", err)
	}
	return &out, nil
}
