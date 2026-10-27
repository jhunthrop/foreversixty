// api/internal/guilds/raids_match_loot.go
//
// Lane G2: loot awarded during a raid joins it when the raid learns its log. An award made live
// has no raid (the raid is not done until its log matches), so every match attaches the guild's
// raid-less awards made in the match window, and a clear detaches what its match attached.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// matchedAward is one award as attaching reads it.
type matchedAward struct {
	id                  int64
	encounterID, itemID int64
	characterKey        string
	awardedBy           *int64
	awardedAt           time.Time
	reportID            *string
	raidID              *int64
	attachedByMatch     bool
}

const matchedAwardColumns = `id, encounter_id, item_id, character_key, awarded_by, awarded_at, report_id, raid_id, attached_by_match`

func scanMatchedAward(row pgx.Row) (matchedAward, error) {
	var a matchedAward
	err := row.Scan(&a.id, &a.encounterID, &a.itemID, &a.characterKey, &a.awardedBy, &a.awardedAt, &a.reportID,
		&a.raidID, &a.attachedByMatch)
	return a, err
}

// attachMatchLoot runs inside a match's transaction, once raidID holds its log: every award of
// the guild with no raid whose awarded_at falls between an hour before the raid's start and six
// hours after the earlier of the log's anchor and the raid's end (a log uploaded days later never
// sweeps a later night's awards) joins the raid, marked attached_by_match. An Onyxia raid only:
// awards are Onyxia's. When the raid already holds an award of the same drop, the newer of the
// two stays on the raid and the older is left with no raid.
func attachMatchLoot(ctx context.Context, tx pgx.Tx, raidID int64) error {
	rows, err := tx.Query(ctx, `
		select `+prefixed("la", matchedAwardColumns)+`
		from guild_raids g
		join reports r on r.id = g.report_id
		join loot_awards la on la.guild_id = g.guild_id and la.raid_id is null
		 and la.awarded_at >= g.starts_at - make_interval(hours => $2)
		 and la.awarded_at <= least(coalesce(r.completed_at, r.created_at),
		                            g.starts_at + make_interval(mins => g.duration_minutes)) + make_interval(hours => $3)
		where g.id = $1 and g.zone = $4
		order by la.awarded_at, la.id
		for update of la`, raidID, matchLeadHours, matchTailHours, onyxiaZoneKey)
	if err != nil {
		return fmt.Errorf("guilds: attach loot: %w", err)
	}
	loose, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (matchedAward, error) { return scanMatchedAward(row) })
	if err != nil {
		return fmt.Errorf("guilds: attach loot: %w", err)
	}
	for _, a := range loose {
		if err := attachAward(ctx, tx, raidID, a); err != nil {
			return err
		}
	}
	return nil
}

// attachAward puts one raid-less award on raidID, minding the one-award-per-drop-per-raid index.
func attachAward(ctx context.Context, tx pgx.Tx, raidID int64, a matchedAward) error {
	held, err := scanMatchedAward(tx.QueryRow(ctx, `
		select `+matchedAwardColumns+` from loot_awards
		where raid_id = $1 and encounter_id = $2 and item_id = $3 for update`, raidID, a.encounterID, a.itemID))
	if errors.Is(err, pgx.ErrNoRows) {
		if _, err := tx.Exec(ctx, `update loot_awards set raid_id = $2, attached_by_match = true where id = $1`,
			a.id, raidID); err != nil {
			return fmt.Errorf("guilds: attach loot: %w", err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("guilds: attach loot: %w", err)
	}
	if !newerAward(a, held) {
		return nil // the raid keeps its newer award; this older one stays raid-less
	}
	return swapAwardOntoRaid(ctx, tx, raidID, a, held)
}

func newerAward(a, b matchedAward) bool {
	return a.awardedAt.After(b.awardedAt) || (a.awardedAt.Equal(b.awardedAt) && a.id > b.id)
}

// swapAwardOntoRaid moves the newer raid-less award loose onto raidID and leaves the raid's
// older award held with no raid. Neither unique index lets one row move while the other still
// stands, so loose is taken out, held moves off, and loose returns under its own id.
func swapAwardOntoRaid(ctx context.Context, tx pgx.Tx, raidID int64, loose, held matchedAward) error {
	if _, err := tx.Exec(ctx, `delete from loot_awards where id = $1`, loose.id); err != nil {
		return fmt.Errorf("guilds: attach loot: %w", err)
	}
	if _, err := tx.Exec(ctx, `update loot_awards set raid_id = null, attached_by_match = false where id = $1`,
		held.id); err != nil {
		return fmt.Errorf("guilds: attach loot: %w", err)
	}
	if _, err := tx.Exec(ctx, `
		insert into loot_awards (id, guild_id, encounter_id, item_id, character_key, awarded_by, awarded_at,
			report_id, raid_id, attached_by_match)
		select $1, guild_id, $2, $3, $4, $5, $6, $7, $8, true from guild_raids where id = $8`,
		loose.id, loose.encounterID, loose.itemID, loose.characterKey, loose.awardedBy, loose.awardedAt,
		loose.reportID, raidID); err != nil {
		return fmt.Errorf("guilds: attach loot: %w", err)
	}
	return nil
}

// detachMatchLoot runs inside a clear's transaction: the awards raidID's match attached go
// back to no raid. An explicit award (filed under the raid itself) stays. An attached award
// whose drop already has another raid-less award stays on the raid rather than lose either:
// the one-raid-less-award-per-drop index has room for one.
func detachMatchLoot(ctx context.Context, tx pgx.Tx, raidID int64) error {
	if _, err := tx.Exec(ctx, `
		update loot_awards la set raid_id = null, attached_by_match = false
		where la.raid_id = $1 and la.attached_by_match
		  and not exists (select 1 from loot_awards o
		                  where o.guild_id = la.guild_id and o.raid_id is null
		                    and o.encounter_id = la.encounter_id and o.item_id = la.item_id)`, raidID); err != nil {
		return fmt.Errorf("guilds: detach loot: %w", err)
	}
	return nil
}

// prefixed qualifies each of a comma-separated column list with alias.
func prefixed(alias, columns string) string {
	cols := strings.Split(columns, ",")
	for i, c := range cols {
		cols[i] = alias + "." + strings.TrimSpace(c)
	}
	return strings.Join(cols, ", ")
}
