// api/internal/guilds/raid_drops.go
//
// What dropped (guild slice 7): the boss corpses the addon's loot window read in a raid arrive in
// the character's export as drops= (api/internal/fs1/drops.go), and every upload stores them for
// the guild the character is a verified member of, one raid_drops row per item of a corpse. Several
// raiders export the same corpse; the corpse key (the client's GUID, else the kill's encounter and
// day) keeps it once, the larger count and the longer winner list standing.
//
// A drop's boss is the loot table's for a table item (the table, not the addon, names a boss), else
// the encounter the addon saw; an item of no table boss with no encounter is trash and is not kept.
// A new drop is filed under its night as it is stored: the guild's planned raid of the drop's zone
// whose match window holds it, else the guild's unflagged log that killed the boss near it (the
// planned raid holding that log when there is one), else none until a match attaches it
// (raid_drops_match.go).
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// Drop time bounds: a corpse older than dropMaxAge or further ahead than dropMaxSkew is a bad clock
// or a stale record and is not stored.
const (
	dropMaxAge  = 30 * 24 * time.Hour
	dropMaxSkew = time.Hour
	// dropKillLead and dropKillTail are how far a corpse may be looted after, and how far before,
	// the logged kill of its boss for the log to be its night.
	dropKillLead = 3 * time.Hour
	dropKillTail = 30 * time.Minute
)

// raidDropRow is one item of one corpse, ready to store.
type raidDropRow struct {
	zone        string
	instanceID  int
	encounterID int64
	itemID      int
	count       int
	winners     []string
	method      string
	corpseKey   string
	droppedAt   time.Time
}

// RecordExportDrops stores the drops= section of characterKey's export for the guild the character
// is a verified member of, in tx. An export without the section, or a character in no guild or not
// yet approved, stores nothing.
func RecordExportDrops(ctx context.Context, tx pgx.Tx, userID int64, characterKey, export string, now time.Time) error {
	d, ok := fs1.Decode(export)
	if !ok || len(d.Drops) == 0 {
		return nil
	}
	var guildID int64
	err := tx.QueryRow(ctx, `select guild_id from guild_characters where character_key = $1 and verified_at is not null`,
		characterKey).Scan(&guildID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("guilds: record drops: %w", err)
	}
	for _, row := range dropRows(d.Drops, now) {
		if err := storeDrop(ctx, tx, guildID, userID, characterKey, row); err != nil {
			return err
		}
	}
	return nil
}

// dropRows is every storable item of corpses: within the time bounds, with a boss.
func dropRows(corpses []fs1.DropCorpse, now time.Time) []raidDropRow {
	var out []raidDropRow
	for _, c := range corpses {
		at := time.Unix(c.At, 0).UTC()
		if at.Before(now.Add(-dropMaxAge)) || at.After(now.Add(dropMaxSkew)) {
			continue
		}
		for _, it := range c.Items {
			encounterID := dropBoss(it.ItemID, c.EncounterID)
			if encounterID == 0 {
				continue
			}
			out = append(out, raidDropRow{zone: bossZone(encounterID), instanceID: c.InstanceID, encounterID: encounterID,
				itemID: it.ItemID, count: it.Count, winners: it.Winners, method: c.Method,
				corpseKey: corpseKey(c.Source, encounterID, at), droppedAt: at})
		}
	}
	return out
}

// dropBoss is the encounter an item dropped from: the loot table's for a table item, else the one
// the addon saw (0 for none).
func dropBoss(itemID int, seen int64) int64 {
	if enc := dropEncounter(itemID); enc != 0 {
		return enc
	}
	return seen
}

// bossZone is the zone key of encounterID's table boss, "" for one the table does not hold.
func bossZone(encounterID int64) string {
	b, _ := lootBossByEncounter(encounterID)
	return b.RaidKey
}

// corpseKey is the corpse's GUID, else its kill: the boss and the UTC day it was looted.
func corpseKey(source string, encounterID int64, at time.Time) string {
	if source != "" {
		return source
	}
	return "kill:" + strconv.FormatInt(encounterID, 10) + ":" + at.Format(exportDateLayout)
}

// storeDrop writes one row: a new drop under its night, an exported one again kept with the larger
// count, the longer winner list and the first known loot method.
func storeDrop(ctx context.Context, tx pgx.Tx, guildID, userID int64, characterKey string, row raidDropRow) error {
	night, err := dropNight(ctx, tx, guildID, row)
	if err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `
		insert into raid_drops (guild_id, raid_id, report_id, zone, instance_id, encounter_id, item_id, count, winners,
			loot_method, corpse_key, dropped_at, reported_by, reported_character)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		on conflict (guild_id, encounter_id, item_id, corpse_key) do update set
		  count = greatest(raid_drops.count, excluded.count),
		  winners = case when cardinality(excluded.winners) > cardinality(raid_drops.winners)
		                 then excluded.winners else raid_drops.winners end,
		  loot_method = case when raid_drops.loot_method = '' then excluded.loot_method else raid_drops.loot_method end`,
		guildID, night.raidID, night.reportID, row.zone, row.instanceID, row.encounterID, row.itemID, row.count,
		row.winners, row.method, row.corpseKey, row.droppedAt, userID, characterKey); err != nil {
		return fmt.Errorf("guilds: record drop: %w", err)
	}
	return nil
}

// dropNight is the night a new drop is filed under: the planned raid of its zone whose match window
// holds it, else the log that killed its boss near it (or the planned raid holding that log), else none.
func dropNight(ctx context.Context, tx pgx.Tx, guildID int64, row raidDropRow) (awardNight, error) {
	var night awardNight
	err := tx.QueryRow(ctx, `
		select id from guild_raids
		where guild_id = $1 and zone = $2 and status <> 'cancelled'
		  and $3 >= starts_at - make_interval(hours => $4)
		  and $3 <= starts_at + make_interval(mins => duration_minutes) + make_interval(hours => $5)
		order by abs(extract(epoch from (starts_at - $3))), id limit 1`,
		guildID, row.zone, row.droppedAt, matchLeadHours, matchTailHours).Scan(&night.raidID)
	if err == nil || !errors.Is(err, pgx.ErrNoRows) {
		return night, wrapDropNight(err)
	}
	err = tx.QueryRow(ctx, `
		select (select g.id from guild_raids g
		        where g.report_id = r.id and g.guild_id = $1 and g.zone = $5 and g.status <> 'cancelled'
		        order by g.starts_at desc limit 1), r.id
		from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.kill and f.encounter_id = $2 and `+presence.NotFlagged("r")+`
		  and to_timestamp(f.start_ms / 1000.0) between $3::timestamptz - make_interval(secs => $4)
		                                            and $3::timestamptz + make_interval(secs => $6)
		order by abs(extract(epoch from (to_timestamp(f.start_ms / 1000.0) - $3::timestamptz))) limit 1`,
		guildID, row.encounterID, row.droppedAt, dropKillLead.Seconds(), row.zone, dropKillTail.Seconds()).
		Scan(&night.raidID, &night.reportID)
	if errors.Is(err, pgx.ErrNoRows) {
		return awardNight{}, nil
	}
	if night.raidID != nil {
		night.reportID = nil
	}
	return night, wrapDropNight(err)
}

func wrapDropNight(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("guilds: drop night: %w", err)
}
