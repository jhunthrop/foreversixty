// api/internal/guilds/raid_drops_read.go
//
// The recorded drops as the council's nights read them (loot_council_nights.go): each drop the
// addon recorded is a drop of its night, open until awarded, with the copies the loot window saw and
// who the loot chat says received them ("Looted by").
package guilds

import (
	"context"
	"fmt"
	"time"
)

// recordedDrop is one stored drop as the nights place it.
type recordedDrop struct {
	raidID    *int64
	reportID  *string
	itemID    int
	count     int
	winners   []string
	droppedAt time.Time
}

// recordedTally is an item's recorded drops on one night: the copies and their winners.
type recordedTally struct {
	count   int
	winners []string
}

// recordedDrops reads the guild's drops of table bosses, oldest first.
func (s *Store) recordedDrops(ctx context.Context, guildID int64) ([]recordedDrop, error) {
	rows, err := s.Pool.Query(ctx, `
		select raid_id, report_id, item_id, count, winners, dropped_at from raid_drops
		where guild_id = $1 and encounter_id = any($2) order by dropped_at, id`, guildID, allLoot.encounters)
	if err != nil {
		return nil, fmt.Errorf("guilds: recorded drops: %w", err)
	}
	defer rows.Close()
	var out []recordedDrop
	for rows.Next() {
		var d recordedDrop
		if err := rows.Scan(&d.raidID, &d.reportID, &d.itemID, &d.count, &d.winners, &d.droppedAt); err != nil {
			return nil, fmt.Errorf("guilds: recorded drops: %w", err)
		}
		out = append(out, d)
	}
	return out, rows.Err()
}

// recordedOn is every item recorded on night, by item id.
func recordedOn(drops []recordedDrop, night awardNight, now time.Time) map[int]recordedTally {
	out := map[int]recordedTally{}
	for _, d := range drops {
		if !onNight(d.raidID, d.reportID, d.droppedAt, night, now) {
			continue
		}
		t := out[d.itemID]
		t.count += d.count
		t.winners = append(t.winners, d.winners...)
		out[d.itemID] = t
	}
	return out
}

// withRecorded is drop with what the addon recorded of it on the night, if anything.
func withRecorded(drop LootNightDrop, tally recordedTally, recorded bool) LootNightDrop {
	drop.LootedBy = []string{}
	if !recorded {
		return drop
	}
	drop.Recorded, drop.Count = true, tally.count
	drop.LootedBy = append(drop.LootedBy, tally.winners...)
	return drop
}
